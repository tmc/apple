//go:build darwin

package model

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"unsafe"

	"github.com/tmc/apple/coregraphics"
	appleiosurface "github.com/tmc/apple/iosurface"
	xane "github.com/tmc/apple/x/ane"
	xanetelemetry "github.com/tmc/apple/x/ane/telemetry"
)

const defaultQoS = uint32(21)

var compileRuntime struct {
	mu sync.Mutex
	rt *xane.Client
}

type CompileOptions struct {
	MILText       string
	WeightBlob    []byte
	WeightPath    string
	WeightFiles   []WeightFile
	PackagePath   string
	ModelKey      string
	SharedModel   bool
	QoS           uint32
	PerfStatsMask uint32
}

type WeightFile struct {
	Path string
	Blob []byte
}

type CompileStats struct {
	CompileNS int64
	LoadNS    int64
	TotalNS   int64
}

type EvalStats struct {
	HWExecutionNS uint64
	Metrics       map[string]float64
}

// Kernel adapts github.com/tmc/apple/x/ane for the local model API.
type Kernel struct {
	rt     *xane.Client
	k      *xane.Model
	shared *sharedMILHandle

	inputBytes   []int
	outputBytes  []int
	inputAlloc   [][]byte
	outputAlloc  [][]byte
	inputLayout  []xane.TensorLayout
	outputLayout []xane.TensorLayout
}

// CopyOutputChannelsToInput copies a contiguous block of channels from src output
// into dst input without converting through float32.
//
// Both tensors must have matching element size and spatial shape. This mirrors
// the native io_copy helper used to chain ANE kernels through IOSurfaces.
func CopyOutputChannelsToInput(dst *Kernel, dstInput, dstChannel int, src *Kernel, srcOutput, srcChannel, channels int) error {
	return CopyOutputRangeToInput(dst, dstInput, dstChannel, 0, src, srcOutput, srcChannel, 0, channels, -1)
}

// CopyOutputRangeToInput copies a contiguous width range for a contiguous block
// of channels from src output into dst input without converting through float32.
//
// Offsets and width are expressed in elements along the width axis. A width of
// -1 copies the full logical row width.
func CopyOutputRangeToInput(dst *Kernel, dstInput, dstChannel, dstOffset int, src *Kernel, srcOutput, srcChannel, srcOffset, channels, width int) error {
	if dst == nil || dst.closed() {
		return fmt.Errorf("copy output to input: destination kernel is closed")
	}
	if src == nil || src.closed() {
		return fmt.Errorf("copy output to input: source kernel is closed")
	}
	if channels < 0 {
		return fmt.Errorf("copy output to input: channels=%d must be >= 0", channels)
	}
	if channels == 0 {
		return nil
	}
	if dstInput < 0 || dstInput >= len(dst.inputLayout) {
		return fmt.Errorf("copy output to input: destination input index %d out of range", dstInput)
	}
	if srcOutput < 0 || srcOutput >= len(src.outputLayout) {
		return fmt.Errorf("copy output to input: source output index %d out of range", srcOutput)
	}
	dstLayout := dst.inputLayout[dstInput]
	srcLayout := src.outputLayout[srcOutput]
	if dstLayout.Height != 1 || srcLayout.Height != 1 {
		return fmt.Errorf("copy output to input: height > 1 is not supported")
	}
	if width < 0 {
		if dstOffset != 0 || srcOffset != 0 {
			return fmt.Errorf("copy output to input: full-width copy requires zero offsets")
		}
		if dstLayout.Width != srcLayout.Width {
			return fmt.Errorf("copy output to input: width mismatch dst=%d src=%d", dstLayout.Width, srcLayout.Width)
		}
		width = srcLayout.Width
	}
	if dstLayout.ElemSize != srcLayout.ElemSize {
		return fmt.Errorf("copy output to input: elem size mismatch dst=%d src=%d", dstLayout.ElemSize, srcLayout.ElemSize)
	}
	if dstOffset < 0 || dstOffset+width > dstLayout.Width {
		return fmt.Errorf("copy output to input: destination width range [%d,%d) out of range [0,%d)", dstOffset, dstOffset+width, dstLayout.Width)
	}
	if srcOffset < 0 || srcOffset+width > srcLayout.Width {
		return fmt.Errorf("copy output to input: source width range [%d,%d) out of range [0,%d)", srcOffset, srcOffset+width, srcLayout.Width)
	}
	if srcChannel < 0 || srcChannel+channels > srcLayout.Channels {
		return fmt.Errorf("copy output to input: source channels [%d,%d) out of range [0,%d)", srcChannel, srcChannel+channels, srcLayout.Channels)
	}
	if dstChannel < 0 || dstChannel+channels > dstLayout.Channels {
		return fmt.Errorf("copy output to input: destination channels [%d,%d) out of range [0,%d)", dstChannel, dstChannel+channels, dstLayout.Channels)
	}
	return copySurfaceRange(
		dst.InputSurface(dstInput), dstLayout, dstChannel, dstOffset,
		src.OutputSurface(srcOutput), srcLayout, srcChannel, srcOffset,
		channels, width,
	)
}

func Compile(opts CompileOptions) (*Kernel, error) {
	k, _, err := CompileWithStats(opts)
	return k, err
}

func CompileWithStats(opts CompileOptions) (*Kernel, CompileStats, error) {
	opts = compileOptionsWithDefaults(opts)
	if opts.PackagePath == "" && opts.MILText == "" {
		return nil, CompileStats{}, fmt.Errorf("compile: MILText or PackagePath is required")
	}
	if opts.SharedModel && opts.PackagePath == "" {
		return compileSharedMILWithStats(opts)
	}

	rt, err := acquireCompileRuntime()
	if err != nil {
		return nil, CompileStats{}, err
	}
	k, st, err := rt.CompileWithStats(xaneCompileOptions(opts))
	if err != nil {
		return nil, CompileStats{}, fmt.Errorf("compile: %w", err)
	}

	out := &Kernel{
		rt: rt,
		k:  k,
	}
	if err := out.initIO(); err != nil {
		out.Close()
		return nil, CompileStats{}, err
	}
	runtime.SetFinalizer(out, (*Kernel).Close)
	return out, adaptCompileStats(st), nil
}

func (k *Kernel) InputBytes(i int) int {
	if k == nil || i < 0 || i >= len(k.inputBytes) {
		return 0
	}
	return k.inputBytes[i]
}

func (k *Kernel) NumInputs() int {
	if k == nil || k.closed() {
		return 0
	}
	return len(k.inputLayout)
}

func (k *Kernel) OutputBytes(i int) int {
	if k == nil || i < 0 || i >= len(k.outputBytes) {
		return 0
	}
	return k.outputBytes[i]
}

func (k *Kernel) NumOutputs() int {
	if k == nil || k.closed() {
		return 0
	}
	return len(k.outputLayout)
}

func (k *Kernel) InputSurface(i int) coregraphics.IOSurfaceRef {
	if k == nil || k.closed() || i < 0 || i >= len(k.inputLayout) {
		return 0
	}
	if k.shared != nil {
		return k.shared.inputSurface(i)
	}
	return k.k.InputSurface(i)
}

func (k *Kernel) InputLayout(i int) xane.TensorLayout {
	if k == nil || i < 0 || i >= len(k.inputLayout) {
		return xane.TensorLayout{}
	}
	return k.inputLayout[i]
}

func (k *Kernel) OutputSurface(i int) coregraphics.IOSurfaceRef {
	if k == nil || k.closed() || i < 0 || i >= len(k.outputLayout) {
		return 0
	}
	if k.shared != nil {
		return k.shared.outputSurface(i)
	}
	return k.k.OutputSurface(i)
}

func (k *Kernel) OutputLayout(i int) xane.TensorLayout {
	if k == nil || i < 0 || i >= len(k.outputLayout) {
		return xane.TensorLayout{}
	}
	return k.outputLayout[i]
}

func (k *Kernel) WriteInput(i int, b []byte) error {
	if k == nil || k.closed() {
		return fmt.Errorf("write input: kernel is closed")
	}
	if i < 0 || i >= len(k.inputLayout) {
		return fmt.Errorf("write input: index %d out of range", i)
	}
	if len(b) != k.InputBytes(i) {
		return fmt.Errorf("write input: got %d bytes, want %d", len(b), k.InputBytes(i))
	}
	ref := k.InputSurface(i)
	if ref == 0 {
		return fmt.Errorf("write input: input surface %d is nil", i)
	}
	if buf := k.inputAlloc[i]; buf != nil {
		copy(buf, b)
		if k.shared != nil {
			return writeSurface(ref, buf)
		}
		return k.k.WriteInput(i, buf)
	}
	if k.shared != nil {
		return writeSurface(ref, b)
	}
	return k.k.WriteInput(i, b)
}

func (k *Kernel) ReadOutput(i int, b []byte) error {
	if k == nil || k.closed() {
		return fmt.Errorf("read output: kernel is closed")
	}
	if i < 0 || i >= len(k.outputLayout) {
		return fmt.Errorf("read output: index %d out of range", i)
	}
	if len(b) != k.OutputBytes(i) {
		return fmt.Errorf("read output: got %d bytes, want %d", len(b), k.OutputBytes(i))
	}
	ref := k.OutputSurface(i)
	if ref == 0 {
		return fmt.Errorf("read output: output surface %d is nil", i)
	}
	if buf := k.outputAlloc[i]; buf != nil {
		if k.shared != nil {
			if err := readSurface(ref, buf); err != nil {
				return err
			}
		} else {
			if err := k.k.ReadOutput(i, buf); err != nil {
				return err
			}
		}
		copy(b, buf[:len(b)])
		return nil
	}
	if k.shared != nil {
		return readSurface(ref, b)
	}
	return k.k.ReadOutput(i, b)
}

func (k *Kernel) WriteInputF32(i int, data []float32) error {
	if k == nil || k.closed() {
		return fmt.Errorf("write input f32: kernel is closed")
	}
	if k.shared != nil {
		if i < 0 || i >= len(k.inputLayout) {
			return fmt.Errorf("write input f32: index %d out of range", i)
		}
		layout := k.inputLayout[i]
		if len(data) != layout.LogicalElements() {
			return fmt.Errorf("write input f32: got %d elements, want %d", len(data), layout.LogicalElements())
		}
		return writeStridedF32WithLayout(k.InputSurface(i), data, layout)
	}
	return k.k.WriteInputF32(i, data)
}

func (k *Kernel) ReadOutputF32(i int, data []float32) error {
	if k == nil || k.closed() {
		return fmt.Errorf("read output f32: kernel is closed")
	}
	if k.shared != nil {
		if i < 0 || i >= len(k.outputLayout) {
			return fmt.Errorf("read output f32: index %d out of range", i)
		}
		layout := k.outputLayout[i]
		if len(data) != layout.LogicalElements() {
			return fmt.Errorf("read output f32: got %d elements, want %d", len(data), layout.LogicalElements())
		}
		return readStridedF32WithLayout(k.OutputSurface(i), data, layout)
	}
	return k.k.ReadOutputF32(i, data)
}

func (k *Kernel) WriteInputFP16(i int, data []float32) error {
	if k == nil || k.closed() {
		return fmt.Errorf("write input fp16: kernel is closed")
	}
	if k.shared != nil {
		if i < 0 || i >= len(k.inputLayout) {
			return fmt.Errorf("write input fp16: index %d out of range", i)
		}
		layout := k.inputLayout[i]
		if len(data) != layout.LogicalElements() {
			return fmt.Errorf("write input fp16: got %d elements, want %d", len(data), layout.LogicalElements())
		}
		return writeStridedFP16WithLayout(k.InputSurface(i), data, layout)
	}
	return k.k.WriteInputFP16(i, data)
}

func (k *Kernel) WriteInputFP16Channels(i, channel int, data []float32) error {
	if k == nil || k.closed() {
		return fmt.Errorf("write input fp16 channels: kernel is closed")
	}
	if k.shared != nil {
		if i < 0 || i >= len(k.inputLayout) {
			return fmt.Errorf("write input fp16 channels: index %d out of range", i)
		}
		return writeStridedFP16ChannelsWithLayout(k.InputSurface(i), data, k.inputLayout[i], channel)
	}
	return k.k.WriteInputFP16Channels(i, channel, data)
}

func (k *Kernel) ReadOutputFP16(i int, data []float32) error {
	if k == nil || k.closed() {
		return fmt.Errorf("read output fp16: kernel is closed")
	}
	if k.shared != nil {
		if i < 0 || i >= len(k.outputLayout) {
			return fmt.Errorf("read output fp16: index %d out of range", i)
		}
		layout := k.outputLayout[i]
		if len(data) != layout.LogicalElements() {
			return fmt.Errorf("read output fp16: got %d elements, want %d", len(data), layout.LogicalElements())
		}
		return readStridedFP16WithLayout(k.OutputSurface(i), data, layout)
	}
	return k.k.ReadOutputFP16(i, data)
}

func (k *Kernel) ReadOutputFP16Channels(i, channel int, data []float32) error {
	if k == nil || k.closed() {
		return fmt.Errorf("read output fp16 channels: kernel is closed")
	}
	if k.shared != nil {
		if i < 0 || i >= len(k.outputLayout) {
			return fmt.Errorf("read output fp16 channels: index %d out of range", i)
		}
		return readStridedFP16ChannelsWithLayout(k.OutputSurface(i), data, k.outputLayout[i], channel)
	}
	return k.k.ReadOutputFP16Channels(i, channel, data)
}

func (k *Kernel) Eval() error {
	if k == nil || k.closed() {
		return fmt.Errorf("eval: kernel is closed")
	}
	if k.shared != nil {
		return k.sharedEval()
	}
	return k.k.Eval()
}

func (k *Kernel) EvalWithStats() (EvalStats, error) {
	if k == nil || k.closed() {
		return EvalStats{}, fmt.Errorf("eval with stats: kernel is closed")
	}
	if k.shared != nil {
		return k.sharedEvalWithStats()
	}
	st, err := xanetelemetry.EvalWithStats(k.k)
	if err != nil {
		return EvalStats{}, err
	}
	return EvalStats{
		HWExecutionNS: st.HWExecutionNS,
		Metrics:       evalStatsMetrics(st),
	}, nil
}

// EvalHWExecutionNS executes the kernel and returns only hardware execution
// time. It skips metric-map materialization for the fast training path.
func (k *Kernel) EvalHWExecutionNS() (uint64, error) {
	if k == nil || k.closed() {
		return 0, fmt.Errorf("eval hw execution ns: kernel is closed")
	}
	if k.shared != nil {
		st, err := k.sharedEvalWithStats()
		if err != nil {
			return 0, err
		}
		return st.HWExecutionNS, nil
	}
	st, err := xanetelemetry.EvalWithStats(k.k)
	if err != nil {
		return 0, err
	}
	return st.HWExecutionNS, nil
}

func evalStatsMetrics(st xanetelemetry.EvalStats) map[string]float64 {
	rv := reflect.ValueOf(st)
	rt := rv.Type()
	var metrics map[string]float64
	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		if !field.IsExported() || skipEvalMetricField(field.Name) {
			continue
		}
		val, ok := numericEvalMetric(rv.Field(i))
		if !ok {
			continue
		}
		if metrics == nil {
			metrics = make(map[string]float64)
		}
		metrics[field.Name] = val
	}
	metrics = addEvalStatsBytes(metrics, st)
	metrics = addPerfCounterMetrics(metrics, st.PerfCounters)
	if st.PerfCountersTruncated {
		metrics = addEvalMetric(metrics, "PerfCountersTruncated", 1)
	}
	return metrics
}

func numericEvalMetric(v reflect.Value) (float64, bool) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(v.Uint()), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	default:
		return 0, false
	}
}

func skipEvalMetricField(name string) bool {
	switch name {
	case "HWExecutionNS", "PerfCounterData", "RawStatsData", "PerfCounters", "PerfCountersTruncated":
		return true
	default:
		return false
	}
}

func addEvalStatsBytes(metrics map[string]float64, st xanetelemetry.EvalStats) map[string]float64 {
	if n := len(st.PerfCounterData); n > 0 {
		metrics = addEvalMetric(metrics, "PerfCounterBytes", float64(n))
	}
	if n := len(st.RawStatsData); n > 0 {
		metrics = addEvalMetric(metrics, "RawStatsBytes", float64(n))
	}
	return metrics
}

func addPerfCounterMetrics(metrics map[string]float64, counters []xanetelemetry.PerfCounter) map[string]float64 {
	for _, counter := range counters {
		name := counter.Name
		if name == "" {
			name = fmt.Sprintf("%d", counter.Index)
		}
		metrics = addEvalMetric(metrics, "PerfCounter."+name, float64(counter.Value))
	}
	return metrics
}

func addEvalMetric(metrics map[string]float64, name string, value float64) map[string]float64 {
	if metrics == nil {
		metrics = make(map[string]float64)
	}
	metrics[name] += value
	return metrics
}

func adaptCompileStats(st xane.CompileStats) CompileStats {
	return CompileStats{
		CompileNS: st.CompileNS,
		LoadNS:    st.LoadNS,
		TotalNS:   st.TotalNS,
	}
}

func (k *Kernel) Diagnostics() xanetelemetry.Diagnostics {
	if k == nil || k.k == nil || k.shared != nil || k.closed() {
		return xanetelemetry.Diagnostics{}
	}
	return xanetelemetry.ProbeDiagnostics(k.k)
}

func (k *Kernel) EvalWithSignalEvent(signalPort uint32, signalValue uint64, cfg xane.SharedEventEvalOptions) error {
	if k == nil || k.k == nil || k.shared != nil || k.closed() {
		return fmt.Errorf("eval with signal event: kernel is closed")
	}
	return k.k.EvalWithSignalEvent(signalPort, signalValue, cfg)
}

func (k *Kernel) EvalBidirectional(waitPort uint32, waitValue uint64, signalPort uint32, signalValue uint64, cfg xane.SharedEventEvalOptions) error {
	if k == nil || k.k == nil || k.shared != nil || k.closed() {
		return fmt.Errorf("eval bidirectional: kernel is closed")
	}
	return k.k.EvalBidirectional(waitPort, waitValue, signalPort, signalValue, cfg)
}

func (k *Kernel) Close() {
	if k == nil {
		return
	}
	runtime.SetFinalizer(k, nil)
	if k.shared != nil {
		k.shared.close()
		k.shared = nil
	}
	if k.k != nil {
		_ = k.k.Close()
		k.k = nil
	}
	k.rt = nil
}

func (k *Kernel) closed() bool {
	return k == nil || (k.k == nil && k.shared == nil)
}

func acquireCompileRuntime() (*xane.Client, error) {
	compileRuntime.mu.Lock()
	defer compileRuntime.mu.Unlock()
	if compileRuntime.rt != nil {
		return compileRuntime.rt, nil
	}
	rt, err := xane.Open()
	if err != nil {
		return nil, fmt.Errorf("compile: open runtime: %w", err)
	}
	compileRuntime.rt = rt
	return rt, nil
}

func xaneCompileOptions(opts CompileOptions) xane.CompileOptions {
	xc := xane.CompileOptions{
		QoS:           opts.QoS,
		PerfStatsMask: opts.PerfStatsMask,
	}
	if opts.PackagePath != "" {
		xc.ModelType = xane.ModelTypePackage
		xc.PackagePath = opts.PackagePath
		xc.ModelKey = opts.ModelKey
		return xc
	}
	xc.ModelType = xane.ModelTypeMIL
	xc.MILText = []byte(opts.MILText)
	xc.WeightBlob = opts.WeightBlob
	xc.WeightPath = opts.WeightPath
	if len(opts.WeightFiles) > 0 {
		xc.WeightFiles = make([]xane.WeightFile, len(opts.WeightFiles))
		for i, wf := range opts.WeightFiles {
			xc.WeightFiles[i] = xane.WeightFile{Path: wf.Path, Blob: wf.Blob}
		}
	}
	return xc
}

func compileOptionsWithDefaults(opts CompileOptions) CompileOptions {
	if opts.QoS == 0 {
		opts.QoS = defaultQoS
	}
	if opts.PerfStatsMask == 0 {
		opts.PerfStatsMask = defaultPerfStatsMask()
	}
	return opts
}

func defaultPerfStatsMask() uint32 {
	if s := os.Getenv("ANE_PERF_STATS_MASK"); s != "" {
		if v, err := strconv.ParseUint(s, 0, 32); err == nil {
			return uint32(v)
		}
	}
	if os.Getenv("ANE_BENCH") == "1" {
		return ^uint32(0)
	}
	return 0
}

func (k *Kernel) initIO() error {
	nIn := len(k.inputLayout)
	nOut := len(k.outputLayout)
	if k.k != nil {
		nIn = k.k.NumInputs()
		nOut = k.k.NumOutputs()
	}
	if nIn == 0 || nOut == 0 {
		return fmt.Errorf("compile: compiled model reported %d inputs and %d outputs", nIn, nOut)
	}

	k.inputBytes = make([]int, nIn)
	k.outputBytes = make([]int, nOut)
	k.inputAlloc = make([][]byte, nIn)
	k.outputAlloc = make([][]byte, nOut)
	if len(k.inputLayout) != nIn {
		k.inputLayout = make([]xane.TensorLayout, nIn)
	}
	if len(k.outputLayout) != nOut {
		k.outputLayout = make([]xane.TensorLayout, nOut)
	}

	for i := range nIn {
		layout := k.inputLayout[i]
		if k.k != nil {
			layout = k.k.InputLayout(i)
		}
		if layout.LogicalBytes() <= 0 {
			return fmt.Errorf("compile: input[%d] has invalid logical size %d", i, layout.LogicalBytes())
		}
		k.inputLayout[i] = layout
		k.inputBytes[i] = layout.LogicalBytes()
		alloc := layout.AllocSize()
		if k.k != nil {
			alloc = k.k.InputAllocSize(i)
		}
		if alloc > layout.LogicalBytes() {
			k.inputAlloc[i] = make([]byte, alloc)
		}
	}
	for i := range nOut {
		layout := k.outputLayout[i]
		if k.k != nil {
			layout = k.k.OutputLayout(i)
		}
		if layout.LogicalBytes() <= 0 {
			return fmt.Errorf("compile: output[%d] has invalid logical size %d", i, layout.LogicalBytes())
		}
		k.outputLayout[i] = layout
		k.outputBytes[i] = layout.LogicalBytes()
		alloc := layout.AllocSize()
		if k.k != nil {
			alloc = k.k.OutputAllocSize(i)
		}
		if alloc > layout.LogicalBytes() {
			k.outputAlloc[i] = make([]byte, alloc)
		}
	}
	return nil
}

func copySurfaceRange(
	dstRef coregraphics.IOSurfaceRef,
	dstLayout xane.TensorLayout,
	dstChannel int,
	dstOffset int,
	srcRef coregraphics.IOSurfaceRef,
	srcLayout xane.TensorLayout,
	srcChannel int,
	srcOffset int,
	channels int,
	width int,
) error {
	if channels == 0 {
		return nil
	}
	rowBytes := width * srcLayout.ElemSize
	if rowBytes <= 0 {
		return fmt.Errorf("copy output to input: invalid row bytes %d", rowBytes)
	}
	dstSurf := appleiosurface.IOSurfaceRef(dstRef)
	srcSurf := appleiosurface.IOSurfaceRef(srcRef)
	appleiosurface.IOSurfaceLock(dstSurf, 0, nil)
	appleiosurface.IOSurfaceLock(srcSurf, appleiosurface.KIOSurfaceLockReadOnly, nil)
	defer appleiosurface.IOSurfaceUnlock(srcSurf, appleiosurface.KIOSurfaceLockReadOnly, nil)
	defer appleiosurface.IOSurfaceUnlock(dstSurf, 0, nil)

	dstBase := appleiosurface.IOSurfaceGetBaseAddress(dstSurf)
	srcBase := appleiosurface.IOSurfaceGetBaseAddress(srcSurf)
	if dstBase == nil || srcBase == nil {
		return fmt.Errorf("copy output to input: nil IOSurface base address")
	}
	dstAlloc := dstLayout.AllocSize()
	srcAlloc := srcLayout.AllocSize()
	dstBytes := unsafe.Slice((*byte)(dstBase), dstAlloc)
	srcBytes := unsafe.Slice((*byte)(srcBase), srcAlloc)
	for c := 0; c < channels; c++ {
		dstOff := (dstChannel+c)*dstLayout.PlaneStride + dstOffset*dstLayout.ElemSize
		srcOff := (srcChannel+c)*srcLayout.PlaneStride + srcOffset*srcLayout.ElemSize
		if dstOff < 0 || dstOff+rowBytes > len(dstBytes) {
			return fmt.Errorf("copy output to input: destination offset %d out of range", dstOff)
		}
		if srcOff < 0 || srcOff+rowBytes > len(srcBytes) {
			return fmt.Errorf("copy output to input: source offset %d out of range", srcOff)
		}
		copy(dstBytes[dstOff:dstOff+rowBytes], srcBytes[srcOff:srcOff+rowBytes])
	}
	return nil
}

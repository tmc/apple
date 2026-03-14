//go:build darwin

package model

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	appleiosurface "github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/private/appleneuralengine"
	xane "github.com/tmc/apple/x/ane"
)

var sharedMILCache struct {
	mu sync.Mutex
	m  map[string]*sharedMILProgram
	p  map[string]*sharedMILPending
}

type sharedMILPending struct {
	done chan struct{}
	err  error
}

type sharedMILProgram struct {
	key string

	inMemModel    appleneuralengine.ANEInMemoryModel
	inputLayouts  []xane.TensorLayout
	outputLayouts []xane.TensorLayout
	qos           uint32
	mapMu         sync.Mutex
	mapMode       atomic.Uint32

	owner *xane.Model
	refs  int
}

type sharedMILHandle struct {
	program *sharedMILProgram
	request appleneuralengine.ANERequest
	inputs  []coregraphics.IOSurfaceRef
	outputs []coregraphics.IOSurfaceRef
}

func (h *sharedMILHandle) inputSurface(i int) coregraphics.IOSurfaceRef {
	if h == nil || i < 0 || i >= len(h.inputs) {
		return 0
	}
	return h.inputs[i]
}

func (h *sharedMILHandle) outputSurface(i int) coregraphics.IOSurfaceRef {
	if h == nil || i < 0 || i >= len(h.outputs) {
		return 0
	}
	return h.outputs[i]
}

func compileSharedMILWithStats(opts CompileOptions) (*Kernel, CompileStats, error) {
	key := sharedMILCacheKey(opts)

	sharedMILCache.mu.Lock()
	if sharedMILCache.m == nil {
		sharedMILCache.m = make(map[string]*sharedMILProgram)
	}
	if prog := sharedMILCache.m[key]; prog != nil {
		prog.refs++
		sharedMILCache.mu.Unlock()
		return prog.newCloneKernel()
	}
	if p := sharedMILCache.p[key]; p != nil {
		sharedMILCache.mu.Unlock()
		<-p.done
		if p.err != nil {
			return nil, CompileStats{}, p.err
		}
		sharedMILCache.mu.Lock()
		prog := sharedMILCache.m[key]
		if prog != nil {
			prog.refs++
		}
		sharedMILCache.mu.Unlock()
		if prog == nil {
			return nil, CompileStats{}, fmt.Errorf("compile: shared model cache miss after wait")
		}
		return prog.newCloneKernel()
	}
	if sharedMILCache.p == nil {
		sharedMILCache.p = make(map[string]*sharedMILPending)
	}
	p := &sharedMILPending{done: make(chan struct{})}
	sharedMILCache.p[key] = p
	sharedMILCache.mu.Unlock()

	rt, err := acquireCompileRuntime()
	if err != nil {
		sharedMILCache.mu.Lock()
		delete(sharedMILCache.p, key)
		p.err = err
		close(p.done)
		sharedMILCache.mu.Unlock()
		return nil, CompileStats{}, err
	}
	xk, st, err := rt.CompileWithStats(xaneCompileOptions(opts))
	if err != nil {
		sharedMILCache.mu.Lock()
		delete(sharedMILCache.p, key)
		p.err = fmt.Errorf("compile: %w", err)
		close(p.done)
		sharedMILCache.mu.Unlock()
		return nil, CompileStats{}, fmt.Errorf("compile: %w", err)
	}

	prog := &sharedMILProgram{
		key:           key,
		inMemModel:    appleneuralengine.ANEInMemoryModelFromID(objc.ID(xk.InMemModelObjcID())),
		inputLayouts:  cloneLayoutsFromXANE(xk, true),
		outputLayouts: cloneLayoutsFromXANE(xk, false),
		qos:           opts.QoS,
		owner:         xk,
		refs:          1,
	}

	sharedMILCache.mu.Lock()
	delete(sharedMILCache.p, key)
	if existing := sharedMILCache.m[key]; existing != nil {
		existing.refs++
		p.err = nil
		close(p.done)
		sharedMILCache.mu.Unlock()
		_ = xk.Close()
		return existing.newCloneKernel()
	}
	sharedMILCache.m[key] = prog
	p.err = nil
	close(p.done)
	sharedMILCache.mu.Unlock()
	k, _, err := prog.newCloneKernel()
	if err != nil {
		return nil, CompileStats{}, err
	}
	k.rt = rt
	runtime.SetFinalizer(k, (*Kernel).Close)
	return k, adaptCompileStats(st), nil
}

func (p *sharedMILProgram) newCloneKernel() (*Kernel, CompileStats, error) {
	request, inputs, outputs, err := createRequestAndSurfaces(p.inputLayouts, p.outputLayouts)
	if err != nil {
		p.release()
		return nil, CompileStats{}, err
	}
	if err := p.mapRequest(request); err != nil {
		p.release()
		return nil, CompileStats{}, err
	}
	k := &Kernel{
		shared: &sharedMILHandle{
			program: p,
			request: request,
			inputs:  inputs,
			outputs: outputs,
		},
		inputLayout:  append([]xane.TensorLayout(nil), p.inputLayouts...),
		outputLayout: append([]xane.TensorLayout(nil), p.outputLayouts...),
	}
	if err := k.initIO(); err != nil {
		k.Close()
		return nil, CompileStats{}, err
	}
	return k, CompileStats{}, nil
}

func (p *sharedMILProgram) mapRequest(request appleneuralengine.ANERequest) error {
	switch p.mapMode.Load() {
	case 1:
		if ok, err := p.inMemModel.MapIOSurfacesWithRequestCacheInferenceError(request, false); err == nil && ok {
			return nil
		}
	case 2:
		if ok, err := p.inMemModel.MapIOSurfacesWithRequestCacheInferenceError(request, true); err == nil && ok {
			return nil
		}
	}
	return p.mapRequestSlow(request)
}

func (p *sharedMILProgram) mapRequestSlow(request appleneuralengine.ANERequest) error {
	p.mapMu.Lock()
	defer p.mapMu.Unlock()

	mode := p.mapMode.Load()
	try := []bool{true, false}
	switch mode {
	case 1:
		try = []bool{false, true}
	case 2:
		try = []bool{true, false}
	}
	var lastErr error
	for _, cacheInference := range try {
		ok, err := p.inMemModel.MapIOSurfacesWithRequestCacheInferenceError(request, cacheInference)
		if err == nil && ok {
			if cacheInference {
				p.mapMode.Store(2)
			} else {
				p.mapMode.Store(1)
			}
			return nil
		}
		if err != nil {
			lastErr = err
		}
	}
	if lastErr != nil {
		return fmt.Errorf("map IOSurfaces failed: %w", lastErr)
	}
	return fmt.Errorf("map IOSurfaces failed")
}

func (h *sharedMILHandle) eval() error {
	if h == nil || h.program == nil {
		return fmt.Errorf("eval: kernel is closed")
	}
	ok, err := h.program.inMemModel.EvaluateWithQoSOptionsRequestError(h.program.qos, nil, h.request)
	if err == nil && ok {
		return nil
	}
	return fmt.Errorf("eval: %w", err)
}

func (h *sharedMILHandle) close() {
	if h == nil || h.program == nil {
		return
	}
	prog := h.program
	prog.mapMu.Lock()
	prog.inMemModel.UnmapIOSurfacesWithRequest(h.request)
	prog.mapMu.Unlock()
	h.program = nil
	h.request = appleneuralengine.ANERequest{}
	h.inputs = nil
	h.outputs = nil
	prog.release()
}

func (p *sharedMILProgram) release() {
	sharedMILCache.mu.Lock()
	defer sharedMILCache.mu.Unlock()
	if p.refs > 0 {
		p.refs--
	}
}

func sharedMILCacheKey(opts CompileOptions) string {
	h := sha256.New()
	writeHashU32(h, opts.QoS)
	writeHashU32(h, opts.PerfStatsMask)
	h.Write([]byte(opts.MILText))
	if len(opts.WeightBlob) > 0 {
		h.Write([]byte{0})
		h.Write([]byte(opts.WeightPath))
		h.Write(opts.WeightBlob)
	}
	for _, wf := range opts.WeightFiles {
		h.Write([]byte{1})
		h.Write([]byte(wf.Path))
		h.Write(wf.Blob)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func writeHashU32(h interface{ Write([]byte) (int, error) }, v uint32) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], v)
	_, _ = h.Write(buf[:])
}

func cloneLayoutsFromXANE(k *xane.Model, inputs bool) []xane.TensorLayout {
	n := k.NumOutputs()
	if inputs {
		n = k.NumInputs()
	}
	out := make([]xane.TensorLayout, n)
	for i := range n {
		if inputs {
			out[i] = k.InputLayout(i)
		} else {
			out[i] = k.OutputLayout(i)
		}
	}
	return out
}

func runtimeNanoNow() int64 {
	return time.Now().UnixNano()
}

func (k *Kernel) sharedEval() error {
	if k == nil || k.shared == nil {
		return fmt.Errorf("eval: kernel is closed")
	}
	return k.shared.eval()
}

func (k *Kernel) sharedEvalWithStats() (EvalStats, error) {
	start := runtimeNanoNow()
	if err := k.sharedEval(); err != nil {
		return EvalStats{}, err
	}
	return EvalStats{HWExecutionNS: uint64(runtimeNanoNow() - start)}, nil
}

func createRequestAndSurfaces(inputLayouts, outputLayouts []xane.TensorLayout) (appleneuralengine.ANERequest, []coregraphics.IOSurfaceRef, []coregraphics.IOSurfaceRef, error) {
	if len(inputLayouts) == 0 || len(outputLayouts) == 0 {
		return appleneuralengine.ANERequest{}, nil, nil, fmt.Errorf("input and output layouts must be specified")
	}

	ioClass := appleneuralengine.GetANEIOSurfaceObjectClass()
	inputs := make([]coregraphics.IOSurfaceRef, len(inputLayouts))
	inputArr := foundation.NewNSMutableArray()
	inputIdxArr := foundation.NewNSMutableArray()
	for i, layout := range inputLayouts {
		ref, err := createSurfaceForLayout(layout)
		if err != nil {
			return appleneuralengine.ANERequest{}, nil, nil, err
		}
		inputs[i] = ref
		inputArr.AddObject(ioClass.ObjectWithIOSurface(ref))
		inputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(i))
	}

	outputs := make([]coregraphics.IOSurfaceRef, len(outputLayouts))
	outputArr := foundation.NewNSMutableArray()
	outputIdxArr := foundation.NewNSMutableArray()
	for i, layout := range outputLayouts {
		ref, err := createSurfaceForLayout(layout)
		if err != nil {
			return appleneuralengine.ANERequest{}, nil, nil, err
		}
		outputs[i] = ref
		outputArr.AddObject(ioClass.ObjectWithIOSurface(ref))
		outputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(i))
	}

	procIdx := foundation.GetNSNumberClass().NumberWithInt(0)
	txnHandle := foundation.GetNSNumberClass().NumberWithUnsignedLongLong(1)
	reqClass := appleneuralengine.GetANERequestClass()
	reqObj := reqClass.RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(
		inputArr, inputIdxArr, outputArr, outputIdxArr, nil, nil, procIdx, nil, txnHandle,
	)
	if reqObj == nil || reqObj.GetID() == 0 {
		reqObj = reqClass.RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndex(
			inputArr, inputIdxArr, outputArr, outputIdxArr, nil, nil, procIdx,
		)
	}
	if reqObj == nil || reqObj.GetID() == 0 {
		return appleneuralengine.ANERequest{}, nil, nil, fmt.Errorf("failed to create request")
	}
	request := appleneuralengine.ANERequestFromID(reqObj.GetID())
	if request.TransactionHandle().GetID() == 0 {
		request.SetTransactionHandle(txnHandle)
	}
	return request, inputs, outputs, nil
}

func createSurfaceForLayout(layout xane.TensorLayout) (coregraphics.IOSurfaceRef, error) {
	if layout.RowStride%64 != 0 {
		return 0, fmt.Errorf("RowStride %d is not 64-byte aligned", layout.RowStride)
	}
	minStride := xane.RowStrideFor(layout.Width, layout.ElemSize)
	if layout.RowStride < minStride {
		return 0, fmt.Errorf("RowStride %d < minimum %d", layout.RowStride, minStride)
	}
	alloc := layout.AllocSize()
	surfHeight := layout.Channels * layout.Height
	keys := [6]unsafe.Pointer{
		cfString("IOSurfaceWidth"),
		cfString("IOSurfaceHeight"),
		cfString("IOSurfaceBytesPerElement"),
		cfString("IOSurfaceBytesPerRow"),
		cfString("IOSurfaceAllocSize"),
		cfString("IOSurfacePixelFormat"),
	}
	values := [6]unsafe.Pointer{
		cfInt(layout.Width),
		cfInt(surfHeight),
		cfInt(layout.ElemSize),
		cfInt(layout.RowStride),
		cfInt(alloc),
		cfInt(0),
	}
	dict := corefoundation.CFDictionaryCreate(0, unsafe.Pointer(&keys[0]), unsafe.Pointer(&values[0]), 6, nil, nil)
	ref := appleiosurface.IOSurfaceCreate(corefoundation.CFDictionaryRef(dict))
	if ref == 0 {
		return 0, fmt.Errorf("failed to create IOSurface")
	}
	return coregraphics.IOSurfaceRef(ref), nil
}

func writeSurface(ref coregraphics.IOSurfaceRef, src []byte) error {
	surfRef := appleiosurface.IOSurfaceRef(ref)
	appleiosurface.IOSurfaceLock(surfRef, 0, nil)
	base := appleiosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		appleiosurface.IOSurfaceUnlock(surfRef, 0, nil)
		return fmt.Errorf("IOSurface base address is nil")
	}
	dst := unsafe.Slice((*byte)(base), len(src))
	copy(dst, src)
	appleiosurface.IOSurfaceUnlock(surfRef, 0, nil)
	return nil
}

func readSurface(ref coregraphics.IOSurfaceRef, dst []byte) error {
	surfRef := appleiosurface.IOSurfaceRef(ref)
	appleiosurface.IOSurfaceLock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
	base := appleiosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		appleiosurface.IOSurfaceUnlock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
		return fmt.Errorf("IOSurface base address is nil")
	}
	src := unsafe.Slice((*byte)(base), len(dst))
	copy(dst, src)
	appleiosurface.IOSurfaceUnlock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
	return nil
}

func writeStridedFP16WithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout xane.TensorLayout) error {
	surfRef := appleiosurface.IOSurfaceRef(ref)
	appleiosurface.IOSurfaceLock(surfRef, 0, nil)
	base := appleiosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		appleiosurface.IOSurfaceUnlock(surfRef, 0, nil)
		return fmt.Errorf("IOSurface base address is nil")
	}
	dst := unsafe.Slice((*byte)(base), layout.AllocSize())
	writeStridedFP16Bytes(dst, data, layout)
	appleiosurface.IOSurfaceUnlock(surfRef, 0, nil)
	return nil
}

func writeStridedFP16ChannelsWithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout xane.TensorLayout, channel int) error {
	channelElems := layout.Height * layout.Width
	if channelElems <= 0 {
		return fmt.Errorf("invalid fp16 layout channel elements %d", channelElems)
	}
	if len(data)%channelElems != 0 {
		return fmt.Errorf("fp16 channel data length %d is not a multiple of channel size %d", len(data), channelElems)
	}
	channels := len(data) / channelElems
	if channel < 0 || channel+channels > layout.Channels {
		return fmt.Errorf("fp16 channel range [%d,%d) out of bounds [0,%d)", channel, channel+channels, layout.Channels)
	}
	sub := layout
	sub.Channels = channels
	surfRef := appleiosurface.IOSurfaceRef(ref)
	appleiosurface.IOSurfaceLock(surfRef, 0, nil)
	base := appleiosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		appleiosurface.IOSurfaceUnlock(surfRef, 0, nil)
		return fmt.Errorf("IOSurface base address is nil")
	}
	dst := unsafe.Slice((*byte)(base), layout.AllocSize())
	writeStridedFP16Bytes(dst[channel*layout.PlaneStride:], data, sub)
	appleiosurface.IOSurfaceUnlock(surfRef, 0, nil)
	return nil
}

func readStridedFP16WithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout xane.TensorLayout) error {
	surfRef := appleiosurface.IOSurfaceRef(ref)
	appleiosurface.IOSurfaceLock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
	base := appleiosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		appleiosurface.IOSurfaceUnlock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
		return fmt.Errorf("IOSurface base address is nil")
	}
	src := unsafe.Slice((*byte)(base), layout.AllocSize())
	readStridedFP16Bytes(data, src, layout)
	appleiosurface.IOSurfaceUnlock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
	return nil
}

func readStridedFP16ChannelsWithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout xane.TensorLayout, channel int) error {
	channelElems := layout.Height * layout.Width
	if channelElems <= 0 {
		return fmt.Errorf("invalid fp16 layout channel elements %d", channelElems)
	}
	if len(data)%channelElems != 0 {
		return fmt.Errorf("fp16 channel data length %d is not a multiple of channel size %d", len(data), channelElems)
	}
	channels := len(data) / channelElems
	if channel < 0 || channel+channels > layout.Channels {
		return fmt.Errorf("fp16 channel range [%d,%d) out of bounds [0,%d)", channel, channel+channels, layout.Channels)
	}
	sub := layout
	sub.Channels = channels
	surfRef := appleiosurface.IOSurfaceRef(ref)
	appleiosurface.IOSurfaceLock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
	base := appleiosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		appleiosurface.IOSurfaceUnlock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
		return fmt.Errorf("IOSurface base address is nil")
	}
	src := unsafe.Slice((*byte)(base), layout.AllocSize())
	readStridedFP16Bytes(data, src[channel*layout.PlaneStride:], sub)
	appleiosurface.IOSurfaceUnlock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
	return nil
}

func writeStridedF32WithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout xane.TensorLayout) error {
	surfRef := appleiosurface.IOSurfaceRef(ref)
	appleiosurface.IOSurfaceLock(surfRef, 0, nil)
	base := appleiosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		appleiosurface.IOSurfaceUnlock(surfRef, 0, nil)
		return fmt.Errorf("IOSurface base address is nil")
	}
	dst := unsafe.Slice((*byte)(base), layout.AllocSize())
	writeStridedF32Bytes(dst, data, layout)
	appleiosurface.IOSurfaceUnlock(surfRef, 0, nil)
	return nil
}

func readStridedF32WithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout xane.TensorLayout) error {
	surfRef := appleiosurface.IOSurfaceRef(ref)
	appleiosurface.IOSurfaceLock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
	base := appleiosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		appleiosurface.IOSurfaceUnlock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
		return fmt.Errorf("IOSurface base address is nil")
	}
	src := unsafe.Slice((*byte)(base), layout.AllocSize())
	readStridedF32Bytes(data, src, layout)
	appleiosurface.IOSurfaceUnlock(surfRef, appleiosurface.KIOSurfaceLockReadOnly, nil)
	return nil
}

func logicalElementLimit(dataLen int, layout xane.TensorLayout) int {
	n := layout.LogicalElements()
	if dataLen < n {
		return dataLen
	}
	return n
}

func writeStridedF32Bytes(dst []byte, data []float32, layout xane.TensorLayout) {
	if layout.Width <= 0 || layout.Height <= 0 || layout.Channels <= 0 {
		return
	}
	limit := logicalElementLimit(len(data), layout)
	if limit == 0 {
		return
	}
	hw := layout.Height * layout.Width
	for c := range layout.Channels {
		for h := range layout.Height {
			srcIdx := c*hw + h*layout.Width
			if srcIdx >= limit {
				return
			}
			n := layout.Width
			if remain := limit - srcIdx; remain < n {
				n = remain
			}
			off := c*layout.PlaneStride + h*layout.RowStride
			if off+n*4 > len(dst) {
				return
			}
			row := unsafe.Slice((*float32)(unsafe.Pointer(&dst[off])), n)
			copy(row, data[srcIdx:srcIdx+n])
		}
	}
}

func readStridedF32Bytes(data []float32, src []byte, layout xane.TensorLayout) {
	if layout.Width <= 0 || layout.Height <= 0 || layout.Channels <= 0 {
		return
	}
	limit := logicalElementLimit(len(data), layout)
	if limit == 0 {
		return
	}
	hw := layout.Height * layout.Width
	for c := range layout.Channels {
		for h := range layout.Height {
			dstIdx := c*hw + h*layout.Width
			if dstIdx >= limit {
				return
			}
			n := layout.Width
			if remain := limit - dstIdx; remain < n {
				n = remain
			}
			off := c*layout.PlaneStride + h*layout.RowStride
			if off+n*4 > len(src) {
				return
			}
			row := unsafe.Slice((*float32)(unsafe.Pointer(&src[off])), n)
			copy(data[dstIdx:dstIdx+n], row)
		}
	}
}

func writeStridedFP16Bytes(dst []byte, data []float32, layout xane.TensorLayout) {
	if layout.Width <= 0 || layout.Height <= 0 || layout.Channels <= 0 {
		return
	}
	limit := logicalElementLimit(len(data), layout)
	if limit == 0 {
		return
	}
	hw := layout.Height * layout.Width
	for c := range layout.Channels {
		for h := range layout.Height {
			srcIdx := c*hw + h*layout.Width
			if srcIdx >= limit {
				return
			}
			n := layout.Width
			if remain := limit - srcIdx; remain < n {
				n = remain
			}
			off := c*layout.PlaneStride + h*layout.RowStride
			if off+n*2 > len(dst) {
				return
			}
			row := unsafe.Slice((*uint16)(unsafe.Pointer(&dst[off])), n)
			for i := range n {
				row[i] = float32ToFP16(data[srcIdx+i])
			}
		}
	}
}

func readStridedFP16Bytes(data []float32, src []byte, layout xane.TensorLayout) {
	if layout.Width <= 0 || layout.Height <= 0 || layout.Channels <= 0 {
		return
	}
	limit := logicalElementLimit(len(data), layout)
	if limit == 0 {
		return
	}
	hw := layout.Height * layout.Width
	for c := range layout.Channels {
		for h := range layout.Height {
			dstIdx := c*hw + h*layout.Width
			if dstIdx >= limit {
				return
			}
			n := layout.Width
			if remain := limit - dstIdx; remain < n {
				n = remain
			}
			off := c*layout.PlaneStride + h*layout.RowStride
			if off+n*2 > len(src) {
				return
			}
			row := unsafe.Slice((*uint16)(unsafe.Pointer(&src[off])), n)
			for i := range n {
				data[dstIdx+i] = fp16ToFloat32(row[i])
			}
		}
	}
}

func float32ToFP16(f float32) uint16 {
	b := math.Float32bits(f)
	sign := (b >> 16) & 0x8000
	exp := int((b>>23)&0xFF) - 127 + 15
	frac := b & 0x7FFFFF
	switch {
	case exp <= 0:
		return uint16(sign)
	case exp >= 31:
		return uint16(sign | 0x7C00)
	default:
		return uint16(sign | uint32(exp)<<10 | (frac >> 13))
	}
}

func fp16ToFloat32(h uint16) float32 {
	sign := uint32(h>>15) & 1
	exp := uint32(h>>10) & 0x1F
	frac := uint32(h) & 0x3FF
	switch {
	case exp == 0:
		if frac == 0 {
			return math.Float32frombits(sign << 31)
		}
		for frac&0x400 == 0 {
			frac <<= 1
			exp--
		}
		exp++
		frac &= 0x3FF
		fallthrough
	case exp < 31:
		return math.Float32frombits(sign<<31 | (exp+127-15)<<23 | frac<<13)
	default:
		return math.Float32frombits(sign<<31 | 0x7F800000 | frac<<13)
	}
}

func cfString(s string) unsafe.Pointer {
	b := append([]byte(s), 0)
	ref := corefoundation.CFStringCreateWithCString(0, &b[0], 0x08000100)
	p := uintptr(ref)
	return *(*unsafe.Pointer)(unsafe.Pointer(&p))
}

func cfInt(v int) unsafe.Pointer {
	val := int64(v)
	ref := corefoundation.CFNumberCreate(0, corefoundation.KCFNumberSInt64Type, unsafe.Pointer(&val))
	p := uintptr(ref)
	return *(*unsafe.Pointer)(unsafe.Pointer(&p))
}

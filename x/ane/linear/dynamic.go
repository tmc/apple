package linear

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/tmc/apple/x/ane/dynamicmatmul"
)

// DynamicExecutor caches compile-once ANE kernels keyed by shape and accepts
// runtime-provided weights on each call.
//
// It preserves the row-major linear API used by Executor:
//   - x is [batch, inDim]
//   - w is [outDim, inDim]
//   - y is [batch, outDim]
type DynamicExecutor struct {
	qos uint32

	mu      sync.Mutex
	kernels map[string]*dynamicCompiledKernel
	stats   Stats
}

type dynamicCompiledKernel struct {
	mu sync.Mutex
	k  *dynamicmatmul.Executor

	weightsIO []float32
}

// NewDynamic creates a dynamic linear executor.
func NewDynamic(opts Options) *DynamicExecutor {
	qos := opts.QoS
	if qos == 0 {
		qos = defaultQoS
	}
	return &DynamicExecutor{
		qos:     qos,
		kernels: make(map[string]*dynamicCompiledKernel),
	}
}

// Stats returns a snapshot of executor counters.
func (e *DynamicExecutor) Stats() Stats {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.stats
}

// Close releases all cached kernels.
func (e *DynamicExecutor) Close() {
	if e == nil {
		return
	}
	e.mu.Lock()
	ks := make([]*dynamicCompiledKernel, 0, len(e.kernels))
	for _, k := range e.kernels {
		ks = append(ks, k)
	}
	e.kernels = make(map[string]*dynamicCompiledKernel)
	e.stats.Kernels = 0
	e.mu.Unlock()

	for _, ck := range ks {
		if ck != nil && ck.k != nil {
			ck.k.Close()
		}
	}
}

// Linear computes x*w^T where x is [batch,inDim] and w is [outDim,inDim].
func (e *DynamicExecutor) Linear(ctx context.Context, x, w []float32, batch, inDim, outDim int) ([]float32, error) {
	out := make([]float32, batch*outDim)
	_, err := e.LinearIntoWithStats(ctx, out, x, w, batch, inDim, outDim)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinearWithStats computes x*w^T and returns per-call execution stats.
func (e *DynamicExecutor) LinearWithStats(ctx context.Context, x, w []float32, batch, inDim, outDim int) ([]float32, CallStats, error) {
	out := make([]float32, batch*outDim)
	st, err := e.LinearIntoWithStats(ctx, out, x, w, batch, inDim, outDim)
	if err != nil {
		return nil, st, err
	}
	return out, st, nil
}

// LinearIntoWithStats computes x*w^T into dst and returns per-call execution stats.
func (e *DynamicExecutor) LinearIntoWithStats(ctx context.Context, dst, x, w []float32, batch, inDim, outDim int) (CallStats, error) {
	var st CallStats
	if e == nil {
		return st, fmt.Errorf("linear dynamic executor is nil")
	}
	if err := ctxErr(ctx); err != nil {
		return st, err
	}
	if batch <= 0 || inDim <= 0 || outDim <= 0 {
		return st, fmt.Errorf("invalid shape: batch=%d inDim=%d outDim=%d", batch, inDim, outDim)
	}
	if len(x) != batch*inDim {
		return st, fmt.Errorf("input length=%d want=%d", len(x), batch*inDim)
	}
	if len(w) != outDim*inDim {
		return st, fmt.Errorf("weight length=%d want=%d", len(w), outDim*inDim)
	}
	if len(dst) != batch*outDim {
		return st, fmt.Errorf("output length=%d want=%d", len(dst), batch*outDim)
	}

	return e.linearIOIntoWithStats(ctx, dst, x, w, batch, inDim, outDim, true)
}

// LinearIOIntoWithStats computes x*w where w is already laid out as [inDim,outDim].
//
// This is a fast path for callers that maintain the dynamic kernel's native
// weight layout and want to avoid transposing [outDim,inDim] on every call.
func (e *DynamicExecutor) LinearIOIntoWithStats(ctx context.Context, dst, x, wIO []float32, batch, inDim, outDim int) (CallStats, error) {
	if e == nil {
		return CallStats{}, fmt.Errorf("linear dynamic executor is nil")
	}
	if err := ctxErr(ctx); err != nil {
		return CallStats{}, err
	}
	if batch <= 0 || inDim <= 0 || outDim <= 0 {
		return CallStats{}, fmt.Errorf("invalid shape: batch=%d inDim=%d outDim=%d", batch, inDim, outDim)
	}
	if len(x) != batch*inDim {
		return CallStats{}, fmt.Errorf("input length=%d want=%d", len(x), batch*inDim)
	}
	if len(wIO) != inDim*outDim {
		return CallStats{}, fmt.Errorf("io weight length=%d want=%d", len(wIO), inDim*outDim)
	}
	if len(dst) != batch*outDim {
		return CallStats{}, fmt.Errorf("output length=%d want=%d", len(dst), batch*outDim)
	}
	return e.linearIOIntoWithStats(ctx, dst, x, wIO, batch, inDim, outDim, false)
}

// LinearOneHotIOIntoWithStats computes y = x*w for one-hot activations encoded
// by xs and a previously primed IO-layout weight matrix.
func (e *DynamicExecutor) LinearOneHotIOIntoWithStats(ctx context.Context, dst []float32, xs []int, batch, inDim, outDim int) (CallStats, error) {
	var st CallStats
	if e == nil {
		return st, fmt.Errorf("linear dynamic executor is nil")
	}
	if err := ctxErr(ctx); err != nil {
		return st, err
	}
	if batch <= 0 || inDim <= 0 || outDim <= 0 {
		return st, fmt.Errorf("invalid shape: batch=%d inDim=%d outDim=%d", batch, inDim, outDim)
	}
	if len(xs) > batch {
		return st, fmt.Errorf("one-hot batch len=%d want <= %d", len(xs), batch)
	}
	if len(dst) != batch*outDim && len(dst) != len(xs)*outDim {
		return st, fmt.Errorf("output length=%d want=%d or %d", len(dst), batch*outDim, len(xs)*outDim)
	}

	ck, compiled, err := e.kernelFor(batch, inDim, outDim)
	if err != nil {
		return st, err
	}
	st.Compiled = compiled
	if err := ctxErr(ctx); err != nil {
		return st, err
	}

	ck.mu.Lock()
	hwNS, err := ck.k.EvalOneHotIOIntoHW(dst, xs)
	ck.mu.Unlock()
	if err != nil {
		return st, fmt.Errorf("linear dynamic one-hot eval: %w", err)
	}
	st.HWExecutionNS = hwNS
	return st, nil
}

// PrimeWeightsIO stores an IO-layout weight matrix in the cached dynamic kernel.
func (e *DynamicExecutor) PrimeWeightsIO(batch, inDim, outDim int, wIO []float32) error {
	if e == nil {
		return fmt.Errorf("linear dynamic executor is nil")
	}
	if len(wIO) != inDim*outDim {
		return fmt.Errorf("io weight length=%d want=%d", len(wIO), inDim*outDim)
	}
	ck, _, err := e.kernelFor(batch, inDim, outDim)
	if err != nil {
		return err
	}
	ck.mu.Lock()
	err = ck.k.PrimeWeightsIO(wIO)
	ck.mu.Unlock()
	if err != nil {
		return fmt.Errorf("prime io weights: %w", err)
	}
	return nil
}

// UpdateWeightsIORows patches a subset of IO-layout weight rows in the cached kernel.
func (e *DynamicExecutor) UpdateWeightsIORows(batch, inDim, outDim int, wIO []float32, rows []int) error {
	if e == nil {
		return fmt.Errorf("linear dynamic executor is nil")
	}
	if len(wIO) != inDim*outDim {
		return fmt.Errorf("io weight length=%d want=%d", len(wIO), inDim*outDim)
	}
	ck, _, err := e.kernelFor(batch, inDim, outDim)
	if err != nil {
		return err
	}
	ck.mu.Lock()
	err = ck.k.UpdateWeightsIORows(wIO, rows)
	ck.mu.Unlock()
	if err != nil {
		return fmt.Errorf("update io weight rows: %w", err)
	}
	return nil
}

func (e *DynamicExecutor) linearIOIntoWithStats(ctx context.Context, dst, x, w []float32, batch, inDim, outDim int, transpose bool) (CallStats, error) {
	var st CallStats
	if e == nil {
		return st, fmt.Errorf("linear dynamic executor is nil")
	}
	if err := ctxErr(ctx); err != nil {
		return st, err
	}
	if batch <= 0 || inDim <= 0 || outDim <= 0 {
		return st, fmt.Errorf("invalid shape: batch=%d inDim=%d outDim=%d", batch, inDim, outDim)
	}
	if len(x) != batch*inDim {
		return st, fmt.Errorf("input length=%d want=%d", len(x), batch*inDim)
	}
	if len(dst) != batch*outDim {
		return st, fmt.Errorf("output length=%d want=%d", len(dst), batch*outDim)
	}

	ck, compiled, err := e.kernelFor(batch, inDim, outDim)
	if err != nil {
		return st, err
	}
	st.Compiled = compiled
	if err := ctxErr(ctx); err != nil {
		return st, err
	}

	ck.mu.Lock()
	weights := w
	if transpose {
		ck.ensureWeightScratch(inDim * outDim)
		transposeWeightsRowMajorOIToIO(ck.weightsIO, w, inDim, outDim)
		weights = ck.weightsIO
	}
	est, err := ck.k.EvalInto(dst, x, weights)
	ck.mu.Unlock()
	if err != nil {
		return st, fmt.Errorf("linear dynamic eval: %w", err)
	}
	st.HWExecutionNS = est.HWExecutionNS
	return st, nil
}

// Prepare compiles and caches a dynamic kernel for the provided shape.
func (e *DynamicExecutor) Prepare(batch, inDim, outDim int) error {
	if e == nil {
		return fmt.Errorf("linear dynamic executor is nil")
	}
	if batch <= 0 || inDim <= 0 || outDim <= 0 {
		return fmt.Errorf("invalid shape: batch=%d inDim=%d outDim=%d", batch, inDim, outDim)
	}
	_, _, err := e.kernelFor(batch, inDim, outDim)
	if err != nil {
		return fmt.Errorf("prepare dynamic kernel: %w", err)
	}
	return nil
}

func (e *DynamicExecutor) kernelFor(batch, inDim, outDim int) (*dynamicCompiledKernel, bool, error) {
	key := dynamicKernelKey(batch, inDim, outDim)

	e.mu.Lock()
	if ck := e.kernels[key]; ck != nil {
		e.stats.CacheHits++
		e.mu.Unlock()
		return ck, false, nil
	}
	e.mu.Unlock()

	k, err := dynamicmatmul.New(batch, inDim, outDim, dynamicmatmul.Options{QoS: e.qos})
	if err != nil {
		return nil, false, fmt.Errorf("compile dynamic kernel: %w", err)
	}
	candidate := &dynamicCompiledKernel{k: k}

	e.mu.Lock()
	if ck := e.kernels[key]; ck != nil {
		e.stats.CacheHits++
		e.mu.Unlock()
		k.Close()
		return ck, false, nil
	}
	e.kernels[key] = candidate
	e.stats.Compiles++
	e.stats.Kernels = len(e.kernels)
	e.mu.Unlock()
	return candidate, true, nil
}

func (ck *dynamicCompiledKernel) ensureWeightScratch(n int) {
	if cap(ck.weightsIO) < n {
		ck.weightsIO = make([]float32, n)
	} else {
		ck.weightsIO = ck.weightsIO[:n]
	}
}

func transposeWeightsRowMajorOIToIO(dst, src []float32, inDim, outDim int) {
	for out := 0; out < outDim; out++ {
		row := src[out*inDim : (out+1)*inDim]
		for in := 0; in < inDim; in++ {
			dst[in*outDim+out] = row[in]
		}
	}
}

func dynamicKernelKey(batch, inDim, outDim int) string {
	var b strings.Builder
	b.Grow(32)
	b.WriteString(strconv.Itoa(batch))
	b.WriteByte(':')
	b.WriteString(strconv.Itoa(inDim))
	b.WriteByte(':')
	b.WriteString(strconv.Itoa(outDim))
	return b.String()
}

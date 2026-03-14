package linear

import (
	"context"
	"fmt"
	"hash/fnv"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/tmc/apple/x/ane/mil"
	"github.com/tmc/apple/x/ane/model"
)

const defaultQoS = uint32(21)

// Options configures linear execution.
type Options struct {
	QoS uint32
}

// Stats reports cache and compile counters.
type Stats struct {
	Compiles  int
	CacheHits int
	Kernels   int
}

// CallStats reports per-call execution details.
type CallStats struct {
	Compiled      bool
	HWExecutionNS uint64
}

// Executor caches compiled kernels and runs linear forwards.
type Executor struct {
	qos uint32

	mu      sync.Mutex
	kernels map[string]*compiledKernel
	stats   Stats
}

type compiledKernel struct {
	mu sync.Mutex
	k  *model.Kernel

	inCF  []float32
	outCF []float32
}

// New creates a linear executor.
func New(opts Options) *Executor {
	qos := opts.QoS
	if qos == 0 {
		qos = defaultQoS
	}
	return &Executor{
		qos:     qos,
		kernels: make(map[string]*compiledKernel),
	}
}

// Stats returns a snapshot of executor counters.
func (e *Executor) Stats() Stats {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.stats
}

// Close releases all cached kernels.
func (e *Executor) Close() {
	if e == nil {
		return
	}
	e.mu.Lock()
	ks := make([]*compiledKernel, 0, len(e.kernels))
	for _, k := range e.kernels {
		ks = append(ks, k)
	}
	e.kernels = make(map[string]*compiledKernel)
	e.stats.Kernels = 0
	e.mu.Unlock()

	for _, ck := range ks {
		if ck != nil && ck.k != nil {
			ck.k.Close()
		}
	}
}

// Linear computes x*w^T where x is [batch,inDim] and w is [outDim,inDim].
//
// x and w use row-major layout.
func (e *Executor) Linear(ctx context.Context, x, w []float32, batch, inDim, outDim int) ([]float32, error) {
	out := make([]float32, batch*outDim)
	_, err := e.LinearIntoWithStats(ctx, out, x, w, batch, inDim, outDim)
	if err != nil {
		return nil, err
	}
	return out, err
}

// LinearWithStats computes x*w^T and returns per-call execution stats.
func (e *Executor) LinearWithStats(ctx context.Context, x, w []float32, batch, inDim, outDim int) ([]float32, CallStats, error) {
	out := make([]float32, batch*outDim)
	st, err := e.LinearIntoWithStats(ctx, out, x, w, batch, inDim, outDim)
	if err != nil {
		return nil, st, err
	}
	return out, st, nil
}

// LinearIntoWithStats computes x*w^T into dst and returns per-call execution stats.
//
// dst must have length batch*outDim and uses row-major layout.
func (e *Executor) LinearIntoWithStats(ctx context.Context, dst, x, w []float32, batch, inDim, outDim int) (CallStats, error) {
	var st CallStats
	if e == nil {
		return st, fmt.Errorf("linear executor is nil")
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

	ck, compiled, err := e.kernelFor(w, batch, inDim, outDim)
	if err != nil {
		return st, err
	}
	st.Compiled = compiled
	if err := ctxErr(ctx); err != nil {
		return st, err
	}

	ck.mu.Lock()
	est, err := ck.evalInto(dst, x, batch, inDim, outDim)
	ck.mu.Unlock()
	if err != nil {
		return st, fmt.Errorf("linear eval: %w", err)
	}
	st.HWExecutionNS = est.HWExecutionNS
	return st, nil
}

// Prepare compiles and caches a kernel for the provided shape and weights.
func (e *Executor) Prepare(w []float32, batch, inDim, outDim int) error {
	if e == nil {
		return fmt.Errorf("linear executor is nil")
	}
	if batch <= 0 || inDim <= 0 || outDim <= 0 {
		return fmt.Errorf("invalid shape: batch=%d inDim=%d outDim=%d", batch, inDim, outDim)
	}
	if len(w) != outDim*inDim {
		return fmt.Errorf("weight length=%d want=%d", len(w), outDim*inDim)
	}
	_, _, err := e.kernelFor(w, batch, inDim, outDim)
	if err != nil {
		return fmt.Errorf("prepare kernel: %w", err)
	}
	return nil
}

func (e *Executor) kernelFor(w []float32, batch, inDim, outDim int) (*compiledKernel, bool, error) {
	key := kernelKey(w, batch, inDim, outDim)

	e.mu.Lock()
	if ck := e.kernels[key]; ck != nil {
		e.stats.CacheHits++
		e.mu.Unlock()
		return ck, false, nil
	}
	e.mu.Unlock()

	k, err := compileKernel(w, batch, inDim, outDim, e.qos)
	if err != nil {
		return nil, false, err
	}
	candidate := &compiledKernel{k: k}

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

func compileKernel(w []float32, batch, inDim, outDim int, qos uint32) (*model.Kernel, error) {
	blob, err := mil.BuildWeightBlob(w, outDim, inDim)
	if err != nil {
		return nil, fmt.Errorf("build weight blob: %w", err)
	}
	k, err := model.Compile(model.CompileOptions{
		MILText:       mil.GenConv(inDim, outDim, batch),
		WeightBlob:    blob,
		QoS:           qos,
		PerfStatsMask: 1,
	})
	if err != nil {
		return nil, fmt.Errorf("compile kernel: %w", err)
	}
	return k, nil
}

func (ck *compiledKernel) evalInto(dst, x []float32, batch, inDim, outDim int) (model.EvalStats, error) {
	ck.ensureScratch(batch*inDim, batch*outDim)
	toChannelFirst(ck.inCF, x, batch, inDim)
	if err := ck.k.WriteInputF32(0, ck.inCF); err != nil {
		return model.EvalStats{}, fmt.Errorf("write input: %w", err)
	}
	st, err := ck.k.EvalWithStats()
	if err != nil {
		return model.EvalStats{}, fmt.Errorf("eval: %w", err)
	}
	if err := ck.k.ReadOutputF32(0, ck.outCF); err != nil {
		return model.EvalStats{}, fmt.Errorf("read output: %w", err)
	}
	fromChannelFirst(dst, ck.outCF, batch, outDim)
	return st, nil
}

func (ck *compiledKernel) ensureScratch(inElems, outElems int) {
	if cap(ck.inCF) < inElems {
		ck.inCF = make([]float32, inElems)
	} else {
		ck.inCF = ck.inCF[:inElems]
	}
	if cap(ck.outCF) < outElems {
		ck.outCF = make([]float32, outElems)
	} else {
		ck.outCF = ck.outCF[:outElems]
	}
}

func toChannelFirst(dst, src []float32, s, d int) {
	for t := 0; t < s; t++ {
		for i := 0; i < d; i++ {
			dst[i*s+t] = src[t*d+i]
		}
	}
}

func fromChannelFirst(dst, src []float32, s, d int) {
	for t := 0; t < s; t++ {
		for i := 0; i < d; i++ {
			dst[t*d+i] = src[i*s+t]
		}
	}
}

func kernelKey(w []float32, batch, inDim, outDim int) string {
	h := fnv.New64a()
	for _, f := range w {
		u := math.Float32bits(f)
		var b [4]byte
		b[0] = byte(u)
		b[1] = byte(u >> 8)
		b[2] = byte(u >> 16)
		b[3] = byte(u >> 24)
		_, _ = h.Write(b[:])
	}
	var b strings.Builder
	b.Grow(64)
	b.WriteString(strconv.Itoa(batch))
	b.WriteByte(':')
	b.WriteString(strconv.Itoa(inDim))
	b.WriteByte(':')
	b.WriteString(strconv.Itoa(outDim))
	b.WriteByte(':')
	b.WriteString(strconv.FormatUint(h.Sum64(), 16))
	return b.String()
}

func ctxErr(ctx context.Context) error {
	if ctx == nil {
		return nil
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

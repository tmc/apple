//go:build darwin

package linear

import (
	"fmt"
	"hash/fnv"
	"math"
	"sync"
	"unsafe"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
)

// Executor manages a cache of ANE kernels for linear operations.
type Executor struct {
	rt    *ane.Runtime
	mu    sync.Mutex
	cache map[cacheKey]*ane.Kernel
}

type cacheKey struct {
	batch, inDim, outDim int
	weightHash           uint64
}

// Stats reports cache statistics.
type Stats struct {
	CachedKernels int
}

// New creates a new linear executor.
func New(rt *ane.Runtime) *Executor {
	return &Executor{
		rt:    rt,
		cache: make(map[cacheKey]*ane.Kernel),
	}
}

// Linear computes y = x @ W^T using the ANE.
//
// x has shape [batch, inDim] (row-major), w has shape [outDim, inDim] (row-major),
// and the result has shape [batch, outDim].
func (e *Executor) Linear(x, w []float32, batch, inDim, outDim int) ([]float32, error) {
	if len(x) != batch*inDim {
		return nil, fmt.Errorf("linear: x length %d != batch*inDim (%d)", len(x), batch*inDim)
	}
	if len(w) != outDim*inDim {
		return nil, fmt.Errorf("linear: w length %d != outDim*inDim (%d)", len(w), outDim*inDim)
	}

	k, err := e.getOrCompile(w, batch, inDim, outDim)
	if err != nil {
		return nil, err
	}

	// Convert input to channel-first layout for the ANE: [1, inDim, batch, 1].
	chFirst := toChannelFirst(x, batch, inDim)
	if err := k.WriteInputF32(0, chFirst); err != nil {
		return nil, fmt.Errorf("linear: write input: %w", err)
	}

	if err := k.Eval(); err != nil {
		return nil, fmt.Errorf("linear: eval: %w", err)
	}

	// Read output in channel-first layout: [1, outDim, batch, 1].
	outBuf := make([]float32, batch*outDim)
	if err := k.ReadOutputF32(0, outBuf); err != nil {
		return nil, fmt.Errorf("linear: read output: %w", err)
	}

	// Convert back to row-major: [batch, outDim].
	return fromChannelFirst(outBuf, batch, outDim), nil
}

// Prepare pre-compiles a kernel for the given weight matrix and shape.
func (e *Executor) Prepare(w []float32, batch, inDim, outDim int) error {
	_, err := e.getOrCompile(w, batch, inDim, outDim)
	return err
}

// Stats returns cache statistics.
func (e *Executor) Stats() Stats {
	e.mu.Lock()
	defer e.mu.Unlock()
	return Stats{CachedKernels: len(e.cache)}
}

// Close releases all cached kernels.
func (e *Executor) Close() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, k := range e.cache {
		k.Close()
	}
	e.cache = nil
	return nil
}

func (e *Executor) getOrCompile(w []float32, batch, inDim, outDim int) (*ane.Kernel, error) {
	key := cacheKey{
		batch:      batch,
		inDim:      inDim,
		outDim:     outDim,
		weightHash: hashWeights(w),
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	if k, ok := e.cache[key]; ok {
		return k, nil
	}

	milText := mil.GenConv(inDim, outDim, batch)
	blob, err := mil.BuildWeightBlob(w, outDim, inDim)
	if err != nil {
		return nil, fmt.Errorf("linear: build weights: %w", err)
	}

	k, err := e.rt.Compile(ane.CompileOptions{
		ModelType:  ane.ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		return nil, fmt.Errorf("linear: compile: %w", err)
	}

	e.cache[key] = k
	return k, nil
}

func hashWeights(w []float32) uint64 {
	h := fnv.New64a()
	b := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(w))), len(w)*4)
	h.Write(b)
	return h.Sum64()
}

// toChannelFirst converts row-major [batch, dim] to channel-first [1, dim, batch, 1].
// out[i*batch+t] = in[t*dim+i]
func toChannelFirst(in []float32, batch, dim int) []float32 {
	out := make([]float32, len(in))
	for t := 0; t < batch; t++ {
		for i := 0; i < dim; i++ {
			out[i*batch+t] = in[t*dim+i]
		}
	}
	return out
}

// fromChannelFirst converts channel-first [1, dim, batch, 1] to row-major [batch, dim].
// out[t*dim+i] = in[i*batch+t]
func fromChannelFirst(in []float32, batch, dim int) []float32 {
	out := make([]float32, len(in))
	for t := 0; t < batch; t++ {
		for i := 0; i < dim; i++ {
			out[t*dim+i] = in[i*batch+t]
		}
	}
	return out
}

// linearCPU computes y = x @ W^T on the CPU for comparison.
func linearCPU(x, w []float32, batch, inDim, outDim int) []float32 {
	y := make([]float32, batch*outDim)
	for b := 0; b < batch; b++ {
		for o := 0; o < outDim; o++ {
			var sum float64
			for i := 0; i < inDim; i++ {
				sum += float64(x[b*inDim+i]) * float64(w[o*inDim+i])
			}
			y[b*outDim+o] = float32(sum)
		}
	}
	return y
}

// maxAbsDiff returns the maximum absolute difference between two slices.
func maxAbsDiff(a, b []float32) float32 {
	var max float32
	for i := range a {
		d := float32(math.Abs(float64(a[i] - b[i])))
		if d > max {
			max = d
		}
	}
	return max
}

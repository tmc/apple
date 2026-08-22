//go:build darwin

package engineattest

import (
	"errors"
	"testing"
	"unsafe"

	"github.com/tmc/apple/metal"
)

const testKernel = `
#include <metal_stdlib>
using namespace metal;
kernel void double_it(device float *v [[buffer(0)]],
                      uint i [[thread_position_in_grid]]) {
	v[i] = v[i] * 2.0f;
}
`

type gpuRig struct {
	device   metal.MTLDeviceObject
	queue    *Queue
	pipeline metal.MTLComputePipelineState
	buf      metal.MTLBuffer
	data     []float32
}

func newGPURig(t *testing.T) *gpuRig {
	t.Helper()
	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		t.Skip("no Metal device")
	}
	lib, err := device.NewLibraryWithSourceOptionsError(testKernel, nil)
	if err != nil {
		t.Fatalf("shader compilation: %v", err)
	}
	fn := lib.NewFunctionWithName("double_it")
	if fn.GetID() == 0 {
		t.Fatal("kernel double_it not found")
	}
	pipeline, err := device.NewComputePipelineStateWithFunctionError(fn)
	if err != nil {
		t.Fatalf("pipeline: %v", err)
	}
	const n = 1024
	buf := device.NewBufferWithLengthOptions(n*4, metal.MTLResourceStorageModeShared)
	if buf.GetID() == 0 {
		t.Fatal("buffer allocation failed")
	}
	data := unsafe.Slice((*float32)(buf.Contents()), n)
	for i := range data {
		data[i] = float32(i)
	}
	return &gpuRig{
		device:   device,
		queue:    NewQueue(device.NewCommandQueue()),
		pipeline: pipeline,
		buf:      buf,
		data:     data,
	}
}

// dispatch runs one real compute pass through the tracked queue.
func (r *gpuRig) dispatch() error {
	cb := r.queue.CommandBuffer()
	enc := cb.ComputeCommandEncoder()
	enc.SetComputePipelineState(r.pipeline)
	enc.SetBufferWithOffsetAtIndex(r.buf, 0, 0)
	enc.DispatchThreadsThreadsPerThreadgroup(
		metal.MTLSize{Width: uint(len(r.data)), Height: 1, Depth: 1},
		metal.MTLSize{Width: 64, Height: 1, Depth: 1})
	enc.EndEncoding()
	cb.Commit()
	cb.WaitUntilCompleted()
	return nil
}

func TestGPURealDispatchPasses(t *testing.T) {
	r := newGPURig(t)
	if err := r.queue.GPU(r.dispatch); err != nil {
		t.Errorf("real dispatch failed attestation: %v", err)
	}
	if got, want := r.data[3], float32(6); got != want {
		t.Errorf("data[3] = %v, want %v (kernel did not run?)", got, want)
	}
}

// TestGPUImpostorFails: a function that doubles the values on the CPU
// and claims the GPU ran must fail the assertion.
func TestGPUImpostorFails(t *testing.T) {
	r := newGPURig(t)
	err := r.queue.GPU(func() error {
		for i := range r.data {
			r.data[i] *= 2
		}
		return nil
	})
	if !errors.Is(err, ErrDidNotRun) {
		t.Errorf("impostor: got %v, want ErrDidNotRun", err)
	}
}

// TestGPUUncommittedFails: creating a command buffer without committing
// it — encoding work that never reaches the GPU — must fail.
func TestGPUUncommittedFails(t *testing.T) {
	r := newGPURig(t)
	err := r.queue.GPU(func() error {
		cb := r.queue.CommandBuffer()
		enc := cb.ComputeCommandEncoder()
		enc.SetComputePipelineState(r.pipeline)
		enc.SetBufferWithOffsetAtIndex(r.buf, 0, 0)
		enc.DispatchThreadsThreadsPerThreadgroup(
			metal.MTLSize{Width: uint(len(r.data)), Height: 1, Depth: 1},
			metal.MTLSize{Width: 64, Height: 1, Depth: 1})
		enc.EndEncoding()
		return nil // never committed
	})
	if !errors.Is(err, ErrDidNotRun) {
		t.Errorf("uncommitted: got %v, want ErrDidNotRun", err)
	}
}

func TestGPUErrorPropagates(t *testing.T) {
	r := newGPURig(t)
	want := errors.New("workload failed")
	err := r.queue.GPU(func() error { return want })
	if !errors.Is(err, want) {
		t.Errorf("got %v, want the workload's own error", err)
	}
}

// TestGPUWorkOutsideRegionFails: a dispatch before the attested region
// must not count as evidence for it.
func TestGPUWorkOutsideRegionFails(t *testing.T) {
	r := newGPURig(t)
	if err := r.dispatch(); err != nil {
		t.Fatal(err)
	}
	err := r.queue.GPU(func() error { return nil })
	if !errors.Is(err, ErrDidNotRun) {
		t.Errorf("outside-region work counted: got %v, want ErrDidNotRun", err)
	}
}

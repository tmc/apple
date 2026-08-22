// Command matrixmultiply multiplies two matrices with MPSMatrixMultiplication.
//
// MPS exposes BLAS-style GEMM as an MPSKernel: the operands are MPSMatrix
// views onto MTLBuffers, described by an MPSMatrixDescriptor that carries the
// row stride, and the kernel is encoded into a command buffer like any other
// GPU work.
//
// The kernel computes
//
//	C = alpha * A * B + beta * C
//
// which this example runs with alpha=1, beta=0 so C is just A*B. The result is
// compared against a plain Go reference implementation and the largest
// absolute difference is printed.
//
// Usage:
//
//	matrixmultiply
package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/metal"
	mps "github.com/tmc/apple/metalperformanceshaders"
)

func init() {
	runtime.LockOSThread()
}

// The shapes to multiply: A is m x k, B is k x n, C is m x n. The values are
// deliberately not square so a row-stride mistake cannot pass unnoticed.
const (
	m = 64
	k = 48
	n = 32
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "matrixmultiply: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		return fmt.Errorf("no Metal device available; this example needs a Metal-capable GPU")
	}
	queue := device.NewCommandQueue()
	if queue.GetID() == 0 {
		return fmt.Errorf("could not create a Metal command queue")
	}
	fmt.Printf("device: %s\n", device.Name())

	a := fill(m, k, 1)
	b := fill(k, n, 2)

	left, err := newMatrix(device, a, m, k)
	if err != nil {
		return fmt.Errorf("left matrix: %w", err)
	}
	right, err := newMatrix(device, b, k, n)
	if err != nil {
		return fmt.Errorf("right matrix: %w", err)
	}
	result, err := newMatrix(device, make([]float32, m*n), m, n)
	if err != nil {
		return fmt.Errorf("result matrix: %w", err)
	}

	// alpha=1, beta=0 makes this a plain C = A*B.
	gemm := mps.NewMatrixMultiplicationWithDeviceTransposeLeftTransposeRightResultRowsResultColumnsInteriorColumnsAlphaBeta(
		device, false, false, m, n, k, 1, 0)
	if gemm.GetID() == 0 {
		return fmt.Errorf("could not create MPSMatrixMultiplication")
	}

	buf := queue.CommandBuffer()
	if buf.GetID() == 0 {
		return fmt.Errorf("could not create a command buffer")
	}
	gemm.EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix(buf, left.matrix, right.matrix, result.matrix)
	buf.Commit()
	buf.WaitUntilCompleted()

	got := readFloats(result.buffer, m*n)
	want := reference(a, b)

	maxDiff := 0.0
	for i := range want {
		if d := math.Abs(float64(got[i] - want[i])); d > maxDiff {
			maxDiff = d
		}
	}
	fmt.Printf("C = A*B with A %dx%d, B %dx%d\n", m, k, k, n)
	fmt.Printf("  C[%d][%d] = %g (cpu %g)\n", 0, 0, got[0], want[0])
	fmt.Printf("  C[%d][%d] = %g (cpu %g)\n", m-1, n-1, got[m*n-1], want[m*n-1])
	fmt.Printf("max |gpu-cpu| = %g\n", maxDiff)

	// float32 accumulation over k=48 terms of magnitude ~1 leaves a few ulps
	// of slack; anything larger is a real disagreement.
	const tol = 1e-4
	if maxDiff > tol {
		return fmt.Errorf("GPU result disagrees with the CPU reference by %g (tolerance %g)", maxDiff, tol)
	}
	fmt.Println("PASS")
	return nil
}

// matrix pairs an MPSMatrix with the buffer that backs it, so callers can read
// the contents back after the GPU has written them.
type matrix struct {
	matrix mps.MPSMatrix
	buffer metal.MTLBuffer
}

// newMatrix returns a rows x columns MPSMatrix over a shared-storage buffer
// holding vals in row-major order.
func newMatrix(device metal.MTLDeviceObject, vals []float32, rows, columns int) (matrix, error) {
	rowBytes := uint(columns) * 4
	buf := device.NewBufferWithBytesLengthOptions(unsafe.Pointer(&vals[0]), uint(rows)*rowBytes, metal.MTLResourceStorageModeShared)
	runtime.KeepAlive(vals)
	if buf.GetID() == 0 {
		return matrix{}, fmt.Errorf("could not allocate a %dx%d buffer", rows, columns)
	}
	desc := mps.NewMatrixDescriptorWithRowsColumnsRowBytesDataType(uint(rows), uint(columns), rowBytes, mps.MPSDataTypeFloat32)
	if desc.GetID() == 0 {
		return matrix{}, fmt.Errorf("could not create a %dx%d matrix descriptor", rows, columns)
	}
	mat := mps.NewMatrixWithBufferDescriptor(buf, desc)
	if mat.GetID() == 0 {
		return matrix{}, fmt.Errorf("could not create a %dx%d MPSMatrix", rows, columns)
	}
	return matrix{matrix: mat, buffer: buf}, nil
}

// fill returns a rows x columns matrix with values that vary along both axes,
// so a transposed or mis-strided read produces visibly wrong numbers.
func fill(rows, columns, seed int) []float32 {
	out := make([]float32, rows*columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			out[i*columns+j] = float32(math.Sin(float64(seed*97 + i*13 + j*7)))
		}
	}
	return out
}

// reference computes A*B in Go, as a check on the GPU result.
func reference(a, b []float32) []float32 {
	out := make([]float32, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var sum float32
			for p := 0; p < k; p++ {
				sum += a[i*k+p] * b[p*n+j]
			}
			out[i*n+j] = sum
		}
	}
	return out
}

// readFloats copies count float32 values out of a shared-storage buffer.
func readFloats(buf metal.MTLBuffer, count int) []float32 {
	out := make([]float32, count)
	copy(out, unsafe.Slice((*float32)(buf.Contents()), count))
	return out
}

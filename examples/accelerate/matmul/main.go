// Command matmul benchmarks matrix multiplication using Apple Accelerate BLAS
// (cblas_sgemm) versus a naive pure-Go implementation.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/tmc/apple/accelerate"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	sizes := []int{64, 256, 512, 1024}

	fmt.Println("Matrix Multiplication Benchmark: Accelerate BLAS vs Naive Go")
	fmt.Println()
	fmt.Printf("%-6s  %12s  %12s  %8s\n", "Size", "BLAS (GFLOPS)", "Go (GFLOPS)", "Speedup")
	fmt.Println("------  ------------  ------------  --------")

	for _, n := range sizes {
		a := randMatrix(n * n)
		b := randMatrix(n * n)

		blasGFLOPS := benchBLAS(a, b, n)
		goGFLOPS := benchNaive(a, b, n)

		speedup := blasGFLOPS / goGFLOPS
		fmt.Printf("%-6d  %12.2f  %12.2f  %7.1fx\n", n, blasGFLOPS, goGFLOPS, speedup)
	}
}

func benchBLAS(a, b []float32, n int) float64 {
	c := make([]float32, n*n)
	flops := 2.0 * float64(n) * float64(n) * float64(n)

	// Warm up.
	accelerate.Cblas_sgemm(
		accelerate.CblasRowMajor,
		accelerate.CblasNoTrans,
		accelerate.CblasNoTrans,
		n, n, n,
		1.0, a, n,
		b, n,
		0.0, c, n,
	)

	iters := calibrateIters(func() {
		accelerate.Cblas_sgemm(
			accelerate.CblasRowMajor,
			accelerate.CblasNoTrans,
			accelerate.CblasNoTrans,
			n, n, n,
			1.0, a, n,
			b, n,
			0.0, c, n,
		)
	})

	start := time.Now()
	for range iters {
		accelerate.Cblas_sgemm(
			accelerate.CblasRowMajor,
			accelerate.CblasNoTrans,
			accelerate.CblasNoTrans,
			n, n, n,
			1.0, a, n,
			b, n,
			0.0, c, n,
		)
	}
	elapsed := time.Since(start).Seconds()
	return (flops * float64(iters)) / elapsed / 1e9
}

func benchNaive(a, b []float32, n int) float64 {
	c := make([]float32, n*n)
	flops := 2.0 * float64(n) * float64(n) * float64(n)

	// Warm up.
	naiveMatmul(a, b, c, n)

	iters := calibrateIters(func() {
		naiveMatmul(a, b, c, n)
	})

	start := time.Now()
	for range iters {
		naiveMatmul(a, b, c, n)
	}
	elapsed := time.Since(start).Seconds()
	return (flops * float64(iters)) / elapsed / 1e9
}

func naiveMatmul(a, b, c []float32, n int) {
	for i := range n {
		for j := range n {
			var sum float32
			for k := range n {
				sum += a[i*n+k] * b[k*n+j]
			}
			c[i*n+j] = sum
		}
	}
}

// calibrateIters finds an iteration count that runs for at least 200ms.
func calibrateIters(fn func()) int {
	iters := 1
	for {
		start := time.Now()
		for range iters {
			fn()
		}
		if time.Since(start) >= 200*time.Millisecond {
			return iters
		}
		iters *= 2
	}
}

func randMatrix(size int) []float32 {
	m := make([]float32, size)
	for i := range m {
		m[i] = rand.Float32()*2 - 1
	}
	return m
}

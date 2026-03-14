// Package linear provides reusable ANE-backed linear forward execution.
//
// It compiles and caches in-memory kernels keyed by shape and weights, then
// executes row-major x*w^T using ANE runtime plumbing from this repository.
//
// DynamicExecutor provides a compile-once variant for workloads whose weights
// change across evaluations.
package linear

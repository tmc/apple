// Package dynamicmatmul provides compile-once ANE matmul kernels with
// runtime-provided weights.
//
// It stages activations and weights into a single ANE input surface, then
// evaluates y = x*w without recompiling when w changes.
package dynamicmatmul

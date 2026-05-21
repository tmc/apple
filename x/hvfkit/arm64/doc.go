//go:build darwin && arm64

// Package arm64 describes ARM64 registers and trap syndromes used with
// Hypervisor.framework.
//
// The tables in this package name common vCPU, system, timer, SIMD/FP, and GIC
// registers. The values are the generated symbols from
// github.com/tmc/apple/hypervisor.
package arm64

//go:build darwin && arm64

// Package hvfkit provides small helpers around Apple's Hypervisor.framework.
//
// The package keeps the generated hypervisor package as the source of truth for
// constants, register types, and raw calls. It adds error-returning wrappers for
// common VM, vCPU, memory, and GIC operations.
package hvfkit

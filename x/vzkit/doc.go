// Package vzkit provides shared infrastructure for building macOS and Linux
// virtual machines using Apple's Virtualization framework via purego.
//
// It extracts common VM plumbing — dispatch queues, run loops, completion
// handlers, vsock, VirtioFS, networking, and VM lifecycle management — so
// that higher-level tools (vz-macos, vz-container, etc.) can share one
// tested implementation instead of duplicating it.
//
// All Objective-C interop uses purego (cgo-free). The generated bindings
// from github.com/tmc/apple provide the type-safe wrappers.
package vzkit

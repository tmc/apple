// Package vzkit provides shared infrastructure for building macOS and Linux
// virtual machines using Apple's Virtualization framework via purego.
//
// It extracts common VM plumbing — dispatch queues, run loops, completion
// handlers, vsock, VirtioFS, networking, and VM lifecycle management — so
// that higher-level tools (vz-macos, vz-container, etc.) can share one
// tested implementation instead of duplicating it.
//
// The root package is the high-level path for basic virtual machine setup and
// lifecycle management. More focused domains live in subpackages such as
// audio, display, network, virtiofs, storage, restore, snapshot, vsock,
// balloon, vm, and privatevm. Lower-level private-API wrappers such as
// clipboard, capture, vnc, framebuffer, vminput, storagehotplug,
// usbpassthrough, debugstub, configcodec, and spice build on the generated
// bindings from github.com/tmc/apple. Experimental private-API wrappers live
// under the exp tree.
//
// All Objective-C interop uses purego (cgo-free). The generated bindings
// from github.com/tmc/apple provide the type-safe wrappers.
package vzkit

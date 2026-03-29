// Package disk provides helpers for managing macOS disk images via hdiutil.
//
// It can query whether a disk image is currently attached, detach it safely
// with escalating retry strategies, and poll until a disk becomes available
// after a VM stops.
//
// All operations shell out to hdiutil and diskutil; no cgo or framework
// bindings are required.
package disk

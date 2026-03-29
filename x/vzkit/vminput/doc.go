// Package vminput provides direct virtual-machine input injection helpers.
//
// It is intended to wrap the private keyboard, pointer, scroll, and multitouch
// event sender APIs exposed by VZVirtualMachine so callers can inject input
// without depending on AppKit focus or CGEvent delivery through the host window
// server.
//
// The package should expose a small set of helpers for:
//
//	Keyboard events and text input
//	Pointer movement, button events, and scrolling
//	Multitouch and gesture delivery
//
// All VM operations should run on the caller-supplied [vm.Queue] so the API
// remains safe to use from arbitrary goroutines.
package vminput

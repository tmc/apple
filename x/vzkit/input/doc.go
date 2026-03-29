// Package input provides CGEvent-based keyboard and mouse input helpers.
//
// It wraps CoreGraphics CGEvent functions via purego (cgo-free) for creating
// and posting keyboard and mouse events. Events can be posted to a specific
// process (e.g., a VM window) or to the system HID event tap.
//
// Basic usage:
//
//	ev, err := input.CreateKeyboardEvent(0, input.KeyReturn, true)
//	if err != nil {
//		log.Fatal(err)
//	}
//	input.PostToSelf(ev)
package input

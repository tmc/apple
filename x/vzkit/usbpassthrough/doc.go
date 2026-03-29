// Package usbpassthrough provides runtime USB passthrough helpers.
//
// It is intended to wrap the private USB controller, XHCI controller, and host
// passthrough device configuration APIs so callers can capture host USB devices
// and attach or detach them from a running virtual machine.
//
// The package should cover:
//
//	Host device discovery and lookup
//	Passthrough device capture and release
//	Live attach and detach through a controller
//
// This package handles runtime USB passthrough. Static USB-backed storage
// configuration remains separate.
package usbpassthrough

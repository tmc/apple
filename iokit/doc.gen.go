// Code generated from Apple documentation for iokit. DO NOT EDIT.

// Package iokit provides Go bindings for the iokit framework.
//
// Access hardware devices and drivers from your apps and services.
//
// The IOKit framework implements nonkernel access to IOKit objects such drivers and nubs through the device-interface mechanism.
//
// # Serial Ports
//
//   - Communicating with a Modem on a Serial Port: Find and connect to a modem attached to a serial port using IOKit.
//
// [iokit Documentation]: https://developer.apple.com/documentation/iokit
package iokit

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the iokit library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/iokit.framework/iokit",
	"/usr/lib/libiokit.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: iokit: failed to load framework from any known path\n")
}

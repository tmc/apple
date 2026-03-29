
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

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/iokit.framework/iokit"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: iokit: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}


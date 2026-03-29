
// Code generated from Apple documentation for coreservices. DO NOT EDIT.

// Package coreservices provides Go bindings for the coreservices framework.
//
// Access and manage key operating system services, such as launch and identity services.
//
// This collection of documents provides the API reference for the Core Services framework, which encompasses many fundamental operating system services used by Carbon applications.
//
// [coreservices Documentation]: https://developer.apple.com/documentation/coreservices
package coreservices

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/coreservices.framework/coreservices"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: coreservices: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}


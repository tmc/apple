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

// frameworkPaths lists paths to try when loading the coreservices library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/coreservices.framework/coreservices",
	"/usr/lib/libcoreservices.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: coreservices: failed to load framework from any known path\n")
}

// Code generated from Apple documentation for Network. DO NOT EDIT.

// Package network provides Go bindings for the Network framework.
//
// # Key Types
//
//   - [NWParameters]
//   - [NWPath]
//   - [NWInterface]
//   - [NWConnection]
//   - [NWEndpoint]
//   - [NWBrowseDescriptor]
//   - [NWAdvertiseDescriptor]
//   - [NWBrowser]
package network

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Network library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/PrivateFrameworks/Network.framework/Network"}

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
	fmt.Fprintf(os.Stderr, "warning: Network: failed to load framework from any known path\n")
}

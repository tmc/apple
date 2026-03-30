// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

// Package remotecoreml provides Go bindings for the remotecoreml framework.
//
// # Key Types
//
//   - [MLNetworkHeaderEncoding]
//   - [MLNetworking]
//   - [MLRemoteConnection]
//   - [MLServer]
//   - [MLNetworkOptions]
//   - [MLNetworkPacket]
//   - [CoreMLModelSecurityServiceToClient]
//   - [MLNetworkUtilities]
//   - [MLLog]
//   - [CoreMLVersion]
package remotecoreml

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the remotecoreml library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/PrivateFrameworks/RemoteCoreML.framework/RemoteCoreML"}

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
	fmt.Fprintf(os.Stderr, "warning: remotecoreml: failed to load framework from any known path\n")
}

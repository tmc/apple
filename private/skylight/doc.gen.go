// Code generated from Apple documentation for skylight. DO NOT EDIT.

// Package skylight provides Go bindings for the skylight framework.
//
// # Key Types
//
//   - [SLSDisplayController]
//   - [SLSSpaceWindowManager]
//   - [SLDisplayPresetDevice]
//   - [SLSXPCService]
//   - [SLContentStream]
//   - [SLSBrightnessControlClient]
//   - [SLContentFilter]
//   - [SLSEventAuthenticationMessage]
//   - [SLSSkyLightGestureEventAuthenticationMessage]
//   - [SLSSkyLightKeyEventAuthenticationMessage]
package skylight

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the skylight library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/PrivateFrameworks/SkyLight.framework/SkyLight"}

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
	fmt.Fprintf(os.Stderr, "warning: skylight: failed to load framework from any known path\n")
}

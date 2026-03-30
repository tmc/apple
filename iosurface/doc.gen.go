// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

// Package iosurface provides Go bindings for the IOSurface framework.
//
// Share hardware-accelerated buffer data (framebuffers and textures) across multiple processes. Manage image memory more efficiently.
//
// The IOSurface framework provides a framebuffer object suitable for sharing across process boundaries. It is commonly used to allow applications to move complex image decompression and draw logic into a separate process to enhance security.
//
// # Classes
//
//   - IOSurface: Data type representing an IOSurface opaque object.
//   - IOSurfaceRef: Data type representing an IOSurface opaque object.
//
// # Variables
//
//   - kIOSurfaceContentHeadroom
//   - kIOSurfaceCopybackCache
//   - kIOSurfaceCopybackInnerCache
//   - kIOSurfaceDefaultCache
//   - kIOSurfaceInhibitCache
//   - kIOSurfaceMapCacheShift
//   - kIOSurfaceMapCopybackCache
//   - kIOSurfaceMapCopybackInnerCache
//   - kIOSurfaceMapDefaultCache
//   - kIOSurfaceMapInhibitCache
//   - kIOSurfaceMapWriteCombineCache
//   - kIOSurfaceMapWriteThruCache
//   - kIOSurfaceWriteCombineCache
//   - kIOSurfaceWriteThruCache
//
// # Key Types
//
//   - [IOSurface] - Data type representing an IOSurface opaque object.
//
// [IOSurface Documentation]: https://developer.apple.com/documentation/IOSurface
package iosurface

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the IOSurface library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/IOSurface.framework/IOSurface",
	"/usr/lib/libIOSurface.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: IOSurface: failed to load framework from any known path\n")
}

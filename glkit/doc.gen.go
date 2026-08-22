// Code generated from Apple documentation for GLKit. DO NOT EDIT.

// Package glkit provides Go bindings for the GLKit framework.
//
// Speed up OpenGL ES or OpenGL app development. Use math libraries,
// background texture loading, pre-created shader effects, and a standard view
// and view controller to implement your rendering loop.
//
// The GLKit framework provides functions and classes that reduce the effort
// required to create new shader-based apps or to port existing apps that rely
// on fixed-function vertex or fragment processing provided by earlier
// versions of OpenGL ES or OpenGL.
//
// # Shader-Based Rendering Effects
//
//   - [GLKNamedEffect]: A standard interface for objects that provide shader-based OpenGL rendering effects.
//
// # Rendering Effect Parameters
//
//   - [GLKit Effects Constants] ([GLKVertexAttrib])
//
// # Math Utilties
//
//   - [GLKMatrixStackRef]: An opaque type that represents a stack of 4 x 4 matrices, providing support for hierarchical transform modeling and similar tasks.
//   - [GLKMatrix3] ([GLKMatrix2], [GLKMatrix3])
//   - [GLKMatrix4] ([GLKMatrix4])
//   - [GLKVector2] ([GLKVector2])
//   - [GLKVector3] ([GLKVector3])
//   - [GLKVector4] ([GLKVector4])
//   - [GLKQuaternion] ([GLKQuaternion])
//   - [GLKit Math Utilities]//
//
// [GLKit Effects Constants]: https://developer.apple.com/documentation/glkit/glkit-effects-constants
// [GLKit Math Utilities]: https://developer.apple.com/documentation/glkit/glkit-math-utilities
package glkit

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the GLKit library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/GLKit.framework/GLKit",
	"/usr/lib/libGLKit.dylib",
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
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: GLKit: failed to load framework from any known path\n")
	}
}

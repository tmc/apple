// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

// Package metalkit provides Go bindings for the MetalKit framework.
//
// Build Metal apps quicker and easier using a common set of utility classes.
//
// # View Management
//
//   - MTKView: A specialized view that creates, configures, and displays Metal objects.
//   - MTKViewDelegate: Methods for responding to a MetalKit view’s drawing and resizing events.
//
// # Texture Loading
//
//   - MTKTextureLoader: An object that creates textures from existing data in common image formats.
//
// # Model Handling
//
//   - MTKMesh: A container for the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//   - MTKMeshBuffer: A buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//   - MTKMeshBufferAllocator: An interface for allocating a MetalKit buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//   - MTKSubmesh: A container for the index data of a Model I/O submesh, suitable for use in a Metal app.
//   - Conversion Functions: Convert between Metal and Model I/O vertex representations.
//   - Model Errors: Learn about errors thrown by model handling methods. ([MTKModelError])
//
// # Key Types
//
//   - [MTKView] - A specialized view that creates, configures, and displays Metal objects.
//   - [MTKTextureLoader] - An object that creates textures from existing data in common image formats.
//   - [MTKMesh] - A container for the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//   - [MTKMeshBuffer] - A buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//   - [MTKSubmesh] - A container for the index data of a Model I/O submesh, suitable for use in a Metal app.
//   - [MTKMeshBufferAllocator] - An interface for allocating a MetalKit buffer that backs the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//
// [MetalKit Documentation]: https://developer.apple.com/documentation/MetalKit
package metalkit

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the MetalKit library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/MetalKit.framework/MetalKit",
	"/usr/lib/libMetalKit.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: MetalKit: failed to load framework from any known path\n")
}

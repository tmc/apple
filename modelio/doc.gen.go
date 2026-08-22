// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

// Package modelio provides Go bindings for the ModelIO framework.
//
// Import, export, and manipulate 3D models using a common infrastructure that
// integrates MetalKit, GLKit, and SceneKit.
//
// The Model I/O framework provides a system-level understanding of 3D model
// assets and related resources. You can use this framework to import and
// export assets from and to a variety of industry standard file formats
// supported by popular authoring tools and game engines. You can also use
// Model I/O to generate or process model and texture data—for example, to
// create subdivision surfaces, to bake ambient occlusion textures, or to
// generate light probes. Model I/O can share data buffers with the MetalKit,
// GLKit, and SceneKit frameworks to help you load, process, and render 3D
// assets efficiently.
//
// # 3D Asset Basics
//
//   - MDLAsset: An indexed container for 3D objects and associated information, such as transform hierarchies, meshes, cameras, and lights. ([MDLProbePlacement])
//   - MDLObject: The base class for objects that are part of a 3D asset, including meshes, cameras, and lights. ([MDLAxisAlignedBoundingBox])
//   - MDLSubmesh: A container for index buffer data and material information to be used in rendering all or part of a 3D object. ([MDLIndexBitDepth], [MDLGeometryType])
//   - [MDLNamed]: The common interface for Model I/O objects that expose a human-readable name.
//
// # Managing Mesh Data
//
//   - [MDLMeshBuffer]: The general interface for managing storage of vertex and index data used in loading, processing, and rendering meshes. ([MDLMeshBufferType])
//   - [MDLMeshBufferAllocator]: The general interface for managing allocation of data buffers to be used in loading, processing, and rendering meshes.
//   - [MDLMeshBufferZone]: The general interface for logical pools of memory used in allocation of related mesh data buffers.
//   - MDLVertexAttribute: A description of the format of per-vertex data for a single vertex attribute in a mesh object. ([MDLVertexFormat])
//
// # Materials
//
//   - MDLMaterialProperty: A definition for one specific aspect of the rendering parameters for a material. ([MDLMaterialSemantic], [MDLMaterialPropertyType])
//
// # Textures
//
//   - MDLTexture: A source of texel data to be used in rendering material surface appearances. ([MDLTextureChannelEncoding])
//   - MDLTextureFilter: A description of filtering modes for a renderer to use when sampling from a texture. ([MDLMaterialTextureWrapMode], [MDLMaterialTextureFilterMode], [MDLMaterialMipMapFilterMode])
//
// # Lights
//
//   - MDLLight: The abstract superclass for objects that describe light sources in a scene. ([MDLLightType])
//   - [MDLLightProbeIrradianceDataSource]: Adopt this protocol to provide information for use in automatic placement of light probes around a scene.
//
// # Cameras
//
//   - MDLCamera: A point of view for rendering a 3D scene, along with a set of parameters describing an intended appearance for rendering. ([MDLCameraProjection])
//
// # Extensible Asset Format Support
//
//   - [MDLComponent]: The base protocol for extensible file format support in Model I/O.
//   - [MDLObjectContainerComponent]: The general interface for classes that can act as containers in an object hierarchy.
//   - [MDLTransformComponent]: The general interface for classes that manage local coordinate space transforms for 3D objects
//
// # Volumetric Representations
//
//   - MDLVoxelArray: A model of a 3D object’s solid volume as a collection of , or cubic units. ([MDLVoxelIndex], [MDLVoxelIndexExtent])
//
// # Protocols
//
//   - [MDLAssetResolver]
//   - [MDLJointAnimation]
//   - [MDLTransformOp]
package modelio

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ModelIO library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/ModelIO.framework/ModelIO",
	"/usr/lib/libModelIO.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: ModelIO: failed to load framework from any known path\n")
	}
}

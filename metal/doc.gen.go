// Code generated from Apple documentation for Metal. DO NOT EDIT.

// Package metal provides Go bindings for the Metal framework.
//
// Render advanced 3D graphics and compute data in parallel with graphics processors.
//
// The Metal framework gives your app direct access to a device’s graphics processing unit (GPU). With Metal, apps can leverage a GPU to quickly render complex scenes and run computational tasks in parallel. For example, apps in these categories use Metal to maximize their performance:
//
// # Essentials
//
//   - Understanding the Metal 4 core API: Discover the features and functionality in the Metal 4 foundational APIs.
//   - Drawing a triangle with Metal 4: Render a colorful, rotating 2D triangle by running draw commands with a render pipeline on a GPU.
//   - Performing calculations on a GPU: Use Metal to find GPUs and perform calculations on them.
//   - Using Metal to draw a view’s contents: Create a MetalKit view and a render pass to draw the view’s contents.
//
// # Samples
//
//   - Metal sample code library: Explore the complete set of Metal samples.
//
// # GPU devices
//
//   - GPU devices and work submission: Find any available GPU, submit work to it with command buffers, suspend work, and coordinate between multiple GPUs. ([MTLDevice], [MTL4CommandQueue], [MTL4CommandQueueDescriptor], [MTL4CommandQueueError], [MTL4CommandBuffer])
//
// # Command encoders
//
//   - Render passes: Encode a render pass to draw graphics into an image. ([MTL4RenderCommandEncoder], [MTLRenderCommandEncoder], [MTL4RenderEncoderOptions], [MTLTriangleFillMode], [MTLWinding])
//   - Compute passes: Encode a compute pass that runs computations in parallel on a thread grid, processing and manipulating Metal resource data on multiple cores of a GPU. ([MTL4ComputeCommandEncoder], [MTLComputeCommandEncoder], [MTL4ComputePipelineDescriptor], [MTLComputePipelineDescriptor], [MTLComputePipelineState])
//   - Machine learning passes: Add machine learning model inference to your Metal app’s GPU workflow. ([MTL4MachineLearningCommandEncoder], [MTL4MachineLearningPipelineState], [MTL4MachineLearningPipelineDescriptor], [MTL4MachineLearningPipelineReflection])
//   - Blit passes: Encode a block information transfer pass to adjust and copy data to and from GPU resources, such as buffers and textures. ([MTLBlitCommandEncoder], [MTLBlitOption], [MTLBlitPassDescriptor], [MTLBlitPassSampleBufferAttachmentDescriptor], [MTLBlitPassSampleBufferAttachmentDescriptorArray])
//   - Indirect command encoding: Store draw commands in Metal buffers and run them at a later time on the GPU, either once or repeatedly. ([MTLIndirectCommandBuffer], [MTLIndirectCommandBufferDescriptor], [MTLIndirectCommandType], [MTLIndirectCommandBufferExecutionRange], [MTLIndirectComputeCommand])
//   - Ray tracing with acceleration structures: Build a representation of your scene’s geometry using triangles and bounding volumes to quickly trace rays through the scene. ([MTLAccelerationStructure], [MTL4AccelerationStructureDescriptor], [MTLAccelerationStructureDescriptor], [MTL4PrimitiveAccelerationStructureDescriptor], [MTLPrimitiveAccelerationStructureDescriptor])
//
// # Resources
//
//   - Resource fundamentals: Control the common attributes of all Metal memory resources, including buffers and textures, and how to configure their underlying memory. ([MTLResidencySet], [MTLResidencySetDescriptor], [MTLResourceViewPool], [MTLResourceViewPoolDescriptor], [MTLTextureViewPool])
//   - Buffers: Create and manage untyped data your app uses to exchange information with its shader functions. ([MTLBuffer], [MTLArgumentDescriptor], [MTLArgumentEncoder])
//   - Textures: Create and manage typed data your app uses to exchange information with its shader functions. ([MTLTexture], [MTLTextureCompressionType], [MTLTextureDescriptor], [MTLSharedTextureHandle], [MTLPixelFormat])
//   - Memory heaps: Take control of your app’s GPU memory management by creating a large memory allocation for various buffers, textures, and other resources. ([MTLHeap], [MTLHeapDescriptor], [MTLHeapType], [MTLSizeAndAlign])
//   - Resource loading: Load assets in your games and apps quickly by running a dedicated input/output queue alongside your GPU tasks. ([MTLIOCommandQueue], [MTLIOCommandQueueDescriptor], [MTLIOPriority], [MTLIOCommandQueueType], [MTLIOScratchBufferAllocator])
//   - Resource synchronization: Prevent multiple commands that can access the same resources simultaneously by coordinating those reads and writes with barriers, fences, or events. ([MTLStages], [MTLFence], [MTLRenderStages], [MTLBarrierScope], [MTL4VisibilityOptions])
//
// # Shader compilation and libraries
//
//   - Using the Metal 4 compilation API: Control when and how you compile an app’s shaders.
//   - Shader libraries: Manage and load your app’s Metal shaders. ([MTL4Compiler], [MTL4CompilerDescriptor], [MTL4CompilerTaskOptions], [MTL4CompilerTaskStatus], [MTL4Archive])
//   - Using function specialization to build pipeline variants: Create pipelines for different levels of detail from a common shader source.
//
// # Presentation
//
//   - Managing your game window for Metal in macOS: Set up a window and view for optimally displaying your Metal content.
//   - Managing your Metal app window in iPadOS: Set up a window that handles dynamically resizing your Metal content.
//   - Adapting your game interface for smaller screens: Make text legible on all devices the player chooses to run your game on.
//   - Onscreen presentation: Show the output from a GPU’s rendering pass to the user in your app.
//   - HDR content: Take advantage of high dynamic range to present more vibrant colors in your apps and games.
//
// # Developer tools
//
//   - Supporting Simulator in a Metal app: Configure alternative render paths in your Metal app to enable running your app in Simulator.
//   - Capturing Metal commands programmatically: Invoke a Metal frame capture from your app, then save the resulting GPU trace to a file or view it in Xcode.
//   - Logging shader debug messages: Print debugging messages that a shader generates using shader logging.
//   - Developing Metal apps that run in Simulator: Prototype and test your Metal apps in Simulator.
//   - Improving your game’s graphics performance and settings: Fix performance glitches and develop default settings for smooth experiences on Apple platforms using the powerful suite of Metal development tools.
//   - Metal debugger: Debug and profile your Metal workload with a GPU trace.
//   - Metal developer workflows: Locate and fix issues related to your app’s use of the Metal API and GPU functions.
//   - GPU counters and counter sample buffers: Retrieve runtime data from a GPU device by sampling one or more of its counters. ([MTLCounterSet], [MTLCommonCounterSet], [MTLCounter], [MTLCommonCounter], [MTLCounterSampleBufferDescriptor])
//   - Metal debugging types: Create capture managers and capture scopes, and review a GPU device’s log after it runs a command buffer. ([MTLCaptureDescriptor], [MTLCaptureManager], [MTLCaptureDestination], [MTLCaptureScope], [MTLCaptureError])
//
// # Apple silicon
//
//   - Porting your Metal code to Apple silicon: Create a version of your Metal app that runs on both Apple silicon and Intel-based Mac computers.
//   - Tailor your apps for Apple GPUs and tile-based deferred rendering: Learn about characteristic Apple GPU features, including imageblocks, tile shaders, and raster order groups.
//
// # Variables
//
//   - MTLDeviceErrorDomain
//
// # Key Types
//
//   - [MTLRenderPipelineDescriptor] - An argument of options you pass to a GPU device to get a render pipeline state.
//   - [MTLMeshRenderPipelineDescriptor] - An object that configures new render pipeline state objects for mesh shading.
//   - [MTL4MeshRenderPipelineDescriptor] - Groups together properties you use to create a mesh render pipeline state object.
//   - [MTLRenderPassDescriptor] - A group of render targets that hold the results of a render pass.
//   - [MTLTextureDescriptor] - An instance that you use to configure new Metal texture instances.
//   - [MTLIndirectCommandBufferDescriptor] - A configuration you create to customize an indirect command buffer.
//   - [MTL4RenderPassDescriptor] - Describes a render pass.
//   - [MTL4RenderPipelineDescriptor] - Groups together properties to create a render pipeline state object.
//   - [MTLAccelerationStructureCurveGeometryDescriptor] - A descriptor you configure with curve geometry for building acceleration structures.
//   - [MTLSamplerDescriptor] - An object that you use to configure a texture sampler.
//
// [Metal Documentation]: https://developer.apple.com/documentation/Metal
package metal

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Metal library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Metal.framework/Metal",
	"/usr/lib/libMetal.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: Metal: failed to load framework from any known path\n")
}

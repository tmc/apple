// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objectivec"
)

// The main Metal interface to a GPU that apps use to draw graphics and run computations in parallel.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice
type MTLDevice interface {
	objectivec.IObject

	// The maximum number of concurrent compilation tasks the device is running.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/maximumConcurrentCompilationTaskCount
	MaximumConcurrentCompilationTaskCount() uint

	// A Boolean value that indicates whether the device uses additional CPU threads for compilation tasks.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/shouldMaximizeConcurrentCompilation
	ShouldMaximizeConcurrentCompilation() bool

	// FunctionHandleWithFunction protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/functionHandle(function:)-4bw39
	FunctionHandleWithFunction(function MTLFunction) MTLFunctionHandle

	// Get the function handle for the specified binary-linked function from the pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/functionHandle(function:)-w9ia
	FunctionHandleWithBinaryFunction(function MTL4BinaryFunction) MTLFunctionHandle

	// Creates a new archive from data available at an [NSURL] address.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeArchive(url:)
	NewArchiveWithURLError(url foundation.INSURL) (MTL4Archive, error)

	// Creates a new argument table from an argument table descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeArgumentTable(descriptor:)
	NewArgumentTableWithDescriptorError(descriptor IMTL4ArgumentTableDescriptor) (MTL4ArgumentTable, error)

	// Creates a new placement sparse buffer of a specific length.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeBuffer(length:options:placementSparsePageSize:)
	NewBufferWithLengthOptionsPlacementSparsePageSize(length uint, options MTLResourceOptions, placementSparsePageSize MTLSparsePageSize) MTLBuffer

	// Creates a new command allocator.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandAllocator()
	NewCommandAllocator() MTL4CommandAllocator

	// Creates a new command allocator from a command allocator descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandAllocator(descriptor:)
	NewCommandAllocatorWithDescriptorError(descriptor IMTL4CommandAllocatorDescriptor) (MTL4CommandAllocator, error)

	// Creates a new command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandBuffer()
	NewCommandBuffer() MTL4CommandBuffer

	// Creates a command queue with the provided configuration.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandQueue(descriptor:)
	NewCommandQueueWithDescriptor(descriptor IMTLCommandQueueDescriptor) MTLCommandQueue

	// Creates a new compiler from a compiler descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCompiler(descriptor:)
	NewCompilerWithDescriptorError(descriptor IMTL4CompilerDescriptor) (MTL4Compiler, error)

	// Creates a new counter heap configured from a counter heap descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCounterHeap(descriptor:)
	NewCounterHeapWithDescriptorError(descriptor IMTL4CounterHeapDescriptor) (MTL4CounterHeap, error)

	// Creates a shader log state with the provided configuration.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLogState(descriptor:)
	NewLogStateWithDescriptorError(descriptor IMTLLogStateDescriptor) (MTLLogState, error)

	// Creates a new command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeMTL4CommandQueue()
	NewMTL4CommandQueue() MTL4CommandQueue

	// Creates a new command queue from a queue descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeMTL4CommandQueue(descriptor:)
	NewMTL4CommandQueueWithDescriptorError(descriptor IMTL4CommandQueueDescriptor) (MTL4CommandQueue, error)

	// Creates a new pipeline data set serializer instance from a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makePipelineDataSetSerializer(descriptor:)
	NewPipelineDataSetSerializerWithDescriptor(descriptor IMTL4PipelineDataSetSerializerDescriptor) MTL4PipelineDataSetSerializer

	// Creates a tensor by allocating new memory.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeTensor(descriptor:)
	NewTensorWithDescriptorError(descriptor IMTLTensorDescriptor) (MTLTensor, error)

	// Creates a new texture view pool from a resource view pool descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeTextureViewPool(descriptor:)
	NewTextureViewPoolWithDescriptorError(descriptor IMTLResourceViewPoolDescriptor) (MTLTextureViewPool, error)

	// Queries the frequency of the GPU timestamp in ticks per second.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/queryTimestampFrequency()
	QueryTimestampFrequency() uint64

	// Returns the size, in bytes, of each entry in a counter heap of a specific counter heap type when your app resolves it into a usable format.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/size(ofCounterHeapEntry:)
	SizeOfCounterHeapEntry(type_ MTL4CounterHeapType) uint

	// Determines the size and alignment required to hold the data of a tensor you create with a descriptor in a buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/tensorSizeAndAlign(descriptor:)
	TensorSizeAndAlignWithDescriptor(descriptor IMTLTensorDescriptor) MTLSizeAndAlign

	// Returns the buffer sizes the GPU device needs to build, refit, and store an acceleration structure.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/accelerationStructureSizes(descriptor:)
	AccelerationStructureSizesWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLAccelerationStructureSizes

	// Retrieves the default sample positions for a specific sample count.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/getDefaultSamplePositions:count:
	GetDefaultSamplePositionsCount(positions *MTLSamplePosition, count uint)

	// Returns the size and alignment, in bytes, of an acceleration structure if you create it from a heap with a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/heapAccelerationStructureSizeAndAlign(descriptor:)
	HeapAccelerationStructureSizeAndAlignWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLSizeAndAlign

	// Returns the size and alignment, in bytes, of an acceleration structure if you create it from a heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/heapAccelerationStructureSizeAndAlign(size:)
	HeapAccelerationStructureSizeAndAlignWithSize(size uint) MTLSizeAndAlign

	// Returns the size and alignment, in bytes, of a buffer if you create it from a heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/heapBufferSizeAndAlign(length:options:)
	HeapBufferSizeAndAlignWithLengthOptions(length uint, options MTLResourceOptions) MTLSizeAndAlign

	// Returns the size and alignment, in bytes, of a texture if you create it from a heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/heapTextureSizeAndAlign(descriptor:)
	HeapTextureSizeAndAlignWithDescriptor(desc IMTLTextureDescriptor) MTLSizeAndAlign

	// Returns the minimum alignment the GPU device requires to create a linear texture from a buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/minimumLinearTextureAlignment(for:)
	MinimumLinearTextureAlignmentForPixelFormat(format MTLPixelFormat) uint

	// Returns the minimum alignment the GPU device requires to create a texture buffer from a buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/minimumTextureBufferAlignment(for:)
	MinimumTextureBufferAlignmentForPixelFormat(format MTLPixelFormat) uint

	// Creates a new ray-tracing acceleration structure from a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeAccelerationStructure(descriptor:)
	NewAccelerationStructureWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLAccelerationStructure

	// Creates a new acceleration structure with a specific size.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeAccelerationStructure(size:)
	NewAccelerationStructureWithSize(size uint) MTLAccelerationStructure

	// Creates a new argument encoder for an array of arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeArgumentEncoder(arguments:)
	NewArgumentEncoderWithArguments(arguments []MTLArgumentDescriptor) MTLArgumentEncoder

	// Creates a new argument encoder for a buffer binding.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeArgumentEncoder(bufferBinding:)
	NewArgumentEncoderWithBufferBinding(bufferBinding MTLBufferBinding) MTLArgumentEncoder

	// Creates a Metal binary archive instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeBinaryArchive(descriptor:)
	NewBinaryArchiveWithDescriptorError(descriptor IMTLBinaryArchiveDescriptor) (MTLBinaryArchive, error)

	// Allocates a new buffer of a given length and initializes its contents by copying existing data into it.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeBuffer(bytes:length:options:)
	NewBufferWithBytesLengthOptions(pointer unsafe.Pointer, length uint, options MTLResourceOptions) MTLBuffer

	// Creates a buffer that wraps an existing contiguous memory allocation.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeBuffer(bytesNoCopy:length:options:deallocator:)
	NewBufferWithBytesNoCopyLengthOptionsDeallocator(pointer unsafe.Pointer, length uint, options MTLResourceOptions, deallocator unsafe.Pointer) MTLBuffer

	// Creates a buffer the method clears with zero values.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeBuffer(length:options:)
	NewBufferWithLengthOptions(length uint, options MTLResourceOptions) MTLBuffer

	// Creates a queue you use to submit rendering and computation commands to a GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandQueue()
	NewCommandQueue() MTLCommandQueue

	// Creates a queue you use to submit rendering and computation commands to a GPU that has a fixed number of uncompleted command buffers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandQueue(maxCommandBufferCount:)
	NewCommandQueueWithMaxCommandBufferCount(maxCommandBufferCount uint) MTLCommandQueue

	// Asynchronously creates a compute pipeline state and reflection information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(descriptor:options:completionHandler:)
	NewComputePipelineStateWithDescriptorOptionsCompletionHandler(descriptor IMTLComputePipelineDescriptor, options MTLPipelineOption, completionHandler ErrorHandler)

	// Synchronously creates a compute pipeline state and reflection information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(descriptor:options:reflection:)
	NewComputePipelineStateWithDescriptorOptionsReflectionError(descriptor IMTLComputePipelineDescriptor, options MTLPipelineOption, reflection *MTLAutoreleasedComputePipelineReflection) (MTLComputePipelineState, error)

	// Asynchronously creates a compute pipeline state with a function instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(function:completionHandler:)
	NewComputePipelineStateWithFunctionCompletionHandler(computeFunction MTLFunction, completionHandler ErrorHandler)

	// Synchronously creates a compute pipeline state with a function instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(function:)
	NewComputePipelineStateWithFunctionError(computeFunction MTLFunction) (MTLComputePipelineState, error)

	// Asynchronously creates a compute pipeline state and reflection with a function instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(function:options:completionHandler:)
	NewComputePipelineStateWithFunctionOptionsCompletionHandler(computeFunction MTLFunction, options MTLPipelineOption, completionHandler ErrorHandler)

	// Synchronously creates a compute pipeline state and reflection with a function instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(function:options:reflection:)
	NewComputePipelineStateWithFunctionOptionsReflectionError(computeFunction MTLFunction, options MTLPipelineOption, reflection *MTLAutoreleasedComputePipelineReflection) (MTLComputePipelineState, error)

	// Creates a counter sample buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCounterSampleBuffer(descriptor:)
	NewCounterSampleBufferWithDescriptorError(descriptor IMTLCounterSampleBufferDescriptor) (MTLCounterSampleBuffer, error)

	// Creates a Metal library instance that contains the functions from your app’s default Metal library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeDefaultLibrary()
	NewDefaultLibrary() MTLLibrary

	// Creates a Metal library instance that contains the functions in a bundle’s default Metal library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeDefaultLibrary(bundle:)
	NewDefaultLibraryWithBundleError(bundle foundation.NSBundle) (MTLLibrary, error)

	// Creates a depth-stencil state instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeDepthStencilState(descriptor:)
	NewDepthStencilStateWithDescriptor(descriptor IMTLDepthStencilDescriptor) MTLDepthStencilState

	// Creates a Metal dynamic library instance from a Metal library instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeDynamicLibrary(library:)
	NewDynamicLibraryError(library MTLLibrary) (MTLDynamicLibrary, error)

	// Creates a Metal dynamic library instance that contains the functions in the Metal library file at a URL.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeDynamicLibrary(url:)
	NewDynamicLibraryWithURLError(url foundation.INSURL) (MTLDynamicLibrary, error)

	// Creates a new event instance that you can use to synchronize commands and resources within the same GPU device.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeEvent()
	NewEvent() MTLEvent

	// Creates a new memory fence instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeFence()
	NewFence() MTLFence

	// Creates a new GPU heap instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeHeap(descriptor:)
	NewHeapWithDescriptor(descriptor IMTLHeapDescriptor) MTLHeap

	// Creates an input/output command queue you use to submit commands that load assets from the file system into GPU resources or system memory.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeIOCommandQueue(descriptor:)
	NewIOCommandQueueWithDescriptorError(descriptor IMTLIOCommandQueueDescriptor) (MTLIOCommandQueue, error)

	// Creates an input/output file handle instance that represents a compressed file at a URL.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeIOFileHandle(url:compressionMethod:)
	NewIOFileHandleWithURLCompressionMethodError(url foundation.INSURL, compressionMethod MTLIOCompressionMethod) (MTLIOFileHandle, error)

	// Creates an input/output file handle instance that represents a file at a URL.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeIOFileHandle(url:)
	NewIOFileHandleWithURLError(url foundation.INSURL) (MTLIOFileHandle, error)

	// Creates an indirect command buffer instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeIndirectCommandBuffer(descriptor:maxCommandCount:options:)
	NewIndirectCommandBufferWithDescriptorMaxCommandCountOptions(descriptor IMTLIndirectCommandBufferDescriptor, maxCount uint, options MTLResourceOptions) MTLIndirectCommandBuffer

	// Creates a Metal library instance from a dispatch-data instance that contains the functions in a precompiled Metal library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(data:)
	NewLibraryWithDataError(data dispatch.Data) (MTLLibrary, error)

	// Asynchronously creates a Metal library instance by compiling the functions in a source string.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(source:options:completionHandler:)
	NewLibraryWithSourceOptionsCompletionHandler(source string, options IMTLCompileOptions, completionHandler ErrorHandler)

	// Synchronously creates a Metal library instance by compiling the functions in a source string.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(source:options:)
	NewLibraryWithSourceOptionsError(source string, options IMTLCompileOptions) (MTLLibrary, error)

	// Asynchronously creates a Metal library from the function stitching graphs in a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(stitchedDescriptor:completionHandler:)
	NewLibraryWithStitchedDescriptorCompletionHandler(descriptor IMTLStitchedLibraryDescriptor, completionHandler ErrorHandler)

	// Synchronously creates a Metal library from the function stitching graphs in a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(stitchedDescriptor:)
	NewLibraryWithStitchedDescriptorError(descriptor IMTLStitchedLibraryDescriptor) (MTLLibrary, error)

	// Creates a Metal library instance that contains the functions in the Metal library file at a URL.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(URL:)
	NewLibraryWithURLError(url foundation.INSURL) (MTLLibrary, error)

	// Creates a rasterization rate map instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRasterizationRateMap(descriptor:)
	NewRasterizationRateMapWithDescriptor(descriptor IMTLRasterizationRateMapDescriptor) MTLRasterizationRateMap

	// Asynchronously creates a render pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(descriptor:completionHandler:)
	NewRenderPipelineStateWithDescriptorCompletionHandler(descriptor IMTLRenderPipelineDescriptor, completionHandler ErrorHandler)

	// Synchronously creates a render pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(descriptor:)
	NewRenderPipelineStateWithDescriptorError(descriptor IMTLRenderPipelineDescriptor) (MTLRenderPipelineState, error)

	// Asynchronously creates a render pipeline state and reflection information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(descriptor:options:completionHandler:)-5gdww
	NewRenderPipelineStateWithDescriptorOptionsCompletionHandler(descriptor IMTLRenderPipelineDescriptor, options MTLPipelineOption, completionHandler ErrorHandler)

	// Synchronously creates a render pipeline state and reflection information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(descriptor:options:reflection:)
	NewRenderPipelineStateWithDescriptorOptionsReflectionError(descriptor IMTLRenderPipelineDescriptor, options MTLPipelineOption, reflection *MTLAutoreleasedRenderPipelineReflection) (MTLRenderPipelineState, error)

	// Asynchronously creates a mesh render pipeline state and reflection information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(descriptor:options:completionHandler:)-1wvya
	NewRenderPipelineStateWithMeshDescriptorOptionsCompletionHandler(descriptor IMTLMeshRenderPipelineDescriptor, options MTLPipelineOption, completionHandler ErrorHandler)

	// Synchronously creates a mesh render pipeline state and reflection information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/newRenderPipelineStateWithMeshDescriptor:options:reflection:error:
	NewRenderPipelineStateWithMeshDescriptorOptionsReflectionError(descriptor IMTLMeshRenderPipelineDescriptor, options MTLPipelineOption, reflection *MTLAutoreleasedRenderPipelineReflection) (MTLRenderPipelineState, error)

	// Asynchronously creates a tile shader’s render pipeline state and reflection information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(tileDescriptor:options:completionHandler:)
	NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler(descriptor IMTLTileRenderPipelineDescriptor, options MTLPipelineOption, completionHandler ErrorHandler)

	// Synchronously creates a tile shader’s render pipeline state and reflection information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(tileDescriptor:options:reflection:)
	NewRenderPipelineStateWithTileDescriptorOptionsReflectionError(descriptor IMTLTileRenderPipelineDescriptor, options MTLPipelineOption, reflection *MTLAutoreleasedRenderPipelineReflection) (MTLRenderPipelineState, error)

	// Creates a residency set, which can move resources in and out of memory residency.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeResidencySet(descriptor:)
	NewResidencySetWithDescriptorError(desc IMTLResidencySetDescriptor) (MTLResidencySet, error)

	// Creates a sampler state instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeSamplerState(descriptor:)
	NewSamplerStateWithDescriptor(descriptor IMTLSamplerDescriptor) MTLSamplerState

	// Creates a new shared event instance that you can use to synchronize commands and resources across different GPU devices.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeSharedEvent()
	NewSharedEvent() MTLSharedEvent

	// Recreates a shared event from a handle.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeSharedEvent(handle:)
	NewSharedEventWithHandle(sharedEventHandle IMTLSharedEventHandle) MTLSharedEvent

	// Creates a texture that you can share across process boundaries.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeSharedTexture(descriptor:)
	NewSharedTextureWithDescriptor(descriptor IMTLTextureDescriptor) MTLTexture

	// Creates a texture that references a shared texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeSharedTexture(handle:)
	NewSharedTextureWithHandle(sharedHandle IMTLSharedTextureHandle) MTLTexture

	// Creates a new texture instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeTexture(descriptor:)
	NewTextureWithDescriptor(descriptor IMTLTextureDescriptor) MTLTexture

	// Creates a texture instance that uses I/O surface to store its underlying data.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeTexture(descriptor:iosurface:plane:)
	NewTextureWithDescriptorIosurfacePlane(descriptor IMTLTextureDescriptor, iosurface iosurface.IOSurfaceRef, plane uint) MTLTexture

	// Captures and returns a CPU timestamp and a GPU timestamp from the same moment in time.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/sampleTimestamps:gpuTimestamp:
	SampleTimestampsGpuTimestamp(cpuTimestamp *MTLTimestamp, gpuTimestamp *MTLTimestamp)

	// Returns the size, in bytes, of a sparse tile the GPU device creates with a specific page size.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/sparseTileSizeInBytes(sparsePageSize:)
	SparseTileSizeInBytesForSparsePageSize(sparsePageSize MTLSparsePageSize) uint

	// Returns the dimensions of a sparse tile for a texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/sparseTileSize(with:pixelFormat:sampleCount:)
	SparseTileSizeWithTextureTypePixelFormatSampleCount(textureType MTLTextureType, pixelFormat MTLPixelFormat, sampleCount uint) MTLSize

	// Returns the dimensions of a sparse tile for a texture that has a specific sparse page size.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/sparseTileSize(textureType:pixelFormat:sampleCount:sparsePageSize:)
	SparseTileSizeWithTextureTypePixelFormatSampleCountSparsePageSize(textureType MTLTextureType, pixelFormat MTLPixelFormat, sampleCount uint, sparsePageSize MTLSparsePageSize) MTLSize

	// Returns a Boolean value that indicates whether you can read GPU counters at the specified command boundary.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsCounterSampling(_:)
	SupportsCounterSampling(samplingPoint MTLCounterSamplingPoint) bool

	// Returns a Boolean value that indicates whether the GPU device supports the feature set of a specific GPU family.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsFamily(_:)
	SupportsFamily(gpuFamily MTLGPUFamily) bool

	// Returns a Boolean value that indicates whether the GPU can create a rasterization rate map with a specific number of layers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsRasterizationRateMap(layerCount:)
	SupportsRasterizationRateMapWithLayerCount(layerCount uint) bool

	// Returns a Boolean value that indicates whether the GPU can sample a texture with a specific number of sample points.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsTextureSampleCount(_:)
	SupportsTextureSampleCount(sampleCount uint) bool

	// Returns a Boolean value that indicates whether the GPU supports an amplification factor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsVertexAmplificationCount(_:)
	SupportsVertexAmplificationCount(count uint) bool

	// A Boolean value that indicates whether the device uses additional CPU threads for compilation tasks.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/shouldMaximizeConcurrentCompilation
	SetShouldMaximizeConcurrentCompilation(value bool)

	// The architectural details of the GPU device.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/architecture
	Architecture() IMTLArchitecture

	// Returns the GPU device’s support tier for argument buffers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/argumentBuffersSupport
	ArgumentBuffersSupport() MTLArgumentBuffersTier

	// The counter sets supported by the device object.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/counterSets
	CounterSets() []objectivec.IObject

	// The total amount of memory, in bytes, the GPU device is using for all of its resources.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/currentAllocatedSize
	CurrentAllocatedSize() uint

	// A Boolean value that indicates whether a device supports a packed depth-and-stencil pixel format.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/isDepth24Stencil8PixelFormatSupported
	Depth24Stencil8PixelFormatSupported() bool

	// A Boolean value that indicates whether the GPU shares all of its memory with the CPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/hasUnifiedMemory
	HasUnifiedMemory() bool

	// A Boolean value that indicates whether a GPU device doesn’t have a connection to a display.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/isHeadless
	Headless() bool

	// The physical location of the GPU relative to the system.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/location
	Location() MTLDeviceLocation

	// A specific GPU position based on its general location.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/locationNumber
	LocationNumber() uint

	// A Boolean value that indicates whether the GPU lowers its performance to conserve energy.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/isLowPower
	LowPower() bool

	// The maximum number of unique argument buffer samplers per app.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/maxArgumentBufferSamplerCount
	MaxArgumentBufferSamplerCount() uint

	// The largest amount of memory, in bytes, that a GPU device can allocate to a buffer instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/maxBufferLength
	MaxBufferLength() uint

	// The maximum threadgroup memory available to a compute kernel, in bytes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/maxThreadgroupMemoryLength
	MaxThreadgroupMemoryLength() uint

	// The maximum number of threads along each dimension of a threadgroup.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/maxThreadsPerThreadgroup
	MaxThreadsPerThreadgroup() MTLSize

	// The highest theoretical rate, in bytes per second, the system can copy between system memory and the GPU’s dedicated memory (VRAM).
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/maxTransferRate
	MaxTransferRate() uint64

	// The full name of the GPU device.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/name
	Name() string

	// The total number of GPUs in the peer group, if applicable.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/peerCount
	PeerCount() uint32

	// The peer group ID the GPU belongs to, if applicable.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/peerGroupID
	PeerGroupID() uint64

	// The unique identifier for a GPU in a peer group.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/peerIndex
	PeerIndex() uint32

	// A Boolean value that indicates whether the GPU supports programmable sample positions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/areProgrammableSamplePositionsSupported
	ProgrammableSamplePositionsSupported() bool

	// A Boolean value that indicates whether the GPU supports raster order groups.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/areRasterOrderGroupsSupported
	RasterOrderGroupsSupported() bool

	// The GPU device’s texture support tier.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/readWriteTextureSupport
	ReadWriteTextureSupport() MTLReadWriteTextureTier

	// An approximation of how much memory, in bytes, this GPU device can allocate without affecting its runtime performance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/recommendedMaxWorkingSetSize
	RecommendedMaxWorkingSetSize() uint64

	// The GPU device’s registry identifier.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/registryID
	RegistryID() uint64

	// A Boolean value that indicates whether the GPU is removable.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/isRemovable
	Removable() bool

	// Returns the size, in bytes, of a sparse tile the GPU device creates using a default page size.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/sparseTileSizeInBytes
	SparseTileSizeInBytes() uint

	// A Boolean value that indicates whether the GPU can filter a texture with a 32-bit floating-point format.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supports32BitFloatFiltering
	Supports32BitFloatFiltering() bool

	// A Boolean value that indicates whether the GPU can allocate 32-bit integer texture formats and resolve to 32-bit floating-point texture formats.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supports32BitMSAA
	Supports32BitMSAA() bool

	// A Boolean value that indicates whether you can use textures that use BC compression.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsBCTextureCompression
	SupportsBCTextureCompression() bool

	// A Boolean value that indicates whether the GPU device can create and use dynamic libraries in compute pipelines.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsDynamicLibraries
	SupportsDynamicLibraries() bool

	// A Boolean value that indicates whether the device supports function pointers in compute kernel functions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsFunctionPointers
	SupportsFunctionPointers() bool

	// A Boolean value that indicates whether the device supports function pointers in render functions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsFunctionPointersFromRender
	SupportsFunctionPointersFromRender() bool

	// A Boolean value that indicates whether the device supports placement sparse resources.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsPlacementSparse
	SupportsPlacementSparse() bool

	// A Boolean value that indicates whether the GPU device supports motion blur for ray tracing.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsPrimitiveMotionBlur
	SupportsPrimitiveMotionBlur() bool

	// A Boolean value that indicates whether the GPU can compute multiple interpolations of a fragment function’s input.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsPullModelInterpolation
	SupportsPullModelInterpolation() bool

	// A Boolean value that indicates whether you can query the texture level of detail from within a shader.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsQueryTextureLOD
	SupportsQueryTextureLOD() bool

	// A Boolean value that indicates whether the GPU device supports ray tracing.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsRaytracing
	SupportsRaytracing() bool

	// A Boolean value that indicates whether you can call ray-tracing functions from a vertex or fragment shader.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsRaytracingFromRender
	SupportsRaytracingFromRender() bool

	// A Boolean value that indicates whether the GPU device can create and use dynamic libraries in render pipelines.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsRenderDynamicLibraries
	SupportsRenderDynamicLibraries() bool

	// A Boolean value that indicates whether the GPU supports barycentric coordinates.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsShaderBarycentricCoordinates
	SupportsShaderBarycentricCoordinates() bool
}

// MTLDeviceObject wraps an existing Objective-C object that conforms to the MTLDevice protocol.
type MTLDeviceObject struct {
	objectivec.Object
}
func (o MTLDeviceObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLDeviceObjectFromID constructs a [MTLDeviceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLDeviceObjectFromID(id objc.ID) MTLDeviceObject {
	return MTLDeviceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The maximum number of concurrent compilation tasks the device is running.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/maximumConcurrentCompilationTaskCount

func (o MTLDeviceObject) MaximumConcurrentCompilationTaskCount() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("maximumConcurrentCompilationTaskCount"))
	return rv
	}

// A Boolean value that indicates whether the device uses additional CPU
// threads for compilation tasks.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/shouldMaximizeConcurrentCompilation

func (o MTLDeviceObject) ShouldMaximizeConcurrentCompilation() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("shouldMaximizeConcurrentCompilation"))
	return rv
	}

//
// # Discussion
// 
// Returns the function handle for a function that was compiled with
// MTLFunctionOptionPipelineIndependent and MTLFunctionOptionCompileToBinary.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/functionHandle(function:)-4bw39

func (o MTLDeviceObject) FunctionHandleWithFunction(function MTLFunction) MTLFunctionHandle {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionHandleWithFunction:"), function)
	return MTLFunctionHandleObjectFromID(rv)
	}

// Get the function handle for the specified binary-linked function from the
// pipeline state.
//
// function: A [MTL4BinaryFunction] instance representing the function binary.
//
// # Return Value
// 
// A [MTLFunctionHandle] instance for a binary function that was compiled with
// [MTLFunctionOptionPipelineIndependent], otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/functionHandle(function:)-w9ia

func (o MTLDeviceObject) FunctionHandleWithBinaryFunction(function MTL4BinaryFunction) MTLFunctionHandle {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionHandleWithBinaryFunction:"), function)
	return MTLFunctionHandleObjectFromID(rv)
	}

// Creates a new archive from data available at an [NSURL] address.
//
// url: An [NSURL] instance that represents the path from which the device loads
// the [MTL4Archive].
//
// # Return Value
// 
// A [MTL4Archive] instance, or `nil` if the function failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeArchive(url:)

func (o MTLDeviceObject) NewArchiveWithURLError(url foundation.INSURL) (MTL4Archive, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newArchiveWithURL:error:"), url)
	if err != nil {
		return nil, err
	}
	return MTL4ArchiveObjectFromID(rv), nil
	}

// Creates a new argument table from an argument table descriptor.
//
// descriptor: A [MTL4ArgumentTableDescriptor] instance that configures the
// [MTL4ArgumentTable] instance.
//
// # Return Value
// 
// A [MTL4ArgumentTable] instance, or `nil` if the function failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeArgumentTable(descriptor:)

func (o MTLDeviceObject) NewArgumentTableWithDescriptorError(descriptor IMTL4ArgumentTableDescriptor) (MTL4ArgumentTable, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newArgumentTableWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTL4ArgumentTableObjectFromID(rv), nil
	}

// Creates a new placement sparse buffer of a specific length.
//
// length: The size of the [MTLBuffer], in bytes.
//
// options: A [MTLResourceOptions] instance that establishes the buffer’s storage
// modes.
// //
// [MTLResourceOptions]: https://developer.apple.com/documentation/Metal/MTLResourceOptions
//
// placementSparsePageSize: [MTLSparsePageSize] to use for the placement sparse buffer.
// //
// [MTLSparsePageSize]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize
//
// # Return Value
// 
// A [MTLBuffer] instance, or `nil` if the function failed.
//
// # Discussion
// 
// This method creates a new placement sparse [MTLBuffer] of a specific
// length. You assign memory to placement sparse buffers using a [MTLHeap] of
// type [HeapTypePlacement].
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeBuffer(length:options:placementSparsePageSize:)

func (o MTLDeviceObject) NewBufferWithLengthOptionsPlacementSparsePageSize(length uint, options MTLResourceOptions, placementSparsePageSize MTLSparsePageSize) MTLBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newBufferWithLength:options:placementSparsePageSize:"), length, options, placementSparsePageSize)
	return MTLBufferObjectFromID(rv)
	}

// Creates a new command allocator.
//
// # Return Value
// 
// A [MTL4CommandAllocator] instance, or `nil` if the function failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandAllocator()

func (o MTLDeviceObject) NewCommandAllocator() MTL4CommandAllocator {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newCommandAllocator"))
	return MTL4CommandAllocatorObjectFromID(rv)
	}

// Creates a new command allocator from a command allocator descriptor.
//
// descriptor: A [MTL4CommandAllocatorDescriptor] instance that configures the
// [MTL4CommandAllocator] instance.
//
// # Return Value
// 
// A [MTL4CommandAllocator] instance, or `nil` if the function failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandAllocator(descriptor:)

func (o MTLDeviceObject) NewCommandAllocatorWithDescriptorError(descriptor IMTL4CommandAllocatorDescriptor) (MTL4CommandAllocator, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newCommandAllocatorWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTL4CommandAllocatorObjectFromID(rv), nil
	}

// Creates a new command buffer.
//
// # Return Value
// 
// A [MTL4CommandBuffer] instance, or `nil` if the function failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandBuffer()

func (o MTLDeviceObject) NewCommandBuffer() MTL4CommandBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newCommandBuffer"))
	return MTL4CommandBufferObjectFromID(rv)
	}

// Creates a command queue with the provided configuration.
//
// descriptor: The configuration for the new command queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandQueue(descriptor:)

func (o MTLDeviceObject) NewCommandQueueWithDescriptor(descriptor IMTLCommandQueueDescriptor) MTLCommandQueue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newCommandQueueWithDescriptor:"), descriptor)
	return MTLCommandQueueObjectFromID(rv)
	}

// Creates a new compiler from a compiler descriptor.
//
// descriptor: A [MTL4CompilerDescriptor] instance that configures the [MTL4Compiler]
// instance.
//
// # Return Value
// 
// A [MTL4Compiler] instance, or `nil` if the function failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCompiler(descriptor:)

func (o MTLDeviceObject) NewCompilerWithDescriptorError(descriptor IMTL4CompilerDescriptor) (MTL4Compiler, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newCompilerWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTL4CompilerObjectFromID(rv), nil
	}

// Creates a new counter heap configured from a counter heap descriptor.
//
// descriptor: [MTL4CounterHeapDescriptor] instance that configures the [MTL4CounterHeap]
// instance.
//
// # Return Value
// 
// A [MTL4CounterHeap] instance, or `nil` if the function failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCounterHeap(descriptor:)

func (o MTLDeviceObject) NewCounterHeapWithDescriptorError(descriptor IMTL4CounterHeapDescriptor) (MTL4CounterHeap, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newCounterHeapWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTL4CounterHeapObjectFromID(rv), nil
	}

// Creates a shader log state with the provided configuration.
//
// descriptor: The configuration for the new shader log state.
//
// # Return Value
// 
// A new [MTLLogState] instance if the method completes successfully;
// otherwise Swift throws an error and Objective-C returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLogState(descriptor:)

func (o MTLDeviceObject) NewLogStateWithDescriptorError(descriptor IMTLLogStateDescriptor) (MTLLogState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newLogStateWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLLogStateObjectFromID(rv), nil
	}

// Creates a new command queue.
//
// # Return Value
// 
// A [MTL4CommandQueue] instance, or `nil` if the function failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeMTL4CommandQueue()

func (o MTLDeviceObject) NewMTL4CommandQueue() MTL4CommandQueue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newMTL4CommandQueue"))
	return MTL4CommandQueueObjectFromID(rv)
	}

// Creates a new command queue from a queue descriptor.
//
// descriptor: A [MTL4CommandQueueDescriptor] instance that configures the
// [MTL4CommandQueue] instance.
//
// # Return Value
// 
// A [MTL4CommandQueue] instance, or `nil` if the function failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeMTL4CommandQueue(descriptor:)

func (o MTLDeviceObject) NewMTL4CommandQueueWithDescriptorError(descriptor IMTL4CommandQueueDescriptor) (MTL4CommandQueue, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newMTL4CommandQueueWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTL4CommandQueueObjectFromID(rv), nil
	}

// Creates a new pipeline data set serializer instance from a descriptor.
//
// descriptor: A [MTL4PipelineDataSetSerializerDescriptor] instance that configures the
// new [MTL4PipelineDataSetSerializer] instance.
//
// # Return Value
// 
// A [MTL4PipelineDataSetSerializer] instance, or `nil` if the function
// failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makePipelineDataSetSerializer(descriptor:)

func (o MTLDeviceObject) NewPipelineDataSetSerializerWithDescriptor(descriptor IMTL4PipelineDataSetSerializerDescriptor) MTL4PipelineDataSetSerializer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newPipelineDataSetSerializerWithDescriptor:"), descriptor)
	return MTL4PipelineDataSetSerializerObjectFromID(rv)
	}

// Creates a tensor by allocating new memory.
//
// descriptor: A description of the properties for the new tensor.
//
// # Return Value
// 
// A new tensor instance that Metal configures using `descriptor` or `nil` if
// an error occurred.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeTensor(descriptor:)

func (o MTLDeviceObject) NewTensorWithDescriptorError(descriptor IMTLTensorDescriptor) (MTLTensor, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newTensorWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLTensorObjectFromID(rv), nil
	}

// Creates a new texture view pool from a resource view pool descriptor.
//
// descriptor: A [MTLResourceViewPoolDescriptor] instance that configures the
// [MTLTextureViewPool] instance.
//
// # Return Value
// 
// A [MTLTextureViewPool] instance, or `nil` if the function failed.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeTextureViewPool(descriptor:)

func (o MTLDeviceObject) NewTextureViewPoolWithDescriptorError(descriptor IMTLResourceViewPoolDescriptor) (MTLTextureViewPool, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newTextureViewPoolWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLTextureViewPoolObjectFromID(rv), nil
	}

// Queries the frequency of the GPU timestamp in ticks per second.
//
// # Return Value
// 
// The frequency of the GPU timestamp in ticks per second.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/queryTimestampFrequency()

func (o MTLDeviceObject) QueryTimestampFrequency() uint64 {
	
	rv := objc.Send[uint64](o.ID, objc.Sel("queryTimestampFrequency"))
	return rv
	}

// Returns the size, in bytes, of each entry in a counter heap of a specific
// counter heap type when your app resolves it into a usable format.
//
// type: [MTL4CounterHeapType] value that represents the type of the
// [MTL4CounterHeap] to resolve.
// //
// [MTL4CounterHeapType]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapType
//
// # Return Value
// 
// The size of the post-transformation entry from a [MTL4CounterHeap] of type
// [MTL4CounterHeapType].
//
// [MTL4CounterHeapType]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapType
//
// # Discussion
// 
// In order to use the data available in a [MTL4CounterHeap], your app first
// resolves it either in the CPU timeline or in the GPU timeline. When your
// app calls [ResolveCounterHeapWithRangeIntoBufferWaitFenceUpdateFence] to
// resolve counter data in the GPU timeline, Metal writes the data into a
// [MTLBuffer].
// 
// During this process, Metal transform the data in the heap into a format
// consisting of entries of the size that this method advertises, based on the
// [MTL4CounterHeapType].
//
// [MTL4CounterHeapType]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapType
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/size(ofCounterHeapEntry:)

func (o MTLDeviceObject) SizeOfCounterHeapEntry(type_ MTL4CounterHeapType) uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("sizeOfCounterHeapEntry:"), type_)
	return rv
	}

// Determines the size and alignment required to hold the data of a tensor you
// create with a descriptor in a buffer.
//
// descriptor: A description of the properties for the new tensor.
//
// # Return Value
// 
// The size and alignment required to hold the data of a tensor you create
// with `descriptor` in a buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/tensorSizeAndAlign(descriptor:)

func (o MTLDeviceObject) TensorSizeAndAlignWithDescriptor(descriptor IMTLTensorDescriptor) MTLSizeAndAlign {
	
	rv := objc.Send[MTLSizeAndAlign](o.ID, objc.Sel("tensorSizeAndAlignWithDescriptor:"), descriptor)
	return rv
	}

// Returns the buffer sizes the GPU device needs to build, refit, and store an
// acceleration structure.
//
// descriptor: An [MTLAccelerationStructureDescriptor] instance.
//
// # Return Value
// 
// A new [MTLAccelerationStructureSizes] instance.
//
// [MTLAccelerationStructureSizes]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureSizes
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/accelerationStructureSizes(descriptor:)

func (o MTLDeviceObject) AccelerationStructureSizesWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLAccelerationStructureSizes {
	
	rv := objc.Send[MTLAccelerationStructureSizes](o.ID, objc.Sel("accelerationStructureSizesWithDescriptor:"), descriptor)
	return rv
	}

// Retrieves the default sample positions for a specific sample count.
//
// positions: A pointer to a destination C array of [MTLSamplePosition] instances —
// with at least `count` elements — the method writes the default positions
// to.
// //
// [MTLSamplePosition]: https://developer.apple.com/documentation/Metal/MTLSamplePosition
//
// count: The number of points a GPU can sample from a texture. Ensure the GPU can
// support the `count` value by first calling the device’s
// [SupportsTextureSampleCount] method.
//
// # Discussion
// 
// The default sample positions are the same on all GPUs that support
// programmable sample positions (see
// [areProgrammableSamplePositionsSupported]).
// 
// The default sample position for GPUs that can sample one time is at the
// pixel’s center.
// 
// [positioning-samples-programmatically-2]
// 
// The default sample positions for GPUs that can sample two times have
// locations in the center of the pixel’s second quadrant and fourth
// quadrants.
// 
// [getDefaultSamplePositions-2]
// 
// The default sample positions for GPUs that can sample four times have one
// location in each of the pixel’s quadrants. Each location is at the center
// of one of that quadrant’s subquadrants.
// 
// [getDefaultSamplePositions-3]
// 
// The default sample positions for GPUs that can sample eight times have two
// locations in each of the pixel’s quadrants.
// 
// [getDefaultSamplePositions-4]
// 
// The table lists the indices and default locations for GPUs that support 1,
// 2, 4, or 8 sample positions.
// 
// [Table data omitted]
//
// [areProgrammableSamplePositionsSupported]: https://developer.apple.com/documentation/Metal/MTLDevice/areProgrammableSamplePositionsSupported
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/getDefaultSamplePositions:count:

func (o MTLDeviceObject) GetDefaultSamplePositionsCount(positions *MTLSamplePosition, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("getDefaultSamplePositions:count:"), objc.CArray(positions), count)
	}

// Returns the size and alignment, in bytes, of an acceleration structure if
// you create it from a heap with a descriptor.
//
// descriptor: An [MTLAccelerationStructureDescriptor] instance.
//
// # Return Value
// 
// An [MTLSizeAndAlign] instance.
//
// [MTLSizeAndAlign]: https://developer.apple.com/documentation/Metal/MTLSizeAndAlign
//
// # Discussion
// 
// Use this method to help estimate an appropriate size for a new heap before
// you create it.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/heapAccelerationStructureSizeAndAlign(descriptor:)

func (o MTLDeviceObject) HeapAccelerationStructureSizeAndAlignWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLSizeAndAlign {
	
	rv := objc.Send[MTLSizeAndAlign](o.ID, objc.Sel("heapAccelerationStructureSizeAndAlignWithDescriptor:"), descriptor)
	return rv
	}

// Returns the size and alignment, in bytes, of an acceleration structure if
// you create it from a heap.
//
// size: The size of an acceleration structure, in bytes.
//
// # Return Value
// 
// An [MTLSizeAndAlign] instance
//
// [MTLSizeAndAlign]: https://developer.apple.com/documentation/Metal/MTLSizeAndAlign
//
// # Discussion
// 
// Use this method to help estimate an appropriate size for a new heap before
// you create it.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/heapAccelerationStructureSizeAndAlign(size:)

func (o MTLDeviceObject) HeapAccelerationStructureSizeAndAlignWithSize(size uint) MTLSizeAndAlign {
	
	rv := objc.Send[MTLSizeAndAlign](o.ID, objc.Sel("heapAccelerationStructureSizeAndAlignWithSize:"), size)
	return rv
	}

// Returns the size and alignment, in bytes, of a buffer if you create it from
// a heap.
//
// length: The size of the buffer, in bytes.
//
// options: An [MTLResourceOptions] instance for a would-be buffer’s storage and
// hazard tracking modes. See [Resource fundamentals] and [Setting resource
// storage modes] for more information.
// //
// [MTLResourceOptions]: https://developer.apple.com/documentation/Metal/MTLResourceOptions
// [Resource fundamentals]: https://developer.apple.com/documentation/Metal/resource-fundamentals
// [Setting resource storage modes]: https://developer.apple.com/documentation/Metal/setting-resource-storage-modes
//
// # Return Value
// 
// An [MTLSizeAndAlign] instance.
//
// [MTLSizeAndAlign]: https://developer.apple.com/documentation/Metal/MTLSizeAndAlign
//
// # Discussion
// 
// Use this method to help estimate an appropriate size for a new heap before
// you create it.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/heapBufferSizeAndAlign(length:options:)

func (o MTLDeviceObject) HeapBufferSizeAndAlignWithLengthOptions(length uint, options MTLResourceOptions) MTLSizeAndAlign {
	
	rv := objc.Send[MTLSizeAndAlign](o.ID, objc.Sel("heapBufferSizeAndAlignWithLength:options:"), length, options)
	return rv
	}

// Returns the size and alignment, in bytes, of a texture if you create it
// from a heap.
//
// desc: An [MTLTextureDescriptor] instance.
//
// # Return Value
// 
// An [MTLSizeAndAlign] instance.
//
// [MTLSizeAndAlign]: https://developer.apple.com/documentation/Metal/MTLSizeAndAlign
//
// # Discussion
// 
// Use this method to help estimate an appropriate size for a new heap before
// you create it.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/heapTextureSizeAndAlign(descriptor:)

func (o MTLDeviceObject) HeapTextureSizeAndAlignWithDescriptor(desc IMTLTextureDescriptor) MTLSizeAndAlign {
	
	rv := objc.Send[MTLSizeAndAlign](o.ID, objc.Sel("heapTextureSizeAndAlignWithDescriptor:"), desc)
	return rv
	}

// Returns the minimum alignment the GPU device requires to create a linear
// texture from a buffer.
//
// format: An [MTLPixelFormat] instance that can’t be any of the depth, stencil, or
// compressed pixel formats.
// //
// [MTLPixelFormat]: https://developer.apple.com/documentation/Metal/MTLPixelFormat
//
// # Discussion
// 
// Metal aligns linear textures to their minimum alignment value, which
// directly affects the [NewTextureWithDescriptorOffsetBytesPerRow] method’s
// `offset` and `bytesPerRow` parameters.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/minimumLinearTextureAlignment(for:)

func (o MTLDeviceObject) MinimumLinearTextureAlignmentForPixelFormat(format MTLPixelFormat) uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("minimumLinearTextureAlignmentForPixelFormat:"), format)
	return rv
	}

// Returns the minimum alignment the GPU device requires to create a texture
// buffer from a buffer.
//
// format: An [MTLPixelFormat] instance.
// //
// [MTLPixelFormat]: https://developer.apple.com/documentation/Metal/MTLPixelFormat
//
// # Discussion
// 
// Metal aligns textures to their minimum alignment value, which directly
// affects the [NewTextureWithDescriptorOffsetBytesPerRow] method’s `offset`
// and `bytesPerRow` parameters.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/minimumTextureBufferAlignment(for:)

func (o MTLDeviceObject) MinimumTextureBufferAlignmentForPixelFormat(format MTLPixelFormat) uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("minimumTextureBufferAlignmentForPixelFormat:"), format)
	return rv
	}

// Creates a new ray-tracing acceleration structure from a descriptor.
//
// descriptor: An [MTLAccelerationStructureDescriptor] instance.
//
// # Return Value
// 
// A new [MTLAccelerationStructure] instance if the method completed
// successfully; otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeAccelerationStructure(descriptor:)

func (o MTLDeviceObject) NewAccelerationStructureWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLAccelerationStructure {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newAccelerationStructureWithDescriptor:"), descriptor)
	return MTLAccelerationStructureObjectFromID(rv)
	}

// Creates a new acceleration structure with a specific size.
//
// size: The size of the new acceleration structure, in bytes.
//
// # Return Value
// 
// A new [MTLAccelerationStructure] instance if the method completed
// successfully; otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeAccelerationStructure(size:)

func (o MTLDeviceObject) NewAccelerationStructureWithSize(size uint) MTLAccelerationStructure {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newAccelerationStructureWithSize:"), size)
	return MTLAccelerationStructureObjectFromID(rv)
	}

// Creates a new argument encoder for an array of arguments.
//
// arguments: An array of [MTLArgumentDescriptor] instances that you need to sort by
// their [Index] properties in monotonically increasing order.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeArgumentEncoder(arguments:)

func (o MTLDeviceObject) NewArgumentEncoderWithArguments(arguments []MTLArgumentDescriptor) MTLArgumentEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newArgumentEncoderWithArguments:"), objectivec.IObjectSliceToNSArray(arguments))
	return MTLArgumentEncoderObjectFromID(rv)
	}

// Creates a new argument encoder for a buffer binding.
//
// bufferBinding: An [MTLBufferBinding] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeArgumentEncoder(bufferBinding:)

func (o MTLDeviceObject) NewArgumentEncoderWithBufferBinding(bufferBinding MTLBufferBinding) MTLArgumentEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newArgumentEncoderWithBufferBinding:"), bufferBinding)
	return MTLArgumentEncoderObjectFromID(rv)
	}

// Creates a Metal binary archive instance.
//
// descriptor: An [MTLBinaryArchiveDescriptor] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeBinaryArchive(descriptor:)

func (o MTLDeviceObject) NewBinaryArchiveWithDescriptorError(descriptor IMTLBinaryArchiveDescriptor) (MTLBinaryArchive, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newBinaryArchiveWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLBinaryArchiveObjectFromID(rv), nil
	}

// Allocates a new buffer of a given length and initializes its contents by
// copying existing data into it.
//
// pointer: A pointer to the starting memory address the method copies the
// initialization data from.
//
// length: The size of the new buffer, in bytes, and the number of bytes the method
// copies from `pointer`.
//
// options: An [MTLResourceOptions] instance that sets the buffer’s storage and
// hazard-tracking modes. See [Resource fundamentals] and [Setting resource
// storage modes] for more information.
// //
// [MTLResourceOptions]: https://developer.apple.com/documentation/Metal/MTLResourceOptions
// [Resource fundamentals]: https://developer.apple.com/documentation/Metal/resource-fundamentals
// [Setting resource storage modes]: https://developer.apple.com/documentation/Metal/setting-resource-storage-modes
//
// # Return Value
// 
// A new [MTLBuffer] instance if the method completes successfully; otherwise
// `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeBuffer(bytes:length:options:)

func (o MTLDeviceObject) NewBufferWithBytesLengthOptions(pointer unsafe.Pointer, length uint, options MTLResourceOptions) MTLBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newBufferWithBytes:length:options:"), pointer, length, options)
	return MTLBufferObjectFromID(rv)
	}

// Creates a buffer that wraps an existing contiguous memory allocation.
//
// pointer: A page-aligned pointer to the starting memory address.
//
// length: The size of the new buffer, in bytes, that results in a page-aligned region
// of memory.
//
// options: An [MTLResourceOptions] instance that sets the buffer’s storage and
// hazard-tracking modes. See [Resource fundamentals] and [Setting resource
// storage modes] for more information.
// //
// [MTLResourceOptions]: https://developer.apple.com/documentation/Metal/MTLResourceOptions
// [Resource fundamentals]: https://developer.apple.com/documentation/Metal/resource-fundamentals
// [Setting resource storage modes]: https://developer.apple.com/documentation/Metal/setting-resource-storage-modes
//
// deallocator: A block the framework invokes when it deallocates the buffer so that your
// app can release the underlying memory; otherwise `nil` to opt out.
//
// # Return Value
// 
// A new [MTLBuffer] instance if the method completes successfully; otherwise
// `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeBuffer(bytesNoCopy:length:options:deallocator:)

func (o MTLDeviceObject) NewBufferWithBytesNoCopyLengthOptionsDeallocator(pointer unsafe.Pointer, length uint, options MTLResourceOptions, deallocator unsafe.Pointer) MTLBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newBufferWithBytesNoCopy:length:options:deallocator:"), pointer, length, options, deallocator)
	return MTLBufferObjectFromID(rv)
	}

// Creates a buffer the method clears with zero values.
//
// length: The size of the new buffer, in bytes.
//
// options: An [MTLResourceOptions] instance that sets the buffer’s storage and
// hazard-tracking modes. See [Resource fundamentals] and [Setting resource
// storage modes] for more information.
// //
// [MTLResourceOptions]: https://developer.apple.com/documentation/Metal/MTLResourceOptions
// [Resource fundamentals]: https://developer.apple.com/documentation/Metal/resource-fundamentals
// [Setting resource storage modes]: https://developer.apple.com/documentation/Metal/setting-resource-storage-modes
//
// # Return Value
// 
// A new [MTLBuffer] instance if the method completed successfully; otherwise
// `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeBuffer(length:options:)

func (o MTLDeviceObject) NewBufferWithLengthOptions(length uint, options MTLResourceOptions) MTLBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newBufferWithLength:options:"), length, options)
	return MTLBufferObjectFromID(rv)
	}

// Creates a queue you use to submit rendering and computation commands to a
// GPU.
//
// # Return Value
// 
// A new [MTLCommandQueue] instance if the method completed successfully;
// otherwise `nil`.
//
// # Discussion
// 
// A command queue can only submit commands to the GPU device instance that
// created it.
// 
// This method is the equivalent of passing `64` to the
// [NewCommandQueueWithMaxCommandBufferCount] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandQueue()

func (o MTLDeviceObject) NewCommandQueue() MTLCommandQueue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newCommandQueue"))
	return MTLCommandQueueObjectFromID(rv)
	}

// Creates a queue you use to submit rendering and computation commands to a
// GPU that has a fixed number of uncompleted command buffers.
//
// maxCommandBufferCount: An integer that sets the maximum number of uncompleted command buffers the
// queue can allow.
//
// # Return Value
// 
// A new [MTLCommandQueue] instance if the method completed successfully;
// otherwise `nil`.
//
// # Discussion
// 
// A Command queue can only submit commands to the GPU device instance that
// created it.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCommandQueue(maxCommandBufferCount:)

func (o MTLDeviceObject) NewCommandQueueWithMaxCommandBufferCount(maxCommandBufferCount uint) MTLCommandQueue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newCommandQueueWithMaxCommandBufferCount:"), maxCommandBufferCount)
	return MTLCommandQueueObjectFromID(rv)
	}

// Asynchronously creates a compute pipeline state and reflection information.
//
// descriptor: An [MTLComputePipelineDescriptor] instance.
//
// options: An [MTLPipelineOption] instance that represents the reflection information
// you want the method to generate.
// //
// [MTLPipelineOption]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
//
// completionHandler: A Swift closure or an Objective-C block the method calls when it finishes
// creating the compute pipeline state.
//
// # Discussion
// 
// Use the compute pipeline state to configure a compute pass by calling the
// [SetComputePipelineState] method of an [MTLComputeCommandEncoder] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(descriptor:options:completionHandler:)

func (o MTLDeviceObject) NewComputePipelineStateWithDescriptorOptionsCompletionHandler(descriptor IMTLComputePipelineDescriptor, options MTLPipelineOption, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newComputePipelineStateWithDescriptor:options:completionHandler:"), descriptor, options, completionHandler)
	}

// Synchronously creates a compute pipeline state and reflection information.
//
// descriptor: An [MTLComputePipelineDescriptor] instance.
//
// options: An [MTLPipelineOption] instance that represents the reflection information
// you want the method to generate.
// //
// [MTLPipelineOption]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
//
// reflection: In Swift, an optional pointer to an
// [MTLAutoreleasedComputePipelineReflection] optional. In Objective-C, a
// pointer to an [MTLAutoreleasedComputePipelineReflection] instance.
// 
// Pass `nil` in either language when you don’t need reflection data.
// Otherwise on return, if the method completes successfully, it assigns an
// [MTLComputePipelineReflection] instance to the pointee, which contains the
// details about the function arguments.
//
// # Discussion
// 
// Use the compute pipeline state to configure a compute pass by calling the
// [SetComputePipelineState] method of an [MTLComputeCommandEncoder] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(descriptor:options:reflection:)

func (o MTLDeviceObject) NewComputePipelineStateWithDescriptorOptionsReflectionError(descriptor IMTLComputePipelineDescriptor, options MTLPipelineOption, reflection *MTLAutoreleasedComputePipelineReflection) (MTLComputePipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithDescriptor:options:reflection:error:"), descriptor, options, reflection)
	if err != nil {
		return nil, err
	}
	return MTLComputePipelineStateObjectFromID(rv), nil
	}

// Asynchronously creates a compute pipeline state with a function instance.
//
// computeFunction: An [MTLFunction] instance.
//
// completionHandler: A Swift closure or an Objective-C block the method calls when it finishes
// creating the compute pipeline state.
//
// # Discussion
// 
// Use the compute pipeline state to configure a compute pass by calling the
// [SetComputePipelineState] method of an [MTLComputeCommandEncoder] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(function:completionHandler:)

func (o MTLDeviceObject) NewComputePipelineStateWithFunctionCompletionHandler(computeFunction MTLFunction, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newComputePipelineStateWithFunction:completionHandler:"), computeFunction, completionHandler)
	}

// Synchronously creates a compute pipeline state with a function instance.
//
// computeFunction: An [MTLFunction] instance.
//
// # Discussion
// 
// Use the compute pipeline state to configure a compute pass by calling the
// [SetComputePipelineState] method of an [MTLComputeCommandEncoder] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(function:)

func (o MTLDeviceObject) NewComputePipelineStateWithFunctionError(computeFunction MTLFunction) (MTLComputePipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithFunction:error:"), computeFunction)
	if err != nil {
		return nil, err
	}
	return MTLComputePipelineStateObjectFromID(rv), nil
	}

// Asynchronously creates a compute pipeline state and reflection with a
// function instance.
//
// computeFunction: An [MTLFunction] instance.
//
// options: An [MTLPipelineOption] instance that represents the reflection information
// you want the method to generate.
// //
// [MTLPipelineOption]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
//
// completionHandler: A Swift closure or an Objective-C block the method calls when it finishes
// creating the compute pipeline state.
//
// # Discussion
// 
// Use the compute pipeline state to configure a compute pass by calling the
// [SetComputePipelineState] method of an [MTLComputeCommandEncoder] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(function:options:completionHandler:)

func (o MTLDeviceObject) NewComputePipelineStateWithFunctionOptionsCompletionHandler(computeFunction MTLFunction, options MTLPipelineOption, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newComputePipelineStateWithFunction:options:completionHandler:"), computeFunction, options, completionHandler)
	}

// Synchronously creates a compute pipeline state and reflection with a
// function instance.
//
// computeFunction: An [MTLFunction] instance.
//
// options: An [MTLPipelineOption] instance that represents the reflection information
// you want the method to generate.
// //
// [MTLPipelineOption]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
//
// reflection: In Swift, an optional pointer to an
// [MTLAutoreleasedComputePipelineReflection] optional. In Objective-C, a
// pointer to an [MTLAutoreleasedComputePipelineReflection] instance.
//
// # Discussion
// 
// Use the compute pipeline state to configure a compute pass by calling the
// [SetComputePipelineState] method of an [MTLComputeCommandEncoder] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeComputePipelineState(function:options:reflection:)

func (o MTLDeviceObject) NewComputePipelineStateWithFunctionOptionsReflectionError(computeFunction MTLFunction, options MTLPipelineOption, reflection *MTLAutoreleasedComputePipelineReflection) (MTLComputePipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithFunction:options:reflection:error:"), computeFunction, options, reflection)
	if err != nil {
		return nil, err
	}
	return MTLComputePipelineStateObjectFromID(rv), nil
	}

// Creates a counter sample buffer.
//
// descriptor: An [MTLCounterSampleBufferDescriptor] instance.
//
// # Return Value
// 
// A new [MTLCounterSampleBuffer] instance if the method completes
// successfully; otherwise Swift throws an error and Objective-C returns
// `nil`.
//
// # Discussion
// 
// The method may produce an error if the GPU driver has exhausted its
// underlying resources for counter sample buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeCounterSampleBuffer(descriptor:)

func (o MTLDeviceObject) NewCounterSampleBufferWithDescriptorError(descriptor IMTLCounterSampleBufferDescriptor) (MTLCounterSampleBuffer, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newCounterSampleBufferWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLCounterSampleBufferObjectFromID(rv), nil
	}

// Creates a Metal library instance that contains the functions from your
// app’s default Metal library.
//
// # Return Value
// 
// A new [MTLLibrary] instance if the method completes successfully; otherwise
// `nil`.
//
// # Discussion
// 
// Xcode compiles all the Metal source files (ending in
// `XCUIElementTypeMetal`) in an Xcode project into a single default library.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeDefaultLibrary()

func (o MTLDeviceObject) NewDefaultLibrary() MTLLibrary {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newDefaultLibrary"))
	return MTLLibraryObjectFromID(rv)
	}

// Creates a Metal library instance that contains the functions in a
// bundle’s default Metal library.
//
// bundle: A bundle instance.
//
// # Return Value
// 
// A new [MTLLibrary] instance if the method completes successfully; otherwise
// Swift throws an error and Objective-C returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeDefaultLibrary(bundle:)

func (o MTLDeviceObject) NewDefaultLibraryWithBundleError(bundle foundation.NSBundle) (MTLLibrary, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newDefaultLibraryWithBundle:error:"), bundle)
	if err != nil {
		return nil, err
	}
	return MTLLibraryObjectFromID(rv), nil
	}

// Creates a depth-stencil state instance.
//
// descriptor: An [MTLDepthStencilDescriptor] instance.
//
// # Return Value
// 
// A new [MTLDepthStencilState] instance if the method completed successfully;
// otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeDepthStencilState(descriptor:)

func (o MTLDeviceObject) NewDepthStencilStateWithDescriptor(descriptor IMTLDepthStencilDescriptor) MTLDepthStencilState {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newDepthStencilStateWithDescriptor:"), descriptor)
	return MTLDepthStencilStateObjectFromID(rv)
	}

// Creates a Metal dynamic library instance from a Metal library instance.
//
// library: An [MTLLibrary] instance.
//
// # Return Value
// 
// A new [MTLDynamicLibrary] instance if the method completes successfully;
// otherwise Swift throws an error and Objective-C returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeDynamicLibrary(library:)

func (o MTLDeviceObject) NewDynamicLibraryError(library MTLLibrary) (MTLDynamicLibrary, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newDynamicLibrary:error:"), library)
	if err != nil {
		return nil, err
	}
	return MTLDynamicLibraryObjectFromID(rv), nil
	}

// Creates a Metal dynamic library instance that contains the functions in the
// Metal library file at a URL.
//
// url: A URL to a Metal library file (ending in `XCUIElementTypeMetallib`).
//
// # Return Value
// 
// A new [MTLDynamicLibrary] instance if the method completes successfully;
// otherwise Swift throws an error and Objective-C returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeDynamicLibrary(url:)

func (o MTLDeviceObject) NewDynamicLibraryWithURLError(url foundation.INSURL) (MTLDynamicLibrary, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newDynamicLibraryWithURL:error:"), url)
	if err != nil {
		return nil, err
	}
	return MTLDynamicLibraryObjectFromID(rv), nil
	}

// Creates a new event instance that you can use to synchronize commands and
// resources within the same GPU device.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeEvent()

func (o MTLDeviceObject) NewEvent() MTLEvent {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newEvent"))
	return MTLEventObjectFromID(rv)
	}

// Creates a new memory fence instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeFence()

func (o MTLDeviceObject) NewFence() MTLFence {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newFence"))
	return MTLFenceObjectFromID(rv)
	}

// Creates a new GPU heap instance.
//
// descriptor: An [MTLHeapDescriptor] instance.
//
// # Return Value
// 
// A new [MTLHeap] instance if the method completed successfully; otherwise
// nil.
//
// # Discussion
// 
// For more information about using heaps, see [Memory heaps].
//
// [Memory heaps]: https://developer.apple.com/documentation/Metal/memory-heaps
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeHeap(descriptor:)

func (o MTLDeviceObject) NewHeapWithDescriptor(descriptor IMTLHeapDescriptor) MTLHeap {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newHeapWithDescriptor:"), descriptor)
	return MTLHeapObjectFromID(rv)
	}

// Creates an input/output command queue you use to submit commands that load
// assets from the file system into GPU resources or system memory.
//
// descriptor: A descriptor instance that configures the command queue.
//
// # Return Value
// 
// A new [MTLIOCommandQueue] instance if the method completes successfully;
// otherwise Swift throws an error and Objective-C returns `nil`.
//
// # Discussion
// 
// For information about using input/output command queues and file handles,
// see [Resource loading].
//
// [Resource loading]: https://developer.apple.com/documentation/Metal/resource-loading
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeIOCommandQueue(descriptor:)

func (o MTLDeviceObject) NewIOCommandQueueWithDescriptorError(descriptor IMTLIOCommandQueueDescriptor) (MTLIOCommandQueue, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newIOCommandQueueWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLIOCommandQueueObjectFromID(rv), nil
	}

// Creates an input/output file handle instance that represents a compressed
// file at a URL.
//
// url: A location URL to a compressed file in the file system.
//
// compressionMethod: The file’s compression format.
//
// # Return Value
// 
// A new [MTLIOFileHandle] instance if the method completes successfully;
// otherwise Swift throws an error and Objective-C returns `nil`.
//
// # Discussion
// 
// For information about using input/output command queues and file handles,
// see [Resource loading].
//
// [Resource loading]: https://developer.apple.com/documentation/Metal/resource-loading
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeIOFileHandle(url:compressionMethod:)

func (o MTLDeviceObject) NewIOFileHandleWithURLCompressionMethodError(url foundation.INSURL, compressionMethod MTLIOCompressionMethod) (MTLIOFileHandle, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newIOFileHandleWithURL:compressionMethod:error:"), url, compressionMethod)
	if err != nil {
		return nil, err
	}
	return MTLIOFileHandleObjectFromID(rv), nil
	}

// Creates an input/output file handle instance that represents a file at a
// URL.
//
// url: The URL to a resource file in the file system.
//
// # Return Value
// 
// A new [MTLIOFileHandle] instance if the method completes successfully;
// otherwise Swift throws an error and Objective-C returns `nil`.
//
// # Discussion
// 
// For information about using input/output command queues and file handles,
// see [Resource loading].
//
// [Resource loading]: https://developer.apple.com/documentation/Metal/resource-loading
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeIOFileHandle(url:)

func (o MTLDeviceObject) NewIOFileHandleWithURLError(url foundation.INSURL) (MTLIOFileHandle, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newIOFileHandleWithURL:error:"), url)
	if err != nil {
		return nil, err
	}
	return MTLIOFileHandleObjectFromID(rv), nil
	}

// Creates an indirect command buffer instance.
//
// descriptor: An [MTLIndirectCommandBufferDescriptor] instance.
//
// maxCount: The largest number of commands you can store in the buffer.
//
// options: An [MTLResourceOptions] instance.
// //
// [MTLResourceOptions]: https://developer.apple.com/documentation/Metal/MTLResourceOptions
//
// # Return Value
// 
// A new [MTLIndirectCommandBuffer] instance if the method completed
// successfully; otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeIndirectCommandBuffer(descriptor:maxCommandCount:options:)

func (o MTLDeviceObject) NewIndirectCommandBufferWithDescriptorMaxCommandCountOptions(descriptor IMTLIndirectCommandBufferDescriptor, maxCount uint, options MTLResourceOptions) MTLIndirectCommandBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newIndirectCommandBufferWithDescriptor:maxCommandCount:options:"), descriptor, maxCount, options)
	return MTLIndirectCommandBufferObjectFromID(rv)
	}

// Creates a Metal library instance from a dispatch-data instance that
// contains the functions in a precompiled Metal library.
//
// data: The data from a precompiled Metal library. For more information, see
// [Building a shader library by precompiling source files].
// //
// [Building a shader library by precompiling source files]: https://developer.apple.com/documentation/Metal/building-a-shader-library-by-precompiling-source-files
//
// # Return Value
// 
// A new [MTLLibrary] instance if the method completes successfully; otherwise
// Swift throws an error and Objective-C returns `nil`.
//
// # Discussion
// 
// In Swift, you can also use the [makeLibrary(data:)] default implementation,
// which has a [DispatchData] parameter.
// 
// Use either method if your application manages its own archiving system for
// libraries — for example, if your app uses a single file that contains
// several libraries.
//
// [DispatchData]: https://developer.apple.com/documentation/Dispatch/DispatchData
// [makeLibrary(data:)]: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(data:)-7khmh
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(data:)

func (o MTLDeviceObject) NewLibraryWithDataError(data dispatch.Data) (MTLLibrary, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newLibraryWithData:error:"), uintptr(data.Handle()))
	if err != nil {
		return nil, err
	}
	return MTLLibraryObjectFromID(rv), nil
	}

// Asynchronously creates a Metal library instance by compiling the functions
// in a source string.
//
// source: A string that contains source code for one or more Metal functions. For
// information about writing source in Metal Shading Language (MSL), see the
// [Metal Shading Language Specification].
// //
// [Metal Shading Language Specification]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// options: An [MTLCompileOptions] instance that affects the compilation of the source
// code in the string.
//
// completionHandler: A Swift closure or an Objective-C block the method calls when the library
// finishes loading.
//
// # Discussion
// 
// Because there’s no search path to find other functions, the source may
// only import the Metal default library.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(source:options:completionHandler:)

func (o MTLDeviceObject) NewLibraryWithSourceOptionsCompletionHandler(source string, options IMTLCompileOptions, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newLibraryWithSource:options:completionHandler:"), objc.String(source), options, completionHandler)
	}

// Synchronously creates a Metal library instance by compiling the functions
// in a source string.
//
// source: A string that contains source code for one or more Metal functions. For
// information about writing source in Metal Shading Language (MSL), see the
// [Metal Shading Language Specification].
// //
// [Metal Shading Language Specification]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// options: An [MTLCompileOptions] instance that affects the compilation of the source
// code in the string.
//
// # Return Value
// 
// A new [MTLLibrary] instance if the method completes successfully; otherwise
// Swift throws an error and Objective-C returns `nil`.
//
// # Discussion
// 
// Because there’s no search path to find other functions, the source may
// only import the Metal default library.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(source:options:)

func (o MTLDeviceObject) NewLibraryWithSourceOptionsError(source string, options IMTLCompileOptions) (MTLLibrary, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newLibraryWithSource:options:error:"), objc.String(source), options)
	if err != nil {
		return nil, err
	}
	return MTLLibraryObjectFromID(rv), nil
	}

// Asynchronously creates a Metal library from the function stitching graphs
// in a descriptor.
//
// descriptor: An [MTLStitchedLibraryDescriptor] instance.
//
// completionHandler: A Swift closure or Objective-C block the method calls when the library
// finishes loading.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(stitchedDescriptor:completionHandler:)

func (o MTLDeviceObject) NewLibraryWithStitchedDescriptorCompletionHandler(descriptor IMTLStitchedLibraryDescriptor, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newLibraryWithStitchedDescriptor:completionHandler:"), descriptor, completionHandler)
	}

// Synchronously creates a Metal library from the function stitching graphs in
// a descriptor.
//
// descriptor: An [MTLStitchedLibraryDescriptor] instance.
//
// # Return Value
// 
// A new [MTLLibrary] instance if the method completes successfully; otherwise
// Swift throws an error and Objective-C returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(stitchedDescriptor:)

func (o MTLDeviceObject) NewLibraryWithStitchedDescriptorError(descriptor IMTLStitchedLibraryDescriptor) (MTLLibrary, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newLibraryWithStitchedDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLLibraryObjectFromID(rv), nil
	}

// Creates a Metal library instance that contains the functions in the Metal
// library file at a URL.
//
// url: A URL to a Metal library file (ending in `XCUIElementTypeMetallib`).
//
// # Return Value
// 
// A new [MTLLibrary] instance if the method completes successfully; otherwise
// Swift throws an error and Objective-C returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeLibrary(URL:)

func (o MTLDeviceObject) NewLibraryWithURLError(url foundation.INSURL) (MTLLibrary, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newLibraryWithURL:error:"), url)
	if err != nil {
		return nil, err
	}
	return MTLLibraryObjectFromID(rv), nil
	}

// Creates a rasterization rate map instance.
//
// descriptor: An [MTLRasterizationRateMapDescriptor] instance.
//
// # Return Value
// 
// A new [MTLRasterizationRateMapDescriptor] instance if the method completes
// successfully; otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRasterizationRateMap(descriptor:)

func (o MTLDeviceObject) NewRasterizationRateMapWithDescriptor(descriptor IMTLRasterizationRateMapDescriptor) MTLRasterizationRateMap {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newRasterizationRateMapWithDescriptor:"), descriptor)
	return MTLRasterizationRateMapObjectFromID(rv)
	}

// Asynchronously creates a render pipeline state.
//
// descriptor: An [MTLRenderPipelineDescriptor] instance.
//
// completionHandler: A Swift closure or an Objective-C block the method calls when it finishes
// creating the render pipeline state.
//
// # Discussion
// 
// Use the graphics-rendering pipeline state to configure a render pass by
// calling the [SetRenderPipelineState] method of an [MTLRenderCommandEncoder]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(descriptor:completionHandler:)

func (o MTLDeviceObject) NewRenderPipelineStateWithDescriptorCompletionHandler(descriptor IMTLRenderPipelineDescriptor, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newRenderPipelineStateWithDescriptor:completionHandler:"), descriptor, completionHandler)
	}

// Synchronously creates a render pipeline state.
//
// descriptor: An [MTLRenderPipelineDescriptor] instance.
//
// # Return Value
// 
// A new [MTLRenderPipelineState] instance if the method completes
// successfully; otherwise Swift throws an error and Objective-C returns
// `nil`.
//
// # Discussion
// 
// Use the graphics-rendering pipeline state to configure a render pass by
// calling the [SetRenderPipelineState] method of an [MTLRenderCommandEncoder]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(descriptor:)

func (o MTLDeviceObject) NewRenderPipelineStateWithDescriptorError(descriptor IMTLRenderPipelineDescriptor) (MTLRenderPipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
	}

// Asynchronously creates a render pipeline state and reflection information.
//
// descriptor: An [MTLRenderPipelineDescriptor] instance.
//
// options: An [MTLPipelineOption] instance that represents the reflection information
// you want the method to generate.
// //
// [MTLPipelineOption]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
//
// completionHandler: A Swift closure or an Objective-C block the method calls when it finishes
// creating the render pipeline state.
//
// # Discussion
// 
// Use the graphics-rendering pipeline state to configure a render pass by
// calling the [SetRenderPipelineState] method of an [MTLRenderCommandEncoder]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(descriptor:options:completionHandler:)-5gdww

func (o MTLDeviceObject) NewRenderPipelineStateWithDescriptorOptionsCompletionHandler(descriptor IMTLRenderPipelineDescriptor, options MTLPipelineOption, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newRenderPipelineStateWithDescriptor:options:completionHandler:"), descriptor, options, completionHandler)
	}

// Synchronously creates a render pipeline state and reflection information.
//
// descriptor: An [MTLRenderPipelineDescriptor] instance.
//
// options: An [MTLPipelineOption] instance that represents the reflection information
// you want the method to generate.
// //
// [MTLPipelineOption]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
//
// reflection: In Swift, an optional pointer to an
// [MTLAutoreleasedRenderPipelineReflection] optional. In Objective-C, a
// pointer to an [MTLAutoreleasedRenderPipelineReflection] instance.
// 
// Pass `nil` in either language when you don’t need reflection data.
// Otherwise on return, if the method completes successfully, it assigns an
// [MTLRenderPipelineReflection] instance to the pointee, which contains the
// details about the function arguments.
//
// # Return Value
// 
// A new [MTLRenderPipelineState] instance if the method completes
// successfully; otherwise Swift throws an error and Objective-C returns
// `nil`.
//
// # Discussion
// 
// Use the graphics-rendering pipeline state to configure a render pass by
// calling the [SetRenderPipelineState] method of an [MTLRenderCommandEncoder]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(descriptor:options:reflection:)

func (o MTLDeviceObject) NewRenderPipelineStateWithDescriptorOptionsReflectionError(descriptor IMTLRenderPipelineDescriptor, options MTLPipelineOption, reflection *MTLAutoreleasedRenderPipelineReflection) (MTLRenderPipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithDescriptor:options:reflection:error:"), descriptor, options, reflection)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
	}

// Asynchronously creates a mesh render pipeline state and reflection
// information.
//
// descriptor: An [MTLMeshRenderPipelineDescriptor] instance.
//
// options: An [MTLPipelineOption] instance that represents the reflection information
// you want the method to generate.
// //
// [MTLPipelineOption]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
//
// completionHandler: A Swift closure or an Objective-C block the method calls when it finishes
// creating the render pipeline state.
//
// # Discussion
// 
// Use the graphics-rendering pipeline state to configure a render pass by
// calling the [SetRenderPipelineState] method of an [MTLRenderCommandEncoder]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(descriptor:options:completionHandler:)-1wvya

func (o MTLDeviceObject) NewRenderPipelineStateWithMeshDescriptorOptionsCompletionHandler(descriptor IMTLMeshRenderPipelineDescriptor, options MTLPipelineOption, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newRenderPipelineStateWithMeshDescriptor:options:completionHandler:"), descriptor, options, completionHandler)
	}

// Synchronously creates a mesh render pipeline state and reflection
// information.
//
// descriptor: An [MTLMeshRenderPipelineDescriptor] instance.
//
// options: An [MTLPipelineOption] instance that represents the reflection information
// you want the method to generate.
// //
// [MTLPipelineOption]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
//
// reflection: In Swift, an optional pointer to an
// [MTLAutoreleasedRenderPipelineReflection] optional. In Objective-C, a
// pointer to an [MTLAutoreleasedRenderPipelineReflection] instance.
// 
// Pass `nil` in either language when you don’t need reflection data.
// Otherwise on return, if the method completes successfully, it assigns an
// [MTLRenderPipelineReflection] instance to the pointee, which contains the
// details about the function arguments.
//
// error: On return, if an error occurs, a pointer to an error information instance;
// otherwise, `nil`.
//
// # Return Value
// 
// A new [MTLRenderPipelineState] instance if the method completes
// successfully; otherwise, `nil`.
//
// # Discussion
// 
// Use the graphics-rendering pipeline state to configure a render pass by
// calling the [SetRenderPipelineState] method of an [MTLRenderCommandEncoder]
// instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/newRenderPipelineStateWithMeshDescriptor:options:reflection:error:

func (o MTLDeviceObject) NewRenderPipelineStateWithMeshDescriptorOptionsReflectionError(descriptor IMTLMeshRenderPipelineDescriptor, options MTLPipelineOption, reflection *MTLAutoreleasedRenderPipelineReflection) (MTLRenderPipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithMeshDescriptor:options:reflection:error:"), descriptor, options, reflection)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
	}

// Asynchronously creates a tile shader’s render pipeline state and
// reflection information.
//
// descriptor: An [MTLTileRenderPipelineDescriptor] instance.
//
// options: An [MTLPipelineOption] instance that represents the reflection information
// you want the method to generate.
// //
// [MTLPipelineOption]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
//
// completionHandler: A Swift closure or an Objective-C block the method calls when it finishes
// creating the render pipeline state.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(tileDescriptor:options:completionHandler:)

func (o MTLDeviceObject) NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler(descriptor IMTLTileRenderPipelineDescriptor, options MTLPipelineOption, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newRenderPipelineStateWithTileDescriptor:options:completionHandler:"), descriptor, options, completionHandler)
	}

// Synchronously creates a tile shader’s render pipeline state and
// reflection information.
//
// descriptor: An [MTLTileRenderPipelineDescriptor] instance.
//
// options: An [MTLPipelineOption] instance that represents the reflection information
// you want the method to generate.
// //
// [MTLPipelineOption]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
//
// reflection: In Swift, an optional pointer to an
// [MTLAutoreleasedRenderPipelineReflection] optional. In Objective-C, a
// pointer to an [MTLAutoreleasedRenderPipelineReflection] instance.
// 
// Pass `nil` in either language when you don’t need reflection data.
// Otherwise on return, if the method completes successfully, it assigns an
// [MTLRenderPipelineReflection] instance to the pointee, which contains the
// details about the function arguments.
//
// # Return Value
// 
// A new [MTLRenderPipelineState] instance if the method completes
// successfully; otherwise Swift throws an error and Objective-C returns
// `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeRenderPipelineState(tileDescriptor:options:reflection:)

func (o MTLDeviceObject) NewRenderPipelineStateWithTileDescriptorOptionsReflectionError(descriptor IMTLTileRenderPipelineDescriptor, options MTLPipelineOption, reflection *MTLAutoreleasedRenderPipelineReflection) (MTLRenderPipelineState, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithTileDescriptor:options:reflection:error:"), descriptor, options, reflection)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
	}

// Creates a residency set, which can move resources in and out of memory
// residency.
//
// desc: A descriptor instance that configures the residency set the method creates.
//
// # Return Value
// 
// A new [MTLResidencySet] instance if the method completes successfully;
// otherwise Swift throws an error and Objective-C returns `nil`.
//
// # Discussion
// 
// Create an [MTLResidencySet] by creating and configuring an
// [MTLResidencySetDescriptor] instance and pass it to this method.
// 
// See [Simplifying GPU resource management with residency sets] for more
// information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeResidencySet(descriptor:)

func (o MTLDeviceObject) NewResidencySetWithDescriptorError(desc IMTLResidencySetDescriptor) (MTLResidencySet, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newResidencySetWithDescriptor:error:"), desc)
	if err != nil {
		return nil, err
	}
	return MTLResidencySetObjectFromID(rv), nil
	}

// Creates a sampler state instance.
//
// descriptor: An [MTLSamplerDescriptor] instance.
//
// # Return Value
// 
// A new [MTLSamplerState] instance if the method completed successfully;
// otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeSamplerState(descriptor:)

func (o MTLDeviceObject) NewSamplerStateWithDescriptor(descriptor IMTLSamplerDescriptor) MTLSamplerState {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newSamplerStateWithDescriptor:"), descriptor)
	return MTLSamplerStateObjectFromID(rv)
	}

// Creates a new shared event instance that you can use to synchronize
// commands and resources across different GPU devices.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeSharedEvent()

func (o MTLDeviceObject) NewSharedEvent() MTLSharedEvent {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newSharedEvent"))
	return MTLSharedEventObjectFromID(rv)
	}

// Recreates a shared event from a handle.
//
// sharedEventHandle: An [MTLSharedEventHandle] instance from another GPU device or process.
//
// # Return Value
// 
// A new [MTLSharedEvent] instance if the method completed successfully;
// otherwise nil.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeSharedEvent(handle:)

func (o MTLDeviceObject) NewSharedEventWithHandle(sharedEventHandle IMTLSharedEventHandle) MTLSharedEvent {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newSharedEventWithHandle:"), sharedEventHandle)
	return MTLSharedEventObjectFromID(rv)
	}

// Creates a texture that you can share across process boundaries.
//
// descriptor: An [MTLTextureDescriptor] instance.
//
// # Return Value
// 
// A new [MTLTexture] instance if the method completed successfully; otherwise
// `nil`.
//
// # Discussion
// 
// You can create a shared texture but only with [ResourceStorageModePrivate].
// You can share the texture with another process by:
// 
// - Creating a texture handle (see [NewSharedTextureHandle]) - Passing the
// texture handle to the other process - Creating a texture in the other
// process by calling the [NewSharedTextureWithHandle]method
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeSharedTexture(descriptor:)

func (o MTLDeviceObject) NewSharedTextureWithDescriptor(descriptor IMTLTextureDescriptor) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newSharedTextureWithDescriptor:"), descriptor)
	return MTLTextureObjectFromID(rv)
	}

// Creates a texture that references a shared texture.
//
// sharedHandle: An [MTLSharedTextureHandle] instance, typically from another process using
// the same GPU device.
//
// # Return Value
// 
// A new [MTLTexture] instance if the method completed successfully; otherwise
// `nil`.
//
// # Discussion
// 
// Call this method from the same [MTLDevice] instance that created the shared
// texture instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeSharedTexture(handle:)

func (o MTLDeviceObject) NewSharedTextureWithHandle(sharedHandle IMTLSharedTextureHandle) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newSharedTextureWithHandle:"), sharedHandle)
	return MTLTextureObjectFromID(rv)
	}

// Creates a new texture instance.
//
// descriptor: An [MTLTextureDescriptor] instance.
//
// # Return Value
// 
// A new [MTLTexture] instance if the method completed successfully; otherwise
// `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeTexture(descriptor:)

func (o MTLDeviceObject) NewTextureWithDescriptor(descriptor IMTLTextureDescriptor) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newTextureWithDescriptor:"), descriptor)
	return MTLTextureObjectFromID(rv)
	}

// Creates a texture instance that uses I/O surface to store its underlying
// data.
//
// descriptor: An [MTLTextureDescriptor] instance.
//
// iosurface: An [IOSurfaceRef] instance.
//
// plane: A plane within i`osurface` the method sets as the texture’s underlying
// data.
//
// # Return Value
// 
// A new [MTLTexture] instance if the method completed successfully; otherwise
// `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/makeTexture(descriptor:iosurface:plane:)

func (o MTLDeviceObject) NewTextureWithDescriptorIosurfacePlane(descriptor IMTLTextureDescriptor, iosurface iosurface.IOSurfaceRef, plane uint) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newTextureWithDescriptor:iosurface:plane:"), descriptor, iosurface, plane)
	return MTLTextureObjectFromID(rv)
	}

// Captures and returns a CPU timestamp and a GPU timestamp from the same
// moment in time.
//
// cpuTimestamp: A pointer the method uses to save a timestamp from the CPU.
//
// gpuTimestamp: A pointer the method uses to save a timestamp from the GPU the device
// instance represents.
//
// # Discussion
// 
// For an example of how and when to use corresponding timestamps from the CPU
// and GPU, see [Converting GPU timestamps into CPU time].
//
// [Converting GPU timestamps into CPU time]: https://developer.apple.com/documentation/Metal/converting-gpu-timestamps-into-cpu-time
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/sampleTimestamps:gpuTimestamp:

func (o MTLDeviceObject) SampleTimestampsGpuTimestamp(cpuTimestamp *MTLTimestamp, gpuTimestamp *MTLTimestamp) {
	
	objc.Send[struct{}](o.ID, objc.Sel("sampleTimestamps:gpuTimestamp:"), cpuTimestamp, gpuTimestamp)
	}

// Returns the size, in bytes, of a sparse tile the GPU device creates with a
// specific page size.
//
// sparsePageSize: An [MTLSparsePageSize] instance.
// //
// [MTLSparsePageSize]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/sparseTileSizeInBytes(sparsePageSize:)

func (o MTLDeviceObject) SparseTileSizeInBytesForSparsePageSize(sparsePageSize MTLSparsePageSize) uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("sparseTileSizeInBytesForSparsePageSize:"), sparsePageSize)
	return rv
	}

// Returns the dimensions of a sparse tile for a texture.
//
// textureType: An [MTLTextureType] instance.
// //
// [MTLTextureType]: https://developer.apple.com/documentation/Metal/MTLTextureType
//
// pixelFormat: An [MTLPixelFormat] instance.
// //
// [MTLPixelFormat]: https://developer.apple.com/documentation/Metal/MTLPixelFormat
//
// sampleCount: The number of samples for each pixel.
//
// # Return Value
// 
// A new [MTLSize] instance.
//
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// # Discussion
// 
// The size of a sparse tile, in bytes, is the same for all sparse textures on
// a GPU device object. Because the size of pixels may vary, the actual
// dimensions of a sparse tile vary based on the texture and the pixel format.
// Use this method to get the dimensions of the tile for a particular format.
// Use these dimensions when converting regions from pixel-based units to
// sparse tile units and vice versa.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/sparseTileSize(with:pixelFormat:sampleCount:)

func (o MTLDeviceObject) SparseTileSizeWithTextureTypePixelFormatSampleCount(textureType MTLTextureType, pixelFormat MTLPixelFormat, sampleCount uint) MTLSize {
	
	rv := objc.Send[MTLSize](o.ID, objc.Sel("sparseTileSizeWithTextureType:pixelFormat:sampleCount:"), textureType, pixelFormat, sampleCount)
	return rv
	}

// Returns the dimensions of a sparse tile for a texture that has a specific
// sparse page size.
//
// textureType: An [MTLTextureType] instance.
// //
// [MTLTextureType]: https://developer.apple.com/documentation/Metal/MTLTextureType
//
// pixelFormat: An [MTLPixelFormat] instance.
// //
// [MTLPixelFormat]: https://developer.apple.com/documentation/Metal/MTLPixelFormat
//
// sampleCount: The number of samples for each pixel.
//
// sparsePageSize: An [MTLSparsePageSize] instance.
// //
// [MTLSparsePageSize]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize
//
// # Return Value
// 
// A new [MTLSize] instance.
//
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/sparseTileSize(textureType:pixelFormat:sampleCount:sparsePageSize:)

func (o MTLDeviceObject) SparseTileSizeWithTextureTypePixelFormatSampleCountSparsePageSize(textureType MTLTextureType, pixelFormat MTLPixelFormat, sampleCount uint, sparsePageSize MTLSparsePageSize) MTLSize {
	
	rv := objc.Send[MTLSize](o.ID, objc.Sel("sparseTileSizeWithTextureType:pixelFormat:sampleCount:sparsePageSize:"), textureType, pixelFormat, sampleCount, sparsePageSize)
	return rv
	}

// Returns a Boolean value that indicates whether you can read GPU counters at
// the specified command boundary.
//
// samplingPoint: The command boundary to test.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsCounterSampling(_:)

func (o MTLDeviceObject) SupportsCounterSampling(samplingPoint MTLCounterSamplingPoint) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("supportsCounterSampling:"), samplingPoint)
	return rv
	}

// Returns a Boolean value that indicates whether the GPU device supports the
// feature set of a specific GPU family.
//
// gpuFamily: An [MTLGPUFamily] instance.
// //
// [MTLGPUFamily]: https://developer.apple.com/documentation/Metal/MTLGPUFamily
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsFamily(_:)

func (o MTLDeviceObject) SupportsFamily(gpuFamily MTLGPUFamily) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("supportsFamily:"), gpuFamily)
	return rv
	}

// Returns a Boolean value that indicates whether the GPU can create a
// rasterization rate map with a specific number of layers.
//
// layerCount: The number of layers for a rasterization rate map.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsRasterizationRateMap(layerCount:)

func (o MTLDeviceObject) SupportsRasterizationRateMapWithLayerCount(layerCount uint) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("supportsRasterizationRateMapWithLayerCount:"), layerCount)
	return rv
	}

// Returns a Boolean value that indicates whether the GPU can sample a texture
// with a specific number of sample points.
//
// sampleCount: The number of points a GPU can sample from a texture.
//
// # Discussion
// 
// The number of points the GPU can sample a texture varies by device:
// 
// [Table data omitted]
// 
// Consider a GPU device’s limitations for sample count by checking
// [MTLTexture]`.`[SampleCount] when configuring these properties:
// 
// - [MTLTextureDescriptor]`.`[SampleCount] -
// [MTLRenderPipelineDescriptor]`.`[RasterSampleCount] -
// [MTLTileRenderPipelineDescriptor]`.`[RasterSampleCount] -
// [MTLMeshRenderPipelineDescriptor]`.`[RasterSampleCount] -
// [MTKView]`.`[sampleCount]
//
// [MTKView]: https://developer.apple.com/documentation/MetalKit/MTKView
// [sampleCount]: https://developer.apple.com/documentation/MetalKit/MTKView/sampleCount
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsTextureSampleCount(_:)

func (o MTLDeviceObject) SupportsTextureSampleCount(sampleCount uint) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("supportsTextureSampleCount:"), sampleCount)
	return rv
	}

// Returns a Boolean value that indicates whether the GPU supports an
// amplification factor.
//
// count: An integer that represents the number of output streams you want the GPU to
// generate from an input stream.
//
// # Discussion
// 
// A vertex amplification factor of `1` has no effect because it effectively
// disables vertex amplification.
// 
// For more information about vertex amplification, see [Improving rendering
// performance with vertex amplification].
//
// [Improving rendering performance with vertex amplification]: https://developer.apple.com/documentation/Metal/improving-rendering-performance-with-vertex-amplification
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsVertexAmplificationCount(_:)

func (o MTLDeviceObject) SupportsVertexAmplificationCount(count uint) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("supportsVertexAmplificationCount:"), count)
	return rv
	}

// Converts a list of sparse pixel regions to tile regions.
//
// pixelRegions: A pointer to a C array of pixel [MTLRegion] instances.
// //
// [MTLRegion]: https://developer.apple.com/documentation/Metal/MTLRegion
//
// tileRegions: A pointer to a C array of tile [MTLRegion] instances.
// //
// [MTLRegion]: https://developer.apple.com/documentation/Metal/MTLRegion
//
// tileSize: An [MTLSize] instance that represents a sparse tile’s size, in pixels.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// mode: An [MTLSparseTextureRegionAlignmentMode] instance.
// //
// [MTLSparseTextureRegionAlignmentMode]: https://developer.apple.com/documentation/Metal/MTLSparseTextureRegionAlignmentMode
//
// numRegions: The number of regions you want the method to convert.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/convertSparsePixelRegions(_:toTileRegions:withTileSize:alignmentMode:numRegions:)

func (o MTLDeviceObject) ConvertSparsePixelRegionsToTileRegionsWithTileSizeAlignmentModeNumRegions(pixelRegions []MTLRegion, tileRegions []MTLRegion, tileSize MTLSize, mode MTLSparseTextureRegionAlignmentMode, numRegions uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("convertSparsePixelRegions:toTileRegions:withTileSize:alignmentMode:numRegions:"), pixelRegions, tileRegions, tileSize, mode, numRegions)
	}

// Converts a list of sparse tile regions to pixel regions.
//
// tileRegions: A pointer to a C array of tile [MTLRegion] instances.
// //
// [MTLRegion]: https://developer.apple.com/documentation/Metal/MTLRegion
//
// pixelRegions: A pointer to a C array of pixel [MTLRegion] instances.
// //
// [MTLRegion]: https://developer.apple.com/documentation/Metal/MTLRegion
//
// tileSize: An [MTLSize] instance that represents a sparse tile’s size, in pixels.
// //
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
//
// numRegions: The number of regions you want the method to convert.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/convertSparseTileRegions(_:toPixelRegions:withTileSize:numRegions:)

func (o MTLDeviceObject) ConvertSparseTileRegionsToPixelRegionsWithTileSizeNumRegions(tileRegions []MTLRegion, pixelRegions []MTLRegion, tileSize MTLSize, numRegions uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("convertSparseTileRegions:toPixelRegions:withTileSize:numRegions:"), tileRegions, pixelRegions, tileSize, numRegions)
	}

func (o MTLDeviceObject) SetShouldMaximizeConcurrentCompilation(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setShouldMaximizeConcurrentCompilation:"), value)
}

// The architectural details of the GPU device.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/architecture
func (o MTLDeviceObject) Architecture() IMTLArchitecture {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("architecture"))
	return MTLArchitectureFromID(rv)
}

// Returns the GPU device’s support tier for argument buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/argumentBuffersSupport
func (o MTLDeviceObject) ArgumentBuffersSupport() MTLArgumentBuffersTier {
	rv := objc.Send[MTLArgumentBuffersTier](o.ID, objc.Sel("argumentBuffersSupport"))
	return MTLArgumentBuffersTier(rv)
}

// The counter sets supported by the device object.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/counterSets
func (o MTLDeviceObject) CounterSets() []objectivec.IObject {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("counterSets"))
	result := make([]objectivec.IObject, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = objectivec.Object{ID: id}
	}
	return result
}

// The total amount of memory, in bytes, the GPU device is using for all of
// its resources.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/currentAllocatedSize
func (o MTLDeviceObject) CurrentAllocatedSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("currentAllocatedSize"))
	return uint(rv)
}

// A Boolean value that indicates whether a device supports a packed
// depth-and-stencil pixel format.
//
// # Discussion
// 
// If the value is [true], the device supports the
// [PixelFormatDepth24Unorm_Stencil8] pixel format.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/isDepth24Stencil8PixelFormatSupported
func (o MTLDeviceObject) Depth24Stencil8PixelFormatSupported() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isDepth24Stencil8PixelFormatSupported"))
	return bool(rv)
}

// A Boolean value that indicates whether the GPU shares all of its memory
// with the CPU.
//
// # Discussion
// 
// A GPU with unified memory ([true]) is typically an integrated GPU. A GPU
// with dedicated memory ([false]) may take additional time to synchronize
// managed resources or copy data into private GPU resources.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/hasUnifiedMemory
func (o MTLDeviceObject) HasUnifiedMemory() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hasUnifiedMemory"))
	return bool(rv)
}

// A Boolean value that indicates whether a GPU device doesn’t have a
// connection to a display.
//
// # Discussion
// 
// The value is [true] when the GPU is , which means it isn’t connected to
// any displays.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/isHeadless
func (o MTLDeviceObject) Headless() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isHeadless"))
	return bool(rv)
}

// The physical location of the GPU relative to the system.
//
// # Discussion
// 
// The value indicates whether the GPU connects to the system through a
// built-in connection, an internal card slot, or an external connection.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/location
func (o MTLDeviceObject) Location() MTLDeviceLocation {
	rv := objc.Send[MTLDeviceLocation](o.ID, objc.Sel("location"))
	return MTLDeviceLocation(rv)
}

// A specific GPU position based on its general location.
//
// # Discussion
// 
// The meaning of the location number depends on a device’s [location]
// property:
// 
// - For [DeviceLocationBuiltIn], the location number is `0` for low-power
// GPUs (see [isLowPower]) and `1` for other GPUs. - For [DeviceLocationSlot],
// the location number represents the slot. - For [DeviceLocationExternal],
// the location number represents the Thunderbolt port.
//
// [isLowPower]: https://developer.apple.com/documentation/Metal/MTLDevice/isLowPower
// [location]: https://developer.apple.com/documentation/Metal/MTLDevice/location
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/locationNumber
func (o MTLDeviceObject) LocationNumber() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("locationNumber"))
	return uint(rv)
}

// A Boolean value that indicates whether the GPU lowers its performance to
// conserve energy.
//
// # Discussion
// 
// Some systems contain multiple GPUs that run with different performance and
// energy characteristics. At runtime, choose a GPU that best matches your
// performance needs while considering the current state of the system. For
// example, your app may choose a lower-power GPU if it doesn’t need the
// best possible performance on a MacBook Pro that’s running on battery
// power. For more information on discovering and selecting GPUs at runtime,
// see [Multi-GPU systems].
// 
// The property is typically [true] for integrated GPUs and [false] for
// discrete GPUs. However, an Apple silicon GPU on a Mac sets the property to
// [false] because it doesn’t need to lower its performance to conserve
// energy.
//
// [Multi-GPU systems]: https://developer.apple.com/documentation/Metal/multi-gpu-systems
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/isLowPower
func (o MTLDeviceObject) LowPower() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isLowPower"))
	return bool(rv)
}

// The maximum number of unique argument buffer samplers per app.
//
// # Discussion
// 
// This limit only applies to samplers that support argument buffers (see
// [SupportArgumentBuffers]). An [MTLSamplerState] instance is only unique if
// the properties of the [MTLSamplerDescriptor] instance that created it are
// unique. For example, two samplers with equal [MinFilter] values but
// different [MagFilter] values are unique.
// 
// See [Improving CPU performance by using argument buffers] for more
// information about argument buffer tiers, limits, and capabilities.
//
// [Improving CPU performance by using argument buffers]: https://developer.apple.com/documentation/Metal/improving-cpu-performance-by-using-argument-buffers
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/maxArgumentBufferSamplerCount
func (o MTLDeviceObject) MaxArgumentBufferSamplerCount() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("maxArgumentBufferSamplerCount"))
	return uint(rv)
}

// The largest amount of memory, in bytes, that a GPU device can allocate to a
// buffer instance.
//
// # Discussion
// 
// The property’s value is at least 256 MB (268,435,456 bytes).
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/maxBufferLength
func (o MTLDeviceObject) MaxBufferLength() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("maxBufferLength"))
	return uint(rv)
}

// The maximum threadgroup memory available to a compute kernel, in bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/maxThreadgroupMemoryLength
func (o MTLDeviceObject) MaxThreadgroupMemoryLength() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("maxThreadgroupMemoryLength"))
	return uint(rv)
}

// The maximum number of threads along each dimension of a threadgroup.
//
// # Discussion
// 
// This property reports the maximum thread group size the device can support
// for a trivial shader. This size isn’t guaranteed for all shaders. For the
// actual thread group size of a specific compute shader, see the
// [MaxTotalThreadsPerThreadgroup] property of your compute pipeline state.
// 
// For more information on the specific threadgroup limits of each GPU family,
// see the Metal feature set tables:
// 
// - [Metal feature set tables (PDF)]
// - [Metal feature set tables (Numbers)]
//
// [Metal feature set tables (Numbers)]: https://developer.apple.com/metal/metal-feature-set-tables.zip
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/maxThreadsPerThreadgroup
func (o MTLDeviceObject) MaxThreadsPerThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](o.ID, objc.Sel("maxThreadsPerThreadgroup"))
	return MTLSize(rv)
}

// The highest theoretical rate, in bytes per second, the system can copy
// between system memory and the GPU’s dedicated memory (VRAM).
//
// # Discussion
// 
// Metal calculates this value from the raw data-clock rate, and the GPU may
// not be able to reach this speed in real-world conditions.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/maxTransferRate
func (o MTLDeviceObject) MaxTransferRate() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("maxTransferRate"))
	return uint64(rv)
}

// The full name of the GPU device.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/name
func (o MTLDeviceObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The total number of GPUs in the peer group, if applicable.
//
// # Discussion
// 
// A peer count value of `0` indicates the GPU isn’t in a peer group.
// Otherwise, the GPU is in a peer group and the value represents the total
// number of GPUs in the group, including this one.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/peerCount
func (o MTLDeviceObject) PeerCount() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("peerCount"))
	return uint32(rv)
}

// The peer group ID the GPU belongs to, if applicable.
//
// # Discussion
// 
// A group ID value of `0` indicates the GPU isn’t in a peer group.
// Otherwise, the GPU is in a peer group and the value is the group’s ID.
// All other GPUs in the same peer group have the same group ID.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/peerGroupID
func (o MTLDeviceObject) PeerGroupID() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("peerGroupID"))
	return uint64(rv)
}

// The unique identifier for a GPU in a peer group.
//
// # Discussion
// 
// If the GPU is part of a peer group (see [peerGroupID] or [peerCount]) the
// peer index is the GPU’s unique value within the group in the range `[0,
// `[peerCount]`)`.
//
// [0, `[peerCount]: https://developer.apple.com/documentation/Metal/MTLDevice/peerCount
// [peerCount]: https://developer.apple.com/documentation/Metal/MTLDevice/peerCount
// [peerGroupID]: https://developer.apple.com/documentation/Metal/MTLDevice/peerGroupID
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/peerIndex
func (o MTLDeviceObject) PeerIndex() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("peerIndex"))
	return uint32(rv)
}

// A Boolean value that indicates whether the GPU supports programmable sample
// positions.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/areProgrammableSamplePositionsSupported
func (o MTLDeviceObject) ProgrammableSamplePositionsSupported() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("areProgrammableSamplePositionsSupported"))
	return bool(rv)
}

// A Boolean value that indicates whether the GPU supports raster order
// groups.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/areRasterOrderGroupsSupported
func (o MTLDeviceObject) RasterOrderGroupsSupported() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("areRasterOrderGroupsSupported"))
	return bool(rv)
}

// The GPU device’s texture support tier.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/readWriteTextureSupport
func (o MTLDeviceObject) ReadWriteTextureSupport() MTLReadWriteTextureTier {
	rv := objc.Send[MTLReadWriteTextureTier](o.ID, objc.Sel("readWriteTextureSupport"))
	return MTLReadWriteTextureTier(rv)
}

// An approximation of how much memory, in bytes, this GPU device can allocate
// without affecting its runtime performance.
//
// # Discussion
// 
// You can help the GPU maintain its performance by keeping the total memory
// footprint of its resources and heaps less than this threshold value.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/recommendedMaxWorkingSetSize
func (o MTLDeviceObject) RecommendedMaxWorkingSetSize() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("recommendedMaxWorkingSetSize"))
	return uint64(rv)
}

// The GPU device’s registry identifier.
//
// # Discussion
// 
// You can use the value to identify the same GPU across task boundaries
// because it’s global to all tasks.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/registryID
func (o MTLDeviceObject) RegistryID() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("registryID"))
	return uint64(rv)
}

// A Boolean value that indicates whether the GPU is removable.
//
// # Discussion
// 
// You can respond to GPU removal notifications by registering with the
// [MTLCopyAllDevicesWithObserver(handler:)] function in Swift, or the
// [MTLCopyAllDevicesWithObserver] function in Objective-C, and responding to
// the [removalRequested] and [wasRemoved] device notification names.
//
// [MTLCopyAllDevicesWithObserver(handler:)]: https://developer.apple.com/documentation/Metal/MTLCopyAllDevicesWithObserver(handler:)
// [removalRequested]: https://developer.apple.com/documentation/Metal/MTLDeviceNotificationName/removalRequested
// [wasRemoved]: https://developer.apple.com/documentation/Metal/MTLDeviceNotificationName/wasRemoved
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/isRemovable
func (o MTLDeviceObject) Removable() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isRemovable"))
	return bool(rv)
}

// Returns the size, in bytes, of a sparse tile the GPU device creates using a
// default page size.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/sparseTileSizeInBytes
func (o MTLDeviceObject) SparseTileSizeInBytes() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("sparseTileSizeInBytes"))
	return uint(rv)
}

// A Boolean value that indicates whether the GPU can filter a texture with a
// 32-bit floating-point format.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supports32BitFloatFiltering
func (o MTLDeviceObject) Supports32BitFloatFiltering() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supports32BitFloatFiltering"))
	return bool(rv)
}

// A Boolean value that indicates whether the GPU can allocate 32-bit integer
// texture formats and resolve to 32-bit floating-point texture formats.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supports32BitMSAA
func (o MTLDeviceObject) Supports32BitMSAA() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supports32BitMSAA"))
	return bool(rv)
}

// A Boolean value that indicates whether you can use textures that use BC
// compression.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsBCTextureCompression
func (o MTLDeviceObject) SupportsBCTextureCompression() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsBCTextureCompression"))
	return bool(rv)
}

// A Boolean value that indicates whether the GPU device can create and use
// dynamic libraries in compute pipelines.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsDynamicLibraries
func (o MTLDeviceObject) SupportsDynamicLibraries() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsDynamicLibraries"))
	return bool(rv)
}

// A Boolean value that indicates whether the device supports function
// pointers in compute kernel functions.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsFunctionPointers
func (o MTLDeviceObject) SupportsFunctionPointers() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsFunctionPointers"))
	return bool(rv)
}

// A Boolean value that indicates whether the device supports function
// pointers in render functions.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsFunctionPointersFromRender
func (o MTLDeviceObject) SupportsFunctionPointersFromRender() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsFunctionPointersFromRender"))
	return bool(rv)
}

// A Boolean value that indicates whether the device supports placement sparse
// resources.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsPlacementSparse
func (o MTLDeviceObject) SupportsPlacementSparse() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsPlacementSparse"))
	return bool(rv)
}

// A Boolean value that indicates whether the GPU device supports motion blur
// for ray tracing.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsPrimitiveMotionBlur
func (o MTLDeviceObject) SupportsPrimitiveMotionBlur() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsPrimitiveMotionBlur"))
	return bool(rv)
}

// A Boolean value that indicates whether the GPU can compute multiple
// interpolations of a fragment function’s input.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsPullModelInterpolation
func (o MTLDeviceObject) SupportsPullModelInterpolation() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsPullModelInterpolation"))
	return bool(rv)
}

// A Boolean value that indicates whether you can query the texture level of
// detail from within a shader.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsQueryTextureLOD
func (o MTLDeviceObject) SupportsQueryTextureLOD() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsQueryTextureLOD"))
	return bool(rv)
}

// A Boolean value that indicates whether the GPU device supports ray tracing.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsRaytracing
func (o MTLDeviceObject) SupportsRaytracing() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsRaytracing"))
	return bool(rv)
}

// A Boolean value that indicates whether you can call ray-tracing functions
// from a vertex or fragment shader.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsRaytracingFromRender
func (o MTLDeviceObject) SupportsRaytracingFromRender() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsRaytracingFromRender"))
	return bool(rv)
}

// A Boolean value that indicates whether the GPU device can create and use
// dynamic libraries in render pipelines.
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsRenderDynamicLibraries
func (o MTLDeviceObject) SupportsRenderDynamicLibraries() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsRenderDynamicLibraries"))
	return bool(rv)
}

// A Boolean value that indicates whether the GPU supports barycentric
// coordinates.
//
// # Discussion
// 
// If a GPU device supports barycentric coordinates, a fragment shader can
// receive them by adding the `[[barycentric_coord]]` attribute to one of its
// arguments. See the [Metal Shading Language specification] and [Detecting
// GPU features and Metal software versions] for more information.
//
// [Detecting GPU features and Metal software versions]: https://developer.apple.com/documentation/Metal/detecting-gpu-features-and-metal-software-versions
// [Metal Shading Language specification]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLDevice/supportsShaderBarycentricCoordinates
func (o MTLDeviceObject) SupportsShaderBarycentricCoordinates() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsShaderBarycentricCoordinates"))
	return bool(rv)
}


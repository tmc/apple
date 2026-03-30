// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that represents a graphics pipeline configuration for a render pass, which the pass applies to the draw commands you encode.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState
type MTLRenderPipelineState interface {
	objectivec.IObject
	MTLAllocation

	// The device instance that creates the pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/device
	Device() MTLDevice

	// A string that helps you identify the render pipeline state during debugging.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/label
	Label() string

	// An unique identifier that represents the pipeline state, which you can add to an argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/gpuResourceID
	GpuResourceID() MTLResourceID

	// The largest number of threads the pipeline state can have in a single object shader threadgroup.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/maxTotalThreadsPerObjectThreadgroup
	MaxTotalThreadsPerObjectThreadgroup() uint

	// The number of threads the render pass applies to a SIMD group for an object shader.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/objectThreadExecutionWidth
	ObjectThreadExecutionWidth() uint

	// The largest number of threads the pipeline state can have in a single mesh shader threadgroup.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/maxTotalThreadsPerMeshThreadgroup
	MaxTotalThreadsPerMeshThreadgroup() uint

	// The largest number of threadgroups the pipeline state can have in a single mesh shader grid.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/maxTotalThreadgroupsPerMeshGrid
	MaxTotalThreadgroupsPerMeshGrid() uint

	// The number of threads the render pass applies to a SIMD group for a mesh shader.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/meshThreadExecutionWidth
	MeshThreadExecutionWidth() uint

	// The largest number of threads the pipeline state can have in a single tile shader threadgroup.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/maxTotalThreadsPerThreadgroup
	MaxTotalThreadsPerThreadgroup() uint

	// A Boolean value that indicates whether the pipeline state needs a threadgroup’s size to equal a tile’s size.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/threadgroupSizeMatchesTileSize
	ThreadgroupSizeMatchesTileSize() bool

	// The memory size, in byes, of the render pipeline’s imageblock for a single sample.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/imageblockSampleLength
	ImageblockSampleLength() uint

	// Returns the length of an imageblock’s memory for the specified imageblock dimensions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/imageblockMemoryLength(forDimensions:)
	ImageblockMemoryLengthForDimensions(imageblockDimensions MTLSize) uint

	// A Boolean value that indicates whether the render pipeline supports encoding commands into an indirect command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/supportIndirectCommandBuffers
	SupportIndirectCommandBuffers() bool

	// The current state of shader validation for the pipeline.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/shaderValidation
	ShaderValidation() MTLShaderValidation

	// Creates a function handle for a shader.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/functionHandle(function:stage:)-7uvul
	FunctionHandleWithFunctionStage(function MTLFunction, stage MTLRenderStages) MTLFunctionHandle

	// Creates a new visible function table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/makeVisibleFunctionTable(descriptor:stage:)
	NewVisibleFunctionTableWithDescriptorStage(descriptor IMTLVisibleFunctionTableDescriptor, stage MTLRenderStages) MTLVisibleFunctionTable

	// Creates a new intersection function table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/makeIntersectionFunctionTable(descriptor:stage:)
	NewIntersectionFunctionTableWithDescriptorStage(descriptor IMTLIntersectionFunctionTableDescriptor, stage MTLRenderStages) MTLIntersectionFunctionTable

	// Creates a new pipeline state that’s a copy of the current pipeline state with additional shaders.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/makeRenderPipelineState(additionalBinaryFunctions:)-84te1
	NewRenderPipelineStateWithAdditionalBinaryFunctionsError(additionalBinaryFunctions IMTLRenderPipelineFunctionsDescriptor) (MTLRenderPipelineState, error)

	// Obtains a reflection object for this render pipeline.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/reflection
	Reflection() IMTLRenderPipelineReflection

	// RequiredThreadsPerMeshThreadgroup protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/requiredThreadsPerMeshThreadgroup
	RequiredThreadsPerMeshThreadgroup() MTLSize

	// RequiredThreadsPerObjectThreadgroup protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/requiredThreadsPerObjectThreadgroup
	RequiredThreadsPerObjectThreadgroup() MTLSize

	// RequiredThreadsPerTileThreadgroup protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/requiredThreadsPerTileThreadgroup
	RequiredThreadsPerTileThreadgroup() MTLSize

	// Obtains the function handle for a specific function this pipeline state links at the binary level.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/functionHandle(function:stage:)-1pgxo
	FunctionHandleWithBinaryFunctionStage(function MTL4BinaryFunction, stage MTLRenderStages) MTLFunctionHandle

	// Obtains a function handle for the a specific function this pipeline links at the Metal IR level.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/functionHandle(withName:stage:)
	FunctionHandleWithNameStage(name string, stage MTLRenderStages) MTLFunctionHandle

	// Creates a render pipeline descriptor from this pipeline that you can use for pipeline specialization.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/makeRenderPipelineDescriptorForSpecialization()
	NewRenderPipelineDescriptorForSpecialization() IMTL4PipelineDescriptor

	// Creates a new render pipeline state by adding binary functions to each stage of this pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/makeRenderPipelineState(additionalBinaryFunctions:)-49r1w
	NewRenderPipelineStateWithBinaryFunctionsError(binaryFunctionsDescriptor IMTL4RenderPipelineBinaryFunctionsDescriptor) (MTLRenderPipelineState, error)
}

// MTLRenderPipelineStateObject wraps an existing Objective-C object that conforms to the MTLRenderPipelineState protocol.
type MTLRenderPipelineStateObject struct {
	objectivec.Object
}

func (o MTLRenderPipelineStateObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLRenderPipelineStateObjectFromID constructs a [MTLRenderPipelineStateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLRenderPipelineStateObjectFromID(id objc.ID) MTLRenderPipelineStateObject {
	return MTLRenderPipelineStateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The device instance that creates the pipeline state.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/device
func (o MTLRenderPipelineStateObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}

// A string that helps you identify the render pipeline state during
// debugging.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/label
func (o MTLRenderPipelineStateObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// An unique identifier that represents the pipeline state, which you can add
// to an argument buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/gpuResourceID
func (o MTLRenderPipelineStateObject) GpuResourceID() MTLResourceID {
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("gpuResourceID"))
	return rv
}

// The largest number of threads the pipeline state can have in a single
// object shader threadgroup.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/maxTotalThreadsPerObjectThreadgroup
func (o MTLRenderPipelineStateObject) MaxTotalThreadsPerObjectThreadgroup() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("maxTotalThreadsPerObjectThreadgroup"))
	return rv
}

// The number of threads the render pass applies to a SIMD group for an object
// shader.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/objectThreadExecutionWidth
func (o MTLRenderPipelineStateObject) ObjectThreadExecutionWidth() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("objectThreadExecutionWidth"))
	return rv
}

// The largest number of threads the pipeline state can have in a single mesh
// shader threadgroup.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/maxTotalThreadsPerMeshThreadgroup
func (o MTLRenderPipelineStateObject) MaxTotalThreadsPerMeshThreadgroup() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("maxTotalThreadsPerMeshThreadgroup"))
	return rv
}

// The largest number of threadgroups the pipeline state can have in a single
// mesh shader grid.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/maxTotalThreadgroupsPerMeshGrid
func (o MTLRenderPipelineStateObject) MaxTotalThreadgroupsPerMeshGrid() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("maxTotalThreadgroupsPerMeshGrid"))
	return rv
}

// The number of threads the render pass applies to a SIMD group for a mesh
// shader.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/meshThreadExecutionWidth
func (o MTLRenderPipelineStateObject) MeshThreadExecutionWidth() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("meshThreadExecutionWidth"))
	return rv
}

// The largest number of threads the pipeline state can have in a single tile
// shader threadgroup.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/maxTotalThreadsPerThreadgroup
func (o MTLRenderPipelineStateObject) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}

// A Boolean value that indicates whether the pipeline state needs a
// threadgroup’s size to equal a tile’s size.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/threadgroupSizeMatchesTileSize
func (o MTLRenderPipelineStateObject) ThreadgroupSizeMatchesTileSize() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("threadgroupSizeMatchesTileSize"))
	return rv
}

// The memory size, in byes, of the render pipeline’s imageblock for a
// single sample.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/imageblockSampleLength
func (o MTLRenderPipelineStateObject) ImageblockSampleLength() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("imageblockSampleLength"))
	return rv
}

// Returns the length of an imageblock’s memory for the specified imageblock
// dimensions.
//
// imageblockDimensions: An [MTLSize] instance that represent the dimensions of an imageblock.
//
// # Discussion
//
// The imageblock dimensions need to match a valid tile size, such as one of
// the following:
//
// - 32 x 32
// - 32 x 16
// - 16 x 16
//
// The GPU partitions tile memory between imageblocks and threadgroup memory,
//
// For information about identifying tile memory limits for GPU devices, see
// either of the following:
//
// - [Metal Feature Set Tables (PDF)]
// - [Metal Feature Set Tables (Numbers)]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/imageblockMemoryLength(forDimensions:)
//
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
// [Metal Feature Set Tables (Numbers)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.zip
// [Metal Feature Set Tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
func (o MTLRenderPipelineStateObject) ImageblockMemoryLengthForDimensions(imageblockDimensions MTLSize) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("imageblockMemoryLengthForDimensions:"), imageblockDimensions)
	return rv
}

// A Boolean value that indicates whether the render pipeline supports
// encoding commands into an indirect command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/supportIndirectCommandBuffers
func (o MTLRenderPipelineStateObject) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}

// The current state of shader validation for the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/shaderValidation
func (o MTLRenderPipelineStateObject) ShaderValidation() MTLShaderValidation {
	rv := objc.Send[MTLShaderValidation](o.ID, objc.Sel("shaderValidation"))
	return rv
}

// Creates a function handle for a shader.
//
// function: An [MTLFunction] instance that represents the shader the method creates a
// handle for.
//
// stage: An [MTLRenderStages] instance that represents the rendering stage that
// invokes the shader that `function` represents.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/functionHandle(function:stage:)-7uvul
//
// [MTLRenderStages]: https://developer.apple.com/documentation/Metal/MTLRenderStages
func (o MTLRenderPipelineStateObject) FunctionHandleWithFunctionStage(function MTLFunction, stage MTLRenderStages) MTLFunctionHandle {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionHandleWithFunction:stage:"), function, stage)
	return MTLFunctionHandleObjectFromID(rv)
}

// Creates a new visible function table.
//
// descriptor: An [MTLVisibleFunctionTableDescriptor] instance that configures the visible
// function table the method creates.
//
// stage: An [MTLRenderStages] instance that represents the render pass stage the
// visible function table applies to.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/makeVisibleFunctionTable(descriptor:stage:)
//
// [MTLRenderStages]: https://developer.apple.com/documentation/Metal/MTLRenderStages
func (o MTLRenderPipelineStateObject) NewVisibleFunctionTableWithDescriptorStage(descriptor IMTLVisibleFunctionTableDescriptor, stage MTLRenderStages) MTLVisibleFunctionTable {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newVisibleFunctionTableWithDescriptor:stage:"), descriptor, stage)
	return MTLVisibleFunctionTableObjectFromID(rv)
}

// Creates a new intersection function table.
//
// descriptor: An [MTLIntersectionFunctionTableDescriptor] instance that configures the
// visible function table the method creates.
//
// stage: An [MTLRenderStages] instance that represents the render pass stage the
// intersection function table applies to.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/makeIntersectionFunctionTable(descriptor:stage:)
//
// [MTLRenderStages]: https://developer.apple.com/documentation/Metal/MTLRenderStages
func (o MTLRenderPipelineStateObject) NewIntersectionFunctionTableWithDescriptorStage(descriptor IMTLIntersectionFunctionTableDescriptor, stage MTLRenderStages) MTLIntersectionFunctionTable {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newIntersectionFunctionTableWithDescriptor:stage:"), descriptor, stage)
	return MTLIntersectionFunctionTableObjectFromID(rv)
}

// Creates a new pipeline state that’s a copy of the current pipeline state
// with additional shaders.
//
// additionalBinaryFunctions: An [MTLRenderPipelineFunctionsDescriptor] instance, which contains
// [MTLFunction] arrays for vertex, fragment, and tile shaders.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/makeRenderPipelineState(additionalBinaryFunctions:)-84te1
func (o MTLRenderPipelineStateObject) NewRenderPipelineStateWithAdditionalBinaryFunctionsError(additionalBinaryFunctions IMTLRenderPipelineFunctionsDescriptor) (MTLRenderPipelineState, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithAdditionalBinaryFunctions:error:"), additionalBinaryFunctions)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
}

// Obtains a reflection object for this render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/reflection
func (o MTLRenderPipelineStateObject) Reflection() IMTLRenderPipelineReflection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("reflection"))
	return MTLRenderPipelineReflectionFromID(rv)
}

// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/requiredThreadsPerMeshThreadgroup
func (o MTLRenderPipelineStateObject) RequiredThreadsPerMeshThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](o.ID, objc.Sel("requiredThreadsPerMeshThreadgroup"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/requiredThreadsPerObjectThreadgroup
func (o MTLRenderPipelineStateObject) RequiredThreadsPerObjectThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](o.ID, objc.Sel("requiredThreadsPerObjectThreadgroup"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/requiredThreadsPerTileThreadgroup
func (o MTLRenderPipelineStateObject) RequiredThreadsPerTileThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](o.ID, objc.Sel("requiredThreadsPerTileThreadgroup"))
	return rv
}

// Obtains the function handle for a specific function this pipeline state
// links at the binary level.
//
// function: A binary function to retrieve the handle.
//
// stage: The shader stage that uses the function.
//
// # Return Value
//
// A function handle representing the function if present, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/functionHandle(function:stage:)-1pgxo
func (o MTLRenderPipelineStateObject) FunctionHandleWithBinaryFunctionStage(function MTL4BinaryFunction, stage MTLRenderStages) MTLFunctionHandle {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionHandleWithBinaryFunction:stage:"), function, stage)
	return MTLFunctionHandleObjectFromID(rv)
}

// Obtains a function handle for the a specific function this pipeline links
// at the Metal IR level.
//
// name: A string containing the name of the function.
//
// stage: The shader stage that uses the function.
//
// # Return Value
//
// A function handle representing the function if present, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/functionHandle(withName:stage:)
func (o MTLRenderPipelineStateObject) FunctionHandleWithNameStage(name string, stage MTLRenderStages) MTLFunctionHandle {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionHandleWithName:stage:"), objc.String(name), stage)
	return MTLFunctionHandleObjectFromID(rv)
}

// Creates a render pipeline descriptor from this pipeline that you can use
// for pipeline specialization.
//
// # Return Value
//
// A new pipeline descriptor that you use for pipeline state specialization.
//
// # Discussion
//
// Use this method to obtain a new [MTL4PipelineDescriptor] instance that you
// can use to specialize any unspecialized properties in this pipeline state
// object.
//
// The returned descriptor contains every unspecialized field in the current
// pipeline state object, set to unspecialized. It may, however, not contain
// valid or accurate properties in any other field.
//
// This descriptor is only valid for the purpose of calling specialization
// functions on the [MTL4Compiler] to specialize this pipeline, for example:
// [NewRenderPipelineStateBySpecializationWithDescriptorPipelineError].
//
// Although this method returns the [MTL4PipelineDescriptor] base class, the
// concrete instance this method returns corresponds to the specific
// descriptor type for the creation of this pipeline state, for example if a
// [MTL4Compiler] instance creates this current pipeline form a
// [MTLTileRenderPipelineDescriptor], this method returns a concrete
// [MTLTileRenderPipelineDescriptor] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/makeRenderPipelineDescriptorForSpecialization()
func (o MTLRenderPipelineStateObject) NewRenderPipelineDescriptorForSpecialization() IMTL4PipelineDescriptor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newRenderPipelineDescriptorForSpecialization"))
	return MTL4PipelineDescriptorFromID(rv)
}

// Creates a new render pipeline state by adding binary functions to each
// stage of this pipeline state.
//
// binaryFunctionsDescriptor: A non-`nil` dynamic linking descriptor.
//
// # Return Value
//
// A new render pipeline state upon success, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineState/makeRenderPipelineState(additionalBinaryFunctions:)-49r1w
func (o MTLRenderPipelineStateObject) NewRenderPipelineStateWithBinaryFunctionsError(binaryFunctionsDescriptor IMTL4RenderPipelineBinaryFunctionsDescriptor) (MTLRenderPipelineState, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newRenderPipelineStateWithBinaryFunctions:error:"), binaryFunctionsDescriptor)
	if err != nil {
		return nil, err
	}
	return MTLRenderPipelineStateObjectFromID(rv), nil
}

// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize
func (o MTLRenderPipelineStateObject) AllocatedSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
}

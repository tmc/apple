// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that represents a GPU pipeline configuration for running kernels in a compute pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState
type MTLComputePipelineState interface {
	objectivec.IObject
	MTLAllocation

	// The device instance that created the pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/device
	Device() MTLDevice

	// An unique identifier that represents the pipeline state, which you can add to an argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/gpuResourceID
	GpuResourceID() MTLResourceID

	// A string that helps you identify the compute pipeline state during debugging.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/label
	Label() string

	// The maximum number of threads in a threadgroup that you can dispatch to the pipeline.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/maxTotalThreadsPerThreadgroup
	MaxTotalThreadsPerThreadgroup() uint

	// The number of threads that the GPU executes simultaneously.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/threadExecutionWidth
	ThreadExecutionWidth() uint

	// The length, in bytes, of statically allocated threadgroup memory.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/staticThreadgroupMemoryLength
	StaticThreadgroupMemoryLength() uint

	// Returns the length of reserved memory for an imageblock of a given size.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/imageblockMemoryLength(forDimensions:)
	ImageblockMemoryLengthForDimensions(imageblockDimensions MTLSize) uint

	// A Boolean value that indicates whether the compute pipeline supports indirect command buffers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/supportIndirectCommandBuffers
	SupportIndirectCommandBuffers() bool

	// The current state of shader validation for the pipeline.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/shaderValidation
	ShaderValidation() MTLShaderValidation

	// Creates a function handle for a visible function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/functionHandle(function:)-7d523
	FunctionHandleWithFunction(function MTLFunction) MTLFunctionHandle

	// Creates a new pipeline state object with additional callable functions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/makeComputePipelineStateWithAdditionalBinaryFunctions(functions:)
	NewComputePipelineStateWithAdditionalBinaryFunctionsError(functions []objectivec.IObject) (MTLComputePipelineState, error)

	// Creates a new visible function table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/makeVisibleFunctionTable(descriptor:)
	NewVisibleFunctionTableWithDescriptor(descriptor IMTLVisibleFunctionTableDescriptor) MTLVisibleFunctionTable

	// Creates a new intersection function table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/makeIntersectionFunctionTable(descriptor:)
	NewIntersectionFunctionTableWithDescriptor(descriptor IMTLIntersectionFunctionTableDescriptor) MTLIntersectionFunctionTable

	// Provides access to this compute pipeline’s reflection.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/reflection
	Reflection() IMTLComputePipelineReflection

	// RequiredThreadsPerThreadgroup protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/requiredThreadsPerThreadgroup
	RequiredThreadsPerThreadgroup() MTLSize

	// Gets the function handle for a function this pipeline links at the binary level.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/functionHandle(function:)-8spaa
	FunctionHandleWithBinaryFunction(function MTL4BinaryFunction) MTLFunctionHandle

	// Gets the function handle for a function this pipeline links at the Metal IR level by name.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/functionHandle(withName:)
	FunctionHandleWithName(name string) MTLFunctionHandle

	// Allocates a new compute pipeline state by adding binary functions to this pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/makeComputePipelineState(additionalBinaryFunctions:)
	NewComputePipelineStateWithBinaryFunctionsError(additionalBinaryFunctions []objectivec.IObject) (MTLComputePipelineState, error)
}

// MTLComputePipelineStateObject wraps an existing Objective-C object that conforms to the MTLComputePipelineState protocol.
type MTLComputePipelineStateObject struct {
	objectivec.Object
}

func (o MTLComputePipelineStateObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLComputePipelineStateObjectFromID constructs a [MTLComputePipelineStateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLComputePipelineStateObjectFromID(id objc.ID) MTLComputePipelineStateObject {
	return MTLComputePipelineStateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The device instance that created the pipeline state.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/device
func (o MTLComputePipelineStateObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}

// An unique identifier that represents the pipeline state, which you can add
// to an argument buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/gpuResourceID
func (o MTLComputePipelineStateObject) GpuResourceID() MTLResourceID {
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("gpuResourceID"))
	return rv
}

// A string that helps you identify the compute pipeline state during
// debugging.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/label
func (o MTLComputePipelineStateObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// The maximum number of threads in a threadgroup that you can dispatch to the
// pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/maxTotalThreadsPerThreadgroup
func (o MTLComputePipelineStateObject) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}

// The number of threads that the GPU executes simultaneously.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/threadExecutionWidth
func (o MTLComputePipelineStateObject) ThreadExecutionWidth() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("threadExecutionWidth"))
	return rv
}

// The length, in bytes, of statically allocated threadgroup memory.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/staticThreadgroupMemoryLength
func (o MTLComputePipelineStateObject) StaticThreadgroupMemoryLength() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("staticThreadgroupMemoryLength"))
	return rv
}

// Returns the length of reserved memory for an imageblock of a given size.
//
// imageblockDimensions: An [MTLSize] instance that represents the dimensions of an imageblock.
//
// # Return Value
//
// The length, in bytes, occupied by the image block in memory.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/imageblockMemoryLength(forDimensions:)
//
// [MTLSize]: https://developer.apple.com/documentation/Metal/MTLSize
func (o MTLComputePipelineStateObject) ImageblockMemoryLengthForDimensions(imageblockDimensions MTLSize) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("imageblockMemoryLengthForDimensions:"), imageblockDimensions)
	return rv
}

// A Boolean value that indicates whether the compute pipeline supports
// indirect command buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/supportIndirectCommandBuffers
func (o MTLComputePipelineStateObject) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}

// The current state of shader validation for the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/shaderValidation
func (o MTLComputePipelineStateObject) ShaderValidation() MTLShaderValidation {
	rv := objc.Send[MTLShaderValidation](o.ID, objc.Sel("shaderValidation"))
	return rv
}

// Creates a function handle for a visible function.
//
// function: An [MTLFunction] instance that represents the visible function to create a
// handle for.
//
// # Return Value
//
// A handle to the visible function. When this value is `nil`, an error
// occurred during handle creation.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/functionHandle(function:)-7d523
func (o MTLComputePipelineStateObject) FunctionHandleWithFunction(function MTLFunction) MTLFunctionHandle {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionHandleWithFunction:"), function)
	return MTLFunctionHandleObjectFromID(rv)
}

// Creates a new pipeline state object with additional callable functions.
//
// functions: The list of additional functions that you want to be able to call.
//
// # Return Value
//
// A new compute pipeline state with access to the provided functions. When
// this value is `nil`, an error occurred during handle creation.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/makeComputePipelineStateWithAdditionalBinaryFunctions(functions:)
func (o MTLComputePipelineStateObject) NewComputePipelineStateWithAdditionalBinaryFunctionsError(functions []objectivec.IObject) (MTLComputePipelineState, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithAdditionalBinaryFunctions:error:"), objectivec.IObjectSliceToNSArray(functions))
	if err != nil {
		return nil, err
	}
	return MTLComputePipelineStateObjectFromID(rv), nil
}

// Creates a new visible function table.
//
// descriptor: An [MTLVisibleFunctionTableDescriptor] instance that configures the created
// table.
//
// # Return Value
//
// A new visible function table, or `nil` if an error occurred in creation.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/makeVisibleFunctionTable(descriptor:)
func (o MTLComputePipelineStateObject) NewVisibleFunctionTableWithDescriptor(descriptor IMTLVisibleFunctionTableDescriptor) MTLVisibleFunctionTable {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newVisibleFunctionTableWithDescriptor:"), descriptor)
	return MTLVisibleFunctionTableObjectFromID(rv)
}

// Creates a new intersection function table.
//
// descriptor: An [MTLIntersectionFunctionTableDescriptor] instance that configures the
// created table.
//
// # Return Value
//
// A new intersection function table, or `nil` if an error occurred in
// creation.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/makeIntersectionFunctionTable(descriptor:)
func (o MTLComputePipelineStateObject) NewIntersectionFunctionTableWithDescriptor(descriptor IMTLIntersectionFunctionTableDescriptor) MTLIntersectionFunctionTable {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newIntersectionFunctionTableWithDescriptor:"), descriptor)
	return MTLIntersectionFunctionTableObjectFromID(rv)
}

// Provides access to this compute pipeline’s reflection.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/reflection
func (o MTLComputePipelineStateObject) Reflection() IMTLComputePipelineReflection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("reflection"))
	return MTLComputePipelineReflectionFromID(rv)
}

// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/requiredThreadsPerThreadgroup
func (o MTLComputePipelineStateObject) RequiredThreadsPerThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](o.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}

// Gets the function handle for a function this pipeline links at the binary
// level.
//
// function: A binary function object representing the function binary to find.
//
// # Return Value
//
// A function handle corresponding to the function if the binary function
// matches a function in this pipeline state, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/functionHandle(function:)-8spaa
func (o MTLComputePipelineStateObject) FunctionHandleWithBinaryFunction(function MTL4BinaryFunction) MTLFunctionHandle {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionHandleWithBinaryFunction:"), function)
	return MTLFunctionHandleObjectFromID(rv)
}

// Gets the function handle for a function this pipeline links at the Metal IR
// level by name.
//
// name: A string representing the name of the function.
//
// # Return Value
//
// A function handle corresponding to the function if the name matches a
// function in this pipeline state, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/functionHandle(withName:)
func (o MTLComputePipelineStateObject) FunctionHandleWithName(name string) MTLFunctionHandle {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionHandleWithName:"), objc.String(name))
	return MTLFunctionHandleObjectFromID(rv)
}

// Allocates a new compute pipeline state by adding binary functions to this
// pipeline state.
//
// additionalBinaryFunctions: A non-`nil` array containing binary functions to add to this pipeline.
//
// # Return Value
//
// A new compute pipeline state upon success, otherwise `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineState/makeComputePipelineState(additionalBinaryFunctions:)
func (o MTLComputePipelineStateObject) NewComputePipelineStateWithBinaryFunctionsError(additionalBinaryFunctions []objectivec.IObject) (MTLComputePipelineState, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newComputePipelineStateWithBinaryFunctions:error:"), objectivec.IObjectSliceToNSArray(additionalBinaryFunctions))
	if err != nil {
		return nil, err
	}
	return MTLComputePipelineStateObjectFromID(rv), nil
}

// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize
func (o MTLComputePipelineStateObject) AllocatedSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
}

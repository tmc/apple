// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLComputePipelineDescriptor] class.
var (
	_MTLComputePipelineDescriptorClass     MTLComputePipelineDescriptorClass
	_MTLComputePipelineDescriptorClassOnce sync.Once
)

func getMTLComputePipelineDescriptorClass() MTLComputePipelineDescriptorClass {
	_MTLComputePipelineDescriptorClassOnce.Do(func() {
		_MTLComputePipelineDescriptorClass = MTLComputePipelineDescriptorClass{class: objc.GetClass("MTLComputePipelineDescriptor")}
	})
	return _MTLComputePipelineDescriptorClass
}

// GetMTLComputePipelineDescriptorClass returns the class object for MTLComputePipelineDescriptor.
func GetMTLComputePipelineDescriptorClass() MTLComputePipelineDescriptorClass {
	return getMTLComputePipelineDescriptorClass()
}

type MTLComputePipelineDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLComputePipelineDescriptorClass) Alloc() MTLComputePipelineDescriptor {
	rv := objc.Send[MTLComputePipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// An instance describing the desired GPU state for a kernel call in a compute
// pass.
//
// # Overview
// 
// A pipeline descriptor provides information necessary for creating an
// [MTLComputePipelineState] instance.
//
// # Configuring the compute execution environment
//
//   - [MTLComputePipelineDescriptor.ComputeFunction]: The compute kernel the pipeline calls.
//   - [MTLComputePipelineDescriptor.SetComputeFunction]
//   - [MTLComputePipelineDescriptor.ThreadGroupSizeIsMultipleOfThreadExecutionWidth]: A Boolean value that indicates whether the threadgroup size is always a multiple of the thread execution width.
//   - [MTLComputePipelineDescriptor.SetThreadGroupSizeIsMultipleOfThreadExecutionWidth]
//   - [MTLComputePipelineDescriptor.MaxTotalThreadsPerThreadgroup]: A property that limits the number of threads you can dispatch in a threadgroup for the compute function.
//   - [MTLComputePipelineDescriptor.SetMaxTotalThreadsPerThreadgroup]
//   - [MTLComputePipelineDescriptor.MaxCallStackDepth]: The maximum call stack depth for indirect function calls in compute shaders.
//   - [MTLComputePipelineDescriptor.SetMaxCallStackDepth]
//
// # Configuring compute pass inputs
//
//   - [MTLComputePipelineDescriptor.StageInputDescriptor]: The organization of input and output data for the next kernel call.
//   - [MTLComputePipelineDescriptor.SetStageInputDescriptor]
//
// # Configuring buffer mutability
//
//   - [MTLComputePipelineDescriptor.Buffers]: The buffer mutability options to apply to the next kernel call.
//
// # Identifying the pipeline state object
//
//   - [MTLComputePipelineDescriptor.Label]: A string that identifies the instance.
//   - [MTLComputePipelineDescriptor.SetLabel]
//
// # Configuring indirect command buffers
//
//   - [MTLComputePipelineDescriptor.SupportIndirectCommandBuffers]: A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer.
//   - [MTLComputePipelineDescriptor.SetSupportIndirectCommandBuffers]
//
// # Configuring shader validation
//
//   - [MTLComputePipelineDescriptor.ShaderValidation]: A value that enables or disables shader validation for the pipeline.
//   - [MTLComputePipelineDescriptor.SetShaderValidation]
//
// # Reset to defaults
//
//   - [MTLComputePipelineDescriptor.Reset]: Resets all compute pipeline descriptor properties to their default values.
//
// # Loading dynamic libraries to link at runtime
//
//   - [MTLComputePipelineDescriptor.PreloadedLibraries]: The dynamic libraries that contain precompiled shader functions you want to link.
//   - [MTLComputePipelineDescriptor.SetPreloadedLibraries]
//
// # Setting callable functions
//
//   - [MTLComputePipelineDescriptor.LinkedFunctions]: The functions with available function pointers for the next kernel call.
//   - [MTLComputePipelineDescriptor.SetLinkedFunctions]
//
// # Loading binary archives
//
//   - [MTLComputePipelineDescriptor.SupportAddingBinaryFunctions]: A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//   - [MTLComputePipelineDescriptor.SetSupportAddingBinaryFunctions]
//   - [MTLComputePipelineDescriptor.BinaryArchives]: The binary archives that contain any precompiled shader functions to link.
//   - [MTLComputePipelineDescriptor.SetBinaryArchives]
//
// # Instance Properties
//
//   - [MTLComputePipelineDescriptor.RequiredThreadsPerThreadgroup]
//   - [MTLComputePipelineDescriptor.SetRequiredThreadsPerThreadgroup]
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor
type MTLComputePipelineDescriptor struct {
	objectivec.Object
}

// MTLComputePipelineDescriptorFromID constructs a [MTLComputePipelineDescriptor] from an objc.ID.
//
// An instance describing the desired GPU state for a kernel call in a compute
// pass.
func MTLComputePipelineDescriptorFromID(id objc.ID) MTLComputePipelineDescriptor {
	return MTLComputePipelineDescriptor{objectivec.Object{id}}
}
// NOTE: MTLComputePipelineDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLComputePipelineDescriptor] class.
//
// # Configuring the compute execution environment
//
//   - [IMTLComputePipelineDescriptor.ComputeFunction]: The compute kernel the pipeline calls.
//   - [IMTLComputePipelineDescriptor.SetComputeFunction]
//   - [IMTLComputePipelineDescriptor.ThreadGroupSizeIsMultipleOfThreadExecutionWidth]: A Boolean value that indicates whether the threadgroup size is always a multiple of the thread execution width.
//   - [IMTLComputePipelineDescriptor.SetThreadGroupSizeIsMultipleOfThreadExecutionWidth]
//   - [IMTLComputePipelineDescriptor.MaxTotalThreadsPerThreadgroup]: A property that limits the number of threads you can dispatch in a threadgroup for the compute function.
//   - [IMTLComputePipelineDescriptor.SetMaxTotalThreadsPerThreadgroup]
//   - [IMTLComputePipelineDescriptor.MaxCallStackDepth]: The maximum call stack depth for indirect function calls in compute shaders.
//   - [IMTLComputePipelineDescriptor.SetMaxCallStackDepth]
//
// # Configuring compute pass inputs
//
//   - [IMTLComputePipelineDescriptor.StageInputDescriptor]: The organization of input and output data for the next kernel call.
//   - [IMTLComputePipelineDescriptor.SetStageInputDescriptor]
//
// # Configuring buffer mutability
//
//   - [IMTLComputePipelineDescriptor.Buffers]: The buffer mutability options to apply to the next kernel call.
//
// # Identifying the pipeline state object
//
//   - [IMTLComputePipelineDescriptor.Label]: A string that identifies the instance.
//   - [IMTLComputePipelineDescriptor.SetLabel]
//
// # Configuring indirect command buffers
//
//   - [IMTLComputePipelineDescriptor.SupportIndirectCommandBuffers]: A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer.
//   - [IMTLComputePipelineDescriptor.SetSupportIndirectCommandBuffers]
//
// # Configuring shader validation
//
//   - [IMTLComputePipelineDescriptor.ShaderValidation]: A value that enables or disables shader validation for the pipeline.
//   - [IMTLComputePipelineDescriptor.SetShaderValidation]
//
// # Reset to defaults
//
//   - [IMTLComputePipelineDescriptor.Reset]: Resets all compute pipeline descriptor properties to their default values.
//
// # Loading dynamic libraries to link at runtime
//
//   - [IMTLComputePipelineDescriptor.PreloadedLibraries]: The dynamic libraries that contain precompiled shader functions you want to link.
//   - [IMTLComputePipelineDescriptor.SetPreloadedLibraries]
//
// # Setting callable functions
//
//   - [IMTLComputePipelineDescriptor.LinkedFunctions]: The functions with available function pointers for the next kernel call.
//   - [IMTLComputePipelineDescriptor.SetLinkedFunctions]
//
// # Loading binary archives
//
//   - [IMTLComputePipelineDescriptor.SupportAddingBinaryFunctions]: A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//   - [IMTLComputePipelineDescriptor.SetSupportAddingBinaryFunctions]
//   - [IMTLComputePipelineDescriptor.BinaryArchives]: The binary archives that contain any precompiled shader functions to link.
//   - [IMTLComputePipelineDescriptor.SetBinaryArchives]
//
// # Instance Properties
//
//   - [IMTLComputePipelineDescriptor.RequiredThreadsPerThreadgroup]
//   - [IMTLComputePipelineDescriptor.SetRequiredThreadsPerThreadgroup]
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor
type IMTLComputePipelineDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the compute execution environment

	// The compute kernel the pipeline calls.
	ComputeFunction() MTLFunction
	SetComputeFunction(value MTLFunction)
	// A Boolean value that indicates whether the threadgroup size is always a multiple of the thread execution width.
	ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool)
	// A property that limits the number of threads you can dispatch in a threadgroup for the compute function.
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	// The maximum call stack depth for indirect function calls in compute shaders.
	MaxCallStackDepth() uint
	SetMaxCallStackDepth(value uint)

	// Topic: Configuring compute pass inputs

	// The organization of input and output data for the next kernel call.
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)

	// Topic: Configuring buffer mutability

	// The buffer mutability options to apply to the next kernel call.
	Buffers() IMTLPipelineBufferDescriptorArray

	// Topic: Identifying the pipeline state object

	// A string that identifies the instance.
	Label() string
	SetLabel(value string)

	// Topic: Configuring indirect command buffers

	// A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer.
	SupportIndirectCommandBuffers() bool
	SetSupportIndirectCommandBuffers(value bool)

	// Topic: Configuring shader validation

	// A value that enables or disables shader validation for the pipeline.
	ShaderValidation() MTLShaderValidation
	SetShaderValidation(value MTLShaderValidation)

	// Topic: Reset to defaults

	// Resets all compute pipeline descriptor properties to their default values.
	Reset()

	// Topic: Loading dynamic libraries to link at runtime

	// The dynamic libraries that contain precompiled shader functions you want to link.
	PreloadedLibraries() []objectivec.IObject
	SetPreloadedLibraries(value []objectivec.IObject)

	// Topic: Setting callable functions

	// The functions with available function pointers for the next kernel call.
	LinkedFunctions() IMTLLinkedFunctions
	SetLinkedFunctions(value IMTLLinkedFunctions)

	// Topic: Loading binary archives

	// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
	SupportAddingBinaryFunctions() bool
	SetSupportAddingBinaryFunctions(value bool)
	// The binary archives that contain any precompiled shader functions to link.
	BinaryArchives() []objectivec.IObject
	SetBinaryArchives(value []objectivec.IObject)

	// Topic: Instance Properties

	RequiredThreadsPerThreadgroup() MTLSize
	SetRequiredThreadsPerThreadgroup(value MTLSize)
}





// Init initializes the instance.
func (c MTLComputePipelineDescriptor) Init() MTLComputePipelineDescriptor {
	rv := objc.Send[MTLComputePipelineDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLComputePipelineDescriptor) Autorelease() MTLComputePipelineDescriptor {
	rv := objc.Send[MTLComputePipelineDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLComputePipelineDescriptor creates a new MTLComputePipelineDescriptor instance.
func NewMTLComputePipelineDescriptor() MTLComputePipelineDescriptor {
	class := getMTLComputePipelineDescriptorClass()
	rv := objc.Send[MTLComputePipelineDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Resets all compute pipeline descriptor properties to their default values.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/reset()
func (c MTLComputePipelineDescriptor) Reset() {
	objc.Send[objc.ID](c.ID, objc.Sel("reset"))
}












// The compute kernel the pipeline calls.
//
// # Discussion
// 
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/computeFunction
func (c MTLComputePipelineDescriptor) ComputeFunction() MTLFunction {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("computeFunction"))
	return MTLFunctionObjectFromID(rv)
}
func (c MTLComputePipelineDescriptor) SetComputeFunction(value MTLFunction) {
	objc.Send[struct{}](c.ID, objc.Sel("setComputeFunction:"), value)
}



// A Boolean value that indicates whether the threadgroup size is always a
// multiple of the thread execution width.
//
// # Discussion
// 
// If you can guarantee that the threadgroup size used by all compute commands
// in this pipeline is a multiple of [ThreadExecutionWidth], set this property
// to `true` to take advantage of additional Metal optimizations.
// 
// The default value is `false`.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/threadGroupSizeIsMultipleOfThreadExecutionWidth
func (c MTLComputePipelineDescriptor) ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("threadGroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}
func (c MTLComputePipelineDescriptor) SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setThreadGroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}



// A property that limits the number of threads you can dispatch in a
// threadgroup for the compute function.
//
// # Discussion
// 
// Metal automatically selects a maximum threadgroup size when you set this
// value to `0`.
// 
// Your shader can also configure the maximum number of threads per
// threadgroup with the `[[max_total_threads_per_threadgroup]]` attribute. See
// the [Metal Shading Language Specification] for more information.
// 
// By default, this property’s value is `0`, which instructs Metal to
// calculate the maximum number of threads per threadgroup based on the
// device’s capabilities and the compute shader’s memory usage.
// 
// The [MaxTotalThreadsPerThreadgroup] property of an
// [MTLComputePipelineState] instance reports the maximum number of threads
// you can dispatch in a threadgroup for that specific compute shader.
// 
// Metal may return an error if this value exceeds the available resources for
// the device, or Metal may lower the thread limit when creating the compute
// pipeline state, which can reduce runtime performance.
//
// [Metal Shading Language Specification]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (c MTLComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}
func (c MTLComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}



// The maximum call stack depth for indirect function calls in compute
// shaders.
//
// # Discussion
// 
// The property’s default value is `1`. Change its value if you use
// recursive functions in your compute pass.
// 
// The maximum call stack depth applies only to indirect function calls in
// your shader, and affects the upper bound of stack memory for each thread.
// Indirect function calls include those to visible functions, intersection
// functions, and to dynamic libraries.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/maxCallStackDepth
func (c MTLComputePipelineDescriptor) MaxCallStackDepth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("maxCallStackDepth"))
	return rv
}
func (c MTLComputePipelineDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxCallStackDepth:"), value)
}



// The organization of input and output data for the next kernel call.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/stageInputDescriptor
func (c MTLComputePipelineDescriptor) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("stageInputDescriptor"))
	return MTLStageInputOutputDescriptorFromID(objc.ID(rv))
}
func (c MTLComputePipelineDescriptor) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[struct{}](c.ID, objc.Sel("setStageInputDescriptor:"), value)
}



// The buffer mutability options to apply to the next kernel call.
//
// # Discussion
// 
// This property holds an array of [MTLPipelineBufferDescriptor] instances,
// where each index corresponds to the same entry in the buffer argument
// table.
// 
// Metal can perform additional optimizations if you guarantee that neither
// the CPU nor the GPU modify a buffer’s contents after set in a
// function’s argument table and before its command buffer completes. Use
// immutable buffers as much as possible, for either regular buffers or
// argument buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/buffers
func (c MTLComputePipelineDescriptor) Buffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("buffers"))
	return MTLPipelineBufferDescriptorArrayFromID(objc.ID(rv))
}



// A string that identifies the instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/label
func (c MTLComputePipelineDescriptor) Label() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (c MTLComputePipelineDescriptor) SetLabel(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLabel:"), objc.String(value))
}



// A Boolean value that indicates whether you can encode commands that
// reference the pipeline state object into an indirect command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/supportIndirectCommandBuffers
func (c MTLComputePipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}
func (c MTLComputePipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}



// A value that enables or disables shader validation for the pipeline.
//
// # Discussion
// 
// You can override the value using either of these environment variables:
// `MTL_SHADER_VALIDATION_ENABLE_PIPELINES` or
// `MTL_SHADER_VALIDATION_DISABLE_PIPELINES.`
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/shaderValidation
func (c MTLComputePipelineDescriptor) ShaderValidation() MTLShaderValidation {
	rv := objc.Send[MTLShaderValidation](c.ID, objc.Sel("shaderValidation"))
	return MTLShaderValidation(rv)
}
func (c MTLComputePipelineDescriptor) SetShaderValidation(value MTLShaderValidation) {
	objc.Send[struct{}](c.ID, objc.Sel("setShaderValidation:"), value)
}



// The dynamic libraries that contain precompiled shader functions you want to
// link.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/preloadedLibraries
func (c MTLComputePipelineDescriptor) PreloadedLibraries() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("preloadedLibraries"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (c MTLComputePipelineDescriptor) SetPreloadedLibraries(value []objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreloadedLibraries:"), objectivec.IObjectSliceToNSArray(value))
}



// The functions with available function pointers for the next kernel call.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/linkedFunctions
func (c MTLComputePipelineDescriptor) LinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("linkedFunctions"))
	return MTLLinkedFunctionsFromID(objc.ID(rv))
}
func (c MTLComputePipelineDescriptor) SetLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[struct{}](c.ID, objc.Sel("setLinkedFunctions:"), value)
}



// A Boolean value that indicates whether you can use the pipeline to create
// new pipelines by adding binary functions to its callable functions list.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/supportAddingBinaryFunctions
func (c MTLComputePipelineDescriptor) SupportAddingBinaryFunctions() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("supportAddingBinaryFunctions"))
	return rv
}
func (c MTLComputePipelineDescriptor) SetSupportAddingBinaryFunctions(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSupportAddingBinaryFunctions:"), value)
}



// The binary archives that contain any precompiled shader functions to link.
//
// # Discussion
// 
// The default value is `nil`.
// 
// When you create a Metal library, Metal compiles shader functions into an
// intermediate representation. When you create the pipeline state object, the
// GPU compiles this intermediate code.
// 
// By providing a set of binary archives, when Metal creates the pipeline
// state object, it first checks the archives to see if there’s already a
// compiled function. If so, Metal uses it instead.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/binaryArchives
func (c MTLComputePipelineDescriptor) BinaryArchives() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("binaryArchives"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (c MTLComputePipelineDescriptor) SetBinaryArchives(value []objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setBinaryArchives:"), objectivec.IObjectSliceToNSArray(value))
}



//
// # Discussion
// 
// Sets the required threads-per-threadgroup during dispatches. The
// `threadsPerThreadgroup` argument of any dispatch must match this value if
// it is set. Optional, unless the pipeline is going to use CooperativeTensors
// in which case this must be set. Setting this to a size of 0 in every
// dimension disables this property
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/requiredThreadsPerThreadgroup
func (c MTLComputePipelineDescriptor) RequiredThreadsPerThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](c.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return MTLSize(rv)
}
func (c MTLComputePipelineDescriptor) SetRequiredThreadsPerThreadgroup(value MTLSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}

























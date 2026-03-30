// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4ComputePipelineDescriptor] class.
var (
	_MTL4ComputePipelineDescriptorClass     MTL4ComputePipelineDescriptorClass
	_MTL4ComputePipelineDescriptorClassOnce sync.Once
)

func getMTL4ComputePipelineDescriptorClass() MTL4ComputePipelineDescriptorClass {
	_MTL4ComputePipelineDescriptorClassOnce.Do(func() {
		_MTL4ComputePipelineDescriptorClass = MTL4ComputePipelineDescriptorClass{class: objc.GetClass("MTL4ComputePipelineDescriptor")}
	})
	return _MTL4ComputePipelineDescriptorClass
}

// GetMTL4ComputePipelineDescriptorClass returns the class object for MTL4ComputePipelineDescriptor.
func GetMTL4ComputePipelineDescriptorClass() MTL4ComputePipelineDescriptorClass {
	return getMTL4ComputePipelineDescriptorClass()
}

type MTL4ComputePipelineDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4ComputePipelineDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4ComputePipelineDescriptorClass) Alloc() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Describes a compute pipeline state.
//
// # Instance Properties
//
//   - [MTL4ComputePipelineDescriptor.ComputeFunctionDescriptor]: A descriptor representing the compute pipeline’s function.
//   - [MTL4ComputePipelineDescriptor.SetComputeFunctionDescriptor]
//   - [MTL4ComputePipelineDescriptor.MaxTotalThreadsPerThreadgroup]: The maximum total number of threads that Metal can execute in a single threadgroup for the compute function.
//   - [MTL4ComputePipelineDescriptor.SetMaxTotalThreadsPerThreadgroup]
//   - [MTL4ComputePipelineDescriptor.RequiredThreadsPerThreadgroup]: The required number of threads per threadgroup for compute dispatches.
//   - [MTL4ComputePipelineDescriptor.SetRequiredThreadsPerThreadgroup]
//   - [MTL4ComputePipelineDescriptor.StaticLinkingDescriptor]: An object that contains information about functions to link to the compute pipeline.
//   - [MTL4ComputePipelineDescriptor.SetStaticLinkingDescriptor]
//   - [MTL4ComputePipelineDescriptor.SupportBinaryLinking]: A boolean value indicating whether the compute pipeline supports linking binary functions.
//   - [MTL4ComputePipelineDescriptor.SetSupportBinaryLinking]
//   - [MTL4ComputePipelineDescriptor.SupportIndirectCommandBuffers]: A value indicating whether the pipeline supports Metal indirect command buffers.
//   - [MTL4ComputePipelineDescriptor.SetSupportIndirectCommandBuffers]
//   - [MTL4ComputePipelineDescriptor.ThreadGroupSizeIsMultipleOfThreadExecutionWidth]: A boolean value indicating whether each dimension of the threadgroup size is a multiple of its corresponding thread execution width.
//   - [MTL4ComputePipelineDescriptor.SetThreadGroupSizeIsMultipleOfThreadExecutionWidth]
//
// # Instance Methods
//
//   - [MTL4ComputePipelineDescriptor.Reset]: Resets the descriptor to its default values.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor
type MTL4ComputePipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4ComputePipelineDescriptorFromID constructs a [MTL4ComputePipelineDescriptor] from an objc.ID.
//
// Describes a compute pipeline state.
func MTL4ComputePipelineDescriptorFromID(id objc.ID) MTL4ComputePipelineDescriptor {
	return MTL4ComputePipelineDescriptor{MTL4PipelineDescriptor: MTL4PipelineDescriptorFromID(id)}
}

// NOTE: MTL4ComputePipelineDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4ComputePipelineDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4ComputePipelineDescriptor.ComputeFunctionDescriptor]: A descriptor representing the compute pipeline’s function.
//   - [IMTL4ComputePipelineDescriptor.SetComputeFunctionDescriptor]
//   - [IMTL4ComputePipelineDescriptor.MaxTotalThreadsPerThreadgroup]: The maximum total number of threads that Metal can execute in a single threadgroup for the compute function.
//   - [IMTL4ComputePipelineDescriptor.SetMaxTotalThreadsPerThreadgroup]
//   - [IMTL4ComputePipelineDescriptor.RequiredThreadsPerThreadgroup]: The required number of threads per threadgroup for compute dispatches.
//   - [IMTL4ComputePipelineDescriptor.SetRequiredThreadsPerThreadgroup]
//   - [IMTL4ComputePipelineDescriptor.StaticLinkingDescriptor]: An object that contains information about functions to link to the compute pipeline.
//   - [IMTL4ComputePipelineDescriptor.SetStaticLinkingDescriptor]
//   - [IMTL4ComputePipelineDescriptor.SupportBinaryLinking]: A boolean value indicating whether the compute pipeline supports linking binary functions.
//   - [IMTL4ComputePipelineDescriptor.SetSupportBinaryLinking]
//   - [IMTL4ComputePipelineDescriptor.SupportIndirectCommandBuffers]: A value indicating whether the pipeline supports Metal indirect command buffers.
//   - [IMTL4ComputePipelineDescriptor.SetSupportIndirectCommandBuffers]
//   - [IMTL4ComputePipelineDescriptor.ThreadGroupSizeIsMultipleOfThreadExecutionWidth]: A boolean value indicating whether each dimension of the threadgroup size is a multiple of its corresponding thread execution width.
//   - [IMTL4ComputePipelineDescriptor.SetThreadGroupSizeIsMultipleOfThreadExecutionWidth]
//
// # Instance Methods
//
//   - [IMTL4ComputePipelineDescriptor.Reset]: Resets the descriptor to its default values.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor
type IMTL4ComputePipelineDescriptor interface {
	IMTL4PipelineDescriptor

	// Topic: Instance Properties

	// A descriptor representing the compute pipeline’s function.
	ComputeFunctionDescriptor() IMTL4FunctionDescriptor
	SetComputeFunctionDescriptor(value IMTL4FunctionDescriptor)
	// The maximum total number of threads that Metal can execute in a single threadgroup for the compute function.
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	// The required number of threads per threadgroup for compute dispatches.
	RequiredThreadsPerThreadgroup() MTLSize
	SetRequiredThreadsPerThreadgroup(value MTLSize)
	// An object that contains information about functions to link to the compute pipeline.
	StaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	// A boolean value indicating whether the compute pipeline supports linking binary functions.
	SupportBinaryLinking() bool
	SetSupportBinaryLinking(value bool)
	// A value indicating whether the pipeline supports Metal indirect command buffers.
	SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState
	SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState)
	// A boolean value indicating whether each dimension of the threadgroup size is a multiple of its corresponding thread execution width.
	ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool)

	// Topic: Instance Methods

	// Resets the descriptor to its default values.
	Reset()
}

// Init initializes the instance.
func (m MTL4ComputePipelineDescriptor) Init() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4ComputePipelineDescriptor) Autorelease() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4ComputePipelineDescriptor creates a new MTL4ComputePipelineDescriptor instance.
func NewMTL4ComputePipelineDescriptor() MTL4ComputePipelineDescriptor {
	class := getMTL4ComputePipelineDescriptorClass()
	rv := objc.Send[MTL4ComputePipelineDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Resets the descriptor to its default values.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/reset()
func (m MTL4ComputePipelineDescriptor) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// A descriptor representing the compute pipeline’s function.
//
// # Discussion
//
// You don’t assign instances of [MTL4FunctionDescriptor] to this property
// directly, instead assign an instance of one of its subclasses, such as
// [MTL4LibraryFunctionDescriptor], which represents a function from a Metal
// library.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/computeFunctionDescriptor
func (m MTL4ComputePipelineDescriptor) ComputeFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("computeFunctionDescriptor"))
	return MTL4FunctionDescriptorFromID(objc.ID(rv))
}
func (m MTL4ComputePipelineDescriptor) SetComputeFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setComputeFunctionDescriptor:"), value)
}

// The maximum total number of threads that Metal can execute in a single
// threadgroup for the compute function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m MTL4ComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}
func (m MTL4ComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}

// The required number of threads per threadgroup for compute dispatches.
//
// # Discussion
//
// When you set this value, you are responsible for ensuring that the
// `threadsPerThreadgroup` argument of any compute dispatch matches it.
//
// Setting this property is optional, except in cases where the pipeline uses
// .
//
// This property’s default value is `0`, which disables its effect.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/requiredThreadsPerThreadgroup
func (m MTL4ComputePipelineDescriptor) RequiredThreadsPerThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](m.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return MTLSize(rv)
}
func (m MTL4ComputePipelineDescriptor) SetRequiredThreadsPerThreadgroup(value MTLSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}

// An object that contains information about functions to link to the compute
// pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/staticLinkingDescriptor
func (m MTL4ComputePipelineDescriptor) StaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("staticLinkingDescriptor"))
	return MTL4StaticLinkingDescriptorFromID(objc.ID(rv))
}
func (m MTL4ComputePipelineDescriptor) SetStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setStaticLinkingDescriptor:"), value)
}

// A boolean value indicating whether the compute pipeline supports linking
// binary functions.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/supportBinaryLinking
func (m MTL4ComputePipelineDescriptor) SupportBinaryLinking() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportBinaryLinking"))
	return rv
}
func (m MTL4ComputePipelineDescriptor) SetSupportBinaryLinking(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportBinaryLinking:"), value)
}

// A value indicating whether the pipeline supports Metal indirect command
// buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/supportIndirectCommandBuffers
func (m MTL4ComputePipelineDescriptor) SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState {
	rv := objc.Send[MTL4IndirectCommandBufferSupportState](m.ID, objc.Sel("supportIndirectCommandBuffers"))
	return MTL4IndirectCommandBufferSupportState(rv)
}
func (m MTL4ComputePipelineDescriptor) SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}

// A boolean value indicating whether each dimension of the threadgroup size
// is a multiple of its corresponding thread execution width.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/threadGroupSizeIsMultipleOfThreadExecutionWidth
func (m MTL4ComputePipelineDescriptor) ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("threadGroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}
func (m MTL4ComputePipelineDescriptor) SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setThreadGroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}

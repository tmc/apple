// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLCommandBufferDescriptor] class.
var (
	_MTLCommandBufferDescriptorClass     MTLCommandBufferDescriptorClass
	_MTLCommandBufferDescriptorClassOnce sync.Once
)

func getMTLCommandBufferDescriptorClass() MTLCommandBufferDescriptorClass {
	_MTLCommandBufferDescriptorClassOnce.Do(func() {
		_MTLCommandBufferDescriptorClass = MTLCommandBufferDescriptorClass{class: objc.GetClass("MTLCommandBufferDescriptor")}
	})
	return _MTLCommandBufferDescriptorClass
}

// GetMTLCommandBufferDescriptorClass returns the class object for MTLCommandBufferDescriptor.
func GetMTLCommandBufferDescriptorClass() MTLCommandBufferDescriptorClass {
	return getMTLCommandBufferDescriptorClass()
}

type MTLCommandBufferDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLCommandBufferDescriptorClass) Alloc() MTLCommandBufferDescriptor {
	rv := objc.Send[MTLCommandBufferDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A configuration that customizes the behavior for a new command buffer.
//
// # Overview
// 
// Create a command buffer with a custom configuration by creating an
// [MTLCommandBufferDescriptor] instance and passing it to an
// [MTLCommandQueue] instance’s [CommandBufferWithDescriptor] method. You
// can configure whether the command buffer retains references to resources
// that its commands refer to with the [MTLCommandBufferDescriptor.RetainedReferences] property. The
// command buffer can save extra error information, which is useful during
// development, by setting its [MTLCommandBufferDescriptor.ErrorOptions] property to
// [CommandBufferErrorOptionEncoderExecutionStatus].
//
// # Configuring the command buffer
//
//   - [MTLCommandBufferDescriptor.LogState]: The shader logging configuration that the command buffer uses.
//   - [MTLCommandBufferDescriptor.SetLogState]
//   - [MTLCommandBufferDescriptor.RetainedReferences]: A Boolean value that indicates whether the command buffer the descriptor creates maintains strong references to the resources it uses.
//   - [MTLCommandBufferDescriptor.SetRetainedReferences]
//   - [MTLCommandBufferDescriptor.ErrorOptions]: The reporting configuration that indicates which information the GPU driver stores in a command buffer’s error property.
//   - [MTLCommandBufferDescriptor.SetErrorOptions]
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor
type MTLCommandBufferDescriptor struct {
	objectivec.Object
}

// MTLCommandBufferDescriptorFromID constructs a [MTLCommandBufferDescriptor] from an objc.ID.
//
// A configuration that customizes the behavior for a new command buffer.
func MTLCommandBufferDescriptorFromID(id objc.ID) MTLCommandBufferDescriptor {
	return MTLCommandBufferDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLCommandBufferDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLCommandBufferDescriptor] class.
//
// # Configuring the command buffer
//
//   - [IMTLCommandBufferDescriptor.LogState]: The shader logging configuration that the command buffer uses.
//   - [IMTLCommandBufferDescriptor.SetLogState]
//   - [IMTLCommandBufferDescriptor.RetainedReferences]: A Boolean value that indicates whether the command buffer the descriptor creates maintains strong references to the resources it uses.
//   - [IMTLCommandBufferDescriptor.SetRetainedReferences]
//   - [IMTLCommandBufferDescriptor.ErrorOptions]: The reporting configuration that indicates which information the GPU driver stores in a command buffer’s error property.
//   - [IMTLCommandBufferDescriptor.SetErrorOptions]
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor
type IMTLCommandBufferDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the command buffer

	// The shader logging configuration that the command buffer uses.
	LogState() MTLLogState
	SetLogState(value MTLLogState)
	// A Boolean value that indicates whether the command buffer the descriptor creates maintains strong references to the resources it uses.
	RetainedReferences() bool
	SetRetainedReferences(value bool)
	// The reporting configuration that indicates which information the GPU driver stores in a command buffer’s error property.
	ErrorOptions() MTLCommandBufferErrorOption
	SetErrorOptions(value MTLCommandBufferErrorOption)

	// The domain for Metal command buffer errors.
	MTLCommandBufferErrorDomain() string
	// An option that instructs a command buffer to save additional details about a GPU runtime error.
	EncoderExecutionStatus() MTLCommandBufferErrorOption
	SetEncoderExecutionStatus(value MTLCommandBufferErrorOption)
}

// Init initializes the instance.
func (c MTLCommandBufferDescriptor) Init() MTLCommandBufferDescriptor {
	rv := objc.Send[MTLCommandBufferDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLCommandBufferDescriptor) Autorelease() MTLCommandBufferDescriptor {
	rv := objc.Send[MTLCommandBufferDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLCommandBufferDescriptor creates a new MTLCommandBufferDescriptor instance.
func NewMTLCommandBufferDescriptor() MTLCommandBufferDescriptor {
	class := getMTLCommandBufferDescriptorClass()
	rv := objc.Send[MTLCommandBufferDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The shader logging configuration that the command buffer uses.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor/logState
func (c MTLCommandBufferDescriptor) LogState() MTLLogState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("logState"))
	return MTLLogStateObjectFromID(rv)
}
func (c MTLCommandBufferDescriptor) SetLogState(value MTLLogState) {
	objc.Send[struct{}](c.ID, objc.Sel("setLogState:"), value)
}

// A Boolean value that indicates whether the command buffer the descriptor
// creates maintains strong references to the resources it uses.
//
// # Discussion
// 
// Set this property to [true] (its default) to create a command buffer that
// maintains strong references to resource instances that its commands need.
// Otherwise, set it to [false] to create a command buffer that doesn’t
// maintain strong references to its resources.
// 
// Apps typically create command buffers that don’t maintain references to
// resources for extremely performance-critical situations. Even though the
// runtime cost for retaining or releasing a single resource is trivial, the
// aggregate time savings may be worth it.
// 
// It’s your app’s responsibility to maintain strong references to all the
// resources the command buffer uses until it finishes running on the GPU.
// 
// You can determine whether an existing command buffer retains references by
// checking its [retainedReferences] property.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [retainedReferences]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/retainedReferences
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor/retainedReferences
func (c MTLCommandBufferDescriptor) RetainedReferences() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("retainedReferences"))
	return rv
}
func (c MTLCommandBufferDescriptor) SetRetainedReferences(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setRetainedReferences:"), value)
}

// The reporting configuration that indicates which information the GPU driver
// stores in a command buffer’s error property.
//
// # Discussion
// 
// By default, a GPU driver doesn’t report additional error information.
// 
// To create a command buffer that saves additional GPU runtime error
// information, add the [CommandBufferErrorOptionEncoderExecutionStatus]
// option to this property. If the GPU encounters an error as it runs the
// command buffer, you can retrieve the additional information from the
// command buffer’s [error] property.
//
// [error]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/error
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferDescriptor/errorOptions
func (c MTLCommandBufferDescriptor) ErrorOptions() MTLCommandBufferErrorOption {
	rv := objc.Send[MTLCommandBufferErrorOption](c.ID, objc.Sel("errorOptions"))
	return MTLCommandBufferErrorOption(rv)
}
func (c MTLCommandBufferDescriptor) SetErrorOptions(value MTLCommandBufferErrorOption) {
	objc.Send[struct{}](c.ID, objc.Sel("setErrorOptions:"), value)
}

// The domain for Metal command buffer errors.
//
// See: https://developer.apple.com/documentation/metal/mtlcommandbuffererrordomain
func (c MTLCommandBufferDescriptor) MTLCommandBufferErrorDomain() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("MTLCommandBufferErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// An option that instructs a command buffer to save additional details about
// a GPU runtime error.
//
// See: https://developer.apple.com/documentation/metal/mtlcommandbuffererroroption/encoderexecutionstatus
func (c MTLCommandBufferDescriptor) EncoderExecutionStatus() MTLCommandBufferErrorOption {
	rv := objc.Send[MTLCommandBufferErrorOption](c.ID, objc.Sel("MTLCommandBufferErrorOptionEncoderExecutionStatus"))
	return MTLCommandBufferErrorOption(rv)
}
func (c MTLCommandBufferDescriptor) SetEncoderExecutionStatus(value MTLCommandBufferErrorOption) {
	objc.Send[struct{}](c.ID, objc.Sel("setMTLCommandBufferErrorOptionEncoderExecutionStatus:"), value)
}


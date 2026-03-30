// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLComputePassDescriptor] class.
var (
	_MTLComputePassDescriptorClass     MTLComputePassDescriptorClass
	_MTLComputePassDescriptorClassOnce sync.Once
)

func getMTLComputePassDescriptorClass() MTLComputePassDescriptorClass {
	_MTLComputePassDescriptorClassOnce.Do(func() {
		_MTLComputePassDescriptorClass = MTLComputePassDescriptorClass{class: objc.GetClass("MTLComputePassDescriptor")}
	})
	return _MTLComputePassDescriptorClass
}

// GetMTLComputePassDescriptorClass returns the class object for MTLComputePassDescriptor.
func GetMTLComputePassDescriptorClass() MTLComputePassDescriptorClass {
	return getMTLComputePassDescriptorClass()
}

type MTLComputePassDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLComputePassDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLComputePassDescriptorClass) Alloc() MTLComputePassDescriptor {
	rv := objc.Send[MTLComputePassDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of how to dispatch execution of pass commands and GPU
// performance sampling.
//
// # Configuring the dispatch mechanism
//
//   - [MTLComputePassDescriptor.DispatchType]: The strategy for dispatching any compute commands encoded in the compute pass.
//   - [MTLComputePassDescriptor.SetDispatchType]
//
// # Specifying sample buffers for GPU counters
//
//   - [MTLComputePassDescriptor.SampleBufferAttachments]: The sample buffers that the compute pass can access.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassDescriptor
type MTLComputePassDescriptor struct {
	objectivec.Object
}

// MTLComputePassDescriptorFromID constructs a [MTLComputePassDescriptor] from an objc.ID.
//
// A description of how to dispatch execution of pass commands and GPU
// performance sampling.
func MTLComputePassDescriptorFromID(id objc.ID) MTLComputePassDescriptor {
	return MTLComputePassDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTLComputePassDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLComputePassDescriptor] class.
//
// # Configuring the dispatch mechanism
//
//   - [IMTLComputePassDescriptor.DispatchType]: The strategy for dispatching any compute commands encoded in the compute pass.
//   - [IMTLComputePassDescriptor.SetDispatchType]
//
// # Specifying sample buffers for GPU counters
//
//   - [IMTLComputePassDescriptor.SampleBufferAttachments]: The sample buffers that the compute pass can access.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassDescriptor
type IMTLComputePassDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the dispatch mechanism

	// The strategy for dispatching any compute commands encoded in the compute pass.
	DispatchType() MTLDispatchType
	SetDispatchType(value MTLDispatchType)

	// Topic: Specifying sample buffers for GPU counters

	// The sample buffers that the compute pass can access.
	SampleBufferAttachments() IMTLComputePassSampleBufferAttachmentDescriptorArray
}

// Init initializes the instance.
func (c MTLComputePassDescriptor) Init() MTLComputePassDescriptor {
	rv := objc.Send[MTLComputePassDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLComputePassDescriptor) Autorelease() MTLComputePassDescriptor {
	rv := objc.Send[MTLComputePassDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLComputePassDescriptor creates a new MTLComputePassDescriptor instance.
func NewMTLComputePassDescriptor() MTLComputePassDescriptor {
	class := getMTLComputePassDescriptorClass()
	rv := objc.Send[MTLComputePassDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The strategy for dispatching any compute commands encoded in the compute
// pass.
//
// # Discussion
//
// The default dispatch type is [MTLDispatchTypeSerial].
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassDescriptor/dispatchType
func (c MTLComputePassDescriptor) DispatchType() MTLDispatchType {
	rv := objc.Send[MTLDispatchType](c.ID, objc.Sel("dispatchType"))
	return MTLDispatchType(rv)
}
func (c MTLComputePassDescriptor) SetDispatchType(value MTLDispatchType) {
	objc.Send[struct{}](c.ID, objc.Sel("setDispatchType:"), value)
}

// The sample buffers that the compute pass can access.
//
// # Discussion
//
// The GPU uses sample buffers to record performance information. See [GPU
// counters and counter sample buffers], [Sampling GPU data into counter
// sample buffers], and [MTLCounter] for more information.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassDescriptor/sampleBufferAttachments
//
// [GPU counters and counter sample buffers]: https://developer.apple.com/documentation/Metal/gpu-counters-and-counter-sample-buffers
// [Sampling GPU data into counter sample buffers]: https://developer.apple.com/documentation/Metal/sampling-gpu-data-into-counter-sample-buffers
func (c MTLComputePassDescriptor) SampleBufferAttachments() IMTLComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sampleBufferAttachments"))
	return MTLComputePassSampleBufferAttachmentDescriptorArrayFromID(objc.ID(rv))
}

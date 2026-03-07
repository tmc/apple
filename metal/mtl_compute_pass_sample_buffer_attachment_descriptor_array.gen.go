// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLComputePassSampleBufferAttachmentDescriptorArray] class.
var (
	_MTLComputePassSampleBufferAttachmentDescriptorArrayClass     MTLComputePassSampleBufferAttachmentDescriptorArrayClass
	_MTLComputePassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTLComputePassSampleBufferAttachmentDescriptorArrayClass() MTLComputePassSampleBufferAttachmentDescriptorArrayClass {
	_MTLComputePassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		_MTLComputePassSampleBufferAttachmentDescriptorArrayClass = MTLComputePassSampleBufferAttachmentDescriptorArrayClass{class: objc.GetClass("MTLComputePassSampleBufferAttachmentDescriptorArray")}
	})
	return _MTLComputePassSampleBufferAttachmentDescriptorArrayClass
}

// GetMTLComputePassSampleBufferAttachmentDescriptorArrayClass returns the class object for MTLComputePassSampleBufferAttachmentDescriptorArray.
func GetMTLComputePassSampleBufferAttachmentDescriptorArrayClass() MTLComputePassSampleBufferAttachmentDescriptorArrayClass {
	return getMTLComputePassSampleBufferAttachmentDescriptorArrayClass()
}

type MTLComputePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLComputePassSampleBufferAttachmentDescriptorArrayClass) Alloc() MTLComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLComputePassSampleBufferAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A container that stores an array of sample buffer attachments for a compute
// pass.
//
// # Overview
// 
// The number of elements in the array is at least the number of elements in
// an [MTLDevice] instance’s [counterSets] property.
//
// [counterSets]: https://developer.apple.com/documentation/Metal/MTLDevice/counterSets
//
// # Accessing a sample buffer attachment
//
//   - [MTLComputePassSampleBufferAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the descriptor object for the specified sample buffer attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptorArray
type MTLComputePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTLComputePassSampleBufferAttachmentDescriptorArrayFromID constructs a [MTLComputePassSampleBufferAttachmentDescriptorArray] from an objc.ID.
//
// A container that stores an array of sample buffer attachments for a compute
// pass.
func MTLComputePassSampleBufferAttachmentDescriptorArrayFromID(id objc.ID) MTLComputePassSampleBufferAttachmentDescriptorArray {
	return MTLComputePassSampleBufferAttachmentDescriptorArray{objectivec.Object{id}}
}
// NOTE: MTLComputePassSampleBufferAttachmentDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLComputePassSampleBufferAttachmentDescriptorArray] class.
//
// # Accessing a sample buffer attachment
//
//   - [IMTLComputePassSampleBufferAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the descriptor object for the specified sample buffer attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptorArray
type IMTLComputePassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject

	// Topic: Accessing a sample buffer attachment

	// Returns the descriptor object for the specified sample buffer attachment.
	ObjectAtIndexedSubscript(attachmentIndex uint) IMTLComputePassSampleBufferAttachmentDescriptor

	// The counter sets supported by the device object.
	CounterSets() MTLCounterSet
	SetCounterSets(value MTLCounterSet)
	// Sets the descriptor object for the specified sample buffer attachment.
	SetObjectAtIndexedSubscript(attachment IMTLComputePassSampleBufferAttachmentDescriptor, attachmentIndex uint)
}





// Init initializes the instance.
func (c MTLComputePassSampleBufferAttachmentDescriptorArray) Init() MTLComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLComputePassSampleBufferAttachmentDescriptorArray](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLComputePassSampleBufferAttachmentDescriptorArray) Autorelease() MTLComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLComputePassSampleBufferAttachmentDescriptorArray](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLComputePassSampleBufferAttachmentDescriptorArray creates a new MTLComputePassSampleBufferAttachmentDescriptorArray instance.
func NewMTLComputePassSampleBufferAttachmentDescriptorArray() MTLComputePassSampleBufferAttachmentDescriptorArray {
	class := getMTLComputePassSampleBufferAttachmentDescriptorArrayClass()
	rv := objc.Send[MTLComputePassSampleBufferAttachmentDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the descriptor object for the specified sample buffer attachment.
//
// attachmentIndex: An index for the sample buffer attachment to fetch.
//
// # Return Value
// 
// The requested [MTLComputePassSampleBufferAttachmentDescriptor] object.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (c MTLComputePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IMTLComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return MTLComputePassSampleBufferAttachmentDescriptorFromID(rv)
}

// Sets the descriptor object for the specified sample buffer attachment.
//
// attachment: A sample buffer attachment descriptor. When set to `nil`, removes any
// existing buffer attachment descriptor at `attachmentIndex`.
//
// attachmentIndex: The attachment in the array to replace.
//
// # Discussion
// 
// The method copies the `attachment` parameter’s contents into the
// attachment at the specified index.
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (c MTLComputePassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLComputePassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}












// The counter sets supported by the device object.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/countersets
func (c MTLComputePassSampleBufferAttachmentDescriptorArray) CounterSets() MTLCounterSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("counterSets"))
	return MTLCounterSetObjectFromID(rv)
}
func (c MTLComputePassSampleBufferAttachmentDescriptorArray) SetCounterSets(value MTLCounterSet) {
	objc.Send[struct{}](c.ID, objc.Sel("setCounterSets:"), value)
}
























// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray] class.
var (
	_MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass     MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass
	_MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass {
	_MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		_MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass = MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass{class: objc.GetClass("MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray")}
	})
	return _MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass
}

// GetMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass returns the class object for MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray.
func GetMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass {
	return getMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass()
}

type MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass) Alloc() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Subscripts
//
//   - [MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray.ObjectAtIndexedSubscript]
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray
type MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayFromID constructs a [MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray] from an objc.ID.
func MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayFromID(id objc.ID) MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	return MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray{objectivec.Object{ID: id}}
}
// NOTE: MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray] class.
//
// # Subscripts
//
//   - [IMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray.ObjectAtIndexedSubscript]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray
type IMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject

	// Topic: Subscripts

	ObjectAtIndexedSubscript(attachmentIndex uint) IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor

	SetObjectAtIndexedSubscript(attachment IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor, attachmentIndex uint)
}

// Init initializes the instance.
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray) Init() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray) Autorelease() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray creates a new MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray instance.
func NewMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	class := getMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass()
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return MTLAccelerationStructurePassSampleBufferAttachmentDescriptorFromID(rv)
}
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](a.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}


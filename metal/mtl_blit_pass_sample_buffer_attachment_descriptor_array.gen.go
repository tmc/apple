// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLBlitPassSampleBufferAttachmentDescriptorArray] class.
var (
	_MTLBlitPassSampleBufferAttachmentDescriptorArrayClass     MTLBlitPassSampleBufferAttachmentDescriptorArrayClass
	_MTLBlitPassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTLBlitPassSampleBufferAttachmentDescriptorArrayClass() MTLBlitPassSampleBufferAttachmentDescriptorArrayClass {
	_MTLBlitPassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		_MTLBlitPassSampleBufferAttachmentDescriptorArrayClass = MTLBlitPassSampleBufferAttachmentDescriptorArrayClass{class: objc.GetClass("MTLBlitPassSampleBufferAttachmentDescriptorArray")}
	})
	return _MTLBlitPassSampleBufferAttachmentDescriptorArrayClass
}

// GetMTLBlitPassSampleBufferAttachmentDescriptorArrayClass returns the class object for MTLBlitPassSampleBufferAttachmentDescriptorArray.
func GetMTLBlitPassSampleBufferAttachmentDescriptorArrayClass() MTLBlitPassSampleBufferAttachmentDescriptorArrayClass {
	return getMTLBlitPassSampleBufferAttachmentDescriptorArrayClass()
}

type MTLBlitPassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLBlitPassSampleBufferAttachmentDescriptorArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLBlitPassSampleBufferAttachmentDescriptorArrayClass) Alloc() MTLBlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLBlitPassSampleBufferAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A container that stores an array of sample buffer attachments for a blit
// pass.
//
// # Overview
// 
// The number of elements in the array is at least the number of elements in
// an [MTLDevice] instance’s [counterSets] property.
//
// [counterSets]: https://developer.apple.com/documentation/Metal/MTLDevice/counterSets
//
// # Accessing a sample buffer attachment descriptor
//
//   - [MTLBlitPassSampleBufferAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Accesses one of the array’s blit pass sample buffer attachment descriptor instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptorArray
type MTLBlitPassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTLBlitPassSampleBufferAttachmentDescriptorArrayFromID constructs a [MTLBlitPassSampleBufferAttachmentDescriptorArray] from an objc.ID.
//
// A container that stores an array of sample buffer attachments for a blit
// pass.
func MTLBlitPassSampleBufferAttachmentDescriptorArrayFromID(id objc.ID) MTLBlitPassSampleBufferAttachmentDescriptorArray {
	return MTLBlitPassSampleBufferAttachmentDescriptorArray{objectivec.Object{ID: id}}
}
// NOTE: MTLBlitPassSampleBufferAttachmentDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLBlitPassSampleBufferAttachmentDescriptorArray] class.
//
// # Accessing a sample buffer attachment descriptor
//
//   - [IMTLBlitPassSampleBufferAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Accesses one of the array’s blit pass sample buffer attachment descriptor instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptorArray
type IMTLBlitPassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject

	// Topic: Accessing a sample buffer attachment descriptor

	// Accesses one of the array’s blit pass sample buffer attachment descriptor instances.
	ObjectAtIndexedSubscript(attachmentIndex uint) IMTLBlitPassSampleBufferAttachmentDescriptor

	// The counter sets supported by the device object.
	CounterSets() MTLCounterSet
	SetCounterSets(value MTLCounterSet)
	// Copies the properties of a blit pass sample buffer attachment descriptor instance to the properties of one of the array’s instances.
	SetObjectAtIndexedSubscript(attachment IMTLBlitPassSampleBufferAttachmentDescriptor, attachmentIndex uint)
}

// Init initializes the instance.
func (b MTLBlitPassSampleBufferAttachmentDescriptorArray) Init() MTLBlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLBlitPassSampleBufferAttachmentDescriptorArray](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MTLBlitPassSampleBufferAttachmentDescriptorArray) Autorelease() MTLBlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLBlitPassSampleBufferAttachmentDescriptorArray](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLBlitPassSampleBufferAttachmentDescriptorArray creates a new MTLBlitPassSampleBufferAttachmentDescriptorArray instance.
func NewMTLBlitPassSampleBufferAttachmentDescriptorArray() MTLBlitPassSampleBufferAttachmentDescriptorArray {
	class := getMTLBlitPassSampleBufferAttachmentDescriptorArrayClass()
	rv := objc.Send[MTLBlitPassSampleBufferAttachmentDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Accesses one of the array’s blit pass sample buffer attachment descriptor
// instances.
//
// attachmentIndex: An index of one of the array’s
// [MTLBlitPassSampleBufferAttachmentDescriptor] instances.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (b MTLBlitPassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IMTLBlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return MTLBlitPassSampleBufferAttachmentDescriptorFromID(rv)
}
// Copies the properties of a blit pass sample buffer attachment descriptor
// instance to the properties of one of the array’s instances.
//
// attachment: An [MTLBlitPassSampleBufferAttachmentDescriptor] instance that the method
// assigns its properties values to the properties of the array’s instance
// at `attachmentIndex`.
// 
// You can reset the property configuration of the array’s instance at
// `attachmentIndex` to its default values by passing `nil`.
//
// attachmentIndex: An index into the array’s copies of attachment descriptor instances.
//
// # Discussion
// 
// The array has at
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (b MTLBlitPassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLBlitPassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](b.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}

// The counter sets supported by the device object.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/countersets
func (b MTLBlitPassSampleBufferAttachmentDescriptorArray) CounterSets() MTLCounterSet {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("counterSets"))
	return MTLCounterSetObjectFromID(rv)
}
func (b MTLBlitPassSampleBufferAttachmentDescriptorArray) SetCounterSets(value MTLCounterSet) {
	objc.Send[struct{}](b.ID, objc.Sel("setCounterSets:"), value)
}


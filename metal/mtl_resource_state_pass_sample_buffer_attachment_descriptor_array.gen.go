// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLResourceStatePassSampleBufferAttachmentDescriptorArray] class.
var (
	_MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass     MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass
	_MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass() MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass {
	_MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		_MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass = MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass{class: objc.GetClass("MTLResourceStatePassSampleBufferAttachmentDescriptorArray")}
	})
	return _MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass
}

// GetMTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass returns the class object for MTLResourceStatePassSampleBufferAttachmentDescriptorArray.
func GetMTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass() MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass {
	return getMTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass()
}

type MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass) Alloc() MTLResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLResourceStatePassSampleBufferAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An array of sample buffer attachments for a resource state pass.
//
// # Accessing a sample buffer attachment
//
//   - [MTLResourceStatePassSampleBufferAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the descriptor object for the specified sample buffer attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptorArray
type MTLResourceStatePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTLResourceStatePassSampleBufferAttachmentDescriptorArrayFromID constructs a [MTLResourceStatePassSampleBufferAttachmentDescriptorArray] from an objc.ID.
//
// An array of sample buffer attachments for a resource state pass.
func MTLResourceStatePassSampleBufferAttachmentDescriptorArrayFromID(id objc.ID) MTLResourceStatePassSampleBufferAttachmentDescriptorArray {
	return MTLResourceStatePassSampleBufferAttachmentDescriptorArray{objectivec.Object{ID: id}}
}

// NOTE: MTLResourceStatePassSampleBufferAttachmentDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLResourceStatePassSampleBufferAttachmentDescriptorArray] class.
//
// # Accessing a sample buffer attachment
//
//   - [IMTLResourceStatePassSampleBufferAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the descriptor object for the specified sample buffer attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptorArray
type IMTLResourceStatePassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject

	// Topic: Accessing a sample buffer attachment

	// Returns the descriptor object for the specified sample buffer attachment.
	ObjectAtIndexedSubscript(attachmentIndex uint) IMTLResourceStatePassSampleBufferAttachmentDescriptor

	// Sets the descriptor object for the specified sample buffer attachment.
	SetObjectAtIndexedSubscript(attachment IMTLResourceStatePassSampleBufferAttachmentDescriptor, attachmentIndex uint)
}

// Init initializes the instance.
func (r MTLResourceStatePassSampleBufferAttachmentDescriptorArray) Init() MTLResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLResourceStatePassSampleBufferAttachmentDescriptorArray](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLResourceStatePassSampleBufferAttachmentDescriptorArray) Autorelease() MTLResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLResourceStatePassSampleBufferAttachmentDescriptorArray](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLResourceStatePassSampleBufferAttachmentDescriptorArray creates a new MTLResourceStatePassSampleBufferAttachmentDescriptorArray instance.
func NewMTLResourceStatePassSampleBufferAttachmentDescriptorArray() MTLResourceStatePassSampleBufferAttachmentDescriptorArray {
	class := getMTLResourceStatePassSampleBufferAttachmentDescriptorArrayClass()
	rv := objc.Send[MTLResourceStatePassSampleBufferAttachmentDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the descriptor object for the specified sample buffer attachment.
//
// attachmentIndex: An index for the object to fetch.
//
// # Return Value
//
// The requested [MTLResourceStatePassSampleBufferAttachmentDescriptor]
// object.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (r MTLResourceStatePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IMTLResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return MTLResourceStatePassSampleBufferAttachmentDescriptorFromID(rv)
}

// Sets the descriptor object for the specified sample buffer attachment.
//
// attachment: A sample buffer attachment descriptor. If you specify `nil`, the attachment
// descriptor resets its configuration to default values.
//
// attachmentIndex: The item in the array to replace.
//
// # Discussion
//
// The method copies the `attachment` parameter’s contents into the
// attachment at the specified index..
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (r MTLResourceStatePassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLResourceStatePassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](r.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}

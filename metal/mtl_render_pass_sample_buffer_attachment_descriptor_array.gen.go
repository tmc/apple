// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRenderPassSampleBufferAttachmentDescriptorArray] class.
var (
	_MTLRenderPassSampleBufferAttachmentDescriptorArrayClass     MTLRenderPassSampleBufferAttachmentDescriptorArrayClass
	_MTLRenderPassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTLRenderPassSampleBufferAttachmentDescriptorArrayClass() MTLRenderPassSampleBufferAttachmentDescriptorArrayClass {
	_MTLRenderPassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		_MTLRenderPassSampleBufferAttachmentDescriptorArrayClass = MTLRenderPassSampleBufferAttachmentDescriptorArrayClass{class: objc.GetClass("MTLRenderPassSampleBufferAttachmentDescriptorArray")}
	})
	return _MTLRenderPassSampleBufferAttachmentDescriptorArrayClass
}

// GetMTLRenderPassSampleBufferAttachmentDescriptorArrayClass returns the class object for MTLRenderPassSampleBufferAttachmentDescriptorArray.
func GetMTLRenderPassSampleBufferAttachmentDescriptorArrayClass() MTLRenderPassSampleBufferAttachmentDescriptorArrayClass {
	return getMTLRenderPassSampleBufferAttachmentDescriptorArrayClass()
}

type MTLRenderPassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRenderPassSampleBufferAttachmentDescriptorArrayClass) Alloc() MTLRenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLRenderPassSampleBufferAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// An array of sample buffer attachments for a render pass.
//
// # Accessing a sample buffer attachment
//
//   - [MTLRenderPassSampleBufferAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the descriptor object for the specified sample buffer attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptorArray
type MTLRenderPassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTLRenderPassSampleBufferAttachmentDescriptorArrayFromID constructs a [MTLRenderPassSampleBufferAttachmentDescriptorArray] from an objc.ID.
//
// An array of sample buffer attachments for a render pass.
func MTLRenderPassSampleBufferAttachmentDescriptorArrayFromID(id objc.ID) MTLRenderPassSampleBufferAttachmentDescriptorArray {
	return MTLRenderPassSampleBufferAttachmentDescriptorArray{objectivec.Object{id}}
}
// NOTE: MTLRenderPassSampleBufferAttachmentDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLRenderPassSampleBufferAttachmentDescriptorArray] class.
//
// # Accessing a sample buffer attachment
//
//   - [IMTLRenderPassSampleBufferAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the descriptor object for the specified sample buffer attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptorArray
type IMTLRenderPassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject

	// Topic: Accessing a sample buffer attachment

	// Returns the descriptor object for the specified sample buffer attachment.
	ObjectAtIndexedSubscript(attachmentIndex uint) IMTLRenderPassSampleBufferAttachmentDescriptor

	// Sets the descriptor object for the specified sample buffer attachment.
	SetObjectAtIndexedSubscript(attachment IMTLRenderPassSampleBufferAttachmentDescriptor, attachmentIndex uint)
}





// Init initializes the instance.
func (r MTLRenderPassSampleBufferAttachmentDescriptorArray) Init() MTLRenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLRenderPassSampleBufferAttachmentDescriptorArray](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRenderPassSampleBufferAttachmentDescriptorArray) Autorelease() MTLRenderPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLRenderPassSampleBufferAttachmentDescriptorArray](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRenderPassSampleBufferAttachmentDescriptorArray creates a new MTLRenderPassSampleBufferAttachmentDescriptorArray instance.
func NewMTLRenderPassSampleBufferAttachmentDescriptorArray() MTLRenderPassSampleBufferAttachmentDescriptorArray {
	class := getMTLRenderPassSampleBufferAttachmentDescriptorArrayClass()
	rv := objc.Send[MTLRenderPassSampleBufferAttachmentDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the descriptor object for the specified sample buffer attachment.
//
// attachmentIndex: An index for the object to fetch.
//
// # Return Value
// 
// The [MTLRenderPassSampleBufferAttachmentDescriptor] at the specified index.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (r MTLRenderPassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IMTLRenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return MTLRenderPassSampleBufferAttachmentDescriptorFromID(rv)
}

// Sets the descriptor object for the specified sample buffer attachment.
//
// attachment: A sample buffer attachment descriptor. Specify `nil` to resets the
// attachment to default values.
//
// attachmentIndex: The item in the array to replace.
//
// # Discussion
// 
// The method copies the parameter contents into the attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (r MTLRenderPassSampleBufferAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLRenderPassSampleBufferAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](r.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}

































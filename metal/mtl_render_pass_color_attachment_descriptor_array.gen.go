// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRenderPassColorAttachmentDescriptorArray] class.
var (
	_MTLRenderPassColorAttachmentDescriptorArrayClass     MTLRenderPassColorAttachmentDescriptorArrayClass
	_MTLRenderPassColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTLRenderPassColorAttachmentDescriptorArrayClass() MTLRenderPassColorAttachmentDescriptorArrayClass {
	_MTLRenderPassColorAttachmentDescriptorArrayClassOnce.Do(func() {
		_MTLRenderPassColorAttachmentDescriptorArrayClass = MTLRenderPassColorAttachmentDescriptorArrayClass{class: objc.GetClass("MTLRenderPassColorAttachmentDescriptorArray")}
	})
	return _MTLRenderPassColorAttachmentDescriptorArrayClass
}

// GetMTLRenderPassColorAttachmentDescriptorArrayClass returns the class object for MTLRenderPassColorAttachmentDescriptorArray.
func GetMTLRenderPassColorAttachmentDescriptorArrayClass() MTLRenderPassColorAttachmentDescriptorArrayClass {
	return getMTLRenderPassColorAttachmentDescriptorArrayClass()
}

type MTLRenderPassColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRenderPassColorAttachmentDescriptorArrayClass) Alloc() MTLRenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[MTLRenderPassColorAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// An array of render pass color attachment descriptor objects.
//
// # Accessing the description of a color attachment
//
//   - [MTLRenderPassColorAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the descriptor object for the specified color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassColorAttachmentDescriptorArray
type MTLRenderPassColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTLRenderPassColorAttachmentDescriptorArrayFromID constructs a [MTLRenderPassColorAttachmentDescriptorArray] from an objc.ID.
//
// An array of render pass color attachment descriptor objects.
func MTLRenderPassColorAttachmentDescriptorArrayFromID(id objc.ID) MTLRenderPassColorAttachmentDescriptorArray {
	return MTLRenderPassColorAttachmentDescriptorArray{objectivec.Object{id}}
}
// NOTE: MTLRenderPassColorAttachmentDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLRenderPassColorAttachmentDescriptorArray] class.
//
// # Accessing the description of a color attachment
//
//   - [IMTLRenderPassColorAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the descriptor object for the specified color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassColorAttachmentDescriptorArray
type IMTLRenderPassColorAttachmentDescriptorArray interface {
	objectivec.IObject

	// Topic: Accessing the description of a color attachment

	// Returns the descriptor object for the specified color attachment.
	ObjectAtIndexedSubscript(attachmentIndex uint) IMTLRenderPassColorAttachmentDescriptor

	// Sets the descriptor for the specified color attachment.
	SetObjectAtIndexedSubscript(attachment IMTLRenderPassColorAttachmentDescriptor, attachmentIndex uint)
}





// Init initializes the instance.
func (r MTLRenderPassColorAttachmentDescriptorArray) Init() MTLRenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[MTLRenderPassColorAttachmentDescriptorArray](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRenderPassColorAttachmentDescriptorArray) Autorelease() MTLRenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[MTLRenderPassColorAttachmentDescriptorArray](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRenderPassColorAttachmentDescriptorArray creates a new MTLRenderPassColorAttachmentDescriptorArray instance.
func NewMTLRenderPassColorAttachmentDescriptorArray() MTLRenderPassColorAttachmentDescriptorArray {
	class := getMTLRenderPassColorAttachmentDescriptorArrayClass()
	rv := objc.Send[MTLRenderPassColorAttachmentDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the descriptor object for the specified color attachment.
//
// attachmentIndex: An index in the color attachment array.
//
// # Return Value
// 
// A descriptor object that contains color attachment information.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassColorAttachmentDescriptorArray/subscript(_:)
func (r MTLRenderPassColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IMTLRenderPassColorAttachmentDescriptor {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return MTLRenderPassColorAttachmentDescriptorFromID(rv)
}

// Sets the descriptor for the specified color attachment.
//
// attachment: A descriptor that contains color attachment information. Specify `nil` to
// reset the attachment to its default values.
//
// attachmentIndex: An index in the color attachment array.
//
// # Discussion
// 
// This method copies the color attachment information from the descriptor
// into the specified attachment in the array. Because the method copies the
// information, you can modify and reuse the descriptor without affecting a
// previously set attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (r MTLRenderPassColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLRenderPassColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](r.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}

































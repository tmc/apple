// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRenderPipelineColorAttachmentDescriptorArray] class.
var (
	_MTLRenderPipelineColorAttachmentDescriptorArrayClass     MTLRenderPipelineColorAttachmentDescriptorArrayClass
	_MTLRenderPipelineColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTLRenderPipelineColorAttachmentDescriptorArrayClass() MTLRenderPipelineColorAttachmentDescriptorArrayClass {
	_MTLRenderPipelineColorAttachmentDescriptorArrayClassOnce.Do(func() {
		_MTLRenderPipelineColorAttachmentDescriptorArrayClass = MTLRenderPipelineColorAttachmentDescriptorArrayClass{class: objc.GetClass("MTLRenderPipelineColorAttachmentDescriptorArray")}
	})
	return _MTLRenderPipelineColorAttachmentDescriptorArrayClass
}

// GetMTLRenderPipelineColorAttachmentDescriptorArrayClass returns the class object for MTLRenderPipelineColorAttachmentDescriptorArray.
func GetMTLRenderPipelineColorAttachmentDescriptorArrayClass() MTLRenderPipelineColorAttachmentDescriptorArrayClass {
	return getMTLRenderPipelineColorAttachmentDescriptorArrayClass()
}

type MTLRenderPipelineColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLRenderPipelineColorAttachmentDescriptorArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRenderPipelineColorAttachmentDescriptorArrayClass) Alloc() MTLRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTLRenderPipelineColorAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An array of render pipeline color attachment descriptor objects.
//
// # Accessing render pipeline state for a color attachment
//
//   - [MTLRenderPipelineColorAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the render pipeline state for the specified color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptorArray
type MTLRenderPipelineColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTLRenderPipelineColorAttachmentDescriptorArrayFromID constructs a [MTLRenderPipelineColorAttachmentDescriptorArray] from an objc.ID.
//
// An array of render pipeline color attachment descriptor objects.
func MTLRenderPipelineColorAttachmentDescriptorArrayFromID(id objc.ID) MTLRenderPipelineColorAttachmentDescriptorArray {
	return MTLRenderPipelineColorAttachmentDescriptorArray{objectivec.Object{ID: id}}
}
// NOTE: MTLRenderPipelineColorAttachmentDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLRenderPipelineColorAttachmentDescriptorArray] class.
//
// # Accessing render pipeline state for a color attachment
//
//   - [IMTLRenderPipelineColorAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the render pipeline state for the specified color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptorArray
type IMTLRenderPipelineColorAttachmentDescriptorArray interface {
	objectivec.IObject

	// Topic: Accessing render pipeline state for a color attachment

	// Returns the render pipeline state for the specified color attachment.
	ObjectAtIndexedSubscript(attachmentIndex uint) IMTLRenderPipelineColorAttachmentDescriptor

	// Sets the render pipeline state for a specified color attachment.
	SetObjectAtIndexedSubscript(attachment IMTLRenderPipelineColorAttachmentDescriptor, attachmentIndex uint)
}

// Init initializes the instance.
func (r MTLRenderPipelineColorAttachmentDescriptorArray) Init() MTLRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTLRenderPipelineColorAttachmentDescriptorArray](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRenderPipelineColorAttachmentDescriptorArray) Autorelease() MTLRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTLRenderPipelineColorAttachmentDescriptorArray](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRenderPipelineColorAttachmentDescriptorArray creates a new MTLRenderPipelineColorAttachmentDescriptorArray instance.
func NewMTLRenderPipelineColorAttachmentDescriptorArray() MTLRenderPipelineColorAttachmentDescriptorArray {
	class := getMTLRenderPipelineColorAttachmentDescriptorArrayClass()
	rv := objc.Send[MTLRenderPipelineColorAttachmentDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the render pipeline state for the specified color attachment.
//
// attachmentIndex: An index in the color attachment array.
//
// # Return Value
// 
// An [MTLRenderPipelineColorAttachmentDescriptor] instance that describes the
// render pipeline information for a color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptorArray/subscript(_:)
func (r MTLRenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IMTLRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return MTLRenderPipelineColorAttachmentDescriptorFromID(rv)
}
// Sets the render pipeline state for a specified color attachment.
//
// attachment: A descriptor that contains the render pipeline description for a color
// attachment.
//
// attachmentIndex: An index in the color attachment array.
//
// # Discussion
// 
// This method copies the pipeline state from the descriptor into the
// specified attachment in the array. The descriptor passed into this method
// can be modified and reused without affecting a previously set attachment.
// 
// If this method is called with `nil` for `attachment` for any legal index,
// its attachment descriptor state is set to the default values.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (r MTLRenderPipelineColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLRenderPipelineColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](r.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}


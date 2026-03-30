// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4RenderPipelineColorAttachmentDescriptorArray] class.
var (
	_MTL4RenderPipelineColorAttachmentDescriptorArrayClass     MTL4RenderPipelineColorAttachmentDescriptorArrayClass
	_MTL4RenderPipelineColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTL4RenderPipelineColorAttachmentDescriptorArrayClass() MTL4RenderPipelineColorAttachmentDescriptorArrayClass {
	_MTL4RenderPipelineColorAttachmentDescriptorArrayClassOnce.Do(func() {
		_MTL4RenderPipelineColorAttachmentDescriptorArrayClass = MTL4RenderPipelineColorAttachmentDescriptorArrayClass{class: objc.GetClass("MTL4RenderPipelineColorAttachmentDescriptorArray")}
	})
	return _MTL4RenderPipelineColorAttachmentDescriptorArrayClass
}

// GetMTL4RenderPipelineColorAttachmentDescriptorArrayClass returns the class object for MTL4RenderPipelineColorAttachmentDescriptorArray.
func GetMTL4RenderPipelineColorAttachmentDescriptorArrayClass() MTL4RenderPipelineColorAttachmentDescriptorArrayClass {
	return getMTL4RenderPipelineColorAttachmentDescriptorArrayClass()
}

type MTL4RenderPipelineColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4RenderPipelineColorAttachmentDescriptorArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4RenderPipelineColorAttachmentDescriptorArrayClass) Alloc() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An array of color attachment descriptions for a render pipeline.
//
// # Instance Methods
//
//   - [MTL4RenderPipelineColorAttachmentDescriptorArray.Reset]: Resets the elements of the descriptor array
//
// # Subscripts
//
//   - [MTL4RenderPipelineColorAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Accesses a color attachment at a specific index.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray
type MTL4RenderPipelineColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTL4RenderPipelineColorAttachmentDescriptorArrayFromID constructs a [MTL4RenderPipelineColorAttachmentDescriptorArray] from an objc.ID.
//
// An array of color attachment descriptions for a render pipeline.
func MTL4RenderPipelineColorAttachmentDescriptorArrayFromID(id objc.ID) MTL4RenderPipelineColorAttachmentDescriptorArray {
	return MTL4RenderPipelineColorAttachmentDescriptorArray{objectivec.Object{ID: id}}
}

// NOTE: MTL4RenderPipelineColorAttachmentDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4RenderPipelineColorAttachmentDescriptorArray] class.
//
// # Instance Methods
//
//   - [IMTL4RenderPipelineColorAttachmentDescriptorArray.Reset]: Resets the elements of the descriptor array
//
// # Subscripts
//
//   - [IMTL4RenderPipelineColorAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Accesses a color attachment at a specific index.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray
type IMTL4RenderPipelineColorAttachmentDescriptorArray interface {
	objectivec.IObject

	// Topic: Instance Methods

	// Resets the elements of the descriptor array
	Reset()

	// Topic: Subscripts

	// Accesses a color attachment at a specific index.
	ObjectAtIndexedSubscript(attachmentIndex uint) IMTL4RenderPipelineColorAttachmentDescriptor

	// Sets an attachment at an index.
	SetObjectAtIndexedSubscript(attachment IMTL4RenderPipelineColorAttachmentDescriptor, attachmentIndex uint)
}

// Init initializes the instance.
func (m MTL4RenderPipelineColorAttachmentDescriptorArray) Init() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4RenderPipelineColorAttachmentDescriptorArray) Autorelease() MTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineColorAttachmentDescriptorArray creates a new MTL4RenderPipelineColorAttachmentDescriptorArray instance.
func NewMTL4RenderPipelineColorAttachmentDescriptorArray() MTL4RenderPipelineColorAttachmentDescriptorArray {
	class := getMTL4RenderPipelineColorAttachmentDescriptorArrayClass()
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Resets the elements of the descriptor array
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray/reset()
func (m MTL4RenderPipelineColorAttachmentDescriptorArray) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// Accesses a color attachment at a specific index.
//
// attachmentIndex: Index of the attachment to access.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray/subscript(_:)
func (m MTL4RenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IMTL4RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return MTL4RenderPipelineColorAttachmentDescriptorFromID(rv)
}

// Sets an attachment at an index.
//
// attachment: The descriptor of the attachment to set.
//
// attachmentIndex: The index of the attachment within the array.
//
// # Discussion
//
// This function offers ‘copy’ semantics.
//
// You can safely set the color attachment at any legal index to nil. This has
// the effect of resetting that attachment descriptor’s state to its default
// values.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (m MTL4RenderPipelineColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTL4RenderPipelineColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}

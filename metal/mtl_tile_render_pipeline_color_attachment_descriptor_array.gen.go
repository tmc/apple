// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTileRenderPipelineColorAttachmentDescriptorArray] class.
var (
	_MTLTileRenderPipelineColorAttachmentDescriptorArrayClass     MTLTileRenderPipelineColorAttachmentDescriptorArrayClass
	_MTLTileRenderPipelineColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getMTLTileRenderPipelineColorAttachmentDescriptorArrayClass() MTLTileRenderPipelineColorAttachmentDescriptorArrayClass {
	_MTLTileRenderPipelineColorAttachmentDescriptorArrayClassOnce.Do(func() {
		_MTLTileRenderPipelineColorAttachmentDescriptorArrayClass = MTLTileRenderPipelineColorAttachmentDescriptorArrayClass{class: objc.GetClass("MTLTileRenderPipelineColorAttachmentDescriptorArray")}
	})
	return _MTLTileRenderPipelineColorAttachmentDescriptorArrayClass
}

// GetMTLTileRenderPipelineColorAttachmentDescriptorArrayClass returns the class object for MTLTileRenderPipelineColorAttachmentDescriptorArray.
func GetMTLTileRenderPipelineColorAttachmentDescriptorArrayClass() MTLTileRenderPipelineColorAttachmentDescriptorArrayClass {
	return getMTLTileRenderPipelineColorAttachmentDescriptorArrayClass()
}

type MTLTileRenderPipelineColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTileRenderPipelineColorAttachmentDescriptorArrayClass) Alloc() MTLTileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTLTileRenderPipelineColorAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An array of color attachment descriptors for the tile render pipeline.
//
// # Instance methods
//
//   - [MTLTileRenderPipelineColorAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the render pipeline state for the specified color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptorArray
type MTLTileRenderPipelineColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTLTileRenderPipelineColorAttachmentDescriptorArrayFromID constructs a [MTLTileRenderPipelineColorAttachmentDescriptorArray] from an objc.ID.
//
// An array of color attachment descriptors for the tile render pipeline.
func MTLTileRenderPipelineColorAttachmentDescriptorArrayFromID(id objc.ID) MTLTileRenderPipelineColorAttachmentDescriptorArray {
	return MTLTileRenderPipelineColorAttachmentDescriptorArray{objectivec.Object{ID: id}}
}
// NOTE: MTLTileRenderPipelineColorAttachmentDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTileRenderPipelineColorAttachmentDescriptorArray] class.
//
// # Instance methods
//
//   - [IMTLTileRenderPipelineColorAttachmentDescriptorArray.ObjectAtIndexedSubscript]: Returns the render pipeline state for the specified color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptorArray
type IMTLTileRenderPipelineColorAttachmentDescriptorArray interface {
	objectivec.IObject

	// Topic: Instance methods

	// Returns the render pipeline state for the specified color attachment.
	ObjectAtIndexedSubscript(attachmentIndex uint) IMTLTileRenderPipelineColorAttachmentDescriptor

	// Sets the render pipeline state for a specified color attachment.
	SetObjectAtIndexedSubscript(attachment IMTLTileRenderPipelineColorAttachmentDescriptor, attachmentIndex uint)
}

// Init initializes the instance.
func (t MTLTileRenderPipelineColorAttachmentDescriptorArray) Init() MTLTileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTLTileRenderPipelineColorAttachmentDescriptorArray](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTileRenderPipelineColorAttachmentDescriptorArray) Autorelease() MTLTileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTLTileRenderPipelineColorAttachmentDescriptorArray](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTileRenderPipelineColorAttachmentDescriptorArray creates a new MTLTileRenderPipelineColorAttachmentDescriptorArray instance.
func NewMTLTileRenderPipelineColorAttachmentDescriptorArray() MTLTileRenderPipelineColorAttachmentDescriptorArray {
	class := getMTLTileRenderPipelineColorAttachmentDescriptorArrayClass()
	rv := objc.Send[MTLTileRenderPipelineColorAttachmentDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the render pipeline state for the specified color attachment.
//
// attachmentIndex: An index in the color attachment array.
//
// # Return Value
// 
// An [MTLTileRenderPipelineColorAttachmentDescriptor] that describes the
// render pipeline information for a color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptorArray/subscript(_:)
func (t MTLTileRenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) IMTLTileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return MTLTileRenderPipelineColorAttachmentDescriptorFromID(rv)
}

// Sets the render pipeline state for a specified color attachment.
//
// attachment: A descriptor that contains the render pipeline description for a color
// attachment. Specify `nil` to reset the entry to default values.
//
// attachmentIndex: An index in the color attachment array.
//
// # Discussion
// 
// This method copies the pipeline state from the descriptor into the
// specified attachment in the array. Afterwards, you can modify and reuse the
// descriptior without affecting a previously set attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (t MTLTileRenderPipelineColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLTileRenderPipelineColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](t.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}


// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTileRenderPipelineColorAttachmentDescriptor] class.
var (
	_MTLTileRenderPipelineColorAttachmentDescriptorClass     MTLTileRenderPipelineColorAttachmentDescriptorClass
	_MTLTileRenderPipelineColorAttachmentDescriptorClassOnce sync.Once
)

func getMTLTileRenderPipelineColorAttachmentDescriptorClass() MTLTileRenderPipelineColorAttachmentDescriptorClass {
	_MTLTileRenderPipelineColorAttachmentDescriptorClassOnce.Do(func() {
		_MTLTileRenderPipelineColorAttachmentDescriptorClass = MTLTileRenderPipelineColorAttachmentDescriptorClass{class: objc.GetClass("MTLTileRenderPipelineColorAttachmentDescriptor")}
	})
	return _MTLTileRenderPipelineColorAttachmentDescriptorClass
}

// GetMTLTileRenderPipelineColorAttachmentDescriptorClass returns the class object for MTLTileRenderPipelineColorAttachmentDescriptor.
func GetMTLTileRenderPipelineColorAttachmentDescriptorClass() MTLTileRenderPipelineColorAttachmentDescriptorClass {
	return getMTLTileRenderPipelineColorAttachmentDescriptorClass()
}

type MTLTileRenderPipelineColorAttachmentDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLTileRenderPipelineColorAttachmentDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTileRenderPipelineColorAttachmentDescriptorClass) Alloc() MTLTileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTLTileRenderPipelineColorAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a tile-shading render pipeline’s color render target.
//
// # Specifying pixel format
//
//   - [MTLTileRenderPipelineColorAttachmentDescriptor.PixelFormat]: The pixel format associated with the tile shading render pipeline.
//   - [MTLTileRenderPipelineColorAttachmentDescriptor.SetPixelFormat]
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptor
type MTLTileRenderPipelineColorAttachmentDescriptor struct {
	objectivec.Object
}

// MTLTileRenderPipelineColorAttachmentDescriptorFromID constructs a [MTLTileRenderPipelineColorAttachmentDescriptor] from an objc.ID.
//
// A description of a tile-shading render pipeline’s color render target.
func MTLTileRenderPipelineColorAttachmentDescriptorFromID(id objc.ID) MTLTileRenderPipelineColorAttachmentDescriptor {
	return MTLTileRenderPipelineColorAttachmentDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTLTileRenderPipelineColorAttachmentDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTileRenderPipelineColorAttachmentDescriptor] class.
//
// # Specifying pixel format
//
//   - [IMTLTileRenderPipelineColorAttachmentDescriptor.PixelFormat]: The pixel format associated with the tile shading render pipeline.
//   - [IMTLTileRenderPipelineColorAttachmentDescriptor.SetPixelFormat]
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptor
type IMTLTileRenderPipelineColorAttachmentDescriptor interface {
	objectivec.IObject

	// Topic: Specifying pixel format

	// The pixel format associated with the tile shading render pipeline.
	PixelFormat() MTLPixelFormat
	SetPixelFormat(value MTLPixelFormat)
}

// Init initializes the instance.
func (t MTLTileRenderPipelineColorAttachmentDescriptor) Init() MTLTileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTLTileRenderPipelineColorAttachmentDescriptor](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTileRenderPipelineColorAttachmentDescriptor) Autorelease() MTLTileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTLTileRenderPipelineColorAttachmentDescriptor](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTileRenderPipelineColorAttachmentDescriptor creates a new MTLTileRenderPipelineColorAttachmentDescriptor instance.
func NewMTLTileRenderPipelineColorAttachmentDescriptor() MTLTileRenderPipelineColorAttachmentDescriptor {
	class := getMTLTileRenderPipelineColorAttachmentDescriptorClass()
	rv := objc.Send[MTLTileRenderPipelineColorAttachmentDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The pixel format associated with the tile shading render pipeline.
//
// # Discussion
//
// The default value is [MTLPixelFormatInvalid].
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptor/pixelFormat
func (t MTLTileRenderPipelineColorAttachmentDescriptor) PixelFormat() MTLPixelFormat {
	rv := objc.Send[MTLPixelFormat](t.ID, objc.Sel("pixelFormat"))
	return MTLPixelFormat(rv)
}
func (t MTLTileRenderPipelineColorAttachmentDescriptor) SetPixelFormat(value MTLPixelFormat) {
	objc.Send[struct{}](t.ID, objc.Sel("setPixelFormat:"), value)
}

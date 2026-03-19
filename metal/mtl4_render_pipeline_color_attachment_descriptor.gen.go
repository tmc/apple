// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4RenderPipelineColorAttachmentDescriptor] class.
var (
	_MTL4RenderPipelineColorAttachmentDescriptorClass     MTL4RenderPipelineColorAttachmentDescriptorClass
	_MTL4RenderPipelineColorAttachmentDescriptorClassOnce sync.Once
)

func getMTL4RenderPipelineColorAttachmentDescriptorClass() MTL4RenderPipelineColorAttachmentDescriptorClass {
	_MTL4RenderPipelineColorAttachmentDescriptorClassOnce.Do(func() {
		_MTL4RenderPipelineColorAttachmentDescriptorClass = MTL4RenderPipelineColorAttachmentDescriptorClass{class: objc.GetClass("MTL4RenderPipelineColorAttachmentDescriptor")}
	})
	return _MTL4RenderPipelineColorAttachmentDescriptorClass
}

// GetMTL4RenderPipelineColorAttachmentDescriptorClass returns the class object for MTL4RenderPipelineColorAttachmentDescriptor.
func GetMTL4RenderPipelineColorAttachmentDescriptorClass() MTL4RenderPipelineColorAttachmentDescriptorClass {
	return getMTL4RenderPipelineColorAttachmentDescriptorClass()
}

type MTL4RenderPipelineColorAttachmentDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4RenderPipelineColorAttachmentDescriptorClass) Alloc() MTL4RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Instance Properties
//
//   - [MTL4RenderPipelineColorAttachmentDescriptor.AlphaBlendOperation]: Configures the alpha blending operation.
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SetAlphaBlendOperation]
//   - [MTL4RenderPipelineColorAttachmentDescriptor.BlendingState]: Configure the blend state for color attachments the pipeline state uses.
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SetBlendingState]
//   - [MTL4RenderPipelineColorAttachmentDescriptor.DestinationAlphaBlendFactor]: Configures the destination-alpha blend factor.
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SetDestinationAlphaBlendFactor]
//   - [MTL4RenderPipelineColorAttachmentDescriptor.DestinationRGBBlendFactor]: Configures the destination RGB blend factor.
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SetDestinationRGBBlendFactor]
//   - [MTL4RenderPipelineColorAttachmentDescriptor.PixelFormat]: Configures the pixel format.
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SetPixelFormat]
//   - [MTL4RenderPipelineColorAttachmentDescriptor.RgbBlendOperation]: Configures the RGB blend operation.
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SetRgbBlendOperation]
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SourceAlphaBlendFactor]: Configures the source-alpha blend factor.
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SetSourceAlphaBlendFactor]
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SourceRGBBlendFactor]: Configures the source RGB blend factor.
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SetSourceRGBBlendFactor]
//   - [MTL4RenderPipelineColorAttachmentDescriptor.WriteMask]: Configures the color write mask.
//   - [MTL4RenderPipelineColorAttachmentDescriptor.SetWriteMask]
//
// # Instance Methods
//
//   - [MTL4RenderPipelineColorAttachmentDescriptor.Reset]: Resets this descriptor to its default state.
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor
type MTL4RenderPipelineColorAttachmentDescriptor struct {
	objectivec.Object
}

// MTL4RenderPipelineColorAttachmentDescriptorFromID constructs a [MTL4RenderPipelineColorAttachmentDescriptor] from an objc.ID.
func MTL4RenderPipelineColorAttachmentDescriptorFromID(id objc.ID) MTL4RenderPipelineColorAttachmentDescriptor {
	return MTL4RenderPipelineColorAttachmentDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTL4RenderPipelineColorAttachmentDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4RenderPipelineColorAttachmentDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.AlphaBlendOperation]: Configures the alpha blending operation.
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SetAlphaBlendOperation]
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.BlendingState]: Configure the blend state for color attachments the pipeline state uses.
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SetBlendingState]
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.DestinationAlphaBlendFactor]: Configures the destination-alpha blend factor.
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SetDestinationAlphaBlendFactor]
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.DestinationRGBBlendFactor]: Configures the destination RGB blend factor.
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SetDestinationRGBBlendFactor]
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.PixelFormat]: Configures the pixel format.
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SetPixelFormat]
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.RgbBlendOperation]: Configures the RGB blend operation.
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SetRgbBlendOperation]
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SourceAlphaBlendFactor]: Configures the source-alpha blend factor.
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SetSourceAlphaBlendFactor]
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SourceRGBBlendFactor]: Configures the source RGB blend factor.
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SetSourceRGBBlendFactor]
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.WriteMask]: Configures the color write mask.
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.SetWriteMask]
//
// # Instance Methods
//
//   - [IMTL4RenderPipelineColorAttachmentDescriptor.Reset]: Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor
type IMTL4RenderPipelineColorAttachmentDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Configures the alpha blending operation.
	AlphaBlendOperation() MTLBlendOperation
	SetAlphaBlendOperation(value MTLBlendOperation)
	// Configure the blend state for color attachments the pipeline state uses.
	BlendingState() MTL4BlendState
	SetBlendingState(value MTL4BlendState)
	// Configures the destination-alpha blend factor.
	DestinationAlphaBlendFactor() MTLBlendFactor
	SetDestinationAlphaBlendFactor(value MTLBlendFactor)
	// Configures the destination RGB blend factor.
	DestinationRGBBlendFactor() MTLBlendFactor
	SetDestinationRGBBlendFactor(value MTLBlendFactor)
	// Configures the pixel format.
	PixelFormat() MTLPixelFormat
	SetPixelFormat(value MTLPixelFormat)
	// Configures the RGB blend operation.
	RgbBlendOperation() MTLBlendOperation
	SetRgbBlendOperation(value MTLBlendOperation)
	// Configures the source-alpha blend factor.
	SourceAlphaBlendFactor() MTLBlendFactor
	SetSourceAlphaBlendFactor(value MTLBlendFactor)
	// Configures the source RGB blend factor.
	SourceRGBBlendFactor() MTLBlendFactor
	SetSourceRGBBlendFactor(value MTLBlendFactor)
	// Configures the color write mask.
	WriteMask() MTLColorWriteMask
	SetWriteMask(value MTLColorWriteMask)

	// Topic: Instance Methods

	// Resets this descriptor to its default state.
	Reset()
}

// Init initializes the instance.
func (m MTL4RenderPipelineColorAttachmentDescriptor) Init() MTL4RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4RenderPipelineColorAttachmentDescriptor) Autorelease() MTL4RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineColorAttachmentDescriptor creates a new MTL4RenderPipelineColorAttachmentDescriptor instance.
func NewMTL4RenderPipelineColorAttachmentDescriptor() MTL4RenderPipelineColorAttachmentDescriptor {
	class := getMTL4RenderPipelineColorAttachmentDescriptorClass()
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/reset()
func (m MTL4RenderPipelineColorAttachmentDescriptor) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// Configures the alpha blending operation.
//
// # Discussion
// 
// This property defaults to [MTLBlendOperationAdd].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (m MTL4RenderPipelineColorAttachmentDescriptor) AlphaBlendOperation() MTLBlendOperation {
	rv := objc.Send[MTLBlendOperation](m.ID, objc.Sel("alphaBlendOperation"))
	return MTLBlendOperation(rv)
}
func (m MTL4RenderPipelineColorAttachmentDescriptor) SetAlphaBlendOperation(value MTLBlendOperation) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlphaBlendOperation:"), value)
}
// Configure the blend state for color attachments the pipeline state uses.
//
// # Discussion
// 
// This property’s default value is [MTL4BlendStateDisabled].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/blendingState
func (m MTL4RenderPipelineColorAttachmentDescriptor) BlendingState() MTL4BlendState {
	rv := objc.Send[MTL4BlendState](m.ID, objc.Sel("blendingState"))
	return MTL4BlendState(rv)
}
func (m MTL4RenderPipelineColorAttachmentDescriptor) SetBlendingState(value MTL4BlendState) {
	objc.Send[struct{}](m.ID, objc.Sel("setBlendingState:"), value)
}
// Configures the destination-alpha blend factor.
//
// # Discussion
// 
// This property defaults to [MTLBlendFactorZero].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (m MTL4RenderPipelineColorAttachmentDescriptor) DestinationAlphaBlendFactor() MTLBlendFactor {
	rv := objc.Send[MTLBlendFactor](m.ID, objc.Sel("destinationAlphaBlendFactor"))
	return MTLBlendFactor(rv)
}
func (m MTL4RenderPipelineColorAttachmentDescriptor) SetDestinationAlphaBlendFactor(value MTLBlendFactor) {
	objc.Send[struct{}](m.ID, objc.Sel("setDestinationAlphaBlendFactor:"), value)
}
// Configures the destination RGB blend factor.
//
// # Discussion
// 
// This property defaults to [MTLBlendFactorZero].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (m MTL4RenderPipelineColorAttachmentDescriptor) DestinationRGBBlendFactor() MTLBlendFactor {
	rv := objc.Send[MTLBlendFactor](m.ID, objc.Sel("destinationRGBBlendFactor"))
	return MTLBlendFactor(rv)
}
func (m MTL4RenderPipelineColorAttachmentDescriptor) SetDestinationRGBBlendFactor(value MTLBlendFactor) {
	objc.Send[struct{}](m.ID, objc.Sel("setDestinationRGBBlendFactor:"), value)
}
// Configures the pixel format.
//
// # Discussion
// 
// This property defaults to [MTLPixelFormatInvalid].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/pixelFormat
func (m MTL4RenderPipelineColorAttachmentDescriptor) PixelFormat() MTLPixelFormat {
	rv := objc.Send[MTLPixelFormat](m.ID, objc.Sel("pixelFormat"))
	return MTLPixelFormat(rv)
}
func (m MTL4RenderPipelineColorAttachmentDescriptor) SetPixelFormat(value MTLPixelFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setPixelFormat:"), value)
}
// Configures the RGB blend operation.
//
// # Discussion
// 
// This property defaults to [MTLBlendOperationAdd].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (m MTL4RenderPipelineColorAttachmentDescriptor) RgbBlendOperation() MTLBlendOperation {
	rv := objc.Send[MTLBlendOperation](m.ID, objc.Sel("rgbBlendOperation"))
	return MTLBlendOperation(rv)
}
func (m MTL4RenderPipelineColorAttachmentDescriptor) SetRgbBlendOperation(value MTLBlendOperation) {
	objc.Send[struct{}](m.ID, objc.Sel("setRgbBlendOperation:"), value)
}
// Configures the source-alpha blend factor.
//
// # Discussion
// 
// This property defaults to [MTLBlendFactorOne].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (m MTL4RenderPipelineColorAttachmentDescriptor) SourceAlphaBlendFactor() MTLBlendFactor {
	rv := objc.Send[MTLBlendFactor](m.ID, objc.Sel("sourceAlphaBlendFactor"))
	return MTLBlendFactor(rv)
}
func (m MTL4RenderPipelineColorAttachmentDescriptor) SetSourceAlphaBlendFactor(value MTLBlendFactor) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceAlphaBlendFactor:"), value)
}
// Configures the source RGB blend factor.
//
// # Discussion
// 
// This property defaults to [MTLBlendFactorOne].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (m MTL4RenderPipelineColorAttachmentDescriptor) SourceRGBBlendFactor() MTLBlendFactor {
	rv := objc.Send[MTLBlendFactor](m.ID, objc.Sel("sourceRGBBlendFactor"))
	return MTLBlendFactor(rv)
}
func (m MTL4RenderPipelineColorAttachmentDescriptor) SetSourceRGBBlendFactor(value MTLBlendFactor) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceRGBBlendFactor:"), value)
}
// Configures the color write mask.
//
// # Discussion
// 
// This property defaults to [MTLColorWriteMaskAll].
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/writeMask
func (m MTL4RenderPipelineColorAttachmentDescriptor) WriteMask() MTLColorWriteMask {
	rv := objc.Send[MTLColorWriteMask](m.ID, objc.Sel("writeMask"))
	return MTLColorWriteMask(rv)
}
func (m MTL4RenderPipelineColorAttachmentDescriptor) SetWriteMask(value MTLColorWriteMask) {
	objc.Send[struct{}](m.ID, objc.Sel("setWriteMask:"), value)
}


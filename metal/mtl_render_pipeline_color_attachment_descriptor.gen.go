// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRenderPipelineColorAttachmentDescriptor] class.
var (
	_MTLRenderPipelineColorAttachmentDescriptorClass     MTLRenderPipelineColorAttachmentDescriptorClass
	_MTLRenderPipelineColorAttachmentDescriptorClassOnce sync.Once
)

func getMTLRenderPipelineColorAttachmentDescriptorClass() MTLRenderPipelineColorAttachmentDescriptorClass {
	_MTLRenderPipelineColorAttachmentDescriptorClassOnce.Do(func() {
		_MTLRenderPipelineColorAttachmentDescriptorClass = MTLRenderPipelineColorAttachmentDescriptorClass{class: objc.GetClass("MTLRenderPipelineColorAttachmentDescriptor")}
	})
	return _MTLRenderPipelineColorAttachmentDescriptorClass
}

// GetMTLRenderPipelineColorAttachmentDescriptorClass returns the class object for MTLRenderPipelineColorAttachmentDescriptor.
func GetMTLRenderPipelineColorAttachmentDescriptorClass() MTLRenderPipelineColorAttachmentDescriptorClass {
	return getMTLRenderPipelineColorAttachmentDescriptorClass()
}

type MTLRenderPipelineColorAttachmentDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRenderPipelineColorAttachmentDescriptorClass) Alloc() MTLRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTLRenderPipelineColorAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A color render target that specifies the color configuration and color
// operations for a render pipeline.
//
// # Overview
// 
// An [MTLRenderPipelineColorAttachmentDescriptor] instance defines the
// configuration of a color attachment associated with a rendering pipeline.
// 
// The [MTLRenderPipelineColorAttachmentDescriptor.PixelFormat] property needs to be specified for the rendering pipeline
// state at the color attachment.
// 
// Blend operations determine how a source fragment is combined with a
// destination value in a color attachment to determine the pixel value to be
// written. The following properties define whether and how blending is
// performed:
// 
// - The [MTLRenderPipelineColorAttachmentDescriptor.BlendingEnabled] property enables blending. The default value is
// [false]. - The [MTLRenderPipelineColorAttachmentDescriptor.WriteMask] property identifies which color channels are
// blended. The default value is [ColorWriteMaskAll], which allows all color
// channels to be blended. - The [MTLRenderPipelineColorAttachmentDescriptor.RgbBlendOperation] and [MTLRenderPipelineColorAttachmentDescriptor.AlphaBlendOperation]
// properties assign the blend operations for RGB and alpha pixel data. The
// default value for both properties is [BlendOperationAdd]. - The
// [MTLRenderPipelineColorAttachmentDescriptor.SourceRGBBlendFactor], [MTLRenderPipelineColorAttachmentDescriptor.SourceAlphaBlendFactor],
// [MTLRenderPipelineColorAttachmentDescriptor.DestinationRGBBlendFactor], and [MTLRenderPipelineColorAttachmentDescriptor.DestinationAlphaBlendFactor] properties
// assign the source and destination blend factors. The default value for
// [MTLRenderPipelineColorAttachmentDescriptor.SourceRGBBlendFactor] and [MTLRenderPipelineColorAttachmentDescriptor.SourceAlphaBlendFactor] is [BlendFactorOne].
// The default value for [MTLRenderPipelineColorAttachmentDescriptor.DestinationRGBBlendFactor] and
// [MTLRenderPipelineColorAttachmentDescriptor.DestinationAlphaBlendFactor] is [BlendFactorZero].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Configuring render pipeline states
//
//   - [MTLRenderPipelineColorAttachmentDescriptor.PixelFormat]: The pixel format of the color attachment’s texture.
//   - [MTLRenderPipelineColorAttachmentDescriptor.SetPixelFormat]
//   - [MTLRenderPipelineColorAttachmentDescriptor.WriteMask]: A bitmask that restricts which color channels are written into the texture.
//   - [MTLRenderPipelineColorAttachmentDescriptor.SetWriteMask]
//
// # Controlling blend operations
//
//   - [MTLRenderPipelineColorAttachmentDescriptor.BlendingEnabled]: A Boolean value that determines whether blending is enabled.
//   - [MTLRenderPipelineColorAttachmentDescriptor.SetBlendingEnabled]
//   - [MTLRenderPipelineColorAttachmentDescriptor.AlphaBlendOperation]: The blend operation assigned for the alpha data.
//   - [MTLRenderPipelineColorAttachmentDescriptor.SetAlphaBlendOperation]
//   - [MTLRenderPipelineColorAttachmentDescriptor.RgbBlendOperation]: The blend operation assigned for the RGB data.
//   - [MTLRenderPipelineColorAttachmentDescriptor.SetRgbBlendOperation]
//
// # Configuring blend factors
//
//   - [MTLRenderPipelineColorAttachmentDescriptor.DestinationAlphaBlendFactor]: The destination blend factor (DBF) used by the alpha blend operation.
//   - [MTLRenderPipelineColorAttachmentDescriptor.SetDestinationAlphaBlendFactor]
//   - [MTLRenderPipelineColorAttachmentDescriptor.DestinationRGBBlendFactor]: The destination blend factor (DBF) used by the RGB blend operation.
//   - [MTLRenderPipelineColorAttachmentDescriptor.SetDestinationRGBBlendFactor]
//   - [MTLRenderPipelineColorAttachmentDescriptor.SourceAlphaBlendFactor]: The source blend factor (SBF) used by the alpha blend operation.
//   - [MTLRenderPipelineColorAttachmentDescriptor.SetSourceAlphaBlendFactor]
//   - [MTLRenderPipelineColorAttachmentDescriptor.SourceRGBBlendFactor]: The source blend factor (SBF) used by the RGB blend operation.
//   - [MTLRenderPipelineColorAttachmentDescriptor.SetSourceRGBBlendFactor]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor
type MTLRenderPipelineColorAttachmentDescriptor struct {
	objectivec.Object
}

// MTLRenderPipelineColorAttachmentDescriptorFromID constructs a [MTLRenderPipelineColorAttachmentDescriptor] from an objc.ID.
//
// A color render target that specifies the color configuration and color
// operations for a render pipeline.
func MTLRenderPipelineColorAttachmentDescriptorFromID(id objc.ID) MTLRenderPipelineColorAttachmentDescriptor {
	return MTLRenderPipelineColorAttachmentDescriptor{objectivec.Object{id}}
}
// NOTE: MTLRenderPipelineColorAttachmentDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLRenderPipelineColorAttachmentDescriptor] class.
//
// # Configuring render pipeline states
//
//   - [IMTLRenderPipelineColorAttachmentDescriptor.PixelFormat]: The pixel format of the color attachment’s texture.
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SetPixelFormat]
//   - [IMTLRenderPipelineColorAttachmentDescriptor.WriteMask]: A bitmask that restricts which color channels are written into the texture.
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SetWriteMask]
//
// # Controlling blend operations
//
//   - [IMTLRenderPipelineColorAttachmentDescriptor.BlendingEnabled]: A Boolean value that determines whether blending is enabled.
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SetBlendingEnabled]
//   - [IMTLRenderPipelineColorAttachmentDescriptor.AlphaBlendOperation]: The blend operation assigned for the alpha data.
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SetAlphaBlendOperation]
//   - [IMTLRenderPipelineColorAttachmentDescriptor.RgbBlendOperation]: The blend operation assigned for the RGB data.
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SetRgbBlendOperation]
//
// # Configuring blend factors
//
//   - [IMTLRenderPipelineColorAttachmentDescriptor.DestinationAlphaBlendFactor]: The destination blend factor (DBF) used by the alpha blend operation.
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SetDestinationAlphaBlendFactor]
//   - [IMTLRenderPipelineColorAttachmentDescriptor.DestinationRGBBlendFactor]: The destination blend factor (DBF) used by the RGB blend operation.
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SetDestinationRGBBlendFactor]
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SourceAlphaBlendFactor]: The source blend factor (SBF) used by the alpha blend operation.
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SetSourceAlphaBlendFactor]
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SourceRGBBlendFactor]: The source blend factor (SBF) used by the RGB blend operation.
//   - [IMTLRenderPipelineColorAttachmentDescriptor.SetSourceRGBBlendFactor]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor
type IMTLRenderPipelineColorAttachmentDescriptor interface {
	objectivec.IObject

	// Topic: Configuring render pipeline states

	// The pixel format of the color attachment’s texture.
	PixelFormat() MTLPixelFormat
	SetPixelFormat(value MTLPixelFormat)
	// A bitmask that restricts which color channels are written into the texture.
	WriteMask() MTLColorWriteMask
	SetWriteMask(value MTLColorWriteMask)

	// Topic: Controlling blend operations

	// A Boolean value that determines whether blending is enabled.
	BlendingEnabled() bool
	SetBlendingEnabled(value bool)
	// The blend operation assigned for the alpha data.
	AlphaBlendOperation() MTLBlendOperation
	SetAlphaBlendOperation(value MTLBlendOperation)
	// The blend operation assigned for the RGB data.
	RgbBlendOperation() MTLBlendOperation
	SetRgbBlendOperation(value MTLBlendOperation)

	// Topic: Configuring blend factors

	// The destination blend factor (DBF) used by the alpha blend operation.
	DestinationAlphaBlendFactor() MTLBlendFactor
	SetDestinationAlphaBlendFactor(value MTLBlendFactor)
	// The destination blend factor (DBF) used by the RGB blend operation.
	DestinationRGBBlendFactor() MTLBlendFactor
	SetDestinationRGBBlendFactor(value MTLBlendFactor)
	// The source blend factor (SBF) used by the alpha blend operation.
	SourceAlphaBlendFactor() MTLBlendFactor
	SetSourceAlphaBlendFactor(value MTLBlendFactor)
	// The source blend factor (SBF) used by the RGB blend operation.
	SourceRGBBlendFactor() MTLBlendFactor
	SetSourceRGBBlendFactor(value MTLBlendFactor)

	// All color channels are enabled.
	All() MTLColorWriteMask
	SetAll(value MTLColorWriteMask)
}





// Init initializes the instance.
func (r MTLRenderPipelineColorAttachmentDescriptor) Init() MTLRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTLRenderPipelineColorAttachmentDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRenderPipelineColorAttachmentDescriptor) Autorelease() MTLRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTLRenderPipelineColorAttachmentDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRenderPipelineColorAttachmentDescriptor creates a new MTLRenderPipelineColorAttachmentDescriptor instance.
func NewMTLRenderPipelineColorAttachmentDescriptor() MTLRenderPipelineColorAttachmentDescriptor {
	class := getMTLRenderPipelineColorAttachmentDescriptorClass()
	rv := objc.Send[MTLRenderPipelineColorAttachmentDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The pixel format of the color attachment’s texture.
//
// # Discussion
// 
// The pixel format of the rendering pipeline state needs to be set to match
// the pixel format of the texture used by the selected color attachment;
// otherwise, an error occurs.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/pixelFormat
func (r MTLRenderPipelineColorAttachmentDescriptor) PixelFormat() MTLPixelFormat {
	rv := objc.Send[MTLPixelFormat](r.ID, objc.Sel("pixelFormat"))
	return MTLPixelFormat(rv)
}
func (r MTLRenderPipelineColorAttachmentDescriptor) SetPixelFormat(value MTLPixelFormat) {
	objc.Send[struct{}](r.ID, objc.Sel("setPixelFormat:"), value)
}



// A bitmask that restricts which color channels are written into the texture.
//
// # Discussion
// 
// The default value of `writeMask` is all ones, [ColorWriteMaskAll], which
// allows all color channels to be blended. The [MTLColorWriteMask] values
// [MTLColorWriteMaskRed], [MTLColorWriteMaskGreen], [MTLColorWriteMaskBlue],
// and [MTLColorWriteMaskAlpha] limit blending to one color channel, and these
// values can be bitwise combined. [MTLColorWriteMaskNone] does not allow any
// color channels to be blended.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/writeMask
func (r MTLRenderPipelineColorAttachmentDescriptor) WriteMask() MTLColorWriteMask {
	rv := objc.Send[MTLColorWriteMask](r.ID, objc.Sel("writeMask"))
	return MTLColorWriteMask(rv)
}
func (r MTLRenderPipelineColorAttachmentDescriptor) SetWriteMask(value MTLColorWriteMask) {
	objc.Send[struct{}](r.ID, objc.Sel("setWriteMask:"), value)
}



// A Boolean value that determines whether blending is enabled.
//
// # Discussion
// 
// The default value is [false], meaning blending is disabled and pixel values
// are unaffected by blending. Disabled blending is effectively the same as
// the [MTLBlendOperationAdd] blend operation with a source blend factor of
// `1.0` and a destination blend factor of `0.0` for both RGB and alpha.
// 
// If the value is [true], blending is enabled and the blend descriptor
// property values are used to determine how source and destination color
// values are combined.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/isBlendingEnabled
func (r MTLRenderPipelineColorAttachmentDescriptor) BlendingEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isBlendingEnabled"))
	return rv
}
func (r MTLRenderPipelineColorAttachmentDescriptor) SetBlendingEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setBlendingEnabled:"), value)
}



// The blend operation assigned for the alpha data.
//
// # Discussion
// 
// The default value is [BlendOperationAdd].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (r MTLRenderPipelineColorAttachmentDescriptor) AlphaBlendOperation() MTLBlendOperation {
	rv := objc.Send[MTLBlendOperation](r.ID, objc.Sel("alphaBlendOperation"))
	return MTLBlendOperation(rv)
}
func (r MTLRenderPipelineColorAttachmentDescriptor) SetAlphaBlendOperation(value MTLBlendOperation) {
	objc.Send[struct{}](r.ID, objc.Sel("setAlphaBlendOperation:"), value)
}



// The blend operation assigned for the RGB data.
//
// # Discussion
// 
// The default value is [BlendOperationAdd].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (r MTLRenderPipelineColorAttachmentDescriptor) RgbBlendOperation() MTLBlendOperation {
	rv := objc.Send[MTLBlendOperation](r.ID, objc.Sel("rgbBlendOperation"))
	return MTLBlendOperation(rv)
}
func (r MTLRenderPipelineColorAttachmentDescriptor) SetRgbBlendOperation(value MTLBlendOperation) {
	objc.Send[struct{}](r.ID, objc.Sel("setRgbBlendOperation:"), value)
}



// The destination blend factor (DBF) used by the alpha blend operation.
//
// # Discussion
// 
// The default value is [BlendFactorZero].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (r MTLRenderPipelineColorAttachmentDescriptor) DestinationAlphaBlendFactor() MTLBlendFactor {
	rv := objc.Send[MTLBlendFactor](r.ID, objc.Sel("destinationAlphaBlendFactor"))
	return MTLBlendFactor(rv)
}
func (r MTLRenderPipelineColorAttachmentDescriptor) SetDestinationAlphaBlendFactor(value MTLBlendFactor) {
	objc.Send[struct{}](r.ID, objc.Sel("setDestinationAlphaBlendFactor:"), value)
}



// The destination blend factor (DBF) used by the RGB blend operation.
//
// # Discussion
// 
// The default value is [BlendFactorZero].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (r MTLRenderPipelineColorAttachmentDescriptor) DestinationRGBBlendFactor() MTLBlendFactor {
	rv := objc.Send[MTLBlendFactor](r.ID, objc.Sel("destinationRGBBlendFactor"))
	return MTLBlendFactor(rv)
}
func (r MTLRenderPipelineColorAttachmentDescriptor) SetDestinationRGBBlendFactor(value MTLBlendFactor) {
	objc.Send[struct{}](r.ID, objc.Sel("setDestinationRGBBlendFactor:"), value)
}



// The source blend factor (SBF) used by the alpha blend operation.
//
// # Discussion
// 
// The default value is [BlendFactorOne].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (r MTLRenderPipelineColorAttachmentDescriptor) SourceAlphaBlendFactor() MTLBlendFactor {
	rv := objc.Send[MTLBlendFactor](r.ID, objc.Sel("sourceAlphaBlendFactor"))
	return MTLBlendFactor(rv)
}
func (r MTLRenderPipelineColorAttachmentDescriptor) SetSourceAlphaBlendFactor(value MTLBlendFactor) {
	objc.Send[struct{}](r.ID, objc.Sel("setSourceAlphaBlendFactor:"), value)
}



// The source blend factor (SBF) used by the RGB blend operation.
//
// # Discussion
// 
// The default value is [BlendFactorOne].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (r MTLRenderPipelineColorAttachmentDescriptor) SourceRGBBlendFactor() MTLBlendFactor {
	rv := objc.Send[MTLBlendFactor](r.ID, objc.Sel("sourceRGBBlendFactor"))
	return MTLBlendFactor(rv)
}
func (r MTLRenderPipelineColorAttachmentDescriptor) SetSourceRGBBlendFactor(value MTLBlendFactor) {
	objc.Send[struct{}](r.ID, objc.Sel("setSourceRGBBlendFactor:"), value)
}



// All color channels are enabled.
//
// See: https://developer.apple.com/documentation/metal/mtlcolorwritemask/all
func (r MTLRenderPipelineColorAttachmentDescriptor) All() MTLColorWriteMask {
	rv := objc.Send[MTLColorWriteMask](r.ID, objc.Sel("MTLColorWriteMaskAll"))
	return MTLColorWriteMask(rv)
}
func (r MTLRenderPipelineColorAttachmentDescriptor) SetAll(value MTLColorWriteMask) {
	objc.Send[struct{}](r.ID, objc.Sel("setMTLColorWriteMaskAll:"), value)
}

























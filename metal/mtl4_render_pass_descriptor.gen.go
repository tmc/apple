// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4RenderPassDescriptor] class.
var (
	_MTL4RenderPassDescriptorClass     MTL4RenderPassDescriptorClass
	_MTL4RenderPassDescriptorClassOnce sync.Once
)

func getMTL4RenderPassDescriptorClass() MTL4RenderPassDescriptorClass {
	_MTL4RenderPassDescriptorClassOnce.Do(func() {
		_MTL4RenderPassDescriptorClass = MTL4RenderPassDescriptorClass{class: objc.GetClass("MTL4RenderPassDescriptor")}
	})
	return _MTL4RenderPassDescriptorClass
}

// GetMTL4RenderPassDescriptorClass returns the class object for MTL4RenderPassDescriptor.
func GetMTL4RenderPassDescriptorClass() MTL4RenderPassDescriptorClass {
	return getMTL4RenderPassDescriptorClass()
}

type MTL4RenderPassDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4RenderPassDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4RenderPassDescriptorClass) Alloc() MTL4RenderPassDescriptor {
	rv := objc.Send[MTL4RenderPassDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Describes a render pass.
//
// # Overview
//
// You use render pass descriptors to create instances of
// [MTL4RenderCommandEncoder] and encode draw commands into instances of
// [MTL4CommandBuffer].
//
// To create render command encoders, you typically call
// [RenderCommandEncoderWithDescriptor]. The
// [RenderCommandEncoderWithDescriptorOptions] variant of this method allows
// you to specify additional options to encode a render pass in parallel from
// multiple CPU cores by creating and render passes.
//
// # Instance Properties
//
//   - [MTL4RenderPassDescriptor.ColorAttachments]: Accesses the array of state information for render attachments that store color data.
//   - [MTL4RenderPassDescriptor.DefaultRasterSampleCount]: Sets the default raster sample count for the render pass when it references no attachments.
//   - [MTL4RenderPassDescriptor.SetDefaultRasterSampleCount]
//   - [MTL4RenderPassDescriptor.DepthAttachment]: Accesses state information for a render attachment that stores depth data.
//   - [MTL4RenderPassDescriptor.SetDepthAttachment]
//   - [MTL4RenderPassDescriptor.ImageblockSampleLength]: Assigns the per-sample size, in bytes, of the largest explicit imageblock layout in the render pass.
//   - [MTL4RenderPassDescriptor.SetImageblockSampleLength]
//   - [MTL4RenderPassDescriptor.RasterizationRateMap]: Assigns an optional variable rasterization rate map that Metal uses in the render pass.
//   - [MTL4RenderPassDescriptor.SetRasterizationRateMap]
//   - [MTL4RenderPassDescriptor.RenderTargetArrayLength]: Assigns the number of layers that all attachments this descriptor references have.
//   - [MTL4RenderPassDescriptor.SetRenderTargetArrayLength]
//   - [MTL4RenderPassDescriptor.RenderTargetHeight]: Sets the height, in pixels, to which Metal constrains the render target.
//   - [MTL4RenderPassDescriptor.SetRenderTargetHeight]
//   - [MTL4RenderPassDescriptor.RenderTargetWidth]: Sets the width, in pixels, to which Metal constrains the render target.
//   - [MTL4RenderPassDescriptor.SetRenderTargetWidth]
//   - [MTL4RenderPassDescriptor.SamplePositions]: Configures the custom sample positions to use in MSAA rendering.
//   - [MTL4RenderPassDescriptor.SetSamplePositions]
//   - [MTL4RenderPassDescriptor.StencilAttachment]: Accesses state information for a render attachment that stores stencil data.
//   - [MTL4RenderPassDescriptor.SetStencilAttachment]
//   - [MTL4RenderPassDescriptor.SupportColorAttachmentMapping]: Controls if the render pass supports color attachment mapping.
//   - [MTL4RenderPassDescriptor.SetSupportColorAttachmentMapping]
//   - [MTL4RenderPassDescriptor.ThreadgroupMemoryLength]: Assigns the per-tile size, in bytes, of the persistent threadgroup memory allocation of this render pass.
//   - [MTL4RenderPassDescriptor.SetThreadgroupMemoryLength]
//   - [MTL4RenderPassDescriptor.TileHeight]: The height of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
//   - [MTL4RenderPassDescriptor.SetTileHeight]
//   - [MTL4RenderPassDescriptor.TileWidth]: The width of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
//   - [MTL4RenderPassDescriptor.SetTileWidth]
//   - [MTL4RenderPassDescriptor.VisibilityResultBuffer]: Configures a buffer into which Metal writes counts of fragments (pixels) passing the depth and stencil tests.
//   - [MTL4RenderPassDescriptor.SetVisibilityResultBuffer]
//   - [MTL4RenderPassDescriptor.VisibilityResultType]: Determines if Metal accumulates visibility results between render encoders or resets them.
//   - [MTL4RenderPassDescriptor.SetVisibilityResultType]
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor
type MTL4RenderPassDescriptor struct {
	objectivec.Object
}

// MTL4RenderPassDescriptorFromID constructs a [MTL4RenderPassDescriptor] from an objc.ID.
//
// Describes a render pass.
func MTL4RenderPassDescriptorFromID(id objc.ID) MTL4RenderPassDescriptor {
	return MTL4RenderPassDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTL4RenderPassDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4RenderPassDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4RenderPassDescriptor.ColorAttachments]: Accesses the array of state information for render attachments that store color data.
//   - [IMTL4RenderPassDescriptor.DefaultRasterSampleCount]: Sets the default raster sample count for the render pass when it references no attachments.
//   - [IMTL4RenderPassDescriptor.SetDefaultRasterSampleCount]
//   - [IMTL4RenderPassDescriptor.DepthAttachment]: Accesses state information for a render attachment that stores depth data.
//   - [IMTL4RenderPassDescriptor.SetDepthAttachment]
//   - [IMTL4RenderPassDescriptor.ImageblockSampleLength]: Assigns the per-sample size, in bytes, of the largest explicit imageblock layout in the render pass.
//   - [IMTL4RenderPassDescriptor.SetImageblockSampleLength]
//   - [IMTL4RenderPassDescriptor.RasterizationRateMap]: Assigns an optional variable rasterization rate map that Metal uses in the render pass.
//   - [IMTL4RenderPassDescriptor.SetRasterizationRateMap]
//   - [IMTL4RenderPassDescriptor.RenderTargetArrayLength]: Assigns the number of layers that all attachments this descriptor references have.
//   - [IMTL4RenderPassDescriptor.SetRenderTargetArrayLength]
//   - [IMTL4RenderPassDescriptor.RenderTargetHeight]: Sets the height, in pixels, to which Metal constrains the render target.
//   - [IMTL4RenderPassDescriptor.SetRenderTargetHeight]
//   - [IMTL4RenderPassDescriptor.RenderTargetWidth]: Sets the width, in pixels, to which Metal constrains the render target.
//   - [IMTL4RenderPassDescriptor.SetRenderTargetWidth]
//   - [IMTL4RenderPassDescriptor.SamplePositions]: Configures the custom sample positions to use in MSAA rendering.
//   - [IMTL4RenderPassDescriptor.SetSamplePositions]
//   - [IMTL4RenderPassDescriptor.StencilAttachment]: Accesses state information for a render attachment that stores stencil data.
//   - [IMTL4RenderPassDescriptor.SetStencilAttachment]
//   - [IMTL4RenderPassDescriptor.SupportColorAttachmentMapping]: Controls if the render pass supports color attachment mapping.
//   - [IMTL4RenderPassDescriptor.SetSupportColorAttachmentMapping]
//   - [IMTL4RenderPassDescriptor.ThreadgroupMemoryLength]: Assigns the per-tile size, in bytes, of the persistent threadgroup memory allocation of this render pass.
//   - [IMTL4RenderPassDescriptor.SetThreadgroupMemoryLength]
//   - [IMTL4RenderPassDescriptor.TileHeight]: The height of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
//   - [IMTL4RenderPassDescriptor.SetTileHeight]
//   - [IMTL4RenderPassDescriptor.TileWidth]: The width of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
//   - [IMTL4RenderPassDescriptor.SetTileWidth]
//   - [IMTL4RenderPassDescriptor.VisibilityResultBuffer]: Configures a buffer into which Metal writes counts of fragments (pixels) passing the depth and stencil tests.
//   - [IMTL4RenderPassDescriptor.SetVisibilityResultBuffer]
//   - [IMTL4RenderPassDescriptor.VisibilityResultType]: Determines if Metal accumulates visibility results between render encoders or resets them.
//   - [IMTL4RenderPassDescriptor.SetVisibilityResultType]
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor
type IMTL4RenderPassDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Accesses the array of state information for render attachments that store color data.
	ColorAttachments() IMTLRenderPassColorAttachmentDescriptorArray
	// Sets the default raster sample count for the render pass when it references no attachments.
	DefaultRasterSampleCount() uint
	SetDefaultRasterSampleCount(value uint)
	// Accesses state information for a render attachment that stores depth data.
	DepthAttachment() IMTLRenderPassDepthAttachmentDescriptor
	SetDepthAttachment(value IMTLRenderPassDepthAttachmentDescriptor)
	// Assigns the per-sample size, in bytes, of the largest explicit imageblock layout in the render pass.
	ImageblockSampleLength() uint
	SetImageblockSampleLength(value uint)
	// Assigns an optional variable rasterization rate map that Metal uses in the render pass.
	RasterizationRateMap() MTLRasterizationRateMap
	SetRasterizationRateMap(value MTLRasterizationRateMap)
	// Assigns the number of layers that all attachments this descriptor references have.
	RenderTargetArrayLength() uint
	SetRenderTargetArrayLength(value uint)
	// Sets the height, in pixels, to which Metal constrains the render target.
	RenderTargetHeight() uint
	SetRenderTargetHeight(value uint)
	// Sets the width, in pixels, to which Metal constrains the render target.
	RenderTargetWidth() uint
	SetRenderTargetWidth(value uint)
	// Configures the custom sample positions to use in MSAA rendering.
	SamplePositions() MTLSamplePosition
	SetSamplePositions(value MTLSamplePosition)
	// Accesses state information for a render attachment that stores stencil data.
	StencilAttachment() IMTLRenderPassStencilAttachmentDescriptor
	SetStencilAttachment(value IMTLRenderPassStencilAttachmentDescriptor)
	// Controls if the render pass supports color attachment mapping.
	SupportColorAttachmentMapping() bool
	SetSupportColorAttachmentMapping(value bool)
	// Assigns the per-tile size, in bytes, of the persistent threadgroup memory allocation of this render pass.
	ThreadgroupMemoryLength() uint
	SetThreadgroupMemoryLength(value uint)
	// The height of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
	TileHeight() uint
	SetTileHeight(value uint)
	// The width of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
	TileWidth() uint
	SetTileWidth(value uint)
	// Configures a buffer into which Metal writes counts of fragments (pixels) passing the depth and stencil tests.
	VisibilityResultBuffer() MTLBuffer
	SetVisibilityResultBuffer(value MTLBuffer)
	// Determines if Metal accumulates visibility results between render encoders or resets them.
	VisibilityResultType() MTLVisibilityResultType
	SetVisibilityResultType(value MTLVisibilityResultType)

	// Configures the custom sample positions to use in MSAA rendering.
	SetSamplePositionsCount(positions []MTLSamplePosition, count uint)
}

// Init initializes the instance.
func (m MTL4RenderPassDescriptor) Init() MTL4RenderPassDescriptor {
	rv := objc.Send[MTL4RenderPassDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4RenderPassDescriptor) Autorelease() MTL4RenderPassDescriptor {
	rv := objc.Send[MTL4RenderPassDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPassDescriptor creates a new MTL4RenderPassDescriptor instance.
func NewMTL4RenderPassDescriptor() MTL4RenderPassDescriptor {
	class := getMTL4RenderPassDescriptorClass()
	rv := objc.Send[MTL4RenderPassDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Configures the custom sample positions to use in MSAA rendering.
//
// positions: Array of [MTLSamplePosition] instances.
//
// count: Number of [MTLSamplePosition] instances in the array. This value needs to
// be a valid sample count, or `0` to disable custom sample positions.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/setSamplePositions:count:
//
// [MTLSamplePosition]: https://developer.apple.com/documentation/Metal/MTLSamplePosition
//
// [MTLSamplePosition]: https://developer.apple.com/documentation/Metal/MTLSamplePosition
func (m MTL4RenderPassDescriptor) SetSamplePositionsCount(positions []MTLSamplePosition, count uint) {
	objc.Send[objc.ID](m.ID, objc.Sel("setSamplePositions:count:"), objc.CArray(positions), count)
}

// Accesses the array of state information for render attachments that store
// color data.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/colorAttachments
func (m MTL4RenderPassDescriptor) ColorAttachments() IMTLRenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("colorAttachments"))
	return MTLRenderPassColorAttachmentDescriptorArrayFromID(objc.ID(rv))
}

// Sets the default raster sample count for the render pass when it references
// no attachments.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/defaultRasterSampleCount
func (m MTL4RenderPassDescriptor) DefaultRasterSampleCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("defaultRasterSampleCount"))
	return rv
}
func (m MTL4RenderPassDescriptor) SetDefaultRasterSampleCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setDefaultRasterSampleCount:"), value)
}

// Accesses state information for a render attachment that stores depth data.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/depthAttachment
func (m MTL4RenderPassDescriptor) DepthAttachment() IMTLRenderPassDepthAttachmentDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("depthAttachment"))
	return MTLRenderPassDepthAttachmentDescriptorFromID(objc.ID(rv))
}
func (m MTL4RenderPassDescriptor) SetDepthAttachment(value IMTLRenderPassDepthAttachmentDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setDepthAttachment:"), value)
}

// Assigns the per-sample size, in bytes, of the largest explicit imageblock
// layout in the render pass.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/imageblockSampleLength
func (m MTL4RenderPassDescriptor) ImageblockSampleLength() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("imageblockSampleLength"))
	return rv
}
func (m MTL4RenderPassDescriptor) SetImageblockSampleLength(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setImageblockSampleLength:"), value)
}

// Assigns an optional variable rasterization rate map that Metal uses in the
// render pass.
//
// # Discussion
//
// Enabling variable rasterization rate allows Metal to decrease the
// rasterization rate, typically in unimportant regions of color attachments,
// to accelerate processing.
//
// When set to `nil`, the default, Metal doesn’t use variable rasterization
// rate.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/rasterizationRateMap
func (m MTL4RenderPassDescriptor) RasterizationRateMap() MTLRasterizationRateMap {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("rasterizationRateMap"))
	return MTLRasterizationRateMapObjectFromID(rv)
}
func (m MTL4RenderPassDescriptor) SetRasterizationRateMap(value MTLRasterizationRateMap) {
	objc.Send[struct{}](m.ID, objc.Sel("setRasterizationRateMap:"), value)
}

// Assigns the number of layers that all attachments this descriptor
// references have.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetArrayLength
func (m MTL4RenderPassDescriptor) RenderTargetArrayLength() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("renderTargetArrayLength"))
	return rv
}
func (m MTL4RenderPassDescriptor) SetRenderTargetArrayLength(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRenderTargetArrayLength:"), value)
}

// Sets the height, in pixels, to which Metal constrains the render target.
//
// # Discussion
//
// When this value is non-zero, you need to assign it to be smaller than or
// equal to the minimum height of all attachments.
//
// The default value of this property is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetHeight
func (m MTL4RenderPassDescriptor) RenderTargetHeight() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("renderTargetHeight"))
	return rv
}
func (m MTL4RenderPassDescriptor) SetRenderTargetHeight(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRenderTargetHeight:"), value)
}

// Sets the width, in pixels, to which Metal constrains the render target.
//
// # Discussion
//
// When this value is non-zero, you need to assign it to be smaller than or
// equal to the minimum width of all attachments.
//
// The default value of this property is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetWidth
func (m MTL4RenderPassDescriptor) RenderTargetWidth() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("renderTargetWidth"))
	return rv
}
func (m MTL4RenderPassDescriptor) SetRenderTargetWidth(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRenderTargetWidth:"), value)
}

// Configures the custom sample positions to use in MSAA rendering.
//
// See: https://developer.apple.com/documentation/metal/mtl4renderpassdescriptor/samplepositions
func (m MTL4RenderPassDescriptor) SamplePositions() MTLSamplePosition {
	rv := objc.Send[MTLSamplePosition](m.ID, objc.Sel("samplePositions"))
	return MTLSamplePosition(rv)
}
func (m MTL4RenderPassDescriptor) SetSamplePositions(value MTLSamplePosition) {
	objc.Send[struct{}](m.ID, objc.Sel("setSamplePositions:"), value)
}

// Accesses state information for a render attachment that stores stencil
// data.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/stencilAttachment
func (m MTL4RenderPassDescriptor) StencilAttachment() IMTLRenderPassStencilAttachmentDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stencilAttachment"))
	return MTLRenderPassStencilAttachmentDescriptorFromID(objc.ID(rv))
}
func (m MTL4RenderPassDescriptor) SetStencilAttachment(value IMTLRenderPassStencilAttachmentDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setStencilAttachment:"), value)
}

// Controls if the render pass supports color attachment mapping.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/supportColorAttachmentMapping
func (m MTL4RenderPassDescriptor) SupportColorAttachmentMapping() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportColorAttachmentMapping"))
	return rv
}
func (m MTL4RenderPassDescriptor) SetSupportColorAttachmentMapping(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportColorAttachmentMapping:"), value)
}

// Assigns the per-tile size, in bytes, of the persistent threadgroup memory
// allocation of this render pass.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/threadgroupMemoryLength
func (m MTL4RenderPassDescriptor) ThreadgroupMemoryLength() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("threadgroupMemoryLength"))
	return rv
}
func (m MTL4RenderPassDescriptor) SetThreadgroupMemoryLength(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setThreadgroupMemoryLength:"), value)
}

// The height of the tiles, in pixels, a render pass you create with this
// descriptor applies to its attachments.
//
// # Discussion
//
// For tile-based rendering, Metal divides each render attachment into smaller
// regions, or . The property’s default is `0`, which tells Metal to select
// a size that fits in tile memory.
//
// See [Tailor your apps for Apple GPUs and tile-based deferred rendering] for
// more information about tiles, tile memory, and deferred rendering.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/tileHeight
//
// [Tailor your apps for Apple GPUs and tile-based deferred rendering]: https://developer.apple.com/documentation/Metal/tailor-your-apps-for-apple-gpus-and-tile-based-deferred-rendering
func (m MTL4RenderPassDescriptor) TileHeight() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("tileHeight"))
	return rv
}
func (m MTL4RenderPassDescriptor) SetTileHeight(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setTileHeight:"), value)
}

// The width of the tiles, in pixels, a render pass you create with this
// descriptor applies to its attachments.
//
// # Discussion
//
// For tile-based rendering, Metal divides each render attachment into smaller
// regions, or . The property’s default is `0`, which tells Metal to select
// a size that fits in tile memory.
//
// See [Tailor your apps for Apple GPUs and tile-based deferred rendering] for
// more information about tiles, tile memory, and deferred rendering.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/tileWidth
//
// [Tailor your apps for Apple GPUs and tile-based deferred rendering]: https://developer.apple.com/documentation/Metal/tailor-your-apps-for-apple-gpus-and-tile-based-deferred-rendering
func (m MTL4RenderPassDescriptor) TileWidth() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("tileWidth"))
	return rv
}
func (m MTL4RenderPassDescriptor) SetTileWidth(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setTileWidth:"), value)
}

// Configures a buffer into which Metal writes counts of fragments (pixels)
// passing the depth and stencil tests.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/visibilityResultBuffer
func (m MTL4RenderPassDescriptor) VisibilityResultBuffer() MTLBuffer {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("visibilityResultBuffer"))
	return MTLBufferObjectFromID(rv)
}
func (m MTL4RenderPassDescriptor) SetVisibilityResultBuffer(value MTLBuffer) {
	objc.Send[struct{}](m.ID, objc.Sel("setVisibilityResultBuffer:"), value)
}

// Determines if Metal accumulates visibility results between render encoders
// or resets them.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/visibilityResultType
func (m MTL4RenderPassDescriptor) VisibilityResultType() MTLVisibilityResultType {
	rv := objc.Send[MTLVisibilityResultType](m.ID, objc.Sel("visibilityResultType"))
	return MTLVisibilityResultType(rv)
}
func (m MTL4RenderPassDescriptor) SetVisibilityResultType(value MTLVisibilityResultType) {
	objc.Send[struct{}](m.ID, objc.Sel("setVisibilityResultType:"), value)
}

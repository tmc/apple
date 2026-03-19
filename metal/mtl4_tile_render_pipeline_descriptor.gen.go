// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4TileRenderPipelineDescriptor] class.
var (
	_MTL4TileRenderPipelineDescriptorClass     MTL4TileRenderPipelineDescriptorClass
	_MTL4TileRenderPipelineDescriptorClassOnce sync.Once
)

func getMTL4TileRenderPipelineDescriptorClass() MTL4TileRenderPipelineDescriptorClass {
	_MTL4TileRenderPipelineDescriptorClassOnce.Do(func() {
		_MTL4TileRenderPipelineDescriptorClass = MTL4TileRenderPipelineDescriptorClass{class: objc.GetClass("MTL4TileRenderPipelineDescriptor")}
	})
	return _MTL4TileRenderPipelineDescriptorClass
}

// GetMTL4TileRenderPipelineDescriptorClass returns the class object for MTL4TileRenderPipelineDescriptor.
func GetMTL4TileRenderPipelineDescriptorClass() MTL4TileRenderPipelineDescriptorClass {
	return getMTL4TileRenderPipelineDescriptorClass()
}

type MTL4TileRenderPipelineDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4TileRenderPipelineDescriptorClass) Alloc() MTL4TileRenderPipelineDescriptor {
	rv := objc.Send[MTL4TileRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together properties you use to create a tile render pipeline state
// object.
//
// # Instance Properties
//
//   - [MTL4TileRenderPipelineDescriptor.ColorAttachments]: Access an array of descriptors that configure the properties of each color attachment in the tile render pipeline.
//   - [MTL4TileRenderPipelineDescriptor.MaxTotalThreadsPerThreadgroup]: Sets the maximum number of threads that the GPU can execute simultaneously within a single threadgroup in the tile render pipeline.
//   - [MTL4TileRenderPipelineDescriptor.SetMaxTotalThreadsPerThreadgroup]
//   - [MTL4TileRenderPipelineDescriptor.RasterSampleCount]: Configures the number of samples per pixel used for multisampling.
//   - [MTL4TileRenderPipelineDescriptor.SetRasterSampleCount]
//   - [MTL4TileRenderPipelineDescriptor.RequiredThreadsPerThreadgroup]: Sets the required number of threads per threadgroup for tile dispatches.
//   - [MTL4TileRenderPipelineDescriptor.SetRequiredThreadsPerThreadgroup]
//   - [MTL4TileRenderPipelineDescriptor.StaticLinkingDescriptor]: Configures an object that contains information about functions to link to the tile render pipeline when Metal builds it.
//   - [MTL4TileRenderPipelineDescriptor.SetStaticLinkingDescriptor]
//   - [MTL4TileRenderPipelineDescriptor.SupportBinaryLinking]: Indicates whether the pipeline supports linking binary functions.
//   - [MTL4TileRenderPipelineDescriptor.SetSupportBinaryLinking]
//   - [MTL4TileRenderPipelineDescriptor.ThreadgroupSizeMatchesTileSize]: Indicating whether the size of the threadgroup matches the size of a tile in the render pipeline.
//   - [MTL4TileRenderPipelineDescriptor.SetThreadgroupSizeMatchesTileSize]
//   - [MTL4TileRenderPipelineDescriptor.TileFunctionDescriptor]: Configures the tile function that the render pipeline executes for each tile in the tile shader stage.
//   - [MTL4TileRenderPipelineDescriptor.SetTileFunctionDescriptor]
//
// # Instance Methods
//
//   - [MTL4TileRenderPipelineDescriptor.Reset]: Resets the descriptor to the default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor
type MTL4TileRenderPipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4TileRenderPipelineDescriptorFromID constructs a [MTL4TileRenderPipelineDescriptor] from an objc.ID.
//
// Groups together properties you use to create a tile render pipeline state
// object.
func MTL4TileRenderPipelineDescriptorFromID(id objc.ID) MTL4TileRenderPipelineDescriptor {
	return MTL4TileRenderPipelineDescriptor{MTL4PipelineDescriptor: MTL4PipelineDescriptorFromID(id)}
}
// NOTE: MTL4TileRenderPipelineDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4TileRenderPipelineDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4TileRenderPipelineDescriptor.ColorAttachments]: Access an array of descriptors that configure the properties of each color attachment in the tile render pipeline.
//   - [IMTL4TileRenderPipelineDescriptor.MaxTotalThreadsPerThreadgroup]: Sets the maximum number of threads that the GPU can execute simultaneously within a single threadgroup in the tile render pipeline.
//   - [IMTL4TileRenderPipelineDescriptor.SetMaxTotalThreadsPerThreadgroup]
//   - [IMTL4TileRenderPipelineDescriptor.RasterSampleCount]: Configures the number of samples per pixel used for multisampling.
//   - [IMTL4TileRenderPipelineDescriptor.SetRasterSampleCount]
//   - [IMTL4TileRenderPipelineDescriptor.RequiredThreadsPerThreadgroup]: Sets the required number of threads per threadgroup for tile dispatches.
//   - [IMTL4TileRenderPipelineDescriptor.SetRequiredThreadsPerThreadgroup]
//   - [IMTL4TileRenderPipelineDescriptor.StaticLinkingDescriptor]: Configures an object that contains information about functions to link to the tile render pipeline when Metal builds it.
//   - [IMTL4TileRenderPipelineDescriptor.SetStaticLinkingDescriptor]
//   - [IMTL4TileRenderPipelineDescriptor.SupportBinaryLinking]: Indicates whether the pipeline supports linking binary functions.
//   - [IMTL4TileRenderPipelineDescriptor.SetSupportBinaryLinking]
//   - [IMTL4TileRenderPipelineDescriptor.ThreadgroupSizeMatchesTileSize]: Indicating whether the size of the threadgroup matches the size of a tile in the render pipeline.
//   - [IMTL4TileRenderPipelineDescriptor.SetThreadgroupSizeMatchesTileSize]
//   - [IMTL4TileRenderPipelineDescriptor.TileFunctionDescriptor]: Configures the tile function that the render pipeline executes for each tile in the tile shader stage.
//   - [IMTL4TileRenderPipelineDescriptor.SetTileFunctionDescriptor]
//
// # Instance Methods
//
//   - [IMTL4TileRenderPipelineDescriptor.Reset]: Resets the descriptor to the default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor
type IMTL4TileRenderPipelineDescriptor interface {
	IMTL4PipelineDescriptor

	// Topic: Instance Properties

	// Access an array of descriptors that configure the properties of each color attachment in the tile render pipeline.
	ColorAttachments() IMTLTileRenderPipelineColorAttachmentDescriptorArray
	// Sets the maximum number of threads that the GPU can execute simultaneously within a single threadgroup in the tile render pipeline.
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	// Configures the number of samples per pixel used for multisampling.
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	// Sets the required number of threads per threadgroup for tile dispatches.
	RequiredThreadsPerThreadgroup() MTLSize
	SetRequiredThreadsPerThreadgroup(value MTLSize)
	// Configures an object that contains information about functions to link to the tile render pipeline when Metal builds it.
	StaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	// Indicates whether the pipeline supports linking binary functions.
	SupportBinaryLinking() bool
	SetSupportBinaryLinking(value bool)
	// Indicating whether the size of the threadgroup matches the size of a tile in the render pipeline.
	ThreadgroupSizeMatchesTileSize() bool
	SetThreadgroupSizeMatchesTileSize(value bool)
	// Configures the tile function that the render pipeline executes for each tile in the tile shader stage.
	TileFunctionDescriptor() IMTL4FunctionDescriptor
	SetTileFunctionDescriptor(value IMTL4FunctionDescriptor)

	// Topic: Instance Methods

	// Resets the descriptor to the default state.
	Reset()
}

// Init initializes the instance.
func (m MTL4TileRenderPipelineDescriptor) Init() MTL4TileRenderPipelineDescriptor {
	rv := objc.Send[MTL4TileRenderPipelineDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4TileRenderPipelineDescriptor) Autorelease() MTL4TileRenderPipelineDescriptor {
	rv := objc.Send[MTL4TileRenderPipelineDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4TileRenderPipelineDescriptor creates a new MTL4TileRenderPipelineDescriptor instance.
func NewMTL4TileRenderPipelineDescriptor() MTL4TileRenderPipelineDescriptor {
	class := getMTL4TileRenderPipelineDescriptorClass()
	rv := objc.Send[MTL4TileRenderPipelineDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Resets the descriptor to the default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/reset()
func (m MTL4TileRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// Access an array of descriptors that configure the properties of each color
// attachment in the tile render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/colorAttachments
func (m MTL4TileRenderPipelineDescriptor) ColorAttachments() IMTLTileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("colorAttachments"))
	return MTLTileRenderPipelineColorAttachmentDescriptorArrayFromID(objc.ID(rv))
}
// Sets the maximum number of threads that the GPU can execute simultaneously
// within a single threadgroup in the tile render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m MTL4TileRenderPipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}
func (m MTL4TileRenderPipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}
// Configures the number of samples per pixel used for multisampling.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/rasterSampleCount
func (m MTL4TileRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("rasterSampleCount"))
	return rv
}
func (m MTL4TileRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRasterSampleCount:"), value)
}
// Sets the required number of threads per threadgroup for tile dispatches.
//
// # Discussion
// 
// This value is typically optional, except in the cases where the tile
// function that [TileFunctionDescriptor] references uses
// [CooperativeTensors]. In this case, you need to provide a non-zero value to
// this property.
// 
// Additionally, when you set this value, the `threadsPerTile` argument of any
// tile dispatch needs to match it.
// 
// Setting this value to a size of 0 in every dimension disables this
// property.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/requiredThreadsPerThreadgroup
func (m MTL4TileRenderPipelineDescriptor) RequiredThreadsPerThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](m.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return MTLSize(rv)
}
func (m MTL4TileRenderPipelineDescriptor) SetRequiredThreadsPerThreadgroup(value MTLSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}
// Configures an object that contains information about functions to link to
// the tile render pipeline when Metal builds it.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/staticLinkingDescriptor
func (m MTL4TileRenderPipelineDescriptor) StaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("staticLinkingDescriptor"))
	return MTL4StaticLinkingDescriptorFromID(objc.ID(rv))
}
func (m MTL4TileRenderPipelineDescriptor) SetStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setStaticLinkingDescriptor:"), value)
}
// Indicates whether the pipeline supports linking binary functions.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/supportBinaryLinking
func (m MTL4TileRenderPipelineDescriptor) SupportBinaryLinking() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportBinaryLinking"))
	return rv
}
func (m MTL4TileRenderPipelineDescriptor) SetSupportBinaryLinking(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportBinaryLinking:"), value)
}
// Indicating whether the size of the threadgroup matches the size of a tile
// in the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/threadgroupSizeMatchesTileSize
func (m MTL4TileRenderPipelineDescriptor) ThreadgroupSizeMatchesTileSize() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("threadgroupSizeMatchesTileSize"))
	return rv
}
func (m MTL4TileRenderPipelineDescriptor) SetThreadgroupSizeMatchesTileSize(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setThreadgroupSizeMatchesTileSize:"), value)
}
// Configures the tile function that the render pipeline executes for each
// tile in the tile shader stage.
//
// See: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/tileFunctionDescriptor
func (m MTL4TileRenderPipelineDescriptor) TileFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("tileFunctionDescriptor"))
	return MTL4FunctionDescriptorFromID(objc.ID(rv))
}
func (m MTL4TileRenderPipelineDescriptor) SetTileFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setTileFunctionDescriptor:"), value)
}


// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTextureDescriptor] class.
var (
	_MTLTextureDescriptorClass     MTLTextureDescriptorClass
	_MTLTextureDescriptorClassOnce sync.Once
)

func getMTLTextureDescriptorClass() MTLTextureDescriptorClass {
	_MTLTextureDescriptorClassOnce.Do(func() {
		_MTLTextureDescriptorClass = MTLTextureDescriptorClass{class: objc.GetClass("MTLTextureDescriptor")}
	})
	return _MTLTextureDescriptorClass
}

// GetMTLTextureDescriptorClass returns the class object for MTLTextureDescriptor.
func GetMTLTextureDescriptorClass() MTLTextureDescriptorClass {
	return getMTLTextureDescriptorClass()
}

type MTLTextureDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLTextureDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTextureDescriptorClass) Alloc() MTLTextureDescriptor {
	rv := objc.Send[MTLTextureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An instance that you use to configure new Metal texture instances.
//
// # Overview
// 
// To create a new texture, first create an [MTLTextureDescriptor] instance
// and set its property values. Then, call either the
// [NewTextureWithDescriptor] or [NewTextureWithDescriptorIosurfacePlane]
// method of an [MTLDevice] instance, or the
// [NewTextureWithDescriptorOffsetBytesPerRow] method of an [MTLBuffer]
// instance.
// 
// When you create a texture, Metal copies property values from the descriptor
// into the new texture. You can reuse an [MTLTextureDescriptor] instance,
// modifying its property values as needed, to create more [MTLTexture]
// instances, without affecting any textures you already created.
//
// # Specifying texture attributes
//
//   - [MTLTextureDescriptor.TextureType]: The dimension and arrangement of texture image data.
//   - [MTLTextureDescriptor.SetTextureType]
//   - [MTLTextureDescriptor.PixelFormat]: The size and bit layout of all pixels in the texture.
//   - [MTLTextureDescriptor.SetPixelFormat]
//   - [MTLTextureDescriptor.Width]: The width of the texture image for the base level mipmap, in pixels.
//   - [MTLTextureDescriptor.SetWidth]
//   - [MTLTextureDescriptor.Height]: The height of the texture image for the base level mipmap, in pixels.
//   - [MTLTextureDescriptor.SetHeight]
//   - [MTLTextureDescriptor.Depth]: The depth of the texture image for the base level mipmap, in pixels.
//   - [MTLTextureDescriptor.SetDepth]
//   - [MTLTextureDescriptor.MipmapLevelCount]: The number of mipmap levels for this texture.
//   - [MTLTextureDescriptor.SetMipmapLevelCount]
//   - [MTLTextureDescriptor.SampleCount]: The number of samples in each fragment.
//   - [MTLTextureDescriptor.SetSampleCount]
//   - [MTLTextureDescriptor.ArrayLength]: The number of array elements for this texture.
//   - [MTLTextureDescriptor.SetArrayLength]
//   - [MTLTextureDescriptor.ResourceOptions]: The behavior of a new memory allocation.
//   - [MTLTextureDescriptor.SetResourceOptions]
//   - [MTLTextureDescriptor.CpuCacheMode]: The CPU cache mode used for the CPU mapping of the texture.
//   - [MTLTextureDescriptor.SetCpuCacheMode]
//   - [MTLTextureDescriptor.StorageMode]: The location and access permissions of the texture.
//   - [MTLTextureDescriptor.SetStorageMode]
//   - [MTLTextureDescriptor.HazardTrackingMode]: The texture’s hazard tracking mode.
//   - [MTLTextureDescriptor.SetHazardTrackingMode]
//   - [MTLTextureDescriptor.AllowGPUOptimizedContents]: A Boolean value indicating whether the GPU is allowed to adjust the texture’s contents to improve GPU performance.
//   - [MTLTextureDescriptor.SetAllowGPUOptimizedContents]
//   - [MTLTextureDescriptor.Usage]: Options that determine how you can use the texture.
//   - [MTLTextureDescriptor.SetUsage]
//   - [MTLTextureDescriptor.Swizzle]: The pattern you want the GPU to apply to pixels when you read or sample pixels from the texture.
//   - [MTLTextureDescriptor.SetSwizzle]
//
// # Instance Properties
//
//   - [MTLTextureDescriptor.CompressionType]
//   - [MTLTextureDescriptor.SetCompressionType]
//   - [MTLTextureDescriptor.PlacementSparsePageSize]: Determines the page size for a placement sparse texture.
//   - [MTLTextureDescriptor.SetPlacementSparsePageSize]
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor
type MTLTextureDescriptor struct {
	objectivec.Object
}

// MTLTextureDescriptorFromID constructs a [MTLTextureDescriptor] from an objc.ID.
//
// An instance that you use to configure new Metal texture instances.
func MTLTextureDescriptorFromID(id objc.ID) MTLTextureDescriptor {
	return MTLTextureDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLTextureDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTextureDescriptor] class.
//
// # Specifying texture attributes
//
//   - [IMTLTextureDescriptor.TextureType]: The dimension and arrangement of texture image data.
//   - [IMTLTextureDescriptor.SetTextureType]
//   - [IMTLTextureDescriptor.PixelFormat]: The size and bit layout of all pixels in the texture.
//   - [IMTLTextureDescriptor.SetPixelFormat]
//   - [IMTLTextureDescriptor.Width]: The width of the texture image for the base level mipmap, in pixels.
//   - [IMTLTextureDescriptor.SetWidth]
//   - [IMTLTextureDescriptor.Height]: The height of the texture image for the base level mipmap, in pixels.
//   - [IMTLTextureDescriptor.SetHeight]
//   - [IMTLTextureDescriptor.Depth]: The depth of the texture image for the base level mipmap, in pixels.
//   - [IMTLTextureDescriptor.SetDepth]
//   - [IMTLTextureDescriptor.MipmapLevelCount]: The number of mipmap levels for this texture.
//   - [IMTLTextureDescriptor.SetMipmapLevelCount]
//   - [IMTLTextureDescriptor.SampleCount]: The number of samples in each fragment.
//   - [IMTLTextureDescriptor.SetSampleCount]
//   - [IMTLTextureDescriptor.ArrayLength]: The number of array elements for this texture.
//   - [IMTLTextureDescriptor.SetArrayLength]
//   - [IMTLTextureDescriptor.ResourceOptions]: The behavior of a new memory allocation.
//   - [IMTLTextureDescriptor.SetResourceOptions]
//   - [IMTLTextureDescriptor.CpuCacheMode]: The CPU cache mode used for the CPU mapping of the texture.
//   - [IMTLTextureDescriptor.SetCpuCacheMode]
//   - [IMTLTextureDescriptor.StorageMode]: The location and access permissions of the texture.
//   - [IMTLTextureDescriptor.SetStorageMode]
//   - [IMTLTextureDescriptor.HazardTrackingMode]: The texture’s hazard tracking mode.
//   - [IMTLTextureDescriptor.SetHazardTrackingMode]
//   - [IMTLTextureDescriptor.AllowGPUOptimizedContents]: A Boolean value indicating whether the GPU is allowed to adjust the texture’s contents to improve GPU performance.
//   - [IMTLTextureDescriptor.SetAllowGPUOptimizedContents]
//   - [IMTLTextureDescriptor.Usage]: Options that determine how you can use the texture.
//   - [IMTLTextureDescriptor.SetUsage]
//   - [IMTLTextureDescriptor.Swizzle]: The pattern you want the GPU to apply to pixels when you read or sample pixels from the texture.
//   - [IMTLTextureDescriptor.SetSwizzle]
//
// # Instance Properties
//
//   - [IMTLTextureDescriptor.CompressionType]
//   - [IMTLTextureDescriptor.SetCompressionType]
//   - [IMTLTextureDescriptor.PlacementSparsePageSize]: Determines the page size for a placement sparse texture.
//   - [IMTLTextureDescriptor.SetPlacementSparsePageSize]
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor
type IMTLTextureDescriptor interface {
	objectivec.IObject

	// Topic: Specifying texture attributes

	// The dimension and arrangement of texture image data.
	TextureType() MTLTextureType
	SetTextureType(value MTLTextureType)
	// The size and bit layout of all pixels in the texture.
	PixelFormat() MTLPixelFormat
	SetPixelFormat(value MTLPixelFormat)
	// The width of the texture image for the base level mipmap, in pixels.
	Width() uint
	SetWidth(value uint)
	// The height of the texture image for the base level mipmap, in pixels.
	Height() uint
	SetHeight(value uint)
	// The depth of the texture image for the base level mipmap, in pixels.
	Depth() uint
	SetDepth(value uint)
	// The number of mipmap levels for this texture.
	MipmapLevelCount() uint
	SetMipmapLevelCount(value uint)
	// The number of samples in each fragment.
	SampleCount() uint
	SetSampleCount(value uint)
	// The number of array elements for this texture.
	ArrayLength() uint
	SetArrayLength(value uint)
	// The behavior of a new memory allocation.
	ResourceOptions() MTLResourceOptions
	SetResourceOptions(value MTLResourceOptions)
	// The CPU cache mode used for the CPU mapping of the texture.
	CpuCacheMode() MTLCPUCacheMode
	SetCpuCacheMode(value MTLCPUCacheMode)
	// The location and access permissions of the texture.
	StorageMode() MTLStorageMode
	SetStorageMode(value MTLStorageMode)
	// The texture’s hazard tracking mode.
	HazardTrackingMode() MTLHazardTrackingMode
	SetHazardTrackingMode(value MTLHazardTrackingMode)
	// A Boolean value indicating whether the GPU is allowed to adjust the texture’s contents to improve GPU performance.
	AllowGPUOptimizedContents() bool
	SetAllowGPUOptimizedContents(value bool)
	// Options that determine how you can use the texture.
	Usage() MTLTextureUsage
	SetUsage(value MTLTextureUsage)
	// The pattern you want the GPU to apply to pixels when you read or sample pixels from the texture.
	Swizzle() MTLTextureSwizzleChannels
	SetSwizzle(value MTLTextureSwizzleChannels)

	// Topic: Instance Properties

	CompressionType() MTLTextureCompressionType
	SetCompressionType(value MTLTextureCompressionType)
	// Determines the page size for a placement sparse texture.
	PlacementSparsePageSize() MTLSparsePageSize
	SetPlacementSparsePageSize(value MTLSparsePageSize)
}

// Init initializes the instance.
func (t MTLTextureDescriptor) Init() MTLTextureDescriptor {
	rv := objc.Send[MTLTextureDescriptor](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTextureDescriptor) Autorelease() MTLTextureDescriptor {
	rv := objc.Send[MTLTextureDescriptor](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTextureDescriptor creates a new MTLTextureDescriptor instance.
func NewMTLTextureDescriptor() MTLTextureDescriptor {
	class := getMTLTextureDescriptorClass()
	rv := objc.Send[MTLTextureDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a texture descriptor object for a 2D texture.
//
// pixelFormat: The format describing how every pixel on the texture image is stored. The
// default value is [PixelFormatRGBA8Unorm].
//
// width: The width of the 2D texture image. The value needs to be greater than or
// equal to `1`.
//
// height: The height of the 2D texture image. The value needs to be greater than or
// equal to `1`.
//
// mipmapped: A Boolean indicating whether the resulting image should be mipmapped. If
// [true], then the [MipmapLevelCount] property in the returned descriptor is
// computed from `width` and `height`. If [false], then [MipmapLevelCount] is
// `1`.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A pointer to a texture descriptor object for a 2D texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/texture2DDescriptor(pixelFormat:width:height:mipmapped:)
func (_MTLTextureDescriptorClass MTLTextureDescriptorClass) Texture2DDescriptorWithPixelFormatWidthHeightMipmapped(pixelFormat MTLPixelFormat, width uint, height uint, mipmapped bool) MTLTextureDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLTextureDescriptorClass.class), objc.Sel("texture2DDescriptorWithPixelFormat:width:height:mipmapped:"), pixelFormat, width, height, mipmapped)
	return MTLTextureDescriptorFromID(rv)
}
// Creates a texture descriptor object for a cube texture.
//
// pixelFormat: The format describing how every pixel on the texture image is stored. The
// default value is [PixelFormatRGBA8Unorm].
//
// size: The width and height of each slice of the cube texture. The value needs to
// be greater than or equal to `1`.
//
// mipmapped: A Boolean indicating whether the resulting image should be mipmapped. If
// [true], then the [MipmapLevelCount] property in the returned descriptor is
// computed from `width` and `height`. If [false], then [MipmapLevelCount] is
// `1`.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A pointer to a texture descriptor object for a cube texture.
//
// # Discussion
// 
// For a cube texture, the property values describe one slice, which is any
// one of its six sides. Each slice is a square.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureCubeDescriptor(pixelFormat:size:mipmapped:)
func (_MTLTextureDescriptorClass MTLTextureDescriptorClass) TextureCubeDescriptorWithPixelFormatSizeMipmapped(pixelFormat MTLPixelFormat, size uint, mipmapped bool) MTLTextureDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLTextureDescriptorClass.class), objc.Sel("textureCubeDescriptorWithPixelFormat:size:mipmapped:"), pixelFormat, size, mipmapped)
	return MTLTextureDescriptorFromID(rv)
}
// Creates a texture descriptor object for a texture buffer.
//
// pixelFormat: The format describing how every pixel on the texture buffer is stored. The
// default value is [PixelFormatRGBA8Unorm].
//
// width: The width of the texture buffer. The value needs to be greater than or
// equal to `1`.
//
// resourceOptions: The access options to use for the new texture buffer.
//
// usage: The allowed usage of the new texture buffer.
//
// # Return Value
// 
// A pointer to a texture descriptor object for a texture buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureBufferDescriptor(with:width:resourceOptions:usage:)
func (_MTLTextureDescriptorClass MTLTextureDescriptorClass) TextureBufferDescriptorWithPixelFormatWidthResourceOptionsUsage(pixelFormat MTLPixelFormat, width uint, resourceOptions MTLResourceOptions, usage MTLTextureUsage) MTLTextureDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLTextureDescriptorClass.class), objc.Sel("textureBufferDescriptorWithPixelFormat:width:resourceOptions:usage:"), pixelFormat, width, resourceOptions, usage)
	return MTLTextureDescriptorFromID(rv)
}

// The dimension and arrangement of texture image data.
//
// # Discussion
// 
// The default value is [MTLTexture2D].
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/textureType
func (t MTLTextureDescriptor) TextureType() MTLTextureType {
	rv := objc.Send[MTLTextureType](t.ID, objc.Sel("textureType"))
	return MTLTextureType(rv)
}
func (t MTLTextureDescriptor) SetTextureType(value MTLTextureType) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextureType:"), value)
}
// The size and bit layout of all pixels in the texture.
//
// # Discussion
// 
// The default value is [MTLPixelFormatRGBA8Unorm].
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/pixelFormat
func (t MTLTextureDescriptor) PixelFormat() MTLPixelFormat {
	rv := objc.Send[MTLPixelFormat](t.ID, objc.Sel("pixelFormat"))
	return MTLPixelFormat(rv)
}
func (t MTLTextureDescriptor) SetPixelFormat(value MTLPixelFormat) {
	objc.Send[struct{}](t.ID, objc.Sel("setPixelFormat:"), value)
}
// The width of the texture image for the base level mipmap, in pixels.
//
// # Discussion
// 
// The default value is `1`. The value needs to be greater than or equal to
// `1`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/width
func (t MTLTextureDescriptor) Width() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("width"))
	return rv
}
func (t MTLTextureDescriptor) SetWidth(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setWidth:"), value)
}
// The height of the texture image for the base level mipmap, in pixels.
//
// # Discussion
// 
// The default value is `1`. The value needs to be greater than or equal to
// `1`. For a 1D texture, the value needs to be `1`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/height
func (t MTLTextureDescriptor) Height() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("height"))
	return rv
}
func (t MTLTextureDescriptor) SetHeight(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setHeight:"), value)
}
// The depth of the texture image for the base level mipmap, in pixels.
//
// # Discussion
// 
// The default value is `1`. The value needs to be greater than or equal to
// `1`. For 1D, 2D, and cube textures, the value needs to be `1`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/depth
func (t MTLTextureDescriptor) Depth() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("depth"))
	return rv
}
func (t MTLTextureDescriptor) SetDepth(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setDepth:"), value)
}
// The number of mipmap levels for this texture.
//
// # Discussion
// 
// The default value is `1`. For a buffer-backed or multisample textures, the
// value needs to be `1`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/mipmapLevelCount
func (t MTLTextureDescriptor) MipmapLevelCount() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("mipmapLevelCount"))
	return rv
}
func (t MTLTextureDescriptor) SetMipmapLevelCount(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setMipmapLevelCount:"), value)
}
// The number of samples in each fragment.
//
// # Discussion
// 
// The default value is `1`. If [TextureType] is not
// [TextureType2DMultisample] or [TextureType2DMultisampleArray], this value
// needs to be `1`.
// 
// Support for different sample count values varies by device. Call the
// [SupportsTextureSampleCount] method to determine if your desired sample
// count value is supported.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/sampleCount
func (t MTLTextureDescriptor) SampleCount() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("sampleCount"))
	return rv
}
func (t MTLTextureDescriptor) SetSampleCount(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setSampleCount:"), value)
}
// The number of array elements for this texture.
//
// # Discussion
// 
// The value of this property needs to be between `1` and `2048`, inclusive.
// The default value is `1`.
// 
// This value is `1` if the texture type is not an array.
// 
// This value can be between `1` and `2048` if the texture type is one of the
// following array types:
// 
// - [TextureType1DArray] - [TextureType2DArray] -
// [TextureType2DMultisampleArray] - [TextureTypeCubeArray]
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/arrayLength
func (t MTLTextureDescriptor) ArrayLength() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("arrayLength"))
	return rv
}
func (t MTLTextureDescriptor) SetArrayLength(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setArrayLength:"), value)
}
// The behavior of a new memory allocation.
//
// # Discussion
// 
// This property only has an effect when you are allocating a new texture. If
// you are creating a texture whose data comes from another [MTLResource]
// object, this property value is ignored, and the value of the original
// resource is used instead.
// 
// The value of this property aggregates the values of [StorageMode],
// [CpuCacheMode], and [HazardTrackingMode]. If you modify this property, the
// other properties also change, and vice versa.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/resourceOptions
func (t MTLTextureDescriptor) ResourceOptions() MTLResourceOptions {
	rv := objc.Send[MTLResourceOptions](t.ID, objc.Sel("resourceOptions"))
	return MTLResourceOptions(rv)
}
func (t MTLTextureDescriptor) SetResourceOptions(value MTLResourceOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setResourceOptions:"), value)
}
// The CPU cache mode used for the CPU mapping of the texture.
//
// # Discussion
// 
// The default value is [CPUCacheModeDefaultCache].
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/cpuCacheMode
func (t MTLTextureDescriptor) CpuCacheMode() MTLCPUCacheMode {
	rv := objc.Send[MTLCPUCacheMode](t.ID, objc.Sel("cpuCacheMode"))
	return MTLCPUCacheMode(rv)
}
func (t MTLTextureDescriptor) SetCpuCacheMode(value MTLCPUCacheMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setCpuCacheMode:"), value)
}
// The location and access permissions of the texture.
//
// # Discussion
// 
// In iOS and tvOS, the default value is [StorageModeShared]. In macOS, the
// default value is [StorageModeManaged].
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/storageMode
func (t MTLTextureDescriptor) StorageMode() MTLStorageMode {
	rv := objc.Send[MTLStorageMode](t.ID, objc.Sel("storageMode"))
	return MTLStorageMode(rv)
}
func (t MTLTextureDescriptor) SetStorageMode(value MTLStorageMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setStorageMode:"), value)
}
// The texture’s hazard tracking mode.
//
// # Discussion
// 
// The default value is [HazardTrackingModeDefault].
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/hazardTrackingMode
func (t MTLTextureDescriptor) HazardTrackingMode() MTLHazardTrackingMode {
	rv := objc.Send[MTLHazardTrackingMode](t.ID, objc.Sel("hazardTrackingMode"))
	return MTLHazardTrackingMode(rv)
}
func (t MTLTextureDescriptor) SetHazardTrackingMode(value MTLHazardTrackingMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setHazardTrackingMode:"), value)
}
// A Boolean value indicating whether the GPU is allowed to adjust the
// texture’s contents to improve GPU performance.
//
// # Discussion
// 
// The default value is `true`, which means that the Metal device is allowed
// to adjust the private layout of the texture in memory to improve GPU
// performance. For a shared or managed texture, this optimization can cause
// slower performance when accessing the texture from the CPU. Setting this
// property to `false` improves CPU performance at the cost of some GPU
// performance.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/allowGPUOptimizedContents
func (t MTLTextureDescriptor) AllowGPUOptimizedContents() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowGPUOptimizedContents"))
	return rv
}
func (t MTLTextureDescriptor) SetAllowGPUOptimizedContents(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowGPUOptimizedContents:"), value)
}
// Options that determine how you can use the texture.
//
// # Discussion
// 
// The default value for this property is [TextureUsageShaderRead]. If the
// given texture has multiple uses in your app, you can combine multiple usage
// options for that texture. After you set a texture’s usage options, you
// can use it only in the ways that you specified.
// 
// Metal can optimize operations for a given texture, based on its intended
// use. Set explicit usage options for a texture, if you know them in advance,
// before you use the texture. Only set usage options that correspond to a
// texture’s intended use.
// 
// In iOS devices with GPU family 5, Metal doesn’t apply lossless
// compression to a given texture if you set any of these options:
// 
// - [TextureUsageUnknown] - [TextureUsageShaderWrite] -
// [TextureUsagePixelFormatView]
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/usage
func (t MTLTextureDescriptor) Usage() MTLTextureUsage {
	rv := objc.Send[MTLTextureUsage](t.ID, objc.Sel("usage"))
	return MTLTextureUsage(rv)
}
func (t MTLTextureDescriptor) SetUsage(value MTLTextureUsage) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsage:"), value)
}
// The pattern you want the GPU to apply to pixels when you read or sample
// pixels from the texture.
//
// # Discussion
// 
// The default value does not apply a transformation to pixels sampled or read
// from the texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/swizzle
func (t MTLTextureDescriptor) Swizzle() MTLTextureSwizzleChannels {
	rv := objc.Send[MTLTextureSwizzleChannels](t.ID, objc.Sel("swizzle"))
	return MTLTextureSwizzleChannels(rv)
}
func (t MTLTextureDescriptor) SetSwizzle(value MTLTextureSwizzleChannels) {
	objc.Send[struct{}](t.ID, objc.Sel("setSwizzle:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/compressionType
func (t MTLTextureDescriptor) CompressionType() MTLTextureCompressionType {
	rv := objc.Send[MTLTextureCompressionType](t.ID, objc.Sel("compressionType"))
	return MTLTextureCompressionType(rv)
}
func (t MTLTextureDescriptor) SetCompressionType(value MTLTextureCompressionType) {
	objc.Send[struct{}](t.ID, objc.Sel("setCompressionType:"), value)
}
// Determines the page size for a placement sparse texture.
//
// # Discussion
// 
// Set this property to a non-zero value to create a .
// 
// Placement sparse textures are instances of [MTLTexture] that you assign
// memory to using a [MTLHeap] instance of type [HeapTypePlacement] and a
// [MaxCompatiblePlacementSparsePageSize] at least as large as the
// [MTLSparsePageSize] value you assign to this property.
// 
// This value is 0 by default.
//
// [MTLSparsePageSize]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureDescriptor/placementSparsePageSize
func (t MTLTextureDescriptor) PlacementSparsePageSize() MTLSparsePageSize {
	rv := objc.Send[MTLSparsePageSize](t.ID, objc.Sel("placementSparsePageSize"))
	return MTLSparsePageSize(rv)
}
func (t MTLTextureDescriptor) SetPlacementSparsePageSize(value MTLSparsePageSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setPlacementSparsePageSize:"), value)
}


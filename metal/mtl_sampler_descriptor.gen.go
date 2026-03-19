// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLSamplerDescriptor] class.
var (
	_MTLSamplerDescriptorClass     MTLSamplerDescriptorClass
	_MTLSamplerDescriptorClassOnce sync.Once
)

func getMTLSamplerDescriptorClass() MTLSamplerDescriptorClass {
	_MTLSamplerDescriptorClassOnce.Do(func() {
		_MTLSamplerDescriptorClass = MTLSamplerDescriptorClass{class: objc.GetClass("MTLSamplerDescriptor")}
	})
	return _MTLSamplerDescriptorClass
}

// GetMTLSamplerDescriptorClass returns the class object for MTLSamplerDescriptor.
func GetMTLSamplerDescriptorClass() MTLSamplerDescriptorClass {
	return getMTLSamplerDescriptorClass()
}

type MTLSamplerDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLSamplerDescriptorClass) Alloc() MTLSamplerDescriptor {
	rv := objc.Send[MTLSamplerDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that you use to configure a texture sampler.
//
// # Overview
// 
// To make a sampler, create and configure an [MTLSamplerDescriptor] instance
// and then call an [MTLDevice] instance’s [NewSamplerStateWithDescriptor]
// method. After you create the sampler, you can release the descriptor or
// reconfigure its properties to create other samplers.
//
// # Declaring the coordinate space
//
//   - [MTLSamplerDescriptor.NormalizedCoordinates]: A Boolean value that indicates whether texture coordinates are normalized to the range `[0.0, 1.0]`.
//   - [MTLSamplerDescriptor.SetNormalizedCoordinates]
//
// # Declaring addressing modes
//
//   - [MTLSamplerDescriptor.RAddressMode]: The address mode for the texture depth (r) coordinate.
//   - [MTLSamplerDescriptor.SetRAddressMode]
//   - [MTLSamplerDescriptor.SAddressMode]: The address mode for the texture width (s) coordinate.
//   - [MTLSamplerDescriptor.SetSAddressMode]
//   - [MTLSamplerDescriptor.TAddressMode]: The address mode for the texture height (t) coordinate.
//   - [MTLSamplerDescriptor.SetTAddressMode]
//   - [MTLSamplerDescriptor.BorderColor]: The border color for clamped texture values.
//   - [MTLSamplerDescriptor.SetBorderColor]
//
// # Declaring filter modes
//
//   - [MTLSamplerDescriptor.MinFilter]: The filtering option for combining pixels within one mipmap level when the sample footprint is larger than a pixel (minification).
//   - [MTLSamplerDescriptor.SetMinFilter]
//   - [MTLSamplerDescriptor.MagFilter]: The filtering operation for combining pixels within one mipmap level when the sample footprint is smaller than a pixel (magnification).
//   - [MTLSamplerDescriptor.SetMagFilter]
//   - [MTLSamplerDescriptor.MipFilter]: The filtering option for combining pixels between two mipmap levels.
//   - [MTLSamplerDescriptor.SetMipFilter]
//   - [MTLSamplerDescriptor.LodMinClamp]: The minimum level of detail (LOD) to use when sampling from a texture.
//   - [MTLSamplerDescriptor.SetLodMinClamp]
//   - [MTLSamplerDescriptor.LodMaxClamp]: The maximum level of detail (LOD) to use when sampling from a texture.
//   - [MTLSamplerDescriptor.SetLodMaxClamp]
//   - [MTLSamplerDescriptor.LodAverage]: A Boolean value that specifies whether the GPU can use an average level of detail (LOD) when sampling from a texture.
//   - [MTLSamplerDescriptor.SetLodAverage]
//   - [MTLSamplerDescriptor.MaxAnisotropy]: The number of samples that can be taken to improve the quality of sample footprints that are anisotropic.
//   - [MTLSamplerDescriptor.SetMaxAnisotropy]
//
// # Declaring the depth comparison mode
//
//   - [MTLSamplerDescriptor.CompareFunction]: The sampler comparison function used when performing a sample compare operation on a depth texture.
//   - [MTLSamplerDescriptor.SetCompareFunction]
//
// # Declaring whether the sampler can be used in argument buffers
//
//   - [MTLSamplerDescriptor.SupportArgumentBuffers]: A Boolean value that indicates whether you can reference a sampler, that you make with this descriptor, by its resource ID from an argument buffer.
//   - [MTLSamplerDescriptor.SetSupportArgumentBuffers]
//
// # Identifying the sampler
//
//   - [MTLSamplerDescriptor.Label]: A string that identifies the sampler.
//   - [MTLSamplerDescriptor.SetLabel]
//
// # Instance Properties
//
//   - [MTLSamplerDescriptor.LodBias]: Sets the level-of-detail (lod) bias when sampling from a texture.
//   - [MTLSamplerDescriptor.SetLodBias]
//   - [MTLSamplerDescriptor.ReductionMode]: Sets the reduction mode for filtering contributing samples.
//   - [MTLSamplerDescriptor.SetReductionMode]
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor
type MTLSamplerDescriptor struct {
	objectivec.Object
}

// MTLSamplerDescriptorFromID constructs a [MTLSamplerDescriptor] from an objc.ID.
//
// An object that you use to configure a texture sampler.
func MTLSamplerDescriptorFromID(id objc.ID) MTLSamplerDescriptor {
	return MTLSamplerDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLSamplerDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLSamplerDescriptor] class.
//
// # Declaring the coordinate space
//
//   - [IMTLSamplerDescriptor.NormalizedCoordinates]: A Boolean value that indicates whether texture coordinates are normalized to the range `[0.0, 1.0]`.
//   - [IMTLSamplerDescriptor.SetNormalizedCoordinates]
//
// # Declaring addressing modes
//
//   - [IMTLSamplerDescriptor.RAddressMode]: The address mode for the texture depth (r) coordinate.
//   - [IMTLSamplerDescriptor.SetRAddressMode]
//   - [IMTLSamplerDescriptor.SAddressMode]: The address mode for the texture width (s) coordinate.
//   - [IMTLSamplerDescriptor.SetSAddressMode]
//   - [IMTLSamplerDescriptor.TAddressMode]: The address mode for the texture height (t) coordinate.
//   - [IMTLSamplerDescriptor.SetTAddressMode]
//   - [IMTLSamplerDescriptor.BorderColor]: The border color for clamped texture values.
//   - [IMTLSamplerDescriptor.SetBorderColor]
//
// # Declaring filter modes
//
//   - [IMTLSamplerDescriptor.MinFilter]: The filtering option for combining pixels within one mipmap level when the sample footprint is larger than a pixel (minification).
//   - [IMTLSamplerDescriptor.SetMinFilter]
//   - [IMTLSamplerDescriptor.MagFilter]: The filtering operation for combining pixels within one mipmap level when the sample footprint is smaller than a pixel (magnification).
//   - [IMTLSamplerDescriptor.SetMagFilter]
//   - [IMTLSamplerDescriptor.MipFilter]: The filtering option for combining pixels between two mipmap levels.
//   - [IMTLSamplerDescriptor.SetMipFilter]
//   - [IMTLSamplerDescriptor.LodMinClamp]: The minimum level of detail (LOD) to use when sampling from a texture.
//   - [IMTLSamplerDescriptor.SetLodMinClamp]
//   - [IMTLSamplerDescriptor.LodMaxClamp]: The maximum level of detail (LOD) to use when sampling from a texture.
//   - [IMTLSamplerDescriptor.SetLodMaxClamp]
//   - [IMTLSamplerDescriptor.LodAverage]: A Boolean value that specifies whether the GPU can use an average level of detail (LOD) when sampling from a texture.
//   - [IMTLSamplerDescriptor.SetLodAverage]
//   - [IMTLSamplerDescriptor.MaxAnisotropy]: The number of samples that can be taken to improve the quality of sample footprints that are anisotropic.
//   - [IMTLSamplerDescriptor.SetMaxAnisotropy]
//
// # Declaring the depth comparison mode
//
//   - [IMTLSamplerDescriptor.CompareFunction]: The sampler comparison function used when performing a sample compare operation on a depth texture.
//   - [IMTLSamplerDescriptor.SetCompareFunction]
//
// # Declaring whether the sampler can be used in argument buffers
//
//   - [IMTLSamplerDescriptor.SupportArgumentBuffers]: A Boolean value that indicates whether you can reference a sampler, that you make with this descriptor, by its resource ID from an argument buffer.
//   - [IMTLSamplerDescriptor.SetSupportArgumentBuffers]
//
// # Identifying the sampler
//
//   - [IMTLSamplerDescriptor.Label]: A string that identifies the sampler.
//   - [IMTLSamplerDescriptor.SetLabel]
//
// # Instance Properties
//
//   - [IMTLSamplerDescriptor.LodBias]: Sets the level-of-detail (lod) bias when sampling from a texture.
//   - [IMTLSamplerDescriptor.SetLodBias]
//   - [IMTLSamplerDescriptor.ReductionMode]: Sets the reduction mode for filtering contributing samples.
//   - [IMTLSamplerDescriptor.SetReductionMode]
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor
type IMTLSamplerDescriptor interface {
	objectivec.IObject

	// Topic: Declaring the coordinate space

	// A Boolean value that indicates whether texture coordinates are normalized to the range `[0.0, 1.0]`.
	NormalizedCoordinates() bool
	SetNormalizedCoordinates(value bool)

	// Topic: Declaring addressing modes

	// The address mode for the texture depth (r) coordinate.
	RAddressMode() MTLSamplerAddressMode
	SetRAddressMode(value MTLSamplerAddressMode)
	// The address mode for the texture width (s) coordinate.
	SAddressMode() MTLSamplerAddressMode
	SetSAddressMode(value MTLSamplerAddressMode)
	// The address mode for the texture height (t) coordinate.
	TAddressMode() MTLSamplerAddressMode
	SetTAddressMode(value MTLSamplerAddressMode)
	// The border color for clamped texture values.
	BorderColor() MTLSamplerBorderColor
	SetBorderColor(value MTLSamplerBorderColor)

	// Topic: Declaring filter modes

	// The filtering option for combining pixels within one mipmap level when the sample footprint is larger than a pixel (minification).
	MinFilter() MTLSamplerMinMagFilter
	SetMinFilter(value MTLSamplerMinMagFilter)
	// The filtering operation for combining pixels within one mipmap level when the sample footprint is smaller than a pixel (magnification).
	MagFilter() MTLSamplerMinMagFilter
	SetMagFilter(value MTLSamplerMinMagFilter)
	// The filtering option for combining pixels between two mipmap levels.
	MipFilter() MTLSamplerMipFilter
	SetMipFilter(value MTLSamplerMipFilter)
	// The minimum level of detail (LOD) to use when sampling from a texture.
	LodMinClamp() float32
	SetLodMinClamp(value float32)
	// The maximum level of detail (LOD) to use when sampling from a texture.
	LodMaxClamp() float32
	SetLodMaxClamp(value float32)
	// A Boolean value that specifies whether the GPU can use an average level of detail (LOD) when sampling from a texture.
	LodAverage() bool
	SetLodAverage(value bool)
	// The number of samples that can be taken to improve the quality of sample footprints that are anisotropic.
	MaxAnisotropy() uint
	SetMaxAnisotropy(value uint)

	// Topic: Declaring the depth comparison mode

	// The sampler comparison function used when performing a sample compare operation on a depth texture.
	CompareFunction() MTLCompareFunction
	SetCompareFunction(value MTLCompareFunction)

	// Topic: Declaring whether the sampler can be used in argument buffers

	// A Boolean value that indicates whether you can reference a sampler, that you make with this descriptor, by its resource ID from an argument buffer.
	SupportArgumentBuffers() bool
	SetSupportArgumentBuffers(value bool)

	// Topic: Identifying the sampler

	// A string that identifies the sampler.
	Label() string
	SetLabel(value string)

	// Topic: Instance Properties

	// Sets the level-of-detail (lod) bias when sampling from a texture.
	LodBias() float32
	SetLodBias(value float32)
	// Sets the reduction mode for filtering contributing samples.
	ReductionMode() MTLSamplerReductionMode
	SetReductionMode(value MTLSamplerReductionMode)
}

// Init initializes the instance.
func (s MTLSamplerDescriptor) Init() MTLSamplerDescriptor {
	rv := objc.Send[MTLSamplerDescriptor](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MTLSamplerDescriptor) Autorelease() MTLSamplerDescriptor {
	rv := objc.Send[MTLSamplerDescriptor](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLSamplerDescriptor creates a new MTLSamplerDescriptor instance.
func NewMTLSamplerDescriptor() MTLSamplerDescriptor {
	class := getMTLSamplerDescriptorClass()
	rv := objc.Send[MTLSamplerDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether texture coordinates are normalized
// to the range `[0.0, 1.0]`.
//
// # Discussion
// 
// If [true], texture coordinates are from `0.0` to `1.0`. If [false], texture
// coordinates are from `0` to `width` for horizontal coordinates and `0` to
// `height` for vertical coordinates. The default value is [true].
// 
// Non-normalized texture coordinates should only be used with 1D and 2D
// textures with the following conditions; otherwise, the results of sampling
// are undefined.
// 
// - The [SamplerAddressModeClampToEdge] or [SamplerAddressModeClampToZero]
// address mode. - The [SamplerMipFilterNotMipmapped] mipmap filtering option.
// - [MinFilter] and [MagFilter] need to be equal to each other. -
// [MaxAnisotropy] needs to be `1`.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/normalizedCoordinates
func (s MTLSamplerDescriptor) NormalizedCoordinates() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("normalizedCoordinates"))
	return rv
}
func (s MTLSamplerDescriptor) SetNormalizedCoordinates(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setNormalizedCoordinates:"), value)
}
// The address mode for the texture depth (r) coordinate.
//
// # Discussion
// 
// The default value is [SamplerAddressModeClampToEdge].
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/rAddressMode
func (s MTLSamplerDescriptor) RAddressMode() MTLSamplerAddressMode {
	rv := objc.Send[MTLSamplerAddressMode](s.ID, objc.Sel("rAddressMode"))
	return MTLSamplerAddressMode(rv)
}
func (s MTLSamplerDescriptor) SetRAddressMode(value MTLSamplerAddressMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setRAddressMode:"), value)
}
// The address mode for the texture width (s) coordinate.
//
// # Discussion
// 
// The default value is [SamplerAddressModeClampToEdge].
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/sAddressMode
func (s MTLSamplerDescriptor) SAddressMode() MTLSamplerAddressMode {
	rv := objc.Send[MTLSamplerAddressMode](s.ID, objc.Sel("sAddressMode"))
	return MTLSamplerAddressMode(rv)
}
func (s MTLSamplerDescriptor) SetSAddressMode(value MTLSamplerAddressMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setSAddressMode:"), value)
}
// The address mode for the texture height (t) coordinate.
//
// # Discussion
// 
// The default value is [SamplerAddressModeClampToEdge].
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/tAddressMode
func (s MTLSamplerDescriptor) TAddressMode() MTLSamplerAddressMode {
	rv := objc.Send[MTLSamplerAddressMode](s.ID, objc.Sel("tAddressMode"))
	return MTLSamplerAddressMode(rv)
}
func (s MTLSamplerDescriptor) SetTAddressMode(value MTLSamplerAddressMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setTAddressMode:"), value)
}
// The border color for clamped texture values.
//
// # Discussion
// 
// This value is only used when the sampler address mode is
// [SamplerAddressModeClampToBorderColor].
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/borderColor
func (s MTLSamplerDescriptor) BorderColor() MTLSamplerBorderColor {
	rv := objc.Send[MTLSamplerBorderColor](s.ID, objc.Sel("borderColor"))
	return MTLSamplerBorderColor(rv)
}
func (s MTLSamplerDescriptor) SetBorderColor(value MTLSamplerBorderColor) {
	objc.Send[struct{}](s.ID, objc.Sel("setBorderColor:"), value)
}
// The filtering option for combining pixels within one mipmap level when the
// sample footprint is larger than a pixel (minification).
//
// # Discussion
// 
// The default value is [SamplerMinMagFilterNearest].
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/minFilter
func (s MTLSamplerDescriptor) MinFilter() MTLSamplerMinMagFilter {
	rv := objc.Send[MTLSamplerMinMagFilter](s.ID, objc.Sel("minFilter"))
	return MTLSamplerMinMagFilter(rv)
}
func (s MTLSamplerDescriptor) SetMinFilter(value MTLSamplerMinMagFilter) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinFilter:"), value)
}
// The filtering operation for combining pixels within one mipmap level when
// the sample footprint is smaller than a pixel (magnification).
//
// # Discussion
// 
// The default value is [SamplerMinMagFilterNearest].
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/magFilter
func (s MTLSamplerDescriptor) MagFilter() MTLSamplerMinMagFilter {
	rv := objc.Send[MTLSamplerMinMagFilter](s.ID, objc.Sel("magFilter"))
	return MTLSamplerMinMagFilter(rv)
}
func (s MTLSamplerDescriptor) SetMagFilter(value MTLSamplerMinMagFilter) {
	objc.Send[struct{}](s.ID, objc.Sel("setMagFilter:"), value)
}
// The filtering option for combining pixels between two mipmap levels.
//
// # Discussion
// 
// The default value is [SamplerMipFilterNotMipmapped].
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/mipFilter
func (s MTLSamplerDescriptor) MipFilter() MTLSamplerMipFilter {
	rv := objc.Send[MTLSamplerMipFilter](s.ID, objc.Sel("mipFilter"))
	return MTLSamplerMipFilter(rv)
}
func (s MTLSamplerDescriptor) SetMipFilter(value MTLSamplerMipFilter) {
	objc.Send[struct{}](s.ID, objc.Sel("setMipFilter:"), value)
}
// The minimum level of detail (LOD) to use when sampling from a texture.
//
// # Discussion
// 
// The default value is `0.0`. Clamp values are always applied, even when
// using an explicit LOD.
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodMinClamp
func (s MTLSamplerDescriptor) LodMinClamp() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("lodMinClamp"))
	return rv
}
func (s MTLSamplerDescriptor) SetLodMinClamp(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setLodMinClamp:"), value)
}
// The maximum level of detail (LOD) to use when sampling from a texture.
//
// # Discussion
// 
// The default value is `FLT_MAX`. Clamp values are always applied, even when
// using an explicit LOD.
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodMaxClamp
func (s MTLSamplerDescriptor) LodMaxClamp() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("lodMaxClamp"))
	return rv
}
func (s MTLSamplerDescriptor) SetLodMaxClamp(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setLodMaxClamp:"), value)
}
// A Boolean value that specifies whether the GPU can use an average level of
// detail (LOD) when sampling from a texture.
//
// # Discussion
// 
// If this value is [true], an average LOD may be used across four fragment
// shader threads. If this value is [false], no averaging is performed and
// each thread accesses its own LOD.
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodAverage
func (s MTLSamplerDescriptor) LodAverage() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("lodAverage"))
	return rv
}
func (s MTLSamplerDescriptor) SetLodAverage(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setLodAverage:"), value)
}
// The number of samples that can be taken to improve the quality of sample
// footprints that are anisotropic.
//
// # Discussion
// 
// Values need to be between `1` and `16`, inclusive. The default value is
// `1`.
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/maxAnisotropy
func (s MTLSamplerDescriptor) MaxAnisotropy() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("maxAnisotropy"))
	return rv
}
func (s MTLSamplerDescriptor) SetMaxAnisotropy(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaxAnisotropy:"), value)
}
// The sampler comparison function used when performing a sample compare
// operation on a depth texture.
//
// # Discussion
// 
// The default value is [CompareFunctionNever].
// 
// The [FeatureSet_iOS_GPUFamily3_v1] and [FeatureSet_iOS_GPUFamily1_v1]
// feature sets allow you to define a framework-side sampler comparison
// function for an [MTLSamplerState] instance. All feature sets support
// shader-side sampler comparison functions, as described in the [Metal
// Shading Language Specification].
//
// [Metal Shading Language Specification]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/compareFunction
func (s MTLSamplerDescriptor) CompareFunction() MTLCompareFunction {
	rv := objc.Send[MTLCompareFunction](s.ID, objc.Sel("compareFunction"))
	return MTLCompareFunction(rv)
}
func (s MTLSamplerDescriptor) SetCompareFunction(value MTLCompareFunction) {
	objc.Send[struct{}](s.ID, objc.Sel("setCompareFunction:"), value)
}
// A Boolean value that indicates whether you can reference a sampler, that
// you make with this descriptor, by its resource ID from an argument buffer.
//
// # Discussion
// 
// The default value is [false], which means that you can only encode the
// samplers you make with this descriptor as individual resources in the
// sampler state argument table.
// 
// Your app can encode samplers into an argument buffer if you create them
// with an [MTLSamplerDescriptor] instance that has this property equal to
// [true].
// 
// Each unique configuration of an [MTLSamplerDescriptor] instance’s
// properties creates a unique [MTLSamplerState] instance. For example, you
// can create unique samplers with the same [MTLSamplerDescriptor] instance by
// changing one or more values of its properties, such as [MinFilter] or
// [MagFilter] before creating another instance.
// 
// Conversely, creating secondary sampler instances with the same descriptor
// property values doesn’t create any additional, unique samplers. Instead,
// they refer to the same underlying sampler, even if you create it with a
// difference descriptor instance because the configuration is the same.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/supportArgumentBuffers
func (s MTLSamplerDescriptor) SupportArgumentBuffers() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("supportArgumentBuffers"))
	return rv
}
func (s MTLSamplerDescriptor) SetSupportArgumentBuffers(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSupportArgumentBuffers:"), value)
}
// A string that identifies the sampler.
//
// # Discussion
// 
// Object and command labels are useful identifiers at runtime or when
// profiling and debugging your app using any Metal tool. See [Naming
// resources and commands].
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/label
func (s MTLSamplerDescriptor) Label() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (s MTLSamplerDescriptor) SetLabel(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setLabel:"), objc.String(value))
}
// Sets the level-of-detail (lod) bias when sampling from a texture.
//
// # Discussion
// 
// The property’s default value is `0.0f`. The precision format is `S4.6`,
// and the range is `[-16.0, 15.999]`.
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/lodBias
func (s MTLSamplerDescriptor) LodBias() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("lodBias"))
	return rv
}
func (s MTLSamplerDescriptor) SetLodBias(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setLodBias:"), value)
}
// Sets the reduction mode for filtering contributing samples.
//
// # Discussion
// 
// The property’s default value is [MTLSamplerReductionModeWeightedAverage].
// The sampler ignores this property if any of the following property values
// are equal to a specific value:
// 
// - The sampler’s [MipFilter] property is equal to
// [MTLSamplerMipFilterNotMipmapped]. - The sampler’s [MipFilter] property
// is equal to [MTLSamplerMipFilterNearest]. - The sampler’s [MinFilter]
// property is equal to [MTLSamplerMinMagFilterNearest]. - The sampler’s
// [MagFilter] property is equal to [MTLSamplerMinMagFilterNearest].
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerDescriptor/reductionMode
func (s MTLSamplerDescriptor) ReductionMode() MTLSamplerReductionMode {
	rv := objc.Send[MTLSamplerReductionMode](s.ID, objc.Sel("reductionMode"))
	return MTLSamplerReductionMode(rv)
}
func (s MTLSamplerDescriptor) SetReductionMode(value MTLSamplerReductionMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setReductionMode:"), value)
}


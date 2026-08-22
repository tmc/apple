// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSImageDescriptor] class.
var (
	_MPSImageDescriptorClass     MPSImageDescriptorClass
	_MPSImageDescriptorClassOnce sync.Once
)

func getMPSImageDescriptorClass() MPSImageDescriptorClass {
	_MPSImageDescriptorClassOnce.Do(func() {
		_MPSImageDescriptorClass = MPSImageDescriptorClass{class: objc.GetClass("MPSImageDescriptor")}
	})
	return _MPSImageDescriptorClass
}

// GetMPSImageDescriptorClass returns the class object for MPSImageDescriptor.
func GetMPSImageDescriptorClass() MPSImageDescriptorClass {
	return getMPSImageDescriptorClass()
}

type MPSImageDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageDescriptorClass) Alloc() MPSImageDescriptor {
	rv := objc.Send[MPSImageDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of the attributes used to create an [MPSImage].
//
// # Overview
//
// You use an [MPSImageDescriptor] to describe and create the properties of an
// [MPSImage] such as its size, pixel format and CPU cache mode.
//
// # Properties
//
//   - [MPSImageDescriptor.Width]: The width of the image.
//   - [MPSImageDescriptor.SetWidth]
//   - [MPSImageDescriptor.Height]: The height of the image.
//   - [MPSImageDescriptor.SetHeight]
//   - [MPSImageDescriptor.FeatureChannels]: The number of feature channels per pixel.
//   - [MPSImageDescriptor.SetFeatureChannels]
//   - [MPSImageDescriptor.NumberOfImages]: The number of images for batch processing.
//   - [MPSImageDescriptor.SetNumberOfImages]
//   - [MPSImageDescriptor.PixelFormat]: The pixel format for the underlying texture.
//   - [MPSImageDescriptor.ChannelFormat]: The storage format to use for each channel in the image.
//   - [MPSImageDescriptor.SetChannelFormat]
//   - [MPSImageDescriptor.CpuCacheMode]: The CPU cache mode of the underlying texture.
//   - [MPSImageDescriptor.SetCpuCacheMode]
//   - [MPSImageDescriptor.StorageMode]: The storage mode of underlying texture.
//   - [MPSImageDescriptor.SetStorageMode]
//   - [MPSImageDescriptor.Usage]: Options to specify the intended usage of the underlying texture.
//   - [MPSImageDescriptor.SetUsage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor
type MPSImageDescriptor struct {
	objectivec.Object
}

// MPSImageDescriptorFromID constructs a [MPSImageDescriptor] from an objc.ID.
//
// A description of the attributes used to create an [MPSImage].
func MPSImageDescriptorFromID(id objc.ID) MPSImageDescriptor {
	return MPSImageDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSImageDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageDescriptor] class.
//
// # Properties
//
//   - [IMPSImageDescriptor.Width]: The width of the image.
//   - [IMPSImageDescriptor.SetWidth]
//   - [IMPSImageDescriptor.Height]: The height of the image.
//   - [IMPSImageDescriptor.SetHeight]
//   - [IMPSImageDescriptor.FeatureChannels]: The number of feature channels per pixel.
//   - [IMPSImageDescriptor.SetFeatureChannels]
//   - [IMPSImageDescriptor.NumberOfImages]: The number of images for batch processing.
//   - [IMPSImageDescriptor.SetNumberOfImages]
//   - [IMPSImageDescriptor.PixelFormat]: The pixel format for the underlying texture.
//   - [IMPSImageDescriptor.ChannelFormat]: The storage format to use for each channel in the image.
//   - [IMPSImageDescriptor.SetChannelFormat]
//   - [IMPSImageDescriptor.CpuCacheMode]: The CPU cache mode of the underlying texture.
//   - [IMPSImageDescriptor.SetCpuCacheMode]
//   - [IMPSImageDescriptor.StorageMode]: The storage mode of underlying texture.
//   - [IMPSImageDescriptor.SetStorageMode]
//   - [IMPSImageDescriptor.Usage]: Options to specify the intended usage of the underlying texture.
//   - [IMPSImageDescriptor.SetUsage]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor
type IMPSImageDescriptor interface {
	objectivec.IObject

	// Topic: Properties

	// The width of the image.
	Width() uint
	SetWidth(value uint)
	// The height of the image.
	Height() uint
	SetHeight(value uint)
	// The number of feature channels per pixel.
	FeatureChannels() uint
	SetFeatureChannels(value uint)
	// The number of images for batch processing.
	NumberOfImages() uint
	SetNumberOfImages(value uint)
	// The pixel format for the underlying texture.
	PixelFormat() metal.MTLPixelFormat
	// The storage format to use for each channel in the image.
	ChannelFormat() MPSImageFeatureChannelFormat
	SetChannelFormat(value MPSImageFeatureChannelFormat)
	// The CPU cache mode of the underlying texture.
	CpuCacheMode() metal.MTLCPUCacheMode
	SetCpuCacheMode(value metal.MTLCPUCacheMode)
	// The storage mode of underlying texture.
	StorageMode() metal.MTLStorageMode
	SetStorageMode(value metal.MTLStorageMode)
	// Options to specify the intended usage of the underlying texture.
	Usage() metal.MTLTextureUsage
	SetUsage(value metal.MTLTextureUsage)
}

// Init initializes the instance.
func (i MPSImageDescriptor) Init() MPSImageDescriptor {
	rv := objc.Send[MPSImageDescriptor](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageDescriptor) Autorelease() MPSImageDescriptor {
	rv := objc.Send[MPSImageDescriptor](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageDescriptor creates a new MPSImageDescriptor instance.
func NewMPSImageDescriptor() MPSImageDescriptor {
	class := getMPSImageDescriptorClass()
	rv := objc.Send[MPSImageDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an image descriptor for a single image.
//
// channelFormat: The storage format to use for each channel in the image.
//
// width: The width of the image.
//
// height: The height of the image.
//
// featureChannels: The number of feature channels per pixel.
//
// # Return Value
//
// A valid [MPSImageDescriptor] object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/init(channelFormat:width:height:featureChannels:)
func NewImageDescriptorWithChannelFormatWidthHeightFeatureChannels(channelFormat MPSImageFeatureChannelFormat, width uint, height uint, featureChannels uint) MPSImageDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSImageDescriptorClass().class), objc.Sel("imageDescriptorWithChannelFormat:width:height:featureChannels:"), channelFormat, width, height, featureChannels)
	return MPSImageDescriptorFromID(rv)
}

// Creates an image descriptor for an image container with options to set
// texture usage and batch size (number of images).
//
// channelFormat: The storage format to use for each channel in the image.
//
// width: The width of the image.
//
// height: The height of the image.
//
// featureChannels: The number of feature channels per pixel.
//
// numberOfImages: The number of images for batch processing.
//
// usage: The intended usage of the underlying texture.
//
// # Return Value
//
// A valid [MPSImageDescriptor] object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/init(channelFormat:width:height:featureChannels:numberOfImages:usage:)
func NewImageDescriptorWithChannelFormatWidthHeightFeatureChannelsNumberOfImagesUsage(channelFormat MPSImageFeatureChannelFormat, width uint, height uint, featureChannels uint, numberOfImages uint, usage metal.MTLTextureUsage) MPSImageDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSImageDescriptorClass().class), objc.Sel("imageDescriptorWithChannelFormat:width:height:featureChannels:numberOfImages:usage:"), channelFormat, width, height, featureChannels, numberOfImages, usage)
	return MPSImageDescriptorFromID(rv)
}

// The width of the image.
//
// # Discussion
//
// The formal width of the image, in pixels. The default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/width
func (i MPSImageDescriptor) Width() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("width"))
	return rv
}
func (i MPSImageDescriptor) SetWidth(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setWidth:"), value)
}

// The height of the image.
//
// # Discussion
//
// The formal height of the image, in pixels. The default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/height
func (i MPSImageDescriptor) Height() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("height"))
	return rv
}
func (i MPSImageDescriptor) SetHeight(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setHeight:"), value)
}

// The number of feature channels per pixel.
//
// # Discussion
//
// The default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/featureChannels
func (i MPSImageDescriptor) FeatureChannels() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("featureChannels"))
	return rv
}
func (i MPSImageDescriptor) SetFeatureChannels(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setFeatureChannels:"), value)
}

// The number of images for batch processing.
//
// # Discussion
//
// The default value is 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/numberOfImages
func (i MPSImageDescriptor) NumberOfImages() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("numberOfImages"))
	return rv
}
func (i MPSImageDescriptor) SetNumberOfImages(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setNumberOfImages:"), value)
}

// The pixel format for the underlying texture.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/pixelFormat
func (i MPSImageDescriptor) PixelFormat() metal.MTLPixelFormat {
	rv := objc.Send[metal.MTLPixelFormat](i.ID, objc.Sel("pixelFormat"))
	return metal.MTLPixelFormat(rv)
}

// The storage format to use for each channel in the image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/channelFormat
func (i MPSImageDescriptor) ChannelFormat() MPSImageFeatureChannelFormat {
	rv := objc.Send[MPSImageFeatureChannelFormat](i.ID, objc.Sel("channelFormat"))
	return MPSImageFeatureChannelFormat(rv)
}
func (i MPSImageDescriptor) SetChannelFormat(value MPSImageFeatureChannelFormat) {
	objc.Send[struct{}](i.ID, objc.Sel("setChannelFormat:"), value)
}

// The CPU cache mode of the underlying texture.
//
// # Discussion
//
// The default value is [MTLCPUCacheMode.defaultCache].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/cpuCacheMode
//
// [MTLCPUCacheMode.defaultCache]: https://developer.apple.com/documentation/Metal/MTLCPUCacheMode/defaultCache
func (i MPSImageDescriptor) CpuCacheMode() metal.MTLCPUCacheMode {
	rv := objc.Send[metal.MTLCPUCacheMode](i.ID, objc.Sel("cpuCacheMode"))
	return metal.MTLCPUCacheMode(rv)
}
func (i MPSImageDescriptor) SetCpuCacheMode(value metal.MTLCPUCacheMode) {
	objc.Send[struct{}](i.ID, objc.Sel("setCpuCacheMode:"), value)
}

// The storage mode of underlying texture.
//
// # Discussion
//
// The default value is [MTLStorageMode.shared].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/storageMode
//
// [MTLStorageMode.shared]: https://developer.apple.com/documentation/Metal/MTLStorageMode/shared
func (i MPSImageDescriptor) StorageMode() metal.MTLStorageMode {
	rv := objc.Send[metal.MTLStorageMode](i.ID, objc.Sel("storageMode"))
	return metal.MTLStorageMode(rv)
}
func (i MPSImageDescriptor) SetStorageMode(value metal.MTLStorageMode) {
	objc.Send[struct{}](i.ID, objc.Sel("setStorageMode:"), value)
}

// Options to specify the intended usage of the underlying texture.
//
// # Discussion
//
// The default value is [shaderRead]`|`[shaderWrite].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDescriptor/usage
//
// [shaderRead]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderRead
// [shaderWrite]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderWrite
func (i MPSImageDescriptor) Usage() metal.MTLTextureUsage {
	rv := objc.Send[metal.MTLTextureUsage](i.ID, objc.Sel("usage"))
	return metal.MTLTextureUsage(rv)
}
func (i MPSImageDescriptor) SetUsage(value metal.MTLTextureUsage) {
	objc.Send[struct{}](i.ID, objc.Sel("setUsage:"), value)
}

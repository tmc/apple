// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSImage] class.
var (
	_MPSImageClass     MPSImageClass
	_MPSImageClassOnce sync.Once
)

func getMPSImageClass() MPSImageClass {
	_MPSImageClassOnce.Do(func() {
		_MPSImageClass = MPSImageClass{class: objc.GetClass("MPSImage")}
	})
	return _MPSImageClass
}

// GetMPSImageClass returns the class object for MPSImage.
func GetMPSImageClass() MPSImageClass {
	return getMPSImageClass()
}

type MPSImageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageClass) Alloc() MPSImage {
	rv := objc.Send[MPSImage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A texture that may have more than four channels for use in convolutional
// neural networks.
//
// # Overview
//
// Some image types, such as those found in convolutional neural networks
// (CNN), differ from a standard texture in that they may have more than 4
// channels per pixel. While the channels could hold RGBA data, they will more
// commonly hold a number of structural permutations upon an RGBA image as the
// neural network progresses. It is not uncommon for each pixel to have 32 or
// 64 channels in it.
//
// Since a standard [MTLTexture] object cannot have more than 4 channels, the
// additional channels are stored in slices of a 2D texture array (i.e. a
// texture of type [MTLTextureType.type2DArray]) such that 4 consecutive
// channels are stored in each slice of this array. If the number of feature
// channels is [N], the number of array slices needed is `(N+3)/4`. For
// example, a 9-channel CNN image with a width of 3 and a height of 2 will be
// stored as follows:
//
// [media-2556907]
//
// Thus, the width and height of the underlying 2D texture array is the same
// as the width and height of the [MPSImage] object and the array length is
// equal to `(` [MPSImageDescriptor.FeatureChannels] `+3)/4`. (Channels marked
// with a `?` are just for padding and should not contain [NaN] or [INF]
// values.)
//
// An [MPSImage] object can contain multiple CNN images for batch processing.
// In order to create an [MPSImage] object that contains [N] images, create an
// [MPSImageDescriptor] object with the [MPSImageDescriptor.NumberOfImages]
// property set to [N]. The length of the 2D texture array (i.e. the number of
// slices) will be equal to `((` [MPSImageDescriptor.FeatureChannels]
// `+3)/4)*` [MPSImageDescriptor.NumberOfImages], where consecutive
// `(featureChannels+3)/4` slices of this array represent one image.
//
// Although an [MPSImage] object can contain more than one image, the actual
// number of images among these processed by an [MPSCNNKernel] object is
// controlled by the `z` dimension of the [MPSCNNKernel.ClipRect] property. (A
// kernel processes `n=clipRect.SizeXCUIElementTypeDepth()` images from this
// collection.)
//
// The starting index of the image to process from the source [MPSImage]
// object is given by `offset.Z()`. The starting index of the image in the
// destination [MPSImage] object where this processed image is written to is
// given by `clipRect.OriginXCUIElementTypeZ()`. Thus, an [MPSCNNKernel]
// object takes the `n=clipRect.SizeXCUIElementTypeDepth()` image from the
// source at indices `[offset.Z(), offset.Z()+n]`, processes each
// independently, and stores the result in the destination at indices
// `[clipRect.OriginXCUIElementTypeZ(), clipRect.OriginXCUIElementTypeZ()+n]`
// respectively. Thus, `offset.Z()+n` should be `=0`.
//
// For example, suppose an [MPSCNNConvolution] object takes an input image
// with 16 channels and outputs an image with 32 channels. The number of
// slices needed in the source 2D texture array is 4 and the number of slices
// needed in the destination 2D texture array is 8. Suppose the source batch
// size is 5 and the destination batch size is 4. Thus, the number of source
// slices will be `4*5=20` and the number of destination slices will be
// `8*4=32`. If you want to process image 2 and 3 of the source and store the
// result at index 1 and 2 in the destination, you can achieve this by setting
// `offset.Z()=2`, `clipRect.OriginXCUIElementTypeZ()=1`, and
// `clipRect.SizeXCUIElementTypeDepth()=2`. The [MPSCNNConvolution] object
// will take, in this case, slices 4 and 5 of the source and produce slices 4
// to 7 of the destination. Similarly, slices 6 and 7 will be used to produce
// slices 8 to 11 of the destination.
//
// All [MPSCNNKernel] objects process images in the batch independently. That
// is, calling a [MPSCNNKernel] object on a batch is formally the same as
// calling it on each image in the batch sequentially. Computational and GPU
// work submission overhead will be amortized over more work if batch
// processing is used. This is especially important for better performance on
// small images.
//
// If `featureChannels<=4` and `numberOfImages=1` (i.e. only one slice is
// needed to represent the image), the underlying metal texture type is chosen
// to be [MTLTextureType.type2D] rather than [MTLTextureType.type2DArray] as
// explained above.
//
// The framework also provides [MPSTemporaryImage] objects, intended for very
// short-lived image data that is produced and consumed immediately in the
// same [MTLCommandBuffer] object. They are a useful way to minimize CPU-side
// texture allocation costs and greatly reduce the amount of memory used by
// your image pipeline.
//
// Creation of the underlying texture may occur lazily in some cases. In
// general, you should avoid calling the [MPSImage.Texture] property to avoid
// materializing memory for longer than necessary. When possible, use the
// other [MPSImage] properties to get information about the object instead.
//
// # The MPSImage Class
//
// [MTLBuffer] and [MTLTexture] objects are commonly used in Metal apps and
// are used directly by the Metal Performance Shaders framework when possible.
// In apps that use CNN, kernels may need more than the four data channels
// that a [MTLTexture] object can provide. In these cases, an [MPSImage]
// object is used instead as an abstraction layer on top of a [MTLTexture]
// object. When more than 4 channels are needed, additional textures in the 2D
// texture array are added to hold additional channels in sets of four. An
// [MPSImage] object tracks this information as the number of feature channels
// in an image.
//
// # CNN Images
//
// [MPSCNNKernel] objects operate on [MPSImage] objects. [MPSImage] objects
// are at their core [MTLTexture] objects; however, whereas [MTLTexture]
// objects commonly represent image or texel data, an [MPSImage] object is a
// more abstract representation of image features. The channels within an
// [MPSImage] do not necessarily correspond to colors in a color space
// (although they can, if necessary). As a result, there can be many more than
// four of them. Having 32 or 64 channels per pixel is not uncommon in CNN.
// This is achieved on the [MTLTexture] object abstraction by inserting extra
// RGBA pixels to handle the additional feature channels (if any) beyond 4.
// These extra pixels are stored as multiple slices of a 2D image array. Thus,
// each CNN pixel in a 32-channel image is represented as 8 array slices, with
// 4-channels stored per-pixel in each slice. The width and height of the
// [MTLTexture] object is the same as the width and height of the [MPSImage]
// object. The number of slices in the [MTLTexture] object is given by the
// number of feature channels rounded up to a multiple of 4.
//
// [MPSImage] objects can be created from existing [MTLTexture] objects. They
// may also be created anew from an [MPSImageDescriptor] and backed with
// either standard texture memory, or as [MPSTemporaryImage] objects using
// memory drawn from the framework’s internal cached texture backing store.
// [MPSTemporaryImage] objects can provide great memory usage and CPU time
// savings, but come with significant restrictions that should be understood
// before using them. For example, their contents are only valid during the
// GPU-side execution of a single [MTLCommandBuffer] object and can not be
// read from or written to by the CPU. They are provided as an efficient way
// to hold CNN computations that are used immediately within the scope of the
// same [MTLCommandBuffer] object and then discarded. Concatenation is also
// supported by allowing you to define from which destination feature channel
// to start writing the output of the current layer. In this way, your app can
// make a large [MPSImage] or [MPSTemporaryImage] object and fill in parts of
// it with multiple layers (as long as the destination feature channel offset
// is a multiple of 4).
//
// # Supported Pixel Formats
//
// The following table shows pixel formats supported by [MPSImage].
//
// [Table data omitted]
//
// # Initializers
//
//   - [MPSImage.InitWithDeviceImageDescriptor]: Initializes an empty image.
//   - [MPSImage.InitWithTextureFeatureChannels]: Initializes an image from a texture. The user-allocated texture has been created for a specific number of feature channels and number of images.
//   - [MPSImage.InitWithParentImageSliceRangeFeatureChannels]
//
// # Methods
//
//   - [MPSImage.SetPurgeableState]: Set (or query) the purgeable state of the image’s underlying texture.
//
// # Methods to Read and Write Raw Data
//
//   - [MPSImage.ReadBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex]
//   - [MPSImage.ReadBytesDataLayoutImageIndex]
//   - [MPSImage.WriteBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex]
//   - [MPSImage.WriteBytesDataLayoutImageIndex]
//
// # Properties
//
//   - [MPSImage.Device]: The device on which the image will be used.
//   - [MPSImage.Width]: The formal width of the image, in pixels.
//   - [MPSImage.Height]: The formal height of the image, in pixels.
//   - [MPSImage.FeatureChannels]: The number of feature channels per pixel.
//   - [MPSImage.NumberOfImages]: The number of images for batch processing.
//   - [MPSImage.TextureType]: The type of the underlying texture.
//   - [MPSImage.PixelFormat]: The pixel format of the underlying texture.
//   - [MPSImage.Precision]: The number of bits of numeric precision available for each feature channel.
//   - [MPSImage.Usage]: The intended usage of the underlying texture.
//   - [MPSImage.PixelSize]: The number of bytes from the first byte of one pixel to the first byte of the next pixel, in storage order. (Includes padding.)
//   - [MPSImage.Texture]: The underlying texture.
//   - [MPSImage.Label]: A string to help identify this object.
//   - [MPSImage.SetLabel]
//
// # Instance Properties
//
//   - [MPSImage.FeatureChannelFormat]
//   - [MPSImage.Parent]
//
// # Instance Methods
//
//   - [MPSImage.BatchRepresentation]
//   - [MPSImage.BatchRepresentationWithSubRange]
//   - [MPSImage.ReadBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex]
//   - [MPSImage.ResourceSize]
//   - [MPSImage.SubImageWithFeatureChannelRange]
//   - [MPSImage.SynchronizeOnCommandBuffer]
//   - [MPSImage.WriteBytesDataLayoutBytesPerColumnBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex]
//   - [MPSImage.WriteBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage
//
// [MTLBuffer]: https://developer.apple.com/documentation/Metal/MTLBuffer
// [MTLCommandBuffer]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer
// [MTLTextureType.type2DArray]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DArray
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
type MPSImage struct {
	objectivec.Object
}

// MPSImageFromID constructs a [MPSImage] from an objc.ID.
//
// A texture that may have more than four channels for use in convolutional
// neural networks.
func MPSImageFromID(id objc.ID) MPSImage {
	return MPSImage{objectivec.Object{ID: id}}
}

// NOTE: MPSImage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImage] class.
//
// # Initializers
//
//   - [IMPSImage.InitWithDeviceImageDescriptor]: Initializes an empty image.
//   - [IMPSImage.InitWithTextureFeatureChannels]: Initializes an image from a texture. The user-allocated texture has been created for a specific number of feature channels and number of images.
//   - [IMPSImage.InitWithParentImageSliceRangeFeatureChannels]
//
// # Methods
//
//   - [IMPSImage.SetPurgeableState]: Set (or query) the purgeable state of the image’s underlying texture.
//
// # Methods to Read and Write Raw Data
//
//   - [IMPSImage.ReadBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex]
//   - [IMPSImage.ReadBytesDataLayoutImageIndex]
//   - [IMPSImage.WriteBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex]
//   - [IMPSImage.WriteBytesDataLayoutImageIndex]
//
// # Properties
//
//   - [IMPSImage.Device]: The device on which the image will be used.
//   - [IMPSImage.Width]: The formal width of the image, in pixels.
//   - [IMPSImage.Height]: The formal height of the image, in pixels.
//   - [IMPSImage.FeatureChannels]: The number of feature channels per pixel.
//   - [IMPSImage.NumberOfImages]: The number of images for batch processing.
//   - [IMPSImage.TextureType]: The type of the underlying texture.
//   - [IMPSImage.PixelFormat]: The pixel format of the underlying texture.
//   - [IMPSImage.Precision]: The number of bits of numeric precision available for each feature channel.
//   - [IMPSImage.Usage]: The intended usage of the underlying texture.
//   - [IMPSImage.PixelSize]: The number of bytes from the first byte of one pixel to the first byte of the next pixel, in storage order. (Includes padding.)
//   - [IMPSImage.Texture]: The underlying texture.
//   - [IMPSImage.Label]: A string to help identify this object.
//   - [IMPSImage.SetLabel]
//
// # Instance Properties
//
//   - [IMPSImage.FeatureChannelFormat]
//   - [IMPSImage.Parent]
//
// # Instance Methods
//
//   - [IMPSImage.BatchRepresentation]
//   - [IMPSImage.BatchRepresentationWithSubRange]
//   - [IMPSImage.ReadBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex]
//   - [IMPSImage.ResourceSize]
//   - [IMPSImage.SubImageWithFeatureChannelRange]
//   - [IMPSImage.SynchronizeOnCommandBuffer]
//   - [IMPSImage.WriteBytesDataLayoutBytesPerColumnBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex]
//   - [IMPSImage.WriteBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage
type IMPSImage interface {
	objectivec.IObject

	// Topic: Initializers

	// Initializes an empty image.
	InitWithDeviceImageDescriptor(device metal.MTLDevice, imageDescriptor IMPSImageDescriptor) MPSImage
	// Initializes an image from a texture. The user-allocated texture has been created for a specific number of feature channels and number of images.
	InitWithTextureFeatureChannels(texture metal.MTLTexture, featureChannels uint) MPSImage
	InitWithParentImageSliceRangeFeatureChannels(parent IMPSImage, sliceRange foundation.NSRange, featureChannels uint) MPSImage

	// Topic: Methods

	// Set (or query) the purgeable state of the image’s underlying texture.
	SetPurgeableState(state MPSPurgeableState) MPSPurgeableState

	// Topic: Methods to Read and Write Raw Data

	ReadBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, bytesPerRow uint, region metal.MTLRegion, featureChannelInfo MPSImageReadWriteParams, imageIndex uint)
	ReadBytesDataLayoutImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, imageIndex uint)
	WriteBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, bytesPerRow uint, region metal.MTLRegion, featureChannelInfo MPSImageReadWriteParams, imageIndex uint)
	WriteBytesDataLayoutImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, imageIndex uint)

	// Topic: Properties

	// The device on which the image will be used.
	Device() metal.MTLDevice
	// The formal width of the image, in pixels.
	Width() uint
	// The formal height of the image, in pixels.
	Height() uint
	// The number of feature channels per pixel.
	FeatureChannels() uint
	// The number of images for batch processing.
	NumberOfImages() uint
	// The type of the underlying texture.
	TextureType() metal.MTLTextureType
	// The pixel format of the underlying texture.
	PixelFormat() metal.MTLPixelFormat
	// The number of bits of numeric precision available for each feature channel.
	Precision() uint
	// The intended usage of the underlying texture.
	Usage() metal.MTLTextureUsage
	// The number of bytes from the first byte of one pixel to the first byte of the next pixel, in storage order. (Includes padding.)
	PixelSize() uintptr
	// The underlying texture.
	Texture() metal.MTLTexture
	// A string to help identify this object.
	Label() string
	SetLabel(value string)

	// Topic: Instance Properties

	FeatureChannelFormat() MPSImageFeatureChannelFormat
	Parent() IMPSImage

	// Topic: Instance Methods

	BatchRepresentation() MPSImageBatch
	BatchRepresentationWithSubRange(subRange foundation.NSRange) MPSImageBatch
	ReadBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, bytesPerRow uint, bytesPerImage uint, region metal.MTLRegion, featureChannelInfo MPSImageReadWriteParams, imageIndex uint)
	ResourceSize() uint
	SubImageWithFeatureChannelRange(range_ foundation.NSRange) IMPSImage
	SynchronizeOnCommandBuffer(commandBuffer metal.MTLCommandBuffer)
	WriteBytesDataLayoutBytesPerColumnBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, bytesPerColumn uint, bytesPerRow uint, bytesPerImage uint, region metal.MTLRegion, featureChannelInfo MPSImageReadWriteParams, imageIndex uint)
	WriteBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, bytesPerRow uint, bytesPerImage uint, region metal.MTLRegion, featureChannelInfo MPSImageReadWriteParams, imageIndex uint)
}

// Init initializes the instance.
func (i MPSImage) Init() MPSImage {
	rv := objc.Send[MPSImage](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImage) Autorelease() MPSImage {
	rv := objc.Send[MPSImage](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImage creates a new MPSImage instance.
func NewMPSImage() MPSImage {
	class := getMPSImageClass()
	rv := objc.Send[MPSImage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an empty image.
//
// device: The device on which the image will be used.
//
// imageDescriptor: The image descriptor.
//
// # Return Value
//
// A valid [MPSImage] object or `nil`, if failure.
//
// # Discussion
//
// Storage for the image data is allocated lazily on the first use of the
// [MPSImage] object, or when the [MPSImage.Texture] property is first read.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/init(device:imageDescriptor:)
func NewImageWithDeviceImageDescriptor(device metal.MTLDevice, imageDescriptor IMPSImageDescriptor) MPSImage {
	instance := getMPSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:imageDescriptor:"), device, imageDescriptor)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/init(parentImage:sliceRange:featureChannels:)
func NewImageWithParentImageSliceRangeFeatureChannels(parent IMPSImage, sliceRange foundation.NSRange, featureChannels uint) MPSImage {
	instance := getMPSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParentImage:sliceRange:featureChannels:"), parent, sliceRange, featureChannels)
	return MPSImageFromID(rv)
}

// Initializes an image from a texture. The user-allocated texture has been
// created for a specific number of feature channels and number of images.
//
// texture: The texture allocated by the user to be used as a backing storage for the
// image.
//
// featureChannels: The number of feature channels the texture contains.
//
// # Return Value
//
// A valid [MPSImage] object or `nil`, if failure.
//
// # Discussion
//
// In a memory-intensive app, you can save memory (and allocation/deallocation
// time) by using an [MPSTemporaryImage] object, where the framework
// aggressively reuses underlying texture memory within the same command
// buffer. However, in certain cases, you may want more control on the
// allocation, placement, reuse, and recycling of memory-backing textures used
// in your app by using the Metal Resource Heaps API. In this case, an app can
// create an [MPSImage] object from a pre-allocated texture by calling this
// method.
//
// The [textureType] property of the given texture can be of type
// [MTLTextureType.type2D] only if `featureChannels<=4` (meaning that
// `numberOfImages=1`). Otherwise, the texture type should be
// [MTLTextureType.type2DArray] with the [arrayLength] property of the given
// texture being equal to `numberOfImages*((featureChannels+3)/4)`.
//
// For textures containing typical image data, the `featureChannels` parameter
// should be set to the number of valid color channels (e.g. for RGB data,
// even though the pixel format is a form of [MTLPixelFormatRGBA],
// `featureChannels` should be set to 3.).
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/init(texture:featureChannels:)
//
// [MTLTextureType.type2DArray]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DArray
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
// [arrayLength]: https://developer.apple.com/documentation/Metal/MTLTexture/arrayLength
// [textureType]: https://developer.apple.com/documentation/Metal/MTLTexture/textureType
func NewImageWithTextureFeatureChannels(texture metal.MTLTexture, featureChannels uint) MPSImage {
	instance := getMPSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTexture:featureChannels:"), texture, featureChannels)
	return MPSImageFromID(rv)
}

// Initializes an empty image.
//
// device: The device on which the image will be used.
//
// imageDescriptor: The image descriptor.
//
// # Return Value
//
// A valid [MPSImage] object or `nil`, if failure.
//
// # Discussion
//
// Storage for the image data is allocated lazily on the first use of the
// [MPSImage] object, or when the [MPSImage.Texture] property is first read.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/init(device:imageDescriptor:)
func (i MPSImage) InitWithDeviceImageDescriptor(device metal.MTLDevice, imageDescriptor IMPSImageDescriptor) MPSImage {
	rv := objc.Send[MPSImage](i.ID, objc.Sel("initWithDevice:imageDescriptor:"), device, imageDescriptor)
	return rv
}

// Initializes an image from a texture. The user-allocated texture has been
// created for a specific number of feature channels and number of images.
//
// texture: The texture allocated by the user to be used as a backing storage for the
// image.
//
// featureChannels: The number of feature channels the texture contains.
//
// # Return Value
//
// A valid [MPSImage] object or `nil`, if failure.
//
// # Discussion
//
// In a memory-intensive app, you can save memory (and allocation/deallocation
// time) by using an [MPSTemporaryImage] object, where the framework
// aggressively reuses underlying texture memory within the same command
// buffer. However, in certain cases, you may want more control on the
// allocation, placement, reuse, and recycling of memory-backing textures used
// in your app by using the Metal Resource Heaps API. In this case, an app can
// create an [MPSImage] object from a pre-allocated texture by calling this
// method.
//
// The [textureType] property of the given texture can be of type
// [MTLTextureType.type2D] only if `featureChannels<=4` (meaning that
// `numberOfImages=1`). Otherwise, the texture type should be
// [MTLTextureType.type2DArray] with the [arrayLength] property of the given
// texture being equal to `numberOfImages*((featureChannels+3)/4)`.
//
// For textures containing typical image data, the `featureChannels` parameter
// should be set to the number of valid color channels (e.g. for RGB data,
// even though the pixel format is a form of [MTLPixelFormatRGBA],
// `featureChannels` should be set to 3.).
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/init(texture:featureChannels:)
//
// [MTLTextureType.type2DArray]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DArray
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
// [arrayLength]: https://developer.apple.com/documentation/Metal/MTLTexture/arrayLength
// [textureType]: https://developer.apple.com/documentation/Metal/MTLTexture/textureType
func (i MPSImage) InitWithTextureFeatureChannels(texture metal.MTLTexture, featureChannels uint) MPSImage {
	rv := objc.Send[MPSImage](i.ID, objc.Sel("initWithTexture:featureChannels:"), texture, featureChannels)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/init(parentImage:sliceRange:featureChannels:)
func (i MPSImage) InitWithParentImageSliceRangeFeatureChannels(parent IMPSImage, sliceRange foundation.NSRange, featureChannels uint) MPSImage {
	rv := objc.Send[MPSImage](i.ID, objc.Sel("initWithParentImage:sliceRange:featureChannels:"), parent, sliceRange, featureChannels)
	return rv
}

// Set (or query) the purgeable state of the image’s underlying texture.
//
// state: The desired purgeable state of the image’s underlying texture.
//
// # Return Value
//
// Returns the prior purgeable state of the image’s underlying texture.
//
// # Discussion
//
// This method behaves the same as the [setPurgeableState(_:)] method of the
// [MTLResource] class, except that the state might be
// [MPSPurgeableStateAllocationDeferred], which means there is no underlying
// texture to mark as volatile or non-volatile. Attempts to set a purgeable
// state on [MTLTexture] objects that have not yet been allocated will be
// ignored.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/setPurgeableState(_:)
//
// [MTLResource]: https://developer.apple.com/documentation/Metal/MTLResource
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
// [setPurgeableState(_:)]: https://developer.apple.com/documentation/Metal/MTLResource/setPurgeableState(_:)
func (i MPSImage) SetPurgeableState(state MPSPurgeableState) MPSPurgeableState {
	rv := objc.Send[MPSPurgeableState](i.ID, objc.Sel("setPurgeableState:"), state)
	return MPSPurgeableState(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/readBytes(_:dataLayout:bytesPerRow:region:featureChannelInfo:imageIndex:)
func (i MPSImage) ReadBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, bytesPerRow uint, region metal.MTLRegion, featureChannelInfo MPSImageReadWriteParams, imageIndex uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("readBytes:dataLayout:bytesPerRow:region:featureChannelInfo:imageIndex:"), dataBytes, dataLayout, bytesPerRow, region, featureChannelInfo, imageIndex)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/readBytes(_:dataLayout:imageIndex:)
func (i MPSImage) ReadBytesDataLayoutImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, imageIndex uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("readBytes:dataLayout:imageIndex:"), dataBytes, dataLayout, imageIndex)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/writeBytes(_:dataLayout:bytesPerRow:region:featureChannelInfo:imageIndex:)
func (i MPSImage) WriteBytesDataLayoutBytesPerRowRegionFeatureChannelInfoImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, bytesPerRow uint, region metal.MTLRegion, featureChannelInfo MPSImageReadWriteParams, imageIndex uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("writeBytes:dataLayout:bytesPerRow:region:featureChannelInfo:imageIndex:"), dataBytes, dataLayout, bytesPerRow, region, featureChannelInfo, imageIndex)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/writeBytes(_:dataLayout:imageIndex:)
func (i MPSImage) WriteBytesDataLayoutImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, imageIndex uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("writeBytes:dataLayout:imageIndex:"), dataBytes, dataLayout, imageIndex)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/batchRepresentation()
func (i MPSImage) BatchRepresentation() MPSImageBatch {
	rv := objc.Send[MPSImageBatch](i.ID, objc.Sel("batchRepresentation"))
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/batchRepresentation(withSubRange:)
func (i MPSImage) BatchRepresentationWithSubRange(subRange foundation.NSRange) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](i.ID, objc.Sel("batchRepresentationWithSubRange:"), subRange)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/readBytes(_:dataLayout:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:)
func (i MPSImage) ReadBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, bytesPerRow uint, bytesPerImage uint, region metal.MTLRegion, featureChannelInfo MPSImageReadWriteParams, imageIndex uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("readBytes:dataLayout:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:"), dataBytes, dataLayout, bytesPerRow, bytesPerImage, region, featureChannelInfo, imageIndex)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/resourceSize()
func (i MPSImage) ResourceSize() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("resourceSize"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/subImage(withFeatureChannelRange:)
func (i MPSImage) SubImageWithFeatureChannelRange(range_ foundation.NSRange) IMPSImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("subImageWithFeatureChannelRange:"), range_)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/synchronize(on:)
func (i MPSImage) SynchronizeOnCommandBuffer(commandBuffer metal.MTLCommandBuffer) {
	objc.Send[objc.ID](i.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/writeBytes(_:dataLayout:bytesPerColumn:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:)
func (i MPSImage) WriteBytesDataLayoutBytesPerColumnBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, bytesPerColumn uint, bytesPerRow uint, bytesPerImage uint, region metal.MTLRegion, featureChannelInfo MPSImageReadWriteParams, imageIndex uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("writeBytes:dataLayout:bytesPerColumn:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:"), dataBytes, dataLayout, bytesPerColumn, bytesPerRow, bytesPerImage, region, featureChannelInfo, imageIndex)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/writeBytes(_:dataLayout:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:)
func (i MPSImage) WriteBytesDataLayoutBytesPerRowBytesPerImageRegionFeatureChannelInfoImageIndex(dataBytes unsafe.Pointer, dataLayout MPSDataLayout, bytesPerRow uint, bytesPerImage uint, region metal.MTLRegion, featureChannelInfo MPSImageReadWriteParams, imageIndex uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("writeBytes:dataLayout:bytesPerRow:bytesPerImage:region:featureChannelInfo:imageIndex:"), dataBytes, dataLayout, bytesPerRow, bytesPerImage, region, featureChannelInfo, imageIndex)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/defaultAllocator()
func (_MPSImageClass MPSImageClass) DefaultAllocator() MPSImageAllocator {
	rv := objc.Send[objc.ID](objc.ID(_MPSImageClass.class), objc.Sel("defaultAllocator"))
	return MPSImageAllocatorObjectFromID(rv)
}

// The device on which the image will be used.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/device
func (i MPSImage) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}

// The formal width of the image, in pixels.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/width
func (i MPSImage) Width() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("width"))
	return rv
}

// The formal height of the image, in pixels.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/height
func (i MPSImage) Height() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("height"))
	return rv
}

// The number of feature channels per pixel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/featureChannels
func (i MPSImage) FeatureChannels() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("featureChannels"))
	return rv
}

// The number of images for batch processing.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/numberOfImages
func (i MPSImage) NumberOfImages() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("numberOfImages"))
	return rv
}

// The type of the underlying texture.
//
// # Discussion
//
// A property that defines the type of texture the image represents. In most
// cases, this will be [MTLTextureType.type2D] or
// [MTLTextureType.type2DArray].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/textureType
//
// [MTLTextureType.type2DArray]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DArray
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
func (i MPSImage) TextureType() metal.MTLTextureType {
	rv := objc.Send[metal.MTLTextureType](i.ID, objc.Sel("textureType"))
	return metal.MTLTextureType(rv)
}

// The pixel format of the underlying texture.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/pixelFormat
func (i MPSImage) PixelFormat() metal.MTLPixelFormat {
	rv := objc.Send[metal.MTLPixelFormat](i.ID, objc.Sel("pixelFormat"))
	return metal.MTLPixelFormat(rv)
}

// The number of bits of numeric precision available for each feature channel.
//
// # Discussion
//
// This is precision, not size (float is 24 bits, not 32; half-precision
// floating-point is 11 bits, not 16; [Snorm] pixel formats have one less bit
// of precision for the sign bit, etc.). For formats like
// [MTLPixelFormat.b5g6r5Unorm], this value is the precision of the most
// precise channel (which is 6 in this case). When this information is
// unavailable, typically for compressed formats, this value is 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/precision
//
// [MTLPixelFormat.b5g6r5Unorm]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/b5g6r5Unorm
func (i MPSImage) Precision() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("precision"))
	return rv
}

// The intended usage of the underlying texture.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/usage
func (i MPSImage) Usage() metal.MTLTextureUsage {
	rv := objc.Send[metal.MTLTextureUsage](i.ID, objc.Sel("usage"))
	return metal.MTLTextureUsage(rv)
}

// The number of bytes from the first byte of one pixel to the first byte of
// the next pixel, in storage order. (Includes padding.)
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/pixelSize
func (i MPSImage) PixelSize() uintptr {
	rv := objc.Send[uintptr](i.ID, objc.Sel("pixelSize"))
	return rv
}

// The underlying texture.
//
// # Discussion
//
// This is a 2D texture if `numberOfImages=1` and `featureChannels<=4`. It is
// a 2D texture array otherwise.
//
// To avoid the high cost of premature allocation of the underlying texture,
// avoid accessing this property except when strictly necessary. Calls to the
// `encode` methods of an [MPSCNNKernel] object typically cause their
// arguments to become allocated. Likewise, [MPSImage] objects initialized
// with the [MPSTemporaryImage.InitWithTextureFeatureChannels] method have
// already been allocated.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/texture
func (i MPSImage) Texture() metal.MTLTexture {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("texture"))
	return metal.MTLTextureObjectFromID(rv)
}

// A string to help identify this object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/label
func (i MPSImage) Label() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (i MPSImage) SetLabel(value string) {
	objc.Send[struct{}](i.ID, objc.Sel("setLabel:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/featureChannelFormat
func (i MPSImage) FeatureChannelFormat() MPSImageFeatureChannelFormat {
	rv := objc.Send[MPSImageFeatureChannelFormat](i.ID, objc.Sel("featureChannelFormat"))
	return MPSImageFeatureChannelFormat(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/parent
func (i MPSImage) Parent() IMPSImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("parent"))
	return MPSImageFromID(objc.ID(rv))
}

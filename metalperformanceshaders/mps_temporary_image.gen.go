// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSTemporaryImage] class.
var (
	_MPSTemporaryImageClass     MPSTemporaryImageClass
	_MPSTemporaryImageClassOnce sync.Once
)

func getMPSTemporaryImageClass() MPSTemporaryImageClass {
	_MPSTemporaryImageClassOnce.Do(func() {
		_MPSTemporaryImageClass = MPSTemporaryImageClass{class: objc.GetClass("MPSTemporaryImage")}
	})
	return _MPSTemporaryImageClass
}

// GetMPSTemporaryImageClass returns the class object for MPSTemporaryImage.
func GetMPSTemporaryImageClass() MPSTemporaryImageClass {
	return getMPSTemporaryImageClass()
}

type MPSTemporaryImageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSTemporaryImageClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSTemporaryImageClass) Alloc() MPSTemporaryImage {
	rv := objc.Send[MPSTemporaryImage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A texture for use in convolutional neural networks that stores transient
// data to be used and discarded promptly.
//
// # Overview
//
// [MPSTemporaryImage] objects can provide a profound reduction in the
// aggregate texture memory and associated CPU-side allocation cost in your
// app. Metal Performance Shaders achieves this by automatically identifying
// [MPSTemporaryImage] objects that do not overlap in time over the course of
// a [MTLCommandBuffer] object’s lifetime and can therefore reuse the same
// memory. [MPSTemporaryImage] objects leverage an internal cache of
// preallocated reusable memory to hold pixel data to avoid typical memory
// allocation performance penalties common to ordinary [MPSImage] and
// [MTLTexture] objects.
//
// To avoid data corruption due to aliasing, [MPSTemporaryImage] objects
// impose some important restrictions:
//
// - The underlying texture storage mode is [MTLStorageMode.private]. You
// cannot, for example, use the [getBytes(_:bytesPerRow:from:mipmapLevel:)] or
// [replace(region:mipmapLevel:withBytes:bytesPerRow:)] methods with them.
// Temporary images are strictly read and written by the GPU. - The temporary
// image may be used only on a single [MTLCommandBuffer] object. This limits
// the chronology to a single linear time stream. - The
// [MPSTemporaryImage.ReadCount] property must be managed correctly. -
// Temporary images must also adhere to the general pixel format restrictions
// for [MPSImage] objects.
//
// Since temporary images can only be used with a single command buffer, and
// can not be used off the GPU, they generally should not be kept around past
// the completion of their associated command buffer. The lifetime of a
// temporary image is typically expected to be extremely short, perhaps
// spanning only a few lines of code.
//
// To keep the lifetime of the underlying texture allocation as short as
// possible, the texture is not allocated until the first time the
// [MPSTemporaryImage] object is used by an [MPSCNNKernel] object or until the
// first time the [MPSImage.Texture] property is read. The
// [MPSTemporaryImage.ReadCount] property serves to limit the lifetime of the
// texture on deallocation.
//
// You may use the [MPSImage.Texture] property with the `encode` methods of an
// [MPSUnaryImageKernel] subclass, if `featureChannels<=4` and the texture
// conforms to the requirements of the given kernel. In such cases, the
// [MPSTemporaryImage.ReadCount] property is not modified, since the enclosing
// object is not available. There is no locking mechanism provided to prevent
// a [MTLTexture] object returned from the [MPSImage.Texture] property from
// becoming invalid when the value of the [MPSTemporaryImage.ReadCount]
// property reaches 0.
//
// [MPSTemporaryImage] objects can otherwise be used wherever [MPSImage]
// objects are used.
//
// # The MPSTemporaryImage Class
//
// The [MPSTemporaryImage] class extends the [MPSImage] class to provide
// advanced caching of unused memory, in order to increase performance and
// reduce memory footprint. [MPSTemporaryImage] objects are intended as fast
// GPU-only storage for intermediate image data needed only transiently within
// a single [MTLCommandBuffer] object. They accelerate the common case of
// image data which is created only to be consumed and destroyed immediately
// by the next operation(s) encoded in a command buffer. [MPSTemporaryImage]
// objects provide a convenient and simple way to save memory by automatically
// aliasing other [MPSTemporaryImage] objects in the same command buffer.
// Because they alias (i.e., share texel storage with) other textures in the
// same command buffer, the valid lifetime of the data in an
// [MPSTemporaryImage] object is extremely short, limited to a portion of a
// the command buffer itself.
//
// You can not read or write data to an [MPSTemporaryImage] using the CPU, or
// use the data in other [MTLCommandBuffer] objects. Use regular [MPSImage]
// objects for more persistent storage.
//
// # Properties
//
//   - [MPSTemporaryImage.ReadCount]: The number of times a temporary image may be read by a CNN kernel before its contents become undefined.
//   - [MPSTemporaryImage.SetReadCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryImage
//
// [MTLCommandBuffer]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer
// [MTLStorageMode.private]: https://developer.apple.com/documentation/Metal/MTLStorageMode/private
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
// [getBytes(_:bytesPerRow:from:mipmapLevel:)]: https://developer.apple.com/documentation/Metal/MTLTexture/getBytes(_:bytesPerRow:from:mipmapLevel:)
// [replace(region:mipmapLevel:withBytes:bytesPerRow:)]: https://developer.apple.com/documentation/Metal/MTLTexture/replace(region:mipmapLevel:withBytes:bytesPerRow:)
type MPSTemporaryImage struct {
	MPSImage
}

// MPSTemporaryImageFromID constructs a [MPSTemporaryImage] from an objc.ID.
//
// A texture for use in convolutional neural networks that stores transient
// data to be used and discarded promptly.
func MPSTemporaryImageFromID(id objc.ID) MPSTemporaryImage {
	return MPSTemporaryImage{MPSImage: MPSImageFromID(id)}
}

// NOTE: MPSTemporaryImage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSTemporaryImage] class.
//
// # Properties
//
//   - [IMPSTemporaryImage.ReadCount]: The number of times a temporary image may be read by a CNN kernel before its contents become undefined.
//   - [IMPSTemporaryImage.SetReadCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryImage
type IMPSTemporaryImage interface {
	IMPSImage

	// Topic: Properties

	// The number of times a temporary image may be read by a CNN kernel before its contents become undefined.
	ReadCount() uint
	SetReadCount(value uint)
}

// Init initializes the instance.
func (t MPSTemporaryImage) Init() MPSTemporaryImage {
	rv := objc.Send[MPSTemporaryImage](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MPSTemporaryImage) Autorelease() MPSTemporaryImage {
	rv := objc.Send[MPSTemporaryImage](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSTemporaryImage creates a new MPSTemporaryImage instance.
func NewMPSTemporaryImage() MPSTemporaryImage {
	class := getMPSTemporaryImageClass()
	rv := objc.Send[MPSTemporaryImage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a temporary image for use on a command buffer.
//
// commandBuffer: The command buffer on which the temporary image will be exclusively used.
//
// imageDescriptor: An image descriptor that describes the image to create.
//
// # Return Value
//
// A valid [MPSTemporaryImage] object.
//
// # Discussion
//
// The temporary image will be released when the command buffer is committed.
// The underlying texture will become invalid before this time due to the
// action of the [MPSTemporaryImage.ReadCount] property.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryImage/init(commandBuffer:imageDescriptor:)
func NewTemporaryImageWithCommandBufferImageDescriptor(commandBuffer metal.MTLCommandBuffer, imageDescriptor IMPSImageDescriptor) MPSTemporaryImage {
	rv := objc.Send[objc.ID](objc.ID(getMPSTemporaryImageClass().class), objc.Sel("temporaryImageWithCommandBuffer:imageDescriptor:"), commandBuffer, imageDescriptor)
	return MPSTemporaryImageFromID(rv)
}

// Low-level interface for creating a temporary image using a texture
// descriptor.
//
// commandBuffer: The command buffer on which the temporary image will be exclusively used.
//
// textureDescriptor: A texture descriptor that describes the temporary image texture to create.
//
// # Return Value
//
// A valid [MPSTemporaryImage] object.
//
// # Discussion
//
// The temporary image will be released when the command buffer is committed.
// The underlying texture will become invalid before this time due to the
// action of the [MPSTemporaryImage.ReadCount] property.
//
// This function provides access to pixel formats not typically covered by the
// [MPSTemporaryImageClass.TemporaryImageWithCommandBufferImageDescriptor]
// method. The feature channels will be inferred from the pixel format without
// changing the width. The following restrictions apply:
//
// - The texture type must be [MTLTextureType.type2D] or
// [MTLTextureType.type2DArray]. - The texture usage must contain at least one
// of [shaderRead] or [shaderWrite]. - The storage mode must be
// [MTLStorageMode.private]. - The depth must be 1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryImage/init(commandBuffer:textureDescriptor:)
//
// [MTLStorageMode.private]: https://developer.apple.com/documentation/Metal/MTLStorageMode/private
// [MTLTextureType.type2DArray]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DArray
// [MTLTextureType.type2D]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
// [shaderRead]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderRead
// [shaderWrite]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderWrite
func NewTemporaryImageWithCommandBufferTextureDescriptor(commandBuffer metal.MTLCommandBuffer, textureDescriptor metal.MTLTextureDescriptor) MPSTemporaryImage {
	rv := objc.Send[objc.ID](objc.ID(getMPSTemporaryImageClass().class), objc.Sel("temporaryImageWithCommandBuffer:textureDescriptor:"), commandBuffer, textureDescriptor)
	return MPSTemporaryImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryImage/init(commandBuffer:textureDescriptor:featureChannels:)
func NewTemporaryImageWithCommandBufferTextureDescriptorFeatureChannels(commandBuffer metal.MTLCommandBuffer, textureDescriptor metal.MTLTextureDescriptor, featureChannels uint) MPSTemporaryImage {
	rv := objc.Send[objc.ID](objc.ID(getMPSTemporaryImageClass().class), objc.Sel("temporaryImageWithCommandBuffer:textureDescriptor:featureChannels:"), commandBuffer, textureDescriptor, featureChannels)
	return MPSTemporaryImageFromID(rv)
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
func NewTemporaryImageWithDeviceImageDescriptor(device metal.MTLDevice, imageDescriptor IMPSImageDescriptor) MPSTemporaryImage {
	instance := getMPSTemporaryImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:imageDescriptor:"), device, imageDescriptor)
	return MPSTemporaryImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImage/init(parentImage:sliceRange:featureChannels:)
func NewTemporaryImageWithParentImageSliceRangeFeatureChannels(parent IMPSImage, sliceRange foundation.NSRange, featureChannels uint) MPSTemporaryImage {
	instance := getMPSTemporaryImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParentImage:sliceRange:featureChannels:"), parent, sliceRange, featureChannels)
	return MPSTemporaryImageFromID(rv)
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
func NewTemporaryImageWithTextureFeatureChannels(texture metal.MTLTexture, featureChannels uint) MPSTemporaryImage {
	instance := getMPSTemporaryImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTexture:featureChannels:"), texture, featureChannels)
	return MPSTemporaryImageFromID(rv)
}

// A method that helps the framework decide which allocations to make ahead of
// time.
//
// commandBuffer: The command buffer on which the temporary images will be exclusively used.
//
// descriptorList: An array of image descriptors that describe the temporary images that will
// be created.
//
// # Discussion
//
// The texture cache that underlies the temporary images can automatically
// allocate new storage as needed, whenever you create new temporary images.
// However, sometimes a more global view of what you plan to make is useful
// for maximizing memory reuse to get the most efficient operation. Calling
// this class method provides a hint to the texture cache about what the list
// of temporary images will be.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryImage/prefetchStorage(with:imageDescriptorList:)
func (_MPSTemporaryImageClass MPSTemporaryImageClass) PrefetchStorageWithCommandBufferImageDescriptorList(commandBuffer metal.MTLCommandBuffer, descriptorList []MPSImageDescriptor) {
	objc.Send[objc.ID](objc.ID(_MPSTemporaryImageClass.class), objc.Sel("prefetchStorageWithCommandBuffer:imageDescriptorList:"), commandBuffer, objectivec.IObjectSliceToNSArray(descriptorList))
}

// The number of times a temporary image may be read by a CNN kernel before
// its contents become undefined.
//
// # Discussion
//
// Temporary images must release their underlying textures for reuse
// immediately after last use. In order to facilitate prompt and convenient
// memory recycling, each time a [MPSTemporaryImage] object is read by an
// `encode` method of an [MPSCNNKernel] object, the value of its
// [MPSTemporaryImage.ReadCount] property is automatically decremented. When
// the value of [MPSTemporaryImage.ReadCount] reaches 0, the underlying
// texture is automatically made available and reusable to the framework for
// its own needs (and for other [MPSTemporaryImage] objects prior to return
// from the `encode` method). The contents of the underlying texture become
// undefined at this time.
//
// By default, the value of [MPSTemporaryImage.ReadCount] is initialized to 1,
// indicating a temporary image that may be overwritten any number of times,
// but read only once.
//
// You may change the value of [MPSTemporaryImage.ReadCount] as desired to
// allow [MPSCNNKernel] objects to read the [MPSTemporaryImage] object
// additional times. However, it is an error to change the value of
// [MPSTemporaryImage.ReadCount] once it reaches 0 (it is also an error to
// read or write to a temporary image with a [MPSTemporaryImage.ReadCount]
// value of 0). You may set the value of [MPSTemporaryImage.ReadCount] to 0
// yourself to cause the underlying texture to be returned to the framework.
// Writing to a temporary image does not adjust the value of
// [MPSTemporaryImage.ReadCount].
//
// The Metal API Validation layer will assert if a temporary image is
// deallocated with a non-zero [MPSTemporaryImage.ReadCount] value to help
// identify cases when resources are not returned promptly.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryImage/readCount
func (t MPSTemporaryImage) ReadCount() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("readCount"))
	return rv
}
func (t MPSTemporaryImage) SetReadCount(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setReadCount:"), value)
}

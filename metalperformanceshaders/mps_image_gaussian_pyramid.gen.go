// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageGaussianPyramid] class.
var (
	_MPSImageGaussianPyramidClass     MPSImageGaussianPyramidClass
	_MPSImageGaussianPyramidClassOnce sync.Once
)

func getMPSImageGaussianPyramidClass() MPSImageGaussianPyramidClass {
	_MPSImageGaussianPyramidClassOnce.Do(func() {
		_MPSImageGaussianPyramidClass = MPSImageGaussianPyramidClass{class: objc.GetClass("MPSImageGaussianPyramid")}
	})
	return _MPSImageGaussianPyramidClass
}

// GetMPSImageGaussianPyramidClass returns the class object for MPSImageGaussianPyramid.
func GetMPSImageGaussianPyramidClass() MPSImageGaussianPyramidClass {
	return getMPSImageGaussianPyramidClass()
}

type MPSImageGaussianPyramidClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageGaussianPyramidClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageGaussianPyramidClass) Alloc() MPSImageGaussianPyramid {
	rv := objc.Send[MPSImageGaussianPyramid](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that convolves an image with a Gaussian pyramid.
//
// # Overview
//
// The Gaussian image pyramid kernel is enqueued as an in-place operation
// using the
// [MPSUnaryImageKernel.EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator]
// method. All mip-map levels (after level 1) present in the provided image
// are filled using the provided filter kernel. The `fallbackCopyAllocator`
// parameter is not used. The Gaussian image pyramid kernel ignores the
// [MPSUnaryImageKernel.ClipRect] and [MPSUnaryImageKernel.Offset] properties,
// and fills the entirety of the mip-map levels. Recall the size of the nth
// mip-map level as:
//
// - `w_n = max(1, floor(w_0 / 2^n))`
// - `h_n = max(1, floor(h_0 / 2^n))`
//
// Where `w_0` and `h_0` are the width and height of the 0th level,
// respectively (i.e. the image dimensions themselves).
//
// The Gaussian image pyramid is constructed as follows:
//
// - First, the 0th level mip-map of the input image is filtered with the
// specified convolution kernel. The default convolution filter kernel is `k =
// ww^T`, where `w = [1/16, 1/4, 3/8, 1/4, 1/16 ]^T`. You may also modify this
// kernel with a `centerWeight` parameter of `a` resulting in `k = ww^T`,
// where `w = [(1/4 - a/2), 1/4, a, 1/4, (1/4 - a/2) ]^T`, or you may provide
// a completely custom kernel. - Afterwards, the image is down-sampled by
// removing all odd rows and columns, which defines the next level in the
// Gaussian image pyramid. - This procedure is continued until every mip-map
// level present in the image is filled with all the pyramid levels.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianPyramid
type MPSImageGaussianPyramid struct {
	MPSImagePyramid
}

// MPSImageGaussianPyramidFromID constructs a [MPSImageGaussianPyramid] from an objc.ID.
//
// A filter that convolves an image with a Gaussian pyramid.
func MPSImageGaussianPyramidFromID(id objc.ID) MPSImageGaussianPyramid {
	return MPSImageGaussianPyramid{MPSImagePyramid: MPSImagePyramidFromID(id)}
}

// NOTE: MPSImageGaussianPyramid adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageGaussianPyramid] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianPyramid
type IMPSImageGaussianPyramid interface {
	IMPSImagePyramid
}

// Init initializes the instance.
func (i MPSImageGaussianPyramid) Init() MPSImageGaussianPyramid {
	rv := objc.Send[MPSImageGaussianPyramid](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageGaussianPyramid) Autorelease() MPSImageGaussianPyramid {
	rv := objc.Send[MPSImageGaussianPyramid](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageGaussianPyramid creates a new MPSImageGaussianPyramid instance.
func NewMPSImageGaussianPyramid() MPSImageGaussianPyramid {
	class := getMPSImageGaussianPyramidClass()
	rv := objc.Send[MPSImageGaussianPyramid](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageGaussianPyramidWithCoder(aDecoder foundation.INSCoder) MPSImageGaussianPyramid {
	instance := getMPSImageGaussianPyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageGaussianPyramidFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(coder:device:)
func NewImageGaussianPyramidWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageGaussianPyramid {
	instance := getMPSImageGaussianPyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageGaussianPyramidFromID(rv)
}

// Initializes a downwards 5-tap image pyramid with the default filter kernel
// and device.
//
// device: The device the filter will run on.
//
// # Return Value
//
// A valid [MPSImagePyramid] object or `nil`, if failure.
//
// # Discussion
//
// The filter kernel is the outer product of `w = [1/16, 1/4, 3/8, 1/4,
// 1/16]^T`, with itself.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:)
func NewImageGaussianPyramidWithDevice(device metal.MTLDevice) MPSImageGaussianPyramid {
	instance := getMPSImageGaussianPyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageGaussianPyramidFromID(rv)
}

// Initialize a downwards 5-tap image pyramid with a central weight parameter
// and device.
//
// device: The device the filter will run on.
//
// centerWeight: Defines the form of the filter kernel through the outer product `ww^T`,
// where `w = [(1/4 - a/2), 1/4, a, 1/4, (1/4 - a/2)]^T` and `a` is the value
// of `centerWeight`.
//
// # Return Value
//
// A valid [MPSImagePyramid] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:centerWeight:)
func NewImageGaussianPyramidWithDeviceCenterWeight(device metal.MTLDevice, centerWeight float32) MPSImageGaussianPyramid {
	instance := getMPSImageGaussianPyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:centerWeight:"), device, centerWeight)
	return MPSImageGaussianPyramidFromID(rv)
}

// Initialize a downwards n-tap image pyramid with a custom filter kernel and
// device.
//
// device: The device the filter will run on.
//
// kernelWidth: The width of the filter kernel.
//
// kernelHeight: The height of the filter kernel.
//
// kernelWeights: A pointer to an array of `kernelWidth*kernelHeight` values to be used as
// the kernel. These values are in row-major order.
//
// # Return Value
//
// A valid [MPSImagePyramid] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:kernelWidth:kernelHeight:weights:)
func NewImageGaussianPyramidWithDeviceKernelWidthKernelHeightWeights(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float32) MPSImageGaussianPyramid {
	instance := getMPSImageGaussianPyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), device, kernelWidth, kernelHeight, kernelWeights)
	return MPSImageGaussianPyramidFromID(rv)
}

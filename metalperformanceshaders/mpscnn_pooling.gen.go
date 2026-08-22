// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPooling] class.
var (
	_MPSCNNPoolingClass     MPSCNNPoolingClass
	_MPSCNNPoolingClassOnce sync.Once
)

func getMPSCNNPoolingClass() MPSCNNPoolingClass {
	_MPSCNNPoolingClassOnce.Do(func() {
		_MPSCNNPoolingClass = MPSCNNPoolingClass{class: objc.GetClass("MPSCNNPooling")}
	})
	return _MPSCNNPoolingClass
}

// GetMPSCNNPoolingClass returns the class object for MPSCNNPooling.
func GetMPSCNNPoolingClass() MPSCNNPoolingClass {
	return getMPSCNNPoolingClass()
}

type MPSCNNPoolingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingClass) Alloc() MPSCNNPooling {
	rv := objc.Send[MPSCNNPooling](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A pooling kernel.
//
// # Overview
//
// Pooling is a form of non-linear sub-sampling. Pooling partitions the input
// image into a set of rectangles (overlapping or non-overlapping) and, for
// each such sub-region, outputs a value. The pooling operation is used in
// computer vision to reduce the dimensionality of intermediate
// representations.
//
// The encode methods in the [MPSCNNKernel] class can be used to encode an
// [MPSCNNPooling] object to a [MTLCommandBuffer] object. The exact location
// of the pooling window for each output value is determined as follows:
//
// - The pooling window center for the first (top left) output pixel of the
// clip rectangle is at spatial coordinates `(offset.X(), offset.Y())` in the
// input image. - From this, the top left corner of the pooling window is at
// `(offset.x - floor(kernelWidth/2)`, `offset.y - floor(kernelHeight/2))` and
// extends `(kernelWidth, kernelHeight)` pixels to the right and down
// direction, which means that the last pixel to be included into the pooling
// window is at `(offset.x + floor((kernelWidth-1)/2)`, `offset.y +
// floor((kernelHeight-1)/2))`, so that for even kernel sizes the pooling
// window extends one pixel more into the left and up direction. - The
// following pooling windows can be then easily deduced from the first one by
// simple shifting the source coordinates according to the values of the
// `strideInPixelsX` and `strideInPixelsY` properties.
//
// For example, the pooling window center `w(x,y)` for the output value at
// coordinate `(x,y)` of the destination clip rectangle (`(x,y)` computed with
// regard to clipping rectangle origin) is at `w(x,y) = (offset.X() +
// strideInPixelsX * x , offset.Y() + strideInPixelsY * y)`.
//
// Quite often it is desirable to distribute the pooling windows as evenly as
// possible in the input image. As explained above, if the `offset` is zero,
// then the center of the first pooling window is at the top left corner of
// the input image, which means that the left and top stripes of the pooling
// window are read from outside the input image boundaries (when filter size
// is larger than unity). Also it may mean that some values from the bottom
// and right stripes are not included at all in the pooling, resulting in loss
// of valuable information.
//
// A scheme used in some common libraries is to shift the source `offset`
// according to the following formula:
//
// - `offset.xy += {(int)ceil(((L.xy - 1) % s.xy) / 2)}`, for odd `f.Xy()` -
// `offset.xy += {(int)floor(((L.xy - 1) % s.xy) / 2) + 1},` for even `f.Xy()`
//
// Where [L] is the size of the input image (or more accurately the size
// corresponding to the scaled `clipRect` value in source coordinates, which
// commonly coincides with the source image itself), `s.Xy()` is
// `(“strideInPixelsX`, `strideInPixelsY“)` and `f.Xy()` is `(kernelWidth,
// kernelHeight)`.
//
// This offset distributes the pooling window centers evenly in the effective
// source `clipRect`, when the output size is rounded up with regards to
// stride (`output size = ceil(input size / stride)`) and is commonly used in
// CNN libraries (for example TensorFlow uses this offset scheme in its
// maximum pooling implementation `tf.NnXCUIElementTypeMax_pool()` with
// `'S“AME“'` - padding, for `'VALID'` padding one can simply set `offset.xy
// += floor(f.xy/2)` to get the first pooling window inside the source image
// completely).
//
// For an [MPSCNNPoolingMax] object, the way the input image borders are
// handled can become important: if there are negative values in the source
// image near the borders of the image and the pooling window crosses the
// borders, then using a [MPSImageEdgeModeZero] edge modemay cause the maximum
// pooling operation to override the negative input data values with zeros
// coming from outside the source image borders, resulting in large boundary
// effects. A simple way to avoid this is to use a [MPSImageEdgeModeClamp]
// edge mode, which for an [MPSCNNPoolingMax] object effectively causes all
// pooling windows to remain within the source image.
//
// # Instance Methods
//
//   - [MPSCNNPooling.InitWithDeviceKernelWidthKernelHeight]: Initializes a pooling filter.
//   - [MPSCNNPooling.InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY]: Initializes a pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling
//
// [MTLCommandBuffer]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer
type MPSCNNPooling struct {
	MPSCNNKernel
}

// MPSCNNPoolingFromID constructs a [MPSCNNPooling] from an objc.ID.
//
// A pooling kernel.
func MPSCNNPoolingFromID(id objc.ID) MPSCNNPooling {
	return MPSCNNPooling{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNPooling adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPooling] class.
//
// # Instance Methods
//
//   - [IMPSCNNPooling.InitWithDeviceKernelWidthKernelHeight]: Initializes a pooling filter.
//   - [IMPSCNNPooling.InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY]: Initializes a pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling
type IMPSCNNPooling interface {
	IMPSCNNKernel

	// Topic: Instance Methods

	// Initializes a pooling filter.
	InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPooling
	// Initializes a pooling filter.
	InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPooling
}

// Init initializes the instance.
func (c MPSCNNPooling) Init() MPSCNNPooling {
	rv := objc.Send[MPSCNNPooling](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPooling) Autorelease() MPSCNNPooling {
	rv := objc.Send[MPSCNNPooling](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPooling creates a new MPSCNNPooling instance.
func NewMPSCNNPooling() MPSCNNPooling {
	class := getMPSCNNPoolingClass()
	rv := objc.Send[MPSCNNPooling](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNPoolingWithCoder(aDecoder foundation.INSCoder) MPSCNNPooling {
	instance := getMPSCNNPoolingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNPoolingFromID(rv)
}

// Initializes a pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling/init(coder:device:)
func NewCNNPoolingWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNPooling {
	instance := getMPSCNNPoolingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNPoolingFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNPoolingWithDevice(device metal.MTLDevice) MPSCNNPooling {
	instance := getMPSCNNPoolingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNPoolingFromID(rv)
}

// Initializes a pooling filter.
//
// device: The device the kernel will run on.
//
// kernelWidth: The width of the kernel.
//
// This value can be odd or even.
//
// kernelHeight: The height of the kernel.
//
// This value can be odd or even.
//
// # Return Value
//
// A valid [MPSCNNPooling] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling/init(device:kernelWidth:kernelHeight:)
func NewCNNPoolingWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPooling {
	instance := getMPSCNNPoolingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNPoolingFromID(rv)
}

// Initializes a pooling filter.
//
// device: The device the kernel will run on.
//
// kernelWidth: The width of the kernel.
//
// This value can be odd or even.
//
// kernelHeight: The height of the kernel.
//
// This value can be odd or even.
//
// strideInPixelsX: The output stride (downsampling factor) in the x dimension.
//
// strideInPixelsY: The output stride (downsampling factor) in the y dimension.
//
// # Return Value
//
// A valid [MPSCNNPooling] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPooling {
	instance := getMPSCNNPoolingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingFromID(rv)
}

// Initializes a pooling filter.
//
// device: The device the kernel will run on.
//
// kernelWidth: The width of the kernel.
//
// This value can be odd or even.
//
// kernelHeight: The height of the kernel.
//
// This value can be odd or even.
//
// # Return Value
//
// A valid [MPSCNNPooling] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling/init(device:kernelWidth:kernelHeight:)
func (c MPSCNNPooling) InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNPooling {
	rv := objc.Send[MPSCNNPooling](c.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return rv
}

// Initializes a pooling filter.
//
// device: The device the kernel will run on.
//
// kernelWidth: The width of the kernel.
//
// This value can be odd or even.
//
// kernelHeight: The height of the kernel.
//
// This value can be odd or even.
//
// strideInPixelsX: The output stride (downsampling factor) in the x dimension.
//
// strideInPixelsY: The output stride (downsampling factor) in the y dimension.
//
// # Return Value
//
// A valid [MPSCNNPooling] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling/init(device:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func (c MPSCNNPooling) InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPooling {
	rv := objc.Send[MPSCNNPooling](c.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

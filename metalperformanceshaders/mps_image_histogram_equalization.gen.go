// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageHistogramEqualization] class.
var (
	_MPSImageHistogramEqualizationClass     MPSImageHistogramEqualizationClass
	_MPSImageHistogramEqualizationClassOnce sync.Once
)

func getMPSImageHistogramEqualizationClass() MPSImageHistogramEqualizationClass {
	_MPSImageHistogramEqualizationClassOnce.Do(func() {
		_MPSImageHistogramEqualizationClass = MPSImageHistogramEqualizationClass{class: objc.GetClass("MPSImageHistogramEqualization")}
	})
	return _MPSImageHistogramEqualizationClass
}

// GetMPSImageHistogramEqualizationClass returns the class object for MPSImageHistogramEqualization.
func GetMPSImageHistogramEqualizationClass() MPSImageHistogramEqualizationClass {
	return getMPSImageHistogramEqualizationClass()
}

type MPSImageHistogramEqualizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageHistogramEqualizationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageHistogramEqualizationClass) Alloc() MPSImageHistogramEqualization {
	rv := objc.Send[MPSImageHistogramEqualization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that equalizes the histogram of an image.
//
// # Overview
//
// The process is divided into three steps:
//
// - Call the [MPSImageHistogramEqualization.InitWithDeviceHistogramInfo]
// method to create a [MPSImageHistogramEqualization] object. - Call the
// [MPSImageHistogramEqualization.EncodeTransformToCommandBufferSourceTextureHistogramHistogramOffset]
// method. This creates a privately held image transform (i.e. a cumulative
// distribution function of the histogram) which will be used to equalize the
// distribution of the histogram of the source image. This process runs on a
// command buffer when it is committed to a command queue. It must complete
// before the next step can be run. It may be performed on the same command
// buffer. The `histogram` argument specifies the histogram buffer which
// contains the histogram values for the source texture. The `sourceTexture`
// argument is used by the method to determine the number of channels and
// therefore which histogram data in the histogram buffer to use. The
// histogram for the source texture must have been computed either on the CPU
// or using the [MPSImageHistogram] kernel. - Call the
// [MPSUnaryImageKernel.EncodeToCommandBufferSourceTextureDestinationTexture]
// method to read data from the source texture, apply the equalization
// transform to it, and write to the destination texture. This step is also
// done on the GPU on a command queue.
//
// # Methods
//
//   - [MPSImageHistogramEqualization.InitWithDeviceHistogramInfo]: Initializes a histogram with specific information.
//   - [MPSImageHistogramEqualization.EncodeTransformToCommandBufferSourceTextureHistogramHistogramOffset]: Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
//
// # Properties
//
//   - [MPSImageHistogramEqualization.HistogramInfo]: A structure describing the histogram content.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramEqualization
type MPSImageHistogramEqualization struct {
	MPSUnaryImageKernel
}

// MPSImageHistogramEqualizationFromID constructs a [MPSImageHistogramEqualization] from an objc.ID.
//
// A filter that equalizes the histogram of an image.
func MPSImageHistogramEqualizationFromID(id objc.ID) MPSImageHistogramEqualization {
	return MPSImageHistogramEqualization{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageHistogramEqualization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageHistogramEqualization] class.
//
// # Methods
//
//   - [IMPSImageHistogramEqualization.InitWithDeviceHistogramInfo]: Initializes a histogram with specific information.
//   - [IMPSImageHistogramEqualization.EncodeTransformToCommandBufferSourceTextureHistogramHistogramOffset]: Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
//
// # Properties
//
//   - [IMPSImageHistogramEqualization.HistogramInfo]: A structure describing the histogram content.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramEqualization
type IMPSImageHistogramEqualization interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes a histogram with specific information.
	InitWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageHistogramEqualization
	// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
	EncodeTransformToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, histogram metal.MTLBuffer, histogramOffset uint)

	// Topic: Properties

	// A structure describing the histogram content.
	HistogramInfo() MPSImageHistogramInfo
}

// Init initializes the instance.
func (i MPSImageHistogramEqualization) Init() MPSImageHistogramEqualization {
	rv := objc.Send[MPSImageHistogramEqualization](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageHistogramEqualization) Autorelease() MPSImageHistogramEqualization {
	rv := objc.Send[MPSImageHistogramEqualization](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageHistogramEqualization creates a new MPSImageHistogramEqualization instance.
func NewMPSImageHistogramEqualization() MPSImageHistogramEqualization {
	class := getMPSImageHistogramEqualizationClass()
	rv := objc.Send[MPSImageHistogramEqualization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageHistogramEqualizationWithCoder(aDecoder foundation.INSCoder) MPSImageHistogramEqualization {
	instance := getMPSImageHistogramEqualizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageHistogramEqualizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramEqualization/init(coder:device:)
func NewImageHistogramEqualizationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageHistogramEqualization {
	instance := getMPSImageHistogramEqualizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageHistogramEqualizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageHistogramEqualizationWithDevice(device metal.MTLDevice) MPSImageHistogramEqualization {
	instance := getMPSImageHistogramEqualizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageHistogramEqualizationFromID(rv)
}

// Initializes a histogram with specific information.
//
// device: The Metal device the filter will run on.
//
// histogramInfo: A pointer to a structure describing the histogram content.
//
// # Return Value
//
// An initialized histogram object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramEqualization/init(device:histogramInfo:)
func NewImageHistogramEqualizationWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageHistogramEqualization {
	instance := getMPSImageHistogramEqualizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:histogramInfo:"), device, unsafe.Pointer(histogramInfo))
	return MPSImageHistogramEqualizationFromID(rv)
}

// Initializes a histogram with specific information.
//
// device: The Metal device the filter will run on.
//
// histogramInfo: A pointer to a structure describing the histogram content.
//
// # Return Value
//
// An initialized histogram object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramEqualization/init(device:histogramInfo:)
func (i MPSImageHistogramEqualization) InitWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageHistogramEqualization {
	rv := objc.Send[MPSImageHistogramEqualization](i.ID, objc.Sel("initWithDevice:histogramInfo:"), device, unsafe.Pointer(histogramInfo))
	return rv
}

// Encodes the transform function to a command buffer using a compute command
// encoder. The transform function computes the equalization lookup table.
//
// commandBuffer: A valid command buffer.
//
// source: A valid texture containing the source image for the filter.
//
// histogram: A valid buffer containing the histogram results for an image. This filter
// will use these histogram results to generate the cumulative histogram for
// equalizing the image. The histogram results per channel are stored
// together. The number of channels for which histogram results are stored is
// determined by the number of channels in the image. If the
// `histogramForAlpha` value of the
// [MPSImageHistogramEqualization.HistogramInfo] property is false and the
// source image is RGBA, then only histogram results for RGB channels are
// stored.
//
// histogramOffset: The byte offset into the histogram buffer where the histogram starts. Must
// conform to alignment requirements for the `offset` parameter of the
// [setBuffer(_:offset:index:)] method.
//
// # Discussion
//
// The transform function will not begin to execute until after the command
// buffer has been enqueued and committed. This step will need to be repeated
// with the new [MPSKernel] object if the [MPSKernel.CopyWithZoneDevice] or
// [copy(with:)] method is called. The transform is stored as internal state
// to the object. You still need to call the
// [MPSUnaryImageKernel.EncodeToCommandBufferSourceTextureDestinationTexture]
// afterward to apply the transform to produce a result texture.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramEqualization/encodeTransform(to:sourceTexture:histogram:histogramOffset:)
//
// [setBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffer(_:offset:index:)
// [copy(with:)]: https://developer.apple.com/documentation/Foundation/NSCopying/copy(with:)
func (i MPSImageHistogramEqualization) EncodeTransformToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, histogram metal.MTLBuffer, histogramOffset uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeTransformToCommandBuffer:sourceTexture:histogram:histogramOffset:"), commandBuffer, source, histogram, histogramOffset)
}

// A structure describing the histogram content.
//
// # Discussion
//
// Returns a structure describing the format of the histogram.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramEqualization/histogramInfo
func (i MPSImageHistogramEqualization) HistogramInfo() MPSImageHistogramInfo {
	rv := objc.Send[MPSImageHistogramInfo](i.ID, objc.Sel("histogramInfo"))
	return MPSImageHistogramInfo(rv)
}

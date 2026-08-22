// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageHistogramSpecification] class.
var (
	_MPSImageHistogramSpecificationClass     MPSImageHistogramSpecificationClass
	_MPSImageHistogramSpecificationClassOnce sync.Once
)

func getMPSImageHistogramSpecificationClass() MPSImageHistogramSpecificationClass {
	_MPSImageHistogramSpecificationClassOnce.Do(func() {
		_MPSImageHistogramSpecificationClass = MPSImageHistogramSpecificationClass{class: objc.GetClass("MPSImageHistogramSpecification")}
	})
	return _MPSImageHistogramSpecificationClass
}

// GetMPSImageHistogramSpecificationClass returns the class object for MPSImageHistogramSpecification.
func GetMPSImageHistogramSpecificationClass() MPSImageHistogramSpecificationClass {
	return getMPSImageHistogramSpecificationClass()
}

type MPSImageHistogramSpecificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageHistogramSpecificationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageHistogramSpecificationClass) Alloc() MPSImageHistogramSpecification {
	rv := objc.Send[MPSImageHistogramSpecification](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that performs a histogram specification operation on an image.
//
// # Overview
//
// [MPSImageHistogramSpecification] is a generalized version of histogram
// equalization operation. The histogram specification filter converts the
// image so that its histogram matches the desired histogram.
//
// The process is divided into three steps:
//
// - Call the [MPSImageHistogramSpecification.InitWithDeviceHistogramInfo]
// method to create a [MPSImageHistogramSpecification] object. - Call the
// [MPSImageHistogramSpecification.EncodeTransformToCommandBufferSourceTextureSourceHistogramSourceHistogramOffsetDesiredHistogramDesiredHistogramOffset]
// method. This creates a privately held image transform which will convert
// the distribution of the source histogram to the desired histogram. This
// process runs on a command buffer when it is committed to a command queue.
// It must complete before the next step can be run. It may be performed on
// the same command buffer. The `sourceTexture` argument is used by the method
// to determine the number of channels and therefore which histogram data in
// the source histogram buffer to use. The source histogram and desired
// histogram must have been computed either on the CPU or using the
// [MPSImageHistogram] kernel. - Call the
// [MPSUnaryImageKernel.EncodeToCommandBufferSourceTextureDestinationTexture]
// method to read data from the source texture, apply the equalization
// transform to it, and write to the destination texture. This step is also
// done on the GPU on a command queue.
//
// # Methods
//
//   - [MPSImageHistogramSpecification.InitWithDeviceHistogramInfo]: Initializes a histogram with specific information.
//   - [MPSImageHistogramSpecification.EncodeTransformToCommandBufferSourceTextureSourceHistogramSourceHistogramOffsetDesiredHistogramDesiredHistogramOffset]: Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
//
// # Properties
//
//   - [MPSImageHistogramSpecification.HistogramInfo]: A structure describing the histogram content.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramSpecification
type MPSImageHistogramSpecification struct {
	MPSUnaryImageKernel
}

// MPSImageHistogramSpecificationFromID constructs a [MPSImageHistogramSpecification] from an objc.ID.
//
// A filter that performs a histogram specification operation on an image.
func MPSImageHistogramSpecificationFromID(id objc.ID) MPSImageHistogramSpecification {
	return MPSImageHistogramSpecification{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageHistogramSpecification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageHistogramSpecification] class.
//
// # Methods
//
//   - [IMPSImageHistogramSpecification.InitWithDeviceHistogramInfo]: Initializes a histogram with specific information.
//   - [IMPSImageHistogramSpecification.EncodeTransformToCommandBufferSourceTextureSourceHistogramSourceHistogramOffsetDesiredHistogramDesiredHistogramOffset]: Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
//
// # Properties
//
//   - [IMPSImageHistogramSpecification.HistogramInfo]: A structure describing the histogram content.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramSpecification
type IMPSImageHistogramSpecification interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes a histogram with specific information.
	InitWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageHistogramSpecification
	// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
	EncodeTransformToCommandBufferSourceTextureSourceHistogramSourceHistogramOffsetDesiredHistogramDesiredHistogramOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, sourceHistogram metal.MTLBuffer, sourceHistogramOffset uint, desiredHistogram metal.MTLBuffer, desiredHistogramOffset uint)

	// Topic: Properties

	// A structure describing the histogram content.
	HistogramInfo() MPSImageHistogramInfo
}

// Init initializes the instance.
func (i MPSImageHistogramSpecification) Init() MPSImageHistogramSpecification {
	rv := objc.Send[MPSImageHistogramSpecification](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageHistogramSpecification) Autorelease() MPSImageHistogramSpecification {
	rv := objc.Send[MPSImageHistogramSpecification](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageHistogramSpecification creates a new MPSImageHistogramSpecification instance.
func NewMPSImageHistogramSpecification() MPSImageHistogramSpecification {
	class := getMPSImageHistogramSpecificationClass()
	rv := objc.Send[MPSImageHistogramSpecification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageHistogramSpecificationWithCoder(aDecoder foundation.INSCoder) MPSImageHistogramSpecification {
	instance := getMPSImageHistogramSpecificationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageHistogramSpecificationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramSpecification/init(coder:device:)
func NewImageHistogramSpecificationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageHistogramSpecification {
	instance := getMPSImageHistogramSpecificationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageHistogramSpecificationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageHistogramSpecificationWithDevice(device metal.MTLDevice) MPSImageHistogramSpecification {
	instance := getMPSImageHistogramSpecificationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageHistogramSpecificationFromID(rv)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramSpecification/init(device:histogramInfo:)
func NewImageHistogramSpecificationWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageHistogramSpecification {
	instance := getMPSImageHistogramSpecificationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:histogramInfo:"), device, unsafe.Pointer(histogramInfo))
	return MPSImageHistogramSpecificationFromID(rv)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramSpecification/init(device:histogramInfo:)
func (i MPSImageHistogramSpecification) InitWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageHistogramSpecification {
	rv := objc.Send[MPSImageHistogramSpecification](i.ID, objc.Sel("initWithDevice:histogramInfo:"), device, unsafe.Pointer(histogramInfo))
	return rv
}

// Encodes the transform function to a command buffer using a compute command
// encoder. The transform function computes the equalization lookup table.
//
// commandBuffer: A valid command buffer.
//
// source: A valid texture containing the source image for the filter.
//
// sourceHistogram: A valid buffer containing the histogram results for the source image. This
// filter will use these histogram results to generate the cumulative
// histogram for equalizing the image. The histogram results per channel are
// stored together. The number of channels for which histogram results are
// stored is determined by the number of channels in the image. If the
// `histogramForAlpha` value of the
// [MPSImageHistogramSpecification.HistogramInfo] property is false and the
// source image is RGBA, then only histogram results for RGB channels are
// stored.
//
// sourceHistogramOffset: The byte offset into the source histogram buffer where the histogram
// starts. Must conform to alignment requirements for the `offset` parameter
// of the [setBuffer(_:offset:index:)] method.
//
// desiredHistogram: A valid buffer containing the desired histogram results for the source
// image. The histogram results per channel are stored together. The number of
// channels for which histogram results are stored is determined by the number
// of channels in the image. If the `histogramForAlpha` value of the
// [MPSImageHistogramSpecification.HistogramInfo] property is false and the
// source image is RGBA, then only histogram results for RGB channels are
// stored.
//
// desiredHistogramOffset: The byte offset into the desired histogram buffer where the histogram
// starts. Must conform to alignment requirements for the `offset` parameter
// of the [setBuffer(_:offset:index:)] method.
//
// # Discussion
//
// The transform function will not begin to execute until after the command
// buffer has been enqueued and committed. This step will need to be repeated
// with the new [MPSKernel] object if the [MPSKernel.CopyWithZoneDevice] or
// [copy(with:)] method is called.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramSpecification/encodeTransform(to:sourceTexture:sourceHistogram:sourceHistogramOffset:desiredHistogram:desiredHistogramOffset:)
//
// [setBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffer(_:offset:index:)
// [copy(with:)]: https://developer.apple.com/documentation/Foundation/NSCopying/copy(with:)
//
// [setBuffer(_:offset:index:)]: https://developer.apple.com/documentation/Metal/MTLComputeCommandEncoder/setBuffer(_:offset:index:)
func (i MPSImageHistogramSpecification) EncodeTransformToCommandBufferSourceTextureSourceHistogramSourceHistogramOffsetDesiredHistogramDesiredHistogramOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, sourceHistogram metal.MTLBuffer, sourceHistogramOffset uint, desiredHistogram metal.MTLBuffer, desiredHistogramOffset uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeTransformToCommandBuffer:sourceTexture:sourceHistogram:sourceHistogramOffset:desiredHistogram:desiredHistogramOffset:"), commandBuffer, source, sourceHistogram, sourceHistogramOffset, desiredHistogram, desiredHistogramOffset)
}

// A structure describing the histogram content.
//
// # Discussion
//
// Returns a structure describing the format of the histogram.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramSpecification/histogramInfo
func (i MPSImageHistogramSpecification) HistogramInfo() MPSImageHistogramInfo {
	rv := objc.Send[MPSImageHistogramInfo](i.ID, objc.Sel("histogramInfo"))
	return MPSImageHistogramInfo(rv)
}

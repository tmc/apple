// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageHistogram] class.
var (
	_MPSImageHistogramClass     MPSImageHistogramClass
	_MPSImageHistogramClassOnce sync.Once
)

func getMPSImageHistogramClass() MPSImageHistogramClass {
	_MPSImageHistogramClassOnce.Do(func() {
		_MPSImageHistogramClass = MPSImageHistogramClass{class: objc.GetClass("MPSImageHistogram")}
	})
	return _MPSImageHistogramClass
}

// GetMPSImageHistogramClass returns the class object for MPSImageHistogram.
func GetMPSImageHistogramClass() MPSImageHistogramClass {
	return getMPSImageHistogramClass()
}

type MPSImageHistogramClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageHistogramClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageHistogramClass) Alloc() MPSImageHistogram {
	rv := objc.Send[MPSImageHistogram](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that computes the histogram of an image.
//
// # Overview
//
// Typically, you use an [MPSImageHistogram] filter to calculate an image’s
// histogram that is passed to a subsequent filter such as
// [MPSImageHistogramEqualization] or [MPSImageHistogramSpecification].
//
// The following listing shows how you can create a histogram filter to
// calculate the histogram of the [MTLTexture], `sourceTexture`. The filter is
// passed an instance of [MPSImageHistogramInfo] that specifies information to
// compute the histogram for the channels of an image. After encoding,
// `histogramInfoBuffer` contains the histogram information and can be used
// for further operations such as equalization or specification.
//
// Listing 1. Creating a histogram filter
//
// # Methods
//
//   - [MPSImageHistogram.InitWithDeviceHistogramInfo]: Initializes a histogram with specific information.
//   - [MPSImageHistogram.EncodeToCommandBufferSourceTextureHistogramHistogramOffset]: Encodes the filter to a command buffer using a compute command encoder.
//   - [MPSImageHistogram.HistogramSizeForSourceFormat]: The amount of space the histogram will take up in the output buffer.
//
// # Properties
//
//   - [MPSImageHistogram.ClipRectSource]: The source rectangle to use when reading data.
//   - [MPSImageHistogram.SetClipRectSource]
//   - [MPSImageHistogram.ZeroHistogram]: Determines whether to zero-initialize the histogram results.
//   - [MPSImageHistogram.SetZeroHistogram]
//   - [MPSImageHistogram.HistogramInfo]: A structure describing the histogram content.
//   - [MPSImageHistogram.MinPixelThresholdValue]
//   - [MPSImageHistogram.SetMinPixelThresholdValue]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram
//
// [MPSImageHistogramInfo]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramInfo
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
type MPSImageHistogram struct {
	MPSKernel
}

// MPSImageHistogramFromID constructs a [MPSImageHistogram] from an objc.ID.
//
// A filter that computes the histogram of an image.
func MPSImageHistogramFromID(id objc.ID) MPSImageHistogram {
	return MPSImageHistogram{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSImageHistogram adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageHistogram] class.
//
// # Methods
//
//   - [IMPSImageHistogram.InitWithDeviceHistogramInfo]: Initializes a histogram with specific information.
//   - [IMPSImageHistogram.EncodeToCommandBufferSourceTextureHistogramHistogramOffset]: Encodes the filter to a command buffer using a compute command encoder.
//   - [IMPSImageHistogram.HistogramSizeForSourceFormat]: The amount of space the histogram will take up in the output buffer.
//
// # Properties
//
//   - [IMPSImageHistogram.ClipRectSource]: The source rectangle to use when reading data.
//   - [IMPSImageHistogram.SetClipRectSource]
//   - [IMPSImageHistogram.ZeroHistogram]: Determines whether to zero-initialize the histogram results.
//   - [IMPSImageHistogram.SetZeroHistogram]
//   - [IMPSImageHistogram.HistogramInfo]: A structure describing the histogram content.
//   - [IMPSImageHistogram.MinPixelThresholdValue]
//   - [IMPSImageHistogram.SetMinPixelThresholdValue]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram
type IMPSImageHistogram interface {
	IMPSKernel

	// Topic: Methods

	// Initializes a histogram with specific information.
	InitWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageHistogram
	// Encodes the filter to a command buffer using a compute command encoder.
	EncodeToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, histogram metal.MTLBuffer, histogramOffset uint)
	// The amount of space the histogram will take up in the output buffer.
	HistogramSizeForSourceFormat(sourceFormat metal.MTLPixelFormat) uintptr

	// Topic: Properties

	// The source rectangle to use when reading data.
	ClipRectSource() metal.MTLRegion
	SetClipRectSource(value metal.MTLRegion)
	// Determines whether to zero-initialize the histogram results.
	ZeroHistogram() bool
	SetZeroHistogram(value bool)
	// A structure describing the histogram content.
	HistogramInfo() MPSImageHistogramInfo
	MinPixelThresholdValue() [4]float32
	SetMinPixelThresholdValue(value [4]float32)
}

// Init initializes the instance.
func (i MPSImageHistogram) Init() MPSImageHistogram {
	rv := objc.Send[MPSImageHistogram](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageHistogram) Autorelease() MPSImageHistogram {
	rv := objc.Send[MPSImageHistogram](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageHistogram creates a new MPSImageHistogram instance.
func NewMPSImageHistogram() MPSImageHistogram {
	class := getMPSImageHistogramClass()
	rv := objc.Send[MPSImageHistogram](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageHistogramWithCoder(aDecoder foundation.INSCoder) MPSImageHistogram {
	instance := getMPSImageHistogramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageHistogramFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram/init(coder:device:)
func NewImageHistogramWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageHistogram {
	instance := getMPSImageHistogramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageHistogramFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewImageHistogramWithDevice(device metal.MTLDevice) MPSImageHistogram {
	instance := getMPSImageHistogramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageHistogramFromID(rv)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram/init(device:histogramInfo:)
func NewImageHistogramWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageHistogram {
	instance := getMPSImageHistogramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:histogramInfo:"), device, unsafe.Pointer(histogramInfo))
	return MPSImageHistogramFromID(rv)
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
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram/init(device:histogramInfo:)
func (i MPSImageHistogram) InitWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageHistogram {
	rv := objc.Send[MPSImageHistogram](i.ID, objc.Sel("initWithDevice:histogramInfo:"), device, unsafe.Pointer(histogramInfo))
	return rv
}

// Encodes the filter to a command buffer using a compute command encoder.
//
// commandBuffer: A valid command buffer.
//
// source: A valid texture containing the source image for the filter.
//
// histogram: A valid buffer to receive the histogram results.
//
// histogramOffset: The byte offset into the histogram buffer at which to write the histogram
// results. Must be a multiple of 32 bytes. The histogram results per channel
// are stored together. The number of channels for which histogram results are
// stored is determined by the number of channels in the image. If the
// `histogramForAlpha` value of the [MPSImageHistogram.HistogramInfo] property
// is false and the source image is RGBA, then only histogram results for RGB
// channels are stored.
//
// # Discussion
//
// The filter will not begin to execute until after the command buffer has
// been enqueued and committed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram/encode(to:sourceTexture:histogram:histogramOffset:)
func (i MPSImageHistogram) EncodeToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, histogram metal.MTLBuffer, histogramOffset uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:histogram:histogramOffset:"), commandBuffer, source, histogram, histogramOffset)
}

// The amount of space the histogram will take up in the output buffer.
//
// sourceFormat: The pixel format of the source image, corresponding to the `sourceTexture`
// object of the
// [MPSImageHistogram.EncodeToCommandBufferSourceTextureHistogramHistogramOffset]
// method.
//
// # Return Value
//
// The number of bytes needed to store the histogram results.
//
// # Discussion
//
// This convenience function calculates the minimum amount of space needed in
// the output histogram for the results. The buffer should be at least this
// length and longer if the `histogramOffset` value in the
// [MPSImageHistogram.EncodeToCommandBufferSourceTextureHistogramHistogramOffset]
// method is non-zero.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram/histogramSize(forSourceFormat:)
func (i MPSImageHistogram) HistogramSizeForSourceFormat(sourceFormat metal.MTLPixelFormat) uintptr {
	rv := objc.Send[uintptr](i.ID, objc.Sel("histogramSizeForSourceFormat:"), sourceFormat)
	return rv
}

// The source rectangle to use when reading data.
//
// # Discussion
//
// This value indicates which part of the source image to read from. If the
// value of `clipRectSource` does not lie completely within the source image,
// then the intersection of the image bounds and the value of `clipRectSource`
// will be used. The value of `clipRectSource` replaces the
// [MPSUnaryImageKernel.Offset] value for this filter, which is ignored.
//
// The default value is [MPSRectNoClip], indicating that the entire source
// texture is used.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram/clipRectSource
//
// [MPSRectNoClip]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRectNoClip
func (i MPSImageHistogram) ClipRectSource() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](i.ID, objc.Sel("clipRectSource"))
	return metal.MTLRegion(rv)
}
func (i MPSImageHistogram) SetClipRectSource(value metal.MTLRegion) {
	objc.Send[struct{}](i.ID, objc.Sel("setClipRectSource:"), value)
}

// Determines whether to zero-initialize the histogram results.
//
// # Discussion
//
// Determines whether the memory region in which the histogram results are to
// be written in the histogram buffer are to be zero-initialized or not.
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram/zeroHistogram
func (i MPSImageHistogram) ZeroHistogram() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("zeroHistogram"))
	return rv
}
func (i MPSImageHistogram) SetZeroHistogram(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setZeroHistogram:"), value)
}

// A structure describing the histogram content.
//
// # Discussion
//
// Returns a structure describing the format of the histogram.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram/histogramInfo
func (i MPSImageHistogram) HistogramInfo() MPSImageHistogramInfo {
	rv := objc.Send[MPSImageHistogramInfo](i.ID, objc.Sel("histogramInfo"))
	return MPSImageHistogramInfo(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram/minPixelThresholdValue
func (i MPSImageHistogram) MinPixelThresholdValue() [4]float32 {
	rv := objc.Send[[4]float32](i.ID, objc.Sel("minPixelThresholdValue"))
	return [4]float32(rv)
}
func (i MPSImageHistogram) SetMinPixelThresholdValue(value [4]float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setMinPixelThresholdValue:"), value)
}

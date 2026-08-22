// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageNormalizedHistogram] class.
var (
	_MPSImageNormalizedHistogramClass     MPSImageNormalizedHistogramClass
	_MPSImageNormalizedHistogramClassOnce sync.Once
)

func getMPSImageNormalizedHistogramClass() MPSImageNormalizedHistogramClass {
	_MPSImageNormalizedHistogramClassOnce.Do(func() {
		_MPSImageNormalizedHistogramClass = MPSImageNormalizedHistogramClass{class: objc.GetClass("MPSImageNormalizedHistogram")}
	})
	return _MPSImageNormalizedHistogramClass
}

// GetMPSImageNormalizedHistogramClass returns the class object for MPSImageNormalizedHistogram.
func GetMPSImageNormalizedHistogramClass() MPSImageNormalizedHistogramClass {
	return getMPSImageNormalizedHistogramClass()
}

type MPSImageNormalizedHistogramClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageNormalizedHistogramClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageNormalizedHistogramClass) Alloc() MPSImageNormalizedHistogram {
	rv := objc.Send[MPSImageNormalizedHistogram](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that computes the normalized histogram of an image.
//
// # Initializers
//
//   - [MPSImageNormalizedHistogram.InitWithDeviceHistogramInfo]
//
// # Instance Properties
//
//   - [MPSImageNormalizedHistogram.ClipRectSource]
//   - [MPSImageNormalizedHistogram.SetClipRectSource]
//   - [MPSImageNormalizedHistogram.HistogramInfo]
//   - [MPSImageNormalizedHistogram.ZeroHistogram]
//   - [MPSImageNormalizedHistogram.SetZeroHistogram]
//
// # Instance Methods
//
//   - [MPSImageNormalizedHistogram.EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset]
//   - [MPSImageNormalizedHistogram.HistogramSizeForSourceFormat]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram
type MPSImageNormalizedHistogram struct {
	MPSKernel
}

// MPSImageNormalizedHistogramFromID constructs a [MPSImageNormalizedHistogram] from an objc.ID.
//
// A filter that computes the normalized histogram of an image.
func MPSImageNormalizedHistogramFromID(id objc.ID) MPSImageNormalizedHistogram {
	return MPSImageNormalizedHistogram{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSImageNormalizedHistogram adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageNormalizedHistogram] class.
//
// # Initializers
//
//   - [IMPSImageNormalizedHistogram.InitWithDeviceHistogramInfo]
//
// # Instance Properties
//
//   - [IMPSImageNormalizedHistogram.ClipRectSource]
//   - [IMPSImageNormalizedHistogram.SetClipRectSource]
//   - [IMPSImageNormalizedHistogram.HistogramInfo]
//   - [IMPSImageNormalizedHistogram.ZeroHistogram]
//   - [IMPSImageNormalizedHistogram.SetZeroHistogram]
//
// # Instance Methods
//
//   - [IMPSImageNormalizedHistogram.EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset]
//   - [IMPSImageNormalizedHistogram.HistogramSizeForSourceFormat]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram
type IMPSImageNormalizedHistogram interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageNormalizedHistogram

	// Topic: Instance Properties

	ClipRectSource() metal.MTLRegion
	SetClipRectSource(value metal.MTLRegion)
	HistogramInfo() MPSImageHistogramInfo
	ZeroHistogram() bool
	SetZeroHistogram(value bool)

	// Topic: Instance Methods

	EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, minmaxTexture metal.MTLTexture, histogram metal.MTLBuffer, histogramOffset uint)
	HistogramSizeForSourceFormat(sourceFormat metal.MTLPixelFormat) uintptr
}

// Init initializes the instance.
func (i MPSImageNormalizedHistogram) Init() MPSImageNormalizedHistogram {
	rv := objc.Send[MPSImageNormalizedHistogram](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageNormalizedHistogram) Autorelease() MPSImageNormalizedHistogram {
	rv := objc.Send[MPSImageNormalizedHistogram](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageNormalizedHistogram creates a new MPSImageNormalizedHistogram instance.
func NewMPSImageNormalizedHistogram() MPSImageNormalizedHistogram {
	class := getMPSImageNormalizedHistogramClass()
	rv := objc.Send[MPSImageNormalizedHistogram](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageNormalizedHistogramWithCoder(aDecoder foundation.INSCoder) MPSImageNormalizedHistogram {
	instance := getMPSImageNormalizedHistogramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageNormalizedHistogramFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram/init(coder:device:)
func NewImageNormalizedHistogramWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageNormalizedHistogram {
	instance := getMPSImageNormalizedHistogramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageNormalizedHistogramFromID(rv)
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
func NewImageNormalizedHistogramWithDevice(device metal.MTLDevice) MPSImageNormalizedHistogram {
	instance := getMPSImageNormalizedHistogramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageNormalizedHistogramFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram/init(device:histogramInfo:)
func NewImageNormalizedHistogramWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageNormalizedHistogram {
	instance := getMPSImageNormalizedHistogramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:histogramInfo:"), device, unsafe.Pointer(histogramInfo))
	return MPSImageNormalizedHistogramFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram/init(device:histogramInfo:)
func (i MPSImageNormalizedHistogram) InitWithDeviceHistogramInfo(device metal.MTLDevice, histogramInfo *MPSImageHistogramInfo) MPSImageNormalizedHistogram {
	rv := objc.Send[MPSImageNormalizedHistogram](i.ID, objc.Sel("initWithDevice:histogramInfo:"), device, unsafe.Pointer(histogramInfo))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram/encode(to:sourceTexture:minmaxTexture:histogram:histogramOffset:)
func (i MPSImageNormalizedHistogram) EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, minmaxTexture metal.MTLTexture, histogram metal.MTLBuffer, histogramOffset uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:minmaxTexture:histogram:histogramOffset:"), commandBuffer, source, minmaxTexture, histogram, histogramOffset)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram/histogramSize(forSourceFormat:)
func (i MPSImageNormalizedHistogram) HistogramSizeForSourceFormat(sourceFormat metal.MTLPixelFormat) uintptr {
	rv := objc.Send[uintptr](i.ID, objc.Sel("histogramSizeForSourceFormat:"), sourceFormat)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram/clipRectSource
func (i MPSImageNormalizedHistogram) ClipRectSource() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](i.ID, objc.Sel("clipRectSource"))
	return metal.MTLRegion(rv)
}
func (i MPSImageNormalizedHistogram) SetClipRectSource(value metal.MTLRegion) {
	objc.Send[struct{}](i.ID, objc.Sel("setClipRectSource:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram/histogramInfo
func (i MPSImageNormalizedHistogram) HistogramInfo() MPSImageHistogramInfo {
	rv := objc.Send[MPSImageHistogramInfo](i.ID, objc.Sel("histogramInfo"))
	return MPSImageHistogramInfo(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram/zeroHistogram
func (i MPSImageNormalizedHistogram) ZeroHistogram() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("zeroHistogram"))
	return rv
}
func (i MPSImageNormalizedHistogram) SetZeroHistogram(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setZeroHistogram:"), value)
}

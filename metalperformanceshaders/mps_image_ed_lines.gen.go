// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageEDLines] class.
var (
	_MPSImageEDLinesClass     MPSImageEDLinesClass
	_MPSImageEDLinesClassOnce sync.Once
)

func getMPSImageEDLinesClass() MPSImageEDLinesClass {
	_MPSImageEDLinesClassOnce.Do(func() {
		_MPSImageEDLinesClass = MPSImageEDLinesClass{class: objc.GetClass("MPSImageEDLines")}
	})
	return _MPSImageEDLinesClass
}

// GetMPSImageEDLinesClass returns the class object for MPSImageEDLines.
func GetMPSImageEDLinesClass() MPSImageEDLinesClass {
	return getMPSImageEDLinesClass()
}

type MPSImageEDLinesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageEDLinesClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageEDLinesClass) Alloc() MPSImageEDLines {
	rv := objc.Send[MPSImageEDLines](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSImageEDLines.InitWithDeviceGaussianSigmaMinLineLengthMaxLinesDetailRatioGradientThresholdLineErrorThresholdMergeLocalityThreshold]
//
// # Instance Properties
//
//   - [MPSImageEDLines.ClipRectSource]
//   - [MPSImageEDLines.SetClipRectSource]
//   - [MPSImageEDLines.DetailRatio]
//   - [MPSImageEDLines.SetDetailRatio]
//   - [MPSImageEDLines.GaussianSigma]
//   - [MPSImageEDLines.GradientThreshold]
//   - [MPSImageEDLines.SetGradientThreshold]
//   - [MPSImageEDLines.LineErrorThreshold]
//   - [MPSImageEDLines.SetLineErrorThreshold]
//   - [MPSImageEDLines.MaxLines]
//   - [MPSImageEDLines.SetMaxLines]
//   - [MPSImageEDLines.MergeLocalityThreshold]
//   - [MPSImageEDLines.SetMergeLocalityThreshold]
//   - [MPSImageEDLines.MinLineLength]
//   - [MPSImageEDLines.SetMinLineLength]
//
// # Instance Methods
//
//   - [MPSImageEDLines.EncodeToCommandBufferSourceTextureDestinationTextureEndpointBufferEndpointOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines
type MPSImageEDLines struct {
	MPSKernel
}

// MPSImageEDLinesFromID constructs a [MPSImageEDLines] from an objc.ID.
func MPSImageEDLinesFromID(id objc.ID) MPSImageEDLines {
	return MPSImageEDLines{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSImageEDLines adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageEDLines] class.
//
// # Initializers
//
//   - [IMPSImageEDLines.InitWithDeviceGaussianSigmaMinLineLengthMaxLinesDetailRatioGradientThresholdLineErrorThresholdMergeLocalityThreshold]
//
// # Instance Properties
//
//   - [IMPSImageEDLines.ClipRectSource]
//   - [IMPSImageEDLines.SetClipRectSource]
//   - [IMPSImageEDLines.DetailRatio]
//   - [IMPSImageEDLines.SetDetailRatio]
//   - [IMPSImageEDLines.GaussianSigma]
//   - [IMPSImageEDLines.GradientThreshold]
//   - [IMPSImageEDLines.SetGradientThreshold]
//   - [IMPSImageEDLines.LineErrorThreshold]
//   - [IMPSImageEDLines.SetLineErrorThreshold]
//   - [IMPSImageEDLines.MaxLines]
//   - [IMPSImageEDLines.SetMaxLines]
//   - [IMPSImageEDLines.MergeLocalityThreshold]
//   - [IMPSImageEDLines.SetMergeLocalityThreshold]
//   - [IMPSImageEDLines.MinLineLength]
//   - [IMPSImageEDLines.SetMinLineLength]
//
// # Instance Methods
//
//   - [IMPSImageEDLines.EncodeToCommandBufferSourceTextureDestinationTextureEndpointBufferEndpointOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines
type IMPSImageEDLines interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceGaussianSigmaMinLineLengthMaxLinesDetailRatioGradientThresholdLineErrorThresholdMergeLocalityThreshold(device metal.MTLDevice, gaussianSigma float32, minLineLength uint16, maxLines uint, detailRatio uint16, gradientThreshold float32, lineErrorThreshold float32, mergeLocalityThreshold float32) MPSImageEDLines

	// Topic: Instance Properties

	ClipRectSource() metal.MTLRegion
	SetClipRectSource(value metal.MTLRegion)
	DetailRatio() uint16
	SetDetailRatio(value uint16)
	GaussianSigma() float32
	GradientThreshold() float32
	SetGradientThreshold(value float32)
	LineErrorThreshold() float32
	SetLineErrorThreshold(value float32)
	MaxLines() uint
	SetMaxLines(value uint)
	MergeLocalityThreshold() float32
	SetMergeLocalityThreshold(value float32)
	MinLineLength() uint16
	SetMinLineLength(value uint16)

	// Topic: Instance Methods

	EncodeToCommandBufferSourceTextureDestinationTextureEndpointBufferEndpointOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, dest metal.MTLTexture, endpointBuffer metal.MTLBuffer, endpointOffset uint)
}

// Init initializes the instance.
func (i MPSImageEDLines) Init() MPSImageEDLines {
	rv := objc.Send[MPSImageEDLines](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageEDLines) Autorelease() MPSImageEDLines {
	rv := objc.Send[MPSImageEDLines](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageEDLines creates a new MPSImageEDLines instance.
func NewMPSImageEDLines() MPSImageEDLines {
	class := getMPSImageEDLinesClass()
	rv := objc.Send[MPSImageEDLines](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageEDLinesWithCoder(aDecoder foundation.INSCoder) MPSImageEDLines {
	instance := getMPSImageEDLinesClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageEDLinesFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/init(coder:device:)
func NewImageEDLinesWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageEDLines {
	instance := getMPSImageEDLinesClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageEDLinesFromID(rv)
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
func NewImageEDLinesWithDevice(device metal.MTLDevice) MPSImageEDLines {
	instance := getMPSImageEDLinesClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageEDLinesFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/init(device:gaussianSigma:minLineLength:maxLines:detailRatio:gradientThreshold:lineErrorThreshold:mergeLocalityThreshold:)
func NewImageEDLinesWithDeviceGaussianSigmaMinLineLengthMaxLinesDetailRatioGradientThresholdLineErrorThresholdMergeLocalityThreshold(device metal.MTLDevice, gaussianSigma float32, minLineLength uint16, maxLines uint, detailRatio uint16, gradientThreshold float32, lineErrorThreshold float32, mergeLocalityThreshold float32) MPSImageEDLines {
	instance := getMPSImageEDLinesClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:gaussianSigma:minLineLength:maxLines:detailRatio:gradientThreshold:lineErrorThreshold:mergeLocalityThreshold:"), device, gaussianSigma, minLineLength, maxLines, detailRatio, gradientThreshold, lineErrorThreshold, mergeLocalityThreshold)
	return MPSImageEDLinesFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/init(device:gaussianSigma:minLineLength:maxLines:detailRatio:gradientThreshold:lineErrorThreshold:mergeLocalityThreshold:)
func (i MPSImageEDLines) InitWithDeviceGaussianSigmaMinLineLengthMaxLinesDetailRatioGradientThresholdLineErrorThresholdMergeLocalityThreshold(device metal.MTLDevice, gaussianSigma float32, minLineLength uint16, maxLines uint, detailRatio uint16, gradientThreshold float32, lineErrorThreshold float32, mergeLocalityThreshold float32) MPSImageEDLines {
	rv := objc.Send[MPSImageEDLines](i.ID, objc.Sel("initWithDevice:gaussianSigma:minLineLength:maxLines:detailRatio:gradientThreshold:lineErrorThreshold:mergeLocalityThreshold:"), device, gaussianSigma, minLineLength, maxLines, detailRatio, gradientThreshold, lineErrorThreshold, mergeLocalityThreshold)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/encode(to:sourceTexture:destinationTexture:endpointBuffer:endpointOffset:)
func (i MPSImageEDLines) EncodeToCommandBufferSourceTextureDestinationTextureEndpointBufferEndpointOffset(commandBuffer metal.MTLCommandBuffer, source metal.MTLTexture, dest metal.MTLTexture, endpointBuffer metal.MTLBuffer, endpointOffset uint) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:endpointBuffer:endpointOffset:"), commandBuffer, source, dest, endpointBuffer, endpointOffset)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/clipRectSource
func (i MPSImageEDLines) ClipRectSource() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](i.ID, objc.Sel("clipRectSource"))
	return metal.MTLRegion(rv)
}
func (i MPSImageEDLines) SetClipRectSource(value metal.MTLRegion) {
	objc.Send[struct{}](i.ID, objc.Sel("setClipRectSource:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/detailRatio
func (i MPSImageEDLines) DetailRatio() uint16 {
	rv := objc.Send[uint16](i.ID, objc.Sel("detailRatio"))
	return rv
}
func (i MPSImageEDLines) SetDetailRatio(value uint16) {
	objc.Send[struct{}](i.ID, objc.Sel("setDetailRatio:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/gaussianSigma
func (i MPSImageEDLines) GaussianSigma() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("gaussianSigma"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/gradientThreshold
func (i MPSImageEDLines) GradientThreshold() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("gradientThreshold"))
	return rv
}
func (i MPSImageEDLines) SetGradientThreshold(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setGradientThreshold:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/lineErrorThreshold
func (i MPSImageEDLines) LineErrorThreshold() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("lineErrorThreshold"))
	return rv
}
func (i MPSImageEDLines) SetLineErrorThreshold(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setLineErrorThreshold:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/maxLines
func (i MPSImageEDLines) MaxLines() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxLines"))
	return rv
}
func (i MPSImageEDLines) SetMaxLines(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxLines:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/mergeLocalityThreshold
func (i MPSImageEDLines) MergeLocalityThreshold() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("mergeLocalityThreshold"))
	return rv
}
func (i MPSImageEDLines) SetMergeLocalityThreshold(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setMergeLocalityThreshold:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines/minLineLength
func (i MPSImageEDLines) MinLineLength() uint16 {
	rv := objc.Send[uint16](i.ID, objc.Sel("minLineLength"))
	return rv
}
func (i MPSImageEDLines) SetMinLineLength(value uint16) {
	objc.Send[struct{}](i.ID, objc.Sel("setMinLineLength:"), value)
}

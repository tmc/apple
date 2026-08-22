// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSSVGF] class.
var (
	_MPSSVGFClass     MPSSVGFClass
	_MPSSVGFClassOnce sync.Once
)

func getMPSSVGFClass() MPSSVGFClass {
	_MPSSVGFClassOnce.Do(func() {
		_MPSSVGFClass = MPSSVGFClass{class: objc.GetClass("MPSSVGF")}
	})
	return _MPSSVGFClass
}

// GetMPSSVGFClass returns the class object for MPSSVGF.
func GetMPSSVGFClass() MPSSVGFClass {
	return getMPSSVGFClass()
}

type MPSSVGFClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSSVGFClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSSVGFClass) Alloc() MPSSVGF {
	rv := objc.Send[MPSSVGF](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSSVGF.BilateralFilterRadius]
//   - [MPSSVGF.SetBilateralFilterRadius]
//   - [MPSSVGF.BilateralFilterSigma]
//   - [MPSSVGF.SetBilateralFilterSigma]
//   - [MPSSVGF.ChannelCount]
//   - [MPSSVGF.SetChannelCount]
//   - [MPSSVGF.ChannelCount2]
//   - [MPSSVGF.SetChannelCount2]
//   - [MPSSVGF.DepthWeight]
//   - [MPSSVGF.SetDepthWeight]
//   - [MPSSVGF.LuminanceWeight]
//   - [MPSSVGF.SetLuminanceWeight]
//   - [MPSSVGF.MinimumFramesForVarianceEstimation]
//   - [MPSSVGF.SetMinimumFramesForVarianceEstimation]
//   - [MPSSVGF.NormalWeight]
//   - [MPSSVGF.SetNormalWeight]
//   - [MPSSVGF.ReprojectionThreshold]
//   - [MPSSVGF.SetReprojectionThreshold]
//   - [MPSSVGF.TemporalReprojectionBlendFactor]
//   - [MPSSVGF.SetTemporalReprojectionBlendFactor]
//   - [MPSSVGF.TemporalWeighting]
//   - [MPSSVGF.SetTemporalWeighting]
//   - [MPSSVGF.VarianceEstimationRadius]
//   - [MPSSVGF.SetVarianceEstimationRadius]
//   - [MPSSVGF.VarianceEstimationSigma]
//   - [MPSSVGF.SetVarianceEstimationSigma]
//   - [MPSSVGF.VariancePrefilterRadius]
//   - [MPSSVGF.SetVariancePrefilterRadius]
//   - [MPSSVGF.VariancePrefilterSigma]
//   - [MPSSVGF.SetVariancePrefilterSigma]
//
// # Instance Methods
//
//   - [MPSSVGF.EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureDepthNormalTexture]
//   - [MPSSVGF.EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureSourceTexture2DestinationTexture2DepthNormalTexture]
//   - [MPSSVGF.EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureFrameCountTextureDepthNormalTexture]
//   - [MPSSVGF.EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureSourceTexture2LuminanceMomentsTexture2DestinationTexture2FrameCountTextureDepthNormalTexture]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF
type MPSSVGF struct {
	MPSKernel
}

// MPSSVGFFromID constructs a [MPSSVGF] from an objc.ID.
func MPSSVGFFromID(id objc.ID) MPSSVGF {
	return MPSSVGF{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSSVGF adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSSVGF] class.
//
// # Instance Properties
//
//   - [IMPSSVGF.BilateralFilterRadius]
//   - [IMPSSVGF.SetBilateralFilterRadius]
//   - [IMPSSVGF.BilateralFilterSigma]
//   - [IMPSSVGF.SetBilateralFilterSigma]
//   - [IMPSSVGF.ChannelCount]
//   - [IMPSSVGF.SetChannelCount]
//   - [IMPSSVGF.ChannelCount2]
//   - [IMPSSVGF.SetChannelCount2]
//   - [IMPSSVGF.DepthWeight]
//   - [IMPSSVGF.SetDepthWeight]
//   - [IMPSSVGF.LuminanceWeight]
//   - [IMPSSVGF.SetLuminanceWeight]
//   - [IMPSSVGF.MinimumFramesForVarianceEstimation]
//   - [IMPSSVGF.SetMinimumFramesForVarianceEstimation]
//   - [IMPSSVGF.NormalWeight]
//   - [IMPSSVGF.SetNormalWeight]
//   - [IMPSSVGF.ReprojectionThreshold]
//   - [IMPSSVGF.SetReprojectionThreshold]
//   - [IMPSSVGF.TemporalReprojectionBlendFactor]
//   - [IMPSSVGF.SetTemporalReprojectionBlendFactor]
//   - [IMPSSVGF.TemporalWeighting]
//   - [IMPSSVGF.SetTemporalWeighting]
//   - [IMPSSVGF.VarianceEstimationRadius]
//   - [IMPSSVGF.SetVarianceEstimationRadius]
//   - [IMPSSVGF.VarianceEstimationSigma]
//   - [IMPSSVGF.SetVarianceEstimationSigma]
//   - [IMPSSVGF.VariancePrefilterRadius]
//   - [IMPSSVGF.SetVariancePrefilterRadius]
//   - [IMPSSVGF.VariancePrefilterSigma]
//   - [IMPSSVGF.SetVariancePrefilterSigma]
//
// # Instance Methods
//
//   - [IMPSSVGF.EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureDepthNormalTexture]
//   - [IMPSSVGF.EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureSourceTexture2DestinationTexture2DepthNormalTexture]
//   - [IMPSSVGF.EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureFrameCountTextureDepthNormalTexture]
//   - [IMPSSVGF.EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureSourceTexture2LuminanceMomentsTexture2DestinationTexture2FrameCountTextureDepthNormalTexture]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF
type IMPSSVGF interface {
	IMPSKernel

	// Topic: Instance Properties

	BilateralFilterRadius() uint
	SetBilateralFilterRadius(value uint)
	BilateralFilterSigma() float32
	SetBilateralFilterSigma(value float32)
	ChannelCount() uint
	SetChannelCount(value uint)
	ChannelCount2() uint
	SetChannelCount2(value uint)
	DepthWeight() float32
	SetDepthWeight(value float32)
	LuminanceWeight() float32
	SetLuminanceWeight(value float32)
	MinimumFramesForVarianceEstimation() uint
	SetMinimumFramesForVarianceEstimation(value uint)
	NormalWeight() float32
	SetNormalWeight(value float32)
	ReprojectionThreshold() float32
	SetReprojectionThreshold(value float32)
	TemporalReprojectionBlendFactor() float32
	SetTemporalReprojectionBlendFactor(value float32)
	TemporalWeighting() MPSTemporalWeighting
	SetTemporalWeighting(value MPSTemporalWeighting)
	VarianceEstimationRadius() uint
	SetVarianceEstimationRadius(value uint)
	VarianceEstimationSigma() float32
	SetVarianceEstimationSigma(value float32)
	VariancePrefilterRadius() uint
	SetVariancePrefilterRadius(value uint)
	VariancePrefilterSigma() float32
	SetVariancePrefilterSigma(value float32)

	// Topic: Instance Methods

	EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureDepthNormalTexture(commandBuffer metal.MTLCommandBuffer, stepDistance uint, sourceTexture metal.MTLTexture, destinationTexture metal.MTLTexture, depthNormalTexture metal.MTLTexture)
	EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureSourceTexture2DestinationTexture2DepthNormalTexture(commandBuffer metal.MTLCommandBuffer, stepDistance uint, sourceTexture metal.MTLTexture, destinationTexture metal.MTLTexture, sourceTexture2 metal.MTLTexture, destinationTexture2 metal.MTLTexture, depthNormalTexture metal.MTLTexture)
	EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureFrameCountTextureDepthNormalTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, luminanceMomentsTexture metal.MTLTexture, destinationTexture metal.MTLTexture, frameCountTexture metal.MTLTexture, depthNormalTexture metal.MTLTexture)
	EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureSourceTexture2LuminanceMomentsTexture2DestinationTexture2FrameCountTextureDepthNormalTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, luminanceMomentsTexture metal.MTLTexture, destinationTexture metal.MTLTexture, sourceTexture2 metal.MTLTexture, luminanceMomentsTexture2 metal.MTLTexture, destinationTexture2 metal.MTLTexture, frameCountTexture metal.MTLTexture, depthNormalTexture metal.MTLTexture)
}

// Init initializes the instance.
func (s MPSSVGF) Init() MPSSVGF {
	rv := objc.Send[MPSSVGF](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSSVGF) Autorelease() MPSSVGF {
	rv := objc.Send[MPSSVGF](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSSVGF creates a new MPSSVGF instance.
func NewMPSSVGF() MPSSVGF {
	class := getMPSSVGFClass()
	rv := objc.Send[MPSSVGF](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewSVGFWithCoder(aDecoder foundation.INSCoder) MPSSVGF {
	instance := getMPSSVGFClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSSVGFFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/init(coder:device:)
func NewSVGFWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSSVGF {
	instance := getMPSSVGFClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSSVGFFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/init(device:)
func NewSVGFWithDevice(device metal.MTLDevice) MPSSVGF {
	instance := getMPSSVGFClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSSVGFFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/encodeBilateralFilter(to:stepDistance:sourceTexture:destinationTexture:depthNormalTexture:)
func (s MPSSVGF) EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureDepthNormalTexture(commandBuffer metal.MTLCommandBuffer, stepDistance uint, sourceTexture metal.MTLTexture, destinationTexture metal.MTLTexture, depthNormalTexture metal.MTLTexture) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeBilateralFilterToCommandBuffer:stepDistance:sourceTexture:destinationTexture:depthNormalTexture:"), commandBuffer, stepDistance, sourceTexture, destinationTexture, depthNormalTexture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/encodeBilateralFilter(to:stepDistance:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:depthNormalTexture:)
func (s MPSSVGF) EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureSourceTexture2DestinationTexture2DepthNormalTexture(commandBuffer metal.MTLCommandBuffer, stepDistance uint, sourceTexture metal.MTLTexture, destinationTexture metal.MTLTexture, sourceTexture2 metal.MTLTexture, destinationTexture2 metal.MTLTexture, depthNormalTexture metal.MTLTexture) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeBilateralFilterToCommandBuffer:stepDistance:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:depthNormalTexture:"), commandBuffer, stepDistance, sourceTexture, destinationTexture, sourceTexture2, destinationTexture2, depthNormalTexture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/encodeVarianceEstimation(to:sourceTexture:luminanceMomentsTexture:destinationTexture:frameCount:depthNormalTexture:)
func (s MPSSVGF) EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureFrameCountTextureDepthNormalTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, luminanceMomentsTexture metal.MTLTexture, destinationTexture metal.MTLTexture, frameCountTexture metal.MTLTexture, depthNormalTexture metal.MTLTexture) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeVarianceEstimationToCommandBuffer:sourceTexture:luminanceMomentsTexture:destinationTexture:frameCountTexture:depthNormalTexture:"), commandBuffer, sourceTexture, luminanceMomentsTexture, destinationTexture, frameCountTexture, depthNormalTexture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/encodeVarianceEstimation(to:sourceTexture:luminanceMomentsTexture:destinationTexture:sourceTexture2:luminanceMomentsTexture2:destinationTexture2:frameCount:depthNormalTexture:)
func (s MPSSVGF) EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureSourceTexture2LuminanceMomentsTexture2DestinationTexture2FrameCountTextureDepthNormalTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, luminanceMomentsTexture metal.MTLTexture, destinationTexture metal.MTLTexture, sourceTexture2 metal.MTLTexture, luminanceMomentsTexture2 metal.MTLTexture, destinationTexture2 metal.MTLTexture, frameCountTexture metal.MTLTexture, depthNormalTexture metal.MTLTexture) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeVarianceEstimationToCommandBuffer:sourceTexture:luminanceMomentsTexture:destinationTexture:sourceTexture2:luminanceMomentsTexture2:destinationTexture2:frameCountTexture:depthNormalTexture:"), commandBuffer, sourceTexture, luminanceMomentsTexture, destinationTexture, sourceTexture2, luminanceMomentsTexture2, destinationTexture2, frameCountTexture, depthNormalTexture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/bilateralFilterRadius
func (s MPSSVGF) BilateralFilterRadius() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("bilateralFilterRadius"))
	return rv
}
func (s MPSSVGF) SetBilateralFilterRadius(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setBilateralFilterRadius:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/bilateralFilterSigma
func (s MPSSVGF) BilateralFilterSigma() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("bilateralFilterSigma"))
	return rv
}
func (s MPSSVGF) SetBilateralFilterSigma(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setBilateralFilterSigma:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/channelCount
func (s MPSSVGF) ChannelCount() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("channelCount"))
	return rv
}
func (s MPSSVGF) SetChannelCount(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setChannelCount:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/channelCount2
func (s MPSSVGF) ChannelCount2() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("channelCount2"))
	return rv
}
func (s MPSSVGF) SetChannelCount2(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setChannelCount2:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/depthWeight
func (s MPSSVGF) DepthWeight() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("depthWeight"))
	return rv
}
func (s MPSSVGF) SetDepthWeight(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setDepthWeight:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/luminanceWeight
func (s MPSSVGF) LuminanceWeight() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("luminanceWeight"))
	return rv
}
func (s MPSSVGF) SetLuminanceWeight(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setLuminanceWeight:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/minimumFramesForVarianceEstimation
func (s MPSSVGF) MinimumFramesForVarianceEstimation() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("minimumFramesForVarianceEstimation"))
	return rv
}
func (s MPSSVGF) SetMinimumFramesForVarianceEstimation(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinimumFramesForVarianceEstimation:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/normalWeight
func (s MPSSVGF) NormalWeight() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("normalWeight"))
	return rv
}
func (s MPSSVGF) SetNormalWeight(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setNormalWeight:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/reprojectionThreshold
func (s MPSSVGF) ReprojectionThreshold() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("reprojectionThreshold"))
	return rv
}
func (s MPSSVGF) SetReprojectionThreshold(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setReprojectionThreshold:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/temporalReprojectionBlendFactor
func (s MPSSVGF) TemporalReprojectionBlendFactor() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("temporalReprojectionBlendFactor"))
	return rv
}
func (s MPSSVGF) SetTemporalReprojectionBlendFactor(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setTemporalReprojectionBlendFactor:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/temporalWeighting
func (s MPSSVGF) TemporalWeighting() MPSTemporalWeighting {
	rv := objc.Send[MPSTemporalWeighting](s.ID, objc.Sel("temporalWeighting"))
	return MPSTemporalWeighting(rv)
}
func (s MPSSVGF) SetTemporalWeighting(value MPSTemporalWeighting) {
	objc.Send[struct{}](s.ID, objc.Sel("setTemporalWeighting:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/varianceEstimationRadius
func (s MPSSVGF) VarianceEstimationRadius() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("varianceEstimationRadius"))
	return rv
}
func (s MPSSVGF) SetVarianceEstimationRadius(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setVarianceEstimationRadius:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/varianceEstimationSigma
func (s MPSSVGF) VarianceEstimationSigma() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("varianceEstimationSigma"))
	return rv
}
func (s MPSSVGF) SetVarianceEstimationSigma(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setVarianceEstimationSigma:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/variancePrefilterRadius
func (s MPSSVGF) VariancePrefilterRadius() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("variancePrefilterRadius"))
	return rv
}
func (s MPSSVGF) SetVariancePrefilterRadius(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setVariancePrefilterRadius:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/variancePrefilterSigma
func (s MPSSVGF) VariancePrefilterSigma() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("variancePrefilterSigma"))
	return rv
}
func (s MPSSVGF) SetVariancePrefilterSigma(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setVariancePrefilterSigma:"), value)
}

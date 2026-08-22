// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSTemporalAA] class.
var (
	_MPSTemporalAAClass     MPSTemporalAAClass
	_MPSTemporalAAClassOnce sync.Once
)

func getMPSTemporalAAClass() MPSTemporalAAClass {
	_MPSTemporalAAClassOnce.Do(func() {
		_MPSTemporalAAClass = MPSTemporalAAClass{class: objc.GetClass("MPSTemporalAA")}
	})
	return _MPSTemporalAAClass
}

// GetMPSTemporalAAClass returns the class object for MPSTemporalAA.
func GetMPSTemporalAAClass() MPSTemporalAAClass {
	return getMPSTemporalAAClass()
}

type MPSTemporalAAClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSTemporalAAClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSTemporalAAClass) Alloc() MPSTemporalAA {
	rv := objc.Send[MPSTemporalAA](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSTemporalAA.BlendFactor]
//   - [MPSTemporalAA.SetBlendFactor]
//
// # Instance Methods
//
//   - [MPSTemporalAA.EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalAA
type MPSTemporalAA struct {
	MPSKernel
}

// MPSTemporalAAFromID constructs a [MPSTemporalAA] from an objc.ID.
func MPSTemporalAAFromID(id objc.ID) MPSTemporalAA {
	return MPSTemporalAA{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSTemporalAA adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSTemporalAA] class.
//
// # Instance Properties
//
//   - [IMPSTemporalAA.BlendFactor]
//   - [IMPSTemporalAA.SetBlendFactor]
//
// # Instance Methods
//
//   - [IMPSTemporalAA.EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalAA
type IMPSTemporalAA interface {
	IMPSKernel

	// Topic: Instance Properties

	BlendFactor() float32
	SetBlendFactor(value float32)

	// Topic: Instance Methods

	EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, previousTexture metal.MTLTexture, destinationTexture metal.MTLTexture, motionVectorTexture metal.MTLTexture, depthTexture metal.MTLTexture)
}

// Init initializes the instance.
func (t MPSTemporalAA) Init() MPSTemporalAA {
	rv := objc.Send[MPSTemporalAA](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MPSTemporalAA) Autorelease() MPSTemporalAA {
	rv := objc.Send[MPSTemporalAA](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSTemporalAA creates a new MPSTemporalAA instance.
func NewMPSTemporalAA() MPSTemporalAA {
	class := getMPSTemporalAAClass()
	rv := objc.Send[MPSTemporalAA](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewTemporalAAWithCoder(aDecoder foundation.INSCoder) MPSTemporalAA {
	instance := getMPSTemporalAAClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSTemporalAAFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalAA/init(coder:device:)
func NewTemporalAAWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSTemporalAA {
	instance := getMPSTemporalAAClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSTemporalAAFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalAA/init(device:)
func NewTemporalAAWithDevice(device metal.MTLDevice) MPSTemporalAA {
	instance := getMPSTemporalAAClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSTemporalAAFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalAA/encode(to:sourceTexture:previousTexture:destinationTexture:motionVectorTexture:depthTexture:)
func (t MPSTemporalAA) EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, previousTexture metal.MTLTexture, destinationTexture metal.MTLTexture, motionVectorTexture metal.MTLTexture, depthTexture metal.MTLTexture) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:previousTexture:destinationTexture:motionVectorTexture:depthTexture:"), commandBuffer, sourceTexture, previousTexture, destinationTexture, motionVectorTexture, depthTexture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalAA/blendFactor
func (t MPSTemporalAA) BlendFactor() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("blendFactor"))
	return rv
}
func (t MPSTemporalAA) SetBlendFactor(value float32) {
	objc.Send[struct{}](t.ID, objc.Sel("setBlendFactor:"), value)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSSVGFDenoiser] class.
var (
	_MPSSVGFDenoiserClass     MPSSVGFDenoiserClass
	_MPSSVGFDenoiserClassOnce sync.Once
)

func getMPSSVGFDenoiserClass() MPSSVGFDenoiserClass {
	_MPSSVGFDenoiserClassOnce.Do(func() {
		_MPSSVGFDenoiserClass = MPSSVGFDenoiserClass{class: objc.GetClass("MPSSVGFDenoiser")}
	})
	return _MPSSVGFDenoiserClass
}

// GetMPSSVGFDenoiserClass returns the class object for MPSSVGFDenoiser.
func GetMPSSVGFDenoiserClass() MPSSVGFDenoiserClass {
	return getMPSSVGFDenoiserClass()
}

type MPSSVGFDenoiserClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSSVGFDenoiserClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSSVGFDenoiserClass) Alloc() MPSSVGFDenoiser {
	rv := objc.Send[MPSSVGFDenoiser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSSVGFDenoiser.InitWithSVGFTextureAllocator]
//   - [MPSSVGFDenoiser.InitWithDevice]
//
// # Instance Properties
//
//   - [MPSSVGFDenoiser.BilateralFilterIterations]
//   - [MPSSVGFDenoiser.SetBilateralFilterIterations]
//   - [MPSSVGFDenoiser.Svgf]
//   - [MPSSVGFDenoiser.TextureAllocator]
//
// # Instance Methods
//
//   - [MPSSVGFDenoiser.ClearTemporalHistory]
//   - [MPSSVGFDenoiser.EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture]
//   - [MPSSVGFDenoiser.EncodeToCommandBufferSourceTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture]
//   - [MPSSVGFDenoiser.ReleaseTemporaryTextures]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser
type MPSSVGFDenoiser struct {
	objectivec.Object
}

// MPSSVGFDenoiserFromID constructs a [MPSSVGFDenoiser] from an objc.ID.
func MPSSVGFDenoiserFromID(id objc.ID) MPSSVGFDenoiser {
	return MPSSVGFDenoiser{objectivec.Object{ID: id}}
}

// NOTE: MPSSVGFDenoiser adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSSVGFDenoiser] class.
//
// # Initializers
//
//   - [IMPSSVGFDenoiser.InitWithSVGFTextureAllocator]
//   - [IMPSSVGFDenoiser.InitWithDevice]
//
// # Instance Properties
//
//   - [IMPSSVGFDenoiser.BilateralFilterIterations]
//   - [IMPSSVGFDenoiser.SetBilateralFilterIterations]
//   - [IMPSSVGFDenoiser.Svgf]
//   - [IMPSSVGFDenoiser.TextureAllocator]
//
// # Instance Methods
//
//   - [IMPSSVGFDenoiser.ClearTemporalHistory]
//   - [IMPSSVGFDenoiser.EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture]
//   - [IMPSSVGFDenoiser.EncodeToCommandBufferSourceTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture]
//   - [IMPSSVGFDenoiser.ReleaseTemporaryTextures]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser
type IMPSSVGFDenoiser interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithSVGFTextureAllocator(svgf IMPSSVGF, textureAllocator MPSSVGFTextureAllocator) MPSSVGFDenoiser
	InitWithDevice(device metal.MTLDevice) MPSSVGFDenoiser

	// Topic: Instance Properties

	BilateralFilterIterations() uint
	SetBilateralFilterIterations(value uint)
	Svgf() IMPSSVGF
	TextureAllocator() MPSSVGFTextureAllocator

	// Topic: Instance Methods

	ClearTemporalHistory()
	EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, destinationTexture metal.MTLTexture, sourceTexture2 metal.MTLTexture, destinationTexture2 metal.MTLTexture, motionVectorTexture metal.MTLTexture, depthNormalTexture metal.MTLTexture, previousDepthNormalTexture metal.MTLTexture)
	EncodeToCommandBufferSourceTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, motionVectorTexture metal.MTLTexture, depthNormalTexture metal.MTLTexture, previousDepthNormalTexture metal.MTLTexture) metal.MTLTexture
	ReleaseTemporaryTextures()
}

// Init initializes the instance.
func (s MPSSVGFDenoiser) Init() MPSSVGFDenoiser {
	rv := objc.Send[MPSSVGFDenoiser](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSSVGFDenoiser) Autorelease() MPSSVGFDenoiser {
	rv := objc.Send[MPSSVGFDenoiser](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSSVGFDenoiser creates a new MPSSVGFDenoiser instance.
func NewMPSSVGFDenoiser() MPSSVGFDenoiser {
	class := getMPSSVGFDenoiserClass()
	rv := objc.Send[MPSSVGFDenoiser](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/init(device:)
func NewSVGFDenoiserWithDevice(device metal.MTLDevice) MPSSVGFDenoiser {
	instance := getMPSSVGFDenoiserClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSSVGFDenoiserFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/init(SVGF:textureAllocator:)
func NewSVGFDenoiserWithSVGFTextureAllocator(svgf IMPSSVGF, textureAllocator MPSSVGFTextureAllocator) MPSSVGFDenoiser {
	instance := getMPSSVGFDenoiserClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSVGF:textureAllocator:"), svgf, textureAllocator)
	return MPSSVGFDenoiserFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/init(SVGF:textureAllocator:)
func (s MPSSVGFDenoiser) InitWithSVGFTextureAllocator(svgf IMPSSVGF, textureAllocator MPSSVGFTextureAllocator) MPSSVGFDenoiser {
	rv := objc.Send[MPSSVGFDenoiser](s.ID, objc.Sel("initWithSVGF:textureAllocator:"), svgf, textureAllocator)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/init(device:)
func (s MPSSVGFDenoiser) InitWithDevice(device metal.MTLDevice) MPSSVGFDenoiser {
	rv := objc.Send[MPSSVGFDenoiser](s.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/clearTemporalHistory()
func (s MPSSVGFDenoiser) ClearTemporalHistory() {
	objc.Send[objc.ID](s.ID, objc.Sel("clearTemporalHistory"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/encode(commandBuffer:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:)
func (s MPSSVGFDenoiser) EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, destinationTexture metal.MTLTexture, sourceTexture2 metal.MTLTexture, destinationTexture2 metal.MTLTexture, motionVectorTexture metal.MTLTexture, depthNormalTexture metal.MTLTexture, previousDepthNormalTexture metal.MTLTexture) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), commandBuffer, sourceTexture, destinationTexture, sourceTexture2, destinationTexture2, motionVectorTexture, depthNormalTexture, previousDepthNormalTexture)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/encode(commandBuffer:sourceTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:)
func (s MPSSVGFDenoiser) EncodeToCommandBufferSourceTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.MTLCommandBuffer, sourceTexture metal.MTLTexture, motionVectorTexture metal.MTLTexture, depthNormalTexture metal.MTLTexture, previousDepthNormalTexture metal.MTLTexture) metal.MTLTexture {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), commandBuffer, sourceTexture, motionVectorTexture, depthNormalTexture, previousDepthNormalTexture)
	return metal.MTLTextureObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/releaseTemporaryTextures()
func (s MPSSVGFDenoiser) ReleaseTemporaryTextures() {
	objc.Send[objc.ID](s.ID, objc.Sel("releaseTemporaryTextures"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/bilateralFilterIterations
func (s MPSSVGFDenoiser) BilateralFilterIterations() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("bilateralFilterIterations"))
	return rv
}
func (s MPSSVGFDenoiser) SetBilateralFilterIterations(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setBilateralFilterIterations:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/svgf
func (s MPSSVGFDenoiser) Svgf() IMPSSVGF {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("svgf"))
	return MPSSVGFFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/textureAllocator
func (s MPSSVGFDenoiser) TextureAllocator() MPSSVGFTextureAllocator {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("textureAllocator"))
	return MPSSVGFTextureAllocatorObjectFromID(rv)
}

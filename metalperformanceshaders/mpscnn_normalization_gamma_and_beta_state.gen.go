// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNNormalizationGammaAndBetaState] class.
var (
	_MPSCNNNormalizationGammaAndBetaStateClass     MPSCNNNormalizationGammaAndBetaStateClass
	_MPSCNNNormalizationGammaAndBetaStateClassOnce sync.Once
)

func getMPSCNNNormalizationGammaAndBetaStateClass() MPSCNNNormalizationGammaAndBetaStateClass {
	_MPSCNNNormalizationGammaAndBetaStateClassOnce.Do(func() {
		_MPSCNNNormalizationGammaAndBetaStateClass = MPSCNNNormalizationGammaAndBetaStateClass{class: objc.GetClass("MPSCNNNormalizationGammaAndBetaState")}
	})
	return _MPSCNNNormalizationGammaAndBetaStateClass
}

// GetMPSCNNNormalizationGammaAndBetaStateClass returns the class object for MPSCNNNormalizationGammaAndBetaState.
func GetMPSCNNNormalizationGammaAndBetaStateClass() MPSCNNNormalizationGammaAndBetaStateClass {
	return getMPSCNNNormalizationGammaAndBetaStateClass()
}

type MPSCNNNormalizationGammaAndBetaStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNormalizationGammaAndBetaStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNormalizationGammaAndBetaStateClass) Alloc() MPSCNNNormalizationGammaAndBetaState {
	rv := objc.Send[MPSCNNNormalizationGammaAndBetaState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that stores gamma and beta terms used to apply a scale and bias
// in instance- or batch-normalization operations.
//
// # Initializers
//
//   - [MPSCNNNormalizationGammaAndBetaState.InitWithGammaBeta]
//
// # Instance Properties
//
//   - [MPSCNNNormalizationGammaAndBetaState.Beta]
//   - [MPSCNNNormalizationGammaAndBetaState.Gamma]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationGammaAndBetaState
type MPSCNNNormalizationGammaAndBetaState struct {
	MPSState
}

// MPSCNNNormalizationGammaAndBetaStateFromID constructs a [MPSCNNNormalizationGammaAndBetaState] from an objc.ID.
//
// An object that stores gamma and beta terms used to apply a scale and bias
// in instance- or batch-normalization operations.
func MPSCNNNormalizationGammaAndBetaStateFromID(id objc.ID) MPSCNNNormalizationGammaAndBetaState {
	return MPSCNNNormalizationGammaAndBetaState{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSCNNNormalizationGammaAndBetaState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNormalizationGammaAndBetaState] class.
//
// # Initializers
//
//   - [IMPSCNNNormalizationGammaAndBetaState.InitWithGammaBeta]
//
// # Instance Properties
//
//   - [IMPSCNNNormalizationGammaAndBetaState.Beta]
//   - [IMPSCNNNormalizationGammaAndBetaState.Gamma]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationGammaAndBetaState
type IMPSCNNNormalizationGammaAndBetaState interface {
	IMPSState

	// Topic: Initializers

	InitWithGammaBeta(gamma metal.MTLBuffer, beta metal.MTLBuffer) MPSCNNNormalizationGammaAndBetaState

	// Topic: Instance Properties

	Beta() metal.MTLBuffer
	Gamma() metal.MTLBuffer
}

// Init initializes the instance.
func (c MPSCNNNormalizationGammaAndBetaState) Init() MPSCNNNormalizationGammaAndBetaState {
	rv := objc.Send[MPSCNNNormalizationGammaAndBetaState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNormalizationGammaAndBetaState) Autorelease() MPSCNNNormalizationGammaAndBetaState {
	rv := objc.Send[MPSCNNNormalizationGammaAndBetaState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNormalizationGammaAndBetaState creates a new MPSCNNNormalizationGammaAndBetaState instance.
func NewMPSCNNNormalizationGammaAndBetaState() MPSCNNNormalizationGammaAndBetaState {
	class := getMPSCNNNormalizationGammaAndBetaStateClass()
	rv := objc.Send[MPSCNNNormalizationGammaAndBetaState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNNormalizationGammaAndBetaStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNNormalizationGammaAndBetaState {
	instance := getMPSCNNNormalizationGammaAndBetaStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNNormalizationGammaAndBetaStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNNormalizationGammaAndBetaStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNNormalizationGammaAndBetaState {
	instance := getMPSCNNNormalizationGammaAndBetaStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNNormalizationGammaAndBetaStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNNormalizationGammaAndBetaStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNNormalizationGammaAndBetaState {
	instance := getMPSCNNNormalizationGammaAndBetaStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNNormalizationGammaAndBetaStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationGammaAndBetaState/init(gamma:beta:)
func NewCNNNormalizationGammaAndBetaStateWithGammaBeta(gamma metal.MTLBuffer, beta metal.MTLBuffer) MPSCNNNormalizationGammaAndBetaState {
	instance := getMPSCNNNormalizationGammaAndBetaStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGamma:beta:"), gamma, beta)
	return MPSCNNNormalizationGammaAndBetaStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNNormalizationGammaAndBetaStateWithResource(resource metal.MTLResource) MPSCNNNormalizationGammaAndBetaState {
	instance := getMPSCNNNormalizationGammaAndBetaStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNNormalizationGammaAndBetaStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNNormalizationGammaAndBetaStateWithResources(resources []objectivec.IObject) MPSCNNNormalizationGammaAndBetaState {
	instance := getMPSCNNNormalizationGammaAndBetaStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNNormalizationGammaAndBetaStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationGammaAndBetaState/init(gamma:beta:)
func (c MPSCNNNormalizationGammaAndBetaState) InitWithGammaBeta(gamma metal.MTLBuffer, beta metal.MTLBuffer) MPSCNNNormalizationGammaAndBetaState {
	rv := objc.Send[MPSCNNNormalizationGammaAndBetaState](c.ID, objc.Sel("initWithGamma:beta:"), gamma, beta)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationGammaAndBetaState/temporaryState(with:numberOfFeatureChannels:)
func (_MPSCNNNormalizationGammaAndBetaStateClass MPSCNNNormalizationGammaAndBetaStateClass) TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer metal.MTLCommandBuffer, numberOfFeatureChannels uint) MPSCNNNormalizationGammaAndBetaState {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNormalizationGammaAndBetaStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:numberOfFeatureChannels:"), commandBuffer, numberOfFeatureChannels)
	return MPSCNNNormalizationGammaAndBetaStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationGammaAndBetaState/beta
func (c MPSCNNNormalizationGammaAndBetaState) Beta() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("beta"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationGammaAndBetaState/gamma
func (c MPSCNNNormalizationGammaAndBetaState) Gamma() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gamma"))
	return metal.MTLBufferObjectFromID(rv)
}

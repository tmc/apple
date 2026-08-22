// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNNormalizationMeanAndVarianceState] class.
var (
	_MPSCNNNormalizationMeanAndVarianceStateClass     MPSCNNNormalizationMeanAndVarianceStateClass
	_MPSCNNNormalizationMeanAndVarianceStateClassOnce sync.Once
)

func getMPSCNNNormalizationMeanAndVarianceStateClass() MPSCNNNormalizationMeanAndVarianceStateClass {
	_MPSCNNNormalizationMeanAndVarianceStateClassOnce.Do(func() {
		_MPSCNNNormalizationMeanAndVarianceStateClass = MPSCNNNormalizationMeanAndVarianceStateClass{class: objc.GetClass("MPSCNNNormalizationMeanAndVarianceState")}
	})
	return _MPSCNNNormalizationMeanAndVarianceStateClass
}

// GetMPSCNNNormalizationMeanAndVarianceStateClass returns the class object for MPSCNNNormalizationMeanAndVarianceState.
func GetMPSCNNNormalizationMeanAndVarianceStateClass() MPSCNNNormalizationMeanAndVarianceStateClass {
	return getMPSCNNNormalizationMeanAndVarianceStateClass()
}

type MPSCNNNormalizationMeanAndVarianceStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNormalizationMeanAndVarianceStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNormalizationMeanAndVarianceStateClass) Alloc() MPSCNNNormalizationMeanAndVarianceState {
	rv := objc.Send[MPSCNNNormalizationMeanAndVarianceState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that stores mean and variance terms used to execute batch
// normalization.
//
// # Initializers
//
//   - [MPSCNNNormalizationMeanAndVarianceState.InitWithMeanVariance]
//
// # Instance Properties
//
//   - [MPSCNNNormalizationMeanAndVarianceState.Mean]
//   - [MPSCNNNormalizationMeanAndVarianceState.Variance]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationMeanAndVarianceState
type MPSCNNNormalizationMeanAndVarianceState struct {
	MPSState
}

// MPSCNNNormalizationMeanAndVarianceStateFromID constructs a [MPSCNNNormalizationMeanAndVarianceState] from an objc.ID.
//
// An object that stores mean and variance terms used to execute batch
// normalization.
func MPSCNNNormalizationMeanAndVarianceStateFromID(id objc.ID) MPSCNNNormalizationMeanAndVarianceState {
	return MPSCNNNormalizationMeanAndVarianceState{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSCNNNormalizationMeanAndVarianceState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNormalizationMeanAndVarianceState] class.
//
// # Initializers
//
//   - [IMPSCNNNormalizationMeanAndVarianceState.InitWithMeanVariance]
//
// # Instance Properties
//
//   - [IMPSCNNNormalizationMeanAndVarianceState.Mean]
//   - [IMPSCNNNormalizationMeanAndVarianceState.Variance]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationMeanAndVarianceState
type IMPSCNNNormalizationMeanAndVarianceState interface {
	IMPSState

	// Topic: Initializers

	InitWithMeanVariance(mean metal.MTLBuffer, variance metal.MTLBuffer) MPSCNNNormalizationMeanAndVarianceState

	// Topic: Instance Properties

	Mean() metal.MTLBuffer
	Variance() metal.MTLBuffer
}

// Init initializes the instance.
func (c MPSCNNNormalizationMeanAndVarianceState) Init() MPSCNNNormalizationMeanAndVarianceState {
	rv := objc.Send[MPSCNNNormalizationMeanAndVarianceState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNormalizationMeanAndVarianceState) Autorelease() MPSCNNNormalizationMeanAndVarianceState {
	rv := objc.Send[MPSCNNNormalizationMeanAndVarianceState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNormalizationMeanAndVarianceState creates a new MPSCNNNormalizationMeanAndVarianceState instance.
func NewMPSCNNNormalizationMeanAndVarianceState() MPSCNNNormalizationMeanAndVarianceState {
	class := getMPSCNNNormalizationMeanAndVarianceStateClass()
	rv := objc.Send[MPSCNNNormalizationMeanAndVarianceState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNNormalizationMeanAndVarianceStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNNormalizationMeanAndVarianceState {
	instance := getMPSCNNNormalizationMeanAndVarianceStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNNormalizationMeanAndVarianceStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNNormalizationMeanAndVarianceStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNNormalizationMeanAndVarianceState {
	instance := getMPSCNNNormalizationMeanAndVarianceStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNNormalizationMeanAndVarianceStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNNormalizationMeanAndVarianceStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNNormalizationMeanAndVarianceState {
	instance := getMPSCNNNormalizationMeanAndVarianceStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNNormalizationMeanAndVarianceStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationMeanAndVarianceState/init(mean:variance:)
func NewCNNNormalizationMeanAndVarianceStateWithMeanVariance(mean metal.MTLBuffer, variance metal.MTLBuffer) MPSCNNNormalizationMeanAndVarianceState {
	instance := getMPSCNNNormalizationMeanAndVarianceStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMean:variance:"), mean, variance)
	return MPSCNNNormalizationMeanAndVarianceStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNNormalizationMeanAndVarianceStateWithResource(resource metal.MTLResource) MPSCNNNormalizationMeanAndVarianceState {
	instance := getMPSCNNNormalizationMeanAndVarianceStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNNormalizationMeanAndVarianceStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNNormalizationMeanAndVarianceStateWithResources(resources []objectivec.IObject) MPSCNNNormalizationMeanAndVarianceState {
	instance := getMPSCNNNormalizationMeanAndVarianceStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNNormalizationMeanAndVarianceStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationMeanAndVarianceState/init(mean:variance:)
func (c MPSCNNNormalizationMeanAndVarianceState) InitWithMeanVariance(mean metal.MTLBuffer, variance metal.MTLBuffer) MPSCNNNormalizationMeanAndVarianceState {
	rv := objc.Send[MPSCNNNormalizationMeanAndVarianceState](c.ID, objc.Sel("initWithMean:variance:"), mean, variance)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationMeanAndVarianceState/temporaryState(with:numberOfFeatureChannels:)
func (_MPSCNNNormalizationMeanAndVarianceStateClass MPSCNNNormalizationMeanAndVarianceStateClass) TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer metal.MTLCommandBuffer, numberOfFeatureChannels uint) MPSCNNNormalizationMeanAndVarianceState {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNormalizationMeanAndVarianceStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:numberOfFeatureChannels:"), commandBuffer, numberOfFeatureChannels)
	return MPSCNNNormalizationMeanAndVarianceStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationMeanAndVarianceState/mean
func (c MPSCNNNormalizationMeanAndVarianceState) Mean() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("mean"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationMeanAndVarianceState/variance
func (c MPSCNNNormalizationMeanAndVarianceState) Variance() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("variance"))
	return metal.MTLBufferObjectFromID(rv)
}

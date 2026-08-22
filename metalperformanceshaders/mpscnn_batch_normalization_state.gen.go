// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNBatchNormalizationState] class.
var (
	_MPSCNNBatchNormalizationStateClass     MPSCNNBatchNormalizationStateClass
	_MPSCNNBatchNormalizationStateClassOnce sync.Once
)

func getMPSCNNBatchNormalizationStateClass() MPSCNNBatchNormalizationStateClass {
	_MPSCNNBatchNormalizationStateClassOnce.Do(func() {
		_MPSCNNBatchNormalizationStateClass = MPSCNNBatchNormalizationStateClass{class: objc.GetClass("MPSCNNBatchNormalizationState")}
	})
	return _MPSCNNBatchNormalizationStateClass
}

// GetMPSCNNBatchNormalizationStateClass returns the class object for MPSCNNBatchNormalizationState.
func GetMPSCNNBatchNormalizationStateClass() MPSCNNBatchNormalizationStateClass {
	return getMPSCNNBatchNormalizationStateClass()
}

type MPSCNNBatchNormalizationStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBatchNormalizationStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBatchNormalizationStateClass) Alloc() MPSCNNBatchNormalizationState {
	rv := objc.Send[MPSCNNBatchNormalizationState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that stores data required to execute batch normalization.
//
// # Instance Properties
//
//   - [MPSCNNBatchNormalizationState.BatchNormalization]
//
// # Instance Methods
//
//   - [MPSCNNBatchNormalizationState.Beta]
//   - [MPSCNNBatchNormalizationState.Gamma]
//   - [MPSCNNBatchNormalizationState.GradientForBeta]
//   - [MPSCNNBatchNormalizationState.GradientForGamma]
//   - [MPSCNNBatchNormalizationState.Mean]
//   - [MPSCNNBatchNormalizationState.Reset]
//   - [MPSCNNBatchNormalizationState.Variance]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState
type MPSCNNBatchNormalizationState struct {
	MPSNNGradientState
}

// MPSCNNBatchNormalizationStateFromID constructs a [MPSCNNBatchNormalizationState] from an objc.ID.
//
// An object that stores data required to execute batch normalization.
func MPSCNNBatchNormalizationStateFromID(id objc.ID) MPSCNNBatchNormalizationState {
	return MPSCNNBatchNormalizationState{MPSNNGradientState: MPSNNGradientStateFromID(id)}
}

// NOTE: MPSCNNBatchNormalizationState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBatchNormalizationState] class.
//
// # Instance Properties
//
//   - [IMPSCNNBatchNormalizationState.BatchNormalization]
//
// # Instance Methods
//
//   - [IMPSCNNBatchNormalizationState.Beta]
//   - [IMPSCNNBatchNormalizationState.Gamma]
//   - [IMPSCNNBatchNormalizationState.GradientForBeta]
//   - [IMPSCNNBatchNormalizationState.GradientForGamma]
//   - [IMPSCNNBatchNormalizationState.Mean]
//   - [IMPSCNNBatchNormalizationState.Reset]
//   - [IMPSCNNBatchNormalizationState.Variance]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState
type IMPSCNNBatchNormalizationState interface {
	IMPSNNGradientState

	// Topic: Instance Properties

	BatchNormalization() IMPSCNNBatchNormalization

	// Topic: Instance Methods

	Beta() metal.MTLBuffer
	Gamma() metal.MTLBuffer
	GradientForBeta() metal.MTLBuffer
	GradientForGamma() metal.MTLBuffer
	Mean() metal.MTLBuffer
	Reset()
	Variance() metal.MTLBuffer
}

// Init initializes the instance.
func (c MPSCNNBatchNormalizationState) Init() MPSCNNBatchNormalizationState {
	rv := objc.Send[MPSCNNBatchNormalizationState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBatchNormalizationState) Autorelease() MPSCNNBatchNormalizationState {
	rv := objc.Send[MPSCNNBatchNormalizationState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBatchNormalizationState creates a new MPSCNNBatchNormalizationState instance.
func NewMPSCNNBatchNormalizationState() MPSCNNBatchNormalizationState {
	class := getMPSCNNBatchNormalizationStateClass()
	rv := objc.Send[MPSCNNBatchNormalizationState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNBatchNormalizationStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNBatchNormalizationState {
	instance := getMPSCNNBatchNormalizationStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNBatchNormalizationStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNBatchNormalizationStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNBatchNormalizationState {
	instance := getMPSCNNBatchNormalizationStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNBatchNormalizationStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNBatchNormalizationStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNBatchNormalizationState {
	instance := getMPSCNNBatchNormalizationStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNBatchNormalizationStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNBatchNormalizationStateWithResource(resource metal.MTLResource) MPSCNNBatchNormalizationState {
	instance := getMPSCNNBatchNormalizationStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNBatchNormalizationStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNBatchNormalizationStateWithResources(resources []objectivec.IObject) MPSCNNBatchNormalizationState {
	instance := getMPSCNNBatchNormalizationStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNBatchNormalizationStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState/beta()
func (c MPSCNNBatchNormalizationState) Beta() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("beta"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState/gamma()
func (c MPSCNNBatchNormalizationState) Gamma() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gamma"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState/gradientForBeta()
func (c MPSCNNBatchNormalizationState) GradientForBeta() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gradientForBeta"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState/gradientForGamma()
func (c MPSCNNBatchNormalizationState) GradientForGamma() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gradientForGamma"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState/mean()
func (c MPSCNNBatchNormalizationState) Mean() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("mean"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState/reset()
func (c MPSCNNBatchNormalizationState) Reset() {
	objc.Send[objc.ID](c.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState/variance()
func (c MPSCNNBatchNormalizationState) Variance() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("variance"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState/batchNormalization
func (c MPSCNNBatchNormalizationState) BatchNormalization() IMPSCNNBatchNormalization {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("batchNormalization"))
	return MPSCNNBatchNormalizationFromID(objc.ID(rv))
}

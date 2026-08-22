// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNInstanceNormalizationGradientState] class.
var (
	_MPSCNNInstanceNormalizationGradientStateClass     MPSCNNInstanceNormalizationGradientStateClass
	_MPSCNNInstanceNormalizationGradientStateClassOnce sync.Once
)

func getMPSCNNInstanceNormalizationGradientStateClass() MPSCNNInstanceNormalizationGradientStateClass {
	_MPSCNNInstanceNormalizationGradientStateClassOnce.Do(func() {
		_MPSCNNInstanceNormalizationGradientStateClass = MPSCNNInstanceNormalizationGradientStateClass{class: objc.GetClass("MPSCNNInstanceNormalizationGradientState")}
	})
	return _MPSCNNInstanceNormalizationGradientStateClass
}

// GetMPSCNNInstanceNormalizationGradientStateClass returns the class object for MPSCNNInstanceNormalizationGradientState.
func GetMPSCNNInstanceNormalizationGradientStateClass() MPSCNNInstanceNormalizationGradientStateClass {
	return getMPSCNNInstanceNormalizationGradientStateClass()
}

type MPSCNNInstanceNormalizationGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNInstanceNormalizationGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNInstanceNormalizationGradientStateClass) Alloc() MPSCNNInstanceNormalizationGradientState {
	rv := objc.Send[MPSCNNInstanceNormalizationGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that stores information required to execute a gradient pass for
// instance normalization.
//
// # Instance Properties
//
//   - [MPSCNNInstanceNormalizationGradientState.Beta]
//   - [MPSCNNInstanceNormalizationGradientState.Gamma]
//   - [MPSCNNInstanceNormalizationGradientState.GradientForBeta]
//   - [MPSCNNInstanceNormalizationGradientState.GradientForGamma]
//   - [MPSCNNInstanceNormalizationGradientState.InstanceNormalization]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientState
type MPSCNNInstanceNormalizationGradientState struct {
	MPSNNGradientState
}

// MPSCNNInstanceNormalizationGradientStateFromID constructs a [MPSCNNInstanceNormalizationGradientState] from an objc.ID.
//
// An object that stores information required to execute a gradient pass for
// instance normalization.
func MPSCNNInstanceNormalizationGradientStateFromID(id objc.ID) MPSCNNInstanceNormalizationGradientState {
	return MPSCNNInstanceNormalizationGradientState{MPSNNGradientState: MPSNNGradientStateFromID(id)}
}

// NOTE: MPSCNNInstanceNormalizationGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNInstanceNormalizationGradientState] class.
//
// # Instance Properties
//
//   - [IMPSCNNInstanceNormalizationGradientState.Beta]
//   - [IMPSCNNInstanceNormalizationGradientState.Gamma]
//   - [IMPSCNNInstanceNormalizationGradientState.GradientForBeta]
//   - [IMPSCNNInstanceNormalizationGradientState.GradientForGamma]
//   - [IMPSCNNInstanceNormalizationGradientState.InstanceNormalization]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientState
type IMPSCNNInstanceNormalizationGradientState interface {
	IMPSNNGradientState

	// Topic: Instance Properties

	Beta() metal.MTLBuffer
	Gamma() metal.MTLBuffer
	GradientForBeta() metal.MTLBuffer
	GradientForGamma() metal.MTLBuffer
	InstanceNormalization() IMPSCNNInstanceNormalization
}

// Init initializes the instance.
func (c MPSCNNInstanceNormalizationGradientState) Init() MPSCNNInstanceNormalizationGradientState {
	rv := objc.Send[MPSCNNInstanceNormalizationGradientState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNInstanceNormalizationGradientState) Autorelease() MPSCNNInstanceNormalizationGradientState {
	rv := objc.Send[MPSCNNInstanceNormalizationGradientState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNInstanceNormalizationGradientState creates a new MPSCNNInstanceNormalizationGradientState instance.
func NewMPSCNNInstanceNormalizationGradientState() MPSCNNInstanceNormalizationGradientState {
	class := getMPSCNNInstanceNormalizationGradientStateClass()
	rv := objc.Send[MPSCNNInstanceNormalizationGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNInstanceNormalizationGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNInstanceNormalizationGradientState {
	instance := getMPSCNNInstanceNormalizationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNInstanceNormalizationGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNInstanceNormalizationGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNInstanceNormalizationGradientState {
	instance := getMPSCNNInstanceNormalizationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNInstanceNormalizationGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNInstanceNormalizationGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNInstanceNormalizationGradientState {
	instance := getMPSCNNInstanceNormalizationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNInstanceNormalizationGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNInstanceNormalizationGradientStateWithResource(resource metal.MTLResource) MPSCNNInstanceNormalizationGradientState {
	instance := getMPSCNNInstanceNormalizationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNInstanceNormalizationGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNInstanceNormalizationGradientStateWithResources(resources []objectivec.IObject) MPSCNNInstanceNormalizationGradientState {
	instance := getMPSCNNInstanceNormalizationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNInstanceNormalizationGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientState/beta
func (c MPSCNNInstanceNormalizationGradientState) Beta() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("beta"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientState/gamma
func (c MPSCNNInstanceNormalizationGradientState) Gamma() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gamma"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientState/gradientForBeta
func (c MPSCNNInstanceNormalizationGradientState) GradientForBeta() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gradientForBeta"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientState/gradientForGamma
func (c MPSCNNInstanceNormalizationGradientState) GradientForGamma() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gradientForGamma"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientState/instanceNormalization
func (c MPSCNNInstanceNormalizationGradientState) InstanceNormalization() IMPSCNNInstanceNormalization {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("instanceNormalization"))
	return MPSCNNInstanceNormalizationFromID(objc.ID(rv))
}

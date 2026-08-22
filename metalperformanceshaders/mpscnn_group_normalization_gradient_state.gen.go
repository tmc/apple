// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNGroupNormalizationGradientState] class.
var (
	_MPSCNNGroupNormalizationGradientStateClass     MPSCNNGroupNormalizationGradientStateClass
	_MPSCNNGroupNormalizationGradientStateClassOnce sync.Once
)

func getMPSCNNGroupNormalizationGradientStateClass() MPSCNNGroupNormalizationGradientStateClass {
	_MPSCNNGroupNormalizationGradientStateClassOnce.Do(func() {
		_MPSCNNGroupNormalizationGradientStateClass = MPSCNNGroupNormalizationGradientStateClass{class: objc.GetClass("MPSCNNGroupNormalizationGradientState")}
	})
	return _MPSCNNGroupNormalizationGradientStateClass
}

// GetMPSCNNGroupNormalizationGradientStateClass returns the class object for MPSCNNGroupNormalizationGradientState.
func GetMPSCNNGroupNormalizationGradientStateClass() MPSCNNGroupNormalizationGradientStateClass {
	return getMPSCNNGroupNormalizationGradientStateClass()
}

type MPSCNNGroupNormalizationGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNGroupNormalizationGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNGroupNormalizationGradientStateClass) Alloc() MPSCNNGroupNormalizationGradientState {
	rv := objc.Send[MPSCNNGroupNormalizationGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSCNNGroupNormalizationGradientState.Beta]
//   - [MPSCNNGroupNormalizationGradientState.Gamma]
//   - [MPSCNNGroupNormalizationGradientState.GradientForBeta]
//   - [MPSCNNGroupNormalizationGradientState.GradientForGamma]
//   - [MPSCNNGroupNormalizationGradientState.GroupNormalization]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState
type MPSCNNGroupNormalizationGradientState struct {
	MPSNNGradientState
}

// MPSCNNGroupNormalizationGradientStateFromID constructs a [MPSCNNGroupNormalizationGradientState] from an objc.ID.
func MPSCNNGroupNormalizationGradientStateFromID(id objc.ID) MPSCNNGroupNormalizationGradientState {
	return MPSCNNGroupNormalizationGradientState{MPSNNGradientState: MPSNNGradientStateFromID(id)}
}

// NOTE: MPSCNNGroupNormalizationGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNGroupNormalizationGradientState] class.
//
// # Instance Properties
//
//   - [IMPSCNNGroupNormalizationGradientState.Beta]
//   - [IMPSCNNGroupNormalizationGradientState.Gamma]
//   - [IMPSCNNGroupNormalizationGradientState.GradientForBeta]
//   - [IMPSCNNGroupNormalizationGradientState.GradientForGamma]
//   - [IMPSCNNGroupNormalizationGradientState.GroupNormalization]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState
type IMPSCNNGroupNormalizationGradientState interface {
	IMPSNNGradientState

	// Topic: Instance Properties

	Beta() metal.MTLBuffer
	Gamma() metal.MTLBuffer
	GradientForBeta() metal.MTLBuffer
	GradientForGamma() metal.MTLBuffer
	GroupNormalization() IMPSCNNGroupNormalization
}

// Init initializes the instance.
func (c MPSCNNGroupNormalizationGradientState) Init() MPSCNNGroupNormalizationGradientState {
	rv := objc.Send[MPSCNNGroupNormalizationGradientState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNGroupNormalizationGradientState) Autorelease() MPSCNNGroupNormalizationGradientState {
	rv := objc.Send[MPSCNNGroupNormalizationGradientState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNGroupNormalizationGradientState creates a new MPSCNNGroupNormalizationGradientState instance.
func NewMPSCNNGroupNormalizationGradientState() MPSCNNGroupNormalizationGradientState {
	class := getMPSCNNGroupNormalizationGradientStateClass()
	rv := objc.Send[MPSCNNGroupNormalizationGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNGroupNormalizationGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNGroupNormalizationGradientState {
	instance := getMPSCNNGroupNormalizationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNGroupNormalizationGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNGroupNormalizationGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNGroupNormalizationGradientState {
	instance := getMPSCNNGroupNormalizationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNGroupNormalizationGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNGroupNormalizationGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNGroupNormalizationGradientState {
	instance := getMPSCNNGroupNormalizationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNGroupNormalizationGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNGroupNormalizationGradientStateWithResource(resource metal.MTLResource) MPSCNNGroupNormalizationGradientState {
	instance := getMPSCNNGroupNormalizationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNGroupNormalizationGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNGroupNormalizationGradientStateWithResources(resources []objectivec.IObject) MPSCNNGroupNormalizationGradientState {
	instance := getMPSCNNGroupNormalizationGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNGroupNormalizationGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState/beta
func (c MPSCNNGroupNormalizationGradientState) Beta() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("beta"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState/gamma
func (c MPSCNNGroupNormalizationGradientState) Gamma() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gamma"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState/gradientForBeta
func (c MPSCNNGroupNormalizationGradientState) GradientForBeta() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gradientForBeta"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState/gradientForGamma
func (c MPSCNNGroupNormalizationGradientState) GradientForGamma() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gradientForGamma"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState/groupNormalization
func (c MPSCNNGroupNormalizationGradientState) GroupNormalization() IMPSCNNGroupNormalization {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("groupNormalization"))
	return MPSCNNGroupNormalizationFromID(objc.ID(rv))
}

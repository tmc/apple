// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGradientState] class.
var (
	_MPSNNGradientStateClass     MPSNNGradientStateClass
	_MPSNNGradientStateClassOnce sync.Once
)

func getMPSNNGradientStateClass() MPSNNGradientStateClass {
	_MPSNNGradientStateClassOnce.Do(func() {
		_MPSNNGradientStateClass = MPSNNGradientStateClass{class: objc.GetClass("MPSNNGradientState")}
	})
	return _MPSNNGradientStateClass
}

// GetMPSNNGradientStateClass returns the class object for MPSNNGradientState.
func GetMPSNNGradientStateClass() MPSNNGradientStateClass {
	return getMPSNNGradientStateClass()
}

type MPSNNGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGradientStateClass) Alloc() MPSNNGradientState {
	rv := objc.Send[MPSNNGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class representing the state of a gradient kernel when it was encoded.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientState
type MPSNNGradientState struct {
	MPSState
}

// MPSNNGradientStateFromID constructs a [MPSNNGradientState] from an objc.ID.
//
// A class representing the state of a gradient kernel when it was encoded.
func MPSNNGradientStateFromID(id objc.ID) MPSNNGradientState {
	return MPSNNGradientState{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSNNGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNGradientState] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientState
type IMPSNNGradientState interface {
	IMPSState
}

// Init initializes the instance.
func (g MPSNNGradientState) Init() MPSNNGradientState {
	rv := objc.Send[MPSNNGradientState](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGradientState) Autorelease() MPSNNGradientState {
	rv := objc.Send[MPSNNGradientState](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGradientState creates a new MPSNNGradientState instance.
func NewMPSNNGradientState() MPSNNGradientState {
	class := getMPSNNGradientStateClass()
	rv := objc.Send[MPSNNGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSNNGradientState {
	instance := getMPSNNGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSNNGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSNNGradientState {
	instance := getMPSNNGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSNNGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSNNGradientState {
	instance := getMPSNNGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSNNGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewGradientStateWithResource(resource metal.MTLResource) MPSNNGradientState {
	instance := getMPSNNGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSNNGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewGradientStateWithResources(resources []objectivec.IObject) MPSNNGradientState {
	instance := getMPSNNGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSNNGradientStateFromID(rv)
}

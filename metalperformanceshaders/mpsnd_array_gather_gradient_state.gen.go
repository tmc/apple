// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNDArrayGatherGradientState] class.
var (
	_MPSNDArrayGatherGradientStateClass     MPSNDArrayGatherGradientStateClass
	_MPSNDArrayGatherGradientStateClassOnce sync.Once
)

func getMPSNDArrayGatherGradientStateClass() MPSNDArrayGatherGradientStateClass {
	_MPSNDArrayGatherGradientStateClassOnce.Do(func() {
		_MPSNDArrayGatherGradientStateClass = MPSNDArrayGatherGradientStateClass{class: objc.GetClass("MPSNDArrayGatherGradientState")}
	})
	return _MPSNDArrayGatherGradientStateClass
}

// GetMPSNDArrayGatherGradientStateClass returns the class object for MPSNDArrayGatherGradientState.
func GetMPSNDArrayGatherGradientStateClass() MPSNDArrayGatherGradientStateClass {
	return getMPSNDArrayGatherGradientStateClass()
}

type MPSNDArrayGatherGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayGatherGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayGatherGradientStateClass) Alloc() MPSNDArrayGatherGradientState {
	rv := objc.Send[MPSNDArrayGatherGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGatherGradientState
type MPSNDArrayGatherGradientState struct {
	MPSNDArrayGradientState
}

// MPSNDArrayGatherGradientStateFromID constructs a [MPSNDArrayGatherGradientState] from an objc.ID.
func MPSNDArrayGatherGradientStateFromID(id objc.ID) MPSNDArrayGatherGradientState {
	return MPSNDArrayGatherGradientState{MPSNDArrayGradientState: MPSNDArrayGradientStateFromID(id)}
}

// NOTE: MPSNDArrayGatherGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayGatherGradientState] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGatherGradientState
type IMPSNDArrayGatherGradientState interface {
	IMPSNDArrayGradientState
}

// Init initializes the instance.
func (n MPSNDArrayGatherGradientState) Init() MPSNDArrayGatherGradientState {
	rv := objc.Send[MPSNDArrayGatherGradientState](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayGatherGradientState) Autorelease() MPSNDArrayGatherGradientState {
	rv := objc.Send[MPSNDArrayGatherGradientState](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayGatherGradientState creates a new MPSNDArrayGatherGradientState instance.
func NewMPSNDArrayGatherGradientState() MPSNDArrayGatherGradientState {
	class := getMPSNDArrayGatherGradientStateClass()
	rv := objc.Send[MPSNDArrayGatherGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewNDArrayGatherGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSNDArrayGatherGradientState {
	instance := getMPSNDArrayGatherGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSNDArrayGatherGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewNDArrayGatherGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSNDArrayGatherGradientState {
	instance := getMPSNDArrayGatherGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSNDArrayGatherGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewNDArrayGatherGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSNDArrayGatherGradientState {
	instance := getMPSNDArrayGatherGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSNDArrayGatherGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewNDArrayGatherGradientStateWithResource(resource metal.MTLResource) MPSNDArrayGatherGradientState {
	instance := getMPSNDArrayGatherGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSNDArrayGatherGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewNDArrayGatherGradientStateWithResources(resources []objectivec.IObject) MPSNDArrayGatherGradientState {
	instance := getMPSNDArrayGatherGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSNDArrayGatherGradientStateFromID(rv)
}

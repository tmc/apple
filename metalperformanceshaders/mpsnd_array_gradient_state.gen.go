// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNDArrayGradientState] class.
var (
	_MPSNDArrayGradientStateClass     MPSNDArrayGradientStateClass
	_MPSNDArrayGradientStateClassOnce sync.Once
)

func getMPSNDArrayGradientStateClass() MPSNDArrayGradientStateClass {
	_MPSNDArrayGradientStateClassOnce.Do(func() {
		_MPSNDArrayGradientStateClass = MPSNDArrayGradientStateClass{class: objc.GetClass("MPSNDArrayGradientState")}
	})
	return _MPSNDArrayGradientStateClass
}

// GetMPSNDArrayGradientStateClass returns the class object for MPSNDArrayGradientState.
func GetMPSNDArrayGradientStateClass() MPSNDArrayGradientStateClass {
	return getMPSNDArrayGradientStateClass()
}

type MPSNDArrayGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayGradientStateClass) Alloc() MPSNDArrayGradientState {
	rv := objc.Send[MPSNDArrayGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGradientState
type MPSNDArrayGradientState struct {
	MPSState
}

// MPSNDArrayGradientStateFromID constructs a [MPSNDArrayGradientState] from an objc.ID.
func MPSNDArrayGradientStateFromID(id objc.ID) MPSNDArrayGradientState {
	return MPSNDArrayGradientState{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSNDArrayGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayGradientState] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGradientState
type IMPSNDArrayGradientState interface {
	IMPSState
}

// Init initializes the instance.
func (n MPSNDArrayGradientState) Init() MPSNDArrayGradientState {
	rv := objc.Send[MPSNDArrayGradientState](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayGradientState) Autorelease() MPSNDArrayGradientState {
	rv := objc.Send[MPSNDArrayGradientState](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayGradientState creates a new MPSNDArrayGradientState instance.
func NewMPSNDArrayGradientState() MPSNDArrayGradientState {
	class := getMPSNDArrayGradientStateClass()
	rv := objc.Send[MPSNDArrayGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewNDArrayGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSNDArrayGradientState {
	instance := getMPSNDArrayGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSNDArrayGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewNDArrayGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSNDArrayGradientState {
	instance := getMPSNDArrayGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSNDArrayGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewNDArrayGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSNDArrayGradientState {
	instance := getMPSNDArrayGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSNDArrayGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewNDArrayGradientStateWithResource(resource metal.MTLResource) MPSNDArrayGradientState {
	instance := getMPSNDArrayGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSNDArrayGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewNDArrayGradientStateWithResources(resources []objectivec.IObject) MPSNDArrayGradientState {
	instance := getMPSNDArrayGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSNDArrayGradientStateFromID(rv)
}

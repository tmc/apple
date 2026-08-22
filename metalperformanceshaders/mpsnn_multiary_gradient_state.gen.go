// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNMultiaryGradientState] class.
var (
	_MPSNNMultiaryGradientStateClass     MPSNNMultiaryGradientStateClass
	_MPSNNMultiaryGradientStateClassOnce sync.Once
)

func getMPSNNMultiaryGradientStateClass() MPSNNMultiaryGradientStateClass {
	_MPSNNMultiaryGradientStateClassOnce.Do(func() {
		_MPSNNMultiaryGradientStateClass = MPSNNMultiaryGradientStateClass{class: objc.GetClass("MPSNNMultiaryGradientState")}
	})
	return _MPSNNMultiaryGradientStateClass
}

// GetMPSNNMultiaryGradientStateClass returns the class object for MPSNNMultiaryGradientState.
func GetMPSNNMultiaryGradientStateClass() MPSNNMultiaryGradientStateClass {
	return getMPSNNMultiaryGradientStateClass()
}

type MPSNNMultiaryGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNMultiaryGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNMultiaryGradientStateClass) Alloc() MPSNNMultiaryGradientState {
	rv := objc.Send[MPSNNMultiaryGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiaryGradientState
type MPSNNMultiaryGradientState struct {
	MPSState
}

// MPSNNMultiaryGradientStateFromID constructs a [MPSNNMultiaryGradientState] from an objc.ID.
func MPSNNMultiaryGradientStateFromID(id objc.ID) MPSNNMultiaryGradientState {
	return MPSNNMultiaryGradientState{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSNNMultiaryGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNMultiaryGradientState] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiaryGradientState
type IMPSNNMultiaryGradientState interface {
	IMPSState
}

// Init initializes the instance.
func (m MPSNNMultiaryGradientState) Init() MPSNNMultiaryGradientState {
	rv := objc.Send[MPSNNMultiaryGradientState](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSNNMultiaryGradientState) Autorelease() MPSNNMultiaryGradientState {
	rv := objc.Send[MPSNNMultiaryGradientState](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNMultiaryGradientState creates a new MPSNNMultiaryGradientState instance.
func NewMPSNNMultiaryGradientState() MPSNNMultiaryGradientState {
	class := getMPSNNMultiaryGradientStateClass()
	rv := objc.Send[MPSNNMultiaryGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewMultiaryGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSNNMultiaryGradientState {
	instance := getMPSNNMultiaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSNNMultiaryGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewMultiaryGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSNNMultiaryGradientState {
	instance := getMPSNNMultiaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSNNMultiaryGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewMultiaryGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSNNMultiaryGradientState {
	instance := getMPSNNMultiaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSNNMultiaryGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewMultiaryGradientStateWithResource(resource metal.MTLResource) MPSNNMultiaryGradientState {
	instance := getMPSNNMultiaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSNNMultiaryGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewMultiaryGradientStateWithResources(resources []objectivec.IObject) MPSNNMultiaryGradientState {
	instance := getMPSNNMultiaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSNNMultiaryGradientStateFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNBinaryGradientState] class.
var (
	_MPSNNBinaryGradientStateClass     MPSNNBinaryGradientStateClass
	_MPSNNBinaryGradientStateClassOnce sync.Once
)

func getMPSNNBinaryGradientStateClass() MPSNNBinaryGradientStateClass {
	_MPSNNBinaryGradientStateClassOnce.Do(func() {
		_MPSNNBinaryGradientStateClass = MPSNNBinaryGradientStateClass{class: objc.GetClass("MPSNNBinaryGradientState")}
	})
	return _MPSNNBinaryGradientStateClass
}

// GetMPSNNBinaryGradientStateClass returns the class object for MPSNNBinaryGradientState.
func GetMPSNNBinaryGradientStateClass() MPSNNBinaryGradientStateClass {
	return getMPSNNBinaryGradientStateClass()
}

type MPSNNBinaryGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNBinaryGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNBinaryGradientStateClass) Alloc() MPSNNBinaryGradientState {
	rv := objc.Send[MPSNNBinaryGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class representing the state of a gradient binary kernel when it was
// encoded.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryGradientState
type MPSNNBinaryGradientState struct {
	MPSState
}

// MPSNNBinaryGradientStateFromID constructs a [MPSNNBinaryGradientState] from an objc.ID.
//
// A class representing the state of a gradient binary kernel when it was
// encoded.
func MPSNNBinaryGradientStateFromID(id objc.ID) MPSNNBinaryGradientState {
	return MPSNNBinaryGradientState{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSNNBinaryGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNBinaryGradientState] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryGradientState
type IMPSNNBinaryGradientState interface {
	IMPSState
}

// Init initializes the instance.
func (b MPSNNBinaryGradientState) Init() MPSNNBinaryGradientState {
	rv := objc.Send[MPSNNBinaryGradientState](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MPSNNBinaryGradientState) Autorelease() MPSNNBinaryGradientState {
	rv := objc.Send[MPSNNBinaryGradientState](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNBinaryGradientState creates a new MPSNNBinaryGradientState instance.
func NewMPSNNBinaryGradientState() MPSNNBinaryGradientState {
	class := getMPSNNBinaryGradientStateClass()
	rv := objc.Send[MPSNNBinaryGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewBinaryGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSNNBinaryGradientState {
	instance := getMPSNNBinaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSNNBinaryGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewBinaryGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSNNBinaryGradientState {
	instance := getMPSNNBinaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSNNBinaryGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewBinaryGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSNNBinaryGradientState {
	instance := getMPSNNBinaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSNNBinaryGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewBinaryGradientStateWithResource(resource metal.MTLResource) MPSNNBinaryGradientState {
	instance := getMPSNNBinaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSNNBinaryGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewBinaryGradientStateWithResources(resources []objectivec.IObject) MPSNNBinaryGradientState {
	instance := getMPSNNBinaryGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSNNBinaryGradientStateFromID(rv)
}

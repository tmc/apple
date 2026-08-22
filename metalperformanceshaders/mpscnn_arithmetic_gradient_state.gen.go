// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNArithmeticGradientState] class.
var (
	_MPSCNNArithmeticGradientStateClass     MPSCNNArithmeticGradientStateClass
	_MPSCNNArithmeticGradientStateClassOnce sync.Once
)

func getMPSCNNArithmeticGradientStateClass() MPSCNNArithmeticGradientStateClass {
	_MPSCNNArithmeticGradientStateClassOnce.Do(func() {
		_MPSCNNArithmeticGradientStateClass = MPSCNNArithmeticGradientStateClass{class: objc.GetClass("MPSCNNArithmeticGradientState")}
	})
	return _MPSCNNArithmeticGradientStateClass
}

// GetMPSCNNArithmeticGradientStateClass returns the class object for MPSCNNArithmeticGradientState.
func GetMPSCNNArithmeticGradientStateClass() MPSCNNArithmeticGradientStateClass {
	return getMPSCNNArithmeticGradientStateClass()
}

type MPSCNNArithmeticGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNArithmeticGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNArithmeticGradientStateClass) Alloc() MPSCNNArithmeticGradientState {
	rv := objc.Send[MPSCNNArithmeticGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that stores the clamp mask used by gradient arithmetic operators.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradientState
type MPSCNNArithmeticGradientState struct {
	MPSNNBinaryGradientState
}

// MPSCNNArithmeticGradientStateFromID constructs a [MPSCNNArithmeticGradientState] from an objc.ID.
//
// An object that stores the clamp mask used by gradient arithmetic operators.
func MPSCNNArithmeticGradientStateFromID(id objc.ID) MPSCNNArithmeticGradientState {
	return MPSCNNArithmeticGradientState{MPSNNBinaryGradientState: MPSNNBinaryGradientStateFromID(id)}
}

// NOTE: MPSCNNArithmeticGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNArithmeticGradientState] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradientState
type IMPSCNNArithmeticGradientState interface {
	IMPSNNBinaryGradientState
}

// Init initializes the instance.
func (c MPSCNNArithmeticGradientState) Init() MPSCNNArithmeticGradientState {
	rv := objc.Send[MPSCNNArithmeticGradientState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNArithmeticGradientState) Autorelease() MPSCNNArithmeticGradientState {
	rv := objc.Send[MPSCNNArithmeticGradientState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNArithmeticGradientState creates a new MPSCNNArithmeticGradientState instance.
func NewMPSCNNArithmeticGradientState() MPSCNNArithmeticGradientState {
	class := getMPSCNNArithmeticGradientStateClass()
	rv := objc.Send[MPSCNNArithmeticGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNArithmeticGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNArithmeticGradientState {
	instance := getMPSCNNArithmeticGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNArithmeticGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNArithmeticGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNArithmeticGradientState {
	instance := getMPSCNNArithmeticGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNArithmeticGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNArithmeticGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNArithmeticGradientState {
	instance := getMPSCNNArithmeticGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNArithmeticGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNArithmeticGradientStateWithResource(resource metal.MTLResource) MPSCNNArithmeticGradientState {
	instance := getMPSCNNArithmeticGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNArithmeticGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNArithmeticGradientStateWithResources(resources []objectivec.IObject) MPSCNNArithmeticGradientState {
	instance := getMPSCNNArithmeticGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNArithmeticGradientStateFromID(rv)
}

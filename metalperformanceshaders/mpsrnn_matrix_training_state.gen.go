// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSRNNMatrixTrainingState] class.
var (
	_MPSRNNMatrixTrainingStateClass     MPSRNNMatrixTrainingStateClass
	_MPSRNNMatrixTrainingStateClassOnce sync.Once
)

func getMPSRNNMatrixTrainingStateClass() MPSRNNMatrixTrainingStateClass {
	_MPSRNNMatrixTrainingStateClassOnce.Do(func() {
		_MPSRNNMatrixTrainingStateClass = MPSRNNMatrixTrainingStateClass{class: objc.GetClass("MPSRNNMatrixTrainingState")}
	})
	return _MPSRNNMatrixTrainingStateClass
}

// GetMPSRNNMatrixTrainingStateClass returns the class object for MPSRNNMatrixTrainingState.
func GetMPSRNNMatrixTrainingStateClass() MPSRNNMatrixTrainingStateClass {
	return getMPSRNNMatrixTrainingStateClass()
}

type MPSRNNMatrixTrainingStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSRNNMatrixTrainingStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSRNNMatrixTrainingStateClass) Alloc() MPSRNNMatrixTrainingState {
	rv := objc.Send[MPSRNNMatrixTrainingState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that holds data from a forward pass to be used in a backward pass.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingState
type MPSRNNMatrixTrainingState struct {
	MPSState
}

// MPSRNNMatrixTrainingStateFromID constructs a [MPSRNNMatrixTrainingState] from an objc.ID.
//
// A class that holds data from a forward pass to be used in a backward pass.
func MPSRNNMatrixTrainingStateFromID(id objc.ID) MPSRNNMatrixTrainingState {
	return MPSRNNMatrixTrainingState{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSRNNMatrixTrainingState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSRNNMatrixTrainingState] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingState
type IMPSRNNMatrixTrainingState interface {
	IMPSState
}

// Init initializes the instance.
func (r MPSRNNMatrixTrainingState) Init() MPSRNNMatrixTrainingState {
	rv := objc.Send[MPSRNNMatrixTrainingState](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSRNNMatrixTrainingState) Autorelease() MPSRNNMatrixTrainingState {
	rv := objc.Send[MPSRNNMatrixTrainingState](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSRNNMatrixTrainingState creates a new MPSRNNMatrixTrainingState instance.
func NewMPSRNNMatrixTrainingState() MPSRNNMatrixTrainingState {
	class := getMPSRNNMatrixTrainingStateClass()
	rv := objc.Send[MPSRNNMatrixTrainingState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewRNNMatrixTrainingStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSRNNMatrixTrainingState {
	instance := getMPSRNNMatrixTrainingStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSRNNMatrixTrainingStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewRNNMatrixTrainingStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSRNNMatrixTrainingState {
	instance := getMPSRNNMatrixTrainingStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSRNNMatrixTrainingStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewRNNMatrixTrainingStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSRNNMatrixTrainingState {
	instance := getMPSRNNMatrixTrainingStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSRNNMatrixTrainingStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewRNNMatrixTrainingStateWithResource(resource metal.MTLResource) MPSRNNMatrixTrainingState {
	instance := getMPSRNNMatrixTrainingStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSRNNMatrixTrainingStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewRNNMatrixTrainingStateWithResources(resources []objectivec.IObject) MPSRNNMatrixTrainingState {
	instance := getMPSRNNMatrixTrainingStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSRNNMatrixTrainingStateFromID(rv)
}

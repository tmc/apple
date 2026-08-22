// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNDropoutGradientState] class.
var (
	_MPSCNNDropoutGradientStateClass     MPSCNNDropoutGradientStateClass
	_MPSCNNDropoutGradientStateClassOnce sync.Once
)

func getMPSCNNDropoutGradientStateClass() MPSCNNDropoutGradientStateClass {
	_MPSCNNDropoutGradientStateClassOnce.Do(func() {
		_MPSCNNDropoutGradientStateClass = MPSCNNDropoutGradientStateClass{class: objc.GetClass("MPSCNNDropoutGradientState")}
	})
	return _MPSCNNDropoutGradientStateClass
}

// GetMPSCNNDropoutGradientStateClass returns the class object for MPSCNNDropoutGradientState.
func GetMPSCNNDropoutGradientStateClass() MPSCNNDropoutGradientStateClass {
	return getMPSCNNDropoutGradientStateClass()
}

type MPSCNNDropoutGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDropoutGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDropoutGradientStateClass) Alloc() MPSCNNDropoutGradientState {
	rv := objc.Send[MPSCNNDropoutGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that stores the mask used by dropout and gradient dropout filters.
//
// # Instance Methods
//
//   - [MPSCNNDropoutGradientState.MaskData]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientState
type MPSCNNDropoutGradientState struct {
	MPSNNGradientState
}

// MPSCNNDropoutGradientStateFromID constructs a [MPSCNNDropoutGradientState] from an objc.ID.
//
// A class that stores the mask used by dropout and gradient dropout filters.
func MPSCNNDropoutGradientStateFromID(id objc.ID) MPSCNNDropoutGradientState {
	return MPSCNNDropoutGradientState{MPSNNGradientState: MPSNNGradientStateFromID(id)}
}

// NOTE: MPSCNNDropoutGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDropoutGradientState] class.
//
// # Instance Methods
//
//   - [IMPSCNNDropoutGradientState.MaskData]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientState
type IMPSCNNDropoutGradientState interface {
	IMPSNNGradientState

	// Topic: Instance Methods

	MaskData() foundation.NSData
}

// Init initializes the instance.
func (c MPSCNNDropoutGradientState) Init() MPSCNNDropoutGradientState {
	rv := objc.Send[MPSCNNDropoutGradientState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDropoutGradientState) Autorelease() MPSCNNDropoutGradientState {
	rv := objc.Send[MPSCNNDropoutGradientState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDropoutGradientState creates a new MPSCNNDropoutGradientState instance.
func NewMPSCNNDropoutGradientState() MPSCNNDropoutGradientState {
	class := getMPSCNNDropoutGradientStateClass()
	rv := objc.Send[MPSCNNDropoutGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNDropoutGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNDropoutGradientState {
	instance := getMPSCNNDropoutGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNDropoutGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNDropoutGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNDropoutGradientState {
	instance := getMPSCNNDropoutGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNDropoutGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNDropoutGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNDropoutGradientState {
	instance := getMPSCNNDropoutGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNDropoutGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNDropoutGradientStateWithResource(resource metal.MTLResource) MPSCNNDropoutGradientState {
	instance := getMPSCNNDropoutGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNDropoutGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNDropoutGradientStateWithResources(resources []objectivec.IObject) MPSCNNDropoutGradientState {
	instance := getMPSCNNDropoutGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNDropoutGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientState/maskData()
func (c MPSCNNDropoutGradientState) MaskData() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("maskData"))
	return foundation.NSDataFromID(rv)
}

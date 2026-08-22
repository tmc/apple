// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNConvolutionTransposeGradientState] class.
var (
	_MPSCNNConvolutionTransposeGradientStateClass     MPSCNNConvolutionTransposeGradientStateClass
	_MPSCNNConvolutionTransposeGradientStateClassOnce sync.Once
)

func getMPSCNNConvolutionTransposeGradientStateClass() MPSCNNConvolutionTransposeGradientStateClass {
	_MPSCNNConvolutionTransposeGradientStateClassOnce.Do(func() {
		_MPSCNNConvolutionTransposeGradientStateClass = MPSCNNConvolutionTransposeGradientStateClass{class: objc.GetClass("MPSCNNConvolutionTransposeGradientState")}
	})
	return _MPSCNNConvolutionTransposeGradientStateClass
}

// GetMPSCNNConvolutionTransposeGradientStateClass returns the class object for MPSCNNConvolutionTransposeGradientState.
func GetMPSCNNConvolutionTransposeGradientStateClass() MPSCNNConvolutionTransposeGradientStateClass {
	return getMPSCNNConvolutionTransposeGradientStateClass()
}

type MPSCNNConvolutionTransposeGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionTransposeGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionTransposeGradientStateClass) Alloc() MPSCNNConvolutionTransposeGradientState {
	rv := objc.Send[MPSCNNConvolutionTransposeGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSCNNConvolutionTransposeGradientState.ConvolutionTranspose]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientState
type MPSCNNConvolutionTransposeGradientState struct {
	MPSCNNConvolutionGradientState
}

// MPSCNNConvolutionTransposeGradientStateFromID constructs a [MPSCNNConvolutionTransposeGradientState] from an objc.ID.
func MPSCNNConvolutionTransposeGradientStateFromID(id objc.ID) MPSCNNConvolutionTransposeGradientState {
	return MPSCNNConvolutionTransposeGradientState{MPSCNNConvolutionGradientState: MPSCNNConvolutionGradientStateFromID(id)}
}

// NOTE: MPSCNNConvolutionTransposeGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionTransposeGradientState] class.
//
// # Instance Properties
//
//   - [IMPSCNNConvolutionTransposeGradientState.ConvolutionTranspose]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientState
type IMPSCNNConvolutionTransposeGradientState interface {
	IMPSCNNConvolutionGradientState

	// Topic: Instance Properties

	ConvolutionTranspose() IMPSCNNConvolutionTranspose
}

// Init initializes the instance.
func (c MPSCNNConvolutionTransposeGradientState) Init() MPSCNNConvolutionTransposeGradientState {
	rv := objc.Send[MPSCNNConvolutionTransposeGradientState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionTransposeGradientState) Autorelease() MPSCNNConvolutionTransposeGradientState {
	rv := objc.Send[MPSCNNConvolutionTransposeGradientState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionTransposeGradientState creates a new MPSCNNConvolutionTransposeGradientState instance.
func NewMPSCNNConvolutionTransposeGradientState() MPSCNNConvolutionTransposeGradientState {
	class := getMPSCNNConvolutionTransposeGradientStateClass()
	rv := objc.Send[MPSCNNConvolutionTransposeGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNConvolutionTransposeGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNConvolutionTransposeGradientState {
	instance := getMPSCNNConvolutionTransposeGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNConvolutionTransposeGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNConvolutionTransposeGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNConvolutionTransposeGradientState {
	instance := getMPSCNNConvolutionTransposeGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNConvolutionTransposeGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNConvolutionTransposeGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNConvolutionTransposeGradientState {
	instance := getMPSCNNConvolutionTransposeGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNConvolutionTransposeGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNConvolutionTransposeGradientStateWithResource(resource metal.MTLResource) MPSCNNConvolutionTransposeGradientState {
	instance := getMPSCNNConvolutionTransposeGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNConvolutionTransposeGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNConvolutionTransposeGradientStateWithResources(resources []objectivec.IObject) MPSCNNConvolutionTransposeGradientState {
	instance := getMPSCNNConvolutionTransposeGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNConvolutionTransposeGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientState/convolutionTranspose
func (c MPSCNNConvolutionTransposeGradientState) ConvolutionTranspose() IMPSCNNConvolutionTranspose {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("convolutionTranspose"))
	return MPSCNNConvolutionTransposeFromID(objc.ID(rv))
}

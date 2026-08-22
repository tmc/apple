// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNConvolutionGradientState] class.
var (
	_MPSCNNConvolutionGradientStateClass     MPSCNNConvolutionGradientStateClass
	_MPSCNNConvolutionGradientStateClassOnce sync.Once
)

func getMPSCNNConvolutionGradientStateClass() MPSCNNConvolutionGradientStateClass {
	_MPSCNNConvolutionGradientStateClassOnce.Do(func() {
		_MPSCNNConvolutionGradientStateClass = MPSCNNConvolutionGradientStateClass{class: objc.GetClass("MPSCNNConvolutionGradientState")}
	})
	return _MPSCNNConvolutionGradientStateClass
}

// GetMPSCNNConvolutionGradientStateClass returns the class object for MPSCNNConvolutionGradientState.
func GetMPSCNNConvolutionGradientStateClass() MPSCNNConvolutionGradientStateClass {
	return getMPSCNNConvolutionGradientStateClass()
}

type MPSCNNConvolutionGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionGradientStateClass) Alloc() MPSCNNConvolutionGradientState {
	rv := objc.Send[MPSCNNConvolutionGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that exposes a gradient convolution kernel’s gradient with
// respect to weights and biases.
//
// # Instance Properties
//
//   - [MPSCNNConvolutionGradientState.Convolution]
//   - [MPSCNNConvolutionGradientState.GradientForBiases]
//   - [MPSCNNConvolutionGradientState.GradientForWeights]
//   - [MPSCNNConvolutionGradientState.GradientForWeightsLayout]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientState
type MPSCNNConvolutionGradientState struct {
	MPSNNGradientState
}

// MPSCNNConvolutionGradientStateFromID constructs a [MPSCNNConvolutionGradientState] from an objc.ID.
//
// An object that exposes a gradient convolution kernel’s gradient with
// respect to weights and biases.
func MPSCNNConvolutionGradientStateFromID(id objc.ID) MPSCNNConvolutionGradientState {
	return MPSCNNConvolutionGradientState{MPSNNGradientState: MPSNNGradientStateFromID(id)}
}

// NOTE: MPSCNNConvolutionGradientState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionGradientState] class.
//
// # Instance Properties
//
//   - [IMPSCNNConvolutionGradientState.Convolution]
//   - [IMPSCNNConvolutionGradientState.GradientForBiases]
//   - [IMPSCNNConvolutionGradientState.GradientForWeights]
//   - [IMPSCNNConvolutionGradientState.GradientForWeightsLayout]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientState
type IMPSCNNConvolutionGradientState interface {
	IMPSNNGradientState

	// Topic: Instance Properties

	Convolution() IMPSCNNConvolution
	GradientForBiases() metal.MTLBuffer
	GradientForWeights() metal.MTLBuffer
	GradientForWeightsLayout() MPSCNNConvolutionWeightsLayout

	SourceHeight() uint
	SourceWidth() uint
}

// Init initializes the instance.
func (c MPSCNNConvolutionGradientState) Init() MPSCNNConvolutionGradientState {
	rv := objc.Send[MPSCNNConvolutionGradientState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionGradientState) Autorelease() MPSCNNConvolutionGradientState {
	rv := objc.Send[MPSCNNConvolutionGradientState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionGradientState creates a new MPSCNNConvolutionGradientState instance.
func NewMPSCNNConvolutionGradientState() MPSCNNConvolutionGradientState {
	class := getMPSCNNConvolutionGradientStateClass()
	rv := objc.Send[MPSCNNConvolutionGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNConvolutionGradientStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNConvolutionGradientState {
	instance := getMPSCNNConvolutionGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNConvolutionGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNConvolutionGradientStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNConvolutionGradientState {
	instance := getMPSCNNConvolutionGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNConvolutionGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNConvolutionGradientStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNConvolutionGradientState {
	instance := getMPSCNNConvolutionGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNConvolutionGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNConvolutionGradientStateWithResource(resource metal.MTLResource) MPSCNNConvolutionGradientState {
	instance := getMPSCNNConvolutionGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNConvolutionGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNConvolutionGradientStateWithResources(resources []objectivec.IObject) MPSCNNConvolutionGradientState {
	instance := getMPSCNNConvolutionGradientStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNConvolutionGradientStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSizeEncodingState/sourceHeight
func (c MPSCNNConvolutionGradientState) SourceHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSizeEncodingState/sourceWidth
func (c MPSCNNConvolutionGradientState) SourceWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sourceWidth"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientState/convolution
func (c MPSCNNConvolutionGradientState) Convolution() IMPSCNNConvolution {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("convolution"))
	return MPSCNNConvolutionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientState/gradientForBiases
func (c MPSCNNConvolutionGradientState) GradientForBiases() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gradientForBiases"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientState/gradientForWeights
func (c MPSCNNConvolutionGradientState) GradientForWeights() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("gradientForWeights"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientState/gradientForWeightsLayout
func (c MPSCNNConvolutionGradientState) GradientForWeightsLayout() MPSCNNConvolutionWeightsLayout {
	rv := objc.Send[MPSCNNConvolutionWeightsLayout](c.ID, objc.Sel("gradientForWeightsLayout"))
	return MPSCNNConvolutionWeightsLayout(rv)
}

// Protocol methods for MPSImageSizeEncodingState

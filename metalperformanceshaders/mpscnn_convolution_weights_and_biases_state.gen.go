// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNConvolutionWeightsAndBiasesState] class.
var (
	_MPSCNNConvolutionWeightsAndBiasesStateClass     MPSCNNConvolutionWeightsAndBiasesStateClass
	_MPSCNNConvolutionWeightsAndBiasesStateClassOnce sync.Once
)

func getMPSCNNConvolutionWeightsAndBiasesStateClass() MPSCNNConvolutionWeightsAndBiasesStateClass {
	_MPSCNNConvolutionWeightsAndBiasesStateClassOnce.Do(func() {
		_MPSCNNConvolutionWeightsAndBiasesStateClass = MPSCNNConvolutionWeightsAndBiasesStateClass{class: objc.GetClass("MPSCNNConvolutionWeightsAndBiasesState")}
	})
	return _MPSCNNConvolutionWeightsAndBiasesStateClass
}

// GetMPSCNNConvolutionWeightsAndBiasesStateClass returns the class object for MPSCNNConvolutionWeightsAndBiasesState.
func GetMPSCNNConvolutionWeightsAndBiasesStateClass() MPSCNNConvolutionWeightsAndBiasesStateClass {
	return getMPSCNNConvolutionWeightsAndBiasesStateClass()
}

type MPSCNNConvolutionWeightsAndBiasesStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionWeightsAndBiasesStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionWeightsAndBiasesStateClass) Alloc() MPSCNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[MPSCNNConvolutionWeightsAndBiasesState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that stores weights and biases.
//
// # Initializers
//
//   - [MPSCNNConvolutionWeightsAndBiasesState.InitWithDeviceCnnConvolutionDescriptor]
//   - [MPSCNNConvolutionWeightsAndBiasesState.InitWithWeightsBiases]
//   - [MPSCNNConvolutionWeightsAndBiasesState.InitWithWeightsWeightsOffsetBiasesBiasesOffsetCnnConvolutionDescriptor]
//
// # Instance Properties
//
//   - [MPSCNNConvolutionWeightsAndBiasesState.Biases]
//   - [MPSCNNConvolutionWeightsAndBiasesState.BiasesOffset]
//   - [MPSCNNConvolutionWeightsAndBiasesState.Weights]
//   - [MPSCNNConvolutionWeightsAndBiasesState.WeightsOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState
type MPSCNNConvolutionWeightsAndBiasesState struct {
	MPSState
}

// MPSCNNConvolutionWeightsAndBiasesStateFromID constructs a [MPSCNNConvolutionWeightsAndBiasesState] from an objc.ID.
//
// A class that stores weights and biases.
func MPSCNNConvolutionWeightsAndBiasesStateFromID(id objc.ID) MPSCNNConvolutionWeightsAndBiasesState {
	return MPSCNNConvolutionWeightsAndBiasesState{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSCNNConvolutionWeightsAndBiasesState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionWeightsAndBiasesState] class.
//
// # Initializers
//
//   - [IMPSCNNConvolutionWeightsAndBiasesState.InitWithDeviceCnnConvolutionDescriptor]
//   - [IMPSCNNConvolutionWeightsAndBiasesState.InitWithWeightsBiases]
//   - [IMPSCNNConvolutionWeightsAndBiasesState.InitWithWeightsWeightsOffsetBiasesBiasesOffsetCnnConvolutionDescriptor]
//
// # Instance Properties
//
//   - [IMPSCNNConvolutionWeightsAndBiasesState.Biases]
//   - [IMPSCNNConvolutionWeightsAndBiasesState.BiasesOffset]
//   - [IMPSCNNConvolutionWeightsAndBiasesState.Weights]
//   - [IMPSCNNConvolutionWeightsAndBiasesState.WeightsOffset]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState
type IMPSCNNConvolutionWeightsAndBiasesState interface {
	IMPSState

	// Topic: Initializers

	InitWithDeviceCnnConvolutionDescriptor(device metal.MTLDevice, descriptor IMPSCNNConvolutionDescriptor) MPSCNNConvolutionWeightsAndBiasesState
	InitWithWeightsBiases(weights metal.MTLBuffer, biases metal.MTLBuffer) MPSCNNConvolutionWeightsAndBiasesState
	InitWithWeightsWeightsOffsetBiasesBiasesOffsetCnnConvolutionDescriptor(weights metal.MTLBuffer, weightsOffset uint, biases metal.MTLBuffer, biasesOffset uint, descriptor IMPSCNNConvolutionDescriptor) MPSCNNConvolutionWeightsAndBiasesState

	// Topic: Instance Properties

	Biases() metal.MTLBuffer
	BiasesOffset() uint
	Weights() metal.MTLBuffer
	WeightsOffset() uint
}

// Init initializes the instance.
func (c MPSCNNConvolutionWeightsAndBiasesState) Init() MPSCNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[MPSCNNConvolutionWeightsAndBiasesState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionWeightsAndBiasesState) Autorelease() MPSCNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[MPSCNNConvolutionWeightsAndBiasesState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionWeightsAndBiasesState creates a new MPSCNNConvolutionWeightsAndBiasesState instance.
func NewMPSCNNConvolutionWeightsAndBiasesState() MPSCNNConvolutionWeightsAndBiasesState {
	class := getMPSCNNConvolutionWeightsAndBiasesStateClass()
	rv := objc.Send[MPSCNNConvolutionWeightsAndBiasesState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewCNNConvolutionWeightsAndBiasesStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSCNNConvolutionWeightsAndBiasesState {
	instance := getMPSCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/init(device:cnnConvolutionDescriptor:)
func NewCNNConvolutionWeightsAndBiasesStateWithDeviceCnnConvolutionDescriptor(device metal.MTLDevice, descriptor IMPSCNNConvolutionDescriptor) MPSCNNConvolutionWeightsAndBiasesState {
	instance := getMPSCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:cnnConvolutionDescriptor:"), device, descriptor)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewCNNConvolutionWeightsAndBiasesStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSCNNConvolutionWeightsAndBiasesState {
	instance := getMPSCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewCNNConvolutionWeightsAndBiasesStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSCNNConvolutionWeightsAndBiasesState {
	instance := getMPSCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewCNNConvolutionWeightsAndBiasesStateWithResource(resource metal.MTLResource) MPSCNNConvolutionWeightsAndBiasesState {
	instance := getMPSCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewCNNConvolutionWeightsAndBiasesStateWithResources(resources []objectivec.IObject) MPSCNNConvolutionWeightsAndBiasesState {
	instance := getMPSCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/init(weights:biases:)
func NewCNNConvolutionWeightsAndBiasesStateWithWeightsBiases(weights metal.MTLBuffer, biases metal.MTLBuffer) MPSCNNConvolutionWeightsAndBiasesState {
	instance := getMPSCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWeights:biases:"), weights, biases)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/init(weights:weightsOffset:biases:biasesOffset:cnnConvolutionDescriptor:)
func NewCNNConvolutionWeightsAndBiasesStateWithWeightsWeightsOffsetBiasesBiasesOffsetCnnConvolutionDescriptor(weights metal.MTLBuffer, weightsOffset uint, biases metal.MTLBuffer, biasesOffset uint, descriptor IMPSCNNConvolutionDescriptor) MPSCNNConvolutionWeightsAndBiasesState {
	instance := getMPSCNNConvolutionWeightsAndBiasesStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWeights:weightsOffset:biases:biasesOffset:cnnConvolutionDescriptor:"), weights, weightsOffset, biases, biasesOffset, descriptor)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/init(device:cnnConvolutionDescriptor:)
func (c MPSCNNConvolutionWeightsAndBiasesState) InitWithDeviceCnnConvolutionDescriptor(device metal.MTLDevice, descriptor IMPSCNNConvolutionDescriptor) MPSCNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[MPSCNNConvolutionWeightsAndBiasesState](c.ID, objc.Sel("initWithDevice:cnnConvolutionDescriptor:"), device, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/init(weights:biases:)
func (c MPSCNNConvolutionWeightsAndBiasesState) InitWithWeightsBiases(weights metal.MTLBuffer, biases metal.MTLBuffer) MPSCNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[MPSCNNConvolutionWeightsAndBiasesState](c.ID, objc.Sel("initWithWeights:biases:"), weights, biases)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/init(weights:weightsOffset:biases:biasesOffset:cnnConvolutionDescriptor:)
func (c MPSCNNConvolutionWeightsAndBiasesState) InitWithWeightsWeightsOffsetBiasesBiasesOffsetCnnConvolutionDescriptor(weights metal.MTLBuffer, weightsOffset uint, biases metal.MTLBuffer, biasesOffset uint, descriptor IMPSCNNConvolutionDescriptor) MPSCNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[MPSCNNConvolutionWeightsAndBiasesState](c.ID, objc.Sel("initWithWeights:weightsOffset:biases:biasesOffset:cnnConvolutionDescriptor:"), weights, weightsOffset, biases, biasesOffset, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/temporaryCNNConvolutionWeightsAndBiasesState(with:cnnConvolutionDescriptor:)
func (_MPSCNNConvolutionWeightsAndBiasesStateClass MPSCNNConvolutionWeightsAndBiasesStateClass) TemporaryCNNConvolutionWeightsAndBiasesStateWithCommandBufferCnnConvolutionDescriptor(commandBuffer metal.MTLCommandBuffer, descriptor IMPSCNNConvolutionDescriptor) MPSCNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNConvolutionWeightsAndBiasesStateClass.class), objc.Sel("temporaryCNNConvolutionWeightsAndBiasesStateWithCommandBuffer:cnnConvolutionDescriptor:"), commandBuffer, descriptor)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/biases
func (c MPSCNNConvolutionWeightsAndBiasesState) Biases() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("biases"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/biasesOffset
func (c MPSCNNConvolutionWeightsAndBiasesState) BiasesOffset() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("biasesOffset"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/weights
func (c MPSCNNConvolutionWeightsAndBiasesState) Weights() metal.MTLBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("weights"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionWeightsAndBiasesState/weightsOffset
func (c MPSCNNConvolutionWeightsAndBiasesState) WeightsOffset() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("weightsOffset"))
	return rv
}

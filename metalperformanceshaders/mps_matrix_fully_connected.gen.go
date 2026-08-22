// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixFullyConnected] class.
var (
	_MPSMatrixFullyConnectedClass     MPSMatrixFullyConnectedClass
	_MPSMatrixFullyConnectedClassOnce sync.Once
)

func getMPSMatrixFullyConnectedClass() MPSMatrixFullyConnectedClass {
	_MPSMatrixFullyConnectedClassOnce.Do(func() {
		_MPSMatrixFullyConnectedClass = MPSMatrixFullyConnectedClass{class: objc.GetClass("MPSMatrixFullyConnected")}
	})
	return _MPSMatrixFullyConnectedClass
}

// GetMPSMatrixFullyConnectedClass returns the class object for MPSMatrixFullyConnected.
func GetMPSMatrixFullyConnectedClass() MPSMatrixFullyConnectedClass {
	return getMPSMatrixFullyConnectedClass()
}

type MPSMatrixFullyConnectedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixFullyConnectedClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixFullyConnectedClass) Alloc() MPSMatrixFullyConnected {
	rv := objc.Send[MPSMatrixFullyConnected](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel for applying a fully connected neural network layer.
//
// # Instance Properties
//
//   - [MPSMatrixFullyConnected.Alpha]
//   - [MPSMatrixFullyConnected.SetAlpha]
//   - [MPSMatrixFullyConnected.SourceInputFeatureChannels]
//   - [MPSMatrixFullyConnected.SetSourceInputFeatureChannels]
//   - [MPSMatrixFullyConnected.SourceNumberOfFeatureVectors]
//   - [MPSMatrixFullyConnected.SetSourceNumberOfFeatureVectors]
//   - [MPSMatrixFullyConnected.SourceOutputFeatureChannels]
//   - [MPSMatrixFullyConnected.SetSourceOutputFeatureChannels]
//
// # Instance Methods
//
//   - [MPSMatrixFullyConnected.EncodeToCommandBufferInputMatrixWeightMatrixBiasVectorResultMatrix]
//   - [MPSMatrixFullyConnected.NeuronParameterA]
//   - [MPSMatrixFullyConnected.NeuronParameterB]
//   - [MPSMatrixFullyConnected.NeuronParameterC]
//   - [MPSMatrixFullyConnected.NeuronType]
//   - [MPSMatrixFullyConnected.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected
type MPSMatrixFullyConnected struct {
	MPSMatrixBinaryKernel
}

// MPSMatrixFullyConnectedFromID constructs a [MPSMatrixFullyConnected] from an objc.ID.
//
// A kernel for applying a fully connected neural network layer.
func MPSMatrixFullyConnectedFromID(id objc.ID) MPSMatrixFullyConnected {
	return MPSMatrixFullyConnected{MPSMatrixBinaryKernel: MPSMatrixBinaryKernelFromID(id)}
}

// NOTE: MPSMatrixFullyConnected adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixFullyConnected] class.
//
// # Instance Properties
//
//   - [IMPSMatrixFullyConnected.Alpha]
//   - [IMPSMatrixFullyConnected.SetAlpha]
//   - [IMPSMatrixFullyConnected.SourceInputFeatureChannels]
//   - [IMPSMatrixFullyConnected.SetSourceInputFeatureChannels]
//   - [IMPSMatrixFullyConnected.SourceNumberOfFeatureVectors]
//   - [IMPSMatrixFullyConnected.SetSourceNumberOfFeatureVectors]
//   - [IMPSMatrixFullyConnected.SourceOutputFeatureChannels]
//   - [IMPSMatrixFullyConnected.SetSourceOutputFeatureChannels]
//
// # Instance Methods
//
//   - [IMPSMatrixFullyConnected.EncodeToCommandBufferInputMatrixWeightMatrixBiasVectorResultMatrix]
//   - [IMPSMatrixFullyConnected.NeuronParameterA]
//   - [IMPSMatrixFullyConnected.NeuronParameterB]
//   - [IMPSMatrixFullyConnected.NeuronParameterC]
//   - [IMPSMatrixFullyConnected.NeuronType]
//   - [IMPSMatrixFullyConnected.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected
type IMPSMatrixFullyConnected interface {
	IMPSMatrixBinaryKernel

	// Topic: Instance Properties

	Alpha() float64
	SetAlpha(value float64)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)
	SourceOutputFeatureChannels() uint
	SetSourceOutputFeatureChannels(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferInputMatrixWeightMatrixBiasVectorResultMatrix(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, weightMatrix IMPSMatrix, biasVector IMPSVector, resultMatrix IMPSMatrix)
	NeuronParameterA() float32
	NeuronParameterB() float32
	NeuronParameterC() float32
	NeuronType() MPSCNNNeuronType
	SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32)
}

// Init initializes the instance.
func (m MPSMatrixFullyConnected) Init() MPSMatrixFullyConnected {
	rv := objc.Send[MPSMatrixFullyConnected](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixFullyConnected) Autorelease() MPSMatrixFullyConnected {
	rv := objc.Send[MPSMatrixFullyConnected](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixFullyConnected creates a new MPSMatrixFullyConnected instance.
func NewMPSMatrixFullyConnected() MPSMatrixFullyConnected {
	class := getMPSMatrixFullyConnectedClass()
	rv := objc.Send[MPSMatrixFullyConnected](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixFullyConnectedWithCoder(aDecoder foundation.INSCoder) MPSMatrixFullyConnected {
	instance := getMPSMatrixFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixFullyConnectedFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/init(coder:device:)
func NewMatrixFullyConnectedWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixFullyConnected {
	instance := getMPSMatrixFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixFullyConnectedFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/init(device:)
func NewMatrixFullyConnectedWithDevice(device metal.MTLDevice) MPSMatrixFullyConnected {
	instance := getMPSMatrixFullyConnectedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixFullyConnectedFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/encode(commandBuffer:inputMatrix:weightMatrix:biasVector:resultMatrix:)
func (m MPSMatrixFullyConnected) EncodeToCommandBufferInputMatrixWeightMatrixBiasVectorResultMatrix(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, weightMatrix IMPSMatrix, biasVector IMPSVector, resultMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:weightMatrix:biasVector:resultMatrix:"), commandBuffer, inputMatrix, weightMatrix, biasVector, resultMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/neuronParameterA()
func (m MPSMatrixFullyConnected) NeuronParameterA() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterA"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/neuronParameterB()
func (m MPSMatrixFullyConnected) NeuronParameterB() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterB"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/neuronParameterC()
func (m MPSMatrixFullyConnected) NeuronParameterC() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterC"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/neuronType()
func (m MPSMatrixFullyConnected) NeuronType() MPSCNNNeuronType {
	rv := objc.Send[MPSCNNNeuronType](m.ID, objc.Sel("neuronType"))
	return MPSCNNNeuronType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/setNeuronType(_:parameterA:parameterB:parameterC:)
func (m MPSMatrixFullyConnected) SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Send[objc.ID](m.ID, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/alpha
func (m MPSMatrixFullyConnected) Alpha() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("alpha"))
	return rv
}
func (m MPSMatrixFullyConnected) SetAlpha(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/sourceInputFeatureChannels
func (m MPSMatrixFullyConnected) SourceInputFeatureChannels() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}
func (m MPSMatrixFullyConnected) SetSourceInputFeatureChannels(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/sourceNumberOfFeatureVectors
func (m MPSMatrixFullyConnected) SourceNumberOfFeatureVectors() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}
func (m MPSMatrixFullyConnected) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnected/sourceOutputFeatureChannels
func (m MPSMatrixFullyConnected) SourceOutputFeatureChannels() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceOutputFeatureChannels"))
	return rv
}
func (m MPSMatrixFullyConnected) SetSourceOutputFeatureChannels(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceOutputFeatureChannels:"), value)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixNeuron] class.
var (
	_MPSMatrixNeuronClass     MPSMatrixNeuronClass
	_MPSMatrixNeuronClassOnce sync.Once
)

func getMPSMatrixNeuronClass() MPSMatrixNeuronClass {
	_MPSMatrixNeuronClassOnce.Do(func() {
		_MPSMatrixNeuronClass = MPSMatrixNeuronClass{class: objc.GetClass("MPSMatrixNeuron")}
	})
	return _MPSMatrixNeuronClass
}

// GetMPSMatrixNeuronClass returns the class object for MPSMatrixNeuron.
func GetMPSMatrixNeuronClass() MPSMatrixNeuronClass {
	return getMPSMatrixNeuronClass()
}

type MPSMatrixNeuronClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixNeuronClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixNeuronClass) Alloc() MPSMatrixNeuron {
	rv := objc.Send[MPSMatrixNeuron](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A neuron activation kernel that operates on matrices.
//
// # Instance Properties
//
//   - [MPSMatrixNeuron.Alpha]
//   - [MPSMatrixNeuron.SetAlpha]
//   - [MPSMatrixNeuron.SourceInputFeatureChannels]
//   - [MPSMatrixNeuron.SetSourceInputFeatureChannels]
//   - [MPSMatrixNeuron.SourceNumberOfFeatureVectors]
//   - [MPSMatrixNeuron.SetSourceNumberOfFeatureVectors]
//
// # Instance Methods
//
//   - [MPSMatrixNeuron.EncodeToCommandBufferInputMatrixBiasVectorResultMatrix]
//   - [MPSMatrixNeuron.NeuronParameterA]
//   - [MPSMatrixNeuron.NeuronParameterB]
//   - [MPSMatrixNeuron.NeuronParameterC]
//   - [MPSMatrixNeuron.NeuronType]
//   - [MPSMatrixNeuron.SetNeuronToPReLUWithParametersA]
//   - [MPSMatrixNeuron.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron
type MPSMatrixNeuron struct {
	MPSMatrixUnaryKernel
}

// MPSMatrixNeuronFromID constructs a [MPSMatrixNeuron] from an objc.ID.
//
// A neuron activation kernel that operates on matrices.
func MPSMatrixNeuronFromID(id objc.ID) MPSMatrixNeuron {
	return MPSMatrixNeuron{MPSMatrixUnaryKernel: MPSMatrixUnaryKernelFromID(id)}
}

// NOTE: MPSMatrixNeuron adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixNeuron] class.
//
// # Instance Properties
//
//   - [IMPSMatrixNeuron.Alpha]
//   - [IMPSMatrixNeuron.SetAlpha]
//   - [IMPSMatrixNeuron.SourceInputFeatureChannels]
//   - [IMPSMatrixNeuron.SetSourceInputFeatureChannels]
//   - [IMPSMatrixNeuron.SourceNumberOfFeatureVectors]
//   - [IMPSMatrixNeuron.SetSourceNumberOfFeatureVectors]
//
// # Instance Methods
//
//   - [IMPSMatrixNeuron.EncodeToCommandBufferInputMatrixBiasVectorResultMatrix]
//   - [IMPSMatrixNeuron.NeuronParameterA]
//   - [IMPSMatrixNeuron.NeuronParameterB]
//   - [IMPSMatrixNeuron.NeuronParameterC]
//   - [IMPSMatrixNeuron.NeuronType]
//   - [IMPSMatrixNeuron.SetNeuronToPReLUWithParametersA]
//   - [IMPSMatrixNeuron.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron
type IMPSMatrixNeuron interface {
	IMPSMatrixUnaryKernel

	// Topic: Instance Properties

	Alpha() float64
	SetAlpha(value float64)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferInputMatrixBiasVectorResultMatrix(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, biasVector IMPSVector, resultMatrix IMPSMatrix)
	NeuronParameterA() float32
	NeuronParameterB() float32
	NeuronParameterC() float32
	NeuronType() MPSCNNNeuronType
	SetNeuronToPReLUWithParametersA(A foundation.NSData)
	SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32)
}

// Init initializes the instance.
func (m MPSMatrixNeuron) Init() MPSMatrixNeuron {
	rv := objc.Send[MPSMatrixNeuron](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixNeuron) Autorelease() MPSMatrixNeuron {
	rv := objc.Send[MPSMatrixNeuron](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixNeuron creates a new MPSMatrixNeuron instance.
func NewMPSMatrixNeuron() MPSMatrixNeuron {
	class := getMPSMatrixNeuronClass()
	rv := objc.Send[MPSMatrixNeuron](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixNeuronWithCoder(aDecoder foundation.INSCoder) MPSMatrixNeuron {
	instance := getMPSMatrixNeuronClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixNeuronFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/init(coder:device:)
func NewMatrixNeuronWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixNeuron {
	instance := getMPSMatrixNeuronClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixNeuronFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/init(device:)
func NewMatrixNeuronWithDevice(device metal.MTLDevice) MPSMatrixNeuron {
	instance := getMPSMatrixNeuronClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixNeuronFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/encode(commandBuffer:inputMatrix:biasVector:resultMatrix:)
func (m MPSMatrixNeuron) EncodeToCommandBufferInputMatrixBiasVectorResultMatrix(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, biasVector IMPSVector, resultMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:biasVector:resultMatrix:"), commandBuffer, inputMatrix, biasVector, resultMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/neuronParameterA()
func (m MPSMatrixNeuron) NeuronParameterA() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterA"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/neuronParameterB()
func (m MPSMatrixNeuron) NeuronParameterB() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterB"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/neuronParameterC()
func (m MPSMatrixNeuron) NeuronParameterC() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterC"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/neuronType()
func (m MPSMatrixNeuron) NeuronType() MPSCNNNeuronType {
	rv := objc.Send[MPSCNNNeuronType](m.ID, objc.Sel("neuronType"))
	return MPSCNNNeuronType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/setNeuronToPReLUWithParametersA(_:)
func (m MPSMatrixNeuron) SetNeuronToPReLUWithParametersA(A foundation.NSData) {
	objc.Send[objc.ID](m.ID, objc.Sel("setNeuronToPReLUWithParametersA:"), A)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/setNeuronType(_:parameterA:parameterB:parameterC:)
func (m MPSMatrixNeuron) SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Send[objc.ID](m.ID, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/alpha
func (m MPSMatrixNeuron) Alpha() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("alpha"))
	return rv
}
func (m MPSMatrixNeuron) SetAlpha(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/sourceInputFeatureChannels
func (m MPSMatrixNeuron) SourceInputFeatureChannels() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}
func (m MPSMatrixNeuron) SetSourceInputFeatureChannels(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron/sourceNumberOfFeatureVectors
func (m MPSMatrixNeuron) SourceNumberOfFeatureVectors() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}
func (m MPSMatrixNeuron) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

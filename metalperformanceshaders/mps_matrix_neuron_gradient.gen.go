// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixNeuronGradient] class.
var (
	_MPSMatrixNeuronGradientClass     MPSMatrixNeuronGradientClass
	_MPSMatrixNeuronGradientClassOnce sync.Once
)

func getMPSMatrixNeuronGradientClass() MPSMatrixNeuronGradientClass {
	_MPSMatrixNeuronGradientClassOnce.Do(func() {
		_MPSMatrixNeuronGradientClass = MPSMatrixNeuronGradientClass{class: objc.GetClass("MPSMatrixNeuronGradient")}
	})
	return _MPSMatrixNeuronGradientClass
}

// GetMPSMatrixNeuronGradientClass returns the class object for MPSMatrixNeuronGradient.
func GetMPSMatrixNeuronGradientClass() MPSMatrixNeuronGradientClass {
	return getMPSMatrixNeuronGradientClass()
}

type MPSMatrixNeuronGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixNeuronGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixNeuronGradientClass) Alloc() MPSMatrixNeuronGradient {
	rv := objc.Send[MPSMatrixNeuronGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient neuron activation kernel that operates on matrices.
//
// # Instance Properties
//
//   - [MPSMatrixNeuronGradient.Alpha]
//   - [MPSMatrixNeuronGradient.SetAlpha]
//   - [MPSMatrixNeuronGradient.SourceInputFeatureChannels]
//   - [MPSMatrixNeuronGradient.SetSourceInputFeatureChannels]
//   - [MPSMatrixNeuronGradient.SourceNumberOfFeatureVectors]
//   - [MPSMatrixNeuronGradient.SetSourceNumberOfFeatureVectors]
//
// # Instance Methods
//
//   - [MPSMatrixNeuronGradient.EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector]
//   - [MPSMatrixNeuronGradient.NeuronParameterA]
//   - [MPSMatrixNeuronGradient.NeuronParameterB]
//   - [MPSMatrixNeuronGradient.NeuronParameterC]
//   - [MPSMatrixNeuronGradient.NeuronType]
//   - [MPSMatrixNeuronGradient.SetNeuronToPReLUWithParametersA]
//   - [MPSMatrixNeuronGradient.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient
type MPSMatrixNeuronGradient struct {
	MPSMatrixBinaryKernel
}

// MPSMatrixNeuronGradientFromID constructs a [MPSMatrixNeuronGradient] from an objc.ID.
//
// A gradient neuron activation kernel that operates on matrices.
func MPSMatrixNeuronGradientFromID(id objc.ID) MPSMatrixNeuronGradient {
	return MPSMatrixNeuronGradient{MPSMatrixBinaryKernel: MPSMatrixBinaryKernelFromID(id)}
}

// NOTE: MPSMatrixNeuronGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixNeuronGradient] class.
//
// # Instance Properties
//
//   - [IMPSMatrixNeuronGradient.Alpha]
//   - [IMPSMatrixNeuronGradient.SetAlpha]
//   - [IMPSMatrixNeuronGradient.SourceInputFeatureChannels]
//   - [IMPSMatrixNeuronGradient.SetSourceInputFeatureChannels]
//   - [IMPSMatrixNeuronGradient.SourceNumberOfFeatureVectors]
//   - [IMPSMatrixNeuronGradient.SetSourceNumberOfFeatureVectors]
//
// # Instance Methods
//
//   - [IMPSMatrixNeuronGradient.EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector]
//   - [IMPSMatrixNeuronGradient.NeuronParameterA]
//   - [IMPSMatrixNeuronGradient.NeuronParameterB]
//   - [IMPSMatrixNeuronGradient.NeuronParameterC]
//   - [IMPSMatrixNeuronGradient.NeuronType]
//   - [IMPSMatrixNeuronGradient.SetNeuronToPReLUWithParametersA]
//   - [IMPSMatrixNeuronGradient.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient
type IMPSMatrixNeuronGradient interface {
	IMPSMatrixBinaryKernel

	// Topic: Instance Properties

	Alpha() float64
	SetAlpha(value float64)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector(commandBuffer metal.MTLCommandBuffer, gradientMatrix IMPSMatrix, inputMatrix IMPSMatrix, biasVector IMPSVector, resultGradientForDataMatrix IMPSMatrix, resultGradientForBiasVector IMPSVector)
	NeuronParameterA() float32
	NeuronParameterB() float32
	NeuronParameterC() float32
	NeuronType() MPSCNNNeuronType
	SetNeuronToPReLUWithParametersA(A foundation.NSData)
	SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32)
}

// Init initializes the instance.
func (m MPSMatrixNeuronGradient) Init() MPSMatrixNeuronGradient {
	rv := objc.Send[MPSMatrixNeuronGradient](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixNeuronGradient) Autorelease() MPSMatrixNeuronGradient {
	rv := objc.Send[MPSMatrixNeuronGradient](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixNeuronGradient creates a new MPSMatrixNeuronGradient instance.
func NewMPSMatrixNeuronGradient() MPSMatrixNeuronGradient {
	class := getMPSMatrixNeuronGradientClass()
	rv := objc.Send[MPSMatrixNeuronGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixNeuronGradientWithCoder(aDecoder foundation.INSCoder) MPSMatrixNeuronGradient {
	instance := getMPSMatrixNeuronGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixNeuronGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/init(coder:device:)
func NewMatrixNeuronGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixNeuronGradient {
	instance := getMPSMatrixNeuronGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixNeuronGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/init(device:)
func NewMatrixNeuronGradientWithDevice(device metal.MTLDevice) MPSMatrixNeuronGradient {
	instance := getMPSMatrixNeuronGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixNeuronGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/encode(to:gradientMatrix:inputMatrix:biasVector:resultGradientForDataMatrix:resultGradientForBiasVector:)
func (m MPSMatrixNeuronGradient) EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector(commandBuffer metal.MTLCommandBuffer, gradientMatrix IMPSMatrix, inputMatrix IMPSMatrix, biasVector IMPSVector, resultGradientForDataMatrix IMPSMatrix, resultGradientForBiasVector IMPSVector) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:gradientMatrix:inputMatrix:biasVector:resultGradientForDataMatrix:resultGradientForBiasVector:"), commandBuffer, gradientMatrix, inputMatrix, biasVector, resultGradientForDataMatrix, resultGradientForBiasVector)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/neuronParameterA()
func (m MPSMatrixNeuronGradient) NeuronParameterA() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterA"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/neuronParameterB()
func (m MPSMatrixNeuronGradient) NeuronParameterB() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterB"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/neuronParameterC()
func (m MPSMatrixNeuronGradient) NeuronParameterC() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterC"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/neuronType()
func (m MPSMatrixNeuronGradient) NeuronType() MPSCNNNeuronType {
	rv := objc.Send[MPSCNNNeuronType](m.ID, objc.Sel("neuronType"))
	return MPSCNNNeuronType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/setNeuronToPReLUWithParametersA(_:)
func (m MPSMatrixNeuronGradient) SetNeuronToPReLUWithParametersA(A foundation.NSData) {
	objc.Send[objc.ID](m.ID, objc.Sel("setNeuronToPReLUWithParametersA:"), A)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/setNeuronType(_:parameterA:parameterB:parameterC:)
func (m MPSMatrixNeuronGradient) SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Send[objc.ID](m.ID, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/alpha
func (m MPSMatrixNeuronGradient) Alpha() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("alpha"))
	return rv
}
func (m MPSMatrixNeuronGradient) SetAlpha(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/sourceInputFeatureChannels
func (m MPSMatrixNeuronGradient) SourceInputFeatureChannels() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}
func (m MPSMatrixNeuronGradient) SetSourceInputFeatureChannels(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuronGradient/sourceNumberOfFeatureVectors
func (m MPSMatrixNeuronGradient) SourceNumberOfFeatureVectors() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}
func (m MPSMatrixNeuronGradient) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

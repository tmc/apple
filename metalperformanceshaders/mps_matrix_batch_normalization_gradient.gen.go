// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixBatchNormalizationGradient] class.
var (
	_MPSMatrixBatchNormalizationGradientClass     MPSMatrixBatchNormalizationGradientClass
	_MPSMatrixBatchNormalizationGradientClassOnce sync.Once
)

func getMPSMatrixBatchNormalizationGradientClass() MPSMatrixBatchNormalizationGradientClass {
	_MPSMatrixBatchNormalizationGradientClassOnce.Do(func() {
		_MPSMatrixBatchNormalizationGradientClass = MPSMatrixBatchNormalizationGradientClass{class: objc.GetClass("MPSMatrixBatchNormalizationGradient")}
	})
	return _MPSMatrixBatchNormalizationGradientClass
}

// GetMPSMatrixBatchNormalizationGradientClass returns the class object for MPSMatrixBatchNormalizationGradient.
func GetMPSMatrixBatchNormalizationGradientClass() MPSMatrixBatchNormalizationGradientClass {
	return getMPSMatrixBatchNormalizationGradientClass()
}

type MPSMatrixBatchNormalizationGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixBatchNormalizationGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixBatchNormalizationGradientClass) Alloc() MPSMatrixBatchNormalizationGradient {
	rv := objc.Send[MPSMatrixBatchNormalizationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A batch normalization gradient kernel that operates on matrices.
//
// # Instance Properties
//
//   - [MPSMatrixBatchNormalizationGradient.Epsilon]
//   - [MPSMatrixBatchNormalizationGradient.SetEpsilon]
//   - [MPSMatrixBatchNormalizationGradient.SourceInputFeatureChannels]
//   - [MPSMatrixBatchNormalizationGradient.SetSourceInputFeatureChannels]
//   - [MPSMatrixBatchNormalizationGradient.SourceNumberOfFeatureVectors]
//   - [MPSMatrixBatchNormalizationGradient.SetSourceNumberOfFeatureVectors]
//
// # Instance Methods
//
//   - [MPSMatrixBatchNormalizationGradient.EncodeToCommandBufferGradientMatrixInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultGradientForDataMatrixResultGradientForGammaVectorResultGradientForBetaVector]
//   - [MPSMatrixBatchNormalizationGradient.NeuronParameterA]
//   - [MPSMatrixBatchNormalizationGradient.NeuronParameterB]
//   - [MPSMatrixBatchNormalizationGradient.NeuronParameterC]
//   - [MPSMatrixBatchNormalizationGradient.NeuronType]
//   - [MPSMatrixBatchNormalizationGradient.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient
type MPSMatrixBatchNormalizationGradient struct {
	MPSMatrixBinaryKernel
}

// MPSMatrixBatchNormalizationGradientFromID constructs a [MPSMatrixBatchNormalizationGradient] from an objc.ID.
//
// A batch normalization gradient kernel that operates on matrices.
func MPSMatrixBatchNormalizationGradientFromID(id objc.ID) MPSMatrixBatchNormalizationGradient {
	return MPSMatrixBatchNormalizationGradient{MPSMatrixBinaryKernel: MPSMatrixBinaryKernelFromID(id)}
}

// NOTE: MPSMatrixBatchNormalizationGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixBatchNormalizationGradient] class.
//
// # Instance Properties
//
//   - [IMPSMatrixBatchNormalizationGradient.Epsilon]
//   - [IMPSMatrixBatchNormalizationGradient.SetEpsilon]
//   - [IMPSMatrixBatchNormalizationGradient.SourceInputFeatureChannels]
//   - [IMPSMatrixBatchNormalizationGradient.SetSourceInputFeatureChannels]
//   - [IMPSMatrixBatchNormalizationGradient.SourceNumberOfFeatureVectors]
//   - [IMPSMatrixBatchNormalizationGradient.SetSourceNumberOfFeatureVectors]
//
// # Instance Methods
//
//   - [IMPSMatrixBatchNormalizationGradient.EncodeToCommandBufferGradientMatrixInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultGradientForDataMatrixResultGradientForGammaVectorResultGradientForBetaVector]
//   - [IMPSMatrixBatchNormalizationGradient.NeuronParameterA]
//   - [IMPSMatrixBatchNormalizationGradient.NeuronParameterB]
//   - [IMPSMatrixBatchNormalizationGradient.NeuronParameterC]
//   - [IMPSMatrixBatchNormalizationGradient.NeuronType]
//   - [IMPSMatrixBatchNormalizationGradient.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient
type IMPSMatrixBatchNormalizationGradient interface {
	IMPSMatrixBinaryKernel

	// Topic: Instance Properties

	Epsilon() float32
	SetEpsilon(value float32)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferGradientMatrixInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultGradientForDataMatrixResultGradientForGammaVectorResultGradientForBetaVector(commandBuffer metal.MTLCommandBuffer, gradientMatrix IMPSMatrix, inputMatrix IMPSMatrix, meanVector IMPSVector, varianceVector IMPSVector, gammaVector IMPSVector, betaVector IMPSVector, resultGradientForDataMatrix IMPSMatrix, resultGradientForGammaVector IMPSVector, resultGradientForBetaVector IMPSVector)
	NeuronParameterA() float32
	NeuronParameterB() float32
	NeuronParameterC() float32
	NeuronType() MPSCNNNeuronType
	SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32)
}

// Init initializes the instance.
func (m MPSMatrixBatchNormalizationGradient) Init() MPSMatrixBatchNormalizationGradient {
	rv := objc.Send[MPSMatrixBatchNormalizationGradient](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixBatchNormalizationGradient) Autorelease() MPSMatrixBatchNormalizationGradient {
	rv := objc.Send[MPSMatrixBatchNormalizationGradient](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixBatchNormalizationGradient creates a new MPSMatrixBatchNormalizationGradient instance.
func NewMPSMatrixBatchNormalizationGradient() MPSMatrixBatchNormalizationGradient {
	class := getMPSMatrixBatchNormalizationGradientClass()
	rv := objc.Send[MPSMatrixBatchNormalizationGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixBatchNormalizationGradientWithCoder(aDecoder foundation.INSCoder) MPSMatrixBatchNormalizationGradient {
	instance := getMPSMatrixBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixBatchNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/init(coder:device:)
func NewMatrixBatchNormalizationGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixBatchNormalizationGradient {
	instance := getMPSMatrixBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixBatchNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/init(device:)
func NewMatrixBatchNormalizationGradientWithDevice(device metal.MTLDevice) MPSMatrixBatchNormalizationGradient {
	instance := getMPSMatrixBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixBatchNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/encode(to:gradientMatrix:inputMatrix:mean:varianceVector:gammaVector:betaVector:resultGradientForDataMatrix:resultGradientForGammaVector:resultGradientForBetaVector:)
func (m MPSMatrixBatchNormalizationGradient) EncodeToCommandBufferGradientMatrixInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultGradientForDataMatrixResultGradientForGammaVectorResultGradientForBetaVector(commandBuffer metal.MTLCommandBuffer, gradientMatrix IMPSMatrix, inputMatrix IMPSMatrix, meanVector IMPSVector, varianceVector IMPSVector, gammaVector IMPSVector, betaVector IMPSVector, resultGradientForDataMatrix IMPSMatrix, resultGradientForGammaVector IMPSVector, resultGradientForBetaVector IMPSVector) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:gradientMatrix:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultGradientForDataMatrix:resultGradientForGammaVector:resultGradientForBetaVector:"), commandBuffer, gradientMatrix, inputMatrix, meanVector, varianceVector, gammaVector, betaVector, resultGradientForDataMatrix, resultGradientForGammaVector, resultGradientForBetaVector)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/neuronParameterA()
func (m MPSMatrixBatchNormalizationGradient) NeuronParameterA() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterA"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/neuronParameterB()
func (m MPSMatrixBatchNormalizationGradient) NeuronParameterB() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterB"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/neuronParameterC()
func (m MPSMatrixBatchNormalizationGradient) NeuronParameterC() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterC"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/neuronType()
func (m MPSMatrixBatchNormalizationGradient) NeuronType() MPSCNNNeuronType {
	rv := objc.Send[MPSCNNNeuronType](m.ID, objc.Sel("neuronType"))
	return MPSCNNNeuronType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/setNeuronType(_:parameterA:parameterB:parameterC:)
func (m MPSMatrixBatchNormalizationGradient) SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Send[objc.ID](m.ID, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/epsilon
func (m MPSMatrixBatchNormalizationGradient) Epsilon() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("epsilon"))
	return rv
}
func (m MPSMatrixBatchNormalizationGradient) SetEpsilon(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setEpsilon:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/sourceInputFeatureChannels
func (m MPSMatrixBatchNormalizationGradient) SourceInputFeatureChannels() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}
func (m MPSMatrixBatchNormalizationGradient) SetSourceInputFeatureChannels(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalizationGradient/sourceNumberOfFeatureVectors
func (m MPSMatrixBatchNormalizationGradient) SourceNumberOfFeatureVectors() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}
func (m MPSMatrixBatchNormalizationGradient) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixBatchNormalization] class.
var (
	_MPSMatrixBatchNormalizationClass     MPSMatrixBatchNormalizationClass
	_MPSMatrixBatchNormalizationClassOnce sync.Once
)

func getMPSMatrixBatchNormalizationClass() MPSMatrixBatchNormalizationClass {
	_MPSMatrixBatchNormalizationClassOnce.Do(func() {
		_MPSMatrixBatchNormalizationClass = MPSMatrixBatchNormalizationClass{class: objc.GetClass("MPSMatrixBatchNormalization")}
	})
	return _MPSMatrixBatchNormalizationClass
}

// GetMPSMatrixBatchNormalizationClass returns the class object for MPSMatrixBatchNormalization.
func GetMPSMatrixBatchNormalizationClass() MPSMatrixBatchNormalizationClass {
	return getMPSMatrixBatchNormalizationClass()
}

type MPSMatrixBatchNormalizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixBatchNormalizationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixBatchNormalizationClass) Alloc() MPSMatrixBatchNormalization {
	rv := objc.Send[MPSMatrixBatchNormalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A batch normalization kernel that operates on matrices.
//
// # Instance Properties
//
//   - [MPSMatrixBatchNormalization.ComputeStatistics]
//   - [MPSMatrixBatchNormalization.SetComputeStatistics]
//   - [MPSMatrixBatchNormalization.Epsilon]
//   - [MPSMatrixBatchNormalization.SetEpsilon]
//   - [MPSMatrixBatchNormalization.SourceInputFeatureChannels]
//   - [MPSMatrixBatchNormalization.SetSourceInputFeatureChannels]
//   - [MPSMatrixBatchNormalization.SourceNumberOfFeatureVectors]
//   - [MPSMatrixBatchNormalization.SetSourceNumberOfFeatureVectors]
//
// # Instance Methods
//
//   - [MPSMatrixBatchNormalization.EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix]
//   - [MPSMatrixBatchNormalization.NeuronParameterA]
//   - [MPSMatrixBatchNormalization.NeuronParameterB]
//   - [MPSMatrixBatchNormalization.NeuronParameterC]
//   - [MPSMatrixBatchNormalization.NeuronType]
//   - [MPSMatrixBatchNormalization.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization
type MPSMatrixBatchNormalization struct {
	MPSMatrixUnaryKernel
}

// MPSMatrixBatchNormalizationFromID constructs a [MPSMatrixBatchNormalization] from an objc.ID.
//
// A batch normalization kernel that operates on matrices.
func MPSMatrixBatchNormalizationFromID(id objc.ID) MPSMatrixBatchNormalization {
	return MPSMatrixBatchNormalization{MPSMatrixUnaryKernel: MPSMatrixUnaryKernelFromID(id)}
}

// NOTE: MPSMatrixBatchNormalization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixBatchNormalization] class.
//
// # Instance Properties
//
//   - [IMPSMatrixBatchNormalization.ComputeStatistics]
//   - [IMPSMatrixBatchNormalization.SetComputeStatistics]
//   - [IMPSMatrixBatchNormalization.Epsilon]
//   - [IMPSMatrixBatchNormalization.SetEpsilon]
//   - [IMPSMatrixBatchNormalization.SourceInputFeatureChannels]
//   - [IMPSMatrixBatchNormalization.SetSourceInputFeatureChannels]
//   - [IMPSMatrixBatchNormalization.SourceNumberOfFeatureVectors]
//   - [IMPSMatrixBatchNormalization.SetSourceNumberOfFeatureVectors]
//
// # Instance Methods
//
//   - [IMPSMatrixBatchNormalization.EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix]
//   - [IMPSMatrixBatchNormalization.NeuronParameterA]
//   - [IMPSMatrixBatchNormalization.NeuronParameterB]
//   - [IMPSMatrixBatchNormalization.NeuronParameterC]
//   - [IMPSMatrixBatchNormalization.NeuronType]
//   - [IMPSMatrixBatchNormalization.SetNeuronTypeParameterAParameterBParameterC]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization
type IMPSMatrixBatchNormalization interface {
	IMPSMatrixUnaryKernel

	// Topic: Instance Properties

	ComputeStatistics() bool
	SetComputeStatistics(value bool)
	Epsilon() float32
	SetEpsilon(value float32)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, meanVector IMPSVector, varianceVector IMPSVector, gammaVector IMPSVector, betaVector IMPSVector, resultMatrix IMPSMatrix)
	NeuronParameterA() float32
	NeuronParameterB() float32
	NeuronParameterC() float32
	NeuronType() MPSCNNNeuronType
	SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32)
}

// Init initializes the instance.
func (m MPSMatrixBatchNormalization) Init() MPSMatrixBatchNormalization {
	rv := objc.Send[MPSMatrixBatchNormalization](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixBatchNormalization) Autorelease() MPSMatrixBatchNormalization {
	rv := objc.Send[MPSMatrixBatchNormalization](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixBatchNormalization creates a new MPSMatrixBatchNormalization instance.
func NewMPSMatrixBatchNormalization() MPSMatrixBatchNormalization {
	class := getMPSMatrixBatchNormalizationClass()
	rv := objc.Send[MPSMatrixBatchNormalization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixBatchNormalizationWithCoder(aDecoder foundation.INSCoder) MPSMatrixBatchNormalization {
	instance := getMPSMatrixBatchNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixBatchNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/init(coder:device:)
func NewMatrixBatchNormalizationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixBatchNormalization {
	instance := getMPSMatrixBatchNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixBatchNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/init(device:)
func NewMatrixBatchNormalizationWithDevice(device metal.MTLDevice) MPSMatrixBatchNormalization {
	instance := getMPSMatrixBatchNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixBatchNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/encode(commandBuffer:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultMatrix:)
func (m MPSMatrixBatchNormalization) EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBuffer metal.MTLCommandBuffer, inputMatrix IMPSMatrix, meanVector IMPSVector, varianceVector IMPSVector, gammaVector IMPSVector, betaVector IMPSVector, resultMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultMatrix:"), commandBuffer, inputMatrix, meanVector, varianceVector, gammaVector, betaVector, resultMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/neuronParameterA()
func (m MPSMatrixBatchNormalization) NeuronParameterA() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterA"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/neuronParameterB()
func (m MPSMatrixBatchNormalization) NeuronParameterB() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterB"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/neuronParameterC()
func (m MPSMatrixBatchNormalization) NeuronParameterC() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("neuronParameterC"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/neuronType()
func (m MPSMatrixBatchNormalization) NeuronType() MPSCNNNeuronType {
	rv := objc.Send[MPSCNNNeuronType](m.ID, objc.Sel("neuronType"))
	return MPSCNNNeuronType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/setNeuronType(_:parameterA:parameterB:parameterC:)
func (m MPSMatrixBatchNormalization) SetNeuronTypeParameterAParameterBParameterC(neuronType MPSCNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Send[objc.ID](m.ID, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/computeStatistics
func (m MPSMatrixBatchNormalization) ComputeStatistics() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("computeStatistics"))
	return rv
}
func (m MPSMatrixBatchNormalization) SetComputeStatistics(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setComputeStatistics:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/epsilon
func (m MPSMatrixBatchNormalization) Epsilon() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("epsilon"))
	return rv
}
func (m MPSMatrixBatchNormalization) SetEpsilon(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setEpsilon:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/sourceInputFeatureChannels
func (m MPSMatrixBatchNormalization) SourceInputFeatureChannels() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}
func (m MPSMatrixBatchNormalization) SetSourceInputFeatureChannels(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBatchNormalization/sourceNumberOfFeatureVectors
func (m MPSMatrixBatchNormalization) SourceNumberOfFeatureVectors() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}
func (m MPSMatrixBatchNormalization) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

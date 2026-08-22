// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNOptimizerStochasticGradientDescent] class.
var (
	_MPSNNOptimizerStochasticGradientDescentClass     MPSNNOptimizerStochasticGradientDescentClass
	_MPSNNOptimizerStochasticGradientDescentClassOnce sync.Once
)

func getMPSNNOptimizerStochasticGradientDescentClass() MPSNNOptimizerStochasticGradientDescentClass {
	_MPSNNOptimizerStochasticGradientDescentClassOnce.Do(func() {
		_MPSNNOptimizerStochasticGradientDescentClass = MPSNNOptimizerStochasticGradientDescentClass{class: objc.GetClass("MPSNNOptimizerStochasticGradientDescent")}
	})
	return _MPSNNOptimizerStochasticGradientDescentClass
}

// GetMPSNNOptimizerStochasticGradientDescentClass returns the class object for MPSNNOptimizerStochasticGradientDescent.
func GetMPSNNOptimizerStochasticGradientDescentClass() MPSNNOptimizerStochasticGradientDescentClass {
	return getMPSNNOptimizerStochasticGradientDescentClass()
}

type MPSNNOptimizerStochasticGradientDescentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNOptimizerStochasticGradientDescentClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNOptimizerStochasticGradientDescentClass) Alloc() MPSNNOptimizerStochasticGradientDescent {
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An optimization layer that performs a gradient descent with an optional
// momentum update.
//
// # Initializers
//
//   - [MPSNNOptimizerStochasticGradientDescent.InitWithDeviceLearningRate]
//   - [MPSNNOptimizerStochasticGradientDescent.InitWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor]
//   - [MPSNNOptimizerStochasticGradientDescent.InitWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor]
//
// # Instance Properties
//
//   - [MPSNNOptimizerStochasticGradientDescent.MomentumScale]
//   - [MPSNNOptimizerStochasticGradientDescent.UseNesterovMomentum]
//   - [MPSNNOptimizerStochasticGradientDescent.UseNestrovMomentum]
//
// # Instance Methods
//
//   - [MPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState]
//   - [MPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState]
//   - [MPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState]
//   - [MPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix]
//   - [MPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent
type MPSNNOptimizerStochasticGradientDescent struct {
	MPSNNOptimizer
}

// MPSNNOptimizerStochasticGradientDescentFromID constructs a [MPSNNOptimizerStochasticGradientDescent] from an objc.ID.
//
// An optimization layer that performs a gradient descent with an optional
// momentum update.
func MPSNNOptimizerStochasticGradientDescentFromID(id objc.ID) MPSNNOptimizerStochasticGradientDescent {
	return MPSNNOptimizerStochasticGradientDescent{MPSNNOptimizer: MPSNNOptimizerFromID(id)}
}

// NOTE: MPSNNOptimizerStochasticGradientDescent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNOptimizerStochasticGradientDescent] class.
//
// # Initializers
//
//   - [IMPSNNOptimizerStochasticGradientDescent.InitWithDeviceLearningRate]
//   - [IMPSNNOptimizerStochasticGradientDescent.InitWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor]
//   - [IMPSNNOptimizerStochasticGradientDescent.InitWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor]
//
// # Instance Properties
//
//   - [IMPSNNOptimizerStochasticGradientDescent.MomentumScale]
//   - [IMPSNNOptimizerStochasticGradientDescent.UseNesterovMomentum]
//   - [IMPSNNOptimizerStochasticGradientDescent.UseNestrovMomentum]
//
// # Instance Methods
//
//   - [IMPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState]
//   - [IMPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState]
//   - [IMPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState]
//   - [IMPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix]
//   - [IMPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent
type IMPSNNOptimizerStochasticGradientDescent interface {
	IMPSNNOptimizer

	// Topic: Initializers

	InitWithDeviceLearningRate(device metal.MTLDevice, learningRate float32) MPSNNOptimizerStochasticGradientDescent
	InitWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor(device metal.MTLDevice, momentumScale float32, useNesterovMomentum bool, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerStochasticGradientDescent
	InitWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor(device metal.MTLDevice, momentumScale float32, useNestrovMomentum bool, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerStochasticGradientDescent

	// Topic: Instance Properties

	MomentumScale() float32
	UseNesterovMomentum() bool
	UseNestrovMomentum() bool

	// Topic: Instance Methods

	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationGradientState IMPSCNNBatchNormalizationState, batchNormalizationSourceState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState(commandBuffer metal.MTLCommandBuffer, convolutionGradientState IMPSCNNConvolutionGradientState, convolutionSourceState IMPSCNNConvolutionWeightsAndBiasesState, inputMomentumVectors []MPSVector, resultState IMPSCNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBuffer metal.MTLCommandBuffer, inputGradientMatrix IMPSMatrix, inputValuesMatrix IMPSMatrix, inputMomentumMatrix IMPSMatrix, resultValuesMatrix IMPSMatrix)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector(commandBuffer metal.MTLCommandBuffer, inputGradientVector IMPSVector, inputValuesVector IMPSVector, inputMomentumVector IMPSVector, resultValuesVector IMPSVector)
}

// Init initializes the instance.
func (o MPSNNOptimizerStochasticGradientDescent) Init() MPSNNOptimizerStochasticGradientDescent {
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o MPSNNOptimizerStochasticGradientDescent) Autorelease() MPSNNOptimizerStochasticGradientDescent {
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNOptimizerStochasticGradientDescent creates a new MPSNNOptimizerStochasticGradientDescent instance.
func NewMPSNNOptimizerStochasticGradientDescent() MPSNNOptimizerStochasticGradientDescent {
	class := getMPSNNOptimizerStochasticGradientDescentClass()
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewOptimizerStochasticGradientDescentWithCoder(aDecoder foundation.INSCoder) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewOptimizerStochasticGradientDescentWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewOptimizerStochasticGradientDescentWithDevice(device metal.MTLDevice) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/init(device:learningRate:)
func NewOptimizerStochasticGradientDescentWithDeviceLearningRate(device metal.MTLDevice, learningRate float32) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:learningRate:"), device, learningRate)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/init(device:momentumScale:useNesterovMomentum:optimizerDescriptor:)
func NewOptimizerStochasticGradientDescentWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor(device metal.MTLDevice, momentumScale float32, useNesterovMomentum bool, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:momentumScale:useNesterovMomentum:optimizerDescriptor:"), device, momentumScale, useNesterovMomentum, optimizerDescriptor)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/init(device:momentumScale:useNestrovMomentum:optimizerDescriptor:)
func NewOptimizerStochasticGradientDescentWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor(device metal.MTLDevice, momentumScale float32, useNestrovMomentum bool, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:momentumScale:useNestrovMomentum:optimizerDescriptor:"), device, momentumScale, useNestrovMomentum, optimizerDescriptor)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/init(device:learningRate:)
func (o MPSNNOptimizerStochasticGradientDescent) InitWithDeviceLearningRate(device metal.MTLDevice, learningRate float32) MPSNNOptimizerStochasticGradientDescent {
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](o.ID, objc.Sel("initWithDevice:learningRate:"), device, learningRate)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/init(device:momentumScale:useNesterovMomentum:optimizerDescriptor:)
func (o MPSNNOptimizerStochasticGradientDescent) InitWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor(device metal.MTLDevice, momentumScale float32, useNesterovMomentum bool, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerStochasticGradientDescent {
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](o.ID, objc.Sel("initWithDevice:momentumScale:useNesterovMomentum:optimizerDescriptor:"), device, momentumScale, useNesterovMomentum, optimizerDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/init(device:momentumScale:useNestrovMomentum:optimizerDescriptor:)
func (o MPSNNOptimizerStochasticGradientDescent) InitWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor(device metal.MTLDevice, momentumScale float32, useNestrovMomentum bool, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerStochasticGradientDescent {
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](o.ID, objc.Sel("initWithDevice:momentumScale:useNestrovMomentum:optimizerDescriptor:"), device, momentumScale, useNestrovMomentum, optimizerDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/encode(commandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:resultState:)
func (o MPSNNOptimizerStochasticGradientDescent) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationGradientState IMPSCNNBatchNormalizationState, batchNormalizationSourceState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:resultState:"), commandBuffer, batchNormalizationGradientState, batchNormalizationSourceState, objectivec.IObjectSliceToNSArray(inputMomentumVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/encode(commandBuffer:batchNormalizationState:inputMomentumVectors:resultState:)
func (o MPSNNOptimizerStochasticGradientDescent) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:resultState:"), commandBuffer, batchNormalizationState, objectivec.IObjectSliceToNSArray(inputMomentumVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/encode(commandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:resultState:)
func (o MPSNNOptimizerStochasticGradientDescent) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState(commandBuffer metal.MTLCommandBuffer, convolutionGradientState IMPSCNNConvolutionGradientState, convolutionSourceState IMPSCNNConvolutionWeightsAndBiasesState, inputMomentumVectors []MPSVector, resultState IMPSCNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:resultState:"), commandBuffer, convolutionGradientState, convolutionSourceState, objectivec.IObjectSliceToNSArray(inputMomentumVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/encode(commandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:resultValuesMatrix:)
func (o MPSNNOptimizerStochasticGradientDescent) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBuffer metal.MTLCommandBuffer, inputGradientMatrix IMPSMatrix, inputValuesMatrix IMPSMatrix, inputMomentumMatrix IMPSMatrix, resultValuesMatrix IMPSMatrix) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:resultValuesMatrix:"), commandBuffer, inputGradientMatrix, inputValuesMatrix, inputMomentumMatrix, resultValuesMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/encode(commandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:resultValuesVector:)
func (o MPSNNOptimizerStochasticGradientDescent) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector(commandBuffer metal.MTLCommandBuffer, inputGradientVector IMPSVector, inputValuesVector IMPSVector, inputMomentumVector IMPSVector, resultValuesVector IMPSVector) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:resultValuesVector:"), commandBuffer, inputGradientVector, inputValuesVector, inputMomentumVector, resultValuesVector)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/momentumScale
func (o MPSNNOptimizerStochasticGradientDescent) MomentumScale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("momentumScale"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/useNesterovMomentum
func (o MPSNNOptimizerStochasticGradientDescent) UseNesterovMomentum() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("useNesterovMomentum"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerStochasticGradientDescent/useNestrovMomentum
func (o MPSNNOptimizerStochasticGradientDescent) UseNestrovMomentum() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("useNestrovMomentum"))
	return rv
}

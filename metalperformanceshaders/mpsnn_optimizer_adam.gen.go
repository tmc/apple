// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNOptimizerAdam] class.
var (
	_MPSNNOptimizerAdamClass     MPSNNOptimizerAdamClass
	_MPSNNOptimizerAdamClassOnce sync.Once
)

func getMPSNNOptimizerAdamClass() MPSNNOptimizerAdamClass {
	_MPSNNOptimizerAdamClassOnce.Do(func() {
		_MPSNNOptimizerAdamClass = MPSNNOptimizerAdamClass{class: objc.GetClass("MPSNNOptimizerAdam")}
	})
	return _MPSNNOptimizerAdamClass
}

// GetMPSNNOptimizerAdamClass returns the class object for MPSNNOptimizerAdam.
func GetMPSNNOptimizerAdamClass() MPSNNOptimizerAdamClass {
	return getMPSNNOptimizerAdamClass()
}

type MPSNNOptimizerAdamClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNOptimizerAdamClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNOptimizerAdamClass) Alloc() MPSNNOptimizerAdam {
	rv := objc.Send[MPSNNOptimizerAdam](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An optimization layer that performs an Adam pdate.
//
// # Initializers
//
//   - [MPSNNOptimizerAdam.InitWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor]
//   - [MPSNNOptimizerAdam.InitWithDeviceLearningRate]
//
// # Instance Properties
//
//   - [MPSNNOptimizerAdam.Beta1]
//   - [MPSNNOptimizerAdam.Beta2]
//   - [MPSNNOptimizerAdam.Epsilon]
//   - [MPSNNOptimizerAdam.TimeStep]
//   - [MPSNNOptimizerAdam.SetTimeStep]
//
// # Instance Methods
//
//   - [MPSNNOptimizerAdam.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState]
//   - [MPSNNOptimizerAdam.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState]
//   - [MPSNNOptimizerAdam.EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState]
//   - [MPSNNOptimizerAdam.EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState]
//   - [MPSNNOptimizerAdam.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState]
//   - [MPSNNOptimizerAdam.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState]
//   - [MPSNNOptimizerAdam.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix]
//   - [MPSNNOptimizerAdam.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix]
//   - [MPSNNOptimizerAdam.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector]
//   - [MPSNNOptimizerAdam.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam
type MPSNNOptimizerAdam struct {
	MPSNNOptimizer
}

// MPSNNOptimizerAdamFromID constructs a [MPSNNOptimizerAdam] from an objc.ID.
//
// An optimization layer that performs an Adam pdate.
func MPSNNOptimizerAdamFromID(id objc.ID) MPSNNOptimizerAdam {
	return MPSNNOptimizerAdam{MPSNNOptimizer: MPSNNOptimizerFromID(id)}
}

// NOTE: MPSNNOptimizerAdam adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNOptimizerAdam] class.
//
// # Initializers
//
//   - [IMPSNNOptimizerAdam.InitWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor]
//   - [IMPSNNOptimizerAdam.InitWithDeviceLearningRate]
//
// # Instance Properties
//
//   - [IMPSNNOptimizerAdam.Beta1]
//   - [IMPSNNOptimizerAdam.Beta2]
//   - [IMPSNNOptimizerAdam.Epsilon]
//   - [IMPSNNOptimizerAdam.TimeStep]
//   - [IMPSNNOptimizerAdam.SetTimeStep]
//
// # Instance Methods
//
//   - [IMPSNNOptimizerAdam.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState]
//   - [IMPSNNOptimizerAdam.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState]
//   - [IMPSNNOptimizerAdam.EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState]
//   - [IMPSNNOptimizerAdam.EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState]
//   - [IMPSNNOptimizerAdam.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState]
//   - [IMPSNNOptimizerAdam.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState]
//   - [IMPSNNOptimizerAdam.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix]
//   - [IMPSNNOptimizerAdam.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix]
//   - [IMPSNNOptimizerAdam.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector]
//   - [IMPSNNOptimizerAdam.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam
type IMPSNNOptimizerAdam interface {
	IMPSNNOptimizer

	// Topic: Initializers

	InitWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor(device metal.MTLDevice, beta1 float64, beta2 float64, epsilon float32, timeStep uint, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerAdam
	InitWithDeviceLearningRate(device metal.MTLDevice, learningRate float32) MPSNNOptimizerAdam

	// Topic: Instance Properties

	Beta1() float64
	Beta2() float64
	Epsilon() float32
	TimeStep() uint
	SetTimeStep(value uint)

	// Topic: Instance Methods

	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationGradientState IMPSCNNBatchNormalizationState, batchNormalizationSourceState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, maximumVelocityVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationGradientState IMPSCNNBatchNormalizationState, batchNormalizationSourceState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, maximumVelocityVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, convolutionGradientState IMPSCNNConvolutionGradientState, convolutionSourceState IMPSCNNConvolutionWeightsAndBiasesState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, maximumVelocityVectors []MPSVector, resultState IMPSCNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, convolutionGradientState IMPSCNNConvolutionGradientState, convolutionSourceState IMPSCNNConvolutionWeightsAndBiasesState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, resultState IMPSCNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix(commandBuffer metal.MTLCommandBuffer, inputGradientMatrix IMPSMatrix, inputValuesMatrix IMPSMatrix, inputMomentumMatrix IMPSMatrix, inputVelocityMatrix IMPSMatrix, maximumVelocityMatrix IMPSMatrix, resultValuesMatrix IMPSMatrix)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBuffer metal.MTLCommandBuffer, inputGradientMatrix IMPSMatrix, inputValuesMatrix IMPSMatrix, inputMomentumMatrix IMPSMatrix, inputVelocityMatrix IMPSMatrix, resultValuesMatrix IMPSMatrix)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector(commandBuffer metal.MTLCommandBuffer, inputGradientVector IMPSVector, inputValuesVector IMPSVector, inputMomentumVector IMPSVector, inputVelocityVector IMPSVector, maximumVelocityVector IMPSVector, resultValuesVector IMPSVector)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector(commandBuffer metal.MTLCommandBuffer, inputGradientVector IMPSVector, inputValuesVector IMPSVector, inputMomentumVector IMPSVector, inputVelocityVector IMPSVector, resultValuesVector IMPSVector)
}

// Init initializes the instance.
func (o MPSNNOptimizerAdam) Init() MPSNNOptimizerAdam {
	rv := objc.Send[MPSNNOptimizerAdam](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o MPSNNOptimizerAdam) Autorelease() MPSNNOptimizerAdam {
	rv := objc.Send[MPSNNOptimizerAdam](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNOptimizerAdam creates a new MPSNNOptimizerAdam instance.
func NewMPSNNOptimizerAdam() MPSNNOptimizerAdam {
	class := getMPSNNOptimizerAdamClass()
	rv := objc.Send[MPSNNOptimizerAdam](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewOptimizerAdamWithCoder(aDecoder foundation.INSCoder) MPSNNOptimizerAdam {
	instance := getMPSNNOptimizerAdamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNOptimizerAdamFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewOptimizerAdamWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNOptimizerAdam {
	instance := getMPSNNOptimizerAdamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNOptimizerAdamFromID(rv)
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
func NewOptimizerAdamWithDevice(device metal.MTLDevice) MPSNNOptimizerAdam {
	instance := getMPSNNOptimizerAdamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNOptimizerAdamFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/init(device:beta1:beta2:epsilon:timeStep:optimizerDescriptor:)
func NewOptimizerAdamWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor(device metal.MTLDevice, beta1 float64, beta2 float64, epsilon float32, timeStep uint, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerAdam {
	instance := getMPSNNOptimizerAdamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:beta1:beta2:epsilon:timeStep:optimizerDescriptor:"), device, beta1, beta2, epsilon, timeStep, optimizerDescriptor)
	return MPSNNOptimizerAdamFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/init(device:learningRate:)
func NewOptimizerAdamWithDeviceLearningRate(device metal.MTLDevice, learningRate float32) MPSNNOptimizerAdam {
	instance := getMPSNNOptimizerAdamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:learningRate:"), device, learningRate)
	return MPSNNOptimizerAdamFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/init(device:beta1:beta2:epsilon:timeStep:optimizerDescriptor:)
func (o MPSNNOptimizerAdam) InitWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor(device metal.MTLDevice, beta1 float64, beta2 float64, epsilon float32, timeStep uint, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerAdam {
	rv := objc.Send[MPSNNOptimizerAdam](o.ID, objc.Sel("initWithDevice:beta1:beta2:epsilon:timeStep:optimizerDescriptor:"), device, beta1, beta2, epsilon, timeStep, optimizerDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/init(device:learningRate:)
func (o MPSNNOptimizerAdam) InitWithDeviceLearningRate(device metal.MTLDevice, learningRate float32) MPSNNOptimizerAdam {
	rv := objc.Send[MPSNNOptimizerAdam](o.ID, objc.Sel("initWithDevice:learningRate:"), device, learningRate)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/encode(commandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:)
func (o MPSNNOptimizerAdam) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationGradientState IMPSCNNBatchNormalizationState, batchNormalizationSourceState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, maximumVelocityVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), commandBuffer, batchNormalizationGradientState, batchNormalizationSourceState, objectivec.IObjectSliceToNSArray(inputMomentumVectors), objectivec.IObjectSliceToNSArray(inputVelocityVectors), objectivec.IObjectSliceToNSArray(maximumVelocityVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/encode(commandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:resultState:)
func (o MPSNNOptimizerAdam) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationGradientState IMPSCNNBatchNormalizationState, batchNormalizationSourceState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:resultState:"), commandBuffer, batchNormalizationGradientState, batchNormalizationSourceState, objectivec.IObjectSliceToNSArray(inputMomentumVectors), objectivec.IObjectSliceToNSArray(inputVelocityVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/encode(commandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:)
func (o MPSNNOptimizerAdam) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, maximumVelocityVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), commandBuffer, batchNormalizationState, objectivec.IObjectSliceToNSArray(inputMomentumVectors), objectivec.IObjectSliceToNSArray(inputVelocityVectors), objectivec.IObjectSliceToNSArray(maximumVelocityVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/encode(commandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:resultState:)
func (o MPSNNOptimizerAdam) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationState IMPSCNNBatchNormalizationState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:resultState:"), commandBuffer, batchNormalizationState, objectivec.IObjectSliceToNSArray(inputMomentumVectors), objectivec.IObjectSliceToNSArray(inputVelocityVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/encode(commandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:)
func (o MPSNNOptimizerAdam) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, convolutionGradientState IMPSCNNConvolutionGradientState, convolutionSourceState IMPSCNNConvolutionWeightsAndBiasesState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, maximumVelocityVectors []MPSVector, resultState IMPSCNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), commandBuffer, convolutionGradientState, convolutionSourceState, objectivec.IObjectSliceToNSArray(inputMomentumVectors), objectivec.IObjectSliceToNSArray(inputVelocityVectors), objectivec.IObjectSliceToNSArray(maximumVelocityVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/encode(commandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:resultState:)
func (o MPSNNOptimizerAdam) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState(commandBuffer metal.MTLCommandBuffer, convolutionGradientState IMPSCNNConvolutionGradientState, convolutionSourceState IMPSCNNConvolutionWeightsAndBiasesState, inputMomentumVectors []MPSVector, inputVelocityVectors []MPSVector, resultState IMPSCNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:resultState:"), commandBuffer, convolutionGradientState, convolutionSourceState, objectivec.IObjectSliceToNSArray(inputMomentumVectors), objectivec.IObjectSliceToNSArray(inputVelocityVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/encode(commandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:maximumVelocityMatrix:resultValuesMatrix:)
func (o MPSNNOptimizerAdam) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix(commandBuffer metal.MTLCommandBuffer, inputGradientMatrix IMPSMatrix, inputValuesMatrix IMPSMatrix, inputMomentumMatrix IMPSMatrix, inputVelocityMatrix IMPSMatrix, maximumVelocityMatrix IMPSMatrix, resultValuesMatrix IMPSMatrix) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:maximumVelocityMatrix:resultValuesMatrix:"), commandBuffer, inputGradientMatrix, inputValuesMatrix, inputMomentumMatrix, inputVelocityMatrix, maximumVelocityMatrix, resultValuesMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/encode(commandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:resultValuesMatrix:)
func (o MPSNNOptimizerAdam) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(commandBuffer metal.MTLCommandBuffer, inputGradientMatrix IMPSMatrix, inputValuesMatrix IMPSMatrix, inputMomentumMatrix IMPSMatrix, inputVelocityMatrix IMPSMatrix, resultValuesMatrix IMPSMatrix) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:resultValuesMatrix:"), commandBuffer, inputGradientMatrix, inputValuesMatrix, inputMomentumMatrix, inputVelocityMatrix, resultValuesMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/encode(commandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:maximumVelocityVector:resultValuesVector:)
func (o MPSNNOptimizerAdam) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector(commandBuffer metal.MTLCommandBuffer, inputGradientVector IMPSVector, inputValuesVector IMPSVector, inputMomentumVector IMPSVector, inputVelocityVector IMPSVector, maximumVelocityVector IMPSVector, resultValuesVector IMPSVector) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:maximumVelocityVector:resultValuesVector:"), commandBuffer, inputGradientVector, inputValuesVector, inputMomentumVector, inputVelocityVector, maximumVelocityVector, resultValuesVector)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/encode(commandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:resultValuesVector:)
func (o MPSNNOptimizerAdam) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector(commandBuffer metal.MTLCommandBuffer, inputGradientVector IMPSVector, inputValuesVector IMPSVector, inputMomentumVector IMPSVector, inputVelocityVector IMPSVector, resultValuesVector IMPSVector) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:resultValuesVector:"), commandBuffer, inputGradientVector, inputValuesVector, inputMomentumVector, inputVelocityVector, resultValuesVector)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/beta1
func (o MPSNNOptimizerAdam) Beta1() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("beta1"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/beta2
func (o MPSNNOptimizerAdam) Beta2() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("beta2"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/epsilon
func (o MPSNNOptimizerAdam) Epsilon() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("epsilon"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerAdam/timeStep
func (o MPSNNOptimizerAdam) TimeStep() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("timeStep"))
	return rv
}
func (o MPSNNOptimizerAdam) SetTimeStep(value uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTimeStep:"), value)
}

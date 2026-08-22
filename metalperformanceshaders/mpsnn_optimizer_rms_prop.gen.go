// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNOptimizerRMSProp] class.
var (
	_MPSNNOptimizerRMSPropClass     MPSNNOptimizerRMSPropClass
	_MPSNNOptimizerRMSPropClassOnce sync.Once
)

func getMPSNNOptimizerRMSPropClass() MPSNNOptimizerRMSPropClass {
	_MPSNNOptimizerRMSPropClassOnce.Do(func() {
		_MPSNNOptimizerRMSPropClass = MPSNNOptimizerRMSPropClass{class: objc.GetClass("MPSNNOptimizerRMSProp")}
	})
	return _MPSNNOptimizerRMSPropClass
}

// GetMPSNNOptimizerRMSPropClass returns the class object for MPSNNOptimizerRMSProp.
func GetMPSNNOptimizerRMSPropClass() MPSNNOptimizerRMSPropClass {
	return getMPSNNOptimizerRMSPropClass()
}

type MPSNNOptimizerRMSPropClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNOptimizerRMSPropClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNOptimizerRMSPropClass) Alloc() MPSNNOptimizerRMSProp {
	rv := objc.Send[MPSNNOptimizerRMSProp](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An optimization layer that performs a root mean square propagation update.
//
// # Initializers
//
//   - [MPSNNOptimizerRMSProp.InitWithDeviceDecayEpsilonOptimizerDescriptor]
//   - [MPSNNOptimizerRMSProp.InitWithDeviceLearningRate]
//
// # Instance Properties
//
//   - [MPSNNOptimizerRMSProp.Decay]
//   - [MPSNNOptimizerRMSProp.Epsilon]
//
// # Instance Methods
//
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp
type MPSNNOptimizerRMSProp struct {
	MPSNNOptimizer
}

// MPSNNOptimizerRMSPropFromID constructs a [MPSNNOptimizerRMSProp] from an objc.ID.
//
// An optimization layer that performs a root mean square propagation update.
func MPSNNOptimizerRMSPropFromID(id objc.ID) MPSNNOptimizerRMSProp {
	return MPSNNOptimizerRMSProp{MPSNNOptimizer: MPSNNOptimizerFromID(id)}
}

// NOTE: MPSNNOptimizerRMSProp adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNOptimizerRMSProp] class.
//
// # Initializers
//
//   - [IMPSNNOptimizerRMSProp.InitWithDeviceDecayEpsilonOptimizerDescriptor]
//   - [IMPSNNOptimizerRMSProp.InitWithDeviceLearningRate]
//
// # Instance Properties
//
//   - [IMPSNNOptimizerRMSProp.Decay]
//   - [IMPSNNOptimizerRMSProp.Epsilon]
//
// # Instance Methods
//
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp
type IMPSNNOptimizerRMSProp interface {
	IMPSNNOptimizer

	// Topic: Initializers

	InitWithDeviceDecayEpsilonOptimizerDescriptor(device metal.MTLDevice, decay float64, epsilon float32, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerRMSProp
	InitWithDeviceLearningRate(device metal.MTLDevice, learningRate float32) MPSNNOptimizerRMSProp

	// Topic: Instance Properties

	Decay() float64
	Epsilon() float32

	// Topic: Instance Methods

	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationGradientState IMPSCNNBatchNormalizationState, batchNormalizationSourceState IMPSCNNBatchNormalizationState, inputSumOfSquaresVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationState IMPSCNNBatchNormalizationState, inputSumOfSquaresVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState(commandBuffer metal.MTLCommandBuffer, convolutionGradientState IMPSCNNConvolutionGradientState, convolutionSourceState IMPSCNNConvolutionWeightsAndBiasesState, inputSumOfSquaresVectors []MPSVector, resultState IMPSCNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBuffer metal.MTLCommandBuffer, inputGradientMatrix IMPSMatrix, inputValuesMatrix IMPSMatrix, inputSumOfSquaresMatrix IMPSMatrix, resultValuesMatrix IMPSMatrix)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector(commandBuffer metal.MTLCommandBuffer, inputGradientVector IMPSVector, inputValuesVector IMPSVector, inputSumOfSquaresVector IMPSVector, resultValuesVector IMPSVector)
}

// Init initializes the instance.
func (o MPSNNOptimizerRMSProp) Init() MPSNNOptimizerRMSProp {
	rv := objc.Send[MPSNNOptimizerRMSProp](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o MPSNNOptimizerRMSProp) Autorelease() MPSNNOptimizerRMSProp {
	rv := objc.Send[MPSNNOptimizerRMSProp](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNOptimizerRMSProp creates a new MPSNNOptimizerRMSProp instance.
func NewMPSNNOptimizerRMSProp() MPSNNOptimizerRMSProp {
	class := getMPSNNOptimizerRMSPropClass()
	rv := objc.Send[MPSNNOptimizerRMSProp](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewOptimizerRMSPropWithCoder(aDecoder foundation.INSCoder) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNOptimizerRMSPropFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewOptimizerRMSPropWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNOptimizerRMSPropFromID(rv)
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
func NewOptimizerRMSPropWithDevice(device metal.MTLDevice) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNOptimizerRMSPropFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/init(device:decay:epsilon:optimizerDescriptor:)
func NewOptimizerRMSPropWithDeviceDecayEpsilonOptimizerDescriptor(device metal.MTLDevice, decay float64, epsilon float32, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:decay:epsilon:optimizerDescriptor:"), device, decay, epsilon, optimizerDescriptor)
	return MPSNNOptimizerRMSPropFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/init(device:learningRate:)
func NewOptimizerRMSPropWithDeviceLearningRate(device metal.MTLDevice, learningRate float32) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:learningRate:"), device, learningRate)
	return MPSNNOptimizerRMSPropFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/init(device:decay:epsilon:optimizerDescriptor:)
func (o MPSNNOptimizerRMSProp) InitWithDeviceDecayEpsilonOptimizerDescriptor(device metal.MTLDevice, decay float64, epsilon float32, optimizerDescriptor IMPSNNOptimizerDescriptor) MPSNNOptimizerRMSProp {
	rv := objc.Send[MPSNNOptimizerRMSProp](o.ID, objc.Sel("initWithDevice:decay:epsilon:optimizerDescriptor:"), device, decay, epsilon, optimizerDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/init(device:learningRate:)
func (o MPSNNOptimizerRMSProp) InitWithDeviceLearningRate(device metal.MTLDevice, learningRate float32) MPSNNOptimizerRMSProp {
	rv := objc.Send[MPSNNOptimizerRMSProp](o.ID, objc.Sel("initWithDevice:learningRate:"), device, learningRate)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/encode(commandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputSumOfSquaresVectors:resultState:)
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationGradientState IMPSCNNBatchNormalizationState, batchNormalizationSourceState IMPSCNNBatchNormalizationState, inputSumOfSquaresVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputSumOfSquaresVectors:resultState:"), commandBuffer, batchNormalizationGradientState, batchNormalizationSourceState, objectivec.IObjectSliceToNSArray(inputSumOfSquaresVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/encode(commandBuffer:batchNormalizationState:inputSumOfSquaresVectors:resultState:)
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState(commandBuffer metal.MTLCommandBuffer, batchNormalizationState IMPSCNNBatchNormalizationState, inputSumOfSquaresVectors []MPSVector, resultState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputSumOfSquaresVectors:resultState:"), commandBuffer, batchNormalizationState, objectivec.IObjectSliceToNSArray(inputSumOfSquaresVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/encode(commandBuffer:convolutionGradientState:convolutionSourceState:inputSumOfSquaresVectors:resultState:)
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState(commandBuffer metal.MTLCommandBuffer, convolutionGradientState IMPSCNNConvolutionGradientState, convolutionSourceState IMPSCNNConvolutionWeightsAndBiasesState, inputSumOfSquaresVectors []MPSVector, resultState IMPSCNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputSumOfSquaresVectors:resultState:"), commandBuffer, convolutionGradientState, convolutionSourceState, objectivec.IObjectSliceToNSArray(inputSumOfSquaresVectors), resultState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/encode(commandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:resultValuesMatrix:)
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBuffer metal.MTLCommandBuffer, inputGradientMatrix IMPSMatrix, inputValuesMatrix IMPSMatrix, inputSumOfSquaresMatrix IMPSMatrix, resultValuesMatrix IMPSMatrix) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:resultValuesMatrix:"), commandBuffer, inputGradientMatrix, inputValuesMatrix, inputSumOfSquaresMatrix, resultValuesMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/encode(commandBuffer:inputGradientVector:inputValuesVector:inputSumOfSquaresVector:resultValuesVector:)
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector(commandBuffer metal.MTLCommandBuffer, inputGradientVector IMPSVector, inputValuesVector IMPSVector, inputSumOfSquaresVector IMPSVector, resultValuesVector IMPSVector) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputSumOfSquaresVector:resultValuesVector:"), commandBuffer, inputGradientVector, inputValuesVector, inputSumOfSquaresVector, resultValuesVector)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/decay
func (o MPSNNOptimizerRMSProp) Decay() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("decay"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerRMSProp/epsilon
func (o MPSNNOptimizerRMSProp) Epsilon() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("epsilon"))
	return rv
}

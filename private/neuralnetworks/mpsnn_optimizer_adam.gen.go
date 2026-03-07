// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
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

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNOptimizerAdamClass) Alloc() MPSNNOptimizerAdam {
	rv := objc.Send[MPSNNOptimizerAdam](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNOptimizerAdam.Beta1]
//   - [MPSNNOptimizerAdam.Beta2]
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
//   - [MPSNNOptimizerAdam.Epsilon]
//   - [MPSNNOptimizerAdam.TimeStep]
//   - [MPSNNOptimizerAdam.SetTimeStep]
//   - [MPSNNOptimizerAdam.InitWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor]
//   - [MPSNNOptimizerAdam.InitWithDeviceLearningRate]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam
type MPSNNOptimizerAdam struct {
	MPSNNOptimizer
}

// MPSNNOptimizerAdamFromID constructs a [MPSNNOptimizerAdam] from an objc.ID.
func MPSNNOptimizerAdamFromID(id objc.ID) MPSNNOptimizerAdam {
	return MPSNNOptimizerAdam{MPSNNOptimizer: MPSNNOptimizerFromID(id)}
}
// Ensure MPSNNOptimizerAdam implements IMPSNNOptimizerAdam.
var _ IMPSNNOptimizerAdam = MPSNNOptimizerAdam{}





// An interface definition for the [MPSNNOptimizerAdam] class.
//
// # Methods
//
//   - [IMPSNNOptimizerAdam.Beta1]
//   - [IMPSNNOptimizerAdam.Beta2]
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
//   - [IMPSNNOptimizerAdam.Epsilon]
//   - [IMPSNNOptimizerAdam.TimeStep]
//   - [IMPSNNOptimizerAdam.SetTimeStep]
//   - [IMPSNNOptimizerAdam.InitWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor]
//   - [IMPSNNOptimizerAdam.InitWithDeviceLearningRate]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam
type IMPSNNOptimizerAdam interface {
	IMPSNNOptimizer

	// Topic: Methods

	Beta1() float64
	Beta2() float64
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, vectors3 objectivec.IObject, state3 objectivec.IObject)
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, state3 objectivec.IObject)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, vectors3 objectivec.IObject, state2 objectivec.IObject)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, state2 objectivec.IObject)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, vectors3 objectivec.IObject, state3 objectivec.IObject)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, state3 objectivec.IObject)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix(buffer objectivec.IObject, matrix objectivec.IObject, matrix2 objectivec.IObject, matrix3 objectivec.IObject, matrix4 objectivec.IObject, matrix5 objectivec.IObject, matrix6 objectivec.IObject)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(buffer objectivec.IObject, matrix objectivec.IObject, matrix2 objectivec.IObject, matrix3 objectivec.IObject, matrix4 objectivec.IObject, matrix5 objectivec.IObject)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector(buffer objectivec.IObject, vector objectivec.IObject, vector2 objectivec.IObject, vector3 objectivec.IObject, vector4 objectivec.IObject, vector5 objectivec.IObject, vector6 objectivec.IObject)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector(buffer objectivec.IObject, vector objectivec.IObject, vector2 objectivec.IObject, vector3 objectivec.IObject, vector4 objectivec.IObject, vector5 objectivec.IObject)
	Epsilon() float32
	TimeStep() uint64
	SetTimeStep(value uint64)
	InitWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor(device objectivec.IObject, beta1 float64, beta2 float64, epsilon float32, step uint64, descriptor objectivec.IObject) MPSNNOptimizerAdam
	InitWithDeviceLearningRate(device objectivec.IObject, rate float32) MPSNNOptimizerAdam
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






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/initWithCoder:device:
func NewOptimizerAdamWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNOptimizerAdam {
	instance := getMPSNNOptimizerAdamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNOptimizerAdamFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/initWithDevice:
func NewOptimizerAdamWithDevice(device objectivec.IObject) MPSNNOptimizerAdam {
	instance := getMPSNNOptimizerAdamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNOptimizerAdamFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/initWithDevice:beta1:beta2:epsilon:timeStep:optimizerDescriptor:
func NewOptimizerAdamWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor(device objectivec.IObject, beta1 float64, beta2 float64, epsilon float32, step uint64, descriptor objectivec.IObject) MPSNNOptimizerAdam {
	instance := getMPSNNOptimizerAdamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:beta1:beta2:epsilon:timeStep:optimizerDescriptor:"), device, beta1, beta2, epsilon, step, descriptor)
	return MPSNNOptimizerAdamFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/initWithDevice:learningRate:
func NewOptimizerAdamWithDeviceLearningRate(device objectivec.IObject, rate float32) MPSNNOptimizerAdam {
	instance := getMPSNNOptimizerAdamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:learningRate:"), device, rate)
	return MPSNNOptimizerAdamFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/initWithDevice:optimizerDescriptor:
func NewOptimizerAdamWithDeviceOptimizerDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNOptimizerAdam {
	instance := getMPSNNOptimizerAdamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:optimizerDescriptor:"), device, descriptor)
	return MPSNNOptimizerAdamFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:
func (o MPSNNOptimizerAdam) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, vectors3 objectivec.IObject, state3 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), buffer, state, state2, vectors, vectors2, vectors3, state3)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:resultState:
func (o MPSNNOptimizerAdam) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsInputVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, state3 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:inputVelocityVectors:resultState:"), buffer, state, state2, vectors, vectors2, state3)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:
func (o MPSNNOptimizerAdam) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, vectors3 objectivec.IObject, state2 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), buffer, state, vectors, vectors2, vectors3, state2)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:resultState:
func (o MPSNNOptimizerAdam) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsInputVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, state2 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:inputVelocityVectors:resultState:"), buffer, state, vectors, vectors2, state2)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:
func (o MPSNNOptimizerAdam) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsMaximumVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, vectors3 objectivec.IObject, state3 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:maximumVelocityVectors:resultState:"), buffer, state, state2, vectors, vectors2, vectors3, state3)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:resultState:
func (o MPSNNOptimizerAdam) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsInputVelocityVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, vectors2 objectivec.IObject, state3 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:inputVelocityVectors:resultState:"), buffer, state, state2, vectors, vectors2, state3)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:maximumVelocityMatrix:resultValuesMatrix:
func (o MPSNNOptimizerAdam) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixMaximumVelocityMatrixResultValuesMatrix(buffer objectivec.IObject, matrix objectivec.IObject, matrix2 objectivec.IObject, matrix3 objectivec.IObject, matrix4 objectivec.IObject, matrix5 objectivec.IObject, matrix6 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:maximumVelocityMatrix:resultValuesMatrix:"), buffer, matrix, matrix2, matrix3, matrix4, matrix5, matrix6)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:resultValuesMatrix:
func (o MPSNNOptimizerAdam) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixInputVelocityMatrixResultValuesMatrix(buffer objectivec.IObject, matrix objectivec.IObject, matrix2 objectivec.IObject, matrix3 objectivec.IObject, matrix4 objectivec.IObject, matrix5 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:inputVelocityMatrix:resultValuesMatrix:"), buffer, matrix, matrix2, matrix3, matrix4, matrix5)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:maximumVelocityVector:resultValuesVector:
func (o MPSNNOptimizerAdam) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorMaximumVelocityVectorResultValuesVector(buffer objectivec.IObject, vector objectivec.IObject, vector2 objectivec.IObject, vector3 objectivec.IObject, vector4 objectivec.IObject, vector5 objectivec.IObject, vector6 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:maximumVelocityVector:resultValuesVector:"), buffer, vector, vector2, vector3, vector4, vector5, vector6)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:resultValuesVector:
func (o MPSNNOptimizerAdam) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorInputVelocityVectorResultValuesVector(buffer objectivec.IObject, vector objectivec.IObject, vector2 objectivec.IObject, vector3 objectivec.IObject, vector4 objectivec.IObject, vector5 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:inputVelocityVector:resultValuesVector:"), buffer, vector, vector2, vector3, vector4, vector5)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/initWithDevice:beta1:beta2:epsilon:timeStep:optimizerDescriptor:
func (o MPSNNOptimizerAdam) InitWithDeviceBeta1Beta2EpsilonTimeStepOptimizerDescriptor(device objectivec.IObject, beta1 float64, beta2 float64, epsilon float32, step uint64, descriptor objectivec.IObject) MPSNNOptimizerAdam {
	rv := objc.Send[MPSNNOptimizerAdam](o.ID, objc.Sel("initWithDevice:beta1:beta2:epsilon:timeStep:optimizerDescriptor:"), device, beta1, beta2, epsilon, step, descriptor)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/initWithDevice:learningRate:
func (o MPSNNOptimizerAdam) InitWithDeviceLearningRate(device objectivec.IObject, rate float32) MPSNNOptimizerAdam {
	rv := objc.Send[MPSNNOptimizerAdam](o.ID, objc.Sel("initWithDevice:learningRate:"), device, rate)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/libraryInfo:
func (_MPSNNOptimizerAdamClass MPSNNOptimizerAdamClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNOptimizerAdamClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/beta1
func (o MPSNNOptimizerAdam) Beta1() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("beta1"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/beta2
func (o MPSNNOptimizerAdam) Beta2() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("beta2"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/epsilon
func (o MPSNNOptimizerAdam) Epsilon() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("epsilon"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerAdam/timeStep
func (o MPSNNOptimizerAdam) TimeStep() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("timeStep"))
	return rv
}
func (o MPSNNOptimizerAdam) SetTimeStep(value uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setTimeStep:"), value)
}


















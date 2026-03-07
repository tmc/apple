// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
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

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNOptimizerRMSPropClass) Alloc() MPSNNOptimizerRMSProp {
	rv := objc.Send[MPSNNOptimizerRMSProp](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNOptimizerRMSProp.Centered]
//   - [MPSNNOptimizerRMSProp.Decay]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixInputWeightedSumMatrixInputMomentumMatrixResultValuesMatrix]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorInputWeightedSumVectorInputMomentumVectorResultValuesVector]
//   - [MPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector]
//   - [MPSNNOptimizerRMSProp.Epsilon]
//   - [MPSNNOptimizerRMSProp.MomentumScale]
//   - [MPSNNOptimizerRMSProp.InitWithDeviceDecayEpsilonMomentumScaleCenteredOptimizerDescriptor]
//   - [MPSNNOptimizerRMSProp.InitWithDeviceDecayEpsilonOptimizerDescriptor]
//   - [MPSNNOptimizerRMSProp.InitWithDeviceLearningRate]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp
type MPSNNOptimizerRMSProp struct {
	MPSNNOptimizer
}

// MPSNNOptimizerRMSPropFromID constructs a [MPSNNOptimizerRMSProp] from an objc.ID.
func MPSNNOptimizerRMSPropFromID(id objc.ID) MPSNNOptimizerRMSProp {
	return MPSNNOptimizerRMSProp{MPSNNOptimizer: MPSNNOptimizerFromID(id)}
}
// Ensure MPSNNOptimizerRMSProp implements IMPSNNOptimizerRMSProp.
var _ IMPSNNOptimizerRMSProp = MPSNNOptimizerRMSProp{}





// An interface definition for the [MPSNNOptimizerRMSProp] class.
//
// # Methods
//
//   - [IMPSNNOptimizerRMSProp.Centered]
//   - [IMPSNNOptimizerRMSProp.Decay]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixInputWeightedSumMatrixInputMomentumMatrixResultValuesMatrix]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorInputWeightedSumVectorInputMomentumVectorResultValuesVector]
//   - [IMPSNNOptimizerRMSProp.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector]
//   - [IMPSNNOptimizerRMSProp.Epsilon]
//   - [IMPSNNOptimizerRMSProp.MomentumScale]
//   - [IMPSNNOptimizerRMSProp.InitWithDeviceDecayEpsilonMomentumScaleCenteredOptimizerDescriptor]
//   - [IMPSNNOptimizerRMSProp.InitWithDeviceDecayEpsilonOptimizerDescriptor]
//   - [IMPSNNOptimizerRMSProp.InitWithDeviceLearningRate]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp
type IMPSNNOptimizerRMSProp interface {
	IMPSNNOptimizer

	// Topic: Methods

	Centered() bool
	Decay() float64
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, state3 objectivec.IObject)
	EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, vectors objectivec.IObject, state2 objectivec.IObject)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, state3 objectivec.IObject)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixInputWeightedSumMatrixInputMomentumMatrixResultValuesMatrix(buffer objectivec.IObject, matrix objectivec.IObject, matrix2 objectivec.IObject, matrix3 objectivec.IObject, matrix4 objectivec.IObject, matrix5 objectivec.IObject, matrix6 objectivec.IObject)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(buffer objectivec.IObject, matrix objectivec.IObject, matrix2 objectivec.IObject, matrix3 objectivec.IObject, matrix4 objectivec.IObject)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorInputWeightedSumVectorInputMomentumVectorResultValuesVector(buffer objectivec.IObject, vector objectivec.IObject, vector2 objectivec.IObject, vector3 objectivec.IObject, vector4 objectivec.IObject, vector5 objectivec.IObject, vector6 objectivec.IObject)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector(buffer objectivec.IObject, vector objectivec.IObject, vector2 objectivec.IObject, vector3 objectivec.IObject, vector4 objectivec.IObject)
	Epsilon() float32
	MomentumScale() float64
	InitWithDeviceDecayEpsilonMomentumScaleCenteredOptimizerDescriptor(device objectivec.IObject, decay float64, epsilon float32, scale float64, centered bool, descriptor objectivec.IObject) MPSNNOptimizerRMSProp
	InitWithDeviceDecayEpsilonOptimizerDescriptor(device objectivec.IObject, decay float64, epsilon float32, descriptor objectivec.IObject) MPSNNOptimizerRMSProp
	InitWithDeviceLearningRate(device objectivec.IObject, rate float32) MPSNNOptimizerRMSProp
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






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/initWithCoder:device:
func NewOptimizerRMSPropWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNOptimizerRMSPropFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/initWithDevice:
func NewOptimizerRMSPropWithDevice(device objectivec.IObject) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNOptimizerRMSPropFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/initWithDevice:decay:epsilon:momentumScale:centered:optimizerDescriptor:
func NewOptimizerRMSPropWithDeviceDecayEpsilonMomentumScaleCenteredOptimizerDescriptor(device objectivec.IObject, decay float64, epsilon float32, scale float64, centered bool, descriptor objectivec.IObject) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:decay:epsilon:momentumScale:centered:optimizerDescriptor:"), device, decay, epsilon, scale, centered, descriptor)
	return MPSNNOptimizerRMSPropFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/initWithDevice:decay:epsilon:optimizerDescriptor:
func NewOptimizerRMSPropWithDeviceDecayEpsilonOptimizerDescriptor(device objectivec.IObject, decay float64, epsilon float32, descriptor objectivec.IObject) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:decay:epsilon:optimizerDescriptor:"), device, decay, epsilon, descriptor)
	return MPSNNOptimizerRMSPropFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/initWithDevice:learningRate:
func NewOptimizerRMSPropWithDeviceLearningRate(device objectivec.IObject, rate float32) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:learningRate:"), device, rate)
	return MPSNNOptimizerRMSPropFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/initWithDevice:optimizerDescriptor:
func NewOptimizerRMSPropWithDeviceOptimizerDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNOptimizerRMSProp {
	instance := getMPSNNOptimizerRMSPropClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:optimizerDescriptor:"), device, descriptor)
	return MPSNNOptimizerRMSPropFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputSumOfSquaresVectors:resultState:
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputSumOfSquaresVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, state3 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputSumOfSquaresVectors:resultState:"), buffer, state, state2, vectors, state3)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/encodeToCommandBuffer:batchNormalizationState:inputSumOfSquaresVectors:resultState:
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferBatchNormalizationStateInputSumOfSquaresVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, vectors objectivec.IObject, state2 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputSumOfSquaresVectors:resultState:"), buffer, state, vectors, state2)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputSumOfSquaresVectors:resultState:
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputSumOfSquaresVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, state3 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputSumOfSquaresVectors:resultState:"), buffer, state, state2, vectors, state3)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:inputWeightedSumMatrix:inputMomentumMatrix:resultValuesMatrix:
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixInputWeightedSumMatrixInputMomentumMatrixResultValuesMatrix(buffer objectivec.IObject, matrix objectivec.IObject, matrix2 objectivec.IObject, matrix3 objectivec.IObject, matrix4 objectivec.IObject, matrix5 objectivec.IObject, matrix6 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:inputWeightedSumMatrix:inputMomentumMatrix:resultValuesMatrix:"), buffer, matrix, matrix2, matrix3, matrix4, matrix5, matrix6)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:resultValuesMatrix:
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(buffer objectivec.IObject, matrix objectivec.IObject, matrix2 objectivec.IObject, matrix3 objectivec.IObject, matrix4 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:resultValuesMatrix:"), buffer, matrix, matrix2, matrix3, matrix4)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputSumOfSquaresVector:inputWeightedSumVector:inputMomentumVector:resultValuesVector:
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorInputWeightedSumVectorInputMomentumVectorResultValuesVector(buffer objectivec.IObject, vector objectivec.IObject, vector2 objectivec.IObject, vector3 objectivec.IObject, vector4 objectivec.IObject, vector5 objectivec.IObject, vector6 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputSumOfSquaresVector:inputWeightedSumVector:inputMomentumVector:resultValuesVector:"), buffer, vector, vector2, vector3, vector4, vector5, vector6)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputSumOfSquaresVector:resultValuesVector:
func (o MPSNNOptimizerRMSProp) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputSumOfSquaresVectorResultValuesVector(buffer objectivec.IObject, vector objectivec.IObject, vector2 objectivec.IObject, vector3 objectivec.IObject, vector4 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputSumOfSquaresVector:resultValuesVector:"), buffer, vector, vector2, vector3, vector4)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/initWithDevice:decay:epsilon:momentumScale:centered:optimizerDescriptor:
func (o MPSNNOptimizerRMSProp) InitWithDeviceDecayEpsilonMomentumScaleCenteredOptimizerDescriptor(device objectivec.IObject, decay float64, epsilon float32, scale float64, centered bool, descriptor objectivec.IObject) MPSNNOptimizerRMSProp {
	rv := objc.Send[MPSNNOptimizerRMSProp](o.ID, objc.Sel("initWithDevice:decay:epsilon:momentumScale:centered:optimizerDescriptor:"), device, decay, epsilon, scale, centered, descriptor)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/initWithDevice:decay:epsilon:optimizerDescriptor:
func (o MPSNNOptimizerRMSProp) InitWithDeviceDecayEpsilonOptimizerDescriptor(device objectivec.IObject, decay float64, epsilon float32, descriptor objectivec.IObject) MPSNNOptimizerRMSProp {
	rv := objc.Send[MPSNNOptimizerRMSProp](o.ID, objc.Sel("initWithDevice:decay:epsilon:optimizerDescriptor:"), device, decay, epsilon, descriptor)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/initWithDevice:learningRate:
func (o MPSNNOptimizerRMSProp) InitWithDeviceLearningRate(device objectivec.IObject, rate float32) MPSNNOptimizerRMSProp {
	rv := objc.Send[MPSNNOptimizerRMSProp](o.ID, objc.Sel("initWithDevice:learningRate:"), device, rate)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/libraryInfo:
func (_MPSNNOptimizerRMSPropClass MPSNNOptimizerRMSPropClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNOptimizerRMSPropClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/centered
func (o MPSNNOptimizerRMSProp) Centered() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("centered"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/decay
func (o MPSNNOptimizerRMSProp) Decay() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("decay"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/epsilon
func (o MPSNNOptimizerRMSProp) Epsilon() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("epsilon"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerRMSProp/momentumScale
func (o MPSNNOptimizerRMSProp) MomentumScale() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("momentumScale"))
	return rv
}


















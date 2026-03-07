// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
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

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNOptimizerStochasticGradientDescentClass) Alloc() MPSNNOptimizerStochasticGradientDescent {
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState]
//   - [MPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState]
//   - [MPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState]
//   - [MPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix]
//   - [MPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector]
//   - [MPSNNOptimizerStochasticGradientDescent.MomentumScale]
//   - [MPSNNOptimizerStochasticGradientDescent.UseNesterovMomentum]
//   - [MPSNNOptimizerStochasticGradientDescent.UseNestrovMomentum]
//   - [MPSNNOptimizerStochasticGradientDescent.InitWithDeviceLearningRate]
//   - [MPSNNOptimizerStochasticGradientDescent.InitWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor]
//   - [MPSNNOptimizerStochasticGradientDescent.InitWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent
type MPSNNOptimizerStochasticGradientDescent struct {
	MPSNNOptimizer
}

// MPSNNOptimizerStochasticGradientDescentFromID constructs a [MPSNNOptimizerStochasticGradientDescent] from an objc.ID.
func MPSNNOptimizerStochasticGradientDescentFromID(id objc.ID) MPSNNOptimizerStochasticGradientDescent {
	return MPSNNOptimizerStochasticGradientDescent{MPSNNOptimizer: MPSNNOptimizerFromID(id)}
}
// Ensure MPSNNOptimizerStochasticGradientDescent implements IMPSNNOptimizerStochasticGradientDescent.
var _ IMPSNNOptimizerStochasticGradientDescent = MPSNNOptimizerStochasticGradientDescent{}





// An interface definition for the [MPSNNOptimizerStochasticGradientDescent] class.
//
// # Methods
//
//   - [IMPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState]
//   - [IMPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState]
//   - [IMPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState]
//   - [IMPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix]
//   - [IMPSNNOptimizerStochasticGradientDescent.EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector]
//   - [IMPSNNOptimizerStochasticGradientDescent.MomentumScale]
//   - [IMPSNNOptimizerStochasticGradientDescent.UseNesterovMomentum]
//   - [IMPSNNOptimizerStochasticGradientDescent.UseNestrovMomentum]
//   - [IMPSNNOptimizerStochasticGradientDescent.InitWithDeviceLearningRate]
//   - [IMPSNNOptimizerStochasticGradientDescent.InitWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor]
//   - [IMPSNNOptimizerStochasticGradientDescent.InitWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent
type IMPSNNOptimizerStochasticGradientDescent interface {
	IMPSNNOptimizer

	// Topic: Methods

	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, state3 objectivec.IObject)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, vectors objectivec.IObject, state2 objectivec.IObject)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, state3 objectivec.IObject)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(buffer objectivec.IObject, matrix objectivec.IObject, matrix2 objectivec.IObject, matrix3 objectivec.IObject, matrix4 objectivec.IObject)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector(buffer objectivec.IObject, vector objectivec.IObject, vector2 objectivec.IObject, vector3 objectivec.IObject, vector4 objectivec.IObject)
	MomentumScale() float32
	UseNesterovMomentum() bool
	UseNestrovMomentum() bool
	InitWithDeviceLearningRate(device objectivec.IObject, rate float32) MPSNNOptimizerStochasticGradientDescent
	InitWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor(device objectivec.IObject, scale float32, momentum bool, descriptor objectivec.IObject) MPSNNOptimizerStochasticGradientDescent
	InitWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor(device objectivec.IObject, scale float32, momentum bool, descriptor objectivec.IObject) MPSNNOptimizerStochasticGradientDescent
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






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/initWithCoder:device:
func NewOptimizerStochasticGradientDescentWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/initWithDevice:
func NewOptimizerStochasticGradientDescentWithDevice(device objectivec.IObject) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/initWithDevice:learningRate:
func NewOptimizerStochasticGradientDescentWithDeviceLearningRate(device objectivec.IObject, rate float32) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:learningRate:"), device, rate)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/initWithDevice:momentumScale:useNesterovMomentum:optimizerDescriptor:
func NewOptimizerStochasticGradientDescentWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor(device objectivec.IObject, scale float32, momentum bool, descriptor objectivec.IObject) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:momentumScale:useNesterovMomentum:optimizerDescriptor:"), device, scale, momentum, descriptor)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/initWithDevice:momentumScale:useNestrovMomentum:optimizerDescriptor:
func NewOptimizerStochasticGradientDescentWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor(device objectivec.IObject, scale float32, momentum bool, descriptor objectivec.IObject) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:momentumScale:useNestrovMomentum:optimizerDescriptor:"), device, scale, momentum, descriptor)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/initWithDevice:optimizerDescriptor:
func NewOptimizerStochasticGradientDescentWithDeviceOptimizerDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNOptimizerStochasticGradientDescent {
	instance := getMPSNNOptimizerStochasticGradientDescentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:optimizerDescriptor:"), device, descriptor)
	return MPSNNOptimizerStochasticGradientDescentFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:resultState:
func (o MPSNNOptimizerStochasticGradientDescent) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, state3 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:resultState:"), buffer, state, state2, vectors, state3)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:resultState:
func (o MPSNNOptimizerStochasticGradientDescent) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, vectors objectivec.IObject, state2 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:resultState:"), buffer, state, vectors, state2)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:resultState:
func (o MPSNNOptimizerStochasticGradientDescent) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState(buffer objectivec.IObject, state objectivec.IObject, state2 objectivec.IObject, vectors objectivec.IObject, state3 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:resultState:"), buffer, state, state2, vectors, state3)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:resultValuesMatrix:
func (o MPSNNOptimizerStochasticGradientDescent) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(buffer objectivec.IObject, matrix objectivec.IObject, matrix2 objectivec.IObject, matrix3 objectivec.IObject, matrix4 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:resultValuesMatrix:"), buffer, matrix, matrix2, matrix3, matrix4)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:resultValuesVector:
func (o MPSNNOptimizerStochasticGradientDescent) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector(buffer objectivec.IObject, vector objectivec.IObject, vector2 objectivec.IObject, vector3 objectivec.IObject, vector4 objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:resultValuesVector:"), buffer, vector, vector2, vector3, vector4)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/initWithDevice:learningRate:
func (o MPSNNOptimizerStochasticGradientDescent) InitWithDeviceLearningRate(device objectivec.IObject, rate float32) MPSNNOptimizerStochasticGradientDescent {
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](o.ID, objc.Sel("initWithDevice:learningRate:"), device, rate)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/initWithDevice:momentumScale:useNesterovMomentum:optimizerDescriptor:
func (o MPSNNOptimizerStochasticGradientDescent) InitWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor(device objectivec.IObject, scale float32, momentum bool, descriptor objectivec.IObject) MPSNNOptimizerStochasticGradientDescent {
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](o.ID, objc.Sel("initWithDevice:momentumScale:useNesterovMomentum:optimizerDescriptor:"), device, scale, momentum, descriptor)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/initWithDevice:momentumScale:useNestrovMomentum:optimizerDescriptor:
func (o MPSNNOptimizerStochasticGradientDescent) InitWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor(device objectivec.IObject, scale float32, momentum bool, descriptor objectivec.IObject) MPSNNOptimizerStochasticGradientDescent {
	rv := objc.Send[MPSNNOptimizerStochasticGradientDescent](o.ID, objc.Sel("initWithDevice:momentumScale:useNestrovMomentum:optimizerDescriptor:"), device, scale, momentum, descriptor)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/libraryInfo:
func (_MPSNNOptimizerStochasticGradientDescentClass MPSNNOptimizerStochasticGradientDescentClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNOptimizerStochasticGradientDescentClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/momentumScale
func (o MPSNNOptimizerStochasticGradientDescent) MomentumScale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("momentumScale"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/useNesterovMomentum
func (o MPSNNOptimizerStochasticGradientDescent) UseNesterovMomentum() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("useNesterovMomentum"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerStochasticGradientDescent/useNestrovMomentum
func (o MPSNNOptimizerStochasticGradientDescent) UseNestrovMomentum() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("useNestrovMomentum"))
	return rv
}


















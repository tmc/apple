// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralNetworkMLComputeLayer] class.
var (
	_MLNeuralNetworkMLComputeLayerClass     MLNeuralNetworkMLComputeLayerClass
	_MLNeuralNetworkMLComputeLayerClassOnce sync.Once
)

func getMLNeuralNetworkMLComputeLayerClass() MLNeuralNetworkMLComputeLayerClass {
	_MLNeuralNetworkMLComputeLayerClassOnce.Do(func() {
		_MLNeuralNetworkMLComputeLayerClass = MLNeuralNetworkMLComputeLayerClass{class: objc.GetClass("MLNeuralNetworkMLComputeLayer")}
	})
	return _MLNeuralNetworkMLComputeLayerClass
}

// GetMLNeuralNetworkMLComputeLayerClass returns the class object for MLNeuralNetworkMLComputeLayer.
func GetMLNeuralNetworkMLComputeLayerClass() MLNeuralNetworkMLComputeLayerClass {
	return getMLNeuralNetworkMLComputeLayerClass()
}

type MLNeuralNetworkMLComputeLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkMLComputeLayerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkMLComputeLayerClass) Alloc() MLNeuralNetworkMLComputeLayer {
	rv := objc.Send[MLNeuralNetworkMLComputeLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer
type MLNeuralNetworkMLComputeLayer struct {
	objectivec.Object
}

// MLNeuralNetworkMLComputeLayerFromID constructs a [MLNeuralNetworkMLComputeLayer] from an objc.ID.
func MLNeuralNetworkMLComputeLayerFromID(id objc.ID) MLNeuralNetworkMLComputeLayer {
	return MLNeuralNetworkMLComputeLayer{objectivec.Object{ID: id}}
}
// Ensure MLNeuralNetworkMLComputeLayer implements IMLNeuralNetworkMLComputeLayer.
var _ IMLNeuralNetworkMLComputeLayer = MLNeuralNetworkMLComputeLayer{}

// An interface definition for the [MLNeuralNetworkMLComputeLayer] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer
type IMLNeuralNetworkMLComputeLayer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (n MLNeuralNetworkMLComputeLayer) Init() MLNeuralNetworkMLComputeLayer {
	rv := objc.Send[MLNeuralNetworkMLComputeLayer](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralNetworkMLComputeLayer) Autorelease() MLNeuralNetworkMLComputeLayer {
	rv := objc.Send[MLNeuralNetworkMLComputeLayer](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkMLComputeLayer creates a new MLNeuralNetworkMLComputeLayer instance.
func NewMLNeuralNetworkMLComputeLayer() MLNeuralNetworkMLComputeLayer {
	class := getMLNeuralNetworkMLComputeLayerClass()
	rv := objc.Send[MLNeuralNetworkMLComputeLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/adamOptimizerWithLearningRate:beta1:beta2:epsilon:timeStep:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) AdamOptimizerWithLearningRateBeta1Beta2EpsilonTimeStep(rate float64, beta1 float64, beta2 float64, epsilon float64, step uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("adamOptimizerWithLearningRate:beta1:beta2:epsilon:timeStep:"), rate, beta1, beta2, epsilon, step)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/batchNormLayerWithFeatureChannels:varianceEpsilon:gamma:beta:mean:variance:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) BatchNormLayerWithFeatureChannelsVarianceEpsilonGammaBetaMeanVariance(channels uint64, epsilon float32, gamma objectivec.IObject, beta objectivec.IObject, mean objectivec.IObject, variance objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("batchNormLayerWithFeatureChannels:varianceEpsilon:gamma:beta:mean:variance:"), channels, epsilon, gamma, beta, mean, variance)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/categoricalCrossEntropyLossWithNumberOfClasses:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) CategoricalCrossEntropyLossWithNumberOfClasses(classes uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("categoricalCrossEntropyLossWithNumberOfClasses:"), classes)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/concatNDLayerWithAxis:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) ConcatNDLayerWithAxis(axis uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("concatNDLayerWithAxis:"), axis)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/convolutionalLayerWithKernelHeight:kernelWidth:inputChannels:outputChannels:strideInX:strideInY:dilationRateInX:dilationRateInY:nGroups:weight:bias:paddingPolicy:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) ConvolutionalLayerWithKernelHeightKernelWidthInputChannelsOutputChannelsStrideInXStrideInYDilationRateInXDilationRateInYNGroupsWeightBiasPaddingPolicy(height uint64, width uint64, channels uint64, channels2 uint64, x uint64, y uint64, x2 uint64, y2 uint64, groups uint64, weight objectivec.IObject, bias objectivec.IObject, policy int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("convolutionalLayerWithKernelHeight:kernelWidth:inputChannels:outputChannels:strideInX:strideInY:dilationRateInX:dilationRateInY:nGroups:weight:bias:paddingPolicy:"), height, width, channels, channels2, x, y, x2, y2, groups, weight, bias, policy)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/fullyConnectedLayerWithKernelWidth:kernelHeight:inputChannels:outputChannels:weight:bias:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) FullyConnectedLayerWithKernelWidthKernelHeightInputChannelsOutputChannelsWeightBias(width uint64, height uint64, channels uint64, channels2 uint64, weight objectivec.IObject, bias objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("fullyConnectedLayerWithKernelWidth:kernelHeight:inputChannels:outputChannels:weight:bias:"), width, height, channels, channels2, weight, bias)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/leakyRelu
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) LeakyRelu() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("leakyRelu"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/matMulLayerWithTransposesX:transposesY:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) MatMulLayerWithTransposesXTransposesY(x bool, y bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("matMulLayerWithTransposesX:transposesY:"), x, y)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/meanSquaredErrorLoss
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) MeanSquaredErrorLoss() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("meanSquaredErrorLoss"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/poolingLayerWithKernelWidth:kernelHeight:strideInX:strideInY:paddingPolicy:poolingType:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) PoolingLayerWithKernelWidthKernelHeightStrideInXStrideInYPaddingPolicyPoolingType(width uint64, height uint64, x uint64, y uint64, policy int, type_ int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("poolingLayerWithKernelWidth:kernelHeight:strideInX:strideInY:paddingPolicy:poolingType:"), width, height, x, y, policy, type_)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/relu
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) Relu() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("relu"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/reluNLayerWithAlpha:beta:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) ReluNLayerWithAlphaBeta(alpha float32, beta float32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("reluNLayerWithAlpha:beta:"), alpha, beta)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/reshapeLayerWithChannelFirstOrderingAndTargetShape:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) ReshapeLayerWithChannelFirstOrderingAndTargetShape(shape objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("reshapeLayerWithChannelFirstOrderingAndTargetShape:"), shape)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/sgdOptimizerWithLearningRate:momentum:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) SgdOptimizerWithLearningRateMomentum(rate float64, momentum float64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("sgdOptimizerWithLearningRate:momentum:"), rate, momentum)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/sigmoid
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) Sigmoid() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("sigmoid"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/softmax
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) Softmax() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("softmax"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/transposeLayerWithAxes:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) TransposeLayerWithAxes(axes objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("transposeLayerWithAxes:"), axes)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeLayer/uniDirectionalLSTMLayerWithInputSize:outputSize:inputWeightsData:recursionWeightsData:hasBias:biasTermsData:sequenceOutput:inputGateActivation:cellGateActivation:hiddenOutputActivation:
func (_MLNeuralNetworkMLComputeLayerClass MLNeuralNetworkMLComputeLayerClass) UniDirectionalLSTMLayerWithInputSizeOutputSizeInputWeightsDataRecursionWeightsDataHasBiasBiasTermsDataSequenceOutputInputGateActivationCellGateActivationHiddenOutputActivation(size uint64, size2 uint64, data objectivec.IObject, data2 objectivec.IObject, bias bool, data3 objectivec.IObject, output bool, activation objectivec.IObject, activation2 objectivec.IObject, activation3 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeLayerClass.class), objc.Sel("uniDirectionalLSTMLayerWithInputSize:outputSize:inputWeightsData:recursionWeightsData:hasBias:biasTermsData:sequenceOutput:inputGateActivation:cellGateActivation:hiddenOutputActivation:"), size, size2, data, data2, bias, data3, output, activation, activation2, activation3)
	return objectivec.Object{ID: rv}
}


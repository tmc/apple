// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNConvolutionGradientNode] class.
var (
	_MPSCNNConvolutionGradientNodeClass     MPSCNNConvolutionGradientNodeClass
	_MPSCNNConvolutionGradientNodeClassOnce sync.Once
)

func getMPSCNNConvolutionGradientNodeClass() MPSCNNConvolutionGradientNodeClass {
	_MPSCNNConvolutionGradientNodeClassOnce.Do(func() {
		_MPSCNNConvolutionGradientNodeClass = MPSCNNConvolutionGradientNodeClass{class: objc.GetClass("MPSCNNConvolutionGradientNode")}
	})
	return _MPSCNNConvolutionGradientNodeClass
}

// GetMPSCNNConvolutionGradientNodeClass returns the class object for MPSCNNConvolutionGradientNode.
func GetMPSCNNConvolutionGradientNodeClass() MPSCNNConvolutionGradientNodeClass {
	return getMPSCNNConvolutionGradientNodeClass()
}

type MPSCNNConvolutionGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionGradientNodeClass) Alloc() MPSCNNConvolutionGradientNode {
	rv := objc.Send[MPSCNNConvolutionGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient convolution kernel.
//
// # Initializers
//
//   - [MPSCNNConvolutionGradientNode.InitWithSourceGradientSourceImageConvolutionGradientStateWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientNode
type MPSCNNConvolutionGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNConvolutionGradientNodeFromID constructs a [MPSCNNConvolutionGradientNode] from an objc.ID.
//
// A representation of a gradient convolution kernel.
func MPSCNNConvolutionGradientNodeFromID(id objc.ID) MPSCNNConvolutionGradientNode {
	return MPSCNNConvolutionGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNConvolutionGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNConvolutionGradientNode.InitWithSourceGradientSourceImageConvolutionGradientStateWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientNode
type IMPSCNNConvolutionGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionGradientNode

	TrainingStyle() MPSNNTrainingStyle
}

// Init initializes the instance.
func (c MPSCNNConvolutionGradientNode) Init() MPSCNNConvolutionGradientNode {
	rv := objc.Send[MPSCNNConvolutionGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionGradientNode) Autorelease() MPSCNNConvolutionGradientNode {
	rv := objc.Send[MPSCNNConvolutionGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionGradientNode creates a new MPSCNNConvolutionGradientNode instance.
func NewMPSCNNConvolutionGradientNode() MPSCNNConvolutionGradientNode {
	class := getMPSCNNConvolutionGradientNodeClass()
	rv := objc.Send[MPSCNNConvolutionGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientNode/init(sourceGradient:sourceImage:convolutionGradientState:weights:)
func NewCNNConvolutionGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionGradientNode {
	instance := getMPSCNNConvolutionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return MPSCNNConvolutionGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientNode/init(sourceGradient:sourceImage:convolutionGradientState:weights:)
func (c MPSCNNConvolutionGradientNode) InitWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionGradientNode {
	rv := objc.Send[MPSCNNConvolutionGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode/trainingStyle
func (c MPSCNNConvolutionGradientNode) TrainingStyle() MPSNNTrainingStyle {
	rv := objc.Send[MPSNNTrainingStyle](c.ID, objc.Sel("trainingStyle"))
	return MPSNNTrainingStyle(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientNode/nodeWithSourceGradient:sourceImage:convolutionGradientState:weights:
func (_MPSCNNConvolutionGradientNodeClass MPSCNNConvolutionGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNConvolutionGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return MPSCNNConvolutionGradientNodeFromID(rv)
}

// Protocol methods for MPSNNTrainableNode

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode/trainingStyle
func (o MPSCNNConvolutionGradientNode) SetTrainingStyle(value MPSNNTrainingStyle) {
	objc.Send[struct{}](o.ID, objc.Sel("setTrainingStyle:"), value)
}

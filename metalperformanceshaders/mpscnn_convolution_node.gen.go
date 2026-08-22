// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNConvolutionNode] class.
var (
	_MPSCNNConvolutionNodeClass     MPSCNNConvolutionNodeClass
	_MPSCNNConvolutionNodeClassOnce sync.Once
)

func getMPSCNNConvolutionNodeClass() MPSCNNConvolutionNodeClass {
	_MPSCNNConvolutionNodeClassOnce.Do(func() {
		_MPSCNNConvolutionNodeClass = MPSCNNConvolutionNodeClass{class: objc.GetClass("MPSCNNConvolutionNode")}
	})
	return _MPSCNNConvolutionNodeClass
}

// GetMPSCNNConvolutionNodeClass returns the class object for MPSCNNConvolutionNode.
func GetMPSCNNConvolutionNodeClass() MPSCNNConvolutionNodeClass {
	return getMPSCNNConvolutionNodeClass()
}

type MPSCNNConvolutionNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionNodeClass) Alloc() MPSCNNConvolutionNode {
	rv := objc.Send[MPSCNNConvolutionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a convolution kernel.
//
// # Initializers
//
//   - [MPSCNNConvolutionNode.InitWithSourceWeights]
//
// # Instance Properties
//
//   - [MPSCNNConvolutionNode.AccumulatorPrecision]
//   - [MPSCNNConvolutionNode.SetAccumulatorPrecision]
//   - [MPSCNNConvolutionNode.ConvolutionGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode
type MPSCNNConvolutionNode struct {
	MPSNNFilterNode
}

// MPSCNNConvolutionNodeFromID constructs a [MPSCNNConvolutionNode] from an objc.ID.
//
// A representation of a convolution kernel.
func MPSCNNConvolutionNodeFromID(id objc.ID) MPSCNNConvolutionNode {
	return MPSCNNConvolutionNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNConvolutionNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionNode] class.
//
// # Initializers
//
//   - [IMPSCNNConvolutionNode.InitWithSourceWeights]
//
// # Instance Properties
//
//   - [IMPSCNNConvolutionNode.AccumulatorPrecision]
//   - [IMPSCNNConvolutionNode.SetAccumulatorPrecision]
//   - [IMPSCNNConvolutionNode.ConvolutionGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode
type IMPSCNNConvolutionNode interface {
	IMPSNNFilterNode
	MPSNNTrainableNode

	// Topic: Initializers

	InitWithSourceWeights(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionNode

	// Topic: Instance Properties

	AccumulatorPrecision() MPSNNConvolutionAccumulatorPrecisionOption
	SetAccumulatorPrecision(value MPSNNConvolutionAccumulatorPrecisionOption)
	ConvolutionGradientState() IMPSCNNConvolutionGradientStateNode
}

// Init initializes the instance.
func (c MPSCNNConvolutionNode) Init() MPSCNNConvolutionNode {
	rv := objc.Send[MPSCNNConvolutionNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionNode) Autorelease() MPSCNNConvolutionNode {
	rv := objc.Send[MPSCNNConvolutionNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionNode creates a new MPSCNNConvolutionNode instance.
func NewMPSCNNConvolutionNode() MPSCNNConvolutionNode {
	class := getMPSCNNConvolutionNodeClass()
	rv := objc.Send[MPSCNNConvolutionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode/init(source:weights:)
func NewCNNConvolutionNodeWithSourceWeights(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionNode {
	instance := getMPSCNNConvolutionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:weights:"), sourceNode, weights)
	return MPSCNNConvolutionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode/init(source:weights:)
func (c MPSCNNConvolutionNode) InitWithSourceWeights(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionNode {
	rv := objc.Send[MPSCNNConvolutionNode](c.ID, objc.Sel("initWithSource:weights:"), sourceNode, weights)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode/nodeWithSource:weights:
func (_MPSCNNConvolutionNodeClass MPSCNNConvolutionNodeClass) NodeWithSourceWeights(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNConvolutionNodeClass.class), objc.Sel("nodeWithSource:weights:"), sourceNode, weights)
	return MPSCNNConvolutionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode/accumulatorPrecision
func (c MPSCNNConvolutionNode) AccumulatorPrecision() MPSNNConvolutionAccumulatorPrecisionOption {
	rv := objc.Send[MPSNNConvolutionAccumulatorPrecisionOption](c.ID, objc.Sel("accumulatorPrecision"))
	return MPSNNConvolutionAccumulatorPrecisionOption(rv)
}
func (c MPSCNNConvolutionNode) SetAccumulatorPrecision(value MPSNNConvolutionAccumulatorPrecisionOption) {
	objc.Send[struct{}](c.ID, objc.Sel("setAccumulatorPrecision:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode/convolutionGradientState
func (c MPSCNNConvolutionNode) ConvolutionGradientState() IMPSCNNConvolutionGradientStateNode {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("convolutionGradientState"))
	return MPSCNNConvolutionGradientStateNodeFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode/trainingStyle
func (c MPSCNNConvolutionNode) TrainingStyle() MPSNNTrainingStyle {
	rv := objc.Send[MPSNNTrainingStyle](c.ID, objc.Sel("trainingStyle"))
	return MPSNNTrainingStyle(rv)
}
func (c MPSCNNConvolutionNode) SetTrainingStyle(value MPSNNTrainingStyle) {
	objc.Send[struct{}](c.ID, objc.Sel("setTrainingStyle:"), value)
}

// Protocol methods for MPSNNTrainableNode

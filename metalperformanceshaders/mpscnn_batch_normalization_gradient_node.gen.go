// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNBatchNormalizationGradientNode] class.
var (
	_MPSCNNBatchNormalizationGradientNodeClass     MPSCNNBatchNormalizationGradientNodeClass
	_MPSCNNBatchNormalizationGradientNodeClassOnce sync.Once
)

func getMPSCNNBatchNormalizationGradientNodeClass() MPSCNNBatchNormalizationGradientNodeClass {
	_MPSCNNBatchNormalizationGradientNodeClassOnce.Do(func() {
		_MPSCNNBatchNormalizationGradientNodeClass = MPSCNNBatchNormalizationGradientNodeClass{class: objc.GetClass("MPSCNNBatchNormalizationGradientNode")}
	})
	return _MPSCNNBatchNormalizationGradientNodeClass
}

// GetMPSCNNBatchNormalizationGradientNodeClass returns the class object for MPSCNNBatchNormalizationGradientNode.
func GetMPSCNNBatchNormalizationGradientNodeClass() MPSCNNBatchNormalizationGradientNodeClass {
	return getMPSCNNBatchNormalizationGradientNodeClass()
}

type MPSCNNBatchNormalizationGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBatchNormalizationGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBatchNormalizationGradientNodeClass) Alloc() MPSCNNBatchNormalizationGradientNode {
	rv := objc.Send[MPSCNNBatchNormalizationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient batch normalization kernel.
//
// # Initializers
//
//   - [MPSCNNBatchNormalizationGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradientNode
type MPSCNNBatchNormalizationGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNBatchNormalizationGradientNodeFromID constructs a [MPSCNNBatchNormalizationGradientNode] from an objc.ID.
//
// A representation of a gradient batch normalization kernel.
func MPSCNNBatchNormalizationGradientNodeFromID(id objc.ID) MPSCNNBatchNormalizationGradientNode {
	return MPSCNNBatchNormalizationGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNBatchNormalizationGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBatchNormalizationGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNBatchNormalizationGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradientNode
type IMPSCNNBatchNormalizationGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNBatchNormalizationGradientNode

	TrainingStyle() MPSNNTrainingStyle
}

// Init initializes the instance.
func (c MPSCNNBatchNormalizationGradientNode) Init() MPSCNNBatchNormalizationGradientNode {
	rv := objc.Send[MPSCNNBatchNormalizationGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBatchNormalizationGradientNode) Autorelease() MPSCNNBatchNormalizationGradientNode {
	rv := objc.Send[MPSCNNBatchNormalizationGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBatchNormalizationGradientNode creates a new MPSCNNBatchNormalizationGradientNode instance.
func NewMPSCNNBatchNormalizationGradientNode() MPSCNNBatchNormalizationGradientNode {
	class := getMPSCNNBatchNormalizationGradientNodeClass()
	rv := objc.Send[MPSCNNBatchNormalizationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewCNNBatchNormalizationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNBatchNormalizationGradientNode {
	instance := getMPSCNNBatchNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSCNNBatchNormalizationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:)
func (c MPSCNNBatchNormalizationGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNBatchNormalizationGradientNode {
	rv := objc.Send[MPSCNNBatchNormalizationGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode/trainingStyle
func (c MPSCNNBatchNormalizationGradientNode) TrainingStyle() MPSNNTrainingStyle {
	rv := objc.Send[MPSNNTrainingStyle](c.ID, objc.Sel("trainingStyle"))
	return MPSNNTrainingStyle(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSCNNBatchNormalizationGradientNodeClass MPSCNNBatchNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNBatchNormalizationGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNBatchNormalizationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSCNNBatchNormalizationGradientNodeFromID(rv)
}

// Protocol methods for MPSNNTrainableNode

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode/trainingStyle
func (o MPSCNNBatchNormalizationGradientNode) SetTrainingStyle(value MPSNNTrainingStyle) {
	objc.Send[struct{}](o.ID, objc.Sel("setTrainingStyle:"), value)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNInstanceNormalizationGradientNode] class.
var (
	_MPSCNNInstanceNormalizationGradientNodeClass     MPSCNNInstanceNormalizationGradientNodeClass
	_MPSCNNInstanceNormalizationGradientNodeClassOnce sync.Once
)

func getMPSCNNInstanceNormalizationGradientNodeClass() MPSCNNInstanceNormalizationGradientNodeClass {
	_MPSCNNInstanceNormalizationGradientNodeClassOnce.Do(func() {
		_MPSCNNInstanceNormalizationGradientNodeClass = MPSCNNInstanceNormalizationGradientNodeClass{class: objc.GetClass("MPSCNNInstanceNormalizationGradientNode")}
	})
	return _MPSCNNInstanceNormalizationGradientNodeClass
}

// GetMPSCNNInstanceNormalizationGradientNodeClass returns the class object for MPSCNNInstanceNormalizationGradientNode.
func GetMPSCNNInstanceNormalizationGradientNodeClass() MPSCNNInstanceNormalizationGradientNodeClass {
	return getMPSCNNInstanceNormalizationGradientNodeClass()
}

type MPSCNNInstanceNormalizationGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNInstanceNormalizationGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNInstanceNormalizationGradientNodeClass) Alloc() MPSCNNInstanceNormalizationGradientNode {
	rv := objc.Send[MPSCNNInstanceNormalizationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient instance normalization kernel.
//
// # Initializers
//
//   - [MPSCNNInstanceNormalizationGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientNode
type MPSCNNInstanceNormalizationGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNInstanceNormalizationGradientNodeFromID constructs a [MPSCNNInstanceNormalizationGradientNode] from an objc.ID.
//
// A representation of a gradient instance normalization kernel.
func MPSCNNInstanceNormalizationGradientNodeFromID(id objc.ID) MPSCNNInstanceNormalizationGradientNode {
	return MPSCNNInstanceNormalizationGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNInstanceNormalizationGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNInstanceNormalizationGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNInstanceNormalizationGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientNode
type IMPSCNNInstanceNormalizationGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNInstanceNormalizationGradientNode

	TrainingStyle() MPSNNTrainingStyle
}

// Init initializes the instance.
func (c MPSCNNInstanceNormalizationGradientNode) Init() MPSCNNInstanceNormalizationGradientNode {
	rv := objc.Send[MPSCNNInstanceNormalizationGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNInstanceNormalizationGradientNode) Autorelease() MPSCNNInstanceNormalizationGradientNode {
	rv := objc.Send[MPSCNNInstanceNormalizationGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNInstanceNormalizationGradientNode creates a new MPSCNNInstanceNormalizationGradientNode instance.
func NewMPSCNNInstanceNormalizationGradientNode() MPSCNNInstanceNormalizationGradientNode {
	class := getMPSCNNInstanceNormalizationGradientNodeClass()
	rv := objc.Send[MPSCNNInstanceNormalizationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewCNNInstanceNormalizationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNInstanceNormalizationGradientNode {
	instance := getMPSCNNInstanceNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSCNNInstanceNormalizationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:)
func (c MPSCNNInstanceNormalizationGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNInstanceNormalizationGradientNode {
	rv := objc.Send[MPSCNNInstanceNormalizationGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode/trainingStyle
func (c MPSCNNInstanceNormalizationGradientNode) TrainingStyle() MPSNNTrainingStyle {
	rv := objc.Send[MPSNNTrainingStyle](c.ID, objc.Sel("trainingStyle"))
	return MPSNNTrainingStyle(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSCNNInstanceNormalizationGradientNodeClass MPSCNNInstanceNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNInstanceNormalizationGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNInstanceNormalizationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSCNNInstanceNormalizationGradientNodeFromID(rv)
}

// Protocol methods for MPSNNTrainableNode

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode/trainingStyle
func (o MPSCNNInstanceNormalizationGradientNode) SetTrainingStyle(value MPSNNTrainingStyle) {
	objc.Send[struct{}](o.ID, objc.Sel("setTrainingStyle:"), value)
}

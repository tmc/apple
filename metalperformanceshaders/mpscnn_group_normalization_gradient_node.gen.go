// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNGroupNormalizationGradientNode] class.
var (
	_MPSCNNGroupNormalizationGradientNodeClass     MPSCNNGroupNormalizationGradientNodeClass
	_MPSCNNGroupNormalizationGradientNodeClassOnce sync.Once
)

func getMPSCNNGroupNormalizationGradientNodeClass() MPSCNNGroupNormalizationGradientNodeClass {
	_MPSCNNGroupNormalizationGradientNodeClassOnce.Do(func() {
		_MPSCNNGroupNormalizationGradientNodeClass = MPSCNNGroupNormalizationGradientNodeClass{class: objc.GetClass("MPSCNNGroupNormalizationGradientNode")}
	})
	return _MPSCNNGroupNormalizationGradientNodeClass
}

// GetMPSCNNGroupNormalizationGradientNodeClass returns the class object for MPSCNNGroupNormalizationGradientNode.
func GetMPSCNNGroupNormalizationGradientNodeClass() MPSCNNGroupNormalizationGradientNodeClass {
	return getMPSCNNGroupNormalizationGradientNodeClass()
}

type MPSCNNGroupNormalizationGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNGroupNormalizationGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNGroupNormalizationGradientNodeClass) Alloc() MPSCNNGroupNormalizationGradientNode {
	rv := objc.Send[MPSCNNGroupNormalizationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSCNNGroupNormalizationGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientNode
type MPSCNNGroupNormalizationGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNGroupNormalizationGradientNodeFromID constructs a [MPSCNNGroupNormalizationGradientNode] from an objc.ID.
func MPSCNNGroupNormalizationGradientNodeFromID(id objc.ID) MPSCNNGroupNormalizationGradientNode {
	return MPSCNNGroupNormalizationGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNGroupNormalizationGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNGroupNormalizationGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNGroupNormalizationGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientNode
type IMPSCNNGroupNormalizationGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNGroupNormalizationGradientNode

	TrainingStyle() MPSNNTrainingStyle
}

// Init initializes the instance.
func (c MPSCNNGroupNormalizationGradientNode) Init() MPSCNNGroupNormalizationGradientNode {
	rv := objc.Send[MPSCNNGroupNormalizationGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNGroupNormalizationGradientNode) Autorelease() MPSCNNGroupNormalizationGradientNode {
	rv := objc.Send[MPSCNNGroupNormalizationGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNGroupNormalizationGradientNode creates a new MPSCNNGroupNormalizationGradientNode instance.
func NewMPSCNNGroupNormalizationGradientNode() MPSCNNGroupNormalizationGradientNode {
	class := getMPSCNNGroupNormalizationGradientNodeClass()
	rv := objc.Send[MPSCNNGroupNormalizationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewCNNGroupNormalizationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNGroupNormalizationGradientNode {
	instance := getMPSCNNGroupNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSCNNGroupNormalizationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:)
func (c MPSCNNGroupNormalizationGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNGroupNormalizationGradientNode {
	rv := objc.Send[MPSCNNGroupNormalizationGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode/trainingStyle
func (c MPSCNNGroupNormalizationGradientNode) TrainingStyle() MPSNNTrainingStyle {
	rv := objc.Send[MPSNNTrainingStyle](c.ID, objc.Sel("trainingStyle"))
	return MPSNNTrainingStyle(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSCNNGroupNormalizationGradientNodeClass MPSCNNGroupNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNGroupNormalizationGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNGroupNormalizationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSCNNGroupNormalizationGradientNodeFromID(rv)
}

// Protocol methods for MPSNNTrainableNode

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNTrainableNode/trainingStyle
func (o MPSCNNGroupNormalizationGradientNode) SetTrainingStyle(value MPSNNTrainingStyle) {
	objc.Send[struct{}](o.ID, objc.Sel("setTrainingStyle:"), value)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReductionSpatialMeanGradientNode] class.
var (
	_MPSNNReductionSpatialMeanGradientNodeClass     MPSNNReductionSpatialMeanGradientNodeClass
	_MPSNNReductionSpatialMeanGradientNodeClassOnce sync.Once
)

func getMPSNNReductionSpatialMeanGradientNodeClass() MPSNNReductionSpatialMeanGradientNodeClass {
	_MPSNNReductionSpatialMeanGradientNodeClassOnce.Do(func() {
		_MPSNNReductionSpatialMeanGradientNodeClass = MPSNNReductionSpatialMeanGradientNodeClass{class: objc.GetClass("MPSNNReductionSpatialMeanGradientNode")}
	})
	return _MPSNNReductionSpatialMeanGradientNodeClass
}

// GetMPSNNReductionSpatialMeanGradientNodeClass returns the class object for MPSNNReductionSpatialMeanGradientNode.
func GetMPSNNReductionSpatialMeanGradientNodeClass() MPSNNReductionSpatialMeanGradientNodeClass {
	return getMPSNNReductionSpatialMeanGradientNodeClass()
}

type MPSNNReductionSpatialMeanGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReductionSpatialMeanGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReductionSpatialMeanGradientNodeClass) Alloc() MPSNNReductionSpatialMeanGradientNode {
	rv := objc.Send[MPSNNReductionSpatialMeanGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNReductionSpatialMeanGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionSpatialMeanGradientNode
type MPSNNReductionSpatialMeanGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNReductionSpatialMeanGradientNodeFromID constructs a [MPSNNReductionSpatialMeanGradientNode] from an objc.ID.
func MPSNNReductionSpatialMeanGradientNodeFromID(id objc.ID) MPSNNReductionSpatialMeanGradientNode {
	return MPSNNReductionSpatialMeanGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSNNReductionSpatialMeanGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReductionSpatialMeanGradientNode] class.
//
// # Initializers
//
//   - [IMPSNNReductionSpatialMeanGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionSpatialMeanGradientNode
type IMPSNNReductionSpatialMeanGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNReductionSpatialMeanGradientNode
}

// Init initializes the instance.
func (r MPSNNReductionSpatialMeanGradientNode) Init() MPSNNReductionSpatialMeanGradientNode {
	rv := objc.Send[MPSNNReductionSpatialMeanGradientNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReductionSpatialMeanGradientNode) Autorelease() MPSNNReductionSpatialMeanGradientNode {
	rv := objc.Send[MPSNNReductionSpatialMeanGradientNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReductionSpatialMeanGradientNode creates a new MPSNNReductionSpatialMeanGradientNode instance.
func NewMPSNNReductionSpatialMeanGradientNode() MPSNNReductionSpatialMeanGradientNode {
	class := getMPSNNReductionSpatialMeanGradientNodeClass()
	rv := objc.Send[MPSNNReductionSpatialMeanGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionSpatialMeanGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewReductionSpatialMeanGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNReductionSpatialMeanGradientNode {
	instance := getMPSNNReductionSpatialMeanGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSNNReductionSpatialMeanGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionSpatialMeanGradientNode/init(sourceGradient:sourceImage:gradientState:)
func (r MPSNNReductionSpatialMeanGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNReductionSpatialMeanGradientNode {
	rv := objc.Send[MPSNNReductionSpatialMeanGradientNode](r.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReductionSpatialMeanGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNReductionSpatialMeanGradientNodeClass MPSNNReductionSpatialMeanGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNReductionSpatialMeanGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNReductionSpatialMeanGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSNNReductionSpatialMeanGradientNodeFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSoftMaxGradientNode] class.
var (
	_MPSCNNSoftMaxGradientNodeClass     MPSCNNSoftMaxGradientNodeClass
	_MPSCNNSoftMaxGradientNodeClassOnce sync.Once
)

func getMPSCNNSoftMaxGradientNodeClass() MPSCNNSoftMaxGradientNodeClass {
	_MPSCNNSoftMaxGradientNodeClassOnce.Do(func() {
		_MPSCNNSoftMaxGradientNodeClass = MPSCNNSoftMaxGradientNodeClass{class: objc.GetClass("MPSCNNSoftMaxGradientNode")}
	})
	return _MPSCNNSoftMaxGradientNodeClass
}

// GetMPSCNNSoftMaxGradientNodeClass returns the class object for MPSCNNSoftMaxGradientNode.
func GetMPSCNNSoftMaxGradientNodeClass() MPSCNNSoftMaxGradientNodeClass {
	return getMPSCNNSoftMaxGradientNodeClass()
}

type MPSCNNSoftMaxGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSoftMaxGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSoftMaxGradientNodeClass) Alloc() MPSCNNSoftMaxGradientNode {
	rv := objc.Send[MPSCNNSoftMaxGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient softmax filter.
//
// # Initializers
//
//   - [MPSCNNSoftMaxGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradientNode
type MPSCNNSoftMaxGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNSoftMaxGradientNodeFromID constructs a [MPSCNNSoftMaxGradientNode] from an objc.ID.
//
// A representation of a gradient softmax filter.
func MPSCNNSoftMaxGradientNodeFromID(id objc.ID) MPSCNNSoftMaxGradientNode {
	return MPSCNNSoftMaxGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNSoftMaxGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSoftMaxGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNSoftMaxGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradientNode
type IMPSCNNSoftMaxGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNSoftMaxGradientNode
}

// Init initializes the instance.
func (c MPSCNNSoftMaxGradientNode) Init() MPSCNNSoftMaxGradientNode {
	rv := objc.Send[MPSCNNSoftMaxGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSoftMaxGradientNode) Autorelease() MPSCNNSoftMaxGradientNode {
	rv := objc.Send[MPSCNNSoftMaxGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSoftMaxGradientNode creates a new MPSCNNSoftMaxGradientNode instance.
func NewMPSCNNSoftMaxGradientNode() MPSCNNSoftMaxGradientNode {
	class := getMPSCNNSoftMaxGradientNodeClass()
	rv := objc.Send[MPSCNNSoftMaxGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewCNNSoftMaxGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNSoftMaxGradientNode {
	instance := getMPSCNNSoftMaxGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSCNNSoftMaxGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradientNode/init(sourceGradient:sourceImage:gradientState:)
func (c MPSCNNSoftMaxGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNSoftMaxGradientNode {
	rv := objc.Send[MPSCNNSoftMaxGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSCNNSoftMaxGradientNodeClass MPSCNNSoftMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNSoftMaxGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNSoftMaxGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSCNNSoftMaxGradientNodeFromID(rv)
}

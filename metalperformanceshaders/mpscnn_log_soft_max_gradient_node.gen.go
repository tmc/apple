// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNLogSoftMaxGradientNode] class.
var (
	_MPSCNNLogSoftMaxGradientNodeClass     MPSCNNLogSoftMaxGradientNodeClass
	_MPSCNNLogSoftMaxGradientNodeClassOnce sync.Once
)

func getMPSCNNLogSoftMaxGradientNodeClass() MPSCNNLogSoftMaxGradientNodeClass {
	_MPSCNNLogSoftMaxGradientNodeClassOnce.Do(func() {
		_MPSCNNLogSoftMaxGradientNodeClass = MPSCNNLogSoftMaxGradientNodeClass{class: objc.GetClass("MPSCNNLogSoftMaxGradientNode")}
	})
	return _MPSCNNLogSoftMaxGradientNodeClass
}

// GetMPSCNNLogSoftMaxGradientNodeClass returns the class object for MPSCNNLogSoftMaxGradientNode.
func GetMPSCNNLogSoftMaxGradientNodeClass() MPSCNNLogSoftMaxGradientNodeClass {
	return getMPSCNNLogSoftMaxGradientNodeClass()
}

type MPSCNNLogSoftMaxGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLogSoftMaxGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLogSoftMaxGradientNodeClass) Alloc() MPSCNNLogSoftMaxGradientNode {
	rv := objc.Send[MPSCNNLogSoftMaxGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient logarithmic softmax filter kernel.
//
// # Initializers
//
//   - [MPSCNNLogSoftMaxGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradientNode
type MPSCNNLogSoftMaxGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNLogSoftMaxGradientNodeFromID constructs a [MPSCNNLogSoftMaxGradientNode] from an objc.ID.
//
// A representation of a gradient logarithmic softmax filter kernel.
func MPSCNNLogSoftMaxGradientNodeFromID(id objc.ID) MPSCNNLogSoftMaxGradientNode {
	return MPSCNNLogSoftMaxGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNLogSoftMaxGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLogSoftMaxGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNLogSoftMaxGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradientNode
type IMPSCNNLogSoftMaxGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNLogSoftMaxGradientNode
}

// Init initializes the instance.
func (c MPSCNNLogSoftMaxGradientNode) Init() MPSCNNLogSoftMaxGradientNode {
	rv := objc.Send[MPSCNNLogSoftMaxGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLogSoftMaxGradientNode) Autorelease() MPSCNNLogSoftMaxGradientNode {
	rv := objc.Send[MPSCNNLogSoftMaxGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLogSoftMaxGradientNode creates a new MPSCNNLogSoftMaxGradientNode instance.
func NewMPSCNNLogSoftMaxGradientNode() MPSCNNLogSoftMaxGradientNode {
	class := getMPSCNNLogSoftMaxGradientNodeClass()
	rv := objc.Send[MPSCNNLogSoftMaxGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewCNNLogSoftMaxGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNLogSoftMaxGradientNode {
	instance := getMPSCNNLogSoftMaxGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSCNNLogSoftMaxGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradientNode/init(sourceGradient:sourceImage:gradientState:)
func (c MPSCNNLogSoftMaxGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNLogSoftMaxGradientNode {
	rv := objc.Send[MPSCNNLogSoftMaxGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSCNNLogSoftMaxGradientNodeClass MPSCNNLogSoftMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSCNNLogSoftMaxGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNLogSoftMaxGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSCNNLogSoftMaxGradientNodeFromID(rv)
}

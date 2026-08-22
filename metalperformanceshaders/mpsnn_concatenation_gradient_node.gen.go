// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNConcatenationGradientNode] class.
var (
	_MPSNNConcatenationGradientNodeClass     MPSNNConcatenationGradientNodeClass
	_MPSNNConcatenationGradientNodeClassOnce sync.Once
)

func getMPSNNConcatenationGradientNodeClass() MPSNNConcatenationGradientNodeClass {
	_MPSNNConcatenationGradientNodeClassOnce.Do(func() {
		_MPSNNConcatenationGradientNodeClass = MPSNNConcatenationGradientNodeClass{class: objc.GetClass("MPSNNConcatenationGradientNode")}
	})
	return _MPSNNConcatenationGradientNodeClass
}

// GetMPSNNConcatenationGradientNodeClass returns the class object for MPSNNConcatenationGradientNode.
func GetMPSNNConcatenationGradientNodeClass() MPSNNConcatenationGradientNodeClass {
	return getMPSNNConcatenationGradientNodeClass()
}

type MPSNNConcatenationGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNConcatenationGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNConcatenationGradientNodeClass) Alloc() MPSNNConcatenationGradientNode {
	rv := objc.Send[MPSNNConcatenationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of the results from one or more gradient kernels.
//
// # Initializers
//
//   - [MPSNNConcatenationGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationGradientNode
type MPSNNConcatenationGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNConcatenationGradientNodeFromID constructs a [MPSNNConcatenationGradientNode] from an objc.ID.
//
// A representation of the results from one or more gradient kernels.
func MPSNNConcatenationGradientNodeFromID(id objc.ID) MPSNNConcatenationGradientNode {
	return MPSNNConcatenationGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSNNConcatenationGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNConcatenationGradientNode] class.
//
// # Initializers
//
//   - [IMPSNNConcatenationGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationGradientNode
type IMPSNNConcatenationGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientState(gradientSourceNode IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNConcatenationGradientNode
}

// Init initializes the instance.
func (c MPSNNConcatenationGradientNode) Init() MPSNNConcatenationGradientNode {
	rv := objc.Send[MPSNNConcatenationGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNConcatenationGradientNode) Autorelease() MPSNNConcatenationGradientNode {
	rv := objc.Send[MPSNNConcatenationGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNConcatenationGradientNode creates a new MPSNNConcatenationGradientNode instance.
func NewMPSNNConcatenationGradientNode() MPSNNConcatenationGradientNode {
	class := getMPSNNConcatenationGradientNodeClass()
	rv := objc.Send[MPSNNConcatenationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewConcatenationGradientNodeWithSourceGradientSourceImageGradientState(gradientSourceNode IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNConcatenationGradientNode {
	instance := getMPSNNConcatenationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradientSourceNode, sourceImage, gradientState)
	return MPSNNConcatenationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationGradientNode/init(sourceGradient:sourceImage:gradientState:)
func (c MPSNNConcatenationGradientNode) InitWithSourceGradientSourceImageGradientState(gradientSourceNode IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNConcatenationGradientNode {
	rv := objc.Send[MPSNNConcatenationGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), gradientSourceNode, sourceImage, gradientState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNConcatenationGradientNodeClass MPSNNConcatenationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(gradientSourceNode IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNConcatenationGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNConcatenationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), gradientSourceNode, sourceImage, gradientState)
	return MPSNNConcatenationGradientNodeFromID(rv)
}

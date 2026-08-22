// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNPadGradientNode] class.
var (
	_MPSNNPadGradientNodeClass     MPSNNPadGradientNodeClass
	_MPSNNPadGradientNodeClassOnce sync.Once
)

func getMPSNNPadGradientNodeClass() MPSNNPadGradientNodeClass {
	_MPSNNPadGradientNodeClassOnce.Do(func() {
		_MPSNNPadGradientNodeClass = MPSNNPadGradientNodeClass{class: objc.GetClass("MPSNNPadGradientNode")}
	})
	return _MPSNNPadGradientNodeClass
}

// GetMPSNNPadGradientNodeClass returns the class object for MPSNNPadGradientNode.
func GetMPSNNPadGradientNodeClass() MPSNNPadGradientNodeClass {
	return getMPSNNPadGradientNodeClass()
}

type MPSNNPadGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNPadGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPadGradientNodeClass) Alloc() MPSNNPadGradientNode {
	rv := objc.Send[MPSNNPadGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNPadGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradientNode
type MPSNNPadGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNPadGradientNodeFromID constructs a [MPSNNPadGradientNode] from an objc.ID.
func MPSNNPadGradientNodeFromID(id objc.ID) MPSNNPadGradientNode {
	return MPSNNPadGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSNNPadGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNPadGradientNode] class.
//
// # Initializers
//
//   - [IMPSNNPadGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradientNode
type IMPSNNPadGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNPadGradientNode
}

// Init initializes the instance.
func (p MPSNNPadGradientNode) Init() MPSNNPadGradientNode {
	rv := objc.Send[MPSNNPadGradientNode](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPadGradientNode) Autorelease() MPSNNPadGradientNode {
	rv := objc.Send[MPSNNPadGradientNode](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPadGradientNode creates a new MPSNNPadGradientNode instance.
func NewMPSNNPadGradientNode() MPSNNPadGradientNode {
	class := getMPSNNPadGradientNodeClass()
	rv := objc.Send[MPSNNPadGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewPadGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNPadGradientNode {
	instance := getMPSNNPadGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSNNPadGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradientNode/init(sourceGradient:sourceImage:gradientState:)
func (p MPSNNPadGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNPadGradientNode {
	rv := objc.Send[MPSNNPadGradientNode](p.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNPadGradientNodeClass MPSNNPadGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNPadGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNPadGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSNNPadGradientNodeFromID(rv)
}

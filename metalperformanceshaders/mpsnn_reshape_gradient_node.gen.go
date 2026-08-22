// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReshapeGradientNode] class.
var (
	_MPSNNReshapeGradientNodeClass     MPSNNReshapeGradientNodeClass
	_MPSNNReshapeGradientNodeClassOnce sync.Once
)

func getMPSNNReshapeGradientNodeClass() MPSNNReshapeGradientNodeClass {
	_MPSNNReshapeGradientNodeClassOnce.Do(func() {
		_MPSNNReshapeGradientNodeClass = MPSNNReshapeGradientNodeClass{class: objc.GetClass("MPSNNReshapeGradientNode")}
	})
	return _MPSNNReshapeGradientNodeClass
}

// GetMPSNNReshapeGradientNodeClass returns the class object for MPSNNReshapeGradientNode.
func GetMPSNNReshapeGradientNodeClass() MPSNNReshapeGradientNodeClass {
	return getMPSNNReshapeGradientNodeClass()
}

type MPSNNReshapeGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReshapeGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReshapeGradientNodeClass) Alloc() MPSNNReshapeGradientNode {
	rv := objc.Send[MPSNNReshapeGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNReshapeGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradientNode
type MPSNNReshapeGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNReshapeGradientNodeFromID constructs a [MPSNNReshapeGradientNode] from an objc.ID.
func MPSNNReshapeGradientNodeFromID(id objc.ID) MPSNNReshapeGradientNode {
	return MPSNNReshapeGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSNNReshapeGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReshapeGradientNode] class.
//
// # Initializers
//
//   - [IMPSNNReshapeGradientNode.InitWithSourceGradientSourceImageGradientState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradientNode
type IMPSNNReshapeGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNReshapeGradientNode
}

// Init initializes the instance.
func (r MPSNNReshapeGradientNode) Init() MPSNNReshapeGradientNode {
	rv := objc.Send[MPSNNReshapeGradientNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReshapeGradientNode) Autorelease() MPSNNReshapeGradientNode {
	rv := objc.Send[MPSNNReshapeGradientNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReshapeGradientNode creates a new MPSNNReshapeGradientNode instance.
func NewMPSNNReshapeGradientNode() MPSNNReshapeGradientNode {
	class := getMPSNNReshapeGradientNodeClass()
	rv := objc.Send[MPSNNReshapeGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewReshapeGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNReshapeGradientNode {
	instance := getMPSNNReshapeGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSNNReshapeGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradientNode/init(sourceGradient:sourceImage:gradientState:)
func (r MPSNNReshapeGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNReshapeGradientNode {
	rv := objc.Send[MPSNNReshapeGradientNode](r.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNReshapeGradientNodeClass MPSNNReshapeGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNReshapeGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNReshapeGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSNNReshapeGradientNodeFromID(rv)
}

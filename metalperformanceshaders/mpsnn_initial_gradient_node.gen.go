// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNInitialGradientNode] class.
var (
	_MPSNNInitialGradientNodeClass     MPSNNInitialGradientNodeClass
	_MPSNNInitialGradientNodeClassOnce sync.Once
)

func getMPSNNInitialGradientNodeClass() MPSNNInitialGradientNodeClass {
	_MPSNNInitialGradientNodeClassOnce.Do(func() {
		_MPSNNInitialGradientNodeClass = MPSNNInitialGradientNodeClass{class: objc.GetClass("MPSNNInitialGradientNode")}
	})
	return _MPSNNInitialGradientNodeClass
}

// GetMPSNNInitialGradientNodeClass returns the class object for MPSNNInitialGradientNode.
func GetMPSNNInitialGradientNodeClass() MPSNNInitialGradientNodeClass {
	return getMPSNNInitialGradientNodeClass()
}

type MPSNNInitialGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNInitialGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNInitialGradientNodeClass) Alloc() MPSNNInitialGradientNode {
	rv := objc.Send[MPSNNInitialGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNInitialGradientNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradientNode
type MPSNNInitialGradientNode struct {
	MPSNNFilterNode
}

// MPSNNInitialGradientNodeFromID constructs a [MPSNNInitialGradientNode] from an objc.ID.
func MPSNNInitialGradientNodeFromID(id objc.ID) MPSNNInitialGradientNode {
	return MPSNNInitialGradientNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSNNInitialGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNInitialGradientNode] class.
//
// # Initializers
//
//   - [IMPSNNInitialGradientNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradientNode
type IMPSNNInitialGradientNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSource(source IMPSNNImageNode) MPSNNInitialGradientNode
}

// Init initializes the instance.
func (i MPSNNInitialGradientNode) Init() MPSNNInitialGradientNode {
	rv := objc.Send[MPSNNInitialGradientNode](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSNNInitialGradientNode) Autorelease() MPSNNInitialGradientNode {
	rv := objc.Send[MPSNNInitialGradientNode](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNInitialGradientNode creates a new MPSNNInitialGradientNode instance.
func NewMPSNNInitialGradientNode() MPSNNInitialGradientNode {
	class := getMPSNNInitialGradientNodeClass()
	rv := objc.Send[MPSNNInitialGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradientNode/init(source:)
func NewInitialGradientNodeWithSource(source IMPSNNImageNode) MPSNNInitialGradientNode {
	instance := getMPSNNInitialGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSNNInitialGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradientNode/init(source:)
func (i MPSNNInitialGradientNode) InitWithSource(source IMPSNNImageNode) MPSNNInitialGradientNode {
	rv := objc.Send[MPSNNInitialGradientNode](i.ID, objc.Sel("initWithSource:"), source)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradientNode/nodeWithSource:
func (_MPSNNInitialGradientNodeClass MPSNNInitialGradientNodeClass) NodeWithSource(source IMPSNNImageNode) MPSNNInitialGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNInitialGradientNodeClass.class), objc.Sel("nodeWithSource:"), source)
	return MPSNNInitialGradientNodeFromID(rv)
}

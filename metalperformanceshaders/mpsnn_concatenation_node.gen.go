// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNConcatenationNode] class.
var (
	_MPSNNConcatenationNodeClass     MPSNNConcatenationNodeClass
	_MPSNNConcatenationNodeClassOnce sync.Once
)

func getMPSNNConcatenationNodeClass() MPSNNConcatenationNodeClass {
	_MPSNNConcatenationNodeClassOnce.Do(func() {
		_MPSNNConcatenationNodeClass = MPSNNConcatenationNodeClass{class: objc.GetClass("MPSNNConcatenationNode")}
	})
	return _MPSNNConcatenationNodeClass
}

// GetMPSNNConcatenationNodeClass returns the class object for MPSNNConcatenationNode.
func GetMPSNNConcatenationNodeClass() MPSNNConcatenationNodeClass {
	return getMPSNNConcatenationNodeClass()
}

type MPSNNConcatenationNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNConcatenationNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNConcatenationNodeClass) Alloc() MPSNNConcatenationNode {
	rv := objc.Send[MPSNNConcatenationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of the results from one or more kernels.
//
// # Initializers
//
//   - [MPSNNConcatenationNode.InitWithSources]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationNode
type MPSNNConcatenationNode struct {
	MPSNNFilterNode
}

// MPSNNConcatenationNodeFromID constructs a [MPSNNConcatenationNode] from an objc.ID.
//
// A representation of the results from one or more kernels.
func MPSNNConcatenationNodeFromID(id objc.ID) MPSNNConcatenationNode {
	return MPSNNConcatenationNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSNNConcatenationNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNConcatenationNode] class.
//
// # Initializers
//
//   - [IMPSNNConcatenationNode.InitWithSources]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationNode
type IMPSNNConcatenationNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSources(sourceNodes []MPSNNImageNode) MPSNNConcatenationNode
}

// Init initializes the instance.
func (c MPSNNConcatenationNode) Init() MPSNNConcatenationNode {
	rv := objc.Send[MPSNNConcatenationNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNConcatenationNode) Autorelease() MPSNNConcatenationNode {
	rv := objc.Send[MPSNNConcatenationNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNConcatenationNode creates a new MPSNNConcatenationNode instance.
func NewMPSNNConcatenationNode() MPSNNConcatenationNode {
	class := getMPSNNConcatenationNodeClass()
	rv := objc.Send[MPSNNConcatenationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationNode/init(sources:)
func NewConcatenationNodeWithSources(sourceNodes []MPSNNImageNode) MPSNNConcatenationNode {
	instance := getMPSNNConcatenationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), objectivec.IObjectSliceToNSArray(sourceNodes))
	return MPSNNConcatenationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationNode/init(sources:)
func (c MPSNNConcatenationNode) InitWithSources(sourceNodes []MPSNNImageNode) MPSNNConcatenationNode {
	rv := objc.Send[MPSNNConcatenationNode](c.ID, objc.Sel("initWithSources:"), objectivec.IObjectSliceToNSArray(sourceNodes))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationNode/nodeWithSources:
func (_MPSNNConcatenationNodeClass MPSNNConcatenationNodeClass) NodeWithSources(sourceNodes []MPSNNImageNode) MPSNNConcatenationNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNConcatenationNodeClass.class), objc.Sel("nodeWithSources:"), objectivec.IObjectSliceToNSArray(sourceNodes))
	return MPSNNConcatenationNodeFromID(rv)
}

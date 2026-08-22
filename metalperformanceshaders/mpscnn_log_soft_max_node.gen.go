// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNLogSoftMaxNode] class.
var (
	_MPSCNNLogSoftMaxNodeClass     MPSCNNLogSoftMaxNodeClass
	_MPSCNNLogSoftMaxNodeClassOnce sync.Once
)

func getMPSCNNLogSoftMaxNodeClass() MPSCNNLogSoftMaxNodeClass {
	_MPSCNNLogSoftMaxNodeClassOnce.Do(func() {
		_MPSCNNLogSoftMaxNodeClass = MPSCNNLogSoftMaxNodeClass{class: objc.GetClass("MPSCNNLogSoftMaxNode")}
	})
	return _MPSCNNLogSoftMaxNodeClass
}

// GetMPSCNNLogSoftMaxNodeClass returns the class object for MPSCNNLogSoftMaxNode.
func GetMPSCNNLogSoftMaxNodeClass() MPSCNNLogSoftMaxNodeClass {
	return getMPSCNNLogSoftMaxNodeClass()
}

type MPSCNNLogSoftMaxNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLogSoftMaxNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLogSoftMaxNodeClass) Alloc() MPSCNNLogSoftMaxNode {
	rv := objc.Send[MPSCNNLogSoftMaxNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a logarithmic softmax filter kernel.
//
// # Initializers
//
//   - [MPSCNNLogSoftMaxNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxNode
type MPSCNNLogSoftMaxNode struct {
	MPSNNFilterNode
}

// MPSCNNLogSoftMaxNodeFromID constructs a [MPSCNNLogSoftMaxNode] from an objc.ID.
//
// A representation of a logarithmic softmax filter kernel.
func MPSCNNLogSoftMaxNodeFromID(id objc.ID) MPSCNNLogSoftMaxNode {
	return MPSCNNLogSoftMaxNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNLogSoftMaxNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLogSoftMaxNode] class.
//
// # Initializers
//
//   - [IMPSCNNLogSoftMaxNode.InitWithSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxNode
type IMPSCNNLogSoftMaxNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNLogSoftMaxNode
}

// Init initializes the instance.
func (c MPSCNNLogSoftMaxNode) Init() MPSCNNLogSoftMaxNode {
	rv := objc.Send[MPSCNNLogSoftMaxNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLogSoftMaxNode) Autorelease() MPSCNNLogSoftMaxNode {
	rv := objc.Send[MPSCNNLogSoftMaxNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLogSoftMaxNode creates a new MPSCNNLogSoftMaxNode instance.
func NewMPSCNNLogSoftMaxNode() MPSCNNLogSoftMaxNode {
	class := getMPSCNNLogSoftMaxNodeClass()
	rv := objc.Send[MPSCNNLogSoftMaxNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxNode/init(source:)
func NewCNNLogSoftMaxNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNLogSoftMaxNode {
	instance := getMPSCNNLogSoftMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNLogSoftMaxNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxNode/init(source:)
func (c MPSCNNLogSoftMaxNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNLogSoftMaxNode {
	rv := objc.Send[MPSCNNLogSoftMaxNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxNode/nodeWithSource:
func (_MPSCNNLogSoftMaxNodeClass MPSCNNLogSoftMaxNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNLogSoftMaxNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNLogSoftMaxNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNLogSoftMaxNodeFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNUnaryReductionNode] class.
var (
	_MPSNNUnaryReductionNodeClass     MPSNNUnaryReductionNodeClass
	_MPSNNUnaryReductionNodeClassOnce sync.Once
)

func getMPSNNUnaryReductionNodeClass() MPSNNUnaryReductionNodeClass {
	_MPSNNUnaryReductionNodeClassOnce.Do(func() {
		_MPSNNUnaryReductionNodeClass = MPSNNUnaryReductionNodeClass{class: objc.GetClass("MPSNNUnaryReductionNode")}
	})
	return _MPSNNUnaryReductionNodeClass
}

// GetMPSNNUnaryReductionNodeClass returns the class object for MPSNNUnaryReductionNode.
func GetMPSNNUnaryReductionNodeClass() MPSNNUnaryReductionNodeClass {
	return getMPSNNUnaryReductionNodeClass()
}

type MPSNNUnaryReductionNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNUnaryReductionNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNUnaryReductionNodeClass) Alloc() MPSNNUnaryReductionNode {
	rv := objc.Send[MPSNNUnaryReductionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNUnaryReductionNode.InitWithSource]
//
// # Instance Properties
//
//   - [MPSNNUnaryReductionNode.ClipRectSource]
//   - [MPSNNUnaryReductionNode.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode
type MPSNNUnaryReductionNode struct {
	MPSNNFilterNode
}

// MPSNNUnaryReductionNodeFromID constructs a [MPSNNUnaryReductionNode] from an objc.ID.
func MPSNNUnaryReductionNodeFromID(id objc.ID) MPSNNUnaryReductionNode {
	return MPSNNUnaryReductionNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSNNUnaryReductionNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNUnaryReductionNode] class.
//
// # Initializers
//
//   - [IMPSNNUnaryReductionNode.InitWithSource]
//
// # Instance Properties
//
//   - [IMPSNNUnaryReductionNode.ClipRectSource]
//   - [IMPSNNUnaryReductionNode.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode
type IMPSNNUnaryReductionNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSNNUnaryReductionNode

	// Topic: Instance Properties

	ClipRectSource() metal.MTLRegion
	SetClipRectSource(value metal.MTLRegion)
}

// Init initializes the instance.
func (u MPSNNUnaryReductionNode) Init() MPSNNUnaryReductionNode {
	rv := objc.Send[MPSNNUnaryReductionNode](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u MPSNNUnaryReductionNode) Autorelease() MPSNNUnaryReductionNode {
	rv := objc.Send[MPSNNUnaryReductionNode](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNUnaryReductionNode creates a new MPSNNUnaryReductionNode instance.
func NewMPSNNUnaryReductionNode() MPSNNUnaryReductionNode {
	class := getMPSNNUnaryReductionNodeClass()
	rv := objc.Send[MPSNNUnaryReductionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewUnaryReductionNodeWithSource(sourceNode IMPSNNImageNode) MPSNNUnaryReductionNode {
	instance := getMPSNNUnaryReductionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNUnaryReductionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func (u MPSNNUnaryReductionNode) InitWithSource(sourceNode IMPSNNImageNode) MPSNNUnaryReductionNode {
	rv := objc.Send[MPSNNUnaryReductionNode](u.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/nodeWithSource:
func (_MPSNNUnaryReductionNodeClass MPSNNUnaryReductionNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSNNUnaryReductionNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNUnaryReductionNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSNNUnaryReductionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/clipRectSource
func (u MPSNNUnaryReductionNode) ClipRectSource() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](u.ID, objc.Sel("clipRectSource"))
	return metal.MTLRegion(rv)
}
func (u MPSNNUnaryReductionNode) SetClipRectSource(value metal.MTLRegion) {
	objc.Send[struct{}](u.ID, objc.Sel("setClipRectSource:"), value)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNPadNode] class.
var (
	_MPSNNPadNodeClass     MPSNNPadNodeClass
	_MPSNNPadNodeClassOnce sync.Once
)

func getMPSNNPadNodeClass() MPSNNPadNodeClass {
	_MPSNNPadNodeClassOnce.Do(func() {
		_MPSNNPadNodeClass = MPSNNPadNodeClass{class: objc.GetClass("MPSNNPadNode")}
	})
	return _MPSNNPadNodeClass
}

// GetMPSNNPadNodeClass returns the class object for MPSNNPadNode.
func GetMPSNNPadNodeClass() MPSNNPadNodeClass {
	return getMPSNNPadNodeClass()
}

type MPSNNPadNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNPadNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPadNodeClass) Alloc() MPSNNPadNode {
	rv := objc.Send[MPSNNPadNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNPadNode.InitWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode]
//
// # Instance Properties
//
//   - [MPSNNPadNode.FillValue]
//   - [MPSNNPadNode.SetFillValue]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadNode
type MPSNNPadNode struct {
	MPSNNFilterNode
}

// MPSNNPadNodeFromID constructs a [MPSNNPadNode] from an objc.ID.
func MPSNNPadNodeFromID(id objc.ID) MPSNNPadNode {
	return MPSNNPadNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSNNPadNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNPadNode] class.
//
// # Initializers
//
//   - [IMPSNNPadNode.InitWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode]
//
// # Instance Properties
//
//   - [IMPSNNPadNode.FillValue]
//   - [IMPSNNPadNode.SetFillValue]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadNode
type IMPSNNPadNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source IMPSNNImageNode, paddingSizeBefore MPSImageCoordinate, paddingSizeAfter MPSImageCoordinate, edgeMode MPSImageEdgeMode) MPSNNPadNode

	// Topic: Instance Properties

	FillValue() float32
	SetFillValue(value float32)
}

// Init initializes the instance.
func (p MPSNNPadNode) Init() MPSNNPadNode {
	rv := objc.Send[MPSNNPadNode](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPadNode) Autorelease() MPSNNPadNode {
	rv := objc.Send[MPSNNPadNode](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPadNode creates a new MPSNNPadNode instance.
func NewMPSNNPadNode() MPSNNPadNode {
	class := getMPSNNPadNodeClass()
	rv := objc.Send[MPSNNPadNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadNode/init(source:paddingSizeBefore:paddingSizeAfter:edgeMode:)
func NewPadNodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source IMPSNNImageNode, paddingSizeBefore MPSImageCoordinate, paddingSizeAfter MPSImageCoordinate, edgeMode MPSImageEdgeMode) MPSNNPadNode {
	instance := getMPSNNPadNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), source, paddingSizeBefore, paddingSizeAfter, edgeMode)
	return MPSNNPadNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadNode/init(source:paddingSizeBefore:paddingSizeAfter:edgeMode:)
func (p MPSNNPadNode) InitWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source IMPSNNImageNode, paddingSizeBefore MPSImageCoordinate, paddingSizeAfter MPSImageCoordinate, edgeMode MPSImageEdgeMode) MPSNNPadNode {
	rv := objc.Send[MPSNNPadNode](p.ID, objc.Sel("initWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), source, paddingSizeBefore, paddingSizeAfter, edgeMode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadNode/nodeWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:
func (_MPSNNPadNodeClass MPSNNPadNodeClass) NodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source IMPSNNImageNode, paddingSizeBefore MPSImageCoordinate, paddingSizeAfter MPSImageCoordinate, edgeMode MPSImageEdgeMode) MPSNNPadNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNPadNodeClass.class), objc.Sel("nodeWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), source, paddingSizeBefore, paddingSizeAfter, edgeMode)
	return MPSNNPadNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadNode/fillValue
func (p MPSNNPadNode) FillValue() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("fillValue"))
	return rv
}
func (p MPSNNPadNode) SetFillValue(value float32) {
	objc.Send[struct{}](p.ID, objc.Sel("setFillValue:"), value)
}

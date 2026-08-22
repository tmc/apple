// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNGradientFilterNode] class.
var (
	_MPSNNGradientFilterNodeClass     MPSNNGradientFilterNodeClass
	_MPSNNGradientFilterNodeClassOnce sync.Once
)

func getMPSNNGradientFilterNodeClass() MPSNNGradientFilterNodeClass {
	_MPSNNGradientFilterNodeClassOnce.Do(func() {
		_MPSNNGradientFilterNodeClass = MPSNNGradientFilterNodeClass{class: objc.GetClass("MPSNNGradientFilterNode")}
	})
	return _MPSNNGradientFilterNodeClass
}

// GetMPSNNGradientFilterNodeClass returns the class object for MPSNNGradientFilterNode.
func GetMPSNNGradientFilterNodeClass() MPSNNGradientFilterNodeClass {
	return getMPSNNGradientFilterNodeClass()
}

type MPSNNGradientFilterNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNGradientFilterNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGradientFilterNodeClass) Alloc() MPSNNGradientFilterNode {
	rv := objc.Send[MPSNNGradientFilterNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientFilterNode
type MPSNNGradientFilterNode struct {
	MPSNNFilterNode
}

// MPSNNGradientFilterNodeFromID constructs a [MPSNNGradientFilterNode] from an objc.ID.
//
// A representation of a gradient filter.
func MPSNNGradientFilterNodeFromID(id objc.ID) MPSNNGradientFilterNode {
	return MPSNNGradientFilterNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSNNGradientFilterNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNGradientFilterNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientFilterNode
type IMPSNNGradientFilterNode interface {
	IMPSNNFilterNode
}

// Init initializes the instance.
func (g MPSNNGradientFilterNode) Init() MPSNNGradientFilterNode {
	rv := objc.Send[MPSNNGradientFilterNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGradientFilterNode) Autorelease() MPSNNGradientFilterNode {
	rv := objc.Send[MPSNNGradientFilterNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGradientFilterNode creates a new MPSNNGradientFilterNode instance.
func NewMPSNNGradientFilterNode() MPSNNGradientFilterNode {
	class := getMPSNNGradientFilterNodeClass()
	rv := objc.Send[MPSNNGradientFilterNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

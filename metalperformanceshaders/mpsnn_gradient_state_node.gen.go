// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNGradientStateNode] class.
var (
	_MPSNNGradientStateNodeClass     MPSNNGradientStateNodeClass
	_MPSNNGradientStateNodeClassOnce sync.Once
)

func getMPSNNGradientStateNodeClass() MPSNNGradientStateNodeClass {
	_MPSNNGradientStateNodeClassOnce.Do(func() {
		_MPSNNGradientStateNodeClass = MPSNNGradientStateNodeClass{class: objc.GetClass("MPSNNGradientStateNode")}
	})
	return _MPSNNGradientStateNodeClass
}

// GetMPSNNGradientStateNodeClass returns the class object for MPSNNGradientStateNode.
func GetMPSNNGradientStateNodeClass() MPSNNGradientStateNodeClass {
	return getMPSNNGradientStateNodeClass()
}

type MPSNNGradientStateNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNGradientStateNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGradientStateNodeClass) Alloc() MPSNNGradientStateNode {
	rv := objc.Send[MPSNNGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of the state created to record the properties of a
// gradient kernel at the time it was encoded.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientStateNode
type MPSNNGradientStateNode struct {
	MPSNNStateNode
}

// MPSNNGradientStateNodeFromID constructs a [MPSNNGradientStateNode] from an objc.ID.
//
// A representation of the state created to record the properties of a
// gradient kernel at the time it was encoded.
func MPSNNGradientStateNodeFromID(id objc.ID) MPSNNGradientStateNode {
	return MPSNNGradientStateNode{MPSNNStateNode: MPSNNStateNodeFromID(id)}
}

// NOTE: MPSNNGradientStateNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNGradientStateNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientStateNode
type IMPSNNGradientStateNode interface {
	IMPSNNStateNode
}

// Init initializes the instance.
func (g MPSNNGradientStateNode) Init() MPSNNGradientStateNode {
	rv := objc.Send[MPSNNGradientStateNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGradientStateNode) Autorelease() MPSNNGradientStateNode {
	rv := objc.Send[MPSNNGradientStateNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGradientStateNode creates a new MPSNNGradientStateNode instance.
func NewMPSNNGradientStateNode() MPSNNGradientStateNode {
	class := getMPSNNGradientStateNodeClass()
	rv := objc.Send[MPSNNGradientStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

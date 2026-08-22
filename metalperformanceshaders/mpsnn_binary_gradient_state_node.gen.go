// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNBinaryGradientStateNode] class.
var (
	_MPSNNBinaryGradientStateNodeClass     MPSNNBinaryGradientStateNodeClass
	_MPSNNBinaryGradientStateNodeClassOnce sync.Once
)

func getMPSNNBinaryGradientStateNodeClass() MPSNNBinaryGradientStateNodeClass {
	_MPSNNBinaryGradientStateNodeClassOnce.Do(func() {
		_MPSNNBinaryGradientStateNodeClass = MPSNNBinaryGradientStateNodeClass{class: objc.GetClass("MPSNNBinaryGradientStateNode")}
	})
	return _MPSNNBinaryGradientStateNodeClass
}

// GetMPSNNBinaryGradientStateNodeClass returns the class object for MPSNNBinaryGradientStateNode.
func GetMPSNNBinaryGradientStateNodeClass() MPSNNBinaryGradientStateNodeClass {
	return getMPSNNBinaryGradientStateNodeClass()
}

type MPSNNBinaryGradientStateNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNBinaryGradientStateNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNBinaryGradientStateNodeClass) Alloc() MPSNNBinaryGradientStateNode {
	rv := objc.Send[MPSNNBinaryGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of the state created to record the properties of a binary
// gradient kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryGradientStateNode
type MPSNNBinaryGradientStateNode struct {
	MPSNNStateNode
}

// MPSNNBinaryGradientStateNodeFromID constructs a [MPSNNBinaryGradientStateNode] from an objc.ID.
//
// A representation of the state created to record the properties of a binary
// gradient kernel.
func MPSNNBinaryGradientStateNodeFromID(id objc.ID) MPSNNBinaryGradientStateNode {
	return MPSNNBinaryGradientStateNode{MPSNNStateNode: MPSNNStateNodeFromID(id)}
}

// NOTE: MPSNNBinaryGradientStateNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNBinaryGradientStateNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryGradientStateNode
type IMPSNNBinaryGradientStateNode interface {
	IMPSNNStateNode
}

// Init initializes the instance.
func (b MPSNNBinaryGradientStateNode) Init() MPSNNBinaryGradientStateNode {
	rv := objc.Send[MPSNNBinaryGradientStateNode](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MPSNNBinaryGradientStateNode) Autorelease() MPSNNBinaryGradientStateNode {
	rv := objc.Send[MPSNNBinaryGradientStateNode](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNBinaryGradientStateNode creates a new MPSNNBinaryGradientStateNode instance.
func NewMPSNNBinaryGradientStateNode() MPSNNBinaryGradientStateNode {
	class := getMPSNNBinaryGradientStateNodeClass()
	rv := objc.Send[MPSNNBinaryGradientStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

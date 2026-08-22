// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNArithmeticGradientStateNode] class.
var (
	_MPSNNArithmeticGradientStateNodeClass     MPSNNArithmeticGradientStateNodeClass
	_MPSNNArithmeticGradientStateNodeClassOnce sync.Once
)

func getMPSNNArithmeticGradientStateNodeClass() MPSNNArithmeticGradientStateNodeClass {
	_MPSNNArithmeticGradientStateNodeClassOnce.Do(func() {
		_MPSNNArithmeticGradientStateNodeClass = MPSNNArithmeticGradientStateNodeClass{class: objc.GetClass("MPSNNArithmeticGradientStateNode")}
	})
	return _MPSNNArithmeticGradientStateNodeClass
}

// GetMPSNNArithmeticGradientStateNodeClass returns the class object for MPSNNArithmeticGradientStateNode.
func GetMPSNNArithmeticGradientStateNodeClass() MPSNNArithmeticGradientStateNodeClass {
	return getMPSNNArithmeticGradientStateNodeClass()
}

type MPSNNArithmeticGradientStateNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNArithmeticGradientStateNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNArithmeticGradientStateNodeClass) Alloc() MPSNNArithmeticGradientStateNode {
	rv := objc.Send[MPSNNArithmeticGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of the clamp mask used by gradient arithmetic operators.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientStateNode
type MPSNNArithmeticGradientStateNode struct {
	MPSNNBinaryGradientStateNode
}

// MPSNNArithmeticGradientStateNodeFromID constructs a [MPSNNArithmeticGradientStateNode] from an objc.ID.
//
// A representation of the clamp mask used by gradient arithmetic operators.
func MPSNNArithmeticGradientStateNodeFromID(id objc.ID) MPSNNArithmeticGradientStateNode {
	return MPSNNArithmeticGradientStateNode{MPSNNBinaryGradientStateNode: MPSNNBinaryGradientStateNodeFromID(id)}
}

// NOTE: MPSNNArithmeticGradientStateNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNArithmeticGradientStateNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientStateNode
type IMPSNNArithmeticGradientStateNode interface {
	IMPSNNBinaryGradientStateNode
}

// Init initializes the instance.
func (a MPSNNArithmeticGradientStateNode) Init() MPSNNArithmeticGradientStateNode {
	rv := objc.Send[MPSNNArithmeticGradientStateNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MPSNNArithmeticGradientStateNode) Autorelease() MPSNNArithmeticGradientStateNode {
	rv := objc.Send[MPSNNArithmeticGradientStateNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNArithmeticGradientStateNode creates a new MPSNNArithmeticGradientStateNode instance.
func NewMPSNNArithmeticGradientStateNode() MPSNNArithmeticGradientStateNode {
	class := getMPSNNArithmeticGradientStateNodeClass()
	rv := objc.Send[MPSNNArithmeticGradientStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

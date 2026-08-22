// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNConvolutionGradientStateNode] class.
var (
	_MPSCNNConvolutionGradientStateNodeClass     MPSCNNConvolutionGradientStateNodeClass
	_MPSCNNConvolutionGradientStateNodeClassOnce sync.Once
)

func getMPSCNNConvolutionGradientStateNodeClass() MPSCNNConvolutionGradientStateNodeClass {
	_MPSCNNConvolutionGradientStateNodeClassOnce.Do(func() {
		_MPSCNNConvolutionGradientStateNodeClass = MPSCNNConvolutionGradientStateNodeClass{class: objc.GetClass("MPSCNNConvolutionGradientStateNode")}
	})
	return _MPSCNNConvolutionGradientStateNodeClass
}

// GetMPSCNNConvolutionGradientStateNodeClass returns the class object for MPSCNNConvolutionGradientStateNode.
func GetMPSCNNConvolutionGradientStateNodeClass() MPSCNNConvolutionGradientStateNodeClass {
	return getMPSCNNConvolutionGradientStateNodeClass()
}

type MPSCNNConvolutionGradientStateNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionGradientStateNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionGradientStateNodeClass) Alloc() MPSCNNConvolutionGradientStateNode {
	rv := objc.Send[MPSCNNConvolutionGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient convolution state.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientStateNode
type MPSCNNConvolutionGradientStateNode struct {
	MPSNNGradientStateNode
}

// MPSCNNConvolutionGradientStateNodeFromID constructs a [MPSCNNConvolutionGradientStateNode] from an objc.ID.
//
// A representation of a gradient convolution state.
func MPSCNNConvolutionGradientStateNodeFromID(id objc.ID) MPSCNNConvolutionGradientStateNode {
	return MPSCNNConvolutionGradientStateNode{MPSNNGradientStateNode: MPSNNGradientStateNodeFromID(id)}
}

// NOTE: MPSCNNConvolutionGradientStateNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionGradientStateNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientStateNode
type IMPSCNNConvolutionGradientStateNode interface {
	IMPSNNGradientStateNode
}

// Init initializes the instance.
func (c MPSCNNConvolutionGradientStateNode) Init() MPSCNNConvolutionGradientStateNode {
	rv := objc.Send[MPSCNNConvolutionGradientStateNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionGradientStateNode) Autorelease() MPSCNNConvolutionGradientStateNode {
	rv := objc.Send[MPSCNNConvolutionGradientStateNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionGradientStateNode creates a new MPSCNNConvolutionGradientStateNode instance.
func NewMPSCNNConvolutionGradientStateNode() MPSCNNConvolutionGradientStateNode {
	class := getMPSCNNConvolutionGradientStateNodeClass()
	rv := objc.Send[MPSCNNConvolutionGradientStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

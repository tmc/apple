// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNConvolutionTransposeGradientStateNode] class.
var (
	_MPSCNNConvolutionTransposeGradientStateNodeClass     MPSCNNConvolutionTransposeGradientStateNodeClass
	_MPSCNNConvolutionTransposeGradientStateNodeClassOnce sync.Once
)

func getMPSCNNConvolutionTransposeGradientStateNodeClass() MPSCNNConvolutionTransposeGradientStateNodeClass {
	_MPSCNNConvolutionTransposeGradientStateNodeClassOnce.Do(func() {
		_MPSCNNConvolutionTransposeGradientStateNodeClass = MPSCNNConvolutionTransposeGradientStateNodeClass{class: objc.GetClass("MPSCNNConvolutionTransposeGradientStateNode")}
	})
	return _MPSCNNConvolutionTransposeGradientStateNodeClass
}

// GetMPSCNNConvolutionTransposeGradientStateNodeClass returns the class object for MPSCNNConvolutionTransposeGradientStateNode.
func GetMPSCNNConvolutionTransposeGradientStateNodeClass() MPSCNNConvolutionTransposeGradientStateNodeClass {
	return getMPSCNNConvolutionTransposeGradientStateNodeClass()
}

type MPSCNNConvolutionTransposeGradientStateNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionTransposeGradientStateNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionTransposeGradientStateNodeClass) Alloc() MPSCNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[MPSCNNConvolutionTransposeGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientStateNode
type MPSCNNConvolutionTransposeGradientStateNode struct {
	MPSCNNConvolutionGradientStateNode
}

// MPSCNNConvolutionTransposeGradientStateNodeFromID constructs a [MPSCNNConvolutionTransposeGradientStateNode] from an objc.ID.
func MPSCNNConvolutionTransposeGradientStateNodeFromID(id objc.ID) MPSCNNConvolutionTransposeGradientStateNode {
	return MPSCNNConvolutionTransposeGradientStateNode{MPSCNNConvolutionGradientStateNode: MPSCNNConvolutionGradientStateNodeFromID(id)}
}

// NOTE: MPSCNNConvolutionTransposeGradientStateNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionTransposeGradientStateNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientStateNode
type IMPSCNNConvolutionTransposeGradientStateNode interface {
	IMPSCNNConvolutionGradientStateNode
}

// Init initializes the instance.
func (c MPSCNNConvolutionTransposeGradientStateNode) Init() MPSCNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[MPSCNNConvolutionTransposeGradientStateNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionTransposeGradientStateNode) Autorelease() MPSCNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[MPSCNNConvolutionTransposeGradientStateNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionTransposeGradientStateNode creates a new MPSCNNConvolutionTransposeGradientStateNode instance.
func NewMPSCNNConvolutionTransposeGradientStateNode() MPSCNNConvolutionTransposeGradientStateNode {
	class := getMPSCNNConvolutionTransposeGradientStateNodeClass()
	rv := objc.Send[MPSCNNConvolutionTransposeGradientStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

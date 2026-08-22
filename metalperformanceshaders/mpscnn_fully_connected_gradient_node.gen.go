// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNFullyConnectedGradientNode] class.
var (
	_MPSCNNFullyConnectedGradientNodeClass     MPSCNNFullyConnectedGradientNodeClass
	_MPSCNNFullyConnectedGradientNodeClassOnce sync.Once
)

func getMPSCNNFullyConnectedGradientNodeClass() MPSCNNFullyConnectedGradientNodeClass {
	_MPSCNNFullyConnectedGradientNodeClassOnce.Do(func() {
		_MPSCNNFullyConnectedGradientNodeClass = MPSCNNFullyConnectedGradientNodeClass{class: objc.GetClass("MPSCNNFullyConnectedGradientNode")}
	})
	return _MPSCNNFullyConnectedGradientNodeClass
}

// GetMPSCNNFullyConnectedGradientNodeClass returns the class object for MPSCNNFullyConnectedGradientNode.
func GetMPSCNNFullyConnectedGradientNodeClass() MPSCNNFullyConnectedGradientNodeClass {
	return getMPSCNNFullyConnectedGradientNodeClass()
}

type MPSCNNFullyConnectedGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNFullyConnectedGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNFullyConnectedGradientNodeClass) Alloc() MPSCNNFullyConnectedGradientNode {
	rv := objc.Send[MPSCNNFullyConnectedGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradientNode
type MPSCNNFullyConnectedGradientNode struct {
	MPSCNNConvolutionGradientNode
}

// MPSCNNFullyConnectedGradientNodeFromID constructs a [MPSCNNFullyConnectedGradientNode] from an objc.ID.
func MPSCNNFullyConnectedGradientNodeFromID(id objc.ID) MPSCNNFullyConnectedGradientNode {
	return MPSCNNFullyConnectedGradientNode{MPSCNNConvolutionGradientNode: MPSCNNConvolutionGradientNodeFromID(id)}
}

// NOTE: MPSCNNFullyConnectedGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNFullyConnectedGradientNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradientNode
type IMPSCNNFullyConnectedGradientNode interface {
	IMPSCNNConvolutionGradientNode
}

// Init initializes the instance.
func (c MPSCNNFullyConnectedGradientNode) Init() MPSCNNFullyConnectedGradientNode {
	rv := objc.Send[MPSCNNFullyConnectedGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNFullyConnectedGradientNode) Autorelease() MPSCNNFullyConnectedGradientNode {
	rv := objc.Send[MPSCNNFullyConnectedGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNFullyConnectedGradientNode creates a new MPSCNNFullyConnectedGradientNode instance.
func NewMPSCNNFullyConnectedGradientNode() MPSCNNFullyConnectedGradientNode {
	class := getMPSCNNFullyConnectedGradientNodeClass()
	rv := objc.Send[MPSCNNFullyConnectedGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradientNode/init(sourceGradient:sourceImage:convolutionGradientState:weights:)
func NewCNNFullyConnectedGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNFullyConnectedGradientNode {
	instance := getMPSCNNFullyConnectedGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return MPSCNNFullyConnectedGradientNodeFromID(rv)
}

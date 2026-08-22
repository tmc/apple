// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNFullyConnectedNode] class.
var (
	_MPSCNNFullyConnectedNodeClass     MPSCNNFullyConnectedNodeClass
	_MPSCNNFullyConnectedNodeClassOnce sync.Once
)

func getMPSCNNFullyConnectedNodeClass() MPSCNNFullyConnectedNodeClass {
	_MPSCNNFullyConnectedNodeClassOnce.Do(func() {
		_MPSCNNFullyConnectedNodeClass = MPSCNNFullyConnectedNodeClass{class: objc.GetClass("MPSCNNFullyConnectedNode")}
	})
	return _MPSCNNFullyConnectedNodeClass
}

// GetMPSCNNFullyConnectedNodeClass returns the class object for MPSCNNFullyConnectedNode.
func GetMPSCNNFullyConnectedNodeClass() MPSCNNFullyConnectedNodeClass {
	return getMPSCNNFullyConnectedNodeClass()
}

type MPSCNNFullyConnectedNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNFullyConnectedNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNFullyConnectedNodeClass) Alloc() MPSCNNFullyConnectedNode {
	rv := objc.Send[MPSCNNFullyConnectedNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a fully connected convolution layer, also known as an
// inner product layer.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedNode
type MPSCNNFullyConnectedNode struct {
	MPSCNNConvolutionNode
}

// MPSCNNFullyConnectedNodeFromID constructs a [MPSCNNFullyConnectedNode] from an objc.ID.
//
// A representation of a fully connected convolution layer, also known as an
// inner product layer.
func MPSCNNFullyConnectedNodeFromID(id objc.ID) MPSCNNFullyConnectedNode {
	return MPSCNNFullyConnectedNode{MPSCNNConvolutionNode: MPSCNNConvolutionNodeFromID(id)}
}

// NOTE: MPSCNNFullyConnectedNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNFullyConnectedNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedNode
type IMPSCNNFullyConnectedNode interface {
	IMPSCNNConvolutionNode
}

// Init initializes the instance.
func (c MPSCNNFullyConnectedNode) Init() MPSCNNFullyConnectedNode {
	rv := objc.Send[MPSCNNFullyConnectedNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNFullyConnectedNode) Autorelease() MPSCNNFullyConnectedNode {
	rv := objc.Send[MPSCNNFullyConnectedNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNFullyConnectedNode creates a new MPSCNNFullyConnectedNode instance.
func NewMPSCNNFullyConnectedNode() MPSCNNFullyConnectedNode {
	class := getMPSCNNFullyConnectedNodeClass()
	rv := objc.Send[MPSCNNFullyConnectedNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedNode/init(source:weights:)
func NewCNNFullyConnectedNodeWithSourceWeights(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource) MPSCNNFullyConnectedNode {
	instance := getMPSCNNFullyConnectedNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:weights:"), sourceNode, weights)
	return MPSCNNFullyConnectedNodeFromID(rv)
}

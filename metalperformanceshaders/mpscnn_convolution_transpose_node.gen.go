// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNConvolutionTransposeNode] class.
var (
	_MPSCNNConvolutionTransposeNodeClass     MPSCNNConvolutionTransposeNodeClass
	_MPSCNNConvolutionTransposeNodeClassOnce sync.Once
)

func getMPSCNNConvolutionTransposeNodeClass() MPSCNNConvolutionTransposeNodeClass {
	_MPSCNNConvolutionTransposeNodeClassOnce.Do(func() {
		_MPSCNNConvolutionTransposeNodeClass = MPSCNNConvolutionTransposeNodeClass{class: objc.GetClass("MPSCNNConvolutionTransposeNode")}
	})
	return _MPSCNNConvolutionTransposeNodeClass
}

// GetMPSCNNConvolutionTransposeNodeClass returns the class object for MPSCNNConvolutionTransposeNode.
func GetMPSCNNConvolutionTransposeNodeClass() MPSCNNConvolutionTransposeNodeClass {
	return getMPSCNNConvolutionTransposeNodeClass()
}

type MPSCNNConvolutionTransposeNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionTransposeNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionTransposeNodeClass) Alloc() MPSCNNConvolutionTransposeNode {
	rv := objc.Send[MPSCNNConvolutionTransposeNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a transposed convolution.
//
// # Initializers
//
//   - [MPSCNNConvolutionTransposeNode.InitWithSourceConvolutionGradientStateWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeNode
type MPSCNNConvolutionTransposeNode struct {
	MPSCNNConvolutionNode
}

// MPSCNNConvolutionTransposeNodeFromID constructs a [MPSCNNConvolutionTransposeNode] from an objc.ID.
//
// A representation of a transposed convolution.
func MPSCNNConvolutionTransposeNodeFromID(id objc.ID) MPSCNNConvolutionTransposeNode {
	return MPSCNNConvolutionTransposeNode{MPSCNNConvolutionNode: MPSCNNConvolutionNodeFromID(id)}
}

// NOTE: MPSCNNConvolutionTransposeNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionTransposeNode] class.
//
// # Initializers
//
//   - [IMPSCNNConvolutionTransposeNode.InitWithSourceConvolutionGradientStateWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeNode
type IMPSCNNConvolutionTransposeNode interface {
	IMPSCNNConvolutionNode

	// Topic: Initializers

	InitWithSourceConvolutionGradientStateWeights(sourceNode IMPSNNImageNode, convolutionGradientState IMPSCNNConvolutionGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeNode
}

// Init initializes the instance.
func (c MPSCNNConvolutionTransposeNode) Init() MPSCNNConvolutionTransposeNode {
	rv := objc.Send[MPSCNNConvolutionTransposeNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionTransposeNode) Autorelease() MPSCNNConvolutionTransposeNode {
	rv := objc.Send[MPSCNNConvolutionTransposeNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionTransposeNode creates a new MPSCNNConvolutionTransposeNode instance.
func NewMPSCNNConvolutionTransposeNode() MPSCNNConvolutionTransposeNode {
	class := getMPSCNNConvolutionTransposeNodeClass()
	rv := objc.Send[MPSCNNConvolutionTransposeNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeNode/init(source:convolutionGradientState:weights:)
func NewCNNConvolutionTransposeNodeWithSourceConvolutionGradientStateWeights(sourceNode IMPSNNImageNode, convolutionGradientState IMPSCNNConvolutionGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeNode {
	instance := getMPSCNNConvolutionTransposeNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:convolutionGradientState:weights:"), sourceNode, convolutionGradientState, weights)
	return MPSCNNConvolutionTransposeNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode/init(source:weights:)
func NewCNNConvolutionTransposeNodeWithSourceWeights(sourceNode IMPSNNImageNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeNode {
	instance := getMPSCNNConvolutionTransposeNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:weights:"), sourceNode, weights)
	return MPSCNNConvolutionTransposeNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeNode/init(source:convolutionGradientState:weights:)
func (c MPSCNNConvolutionTransposeNode) InitWithSourceConvolutionGradientStateWeights(sourceNode IMPSNNImageNode, convolutionGradientState IMPSCNNConvolutionGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeNode {
	rv := objc.Send[MPSCNNConvolutionTransposeNode](c.ID, objc.Sel("initWithSource:convolutionGradientState:weights:"), sourceNode, convolutionGradientState, weights)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeNode/nodeWithSource:convolutionGradientState:weights:
func (_MPSCNNConvolutionTransposeNodeClass MPSCNNConvolutionTransposeNodeClass) NodeWithSourceConvolutionGradientStateWeights(sourceNode IMPSNNImageNode, convolutionGradientState IMPSCNNConvolutionGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNConvolutionTransposeNodeClass.class), objc.Sel("nodeWithSource:convolutionGradientState:weights:"), sourceNode, convolutionGradientState, weights)
	return MPSCNNConvolutionTransposeNodeFromID(rv)
}

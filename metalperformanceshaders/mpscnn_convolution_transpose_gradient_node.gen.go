// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNConvolutionTransposeGradientNode] class.
var (
	_MPSCNNConvolutionTransposeGradientNodeClass     MPSCNNConvolutionTransposeGradientNodeClass
	_MPSCNNConvolutionTransposeGradientNodeClassOnce sync.Once
)

func getMPSCNNConvolutionTransposeGradientNodeClass() MPSCNNConvolutionTransposeGradientNodeClass {
	_MPSCNNConvolutionTransposeGradientNodeClassOnce.Do(func() {
		_MPSCNNConvolutionTransposeGradientNodeClass = MPSCNNConvolutionTransposeGradientNodeClass{class: objc.GetClass("MPSCNNConvolutionTransposeGradientNode")}
	})
	return _MPSCNNConvolutionTransposeGradientNodeClass
}

// GetMPSCNNConvolutionTransposeGradientNodeClass returns the class object for MPSCNNConvolutionTransposeGradientNode.
func GetMPSCNNConvolutionTransposeGradientNodeClass() MPSCNNConvolutionTransposeGradientNodeClass {
	return getMPSCNNConvolutionTransposeGradientNodeClass()
}

type MPSCNNConvolutionTransposeGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionTransposeGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionTransposeGradientNodeClass) Alloc() MPSCNNConvolutionTransposeGradientNode {
	rv := objc.Send[MPSCNNConvolutionTransposeGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSCNNConvolutionTransposeGradientNode.InitWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientNode
type MPSCNNConvolutionTransposeGradientNode struct {
	MPSCNNConvolutionGradientNode
}

// MPSCNNConvolutionTransposeGradientNodeFromID constructs a [MPSCNNConvolutionTransposeGradientNode] from an objc.ID.
func MPSCNNConvolutionTransposeGradientNodeFromID(id objc.ID) MPSCNNConvolutionTransposeGradientNode {
	return MPSCNNConvolutionTransposeGradientNode{MPSCNNConvolutionGradientNode: MPSCNNConvolutionGradientNodeFromID(id)}
}

// NOTE: MPSCNNConvolutionTransposeGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionTransposeGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNConvolutionTransposeGradientNode.InitWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientNode
type IMPSCNNConvolutionTransposeGradientNode interface {
	IMPSCNNConvolutionGradientNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionTransposeGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeGradientNode
}

// Init initializes the instance.
func (c MPSCNNConvolutionTransposeGradientNode) Init() MPSCNNConvolutionTransposeGradientNode {
	rv := objc.Send[MPSCNNConvolutionTransposeGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionTransposeGradientNode) Autorelease() MPSCNNConvolutionTransposeGradientNode {
	rv := objc.Send[MPSCNNConvolutionTransposeGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionTransposeGradientNode creates a new MPSCNNConvolutionTransposeGradientNode instance.
func NewMPSCNNConvolutionTransposeGradientNode() MPSCNNConvolutionTransposeGradientNode {
	class := getMPSCNNConvolutionTransposeGradientNodeClass()
	rv := objc.Send[MPSCNNConvolutionTransposeGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientNode/init(sourceGradient:sourceImage:convolutionGradientState:weights:)
func NewCNNConvolutionTransposeGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeGradientNode {
	instance := getMPSCNNConvolutionTransposeGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return MPSCNNConvolutionTransposeGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientNode/init(sourceGradient:sourceImage:convolutionTransposeGradientState:weights:)
func NewCNNConvolutionTransposeGradientNodeWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionTransposeGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeGradientNode {
	instance := getMPSCNNConvolutionTransposeGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionTransposeGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return MPSCNNConvolutionTransposeGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientNode/init(sourceGradient:sourceImage:convolutionTransposeGradientState:weights:)
func (c MPSCNNConvolutionTransposeGradientNode) InitWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionTransposeGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeGradientNode {
	rv := objc.Send[MPSCNNConvolutionTransposeGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:convolutionTransposeGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientNode/nodeWithSourceGradient:sourceImage:convolutionTransposeGradientState:weights:
func (_MPSCNNConvolutionTransposeGradientNodeClass MPSCNNConvolutionTransposeGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSCNNConvolutionTransposeGradientStateNode, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTransposeGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNConvolutionTransposeGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:convolutionTransposeGradientState:weights:"), sourceGradient, sourceImage, gradientState, weights)
	return MPSCNNConvolutionTransposeGradientNodeFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNGramMatrixCalculationGradientNode] class.
var (
	_MPSNNGramMatrixCalculationGradientNodeClass     MPSNNGramMatrixCalculationGradientNodeClass
	_MPSNNGramMatrixCalculationGradientNodeClassOnce sync.Once
)

func getMPSNNGramMatrixCalculationGradientNodeClass() MPSNNGramMatrixCalculationGradientNodeClass {
	_MPSNNGramMatrixCalculationGradientNodeClassOnce.Do(func() {
		_MPSNNGramMatrixCalculationGradientNodeClass = MPSNNGramMatrixCalculationGradientNodeClass{class: objc.GetClass("MPSNNGramMatrixCalculationGradientNode")}
	})
	return _MPSNNGramMatrixCalculationGradientNodeClass
}

// GetMPSNNGramMatrixCalculationGradientNodeClass returns the class object for MPSNNGramMatrixCalculationGradientNode.
func GetMPSNNGramMatrixCalculationGradientNodeClass() MPSNNGramMatrixCalculationGradientNodeClass {
	return getMPSNNGramMatrixCalculationGradientNodeClass()
}

type MPSNNGramMatrixCalculationGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNGramMatrixCalculationGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGramMatrixCalculationGradientNodeClass) Alloc() MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNGramMatrixCalculationGradientNode.InitWithSourceGradientSourceImageGradientState]
//   - [MPSNNGramMatrixCalculationGradientNode.InitWithSourceGradientSourceImageGradientStateAlpha]
//
// # Instance Properties
//
//   - [MPSNNGramMatrixCalculationGradientNode.Alpha]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode
type MPSNNGramMatrixCalculationGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNGramMatrixCalculationGradientNodeFromID constructs a [MPSNNGramMatrixCalculationGradientNode] from an objc.ID.
func MPSNNGramMatrixCalculationGradientNodeFromID(id objc.ID) MPSNNGramMatrixCalculationGradientNode {
	return MPSNNGramMatrixCalculationGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSNNGramMatrixCalculationGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNGramMatrixCalculationGradientNode] class.
//
// # Initializers
//
//   - [IMPSNNGramMatrixCalculationGradientNode.InitWithSourceGradientSourceImageGradientState]
//   - [IMPSNNGramMatrixCalculationGradientNode.InitWithSourceGradientSourceImageGradientStateAlpha]
//
// # Instance Properties
//
//   - [IMPSNNGramMatrixCalculationGradientNode.Alpha]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode
type IMPSNNGramMatrixCalculationGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNGramMatrixCalculationGradientNode
	InitWithSourceGradientSourceImageGradientStateAlpha(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, alpha float32) MPSNNGramMatrixCalculationGradientNode

	// Topic: Instance Properties

	Alpha() float32
}

// Init initializes the instance.
func (g MPSNNGramMatrixCalculationGradientNode) Init() MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGramMatrixCalculationGradientNode) Autorelease() MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGramMatrixCalculationGradientNode creates a new MPSNNGramMatrixCalculationGradientNode instance.
func NewMPSNNGramMatrixCalculationGradientNode() MPSNNGramMatrixCalculationGradientNode {
	class := getMPSNNGramMatrixCalculationGradientNodeClass()
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode/init(sourceGradient:sourceImage:gradientState:)
func NewGramMatrixCalculationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNGramMatrixCalculationGradientNode {
	instance := getMPSNNGramMatrixCalculationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSNNGramMatrixCalculationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode/init(sourceGradient:sourceImage:gradientState:alpha:)
func NewGramMatrixCalculationGradientNodeWithSourceGradientSourceImageGradientStateAlpha(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, alpha float32) MPSNNGramMatrixCalculationGradientNode {
	instance := getMPSNNGramMatrixCalculationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:alpha:"), sourceGradient, sourceImage, gradientState, alpha)
	return MPSNNGramMatrixCalculationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode/init(sourceGradient:sourceImage:gradientState:)
func (g MPSNNGramMatrixCalculationGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](g.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode/init(sourceGradient:sourceImage:gradientState:alpha:)
func (g MPSNNGramMatrixCalculationGradientNode) InitWithSourceGradientSourceImageGradientStateAlpha(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, alpha float32) MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[MPSNNGramMatrixCalculationGradientNode](g.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:alpha:"), sourceGradient, sourceImage, gradientState, alpha)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:
func (_MPSNNGramMatrixCalculationGradientNodeClass MPSNNGramMatrixCalculationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode) MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGramMatrixCalculationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return MPSNNGramMatrixCalculationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:alpha:
func (_MPSNNGramMatrixCalculationGradientNodeClass MPSNNGramMatrixCalculationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateAlpha(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, alpha float32) MPSNNGramMatrixCalculationGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGramMatrixCalculationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:alpha:"), sourceGradient, sourceImage, gradientState, alpha)
	return MPSNNGramMatrixCalculationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode/alpha
func (g MPSNNGramMatrixCalculationGradientNode) Alpha() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("alpha"))
	return rv
}

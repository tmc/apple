// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNGramMatrixCalculationNode] class.
var (
	_MPSNNGramMatrixCalculationNodeClass     MPSNNGramMatrixCalculationNodeClass
	_MPSNNGramMatrixCalculationNodeClassOnce sync.Once
)

func getMPSNNGramMatrixCalculationNodeClass() MPSNNGramMatrixCalculationNodeClass {
	_MPSNNGramMatrixCalculationNodeClassOnce.Do(func() {
		_MPSNNGramMatrixCalculationNodeClass = MPSNNGramMatrixCalculationNodeClass{class: objc.GetClass("MPSNNGramMatrixCalculationNode")}
	})
	return _MPSNNGramMatrixCalculationNodeClass
}

// GetMPSNNGramMatrixCalculationNodeClass returns the class object for MPSNNGramMatrixCalculationNode.
func GetMPSNNGramMatrixCalculationNodeClass() MPSNNGramMatrixCalculationNodeClass {
	return getMPSNNGramMatrixCalculationNodeClass()
}

type MPSNNGramMatrixCalculationNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNGramMatrixCalculationNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGramMatrixCalculationNodeClass) Alloc() MPSNNGramMatrixCalculationNode {
	rv := objc.Send[MPSNNGramMatrixCalculationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNGramMatrixCalculationNode.InitWithSource]
//   - [MPSNNGramMatrixCalculationNode.InitWithSourceAlpha]
//
// # Instance Properties
//
//   - [MPSNNGramMatrixCalculationNode.Alpha]
//   - [MPSNNGramMatrixCalculationNode.PropertyCallBack]
//   - [MPSNNGramMatrixCalculationNode.SetPropertyCallBack]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode
type MPSNNGramMatrixCalculationNode struct {
	MPSNNFilterNode
}

// MPSNNGramMatrixCalculationNodeFromID constructs a [MPSNNGramMatrixCalculationNode] from an objc.ID.
func MPSNNGramMatrixCalculationNodeFromID(id objc.ID) MPSNNGramMatrixCalculationNode {
	return MPSNNGramMatrixCalculationNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSNNGramMatrixCalculationNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNGramMatrixCalculationNode] class.
//
// # Initializers
//
//   - [IMPSNNGramMatrixCalculationNode.InitWithSource]
//   - [IMPSNNGramMatrixCalculationNode.InitWithSourceAlpha]
//
// # Instance Properties
//
//   - [IMPSNNGramMatrixCalculationNode.Alpha]
//   - [IMPSNNGramMatrixCalculationNode.PropertyCallBack]
//   - [IMPSNNGramMatrixCalculationNode.SetPropertyCallBack]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode
type IMPSNNGramMatrixCalculationNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSNNGramMatrixCalculationNode
	InitWithSourceAlpha(sourceNode IMPSNNImageNode, alpha float32) MPSNNGramMatrixCalculationNode

	// Topic: Instance Properties

	Alpha() float32
	PropertyCallBack() MPSNNGramMatrixCallback
	SetPropertyCallBack(value MPSNNGramMatrixCallback)
}

// Init initializes the instance.
func (g MPSNNGramMatrixCalculationNode) Init() MPSNNGramMatrixCalculationNode {
	rv := objc.Send[MPSNNGramMatrixCalculationNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGramMatrixCalculationNode) Autorelease() MPSNNGramMatrixCalculationNode {
	rv := objc.Send[MPSNNGramMatrixCalculationNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGramMatrixCalculationNode creates a new MPSNNGramMatrixCalculationNode instance.
func NewMPSNNGramMatrixCalculationNode() MPSNNGramMatrixCalculationNode {
	class := getMPSNNGramMatrixCalculationNodeClass()
	rv := objc.Send[MPSNNGramMatrixCalculationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/init(source:)
func NewGramMatrixCalculationNodeWithSource(sourceNode IMPSNNImageNode) MPSNNGramMatrixCalculationNode {
	instance := getMPSNNGramMatrixCalculationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSNNGramMatrixCalculationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/init(source:alpha:)
func NewGramMatrixCalculationNodeWithSourceAlpha(sourceNode IMPSNNImageNode, alpha float32) MPSNNGramMatrixCalculationNode {
	instance := getMPSNNGramMatrixCalculationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:alpha:"), sourceNode, alpha)
	return MPSNNGramMatrixCalculationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/init(source:)
func (g MPSNNGramMatrixCalculationNode) InitWithSource(sourceNode IMPSNNImageNode) MPSNNGramMatrixCalculationNode {
	rv := objc.Send[MPSNNGramMatrixCalculationNode](g.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/init(source:alpha:)
func (g MPSNNGramMatrixCalculationNode) InitWithSourceAlpha(sourceNode IMPSNNImageNode, alpha float32) MPSNNGramMatrixCalculationNode {
	rv := objc.Send[MPSNNGramMatrixCalculationNode](g.ID, objc.Sel("initWithSource:alpha:"), sourceNode, alpha)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/nodeWithSource:
func (_MPSNNGramMatrixCalculationNodeClass MPSNNGramMatrixCalculationNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSNNGramMatrixCalculationNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGramMatrixCalculationNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSNNGramMatrixCalculationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/nodeWithSource:alpha:
func (_MPSNNGramMatrixCalculationNodeClass MPSNNGramMatrixCalculationNodeClass) NodeWithSourceAlpha(sourceNode IMPSNNImageNode, alpha float32) MPSNNGramMatrixCalculationNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGramMatrixCalculationNodeClass.class), objc.Sel("nodeWithSource:alpha:"), sourceNode, alpha)
	return MPSNNGramMatrixCalculationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/alpha
func (g MPSNNGramMatrixCalculationNode) Alpha() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("alpha"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/propertyCallBack
func (g MPSNNGramMatrixCalculationNode) PropertyCallBack() MPSNNGramMatrixCallback {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("propertyCallBack"))
	return MPSNNGramMatrixCallbackObjectFromID(rv)
}
func (g MPSNNGramMatrixCalculationNode) SetPropertyCallBack(value MPSNNGramMatrixCallback) {
	objc.Send[struct{}](g.ID, objc.Sel("setPropertyCallBack:"), value)
}

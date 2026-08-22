// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNormalizationNode] class.
var (
	_MPSCNNNormalizationNodeClass     MPSCNNNormalizationNodeClass
	_MPSCNNNormalizationNodeClassOnce sync.Once
)

func getMPSCNNNormalizationNodeClass() MPSCNNNormalizationNodeClass {
	_MPSCNNNormalizationNodeClassOnce.Do(func() {
		_MPSCNNNormalizationNodeClass = MPSCNNNormalizationNodeClass{class: objc.GetClass("MPSCNNNormalizationNode")}
	})
	return _MPSCNNNormalizationNodeClass
}

// GetMPSCNNNormalizationNodeClass returns the class object for MPSCNNNormalizationNode.
func GetMPSCNNNormalizationNodeClass() MPSCNNNormalizationNodeClass {
	return getMPSCNNNormalizationNodeClass()
}

type MPSCNNNormalizationNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNormalizationNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNormalizationNodeClass) Alloc() MPSCNNNormalizationNode {
	rv := objc.Send[MPSCNNNormalizationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Virtual base class for CNN normalization nodes.
//
// # Initializers
//
//   - [MPSCNNNormalizationNode.InitWithSource]
//
// # Instance Properties
//
//   - [MPSCNNNormalizationNode.Alpha]
//   - [MPSCNNNormalizationNode.SetAlpha]
//   - [MPSCNNNormalizationNode.Beta]
//   - [MPSCNNNormalizationNode.SetBeta]
//   - [MPSCNNNormalizationNode.Delta]
//   - [MPSCNNNormalizationNode.SetDelta]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationNode
type MPSCNNNormalizationNode struct {
	MPSNNFilterNode
}

// MPSCNNNormalizationNodeFromID constructs a [MPSCNNNormalizationNode] from an objc.ID.
//
// Virtual base class for CNN normalization nodes.
func MPSCNNNormalizationNodeFromID(id objc.ID) MPSCNNNormalizationNode {
	return MPSCNNNormalizationNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNNormalizationNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNormalizationNode] class.
//
// # Initializers
//
//   - [IMPSCNNNormalizationNode.InitWithSource]
//
// # Instance Properties
//
//   - [IMPSCNNNormalizationNode.Alpha]
//   - [IMPSCNNNormalizationNode.SetAlpha]
//   - [IMPSCNNNormalizationNode.Beta]
//   - [IMPSCNNNormalizationNode.SetBeta]
//   - [IMPSCNNNormalizationNode.Delta]
//   - [IMPSCNNNormalizationNode.SetDelta]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationNode
type IMPSCNNNormalizationNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNormalizationNode

	// Topic: Instance Properties

	Alpha() float32
	SetAlpha(value float32)
	Beta() float32
	SetBeta(value float32)
	Delta() float32
	SetDelta(value float32)
}

// Init initializes the instance.
func (c MPSCNNNormalizationNode) Init() MPSCNNNormalizationNode {
	rv := objc.Send[MPSCNNNormalizationNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNormalizationNode) Autorelease() MPSCNNNormalizationNode {
	rv := objc.Send[MPSCNNNormalizationNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNormalizationNode creates a new MPSCNNNormalizationNode instance.
func NewMPSCNNNormalizationNode() MPSCNNNormalizationNode {
	class := getMPSCNNNormalizationNodeClass()
	rv := objc.Send[MPSCNNNormalizationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationNode/init(source:)
func NewCNNNormalizationNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNormalizationNode {
	instance := getMPSCNNNormalizationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationNode/init(source:)
func (c MPSCNNNormalizationNode) InitWithSource(sourceNode IMPSNNImageNode) MPSCNNNormalizationNode {
	rv := objc.Send[MPSCNNNormalizationNode](c.ID, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationNode/nodeWithSource:
func (_MPSCNNNormalizationNodeClass MPSCNNNormalizationNodeClass) NodeWithSource(sourceNode IMPSNNImageNode) MPSCNNNormalizationNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNNormalizationNodeClass.class), objc.Sel("nodeWithSource:"), sourceNode)
	return MPSCNNNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationNode/alpha
func (c MPSCNNNormalizationNode) Alpha() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("alpha"))
	return rv
}
func (c MPSCNNNormalizationNode) SetAlpha(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationNode/beta
func (c MPSCNNNormalizationNode) Beta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("beta"))
	return rv
}
func (c MPSCNNNormalizationNode) SetBeta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBeta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationNode/delta
func (c MPSCNNNormalizationNode) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}
func (c MPSCNNNormalizationNode) SetDelta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelta:"), value)
}

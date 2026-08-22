// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNUpsamplingNearestNode] class.
var (
	_MPSCNNUpsamplingNearestNodeClass     MPSCNNUpsamplingNearestNodeClass
	_MPSCNNUpsamplingNearestNodeClassOnce sync.Once
)

func getMPSCNNUpsamplingNearestNodeClass() MPSCNNUpsamplingNearestNodeClass {
	_MPSCNNUpsamplingNearestNodeClassOnce.Do(func() {
		_MPSCNNUpsamplingNearestNodeClass = MPSCNNUpsamplingNearestNodeClass{class: objc.GetClass("MPSCNNUpsamplingNearestNode")}
	})
	return _MPSCNNUpsamplingNearestNodeClass
}

// GetMPSCNNUpsamplingNearestNodeClass returns the class object for MPSCNNUpsamplingNearestNode.
func GetMPSCNNUpsamplingNearestNodeClass() MPSCNNUpsamplingNearestNodeClass {
	return getMPSCNNUpsamplingNearestNodeClass()
}

type MPSCNNUpsamplingNearestNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNUpsamplingNearestNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNUpsamplingNearestNodeClass) Alloc() MPSCNNUpsamplingNearestNode {
	rv := objc.Send[MPSCNNUpsamplingNearestNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a nearest spatial upsampling filter.
//
// # Initializers
//
//   - [MPSCNNUpsamplingNearestNode.InitWithSourceIntegerScaleFactorXIntegerScaleFactorY]
//
// # Instance Properties
//
//   - [MPSCNNUpsamplingNearestNode.ScaleFactorX]
//   - [MPSCNNUpsamplingNearestNode.ScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestNode
type MPSCNNUpsamplingNearestNode struct {
	MPSNNFilterNode
}

// MPSCNNUpsamplingNearestNodeFromID constructs a [MPSCNNUpsamplingNearestNode] from an objc.ID.
//
// A representation of a nearest spatial upsampling filter.
func MPSCNNUpsamplingNearestNodeFromID(id objc.ID) MPSCNNUpsamplingNearestNode {
	return MPSCNNUpsamplingNearestNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNUpsamplingNearestNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNUpsamplingNearestNode] class.
//
// # Initializers
//
//   - [IMPSCNNUpsamplingNearestNode.InitWithSourceIntegerScaleFactorXIntegerScaleFactorY]
//
// # Instance Properties
//
//   - [IMPSCNNUpsamplingNearestNode.ScaleFactorX]
//   - [IMPSCNNUpsamplingNearestNode.ScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestNode
type IMPSCNNUpsamplingNearestNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingNearestNode

	// Topic: Instance Properties

	ScaleFactorX() float64
	ScaleFactorY() float64
}

// Init initializes the instance.
func (c MPSCNNUpsamplingNearestNode) Init() MPSCNNUpsamplingNearestNode {
	rv := objc.Send[MPSCNNUpsamplingNearestNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNUpsamplingNearestNode) Autorelease() MPSCNNUpsamplingNearestNode {
	rv := objc.Send[MPSCNNUpsamplingNearestNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNUpsamplingNearestNode creates a new MPSCNNUpsamplingNearestNode instance.
func NewMPSCNNUpsamplingNearestNode() MPSCNNUpsamplingNearestNode {
	class := getMPSCNNUpsamplingNearestNodeClass()
	rv := objc.Send[MPSCNNUpsamplingNearestNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestNode/init(source:integerScaleFactorX:integerScaleFactorY:)
func NewCNNUpsamplingNearestNodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingNearestNode {
	instance := getMPSCNNUpsamplingNearestNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	return MPSCNNUpsamplingNearestNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestNode/init(source:integerScaleFactorX:integerScaleFactorY:)
func (c MPSCNNUpsamplingNearestNode) InitWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingNearestNode {
	rv := objc.Send[MPSCNNUpsamplingNearestNode](c.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestNode/nodeWithSource:integerScaleFactorX:integerScaleFactorY:
func (_MPSCNNUpsamplingNearestNodeClass MPSCNNUpsamplingNearestNodeClass) NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingNearestNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNUpsamplingNearestNodeClass.class), objc.Sel("nodeWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	return MPSCNNUpsamplingNearestNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestNode/scaleFactorX
func (c MPSCNNUpsamplingNearestNode) ScaleFactorX() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestNode/scaleFactorY
func (c MPSCNNUpsamplingNearestNode) ScaleFactorY() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorY"))
	return rv
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNUpsamplingBilinearNode] class.
var (
	_MPSCNNUpsamplingBilinearNodeClass     MPSCNNUpsamplingBilinearNodeClass
	_MPSCNNUpsamplingBilinearNodeClassOnce sync.Once
)

func getMPSCNNUpsamplingBilinearNodeClass() MPSCNNUpsamplingBilinearNodeClass {
	_MPSCNNUpsamplingBilinearNodeClassOnce.Do(func() {
		_MPSCNNUpsamplingBilinearNodeClass = MPSCNNUpsamplingBilinearNodeClass{class: objc.GetClass("MPSCNNUpsamplingBilinearNode")}
	})
	return _MPSCNNUpsamplingBilinearNodeClass
}

// GetMPSCNNUpsamplingBilinearNodeClass returns the class object for MPSCNNUpsamplingBilinearNode.
func GetMPSCNNUpsamplingBilinearNodeClass() MPSCNNUpsamplingBilinearNodeClass {
	return getMPSCNNUpsamplingBilinearNodeClass()
}

type MPSCNNUpsamplingBilinearNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNUpsamplingBilinearNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNUpsamplingBilinearNodeClass) Alloc() MPSCNNUpsamplingBilinearNode {
	rv := objc.Send[MPSCNNUpsamplingBilinearNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a bilinear spatial upsampling filter.
//
// # Initializers
//
//   - [MPSCNNUpsamplingBilinearNode.InitWithSourceIntegerScaleFactorXIntegerScaleFactorY]
//   - [MPSCNNUpsamplingBilinearNode.InitWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners]
//
// # Instance Properties
//
//   - [MPSCNNUpsamplingBilinearNode.ScaleFactorX]
//   - [MPSCNNUpsamplingBilinearNode.ScaleFactorY]
//   - [MPSCNNUpsamplingBilinearNode.AlignCorners]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode
type MPSCNNUpsamplingBilinearNode struct {
	MPSNNFilterNode
}

// MPSCNNUpsamplingBilinearNodeFromID constructs a [MPSCNNUpsamplingBilinearNode] from an objc.ID.
//
// A representation of a bilinear spatial upsampling filter.
func MPSCNNUpsamplingBilinearNodeFromID(id objc.ID) MPSCNNUpsamplingBilinearNode {
	return MPSCNNUpsamplingBilinearNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNUpsamplingBilinearNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNUpsamplingBilinearNode] class.
//
// # Initializers
//
//   - [IMPSCNNUpsamplingBilinearNode.InitWithSourceIntegerScaleFactorXIntegerScaleFactorY]
//   - [IMPSCNNUpsamplingBilinearNode.InitWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners]
//
// # Instance Properties
//
//   - [IMPSCNNUpsamplingBilinearNode.ScaleFactorX]
//   - [IMPSCNNUpsamplingBilinearNode.ScaleFactorY]
//   - [IMPSCNNUpsamplingBilinearNode.AlignCorners]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode
type IMPSCNNUpsamplingBilinearNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingBilinearNode
	InitWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) MPSCNNUpsamplingBilinearNode

	// Topic: Instance Properties

	ScaleFactorX() float64
	ScaleFactorY() float64
	AlignCorners() bool
}

// Init initializes the instance.
func (c MPSCNNUpsamplingBilinearNode) Init() MPSCNNUpsamplingBilinearNode {
	rv := objc.Send[MPSCNNUpsamplingBilinearNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNUpsamplingBilinearNode) Autorelease() MPSCNNUpsamplingBilinearNode {
	rv := objc.Send[MPSCNNUpsamplingBilinearNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNUpsamplingBilinearNode creates a new MPSCNNUpsamplingBilinearNode instance.
func NewMPSCNNUpsamplingBilinearNode() MPSCNNUpsamplingBilinearNode {
	class := getMPSCNNUpsamplingBilinearNodeClass()
	rv := objc.Send[MPSCNNUpsamplingBilinearNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode/init(source:integerScaleFactorX:integerScaleFactorY:)
func NewCNNUpsamplingBilinearNodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingBilinearNode {
	instance := getMPSCNNUpsamplingBilinearNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	return MPSCNNUpsamplingBilinearNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode/init(source:integerScaleFactorX:integerScaleFactorY:alignCorners:)
func NewCNNUpsamplingBilinearNodeWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) MPSCNNUpsamplingBilinearNode {
	instance := getMPSCNNUpsamplingBilinearNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:alignCorners:"), sourceNode, integerScaleFactorX, integerScaleFactorY, alignCorners)
	return MPSCNNUpsamplingBilinearNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode/init(source:integerScaleFactorX:integerScaleFactorY:)
func (c MPSCNNUpsamplingBilinearNode) InitWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingBilinearNode {
	rv := objc.Send[MPSCNNUpsamplingBilinearNode](c.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode/init(source:integerScaleFactorX:integerScaleFactorY:alignCorners:)
func (c MPSCNNUpsamplingBilinearNode) InitWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) MPSCNNUpsamplingBilinearNode {
	rv := objc.Send[MPSCNNUpsamplingBilinearNode](c.ID, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:alignCorners:"), sourceNode, integerScaleFactorX, integerScaleFactorY, alignCorners)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode/nodeWithSource:integerScaleFactorX:integerScaleFactorY:
func (_MPSCNNUpsamplingBilinearNodeClass MPSCNNUpsamplingBilinearNodeClass) NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) MPSCNNUpsamplingBilinearNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNUpsamplingBilinearNodeClass.class), objc.Sel("nodeWithSource:integerScaleFactorX:integerScaleFactorY:"), sourceNode, integerScaleFactorX, integerScaleFactorY)
	return MPSCNNUpsamplingBilinearNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode/nodeWithSource:integerScaleFactorX:integerScaleFactorY:alignCorners:
func (_MPSCNNUpsamplingBilinearNodeClass MPSCNNUpsamplingBilinearNodeClass) NodeWithSourceIntegerScaleFactorXIntegerScaleFactorYAlignCorners(sourceNode IMPSNNImageNode, integerScaleFactorX uint, integerScaleFactorY uint, alignCorners bool) MPSCNNUpsamplingBilinearNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNUpsamplingBilinearNodeClass.class), objc.Sel("nodeWithSource:integerScaleFactorX:integerScaleFactorY:alignCorners:"), sourceNode, integerScaleFactorX, integerScaleFactorY, alignCorners)
	return MPSCNNUpsamplingBilinearNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode/scaleFactorX
func (c MPSCNNUpsamplingBilinearNode) ScaleFactorX() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode/scaleFactorY
func (c MPSCNNUpsamplingBilinearNode) ScaleFactorY() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorY"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearNode/alignCorners
func (c MPSCNNUpsamplingBilinearNode) AlignCorners() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("alignCorners"))
	return rv
}

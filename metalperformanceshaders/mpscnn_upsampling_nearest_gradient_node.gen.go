// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNUpsamplingNearestGradientNode] class.
var (
	_MPSCNNUpsamplingNearestGradientNodeClass     MPSCNNUpsamplingNearestGradientNodeClass
	_MPSCNNUpsamplingNearestGradientNodeClassOnce sync.Once
)

func getMPSCNNUpsamplingNearestGradientNodeClass() MPSCNNUpsamplingNearestGradientNodeClass {
	_MPSCNNUpsamplingNearestGradientNodeClassOnce.Do(func() {
		_MPSCNNUpsamplingNearestGradientNodeClass = MPSCNNUpsamplingNearestGradientNodeClass{class: objc.GetClass("MPSCNNUpsamplingNearestGradientNode")}
	})
	return _MPSCNNUpsamplingNearestGradientNodeClass
}

// GetMPSCNNUpsamplingNearestGradientNodeClass returns the class object for MPSCNNUpsamplingNearestGradientNode.
func GetMPSCNNUpsamplingNearestGradientNodeClass() MPSCNNUpsamplingNearestGradientNodeClass {
	return getMPSCNNUpsamplingNearestGradientNodeClass()
}

type MPSCNNUpsamplingNearestGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNUpsamplingNearestGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNUpsamplingNearestGradientNodeClass) Alloc() MPSCNNUpsamplingNearestGradientNode {
	rv := objc.Send[MPSCNNUpsamplingNearestGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient nearest spatial upsampling filter.
//
// # Initializers
//
//   - [MPSCNNUpsamplingNearestGradientNode.InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY]
//
// # Instance Properties
//
//   - [MPSCNNUpsamplingNearestGradientNode.ScaleFactorX]
//   - [MPSCNNUpsamplingNearestGradientNode.ScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradientNode
type MPSCNNUpsamplingNearestGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNUpsamplingNearestGradientNodeFromID constructs a [MPSCNNUpsamplingNearestGradientNode] from an objc.ID.
//
// A representation of a gradient nearest spatial upsampling filter.
func MPSCNNUpsamplingNearestGradientNodeFromID(id objc.ID) MPSCNNUpsamplingNearestGradientNode {
	return MPSCNNUpsamplingNearestGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNUpsamplingNearestGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNUpsamplingNearestGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNUpsamplingNearestGradientNode.InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY]
//
// # Instance Properties
//
//   - [IMPSCNNUpsamplingNearestGradientNode.ScaleFactorX]
//   - [IMPSCNNUpsamplingNearestGradientNode.ScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradientNode
type IMPSCNNUpsamplingNearestGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, scaleFactorX float64, scaleFactorY float64) MPSCNNUpsamplingNearestGradientNode

	// Topic: Instance Properties

	ScaleFactorX() float64
	ScaleFactorY() float64
}

// Init initializes the instance.
func (c MPSCNNUpsamplingNearestGradientNode) Init() MPSCNNUpsamplingNearestGradientNode {
	rv := objc.Send[MPSCNNUpsamplingNearestGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNUpsamplingNearestGradientNode) Autorelease() MPSCNNUpsamplingNearestGradientNode {
	rv := objc.Send[MPSCNNUpsamplingNearestGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNUpsamplingNearestGradientNode creates a new MPSCNNUpsamplingNearestGradientNode instance.
func NewMPSCNNUpsamplingNearestGradientNode() MPSCNNUpsamplingNearestGradientNode {
	class := getMPSCNNUpsamplingNearestGradientNodeClass()
	rv := objc.Send[MPSCNNUpsamplingNearestGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradientNode/init(sourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:)
func NewCNNUpsamplingNearestGradientNodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, scaleFactorX float64, scaleFactorY float64) MPSCNNUpsamplingNearestGradientNode {
	instance := getMPSCNNUpsamplingNearestGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	return MPSCNNUpsamplingNearestGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradientNode/init(sourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:)
func (c MPSCNNUpsamplingNearestGradientNode) InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, scaleFactorX float64, scaleFactorY float64) MPSCNNUpsamplingNearestGradientNode {
	rv := objc.Send[MPSCNNUpsamplingNearestGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradientNode/nodeWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:
func (_MPSCNNUpsamplingNearestGradientNodeClass MPSCNNUpsamplingNearestGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, scaleFactorX float64, scaleFactorY float64) MPSCNNUpsamplingNearestGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNUpsamplingNearestGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	return MPSCNNUpsamplingNearestGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradientNode/scaleFactorX
func (c MPSCNNUpsamplingNearestGradientNode) ScaleFactorX() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingNearestGradientNode/scaleFactorY
func (c MPSCNNUpsamplingNearestGradientNode) ScaleFactorY() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorY"))
	return rv
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNUpsamplingBilinearGradientNode] class.
var (
	_MPSCNNUpsamplingBilinearGradientNodeClass     MPSCNNUpsamplingBilinearGradientNodeClass
	_MPSCNNUpsamplingBilinearGradientNodeClassOnce sync.Once
)

func getMPSCNNUpsamplingBilinearGradientNodeClass() MPSCNNUpsamplingBilinearGradientNodeClass {
	_MPSCNNUpsamplingBilinearGradientNodeClassOnce.Do(func() {
		_MPSCNNUpsamplingBilinearGradientNodeClass = MPSCNNUpsamplingBilinearGradientNodeClass{class: objc.GetClass("MPSCNNUpsamplingBilinearGradientNode")}
	})
	return _MPSCNNUpsamplingBilinearGradientNodeClass
}

// GetMPSCNNUpsamplingBilinearGradientNodeClass returns the class object for MPSCNNUpsamplingBilinearGradientNode.
func GetMPSCNNUpsamplingBilinearGradientNodeClass() MPSCNNUpsamplingBilinearGradientNodeClass {
	return getMPSCNNUpsamplingBilinearGradientNodeClass()
}

type MPSCNNUpsamplingBilinearGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNUpsamplingBilinearGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNUpsamplingBilinearGradientNodeClass) Alloc() MPSCNNUpsamplingBilinearGradientNode {
	rv := objc.Send[MPSCNNUpsamplingBilinearGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient bilinear spatial upsampling filter.
//
// # Initializers
//
//   - [MPSCNNUpsamplingBilinearGradientNode.InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY]
//
// # Instance Properties
//
//   - [MPSCNNUpsamplingBilinearGradientNode.ScaleFactorX]
//   - [MPSCNNUpsamplingBilinearGradientNode.ScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradientNode
type MPSCNNUpsamplingBilinearGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNUpsamplingBilinearGradientNodeFromID constructs a [MPSCNNUpsamplingBilinearGradientNode] from an objc.ID.
//
// A representation of a gradient bilinear spatial upsampling filter.
func MPSCNNUpsamplingBilinearGradientNodeFromID(id objc.ID) MPSCNNUpsamplingBilinearGradientNode {
	return MPSCNNUpsamplingBilinearGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNUpsamplingBilinearGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNUpsamplingBilinearGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNUpsamplingBilinearGradientNode.InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY]
//
// # Instance Properties
//
//   - [IMPSCNNUpsamplingBilinearGradientNode.ScaleFactorX]
//   - [IMPSCNNUpsamplingBilinearGradientNode.ScaleFactorY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradientNode
type IMPSCNNUpsamplingBilinearGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, scaleFactorX float64, scaleFactorY float64) MPSCNNUpsamplingBilinearGradientNode

	// Topic: Instance Properties

	ScaleFactorX() float64
	ScaleFactorY() float64
}

// Init initializes the instance.
func (c MPSCNNUpsamplingBilinearGradientNode) Init() MPSCNNUpsamplingBilinearGradientNode {
	rv := objc.Send[MPSCNNUpsamplingBilinearGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNUpsamplingBilinearGradientNode) Autorelease() MPSCNNUpsamplingBilinearGradientNode {
	rv := objc.Send[MPSCNNUpsamplingBilinearGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNUpsamplingBilinearGradientNode creates a new MPSCNNUpsamplingBilinearGradientNode instance.
func NewMPSCNNUpsamplingBilinearGradientNode() MPSCNNUpsamplingBilinearGradientNode {
	class := getMPSCNNUpsamplingBilinearGradientNodeClass()
	rv := objc.Send[MPSCNNUpsamplingBilinearGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradientNode/init(sourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:)
func NewCNNUpsamplingBilinearGradientNodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, scaleFactorX float64, scaleFactorY float64) MPSCNNUpsamplingBilinearGradientNode {
	instance := getMPSCNNUpsamplingBilinearGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	return MPSCNNUpsamplingBilinearGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradientNode/init(sourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:)
func (c MPSCNNUpsamplingBilinearGradientNode) InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, scaleFactorX float64, scaleFactorY float64) MPSCNNUpsamplingBilinearGradientNode {
	rv := objc.Send[MPSCNNUpsamplingBilinearGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradientNode/nodeWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:
func (_MPSCNNUpsamplingBilinearGradientNodeClass MPSCNNUpsamplingBilinearGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, scaleFactorX float64, scaleFactorY float64) MPSCNNUpsamplingBilinearGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNUpsamplingBilinearGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	return MPSCNNUpsamplingBilinearGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradientNode/scaleFactorX
func (c MPSCNNUpsamplingBilinearGradientNode) ScaleFactorX() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNUpsamplingBilinearGradientNode/scaleFactorY
func (c MPSCNNUpsamplingBilinearGradientNode) ScaleFactorY() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactorY"))
	return rv
}

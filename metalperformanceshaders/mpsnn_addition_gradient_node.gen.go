// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNAdditionGradientNode] class.
var (
	_MPSNNAdditionGradientNodeClass     MPSNNAdditionGradientNodeClass
	_MPSNNAdditionGradientNodeClassOnce sync.Once
)

func getMPSNNAdditionGradientNodeClass() MPSNNAdditionGradientNodeClass {
	_MPSNNAdditionGradientNodeClassOnce.Do(func() {
		_MPSNNAdditionGradientNodeClass = MPSNNAdditionGradientNodeClass{class: objc.GetClass("MPSNNAdditionGradientNode")}
	})
	return _MPSNNAdditionGradientNodeClass
}

// GetMPSNNAdditionGradientNodeClass returns the class object for MPSNNAdditionGradientNode.
func GetMPSNNAdditionGradientNodeClass() MPSNNAdditionGradientNodeClass {
	return getMPSNNAdditionGradientNodeClass()
}

type MPSNNAdditionGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNAdditionGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNAdditionGradientNodeClass) Alloc() MPSNNAdditionGradientNode {
	rv := objc.Send[MPSNNAdditionGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient addition operator.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNAdditionGradientNode
type MPSNNAdditionGradientNode struct {
	MPSNNArithmeticGradientNode
}

// MPSNNAdditionGradientNodeFromID constructs a [MPSNNAdditionGradientNode] from an objc.ID.
//
// A representation of a gradient addition operator.
func MPSNNAdditionGradientNodeFromID(id objc.ID) MPSNNAdditionGradientNode {
	return MPSNNAdditionGradientNode{MPSNNArithmeticGradientNode: MPSNNArithmeticGradientNodeFromID(id)}
}

// NOTE: MPSNNAdditionGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNAdditionGradientNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNAdditionGradientNode
type IMPSNNAdditionGradientNode interface {
	IMPSNNArithmeticGradientNode
}

// Init initializes the instance.
func (a MPSNNAdditionGradientNode) Init() MPSNNAdditionGradientNode {
	rv := objc.Send[MPSNNAdditionGradientNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MPSNNAdditionGradientNode) Autorelease() MPSNNAdditionGradientNode {
	rv := objc.Send[MPSNNAdditionGradientNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNAdditionGradientNode creates a new MPSNNAdditionGradientNode instance.
func NewMPSNNAdditionGradientNode() MPSNNAdditionGradientNode {
	class := getMPSNNAdditionGradientNodeClass()
	rv := objc.Send[MPSNNAdditionGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/init(gradientImages:forwardFilter:isSecondarySourceFilter:)
func NewAdditionGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []MPSNNImageNode, filter IMPSNNFilterNode, isSecondarySourceFilter bool) MPSNNAdditionGradientNode {
	instance := getMPSNNAdditionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), objectivec.IObjectSliceToNSArray(gradientImages), filter, isSecondarySourceFilter)
	return MPSNNAdditionGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/init(sourceGradient:sourceImage:gradientState:isSecondarySourceFilter:)
func NewAdditionGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNBinaryGradientStateNode, isSecondarySourceFilter bool) MPSNNAdditionGradientNode {
	instance := getMPSNNAdditionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	return MPSNNAdditionGradientNodeFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNMultiplicationGradientNode] class.
var (
	_MPSNNMultiplicationGradientNodeClass     MPSNNMultiplicationGradientNodeClass
	_MPSNNMultiplicationGradientNodeClassOnce sync.Once
)

func getMPSNNMultiplicationGradientNodeClass() MPSNNMultiplicationGradientNodeClass {
	_MPSNNMultiplicationGradientNodeClassOnce.Do(func() {
		_MPSNNMultiplicationGradientNodeClass = MPSNNMultiplicationGradientNodeClass{class: objc.GetClass("MPSNNMultiplicationGradientNode")}
	})
	return _MPSNNMultiplicationGradientNodeClass
}

// GetMPSNNMultiplicationGradientNodeClass returns the class object for MPSNNMultiplicationGradientNode.
func GetMPSNNMultiplicationGradientNodeClass() MPSNNMultiplicationGradientNodeClass {
	return getMPSNNMultiplicationGradientNodeClass()
}

type MPSNNMultiplicationGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNMultiplicationGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNMultiplicationGradientNodeClass) Alloc() MPSNNMultiplicationGradientNode {
	rv := objc.Send[MPSNNMultiplicationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient multiplication operator.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiplicationGradientNode
type MPSNNMultiplicationGradientNode struct {
	MPSNNArithmeticGradientNode
}

// MPSNNMultiplicationGradientNodeFromID constructs a [MPSNNMultiplicationGradientNode] from an objc.ID.
//
// A representation of a gradient multiplication operator.
func MPSNNMultiplicationGradientNodeFromID(id objc.ID) MPSNNMultiplicationGradientNode {
	return MPSNNMultiplicationGradientNode{MPSNNArithmeticGradientNode: MPSNNArithmeticGradientNodeFromID(id)}
}

// NOTE: MPSNNMultiplicationGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNMultiplicationGradientNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiplicationGradientNode
type IMPSNNMultiplicationGradientNode interface {
	IMPSNNArithmeticGradientNode
}

// Init initializes the instance.
func (m MPSNNMultiplicationGradientNode) Init() MPSNNMultiplicationGradientNode {
	rv := objc.Send[MPSNNMultiplicationGradientNode](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSNNMultiplicationGradientNode) Autorelease() MPSNNMultiplicationGradientNode {
	rv := objc.Send[MPSNNMultiplicationGradientNode](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNMultiplicationGradientNode creates a new MPSNNMultiplicationGradientNode instance.
func NewMPSNNMultiplicationGradientNode() MPSNNMultiplicationGradientNode {
	class := getMPSNNMultiplicationGradientNodeClass()
	rv := objc.Send[MPSNNMultiplicationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/init(gradientImages:forwardFilter:isSecondarySourceFilter:)
func NewMultiplicationGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []MPSNNImageNode, filter IMPSNNFilterNode, isSecondarySourceFilter bool) MPSNNMultiplicationGradientNode {
	instance := getMPSNNMultiplicationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), objectivec.IObjectSliceToNSArray(gradientImages), filter, isSecondarySourceFilter)
	return MPSNNMultiplicationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/init(sourceGradient:sourceImage:gradientState:isSecondarySourceFilter:)
func NewMultiplicationGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNBinaryGradientStateNode, isSecondarySourceFilter bool) MPSNNMultiplicationGradientNode {
	instance := getMPSNNMultiplicationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	return MPSNNMultiplicationGradientNodeFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNSubtractionGradientNode] class.
var (
	_MPSNNSubtractionGradientNodeClass     MPSNNSubtractionGradientNodeClass
	_MPSNNSubtractionGradientNodeClassOnce sync.Once
)

func getMPSNNSubtractionGradientNodeClass() MPSNNSubtractionGradientNodeClass {
	_MPSNNSubtractionGradientNodeClassOnce.Do(func() {
		_MPSNNSubtractionGradientNodeClass = MPSNNSubtractionGradientNodeClass{class: objc.GetClass("MPSNNSubtractionGradientNode")}
	})
	return _MPSNNSubtractionGradientNodeClass
}

// GetMPSNNSubtractionGradientNodeClass returns the class object for MPSNNSubtractionGradientNode.
func GetMPSNNSubtractionGradientNodeClass() MPSNNSubtractionGradientNodeClass {
	return getMPSNNSubtractionGradientNodeClass()
}

type MPSNNSubtractionGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNSubtractionGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNSubtractionGradientNodeClass) Alloc() MPSNNSubtractionGradientNode {
	rv := objc.Send[MPSNNSubtractionGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient subtraction operator.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSubtractionGradientNode
type MPSNNSubtractionGradientNode struct {
	MPSNNArithmeticGradientNode
}

// MPSNNSubtractionGradientNodeFromID constructs a [MPSNNSubtractionGradientNode] from an objc.ID.
//
// A representation of a gradient subtraction operator.
func MPSNNSubtractionGradientNodeFromID(id objc.ID) MPSNNSubtractionGradientNode {
	return MPSNNSubtractionGradientNode{MPSNNArithmeticGradientNode: MPSNNArithmeticGradientNodeFromID(id)}
}

// NOTE: MPSNNSubtractionGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNSubtractionGradientNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSubtractionGradientNode
type IMPSNNSubtractionGradientNode interface {
	IMPSNNArithmeticGradientNode
}

// Init initializes the instance.
func (s MPSNNSubtractionGradientNode) Init() MPSNNSubtractionGradientNode {
	rv := objc.Send[MPSNNSubtractionGradientNode](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSNNSubtractionGradientNode) Autorelease() MPSNNSubtractionGradientNode {
	rv := objc.Send[MPSNNSubtractionGradientNode](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNSubtractionGradientNode creates a new MPSNNSubtractionGradientNode instance.
func NewMPSNNSubtractionGradientNode() MPSNNSubtractionGradientNode {
	class := getMPSNNSubtractionGradientNodeClass()
	rv := objc.Send[MPSNNSubtractionGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/init(gradientImages:forwardFilter:isSecondarySourceFilter:)
func NewSubtractionGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []MPSNNImageNode, filter IMPSNNFilterNode, isSecondarySourceFilter bool) MPSNNSubtractionGradientNode {
	instance := getMPSNNSubtractionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), objectivec.IObjectSliceToNSArray(gradientImages), filter, isSecondarySourceFilter)
	return MPSNNSubtractionGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/init(sourceGradient:sourceImage:gradientState:isSecondarySourceFilter:)
func NewSubtractionGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNBinaryGradientStateNode, isSecondarySourceFilter bool) MPSNNSubtractionGradientNode {
	instance := getMPSNNSubtractionGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	return MPSNNSubtractionGradientNodeFromID(rv)
}

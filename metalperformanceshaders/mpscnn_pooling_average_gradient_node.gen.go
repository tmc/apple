// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingAverageGradientNode] class.
var (
	_MPSCNNPoolingAverageGradientNodeClass     MPSCNNPoolingAverageGradientNodeClass
	_MPSCNNPoolingAverageGradientNodeClassOnce sync.Once
)

func getMPSCNNPoolingAverageGradientNodeClass() MPSCNNPoolingAverageGradientNodeClass {
	_MPSCNNPoolingAverageGradientNodeClassOnce.Do(func() {
		_MPSCNNPoolingAverageGradientNodeClass = MPSCNNPoolingAverageGradientNodeClass{class: objc.GetClass("MPSCNNPoolingAverageGradientNode")}
	})
	return _MPSCNNPoolingAverageGradientNodeClass
}

// GetMPSCNNPoolingAverageGradientNodeClass returns the class object for MPSCNNPoolingAverageGradientNode.
func GetMPSCNNPoolingAverageGradientNodeClass() MPSCNNPoolingAverageGradientNodeClass {
	return getMPSCNNPoolingAverageGradientNodeClass()
}

type MPSCNNPoolingAverageGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingAverageGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingAverageGradientNodeClass) Alloc() MPSCNNPoolingAverageGradientNode {
	rv := objc.Send[MPSCNNPoolingAverageGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient average pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradientNode
type MPSCNNPoolingAverageGradientNode struct {
	MPSCNNPoolingGradientNode
}

// MPSCNNPoolingAverageGradientNodeFromID constructs a [MPSCNNPoolingAverageGradientNode] from an objc.ID.
//
// A representation of a gradient average pooling filter.
func MPSCNNPoolingAverageGradientNodeFromID(id objc.ID) MPSCNNPoolingAverageGradientNode {
	return MPSCNNPoolingAverageGradientNode{MPSCNNPoolingGradientNode: MPSCNNPoolingGradientNodeFromID(id)}
}

// NOTE: MPSCNNPoolingAverageGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingAverageGradientNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradientNode
type IMPSCNNPoolingAverageGradientNode interface {
	IMPSCNNPoolingGradientNode
}

// Init initializes the instance.
func (c MPSCNNPoolingAverageGradientNode) Init() MPSCNNPoolingAverageGradientNode {
	rv := objc.Send[MPSCNNPoolingAverageGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingAverageGradientNode) Autorelease() MPSCNNPoolingAverageGradientNode {
	rv := objc.Send[MPSCNNPoolingAverageGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingAverageGradientNode creates a new MPSCNNPoolingAverageGradientNode instance.
func NewMPSCNNPoolingAverageGradientNode() MPSCNNPoolingAverageGradientNode {
	class := getMPSCNNPoolingAverageGradientNodeClass()
	rv := objc.Send[MPSCNNPoolingAverageGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/init(sourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:)
func NewCNNPoolingAverageGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy MPSNNPadding) MPSCNNPoolingAverageGradientNode {
	instance := getMPSCNNPoolingAverageGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	return MPSCNNPoolingAverageGradientNodeFromID(rv)
}

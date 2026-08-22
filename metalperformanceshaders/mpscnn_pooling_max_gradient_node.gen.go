// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingMaxGradientNode] class.
var (
	_MPSCNNPoolingMaxGradientNodeClass     MPSCNNPoolingMaxGradientNodeClass
	_MPSCNNPoolingMaxGradientNodeClassOnce sync.Once
)

func getMPSCNNPoolingMaxGradientNodeClass() MPSCNNPoolingMaxGradientNodeClass {
	_MPSCNNPoolingMaxGradientNodeClassOnce.Do(func() {
		_MPSCNNPoolingMaxGradientNodeClass = MPSCNNPoolingMaxGradientNodeClass{class: objc.GetClass("MPSCNNPoolingMaxGradientNode")}
	})
	return _MPSCNNPoolingMaxGradientNodeClass
}

// GetMPSCNNPoolingMaxGradientNodeClass returns the class object for MPSCNNPoolingMaxGradientNode.
func GetMPSCNNPoolingMaxGradientNodeClass() MPSCNNPoolingMaxGradientNodeClass {
	return getMPSCNNPoolingMaxGradientNodeClass()
}

type MPSCNNPoolingMaxGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingMaxGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingMaxGradientNodeClass) Alloc() MPSCNNPoolingMaxGradientNode {
	rv := objc.Send[MPSCNNPoolingMaxGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient max pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxGradientNode
type MPSCNNPoolingMaxGradientNode struct {
	MPSCNNPoolingGradientNode
}

// MPSCNNPoolingMaxGradientNodeFromID constructs a [MPSCNNPoolingMaxGradientNode] from an objc.ID.
//
// A representation of a gradient max pooling filter.
func MPSCNNPoolingMaxGradientNodeFromID(id objc.ID) MPSCNNPoolingMaxGradientNode {
	return MPSCNNPoolingMaxGradientNode{MPSCNNPoolingGradientNode: MPSCNNPoolingGradientNodeFromID(id)}
}

// NOTE: MPSCNNPoolingMaxGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingMaxGradientNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxGradientNode
type IMPSCNNPoolingMaxGradientNode interface {
	IMPSCNNPoolingGradientNode
}

// Init initializes the instance.
func (c MPSCNNPoolingMaxGradientNode) Init() MPSCNNPoolingMaxGradientNode {
	rv := objc.Send[MPSCNNPoolingMaxGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingMaxGradientNode) Autorelease() MPSCNNPoolingMaxGradientNode {
	rv := objc.Send[MPSCNNPoolingMaxGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingMaxGradientNode creates a new MPSCNNPoolingMaxGradientNode instance.
func NewMPSCNNPoolingMaxGradientNode() MPSCNNPoolingMaxGradientNode {
	class := getMPSCNNPoolingMaxGradientNodeClass()
	rv := objc.Send[MPSCNNPoolingMaxGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/init(sourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:)
func NewCNNPoolingMaxGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy MPSNNPadding) MPSCNNPoolingMaxGradientNode {
	instance := getMPSCNNPoolingMaxGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	return MPSCNNPoolingMaxGradientNodeFromID(rv)
}

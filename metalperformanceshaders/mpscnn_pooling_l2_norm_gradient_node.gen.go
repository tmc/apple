// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingL2NormGradientNode] class.
var (
	_MPSCNNPoolingL2NormGradientNodeClass     MPSCNNPoolingL2NormGradientNodeClass
	_MPSCNNPoolingL2NormGradientNodeClassOnce sync.Once
)

func getMPSCNNPoolingL2NormGradientNodeClass() MPSCNNPoolingL2NormGradientNodeClass {
	_MPSCNNPoolingL2NormGradientNodeClassOnce.Do(func() {
		_MPSCNNPoolingL2NormGradientNodeClass = MPSCNNPoolingL2NormGradientNodeClass{class: objc.GetClass("MPSCNNPoolingL2NormGradientNode")}
	})
	return _MPSCNNPoolingL2NormGradientNodeClass
}

// GetMPSCNNPoolingL2NormGradientNodeClass returns the class object for MPSCNNPoolingL2NormGradientNode.
func GetMPSCNNPoolingL2NormGradientNodeClass() MPSCNNPoolingL2NormGradientNodeClass {
	return getMPSCNNPoolingL2NormGradientNodeClass()
}

type MPSCNNPoolingL2NormGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingL2NormGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingL2NormGradientNodeClass) Alloc() MPSCNNPoolingL2NormGradientNode {
	rv := objc.Send[MPSCNNPoolingL2NormGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient L2-norm pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormGradientNode
type MPSCNNPoolingL2NormGradientNode struct {
	MPSCNNPoolingGradientNode
}

// MPSCNNPoolingL2NormGradientNodeFromID constructs a [MPSCNNPoolingL2NormGradientNode] from an objc.ID.
//
// A representation of a gradient L2-norm pooling filter.
func MPSCNNPoolingL2NormGradientNodeFromID(id objc.ID) MPSCNNPoolingL2NormGradientNode {
	return MPSCNNPoolingL2NormGradientNode{MPSCNNPoolingGradientNode: MPSCNNPoolingGradientNodeFromID(id)}
}

// NOTE: MPSCNNPoolingL2NormGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingL2NormGradientNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormGradientNode
type IMPSCNNPoolingL2NormGradientNode interface {
	IMPSCNNPoolingGradientNode
}

// Init initializes the instance.
func (c MPSCNNPoolingL2NormGradientNode) Init() MPSCNNPoolingL2NormGradientNode {
	rv := objc.Send[MPSCNNPoolingL2NormGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingL2NormGradientNode) Autorelease() MPSCNNPoolingL2NormGradientNode {
	rv := objc.Send[MPSCNNPoolingL2NormGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingL2NormGradientNode creates a new MPSCNNPoolingL2NormGradientNode instance.
func NewMPSCNNPoolingL2NormGradientNode() MPSCNNPoolingL2NormGradientNode {
	class := getMPSCNNPoolingL2NormGradientNodeClass()
	rv := objc.Send[MPSCNNPoolingL2NormGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/init(sourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:)
func NewCNNPoolingL2NormGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy MPSNNPadding) MPSCNNPoolingL2NormGradientNode {
	instance := getMPSCNNPoolingL2NormGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	return MPSCNNPoolingL2NormGradientNodeFromID(rv)
}

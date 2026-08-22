// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingGradientNode] class.
var (
	_MPSCNNPoolingGradientNodeClass     MPSCNNPoolingGradientNodeClass
	_MPSCNNPoolingGradientNodeClassOnce sync.Once
)

func getMPSCNNPoolingGradientNodeClass() MPSCNNPoolingGradientNodeClass {
	_MPSCNNPoolingGradientNodeClassOnce.Do(func() {
		_MPSCNNPoolingGradientNodeClass = MPSCNNPoolingGradientNodeClass{class: objc.GetClass("MPSCNNPoolingGradientNode")}
	})
	return _MPSCNNPoolingGradientNodeClass
}

// GetMPSCNNPoolingGradientNodeClass returns the class object for MPSCNNPoolingGradientNode.
func GetMPSCNNPoolingGradientNodeClass() MPSCNNPoolingGradientNodeClass {
	return getMPSCNNPoolingGradientNodeClass()
}

type MPSCNNPoolingGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingGradientNodeClass) Alloc() MPSCNNPoolingGradientNode {
	rv := objc.Send[MPSCNNPoolingGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient pooling kernel.
//
// # Initializers
//
//   - [MPSCNNPoolingGradientNode.InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy]
//
// # Instance Properties
//
//   - [MPSCNNPoolingGradientNode.KernelHeight]
//   - [MPSCNNPoolingGradientNode.KernelWidth]
//   - [MPSCNNPoolingGradientNode.StrideInPixelsX]
//   - [MPSCNNPoolingGradientNode.StrideInPixelsY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode
type MPSCNNPoolingGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNPoolingGradientNodeFromID constructs a [MPSCNNPoolingGradientNode] from an objc.ID.
//
// A representation of a gradient pooling kernel.
func MPSCNNPoolingGradientNodeFromID(id objc.ID) MPSCNNPoolingGradientNode {
	return MPSCNNPoolingGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNPoolingGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNPoolingGradientNode.InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy]
//
// # Instance Properties
//
//   - [IMPSCNNPoolingGradientNode.KernelHeight]
//   - [IMPSCNNPoolingGradientNode.KernelWidth]
//   - [IMPSCNNPoolingGradientNode.StrideInPixelsX]
//   - [IMPSCNNPoolingGradientNode.StrideInPixelsY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode
type IMPSCNNPoolingGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy MPSNNPadding) MPSCNNPoolingGradientNode

	// Topic: Instance Properties

	KernelHeight() uint
	KernelWidth() uint
	StrideInPixelsX() uint
	StrideInPixelsY() uint
}

// Init initializes the instance.
func (c MPSCNNPoolingGradientNode) Init() MPSCNNPoolingGradientNode {
	rv := objc.Send[MPSCNNPoolingGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingGradientNode) Autorelease() MPSCNNPoolingGradientNode {
	rv := objc.Send[MPSCNNPoolingGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingGradientNode creates a new MPSCNNPoolingGradientNode instance.
func NewMPSCNNPoolingGradientNode() MPSCNNPoolingGradientNode {
	class := getMPSCNNPoolingGradientNodeClass()
	rv := objc.Send[MPSCNNPoolingGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/init(sourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:)
func NewCNNPoolingGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy MPSNNPadding) MPSCNNPoolingGradientNode {
	instance := getMPSCNNPoolingGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	return MPSCNNPoolingGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/init(sourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:)
func (c MPSCNNPoolingGradientNode) InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy MPSNNPadding) MPSCNNPoolingGradientNode {
	rv := objc.Send[MPSCNNPoolingGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:
func (_MPSCNNPoolingGradientNodeClass MPSCNNPoolingGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy MPSNNPadding) MPSCNNPoolingGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNPoolingGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	return MPSCNNPoolingGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/kernelHeight
func (c MPSCNNPoolingGradientNode) KernelHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/kernelWidth
func (c MPSCNNPoolingGradientNode) KernelWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelWidth"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/strideInPixelsX
func (c MPSCNNPoolingGradientNode) StrideInPixelsX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("strideInPixelsX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/strideInPixelsY
func (c MPSCNNPoolingGradientNode) StrideInPixelsY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("strideInPixelsY"))
	return rv
}

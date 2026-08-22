// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNDilatedPoolingMaxGradientNode] class.
var (
	_MPSCNNDilatedPoolingMaxGradientNodeClass     MPSCNNDilatedPoolingMaxGradientNodeClass
	_MPSCNNDilatedPoolingMaxGradientNodeClassOnce sync.Once
)

func getMPSCNNDilatedPoolingMaxGradientNodeClass() MPSCNNDilatedPoolingMaxGradientNodeClass {
	_MPSCNNDilatedPoolingMaxGradientNodeClassOnce.Do(func() {
		_MPSCNNDilatedPoolingMaxGradientNodeClass = MPSCNNDilatedPoolingMaxGradientNodeClass{class: objc.GetClass("MPSCNNDilatedPoolingMaxGradientNode")}
	})
	return _MPSCNNDilatedPoolingMaxGradientNodeClass
}

// GetMPSCNNDilatedPoolingMaxGradientNodeClass returns the class object for MPSCNNDilatedPoolingMaxGradientNode.
func GetMPSCNNDilatedPoolingMaxGradientNodeClass() MPSCNNDilatedPoolingMaxGradientNodeClass {
	return getMPSCNNDilatedPoolingMaxGradientNodeClass()
}

type MPSCNNDilatedPoolingMaxGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDilatedPoolingMaxGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDilatedPoolingMaxGradientNodeClass) Alloc() MPSCNNDilatedPoolingMaxGradientNode {
	rv := objc.Send[MPSCNNDilatedPoolingMaxGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient dilated max pooling filter.
//
// # Initializers
//
//   - [MPSCNNDilatedPoolingMaxGradientNode.InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY]
//
// # Instance Properties
//
//   - [MPSCNNDilatedPoolingMaxGradientNode.DilationRateX]
//   - [MPSCNNDilatedPoolingMaxGradientNode.DilationRateY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradientNode
type MPSCNNDilatedPoolingMaxGradientNode struct {
	MPSCNNPoolingGradientNode
}

// MPSCNNDilatedPoolingMaxGradientNodeFromID constructs a [MPSCNNDilatedPoolingMaxGradientNode] from an objc.ID.
//
// A representation of a gradient dilated max pooling filter.
func MPSCNNDilatedPoolingMaxGradientNodeFromID(id objc.ID) MPSCNNDilatedPoolingMaxGradientNode {
	return MPSCNNDilatedPoolingMaxGradientNode{MPSCNNPoolingGradientNode: MPSCNNPoolingGradientNodeFromID(id)}
}

// NOTE: MPSCNNDilatedPoolingMaxGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDilatedPoolingMaxGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNDilatedPoolingMaxGradientNode.InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY]
//
// # Instance Properties
//
//   - [IMPSCNNDilatedPoolingMaxGradientNode.DilationRateX]
//   - [IMPSCNNDilatedPoolingMaxGradientNode.DilationRateY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradientNode
type IMPSCNNDilatedPoolingMaxGradientNode interface {
	IMPSCNNPoolingGradientNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) MPSCNNDilatedPoolingMaxGradientNode

	// Topic: Instance Properties

	DilationRateX() uint
	DilationRateY() uint
}

// Init initializes the instance.
func (c MPSCNNDilatedPoolingMaxGradientNode) Init() MPSCNNDilatedPoolingMaxGradientNode {
	rv := objc.Send[MPSCNNDilatedPoolingMaxGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDilatedPoolingMaxGradientNode) Autorelease() MPSCNNDilatedPoolingMaxGradientNode {
	rv := objc.Send[MPSCNNDilatedPoolingMaxGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDilatedPoolingMaxGradientNode creates a new MPSCNNDilatedPoolingMaxGradientNode instance.
func NewMPSCNNDilatedPoolingMaxGradientNode() MPSCNNDilatedPoolingMaxGradientNode {
	class := getMPSCNNDilatedPoolingMaxGradientNodeClass()
	rv := objc.Send[MPSCNNDilatedPoolingMaxGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradientNode/init(sourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:)
func NewCNNDilatedPoolingMaxGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) MPSCNNDilatedPoolingMaxGradientNode {
	instance := getMPSCNNDilatedPoolingMaxGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	return MPSCNNDilatedPoolingMaxGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradientNode/init(sourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:)
func NewCNNDilatedPoolingMaxGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy MPSNNPadding) MPSCNNDilatedPoolingMaxGradientNode {
	instance := getMPSCNNDilatedPoolingMaxGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	return MPSCNNDilatedPoolingMaxGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradientNode/init(sourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:)
func (c MPSCNNDilatedPoolingMaxGradientNode) InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) MPSCNNDilatedPoolingMaxGradientNode {
	rv := objc.Send[MPSCNNDilatedPoolingMaxGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradientNode/nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:
func (_MPSCNNDilatedPoolingMaxGradientNodeClass MPSCNNDilatedPoolingMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) MPSCNNDilatedPoolingMaxGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNDilatedPoolingMaxGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	return MPSCNNDilatedPoolingMaxGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradientNode/dilationRateX
func (c MPSCNNDilatedPoolingMaxGradientNode) DilationRateX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dilationRateX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradientNode/dilationRateY
func (c MPSCNNDilatedPoolingMaxGradientNode) DilationRateY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dilationRateY"))
	return rv
}

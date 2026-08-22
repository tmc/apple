// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNDilatedPoolingMaxNode] class.
var (
	_MPSCNNDilatedPoolingMaxNodeClass     MPSCNNDilatedPoolingMaxNodeClass
	_MPSCNNDilatedPoolingMaxNodeClassOnce sync.Once
)

func getMPSCNNDilatedPoolingMaxNodeClass() MPSCNNDilatedPoolingMaxNodeClass {
	_MPSCNNDilatedPoolingMaxNodeClassOnce.Do(func() {
		_MPSCNNDilatedPoolingMaxNodeClass = MPSCNNDilatedPoolingMaxNodeClass{class: objc.GetClass("MPSCNNDilatedPoolingMaxNode")}
	})
	return _MPSCNNDilatedPoolingMaxNodeClass
}

// GetMPSCNNDilatedPoolingMaxNodeClass returns the class object for MPSCNNDilatedPoolingMaxNode.
func GetMPSCNNDilatedPoolingMaxNodeClass() MPSCNNDilatedPoolingMaxNodeClass {
	return getMPSCNNDilatedPoolingMaxNodeClass()
}

type MPSCNNDilatedPoolingMaxNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDilatedPoolingMaxNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDilatedPoolingMaxNodeClass) Alloc() MPSCNNDilatedPoolingMaxNode {
	rv := objc.Send[MPSCNNDilatedPoolingMaxNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a dilated max pooling filter.
//
// # Initializers
//
//   - [MPSCNNDilatedPoolingMaxNode.InitWithSourceFilterSize]
//   - [MPSCNNDilatedPoolingMaxNode.InitWithSourceFilterSizeStrideDilationRate]
//   - [MPSCNNDilatedPoolingMaxNode.InitWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY]
//
// # Instance Properties
//
//   - [MPSCNNDilatedPoolingMaxNode.DilationRateX]
//   - [MPSCNNDilatedPoolingMaxNode.DilationRateY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode
type MPSCNNDilatedPoolingMaxNode struct {
	MPSNNFilterNode
}

// MPSCNNDilatedPoolingMaxNodeFromID constructs a [MPSCNNDilatedPoolingMaxNode] from an objc.ID.
//
// A representation of a dilated max pooling filter.
func MPSCNNDilatedPoolingMaxNodeFromID(id objc.ID) MPSCNNDilatedPoolingMaxNode {
	return MPSCNNDilatedPoolingMaxNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNDilatedPoolingMaxNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDilatedPoolingMaxNode] class.
//
// # Initializers
//
//   - [IMPSCNNDilatedPoolingMaxNode.InitWithSourceFilterSize]
//   - [IMPSCNNDilatedPoolingMaxNode.InitWithSourceFilterSizeStrideDilationRate]
//   - [IMPSCNNDilatedPoolingMaxNode.InitWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY]
//
// # Instance Properties
//
//   - [IMPSCNNDilatedPoolingMaxNode.DilationRateX]
//   - [IMPSCNNDilatedPoolingMaxNode.DilationRateY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode
type IMPSCNNDilatedPoolingMaxNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNDilatedPoolingMaxNode
	InitWithSourceFilterSizeStrideDilationRate(sourceNode IMPSNNImageNode, size uint, stride uint, dilationRate uint) MPSCNNDilatedPoolingMaxNode
	InitWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceNode IMPSNNImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) MPSCNNDilatedPoolingMaxNode

	// Topic: Instance Properties

	DilationRateX() uint
	DilationRateY() uint
}

// Init initializes the instance.
func (c MPSCNNDilatedPoolingMaxNode) Init() MPSCNNDilatedPoolingMaxNode {
	rv := objc.Send[MPSCNNDilatedPoolingMaxNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDilatedPoolingMaxNode) Autorelease() MPSCNNDilatedPoolingMaxNode {
	rv := objc.Send[MPSCNNDilatedPoolingMaxNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDilatedPoolingMaxNode creates a new MPSCNNDilatedPoolingMaxNode instance.
func NewMPSCNNDilatedPoolingMaxNode() MPSCNNDilatedPoolingMaxNode {
	class := getMPSCNNDilatedPoolingMaxNodeClass()
	rv := objc.Send[MPSCNNDilatedPoolingMaxNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode/init(source:filterSize:)
func NewCNNDilatedPoolingMaxNodeWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNDilatedPoolingMaxNode {
	instance := getMPSCNNDilatedPoolingMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:filterSize:"), sourceNode, size)
	return MPSCNNDilatedPoolingMaxNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode/init(source:filterSize:stride:dilationRate:)
func NewCNNDilatedPoolingMaxNodeWithSourceFilterSizeStrideDilationRate(sourceNode IMPSNNImageNode, size uint, stride uint, dilationRate uint) MPSCNNDilatedPoolingMaxNode {
	instance := getMPSCNNDilatedPoolingMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:filterSize:stride:dilationRate:"), sourceNode, size, stride, dilationRate)
	return MPSCNNDilatedPoolingMaxNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode/init(source:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:)
func NewCNNDilatedPoolingMaxNodeWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceNode IMPSNNImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) MPSCNNDilatedPoolingMaxNode {
	instance := getMPSCNNDilatedPoolingMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:"), sourceNode, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	return MPSCNNDilatedPoolingMaxNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode/init(source:filterSize:)
func (c MPSCNNDilatedPoolingMaxNode) InitWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNDilatedPoolingMaxNode {
	rv := objc.Send[MPSCNNDilatedPoolingMaxNode](c.ID, objc.Sel("initWithSource:filterSize:"), sourceNode, size)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode/init(source:filterSize:stride:dilationRate:)
func (c MPSCNNDilatedPoolingMaxNode) InitWithSourceFilterSizeStrideDilationRate(sourceNode IMPSNNImageNode, size uint, stride uint, dilationRate uint) MPSCNNDilatedPoolingMaxNode {
	rv := objc.Send[MPSCNNDilatedPoolingMaxNode](c.ID, objc.Sel("initWithSource:filterSize:stride:dilationRate:"), sourceNode, size, stride, dilationRate)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode/init(source:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:)
func (c MPSCNNDilatedPoolingMaxNode) InitWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceNode IMPSNNImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) MPSCNNDilatedPoolingMaxNode {
	rv := objc.Send[MPSCNNDilatedPoolingMaxNode](c.ID, objc.Sel("initWithSource:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:"), sourceNode, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode/nodeWithSource:filterSize:
func (_MPSCNNDilatedPoolingMaxNodeClass MPSCNNDilatedPoolingMaxNodeClass) NodeWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNDilatedPoolingMaxNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNDilatedPoolingMaxNodeClass.class), objc.Sel("nodeWithSource:filterSize:"), sourceNode, size)
	return MPSCNNDilatedPoolingMaxNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode/nodeWithSource:filterSize:stride:dilationRate:
func (_MPSCNNDilatedPoolingMaxNodeClass MPSCNNDilatedPoolingMaxNodeClass) NodeWithSourceFilterSizeStrideDilationRate(sourceNode IMPSNNImageNode, size uint, stride uint, dilationRate uint) MPSCNNDilatedPoolingMaxNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNDilatedPoolingMaxNodeClass.class), objc.Sel("nodeWithSource:filterSize:stride:dilationRate:"), sourceNode, size, stride, dilationRate)
	return MPSCNNDilatedPoolingMaxNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode/dilationRateX
func (c MPSCNNDilatedPoolingMaxNode) DilationRateX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dilationRateX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxNode/dilationRateY
func (c MPSCNNDilatedPoolingMaxNode) DilationRateY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("dilationRateY"))
	return rv
}

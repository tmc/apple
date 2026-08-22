// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingNode] class.
var (
	_MPSCNNPoolingNodeClass     MPSCNNPoolingNodeClass
	_MPSCNNPoolingNodeClassOnce sync.Once
)

func getMPSCNNPoolingNodeClass() MPSCNNPoolingNodeClass {
	_MPSCNNPoolingNodeClassOnce.Do(func() {
		_MPSCNNPoolingNodeClass = MPSCNNPoolingNodeClass{class: objc.GetClass("MPSCNNPoolingNode")}
	})
	return _MPSCNNPoolingNodeClass
}

// GetMPSCNNPoolingNodeClass returns the class object for MPSCNNPoolingNode.
func GetMPSCNNPoolingNodeClass() MPSCNNPoolingNodeClass {
	return getMPSCNNPoolingNodeClass()
}

type MPSCNNPoolingNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingNodeClass) Alloc() MPSCNNPoolingNode {
	rv := objc.Send[MPSCNNPoolingNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a MPS CNN pooling kernel.
//
// # Initializers
//
//   - [MPSCNNPoolingNode.InitWithSourceFilterSize]
//   - [MPSCNNPoolingNode.InitWithSourceFilterSizeStride]
//   - [MPSCNNPoolingNode.InitWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY]
//
// # Instance Properties
//
//   - [MPSCNNPoolingNode.KernelHeight]
//   - [MPSCNNPoolingNode.KernelWidth]
//   - [MPSCNNPoolingNode.StrideInPixelsX]
//   - [MPSCNNPoolingNode.StrideInPixelsY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode
type MPSCNNPoolingNode struct {
	MPSNNFilterNode
}

// MPSCNNPoolingNodeFromID constructs a [MPSCNNPoolingNode] from an objc.ID.
//
// A representation of a MPS CNN pooling kernel.
func MPSCNNPoolingNodeFromID(id objc.ID) MPSCNNPoolingNode {
	return MPSCNNPoolingNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNPoolingNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingNode] class.
//
// # Initializers
//
//   - [IMPSCNNPoolingNode.InitWithSourceFilterSize]
//   - [IMPSCNNPoolingNode.InitWithSourceFilterSizeStride]
//   - [IMPSCNNPoolingNode.InitWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY]
//
// # Instance Properties
//
//   - [IMPSCNNPoolingNode.KernelHeight]
//   - [IMPSCNNPoolingNode.KernelWidth]
//   - [IMPSCNNPoolingNode.StrideInPixelsX]
//   - [IMPSCNNPoolingNode.StrideInPixelsY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode
type IMPSCNNPoolingNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNPoolingNode
	InitWithSourceFilterSizeStride(sourceNode IMPSNNImageNode, size uint, stride uint) MPSCNNPoolingNode
	InitWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(sourceNode IMPSNNImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingNode

	// Topic: Instance Properties

	KernelHeight() uint
	KernelWidth() uint
	StrideInPixelsX() uint
	StrideInPixelsY() uint
}

// Init initializes the instance.
func (c MPSCNNPoolingNode) Init() MPSCNNPoolingNode {
	rv := objc.Send[MPSCNNPoolingNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingNode) Autorelease() MPSCNNPoolingNode {
	rv := objc.Send[MPSCNNPoolingNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingNode creates a new MPSCNNPoolingNode instance.
func NewMPSCNNPoolingNode() MPSCNNPoolingNode {
	class := getMPSCNNPoolingNodeClass()
	rv := objc.Send[MPSCNNPoolingNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:filterSize:)
func NewCNNPoolingNodeWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNPoolingNode {
	instance := getMPSCNNPoolingNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:filterSize:"), sourceNode, size)
	return MPSCNNPoolingNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:filterSize:stride:)
func NewCNNPoolingNodeWithSourceFilterSizeStride(sourceNode IMPSNNImageNode, size uint, stride uint) MPSCNNPoolingNode {
	instance := getMPSCNNPoolingNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:filterSize:stride:"), sourceNode, size, stride)
	return MPSCNNPoolingNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingNodeWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(sourceNode IMPSNNImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingNode {
	instance := getMPSCNNPoolingNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), sourceNode, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:filterSize:)
func (c MPSCNNPoolingNode) InitWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNPoolingNode {
	rv := objc.Send[MPSCNNPoolingNode](c.ID, objc.Sel("initWithSource:filterSize:"), sourceNode, size)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:filterSize:stride:)
func (c MPSCNNPoolingNode) InitWithSourceFilterSizeStride(sourceNode IMPSNNImageNode, size uint, stride uint) MPSCNNPoolingNode {
	rv := objc.Send[MPSCNNPoolingNode](c.ID, objc.Sel("initWithSource:filterSize:stride:"), sourceNode, size, stride)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func (c MPSCNNPoolingNode) InitWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(sourceNode IMPSNNImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingNode {
	rv := objc.Send[MPSCNNPoolingNode](c.ID, objc.Sel("initWithSource:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), sourceNode, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/nodeWithSource:filterSize:
func (_MPSCNNPoolingNodeClass MPSCNNPoolingNodeClass) NodeWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNPoolingNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNPoolingNodeClass.class), objc.Sel("nodeWithSource:filterSize:"), sourceNode, size)
	return MPSCNNPoolingNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/nodeWithSource:filterSize:stride:
func (_MPSCNNPoolingNodeClass MPSCNNPoolingNodeClass) NodeWithSourceFilterSizeStride(sourceNode IMPSNNImageNode, size uint, stride uint) MPSCNNPoolingNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNPoolingNodeClass.class), objc.Sel("nodeWithSource:filterSize:stride:"), sourceNode, size, stride)
	return MPSCNNPoolingNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/kernelHeight
func (c MPSCNNPoolingNode) KernelHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/kernelWidth
func (c MPSCNNPoolingNode) KernelWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelWidth"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/strideInPixelsX
func (c MPSCNNPoolingNode) StrideInPixelsX() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("strideInPixelsX"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/strideInPixelsY
func (c MPSCNNPoolingNode) StrideInPixelsY() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("strideInPixelsY"))
	return rv
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingAverageNode] class.
var (
	_MPSCNNPoolingAverageNodeClass     MPSCNNPoolingAverageNodeClass
	_MPSCNNPoolingAverageNodeClassOnce sync.Once
)

func getMPSCNNPoolingAverageNodeClass() MPSCNNPoolingAverageNodeClass {
	_MPSCNNPoolingAverageNodeClassOnce.Do(func() {
		_MPSCNNPoolingAverageNodeClass = MPSCNNPoolingAverageNodeClass{class: objc.GetClass("MPSCNNPoolingAverageNode")}
	})
	return _MPSCNNPoolingAverageNodeClass
}

// GetMPSCNNPoolingAverageNodeClass returns the class object for MPSCNNPoolingAverageNode.
func GetMPSCNNPoolingAverageNodeClass() MPSCNNPoolingAverageNodeClass {
	return getMPSCNNPoolingAverageNodeClass()
}

type MPSCNNPoolingAverageNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingAverageNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingAverageNodeClass) Alloc() MPSCNNPoolingAverageNode {
	rv := objc.Send[MPSCNNPoolingAverageNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of an average pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageNode
type MPSCNNPoolingAverageNode struct {
	MPSCNNPoolingNode
}

// MPSCNNPoolingAverageNodeFromID constructs a [MPSCNNPoolingAverageNode] from an objc.ID.
//
// A representation of an average pooling filter.
func MPSCNNPoolingAverageNodeFromID(id objc.ID) MPSCNNPoolingAverageNode {
	return MPSCNNPoolingAverageNode{MPSCNNPoolingNode: MPSCNNPoolingNodeFromID(id)}
}

// NOTE: MPSCNNPoolingAverageNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingAverageNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageNode
type IMPSCNNPoolingAverageNode interface {
	IMPSCNNPoolingNode
}

// Init initializes the instance.
func (c MPSCNNPoolingAverageNode) Init() MPSCNNPoolingAverageNode {
	rv := objc.Send[MPSCNNPoolingAverageNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingAverageNode) Autorelease() MPSCNNPoolingAverageNode {
	rv := objc.Send[MPSCNNPoolingAverageNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingAverageNode creates a new MPSCNNPoolingAverageNode instance.
func NewMPSCNNPoolingAverageNode() MPSCNNPoolingAverageNode {
	class := getMPSCNNPoolingAverageNodeClass()
	rv := objc.Send[MPSCNNPoolingAverageNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:filterSize:)
func NewCNNPoolingAverageNodeWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNPoolingAverageNode {
	instance := getMPSCNNPoolingAverageNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:filterSize:"), sourceNode, size)
	return MPSCNNPoolingAverageNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:filterSize:stride:)
func NewCNNPoolingAverageNodeWithSourceFilterSizeStride(sourceNode IMPSNNImageNode, size uint, stride uint) MPSCNNPoolingAverageNode {
	instance := getMPSCNNPoolingAverageNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:filterSize:stride:"), sourceNode, size, stride)
	return MPSCNNPoolingAverageNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingAverageNodeWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(sourceNode IMPSNNImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingAverageNode {
	instance := getMPSCNNPoolingAverageNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), sourceNode, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingAverageNodeFromID(rv)
}

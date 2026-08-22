// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingL2NormNode] class.
var (
	_MPSCNNPoolingL2NormNodeClass     MPSCNNPoolingL2NormNodeClass
	_MPSCNNPoolingL2NormNodeClassOnce sync.Once
)

func getMPSCNNPoolingL2NormNodeClass() MPSCNNPoolingL2NormNodeClass {
	_MPSCNNPoolingL2NormNodeClassOnce.Do(func() {
		_MPSCNNPoolingL2NormNodeClass = MPSCNNPoolingL2NormNodeClass{class: objc.GetClass("MPSCNNPoolingL2NormNode")}
	})
	return _MPSCNNPoolingL2NormNodeClass
}

// GetMPSCNNPoolingL2NormNodeClass returns the class object for MPSCNNPoolingL2NormNode.
func GetMPSCNNPoolingL2NormNodeClass() MPSCNNPoolingL2NormNodeClass {
	return getMPSCNNPoolingL2NormNodeClass()
}

type MPSCNNPoolingL2NormNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingL2NormNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingL2NormNodeClass) Alloc() MPSCNNPoolingL2NormNode {
	rv := objc.Send[MPSCNNPoolingL2NormNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a L2-norm pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormNode
type MPSCNNPoolingL2NormNode struct {
	MPSCNNPoolingNode
}

// MPSCNNPoolingL2NormNodeFromID constructs a [MPSCNNPoolingL2NormNode] from an objc.ID.
//
// A representation of a L2-norm pooling filter.
func MPSCNNPoolingL2NormNodeFromID(id objc.ID) MPSCNNPoolingL2NormNode {
	return MPSCNNPoolingL2NormNode{MPSCNNPoolingNode: MPSCNNPoolingNodeFromID(id)}
}

// NOTE: MPSCNNPoolingL2NormNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingL2NormNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormNode
type IMPSCNNPoolingL2NormNode interface {
	IMPSCNNPoolingNode
}

// Init initializes the instance.
func (c MPSCNNPoolingL2NormNode) Init() MPSCNNPoolingL2NormNode {
	rv := objc.Send[MPSCNNPoolingL2NormNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingL2NormNode) Autorelease() MPSCNNPoolingL2NormNode {
	rv := objc.Send[MPSCNNPoolingL2NormNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingL2NormNode creates a new MPSCNNPoolingL2NormNode instance.
func NewMPSCNNPoolingL2NormNode() MPSCNNPoolingL2NormNode {
	class := getMPSCNNPoolingL2NormNodeClass()
	rv := objc.Send[MPSCNNPoolingL2NormNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:filterSize:)
func NewCNNPoolingL2NormNodeWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNPoolingL2NormNode {
	instance := getMPSCNNPoolingL2NormNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:filterSize:"), sourceNode, size)
	return MPSCNNPoolingL2NormNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:filterSize:stride:)
func NewCNNPoolingL2NormNodeWithSourceFilterSizeStride(sourceNode IMPSNNImageNode, size uint, stride uint) MPSCNNPoolingL2NormNode {
	instance := getMPSCNNPoolingL2NormNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:filterSize:stride:"), sourceNode, size, stride)
	return MPSCNNPoolingL2NormNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingL2NormNodeWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(sourceNode IMPSNNImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingL2NormNode {
	instance := getMPSCNNPoolingL2NormNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), sourceNode, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingL2NormNodeFromID(rv)
}

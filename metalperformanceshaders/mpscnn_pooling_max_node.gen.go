// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNPoolingMaxNode] class.
var (
	_MPSCNNPoolingMaxNodeClass     MPSCNNPoolingMaxNodeClass
	_MPSCNNPoolingMaxNodeClassOnce sync.Once
)

func getMPSCNNPoolingMaxNodeClass() MPSCNNPoolingMaxNodeClass {
	_MPSCNNPoolingMaxNodeClassOnce.Do(func() {
		_MPSCNNPoolingMaxNodeClass = MPSCNNPoolingMaxNodeClass{class: objc.GetClass("MPSCNNPoolingMaxNode")}
	})
	return _MPSCNNPoolingMaxNodeClass
}

// GetMPSCNNPoolingMaxNodeClass returns the class object for MPSCNNPoolingMaxNode.
func GetMPSCNNPoolingMaxNodeClass() MPSCNNPoolingMaxNodeClass {
	return getMPSCNNPoolingMaxNodeClass()
}

type MPSCNNPoolingMaxNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNPoolingMaxNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNPoolingMaxNodeClass) Alloc() MPSCNNPoolingMaxNode {
	rv := objc.Send[MPSCNNPoolingMaxNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a max pooling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxNode
type MPSCNNPoolingMaxNode struct {
	MPSCNNPoolingNode
}

// MPSCNNPoolingMaxNodeFromID constructs a [MPSCNNPoolingMaxNode] from an objc.ID.
//
// A representation of a max pooling filter.
func MPSCNNPoolingMaxNodeFromID(id objc.ID) MPSCNNPoolingMaxNode {
	return MPSCNNPoolingMaxNode{MPSCNNPoolingNode: MPSCNNPoolingNodeFromID(id)}
}

// NOTE: MPSCNNPoolingMaxNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNPoolingMaxNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxNode
type IMPSCNNPoolingMaxNode interface {
	IMPSCNNPoolingNode
}

// Init initializes the instance.
func (c MPSCNNPoolingMaxNode) Init() MPSCNNPoolingMaxNode {
	rv := objc.Send[MPSCNNPoolingMaxNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNPoolingMaxNode) Autorelease() MPSCNNPoolingMaxNode {
	rv := objc.Send[MPSCNNPoolingMaxNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNPoolingMaxNode creates a new MPSCNNPoolingMaxNode instance.
func NewMPSCNNPoolingMaxNode() MPSCNNPoolingMaxNode {
	class := getMPSCNNPoolingMaxNodeClass()
	rv := objc.Send[MPSCNNPoolingMaxNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:filterSize:)
func NewCNNPoolingMaxNodeWithSourceFilterSize(sourceNode IMPSNNImageNode, size uint) MPSCNNPoolingMaxNode {
	instance := getMPSCNNPoolingMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:filterSize:"), sourceNode, size)
	return MPSCNNPoolingMaxNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:filterSize:stride:)
func NewCNNPoolingMaxNodeWithSourceFilterSizeStride(sourceNode IMPSNNImageNode, size uint, stride uint) MPSCNNPoolingMaxNode {
	instance := getMPSCNNPoolingMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:filterSize:stride:"), sourceNode, size, stride)
	return MPSCNNPoolingMaxNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingNode/init(source:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:)
func NewCNNPoolingMaxNodeWithSourceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(sourceNode IMPSNNImageNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) MPSCNNPoolingMaxNode {
	instance := getMPSCNNPoolingMaxNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), sourceNode, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return MPSCNNPoolingMaxNodeFromID(rv)
}

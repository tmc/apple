// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSpatialNormalizationNode] class.
var (
	_MPSCNNSpatialNormalizationNodeClass     MPSCNNSpatialNormalizationNodeClass
	_MPSCNNSpatialNormalizationNodeClassOnce sync.Once
)

func getMPSCNNSpatialNormalizationNodeClass() MPSCNNSpatialNormalizationNodeClass {
	_MPSCNNSpatialNormalizationNodeClassOnce.Do(func() {
		_MPSCNNSpatialNormalizationNodeClass = MPSCNNSpatialNormalizationNodeClass{class: objc.GetClass("MPSCNNSpatialNormalizationNode")}
	})
	return _MPSCNNSpatialNormalizationNodeClass
}

// GetMPSCNNSpatialNormalizationNodeClass returns the class object for MPSCNNSpatialNormalizationNode.
func GetMPSCNNSpatialNormalizationNodeClass() MPSCNNSpatialNormalizationNodeClass {
	return getMPSCNNSpatialNormalizationNodeClass()
}

type MPSCNNSpatialNormalizationNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSpatialNormalizationNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSpatialNormalizationNodeClass) Alloc() MPSCNNSpatialNormalizationNode {
	rv := objc.Send[MPSCNNSpatialNormalizationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a spatial normalization kernel.
//
// # Initializers
//
//   - [MPSCNNSpatialNormalizationNode.InitWithSourceKernelSize]
//
// # Instance Properties
//
//   - [MPSCNNSpatialNormalizationNode.KernelHeight]
//   - [MPSCNNSpatialNormalizationNode.SetKernelHeight]
//   - [MPSCNNSpatialNormalizationNode.KernelWidth]
//   - [MPSCNNSpatialNormalizationNode.SetKernelWidth]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationNode
type MPSCNNSpatialNormalizationNode struct {
	MPSCNNNormalizationNode
}

// MPSCNNSpatialNormalizationNodeFromID constructs a [MPSCNNSpatialNormalizationNode] from an objc.ID.
//
// A representation of a spatial normalization kernel.
func MPSCNNSpatialNormalizationNodeFromID(id objc.ID) MPSCNNSpatialNormalizationNode {
	return MPSCNNSpatialNormalizationNode{MPSCNNNormalizationNode: MPSCNNNormalizationNodeFromID(id)}
}

// NOTE: MPSCNNSpatialNormalizationNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSpatialNormalizationNode] class.
//
// # Initializers
//
//   - [IMPSCNNSpatialNormalizationNode.InitWithSourceKernelSize]
//
// # Instance Properties
//
//   - [IMPSCNNSpatialNormalizationNode.KernelHeight]
//   - [IMPSCNNSpatialNormalizationNode.SetKernelHeight]
//   - [IMPSCNNSpatialNormalizationNode.KernelWidth]
//   - [IMPSCNNSpatialNormalizationNode.SetKernelWidth]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationNode
type IMPSCNNSpatialNormalizationNode interface {
	IMPSCNNNormalizationNode

	// Topic: Initializers

	InitWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNSpatialNormalizationNode

	// Topic: Instance Properties

	KernelHeight() uint
	SetKernelHeight(value uint)
	KernelWidth() uint
	SetKernelWidth(value uint)
}

// Init initializes the instance.
func (c MPSCNNSpatialNormalizationNode) Init() MPSCNNSpatialNormalizationNode {
	rv := objc.Send[MPSCNNSpatialNormalizationNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSpatialNormalizationNode) Autorelease() MPSCNNSpatialNormalizationNode {
	rv := objc.Send[MPSCNNSpatialNormalizationNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSpatialNormalizationNode creates a new MPSCNNSpatialNormalizationNode instance.
func NewMPSCNNSpatialNormalizationNode() MPSCNNSpatialNormalizationNode {
	class := getMPSCNNSpatialNormalizationNodeClass()
	rv := objc.Send[MPSCNNSpatialNormalizationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationNode/init(source:)
func NewCNNSpatialNormalizationNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNSpatialNormalizationNode {
	instance := getMPSCNNSpatialNormalizationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNSpatialNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationNode/init(source:kernelSize:)
func NewCNNSpatialNormalizationNodeWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNSpatialNormalizationNode {
	instance := getMPSCNNSpatialNormalizationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:kernelSize:"), sourceNode, kernelSize)
	return MPSCNNSpatialNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationNode/init(source:kernelSize:)
func (c MPSCNNSpatialNormalizationNode) InitWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNSpatialNormalizationNode {
	rv := objc.Send[MPSCNNSpatialNormalizationNode](c.ID, objc.Sel("initWithSource:kernelSize:"), sourceNode, kernelSize)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationNode/nodeWithSource:kernelSize:
func (_MPSCNNSpatialNormalizationNodeClass MPSCNNSpatialNormalizationNodeClass) NodeWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNSpatialNormalizationNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNSpatialNormalizationNodeClass.class), objc.Sel("nodeWithSource:kernelSize:"), sourceNode, kernelSize)
	return MPSCNNSpatialNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationNode/kernelHeight
func (c MPSCNNSpatialNormalizationNode) KernelHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelHeight"))
	return rv
}
func (c MPSCNNSpatialNormalizationNode) SetKernelHeight(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelHeight:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationNode/kernelWidth
func (c MPSCNNSpatialNormalizationNode) KernelWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelWidth"))
	return rv
}
func (c MPSCNNSpatialNormalizationNode) SetKernelWidth(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelWidth:"), value)
}

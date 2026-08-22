// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNCrossChannelNormalizationNode] class.
var (
	_MPSCNNCrossChannelNormalizationNodeClass     MPSCNNCrossChannelNormalizationNodeClass
	_MPSCNNCrossChannelNormalizationNodeClassOnce sync.Once
)

func getMPSCNNCrossChannelNormalizationNodeClass() MPSCNNCrossChannelNormalizationNodeClass {
	_MPSCNNCrossChannelNormalizationNodeClassOnce.Do(func() {
		_MPSCNNCrossChannelNormalizationNodeClass = MPSCNNCrossChannelNormalizationNodeClass{class: objc.GetClass("MPSCNNCrossChannelNormalizationNode")}
	})
	return _MPSCNNCrossChannelNormalizationNodeClass
}

// GetMPSCNNCrossChannelNormalizationNodeClass returns the class object for MPSCNNCrossChannelNormalizationNode.
func GetMPSCNNCrossChannelNormalizationNodeClass() MPSCNNCrossChannelNormalizationNodeClass {
	return getMPSCNNCrossChannelNormalizationNodeClass()
}

type MPSCNNCrossChannelNormalizationNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNCrossChannelNormalizationNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNCrossChannelNormalizationNodeClass) Alloc() MPSCNNCrossChannelNormalizationNode {
	rv := objc.Send[MPSCNNCrossChannelNormalizationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a normalization kernel across feature channels.
//
// # Initializers
//
//   - [MPSCNNCrossChannelNormalizationNode.InitWithSourceKernelSize]
//
// # Instance Properties
//
//   - [MPSCNNCrossChannelNormalizationNode.KernelSizeInFeatureChannels]
//   - [MPSCNNCrossChannelNormalizationNode.SetKernelSizeInFeatureChannels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationNode
type MPSCNNCrossChannelNormalizationNode struct {
	MPSCNNNormalizationNode
}

// MPSCNNCrossChannelNormalizationNodeFromID constructs a [MPSCNNCrossChannelNormalizationNode] from an objc.ID.
//
// A representation of a normalization kernel across feature channels.
func MPSCNNCrossChannelNormalizationNodeFromID(id objc.ID) MPSCNNCrossChannelNormalizationNode {
	return MPSCNNCrossChannelNormalizationNode{MPSCNNNormalizationNode: MPSCNNNormalizationNodeFromID(id)}
}

// NOTE: MPSCNNCrossChannelNormalizationNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNCrossChannelNormalizationNode] class.
//
// # Initializers
//
//   - [IMPSCNNCrossChannelNormalizationNode.InitWithSourceKernelSize]
//
// # Instance Properties
//
//   - [IMPSCNNCrossChannelNormalizationNode.KernelSizeInFeatureChannels]
//   - [IMPSCNNCrossChannelNormalizationNode.SetKernelSizeInFeatureChannels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationNode
type IMPSCNNCrossChannelNormalizationNode interface {
	IMPSCNNNormalizationNode

	// Topic: Initializers

	InitWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNCrossChannelNormalizationNode

	// Topic: Instance Properties

	KernelSizeInFeatureChannels() uint
	SetKernelSizeInFeatureChannels(value uint)
}

// Init initializes the instance.
func (c MPSCNNCrossChannelNormalizationNode) Init() MPSCNNCrossChannelNormalizationNode {
	rv := objc.Send[MPSCNNCrossChannelNormalizationNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNCrossChannelNormalizationNode) Autorelease() MPSCNNCrossChannelNormalizationNode {
	rv := objc.Send[MPSCNNCrossChannelNormalizationNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNCrossChannelNormalizationNode creates a new MPSCNNCrossChannelNormalizationNode instance.
func NewMPSCNNCrossChannelNormalizationNode() MPSCNNCrossChannelNormalizationNode {
	class := getMPSCNNCrossChannelNormalizationNodeClass()
	rv := objc.Send[MPSCNNCrossChannelNormalizationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationNode/init(source:)
func NewCNNCrossChannelNormalizationNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNCrossChannelNormalizationNode {
	instance := getMPSCNNCrossChannelNormalizationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNCrossChannelNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationNode/init(source:kernelSize:)
func NewCNNCrossChannelNormalizationNodeWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNCrossChannelNormalizationNode {
	instance := getMPSCNNCrossChannelNormalizationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:kernelSize:"), sourceNode, kernelSize)
	return MPSCNNCrossChannelNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationNode/init(source:kernelSize:)
func (c MPSCNNCrossChannelNormalizationNode) InitWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNCrossChannelNormalizationNode {
	rv := objc.Send[MPSCNNCrossChannelNormalizationNode](c.ID, objc.Sel("initWithSource:kernelSize:"), sourceNode, kernelSize)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationNode/nodeWithSource:kernelSize:
func (_MPSCNNCrossChannelNormalizationNodeClass MPSCNNCrossChannelNormalizationNodeClass) NodeWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNCrossChannelNormalizationNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNCrossChannelNormalizationNodeClass.class), objc.Sel("nodeWithSource:kernelSize:"), sourceNode, kernelSize)
	return MPSCNNCrossChannelNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationNode/kernelSizeInFeatureChannels
func (c MPSCNNCrossChannelNormalizationNode) KernelSizeInFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelSizeInFeatureChannels"))
	return rv
}
func (c MPSCNNCrossChannelNormalizationNode) SetKernelSizeInFeatureChannels(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelSizeInFeatureChannels:"), value)
}

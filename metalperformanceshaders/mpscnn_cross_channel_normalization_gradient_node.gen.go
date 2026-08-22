// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNCrossChannelNormalizationGradientNode] class.
var (
	_MPSCNNCrossChannelNormalizationGradientNodeClass     MPSCNNCrossChannelNormalizationGradientNodeClass
	_MPSCNNCrossChannelNormalizationGradientNodeClassOnce sync.Once
)

func getMPSCNNCrossChannelNormalizationGradientNodeClass() MPSCNNCrossChannelNormalizationGradientNodeClass {
	_MPSCNNCrossChannelNormalizationGradientNodeClassOnce.Do(func() {
		_MPSCNNCrossChannelNormalizationGradientNodeClass = MPSCNNCrossChannelNormalizationGradientNodeClass{class: objc.GetClass("MPSCNNCrossChannelNormalizationGradientNode")}
	})
	return _MPSCNNCrossChannelNormalizationGradientNodeClass
}

// GetMPSCNNCrossChannelNormalizationGradientNodeClass returns the class object for MPSCNNCrossChannelNormalizationGradientNode.
func GetMPSCNNCrossChannelNormalizationGradientNodeClass() MPSCNNCrossChannelNormalizationGradientNodeClass {
	return getMPSCNNCrossChannelNormalizationGradientNodeClass()
}

type MPSCNNCrossChannelNormalizationGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNCrossChannelNormalizationGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNCrossChannelNormalizationGradientNodeClass) Alloc() MPSCNNCrossChannelNormalizationGradientNode {
	rv := objc.Send[MPSCNNCrossChannelNormalizationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient normalization kernel applied across feature
// channels.
//
// # Initializers
//
//   - [MPSCNNCrossChannelNormalizationGradientNode.InitWithSourceGradientSourceImageGradientStateKernelSize]
//
// # Instance Properties
//
//   - [MPSCNNCrossChannelNormalizationGradientNode.KernelSize]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradientNode
type MPSCNNCrossChannelNormalizationGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNCrossChannelNormalizationGradientNodeFromID constructs a [MPSCNNCrossChannelNormalizationGradientNode] from an objc.ID.
//
// A representation of a gradient normalization kernel applied across feature
// channels.
func MPSCNNCrossChannelNormalizationGradientNodeFromID(id objc.ID) MPSCNNCrossChannelNormalizationGradientNode {
	return MPSCNNCrossChannelNormalizationGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNCrossChannelNormalizationGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNCrossChannelNormalizationGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNCrossChannelNormalizationGradientNode.InitWithSourceGradientSourceImageGradientStateKernelSize]
//
// # Instance Properties
//
//   - [IMPSCNNCrossChannelNormalizationGradientNode.KernelSize]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradientNode
type IMPSCNNCrossChannelNormalizationGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelSize uint) MPSCNNCrossChannelNormalizationGradientNode

	// Topic: Instance Properties

	KernelSize() uint
}

// Init initializes the instance.
func (c MPSCNNCrossChannelNormalizationGradientNode) Init() MPSCNNCrossChannelNormalizationGradientNode {
	rv := objc.Send[MPSCNNCrossChannelNormalizationGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNCrossChannelNormalizationGradientNode) Autorelease() MPSCNNCrossChannelNormalizationGradientNode {
	rv := objc.Send[MPSCNNCrossChannelNormalizationGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNCrossChannelNormalizationGradientNode creates a new MPSCNNCrossChannelNormalizationGradientNode instance.
func NewMPSCNNCrossChannelNormalizationGradientNode() MPSCNNCrossChannelNormalizationGradientNode {
	class := getMPSCNNCrossChannelNormalizationGradientNodeClass()
	rv := objc.Send[MPSCNNCrossChannelNormalizationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:kernelSize:)
func NewCNNCrossChannelNormalizationGradientNodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelSize uint) MPSCNNCrossChannelNormalizationGradientNode {
	instance := getMPSCNNCrossChannelNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelSize:"), sourceGradient, sourceImage, gradientState, kernelSize)
	return MPSCNNCrossChannelNormalizationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:kernelSize:)
func (c MPSCNNCrossChannelNormalizationGradientNode) InitWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelSize uint) MPSCNNCrossChannelNormalizationGradientNode {
	rv := objc.Send[MPSCNNCrossChannelNormalizationGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelSize:"), sourceGradient, sourceImage, gradientState, kernelSize)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:kernelSize:
func (_MPSCNNCrossChannelNormalizationGradientNodeClass MPSCNNCrossChannelNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelSize uint) MPSCNNCrossChannelNormalizationGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNCrossChannelNormalizationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelSize:"), sourceGradient, sourceImage, gradientState, kernelSize)
	return MPSCNNCrossChannelNormalizationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalizationGradientNode/kernelSize
func (c MPSCNNCrossChannelNormalizationGradientNode) KernelSize() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelSize"))
	return rv
}

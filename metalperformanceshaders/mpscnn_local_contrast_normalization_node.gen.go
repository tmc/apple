// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNLocalContrastNormalizationNode] class.
var (
	_MPSCNNLocalContrastNormalizationNodeClass     MPSCNNLocalContrastNormalizationNodeClass
	_MPSCNNLocalContrastNormalizationNodeClassOnce sync.Once
)

func getMPSCNNLocalContrastNormalizationNodeClass() MPSCNNLocalContrastNormalizationNodeClass {
	_MPSCNNLocalContrastNormalizationNodeClassOnce.Do(func() {
		_MPSCNNLocalContrastNormalizationNodeClass = MPSCNNLocalContrastNormalizationNodeClass{class: objc.GetClass("MPSCNNLocalContrastNormalizationNode")}
	})
	return _MPSCNNLocalContrastNormalizationNodeClass
}

// GetMPSCNNLocalContrastNormalizationNodeClass returns the class object for MPSCNNLocalContrastNormalizationNode.
func GetMPSCNNLocalContrastNormalizationNodeClass() MPSCNNLocalContrastNormalizationNodeClass {
	return getMPSCNNLocalContrastNormalizationNodeClass()
}

type MPSCNNLocalContrastNormalizationNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLocalContrastNormalizationNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLocalContrastNormalizationNodeClass) Alloc() MPSCNNLocalContrastNormalizationNode {
	rv := objc.Send[MPSCNNLocalContrastNormalizationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a local-contrast normalization kernel.
//
// # Initializers
//
//   - [MPSCNNLocalContrastNormalizationNode.InitWithSourceKernelSize]
//
// # Instance Properties
//
//   - [MPSCNNLocalContrastNormalizationNode.KernelHeight]
//   - [MPSCNNLocalContrastNormalizationNode.SetKernelHeight]
//   - [MPSCNNLocalContrastNormalizationNode.KernelWidth]
//   - [MPSCNNLocalContrastNormalizationNode.SetKernelWidth]
//   - [MPSCNNLocalContrastNormalizationNode.P0]
//   - [MPSCNNLocalContrastNormalizationNode.SetP0]
//   - [MPSCNNLocalContrastNormalizationNode.Pm]
//   - [MPSCNNLocalContrastNormalizationNode.SetPm]
//   - [MPSCNNLocalContrastNormalizationNode.Ps]
//   - [MPSCNNLocalContrastNormalizationNode.SetPs]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode
type MPSCNNLocalContrastNormalizationNode struct {
	MPSCNNNormalizationNode
}

// MPSCNNLocalContrastNormalizationNodeFromID constructs a [MPSCNNLocalContrastNormalizationNode] from an objc.ID.
//
// A representation of a local-contrast normalization kernel.
func MPSCNNLocalContrastNormalizationNodeFromID(id objc.ID) MPSCNNLocalContrastNormalizationNode {
	return MPSCNNLocalContrastNormalizationNode{MPSCNNNormalizationNode: MPSCNNNormalizationNodeFromID(id)}
}

// NOTE: MPSCNNLocalContrastNormalizationNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLocalContrastNormalizationNode] class.
//
// # Initializers
//
//   - [IMPSCNNLocalContrastNormalizationNode.InitWithSourceKernelSize]
//
// # Instance Properties
//
//   - [IMPSCNNLocalContrastNormalizationNode.KernelHeight]
//   - [IMPSCNNLocalContrastNormalizationNode.SetKernelHeight]
//   - [IMPSCNNLocalContrastNormalizationNode.KernelWidth]
//   - [IMPSCNNLocalContrastNormalizationNode.SetKernelWidth]
//   - [IMPSCNNLocalContrastNormalizationNode.P0]
//   - [IMPSCNNLocalContrastNormalizationNode.SetP0]
//   - [IMPSCNNLocalContrastNormalizationNode.Pm]
//   - [IMPSCNNLocalContrastNormalizationNode.SetPm]
//   - [IMPSCNNLocalContrastNormalizationNode.Ps]
//   - [IMPSCNNLocalContrastNormalizationNode.SetPs]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode
type IMPSCNNLocalContrastNormalizationNode interface {
	IMPSCNNNormalizationNode

	// Topic: Initializers

	InitWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNLocalContrastNormalizationNode

	// Topic: Instance Properties

	KernelHeight() uint
	SetKernelHeight(value uint)
	KernelWidth() uint
	SetKernelWidth(value uint)
	P0() float32
	SetP0(value float32)
	Pm() float32
	SetPm(value float32)
	Ps() float32
	SetPs(value float32)
}

// Init initializes the instance.
func (c MPSCNNLocalContrastNormalizationNode) Init() MPSCNNLocalContrastNormalizationNode {
	rv := objc.Send[MPSCNNLocalContrastNormalizationNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLocalContrastNormalizationNode) Autorelease() MPSCNNLocalContrastNormalizationNode {
	rv := objc.Send[MPSCNNLocalContrastNormalizationNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLocalContrastNormalizationNode creates a new MPSCNNLocalContrastNormalizationNode instance.
func NewMPSCNNLocalContrastNormalizationNode() MPSCNNLocalContrastNormalizationNode {
	class := getMPSCNNLocalContrastNormalizationNodeClass()
	rv := objc.Send[MPSCNNLocalContrastNormalizationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode/init(source:)
func NewCNNLocalContrastNormalizationNodeWithSource(sourceNode IMPSNNImageNode) MPSCNNLocalContrastNormalizationNode {
	instance := getMPSCNNLocalContrastNormalizationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	return MPSCNNLocalContrastNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode/init(source:kernelSize:)
func NewCNNLocalContrastNormalizationNodeWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNLocalContrastNormalizationNode {
	instance := getMPSCNNLocalContrastNormalizationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:kernelSize:"), sourceNode, kernelSize)
	return MPSCNNLocalContrastNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode/init(source:kernelSize:)
func (c MPSCNNLocalContrastNormalizationNode) InitWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNLocalContrastNormalizationNode {
	rv := objc.Send[MPSCNNLocalContrastNormalizationNode](c.ID, objc.Sel("initWithSource:kernelSize:"), sourceNode, kernelSize)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode/nodeWithSource:kernelSize:
func (_MPSCNNLocalContrastNormalizationNodeClass MPSCNNLocalContrastNormalizationNodeClass) NodeWithSourceKernelSize(sourceNode IMPSNNImageNode, kernelSize uint) MPSCNNLocalContrastNormalizationNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNLocalContrastNormalizationNodeClass.class), objc.Sel("nodeWithSource:kernelSize:"), sourceNode, kernelSize)
	return MPSCNNLocalContrastNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode/kernelHeight
func (c MPSCNNLocalContrastNormalizationNode) KernelHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelHeight"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationNode) SetKernelHeight(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelHeight:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode/kernelWidth
func (c MPSCNNLocalContrastNormalizationNode) KernelWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelWidth"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationNode) SetKernelWidth(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelWidth:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode/p0
func (c MPSCNNLocalContrastNormalizationNode) P0() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("p0"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationNode) SetP0(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setP0:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode/pm
func (c MPSCNNLocalContrastNormalizationNode) Pm() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("pm"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationNode) SetPm(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPm:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationNode/ps
func (c MPSCNNLocalContrastNormalizationNode) Ps() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("ps"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationNode) SetPs(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPs:"), value)
}

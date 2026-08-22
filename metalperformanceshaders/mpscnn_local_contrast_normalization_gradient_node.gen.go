// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNLocalContrastNormalizationGradientNode] class.
var (
	_MPSCNNLocalContrastNormalizationGradientNodeClass     MPSCNNLocalContrastNormalizationGradientNodeClass
	_MPSCNNLocalContrastNormalizationGradientNodeClassOnce sync.Once
)

func getMPSCNNLocalContrastNormalizationGradientNodeClass() MPSCNNLocalContrastNormalizationGradientNodeClass {
	_MPSCNNLocalContrastNormalizationGradientNodeClassOnce.Do(func() {
		_MPSCNNLocalContrastNormalizationGradientNodeClass = MPSCNNLocalContrastNormalizationGradientNodeClass{class: objc.GetClass("MPSCNNLocalContrastNormalizationGradientNode")}
	})
	return _MPSCNNLocalContrastNormalizationGradientNodeClass
}

// GetMPSCNNLocalContrastNormalizationGradientNodeClass returns the class object for MPSCNNLocalContrastNormalizationGradientNode.
func GetMPSCNNLocalContrastNormalizationGradientNodeClass() MPSCNNLocalContrastNormalizationGradientNodeClass {
	return getMPSCNNLocalContrastNormalizationGradientNodeClass()
}

type MPSCNNLocalContrastNormalizationGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLocalContrastNormalizationGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLocalContrastNormalizationGradientNodeClass) Alloc() MPSCNNLocalContrastNormalizationGradientNode {
	rv := objc.Send[MPSCNNLocalContrastNormalizationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient local-contrast normalization kernel.
//
// # Initializers
//
//   - [MPSCNNLocalContrastNormalizationGradientNode.InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight]
//
// # Instance Properties
//
//   - [MPSCNNLocalContrastNormalizationGradientNode.Alpha]
//   - [MPSCNNLocalContrastNormalizationGradientNode.SetAlpha]
//   - [MPSCNNLocalContrastNormalizationGradientNode.Beta]
//   - [MPSCNNLocalContrastNormalizationGradientNode.SetBeta]
//   - [MPSCNNLocalContrastNormalizationGradientNode.Delta]
//   - [MPSCNNLocalContrastNormalizationGradientNode.SetDelta]
//   - [MPSCNNLocalContrastNormalizationGradientNode.KernelHeight]
//   - [MPSCNNLocalContrastNormalizationGradientNode.KernelWidth]
//   - [MPSCNNLocalContrastNormalizationGradientNode.P0]
//   - [MPSCNNLocalContrastNormalizationGradientNode.SetP0]
//   - [MPSCNNLocalContrastNormalizationGradientNode.Pm]
//   - [MPSCNNLocalContrastNormalizationGradientNode.SetPm]
//   - [MPSCNNLocalContrastNormalizationGradientNode.Ps]
//   - [MPSCNNLocalContrastNormalizationGradientNode.SetPs]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode
type MPSCNNLocalContrastNormalizationGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNLocalContrastNormalizationGradientNodeFromID constructs a [MPSCNNLocalContrastNormalizationGradientNode] from an objc.ID.
//
// A representation of a gradient local-contrast normalization kernel.
func MPSCNNLocalContrastNormalizationGradientNodeFromID(id objc.ID) MPSCNNLocalContrastNormalizationGradientNode {
	return MPSCNNLocalContrastNormalizationGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNLocalContrastNormalizationGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLocalContrastNormalizationGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNLocalContrastNormalizationGradientNode.InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight]
//
// # Instance Properties
//
//   - [IMPSCNNLocalContrastNormalizationGradientNode.Alpha]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.SetAlpha]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.Beta]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.SetBeta]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.Delta]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.SetDelta]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.KernelHeight]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.KernelWidth]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.P0]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.SetP0]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.Pm]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.SetPm]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.Ps]
//   - [IMPSCNNLocalContrastNormalizationGradientNode.SetPs]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode
type IMPSCNNLocalContrastNormalizationGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint) MPSCNNLocalContrastNormalizationGradientNode

	// Topic: Instance Properties

	Alpha() float32
	SetAlpha(value float32)
	Beta() float32
	SetBeta(value float32)
	Delta() float32
	SetDelta(value float32)
	KernelHeight() uint
	KernelWidth() uint
	P0() float32
	SetP0(value float32)
	Pm() float32
	SetPm(value float32)
	Ps() float32
	SetPs(value float32)
}

// Init initializes the instance.
func (c MPSCNNLocalContrastNormalizationGradientNode) Init() MPSCNNLocalContrastNormalizationGradientNode {
	rv := objc.Send[MPSCNNLocalContrastNormalizationGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLocalContrastNormalizationGradientNode) Autorelease() MPSCNNLocalContrastNormalizationGradientNode {
	rv := objc.Send[MPSCNNLocalContrastNormalizationGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLocalContrastNormalizationGradientNode creates a new MPSCNNLocalContrastNormalizationGradientNode instance.
func NewMPSCNNLocalContrastNormalizationGradientNode() MPSCNNLocalContrastNormalizationGradientNode {
	class := getMPSCNNLocalContrastNormalizationGradientNodeClass()
	rv := objc.Send[MPSCNNLocalContrastNormalizationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:)
func NewCNNLocalContrastNormalizationGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint) MPSCNNLocalContrastNormalizationGradientNode {
	instance := getMPSCNNLocalContrastNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight)
	return MPSCNNLocalContrastNormalizationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:)
func (c MPSCNNLocalContrastNormalizationGradientNode) InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint) MPSCNNLocalContrastNormalizationGradientNode {
	rv := objc.Send[MPSCNNLocalContrastNormalizationGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:
func (_MPSCNNLocalContrastNormalizationGradientNodeClass MPSCNNLocalContrastNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelWidth uint, kernelHeight uint) MPSCNNLocalContrastNormalizationGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNLocalContrastNormalizationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight)
	return MPSCNNLocalContrastNormalizationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/alpha
func (c MPSCNNLocalContrastNormalizationGradientNode) Alpha() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("alpha"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradientNode) SetAlpha(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/beta
func (c MPSCNNLocalContrastNormalizationGradientNode) Beta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("beta"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradientNode) SetBeta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBeta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/delta
func (c MPSCNNLocalContrastNormalizationGradientNode) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradientNode) SetDelta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/kernelHeight
func (c MPSCNNLocalContrastNormalizationGradientNode) KernelHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/kernelWidth
func (c MPSCNNLocalContrastNormalizationGradientNode) KernelWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelWidth"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/p0
func (c MPSCNNLocalContrastNormalizationGradientNode) P0() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("p0"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradientNode) SetP0(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setP0:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/pm
func (c MPSCNNLocalContrastNormalizationGradientNode) Pm() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("pm"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradientNode) SetPm(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPm:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalizationGradientNode/ps
func (c MPSCNNLocalContrastNormalizationGradientNode) Ps() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("ps"))
	return rv
}
func (c MPSCNNLocalContrastNormalizationGradientNode) SetPs(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPs:"), value)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSpatialNormalizationGradientNode] class.
var (
	_MPSCNNSpatialNormalizationGradientNodeClass     MPSCNNSpatialNormalizationGradientNodeClass
	_MPSCNNSpatialNormalizationGradientNodeClassOnce sync.Once
)

func getMPSCNNSpatialNormalizationGradientNodeClass() MPSCNNSpatialNormalizationGradientNodeClass {
	_MPSCNNSpatialNormalizationGradientNodeClassOnce.Do(func() {
		_MPSCNNSpatialNormalizationGradientNodeClass = MPSCNNSpatialNormalizationGradientNodeClass{class: objc.GetClass("MPSCNNSpatialNormalizationGradientNode")}
	})
	return _MPSCNNSpatialNormalizationGradientNodeClass
}

// GetMPSCNNSpatialNormalizationGradientNodeClass returns the class object for MPSCNNSpatialNormalizationGradientNode.
func GetMPSCNNSpatialNormalizationGradientNodeClass() MPSCNNSpatialNormalizationGradientNodeClass {
	return getMPSCNNSpatialNormalizationGradientNodeClass()
}

type MPSCNNSpatialNormalizationGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSpatialNormalizationGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSpatialNormalizationGradientNodeClass) Alloc() MPSCNNSpatialNormalizationGradientNode {
	rv := objc.Send[MPSCNNSpatialNormalizationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient spatial normalization kernel.
//
// # Initializers
//
//   - [MPSCNNSpatialNormalizationGradientNode.InitWithSourceGradientSourceImageGradientStateKernelSize]
//
// # Instance Properties
//
//   - [MPSCNNSpatialNormalizationGradientNode.Alpha]
//   - [MPSCNNSpatialNormalizationGradientNode.SetAlpha]
//   - [MPSCNNSpatialNormalizationGradientNode.Beta]
//   - [MPSCNNSpatialNormalizationGradientNode.SetBeta]
//   - [MPSCNNSpatialNormalizationGradientNode.Delta]
//   - [MPSCNNSpatialNormalizationGradientNode.SetDelta]
//   - [MPSCNNSpatialNormalizationGradientNode.KernelHeight]
//   - [MPSCNNSpatialNormalizationGradientNode.SetKernelHeight]
//   - [MPSCNNSpatialNormalizationGradientNode.KernelWidth]
//   - [MPSCNNSpatialNormalizationGradientNode.SetKernelWidth]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode
type MPSCNNSpatialNormalizationGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNSpatialNormalizationGradientNodeFromID constructs a [MPSCNNSpatialNormalizationGradientNode] from an objc.ID.
//
// A representation of a gradient spatial normalization kernel.
func MPSCNNSpatialNormalizationGradientNodeFromID(id objc.ID) MPSCNNSpatialNormalizationGradientNode {
	return MPSCNNSpatialNormalizationGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNSpatialNormalizationGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSpatialNormalizationGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNSpatialNormalizationGradientNode.InitWithSourceGradientSourceImageGradientStateKernelSize]
//
// # Instance Properties
//
//   - [IMPSCNNSpatialNormalizationGradientNode.Alpha]
//   - [IMPSCNNSpatialNormalizationGradientNode.SetAlpha]
//   - [IMPSCNNSpatialNormalizationGradientNode.Beta]
//   - [IMPSCNNSpatialNormalizationGradientNode.SetBeta]
//   - [IMPSCNNSpatialNormalizationGradientNode.Delta]
//   - [IMPSCNNSpatialNormalizationGradientNode.SetDelta]
//   - [IMPSCNNSpatialNormalizationGradientNode.KernelHeight]
//   - [IMPSCNNSpatialNormalizationGradientNode.SetKernelHeight]
//   - [IMPSCNNSpatialNormalizationGradientNode.KernelWidth]
//   - [IMPSCNNSpatialNormalizationGradientNode.SetKernelWidth]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode
type IMPSCNNSpatialNormalizationGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelSize uint) MPSCNNSpatialNormalizationGradientNode

	// Topic: Instance Properties

	Alpha() float32
	SetAlpha(value float32)
	Beta() float32
	SetBeta(value float32)
	Delta() float32
	SetDelta(value float32)
	KernelHeight() uint
	SetKernelHeight(value uint)
	KernelWidth() uint
	SetKernelWidth(value uint)
}

// Init initializes the instance.
func (c MPSCNNSpatialNormalizationGradientNode) Init() MPSCNNSpatialNormalizationGradientNode {
	rv := objc.Send[MPSCNNSpatialNormalizationGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSpatialNormalizationGradientNode) Autorelease() MPSCNNSpatialNormalizationGradientNode {
	rv := objc.Send[MPSCNNSpatialNormalizationGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSpatialNormalizationGradientNode creates a new MPSCNNSpatialNormalizationGradientNode instance.
func NewMPSCNNSpatialNormalizationGradientNode() MPSCNNSpatialNormalizationGradientNode {
	class := getMPSCNNSpatialNormalizationGradientNodeClass()
	rv := objc.Send[MPSCNNSpatialNormalizationGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:kernelSize:)
func NewCNNSpatialNormalizationGradientNodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelSize uint) MPSCNNSpatialNormalizationGradientNode {
	instance := getMPSCNNSpatialNormalizationGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelSize:"), sourceGradient, sourceImage, gradientState, kernelSize)
	return MPSCNNSpatialNormalizationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode/init(sourceGradient:sourceImage:gradientState:kernelSize:)
func (c MPSCNNSpatialNormalizationGradientNode) InitWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelSize uint) MPSCNNSpatialNormalizationGradientNode {
	rv := objc.Send[MPSCNNSpatialNormalizationGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelSize:"), sourceGradient, sourceImage, gradientState, kernelSize)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode/nodeWithSourceGradient:sourceImage:gradientState:kernelSize:
func (_MPSCNNSpatialNormalizationGradientNodeClass MPSCNNSpatialNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, kernelSize uint) MPSCNNSpatialNormalizationGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNSpatialNormalizationGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelSize:"), sourceGradient, sourceImage, gradientState, kernelSize)
	return MPSCNNSpatialNormalizationGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode/alpha
func (c MPSCNNSpatialNormalizationGradientNode) Alpha() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("alpha"))
	return rv
}
func (c MPSCNNSpatialNormalizationGradientNode) SetAlpha(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode/beta
func (c MPSCNNSpatialNormalizationGradientNode) Beta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("beta"))
	return rv
}
func (c MPSCNNSpatialNormalizationGradientNode) SetBeta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBeta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode/delta
func (c MPSCNNSpatialNormalizationGradientNode) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}
func (c MPSCNNSpatialNormalizationGradientNode) SetDelta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelta:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode/kernelHeight
func (c MPSCNNSpatialNormalizationGradientNode) KernelHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelHeight"))
	return rv
}
func (c MPSCNNSpatialNormalizationGradientNode) SetKernelHeight(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelHeight:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalizationGradientNode/kernelWidth
func (c MPSCNNSpatialNormalizationGradientNode) KernelWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelWidth"))
	return rv
}
func (c MPSCNNSpatialNormalizationGradientNode) SetKernelWidth(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelWidth:"), value)
}

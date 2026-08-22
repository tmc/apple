// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNDropoutGradientNode] class.
var (
	_MPSCNNDropoutGradientNodeClass     MPSCNNDropoutGradientNodeClass
	_MPSCNNDropoutGradientNodeClassOnce sync.Once
)

func getMPSCNNDropoutGradientNodeClass() MPSCNNDropoutGradientNodeClass {
	_MPSCNNDropoutGradientNodeClassOnce.Do(func() {
		_MPSCNNDropoutGradientNodeClass = MPSCNNDropoutGradientNodeClass{class: objc.GetClass("MPSCNNDropoutGradientNode")}
	})
	return _MPSCNNDropoutGradientNodeClass
}

// GetMPSCNNDropoutGradientNodeClass returns the class object for MPSCNNDropoutGradientNode.
func GetMPSCNNDropoutGradientNodeClass() MPSCNNDropoutGradientNodeClass {
	return getMPSCNNDropoutGradientNodeClass()
}

type MPSCNNDropoutGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDropoutGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDropoutGradientNodeClass) Alloc() MPSCNNDropoutGradientNode {
	rv := objc.Send[MPSCNNDropoutGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a gradient dropout filter.
//
// # Initializers
//
//   - [MPSCNNDropoutGradientNode.InitWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels]
//
// # Instance Properties
//
//   - [MPSCNNDropoutGradientNode.KeepProbability]
//   - [MPSCNNDropoutGradientNode.MaskStrideInPixels]
//   - [MPSCNNDropoutGradientNode.Seed]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientNode
type MPSCNNDropoutGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSCNNDropoutGradientNodeFromID constructs a [MPSCNNDropoutGradientNode] from an objc.ID.
//
// A representation of a gradient dropout filter.
func MPSCNNDropoutGradientNodeFromID(id objc.ID) MPSCNNDropoutGradientNode {
	return MPSCNNDropoutGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSCNNDropoutGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDropoutGradientNode] class.
//
// # Initializers
//
//   - [IMPSCNNDropoutGradientNode.InitWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels]
//
// # Instance Properties
//
//   - [IMPSCNNDropoutGradientNode.KeepProbability]
//   - [IMPSCNNDropoutGradientNode.MaskStrideInPixels]
//   - [IMPSCNNDropoutGradientNode.Seed]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientNode
type IMPSCNNDropoutGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutGradientNode

	// Topic: Instance Properties

	KeepProbability() float32
	MaskStrideInPixels() metal.MTLSize
	Seed() uint
}

// Init initializes the instance.
func (c MPSCNNDropoutGradientNode) Init() MPSCNNDropoutGradientNode {
	rv := objc.Send[MPSCNNDropoutGradientNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDropoutGradientNode) Autorelease() MPSCNNDropoutGradientNode {
	rv := objc.Send[MPSCNNDropoutGradientNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDropoutGradientNode creates a new MPSCNNDropoutGradientNode instance.
func NewMPSCNNDropoutGradientNode() MPSCNNDropoutGradientNode {
	class := getMPSCNNDropoutGradientNodeClass()
	rv := objc.Send[MPSCNNDropoutGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientNode/init(sourceGradient:sourceImage:gradientState:keepProbability:seed:maskStrideInPixels:)
func NewCNNDropoutGradientNodeWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutGradientNode {
	instance := getMPSCNNDropoutGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:keepProbability:seed:maskStrideInPixels:"), sourceGradient, sourceImage, gradientState, keepProbability, seed, maskStrideInPixels)
	return MPSCNNDropoutGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientNode/init(sourceGradient:sourceImage:gradientState:keepProbability:seed:maskStrideInPixels:)
func (c MPSCNNDropoutGradientNode) InitWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutGradientNode {
	rv := objc.Send[MPSCNNDropoutGradientNode](c.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:keepProbability:seed:maskStrideInPixels:"), sourceGradient, sourceImage, gradientState, keepProbability, seed, maskStrideInPixels)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientNode/nodeWithSourceGradient:sourceImage:gradientState:keepProbability:seed:maskStrideInPixels:
func (_MPSCNNDropoutGradientNodeClass MPSCNNDropoutGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNGradientStateNode, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNDropoutGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:keepProbability:seed:maskStrideInPixels:"), sourceGradient, sourceImage, gradientState, keepProbability, seed, maskStrideInPixels)
	return MPSCNNDropoutGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientNode/keepProbability
func (c MPSCNNDropoutGradientNode) KeepProbability() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("keepProbability"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientNode/maskStrideInPixels
func (c MPSCNNDropoutGradientNode) MaskStrideInPixels() metal.MTLSize {
	rv := objc.Send[metal.MTLSize](c.ID, objc.Sel("maskStrideInPixels"))
	return metal.MTLSize(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientNode/seed
func (c MPSCNNDropoutGradientNode) Seed() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("seed"))
	return rv
}

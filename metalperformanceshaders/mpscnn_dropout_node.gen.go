// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNDropoutNode] class.
var (
	_MPSCNNDropoutNodeClass     MPSCNNDropoutNodeClass
	_MPSCNNDropoutNodeClassOnce sync.Once
)

func getMPSCNNDropoutNodeClass() MPSCNNDropoutNodeClass {
	_MPSCNNDropoutNodeClassOnce.Do(func() {
		_MPSCNNDropoutNodeClass = MPSCNNDropoutNodeClass{class: objc.GetClass("MPSCNNDropoutNode")}
	})
	return _MPSCNNDropoutNodeClass
}

// GetMPSCNNDropoutNodeClass returns the class object for MPSCNNDropoutNode.
func GetMPSCNNDropoutNodeClass() MPSCNNDropoutNodeClass {
	return getMPSCNNDropoutNodeClass()
}

type MPSCNNDropoutNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDropoutNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDropoutNodeClass) Alloc() MPSCNNDropoutNode {
	rv := objc.Send[MPSCNNDropoutNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a dropout filter.
//
// # Initializers
//
//   - [MPSCNNDropoutNode.InitWithSource]
//   - [MPSCNNDropoutNode.InitWithSourceKeepProbability]
//   - [MPSCNNDropoutNode.InitWithSourceKeepProbabilitySeedMaskStrideInPixels]
//
// # Instance Properties
//
//   - [MPSCNNDropoutNode.KeepProbability]
//   - [MPSCNNDropoutNode.MaskStrideInPixels]
//   - [MPSCNNDropoutNode.Seed]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode
type MPSCNNDropoutNode struct {
	MPSNNFilterNode
}

// MPSCNNDropoutNodeFromID constructs a [MPSCNNDropoutNode] from an objc.ID.
//
// A representation of a dropout filter.
func MPSCNNDropoutNodeFromID(id objc.ID) MPSCNNDropoutNode {
	return MPSCNNDropoutNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNDropoutNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDropoutNode] class.
//
// # Initializers
//
//   - [IMPSCNNDropoutNode.InitWithSource]
//   - [IMPSCNNDropoutNode.InitWithSourceKeepProbability]
//   - [IMPSCNNDropoutNode.InitWithSourceKeepProbabilitySeedMaskStrideInPixels]
//
// # Instance Properties
//
//   - [IMPSCNNDropoutNode.KeepProbability]
//   - [IMPSCNNDropoutNode.MaskStrideInPixels]
//   - [IMPSCNNDropoutNode.Seed]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode
type IMPSCNNDropoutNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSource(source IMPSNNImageNode) MPSCNNDropoutNode
	InitWithSourceKeepProbability(source IMPSNNImageNode, keepProbability float32) MPSCNNDropoutNode
	InitWithSourceKeepProbabilitySeedMaskStrideInPixels(source IMPSNNImageNode, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutNode

	// Topic: Instance Properties

	KeepProbability() float32
	MaskStrideInPixels() metal.MTLSize
	Seed() uint
}

// Init initializes the instance.
func (c MPSCNNDropoutNode) Init() MPSCNNDropoutNode {
	rv := objc.Send[MPSCNNDropoutNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDropoutNode) Autorelease() MPSCNNDropoutNode {
	rv := objc.Send[MPSCNNDropoutNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDropoutNode creates a new MPSCNNDropoutNode instance.
func NewMPSCNNDropoutNode() MPSCNNDropoutNode {
	class := getMPSCNNDropoutNodeClass()
	rv := objc.Send[MPSCNNDropoutNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/init(source:)
func NewCNNDropoutNodeWithSource(source IMPSNNImageNode) MPSCNNDropoutNode {
	instance := getMPSCNNDropoutNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), source)
	return MPSCNNDropoutNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/init(source:keepProbability:)
func NewCNNDropoutNodeWithSourceKeepProbability(source IMPSNNImageNode, keepProbability float32) MPSCNNDropoutNode {
	instance := getMPSCNNDropoutNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:keepProbability:"), source, keepProbability)
	return MPSCNNDropoutNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/init(source:keepProbability:seed:maskStrideInPixels:)
func NewCNNDropoutNodeWithSourceKeepProbabilitySeedMaskStrideInPixels(source IMPSNNImageNode, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutNode {
	instance := getMPSCNNDropoutNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:keepProbability:seed:maskStrideInPixels:"), source, keepProbability, seed, maskStrideInPixels)
	return MPSCNNDropoutNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/init(source:)
func (c MPSCNNDropoutNode) InitWithSource(source IMPSNNImageNode) MPSCNNDropoutNode {
	rv := objc.Send[MPSCNNDropoutNode](c.ID, objc.Sel("initWithSource:"), source)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/init(source:keepProbability:)
func (c MPSCNNDropoutNode) InitWithSourceKeepProbability(source IMPSNNImageNode, keepProbability float32) MPSCNNDropoutNode {
	rv := objc.Send[MPSCNNDropoutNode](c.ID, objc.Sel("initWithSource:keepProbability:"), source, keepProbability)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/init(source:keepProbability:seed:maskStrideInPixels:)
func (c MPSCNNDropoutNode) InitWithSourceKeepProbabilitySeedMaskStrideInPixels(source IMPSNNImageNode, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutNode {
	rv := objc.Send[MPSCNNDropoutNode](c.ID, objc.Sel("initWithSource:keepProbability:seed:maskStrideInPixels:"), source, keepProbability, seed, maskStrideInPixels)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/nodeWithSource:
func (_MPSCNNDropoutNodeClass MPSCNNDropoutNodeClass) NodeWithSource(source IMPSNNImageNode) MPSCNNDropoutNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNDropoutNodeClass.class), objc.Sel("nodeWithSource:"), source)
	return MPSCNNDropoutNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/nodeWithSource:keepProbability:
func (_MPSCNNDropoutNodeClass MPSCNNDropoutNodeClass) NodeWithSourceKeepProbability(source IMPSNNImageNode, keepProbability float32) MPSCNNDropoutNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNDropoutNodeClass.class), objc.Sel("nodeWithSource:keepProbability:"), source, keepProbability)
	return MPSCNNDropoutNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/nodeWithSource:keepProbability:seed:maskStrideInPixels:
func (_MPSCNNDropoutNodeClass MPSCNNDropoutNodeClass) NodeWithSourceKeepProbabilitySeedMaskStrideInPixels(source IMPSNNImageNode, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNDropoutNodeClass.class), objc.Sel("nodeWithSource:keepProbability:seed:maskStrideInPixels:"), source, keepProbability, seed, maskStrideInPixels)
	return MPSCNNDropoutNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/keepProbability
func (c MPSCNNDropoutNode) KeepProbability() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("keepProbability"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/maskStrideInPixels
func (c MPSCNNDropoutNode) MaskStrideInPixels() metal.MTLSize {
	rv := objc.Send[metal.MTLSize](c.ID, objc.Sel("maskStrideInPixels"))
	return metal.MTLSize(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutNode/seed
func (c MPSCNNDropoutNode) Seed() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("seed"))
	return rv
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNScaleNode] class.
var (
	_MPSNNScaleNodeClass     MPSNNScaleNodeClass
	_MPSNNScaleNodeClassOnce sync.Once
)

func getMPSNNScaleNodeClass() MPSNNScaleNodeClass {
	_MPSNNScaleNodeClassOnce.Do(func() {
		_MPSNNScaleNodeClass = MPSNNScaleNodeClass{class: objc.GetClass("MPSNNScaleNode")}
	})
	return _MPSNNScaleNodeClass
}

// GetMPSNNScaleNodeClass returns the class object for MPSNNScaleNode.
func GetMPSNNScaleNodeClass() MPSNNScaleNodeClass {
	return getMPSNNScaleNodeClass()
}

type MPSNNScaleNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNScaleNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNScaleNodeClass) Alloc() MPSNNScaleNode {
	rv := objc.Send[MPSNNScaleNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Abstract node representing an image resampling filter.
//
// # Initializers
//
//   - [MPSNNScaleNode.InitWithSourceOutputSize]
//   - [MPSNNScaleNode.InitWithSourceTransformProviderOutputSize]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode
type MPSNNScaleNode struct {
	MPSNNFilterNode
}

// MPSNNScaleNodeFromID constructs a [MPSNNScaleNode] from an objc.ID.
//
// Abstract node representing an image resampling filter.
func MPSNNScaleNodeFromID(id objc.ID) MPSNNScaleNode {
	return MPSNNScaleNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSNNScaleNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNScaleNode] class.
//
// # Initializers
//
//   - [IMPSNNScaleNode.InitWithSourceOutputSize]
//   - [IMPSNNScaleNode.InitWithSourceTransformProviderOutputSize]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode
type IMPSNNScaleNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSourceOutputSize(sourceNode IMPSNNImageNode, size metal.MTLSize) MPSNNScaleNode
	InitWithSourceTransformProviderOutputSize(sourceNode IMPSNNImageNode, transformProvider MPSImageTransformProvider, size metal.MTLSize) MPSNNScaleNode
}

// Init initializes the instance.
func (s MPSNNScaleNode) Init() MPSNNScaleNode {
	rv := objc.Send[MPSNNScaleNode](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSNNScaleNode) Autorelease() MPSNNScaleNode {
	rv := objc.Send[MPSNNScaleNode](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNScaleNode creates a new MPSNNScaleNode instance.
func NewMPSNNScaleNode() MPSNNScaleNode {
	class := getMPSNNScaleNodeClass()
	rv := objc.Send[MPSNNScaleNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode/init(source:outputSize:)
func NewScaleNodeWithSourceOutputSize(sourceNode IMPSNNImageNode, size metal.MTLSize) MPSNNScaleNode {
	instance := getMPSNNScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:outputSize:"), sourceNode, size)
	return MPSNNScaleNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode/init(source:transformProvider:outputSize:)
func NewScaleNodeWithSourceTransformProviderOutputSize(sourceNode IMPSNNImageNode, transformProvider MPSImageTransformProvider, size metal.MTLSize) MPSNNScaleNode {
	instance := getMPSNNScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:transformProvider:outputSize:"), sourceNode, transformProvider, size)
	return MPSNNScaleNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode/init(source:outputSize:)
func (s MPSNNScaleNode) InitWithSourceOutputSize(sourceNode IMPSNNImageNode, size metal.MTLSize) MPSNNScaleNode {
	rv := objc.Send[MPSNNScaleNode](s.ID, objc.Sel("initWithSource:outputSize:"), sourceNode, size)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode/init(source:transformProvider:outputSize:)
func (s MPSNNScaleNode) InitWithSourceTransformProviderOutputSize(sourceNode IMPSNNImageNode, transformProvider MPSImageTransformProvider, size metal.MTLSize) MPSNNScaleNode {
	rv := objc.Send[MPSNNScaleNode](s.ID, objc.Sel("initWithSource:transformProvider:outputSize:"), sourceNode, transformProvider, size)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode/nodeWithSource:outputSize:
func (_MPSNNScaleNodeClass MPSNNScaleNodeClass) NodeWithSourceOutputSize(sourceNode IMPSNNImageNode, size metal.MTLSize) MPSNNScaleNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNScaleNodeClass.class), objc.Sel("nodeWithSource:outputSize:"), sourceNode, size)
	return MPSNNScaleNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode/nodeWithSource:transformProvider:outputSize:
func (_MPSNNScaleNodeClass MPSNNScaleNodeClass) NodeWithSourceTransformProviderOutputSize(sourceNode IMPSNNImageNode, transformProvider MPSImageTransformProvider, size metal.MTLSize) MPSNNScaleNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNScaleNodeClass.class), objc.Sel("nodeWithSource:transformProvider:outputSize:"), sourceNode, transformProvider, size)
	return MPSNNScaleNodeFromID(rv)
}

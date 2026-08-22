// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNBilinearScaleNode] class.
var (
	_MPSNNBilinearScaleNodeClass     MPSNNBilinearScaleNodeClass
	_MPSNNBilinearScaleNodeClassOnce sync.Once
)

func getMPSNNBilinearScaleNodeClass() MPSNNBilinearScaleNodeClass {
	_MPSNNBilinearScaleNodeClassOnce.Do(func() {
		_MPSNNBilinearScaleNodeClass = MPSNNBilinearScaleNodeClass{class: objc.GetClass("MPSNNBilinearScaleNode")}
	})
	return _MPSNNBilinearScaleNodeClass
}

// GetMPSNNBilinearScaleNodeClass returns the class object for MPSNNBilinearScaleNode.
func GetMPSNNBilinearScaleNodeClass() MPSNNBilinearScaleNodeClass {
	return getMPSNNBilinearScaleNodeClass()
}

type MPSNNBilinearScaleNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNBilinearScaleNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNBilinearScaleNodeClass) Alloc() MPSNNBilinearScaleNode {
	rv := objc.Send[MPSNNBilinearScaleNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a bilinear resampling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBilinearScaleNode
type MPSNNBilinearScaleNode struct {
	MPSNNScaleNode
}

// MPSNNBilinearScaleNodeFromID constructs a [MPSNNBilinearScaleNode] from an objc.ID.
//
// A representation of a bilinear resampling filter.
func MPSNNBilinearScaleNodeFromID(id objc.ID) MPSNNBilinearScaleNode {
	return MPSNNBilinearScaleNode{MPSNNScaleNode: MPSNNScaleNodeFromID(id)}
}

// NOTE: MPSNNBilinearScaleNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNBilinearScaleNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBilinearScaleNode
type IMPSNNBilinearScaleNode interface {
	IMPSNNScaleNode
}

// Init initializes the instance.
func (b MPSNNBilinearScaleNode) Init() MPSNNBilinearScaleNode {
	rv := objc.Send[MPSNNBilinearScaleNode](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MPSNNBilinearScaleNode) Autorelease() MPSNNBilinearScaleNode {
	rv := objc.Send[MPSNNBilinearScaleNode](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNBilinearScaleNode creates a new MPSNNBilinearScaleNode instance.
func NewMPSNNBilinearScaleNode() MPSNNBilinearScaleNode {
	class := getMPSNNBilinearScaleNodeClass()
	rv := objc.Send[MPSNNBilinearScaleNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode/init(source:outputSize:)
func NewBilinearScaleNodeWithSourceOutputSize(sourceNode IMPSNNImageNode, size metal.MTLSize) MPSNNBilinearScaleNode {
	instance := getMPSNNBilinearScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:outputSize:"), sourceNode, size)
	return MPSNNBilinearScaleNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode/init(source:transformProvider:outputSize:)
func NewBilinearScaleNodeWithSourceTransformProviderOutputSize(sourceNode IMPSNNImageNode, transformProvider MPSImageTransformProvider, size metal.MTLSize) MPSNNBilinearScaleNode {
	instance := getMPSNNBilinearScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:transformProvider:outputSize:"), sourceNode, transformProvider, size)
	return MPSNNBilinearScaleNodeFromID(rv)
}

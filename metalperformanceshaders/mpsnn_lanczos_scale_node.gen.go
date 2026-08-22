// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNLanczosScaleNode] class.
var (
	_MPSNNLanczosScaleNodeClass     MPSNNLanczosScaleNodeClass
	_MPSNNLanczosScaleNodeClassOnce sync.Once
)

func getMPSNNLanczosScaleNodeClass() MPSNNLanczosScaleNodeClass {
	_MPSNNLanczosScaleNodeClassOnce.Do(func() {
		_MPSNNLanczosScaleNodeClass = MPSNNLanczosScaleNodeClass{class: objc.GetClass("MPSNNLanczosScaleNode")}
	})
	return _MPSNNLanczosScaleNodeClass
}

// GetMPSNNLanczosScaleNodeClass returns the class object for MPSNNLanczosScaleNode.
func GetMPSNNLanczosScaleNodeClass() MPSNNLanczosScaleNodeClass {
	return getMPSNNLanczosScaleNodeClass()
}

type MPSNNLanczosScaleNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNLanczosScaleNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLanczosScaleNodeClass) Alloc() MPSNNLanczosScaleNode {
	rv := objc.Send[MPSNNLanczosScaleNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a Lanczos resampling filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLanczosScaleNode
type MPSNNLanczosScaleNode struct {
	MPSNNScaleNode
}

// MPSNNLanczosScaleNodeFromID constructs a [MPSNNLanczosScaleNode] from an objc.ID.
//
// A representation of a Lanczos resampling filter.
func MPSNNLanczosScaleNodeFromID(id objc.ID) MPSNNLanczosScaleNode {
	return MPSNNLanczosScaleNode{MPSNNScaleNode: MPSNNScaleNodeFromID(id)}
}

// NOTE: MPSNNLanczosScaleNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNLanczosScaleNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLanczosScaleNode
type IMPSNNLanczosScaleNode interface {
	IMPSNNScaleNode
}

// Init initializes the instance.
func (l MPSNNLanczosScaleNode) Init() MPSNNLanczosScaleNode {
	rv := objc.Send[MPSNNLanczosScaleNode](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLanczosScaleNode) Autorelease() MPSNNLanczosScaleNode {
	rv := objc.Send[MPSNNLanczosScaleNode](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLanczosScaleNode creates a new MPSNNLanczosScaleNode instance.
func NewMPSNNLanczosScaleNode() MPSNNLanczosScaleNode {
	class := getMPSNNLanczosScaleNodeClass()
	rv := objc.Send[MPSNNLanczosScaleNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode/init(source:outputSize:)
func NewLanczosScaleNodeWithSourceOutputSize(sourceNode IMPSNNImageNode, size metal.MTLSize) MPSNNLanczosScaleNode {
	instance := getMPSNNLanczosScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:outputSize:"), sourceNode, size)
	return MPSNNLanczosScaleNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode/init(source:transformProvider:outputSize:)
func NewLanczosScaleNodeWithSourceTransformProviderOutputSize(sourceNode IMPSNNImageNode, transformProvider MPSImageTransformProvider, size metal.MTLSize) MPSNNLanczosScaleNode {
	instance := getMPSNNLanczosScaleNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:transformProvider:outputSize:"), sourceNode, transformProvider, size)
	return MPSNNLanczosScaleNodeFromID(rv)
}

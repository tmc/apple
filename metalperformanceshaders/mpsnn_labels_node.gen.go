// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNLabelsNode] class.
var (
	_MPSNNLabelsNodeClass     MPSNNLabelsNodeClass
	_MPSNNLabelsNodeClassOnce sync.Once
)

func getMPSNNLabelsNodeClass() MPSNNLabelsNodeClass {
	_MPSNNLabelsNodeClassOnce.Do(func() {
		_MPSNNLabelsNodeClass = MPSNNLabelsNodeClass{class: objc.GetClass("MPSNNLabelsNode")}
	})
	return _MPSNNLabelsNodeClass
}

// GetMPSNNLabelsNodeClass returns the class object for MPSNNLabelsNode.
func GetMPSNNLabelsNodeClass() MPSNNLabelsNodeClass {
	return getMPSNNLabelsNodeClass()
}

type MPSNNLabelsNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNLabelsNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLabelsNodeClass) Alloc() MPSNNLabelsNode {
	rv := objc.Send[MPSNNLabelsNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A placeholder node denoting the per-element weight buffer used by loss and
// gradient loss kernels.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLabelsNode
type MPSNNLabelsNode struct {
	MPSNNStateNode
}

// MPSNNLabelsNodeFromID constructs a [MPSNNLabelsNode] from an objc.ID.
//
// A placeholder node denoting the per-element weight buffer used by loss and
// gradient loss kernels.
func MPSNNLabelsNodeFromID(id objc.ID) MPSNNLabelsNode {
	return MPSNNLabelsNode{MPSNNStateNode: MPSNNStateNodeFromID(id)}
}

// NOTE: MPSNNLabelsNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNLabelsNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLabelsNode
type IMPSNNLabelsNode interface {
	IMPSNNStateNode
}

// Init initializes the instance.
func (l MPSNNLabelsNode) Init() MPSNNLabelsNode {
	rv := objc.Send[MPSNNLabelsNode](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLabelsNode) Autorelease() MPSNNLabelsNode {
	rv := objc.Send[MPSNNLabelsNode](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLabelsNode creates a new MPSNNLabelsNode instance.
func NewMPSNNLabelsNode() MPSNNLabelsNode {
	class := getMPSNNLabelsNodeClass()
	rv := objc.Send[MPSNNLabelsNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

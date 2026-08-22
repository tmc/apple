// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNMultiaryGradientStateNode] class.
var (
	_MPSNNMultiaryGradientStateNodeClass     MPSNNMultiaryGradientStateNodeClass
	_MPSNNMultiaryGradientStateNodeClassOnce sync.Once
)

func getMPSNNMultiaryGradientStateNodeClass() MPSNNMultiaryGradientStateNodeClass {
	_MPSNNMultiaryGradientStateNodeClassOnce.Do(func() {
		_MPSNNMultiaryGradientStateNodeClass = MPSNNMultiaryGradientStateNodeClass{class: objc.GetClass("MPSNNMultiaryGradientStateNode")}
	})
	return _MPSNNMultiaryGradientStateNodeClass
}

// GetMPSNNMultiaryGradientStateNodeClass returns the class object for MPSNNMultiaryGradientStateNode.
func GetMPSNNMultiaryGradientStateNodeClass() MPSNNMultiaryGradientStateNodeClass {
	return getMPSNNMultiaryGradientStateNodeClass()
}

type MPSNNMultiaryGradientStateNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNMultiaryGradientStateNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNMultiaryGradientStateNodeClass) Alloc() MPSNNMultiaryGradientStateNode {
	rv := objc.Send[MPSNNMultiaryGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiaryGradientStateNode
type MPSNNMultiaryGradientStateNode struct {
	MPSNNStateNode
}

// MPSNNMultiaryGradientStateNodeFromID constructs a [MPSNNMultiaryGradientStateNode] from an objc.ID.
func MPSNNMultiaryGradientStateNodeFromID(id objc.ID) MPSNNMultiaryGradientStateNode {
	return MPSNNMultiaryGradientStateNode{MPSNNStateNode: MPSNNStateNodeFromID(id)}
}

// NOTE: MPSNNMultiaryGradientStateNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNMultiaryGradientStateNode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiaryGradientStateNode
type IMPSNNMultiaryGradientStateNode interface {
	IMPSNNStateNode
}

// Init initializes the instance.
func (m MPSNNMultiaryGradientStateNode) Init() MPSNNMultiaryGradientStateNode {
	rv := objc.Send[MPSNNMultiaryGradientStateNode](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSNNMultiaryGradientStateNode) Autorelease() MPSNNMultiaryGradientStateNode {
	rv := objc.Send[MPSNNMultiaryGradientStateNode](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNMultiaryGradientStateNode creates a new MPSNNMultiaryGradientStateNode instance.
func NewMPSNNMultiaryGradientStateNode() MPSNNMultiaryGradientStateNode {
	class := getMPSNNMultiaryGradientStateNodeClass()
	rv := objc.Send[MPSNNMultiaryGradientStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

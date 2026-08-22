// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReshapeNode] class.
var (
	_MPSNNReshapeNodeClass     MPSNNReshapeNodeClass
	_MPSNNReshapeNodeClassOnce sync.Once
)

func getMPSNNReshapeNodeClass() MPSNNReshapeNodeClass {
	_MPSNNReshapeNodeClassOnce.Do(func() {
		_MPSNNReshapeNodeClass = MPSNNReshapeNodeClass{class: objc.GetClass("MPSNNReshapeNode")}
	})
	return _MPSNNReshapeNodeClass
}

// GetMPSNNReshapeNodeClass returns the class object for MPSNNReshapeNode.
func GetMPSNNReshapeNodeClass() MPSNNReshapeNodeClass {
	return getMPSNNReshapeNodeClass()
}

type MPSNNReshapeNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReshapeNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReshapeNodeClass) Alloc() MPSNNReshapeNode {
	rv := objc.Send[MPSNNReshapeNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNReshapeNode.InitWithSourceResultWidthResultHeightResultFeatureChannels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeNode
type MPSNNReshapeNode struct {
	MPSNNFilterNode
}

// MPSNNReshapeNodeFromID constructs a [MPSNNReshapeNode] from an objc.ID.
func MPSNNReshapeNodeFromID(id objc.ID) MPSNNReshapeNode {
	return MPSNNReshapeNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSNNReshapeNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReshapeNode] class.
//
// # Initializers
//
//   - [IMPSNNReshapeNode.InitWithSourceResultWidthResultHeightResultFeatureChannels]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeNode
type IMPSNNReshapeNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSourceResultWidthResultHeightResultFeatureChannels(source IMPSNNImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) MPSNNReshapeNode
}

// Init initializes the instance.
func (r MPSNNReshapeNode) Init() MPSNNReshapeNode {
	rv := objc.Send[MPSNNReshapeNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReshapeNode) Autorelease() MPSNNReshapeNode {
	rv := objc.Send[MPSNNReshapeNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReshapeNode creates a new MPSNNReshapeNode instance.
func NewMPSNNReshapeNode() MPSNNReshapeNode {
	class := getMPSNNReshapeNodeClass()
	rv := objc.Send[MPSNNReshapeNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeNode/init(source:resultWidth:resultHeight:resultFeatureChannels:)
func NewReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels(source IMPSNNImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) MPSNNReshapeNode {
	instance := getMPSNNReshapeNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, resultWidth, resultHeight, resultFeatureChannels)
	return MPSNNReshapeNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeNode/init(source:resultWidth:resultHeight:resultFeatureChannels:)
func (r MPSNNReshapeNode) InitWithSourceResultWidthResultHeightResultFeatureChannels(source IMPSNNImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) MPSNNReshapeNode {
	rv := objc.Send[MPSNNReshapeNode](r.ID, objc.Sel("initWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, resultWidth, resultHeight, resultFeatureChannels)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeNode/nodeWithSource:resultWidth:resultHeight:resultFeatureChannels:
func (_MPSNNReshapeNodeClass MPSNNReshapeNodeClass) NodeWithSourceResultWidthResultHeightResultFeatureChannels(source IMPSNNImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) MPSNNReshapeNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNReshapeNodeClass.class), objc.Sel("nodeWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, resultWidth, resultHeight, resultFeatureChannels)
	return MPSNNReshapeNodeFromID(rv)
}

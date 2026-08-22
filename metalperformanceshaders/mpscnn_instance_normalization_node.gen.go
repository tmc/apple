// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNInstanceNormalizationNode] class.
var (
	_MPSCNNInstanceNormalizationNodeClass     MPSCNNInstanceNormalizationNodeClass
	_MPSCNNInstanceNormalizationNodeClassOnce sync.Once
)

func getMPSCNNInstanceNormalizationNodeClass() MPSCNNInstanceNormalizationNodeClass {
	_MPSCNNInstanceNormalizationNodeClassOnce.Do(func() {
		_MPSCNNInstanceNormalizationNodeClass = MPSCNNInstanceNormalizationNodeClass{class: objc.GetClass("MPSCNNInstanceNormalizationNode")}
	})
	return _MPSCNNInstanceNormalizationNodeClass
}

// GetMPSCNNInstanceNormalizationNodeClass returns the class object for MPSCNNInstanceNormalizationNode.
func GetMPSCNNInstanceNormalizationNodeClass() MPSCNNInstanceNormalizationNodeClass {
	return getMPSCNNInstanceNormalizationNodeClass()
}

type MPSCNNInstanceNormalizationNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNInstanceNormalizationNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNInstanceNormalizationNodeClass) Alloc() MPSCNNInstanceNormalizationNode {
	rv := objc.Send[MPSCNNInstanceNormalizationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of an instance normalization kernel.
//
// # Initializers
//
//   - [MPSCNNInstanceNormalizationNode.InitWithSourceDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationNode
type MPSCNNInstanceNormalizationNode struct {
	MPSNNFilterNode
}

// MPSCNNInstanceNormalizationNodeFromID constructs a [MPSCNNInstanceNormalizationNode] from an objc.ID.
//
// A representation of an instance normalization kernel.
func MPSCNNInstanceNormalizationNodeFromID(id objc.ID) MPSCNNInstanceNormalizationNode {
	return MPSCNNInstanceNormalizationNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNInstanceNormalizationNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNInstanceNormalizationNode] class.
//
// # Initializers
//
//   - [IMPSCNNInstanceNormalizationNode.InitWithSourceDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationNode
type IMPSCNNInstanceNormalizationNode interface {
	IMPSNNFilterNode
	MPSNNTrainableNode

	// Topic: Initializers

	InitWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNInstanceNormalizationDataSource) MPSCNNInstanceNormalizationNode
}

// Init initializes the instance.
func (c MPSCNNInstanceNormalizationNode) Init() MPSCNNInstanceNormalizationNode {
	rv := objc.Send[MPSCNNInstanceNormalizationNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNInstanceNormalizationNode) Autorelease() MPSCNNInstanceNormalizationNode {
	rv := objc.Send[MPSCNNInstanceNormalizationNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNInstanceNormalizationNode creates a new MPSCNNInstanceNormalizationNode instance.
func NewMPSCNNInstanceNormalizationNode() MPSCNNInstanceNormalizationNode {
	class := getMPSCNNInstanceNormalizationNodeClass()
	rv := objc.Send[MPSCNNInstanceNormalizationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationNode/init(source:dataSource:)
func NewCNNInstanceNormalizationNodeWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNInstanceNormalizationDataSource) MPSCNNInstanceNormalizationNode {
	instance := getMPSCNNInstanceNormalizationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	return MPSCNNInstanceNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationNode/init(source:dataSource:)
func (c MPSCNNInstanceNormalizationNode) InitWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNInstanceNormalizationDataSource) MPSCNNInstanceNormalizationNode {
	rv := objc.Send[MPSCNNInstanceNormalizationNode](c.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationNode/nodeWithSource:dataSource:
func (_MPSCNNInstanceNormalizationNodeClass MPSCNNInstanceNormalizationNodeClass) NodeWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNInstanceNormalizationDataSource) MPSCNNInstanceNormalizationNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNInstanceNormalizationNodeClass.class), objc.Sel("nodeWithSource:dataSource:"), source, dataSource)
	return MPSCNNInstanceNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationNode/trainingStyle
func (c MPSCNNInstanceNormalizationNode) TrainingStyle() MPSNNTrainingStyle {
	rv := objc.Send[MPSNNTrainingStyle](c.ID, objc.Sel("trainingStyle"))
	return MPSNNTrainingStyle(rv)
}
func (c MPSCNNInstanceNormalizationNode) SetTrainingStyle(value MPSNNTrainingStyle) {
	objc.Send[struct{}](c.ID, objc.Sel("setTrainingStyle:"), value)
}

// Protocol methods for MPSNNTrainableNode

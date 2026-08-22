// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNGroupNormalizationNode] class.
var (
	_MPSCNNGroupNormalizationNodeClass     MPSCNNGroupNormalizationNodeClass
	_MPSCNNGroupNormalizationNodeClassOnce sync.Once
)

func getMPSCNNGroupNormalizationNodeClass() MPSCNNGroupNormalizationNodeClass {
	_MPSCNNGroupNormalizationNodeClassOnce.Do(func() {
		_MPSCNNGroupNormalizationNodeClass = MPSCNNGroupNormalizationNodeClass{class: objc.GetClass("MPSCNNGroupNormalizationNode")}
	})
	return _MPSCNNGroupNormalizationNodeClass
}

// GetMPSCNNGroupNormalizationNodeClass returns the class object for MPSCNNGroupNormalizationNode.
func GetMPSCNNGroupNormalizationNodeClass() MPSCNNGroupNormalizationNodeClass {
	return getMPSCNNGroupNormalizationNodeClass()
}

type MPSCNNGroupNormalizationNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNGroupNormalizationNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNGroupNormalizationNodeClass) Alloc() MPSCNNGroupNormalizationNode {
	rv := objc.Send[MPSCNNGroupNormalizationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSCNNGroupNormalizationNode.InitWithSourceDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode
type MPSCNNGroupNormalizationNode struct {
	MPSNNFilterNode
}

// MPSCNNGroupNormalizationNodeFromID constructs a [MPSCNNGroupNormalizationNode] from an objc.ID.
func MPSCNNGroupNormalizationNodeFromID(id objc.ID) MPSCNNGroupNormalizationNode {
	return MPSCNNGroupNormalizationNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNGroupNormalizationNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNGroupNormalizationNode] class.
//
// # Initializers
//
//   - [IMPSCNNGroupNormalizationNode.InitWithSourceDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode
type IMPSCNNGroupNormalizationNode interface {
	IMPSNNFilterNode
	MPSNNTrainableNode

	// Topic: Initializers

	InitWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNGroupNormalizationDataSource) MPSCNNGroupNormalizationNode
}

// Init initializes the instance.
func (c MPSCNNGroupNormalizationNode) Init() MPSCNNGroupNormalizationNode {
	rv := objc.Send[MPSCNNGroupNormalizationNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNGroupNormalizationNode) Autorelease() MPSCNNGroupNormalizationNode {
	rv := objc.Send[MPSCNNGroupNormalizationNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNGroupNormalizationNode creates a new MPSCNNGroupNormalizationNode instance.
func NewMPSCNNGroupNormalizationNode() MPSCNNGroupNormalizationNode {
	class := getMPSCNNGroupNormalizationNodeClass()
	rv := objc.Send[MPSCNNGroupNormalizationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode/init(source:dataSource:)
func NewCNNGroupNormalizationNodeWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNGroupNormalizationDataSource) MPSCNNGroupNormalizationNode {
	instance := getMPSCNNGroupNormalizationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	return MPSCNNGroupNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode/init(source:dataSource:)
func (c MPSCNNGroupNormalizationNode) InitWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNGroupNormalizationDataSource) MPSCNNGroupNormalizationNode {
	rv := objc.Send[MPSCNNGroupNormalizationNode](c.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode/nodeWithSource:dataSource:
func (_MPSCNNGroupNormalizationNodeClass MPSCNNGroupNormalizationNodeClass) NodeWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNGroupNormalizationDataSource) MPSCNNGroupNormalizationNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNGroupNormalizationNodeClass.class), objc.Sel("nodeWithSource:dataSource:"), source, dataSource)
	return MPSCNNGroupNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationNode/trainingStyle
func (c MPSCNNGroupNormalizationNode) TrainingStyle() MPSNNTrainingStyle {
	rv := objc.Send[MPSNNTrainingStyle](c.ID, objc.Sel("trainingStyle"))
	return MPSNNTrainingStyle(rv)
}
func (c MPSCNNGroupNormalizationNode) SetTrainingStyle(value MPSNNTrainingStyle) {
	objc.Send[struct{}](c.ID, objc.Sel("setTrainingStyle:"), value)
}

// Protocol methods for MPSNNTrainableNode

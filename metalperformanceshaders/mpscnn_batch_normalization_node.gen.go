// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNBatchNormalizationNode] class.
var (
	_MPSCNNBatchNormalizationNodeClass     MPSCNNBatchNormalizationNodeClass
	_MPSCNNBatchNormalizationNodeClassOnce sync.Once
)

func getMPSCNNBatchNormalizationNodeClass() MPSCNNBatchNormalizationNodeClass {
	_MPSCNNBatchNormalizationNodeClassOnce.Do(func() {
		_MPSCNNBatchNormalizationNodeClass = MPSCNNBatchNormalizationNodeClass{class: objc.GetClass("MPSCNNBatchNormalizationNode")}
	})
	return _MPSCNNBatchNormalizationNodeClass
}

// GetMPSCNNBatchNormalizationNodeClass returns the class object for MPSCNNBatchNormalizationNode.
func GetMPSCNNBatchNormalizationNodeClass() MPSCNNBatchNormalizationNodeClass {
	return getMPSCNNBatchNormalizationNodeClass()
}

type MPSCNNBatchNormalizationNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBatchNormalizationNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBatchNormalizationNodeClass) Alloc() MPSCNNBatchNormalizationNode {
	rv := objc.Send[MPSCNNBatchNormalizationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a batch normalization kernel.
//
// # Initializers
//
//   - [MPSCNNBatchNormalizationNode.InitWithSourceDataSource]
//
// # Instance Properties
//
//   - [MPSCNNBatchNormalizationNode.Flags]
//   - [MPSCNNBatchNormalizationNode.SetFlags]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationNode
type MPSCNNBatchNormalizationNode struct {
	MPSNNFilterNode
}

// MPSCNNBatchNormalizationNodeFromID constructs a [MPSCNNBatchNormalizationNode] from an objc.ID.
//
// A representation of a batch normalization kernel.
func MPSCNNBatchNormalizationNodeFromID(id objc.ID) MPSCNNBatchNormalizationNode {
	return MPSCNNBatchNormalizationNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSCNNBatchNormalizationNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBatchNormalizationNode] class.
//
// # Initializers
//
//   - [IMPSCNNBatchNormalizationNode.InitWithSourceDataSource]
//
// # Instance Properties
//
//   - [IMPSCNNBatchNormalizationNode.Flags]
//   - [IMPSCNNBatchNormalizationNode.SetFlags]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationNode
type IMPSCNNBatchNormalizationNode interface {
	IMPSNNFilterNode
	MPSNNTrainableNode

	// Topic: Initializers

	InitWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNBatchNormalizationDataSource) MPSCNNBatchNormalizationNode

	// Topic: Instance Properties

	Flags() MPSCNNBatchNormalizationFlags
	SetFlags(value MPSCNNBatchNormalizationFlags)
}

// Init initializes the instance.
func (c MPSCNNBatchNormalizationNode) Init() MPSCNNBatchNormalizationNode {
	rv := objc.Send[MPSCNNBatchNormalizationNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBatchNormalizationNode) Autorelease() MPSCNNBatchNormalizationNode {
	rv := objc.Send[MPSCNNBatchNormalizationNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBatchNormalizationNode creates a new MPSCNNBatchNormalizationNode instance.
func NewMPSCNNBatchNormalizationNode() MPSCNNBatchNormalizationNode {
	class := getMPSCNNBatchNormalizationNodeClass()
	rv := objc.Send[MPSCNNBatchNormalizationNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationNode/init(source:dataSource:)
func NewCNNBatchNormalizationNodeWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNBatchNormalizationDataSource) MPSCNNBatchNormalizationNode {
	instance := getMPSCNNBatchNormalizationNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	return MPSCNNBatchNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationNode/init(source:dataSource:)
func (c MPSCNNBatchNormalizationNode) InitWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNBatchNormalizationDataSource) MPSCNNBatchNormalizationNode {
	rv := objc.Send[MPSCNNBatchNormalizationNode](c.ID, objc.Sel("initWithSource:dataSource:"), source, dataSource)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationNode/nodeWithSource:dataSource:
func (_MPSCNNBatchNormalizationNodeClass MPSCNNBatchNormalizationNodeClass) NodeWithSourceDataSource(source IMPSNNImageNode, dataSource MPSCNNBatchNormalizationDataSource) MPSCNNBatchNormalizationNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNBatchNormalizationNodeClass.class), objc.Sel("nodeWithSource:dataSource:"), source, dataSource)
	return MPSCNNBatchNormalizationNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationNode/flags
func (c MPSCNNBatchNormalizationNode) Flags() MPSCNNBatchNormalizationFlags {
	rv := objc.Send[MPSCNNBatchNormalizationFlags](c.ID, objc.Sel("flags"))
	return MPSCNNBatchNormalizationFlags(rv)
}
func (c MPSCNNBatchNormalizationNode) SetFlags(value MPSCNNBatchNormalizationFlags) {
	objc.Send[struct{}](c.ID, objc.Sel("setFlags:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationNode/trainingStyle
func (c MPSCNNBatchNormalizationNode) TrainingStyle() MPSNNTrainingStyle {
	rv := objc.Send[MPSNNTrainingStyle](c.ID, objc.Sel("trainingStyle"))
	return MPSNNTrainingStyle(rv)
}
func (c MPSCNNBatchNormalizationNode) SetTrainingStyle(value MPSNNTrainingStyle) {
	objc.Send[struct{}](c.ID, objc.Sel("setTrainingStyle:"), value)
}

// Protocol methods for MPSNNTrainableNode

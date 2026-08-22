// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNForwardLossNode] class.
var (
	_MPSNNForwardLossNodeClass     MPSNNForwardLossNodeClass
	_MPSNNForwardLossNodeClassOnce sync.Once
)

func getMPSNNForwardLossNodeClass() MPSNNForwardLossNodeClass {
	_MPSNNForwardLossNodeClassOnce.Do(func() {
		_MPSNNForwardLossNodeClass = MPSNNForwardLossNodeClass{class: objc.GetClass("MPSNNForwardLossNode")}
	})
	return _MPSNNForwardLossNodeClass
}

// GetMPSNNForwardLossNodeClass returns the class object for MPSNNForwardLossNode.
func GetMPSNNForwardLossNodeClass() MPSNNForwardLossNodeClass {
	return getMPSNNForwardLossNodeClass()
}

type MPSNNForwardLossNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNForwardLossNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNForwardLossNodeClass) Alloc() MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNForwardLossNode.InitWithSourceLabelsLossDescriptor]
//   - [MPSNNForwardLossNode.InitWithSourceLabelsWeightsLossDescriptor]
//   - [MPSNNForwardLossNode.InitWithSourcesLossDescriptor]
//
// # Instance Properties
//
//   - [MPSNNForwardLossNode.Delta]
//   - [MPSNNForwardLossNode.Epsilon]
//   - [MPSNNForwardLossNode.LabelSmoothing]
//   - [MPSNNForwardLossNode.LossType]
//   - [MPSNNForwardLossNode.NumberOfClasses]
//   - [MPSNNForwardLossNode.PropertyCallBack]
//   - [MPSNNForwardLossNode.SetPropertyCallBack]
//   - [MPSNNForwardLossNode.ReduceAcrossBatch]
//   - [MPSNNForwardLossNode.ReductionType]
//   - [MPSNNForwardLossNode.Weight]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode
type MPSNNForwardLossNode struct {
	MPSNNFilterNode
}

// MPSNNForwardLossNodeFromID constructs a [MPSNNForwardLossNode] from an objc.ID.
func MPSNNForwardLossNodeFromID(id objc.ID) MPSNNForwardLossNode {
	return MPSNNForwardLossNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSNNForwardLossNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNForwardLossNode] class.
//
// # Initializers
//
//   - [IMPSNNForwardLossNode.InitWithSourceLabelsLossDescriptor]
//   - [IMPSNNForwardLossNode.InitWithSourceLabelsWeightsLossDescriptor]
//   - [IMPSNNForwardLossNode.InitWithSourcesLossDescriptor]
//
// # Instance Properties
//
//   - [IMPSNNForwardLossNode.Delta]
//   - [IMPSNNForwardLossNode.Epsilon]
//   - [IMPSNNForwardLossNode.LabelSmoothing]
//   - [IMPSNNForwardLossNode.LossType]
//   - [IMPSNNForwardLossNode.NumberOfClasses]
//   - [IMPSNNForwardLossNode.PropertyCallBack]
//   - [IMPSNNForwardLossNode.SetPropertyCallBack]
//   - [IMPSNNForwardLossNode.ReduceAcrossBatch]
//   - [IMPSNNForwardLossNode.ReductionType]
//   - [IMPSNNForwardLossNode.Weight]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode
type IMPSNNForwardLossNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithSourceLabelsLossDescriptor(source IMPSNNImageNode, labels IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode
	InitWithSourceLabelsWeightsLossDescriptor(source IMPSNNImageNode, labels IMPSNNImageNode, weights IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode
	InitWithSourcesLossDescriptor(sourceNodes []MPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode

	// Topic: Instance Properties

	Delta() float32
	Epsilon() float32
	LabelSmoothing() float32
	LossType() MPSCNNLossType
	NumberOfClasses() uint
	PropertyCallBack() MPSNNLossCallback
	SetPropertyCallBack(value MPSNNLossCallback)
	ReduceAcrossBatch() bool
	ReductionType() MPSCNNReductionType
	Weight() float32
}

// Init initializes the instance.
func (f MPSNNForwardLossNode) Init() MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MPSNNForwardLossNode) Autorelease() MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNForwardLossNode creates a new MPSNNForwardLossNode instance.
func NewMPSNNForwardLossNode() MPSNNForwardLossNode {
	class := getMPSNNForwardLossNodeClass()
	rv := objc.Send[MPSNNForwardLossNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/init(source:labels:lossDescriptor:)
func NewForwardLossNodeWithSourceLabelsLossDescriptor(source IMPSNNImageNode, labels IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode {
	instance := getMPSNNForwardLossNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:labels:lossDescriptor:"), source, labels, descriptor)
	return MPSNNForwardLossNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/init(source:labels:weights:lossDescriptor:)-8c2l6
func NewForwardLossNodeWithSourceLabelsWeightsLossDescriptor(source IMPSNNImageNode, labels IMPSNNImageNode, weights IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode {
	instance := getMPSNNForwardLossNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:labels:weights:lossDescriptor:"), source, labels, weights, descriptor)
	return MPSNNForwardLossNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/init(sources:lossDescriptor:)
func NewForwardLossNodeWithSourcesLossDescriptor(sourceNodes []MPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode {
	instance := getMPSNNForwardLossNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:lossDescriptor:"), objectivec.IObjectSliceToNSArray(sourceNodes), descriptor)
	return MPSNNForwardLossNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/init(source:labels:lossDescriptor:)
func (f MPSNNForwardLossNode) InitWithSourceLabelsLossDescriptor(source IMPSNNImageNode, labels IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](f.ID, objc.Sel("initWithSource:labels:lossDescriptor:"), source, labels, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/init(source:labels:weights:lossDescriptor:)-8c2l6
func (f MPSNNForwardLossNode) InitWithSourceLabelsWeightsLossDescriptor(source IMPSNNImageNode, labels IMPSNNImageNode, weights IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](f.ID, objc.Sel("initWithSource:labels:weights:lossDescriptor:"), source, labels, weights, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/init(sources:lossDescriptor:)
func (f MPSNNForwardLossNode) InitWithSourcesLossDescriptor(sourceNodes []MPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode {
	rv := objc.Send[MPSNNForwardLossNode](f.ID, objc.Sel("initWithSources:lossDescriptor:"), objectivec.IObjectSliceToNSArray(sourceNodes), descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/init(source:labels:weights:lossDescriptor:)-9bsd7
func (_MPSNNForwardLossNodeClass MPSNNForwardLossNodeClass) NodeWithSourceLabelsWeightsLossDescriptor(source IMPSNNImageNode, labels IMPSNNImageNode, weights IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNForwardLossNodeClass.class), objc.Sel("nodeWithSource:labels:weights:lossDescriptor:"), source, labels, weights, descriptor)
	return MPSNNForwardLossNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/nodeWithSource:labels:lossDescriptor:
func (_MPSNNForwardLossNodeClass MPSNNForwardLossNodeClass) NodeWithSourceLabelsLossDescriptor(source IMPSNNImageNode, labels IMPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNForwardLossNodeClass.class), objc.Sel("nodeWithSource:labels:lossDescriptor:"), source, labels, descriptor)
	return MPSNNForwardLossNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/nodeWithSources:lossDescriptor:
func (_MPSNNForwardLossNodeClass MPSNNForwardLossNodeClass) NodeWithSourcesLossDescriptor(sourceNodes []MPSNNImageNode, descriptor IMPSCNNLossDescriptor) MPSNNForwardLossNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNForwardLossNodeClass.class), objc.Sel("nodeWithSources:lossDescriptor:"), objectivec.IObjectSliceToNSArray(sourceNodes), descriptor)
	return MPSNNForwardLossNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/delta
func (f MPSNNForwardLossNode) Delta() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("delta"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/epsilon
func (f MPSNNForwardLossNode) Epsilon() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("epsilon"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/labelSmoothing
func (f MPSNNForwardLossNode) LabelSmoothing() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("labelSmoothing"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/lossType
func (f MPSNNForwardLossNode) LossType() MPSCNNLossType {
	rv := objc.Send[MPSCNNLossType](f.ID, objc.Sel("lossType"))
	return MPSCNNLossType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/numberOfClasses
func (f MPSNNForwardLossNode) NumberOfClasses() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("numberOfClasses"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/propertyCallBack
func (f MPSNNForwardLossNode) PropertyCallBack() MPSNNLossCallback {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("propertyCallBack"))
	return MPSNNLossCallbackObjectFromID(rv)
}
func (f MPSNNForwardLossNode) SetPropertyCallBack(value MPSNNLossCallback) {
	objc.Send[struct{}](f.ID, objc.Sel("setPropertyCallBack:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/reduceAcrossBatch
func (f MPSNNForwardLossNode) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/reductionType
func (f MPSNNForwardLossNode) ReductionType() MPSCNNReductionType {
	rv := objc.Send[MPSCNNReductionType](f.ID, objc.Sel("reductionType"))
	return MPSCNNReductionType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/weight
func (f MPSNNForwardLossNode) Weight() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("weight"))
	return rv
}

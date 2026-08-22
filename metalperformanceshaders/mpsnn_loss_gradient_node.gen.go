// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNLossGradientNode] class.
var (
	_MPSNNLossGradientNodeClass     MPSNNLossGradientNodeClass
	_MPSNNLossGradientNodeClassOnce sync.Once
)

func getMPSNNLossGradientNodeClass() MPSNNLossGradientNodeClass {
	_MPSNNLossGradientNodeClassOnce.Do(func() {
		_MPSNNLossGradientNodeClass = MPSNNLossGradientNodeClass{class: objc.GetClass("MPSNNLossGradientNode")}
	})
	return _MPSNNLossGradientNodeClass
}

// GetMPSNNLossGradientNodeClass returns the class object for MPSNNLossGradientNode.
func GetMPSNNLossGradientNodeClass() MPSNNLossGradientNodeClass {
	return getMPSNNLossGradientNodeClass()
}

type MPSNNLossGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNLossGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLossGradientNodeClass) Alloc() MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNLossGradientNode.InitWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter]
//   - [MPSNNLossGradientNode.InitWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter]
//   - [MPSNNLossGradientNode.InitWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter]
//
// # Instance Properties
//
//   - [MPSNNLossGradientNode.Delta]
//   - [MPSNNLossGradientNode.Epsilon]
//   - [MPSNNLossGradientNode.IsLabelsGradientFilter]
//   - [MPSNNLossGradientNode.LabelSmoothing]
//   - [MPSNNLossGradientNode.LossType]
//   - [MPSNNLossGradientNode.NumberOfClasses]
//   - [MPSNNLossGradientNode.PropertyCallBack]
//   - [MPSNNLossGradientNode.SetPropertyCallBack]
//   - [MPSNNLossGradientNode.ReduceAcrossBatch]
//   - [MPSNNLossGradientNode.ReductionType]
//   - [MPSNNLossGradientNode.Weight]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode
type MPSNNLossGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNLossGradientNodeFromID constructs a [MPSNNLossGradientNode] from an objc.ID.
func MPSNNLossGradientNodeFromID(id objc.ID) MPSNNLossGradientNode {
	return MPSNNLossGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSNNLossGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNLossGradientNode] class.
//
// # Initializers
//
//   - [IMPSNNLossGradientNode.InitWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter]
//   - [IMPSNNLossGradientNode.InitWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter]
//   - [IMPSNNLossGradientNode.InitWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter]
//
// # Instance Properties
//
//   - [IMPSNNLossGradientNode.Delta]
//   - [IMPSNNLossGradientNode.Epsilon]
//   - [IMPSNNLossGradientNode.IsLabelsGradientFilter]
//   - [IMPSNNLossGradientNode.LabelSmoothing]
//   - [IMPSNNLossGradientNode.LossType]
//   - [IMPSNNLossGradientNode.NumberOfClasses]
//   - [IMPSNNLossGradientNode.PropertyCallBack]
//   - [IMPSNNLossGradientNode.SetPropertyCallBack]
//   - [IMPSNNLossGradientNode.ReduceAcrossBatch]
//   - [IMPSNNLossGradientNode.ReductionType]
//   - [IMPSNNLossGradientNode.Weight]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode
type IMPSNNLossGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, labels IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode
	InitWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, labels IMPSNNImageNode, weights IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode
	InitWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes []MPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode

	// Topic: Instance Properties

	Delta() float32
	Epsilon() float32
	IsLabelsGradientFilter() bool
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
func (l MPSNNLossGradientNode) Init() MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLossGradientNode) Autorelease() MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLossGradientNode creates a new MPSNNLossGradientNode instance.
func NewMPSNNLossGradientNode() MPSNNLossGradientNode {
	class := getMPSNNLossGradientNodeClass()
	rv := objc.Send[MPSNNLossGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/init(sourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:)
func NewLossGradientNodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, labels IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode {
	instance := getMPSNNLossGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, gradientState, descriptor, isLabelsGradientFilter)
	return MPSNNLossGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/init(sourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:)-3rcen
func NewLossGradientNodeWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, labels IMPSNNImageNode, weights IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode {
	instance := getMPSNNLossGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, weights, gradientState, descriptor, isLabelsGradientFilter)
	return MPSNNLossGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/init(sources:gradientState:lossDescriptor:isLabelsGradientFilter:)
func NewLossGradientNodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes []MPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode {
	instance := getMPSNNLossGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), objectivec.IObjectSliceToNSArray(sourceNodes), gradientState, descriptor, isLabelsGradientFilter)
	return MPSNNLossGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/init(sourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:)
func (l MPSNNLossGradientNode) InitWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, labels IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](l.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/init(sourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:)-3rcen
func (l MPSNNLossGradientNode) InitWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, labels IMPSNNImageNode, weights IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](l.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, weights, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/init(sources:gradientState:lossDescriptor:isLabelsGradientFilter:)
func (l MPSNNLossGradientNode) InitWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes []MPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode {
	rv := objc.Send[MPSNNLossGradientNode](l.ID, objc.Sel("initWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), objectivec.IObjectSliceToNSArray(sourceNodes), gradientState, descriptor, isLabelsGradientFilter)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/init(sourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:)-9eqch
func (_MPSNNLossGradientNodeClass MPSNNLossGradientNodeClass) NodeWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, labels IMPSNNImageNode, weights IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNLossGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, weights, gradientState, descriptor, isLabelsGradientFilter)
	return MPSNNLossGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/nodeWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:
func (_MPSNNLossGradientNodeClass MPSNNLossGradientNodeClass) NodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, labels IMPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNLossGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, gradientState, descriptor, isLabelsGradientFilter)
	return MPSNNLossGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:
func (_MPSNNLossGradientNodeClass MPSNNLossGradientNodeClass) NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes []MPSNNImageNode, gradientState IMPSNNGradientStateNode, descriptor IMPSCNNLossDescriptor, isLabelsGradientFilter bool) MPSNNLossGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNLossGradientNodeClass.class), objc.Sel("nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), objectivec.IObjectSliceToNSArray(sourceNodes), gradientState, descriptor, isLabelsGradientFilter)
	return MPSNNLossGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/delta
func (l MPSNNLossGradientNode) Delta() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("delta"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/epsilon
func (l MPSNNLossGradientNode) Epsilon() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("epsilon"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/isLabelsGradientFilter
func (l MPSNNLossGradientNode) IsLabelsGradientFilter() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isLabelsGradientFilter"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/labelSmoothing
func (l MPSNNLossGradientNode) LabelSmoothing() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("labelSmoothing"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/lossType
func (l MPSNNLossGradientNode) LossType() MPSCNNLossType {
	rv := objc.Send[MPSCNNLossType](l.ID, objc.Sel("lossType"))
	return MPSCNNLossType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/numberOfClasses
func (l MPSNNLossGradientNode) NumberOfClasses() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("numberOfClasses"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/propertyCallBack
func (l MPSNNLossGradientNode) PropertyCallBack() MPSNNLossCallback {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("propertyCallBack"))
	return MPSNNLossCallbackObjectFromID(rv)
}
func (l MPSNNLossGradientNode) SetPropertyCallBack(value MPSNNLossCallback) {
	objc.Send[struct{}](l.ID, objc.Sel("setPropertyCallBack:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/reduceAcrossBatch
func (l MPSNNLossGradientNode) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/reductionType
func (l MPSNNLossGradientNode) ReductionType() MPSCNNReductionType {
	rv := objc.Send[MPSCNNReductionType](l.ID, objc.Sel("reductionType"))
	return MPSCNNReductionType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/weight
func (l MPSNNLossGradientNode) Weight() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("weight"))
	return rv
}

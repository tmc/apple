// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNArithmeticGradientNode] class.
var (
	_MPSNNArithmeticGradientNodeClass     MPSNNArithmeticGradientNodeClass
	_MPSNNArithmeticGradientNodeClassOnce sync.Once
)

func getMPSNNArithmeticGradientNodeClass() MPSNNArithmeticGradientNodeClass {
	_MPSNNArithmeticGradientNodeClassOnce.Do(func() {
		_MPSNNArithmeticGradientNodeClass = MPSNNArithmeticGradientNodeClass{class: objc.GetClass("MPSNNArithmeticGradientNode")}
	})
	return _MPSNNArithmeticGradientNodeClass
}

// GetMPSNNArithmeticGradientNodeClass returns the class object for MPSNNArithmeticGradientNode.
func GetMPSNNArithmeticGradientNodeClass() MPSNNArithmeticGradientNodeClass {
	return getMPSNNArithmeticGradientNodeClass()
}

type MPSNNArithmeticGradientNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNArithmeticGradientNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNArithmeticGradientNodeClass) Alloc() MPSNNArithmeticGradientNode {
	rv := objc.Send[MPSNNArithmeticGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A representation of the base class for gradient arithmetic operators.
//
// # Initializers
//
//   - [MPSNNArithmeticGradientNode.InitWithGradientImagesForwardFilterIsSecondarySourceFilter]
//   - [MPSNNArithmeticGradientNode.InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter]
//
// # Instance Properties
//
//   - [MPSNNArithmeticGradientNode.Bias]
//   - [MPSNNArithmeticGradientNode.SetBias]
//   - [MPSNNArithmeticGradientNode.IsSecondarySourceFilter]
//   - [MPSNNArithmeticGradientNode.MaximumValue]
//   - [MPSNNArithmeticGradientNode.SetMaximumValue]
//   - [MPSNNArithmeticGradientNode.MinimumValue]
//   - [MPSNNArithmeticGradientNode.SetMinimumValue]
//   - [MPSNNArithmeticGradientNode.PrimaryScale]
//   - [MPSNNArithmeticGradientNode.SetPrimaryScale]
//   - [MPSNNArithmeticGradientNode.SecondaryScale]
//   - [MPSNNArithmeticGradientNode.SetSecondaryScale]
//   - [MPSNNArithmeticGradientNode.SecondaryStrideInFeatureChannels]
//   - [MPSNNArithmeticGradientNode.SetSecondaryStrideInFeatureChannels]
//   - [MPSNNArithmeticGradientNode.SecondaryStrideInPixelsX]
//   - [MPSNNArithmeticGradientNode.SetSecondaryStrideInPixelsX]
//   - [MPSNNArithmeticGradientNode.SecondaryStrideInPixelsY]
//   - [MPSNNArithmeticGradientNode.SetSecondaryStrideInPixelsY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode
type MPSNNArithmeticGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNArithmeticGradientNodeFromID constructs a [MPSNNArithmeticGradientNode] from an objc.ID.
//
// A representation of the base class for gradient arithmetic operators.
func MPSNNArithmeticGradientNodeFromID(id objc.ID) MPSNNArithmeticGradientNode {
	return MPSNNArithmeticGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}

// NOTE: MPSNNArithmeticGradientNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNArithmeticGradientNode] class.
//
// # Initializers
//
//   - [IMPSNNArithmeticGradientNode.InitWithGradientImagesForwardFilterIsSecondarySourceFilter]
//   - [IMPSNNArithmeticGradientNode.InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter]
//
// # Instance Properties
//
//   - [IMPSNNArithmeticGradientNode.Bias]
//   - [IMPSNNArithmeticGradientNode.SetBias]
//   - [IMPSNNArithmeticGradientNode.IsSecondarySourceFilter]
//   - [IMPSNNArithmeticGradientNode.MaximumValue]
//   - [IMPSNNArithmeticGradientNode.SetMaximumValue]
//   - [IMPSNNArithmeticGradientNode.MinimumValue]
//   - [IMPSNNArithmeticGradientNode.SetMinimumValue]
//   - [IMPSNNArithmeticGradientNode.PrimaryScale]
//   - [IMPSNNArithmeticGradientNode.SetPrimaryScale]
//   - [IMPSNNArithmeticGradientNode.SecondaryScale]
//   - [IMPSNNArithmeticGradientNode.SetSecondaryScale]
//   - [IMPSNNArithmeticGradientNode.SecondaryStrideInFeatureChannels]
//   - [IMPSNNArithmeticGradientNode.SetSecondaryStrideInFeatureChannels]
//   - [IMPSNNArithmeticGradientNode.SecondaryStrideInPixelsX]
//   - [IMPSNNArithmeticGradientNode.SetSecondaryStrideInPixelsX]
//   - [IMPSNNArithmeticGradientNode.SecondaryStrideInPixelsY]
//   - [IMPSNNArithmeticGradientNode.SetSecondaryStrideInPixelsY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode
type IMPSNNArithmeticGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Initializers

	InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []MPSNNImageNode, filter IMPSNNFilterNode, isSecondarySourceFilter bool) MPSNNArithmeticGradientNode
	InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNBinaryGradientStateNode, isSecondarySourceFilter bool) MPSNNArithmeticGradientNode

	// Topic: Instance Properties

	Bias() float32
	SetBias(value float32)
	IsSecondarySourceFilter() bool
	MaximumValue() float32
	SetMaximumValue(value float32)
	MinimumValue() float32
	SetMinimumValue(value float32)
	PrimaryScale() float32
	SetPrimaryScale(value float32)
	SecondaryScale() float32
	SetSecondaryScale(value float32)
	SecondaryStrideInFeatureChannels() uint
	SetSecondaryStrideInFeatureChannels(value uint)
	SecondaryStrideInPixelsX() uint
	SetSecondaryStrideInPixelsX(value uint)
	SecondaryStrideInPixelsY() uint
	SetSecondaryStrideInPixelsY(value uint)
}

// Init initializes the instance.
func (a MPSNNArithmeticGradientNode) Init() MPSNNArithmeticGradientNode {
	rv := objc.Send[MPSNNArithmeticGradientNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MPSNNArithmeticGradientNode) Autorelease() MPSNNArithmeticGradientNode {
	rv := objc.Send[MPSNNArithmeticGradientNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNArithmeticGradientNode creates a new MPSNNArithmeticGradientNode instance.
func NewMPSNNArithmeticGradientNode() MPSNNArithmeticGradientNode {
	class := getMPSNNArithmeticGradientNodeClass()
	rv := objc.Send[MPSNNArithmeticGradientNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/init(gradientImages:forwardFilter:isSecondarySourceFilter:)
func NewArithmeticGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []MPSNNImageNode, filter IMPSNNFilterNode, isSecondarySourceFilter bool) MPSNNArithmeticGradientNode {
	instance := getMPSNNArithmeticGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), objectivec.IObjectSliceToNSArray(gradientImages), filter, isSecondarySourceFilter)
	return MPSNNArithmeticGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/init(sourceGradient:sourceImage:gradientState:isSecondarySourceFilter:)
func NewArithmeticGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNBinaryGradientStateNode, isSecondarySourceFilter bool) MPSNNArithmeticGradientNode {
	instance := getMPSNNArithmeticGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	return MPSNNArithmeticGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/init(gradientImages:forwardFilter:isSecondarySourceFilter:)
func (a MPSNNArithmeticGradientNode) InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []MPSNNImageNode, filter IMPSNNFilterNode, isSecondarySourceFilter bool) MPSNNArithmeticGradientNode {
	rv := objc.Send[MPSNNArithmeticGradientNode](a.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), objectivec.IObjectSliceToNSArray(gradientImages), filter, isSecondarySourceFilter)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/init(sourceGradient:sourceImage:gradientState:isSecondarySourceFilter:)
func (a MPSNNArithmeticGradientNode) InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNBinaryGradientStateNode, isSecondarySourceFilter bool) MPSNNArithmeticGradientNode {
	rv := objc.Send[MPSNNArithmeticGradientNode](a.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/nodeWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:
func (_MPSNNArithmeticGradientNodeClass MPSNNArithmeticGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient IMPSNNImageNode, sourceImage IMPSNNImageNode, gradientState IMPSNNBinaryGradientStateNode, isSecondarySourceFilter bool) MPSNNArithmeticGradientNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNArithmeticGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	return MPSNNArithmeticGradientNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/bias
func (a MPSNNArithmeticGradientNode) Bias() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("bias"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetBias(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setBias:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/isSecondarySourceFilter
func (a MPSNNArithmeticGradientNode) IsSecondarySourceFilter() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isSecondarySourceFilter"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/maximumValue
func (a MPSNNArithmeticGradientNode) MaximumValue() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("maximumValue"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetMaximumValue(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setMaximumValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/minimumValue
func (a MPSNNArithmeticGradientNode) MinimumValue() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("minimumValue"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetMinimumValue(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setMinimumValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/primaryScale
func (a MPSNNArithmeticGradientNode) PrimaryScale() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("primaryScale"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetPrimaryScale(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPrimaryScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/secondaryScale
func (a MPSNNArithmeticGradientNode) SecondaryScale() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("secondaryScale"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetSecondaryScale(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setSecondaryScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/secondaryStrideInFeatureChannels
func (a MPSNNArithmeticGradientNode) SecondaryStrideInFeatureChannels() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetSecondaryStrideInFeatureChannels(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/secondaryStrideInPixelsX
func (a MPSNNArithmeticGradientNode) SecondaryStrideInPixelsX() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetSecondaryStrideInPixelsX(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientNode/secondaryStrideInPixelsY
func (a MPSNNArithmeticGradientNode) SecondaryStrideInPixelsY() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetSecondaryStrideInPixelsY(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}

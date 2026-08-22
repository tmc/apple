// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNBinaryArithmeticNode] class.
var (
	_MPSNNBinaryArithmeticNodeClass     MPSNNBinaryArithmeticNodeClass
	_MPSNNBinaryArithmeticNodeClassOnce sync.Once
)

func getMPSNNBinaryArithmeticNodeClass() MPSNNBinaryArithmeticNodeClass {
	_MPSNNBinaryArithmeticNodeClassOnce.Do(func() {
		_MPSNNBinaryArithmeticNodeClass = MPSNNBinaryArithmeticNodeClass{class: objc.GetClass("MPSNNBinaryArithmeticNode")}
	})
	return _MPSNNBinaryArithmeticNodeClass
}

// GetMPSNNBinaryArithmeticNodeClass returns the class object for MPSNNBinaryArithmeticNode.
func GetMPSNNBinaryArithmeticNodeClass() MPSNNBinaryArithmeticNodeClass {
	return getMPSNNBinaryArithmeticNodeClass()
}

type MPSNNBinaryArithmeticNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNBinaryArithmeticNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNBinaryArithmeticNodeClass) Alloc() MPSNNBinaryArithmeticNode {
	rv := objc.Send[MPSNNBinaryArithmeticNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Virtual base class for basic arithmetic nodes.
//
// # Initializers
//
//   - [MPSNNBinaryArithmeticNode.InitWithLeftSourceRightSource]
//   - [MPSNNBinaryArithmeticNode.InitWithSources]
//
// # Instance Properties
//
//   - [MPSNNBinaryArithmeticNode.Bias]
//   - [MPSNNBinaryArithmeticNode.SetBias]
//   - [MPSNNBinaryArithmeticNode.MaximumValue]
//   - [MPSNNBinaryArithmeticNode.SetMaximumValue]
//   - [MPSNNBinaryArithmeticNode.MinimumValue]
//   - [MPSNNBinaryArithmeticNode.SetMinimumValue]
//   - [MPSNNBinaryArithmeticNode.PrimaryScale]
//   - [MPSNNBinaryArithmeticNode.SetPrimaryScale]
//   - [MPSNNBinaryArithmeticNode.PrimaryStrideInFeatureChannels]
//   - [MPSNNBinaryArithmeticNode.SetPrimaryStrideInFeatureChannels]
//   - [MPSNNBinaryArithmeticNode.PrimaryStrideInPixelsX]
//   - [MPSNNBinaryArithmeticNode.SetPrimaryStrideInPixelsX]
//   - [MPSNNBinaryArithmeticNode.PrimaryStrideInPixelsY]
//   - [MPSNNBinaryArithmeticNode.SetPrimaryStrideInPixelsY]
//   - [MPSNNBinaryArithmeticNode.SecondaryScale]
//   - [MPSNNBinaryArithmeticNode.SetSecondaryScale]
//   - [MPSNNBinaryArithmeticNode.SecondaryStrideInFeatureChannels]
//   - [MPSNNBinaryArithmeticNode.SetSecondaryStrideInFeatureChannels]
//   - [MPSNNBinaryArithmeticNode.SecondaryStrideInPixelsX]
//   - [MPSNNBinaryArithmeticNode.SetSecondaryStrideInPixelsX]
//   - [MPSNNBinaryArithmeticNode.SecondaryStrideInPixelsY]
//   - [MPSNNBinaryArithmeticNode.SetSecondaryStrideInPixelsY]
//
// # Instance Methods
//
//   - [MPSNNBinaryArithmeticNode.GradientClass]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode
type MPSNNBinaryArithmeticNode struct {
	MPSNNFilterNode
}

// MPSNNBinaryArithmeticNodeFromID constructs a [MPSNNBinaryArithmeticNode] from an objc.ID.
//
// Virtual base class for basic arithmetic nodes.
func MPSNNBinaryArithmeticNodeFromID(id objc.ID) MPSNNBinaryArithmeticNode {
	return MPSNNBinaryArithmeticNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}

// NOTE: MPSNNBinaryArithmeticNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNBinaryArithmeticNode] class.
//
// # Initializers
//
//   - [IMPSNNBinaryArithmeticNode.InitWithLeftSourceRightSource]
//   - [IMPSNNBinaryArithmeticNode.InitWithSources]
//
// # Instance Properties
//
//   - [IMPSNNBinaryArithmeticNode.Bias]
//   - [IMPSNNBinaryArithmeticNode.SetBias]
//   - [IMPSNNBinaryArithmeticNode.MaximumValue]
//   - [IMPSNNBinaryArithmeticNode.SetMaximumValue]
//   - [IMPSNNBinaryArithmeticNode.MinimumValue]
//   - [IMPSNNBinaryArithmeticNode.SetMinimumValue]
//   - [IMPSNNBinaryArithmeticNode.PrimaryScale]
//   - [IMPSNNBinaryArithmeticNode.SetPrimaryScale]
//   - [IMPSNNBinaryArithmeticNode.PrimaryStrideInFeatureChannels]
//   - [IMPSNNBinaryArithmeticNode.SetPrimaryStrideInFeatureChannels]
//   - [IMPSNNBinaryArithmeticNode.PrimaryStrideInPixelsX]
//   - [IMPSNNBinaryArithmeticNode.SetPrimaryStrideInPixelsX]
//   - [IMPSNNBinaryArithmeticNode.PrimaryStrideInPixelsY]
//   - [IMPSNNBinaryArithmeticNode.SetPrimaryStrideInPixelsY]
//   - [IMPSNNBinaryArithmeticNode.SecondaryScale]
//   - [IMPSNNBinaryArithmeticNode.SetSecondaryScale]
//   - [IMPSNNBinaryArithmeticNode.SecondaryStrideInFeatureChannels]
//   - [IMPSNNBinaryArithmeticNode.SetSecondaryStrideInFeatureChannels]
//   - [IMPSNNBinaryArithmeticNode.SecondaryStrideInPixelsX]
//   - [IMPSNNBinaryArithmeticNode.SetSecondaryStrideInPixelsX]
//   - [IMPSNNBinaryArithmeticNode.SecondaryStrideInPixelsY]
//   - [IMPSNNBinaryArithmeticNode.SetSecondaryStrideInPixelsY]
//
// # Instance Methods
//
//   - [IMPSNNBinaryArithmeticNode.GradientClass]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode
type IMPSNNBinaryArithmeticNode interface {
	IMPSNNFilterNode

	// Topic: Initializers

	InitWithLeftSourceRightSource(left IMPSNNImageNode, right IMPSNNImageNode) MPSNNBinaryArithmeticNode
	InitWithSources(sourceNodes []MPSNNImageNode) MPSNNBinaryArithmeticNode

	// Topic: Instance Properties

	Bias() float32
	SetBias(value float32)
	MaximumValue() float32
	SetMaximumValue(value float32)
	MinimumValue() float32
	SetMinimumValue(value float32)
	PrimaryScale() float32
	SetPrimaryScale(value float32)
	PrimaryStrideInFeatureChannels() uint
	SetPrimaryStrideInFeatureChannels(value uint)
	PrimaryStrideInPixelsX() uint
	SetPrimaryStrideInPixelsX(value uint)
	PrimaryStrideInPixelsY() uint
	SetPrimaryStrideInPixelsY(value uint)
	SecondaryScale() float32
	SetSecondaryScale(value float32)
	SecondaryStrideInFeatureChannels() uint
	SetSecondaryStrideInFeatureChannels(value uint)
	SecondaryStrideInPixelsX() uint
	SetSecondaryStrideInPixelsX(value uint)
	SecondaryStrideInPixelsY() uint
	SetSecondaryStrideInPixelsY(value uint)

	// Topic: Instance Methods

	GradientClass() objectivec.Class
}

// Init initializes the instance.
func (b MPSNNBinaryArithmeticNode) Init() MPSNNBinaryArithmeticNode {
	rv := objc.Send[MPSNNBinaryArithmeticNode](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MPSNNBinaryArithmeticNode) Autorelease() MPSNNBinaryArithmeticNode {
	rv := objc.Send[MPSNNBinaryArithmeticNode](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNBinaryArithmeticNode creates a new MPSNNBinaryArithmeticNode instance.
func NewMPSNNBinaryArithmeticNode() MPSNNBinaryArithmeticNode {
	class := getMPSNNBinaryArithmeticNodeClass()
	rv := objc.Send[MPSNNBinaryArithmeticNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(leftSource:rightSource:)
func NewBinaryArithmeticNodeWithLeftSourceRightSource(left IMPSNNImageNode, right IMPSNNImageNode) MPSNNBinaryArithmeticNode {
	instance := getMPSNNBinaryArithmeticNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), left, right)
	return MPSNNBinaryArithmeticNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(sources:)
func NewBinaryArithmeticNodeWithSources(sourceNodes []MPSNNImageNode) MPSNNBinaryArithmeticNode {
	instance := getMPSNNBinaryArithmeticNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), objectivec.IObjectSliceToNSArray(sourceNodes))
	return MPSNNBinaryArithmeticNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(leftSource:rightSource:)
func (b MPSNNBinaryArithmeticNode) InitWithLeftSourceRightSource(left IMPSNNImageNode, right IMPSNNImageNode) MPSNNBinaryArithmeticNode {
	rv := objc.Send[MPSNNBinaryArithmeticNode](b.ID, objc.Sel("initWithLeftSource:rightSource:"), left, right)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/init(sources:)
func (b MPSNNBinaryArithmeticNode) InitWithSources(sourceNodes []MPSNNImageNode) MPSNNBinaryArithmeticNode {
	rv := objc.Send[MPSNNBinaryArithmeticNode](b.ID, objc.Sel("initWithSources:"), objectivec.IObjectSliceToNSArray(sourceNodes))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/gradientClass()
func (b MPSNNBinaryArithmeticNode) GradientClass() objectivec.Class {
	rv := objc.Send[objectivec.Class](b.ID, objc.Sel("gradientClass"))
	return objectivec.Class(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/nodeWithLeftSource:rightSource:
func (_MPSNNBinaryArithmeticNodeClass MPSNNBinaryArithmeticNodeClass) NodeWithLeftSourceRightSource(left IMPSNNImageNode, right IMPSNNImageNode) MPSNNBinaryArithmeticNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNBinaryArithmeticNodeClass.class), objc.Sel("nodeWithLeftSource:rightSource:"), left, right)
	return MPSNNBinaryArithmeticNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/nodeWithSources:
func (_MPSNNBinaryArithmeticNodeClass MPSNNBinaryArithmeticNodeClass) NodeWithSources(sourceNodes []MPSNNImageNode) MPSNNBinaryArithmeticNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNBinaryArithmeticNodeClass.class), objc.Sel("nodeWithSources:"), objectivec.IObjectSliceToNSArray(sourceNodes))
	return MPSNNBinaryArithmeticNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/bias
func (b MPSNNBinaryArithmeticNode) Bias() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("bias"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetBias(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setBias:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/maximumValue
func (b MPSNNBinaryArithmeticNode) MaximumValue() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("maximumValue"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetMaximumValue(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setMaximumValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/minimumValue
func (b MPSNNBinaryArithmeticNode) MinimumValue() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("minimumValue"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetMinimumValue(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setMinimumValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/primaryScale
func (b MPSNNBinaryArithmeticNode) PrimaryScale() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("primaryScale"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetPrimaryScale(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrimaryScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/primaryStrideInFeatureChannels
func (b MPSNNBinaryArithmeticNode) PrimaryStrideInFeatureChannels() uint {
	rv := objc.Send[uint](b.ID, objc.Sel("primaryStrideInFeatureChannels"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetPrimaryStrideInFeatureChannels(value uint) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrimaryStrideInFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/primaryStrideInPixelsX
func (b MPSNNBinaryArithmeticNode) PrimaryStrideInPixelsX() uint {
	rv := objc.Send[uint](b.ID, objc.Sel("primaryStrideInPixelsX"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetPrimaryStrideInPixelsX(value uint) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/primaryStrideInPixelsY
func (b MPSNNBinaryArithmeticNode) PrimaryStrideInPixelsY() uint {
	rv := objc.Send[uint](b.ID, objc.Sel("primaryStrideInPixelsY"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetPrimaryStrideInPixelsY(value uint) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/secondaryScale
func (b MPSNNBinaryArithmeticNode) SecondaryScale() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("secondaryScale"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetSecondaryScale(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setSecondaryScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/secondaryStrideInFeatureChannels
func (b MPSNNBinaryArithmeticNode) SecondaryStrideInFeatureChannels() uint {
	rv := objc.Send[uint](b.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetSecondaryStrideInFeatureChannels(value uint) {
	objc.Send[struct{}](b.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/secondaryStrideInPixelsX
func (b MPSNNBinaryArithmeticNode) SecondaryStrideInPixelsX() uint {
	rv := objc.Send[uint](b.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetSecondaryStrideInPixelsX(value uint) {
	objc.Send[struct{}](b.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode/secondaryStrideInPixelsY
func (b MPSNNBinaryArithmeticNode) SecondaryStrideInPixelsY() uint {
	rv := objc.Send[uint](b.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetSecondaryStrideInPixelsY(value uint) {
	objc.Send[struct{}](b.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}

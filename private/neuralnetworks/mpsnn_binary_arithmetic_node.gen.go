// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

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

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNBinaryArithmeticNodeClass) Alloc() MPSNNBinaryArithmeticNode {
	rv := objc.Send[MPSNNBinaryArithmeticNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
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
//   - [MPSNNBinaryArithmeticNode.InitWithLeftSourceRightSource]
//   - [MPSNNBinaryArithmeticNode.InitWithSources]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode
type MPSNNBinaryArithmeticNode struct {
	MPSNNFilterNode
}

// MPSNNBinaryArithmeticNodeFromID constructs a [MPSNNBinaryArithmeticNode] from an objc.ID.
func MPSNNBinaryArithmeticNodeFromID(id objc.ID) MPSNNBinaryArithmeticNode {
	return MPSNNBinaryArithmeticNode{MPSNNFilterNode: MPSNNFilterNodeFromID(id)}
}
// Ensure MPSNNBinaryArithmeticNode implements IMPSNNBinaryArithmeticNode.
var _ IMPSNNBinaryArithmeticNode = MPSNNBinaryArithmeticNode{}





// An interface definition for the [MPSNNBinaryArithmeticNode] class.
//
// # Methods
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
//   - [IMPSNNBinaryArithmeticNode.InitWithLeftSourceRightSource]
//   - [IMPSNNBinaryArithmeticNode.InitWithSources]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode
type IMPSNNBinaryArithmeticNode interface {
	IMPSNNFilterNode

	// Topic: Methods

	Bias() float32
	SetBias(value float32)
	MaximumValue() float32
	SetMaximumValue(value float32)
	MinimumValue() float32
	SetMinimumValue(value float32)
	PrimaryScale() float32
	SetPrimaryScale(value float32)
	PrimaryStrideInFeatureChannels() uint64
	SetPrimaryStrideInFeatureChannels(value uint64)
	PrimaryStrideInPixelsX() uint64
	SetPrimaryStrideInPixelsX(value uint64)
	PrimaryStrideInPixelsY() uint64
	SetPrimaryStrideInPixelsY(value uint64)
	SecondaryScale() float32
	SetSecondaryScale(value float32)
	SecondaryStrideInFeatureChannels() uint64
	SetSecondaryStrideInFeatureChannels(value uint64)
	SecondaryStrideInPixelsX() uint64
	SetSecondaryStrideInPixelsX(value uint64)
	SecondaryStrideInPixelsY() uint64
	SetSecondaryStrideInPixelsY(value uint64)
	InitWithLeftSourceRightSource(source objectivec.IObject, source2 objectivec.IObject) MPSNNBinaryArithmeticNode
	InitWithSources(sources objectivec.IObject) MPSNNBinaryArithmeticNode
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






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithLeftSource:rightSource:
func NewBinaryArithmeticNodeWithLeftSourceRightSource(source objectivec.IObject, source2 objectivec.IObject) MPSNNBinaryArithmeticNode {
	instance := getMPSNNBinaryArithmeticNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftSource:rightSource:"), source, source2)
	return MPSNNBinaryArithmeticNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewBinaryArithmeticNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNBinaryArithmeticNode {
	instance := getMPSNNBinaryArithmeticNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNBinaryArithmeticNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithSources:
func NewBinaryArithmeticNodeWithSources(sources objectivec.IObject) MPSNNBinaryArithmeticNode {
	instance := getMPSNNBinaryArithmeticNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSources:"), sources)
	return MPSNNBinaryArithmeticNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithLeftSource:rightSource:
func (b MPSNNBinaryArithmeticNode) InitWithLeftSourceRightSource(source objectivec.IObject, source2 objectivec.IObject) MPSNNBinaryArithmeticNode {
	rv := objc.Send[MPSNNBinaryArithmeticNode](b.ID, objc.Sel("initWithLeftSource:rightSource:"), source, source2)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/initWithSources:
func (b MPSNNBinaryArithmeticNode) InitWithSources(sources objectivec.IObject) MPSNNBinaryArithmeticNode {
	rv := objc.Send[MPSNNBinaryArithmeticNode](b.ID, objc.Sel("initWithSources:"), sources)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/nodeWithLeftSource:rightSource:
func (_MPSNNBinaryArithmeticNodeClass MPSNNBinaryArithmeticNodeClass) NodeWithLeftSourceRightSource(source objectivec.IObject, source2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNBinaryArithmeticNodeClass.class), objc.Sel("nodeWithLeftSource:rightSource:"), source, source2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/nodeWithSources:
func (_MPSNNBinaryArithmeticNodeClass MPSNNBinaryArithmeticNodeClass) NodeWithSources(sources objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNBinaryArithmeticNodeClass.class), objc.Sel("nodeWithSources:"), sources)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/bias
func (b MPSNNBinaryArithmeticNode) Bias() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("bias"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetBias(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setBias:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/maximumValue
func (b MPSNNBinaryArithmeticNode) MaximumValue() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("maximumValue"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetMaximumValue(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setMaximumValue:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/minimumValue
func (b MPSNNBinaryArithmeticNode) MinimumValue() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("minimumValue"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetMinimumValue(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setMinimumValue:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/primaryScale
func (b MPSNNBinaryArithmeticNode) PrimaryScale() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("primaryScale"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetPrimaryScale(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrimaryScale:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/primaryStrideInFeatureChannels
func (b MPSNNBinaryArithmeticNode) PrimaryStrideInFeatureChannels() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("primaryStrideInFeatureChannels"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetPrimaryStrideInFeatureChannels(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrimaryStrideInFeatureChannels:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/primaryStrideInPixelsX
func (b MPSNNBinaryArithmeticNode) PrimaryStrideInPixelsX() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("primaryStrideInPixelsX"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetPrimaryStrideInPixelsX(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/primaryStrideInPixelsY
func (b MPSNNBinaryArithmeticNode) PrimaryStrideInPixelsY() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("primaryStrideInPixelsY"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetPrimaryStrideInPixelsY(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/secondaryScale
func (b MPSNNBinaryArithmeticNode) SecondaryScale() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("secondaryScale"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetSecondaryScale(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setSecondaryScale:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/secondaryStrideInFeatureChannels
func (b MPSNNBinaryArithmeticNode) SecondaryStrideInFeatureChannels() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetSecondaryStrideInFeatureChannels(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/secondaryStrideInPixelsX
func (b MPSNNBinaryArithmeticNode) SecondaryStrideInPixelsX() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetSecondaryStrideInPixelsX(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNBinaryArithmeticNode/secondaryStrideInPixelsY
func (b MPSNNBinaryArithmeticNode) SecondaryStrideInPixelsY() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}
func (b MPSNNBinaryArithmeticNode) SetSecondaryStrideInPixelsY(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}


















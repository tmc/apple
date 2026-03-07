// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

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

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNArithmeticGradientNodeClass) Alloc() MPSNNArithmeticGradientNode {
	rv := objc.Send[MPSNNArithmeticGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
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
//   - [MPSNNArithmeticGradientNode.InitWithGradientImagesForwardFilterIsSecondarySourceFilter]
//   - [MPSNNArithmeticGradientNode.InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode
type MPSNNArithmeticGradientNode struct {
	MPSNNGradientFilterNode
}

// MPSNNArithmeticGradientNodeFromID constructs a [MPSNNArithmeticGradientNode] from an objc.ID.
func MPSNNArithmeticGradientNodeFromID(id objc.ID) MPSNNArithmeticGradientNode {
	return MPSNNArithmeticGradientNode{MPSNNGradientFilterNode: MPSNNGradientFilterNodeFromID(id)}
}
// Ensure MPSNNArithmeticGradientNode implements IMPSNNArithmeticGradientNode.
var _ IMPSNNArithmeticGradientNode = MPSNNArithmeticGradientNode{}





// An interface definition for the [MPSNNArithmeticGradientNode] class.
//
// # Methods
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
//   - [IMPSNNArithmeticGradientNode.InitWithGradientImagesForwardFilterIsSecondarySourceFilter]
//   - [IMPSNNArithmeticGradientNode.InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode
type IMPSNNArithmeticGradientNode interface {
	IMPSNNGradientFilterNode

	// Topic: Methods

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
	SecondaryStrideInFeatureChannels() uint64
	SetSecondaryStrideInFeatureChannels(value uint64)
	SecondaryStrideInPixelsX() uint64
	SetSecondaryStrideInPixelsX(value uint64)
	SecondaryStrideInPixelsY() uint64
	SetSecondaryStrideInPixelsY(value uint64)
	InitWithGradientImagesForwardFilterIsSecondarySourceFilter(images objectivec.IObject, filter objectivec.IObject, filter2 bool) MPSNNArithmeticGradientNode
	InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, filter bool) MPSNNArithmeticGradientNode
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






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:forwardFilter:
func NewArithmeticGradientNodeWithGradientImagesForwardFilter(images objectivec.IObject, filter objectivec.IObject) MPSNNArithmeticGradientNode {
	instance := getMPSNNArithmeticGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:"), images, filter)
	return MPSNNArithmeticGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/initWithGradientImages:forwardFilter:isSecondarySourceFilter:
func NewArithmeticGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(images objectivec.IObject, filter objectivec.IObject, filter2 bool) MPSNNArithmeticGradientNode {
	instance := getMPSNNArithmeticGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), images, filter, filter2)
	return MPSNNArithmeticGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:
func NewArithmeticGradientNodeWithGradientImagesSourceImagesBinaryGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNArithmeticGradientNode {
	instance := getMPSNNArithmeticGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:binaryGradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNArithmeticGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGradientFilterNode/initWithGradientImages:sourceImages:gradientState:paddingPolicy:
func NewArithmeticGradientNodeWithGradientImagesSourceImagesGradientStatePaddingPolicy(images objectivec.IObject, images2 objectivec.IObject, state objectivec.IObject, policy objectivec.IObject) MPSNNArithmeticGradientNode {
	instance := getMPSNNArithmeticGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGradientImages:sourceImages:gradientState:paddingPolicy:"), images, images2, state, policy)
	return MPSNNArithmeticGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:
func NewArithmeticGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, filter bool) MPSNNArithmeticGradientNode {
	instance := getMPSNNArithmeticGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), gradient, image, state, filter)
	return MPSNNArithmeticGradientNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNFilterNode/initWithSourceImages:sourceStates:paddingPolicy:
func NewArithmeticGradientNodeWithSourceImagesSourceStatesPaddingPolicy(images objectivec.IObject, states objectivec.IObject, policy objectivec.IObject) MPSNNArithmeticGradientNode {
	instance := getMPSNNArithmeticGradientNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceImages:sourceStates:paddingPolicy:"), images, states, policy)
	return MPSNNArithmeticGradientNodeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/initWithGradientImages:forwardFilter:isSecondarySourceFilter:
func (a MPSNNArithmeticGradientNode) InitWithGradientImagesForwardFilterIsSecondarySourceFilter(images objectivec.IObject, filter objectivec.IObject, filter2 bool) MPSNNArithmeticGradientNode {
	rv := objc.Send[MPSNNArithmeticGradientNode](a.ID, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), images, filter, filter2)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:
func (a MPSNNArithmeticGradientNode) InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, filter bool) MPSNNArithmeticGradientNode {
	rv := objc.Send[MPSNNArithmeticGradientNode](a.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), gradient, image, state, filter)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/nodeWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:
func (_MPSNNArithmeticGradientNodeClass MPSNNArithmeticGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(gradient objectivec.IObject, image objectivec.IObject, state objectivec.IObject, filter bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNArithmeticGradientNodeClass.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), gradient, image, state, filter)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/bias
func (a MPSNNArithmeticGradientNode) Bias() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("bias"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetBias(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setBias:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/isSecondarySourceFilter
func (a MPSNNArithmeticGradientNode) IsSecondarySourceFilter() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isSecondarySourceFilter"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/maximumValue
func (a MPSNNArithmeticGradientNode) MaximumValue() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("maximumValue"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetMaximumValue(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setMaximumValue:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/minimumValue
func (a MPSNNArithmeticGradientNode) MinimumValue() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("minimumValue"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetMinimumValue(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setMinimumValue:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/primaryScale
func (a MPSNNArithmeticGradientNode) PrimaryScale() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("primaryScale"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetPrimaryScale(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPrimaryScale:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/secondaryScale
func (a MPSNNArithmeticGradientNode) SecondaryScale() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("secondaryScale"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetSecondaryScale(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setSecondaryScale:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/secondaryStrideInFeatureChannels
func (a MPSNNArithmeticGradientNode) SecondaryStrideInFeatureChannels() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetSecondaryStrideInFeatureChannels(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/secondaryStrideInPixelsX
func (a MPSNNArithmeticGradientNode) SecondaryStrideInPixelsX() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetSecondaryStrideInPixelsX(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNArithmeticGradientNode/secondaryStrideInPixelsY
func (a MPSNNArithmeticGradientNode) SecondaryStrideInPixelsY() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}
func (a MPSNNArithmeticGradientNode) SetSecondaryStrideInPixelsY(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}


















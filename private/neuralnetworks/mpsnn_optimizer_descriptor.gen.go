// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNOptimizerDescriptor] class.
var (
	_MPSNNOptimizerDescriptorClass     MPSNNOptimizerDescriptorClass
	_MPSNNOptimizerDescriptorClassOnce sync.Once
)

func getMPSNNOptimizerDescriptorClass() MPSNNOptimizerDescriptorClass {
	_MPSNNOptimizerDescriptorClassOnce.Do(func() {
		_MPSNNOptimizerDescriptorClass = MPSNNOptimizerDescriptorClass{class: objc.GetClass("MPSNNOptimizerDescriptor")}
	})
	return _MPSNNOptimizerDescriptorClass
}

// GetMPSNNOptimizerDescriptorClass returns the class object for MPSNNOptimizerDescriptor.
func GetMPSNNOptimizerDescriptorClass() MPSNNOptimizerDescriptorClass {
	return getMPSNNOptimizerDescriptorClass()
}

type MPSNNOptimizerDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNOptimizerDescriptorClass) Alloc() MPSNNOptimizerDescriptor {
	rv := objc.Send[MPSNNOptimizerDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNOptimizerDescriptor.ApplyGradientClipping]
//   - [MPSNNOptimizerDescriptor.SetApplyGradientClipping]
//   - [MPSNNOptimizerDescriptor.GradientClipMax]
//   - [MPSNNOptimizerDescriptor.SetGradientClipMax]
//   - [MPSNNOptimizerDescriptor.GradientClipMin]
//   - [MPSNNOptimizerDescriptor.SetGradientClipMin]
//   - [MPSNNOptimizerDescriptor.GradientRescale]
//   - [MPSNNOptimizerDescriptor.SetGradientRescale]
//   - [MPSNNOptimizerDescriptor.LearningRate]
//   - [MPSNNOptimizerDescriptor.SetLearningRate]
//   - [MPSNNOptimizerDescriptor.RegularizationScale]
//   - [MPSNNOptimizerDescriptor.SetRegularizationScale]
//   - [MPSNNOptimizerDescriptor.RegularizationType]
//   - [MPSNNOptimizerDescriptor.SetRegularizationType]
//   - [MPSNNOptimizerDescriptor.InitWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale]
//   - [MPSNNOptimizerDescriptor.InitWithLearningRateGradientRescaleRegularizationTypeRegularizationScale]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor
type MPSNNOptimizerDescriptor struct {
	objectivec.Object
}

// MPSNNOptimizerDescriptorFromID constructs a [MPSNNOptimizerDescriptor] from an objc.ID.
func MPSNNOptimizerDescriptorFromID(id objc.ID) MPSNNOptimizerDescriptor {
	return MPSNNOptimizerDescriptor{objectivec.Object{id}}
}
// Ensure MPSNNOptimizerDescriptor implements IMPSNNOptimizerDescriptor.
var _ IMPSNNOptimizerDescriptor = MPSNNOptimizerDescriptor{}





// An interface definition for the [MPSNNOptimizerDescriptor] class.
//
// # Methods
//
//   - [IMPSNNOptimizerDescriptor.ApplyGradientClipping]
//   - [IMPSNNOptimizerDescriptor.SetApplyGradientClipping]
//   - [IMPSNNOptimizerDescriptor.GradientClipMax]
//   - [IMPSNNOptimizerDescriptor.SetGradientClipMax]
//   - [IMPSNNOptimizerDescriptor.GradientClipMin]
//   - [IMPSNNOptimizerDescriptor.SetGradientClipMin]
//   - [IMPSNNOptimizerDescriptor.GradientRescale]
//   - [IMPSNNOptimizerDescriptor.SetGradientRescale]
//   - [IMPSNNOptimizerDescriptor.LearningRate]
//   - [IMPSNNOptimizerDescriptor.SetLearningRate]
//   - [IMPSNNOptimizerDescriptor.RegularizationScale]
//   - [IMPSNNOptimizerDescriptor.SetRegularizationScale]
//   - [IMPSNNOptimizerDescriptor.RegularizationType]
//   - [IMPSNNOptimizerDescriptor.SetRegularizationType]
//   - [IMPSNNOptimizerDescriptor.InitWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale]
//   - [IMPSNNOptimizerDescriptor.InitWithLearningRateGradientRescaleRegularizationTypeRegularizationScale]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor
type IMPSNNOptimizerDescriptor interface {
	objectivec.IObject

	// Topic: Methods

	ApplyGradientClipping() bool
	SetApplyGradientClipping(value bool)
	GradientClipMax() float32
	SetGradientClipMax(value float32)
	GradientClipMin() float32
	SetGradientClipMin(value float32)
	GradientRescale() float32
	SetGradientRescale(value float32)
	LearningRate() float32
	SetLearningRate(value float32)
	RegularizationScale() float32
	SetRegularizationScale(value float32)
	RegularizationType() uint64
	SetRegularizationType(value uint64)
	InitWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(rate float32, rescale float32, clipping bool, max float32, min float32, type_ uint64, scale float32) MPSNNOptimizerDescriptor
	InitWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(rate float32, rescale float32, type_ uint64, scale float32) MPSNNOptimizerDescriptor
}





// Init initializes the instance.
func (o MPSNNOptimizerDescriptor) Init() MPSNNOptimizerDescriptor {
	rv := objc.Send[MPSNNOptimizerDescriptor](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o MPSNNOptimizerDescriptor) Autorelease() MPSNNOptimizerDescriptor {
	rv := objc.Send[MPSNNOptimizerDescriptor](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNOptimizerDescriptor creates a new MPSNNOptimizerDescriptor instance.
func NewMPSNNOptimizerDescriptor() MPSNNOptimizerDescriptor {
	class := getMPSNNOptimizerDescriptorClass()
	rv := objc.Send[MPSNNOptimizerDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/initWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:
func NewOptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(rate float32, rescale float32, clipping bool, max float32, min float32, type_ uint64, scale float32) MPSNNOptimizerDescriptor {
	instance := getMPSNNOptimizerDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:"), rate, rescale, clipping, max, min, type_, scale)
	return MPSNNOptimizerDescriptorFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/initWithLearningRate:gradientRescale:regularizationType:regularizationScale:
func NewOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(rate float32, rescale float32, type_ uint64, scale float32) MPSNNOptimizerDescriptor {
	instance := getMPSNNOptimizerDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), rate, rescale, type_, scale)
	return MPSNNOptimizerDescriptorFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/initWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:
func (o MPSNNOptimizerDescriptor) InitWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(rate float32, rescale float32, clipping bool, max float32, min float32, type_ uint64, scale float32) MPSNNOptimizerDescriptor {
	rv := objc.Send[MPSNNOptimizerDescriptor](o.ID, objc.Sel("initWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:"), rate, rescale, clipping, max, min, type_, scale)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/initWithLearningRate:gradientRescale:regularizationType:regularizationScale:
func (o MPSNNOptimizerDescriptor) InitWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(rate float32, rescale float32, type_ uint64, scale float32) MPSNNOptimizerDescriptor {
	rv := objc.Send[MPSNNOptimizerDescriptor](o.ID, objc.Sel("initWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), rate, rescale, type_, scale)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/optimizerDescriptorWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:
func (_MPSNNOptimizerDescriptorClass MPSNNOptimizerDescriptorClass) OptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(rate float32, rescale float32, clipping bool, max float32, min float32, type_ uint64, scale float32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNOptimizerDescriptorClass.class), objc.Sel("optimizerDescriptorWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:"), rate, rescale, clipping, max, min, type_, scale)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/optimizerDescriptorWithLearningRate:gradientRescale:regularizationType:regularizationScale:
func (_MPSNNOptimizerDescriptorClass MPSNNOptimizerDescriptorClass) OptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(rate float32, rescale float32, type_ uint64, scale float32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNOptimizerDescriptorClass.class), objc.Sel("optimizerDescriptorWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), rate, rescale, type_, scale)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/applyGradientClipping
func (o MPSNNOptimizerDescriptor) ApplyGradientClipping() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("applyGradientClipping"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetApplyGradientClipping(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setApplyGradientClipping:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/gradientClipMax
func (o MPSNNOptimizerDescriptor) GradientClipMax() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientClipMax"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetGradientClipMax(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setGradientClipMax:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/gradientClipMin
func (o MPSNNOptimizerDescriptor) GradientClipMin() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientClipMin"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetGradientClipMin(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setGradientClipMin:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/gradientRescale
func (o MPSNNOptimizerDescriptor) GradientRescale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientRescale"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetGradientRescale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setGradientRescale:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/learningRate
func (o MPSNNOptimizerDescriptor) LearningRate() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("learningRate"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetLearningRate(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setLearningRate:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/regularizationScale
func (o MPSNNOptimizerDescriptor) RegularizationScale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("regularizationScale"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetRegularizationScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRegularizationScale:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizerDescriptor/regularizationType
func (o MPSNNOptimizerDescriptor) RegularizationType() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("regularizationType"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetRegularizationType(value uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setRegularizationType:"), value)
}


















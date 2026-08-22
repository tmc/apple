// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

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

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNOptimizerDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNOptimizerDescriptorClass) Alloc() MPSNNOptimizerDescriptor {
	rv := objc.Send[MPSNNOptimizerDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that specifies properties used by an optimizer kernel.
//
// # Initializers
//
//   - [MPSNNOptimizerDescriptor.InitWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale]
//   - [MPSNNOptimizerDescriptor.InitWithLearningRateGradientRescaleRegularizationTypeRegularizationScale]
//
// # Instance Properties
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
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor
type MPSNNOptimizerDescriptor struct {
	objectivec.Object
}

// MPSNNOptimizerDescriptorFromID constructs a [MPSNNOptimizerDescriptor] from an objc.ID.
//
// An object that specifies properties used by an optimizer kernel.
func MPSNNOptimizerDescriptorFromID(id objc.ID) MPSNNOptimizerDescriptor {
	return MPSNNOptimizerDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSNNOptimizerDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNOptimizerDescriptor] class.
//
// # Initializers
//
//   - [IMPSNNOptimizerDescriptor.InitWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale]
//   - [IMPSNNOptimizerDescriptor.InitWithLearningRateGradientRescaleRegularizationTypeRegularizationScale]
//
// # Instance Properties
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
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor
type IMPSNNOptimizerDescriptor interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, applyGradientClipping bool, gradientClipMax float32, gradientClipMin float32, regularizationType MPSNNRegularizationType, regularizationScale float32) MPSNNOptimizerDescriptor
	InitWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, regularizationType MPSNNRegularizationType, regularizationScale float32) MPSNNOptimizerDescriptor

	// Topic: Instance Properties

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
	RegularizationType() MPSNNRegularizationType
	SetRegularizationType(value MPSNNRegularizationType)
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

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/init(learningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:)
func NewOptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, applyGradientClipping bool, gradientClipMax float32, gradientClipMin float32, regularizationType MPSNNRegularizationType, regularizationScale float32) MPSNNOptimizerDescriptor {
	instance := getMPSNNOptimizerDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:"), learningRate, gradientRescale, applyGradientClipping, gradientClipMax, gradientClipMin, regularizationType, regularizationScale)
	return MPSNNOptimizerDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/init(learningRate:gradientRescale:regularizationType:regularizationScale:)
func NewOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, regularizationType MPSNNRegularizationType, regularizationScale float32) MPSNNOptimizerDescriptor {
	instance := getMPSNNOptimizerDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), learningRate, gradientRescale, regularizationType, regularizationScale)
	return MPSNNOptimizerDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/init(learningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:)
func (o MPSNNOptimizerDescriptor) InitWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, applyGradientClipping bool, gradientClipMax float32, gradientClipMin float32, regularizationType MPSNNRegularizationType, regularizationScale float32) MPSNNOptimizerDescriptor {
	rv := objc.Send[MPSNNOptimizerDescriptor](o.ID, objc.Sel("initWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:"), learningRate, gradientRescale, applyGradientClipping, gradientClipMax, gradientClipMin, regularizationType, regularizationScale)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/init(learningRate:gradientRescale:regularizationType:regularizationScale:)
func (o MPSNNOptimizerDescriptor) InitWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, regularizationType MPSNNRegularizationType, regularizationScale float32) MPSNNOptimizerDescriptor {
	rv := objc.Send[MPSNNOptimizerDescriptor](o.ID, objc.Sel("initWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), learningRate, gradientRescale, regularizationType, regularizationScale)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/optimizerDescriptorWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:
func (_MPSNNOptimizerDescriptorClass MPSNNOptimizerDescriptorClass) OptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, applyGradientClipping bool, gradientClipMax float32, gradientClipMin float32, regularizationType MPSNNRegularizationType, regularizationScale float32) MPSNNOptimizerDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNOptimizerDescriptorClass.class), objc.Sel("optimizerDescriptorWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:"), learningRate, gradientRescale, applyGradientClipping, gradientClipMax, gradientClipMin, regularizationType, regularizationScale)
	return MPSNNOptimizerDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/optimizerDescriptorWithLearningRate:gradientRescale:regularizationType:regularizationScale:
func (_MPSNNOptimizerDescriptorClass MPSNNOptimizerDescriptorClass) OptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, regularizationType MPSNNRegularizationType, regularizationScale float32) MPSNNOptimizerDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNOptimizerDescriptorClass.class), objc.Sel("optimizerDescriptorWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), learningRate, gradientRescale, regularizationType, regularizationScale)
	return MPSNNOptimizerDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/applyGradientClipping
func (o MPSNNOptimizerDescriptor) ApplyGradientClipping() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("applyGradientClipping"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetApplyGradientClipping(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setApplyGradientClipping:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/gradientClipMax
func (o MPSNNOptimizerDescriptor) GradientClipMax() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientClipMax"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetGradientClipMax(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setGradientClipMax:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/gradientClipMin
func (o MPSNNOptimizerDescriptor) GradientClipMin() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientClipMin"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetGradientClipMin(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setGradientClipMin:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/gradientRescale
func (o MPSNNOptimizerDescriptor) GradientRescale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientRescale"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetGradientRescale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setGradientRescale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/learningRate
func (o MPSNNOptimizerDescriptor) LearningRate() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("learningRate"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetLearningRate(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setLearningRate:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/regularizationScale
func (o MPSNNOptimizerDescriptor) RegularizationScale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("regularizationScale"))
	return rv
}
func (o MPSNNOptimizerDescriptor) SetRegularizationScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRegularizationScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor/regularizationType
func (o MPSNNOptimizerDescriptor) RegularizationType() MPSNNRegularizationType {
	rv := objc.Send[MPSNNRegularizationType](o.ID, objc.Sel("regularizationType"))
	return MPSNNRegularizationType(rv)
}
func (o MPSNNOptimizerDescriptor) SetRegularizationType(value MPSNNRegularizationType) {
	objc.Send[struct{}](o.ID, objc.Sel("setRegularizationType:"), value)
}

// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNOptimizer] class.
var (
	_MPSNNOptimizerClass     MPSNNOptimizerClass
	_MPSNNOptimizerClassOnce sync.Once
)

func getMPSNNOptimizerClass() MPSNNOptimizerClass {
	_MPSNNOptimizerClassOnce.Do(func() {
		_MPSNNOptimizerClass = MPSNNOptimizerClass{class: objc.GetClass("MPSNNOptimizer")}
	})
	return _MPSNNOptimizerClass
}

// GetMPSNNOptimizerClass returns the class object for MPSNNOptimizer.
func GetMPSNNOptimizerClass() MPSNNOptimizerClass {
	return getMPSNNOptimizerClass()
}

type MPSNNOptimizerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNOptimizerClass) Alloc() MPSNNOptimizer {
	rv := objc.Send[MPSNNOptimizer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNOptimizer.ApplyGradientClipping]
//   - [MPSNNOptimizer.SetApplyGradientClipping]
//   - [MPSNNOptimizer.CopyWithZoneDevice]
//   - [MPSNNOptimizer.GradientClipMax]
//   - [MPSNNOptimizer.GradientClipMin]
//   - [MPSNNOptimizer.GradientRescale]
//   - [MPSNNOptimizer.LearningRate]
//   - [MPSNNOptimizer.RegularizationScale]
//   - [MPSNNOptimizer.RegularizationType]
//   - [MPSNNOptimizer.InitWithCoderDevice]
//   - [MPSNNOptimizer.InitWithDevice]
//   - [MPSNNOptimizer.InitWithDeviceOptimizerDescriptor]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer
type MPSNNOptimizer struct {
	MPSKernel
}

// MPSNNOptimizerFromID constructs a [MPSNNOptimizer] from an objc.ID.
func MPSNNOptimizerFromID(id objc.ID) MPSNNOptimizer {
	return MPSNNOptimizer{MPSKernel: MPSKernelFromID(id)}
}
// Ensure MPSNNOptimizer implements IMPSNNOptimizer.
var _ IMPSNNOptimizer = MPSNNOptimizer{}





// An interface definition for the [MPSNNOptimizer] class.
//
// # Methods
//
//   - [IMPSNNOptimizer.ApplyGradientClipping]
//   - [IMPSNNOptimizer.SetApplyGradientClipping]
//   - [IMPSNNOptimizer.CopyWithZoneDevice]
//   - [IMPSNNOptimizer.GradientClipMax]
//   - [IMPSNNOptimizer.GradientClipMin]
//   - [IMPSNNOptimizer.GradientRescale]
//   - [IMPSNNOptimizer.LearningRate]
//   - [IMPSNNOptimizer.RegularizationScale]
//   - [IMPSNNOptimizer.RegularizationType]
//   - [IMPSNNOptimizer.InitWithCoderDevice]
//   - [IMPSNNOptimizer.InitWithDevice]
//   - [IMPSNNOptimizer.InitWithDeviceOptimizerDescriptor]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer
type IMPSNNOptimizer interface {
	IMPSKernel

	// Topic: Methods

	ApplyGradientClipping() bool
	SetApplyGradientClipping(value bool)
	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	GradientClipMax() float32
	GradientClipMin() float32
	GradientRescale() float32
	LearningRate() float32
	RegularizationScale() float32
	RegularizationType() uint64
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNOptimizer
	InitWithDevice(device objectivec.IObject) MPSNNOptimizer
	InitWithDeviceOptimizerDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNOptimizer
}





// Init initializes the instance.
func (o MPSNNOptimizer) Init() MPSNNOptimizer {
	rv := objc.Send[MPSNNOptimizer](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o MPSNNOptimizer) Autorelease() MPSNNOptimizer {
	rv := objc.Send[MPSNNOptimizer](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNOptimizer creates a new MPSNNOptimizer instance.
func NewMPSNNOptimizer() MPSNNOptimizer {
	class := getMPSNNOptimizerClass()
	rv := objc.Send[MPSNNOptimizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/initWithCoder:device:
func NewOptimizerWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNOptimizer {
	instance := getMPSNNOptimizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNOptimizerFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/initWithDevice:
func NewOptimizerWithDevice(device objectivec.IObject) MPSNNOptimizer {
	instance := getMPSNNOptimizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNOptimizerFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/initWithDevice:optimizerDescriptor:
func NewOptimizerWithDeviceOptimizerDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNOptimizer {
	instance := getMPSNNOptimizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:optimizerDescriptor:"), device, descriptor)
	return MPSNNOptimizerFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/copyWithZone:device:
func (o MPSNNOptimizer) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/setLearningRate:
func (o MPSNNOptimizer) SetLearningRate(rate float32) {
	objc.Send[objc.ID](o.ID, objc.Sel("setLearningRate:"), rate)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/initWithCoder:device:
func (o MPSNNOptimizer) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNOptimizer {
	rv := objc.Send[MPSNNOptimizer](o.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/initWithDevice:
func (o MPSNNOptimizer) InitWithDevice(device objectivec.IObject) MPSNNOptimizer {
	rv := objc.Send[MPSNNOptimizer](o.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/initWithDevice:optimizerDescriptor:
func (o MPSNNOptimizer) InitWithDeviceOptimizerDescriptor(device objectivec.IObject, descriptor objectivec.IObject) MPSNNOptimizer {
	rv := objc.Send[MPSNNOptimizer](o.ID, objc.Sel("initWithDevice:optimizerDescriptor:"), device, descriptor)
	return rv
}












// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/applyGradientClipping
func (o MPSNNOptimizer) ApplyGradientClipping() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("applyGradientClipping"))
	return rv
}
func (o MPSNNOptimizer) SetApplyGradientClipping(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setApplyGradientClipping:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/gradientClipMax
func (o MPSNNOptimizer) GradientClipMax() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientClipMax"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/gradientClipMin
func (o MPSNNOptimizer) GradientClipMin() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientClipMin"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/gradientRescale
func (o MPSNNOptimizer) GradientRescale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientRescale"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/learningRate
func (o MPSNNOptimizer) LearningRate() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("learningRate"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/regularizationScale
func (o MPSNNOptimizer) RegularizationScale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("regularizationScale"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNOptimizer/regularizationType
func (o MPSNNOptimizer) RegularizationType() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("regularizationType"))
	return rv
}


















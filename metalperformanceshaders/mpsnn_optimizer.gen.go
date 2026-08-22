// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNOptimizerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNOptimizerClass) Alloc() MPSNNOptimizer {
	rv := objc.Send[MPSNNOptimizer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The base class for optimization layers.
//
// # Instance Properties
//
//   - [MPSNNOptimizer.ApplyGradientClipping]
//   - [MPSNNOptimizer.SetApplyGradientClipping]
//   - [MPSNNOptimizer.GradientClipMax]
//   - [MPSNNOptimizer.GradientClipMin]
//   - [MPSNNOptimizer.GradientRescale]
//   - [MPSNNOptimizer.LearningRate]
//   - [MPSNNOptimizer.RegularizationScale]
//   - [MPSNNOptimizer.RegularizationType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizer
type MPSNNOptimizer struct {
	MPSKernel
}

// MPSNNOptimizerFromID constructs a [MPSNNOptimizer] from an objc.ID.
//
// The base class for optimization layers.
func MPSNNOptimizerFromID(id objc.ID) MPSNNOptimizer {
	return MPSNNOptimizer{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSNNOptimizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNOptimizer] class.
//
// # Instance Properties
//
//   - [IMPSNNOptimizer.ApplyGradientClipping]
//   - [IMPSNNOptimizer.SetApplyGradientClipping]
//   - [IMPSNNOptimizer.GradientClipMax]
//   - [IMPSNNOptimizer.GradientClipMin]
//   - [IMPSNNOptimizer.GradientRescale]
//   - [IMPSNNOptimizer.LearningRate]
//   - [IMPSNNOptimizer.RegularizationScale]
//   - [IMPSNNOptimizer.RegularizationType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizer
type IMPSNNOptimizer interface {
	IMPSKernel

	// Topic: Instance Properties

	ApplyGradientClipping() bool
	SetApplyGradientClipping(value bool)
	GradientClipMax() float32
	GradientClipMin() float32
	GradientRescale() float32
	LearningRate() float32
	RegularizationScale() float32
	RegularizationType() MPSNNRegularizationType
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

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewOptimizerWithCoder(aDecoder foundation.INSCoder) MPSNNOptimizer {
	instance := getMPSNNOptimizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNOptimizerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewOptimizerWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNOptimizer {
	instance := getMPSNNOptimizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNOptimizerFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewOptimizerWithDevice(device metal.MTLDevice) MPSNNOptimizer {
	instance := getMPSNNOptimizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNOptimizerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizer/applyGradientClipping
func (o MPSNNOptimizer) ApplyGradientClipping() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("applyGradientClipping"))
	return rv
}
func (o MPSNNOptimizer) SetApplyGradientClipping(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setApplyGradientClipping:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizer/gradientClipMax
func (o MPSNNOptimizer) GradientClipMax() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientClipMax"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizer/gradientClipMin
func (o MPSNNOptimizer) GradientClipMin() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientClipMin"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizer/gradientRescale
func (o MPSNNOptimizer) GradientRescale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("gradientRescale"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizer/learningRate
func (o MPSNNOptimizer) LearningRate() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("learningRate"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizer/regularizationScale
func (o MPSNNOptimizer) RegularizationScale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("regularizationScale"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizer/regularizationType
func (o MPSNNOptimizer) RegularizationType() MPSNNRegularizationType {
	rv := objc.Send[MPSNNRegularizationType](o.ID, objc.Sel("regularizationType"))
	return MPSNNRegularizationType(rv)
}

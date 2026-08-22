// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNGramMatrixCalculationGradient] class.
var (
	_MPSNNGramMatrixCalculationGradientClass     MPSNNGramMatrixCalculationGradientClass
	_MPSNNGramMatrixCalculationGradientClassOnce sync.Once
)

func getMPSNNGramMatrixCalculationGradientClass() MPSNNGramMatrixCalculationGradientClass {
	_MPSNNGramMatrixCalculationGradientClassOnce.Do(func() {
		_MPSNNGramMatrixCalculationGradientClass = MPSNNGramMatrixCalculationGradientClass{class: objc.GetClass("MPSNNGramMatrixCalculationGradient")}
	})
	return _MPSNNGramMatrixCalculationGradientClass
}

// GetMPSNNGramMatrixCalculationGradientClass returns the class object for MPSNNGramMatrixCalculationGradient.
func GetMPSNNGramMatrixCalculationGradientClass() MPSNNGramMatrixCalculationGradientClass {
	return getMPSNNGramMatrixCalculationGradientClass()
}

type MPSNNGramMatrixCalculationGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNGramMatrixCalculationGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGramMatrixCalculationGradientClass) Alloc() MPSNNGramMatrixCalculationGradient {
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNGramMatrixCalculationGradient.InitWithDeviceAlpha]
//
// # Instance Properties
//
//   - [MPSNNGramMatrixCalculationGradient.Alpha]
//   - [MPSNNGramMatrixCalculationGradient.SetAlpha]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient
type MPSNNGramMatrixCalculationGradient struct {
	MPSCNNGradientKernel
}

// MPSNNGramMatrixCalculationGradientFromID constructs a [MPSNNGramMatrixCalculationGradient] from an objc.ID.
func MPSNNGramMatrixCalculationGradientFromID(id objc.ID) MPSNNGramMatrixCalculationGradient {
	return MPSNNGramMatrixCalculationGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSNNGramMatrixCalculationGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNGramMatrixCalculationGradient] class.
//
// # Initializers
//
//   - [IMPSNNGramMatrixCalculationGradient.InitWithDeviceAlpha]
//
// # Instance Properties
//
//   - [IMPSNNGramMatrixCalculationGradient.Alpha]
//   - [IMPSNNGramMatrixCalculationGradient.SetAlpha]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient
type IMPSNNGramMatrixCalculationGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceAlpha(device metal.MTLDevice, alpha float32) MPSNNGramMatrixCalculationGradient

	// Topic: Instance Properties

	Alpha() float32
	SetAlpha(value float32)
}

// Init initializes the instance.
func (g MPSNNGramMatrixCalculationGradient) Init() MPSNNGramMatrixCalculationGradient {
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGramMatrixCalculationGradient) Autorelease() MPSNNGramMatrixCalculationGradient {
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGramMatrixCalculationGradient creates a new MPSNNGramMatrixCalculationGradient instance.
func NewMPSNNGramMatrixCalculationGradient() MPSNNGramMatrixCalculationGradient {
	class := getMPSNNGramMatrixCalculationGradientClass()
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewGramMatrixCalculationGradientWithCoder(aDecoder foundation.INSCoder) MPSNNGramMatrixCalculationGradient {
	instance := getMPSNNGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNGramMatrixCalculationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient/init(coder:device:)
func NewGramMatrixCalculationGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNGramMatrixCalculationGradient {
	instance := getMPSNNGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNGramMatrixCalculationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient/init(device:)
func NewGramMatrixCalculationGradientWithDevice(device metal.MTLDevice) MPSNNGramMatrixCalculationGradient {
	instance := getMPSNNGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNGramMatrixCalculationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient/init(device:alpha:)
func NewGramMatrixCalculationGradientWithDeviceAlpha(device metal.MTLDevice, alpha float32) MPSNNGramMatrixCalculationGradient {
	instance := getMPSNNGramMatrixCalculationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:alpha:"), device, alpha)
	return MPSNNGramMatrixCalculationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient/init(device:alpha:)
func (g MPSNNGramMatrixCalculationGradient) InitWithDeviceAlpha(device metal.MTLDevice, alpha float32) MPSNNGramMatrixCalculationGradient {
	rv := objc.Send[MPSNNGramMatrixCalculationGradient](g.ID, objc.Sel("initWithDevice:alpha:"), device, alpha)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradient/alpha
func (g MPSNNGramMatrixCalculationGradient) Alpha() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("alpha"))
	return rv
}
func (g MPSNNGramMatrixCalculationGradient) SetAlpha(value float32) {
	objc.Send[struct{}](g.ID, objc.Sel("setAlpha:"), value)
}

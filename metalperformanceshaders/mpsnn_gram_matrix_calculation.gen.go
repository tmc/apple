// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNGramMatrixCalculation] class.
var (
	_MPSNNGramMatrixCalculationClass     MPSNNGramMatrixCalculationClass
	_MPSNNGramMatrixCalculationClassOnce sync.Once
)

func getMPSNNGramMatrixCalculationClass() MPSNNGramMatrixCalculationClass {
	_MPSNNGramMatrixCalculationClassOnce.Do(func() {
		_MPSNNGramMatrixCalculationClass = MPSNNGramMatrixCalculationClass{class: objc.GetClass("MPSNNGramMatrixCalculation")}
	})
	return _MPSNNGramMatrixCalculationClass
}

// GetMPSNNGramMatrixCalculationClass returns the class object for MPSNNGramMatrixCalculation.
func GetMPSNNGramMatrixCalculationClass() MPSNNGramMatrixCalculationClass {
	return getMPSNNGramMatrixCalculationClass()
}

type MPSNNGramMatrixCalculationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNGramMatrixCalculationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGramMatrixCalculationClass) Alloc() MPSNNGramMatrixCalculation {
	rv := objc.Send[MPSNNGramMatrixCalculation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNGramMatrixCalculation.InitWithDeviceAlpha]
//
// # Instance Properties
//
//   - [MPSNNGramMatrixCalculation.Alpha]
//   - [MPSNNGramMatrixCalculation.SetAlpha]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculation
type MPSNNGramMatrixCalculation struct {
	MPSCNNKernel
}

// MPSNNGramMatrixCalculationFromID constructs a [MPSNNGramMatrixCalculation] from an objc.ID.
func MPSNNGramMatrixCalculationFromID(id objc.ID) MPSNNGramMatrixCalculation {
	return MPSNNGramMatrixCalculation{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSNNGramMatrixCalculation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNGramMatrixCalculation] class.
//
// # Initializers
//
//   - [IMPSNNGramMatrixCalculation.InitWithDeviceAlpha]
//
// # Instance Properties
//
//   - [IMPSNNGramMatrixCalculation.Alpha]
//   - [IMPSNNGramMatrixCalculation.SetAlpha]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculation
type IMPSNNGramMatrixCalculation interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceAlpha(device metal.MTLDevice, alpha float32) MPSNNGramMatrixCalculation

	// Topic: Instance Properties

	Alpha() float32
	SetAlpha(value float32)
}

// Init initializes the instance.
func (g MPSNNGramMatrixCalculation) Init() MPSNNGramMatrixCalculation {
	rv := objc.Send[MPSNNGramMatrixCalculation](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGramMatrixCalculation) Autorelease() MPSNNGramMatrixCalculation {
	rv := objc.Send[MPSNNGramMatrixCalculation](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGramMatrixCalculation creates a new MPSNNGramMatrixCalculation instance.
func NewMPSNNGramMatrixCalculation() MPSNNGramMatrixCalculation {
	class := getMPSNNGramMatrixCalculationClass()
	rv := objc.Send[MPSNNGramMatrixCalculation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewGramMatrixCalculationWithCoder(aDecoder foundation.INSCoder) MPSNNGramMatrixCalculation {
	instance := getMPSNNGramMatrixCalculationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNGramMatrixCalculationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculation/init(coder:device:)
func NewGramMatrixCalculationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNGramMatrixCalculation {
	instance := getMPSNNGramMatrixCalculationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNGramMatrixCalculationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculation/init(device:)
func NewGramMatrixCalculationWithDevice(device metal.MTLDevice) MPSNNGramMatrixCalculation {
	instance := getMPSNNGramMatrixCalculationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNGramMatrixCalculationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculation/init(device:alpha:)
func NewGramMatrixCalculationWithDeviceAlpha(device metal.MTLDevice, alpha float32) MPSNNGramMatrixCalculation {
	instance := getMPSNNGramMatrixCalculationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:alpha:"), device, alpha)
	return MPSNNGramMatrixCalculationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculation/init(device:alpha:)
func (g MPSNNGramMatrixCalculation) InitWithDeviceAlpha(device metal.MTLDevice, alpha float32) MPSNNGramMatrixCalculation {
	rv := objc.Send[MPSNNGramMatrixCalculation](g.ID, objc.Sel("initWithDevice:alpha:"), device, alpha)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculation/alpha
func (g MPSNNGramMatrixCalculation) Alpha() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("alpha"))
	return rv
}
func (g MPSNNGramMatrixCalculation) SetAlpha(value float32) {
	objc.Send[struct{}](g.ID, objc.Sel("setAlpha:"), value)
}

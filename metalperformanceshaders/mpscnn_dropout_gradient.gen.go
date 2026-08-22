// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNDropoutGradient] class.
var (
	_MPSCNNDropoutGradientClass     MPSCNNDropoutGradientClass
	_MPSCNNDropoutGradientClassOnce sync.Once
)

func getMPSCNNDropoutGradientClass() MPSCNNDropoutGradientClass {
	_MPSCNNDropoutGradientClassOnce.Do(func() {
		_MPSCNNDropoutGradientClass = MPSCNNDropoutGradientClass{class: objc.GetClass("MPSCNNDropoutGradient")}
	})
	return _MPSCNNDropoutGradientClass
}

// GetMPSCNNDropoutGradientClass returns the class object for MPSCNNDropoutGradient.
func GetMPSCNNDropoutGradientClass() MPSCNNDropoutGradientClass {
	return getMPSCNNDropoutGradientClass()
}

type MPSCNNDropoutGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDropoutGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDropoutGradientClass) Alloc() MPSCNNDropoutGradient {
	rv := objc.Send[MPSCNNDropoutGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient dropout filter.
//
// # Initializers
//
//   - [MPSCNNDropoutGradient.InitWithDeviceKeepProbabilitySeedMaskStrideInPixels]
//
// # Instance Properties
//
//   - [MPSCNNDropoutGradient.KeepProbability]
//   - [MPSCNNDropoutGradient.MaskStrideInPixels]
//   - [MPSCNNDropoutGradient.Seed]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradient
type MPSCNNDropoutGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNDropoutGradientFromID constructs a [MPSCNNDropoutGradient] from an objc.ID.
//
// A gradient dropout filter.
func MPSCNNDropoutGradientFromID(id objc.ID) MPSCNNDropoutGradient {
	return MPSCNNDropoutGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNDropoutGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDropoutGradient] class.
//
// # Initializers
//
//   - [IMPSCNNDropoutGradient.InitWithDeviceKeepProbabilitySeedMaskStrideInPixels]
//
// # Instance Properties
//
//   - [IMPSCNNDropoutGradient.KeepProbability]
//   - [IMPSCNNDropoutGradient.MaskStrideInPixels]
//   - [IMPSCNNDropoutGradient.Seed]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradient
type IMPSCNNDropoutGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceKeepProbabilitySeedMaskStrideInPixels(device metal.MTLDevice, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutGradient

	// Topic: Instance Properties

	KeepProbability() float32
	MaskStrideInPixels() metal.MTLSize
	Seed() uint
}

// Init initializes the instance.
func (c MPSCNNDropoutGradient) Init() MPSCNNDropoutGradient {
	rv := objc.Send[MPSCNNDropoutGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDropoutGradient) Autorelease() MPSCNNDropoutGradient {
	rv := objc.Send[MPSCNNDropoutGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDropoutGradient creates a new MPSCNNDropoutGradient instance.
func NewMPSCNNDropoutGradient() MPSCNNDropoutGradient {
	class := getMPSCNNDropoutGradientClass()
	rv := objc.Send[MPSCNNDropoutGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNDropoutGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNDropoutGradient {
	instance := getMPSCNNDropoutGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNDropoutGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradient/init(coder:device:)
func NewCNNDropoutGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNDropoutGradient {
	instance := getMPSCNNDropoutGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNDropoutGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNDropoutGradientWithDevice(device metal.MTLDevice) MPSCNNDropoutGradient {
	instance := getMPSCNNDropoutGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNDropoutGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradient/init(device:keepProbability:seed:maskStrideInPixels:)
func NewCNNDropoutGradientWithDeviceKeepProbabilitySeedMaskStrideInPixels(device metal.MTLDevice, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutGradient {
	instance := getMPSCNNDropoutGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:keepProbability:seed:maskStrideInPixels:"), device, keepProbability, seed, maskStrideInPixels)
	return MPSCNNDropoutGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradient/init(device:keepProbability:seed:maskStrideInPixels:)
func (c MPSCNNDropoutGradient) InitWithDeviceKeepProbabilitySeedMaskStrideInPixels(device metal.MTLDevice, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropoutGradient {
	rv := objc.Send[MPSCNNDropoutGradient](c.ID, objc.Sel("initWithDevice:keepProbability:seed:maskStrideInPixels:"), device, keepProbability, seed, maskStrideInPixels)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradient/keepProbability
func (c MPSCNNDropoutGradient) KeepProbability() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("keepProbability"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradient/maskStrideInPixels
func (c MPSCNNDropoutGradient) MaskStrideInPixels() metal.MTLSize {
	rv := objc.Send[metal.MTLSize](c.ID, objc.Sel("maskStrideInPixels"))
	return metal.MTLSize(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradient/seed
func (c MPSCNNDropoutGradient) Seed() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("seed"))
	return rv
}

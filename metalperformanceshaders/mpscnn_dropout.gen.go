// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNDropout] class.
var (
	_MPSCNNDropoutClass     MPSCNNDropoutClass
	_MPSCNNDropoutClassOnce sync.Once
)

func getMPSCNNDropoutClass() MPSCNNDropoutClass {
	_MPSCNNDropoutClassOnce.Do(func() {
		_MPSCNNDropoutClass = MPSCNNDropoutClass{class: objc.GetClass("MPSCNNDropout")}
	})
	return _MPSCNNDropoutClass
}

// GetMPSCNNDropoutClass returns the class object for MPSCNNDropout.
func GetMPSCNNDropoutClass() MPSCNNDropoutClass {
	return getMPSCNNDropoutClass()
}

type MPSCNNDropoutClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDropoutClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDropoutClass) Alloc() MPSCNNDropout {
	rv := objc.Send[MPSCNNDropout](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A dropout filter.
//
// # Initializers
//
//   - [MPSCNNDropout.InitWithDeviceKeepProbabilitySeedMaskStrideInPixels]
//
// # Instance Properties
//
//   - [MPSCNNDropout.KeepProbability]
//   - [MPSCNNDropout.MaskStrideInPixels]
//   - [MPSCNNDropout.Seed]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropout
type MPSCNNDropout struct {
	MPSCNNKernel
}

// MPSCNNDropoutFromID constructs a [MPSCNNDropout] from an objc.ID.
//
// A dropout filter.
func MPSCNNDropoutFromID(id objc.ID) MPSCNNDropout {
	return MPSCNNDropout{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNDropout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDropout] class.
//
// # Initializers
//
//   - [IMPSCNNDropout.InitWithDeviceKeepProbabilitySeedMaskStrideInPixels]
//
// # Instance Properties
//
//   - [IMPSCNNDropout.KeepProbability]
//   - [IMPSCNNDropout.MaskStrideInPixels]
//   - [IMPSCNNDropout.Seed]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropout
type IMPSCNNDropout interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceKeepProbabilitySeedMaskStrideInPixels(device metal.MTLDevice, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropout

	// Topic: Instance Properties

	KeepProbability() float32
	MaskStrideInPixels() metal.MTLSize
	Seed() uint
}

// Init initializes the instance.
func (c MPSCNNDropout) Init() MPSCNNDropout {
	rv := objc.Send[MPSCNNDropout](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDropout) Autorelease() MPSCNNDropout {
	rv := objc.Send[MPSCNNDropout](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDropout creates a new MPSCNNDropout instance.
func NewMPSCNNDropout() MPSCNNDropout {
	class := getMPSCNNDropoutClass()
	rv := objc.Send[MPSCNNDropout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNDropoutWithCoder(aDecoder foundation.INSCoder) MPSCNNDropout {
	instance := getMPSCNNDropoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNDropoutFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropout/init(coder:device:)
func NewCNNDropoutWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNDropout {
	instance := getMPSCNNDropoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNDropoutFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNDropoutWithDevice(device metal.MTLDevice) MPSCNNDropout {
	instance := getMPSCNNDropoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNDropoutFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropout/init(device:keepProbability:seed:maskStrideInPixels:)
func NewCNNDropoutWithDeviceKeepProbabilitySeedMaskStrideInPixels(device metal.MTLDevice, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropout {
	instance := getMPSCNNDropoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:keepProbability:seed:maskStrideInPixels:"), device, keepProbability, seed, maskStrideInPixels)
	return MPSCNNDropoutFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropout/init(device:keepProbability:seed:maskStrideInPixels:)
func (c MPSCNNDropout) InitWithDeviceKeepProbabilitySeedMaskStrideInPixels(device metal.MTLDevice, keepProbability float32, seed uint, maskStrideInPixels metal.MTLSize) MPSCNNDropout {
	rv := objc.Send[MPSCNNDropout](c.ID, objc.Sel("initWithDevice:keepProbability:seed:maskStrideInPixels:"), device, keepProbability, seed, maskStrideInPixels)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropout/keepProbability
func (c MPSCNNDropout) KeepProbability() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("keepProbability"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropout/maskStrideInPixels
func (c MPSCNNDropout) MaskStrideInPixels() metal.MTLSize {
	rv := objc.Send[metal.MTLSize](c.ID, objc.Sel("maskStrideInPixels"))
	return metal.MTLSize(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropout/seed
func (c MPSCNNDropout) Seed() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("seed"))
	return rv
}

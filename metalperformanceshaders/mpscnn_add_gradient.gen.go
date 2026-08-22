// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNAddGradient] class.
var (
	_MPSCNNAddGradientClass     MPSCNNAddGradientClass
	_MPSCNNAddGradientClassOnce sync.Once
)

func getMPSCNNAddGradientClass() MPSCNNAddGradientClass {
	_MPSCNNAddGradientClassOnce.Do(func() {
		_MPSCNNAddGradientClass = MPSCNNAddGradientClass{class: objc.GetClass("MPSCNNAddGradient")}
	})
	return _MPSCNNAddGradientClass
}

// GetMPSCNNAddGradientClass returns the class object for MPSCNNAddGradient.
func GetMPSCNNAddGradientClass() MPSCNNAddGradientClass {
	return getMPSCNNAddGradientClass()
}

type MPSCNNAddGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNAddGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNAddGradientClass) Alloc() MPSCNNAddGradient {
	rv := objc.Send[MPSCNNAddGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient addition operator.
//
// # Initializers
//
//   - [MPSCNNAddGradient.InitWithDeviceIsSecondarySourceFilter]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNAddGradient
type MPSCNNAddGradient struct {
	MPSCNNArithmeticGradient
}

// MPSCNNAddGradientFromID constructs a [MPSCNNAddGradient] from an objc.ID.
//
// A gradient addition operator.
func MPSCNNAddGradientFromID(id objc.ID) MPSCNNAddGradient {
	return MPSCNNAddGradient{MPSCNNArithmeticGradient: MPSCNNArithmeticGradientFromID(id)}
}

// NOTE: MPSCNNAddGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNAddGradient] class.
//
// # Initializers
//
//   - [IMPSCNNAddGradient.InitWithDeviceIsSecondarySourceFilter]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNAddGradient
type IMPSCNNAddGradient interface {
	IMPSCNNArithmeticGradient

	// Topic: Initializers

	InitWithDeviceIsSecondarySourceFilter(device metal.MTLDevice, isSecondarySourceFilter bool) MPSCNNAddGradient
}

// Init initializes the instance.
func (c MPSCNNAddGradient) Init() MPSCNNAddGradient {
	rv := objc.Send[MPSCNNAddGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNAddGradient) Autorelease() MPSCNNAddGradient {
	rv := objc.Send[MPSCNNAddGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNAddGradient creates a new MPSCNNAddGradient instance.
func NewMPSCNNAddGradient() MPSCNNAddGradient {
	class := getMPSCNNAddGradientClass()
	rv := objc.Send[MPSCNNAddGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNAddGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNAddGradient {
	instance := getMPSCNNAddGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNAddGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(coder:device:)
func NewCNNAddGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNAddGradient {
	instance := getMPSCNNAddGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNAddGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNAddGradientWithDevice(device metal.MTLDevice) MPSCNNAddGradient {
	instance := getMPSCNNAddGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNAddGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNAddGradient/init(device:isSecondarySourceFilter:)
func NewCNNAddGradientWithDeviceIsSecondarySourceFilter(device metal.MTLDevice, isSecondarySourceFilter bool) MPSCNNAddGradient {
	instance := getMPSCNNAddGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	return MPSCNNAddGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNAddGradient/init(device:isSecondarySourceFilter:)
func (c MPSCNNAddGradient) InitWithDeviceIsSecondarySourceFilter(device metal.MTLDevice, isSecondarySourceFilter bool) MPSCNNAddGradient {
	rv := objc.Send[MPSCNNAddGradient](c.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	return rv
}

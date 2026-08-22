// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSubtractGradient] class.
var (
	_MPSCNNSubtractGradientClass     MPSCNNSubtractGradientClass
	_MPSCNNSubtractGradientClassOnce sync.Once
)

func getMPSCNNSubtractGradientClass() MPSCNNSubtractGradientClass {
	_MPSCNNSubtractGradientClassOnce.Do(func() {
		_MPSCNNSubtractGradientClass = MPSCNNSubtractGradientClass{class: objc.GetClass("MPSCNNSubtractGradient")}
	})
	return _MPSCNNSubtractGradientClass
}

// GetMPSCNNSubtractGradientClass returns the class object for MPSCNNSubtractGradient.
func GetMPSCNNSubtractGradientClass() MPSCNNSubtractGradientClass {
	return getMPSCNNSubtractGradientClass()
}

type MPSCNNSubtractGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSubtractGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSubtractGradientClass) Alloc() MPSCNNSubtractGradient {
	rv := objc.Send[MPSCNNSubtractGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient subtraction operator.
//
// # Initializers
//
//   - [MPSCNNSubtractGradient.InitWithDeviceIsSecondarySourceFilter]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtractGradient
type MPSCNNSubtractGradient struct {
	MPSCNNArithmeticGradient
}

// MPSCNNSubtractGradientFromID constructs a [MPSCNNSubtractGradient] from an objc.ID.
//
// A gradient subtraction operator.
func MPSCNNSubtractGradientFromID(id objc.ID) MPSCNNSubtractGradient {
	return MPSCNNSubtractGradient{MPSCNNArithmeticGradient: MPSCNNArithmeticGradientFromID(id)}
}

// NOTE: MPSCNNSubtractGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSubtractGradient] class.
//
// # Initializers
//
//   - [IMPSCNNSubtractGradient.InitWithDeviceIsSecondarySourceFilter]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtractGradient
type IMPSCNNSubtractGradient interface {
	IMPSCNNArithmeticGradient

	// Topic: Initializers

	InitWithDeviceIsSecondarySourceFilter(device metal.MTLDevice, isSecondarySourceFilter bool) MPSCNNSubtractGradient
}

// Init initializes the instance.
func (c MPSCNNSubtractGradient) Init() MPSCNNSubtractGradient {
	rv := objc.Send[MPSCNNSubtractGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSubtractGradient) Autorelease() MPSCNNSubtractGradient {
	rv := objc.Send[MPSCNNSubtractGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSubtractGradient creates a new MPSCNNSubtractGradient instance.
func NewMPSCNNSubtractGradient() MPSCNNSubtractGradient {
	class := getMPSCNNSubtractGradientClass()
	rv := objc.Send[MPSCNNSubtractGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNSubtractGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNSubtractGradient {
	instance := getMPSCNNSubtractGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNSubtractGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(coder:device:)
func NewCNNSubtractGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNSubtractGradient {
	instance := getMPSCNNSubtractGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNSubtractGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNSubtractGradientWithDevice(device metal.MTLDevice) MPSCNNSubtractGradient {
	instance := getMPSCNNSubtractGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNSubtractGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtractGradient/init(device:isSecondarySourceFilter:)
func NewCNNSubtractGradientWithDeviceIsSecondarySourceFilter(device metal.MTLDevice, isSecondarySourceFilter bool) MPSCNNSubtractGradient {
	instance := getMPSCNNSubtractGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	return MPSCNNSubtractGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtractGradient/init(device:isSecondarySourceFilter:)
func (c MPSCNNSubtractGradient) InitWithDeviceIsSecondarySourceFilter(device metal.MTLDevice, isSecondarySourceFilter bool) MPSCNNSubtractGradient {
	rv := objc.Send[MPSCNNSubtractGradient](c.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	return rv
}

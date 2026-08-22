// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNMultiplyGradient] class.
var (
	_MPSCNNMultiplyGradientClass     MPSCNNMultiplyGradientClass
	_MPSCNNMultiplyGradientClassOnce sync.Once
)

func getMPSCNNMultiplyGradientClass() MPSCNNMultiplyGradientClass {
	_MPSCNNMultiplyGradientClassOnce.Do(func() {
		_MPSCNNMultiplyGradientClass = MPSCNNMultiplyGradientClass{class: objc.GetClass("MPSCNNMultiplyGradient")}
	})
	return _MPSCNNMultiplyGradientClass
}

// GetMPSCNNMultiplyGradientClass returns the class object for MPSCNNMultiplyGradient.
func GetMPSCNNMultiplyGradientClass() MPSCNNMultiplyGradientClass {
	return getMPSCNNMultiplyGradientClass()
}

type MPSCNNMultiplyGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNMultiplyGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNMultiplyGradientClass) Alloc() MPSCNNMultiplyGradient {
	rv := objc.Send[MPSCNNMultiplyGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient multiply operator.
//
// # Initializers
//
//   - [MPSCNNMultiplyGradient.InitWithDeviceIsSecondarySourceFilter]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiplyGradient
type MPSCNNMultiplyGradient struct {
	MPSCNNArithmeticGradient
}

// MPSCNNMultiplyGradientFromID constructs a [MPSCNNMultiplyGradient] from an objc.ID.
//
// A gradient multiply operator.
func MPSCNNMultiplyGradientFromID(id objc.ID) MPSCNNMultiplyGradient {
	return MPSCNNMultiplyGradient{MPSCNNArithmeticGradient: MPSCNNArithmeticGradientFromID(id)}
}

// NOTE: MPSCNNMultiplyGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNMultiplyGradient] class.
//
// # Initializers
//
//   - [IMPSCNNMultiplyGradient.InitWithDeviceIsSecondarySourceFilter]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiplyGradient
type IMPSCNNMultiplyGradient interface {
	IMPSCNNArithmeticGradient

	// Topic: Initializers

	InitWithDeviceIsSecondarySourceFilter(device metal.MTLDevice, isSecondarySourceFilter bool) MPSCNNMultiplyGradient
}

// Init initializes the instance.
func (c MPSCNNMultiplyGradient) Init() MPSCNNMultiplyGradient {
	rv := objc.Send[MPSCNNMultiplyGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNMultiplyGradient) Autorelease() MPSCNNMultiplyGradient {
	rv := objc.Send[MPSCNNMultiplyGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNMultiplyGradient creates a new MPSCNNMultiplyGradient instance.
func NewMPSCNNMultiplyGradient() MPSCNNMultiplyGradient {
	class := getMPSCNNMultiplyGradientClass()
	rv := objc.Send[MPSCNNMultiplyGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNMultiplyGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNMultiplyGradient {
	instance := getMPSCNNMultiplyGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNMultiplyGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(coder:device:)
func NewCNNMultiplyGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNMultiplyGradient {
	instance := getMPSCNNMultiplyGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNMultiplyGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNMultiplyGradientWithDevice(device metal.MTLDevice) MPSCNNMultiplyGradient {
	instance := getMPSCNNMultiplyGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNMultiplyGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiplyGradient/init(device:isSecondarySourceFilter:)
func NewCNNMultiplyGradientWithDeviceIsSecondarySourceFilter(device metal.MTLDevice, isSecondarySourceFilter bool) MPSCNNMultiplyGradient {
	instance := getMPSCNNMultiplyGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	return MPSCNNMultiplyGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiplyGradient/init(device:isSecondarySourceFilter:)
func (c MPSCNNMultiplyGradient) InitWithDeviceIsSecondarySourceFilter(device metal.MTLDevice, isSecondarySourceFilter bool) MPSCNNMultiplyGradient {
	rv := objc.Send[MPSCNNMultiplyGradient](c.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	return rv
}

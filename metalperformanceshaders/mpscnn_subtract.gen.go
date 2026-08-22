// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSubtract] class.
var (
	_MPSCNNSubtractClass     MPSCNNSubtractClass
	_MPSCNNSubtractClassOnce sync.Once
)

func getMPSCNNSubtractClass() MPSCNNSubtractClass {
	_MPSCNNSubtractClassOnce.Do(func() {
		_MPSCNNSubtractClass = MPSCNNSubtractClass{class: objc.GetClass("MPSCNNSubtract")}
	})
	return _MPSCNNSubtractClass
}

// GetMPSCNNSubtractClass returns the class object for MPSCNNSubtract.
func GetMPSCNNSubtractClass() MPSCNNSubtractClass {
	return getMPSCNNSubtractClass()
}

type MPSCNNSubtractClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSubtractClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSubtractClass) Alloc() MPSCNNSubtract {
	rv := objc.Send[MPSCNNSubtract](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A subtraction operator.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtract
type MPSCNNSubtract struct {
	MPSCNNArithmetic
}

// MPSCNNSubtractFromID constructs a [MPSCNNSubtract] from an objc.ID.
//
// A subtraction operator.
func MPSCNNSubtractFromID(id objc.ID) MPSCNNSubtract {
	return MPSCNNSubtract{MPSCNNArithmetic: MPSCNNArithmeticFromID(id)}
}

// NOTE: MPSCNNSubtract adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSubtract] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtract
type IMPSCNNSubtract interface {
	IMPSCNNArithmetic
}

// Init initializes the instance.
func (c MPSCNNSubtract) Init() MPSCNNSubtract {
	rv := objc.Send[MPSCNNSubtract](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSubtract) Autorelease() MPSCNNSubtract {
	rv := objc.Send[MPSCNNSubtract](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSubtract creates a new MPSCNNSubtract instance.
func NewMPSCNNSubtract() MPSCNNSubtract {
	class := getMPSCNNSubtractClass()
	rv := objc.Send[MPSCNNSubtract](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNSubtractWithCoder(aDecoder foundation.INSCoder) MPSCNNSubtract {
	instance := getMPSCNNSubtractClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNSubtractFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(coder:device:)
func NewCNNSubtractWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNSubtract {
	instance := getMPSCNNSubtractClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNSubtractFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtract/init(device:)
func NewCNNSubtractWithDevice(device metal.MTLDevice) MPSCNNSubtract {
	instance := getMPSCNNSubtractClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNSubtractFromID(rv)
}

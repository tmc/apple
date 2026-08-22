// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageSubtract] class.
var (
	_MPSImageSubtractClass     MPSImageSubtractClass
	_MPSImageSubtractClassOnce sync.Once
)

func getMPSImageSubtractClass() MPSImageSubtractClass {
	_MPSImageSubtractClassOnce.Do(func() {
		_MPSImageSubtractClass = MPSImageSubtractClass{class: objc.GetClass("MPSImageSubtract")}
	})
	return _MPSImageSubtractClass
}

// GetMPSImageSubtractClass returns the class object for MPSImageSubtract.
func GetMPSImageSubtractClass() MPSImageSubtractClass {
	return getMPSImageSubtractClass()
}

type MPSImageSubtractClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageSubtractClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageSubtractClass) Alloc() MPSImageSubtract {
	rv := objc.Send[MPSImageSubtract](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the element-wise difference of its two input images.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSubtract
type MPSImageSubtract struct {
	MPSImageArithmetic
}

// MPSImageSubtractFromID constructs a [MPSImageSubtract] from an objc.ID.
//
// A filter that returns the element-wise difference of its two input images.
func MPSImageSubtractFromID(id objc.ID) MPSImageSubtract {
	return MPSImageSubtract{MPSImageArithmetic: MPSImageArithmeticFromID(id)}
}

// NOTE: MPSImageSubtract adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageSubtract] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSubtract
type IMPSImageSubtract interface {
	IMPSImageArithmetic
}

// Init initializes the instance.
func (i MPSImageSubtract) Init() MPSImageSubtract {
	rv := objc.Send[MPSImageSubtract](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageSubtract) Autorelease() MPSImageSubtract {
	rv := objc.Send[MPSImageSubtract](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageSubtract creates a new MPSImageSubtract instance.
func NewMPSImageSubtract() MPSImageSubtract {
	class := getMPSImageSubtractClass()
	rv := objc.Send[MPSImageSubtract](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageSubtractWithCoder(aDecoder foundation.INSCoder) MPSImageSubtract {
	instance := getMPSImageSubtractClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageSubtractFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/init(coder:device:)
func NewImageSubtractWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageSubtract {
	instance := getMPSImageSubtractClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageSubtractFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSubtract/init(device:)
func NewImageSubtractWithDevice(device metal.MTLDevice) MPSImageSubtract {
	instance := getMPSImageSubtractClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageSubtractFromID(rv)
}

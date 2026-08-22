// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageMultiply] class.
var (
	_MPSImageMultiplyClass     MPSImageMultiplyClass
	_MPSImageMultiplyClassOnce sync.Once
)

func getMPSImageMultiplyClass() MPSImageMultiplyClass {
	_MPSImageMultiplyClassOnce.Do(func() {
		_MPSImageMultiplyClass = MPSImageMultiplyClass{class: objc.GetClass("MPSImageMultiply")}
	})
	return _MPSImageMultiplyClass
}

// GetMPSImageMultiplyClass returns the class object for MPSImageMultiply.
func GetMPSImageMultiplyClass() MPSImageMultiplyClass {
	return getMPSImageMultiplyClass()
}

type MPSImageMultiplyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageMultiplyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageMultiplyClass) Alloc() MPSImageMultiply {
	rv := objc.Send[MPSImageMultiply](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the element-wise product of its two input images.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMultiply
type MPSImageMultiply struct {
	MPSImageArithmetic
}

// MPSImageMultiplyFromID constructs a [MPSImageMultiply] from an objc.ID.
//
// A filter that returns the element-wise product of its two input images.
func MPSImageMultiplyFromID(id objc.ID) MPSImageMultiply {
	return MPSImageMultiply{MPSImageArithmetic: MPSImageArithmeticFromID(id)}
}

// NOTE: MPSImageMultiply adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageMultiply] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMultiply
type IMPSImageMultiply interface {
	IMPSImageArithmetic
}

// Init initializes the instance.
func (i MPSImageMultiply) Init() MPSImageMultiply {
	rv := objc.Send[MPSImageMultiply](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageMultiply) Autorelease() MPSImageMultiply {
	rv := objc.Send[MPSImageMultiply](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageMultiply creates a new MPSImageMultiply instance.
func NewMPSImageMultiply() MPSImageMultiply {
	class := getMPSImageMultiplyClass()
	rv := objc.Send[MPSImageMultiply](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageMultiplyWithCoder(aDecoder foundation.INSCoder) MPSImageMultiply {
	instance := getMPSImageMultiplyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageMultiplyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/init(coder:device:)
func NewImageMultiplyWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageMultiply {
	instance := getMPSImageMultiplyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageMultiplyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMultiply/init(device:)
func NewImageMultiplyWithDevice(device metal.MTLDevice) MPSImageMultiply {
	instance := getMPSImageMultiplyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageMultiplyFromID(rv)
}

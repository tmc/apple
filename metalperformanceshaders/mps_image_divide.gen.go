// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageDivide] class.
var (
	_MPSImageDivideClass     MPSImageDivideClass
	_MPSImageDivideClassOnce sync.Once
)

func getMPSImageDivideClass() MPSImageDivideClass {
	_MPSImageDivideClassOnce.Do(func() {
		_MPSImageDivideClass = MPSImageDivideClass{class: objc.GetClass("MPSImageDivide")}
	})
	return _MPSImageDivideClass
}

// GetMPSImageDivideClass returns the class object for MPSImageDivide.
func GetMPSImageDivideClass() MPSImageDivideClass {
	return getMPSImageDivideClass()
}

type MPSImageDivideClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageDivideClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageDivideClass) Alloc() MPSImageDivide {
	rv := objc.Send[MPSImageDivide](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the element-wise quotient of its two input images.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDivide
type MPSImageDivide struct {
	MPSImageArithmetic
}

// MPSImageDivideFromID constructs a [MPSImageDivide] from an objc.ID.
//
// A filter that returns the element-wise quotient of its two input images.
func MPSImageDivideFromID(id objc.ID) MPSImageDivide {
	return MPSImageDivide{MPSImageArithmetic: MPSImageArithmeticFromID(id)}
}

// NOTE: MPSImageDivide adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageDivide] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDivide
type IMPSImageDivide interface {
	IMPSImageArithmetic
}

// Init initializes the instance.
func (i MPSImageDivide) Init() MPSImageDivide {
	rv := objc.Send[MPSImageDivide](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageDivide) Autorelease() MPSImageDivide {
	rv := objc.Send[MPSImageDivide](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageDivide creates a new MPSImageDivide instance.
func NewMPSImageDivide() MPSImageDivide {
	class := getMPSImageDivideClass()
	rv := objc.Send[MPSImageDivide](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageDivideWithCoder(aDecoder foundation.INSCoder) MPSImageDivide {
	instance := getMPSImageDivideClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageDivideFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/init(coder:device:)
func NewImageDivideWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageDivide {
	instance := getMPSImageDivideClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageDivideFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDivide/init(device:)
func NewImageDivideWithDevice(device metal.MTLDevice) MPSImageDivide {
	instance := getMPSImageDivideClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageDivideFromID(rv)
}

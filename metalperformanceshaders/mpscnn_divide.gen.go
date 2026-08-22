// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNDivide] class.
var (
	_MPSCNNDivideClass     MPSCNNDivideClass
	_MPSCNNDivideClassOnce sync.Once
)

func getMPSCNNDivideClass() MPSCNNDivideClass {
	_MPSCNNDivideClassOnce.Do(func() {
		_MPSCNNDivideClass = MPSCNNDivideClass{class: objc.GetClass("MPSCNNDivide")}
	})
	return _MPSCNNDivideClass
}

// GetMPSCNNDivideClass returns the class object for MPSCNNDivide.
func GetMPSCNNDivideClass() MPSCNNDivideClass {
	return getMPSCNNDivideClass()
}

type MPSCNNDivideClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNDivideClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNDivideClass) Alloc() MPSCNNDivide {
	rv := objc.Send[MPSCNNDivide](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A division operator.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDivide
type MPSCNNDivide struct {
	MPSCNNArithmetic
}

// MPSCNNDivideFromID constructs a [MPSCNNDivide] from an objc.ID.
//
// A division operator.
func MPSCNNDivideFromID(id objc.ID) MPSCNNDivide {
	return MPSCNNDivide{MPSCNNArithmetic: MPSCNNArithmeticFromID(id)}
}

// NOTE: MPSCNNDivide adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNDivide] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDivide
type IMPSCNNDivide interface {
	IMPSCNNArithmetic
}

// Init initializes the instance.
func (c MPSCNNDivide) Init() MPSCNNDivide {
	rv := objc.Send[MPSCNNDivide](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNDivide) Autorelease() MPSCNNDivide {
	rv := objc.Send[MPSCNNDivide](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNDivide creates a new MPSCNNDivide instance.
func NewMPSCNNDivide() MPSCNNDivide {
	class := getMPSCNNDivideClass()
	rv := objc.Send[MPSCNNDivide](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNDivideWithCoder(aDecoder foundation.INSCoder) MPSCNNDivide {
	instance := getMPSCNNDivideClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNDivideFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(coder:device:)
func NewCNNDivideWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNDivide {
	instance := getMPSCNNDivideClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNDivideFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDivide/init(device:)
func NewCNNDivideWithDevice(device metal.MTLDevice) MPSCNNDivide {
	instance := getMPSCNNDivideClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNDivideFromID(rv)
}

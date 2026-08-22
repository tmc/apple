// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNMultiply] class.
var (
	_MPSCNNMultiplyClass     MPSCNNMultiplyClass
	_MPSCNNMultiplyClassOnce sync.Once
)

func getMPSCNNMultiplyClass() MPSCNNMultiplyClass {
	_MPSCNNMultiplyClassOnce.Do(func() {
		_MPSCNNMultiplyClass = MPSCNNMultiplyClass{class: objc.GetClass("MPSCNNMultiply")}
	})
	return _MPSCNNMultiplyClass
}

// GetMPSCNNMultiplyClass returns the class object for MPSCNNMultiply.
func GetMPSCNNMultiplyClass() MPSCNNMultiplyClass {
	return getMPSCNNMultiplyClass()
}

type MPSCNNMultiplyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNMultiplyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNMultiplyClass) Alloc() MPSCNNMultiply {
	rv := objc.Send[MPSCNNMultiply](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A multiply operator.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiply
type MPSCNNMultiply struct {
	MPSCNNArithmetic
}

// MPSCNNMultiplyFromID constructs a [MPSCNNMultiply] from an objc.ID.
//
// A multiply operator.
func MPSCNNMultiplyFromID(id objc.ID) MPSCNNMultiply {
	return MPSCNNMultiply{MPSCNNArithmetic: MPSCNNArithmeticFromID(id)}
}

// NOTE: MPSCNNMultiply adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNMultiply] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiply
type IMPSCNNMultiply interface {
	IMPSCNNArithmetic
}

// Init initializes the instance.
func (c MPSCNNMultiply) Init() MPSCNNMultiply {
	rv := objc.Send[MPSCNNMultiply](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNMultiply) Autorelease() MPSCNNMultiply {
	rv := objc.Send[MPSCNNMultiply](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNMultiply creates a new MPSCNNMultiply instance.
func NewMPSCNNMultiply() MPSCNNMultiply {
	class := getMPSCNNMultiplyClass()
	rv := objc.Send[MPSCNNMultiply](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNMultiplyWithCoder(aDecoder foundation.INSCoder) MPSCNNMultiply {
	instance := getMPSCNNMultiplyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNMultiplyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(coder:device:)
func NewCNNMultiplyWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNMultiply {
	instance := getMPSCNNMultiplyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNMultiplyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiply/init(device:)
func NewCNNMultiplyWithDevice(device metal.MTLDevice) MPSCNNMultiply {
	instance := getMPSCNNMultiplyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNMultiplyFromID(rv)
}

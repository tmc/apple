// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceRowMax] class.
var (
	_MPSNNReduceRowMaxClass     MPSNNReduceRowMaxClass
	_MPSNNReduceRowMaxClassOnce sync.Once
)

func getMPSNNReduceRowMaxClass() MPSNNReduceRowMaxClass {
	_MPSNNReduceRowMaxClassOnce.Do(func() {
		_MPSNNReduceRowMaxClass = MPSNNReduceRowMaxClass{class: objc.GetClass("MPSNNReduceRowMax")}
	})
	return _MPSNNReduceRowMaxClass
}

// GetMPSNNReduceRowMaxClass returns the class object for MPSNNReduceRowMax.
func GetMPSNNReduceRowMaxClass() MPSNNReduceRowMaxClass {
	return getMPSNNReduceRowMaxClass()
}

type MPSNNReduceRowMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceRowMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceRowMaxClass) Alloc() MPSNNReduceRowMax {
	rv := objc.Send[MPSNNReduceRowMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the maximum value for each row in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMax
type MPSNNReduceRowMax struct {
	MPSNNReduceUnary
}

// MPSNNReduceRowMaxFromID constructs a [MPSNNReduceRowMax] from an objc.ID.
//
// A reduction filter that returns the maximum value for each row in an image.
func MPSNNReduceRowMaxFromID(id objc.ID) MPSNNReduceRowMax {
	return MPSNNReduceRowMax{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceRowMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceRowMax] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMax
type IMPSNNReduceRowMax interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceRowMax) Init() MPSNNReduceRowMax {
	rv := objc.Send[MPSNNReduceRowMax](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceRowMax) Autorelease() MPSNNReduceRowMax {
	rv := objc.Send[MPSNNReduceRowMax](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceRowMax creates a new MPSNNReduceRowMax instance.
func NewMPSNNReduceRowMax() MPSNNReduceRowMax {
	class := getMPSNNReduceRowMaxClass()
	rv := objc.Send[MPSNNReduceRowMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceRowMaxWithCoder(aDecoder foundation.INSCoder) MPSNNReduceRowMax {
	instance := getMPSNNReduceRowMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceRowMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMax/init(coder:device:)
func NewReduceRowMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceRowMax {
	instance := getMPSNNReduceRowMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceRowMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMax/init(device:)
func NewReduceRowMaxWithDevice(device metal.MTLDevice) MPSNNReduceRowMax {
	instance := getMPSNNReduceRowMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceRowMaxFromID(rv)
}

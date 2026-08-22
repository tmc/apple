// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceColumnMax] class.
var (
	_MPSNNReduceColumnMaxClass     MPSNNReduceColumnMaxClass
	_MPSNNReduceColumnMaxClassOnce sync.Once
)

func getMPSNNReduceColumnMaxClass() MPSNNReduceColumnMaxClass {
	_MPSNNReduceColumnMaxClassOnce.Do(func() {
		_MPSNNReduceColumnMaxClass = MPSNNReduceColumnMaxClass{class: objc.GetClass("MPSNNReduceColumnMax")}
	})
	return _MPSNNReduceColumnMaxClass
}

// GetMPSNNReduceColumnMaxClass returns the class object for MPSNNReduceColumnMax.
func GetMPSNNReduceColumnMaxClass() MPSNNReduceColumnMaxClass {
	return getMPSNNReduceColumnMaxClass()
}

type MPSNNReduceColumnMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceColumnMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceColumnMaxClass) Alloc() MPSNNReduceColumnMax {
	rv := objc.Send[MPSNNReduceColumnMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the maximum value for each column in an
// image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMax
type MPSNNReduceColumnMax struct {
	MPSNNReduceUnary
}

// MPSNNReduceColumnMaxFromID constructs a [MPSNNReduceColumnMax] from an objc.ID.
//
// A reduction filter that returns the maximum value for each column in an
// image.
func MPSNNReduceColumnMaxFromID(id objc.ID) MPSNNReduceColumnMax {
	return MPSNNReduceColumnMax{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceColumnMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceColumnMax] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMax
type IMPSNNReduceColumnMax interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceColumnMax) Init() MPSNNReduceColumnMax {
	rv := objc.Send[MPSNNReduceColumnMax](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceColumnMax) Autorelease() MPSNNReduceColumnMax {
	rv := objc.Send[MPSNNReduceColumnMax](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceColumnMax creates a new MPSNNReduceColumnMax instance.
func NewMPSNNReduceColumnMax() MPSNNReduceColumnMax {
	class := getMPSNNReduceColumnMaxClass()
	rv := objc.Send[MPSNNReduceColumnMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceColumnMaxWithCoder(aDecoder foundation.INSCoder) MPSNNReduceColumnMax {
	instance := getMPSNNReduceColumnMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceColumnMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMax/init(coder:device:)
func NewReduceColumnMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceColumnMax {
	instance := getMPSNNReduceColumnMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceColumnMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMax/init(device:)
func NewReduceColumnMaxWithDevice(device metal.MTLDevice) MPSNNReduceColumnMax {
	instance := getMPSNNReduceColumnMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceColumnMaxFromID(rv)
}

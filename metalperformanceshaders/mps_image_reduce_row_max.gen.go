// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageReduceRowMax] class.
var (
	_MPSImageReduceRowMaxClass     MPSImageReduceRowMaxClass
	_MPSImageReduceRowMaxClassOnce sync.Once
)

func getMPSImageReduceRowMaxClass() MPSImageReduceRowMaxClass {
	_MPSImageReduceRowMaxClassOnce.Do(func() {
		_MPSImageReduceRowMaxClass = MPSImageReduceRowMaxClass{class: objc.GetClass("MPSImageReduceRowMax")}
	})
	return _MPSImageReduceRowMaxClass
}

// GetMPSImageReduceRowMaxClass returns the class object for MPSImageReduceRowMax.
func GetMPSImageReduceRowMaxClass() MPSImageReduceRowMaxClass {
	return getMPSImageReduceRowMaxClass()
}

type MPSImageReduceRowMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageReduceRowMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageReduceRowMaxClass) Alloc() MPSImageReduceRowMax {
	rv := objc.Send[MPSImageReduceRowMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the maximum value for each row in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMax
type MPSImageReduceRowMax struct {
	MPSImageReduceUnary
}

// MPSImageReduceRowMaxFromID constructs a [MPSImageReduceRowMax] from an objc.ID.
//
// A filter that returns the maximum value for each row in an image.
func MPSImageReduceRowMaxFromID(id objc.ID) MPSImageReduceRowMax {
	return MPSImageReduceRowMax{MPSImageReduceUnary: MPSImageReduceUnaryFromID(id)}
}

// NOTE: MPSImageReduceRowMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageReduceRowMax] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMax
type IMPSImageReduceRowMax interface {
	IMPSImageReduceUnary
}

// Init initializes the instance.
func (i MPSImageReduceRowMax) Init() MPSImageReduceRowMax {
	rv := objc.Send[MPSImageReduceRowMax](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageReduceRowMax) Autorelease() MPSImageReduceRowMax {
	rv := objc.Send[MPSImageReduceRowMax](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageReduceRowMax creates a new MPSImageReduceRowMax instance.
func NewMPSImageReduceRowMax() MPSImageReduceRowMax {
	class := getMPSImageReduceRowMaxClass()
	rv := objc.Send[MPSImageReduceRowMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageReduceRowMaxWithCoder(aDecoder foundation.INSCoder) MPSImageReduceRowMax {
	instance := getMPSImageReduceRowMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageReduceRowMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageReduceRowMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageReduceRowMax {
	instance := getMPSImageReduceRowMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageReduceRowMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMax/init(device:)
func NewImageReduceRowMaxWithDevice(device metal.MTLDevice) MPSImageReduceRowMax {
	instance := getMPSImageReduceRowMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageReduceRowMaxFromID(rv)
}

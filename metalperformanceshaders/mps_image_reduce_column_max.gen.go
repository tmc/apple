// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageReduceColumnMax] class.
var (
	_MPSImageReduceColumnMaxClass     MPSImageReduceColumnMaxClass
	_MPSImageReduceColumnMaxClassOnce sync.Once
)

func getMPSImageReduceColumnMaxClass() MPSImageReduceColumnMaxClass {
	_MPSImageReduceColumnMaxClassOnce.Do(func() {
		_MPSImageReduceColumnMaxClass = MPSImageReduceColumnMaxClass{class: objc.GetClass("MPSImageReduceColumnMax")}
	})
	return _MPSImageReduceColumnMaxClass
}

// GetMPSImageReduceColumnMaxClass returns the class object for MPSImageReduceColumnMax.
func GetMPSImageReduceColumnMaxClass() MPSImageReduceColumnMaxClass {
	return getMPSImageReduceColumnMaxClass()
}

type MPSImageReduceColumnMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageReduceColumnMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageReduceColumnMaxClass) Alloc() MPSImageReduceColumnMax {
	rv := objc.Send[MPSImageReduceColumnMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the maximum value for each column in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMax
type MPSImageReduceColumnMax struct {
	MPSImageReduceUnary
}

// MPSImageReduceColumnMaxFromID constructs a [MPSImageReduceColumnMax] from an objc.ID.
//
// A filter that returns the maximum value for each column in an image.
func MPSImageReduceColumnMaxFromID(id objc.ID) MPSImageReduceColumnMax {
	return MPSImageReduceColumnMax{MPSImageReduceUnary: MPSImageReduceUnaryFromID(id)}
}

// NOTE: MPSImageReduceColumnMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageReduceColumnMax] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMax
type IMPSImageReduceColumnMax interface {
	IMPSImageReduceUnary
}

// Init initializes the instance.
func (i MPSImageReduceColumnMax) Init() MPSImageReduceColumnMax {
	rv := objc.Send[MPSImageReduceColumnMax](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageReduceColumnMax) Autorelease() MPSImageReduceColumnMax {
	rv := objc.Send[MPSImageReduceColumnMax](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageReduceColumnMax creates a new MPSImageReduceColumnMax instance.
func NewMPSImageReduceColumnMax() MPSImageReduceColumnMax {
	class := getMPSImageReduceColumnMaxClass()
	rv := objc.Send[MPSImageReduceColumnMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageReduceColumnMaxWithCoder(aDecoder foundation.INSCoder) MPSImageReduceColumnMax {
	instance := getMPSImageReduceColumnMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageReduceColumnMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageReduceColumnMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageReduceColumnMax {
	instance := getMPSImageReduceColumnMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageReduceColumnMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMax/init(device:)
func NewImageReduceColumnMaxWithDevice(device metal.MTLDevice) MPSImageReduceColumnMax {
	instance := getMPSImageReduceColumnMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageReduceColumnMaxFromID(rv)
}

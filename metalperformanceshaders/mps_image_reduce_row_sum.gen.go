// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageReduceRowSum] class.
var (
	_MPSImageReduceRowSumClass     MPSImageReduceRowSumClass
	_MPSImageReduceRowSumClassOnce sync.Once
)

func getMPSImageReduceRowSumClass() MPSImageReduceRowSumClass {
	_MPSImageReduceRowSumClassOnce.Do(func() {
		_MPSImageReduceRowSumClass = MPSImageReduceRowSumClass{class: objc.GetClass("MPSImageReduceRowSum")}
	})
	return _MPSImageReduceRowSumClass
}

// GetMPSImageReduceRowSumClass returns the class object for MPSImageReduceRowSum.
func GetMPSImageReduceRowSumClass() MPSImageReduceRowSumClass {
	return getMPSImageReduceRowSumClass()
}

type MPSImageReduceRowSumClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageReduceRowSumClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageReduceRowSumClass) Alloc() MPSImageReduceRowSum {
	rv := objc.Send[MPSImageReduceRowSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the sum of all values for a row in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowSum
type MPSImageReduceRowSum struct {
	MPSImageReduceUnary
}

// MPSImageReduceRowSumFromID constructs a [MPSImageReduceRowSum] from an objc.ID.
//
// A filter that returns the sum of all values for a row in an image.
func MPSImageReduceRowSumFromID(id objc.ID) MPSImageReduceRowSum {
	return MPSImageReduceRowSum{MPSImageReduceUnary: MPSImageReduceUnaryFromID(id)}
}

// NOTE: MPSImageReduceRowSum adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageReduceRowSum] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowSum
type IMPSImageReduceRowSum interface {
	IMPSImageReduceUnary
}

// Init initializes the instance.
func (i MPSImageReduceRowSum) Init() MPSImageReduceRowSum {
	rv := objc.Send[MPSImageReduceRowSum](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageReduceRowSum) Autorelease() MPSImageReduceRowSum {
	rv := objc.Send[MPSImageReduceRowSum](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageReduceRowSum creates a new MPSImageReduceRowSum instance.
func NewMPSImageReduceRowSum() MPSImageReduceRowSum {
	class := getMPSImageReduceRowSumClass()
	rv := objc.Send[MPSImageReduceRowSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageReduceRowSumWithCoder(aDecoder foundation.INSCoder) MPSImageReduceRowSum {
	instance := getMPSImageReduceRowSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageReduceRowSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageReduceRowSumWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageReduceRowSum {
	instance := getMPSImageReduceRowSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageReduceRowSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowSum/init(device:)
func NewImageReduceRowSumWithDevice(device metal.MTLDevice) MPSImageReduceRowSum {
	instance := getMPSImageReduceRowSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageReduceRowSumFromID(rv)
}

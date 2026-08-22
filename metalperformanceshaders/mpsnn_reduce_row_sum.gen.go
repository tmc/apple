// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceRowSum] class.
var (
	_MPSNNReduceRowSumClass     MPSNNReduceRowSumClass
	_MPSNNReduceRowSumClassOnce sync.Once
)

func getMPSNNReduceRowSumClass() MPSNNReduceRowSumClass {
	_MPSNNReduceRowSumClassOnce.Do(func() {
		_MPSNNReduceRowSumClass = MPSNNReduceRowSumClass{class: objc.GetClass("MPSNNReduceRowSum")}
	})
	return _MPSNNReduceRowSumClass
}

// GetMPSNNReduceRowSumClass returns the class object for MPSNNReduceRowSum.
func GetMPSNNReduceRowSumClass() MPSNNReduceRowSumClass {
	return getMPSNNReduceRowSumClass()
}

type MPSNNReduceRowSumClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceRowSumClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceRowSumClass) Alloc() MPSNNReduceRowSum {
	rv := objc.Send[MPSNNReduceRowSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the sum of all values for each row in an
// image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowSum
type MPSNNReduceRowSum struct {
	MPSNNReduceUnary
}

// MPSNNReduceRowSumFromID constructs a [MPSNNReduceRowSum] from an objc.ID.
//
// A reduction filter that returns the sum of all values for each row in an
// image.
func MPSNNReduceRowSumFromID(id objc.ID) MPSNNReduceRowSum {
	return MPSNNReduceRowSum{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceRowSum adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceRowSum] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowSum
type IMPSNNReduceRowSum interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceRowSum) Init() MPSNNReduceRowSum {
	rv := objc.Send[MPSNNReduceRowSum](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceRowSum) Autorelease() MPSNNReduceRowSum {
	rv := objc.Send[MPSNNReduceRowSum](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceRowSum creates a new MPSNNReduceRowSum instance.
func NewMPSNNReduceRowSum() MPSNNReduceRowSum {
	class := getMPSNNReduceRowSumClass()
	rv := objc.Send[MPSNNReduceRowSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceRowSumWithCoder(aDecoder foundation.INSCoder) MPSNNReduceRowSum {
	instance := getMPSNNReduceRowSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceRowSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowSum/init(coder:device:)
func NewReduceRowSumWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceRowSum {
	instance := getMPSNNReduceRowSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceRowSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowSum/init(device:)
func NewReduceRowSumWithDevice(device metal.MTLDevice) MPSNNReduceRowSum {
	instance := getMPSNNReduceRowSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceRowSumFromID(rv)
}

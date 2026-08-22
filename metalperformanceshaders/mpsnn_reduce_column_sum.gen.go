// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceColumnSum] class.
var (
	_MPSNNReduceColumnSumClass     MPSNNReduceColumnSumClass
	_MPSNNReduceColumnSumClassOnce sync.Once
)

func getMPSNNReduceColumnSumClass() MPSNNReduceColumnSumClass {
	_MPSNNReduceColumnSumClassOnce.Do(func() {
		_MPSNNReduceColumnSumClass = MPSNNReduceColumnSumClass{class: objc.GetClass("MPSNNReduceColumnSum")}
	})
	return _MPSNNReduceColumnSumClass
}

// GetMPSNNReduceColumnSumClass returns the class object for MPSNNReduceColumnSum.
func GetMPSNNReduceColumnSumClass() MPSNNReduceColumnSumClass {
	return getMPSNNReduceColumnSumClass()
}

type MPSNNReduceColumnSumClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceColumnSumClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceColumnSumClass) Alloc() MPSNNReduceColumnSum {
	rv := objc.Send[MPSNNReduceColumnSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the sum of all values for each column in an
// image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnSum
type MPSNNReduceColumnSum struct {
	MPSNNReduceUnary
}

// MPSNNReduceColumnSumFromID constructs a [MPSNNReduceColumnSum] from an objc.ID.
//
// A reduction filter that returns the sum of all values for each column in an
// image.
func MPSNNReduceColumnSumFromID(id objc.ID) MPSNNReduceColumnSum {
	return MPSNNReduceColumnSum{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceColumnSum adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceColumnSum] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnSum
type IMPSNNReduceColumnSum interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceColumnSum) Init() MPSNNReduceColumnSum {
	rv := objc.Send[MPSNNReduceColumnSum](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceColumnSum) Autorelease() MPSNNReduceColumnSum {
	rv := objc.Send[MPSNNReduceColumnSum](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceColumnSum creates a new MPSNNReduceColumnSum instance.
func NewMPSNNReduceColumnSum() MPSNNReduceColumnSum {
	class := getMPSNNReduceColumnSumClass()
	rv := objc.Send[MPSNNReduceColumnSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceColumnSumWithCoder(aDecoder foundation.INSCoder) MPSNNReduceColumnSum {
	instance := getMPSNNReduceColumnSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceColumnSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnSum/init(coder:device:)
func NewReduceColumnSumWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceColumnSum {
	instance := getMPSNNReduceColumnSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceColumnSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnSum/init(device:)
func NewReduceColumnSumWithDevice(device metal.MTLDevice) MPSNNReduceColumnSum {
	instance := getMPSNNReduceColumnSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceColumnSumFromID(rv)
}

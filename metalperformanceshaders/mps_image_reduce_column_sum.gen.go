// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageReduceColumnSum] class.
var (
	_MPSImageReduceColumnSumClass     MPSImageReduceColumnSumClass
	_MPSImageReduceColumnSumClassOnce sync.Once
)

func getMPSImageReduceColumnSumClass() MPSImageReduceColumnSumClass {
	_MPSImageReduceColumnSumClassOnce.Do(func() {
		_MPSImageReduceColumnSumClass = MPSImageReduceColumnSumClass{class: objc.GetClass("MPSImageReduceColumnSum")}
	})
	return _MPSImageReduceColumnSumClass
}

// GetMPSImageReduceColumnSumClass returns the class object for MPSImageReduceColumnSum.
func GetMPSImageReduceColumnSumClass() MPSImageReduceColumnSumClass {
	return getMPSImageReduceColumnSumClass()
}

type MPSImageReduceColumnSumClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageReduceColumnSumClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageReduceColumnSumClass) Alloc() MPSImageReduceColumnSum {
	rv := objc.Send[MPSImageReduceColumnSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the sum of all values for a column in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnSum
type MPSImageReduceColumnSum struct {
	MPSImageReduceUnary
}

// MPSImageReduceColumnSumFromID constructs a [MPSImageReduceColumnSum] from an objc.ID.
//
// A filter that returns the sum of all values for a column in an image.
func MPSImageReduceColumnSumFromID(id objc.ID) MPSImageReduceColumnSum {
	return MPSImageReduceColumnSum{MPSImageReduceUnary: MPSImageReduceUnaryFromID(id)}
}

// NOTE: MPSImageReduceColumnSum adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageReduceColumnSum] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnSum
type IMPSImageReduceColumnSum interface {
	IMPSImageReduceUnary
}

// Init initializes the instance.
func (i MPSImageReduceColumnSum) Init() MPSImageReduceColumnSum {
	rv := objc.Send[MPSImageReduceColumnSum](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageReduceColumnSum) Autorelease() MPSImageReduceColumnSum {
	rv := objc.Send[MPSImageReduceColumnSum](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageReduceColumnSum creates a new MPSImageReduceColumnSum instance.
func NewMPSImageReduceColumnSum() MPSImageReduceColumnSum {
	class := getMPSImageReduceColumnSumClass()
	rv := objc.Send[MPSImageReduceColumnSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageReduceColumnSumWithCoder(aDecoder foundation.INSCoder) MPSImageReduceColumnSum {
	instance := getMPSImageReduceColumnSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageReduceColumnSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageReduceColumnSumWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageReduceColumnSum {
	instance := getMPSImageReduceColumnSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageReduceColumnSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnSum/init(device:)
func NewImageReduceColumnSumWithDevice(device metal.MTLDevice) MPSImageReduceColumnSum {
	instance := getMPSImageReduceColumnSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageReduceColumnSumFromID(rv)
}

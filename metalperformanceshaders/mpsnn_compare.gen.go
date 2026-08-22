// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNCompare] class.
var (
	_MPSNNCompareClass     MPSNNCompareClass
	_MPSNNCompareClassOnce sync.Once
)

func getMPSNNCompareClass() MPSNNCompareClass {
	_MPSNNCompareClassOnce.Do(func() {
		_MPSNNCompareClass = MPSNNCompareClass{class: objc.GetClass("MPSNNCompare")}
	})
	return _MPSNNCompareClass
}

// GetMPSNNCompareClass returns the class object for MPSNNCompare.
func GetMPSNNCompareClass() MPSNNCompareClass {
	return getMPSNNCompareClass()
}

type MPSNNCompareClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNCompareClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNCompareClass) Alloc() MPSNNCompare {
	rv := objc.Send[MPSNNCompare](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSNNCompare.ComparisonType]
//   - [MPSNNCompare.SetComparisonType]
//   - [MPSNNCompare.Threshold]
//   - [MPSNNCompare.SetThreshold]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCompare
type MPSNNCompare struct {
	MPSCNNArithmetic
}

// MPSNNCompareFromID constructs a [MPSNNCompare] from an objc.ID.
func MPSNNCompareFromID(id objc.ID) MPSNNCompare {
	return MPSNNCompare{MPSCNNArithmetic: MPSCNNArithmeticFromID(id)}
}

// NOTE: MPSNNCompare adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNCompare] class.
//
// # Instance Properties
//
//   - [IMPSNNCompare.ComparisonType]
//   - [IMPSNNCompare.SetComparisonType]
//   - [IMPSNNCompare.Threshold]
//   - [IMPSNNCompare.SetThreshold]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCompare
type IMPSNNCompare interface {
	IMPSCNNArithmetic

	// Topic: Instance Properties

	ComparisonType() MPSNNComparisonType
	SetComparisonType(value MPSNNComparisonType)
	Threshold() float32
	SetThreshold(value float32)
}

// Init initializes the instance.
func (c MPSNNCompare) Init() MPSNNCompare {
	rv := objc.Send[MPSNNCompare](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNCompare) Autorelease() MPSNNCompare {
	rv := objc.Send[MPSNNCompare](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNCompare creates a new MPSNNCompare instance.
func NewMPSNNCompare() MPSNNCompare {
	class := getMPSNNCompareClass()
	rv := objc.Send[MPSNNCompare](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCompareWithCoder(aDecoder foundation.INSCoder) MPSNNCompare {
	instance := getMPSNNCompareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNCompareFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(coder:device:)
func NewCompareWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNCompare {
	instance := getMPSNNCompareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNCompareFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCompare/init(device:)
func NewCompareWithDevice(device metal.MTLDevice) MPSNNCompare {
	instance := getMPSNNCompareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNCompareFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCompare/comparisonType
func (c MPSNNCompare) ComparisonType() MPSNNComparisonType {
	rv := objc.Send[MPSNNComparisonType](c.ID, objc.Sel("comparisonType"))
	return MPSNNComparisonType(rv)
}
func (c MPSNNCompare) SetComparisonType(value MPSNNComparisonType) {
	objc.Send[struct{}](c.ID, objc.Sel("setComparisonType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCompare/threshold
func (c MPSNNCompare) Threshold() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("threshold"))
	return rv
}
func (c MPSNNCompare) SetThreshold(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setThreshold:"), value)
}

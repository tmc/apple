// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceFeatureChannelsArgumentMax] class.
var (
	_MPSNNReduceFeatureChannelsArgumentMaxClass     MPSNNReduceFeatureChannelsArgumentMaxClass
	_MPSNNReduceFeatureChannelsArgumentMaxClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsArgumentMaxClass() MPSNNReduceFeatureChannelsArgumentMaxClass {
	_MPSNNReduceFeatureChannelsArgumentMaxClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsArgumentMaxClass = MPSNNReduceFeatureChannelsArgumentMaxClass{class: objc.GetClass("MPSNNReduceFeatureChannelsArgumentMax")}
	})
	return _MPSNNReduceFeatureChannelsArgumentMaxClass
}

// GetMPSNNReduceFeatureChannelsArgumentMaxClass returns the class object for MPSNNReduceFeatureChannelsArgumentMax.
func GetMPSNNReduceFeatureChannelsArgumentMaxClass() MPSNNReduceFeatureChannelsArgumentMaxClass {
	return getMPSNNReduceFeatureChannelsArgumentMaxClass()
}

type MPSNNReduceFeatureChannelsArgumentMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceFeatureChannelsArgumentMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsArgumentMaxClass) Alloc() MPSNNReduceFeatureChannelsArgumentMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the index of the location of the maximum
// value for each feature channel in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsArgumentMax
type MPSNNReduceFeatureChannelsArgumentMax struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsArgumentMaxFromID constructs a [MPSNNReduceFeatureChannelsArgumentMax] from an objc.ID.
//
// A reduction filter that returns the index of the location of the maximum
// value for each feature channel in an image.
func MPSNNReduceFeatureChannelsArgumentMaxFromID(id objc.ID) MPSNNReduceFeatureChannelsArgumentMax {
	return MPSNNReduceFeatureChannelsArgumentMax{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceFeatureChannelsArgumentMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceFeatureChannelsArgumentMax] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsArgumentMax
type IMPSNNReduceFeatureChannelsArgumentMax interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsArgumentMax) Init() MPSNNReduceFeatureChannelsArgumentMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMax](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsArgumentMax) Autorelease() MPSNNReduceFeatureChannelsArgumentMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMax](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsArgumentMax creates a new MPSNNReduceFeatureChannelsArgumentMax instance.
func NewMPSNNReduceFeatureChannelsArgumentMax() MPSNNReduceFeatureChannelsArgumentMax {
	class := getMPSNNReduceFeatureChannelsArgumentMaxClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceFeatureChannelsArgumentMaxWithCoder(aDecoder foundation.INSCoder) MPSNNReduceFeatureChannelsArgumentMax {
	instance := getMPSNNReduceFeatureChannelsArgumentMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceFeatureChannelsArgumentMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsArgumentMax/init(coder:device:)
func NewReduceFeatureChannelsArgumentMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceFeatureChannelsArgumentMax {
	instance := getMPSNNReduceFeatureChannelsArgumentMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceFeatureChannelsArgumentMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsArgumentMax/init(device:)
func NewReduceFeatureChannelsArgumentMaxWithDevice(device metal.MTLDevice) MPSNNReduceFeatureChannelsArgumentMax {
	instance := getMPSNNReduceFeatureChannelsArgumentMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsArgumentMaxFromID(rv)
}

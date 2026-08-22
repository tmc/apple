// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceFeatureChannelsMax] class.
var (
	_MPSNNReduceFeatureChannelsMaxClass     MPSNNReduceFeatureChannelsMaxClass
	_MPSNNReduceFeatureChannelsMaxClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsMaxClass() MPSNNReduceFeatureChannelsMaxClass {
	_MPSNNReduceFeatureChannelsMaxClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsMaxClass = MPSNNReduceFeatureChannelsMaxClass{class: objc.GetClass("MPSNNReduceFeatureChannelsMax")}
	})
	return _MPSNNReduceFeatureChannelsMaxClass
}

// GetMPSNNReduceFeatureChannelsMaxClass returns the class object for MPSNNReduceFeatureChannelsMax.
func GetMPSNNReduceFeatureChannelsMaxClass() MPSNNReduceFeatureChannelsMaxClass {
	return getMPSNNReduceFeatureChannelsMaxClass()
}

type MPSNNReduceFeatureChannelsMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceFeatureChannelsMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsMaxClass) Alloc() MPSNNReduceFeatureChannelsMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the maximum value for each feature channel
// in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMax
type MPSNNReduceFeatureChannelsMax struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsMaxFromID constructs a [MPSNNReduceFeatureChannelsMax] from an objc.ID.
//
// A reduction filter that returns the maximum value for each feature channel
// in an image.
func MPSNNReduceFeatureChannelsMaxFromID(id objc.ID) MPSNNReduceFeatureChannelsMax {
	return MPSNNReduceFeatureChannelsMax{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceFeatureChannelsMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceFeatureChannelsMax] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMax
type IMPSNNReduceFeatureChannelsMax interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsMax) Init() MPSNNReduceFeatureChannelsMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsMax](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsMax) Autorelease() MPSNNReduceFeatureChannelsMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsMax](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsMax creates a new MPSNNReduceFeatureChannelsMax instance.
func NewMPSNNReduceFeatureChannelsMax() MPSNNReduceFeatureChannelsMax {
	class := getMPSNNReduceFeatureChannelsMaxClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceFeatureChannelsMaxWithCoder(aDecoder foundation.INSCoder) MPSNNReduceFeatureChannelsMax {
	instance := getMPSNNReduceFeatureChannelsMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceFeatureChannelsMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMax/init(coder:device:)
func NewReduceFeatureChannelsMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceFeatureChannelsMax {
	instance := getMPSNNReduceFeatureChannelsMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceFeatureChannelsMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMax/init(device:)
func NewReduceFeatureChannelsMaxWithDevice(device metal.MTLDevice) MPSNNReduceFeatureChannelsMax {
	instance := getMPSNNReduceFeatureChannelsMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsMaxFromID(rv)
}

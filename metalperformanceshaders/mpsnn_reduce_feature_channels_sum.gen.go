// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceFeatureChannelsSum] class.
var (
	_MPSNNReduceFeatureChannelsSumClass     MPSNNReduceFeatureChannelsSumClass
	_MPSNNReduceFeatureChannelsSumClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsSumClass() MPSNNReduceFeatureChannelsSumClass {
	_MPSNNReduceFeatureChannelsSumClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsSumClass = MPSNNReduceFeatureChannelsSumClass{class: objc.GetClass("MPSNNReduceFeatureChannelsSum")}
	})
	return _MPSNNReduceFeatureChannelsSumClass
}

// GetMPSNNReduceFeatureChannelsSumClass returns the class object for MPSNNReduceFeatureChannelsSum.
func GetMPSNNReduceFeatureChannelsSumClass() MPSNNReduceFeatureChannelsSumClass {
	return getMPSNNReduceFeatureChannelsSumClass()
}

type MPSNNReduceFeatureChannelsSumClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceFeatureChannelsSumClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsSumClass) Alloc() MPSNNReduceFeatureChannelsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the sum of all values for each feature
// channel in an image.
//
// # Instance Properties
//
//   - [MPSNNReduceFeatureChannelsSum.Weight]
//   - [MPSNNReduceFeatureChannelsSum.SetWeight]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsSum
type MPSNNReduceFeatureChannelsSum struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsSumFromID constructs a [MPSNNReduceFeatureChannelsSum] from an objc.ID.
//
// A reduction filter that returns the sum of all values for each feature
// channel in an image.
func MPSNNReduceFeatureChannelsSumFromID(id objc.ID) MPSNNReduceFeatureChannelsSum {
	return MPSNNReduceFeatureChannelsSum{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceFeatureChannelsSum adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceFeatureChannelsSum] class.
//
// # Instance Properties
//
//   - [IMPSNNReduceFeatureChannelsSum.Weight]
//   - [IMPSNNReduceFeatureChannelsSum.SetWeight]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsSum
type IMPSNNReduceFeatureChannelsSum interface {
	IMPSNNReduceUnary

	// Topic: Instance Properties

	Weight() float32
	SetWeight(value float32)
}

// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsSum) Init() MPSNNReduceFeatureChannelsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsSum](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsSum) Autorelease() MPSNNReduceFeatureChannelsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsSum](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsSum creates a new MPSNNReduceFeatureChannelsSum instance.
func NewMPSNNReduceFeatureChannelsSum() MPSNNReduceFeatureChannelsSum {
	class := getMPSNNReduceFeatureChannelsSumClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceFeatureChannelsSumWithCoder(aDecoder foundation.INSCoder) MPSNNReduceFeatureChannelsSum {
	instance := getMPSNNReduceFeatureChannelsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceFeatureChannelsSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsSum/init(coder:device:)
func NewReduceFeatureChannelsSumWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceFeatureChannelsSum {
	instance := getMPSNNReduceFeatureChannelsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceFeatureChannelsSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsSum/init(device:)
func NewReduceFeatureChannelsSumWithDevice(device metal.MTLDevice) MPSNNReduceFeatureChannelsSum {
	instance := getMPSNNReduceFeatureChannelsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsSum/weight
func (r MPSNNReduceFeatureChannelsSum) Weight() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("weight"))
	return rv
}
func (r MPSNNReduceFeatureChannelsSum) SetWeight(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setWeight:"), value)
}

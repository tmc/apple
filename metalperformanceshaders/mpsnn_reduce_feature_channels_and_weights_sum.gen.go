// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceFeatureChannelsAndWeightsSum] class.
var (
	_MPSNNReduceFeatureChannelsAndWeightsSumClass     MPSNNReduceFeatureChannelsAndWeightsSumClass
	_MPSNNReduceFeatureChannelsAndWeightsSumClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsAndWeightsSumClass() MPSNNReduceFeatureChannelsAndWeightsSumClass {
	_MPSNNReduceFeatureChannelsAndWeightsSumClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsAndWeightsSumClass = MPSNNReduceFeatureChannelsAndWeightsSumClass{class: objc.GetClass("MPSNNReduceFeatureChannelsAndWeightsSum")}
	})
	return _MPSNNReduceFeatureChannelsAndWeightsSumClass
}

// GetMPSNNReduceFeatureChannelsAndWeightsSumClass returns the class object for MPSNNReduceFeatureChannelsAndWeightsSum.
func GetMPSNNReduceFeatureChannelsAndWeightsSumClass() MPSNNReduceFeatureChannelsAndWeightsSumClass {
	return getMPSNNReduceFeatureChannelsAndWeightsSumClass()
}

type MPSNNReduceFeatureChannelsAndWeightsSumClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceFeatureChannelsAndWeightsSumClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsAndWeightsSumClass) Alloc() MPSNNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the weighted sum of all values for each
// feature channel in an image.
//
// # Initializers
//
//   - [MPSNNReduceFeatureChannelsAndWeightsSum.InitWithDeviceDoWeightedSumByNonZeroWeights]
//
// # Instance Properties
//
//   - [MPSNNReduceFeatureChannelsAndWeightsSum.DoWeightedSumByNonZeroWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsSum
type MPSNNReduceFeatureChannelsAndWeightsSum struct {
	MPSNNReduceBinary
}

// MPSNNReduceFeatureChannelsAndWeightsSumFromID constructs a [MPSNNReduceFeatureChannelsAndWeightsSum] from an objc.ID.
//
// A reduction filter that returns the weighted sum of all values for each
// feature channel in an image.
func MPSNNReduceFeatureChannelsAndWeightsSumFromID(id objc.ID) MPSNNReduceFeatureChannelsAndWeightsSum {
	return MPSNNReduceFeatureChannelsAndWeightsSum{MPSNNReduceBinary: MPSNNReduceBinaryFromID(id)}
}

// NOTE: MPSNNReduceFeatureChannelsAndWeightsSum adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceFeatureChannelsAndWeightsSum] class.
//
// # Initializers
//
//   - [IMPSNNReduceFeatureChannelsAndWeightsSum.InitWithDeviceDoWeightedSumByNonZeroWeights]
//
// # Instance Properties
//
//   - [IMPSNNReduceFeatureChannelsAndWeightsSum.DoWeightedSumByNonZeroWeights]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsSum
type IMPSNNReduceFeatureChannelsAndWeightsSum interface {
	IMPSNNReduceBinary

	// Topic: Initializers

	InitWithDeviceDoWeightedSumByNonZeroWeights(device metal.MTLDevice, doWeightedSumByNonZeroWeights bool) MPSNNReduceFeatureChannelsAndWeightsSum

	// Topic: Instance Properties

	DoWeightedSumByNonZeroWeights() bool
}

// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsAndWeightsSum) Init() MPSNNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsSum](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsAndWeightsSum) Autorelease() MPSNNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsSum](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsAndWeightsSum creates a new MPSNNReduceFeatureChannelsAndWeightsSum instance.
func NewMPSNNReduceFeatureChannelsAndWeightsSum() MPSNNReduceFeatureChannelsAndWeightsSum {
	class := getMPSNNReduceFeatureChannelsAndWeightsSumClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceFeatureChannelsAndWeightsSumWithCoder(aDecoder foundation.INSCoder) MPSNNReduceFeatureChannelsAndWeightsSum {
	instance := getMPSNNReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceFeatureChannelsAndWeightsSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsSum/init(coder:device:)
func NewReduceFeatureChannelsAndWeightsSumWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceFeatureChannelsAndWeightsSum {
	instance := getMPSNNReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceFeatureChannelsAndWeightsSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsSum/init(device:)
func NewReduceFeatureChannelsAndWeightsSumWithDevice(device metal.MTLDevice) MPSNNReduceFeatureChannelsAndWeightsSum {
	instance := getMPSNNReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsAndWeightsSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsSum/init(device:doWeightedSumByNonZeroWeights:)
func NewReduceFeatureChannelsAndWeightsSumWithDeviceDoWeightedSumByNonZeroWeights(device metal.MTLDevice, doWeightedSumByNonZeroWeights bool) MPSNNReduceFeatureChannelsAndWeightsSum {
	instance := getMPSNNReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:doWeightedSumByNonZeroWeights:"), device, doWeightedSumByNonZeroWeights)
	return MPSNNReduceFeatureChannelsAndWeightsSumFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsSum/init(device:doWeightedSumByNonZeroWeights:)
func (r MPSNNReduceFeatureChannelsAndWeightsSum) InitWithDeviceDoWeightedSumByNonZeroWeights(device metal.MTLDevice, doWeightedSumByNonZeroWeights bool) MPSNNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsSum](r.ID, objc.Sel("initWithDevice:doWeightedSumByNonZeroWeights:"), device, doWeightedSumByNonZeroWeights)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsSum/doWeightedSumByNonZeroWeights
func (r MPSNNReduceFeatureChannelsAndWeightsSum) DoWeightedSumByNonZeroWeights() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("doWeightedSumByNonZeroWeights"))
	return rv
}

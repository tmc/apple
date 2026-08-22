// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceFeatureChannelsAndWeightsMean] class.
var (
	_MPSNNReduceFeatureChannelsAndWeightsMeanClass     MPSNNReduceFeatureChannelsAndWeightsMeanClass
	_MPSNNReduceFeatureChannelsAndWeightsMeanClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsAndWeightsMeanClass() MPSNNReduceFeatureChannelsAndWeightsMeanClass {
	_MPSNNReduceFeatureChannelsAndWeightsMeanClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsAndWeightsMeanClass = MPSNNReduceFeatureChannelsAndWeightsMeanClass{class: objc.GetClass("MPSNNReduceFeatureChannelsAndWeightsMean")}
	})
	return _MPSNNReduceFeatureChannelsAndWeightsMeanClass
}

// GetMPSNNReduceFeatureChannelsAndWeightsMeanClass returns the class object for MPSNNReduceFeatureChannelsAndWeightsMean.
func GetMPSNNReduceFeatureChannelsAndWeightsMeanClass() MPSNNReduceFeatureChannelsAndWeightsMeanClass {
	return getMPSNNReduceFeatureChannelsAndWeightsMeanClass()
}

type MPSNNReduceFeatureChannelsAndWeightsMeanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceFeatureChannelsAndWeightsMeanClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsAndWeightsMeanClass) Alloc() MPSNNReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the weighted sum for each feature channel
// in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsMean
type MPSNNReduceFeatureChannelsAndWeightsMean struct {
	MPSNNReduceBinary
}

// MPSNNReduceFeatureChannelsAndWeightsMeanFromID constructs a [MPSNNReduceFeatureChannelsAndWeightsMean] from an objc.ID.
//
// A reduction filter that returns the weighted sum for each feature channel
// in an image.
func MPSNNReduceFeatureChannelsAndWeightsMeanFromID(id objc.ID) MPSNNReduceFeatureChannelsAndWeightsMean {
	return MPSNNReduceFeatureChannelsAndWeightsMean{MPSNNReduceBinary: MPSNNReduceBinaryFromID(id)}
}

// NOTE: MPSNNReduceFeatureChannelsAndWeightsMean adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceFeatureChannelsAndWeightsMean] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsMean
type IMPSNNReduceFeatureChannelsAndWeightsMean interface {
	IMPSNNReduceBinary
}

// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsAndWeightsMean) Init() MPSNNReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsMean](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsAndWeightsMean) Autorelease() MPSNNReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsMean](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsAndWeightsMean creates a new MPSNNReduceFeatureChannelsAndWeightsMean instance.
func NewMPSNNReduceFeatureChannelsAndWeightsMean() MPSNNReduceFeatureChannelsAndWeightsMean {
	class := getMPSNNReduceFeatureChannelsAndWeightsMeanClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceFeatureChannelsAndWeightsMeanWithCoder(aDecoder foundation.INSCoder) MPSNNReduceFeatureChannelsAndWeightsMean {
	instance := getMPSNNReduceFeatureChannelsAndWeightsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceFeatureChannelsAndWeightsMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsMean/init(coder:device:)
func NewReduceFeatureChannelsAndWeightsMeanWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceFeatureChannelsAndWeightsMean {
	instance := getMPSNNReduceFeatureChannelsAndWeightsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceFeatureChannelsAndWeightsMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsMean/init(device:)
func NewReduceFeatureChannelsAndWeightsMeanWithDevice(device metal.MTLDevice) MPSNNReduceFeatureChannelsAndWeightsMean {
	instance := getMPSNNReduceFeatureChannelsAndWeightsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsAndWeightsMeanFromID(rv)
}

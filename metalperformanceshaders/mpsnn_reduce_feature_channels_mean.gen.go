// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceFeatureChannelsMean] class.
var (
	_MPSNNReduceFeatureChannelsMeanClass     MPSNNReduceFeatureChannelsMeanClass
	_MPSNNReduceFeatureChannelsMeanClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsMeanClass() MPSNNReduceFeatureChannelsMeanClass {
	_MPSNNReduceFeatureChannelsMeanClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsMeanClass = MPSNNReduceFeatureChannelsMeanClass{class: objc.GetClass("MPSNNReduceFeatureChannelsMean")}
	})
	return _MPSNNReduceFeatureChannelsMeanClass
}

// GetMPSNNReduceFeatureChannelsMeanClass returns the class object for MPSNNReduceFeatureChannelsMean.
func GetMPSNNReduceFeatureChannelsMeanClass() MPSNNReduceFeatureChannelsMeanClass {
	return getMPSNNReduceFeatureChannelsMeanClass()
}

type MPSNNReduceFeatureChannelsMeanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceFeatureChannelsMeanClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsMeanClass) Alloc() MPSNNReduceFeatureChannelsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the mean value for each feature channel in
// an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMean
type MPSNNReduceFeatureChannelsMean struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsMeanFromID constructs a [MPSNNReduceFeatureChannelsMean] from an objc.ID.
//
// A reduction filter that returns the mean value for each feature channel in
// an image.
func MPSNNReduceFeatureChannelsMeanFromID(id objc.ID) MPSNNReduceFeatureChannelsMean {
	return MPSNNReduceFeatureChannelsMean{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceFeatureChannelsMean adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceFeatureChannelsMean] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMean
type IMPSNNReduceFeatureChannelsMean interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsMean) Init() MPSNNReduceFeatureChannelsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsMean](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsMean) Autorelease() MPSNNReduceFeatureChannelsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsMean](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsMean creates a new MPSNNReduceFeatureChannelsMean instance.
func NewMPSNNReduceFeatureChannelsMean() MPSNNReduceFeatureChannelsMean {
	class := getMPSNNReduceFeatureChannelsMeanClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceFeatureChannelsMeanWithCoder(aDecoder foundation.INSCoder) MPSNNReduceFeatureChannelsMean {
	instance := getMPSNNReduceFeatureChannelsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceFeatureChannelsMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMean/init(coder:device:)
func NewReduceFeatureChannelsMeanWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceFeatureChannelsMean {
	instance := getMPSNNReduceFeatureChannelsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceFeatureChannelsMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsMean/init(device:)
func NewReduceFeatureChannelsMeanWithDevice(device metal.MTLDevice) MPSNNReduceFeatureChannelsMean {
	instance := getMPSNNReduceFeatureChannelsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsMeanFromID(rv)
}

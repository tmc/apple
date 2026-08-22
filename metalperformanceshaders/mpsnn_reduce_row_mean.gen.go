// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceRowMean] class.
var (
	_MPSNNReduceRowMeanClass     MPSNNReduceRowMeanClass
	_MPSNNReduceRowMeanClassOnce sync.Once
)

func getMPSNNReduceRowMeanClass() MPSNNReduceRowMeanClass {
	_MPSNNReduceRowMeanClassOnce.Do(func() {
		_MPSNNReduceRowMeanClass = MPSNNReduceRowMeanClass{class: objc.GetClass("MPSNNReduceRowMean")}
	})
	return _MPSNNReduceRowMeanClass
}

// GetMPSNNReduceRowMeanClass returns the class object for MPSNNReduceRowMean.
func GetMPSNNReduceRowMeanClass() MPSNNReduceRowMeanClass {
	return getMPSNNReduceRowMeanClass()
}

type MPSNNReduceRowMeanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceRowMeanClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceRowMeanClass) Alloc() MPSNNReduceRowMean {
	rv := objc.Send[MPSNNReduceRowMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the mean value for each row in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMean
type MPSNNReduceRowMean struct {
	MPSNNReduceUnary
}

// MPSNNReduceRowMeanFromID constructs a [MPSNNReduceRowMean] from an objc.ID.
//
// A reduction filter that returns the mean value for each row in an image.
func MPSNNReduceRowMeanFromID(id objc.ID) MPSNNReduceRowMean {
	return MPSNNReduceRowMean{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceRowMean adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceRowMean] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMean
type IMPSNNReduceRowMean interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceRowMean) Init() MPSNNReduceRowMean {
	rv := objc.Send[MPSNNReduceRowMean](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceRowMean) Autorelease() MPSNNReduceRowMean {
	rv := objc.Send[MPSNNReduceRowMean](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceRowMean creates a new MPSNNReduceRowMean instance.
func NewMPSNNReduceRowMean() MPSNNReduceRowMean {
	class := getMPSNNReduceRowMeanClass()
	rv := objc.Send[MPSNNReduceRowMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceRowMeanWithCoder(aDecoder foundation.INSCoder) MPSNNReduceRowMean {
	instance := getMPSNNReduceRowMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceRowMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMean/init(coder:device:)
func NewReduceRowMeanWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceRowMean {
	instance := getMPSNNReduceRowMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceRowMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMean/init(device:)
func NewReduceRowMeanWithDevice(device metal.MTLDevice) MPSNNReduceRowMean {
	instance := getMPSNNReduceRowMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceRowMeanFromID(rv)
}

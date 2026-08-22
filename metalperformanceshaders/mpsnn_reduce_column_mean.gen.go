// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceColumnMean] class.
var (
	_MPSNNReduceColumnMeanClass     MPSNNReduceColumnMeanClass
	_MPSNNReduceColumnMeanClassOnce sync.Once
)

func getMPSNNReduceColumnMeanClass() MPSNNReduceColumnMeanClass {
	_MPSNNReduceColumnMeanClassOnce.Do(func() {
		_MPSNNReduceColumnMeanClass = MPSNNReduceColumnMeanClass{class: objc.GetClass("MPSNNReduceColumnMean")}
	})
	return _MPSNNReduceColumnMeanClass
}

// GetMPSNNReduceColumnMeanClass returns the class object for MPSNNReduceColumnMean.
func GetMPSNNReduceColumnMeanClass() MPSNNReduceColumnMeanClass {
	return getMPSNNReduceColumnMeanClass()
}

type MPSNNReduceColumnMeanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceColumnMeanClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceColumnMeanClass) Alloc() MPSNNReduceColumnMean {
	rv := objc.Send[MPSNNReduceColumnMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the mean value for each column in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMean
type MPSNNReduceColumnMean struct {
	MPSNNReduceUnary
}

// MPSNNReduceColumnMeanFromID constructs a [MPSNNReduceColumnMean] from an objc.ID.
//
// A reduction filter that returns the mean value for each column in an image.
func MPSNNReduceColumnMeanFromID(id objc.ID) MPSNNReduceColumnMean {
	return MPSNNReduceColumnMean{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceColumnMean adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceColumnMean] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMean
type IMPSNNReduceColumnMean interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceColumnMean) Init() MPSNNReduceColumnMean {
	rv := objc.Send[MPSNNReduceColumnMean](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceColumnMean) Autorelease() MPSNNReduceColumnMean {
	rv := objc.Send[MPSNNReduceColumnMean](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceColumnMean creates a new MPSNNReduceColumnMean instance.
func NewMPSNNReduceColumnMean() MPSNNReduceColumnMean {
	class := getMPSNNReduceColumnMeanClass()
	rv := objc.Send[MPSNNReduceColumnMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceColumnMeanWithCoder(aDecoder foundation.INSCoder) MPSNNReduceColumnMean {
	instance := getMPSNNReduceColumnMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceColumnMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMean/init(coder:device:)
func NewReduceColumnMeanWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceColumnMean {
	instance := getMPSNNReduceColumnMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceColumnMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMean/init(device:)
func NewReduceColumnMeanWithDevice(device metal.MTLDevice) MPSNNReduceColumnMean {
	instance := getMPSNNReduceColumnMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceColumnMeanFromID(rv)
}

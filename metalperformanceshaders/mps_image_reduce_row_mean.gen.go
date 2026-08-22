// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageReduceRowMean] class.
var (
	_MPSImageReduceRowMeanClass     MPSImageReduceRowMeanClass
	_MPSImageReduceRowMeanClassOnce sync.Once
)

func getMPSImageReduceRowMeanClass() MPSImageReduceRowMeanClass {
	_MPSImageReduceRowMeanClassOnce.Do(func() {
		_MPSImageReduceRowMeanClass = MPSImageReduceRowMeanClass{class: objc.GetClass("MPSImageReduceRowMean")}
	})
	return _MPSImageReduceRowMeanClass
}

// GetMPSImageReduceRowMeanClass returns the class object for MPSImageReduceRowMean.
func GetMPSImageReduceRowMeanClass() MPSImageReduceRowMeanClass {
	return getMPSImageReduceRowMeanClass()
}

type MPSImageReduceRowMeanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageReduceRowMeanClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageReduceRowMeanClass) Alloc() MPSImageReduceRowMean {
	rv := objc.Send[MPSImageReduceRowMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the mean value for each row in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMean
type MPSImageReduceRowMean struct {
	MPSImageReduceUnary
}

// MPSImageReduceRowMeanFromID constructs a [MPSImageReduceRowMean] from an objc.ID.
//
// A filter that returns the mean value for each row in an image.
func MPSImageReduceRowMeanFromID(id objc.ID) MPSImageReduceRowMean {
	return MPSImageReduceRowMean{MPSImageReduceUnary: MPSImageReduceUnaryFromID(id)}
}

// NOTE: MPSImageReduceRowMean adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageReduceRowMean] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMean
type IMPSImageReduceRowMean interface {
	IMPSImageReduceUnary
}

// Init initializes the instance.
func (i MPSImageReduceRowMean) Init() MPSImageReduceRowMean {
	rv := objc.Send[MPSImageReduceRowMean](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageReduceRowMean) Autorelease() MPSImageReduceRowMean {
	rv := objc.Send[MPSImageReduceRowMean](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageReduceRowMean creates a new MPSImageReduceRowMean instance.
func NewMPSImageReduceRowMean() MPSImageReduceRowMean {
	class := getMPSImageReduceRowMeanClass()
	rv := objc.Send[MPSImageReduceRowMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageReduceRowMeanWithCoder(aDecoder foundation.INSCoder) MPSImageReduceRowMean {
	instance := getMPSImageReduceRowMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageReduceRowMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageReduceRowMeanWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageReduceRowMean {
	instance := getMPSImageReduceRowMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageReduceRowMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMean/init(device:)
func NewImageReduceRowMeanWithDevice(device metal.MTLDevice) MPSImageReduceRowMean {
	instance := getMPSImageReduceRowMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageReduceRowMeanFromID(rv)
}

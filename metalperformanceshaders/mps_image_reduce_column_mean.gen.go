// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageReduceColumnMean] class.
var (
	_MPSImageReduceColumnMeanClass     MPSImageReduceColumnMeanClass
	_MPSImageReduceColumnMeanClassOnce sync.Once
)

func getMPSImageReduceColumnMeanClass() MPSImageReduceColumnMeanClass {
	_MPSImageReduceColumnMeanClassOnce.Do(func() {
		_MPSImageReduceColumnMeanClass = MPSImageReduceColumnMeanClass{class: objc.GetClass("MPSImageReduceColumnMean")}
	})
	return _MPSImageReduceColumnMeanClass
}

// GetMPSImageReduceColumnMeanClass returns the class object for MPSImageReduceColumnMean.
func GetMPSImageReduceColumnMeanClass() MPSImageReduceColumnMeanClass {
	return getMPSImageReduceColumnMeanClass()
}

type MPSImageReduceColumnMeanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageReduceColumnMeanClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageReduceColumnMeanClass) Alloc() MPSImageReduceColumnMean {
	rv := objc.Send[MPSImageReduceColumnMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the mean value for each column in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMean
type MPSImageReduceColumnMean struct {
	MPSImageReduceUnary
}

// MPSImageReduceColumnMeanFromID constructs a [MPSImageReduceColumnMean] from an objc.ID.
//
// A filter that returns the mean value for each column in an image.
func MPSImageReduceColumnMeanFromID(id objc.ID) MPSImageReduceColumnMean {
	return MPSImageReduceColumnMean{MPSImageReduceUnary: MPSImageReduceUnaryFromID(id)}
}

// NOTE: MPSImageReduceColumnMean adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageReduceColumnMean] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMean
type IMPSImageReduceColumnMean interface {
	IMPSImageReduceUnary
}

// Init initializes the instance.
func (i MPSImageReduceColumnMean) Init() MPSImageReduceColumnMean {
	rv := objc.Send[MPSImageReduceColumnMean](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageReduceColumnMean) Autorelease() MPSImageReduceColumnMean {
	rv := objc.Send[MPSImageReduceColumnMean](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageReduceColumnMean creates a new MPSImageReduceColumnMean instance.
func NewMPSImageReduceColumnMean() MPSImageReduceColumnMean {
	class := getMPSImageReduceColumnMeanClass()
	rv := objc.Send[MPSImageReduceColumnMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageReduceColumnMeanWithCoder(aDecoder foundation.INSCoder) MPSImageReduceColumnMean {
	instance := getMPSImageReduceColumnMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageReduceColumnMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageReduceColumnMeanWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageReduceColumnMean {
	instance := getMPSImageReduceColumnMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageReduceColumnMeanFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMean/init(device:)
func NewImageReduceColumnMeanWithDevice(device metal.MTLDevice) MPSImageReduceColumnMean {
	instance := getMPSImageReduceColumnMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageReduceColumnMeanFromID(rv)
}

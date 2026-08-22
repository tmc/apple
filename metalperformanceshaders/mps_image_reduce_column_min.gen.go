// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageReduceColumnMin] class.
var (
	_MPSImageReduceColumnMinClass     MPSImageReduceColumnMinClass
	_MPSImageReduceColumnMinClassOnce sync.Once
)

func getMPSImageReduceColumnMinClass() MPSImageReduceColumnMinClass {
	_MPSImageReduceColumnMinClassOnce.Do(func() {
		_MPSImageReduceColumnMinClass = MPSImageReduceColumnMinClass{class: objc.GetClass("MPSImageReduceColumnMin")}
	})
	return _MPSImageReduceColumnMinClass
}

// GetMPSImageReduceColumnMinClass returns the class object for MPSImageReduceColumnMin.
func GetMPSImageReduceColumnMinClass() MPSImageReduceColumnMinClass {
	return getMPSImageReduceColumnMinClass()
}

type MPSImageReduceColumnMinClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageReduceColumnMinClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageReduceColumnMinClass) Alloc() MPSImageReduceColumnMin {
	rv := objc.Send[MPSImageReduceColumnMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the minimum value for each column in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMin
type MPSImageReduceColumnMin struct {
	MPSImageReduceUnary
}

// MPSImageReduceColumnMinFromID constructs a [MPSImageReduceColumnMin] from an objc.ID.
//
// A filter that returns the minimum value for each column in an image.
func MPSImageReduceColumnMinFromID(id objc.ID) MPSImageReduceColumnMin {
	return MPSImageReduceColumnMin{MPSImageReduceUnary: MPSImageReduceUnaryFromID(id)}
}

// NOTE: MPSImageReduceColumnMin adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageReduceColumnMin] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMin
type IMPSImageReduceColumnMin interface {
	IMPSImageReduceUnary
}

// Init initializes the instance.
func (i MPSImageReduceColumnMin) Init() MPSImageReduceColumnMin {
	rv := objc.Send[MPSImageReduceColumnMin](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageReduceColumnMin) Autorelease() MPSImageReduceColumnMin {
	rv := objc.Send[MPSImageReduceColumnMin](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageReduceColumnMin creates a new MPSImageReduceColumnMin instance.
func NewMPSImageReduceColumnMin() MPSImageReduceColumnMin {
	class := getMPSImageReduceColumnMinClass()
	rv := objc.Send[MPSImageReduceColumnMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageReduceColumnMinWithCoder(aDecoder foundation.INSCoder) MPSImageReduceColumnMin {
	instance := getMPSImageReduceColumnMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageReduceColumnMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageReduceColumnMinWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageReduceColumnMin {
	instance := getMPSImageReduceColumnMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageReduceColumnMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMin/init(device:)
func NewImageReduceColumnMinWithDevice(device metal.MTLDevice) MPSImageReduceColumnMin {
	instance := getMPSImageReduceColumnMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageReduceColumnMinFromID(rv)
}

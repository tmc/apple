// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageReduceRowMin] class.
var (
	_MPSImageReduceRowMinClass     MPSImageReduceRowMinClass
	_MPSImageReduceRowMinClassOnce sync.Once
)

func getMPSImageReduceRowMinClass() MPSImageReduceRowMinClass {
	_MPSImageReduceRowMinClassOnce.Do(func() {
		_MPSImageReduceRowMinClass = MPSImageReduceRowMinClass{class: objc.GetClass("MPSImageReduceRowMin")}
	})
	return _MPSImageReduceRowMinClass
}

// GetMPSImageReduceRowMinClass returns the class object for MPSImageReduceRowMin.
func GetMPSImageReduceRowMinClass() MPSImageReduceRowMinClass {
	return getMPSImageReduceRowMinClass()
}

type MPSImageReduceRowMinClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageReduceRowMinClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageReduceRowMinClass) Alloc() MPSImageReduceRowMin {
	rv := objc.Send[MPSImageReduceRowMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the minimum value for each row in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMin
type MPSImageReduceRowMin struct {
	MPSImageReduceUnary
}

// MPSImageReduceRowMinFromID constructs a [MPSImageReduceRowMin] from an objc.ID.
//
// A filter that returns the minimum value for each row in an image.
func MPSImageReduceRowMinFromID(id objc.ID) MPSImageReduceRowMin {
	return MPSImageReduceRowMin{MPSImageReduceUnary: MPSImageReduceUnaryFromID(id)}
}

// NOTE: MPSImageReduceRowMin adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageReduceRowMin] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMin
type IMPSImageReduceRowMin interface {
	IMPSImageReduceUnary
}

// Init initializes the instance.
func (i MPSImageReduceRowMin) Init() MPSImageReduceRowMin {
	rv := objc.Send[MPSImageReduceRowMin](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageReduceRowMin) Autorelease() MPSImageReduceRowMin {
	rv := objc.Send[MPSImageReduceRowMin](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageReduceRowMin creates a new MPSImageReduceRowMin instance.
func NewMPSImageReduceRowMin() MPSImageReduceRowMin {
	class := getMPSImageReduceRowMinClass()
	rv := objc.Send[MPSImageReduceRowMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageReduceRowMinWithCoder(aDecoder foundation.INSCoder) MPSImageReduceRowMin {
	instance := getMPSImageReduceRowMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageReduceRowMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageReduceRowMinWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageReduceRowMin {
	instance := getMPSImageReduceRowMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageReduceRowMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMin/init(device:)
func NewImageReduceRowMinWithDevice(device metal.MTLDevice) MPSImageReduceRowMin {
	instance := getMPSImageReduceRowMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageReduceRowMinFromID(rv)
}

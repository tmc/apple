// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceRowMin] class.
var (
	_MPSNNReduceRowMinClass     MPSNNReduceRowMinClass
	_MPSNNReduceRowMinClassOnce sync.Once
)

func getMPSNNReduceRowMinClass() MPSNNReduceRowMinClass {
	_MPSNNReduceRowMinClassOnce.Do(func() {
		_MPSNNReduceRowMinClass = MPSNNReduceRowMinClass{class: objc.GetClass("MPSNNReduceRowMin")}
	})
	return _MPSNNReduceRowMinClass
}

// GetMPSNNReduceRowMinClass returns the class object for MPSNNReduceRowMin.
func GetMPSNNReduceRowMinClass() MPSNNReduceRowMinClass {
	return getMPSNNReduceRowMinClass()
}

type MPSNNReduceRowMinClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceRowMinClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceRowMinClass) Alloc() MPSNNReduceRowMin {
	rv := objc.Send[MPSNNReduceRowMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the minimum value for each row in an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMin
type MPSNNReduceRowMin struct {
	MPSNNReduceUnary
}

// MPSNNReduceRowMinFromID constructs a [MPSNNReduceRowMin] from an objc.ID.
//
// A reduction filter that returns the minimum value for each row in an image.
func MPSNNReduceRowMinFromID(id objc.ID) MPSNNReduceRowMin {
	return MPSNNReduceRowMin{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceRowMin adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceRowMin] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMin
type IMPSNNReduceRowMin interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceRowMin) Init() MPSNNReduceRowMin {
	rv := objc.Send[MPSNNReduceRowMin](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceRowMin) Autorelease() MPSNNReduceRowMin {
	rv := objc.Send[MPSNNReduceRowMin](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceRowMin creates a new MPSNNReduceRowMin instance.
func NewMPSNNReduceRowMin() MPSNNReduceRowMin {
	class := getMPSNNReduceRowMinClass()
	rv := objc.Send[MPSNNReduceRowMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceRowMinWithCoder(aDecoder foundation.INSCoder) MPSNNReduceRowMin {
	instance := getMPSNNReduceRowMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceRowMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMin/init(coder:device:)
func NewReduceRowMinWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceRowMin {
	instance := getMPSNNReduceRowMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceRowMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceRowMin/init(device:)
func NewReduceRowMinWithDevice(device metal.MTLDevice) MPSNNReduceRowMin {
	instance := getMPSNNReduceRowMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceRowMinFromID(rv)
}

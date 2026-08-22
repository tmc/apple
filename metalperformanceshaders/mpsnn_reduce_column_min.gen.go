// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceColumnMin] class.
var (
	_MPSNNReduceColumnMinClass     MPSNNReduceColumnMinClass
	_MPSNNReduceColumnMinClassOnce sync.Once
)

func getMPSNNReduceColumnMinClass() MPSNNReduceColumnMinClass {
	_MPSNNReduceColumnMinClassOnce.Do(func() {
		_MPSNNReduceColumnMinClass = MPSNNReduceColumnMinClass{class: objc.GetClass("MPSNNReduceColumnMin")}
	})
	return _MPSNNReduceColumnMinClass
}

// GetMPSNNReduceColumnMinClass returns the class object for MPSNNReduceColumnMin.
func GetMPSNNReduceColumnMinClass() MPSNNReduceColumnMinClass {
	return getMPSNNReduceColumnMinClass()
}

type MPSNNReduceColumnMinClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceColumnMinClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceColumnMinClass) Alloc() MPSNNReduceColumnMin {
	rv := objc.Send[MPSNNReduceColumnMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A reduction filter that returns the minimum value for each column in an
// image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMin
type MPSNNReduceColumnMin struct {
	MPSNNReduceUnary
}

// MPSNNReduceColumnMinFromID constructs a [MPSNNReduceColumnMin] from an objc.ID.
//
// A reduction filter that returns the minimum value for each column in an
// image.
func MPSNNReduceColumnMinFromID(id objc.ID) MPSNNReduceColumnMin {
	return MPSNNReduceColumnMin{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}

// NOTE: MPSNNReduceColumnMin adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceColumnMin] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMin
type IMPSNNReduceColumnMin interface {
	IMPSNNReduceUnary
}

// Init initializes the instance.
func (r MPSNNReduceColumnMin) Init() MPSNNReduceColumnMin {
	rv := objc.Send[MPSNNReduceColumnMin](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceColumnMin) Autorelease() MPSNNReduceColumnMin {
	rv := objc.Send[MPSNNReduceColumnMin](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceColumnMin creates a new MPSNNReduceColumnMin instance.
func NewMPSNNReduceColumnMin() MPSNNReduceColumnMin {
	class := getMPSNNReduceColumnMinClass()
	rv := objc.Send[MPSNNReduceColumnMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceColumnMinWithCoder(aDecoder foundation.INSCoder) MPSNNReduceColumnMin {
	instance := getMPSNNReduceColumnMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceColumnMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMin/init(coder:device:)
func NewReduceColumnMinWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceColumnMin {
	instance := getMPSNNReduceColumnMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceColumnMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceColumnMin/init(device:)
func NewReduceColumnMinWithDevice(device metal.MTLDevice) MPSNNReduceColumnMin {
	instance := getMPSNNReduceColumnMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceColumnMinFromID(rv)
}

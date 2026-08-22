// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixLogSoftMax] class.
var (
	_MPSMatrixLogSoftMaxClass     MPSMatrixLogSoftMaxClass
	_MPSMatrixLogSoftMaxClassOnce sync.Once
)

func getMPSMatrixLogSoftMaxClass() MPSMatrixLogSoftMaxClass {
	_MPSMatrixLogSoftMaxClassOnce.Do(func() {
		_MPSMatrixLogSoftMaxClass = MPSMatrixLogSoftMaxClass{class: objc.GetClass("MPSMatrixLogSoftMax")}
	})
	return _MPSMatrixLogSoftMaxClass
}

// GetMPSMatrixLogSoftMaxClass returns the class object for MPSMatrixLogSoftMax.
func GetMPSMatrixLogSoftMaxClass() MPSMatrixLogSoftMaxClass {
	return getMPSMatrixLogSoftMaxClass()
}

type MPSMatrixLogSoftMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixLogSoftMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixLogSoftMaxClass) Alloc() MPSMatrixLogSoftMax {
	rv := objc.Send[MPSMatrixLogSoftMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A logarithmic softmax kernel that operates on matrices.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixLogSoftMax
type MPSMatrixLogSoftMax struct {
	MPSMatrixSoftMax
}

// MPSMatrixLogSoftMaxFromID constructs a [MPSMatrixLogSoftMax] from an objc.ID.
//
// A logarithmic softmax kernel that operates on matrices.
func MPSMatrixLogSoftMaxFromID(id objc.ID) MPSMatrixLogSoftMax {
	return MPSMatrixLogSoftMax{MPSMatrixSoftMax: MPSMatrixSoftMaxFromID(id)}
}

// NOTE: MPSMatrixLogSoftMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixLogSoftMax] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixLogSoftMax
type IMPSMatrixLogSoftMax interface {
	IMPSMatrixSoftMax
}

// Init initializes the instance.
func (m MPSMatrixLogSoftMax) Init() MPSMatrixLogSoftMax {
	rv := objc.Send[MPSMatrixLogSoftMax](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixLogSoftMax) Autorelease() MPSMatrixLogSoftMax {
	rv := objc.Send[MPSMatrixLogSoftMax](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixLogSoftMax creates a new MPSMatrixLogSoftMax instance.
func NewMPSMatrixLogSoftMax() MPSMatrixLogSoftMax {
	class := getMPSMatrixLogSoftMaxClass()
	rv := objc.Send[MPSMatrixLogSoftMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixLogSoftMaxWithCoder(aDecoder foundation.INSCoder) MPSMatrixLogSoftMax {
	instance := getMPSMatrixLogSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixLogSoftMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMax/init(coder:device:)
func NewMatrixLogSoftMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixLogSoftMax {
	instance := getMPSMatrixLogSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixLogSoftMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMax/init(device:)
func NewMatrixLogSoftMaxWithDevice(device metal.MTLDevice) MPSMatrixLogSoftMax {
	instance := getMPSMatrixLogSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixLogSoftMaxFromID(rv)
}

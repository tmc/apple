// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixLogSoftMaxGradient] class.
var (
	_MPSMatrixLogSoftMaxGradientClass     MPSMatrixLogSoftMaxGradientClass
	_MPSMatrixLogSoftMaxGradientClassOnce sync.Once
)

func getMPSMatrixLogSoftMaxGradientClass() MPSMatrixLogSoftMaxGradientClass {
	_MPSMatrixLogSoftMaxGradientClassOnce.Do(func() {
		_MPSMatrixLogSoftMaxGradientClass = MPSMatrixLogSoftMaxGradientClass{class: objc.GetClass("MPSMatrixLogSoftMaxGradient")}
	})
	return _MPSMatrixLogSoftMaxGradientClass
}

// GetMPSMatrixLogSoftMaxGradientClass returns the class object for MPSMatrixLogSoftMaxGradient.
func GetMPSMatrixLogSoftMaxGradientClass() MPSMatrixLogSoftMaxGradientClass {
	return getMPSMatrixLogSoftMaxGradientClass()
}

type MPSMatrixLogSoftMaxGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixLogSoftMaxGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixLogSoftMaxGradientClass) Alloc() MPSMatrixLogSoftMaxGradient {
	rv := objc.Send[MPSMatrixLogSoftMaxGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A logarithmic gradient softmax kernel that operates on matrices.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixLogSoftMaxGradient
type MPSMatrixLogSoftMaxGradient struct {
	MPSMatrixSoftMaxGradient
}

// MPSMatrixLogSoftMaxGradientFromID constructs a [MPSMatrixLogSoftMaxGradient] from an objc.ID.
//
// A logarithmic gradient softmax kernel that operates on matrices.
func MPSMatrixLogSoftMaxGradientFromID(id objc.ID) MPSMatrixLogSoftMaxGradient {
	return MPSMatrixLogSoftMaxGradient{MPSMatrixSoftMaxGradient: MPSMatrixSoftMaxGradientFromID(id)}
}

// NOTE: MPSMatrixLogSoftMaxGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixLogSoftMaxGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixLogSoftMaxGradient
type IMPSMatrixLogSoftMaxGradient interface {
	IMPSMatrixSoftMaxGradient
}

// Init initializes the instance.
func (m MPSMatrixLogSoftMaxGradient) Init() MPSMatrixLogSoftMaxGradient {
	rv := objc.Send[MPSMatrixLogSoftMaxGradient](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixLogSoftMaxGradient) Autorelease() MPSMatrixLogSoftMaxGradient {
	rv := objc.Send[MPSMatrixLogSoftMaxGradient](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixLogSoftMaxGradient creates a new MPSMatrixLogSoftMaxGradient instance.
func NewMPSMatrixLogSoftMaxGradient() MPSMatrixLogSoftMaxGradient {
	class := getMPSMatrixLogSoftMaxGradientClass()
	rv := objc.Send[MPSMatrixLogSoftMaxGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixLogSoftMaxGradientWithCoder(aDecoder foundation.INSCoder) MPSMatrixLogSoftMaxGradient {
	instance := getMPSMatrixLogSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixLogSoftMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient/init(coder:device:)
func NewMatrixLogSoftMaxGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixLogSoftMaxGradient {
	instance := getMPSMatrixLogSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixLogSoftMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient/init(device:)
func NewMatrixLogSoftMaxGradientWithDevice(device metal.MTLDevice) MPSMatrixLogSoftMaxGradient {
	instance := getMPSMatrixLogSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixLogSoftMaxGradientFromID(rv)
}

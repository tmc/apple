// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNLogSoftMaxGradient] class.
var (
	_MPSCNNLogSoftMaxGradientClass     MPSCNNLogSoftMaxGradientClass
	_MPSCNNLogSoftMaxGradientClassOnce sync.Once
)

func getMPSCNNLogSoftMaxGradientClass() MPSCNNLogSoftMaxGradientClass {
	_MPSCNNLogSoftMaxGradientClassOnce.Do(func() {
		_MPSCNNLogSoftMaxGradientClass = MPSCNNLogSoftMaxGradientClass{class: objc.GetClass("MPSCNNLogSoftMaxGradient")}
	})
	return _MPSCNNLogSoftMaxGradientClass
}

// GetMPSCNNLogSoftMaxGradientClass returns the class object for MPSCNNLogSoftMaxGradient.
func GetMPSCNNLogSoftMaxGradientClass() MPSCNNLogSoftMaxGradientClass {
	return getMPSCNNLogSoftMaxGradientClass()
}

type MPSCNNLogSoftMaxGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLogSoftMaxGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLogSoftMaxGradientClass) Alloc() MPSCNNLogSoftMaxGradient {
	rv := objc.Send[MPSCNNLogSoftMaxGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient logarithmic softmax filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradient
type MPSCNNLogSoftMaxGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNLogSoftMaxGradientFromID constructs a [MPSCNNLogSoftMaxGradient] from an objc.ID.
//
// A gradient logarithmic softmax filter.
func MPSCNNLogSoftMaxGradientFromID(id objc.ID) MPSCNNLogSoftMaxGradient {
	return MPSCNNLogSoftMaxGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNLogSoftMaxGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLogSoftMaxGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradient
type IMPSCNNLogSoftMaxGradient interface {
	IMPSCNNGradientKernel
}

// Init initializes the instance.
func (c MPSCNNLogSoftMaxGradient) Init() MPSCNNLogSoftMaxGradient {
	rv := objc.Send[MPSCNNLogSoftMaxGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLogSoftMaxGradient) Autorelease() MPSCNNLogSoftMaxGradient {
	rv := objc.Send[MPSCNNLogSoftMaxGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLogSoftMaxGradient creates a new MPSCNNLogSoftMaxGradient instance.
func NewMPSCNNLogSoftMaxGradient() MPSCNNLogSoftMaxGradient {
	class := getMPSCNNLogSoftMaxGradientClass()
	rv := objc.Send[MPSCNNLogSoftMaxGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNLogSoftMaxGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNLogSoftMaxGradient {
	instance := getMPSCNNLogSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNLogSoftMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradient/init(coder:device:)
func NewCNNLogSoftMaxGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNLogSoftMaxGradient {
	instance := getMPSCNNLogSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNLogSoftMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradient/init(device:)
func NewCNNLogSoftMaxGradientWithDevice(device metal.MTLDevice) MPSCNNLogSoftMaxGradient {
	instance := getMPSCNNLogSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNLogSoftMaxGradientFromID(rv)
}

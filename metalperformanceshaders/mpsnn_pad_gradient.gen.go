// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNPadGradient] class.
var (
	_MPSNNPadGradientClass     MPSNNPadGradientClass
	_MPSNNPadGradientClassOnce sync.Once
)

func getMPSNNPadGradientClass() MPSNNPadGradientClass {
	_MPSNNPadGradientClassOnce.Do(func() {
		_MPSNNPadGradientClass = MPSNNPadGradientClass{class: objc.GetClass("MPSNNPadGradient")}
	})
	return _MPSNNPadGradientClass
}

// GetMPSNNPadGradientClass returns the class object for MPSNNPadGradient.
func GetMPSNNPadGradientClass() MPSNNPadGradientClass {
	return getMPSNNPadGradientClass()
}

type MPSNNPadGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNPadGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNPadGradientClass) Alloc() MPSNNPadGradient {
	rv := objc.Send[MPSNNPadGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradient
type MPSNNPadGradient struct {
	MPSCNNGradientKernel
}

// MPSNNPadGradientFromID constructs a [MPSNNPadGradient] from an objc.ID.
func MPSNNPadGradientFromID(id objc.ID) MPSNNPadGradient {
	return MPSNNPadGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSNNPadGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNPadGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradient
type IMPSNNPadGradient interface {
	IMPSCNNGradientKernel
}

// Init initializes the instance.
func (p MPSNNPadGradient) Init() MPSNNPadGradient {
	rv := objc.Send[MPSNNPadGradient](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MPSNNPadGradient) Autorelease() MPSNNPadGradient {
	rv := objc.Send[MPSNNPadGradient](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNPadGradient creates a new MPSNNPadGradient instance.
func NewMPSNNPadGradient() MPSNNPadGradient {
	class := getMPSNNPadGradientClass()
	rv := objc.Send[MPSNNPadGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewPadGradientWithCoder(aDecoder foundation.INSCoder) MPSNNPadGradient {
	instance := getMPSNNPadGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNPadGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradient/init(coder:device:)
func NewPadGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNPadGradient {
	instance := getMPSNNPadGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNPadGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradient/init(device:)
func NewPadGradientWithDevice(device metal.MTLDevice) MPSNNPadGradient {
	instance := getMPSNNPadGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNPadGradientFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNInitialGradient] class.
var (
	_MPSNNInitialGradientClass     MPSNNInitialGradientClass
	_MPSNNInitialGradientClassOnce sync.Once
)

func getMPSNNInitialGradientClass() MPSNNInitialGradientClass {
	_MPSNNInitialGradientClassOnce.Do(func() {
		_MPSNNInitialGradientClass = MPSNNInitialGradientClass{class: objc.GetClass("MPSNNInitialGradient")}
	})
	return _MPSNNInitialGradientClass
}

// GetMPSNNInitialGradientClass returns the class object for MPSNNInitialGradient.
func GetMPSNNInitialGradientClass() MPSNNInitialGradientClass {
	return getMPSNNInitialGradientClass()
}

type MPSNNInitialGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNInitialGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNInitialGradientClass) Alloc() MPSNNInitialGradient {
	rv := objc.Send[MPSNNInitialGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradient
type MPSNNInitialGradient struct {
	MPSCNNKernel
}

// MPSNNInitialGradientFromID constructs a [MPSNNInitialGradient] from an objc.ID.
func MPSNNInitialGradientFromID(id objc.ID) MPSNNInitialGradient {
	return MPSNNInitialGradient{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSNNInitialGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNInitialGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradient
type IMPSNNInitialGradient interface {
	IMPSCNNKernel
}

// Init initializes the instance.
func (i MPSNNInitialGradient) Init() MPSNNInitialGradient {
	rv := objc.Send[MPSNNInitialGradient](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSNNInitialGradient) Autorelease() MPSNNInitialGradient {
	rv := objc.Send[MPSNNInitialGradient](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNInitialGradient creates a new MPSNNInitialGradient instance.
func NewMPSNNInitialGradient() MPSNNInitialGradient {
	class := getMPSNNInitialGradientClass()
	rv := objc.Send[MPSNNInitialGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewInitialGradientWithCoder(aDecoder foundation.INSCoder) MPSNNInitialGradient {
	instance := getMPSNNInitialGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNInitialGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(coder:device:)
func NewInitialGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNInitialGradient {
	instance := getMPSNNInitialGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNInitialGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradient/init(device:)
func NewInitialGradientWithDevice(device metal.MTLDevice) MPSNNInitialGradient {
	instance := getMPSNNInitialGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNInitialGradientFromID(rv)
}

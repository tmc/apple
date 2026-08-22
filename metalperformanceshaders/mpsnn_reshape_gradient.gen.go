// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReshapeGradient] class.
var (
	_MPSNNReshapeGradientClass     MPSNNReshapeGradientClass
	_MPSNNReshapeGradientClassOnce sync.Once
)

func getMPSNNReshapeGradientClass() MPSNNReshapeGradientClass {
	_MPSNNReshapeGradientClassOnce.Do(func() {
		_MPSNNReshapeGradientClass = MPSNNReshapeGradientClass{class: objc.GetClass("MPSNNReshapeGradient")}
	})
	return _MPSNNReshapeGradientClass
}

// GetMPSNNReshapeGradientClass returns the class object for MPSNNReshapeGradient.
func GetMPSNNReshapeGradientClass() MPSNNReshapeGradientClass {
	return getMPSNNReshapeGradientClass()
}

type MPSNNReshapeGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReshapeGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReshapeGradientClass) Alloc() MPSNNReshapeGradient {
	rv := objc.Send[MPSNNReshapeGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradient
type MPSNNReshapeGradient struct {
	MPSCNNGradientKernel
}

// MPSNNReshapeGradientFromID constructs a [MPSNNReshapeGradient] from an objc.ID.
func MPSNNReshapeGradientFromID(id objc.ID) MPSNNReshapeGradient {
	return MPSNNReshapeGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSNNReshapeGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReshapeGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradient
type IMPSNNReshapeGradient interface {
	IMPSCNNGradientKernel
}

// Init initializes the instance.
func (r MPSNNReshapeGradient) Init() MPSNNReshapeGradient {
	rv := objc.Send[MPSNNReshapeGradient](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReshapeGradient) Autorelease() MPSNNReshapeGradient {
	rv := objc.Send[MPSNNReshapeGradient](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReshapeGradient creates a new MPSNNReshapeGradient instance.
func NewMPSNNReshapeGradient() MPSNNReshapeGradient {
	class := getMPSNNReshapeGradientClass()
	rv := objc.Send[MPSNNReshapeGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReshapeGradientWithCoder(aDecoder foundation.INSCoder) MPSNNReshapeGradient {
	instance := getMPSNNReshapeGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReshapeGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradient/init(coder:device:)
func NewReshapeGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReshapeGradient {
	instance := getMPSNNReshapeGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReshapeGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradient/init(device:)
func NewReshapeGradientWithDevice(device metal.MTLDevice) MPSNNReshapeGradient {
	instance := getMPSNNReshapeGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReshapeGradientFromID(rv)
}

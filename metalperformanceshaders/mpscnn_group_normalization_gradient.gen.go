// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNGroupNormalizationGradient] class.
var (
	_MPSCNNGroupNormalizationGradientClass     MPSCNNGroupNormalizationGradientClass
	_MPSCNNGroupNormalizationGradientClassOnce sync.Once
)

func getMPSCNNGroupNormalizationGradientClass() MPSCNNGroupNormalizationGradientClass {
	_MPSCNNGroupNormalizationGradientClassOnce.Do(func() {
		_MPSCNNGroupNormalizationGradientClass = MPSCNNGroupNormalizationGradientClass{class: objc.GetClass("MPSCNNGroupNormalizationGradient")}
	})
	return _MPSCNNGroupNormalizationGradientClass
}

// GetMPSCNNGroupNormalizationGradientClass returns the class object for MPSCNNGroupNormalizationGradient.
func GetMPSCNNGroupNormalizationGradientClass() MPSCNNGroupNormalizationGradientClass {
	return getMPSCNNGroupNormalizationGradientClass()
}

type MPSCNNGroupNormalizationGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNGroupNormalizationGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNGroupNormalizationGradientClass) Alloc() MPSCNNGroupNormalizationGradient {
	rv := objc.Send[MPSCNNGroupNormalizationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradient
type MPSCNNGroupNormalizationGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNGroupNormalizationGradientFromID constructs a [MPSCNNGroupNormalizationGradient] from an objc.ID.
func MPSCNNGroupNormalizationGradientFromID(id objc.ID) MPSCNNGroupNormalizationGradient {
	return MPSCNNGroupNormalizationGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNGroupNormalizationGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNGroupNormalizationGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradient
type IMPSCNNGroupNormalizationGradient interface {
	IMPSCNNGradientKernel
}

// Init initializes the instance.
func (c MPSCNNGroupNormalizationGradient) Init() MPSCNNGroupNormalizationGradient {
	rv := objc.Send[MPSCNNGroupNormalizationGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNGroupNormalizationGradient) Autorelease() MPSCNNGroupNormalizationGradient {
	rv := objc.Send[MPSCNNGroupNormalizationGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNGroupNormalizationGradient creates a new MPSCNNGroupNormalizationGradient instance.
func NewMPSCNNGroupNormalizationGradient() MPSCNNGroupNormalizationGradient {
	class := getMPSCNNGroupNormalizationGradientClass()
	rv := objc.Send[MPSCNNGroupNormalizationGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNGroupNormalizationGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNGroupNormalizationGradient {
	instance := getMPSCNNGroupNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNGroupNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(coder:device:)
func NewCNNGroupNormalizationGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNGroupNormalizationGradient {
	instance := getMPSCNNGroupNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNGroupNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNGroupNormalizationGradientWithDevice(device metal.MTLDevice) MPSCNNGroupNormalizationGradient {
	instance := getMPSCNNGroupNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNGroupNormalizationGradientFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNInstanceNormalizationGradient] class.
var (
	_MPSCNNInstanceNormalizationGradientClass     MPSCNNInstanceNormalizationGradientClass
	_MPSCNNInstanceNormalizationGradientClassOnce sync.Once
)

func getMPSCNNInstanceNormalizationGradientClass() MPSCNNInstanceNormalizationGradientClass {
	_MPSCNNInstanceNormalizationGradientClassOnce.Do(func() {
		_MPSCNNInstanceNormalizationGradientClass = MPSCNNInstanceNormalizationGradientClass{class: objc.GetClass("MPSCNNInstanceNormalizationGradient")}
	})
	return _MPSCNNInstanceNormalizationGradientClass
}

// GetMPSCNNInstanceNormalizationGradientClass returns the class object for MPSCNNInstanceNormalizationGradient.
func GetMPSCNNInstanceNormalizationGradientClass() MPSCNNInstanceNormalizationGradientClass {
	return getMPSCNNInstanceNormalizationGradientClass()
}

type MPSCNNInstanceNormalizationGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNInstanceNormalizationGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNInstanceNormalizationGradientClass) Alloc() MPSCNNInstanceNormalizationGradient {
	rv := objc.Send[MPSCNNInstanceNormalizationGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient instance normalization kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradient
type MPSCNNInstanceNormalizationGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNInstanceNormalizationGradientFromID constructs a [MPSCNNInstanceNormalizationGradient] from an objc.ID.
//
// A gradient instance normalization kernel.
func MPSCNNInstanceNormalizationGradientFromID(id objc.ID) MPSCNNInstanceNormalizationGradient {
	return MPSCNNInstanceNormalizationGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNInstanceNormalizationGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNInstanceNormalizationGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradient
type IMPSCNNInstanceNormalizationGradient interface {
	IMPSCNNGradientKernel
}

// Init initializes the instance.
func (c MPSCNNInstanceNormalizationGradient) Init() MPSCNNInstanceNormalizationGradient {
	rv := objc.Send[MPSCNNInstanceNormalizationGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNInstanceNormalizationGradient) Autorelease() MPSCNNInstanceNormalizationGradient {
	rv := objc.Send[MPSCNNInstanceNormalizationGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNInstanceNormalizationGradient creates a new MPSCNNInstanceNormalizationGradient instance.
func NewMPSCNNInstanceNormalizationGradient() MPSCNNInstanceNormalizationGradient {
	class := getMPSCNNInstanceNormalizationGradientClass()
	rv := objc.Send[MPSCNNInstanceNormalizationGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNInstanceNormalizationGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNInstanceNormalizationGradient {
	instance := getMPSCNNInstanceNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNInstanceNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(coder:device:)
func NewCNNInstanceNormalizationGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNInstanceNormalizationGradient {
	instance := getMPSCNNInstanceNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNInstanceNormalizationGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNInstanceNormalizationGradientWithDevice(device metal.MTLDevice) MPSCNNInstanceNormalizationGradient {
	instance := getMPSCNNInstanceNormalizationGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNInstanceNormalizationGradientFromID(rv)
}

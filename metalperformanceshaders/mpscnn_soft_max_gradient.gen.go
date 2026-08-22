// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSoftMaxGradient] class.
var (
	_MPSCNNSoftMaxGradientClass     MPSCNNSoftMaxGradientClass
	_MPSCNNSoftMaxGradientClassOnce sync.Once
)

func getMPSCNNSoftMaxGradientClass() MPSCNNSoftMaxGradientClass {
	_MPSCNNSoftMaxGradientClassOnce.Do(func() {
		_MPSCNNSoftMaxGradientClass = MPSCNNSoftMaxGradientClass{class: objc.GetClass("MPSCNNSoftMaxGradient")}
	})
	return _MPSCNNSoftMaxGradientClass
}

// GetMPSCNNSoftMaxGradientClass returns the class object for MPSCNNSoftMaxGradient.
func GetMPSCNNSoftMaxGradientClass() MPSCNNSoftMaxGradientClass {
	return getMPSCNNSoftMaxGradientClass()
}

type MPSCNNSoftMaxGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSoftMaxGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSoftMaxGradientClass) Alloc() MPSCNNSoftMaxGradient {
	rv := objc.Send[MPSCNNSoftMaxGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient softmax filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradient
type MPSCNNSoftMaxGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNSoftMaxGradientFromID constructs a [MPSCNNSoftMaxGradient] from an objc.ID.
//
// A gradient softmax filter.
func MPSCNNSoftMaxGradientFromID(id objc.ID) MPSCNNSoftMaxGradient {
	return MPSCNNSoftMaxGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNSoftMaxGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSoftMaxGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradient
type IMPSCNNSoftMaxGradient interface {
	IMPSCNNGradientKernel
}

// Init initializes the instance.
func (c MPSCNNSoftMaxGradient) Init() MPSCNNSoftMaxGradient {
	rv := objc.Send[MPSCNNSoftMaxGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSoftMaxGradient) Autorelease() MPSCNNSoftMaxGradient {
	rv := objc.Send[MPSCNNSoftMaxGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSoftMaxGradient creates a new MPSCNNSoftMaxGradient instance.
func NewMPSCNNSoftMaxGradient() MPSCNNSoftMaxGradient {
	class := getMPSCNNSoftMaxGradientClass()
	rv := objc.Send[MPSCNNSoftMaxGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNSoftMaxGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNSoftMaxGradient {
	instance := getMPSCNNSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNSoftMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradient/init(coder:device:)
func NewCNNSoftMaxGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNSoftMaxGradient {
	instance := getMPSCNNSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNSoftMaxGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradient/init(device:)
func NewCNNSoftMaxGradientWithDevice(device metal.MTLDevice) MPSCNNSoftMaxGradient {
	instance := getMPSCNNSoftMaxGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNSoftMaxGradientFromID(rv)
}

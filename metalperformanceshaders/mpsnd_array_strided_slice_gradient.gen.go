// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayStridedSliceGradient] class.
var (
	_MPSNDArrayStridedSliceGradientClass     MPSNDArrayStridedSliceGradientClass
	_MPSNDArrayStridedSliceGradientClassOnce sync.Once
)

func getMPSNDArrayStridedSliceGradientClass() MPSNDArrayStridedSliceGradientClass {
	_MPSNDArrayStridedSliceGradientClassOnce.Do(func() {
		_MPSNDArrayStridedSliceGradientClass = MPSNDArrayStridedSliceGradientClass{class: objc.GetClass("MPSNDArrayStridedSliceGradient")}
	})
	return _MPSNDArrayStridedSliceGradientClass
}

// GetMPSNDArrayStridedSliceGradientClass returns the class object for MPSNDArrayStridedSliceGradient.
func GetMPSNDArrayStridedSliceGradientClass() MPSNDArrayStridedSliceGradientClass {
	return getMPSNDArrayStridedSliceGradientClass()
}

type MPSNDArrayStridedSliceGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayStridedSliceGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayStridedSliceGradientClass) Alloc() MPSNDArrayStridedSliceGradient {
	rv := objc.Send[MPSNDArrayStridedSliceGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSliceGradient
type MPSNDArrayStridedSliceGradient struct {
	MPSNDArrayUnaryGradientKernel
}

// MPSNDArrayStridedSliceGradientFromID constructs a [MPSNDArrayStridedSliceGradient] from an objc.ID.
func MPSNDArrayStridedSliceGradientFromID(id objc.ID) MPSNDArrayStridedSliceGradient {
	return MPSNDArrayStridedSliceGradient{MPSNDArrayUnaryGradientKernel: MPSNDArrayUnaryGradientKernelFromID(id)}
}

// NOTE: MPSNDArrayStridedSliceGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayStridedSliceGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSliceGradient
type IMPSNDArrayStridedSliceGradient interface {
	IMPSNDArrayUnaryGradientKernel
}

// Init initializes the instance.
func (n MPSNDArrayStridedSliceGradient) Init() MPSNDArrayStridedSliceGradient {
	rv := objc.Send[MPSNDArrayStridedSliceGradient](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayStridedSliceGradient) Autorelease() MPSNDArrayStridedSliceGradient {
	rv := objc.Send[MPSNDArrayStridedSliceGradient](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayStridedSliceGradient creates a new MPSNDArrayStridedSliceGradient instance.
func NewMPSNDArrayStridedSliceGradient() MPSNDArrayStridedSliceGradient {
	class := getMPSNDArrayStridedSliceGradientClass()
	rv := objc.Send[MPSNDArrayStridedSliceGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayStridedSliceGradientWithCoder(aDecoder foundation.INSCoder) MPSNDArrayStridedSliceGradient {
	instance := getMPSNDArrayStridedSliceGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayStridedSliceGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel/init(coder:device:)
func NewNDArrayStridedSliceGradientWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayStridedSliceGradient {
	instance := getMPSNDArrayStridedSliceGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayStridedSliceGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel/init(device:)
func NewNDArrayStridedSliceGradientWithDevice(device metal.MTLDevice) MPSNDArrayStridedSliceGradient {
	instance := getMPSNDArrayStridedSliceGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayStridedSliceGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/init(device:sourceCount:)
func NewNDArrayStridedSliceGradientWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayStridedSliceGradient {
	instance := getMPSNDArrayStridedSliceGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayStridedSliceGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/init(device:sourceCount:sourceGradientIndex:)
func NewNDArrayStridedSliceGradientWithDeviceSourceCountSourceGradientIndex(device metal.MTLDevice, count uint, sourceGradientIndex uint) MPSNDArrayStridedSliceGradient {
	instance := getMPSNDArrayStridedSliceGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), device, count, sourceGradientIndex)
	return MPSNDArrayStridedSliceGradientFromID(rv)
}

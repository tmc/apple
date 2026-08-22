// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayGatherGradient] class.
var (
	_MPSNDArrayGatherGradientClass     MPSNDArrayGatherGradientClass
	_MPSNDArrayGatherGradientClassOnce sync.Once
)

func getMPSNDArrayGatherGradientClass() MPSNDArrayGatherGradientClass {
	_MPSNDArrayGatherGradientClassOnce.Do(func() {
		_MPSNDArrayGatherGradientClass = MPSNDArrayGatherGradientClass{class: objc.GetClass("MPSNDArrayGatherGradient")}
	})
	return _MPSNDArrayGatherGradientClass
}

// GetMPSNDArrayGatherGradientClass returns the class object for MPSNDArrayGatherGradient.
func GetMPSNDArrayGatherGradientClass() MPSNDArrayGatherGradientClass {
	return getMPSNDArrayGatherGradientClass()
}

type MPSNDArrayGatherGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayGatherGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayGatherGradientClass) Alloc() MPSNDArrayGatherGradient {
	rv := objc.Send[MPSNDArrayGatherGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGatherGradient
type MPSNDArrayGatherGradient struct {
	MPSNDArrayBinaryPrimaryGradientKernel
}

// MPSNDArrayGatherGradientFromID constructs a [MPSNDArrayGatherGradient] from an objc.ID.
func MPSNDArrayGatherGradientFromID(id objc.ID) MPSNDArrayGatherGradient {
	return MPSNDArrayGatherGradient{MPSNDArrayBinaryPrimaryGradientKernel: MPSNDArrayBinaryPrimaryGradientKernelFromID(id)}
}

// NOTE: MPSNDArrayGatherGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayGatherGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGatherGradient
type IMPSNDArrayGatherGradient interface {
	IMPSNDArrayBinaryPrimaryGradientKernel
}

// Init initializes the instance.
func (n MPSNDArrayGatherGradient) Init() MPSNDArrayGatherGradient {
	rv := objc.Send[MPSNDArrayGatherGradient](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayGatherGradient) Autorelease() MPSNDArrayGatherGradient {
	rv := objc.Send[MPSNDArrayGatherGradient](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayGatherGradient creates a new MPSNDArrayGatherGradient instance.
func NewMPSNDArrayGatherGradient() MPSNDArrayGatherGradient {
	class := getMPSNDArrayGatherGradientClass()
	rv := objc.Send[MPSNDArrayGatherGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayGatherGradientWithCoder(aDecoder foundation.INSCoder) MPSNDArrayGatherGradient {
	instance := getMPSNDArrayGatherGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayGatherGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel/init(coder:device:)
func NewNDArrayGatherGradientWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayGatherGradient {
	instance := getMPSNDArrayGatherGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayGatherGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel/init(device:)
func NewNDArrayGatherGradientWithDevice(device metal.MTLDevice) MPSNDArrayGatherGradient {
	instance := getMPSNDArrayGatherGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayGatherGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/init(device:sourceCount:)
func NewNDArrayGatherGradientWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayGatherGradient {
	instance := getMPSNDArrayGatherGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayGatherGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel/init(device:sourceCount:sourceGradientIndex:)
func NewNDArrayGatherGradientWithDeviceSourceCountSourceGradientIndex(device metal.MTLDevice, count uint, sourceGradientIndex uint) MPSNDArrayGatherGradient {
	instance := getMPSNDArrayGatherGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), device, count, sourceGradientIndex)
	return MPSNDArrayGatherGradientFromID(rv)
}

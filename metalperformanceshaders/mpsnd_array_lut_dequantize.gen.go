// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayLUTDequantize] class.
var (
	_MPSNDArrayLUTDequantizeClass     MPSNDArrayLUTDequantizeClass
	_MPSNDArrayLUTDequantizeClassOnce sync.Once
)

func getMPSNDArrayLUTDequantizeClass() MPSNDArrayLUTDequantizeClass {
	_MPSNDArrayLUTDequantizeClassOnce.Do(func() {
		_MPSNDArrayLUTDequantizeClass = MPSNDArrayLUTDequantizeClass{class: objc.GetClass("MPSNDArrayLUTDequantize")}
	})
	return _MPSNDArrayLUTDequantizeClass
}

// GetMPSNDArrayLUTDequantizeClass returns the class object for MPSNDArrayLUTDequantize.
func GetMPSNDArrayLUTDequantizeClass() MPSNDArrayLUTDequantizeClass {
	return getMPSNDArrayLUTDequantizeClass()
}

type MPSNDArrayLUTDequantizeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayLUTDequantizeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayLUTDequantizeClass) Alloc() MPSNDArrayLUTDequantize {
	rv := objc.Send[MPSNDArrayLUTDequantize](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTDequantize
type MPSNDArrayLUTDequantize struct {
	MPSNDArrayMultiaryKernel
}

// MPSNDArrayLUTDequantizeFromID constructs a [MPSNDArrayLUTDequantize] from an objc.ID.
func MPSNDArrayLUTDequantizeFromID(id objc.ID) MPSNDArrayLUTDequantize {
	return MPSNDArrayLUTDequantize{MPSNDArrayMultiaryKernel: MPSNDArrayMultiaryKernelFromID(id)}
}

// NOTE: MPSNDArrayLUTDequantize adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayLUTDequantize] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTDequantize
type IMPSNDArrayLUTDequantize interface {
	IMPSNDArrayMultiaryKernel
}

// Init initializes the instance.
func (n MPSNDArrayLUTDequantize) Init() MPSNDArrayLUTDequantize {
	rv := objc.Send[MPSNDArrayLUTDequantize](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayLUTDequantize) Autorelease() MPSNDArrayLUTDequantize {
	rv := objc.Send[MPSNDArrayLUTDequantize](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayLUTDequantize creates a new MPSNDArrayLUTDequantize instance.
func NewMPSNDArrayLUTDequantize() MPSNDArrayLUTDequantize {
	class := getMPSNDArrayLUTDequantizeClass()
	rv := objc.Send[MPSNDArrayLUTDequantize](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayLUTDequantizeWithCoder(aDecoder foundation.INSCoder) MPSNDArrayLUTDequantize {
	instance := getMPSNDArrayLUTDequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayLUTDequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(coder:device:)
func NewNDArrayLUTDequantizeWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayLUTDequantize {
	instance := getMPSNDArrayLUTDequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayLUTDequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTDequantize/init(device:)
func NewNDArrayLUTDequantizeWithDevice(device metal.MTLDevice) MPSNDArrayLUTDequantize {
	instance := getMPSNDArrayLUTDequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayLUTDequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayLUTDequantizeWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayLUTDequantize {
	instance := getMPSNDArrayLUTDequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayLUTDequantizeFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayAffineInt4Dequantize] class.
var (
	_MPSNDArrayAffineInt4DequantizeClass     MPSNDArrayAffineInt4DequantizeClass
	_MPSNDArrayAffineInt4DequantizeClassOnce sync.Once
)

func getMPSNDArrayAffineInt4DequantizeClass() MPSNDArrayAffineInt4DequantizeClass {
	_MPSNDArrayAffineInt4DequantizeClassOnce.Do(func() {
		_MPSNDArrayAffineInt4DequantizeClass = MPSNDArrayAffineInt4DequantizeClass{class: objc.GetClass("MPSNDArrayAffineInt4Dequantize")}
	})
	return _MPSNDArrayAffineInt4DequantizeClass
}

// GetMPSNDArrayAffineInt4DequantizeClass returns the class object for MPSNDArrayAffineInt4Dequantize.
func GetMPSNDArrayAffineInt4DequantizeClass() MPSNDArrayAffineInt4DequantizeClass {
	return getMPSNDArrayAffineInt4DequantizeClass()
}

type MPSNDArrayAffineInt4DequantizeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayAffineInt4DequantizeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayAffineInt4DequantizeClass) Alloc() MPSNDArrayAffineInt4Dequantize {
	rv := objc.Send[MPSNDArrayAffineInt4Dequantize](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNDArrayAffineInt4Dequantize.InitWithDeviceQuantizationDescriptor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineInt4Dequantize
type MPSNDArrayAffineInt4Dequantize struct {
	MPSNDArrayMultiaryKernel
}

// MPSNDArrayAffineInt4DequantizeFromID constructs a [MPSNDArrayAffineInt4Dequantize] from an objc.ID.
func MPSNDArrayAffineInt4DequantizeFromID(id objc.ID) MPSNDArrayAffineInt4Dequantize {
	return MPSNDArrayAffineInt4Dequantize{MPSNDArrayMultiaryKernel: MPSNDArrayMultiaryKernelFromID(id)}
}

// NOTE: MPSNDArrayAffineInt4Dequantize adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayAffineInt4Dequantize] class.
//
// # Initializers
//
//   - [IMPSNDArrayAffineInt4Dequantize.InitWithDeviceQuantizationDescriptor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineInt4Dequantize
type IMPSNDArrayAffineInt4Dequantize interface {
	IMPSNDArrayMultiaryKernel

	// Topic: Initializers

	InitWithDeviceQuantizationDescriptor(device metal.MTLDevice, quantizationDescriptor IMPSNDArrayAffineQuantizationDescriptor) MPSNDArrayAffineInt4Dequantize
}

// Init initializes the instance.
func (n MPSNDArrayAffineInt4Dequantize) Init() MPSNDArrayAffineInt4Dequantize {
	rv := objc.Send[MPSNDArrayAffineInt4Dequantize](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayAffineInt4Dequantize) Autorelease() MPSNDArrayAffineInt4Dequantize {
	rv := objc.Send[MPSNDArrayAffineInt4Dequantize](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayAffineInt4Dequantize creates a new MPSNDArrayAffineInt4Dequantize instance.
func NewMPSNDArrayAffineInt4Dequantize() MPSNDArrayAffineInt4Dequantize {
	class := getMPSNDArrayAffineInt4DequantizeClass()
	rv := objc.Send[MPSNDArrayAffineInt4Dequantize](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayAffineInt4DequantizeWithCoder(aDecoder foundation.INSCoder) MPSNDArrayAffineInt4Dequantize {
	instance := getMPSNDArrayAffineInt4DequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayAffineInt4DequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(coder:device:)
func NewNDArrayAffineInt4DequantizeWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayAffineInt4Dequantize {
	instance := getMPSNDArrayAffineInt4DequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayAffineInt4DequantizeFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewNDArrayAffineInt4DequantizeWithDevice(device metal.MTLDevice) MPSNDArrayAffineInt4Dequantize {
	instance := getMPSNDArrayAffineInt4DequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayAffineInt4DequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineInt4Dequantize/init(device:quantizationDescriptor:)
func NewNDArrayAffineInt4DequantizeWithDeviceQuantizationDescriptor(device metal.MTLDevice, quantizationDescriptor IMPSNDArrayAffineQuantizationDescriptor) MPSNDArrayAffineInt4Dequantize {
	instance := getMPSNDArrayAffineInt4DequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:quantizationDescriptor:"), device, quantizationDescriptor)
	return MPSNDArrayAffineInt4DequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayAffineInt4DequantizeWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayAffineInt4Dequantize {
	instance := getMPSNDArrayAffineInt4DequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayAffineInt4DequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineInt4Dequantize/init(device:quantizationDescriptor:)
func (n MPSNDArrayAffineInt4Dequantize) InitWithDeviceQuantizationDescriptor(device metal.MTLDevice, quantizationDescriptor IMPSNDArrayAffineQuantizationDescriptor) MPSNDArrayAffineInt4Dequantize {
	rv := objc.Send[MPSNDArrayAffineInt4Dequantize](n.ID, objc.Sel("initWithDevice:quantizationDescriptor:"), device, quantizationDescriptor)
	return rv
}

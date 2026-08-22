// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayVectorLUTDequantize] class.
var (
	_MPSNDArrayVectorLUTDequantizeClass     MPSNDArrayVectorLUTDequantizeClass
	_MPSNDArrayVectorLUTDequantizeClassOnce sync.Once
)

func getMPSNDArrayVectorLUTDequantizeClass() MPSNDArrayVectorLUTDequantizeClass {
	_MPSNDArrayVectorLUTDequantizeClassOnce.Do(func() {
		_MPSNDArrayVectorLUTDequantizeClass = MPSNDArrayVectorLUTDequantizeClass{class: objc.GetClass("MPSNDArrayVectorLUTDequantize")}
	})
	return _MPSNDArrayVectorLUTDequantizeClass
}

// GetMPSNDArrayVectorLUTDequantizeClass returns the class object for MPSNDArrayVectorLUTDequantize.
func GetMPSNDArrayVectorLUTDequantizeClass() MPSNDArrayVectorLUTDequantizeClass {
	return getMPSNDArrayVectorLUTDequantizeClass()
}

type MPSNDArrayVectorLUTDequantizeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayVectorLUTDequantizeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayVectorLUTDequantizeClass) Alloc() MPSNDArrayVectorLUTDequantize {
	rv := objc.Send[MPSNDArrayVectorLUTDequantize](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNDArrayVectorLUTDequantize.InitWithDeviceAxis]
//
// # Instance Properties
//
//   - [MPSNDArrayVectorLUTDequantize.VectorAxis]
//   - [MPSNDArrayVectorLUTDequantize.SetVectorAxis]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayVectorLUTDequantize
type MPSNDArrayVectorLUTDequantize struct {
	MPSNDArrayMultiaryKernel
}

// MPSNDArrayVectorLUTDequantizeFromID constructs a [MPSNDArrayVectorLUTDequantize] from an objc.ID.
func MPSNDArrayVectorLUTDequantizeFromID(id objc.ID) MPSNDArrayVectorLUTDequantize {
	return MPSNDArrayVectorLUTDequantize{MPSNDArrayMultiaryKernel: MPSNDArrayMultiaryKernelFromID(id)}
}

// NOTE: MPSNDArrayVectorLUTDequantize adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayVectorLUTDequantize] class.
//
// # Initializers
//
//   - [IMPSNDArrayVectorLUTDequantize.InitWithDeviceAxis]
//
// # Instance Properties
//
//   - [IMPSNDArrayVectorLUTDequantize.VectorAxis]
//   - [IMPSNDArrayVectorLUTDequantize.SetVectorAxis]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayVectorLUTDequantize
type IMPSNDArrayVectorLUTDequantize interface {
	IMPSNDArrayMultiaryKernel

	// Topic: Initializers

	InitWithDeviceAxis(device metal.MTLDevice, axis uint) MPSNDArrayVectorLUTDequantize

	// Topic: Instance Properties

	VectorAxis() uint
	SetVectorAxis(value uint)
}

// Init initializes the instance.
func (n MPSNDArrayVectorLUTDequantize) Init() MPSNDArrayVectorLUTDequantize {
	rv := objc.Send[MPSNDArrayVectorLUTDequantize](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayVectorLUTDequantize) Autorelease() MPSNDArrayVectorLUTDequantize {
	rv := objc.Send[MPSNDArrayVectorLUTDequantize](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayVectorLUTDequantize creates a new MPSNDArrayVectorLUTDequantize instance.
func NewMPSNDArrayVectorLUTDequantize() MPSNDArrayVectorLUTDequantize {
	class := getMPSNDArrayVectorLUTDequantizeClass()
	rv := objc.Send[MPSNDArrayVectorLUTDequantize](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayVectorLUTDequantizeWithCoder(aDecoder foundation.INSCoder) MPSNDArrayVectorLUTDequantize {
	instance := getMPSNDArrayVectorLUTDequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayVectorLUTDequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(coder:device:)
func NewNDArrayVectorLUTDequantizeWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayVectorLUTDequantize {
	instance := getMPSNDArrayVectorLUTDequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayVectorLUTDequantizeFromID(rv)
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
func NewNDArrayVectorLUTDequantizeWithDevice(device metal.MTLDevice) MPSNDArrayVectorLUTDequantize {
	instance := getMPSNDArrayVectorLUTDequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayVectorLUTDequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayVectorLUTDequantize/init(device:axis:)
func NewNDArrayVectorLUTDequantizeWithDeviceAxis(device metal.MTLDevice, axis uint) MPSNDArrayVectorLUTDequantize {
	instance := getMPSNDArrayVectorLUTDequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:axis:"), device, axis)
	return MPSNDArrayVectorLUTDequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayVectorLUTDequantizeWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayVectorLUTDequantize {
	instance := getMPSNDArrayVectorLUTDequantizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayVectorLUTDequantizeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayVectorLUTDequantize/init(device:axis:)
func (n MPSNDArrayVectorLUTDequantize) InitWithDeviceAxis(device metal.MTLDevice, axis uint) MPSNDArrayVectorLUTDequantize {
	rv := objc.Send[MPSNDArrayVectorLUTDequantize](n.ID, objc.Sel("initWithDevice:axis:"), device, axis)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayVectorLUTDequantize/vectorAxis
func (n MPSNDArrayVectorLUTDequantize) VectorAxis() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("vectorAxis"))
	return rv
}
func (n MPSNDArrayVectorLUTDequantize) SetVectorAxis(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setVectorAxis:"), value)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayQuantizedMatrixMultiplication] class.
var (
	_MPSNDArrayQuantizedMatrixMultiplicationClass     MPSNDArrayQuantizedMatrixMultiplicationClass
	_MPSNDArrayQuantizedMatrixMultiplicationClassOnce sync.Once
)

func getMPSNDArrayQuantizedMatrixMultiplicationClass() MPSNDArrayQuantizedMatrixMultiplicationClass {
	_MPSNDArrayQuantizedMatrixMultiplicationClassOnce.Do(func() {
		_MPSNDArrayQuantizedMatrixMultiplicationClass = MPSNDArrayQuantizedMatrixMultiplicationClass{class: objc.GetClass("MPSNDArrayQuantizedMatrixMultiplication")}
	})
	return _MPSNDArrayQuantizedMatrixMultiplicationClass
}

// GetMPSNDArrayQuantizedMatrixMultiplicationClass returns the class object for MPSNDArrayQuantizedMatrixMultiplication.
func GetMPSNDArrayQuantizedMatrixMultiplicationClass() MPSNDArrayQuantizedMatrixMultiplicationClass {
	return getMPSNDArrayQuantizedMatrixMultiplicationClass()
}

type MPSNDArrayQuantizedMatrixMultiplicationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayQuantizedMatrixMultiplicationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayQuantizedMatrixMultiplicationClass) Alloc() MPSNDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[MPSNDArrayQuantizedMatrixMultiplication](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNDArrayQuantizedMatrixMultiplication.InitWithDeviceLeftQuantizationDescriptorRightQuantizationDescriptor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizedMatrixMultiplication
type MPSNDArrayQuantizedMatrixMultiplication struct {
	MPSNDArrayMatrixMultiplication
}

// MPSNDArrayQuantizedMatrixMultiplicationFromID constructs a [MPSNDArrayQuantizedMatrixMultiplication] from an objc.ID.
func MPSNDArrayQuantizedMatrixMultiplicationFromID(id objc.ID) MPSNDArrayQuantizedMatrixMultiplication {
	return MPSNDArrayQuantizedMatrixMultiplication{MPSNDArrayMatrixMultiplication: MPSNDArrayMatrixMultiplicationFromID(id)}
}

// NOTE: MPSNDArrayQuantizedMatrixMultiplication adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayQuantizedMatrixMultiplication] class.
//
// # Initializers
//
//   - [IMPSNDArrayQuantizedMatrixMultiplication.InitWithDeviceLeftQuantizationDescriptorRightQuantizationDescriptor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizedMatrixMultiplication
type IMPSNDArrayQuantizedMatrixMultiplication interface {
	IMPSNDArrayMatrixMultiplication

	// Topic: Initializers

	InitWithDeviceLeftQuantizationDescriptorRightQuantizationDescriptor(device metal.MTLDevice, leftQuantizationDescriptor IMPSNDArrayQuantizationDescriptor, rightQuantizationDescriptor IMPSNDArrayQuantizationDescriptor) MPSNDArrayQuantizedMatrixMultiplication
}

// Init initializes the instance.
func (n MPSNDArrayQuantizedMatrixMultiplication) Init() MPSNDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[MPSNDArrayQuantizedMatrixMultiplication](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayQuantizedMatrixMultiplication) Autorelease() MPSNDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[MPSNDArrayQuantizedMatrixMultiplication](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayQuantizedMatrixMultiplication creates a new MPSNDArrayQuantizedMatrixMultiplication instance.
func NewMPSNDArrayQuantizedMatrixMultiplication() MPSNDArrayQuantizedMatrixMultiplication {
	class := getMPSNDArrayQuantizedMatrixMultiplicationClass()
	rv := objc.Send[MPSNDArrayQuantizedMatrixMultiplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayQuantizedMatrixMultiplicationWithCoder(aDecoder foundation.INSCoder) MPSNDArrayQuantizedMatrixMultiplication {
	instance := getMPSNDArrayQuantizedMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayQuantizedMatrixMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(coder:device:)
func NewNDArrayQuantizedMatrixMultiplicationWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayQuantizedMatrixMultiplication {
	instance := getMPSNDArrayQuantizedMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayQuantizedMatrixMultiplicationFromID(rv)
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
func NewNDArrayQuantizedMatrixMultiplicationWithDevice(device metal.MTLDevice) MPSNDArrayQuantizedMatrixMultiplication {
	instance := getMPSNDArrayQuantizedMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayQuantizedMatrixMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizedMatrixMultiplication/init(device:leftQuantizationDescriptor:rightQuantizationDescriptor:)
func NewNDArrayQuantizedMatrixMultiplicationWithDeviceLeftQuantizationDescriptorRightQuantizationDescriptor(device metal.MTLDevice, leftQuantizationDescriptor IMPSNDArrayQuantizationDescriptor, rightQuantizationDescriptor IMPSNDArrayQuantizationDescriptor) MPSNDArrayQuantizedMatrixMultiplication {
	instance := getMPSNDArrayQuantizedMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:leftQuantizationDescriptor:rightQuantizationDescriptor:"), device, leftQuantizationDescriptor, rightQuantizationDescriptor)
	return MPSNDArrayQuantizedMatrixMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayQuantizedMatrixMultiplicationWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayQuantizedMatrixMultiplication {
	instance := getMPSNDArrayQuantizedMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayQuantizedMatrixMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizedMatrixMultiplication/init(device:leftQuantizationDescriptor:rightQuantizationDescriptor:)
func (n MPSNDArrayQuantizedMatrixMultiplication) InitWithDeviceLeftQuantizationDescriptorRightQuantizationDescriptor(device metal.MTLDevice, leftQuantizationDescriptor IMPSNDArrayQuantizationDescriptor, rightQuantizationDescriptor IMPSNDArrayQuantizationDescriptor) MPSNDArrayQuantizedMatrixMultiplication {
	rv := objc.Send[MPSNDArrayQuantizedMatrixMultiplication](n.ID, objc.Sel("initWithDevice:leftQuantizationDescriptor:rightQuantizationDescriptor:"), device, leftQuantizationDescriptor, rightQuantizationDescriptor)
	return rv
}

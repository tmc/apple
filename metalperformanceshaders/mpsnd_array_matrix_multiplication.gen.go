// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayMatrixMultiplication] class.
var (
	_MPSNDArrayMatrixMultiplicationClass     MPSNDArrayMatrixMultiplicationClass
	_MPSNDArrayMatrixMultiplicationClassOnce sync.Once
)

func getMPSNDArrayMatrixMultiplicationClass() MPSNDArrayMatrixMultiplicationClass {
	_MPSNDArrayMatrixMultiplicationClassOnce.Do(func() {
		_MPSNDArrayMatrixMultiplicationClass = MPSNDArrayMatrixMultiplicationClass{class: objc.GetClass("MPSNDArrayMatrixMultiplication")}
	})
	return _MPSNDArrayMatrixMultiplicationClass
}

// GetMPSNDArrayMatrixMultiplicationClass returns the class object for MPSNDArrayMatrixMultiplication.
func GetMPSNDArrayMatrixMultiplicationClass() MPSNDArrayMatrixMultiplicationClass {
	return getMPSNDArrayMatrixMultiplicationClass()
}

type MPSNDArrayMatrixMultiplicationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayMatrixMultiplicationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayMatrixMultiplicationClass) Alloc() MPSNDArrayMatrixMultiplication {
	rv := objc.Send[MPSNDArrayMatrixMultiplication](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSNDArrayMatrixMultiplication.Alpha]
//   - [MPSNDArrayMatrixMultiplication.SetAlpha]
//   - [MPSNDArrayMatrixMultiplication.Beta]
//   - [MPSNDArrayMatrixMultiplication.SetBeta]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMatrixMultiplication
type MPSNDArrayMatrixMultiplication struct {
	MPSNDArrayMultiaryKernel
}

// MPSNDArrayMatrixMultiplicationFromID constructs a [MPSNDArrayMatrixMultiplication] from an objc.ID.
func MPSNDArrayMatrixMultiplicationFromID(id objc.ID) MPSNDArrayMatrixMultiplication {
	return MPSNDArrayMatrixMultiplication{MPSNDArrayMultiaryKernel: MPSNDArrayMultiaryKernelFromID(id)}
}

// NOTE: MPSNDArrayMatrixMultiplication adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayMatrixMultiplication] class.
//
// # Instance Properties
//
//   - [IMPSNDArrayMatrixMultiplication.Alpha]
//   - [IMPSNDArrayMatrixMultiplication.SetAlpha]
//   - [IMPSNDArrayMatrixMultiplication.Beta]
//   - [IMPSNDArrayMatrixMultiplication.SetBeta]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMatrixMultiplication
type IMPSNDArrayMatrixMultiplication interface {
	IMPSNDArrayMultiaryKernel

	// Topic: Instance Properties

	Alpha() float64
	SetAlpha(value float64)
	Beta() float64
	SetBeta(value float64)
}

// Init initializes the instance.
func (n MPSNDArrayMatrixMultiplication) Init() MPSNDArrayMatrixMultiplication {
	rv := objc.Send[MPSNDArrayMatrixMultiplication](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayMatrixMultiplication) Autorelease() MPSNDArrayMatrixMultiplication {
	rv := objc.Send[MPSNDArrayMatrixMultiplication](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayMatrixMultiplication creates a new MPSNDArrayMatrixMultiplication instance.
func NewMPSNDArrayMatrixMultiplication() MPSNDArrayMatrixMultiplication {
	class := getMPSNDArrayMatrixMultiplicationClass()
	rv := objc.Send[MPSNDArrayMatrixMultiplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayMatrixMultiplicationWithCoder(aDecoder foundation.INSCoder) MPSNDArrayMatrixMultiplication {
	instance := getMPSNDArrayMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayMatrixMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(coder:device:)
func NewNDArrayMatrixMultiplicationWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayMatrixMultiplication {
	instance := getMPSNDArrayMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayMatrixMultiplicationFromID(rv)
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
func NewNDArrayMatrixMultiplicationWithDevice(device metal.MTLDevice) MPSNDArrayMatrixMultiplication {
	instance := getMPSNDArrayMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayMatrixMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayMatrixMultiplicationWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayMatrixMultiplication {
	instance := getMPSNDArrayMatrixMultiplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayMatrixMultiplicationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMatrixMultiplication/alpha
func (n MPSNDArrayMatrixMultiplication) Alpha() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("alpha"))
	return rv
}
func (n MPSNDArrayMatrixMultiplication) SetAlpha(value float64) {
	objc.Send[struct{}](n.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMatrixMultiplication/beta
func (n MPSNDArrayMatrixMultiplication) Beta() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("beta"))
	return rv
}
func (n MPSNDArrayMatrixMultiplication) SetBeta(value float64) {
	objc.Send[struct{}](n.ID, objc.Sel("setBeta:"), value)
}

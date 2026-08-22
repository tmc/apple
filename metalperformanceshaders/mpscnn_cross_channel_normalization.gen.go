// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNCrossChannelNormalization] class.
var (
	_MPSCNNCrossChannelNormalizationClass     MPSCNNCrossChannelNormalizationClass
	_MPSCNNCrossChannelNormalizationClassOnce sync.Once
)

func getMPSCNNCrossChannelNormalizationClass() MPSCNNCrossChannelNormalizationClass {
	_MPSCNNCrossChannelNormalizationClassOnce.Do(func() {
		_MPSCNNCrossChannelNormalizationClass = MPSCNNCrossChannelNormalizationClass{class: objc.GetClass("MPSCNNCrossChannelNormalization")}
	})
	return _MPSCNNCrossChannelNormalizationClass
}

// GetMPSCNNCrossChannelNormalizationClass returns the class object for MPSCNNCrossChannelNormalization.
func GetMPSCNNCrossChannelNormalizationClass() MPSCNNCrossChannelNormalizationClass {
	return getMPSCNNCrossChannelNormalizationClass()
}

type MPSCNNCrossChannelNormalizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNCrossChannelNormalizationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNCrossChannelNormalizationClass) Alloc() MPSCNNCrossChannelNormalization {
	rv := objc.Send[MPSCNNCrossChannelNormalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A normalization kernel applied across feature channels.
//
// # Overview
//
// The normalization kernel applies the kernel to a local region across nearby
// feature channels, but with no spatial extent (i.e., they have the shape
// `kernel size x 1 x 1`). The normalized output is given by the function:
//
// [media-2903526]
//
// Where the normalizing factor is:
//
// [media-2903524]
//
// Where [N] is the kernel size. The window `Q(k)` itself is defined as:
//
// [media-2903527]
//
// Where `k` is the feature channel index (running from 0 to `D-1`) and [D] is
// the number of feature channels, and the values of
// [MPSCNNCrossChannelNormalization.Alpha],
// [MPSCNNCrossChannelNormalization.Beta], and
// [MPSCNNCrossChannelNormalization.Delta] are set via properties.
//
// It is your responsibility to ensure that the combination of the values of
// the [MPSCNNCrossChannelNormalization.Delta] and
// [MPSCNNCrossChannelNormalization.Alpha] properties does not result in a
// situation where the denominator becomes zero - in such situations the
// resulting pixel-value is undefined.
//
// # Initializers
//
//   - [MPSCNNCrossChannelNormalization.InitWithDeviceKernelSize]: Initializes a normalization kernel in a channel.
//
// # Instance Properties
//
//   - [MPSCNNCrossChannelNormalization.Alpha]: The “alpha” variable of the kernel function.
//   - [MPSCNNCrossChannelNormalization.SetAlpha]
//   - [MPSCNNCrossChannelNormalization.Beta]: The “beta” variable of the kernel function.
//   - [MPSCNNCrossChannelNormalization.SetBeta]
//   - [MPSCNNCrossChannelNormalization.Delta]: The “delta” variable of the kernel function.
//   - [MPSCNNCrossChannelNormalization.SetDelta]
//   - [MPSCNNCrossChannelNormalization.KernelSize]: The size of the square kernel window.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization
type MPSCNNCrossChannelNormalization struct {
	MPSCNNKernel
}

// MPSCNNCrossChannelNormalizationFromID constructs a [MPSCNNCrossChannelNormalization] from an objc.ID.
//
// A normalization kernel applied across feature channels.
func MPSCNNCrossChannelNormalizationFromID(id objc.ID) MPSCNNCrossChannelNormalization {
	return MPSCNNCrossChannelNormalization{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNCrossChannelNormalization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNCrossChannelNormalization] class.
//
// # Initializers
//
//   - [IMPSCNNCrossChannelNormalization.InitWithDeviceKernelSize]: Initializes a normalization kernel in a channel.
//
// # Instance Properties
//
//   - [IMPSCNNCrossChannelNormalization.Alpha]: The “alpha” variable of the kernel function.
//   - [IMPSCNNCrossChannelNormalization.SetAlpha]
//   - [IMPSCNNCrossChannelNormalization.Beta]: The “beta” variable of the kernel function.
//   - [IMPSCNNCrossChannelNormalization.SetBeta]
//   - [IMPSCNNCrossChannelNormalization.Delta]: The “delta” variable of the kernel function.
//   - [IMPSCNNCrossChannelNormalization.SetDelta]
//   - [IMPSCNNCrossChannelNormalization.KernelSize]: The size of the square kernel window.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization
type IMPSCNNCrossChannelNormalization interface {
	IMPSCNNKernel

	// Topic: Initializers

	// Initializes a normalization kernel in a channel.
	InitWithDeviceKernelSize(device metal.MTLDevice, kernelSize uint) MPSCNNCrossChannelNormalization

	// Topic: Instance Properties

	// The “alpha” variable of the kernel function.
	Alpha() float32
	SetAlpha(value float32)
	// The “beta” variable of the kernel function.
	Beta() float32
	SetBeta(value float32)
	// The “delta” variable of the kernel function.
	Delta() float32
	SetDelta(value float32)
	// The size of the square kernel window.
	KernelSize() uint
}

// Init initializes the instance.
func (c MPSCNNCrossChannelNormalization) Init() MPSCNNCrossChannelNormalization {
	rv := objc.Send[MPSCNNCrossChannelNormalization](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNCrossChannelNormalization) Autorelease() MPSCNNCrossChannelNormalization {
	rv := objc.Send[MPSCNNCrossChannelNormalization](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNCrossChannelNormalization creates a new MPSCNNCrossChannelNormalization instance.
func NewMPSCNNCrossChannelNormalization() MPSCNNCrossChannelNormalization {
	class := getMPSCNNCrossChannelNormalizationClass()
	rv := objc.Send[MPSCNNCrossChannelNormalization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNCrossChannelNormalizationWithCoder(aDecoder foundation.INSCoder) MPSCNNCrossChannelNormalization {
	instance := getMPSCNNCrossChannelNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNCrossChannelNormalizationFromID(rv)
}

// Initializes a normalization kernel in a channel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization/init(coder:device:)
func NewCNNCrossChannelNormalizationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNCrossChannelNormalization {
	instance := getMPSCNNCrossChannelNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNCrossChannelNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNCrossChannelNormalizationWithDevice(device metal.MTLDevice) MPSCNNCrossChannelNormalization {
	instance := getMPSCNNCrossChannelNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNCrossChannelNormalizationFromID(rv)
}

// Initializes a normalization kernel in a channel.
//
// device: The device the filter will run on.
//
// kernelSize: The size of the kernel, in both x and y dimensions.
//
// # Return Value
//
// A valid [MPSCNNCrossChannelNormalization] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization/init(device:kernelSize:)
func NewCNNCrossChannelNormalizationWithDeviceKernelSize(device metal.MTLDevice, kernelSize uint) MPSCNNCrossChannelNormalization {
	instance := getMPSCNNCrossChannelNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelSize:"), device, kernelSize)
	return MPSCNNCrossChannelNormalizationFromID(rv)
}

// Initializes a normalization kernel in a channel.
//
// device: The device the filter will run on.
//
// kernelSize: The size of the kernel, in both x and y dimensions.
//
// # Return Value
//
// A valid [MPSCNNCrossChannelNormalization] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization/init(device:kernelSize:)
func (c MPSCNNCrossChannelNormalization) InitWithDeviceKernelSize(device metal.MTLDevice, kernelSize uint) MPSCNNCrossChannelNormalization {
	rv := objc.Send[MPSCNNCrossChannelNormalization](c.ID, objc.Sel("initWithDevice:kernelSize:"), device, kernelSize)
	return rv
}

// The “alpha” variable of the kernel function.
//
// # Discussion
//
// The default value is `1.0`. Values must be non-negative.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization/alpha
func (c MPSCNNCrossChannelNormalization) Alpha() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("alpha"))
	return rv
}
func (c MPSCNNCrossChannelNormalization) SetAlpha(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlpha:"), value)
}

// The “beta” variable of the kernel function.
//
// # Discussion
//
// The default value is `5.0`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization/beta
func (c MPSCNNCrossChannelNormalization) Beta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("beta"))
	return rv
}
func (c MPSCNNCrossChannelNormalization) SetBeta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBeta:"), value)
}

// The “delta” variable of the kernel function.
//
// # Discussion
//
// The default value is `1.0`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization/delta
func (c MPSCNNCrossChannelNormalization) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}
func (c MPSCNNCrossChannelNormalization) SetDelta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelta:"), value)
}

// The size of the square kernel window.
//
// # Discussion
//
// The default value is `5`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNCrossChannelNormalization/kernelSize
func (c MPSCNNCrossChannelNormalization) KernelSize() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("kernelSize"))
	return rv
}

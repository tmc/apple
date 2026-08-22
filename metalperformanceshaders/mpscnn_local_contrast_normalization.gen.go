// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNLocalContrastNormalization] class.
var (
	_MPSCNNLocalContrastNormalizationClass     MPSCNNLocalContrastNormalizationClass
	_MPSCNNLocalContrastNormalizationClassOnce sync.Once
)

func getMPSCNNLocalContrastNormalizationClass() MPSCNNLocalContrastNormalizationClass {
	_MPSCNNLocalContrastNormalizationClassOnce.Do(func() {
		_MPSCNNLocalContrastNormalizationClass = MPSCNNLocalContrastNormalizationClass{class: objc.GetClass("MPSCNNLocalContrastNormalization")}
	})
	return _MPSCNNLocalContrastNormalizationClass
}

// GetMPSCNNLocalContrastNormalizationClass returns the class object for MPSCNNLocalContrastNormalization.
func GetMPSCNNLocalContrastNormalizationClass() MPSCNNLocalContrastNormalizationClass {
	return getMPSCNNLocalContrastNormalizationClass()
}

type MPSCNNLocalContrastNormalizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLocalContrastNormalizationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLocalContrastNormalizationClass) Alloc() MPSCNNLocalContrastNormalization {
	rv := objc.Send[MPSCNNLocalContrastNormalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A local-contrast normalization kernel.
//
// # Overview
//
// The local contrast normalization kernel is quite similar to the spatial
// normalization kernel, described in the [MPSCNNSpatialNormalization] class,
// in that it applies the kernel over local regions which extend spatially,
// but are in separate feature channels (i.e., they have the shape `1 x kernel
// width x kernel height`). However, instead of dividing by the local
// “energy” of the feature, the denominator uses the local variance of the
// feature - effectively the mean value of the feature is subtracted from the
// signal. For each feature channel, the function computes the variance
// `VAR(i,j)` and mean `M(i,j)` of `X(i,j)` inside each rectangle around the
// spatial point `(i,j)`. Then the result is computed for each element of [X]
// as follows:
//
// [media-2903532]
//
// Where `kw` and `kh` are the values of the `kernelWidth` and the
// `kernelHeight` properties, respectively, and the values of the
// [MPSCNNLocalContrastNormalization.Pm],
// [MPSCNNLocalContrastNormalization.Ps], and
// [MPSCNNLocalContrastNormalization.P0] properties can be used to offset and
// scale the result in various ways. For example setting `pm=0`, `ps=1`,
// `p0=1`, `delta=0`, `alpha=1.0` and `beta=0.5` scales input data so that the
// result has unit variance and zero mean, provided that input variance is
// positive.
//
// It is your responsibility to ensure that the combination of the values of
// the [MPSCNNLocalContrastNormalization.Delta] and
// [MPSCNNLocalContrastNormalization.Alpha] properties does not result in a
// situation where the denominator becomes zero - in such situations the
// resulting pixel-value is undefined. A good way to guard against tiny
// variances is to regulate the expression with a small delta value, for
// example `delta=1/1024`.
//
// # Initializers
//
//   - [MPSCNNLocalContrastNormalization.InitWithDeviceKernelWidthKernelHeight]: Initializes a local contrast normalization kernel.
//
// # Instance Properties
//
//   - [MPSCNNLocalContrastNormalization.Alpha]: The “alpha” variable of the kernel function.
//   - [MPSCNNLocalContrastNormalization.SetAlpha]
//   - [MPSCNNLocalContrastNormalization.Beta]: The “beta” variable of the kernel function.
//   - [MPSCNNLocalContrastNormalization.SetBeta]
//   - [MPSCNNLocalContrastNormalization.Delta]: The “delta” variable of the kernel function.
//   - [MPSCNNLocalContrastNormalization.SetDelta]
//   - [MPSCNNLocalContrastNormalization.P0]: The “p0” variable of the kernel function.
//   - [MPSCNNLocalContrastNormalization.SetP0]
//   - [MPSCNNLocalContrastNormalization.Pm]: The “pm” variable of the kernel function.
//   - [MPSCNNLocalContrastNormalization.SetPm]
//   - [MPSCNNLocalContrastNormalization.Ps]: The “ps” variable of the kernel function.
//   - [MPSCNNLocalContrastNormalization.SetPs]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization
type MPSCNNLocalContrastNormalization struct {
	MPSCNNKernel
}

// MPSCNNLocalContrastNormalizationFromID constructs a [MPSCNNLocalContrastNormalization] from an objc.ID.
//
// A local-contrast normalization kernel.
func MPSCNNLocalContrastNormalizationFromID(id objc.ID) MPSCNNLocalContrastNormalization {
	return MPSCNNLocalContrastNormalization{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNLocalContrastNormalization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLocalContrastNormalization] class.
//
// # Initializers
//
//   - [IMPSCNNLocalContrastNormalization.InitWithDeviceKernelWidthKernelHeight]: Initializes a local contrast normalization kernel.
//
// # Instance Properties
//
//   - [IMPSCNNLocalContrastNormalization.Alpha]: The “alpha” variable of the kernel function.
//   - [IMPSCNNLocalContrastNormalization.SetAlpha]
//   - [IMPSCNNLocalContrastNormalization.Beta]: The “beta” variable of the kernel function.
//   - [IMPSCNNLocalContrastNormalization.SetBeta]
//   - [IMPSCNNLocalContrastNormalization.Delta]: The “delta” variable of the kernel function.
//   - [IMPSCNNLocalContrastNormalization.SetDelta]
//   - [IMPSCNNLocalContrastNormalization.P0]: The “p0” variable of the kernel function.
//   - [IMPSCNNLocalContrastNormalization.SetP0]
//   - [IMPSCNNLocalContrastNormalization.Pm]: The “pm” variable of the kernel function.
//   - [IMPSCNNLocalContrastNormalization.SetPm]
//   - [IMPSCNNLocalContrastNormalization.Ps]: The “ps” variable of the kernel function.
//   - [IMPSCNNLocalContrastNormalization.SetPs]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization
type IMPSCNNLocalContrastNormalization interface {
	IMPSCNNKernel

	// Topic: Initializers

	// Initializes a local contrast normalization kernel.
	InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNLocalContrastNormalization

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
	// The “p0” variable of the kernel function.
	P0() float32
	SetP0(value float32)
	// The “pm” variable of the kernel function.
	Pm() float32
	SetPm(value float32)
	// The “ps” variable of the kernel function.
	Ps() float32
	SetPs(value float32)
}

// Init initializes the instance.
func (c MPSCNNLocalContrastNormalization) Init() MPSCNNLocalContrastNormalization {
	rv := objc.Send[MPSCNNLocalContrastNormalization](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLocalContrastNormalization) Autorelease() MPSCNNLocalContrastNormalization {
	rv := objc.Send[MPSCNNLocalContrastNormalization](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLocalContrastNormalization creates a new MPSCNNLocalContrastNormalization instance.
func NewMPSCNNLocalContrastNormalization() MPSCNNLocalContrastNormalization {
	class := getMPSCNNLocalContrastNormalizationClass()
	rv := objc.Send[MPSCNNLocalContrastNormalization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNLocalContrastNormalizationWithCoder(aDecoder foundation.INSCoder) MPSCNNLocalContrastNormalization {
	instance := getMPSCNNLocalContrastNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNLocalContrastNormalizationFromID(rv)
}

// Initializes a local contrast normalization kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization/init(coder:device:)
func NewCNNLocalContrastNormalizationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNLocalContrastNormalization {
	instance := getMPSCNNLocalContrastNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNLocalContrastNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNLocalContrastNormalizationWithDevice(device metal.MTLDevice) MPSCNNLocalContrastNormalization {
	instance := getMPSCNNLocalContrastNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNLocalContrastNormalizationFromID(rv)
}

// Initializes a local contrast normalization kernel.
//
// device: The device the kernel will run on.
//
// kernelWidth: The width of the kernel.
//
// kernelHeight: The height of the kernel.
//
// # Return Value
//
// A valid [MPSCNNLocalContrastNormalization] object or `nil`, if failure.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization/init(device:kernelWidth:kernelHeight:)
func NewCNNLocalContrastNormalizationWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNLocalContrastNormalization {
	instance := getMPSCNNLocalContrastNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNLocalContrastNormalizationFromID(rv)
}

// Initializes a local contrast normalization kernel.
//
// device: The device the kernel will run on.
//
// kernelWidth: The width of the kernel.
//
// kernelHeight: The height of the kernel.
//
// # Return Value
//
// A valid [MPSCNNLocalContrastNormalization] object or `nil`, if failure.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization/init(device:kernelWidth:kernelHeight:)
func (c MPSCNNLocalContrastNormalization) InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNLocalContrastNormalization {
	rv := objc.Send[MPSCNNLocalContrastNormalization](c.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return rv
}

// The “alpha” variable of the kernel function.
//
// # Discussion
//
// The default value is `1.0`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization/alpha
func (c MPSCNNLocalContrastNormalization) Alpha() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("alpha"))
	return rv
}
func (c MPSCNNLocalContrastNormalization) SetAlpha(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlpha:"), value)
}

// The “beta” variable of the kernel function.
//
// # Discussion
//
// The default value is `0.5`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization/beta
func (c MPSCNNLocalContrastNormalization) Beta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("beta"))
	return rv
}
func (c MPSCNNLocalContrastNormalization) SetBeta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBeta:"), value)
}

// The “delta” variable of the kernel function.
//
// # Discussion
//
// The default value is `1/1024`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization/delta
func (c MPSCNNLocalContrastNormalization) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}
func (c MPSCNNLocalContrastNormalization) SetDelta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelta:"), value)
}

// The “p0” variable of the kernel function.
//
// # Discussion
//
// The default value is `1.0`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization/p0
func (c MPSCNNLocalContrastNormalization) P0() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("p0"))
	return rv
}
func (c MPSCNNLocalContrastNormalization) SetP0(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setP0:"), value)
}

// The “pm” variable of the kernel function.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization/pm
func (c MPSCNNLocalContrastNormalization) Pm() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("pm"))
	return rv
}
func (c MPSCNNLocalContrastNormalization) SetPm(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPm:"), value)
}

// The “ps” variable of the kernel function.
//
// # Discussion
//
// The default value is `1.0`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLocalContrastNormalization/ps
func (c MPSCNNLocalContrastNormalization) Ps() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("ps"))
	return rv
}
func (c MPSCNNLocalContrastNormalization) SetPs(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPs:"), value)
}

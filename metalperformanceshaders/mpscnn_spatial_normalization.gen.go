// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSpatialNormalization] class.
var (
	_MPSCNNSpatialNormalizationClass     MPSCNNSpatialNormalizationClass
	_MPSCNNSpatialNormalizationClassOnce sync.Once
)

func getMPSCNNSpatialNormalizationClass() MPSCNNSpatialNormalizationClass {
	_MPSCNNSpatialNormalizationClassOnce.Do(func() {
		_MPSCNNSpatialNormalizationClass = MPSCNNSpatialNormalizationClass{class: objc.GetClass("MPSCNNSpatialNormalization")}
	})
	return _MPSCNNSpatialNormalizationClass
}

// GetMPSCNNSpatialNormalizationClass returns the class object for MPSCNNSpatialNormalization.
func GetMPSCNNSpatialNormalizationClass() MPSCNNSpatialNormalizationClass {
	return getMPSCNNSpatialNormalizationClass()
}

type MPSCNNSpatialNormalizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSpatialNormalizationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSpatialNormalizationClass) Alloc() MPSCNNSpatialNormalization {
	rv := objc.Send[MPSCNNSpatialNormalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A spatial normalization kernel.
//
// # Overview
//
// The spatial normalization for a feature channel applies the kernel over
// local regions which extend spatially, but are in separate feature channels
// (i.e., they have the shape `1 x kernel width x kernel height`).
//
// For each feature channel, the function computes the sum of squares of [X]
// inside each rectangle, `N2(i,j)`. It then divides each element of [X] as
// follows:
//
// [media-2903551]
//
// Where `kw` and `kh` are the values of the `kernelWidth` and `kernelHeight`
// properties, respectively. It is your responsibility to ensure that the
// combination of the values of the [MPSCNNSpatialNormalization.Delta] and
// [MPSCNNSpatialNormalization.Alpha] `kernelWidth` `kernelHeight` properties
// does not result in a situation where the denominator becomes zero (in such
// situations the resulting pixel-value is undefined).
//
// # Initializers
//
//   - [MPSCNNSpatialNormalization.InitWithDeviceKernelWidthKernelHeight]: Initializes a spatial normalization kernel.
//
// # Instance Properties
//
//   - [MPSCNNSpatialNormalization.Alpha]: The “alpha” variable of the kernel function.
//   - [MPSCNNSpatialNormalization.SetAlpha]
//   - [MPSCNNSpatialNormalization.Beta]: The “beta” variable of the kernel function.
//   - [MPSCNNSpatialNormalization.SetBeta]
//   - [MPSCNNSpatialNormalization.Delta]: The “delta” variable of the kernel function.
//   - [MPSCNNSpatialNormalization.SetDelta]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization
type MPSCNNSpatialNormalization struct {
	MPSCNNKernel
}

// MPSCNNSpatialNormalizationFromID constructs a [MPSCNNSpatialNormalization] from an objc.ID.
//
// A spatial normalization kernel.
func MPSCNNSpatialNormalizationFromID(id objc.ID) MPSCNNSpatialNormalization {
	return MPSCNNSpatialNormalization{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNSpatialNormalization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSpatialNormalization] class.
//
// # Initializers
//
//   - [IMPSCNNSpatialNormalization.InitWithDeviceKernelWidthKernelHeight]: Initializes a spatial normalization kernel.
//
// # Instance Properties
//
//   - [IMPSCNNSpatialNormalization.Alpha]: The “alpha” variable of the kernel function.
//   - [IMPSCNNSpatialNormalization.SetAlpha]
//   - [IMPSCNNSpatialNormalization.Beta]: The “beta” variable of the kernel function.
//   - [IMPSCNNSpatialNormalization.SetBeta]
//   - [IMPSCNNSpatialNormalization.Delta]: The “delta” variable of the kernel function.
//   - [IMPSCNNSpatialNormalization.SetDelta]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization
type IMPSCNNSpatialNormalization interface {
	IMPSCNNKernel

	// Topic: Initializers

	// Initializes a spatial normalization kernel.
	InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNSpatialNormalization

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
}

// Init initializes the instance.
func (c MPSCNNSpatialNormalization) Init() MPSCNNSpatialNormalization {
	rv := objc.Send[MPSCNNSpatialNormalization](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSpatialNormalization) Autorelease() MPSCNNSpatialNormalization {
	rv := objc.Send[MPSCNNSpatialNormalization](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSpatialNormalization creates a new MPSCNNSpatialNormalization instance.
func NewMPSCNNSpatialNormalization() MPSCNNSpatialNormalization {
	class := getMPSCNNSpatialNormalizationClass()
	rv := objc.Send[MPSCNNSpatialNormalization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNSpatialNormalizationWithCoder(aDecoder foundation.INSCoder) MPSCNNSpatialNormalization {
	instance := getMPSCNNSpatialNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNSpatialNormalizationFromID(rv)
}

// Initializes a spatial normalization kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization/init(coder:device:)
func NewCNNSpatialNormalizationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNSpatialNormalization {
	instance := getMPSCNNSpatialNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNSpatialNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNSpatialNormalizationWithDevice(device metal.MTLDevice) MPSCNNSpatialNormalization {
	instance := getMPSCNNSpatialNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNSpatialNormalizationFromID(rv)
}

// Initializes a spatial normalization kernel.
//
// device: The device the kernel will run on.
//
// kernelWidth: The width of the kernel.
//
// kernelHeight: The height of the kernel.
//
// # Return Value
//
// A valid [MPSCNNSpatialNormalization] object or `nil`, if failure.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization/init(device:kernelWidth:kernelHeight:)
func NewCNNSpatialNormalizationWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNSpatialNormalization {
	instance := getMPSCNNSpatialNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSCNNSpatialNormalizationFromID(rv)
}

// Initializes a spatial normalization kernel.
//
// device: The device the kernel will run on.
//
// kernelWidth: The width of the kernel.
//
// kernelHeight: The height of the kernel.
//
// # Return Value
//
// A valid [MPSCNNSpatialNormalization] object or `nil`, if failure.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization/init(device:kernelWidth:kernelHeight:)
func (c MPSCNNSpatialNormalization) InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSCNNSpatialNormalization {
	rv := objc.Send[MPSCNNSpatialNormalization](c.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return rv
}

// The “alpha” variable of the kernel function.
//
// # Discussion
//
// The default value is `1.0`. Values must be non-negative.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization/alpha
func (c MPSCNNSpatialNormalization) Alpha() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("alpha"))
	return rv
}
func (c MPSCNNSpatialNormalization) SetAlpha(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlpha:"), value)
}

// The “beta” variable of the kernel function.
//
// # Discussion
//
// The default value is `5.0`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization/beta
func (c MPSCNNSpatialNormalization) Beta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("beta"))
	return rv
}
func (c MPSCNNSpatialNormalization) SetBeta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setBeta:"), value)
}

// The “delta” variable of the kernel function.
//
// # Discussion
//
// The default value is `1.0`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization/delta
func (c MPSCNNSpatialNormalization) Delta() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("delta"))
	return rv
}
func (c MPSCNNSpatialNormalization) SetDelta(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelta:"), value)
}

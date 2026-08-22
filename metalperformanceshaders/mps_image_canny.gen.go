// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageCanny] class.
var (
	_MPSImageCannyClass     MPSImageCannyClass
	_MPSImageCannyClassOnce sync.Once
)

func getMPSImageCannyClass() MPSImageCannyClass {
	_MPSImageCannyClassOnce.Do(func() {
		_MPSImageCannyClass = MPSImageCannyClass{class: objc.GetClass("MPSImageCanny")}
	})
	return _MPSImageCannyClass
}

// GetMPSImageCannyClass returns the class object for MPSImageCanny.
func GetMPSImageCannyClass() MPSImageCannyClass {
	return getMPSImageCannyClass()
}

type MPSImageCannyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageCannyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageCannyClass) Alloc() MPSImageCanny {
	rv := objc.Send[MPSImageCanny](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSImageCanny.InitWithDeviceLinearToGrayScaleTransformSigma]
//
// # Instance Properties
//
//   - [MPSImageCanny.ColorTransform]
//   - [MPSImageCanny.HighThreshold]
//   - [MPSImageCanny.SetHighThreshold]
//   - [MPSImageCanny.LowThreshold]
//   - [MPSImageCanny.SetLowThreshold]
//   - [MPSImageCanny.Sigma]
//   - [MPSImageCanny.UseFastMode]
//   - [MPSImageCanny.SetUseFastMode]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny
type MPSImageCanny struct {
	MPSUnaryImageKernel
}

// MPSImageCannyFromID constructs a [MPSImageCanny] from an objc.ID.
func MPSImageCannyFromID(id objc.ID) MPSImageCanny {
	return MPSImageCanny{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageCanny adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageCanny] class.
//
// # Initializers
//
//   - [IMPSImageCanny.InitWithDeviceLinearToGrayScaleTransformSigma]
//
// # Instance Properties
//
//   - [IMPSImageCanny.ColorTransform]
//   - [IMPSImageCanny.HighThreshold]
//   - [IMPSImageCanny.SetHighThreshold]
//   - [IMPSImageCanny.LowThreshold]
//   - [IMPSImageCanny.SetLowThreshold]
//   - [IMPSImageCanny.Sigma]
//   - [IMPSImageCanny.UseFastMode]
//   - [IMPSImageCanny.SetUseFastMode]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny
type IMPSImageCanny interface {
	IMPSUnaryImageKernel

	// Topic: Initializers

	InitWithDeviceLinearToGrayScaleTransformSigma(device metal.MTLDevice, transform *float32, sigma float32) MPSImageCanny

	// Topic: Instance Properties

	ColorTransform() unsafe.Pointer
	HighThreshold() float32
	SetHighThreshold(value float32)
	LowThreshold() float32
	SetLowThreshold(value float32)
	Sigma() float32
	UseFastMode() bool
	SetUseFastMode(value bool)
}

// Init initializes the instance.
func (i MPSImageCanny) Init() MPSImageCanny {
	rv := objc.Send[MPSImageCanny](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageCanny) Autorelease() MPSImageCanny {
	rv := objc.Send[MPSImageCanny](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageCanny creates a new MPSImageCanny instance.
func NewMPSImageCanny() MPSImageCanny {
	class := getMPSImageCannyClass()
	rv := objc.Send[MPSImageCanny](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageCannyWithCoder(aDecoder foundation.INSCoder) MPSImageCanny {
	instance := getMPSImageCannyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageCannyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/init(coder:device:)
func NewImageCannyWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageCanny {
	instance := getMPSImageCannyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageCannyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/init(device:)
func NewImageCannyWithDevice(device metal.MTLDevice) MPSImageCanny {
	instance := getMPSImageCannyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageCannyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/init(device:linearToGrayScaleTransform:sigma:)
func NewImageCannyWithDeviceLinearToGrayScaleTransformSigma(device metal.MTLDevice, transform *float32, sigma float32) MPSImageCanny {
	instance := getMPSImageCannyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:linearToGrayScaleTransform:sigma:"), device, transform, sigma)
	return MPSImageCannyFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/init(device:linearToGrayScaleTransform:sigma:)
func (i MPSImageCanny) InitWithDeviceLinearToGrayScaleTransformSigma(device metal.MTLDevice, transform *float32, sigma float32) MPSImageCanny {
	rv := objc.Send[MPSImageCanny](i.ID, objc.Sel("initWithDevice:linearToGrayScaleTransform:sigma:"), device, transform, sigma)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/colorTransform
func (i MPSImageCanny) ColorTransform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("colorTransform"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/highThreshold
func (i MPSImageCanny) HighThreshold() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("highThreshold"))
	return rv
}
func (i MPSImageCanny) SetHighThreshold(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setHighThreshold:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/lowThreshold
func (i MPSImageCanny) LowThreshold() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("lowThreshold"))
	return rv
}
func (i MPSImageCanny) SetLowThreshold(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setLowThreshold:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/sigma
func (i MPSImageCanny) Sigma() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("sigma"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny/useFastMode
func (i MPSImageCanny) UseFastMode() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("useFastMode"))
	return rv
}
func (i MPSImageCanny) SetUseFastMode(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setUseFastMode:"), value)
}

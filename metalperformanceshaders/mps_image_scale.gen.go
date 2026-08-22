// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageScale] class.
var (
	_MPSImageScaleClass     MPSImageScaleClass
	_MPSImageScaleClassOnce sync.Once
)

func getMPSImageScaleClass() MPSImageScaleClass {
	_MPSImageScaleClassOnce.Do(func() {
		_MPSImageScaleClass = MPSImageScaleClass{class: objc.GetClass("MPSImageScale")}
	})
	return _MPSImageScaleClass
}

// GetMPSImageScaleClass returns the class object for MPSImageScale.
func GetMPSImageScaleClass() MPSImageScaleClass {
	return getMPSImageScaleClass()
}

type MPSImageScaleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageScaleClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageScaleClass) Alloc() MPSImageScale {
	rv := objc.Send[MPSImageScale](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that resizes and changes the aspect ratio of an image.
//
// # Instance Properties
//
//   - [MPSImageScale.ScaleTransform]
//   - [MPSImageScale.SetScaleTransform]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale
type MPSImageScale struct {
	MPSUnaryImageKernel
}

// MPSImageScaleFromID constructs a [MPSImageScale] from an objc.ID.
//
// A filter that resizes and changes the aspect ratio of an image.
func MPSImageScaleFromID(id objc.ID) MPSImageScale {
	return MPSImageScale{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageScale adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageScale] class.
//
// # Instance Properties
//
//   - [IMPSImageScale.ScaleTransform]
//   - [IMPSImageScale.SetScaleTransform]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale
type IMPSImageScale interface {
	IMPSUnaryImageKernel

	// Topic: Instance Properties

	ScaleTransform() *MPSScaleTransform
	SetScaleTransform(value *MPSScaleTransform)
}

// Init initializes the instance.
func (i MPSImageScale) Init() MPSImageScale {
	rv := objc.Send[MPSImageScale](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageScale) Autorelease() MPSImageScale {
	rv := objc.Send[MPSImageScale](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageScale creates a new MPSImageScale instance.
func NewMPSImageScale() MPSImageScale {
	class := getMPSImageScaleClass()
	rv := objc.Send[MPSImageScale](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageScaleWithCoder(aDecoder foundation.INSCoder) MPSImageScale {
	instance := getMPSImageScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageScaleFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale/init(coder:device:)
func NewImageScaleWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageScale {
	instance := getMPSImageScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageScaleFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale/init(device:)
func NewImageScaleWithDevice(device metal.MTLDevice) MPSImageScale {
	instance := getMPSImageScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageScaleFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale/scaleTransform
func (i MPSImageScale) ScaleTransform() *MPSScaleTransform {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("scaleTransform"))
	return (*MPSScaleTransform)(rv)
}
func (i MPSImageScale) SetScaleTransform(value *MPSScaleTransform) {
	objc.Send[struct{}](i.ID, objc.Sel("setScaleTransform:"), value)
}

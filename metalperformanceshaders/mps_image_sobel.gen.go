// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageSobel] class.
var (
	_MPSImageSobelClass     MPSImageSobelClass
	_MPSImageSobelClassOnce sync.Once
)

func getMPSImageSobelClass() MPSImageSobelClass {
	_MPSImageSobelClassOnce.Do(func() {
		_MPSImageSobelClass = MPSImageSobelClass{class: objc.GetClass("MPSImageSobel")}
	})
	return _MPSImageSobelClass
}

// GetMPSImageSobelClass returns the class object for MPSImageSobel.
func GetMPSImageSobelClass() MPSImageSobelClass {
	return getMPSImageSobelClass()
}

type MPSImageSobelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageSobelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageSobelClass) Alloc() MPSImageSobel {
	rv := objc.Send[MPSImageSobel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that convolves an image with the Sobel operator.
//
// # Overview
//
// When the color model (e.g. RGB, two-channel, grayscale, etc.) of the source
// and destination textures match, the filter is applied to each color channel
// separately. If the destination is single-channel (i.e. monochrome) but the
// source is multi-channel, the pixel values are converted to grayscale before
// applying the Sobel operator by using the linear gray color transform vector
// `v` shown in the code listing below.
//
// # Methods
//
//   - [MPSImageSobel.InitWithDeviceLinearGrayColorTransform]: Initializes a Sobel filter on a given device using a specific color transform.
//
// # Properties
//
//   - [MPSImageSobel.ColorTransform]: The color transform used to initialize the Sobel filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSobel
type MPSImageSobel struct {
	MPSUnaryImageKernel
}

// MPSImageSobelFromID constructs a [MPSImageSobel] from an objc.ID.
//
// A filter that convolves an image with the Sobel operator.
func MPSImageSobelFromID(id objc.ID) MPSImageSobel {
	return MPSImageSobel{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageSobel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageSobel] class.
//
// # Methods
//
//   - [IMPSImageSobel.InitWithDeviceLinearGrayColorTransform]: Initializes a Sobel filter on a given device using a specific color transform.
//
// # Properties
//
//   - [IMPSImageSobel.ColorTransform]: The color transform used to initialize the Sobel filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSobel
type IMPSImageSobel interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes a Sobel filter on a given device using a specific color transform.
	InitWithDeviceLinearGrayColorTransform(device metal.MTLDevice, transform *float32) MPSImageSobel

	// Topic: Properties

	// The color transform used to initialize the Sobel filter.
	ColorTransform() unsafe.Pointer
}

// Init initializes the instance.
func (i MPSImageSobel) Init() MPSImageSobel {
	rv := objc.Send[MPSImageSobel](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageSobel) Autorelease() MPSImageSobel {
	rv := objc.Send[MPSImageSobel](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageSobel creates a new MPSImageSobel instance.
func NewMPSImageSobel() MPSImageSobel {
	class := getMPSImageSobelClass()
	rv := objc.Send[MPSImageSobel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageSobelWithCoder(aDecoder foundation.INSCoder) MPSImageSobel {
	instance := getMPSImageSobelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageSobelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSobel/init(coder:device:)
func NewImageSobelWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageSobel {
	instance := getMPSImageSobelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageSobelFromID(rv)
}

// Initializes a Sobel filter on a given device using the default color
// transform.
//
// device: The Metal device the filter will run on.
//
// # Return Value
//
// An initialized Sobel filter object.
//
// # Discussion
//
// The default color transform matrix is an array of 3 floats set to the
// BT.601/JPEG standard: `{0.299f, 0.587f, 0.114f}`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSobel/init(device:)
func NewImageSobelWithDevice(device metal.MTLDevice) MPSImageSobel {
	instance := getMPSImageSobelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageSobelFromID(rv)
}

// Initializes a Sobel filter on a given device using a specific color
// transform.
//
// device: The Metal device the filter will run on.
//
// transform: The color transform to use. This matrix is an array of 3 floats that
// describes the RGB-to-grayscale color transform:
//
// `Luminance = transform[0] * pixel.X() + transform[1] * pixel.Y() +
// transform[2] * pixel.Z()`
//
// # Return Value
//
// An initialized Sobel filter object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSobel/init(device:linearGrayColorTransform:)
func NewImageSobelWithDeviceLinearGrayColorTransform(device metal.MTLDevice, transform *float32) MPSImageSobel {
	instance := getMPSImageSobelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:linearGrayColorTransform:"), device, transform)
	return MPSImageSobelFromID(rv)
}

// Initializes a Sobel filter on a given device using a specific color
// transform.
//
// device: The Metal device the filter will run on.
//
// transform: The color transform to use. This matrix is an array of 3 floats that
// describes the RGB-to-grayscale color transform:
//
// `Luminance = transform[0] * pixel.X() + transform[1] * pixel.Y() +
// transform[2] * pixel.Z()`
//
// # Return Value
//
// An initialized Sobel filter object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSobel/init(device:linearGrayColorTransform:)
func (i MPSImageSobel) InitWithDeviceLinearGrayColorTransform(device metal.MTLDevice, transform *float32) MPSImageSobel {
	rv := objc.Send[MPSImageSobel](i.ID, objc.Sel("initWithDevice:linearGrayColorTransform:"), device, transform)
	return rv
}

// The color transform used to initialize the Sobel filter.
//
// # Discussion
//
// This property returns a pointer to the array of 3 floats used to convert
// RGBA, RGB or RG source images to the destination texture format when said
// destination is monochrome.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSobel/colorTransform
func (i MPSImageSobel) ColorTransform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("colorTransform"))
	return rv
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageConversion] class.
var (
	_MPSImageConversionClass     MPSImageConversionClass
	_MPSImageConversionClassOnce sync.Once
)

func getMPSImageConversionClass() MPSImageConversionClass {
	_MPSImageConversionClassOnce.Do(func() {
		_MPSImageConversionClass = MPSImageConversionClass{class: objc.GetClass("MPSImageConversion")}
	})
	return _MPSImageConversionClass
}

// GetMPSImageConversionClass returns the class object for MPSImageConversion.
func GetMPSImageConversionClass() MPSImageConversionClass {
	return getMPSImageConversionClass()
}

type MPSImageConversionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageConversionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageConversionClass) Alloc() MPSImageConversion {
	rv := objc.Send[MPSImageConversion](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that performs a conversion of color space, alpha, or pixel format.
//
// # Overview
//
// An [MPSImageConversion] filter allows you to change the alpha encoding or
// color space of an image. For example, you can convert an image with a
// premultiplied alpha to non-premultiplied, or change the color space from
// one variant to another.
//
// As with all Metal Performance Shaders filters, the conversion filter allows
// for source and destination textures with different pixel formats and, in
// that case, will convert the source texture’s format to the destination
// texture’s format. See [Supported Pixel Formats for Image Kernels] for a
// list of supported pixel formats.
//
// The following listing shows how you can create an image conversion filter
// to map the color intensity from the sRGB color space to a linear gamma
// curve.
//
// Listing 1. Mapping color intensity from the sRGB color space to a linear
// gamma curve.
//
// # Methods
//
//   - [MPSImageConversion.InitWithDeviceSrcAlphaDestAlphaBackgroundColorConversionInfo]: Initializes a filter that can convert texture color space, alpha, and pixel format.
//
// # Properties
//
//   - [MPSImageConversion.SourceAlpha]: Premultiplication description for the source texture.
//   - [MPSImageConversion.DestinationAlpha]: Premultiplication description for the destination texture.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConversion
//
// [Supported Pixel Formats for Image Kernels]: https://developer.apple.com/documentation/MetalPerformanceShaders/image-filters#Supported-Pixel-Formats-for-Image-Kernels
type MPSImageConversion struct {
	MPSUnaryImageKernel
}

// MPSImageConversionFromID constructs a [MPSImageConversion] from an objc.ID.
//
// A filter that performs a conversion of color space, alpha, or pixel format.
func MPSImageConversionFromID(id objc.ID) MPSImageConversion {
	return MPSImageConversion{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageConversion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageConversion] class.
//
// # Methods
//
//   - [IMPSImageConversion.InitWithDeviceSrcAlphaDestAlphaBackgroundColorConversionInfo]: Initializes a filter that can convert texture color space, alpha, and pixel format.
//
// # Properties
//
//   - [IMPSImageConversion.SourceAlpha]: Premultiplication description for the source texture.
//   - [IMPSImageConversion.DestinationAlpha]: Premultiplication description for the destination texture.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConversion
type IMPSImageConversion interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes a filter that can convert texture color space, alpha, and pixel format.
	InitWithDeviceSrcAlphaDestAlphaBackgroundColorConversionInfo(device metal.MTLDevice, srcAlpha MPSAlphaType, destAlpha MPSAlphaType, backgroundColor *float64, conversionInfo coregraphics.CGColorConversionInfoRef) MPSImageConversion

	// Topic: Properties

	// Premultiplication description for the source texture.
	SourceAlpha() MPSAlphaType
	// Premultiplication description for the destination texture.
	DestinationAlpha() MPSAlphaType
}

// Init initializes the instance.
func (i MPSImageConversion) Init() MPSImageConversion {
	rv := objc.Send[MPSImageConversion](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageConversion) Autorelease() MPSImageConversion {
	rv := objc.Send[MPSImageConversion](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageConversion creates a new MPSImageConversion instance.
func NewMPSImageConversion() MPSImageConversion {
	class := getMPSImageConversionClass()
	rv := objc.Send[MPSImageConversion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageConversionWithCoder(aDecoder foundation.INSCoder) MPSImageConversion {
	instance := getMPSImageConversionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageConversionFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageConversionWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageConversion {
	instance := getMPSImageConversionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageConversionFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageConversionWithDevice(device metal.MTLDevice) MPSImageConversion {
	instance := getMPSImageConversionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageConversionFromID(rv)
}

// Initializes a filter that can convert texture color space, alpha, and pixel
// format.
//
// device: The device that the filter will run on.
//
// srcAlpha: The alpha encoding for the source texture.
//
// destAlpha: The alpha encoding for the destination texture.
//
// backgroundColor: An array of [CGFloat] values giving the background color to use when
// flattening an image.
//
// The color is in the source color space. The length of the array is the
// number of color channels in the source color space. If this parameter is
// not applicable to your desired conversion, use `{0}`.
//
// conversionInfo: The color space conversion to use. This value may be [NULL], indicating
// that no color space conversions need to be done.
//
// # Return Value
//
// An [MPSImageConversion] object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConversion/init(device:srcAlpha:destAlpha:backgroundColor:conversionInfo:)
//
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
func NewImageConversionWithDeviceSrcAlphaDestAlphaBackgroundColorConversionInfo(device metal.MTLDevice, srcAlpha MPSAlphaType, destAlpha MPSAlphaType, backgroundColor *float64, conversionInfo coregraphics.CGColorConversionInfoRef) MPSImageConversion {
	instance := getMPSImageConversionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:srcAlpha:destAlpha:backgroundColor:conversionInfo:"), device, srcAlpha, destAlpha, backgroundColor, conversionInfo)
	return MPSImageConversionFromID(rv)
}

// Initializes a filter that can convert texture color space, alpha, and pixel
// format.
//
// device: The device that the filter will run on.
//
// srcAlpha: The alpha encoding for the source texture.
//
// destAlpha: The alpha encoding for the destination texture.
//
// backgroundColor: An array of [CGFloat] values giving the background color to use when
// flattening an image.
//
// The color is in the source color space. The length of the array is the
// number of color channels in the source color space. If this parameter is
// not applicable to your desired conversion, use `{0}`.
//
// conversionInfo: The color space conversion to use. This value may be [NULL], indicating
// that no color space conversions need to be done.
//
// # Return Value
//
// An [MPSImageConversion] object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConversion/init(device:srcAlpha:destAlpha:backgroundColor:conversionInfo:)
//
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
func (i MPSImageConversion) InitWithDeviceSrcAlphaDestAlphaBackgroundColorConversionInfo(device metal.MTLDevice, srcAlpha MPSAlphaType, destAlpha MPSAlphaType, backgroundColor *float64, conversionInfo coregraphics.CGColorConversionInfoRef) MPSImageConversion {
	rv := objc.Send[MPSImageConversion](i.ID, objc.Sel("initWithDevice:srcAlpha:destAlpha:backgroundColor:conversionInfo:"), device, srcAlpha, destAlpha, backgroundColor, conversionInfo)
	return rv
}

// Premultiplication description for the source texture.
//
// # Discussion
//
// Most color space conversion operations can not work directly on
// premultiplied data. Use this property to tag premultiplied data so that the
// source texture can be un-premultiplied prior to the application of these
// transforms. The default value is [MPSAlphaTypeAlphaIsOne].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConversion/sourceAlpha
func (i MPSImageConversion) SourceAlpha() MPSAlphaType {
	rv := objc.Send[MPSAlphaType](i.ID, objc.Sel("sourceAlpha"))
	return MPSAlphaType(rv)
}

// Premultiplication description for the destination texture.
//
// # Discussion
//
// Color space conversion operations produce non-premultiplied data. Use this
// property to tag cases where premultiplied results are required. If the
// [MPSAlphaTypeAlphaIsOne] value is used, the alpha channel will be set to 1.
// The default value is [MPSAlphaTypeAlphaIsOne].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConversion/destinationAlpha
func (i MPSImageConversion) DestinationAlpha() MPSAlphaType {
	rv := objc.Send[MPSAlphaType](i.ID, objc.Sel("destinationAlpha"))
	return MPSAlphaType(rv)
}

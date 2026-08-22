// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTLowLatencySuperResolutionScalerConfiguration] class.
var (
	_VTLowLatencySuperResolutionScalerConfigurationClass     VTLowLatencySuperResolutionScalerConfigurationClass
	_VTLowLatencySuperResolutionScalerConfigurationClassOnce sync.Once
)

func getVTLowLatencySuperResolutionScalerConfigurationClass() VTLowLatencySuperResolutionScalerConfigurationClass {
	_VTLowLatencySuperResolutionScalerConfigurationClassOnce.Do(func() {
		_VTLowLatencySuperResolutionScalerConfigurationClass = VTLowLatencySuperResolutionScalerConfigurationClass{class: objc.GetClass("VTLowLatencySuperResolutionScalerConfiguration")}
	})
	return _VTLowLatencySuperResolutionScalerConfigurationClass
}

// GetVTLowLatencySuperResolutionScalerConfigurationClass returns the class object for VTLowLatencySuperResolutionScalerConfiguration.
func GetVTLowLatencySuperResolutionScalerConfigurationClass() VTLowLatencySuperResolutionScalerConfigurationClass {
	return getVTLowLatencySuperResolutionScalerConfigurationClass()
}

type VTLowLatencySuperResolutionScalerConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTLowLatencySuperResolutionScalerConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTLowLatencySuperResolutionScalerConfigurationClass) Alloc() VTLowLatencySuperResolutionScalerConfiguration {
	rv := objc.Send[VTLowLatencySuperResolutionScalerConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object you use to configure frame processor for low-latency
// super-resolution scaler processing.
//
// # Overview
//
// Use this object to configure a [VTFrameProcessor]. Query this interface
// also for important operating details, like the pixel buffer attributes
// required for frames you submit to the processor.
//
// # Creating a super resolution scaler configuration
//
//   - [VTLowLatencySuperResolutionScalerConfiguration.InitWithFrameWidthFrameHeightScaleFactor]: Creates a new low-latency super-resolution scaler configuration with specified frame width and height.
//
// # Inspecting the configuration
//
//   - [VTLowLatencySuperResolutionScalerConfiguration.FrameWidth]: Width of source frame in pixels.
//   - [VTLowLatencySuperResolutionScalerConfiguration.FrameHeight]: Height of source frame in pixels.
//   - [VTLowLatencySuperResolutionScalerConfiguration.ScaleFactor]: Scale factor with which you initialized the configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration
type VTLowLatencySuperResolutionScalerConfiguration struct {
	objectivec.Object
}

// VTLowLatencySuperResolutionScalerConfigurationFromID constructs a [VTLowLatencySuperResolutionScalerConfiguration] from an objc.ID.
//
// An object you use to configure frame processor for low-latency
// super-resolution scaler processing.
func VTLowLatencySuperResolutionScalerConfigurationFromID(id objc.ID) VTLowLatencySuperResolutionScalerConfiguration {
	return VTLowLatencySuperResolutionScalerConfiguration{objectivec.Object{ID: id}}
}

// NOTE: VTLowLatencySuperResolutionScalerConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTLowLatencySuperResolutionScalerConfiguration] class.
//
// # Creating a super resolution scaler configuration
//
//   - [IVTLowLatencySuperResolutionScalerConfiguration.InitWithFrameWidthFrameHeightScaleFactor]: Creates a new low-latency super-resolution scaler configuration with specified frame width and height.
//
// # Inspecting the configuration
//
//   - [IVTLowLatencySuperResolutionScalerConfiguration.FrameWidth]: Width of source frame in pixels.
//   - [IVTLowLatencySuperResolutionScalerConfiguration.FrameHeight]: Height of source frame in pixels.
//   - [IVTLowLatencySuperResolutionScalerConfiguration.ScaleFactor]: Scale factor with which you initialized the configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration
type IVTLowLatencySuperResolutionScalerConfiguration interface {
	objectivec.IObject

	// Topic: Creating a super resolution scaler configuration

	// Creates a new low-latency super-resolution scaler configuration with specified frame width and height.
	InitWithFrameWidthFrameHeightScaleFactor(frameWidth int, frameHeight int, scaleFactor float32) VTLowLatencySuperResolutionScalerConfiguration

	// Topic: Inspecting the configuration

	// Width of source frame in pixels.
	FrameWidth() int
	// Height of source frame in pixels.
	FrameHeight() int
	// Scale factor with which you initialized the configuration.
	ScaleFactor() float32
}

// Init initializes the instance.
func (v VTLowLatencySuperResolutionScalerConfiguration) Init() VTLowLatencySuperResolutionScalerConfiguration {
	rv := objc.Send[VTLowLatencySuperResolutionScalerConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTLowLatencySuperResolutionScalerConfiguration) Autorelease() VTLowLatencySuperResolutionScalerConfiguration {
	rv := objc.Send[VTLowLatencySuperResolutionScalerConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTLowLatencySuperResolutionScalerConfiguration creates a new VTLowLatencySuperResolutionScalerConfiguration instance.
func NewVTLowLatencySuperResolutionScalerConfiguration() VTLowLatencySuperResolutionScalerConfiguration {
	class := getVTLowLatencySuperResolutionScalerConfigurationClass()
	rv := objc.Send[VTLowLatencySuperResolutionScalerConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new low-latency super-resolution scaler configuration with
// specified frame width and height.
//
// frameWidth: Width of source frame in pixels.
//
// frameHeight: Height of source frame in pixels.
//
// scaleFactor: The scale factor to apply. This must be a supported value that
// [VTLowLatencySuperResolutionScalerConfigurationClass.SupportedScaleFactorsForFrameWidthFrameHeight]
// returns.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/init(frameWidth:frameHeight:scaleFactor:)
func NewVTLowLatencySuperResolutionScalerConfigurationWithFrameWidthFrameHeightScaleFactor(frameWidth int, frameHeight int, scaleFactor float32) VTLowLatencySuperResolutionScalerConfiguration {
	instance := getVTLowLatencySuperResolutionScalerConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameWidth:frameHeight:scaleFactor:"), frameWidth, frameHeight, scaleFactor)
	return VTLowLatencySuperResolutionScalerConfigurationFromID(rv)
}

// Creates a new low-latency super-resolution scaler configuration with
// specified frame width and height.
//
// frameWidth: Width of source frame in pixels.
//
// frameHeight: Height of source frame in pixels.
//
// scaleFactor: The scale factor to apply. This must be a supported value that
// [VTLowLatencySuperResolutionScalerConfigurationClass.SupportedScaleFactorsForFrameWidthFrameHeight]
// returns.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/init(frameWidth:frameHeight:scaleFactor:)
func (v VTLowLatencySuperResolutionScalerConfiguration) InitWithFrameWidthFrameHeightScaleFactor(frameWidth int, frameHeight int, scaleFactor float32) VTLowLatencySuperResolutionScalerConfiguration {
	rv := objc.Send[VTLowLatencySuperResolutionScalerConfiguration](v.ID, objc.Sel("initWithFrameWidth:frameHeight:scaleFactor:"), frameWidth, frameHeight, scaleFactor)
	return rv
}

// Returns an array of supported scale factors values, or an empty list if the
// processor doesn’t support the dimensions.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/supportedScaleFactorsForFrameWidth:frameHeight:
func (_VTLowLatencySuperResolutionScalerConfigurationClass VTLowLatencySuperResolutionScalerConfigurationClass) SupportedScaleFactorsForFrameWidthFrameHeight(frameWidth int, frameHeight int) []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](objc.ID(_VTLowLatencySuperResolutionScalerConfigurationClass.class), objc.Sel("supportedScaleFactorsForFrameWidth:frameHeight:"), frameWidth, frameHeight)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Width of source frame in pixels.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/frameWidth
func (v VTLowLatencySuperResolutionScalerConfiguration) FrameWidth() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameWidth"))
	return rv
}

// Height of source frame in pixels.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/frameHeight
func (v VTLowLatencySuperResolutionScalerConfiguration) FrameHeight() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameHeight"))
	return rv
}

// Scale factor with which you initialized the configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/scaleFactor
func (v VTLowLatencySuperResolutionScalerConfiguration) ScaleFactor() float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("scaleFactor"))
	return rv
}

// Pixel buffer attributes dictionary that describes requirements for pixel
// buffers which represent source frames and reference frames.
//
// # Discussion
//
// Use [CVPixelBufferCreateResolvedAttributesDictionary] to combine this
// dictionary with your pixel buffer attributes dictionary.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/sourcePixelBufferAttributes
func (v VTLowLatencySuperResolutionScalerConfiguration) SourcePixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourcePixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// Pixel buffer attributes dictionary that describes requirements for pixel
// buffers which represent destination frames.
//
// # Discussion
//
// Use [CVPixelBufferCreateResolvedAttributesDictionary] to combine this
// dictionary with your pixel buffer attributes dictionary.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/destinationPixelBufferAttributes
func (v VTLowLatencySuperResolutionScalerConfiguration) DestinationPixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationPixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// Available supported pixel formats for source frames for current
// configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/frameSupportedPixelFormats
func (v VTLowLatencySuperResolutionScalerConfiguration) FrameSupportedPixelFormats() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("frameSupportedPixelFormats"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Reports whether the system supports this processor on the current
// configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/isSupported
func (_VTLowLatencySuperResolutionScalerConfigurationClass VTLowLatencySuperResolutionScalerConfigurationClass) IsSupported() bool {
	rv := objc.Send[bool](objc.ID(_VTLowLatencySuperResolutionScalerConfigurationClass.class), objc.Sel("isSupported"))
	return rv
}

// Maximum dimensions for a source frame for the processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/maximumDimensions
func (_VTLowLatencySuperResolutionScalerConfigurationClass VTLowLatencySuperResolutionScalerConfigurationClass) MaximumDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](objc.ID(_VTLowLatencySuperResolutionScalerConfigurationClass.class), objc.Sel("maximumDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// Minimum dimensions for a source frame for the processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerConfiguration/minimumDimensions
func (_VTLowLatencySuperResolutionScalerConfigurationClass VTLowLatencySuperResolutionScalerConfigurationClass) MinimumDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](objc.ID(_VTLowLatencySuperResolutionScalerConfigurationClass.class), objc.Sel("minimumDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// Protocol methods for VTFrameProcessorConfiguration

// Returns the number of “next” frames that this processor requires for
// processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/nextFrameCount-533br
func (o VTLowLatencySuperResolutionScalerConfiguration) NextFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("nextFrameCount"))
	return int(rv)
}

// Returns the number of “previous” frames that this processor requires
// for processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/previousFrameCount-20ke2
func (o VTLowLatencySuperResolutionScalerConfiguration) PreviousFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("previousFrameCount"))
	return int(rv)
}

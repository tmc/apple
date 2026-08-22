// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTTemporalNoiseFilterConfiguration] class.
var (
	_VTTemporalNoiseFilterConfigurationClass     VTTemporalNoiseFilterConfigurationClass
	_VTTemporalNoiseFilterConfigurationClassOnce sync.Once
)

func getVTTemporalNoiseFilterConfigurationClass() VTTemporalNoiseFilterConfigurationClass {
	_VTTemporalNoiseFilterConfigurationClassOnce.Do(func() {
		_VTTemporalNoiseFilterConfigurationClass = VTTemporalNoiseFilterConfigurationClass{class: objc.GetClass("VTTemporalNoiseFilterConfiguration")}
	})
	return _VTTemporalNoiseFilterConfigurationClass
}

// GetVTTemporalNoiseFilterConfigurationClass returns the class object for VTTemporalNoiseFilterConfiguration.
func GetVTTemporalNoiseFilterConfigurationClass() VTTemporalNoiseFilterConfigurationClass {
	return getVTTemporalNoiseFilterConfigurationClass()
}

type VTTemporalNoiseFilterConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTTemporalNoiseFilterConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTTemporalNoiseFilterConfigurationClass) Alloc() VTTemporalNoiseFilterConfiguration {
	rv := objc.Send[VTTemporalNoiseFilterConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A configuration object to initiate a frame processor and use temporal
// noise-filter processor.
//
// # Overview
//
// The class properties of [VTTemporalNoiseFilterConfiguration] help to
// identify the capabilities of temporal noise filter processor on the current
// platform, prior to initiating a session. You can confirm the availability
// of temporal noise-filter processor in the current platform by checking the
// [VTTemporalNoiseFilterConfigurationClass.Supported] class property. Verify
// the processor’s capability to process source frames by ensuring that the
// dimensions are no less than
// [VTTemporalNoiseFilterConfigurationClass.MinimumDimensions] and no greater
// than [VTTemporalNoiseFilterConfigurationClass.MaximumDimensions]. Use the
// instance properties such as
// [VTTemporalNoiseFilterConfiguration.FrameSupportedPixelFormats],
// [VTTemporalNoiseFilterConfiguration.SourcePixelBufferAttributes], and
// [VTTemporalNoiseFilterConfiguration.DestinationPixelBufferAttributes] to
// ensure that the input and output pixel buffer formats and attributes of the
// processor align with the client’s specific requirements. The properties
// [VTTemporalNoiseFilterConfiguration.PreviousFrameCount] and
// [VTTemporalNoiseFilterConfiguration.NextFrameCount] represent the maximum
// number of preceding and subsequent reference frames, used in the processing
// of a source frame, to achieve optimum noise-reduction quality.
//
// # Creating a temporal noise filter configuration
//
//   - [VTTemporalNoiseFilterConfiguration.InitWithFrameWidthFrameHeightSourcePixelFormat]: Creates a new temporal noise-processor configuration.
//
// # Inspecting the configuration
//
//   - [VTTemporalNoiseFilterConfiguration.FrameWidth]: Width of source frame in pixels.
//   - [VTTemporalNoiseFilterConfiguration.FrameHeight]: Height of source frame in pixels.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration
type VTTemporalNoiseFilterConfiguration struct {
	objectivec.Object
}

// VTTemporalNoiseFilterConfigurationFromID constructs a [VTTemporalNoiseFilterConfiguration] from an objc.ID.
//
// A configuration object to initiate a frame processor and use temporal
// noise-filter processor.
func VTTemporalNoiseFilterConfigurationFromID(id objc.ID) VTTemporalNoiseFilterConfiguration {
	return VTTemporalNoiseFilterConfiguration{objectivec.Object{ID: id}}
}

// NOTE: VTTemporalNoiseFilterConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTTemporalNoiseFilterConfiguration] class.
//
// # Creating a temporal noise filter configuration
//
//   - [IVTTemporalNoiseFilterConfiguration.InitWithFrameWidthFrameHeightSourcePixelFormat]: Creates a new temporal noise-processor configuration.
//
// # Inspecting the configuration
//
//   - [IVTTemporalNoiseFilterConfiguration.FrameWidth]: Width of source frame in pixels.
//   - [IVTTemporalNoiseFilterConfiguration.FrameHeight]: Height of source frame in pixels.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration
type IVTTemporalNoiseFilterConfiguration interface {
	objectivec.IObject
	VTFrameProcessorConfiguration

	// Topic: Creating a temporal noise filter configuration

	// Creates a new temporal noise-processor configuration.
	InitWithFrameWidthFrameHeightSourcePixelFormat(frameWidth int, frameHeight int, sourcePixelFormat uint32) VTTemporalNoiseFilterConfiguration

	// Topic: Inspecting the configuration

	// Width of source frame in pixels.
	FrameWidth() int
	// Height of source frame in pixels.
	FrameHeight() int
}

// Init initializes the instance.
func (v VTTemporalNoiseFilterConfiguration) Init() VTTemporalNoiseFilterConfiguration {
	rv := objc.Send[VTTemporalNoiseFilterConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTTemporalNoiseFilterConfiguration) Autorelease() VTTemporalNoiseFilterConfiguration {
	rv := objc.Send[VTTemporalNoiseFilterConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTTemporalNoiseFilterConfiguration creates a new VTTemporalNoiseFilterConfiguration instance.
func NewVTTemporalNoiseFilterConfiguration() VTTemporalNoiseFilterConfiguration {
	class := getVTTemporalNoiseFilterConfigurationClass()
	rv := objc.Send[VTTemporalNoiseFilterConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new temporal noise-processor configuration.
//
// frameWidth: Width of source frame in pixels.
//
// frameHeight: Height of source frame in pixels.
//
// # Discussion
//
// Returns nil if frameWidth, frameHeight, or sourcePixelFormat is
// unsupported.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/init(frameWidth:frameHeight:sourcePixelFormat:)
func NewVTTemporalNoiseFilterConfigurationWithFrameWidthFrameHeightSourcePixelFormat(frameWidth int, frameHeight int, sourcePixelFormat uint32) VTTemporalNoiseFilterConfiguration {
	instance := getVTTemporalNoiseFilterConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameWidth:frameHeight:sourcePixelFormat:"), frameWidth, frameHeight, sourcePixelFormat)
	return VTTemporalNoiseFilterConfigurationFromID(rv)
}

// Creates a new temporal noise-processor configuration.
//
// frameWidth: Width of source frame in pixels.
//
// frameHeight: Height of source frame in pixels.
//
// # Discussion
//
// Returns nil if frameWidth, frameHeight, or sourcePixelFormat is
// unsupported.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/init(frameWidth:frameHeight:sourcePixelFormat:)
func (v VTTemporalNoiseFilterConfiguration) InitWithFrameWidthFrameHeightSourcePixelFormat(frameWidth int, frameHeight int, sourcePixelFormat uint32) VTTemporalNoiseFilterConfiguration {
	rv := objc.Send[VTTemporalNoiseFilterConfiguration](v.ID, objc.Sel("initWithFrameWidth:frameHeight:sourcePixelFormat:"), frameWidth, frameHeight, sourcePixelFormat)
	return rv
}

// Width of source frame in pixels.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/frameWidth
func (v VTTemporalNoiseFilterConfiguration) FrameWidth() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameWidth"))
	return rv
}

// Height of source frame in pixels.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/frameHeight
func (v VTTemporalNoiseFilterConfiguration) FrameHeight() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameHeight"))
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
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/sourcePixelBufferAttributes
func (v VTTemporalNoiseFilterConfiguration) SourcePixelBufferAttributes() foundation.INSDictionary {
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
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/destinationPixelBufferAttributes
func (v VTTemporalNoiseFilterConfiguration) DestinationPixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationPixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// Supported pixel formats for source frames for current configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/frameSupportedPixelFormats
func (v VTTemporalNoiseFilterConfiguration) FrameSupportedPixelFormats() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("frameSupportedPixelFormats"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Maximum number of future reference frames that the processor can use to
// process a source frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/nextFrameCount
func (v VTTemporalNoiseFilterConfiguration) NextFrameCount() int {
	rv := objc.Send[int](v.ID, objc.Sel("nextFrameCount"))
	return rv
}

// Maximum number of past reference frames that the processor can use to
// process a source frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/previousFrameCount
func (v VTTemporalNoiseFilterConfiguration) PreviousFrameCount() int {
	rv := objc.Send[int](v.ID, objc.Sel("previousFrameCount"))
	return rv
}

// Reports whether the system supports this processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/isSupported
func (_VTTemporalNoiseFilterConfigurationClass VTTemporalNoiseFilterConfigurationClass) IsSupported() bool {
	rv := objc.Send[bool](objc.ID(_VTTemporalNoiseFilterConfigurationClass.class), objc.Sel("isSupported"))
	return rv
}

// List of all supported pixel formats for source frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/supportedSourcePixelFormats-1w5i2
func (_VTTemporalNoiseFilterConfigurationClass VTTemporalNoiseFilterConfigurationClass) SupportedSourcePixelFormats() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](objc.ID(_VTTemporalNoiseFilterConfigurationClass.class), objc.Sel("supportedSourcePixelFormats"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// The maximum dimensions of a source frame that the processor supports.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/maximumDimensions
func (_VTTemporalNoiseFilterConfigurationClass VTTemporalNoiseFilterConfigurationClass) MaximumDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](objc.ID(_VTTemporalNoiseFilterConfigurationClass.class), objc.Sel("maximumDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// The minimum dimensions of a source frame that the processor supports.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterConfiguration/minimumDimensions
func (_VTTemporalNoiseFilterConfigurationClass VTTemporalNoiseFilterConfigurationClass) MinimumDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](objc.ID(_VTTemporalNoiseFilterConfigurationClass.class), objc.Sel("minimumDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// Protocol methods for VTFrameProcessorConfiguration

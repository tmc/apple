// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTFrameRateConversionConfiguration] class.
var (
	_VTFrameRateConversionConfigurationClass     VTFrameRateConversionConfigurationClass
	_VTFrameRateConversionConfigurationClassOnce sync.Once
)

func getVTFrameRateConversionConfigurationClass() VTFrameRateConversionConfigurationClass {
	_VTFrameRateConversionConfigurationClassOnce.Do(func() {
		_VTFrameRateConversionConfigurationClass = VTFrameRateConversionConfigurationClass{class: objc.GetClass("VTFrameRateConversionConfiguration")}
	})
	return _VTFrameRateConversionConfigurationClass
}

// GetVTFrameRateConversionConfigurationClass returns the class object for VTFrameRateConversionConfiguration.
func GetVTFrameRateConversionConfigurationClass() VTFrameRateConversionConfigurationClass {
	return getVTFrameRateConversionConfigurationClass()
}

type VTFrameRateConversionConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTFrameRateConversionConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTFrameRateConversionConfigurationClass) Alloc() VTFrameRateConversionConfiguration {
	rv := objc.Send[VTFrameRateConversionConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that enables the frame rate conversion on a frame processing
// session.
//
// # Creating a frame rate conversion configuration
//
//   - [VTFrameRateConversionConfiguration.InitWithFrameWidthFrameHeightUsePrecomputedFlowQualityPrioritizationRevision]: Creates a new frame rate conversion configuration with specified flow width and height.
//
// # Inspecting the configuration
//
//   - [VTFrameRateConversionConfiguration.FrameWidth]: The width of a source frame in pixels.
//   - [VTFrameRateConversionConfiguration.FrameHeight]: The height of a source frame in pixels.
//   - [VTFrameRateConversionConfiguration.UsePrecomputedFlow]: A Boolean value to indicates whether the optical flow will be provided by the user.
//   - [VTFrameRateConversionConfiguration.QualityPrioritization]: A value that specifies whether to prioritize quality or performance.
//
// # Inspecting revision information
//
//   - [VTFrameRateConversionConfiguration.Revision]: The specific algorithm or configuration revision to use to perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration
type VTFrameRateConversionConfiguration struct {
	objectivec.Object
}

// VTFrameRateConversionConfigurationFromID constructs a [VTFrameRateConversionConfiguration] from an objc.ID.
//
// An object that enables the frame rate conversion on a frame processing
// session.
func VTFrameRateConversionConfigurationFromID(id objc.ID) VTFrameRateConversionConfiguration {
	return VTFrameRateConversionConfiguration{objectivec.Object{ID: id}}
}

// NOTE: VTFrameRateConversionConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTFrameRateConversionConfiguration] class.
//
// # Creating a frame rate conversion configuration
//
//   - [IVTFrameRateConversionConfiguration.InitWithFrameWidthFrameHeightUsePrecomputedFlowQualityPrioritizationRevision]: Creates a new frame rate conversion configuration with specified flow width and height.
//
// # Inspecting the configuration
//
//   - [IVTFrameRateConversionConfiguration.FrameWidth]: The width of a source frame in pixels.
//   - [IVTFrameRateConversionConfiguration.FrameHeight]: The height of a source frame in pixels.
//   - [IVTFrameRateConversionConfiguration.UsePrecomputedFlow]: A Boolean value to indicates whether the optical flow will be provided by the user.
//   - [IVTFrameRateConversionConfiguration.QualityPrioritization]: A value that specifies whether to prioritize quality or performance.
//
// # Inspecting revision information
//
//   - [IVTFrameRateConversionConfiguration.Revision]: The specific algorithm or configuration revision to use to perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration
type IVTFrameRateConversionConfiguration interface {
	objectivec.IObject

	// Topic: Creating a frame rate conversion configuration

	// Creates a new frame rate conversion configuration with specified flow width and height.
	InitWithFrameWidthFrameHeightUsePrecomputedFlowQualityPrioritizationRevision(frameWidth int, frameHeight int, usePrecomputedFlow bool, qualityPrioritization VTFrameRateConversionConfigurationQualityPrioritization, revision VTFrameRateConversionConfigurationRevision) VTFrameRateConversionConfiguration

	// Topic: Inspecting the configuration

	// The width of a source frame in pixels.
	FrameWidth() int
	// The height of a source frame in pixels.
	FrameHeight() int
	// A Boolean value to indicates whether the optical flow will be provided by the user.
	UsePrecomputedFlow() bool
	// A value that specifies whether to prioritize quality or performance.
	QualityPrioritization() VTFrameRateConversionConfigurationQualityPrioritization

	// Topic: Inspecting revision information

	// The specific algorithm or configuration revision to use to perform the request.
	Revision() VTFrameRateConversionConfigurationRevision
}

// Init initializes the instance.
func (v VTFrameRateConversionConfiguration) Init() VTFrameRateConversionConfiguration {
	rv := objc.Send[VTFrameRateConversionConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTFrameRateConversionConfiguration) Autorelease() VTFrameRateConversionConfiguration {
	rv := objc.Send[VTFrameRateConversionConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTFrameRateConversionConfiguration creates a new VTFrameRateConversionConfiguration instance.
func NewVTFrameRateConversionConfiguration() VTFrameRateConversionConfiguration {
	class := getVTFrameRateConversionConfigurationClass()
	rv := objc.Send[VTFrameRateConversionConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new frame rate conversion configuration with specified flow width
// and height.
//
// frameWidth: The width of source frame in pixels. The maximum value is 8192 pixels for
// macOS, and 4096 pixels for iOS.
//
// frameHeight: The height of source frame in pixels. The maximum value is 4320 pixels for
// macOS, and 2160 pixels for iOS.
//
// usePrecomputedFlow: If true the optical flow will be provided by the user, else this
// configuration will compute the optical flow on the fly.
//
// qualityPrioritization: Instance to control quality and performance levels. See
// VEFrameRateConversionConfigurationQualityPrioritization for more
// information.
//
// revision: The specific algorithm or configuration revision that is used to perform
// the request.
//
// # Discussion
//
// Initialization fails if the dimensions are out of range or if the revision
// is unsupported.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/init(frameWidth:frameHeight:usePrecomputedFlow:qualityPrioritization:revision:)
func NewVTFrameRateConversionConfigurationWithFrameWidthFrameHeightUsePrecomputedFlowQualityPrioritizationRevision(frameWidth int, frameHeight int, usePrecomputedFlow bool, qualityPrioritization VTFrameRateConversionConfigurationQualityPrioritization, revision VTFrameRateConversionConfigurationRevision) VTFrameRateConversionConfiguration {
	instance := getVTFrameRateConversionConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameWidth:frameHeight:usePrecomputedFlow:qualityPrioritization:revision:"), frameWidth, frameHeight, usePrecomputedFlow, qualityPrioritization, revision)
	return VTFrameRateConversionConfigurationFromID(rv)
}

// Creates a new frame rate conversion configuration with specified flow width
// and height.
//
// frameWidth: The width of source frame in pixels. The maximum value is 8192 pixels for
// macOS, and 4096 pixels for iOS.
//
// frameHeight: The height of source frame in pixels. The maximum value is 4320 pixels for
// macOS, and 2160 pixels for iOS.
//
// usePrecomputedFlow: If true the optical flow will be provided by the user, else this
// configuration will compute the optical flow on the fly.
//
// qualityPrioritization: Instance to control quality and performance levels. See
// VEFrameRateConversionConfigurationQualityPrioritization for more
// information.
//
// revision: The specific algorithm or configuration revision that is used to perform
// the request.
//
// # Discussion
//
// Initialization fails if the dimensions are out of range or if the revision
// is unsupported.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/init(frameWidth:frameHeight:usePrecomputedFlow:qualityPrioritization:revision:)
func (v VTFrameRateConversionConfiguration) InitWithFrameWidthFrameHeightUsePrecomputedFlowQualityPrioritizationRevision(frameWidth int, frameHeight int, usePrecomputedFlow bool, qualityPrioritization VTFrameRateConversionConfigurationQualityPrioritization, revision VTFrameRateConversionConfigurationRevision) VTFrameRateConversionConfiguration {
	rv := objc.Send[VTFrameRateConversionConfiguration](v.ID, objc.Sel("initWithFrameWidth:frameHeight:usePrecomputedFlow:qualityPrioritization:revision:"), frameWidth, frameHeight, usePrecomputedFlow, qualityPrioritization, revision)
	return rv
}

// The width of a source frame in pixels.
//
// # Discussion
//
// The maximum value is 8192 pixels for macOS, and 4096 pixels for iOS.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/frameWidth
func (v VTFrameRateConversionConfiguration) FrameWidth() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameWidth"))
	return rv
}

// The height of a source frame in pixels.
//
// # Discussion
//
// The maximum value is 4320 pixels for macOS, and 2160 pixels for iOS.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/frameHeight
func (v VTFrameRateConversionConfiguration) FrameHeight() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameHeight"))
	return rv
}

// A Boolean value to indicates whether the optical flow will be provided by
// the user.
//
// # Discussion
//
// If `false` this configuration computes the optical flow on the fly.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/usePrecomputedFlow
func (v VTFrameRateConversionConfiguration) UsePrecomputedFlow() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("usePrecomputedFlow"))
	return rv
}

// A dictionary of pixel buffer attributes describing requirements for pixel
// buffers used as source frames and reference frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/sourcePixelBufferAttributes
func (v VTFrameRateConversionConfiguration) SourcePixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourcePixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A dictionary of pixel buffer attributes describing the requirements for
// pixel buffers used as destination frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/destinationPixelBufferAttributes
func (v VTFrameRateConversionConfiguration) DestinationPixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationPixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A value that specifies whether to prioritize quality or performance.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/qualityPrioritization-swift.property
func (v VTFrameRateConversionConfiguration) QualityPrioritization() VTFrameRateConversionConfigurationQualityPrioritization {
	rv := objc.Send[VTFrameRateConversionConfigurationQualityPrioritization](v.ID, objc.Sel("qualityPrioritization"))
	return VTFrameRateConversionConfigurationQualityPrioritization(rv)
}

// The specific algorithm or configuration revision to use to perform the
// request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/revision-swift.property
func (v VTFrameRateConversionConfiguration) Revision() VTFrameRateConversionConfigurationRevision {
	rv := objc.Send[VTFrameRateConversionConfigurationRevision](v.ID, objc.Sel("revision"))
	return VTFrameRateConversionConfigurationRevision(rv)
}

// Supported pixel formats available for source frames for current
// configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/frameSupportedPixelFormats-9w7f9
func (v VTFrameRateConversionConfiguration) FrameSupportedPixelFormats() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("frameSupportedPixelFormats"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Reports whether the system supports this processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/isSupported
func (_VTFrameRateConversionConfigurationClass VTFrameRateConversionConfigurationClass) IsSupported() bool {
	rv := objc.Send[bool](objc.ID(_VTFrameRateConversionConfigurationClass.class), objc.Sel("isSupported"))
	return rv
}

// The default revision of a particular algorithm or configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/defaultRevision
func (_VTFrameRateConversionConfigurationClass VTFrameRateConversionConfigurationClass) DefaultRevision() VTFrameRateConversionConfigurationRevision {
	rv := objc.Send[VTFrameRateConversionConfigurationRevision](objc.ID(_VTFrameRateConversionConfigurationClass.class), objc.Sel("defaultRevision"))
	return VTFrameRateConversionConfigurationRevision(rv)
}

// The collection of currently-supported algorithms or configuration revisions
// for the class of configurations.
//
// # Discussion
//
// Use this property to determine the supported revisions for each
// configuration at runtime.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/supportedRevisions
func (_VTFrameRateConversionConfigurationClass VTFrameRateConversionConfigurationClass) SupportedRevisions() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](objc.ID(_VTFrameRateConversionConfigurationClass.class), objc.Sel("supportedRevisions"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the processor supported on the
// current configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionConfiguration/processorSupported
func (_VTFrameRateConversionConfigurationClass VTFrameRateConversionConfigurationClass) ProcessorSupported() bool {
	rv := objc.Send[bool](objc.ID(_VTFrameRateConversionConfigurationClass.class), objc.Sel("processorSupported"))
	return rv
}

// Protocol methods for VTFrameProcessorConfiguration

// Returns the number of “next” frames that this processor requires for
// processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/nextFrameCount-533br
func (o VTFrameRateConversionConfiguration) NextFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("nextFrameCount"))
	return int(rv)
}

// Returns the number of “previous” frames that this processor requires
// for processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/previousFrameCount-20ke2
func (o VTFrameRateConversionConfiguration) PreviousFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("previousFrameCount"))
	return int(rv)
}

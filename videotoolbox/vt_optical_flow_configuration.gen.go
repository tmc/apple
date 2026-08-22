// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTOpticalFlowConfiguration] class.
var (
	_VTOpticalFlowConfigurationClass     VTOpticalFlowConfigurationClass
	_VTOpticalFlowConfigurationClassOnce sync.Once
)

func getVTOpticalFlowConfigurationClass() VTOpticalFlowConfigurationClass {
	_VTOpticalFlowConfigurationClassOnce.Do(func() {
		_VTOpticalFlowConfigurationClass = VTOpticalFlowConfigurationClass{class: objc.GetClass("VTOpticalFlowConfiguration")}
	})
	return _VTOpticalFlowConfigurationClass
}

// GetVTOpticalFlowConfigurationClass returns the class object for VTOpticalFlowConfiguration.
func GetVTOpticalFlowConfigurationClass() VTOpticalFlowConfigurationClass {
	return getVTOpticalFlowConfigurationClass()
}

type VTOpticalFlowConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTOpticalFlowConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTOpticalFlowConfigurationClass) Alloc() VTOpticalFlowConfiguration {
	rv := objc.Send[VTOpticalFlowConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A configuration object that enables optical flow on a frame processing
// session.
//
// # Creating an optical flow configuration
//
//   - [VTOpticalFlowConfiguration.InitWithFrameWidthFrameHeightQualityPrioritizationRevision]
//
// # Inspecting the configuration
//
//   - [VTOpticalFlowConfiguration.FrameWidth]: The width of a source frame in pixels.
//   - [VTOpticalFlowConfiguration.FrameHeight]: The height of source frame in pixels.
//   - [VTOpticalFlowConfiguration.QualityPrioritization]: A value that specifies whether to prioritize quality or performance.
//
// # Inspecting revision information
//
//   - [VTOpticalFlowConfiguration.Revision]: The specific algorithm or configuration revision that is to be used to perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration
type VTOpticalFlowConfiguration struct {
	objectivec.Object
}

// VTOpticalFlowConfigurationFromID constructs a [VTOpticalFlowConfiguration] from an objc.ID.
//
// A configuration object that enables optical flow on a frame processing
// session.
func VTOpticalFlowConfigurationFromID(id objc.ID) VTOpticalFlowConfiguration {
	return VTOpticalFlowConfiguration{objectivec.Object{ID: id}}
}

// NOTE: VTOpticalFlowConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTOpticalFlowConfiguration] class.
//
// # Creating an optical flow configuration
//
//   - [IVTOpticalFlowConfiguration.InitWithFrameWidthFrameHeightQualityPrioritizationRevision]
//
// # Inspecting the configuration
//
//   - [IVTOpticalFlowConfiguration.FrameWidth]: The width of a source frame in pixels.
//   - [IVTOpticalFlowConfiguration.FrameHeight]: The height of source frame in pixels.
//   - [IVTOpticalFlowConfiguration.QualityPrioritization]: A value that specifies whether to prioritize quality or performance.
//
// # Inspecting revision information
//
//   - [IVTOpticalFlowConfiguration.Revision]: The specific algorithm or configuration revision that is to be used to perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration
type IVTOpticalFlowConfiguration interface {
	objectivec.IObject

	// Topic: Creating an optical flow configuration

	InitWithFrameWidthFrameHeightQualityPrioritizationRevision(frameWidth int, frameHeight int, qualityPrioritization VTOpticalFlowConfigurationQualityPrioritization, revision VTOpticalFlowConfigurationRevision) VTOpticalFlowConfiguration

	// Topic: Inspecting the configuration

	// The width of a source frame in pixels.
	FrameWidth() int
	// The height of source frame in pixels.
	FrameHeight() int
	// A value that specifies whether to prioritize quality or performance.
	QualityPrioritization() VTOpticalFlowConfigurationQualityPrioritization

	// Topic: Inspecting revision information

	// The specific algorithm or configuration revision that is to be used to perform the request.
	Revision() VTOpticalFlowConfigurationRevision
}

// Init initializes the instance.
func (v VTOpticalFlowConfiguration) Init() VTOpticalFlowConfiguration {
	rv := objc.Send[VTOpticalFlowConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTOpticalFlowConfiguration) Autorelease() VTOpticalFlowConfiguration {
	rv := objc.Send[VTOpticalFlowConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTOpticalFlowConfiguration creates a new VTOpticalFlowConfiguration instance.
func NewVTOpticalFlowConfiguration() VTOpticalFlowConfiguration {
	class := getVTOpticalFlowConfigurationClass()
	rv := objc.Send[VTOpticalFlowConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// frameWidth: The width of source frame in pixels. The maximum value is 8192 pixels for
// macOS, and 4096 pixels for iOS.
//
// frameHeight: The height of source frame in pixels. The maximum value is 4320 pixels for
// macOS, and 2160 pixels for iOS.
//
// qualityPrioritization: Instance to control quality and performance levels. See
// VEFrameRateConversionConfigurationQualityPrioritization for more
// information.
//
// revision: The specific algorithm or configuration revision that is to be used to
// perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/init(frameWidth:frameHeight:qualityPrioritization:revision:)
func NewVTOpticalFlowConfigurationWithFrameWidthFrameHeightQualityPrioritizationRevision(frameWidth int, frameHeight int, qualityPrioritization VTOpticalFlowConfigurationQualityPrioritization, revision VTOpticalFlowConfigurationRevision) VTOpticalFlowConfiguration {
	instance := getVTOpticalFlowConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameWidth:frameHeight:qualityPrioritization:revision:"), frameWidth, frameHeight, qualityPrioritization, revision)
	return VTOpticalFlowConfigurationFromID(rv)
}

// frameWidth: The width of source frame in pixels. The maximum value is 8192 pixels for
// macOS, and 4096 pixels for iOS.
//
// frameHeight: The height of source frame in pixels. The maximum value is 4320 pixels for
// macOS, and 2160 pixels for iOS.
//
// qualityPrioritization: Instance to control quality and performance levels. See
// VEFrameRateConversionConfigurationQualityPrioritization for more
// information.
//
// revision: The specific algorithm or configuration revision that is to be used to
// perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/init(frameWidth:frameHeight:qualityPrioritization:revision:)
func (v VTOpticalFlowConfiguration) InitWithFrameWidthFrameHeightQualityPrioritizationRevision(frameWidth int, frameHeight int, qualityPrioritization VTOpticalFlowConfigurationQualityPrioritization, revision VTOpticalFlowConfigurationRevision) VTOpticalFlowConfiguration {
	rv := objc.Send[VTOpticalFlowConfiguration](v.ID, objc.Sel("initWithFrameWidth:frameHeight:qualityPrioritization:revision:"), frameWidth, frameHeight, qualityPrioritization, revision)
	return rv
}

// The width of a source frame in pixels.
//
// # Discussion
//
// The maximum value is 8192 pixels for macOS, and 4096 pixels for iOS.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/frameWidth
func (v VTOpticalFlowConfiguration) FrameWidth() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameWidth"))
	return rv
}

// The height of source frame in pixels.
//
// # Discussion
//
// The maximum value is 8192 pixels for macOS, and 4096 pixels for iOS.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/frameHeight
func (v VTOpticalFlowConfiguration) FrameHeight() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameHeight"))
	return rv
}

// A dictionary of pixel buffer attributes describing requirements for pixel
// buffers used as source frames and reference frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/sourcePixelBufferAttributes
func (v VTOpticalFlowConfiguration) SourcePixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourcePixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A dictionary of pixel buffer attributes describing the requirements for
// pixel buffers used as destination frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/destinationPixelBufferAttributes
func (v VTOpticalFlowConfiguration) DestinationPixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationPixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A value that specifies whether to prioritize quality or performance.
//
// # Discussion
//
// See VEFrameRateConversionConfigurationQualityPrioritization for more info.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/qualityPrioritization-swift.property
func (v VTOpticalFlowConfiguration) QualityPrioritization() VTOpticalFlowConfigurationQualityPrioritization {
	rv := objc.Send[VTOpticalFlowConfigurationQualityPrioritization](v.ID, objc.Sel("qualityPrioritization"))
	return VTOpticalFlowConfigurationQualityPrioritization(rv)
}

// The specific algorithm or configuration revision that is to be used to
// perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/revision-swift.property
func (v VTOpticalFlowConfiguration) Revision() VTOpticalFlowConfigurationRevision {
	rv := objc.Send[VTOpticalFlowConfigurationRevision](v.ID, objc.Sel("revision"))
	return VTOpticalFlowConfigurationRevision(rv)
}

// Supported pixel formats for source frames for current configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/frameSupportedPixelFormats-85kob
func (v VTOpticalFlowConfiguration) FrameSupportedPixelFormats() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("frameSupportedPixelFormats"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Reports whether the system supports this processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/isSupported
func (_VTOpticalFlowConfigurationClass VTOpticalFlowConfigurationClass) IsSupported() bool {
	rv := objc.Send[bool](objc.ID(_VTOpticalFlowConfigurationClass.class), objc.Sel("isSupported"))
	return rv
}

// The default revision of a particular algorithm or configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/defaultRevision
func (_VTOpticalFlowConfigurationClass VTOpticalFlowConfigurationClass) DefaultRevision() VTOpticalFlowConfigurationRevision {
	rv := objc.Send[VTOpticalFlowConfigurationRevision](objc.ID(_VTOpticalFlowConfigurationClass.class), objc.Sel("defaultRevision"))
	return VTOpticalFlowConfigurationRevision(rv)
}

// A boolean value that indicates whether the processor supported on the
// current configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/supportedRevisions
func (_VTOpticalFlowConfigurationClass VTOpticalFlowConfigurationClass) SupportedRevisions() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](objc.ID(_VTOpticalFlowConfigurationClass.class), objc.Sel("supportedRevisions"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// A boolean value that indicates whether the processor supported on the
// current configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowConfiguration/processorSupported
func (_VTOpticalFlowConfigurationClass VTOpticalFlowConfigurationClass) ProcessorSupported() bool {
	rv := objc.Send[bool](objc.ID(_VTOpticalFlowConfigurationClass.class), objc.Sel("processorSupported"))
	return rv
}

// Protocol methods for VTFrameProcessorConfiguration

// Returns the number of “next” frames that this processor requires for
// processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/nextFrameCount-533br
func (o VTOpticalFlowConfiguration) NextFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("nextFrameCount"))
	return int(rv)
}

// Returns the number of “previous” frames that this processor requires
// for processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/previousFrameCount-20ke2
func (o VTOpticalFlowConfiguration) PreviousFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("previousFrameCount"))
	return int(rv)
}

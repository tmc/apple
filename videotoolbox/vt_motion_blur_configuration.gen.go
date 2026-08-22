// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTMotionBlurConfiguration] class.
var (
	_VTMotionBlurConfigurationClass     VTMotionBlurConfigurationClass
	_VTMotionBlurConfigurationClassOnce sync.Once
)

func getVTMotionBlurConfigurationClass() VTMotionBlurConfigurationClass {
	_VTMotionBlurConfigurationClassOnce.Do(func() {
		_VTMotionBlurConfigurationClass = VTMotionBlurConfigurationClass{class: objc.GetClass("VTMotionBlurConfiguration")}
	})
	return _VTMotionBlurConfigurationClass
}

// GetVTMotionBlurConfigurationClass returns the class object for VTMotionBlurConfiguration.
func GetVTMotionBlurConfigurationClass() VTMotionBlurConfigurationClass {
	return getVTMotionBlurConfigurationClass()
}

type VTMotionBlurConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTMotionBlurConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTMotionBlurConfigurationClass) Alloc() VTMotionBlurConfiguration {
	rv := objc.Send[VTMotionBlurConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A configuration object to enable motion blur on a frame processing session.
//
// # Creating a motion blur configuration
//
//   - [VTMotionBlurConfiguration.InitWithFrameWidthFrameHeightUsePrecomputedFlowQualityPrioritizationRevision]: Creates a new motion blur configuration with specified flow width and height.
//
// # Inspecting the configuration
//
//   - [VTMotionBlurConfiguration.FrameWidth]: The width of a source frame in pixels.
//   - [VTMotionBlurConfiguration.FrameHeight]: The height of a source frame in pixels.
//   - [VTMotionBlurConfiguration.UsePrecomputedFlow]: A Boolean value to indicates whether the the optical flow will be provided by the user.
//   - [VTMotionBlurConfiguration.QualityPrioritization]: A value that specifies whether to prioritize quality or performance.
//
// # Inspecting revision information
//
//   - [VTMotionBlurConfiguration.Revision]: The specific algorithm or configuration revision that is to be used to perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration
type VTMotionBlurConfiguration struct {
	objectivec.Object
}

// VTMotionBlurConfigurationFromID constructs a [VTMotionBlurConfiguration] from an objc.ID.
//
// A configuration object to enable motion blur on a frame processing session.
func VTMotionBlurConfigurationFromID(id objc.ID) VTMotionBlurConfiguration {
	return VTMotionBlurConfiguration{objectivec.Object{ID: id}}
}

// NOTE: VTMotionBlurConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTMotionBlurConfiguration] class.
//
// # Creating a motion blur configuration
//
//   - [IVTMotionBlurConfiguration.InitWithFrameWidthFrameHeightUsePrecomputedFlowQualityPrioritizationRevision]: Creates a new motion blur configuration with specified flow width and height.
//
// # Inspecting the configuration
//
//   - [IVTMotionBlurConfiguration.FrameWidth]: The width of a source frame in pixels.
//   - [IVTMotionBlurConfiguration.FrameHeight]: The height of a source frame in pixels.
//   - [IVTMotionBlurConfiguration.UsePrecomputedFlow]: A Boolean value to indicates whether the the optical flow will be provided by the user.
//   - [IVTMotionBlurConfiguration.QualityPrioritization]: A value that specifies whether to prioritize quality or performance.
//
// # Inspecting revision information
//
//   - [IVTMotionBlurConfiguration.Revision]: The specific algorithm or configuration revision that is to be used to perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration
type IVTMotionBlurConfiguration interface {
	objectivec.IObject

	// Topic: Creating a motion blur configuration

	// Creates a new motion blur configuration with specified flow width and height.
	InitWithFrameWidthFrameHeightUsePrecomputedFlowQualityPrioritizationRevision(frameWidth int, frameHeight int, usePrecomputedFlow bool, qualityPrioritization VTMotionBlurConfigurationQualityPrioritization, revision VTMotionBlurConfigurationRevision) VTMotionBlurConfiguration

	// Topic: Inspecting the configuration

	// The width of a source frame in pixels.
	FrameWidth() int
	// The height of a source frame in pixels.
	FrameHeight() int
	// A Boolean value to indicates whether the the optical flow will be provided by the user.
	UsePrecomputedFlow() bool
	// A value that specifies whether to prioritize quality or performance.
	QualityPrioritization() VTMotionBlurConfigurationQualityPrioritization

	// Topic: Inspecting revision information

	// The specific algorithm or configuration revision that is to be used to perform the request.
	Revision() VTMotionBlurConfigurationRevision
}

// Init initializes the instance.
func (v VTMotionBlurConfiguration) Init() VTMotionBlurConfiguration {
	rv := objc.Send[VTMotionBlurConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTMotionBlurConfiguration) Autorelease() VTMotionBlurConfiguration {
	rv := objc.Send[VTMotionBlurConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTMotionBlurConfiguration creates a new VTMotionBlurConfiguration instance.
func NewVTMotionBlurConfiguration() VTMotionBlurConfiguration {
	class := getVTMotionBlurConfigurationClass()
	rv := objc.Send[VTMotionBlurConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new motion blur configuration with specified flow width and
// height.
//
// frameWidth: The width of the source frame in pixels. Maximum value is 8192 pixels for
// macOS, and 4096 pixels for iOS.
//
// frameHeight: The height of the source frame in pixels. Maximum value is 4320 pixels for
// macOS, and 2160 pixels for iOS.
//
// usePrecomputedFlow: If true it indicates that the optical flow will be provided by the user, if
// false this configuration will compute the optical flow on the fly.
//
// qualityPrioritization: Instance to control quality and performance levels. See
// VEMotionBlurConfigurationQualityPrioritization for more information.
//
// revision: The specific algorithm or configuration revision that is to be used to
// perform the request.
//
// # Discussion
//
// Initialization fails if the dimensions are out of range or revision is
// unsupported.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/init(frameWidth:frameHeight:usePrecomputedFlow:qualityPrioritization:revision:)
func NewVTMotionBlurConfigurationWithFrameWidthFrameHeightUsePrecomputedFlowQualityPrioritizationRevision(frameWidth int, frameHeight int, usePrecomputedFlow bool, qualityPrioritization VTMotionBlurConfigurationQualityPrioritization, revision VTMotionBlurConfigurationRevision) VTMotionBlurConfiguration {
	instance := getVTMotionBlurConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameWidth:frameHeight:usePrecomputedFlow:qualityPrioritization:revision:"), frameWidth, frameHeight, usePrecomputedFlow, qualityPrioritization, revision)
	return VTMotionBlurConfigurationFromID(rv)
}

// Creates a new motion blur configuration with specified flow width and
// height.
//
// frameWidth: The width of the source frame in pixels. Maximum value is 8192 pixels for
// macOS, and 4096 pixels for iOS.
//
// frameHeight: The height of the source frame in pixels. Maximum value is 4320 pixels for
// macOS, and 2160 pixels for iOS.
//
// usePrecomputedFlow: If true it indicates that the optical flow will be provided by the user, if
// false this configuration will compute the optical flow on the fly.
//
// qualityPrioritization: Instance to control quality and performance levels. See
// VEMotionBlurConfigurationQualityPrioritization for more information.
//
// revision: The specific algorithm or configuration revision that is to be used to
// perform the request.
//
// # Discussion
//
// Initialization fails if the dimensions are out of range or revision is
// unsupported.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/init(frameWidth:frameHeight:usePrecomputedFlow:qualityPrioritization:revision:)
func (v VTMotionBlurConfiguration) InitWithFrameWidthFrameHeightUsePrecomputedFlowQualityPrioritizationRevision(frameWidth int, frameHeight int, usePrecomputedFlow bool, qualityPrioritization VTMotionBlurConfigurationQualityPrioritization, revision VTMotionBlurConfigurationRevision) VTMotionBlurConfiguration {
	rv := objc.Send[VTMotionBlurConfiguration](v.ID, objc.Sel("initWithFrameWidth:frameHeight:usePrecomputedFlow:qualityPrioritization:revision:"), frameWidth, frameHeight, usePrecomputedFlow, qualityPrioritization, revision)
	return rv
}

// The width of a source frame in pixels.
//
// # Discussion
//
// The maximum value is 8192 pixels for macOS, and 4096 pixels for iOS.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/frameWidth
func (v VTMotionBlurConfiguration) FrameWidth() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameWidth"))
	return rv
}

// The height of a source frame in pixels.
//
// # Discussion
//
// The maximum value is 4320 pixels for macOS, and 2160 pixels for iOS.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/frameHeight
func (v VTMotionBlurConfiguration) FrameHeight() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameHeight"))
	return rv
}

// A Boolean value to indicates whether the the optical flow will be provided
// by the user.
//
// # Discussion
//
// If `false` this configuration computes the optical flow on the fly.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/usePrecomputedFlow
func (v VTMotionBlurConfiguration) UsePrecomputedFlow() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("usePrecomputedFlow"))
	return rv
}

// A dictionary of pixel buffer attributes describing requirements for pixel
// buffers used as source frames and reference frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/sourcePixelBufferAttributes
func (v VTMotionBlurConfiguration) SourcePixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourcePixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A dictionary of pixel buffer attributes describing the requirements for
// pixel buffers used as destination frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/destinationPixelBufferAttributes
func (v VTMotionBlurConfiguration) DestinationPixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationPixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A value that specifies whether to prioritize quality or performance.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/qualityPrioritization-swift.property
func (v VTMotionBlurConfiguration) QualityPrioritization() VTMotionBlurConfigurationQualityPrioritization {
	rv := objc.Send[VTMotionBlurConfigurationQualityPrioritization](v.ID, objc.Sel("qualityPrioritization"))
	return VTMotionBlurConfigurationQualityPrioritization(rv)
}

// The specific algorithm or configuration revision that is to be used to
// perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/revision-swift.property
func (v VTMotionBlurConfiguration) Revision() VTMotionBlurConfigurationRevision {
	rv := objc.Send[VTMotionBlurConfigurationRevision](v.ID, objc.Sel("revision"))
	return VTMotionBlurConfigurationRevision(rv)
}

// Available supported pixel formats for source frames for current
// configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/frameSupportedPixelFormats-5a1iv
func (v VTMotionBlurConfiguration) FrameSupportedPixelFormats() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("frameSupportedPixelFormats"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Reports whether the system supports this processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/isSupported
func (_VTMotionBlurConfigurationClass VTMotionBlurConfigurationClass) IsSupported() bool {
	rv := objc.Send[bool](objc.ID(_VTMotionBlurConfigurationClass.class), objc.Sel("isSupported"))
	return rv
}

// The default revision of a particular algorithm or configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/defaultRevision
func (_VTMotionBlurConfigurationClass VTMotionBlurConfigurationClass) DefaultRevision() VTMotionBlurConfigurationRevision {
	rv := objc.Send[VTMotionBlurConfigurationRevision](objc.ID(_VTMotionBlurConfigurationClass.class), objc.Sel("defaultRevision"))
	return VTMotionBlurConfigurationRevision(rv)
}

// The collection of currently-supported algorithms or configuration revisions
// for the class of configurations.
//
// # Discussion
//
// This property allows you to check what revisions are available for each
// configuration at runtime.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/supportedRevisions
func (_VTMotionBlurConfigurationClass VTMotionBlurConfigurationClass) SupportedRevisions() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](objc.ID(_VTMotionBlurConfigurationClass.class), objc.Sel("supportedRevisions"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the processor is supported.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurConfiguration/processorSupported
func (_VTMotionBlurConfigurationClass VTMotionBlurConfigurationClass) ProcessorSupported() bool {
	rv := objc.Send[bool](objc.ID(_VTMotionBlurConfigurationClass.class), objc.Sel("processorSupported"))
	return rv
}

// Protocol methods for VTFrameProcessorConfiguration

// Returns the number of “next” frames that this processor requires for
// processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/nextFrameCount-533br
func (o VTMotionBlurConfiguration) NextFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("nextFrameCount"))
	return int(rv)
}

// Returns the number of “previous” frames that this processor requires
// for processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/previousFrameCount-20ke2
func (o VTMotionBlurConfiguration) PreviousFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("previousFrameCount"))
	return int(rv)
}

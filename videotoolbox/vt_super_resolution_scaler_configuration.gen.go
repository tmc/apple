// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTSuperResolutionScalerConfiguration] class.
var (
	_VTSuperResolutionScalerConfigurationClass     VTSuperResolutionScalerConfigurationClass
	_VTSuperResolutionScalerConfigurationClassOnce sync.Once
)

func getVTSuperResolutionScalerConfigurationClass() VTSuperResolutionScalerConfigurationClass {
	_VTSuperResolutionScalerConfigurationClassOnce.Do(func() {
		_VTSuperResolutionScalerConfigurationClass = VTSuperResolutionScalerConfigurationClass{class: objc.GetClass("VTSuperResolutionScalerConfiguration")}
	})
	return _VTSuperResolutionScalerConfigurationClass
}

// GetVTSuperResolutionScalerConfigurationClass returns the class object for VTSuperResolutionScalerConfiguration.
func GetVTSuperResolutionScalerConfigurationClass() VTSuperResolutionScalerConfigurationClass {
	return getVTSuperResolutionScalerConfigurationClass()
}

type VTSuperResolutionScalerConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTSuperResolutionScalerConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTSuperResolutionScalerConfigurationClass) Alloc() VTSuperResolutionScalerConfiguration {
	rv := objc.Send[VTSuperResolutionScalerConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Configuration that you use to set up the super-resolution processor.
//
// # Overview
//
// This configuration enables the super-resolution processor on a
// [VTFrameProcessor] session.
//
// # Creating a super resolution scaler configuration
//
//   - [VTSuperResolutionScalerConfiguration.InitWithFrameWidthFrameHeightScaleFactorInputTypeUsePrecomputedFlowQualityPrioritizationRevision]: Creates a new super-resolution scaler processor configuration.
//
// # Inspecting the configuration
//
//   - [VTSuperResolutionScalerConfiguration.FrameWidth]: Width of source frame in pixels.
//   - [VTSuperResolutionScalerConfiguration.FrameHeight]: Height of source frame in pixels.
//   - [VTSuperResolutionScalerConfiguration.ScaleFactor]: Indicates the scale factor between input and output.
//   - [VTSuperResolutionScalerConfiguration.InputType]: Indicates the type of input.
//   - [VTSuperResolutionScalerConfiguration.PrecomputedFlow]: Indicates that you provide optical flow.
//   - [VTSuperResolutionScalerConfiguration.QualityPrioritization]: A parameter to control quality and performance levels.
//
// # Managing the configuration model
//
//   - [VTSuperResolutionScalerConfiguration.ConfigurationModelStatus]: Reports the download status of models that the system needs for the current configuration.
//   - [VTSuperResolutionScalerConfiguration.ConfigurationModelPercentageAvailable]: Returns a floating point value between 0.0 and 1.0 indicating the percentage of required model assets that have been downloaded.
//   - [VTSuperResolutionScalerConfiguration.DownloadConfigurationModelWithCompletionHandler]: Downloads models that the system needs for the current configuration.
//
// # Inspecting revision information
//
//   - [VTSuperResolutionScalerConfiguration.Revision]: The specific algorithm or configuration revision you use to perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration
type VTSuperResolutionScalerConfiguration struct {
	objectivec.Object
}

// VTSuperResolutionScalerConfigurationFromID constructs a [VTSuperResolutionScalerConfiguration] from an objc.ID.
//
// Configuration that you use to set up the super-resolution processor.
func VTSuperResolutionScalerConfigurationFromID(id objc.ID) VTSuperResolutionScalerConfiguration {
	return VTSuperResolutionScalerConfiguration{objectivec.Object{ID: id}}
}

// NOTE: VTSuperResolutionScalerConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTSuperResolutionScalerConfiguration] class.
//
// # Creating a super resolution scaler configuration
//
//   - [IVTSuperResolutionScalerConfiguration.InitWithFrameWidthFrameHeightScaleFactorInputTypeUsePrecomputedFlowQualityPrioritizationRevision]: Creates a new super-resolution scaler processor configuration.
//
// # Inspecting the configuration
//
//   - [IVTSuperResolutionScalerConfiguration.FrameWidth]: Width of source frame in pixels.
//   - [IVTSuperResolutionScalerConfiguration.FrameHeight]: Height of source frame in pixels.
//   - [IVTSuperResolutionScalerConfiguration.ScaleFactor]: Indicates the scale factor between input and output.
//   - [IVTSuperResolutionScalerConfiguration.InputType]: Indicates the type of input.
//   - [IVTSuperResolutionScalerConfiguration.PrecomputedFlow]: Indicates that you provide optical flow.
//   - [IVTSuperResolutionScalerConfiguration.QualityPrioritization]: A parameter to control quality and performance levels.
//
// # Managing the configuration model
//
//   - [IVTSuperResolutionScalerConfiguration.ConfigurationModelStatus]: Reports the download status of models that the system needs for the current configuration.
//   - [IVTSuperResolutionScalerConfiguration.ConfigurationModelPercentageAvailable]: Returns a floating point value between 0.0 and 1.0 indicating the percentage of required model assets that have been downloaded.
//   - [IVTSuperResolutionScalerConfiguration.DownloadConfigurationModelWithCompletionHandler]: Downloads models that the system needs for the current configuration.
//
// # Inspecting revision information
//
//   - [IVTSuperResolutionScalerConfiguration.Revision]: The specific algorithm or configuration revision you use to perform the request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration
type IVTSuperResolutionScalerConfiguration interface {
	objectivec.IObject

	// Topic: Creating a super resolution scaler configuration

	// Creates a new super-resolution scaler processor configuration.
	InitWithFrameWidthFrameHeightScaleFactorInputTypeUsePrecomputedFlowQualityPrioritizationRevision(frameWidth int, frameHeight int, scaleFactor int, inputType VTSuperResolutionScalerConfigurationInputType, usePrecomputedFlow bool, qualityPrioritization VTSuperResolutionScalerConfigurationQualityPrioritization, revision VTSuperResolutionScalerConfigurationRevision) VTSuperResolutionScalerConfiguration

	// Topic: Inspecting the configuration

	// Width of source frame in pixels.
	FrameWidth() int
	// Height of source frame in pixels.
	FrameHeight() int
	// Indicates the scale factor between input and output.
	ScaleFactor() int
	// Indicates the type of input.
	InputType() VTSuperResolutionScalerConfigurationInputType
	// Indicates that you provide optical flow.
	PrecomputedFlow() bool
	// A parameter to control quality and performance levels.
	QualityPrioritization() VTSuperResolutionScalerConfigurationQualityPrioritization

	// Topic: Managing the configuration model

	// Reports the download status of models that the system needs for the current configuration.
	ConfigurationModelStatus() VTSuperResolutionScalerConfigurationModelStatus
	// Returns a floating point value between 0.0 and 1.0 indicating the percentage of required model assets that have been downloaded.
	ConfigurationModelPercentageAvailable() float32
	// Downloads models that the system needs for the current configuration.
	DownloadConfigurationModelWithCompletionHandler(completionHandler ErrorHandler)

	// Topic: Inspecting revision information

	// The specific algorithm or configuration revision you use to perform the request.
	Revision() VTSuperResolutionScalerConfigurationRevision
}

// Init initializes the instance.
func (v VTSuperResolutionScalerConfiguration) Init() VTSuperResolutionScalerConfiguration {
	rv := objc.Send[VTSuperResolutionScalerConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTSuperResolutionScalerConfiguration) Autorelease() VTSuperResolutionScalerConfiguration {
	rv := objc.Send[VTSuperResolutionScalerConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTSuperResolutionScalerConfiguration creates a new VTSuperResolutionScalerConfiguration instance.
func NewVTSuperResolutionScalerConfiguration() VTSuperResolutionScalerConfiguration {
	class := getVTSuperResolutionScalerConfigurationClass()
	rv := objc.Send[VTSuperResolutionScalerConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new super-resolution scaler processor configuration.
//
// frameWidth: Width of source frame in pixels. With
// [VTSuperResolutionScalerConfigurationInputTypeVideo], maximum width is 1920
// on macOS and 1440 on iOS. With
// [VTSuperResolutionScalerConfigurationInputTypeImage], maximum width is
// 1920.
//
// frameHeight: Height of source frame in pixels. With
// [VTSuperResolutionScalerConfigurationInputTypeVideo], maximum height is
// 1080. With [VTSuperResolutionScalerConfigurationInputTypeImage], maximum
// height is 1920 on macOS and 1080 on iOS.
//
// scaleFactor: Indicates the scale factor between input and output.
//
// inputType: Indicates the type of input, either video or image.
//
// usePrecomputedFlow: Boolean value to indicate that you provide optical flow; if false, this
// configuration computes the optical flow on the fly.
//
// qualityPrioritization: A level you use to prioritize quality or performance; for more information
// about supported levels, see
// [VTSuperResolutionScalerConfiguration.QualityPrioritization].
//
// revision: The specific algorithm or configuration revision you use to perform the
// request.
//
// # Discussion
//
// This processor increases resolution of an image or video. Returns `nil` if
// dimensions are out of range or revision is unsupported.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/init(frameWidth:frameHeight:scaleFactor:inputType:usePrecomputedFlow:qualityPrioritization:revision:)
//
// [VTSuperResolutionScalerConfiguration.QualityPrioritization]: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/QualityPrioritization-swift.enum
func NewVTSuperResolutionScalerConfigurationWithFrameWidthFrameHeightScaleFactorInputTypeUsePrecomputedFlowQualityPrioritizationRevision(frameWidth int, frameHeight int, scaleFactor int, inputType VTSuperResolutionScalerConfigurationInputType, usePrecomputedFlow bool, qualityPrioritization VTSuperResolutionScalerConfigurationQualityPrioritization, revision VTSuperResolutionScalerConfigurationRevision) VTSuperResolutionScalerConfiguration {
	instance := getVTSuperResolutionScalerConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameWidth:frameHeight:scaleFactor:inputType:usePrecomputedFlow:qualityPrioritization:revision:"), frameWidth, frameHeight, scaleFactor, inputType, usePrecomputedFlow, qualityPrioritization, revision)
	return VTSuperResolutionScalerConfigurationFromID(rv)
}

// Creates a new super-resolution scaler processor configuration.
//
// frameWidth: Width of source frame in pixels. With
// [VTSuperResolutionScalerConfigurationInputTypeVideo], maximum width is 1920
// on macOS and 1440 on iOS. With
// [VTSuperResolutionScalerConfigurationInputTypeImage], maximum width is
// 1920.
//
// frameHeight: Height of source frame in pixels. With
// [VTSuperResolutionScalerConfigurationInputTypeVideo], maximum height is
// 1080. With [VTSuperResolutionScalerConfigurationInputTypeImage], maximum
// height is 1920 on macOS and 1080 on iOS.
//
// scaleFactor: Indicates the scale factor between input and output.
//
// inputType: Indicates the type of input, either video or image.
//
// usePrecomputedFlow: Boolean value to indicate that you provide optical flow; if false, this
// configuration computes the optical flow on the fly.
//
// qualityPrioritization: A level you use to prioritize quality or performance; for more information
// about supported levels, see
// [VTSuperResolutionScalerConfiguration.QualityPrioritization].
//
// revision: The specific algorithm or configuration revision you use to perform the
// request.
//
// # Discussion
//
// This processor increases resolution of an image or video. Returns `nil` if
// dimensions are out of range or revision is unsupported.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/init(frameWidth:frameHeight:scaleFactor:inputType:usePrecomputedFlow:qualityPrioritization:revision:)
//
// [VTSuperResolutionScalerConfiguration.QualityPrioritization]: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/QualityPrioritization-swift.enum
func (v VTSuperResolutionScalerConfiguration) InitWithFrameWidthFrameHeightScaleFactorInputTypeUsePrecomputedFlowQualityPrioritizationRevision(frameWidth int, frameHeight int, scaleFactor int, inputType VTSuperResolutionScalerConfigurationInputType, usePrecomputedFlow bool, qualityPrioritization VTSuperResolutionScalerConfigurationQualityPrioritization, revision VTSuperResolutionScalerConfigurationRevision) VTSuperResolutionScalerConfiguration {
	rv := objc.Send[VTSuperResolutionScalerConfiguration](v.ID, objc.Sel("initWithFrameWidth:frameHeight:scaleFactor:inputType:usePrecomputedFlow:qualityPrioritization:revision:"), frameWidth, frameHeight, scaleFactor, inputType, usePrecomputedFlow, qualityPrioritization, revision)
	return rv
}

// Downloads models that the system needs for the current configuration.
//
// # Discussion
//
// This method downloads model assets required for the current configuration
// in background. You should call this method if
// [VTSuperResolutionScalerConfiguration.ConfigurationModelStatus] is
// [VTSuperResolutionScalerConfigurationModelStatusDownloadRequired]. After
// this method is called, you can query
// [VTSuperResolutionScalerConfiguration.ConfigurationModelPercentageAvailable]
// to determine progress of model asset download process. If the download
// fails, the completion handler is invoked with an [NSError], and the
// [VTSuperResolutionScalerConfiguration.ConfigurationModelStatus] goes back
// to [VTSuperResolutionScalerConfigurationModelStatusDownloadRequired]. If
// the download succeeds, the completion handler is invoked with `nil`
// NSError.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/downloadConfigurationModel(completionHandler:)
func (v VTSuperResolutionScalerConfiguration) DownloadConfigurationModelWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](v.ID, objc.Sel("downloadConfigurationModelWithCompletionHandler:"), _block0)
}

// Width of source frame in pixels.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/frameWidth
func (v VTSuperResolutionScalerConfiguration) FrameWidth() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameWidth"))
	return rv
}

// Height of source frame in pixels.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/frameHeight
func (v VTSuperResolutionScalerConfiguration) FrameHeight() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameHeight"))
	return rv
}

// Indicates the scale factor between input and output.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/scaleFactor
func (v VTSuperResolutionScalerConfiguration) ScaleFactor() int {
	rv := objc.Send[int](v.ID, objc.Sel("scaleFactor"))
	return rv
}

// Indicates the type of input.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/inputType-swift.property
func (v VTSuperResolutionScalerConfiguration) InputType() VTSuperResolutionScalerConfigurationInputType {
	rv := objc.Send[VTSuperResolutionScalerConfigurationInputType](v.ID, objc.Sel("inputType"))
	return VTSuperResolutionScalerConfigurationInputType(rv)
}

// Indicates that you provide optical flow.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/usesPrecomputedFlow
func (v VTSuperResolutionScalerConfiguration) PrecomputedFlow() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("usesPrecomputedFlow"))
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
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/sourcePixelBufferAttributes
func (v VTSuperResolutionScalerConfiguration) SourcePixelBufferAttributes() foundation.INSDictionary {
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
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/destinationPixelBufferAttributes
func (v VTSuperResolutionScalerConfiguration) DestinationPixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationPixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A parameter to control quality and performance levels.
//
// # Discussion
//
// For more information about supported levels, see
// [VTSuperResolutionScalerConfiguration.QualityPrioritization].
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/qualityPrioritization-swift.property
//
// [VTSuperResolutionScalerConfiguration.QualityPrioritization]: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/QualityPrioritization-swift.enum
func (v VTSuperResolutionScalerConfiguration) QualityPrioritization() VTSuperResolutionScalerConfigurationQualityPrioritization {
	rv := objc.Send[VTSuperResolutionScalerConfigurationQualityPrioritization](v.ID, objc.Sel("qualityPrioritization"))
	return VTSuperResolutionScalerConfigurationQualityPrioritization(rv)
}

// Reports the download status of models that the system needs for the current
// configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/configurationModelStatus
func (v VTSuperResolutionScalerConfiguration) ConfigurationModelStatus() VTSuperResolutionScalerConfigurationModelStatus {
	rv := objc.Send[VTSuperResolutionScalerConfigurationModelStatus](v.ID, objc.Sel("configurationModelStatus"))
	return VTSuperResolutionScalerConfigurationModelStatus(rv)
}

// Returns a floating point value between 0.0 and 1.0 indicating the
// percentage of required model assets that have been downloaded.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/configurationModelPercentageAvailable
func (v VTSuperResolutionScalerConfiguration) ConfigurationModelPercentageAvailable() float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("configurationModelPercentageAvailable"))
	return rv
}

// The specific algorithm or configuration revision you use to perform the
// request.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/revision-swift.property
func (v VTSuperResolutionScalerConfiguration) Revision() VTSuperResolutionScalerConfigurationRevision {
	rv := objc.Send[VTSuperResolutionScalerConfigurationRevision](v.ID, objc.Sel("revision"))
	return VTSuperResolutionScalerConfigurationRevision(rv)
}

// Available supported pixel formats for source frames for current
// configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/frameSupportedPixelFormats
func (v VTSuperResolutionScalerConfiguration) FrameSupportedPixelFormats() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("frameSupportedPixelFormats"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Reports whether the system supports this processor.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/isSupported
func (_VTSuperResolutionScalerConfigurationClass VTSuperResolutionScalerConfigurationClass) IsSupported() bool {
	rv := objc.Send[bool](objc.ID(_VTSuperResolutionScalerConfigurationClass.class), objc.Sel("isSupported"))
	return rv
}

// Reports the set of supported scale factors to use when initializing a
// super-resolution scaler configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/supportedScaleFactors-94rce
func (_VTSuperResolutionScalerConfigurationClass VTSuperResolutionScalerConfigurationClass) SupportedScaleFactors() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](objc.ID(_VTSuperResolutionScalerConfigurationClass.class), objc.Sel("supportedScaleFactors"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Provides the default revision of a specific algorithm or configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/defaultRevision
func (_VTSuperResolutionScalerConfigurationClass VTSuperResolutionScalerConfigurationClass) DefaultRevision() VTSuperResolutionScalerConfigurationRevision {
	rv := objc.Send[VTSuperResolutionScalerConfigurationRevision](objc.ID(_VTSuperResolutionScalerConfigurationClass.class), objc.Sel("defaultRevision"))
	return VTSuperResolutionScalerConfigurationRevision(rv)
}

// Provides the collection of currently supported algorithms or configuration
// revisions for the class of configuration.
//
// # Discussion
//
// A property you use to introspect at runtime which revisions are available
// for each configuration.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerConfiguration/supportedRevisions
func (_VTSuperResolutionScalerConfigurationClass VTSuperResolutionScalerConfigurationClass) SupportedRevisions() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](objc.ID(_VTSuperResolutionScalerConfigurationClass.class), objc.Sel("supportedRevisions"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// Protocol methods for VTFrameProcessorConfiguration

// Returns the number of “next” frames that this processor requires for
// processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/nextFrameCount-533br
func (o VTSuperResolutionScalerConfiguration) NextFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("nextFrameCount"))
	return int(rv)
}

// Returns the number of “previous” frames that this processor requires
// for processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorConfiguration/previousFrameCount-20ke2
func (o VTSuperResolutionScalerConfiguration) PreviousFrameCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("previousFrameCount"))
	return int(rv)
}

// DownloadConfigurationModel is a synchronous wrapper around [VTSuperResolutionScalerConfiguration.DownloadConfigurationModelWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VTSuperResolutionScalerConfiguration) DownloadConfigurationModel(ctx context.Context) error {
	done := make(chan error, 1)
	v.DownloadConfigurationModelWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

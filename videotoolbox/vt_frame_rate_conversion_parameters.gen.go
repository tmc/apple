// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTFrameRateConversionParameters] class.
var (
	_VTFrameRateConversionParametersClass     VTFrameRateConversionParametersClass
	_VTFrameRateConversionParametersClassOnce sync.Once
)

func getVTFrameRateConversionParametersClass() VTFrameRateConversionParametersClass {
	_VTFrameRateConversionParametersClassOnce.Do(func() {
		_VTFrameRateConversionParametersClass = VTFrameRateConversionParametersClass{class: objc.GetClass("VTFrameRateConversionParameters")}
	})
	return _VTFrameRateConversionParametersClass
}

// GetVTFrameRateConversionParametersClass returns the class object for VTFrameRateConversionParameters.
func GetVTFrameRateConversionParametersClass() VTFrameRateConversionParametersClass {
	return getVTFrameRateConversionParametersClass()
}

type VTFrameRateConversionParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTFrameRateConversionParametersClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTFrameRateConversionParametersClass) Alloc() VTFrameRateConversionParameters {
	rv := objc.Send[VTFrameRateConversionParameters](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains the required input and output parameters to run a
// frame rate conversion processor on a frame.
//
// # Overview
//
// This object is used in the processWithParameters call of a VTFrameProcessor
// class. The output parameter is a destinationFrame where the output frame is
// returned as a VTFrameProcessorMutableFrame back to the caller function once
// the processing completes.
//
// The parameters within VTFrameRateConversionParameters are frame level
// parameters.
//
// # Inspecting the parameters
//
//   - [VTFrameRateConversionParameters.NextFrame]: The next source frame in presentation time order.
//   - [VTFrameRateConversionParameters.OpticalFlow]: A property that defines the optical flow for an object.
//   - [VTFrameRateConversionParameters.InterpolationPhase]: An array of floating-point values that indicate which intervals to insert a frame between the current and next frame.
//   - [VTFrameRateConversionParameters.SubmissionMode]: A value describing the processing request in a parameters submission object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters
type VTFrameRateConversionParameters struct {
	objectivec.Object
}

// VTFrameRateConversionParametersFromID constructs a [VTFrameRateConversionParameters] from an objc.ID.
//
// An object that contains the required input and output parameters to run a
// frame rate conversion processor on a frame.
func VTFrameRateConversionParametersFromID(id objc.ID) VTFrameRateConversionParameters {
	return VTFrameRateConversionParameters{objectivec.Object{ID: id}}
}

// NOTE: VTFrameRateConversionParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTFrameRateConversionParameters] class.
//
// # Inspecting the parameters
//
//   - [IVTFrameRateConversionParameters.NextFrame]: The next source frame in presentation time order.
//   - [IVTFrameRateConversionParameters.OpticalFlow]: A property that defines the optical flow for an object.
//   - [IVTFrameRateConversionParameters.InterpolationPhase]: An array of floating-point values that indicate which intervals to insert a frame between the current and next frame.
//   - [IVTFrameRateConversionParameters.SubmissionMode]: A value describing the processing request in a parameters submission object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters
type IVTFrameRateConversionParameters interface {
	objectivec.IObject

	// Topic: Inspecting the parameters

	// The next source frame in presentation time order.
	NextFrame() IVTFrameProcessorFrame
	// A property that defines the optical flow for an object.
	OpticalFlow() IVTFrameProcessorOpticalFlow
	// An array of floating-point values that indicate which intervals to insert a frame between the current and next frame.
	InterpolationPhase() []foundation.NSNumber
	// A value describing the processing request in a parameters submission object.
	SubmissionMode() VTFrameRateConversionParametersSubmissionMode

	// Creates a new frame rate conversion parameters object.
	InitWithSourceFrameNextFrameOpticalFlowInterpolationPhaseSubmissionModeDestinationFrames(sourceFrame IVTFrameProcessorFrame, nextFrame IVTFrameProcessorFrame, opticalFlow IVTFrameProcessorOpticalFlow, interpolationPhase []foundation.NSNumber, submissionMode VTFrameRateConversionParametersSubmissionMode, destinationFrame []VTFrameProcessorFrame) VTFrameRateConversionParameters
}

// Init initializes the instance.
func (v VTFrameRateConversionParameters) Init() VTFrameRateConversionParameters {
	rv := objc.Send[VTFrameRateConversionParameters](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTFrameRateConversionParameters) Autorelease() VTFrameRateConversionParameters {
	rv := objc.Send[VTFrameRateConversionParameters](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTFrameRateConversionParameters creates a new VTFrameRateConversionParameters instance.
func NewVTFrameRateConversionParameters() VTFrameRateConversionParameters {
	class := getVTFrameRateConversionParametersClass()
	rv := objc.Send[VTFrameRateConversionParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new frame rate conversion parameters object.
//
// sourceFrame: The current source frame. This must be a non nil value.
//
// nextFrame: The next source frame in presentation time order. This value can be set to
// nil for the last frame.
//
// opticalFlow: A property that defines the optical flow for an object. An optional
// VTFrameProcessorReadOnlyOpticalFlow contains the forward and backward
// optical flow of the next frame. This value is only needed if the optical
// flow is pre-computed. For the first frame it will always be nil.
//
// interpolationPhase: Array of float numbers to indicate which intervals to insert a frame
// between the current and next frame. Array size indicates how many frames
// are needed to interpolate and needs to match destinationFrames size, where
// there is one interval for each destination frame. Float number values
// should be between 0 and 1, e.g. to insert one frame in the middle, a value
// of 0.5 can be used.
//
// submissionMode: A value describing the processing request in a parameters submission
// object.
//
// destinationFrame: A caller-allocated NSArray of VTFrameProcessorFrame that contains pixel
// buffers that receive the results. The array must contain the same number of
// elements as interpolationPhase NSArray.
//
// # Discussion
//
// Initialization fails if the `sourceFrame` or `nextFrame` arguments are
// [NULL], if `sourceFrame` and reference frames don’t have the same pixel
// format, or if the `interpolationPhase` array count does not match the
// `destinationFrames` array count.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters/initWithSourceFrame:nextFrame:opticalFlow:interpolationPhase:submissionMode:destinationFrames:
func NewVTFrameRateConversionParametersWithSourceFrameNextFrameOpticalFlowInterpolationPhaseSubmissionModeDestinationFrames(sourceFrame IVTFrameProcessorFrame, nextFrame IVTFrameProcessorFrame, opticalFlow IVTFrameProcessorOpticalFlow, interpolationPhase []foundation.NSNumber, submissionMode VTFrameRateConversionParametersSubmissionMode, destinationFrame []VTFrameProcessorFrame) VTFrameRateConversionParameters {
	instance := getVTFrameRateConversionParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceFrame:nextFrame:opticalFlow:interpolationPhase:submissionMode:destinationFrames:"), sourceFrame, nextFrame, opticalFlow, objectivec.IObjectSliceToNSArray(interpolationPhase), submissionMode, objectivec.IObjectSliceToNSArray(destinationFrame))
	return VTFrameRateConversionParametersFromID(rv)
}

// Creates a new frame rate conversion parameters object.
//
// sourceFrame: The current source frame. This must be a non nil value.
//
// nextFrame: The next source frame in presentation time order. This value can be set to
// nil for the last frame.
//
// opticalFlow: A property that defines the optical flow for an object. An optional
// VTFrameProcessorReadOnlyOpticalFlow contains the forward and backward
// optical flow of the next frame. This value is only needed if the optical
// flow is pre-computed. For the first frame it will always be nil.
//
// interpolationPhase: Array of float numbers to indicate which intervals to insert a frame
// between the current and next frame. Array size indicates how many frames
// are needed to interpolate and needs to match destinationFrames size, where
// there is one interval for each destination frame. Float number values
// should be between 0 and 1, e.g. to insert one frame in the middle, a value
// of 0.5 can be used.
//
// submissionMode: A value describing the processing request in a parameters submission
// object.
//
// destinationFrame: A caller-allocated NSArray of VTFrameProcessorFrame that contains pixel
// buffers that receive the results. The array must contain the same number of
// elements as interpolationPhase NSArray.
//
// # Discussion
//
// Initialization fails if the `sourceFrame` or `nextFrame` arguments are
// [NULL], if `sourceFrame` and reference frames don’t have the same pixel
// format, or if the `interpolationPhase` array count does not match the
// `destinationFrames` array count.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters/initWithSourceFrame:nextFrame:opticalFlow:interpolationPhase:submissionMode:destinationFrames:
func (v VTFrameRateConversionParameters) InitWithSourceFrameNextFrameOpticalFlowInterpolationPhaseSubmissionModeDestinationFrames(sourceFrame IVTFrameProcessorFrame, nextFrame IVTFrameProcessorFrame, opticalFlow IVTFrameProcessorOpticalFlow, interpolationPhase []foundation.NSNumber, submissionMode VTFrameRateConversionParametersSubmissionMode, destinationFrame []VTFrameProcessorFrame) VTFrameRateConversionParameters {
	rv := objc.Send[VTFrameRateConversionParameters](v.ID, objc.Sel("initWithSourceFrame:nextFrame:opticalFlow:interpolationPhase:submissionMode:destinationFrames:"), sourceFrame, nextFrame, opticalFlow, objectivec.IObjectSliceToNSArray(interpolationPhase), submissionMode, objectivec.IObjectSliceToNSArray(destinationFrame))
	return rv
}

// The current source frame.
//
// # Discussion
//
// This value must be non-nil.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters/sourceFrame
func (v VTFrameRateConversionParameters) SourceFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourceFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// The next source frame in presentation time order.
//
// # Discussion
//
// For the last frame this value will be nil.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters/nextFrame
func (v VTFrameRateConversionParameters) NextFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("nextFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// A property that defines the optical flow for an object.
//
// # Discussion
//
// An optional VTFrameProcessorReadOnlyOpticalFlow contains the forward and
// backward optical flow of the next frame. This value is only needed if the
// optical flow is pre-computed. For the first frame it will always be nil.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters/opticalFlow
func (v VTFrameRateConversionParameters) OpticalFlow() IVTFrameProcessorOpticalFlow {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("opticalFlow"))
	return VTFrameProcessorOpticalFlowFromID(objc.ID(rv))
}

// An array of floating-point values that indicate which intervals to insert a
// frame between the current and next frame.
//
// # Discussion
//
// Array size indicates how many frames are needed to interpolate and needs to
// match destinationFrames array size, where there is one interval for each
// destination frame. Float number values should be between 0 and 1, e.g. to
// insert one frame in the middle, a value of 0.5 can be used.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters/interpolationPhase-6wdns
func (v VTFrameRateConversionParameters) InterpolationPhase() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("interpolationPhase"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// A value describing the processing request in a parameters submission
// object.
//
// # Discussion
//
// Set to VTFrameRateConversionParametersSubmissionModeSequential to indicate
// that the current submission follows the presentation time order without
// jumping or skipping when compared to the previous submission. Using the
// submission mode sequential will yield better performance. Set to
// VTFrameRateConversionParametersSubmissionModeRandom to indicate a skip or a
// jump in frame sequence. If the submission mode random is set, the internal
// cache will be cleared during the processWithParameters call.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters/submissionMode-swift.property
func (v VTFrameRateConversionParameters) SubmissionMode() VTFrameRateConversionParametersSubmissionMode {
	rv := objc.Send[VTFrameRateConversionParametersSubmissionMode](v.ID, objc.Sel("submissionMode"))
	return VTFrameRateConversionParametersSubmissionMode(rv)
}

// A caller-allocated array of frames that contains the pixel buffers to
// receive the results.
//
// # Discussion
//
// This array must contain the same number of elements as the
// `interpolationPhase` property.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameRateConversionParameters/destinationFrames
func (v VTFrameRateConversionParameters) DestinationFrames() []VTFrameProcessorFrame {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("destinationFrames"))
	return objc.ConvertSlice(rv, func(id objc.ID) VTFrameProcessorFrame {
		return VTFrameProcessorFrameFromID(id)
	})
}

// Protocol methods for VTFrameProcessorParameters

// Destination frame that contains the destination frame for processors which
// output a single processed frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrame-3im3l
func (o VTFrameRateConversionParameters) DestinationFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("destinationFrame"))
	return VTFrameProcessorFrameFromID(rv)
}

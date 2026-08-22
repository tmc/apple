// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTSuperResolutionScalerParameters] class.
var (
	_VTSuperResolutionScalerParametersClass     VTSuperResolutionScalerParametersClass
	_VTSuperResolutionScalerParametersClassOnce sync.Once
)

func getVTSuperResolutionScalerParametersClass() VTSuperResolutionScalerParametersClass {
	_VTSuperResolutionScalerParametersClassOnce.Do(func() {
		_VTSuperResolutionScalerParametersClass = VTSuperResolutionScalerParametersClass{class: objc.GetClass("VTSuperResolutionScalerParameters")}
	})
	return _VTSuperResolutionScalerParametersClass
}

// GetVTSuperResolutionScalerParametersClass returns the class object for VTSuperResolutionScalerParameters.
func GetVTSuperResolutionScalerParametersClass() VTSuperResolutionScalerParametersClass {
	return getVTSuperResolutionScalerParametersClass()
}

type VTSuperResolutionScalerParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTSuperResolutionScalerParametersClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTSuperResolutionScalerParametersClass) Alloc() VTSuperResolutionScalerParameters {
	rv := objc.Send[VTSuperResolutionScalerParameters](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains both input and output parameters that the
// super-resolution processor needs to run on a frame.
//
// # Overview
//
// Use this object in the `processWithParameters` call of the
// [VTFrameProcessor] class. The output parameter for this class is
// `destinationFrame`, where the processor returns the output frame (as
// [VTFrameProcessorFrame]) back to you once `processWithParameters`
// completes.
//
// [VTSuperResolutionScalerParameters] are frame-level parameters.
//
// # Creating a parameters object
//
//   - [VTSuperResolutionScalerParameters.InitWithSourceFramePreviousFramePreviousOutputFrameOpticalFlowSubmissionModeDestinationFrame]: Creates a new super-resolution scaler parameters instance.
//
// # Inspecting the parameters
//
//   - [VTSuperResolutionScalerParameters.PreviousFrame]: Previous source frame in presentation time order, which is `nil` for the first frame.
//   - [VTSuperResolutionScalerParameters.PreviousOutputFrame]: Previous output frame in presentation time order, which is `nil` for the first frame.
//   - [VTSuperResolutionScalerParameters.OpticalFlow]: Optional object that contains forward and backward optical flow with the previous frame.
//   - [VTSuperResolutionScalerParameters.SubmissionMode]: Ordering of the input frames in this submission relative to the previous submission.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters
type VTSuperResolutionScalerParameters struct {
	objectivec.Object
}

// VTSuperResolutionScalerParametersFromID constructs a [VTSuperResolutionScalerParameters] from an objc.ID.
//
// An object that contains both input and output parameters that the
// super-resolution processor needs to run on a frame.
func VTSuperResolutionScalerParametersFromID(id objc.ID) VTSuperResolutionScalerParameters {
	return VTSuperResolutionScalerParameters{objectivec.Object{ID: id}}
}

// NOTE: VTSuperResolutionScalerParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTSuperResolutionScalerParameters] class.
//
// # Creating a parameters object
//
//   - [IVTSuperResolutionScalerParameters.InitWithSourceFramePreviousFramePreviousOutputFrameOpticalFlowSubmissionModeDestinationFrame]: Creates a new super-resolution scaler parameters instance.
//
// # Inspecting the parameters
//
//   - [IVTSuperResolutionScalerParameters.PreviousFrame]: Previous source frame in presentation time order, which is `nil` for the first frame.
//   - [IVTSuperResolutionScalerParameters.PreviousOutputFrame]: Previous output frame in presentation time order, which is `nil` for the first frame.
//   - [IVTSuperResolutionScalerParameters.OpticalFlow]: Optional object that contains forward and backward optical flow with the previous frame.
//   - [IVTSuperResolutionScalerParameters.SubmissionMode]: Ordering of the input frames in this submission relative to the previous submission.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters
type IVTSuperResolutionScalerParameters interface {
	objectivec.IObject

	// Topic: Creating a parameters object

	// Creates a new super-resolution scaler parameters instance.
	InitWithSourceFramePreviousFramePreviousOutputFrameOpticalFlowSubmissionModeDestinationFrame(sourceFrame IVTFrameProcessorFrame, previousFrame IVTFrameProcessorFrame, previousOutputFrame IVTFrameProcessorFrame, opticalFlow IVTFrameProcessorOpticalFlow, submissionMode VTSuperResolutionScalerParametersSubmissionMode, destinationFrame IVTFrameProcessorFrame) VTSuperResolutionScalerParameters

	// Topic: Inspecting the parameters

	// Previous source frame in presentation time order, which is `nil` for the first frame.
	PreviousFrame() IVTFrameProcessorFrame
	// Previous output frame in presentation time order, which is `nil` for the first frame.
	PreviousOutputFrame() IVTFrameProcessorFrame
	// Optional object that contains forward and backward optical flow with the previous frame.
	OpticalFlow() IVTFrameProcessorOpticalFlow
	// Ordering of the input frames in this submission relative to the previous submission.
	SubmissionMode() VTSuperResolutionScalerParametersSubmissionMode
}

// Init initializes the instance.
func (v VTSuperResolutionScalerParameters) Init() VTSuperResolutionScalerParameters {
	rv := objc.Send[VTSuperResolutionScalerParameters](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTSuperResolutionScalerParameters) Autorelease() VTSuperResolutionScalerParameters {
	rv := objc.Send[VTSuperResolutionScalerParameters](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTSuperResolutionScalerParameters creates a new VTSuperResolutionScalerParameters instance.
func NewVTSuperResolutionScalerParameters() VTSuperResolutionScalerParameters {
	class := getVTSuperResolutionScalerParametersClass()
	rv := objc.Send[VTSuperResolutionScalerParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new super-resolution scaler parameters instance.
//
// sourceFrame: Current source frame; must be non `nil`.
//
// previousFrame: The previous source frame in presentation time order. For the first frame
// you can set this to `nil`.
//
// previousOutputFrame: The previous output frame in presentation time order. For the first frame
// you can set this to `nil`.
//
// opticalFlow: Optional [VTFrameProcessorOpticalFlow] object that contains forward and
// backward optical flow between the `sourceFrame` and `previousFrame`. You
// only need this if optical flow is pre-computed.
//
// submissionMode: Provides a hint to let the processor know whether you are submitting frames
// in presentation sequence. For more information about supported modes see
// [VTSuperResolutionScalerParameters.SubmissionMode].
//
// destinationFrame: User-allocated pixel buffer that receives the results.
//
// # Discussion
//
// Returns `nil` if `sourceFrame` or `destinationFrame` is `nil`, or if
// `sourceFrame` and reference frames have different pixel formats.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/init(sourceFrame:previousFrame:previousOutputFrame:opticalFlow:submissionMode:destinationFrame:)
//
// [VTSuperResolutionScalerParameters.SubmissionMode]: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/SubmissionMode-swift.enum
func NewVTSuperResolutionScalerParametersWithSourceFramePreviousFramePreviousOutputFrameOpticalFlowSubmissionModeDestinationFrame(sourceFrame IVTFrameProcessorFrame, previousFrame IVTFrameProcessorFrame, previousOutputFrame IVTFrameProcessorFrame, opticalFlow IVTFrameProcessorOpticalFlow, submissionMode VTSuperResolutionScalerParametersSubmissionMode, destinationFrame IVTFrameProcessorFrame) VTSuperResolutionScalerParameters {
	instance := getVTSuperResolutionScalerParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceFrame:previousFrame:previousOutputFrame:opticalFlow:submissionMode:destinationFrame:"), sourceFrame, previousFrame, previousOutputFrame, opticalFlow, submissionMode, destinationFrame)
	return VTSuperResolutionScalerParametersFromID(rv)
}

// Creates a new super-resolution scaler parameters instance.
//
// sourceFrame: Current source frame; must be non `nil`.
//
// previousFrame: The previous source frame in presentation time order. For the first frame
// you can set this to `nil`.
//
// previousOutputFrame: The previous output frame in presentation time order. For the first frame
// you can set this to `nil`.
//
// opticalFlow: Optional [VTFrameProcessorOpticalFlow] object that contains forward and
// backward optical flow between the `sourceFrame` and `previousFrame`. You
// only need this if optical flow is pre-computed.
//
// submissionMode: Provides a hint to let the processor know whether you are submitting frames
// in presentation sequence. For more information about supported modes see
// [VTSuperResolutionScalerParameters.SubmissionMode].
//
// destinationFrame: User-allocated pixel buffer that receives the results.
//
// # Discussion
//
// Returns `nil` if `sourceFrame` or `destinationFrame` is `nil`, or if
// `sourceFrame` and reference frames have different pixel formats.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/init(sourceFrame:previousFrame:previousOutputFrame:opticalFlow:submissionMode:destinationFrame:)
//
// [VTSuperResolutionScalerParameters.SubmissionMode]: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/SubmissionMode-swift.enum
func (v VTSuperResolutionScalerParameters) InitWithSourceFramePreviousFramePreviousOutputFrameOpticalFlowSubmissionModeDestinationFrame(sourceFrame IVTFrameProcessorFrame, previousFrame IVTFrameProcessorFrame, previousOutputFrame IVTFrameProcessorFrame, opticalFlow IVTFrameProcessorOpticalFlow, submissionMode VTSuperResolutionScalerParametersSubmissionMode, destinationFrame IVTFrameProcessorFrame) VTSuperResolutionScalerParameters {
	rv := objc.Send[VTSuperResolutionScalerParameters](v.ID, objc.Sel("initWithSourceFrame:previousFrame:previousOutputFrame:opticalFlow:submissionMode:destinationFrame:"), sourceFrame, previousFrame, previousOutputFrame, opticalFlow, submissionMode, destinationFrame)
	return rv
}

// Current source frame, which must be non `nil`.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/sourceFrame
func (v VTSuperResolutionScalerParameters) SourceFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourceFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Previous source frame in presentation time order, which is `nil` for the
// first frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/previousFrame
func (v VTSuperResolutionScalerParameters) PreviousFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("previousFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Previous output frame in presentation time order, which is `nil` for the
// first frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/previousOutputFrame
func (v VTSuperResolutionScalerParameters) PreviousOutputFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("previousOutputFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Optional object that contains forward and backward optical flow with the
// previous frame.
//
// # Discussion
//
// You only need this if optical flow is pre-computed. For the first frame
// this is `nil`.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/opticalFlow
func (v VTSuperResolutionScalerParameters) OpticalFlow() IVTFrameProcessorOpticalFlow {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("opticalFlow"))
	return VTFrameProcessorOpticalFlowFromID(objc.ID(rv))
}

// Ordering of the input frames in this submission relative to the previous
// submission.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/submissionMode-swift.property
func (v VTSuperResolutionScalerParameters) SubmissionMode() VTSuperResolutionScalerParametersSubmissionMode {
	rv := objc.Send[VTSuperResolutionScalerParametersSubmissionMode](v.ID, objc.Sel("submissionMode"))
	return VTSuperResolutionScalerParametersSubmissionMode(rv)
}

// Destination frame that contains user-allocated pixel buffer that receives
// the results.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSuperResolutionScalerParameters/destinationFrame
func (v VTSuperResolutionScalerParameters) DestinationFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Protocol methods for VTFrameProcessorParameters

// Array of destination frames for processors which may output more than one
// processed frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrames-8yges
func (o VTSuperResolutionScalerParameters) DestinationFrames() []VTFrameProcessorFrame {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("destinationFrames"))
	result := make([]VTFrameProcessorFrame, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = VTFrameProcessorFrameFromID(id)
	}
	return result
}

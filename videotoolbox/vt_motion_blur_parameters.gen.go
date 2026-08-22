// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTMotionBlurParameters] class.
var (
	_VTMotionBlurParametersClass     VTMotionBlurParametersClass
	_VTMotionBlurParametersClassOnce sync.Once
)

func getVTMotionBlurParametersClass() VTMotionBlurParametersClass {
	_VTMotionBlurParametersClassOnce.Do(func() {
		_VTMotionBlurParametersClass = VTMotionBlurParametersClass{class: objc.GetClass("VTMotionBlurParameters")}
	})
	return _VTMotionBlurParametersClass
}

// GetVTMotionBlurParametersClass returns the class object for VTMotionBlurParameters.
func GetVTMotionBlurParametersClass() VTMotionBlurParametersClass {
	return getVTMotionBlurParametersClass()
}

type VTMotionBlurParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTMotionBlurParametersClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTMotionBlurParametersClass) Alloc() VTMotionBlurParameters {
	rv := objc.Send[VTMotionBlurParameters](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// This object contains both input and output parameters necessary to run the
// motion blur processor on a frame.
//
// # Overview
//
// This object is used in the processWithParameters call of the
// VTFrameProcessor class. The output parameter is a destinationFrame where
// the output frame is returned as a VTFrameProcessorFrame back to the caller
// function once processWithParameters completes.
//
// The parameters within VTMotionBlurParameters are frame level parameters.
//
// # Creating a parameters object
//
//   - [VTMotionBlurParameters.InitWithSourceFrameNextFramePreviousFrameNextOpticalFlowPreviousOpticalFlowMotionBlurStrengthSubmissionModeDestinationFrame]
//
// # Inspecting the parameters
//
//   - [VTMotionBlurParameters.NextFrame]: The next source frame in presentation time order.
//   - [VTMotionBlurParameters.PreviousFrame]: The previous source frame in presentation time order.
//   - [VTMotionBlurParameters.MotionBlurStrength]: A value that indicates the strength of blur to apply.
//   - [VTMotionBlurParameters.NextOpticalFlow]: Optional optical flow object that contains forward and backward optical flow with the next frame.
//   - [VTMotionBlurParameters.PreviousOpticalFlow]: Optional optical flow object that contains forward and backward optical flow with the previous frame.
//   - [VTMotionBlurParameters.SubmissionMode]: A value describing the processing request in a parameters submission object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters
type VTMotionBlurParameters struct {
	objectivec.Object
}

// VTMotionBlurParametersFromID constructs a [VTMotionBlurParameters] from an objc.ID.
//
// This object contains both input and output parameters necessary to run the
// motion blur processor on a frame.
func VTMotionBlurParametersFromID(id objc.ID) VTMotionBlurParameters {
	return VTMotionBlurParameters{objectivec.Object{ID: id}}
}

// NOTE: VTMotionBlurParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTMotionBlurParameters] class.
//
// # Creating a parameters object
//
//   - [IVTMotionBlurParameters.InitWithSourceFrameNextFramePreviousFrameNextOpticalFlowPreviousOpticalFlowMotionBlurStrengthSubmissionModeDestinationFrame]
//
// # Inspecting the parameters
//
//   - [IVTMotionBlurParameters.NextFrame]: The next source frame in presentation time order.
//   - [IVTMotionBlurParameters.PreviousFrame]: The previous source frame in presentation time order.
//   - [IVTMotionBlurParameters.MotionBlurStrength]: A value that indicates the strength of blur to apply.
//   - [IVTMotionBlurParameters.NextOpticalFlow]: Optional optical flow object that contains forward and backward optical flow with the next frame.
//   - [IVTMotionBlurParameters.PreviousOpticalFlow]: Optional optical flow object that contains forward and backward optical flow with the previous frame.
//   - [IVTMotionBlurParameters.SubmissionMode]: A value describing the processing request in a parameters submission object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters
type IVTMotionBlurParameters interface {
	objectivec.IObject

	// Topic: Creating a parameters object

	InitWithSourceFrameNextFramePreviousFrameNextOpticalFlowPreviousOpticalFlowMotionBlurStrengthSubmissionModeDestinationFrame(sourceFrame IVTFrameProcessorFrame, nextFrame IVTFrameProcessorFrame, previousFrame IVTFrameProcessorFrame, nextOpticalFlow IVTFrameProcessorOpticalFlow, previousOpticalFlow IVTFrameProcessorOpticalFlow, motionBlurStrength int, submissionMode VTMotionBlurParametersSubmissionMode, destinationFrame IVTFrameProcessorFrame) VTMotionBlurParameters

	// Topic: Inspecting the parameters

	// The next source frame in presentation time order.
	NextFrame() IVTFrameProcessorFrame
	// The previous source frame in presentation time order.
	PreviousFrame() IVTFrameProcessorFrame
	// A value that indicates the strength of blur to apply.
	MotionBlurStrength() int
	// Optional optical flow object that contains forward and backward optical flow with the next frame.
	NextOpticalFlow() IVTFrameProcessorOpticalFlow
	// Optional optical flow object that contains forward and backward optical flow with the previous frame.
	PreviousOpticalFlow() IVTFrameProcessorOpticalFlow
	// A value describing the processing request in a parameters submission object.
	SubmissionMode() VTMotionBlurParametersSubmissionMode
}

// Init initializes the instance.
func (v VTMotionBlurParameters) Init() VTMotionBlurParameters {
	rv := objc.Send[VTMotionBlurParameters](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTMotionBlurParameters) Autorelease() VTMotionBlurParameters {
	rv := objc.Send[VTMotionBlurParameters](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTMotionBlurParameters creates a new VTMotionBlurParameters instance.
func NewVTMotionBlurParameters() VTMotionBlurParameters {
	class := getVTMotionBlurParametersClass()
	rv := objc.Send[VTMotionBlurParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// sourceFrame: The current source frame. This must be a non-nil value.
//
// nextFrame: The next source frame in presentation time order. This value can be set to
// nil for the last frame.
//
// previousFrame: The previous source frame in presentation time order. This value can be set
// to nil for the first frame.
//
// nextOpticalFlow: Optional optical flow object that contains forward and backward optical
// flow with the next frame. For the last frame this will always be nil. This
// object is only needed if optical flow is pre-computed.
//
// previousOpticalFlow: Optional optical flow object that contains forward and backward optical
// flow with the next frame. For the first frame this will always be nil. This
// object is only needed if optical flow is pre-computed.
//
// motionBlurStrength: Number to indicate the strength of blur to apply. This NSInteger number
// ranges from 1 to 100. The default value is 50.
//
// submissionMode: A value describing the processing request in a parameters submission
// object.
//
// destinationFrame: A user-allocated pixel buffer that will receive the results.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/init(sourceFrame:nextFrame:previousFrame:nextOpticalFlow:previousOpticalFlow:motionBlurStrength:submissionMode:destinationFrame:)
func NewVTMotionBlurParametersWithSourceFrameNextFramePreviousFrameNextOpticalFlowPreviousOpticalFlowMotionBlurStrengthSubmissionModeDestinationFrame(sourceFrame IVTFrameProcessorFrame, nextFrame IVTFrameProcessorFrame, previousFrame IVTFrameProcessorFrame, nextOpticalFlow IVTFrameProcessorOpticalFlow, previousOpticalFlow IVTFrameProcessorOpticalFlow, motionBlurStrength int, submissionMode VTMotionBlurParametersSubmissionMode, destinationFrame IVTFrameProcessorFrame) VTMotionBlurParameters {
	instance := getVTMotionBlurParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceFrame:nextFrame:previousFrame:nextOpticalFlow:previousOpticalFlow:motionBlurStrength:submissionMode:destinationFrame:"), sourceFrame, nextFrame, previousFrame, nextOpticalFlow, previousOpticalFlow, motionBlurStrength, submissionMode, destinationFrame)
	return VTMotionBlurParametersFromID(rv)
}

// sourceFrame: The current source frame. This must be a non-nil value.
//
// nextFrame: The next source frame in presentation time order. This value can be set to
// nil for the last frame.
//
// previousFrame: The previous source frame in presentation time order. This value can be set
// to nil for the first frame.
//
// nextOpticalFlow: Optional optical flow object that contains forward and backward optical
// flow with the next frame. For the last frame this will always be nil. This
// object is only needed if optical flow is pre-computed.
//
// previousOpticalFlow: Optional optical flow object that contains forward and backward optical
// flow with the next frame. For the first frame this will always be nil. This
// object is only needed if optical flow is pre-computed.
//
// motionBlurStrength: Number to indicate the strength of blur to apply. This NSInteger number
// ranges from 1 to 100. The default value is 50.
//
// submissionMode: A value describing the processing request in a parameters submission
// object.
//
// destinationFrame: A user-allocated pixel buffer that will receive the results.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/init(sourceFrame:nextFrame:previousFrame:nextOpticalFlow:previousOpticalFlow:motionBlurStrength:submissionMode:destinationFrame:)
func (v VTMotionBlurParameters) InitWithSourceFrameNextFramePreviousFrameNextOpticalFlowPreviousOpticalFlowMotionBlurStrengthSubmissionModeDestinationFrame(sourceFrame IVTFrameProcessorFrame, nextFrame IVTFrameProcessorFrame, previousFrame IVTFrameProcessorFrame, nextOpticalFlow IVTFrameProcessorOpticalFlow, previousOpticalFlow IVTFrameProcessorOpticalFlow, motionBlurStrength int, submissionMode VTMotionBlurParametersSubmissionMode, destinationFrame IVTFrameProcessorFrame) VTMotionBlurParameters {
	rv := objc.Send[VTMotionBlurParameters](v.ID, objc.Sel("initWithSourceFrame:nextFrame:previousFrame:nextOpticalFlow:previousOpticalFlow:motionBlurStrength:submissionMode:destinationFrame:"), sourceFrame, nextFrame, previousFrame, nextOpticalFlow, previousOpticalFlow, motionBlurStrength, submissionMode, destinationFrame)
	return rv
}

// The current source frame.
//
// # Discussion
//
// This value must be non nil.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/sourceFrame
func (v VTMotionBlurParameters) SourceFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourceFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// The next source frame in presentation time order.
//
// # Discussion
//
// This value can be set to nil for the last frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/nextFrame
func (v VTMotionBlurParameters) NextFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("nextFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// The previous source frame in presentation time order.
//
// # Discussion
//
// This value can be set to nil for the first frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/previousFrame
func (v VTMotionBlurParameters) PreviousFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("previousFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// A value that indicates the strength of blur to apply.
//
// # Discussion
//
// The supported range for this value is from 1 and 100. The default value is
// 50.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/motionBlurStrength
func (v VTMotionBlurParameters) MotionBlurStrength() int {
	rv := objc.Send[int](v.ID, objc.Sel("motionBlurStrength"))
	return rv
}

// Optional optical flow object that contains forward and backward optical
// flow with the next frame.
//
// # Discussion
//
// For the last frame this will always be nil. This object is only needed if
// optical flow is pre-computed.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/nextOpticalFlow
func (v VTMotionBlurParameters) NextOpticalFlow() IVTFrameProcessorOpticalFlow {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("nextOpticalFlow"))
	return VTFrameProcessorOpticalFlowFromID(objc.ID(rv))
}

// Optional optical flow object that contains forward and backward optical
// flow with the previous frame.
//
// # Discussion
//
// For the first frame this will always be nil. This object is only needed if
// optical flow is pre-computed.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/previousOpticalFlow
func (v VTMotionBlurParameters) PreviousOpticalFlow() IVTFrameProcessorOpticalFlow {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("previousOpticalFlow"))
	return VTFrameProcessorOpticalFlowFromID(objc.ID(rv))
}

// A value describing the processing request in a parameters submission
// object.
//
// # Discussion
//
// Set to VTMotionBlurParametersSubmissionModeSequential to indicate that the
// current submission follows the presentation time order without jumping or
// skipping when compared to the previous submission. Using the submission
// mode sequential will yield better performance. Set to
// VTMotionBlurParametersSubmissionModeRandom to indicate a skip or a jump in
// frame sequence. If submission mode random is set, the internal cache will
// be cleared during the processWithParameters call.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/submissionMode-swift.property
func (v VTMotionBlurParameters) SubmissionMode() VTMotionBlurParametersSubmissionMode {
	rv := objc.Send[VTMotionBlurParametersSubmissionMode](v.ID, objc.Sel("submissionMode"))
	return VTMotionBlurParametersSubmissionMode(rv)
}

// A user-allocated pixel buffer that receives the results.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionBlurParameters/destinationFrame
func (v VTMotionBlurParameters) DestinationFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Protocol methods for VTFrameProcessorParameters

// Array of destination frames for processors which may output more than one
// processed frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrames-8yges
func (o VTMotionBlurParameters) DestinationFrames() []VTFrameProcessorFrame {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("destinationFrames"))
	result := make([]VTFrameProcessorFrame, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = VTFrameProcessorFrameFromID(id)
	}
	return result
}

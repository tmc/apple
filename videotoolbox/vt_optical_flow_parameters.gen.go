// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTOpticalFlowParameters] class.
var (
	_VTOpticalFlowParametersClass     VTOpticalFlowParametersClass
	_VTOpticalFlowParametersClassOnce sync.Once
)

func getVTOpticalFlowParametersClass() VTOpticalFlowParametersClass {
	_VTOpticalFlowParametersClassOnce.Do(func() {
		_VTOpticalFlowParametersClass = VTOpticalFlowParametersClass{class: objc.GetClass("VTOpticalFlowParameters")}
	})
	return _VTOpticalFlowParametersClass
}

// GetVTOpticalFlowParametersClass returns the class object for VTOpticalFlowParameters.
func GetVTOpticalFlowParametersClass() VTOpticalFlowParametersClass {
	return getVTOpticalFlowParametersClass()
}

type VTOpticalFlowParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTOpticalFlowParametersClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTOpticalFlowParametersClass) Alloc() VTOpticalFlowParameters {
	rv := objc.Send[VTOpticalFlowParameters](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes frame-level optical flow parameters.
//
// # Creating a parameters object
//
//   - [VTOpticalFlowParameters.InitWithSourceFrameNextFrameSubmissionModeDestinationOpticalFlow]
//
// # Inspecting the parameters
//
//   - [VTOpticalFlowParameters.NextFrame]: The next source frame in presentation time order.
//   - [VTOpticalFlowParameters.DestinationOpticalFlow]: A user allocated mutable optical flow that will receive the results.
//   - [VTOpticalFlowParameters.SubmissionMode]: A value describing the processing request in a parameters submission object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowParameters
type VTOpticalFlowParameters struct {
	objectivec.Object
}

// VTOpticalFlowParametersFromID constructs a [VTOpticalFlowParameters] from an objc.ID.
//
// An object that describes frame-level optical flow parameters.
func VTOpticalFlowParametersFromID(id objc.ID) VTOpticalFlowParameters {
	return VTOpticalFlowParameters{objectivec.Object{ID: id}}
}

// NOTE: VTOpticalFlowParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTOpticalFlowParameters] class.
//
// # Creating a parameters object
//
//   - [IVTOpticalFlowParameters.InitWithSourceFrameNextFrameSubmissionModeDestinationOpticalFlow]
//
// # Inspecting the parameters
//
//   - [IVTOpticalFlowParameters.NextFrame]: The next source frame in presentation time order.
//   - [IVTOpticalFlowParameters.DestinationOpticalFlow]: A user allocated mutable optical flow that will receive the results.
//   - [IVTOpticalFlowParameters.SubmissionMode]: A value describing the processing request in a parameters submission object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowParameters
type IVTOpticalFlowParameters interface {
	objectivec.IObject

	// Topic: Creating a parameters object

	InitWithSourceFrameNextFrameSubmissionModeDestinationOpticalFlow(sourceFrame IVTFrameProcessorFrame, nextFrame IVTFrameProcessorFrame, submissionMode VTOpticalFlowParametersSubmissionMode, destinationOpticalFlow IVTFrameProcessorOpticalFlow) VTOpticalFlowParameters

	// Topic: Inspecting the parameters

	// The next source frame in presentation time order.
	NextFrame() IVTFrameProcessorFrame
	// A user allocated mutable optical flow that will receive the results.
	DestinationOpticalFlow() IVTFrameProcessorOpticalFlow
	// A value describing the processing request in a parameters submission object.
	SubmissionMode() VTOpticalFlowParametersSubmissionMode
}

// Init initializes the instance.
func (v VTOpticalFlowParameters) Init() VTOpticalFlowParameters {
	rv := objc.Send[VTOpticalFlowParameters](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTOpticalFlowParameters) Autorelease() VTOpticalFlowParameters {
	rv := objc.Send[VTOpticalFlowParameters](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTOpticalFlowParameters creates a new VTOpticalFlowParameters instance.
func NewVTOpticalFlowParameters() VTOpticalFlowParameters {
	class := getVTOpticalFlowParametersClass()
	rv := objc.Send[VTOpticalFlowParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// sourceFrame: The current source frame. This must be a non-nil value.
//
// nextFrame: The next source frame in presentation time order. This value can be set to
// nil for the last frame.
//
// submissionMode: A value describing the processing request in a parameters submission
// object.
//
// destinationOpticalFlow: A user allocated mutable optical flow that will receive the results.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowParameters/init(sourceFrame:nextFrame:submissionMode:destinationOpticalFlow:)
func NewVTOpticalFlowParametersWithSourceFrameNextFrameSubmissionModeDestinationOpticalFlow(sourceFrame IVTFrameProcessorFrame, nextFrame IVTFrameProcessorFrame, submissionMode VTOpticalFlowParametersSubmissionMode, destinationOpticalFlow IVTFrameProcessorOpticalFlow) VTOpticalFlowParameters {
	instance := getVTOpticalFlowParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceFrame:nextFrame:submissionMode:destinationOpticalFlow:"), sourceFrame, nextFrame, submissionMode, destinationOpticalFlow)
	return VTOpticalFlowParametersFromID(rv)
}

// sourceFrame: The current source frame. This must be a non-nil value.
//
// nextFrame: The next source frame in presentation time order. This value can be set to
// nil for the last frame.
//
// submissionMode: A value describing the processing request in a parameters submission
// object.
//
// destinationOpticalFlow: A user allocated mutable optical flow that will receive the results.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowParameters/init(sourceFrame:nextFrame:submissionMode:destinationOpticalFlow:)
func (v VTOpticalFlowParameters) InitWithSourceFrameNextFrameSubmissionModeDestinationOpticalFlow(sourceFrame IVTFrameProcessorFrame, nextFrame IVTFrameProcessorFrame, submissionMode VTOpticalFlowParametersSubmissionMode, destinationOpticalFlow IVTFrameProcessorOpticalFlow) VTOpticalFlowParameters {
	rv := objc.Send[VTOpticalFlowParameters](v.ID, objc.Sel("initWithSourceFrame:nextFrame:submissionMode:destinationOpticalFlow:"), sourceFrame, nextFrame, submissionMode, destinationOpticalFlow)
	return rv
}

// The current source frame.
//
// # Discussion
//
// This must be a non-nil value.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowParameters/sourceFrame
func (v VTOpticalFlowParameters) SourceFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourceFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// The next source frame in presentation time order.
//
// # Discussion
//
// This value can be set to nil for the last frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowParameters/nextFrame
func (v VTOpticalFlowParameters) NextFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("nextFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// A user allocated mutable optical flow that will receive the results.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowParameters/destinationOpticalFlow
func (v VTOpticalFlowParameters) DestinationOpticalFlow() IVTFrameProcessorOpticalFlow {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationOpticalFlow"))
	return VTFrameProcessorOpticalFlowFromID(objc.ID(rv))
}

// A value describing the processing request in a parameters submission
// object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTOpticalFlowParameters/submissionMode-swift.property
func (v VTOpticalFlowParameters) SubmissionMode() VTOpticalFlowParametersSubmissionMode {
	rv := objc.Send[VTOpticalFlowParametersSubmissionMode](v.ID, objc.Sel("submissionMode"))
	return VTOpticalFlowParametersSubmissionMode(rv)
}

// Protocol methods for VTFrameProcessorParameters

// Destination frame that contains the destination frame for processors which
// output a single processed frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrame-3im3l
func (o VTOpticalFlowParameters) DestinationFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("destinationFrame"))
	return VTFrameProcessorFrameFromID(rv)
}

// Array of destination frames for processors which may output more than one
// processed frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrames-8yges
func (o VTOpticalFlowParameters) DestinationFrames() []VTFrameProcessorFrame {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("destinationFrames"))
	result := make([]VTFrameProcessorFrame, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = VTFrameProcessorFrameFromID(id)
	}
	return result
}

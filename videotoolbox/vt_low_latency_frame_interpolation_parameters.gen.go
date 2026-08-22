// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTLowLatencyFrameInterpolationParameters] class.
var (
	_VTLowLatencyFrameInterpolationParametersClass     VTLowLatencyFrameInterpolationParametersClass
	_VTLowLatencyFrameInterpolationParametersClassOnce sync.Once
)

func getVTLowLatencyFrameInterpolationParametersClass() VTLowLatencyFrameInterpolationParametersClass {
	_VTLowLatencyFrameInterpolationParametersClassOnce.Do(func() {
		_VTLowLatencyFrameInterpolationParametersClass = VTLowLatencyFrameInterpolationParametersClass{class: objc.GetClass("VTLowLatencyFrameInterpolationParameters")}
	})
	return _VTLowLatencyFrameInterpolationParametersClass
}

// GetVTLowLatencyFrameInterpolationParametersClass returns the class object for VTLowLatencyFrameInterpolationParameters.
func GetVTLowLatencyFrameInterpolationParametersClass() VTLowLatencyFrameInterpolationParametersClass {
	return getVTLowLatencyFrameInterpolationParametersClass()
}

type VTLowLatencyFrameInterpolationParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTLowLatencyFrameInterpolationParametersClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTLowLatencyFrameInterpolationParametersClass) Alloc() VTLowLatencyFrameInterpolationParameters {
	rv := objc.Send[VTLowLatencyFrameInterpolationParameters](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains both input and output parameters that the
// low-latency frame interpolation processor needs.
//
// # Overview
//
// Use this object in the `processWithParameters` call of [VTFrameProcessor]
// class.
//
// [VTLowLatencyFrameInterpolationParameters] are frame-level parameters.
//
// # Inspecting the parameters
//
//   - [VTLowLatencyFrameInterpolationParameters.PreviousFrame]: Previous frame that you provided when creating the low-latency frame interpolation parameters object.
//   - [VTLowLatencyFrameInterpolationParameters.InterpolationPhase]: Array of interpolation phases that you provided when creating the low-latency frame interpolation parameters object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationParameters
type VTLowLatencyFrameInterpolationParameters struct {
	objectivec.Object
}

// VTLowLatencyFrameInterpolationParametersFromID constructs a [VTLowLatencyFrameInterpolationParameters] from an objc.ID.
//
// An object that contains both input and output parameters that the
// low-latency frame interpolation processor needs.
func VTLowLatencyFrameInterpolationParametersFromID(id objc.ID) VTLowLatencyFrameInterpolationParameters {
	return VTLowLatencyFrameInterpolationParameters{objectivec.Object{ID: id}}
}

// NOTE: VTLowLatencyFrameInterpolationParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTLowLatencyFrameInterpolationParameters] class.
//
// # Inspecting the parameters
//
//   - [IVTLowLatencyFrameInterpolationParameters.PreviousFrame]: Previous frame that you provided when creating the low-latency frame interpolation parameters object.
//   - [IVTLowLatencyFrameInterpolationParameters.InterpolationPhase]: Array of interpolation phases that you provided when creating the low-latency frame interpolation parameters object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationParameters
type IVTLowLatencyFrameInterpolationParameters interface {
	objectivec.IObject

	// Topic: Inspecting the parameters

	// Previous frame that you provided when creating the low-latency frame interpolation parameters object.
	PreviousFrame() IVTFrameProcessorFrame
	// Array of interpolation phases that you provided when creating the low-latency frame interpolation parameters object.
	InterpolationPhase() []foundation.NSNumber

	// Creates a new low-latency frame interpolation parameters object.
	InitWithSourceFramePreviousFrameInterpolationPhaseDestinationFrames(sourceFrame IVTFrameProcessorFrame, previousFrame IVTFrameProcessorFrame, interpolationPhase []foundation.NSNumber, destinationFrames []VTFrameProcessorFrame) VTLowLatencyFrameInterpolationParameters
}

// Init initializes the instance.
func (v VTLowLatencyFrameInterpolationParameters) Init() VTLowLatencyFrameInterpolationParameters {
	rv := objc.Send[VTLowLatencyFrameInterpolationParameters](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTLowLatencyFrameInterpolationParameters) Autorelease() VTLowLatencyFrameInterpolationParameters {
	rv := objc.Send[VTLowLatencyFrameInterpolationParameters](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTLowLatencyFrameInterpolationParameters creates a new VTLowLatencyFrameInterpolationParameters instance.
func NewVTLowLatencyFrameInterpolationParameters() VTLowLatencyFrameInterpolationParameters {
	class := getVTLowLatencyFrameInterpolationParametersClass()
	rv := objc.Send[VTLowLatencyFrameInterpolationParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new low-latency frame interpolation parameters object.
//
// sourceFrame: Current frame to use for interpolation; must be non `nil`.
//
// previousFrame: Previous frame used for interpolation; must be non `nil`.
//
// interpolationPhase: Array of float numbers that indicate interpolation phase locations at which
// the processor interpolates the frames. Must be greater than 0 and less than
// 1.0; for example 0.5 is midway between the previous frame and the source
// frame. If you enable spatial scaling, the only supported interpolation
// phase is 0.5.
//
// destinationFrames: Caller-allocated array of [VTFrameProcessorFrame] to receive the
// interpolated frames. This must have the same number of elements as the the
// `interpolationPhase`. If you enable spatial scaling, it must also contain
// an element to hold the scaled version of sourceFrame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationParameters/initWithSourceFrame:previousFrame:interpolationPhase:destinationFrames:
func NewVTLowLatencyFrameInterpolationParametersWithSourceFramePreviousFrameInterpolationPhaseDestinationFrames(sourceFrame IVTFrameProcessorFrame, previousFrame IVTFrameProcessorFrame, interpolationPhase []foundation.NSNumber, destinationFrames []VTFrameProcessorFrame) VTLowLatencyFrameInterpolationParameters {
	instance := getVTLowLatencyFrameInterpolationParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceFrame:previousFrame:interpolationPhase:destinationFrames:"), sourceFrame, previousFrame, objectivec.IObjectSliceToNSArray(interpolationPhase), objectivec.IObjectSliceToNSArray(destinationFrames))
	return VTLowLatencyFrameInterpolationParametersFromID(rv)
}

// Creates a new low-latency frame interpolation parameters object.
//
// sourceFrame: Current frame to use for interpolation; must be non `nil`.
//
// previousFrame: Previous frame used for interpolation; must be non `nil`.
//
// interpolationPhase: Array of float numbers that indicate interpolation phase locations at which
// the processor interpolates the frames. Must be greater than 0 and less than
// 1.0; for example 0.5 is midway between the previous frame and the source
// frame. If you enable spatial scaling, the only supported interpolation
// phase is 0.5.
//
// destinationFrames: Caller-allocated array of [VTFrameProcessorFrame] to receive the
// interpolated frames. This must have the same number of elements as the the
// `interpolationPhase`. If you enable spatial scaling, it must also contain
// an element to hold the scaled version of sourceFrame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationParameters/initWithSourceFrame:previousFrame:interpolationPhase:destinationFrames:
func (v VTLowLatencyFrameInterpolationParameters) InitWithSourceFramePreviousFrameInterpolationPhaseDestinationFrames(sourceFrame IVTFrameProcessorFrame, previousFrame IVTFrameProcessorFrame, interpolationPhase []foundation.NSNumber, destinationFrames []VTFrameProcessorFrame) VTLowLatencyFrameInterpolationParameters {
	rv := objc.Send[VTLowLatencyFrameInterpolationParameters](v.ID, objc.Sel("initWithSourceFrame:previousFrame:interpolationPhase:destinationFrames:"), sourceFrame, previousFrame, objectivec.IObjectSliceToNSArray(interpolationPhase), objectivec.IObjectSliceToNSArray(destinationFrames))
	return rv
}

// Source frame that you provided when creating the low-latency frame
// interpolation parameters object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationParameters/sourceFrame
func (v VTLowLatencyFrameInterpolationParameters) SourceFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourceFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Previous frame that you provided when creating the low-latency frame
// interpolation parameters object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationParameters/previousFrame
func (v VTLowLatencyFrameInterpolationParameters) PreviousFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("previousFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Array of interpolation phases that you provided when creating the
// low-latency frame interpolation parameters object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationParameters/interpolationPhase-33m49
func (v VTLowLatencyFrameInterpolationParameters) InterpolationPhase() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("interpolationPhase"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Array of destination frames that you provided when creating the low-latency
// frame interpolation parameters object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencyFrameInterpolationParameters/destinationFrames
func (v VTLowLatencyFrameInterpolationParameters) DestinationFrames() []VTFrameProcessorFrame {
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
func (o VTLowLatencyFrameInterpolationParameters) DestinationFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("destinationFrame"))
	return VTFrameProcessorFrameFromID(rv)
}

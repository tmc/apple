// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTTemporalNoiseFilterParameters] class.
var (
	_VTTemporalNoiseFilterParametersClass     VTTemporalNoiseFilterParametersClass
	_VTTemporalNoiseFilterParametersClassOnce sync.Once
)

func getVTTemporalNoiseFilterParametersClass() VTTemporalNoiseFilterParametersClass {
	_VTTemporalNoiseFilterParametersClassOnce.Do(func() {
		_VTTemporalNoiseFilterParametersClass = VTTemporalNoiseFilterParametersClass{class: objc.GetClass("VTTemporalNoiseFilterParameters")}
	})
	return _VTTemporalNoiseFilterParametersClass
}

// GetVTTemporalNoiseFilterParametersClass returns the class object for VTTemporalNoiseFilterParameters.
func GetVTTemporalNoiseFilterParametersClass() VTTemporalNoiseFilterParametersClass {
	return getVTTemporalNoiseFilterParametersClass()
}

type VTTemporalNoiseFilterParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTTemporalNoiseFilterParametersClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTTemporalNoiseFilterParametersClass) Alloc() VTTemporalNoiseFilterParameters {
	rv := objc.Send[VTTemporalNoiseFilterParameters](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Encapsulates the frame-level parameters necessary for processing a source
// frame using temporal noise-filter processor.
//
// # Overview
//
// This object is intended for sending input parameters into the
// `processWithParameters` method of the [VTFrameProcessor] class. Temporal
// noise-filter processor utilizes past and future reference frames, provided
// in presentation time order, to reduce noise from the source frame. The
// `previousFrameCount` and `nextFrameCount` properties in
// [VTTemporalNoiseFilterConfiguration] represent the maximum number of past
// and future reference frames that the processor can use to achieve optimum
// noise reduction quality. The number of reference frames provided shall
// depend on their availability, but at a minimum, you must provide one
// reference frame, either past or future. The parameter `destinationFrame`
// stores the output frame that the processor returns to the caller upon the
// successful completion of the `processWithParameters` operation.
//
// # Creating a parameters object
//
//   - [VTTemporalNoiseFilterParameters.InitWithSourceFrameNextFramesPreviousFramesDestinationFrameFilterStrengthHasDiscontinuity]: Creates a new [VTTemporalNoiseFilterParameters] object.
//
// # Inspecting the parameters
//
//   - [VTTemporalNoiseFilterParameters.NextFrames]: Future reference frames in presentation time order that you use to process the source frame.
//   - [VTTemporalNoiseFilterParameters.PreviousFrames]: Past reference frames in presentation time order that you use to process the source frame.
//   - [VTTemporalNoiseFilterParameters.FilterStrength]: A parameter to control the strength of noise-filtering. The value can range from the minimum strength of 0.0 to the maximum strength of 1.0. Change in filter strength causes the processor to flush all frames in the queue prior to processing the source frame.
//   - [VTTemporalNoiseFilterParameters.SetFilterStrength]
//   - [VTTemporalNoiseFilterParameters.HasDiscontinuity]: A Boolean that indicates sequence discontinuity, forcing the processor to reset prior to processing the source frame.
//   - [VTTemporalNoiseFilterParameters.SetHasDiscontinuity]
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterParameters
type VTTemporalNoiseFilterParameters struct {
	objectivec.Object
}

// VTTemporalNoiseFilterParametersFromID constructs a [VTTemporalNoiseFilterParameters] from an objc.ID.
//
// Encapsulates the frame-level parameters necessary for processing a source
// frame using temporal noise-filter processor.
func VTTemporalNoiseFilterParametersFromID(id objc.ID) VTTemporalNoiseFilterParameters {
	return VTTemporalNoiseFilterParameters{objectivec.Object{ID: id}}
}

// NOTE: VTTemporalNoiseFilterParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTTemporalNoiseFilterParameters] class.
//
// # Creating a parameters object
//
//   - [IVTTemporalNoiseFilterParameters.InitWithSourceFrameNextFramesPreviousFramesDestinationFrameFilterStrengthHasDiscontinuity]: Creates a new [VTTemporalNoiseFilterParameters] object.
//
// # Inspecting the parameters
//
//   - [IVTTemporalNoiseFilterParameters.NextFrames]: Future reference frames in presentation time order that you use to process the source frame.
//   - [IVTTemporalNoiseFilterParameters.PreviousFrames]: Past reference frames in presentation time order that you use to process the source frame.
//   - [IVTTemporalNoiseFilterParameters.FilterStrength]: A parameter to control the strength of noise-filtering. The value can range from the minimum strength of 0.0 to the maximum strength of 1.0. Change in filter strength causes the processor to flush all frames in the queue prior to processing the source frame.
//   - [IVTTemporalNoiseFilterParameters.SetFilterStrength]
//   - [IVTTemporalNoiseFilterParameters.HasDiscontinuity]: A Boolean that indicates sequence discontinuity, forcing the processor to reset prior to processing the source frame.
//   - [IVTTemporalNoiseFilterParameters.SetHasDiscontinuity]
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterParameters
type IVTTemporalNoiseFilterParameters interface {
	objectivec.IObject

	// Topic: Creating a parameters object

	// Creates a new [VTTemporalNoiseFilterParameters] object.
	InitWithSourceFrameNextFramesPreviousFramesDestinationFrameFilterStrengthHasDiscontinuity(sourceFrame IVTFrameProcessorFrame, nextFrames []VTFrameProcessorFrame, previousFrames []VTFrameProcessorFrame, destinationFrame IVTFrameProcessorFrame, filterStrength float32, hasDiscontinuity bool) VTTemporalNoiseFilterParameters

	// Topic: Inspecting the parameters

	// Future reference frames in presentation time order that you use to process the source frame.
	NextFrames() []VTFrameProcessorFrame
	// Past reference frames in presentation time order that you use to process the source frame.
	PreviousFrames() []VTFrameProcessorFrame
	// A parameter to control the strength of noise-filtering. The value can range from the minimum strength of 0.0 to the maximum strength of 1.0. Change in filter strength causes the processor to flush all frames in the queue prior to processing the source frame.
	FilterStrength() float32
	SetFilterStrength(value float32)
	// A Boolean that indicates sequence discontinuity, forcing the processor to reset prior to processing the source frame.
	HasDiscontinuity() bool
	SetHasDiscontinuity(value bool)
}

// Init initializes the instance.
func (v VTTemporalNoiseFilterParameters) Init() VTTemporalNoiseFilterParameters {
	rv := objc.Send[VTTemporalNoiseFilterParameters](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTTemporalNoiseFilterParameters) Autorelease() VTTemporalNoiseFilterParameters {
	rv := objc.Send[VTTemporalNoiseFilterParameters](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTTemporalNoiseFilterParameters creates a new VTTemporalNoiseFilterParameters instance.
func NewVTTemporalNoiseFilterParameters() VTTemporalNoiseFilterParameters {
	class := getVTTemporalNoiseFilterParametersClass()
	rv := objc.Send[VTTemporalNoiseFilterParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new [VTTemporalNoiseFilterParameters] object.
//
// sourceFrame: Current source frame; must be non `nil`.
//
// nextFrames: Future reference frames in presentation time order to use for processing
// the source frame. The number of frames can vary from 0 to the number
// specified by [VTTemporalNoiseFilterConfiguration.NextFrameCount] property.
//
// previousFrames: Past reference frames in presentation time order to use for processing the
// source frame. The number of frames can vary from 0 to the number specified
// by [VTTemporalNoiseFilterConfiguration.PreviousFrameCount] property.
//
// destinationFrame: User-allocated pixel buffer that receives the output frame. The pixel
// format of `destinationFrame` must match with that of the `sourceFrame`.
//
// filterStrength: Strength of the noise-filtering to use. The value can range from the
// minimum strength of 0.0 to the maximum strength of 1.0. Change in filter
// strength causes the processor to flush all frames in the queue prior to
// processing the source frame.
//
// hasDiscontinuity: Marks sequence discontinuity, forcing the processor to reset prior to
// processing the source frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterParameters/init(sourceFrame:nextFrames:previousFrames:destinationFrame:filterStrength:hasDiscontinuity:)
func NewVTTemporalNoiseFilterParametersWithSourceFrameNextFramesPreviousFramesDestinationFrameFilterStrengthHasDiscontinuity(sourceFrame IVTFrameProcessorFrame, nextFrames []VTFrameProcessorFrame, previousFrames []VTFrameProcessorFrame, destinationFrame IVTFrameProcessorFrame, filterStrength float32, hasDiscontinuity bool) VTTemporalNoiseFilterParameters {
	instance := getVTTemporalNoiseFilterParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceFrame:nextFrames:previousFrames:destinationFrame:filterStrength:hasDiscontinuity:"), sourceFrame, objectivec.IObjectSliceToNSArray(nextFrames), objectivec.IObjectSliceToNSArray(previousFrames), destinationFrame, filterStrength, hasDiscontinuity)
	return VTTemporalNoiseFilterParametersFromID(rv)
}

// Creates a new [VTTemporalNoiseFilterParameters] object.
//
// sourceFrame: Current source frame; must be non `nil`.
//
// nextFrames: Future reference frames in presentation time order to use for processing
// the source frame. The number of frames can vary from 0 to the number
// specified by [VTTemporalNoiseFilterConfiguration.NextFrameCount] property.
//
// previousFrames: Past reference frames in presentation time order to use for processing the
// source frame. The number of frames can vary from 0 to the number specified
// by [VTTemporalNoiseFilterConfiguration.PreviousFrameCount] property.
//
// destinationFrame: User-allocated pixel buffer that receives the output frame. The pixel
// format of `destinationFrame` must match with that of the `sourceFrame`.
//
// filterStrength: Strength of the noise-filtering to use. The value can range from the
// minimum strength of 0.0 to the maximum strength of 1.0. Change in filter
// strength causes the processor to flush all frames in the queue prior to
// processing the source frame.
//
// hasDiscontinuity: Marks sequence discontinuity, forcing the processor to reset prior to
// processing the source frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterParameters/init(sourceFrame:nextFrames:previousFrames:destinationFrame:filterStrength:hasDiscontinuity:)
func (v VTTemporalNoiseFilterParameters) InitWithSourceFrameNextFramesPreviousFramesDestinationFrameFilterStrengthHasDiscontinuity(sourceFrame IVTFrameProcessorFrame, nextFrames []VTFrameProcessorFrame, previousFrames []VTFrameProcessorFrame, destinationFrame IVTFrameProcessorFrame, filterStrength float32, hasDiscontinuity bool) VTTemporalNoiseFilterParameters {
	rv := objc.Send[VTTemporalNoiseFilterParameters](v.ID, objc.Sel("initWithSourceFrame:nextFrames:previousFrames:destinationFrame:filterStrength:hasDiscontinuity:"), sourceFrame, objectivec.IObjectSliceToNSArray(nextFrames), objectivec.IObjectSliceToNSArray(previousFrames), destinationFrame, filterStrength, hasDiscontinuity)
	return rv
}

// Current source frame; must be non `nil`.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterParameters/sourceFrame
func (v VTTemporalNoiseFilterParameters) SourceFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourceFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Future reference frames in presentation time order that you use to process
// the source frame.
//
// # Discussion
//
// The number of frames can vary from 0 to the number specified by the
// `nextFrameCount` property in [VTTemporalNoiseFilterConfiguration].
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterParameters/nextFrames
func (v VTTemporalNoiseFilterParameters) NextFrames() []VTFrameProcessorFrame {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("nextFrames"))
	return objc.ConvertSlice(rv, func(id objc.ID) VTFrameProcessorFrame {
		return VTFrameProcessorFrameFromID(id)
	})
}

// Past reference frames in presentation time order that you use to process
// the source frame.
//
// # Discussion
//
// The number of frames can vary from 0 to the number specified by the
// `previousFrameCount` property in [VTTemporalNoiseFilterConfiguration].
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterParameters/previousFrames
func (v VTTemporalNoiseFilterParameters) PreviousFrames() []VTFrameProcessorFrame {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("previousFrames"))
	return objc.ConvertSlice(rv, func(id objc.ID) VTFrameProcessorFrame {
		return VTFrameProcessorFrameFromID(id)
	})
}

// A parameter to control the strength of noise-filtering. The value can range
// from the minimum strength of 0.0 to the maximum strength of 1.0. Change in
// filter strength causes the processor to flush all frames in the queue prior
// to processing the source frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterParameters/filterStrength
func (v VTTemporalNoiseFilterParameters) FilterStrength() float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("filterStrength"))
	return rv
}
func (v VTTemporalNoiseFilterParameters) SetFilterStrength(value float32) {
	objc.Send[struct{}](v.ID, objc.Sel("setFilterStrength:"), value)
}

// A Boolean that indicates sequence discontinuity, forcing the processor to
// reset prior to processing the source frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterParameters/hasDiscontinuity
func (v VTTemporalNoiseFilterParameters) HasDiscontinuity() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("hasDiscontinuity"))
	return rv
}
func (v VTTemporalNoiseFilterParameters) SetHasDiscontinuity(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setHasDiscontinuity:"), value)
}

// Destination frame that contains a user-allocated pixel buffer that receives
// the output frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTTemporalNoiseFilterParameters/destinationFrame
func (v VTTemporalNoiseFilterParameters) DestinationFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Protocol methods for VTFrameProcessorParameters

// Array of destination frames for processors which may output more than one
// processed frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrames-8yges
func (o VTTemporalNoiseFilterParameters) DestinationFrames() []VTFrameProcessorFrame {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("destinationFrames"))
	result := make([]VTFrameProcessorFrame, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = VTFrameProcessorFrameFromID(id)
	}
	return result
}

// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTLowLatencySuperResolutionScalerParameters] class.
var (
	_VTLowLatencySuperResolutionScalerParametersClass     VTLowLatencySuperResolutionScalerParametersClass
	_VTLowLatencySuperResolutionScalerParametersClassOnce sync.Once
)

func getVTLowLatencySuperResolutionScalerParametersClass() VTLowLatencySuperResolutionScalerParametersClass {
	_VTLowLatencySuperResolutionScalerParametersClassOnce.Do(func() {
		_VTLowLatencySuperResolutionScalerParametersClass = VTLowLatencySuperResolutionScalerParametersClass{class: objc.GetClass("VTLowLatencySuperResolutionScalerParameters")}
	})
	return _VTLowLatencySuperResolutionScalerParametersClass
}

// GetVTLowLatencySuperResolutionScalerParametersClass returns the class object for VTLowLatencySuperResolutionScalerParameters.
func GetVTLowLatencySuperResolutionScalerParametersClass() VTLowLatencySuperResolutionScalerParametersClass {
	return getVTLowLatencySuperResolutionScalerParametersClass()
}

type VTLowLatencySuperResolutionScalerParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTLowLatencySuperResolutionScalerParametersClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTLowLatencySuperResolutionScalerParametersClass) Alloc() VTLowLatencySuperResolutionScalerParameters {
	rv := objc.Send[VTLowLatencySuperResolutionScalerParameters](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains both input and output parameters that the
// low-latency super-resolution scaler frame processor needs.
//
// # Overview
//
// Use this object in the `processWithParameters` call of [VTFrameProcessor]
// class.
//
// [VTLowLatencySuperResolutionScalerParameters] are frame-level parameters.
//
// # Creating a parameters object
//
//   - [VTLowLatencySuperResolutionScalerParameters.InitWithSourceFrameDestinationFrame]: Creates a new low-latency, super-resolution scaler parameters object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerParameters
type VTLowLatencySuperResolutionScalerParameters struct {
	objectivec.Object
}

// VTLowLatencySuperResolutionScalerParametersFromID constructs a [VTLowLatencySuperResolutionScalerParameters] from an objc.ID.
//
// An object that contains both input and output parameters that the
// low-latency super-resolution scaler frame processor needs.
func VTLowLatencySuperResolutionScalerParametersFromID(id objc.ID) VTLowLatencySuperResolutionScalerParameters {
	return VTLowLatencySuperResolutionScalerParameters{objectivec.Object{ID: id}}
}

// NOTE: VTLowLatencySuperResolutionScalerParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTLowLatencySuperResolutionScalerParameters] class.
//
// # Creating a parameters object
//
//   - [IVTLowLatencySuperResolutionScalerParameters.InitWithSourceFrameDestinationFrame]: Creates a new low-latency, super-resolution scaler parameters object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerParameters
type IVTLowLatencySuperResolutionScalerParameters interface {
	objectivec.IObject

	// Topic: Creating a parameters object

	// Creates a new low-latency, super-resolution scaler parameters object.
	InitWithSourceFrameDestinationFrame(sourceFrame IVTFrameProcessorFrame, destinationFrame IVTFrameProcessorFrame) VTLowLatencySuperResolutionScalerParameters
}

// Init initializes the instance.
func (v VTLowLatencySuperResolutionScalerParameters) Init() VTLowLatencySuperResolutionScalerParameters {
	rv := objc.Send[VTLowLatencySuperResolutionScalerParameters](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTLowLatencySuperResolutionScalerParameters) Autorelease() VTLowLatencySuperResolutionScalerParameters {
	rv := objc.Send[VTLowLatencySuperResolutionScalerParameters](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTLowLatencySuperResolutionScalerParameters creates a new VTLowLatencySuperResolutionScalerParameters instance.
func NewVTLowLatencySuperResolutionScalerParameters() VTLowLatencySuperResolutionScalerParameters {
	class := getVTLowLatencySuperResolutionScalerParametersClass()
	rv := objc.Send[VTLowLatencySuperResolutionScalerParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new low-latency, super-resolution scaler parameters object.
//
// sourceFrame: Current source frame; must be non `nil`.
//
// destinationFrame: User-allocated pixel buffer that receives the scaled processor output; must
// be non `nil`.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerParameters/init(sourceFrame:destinationFrame:)
func NewVTLowLatencySuperResolutionScalerParametersWithSourceFrameDestinationFrame(sourceFrame IVTFrameProcessorFrame, destinationFrame IVTFrameProcessorFrame) VTLowLatencySuperResolutionScalerParameters {
	instance := getVTLowLatencySuperResolutionScalerParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceFrame:destinationFrame:"), sourceFrame, destinationFrame)
	return VTLowLatencySuperResolutionScalerParametersFromID(rv)
}

// Creates a new low-latency, super-resolution scaler parameters object.
//
// sourceFrame: Current source frame; must be non `nil`.
//
// destinationFrame: User-allocated pixel buffer that receives the scaled processor output; must
// be non `nil`.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerParameters/init(sourceFrame:destinationFrame:)
func (v VTLowLatencySuperResolutionScalerParameters) InitWithSourceFrameDestinationFrame(sourceFrame IVTFrameProcessorFrame, destinationFrame IVTFrameProcessorFrame) VTLowLatencySuperResolutionScalerParameters {
	rv := objc.Send[VTLowLatencySuperResolutionScalerParameters](v.ID, objc.Sel("initWithSourceFrame:destinationFrame:"), sourceFrame, destinationFrame)
	return rv
}

// Current source frame, which must be non `nil`.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerParameters/sourceFrame
func (v VTLowLatencySuperResolutionScalerParameters) SourceFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourceFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Destination frame that contains user-allocated pixel buffer that receives
// the scaled processor output.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTLowLatencySuperResolutionScalerParameters/destinationFrame
func (v VTLowLatencySuperResolutionScalerParameters) DestinationFrame() IVTFrameProcessorFrame {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("destinationFrame"))
	return VTFrameProcessorFrameFromID(objc.ID(rv))
}

// Protocol methods for VTFrameProcessorParameters

// Array of destination frames for processors which may output more than one
// processed frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorParameters/destinationFrames-8yges
func (o VTLowLatencySuperResolutionScalerParameters) DestinationFrames() []VTFrameProcessorFrame {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("destinationFrames"))
	result := make([]VTFrameProcessorFrame, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = VTFrameProcessorFrameFromID(id)
	}
	return result
}

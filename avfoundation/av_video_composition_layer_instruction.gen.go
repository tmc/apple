// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVideoCompositionLayerInstruction] class.
var (
	_AVVideoCompositionLayerInstructionClass     AVVideoCompositionLayerInstructionClass
	_AVVideoCompositionLayerInstructionClassOnce sync.Once
)

func getAVVideoCompositionLayerInstructionClass() AVVideoCompositionLayerInstructionClass {
	_AVVideoCompositionLayerInstructionClassOnce.Do(func() {
		_AVVideoCompositionLayerInstructionClass = AVVideoCompositionLayerInstructionClass{class: objc.GetClass("AVVideoCompositionLayerInstruction")}
	})
	return _AVVideoCompositionLayerInstructionClass
}

// GetAVVideoCompositionLayerInstructionClass returns the class object for AVVideoCompositionLayerInstruction.
func GetAVVideoCompositionLayerInstructionClass() AVVideoCompositionLayerInstructionClass {
	return getAVVideoCompositionLayerInstructionClass()
}

type AVVideoCompositionLayerInstructionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVideoCompositionLayerInstructionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVideoCompositionLayerInstructionClass) Alloc() AVVideoCompositionLayerInstruction {
	rv := objc.Send[AVVideoCompositionLayerInstruction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object used to modify the transform, cropping, and opacity ramps applied
// to a given track in a composition.
//
// # Getting the track ID
//
//   - [AVVideoCompositionLayerInstruction.TrackID]: The track identifier of the source track to which the compositor will apply the instruction.
//
// # Getting opacity, transform, and cropping ramps
//
//   - [AVVideoCompositionLayerInstruction.GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange]: Obtains the crop rectangle ramp that includes the specified time.
//   - [AVVideoCompositionLayerInstruction.GetOpacityRampForTimeStartOpacityEndOpacityTimeRange]: Obtains the opacity ramp that includes a specified time.
//   - [AVVideoCompositionLayerInstruction.GetTransformRampForTimeStartTransformEndTransformTimeRange]: Obtains the transform ramp that includes a specified time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction
type AVVideoCompositionLayerInstruction struct {
	objectivec.Object
}

// AVVideoCompositionLayerInstructionFromID constructs a [AVVideoCompositionLayerInstruction] from an objc.ID.
//
// An object used to modify the transform, cropping, and opacity ramps applied
// to a given track in a composition.
func AVVideoCompositionLayerInstructionFromID(id objc.ID) AVVideoCompositionLayerInstruction {
	return AVVideoCompositionLayerInstruction{objectivec.Object{ID: id}}
}

// NOTE: AVVideoCompositionLayerInstruction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVVideoCompositionLayerInstruction] class.
//
// # Getting the track ID
//
//   - [IAVVideoCompositionLayerInstruction.TrackID]: The track identifier of the source track to which the compositor will apply the instruction.
//
// # Getting opacity, transform, and cropping ramps
//
//   - [IAVVideoCompositionLayerInstruction.GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange]: Obtains the crop rectangle ramp that includes the specified time.
//   - [IAVVideoCompositionLayerInstruction.GetOpacityRampForTimeStartOpacityEndOpacityTimeRange]: Obtains the opacity ramp that includes a specified time.
//   - [IAVVideoCompositionLayerInstruction.GetTransformRampForTimeStartTransformEndTransformTimeRange]: Obtains the transform ramp that includes a specified time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction
type IAVVideoCompositionLayerInstruction interface {
	objectivec.IObject

	// Topic: Getting the track ID

	// The track identifier of the source track to which the compositor will apply the instruction.
	TrackID() int32

	// Topic: Getting opacity, transform, and cropping ramps

	// Obtains the crop rectangle ramp that includes the specified time.
	GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange(time coremedia.CMTime, startCropRectangle *corefoundation.CGRect, endCropRectangle *corefoundation.CGRect, timeRange *coremedia.CMTimeRange) bool
	// Obtains the opacity ramp that includes a specified time.
	GetOpacityRampForTimeStartOpacityEndOpacityTimeRange(time coremedia.CMTime, timeRange *coremedia.CMTimeRange) (float32, float32, bool)
	// Obtains the transform ramp that includes a specified time.
	GetTransformRampForTimeStartTransformEndTransformTimeRange(time coremedia.CMTime, startTransform *corefoundation.CGAffineTransform, endTransform *corefoundation.CGAffineTransform, timeRange *coremedia.CMTimeRange) bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (v AVVideoCompositionLayerInstruction) Init() AVVideoCompositionLayerInstruction {
	rv := objc.Send[AVVideoCompositionLayerInstruction](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVideoCompositionLayerInstruction) Autorelease() AVVideoCompositionLayerInstruction {
	rv := objc.Send[AVVideoCompositionLayerInstruction](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVideoCompositionLayerInstruction creates a new AVVideoCompositionLayerInstruction instance.
func NewAVVideoCompositionLayerInstruction() AVVideoCompositionLayerInstruction {
	class := getAVVideoCompositionLayerInstructionClass()
	rv := objc.Send[AVVideoCompositionLayerInstruction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Obtains the crop rectangle ramp that includes the specified time.
//
// time: If a ramp with a time range that contains the specified time has been set,
// information about the effective ramp for that time is supplied. Otherwise,
// information about the first ramp that starts after the specified time is
// supplied.
//
// startCropRectangle: A pointer to a [CGRect] to receive the starting crop rectangle value for
// the crop rectangle ramp.
//
// May be NULL.
//
// endCropRectangle: A pointer to a [CGRect] to receive the ending crop rectangle value for the
// crop rectangle ramp.
//
// This value may be [NULL].
//
// timeRange: A pointer to a [CMTimeRange] to receive the time range of the crop
// rectangle ramp.
//
// This value may be [NULL].
//
// # Return Value
//
// false will be returned if the specified time is beyond the duration of the
// last crop rectangle ramp that has been set.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/getCropRectangleRamp(for:startCropRectangle:endCropRectangle:timeRange:)
func (v AVVideoCompositionLayerInstruction) GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange(time coremedia.CMTime, startCropRectangle *corefoundation.CGRect, endCropRectangle *corefoundation.CGRect, timeRange *coremedia.CMTimeRange) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("getCropRectangleRampForTime:startCropRectangle:endCropRectangle:timeRange:"), time, startCropRectangle, endCropRectangle, timeRange)
	return rv
}

// Obtains the opacity ramp that includes a specified time.
//
// time: If a ramp with a time range that contains the specified time has been set,
// information about the effective ramp for that time is supplied. Otherwise,
// information about the first ramp that starts after the specified time is
// supplied.
//
// startOpacity: A pointer to a float to receive the starting opacity value for the opacity
// ramp.
//
// This value may be [NULL].
//
// endOpacity: A pointer to a float to receive the ending opacity value for the opacity
// ramp.
//
// This value may be [NULL].
//
// timeRange: A pointer to a [CMTimeRange] to receive the time range of the opacity ramp.
//
// This value may be [NULL].
//
// # Return Value
//
// true if values are returned successfully, otherwise false. false is
// returned if `time` is beyond the duration of the last opacity ramp that has
// been set.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/getOpacityRamp(for:startOpacity:endOpacity:timeRange:)
func (v AVVideoCompositionLayerInstruction) GetOpacityRampForTimeStartOpacityEndOpacityTimeRange(time coremedia.CMTime, timeRange *coremedia.CMTimeRange) (float32, float32, bool) {
	var startOpacity float32
	var endOpacity float32
	rv := objc.Send[bool](v.ID, objc.Sel("getOpacityRampForTime:startOpacity:endOpacity:timeRange:"), time, unsafe.Pointer(&startOpacity), unsafe.Pointer(&endOpacity), timeRange)
	return startOpacity, endOpacity, rv
}

// Obtains the transform ramp that includes a specified time.
//
// time: If a ramp with a time range that contains the specified time has been set,
// information about the effective ramp for that time is supplied. Otherwise,
// information about the first ramp that starts after the specified time is
// supplied.
//
// startTransform: A pointer to a float to receive the starting transform value for the
// transform ramp.
//
// This value may be [NULL].
//
// endTransform: A pointer to a float to receive the ending transform value for the
// transform ramp.
//
// This value may be [NULL].
//
// timeRange: A pointer to a [CMTimeRange] to receive the time range of the transform
// ramp.
//
// This value may be [NULL].
//
// # Return Value
//
// true if values are returned successfully, otherwise false. false is
// returned if `time` is beyond the duration of the last transform ramp that
// has been set.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/getTransformRamp(for:start:end:timeRange:)
func (v AVVideoCompositionLayerInstruction) GetTransformRampForTimeStartTransformEndTransformTimeRange(time coremedia.CMTime, startTransform *corefoundation.CGAffineTransform, endTransform *corefoundation.CGAffineTransform, timeRange *coremedia.CMTimeRange) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("getTransformRampForTime:startTransform:endTransform:timeRange:"), time, startTransform, endTransform, timeRange)
	return rv
}
func (v AVVideoCompositionLayerInstruction) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Pass-through initializer, for internal use in AVFoundation only
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/videoCompositionLayerInstructionWithLayerInstruction:
func (_AVVideoCompositionLayerInstructionClass AVVideoCompositionLayerInstructionClass) VideoCompositionLayerInstructionWithLayerInstruction(instruction IAVVideoCompositionLayerInstruction) AVVideoCompositionLayerInstruction {
	rv := objc.Send[objc.ID](objc.ID(_AVVideoCompositionLayerInstructionClass.class), objc.Sel("videoCompositionLayerInstructionWithLayerInstruction:"), instruction)
	return AVVideoCompositionLayerInstructionFromID(rv)
}

// The track identifier of the source track to which the compositor will apply
// the instruction.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/trackID
func (v AVVideoCompositionLayerInstruction) TrackID() int32 {
	rv := objc.Send[int32](v.ID, objc.Sel("trackID"))
	return rv
}

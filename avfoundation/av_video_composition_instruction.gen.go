// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVideoCompositionInstruction] class.
var (
	_AVVideoCompositionInstructionClass     AVVideoCompositionInstructionClass
	_AVVideoCompositionInstructionClassOnce sync.Once
)

func getAVVideoCompositionInstructionClass() AVVideoCompositionInstructionClass {
	_AVVideoCompositionInstructionClassOnce.Do(func() {
		_AVVideoCompositionInstructionClass = AVVideoCompositionInstructionClass{class: objc.GetClass("AVVideoCompositionInstruction")}
	})
	return _AVVideoCompositionInstructionClass
}

// GetAVVideoCompositionInstructionClass returns the class object for AVVideoCompositionInstruction.
func GetAVVideoCompositionInstructionClass() AVVideoCompositionInstructionClass {
	return getAVVideoCompositionInstructionClass()
}

type AVVideoCompositionInstructionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVideoCompositionInstructionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVideoCompositionInstructionClass) Alloc() AVVideoCompositionInstruction {
	rv := objc.Send[AVVideoCompositionInstruction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An operation that a compositor performs.
//
// # Overview
// 
// An [AVVideoComposition] object maintains an array of [AVVideoCompositionInstruction.Instructions] to
// perform its composition.
//
// # Inspecting the instruction
//
//   - [AVVideoCompositionInstruction.BackgroundColor]: The background color of the composition.
//   - [AVVideoCompositionInstruction.LayerInstructions]: Instructions that specify how to layer and compose video frames from source tracks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class
type AVVideoCompositionInstruction struct {
	objectivec.Object
}

// AVVideoCompositionInstructionFromID constructs a [AVVideoCompositionInstruction] from an objc.ID.
//
// An operation that a compositor performs.
func AVVideoCompositionInstructionFromID(id objc.ID) AVVideoCompositionInstruction {
	return AVVideoCompositionInstruction{objectivec.Object{ID: id}}
}
// NOTE: AVVideoCompositionInstruction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVVideoCompositionInstruction] class.
//
// # Inspecting the instruction
//
//   - [IAVVideoCompositionInstruction.BackgroundColor]: The background color of the composition.
//   - [IAVVideoCompositionInstruction.LayerInstructions]: Instructions that specify how to layer and compose video frames from source tracks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class
type IAVVideoCompositionInstruction interface {
	objectivec.IObject

	// Topic: Inspecting the instruction

	// The background color of the composition.
	BackgroundColor() coregraphics.CGColorRef
	// Instructions that specify how to layer and compose video frames from source tracks.
	LayerInstructions() []AVVideoCompositionLayerInstruction

	// The video composition instructions.
	Instructions() objc.ID
	SetInstructions(value objc.ID)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (v AVVideoCompositionInstruction) Init() AVVideoCompositionInstruction {
	rv := objc.Send[AVVideoCompositionInstruction](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVideoCompositionInstruction) Autorelease() AVVideoCompositionInstruction {
	rv := objc.Send[AVVideoCompositionInstruction](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVideoCompositionInstruction creates a new AVVideoCompositionInstruction instance.
func NewAVVideoCompositionInstruction() AVVideoCompositionInstruction {
	class := getAVVideoCompositionInstructionClass()
	rv := objc.Send[AVVideoCompositionInstruction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v AVVideoCompositionInstruction) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Pass-through initializer, for internal use in AVFoundation only
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/videoCompositionInstructionWithInstruction:
func (_AVVideoCompositionInstructionClass AVVideoCompositionInstructionClass) VideoCompositionInstructionWithInstruction(instruction AVVideoCompositionInstruction) AVVideoCompositionInstruction {
	rv := objc.Send[objc.ID](objc.ID(_AVVideoCompositionInstructionClass.class), objc.Sel("videoCompositionInstructionWithInstruction:"), instruction)
	return AVVideoCompositionInstructionFromID(rv)
}

// The background color of the composition.
//
// # Discussion
// 
// Only solid BGRA colors are supported; patterns and other supported colors
// are ignored. If the rendered pixel buffer does not have alpha, the alpha
// value of the background color is ignored.
// 
// If the background color is [NULL], the video compositor uses a default
// background color of opaque black.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/backgroundColor
func (v AVVideoCompositionInstruction) BackgroundColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](v.ID, objc.Sel("backgroundColor"))
	return coregraphics.CGColorRef(rv)
}
// Instructions that specify how to layer and compose video frames from source
// tracks.
//
// # Discussion
// 
// Tracks are layered in the composition according to the top-to-bottom order
// of the `layerInstructions` array; the track with `trackID` of the first
// instruction in the array will be layered on top, with the track with the
// `trackID` of the second instruction immediately underneath, and so on.
// 
// If the property value is `nil`, the output is a fill of the background
// color.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/layerInstructions
func (v AVVideoCompositionInstruction) LayerInstructions() []AVVideoCompositionLayerInstruction {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("layerInstructions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVVideoCompositionLayerInstruction {
		return AVVideoCompositionLayerInstructionFromID(id)
	})
}
// The time range to which the instruction applies.
//
// # Discussion
// 
// If the time range is invalid, the video compositor will ignore it. See also
// the requirements of the [TimeRange] property in the array of objects
// implementing the [AVVideoCompositionInstructionProtocol] protocol as
// described in the [AVVideoComposition] class’s [Instructions] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/timeRange
func (v AVVideoCompositionInstruction) TimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](v.ID, objc.Sel("timeRange"))
	return coremedia.CMTimeRange(rv)
}
// A Boolean value that indicates whether the instruction requires post
// processing.
//
// # Discussion
// 
// A value of [false] indicates that no post processing is required for the
// whole duration of the video composition instruction. The composition
// process is more efficient if the value is [false].
// 
// The value is [true] by default.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/enablePostProcessing
func (v AVVideoCompositionInstruction) EnablePostProcessing() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("enablePostProcessing"))
	return rv
}
// The identifiers of source video tracks that the compositor requires to
// compose frames for the instruction.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/requiredSourceTrackIDs
func (v AVVideoCompositionInstruction) RequiredSourceTrackIDs() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("requiredSourceTrackIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
// The identifiers of source sample data tracks that the compositor requires
// to compose frames for the instruction.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/requiredSourceSampleDataTrackIDs
func (v AVVideoCompositionInstruction) RequiredSourceSampleDataTrackIDs() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("requiredSourceSampleDataTrackIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// The track identifier from an instruction source frame.
//
// # Discussion
// 
// If the video composition result is one of the source frames for the
// duration of the instruction, this property returns the corresponding track
// ID. The compositor won’t be run for the duration of the instruction and
// the proper source frame will be used instead. The value of this property is
// computed from the layer instructions
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/passthroughTrackID
func (v AVVideoCompositionInstruction) PassthroughTrackID() int32 {
	rv := objc.Send[int32](v.ID, objc.Sel("passthroughTrackID"))
	return rv
}
// The video composition instructions.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideocomposition/instructions
func (v AVVideoCompositionInstruction) Instructions() objc.ID {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("instructions"))
	return rv
}
func (v AVVideoCompositionInstruction) SetInstructions(value objc.ID) {
	objc.Send[struct{}](v.ID, objc.Sel("setInstructions:"), value)
}

			// Protocol methods for AVVideoCompositionInstructionProtocol
			
// A Boolean value that indicates whether the composition contains tweening.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstructionProtocol/containsTweening
func (o AVVideoCompositionInstruction) ContainsTweening() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("containsTweening"))
	return rv
	}


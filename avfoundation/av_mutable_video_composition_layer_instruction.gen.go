// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMutableVideoCompositionLayerInstruction] class.
var (
	_AVMutableVideoCompositionLayerInstructionClass     AVMutableVideoCompositionLayerInstructionClass
	_AVMutableVideoCompositionLayerInstructionClassOnce sync.Once
)

func getAVMutableVideoCompositionLayerInstructionClass() AVMutableVideoCompositionLayerInstructionClass {
	_AVMutableVideoCompositionLayerInstructionClassOnce.Do(func() {
		_AVMutableVideoCompositionLayerInstructionClass = AVMutableVideoCompositionLayerInstructionClass{class: objc.GetClass("AVMutableVideoCompositionLayerInstruction")}
	})
	return _AVMutableVideoCompositionLayerInstructionClass
}

// GetAVMutableVideoCompositionLayerInstructionClass returns the class object for AVMutableVideoCompositionLayerInstruction.
func GetAVMutableVideoCompositionLayerInstructionClass() AVMutableVideoCompositionLayerInstructionClass {
	return getAVMutableVideoCompositionLayerInstructionClass()
}

type AVMutableVideoCompositionLayerInstructionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableVideoCompositionLayerInstructionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableVideoCompositionLayerInstructionClass) Alloc() AVMutableVideoCompositionLayerInstruction {
	rv := objc.Send[AVMutableVideoCompositionLayerInstruction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object used to modify the transform, cropping, and opacity ramps applied
// to a given track in a mutable composition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction
type AVMutableVideoCompositionLayerInstruction struct {
	AVVideoCompositionLayerInstruction
}

// AVMutableVideoCompositionLayerInstructionFromID constructs a [AVMutableVideoCompositionLayerInstruction] from an objc.ID.
//
// An object used to modify the transform, cropping, and opacity ramps applied
// to a given track in a mutable composition.
func AVMutableVideoCompositionLayerInstructionFromID(id objc.ID) AVMutableVideoCompositionLayerInstruction {
	return AVMutableVideoCompositionLayerInstruction{AVVideoCompositionLayerInstruction: AVVideoCompositionLayerInstructionFromID(id)}
}

// NOTE: AVMutableVideoCompositionLayerInstruction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableVideoCompositionLayerInstruction] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction
type IAVMutableVideoCompositionLayerInstruction interface {
	IAVVideoCompositionLayerInstruction
}

// Init initializes the instance.
func (m AVMutableVideoCompositionLayerInstruction) Init() AVMutableVideoCompositionLayerInstruction {
	rv := objc.Send[AVMutableVideoCompositionLayerInstruction](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableVideoCompositionLayerInstruction) Autorelease() AVMutableVideoCompositionLayerInstruction {
	rv := objc.Send[AVMutableVideoCompositionLayerInstruction](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableVideoCompositionLayerInstruction creates a new AVMutableVideoCompositionLayerInstruction instance.
func NewAVMutableVideoCompositionLayerInstruction() AVMutableVideoCompositionLayerInstruction {
	class := getAVMutableVideoCompositionLayerInstructionClass()
	rv := objc.Send[AVMutableVideoCompositionLayerInstruction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new mutable video composition layer instruction for the given
// track.
//
// track: The asset track to which to apply the instruction.
//
// # Return Value
//
// A new mutable video composition layer instruction with no transform or
// opacity ramps and [TrackID] initialized to the track ID of `track`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/init(assetTrack:)
func NewMutableVideoCompositionLayerInstructionWithAssetTrack(track IAVAssetTrack) AVMutableVideoCompositionLayerInstruction {
	rv := objc.Send[objc.ID](objc.ID(getAVMutableVideoCompositionLayerInstructionClass().class), objc.Sel("videoCompositionLayerInstructionWithAssetTrack:"), track)
	return AVMutableVideoCompositionLayerInstructionFromID(rv)
}

// Returns a new mutable video composition layer instruction.
//
// # Return Value
//
// A new mutable video composition layer instruction with no transform or
// opacity ramps and [TrackID] initialized to [kCMPersistentTrackID_Invalid].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/videoCompositionLayerInstruction
//
// [kCMPersistentTrackID_Invalid]: https://developer.apple.com/documentation/CoreMedia/kCMPersistentTrackID_Invalid
func (_AVMutableVideoCompositionLayerInstructionClass AVMutableVideoCompositionLayerInstructionClass) VideoCompositionLayerInstruction() AVMutableVideoCompositionLayerInstruction {
	rv := objc.Send[objc.ID](objc.ID(_AVMutableVideoCompositionLayerInstructionClass.class), objc.Sel("videoCompositionLayerInstruction"))
	return AVMutableVideoCompositionLayerInstructionFromID(rv)
}

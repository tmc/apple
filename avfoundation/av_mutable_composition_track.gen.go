// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMutableCompositionTrack] class.
var (
	_AVMutableCompositionTrackClass     AVMutableCompositionTrackClass
	_AVMutableCompositionTrackClassOnce sync.Once
)

func getAVMutableCompositionTrackClass() AVMutableCompositionTrackClass {
	_AVMutableCompositionTrackClassOnce.Do(func() {
		_AVMutableCompositionTrackClass = AVMutableCompositionTrackClass{class: objc.GetClass("AVMutableCompositionTrack")}
	})
	return _AVMutableCompositionTrackClass
}

// GetAVMutableCompositionTrackClass returns the class object for AVMutableCompositionTrack.
func GetAVMutableCompositionTrackClass() AVMutableCompositionTrackClass {
	return getAVMutableCompositionTrackClass()
}

type AVMutableCompositionTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableCompositionTrackClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableCompositionTrackClass) Alloc() AVMutableCompositionTrack {
	rv := objc.Send[AVMutableCompositionTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable track in a composition that you use to insert, remove, and scale
// track segments without affecting their low-level representation.
//
// # Overview
// 
// Use this object to define constraints for the temporal arrangement of the
// track segments. If you set the composition’s track segments, you can test
// whether they meet the constraints by calling the
// [AVMutableCompositionTrack.ValidateTrackSegmentsError] method.
//
// # Configuring track properties
//
//   - [AVMutableCompositionTrack.Enabled]: A Boolean value that indicates whether the tracks is in an enabled state.
//   - [AVMutableCompositionTrack.SetEnabled]
//
// # Managing time ranges
//
//   - [AVMutableCompositionTrack.InsertEmptyTimeRange]: Adds or extends an empty time range within the track.
//   - [AVMutableCompositionTrack.InsertTimeRangeOfTrackAtTimeError]: Inserts a time range of media from a source track into a composition track.
//   - [AVMutableCompositionTrack.InsertTimeRangesOfTracksAtTimeError]: Inserts the time ranges of multiple source tracks into a track of a composition.
//   - [AVMutableCompositionTrack.RemoveTimeRange]: Removes a time range of media from a composition track.
//   - [AVMutableCompositionTrack.ScaleTimeRangeToDuration]: Changes the duration of a time range of the track.
//
// # Associating tracks
//
//   - [AVMutableCompositionTrack.AddTrackAssociationToTrackType]: Establishes a track association of a specific type between two tracks.
//   - [AVMutableCompositionTrack.RemoveTrackAssociationToTrackType]: Removes an association from a composition track.
//
// # Replacing format descriptions
//
//   - [AVMutableCompositionTrack.ReplaceFormatDescriptionWithFormatDescription]: Replaces a format description with another or cancels a previous replacement.
//
// # Validating segments
//
//   - [AVMutableCompositionTrack.ValidateTrackSegmentsError]: Returns a Boolean value that indicates whether a given array of track segments conform to the timing rules for a composition track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack
type AVMutableCompositionTrack struct {
	AVCompositionTrack
}

// AVMutableCompositionTrackFromID constructs a [AVMutableCompositionTrack] from an objc.ID.
//
// A mutable track in a composition that you use to insert, remove, and scale
// track segments without affecting their low-level representation.
func AVMutableCompositionTrackFromID(id objc.ID) AVMutableCompositionTrack {
	return AVMutableCompositionTrack{AVCompositionTrack: AVCompositionTrackFromID(id)}
}
// NOTE: AVMutableCompositionTrack adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableCompositionTrack] class.
//
// # Configuring track properties
//
//   - [IAVMutableCompositionTrack.Enabled]: A Boolean value that indicates whether the tracks is in an enabled state.
//   - [IAVMutableCompositionTrack.SetEnabled]
//
// # Managing time ranges
//
//   - [IAVMutableCompositionTrack.InsertEmptyTimeRange]: Adds or extends an empty time range within the track.
//   - [IAVMutableCompositionTrack.InsertTimeRangeOfTrackAtTimeError]: Inserts a time range of media from a source track into a composition track.
//   - [IAVMutableCompositionTrack.InsertTimeRangesOfTracksAtTimeError]: Inserts the time ranges of multiple source tracks into a track of a composition.
//   - [IAVMutableCompositionTrack.RemoveTimeRange]: Removes a time range of media from a composition track.
//   - [IAVMutableCompositionTrack.ScaleTimeRangeToDuration]: Changes the duration of a time range of the track.
//
// # Associating tracks
//
//   - [IAVMutableCompositionTrack.AddTrackAssociationToTrackType]: Establishes a track association of a specific type between two tracks.
//   - [IAVMutableCompositionTrack.RemoveTrackAssociationToTrackType]: Removes an association from a composition track.
//
// # Replacing format descriptions
//
//   - [IAVMutableCompositionTrack.ReplaceFormatDescriptionWithFormatDescription]: Replaces a format description with another or cancels a previous replacement.
//
// # Validating segments
//
//   - [IAVMutableCompositionTrack.ValidateTrackSegmentsError]: Returns a Boolean value that indicates whether a given array of track segments conform to the timing rules for a composition track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack
type IAVMutableCompositionTrack interface {
	IAVCompositionTrack

	// Topic: Configuring track properties

	// A Boolean value that indicates whether the tracks is in an enabled state.
	Enabled() bool
	SetEnabled(value bool)

	// Topic: Managing time ranges

	// Adds or extends an empty time range within the track.
	InsertEmptyTimeRange(timeRange uintptr)
	// Inserts a time range of media from a source track into a composition track.
	InsertTimeRangeOfTrackAtTimeError(timeRange uintptr, track IAVAssetTrack, startTime uintptr) (bool, error)
	// Inserts the time ranges of multiple source tracks into a track of a composition.
	InsertTimeRangesOfTracksAtTimeError(timeRanges []foundation.NSValue, tracks []AVAssetTrack, startTime uintptr) (bool, error)
	// Removes a time range of media from a composition track.
	RemoveTimeRange(timeRange uintptr)
	// Changes the duration of a time range of the track.
	ScaleTimeRangeToDuration(timeRange uintptr, duration uintptr)

	// Topic: Associating tracks

	// Establishes a track association of a specific type between two tracks.
	AddTrackAssociationToTrackType(compositionTrack IAVCompositionTrack, trackAssociationType AVTrackAssociationType)
	// Removes an association from a composition track.
	RemoveTrackAssociationToTrackType(compositionTrack IAVCompositionTrack, trackAssociationType AVTrackAssociationType)

	// Topic: Replacing format descriptions

	// Replaces a format description with another or cancels a previous replacement.
	ReplaceFormatDescriptionWithFormatDescription(originalFormatDescription uintptr, replacementFormatDescription uintptr)

	// Topic: Validating segments

	// Returns a Boolean value that indicates whether a given array of track segments conform to the timing rules for a composition track.
	ValidateTrackSegmentsError(trackSegments []AVCompositionTrackSegment) (bool, error)
}

// Init initializes the instance.
func (m AVMutableCompositionTrack) Init() AVMutableCompositionTrack {
	rv := objc.Send[AVMutableCompositionTrack](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableCompositionTrack) Autorelease() AVMutableCompositionTrack {
	rv := objc.Send[AVMutableCompositionTrack](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableCompositionTrack creates a new AVMutableCompositionTrack instance.
func NewAVMutableCompositionTrack() AVMutableCompositionTrack {
	class := getAVMutableCompositionTrackClass()
	rv := objc.Send[AVMutableCompositionTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds or extends an empty time range within the track.
//
// timeRange: The empty time range to insert.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/insertEmptyTimeRange(_:)
func (m AVMutableCompositionTrack) InsertEmptyTimeRange(timeRange uintptr) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertEmptyTimeRange:"), timeRange)
}
// Inserts a time range of media from a source track into a composition track.
//
// timeRange: The time range of media in the source track to add.
//
// track: The source asset track that contains the media to add.
//
// startTime: A start time within the composition track to insert the time range.
//
// # Discussion
// 
// The time range you insert presents at its natural duration and rate. If
// necessary, you can scale it to a different duration by calling the
// [ScaleTimeRangeToDuration] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/insertTimeRange(_:of:at:)
func (m AVMutableCompositionTrack) InsertTimeRangeOfTrackAtTimeError(timeRange uintptr, track IAVAssetTrack, startTime uintptr) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("insertTimeRange:ofTrack:atTime:error:"), timeRange, track, startTime, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("insertTimeRange:ofTrack:atTime:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Inserts the time ranges of multiple source tracks into a track of a
// composition.
//
// timeRanges: The time ranges of media in the source tracks to insert.
//
// tracks: The source asset tracks that contain the media to insert.
//
// startTime: A start time within composition the track to insert the time range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/insertTimeRanges(_:of:at:)
func (m AVMutableCompositionTrack) InsertTimeRangesOfTracksAtTimeError(timeRanges []foundation.NSValue, tracks []AVAssetTrack, startTime uintptr) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("insertTimeRanges:ofTracks:atTime:error:"), objectivec.IObjectSliceToNSArray(timeRanges), objectivec.IObjectSliceToNSArray(tracks), startTime, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("insertTimeRanges:ofTracks:atTime:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Removes a time range of media from a composition track.
//
// timeRange: The time range to remove.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/removeTimeRange(_:)
func (m AVMutableCompositionTrack) RemoveTimeRange(timeRange uintptr) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTimeRange:"), timeRange)
}
// Changes the duration of a time range of the track.
//
// timeRange: The time range to scale.
//
// duration: A new duration value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/scaleTimeRange(_:toDuration:)
func (m AVMutableCompositionTrack) ScaleTimeRangeToDuration(timeRange uintptr, duration uintptr) {
	objc.Send[objc.ID](m.ID, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}
// Establishes a track association of a specific type between two tracks.
//
// compositionTrack: A composition track to associate.
//
// trackAssociationType: The type of track association to create between tracks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/addTrackAssociation(to:type:)
func (m AVMutableCompositionTrack) AddTrackAssociationToTrackType(compositionTrack IAVCompositionTrack, trackAssociationType AVTrackAssociationType) {
	objc.Send[objc.ID](m.ID, objc.Sel("addTrackAssociationToTrack:type:"), compositionTrack, objc.String(string(trackAssociationType)))
}
// Removes an association from a composition track.
//
// compositionTrack: A composition track to remove the association from.
//
// trackAssociationType: The type of track association to remove.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/removeTrackAssociation(to:type:)
func (m AVMutableCompositionTrack) RemoveTrackAssociationToTrackType(compositionTrack IAVCompositionTrack, trackAssociationType AVTrackAssociationType) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTrackAssociationToTrack:type:"), compositionTrack, objc.String(string(trackAssociationType)))
}
// Replaces a format description with another or cancels a previous
// replacement.
//
// originalFormatDescription: The original format description.
//
// replacementFormatDescription: A replacement format description.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/replaceFormatDescription(_:with:)
func (m AVMutableCompositionTrack) ReplaceFormatDescriptionWithFormatDescription(originalFormatDescription uintptr, replacementFormatDescription uintptr) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceFormatDescription:withFormatDescription:"), originalFormatDescription, replacementFormatDescription)
}
// Returns a Boolean value that indicates whether a given array of track
// segments conform to the timing rules for a composition track.
//
// trackSegments: The track segments to validate.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/validateSegments(_:)
func (m AVMutableCompositionTrack) ValidateTrackSegmentsError(trackSegments []AVCompositionTrackSegment) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("validateTrackSegments:error:"), objectivec.IObjectSliceToNSArray(trackSegments), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateTrackSegments:error: returned NO with nil NSError")
	}
	return rv, nil

}

// A Boolean value that indicates whether the tracks is in an enabled state.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCompositionTrack/isEnabled
func (m AVMutableCompositionTrack) Enabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEnabled"))
	return rv
}
func (m AVMutableCompositionTrack) SetEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setEnabled:"), value)
}


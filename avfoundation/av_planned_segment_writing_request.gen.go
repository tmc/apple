// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlannedSegmentWritingRequest] class.
var (
	_AVPlannedSegmentWritingRequestClass     AVPlannedSegmentWritingRequestClass
	_AVPlannedSegmentWritingRequestClassOnce sync.Once
)

func getAVPlannedSegmentWritingRequestClass() AVPlannedSegmentWritingRequestClass {
	_AVPlannedSegmentWritingRequestClassOnce.Do(func() {
		_AVPlannedSegmentWritingRequestClass = AVPlannedSegmentWritingRequestClass{class: objc.GetClass("AVPlannedSegmentWritingRequest")}
	})
	return _AVPlannedSegmentWritingRequestClass
}

// GetAVPlannedSegmentWritingRequestClass returns the class object for AVPlannedSegmentWritingRequest.
func GetAVPlannedSegmentWritingRequestClass() AVPlannedSegmentWritingRequestClass {
	return getAVPlannedSegmentWritingRequestClass()
}

type AVPlannedSegmentWritingRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlannedSegmentWritingRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlannedSegmentWritingRequestClass) Alloc() AVPlannedSegmentWritingRequest {
	rv := objc.Send[AVPlannedSegmentWritingRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// AVPlannedSegmentWritingRequest encompasses a request from the
// AVAssetWritingPlanner to the client code to write one incremental track
// segment.
//
// # Overview
//
// The client should respond to this request by writing the specified time
// range of data to a movie file at the specified segmentFileOutputURL, with
// start PTS zero. The client’s writing work may be completed
// asynchronously. If it completes successfully, clients must call the
// `-finish` or `-finishWithClientState` method on the request object. If
// writing the segment fails, clients must call the `-` method on the request
// object. If segment writing needs to be stopped before reaching the end of
// the segment, clients must call `-cancel`.
//
// # Inspecting the request
//
//   - [AVPlannedSegmentWritingRequest.TimeRange]: The PTS range for this segment.
//   - [AVPlannedSegmentWritingRequest.SegmentFileOutputURL]: The URL of the file where this incremental segment should be written to.
//   - [AVPlannedSegmentWritingRequest.Progress]: The current progress for the track identified by assemblyTrackID.
//   - [AVPlannedSegmentWritingRequest.AssemblyTrackID]: The trackID identifies which track should be written to this segment file. This is the same track ID in the AVAssetTrackPlan object. This is also the trackID the AVAssetWritingPlanner uses to build the assembled AVComposition before it calls the completion handler.
//
// # Managing client state
//
//   - [AVPlannedSegmentWritingRequest.ClientStateToRestore]: The client state persisted from the previous segment, if any. Specifically, this is the NSData provided to the previous segment’s finishWithClientState: method. The client is responsible to restore its client state before writing the current segment. For example, clients such as compositors with a temporal element may need some processing history of previous samples in order to generate an output sample at time N. This will be nil for algorithms that are stateless.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentWritingRequest
type AVPlannedSegmentWritingRequest struct {
	objectivec.Object
}

// AVPlannedSegmentWritingRequestFromID constructs a [AVPlannedSegmentWritingRequest] from an objc.ID.
//
// AVPlannedSegmentWritingRequest encompasses a request from the
// AVAssetWritingPlanner to the client code to write one incremental track
// segment.
func AVPlannedSegmentWritingRequestFromID(id objc.ID) AVPlannedSegmentWritingRequest {
	return AVPlannedSegmentWritingRequest{objectivec.Object{ID: id}}
}

// NOTE: AVPlannedSegmentWritingRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlannedSegmentWritingRequest] class.
//
// # Inspecting the request
//
//   - [IAVPlannedSegmentWritingRequest.TimeRange]: The PTS range for this segment.
//   - [IAVPlannedSegmentWritingRequest.SegmentFileOutputURL]: The URL of the file where this incremental segment should be written to.
//   - [IAVPlannedSegmentWritingRequest.Progress]: The current progress for the track identified by assemblyTrackID.
//   - [IAVPlannedSegmentWritingRequest.AssemblyTrackID]: The trackID identifies which track should be written to this segment file. This is the same track ID in the AVAssetTrackPlan object. This is also the trackID the AVAssetWritingPlanner uses to build the assembled AVComposition before it calls the completion handler.
//
// # Managing client state
//
//   - [IAVPlannedSegmentWritingRequest.ClientStateToRestore]: The client state persisted from the previous segment, if any. Specifically, this is the NSData provided to the previous segment’s finishWithClientState: method. The client is responsible to restore its client state before writing the current segment. For example, clients such as compositors with a temporal element may need some processing history of previous samples in order to generate an output sample at time N. This will be nil for algorithms that are stateless.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentWritingRequest
type IAVPlannedSegmentWritingRequest interface {
	objectivec.IObject

	// Topic: Inspecting the request

	// The PTS range for this segment.
	TimeRange() coremedia.CMTimeRange
	// The URL of the file where this incremental segment should be written to.
	SegmentFileOutputURL() foundation.NSURL
	// The current progress for the track identified by assemblyTrackID.
	Progress() float32
	// The trackID identifies which track should be written to this segment file. This is the same track ID in the AVAssetTrackPlan object. This is also the trackID the AVAssetWritingPlanner uses to build the assembled AVComposition before it calls the completion handler.
	AssemblyTrackID() coremedia.CMPersistentTrackID

	// Topic: Managing client state

	// The client state persisted from the previous segment, if any. Specifically, this is the NSData provided to the previous segment’s finishWithClientState: method. The client is responsible to restore its client state before writing the current segment. For example, clients such as compositors with a temporal element may need some processing history of previous samples in order to generate an output sample at time N. This will be nil for algorithms that are stateless.
	ClientStateToRestore() foundation.NSData
}

// Init initializes the instance.
func (p AVPlannedSegmentWritingRequest) Init() AVPlannedSegmentWritingRequest {
	rv := objc.Send[AVPlannedSegmentWritingRequest](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlannedSegmentWritingRequest) Autorelease() AVPlannedSegmentWritingRequest {
	rv := objc.Send[AVPlannedSegmentWritingRequest](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlannedSegmentWritingRequest creates a new AVPlannedSegmentWritingRequest instance.
func NewAVPlannedSegmentWritingRequest() AVPlannedSegmentWritingRequest {
	class := getAVPlannedSegmentWritingRequestClass()
	rv := objc.Send[AVPlannedSegmentWritingRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The PTS range for this segment.
//
// # Discussion
//
// The client is responsible for delivering the appropriate sample
// corresponding to timeRange.start if we are resuming a previous session that
// has already made incremental progress for this track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentWritingRequest/timeRange
func (p AVPlannedSegmentWritingRequest) TimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](p.ID, objc.Sel("timeRange"))
	return coremedia.CMTimeRange(rv)
}

// The URL of the file where this incremental segment should be written to.
//
// # Discussion
//
// AVAssetWritingPlanner will request each incremental segment to be written
// to a different file. If the file already exists from a previous session,
// the client should delete it to allow the subsequent asset writer session to
// succeed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentWritingRequest/segmentFileOutputURL
func (p AVPlannedSegmentWritingRequest) SegmentFileOutputURL() foundation.NSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("segmentFileOutputURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The current progress for the track identified by assemblyTrackID.
//
// # Discussion
//
// Returns a float value between 0.0 and 1.0 representing the percentage of
// duration completed for this track. This value is updated as segments are
// completed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentWritingRequest/progress
func (p AVPlannedSegmentWritingRequest) Progress() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("progress"))
	return rv
}

// The trackID identifies which track should be written to this segment file.
// This is the same track ID in the AVAssetTrackPlan object. This is also the
// trackID the AVAssetWritingPlanner uses to build the assembled AVComposition
// before it calls the completion handler.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentWritingRequest/assemblyTrackID
func (p AVPlannedSegmentWritingRequest) AssemblyTrackID() coremedia.CMPersistentTrackID {
	rv := objc.Send[coremedia.CMPersistentTrackID](p.ID, objc.Sel("assemblyTrackID"))
	return coremedia.CMPersistentTrackID(rv)
}

// The client state persisted from the previous segment, if any. Specifically,
// this is the NSData provided to the previous segment’s
// finishWithClientState: method. The client is responsible to restore its
// client state before writing the current segment. For example, clients such
// as compositors with a temporal element may need some processing history of
// previous samples in order to generate an output sample at time N. This will
// be nil for algorithms that are stateless.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentWritingRequest/clientStateToRestore
func (p AVPlannedSegmentWritingRequest) ClientStateToRestore() foundation.NSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("clientStateToRestore"))
	return foundation.NSDataFromID(objc.ID(rv))
}

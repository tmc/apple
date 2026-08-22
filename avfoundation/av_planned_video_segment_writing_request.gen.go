// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVPlannedVideoSegmentWritingRequest] class.
var (
	_AVPlannedVideoSegmentWritingRequestClass     AVPlannedVideoSegmentWritingRequestClass
	_AVPlannedVideoSegmentWritingRequestClassOnce sync.Once
)

func getAVPlannedVideoSegmentWritingRequestClass() AVPlannedVideoSegmentWritingRequestClass {
	_AVPlannedVideoSegmentWritingRequestClassOnce.Do(func() {
		_AVPlannedVideoSegmentWritingRequestClass = AVPlannedVideoSegmentWritingRequestClass{class: objc.GetClass("AVPlannedVideoSegmentWritingRequest")}
	})
	return _AVPlannedVideoSegmentWritingRequestClass
}

// GetAVPlannedVideoSegmentWritingRequestClass returns the class object for AVPlannedVideoSegmentWritingRequest.
func GetAVPlannedVideoSegmentWritingRequestClass() AVPlannedVideoSegmentWritingRequestClass {
	return getAVPlannedVideoSegmentWritingRequestClass()
}

type AVPlannedVideoSegmentWritingRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlannedVideoSegmentWritingRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlannedVideoSegmentWritingRequestClass) Alloc() AVPlannedVideoSegmentWritingRequest {
	rv := objc.Send[AVPlannedVideoSegmentWritingRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// AVPlannedVideoSegmentWritingRequest encompasses a request from the
// AVAssetWritingPlanner to the client code to write one incremental video
// track segment with compression.
//
// # Overview
//
// The client should respond to this request by writing the specified time
// range of data to a movie file at the specified segmentFileOutputURL, with
// start PTS zero. The client’s writing work may be completed
// asynchronously. If it completes successfully, it must call the `-finish`
// method on the request object. If writing the segment fails, it must call
// the `-` method on the request object.
//
// # Inspecting the request
//
//   - [AVPlannedVideoSegmentWritingRequest.FrameCount]: The number of frames in this planned video segment. This is provided for convenience, and is the same value that was configured for the segment in AVPlannedVideoSegmentConfiguration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedVideoSegmentWritingRequest
type AVPlannedVideoSegmentWritingRequest struct {
	AVPlannedSegmentWritingRequest
}

// AVPlannedVideoSegmentWritingRequestFromID constructs a [AVPlannedVideoSegmentWritingRequest] from an objc.ID.
//
// AVPlannedVideoSegmentWritingRequest encompasses a request from the
// AVAssetWritingPlanner to the client code to write one incremental video
// track segment with compression.
func AVPlannedVideoSegmentWritingRequestFromID(id objc.ID) AVPlannedVideoSegmentWritingRequest {
	return AVPlannedVideoSegmentWritingRequest{AVPlannedSegmentWritingRequest: AVPlannedSegmentWritingRequestFromID(id)}
}

// NOTE: AVPlannedVideoSegmentWritingRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlannedVideoSegmentWritingRequest] class.
//
// # Inspecting the request
//
//   - [IAVPlannedVideoSegmentWritingRequest.FrameCount]: The number of frames in this planned video segment. This is provided for convenience, and is the same value that was configured for the segment in AVPlannedVideoSegmentConfiguration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedVideoSegmentWritingRequest
type IAVPlannedVideoSegmentWritingRequest interface {
	IAVPlannedSegmentWritingRequest

	// Topic: Inspecting the request

	// The number of frames in this planned video segment. This is provided for convenience, and is the same value that was configured for the segment in AVPlannedVideoSegmentConfiguration.
	FrameCount() int
}

// Init initializes the instance.
func (p AVPlannedVideoSegmentWritingRequest) Init() AVPlannedVideoSegmentWritingRequest {
	rv := objc.Send[AVPlannedVideoSegmentWritingRequest](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlannedVideoSegmentWritingRequest) Autorelease() AVPlannedVideoSegmentWritingRequest {
	rv := objc.Send[AVPlannedVideoSegmentWritingRequest](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlannedVideoSegmentWritingRequest creates a new AVPlannedVideoSegmentWritingRequest instance.
func NewAVPlannedVideoSegmentWritingRequest() AVPlannedVideoSegmentWritingRequest {
	class := getAVPlannedVideoSegmentWritingRequestClass()
	rv := objc.Send[AVPlannedVideoSegmentWritingRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The number of frames in this planned video segment. This is provided for
// convenience, and is the same value that was configured for the segment in
// AVPlannedVideoSegmentConfiguration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedVideoSegmentWritingRequest/frameCount
func (p AVPlannedVideoSegmentWritingRequest) FrameCount() int {
	rv := objc.Send[int](p.ID, objc.Sel("frameCount"))
	return rv
}

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetVideoTrackPlan] class.
var (
	_AVAssetVideoTrackPlanClass     AVAssetVideoTrackPlanClass
	_AVAssetVideoTrackPlanClassOnce sync.Once
)

func getAVAssetVideoTrackPlanClass() AVAssetVideoTrackPlanClass {
	_AVAssetVideoTrackPlanClassOnce.Do(func() {
		_AVAssetVideoTrackPlanClass = AVAssetVideoTrackPlanClass{class: objc.GetClass("AVAssetVideoTrackPlan")}
	})
	return _AVAssetVideoTrackPlanClass
}

// GetAVAssetVideoTrackPlanClass returns the class object for AVAssetVideoTrackPlan.
func GetAVAssetVideoTrackPlanClass() AVAssetVideoTrackPlanClass {
	return getAVAssetVideoTrackPlanClass()
}

type AVAssetVideoTrackPlanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetVideoTrackPlanClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetVideoTrackPlanClass) Alloc() AVAssetVideoTrackPlan {
	rv := objc.Send[AVAssetVideoTrackPlan](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// AVAssetVideoTrackPlan holds information about a track and how it should be
// segmented and executed in an incremental writing session.
//
// # Overview
//
// Call AVAssetWritingPlanner’s “planTrack:withSegmentsGeneratedBy:”
// method to add an AVAssetTrackPlan to the planner’s plan to include it in
// the incremental writing session. Use this class instead of the base class
// AVAssetTrackPlan if you are setting up AVAssetWriter with video
// compression. This configuration hints to the planner that it must
// coordinate segment boundaries transitions between segments. This is
// abstracted from the client via using either the
// resumableAssetWriterInputWithMediaType or
// createResumableCompressionSessionWithAllocator helper functions within the
// AVPlannedVideoSegmentWritingRequest.
//
// # Inspecting the video track plan
//
//   - [AVAssetVideoTrackPlan.VideoCodecType]: Video codec type of this track
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVideoTrackPlan
type AVAssetVideoTrackPlan struct {
	AVAssetTrackPlan
}

// AVAssetVideoTrackPlanFromID constructs a [AVAssetVideoTrackPlan] from an objc.ID.
//
// AVAssetVideoTrackPlan holds information about a track and how it should be
// segmented and executed in an incremental writing session.
func AVAssetVideoTrackPlanFromID(id objc.ID) AVAssetVideoTrackPlan {
	return AVAssetVideoTrackPlan{AVAssetTrackPlan: AVAssetTrackPlanFromID(id)}
}

// NOTE: AVAssetVideoTrackPlan adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetVideoTrackPlan] class.
//
// # Inspecting the video track plan
//
//   - [IAVAssetVideoTrackPlan.VideoCodecType]: Video codec type of this track
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVideoTrackPlan
type IAVAssetVideoTrackPlan interface {
	IAVAssetTrackPlan

	// Topic: Inspecting the video track plan

	// Video codec type of this track
	VideoCodecType() AVVideoCodecType
}

// Init initializes the instance.
func (a AVAssetVideoTrackPlan) Init() AVAssetVideoTrackPlan {
	rv := objc.Send[AVAssetVideoTrackPlan](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetVideoTrackPlan) Autorelease() AVAssetVideoTrackPlan {
	rv := objc.Send[AVAssetVideoTrackPlan](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetVideoTrackPlan creates a new AVAssetVideoTrackPlan instance.
func NewAVAssetVideoTrackPlan() AVAssetVideoTrackPlan {
	class := getAVAssetVideoTrackPlanClass()
	rv := objc.Send[AVAssetVideoTrackPlan](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an instance of AVAssetTrackPlan
//
// mediaType: Media type of the track
//
// segmentConfigurations: Segment configurations of the track
//
// trackID: The trackID that identifies this track in the assemblyComposition the
// planner passes to the completion handler of the incremental writing
// session.
//
// # Discussion
//
// This initializer throws NSInvalidArgumentException if trackID is
// kCMPersistentTrackID_Invalid.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackPlan/init(mediaType:segmentConfigurations:assemblyTrackID:)
func NewAssetVideoTrackPlanWithMediaTypeSegmentConfigurationsAssemblyTrackID(mediaType AVMediaType, segmentConfigurations []AVPlannedSegmentConfiguration, trackID coremedia.CMPersistentTrackID) AVAssetVideoTrackPlan {
	instance := getAVAssetVideoTrackPlanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMediaType:segmentConfigurations:assemblyTrackID:"), objc.String(string(mediaType)), objectivec.IObjectSliceToNSArray(segmentConfigurations), trackID)
	return AVAssetVideoTrackPlanFromID(rv)
}

// Video codec type of this track
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVideoTrackPlan/videoCodecType
func (a AVAssetVideoTrackPlan) VideoCodecType() AVVideoCodecType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoCodecType"))
	return AVVideoCodecType(foundation.NSStringFromID(rv).String())
}

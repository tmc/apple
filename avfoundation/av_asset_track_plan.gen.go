// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetTrackPlan] class.
var (
	_AVAssetTrackPlanClass     AVAssetTrackPlanClass
	_AVAssetTrackPlanClassOnce sync.Once
)

func getAVAssetTrackPlanClass() AVAssetTrackPlanClass {
	_AVAssetTrackPlanClassOnce.Do(func() {
		_AVAssetTrackPlanClass = AVAssetTrackPlanClass{class: objc.GetClass("AVAssetTrackPlan")}
	})
	return _AVAssetTrackPlanClass
}

// GetAVAssetTrackPlanClass returns the class object for AVAssetTrackPlan.
func GetAVAssetTrackPlanClass() AVAssetTrackPlanClass {
	return getAVAssetTrackPlanClass()
}

type AVAssetTrackPlanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetTrackPlanClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetTrackPlanClass) Alloc() AVAssetTrackPlan {
	rv := objc.Send[AVAssetTrackPlan](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// AVAssetTrackPlan holds information about a track and how it should be
// segmented and executed in an incremental writing session.
//
// # Overview
//
// Call AVAssetWritingPlanner’s “planTrack:withSegmentsGeneratedBy:”
// method to add an AVAssetTrackPlan to the planner to include it in the
// incremental writing session.
//
// # Creating a track plan
//
//   - [AVAssetTrackPlan.InitWithMediaTypeSegmentConfigurationsAssemblyTrackID]: Returns an instance of AVAssetTrackPlan
//
// # Inspecting the track plan
//
//   - [AVAssetTrackPlan.MediaType]: The media type of this track.
//   - [AVAssetTrackPlan.SegmentConfigurations]: Array of AVPlannedSegmentConfigurations, each element specifying the configuration of a planned segment, ordered in output PTS order.
//   - [AVAssetTrackPlan.AssemblyTrackID]: This is the track ID of this track when it is included in the assemblyComposition the planner passes to the completion handler to assemble all planned segments of all tracks into a single AVComposition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackPlan
type AVAssetTrackPlan struct {
	objectivec.Object
}

// AVAssetTrackPlanFromID constructs a [AVAssetTrackPlan] from an objc.ID.
//
// AVAssetTrackPlan holds information about a track and how it should be
// segmented and executed in an incremental writing session.
func AVAssetTrackPlanFromID(id objc.ID) AVAssetTrackPlan {
	return AVAssetTrackPlan{objectivec.Object{ID: id}}
}

// NOTE: AVAssetTrackPlan adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetTrackPlan] class.
//
// # Creating a track plan
//
//   - [IAVAssetTrackPlan.InitWithMediaTypeSegmentConfigurationsAssemblyTrackID]: Returns an instance of AVAssetTrackPlan
//
// # Inspecting the track plan
//
//   - [IAVAssetTrackPlan.MediaType]: The media type of this track.
//   - [IAVAssetTrackPlan.SegmentConfigurations]: Array of AVPlannedSegmentConfigurations, each element specifying the configuration of a planned segment, ordered in output PTS order.
//   - [IAVAssetTrackPlan.AssemblyTrackID]: This is the track ID of this track when it is included in the assemblyComposition the planner passes to the completion handler to assemble all planned segments of all tracks into a single AVComposition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackPlan
type IAVAssetTrackPlan interface {
	objectivec.IObject

	// Topic: Creating a track plan

	// Returns an instance of AVAssetTrackPlan
	InitWithMediaTypeSegmentConfigurationsAssemblyTrackID(mediaType AVMediaType, segmentConfigurations []AVPlannedSegmentConfiguration, trackID coremedia.CMPersistentTrackID) AVAssetTrackPlan

	// Topic: Inspecting the track plan

	// The media type of this track.
	MediaType() AVMediaType
	// Array of AVPlannedSegmentConfigurations, each element specifying the configuration of a planned segment, ordered in output PTS order.
	SegmentConfigurations() []AVPlannedSegmentConfiguration
	// This is the track ID of this track when it is included in the assemblyComposition the planner passes to the completion handler to assemble all planned segments of all tracks into a single AVComposition.
	AssemblyTrackID() coremedia.CMPersistentTrackID
}

// Init initializes the instance.
func (a AVAssetTrackPlan) Init() AVAssetTrackPlan {
	rv := objc.Send[AVAssetTrackPlan](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetTrackPlan) Autorelease() AVAssetTrackPlan {
	rv := objc.Send[AVAssetTrackPlan](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetTrackPlan creates a new AVAssetTrackPlan instance.
func NewAVAssetTrackPlan() AVAssetTrackPlan {
	class := getAVAssetTrackPlanClass()
	rv := objc.Send[AVAssetTrackPlan](objc.ID(class.class), objc.Sel("new"))
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
func NewAssetTrackPlanWithMediaTypeSegmentConfigurationsAssemblyTrackID(mediaType AVMediaType, segmentConfigurations []AVPlannedSegmentConfiguration, trackID coremedia.CMPersistentTrackID) AVAssetTrackPlan {
	instance := getAVAssetTrackPlanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMediaType:segmentConfigurations:assemblyTrackID:"), objc.String(string(mediaType)), objectivec.IObjectSliceToNSArray(segmentConfigurations), trackID)
	return AVAssetTrackPlanFromID(rv)
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
func (a AVAssetTrackPlan) InitWithMediaTypeSegmentConfigurationsAssemblyTrackID(mediaType AVMediaType, segmentConfigurations []AVPlannedSegmentConfiguration, trackID coremedia.CMPersistentTrackID) AVAssetTrackPlan {
	rv := objc.Send[AVAssetTrackPlan](a.ID, objc.Sel("initWithMediaType:segmentConfigurations:assemblyTrackID:"), objc.String(string(mediaType)), objectivec.IObjectSliceToNSArray(segmentConfigurations), trackID)
	return rv
}

// The media type of this track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackPlan/mediaType
func (a AVAssetTrackPlan) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}

// Array of AVPlannedSegmentConfigurations, each element specifying the
// configuration of a planned segment, ordered in output PTS order.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackPlan/segmentConfigurations
func (a AVAssetTrackPlan) SegmentConfigurations() []AVPlannedSegmentConfiguration {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("segmentConfigurations"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlannedSegmentConfiguration {
		return AVPlannedSegmentConfigurationFromID(id)
	})
}

// This is the track ID of this track when it is included in the
// assemblyComposition the planner passes to the completion handler to
// assemble all planned segments of all tracks into a single AVComposition.
//
// # Discussion
//
// The assemblyTrackID serves the purpose as a unique identifier of the track
// in the incremental writing session. This does not necessarily match the
// trackID of the source asset. The client is responsible for remembering the
// relationship between assemblyTrackID and the trackID in the source asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackPlan/assemblyTrackID
func (a AVAssetTrackPlan) AssemblyTrackID() coremedia.CMPersistentTrackID {
	rv := objc.Send[coremedia.CMPersistentTrackID](a.ID, objc.Sel("assemblyTrackID"))
	return coremedia.CMPersistentTrackID(rv)
}

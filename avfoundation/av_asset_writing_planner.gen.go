// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetWritingPlanner] class.
var (
	_AVAssetWritingPlannerClass     AVAssetWritingPlannerClass
	_AVAssetWritingPlannerClassOnce sync.Once
)

func getAVAssetWritingPlannerClass() AVAssetWritingPlannerClass {
	_AVAssetWritingPlannerClassOnce.Do(func() {
		_AVAssetWritingPlannerClass = AVAssetWritingPlannerClass{class: objc.GetClass("AVAssetWritingPlanner")}
	})
	return _AVAssetWritingPlannerClass
}

// GetAVAssetWritingPlannerClass returns the class object for AVAssetWritingPlanner.
func GetAVAssetWritingPlannerClass() AVAssetWritingPlannerClass {
	return getAVAssetWritingPlannerClass()
}

type AVAssetWritingPlannerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetWritingPlannerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetWritingPlannerClass) Alloc() AVAssetWritingPlanner {
	rv := objc.Send[AVAssetWritingPlanner](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// AVAssetWritingPlanner orchestrates incremental writing of media files.
//
// # Overview
//
// AVAssetWritingPlanner orchestrates an incremental and resumable asset file
// writing session. It keeps track of the progress of the incremental
// segments, and can resume the writing from the last checkpoint. This is NOT
// intended for any real time applications. Also, not all tracks can be
// written incrementally. The workflow is as follows:
//
// - The client creates the planner with a unique directoryForTemporaryFiles.
// - The client tells the planner which tracks are to be written incrementally
// by calling the “planTrack:withSegmentsGeneratedBy:” method, providing a
// callback block that writes one segment per block invocation. - The client
// kicks off the incremental writing session by calling the
// “executePlanWithCompletionHandler” method. - The planner will call the
// writingSegmentCallbackBlock to ask the client to write one incremental
// segment of one track at a time. The client code should write one
// incremental segment according to the “AVPlannedSegmentWritingRequest”
// object passed in to the callback block. Clients must call “finish” or
// “finishWithError” or “finishWithClientState” or “cancel”
// methods on the request object when it finishes the segment successfully, or
// encountered an error, or wants to cancel the writing of the segment. - At
// the end of the writing, after all incremental segments are finished, the
// planner calls the completionHandler. The client can use the
// “assemblyComposition” object passed in to the completionHandler to
// assemble the incremental segments into full tracks and export it to a final
// output file. The completionHandler will also be called when there is any
// irrecoverable error. - The client is responsible for cleaning all files in
// the directoryForTemporaryFiles after the incremental session is done and
// the final output file is written.
//
// AVAssetWritingPlanner is able to recognize when a plan-in-progress matching
// the plan was already saved at directoryForTemporaryFiles, presumably by a
// previous invocation of the client, and possibly aborted due to that client
// being terminated abruptly, and will assist by resuming the plan at the
// first step that wasn’t previously completed.
//
// # Getting progress
//
//   - [AVAssetWritingPlanner.Progress]: The current progress of the AVAssetWritingPlanner.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWritingPlanner
type AVAssetWritingPlanner struct {
	objectivec.Object
}

// AVAssetWritingPlannerFromID constructs a [AVAssetWritingPlanner] from an objc.ID.
//
// AVAssetWritingPlanner orchestrates incremental writing of media files.
func AVAssetWritingPlannerFromID(id objc.ID) AVAssetWritingPlanner {
	return AVAssetWritingPlanner{objectivec.Object{ID: id}}
}

// NOTE: AVAssetWritingPlanner adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetWritingPlanner] class.
//
// # Getting progress
//
//   - [IAVAssetWritingPlanner.Progress]: The current progress of the AVAssetWritingPlanner.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWritingPlanner
type IAVAssetWritingPlanner interface {
	objectivec.IObject

	// Topic: Getting progress

	// The current progress of the AVAssetWritingPlanner.
	Progress() IAVAssetWritingPlannerProgress
}

// Init initializes the instance.
func (a AVAssetWritingPlanner) Init() AVAssetWritingPlanner {
	rv := objc.Send[AVAssetWritingPlanner](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetWritingPlanner) Autorelease() AVAssetWritingPlanner {
	rv := objc.Send[AVAssetWritingPlanner](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetWritingPlanner creates a new AVAssetWritingPlanner instance.
func NewAVAssetWritingPlanner() AVAssetWritingPlanner {
	class := getAVAssetWritingPlannerClass()
	rv := objc.Send[AVAssetWritingPlanner](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns segment boundary recommendations for a given source video asset
// track.
//
// videoAssetTrack: The source video AVAssetTrack to be analyzed.
//
// minimumSegmentDuration: The client selected minimum duration for the segments.
//
// minimumSegmentFrameCount: The minimum number of source frames in a segment.
//
// # Return Value
//
// Array of AVPlannedVideoSegmentConfiguration objects, each element
// specifying the configuration of a planned video segment, ordered in output
// PTS order
//
// # Discussion
//
// This is a convenience method that can help clients to pick optimal
// segmentation boundaries for a given source video AVAssetTrack based on the
// structure of the track and the minimumSegmentDuration and
// minimumSegmentFrameCount values provided.
//
// The client needs to ensure that the minimumSegmentDuration is greater than
// or equal to the segment boundary guidelines for the codec type. The client
// should also ensure that minimumSegmentFrameCount also exceeds the segment
// boundary guidelines.
//
// The segments returned will satisfy both the minimumSegmentDuration and
// minimumSegmentFrameCount requirements. The only exception is the very last
// segment, which may be shorter.
//
// The returned array will ensure that segment boundaries occur on sample
// boundaries.
//
// Clients can use these results to fill in the
// AVPlannedVideoSegmentConfiguration for this asset track, if the output
// maintains the source timing. If the output timing differs from the source,
// then the returned AVPlannedVideoSegmentConfiguration array’s results need
// to be modified accordingly by the client.
//
// This method throws NSInvalidArgumentException if minimumSegmentDuration is
// not numeric or is less than or equal to zero, or if
// minimumSegmentFrameCount is less than or equal to 0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWritingPlanner/segmentBoundaryRecommendations(forVideoTrack:minimumSegmentDuration:minimumSegmentFrameCount:)
func (_AVAssetWritingPlannerClass AVAssetWritingPlannerClass) SegmentBoundaryRecommendationsForVideoAVAssetTrackMinimumSegmentDurationMinimumSegmentFrameCount(videoAssetTrack IAVAssetTrack, minimumSegmentDuration coremedia.CMTime, minimumSegmentFrameCount int) []AVPlannedVideoSegmentConfiguration {
	rv := objc.Send[[]objc.ID](objc.ID(_AVAssetWritingPlannerClass.class), objc.Sel("segmentBoundaryRecommendationsForVideoAVAssetTrack:minimumSegmentDuration:minimumSegmentFrameCount:"), videoAssetTrack, minimumSegmentDuration, minimumSegmentFrameCount)
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlannedVideoSegmentConfiguration {
		return AVPlannedVideoSegmentConfigurationFromID(id)
	})
}

// The current progress of the AVAssetWritingPlanner.
//
// # Discussion
//
// Returns an AVAssetWritingPlannerProgress object that can be queried for
// per-track and overall progress information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWritingPlanner/progress
func (a AVAssetWritingPlanner) Progress() IAVAssetWritingPlannerProgress {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("progress"))
	return AVAssetWritingPlannerProgressFromID(objc.ID(rv))
}

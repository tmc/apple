// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetWritingPlannerProgress] class.
var (
	_AVAssetWritingPlannerProgressClass     AVAssetWritingPlannerProgressClass
	_AVAssetWritingPlannerProgressClassOnce sync.Once
)

func getAVAssetWritingPlannerProgressClass() AVAssetWritingPlannerProgressClass {
	_AVAssetWritingPlannerProgressClassOnce.Do(func() {
		_AVAssetWritingPlannerProgressClass = AVAssetWritingPlannerProgressClass{class: objc.GetClass("AVAssetWritingPlannerProgress")}
	})
	return _AVAssetWritingPlannerProgressClass
}

// GetAVAssetWritingPlannerProgressClass returns the class object for AVAssetWritingPlannerProgress.
func GetAVAssetWritingPlannerProgressClass() AVAssetWritingPlannerProgressClass {
	return getAVAssetWritingPlannerProgressClass()
}

type AVAssetWritingPlannerProgressClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetWritingPlannerProgressClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetWritingPlannerProgressClass) Alloc() AVAssetWritingPlannerProgress {
	rv := objc.Send[AVAssetWritingPlannerProgress](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// AVAssetWritingPlannerProgress tracks the progress of incremental writing
// for each track in an AVAssetWritingPlanner session.
//
// # Overview
//
// This class provides per-track progress information as a percentage of the
// total duration completed. Progress can be queried by assemblyTrackID.
//
// # Getting progress
//
//   - [AVAssetWritingPlannerProgress.OverallProgress]: The overall progress across all tracks.
//   - [AVAssetWritingPlannerProgress.ProgressForTrack]: Returns the progress for a specific track identified by its assemblyTrackID.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWritingPlannerProgress
type AVAssetWritingPlannerProgress struct {
	objectivec.Object
}

// AVAssetWritingPlannerProgressFromID constructs a [AVAssetWritingPlannerProgress] from an objc.ID.
//
// AVAssetWritingPlannerProgress tracks the progress of incremental writing
// for each track in an AVAssetWritingPlanner session.
func AVAssetWritingPlannerProgressFromID(id objc.ID) AVAssetWritingPlannerProgress {
	return AVAssetWritingPlannerProgress{objectivec.Object{ID: id}}
}

// NOTE: AVAssetWritingPlannerProgress adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetWritingPlannerProgress] class.
//
// # Getting progress
//
//   - [IAVAssetWritingPlannerProgress.OverallProgress]: The overall progress across all tracks.
//   - [IAVAssetWritingPlannerProgress.ProgressForTrack]: Returns the progress for a specific track identified by its assemblyTrackID.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWritingPlannerProgress
type IAVAssetWritingPlannerProgress interface {
	objectivec.IObject

	// Topic: Getting progress

	// The overall progress across all tracks.
	OverallProgress() float32
	// Returns the progress for a specific track identified by its assemblyTrackID.
	ProgressForTrack(assemblyTrackID coremedia.CMPersistentTrackID) float32
}

// Init initializes the instance.
func (a AVAssetWritingPlannerProgress) Init() AVAssetWritingPlannerProgress {
	rv := objc.Send[AVAssetWritingPlannerProgress](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetWritingPlannerProgress) Autorelease() AVAssetWritingPlannerProgress {
	rv := objc.Send[AVAssetWritingPlannerProgress](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetWritingPlannerProgress creates a new AVAssetWritingPlannerProgress instance.
func NewAVAssetWritingPlannerProgress() AVAssetWritingPlannerProgress {
	class := getAVAssetWritingPlannerProgressClass()
	rv := objc.Send[AVAssetWritingPlannerProgress](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the progress for a specific track identified by its
// assemblyTrackID.
//
// assemblyTrackID: The track ID to query progress for.
//
// # Return Value
//
// A float value between 0.0 and 1.0 representing the percentage of duration
// completed for the track. Returns 0.0 if the track ID is not found.
//
// # Discussion
//
// The progress is calculated as the ratio of completed duration to total
// duration for the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWritingPlannerProgress/progress(forTrack:)
func (a AVAssetWritingPlannerProgress) ProgressForTrack(assemblyTrackID coremedia.CMPersistentTrackID) float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("progressForTrack:"), assemblyTrackID)
	return rv
}

// The overall progress across all tracks.
//
// # Discussion
//
// Returns a float value between 0.0 and 1.0 representing the overall
// progress. This is calculated as the average progress of all tracks weighted
// by their durations.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWritingPlannerProgress/overallProgress
func (a AVAssetWritingPlannerProgress) OverallProgress() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("overallProgress"))
	return rv
}

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemIntegratedTimeline] class.
var (
	_AVPlayerItemIntegratedTimelineClass     AVPlayerItemIntegratedTimelineClass
	_AVPlayerItemIntegratedTimelineClassOnce sync.Once
)

func getAVPlayerItemIntegratedTimelineClass() AVPlayerItemIntegratedTimelineClass {
	_AVPlayerItemIntegratedTimelineClassOnce.Do(func() {
		_AVPlayerItemIntegratedTimelineClass = AVPlayerItemIntegratedTimelineClass{class: objc.GetClass("AVPlayerItemIntegratedTimeline")}
	})
	return _AVPlayerItemIntegratedTimelineClass
}

// GetAVPlayerItemIntegratedTimelineClass returns the class object for AVPlayerItemIntegratedTimeline.
func GetAVPlayerItemIntegratedTimelineClass() AVPlayerItemIntegratedTimelineClass {
	return getAVPlayerItemIntegratedTimelineClass()
}

type AVPlayerItemIntegratedTimelineClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemIntegratedTimelineClass) Alloc() AVPlayerItemIntegratedTimeline {
	rv := objc.Send[AVPlayerItemIntegratedTimeline](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that models the timeline and playback sequence of a primary
// player item and scheduled interstitial events.
//
// # Overview
// 
// The timeline models all regions to traverse during playback. A player may
// not present portions of the primary item when exiting an interstitial event
// with a positive resumption offset.
//
// # Inspecting snapshots
//
//   - [AVPlayerItemIntegratedTimeline.CurrentSnapshot]: An immutable representation of the timeline state at time of request.
//
// # Inspecting the time and date
//
//   - [AVPlayerItemIntegratedTimeline.CurrentTime]: The current time on the integrated timeline.
//   - [AVPlayerItemIntegratedTimeline.CurrentDate]: The current date of playback.
//
// # Seeking
//
//   - [AVPlayerItemIntegratedTimeline.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]: Seeks to a particular time in the integrated time domain.
//   - [AVPlayerItemIntegratedTimeline.SeekToDateCompletionHandler]: Seeks to a particular date in the integrated time domain.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline
type AVPlayerItemIntegratedTimeline struct {
	objectivec.Object
}

// AVPlayerItemIntegratedTimelineFromID constructs a [AVPlayerItemIntegratedTimeline] from an objc.ID.
//
// An object that models the timeline and playback sequence of a primary
// player item and scheduled interstitial events.
func AVPlayerItemIntegratedTimelineFromID(id objc.ID) AVPlayerItemIntegratedTimeline {
	return AVPlayerItemIntegratedTimeline{objectivec.Object{ID: id}}
}
// NOTE: AVPlayerItemIntegratedTimeline adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemIntegratedTimeline] class.
//
// # Inspecting snapshots
//
//   - [IAVPlayerItemIntegratedTimeline.CurrentSnapshot]: An immutable representation of the timeline state at time of request.
//
// # Inspecting the time and date
//
//   - [IAVPlayerItemIntegratedTimeline.CurrentTime]: The current time on the integrated timeline.
//   - [IAVPlayerItemIntegratedTimeline.CurrentDate]: The current date of playback.
//
// # Seeking
//
//   - [IAVPlayerItemIntegratedTimeline.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]: Seeks to a particular time in the integrated time domain.
//   - [IAVPlayerItemIntegratedTimeline.SeekToDateCompletionHandler]: Seeks to a particular date in the integrated time domain.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline
type IAVPlayerItemIntegratedTimeline interface {
	objectivec.IObject

	// Topic: Inspecting snapshots

	// An immutable representation of the timeline state at time of request.
	CurrentSnapshot() IAVPlayerItemIntegratedTimelineSnapshot

	// Topic: Inspecting the time and date

	// The current time on the integrated timeline.
	CurrentTime() objectivec.IObject
	// The current date of playback.
	CurrentDate() foundation.INSDate

	// Topic: Seeking

	// Seeks to a particular time in the integrated time domain.
	SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time objectivec.IObject, toleranceBefore objectivec.IObject, toleranceAfter objectivec.IObject, completionHandler BoolHandler)
	// Seeks to a particular date in the integrated time domain.
	SeekToDateCompletionHandler(date foundation.INSDate, completionHandler BoolHandler)

	// Requests invocation of a block when traversing an offset in a segment during playback.
	AddBoundaryTimeObserverForSegmentOffsetsIntoSegmentQueueUsingBlock(segment IAVPlayerItemSegment, offsetsIntoSegment foundation.INSArray, queue dispatch.Queue, block BoolHandler) AVPlayerItemIntegratedTimelineObserver
	// Requests invocation of a block during playback to report changing time.
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval objectivec.IObject, queue dispatch.Queue, block CMTimeHandler) AVPlayerItemIntegratedTimelineObserver
	// Cancels a previously registered time observer.
	RemoveTimeObserver(observer AVPlayerItemIntegratedTimelineObserver)
}

// Init initializes the instance.
func (p AVPlayerItemIntegratedTimeline) Init() AVPlayerItemIntegratedTimeline {
	rv := objc.Send[AVPlayerItemIntegratedTimeline](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemIntegratedTimeline) Autorelease() AVPlayerItemIntegratedTimeline {
	rv := objc.Send[AVPlayerItemIntegratedTimeline](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemIntegratedTimeline creates a new AVPlayerItemIntegratedTimeline instance.
func NewAVPlayerItemIntegratedTimeline() AVPlayerItemIntegratedTimeline {
	class := getAVPlayerItemIntegratedTimelineClass()
	rv := objc.Send[AVPlayerItemIntegratedTimeline](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Seeks to a particular time in the integrated time domain.
//
// time: A time represented in the integrated time domain.
//
// toleranceBefore: A tolerance before the target time to allow.
//
// toleranceAfter: A tolerance after the target time to allow.
//
// completionHandler: A callback the system invokes after the seek completes. It passes a Boolean
// value of `true` if the playhead moved to the new time.
//
// time is a [coremedia.CMTime].
//
// toleranceBefore is a [coremedia.CMTime].
//
// toleranceAfter is a [coremedia.CMTime].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/seek(to:toleranceBefore:toleranceAfter:completionHandler:)
func (p AVPlayerItemIntegratedTimeline) SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time objectivec.IObject, toleranceBefore objectivec.IObject, toleranceAfter objectivec.IObject, completionHandler BoolHandler) {
_block3, _cleanup3 := NewBoolBlock(completionHandler)
	defer _cleanup3()
	objc.Send[objc.ID](p.ID, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:completionHandler:"), time, toleranceBefore, toleranceAfter, _block3)
}
// Seeks to a particular date in the integrated time domain.
//
// date: A date represented in the integrated time domain.
//
// completionHandler: A callback the system invokes after the seek completes. It passes a Boolean
// value of `true` if the playhead moved to the new date.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/seek(to:completionHandler:)
func (p AVPlayerItemIntegratedTimeline) SeekToDateCompletionHandler(date foundation.INSDate, completionHandler BoolHandler) {
_block1, _cleanup1 := NewBoolBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](p.ID, objc.Sel("seekToDate:completionHandler:"), date, _block1)
}
// Requests invocation of a block when traversing an offset in a segment
// during playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/addBoundaryTimeObserverForSegment:offsetsIntoSegment:queue:usingBlock:
func (p AVPlayerItemIntegratedTimeline) AddBoundaryTimeObserverForSegmentOffsetsIntoSegmentQueueUsingBlock(segment IAVPlayerItemSegment, offsetsIntoSegment foundation.INSArray, queue dispatch.Queue, block BoolHandler) AVPlayerItemIntegratedTimelineObserver {
_block3, _cleanup3 := NewBoolBlock(block)
	defer _cleanup3()
	rv := objc.Send[objc.ID](p.ID, objc.Sel("addBoundaryTimeObserverForSegment:offsetsIntoSegment:queue:usingBlock:"), segment, offsetsIntoSegment, uintptr(queue.Handle()), _block3)
	return AVPlayerItemIntegratedTimelineObserverObjectFromID(rv)
}
// Requests invocation of a block during playback to report changing time.
//
// interval is a [coremedia.CMTime].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/addPeriodicTimeObserverForInterval:queue:usingBlock:
func (p AVPlayerItemIntegratedTimeline) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval objectivec.IObject, queue dispatch.Queue, block CMTimeHandler) AVPlayerItemIntegratedTimelineObserver {
_block2, _cleanup2 := NewCMTimeBlock(block)
	defer _cleanup2()
	rv := objc.Send[objc.ID](p.ID, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, uintptr(queue.Handle()), _block2)
	return AVPlayerItemIntegratedTimelineObserverObjectFromID(rv)
}
// Cancels a previously registered time observer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/removeTimeObserver:
func (p AVPlayerItemIntegratedTimeline) RemoveTimeObserver(observer AVPlayerItemIntegratedTimelineObserver) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeTimeObserver:"), observer)
}

// An immutable representation of the timeline state at time of request.
//
// # Discussion
// 
// A timeline snapshot provides a read-only view of the details of the
// timeline. Because a snapshot provides a fixed view of the timeline at the
// time of the request, its state doesn’t update as playback continues.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/currentSnapshot
func (p AVPlayerItemIntegratedTimeline) CurrentSnapshot() IAVPlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentSnapshot"))
	return AVPlayerItemIntegratedTimelineSnapshotFromID(objc.ID(rv))
}
// The current time on the integrated timeline.
//
// # Discussion
// 
// During playback of interstitial events that occupy a single point, this
// value doesn’t change.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/currentTime
func (p AVPlayerItemIntegratedTimeline) CurrentTime() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentTime"))
	return objectivec.Object{ID: rv}
}
// The current date of playback.
//
// # Discussion
// 
// This value is `nil` if playback doesn’t map to a date.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/currentDate
func (p AVPlayerItemIntegratedTimeline) CurrentDate() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// A notification the system posts when the snapshot objects provided by this
// timeline become out of sync with the current timeline state.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/snapshotsoutofsyncnotification
func (_AVPlayerItemIntegratedTimelineClass AVPlayerItemIntegratedTimelineClass) SnapshotsOutOfSyncNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_AVPlayerItemIntegratedTimelineClass.class), objc.Sel("AVPlayerIntegratedTimelineSnapshotsOutOfSyncNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// SeekToTimeToleranceBeforeToleranceAfter is a synchronous wrapper around [AVPlayerItemIntegratedTimeline.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayerItemIntegratedTimeline) SeekToTimeToleranceBeforeToleranceAfter(ctx context.Context, time objectivec.IObject, toleranceBefore objectivec.IObject, toleranceAfter objectivec.IObject) (bool, error) {
	done := make(chan bool, 1)
	p.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time, toleranceBefore, toleranceAfter, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// SeekToDate is a synchronous wrapper around [AVPlayerItemIntegratedTimeline.SeekToDateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayerItemIntegratedTimeline) SeekToDate(ctx context.Context, date foundation.INSDate) (bool, error) {
	done := make(chan bool, 1)
	p.SeekToDateCompletionHandler(date, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// AddBoundaryTimeObserverForSegmentOffsetsIntoSegmentQueueUsingBlockSync is a synchronous wrapper around [AVPlayerItemIntegratedTimeline.AddBoundaryTimeObserverForSegmentOffsetsIntoSegmentQueueUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayerItemIntegratedTimeline) AddBoundaryTimeObserverForSegmentOffsetsIntoSegmentQueueUsingBlockSync(ctx context.Context, segment IAVPlayerItemSegment, offsetsIntoSegment foundation.INSArray, queue dispatch.Queue) (bool, error) {
	done := make(chan bool, 1)
	p.AddBoundaryTimeObserverForSegmentOffsetsIntoSegmentQueueUsingBlock(segment, offsetsIntoSegment, queue, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// AddPeriodicTimeObserverForIntervalQueueUsingBlockSync is a synchronous wrapper around [AVPlayerItemIntegratedTimeline.AddPeriodicTimeObserverForIntervalQueueUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayerItemIntegratedTimeline) AddPeriodicTimeObserverForIntervalQueueUsingBlockSync(ctx context.Context, interval objectivec.IObject, queue dispatch.Queue) (objectivec.IObject, error) {
	done := make(chan objectivec.IObject, 1)
	p.AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval, queue, func(val objectivec.IObject) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}


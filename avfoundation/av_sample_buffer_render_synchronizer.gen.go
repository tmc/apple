// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSampleBufferRenderSynchronizer] class.
var (
	_AVSampleBufferRenderSynchronizerClass     AVSampleBufferRenderSynchronizerClass
	_AVSampleBufferRenderSynchronizerClassOnce sync.Once
)

func getAVSampleBufferRenderSynchronizerClass() AVSampleBufferRenderSynchronizerClass {
	_AVSampleBufferRenderSynchronizerClassOnce.Do(func() {
		_AVSampleBufferRenderSynchronizerClass = AVSampleBufferRenderSynchronizerClass{class: objc.GetClass("AVSampleBufferRenderSynchronizer")}
	})
	return _AVSampleBufferRenderSynchronizerClass
}

// GetAVSampleBufferRenderSynchronizerClass returns the class object for AVSampleBufferRenderSynchronizer.
func GetAVSampleBufferRenderSynchronizerClass() AVSampleBufferRenderSynchronizerClass {
	return getAVSampleBufferRenderSynchronizerClass()
}

type AVSampleBufferRenderSynchronizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSampleBufferRenderSynchronizerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSampleBufferRenderSynchronizerClass) Alloc() AVSampleBufferRenderSynchronizer {
	rv := objc.Send[AVSampleBufferRenderSynchronizer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object used to synchronize multiple queued sample buffers to a single
// timeline.
//
// # Overview
//
// This class synchronizes multiple objects that conform to
// [AVQueuedSampleBufferRendering] to a single timeline.
//
// # Managing renderers
//
//   - [AVSampleBufferRenderSynchronizer.Renderers]: An array of queued sample buffer renderers currently attached to the synchronizer.
//   - [AVSampleBufferRenderSynchronizer.AddRenderer]: Adds a renderer to the list of renderers under the synchronizer’s control.
//   - [AVSampleBufferRenderSynchronizer.RemoveRendererAtTimeCompletionHandler]: Removes a renderer from the synchronizer.
//
// # Accessing time information
//
//   - [AVSampleBufferRenderSynchronizer.CurrentTime]: Returns the current time of the synchronizer.
//   - [AVSampleBufferRenderSynchronizer.Timebase]: The synchronizer’s rendering timebase which determines how it interprets timestamps.
//   - [AVSampleBufferRenderSynchronizer.Rate]: The current playback rate.
//   - [AVSampleBufferRenderSynchronizer.SetRate]
//   - [AVSampleBufferRenderSynchronizer.SetRateTime]: Sets the renderer’s time and rate.
//   - [AVSampleBufferRenderSynchronizer.SetRateTimeAtHostTime]: Sets the playback rate and the relationship between the current time and host time.
//   - [AVSampleBufferRenderSynchronizer.DelaysRateChangeUntilHasSufficientMediaData]: A Boolean value that Indicates whether the playback should start immediately on rate change requests.
//   - [AVSampleBufferRenderSynchronizer.SetDelaysRateChangeUntilHasSufficientMediaData]
//
// # Observing time
//
//   - [AVSampleBufferRenderSynchronizer.AddPeriodicTimeObserverForIntervalQueueUsingBlock]: Requests invocation of a block during rendering at specified time intervals.
//   - [AVSampleBufferRenderSynchronizer.AddBoundaryTimeObserverForTimesQueueUsingBlock]: Requests invocation of a block when specified times are traversed during normal rendering.
//   - [AVSampleBufferRenderSynchronizer.RemoveTimeObserver]: Cancels the specified time observer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer
type AVSampleBufferRenderSynchronizer struct {
	objectivec.Object
}

// AVSampleBufferRenderSynchronizerFromID constructs a [AVSampleBufferRenderSynchronizer] from an objc.ID.
//
// An object used to synchronize multiple queued sample buffers to a single
// timeline.
func AVSampleBufferRenderSynchronizerFromID(id objc.ID) AVSampleBufferRenderSynchronizer {
	return AVSampleBufferRenderSynchronizer{objectivec.Object{ID: id}}
}

// NOTE: AVSampleBufferRenderSynchronizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSampleBufferRenderSynchronizer] class.
//
// # Managing renderers
//
//   - [IAVSampleBufferRenderSynchronizer.Renderers]: An array of queued sample buffer renderers currently attached to the synchronizer.
//   - [IAVSampleBufferRenderSynchronizer.AddRenderer]: Adds a renderer to the list of renderers under the synchronizer’s control.
//   - [IAVSampleBufferRenderSynchronizer.RemoveRendererAtTimeCompletionHandler]: Removes a renderer from the synchronizer.
//
// # Accessing time information
//
//   - [IAVSampleBufferRenderSynchronizer.CurrentTime]: Returns the current time of the synchronizer.
//   - [IAVSampleBufferRenderSynchronizer.Timebase]: The synchronizer’s rendering timebase which determines how it interprets timestamps.
//   - [IAVSampleBufferRenderSynchronizer.Rate]: The current playback rate.
//   - [IAVSampleBufferRenderSynchronizer.SetRate]
//   - [IAVSampleBufferRenderSynchronizer.SetRateTime]: Sets the renderer’s time and rate.
//   - [IAVSampleBufferRenderSynchronizer.SetRateTimeAtHostTime]: Sets the playback rate and the relationship between the current time and host time.
//   - [IAVSampleBufferRenderSynchronizer.DelaysRateChangeUntilHasSufficientMediaData]: A Boolean value that Indicates whether the playback should start immediately on rate change requests.
//   - [IAVSampleBufferRenderSynchronizer.SetDelaysRateChangeUntilHasSufficientMediaData]
//
// # Observing time
//
//   - [IAVSampleBufferRenderSynchronizer.AddPeriodicTimeObserverForIntervalQueueUsingBlock]: Requests invocation of a block during rendering at specified time intervals.
//   - [IAVSampleBufferRenderSynchronizer.AddBoundaryTimeObserverForTimesQueueUsingBlock]: Requests invocation of a block when specified times are traversed during normal rendering.
//   - [IAVSampleBufferRenderSynchronizer.RemoveTimeObserver]: Cancels the specified time observer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer
type IAVSampleBufferRenderSynchronizer interface {
	objectivec.IObject

	// Topic: Managing renderers

	// An array of queued sample buffer renderers currently attached to the synchronizer.
	Renderers() []objectivec.IObject
	// Adds a renderer to the list of renderers under the synchronizer’s control.
	AddRenderer(renderer AVQueuedSampleBufferRendering)
	// Removes a renderer from the synchronizer.
	RemoveRendererAtTimeCompletionHandler(renderer AVQueuedSampleBufferRendering, time coremedia.CMTime, completionHandler BoolHandler)

	// Topic: Accessing time information

	// Returns the current time of the synchronizer.
	CurrentTime() coremedia.CMTime
	// The synchronizer’s rendering timebase which determines how it interprets timestamps.
	Timebase() uintptr
	// The current playback rate.
	Rate() float32
	SetRate(value float32)
	// Sets the renderer’s time and rate.
	SetRateTime(rate float32, time coremedia.CMTime)
	// Sets the playback rate and the relationship between the current time and host time.
	SetRateTimeAtHostTime(rate float32, time coremedia.CMTime, hostTime coremedia.CMTime)
	// A Boolean value that Indicates whether the playback should start immediately on rate change requests.
	DelaysRateChangeUntilHasSufficientMediaData() bool
	SetDelaysRateChangeUntilHasSufficientMediaData(value bool)

	// Topic: Observing time

	// Requests invocation of a block during rendering at specified time intervals.
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval coremedia.CMTime, queue dispatch.Queue, block CMTimeHandler) objectivec.IObject
	// Requests invocation of a block when specified times are traversed during normal rendering.
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.NSValue, queue dispatch.Queue, block VoidHandler) objectivec.IObject
	// Cancels the specified time observer.
	RemoveTimeObserver(observer objectivec.IObject)
}

// Init initializes the instance.
func (s AVSampleBufferRenderSynchronizer) Init() AVSampleBufferRenderSynchronizer {
	rv := objc.Send[AVSampleBufferRenderSynchronizer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSampleBufferRenderSynchronizer) Autorelease() AVSampleBufferRenderSynchronizer {
	rv := objc.Send[AVSampleBufferRenderSynchronizer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferRenderSynchronizer creates a new AVSampleBufferRenderSynchronizer instance.
func NewAVSampleBufferRenderSynchronizer() AVSampleBufferRenderSynchronizer {
	class := getAVSampleBufferRenderSynchronizerClass()
	rv := objc.Send[AVSampleBufferRenderSynchronizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds a renderer to the list of renderers under the synchronizer’s
// control.
//
// renderer: The render to be added.
//
// # Discussion
//
// This method can be called while [Rate] is not `0.0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addRenderer(_:)
func (s AVSampleBufferRenderSynchronizer) AddRenderer(renderer AVQueuedSampleBufferRendering) {
	objc.Send[objc.ID](s.ID, objc.Sel("addRenderer:"), renderer)
}

// Removes a renderer from the synchronizer.
//
// time: The time on the timebase’s timeline at which the renderer should be
// removed. If the time is in the past, the renderer is immediately removed.
//
// completionHandler: An optional block to invoke when the renderer is removed from the
// synchronizer. The block takes one argument:
//
// didRemoveRenderer: A Boolean value indicating the whether the renderer was
// removed.
//
// # Discussion
//
// This method removes the renderer asynchronously. The method can be called
// more than once, with a subsequent scheduled removal replacing a previously
// scheduled removal. This method can be called while [Rate] is not `0.0`.
//
// Clients may provide an optional `completionHandler` to be notified when the
// scheduled removal is complete. If provided, the completion handler will
// always be called with the following values for `didRemoveRenderer`:
//
// - If the renderer has not been added to this synchronizer,
// `didRemoveRenderer` is [NO]. - If the removal of a particular renderer is
// scheduled after the same renderer’s removal was previous scheduled but
// not yet occurred, the previously scheduled removal’s completion handler
// is and `didRemoveRenderer` set to [NO]. - When the renderer is removed due
// to a scheduled removal, the completion handler is called and
// `didRemoveRenderer` set to YES.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/removeRenderer(_:at:completionHandler:)
func (s AVSampleBufferRenderSynchronizer) RemoveRendererAtTimeCompletionHandler(renderer AVQueuedSampleBufferRendering, time coremedia.CMTime, completionHandler BoolHandler) {
	_block2, _ := NewBoolBlock(completionHandler)
	objc.Send[objc.ID](s.ID, objc.Sel("removeRenderer:atTime:completionHandler:"), renderer, time, _block2)
}

// Returns the current time of the synchronizer.
//
// # Return Value
//
// A [CMTime] object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/currentTime()
//
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
func (s AVSampleBufferRenderSynchronizer) CurrentTime() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](s.ID, objc.Sel("currentTime"))
	return coremedia.CMTime(rv)
}

// Sets the renderer’s time and rate.
//
// rate: The new timebase rate. This value must be greater than or equal to `0.0`.
//
// time: The new timebase time. This value must be greater than or equal to [zero],
// or [invalid].
//
// # Discussion
//
// This method first sets the new time and then the new rendering rate. A
// `rate` value of `0.0` means that playback has stopped while a `rate` value
// of `1.0` indicates playback should be at the natural rate of the media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/setRate(_:time:)
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
func (s AVSampleBufferRenderSynchronizer) SetRateTime(rate float32, time coremedia.CMTime) {
	objc.Send[objc.ID](s.ID, objc.Sel("setRate:time:"), rate, time)
}

// Sets the playback rate and the relationship between the current time and
// host time.
//
// rate: A new timebase rate. This value must be greater than or equal to `0.0`.
//
// time: A new timebase time. This value must be greater than or equal to [zero], or
// [invalid].
//
// hostTime: A new host time. This value must be greater than or equal to [zero], or
// [invalid].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/setRate(_:time:atHostTime:)
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
func (s AVSampleBufferRenderSynchronizer) SetRateTimeAtHostTime(rate float32, time coremedia.CMTime, hostTime coremedia.CMTime) {
	objc.Send[objc.ID](s.ID, objc.Sel("setRate:time:atHostTime:"), rate, time, hostTime)
}

// Requests invocation of a block during rendering at specified time
// intervals.
//
// interval: The specified time interval requesting block invocation during rendering.
//
// queue: The serial queue the block should be unqueued on. If you pass [NULL], the
// main queue is used. Passing a concurrent queue results in undefined
// behavior.
//
// block: The block to be invoked periodically.
//
// # Return Value
//
// An object that conforms to [NSObject]. You must retain this value as long
// as you want the time observer to be invoked by the synchronizer. Pass this
// object to [RemoveTimeObserver] to cancel time observation.
//
// # Discussion
//
// The block associated with this method is invoked at the specified time
// intervals, interpreted according to the timeline of the timebase. The block
// is also invoked whenever there is a time jump or rendering starts or stops.
//
// If a very short time interval is used, the synchronizer may invoke the
// block less frequently than requested. However, the synchronizer will invoke
// the block often enough for the client to update indications of the current
// time appropriately in its end-user interface.
//
// Always pair a call to this method with a call to [RemoveTimeObserver].
// Releasing the observer without calling `removeTimeObserver(_:)` results in
// undefined behavior.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addPeriodicTimeObserver(forInterval:queue:using:)
//
// [NSObject]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class
func (s AVSampleBufferRenderSynchronizer) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval coremedia.CMTime, queue dispatch.Queue, block CMTimeHandler) objectivec.IObject {
	_block2, _ := NewCMTimeBlock(block)
	rv := objc.Send[objc.ID](s.ID, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, uintptr(queue.Handle()), _block2)
	return objectivec.Object{ID: rv}
}

// Requests invocation of a block when specified times are traversed during
// normal rendering.
//
// times: An array containing the times for which the observer requests notification.
//
// queue: The serial queue the block should be unqueued on. If you pass [NULL], the
// main queue is used. Passing a concurrent queue results in undefined
// behavior.
//
// block: The block to be invoked when any of the specified times is crossed during
// normal rendering.
//
// # Return Value
//
// An object that conforms to [NSObject]. You must retain this value as long
// as you want the time observer to be invoked by the synchronizer. Pass this
// object to [RemoveTimeObserver] to cancel time observation.
//
// # Discussion
//
// Always pair a call to this method with a call to [RemoveTimeObserver].
// Releasing the observer without calling `removeTimeObserver(_:)` results in
// undefined behavior.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addBoundaryTimeObserver(forTimes:queue:using:)
//
// [NSObject]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class
func (s AVSampleBufferRenderSynchronizer) AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.NSValue, queue dispatch.Queue, block VoidHandler) objectivec.IObject {
	_block2, _ := NewVoidBlock(block)
	rv := objc.Send[objc.ID](s.ID, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, uintptr(queue.Handle()), _block2)
	return objectivec.Object{ID: rv}
}

// Cancels the specified time observer.
//
// observer: The time observer to be cancelled.
//
// # Discussion
//
// Use this method to explicitly cancel time observers added using
// [AddPeriodicTimeObserverForIntervalQueueUsingBlock] or
// [AddBoundaryTimeObserverForTimesQueueUsingBlock]
//
// Upon return, the caller is guaranteed that no new time observer blocks will
// begin executing. Depending on the calling thread and the queue used to add
// the time observer, an in-flight block may continue to execute after this
// method returns. You can guarantee synchronous time observer removal by
// enqueuing the call to “ on that queue. Call [sync(execute:)] after “ to
// wait for any in-flight blocks to finish executing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/removeTimeObserver(_:)
//
// [sync(execute:)]: https://developer.apple.com/documentation/Dispatch/DispatchQueue/sync(execute:)-3segw
func (s AVSampleBufferRenderSynchronizer) RemoveTimeObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeTimeObserver:"), observer)
}

// An array of queued sample buffer renderers currently attached to the
// synchronizer.
//
// # Discussion
//
// This property includes all renderers that have been added to the
// synchronizer and haven’t been removed, including renderers that have been
// scheduled for removal, but have yet to be removed. This property is not KVO
// observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/renderers
func (s AVSampleBufferRenderSynchronizer) Renderers() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("renderers"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// The synchronizer’s rendering timebase which determines how it interprets
// timestamps.
//
// # Discussion
//
// The default for this property is the clock for an added
// [AVSampleBufferAudioRenderer] object. If you haven’t added a renderer,
// the timebase is the system host clock.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/timebase
func (s AVSampleBufferRenderSynchronizer) Timebase() uintptr {
	rv := objc.Send[uintptr](s.ID, objc.Sel("timebase"))
	return rv
}

// The current playback rate.
//
// # Discussion
//
// A value of `0.0` means playback has stopped. A value of `1.0` tells the
// renderer to play at the natural rate of the media. This property must be
// greater than or equal to `0.0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/rate
func (s AVSampleBufferRenderSynchronizer) Rate() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("rate"))
	return rv
}
func (s AVSampleBufferRenderSynchronizer) SetRate(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setRate:"), value)
}

// A Boolean value that Indicates whether the playback should start
// immediately on rate change requests.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/delaysRateChangeUntilHasSufficientMediaData
func (s AVSampleBufferRenderSynchronizer) DelaysRateChangeUntilHasSufficientMediaData() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("delaysRateChangeUntilHasSufficientMediaData"))
	return rv
}
func (s AVSampleBufferRenderSynchronizer) SetDelaysRateChangeUntilHasSufficientMediaData(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelaysRateChangeUntilHasSufficientMediaData:"), value)
}

// RemoveRendererAtTime is a synchronous wrapper around [AVSampleBufferRenderSynchronizer.RemoveRendererAtTimeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s AVSampleBufferRenderSynchronizer) RemoveRendererAtTime(ctx context.Context, renderer AVQueuedSampleBufferRendering, time coremedia.CMTime) (bool, error) {
	done := make(chan bool, 1)
	s.RemoveRendererAtTimeCompletionHandler(renderer, time, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// AddPeriodicTimeObserverForIntervalQueueUsingBlockSync is a synchronous wrapper around [AVSampleBufferRenderSynchronizer.AddPeriodicTimeObserverForIntervalQueueUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s AVSampleBufferRenderSynchronizer) AddPeriodicTimeObserverForIntervalQueueUsingBlockSync(ctx context.Context, interval coremedia.CMTime, queue dispatch.Queue) (coremedia.CMTime, error) {
	done := make(chan coremedia.CMTime, 1)
	s.AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval, queue, func(val coremedia.CMTime) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return coremedia.CMTime{}, ctx.Err()
	}
}

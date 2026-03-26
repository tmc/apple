// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetricDownloadSummaryEvent] class.
var (
	_AVMetricDownloadSummaryEventClass     AVMetricDownloadSummaryEventClass
	_AVMetricDownloadSummaryEventClassOnce sync.Once
)

func getAVMetricDownloadSummaryEventClass() AVMetricDownloadSummaryEventClass {
	_AVMetricDownloadSummaryEventClassOnce.Do(func() {
		_AVMetricDownloadSummaryEventClass = AVMetricDownloadSummaryEventClass{class: objc.GetClass("AVMetricDownloadSummaryEvent")}
	})
	return _AVMetricDownloadSummaryEventClass
}

// GetAVMetricDownloadSummaryEventClass returns the class object for AVMetricDownloadSummaryEvent.
func GetAVMetricDownloadSummaryEventClass() AVMetricDownloadSummaryEventClass {
	return getAVMetricDownloadSummaryEventClass()
}

type AVMetricDownloadSummaryEventClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricDownloadSummaryEventClass) Alloc() AVMetricDownloadSummaryEvent {
	rv := objc.Send[AVMetricDownloadSummaryEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Represents a summary metric event with aggregated metrics for the entire
// download task.
//
// # Overview
// 
// Subclasses of this type that are used from Swift must fulfill the
// requirements of a Sendable type.
//
// # Instance Properties
//
//   - [AVMetricDownloadSummaryEvent.BytesDownloadedCount]: Returns the total number of bytes downloaded by the download task.
//   - [AVMetricDownloadSummaryEvent.DownloadDuration]: Returns the total duration of the download in seconds.
//   - [AVMetricDownloadSummaryEvent.ErrorEvent]: Returns the error event if any. If no value is available, returns nil.
//   - [AVMetricDownloadSummaryEvent.MediaResourceRequestCount]: Returns the total number of media requests performed by the download task. This includes playlist requests, media segment requests, and content key requests.
//   - [AVMetricDownloadSummaryEvent.RecoverableErrorCount]: Returns the total count of recoverable errors encountered during the download. If no errors were encountered, returns 0.
//   - [AVMetricDownloadSummaryEvent.Variants]: Returns the variants that were downloaded.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent
type AVMetricDownloadSummaryEvent struct {
	AVMetricEvent
}

// AVMetricDownloadSummaryEventFromID constructs a [AVMetricDownloadSummaryEvent] from an objc.ID.
//
// Represents a summary metric event with aggregated metrics for the entire
// download task.
func AVMetricDownloadSummaryEventFromID(id objc.ID) AVMetricDownloadSummaryEvent {
	return AVMetricDownloadSummaryEvent{AVMetricEvent: AVMetricEventFromID(id)}
}
// NOTE: AVMetricDownloadSummaryEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricDownloadSummaryEvent] class.
//
// # Instance Properties
//
//   - [IAVMetricDownloadSummaryEvent.BytesDownloadedCount]: Returns the total number of bytes downloaded by the download task.
//   - [IAVMetricDownloadSummaryEvent.DownloadDuration]: Returns the total duration of the download in seconds.
//   - [IAVMetricDownloadSummaryEvent.ErrorEvent]: Returns the error event if any. If no value is available, returns nil.
//   - [IAVMetricDownloadSummaryEvent.MediaResourceRequestCount]: Returns the total number of media requests performed by the download task. This includes playlist requests, media segment requests, and content key requests.
//   - [IAVMetricDownloadSummaryEvent.RecoverableErrorCount]: Returns the total count of recoverable errors encountered during the download. If no errors were encountered, returns 0.
//   - [IAVMetricDownloadSummaryEvent.Variants]: Returns the variants that were downloaded.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent
type IAVMetricDownloadSummaryEvent interface {
	IAVMetricEvent

	// Topic: Instance Properties

	// Returns the total number of bytes downloaded by the download task.
	BytesDownloadedCount() int
	// Returns the total duration of the download in seconds.
	DownloadDuration() float64
	// Returns the error event if any. If no value is available, returns nil.
	ErrorEvent() IAVMetricErrorEvent
	// Returns the total number of media requests performed by the download task. This includes playlist requests, media segment requests, and content key requests.
	MediaResourceRequestCount() int
	// Returns the total count of recoverable errors encountered during the download. If no errors were encountered, returns 0.
	RecoverableErrorCount() int
	// Returns the variants that were downloaded.
	Variants() []AVAssetVariant
}

// Init initializes the instance.
func (m AVMetricDownloadSummaryEvent) Init() AVMetricDownloadSummaryEvent {
	rv := objc.Send[AVMetricDownloadSummaryEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricDownloadSummaryEvent) Autorelease() AVMetricDownloadSummaryEvent {
	rv := objc.Send[AVMetricDownloadSummaryEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricDownloadSummaryEvent creates a new AVMetricDownloadSummaryEvent instance.
func NewAVMetricDownloadSummaryEvent() AVMetricDownloadSummaryEvent {
	class := getAVMetricDownloadSummaryEventClass()
	rv := objc.Send[AVMetricDownloadSummaryEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the total number of bytes downloaded by the download task.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/bytesDownloadedCount
func (m AVMetricDownloadSummaryEvent) BytesDownloadedCount() int {
	rv := objc.Send[int](m.ID, objc.Sel("bytesDownloadedCount"))
	return rv
}
// Returns the total duration of the download in seconds.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/downloadDuration
func (m AVMetricDownloadSummaryEvent) DownloadDuration() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("downloadDuration"))
	return rv
}
// Returns the error event if any. If no value is available, returns nil.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/errorEvent
func (m AVMetricDownloadSummaryEvent) ErrorEvent() IAVMetricErrorEvent {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("errorEvent"))
	return AVMetricErrorEventFromID(objc.ID(rv))
}
// Returns the total number of media requests performed by the download task.
// This includes playlist requests, media segment requests, and content key
// requests.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/mediaResourceRequestCount
func (m AVMetricDownloadSummaryEvent) MediaResourceRequestCount() int {
	rv := objc.Send[int](m.ID, objc.Sel("mediaResourceRequestCount"))
	return rv
}
// Returns the total count of recoverable errors encountered during the
// download. If no errors were encountered, returns 0.
//
// # Discussion
// 
// Error counts may not be consistent across OS versions. Comparisons should
// be made within a given OS version, as error reporting is subject to change
// with OS updates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/recoverableErrorCount
func (m AVMetricDownloadSummaryEvent) RecoverableErrorCount() int {
	rv := objc.Send[int](m.ID, objc.Sel("recoverableErrorCount"))
	return rv
}
// Returns the variants that were downloaded.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricDownloadSummaryEvent/variants
func (m AVMetricDownloadSummaryEvent) Variants() []AVAssetVariant {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("variants"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetVariant {
		return AVAssetVariantFromID(id)
	})
}


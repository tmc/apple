// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemAccessLogEvent] class.
var (
	_AVPlayerItemAccessLogEventClass     AVPlayerItemAccessLogEventClass
	_AVPlayerItemAccessLogEventClassOnce sync.Once
)

func getAVPlayerItemAccessLogEventClass() AVPlayerItemAccessLogEventClass {
	_AVPlayerItemAccessLogEventClassOnce.Do(func() {
		_AVPlayerItemAccessLogEventClass = AVPlayerItemAccessLogEventClass{class: objc.GetClass("AVPlayerItemAccessLogEvent")}
	})
	return _AVPlayerItemAccessLogEventClass
}

// GetAVPlayerItemAccessLogEventClass returns the class object for AVPlayerItemAccessLogEvent.
func GetAVPlayerItemAccessLogEventClass() AVPlayerItemAccessLogEventClass {
	return getAVPlayerItemAccessLogEventClass()
}

type AVPlayerItemAccessLogEventClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemAccessLogEventClass) Alloc() AVPlayerItemAccessLogEvent {
	rv := objc.Send[AVPlayerItemAccessLogEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A single entry in a player item’s access log.
//
// # Overview
// 
// This object provides named properties for accessing the data fields of each
// log event. Each event is a single entry in an [AVPlayerItem] object’s
// access log.
// 
// These properties aren’t observable. For more information about key-value
// observing, see [Using Key-Value Observing in Swift].
//
// [Using Key-Value Observing in Swift]: https://developer.apple.com/documentation/Swift/using-key-value-observing-in-swift
//
// # Getting server-related log events
//
//   - [AVPlayerItemAccessLogEvent.URI]: The URI of the playback item.
//   - [AVPlayerItemAccessLogEvent.ServerAddress]: The IP address of the server that was the source of the last delivered media segment.
//   - [AVPlayerItemAccessLogEvent.NumberOfServerAddressChanges]: A count of changes to the server address over the last uninterrupted period of playback.
//   - [AVPlayerItemAccessLogEvent.MediaRequestsWWAN]: The number of network read requests over a WWAN.
//   - [AVPlayerItemAccessLogEvent.TransferDuration]: The accumulated duration, in seconds, of active network transfer of bytes.
//   - [AVPlayerItemAccessLogEvent.NumberOfBytesTransferred]: The accumulated number of bytes transferred by the item.
//   - [AVPlayerItemAccessLogEvent.NumberOfMediaRequests]: The number of media read requests from the server to this client.
//
// # Getting playback-related log events
//
//   - [AVPlayerItemAccessLogEvent.PlaybackStartDate]: The date and time at which playback began for this event.
//   - [AVPlayerItemAccessLogEvent.PlaybackSessionID]: A GUID that identifies the playback session.
//   - [AVPlayerItemAccessLogEvent.PlaybackStartOffset]: The offset, in seconds, in the playlist where the last uninterrupted period of playback began.
//   - [AVPlayerItemAccessLogEvent.PlaybackType]: The playback type.
//   - [AVPlayerItemAccessLogEvent.StartupTime]: The accumulated duration, in seconds, until the player item is ready to play.
//   - [AVPlayerItemAccessLogEvent.DurationWatched]: The accumulated duration, in seconds, of the media played.
//   - [AVPlayerItemAccessLogEvent.NumberOfDroppedVideoFrames]: The total number of dropped video frames
//   - [AVPlayerItemAccessLogEvent.NumberOfStalls]: The total number of playback stalls encountered.
//   - [AVPlayerItemAccessLogEvent.SegmentsDownloadedDuration]: The accumulated duration, in seconds, of the media segments downloaded.
//   - [AVPlayerItemAccessLogEvent.DownloadOverdue]: The total number of times that downloading the segments took too long.
//
// # Getting bit rate log events
//
//   - [AVPlayerItemAccessLogEvent.ObservedBitrateStandardDeviation]: The standard deviation of the observed segment download bit rates.
//   - [AVPlayerItemAccessLogEvent.SwitchBitrate]: The bandwidth value that causes a switch, up or down, in the item’s quality being played.
//   - [AVPlayerItemAccessLogEvent.IndicatedBitrate]: The throughput, in bits per second, required to play the stream, as advertised by the server.
//   - [AVPlayerItemAccessLogEvent.ObservedBitrate]: The empirical throughput, in bits per second, across all media downloaded.
//   - [AVPlayerItemAccessLogEvent.AverageAudioBitrate]: The audio track’s average bit rate, in bits per second.
//   - [AVPlayerItemAccessLogEvent.AverageVideoBitrate]: The video track’s average bit rate, in bits per second.
//   - [AVPlayerItemAccessLogEvent.IndicatedAverageBitrate]: The average throughput, in bits per second, required to play the stream, as advertised by the server.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent
type AVPlayerItemAccessLogEvent struct {
	objectivec.Object
}

// AVPlayerItemAccessLogEventFromID constructs a [AVPlayerItemAccessLogEvent] from an objc.ID.
//
// A single entry in a player item’s access log.
func AVPlayerItemAccessLogEventFromID(id objc.ID) AVPlayerItemAccessLogEvent {
	return AVPlayerItemAccessLogEvent{objectivec.Object{ID: id}}
}
// NOTE: AVPlayerItemAccessLogEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemAccessLogEvent] class.
//
// # Getting server-related log events
//
//   - [IAVPlayerItemAccessLogEvent.URI]: The URI of the playback item.
//   - [IAVPlayerItemAccessLogEvent.ServerAddress]: The IP address of the server that was the source of the last delivered media segment.
//   - [IAVPlayerItemAccessLogEvent.NumberOfServerAddressChanges]: A count of changes to the server address over the last uninterrupted period of playback.
//   - [IAVPlayerItemAccessLogEvent.MediaRequestsWWAN]: The number of network read requests over a WWAN.
//   - [IAVPlayerItemAccessLogEvent.TransferDuration]: The accumulated duration, in seconds, of active network transfer of bytes.
//   - [IAVPlayerItemAccessLogEvent.NumberOfBytesTransferred]: The accumulated number of bytes transferred by the item.
//   - [IAVPlayerItemAccessLogEvent.NumberOfMediaRequests]: The number of media read requests from the server to this client.
//
// # Getting playback-related log events
//
//   - [IAVPlayerItemAccessLogEvent.PlaybackStartDate]: The date and time at which playback began for this event.
//   - [IAVPlayerItemAccessLogEvent.PlaybackSessionID]: A GUID that identifies the playback session.
//   - [IAVPlayerItemAccessLogEvent.PlaybackStartOffset]: The offset, in seconds, in the playlist where the last uninterrupted period of playback began.
//   - [IAVPlayerItemAccessLogEvent.PlaybackType]: The playback type.
//   - [IAVPlayerItemAccessLogEvent.StartupTime]: The accumulated duration, in seconds, until the player item is ready to play.
//   - [IAVPlayerItemAccessLogEvent.DurationWatched]: The accumulated duration, in seconds, of the media played.
//   - [IAVPlayerItemAccessLogEvent.NumberOfDroppedVideoFrames]: The total number of dropped video frames
//   - [IAVPlayerItemAccessLogEvent.NumberOfStalls]: The total number of playback stalls encountered.
//   - [IAVPlayerItemAccessLogEvent.SegmentsDownloadedDuration]: The accumulated duration, in seconds, of the media segments downloaded.
//   - [IAVPlayerItemAccessLogEvent.DownloadOverdue]: The total number of times that downloading the segments took too long.
//
// # Getting bit rate log events
//
//   - [IAVPlayerItemAccessLogEvent.ObservedBitrateStandardDeviation]: The standard deviation of the observed segment download bit rates.
//   - [IAVPlayerItemAccessLogEvent.SwitchBitrate]: The bandwidth value that causes a switch, up or down, in the item’s quality being played.
//   - [IAVPlayerItemAccessLogEvent.IndicatedBitrate]: The throughput, in bits per second, required to play the stream, as advertised by the server.
//   - [IAVPlayerItemAccessLogEvent.ObservedBitrate]: The empirical throughput, in bits per second, across all media downloaded.
//   - [IAVPlayerItemAccessLogEvent.AverageAudioBitrate]: The audio track’s average bit rate, in bits per second.
//   - [IAVPlayerItemAccessLogEvent.AverageVideoBitrate]: The video track’s average bit rate, in bits per second.
//   - [IAVPlayerItemAccessLogEvent.IndicatedAverageBitrate]: The average throughput, in bits per second, required to play the stream, as advertised by the server.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent
type IAVPlayerItemAccessLogEvent interface {
	objectivec.IObject

	// Topic: Getting server-related log events

	// The URI of the playback item.
	URI() string
	// The IP address of the server that was the source of the last delivered media segment.
	ServerAddress() string
	// A count of changes to the server address over the last uninterrupted period of playback.
	NumberOfServerAddressChanges() int
	// The number of network read requests over a WWAN.
	MediaRequestsWWAN() int
	// The accumulated duration, in seconds, of active network transfer of bytes.
	TransferDuration() float64
	// The accumulated number of bytes transferred by the item.
	NumberOfBytesTransferred() int64
	// The number of media read requests from the server to this client.
	NumberOfMediaRequests() int

	// Topic: Getting playback-related log events

	// The date and time at which playback began for this event.
	PlaybackStartDate() foundation.INSDate
	// A GUID that identifies the playback session.
	PlaybackSessionID() string
	// The offset, in seconds, in the playlist where the last uninterrupted period of playback began.
	PlaybackStartOffset() float64
	// The playback type.
	PlaybackType() string
	// The accumulated duration, in seconds, until the player item is ready to play.
	StartupTime() float64
	// The accumulated duration, in seconds, of the media played.
	DurationWatched() float64
	// The total number of dropped video frames
	NumberOfDroppedVideoFrames() int
	// The total number of playback stalls encountered.
	NumberOfStalls() int
	// The accumulated duration, in seconds, of the media segments downloaded.
	SegmentsDownloadedDuration() float64
	// The total number of times that downloading the segments took too long.
	DownloadOverdue() int

	// Topic: Getting bit rate log events

	// The standard deviation of the observed segment download bit rates.
	ObservedBitrateStandardDeviation() float64
	// The bandwidth value that causes a switch, up or down, in the item’s quality being played.
	SwitchBitrate() float64
	// The throughput, in bits per second, required to play the stream, as advertised by the server.
	IndicatedBitrate() float64
	// The empirical throughput, in bits per second, across all media downloaded.
	ObservedBitrate() float64
	// The audio track’s average bit rate, in bits per second.
	AverageAudioBitrate() float64
	// The video track’s average bit rate, in bits per second.
	AverageVideoBitrate() float64
	// The average throughput, in bits per second, required to play the stream, as advertised by the server.
	IndicatedAverageBitrate() float64
}

// Init initializes the instance.
func (p AVPlayerItemAccessLogEvent) Init() AVPlayerItemAccessLogEvent {
	rv := objc.Send[AVPlayerItemAccessLogEvent](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemAccessLogEvent) Autorelease() AVPlayerItemAccessLogEvent {
	rv := objc.Send[AVPlayerItemAccessLogEvent](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemAccessLogEvent creates a new AVPlayerItemAccessLogEvent instance.
func NewAVPlayerItemAccessLogEvent() AVPlayerItemAccessLogEvent {
	class := getAVPlayerItemAccessLogEventClass()
	rv := objc.Send[AVPlayerItemAccessLogEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The URI of the playback item.
//
// # Discussion
// 
// The property corresponds to “uri”.
// 
// The value of this property is `nil` if the URI is unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/uri
func (p AVPlayerItemAccessLogEvent) URI() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URI"))
	return foundation.NSStringFromID(rv).String()
}
// The IP address of the server that was the source of the last delivered
// media segment.
//
// # Discussion
// 
// The property corresponds to “s-ip”.
// 
// The value of this property is `nil` if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/serverAddress
func (p AVPlayerItemAccessLogEvent) ServerAddress() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("serverAddress"))
	return foundation.NSStringFromID(rv).String()
}
// A count of changes to the server address over the last uninterrupted period
// of playback.
//
// # Discussion
// 
// The property corresponds to “s-ip-changes”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfServerAddressChanges
func (p AVPlayerItemAccessLogEvent) NumberOfServerAddressChanges() int {
	rv := objc.Send[int](p.ID, objc.Sel("numberOfServerAddressChanges"))
	return rv
}
// The number of network read requests over a WWAN.
//
// # Discussion
// 
// The value of the property is negative if unknown.
// 
// Corresponds to “sc-wwan-count”.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/mediaRequestsWWAN
func (p AVPlayerItemAccessLogEvent) MediaRequestsWWAN() int {
	rv := objc.Send[int](p.ID, objc.Sel("mediaRequestsWWAN"))
	return rv
}
// The accumulated duration, in seconds, of active network transfer of bytes.
//
// # Discussion
// 
// The value of the property is negative if unknown.
// 
// Corresponds to “c-transfer-duration”.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/transferDuration
func (p AVPlayerItemAccessLogEvent) TransferDuration() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("transferDuration"))
	return rv
}
// The accumulated number of bytes transferred by the item.
//
// # Discussion
// 
// The property corresponds to “bytes”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfBytesTransferred
func (p AVPlayerItemAccessLogEvent) NumberOfBytesTransferred() int64 {
	rv := objc.Send[int64](p.ID, objc.Sel("numberOfBytesTransferred"))
	return rv
}
// The number of media read requests from the server to this client.
//
// # Discussion
// 
// For HTTP live streaming, this property contains the count of media requests
// downloaded from the server. For progressive-style HTTP media downloads, it
// contains a count of HTTP [GET] (byte-range) requests for the resource.
// 
// The property corresponds to “sc-count”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfMediaRequests
func (p AVPlayerItemAccessLogEvent) NumberOfMediaRequests() int {
	rv := objc.Send[int](p.ID, objc.Sel("numberOfMediaRequests"))
	return rv
}
// The date and time at which playback began for this event.
//
// # Discussion
// 
// The property corresponds to “date”.
// 
// The value of this property is `nil` if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/playbackStartDate
func (p AVPlayerItemAccessLogEvent) PlaybackStartDate() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("playbackStartDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// A GUID that identifies the playback session.
//
// # Discussion
// 
// This value is used in HTTP requests.
// 
// The property corresponds to “cs-guid”.
// 
// The value of this property is `nil` if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/playbackSessionID
func (p AVPlayerItemAccessLogEvent) PlaybackSessionID() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("playbackSessionID"))
	return foundation.NSStringFromID(rv).String()
}
// The offset, in seconds, in the playlist where the last uninterrupted period
// of playback began.
//
// # Discussion
// 
// The property corresponds to “c-start-time”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/playbackStartOffset
func (p AVPlayerItemAccessLogEvent) PlaybackStartOffset() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("playbackStartOffset"))
	return rv
}
// The playback type.
//
// # Discussion
// 
// The playback type can be live, VOD, or from a file. If `nil` is returned
// the playback type is unknown.
// 
// Corresponds to “s-playback-type”.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/playbackType
func (p AVPlayerItemAccessLogEvent) PlaybackType() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("playbackType"))
	return foundation.NSStringFromID(rv).String()
}
// The accumulated duration, in seconds, until the player item is ready to
// play.
//
// # Discussion
// 
// The value of the property is negative if unknown.
// 
// Corresponds to “c-startup-time”.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/startupTime
func (p AVPlayerItemAccessLogEvent) StartupTime() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("startupTime"))
	return rv
}
// The accumulated duration, in seconds, of the media played.
//
// # Discussion
// 
// The property corresponds to “c-duration-watched”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/durationWatched
func (p AVPlayerItemAccessLogEvent) DurationWatched() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("durationWatched"))
	return rv
}
// The total number of dropped video frames
//
// # Discussion
// 
// The property corresponds to “c-frames-dropped”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfDroppedVideoFrames
func (p AVPlayerItemAccessLogEvent) NumberOfDroppedVideoFrames() int {
	rv := objc.Send[int](p.ID, objc.Sel("numberOfDroppedVideoFrames"))
	return rv
}
// The total number of playback stalls encountered.
//
// # Discussion
// 
// The property corresponds to “c-stalls”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfStalls
func (p AVPlayerItemAccessLogEvent) NumberOfStalls() int {
	rv := objc.Send[int](p.ID, objc.Sel("numberOfStalls"))
	return rv
}
// The accumulated duration, in seconds, of the media segments downloaded.
//
// # Discussion
// 
// The property corresponds to “c-duration-downloaded”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/segmentsDownloadedDuration
func (p AVPlayerItemAccessLogEvent) SegmentsDownloadedDuration() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("segmentsDownloadedDuration"))
	return rv
}
// The total number of times that downloading the segments took too long.
//
// # Discussion
// 
// The value of the property is negative if unknown.
// 
// This property corresponds to “c-overdue”.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/downloadOverdue
func (p AVPlayerItemAccessLogEvent) DownloadOverdue() int {
	rv := objc.Send[int](p.ID, objc.Sel("downloadOverdue"))
	return rv
}
// The standard deviation of the observed segment download bit rates.
//
// # Discussion
// 
// The value of the property is negative if unknown.
// 
// Corresponds to “c-observed-bitrate-sd”.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/observedBitrateStandardDeviation
func (p AVPlayerItemAccessLogEvent) ObservedBitrateStandardDeviation() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("observedBitrateStandardDeviation"))
	return rv
}
// The bandwidth value that causes a switch, up or down, in the item’s
// quality being played.
//
// # Discussion
// 
// The value of the property is negative if unknown.
// 
// Corresponds to “c-switch-bitrate”.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/switchBitrate
func (p AVPlayerItemAccessLogEvent) SwitchBitrate() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("switchBitrate"))
	return rv
}
// The throughput, in bits per second, required to play the stream, as
// advertised by the server.
//
// # Discussion
// 
// The property corresponds to “sc-indicated-bitrate”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/indicatedBitrate
func (p AVPlayerItemAccessLogEvent) IndicatedBitrate() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("indicatedBitrate"))
	return rv
}
// The empirical throughput, in bits per second, across all media downloaded.
//
// # Discussion
// 
// The property corresponds to “c-observed-bitrate”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/observedBitrate
func (p AVPlayerItemAccessLogEvent) ObservedBitrate() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("observedBitrate"))
	return rv
}
// The audio track’s average bit rate, in bits per second.
//
// # Discussion
// 
// The property corresponds to “c-avg-audio-bitrate”.
// 
// This property returns a non-positive value if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/averageAudioBitrate
func (p AVPlayerItemAccessLogEvent) AverageAudioBitrate() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("averageAudioBitrate"))
	return rv
}
// The video track’s average bit rate, in bits per second.
//
// # Discussion
// 
// The property corresponds to “c-avg-video-bitrate”.
// 
// This property returns the average bitrate of the video track if it is
// unmuxed, or the average bitrate of the combined content if muxed. Measured
// in bits per second.
// 
// The value is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/averageVideoBitrate
func (p AVPlayerItemAccessLogEvent) AverageVideoBitrate() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("averageVideoBitrate"))
	return rv
}
// The average throughput, in bits per second, required to play the stream, as
// advertised by the server.
//
// # Discussion
// 
// The property corresponds to “sc-indicated-avg-bitrate”.
// 
// The value of this property is negative if unknown.
// 
// This property is not compatible with key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/indicatedAverageBitrate
func (p AVPlayerItemAccessLogEvent) IndicatedAverageBitrate() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("indicatedAverageBitrate"))
	return rv
}


// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemErrorLogEvent] class.
var (
	_AVPlayerItemErrorLogEventClass     AVPlayerItemErrorLogEventClass
	_AVPlayerItemErrorLogEventClassOnce sync.Once
)

func getAVPlayerItemErrorLogEventClass() AVPlayerItemErrorLogEventClass {
	_AVPlayerItemErrorLogEventClassOnce.Do(func() {
		_AVPlayerItemErrorLogEventClass = AVPlayerItemErrorLogEventClass{class: objc.GetClass("AVPlayerItemErrorLogEvent")}
	})
	return _AVPlayerItemErrorLogEventClass
}

// GetAVPlayerItemErrorLogEventClass returns the class object for AVPlayerItemErrorLogEvent.
func GetAVPlayerItemErrorLogEventClass() AVPlayerItemErrorLogEventClass {
	return getAVPlayerItemErrorLogEventClass()
}

type AVPlayerItemErrorLogEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemErrorLogEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemErrorLogEventClass) Alloc() AVPlayerItemErrorLogEvent {
	rv := objc.Send[AVPlayerItemErrorLogEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A single item in a player item’s error log.
//
// # Overview
//
// This object provides properties for accessing the data fields of each log
// event. Each event is a single entry in an [AVPlayerItem] object’s error
// log.
//
// These properties aren’t observable. For more information about key-value
// observing, see [Using Key-Value Observing in Swift].
//
// # Getting information about the event
//
//   - [AVPlayerItemErrorLogEvent.Date]: The date and time when the error occurred.
//   - [AVPlayerItemErrorLogEvent.URI]: The URI of the playback item that had an error.
//   - [AVPlayerItemErrorLogEvent.ServerAddress]: The IP address of the server that was the source of the error.
//   - [AVPlayerItemErrorLogEvent.PlaybackSessionID]: A GUID that identifies the playback session that had an error.
//   - [AVPlayerItemErrorLogEvent.ErrorStatusCode]: A unique error code identifier.
//   - [AVPlayerItemErrorLogEvent.ErrorDomain]: The domain of the error.
//   - [AVPlayerItemErrorLogEvent.ErrorComment]: A description of the error encountered.
//   - [AVPlayerItemErrorLogEvent.AllHTTPResponseHeaderFields]: The HTTP header fields the server returns.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent
//
// [Using Key-Value Observing in Swift]: https://developer.apple.com/documentation/Swift/using-key-value-observing-in-swift
type AVPlayerItemErrorLogEvent struct {
	objectivec.Object
}

// AVPlayerItemErrorLogEventFromID constructs a [AVPlayerItemErrorLogEvent] from an objc.ID.
//
// A single item in a player item’s error log.
func AVPlayerItemErrorLogEventFromID(id objc.ID) AVPlayerItemErrorLogEvent {
	return AVPlayerItemErrorLogEvent{objectivec.Object{ID: id}}
}

// NOTE: AVPlayerItemErrorLogEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemErrorLogEvent] class.
//
// # Getting information about the event
//
//   - [IAVPlayerItemErrorLogEvent.Date]: The date and time when the error occurred.
//   - [IAVPlayerItemErrorLogEvent.URI]: The URI of the playback item that had an error.
//   - [IAVPlayerItemErrorLogEvent.ServerAddress]: The IP address of the server that was the source of the error.
//   - [IAVPlayerItemErrorLogEvent.PlaybackSessionID]: A GUID that identifies the playback session that had an error.
//   - [IAVPlayerItemErrorLogEvent.ErrorStatusCode]: A unique error code identifier.
//   - [IAVPlayerItemErrorLogEvent.ErrorDomain]: The domain of the error.
//   - [IAVPlayerItemErrorLogEvent.ErrorComment]: A description of the error encountered.
//   - [IAVPlayerItemErrorLogEvent.AllHTTPResponseHeaderFields]: The HTTP header fields the server returns.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent
type IAVPlayerItemErrorLogEvent interface {
	objectivec.IObject

	// Topic: Getting information about the event

	// The date and time when the error occurred.
	Date() foundation.INSDate
	// The URI of the playback item that had an error.
	URI() string
	// The IP address of the server that was the source of the error.
	ServerAddress() string
	// A GUID that identifies the playback session that had an error.
	PlaybackSessionID() string
	// A unique error code identifier.
	ErrorStatusCode() int
	// The domain of the error.
	ErrorDomain() string
	// A description of the error encountered.
	ErrorComment() string
	// The HTTP header fields the server returns.
	AllHTTPResponseHeaderFields() foundation.INSDictionary
}

// Init initializes the instance.
func (p AVPlayerItemErrorLogEvent) Init() AVPlayerItemErrorLogEvent {
	rv := objc.Send[AVPlayerItemErrorLogEvent](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemErrorLogEvent) Autorelease() AVPlayerItemErrorLogEvent {
	rv := objc.Send[AVPlayerItemErrorLogEvent](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemErrorLogEvent creates a new AVPlayerItemErrorLogEvent instance.
func NewAVPlayerItemErrorLogEvent() AVPlayerItemErrorLogEvent {
	class := getAVPlayerItemErrorLogEventClass()
	rv := objc.Send[AVPlayerItemErrorLogEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The date and time when the error occurred.
//
// # Discussion
//
// The property corresponds to “date”.
//
// The value of this property may be `nil` if the date is unknown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/date
func (p AVPlayerItemErrorLogEvent) Date() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("date"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The URI of the playback item that had an error.
//
// # Discussion
//
// The property corresponds to “uri”.
//
// The value of this property may be `nil` if the URI is unknown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/uri
func (p AVPlayerItemErrorLogEvent) URI() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URI"))
	return foundation.NSStringFromID(rv).String()
}

// The IP address of the server that was the source of the error.
//
// # Discussion
//
// The property corresponds to “s-ip”.
//
// The value of this property can be either an IPv4 or IPv6 address, and may
// be `nil` if the address is unknown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/serverAddress
func (p AVPlayerItemErrorLogEvent) ServerAddress() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("serverAddress"))
	return foundation.NSStringFromID(rv).String()
}

// A GUID that identifies the playback session that had an error.
//
// # Discussion
//
// The property corresponds to “cs-guid”.
//
// The value of this property is used in HTTP requests, and may be `nil` if
// the GUID is unknown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/playbackSessionID
func (p AVPlayerItemErrorLogEvent) PlaybackSessionID() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("playbackSessionID"))
	return foundation.NSStringFromID(rv).String()
}

// A unique error code identifier.
//
// # Discussion
//
// The property corresponds to “status”.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/errorStatusCode
func (p AVPlayerItemErrorLogEvent) ErrorStatusCode() int {
	rv := objc.Send[int](p.ID, objc.Sel("errorStatusCode"))
	return rv
}

// The domain of the error.
//
// # Discussion
//
// The property corresponds to “domain”.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/errorDomain
func (p AVPlayerItemErrorLogEvent) ErrorDomain() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("errorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// A description of the error encountered.
//
// # Discussion
//
// The property corresponds to “comment”.
//
// The value of this property may be `nil` if further information is not
// available.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/errorComment
func (p AVPlayerItemErrorLogEvent) ErrorComment() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("errorComment"))
	return foundation.NSStringFromID(rv).String()
}

// The HTTP header fields the server returns.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/allHTTPResponseHeaderFields
func (p AVPlayerItemErrorLogEvent) AllHTTPResponseHeaderFields() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("allHTTPResponseHeaderFields"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

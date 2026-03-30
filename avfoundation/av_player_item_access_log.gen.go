// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemAccessLog] class.
var (
	_AVPlayerItemAccessLogClass     AVPlayerItemAccessLogClass
	_AVPlayerItemAccessLogClassOnce sync.Once
)

func getAVPlayerItemAccessLogClass() AVPlayerItemAccessLogClass {
	_AVPlayerItemAccessLogClassOnce.Do(func() {
		_AVPlayerItemAccessLogClass = AVPlayerItemAccessLogClass{class: objc.GetClass("AVPlayerItemAccessLog")}
	})
	return _AVPlayerItemAccessLogClass
}

// GetAVPlayerItemAccessLogClass returns the class object for AVPlayerItemAccessLog.
func GetAVPlayerItemAccessLogClass() AVPlayerItemAccessLogClass {
	return getAVPlayerItemAccessLogClass()
}

type AVPlayerItemAccessLogClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemAccessLogClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemAccessLogClass) Alloc() AVPlayerItemAccessLog {
	rv := objc.Send[AVPlayerItemAccessLog](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object used to retrieve the access log associated with a player item.
//
// # Overview
//
// An [AVPlayerItemAccessLog] object accumulates key metrics about network
// playback and presents them as a collection of [AVPlayerItemAccessLogEvent]
// instances. Each event instance collates the data that relates to each
// uninterrupted period of playback.
//
// # Accessing log data
//
//   - [AVPlayerItemAccessLog.Events]: A chronologically ordered array of player item access log events.
//   - [AVPlayerItemAccessLog.ExtendedLogData]: Returns a serialized representation of the access log in the Extended Log File Format.
//   - [AVPlayerItemAccessLog.ExtendedLogDataStringEncoding]: The string encoding of the extended log data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLog
type AVPlayerItemAccessLog struct {
	objectivec.Object
}

// AVPlayerItemAccessLogFromID constructs a [AVPlayerItemAccessLog] from an objc.ID.
//
// An object used to retrieve the access log associated with a player item.
func AVPlayerItemAccessLogFromID(id objc.ID) AVPlayerItemAccessLog {
	return AVPlayerItemAccessLog{objectivec.Object{ID: id}}
}

// NOTE: AVPlayerItemAccessLog adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemAccessLog] class.
//
// # Accessing log data
//
//   - [IAVPlayerItemAccessLog.Events]: A chronologically ordered array of player item access log events.
//   - [IAVPlayerItemAccessLog.ExtendedLogData]: Returns a serialized representation of the access log in the Extended Log File Format.
//   - [IAVPlayerItemAccessLog.ExtendedLogDataStringEncoding]: The string encoding of the extended log data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLog
type IAVPlayerItemAccessLog interface {
	objectivec.IObject

	// Topic: Accessing log data

	// A chronologically ordered array of player item access log events.
	Events() []AVPlayerItemAccessLogEvent
	// Returns a serialized representation of the access log in the Extended Log File Format.
	ExtendedLogData() foundation.INSData
	// The string encoding of the extended log data.
	ExtendedLogDataStringEncoding() uint
}

// Init initializes the instance.
func (p AVPlayerItemAccessLog) Init() AVPlayerItemAccessLog {
	rv := objc.Send[AVPlayerItemAccessLog](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemAccessLog) Autorelease() AVPlayerItemAccessLog {
	rv := objc.Send[AVPlayerItemAccessLog](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemAccessLog creates a new AVPlayerItemAccessLog instance.
func NewAVPlayerItemAccessLog() AVPlayerItemAccessLog {
	class := getAVPlayerItemAccessLogClass()
	rv := objc.Send[AVPlayerItemAccessLog](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a serialized representation of the access log in the Extended Log
// File Format.
//
// # Return Value
//
// A serialized representation of the access log in the Extended Log File
// Format.
//
// # Discussion
//
// This method converts the web server access log into a textual format that
// conforms to the W3C Extended Log File Format for web server log files. For
// more information, see [http://www.w3.org/pub/WWW/TR/WD-logfile.html].
//
// You can generate a string suitable for console output using:
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLog/extendedLogData()
//
// [http://www.w3.org/pub/WWW/TR/WD-logfile.html]: http://www.w3.org/pub/WWW/TR/WD-logfile.html
func (p AVPlayerItemAccessLog) ExtendedLogData() foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("extendedLogData"))
	return foundation.NSDataFromID(rv)
}

// A chronologically ordered array of player item access log events.
//
// # Discussion
//
// The array contains [AVPlayerItemAccessLogEvent] objects that represent the
// chronological sequence of events contained in the access log.
//
// This property isn’t observable. For more information about key-value
// observing, see [Using Key-Value Observing in Swift].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLog/events
//
// [Using Key-Value Observing in Swift]: https://developer.apple.com/documentation/Swift/using-key-value-observing-in-swift
func (p AVPlayerItemAccessLog) Events() []AVPlayerItemAccessLogEvent {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("events"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerItemAccessLogEvent {
		return AVPlayerItemAccessLogEventFromID(id)
	})
}

// The string encoding of the extended log data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLog/extendedLogDataStringEncoding
func (p AVPlayerItemAccessLog) ExtendedLogDataStringEncoding() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("extendedLogDataStringEncoding"))
	return rv
}

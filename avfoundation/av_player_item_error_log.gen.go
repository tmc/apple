// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemErrorLog] class.
var (
	_AVPlayerItemErrorLogClass     AVPlayerItemErrorLogClass
	_AVPlayerItemErrorLogClassOnce sync.Once
)

func getAVPlayerItemErrorLogClass() AVPlayerItemErrorLogClass {
	_AVPlayerItemErrorLogClassOnce.Do(func() {
		_AVPlayerItemErrorLogClass = AVPlayerItemErrorLogClass{class: objc.GetClass("AVPlayerItemErrorLog")}
	})
	return _AVPlayerItemErrorLogClass
}

// GetAVPlayerItemErrorLogClass returns the class object for AVPlayerItemErrorLog.
func GetAVPlayerItemErrorLogClass() AVPlayerItemErrorLogClass {
	return getAVPlayerItemErrorLogClass()
}

type AVPlayerItemErrorLogClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemErrorLogClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemErrorLogClass) Alloc() AVPlayerItemErrorLog {
	rv := objc.Send[AVPlayerItemErrorLog](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// The error log associated with a player item.
//
// # Accessing error data
//
//   - [AVPlayerItemErrorLog.Events]: A chronologically ordered array of player item error log event objects.
//   - [AVPlayerItemErrorLog.ExtendedLogData]: Returns a serialized representation of the error log in the Extended Log File Format.
//   - [AVPlayerItemErrorLog.ExtendedLogDataStringEncoding]: The string encoding of the extended log data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLog
type AVPlayerItemErrorLog struct {
	objectivec.Object
}

// AVPlayerItemErrorLogFromID constructs a [AVPlayerItemErrorLog] from an objc.ID.
//
// The error log associated with a player item.
func AVPlayerItemErrorLogFromID(id objc.ID) AVPlayerItemErrorLog {
	return AVPlayerItemErrorLog{objectivec.Object{ID: id}}
}

// NOTE: AVPlayerItemErrorLog adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemErrorLog] class.
//
// # Accessing error data
//
//   - [IAVPlayerItemErrorLog.Events]: A chronologically ordered array of player item error log event objects.
//   - [IAVPlayerItemErrorLog.ExtendedLogData]: Returns a serialized representation of the error log in the Extended Log File Format.
//   - [IAVPlayerItemErrorLog.ExtendedLogDataStringEncoding]: The string encoding of the extended log data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLog
type IAVPlayerItemErrorLog interface {
	objectivec.IObject

	// Topic: Accessing error data

	// A chronologically ordered array of player item error log event objects.
	Events() []AVPlayerItemErrorLogEvent
	// Returns a serialized representation of the error log in the Extended Log File Format.
	ExtendedLogData() foundation.INSData
	// The string encoding of the extended log data.
	ExtendedLogDataStringEncoding() uint
}

// Init initializes the instance.
func (p AVPlayerItemErrorLog) Init() AVPlayerItemErrorLog {
	rv := objc.Send[AVPlayerItemErrorLog](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemErrorLog) Autorelease() AVPlayerItemErrorLog {
	rv := objc.Send[AVPlayerItemErrorLog](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemErrorLog creates a new AVPlayerItemErrorLog instance.
func NewAVPlayerItemErrorLog() AVPlayerItemErrorLog {
	class := getAVPlayerItemErrorLogClass()
	rv := objc.Send[AVPlayerItemErrorLog](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a serialized representation of the error log in the Extended Log
// File Format.
//
// # Return Value
//
// A serialized representation of the error log in the Extended Log File
// Format.
//
// # Discussion
//
// This method converts the web server error log into a textual format that
// conforms to the W3C Extended Log File Format for web server log files. For
// more information, see [http://www.w3.org/pub/WWW/TR/WD-logfile.html].
//
// You can generate a string suitable for console output using:
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLog/extendedLogData()
//
// [http://www.w3.org/pub/WWW/TR/WD-logfile.html]: http://www.w3.org/pub/WWW/TR/WD-logfile.html
func (p AVPlayerItemErrorLog) ExtendedLogData() foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("extendedLogData"))
	return foundation.NSDataFromID(rv)
}

// A chronologically ordered array of player item error log event objects.
//
// # Discussion
//
// The array contains [AVPlayerItemErrorLogEvent] objects that represent the
// chronological sequence of events contained in the error log.
//
// This property isn’t observable. For more information about key-value
// observing, see [Using Key-Value Observing in Swift].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLog/events
//
// [Using Key-Value Observing in Swift]: https://developer.apple.com/documentation/Swift/using-key-value-observing-in-swift
func (p AVPlayerItemErrorLog) Events() []AVPlayerItemErrorLogEvent {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("events"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerItemErrorLogEvent {
		return AVPlayerItemErrorLogEventFromID(id)
	})
}

// The string encoding of the extended log data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLog/extendedLogDataStringEncoding
func (p AVPlayerItemErrorLog) ExtendedLogDataStringEncoding() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("extendedLogDataStringEncoding"))
	return rv
}

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureTimecodeSource] class.
var (
	_AVCaptureTimecodeSourceClass     AVCaptureTimecodeSourceClass
	_AVCaptureTimecodeSourceClassOnce sync.Once
)

func getAVCaptureTimecodeSourceClass() AVCaptureTimecodeSourceClass {
	_AVCaptureTimecodeSourceClassOnce.Do(func() {
		_AVCaptureTimecodeSourceClass = AVCaptureTimecodeSourceClass{class: objc.GetClass("AVCaptureTimecodeSource")}
	})
	return _AVCaptureTimecodeSourceClass
}

// GetAVCaptureTimecodeSourceClass returns the class object for AVCaptureTimecodeSource.
func GetAVCaptureTimecodeSourceClass() AVCaptureTimecodeSourceClass {
	return getAVCaptureTimecodeSourceClass()
}

type AVCaptureTimecodeSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureTimecodeSourceClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureTimecodeSourceClass) Alloc() AVCaptureTimecodeSource {
	rv := objc.Send[AVCaptureTimecodeSource](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Describes a timecode source that a timecode generator can synchronize to.
//
// # Overview
// 
// [AVCaptureTimecodeSource] provides information about a specific timecode
// source available for synchronization in [AVCaptureTimecodeGenerator]. It
// includes metadata such as the source’s name, type, and unique identifier.
//
// # Inspecting the source
//
//   - [AVCaptureTimecodeSource.DisplayName]: The name of the timecode source.
//   - [AVCaptureTimecodeSource.Type]: The type of timecode source.
//   - [AVCaptureTimecodeSource.Uuid]: A unique identifier for the timecode source.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source
type AVCaptureTimecodeSource struct {
	objectivec.Object
}

// AVCaptureTimecodeSourceFromID constructs a [AVCaptureTimecodeSource] from an objc.ID.
//
// Describes a timecode source that a timecode generator can synchronize to.
func AVCaptureTimecodeSourceFromID(id objc.ID) AVCaptureTimecodeSource {
	return AVCaptureTimecodeSource{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureTimecodeSource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureTimecodeSource] class.
//
// # Inspecting the source
//
//   - [IAVCaptureTimecodeSource.DisplayName]: The name of the timecode source.
//   - [IAVCaptureTimecodeSource.Type]: The type of timecode source.
//   - [IAVCaptureTimecodeSource.Uuid]: A unique identifier for the timecode source.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source
type IAVCaptureTimecodeSource interface {
	objectivec.IObject

	// Topic: Inspecting the source

	// The name of the timecode source.
	DisplayName() string
	// The type of timecode source.
	Type() AVCaptureTimecodeSourceType
	// A unique identifier for the timecode source.
	Uuid() foundation.NSUUID
}

// Init initializes the instance.
func (c AVCaptureTimecodeSource) Init() AVCaptureTimecodeSource {
	rv := objc.Send[AVCaptureTimecodeSource](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureTimecodeSource) Autorelease() AVCaptureTimecodeSource {
	rv := objc.Send[AVCaptureTimecodeSource](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureTimecodeSource creates a new AVCaptureTimecodeSource instance.
func NewAVCaptureTimecodeSource() AVCaptureTimecodeSource {
	class := getAVCaptureTimecodeSourceClass()
	rv := objc.Send[AVCaptureTimecodeSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The name of the timecode source.
//
// # Discussion
// 
// This property provides a descriptive name of the timecode source, useful
// for display in user interfaces or logging.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source/displayName
func (c AVCaptureTimecodeSource) DisplayName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}
// The type of timecode source.
//
// # Discussion
// 
// Indicates the type of timecode source, represented as a value from the
// [AVCaptureTimecodeSynchronizationSourceType] enum. This helps you identify
// the source for specific synchronization use cases, such as frame counter,
// real-time clock, MIDI, or HID.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source/type
func (c AVCaptureTimecodeSource) Type() AVCaptureTimecodeSourceType {
	rv := objc.Send[AVCaptureTimecodeSourceType](c.ID, objc.Sel("type"))
	return AVCaptureTimecodeSourceType(rv)
}
// A unique identifier for the timecode source.
//
// # Discussion
// 
// The UUID uniquely identifies this timecode source. It is particularly
// useful when multiple sources of the same type are available, allowing your
// application to distinguish between them.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source/uuid
func (c AVCaptureTimecodeSource) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}


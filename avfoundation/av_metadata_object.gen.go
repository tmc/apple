// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMetadataObject] class.
var (
	_AVMetadataObjectClass     AVMetadataObjectClass
	_AVMetadataObjectClassOnce sync.Once
)

func getAVMetadataObjectClass() AVMetadataObjectClass {
	_AVMetadataObjectClassOnce.Do(func() {
		_AVMetadataObjectClass = AVMetadataObjectClass{class: objc.GetClass("AVMetadataObject")}
	})
	return _AVMetadataObjectClass
}

// GetAVMetadataObjectClass returns the class object for AVMetadataObject.
func GetAVMetadataObjectClass() AVMetadataObjectClass {
	return getAVMetadataObjectClass()
}

type AVMetadataObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataObjectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataObjectClass) Alloc() AVMetadataObject {
	rv := objc.Send[AVMetadataObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for objects provided by a metadata capture output.
//
// # Overview
// 
// The [AVMetadataObject] class is an abstract class that defines the basic
// properties associated with a piece of metadata. These attributes reflect
// information either about the metadata itself or the media from which the
// metadata originated. Subclasses are responsible for providing appropriate
// values for each of the relevant properties.
// 
// You shouldn’t subclass [AVMetadataObject] directly. Instead, you use one
// of the defined subclasses provided by the AVFoundation framework.
// Similarly, you don’t create instances of this class yourself but use an
// [AVCaptureMetadataOutput] object to retrieve them from the captured data.
//
// # Inspecting the metadata
//
//   - [AVMetadataObject.Bounds]: The bounding rectangle associated with the metadata.
//   - [AVMetadataObject.Duration]: The duration of the media associated with this metadata object.
//   - [AVMetadataObject.Time]: The media time value associated with the metadata object.
//   - [AVMetadataObject.Type]: The type of metadata that this object provides.
//   - [AVMetadataObject.FixedFocus]: A BOOL indicating whether this metadata object represents a fixed focus.
//   - [AVMetadataObject.CinematicVideoFocusMode]: The current focus mode when an object is detected during a Cinematic Video recording.
//   - [AVMetadataObject.GroupID]: An identifier associated with a metadata object used to group it with other metadata objects belonging to a common parent.
//   - [AVMetadataObject.ObjectID]: A unique identifier for each detected object type (face, body, hands, heads and salient objects) in a collection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject
type AVMetadataObject struct {
	objectivec.Object
}

// AVMetadataObjectFromID constructs a [AVMetadataObject] from an objc.ID.
//
// The abstract superclass for objects provided by a metadata capture output.
func AVMetadataObjectFromID(id objc.ID) AVMetadataObject {
	return AVMetadataObject{objectivec.Object{ID: id}}
}
// NOTE: AVMetadataObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataObject] class.
//
// # Inspecting the metadata
//
//   - [IAVMetadataObject.Bounds]: The bounding rectangle associated with the metadata.
//   - [IAVMetadataObject.Duration]: The duration of the media associated with this metadata object.
//   - [IAVMetadataObject.Time]: The media time value associated with the metadata object.
//   - [IAVMetadataObject.Type]: The type of metadata that this object provides.
//   - [IAVMetadataObject.FixedFocus]: A BOOL indicating whether this metadata object represents a fixed focus.
//   - [IAVMetadataObject.CinematicVideoFocusMode]: The current focus mode when an object is detected during a Cinematic Video recording.
//   - [IAVMetadataObject.GroupID]: An identifier associated with a metadata object used to group it with other metadata objects belonging to a common parent.
//   - [IAVMetadataObject.ObjectID]: A unique identifier for each detected object type (face, body, hands, heads and salient objects) in a collection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject
type IAVMetadataObject interface {
	objectivec.IObject

	// Topic: Inspecting the metadata

	// The bounding rectangle associated with the metadata.
	Bounds() corefoundation.CGRect
	// The duration of the media associated with this metadata object.
	Duration() uintptr
	// The media time value associated with the metadata object.
	Time() uintptr
	// The type of metadata that this object provides.
	Type() AVMetadataObjectType
	// A BOOL indicating whether this metadata object represents a fixed focus.
	FixedFocus() bool
	// The current focus mode when an object is detected during a Cinematic Video recording.
	CinematicVideoFocusMode() AVCaptureCinematicVideoFocusMode
	// An identifier associated with a metadata object used to group it with other metadata objects belonging to a common parent.
	GroupID() int
	// A unique identifier for each detected object type (face, body, hands, heads and salient objects) in a collection.
	ObjectID() int
}

// Init initializes the instance.
func (m AVMetadataObject) Init() AVMetadataObject {
	rv := objc.Send[AVMetadataObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataObject) Autorelease() AVMetadataObject {
	rv := objc.Send[AVMetadataObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataObject creates a new AVMetadataObject instance.
func NewAVMetadataObject() AVMetadataObject {
	class := getAVMetadataObjectClass()
	rv := objc.Send[AVMetadataObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The bounding rectangle associated with the metadata.
//
// # Discussion
// 
// The bounding rectangle is specified relative to the picture or video of the
// corresponding media. The rectangle’s origin is always specified in the
// top-left corner, and the x and y axis extend down and to the right.
// 
// If the metadata has no bounding rectangle, the value of this property
// should be [CGRectZero].
// 
// For video content, the bounding rectangle may be expressed using scalar
// values in the range 0.0 to 1.0. Scalar values remain meaningful even when
// the original video has been scaled down.
//
// [CGRectZero]: https://developer.apple.com/documentation/CoreGraphics/CGRectZero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/bounds
func (m AVMetadataObject) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
// The duration of the media associated with this metadata object.
//
// # Discussion
// 
// For metadata originating from a sample buffer ([CMSampleBuffer]), the
// duration reflects the duration of the sample buffer. If there is no valid
// duration value associated with the metadata, this property should contain
// [invalid].
//
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/duration
func (m AVMetadataObject) Duration() uintptr {
	rv := objc.Send[uintptr](m.ID, objc.Sel("duration"))
	return rv
}
// The media time value associated with the metadata object.
//
// # Discussion
// 
// For captured media, this property represents the time when the metadata was
// captured. For metadata originating from a sample buffer ([CMSampleBuffer]),
// the time is the sample buffer’s presentation time. If there is no valid
// time value associated with the metadata, this property should contain
// [invalid].
//
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/time
func (m AVMetadataObject) Time() uintptr {
	rv := objc.Send[uintptr](m.ID, objc.Sel("time"))
	return rv
}
// The type of metadata that this object provides.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/type
func (m AVMetadataObject) Type() AVMetadataObjectType {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("type"))
	return AVMetadataObjectType(foundation.NSStringFromID(rv).String())
}
// A BOOL indicating whether this metadata object represents a fixed focus.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/isFixedFocus
func (m AVMetadataObject) FixedFocus() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isFixedFocus"))
	return rv
}
// The current focus mode when an object is detected during a Cinematic Video
// recording.
//
// # Discussion
// 
// Default is [CaptureCinematicVideoFocusModeNone].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/cinematicVideoFocusMode
func (m AVMetadataObject) CinematicVideoFocusMode() AVCaptureCinematicVideoFocusMode {
	rv := objc.Send[AVCaptureCinematicVideoFocusMode](m.ID, objc.Sel("cinematicVideoFocusMode"))
	return AVCaptureCinematicVideoFocusMode(rv)
}
// An identifier associated with a metadata object used to group it with other
// metadata objects belonging to a common parent.
//
// # Discussion
// 
// When presented with a collection of [AVMetadataObject] instances of
// different types, you may use the objects’ [GroupID] to combine them into
// groups. For example, a human body and face belonging to the same person
// have the same [GroupID]. If an object’s [GroupID] property is set to -1,
// it is invalid. When set to a value of >=0, it is unique across all object
// groups.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/groupID
func (m AVMetadataObject) GroupID() int {
	rv := objc.Send[int](m.ID, objc.Sel("groupID"))
	return rv
}
// A unique identifier for each detected object type (face, body, hands, heads
// and salient objects) in a collection.
//
// # Discussion
// 
// Defaults to a value of -1 when invalid or not available. When used in
// conjunction with an [AVCaptureMetadataOutput], each newly detected object
// that enters the scene is assigned a unique identifier. [ObjectID]s are
// never re-used as objects leave the picture and new ones enter. Objects that
// leave the picture and then re-enter are assigned a new [ObjectID].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/objectID
func (m AVMetadataObject) ObjectID() int {
	rv := objc.Send[int](m.ID, objc.Sel("objectID"))
	return rv
}


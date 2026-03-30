// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetadataBodyObject] class.
var (
	_AVMetadataBodyObjectClass     AVMetadataBodyObjectClass
	_AVMetadataBodyObjectClassOnce sync.Once
)

func getAVMetadataBodyObjectClass() AVMetadataBodyObjectClass {
	_AVMetadataBodyObjectClassOnce.Do(func() {
		_AVMetadataBodyObjectClass = AVMetadataBodyObjectClass{class: objc.GetClass("AVMetadataBodyObject")}
	})
	return _AVMetadataBodyObjectClass
}

// GetAVMetadataBodyObjectClass returns the class object for AVMetadataBodyObject.
func GetAVMetadataBodyObjectClass() AVMetadataBodyObjectClass {
	return getAVMetadataBodyObjectClass()
}

type AVMetadataBodyObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataBodyObjectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataBodyObjectClass) Alloc() AVMetadataBodyObject {
	rv := objc.Send[AVMetadataBodyObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that defines the interface for a metadata body object.
//
// # Overview
//
// A metadata body object represents a single detected body in a picture.
// It’s the base object used to represent bodies, for example
// [AVMetadataHumanBodyObject], [AVMetadataDogBodyObject], and
// [AVMetadataCatBodyObject].
//
// # Inspecting metadata
//
//   - [AVMetadataBodyObject.BodyID]: An integer value that defines the unique identifier of an object in a picture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataBodyObject
type AVMetadataBodyObject struct {
	AVMetadataObject
}

// AVMetadataBodyObjectFromID constructs a [AVMetadataBodyObject] from an objc.ID.
//
// An abstract class that defines the interface for a metadata body object.
func AVMetadataBodyObjectFromID(id objc.ID) AVMetadataBodyObject {
	return AVMetadataBodyObject{AVMetadataObject: AVMetadataObjectFromID(id)}
}

// NOTE: AVMetadataBodyObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataBodyObject] class.
//
// # Inspecting metadata
//
//   - [IAVMetadataBodyObject.BodyID]: An integer value that defines the unique identifier of an object in a picture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataBodyObject
type IAVMetadataBodyObject interface {
	IAVMetadataObject

	// Topic: Inspecting metadata

	// An integer value that defines the unique identifier of an object in a picture.
	BodyID() int
}

// Init initializes the instance.
func (m AVMetadataBodyObject) Init() AVMetadataBodyObject {
	rv := objc.Send[AVMetadataBodyObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataBodyObject) Autorelease() AVMetadataBodyObject {
	rv := objc.Send[AVMetadataBodyObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataBodyObject creates a new AVMetadataBodyObject instance.
func NewAVMetadataBodyObject() AVMetadataBodyObject {
	class := getAVMetadataBodyObjectClass()
	rv := objc.Send[AVMetadataBodyObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An integer value that defines the unique identifier of an object in a
// picture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataBodyObject/bodyID
func (m AVMetadataBodyObject) BodyID() int {
	rv := objc.Send[int](m.ID, objc.Sel("bodyID"))
	return rv
}

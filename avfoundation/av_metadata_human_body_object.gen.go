// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetadataHumanBodyObject] class.
var (
	_AVMetadataHumanBodyObjectClass     AVMetadataHumanBodyObjectClass
	_AVMetadataHumanBodyObjectClassOnce sync.Once
)

func getAVMetadataHumanBodyObjectClass() AVMetadataHumanBodyObjectClass {
	_AVMetadataHumanBodyObjectClassOnce.Do(func() {
		_AVMetadataHumanBodyObjectClass = AVMetadataHumanBodyObjectClass{class: objc.GetClass("AVMetadataHumanBodyObject")}
	})
	return _AVMetadataHumanBodyObjectClass
}

// GetAVMetadataHumanBodyObjectClass returns the class object for AVMetadataHumanBodyObject.
func GetAVMetadataHumanBodyObjectClass() AVMetadataHumanBodyObjectClass {
	return getAVMetadataHumanBodyObjectClass()
}

type AVMetadataHumanBodyObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataHumanBodyObjectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataHumanBodyObjectClass) Alloc() AVMetadataHumanBodyObject {
	rv := objc.Send[AVMetadataHumanBodyObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object representing a single detected human body in a picture.
//
// # Overview
//
// This object is an immutable type that describes the various features found
// in the human body in a picture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataHumanBodyObject
type AVMetadataHumanBodyObject struct {
	AVMetadataBodyObject
}

// AVMetadataHumanBodyObjectFromID constructs a [AVMetadataHumanBodyObject] from an objc.ID.
//
// An object representing a single detected human body in a picture.
func AVMetadataHumanBodyObjectFromID(id objc.ID) AVMetadataHumanBodyObject {
	return AVMetadataHumanBodyObject{AVMetadataBodyObject: AVMetadataBodyObjectFromID(id)}
}

// NOTE: AVMetadataHumanBodyObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataHumanBodyObject] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataHumanBodyObject
type IAVMetadataHumanBodyObject interface {
	IAVMetadataBodyObject
}

// Init initializes the instance.
func (m AVMetadataHumanBodyObject) Init() AVMetadataHumanBodyObject {
	rv := objc.Send[AVMetadataHumanBodyObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataHumanBodyObject) Autorelease() AVMetadataHumanBodyObject {
	rv := objc.Send[AVMetadataHumanBodyObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataHumanBodyObject creates a new AVMetadataHumanBodyObject instance.
func NewAVMetadataHumanBodyObject() AVMetadataHumanBodyObject {
	class := getAVMetadataHumanBodyObjectClass()
	rv := objc.Send[AVMetadataHumanBodyObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

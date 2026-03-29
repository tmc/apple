// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetadataDogBodyObject] class.
var (
	_AVMetadataDogBodyObjectClass     AVMetadataDogBodyObjectClass
	_AVMetadataDogBodyObjectClassOnce sync.Once
)

func getAVMetadataDogBodyObjectClass() AVMetadataDogBodyObjectClass {
	_AVMetadataDogBodyObjectClassOnce.Do(func() {
		_AVMetadataDogBodyObjectClass = AVMetadataDogBodyObjectClass{class: objc.GetClass("AVMetadataDogBodyObject")}
	})
	return _AVMetadataDogBodyObjectClass
}

// GetAVMetadataDogBodyObjectClass returns the class object for AVMetadataDogBodyObject.
func GetAVMetadataDogBodyObjectClass() AVMetadataDogBodyObjectClass {
	return getAVMetadataDogBodyObjectClass()
}

type AVMetadataDogBodyObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataDogBodyObjectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataDogBodyObjectClass) Alloc() AVMetadataDogBodyObject {
	rv := objc.Send[AVMetadataDogBodyObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object representing a single detected dog body in a picture.
//
// # Overview
// 
// This object is an immutable type that describes the various features found
// in the dog body in a picture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataDogBodyObject
type AVMetadataDogBodyObject struct {
	AVMetadataBodyObject
}

// AVMetadataDogBodyObjectFromID constructs a [AVMetadataDogBodyObject] from an objc.ID.
//
// An object representing a single detected dog body in a picture.
func AVMetadataDogBodyObjectFromID(id objc.ID) AVMetadataDogBodyObject {
	return AVMetadataDogBodyObject{AVMetadataBodyObject: AVMetadataBodyObjectFromID(id)}
}
// NOTE: AVMetadataDogBodyObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataDogBodyObject] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataDogBodyObject
type IAVMetadataDogBodyObject interface {
	IAVMetadataBodyObject
}

// Init initializes the instance.
func (m AVMetadataDogBodyObject) Init() AVMetadataDogBodyObject {
	rv := objc.Send[AVMetadataDogBodyObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataDogBodyObject) Autorelease() AVMetadataDogBodyObject {
	rv := objc.Send[AVMetadataDogBodyObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataDogBodyObject creates a new AVMetadataDogBodyObject instance.
func NewAVMetadataDogBodyObject() AVMetadataDogBodyObject {
	class := getAVMetadataDogBodyObjectClass()
	rv := objc.Send[AVMetadataDogBodyObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}


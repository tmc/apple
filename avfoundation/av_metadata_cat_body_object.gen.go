// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetadataCatBodyObject] class.
var (
	_AVMetadataCatBodyObjectClass     AVMetadataCatBodyObjectClass
	_AVMetadataCatBodyObjectClassOnce sync.Once
)

func getAVMetadataCatBodyObjectClass() AVMetadataCatBodyObjectClass {
	_AVMetadataCatBodyObjectClassOnce.Do(func() {
		_AVMetadataCatBodyObjectClass = AVMetadataCatBodyObjectClass{class: objc.GetClass("AVMetadataCatBodyObject")}
	})
	return _AVMetadataCatBodyObjectClass
}

// GetAVMetadataCatBodyObjectClass returns the class object for AVMetadataCatBodyObject.
func GetAVMetadataCatBodyObjectClass() AVMetadataCatBodyObjectClass {
	return getAVMetadataCatBodyObjectClass()
}

type AVMetadataCatBodyObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataCatBodyObjectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataCatBodyObjectClass) Alloc() AVMetadataCatBodyObject {
	rv := objc.Send[AVMetadataCatBodyObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object representing a single detected cat body in a picture.
//
// # Overview
// 
// This object is an immutable type that describes the various features found
// in the cat body in a picture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataCatBodyObject
type AVMetadataCatBodyObject struct {
	AVMetadataBodyObject
}

// AVMetadataCatBodyObjectFromID constructs a [AVMetadataCatBodyObject] from an objc.ID.
//
// An object representing a single detected cat body in a picture.
func AVMetadataCatBodyObjectFromID(id objc.ID) AVMetadataCatBodyObject {
	return AVMetadataCatBodyObject{AVMetadataBodyObject: AVMetadataBodyObjectFromID(id)}
}
// NOTE: AVMetadataCatBodyObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataCatBodyObject] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataCatBodyObject
type IAVMetadataCatBodyObject interface {
	IAVMetadataBodyObject
}

// Init initializes the instance.
func (m AVMetadataCatBodyObject) Init() AVMetadataCatBodyObject {
	rv := objc.Send[AVMetadataCatBodyObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataCatBodyObject) Autorelease() AVMetadataCatBodyObject {
	rv := objc.Send[AVMetadataCatBodyObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataCatBodyObject creates a new AVMetadataCatBodyObject instance.
func NewAVMetadataCatBodyObject() AVMetadataCatBodyObject {
	class := getAVMetadataCatBodyObjectClass()
	rv := objc.Send[AVMetadataCatBodyObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}


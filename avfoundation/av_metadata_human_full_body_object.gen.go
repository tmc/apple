// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetadataHumanFullBodyObject] class.
var (
	_AVMetadataHumanFullBodyObjectClass     AVMetadataHumanFullBodyObjectClass
	_AVMetadataHumanFullBodyObjectClassOnce sync.Once
)

func getAVMetadataHumanFullBodyObjectClass() AVMetadataHumanFullBodyObjectClass {
	_AVMetadataHumanFullBodyObjectClassOnce.Do(func() {
		_AVMetadataHumanFullBodyObjectClass = AVMetadataHumanFullBodyObjectClass{class: objc.GetClass("AVMetadataHumanFullBodyObject")}
	})
	return _AVMetadataHumanFullBodyObjectClass
}

// GetAVMetadataHumanFullBodyObjectClass returns the class object for AVMetadataHumanFullBodyObject.
func GetAVMetadataHumanFullBodyObjectClass() AVMetadataHumanFullBodyObjectClass {
	return getAVMetadataHumanFullBodyObjectClass()
}

type AVMetadataHumanFullBodyObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataHumanFullBodyObjectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataHumanFullBodyObjectClass) Alloc() AVMetadataHumanFullBodyObject {
	rv := objc.Send[AVMetadataHumanFullBodyObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a detected human full body in a picture.
//
// # Overview
//
// On supported platforms, [AVCaptureMetadataOutput] outputs arrays of
// detected human full body objects.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataHumanFullBodyObject
type AVMetadataHumanFullBodyObject struct {
	AVMetadataBodyObject
}

// AVMetadataHumanFullBodyObjectFromID constructs a [AVMetadataHumanFullBodyObject] from an objc.ID.
//
// An object that represents a detected human full body in a picture.
func AVMetadataHumanFullBodyObjectFromID(id objc.ID) AVMetadataHumanFullBodyObject {
	return AVMetadataHumanFullBodyObject{AVMetadataBodyObject: AVMetadataBodyObjectFromID(id)}
}

// NOTE: AVMetadataHumanFullBodyObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataHumanFullBodyObject] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataHumanFullBodyObject
type IAVMetadataHumanFullBodyObject interface {
	IAVMetadataBodyObject
}

// Init initializes the instance.
func (m AVMetadataHumanFullBodyObject) Init() AVMetadataHumanFullBodyObject {
	rv := objc.Send[AVMetadataHumanFullBodyObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataHumanFullBodyObject) Autorelease() AVMetadataHumanFullBodyObject {
	rv := objc.Send[AVMetadataHumanFullBodyObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataHumanFullBodyObject creates a new AVMetadataHumanFullBodyObject instance.
func NewAVMetadataHumanFullBodyObject() AVMetadataHumanFullBodyObject {
	class := getAVMetadataHumanFullBodyObjectClass()
	rv := objc.Send[AVMetadataHumanFullBodyObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

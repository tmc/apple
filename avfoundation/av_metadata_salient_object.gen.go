// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetadataSalientObject] class.
var (
	_AVMetadataSalientObjectClass     AVMetadataSalientObjectClass
	_AVMetadataSalientObjectClassOnce sync.Once
)

func getAVMetadataSalientObjectClass() AVMetadataSalientObjectClass {
	_AVMetadataSalientObjectClassOnce.Do(func() {
		_AVMetadataSalientObjectClass = AVMetadataSalientObjectClass{class: objc.GetClass("AVMetadataSalientObject")}
	})
	return _AVMetadataSalientObjectClass
}

// GetAVMetadataSalientObjectClass returns the class object for AVMetadataSalientObject.
func GetAVMetadataSalientObjectClass() AVMetadataSalientObjectClass {
	return getAVMetadataSalientObjectClass()
}

type AVMetadataSalientObjectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataSalientObjectClass) Alloc() AVMetadataSalientObject {
	rv := objc.Send[AVMetadataSalientObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object representing a single salient area in a picture.
//
// # Overview
// 
// This object is an immutable type that describes the various features of the
// salient object in a picture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataSalientObject
type AVMetadataSalientObject struct {
	AVMetadataObject
}

// AVMetadataSalientObjectFromID constructs a [AVMetadataSalientObject] from an objc.ID.
//
// An object representing a single salient area in a picture.
func AVMetadataSalientObjectFromID(id objc.ID) AVMetadataSalientObject {
	return AVMetadataSalientObject{AVMetadataObject: AVMetadataObjectFromID(id)}
}
// NOTE: AVMetadataSalientObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataSalientObject] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataSalientObject
type IAVMetadataSalientObject interface {
	IAVMetadataObject
}

// Init initializes the instance.
func (m AVMetadataSalientObject) Init() AVMetadataSalientObject {
	rv := objc.Send[AVMetadataSalientObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataSalientObject) Autorelease() AVMetadataSalientObject {
	rv := objc.Send[AVMetadataSalientObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataSalientObject creates a new AVMetadataSalientObject instance.
func NewAVMetadataSalientObject() AVMetadataSalientObject {
	class := getAVMetadataSalientObjectClass()
	rv := objc.Send[AVMetadataSalientObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}


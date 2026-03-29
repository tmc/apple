// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetadataCatHeadObject] class.
var (
	_AVMetadataCatHeadObjectClass     AVMetadataCatHeadObjectClass
	_AVMetadataCatHeadObjectClassOnce sync.Once
)

func getAVMetadataCatHeadObjectClass() AVMetadataCatHeadObjectClass {
	_AVMetadataCatHeadObjectClassOnce.Do(func() {
		_AVMetadataCatHeadObjectClass = AVMetadataCatHeadObjectClass{class: objc.GetClass("AVMetadataCatHeadObject")}
	})
	return _AVMetadataCatHeadObjectClass
}

// GetAVMetadataCatHeadObjectClass returns the class object for AVMetadataCatHeadObject.
func GetAVMetadataCatHeadObjectClass() AVMetadataCatHeadObjectClass {
	return getAVMetadataCatHeadObjectClass()
}

type AVMetadataCatHeadObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataCatHeadObjectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataCatHeadObjectClass) Alloc() AVMetadataCatHeadObject {
	rv := objc.Send[AVMetadataCatHeadObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A concrete metadata object subclass representing a cat head.
//
// # Overview
// 
// [AVMetadataCatHeadObject] is a concrete subclass of [AVMetadataObject]
// representing a cat head.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataCatHeadObject
type AVMetadataCatHeadObject struct {
	AVMetadataObject
}

// AVMetadataCatHeadObjectFromID constructs a [AVMetadataCatHeadObject] from an objc.ID.
//
// A concrete metadata object subclass representing a cat head.
func AVMetadataCatHeadObjectFromID(id objc.ID) AVMetadataCatHeadObject {
	return AVMetadataCatHeadObject{AVMetadataObject: AVMetadataObjectFromID(id)}
}
// NOTE: AVMetadataCatHeadObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataCatHeadObject] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataCatHeadObject
type IAVMetadataCatHeadObject interface {
	IAVMetadataObject
}

// Init initializes the instance.
func (m AVMetadataCatHeadObject) Init() AVMetadataCatHeadObject {
	rv := objc.Send[AVMetadataCatHeadObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataCatHeadObject) Autorelease() AVMetadataCatHeadObject {
	rv := objc.Send[AVMetadataCatHeadObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataCatHeadObject creates a new AVMetadataCatHeadObject instance.
func NewAVMetadataCatHeadObject() AVMetadataCatHeadObject {
	class := getAVMetadataCatHeadObjectClass()
	rv := objc.Send[AVMetadataCatHeadObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}


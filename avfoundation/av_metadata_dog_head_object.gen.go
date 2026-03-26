// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetadataDogHeadObject] class.
var (
	_AVMetadataDogHeadObjectClass     AVMetadataDogHeadObjectClass
	_AVMetadataDogHeadObjectClassOnce sync.Once
)

func getAVMetadataDogHeadObjectClass() AVMetadataDogHeadObjectClass {
	_AVMetadataDogHeadObjectClassOnce.Do(func() {
		_AVMetadataDogHeadObjectClass = AVMetadataDogHeadObjectClass{class: objc.GetClass("AVMetadataDogHeadObject")}
	})
	return _AVMetadataDogHeadObjectClass
}

// GetAVMetadataDogHeadObjectClass returns the class object for AVMetadataDogHeadObject.
func GetAVMetadataDogHeadObjectClass() AVMetadataDogHeadObjectClass {
	return getAVMetadataDogHeadObjectClass()
}

type AVMetadataDogHeadObjectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataDogHeadObjectClass) Alloc() AVMetadataDogHeadObject {
	rv := objc.Send[AVMetadataDogHeadObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A concrete metadata object subclass representing a dog head.
//
// # Overview
// 
// [AVMetadataDogHeadObject] is a concrete subclass of [AVMetadataObject]
// representing a dog head.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataDogHeadObject
type AVMetadataDogHeadObject struct {
	AVMetadataObject
}

// AVMetadataDogHeadObjectFromID constructs a [AVMetadataDogHeadObject] from an objc.ID.
//
// A concrete metadata object subclass representing a dog head.
func AVMetadataDogHeadObjectFromID(id objc.ID) AVMetadataDogHeadObject {
	return AVMetadataDogHeadObject{AVMetadataObject: AVMetadataObjectFromID(id)}
}
// NOTE: AVMetadataDogHeadObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataDogHeadObject] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataDogHeadObject
type IAVMetadataDogHeadObject interface {
	IAVMetadataObject
}

// Init initializes the instance.
func (m AVMetadataDogHeadObject) Init() AVMetadataDogHeadObject {
	rv := objc.Send[AVMetadataDogHeadObject](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataDogHeadObject) Autorelease() AVMetadataDogHeadObject {
	rv := objc.Send[AVMetadataDogHeadObject](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataDogHeadObject creates a new AVMetadataDogHeadObject instance.
func NewAVMetadataDogHeadObject() AVMetadataDogHeadObject {
	class := getAVMetadataDogHeadObjectClass()
	rv := objc.Send[AVMetadataDogHeadObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}


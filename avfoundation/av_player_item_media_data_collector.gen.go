// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemMediaDataCollector] class.
var (
	_AVPlayerItemMediaDataCollectorClass     AVPlayerItemMediaDataCollectorClass
	_AVPlayerItemMediaDataCollectorClassOnce sync.Once
)

func getAVPlayerItemMediaDataCollectorClass() AVPlayerItemMediaDataCollectorClass {
	_AVPlayerItemMediaDataCollectorClassOnce.Do(func() {
		_AVPlayerItemMediaDataCollectorClass = AVPlayerItemMediaDataCollectorClass{class: objc.GetClass("AVPlayerItemMediaDataCollector")}
	})
	return _AVPlayerItemMediaDataCollectorClass
}

// GetAVPlayerItemMediaDataCollectorClass returns the class object for AVPlayerItemMediaDataCollector.
func GetAVPlayerItemMediaDataCollectorClass() AVPlayerItemMediaDataCollectorClass {
	return getAVPlayerItemMediaDataCollectorClass()
}

type AVPlayerItemMediaDataCollectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemMediaDataCollectorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemMediaDataCollectorClass) Alloc() AVPlayerItemMediaDataCollector {
	rv := objc.Send[AVPlayerItemMediaDataCollector](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// The abstract base for media data collectors.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMediaDataCollector
type AVPlayerItemMediaDataCollector struct {
	objectivec.Object
}

// AVPlayerItemMediaDataCollectorFromID constructs a [AVPlayerItemMediaDataCollector] from an objc.ID.
//
// The abstract base for media data collectors.
func AVPlayerItemMediaDataCollectorFromID(id objc.ID) AVPlayerItemMediaDataCollector {
	return AVPlayerItemMediaDataCollector{objectivec.Object{ID: id}}
}
// NOTE: AVPlayerItemMediaDataCollector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemMediaDataCollector] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMediaDataCollector
type IAVPlayerItemMediaDataCollector interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p AVPlayerItemMediaDataCollector) Init() AVPlayerItemMediaDataCollector {
	rv := objc.Send[AVPlayerItemMediaDataCollector](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemMediaDataCollector) Autorelease() AVPlayerItemMediaDataCollector {
	rv := objc.Send[AVPlayerItemMediaDataCollector](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemMediaDataCollector creates a new AVPlayerItemMediaDataCollector instance.
func NewAVPlayerItemMediaDataCollector() AVPlayerItemMediaDataCollector {
	class := getAVPlayerItemMediaDataCollectorClass()
	rv := objc.Send[AVPlayerItemMediaDataCollector](objc.ID(class.class), objc.Sel("new"))
	return rv
}


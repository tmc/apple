// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVideoCompositionRenderHint] class.
var (
	_AVVideoCompositionRenderHintClass     AVVideoCompositionRenderHintClass
	_AVVideoCompositionRenderHintClassOnce sync.Once
)

func getAVVideoCompositionRenderHintClass() AVVideoCompositionRenderHintClass {
	_AVVideoCompositionRenderHintClassOnce.Do(func() {
		_AVVideoCompositionRenderHintClass = AVVideoCompositionRenderHintClass{class: objc.GetClass("AVVideoCompositionRenderHint")}
	})
	return _AVVideoCompositionRenderHintClass
}

// GetAVVideoCompositionRenderHintClass returns the class object for AVVideoCompositionRenderHint.
func GetAVVideoCompositionRenderHintClass() AVVideoCompositionRenderHintClass {
	return getAVVideoCompositionRenderHintClass()
}

type AVVideoCompositionRenderHintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVideoCompositionRenderHintClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVideoCompositionRenderHintClass) Alloc() AVVideoCompositionRenderHint {
	rv := objc.Send[AVVideoCompositionRenderHint](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Information about upcoming composition requests, such as composition start
// time and end time.
//
// # Managing composition timing
//
//   - [AVVideoCompositionRenderHint.StartCompositionTime]: The start time of the upcoming composition requests.
//   - [AVVideoCompositionRenderHint.EndCompositionTime]: The end time of the upcoming composition requests.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderHint
type AVVideoCompositionRenderHint struct {
	objectivec.Object
}

// AVVideoCompositionRenderHintFromID constructs a [AVVideoCompositionRenderHint] from an objc.ID.
//
// Information about upcoming composition requests, such as composition start
// time and end time.
func AVVideoCompositionRenderHintFromID(id objc.ID) AVVideoCompositionRenderHint {
	return AVVideoCompositionRenderHint{objectivec.Object{ID: id}}
}
// NOTE: AVVideoCompositionRenderHint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVVideoCompositionRenderHint] class.
//
// # Managing composition timing
//
//   - [IAVVideoCompositionRenderHint.StartCompositionTime]: The start time of the upcoming composition requests.
//   - [IAVVideoCompositionRenderHint.EndCompositionTime]: The end time of the upcoming composition requests.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderHint
type IAVVideoCompositionRenderHint interface {
	objectivec.IObject

	// Topic: Managing composition timing

	// The start time of the upcoming composition requests.
	StartCompositionTime() coremedia.CMTime
	// The end time of the upcoming composition requests.
	EndCompositionTime() coremedia.CMTime
}

// Init initializes the instance.
func (v AVVideoCompositionRenderHint) Init() AVVideoCompositionRenderHint {
	rv := objc.Send[AVVideoCompositionRenderHint](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVideoCompositionRenderHint) Autorelease() AVVideoCompositionRenderHint {
	rv := objc.Send[AVVideoCompositionRenderHint](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVideoCompositionRenderHint creates a new AVVideoCompositionRenderHint instance.
func NewAVVideoCompositionRenderHint() AVVideoCompositionRenderHint {
	class := getAVVideoCompositionRenderHintClass()
	rv := objc.Send[AVVideoCompositionRenderHint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The start time of the upcoming composition requests.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderHint/startCompositionTime
func (v AVVideoCompositionRenderHint) StartCompositionTime() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](v.ID, objc.Sel("startCompositionTime"))
	return coremedia.CMTime(rv)
}
// The end time of the upcoming composition requests.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderHint/endCompositionTime
func (v AVVideoCompositionRenderHint) EndCompositionTime() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](v.ID, objc.Sel("endCompositionTime"))
	return coremedia.CMTime(rv)
}


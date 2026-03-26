// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVRouteDetector] class.
var (
	_AVRouteDetectorClass     AVRouteDetectorClass
	_AVRouteDetectorClassOnce sync.Once
)

func getAVRouteDetectorClass() AVRouteDetectorClass {
	_AVRouteDetectorClassOnce.Do(func() {
		_AVRouteDetectorClass = AVRouteDetectorClass{class: objc.GetClass("AVRouteDetector")}
	})
	return _AVRouteDetectorClass
}

// GetAVRouteDetectorClass returns the class object for AVRouteDetector.
func GetAVRouteDetectorClass() AVRouteDetectorClass {
	return getAVRouteDetectorClass()
}

type AVRouteDetectorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVRouteDetectorClass) Alloc() AVRouteDetector {
	rv := objc.Send[AVRouteDetector](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that detects available media playback routes.
//
// # Overview
// 
// If you enable route detection, the object reports whether it detects
// multiple playback routes. If it does, use [AVRoutePickerView] to present
// the UI for the user to select an appropriate route.
//
// [AVRoutePickerView]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView
//
// # Detecting routes
//
//   - [AVRouteDetector.RouteDetectionEnabled]: A Boolean value that indicates whether route detection is in an enabled state.
//   - [AVRouteDetector.SetRouteDetectionEnabled]
//   - [AVRouteDetector.MultipleRoutesDetected]: A Boolean value that indicates whether the object detects more than one playback route.
//   - [AVRouteDetector.AVRouteDetectorMultipleRoutesDetectedDidChange]: A notification the system posts when changes occur to its detected routes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector
type AVRouteDetector struct {
	objectivec.Object
}

// AVRouteDetectorFromID constructs a [AVRouteDetector] from an objc.ID.
//
// An object that detects available media playback routes.
func AVRouteDetectorFromID(id objc.ID) AVRouteDetector {
	return AVRouteDetector{objectivec.Object{ID: id}}
}
// NOTE: AVRouteDetector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVRouteDetector] class.
//
// # Detecting routes
//
//   - [IAVRouteDetector.RouteDetectionEnabled]: A Boolean value that indicates whether route detection is in an enabled state.
//   - [IAVRouteDetector.SetRouteDetectionEnabled]
//   - [IAVRouteDetector.MultipleRoutesDetected]: A Boolean value that indicates whether the object detects more than one playback route.
//   - [IAVRouteDetector.AVRouteDetectorMultipleRoutesDetectedDidChange]: A notification the system posts when changes occur to its detected routes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector
type IAVRouteDetector interface {
	objectivec.IObject

	// Topic: Detecting routes

	// A Boolean value that indicates whether route detection is in an enabled state.
	RouteDetectionEnabled() bool
	SetRouteDetectionEnabled(value bool)
	// A Boolean value that indicates whether the object detects more than one playback route.
	MultipleRoutesDetected() bool
	// A notification the system posts when changes occur to its detected routes.
	AVRouteDetectorMultipleRoutesDetectedDidChange() foundation.NSString
}

// Init initializes the instance.
func (r AVRouteDetector) Init() AVRouteDetector {
	rv := objc.Send[AVRouteDetector](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r AVRouteDetector) Autorelease() AVRouteDetector {
	rv := objc.Send[AVRouteDetector](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVRouteDetector creates a new AVRouteDetector instance.
func NewAVRouteDetector() AVRouteDetector {
	class := getAVRouteDetectorClass()
	rv := objc.Send[AVRouteDetector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether route detection is in an enabled
// state.
//
// # Discussion
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector/isRouteDetectionEnabled
func (r AVRouteDetector) RouteDetectionEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isRouteDetectionEnabled"))
	return rv
}
func (r AVRouteDetector) SetRouteDetectionEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setRouteDetectionEnabled:"), value)
}
// A Boolean value that indicates whether the object detects more than one
// playback route.
//
// # Discussion
// 
// The system posts a
// [AVRouteDetectorMultipleRoutesDetectedDidChangeNotification] notification
// when this property value changes.
//
// [AVRouteDetectorMultipleRoutesDetectedDidChangeNotification]: https://developer.apple.com/documentation/AVFoundation/AVRouteDetectorMultipleRoutesDetectedDidChangeNotification
//
// See: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector/multipleRoutesDetected
func (r AVRouteDetector) MultipleRoutesDetected() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("multipleRoutesDetected"))
	return rv
}
// A notification the system posts when changes occur to its detected routes.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/AVRouteDetectorMultipleRoutesDetectedDidChange
func (r AVRouteDetector) AVRouteDetectorMultipleRoutesDetectedDidChange() foundation.NSString {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("AVRouteDetectorMultipleRoutesDetectedDidChange"))
	return foundation.NSStringFromID(objc.ID(rv))
}


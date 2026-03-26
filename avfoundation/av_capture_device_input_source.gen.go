// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureDeviceInputSource] class.
var (
	_AVCaptureDeviceInputSourceClass     AVCaptureDeviceInputSourceClass
	_AVCaptureDeviceInputSourceClassOnce sync.Once
)

func getAVCaptureDeviceInputSourceClass() AVCaptureDeviceInputSourceClass {
	_AVCaptureDeviceInputSourceClassOnce.Do(func() {
		_AVCaptureDeviceInputSourceClass = AVCaptureDeviceInputSourceClass{class: objc.GetClass("AVCaptureDeviceInputSource")}
	})
	return _AVCaptureDeviceInputSourceClass
}

// GetAVCaptureDeviceInputSourceClass returns the class object for AVCaptureDeviceInputSource.
func GetAVCaptureDeviceInputSourceClass() AVCaptureDeviceInputSourceClass {
	return getAVCaptureDeviceInputSourceClass()
}

type AVCaptureDeviceInputSourceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureDeviceInputSourceClass) Alloc() AVCaptureDeviceInputSource {
	rv := objc.Send[AVCaptureDeviceInputSource](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A distinct input source on a capture device.
//
// # Overview
// 
// A capture device may optionally present an array of input sources that
// represent distinct mutually exclusive inputs to the device. For example, an
// audio capture device might have ADAT optical and analog input sources; a
// video capture device might have an HDMI or component input source.
//
// # Accessing properties
//
//   - [AVCaptureDeviceInputSource.InputSourceID]: An identifier for an input source.
//   - [AVCaptureDeviceInputSource.LocalizedName]: A localized, human-readable name for the input source.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/InputSource
type AVCaptureDeviceInputSource struct {
	objectivec.Object
}

// AVCaptureDeviceInputSourceFromID constructs a [AVCaptureDeviceInputSource] from an objc.ID.
//
// A distinct input source on a capture device.
func AVCaptureDeviceInputSourceFromID(id objc.ID) AVCaptureDeviceInputSource {
	return AVCaptureDeviceInputSource{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureDeviceInputSource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureDeviceInputSource] class.
//
// # Accessing properties
//
//   - [IAVCaptureDeviceInputSource.InputSourceID]: An identifier for an input source.
//   - [IAVCaptureDeviceInputSource.LocalizedName]: A localized, human-readable name for the input source.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/InputSource
type IAVCaptureDeviceInputSource interface {
	objectivec.IObject

	// Topic: Accessing properties

	// An identifier for an input source.
	InputSourceID() string
	// A localized, human-readable name for the input source.
	LocalizedName() string

	// The currently active input source of the device.
	ActiveInputSource() IAVCaptureDeviceInputSource
	SetActiveInputSource(value IAVCaptureDeviceInputSource)
	// An array of input sources that the device supports.
	InputSources() IAVCaptureDeviceInputSource
	SetInputSources(value IAVCaptureDeviceInputSource)
}

// Init initializes the instance.
func (c AVCaptureDeviceInputSource) Init() AVCaptureDeviceInputSource {
	rv := objc.Send[AVCaptureDeviceInputSource](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureDeviceInputSource) Autorelease() AVCaptureDeviceInputSource {
	rv := objc.Send[AVCaptureDeviceInputSource](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDeviceInputSource creates a new AVCaptureDeviceInputSource instance.
func NewAVCaptureDeviceInputSource() AVCaptureDeviceInputSource {
	class := getAVCaptureDeviceInputSourceClass()
	rv := objc.Send[AVCaptureDeviceInputSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An identifier for an input source.
//
// # Discussion
// 
// The identifier is unique among the input sources exposed by particular
// capture device instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/InputSource/inputSourceID
func (c AVCaptureDeviceInputSource) InputSourceID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputSourceID"))
	return foundation.NSStringFromID(rv).String()
}
// A localized, human-readable name for the input source.
//
// # Discussion
// 
// You can use this property to display the name of the capture device input
// source in a user interface.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/InputSource/localizedName
func (c AVCaptureDeviceInputSource) LocalizedName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}
// The currently active input source of the device.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeinputsource
func (c AVCaptureDeviceInputSource) ActiveInputSource() IAVCaptureDeviceInputSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activeInputSource"))
	return AVCaptureDeviceInputSourceFromID(objc.ID(rv))
}
func (c AVCaptureDeviceInputSource) SetActiveInputSource(value IAVCaptureDeviceInputSource) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveInputSource:"), value)
}
// An array of input sources that the device supports.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/inputsources
func (c AVCaptureDeviceInputSource) InputSources() IAVCaptureDeviceInputSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputSources"))
	return AVCaptureDeviceInputSourceFromID(objc.ID(rv))
}
func (c AVCaptureDeviceInputSource) SetInputSources(value IAVCaptureDeviceInputSource) {
	objc.Send[struct{}](c.ID, objc.Sel("setInputSources:"), value)
}


// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureControl] class.
var (
	_AVCaptureControlClass     AVCaptureControlClass
	_AVCaptureControlClassOnce sync.Once
)

func getAVCaptureControlClass() AVCaptureControlClass {
	_AVCaptureControlClassOnce.Do(func() {
		_AVCaptureControlClass = AVCaptureControlClass{class: objc.GetClass("AVCaptureControl")}
	})
	return _AVCaptureControlClass
}

// GetAVCaptureControlClass returns the class object for AVCaptureControl.
func GetAVCaptureControlClass() AVCaptureControlClass {
	return getAVCaptureControlClass()
}

type AVCaptureControlClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureControlClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureControlClass) Alloc() AVCaptureControl {
	rv := objc.Send[AVCaptureControl](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class for controls that interact with the camera system.
//
// # Overview
// 
// Capture controls provide the interface for interacting with the camera
// system from the Camera Control button on iPhone 16 devices. The framework
// provides several concrete subclasses of this class that allow apps to
// access built-in functionality and define custom controls.
//
// # Setting the enabled state
//
//   - [AVCaptureControl.Enabled]: A Boolean value that indicates whether this control supports user interaction.
//   - [AVCaptureControl.SetEnabled]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureControl
type AVCaptureControl struct {
	objectivec.Object
}

// AVCaptureControlFromID constructs a [AVCaptureControl] from an objc.ID.
//
// An abstract base class for controls that interact with the camera system.
func AVCaptureControlFromID(id objc.ID) AVCaptureControl {
	return AVCaptureControl{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureControl adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureControl] class.
//
// # Setting the enabled state
//
//   - [IAVCaptureControl.Enabled]: A Boolean value that indicates whether this control supports user interaction.
//   - [IAVCaptureControl.SetEnabled]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureControl
type IAVCaptureControl interface {
	objectivec.IObject

	// Topic: Setting the enabled state

	// A Boolean value that indicates whether this control supports user interaction.
	Enabled() bool
	SetEnabled(value bool)
}

// Init initializes the instance.
func (c AVCaptureControl) Init() AVCaptureControl {
	rv := objc.Send[AVCaptureControl](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureControl) Autorelease() AVCaptureControl {
	rv := objc.Send[AVCaptureControl](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureControl creates a new AVCaptureControl instance.
func NewAVCaptureControl() AVCaptureControl {
	class := getAVCaptureControlClass()
	rv := objc.Send[AVCaptureControl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether this control supports user
// interaction.
//
// # Discussion
// 
// Controls support user interaction by default. You can temporarily disable
// user interaction on a control without removing it from a capture session by
// setting it’s enabled state to `false`.
// 
// The default value is `true`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureControl/isEnabled
func (c AVCaptureControl) Enabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEnabled"))
	return rv
}
func (c AVCaptureControl) SetEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabled:"), value)
}


// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
)

// The class instance for the [AVCaptureScreenInput] class.
var (
	_AVCaptureScreenInputClass     AVCaptureScreenInputClass
	_AVCaptureScreenInputClassOnce sync.Once
)

func getAVCaptureScreenInputClass() AVCaptureScreenInputClass {
	_AVCaptureScreenInputClassOnce.Do(func() {
		_AVCaptureScreenInputClass = AVCaptureScreenInputClass{class: objc.GetClass("AVCaptureScreenInput")}
	})
	return _AVCaptureScreenInputClass
}

// GetAVCaptureScreenInputClass returns the class object for AVCaptureScreenInput.
func GetAVCaptureScreenInputClass() AVCaptureScreenInputClass {
	return getAVCaptureScreenInputClass()
}

type AVCaptureScreenInputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureScreenInputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureScreenInputClass) Alloc() AVCaptureScreenInput {
	rv := objc.Send[AVCaptureScreenInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A capture input for recording from a screen in macOS.
//
// # Overview
// 
// This class is a concrete capture input subclass that provides an interface
// to capture media from a screen or a portion of a screen.
// 
// Use instances of this class as input sources for [AVCaptureSession] objects
// that provide media data from one of the screens connected to the system,
// represented by [CGDirectDisplayID].
//
// [CGDirectDisplayID]: https://developer.apple.com/documentation/CoreGraphics/CGDirectDisplayID
//
// # Initializing a capture screen input
//
//   - [AVCaptureScreenInput.InitWithDisplayID]: Initializes a capture screen input that provides media data from the specified display.
//
// # Setting video capture options
//
//   - [AVCaptureScreenInput.MinFrameDuration]: The screen input’s minimum frame duration.
//   - [AVCaptureScreenInput.SetMinFrameDuration]
//   - [AVCaptureScreenInput.CropRect]: Indicates the bounding rectangle of the screen area to be captured, in pixels.
//   - [AVCaptureScreenInput.SetCropRect]
//   - [AVCaptureScreenInput.ScaleFactor]: Indicates the factor by which video buffers captured from the screen are to be scaled.
//   - [AVCaptureScreenInput.SetScaleFactor]
//
// # Capturing mouse activity
//
//   - [AVCaptureScreenInput.CapturesCursor]: A Boolean value that specifies whether the mouse cursor appears in the captured output.
//   - [AVCaptureScreenInput.SetCapturesCursor]
//   - [AVCaptureScreenInput.CapturesMouseClicks]: A Boolean value that specifies whether mouse clicks appear highlighted in the captured output.
//   - [AVCaptureScreenInput.SetCapturesMouseClicks]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput
type AVCaptureScreenInput struct {
	AVCaptureInput
}

// AVCaptureScreenInputFromID constructs a [AVCaptureScreenInput] from an objc.ID.
//
// A capture input for recording from a screen in macOS.
func AVCaptureScreenInputFromID(id objc.ID) AVCaptureScreenInput {
	return AVCaptureScreenInput{AVCaptureInput: AVCaptureInputFromID(id)}
}
// NOTE: AVCaptureScreenInput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureScreenInput] class.
//
// # Initializing a capture screen input
//
//   - [IAVCaptureScreenInput.InitWithDisplayID]: Initializes a capture screen input that provides media data from the specified display.
//
// # Setting video capture options
//
//   - [IAVCaptureScreenInput.MinFrameDuration]: The screen input’s minimum frame duration.
//   - [IAVCaptureScreenInput.SetMinFrameDuration]
//   - [IAVCaptureScreenInput.CropRect]: Indicates the bounding rectangle of the screen area to be captured, in pixels.
//   - [IAVCaptureScreenInput.SetCropRect]
//   - [IAVCaptureScreenInput.ScaleFactor]: Indicates the factor by which video buffers captured from the screen are to be scaled.
//   - [IAVCaptureScreenInput.SetScaleFactor]
//
// # Capturing mouse activity
//
//   - [IAVCaptureScreenInput.CapturesCursor]: A Boolean value that specifies whether the mouse cursor appears in the captured output.
//   - [IAVCaptureScreenInput.SetCapturesCursor]
//   - [IAVCaptureScreenInput.CapturesMouseClicks]: A Boolean value that specifies whether mouse clicks appear highlighted in the captured output.
//   - [IAVCaptureScreenInput.SetCapturesMouseClicks]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput
type IAVCaptureScreenInput interface {
	IAVCaptureInput

	// Topic: Initializing a capture screen input

	// Initializes a capture screen input that provides media data from the specified display.
	InitWithDisplayID(displayID uint32) AVCaptureScreenInput

	// Topic: Setting video capture options

	// The screen input’s minimum frame duration.
	MinFrameDuration() uintptr
	SetMinFrameDuration(value uintptr)
	// Indicates the bounding rectangle of the screen area to be captured, in pixels.
	CropRect() corefoundation.CGRect
	SetCropRect(value corefoundation.CGRect)
	// Indicates the factor by which video buffers captured from the screen are to be scaled.
	ScaleFactor() float64
	SetScaleFactor(value float64)

	// Topic: Capturing mouse activity

	// A Boolean value that specifies whether the mouse cursor appears in the captured output.
	CapturesCursor() bool
	SetCapturesCursor(value bool)
	// A Boolean value that specifies whether mouse clicks appear highlighted in the captured output.
	CapturesMouseClicks() bool
	SetCapturesMouseClicks(value bool)
}

// Init initializes the instance.
func (c AVCaptureScreenInput) Init() AVCaptureScreenInput {
	rv := objc.Send[AVCaptureScreenInput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureScreenInput) Autorelease() AVCaptureScreenInput {
	rv := objc.Send[AVCaptureScreenInput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureScreenInput creates a new AVCaptureScreenInput instance.
func NewAVCaptureScreenInput() AVCaptureScreenInput {
	class := getAVCaptureScreenInputClass()
	rv := objc.Send[AVCaptureScreenInput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a capture screen input that provides media data from the
// specified display.
//
// displayID: The ID of the display from which to capture video.
// 
// [CGDirectDisplayID] is defined in ``.
//
// # Return Value
// 
// A capture screen input initialized to provide media data from a given
// display. If the display cannot be used (because it is not available on the
// system, for example), returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/init(displayID:)
func NewCaptureScreenInputWithDisplayID(displayID uint32) AVCaptureScreenInput {
	instance := getAVCaptureScreenInputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplayID:"), displayID)
	return AVCaptureScreenInputFromID(rv)
}

// Initializes a capture screen input that provides media data from the
// specified display.
//
// displayID: The ID of the display from which to capture video.
// 
// [CGDirectDisplayID] is defined in ``.
//
// # Return Value
// 
// A capture screen input initialized to provide media data from a given
// display. If the display cannot be used (because it is not available on the
// system, for example), returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/init(displayID:)
func (c AVCaptureScreenInput) InitWithDisplayID(displayID uint32) AVCaptureScreenInput {
	rv := objc.Send[AVCaptureScreenInput](c.ID, objc.Sel("initWithDisplayID:"), displayID)
	return rv
}

// The screen input’s minimum frame duration.
//
// # Discussion
// 
// The `minFrameDuration` is the reciprocal of its maximum frame rate.
// 
// You use this property to request a maximum frame rate at which the input
// produces video frames. The requested rate may not be achievable due to
// overall bandwidth, so actual frame rates may be lower.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/minFrameDuration
func (c AVCaptureScreenInput) MinFrameDuration() uintptr {
	rv := objc.Send[uintptr](c.ID, objc.Sel("minFrameDuration"))
	return rv
}
func (c AVCaptureScreenInput) SetMinFrameDuration(value uintptr) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinFrameDuration:"), value)
}
// Indicates the bounding rectangle of the screen area to be captured, in
// pixels.
//
// # Discussion
// 
// By default, [AVCaptureScreenInput] captures the entire area of the
// displayID with which it is associated.
// 
// Set the value of this property to limit the capture rectangle to a
// subsection of the screen.
// 
// The rectangle should define a smaller section of the screen in the
// screen’s coordinate system. The origin (0,0) is the bottom-left corner of
// the screen.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/cropRect
func (c AVCaptureScreenInput) CropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("cropRect"))
	return corefoundation.CGRect(rv)
}
func (c AVCaptureScreenInput) SetCropRect(value corefoundation.CGRect) {
	objc.Send[struct{}](c.ID, objc.Sel("setCropRect:"), value)
}
// Indicates the factor by which video buffers captured from the screen are to
// be scaled.
//
// # Discussion
// 
// By default, [AVCaptureScreenInput] captures the video buffers from the
// display at a scale factor of 1.0 (no scaling). Set this property to scale
// the buffers by a given factor; for example a 320x240 capture area with a
// scaleFactor of `2.0` produces video buffers at 640x480.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/scaleFactor
func (c AVCaptureScreenInput) ScaleFactor() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scaleFactor"))
	return rv
}
func (c AVCaptureScreenInput) SetScaleFactor(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setScaleFactor:"), value)
}
// A Boolean value that specifies whether the mouse cursor appears in the
// captured output.
//
// # Discussion
// 
// When this property is true (the default), captured video frames include the
// mouse pointer. If you change this property to false, the captured output
// contains only the windows on the screen (that is, the mouse pointer is
// invisible in captured video).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/capturesCursor
func (c AVCaptureScreenInput) CapturesCursor() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("capturesCursor"))
	return rv
}
func (c AVCaptureScreenInput) SetCapturesCursor(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCapturesCursor:"), value)
}
// A Boolean value that specifies whether mouse clicks appear highlighted in
// the captured output.
//
// # Discussion
// 
// By default, [AVCaptureScreenInput] does not highlight mouse clicks in its
// captured output.
// 
// If you set this property is set to [true], mouse clicks are highlighted (a
// circle is drawn around the mouse for the duration of the click) in the
// captured output.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput/capturesMouseClicks
func (c AVCaptureScreenInput) CapturesMouseClicks() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("capturesMouseClicks"))
	return rv
}
func (c AVCaptureScreenInput) SetCapturesMouseClicks(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCapturesMouseClicks:"), value)
}


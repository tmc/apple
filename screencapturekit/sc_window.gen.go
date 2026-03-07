// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCWindow] class.
var (
	_SCWindowClass     SCWindowClass
	_SCWindowClassOnce sync.Once
)

func getSCWindowClass() SCWindowClass {
	_SCWindowClassOnce.Do(func() {
		_SCWindowClass = SCWindowClass{class: objc.GetClass("SCWindow")}
	})
	return _SCWindowClass
}

// GetSCWindowClass returns the class object for SCWindow.
func GetSCWindowClass() SCWindowClass {
	return getSCWindowClass()
}

type SCWindowClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCWindowClass) Alloc() SCWindow {
	rv := objc.Send[SCWindow](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}







// An instance that represents an onscreen window.
//
// # Overview
// 
// Retrieve the available windows from an instance of [SCShareableContent].
// Select one or more windows to capture and use them to create an instance of
// [SCContentFilter]. Apply the filter to an instance of [SCStream] to limit
// its output to content matching your criteria.
//
// # Identifying windows
//
//   - [SCWindow.WindowID]: The Core Graphics window identifier.
//   - [SCWindow.Title]: The string that displays in a window’s title bar.
//   - [SCWindow.OwningApplication]: The app that owns the window.
//   - [SCWindow.WindowLayer]: The layer of the window relative to other windows.
//
// # Accessing dimensions
//
//   - [SCWindow.Frame]: A rectangle the represents the frame of the window within a display.
//
// # Determining visibility
//
//   - [SCWindow.OnScreen]: A Boolean value that indicates whether the window is on screen.
//   - [SCWindow.Active]: A Boolean value that indicates if the window is currently streaming.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow
type SCWindow struct {
	objectivec.Object
}

// SCWindowFromID constructs a [SCWindow] from an objc.ID.
//
// An instance that represents an onscreen window.
func SCWindowFromID(id objc.ID) SCWindow {
	return SCWindow{objectivec.Object{id}}
}
// NOTE: SCWindow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [SCWindow] class.
//
// # Identifying windows
//
//   - [ISCWindow.WindowID]: The Core Graphics window identifier.
//   - [ISCWindow.Title]: The string that displays in a window’s title bar.
//   - [ISCWindow.OwningApplication]: The app that owns the window.
//   - [ISCWindow.WindowLayer]: The layer of the window relative to other windows.
//
// # Accessing dimensions
//
//   - [ISCWindow.Frame]: A rectangle the represents the frame of the window within a display.
//
// # Determining visibility
//
//   - [ISCWindow.OnScreen]: A Boolean value that indicates whether the window is on screen.
//   - [ISCWindow.Active]: A Boolean value that indicates if the window is currently streaming.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow
type ISCWindow interface {
	objectivec.IObject

	// Topic: Identifying windows

	// The Core Graphics window identifier.
	WindowID() uint32
	// The string that displays in a window’s title bar.
	Title() string
	// The app that owns the window.
	OwningApplication() ISCRunningApplication
	// The layer of the window relative to other windows.
	WindowLayer() int

	// Topic: Accessing dimensions

	// A rectangle the represents the frame of the window within a display.
	Frame() corefoundation.CGRect

	// Topic: Determining visibility

	// A Boolean value that indicates whether the window is on screen.
	OnScreen() bool
	// A Boolean value that indicates if the window is currently streaming.
	Active() bool
}





// Init initializes the instance.
func (w SCWindow) Init() SCWindow {
	rv := objc.Send[SCWindow](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w SCWindow) Autorelease() SCWindow {
	rv := objc.Send[SCWindow](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCWindow creates a new SCWindow instance.
func NewSCWindow() SCWindow {
	class := getSCWindowClass()
	rv := objc.Send[SCWindow](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The Core Graphics window identifier.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/windowID
func (w SCWindow) WindowID() uint32 {
	rv := objc.Send[uint32](w.ID, objc.Sel("windowID"))
	return rv
}



// The string that displays in a window’s title bar.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/title
func (w SCWindow) Title() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}



// The app that owns the window.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/owningApplication
func (w SCWindow) OwningApplication() ISCRunningApplication {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("owningApplication"))
	return SCRunningApplicationFromID(objc.ID(rv))
}



// The layer of the window relative to other windows.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/windowLayer
func (w SCWindow) WindowLayer() int {
	rv := objc.Send[int](w.ID, objc.Sel("windowLayer"))
	return rv
}



// A rectangle the represents the frame of the window within a display.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/frame
func (w SCWindow) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}



// A Boolean value that indicates whether the window is on screen.
//
// # Discussion
// 
// This value represents the macOS window server’s onscreen status of the
// window.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/isOnScreen
func (w SCWindow) OnScreen() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isOnScreen"))
	return rv
}



// A Boolean value that indicates if the window is currently streaming.
//
// # Discussion
// 
// When this value is `true`, the window is currently streaming, even if
// offscreen.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/isActive
func (w SCWindow) Active() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isActive"))
	return rv
}
























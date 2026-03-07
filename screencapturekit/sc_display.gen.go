// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCDisplay] class.
var (
	_SCDisplayClass     SCDisplayClass
	_SCDisplayClassOnce sync.Once
)

func getSCDisplayClass() SCDisplayClass {
	_SCDisplayClassOnce.Do(func() {
		_SCDisplayClass = SCDisplayClass{class: objc.GetClass("SCDisplay")}
	})
	return _SCDisplayClass
}

// GetSCDisplayClass returns the class object for SCDisplay.
func GetSCDisplayClass() SCDisplayClass {
	return getSCDisplayClass()
}

type SCDisplayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCDisplayClass) Alloc() SCDisplay {
	rv := objc.Send[SCDisplay](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}







// An instance that represents a display device.
//
// # Overview
// 
// A display object represents a physical display connected to a Mac. Query
// the display to retrieve its unique identifier and onscreen coordinates.
// 
// Retrieve the available displays from an instance of [SCShareableContent].
// Select a display to capture and use it to create an instance of
// [SCContentFilter]. Apply the filter to an instance of [SCStream] to limit
// its output to content matching your criteria.
//
// # Identifying displays
//
//   - [SCDisplay.DisplayID]: The Core Graphics display identifier.
//
// # Accessing dimensions
//
//   - [SCDisplay.Frame]: The frame of the display.
//   - [SCDisplay.Width]: The width of the display in points.
//   - [SCDisplay.Height]: The height of the display in points.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay
type SCDisplay struct {
	objectivec.Object
}

// SCDisplayFromID constructs a [SCDisplay] from an objc.ID.
//
// An instance that represents a display device.
func SCDisplayFromID(id objc.ID) SCDisplay {
	return SCDisplay{objectivec.Object{id}}
}
// NOTE: SCDisplay adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [SCDisplay] class.
//
// # Identifying displays
//
//   - [ISCDisplay.DisplayID]: The Core Graphics display identifier.
//
// # Accessing dimensions
//
//   - [ISCDisplay.Frame]: The frame of the display.
//   - [ISCDisplay.Width]: The width of the display in points.
//   - [ISCDisplay.Height]: The height of the display in points.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay
type ISCDisplay interface {
	objectivec.IObject

	// Topic: Identifying displays

	// The Core Graphics display identifier.
	DisplayID() uint32

	// Topic: Accessing dimensions

	// The frame of the display.
	Frame() corefoundation.CGRect
	// The width of the display in points.
	Width() int
	// The height of the display in points.
	Height() int
}





// Init initializes the instance.
func (d SCDisplay) Init() SCDisplay {
	rv := objc.Send[SCDisplay](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d SCDisplay) Autorelease() SCDisplay {
	rv := objc.Send[SCDisplay](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCDisplay creates a new SCDisplay instance.
func NewSCDisplay() SCDisplay {
	class := getSCDisplayClass()
	rv := objc.Send[SCDisplay](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The Core Graphics display identifier.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay/displayID
func (d SCDisplay) DisplayID() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("displayID"))
	return rv
}



// The frame of the display.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay/frame
func (d SCDisplay) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](d.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}



// The width of the display in points.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay/width
func (d SCDisplay) Width() int {
	rv := objc.Send[int](d.ID, objc.Sel("width"))
	return rv
}



// The height of the display in points.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay/height
func (d SCDisplay) Height() int {
	rv := objc.Send[int](d.ID, objc.Sel("height"))
	return rv
}
























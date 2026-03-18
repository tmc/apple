// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSStatusBar] class.
var (
	_NSStatusBarClass     NSStatusBarClass
	_NSStatusBarClassOnce sync.Once
)

func getNSStatusBarClass() NSStatusBarClass {
	_NSStatusBarClassOnce.Do(func() {
		_NSStatusBarClass = NSStatusBarClass{class: objc.GetClass("NSStatusBar")}
	})
	return _NSStatusBarClass
}

// GetNSStatusBarClass returns the class object for NSStatusBar.
func GetNSStatusBarClass() NSStatusBarClass {
	return getNSStatusBarClass()
}

type NSStatusBarClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStatusBarClass) Alloc() NSStatusBar {
	rv := objc.Send[NSStatusBar](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages a collection of status items displayed within the
// system-wide menu bar.
//
// # Overview
// 
// A status item (an instance of [NSStatusItem]) can be displayed with text or
// an icon, can provide a menu and a target-action message when clicked, or
// can be a fully customized view that you create. Use status items sparingly
// and only if the alternatives (such as a Dock menu, preference pane, or
// status window) are not suitable. Because there is limited space in which to
// display status items, status items are not guaranteed to be available at
// all times. For this reason, do not rely on them being available and always
// provide a user preference for hiding your application’s status items to
// free up space in the menu bar.
//
// # Managing Status items
//
//   - [NSStatusBar.StatusItemWithLength]: Returns a newly created status item that has been allotted a specified space within the status bar.
//   - [NSStatusBar.RemoveStatusItem]: Removes the specified status item from the receiver.
//
// # Getting Status-Bar Attributes
//
//   - [NSStatusBar.Vertical]: A Boolean value indicating whether the status bar has a vertical orientation.
//   - [NSStatusBar.Thickness]: The thickness of the status bar, in pixels.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusBar
type NSStatusBar struct {
	objectivec.Object
}

// NSStatusBarFromID constructs a [NSStatusBar] from an objc.ID.
//
// An object that manages a collection of status items displayed within the
// system-wide menu bar.
func NSStatusBarFromID(id objc.ID) NSStatusBar {
	return NSStatusBar{objectivec.Object{ID: id}}
}
// NOTE: NSStatusBar adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSStatusBar] class.
//
// # Managing Status items
//
//   - [INSStatusBar.StatusItemWithLength]: Returns a newly created status item that has been allotted a specified space within the status bar.
//   - [INSStatusBar.RemoveStatusItem]: Removes the specified status item from the receiver.
//
// # Getting Status-Bar Attributes
//
//   - [INSStatusBar.Vertical]: A Boolean value indicating whether the status bar has a vertical orientation.
//   - [INSStatusBar.Thickness]: The thickness of the status bar, in pixels.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusBar
type INSStatusBar interface {
	objectivec.IObject

	// Topic: Managing Status items

	// Returns a newly created status item that has been allotted a specified space within the status bar.
	StatusItemWithLength(length float64) INSStatusItem
	// Removes the specified status item from the receiver.
	RemoveStatusItem(item INSStatusItem)

	// Topic: Getting Status-Bar Attributes

	// A Boolean value indicating whether the status bar has a vertical orientation.
	Vertical() bool
	// The thickness of the status bar, in pixels.
	Thickness() float64
}

// Init initializes the instance.
func (s NSStatusBar) Init() NSStatusBar {
	rv := objc.Send[NSStatusBar](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStatusBar) Autorelease() NSStatusBar {
	rv := objc.Send[NSStatusBar](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStatusBar creates a new NSStatusBar instance.
func NewNSStatusBar() NSStatusBar {
	class := getNSStatusBarClass()
	rv := objc.Send[NSStatusBar](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a newly created status item that has been allotted a specified
// space within the status bar.
//
// length: A constant that specifies whether the status item is of fixed width, or
// variable width. The valid constants are described in [Status Bar Item
// Length].
// //
// [Status Bar Item Length]: https://developer.apple.com/documentation/AppKit/status-bar-item-length
//
// # Return Value
// 
// An [NSStatusItem] object.
//
// # Discussion
// 
// The receiver does not retain a reference to the status item, so you need to
// retain it. Otherwise, the object is removed from the status bar when it is
// deallocated.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusBar/statusItem(withLength:)
func (s NSStatusBar) StatusItemWithLength(length float64) INSStatusItem {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("statusItemWithLength:"), length)
	return NSStatusItemFromID(rv)
}

// Removes the specified status item from the receiver.
//
// item: The [NSStatusItem] object to remove.
//
// # Discussion
// 
// Status items to the left of the specified one in the status bar shift to
// the right to reclaim its space.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusBar/removeStatusItem(_:)
func (s NSStatusBar) RemoveStatusItem(item INSStatusItem) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeStatusItem:"), item)
}

// A Boolean value indicating whether the status bar has a vertical
// orientation.
//
// # Discussion
// 
// When the value of this property is [true], the status bar has a vertical
// orientation. The status bar returned by the [SystemStatusBar] method is
// horizontal and has the value [false] for this property.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusBar/isVertical
func (s NSStatusBar) Vertical() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isVertical"))
	return rv
}

// The thickness of the status bar, in pixels.
//
// # Discussion
// 
// The default value of this property is `20.0`. The status bar returned by
// the [SystemStatusBar] has a thickness of 22 pixels, which corresponds to
// the thickness of the menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusBar/thickness
func (s NSStatusBar) Thickness() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("thickness"))
	return rv
}

// Returns the system-wide status bar located in the menu bar.
//
// # Return Value
// 
// The shared status bar object.
// 
// # Discussion
// 
// The status bar begins at the right side of the menu bar (to the left of
// Menu Extras and the menu bar clock) and grows to the left as [NSStatusItem]
// objects are added to it.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusBar/system
func (_NSStatusBarClass NSStatusBarClass) SystemStatusBar() NSStatusBar {
	rv := objc.Send[objc.ID](objc.ID(_NSStatusBarClass.class), objc.Sel("systemStatusBar"))
	return NSStatusBarFromID(objc.ID(rv))
}


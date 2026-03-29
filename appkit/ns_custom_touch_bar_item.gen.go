// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSCustomTouchBarItem] class.
var (
	_NSCustomTouchBarItemClass     NSCustomTouchBarItemClass
	_NSCustomTouchBarItemClassOnce sync.Once
)

func getNSCustomTouchBarItemClass() NSCustomTouchBarItemClass {
	_NSCustomTouchBarItemClassOnce.Do(func() {
		_NSCustomTouchBarItemClass = NSCustomTouchBarItemClass{class: objc.GetClass("NSCustomTouchBarItem")}
	})
	return _NSCustomTouchBarItemClass
}

// GetNSCustomTouchBarItemClass returns the class object for NSCustomTouchBarItem.
func GetNSCustomTouchBarItemClass() NSCustomTouchBarItemClass {
	return getNSCustomTouchBarItemClass()
}

type NSCustomTouchBarItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCustomTouchBarItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCustomTouchBarItemClass) Alloc() NSCustomTouchBarItem {
	rv := objc.Send[NSCustomTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A bar item that contains a responder of your choice, such as a view, a
// button, or a scrubber.
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem
type NSCustomTouchBarItem struct {
	NSTouchBarItem
}

// NSCustomTouchBarItemFromID constructs a [NSCustomTouchBarItem] from an objc.ID.
//
// A bar item that contains a responder of your choice, such as a view, a
// button, or a scrubber.
func NSCustomTouchBarItemFromID(id objc.ID) NSCustomTouchBarItem {
	return NSCustomTouchBarItem{NSTouchBarItem: NSTouchBarItemFromID(id)}
}
// NOTE: NSCustomTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCustomTouchBarItem] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem
type INSCustomTouchBarItem interface {
	INSTouchBarItem
}

// Init initializes the instance.
func (c NSCustomTouchBarItem) Init() NSCustomTouchBarItem {
	rv := objc.Send[NSCustomTouchBarItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCustomTouchBarItem) Autorelease() NSCustomTouchBarItem {
	rv := objc.Send[NSCustomTouchBarItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCustomTouchBarItem creates a new NSCustomTouchBarItem instance.
func NewNSCustomTouchBarItem() NSCustomTouchBarItem {
	class := getNSCustomTouchBarItemClass()
	rv := objc.Send[NSCustomTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewCustomTouchBarItemWithCoder(coder foundation.INSCoder) NSCustomTouchBarItem {
	instance := getNSCustomTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCustomTouchBarItemFromID(rv)
}

// Creates a new item with the specified identifier.
//
// # Discussion
// 
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewCustomTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSCustomTouchBarItem {
	instance := getNSCustomTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSCustomTouchBarItemFromID(rv)
}


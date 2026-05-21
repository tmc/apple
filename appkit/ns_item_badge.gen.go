// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSItemBadge] class.
var (
	_NSItemBadgeClass     NSItemBadgeClass
	_NSItemBadgeClassOnce sync.Once
)

func getNSItemBadgeClass() NSItemBadgeClass {
	_NSItemBadgeClassOnce.Do(func() {
		_NSItemBadgeClass = NSItemBadgeClass{class: objc.GetClass("NSItemBadge")}
	})
	return _NSItemBadgeClass
}

// GetNSItemBadgeClass returns the class object for NSItemBadge.
func GetNSItemBadgeClass() NSItemBadgeClass {
	return getNSItemBadgeClass()
}

type NSItemBadgeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSItemBadgeClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSItemBadgeClass) Alloc() NSItemBadge {
	rv := objc.Send[NSItemBadge](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// [NSItemBadge] represents a badge that can be attached to an
// [NSToolbarItem].
//
// # Overview
//
// This badge provides a way to display small visual indicators, such as
// counts and text labels, within a toolbar item. Badges can be used to
// highlight important information, such as unread notifications or status
// indicators.
//
// # Instance Properties
//
//   - [NSItemBadge.Text]: The text to be displayed within the badge.
//
// See: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class
type NSItemBadge struct {
	objectivec.Object
}

// NSItemBadgeFromID constructs a [NSItemBadge] from an objc.ID.
//
// [NSItemBadge] represents a badge that can be attached to an
// [NSToolbarItem].
func NSItemBadgeFromID(id objc.ID) NSItemBadge {
	return NSItemBadge{objectivec.Object{ID: id}}
}

// Ensure NSItemBadge implements INSItemBadge.
var _ INSItemBadge = NSItemBadge{}

// An interface definition for the [NSItemBadge] class.
//
// # Instance Properties
//
//   - [INSItemBadge.Text]: The text to be displayed within the badge.
//
// See: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class
type INSItemBadge interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The text to be displayed within the badge.
	Text() string
}

// Init initializes the instance.
func (i NSItemBadge) Init() NSItemBadge {
	rv := objc.Send[NSItemBadge](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSItemBadge) Autorelease() NSItemBadge {
	rv := objc.Send[NSItemBadge](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSItemBadge creates a new NSItemBadge instance.
func NewNSItemBadge() NSItemBadge {
	class := getNSItemBadgeClass()
	rv := objc.Send[NSItemBadge](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a badge displaying a localized numerical count.
//
// count: The integer value to localize and display in the badge.
//
// # Return Value
//
// A new NSItemBadge instance with the localized specified count.
//
// See: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/badgeWithCount:
func (_NSItemBadgeClass NSItemBadgeClass) BadgeWithCount(count int) NSItemBadge {
	rv := objc.Send[objc.ID](objc.ID(_NSItemBadgeClass.class), objc.Sel("badgeWithCount:"), count)
	return NSItemBadgeFromID(rv)
}

// Creates a badge displaying a text.
//
// text: The text to be displayed inside the badge.
//
// # Return Value
//
// A new [NSItemBadge] instance with the specified text.
//
// See: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/badgeWithText:
func (_NSItemBadgeClass NSItemBadgeClass) BadgeWithText(text string) NSItemBadge {
	rv := objc.Send[objc.ID](objc.ID(_NSItemBadgeClass.class), objc.Sel("badgeWithText:"), objc.String(text))
	return NSItemBadgeFromID(rv)
}

// Creates a badge styled as an indicator. In this context, an indicator is
// simply a badge without any text.
//
// # Return Value
//
// A new [NSItemBadge] instance styled as an indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/indicatorBadge
func (_NSItemBadgeClass NSItemBadgeClass) IndicatorBadge() NSItemBadge {
	rv := objc.Send[objc.ID](objc.ID(_NSItemBadgeClass.class), objc.Sel("indicatorBadge"))
	return NSItemBadgeFromID(rv)
}

// The text to be displayed within the badge.
//
// See: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/text
func (i NSItemBadge) Text() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("text"))
	return foundation.NSStringFromID(rv).String()
}

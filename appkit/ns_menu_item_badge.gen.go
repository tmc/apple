// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMenuItemBadge] class.
var (
	_NSMenuItemBadgeClass     NSMenuItemBadgeClass
	_NSMenuItemBadgeClassOnce sync.Once
)

func getNSMenuItemBadgeClass() NSMenuItemBadgeClass {
	_NSMenuItemBadgeClassOnce.Do(func() {
		_NSMenuItemBadgeClass = NSMenuItemBadgeClass{class: objc.GetClass("NSMenuItemBadge")}
	})
	return _NSMenuItemBadgeClass
}

// GetNSMenuItemBadgeClass returns the class object for NSMenuItemBadge.
func GetNSMenuItemBadgeClass() NSMenuItemBadgeClass {
	return getNSMenuItemBadgeClass()
}

type NSMenuItemBadgeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMenuItemBadgeClass) Alloc() NSMenuItemBadge {
	rv := objc.Send[NSMenuItemBadge](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A control that provides additional quantitative information specific to a
// menu item, such as the number of available updates.
//
// # Overview
// 
// You create a badge using an initializer or a predefined factory method, and
// then you assign it to the [NSMenuItemBadge.Badge] property of a [NSMenuItem] for display.
// 
// [media-4304515]
// 
// For example, to display a badge with a count, use the [NSMenuItemBadge.InitWithCount]
// initalizer, passing in the value of `count` as an [Int].
// 
// To display a badge with a custom string, use the [NSMenuItemBadge.InitWithString]
// initializer, passing in the string you want to display.
// 
// To display a badge using a predefined [NSMenuItemBadge.BadgeType], use a
// factory method such as [NSMenuItemBadge.NewItemsWithCount], passing in the `count` of the
// badge to display.
// 
// The default value of this property is `nil`.
//
// [NSMenuItemBadge.BadgeType]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/BadgeType
//
// # Creating menu item badges
//
//   - [NSMenuItemBadge.InitWithCount]: Creates a badge with a count and an empty string.
//   - [NSMenuItemBadge.InitWithString]: Creates a badge with the provided custom string.
//
// # Accessing menu item badge attributes
//
//   - [NSMenuItemBadge.ItemCount]: The number of items the badge displays.
//   - [NSMenuItemBadge.StringValue]: The string representation of the badge when it displays.
//   - [NSMenuItemBadge.Type]: The type of items the badge displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge
type NSMenuItemBadge struct {
	objectivec.Object
}

// NSMenuItemBadgeFromID constructs a [NSMenuItemBadge] from an objc.ID.
//
// A control that provides additional quantitative information specific to a
// menu item, such as the number of available updates.
func NSMenuItemBadgeFromID(id objc.ID) NSMenuItemBadge {
	return NSMenuItemBadge{objectivec.Object{ID: id}}
}
// NOTE: NSMenuItemBadge adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMenuItemBadge] class.
//
// # Creating menu item badges
//
//   - [INSMenuItemBadge.InitWithCount]: Creates a badge with a count and an empty string.
//   - [INSMenuItemBadge.InitWithString]: Creates a badge with the provided custom string.
//
// # Accessing menu item badge attributes
//
//   - [INSMenuItemBadge.ItemCount]: The number of items the badge displays.
//   - [INSMenuItemBadge.StringValue]: The string representation of the badge when it displays.
//   - [INSMenuItemBadge.Type]: The type of items the badge displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge
type INSMenuItemBadge interface {
	objectivec.IObject

	// Topic: Creating menu item badges

	// Creates a badge with a count and an empty string.
	InitWithCount(itemCount int) NSMenuItemBadge
	// Creates a badge with the provided custom string.
	InitWithString(string_ string) NSMenuItemBadge

	// Topic: Accessing menu item badge attributes

	// The number of items the badge displays.
	ItemCount() int
	// The string representation of the badge when it displays.
	StringValue() string
	// The type of items the badge displays.
	Type() NSMenuItemBadgeType

	Badge() INSMenuItemBadge
	SetBadge(value INSMenuItemBadge)
}

// Init initializes the instance.
func (m NSMenuItemBadge) Init() NSMenuItemBadge {
	rv := objc.Send[NSMenuItemBadge](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMenuItemBadge) Autorelease() NSMenuItemBadge {
	rv := objc.Send[NSMenuItemBadge](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMenuItemBadge creates a new NSMenuItemBadge instance.
func NewNSMenuItemBadge() NSMenuItemBadge {
	class := getNSMenuItemBadgeClass()
	rv := objc.Send[NSMenuItemBadge](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a badge with a count and an empty string.
//
// itemCount: The badge count.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/init(count:)
func NewMenuItemBadgeWithCount(itemCount int) NSMenuItemBadge {
	instance := getNSMenuItemBadgeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCount:"), itemCount)
	return NSMenuItemBadgeFromID(rv)
}

// Creates a badge with the provided custom string.
//
// string: The string label that displays when the badge appears.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/init(string:)
func NewMenuItemBadgeWithString(string_ string) NSMenuItemBadge {
	instance := getNSMenuItemBadgeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(string_))
	return NSMenuItemBadgeFromID(rv)
}

// Creates a badge with a count and an empty string.
//
// itemCount: The badge count.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/init(count:)
func (m NSMenuItemBadge) InitWithCount(itemCount int) NSMenuItemBadge {
	rv := objc.Send[NSMenuItemBadge](m.ID, objc.Sel("initWithCount:"), itemCount)
	return rv
}

// Creates a badge with the provided custom string.
//
// string: The string label that displays when the badge appears.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/init(string:)
func (m NSMenuItemBadge) InitWithString(string_ string) NSMenuItemBadge {
	rv := objc.Send[NSMenuItemBadge](m.ID, objc.Sel("initWithString:"), objc.String(string_))
	return rv
}

// Creates an alert-style badge with an integer count and a predefined label
// that represents the number of alerts.
//
// itemCount: The badge count.
//
// # Return Value
// 
// Returns a new badge item of type [NSMenuItemBadge.BadgeType.alerts].
//
// [NSMenuItemBadge.BadgeType.alerts]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/BadgeType/alerts
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/alerts(count:)
func (_NSMenuItemBadgeClass NSMenuItemBadgeClass) AlertsWithCount(itemCount int) NSMenuItemBadge {
	rv := objc.Send[objc.ID](objc.ID(_NSMenuItemBadgeClass.class), objc.Sel("alertsWithCount:"), itemCount)
	return NSMenuItemBadgeFromID(rv)
}

// Creates a new item-style badge with an integer count and a predefined label
// that represents the number of new items.
//
// itemCount: The badge count.
//
// # Return Value
// 
// Returns a new badge item of type [NSMenuItemBadge.BadgeType.newItems].
//
// [NSMenuItemBadge.BadgeType.newItems]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/BadgeType/newItems
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/newItems(count:)
func (_NSMenuItemBadgeClass NSMenuItemBadgeClass) NewItemsWithCount(itemCount int) NSMenuItemBadge {
	rv := objc.Send[objc.ID](objc.ID(_NSMenuItemBadgeClass.class), objc.Sel("newItemsWithCount:"), itemCount)
	return NSMenuItemBadgeFromID(rv)
}

// Creates an update-style badge with an integer count and a predefined label
// that represents the number of available updates.
//
// itemCount: The badge count.
//
// # Return Value
// 
// Returns a new badge item of type [NSMenuItemBadge.BadgeType.updates].
//
// [NSMenuItemBadge.BadgeType.updates]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/BadgeType/updates
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/updates(count:)
func (_NSMenuItemBadgeClass NSMenuItemBadgeClass) UpdatesWithCount(itemCount int) NSMenuItemBadge {
	rv := objc.Send[objc.ID](objc.ID(_NSMenuItemBadgeClass.class), objc.Sel("updatesWithCount:"), itemCount)
	return NSMenuItemBadgeFromID(rv)
}

// The number of items the badge displays.
//
// # Discussion
// 
// If you create a badge with a custom string, this value is `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/itemCount
func (m NSMenuItemBadge) ItemCount() int {
	rv := objc.Send[int](m.ID, objc.Sel("itemCount"))
	return rv
}

// The string representation of the badge when it displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/stringValue-fc9f
func (m NSMenuItemBadge) StringValue() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stringValue"))
	return foundation.NSStringFromID(rv).String()
}

// The type of items the badge displays.
//
// # Discussion
// 
// If you create a badge with a custom string, this value is
// [NSMenuItemBadge.BadgeType.none].
//
// [NSMenuItemBadge.BadgeType.none]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/BadgeType/none
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/type
func (m NSMenuItemBadge) Type() NSMenuItemBadgeType {
	rv := objc.Send[NSMenuItemBadgeType](m.ID, objc.Sel("type"))
	return NSMenuItemBadgeType(rv)
}

// See: https://developer.apple.com/documentation/appkit/nsmenuitem/badge
func (m NSMenuItemBadge) Badge() INSMenuItemBadge {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("badge"))
	return NSMenuItemBadgeFromID(objc.ID(rv))
}
func (m NSMenuItemBadge) SetBadge(value INSMenuItemBadge) {
	objc.Send[struct{}](m.ID, objc.Sel("setBadge:"), value)
}


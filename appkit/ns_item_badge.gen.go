// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	return NSItemBadge{objectivec.Object{id}}
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

	// A Boolean value that indicates whether the toolbar item has a bordered style.
	IsBordered() bool
	SetIsBordered(value bool)
	// A Boolean value that indicates whether the item is enabled.
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
	IsNavigational() bool
	SetIsNavigational(value bool)
	// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
	IsVisible() bool
	SetIsVisible(value bool)
	// Defines the toolbar item’s appearance. The default style is plain.
	Style() NSToolbarItemStyle
	SetStyle(value NSToolbarItemStyle)
	// An integer tag you can use to identify the toolbar item.
	Tag() int
	SetTag(value int)
	// The display priority associated with the toolbar item.
	VisibilityPriority() NSToolbarItemVisibilityPriority
	SetVisibilityPriority(value NSToolbarItemVisibilityPriority)
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



// A Boolean value that indicates whether the toolbar item has a bordered
// style.
//
// See: https://developer.apple.com/documentation/appkit/nstoolbaritem/isbordered
func (i NSItemBadge) IsBordered() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("bordered"))
	return rv
}
func (i NSItemBadge) SetIsBordered(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setBordered:"), value)
}



// A Boolean value that indicates whether the item is enabled.
//
// See: https://developer.apple.com/documentation/appkit/nstoolbaritem/isenabled
func (i NSItemBadge) IsEnabled() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("enabled"))
	return rv
}
func (i NSItemBadge) SetIsEnabled(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setEnabled:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nstoolbaritem/ishidden
func (i NSItemBadge) IsHidden() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("hidden"))
	return rv
}
func (i NSItemBadge) SetIsHidden(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setHidden:"), value)
}



// A Boolean value that indicates whether the item behaves as a navigation
// item in the toolbar.
//
// See: https://developer.apple.com/documentation/appkit/nstoolbaritem/isnavigational
func (i NSItemBadge) IsNavigational() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("navigational"))
	return rv
}
func (i NSItemBadge) SetIsNavigational(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setNavigational:"), value)
}



// A Boolean value that indicates whether the item is currently visible in the
// toolbar, and not in the overflow menu.
//
// See: https://developer.apple.com/documentation/appkit/nstoolbaritem/isvisible
func (i NSItemBadge) IsVisible() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("visible"))
	return rv
}
func (i NSItemBadge) SetIsVisible(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setVisible:"), value)
}



// Defines the toolbar item’s appearance. The default style is plain.
//
// See: https://developer.apple.com/documentation/appkit/nstoolbaritem/style-swift.property
func (i NSItemBadge) Style() NSToolbarItemStyle {
	rv := objc.Send[NSToolbarItemStyle](i.ID, objc.Sel("style"))
	return NSToolbarItemStyle(rv)
}
func (i NSItemBadge) SetStyle(value NSToolbarItemStyle) {
	objc.Send[struct{}](i.ID, objc.Sel("setStyle:"), value)
}



// An integer tag you can use to identify the toolbar item.
//
// See: https://developer.apple.com/documentation/appkit/nstoolbaritem/tag
func (i NSItemBadge) Tag() int {
	rv := objc.Send[int](i.ID, objc.Sel("tag"))
	return rv
}
func (i NSItemBadge) SetTag(value int) {
	objc.Send[struct{}](i.ID, objc.Sel("setTag:"), value)
}



// The display priority associated with the toolbar item.
//
// See: https://developer.apple.com/documentation/appkit/nstoolbaritem/visibilitypriority-swift.property
func (i NSItemBadge) VisibilityPriority() NSToolbarItemVisibilityPriority {
	rv := objc.Send[NSToolbarItemVisibilityPriority](i.ID, objc.Sel("visibilityPriority"))
	return NSToolbarItemVisibilityPriority(rv)
}
func (i NSItemBadge) SetVisibilityPriority(value NSToolbarItemVisibilityPriority) {
	objc.Send[struct{}](i.ID, objc.Sel("setVisibilityPriority:"), value)
}


















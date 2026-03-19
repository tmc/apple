// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSMenuToolbarItem] class.
var (
	_NSMenuToolbarItemClass     NSMenuToolbarItemClass
	_NSMenuToolbarItemClassOnce sync.Once
)

func getNSMenuToolbarItemClass() NSMenuToolbarItemClass {
	_NSMenuToolbarItemClassOnce.Do(func() {
		_NSMenuToolbarItemClass = NSMenuToolbarItemClass{class: objc.GetClass("NSMenuToolbarItem")}
	})
	return _NSMenuToolbarItemClass
}

// GetNSMenuToolbarItemClass returns the class object for NSMenuToolbarItem.
func GetNSMenuToolbarItemClass() NSMenuToolbarItemClass {
	return getNSMenuToolbarItemClass()
}

type NSMenuToolbarItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMenuToolbarItemClass) Alloc() NSMenuToolbarItem {
	rv := objc.Send[NSMenuToolbarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A control that presents a menu in a window’s toolbar.
//
// # Overview
// 
// If you set an action on an [NSMenuToolbarItem] control item, the user
// invokes the action when clicking on the item through pressing and holding
// to display the menu. If you set an action on the item and [NSMenuToolbarItem.ShowsIndicator]
// to [true], the system displays the indicator as a separate segment so the
// user can invoke the menu with a click on that segment.
// 
// If you don’t set an action on the [NSMenuToolbarItem], a simple click
// invokes the menu, and the indicator is purely decorative.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Configuring a menu toolbar item
//
//   - [NSMenuToolbarItem.ShowsIndicator]: A Boolean value that determines whether the toolbar item displays an indicator of additional functionality.
//   - [NSMenuToolbarItem.SetShowsIndicator]
//   - [NSMenuToolbarItem.Menu]: The menu presented from the toolbar item.
//   - [NSMenuToolbarItem.SetMenu]
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem
type NSMenuToolbarItem struct {
	NSToolbarItem
}

// NSMenuToolbarItemFromID constructs a [NSMenuToolbarItem] from an objc.ID.
//
// A control that presents a menu in a window’s toolbar.
func NSMenuToolbarItemFromID(id objc.ID) NSMenuToolbarItem {
	return NSMenuToolbarItem{NSToolbarItem: NSToolbarItemFromID(id)}
}
// NOTE: NSMenuToolbarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMenuToolbarItem] class.
//
// # Configuring a menu toolbar item
//
//   - [INSMenuToolbarItem.ShowsIndicator]: A Boolean value that determines whether the toolbar item displays an indicator of additional functionality.
//   - [INSMenuToolbarItem.SetShowsIndicator]
//   - [INSMenuToolbarItem.Menu]: The menu presented from the toolbar item.
//   - [INSMenuToolbarItem.SetMenu]
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem
type INSMenuToolbarItem interface {
	INSToolbarItem

	// Topic: Configuring a menu toolbar item

	// A Boolean value that determines whether the toolbar item displays an indicator of additional functionality.
	ShowsIndicator() bool
	SetShowsIndicator(value bool)
	// The menu presented from the toolbar item.
	Menu() INSMenu
	SetMenu(value INSMenu)
}

// Init initializes the instance.
func (m NSMenuToolbarItem) Init() NSMenuToolbarItem {
	rv := objc.Send[NSMenuToolbarItem](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMenuToolbarItem) Autorelease() NSMenuToolbarItem {
	rv := objc.Send[NSMenuToolbarItem](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMenuToolbarItem creates a new NSMenuToolbarItem instance.
func NewNSMenuToolbarItem() NSMenuToolbarItem {
	class := getNSMenuToolbarItemClass()
	rv := objc.Send[NSMenuToolbarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a toolbar item with the specified identifier.
//
// itemIdentifier: The identifier for the toolbar item. You use this value to identify the
// item within your app, so you don’t need to localize it. For example, your
// toolbar delegate uses this value to identify the specific toolbar item.
//
// # Return Value
// 
// A new toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:)
func NewMenuToolbarItemWithItemIdentifier(itemIdentifier NSToolbarItemIdentifier) NSMenuToolbarItem {
	instance := getNSMenuToolbarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItemIdentifier:"), objc.String(string(itemIdentifier)))
	return NSMenuToolbarItemFromID(rv)
}

// A Boolean value that determines whether the toolbar item displays an
// indicator of additional functionality.
//
// # Discussion
// 
// If [true], the menu toolbar item shows an indicator that additional
// functionality is available.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem/showsIndicator
func (m NSMenuToolbarItem) ShowsIndicator() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("showsIndicator"))
	return rv
}
func (m NSMenuToolbarItem) SetShowsIndicator(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setShowsIndicator:"), value)
}
// The menu presented from the toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem/menu
func (m NSMenuToolbarItem) Menu() INSMenu {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("menu"))
	return NSMenuFromID(objc.ID(rv))
}
func (m NSMenuToolbarItem) SetMenu(value INSMenu) {
	objc.Send[struct{}](m.ID, objc.Sel("setMenu:"), value)
}


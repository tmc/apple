// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSearchToolbarItem] class.
var (
	_NSSearchToolbarItemClass     NSSearchToolbarItemClass
	_NSSearchToolbarItemClassOnce sync.Once
)

func getNSSearchToolbarItemClass() NSSearchToolbarItemClass {
	_NSSearchToolbarItemClassOnce.Do(func() {
		_NSSearchToolbarItemClass = NSSearchToolbarItemClass{class: objc.GetClass("NSSearchToolbarItem")}
	})
	return _NSSearchToolbarItemClass
}

// GetNSSearchToolbarItemClass returns the class object for NSSearchToolbarItem.
func GetNSSearchToolbarItemClass() NSSearchToolbarItemClass {
	return getNSSearchToolbarItemClass()
}

type NSSearchToolbarItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSearchToolbarItemClass) Alloc() NSSearchToolbarItem {
	rv := objc.Send[NSSearchToolbarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A toolbar item that contains a search field optimized for performing
// text-based searches.
//
// # Overview
// 
// [NSSearchToolbarItem] automatically resizes to accommodate typing when the
// focus switches to the toolbar item. When the toolbar is low on space, the
// system may collapse the search item into a button representation, which
// then expands to a full search field when the user clicks on it.
//
// # Configuring a search item
//
//   - [NSSearchToolbarItem.PreferredWidthForSearchField]: The preferred width for the toolbar item when it has keyboard focus.
//   - [NSSearchToolbarItem.SetPreferredWidthForSearchField]
//   - [NSSearchToolbarItem.ResignsFirstResponderWithCancel]: A Boolean value that enables the cancel button in the search field to resign the first responder in addition to clearing the contents.
//   - [NSSearchToolbarItem.SetResignsFirstResponderWithCancel]
//   - [NSSearchToolbarItem.SearchField]: The search field inside the toolbar item.
//   - [NSSearchToolbarItem.SetSearchField]
//
// # Controlling search interactions
//
//   - [NSSearchToolbarItem.BeginSearchInteraction]: Starts a search interaction and moves the keyboard focus to the search field.
//   - [NSSearchToolbarItem.EndSearchInteraction]: Ends a search interaction by giving up the first responder and adjusting the size of the search field to the available width for the toolbar item if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem
type NSSearchToolbarItem struct {
	NSToolbarItem
}

// NSSearchToolbarItemFromID constructs a [NSSearchToolbarItem] from an objc.ID.
//
// A toolbar item that contains a search field optimized for performing
// text-based searches.
func NSSearchToolbarItemFromID(id objc.ID) NSSearchToolbarItem {
	return NSSearchToolbarItem{NSToolbarItem: NSToolbarItemFromID(id)}
}
// NOTE: NSSearchToolbarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSearchToolbarItem] class.
//
// # Configuring a search item
//
//   - [INSSearchToolbarItem.PreferredWidthForSearchField]: The preferred width for the toolbar item when it has keyboard focus.
//   - [INSSearchToolbarItem.SetPreferredWidthForSearchField]
//   - [INSSearchToolbarItem.ResignsFirstResponderWithCancel]: A Boolean value that enables the cancel button in the search field to resign the first responder in addition to clearing the contents.
//   - [INSSearchToolbarItem.SetResignsFirstResponderWithCancel]
//   - [INSSearchToolbarItem.SearchField]: The search field inside the toolbar item.
//   - [INSSearchToolbarItem.SetSearchField]
//
// # Controlling search interactions
//
//   - [INSSearchToolbarItem.BeginSearchInteraction]: Starts a search interaction and moves the keyboard focus to the search field.
//   - [INSSearchToolbarItem.EndSearchInteraction]: Ends a search interaction by giving up the first responder and adjusting the size of the search field to the available width for the toolbar item if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem
type INSSearchToolbarItem interface {
	INSToolbarItem

	// Topic: Configuring a search item

	// The preferred width for the toolbar item when it has keyboard focus.
	PreferredWidthForSearchField() float64
	SetPreferredWidthForSearchField(value float64)
	// A Boolean value that enables the cancel button in the search field to resign the first responder in addition to clearing the contents.
	ResignsFirstResponderWithCancel() bool
	SetResignsFirstResponderWithCancel(value bool)
	// The search field inside the toolbar item.
	SearchField() INSSearchField
	SetSearchField(value INSSearchField)

	// Topic: Controlling search interactions

	// Starts a search interaction and moves the keyboard focus to the search field.
	BeginSearchInteraction()
	// Ends a search interaction by giving up the first responder and adjusting the size of the search field to the available width for the toolbar item if necessary.
	EndSearchInteraction()
}

// Init initializes the instance.
func (s NSSearchToolbarItem) Init() NSSearchToolbarItem {
	rv := objc.Send[NSSearchToolbarItem](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSearchToolbarItem) Autorelease() NSSearchToolbarItem {
	rv := objc.Send[NSSearchToolbarItem](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSearchToolbarItem creates a new NSSearchToolbarItem instance.
func NewNSSearchToolbarItem() NSSearchToolbarItem {
	class := getNSSearchToolbarItemClass()
	rv := objc.Send[NSSearchToolbarItem](objc.ID(class.class), objc.Sel("new"))
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
func NewSearchToolbarItemWithItemIdentifier(itemIdentifier NSToolbarItemIdentifier) NSSearchToolbarItem {
	instance := getNSSearchToolbarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItemIdentifier:"), objc.String(string(itemIdentifier)))
	return NSSearchToolbarItemFromID(rv)
}

// Starts a search interaction and moves the keyboard focus to the search
// field.
//
// # Discussion
// 
// If the system displays a compressed search field, starting the search
// interaction expands the field to the width stored in the
// [PreferredWidthForSearchField] property and moves the keyboard focus into
// the search field. Use [BeginSearchInteraction] and [EndSearchInteraction]
// to programmatically control a search.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/beginSearchInteraction()
func (s NSSearchToolbarItem) BeginSearchInteraction() {
	objc.Send[objc.ID](s.ID, objc.Sel("beginSearchInteraction"))
}
// Ends a search interaction by giving up the first responder and adjusting
// the size of the search field to the available width for the toolbar item if
// necessary.
//
// # Discussion
// 
// Use [BeginSearchInteraction] and [EndSearchInteraction] to programmatically
// control a search.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/endSearchInteraction()
func (s NSSearchToolbarItem) EndSearchInteraction() {
	objc.Send[objc.ID](s.ID, objc.Sel("endSearchInteraction"))
}

// The preferred width for the toolbar item when it has keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/preferredWidthForSearchField
func (s NSSearchToolbarItem) PreferredWidthForSearchField() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("preferredWidthForSearchField"))
	return rv
}
func (s NSSearchToolbarItem) SetPreferredWidthForSearchField(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreferredWidthForSearchField:"), value)
}
// A Boolean value that enables the cancel button in the search field to
// resign the first responder in addition to clearing the contents.
//
// # Discussion
// 
// The default value is `true`. If set to `false`, the cancel button only
// clears the contents of the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/resignsFirstResponderWithCancel
func (s NSSearchToolbarItem) ResignsFirstResponderWithCancel() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("resignsFirstResponderWithCancel"))
	return rv
}
func (s NSSearchToolbarItem) SetResignsFirstResponderWithCancel(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setResignsFirstResponderWithCancel:"), value)
}
// The search field inside the toolbar item.
//
// # Discussion
// 
// When you set `searchField` to `nil`, it uses the default configuration for
// the toolbar item, and inherits the item’s properties and layout
// constraints. However, if you want to customize the search field, you’ll
// need to add those settings before assigning it to the toolbar item. For
// more information about customizing a search field, see [NSSearchField].
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchToolbarItem/searchField
func (s NSSearchToolbarItem) SearchField() INSSearchField {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("searchField"))
	return NSSearchFieldFromID(objc.ID(rv))
}
func (s NSSearchToolbarItem) SetSearchField(value INSSearchField) {
	objc.Send[struct{}](s.ID, objc.Sel("setSearchField:"), value)
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSharingServicePickerToolbarItem] class.
var (
	_NSSharingServicePickerToolbarItemClass     NSSharingServicePickerToolbarItemClass
	_NSSharingServicePickerToolbarItemClassOnce sync.Once
)

func getNSSharingServicePickerToolbarItemClass() NSSharingServicePickerToolbarItemClass {
	_NSSharingServicePickerToolbarItemClassOnce.Do(func() {
		_NSSharingServicePickerToolbarItemClass = NSSharingServicePickerToolbarItemClass{class: objc.GetClass("NSSharingServicePickerToolbarItem")}
	})
	return _NSSharingServicePickerToolbarItemClass
}

// GetNSSharingServicePickerToolbarItemClass returns the class object for NSSharingServicePickerToolbarItem.
func GetNSSharingServicePickerToolbarItemClass() NSSharingServicePickerToolbarItemClass {
	return getNSSharingServicePickerToolbarItemClass()
}

type NSSharingServicePickerToolbarItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSharingServicePickerToolbarItemClass) Alloc() NSSharingServicePickerToolbarItem {
	rv := objc.Send[NSSharingServicePickerToolbarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A toolbar item that displays the macOS share sheet.
//
// # Overview
// 
// An [NSSharingServicePickerToolbarItem] object is a standard item you add to
// your window’s toolbar. When someone clicks it, the item displays the
// macOS share sheet. Use this item to share the selected or focal content
// from the current window. For example, you might share the photo someone is
// viewing, the currently selected text, or the window’s associated
// document.
// 
// Provide the items to share using the associated [NSSharingServicePickerToolbarItem.Delegate] object. For an
// app built using Mac Catalyst, provide the items from the object in the
// [NSSharingServicePickerToolbarItem.ActivityItemsConfiguration] property.
//
// # Getting the Toolbar Items
//
//   - [NSSharingServicePickerToolbarItem.Delegate]: The custom object from your app that provides the items to share.
//   - [NSSharingServicePickerToolbarItem.SetDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem
type NSSharingServicePickerToolbarItem struct {
	NSToolbarItem
}

// NSSharingServicePickerToolbarItemFromID constructs a [NSSharingServicePickerToolbarItem] from an objc.ID.
//
// A toolbar item that displays the macOS share sheet.
func NSSharingServicePickerToolbarItemFromID(id objc.ID) NSSharingServicePickerToolbarItem {
	return NSSharingServicePickerToolbarItem{NSToolbarItem: NSToolbarItemFromID(id)}
}
// NOTE: NSSharingServicePickerToolbarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSharingServicePickerToolbarItem] class.
//
// # Getting the Toolbar Items
//
//   - [INSSharingServicePickerToolbarItem.Delegate]: The custom object from your app that provides the items to share.
//   - [INSSharingServicePickerToolbarItem.SetDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem
type INSSharingServicePickerToolbarItem interface {
	INSToolbarItem

	// Topic: Getting the Toolbar Items

	// The custom object from your app that provides the items to share.
	Delegate() NSSharingServicePickerToolbarItemDelegate
	SetDelegate(value NSSharingServicePickerToolbarItemDelegate)
}

// Init initializes the instance.
func (s NSSharingServicePickerToolbarItem) Init() NSSharingServicePickerToolbarItem {
	rv := objc.Send[NSSharingServicePickerToolbarItem](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSharingServicePickerToolbarItem) Autorelease() NSSharingServicePickerToolbarItem {
	rv := objc.Send[NSSharingServicePickerToolbarItem](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSharingServicePickerToolbarItem creates a new NSSharingServicePickerToolbarItem instance.
func NewNSSharingServicePickerToolbarItem() NSSharingServicePickerToolbarItem {
	class := getNSSharingServicePickerToolbarItemClass()
	rv := objc.Send[NSSharingServicePickerToolbarItem](objc.ID(class.class), objc.Sel("new"))
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
func NewSharingServicePickerToolbarItemWithItemIdentifier(itemIdentifier NSToolbarItemIdentifier) NSSharingServicePickerToolbarItem {
	instance := getNSSharingServicePickerToolbarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItemIdentifier:"), objc.String(string(itemIdentifier)))
	return NSSharingServicePickerToolbarItemFromID(rv)
}

// The custom object from your app that provides the items to share.
//
// # Discussion
// 
// Use your delegate object to provide the set of items you want to share from
// your window. If this property is `nil`, AppKit disables the toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem/delegate
func (s NSSharingServicePickerToolbarItem) Delegate() NSSharingServicePickerToolbarItemDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSSharingServicePickerToolbarItemDelegateObjectFromID(rv)
}
func (s NSSharingServicePickerToolbarItem) SetDelegate(value NSSharingServicePickerToolbarItemDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}


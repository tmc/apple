// Code generated from Apple documentation for AppKit. DO NOT EDIT.
// +build ios

package appkit

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/objectivec"
)

// Creates a toolbar item with property values from the specified bar button
// item.
//
// itemIdentifier: The identifier for the toolbar item. You use this value to identify the
// item within your app, so you don’t need to localize it. For example, your
// toolbar delegate uses this value to identify the specific toolbar item.
//
// barButtonItem: The bar button item to use to create the toolbar item.
//
// barButtonItem is a [uikit.UIBarButtonItem].
//
// # Return Value
// 
// A new toolbar item.
//
// # Discussion
// 
// Use this method to create and initialize a toolbar item with property
// values from a [UIBarButtonItem], such as [Title], [Image], [Action], and
// [Target].
//
// [UIBarButtonItem]: https://developer.apple.com/documentation/UIKit/UIBarButtonItem
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:barButtonItem:)
// barButtonItem is a [uikit.UIBarButtonItem].
func (_NSToolbarItemClass NSToolbarItemClass) ItemWithItemIdentifierBarButtonItem(itemIdentifier NSToolbarItemIdentifier, barButtonItem objectivec.IObject) NSToolbarItem {
rv := objc.Send[objc.ID](objc.ID(_NSToolbarItemClass.class), objc.Sel("itemWithItemIdentifier:barButtonItem:"), objc.String(string(itemIdentifier)), barButtonItem)
return NSToolbarItemFromID(rv)
}

// The menu item to use for the toolbar item is in the overflow menu in a Mac
// app built with Mac Catalyst.
//
// # Discussion
// 
// The toolbar provides an initial default menu form representation that uses
// the toolbar item’s label as the menu item’s title. You can customize
// this menu item by changing the title or adding a submenu. When the toolbar
// is in text only mode, this menu item provides the text for the toolbar
// item. If the menu item in this property has a submenu and is visbile,
// clicking the toolbar item displays that submenu. If the toolbar item
// isn’t visible because it’s in the overflow menu, the menu item and
// submenu appear there.
// 
// For a discussion on menu forms, see [Setting a Toolbar Item’s
// Representation].
//
// [Setting a Toolbar Item’s Representation]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Toolbars/Tasks/SettingTBItemRep.html#//apple_ref/doc/uid/20000722
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/itemMenuFormRepresentation
func (t NSToolbarItem) ItemMenuFormRepresentation() objectivec.IObject {
rv := objc.Send[objc.ID](t.ID, objc.Sel("itemMenuFormRepresentation"))
return objectivec.Object{ID: rv}
}
func (t NSToolbarItem) SetItemMenuFormRepresentation(value objectivec.IObject) {
objc.Send[struct{}](t.ID, objc.Sel("setItemMenuFormRepresentation:"), value)
}


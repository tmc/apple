// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSTrackingSeparatorToolbarItem] class.
var (
	_NSTrackingSeparatorToolbarItemClass     NSTrackingSeparatorToolbarItemClass
	_NSTrackingSeparatorToolbarItemClassOnce sync.Once
)

func getNSTrackingSeparatorToolbarItemClass() NSTrackingSeparatorToolbarItemClass {
	_NSTrackingSeparatorToolbarItemClassOnce.Do(func() {
		_NSTrackingSeparatorToolbarItemClass = NSTrackingSeparatorToolbarItemClass{class: objc.GetClass("NSTrackingSeparatorToolbarItem")}
	})
	return _NSTrackingSeparatorToolbarItemClass
}

// GetNSTrackingSeparatorToolbarItemClass returns the class object for NSTrackingSeparatorToolbarItem.
func GetNSTrackingSeparatorToolbarItemClass() NSTrackingSeparatorToolbarItemClass {
	return getNSTrackingSeparatorToolbarItemClass()
}

type NSTrackingSeparatorToolbarItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTrackingSeparatorToolbarItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTrackingSeparatorToolbarItemClass) Alloc() NSTrackingSeparatorToolbarItem {
	rv := objc.Send[NSTrackingSeparatorToolbarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A toolbar separator that aligns with the vertical split view in the same
// window.
//
// # Overview
// 
// Use a [NSTrackingSeparatorToolbarItem] to divide an [NSToolbar] into
// sections that visually align with the views on either side of the divider
// of the [SplitView]. This keeps [NSToolbarItem]s above the content that’s
// the [NSTrackingSeparatorToolbarItem.Target] for the item’s [NSTrackingSeparatorToolbarItem.Target].
// 
// The `splitView` must be in the same window as the toolbar containing this
// item before showing the toolbar.
//
// # configuring a tracking separator
//
//   - [NSTrackingSeparatorToolbarItem.DividerIndex]: The index of the split view divider to align with the tracking separator.
//   - [NSTrackingSeparatorToolbarItem.SetDividerIndex]
//   - [NSTrackingSeparatorToolbarItem.SplitView]: The vertical split view to align with the toolbar separator.
//   - [NSTrackingSeparatorToolbarItem.SetSplitView]
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingSeparatorToolbarItem
type NSTrackingSeparatorToolbarItem struct {
	NSToolbarItem
}

// NSTrackingSeparatorToolbarItemFromID constructs a [NSTrackingSeparatorToolbarItem] from an objc.ID.
//
// A toolbar separator that aligns with the vertical split view in the same
// window.
func NSTrackingSeparatorToolbarItemFromID(id objc.ID) NSTrackingSeparatorToolbarItem {
	return NSTrackingSeparatorToolbarItem{NSToolbarItem: NSToolbarItemFromID(id)}
}
// NOTE: NSTrackingSeparatorToolbarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTrackingSeparatorToolbarItem] class.
//
// # configuring a tracking separator
//
//   - [INSTrackingSeparatorToolbarItem.DividerIndex]: The index of the split view divider to align with the tracking separator.
//   - [INSTrackingSeparatorToolbarItem.SetDividerIndex]
//   - [INSTrackingSeparatorToolbarItem.SplitView]: The vertical split view to align with the toolbar separator.
//   - [INSTrackingSeparatorToolbarItem.SetSplitView]
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingSeparatorToolbarItem
type INSTrackingSeparatorToolbarItem interface {
	INSToolbarItem

	// Topic: configuring a tracking separator

	// The index of the split view divider to align with the tracking separator.
	DividerIndex() int
	SetDividerIndex(value int)
	// The vertical split view to align with the toolbar separator.
	SplitView() INSSplitView
	SetSplitView(value INSSplitView)
}

// Init initializes the instance.
func (t NSTrackingSeparatorToolbarItem) Init() NSTrackingSeparatorToolbarItem {
	rv := objc.Send[NSTrackingSeparatorToolbarItem](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTrackingSeparatorToolbarItem) Autorelease() NSTrackingSeparatorToolbarItem {
	rv := objc.Send[NSTrackingSeparatorToolbarItem](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTrackingSeparatorToolbarItem creates a new NSTrackingSeparatorToolbarItem instance.
func NewNSTrackingSeparatorToolbarItem() NSTrackingSeparatorToolbarItem {
	class := getNSTrackingSeparatorToolbarItemClass()
	rv := objc.Send[NSTrackingSeparatorToolbarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new tracking separator toolbar item and configures it to align
// with the divider of the split view.
//
// identifier: The identifier of the toolbar item.
//
// splitView: The split view to align with the tracking separator.
//
// dividerIndex: The index of the divider in the split view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingSeparatorToolbarItem/init(identifier:splitView:dividerIndex:)
func NewTrackingSeparatorToolbarItemWithIdentifierSplitViewDividerIndex(identifier NSToolbarItemIdentifier, splitView INSSplitView, dividerIndex int) NSTrackingSeparatorToolbarItem {
	rv := objc.Send[objc.ID](objc.ID(getNSTrackingSeparatorToolbarItemClass().class), objc.Sel("trackingSeparatorToolbarItemWithIdentifier:splitView:dividerIndex:"), objc.String(string(identifier)), splitView, dividerIndex)
	return NSTrackingSeparatorToolbarItemFromID(rv)
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
func NewTrackingSeparatorToolbarItemWithItemIdentifier(itemIdentifier NSToolbarItemIdentifier) NSTrackingSeparatorToolbarItem {
	instance := getNSTrackingSeparatorToolbarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItemIdentifier:"), objc.String(string(itemIdentifier)))
	return NSTrackingSeparatorToolbarItemFromID(rv)
}

// The index of the split view divider to align with the tracking separator.
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingSeparatorToolbarItem/dividerIndex
func (t NSTrackingSeparatorToolbarItem) DividerIndex() int {
	rv := objc.Send[int](t.ID, objc.Sel("dividerIndex"))
	return rv
}
func (t NSTrackingSeparatorToolbarItem) SetDividerIndex(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setDividerIndex:"), value)
}
// The vertical split view to align with the toolbar separator.
//
// # Discussion
// 
// The `splitView` must be in the same window as the toolbar containing the
// [NSTrackingSeparatorToolbarItem] before showing the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingSeparatorToolbarItem/splitView
func (t NSTrackingSeparatorToolbarItem) SplitView() INSSplitView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("splitView"))
	return NSSplitViewFromID(objc.ID(rv))
}
func (t NSTrackingSeparatorToolbarItem) SetSplitView(value INSSplitView) {
	objc.Send[struct{}](t.ID, objc.Sel("setSplitView:"), value)
}


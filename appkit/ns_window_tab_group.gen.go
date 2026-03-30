// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSWindowTabGroup] class.
var (
	_NSWindowTabGroupClass     NSWindowTabGroupClass
	_NSWindowTabGroupClassOnce sync.Once
)

func getNSWindowTabGroupClass() NSWindowTabGroupClass {
	_NSWindowTabGroupClassOnce.Do(func() {
		_NSWindowTabGroupClass = NSWindowTabGroupClass{class: objc.GetClass("NSWindowTabGroup")}
	})
	return _NSWindowTabGroupClass
}

// GetNSWindowTabGroupClass returns the class object for NSWindowTabGroup.
func GetNSWindowTabGroupClass() NSWindowTabGroupClass {
	return getNSWindowTabGroupClass()
}

type NSWindowTabGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSWindowTabGroupClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWindowTabGroupClass) Alloc() NSWindowTabGroup {
	rv := objc.Send[NSWindowTabGroup](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A group of windows that display together as a single tabbed window.
//
// # Overview
//
// AppKit automatically creates instances of [NSWindowTabGroup] to reflect the
// tabbing state of your windows. You can access a window’s current tab
// group using the [NSWindowTabGroup.TabGroup] property.
//
// # Checking the Group Identifier
//
//   - [NSWindowTabGroup.Identifier]: The unique identifier for a tabbed window group.
//
// # Configuring the Tab User Interface
//
//   - [NSWindowTabGroup.OverviewVisible]: A Boolean value indicating if the tab overview is currently displayed.
//   - [NSWindowTabGroup.SetOverviewVisible]
//   - [NSWindowTabGroup.TabBarVisible]: A Boolean value indicating whether the tabbed window group currently displays a tab bar.
//
// # Managing Tabbed Windows
//
//   - [NSWindowTabGroup.Windows]: A collection of the windows that are currently grouped together by this window tab group.
//   - [NSWindowTabGroup.SelectedWindow]: The selected, or frontmost, window in the tab group.
//   - [NSWindowTabGroup.SetSelectedWindow]
//   - [NSWindowTabGroup.AddWindow]: Adds a window to the tab group.
//   - [NSWindowTabGroup.InsertWindowAtIndex]: Inserts a window at a specific location within the tab group.
//   - [NSWindowTabGroup.RemoveWindow]: Removes a window from the tab group.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup
type NSWindowTabGroup struct {
	objectivec.Object
}

// NSWindowTabGroupFromID constructs a [NSWindowTabGroup] from an objc.ID.
//
// A group of windows that display together as a single tabbed window.
func NSWindowTabGroupFromID(id objc.ID) NSWindowTabGroup {
	return NSWindowTabGroup{objectivec.Object{ID: id}}
}

// NOTE: NSWindowTabGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSWindowTabGroup] class.
//
// # Checking the Group Identifier
//
//   - [INSWindowTabGroup.Identifier]: The unique identifier for a tabbed window group.
//
// # Configuring the Tab User Interface
//
//   - [INSWindowTabGroup.OverviewVisible]: A Boolean value indicating if the tab overview is currently displayed.
//   - [INSWindowTabGroup.SetOverviewVisible]
//   - [INSWindowTabGroup.TabBarVisible]: A Boolean value indicating whether the tabbed window group currently displays a tab bar.
//
// # Managing Tabbed Windows
//
//   - [INSWindowTabGroup.Windows]: A collection of the windows that are currently grouped together by this window tab group.
//   - [INSWindowTabGroup.SelectedWindow]: The selected, or frontmost, window in the tab group.
//   - [INSWindowTabGroup.SetSelectedWindow]
//   - [INSWindowTabGroup.AddWindow]: Adds a window to the tab group.
//   - [INSWindowTabGroup.InsertWindowAtIndex]: Inserts a window at a specific location within the tab group.
//   - [INSWindowTabGroup.RemoveWindow]: Removes a window from the tab group.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup
type INSWindowTabGroup interface {
	objectivec.IObject

	// Topic: Checking the Group Identifier

	// The unique identifier for a tabbed window group.
	Identifier() NSWindowTabbingIdentifier

	// Topic: Configuring the Tab User Interface

	// A Boolean value indicating if the tab overview is currently displayed.
	OverviewVisible() bool
	SetOverviewVisible(value bool)
	// A Boolean value indicating whether the tabbed window group currently displays a tab bar.
	TabBarVisible() bool

	// Topic: Managing Tabbed Windows

	// A collection of the windows that are currently grouped together by this window tab group.
	Windows() []NSWindow
	// The selected, or frontmost, window in the tab group.
	SelectedWindow() INSWindow
	SetSelectedWindow(value INSWindow)
	// Adds a window to the tab group.
	AddWindow(window INSWindow)
	// Inserts a window at a specific location within the tab group.
	InsertWindowAtIndex(window INSWindow, index int)
	// Removes a window from the tab group.
	RemoveWindow(window INSWindow)

	// A group of windows that display together as a tab group.
	TabGroup() INSWindowTabGroup
	SetTabGroup(value INSWindowTabGroup)
}

// Init initializes the instance.
func (w NSWindowTabGroup) Init() NSWindowTabGroup {
	rv := objc.Send[NSWindowTabGroup](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWindowTabGroup) Autorelease() NSWindowTabGroup {
	rv := objc.Send[NSWindowTabGroup](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWindowTabGroup creates a new NSWindowTabGroup instance.
func NewNSWindowTabGroup() NSWindowTabGroup {
	class := getNSWindowTabGroupClass()
	rv := objc.Send[NSWindowTabGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds a window to the tab group.
//
// window: The window to append to the tab group.
//
// # Discussion
//
// This method appends the window to the end of the tab group. If the window
// is already a member of another tab group, it is first removed from that
// group.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/addWindow(_:)
func (w NSWindowTabGroup) AddWindow(window INSWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("addWindow:"), window)
}

// Inserts a window at a specific location within the tab group.
//
// window: The window to insert into the tab group.
//
// index: The location in the tab group at which to insert window. This value must
// not be negative, and must not be greater than the number of windows in the
// tab group.
//
// Raises an [NSInternalInconsistencyException] if the index is negative or
// larger than the number of tabs in the group.
//
// # Discussion
//
// Inserts the window at the specified location within the tab group. If the
// window is already a member of another tab group, it is first removed from
// that group.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/insertWindow(_:at:)
func (w NSWindowTabGroup) InsertWindowAtIndex(window INSWindow, index int) {
	objc.Send[objc.ID](w.ID, objc.Sel("insertWindow:atIndex:"), window, index)
}

// Removes a window from the tab group.
//
// window: The window to remove from the tab group. This window must already be a
// member of the tab group.
//
// Raises an [NSInternalInconsistencyException] if the window is not a member
// of the tab group.
//
// # Discussion
//
// You can use [RemoveWindow] to explicitly remove a window from the tab
// group. Windows are implicity removed from their associated tab groups when
// they order out.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/removeWindow(_:)
func (w NSWindowTabGroup) RemoveWindow(window INSWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("removeWindow:"), window)
}

// The unique identifier for a tabbed window group.
//
// # Discussion
//
// This identifier is shared by all windows that participate in this tab
// group.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/identifier
func (w NSWindowTabGroup) Identifier() NSWindowTabbingIdentifier {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("identifier"))
	return NSWindowTabbingIdentifier(foundation.NSStringFromID(rv).String())
}

// A Boolean value indicating if the tab overview is currently displayed.
//
// # Discussion
//
// The tab overview provides a visual overview of the windows that make up a
// tabbing group. This property indicates whether the tab overview is
// currently displayed. Setting the property either shows or hides the
// overview.
//
// You can monitor this property for changes using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/isOverviewVisible
func (w NSWindowTabGroup) OverviewVisible() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isOverviewVisible"))
	return rv
}
func (w NSWindowTabGroup) SetOverviewVisible(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setOverviewVisible:"), value)
}

// A Boolean value indicating whether the tabbed window group currently
// displays a tab bar.
//
// # Discussion
//
// Typically, a tabbed window displays a tab bar if there is more than one
// window in the tabbing group. The tab bar can also be manually toggled using
// the [ToggleTabBar] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/isTabBarVisible
func (w NSWindowTabGroup) TabBarVisible() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isTabBarVisible"))
	return rv
}

// A collection of the windows that are currently grouped together by this
// window tab group.
//
// # Discussion
//
// The order of this array corresponds to the visual order of the tabs in this
// tab group, arranged from the leading edge of the tab bar to the trailing
// edge. You can monitor this property for changes using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/windows
func (w NSWindowTabGroup) Windows() []NSWindow {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("windows"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSWindow {
		return NSWindowFromID(id)
	})
}

// The selected, or frontmost, window in the tab group.
//
// # Discussion
//
// This property returns the currently selected window within the tabbed
// window group. Setting this property visually changes the selected tab.
//
// You can monitor this property for changes using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/selectedWindow
func (w NSWindowTabGroup) SelectedWindow() INSWindow {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("selectedWindow"))
	return NSWindowFromID(objc.ID(rv))
}
func (w NSWindowTabGroup) SetSelectedWindow(value INSWindow) {
	objc.Send[struct{}](w.ID, objc.Sel("setSelectedWindow:"), value)
}

// A group of windows that display together as a tab group.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/tabgroup
func (w NSWindowTabGroup) TabGroup() INSWindowTabGroup {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("tabGroup"))
	return NSWindowTabGroupFromID(objc.ID(rv))
}
func (w NSWindowTabGroup) SetTabGroup(value INSWindowTabGroup) {
	objc.Send[struct{}](w.ID, objc.Sel("setTabGroup:"), value)
}

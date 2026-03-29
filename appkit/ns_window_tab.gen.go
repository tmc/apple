// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSWindowTab] class.
var (
	_NSWindowTabClass     NSWindowTabClass
	_NSWindowTabClassOnce sync.Once
)

func getNSWindowTabClass() NSWindowTabClass {
	_NSWindowTabClassOnce.Do(func() {
		_NSWindowTabClass = NSWindowTabClass{class: objc.GetClass("NSWindowTab")}
	})
	return _NSWindowTabClass
}

// GetNSWindowTabClass returns the class object for NSWindowTab.
func GetNSWindowTabClass() NSWindowTabClass {
	return getNSWindowTabClass()
}

type NSWindowTabClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSWindowTabClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWindowTabClass) Alloc() NSWindowTab {
	rv := objc.Send[NSWindowTab](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A tab associated with a window that is part of a tabbing group.
//
// # Overview
// 
// [NSWindowTab] describes the way a window displays as part of a tabbed
// window group. The properties of [NSWindowTab] are configurable at any time,
// but only take effect when the associated [NSWindow] displays in a tab.
// 
// AppKit automatically creates an instance of [NSWindowTab] for each
// [NSWindow]. You can access a window’s tab object using the [NSWindowTab.Tab]
// property.
//
// # Customizing the Title
//
//   - [NSWindowTab.Title]: The title for the window tab.
//   - [NSWindowTab.SetTitle]
//   - [NSWindowTab.AttributedTitle]: The title for the window tab, specified as an attributed string.
//   - [NSWindowTab.SetAttributedTitle]
//
// # Customizing the Tooltip
//
//   - [NSWindowTab.ToolTip]: The tooltip for this window tab.
//   - [NSWindowTab.SetToolTip]
//
// # Adding an Accessory View
//
//   - [NSWindowTab.AccessoryView]: An optional accessory view for the tab.
//   - [NSWindowTab.SetAccessoryView]
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTab
type NSWindowTab struct {
	objectivec.Object
}

// NSWindowTabFromID constructs a [NSWindowTab] from an objc.ID.
//
// A tab associated with a window that is part of a tabbing group.
func NSWindowTabFromID(id objc.ID) NSWindowTab {
	return NSWindowTab{objectivec.Object{ID: id}}
}
// NOTE: NSWindowTab adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSWindowTab] class.
//
// # Customizing the Title
//
//   - [INSWindowTab.Title]: The title for the window tab.
//   - [INSWindowTab.SetTitle]
//   - [INSWindowTab.AttributedTitle]: The title for the window tab, specified as an attributed string.
//   - [INSWindowTab.SetAttributedTitle]
//
// # Customizing the Tooltip
//
//   - [INSWindowTab.ToolTip]: The tooltip for this window tab.
//   - [INSWindowTab.SetToolTip]
//
// # Adding an Accessory View
//
//   - [INSWindowTab.AccessoryView]: An optional accessory view for the tab.
//   - [INSWindowTab.SetAccessoryView]
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTab
type INSWindowTab interface {
	objectivec.IObject

	// Topic: Customizing the Title

	// The title for the window tab.
	Title() string
	SetTitle(value string)
	// The title for the window tab, specified as an attributed string.
	AttributedTitle() foundation.NSAttributedString
	SetAttributedTitle(value foundation.NSAttributedString)

	// Topic: Customizing the Tooltip

	// The tooltip for this window tab.
	ToolTip() string
	SetToolTip(value string)

	// Topic: Adding an Accessory View

	// An optional accessory view for the tab.
	AccessoryView() INSView
	SetAccessoryView(value INSView)

	// An object that represents information about a window when it displays as a tab.
	Tab() INSWindowTab
	SetTab(value INSWindowTab)
	// A value that allows a group of related windows.
	TabbingIdentifier() NSWindowTabbingIdentifier
	SetTabbingIdentifier(value NSWindowTabbingIdentifier)
}

// Init initializes the instance.
func (w NSWindowTab) Init() NSWindowTab {
	rv := objc.Send[NSWindowTab](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWindowTab) Autorelease() NSWindowTab {
	rv := objc.Send[NSWindowTab](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWindowTab creates a new NSWindowTab instance.
func NewNSWindowTab() NSWindowTab {
	class := getNSWindowTabClass()
	rv := objc.Send[NSWindowTab](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The title for the window tab.
//
// # Discussion
// 
// The title displays within the window tab when the associated window is part
// of a tabbing group.
// 
// By default, the title of the window tab follows the title of its associated
// window, but it may be customized using the [Title] property. If the title
// has been customized, setting the [Title] property to [nil] causes it to
// follow the window’s title again.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTab/title
func (w NSWindowTab) Title() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (w NSWindowTab) SetTitle(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setTitle:"), objc.String(value))
}
// The title for the window tab, specified as an attributed string.
//
// # Discussion
// 
// If you provide an attributed title, the window tab uses all of the
// attributes that are explicitly provided in the attributed string.
// Attributes that are left unspecified, including the font and foreground
// color, are automatically filled in using default values appropriate for the
// window tab.
// 
// If the [AttributedTitle] property is nil, the window tab uses the [Title]
// property instead. The default value is [nil].
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTab/attributedTitle
func (w NSWindowTab) AttributedTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("attributedTitle"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (w NSWindowTab) SetAttributedTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](w.ID, objc.Sel("setAttributedTitle:"), value)
}
// The tooltip for this window tab.
//
// # Discussion
// 
// By default, the window tab’s tooltip displays its [Title] string. Once
// customized, setting the [ToolTip] property to [nil] causes it to follow the
// [Title] property again.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTab/toolTip
func (w NSWindowTab) ToolTip() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("toolTip"))
	return foundation.NSStringFromID(rv).String()
}
func (w NSWindowTab) SetToolTip(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setToolTip:"), objc.String(value))
}
// An optional accessory view for the tab.
//
// # Discussion
// 
// You can customize the window tab by adding an accessory view that displays
// alongside the tab’s title.
// 
// The [translatesAutoresizingMaskIntoConstraints] property is automatically
// set to [false] on the view. Constraints can be created and activated to
// specify the view’s width and height values. A constraint is automatically
// added to vertically center the view, and to right align the view within the
// tab.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [translatesAutoresizingMaskIntoConstraints]: https://developer.apple.com/documentation/AppKit/NSView/translatesAutoresizingMaskIntoConstraints
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowTab/accessoryView
func (w NSWindowTab) AccessoryView() INSView {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("accessoryView"))
	return NSViewFromID(objc.ID(rv))
}
func (w NSWindowTab) SetAccessoryView(value INSView) {
	objc.Send[struct{}](w.ID, objc.Sel("setAccessoryView:"), value)
}
// An object that represents information about a window when it displays as a
// tab.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/tab
func (w NSWindowTab) Tab() INSWindowTab {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("tab"))
	return NSWindowTabFromID(objc.ID(rv))
}
func (w NSWindowTab) SetTab(value INSWindowTab) {
	objc.Send[struct{}](w.ID, objc.Sel("setTab:"), value)
}
// A value that allows a group of related windows.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/tabbingidentifier-swift.property
func (w NSWindowTab) TabbingIdentifier() NSWindowTabbingIdentifier {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("tabbingIdentifier"))
	return NSWindowTabbingIdentifier(foundation.NSStringFromID(rv).String())
}
func (w NSWindowTab) SetTabbingIdentifier(value NSWindowTabbingIdentifier) {
	objc.Send[struct{}](w.ID, objc.Sel("setTabbingIdentifier:"), objc.String(string(value)))
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSStatusItem] class.
var (
	_NSStatusItemClass     NSStatusItemClass
	_NSStatusItemClassOnce sync.Once
)

func getNSStatusItemClass() NSStatusItemClass {
	_NSStatusItemClassOnce.Do(func() {
		_NSStatusItemClass = NSStatusItemClass{class: objc.GetClass("NSStatusItem")}
	})
	return _NSStatusItemClass
}

// GetNSStatusItemClass returns the class object for NSStatusItem.
func GetNSStatusItemClass() NSStatusItemClass {
	return getNSStatusItemClass()
}

type NSStatusItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSStatusItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStatusItemClass) Alloc() NSStatusItem {
	rv := objc.Send[NSStatusItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An individual element displayed in the system menu bar.
//
// # Overview
//
// The [NSStatusBar] method [StatusItemWithLength] creates instances of this
// class and automatically adds them to the menu bar. Use the [Button]
// property to customize the appearance and behavior of the status item.
//
// # Getting the Item’s Status Bar
//
//   - [NSStatusItem.StatusBar]: The status bar that displays the status item.
//
// # Managing the Status Item’s Behavior
//
//   - [NSStatusItem.Behavior]: The set of allowed behaviors for the status item.
//   - [NSStatusItem.SetBehavior]
//   - [NSStatusItem.Button]: The button displayed in the status bar.
//   - [NSStatusItem.Menu]: The pull-down menu displayed when the user clicks the status item.
//   - [NSStatusItem.SetMenu]
//
// # Configuring the Status Item’s Appearance
//
//   - [NSStatusItem.Visible]: A Boolean value indicating if the menu bar currently displays the status item.
//   - [NSStatusItem.SetVisible]
//   - [NSStatusItem.Length]: The amount of space in the status bar that should be allocated to the status item.
//   - [NSStatusItem.SetLength]
//
// # Setting the Autosave Name
//
//   - [NSStatusItem.AutosaveName]: A unique name for saving and restoring information about a status item.
//   - [NSStatusItem.SetAutosaveName]
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusItem
type NSStatusItem struct {
	objectivec.Object
}

// NSStatusItemFromID constructs a [NSStatusItem] from an objc.ID.
//
// An individual element displayed in the system menu bar.
func NSStatusItemFromID(id objc.ID) NSStatusItem {
	return NSStatusItem{objectivec.Object{ID: id}}
}

// NOTE: NSStatusItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSStatusItem] class.
//
// # Getting the Item’s Status Bar
//
//   - [INSStatusItem.StatusBar]: The status bar that displays the status item.
//
// # Managing the Status Item’s Behavior
//
//   - [INSStatusItem.Behavior]: The set of allowed behaviors for the status item.
//   - [INSStatusItem.SetBehavior]
//   - [INSStatusItem.Button]: The button displayed in the status bar.
//   - [INSStatusItem.Menu]: The pull-down menu displayed when the user clicks the status item.
//   - [INSStatusItem.SetMenu]
//
// # Configuring the Status Item’s Appearance
//
//   - [INSStatusItem.Visible]: A Boolean value indicating if the menu bar currently displays the status item.
//   - [INSStatusItem.SetVisible]
//   - [INSStatusItem.Length]: The amount of space in the status bar that should be allocated to the status item.
//   - [INSStatusItem.SetLength]
//
// # Setting the Autosave Name
//
//   - [INSStatusItem.AutosaveName]: A unique name for saving and restoring information about a status item.
//   - [INSStatusItem.SetAutosaveName]
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusItem
type INSStatusItem interface {
	objectivec.IObject

	// Topic: Getting the Item’s Status Bar

	// The status bar that displays the status item.
	StatusBar() INSStatusBar

	// Topic: Managing the Status Item’s Behavior

	// The set of allowed behaviors for the status item.
	Behavior() NSStatusItemBehavior
	SetBehavior(value NSStatusItemBehavior)
	// The button displayed in the status bar.
	Button() INSStatusBarButton
	// The pull-down menu displayed when the user clicks the status item.
	Menu() INSMenu
	SetMenu(value INSMenu)

	// Topic: Configuring the Status Item’s Appearance

	// A Boolean value indicating if the menu bar currently displays the status item.
	Visible() bool
	SetVisible(value bool)
	// The amount of space in the status bar that should be allocated to the status item.
	Length() float64
	SetLength(value float64)

	// Topic: Setting the Autosave Name

	// A unique name for saving and restoring information about a status item.
	AutosaveName() NSStatusItemAutosaveName
	SetAutosaveName(value NSStatusItemAutosaveName)
}

// Init initializes the instance.
func (s NSStatusItem) Init() NSStatusItem {
	rv := objc.Send[NSStatusItem](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStatusItem) Autorelease() NSStatusItem {
	rv := objc.Send[NSStatusItem](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStatusItem creates a new NSStatusItem instance.
func NewNSStatusItem() NSStatusItem {
	class := getNSStatusItemClass()
	rv := objc.Send[NSStatusItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The status bar that displays the status item.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/statusBar
func (s NSStatusItem) StatusBar() INSStatusBar {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("statusBar"))
	return NSStatusBarFromID(objc.ID(rv))
}

// The set of allowed behaviors for the status item.
//
// # Discussion
//
// By default, this property includes no behavior options. See
// [NSStatusItem.Behavior] for a list of available behavior options and their
// effects.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/behavior-swift.property
//
// [NSStatusItem.Behavior]: https://developer.apple.com/documentation/AppKit/NSStatusItem/Behavior-swift.struct
func (s NSStatusItem) Behavior() NSStatusItemBehavior {
	rv := objc.Send[NSStatusItemBehavior](s.ID, objc.Sel("behavior"))
	return NSStatusItemBehavior(rv)
}
func (s NSStatusItem) SetBehavior(value NSStatusItemBehavior) {
	objc.Send[struct{}](s.ID, objc.Sel("setBehavior:"), value)
}

// The button displayed in the status bar.
//
// # Discussion
//
// The status item automatically creates this button by default. Use this
// property to customize the appearance and behavior of the button, such as
// its [Image], [Target], [Action], [ToolTip], and so on.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/button
func (s NSStatusItem) Button() INSStatusBarButton {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("button"))
	return NSStatusBarButtonFromID(objc.ID(rv))
}

// The pull-down menu displayed when the user clicks the status item.
//
// # Discussion
//
// When non-`nil`, the status item’s single click action behavior is not
// used. Setting the value of this property to `nil` removes the menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/menu
func (s NSStatusItem) Menu() INSMenu {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("menu"))
	return NSMenuFromID(objc.ID(rv))
}
func (s NSStatusItem) SetMenu(value INSMenu) {
	objc.Send[struct{}](s.ID, objc.Sel("setMenu:"), value)
}

// A Boolean value indicating if the menu bar currently displays the status
// item.
//
// # Discussion
//
// Setting this property either shows or hides the status item within the menu
// bar. The item’s visiblity may also change if the user removes the item
// manually, and you can watch for changes in visibility using key-value
// observation. The status item’s visiblity persists and restores
// automatically based on the value of [AutosaveName].
//
// This property returns true even if the status item is temporarily hidden
// due to insufficient space in the menu bar. The default value is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/isVisible
func (s NSStatusItem) Visible() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isVisible"))
	return rv
}
func (s NSStatusItem) SetVisible(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setVisible:"), value)
}

// The amount of space in the status bar that should be allocated to the
// status item.
//
// # Discussion
//
// If the status bar is horizontal, the value of this property is the width of
// the status item. In addition to a fixed length, this value can be
// NSSquareStatusItemLength or NSVariableStatusItemLength (see [NSStatusBar]
// Constants) to allow the status bar to allocate (and adjust) the space
// according to either the status bar’s thickness or the status item’s
// true size.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/length
func (s NSStatusItem) Length() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("length"))
	return rv
}
func (s NSStatusItem) SetLength(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setLength:"), value)
}

// A unique name for saving and restoring information about a status item.
//
// # Discussion
//
// If you do not provide an autosave name for a status item, the system
// automatically chooses a unique name. Setting this property to nil resets it
// to the automatically chosen name.
//
// Applications with multiple status items should set an autosave name after
// creating each item.
//
// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/autosaveName-swift.property
func (s NSStatusItem) AutosaveName() NSStatusItemAutosaveName {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("autosaveName"))
	return NSStatusItemAutosaveName(foundation.NSStringFromID(rv).String())
}
func (s NSStatusItem) SetAutosaveName(value NSStatusItemAutosaveName) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutosaveName:"), objc.String(string(value)))
}

// A status item length that is equal to the status bar’s thickness.
//
// See: https://developer.apple.com/documentation/appkit/nsstatusitem/squarelength
func (_NSStatusItemClass NSStatusItemClass) SquareLength() float64 {
	rv := objc.Send[float64](objc.ID(_NSStatusItemClass.class), objc.Sel("NSSquareStatusItemLength"))
	return rv
}

// A status item length that dynamically adjusts to the width of its contents.
//
// See: https://developer.apple.com/documentation/appkit/nsstatusitem/variablelength
func (_NSStatusItemClass NSStatusItemClass) VariableLength() float64 {
	rv := objc.Send[float64](objc.ID(_NSStatusItemClass.class), objc.Sel("NSVariableStatusItemLength"))
	return rv
}

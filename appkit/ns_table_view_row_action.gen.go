// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTableViewRowAction] class.
var (
	_NSTableViewRowActionClass     NSTableViewRowActionClass
	_NSTableViewRowActionClassOnce sync.Once
)

func getNSTableViewRowActionClass() NSTableViewRowActionClass {
	_NSTableViewRowActionClassOnce.Do(func() {
		_NSTableViewRowActionClass = NSTableViewRowActionClass{class: objc.GetClass("NSTableViewRowAction")}
	})
	return _NSTableViewRowActionClass
}

// GetNSTableViewRowActionClass returns the class object for NSTableViewRowAction.
func GetNSTableViewRowActionClass() NSTableViewRowActionClass {
	return getNSTableViewRowActionClass()
}

type NSTableViewRowActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTableViewRowActionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTableViewRowActionClass) Alloc() NSTableViewRowAction {
	rv := objc.Send[NSTableViewRowAction](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A single action to present when the user swipes horizontally on a table
// row.
//
// # Overview
// 
// In an editable table, performing a horizontal swipe on a row reveals a
// button to delete the row by default. This class lets you define one or more
// custom actions to display for a given row in your table. Each instance of
// this class represents a single action to perform and includes the text,
// formatting information, and behavior for the corresponding button.
// 
// To add custom actions to your table view’s rows, implement the
// [TableViewRowActionsForRowEdge] method in your table view’s delegate
// object. In that method, create and return an array of actions for the
// specified row. The table handles the remaining work of displaying the
// action buttons and executing the appropriate handler block when the user
// clicks the button.
//
// # Configuring the Action’s Appearance
//
//   - [NSTableViewRowAction.Style]: The style applied to the action button.
//   - [NSTableViewRowAction.Title]: The title of the action button.
//   - [NSTableViewRowAction.SetTitle]
//   - [NSTableViewRowAction.BackgroundColor]: The background color of the action button.
//   - [NSTableViewRowAction.SetBackgroundColor]
//
// # Instance Properties
//
//   - [NSTableViewRowAction.Image]
//   - [NSTableViewRowAction.SetImage]
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction
type NSTableViewRowAction struct {
	objectivec.Object
}

// NSTableViewRowActionFromID constructs a [NSTableViewRowAction] from an objc.ID.
//
// A single action to present when the user swipes horizontally on a table
// row.
func NSTableViewRowActionFromID(id objc.ID) NSTableViewRowAction {
	return NSTableViewRowAction{objectivec.Object{ID: id}}
}
// NOTE: NSTableViewRowAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTableViewRowAction] class.
//
// # Configuring the Action’s Appearance
//
//   - [INSTableViewRowAction.Style]: The style applied to the action button.
//   - [INSTableViewRowAction.Title]: The title of the action button.
//   - [INSTableViewRowAction.SetTitle]
//   - [INSTableViewRowAction.BackgroundColor]: The background color of the action button.
//   - [INSTableViewRowAction.SetBackgroundColor]
//
// # Instance Properties
//
//   - [INSTableViewRowAction.Image]
//   - [INSTableViewRowAction.SetImage]
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction
type INSTableViewRowAction interface {
	objectivec.IObject

	// Topic: Configuring the Action’s Appearance

	// The style applied to the action button.
	Style() NSTableViewRowActionStyle
	// The title of the action button.
	Title() string
	SetTitle(value string)
	// The background color of the action button.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)

	// Topic: Instance Properties

	Image() INSImage
	SetImage(value INSImage)
}

// Init initializes the instance.
func (t NSTableViewRowAction) Init() NSTableViewRowAction {
	rv := objc.Send[NSTableViewRowAction](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTableViewRowAction) Autorelease() NSTableViewRowAction {
	rv := objc.Send[NSTableViewRowAction](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTableViewRowAction creates a new NSTableViewRowAction instance.
func NewNSTableViewRowAction() NSTableViewRowAction {
	class := getNSTableViewRowActionClass()
	rv := objc.Send[NSTableViewRowAction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns a new table view row action object.
//
// style: The style characteristics to apply to the button. Use this value to apply
// default appearance characteristics to the button. These characteristics
// visually communicate, such as by color, information about what the button
// does. For example, specify a style of
// [NSTableViewRowAction.Style.destructive] to indicate an action is
// destructive to the underlying data. For a list of possible style values,
// see [NSTableViewRowAction.Style].
// //
// [NSTableViewRowAction.Style.destructive]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/Style-swift.enum/destructive
// [NSTableViewRowAction.Style]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/Style-swift.enum
//
// title: The string to display in the button. Specify a string localized for the
// user’s current language.
//
// handler: The block to execute when the user clicks the button associated with this
// action. AppKit makes a copy of the block you provide. When the user selects
// the action represented by this object, AppKit executes your `handler` block
// on the app’s main thread. This parameter must not be `nil`. This block
// has no return value and takes the following parameters:
// 
// action: The action object representing the action that the user selected.
// indexPath: The table row that the user acted on.
//
// # Return Value
// 
// A new table row action object that you can return from your table view’s
// delegate method.
//
// # Discussion
// 
// The style and handler block you specify cannot be changed later. You can
// change the title of the action button. You can also configure other
// appearance-related properties of the button using the properties of this
// class.
// 
// You can assign the same row action object to multiple rows of your table.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/init(style:title:handler:)
func (_NSTableViewRowActionClass NSTableViewRowActionClass) RowActionWithStyleTitleHandler(style NSTableViewRowActionStyle, title string, handler TableViewRowActionHandler) NSTableViewRowAction {
_block2, _ := NewTableViewRowActionBlock(handler)
	rv := objc.Send[objc.ID](objc.ID(_NSTableViewRowActionClass.class), objc.Sel("rowActionWithStyle:title:handler:"), style, objc.String(title), _block2)
	return NSTableViewRowActionFromID(rv)
}

// The style applied to the action button.
//
// # Discussion
// 
// The value of this property is set at creation time and cannot be changed
// later.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/style-swift.property
func (t NSTableViewRowAction) Style() NSTableViewRowActionStyle {
	rv := objc.Send[NSTableViewRowActionStyle](t.ID, objc.Sel("style"))
	return NSTableViewRowActionStyle(rv)
}
// The title of the action button.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/title
func (t NSTableViewRowAction) Title() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTableViewRowAction) SetTitle(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setTitle:"), objc.String(value))
}
// The background color of the action button.
//
// # Discussion
// 
// Use this property to specify the background color for your button. If you
// do not specify a value for this property, AppKit assigns a default color
// based on the value in the [Style] property. Generally, this color is red
// for destructive actions and blue for nondestructive actions.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/backgroundColor
func (t NSTableViewRowAction) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTableViewRowAction) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setBackgroundColor:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/image
func (t NSTableViewRowAction) Image() INSImage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (t NSTableViewRowAction) SetImage(value INSImage) {
	objc.Send[struct{}](t.ID, objc.Sel("setImage:"), value)
}


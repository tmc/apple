// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTableCellView] class.
var (
	_NSTableCellViewClass     NSTableCellViewClass
	_NSTableCellViewClassOnce sync.Once
)

func getNSTableCellViewClass() NSTableCellViewClass {
	_NSTableCellViewClassOnce.Do(func() {
		_NSTableCellViewClass = NSTableCellViewClass{class: objc.GetClass("NSTableCellView")}
	})
	return _NSTableCellViewClass
}

// GetNSTableCellViewClass returns the class object for NSTableCellView.
func GetNSTableCellViewClass() NSTableCellViewClass {
	return getNSTableCellViewClass()
}

type NSTableCellViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTableCellViewClass) Alloc() NSTableCellView {
	rv := objc.Send[NSTableCellView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A reusable container view shown for a particular cell in a table view that
// uses rows for content.
//
// # Overview
// 
// The [ImageView] and [TextField] properties are connected in Interface
// Builder. Additional properties can be added by subclassing
// [NSTableCellView] and adding the required properties and connecting them
// programmatically or in Interface Builder.
// 
// The `objectValue` is used when setting the value of the view cell by the
// [TableViewObjectValueForTableColumnRow] method in the
// [NSTableViewDataSource]. If you use your own custom view cells that are not
// based on [NSTableCellView] you should implement this property in order to
// be able to receive changes to cell values.
//
// # Represented Object
//
//   - [NSTableCellView.ObjectValue]: The object that represents the cell data.
//   - [NSTableCellView.SetObjectValue]
//
// # Displayed Items
//
//   - [NSTableCellView.ImageView]: Image displayed by the cell.
//   - [NSTableCellView.SetImageView]
//   - [NSTableCellView.TextField]: Text displayed by the cell.
//   - [NSTableCellView.SetTextField]
//
// # Getting and Setting the Background Style
//
//   - [NSTableCellView.BackgroundStyle]: This property is automatically set by the enclosing row view to let this view know what its background looks like.
//   - [NSTableCellView.SetBackgroundStyle]
//
// # Getting and Setting the Row Size Style
//
//   - [NSTableCellView.RowSizeStyle]: Returns the row size style.
//   - [NSTableCellView.SetRowSizeStyle]
//
// # Dragging Images
//
//   - [NSTableCellView.DraggingImageComponents]: Returns dragging images for the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableCellView
type NSTableCellView struct {
	NSView
}

// NSTableCellViewFromID constructs a [NSTableCellView] from an objc.ID.
//
// A reusable container view shown for a particular cell in a table view that
// uses rows for content.
func NSTableCellViewFromID(id objc.ID) NSTableCellView {
	return NSTableCellView{NSView: NSViewFromID(id)}
}
// NOTE: NSTableCellView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTableCellView] class.
//
// # Represented Object
//
//   - [INSTableCellView.ObjectValue]: The object that represents the cell data.
//   - [INSTableCellView.SetObjectValue]
//
// # Displayed Items
//
//   - [INSTableCellView.ImageView]: Image displayed by the cell.
//   - [INSTableCellView.SetImageView]
//   - [INSTableCellView.TextField]: Text displayed by the cell.
//   - [INSTableCellView.SetTextField]
//
// # Getting and Setting the Background Style
//
//   - [INSTableCellView.BackgroundStyle]: This property is automatically set by the enclosing row view to let this view know what its background looks like.
//   - [INSTableCellView.SetBackgroundStyle]
//
// # Getting and Setting the Row Size Style
//
//   - [INSTableCellView.RowSizeStyle]: Returns the row size style.
//   - [INSTableCellView.SetRowSizeStyle]
//
// # Dragging Images
//
//   - [INSTableCellView.DraggingImageComponents]: Returns dragging images for the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableCellView
type INSTableCellView interface {
	INSView

	// Topic: Represented Object

	// The object that represents the cell data.
	ObjectValue() objectivec.IObject
	SetObjectValue(value objectivec.IObject)

	// Topic: Displayed Items

	// Image displayed by the cell.
	ImageView() INSImageView
	SetImageView(value INSImageView)
	// Text displayed by the cell.
	TextField() INSTextField
	SetTextField(value INSTextField)

	// Topic: Getting and Setting the Background Style

	// This property is automatically set by the enclosing row view to let this view know what its background looks like.
	BackgroundStyle() NSBackgroundStyle
	SetBackgroundStyle(value NSBackgroundStyle)

	// Topic: Getting and Setting the Row Size Style

	// Returns the row size style.
	RowSizeStyle() NSTableViewRowSizeStyle
	SetRowSizeStyle(value NSTableViewRowSizeStyle)

	// Topic: Dragging Images

	// Returns dragging images for the cell.
	DraggingImageComponents() []NSDraggingImageComponent
}

// Init initializes the instance.
func (t NSTableCellView) Init() NSTableCellView {
	rv := objc.Send[NSTableCellView](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTableCellView) Autorelease() NSTableCellView {
	rv := objc.Send[NSTableCellView](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTableCellView creates a new NSTableCellView instance.
func NewNSTableCellView() NSTableCellView {
	class := getNSTableCellViewClass()
	rv := objc.Send[NSTableCellView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a view using from data in the specified coder object.
//
// coder: The coder object that contains the view’s configuration details.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewTableCellViewWithCoder(coder foundation.INSCoder) NSTableCellView {
	instance := getNSTableCellViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTableCellViewFromID(rv)
}

// Initializes and returns a newly allocated [NSView] object with a specified
// frame rectangle.
//
// frameRect: The frame rectangle for the created view object.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// # Discussion
// 
// Insert the view into your window’s view hieararchy before you can do
// anything with it. This method is the designated initializer for the
// [NSView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewTableCellViewWithFrame(frameRect corefoundation.CGRect) NSTableCellView {
	instance := getNSTableCellViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSTableCellViewFromID(rv)
}

// The object that represents the cell data.
//
// # Discussion
// 
// The `objectValue` is automatically set by the table when using bindings or
// is the object returned by the [NSTableViewDataSource] protocol method
// [TableViewObjectValueForTableColumnRow].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableCellView/objectValue
func (t NSTableCellView) ObjectValue() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("objectValue"))
	return objectivec.Object{ID: rv}
}
func (t NSTableCellView) SetObjectValue(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setObjectValue:"), value)
}
// Image displayed by the cell.
//
// # Discussion
// 
// This property is typically configured when the row is created in the
// [NSTableViewDataSource] protocol method [TableViewViewForTableColumnRow].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableCellView/imageView
func (t NSTableCellView) ImageView() INSImageView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("imageView"))
	return NSImageViewFromID(objc.ID(rv))
}
func (t NSTableCellView) SetImageView(value INSImageView) {
	objc.Send[struct{}](t.ID, objc.Sel("setImageView:"), value)
}
// Text displayed by the cell.
//
// # Discussion
// 
// This property is typically configured when the row is created in the
// [NSTableViewDelegate] protocol method [TableViewViewForTableColumnRow].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableCellView/textField
func (t NSTableCellView) TextField() INSTextField {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textField"))
	return NSTextFieldFromID(objc.ID(rv))
}
func (t NSTableCellView) SetTextField(value INSTextField) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextField:"), value)
}
// This property is automatically set by the enclosing row view to let this
// view know what its background looks like.
//
// # Discussion
// 
// The property is automatically set by the enclosing [NSTableRowView] to let
// this view know what its background looks like. For instance, when the
// `backgroundStyle` is NSBackgroundStyleDark, the view should use a light
// text color.
// 
// The default implementation automatically forwards calls to all subviews
// that implement `` or are an NSControl, which have [NSCell] classes that
// respond to [BackgroundStyle].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableCellView/backgroundStyle
func (t NSTableCellView) BackgroundStyle() NSBackgroundStyle {
	rv := objc.Send[NSBackgroundStyle](t.ID, objc.Sel("backgroundStyle"))
	return NSBackgroundStyle(rv)
}
func (t NSTableCellView) SetBackgroundStyle(value NSBackgroundStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setBackgroundStyle:"), value)
}
// Returns the row size style.
//
// # Discussion
// 
// The `rowSizeStyle` property is set by the [NSTableView] to its
// [EffectiveRowSizeStyle]. The cell view will layout the [TextField] and
// [ImageView] based on the `rowSizeStyle`.
// 
// A value of [NSTableView.RowSizeStyle.default] should never be set on the
// cell view, as it is an appropriate value only for the table as it returns
// the effective row size style for the table.
//
// [NSTableView.RowSizeStyle.default]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/default
//
// See: https://developer.apple.com/documentation/AppKit/NSTableCellView/rowSizeStyle
func (t NSTableCellView) RowSizeStyle() NSTableViewRowSizeStyle {
	rv := objc.Send[NSTableViewRowSizeStyle](t.ID, objc.Sel("rowSizeStyle"))
	return NSTableViewRowSizeStyle(rv)
}
func (t NSTableCellView) SetRowSizeStyle(value NSTableViewRowSizeStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setRowSizeStyle:"), value)
}
// Returns dragging images for the cell.
//
// # Discussion
// 
// The default implementation of this method returns an array of up to two
// [NSDraggingImageComponent] instances – one for the [ImageView] and
// another for the [TextField] (unless the property is `nil`).
// 
// These method can be subclassed and overridden to provide a custom set of
// [NSDraggingImageComponent] objects to create the drag image from this view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableCellView/draggingImageComponents
func (t NSTableCellView) DraggingImageComponents() []NSDraggingImageComponent {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("draggingImageComponents"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSDraggingImageComponent {
		return NSDraggingImageComponentFromID(id)
	})
}


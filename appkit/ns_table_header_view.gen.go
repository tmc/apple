// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTableHeaderView] class.
var (
	_NSTableHeaderViewClass     NSTableHeaderViewClass
	_NSTableHeaderViewClassOnce sync.Once
)

func getNSTableHeaderViewClass() NSTableHeaderViewClass {
	_NSTableHeaderViewClassOnce.Do(func() {
		_NSTableHeaderViewClass = NSTableHeaderViewClass{class: objc.GetClass("NSTableHeaderView")}
	})
	return _NSTableHeaderViewClass
}

// GetNSTableHeaderViewClass returns the class object for NSTableHeaderView.
func GetNSTableHeaderViewClass() NSTableHeaderViewClass {
	return getNSTableHeaderViewClass()
}

type NSTableHeaderViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTableHeaderViewClass) Alloc() NSTableHeaderView {
	rv := objc.Send[NSTableHeaderView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that draws headers over a table view’s columns and handles
// mouse events in those headers.
//
// # Overview
// 
// [NSTableHeaderView] uses [NSTableHeaderCell] to implement its user
// interface.
//
// # Setting the table view
//
//   - [NSTableHeaderView.TableView]: The [NSTableView](<doc://com.apple.appkit/documentation/AppKit/NSTableView>) instance that this table header view belongs to.
//   - [NSTableHeaderView.SetTableView]
//
// # Checking altered columns
//
//   - [NSTableHeaderView.DraggedColumn]: The index of the column that the user is dragging.
//   - [NSTableHeaderView.DraggedDistance]: The horizontal distance that the user has dragged a column.
//   - [NSTableHeaderView.ResizedColumn]: The index of the column that the user is resizing.
//
// # Utility methods
//
//   - [NSTableHeaderView.ColumnAtPoint]: Returns the index of the column whose header lies under `aPoint` in the receiver, or –1 if no such column is found.
//   - [NSTableHeaderView.HeaderRectOfColumn]: Returns the rectangle containing the header tile for the column at `columnIndex`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderView
type NSTableHeaderView struct {
	NSView
}

// NSTableHeaderViewFromID constructs a [NSTableHeaderView] from an objc.ID.
//
// An object that draws headers over a table view’s columns and handles
// mouse events in those headers.
func NSTableHeaderViewFromID(id objc.ID) NSTableHeaderView {
	return NSTableHeaderView{NSView: NSViewFromID(id)}
}
// NOTE: NSTableHeaderView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTableHeaderView] class.
//
// # Setting the table view
//
//   - [INSTableHeaderView.TableView]: The [NSTableView](<doc://com.apple.appkit/documentation/AppKit/NSTableView>) instance that this table header view belongs to.
//   - [INSTableHeaderView.SetTableView]
//
// # Checking altered columns
//
//   - [INSTableHeaderView.DraggedColumn]: The index of the column that the user is dragging.
//   - [INSTableHeaderView.DraggedDistance]: The horizontal distance that the user has dragged a column.
//   - [INSTableHeaderView.ResizedColumn]: The index of the column that the user is resizing.
//
// # Utility methods
//
//   - [INSTableHeaderView.ColumnAtPoint]: Returns the index of the column whose header lies under `aPoint` in the receiver, or –1 if no such column is found.
//   - [INSTableHeaderView.HeaderRectOfColumn]: Returns the rectangle containing the header tile for the column at `columnIndex`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderView
type INSTableHeaderView interface {
	INSView
	NSViewToolTipOwner

	// Topic: Setting the table view

	// The [NSTableView](<doc://com.apple.appkit/documentation/AppKit/NSTableView>) instance that this table header view belongs to.
	TableView() INSTableView
	SetTableView(value INSTableView)

	// Topic: Checking altered columns

	// The index of the column that the user is dragging.
	DraggedColumn() int
	// The horizontal distance that the user has dragged a column.
	DraggedDistance() float64
	// The index of the column that the user is resizing.
	ResizedColumn() int

	// Topic: Utility methods

	// Returns the index of the column whose header lies under `aPoint` in the receiver, or –1 if no such column is found.
	ColumnAtPoint(point corefoundation.CGPoint) int
	// Returns the rectangle containing the header tile for the column at `columnIndex`.
	HeaderRectOfColumn(column int) corefoundation.CGRect
}

// Init initializes the instance.
func (t NSTableHeaderView) Init() NSTableHeaderView {
	rv := objc.Send[NSTableHeaderView](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTableHeaderView) Autorelease() NSTableHeaderView {
	rv := objc.Send[NSTableHeaderView](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTableHeaderView creates a new NSTableHeaderView instance.
func NewNSTableHeaderView() NSTableHeaderView {
	class := getNSTableHeaderViewClass()
	rv := objc.Send[NSTableHeaderView](objc.ID(class.class), objc.Sel("new"))
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
func NewTableHeaderViewWithCoder(coder foundation.INSCoder) NSTableHeaderView {
	instance := getNSTableHeaderViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTableHeaderViewFromID(rv)
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
func NewTableHeaderViewWithFrame(frameRect corefoundation.CGRect) NSTableHeaderView {
	instance := getNSTableHeaderViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSTableHeaderViewFromID(rv)
}

// Returns the index of the column whose header lies under `aPoint` in the
// receiver, or –1 if no such column is found.
//
// # Discussion
// 
// `aPoint` is expressed in the receiver’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/column(at:)
func (t NSTableHeaderView) ColumnAtPoint(point corefoundation.CGPoint) int {
	rv := objc.Send[int](t.ID, objc.Sel("columnAtPoint:"), point)
	return rv
}
// Returns the rectangle containing the header tile for the column at
// `columnIndex`.
//
// # Discussion
// 
// Raises an [NSInternalInconsistencyException] if `columnIndex` is out of
// bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/headerRect(ofColumn:)
func (t NSTableHeaderView) HeaderRectOfColumn(column int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("headerRectOfColumn:"), column)
	return corefoundation.CGRect(rv)
}
// Returns the tool tip string to be displayed due to the cursor pausing at
// location `point` within the tool tip rectangle identified by `tag` in the
// view `view`.
//
// # Discussion
// 
// `userData` is additional information provided by the creator of the tool
// tip rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewToolTipOwner/view(_:stringForToolTip:point:userData:)
func (t NSTableHeaderView) ViewStringForToolTipPointUserData(view INSView, tag objectivec.IObject, point corefoundation.CGPoint, data unsafe.Pointer) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("view:stringForToolTip:point:userData:"), view, tag, point, data)
	return foundation.NSStringFromID(rv).String()
}

// The [NSTableView] instance that this table header view belongs to.
//
// # Discussion
// 
// You should never need to set this property; it’s assigned automatically
// when you set the header view for an [NSTableView].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/tableView
func (t NSTableHeaderView) TableView() INSTableView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tableView"))
	return NSTableViewFromID(objc.ID(rv))
}
func (t NSTableHeaderView) SetTableView(value INSTableView) {
	objc.Send[struct{}](t.ID, objc.Sel("setTableView:"), value)
}
// The index of the column that the user is dragging.
//
// # Discussion
// 
// If the user is dragging a column, this property contains the index of that
// column; otherwise, it contains `-1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/draggedColumn
func (t NSTableHeaderView) DraggedColumn() int {
	rv := objc.Send[int](t.ID, objc.Sel("draggedColumn"))
	return rv
}
// The horizontal distance that the user has dragged a column.
//
// # Discussion
// 
// If the user is dragging a column, this property contains that column’s
// horizontal distance from its original position; otherwise, the property’s
// value is undefined.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/draggedDistance
func (t NSTableHeaderView) DraggedDistance() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("draggedDistance"))
	return rv
}
// The index of the column that the user is resizing.
//
// # Discussion
// 
// If the user is resizing a column, this property contains the index of that
// column; otherwise, it contains `-1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderView/resizedColumn
func (t NSTableHeaderView) ResizedColumn() int {
	rv := objc.Send[int](t.ID, objc.Sel("resizedColumn"))
	return rv
}

			// Protocol methods for NSViewToolTipOwner
			


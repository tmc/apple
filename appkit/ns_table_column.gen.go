// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTableColumn] class.
var (
	_NSTableColumnClass     NSTableColumnClass
	_NSTableColumnClassOnce sync.Once
)

func getNSTableColumnClass() NSTableColumnClass {
	_NSTableColumnClassOnce.Do(func() {
		_NSTableColumnClass = NSTableColumnClass{class: objc.GetClass("NSTableColumn")}
	})
	return _NSTableColumnClass
}

// GetNSTableColumnClass returns the class object for NSTableColumn.
func GetNSTableColumnClass() NSTableColumnClass {
	return getNSTableColumnClass()
}

type NSTableColumnClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTableColumnClass) Alloc() NSTableColumn {
	rv := objc.Send[NSTableColumn](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The display characteristics and identifier for a column in a table view.
//
// # Overview
// 
// A table column object determines the width (including the maximum and
// minimum widths) of its column in the table view and specifies the
// column’s resizing and editing behavior. A table column stores two cell
// objects: the header cell, which is used to draw the column header, and the
// data cell, which is used to draw the values for each row. In a cell-based
// table, you can control the display of the column by specifying subclasses
// of [NSCell] to use and by setting the font and other display
// characteristics for these cells. For example, you can use an
// [NSTextFieldCell] to display string values or substitute an [NSImageCell]
// to display pictures.
//
// # Creating a Table Column
//
//   - [NSTableColumn.InitWithIdentifier]: Initializes a newly created table column with a string identifier.
//
// # Setting the Table View
//
//   - [NSTableColumn.TableView]: The table view that contains the table column.
//   - [NSTableColumn.SetTableView]
//
// # Controlling Size
//
//   - [NSTableColumn.Width]: The table column’s width, in points.
//   - [NSTableColumn.SetWidth]
//   - [NSTableColumn.MinWidth]: The table column’s minimum width, in points.
//   - [NSTableColumn.SetMinWidth]
//   - [NSTableColumn.MaxWidth]: The table column’s maximum width, in points.
//   - [NSTableColumn.SetMaxWidth]
//   - [NSTableColumn.ResizingMask]: The table column’s resizing mask.
//   - [NSTableColumn.SetResizingMask]
//   - [NSTableColumn.SizeToFit]: Resizes the table column to fit the width of its header cell.
//
// # Setting the Header
//
//   - [NSTableColumn.Title]: The title of the table column’s header.
//   - [NSTableColumn.SetTitle]
//   - [NSTableColumn.HeaderCell]: The cell used to draw the table column’s header.
//   - [NSTableColumn.SetHeaderCell]
//
// # Controlling Editability in a Cell-Based Table
//
//   - [NSTableColumn.Editable]: A Boolean that indicates whether a cell-based table’s column cells are user editable.
//   - [NSTableColumn.SetEditable]
//
// # Sorting
//
//   - [NSTableColumn.SortDescriptorPrototype]: The table column’s sort descriptor prototype.
//   - [NSTableColumn.SetSortDescriptorPrototype]
//
// # Setting Column Visibility
//
//   - [NSTableColumn.Hidden]: A Boolean that indicates whether the table column is hidden.
//   - [NSTableColumn.SetHidden]
//
// # Setting Tooltips
//
//   - [NSTableColumn.HeaderToolTip]: The string that’s displayed in a help tag over the table column header.
//   - [NSTableColumn.SetHeaderToolTip]
//
// # Deprecated Methods
//
//   - [NSTableColumn.DataCell]: The cell prototype used by the table column to draw individual cells.
//   - [NSTableColumn.SetDataCell]
//   - [NSTableColumn.DataCellForRow]: Returns the cell object used to display values in the specified row of the table column.
//
// # Initializers
//
//   - [NSTableColumn.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn
type NSTableColumn struct {
	objectivec.Object
}

// NSTableColumnFromID constructs a [NSTableColumn] from an objc.ID.
//
// The display characteristics and identifier for a column in a table view.
func NSTableColumnFromID(id objc.ID) NSTableColumn {
	return NSTableColumn{objectivec.Object{ID: id}}
}
// NOTE: NSTableColumn adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTableColumn] class.
//
// # Creating a Table Column
//
//   - [INSTableColumn.InitWithIdentifier]: Initializes a newly created table column with a string identifier.
//
// # Setting the Table View
//
//   - [INSTableColumn.TableView]: The table view that contains the table column.
//   - [INSTableColumn.SetTableView]
//
// # Controlling Size
//
//   - [INSTableColumn.Width]: The table column’s width, in points.
//   - [INSTableColumn.SetWidth]
//   - [INSTableColumn.MinWidth]: The table column’s minimum width, in points.
//   - [INSTableColumn.SetMinWidth]
//   - [INSTableColumn.MaxWidth]: The table column’s maximum width, in points.
//   - [INSTableColumn.SetMaxWidth]
//   - [INSTableColumn.ResizingMask]: The table column’s resizing mask.
//   - [INSTableColumn.SetResizingMask]
//   - [INSTableColumn.SizeToFit]: Resizes the table column to fit the width of its header cell.
//
// # Setting the Header
//
//   - [INSTableColumn.Title]: The title of the table column’s header.
//   - [INSTableColumn.SetTitle]
//   - [INSTableColumn.HeaderCell]: The cell used to draw the table column’s header.
//   - [INSTableColumn.SetHeaderCell]
//
// # Controlling Editability in a Cell-Based Table
//
//   - [INSTableColumn.Editable]: A Boolean that indicates whether a cell-based table’s column cells are user editable.
//   - [INSTableColumn.SetEditable]
//
// # Sorting
//
//   - [INSTableColumn.SortDescriptorPrototype]: The table column’s sort descriptor prototype.
//   - [INSTableColumn.SetSortDescriptorPrototype]
//
// # Setting Column Visibility
//
//   - [INSTableColumn.Hidden]: A Boolean that indicates whether the table column is hidden.
//   - [INSTableColumn.SetHidden]
//
// # Setting Tooltips
//
//   - [INSTableColumn.HeaderToolTip]: The string that’s displayed in a help tag over the table column header.
//   - [INSTableColumn.SetHeaderToolTip]
//
// # Deprecated Methods
//
//   - [INSTableColumn.DataCell]: The cell prototype used by the table column to draw individual cells.
//   - [INSTableColumn.SetDataCell]
//   - [INSTableColumn.DataCellForRow]: Returns the cell object used to display values in the specified row of the table column.
//
// # Initializers
//
//   - [INSTableColumn.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn
type INSTableColumn interface {
	objectivec.IObject
	NSUserInterfaceItemIdentification

	// Topic: Creating a Table Column

	// Initializes a newly created table column with a string identifier.
	InitWithIdentifier(identifier NSUserInterfaceItemIdentifier) NSTableColumn

	// Topic: Setting the Table View

	// The table view that contains the table column.
	TableView() INSTableView
	SetTableView(value INSTableView)

	// Topic: Controlling Size

	// The table column’s width, in points.
	Width() float64
	SetWidth(value float64)
	// The table column’s minimum width, in points.
	MinWidth() float64
	SetMinWidth(value float64)
	// The table column’s maximum width, in points.
	MaxWidth() float64
	SetMaxWidth(value float64)
	// The table column’s resizing mask.
	ResizingMask() NSTableColumnResizingOptions
	SetResizingMask(value NSTableColumnResizingOptions)
	// Resizes the table column to fit the width of its header cell.
	SizeToFit()

	// Topic: Setting the Header

	// The title of the table column’s header.
	Title() string
	SetTitle(value string)
	// The cell used to draw the table column’s header.
	HeaderCell() INSTableHeaderCell
	SetHeaderCell(value INSTableHeaderCell)

	// Topic: Controlling Editability in a Cell-Based Table

	// A Boolean that indicates whether a cell-based table’s column cells are user editable.
	Editable() bool
	SetEditable(value bool)

	// Topic: Sorting

	// The table column’s sort descriptor prototype.
	SortDescriptorPrototype() foundation.INSSortDescriptor
	SetSortDescriptorPrototype(value foundation.INSSortDescriptor)

	// Topic: Setting Column Visibility

	// A Boolean that indicates whether the table column is hidden.
	Hidden() bool
	SetHidden(value bool)

	// Topic: Setting Tooltips

	// The string that’s displayed in a help tag over the table column header.
	HeaderToolTip() string
	SetHeaderToolTip(value string)

	// Topic: Deprecated Methods

	// The cell prototype used by the table column to draw individual cells.
	DataCell() objectivec.IObject
	SetDataCell(value objectivec.IObject)
	// Returns the cell object used to display values in the specified row of the table column.
	DataCellForRow(row int) objectivec.IObject

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) NSTableColumn

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTableColumn) Init() NSTableColumn {
	rv := objc.Send[NSTableColumn](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTableColumn) Autorelease() NSTableColumn {
	rv := objc.Send[NSTableColumn](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTableColumn creates a new NSTableColumn instance.
func NewNSTableColumn() NSTableColumn {
	class := getNSTableColumnClass()
	rv := objc.Send[NSTableColumn](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/init(coder:)
func NewTableColumnWithCoder(coder foundation.INSCoder) NSTableColumn {
	instance := getNSTableColumnClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTableColumnFromID(rv)
}

// Initializes a newly created table column with a string identifier.
//
// identifier: The string identifier for the column.
//
// # Return Value
// 
// An initialized table column instance with an [NSTextFieldCell] instance as
// its default cell.
//
// # Discussion
// 
// You can set the table column title using the [Title] property.
// 
// This method is the designated initializer for the [NSTableColumn] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/init(identifier:)
func NewTableColumnWithIdentifier(identifier NSUserInterfaceItemIdentifier) NSTableColumn {
	instance := getNSTableColumnClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSTableColumnFromID(rv)
}

// Initializes a newly created table column with a string identifier.
//
// identifier: The string identifier for the column.
//
// # Return Value
// 
// An initialized table column instance with an [NSTextFieldCell] instance as
// its default cell.
//
// # Discussion
// 
// You can set the table column title using the [Title] property.
// 
// This method is the designated initializer for the [NSTableColumn] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/init(identifier:)
func (t NSTableColumn) InitWithIdentifier(identifier NSUserInterfaceItemIdentifier) NSTableColumn {
	rv := objc.Send[NSTableColumn](t.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return rv
}
// Resizes the table column to fit the width of its header cell.
//
// # Discussion
// 
// If the table column’s maximum width is less than the width of the header,
// the maximum is increased to the header’s width. Similarly, if the table
// column’s minimum width is greater than the width of the header, the
// minimum is reduced to the header’s width.
// 
// If this method causes the table column’s width to change, the column’s
// table view is marked as needing display.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/sizeToFit()
func (t NSTableColumn) SizeToFit() {
	objc.Send[objc.ID](t.ID, objc.Sel("sizeToFit"))
}
// Returns the cell object used to display values in the specified row of the
// table column.
//
// row: The table column row.
//
// # Return Value
// 
// The data cell object.
//
// # Discussion
// 
// Returns the [NSCell] object used by the table view to draw values for the
// receiver. The table view calls this method when drawing the row, so you
// shouldn’t need to call it directly. By default, this method just accesses
// [DataCell].
// 
// To enable per-row customization of the cell used by the table column, you
// can override this method or use the [NSTableViewDelegate] method
// [TableViewDataCellForTableColumnRow]. In both cases, the cell that’s
// returned should properly implement [copy(with:)], because the table view
// may copy the cell during certain operations.
// 
// Subclasses should be prepared for this method to be called with `row` equal
// to –1 in cases where no actual row is involved but the table view needs
// to get some generic cell information.
//
// [copy(with:)]: https://developer.apple.com/documentation/Foundation/NSCopying/copy(with:)
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/dataCell(forRow:)
func (t NSTableColumn) DataCellForRow(row int) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dataCellForRow:"), row)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/init(coder:)
func (t NSTableColumn) InitWithCoder(coder foundation.INSCoder) NSTableColumn {
	rv := objc.Send[NSTableColumn](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (t NSTableColumn) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The table view that contains the table column.
//
// # Discussion
// 
// You should never need to set this property; it’s set automatically when
// you add a table column to a table view using the [NSTableView] class’s
// method [AddTableColumn].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/tableView
func (t NSTableColumn) TableView() INSTableView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tableView"))
	return NSTableViewFromID(objc.ID(rv))
}
func (t NSTableColumn) SetTableView(value INSTableView) {
	objc.Send[struct{}](t.ID, objc.Sel("setTableView:"), value)
}
// The table column’s width, in points.
//
// # Discussion
// 
// The default value of this property is `100.0`.
// 
// If the value of this property exceeds the minimum or maximum width, it’s
// adjusted to the appropriate limiting value.
// 
// This property posts [columnDidResizeNotification] on behalf of the table
// column’s [NSTableView] and marks the table view as needing display.
//
// [columnDidResizeNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/columnDidResizeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/width
func (t NSTableColumn) Width() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("width"))
	return rv
}
func (t NSTableColumn) SetWidth(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setWidth:"), value)
}
// The table column’s minimum width, in points.
//
// # Discussion
// 
// The default value of this property is `10.0`.
// 
// The table column width can’t be less than the value of this property,
// whether the column is resized by the user or programmatically. If the table
// column’s current width is less than the value of this property, the width
// is set to the value of this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/minWidth
func (t NSTableColumn) MinWidth() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("minWidth"))
	return rv
}
func (t NSTableColumn) SetMinWidth(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setMinWidth:"), value)
}
// The table column’s maximum width, in points.
//
// # Discussion
// 
// The default value of this property is [MAXFLOAT].
// 
// The table column width can’t be greater than the value of this property,
// whether the column is resized by the user or programmatically. If the table
// column’s current width is greater than the value of this property, the
// width is set to the value of this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/maxWidth
func (t NSTableColumn) MaxWidth() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("maxWidth"))
	return rv
}
func (t NSTableColumn) SetMaxWidth(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setMaxWidth:"), value)
}
// The table column’s resizing mask.
//
// # Discussion
// 
// The value of this property specifies the resizability of the table column.
// See [Resizing Modes] for possible values. Values can be combined using the
// C bitwise OR operator.
// 
// When the value of this property is `0`, the column is not resizable. The
// default value of this property is [TableColumnUserResizingMask] |
// [TableColumnAutoresizingMask].
//
// [Resizing Modes]: https://developer.apple.com/documentation/AppKit/resizing-modes
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/resizingMask
func (t NSTableColumn) ResizingMask() NSTableColumnResizingOptions {
	rv := objc.Send[NSTableColumnResizingOptions](t.ID, objc.Sel("resizingMask"))
	return NSTableColumnResizingOptions(rv)
}
func (t NSTableColumn) SetResizingMask(value NSTableColumnResizingOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setResizingMask:"), value)
}
// The title of the table column’s header.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/title
func (t NSTableColumn) Title() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTableColumn) SetTitle(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setTitle:"), objc.String(value))
}
// The cell used to draw the table column’s header.
//
// # Discussion
// 
// The value of this property must not be `nil`. It’s recommended that the
// value of this property be an instance or subclass of [NSTableHeaderCell].
// 
// You can set the table column title using the [Title] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/headerCell
func (t NSTableColumn) HeaderCell() INSTableHeaderCell {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("headerCell"))
	return NSTableHeaderCellFromID(objc.ID(rv))
}
func (t NSTableColumn) SetHeaderCell(value INSTableHeaderCell) {
	objc.Send[struct{}](t.ID, objc.Sel("setHeaderCell:"), value)
}
// The identifier string for the table column.
//
// # Discussion
// 
// This string is used by the data source to identify the table column.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/identifier
func (t NSTableColumn) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (t NSTableColumn) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](t.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}
// A Boolean that indicates whether a cell-based table’s column cells are
// user editable.
//
// # Discussion
// 
// When the value of this property is [true], the user can edit cells in the
// cell-based table’s column. The default value of this property is [true].
// 
// To initiate editing programmatically regardless of the value of this
// property, use the [NSTableView] [EditColumnRowWithEventSelect] method.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/isEditable
func (t NSTableColumn) Editable() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isEditable"))
	return rv
}
func (t NSTableColumn) SetEditable(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setEditable:"), value)
}
// The table column’s sort descriptor prototype.
//
// # Discussion
// 
// A table column is considered sortable if it has a sort descriptor that
// specifies the sorting direction, a key to sort by, and a selector that
// defines how to sort.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/sortDescriptorPrototype
func (t NSTableColumn) SortDescriptorPrototype() foundation.INSSortDescriptor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("sortDescriptorPrototype"))
	return foundation.NSSortDescriptorFromID(objc.ID(rv))
}
func (t NSTableColumn) SetSortDescriptorPrototype(value foundation.INSSortDescriptor) {
	objc.Send[struct{}](t.ID, objc.Sel("setSortDescriptorPrototype:"), value)
}
// A Boolean that indicates whether the table column is hidden.
//
// # Discussion
// 
// When the value of this property is [true], the table column is hidden. The
// default value is [false].
// 
// Columns that are hidden still exist in the table view object’s
// [TableColumns] array and are included in the table view’s
// [NumberOfColumns] count.
// 
// The hidden state is stored when the table view autosaves the table column
// state.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/isHidden
func (t NSTableColumn) Hidden() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isHidden"))
	return rv
}
func (t NSTableColumn) SetHidden(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setHidden:"), value)
}
// The string that’s displayed in a help tag over the table column header.
//
// # Discussion
// 
// When the value of this property is `nil`, the table column header doesn’t
// display a help tag (also known as a tooltip). Otherwise, the string is
// displayed in a help tag when the pointer pauses over the header of the
// table column. The default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/headerToolTip
func (t NSTableColumn) HeaderToolTip() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("headerToolTip"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTableColumn) SetHeaderToolTip(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setHeaderToolTip:"), objc.String(value))
}
// The cell prototype used by the table column to draw individual cells.
//
// # Discussion
// 
// You can use this property to control the font, alignment, and other text
// attributes for the table column contents.
// 
// You can also assign a cell that displays things other than text—for
// example, you can display images by setting the cell to [NSImageCell].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableColumn/dataCell
func (t NSTableColumn) DataCell() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dataCell"))
	return objectivec.Object{ID: rv}
}
func (t NSTableColumn) SetDataCell(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setDataCell:"), value)
}

			// Protocol methods for NSUserInterfaceItemIdentification
			


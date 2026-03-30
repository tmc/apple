// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGridView] class.
var (
	_NSGridViewClass     NSGridViewClass
	_NSGridViewClassOnce sync.Once
)

func getNSGridViewClass() NSGridViewClass {
	_NSGridViewClassOnce.Do(func() {
		_NSGridViewClass = NSGridViewClass{class: objc.GetClass("NSGridView")}
	})
	return _NSGridViewClass
}

// GetNSGridViewClass returns the class object for NSGridView.
func GetNSGridViewClass() NSGridViewClass {
	return getNSGridViewClass()
}

type NSGridViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSGridViewClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGridViewClass) Alloc() NSGridView {
	rv := objc.Send[NSGridView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A container that aligns views in a flexible grid of rows and columns.
//
// # Overview
//
// A grid view helps you lay out content, such as photos or thumbnails, in a
// row-column arrangement similar to a spreadsheet. Within a grid view, an
// item that occupies a single row-column intersection is represented by an
// [NSGridCell] object.
//
// # Getting Information About the Grid
//
//   - [NSGridView.NumberOfRows]: The number of rows in the grid view.
//   - [NSGridView.NumberOfColumns]: The number of columns in the grid view.
//   - [NSGridView.IndexOfColumn]: Returns the index of the specified grid column.
//   - [NSGridView.RowAtIndex]: Returns the grid row object at the specified index.
//   - [NSGridView.ColumnAtIndex]: Returns the grid column object at the specified index.
//   - [NSGridView.IndexOfRow]: Returns the index of the specified grid row.
//
// # Adding, Removing, and Moving Rows
//
//   - [NSGridView.AddRowWithViews]: Adds an array of views to a new row.
//   - [NSGridView.InsertRowAtIndexWithViews]: Inserts the array of view objects into the grid view at the index.
//   - [NSGridView.RemoveRowAtIndex]: Removes the row from the grid view at the index.
//   - [NSGridView.MoveRowAtIndexToIndex]: Moves the specified row to the new row location.
//
// # Adding, Removing, and Moving Columns
//
//   - [NSGridView.AddColumnWithViews]: Adds a new column containing the array of views.
//   - [NSGridView.InsertColumnAtIndexWithViews]: Inserts the array of view objects at the specified index.
//   - [NSGridView.RemoveColumnAtIndex]: Removes the column from the grid view at the specified index.
//   - [NSGridView.MoveColumnAtIndexToIndex]: Moves the specified column to a new column location.
//
// # Managing Grid Spacing and Alignment
//
//   - [NSGridView.ColumnSpacing]: The column spacing for the grid view.
//   - [NSGridView.SetColumnSpacing]
//   - [NSGridView.RowSpacing]: The row spacing for the grid view.
//   - [NSGridView.SetRowSpacing]
//   - [NSGridView.RowAlignment]: The row alignment for the grid view.
//   - [NSGridView.SetRowAlignment]
//   - [NSGridView.XPlacement]: The placement of the cell within the grid column.
//   - [NSGridView.SetXPlacement]
//   - [NSGridView.YPlacement]: The placement of the cell within the grid row.
//   - [NSGridView.SetYPlacement]
//
// # Creating and Merging Cells
//
//   - [NSGridView.CellAtColumnIndexRowIndex]: Returns the grid cell object at the specified column and row index.
//   - [NSGridView.CellForView]: Returns the grid cell object that contains the given view or one of its ancestors.
//   - [NSGridView.MergeCellsInHorizontalRangeVerticalRange]: Expands the cell at the top-leading corner of the horizontal and vertical range to cover the entire area.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView
type NSGridView struct {
	NSView
}

// NSGridViewFromID constructs a [NSGridView] from an objc.ID.
//
// A container that aligns views in a flexible grid of rows and columns.
func NSGridViewFromID(id objc.ID) NSGridView {
	return NSGridView{NSView: NSViewFromID(id)}
}

// NOTE: NSGridView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGridView] class.
//
// # Getting Information About the Grid
//
//   - [INSGridView.NumberOfRows]: The number of rows in the grid view.
//   - [INSGridView.NumberOfColumns]: The number of columns in the grid view.
//   - [INSGridView.IndexOfColumn]: Returns the index of the specified grid column.
//   - [INSGridView.RowAtIndex]: Returns the grid row object at the specified index.
//   - [INSGridView.ColumnAtIndex]: Returns the grid column object at the specified index.
//   - [INSGridView.IndexOfRow]: Returns the index of the specified grid row.
//
// # Adding, Removing, and Moving Rows
//
//   - [INSGridView.AddRowWithViews]: Adds an array of views to a new row.
//   - [INSGridView.InsertRowAtIndexWithViews]: Inserts the array of view objects into the grid view at the index.
//   - [INSGridView.RemoveRowAtIndex]: Removes the row from the grid view at the index.
//   - [INSGridView.MoveRowAtIndexToIndex]: Moves the specified row to the new row location.
//
// # Adding, Removing, and Moving Columns
//
//   - [INSGridView.AddColumnWithViews]: Adds a new column containing the array of views.
//   - [INSGridView.InsertColumnAtIndexWithViews]: Inserts the array of view objects at the specified index.
//   - [INSGridView.RemoveColumnAtIndex]: Removes the column from the grid view at the specified index.
//   - [INSGridView.MoveColumnAtIndexToIndex]: Moves the specified column to a new column location.
//
// # Managing Grid Spacing and Alignment
//
//   - [INSGridView.ColumnSpacing]: The column spacing for the grid view.
//   - [INSGridView.SetColumnSpacing]
//   - [INSGridView.RowSpacing]: The row spacing for the grid view.
//   - [INSGridView.SetRowSpacing]
//   - [INSGridView.RowAlignment]: The row alignment for the grid view.
//   - [INSGridView.SetRowAlignment]
//   - [INSGridView.XPlacement]: The placement of the cell within the grid column.
//   - [INSGridView.SetXPlacement]
//   - [INSGridView.YPlacement]: The placement of the cell within the grid row.
//   - [INSGridView.SetYPlacement]
//
// # Creating and Merging Cells
//
//   - [INSGridView.CellAtColumnIndexRowIndex]: Returns the grid cell object at the specified column and row index.
//   - [INSGridView.CellForView]: Returns the grid cell object that contains the given view or one of its ancestors.
//   - [INSGridView.MergeCellsInHorizontalRangeVerticalRange]: Expands the cell at the top-leading corner of the horizontal and vertical range to cover the entire area.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView
type INSGridView interface {
	INSView

	// Topic: Getting Information About the Grid

	// The number of rows in the grid view.
	NumberOfRows() int
	// The number of columns in the grid view.
	NumberOfColumns() int
	// Returns the index of the specified grid column.
	IndexOfColumn(column INSGridColumn) int
	// Returns the grid row object at the specified index.
	RowAtIndex(index int) INSGridRow
	// Returns the grid column object at the specified index.
	ColumnAtIndex(index int) INSGridColumn
	// Returns the index of the specified grid row.
	IndexOfRow(row INSGridRow) int

	// Topic: Adding, Removing, and Moving Rows

	// Adds an array of views to a new row.
	AddRowWithViews(views []NSView) INSGridRow
	// Inserts the array of view objects into the grid view at the index.
	InsertRowAtIndexWithViews(index int, views []NSView) INSGridRow
	// Removes the row from the grid view at the index.
	RemoveRowAtIndex(index int)
	// Moves the specified row to the new row location.
	MoveRowAtIndexToIndex(fromIndex int, toIndex int)

	// Topic: Adding, Removing, and Moving Columns

	// Adds a new column containing the array of views.
	AddColumnWithViews(views []NSView) INSGridColumn
	// Inserts the array of view objects at the specified index.
	InsertColumnAtIndexWithViews(index int, views []NSView) INSGridColumn
	// Removes the column from the grid view at the specified index.
	RemoveColumnAtIndex(index int)
	// Moves the specified column to a new column location.
	MoveColumnAtIndexToIndex(fromIndex int, toIndex int)

	// Topic: Managing Grid Spacing and Alignment

	// The column spacing for the grid view.
	ColumnSpacing() float64
	SetColumnSpacing(value float64)
	// The row spacing for the grid view.
	RowSpacing() float64
	SetRowSpacing(value float64)
	// The row alignment for the grid view.
	RowAlignment() NSGridRowAlignment
	SetRowAlignment(value NSGridRowAlignment)
	// The placement of the cell within the grid column.
	XPlacement() NSGridCellPlacement
	SetXPlacement(value NSGridCellPlacement)
	// The placement of the cell within the grid row.
	YPlacement() NSGridCellPlacement
	SetYPlacement(value NSGridCellPlacement)

	// Topic: Creating and Merging Cells

	// Returns the grid cell object at the specified column and row index.
	CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) INSGridCell
	// Returns the grid cell object that contains the given view or one of its ancestors.
	CellForView(view INSView) INSGridCell
	// Expands the cell at the top-leading corner of the horizontal and vertical range to cover the entire area.
	MergeCellsInHorizontalRangeVerticalRange(hRange foundation.NSRange, vRange foundation.NSRange)
}

// Init initializes the instance.
func (g NSGridView) Init() NSGridView {
	rv := objc.Send[NSGridView](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGridView) Autorelease() NSGridView {
	rv := objc.Send[NSGridView](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGridView creates a new NSGridView instance.
func NewNSGridView() NSGridView {
	class := getNSGridViewClass()
	rv := objc.Send[NSGridView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a newly allocated grid view object from the coder.
//
// coder: The coder object.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/init(coder:)
func NewGridViewWithCoder(coder foundation.INSCoder) NSGridView {
	instance := getNSGridViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSGridViewFromID(rv)
}

// Creates a newly allocated grid view object with the specified frame
// rectangle.
//
// frameRect: The frame rectangle for the view, measured in points. The origin of the
// frame is relative to the superview in which you plan to add it.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/init(frame:)
func NewGridViewWithFrame(frameRect corefoundation.CGRect) NSGridView {
	instance := getNSGridViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSGridViewFromID(rv)
}

// Creates a newly allocated grid view object with the specified number of
// columns and rows.
//
// columnCount: The number of columns for this grid view.
//
// rowCount: The number of rows for this grid view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/init(numberOfColumns:rows:)
func NewGridViewWithNumberOfColumnsRows(columnCount int, rowCount int) NSGridView {
	rv := objc.Send[objc.ID](objc.ID(getNSGridViewClass().class), objc.Sel("gridViewWithNumberOfColumns:rows:"), columnCount, rowCount)
	return NSGridViewFromID(rv)
}

// Creates a newly allocated grid view object with the specified array of
// arrays of views.
//
// rows: An array of arrays of grid view row objects.
//
// # Discussion
//
// This method creates an autoreleased grid view large enough to hold the
// passed array of rows. Each element in the array is itself an array of views
// for that row.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/init(views:)
func NewGridViewWithViews(rows []foundation.NSArray) NSGridView {
	rv := objc.Send[objc.ID](objc.ID(getNSGridViewClass().class), objc.Sel("gridViewWithViews:"), objectivec.IObjectSliceToNSArray(rows))
	return NSGridViewFromID(rv)
}

// Returns the index of the specified grid column.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/index(of:)-32sdd
func (g NSGridView) IndexOfColumn(column INSGridColumn) int {
	rv := objc.Send[int](g.ID, objc.Sel("indexOfColumn:"), column)
	return rv
}

// Returns the grid row object at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/row(at:)
func (g NSGridView) RowAtIndex(index int) INSGridRow {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("rowAtIndex:"), index)
	return NSGridRowFromID(rv)
}

// Returns the grid column object at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/column(at:)
func (g NSGridView) ColumnAtIndex(index int) INSGridColumn {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("columnAtIndex:"), index)
	return NSGridColumnFromID(rv)
}

// Returns the index of the specified grid row.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/index(of:)-6zs2o
func (g NSGridView) IndexOfRow(row INSGridRow) int {
	rv := objc.Send[int](g.ID, objc.Sel("indexOfRow:"), row)
	return rv
}

// Adds an array of views to a new row.
//
// # Discussion
//
// You can insert and remove rows and columns dynamically in a grid view. The
// grid is enlarged as needed to hold the specified views.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/addRow(with:)
func (g NSGridView) AddRowWithViews(views []NSView) INSGridRow {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("addRowWithViews:"), objectivec.IObjectSliceToNSArray(views))
	return NSGridRowFromID(rv)
}

// Inserts the array of view objects into the grid view at the index.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/insertRow(at:with:)
func (g NSGridView) InsertRowAtIndexWithViews(index int, views []NSView) INSGridRow {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("insertRowAtIndex:withViews:"), index, objectivec.IObjectSliceToNSArray(views))
	return NSGridRowFromID(rv)
}

// Removes the row from the grid view at the index.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/removeRow(at:)
func (g NSGridView) RemoveRowAtIndex(index int) {
	objc.Send[objc.ID](g.ID, objc.Sel("removeRowAtIndex:"), index)
}

// Moves the specified row to the new row location.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/moveRow(at:to:)
func (g NSGridView) MoveRowAtIndexToIndex(fromIndex int, toIndex int) {
	objc.Send[objc.ID](g.ID, objc.Sel("moveRowAtIndex:toIndex:"), fromIndex, toIndex)
}

// Adds a new column containing the array of views.
//
// # Return Value
//
// The newly created grid column.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/addColumn(with:)
func (g NSGridView) AddColumnWithViews(views []NSView) INSGridColumn {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("addColumnWithViews:"), objectivec.IObjectSliceToNSArray(views))
	return NSGridColumnFromID(rv)
}

// Inserts the array of view objects at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/insertColumn(at:with:)
func (g NSGridView) InsertColumnAtIndexWithViews(index int, views []NSView) INSGridColumn {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("insertColumnAtIndex:withViews:"), index, objectivec.IObjectSliceToNSArray(views))
	return NSGridColumnFromID(rv)
}

// Removes the column from the grid view at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/removeColumn(at:)
func (g NSGridView) RemoveColumnAtIndex(index int) {
	objc.Send[objc.ID](g.ID, objc.Sel("removeColumnAtIndex:"), index)
}

// Moves the specified column to a new column location.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/moveColumn(at:to:)
func (g NSGridView) MoveColumnAtIndexToIndex(fromIndex int, toIndex int) {
	objc.Send[objc.ID](g.ID, objc.Sel("moveColumnAtIndex:toIndex:"), fromIndex, toIndex)
}

// Returns the grid cell object at the specified column and row index.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/cell(atColumnIndex:rowIndex:)
func (g NSGridView) CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) INSGridCell {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cellAtColumnIndex:rowIndex:"), columnIndex, rowIndex)
	return NSGridCellFromID(rv)
}

// Returns the grid cell object that contains the given view or one of its
// ancestors.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/cell(for:)
func (g NSGridView) CellForView(view INSView) INSGridCell {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cellForView:"), view)
	return NSGridCellFromID(rv)
}

// Expands the cell at the top-leading corner of the horizontal and vertical
// range to cover the entire area.
//
// # Discussion
//
// This function invalidates other cells in the range, and they no longer
// maintain their layout, constraints, or content views. Cell merging has no
// effect on the base cell coordinate system of the grid view, and cell
// references within a merged region refer to the single merged cell.
//
// Use this method to configure the grid geometry before installing views. If
// the cells being merged contain content views, only the top-leading views
// are kept.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/mergeCells(inHorizontalRange:verticalRange:)
func (g NSGridView) MergeCellsInHorizontalRangeVerticalRange(hRange foundation.NSRange, vRange foundation.NSRange) {
	objc.Send[objc.ID](g.ID, objc.Sel("mergeCellsInHorizontalRange:verticalRange:"), hRange, vRange)
}

// The number of rows in the grid view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/numberOfRows
func (g NSGridView) NumberOfRows() int {
	rv := objc.Send[int](g.ID, objc.Sel("numberOfRows"))
	return rv
}

// The number of columns in the grid view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/numberOfColumns
func (g NSGridView) NumberOfColumns() int {
	rv := objc.Send[int](g.ID, objc.Sel("numberOfColumns"))
	return rv
}

// The column spacing for the grid view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/columnSpacing
func (g NSGridView) ColumnSpacing() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("columnSpacing"))
	return rv
}
func (g NSGridView) SetColumnSpacing(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setColumnSpacing:"), value)
}

// The row spacing for the grid view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/rowSpacing
func (g NSGridView) RowSpacing() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("rowSpacing"))
	return rv
}
func (g NSGridView) SetRowSpacing(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setRowSpacing:"), value)
}

// The row alignment for the grid view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/rowAlignment
func (g NSGridView) RowAlignment() NSGridRowAlignment {
	rv := objc.Send[NSGridRowAlignment](g.ID, objc.Sel("rowAlignment"))
	return NSGridRowAlignment(rv)
}
func (g NSGridView) SetRowAlignment(value NSGridRowAlignment) {
	objc.Send[struct{}](g.ID, objc.Sel("setRowAlignment:"), value)
}

// The placement of the cell within the grid column.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/xPlacement
func (g NSGridView) XPlacement() NSGridCellPlacement {
	rv := objc.Send[NSGridCellPlacement](g.ID, objc.Sel("xPlacement"))
	return NSGridCellPlacement(rv)
}
func (g NSGridView) SetXPlacement(value NSGridCellPlacement) {
	objc.Send[struct{}](g.ID, objc.Sel("setXPlacement:"), value)
}

// The placement of the cell within the grid row.
//
// See: https://developer.apple.com/documentation/AppKit/NSGridView/yPlacement
func (g NSGridView) YPlacement() NSGridCellPlacement {
	rv := objc.Send[NSGridCellPlacement](g.ID, objc.Sel("yPlacement"))
	return NSGridCellPlacement(rv)
}
func (g NSGridView) SetYPlacement(value NSGridCellPlacement) {
	objc.Send[struct{}](g.ID, objc.Sel("setYPlacement:"), value)
}

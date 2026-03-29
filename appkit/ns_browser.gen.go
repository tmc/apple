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

// The class instance for the [NSBrowser] class.
var (
	_NSBrowserClass     NSBrowserClass
	_NSBrowserClassOnce sync.Once
)

func getNSBrowserClass() NSBrowserClass {
	_NSBrowserClassOnce.Do(func() {
		_NSBrowserClass = NSBrowserClass{class: objc.GetClass("NSBrowser")}
	})
	return _NSBrowserClass
}

// GetNSBrowserClass returns the class object for NSBrowser.
func GetNSBrowserClass() NSBrowserClass {
	return getNSBrowserClass()
}

type NSBrowserClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBrowserClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBrowserClass) Alloc() NSBrowser {
	rv := objc.Send[NSBrowser](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An interface that displays a hierarchically organized list of data items
// that can be navigated and selected.
//
// # Overview
// 
// A browser displays information using a set of columns, which are indexed
// from left to right. Each successive column displays the next level down in
// the data hierarchy. This class uses the [NSBrowserCell] class to implement
// its user interface.
// 
// Browsers have the following components:
// 
// - Columns
// - Scroll views
// - Matrices
// - Browser cells
// 
// To the user, browsers display data in columns and rows within each column.
// These components are arranged in the following component hierarchy:
// 
// # Superclass overrides
// 
// - [isOpaque] returns [true] when the browser doesn’t have a title and its
// background color’s alpha component is `1.0`; otherwise, it returns
// [false].
// 
// # Protocol implementations
// 
// - The [NSBrowser] implementation of
// [namesOfPromisedFilesDropped(atDestination:)] provides the names of the
// files that the browser promises to create at a specified location, the
// result of sending `` to the delegate.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [isOpaque]: https://developer.apple.com/documentation/AppKit/NSView/isOpaque
// [namesOfPromisedFilesDropped(atDestination:)]: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/namesOfPromisedFilesDropped(atDestination:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Configuring Browsers
//
//   - [NSBrowser.ReusesColumns]: A Boolean that indicates whether the browser reuses matrix objects after their columns are unloaded.
//   - [NSBrowser.SetReusesColumns]
//   - [NSBrowser.MaxVisibleColumns]: The maximum number of visible columns.
//   - [NSBrowser.SetMaxVisibleColumns]
//   - [NSBrowser.AutohidesScroller]: A Boolean that indicates whether the browser automatically hides its scroller.
//   - [NSBrowser.SetAutohidesScroller]
//   - [NSBrowser.BackgroundColor]: The browser’s background color.
//   - [NSBrowser.SetBackgroundColor]
//   - [NSBrowser.MinColumnWidth]: The minimum column width, in pixels.
//   - [NSBrowser.SetMinColumnWidth]
//   - [NSBrowser.SeparatesColumns]: A Boolean that indicates whether columns are separated by bezeled borders.
//   - [NSBrowser.SetSeparatesColumns]
//   - [NSBrowser.TakesTitleFromPreviousColumn]: A Boolean that indicates whether a column takes its title from the selected cell in the previous column.
//   - [NSBrowser.SetTakesTitleFromPreviousColumn]
//   - [NSBrowser.Tile]: Adjusts the various subviews of the browser—scrollers, columns, titles, and so on—without redrawing.
//   - [NSBrowser.Delegate]: The browser’s delegate.
//   - [NSBrowser.SetDelegate]
//
// # Managing Component Types
//
//   - [NSBrowser.CellPrototype]: The prototype [NSCell] for displaying items in the matrices in the columns of the browser.
//   - [NSBrowser.SetCellPrototype]
//
// # Managing Selection Behavior
//
//   - [NSBrowser.AllowsBranchSelection]: A Boolean that indicates whether the user can select branch items.
//   - [NSBrowser.SetAllowsBranchSelection]
//   - [NSBrowser.AllowsEmptySelection]: A Boolean that indicates whether there can be nothing selected.
//   - [NSBrowser.SetAllowsEmptySelection]
//   - [NSBrowser.AllowsMultipleSelection]: A Boolean that indicates whether the user can select multiple items.
//   - [NSBrowser.SetAllowsMultipleSelection]
//   - [NSBrowser.SelectedRowIndexesInColumn]: Provides the indexes of the selected rows in a given column of the browser.
//   - [NSBrowser.SelectRowIndexesInColumn]: Specifies the selected rows in a given column of the browser.
//   - [NSBrowser.AllowsTypeSelect]: A Boolean that indicates whether the browser allows keystroke-based selection (type select).
//   - [NSBrowser.SetAllowsTypeSelect]
//
// # Managing Selection
//
//   - [NSBrowser.SelectedCellInColumn]: Returns the last (lowest) cell selected in the given column.
//   - [NSBrowser.SelectedCells]: All cells selected in the rightmost column.
//   - [NSBrowser.SelectedRowInColumn]: Returns the row index of the selected cell in the specified column.
//   - [NSBrowser.SelectRowInColumn]: Selects the cell at the specified row and column index.
//   - [NSBrowser.SelectionIndexPath]: The index path of the item selected in the browser.
//   - [NSBrowser.SetSelectionIndexPath]
//   - [NSBrowser.SelectionIndexPaths]: An array containing the index paths of all items selected in the browser.
//   - [NSBrowser.SetSelectionIndexPaths]
//
// # Accessing Components
//
//   - [NSBrowser.LoadedCellAtRowColumn]: Loads, if necessary, and returns the cell at the specified row and column location.
//   - [NSBrowser.EditItemAtIndexPathWithEventSelect]: Begins editing the item at the specified path.
//   - [NSBrowser.ItemAtIndexPath]: Returns the item at the specified index path.
//   - [NSBrowser.ItemAtRowInColumn]: Returns the item located at the specified row and column.
//   - [NSBrowser.IndexPathForColumn]: Returns the index path of the item whose children are displayed in the given column.
//   - [NSBrowser.IsLeafItem]: Returns whether the specified item is a leaf item.
//   - [NSBrowser.ParentForItemsInColumn]: Returns the item that contains the children located in the specified column.
//
// # Managing the Path
//
//   - [NSBrowser.Path]: Returns a string representing the browser’s current path.
//   - [NSBrowser.SetPath]: Sets the path to be displayed by the browser.
//   - [NSBrowser.PathToColumn]: Returns a string representing the path from the first column up to, but not including, the column at the given index.
//   - [NSBrowser.PathSeparator]: The path separator.
//   - [NSBrowser.SetPathSeparator]
//
// # Managing Columns
//
//   - [NSBrowser.AddColumn]: Adds a column to the right of the last column.
//   - [NSBrowser.SelectedColumn]: The index of the last column with a selected item.
//   - [NSBrowser.LastColumn]: The index of the last column loaded.
//   - [NSBrowser.SetLastColumn]
//   - [NSBrowser.FirstVisibleColumn]: The index of the first visible column.
//   - [NSBrowser.NumberOfVisibleColumns]: The number of visible columns.
//   - [NSBrowser.LastVisibleColumn]: The index of the last visible column.
//   - [NSBrowser.ValidateVisibleColumns]: Validates the browser’s visible columns.
//   - [NSBrowser.Loaded]: A Boolean that indicates whether column 0 is loaded.
//   - [NSBrowser.LoadColumnZero]: Loads column 0; unloads previously loaded columns.
//   - [NSBrowser.ReloadColumn]: Reloads the given column.
//
// # Accessing Column Titles
//
//   - [NSBrowser.TitleOfColumn]: Returns the title displayed for the given column.
//   - [NSBrowser.SetTitleOfColumn]: Sets the title of the given column.
//   - [NSBrowser.Titled]: A Boolean that indicates whether columns display titles.
//   - [NSBrowser.SetTitled]
//   - [NSBrowser.DrawTitleOfColumnInRect]: Draws the title for the specified column within the given rectangle.
//   - [NSBrowser.TitleHeight]: The height of the column titles for the browser.
//   - [NSBrowser.TitleFrameOfColumn]: Returns the bounds of the title frame for the specified column.
//
// # Updating Browsers
//
//   - [NSBrowser.NoteHeightOfRowsWithIndexesChangedInColumn]: Immediately retiles the browser’s columns using row heights specified by the browser’s delegate.
//   - [NSBrowser.ReloadDataForRowIndexesInColumn]: Updates the rows in the column with the specified column index with indexes in the specified set.
//
// # Scrolling
//
//   - [NSBrowser.HasHorizontalScroller]: A Boolean that indicates whether the browser has a horizontal scroller.
//   - [NSBrowser.SetHasHorizontalScroller]
//   - [NSBrowser.ScrollColumnToVisible]: Scrolls to make the specified column visible.
//   - [NSBrowser.ScrollColumnsLeftBy]: Scrolls columns left by the specified number of columns.
//   - [NSBrowser.ScrollColumnsRightBy]: Scrolls columns right by the specified number of columns.
//   - [NSBrowser.ScrollRowToVisibleInColumn]: Scrolls the specified row to be visible within the specified column.
//
// # Dragging
//
//   - [NSBrowser.SetDraggingSourceOperationMaskForLocal]: Specifies the drag-operation mask for dragging operations with local or external destinations.
//   - [NSBrowser.CanDragRowsWithIndexesInColumnWithEvent]: Indicates whether the browser can attempt to initiate a drag of the given rows for the given event.
//   - [NSBrowser.DraggingImageForRowsWithIndexesInColumnWithEventOffset]: Provides an image to represent dragged rows during a drag operation on the browser.
//
// # Getting Column Frames
//
//   - [NSBrowser.FrameOfColumn]: Returns the rectangle containing the given column.
//   - [NSBrowser.FrameOfInsideOfColumn]: Returns the rectangle containing the specified column, not including borders.
//
// # Getting Row Frames
//
//   - [NSBrowser.FrameOfRowInColumn]: Returns the frame of the cell at the specified location, including the expandable arrow.
//   - [NSBrowser.GetRowColumnForPoint]: Gets the row and column coordinates for the specified point, if a cell exists at that point.
//
// # Managing Actions
//
//   - [NSBrowser.DoubleAction]: The browser’s double-click action method.
//   - [NSBrowser.SetDoubleAction]
//   - [NSBrowser.SendsActionOnArrowKeys]: A Boolean that indicates whether pressing an arrow key causes an action message to be sent.
//   - [NSBrowser.SetSendsActionOnArrowKeys]
//   - [NSBrowser.SendAction]: Sends the action message to the target.
//
// # Handling Mouse-Click Events
//
//   - [NSBrowser.DoClick]: Responds to (single) mouse clicks in a column of the browser.
//   - [NSBrowser.DoDoubleClick]: Responds to double clicks in a column of the browser.
//   - [NSBrowser.ClickedColumn]: The column number of the cell that the user clicked to display a context menu.
//   - [NSBrowser.ClickedRow]: The row number of the cell that the user clicked to display a context menu.
//
// # Sizing
//
//   - [NSBrowser.ColumnsAutosaveName]: The name used to automatically save the browser’s column configuration.
//   - [NSBrowser.SetColumnsAutosaveName]
//   - [NSBrowser.ColumnContentWidthForColumnWidth]: Returns the content width for a given column width.
//   - [NSBrowser.ColumnWidthForColumnContentWidth]: Returns the column width for the width of the given column’s content.
//   - [NSBrowser.ColumnResizingType]: A constant indicating the browser’s column resizing type.
//   - [NSBrowser.SetColumnResizingType]
//   - [NSBrowser.PrefersAllColumnUserResizing]: A Boolean that indicates whether the browser is set to resize all columns simultaneously rather than resizing a single column at a time.
//   - [NSBrowser.SetPrefersAllColumnUserResizing]
//   - [NSBrowser.WidthOfColumn]: Returns the width of the specified column.
//   - [NSBrowser.SetWidthOfColumn]: Sets the width of the specified column.
//   - [NSBrowser.DefaultColumnWidth]: Returns the default column width of the browser’s columns.
//   - [NSBrowser.SetDefaultColumnWidth]: Sets the default column width for new browser columns that do not otherwise have an initial width from defaults or the browser’s delegate.
//   - [NSBrowser.RowHeight]: The height of the browser’s rows.
//   - [NSBrowser.SetRowHeight]
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser
type NSBrowser struct {
	NSControl
}

// NSBrowserFromID constructs a [NSBrowser] from an objc.ID.
//
// An interface that displays a hierarchically organized list of data items
// that can be navigated and selected.
func NSBrowserFromID(id objc.ID) NSBrowser {
	return NSBrowser{NSControl: NSControlFromID(id)}
}
// NOTE: NSBrowser adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBrowser] class.
//
// # Configuring Browsers
//
//   - [INSBrowser.ReusesColumns]: A Boolean that indicates whether the browser reuses matrix objects after their columns are unloaded.
//   - [INSBrowser.SetReusesColumns]
//   - [INSBrowser.MaxVisibleColumns]: The maximum number of visible columns.
//   - [INSBrowser.SetMaxVisibleColumns]
//   - [INSBrowser.AutohidesScroller]: A Boolean that indicates whether the browser automatically hides its scroller.
//   - [INSBrowser.SetAutohidesScroller]
//   - [INSBrowser.BackgroundColor]: The browser’s background color.
//   - [INSBrowser.SetBackgroundColor]
//   - [INSBrowser.MinColumnWidth]: The minimum column width, in pixels.
//   - [INSBrowser.SetMinColumnWidth]
//   - [INSBrowser.SeparatesColumns]: A Boolean that indicates whether columns are separated by bezeled borders.
//   - [INSBrowser.SetSeparatesColumns]
//   - [INSBrowser.TakesTitleFromPreviousColumn]: A Boolean that indicates whether a column takes its title from the selected cell in the previous column.
//   - [INSBrowser.SetTakesTitleFromPreviousColumn]
//   - [INSBrowser.Tile]: Adjusts the various subviews of the browser—scrollers, columns, titles, and so on—without redrawing.
//   - [INSBrowser.Delegate]: The browser’s delegate.
//   - [INSBrowser.SetDelegate]
//
// # Managing Component Types
//
//   - [INSBrowser.CellPrototype]: The prototype [NSCell] for displaying items in the matrices in the columns of the browser.
//   - [INSBrowser.SetCellPrototype]
//
// # Managing Selection Behavior
//
//   - [INSBrowser.AllowsBranchSelection]: A Boolean that indicates whether the user can select branch items.
//   - [INSBrowser.SetAllowsBranchSelection]
//   - [INSBrowser.AllowsEmptySelection]: A Boolean that indicates whether there can be nothing selected.
//   - [INSBrowser.SetAllowsEmptySelection]
//   - [INSBrowser.AllowsMultipleSelection]: A Boolean that indicates whether the user can select multiple items.
//   - [INSBrowser.SetAllowsMultipleSelection]
//   - [INSBrowser.SelectedRowIndexesInColumn]: Provides the indexes of the selected rows in a given column of the browser.
//   - [INSBrowser.SelectRowIndexesInColumn]: Specifies the selected rows in a given column of the browser.
//   - [INSBrowser.AllowsTypeSelect]: A Boolean that indicates whether the browser allows keystroke-based selection (type select).
//   - [INSBrowser.SetAllowsTypeSelect]
//
// # Managing Selection
//
//   - [INSBrowser.SelectedCellInColumn]: Returns the last (lowest) cell selected in the given column.
//   - [INSBrowser.SelectedCells]: All cells selected in the rightmost column.
//   - [INSBrowser.SelectedRowInColumn]: Returns the row index of the selected cell in the specified column.
//   - [INSBrowser.SelectRowInColumn]: Selects the cell at the specified row and column index.
//   - [INSBrowser.SelectionIndexPath]: The index path of the item selected in the browser.
//   - [INSBrowser.SetSelectionIndexPath]
//   - [INSBrowser.SelectionIndexPaths]: An array containing the index paths of all items selected in the browser.
//   - [INSBrowser.SetSelectionIndexPaths]
//
// # Accessing Components
//
//   - [INSBrowser.LoadedCellAtRowColumn]: Loads, if necessary, and returns the cell at the specified row and column location.
//   - [INSBrowser.EditItemAtIndexPathWithEventSelect]: Begins editing the item at the specified path.
//   - [INSBrowser.ItemAtIndexPath]: Returns the item at the specified index path.
//   - [INSBrowser.ItemAtRowInColumn]: Returns the item located at the specified row and column.
//   - [INSBrowser.IndexPathForColumn]: Returns the index path of the item whose children are displayed in the given column.
//   - [INSBrowser.IsLeafItem]: Returns whether the specified item is a leaf item.
//   - [INSBrowser.ParentForItemsInColumn]: Returns the item that contains the children located in the specified column.
//
// # Managing the Path
//
//   - [INSBrowser.Path]: Returns a string representing the browser’s current path.
//   - [INSBrowser.SetPath]: Sets the path to be displayed by the browser.
//   - [INSBrowser.PathToColumn]: Returns a string representing the path from the first column up to, but not including, the column at the given index.
//   - [INSBrowser.PathSeparator]: The path separator.
//   - [INSBrowser.SetPathSeparator]
//
// # Managing Columns
//
//   - [INSBrowser.AddColumn]: Adds a column to the right of the last column.
//   - [INSBrowser.SelectedColumn]: The index of the last column with a selected item.
//   - [INSBrowser.LastColumn]: The index of the last column loaded.
//   - [INSBrowser.SetLastColumn]
//   - [INSBrowser.FirstVisibleColumn]: The index of the first visible column.
//   - [INSBrowser.NumberOfVisibleColumns]: The number of visible columns.
//   - [INSBrowser.LastVisibleColumn]: The index of the last visible column.
//   - [INSBrowser.ValidateVisibleColumns]: Validates the browser’s visible columns.
//   - [INSBrowser.Loaded]: A Boolean that indicates whether column 0 is loaded.
//   - [INSBrowser.LoadColumnZero]: Loads column 0; unloads previously loaded columns.
//   - [INSBrowser.ReloadColumn]: Reloads the given column.
//
// # Accessing Column Titles
//
//   - [INSBrowser.TitleOfColumn]: Returns the title displayed for the given column.
//   - [INSBrowser.SetTitleOfColumn]: Sets the title of the given column.
//   - [INSBrowser.Titled]: A Boolean that indicates whether columns display titles.
//   - [INSBrowser.SetTitled]
//   - [INSBrowser.DrawTitleOfColumnInRect]: Draws the title for the specified column within the given rectangle.
//   - [INSBrowser.TitleHeight]: The height of the column titles for the browser.
//   - [INSBrowser.TitleFrameOfColumn]: Returns the bounds of the title frame for the specified column.
//
// # Updating Browsers
//
//   - [INSBrowser.NoteHeightOfRowsWithIndexesChangedInColumn]: Immediately retiles the browser’s columns using row heights specified by the browser’s delegate.
//   - [INSBrowser.ReloadDataForRowIndexesInColumn]: Updates the rows in the column with the specified column index with indexes in the specified set.
//
// # Scrolling
//
//   - [INSBrowser.HasHorizontalScroller]: A Boolean that indicates whether the browser has a horizontal scroller.
//   - [INSBrowser.SetHasHorizontalScroller]
//   - [INSBrowser.ScrollColumnToVisible]: Scrolls to make the specified column visible.
//   - [INSBrowser.ScrollColumnsLeftBy]: Scrolls columns left by the specified number of columns.
//   - [INSBrowser.ScrollColumnsRightBy]: Scrolls columns right by the specified number of columns.
//   - [INSBrowser.ScrollRowToVisibleInColumn]: Scrolls the specified row to be visible within the specified column.
//
// # Dragging
//
//   - [INSBrowser.SetDraggingSourceOperationMaskForLocal]: Specifies the drag-operation mask for dragging operations with local or external destinations.
//   - [INSBrowser.CanDragRowsWithIndexesInColumnWithEvent]: Indicates whether the browser can attempt to initiate a drag of the given rows for the given event.
//   - [INSBrowser.DraggingImageForRowsWithIndexesInColumnWithEventOffset]: Provides an image to represent dragged rows during a drag operation on the browser.
//
// # Getting Column Frames
//
//   - [INSBrowser.FrameOfColumn]: Returns the rectangle containing the given column.
//   - [INSBrowser.FrameOfInsideOfColumn]: Returns the rectangle containing the specified column, not including borders.
//
// # Getting Row Frames
//
//   - [INSBrowser.FrameOfRowInColumn]: Returns the frame of the cell at the specified location, including the expandable arrow.
//   - [INSBrowser.GetRowColumnForPoint]: Gets the row and column coordinates for the specified point, if a cell exists at that point.
//
// # Managing Actions
//
//   - [INSBrowser.DoubleAction]: The browser’s double-click action method.
//   - [INSBrowser.SetDoubleAction]
//   - [INSBrowser.SendsActionOnArrowKeys]: A Boolean that indicates whether pressing an arrow key causes an action message to be sent.
//   - [INSBrowser.SetSendsActionOnArrowKeys]
//   - [INSBrowser.SendAction]: Sends the action message to the target.
//
// # Handling Mouse-Click Events
//
//   - [INSBrowser.DoClick]: Responds to (single) mouse clicks in a column of the browser.
//   - [INSBrowser.DoDoubleClick]: Responds to double clicks in a column of the browser.
//   - [INSBrowser.ClickedColumn]: The column number of the cell that the user clicked to display a context menu.
//   - [INSBrowser.ClickedRow]: The row number of the cell that the user clicked to display a context menu.
//
// # Sizing
//
//   - [INSBrowser.ColumnsAutosaveName]: The name used to automatically save the browser’s column configuration.
//   - [INSBrowser.SetColumnsAutosaveName]
//   - [INSBrowser.ColumnContentWidthForColumnWidth]: Returns the content width for a given column width.
//   - [INSBrowser.ColumnWidthForColumnContentWidth]: Returns the column width for the width of the given column’s content.
//   - [INSBrowser.ColumnResizingType]: A constant indicating the browser’s column resizing type.
//   - [INSBrowser.SetColumnResizingType]
//   - [INSBrowser.PrefersAllColumnUserResizing]: A Boolean that indicates whether the browser is set to resize all columns simultaneously rather than resizing a single column at a time.
//   - [INSBrowser.SetPrefersAllColumnUserResizing]
//   - [INSBrowser.WidthOfColumn]: Returns the width of the specified column.
//   - [INSBrowser.SetWidthOfColumn]: Sets the width of the specified column.
//   - [INSBrowser.DefaultColumnWidth]: Returns the default column width of the browser’s columns.
//   - [INSBrowser.SetDefaultColumnWidth]: Sets the default column width for new browser columns that do not otherwise have an initial width from defaults or the browser’s delegate.
//   - [INSBrowser.RowHeight]: The height of the browser’s rows.
//   - [INSBrowser.SetRowHeight]
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser
type INSBrowser interface {
	INSControl

	// Topic: Configuring Browsers

	// A Boolean that indicates whether the browser reuses matrix objects after their columns are unloaded.
	ReusesColumns() bool
	SetReusesColumns(value bool)
	// The maximum number of visible columns.
	MaxVisibleColumns() int
	SetMaxVisibleColumns(value int)
	// A Boolean that indicates whether the browser automatically hides its scroller.
	AutohidesScroller() bool
	SetAutohidesScroller(value bool)
	// The browser’s background color.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// The minimum column width, in pixels.
	MinColumnWidth() float64
	SetMinColumnWidth(value float64)
	// A Boolean that indicates whether columns are separated by bezeled borders.
	SeparatesColumns() bool
	SetSeparatesColumns(value bool)
	// A Boolean that indicates whether a column takes its title from the selected cell in the previous column.
	TakesTitleFromPreviousColumn() bool
	SetTakesTitleFromPreviousColumn(value bool)
	// Adjusts the various subviews of the browser—scrollers, columns, titles, and so on—without redrawing.
	Tile()
	// The browser’s delegate.
	Delegate() NSBrowserDelegate
	SetDelegate(value NSBrowserDelegate)

	// Topic: Managing Component Types

	// The prototype [NSCell] for displaying items in the matrices in the columns of the browser.
	CellPrototype() objectivec.IObject
	SetCellPrototype(value objectivec.IObject)

	// Topic: Managing Selection Behavior

	// A Boolean that indicates whether the user can select branch items.
	AllowsBranchSelection() bool
	SetAllowsBranchSelection(value bool)
	// A Boolean that indicates whether there can be nothing selected.
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	// A Boolean that indicates whether the user can select multiple items.
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	// Provides the indexes of the selected rows in a given column of the browser.
	SelectedRowIndexesInColumn(column int) foundation.NSIndexSet
	// Specifies the selected rows in a given column of the browser.
	SelectRowIndexesInColumn(indexes foundation.NSIndexSet, column int)
	// A Boolean that indicates whether the browser allows keystroke-based selection (type select).
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)

	// Topic: Managing Selection

	// Returns the last (lowest) cell selected in the given column.
	SelectedCellInColumn(column int) objectivec.IObject
	// All cells selected in the rightmost column.
	SelectedCells() []NSCell
	// Returns the row index of the selected cell in the specified column.
	SelectedRowInColumn(column int) int
	// Selects the cell at the specified row and column index.
	SelectRowInColumn(row int, column int)
	// The index path of the item selected in the browser.
	SelectionIndexPath() objc.ID
	SetSelectionIndexPath(value objc.ID)
	// An array containing the index paths of all items selected in the browser.
	SelectionIndexPaths() []objc.ID
	SetSelectionIndexPaths(value []objc.ID)

	// Topic: Accessing Components

	// Loads, if necessary, and returns the cell at the specified row and column location.
	LoadedCellAtRowColumn(row int, col int) objectivec.IObject
	// Begins editing the item at the specified path.
	EditItemAtIndexPathWithEventSelect(indexPath foundation.INSIndexPath, event INSEvent, select_ bool)
	// Returns the item at the specified index path.
	ItemAtIndexPath(indexPath foundation.INSIndexPath) objectivec.IObject
	// Returns the item located at the specified row and column.
	ItemAtRowInColumn(row int, column int) objectivec.IObject
	// Returns the index path of the item whose children are displayed in the given column.
	IndexPathForColumn(column int) objc.ID
	// Returns whether the specified item is a leaf item.
	IsLeafItem(item objectivec.IObject) bool
	// Returns the item that contains the children located in the specified column.
	ParentForItemsInColumn(column int) objectivec.IObject

	// Topic: Managing the Path

	// Returns a string representing the browser’s current path.
	Path() string
	// Sets the path to be displayed by the browser.
	SetPath(path string) bool
	// Returns a string representing the path from the first column up to, but not including, the column at the given index.
	PathToColumn(column int) string
	// The path separator.
	PathSeparator() string
	SetPathSeparator(value string)

	// Topic: Managing Columns

	// Adds a column to the right of the last column.
	AddColumn()
	// The index of the last column with a selected item.
	SelectedColumn() int
	// The index of the last column loaded.
	LastColumn() int
	SetLastColumn(value int)
	// The index of the first visible column.
	FirstVisibleColumn() int
	// The number of visible columns.
	NumberOfVisibleColumns() int
	// The index of the last visible column.
	LastVisibleColumn() int
	// Validates the browser’s visible columns.
	ValidateVisibleColumns()
	// A Boolean that indicates whether column 0 is loaded.
	Loaded() bool
	// Loads column 0; unloads previously loaded columns.
	LoadColumnZero()
	// Reloads the given column.
	ReloadColumn(column int)

	// Topic: Accessing Column Titles

	// Returns the title displayed for the given column.
	TitleOfColumn(column int) string
	// Sets the title of the given column.
	SetTitleOfColumn(string_ string, column int)
	// A Boolean that indicates whether columns display titles.
	Titled() bool
	SetTitled(value bool)
	// Draws the title for the specified column within the given rectangle.
	DrawTitleOfColumnInRect(column int, rect corefoundation.CGRect)
	// The height of the column titles for the browser.
	TitleHeight() float64
	// Returns the bounds of the title frame for the specified column.
	TitleFrameOfColumn(column int) corefoundation.CGRect

	// Topic: Updating Browsers

	// Immediately retiles the browser’s columns using row heights specified by the browser’s delegate.
	NoteHeightOfRowsWithIndexesChangedInColumn(indexSet foundation.NSIndexSet, columnIndex int)
	// Updates the rows in the column with the specified column index with indexes in the specified set.
	ReloadDataForRowIndexesInColumn(rowIndexes foundation.NSIndexSet, column int)

	// Topic: Scrolling

	// A Boolean that indicates whether the browser has a horizontal scroller.
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	// Scrolls to make the specified column visible.
	ScrollColumnToVisible(column int)
	// Scrolls columns left by the specified number of columns.
	ScrollColumnsLeftBy(shiftAmount int)
	// Scrolls columns right by the specified number of columns.
	ScrollColumnsRightBy(shiftAmount int)
	// Scrolls the specified row to be visible within the specified column.
	ScrollRowToVisibleInColumn(row int, column int)

	// Topic: Dragging

	// Specifies the drag-operation mask for dragging operations with local or external destinations.
	SetDraggingSourceOperationMaskForLocal(mask NSDragOperation, isLocal bool)
	// Indicates whether the browser can attempt to initiate a drag of the given rows for the given event.
	CanDragRowsWithIndexesInColumnWithEvent(rowIndexes foundation.NSIndexSet, column int, event INSEvent) bool
	// Provides an image to represent dragged rows during a drag operation on the browser.
	DraggingImageForRowsWithIndexesInColumnWithEventOffset(rowIndexes foundation.NSIndexSet, column int, event INSEvent, dragImageOffset foundation.NSPoint) INSImage

	// Topic: Getting Column Frames

	// Returns the rectangle containing the given column.
	FrameOfColumn(column int) corefoundation.CGRect
	// Returns the rectangle containing the specified column, not including borders.
	FrameOfInsideOfColumn(column int) corefoundation.CGRect

	// Topic: Getting Row Frames

	// Returns the frame of the cell at the specified location, including the expandable arrow.
	FrameOfRowInColumn(row int, column int) corefoundation.CGRect
	// Gets the row and column coordinates for the specified point, if a cell exists at that point.
	GetRowColumnForPoint(point corefoundation.CGPoint) (int, int, bool)

	// Topic: Managing Actions

	// The browser’s double-click action method.
	DoubleAction() objc.SEL
	SetDoubleAction(value objc.SEL)
	// A Boolean that indicates whether pressing an arrow key causes an action message to be sent.
	SendsActionOnArrowKeys() bool
	SetSendsActionOnArrowKeys(value bool)
	// Sends the action message to the target.
	SendAction() bool

	// Topic: Handling Mouse-Click Events

	// Responds to (single) mouse clicks in a column of the browser.
	DoClick(sender objectivec.IObject)
	// Responds to double clicks in a column of the browser.
	DoDoubleClick(sender objectivec.IObject)
	// The column number of the cell that the user clicked to display a context menu.
	ClickedColumn() int
	// The row number of the cell that the user clicked to display a context menu.
	ClickedRow() int

	// Topic: Sizing

	// The name used to automatically save the browser’s column configuration.
	ColumnsAutosaveName() NSBrowserColumnsAutosaveName
	SetColumnsAutosaveName(value NSBrowserColumnsAutosaveName)
	// Returns the content width for a given column width.
	ColumnContentWidthForColumnWidth(columnWidth float64) float64
	// Returns the column width for the width of the given column’s content.
	ColumnWidthForColumnContentWidth(columnContentWidth float64) float64
	// A constant indicating the browser’s column resizing type.
	ColumnResizingType() NSBrowserColumnResizingType
	SetColumnResizingType(value NSBrowserColumnResizingType)
	// A Boolean that indicates whether the browser is set to resize all columns simultaneously rather than resizing a single column at a time.
	PrefersAllColumnUserResizing() bool
	SetPrefersAllColumnUserResizing(value bool)
	// Returns the width of the specified column.
	WidthOfColumn(column int) float64
	// Sets the width of the specified column.
	SetWidthOfColumn(columnWidth float64, columnIndex int)
	// Returns the default column width of the browser’s columns.
	DefaultColumnWidth() float64
	// Sets the default column width for new browser columns that do not otherwise have an initial width from defaults or the browser’s delegate.
	SetDefaultColumnWidth(columnWidth float64)
	// The height of the browser’s rows.
	RowHeight() float64
	SetRowHeight(value float64)
}

// Init initializes the instance.
func (b NSBrowser) Init() NSBrowser {
	rv := objc.Send[NSBrowser](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBrowser) Autorelease() NSBrowser {
	rv := objc.Send[NSBrowser](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBrowser creates a new NSBrowser instance.
func NewNSBrowser() NSBrowser {
	class := getNSBrowserClass()
	rv := objc.Send[NSBrowser](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewBrowserWithCoder(coder foundation.INSCoder) NSBrowser {
	instance := getNSBrowserClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSBrowserFromID(rv)
}

// Initializes a control with the specified frame rectangle.
//
// frameRect: The rectangle of the control, specified in points in the coordinate space
// of the enclosing view.
//
// # Return Value
// 
// An initialized control object, or `nil` if the object couldn’t be
// initialized.
//
// # Discussion
// 
// If a cell has been specified for controls of this type, this method also
// creates an instance of the cell. Because [NSControl] is an abstract class,
// invocations of this method should appear only in the designated
// initializers of subclasses; that is, there should always be a more specific
// designated initializer for the subclass, as this method is the designated
// initializer for [NSControl].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewBrowserWithFrame(frameRect corefoundation.CGRect) NSBrowser {
	instance := getNSBrowserClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSBrowserFromID(rv)
}

// Adjusts the various subviews of the browser—scrollers, columns, titles,
// and so on—without redrawing.
//
// # Discussion
// 
// Your code shouldn’t send this message. It’s invoked any time the
// appearance of the browser changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/tile()
func (b NSBrowser) Tile() {
	objc.Send[objc.ID](b.ID, objc.Sel("tile"))
}
// Provides the indexes of the selected rows in a given column of the browser.
//
// column: The column whose selected rows are provided.
//
// # Return Value
// 
// Rows selected in column `columnIndex`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedRowIndexes(inColumn:)
func (b NSBrowser) SelectedRowIndexesInColumn(column int) foundation.NSIndexSet {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("selectedRowIndexesInColumn:"), column)
	return foundation.NSIndexSetFromID(rv)
}
// Specifies the selected rows in a given column of the browser.
//
// indexes: Rows to be selected in column `columnIndex`.
//
// column: Column in which to select rows `rowIndexes`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/selectRowIndexes(_:inColumn:)
func (b NSBrowser) SelectRowIndexesInColumn(indexes foundation.NSIndexSet, column int) {
	objc.Send[objc.ID](b.ID, objc.Sel("selectRowIndexes:inColumn:"), indexes, column)
}
// Returns the last (lowest) cell selected in the given column.
//
// column: The column whose last selected cell is to be returned.
//
// # Return Value
// 
// The last (or lowest) selected cell. Returns `nil` if there is no selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedCell(inColumn:)
func (b NSBrowser) SelectedCellInColumn(column int) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("selectedCellInColumn:"), column)
	return objectivec.Object{ID: rv}
}
// Returns the row index of the selected cell in the specified column.
//
// column: The column index specifying the column for which to return the selected
// row.
//
// # Return Value
// 
// The row index of the selected cell in the specified column. Returns `-1` if
// there is no selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedRow(inColumn:)
func (b NSBrowser) SelectedRowInColumn(column int) int {
	rv := objc.Send[int](b.ID, objc.Sel("selectedRowInColumn:"), column)
	return rv
}
// Selects the cell at the specified row and column index.
//
// row: The row index of the cell to select.
//
// column: The column index of the cell to select.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/selectRow(_:inColumn:)
func (b NSBrowser) SelectRowInColumn(row int, column int) {
	objc.Send[objc.ID](b.ID, objc.Sel("selectRow:inColumn:"), row, column)
}
// Loads, if necessary, and returns the cell at the specified row and column
// location.
//
// row: The row index of the cell to return.
//
// col: The column index of the cell to return.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/loadedCell(atRow:column:)
func (b NSBrowser) LoadedCellAtRowColumn(row int, col int) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("loadedCellAtRow:column:"), row, col)
	return objectivec.Object{ID: rv}
}
// Begins editing the item at the specified path.
//
// indexPath: The path of the item.
//
// event: The event to use when beginning the edit.
//
// select: If [true], the cells contents will be selected; if [false], they will not
// be selected.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/editItem(at:with:select:)
func (b NSBrowser) EditItemAtIndexPathWithEventSelect(indexPath foundation.INSIndexPath, event INSEvent, select_ bool) {
	objc.Send[objc.ID](b.ID, objc.Sel("editItemAtIndexPath:withEvent:select:"), indexPath, event, select_)
}
// Returns the item at the specified index path.
//
// indexPath: The index path of the item to return.
//
// # Return Value
// 
// The item.
//
// # Discussion
// 
// This method can only be used if the delegate implements the item data
// source methods. The specified index path must be displayable in the
// browser.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/item(at:)
func (b NSBrowser) ItemAtIndexPath(indexPath foundation.INSIndexPath) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("itemAtIndexPath:"), indexPath)
	return objectivec.Object{ID: rv}
}
// Returns the item located at the specified row and column.
//
// row: The row of the item.
//
// column: The column of the item.
//
// # Return Value
// 
// The item.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/item(atRow:inColumn:)
func (b NSBrowser) ItemAtRowInColumn(row int, column int) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("itemAtRow:inColumn:"), row, column)
	return objectivec.Object{ID: rv}
}
// Returns the index path of the item whose children are displayed in the
// given column.
//
// column: The column to find the index path for.
//
// # Return Value
// 
// The index path of the column.
//
// # Discussion
// 
// This method can only be used if the delegate implements the item data
// source methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/indexPath(forColumn:)
func (b NSBrowser) IndexPathForColumn(column int) objc.ID {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("indexPathForColumn:"), column)
	return rv
}
// Returns whether the specified item is a leaf item.
//
// item: The item to be checked.
//
// # Return Value
// 
// [true] if the item is a leaf item; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method may return [false] if the item has never been displayed in the
// browser or accessed via [ItemAtIndexPath]. Overriding this method has no
// effect. It may be used only if the browser’s delegate implements the item
// data source methods.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/isLeafItem(_:)
func (b NSBrowser) IsLeafItem(item objectivec.IObject) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isLeafItem:"), item)
	return rv
}
// Returns the item that contains the children located in the specified
// column.
//
// column: The column.
//
// # Return Value
// 
// The parent item for the specified column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/parentForItems(inColumn:)
func (b NSBrowser) ParentForItemsInColumn(column int) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("parentForItemsInColumn:"), column)
	return objectivec.Object{ID: rv}
}
// Returns a string representing the browser’s current path.
//
// # Return Value
// 
// The path representing the current selection. The components of this path
// are separated with the string returned by [PathSeparator].
//
// # Discussion
// 
// Invoking this method is equivalent to invoking [PathToColumn] for all
// columns.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/path()
func (b NSBrowser) Path() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("path"))
	return foundation.NSStringFromID(rv).String()
}
// Sets the path to be displayed by the browser.
//
// path: The path to display. If `path` is prefixed by the path separator, the path
// is absolute, containing the full path from the browser’s first column.
// Otherwise, the path is relative, extending the browser’s current path
// starting at the last column.
//
// # Return Value
// 
// [true] if the given path is valid; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// While parsing `path`, the browser compares each component with the entries
// in the current column. If an exact match is found, the matching entry is
// selected, and the next component is compared to the next column’s
// entries. If no match is found for a component, the method exits and returns
// [false]; the final path is set to the valid portion of `path`. If each
// component of `path` specifies a valid branch or leaf in the browser’s
// hierarchy, the method returns [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/setPath(_:)
func (b NSBrowser) SetPath(path string) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("setPath:"), objc.String(path))
	return rv
}
// Returns a string representing the path from the first column up to, but not
// including, the column at the given index.
//
// column: The index of the column at which the path stops.
//
// # Return Value
// 
// The path of the current selection up to, but not including, the specified
// column. The components of this path are separated with the string returned
// by [PathSeparator].
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/path(toColumn:)
func (b NSBrowser) PathToColumn(column int) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("pathToColumn:"), column)
	return foundation.NSStringFromID(rv).String()
}
// Adds a column to the right of the last column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/addColumn()
func (b NSBrowser) AddColumn() {
	objc.Send[objc.ID](b.ID, objc.Sel("addColumn"))
}
// Validates the browser’s visible columns.
//
// # Discussion
// 
// This method invokes the delegate method [BrowserIsColumnValid]
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/validateVisibleColumns()
func (b NSBrowser) ValidateVisibleColumns() {
	objc.Send[objc.ID](b.ID, objc.Sel("validateVisibleColumns"))
}
// Loads column 0; unloads previously loaded columns.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/loadColumnZero()
func (b NSBrowser) LoadColumnZero() {
	objc.Send[objc.ID](b.ID, objc.Sel("loadColumnZero"))
}
// Reloads the given column.
//
// column: The index of the column to reload.
//
// # Discussion
// 
// If after reloading the selected item no longer exists in the column, the
// column is set to be the last column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/reloadColumn(_:)
func (b NSBrowser) ReloadColumn(column int) {
	objc.Send[objc.ID](b.ID, objc.Sel("reloadColumn:"), column)
}
// Returns the title displayed for the given column.
//
// column: The index of the column for which to get the title.
//
// # Return Value
// 
// The title of the specified column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/title(ofColumn:)
func (b NSBrowser) TitleOfColumn(column int) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("titleOfColumn:"), column)
	return foundation.NSStringFromID(rv).String()
}
// Sets the title of the given column.
//
// string: The title of the column.
//
// column: The index of the column whose title should be set.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/setTitle(_:ofColumn:)
func (b NSBrowser) SetTitleOfColumn(string_ string, column int) {
	objc.Send[objc.ID](b.ID, objc.Sel("setTitle:ofColumn:"), objc.String(string_), column)
}
// Draws the title for the specified column within the given rectangle.
//
// column: The index of the column for which to draw the title.
//
// rect: The rectangle within which to draw the title.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/drawTitle(ofColumn:in:)
func (b NSBrowser) DrawTitleOfColumnInRect(column int, rect corefoundation.CGRect) {
	objc.Send[objc.ID](b.ID, objc.Sel("drawTitleOfColumn:inRect:"), column, rect)
}
// Returns the bounds of the title frame for the specified column.
//
// column: The index of the column for which to return the title frame.
//
// # Return Value
// 
// The rectangle specifying the bounds of the column’s title frame.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/titleFrame(ofColumn:)
func (b NSBrowser) TitleFrameOfColumn(column int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b.ID, objc.Sel("titleFrameOfColumn:"), column)
	return corefoundation.CGRect(rv)
}
// Immediately retiles the browser’s columns using row heights specified by
// the browser’s delegate.
//
// indexSet: The indexes of the rows to resize.
//
// columnIndex: The column to retile.
//
// # Discussion
// 
// The browser’s delegate must implement
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/noteHeightOfRowsWithIndexesChanged(_:inColumn:)
func (b NSBrowser) NoteHeightOfRowsWithIndexesChangedInColumn(indexSet foundation.NSIndexSet, columnIndex int) {
	objc.Send[objc.ID](b.ID, objc.Sel("noteHeightOfRowsWithIndexesChanged:inColumn:"), indexSet, columnIndex)
}
// Updates the rows in the column with the specified column index with indexes
// in the specified set.
//
// rowIndexes: The set of row indexes of the rows to be updated.
//
// column: The column containing the rows to be updated.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/reloadData(forRowIndexes:inColumn:)
func (b NSBrowser) ReloadDataForRowIndexesInColumn(rowIndexes foundation.NSIndexSet, column int) {
	objc.Send[objc.ID](b.ID, objc.Sel("reloadDataForRowIndexes:inColumn:"), rowIndexes, column)
}
// Scrolls to make the specified column visible.
//
// column: The index of the column to scroll to.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/scrollColumnToVisible(_:)
func (b NSBrowser) ScrollColumnToVisible(column int) {
	objc.Send[objc.ID](b.ID, objc.Sel("scrollColumnToVisible:"), column)
}
// Scrolls columns left by the specified number of columns.
//
// shiftAmount: The number of columns by which to scroll the browser.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/scrollColumnsLeft(by:)
func (b NSBrowser) ScrollColumnsLeftBy(shiftAmount int) {
	objc.Send[objc.ID](b.ID, objc.Sel("scrollColumnsLeftBy:"), shiftAmount)
}
// Scrolls columns right by the specified number of columns.
//
// shiftAmount: The number of columns by which to scroll the browser.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/scrollColumnsRight(by:)
func (b NSBrowser) ScrollColumnsRightBy(shiftAmount int) {
	objc.Send[objc.ID](b.ID, objc.Sel("scrollColumnsRightBy:"), shiftAmount)
}
// Scrolls the specified row to be visible within the specified column.
//
// row: The index of the row to scroll.
//
// column: The index of the column containing the row to scroll.
//
// # Discussion
// 
// The row’s column will not be scrolled to visible via this method. To
// scroll the column to visible, use [ScrollColumnToVisible].
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/scrollRowToVisible(_:inColumn:)
func (b NSBrowser) ScrollRowToVisibleInColumn(row int, column int) {
	objc.Send[objc.ID](b.ID, objc.Sel("scrollRowToVisible:inColumn:"), row, column)
}
// Specifies the drag-operation mask for dragging operations with local or
// external destinations.
//
// mask: Dragging operation mask to use for either local or external drag
// operations, as specified by localDestination.
//
// isLocal: Indicates the location of the dragging operation’s destination object:
// 
// [true] for this application; [false] for another application.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/setDraggingSourceOperationMask(_:forLocal:)
func (b NSBrowser) SetDraggingSourceOperationMaskForLocal(mask NSDragOperation, isLocal bool) {
	objc.Send[objc.ID](b.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}
// Indicates whether the browser can attempt to initiate a drag of the given
// rows for the given event.
//
// rowIndexes: Rows the user is dragging
//
// column: Column containing the rows the user is dragging.
//
// event: Mouse-drag event.
//
// # Return Value
// 
// [true] when `rowIndexes` identifies at least one row and all the identified
// rows are enabled; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/canDragRows(with:inColumn:with:)
func (b NSBrowser) CanDragRowsWithIndexesInColumnWithEvent(rowIndexes foundation.NSIndexSet, column int, event INSEvent) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("canDragRowsWithIndexes:inColumn:withEvent:"), rowIndexes, column, event)
	return rv
}
// Provides an image to represent dragged rows during a drag operation on the
// browser.
//
// rowIndexes: Rows the user is dragging.
//
// column: Column with the rows the user is dragging.
//
// event: Mouse drag event.
//
// dragImageOffset: Offset for the returned image:
// 
// - [NSZeroPoint]: The image is centered under the pointer.
// //
// [NSZeroPoint]: https://developer.apple.com/documentation/Foundation/NSZeroPoint
//
// # Return Value
// 
// Image representing the visible cells identified by rowIndexes.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/draggingImageForRows(with:inColumn:with:offset:)
func (b NSBrowser) DraggingImageForRowsWithIndexesInColumnWithEventOffset(rowIndexes foundation.NSIndexSet, column int, event INSEvent, dragImageOffset foundation.NSPoint) INSImage {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("draggingImageForRowsWithIndexes:inColumn:withEvent:offset:"), rowIndexes, column, event, dragImageOffset)
	return NSImageFromID(rv)
}
// Returns the rectangle containing the given column.
//
// column: The index of the column for which to retrieve the frame.
//
// # Return Value
// 
// The rectangle containing the specified column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/frame(ofColumn:)
func (b NSBrowser) FrameOfColumn(column int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b.ID, objc.Sel("frameOfColumn:"), column)
	return corefoundation.CGRect(rv)
}
// Returns the rectangle containing the specified column, not including
// borders.
//
// column: The index of the column for which to retrieve the inside frame.
//
// # Return Value
// 
// The rectangle containing the column, not including the column borders.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/frame(ofInsideOfColumn:)
func (b NSBrowser) FrameOfInsideOfColumn(column int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b.ID, objc.Sel("frameOfInsideOfColumn:"), column)
	return corefoundation.CGRect(rv)
}
// Returns the frame of the cell at the specified location, including the
// expandable arrow.
//
// row: The row of the cell.
//
// column: The column of the cell.
//
// # Return Value
// 
// The frame of the cell, in the [NSBrowser] coordinate space.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/frame(ofRow:inColumn:)
func (b NSBrowser) FrameOfRowInColumn(row int, column int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b.ID, objc.Sel("frameOfRow:inColumn:"), row, column)
	return corefoundation.CGRect(rv)
}
// Gets the row and column coordinates for the specified point, if a cell
// exists at that point.
//
// row: On output, the row number of the cell at the specified point, or `-1` if
// there is no cell at the point.
//
// column: On output, he column number of the cell at the specified point, or `-1` if
// there is no cell at the point.
//
// point: The point to test.
//
// # Return Value
// 
// [true] if a cell exists at the specified point; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If a row does not exist at `point`, then `-1` is set for the row. If a
// column does not exist at `point`, then `-1` is set for the column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/getRow(_:column:for:)
func (b NSBrowser) GetRowColumnForPoint(point corefoundation.CGPoint) (int, int, bool) {
	var row int
	var column int
	rv := objc.Send[bool](b.ID, objc.Sel("getRow:column:forPoint:"), unsafe.Pointer(&row), unsafe.Pointer(&column), point)
	return row, column, rv
}
// Sends the action message to the target.
//
// # Return Value
// 
// [true] if successful; [false] if no target for the message could be found.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/sendAction()
func (b NSBrowser) SendAction() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("sendAction"))
	return rv
}
// Responds to (single) mouse clicks in a column of the browser.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/doClick(_:)
func (b NSBrowser) DoClick(sender objectivec.IObject) {
	objc.Send[objc.ID](b.ID, objc.Sel("doClick:"), sender)
}
// Responds to double clicks in a column of the browser.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/doDoubleClick(_:)
func (b NSBrowser) DoDoubleClick(sender objectivec.IObject) {
	objc.Send[objc.ID](b.ID, objc.Sel("doDoubleClick:"), sender)
}
// Returns the content width for a given column width.
//
// columnWidth: The width of the column. This width is the entire scrolling text view.
//
// # Return Value
// 
// The width of the content for the column. This is the width of the matrix in
// the column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/columnContentWidth(forColumnWidth:)
func (b NSBrowser) ColumnContentWidthForColumnWidth(columnWidth float64) float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("columnContentWidthForColumnWidth:"), columnWidth)
	return rv
}
// Returns the column width for the width of the given column’s content.
//
// columnContentWidth: The width of the column’s content (the width of the matrix in the
// column).
//
// # Return Value
// 
// The width of the column (the width of the entire scrolling text view).
//
// # Discussion
// 
// For example, to guarantee that 16 pixels of your browser cell are always
// visible, call:
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/columnWidth(forColumnContentWidth:)
func (b NSBrowser) ColumnWidthForColumnContentWidth(columnContentWidth float64) float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("columnWidthForColumnContentWidth:"), columnContentWidth)
	return rv
}
// Returns the width of the specified column.
//
// column: The index of the column for which to retrieve the width.
//
// # Return Value
// 
// The width of the column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/width(ofColumn:)
func (b NSBrowser) WidthOfColumn(column int) float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("widthOfColumn:"), column)
	return rv
}
// Sets the width of the specified column.
//
// columnWidth: The new width of the specified column.
//
// columnIndex: The index of the column for which to set the width.
//
// # Discussion
// 
// This method can be used to set the initial width of browser columns unless
// the column sizing is automatic; [SetWidthOfColumn] does nothing if
// [ColumnResizingType] is [NSBrowser.ColumnResizingType.autoColumnResizing].
// To set the default width for new columns (that don’t otherwise have
// initial widths from defaults or via the delegate), use a `columnIndex` of
// –1. A value set for `columnIndex` of –1 is persistent. An
// [columnConfigurationDidChangeNotification] notification is posted (not
// immediately), if necessary, so that the browser can autosave the new column
// configuration.
//
// [NSBrowser.ColumnResizingType.autoColumnResizing]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum/autoColumnResizing
// [columnConfigurationDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSBrowser/columnConfigurationDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/setWidth(_:ofColumn:)
func (b NSBrowser) SetWidthOfColumn(columnWidth float64, columnIndex int) {
	objc.Send[objc.ID](b.ID, objc.Sel("setWidth:ofColumn:"), columnWidth, columnIndex)
}
// Returns the default column width of the browser’s columns.
//
// # Return Value
// 
// The default column width.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/defaultColumnWidth()
func (b NSBrowser) DefaultColumnWidth() float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("defaultColumnWidth"))
	return rv
}
// Sets the default column width for new browser columns that do not otherwise
// have an initial width from defaults or the browser’s delegate.
//
// columnWidth: The default column width to set.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/setDefaultColumnWidth(_:)
func (b NSBrowser) SetDefaultColumnWidth(columnWidth float64) {
	objc.Send[objc.ID](b.ID, objc.Sel("setDefaultColumnWidth:"), columnWidth)
}

// Removes the column configuration data stored under the given name from the
// application’s user defaults.
//
// name: The name of the column configuration data to remove.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/removeSavedColumns(withAutosaveName:)
func (_NSBrowserClass NSBrowserClass) RemoveSavedColumnsWithAutosaveName(name NSBrowserColumnsAutosaveName) {
	objc.Send[objc.ID](objc.ID(_NSBrowserClass.class), objc.Sel("removeSavedColumnsWithAutosaveName:"), objc.String(string(name)))
}

// A Boolean that indicates whether the browser reuses matrix objects after
// their columns are unloaded.
//
// # Discussion
// 
// When the value of this property is [true], the [NSMatrix] objects aren’t
// freed when their columns are unloaded, so they can be reused.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/reusesColumns
func (b NSBrowser) ReusesColumns() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("reusesColumns"))
	return rv
}
func (b NSBrowser) SetReusesColumns(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setReusesColumns:"), value)
}
// The maximum number of visible columns.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/maxVisibleColumns
func (b NSBrowser) MaxVisibleColumns() int {
	rv := objc.Send[int](b.ID, objc.Sel("maxVisibleColumns"))
	return rv
}
func (b NSBrowser) SetMaxVisibleColumns(value int) {
	objc.Send[struct{}](b.ID, objc.Sel("setMaxVisibleColumns:"), value)
}
// A Boolean that indicates whether the browser automatically hides its
// scroller.
//
// # Discussion
// 
// When the value of this property is [true], the scroller is automatically
// hidden.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/autohidesScroller
func (b NSBrowser) AutohidesScroller() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("autohidesScroller"))
	return rv
}
func (b NSBrowser) SetAutohidesScroller(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setAutohidesScroller:"), value)
}
// The browser’s background color.
//
// # Discussion
// 
// The default value of this property is `[NSColor whiteColor]`. `[NSColor
// clearColor]` specifies a transparent background.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/backgroundColor
func (b NSBrowser) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (b NSBrowser) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](b.ID, objc.Sel("setBackgroundColor:"), value)
}
// The minimum column width, in pixels.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/minColumnWidth
func (b NSBrowser) MinColumnWidth() float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("minColumnWidth"))
	return rv
}
func (b NSBrowser) SetMinColumnWidth(value float64) {
	objc.Send[struct{}](b.ID, objc.Sel("setMinColumnWidth:"), value)
}
// A Boolean that indicates whether columns are separated by bezeled borders.
//
// # Discussion
// 
// When the value of this property is [true], the browser’s columns are
// separated by bezeled borders.
// 
// This value is ignored if [Titled] does not return [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/separatesColumns
func (b NSBrowser) SeparatesColumns() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("separatesColumns"))
	return rv
}
func (b NSBrowser) SetSeparatesColumns(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setSeparatesColumns:"), value)
}
// A Boolean that indicates whether a column takes its title from the selected
// cell in the previous column.
//
// # Discussion
// 
// When the value of this property is [true], the title of a column is set to
// the string value of the selected [NSCell] in the previous column.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/takesTitleFromPreviousColumn
func (b NSBrowser) TakesTitleFromPreviousColumn() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("takesTitleFromPreviousColumn"))
	return rv
}
func (b NSBrowser) SetTakesTitleFromPreviousColumn(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setTakesTitleFromPreviousColumn:"), value)
}
// The browser’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/delegate
func (b NSBrowser) Delegate() NSBrowserDelegate {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("delegate"))
	return NSBrowserDelegateObjectFromID(rv)
}
func (b NSBrowser) SetDelegate(value NSBrowserDelegate) {
	objc.Send[struct{}](b.ID, objc.Sel("setDelegate:"), value)
}
// The prototype [NSCell] for displaying items in the matrices in the columns
// of the browser.
//
// # Discussion
// 
// The prototype [NSCell] instance is copied to display items in the matrices
// of the browser.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/cellPrototype
func (b NSBrowser) CellPrototype() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("cellPrototype"))
	return objectivec.Object{ID: rv}
}
func (b NSBrowser) SetCellPrototype(value objectivec.IObject) {
	objc.Send[struct{}](b.ID, objc.Sel("setCellPrototype:"), value)
}
// A Boolean that indicates whether the user can select branch items.
//
// # Discussion
// 
// When the value of this property is [true], the user can select branch items
// when multiple selection is enabled.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsBranchSelection
func (b NSBrowser) AllowsBranchSelection() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("allowsBranchSelection"))
	return rv
}
func (b NSBrowser) SetAllowsBranchSelection(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setAllowsBranchSelection:"), value)
}
// A Boolean that indicates whether there can be nothing selected.
//
// # Discussion
// 
// When the value of this property is [true], the browser allows the selection
// to be empty.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsEmptySelection
func (b NSBrowser) AllowsEmptySelection() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("allowsEmptySelection"))
	return rv
}
func (b NSBrowser) SetAllowsEmptySelection(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setAllowsEmptySelection:"), value)
}
// A Boolean that indicates whether the user can select multiple items.
//
// # Discussion
// 
// When the value of this property is [true], the browser allows the user to
// select multiple items at once.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsMultipleSelection
func (b NSBrowser) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}
func (b NSBrowser) SetAllowsMultipleSelection(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}
// A Boolean that indicates whether the browser allows keystroke-based
// selection (type select).
//
// # Discussion
// 
// When the value of this property is [true], the browser allows
// keystroke-based selection. The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/allowsTypeSelect
func (b NSBrowser) AllowsTypeSelect() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("allowsTypeSelect"))
	return rv
}
func (b NSBrowser) SetAllowsTypeSelect(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setAllowsTypeSelect:"), value)
}
// All cells selected in the rightmost column.
//
// # Discussion
// 
// When the array is empty, there is no selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedCells
func (b NSBrowser) SelectedCells() []NSCell {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("selectedCells"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSCell {
		return NSCellFromID(id)
	})
}
// The index path of the item selected in the browser.
//
// # Discussion
// 
// When the value of this property is `nil`, there is no selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/selectionIndexPath
func (b NSBrowser) SelectionIndexPath() objc.ID {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("selectionIndexPath"))
	return rv
}
func (b NSBrowser) SetSelectionIndexPath(value objc.ID) {
	objc.Send[struct{}](b.ID, objc.Sel("setSelectionIndexPath:"), value)
}
// An array containing the index paths of all items selected in the browser.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/selectionIndexPaths
func (b NSBrowser) SelectionIndexPaths() []objc.ID {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("selectionIndexPaths"))
	return rv
}
func (b NSBrowser) SetSelectionIndexPaths(value []objc.ID) {
	objc.Send[struct{}](b.ID, objc.Sel("setSelectionIndexPaths:"), value)
}
// The path separator.
//
// # Discussion
// 
// The default value of this property is `/`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/pathSeparator
func (b NSBrowser) PathSeparator() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("pathSeparator"))
	return foundation.NSStringFromID(rv).String()
}
func (b NSBrowser) SetPathSeparator(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setPathSeparator:"), objc.String(value))
}
// The index of the last column with a selected item.
//
// # Discussion
// 
// When the value of this property is `-1`, there is no column selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/selectedColumn
func (b NSBrowser) SelectedColumn() int {
	rv := objc.Send[int](b.ID, objc.Sel("selectedColumn"))
	return rv
}
// The index of the last column loaded.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/lastColumn
func (b NSBrowser) LastColumn() int {
	rv := objc.Send[int](b.ID, objc.Sel("lastColumn"))
	return rv
}
func (b NSBrowser) SetLastColumn(value int) {
	objc.Send[struct{}](b.ID, objc.Sel("setLastColumn:"), value)
}
// The index of the first visible column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/firstVisibleColumn
func (b NSBrowser) FirstVisibleColumn() int {
	rv := objc.Send[int](b.ID, objc.Sel("firstVisibleColumn"))
	return rv
}
// The number of visible columns.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/numberOfVisibleColumns
func (b NSBrowser) NumberOfVisibleColumns() int {
	rv := objc.Send[int](b.ID, objc.Sel("numberOfVisibleColumns"))
	return rv
}
// The index of the last visible column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/lastVisibleColumn
func (b NSBrowser) LastVisibleColumn() int {
	rv := objc.Send[int](b.ID, objc.Sel("lastVisibleColumn"))
	return rv
}
// A Boolean that indicates whether column 0 is loaded.
//
// # Discussion
// 
// When the value of this property is [true], column 0 is loaded.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/isLoaded
func (b NSBrowser) Loaded() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isLoaded"))
	return rv
}
// A Boolean that indicates whether columns display titles.
//
// # Discussion
// 
// When the value of this property is [true], the columns in a browser display
// titles.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/isTitled
func (b NSBrowser) Titled() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isTitled"))
	return rv
}
func (b NSBrowser) SetTitled(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setTitled:"), value)
}
// The height of the column titles for the browser.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/titleHeight
func (b NSBrowser) TitleHeight() float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("titleHeight"))
	return rv
}
// A Boolean that indicates whether the browser has a horizontal scroller.
//
// # Discussion
// 
// When the value of this property is [true], the browser uses an [NSScroller]
// object to scroll horizontally.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/hasHorizontalScroller
func (b NSBrowser) HasHorizontalScroller() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("hasHorizontalScroller"))
	return rv
}
func (b NSBrowser) SetHasHorizontalScroller(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setHasHorizontalScroller:"), value)
}
// The browser’s double-click action method.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/doubleAction
func (b NSBrowser) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](b.ID, objc.Sel("doubleAction"))
	return rv
}
func (b NSBrowser) SetDoubleAction(value objc.SEL) {
	objc.Send[struct{}](b.ID, objc.Sel("setDoubleAction:"), value)
}
// A Boolean that indicates whether pressing an arrow key causes an action
// message to be sent.
//
// # Discussion
// 
// When the value of this property is [false], pressing an arrow key scrolls
// the browser. When the value of this property is [true], it also sends the
// action message specified by [Action].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/sendsActionOnArrowKeys
func (b NSBrowser) SendsActionOnArrowKeys() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("sendsActionOnArrowKeys"))
	return rv
}
func (b NSBrowser) SetSendsActionOnArrowKeys(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setSendsActionOnArrowKeys:"), value)
}
// The column number of the cell that the user clicked to display a context
// menu.
//
// # Discussion
// 
// The value of this property is `-1` if no context menu is active.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/clickedColumn
func (b NSBrowser) ClickedColumn() int {
	rv := objc.Send[int](b.ID, objc.Sel("clickedColumn"))
	return rv
}
// The row number of the cell that the user clicked to display a context menu.
//
// # Discussion
// 
// The value of this property is `-1` if no context menu is active.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/clickedRow
func (b NSBrowser) ClickedRow() int {
	rv := objc.Send[int](b.ID, objc.Sel("clickedRow"))
	return rv
}
// The name used to automatically save the browser’s column configuration.
//
// # Discussion
// 
// Column configuration is defined as an array of column content widths. One
// width is saved for each level the user has reached. That is, the browser
// saves column width based on depth, not on unique paths. To do more complex
// column persistence, you should register for
// [columnConfigurationDidChangeNotification] and handle persistence yourself.
// This setting is persistent.
// 
// When this property is set to a value different than its current value, this
// property also reads in any column configuration data previously saved under
// the new value and applies the values to the browser.
//
// [columnConfigurationDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSBrowser/columnConfigurationDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/columnsAutosaveName-swift.property
func (b NSBrowser) ColumnsAutosaveName() NSBrowserColumnsAutosaveName {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("columnsAutosaveName"))
	return NSBrowserColumnsAutosaveName(foundation.NSStringFromID(rv).String())
}
func (b NSBrowser) SetColumnsAutosaveName(value NSBrowserColumnsAutosaveName) {
	objc.Send[struct{}](b.ID, objc.Sel("setColumnsAutosaveName:"), objc.String(string(value)))
}
// A constant indicating the browser’s column resizing type.
//
// # Discussion
// 
// See [NSBrowser.ColumnResizingType] for possible values. The default value
// of this property is [NSBrowser.ColumnResizingType.autoColumnResizing].
//
// [NSBrowser.ColumnResizingType.autoColumnResizing]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum/autoColumnResizing
// [NSBrowser.ColumnResizingType]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/columnResizingType-swift.property
func (b NSBrowser) ColumnResizingType() NSBrowserColumnResizingType {
	rv := objc.Send[NSBrowserColumnResizingType](b.ID, objc.Sel("columnResizingType"))
	return NSBrowserColumnResizingType(rv)
}
func (b NSBrowser) SetColumnResizingType(value NSBrowserColumnResizingType) {
	objc.Send[struct{}](b.ID, objc.Sel("setColumnResizingType:"), value)
}
// A Boolean that indicates whether the browser is set to resize all columns
// simultaneously rather than resizing a single column at a time.
//
// # Discussion
// 
// When the value of this property is [true], the browser is set to resize all
// columns simultaneously. The default value of this property is [false].
// 
// This setting applies only to browsers that allow the user to resize columns
// (see the constant [NSBrowser.ColumnResizingType.userColumnResizing]).
// Holding down the Option key while resizing switches the type of resizing
// used. This setting is persistent.
//
// [NSBrowser.ColumnResizingType.userColumnResizing]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum/userColumnResizing
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/prefersAllColumnUserResizing
func (b NSBrowser) PrefersAllColumnUserResizing() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("prefersAllColumnUserResizing"))
	return rv
}
func (b NSBrowser) SetPrefersAllColumnUserResizing(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setPrefersAllColumnUserResizing:"), value)
}
// The height of the browser’s rows.
//
// # Discussion
// 
// The value of this property must be greater than `0`. The default value of
// this property is `17.0`. Any fractional value will be forced to an integral
// value for drawing. For variable row height browsers (ones whose delegates
// implement [BrowserHeightOfRowInColumn]), the row height will be used to
// draw alternating rows past the last row in each browser column.
// 
// This property is only available when using the item delegate methods. An
// exception is thrown if you are using the matrix delegate methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowser/rowHeight
func (b NSBrowser) RowHeight() float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("rowHeight"))
	return rv
}
func (b NSBrowser) SetRowHeight(value float64) {
	objc.Send[struct{}](b.ID, objc.Sel("setRowHeight:"), value)
}


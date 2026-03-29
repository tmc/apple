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

// The class instance for the [NSTableView] class.
var (
	_NSTableViewClass     NSTableViewClass
	_NSTableViewClassOnce sync.Once
)

func getNSTableViewClass() NSTableViewClass {
	_NSTableViewClassOnce.Do(func() {
		_NSTableViewClass = NSTableViewClass{class: objc.GetClass("NSTableView")}
	})
	return _NSTableViewClass
}

// GetNSTableViewClass returns the class object for NSTableView.
func GetNSTableViewClass() NSTableViewClass {
	return getNSTableViewClass()
}

type NSTableViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTableViewClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTableViewClass) Alloc() NSTableView {
	rv := objc.Send[NSTableView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A set of related records, displayed in rows that represent individual
// records and columns that represent the attributes of those records.
//
// # Overview
// 
// Table views are displayed in scroll views. Beginning with macOS v10.7, you
// can use [NSView] objects (most commonly customized [NSTableCellView]
// objects) instead of cells for specifying rows and columns. You can still
// use [NSCell] objects for each row and column item if you prefer.
// 
// A table view does not store its own data; it retrieves data values as
// needed from a data source to which it has a weak reference. You should not,
// therefore, directly set data values programmatically in the table view;
// instead, modify the values in the data source and allow the changes to be
// reflected in the table view. To learn about the methods that an
// [NSTableView] object uses to provide and access the contents of its data
// source object, see [NSTableViewDataSource].
// 
// To customize a table view’s behavior without subclassing [NSTableView],
// use the methods defined by the [NSTableViewDelegate] protocol. For example,
// the delegate supports table column management, type-to-select
// functionality, row selection and editing, custom tracking, and custom views
// for individual columns and rows. To learn more about the table view
// delegate, see [NSTableViewDelegate].
// 
// # Subclassing
// 
// Subclassing [NSTableView] is usually not necessary. Instead, you customize
// the table view using a delegate object (an object conforming to the
// [NSTableViewDelegate] protocol) and a data source object (conforming to the
// [NSTableViewDataSource] protocol), or by subclassing one of the following
// subcomponents: cells (when using [NSCell]-based table views), the row cell
// view or the row view (when using [NSView]-based table views), the table
// column class, or table column header classes.
// 
// # Enabling the Table View
// 
// Use the [Enabled] property to enable or disable the table view, which the
// view inherits from [NSControl]. This property affects the visual appearance
// of the table view differently depending on whether you use a view- or a
// cell-based table view. When you change the property’s value for a
// cell-based table view, the system manages the visual appearance of that
// table view’s rows, and updates them to a state that reflects the value.
// Because view-based table views permit complex items in their cells, it’s
// the developer’s responsibility to update each cell’s appearance as
// appropriate.
//
// # Managing the Table’s Data
//
//   - [NSTableView.DataSource]: The object that provides the data displayed by the table view.
//   - [NSTableView.SetDataSource]
//   - [NSTableView.UsesStaticContents]: A Boolean value indicating whether the table uses static data.
//   - [NSTableView.SetUsesStaticContents]
//   - [NSTableView.ReloadData]: Marks the table view as needing redisplay, so it will reload the data for visible cells and draw the new values.
//   - [NSTableView.ReloadDataForRowIndexesColumnIndexes]: Reloads the data for only the specified rows and columns.
//
// # Creating Views to Display
//
//   - [NSTableView.MakeViewWithIdentifierOwner]: Returns a new or existing view with the specified identifier.
//   - [NSTableView.RowViewAtRowMakeIfNecessary]: Returns a row view at the specified index, creating one if necessary.
//   - [NSTableView.ViewAtColumnRowMakeIfNecessary]: Returns a view at the specified row and column indexes, creating one if necessary.
//
// # Updating the Table View Arrangement
//
//   - [NSTableView.BeginUpdates]: Begins a group of updates for the table view.
//   - [NSTableView.EndUpdates]: Ends the group of updates for the table view.
//   - [NSTableView.MoveRowAtIndexToIndex]: Moves the specified row to the new row location using animation.
//   - [NSTableView.InsertRowsAtIndexesWithAnimation]: Inserts the rows using the specified animation.
//   - [NSTableView.RemoveRowsAtIndexesWithAnimation]: Removes the rows using the specified animation.
//   - [NSTableView.RowForView]: Returns the index of the row for the specified view.
//   - [NSTableView.ColumnForView]: Returns the column index for the specified view.
//
// # NSView-Based Table Nib File Registration
//
//   - [NSTableView.RegisterNibForIdentifier]: Registers a NIB for the specified identifier, so that view-based table views can use it to instantiate views.
//   - [NSTableView.RegisteredNibsByIdentifier]: The dictionary of all registered nib files for view-based table view identifiers.
//
// # Target-action Behavior
//
//   - [NSTableView.DoubleAction]: The message sent to the table view’s target when the user double-clicks a cell or column header.
//   - [NSTableView.SetDoubleAction]
//   - [NSTableView.ClickedColumn]: The index of the column the user clicked.
//   - [NSTableView.ClickedRow]: The index of the row the user clicked.
//
// # Configuring Behavior
//
//   - [NSTableView.AllowsColumnReordering]: A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers.
//   - [NSTableView.SetAllowsColumnReordering]
//   - [NSTableView.AllowsColumnResizing]: A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers.
//   - [NSTableView.SetAllowsColumnResizing]
//   - [NSTableView.AllowsMultipleSelection]: A Boolean value indicating whether the table view allows the user to select more than one column or row at a time.
//   - [NSTableView.SetAllowsMultipleSelection]
//   - [NSTableView.AllowsEmptySelection]: A Boolean value indicating whether the table view allows the user to select zero columns or rows.
//   - [NSTableView.SetAllowsEmptySelection]
//   - [NSTableView.AllowsColumnSelection]: A Boolean value indicating whether the table view allows the user to select columns by clicking their headers.
//   - [NSTableView.SetAllowsColumnSelection]
//   - [NSTableView.UsesAutomaticRowHeights]: A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows.
//   - [NSTableView.SetUsesAutomaticRowHeights]
//
// # Setting Display Attributes
//
//   - [NSTableView.IntercellSpacing]: The horizontal and vertical spacing between cells.
//   - [NSTableView.SetIntercellSpacing]
//   - [NSTableView.RowHeight]: The height of each row in the table.
//   - [NSTableView.SetRowHeight]
//   - [NSTableView.BackgroundColor]: The color used to draw the background of the table.
//   - [NSTableView.SetBackgroundColor]
//   - [NSTableView.UsesAlternatingRowBackgroundColors]: A Boolean value indicating whether the table view uses alternating row colors for its background.
//   - [NSTableView.SetUsesAlternatingRowBackgroundColors]
//   - [NSTableView.Style]: The style that the table view uses.
//   - [NSTableView.SetStyle]
//   - [NSTableView.EffectiveStyle]: The effective style that the table uses.
//   - [NSTableView.SelectionHighlightStyle]: The selection highlight style used by the table view to indicate row and column selection.
//   - [NSTableView.SetSelectionHighlightStyle]
//   - [NSTableView.GridColor]: The color used to draw grid lines.
//   - [NSTableView.SetGridColor]
//   - [NSTableView.GridStyleMask]: The grid lines drawn by the table view.
//   - [NSTableView.SetGridStyleMask]
//   - [NSTableView.IndicatorImageInTableColumn]: Returns the indicator image of the specified table column.
//   - [NSTableView.SetIndicatorImageInTableColumn]: Sets the indicator image of the specified column.
//
// # Getting and Setting Row Size Styles
//
//   - [NSTableView.EffectiveRowSizeStyle]: The effective row size style for the table.
//   - [NSTableView.RowSizeStyle]: The row size style (small, medium, large, or custom) used by the table view.
//   - [NSTableView.SetRowSizeStyle]
//
// # Column Management
//
//   - [NSTableView.AddTableColumn]: Adds the specified column as the last column of the table view.
//   - [NSTableView.RemoveTableColumn]: Removes the specified column from the table view.
//   - [NSTableView.MoveColumnToColumn]: Moves the column and heading at the specified index to the new specified index.
//   - [NSTableView.TableColumns]: An array containing the current table column objects.
//   - [NSTableView.ColumnWithIdentifier]: Returns the index of the first column in the table view whose identifier is equal to the specified identifier.
//   - [NSTableView.TableColumnWithIdentifier]: Returns the [NSTableColumn] object for the first column whose identifier is equal to the specified object.
//
// # Selecting Columns and Rows
//
//   - [NSTableView.SelectColumnIndexesByExtendingSelection]: Sets the column selection using `indexes` possibly extending the selection.
//   - [NSTableView.SelectedColumn]: The index of the last selected column (or the last column added to the selection).
//   - [NSTableView.SelectedColumnIndexes]: An index set containing the indexes of the selected columns.
//   - [NSTableView.DeselectColumn]: Deselects the column at the specified index if it’s selected.
//   - [NSTableView.NumberOfSelectedColumns]: The number of selected columns.
//   - [NSTableView.IsColumnSelected]: Returns a Boolean value that indicates whether the column at the specified index is selected.
//   - [NSTableView.SelectRowIndexesByExtendingSelection]: Sets the row selection using `indexes` extending the selection if specified.
//   - [NSTableView.SelectedRow]: The index of the last selected row (or the last row added to the selection).
//   - [NSTableView.SelectedRowIndexes]: An index set containing the indexes of the selected rows.
//   - [NSTableView.DeselectRow]: Deselects the row at the specified index if it’s selected.
//   - [NSTableView.NumberOfSelectedRows]: The number of selected rows.
//   - [NSTableView.IsRowSelected]: Returns a Boolean value that indicates whether the row at the specified index is selected.
//   - [NSTableView.DeselectAll]: Deselects all selected rows or columns if empty selection is allowed; otherwise does nothing.
//
// # Enumerating Table Rows
//
//   - [NSTableView.EnumerateAvailableRowViewsUsingBlock]: Allows the enumeration of all the table rows that are known to the table view.
//
// # Managing Type Select
//
//   - [NSTableView.AllowsTypeSelect]: A Boolean value indicating whether the table view allows the user to type characters to select rows.
//   - [NSTableView.SetAllowsTypeSelect]
//
// # Table Dimensions
//
//   - [NSTableView.NumberOfColumns]: The number of columns in the table.
//   - [NSTableView.NumberOfRows]: The number of rows in the table.
//
// # Getting and Setting Floating Rows
//
//   - [NSTableView.FloatsGroupRows]: A Boolean value indicating whether the table view draws grouped rows as if they are floating.
//   - [NSTableView.SetFloatsGroupRows]
//
// # Editing Cells
//
//   - [NSTableView.EditColumnRowWithEventSelect]: Edits the cell at the specified column and row using the specified event and selection behavior.
//   - [NSTableView.EditedColumn]: The index of the column being edited.
//   - [NSTableView.EditedRow]: The index of the row being edited.
//
// # Adding and Deleting Row Views
//
//   - [NSTableView.DidAddRowViewForRow]: Invoked when a row view is added to the table.
//   - [NSTableView.DidRemoveRowViewForRow]: Invoked when a row view is removed from the table.
//
// # Setting Auxiliary Views
//
//   - [NSTableView.HeaderView]: The view object used to draw headers over columns.
//   - [NSTableView.SetHeaderView]
//   - [NSTableView.CornerView]: The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view.
//   - [NSTableView.SetCornerView]
//
// # Layout Support
//
//   - [NSTableView.RectOfColumn]: Returns the rectangle containing the column at the specified index.
//   - [NSTableView.RectOfRow]: Returns the rectangle containing the row at the specified index.
//   - [NSTableView.RowsInRect]: Returns a range of indexes for the rows that lie wholly or partially within the vertical boundaries of the specified rectangle.
//   - [NSTableView.ColumnIndexesInRect]: Returns the indexes of the table view’s columns that intersect the specified rectangle.
//   - [NSTableView.ColumnAtPoint]: Returns the index of the column the specified point lies in.
//   - [NSTableView.RowAtPoint]: Returns the index of the row the specified point lies in.
//   - [NSTableView.FrameOfCellAtColumnRow]: Returns a rectangle locating the cell that lies at the intersection of the specified column and row.
//   - [NSTableView.ColumnAutoresizingStyle]: The table view’s column autoresizing style.
//   - [NSTableView.SetColumnAutoresizingStyle]
//   - [NSTableView.SizeLastColumnToFit]: Resizes the last column so the table view fits exactly within its enclosing clip view.
//   - [NSTableView.NoteNumberOfRowsChanged]: Informs the table view that the number of records in its data source has changed.
//   - [NSTableView.Tile]: Properly sizes the table view and its header view and marks it as needing display.
//   - [NSTableView.NoteHeightOfRowsWithIndexesChanged]: Informs the table view that the rows specified in `indexSet` have changed height.
//
// # Drawing
//
//   - [NSTableView.DrawRowClipRect]: Draws the cells for the row at `rowIndex` in the columns that intersect `clipRect`.
//   - [NSTableView.DrawGridInClipRect]: Draws the grid lines within the supplied rectangle.
//   - [NSTableView.HighlightSelectionInClipRect]: Highlights the region of the table view in the specified rectangle.
//   - [NSTableView.DrawBackgroundInClipRect]: Draws the background of the table view in the clip rect specified by the rectangle.
//
// # Scrolling
//
//   - [NSTableView.ScrollRowToVisible]: Scrolls the view so the specified row is visible.
//   - [NSTableView.ScrollColumnToVisible]: Scrolls the view so the specified column is visible.
//
// # Table Column State Persistence
//
//   - [NSTableView.AutosaveTableColumns]: A Boolean value indicating whether the order and width of the table view’s columns are automatically saved.
//   - [NSTableView.SetAutosaveTableColumns]
//   - [NSTableView.AutosaveName]: The name under which table information is automatically saved.
//   - [NSTableView.SetAutosaveName]
//
// # Accessing the Delegate
//
//   - [NSTableView.Delegate]: The table view’s delegate.
//   - [NSTableView.SetDelegate]
//
// # Highlightable Column Headers
//
//   - [NSTableView.HighlightedTableColumn]: The column highlighted in the table.
//   - [NSTableView.SetHighlightedTableColumn]
//
// # Dragging
//
//   - [NSTableView.DragImageForRowsWithIndexesTableColumnsEventOffset]: Computes and returns an image to use for dragging.
//   - [NSTableView.CanDragRowsWithIndexesAtPoint]: Returns a Boolean value indicating whether the table view allows dragging the rows with the drag initiated at the specified point.
//   - [NSTableView.SetDraggingSourceOperationMaskForLocal]: Sets the default operation mask returned by `` to `mask`.
//   - [NSTableView.VerticalMotionCanBeginDrag]: A Boolean value indicating whether vertical motion is treated as a drag or selection change.
//   - [NSTableView.SetVerticalMotionCanBeginDrag]
//   - [NSTableView.DraggingDestinationFeedbackStyle]: The feedback style displayed when the user drags over the table view.
//   - [NSTableView.SetDraggingDestinationFeedbackStyle]
//   - [NSTableView.SetDropRowDropOperation]: Retargets the proposed drop operation.
//
// # Sorting
//
//   - [NSTableView.SortDescriptors]: The table view’s sort descriptors.
//   - [NSTableView.SetSortDescriptors]
//
// # Row Actions
//
//   - [NSTableView.RowActionsVisible]: A Boolean value indicating whether a table row’s actions are visible.
//   - [NSTableView.SetRowActionsVisible]
//
// # Hiding and Showing Table Rows
//
//   - [NSTableView.HideRowsAtIndexesWithAnimation]: Hides the specified table rows.
//   - [NSTableView.UnhideRowsAtIndexesWithAnimation]: Unhides the specified table rows.
//   - [NSTableView.HiddenRowIndexes]: The indexes of all hidden table rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView
type NSTableView struct {
	NSControl
}

// NSTableViewFromID constructs a [NSTableView] from an objc.ID.
//
// A set of related records, displayed in rows that represent individual
// records and columns that represent the attributes of those records.
func NSTableViewFromID(id objc.ID) NSTableView {
	return NSTableView{NSControl: NSControlFromID(id)}
}
// NOTE: NSTableView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTableView] class.
//
// # Managing the Table’s Data
//
//   - [INSTableView.DataSource]: The object that provides the data displayed by the table view.
//   - [INSTableView.SetDataSource]
//   - [INSTableView.UsesStaticContents]: A Boolean value indicating whether the table uses static data.
//   - [INSTableView.SetUsesStaticContents]
//   - [INSTableView.ReloadData]: Marks the table view as needing redisplay, so it will reload the data for visible cells and draw the new values.
//   - [INSTableView.ReloadDataForRowIndexesColumnIndexes]: Reloads the data for only the specified rows and columns.
//
// # Creating Views to Display
//
//   - [INSTableView.MakeViewWithIdentifierOwner]: Returns a new or existing view with the specified identifier.
//   - [INSTableView.RowViewAtRowMakeIfNecessary]: Returns a row view at the specified index, creating one if necessary.
//   - [INSTableView.ViewAtColumnRowMakeIfNecessary]: Returns a view at the specified row and column indexes, creating one if necessary.
//
// # Updating the Table View Arrangement
//
//   - [INSTableView.BeginUpdates]: Begins a group of updates for the table view.
//   - [INSTableView.EndUpdates]: Ends the group of updates for the table view.
//   - [INSTableView.MoveRowAtIndexToIndex]: Moves the specified row to the new row location using animation.
//   - [INSTableView.InsertRowsAtIndexesWithAnimation]: Inserts the rows using the specified animation.
//   - [INSTableView.RemoveRowsAtIndexesWithAnimation]: Removes the rows using the specified animation.
//   - [INSTableView.RowForView]: Returns the index of the row for the specified view.
//   - [INSTableView.ColumnForView]: Returns the column index for the specified view.
//
// # NSView-Based Table Nib File Registration
//
//   - [INSTableView.RegisterNibForIdentifier]: Registers a NIB for the specified identifier, so that view-based table views can use it to instantiate views.
//   - [INSTableView.RegisteredNibsByIdentifier]: The dictionary of all registered nib files for view-based table view identifiers.
//
// # Target-action Behavior
//
//   - [INSTableView.DoubleAction]: The message sent to the table view’s target when the user double-clicks a cell or column header.
//   - [INSTableView.SetDoubleAction]
//   - [INSTableView.ClickedColumn]: The index of the column the user clicked.
//   - [INSTableView.ClickedRow]: The index of the row the user clicked.
//
// # Configuring Behavior
//
//   - [INSTableView.AllowsColumnReordering]: A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers.
//   - [INSTableView.SetAllowsColumnReordering]
//   - [INSTableView.AllowsColumnResizing]: A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers.
//   - [INSTableView.SetAllowsColumnResizing]
//   - [INSTableView.AllowsMultipleSelection]: A Boolean value indicating whether the table view allows the user to select more than one column or row at a time.
//   - [INSTableView.SetAllowsMultipleSelection]
//   - [INSTableView.AllowsEmptySelection]: A Boolean value indicating whether the table view allows the user to select zero columns or rows.
//   - [INSTableView.SetAllowsEmptySelection]
//   - [INSTableView.AllowsColumnSelection]: A Boolean value indicating whether the table view allows the user to select columns by clicking their headers.
//   - [INSTableView.SetAllowsColumnSelection]
//   - [INSTableView.UsesAutomaticRowHeights]: A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows.
//   - [INSTableView.SetUsesAutomaticRowHeights]
//
// # Setting Display Attributes
//
//   - [INSTableView.IntercellSpacing]: The horizontal and vertical spacing between cells.
//   - [INSTableView.SetIntercellSpacing]
//   - [INSTableView.RowHeight]: The height of each row in the table.
//   - [INSTableView.SetRowHeight]
//   - [INSTableView.BackgroundColor]: The color used to draw the background of the table.
//   - [INSTableView.SetBackgroundColor]
//   - [INSTableView.UsesAlternatingRowBackgroundColors]: A Boolean value indicating whether the table view uses alternating row colors for its background.
//   - [INSTableView.SetUsesAlternatingRowBackgroundColors]
//   - [INSTableView.Style]: The style that the table view uses.
//   - [INSTableView.SetStyle]
//   - [INSTableView.EffectiveStyle]: The effective style that the table uses.
//   - [INSTableView.SelectionHighlightStyle]: The selection highlight style used by the table view to indicate row and column selection.
//   - [INSTableView.SetSelectionHighlightStyle]
//   - [INSTableView.GridColor]: The color used to draw grid lines.
//   - [INSTableView.SetGridColor]
//   - [INSTableView.GridStyleMask]: The grid lines drawn by the table view.
//   - [INSTableView.SetGridStyleMask]
//   - [INSTableView.IndicatorImageInTableColumn]: Returns the indicator image of the specified table column.
//   - [INSTableView.SetIndicatorImageInTableColumn]: Sets the indicator image of the specified column.
//
// # Getting and Setting Row Size Styles
//
//   - [INSTableView.EffectiveRowSizeStyle]: The effective row size style for the table.
//   - [INSTableView.RowSizeStyle]: The row size style (small, medium, large, or custom) used by the table view.
//   - [INSTableView.SetRowSizeStyle]
//
// # Column Management
//
//   - [INSTableView.AddTableColumn]: Adds the specified column as the last column of the table view.
//   - [INSTableView.RemoveTableColumn]: Removes the specified column from the table view.
//   - [INSTableView.MoveColumnToColumn]: Moves the column and heading at the specified index to the new specified index.
//   - [INSTableView.TableColumns]: An array containing the current table column objects.
//   - [INSTableView.ColumnWithIdentifier]: Returns the index of the first column in the table view whose identifier is equal to the specified identifier.
//   - [INSTableView.TableColumnWithIdentifier]: Returns the [NSTableColumn] object for the first column whose identifier is equal to the specified object.
//
// # Selecting Columns and Rows
//
//   - [INSTableView.SelectColumnIndexesByExtendingSelection]: Sets the column selection using `indexes` possibly extending the selection.
//   - [INSTableView.SelectedColumn]: The index of the last selected column (or the last column added to the selection).
//   - [INSTableView.SelectedColumnIndexes]: An index set containing the indexes of the selected columns.
//   - [INSTableView.DeselectColumn]: Deselects the column at the specified index if it’s selected.
//   - [INSTableView.NumberOfSelectedColumns]: The number of selected columns.
//   - [INSTableView.IsColumnSelected]: Returns a Boolean value that indicates whether the column at the specified index is selected.
//   - [INSTableView.SelectRowIndexesByExtendingSelection]: Sets the row selection using `indexes` extending the selection if specified.
//   - [INSTableView.SelectedRow]: The index of the last selected row (or the last row added to the selection).
//   - [INSTableView.SelectedRowIndexes]: An index set containing the indexes of the selected rows.
//   - [INSTableView.DeselectRow]: Deselects the row at the specified index if it’s selected.
//   - [INSTableView.NumberOfSelectedRows]: The number of selected rows.
//   - [INSTableView.IsRowSelected]: Returns a Boolean value that indicates whether the row at the specified index is selected.
//   - [INSTableView.DeselectAll]: Deselects all selected rows or columns if empty selection is allowed; otherwise does nothing.
//
// # Enumerating Table Rows
//
//   - [INSTableView.EnumerateAvailableRowViewsUsingBlock]: Allows the enumeration of all the table rows that are known to the table view.
//
// # Managing Type Select
//
//   - [INSTableView.AllowsTypeSelect]: A Boolean value indicating whether the table view allows the user to type characters to select rows.
//   - [INSTableView.SetAllowsTypeSelect]
//
// # Table Dimensions
//
//   - [INSTableView.NumberOfColumns]: The number of columns in the table.
//   - [INSTableView.NumberOfRows]: The number of rows in the table.
//
// # Getting and Setting Floating Rows
//
//   - [INSTableView.FloatsGroupRows]: A Boolean value indicating whether the table view draws grouped rows as if they are floating.
//   - [INSTableView.SetFloatsGroupRows]
//
// # Editing Cells
//
//   - [INSTableView.EditColumnRowWithEventSelect]: Edits the cell at the specified column and row using the specified event and selection behavior.
//   - [INSTableView.EditedColumn]: The index of the column being edited.
//   - [INSTableView.EditedRow]: The index of the row being edited.
//
// # Adding and Deleting Row Views
//
//   - [INSTableView.DidAddRowViewForRow]: Invoked when a row view is added to the table.
//   - [INSTableView.DidRemoveRowViewForRow]: Invoked when a row view is removed from the table.
//
// # Setting Auxiliary Views
//
//   - [INSTableView.HeaderView]: The view object used to draw headers over columns.
//   - [INSTableView.SetHeaderView]
//   - [INSTableView.CornerView]: The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view.
//   - [INSTableView.SetCornerView]
//
// # Layout Support
//
//   - [INSTableView.RectOfColumn]: Returns the rectangle containing the column at the specified index.
//   - [INSTableView.RectOfRow]: Returns the rectangle containing the row at the specified index.
//   - [INSTableView.RowsInRect]: Returns a range of indexes for the rows that lie wholly or partially within the vertical boundaries of the specified rectangle.
//   - [INSTableView.ColumnIndexesInRect]: Returns the indexes of the table view’s columns that intersect the specified rectangle.
//   - [INSTableView.ColumnAtPoint]: Returns the index of the column the specified point lies in.
//   - [INSTableView.RowAtPoint]: Returns the index of the row the specified point lies in.
//   - [INSTableView.FrameOfCellAtColumnRow]: Returns a rectangle locating the cell that lies at the intersection of the specified column and row.
//   - [INSTableView.ColumnAutoresizingStyle]: The table view’s column autoresizing style.
//   - [INSTableView.SetColumnAutoresizingStyle]
//   - [INSTableView.SizeLastColumnToFit]: Resizes the last column so the table view fits exactly within its enclosing clip view.
//   - [INSTableView.NoteNumberOfRowsChanged]: Informs the table view that the number of records in its data source has changed.
//   - [INSTableView.Tile]: Properly sizes the table view and its header view and marks it as needing display.
//   - [INSTableView.NoteHeightOfRowsWithIndexesChanged]: Informs the table view that the rows specified in `indexSet` have changed height.
//
// # Drawing
//
//   - [INSTableView.DrawRowClipRect]: Draws the cells for the row at `rowIndex` in the columns that intersect `clipRect`.
//   - [INSTableView.DrawGridInClipRect]: Draws the grid lines within the supplied rectangle.
//   - [INSTableView.HighlightSelectionInClipRect]: Highlights the region of the table view in the specified rectangle.
//   - [INSTableView.DrawBackgroundInClipRect]: Draws the background of the table view in the clip rect specified by the rectangle.
//
// # Scrolling
//
//   - [INSTableView.ScrollRowToVisible]: Scrolls the view so the specified row is visible.
//   - [INSTableView.ScrollColumnToVisible]: Scrolls the view so the specified column is visible.
//
// # Table Column State Persistence
//
//   - [INSTableView.AutosaveTableColumns]: A Boolean value indicating whether the order and width of the table view’s columns are automatically saved.
//   - [INSTableView.SetAutosaveTableColumns]
//   - [INSTableView.AutosaveName]: The name under which table information is automatically saved.
//   - [INSTableView.SetAutosaveName]
//
// # Accessing the Delegate
//
//   - [INSTableView.Delegate]: The table view’s delegate.
//   - [INSTableView.SetDelegate]
//
// # Highlightable Column Headers
//
//   - [INSTableView.HighlightedTableColumn]: The column highlighted in the table.
//   - [INSTableView.SetHighlightedTableColumn]
//
// # Dragging
//
//   - [INSTableView.DragImageForRowsWithIndexesTableColumnsEventOffset]: Computes and returns an image to use for dragging.
//   - [INSTableView.CanDragRowsWithIndexesAtPoint]: Returns a Boolean value indicating whether the table view allows dragging the rows with the drag initiated at the specified point.
//   - [INSTableView.SetDraggingSourceOperationMaskForLocal]: Sets the default operation mask returned by `` to `mask`.
//   - [INSTableView.VerticalMotionCanBeginDrag]: A Boolean value indicating whether vertical motion is treated as a drag or selection change.
//   - [INSTableView.SetVerticalMotionCanBeginDrag]
//   - [INSTableView.DraggingDestinationFeedbackStyle]: The feedback style displayed when the user drags over the table view.
//   - [INSTableView.SetDraggingDestinationFeedbackStyle]
//   - [INSTableView.SetDropRowDropOperation]: Retargets the proposed drop operation.
//
// # Sorting
//
//   - [INSTableView.SortDescriptors]: The table view’s sort descriptors.
//   - [INSTableView.SetSortDescriptors]
//
// # Row Actions
//
//   - [INSTableView.RowActionsVisible]: A Boolean value indicating whether a table row’s actions are visible.
//   - [INSTableView.SetRowActionsVisible]
//
// # Hiding and Showing Table Rows
//
//   - [INSTableView.HideRowsAtIndexesWithAnimation]: Hides the specified table rows.
//   - [INSTableView.UnhideRowsAtIndexesWithAnimation]: Unhides the specified table rows.
//   - [INSTableView.HiddenRowIndexes]: The indexes of all hidden table rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView
type INSTableView interface {
	INSControl
	NSAccessibilityGroup
	NSAccessibilityTable
	NSDraggingSource
	NSTextDelegate
	NSTextViewDelegate
	NSUserInterfaceValidations

	// Topic: Managing the Table’s Data

	// The object that provides the data displayed by the table view.
	DataSource() NSTableViewDataSource
	SetDataSource(value NSTableViewDataSource)
	// A Boolean value indicating whether the table uses static data.
	UsesStaticContents() bool
	SetUsesStaticContents(value bool)
	// Marks the table view as needing redisplay, so it will reload the data for visible cells and draw the new values.
	ReloadData()
	// Reloads the data for only the specified rows and columns.
	ReloadDataForRowIndexesColumnIndexes(rowIndexes foundation.NSIndexSet, columnIndexes foundation.NSIndexSet)

	// Topic: Creating Views to Display

	// Returns a new or existing view with the specified identifier.
	MakeViewWithIdentifierOwner(identifier NSUserInterfaceItemIdentifier, owner objectivec.IObject) INSView
	// Returns a row view at the specified index, creating one if necessary.
	RowViewAtRowMakeIfNecessary(row int, makeIfNecessary bool) INSTableRowView
	// Returns a view at the specified row and column indexes, creating one if necessary.
	ViewAtColumnRowMakeIfNecessary(column int, row int, makeIfNecessary bool) INSView

	// Topic: Updating the Table View Arrangement

	// Begins a group of updates for the table view.
	BeginUpdates()
	// Ends the group of updates for the table view.
	EndUpdates()
	// Moves the specified row to the new row location using animation.
	MoveRowAtIndexToIndex(oldIndex int, newIndex int)
	// Inserts the rows using the specified animation.
	InsertRowsAtIndexesWithAnimation(indexes foundation.NSIndexSet, animationOptions NSTableViewAnimationOptions)
	// Removes the rows using the specified animation.
	RemoveRowsAtIndexesWithAnimation(indexes foundation.NSIndexSet, animationOptions NSTableViewAnimationOptions)
	// Returns the index of the row for the specified view.
	RowForView(view INSView) int
	// Returns the column index for the specified view.
	ColumnForView(view INSView) int

	// Topic: NSView-Based Table Nib File Registration

	// Registers a NIB for the specified identifier, so that view-based table views can use it to instantiate views.
	RegisterNibForIdentifier(nib INSNib, identifier NSUserInterfaceItemIdentifier)
	// The dictionary of all registered nib files for view-based table view identifiers.
	RegisteredNibsByIdentifier() foundation.INSDictionary

	// Topic: Target-action Behavior

	// The message sent to the table view’s target when the user double-clicks a cell or column header.
	DoubleAction() objc.SEL
	SetDoubleAction(value objc.SEL)
	// The index of the column the user clicked.
	ClickedColumn() int
	// The index of the row the user clicked.
	ClickedRow() int

	// Topic: Configuring Behavior

	// A Boolean value indicating whether the table view allows the user to rearrange columns by dragging their headers.
	AllowsColumnReordering() bool
	SetAllowsColumnReordering(value bool)
	// A Boolean value indicating whether the table view allows the user to resize columns by dragging between their headers.
	AllowsColumnResizing() bool
	SetAllowsColumnResizing(value bool)
	// A Boolean value indicating whether the table view allows the user to select more than one column or row at a time.
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	// A Boolean value indicating whether the table view allows the user to select zero columns or rows.
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	// A Boolean value indicating whether the table view allows the user to select columns by clicking their headers.
	AllowsColumnSelection() bool
	SetAllowsColumnSelection(value bool)
	// A Boolean value that indicates whether the table view uses autolayout to calculate the height of rows.
	UsesAutomaticRowHeights() bool
	SetUsesAutomaticRowHeights(value bool)

	// Topic: Setting Display Attributes

	// The horizontal and vertical spacing between cells.
	IntercellSpacing() corefoundation.CGSize
	SetIntercellSpacing(value corefoundation.CGSize)
	// The height of each row in the table.
	RowHeight() float64
	SetRowHeight(value float64)
	// The color used to draw the background of the table.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// A Boolean value indicating whether the table view uses alternating row colors for its background.
	UsesAlternatingRowBackgroundColors() bool
	SetUsesAlternatingRowBackgroundColors(value bool)
	// The style that the table view uses.
	Style() NSTableViewStyle
	SetStyle(value NSTableViewStyle)
	// The effective style that the table uses.
	EffectiveStyle() NSTableViewStyle
	// The selection highlight style used by the table view to indicate row and column selection.
	SelectionHighlightStyle() NSTableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value NSTableViewSelectionHighlightStyle)
	// The color used to draw grid lines.
	GridColor() INSColor
	SetGridColor(value INSColor)
	// The grid lines drawn by the table view.
	GridStyleMask() NSTableViewGridLineStyle
	SetGridStyleMask(value NSTableViewGridLineStyle)
	// Returns the indicator image of the specified table column.
	IndicatorImageInTableColumn(tableColumn INSTableColumn) INSImage
	// Sets the indicator image of the specified column.
	SetIndicatorImageInTableColumn(image INSImage, tableColumn INSTableColumn)

	// Topic: Getting and Setting Row Size Styles

	// The effective row size style for the table.
	EffectiveRowSizeStyle() NSTableViewRowSizeStyle
	// The row size style (small, medium, large, or custom) used by the table view.
	RowSizeStyle() NSTableViewRowSizeStyle
	SetRowSizeStyle(value NSTableViewRowSizeStyle)

	// Topic: Column Management

	// Adds the specified column as the last column of the table view.
	AddTableColumn(tableColumn INSTableColumn)
	// Removes the specified column from the table view.
	RemoveTableColumn(tableColumn INSTableColumn)
	// Moves the column and heading at the specified index to the new specified index.
	MoveColumnToColumn(oldIndex int, newIndex int)
	// An array containing the current table column objects.
	TableColumns() []NSTableColumn
	// Returns the index of the first column in the table view whose identifier is equal to the specified identifier.
	ColumnWithIdentifier(identifier NSUserInterfaceItemIdentifier) int
	// Returns the [NSTableColumn] object for the first column whose identifier is equal to the specified object.
	TableColumnWithIdentifier(identifier NSUserInterfaceItemIdentifier) INSTableColumn

	// Topic: Selecting Columns and Rows

	// Sets the column selection using `indexes` possibly extending the selection.
	SelectColumnIndexesByExtendingSelection(indexes foundation.NSIndexSet, extend bool)
	// The index of the last selected column (or the last column added to the selection).
	SelectedColumn() int
	// An index set containing the indexes of the selected columns.
	SelectedColumnIndexes() foundation.NSIndexSet
	// Deselects the column at the specified index if it’s selected.
	DeselectColumn(column int)
	// The number of selected columns.
	NumberOfSelectedColumns() int
	// Returns a Boolean value that indicates whether the column at the specified index is selected.
	IsColumnSelected(column int) bool
	// Sets the row selection using `indexes` extending the selection if specified.
	SelectRowIndexesByExtendingSelection(indexes foundation.NSIndexSet, extend bool)
	// The index of the last selected row (or the last row added to the selection).
	SelectedRow() int
	// An index set containing the indexes of the selected rows.
	SelectedRowIndexes() foundation.NSIndexSet
	// Deselects the row at the specified index if it’s selected.
	DeselectRow(row int)
	// The number of selected rows.
	NumberOfSelectedRows() int
	// Returns a Boolean value that indicates whether the row at the specified index is selected.
	IsRowSelected(row int) bool
	// Deselects all selected rows or columns if empty selection is allowed; otherwise does nothing.
	DeselectAll(sender objectivec.IObject)

	// Topic: Enumerating Table Rows

	// Allows the enumeration of all the table rows that are known to the table view.
	EnumerateAvailableRowViewsUsingBlock(handler TableRowViewHandler)

	// Topic: Managing Type Select

	// A Boolean value indicating whether the table view allows the user to type characters to select rows.
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)

	// Topic: Table Dimensions

	// The number of columns in the table.
	NumberOfColumns() int
	// The number of rows in the table.
	NumberOfRows() int

	// Topic: Getting and Setting Floating Rows

	// A Boolean value indicating whether the table view draws grouped rows as if they are floating.
	FloatsGroupRows() bool
	SetFloatsGroupRows(value bool)

	// Topic: Editing Cells

	// Edits the cell at the specified column and row using the specified event and selection behavior.
	EditColumnRowWithEventSelect(column int, row int, event INSEvent, select_ bool)
	// The index of the column being edited.
	EditedColumn() int
	// The index of the row being edited.
	EditedRow() int

	// Topic: Adding and Deleting Row Views

	// Invoked when a row view is added to the table.
	DidAddRowViewForRow(rowView INSTableRowView, row int)
	// Invoked when a row view is removed from the table.
	DidRemoveRowViewForRow(rowView INSTableRowView, row int)

	// Topic: Setting Auxiliary Views

	// The view object used to draw headers over columns.
	HeaderView() INSTableHeaderView
	SetHeaderView(value INSTableHeaderView)
	// The view used to draw the area to the right of the column headers and above the vertical scroller of the enclosing scroll view.
	CornerView() INSView
	SetCornerView(value INSView)

	// Topic: Layout Support

	// Returns the rectangle containing the column at the specified index.
	RectOfColumn(column int) corefoundation.CGRect
	// Returns the rectangle containing the row at the specified index.
	RectOfRow(row int) corefoundation.CGRect
	// Returns a range of indexes for the rows that lie wholly or partially within the vertical boundaries of the specified rectangle.
	RowsInRect(rect corefoundation.CGRect) foundation.NSRange
	// Returns the indexes of the table view’s columns that intersect the specified rectangle.
	ColumnIndexesInRect(rect corefoundation.CGRect) foundation.NSIndexSet
	// Returns the index of the column the specified point lies in.
	ColumnAtPoint(point corefoundation.CGPoint) int
	// Returns the index of the row the specified point lies in.
	RowAtPoint(point corefoundation.CGPoint) int
	// Returns a rectangle locating the cell that lies at the intersection of the specified column and row.
	FrameOfCellAtColumnRow(column int, row int) corefoundation.CGRect
	// The table view’s column autoresizing style.
	ColumnAutoresizingStyle() NSTableViewColumnAutoresizingStyle
	SetColumnAutoresizingStyle(value NSTableViewColumnAutoresizingStyle)
	// Resizes the last column so the table view fits exactly within its enclosing clip view.
	SizeLastColumnToFit()
	// Informs the table view that the number of records in its data source has changed.
	NoteNumberOfRowsChanged()
	// Properly sizes the table view and its header view and marks it as needing display.
	Tile()
	// Informs the table view that the rows specified in `indexSet` have changed height.
	NoteHeightOfRowsWithIndexesChanged(indexSet foundation.NSIndexSet)

	// Topic: Drawing

	// Draws the cells for the row at `rowIndex` in the columns that intersect `clipRect`.
	DrawRowClipRect(row int, clipRect corefoundation.CGRect)
	// Draws the grid lines within the supplied rectangle.
	DrawGridInClipRect(clipRect corefoundation.CGRect)
	// Highlights the region of the table view in the specified rectangle.
	HighlightSelectionInClipRect(clipRect corefoundation.CGRect)
	// Draws the background of the table view in the clip rect specified by the rectangle.
	DrawBackgroundInClipRect(clipRect corefoundation.CGRect)

	// Topic: Scrolling

	// Scrolls the view so the specified row is visible.
	ScrollRowToVisible(row int)
	// Scrolls the view so the specified column is visible.
	ScrollColumnToVisible(column int)

	// Topic: Table Column State Persistence

	// A Boolean value indicating whether the order and width of the table view’s columns are automatically saved.
	AutosaveTableColumns() bool
	SetAutosaveTableColumns(value bool)
	// The name under which table information is automatically saved.
	AutosaveName() NSTableViewAutosaveName
	SetAutosaveName(value NSTableViewAutosaveName)

	// Topic: Accessing the Delegate

	// The table view’s delegate.
	Delegate() NSTableViewDelegate
	SetDelegate(value NSTableViewDelegate)

	// Topic: Highlightable Column Headers

	// The column highlighted in the table.
	HighlightedTableColumn() INSTableColumn
	SetHighlightedTableColumn(value INSTableColumn)

	// Topic: Dragging

	// Computes and returns an image to use for dragging.
	DragImageForRowsWithIndexesTableColumnsEventOffset(dragRows foundation.NSIndexSet, tableColumns []NSTableColumn, dragEvent INSEvent, dragImageOffset foundation.NSPoint) INSImage
	// Returns a Boolean value indicating whether the table view allows dragging the rows with the drag initiated at the specified point.
	CanDragRowsWithIndexesAtPoint(rowIndexes foundation.NSIndexSet, mouseDownPoint corefoundation.CGPoint) bool
	// Sets the default operation mask returned by `` to `mask`.
	SetDraggingSourceOperationMaskForLocal(mask NSDragOperation, isLocal bool)
	// A Boolean value indicating whether vertical motion is treated as a drag or selection change.
	VerticalMotionCanBeginDrag() bool
	SetVerticalMotionCanBeginDrag(value bool)
	// The feedback style displayed when the user drags over the table view.
	DraggingDestinationFeedbackStyle() NSTableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value NSTableViewDraggingDestinationFeedbackStyle)
	// Retargets the proposed drop operation.
	SetDropRowDropOperation(row int, dropOperation NSTableViewDropOperation)

	// Topic: Sorting

	// The table view’s sort descriptors.
	SortDescriptors() []foundation.NSSortDescriptor
	SetSortDescriptors(value []foundation.NSSortDescriptor)

	// Topic: Row Actions

	// A Boolean value indicating whether a table row’s actions are visible.
	RowActionsVisible() bool
	SetRowActionsVisible(value bool)

	// Topic: Hiding and Showing Table Rows

	// Hides the specified table rows.
	HideRowsAtIndexesWithAnimation(indexes foundation.NSIndexSet, rowAnimation NSTableViewAnimationOptions)
	// Unhides the specified table rows.
	UnhideRowsAtIndexesWithAnimation(indexes foundation.NSIndexSet, rowAnimation NSTableViewAnimationOptions)
	// The indexes of all hidden table rows.
	HiddenRowIndexes() foundation.NSIndexSet

	// A Boolean value that indicates whether the receiver reacts to mouse events.
	IsEnabled() bool
	SetIsEnabled(value bool)
}

// Init initializes the instance.
func (t NSTableView) Init() NSTableView {
	rv := objc.Send[NSTableView](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTableView) Autorelease() NSTableView {
	rv := objc.Send[NSTableView](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTableView creates a new NSTableView instance.
func NewNSTableView() NSTableView {
	class := getNSTableViewClass()
	rv := objc.Send[NSTableView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/init(coder:)
func NewTableViewWithCoder(coder foundation.INSCoder) NSTableView {
	instance := getNSTableViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTableViewFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/init(frame:)
func NewTableViewWithFrame(frameRect corefoundation.CGRect) NSTableView {
	instance := getNSTableViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSTableViewFromID(rv)
}

// Marks the table view as needing redisplay, so it will reload the data for
// visible cells and draw the new values.
//
// # Discussion
// 
// This method forces a redraw of all the visible cells in the table view. If
// you want to update the value in a single cell, column, or row, it is more
// efficient to use [FrameOfCellAtColumnRow], [RectOfColumn], or [RectOfRow]
// in conjunction with the [SetNeedsDisplayInRect] method of [NSView]. If you
// just want to update the scroller, use [NoteNumberOfRowsChanged]; if the
// height of a set of rows changes, use [NoteHeightOfRowsWithIndexesChanged].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/reloadData()
func (t NSTableView) ReloadData() {
	objc.Send[objc.ID](t.ID, objc.Sel("reloadData"))
}
// Reloads the data for only the specified rows and columns.
//
// rowIndexes: The indexes of the rows to update.
//
// columnIndexes: The indexes of the columns to update.
//
// # Discussion
// 
// For cells that are visible, the appropriate [DataSource] and [Delegate]
// methods are called and the cells are redrawn.
// 
// For tables that support variable row heights, the row height is not
// re-queried from the delegate; it is your responsibility to invoke
// [NoteHeightOfRowsWithIndexesChanged] if a row height change is required.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/reloadData(forRowIndexes:columnIndexes:)
func (t NSTableView) ReloadDataForRowIndexesColumnIndexes(rowIndexes foundation.NSIndexSet, columnIndexes foundation.NSIndexSet) {
	objc.Send[objc.ID](t.ID, objc.Sel("reloadDataForRowIndexes:columnIndexes:"), rowIndexes, columnIndexes)
}
// Returns a new or existing view with the specified identifier.
//
// identifier: The view identifier. Must not be `nil`.
//
// owner: The owner of the NIB that may be loaded and instantiated to create a new
// view with the specified identifier.
//
// # Return Value
// 
// A view for the row.
//
// # Discussion
// 
// Typically, `identifier` is associated with a cell view that’s contained
// in a table’s [Nib file]. When this method is called, the table view
// automatically instantiates the cell view with the specified owner, which is
// usually the table view’s delegate. (The owner is useful in setting up
// outlets and target/actions from the view.) Note that a cell view’s
// identifier must be the same as its table column’s identifier for bindings
// to work. If you’re using bindings, it’s recommended that you use the
// Automatic identifier setting in Interface Builder.
// 
// This method may also return a reused view with the same `identifier` that
// is no longer available on screen. If a view with the specified identifier
// can’t be instantiated from the nib file or found in the reuse queue, this
// method returns `nil`.
// 
// This method is usually called by the delegate in
// [TableViewViewForTableColumnRow], but it can also be overridden to provide
// custom views for the `identifier`. Note that [awakeFromNib()] is called
// each time this method is called, which means that `awakeFromNib` is also
// called on `owner`, even though the owner is already awake.
//
// [Nib file]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/NibFile.html#//apple_ref/doc/uid/TP40008195-CH34
// [awakeFromNib()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/awakeFromNib()
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/makeView(withIdentifier:owner:)
func (t NSTableView) MakeViewWithIdentifierOwner(identifier NSUserInterfaceItemIdentifier, owner objectivec.IObject) INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("makeViewWithIdentifier:owner:"), objc.String(string(identifier)), owner)
	return NSViewFromID(rv)
}
// Returns a row view at the specified index, creating one if necessary.
//
// row: The row index.
//
// makeIfNecessary: [true] if a view is required, [false] if you want to update properties on a
// view, if one is available.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An instance, or subclass, of [NSTableRowView]. Returning `nil` is also
// valid if `makeIfNecessary` is [false] and the view did not exist.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Discussion
// 
// This method first attempts to return a currently displayed view in the
// visible area. If there is no visible view, and `makeIfNecessary` is [true],
// a prepared temporary view is returned. If `makeIfNecessary` is [false], and
// the view is not visible, `nil` is returned.
// 
// In general, `makeIfNecessary` should be [true] if you require a resulting
// view, and [false] if you want to update properties on a view only if it is
// available (generally this means it is visible).
// 
// An exception is thrown if `row` falls outside of the number of rows in the
// table ([NumberOfRows]). The returned result should generally not be held
// onto for longer than the current run loop cycle. It’s better to call
// [RowViewAtRowMakeIfNecessary] whenever a view is required.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/rowView(atRow:makeIfNecessary:)
func (t NSTableView) RowViewAtRowMakeIfNecessary(row int, makeIfNecessary bool) INSTableRowView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("rowViewAtRow:makeIfNecessary:"), row, makeIfNecessary)
	return NSTableRowViewFromID(rv)
}
// Returns a view at the specified row and column indexes, creating one if
// necessary.
//
// column: The index of the column in the [TableColumns] array.
//
// row: The row index.
//
// makeIfNecessary: [true] if a view is required, [false] if you want to update properties on a
// view, if one is available.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An instance of [NSView].
//
// # Discussion
// 
// This method first attempts to return an available view, which is generally
// in the visible area. If there is no available view, and `makeIfNecessary`
// is [true], a prepared temporary view is returned. If `makeIfNecessary` is
// [false], and the view is not available, `nil` will be returned.
// 
// In general, `makeIfNecessary` should be [true] if you require a resulting
// view, and [false] if you only want to update properties on a view only if
// it is available (generally this means it is visible).
// 
// An exception will be thrown if `row` is not within the [NumberOfRows]. The
// returned result should generally not be held onto for longer than the
// current run loop cycle. Instead they should re-query the table view for the
// row view.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/view(atColumn:row:makeIfNecessary:)
func (t NSTableView) ViewAtColumnRowMakeIfNecessary(column int, row int, makeIfNecessary bool) INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("viewAtColumn:row:makeIfNecessary:"), column, row, makeIfNecessary)
	return NSViewFromID(rv)
}
// Begins a group of updates for the table view.
//
// # Discussion
// 
// For [NSView]-based table views, multiple row changes—that is, insertions,
// deletions, and moves—are animated simultaneously by surrounding calls to
// those method calls with [BeginUpdates] and [EndUpdates]. These methods are
// nestable.
// 
// The selected rows are maintained during the series of insertions,
// deletions, moves, and scrolling. If a selected row is deleted, a selection
// changed notification occurs after [RemoveRowsAtIndexesWithAnimation] is
// called.
// 
// It is not necessary to call [BeginUpdates] and [EndUpdates] if only one
// insertion, deletion, or move is occurring and the table view is an
// [NSView]-based table view. When using an [NSCell]-based table view, you
// must surround any insertion, deletion, or move in an update block for
// animations to occur.
// 
// The main reason for doing a batch update of changes to a table view is to
// avoid having the table animate unnecessarily.
// 
// Note that these methods should be called to reflect changes in your model;
// they do not make any underlying model changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/beginUpdates()
func (t NSTableView) BeginUpdates() {
	objc.Send[objc.ID](t.ID, objc.Sel("beginUpdates"))
}
// Ends the group of updates for the table view.
//
// # Discussion
// 
// Ends the group of updates for the table view. This method, like
// [BeginUpdates], is nestable. See [BeginUpdates] for details.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/endUpdates()
func (t NSTableView) EndUpdates() {
	objc.Send[objc.ID](t.ID, objc.Sel("endUpdates"))
}
// Moves the specified row to the new row location using animation.
//
// oldIndex: Initial row index.
//
// newIndex: New row index.
//
// # Discussion
// 
// This is similar to removing a row at `oldIndex` and inserting it at
// `newIndex`, except the same view is used and simply has its position
// updated to the new location.
// 
// Changes happen incrementally as they are sent to the table, so as soon as
// this method is called the row can be considered moved. However the
// underlying view is not moved until [EndUpdates] has been called.
// 
// This method can be called multiple times within the same [BeginUpdates] and
// [EndUpdates] block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/moveRow(at:to:)
func (t NSTableView) MoveRowAtIndexToIndex(oldIndex int, newIndex int) {
	objc.Send[objc.ID](t.ID, objc.Sel("moveRowAtIndex:toIndex:"), oldIndex, newIndex)
}
// Inserts the rows using the specified animation.
//
// indexes: The final positions of the new rows to be inserted.
//
// animationOptions: The animation displayed during the insert. See
// [NSTableView.AnimationOptions] for the possible values that can be combined
// using the C bitwise OR operator.
// //
// [NSTableView.AnimationOptions]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions
//
// # Discussion
// 
// The [NumberOfRows] in the table view is automatically increased by the
// count of `indexes`.
// 
// Calling this method multiple times within the same [BeginUpdates] and
// [EndUpdates] block is allowed, and changes are processed incrementally.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/insertRows(at:withAnimation:)
func (t NSTableView) InsertRowsAtIndexesWithAnimation(indexes foundation.NSIndexSet, animationOptions NSTableViewAnimationOptions) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertRowsAtIndexes:withAnimation:"), indexes, animationOptions)
}
// Removes the rows using the specified animation.
//
// indexes: An index set containing the rows to remove.
//
// animationOptions: The animation displayed during the insert. See
// [NSTableView.AnimationOptions] for the possible values that can be combined
// using the C bitwise OR operator.
// //
// [NSTableView.AnimationOptions]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions
//
// # Discussion
// 
// This method deletes from the table the rows represented at `indexes` and
// automatically decreases [NumberOfRows] by the count of `indexes`.
// 
// The row indexes should be with respect to the current state displayed in
// the table view, and not the final state, because the specified rows do not
// exist in the final state.
// 
// Calling this method multiple times within the same [BeginUpdates] and
// [EndUpdates] block is allowed, and changes are processed incrementally.
// 
// Changes are processed incrementally as the
// [InsertRowsAtIndexesWithAnimation], [RemoveRowsAtIndexesWithAnimation], and
// the [MoveRowAtIndexToIndex] methods are called. It is acceptable to delete
// row `0` multiple times, as long as there is still a row available.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/removeRows(at:withAnimation:)
func (t NSTableView) RemoveRowsAtIndexesWithAnimation(indexes foundation.NSIndexSet, animationOptions NSTableViewAnimationOptions) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeRowsAtIndexes:withAnimation:"), indexes, animationOptions)
}
// Returns the index of the row for the specified view.
//
// view: The view for which to retrieve the row.
//
// # Return Value
// 
// The index of the row containing to `view`. This method returns `-1` if the
// view is not in the table view. This method may also return `-1` if the row
// containing the view is being animated away, such as during the deletion of
// a row.
//
// # Discussion
// 
// This method is typically called in the action method for an [NSButton] (or
// [NSControl]) to find out what row (and column) the action should be
// performed on.
// 
// The implementation is `O(n)` where is the number of visible rows, so this
// method should generally not be called within a loop.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/row(for:)
func (t NSTableView) RowForView(view INSView) int {
	rv := objc.Send[int](t.ID, objc.Sel("rowForView:"), view)
	return rv
}
// Returns the column index for the specified view.
//
// view: The view for which to retrieve the column.
//
// # Return Value
// 
// The index of the column containing `view` in the [TableColumns] array. This
// method returns `-1` if the view is not in the table view. This method may
// also return `-1` if the row containing the view is being animated away,
// such as during the deletion of a row.
//
// # Discussion
// 
// This method is typically called in the action method of an [NSButton] (or
// [NSControl]) to find out what row (and column) the action should be
// performed on.
// 
// The implementation is `O(n)` where is the number of visible rows, so this
// method should generally not be called within a loop.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/column(for:)
func (t NSTableView) ColumnForView(view INSView) int {
	rv := objc.Send[int](t.ID, objc.Sel("columnForView:"), view)
	return rv
}
// Registers a NIB for the specified identifier, so that view-based table
// views can use it to instantiate views.
//
// nib: The nib containing the view.
//
// identifier: The identifier of the view to create.
//
// # Discussion
// 
// Use this method to associate one of the NIB’s cell views with
// `identifier` so that the table can instantiate this view when requested.
// This method is used when [ViewWithIdentifierOwner] is called, and there was
// no NIB created at design time for the specified identifier. This allows
// dynamic loading of NIBs that can be associated with the table.
// 
// Because a NIB can contain multiple views, you can associate the same NIB
// with multiple identifiers. To remove a previously associated NIB for
// `identifier`, pass in `nil` for the `nib` value.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/register(_:forIdentifier:)
func (t NSTableView) RegisterNibForIdentifier(nib INSNib, identifier NSUserInterfaceItemIdentifier) {
	objc.Send[objc.ID](t.ID, objc.Sel("registerNib:forIdentifier:"), nib, objc.String(string(identifier)))
}
// Returns the indicator image of the specified table column.
//
// tableColumn: A table column in the table view.
//
// # Discussion
// 
// An indicator image is an arbitrary (small) image that is rendered on the
// right side of the column header. An example of its use is in Mail to
// indicate the sorting direction of the currently sorted column in a mailbox.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/indicatorImage(in:)
func (t NSTableView) IndicatorImageInTableColumn(tableColumn INSTableColumn) INSImage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("indicatorImageInTableColumn:"), tableColumn)
	return NSImageFromID(rv)
}
// Sets the indicator image of the specified column.
//
// image: The indicator image for the column.
//
// tableColumn: The table column.
//
// # Discussion
// 
// The default sorting order indicators are available as named [NSImage]
// objects. These images are accessed using `[NSImage ]` passing either
// `@"NSAscendingSortIndicator"` (the “^” icon), and
// `@"NSDescendingSortIndicator"` (the “v” icon).
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/setIndicatorImage(_:in:)
func (t NSTableView) SetIndicatorImageInTableColumn(image INSImage, tableColumn INSTableColumn) {
	objc.Send[objc.ID](t.ID, objc.Sel("setIndicatorImage:inTableColumn:"), image, tableColumn)
}
// Adds the specified column as the last column of the table view.
//
// tableColumn: The column to add to the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/addTableColumn(_:)
func (t NSTableView) AddTableColumn(tableColumn INSTableColumn) {
	objc.Send[objc.ID](t.ID, objc.Sel("addTableColumn:"), tableColumn)
}
// Removes the specified column from the table view.
//
// tableColumn: The column to remove from the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/removeTableColumn(_:)
func (t NSTableView) RemoveTableColumn(tableColumn INSTableColumn) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeTableColumn:"), tableColumn)
}
// Moves the column and heading at the specified index to the new specified
// index.
//
// oldIndex: The current index in the [TableColumns] array of the column to move.
//
// newIndex: The new index in the [TableColumns] array for the moved column.
//
// # Discussion
// 
// This method posts [columnDidMoveNotification] to the default notification
// center.
//
// [columnDidMoveNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/columnDidMoveNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/moveColumn(_:toColumn:)
func (t NSTableView) MoveColumnToColumn(oldIndex int, newIndex int) {
	objc.Send[objc.ID](t.ID, objc.Sel("moveColumn:toColumn:"), oldIndex, newIndex)
}
// Returns the index of the first column in the table view whose identifier is
// equal to the specified identifier.
//
// identifier: A column identifier.
//
// # Return Value
// 
// The index in the [TableColumns] array of the first column in the table view
// whose identifier is equal to `anObject` (when compared using ``), or `–1`
// if no columns are found with the specified identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/column(withIdentifier:)
func (t NSTableView) ColumnWithIdentifier(identifier NSUserInterfaceItemIdentifier) int {
	rv := objc.Send[int](t.ID, objc.Sel("columnWithIdentifier:"), objc.String(string(identifier)))
	return rv
}
// Returns the [NSTableColumn] object for the first column whose identifier is
// equal to the specified object.
//
// identifier: A column identifier.
//
// # Return Value
// 
// The [NSTableColumn] object for the first column whose identifier is equal
// to `anObject` (when compared using ``), or `nil` if no columns are found
// with the specified identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/tableColumn(withIdentifier:)
func (t NSTableView) TableColumnWithIdentifier(identifier NSUserInterfaceItemIdentifier) INSTableColumn {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tableColumnWithIdentifier:"), objc.String(string(identifier)))
	return NSTableColumnFromID(rv)
}
// Sets the column selection using `indexes` possibly extending the selection.
//
// indexes: The column indexes to select.
//
// extend: [true] if the selection should be extended, [false] if the current
// selection should be changed.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Replaces the deprecated [selectColumn:byExtendingSelection:] method.
//
// [selectColumn:byExtendingSelection:]: https://developer.apple.com/documentation/AppKit/NSTableView/selectColumn:byExtendingSelection:
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/selectColumnIndexes(_:byExtendingSelection:)
func (t NSTableView) SelectColumnIndexesByExtendingSelection(indexes foundation.NSIndexSet, extend bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("selectColumnIndexes:byExtendingSelection:"), indexes, extend)
}
// Deselects the column at the specified index if it’s selected.
//
// column: The index in the [TableColumns] array of the column to deselect.
//
// # Discussion
// 
// Deselects the column at `columnIndex` if it’s selected, regardless of
// whether empty selection is allowed.
// 
// If the selection does in fact change, this method posts
// [selectionDidChangeNotification] to the default notification center.
// 
// If the indicated column was the last column selected by the user, the
// column nearest it effectively becomes the last selected column. In case of
// a tie, priority is given to the column on the left.
// 
// This method doesn’t check with the delegate before changing the
// selection.
//
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/deselectColumn(_:)
func (t NSTableView) DeselectColumn(column int) {
	objc.Send[objc.ID](t.ID, objc.Sel("deselectColumn:"), column)
}
// Returns a Boolean value that indicates whether the column at the specified
// index is selected.
//
// column: The index into the [TableColumns] array that represents the column to test.
//
// # Return Value
// 
// [true] if the column at `columnIndex` is selected, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/isColumnSelected(_:)
func (t NSTableView) IsColumnSelected(column int) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isColumnSelected:"), column)
	return rv
}
// Sets the row selection using `indexes` extending the selection if
// specified.
//
// indexes: The indexes to select.
//
// extend: [true] if the selection should be extended, [false] if the current
// selection should be changed.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Replaces the deprecated [selectRow:byExtendingSelection:] method.
//
// [selectRow:byExtendingSelection:]: https://developer.apple.com/documentation/AppKit/NSTableView/selectRow:byExtendingSelection:
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/selectRowIndexes(_:byExtendingSelection:)
func (t NSTableView) SelectRowIndexesByExtendingSelection(indexes foundation.NSIndexSet, extend bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("selectRowIndexes:byExtendingSelection:"), indexes, extend)
}
// Deselects the row at the specified index if it’s selected.
//
// row: The index of the row to deselect.
//
// # Discussion
// 
// Deselects the row at `rowIndex` if it’s selected, regardless of whether
// empty selection is allowed.
// 
// If the selection does in fact change, posts
// [selectionDidChangeNotification] to the default notification center.
// 
// If the indicated row was the last row selected by the user, the row nearest
// it effectively becomes the last selected row. In case of a tie, priority is
// given to the row above.
// 
// This method doesn’t check with the delegate before changing the
// selection.
//
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/deselectRow(_:)
func (t NSTableView) DeselectRow(row int) {
	objc.Send[objc.ID](t.ID, objc.Sel("deselectRow:"), row)
}
// Returns a Boolean value that indicates whether the row at the specified
// index is selected.
//
// row: The index of the row to test.
//
// # Return Value
// 
// [true] if the row at `rowIndex` is selected, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/isRowSelected(_:)
func (t NSTableView) IsRowSelected(row int) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isRowSelected:"), row)
	return rv
}
// Deselects all selected rows or columns if empty selection is allowed;
// otherwise does nothing.
//
// sender: Typically the object that sent the message.
//
// # Discussion
// 
// Posts [selectionDidChangeNotification] to the default notification center
// if the selection does in fact change.
// 
// As a target-action method, [DeselectAll] checks with the delegate before
// changing the selection, using [SelectionShouldChangeInTableView].
//
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/deselectAll(_:)
func (t NSTableView) DeselectAll(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("deselectAll:"), sender)
}
// Allows the enumeration of all the table rows that are known to the table
// view.
//
// handler: The [Block] to apply to elements in the set.
// 
// The [Block] takes two arguments:
// 
// rowView: The view for the row.
// row: The index of the row.
//
// # Discussion
// 
// The enumeration includes all views in the [visibleRect]; however, it may
// also include ones that are “in flight” due to animations or other
// attributes of the table.
// 
// It is preferred to use this method to efficiently make changes over all
// views that exist in the table.
//
// [visibleRect]: https://developer.apple.com/documentation/AppKit/NSView/visibleRect
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/enumerateAvailableRowViews(_:)
func (t NSTableView) EnumerateAvailableRowViewsUsingBlock(handler TableRowViewHandler) {
_block0, _ := NewTableRowViewBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("enumerateAvailableRowViewsUsingBlock:"), _block0)
}
// Edits the cell at the specified column and row using the specified event
// and selection behavior.
//
// column: The index of the column in the [TableColumns] array.
//
// row: The row index.
//
// event: The event.
//
// select: [true] if the entered contents should be selected, otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is invoked automatically in response to user actions; you
// should rarely need to invoke it directly. `theEvent` is usually the mouse
// event that triggered editing; it can be `nil` when starting an edit
// programmatically.
// 
// This method scrolls the table view so that the cell is visible and sets up
// the field editor. If `flag` is [false], it calls the
// [EditWithFrameInViewEditorDelegateEvent] method of the field editor’s
// [NSCell] object, providing the [NSTableView] as the text delegate. If
// `flag` is [true], this method calls the
// [SelectWithFrameInViewEditorDelegateStartLength] method instead.
// 
// This method can be overridden to customize drawing for `rowIndex` when
// using [NSCell]-based table views.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/editColumn(_:row:with:select:)
func (t NSTableView) EditColumnRowWithEventSelect(column int, row int, event INSEvent, select_ bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("editColumn:row:withEvent:select:"), column, row, event, select_)
}
// Invoked when a row view is added to the table.
//
// rowView: The row view.
//
// row: The row index.
//
// # Discussion
// 
// The subclass can implement this method to be alerted when `rowView` has
// been added to the table. At this point, the subclass can choose to add in
// extra views, or modify any properties of `rowView`. Subclasses must be sure
// to call super.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/didAdd(_:forRow:)
func (t NSTableView) DidAddRowViewForRow(rowView INSTableRowView, row int) {
	objc.Send[objc.ID](t.ID, objc.Sel("didAddRowView:forRow:"), rowView, row)
}
// Invoked when a row view is removed from the table.
//
// rowView: The row view.
//
// row: The row index. The index is `-1` for rows that are being deleted from the
// table, and no longer have a valid row; otherwise it is the valid row that
// is being removed due to it being moved off screen.
//
// # Discussion
// 
// The subclass can implement this method to be alerted when `rowView` has
// been removed from the table. The removed `rowView` may be reused by the
// table, so any additionally inserted views should be removed at this point.
// Subclasses must be sure to call `super`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/didRemove(_:forRow:)
func (t NSTableView) DidRemoveRowViewForRow(rowView INSTableRowView, row int) {
	objc.Send[objc.ID](t.ID, objc.Sel("didRemoveRowView:forRow:"), rowView, row)
}
// Returns the rectangle containing the column at the specified index.
//
// column: The index in the [TableColumns] array of a column in the table view.
//
// # Return Value
// 
// The rectangle containing the column at `columnIndex`. Returns [NSZeroRect]
// if `columnIndex` lies outside the range of valid column indexes for the
// table view.
//
// # Discussion
// 
// You can use this method to update a single column more efficiently than
// sending the table view a [ReloadData] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/rect(ofColumn:)
func (t NSTableView) RectOfColumn(column int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("rectOfColumn:"), column)
	return corefoundation.CGRect(rv)
}
// Returns the rectangle containing the row at the specified index.
//
// # Return Value
// 
// The rectangle containing the row at `rowIndex`. Returns [NSZeroRect] if
// `rowIndex` lies outside the range of valid row indexes for the table view.
//
// # Discussion
// 
// You can use this method to update a single row more efficiently than
// sending the table view a [ReloadData] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/rect(ofRow:)
func (t NSTableView) RectOfRow(row int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("rectOfRow:"), row)
	return corefoundation.CGRect(rv)
}
// Returns a range of indexes for the rows that lie wholly or partially within
// the vertical boundaries of the specified rectangle.
//
// rect: A rectangle in the coordinate system of the table view.
//
// # Return Value
// 
// A range of indexes for the table view’s rows that lie wholly or partially
// within the horizontal boundaries of `aRect`. If the width or height of
// `aRect` is `0`, this method returns an [NSRange] whose length is `0`.
//
// # Discussion
// 
// The location of the range is the index of the first row in the rectangle,
// and the length is the number of rows that lie in the rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/rows(in:)
func (t NSTableView) RowsInRect(rect corefoundation.CGRect) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("rowsInRect:"), rect)
	return foundation.NSRange(rv)
}
// Returns the indexes of the table view’s columns that intersect the
// specified rectangle.
//
// rect: The rectangle in the table view’s coordinate system to test for column
// enclosure.
//
// # Return Value
// 
// New [NSIndexSet] object containing the indexes of the table view’s
// columns that intersect with `rect`.
//
// [NSIndexSet]: https://developer.apple.com/documentation/Foundation/NSIndexSet
//
// # Discussion
// 
// Columns that return [true] for the [NSTableColumn] method [Hidden] are
// excluded from the results.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/columnIndexes(in:)
func (t NSTableView) ColumnIndexesInRect(rect corefoundation.CGRect) foundation.NSIndexSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("columnIndexesInRect:"), rect)
	return foundation.NSIndexSetFromID(rv)
}
// Returns the index of the column the specified point lies in.
//
// point: A point in the coordinate system of the table view.
//
// # Return Value
// 
// The index in the [TableColumns] array of the column `aPoint` lies in, or
// `–1` if `aPoint` lies outside the table view’s bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/column(at:)
func (t NSTableView) ColumnAtPoint(point corefoundation.CGPoint) int {
	rv := objc.Send[int](t.ID, objc.Sel("columnAtPoint:"), point)
	return rv
}
// Returns the index of the row the specified point lies in.
//
// point: A point in the coordinate system of the table view.
//
// # Return Value
// 
// The index of the row `aPoint` lies in, or `–1` if `aPoint` lies outside
// the table view’s bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/row(at:)
func (t NSTableView) RowAtPoint(point corefoundation.CGPoint) int {
	rv := objc.Send[int](t.ID, objc.Sel("rowAtPoint:"), point)
	return rv
}
// Returns a rectangle locating the cell that lies at the intersection of the
// specified column and row.
//
// column: The index in the [TableColumns] array of the column containing the cell
// whose rectangle you want.
//
// row: The index of the row containing the cell whose rectangle you want.
//
// # Return Value
// 
// A rectangle locating the cell that lies at the intersection of
// `columnIndex` and `rowIndex`. This method returns [NSZeroRect] if
// `columnIndex` or `rowIndex` is greater than the number of columns or rows
// in the table view.
//
// # Discussion
// 
// You can use this method to update a single cell more efficiently than
// sending the table view a [ReloadData] message using
// [ReloadDataForRowIndexesColumnIndexes]
// 
// The result of this method is used in a [DrawWithFrameInView] message to the
// table column’s data cell. You can subclass and override this method to
// customize the frame of a particular cell. However, never return a frame
// larger than the default implementation returns.
// 
// The default frame is computed to have a height equal to the [RectOfRow] for
// `rowIndex`, minus the half [IntercellSpacing] height on the top and half on
// the bottom. The width of frame is equal to the with of the table column
// minus half the [IntercellSpacing] width on the left, and half on the right.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/frameOfCell(atColumn:row:)
func (t NSTableView) FrameOfCellAtColumnRow(column int, row int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("frameOfCellAtColumn:row:"), column, row)
	return corefoundation.CGRect(rv)
}
// Resizes the last column so the table view fits exactly within its enclosing
// clip view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/sizeLastColumnToFit()
func (t NSTableView) SizeLastColumnToFit() {
	objc.Send[objc.ID](t.ID, objc.Sel("sizeLastColumnToFit"))
}
// Informs the table view that the number of records in its data source has
// changed.
//
// # Discussion
// 
// This method allows the table view to update the scrollers in its scroll
// view without actually reloading data into the table view. It’s useful for
// a data source that continually receives data in the background over a
// period of time, in which case the table view can remain responsive to the
// user while the data is received.
// 
// See the [NSTableViewDataSource] protocol specification for information on
// the messages an [NSTableView] object sends to its data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/noteNumberOfRowsChanged()
func (t NSTableView) NoteNumberOfRowsChanged() {
	objc.Send[objc.ID](t.ID, objc.Sel("noteNumberOfRowsChanged"))
}
// Properly sizes the table view and its header view and marks it as needing
// display.
//
// # Discussion
// 
// Also resets cursor rectangles for the header view and line scroll amounts
// for the [NSScrollView] object.
// 
// For performance reasons, calling this method is generally not recommended.
// Instead, the table will call it automatically when necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/tile()
func (t NSTableView) Tile() {
	objc.Send[objc.ID](t.ID, objc.Sel("tile"))
}
// Informs the table view that the rows specified in `indexSet` have changed
// height.
//
// indexSet: Index set of rows that have changed their height.
//
// # Discussion
// 
// If the delegate implements [TableViewHeightOfRow] this method immediately
// retiles the table view using the row heights the delegate provides.
// 
// For [NSView]-based tables, this method will animate. To turn off the
// animation, create an [NSAnimationContext] grouping and set the [Duration]
// to 0. Then call this method and end the grouping.
// 
// For [NSCell]-based tables, this method normally doesn’t animate. However,
// it will animate if you call it inside a [BeginUpdates]/[EndUpdates] block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/noteHeightOfRows(withIndexesChanged:)
func (t NSTableView) NoteHeightOfRowsWithIndexesChanged(indexSet foundation.NSIndexSet) {
	objc.Send[objc.ID](t.ID, objc.Sel("noteHeightOfRowsWithIndexesChanged:"), indexSet)
}
// Draws the cells for the row at `rowIndex` in the columns that intersect
// `clipRect`.
//
// row: The row index.
//
// clipRect: The intersecting rectangle.
//
// # Discussion
// 
// [NSCell]-based table views can override this method to customize the
// drawing of the rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/drawRow(_:clipRect:)
func (t NSTableView) DrawRowClipRect(row int, clipRect corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawRow:clipRect:"), row, clipRect)
}
// Draws the grid lines within the supplied rectangle.
//
// clipRect: The rectangle in the table view’s coordinate system.
//
// # Discussion
// 
// Draws the grid lines within `clipRect`, using the grid color set with
// [GridColor].
// 
// Subclasses can override this method to draw grid lines other than the
// standard ones. This method draws a grid regardless of whether the table
// view is set to draw one automatically.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/drawGrid(inClipRect:)
func (t NSTableView) DrawGridInClipRect(clipRect corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawGridInClipRect:"), clipRect)
}
// Highlights the region of the table view in the specified rectangle.
//
// clipRect: The rectangle, in the table view view’s coordinate system.
//
// # Discussion
// 
// This method is invoked before [DrawRowClipRect].
// 
// [NSCell]-based table views can override this method to change the manner in
// which they highlight selections.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/highlightSelection(inClipRect:)
func (t NSTableView) HighlightSelectionInClipRect(clipRect corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("highlightSelectionInClipRect:"), clipRect)
}
// Draws the background of the table view in the clip rect specified by the
// rectangle.
//
// clipRect: The rectangle, in the table view’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/drawBackground(inClipRect:)
func (t NSTableView) DrawBackgroundInClipRect(clipRect corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawBackgroundInClipRect:"), clipRect)
}
// Scrolls the view so the specified row is visible.
//
// row: The row index.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/scrollRowToVisible(_:)
func (t NSTableView) ScrollRowToVisible(row int) {
	objc.Send[objc.ID](t.ID, objc.Sel("scrollRowToVisible:"), row)
}
// Scrolls the view so the specified column is visible.
//
// column: The index of the column in the [TableColumns] array.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/scrollColumnToVisible(_:)
func (t NSTableView) ScrollColumnToVisible(column int) {
	objc.Send[objc.ID](t.ID, objc.Sel("scrollColumnToVisible:"), column)
}
// Computes and returns an image to use for dragging.
//
// dragRows: An index set containing the row indexes that should be in the image.
//
// tableColumns: An array of table columns that should be in the image.
//
// dragEvent: The event that initiated the drag.
//
// dragImageOffset: An in/out parameter specifying the offset of the cursor in the image, the
// default value is [NSZeroPoint]. Returning [NSZeroPoint] causes the cursor
// to be centered.
// //
// [NSZeroPoint]: https://developer.apple.com/documentation/Foundation/NSZeroPoint
//
// # Return Value
// 
// An [NSImage] containing a custom image for the specified rows and columns
// participating in the drag.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/dragImageForRows(with:tableColumns:event:offset:)
func (t NSTableView) DragImageForRowsWithIndexesTableColumnsEventOffset(dragRows foundation.NSIndexSet, tableColumns []NSTableColumn, dragEvent INSEvent, dragImageOffset foundation.NSPoint) INSImage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dragImageForRowsWithIndexes:tableColumns:event:offset:"), dragRows, objectivec.IObjectSliceToNSArray(tableColumns), dragEvent, dragImageOffset)
	return NSImageFromID(rv)
}
// Returns a Boolean value indicating whether the table view allows dragging
// the rows with the drag initiated at the specified point.
//
// rowIndexes: The row indexes to drag.
//
// mouseDownPoint: The location where the drag was initiated.
//
// # Return Value
// 
// [false] to disallow the drag.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/canDragRows(with:at:)
func (t NSTableView) CanDragRowsWithIndexesAtPoint(rowIndexes foundation.NSIndexSet, mouseDownPoint corefoundation.CGPoint) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("canDragRowsWithIndexes:atPoint:"), rowIndexes, mouseDownPoint)
	return rv
}
// Sets the default operation mask returned by `` to `mask`.
//
// mask: The drag operation mask. See [NSDragOperation] for the supported values.
// //
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// isLocal: [true] if the destination is the same application, otherwise [false]. In
// either case the specified `mask` value is archived and used.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/setDraggingSourceOperationMask(_:forLocal:)
func (t NSTableView) SetDraggingSourceOperationMaskForLocal(mask NSDragOperation, isLocal bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}
// Retargets the proposed drop operation.
//
// row: The target row index.
//
// dropOperation: The drop operation. Supported values are specified by
// [NSTableView.DropOperation].
// //
// [NSTableView.DropOperation]: https://developer.apple.com/documentation/AppKit/NSTableView/DropOperation
//
// # Discussion
// 
// For example, to specify a drop on the second row, specify `row` as 1, and
// `operation` as [NSTableViewDropOn]. To specify a drop below the last row,
// specify `row` as `[self numberOfRows]` and `operation` as
// [NSTableViewDropAbove].
// 
// Passing a value of `–1` for `row` and [NSTableViewDropOn] as the
// `operation` causes the entire table view to be highlighted rather than a
// specific row. This is useful if the data displayed by the table view does
// not allow the user to drop items at a specific row location.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/setDropRow(_:dropOperation:)
func (t NSTableView) SetDropRowDropOperation(row int, dropOperation NSTableViewDropOperation) {
	objc.Send[objc.ID](t.ID, objc.Sel("setDropRow:dropOperation:"), row, dropOperation)
}
// Hides the specified table rows.
//
// indexes: An index set containing indexes of the rows to be hidden.
//
// rowAnimation: An animation effect to be applied when the rows are hidden.
//
// # Discussion
// 
// Use this method when you no longer want the data to be visible to the user,
// but you don’t want to permanently remove the data. Hidden table rows have
// a height of zero and cannot be selected by the user. However, if a selected
// table row is hidden, it will remain selected.
// 
// Hiding a table row causes the [TableViewDidRemoveRowViewForRow] delegate
// method to be invoked.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/hideRows(at:withAnimation:)
func (t NSTableView) HideRowsAtIndexesWithAnimation(indexes foundation.NSIndexSet, rowAnimation NSTableViewAnimationOptions) {
	objc.Send[objc.ID](t.ID, objc.Sel("hideRowsAtIndexes:withAnimation:"), indexes, rowAnimation)
}
// Unhides the specified table rows.
//
// indexes: An index set containing indexes of the hidden rows to be shown again.
//
// rowAnimation: An animation effect to be applied when the rows are hidden.
//
// # Discussion
// 
// Unhiding a table row causes the [TableViewDidAddRowViewForRow] delegate
// method to be invoked.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/unhideRows(at:withAnimation:)
func (t NSTableView) UnhideRowsAtIndexesWithAnimation(indexes foundation.NSIndexSet, rowAnimation NSTableViewAnimationOptions) {
	objc.Send[objc.ID](t.ID, objc.Sel("unhideRowsAtIndexes:withAnimation:"), indexes, rowAnimation)
}
// Returns the column header accessibility elements for the table.
//
// # Return Value
// 
// The column header element.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityColumnHeaderUIElements] property.
//
// [accessibilityColumnHeaderUIElements]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnHeaderUIElements
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityColumnHeaderUIElements()
func (t NSTableView) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}
// Returns the column accessibility elements for the table.
//
// # Return Value
// 
// An array containing the table’s column elements.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityColumns] property.
//
// [accessibilityColumns]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumns
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityColumns()
func (t NSTableView) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}
// Returns the row header accessibility elements for the table.
//
// # Return Value
// 
// The row header elements for the table.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityRowHeaderUIElements] property.
//
// [accessibilityRowHeaderUIElements]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowHeaderUIElements
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityRowHeaderUIElements()
func (t NSTableView) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}
// Returns the row accessibility elements for the table.
//
// # Return Value
// 
// An array containing the table’s row elements.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityRows] property.
//
// [accessibilityRows]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRows
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityRows()
func (t NSTableView) AccessibilityRows() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("accessibilityRows"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// The currently selected cells for the table.
//
// # Return Value
// 
// An array containing the currently selected cells for the table.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilitySelectedCells] property. Additionally, your class needs to
// send a [selectedCellsChanged] notification whenever the table’s selected
// cells change.
//
// [accessibilitySelectedCells]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedCells
// [selectedCellsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedCellsChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilitySelectedCells()
func (t NSTableView) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}
// Returns the currently selected columns for the table.
//
// # Return Value
// 
// An array containing the currently selected columns for the table.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilitySelectedColumns] property. Additionally, your class needs to
// send a [selectedColumnsChanged] notification whenever the table’s
// selected columns change.
//
// [accessibilitySelectedColumns]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedColumns
// [selectedColumnsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedColumnsChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilitySelectedColumns()
func (t NSTableView) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}
// Returns the currently selected rows for the table.
//
// # Return Value
// 
// An array containing the currently selected rows for the table.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilitySelectedRows] property. Additionally, your class needs to
// send a [selectedRowsChanged] notification whenever the table’s selected
// rows change.
//
// [accessibilitySelectedRows]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedRows
// [selectedRowsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedRowsChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilitySelectedRows()
func (t NSTableView) AccessibilitySelectedRows() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("accessibilitySelectedRows"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// Returns the visible cells for the table.
//
// # Return Value
// 
// An array containing the currently visible cells.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityVisibleCells] property.
//
// [accessibilityVisibleCells]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCells
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityVisibleCells()
func (t NSTableView) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}
// Returns the visible columns for the table.
//
// # Return Value
// 
// An array containing the currently visible columns.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityVisibleColumns] property.
//
// [accessibilityVisibleColumns]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleColumns
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityVisibleColumns()
func (t NSTableView) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}
// Returns the visible rows for the table.
//
// # Return Value
// 
// An array containing the currently visible rows.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityVisibleRows] property.
//
// [accessibilityVisibleRows]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleRows
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityVisibleRows()
func (t NSTableView) AccessibilityVisibleRows() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("accessibilityVisibleRows"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// Invoked when the dragging session has completed.
//
// session: The dragging session.
//
// screenPoint: The point where the drag ended, in screen coordinates.
//
// operation: The drag operation. See [NSDragOperation] for drag operation types.
// //
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:endedAt:operation:)
func (t NSTableView) DraggingSessionEndedAtPointOperation(session INSDraggingSession, screenPoint corefoundation.CGPoint, operation NSDragOperation) {
	objc.Send[objc.ID](t.ID, objc.Sel("draggingSession:endedAtPoint:operation:"), session, screenPoint, operation)
}
// Invoked when the drag moves on the screen.
//
// session: The dragging session.
//
// screenPoint: The point where the drag moved to, in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:movedTo:)
func (t NSTableView) DraggingSessionMovedToPoint(session INSDraggingSession, screenPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](t.ID, objc.Sel("draggingSession:movedToPoint:"), session, screenPoint)
}
// Declares the types of operations the source allows to be performed.
//
// session: The dragging session.
//
// context: The dragging context. See [NSDraggingContext] for the supported values.
// //
// [NSDraggingContext]: https://developer.apple.com/documentation/AppKit/NSDraggingContext
//
// # Return Value
// 
// The appropriate dragging operation as defined in
//
// # Discussion
// 
// In the future Apple may provide more specific “within” values in the
// future. To account for this, for unrecognized localities, return the
// operation mask for the most specific context that you are concerned with.
// The following code is an example of how to implement this functionality:
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:sourceOperationMaskFor:)
func (t NSTableView) DraggingSessionSourceOperationMaskForDraggingContext(session INSDraggingSession, context NSDraggingContext) NSDragOperation {
	rv := objc.Send[NSDragOperation](t.ID, objc.Sel("draggingSession:sourceOperationMaskForDraggingContext:"), session, context)
	return NSDragOperation(rv)
}
// Invoked when the drag will begin.
//
// session: The dragging session.
//
// screenPoint: The point where the drag will begin, in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:willBeginAt:)
func (t NSTableView) DraggingSessionWillBeginAtPoint(session INSDraggingSession, screenPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](t.ID, objc.Sel("draggingSession:willBeginAtPoint:"), session, screenPoint)
}
// Returns whether the modifier keys will be ignored for this dragging
// session.
//
// session: The dragging session.
//
// # Return Value
// 
// [true] if the modifier keys will be ignored, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/ignoreModifierKeys(for:)
func (t NSTableView) IgnoreModifierKeysForDraggingSession(session INSDraggingSession) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("ignoreModifierKeysForDraggingSession:"), session)
	return rv
}
// Sets the table’s currently selected rows.
//
// selectedRows: An array containing the row elements to be selected.
//
// # Discussion
// 
// This method is the setter for the [NSAccessibilityProtocol] protocol’s
// [accessibilitySelectedRows] property. Implementing this method allows the
// user to change the selected row using an accessibility client.
// Additionally, your class needs to send a [selectedRowsChanged] notification
// whenever the table’s selected rows change.
//
// [accessibilitySelectedRows]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedRows
// [selectedRowsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedRowsChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/setAccessibilitySelectedRows(_:)
func (t NSTableView) SetAccessibilitySelectedRows(selectedRows []objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("setAccessibilitySelectedRows:"), objectivec.IObjectSliceToNSArray(selectedRows))
}
// Informs the delegate that the text object has begun editing (that the user
// has begun changing it).
//
// # Discussion
// 
// The name of `aNotification` is [didBeginEditingNotification].
//
// [didBeginEditingNotification]: https://developer.apple.com/documentation/AppKit/NSText/didBeginEditingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textDidBeginEditing(_:)
func (t NSTableView) TextDidBeginEditing(notification foundation.NSNotification) {
	objc.Send[objc.ID](t.ID, objc.Sel("textDidBeginEditing:"), notification)
}
// Informs the delegate that the text object has changed its characters or
// formatting attributes.
//
// # Discussion
// 
// The name of `aNotification` is [didChangeNotification].
//
// [didChangeNotification]: https://developer.apple.com/documentation/AppKit/NSText/didChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textDidChange(_:)
func (t NSTableView) TextDidChange(notification foundation.NSNotification) {
	objc.Send[objc.ID](t.ID, objc.Sel("textDidChange:"), notification)
}
// Informs the delegate that the text object has finished editing (that it has
// resigned first responder status).
//
// # Discussion
// 
// The name of `aNotification` is [didEndEditingNotification].
//
// [didEndEditingNotification]: https://developer.apple.com/documentation/AppKit/NSText/didEndEditingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textDidEndEditing(_:)
func (t NSTableView) TextDidEndEditing(notification foundation.NSNotification) {
	objc.Send[objc.ID](t.ID, objc.Sel("textDidEndEditing:"), notification)
}
// Invoked when a text object begins to change its text, this method requests
// permission for `aTextObject` to begin editing.
//
// # Discussion
// 
// If the delegate returns [true], the text object proceeds to make changes.
// If the delegate returns [false], the text object abandons the editing
// operation. This method is also invoked when the user drags and drops a file
// onto the text object.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textShouldBeginEditing(_:)
func (t NSTableView) TextShouldBeginEditing(textObject INSText) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("textShouldBeginEditing:"), textObject)
	return rv
}
// Invoked from a text object’s implementation of [ResignFirstResponder],
// this method requests permission for `aTextObject` to end editing.
//
// # Discussion
// 
// If the delegate returns [true], the text object proceeds to finish editing
// and resign first responder status. If the delegate returns [false], the
// text object selects all of its text and remains the first responder.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textShouldEndEditing(_:)
func (t NSTableView) TextShouldEndEditing(textObject INSText) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("textShouldEndEditing:"), textObject)
	return rv
}
// Returns an array of text objects to include in a text selection.
//
// # Return Value
// 
// An array of [NSTextCheckingResult] objects.
//
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:candidates:forSelectedRange:)
func (t NSTableView) TextViewWithCandidatesForSelectedRange(textView INSTextView, candidates []foundation.NSTextCheckingResult, selectedRange foundation.NSRange) []foundation.NSTextCheckingResult {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textView:candidates:forSelectedRange:"), textView, objectivec.IObjectSliceToNSArray(candidates), selectedRange)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSTextCheckingResult {
		return foundation.NSTextCheckingResultFromID(id)
	})
}
// Returns an array of objects that represent the elements of a selection.
//
// # Return Value
// 
// An array of objects that represent the selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:candidatesForSelectedRange:)
func (t NSTableView) TextViewCandidatesForSelectedRange(textView INSTextView, selectedRange foundation.NSRange) foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textView:candidatesForSelectedRange:"), textView, selectedRange)
	return foundation.NSArrayFromID(rv)
}
// Sent when the user clicks a cell.
//
// textView: The text view sending the message.
//
// cell: The cell clicked by the user.
//
// cellFrame: The frame of the clicked cell.
//
// charIndex: The character index of the clicked cell.
//
// # Discussion
// 
// The delegate can use this message as its cue to perform an action or select
// the attachment cell’s character. `aTextView` is the first text view in a
// series shared by a layout manager, not necessarily the one that draws
// `cell`.
// 
// The delegate may subsequently receive a
// [TextViewDoubleClickedOnCellInRectAtIndex] message if the user continues to
// perform a double click.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:clickedOn:in:at:)
func (t NSTableView) TextViewClickedOnCellInRectAtIndex(textView INSTextView, cell NSTextAttachmentCell, cellFrame corefoundation.CGRect, charIndex uint) {
	objc.Send[objc.ID](t.ID, objc.Sel("textView:clickedOnCell:inRect:atIndex:"), textView, cell, cellFrame, charIndex)
}
// Sent after the user clicks a link.
//
// textView: The text view sending the message.
//
// link: The link that was clicked; the value of [link].
// //
// [link]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/link
//
// charIndex: The character index where the click occurred, indexed within the text
// storage.
//
// # Return Value
// 
// [true] if the click was handled; otherwise, [false] to allow the next
// responder to handle it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate can use this method to handle the click on the link. It is
// invoked by [ClickedOnLinkAtIndex].
// 
// The `charIndex` parameter is a character index somewhere in the range of
// the link attribute. If the user actually physically clicked the link, then
// it should be the character that was originally clicked. In some cases a
// link may be opened indirectly or programmatically, in which case a
// character index somewhere in the range of the link attribute is supplied.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:clickedOnLink:at:)
func (t NSTableView) TextViewClickedOnLinkAtIndex(textView INSTextView, link objectivec.IObject, charIndex uint) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("textView:clickedOnLink:atIndex:"), textView, link, charIndex)
	return rv
}
// Returns the actual completions for a partial word.
//
// textView: The text view sending the message.
//
// words: The proposed array of completions.
//
// charRange: The range of characters to be completed.
//
// index: On return, the index of the initially selected completion. The default is
// 0, and –1 indicates no selection.
//
// # Return Value
// 
// The actual array of completions that will be presented for the partial word
// at the given range. Returning `nil` or a zero-length array suppresses
// completion.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:completions:forPartialWordRange:indexOfSelectedItem:)
func (t NSTableView) TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView INSTextView, words []string, charRange foundation.NSRange, index unsafe.Pointer) []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textView:completions:forPartialWordRange:indexOfSelectedItem:"), textView, objectivec.StringSliceToNSArray(words), charRange, index)
	return objc.ConvertSliceToStrings(rv)
}
// Sent when the selection changes in the text view.
//
// notification: A notification named [didChangeSelectionNotification].
// //
// [didChangeSelectionNotification]: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeSelectionNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textViewDidChangeSelection(_:)
func (t NSTableView) TextViewDidChangeSelection(notification foundation.NSNotification) {
	objc.Send[objc.ID](t.ID, objc.Sel("textViewDidChangeSelection:"), notification)
}
// Sent when a text view’s typing attributes change.
//
// notification: A notification named [didChangeTypingAttributesNotification].
// //
// [didChangeTypingAttributesNotification]: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeTypingAttributesNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textViewDidChangeTypingAttributes(_:)
func (t NSTableView) TextViewDidChangeTypingAttributes(notification foundation.NSNotification) {
	objc.Send[objc.ID](t.ID, objc.Sel("textViewDidChangeTypingAttributes:"), notification)
}
// Invoked to allow the delegate to modify the text checking results after
// checking has occurred.
//
// view: The text view sending the message.
//
// range: The range that was checked.
//
// checkingTypes: The type of checking that was performed. The possible constants are listed
// in [NSTextCheckingTypes] and can be combined using the C bit-wise [OR]
// operator to perform multiple checks at the same time.
// //
// [NSTextCheckingTypes]: https://developer.apple.com/documentation/Foundation/NSTextCheckingTypes
//
// options: A dictionary of values used during the checking process to perform. See
// Spell Checking Option Dictionary Keys for the supported values.
//
// results: An array of [NSTextCheckingResult] instances.
// //
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
//
// orthography: The orthography of the text.
//
// wordCount: The number of words checked.
//
// # Return Value
// 
// An array of [NSTextCheckingResult] instances. You can return the results
// array as is, or an altered array of [NSTextCheckingResult] objects.
//
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
//
// # Discussion
// 
// Invoked by
// [HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount], this
// method allows observation of text checking, or modification of the results
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:didCheckTextIn:types:options:results:orthography:wordCount:)
func (t NSTableView) TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view INSTextView, range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary, results []foundation.NSTextCheckingResult, orthography foundation.NSOrthography, wordCount int) []foundation.NSTextCheckingResult {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textView:didCheckTextInRange:types:options:results:orthography:wordCount:"), view, range_, checkingTypes, options, objectivec.IObjectSliceToNSArray(results), orthography, wordCount)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSTextCheckingResult {
		return foundation.NSTextCheckingResultFromID(id)
	})
}
// Sent to allow the delegate to perform the command for the text view.
//
// textView: The text view sending the message. This is the first text view in a series
// shared by a layout manager.
//
// commandSelector: The selector.
//
// # Return Value
// 
// [true] indicates that the delegate handled the command and the text view
// will not attempt to perform it; [false] indicates that the delegate did not
// handle the command the text view will attempt to perform it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is invoked by [NSTextView]’s `doCommand()` method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:doCommandBy:)
func (t NSTableView) TextViewDoCommandBySelector(textView INSTextView, commandSelector objc.SEL) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("textView:doCommandBySelector:"), textView, commandSelector)
	return rv
}
// Sent when the user double-clicks a cell.
//
// textView: The text view sending the message.
//
// cell: The cell double-clicked by the user.
//
// cellFrame: The frame of the double-clicked cell.
//
// charIndex: The character index of the double-clicked cell.
//
// # Discussion
// 
// The delegate can use this message as its cue to perform an action, such as
// opening the file represented by the attachment. `aTextView` is the first
// text view in a series shared by a layout manager, not necessarily the one
// that draws `cell`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:doubleClickedOn:in:at:)
func (t NSTableView) TextViewDoubleClickedOnCellInRectAtIndex(textView INSTextView, cell NSTextAttachmentCell, cellFrame corefoundation.CGRect, charIndex uint) {
	objc.Send[objc.ID](t.ID, objc.Sel("textView:doubleClickedOnCell:inRect:atIndex:"), textView, cell, cellFrame, charIndex)
}
// Sent when the user attempts to drag a cell.
//
// view: The text view sending the message.
//
// cell: The cell being dragged.
//
// rect: The rectangle from which the cell was dragged.
//
// event: The mouse-down event that preceded the mouse-dragged event.
//
// charIndex: The character position where the mouse button was clicked.
//
// # Discussion
// 
// The delegate can use this message as its cue to initiate a dragging
// operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:draggedCell:in:event:at:)
func (t NSTableView) TextViewDraggedCellInRectEventAtIndex(view INSTextView, cell NSTextAttachmentCell, rect corefoundation.CGRect, event INSEvent, charIndex uint) {
	objc.Send[objc.ID](t.ID, objc.Sel("textView:draggedCell:inRect:event:atIndex:"), view, cell, rect, event, charIndex)
}
// Allows delegate to control the context menu returned by the text view.
//
// view: The text view sending the message.
//
// menu: The proposed contextual menu.
//
// event: The mouse-down event that initiated the contextual menu’s display.
//
// charIndex: The character position where the mouse button was clicked.
//
// # Return Value
// 
// A menu to use as the contextual menu. You can return `menu` unaltered, or
// you can return a customized menu.
//
// # Discussion
// 
// This method allows the delegate to control the context menu returned by
// [MenuForEvent].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:menu:for:at:)
func (t NSTableView) TextViewMenuForEventAtIndex(view INSTextView, menu INSMenu, event INSEvent, charIndex uint) INSMenu {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textView:menu:forEvent:atIndex:"), view, menu, event, charIndex)
	return NSMenuFromID(rv)
}
// Sent when a text view needs to determine if text in a specified range
// should be changed.
//
// textView: The text view sending the message. This is the first text view in a series
// shared by a layout manager, not necessarily the text view displaying the
// selected text.
//
// affectedCharRange: The range of characters to be replaced.
//
// replacementString: The characters that will replace the characters in `affectedCharRange`;
// `nil` if only text attributes are being changed.
//
// # Return Value
// 
// [true] to allow the replacement, or [false] to reject the change.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If a delegate implements this method and not its multiple-selection
// replacement, [TextViewShouldChangeTextInRangesReplacementStrings], it is
// called with an appropriate range and string. If a delegate implements the
// new method, then this one is ignored.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldChangeTextIn:replacementString:)
func (t NSTableView) TextViewShouldChangeTextInRangeReplacementString(textView INSTextView, affectedCharRange foundation.NSRange, replacementString string) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("textView:shouldChangeTextInRange:replacementString:"), textView, affectedCharRange, objc.String(replacementString))
	return rv
}
// Sent when a text view needs to determine if text in an array of specified
// ranges should be changed.
//
// textView: The text view sending the message. This is the first text view in a series
// shared by a layout manager, not necessarily the text view displaying the
// selected text.
//
// affectedRanges: The array of ranges of characters to be replaced. This array must be a
// non-nil, non-empty array of objects responding to the NSValue `rangeValue`
// method, and in addition its elements must be sorted, non-overlapping,
// non-contiguous, and (except for the case of a single range) have
// non-zero-length.
//
// replacementStrings: The array of strings that will replace the characters in `affectedRanges`,
// one string for each range; `nil` if only text attributes are being changed.
//
// # Return Value
// 
// [true] to allow the replacement, or [false] to reject the change.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldChangeTextInRanges:replacementStrings:)
func (t NSTableView) TextViewShouldChangeTextInRangesReplacementStrings(textView INSTextView, affectedRanges []foundation.NSValue, replacementStrings []string) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("textView:shouldChangeTextInRanges:replacementStrings:"), textView, objectivec.IObjectSliceToNSArray(affectedRanges), objectivec.StringSliceToNSArray(replacementStrings))
	return rv
}
// Sent when the typing attributes are changed.
//
// textView: The text view sending the message.
//
// oldTypingAttributes: The old typing attributes.
//
// newTypingAttributes: The proposed typing attributes.
//
// # Return Value
// 
// The actual new typing attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldChangeTypingAttributes:toAttributes:)
func (t NSTableView) TextViewShouldChangeTypingAttributesToAttributes(textView INSTextView, oldTypingAttributes foundation.INSDictionary, newTypingAttributes foundation.INSDictionary) foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textView:shouldChangeTypingAttributes:toAttributes:"), textView, oldTypingAttributes, newTypingAttributes)
	return foundation.NSDictionaryFromID(rv)
}
// Returns a Boolean value that indicates whether to select the text object at
// the index.
//
// textView: The text view that sent the message.
//
// index: The index that represents the start of the candidate text to evaluate.
//
// # Return Value
// 
// Returns [true] if the framework selects the text.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldSelectCandidateAt:)
func (t NSTableView) TextViewShouldSelectCandidateAtIndex(textView INSTextView, index uint) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("textView:shouldSelectCandidateAtIndex:"), textView, index)
	return rv
}
// Sent when the spelling state is changed.
//
// textView: The text view sending the message.
//
// value: The proposed spelling state value to set. Possible values, for the
// temporary attribute on the layout manager using the key
// NSSpellingStateAttributeName, are:
// 
// - [NSSpellingStateSpellingFlag] to highlight spelling issues. -
// [NSSpellingStateGrammarFlag] to highlight grammar issues.
// //
// [NSSpellingStateGrammarFlag]: https://developer.apple.com/documentation/AppKit/NSSpellingState/NSSpellingStateGrammarFlag
// [NSSpellingStateSpellingFlag]: https://developer.apple.com/documentation/AppKit/NSSpellingState/NSSpellingStateSpellingFlag
//
// affectedCharRange: The character range over which to set the given spelling state.
//
// # Return Value
// 
// The actual spelling state to set.
//
// # Discussion
// 
// Delegate only. Allows delegate to control the setting of spelling and
// grammar indicators.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldSetSpellingState:range:)
func (t NSTableView) TextViewShouldSetSpellingStateRange(textView INSTextView, value int, affectedCharRange foundation.NSRange) int {
	rv := objc.Send[int](t.ID, objc.Sel("textView:shouldSetSpellingState:range:"), textView, value, affectedCharRange)
	return rv
}
// Returns and array of touch bar elements for the framework to update.
//
// textView: The text view that sent the message.
//
// identifiers: An array of touch bar identifiers to evaluate.
//
// # Return Value
// 
// Returns an array of [NSTouchBarItemIdentifier] elements for framework to
// update.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldUpdateTouchBarItemIdentifiers:)
func (t NSTableView) TextViewShouldUpdateTouchBarItemIdentifiers(textView INSTextView, identifiers []string) []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textView:shouldUpdateTouchBarItemIdentifiers:"), textView, objectivec.StringSliceToNSArray(identifiers))
	return objc.ConvertSliceToStrings(rv)
}
// Returns a URL representing the document contents for a text attachment.
//
// textView: The text view sending the message.
//
// textAttachment: The text attachment object containing an [NSFileWrapper] object that holds
// the contents of the attached file.
//
// charIndex: The character index of the text attachment.
//
// # Return Value
// 
// The absolute URL for the document contents represented by `textAttachment`.
//
// # Discussion
// 
// The returned [NSURL] object is used by the text view to provide default
// behaviors involving text attachments such as Quick Look and
// double-clicking. For example, the [NSTextView] method
// [QuickLookPreviewableItemsInRanges] uses this method for mapping text
// attachments to their corresponding document URLs, and [NSTextView] invokes
// the [NSWorkspace] method [OpenURL] with the URL returned from this method
// when the delegate has no [TextViewDoubleClickedOnCellInRectAtIndex]
// implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:urlForContentsOf:at:)
func (t NSTableView) TextViewURLForContentsOfTextAttachmentAtIndex(textView INSTextView, textAttachment INSTextAttachment, charIndex uint) foundation.INSURL {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textView:URLForContentsOfTextAttachment:atIndex:"), textView, textAttachment, charIndex)
	return foundation.NSURLFromID(rv)
}
// Returns the actual range to select.
//
// textView: The text view sending the message. This is the first text view in a series
// shared by a layout manager, not necessarily the text view displaying the
// selected text.
//
// oldSelectedCharRange: The original range of the selection.
//
// newSelectedCharRange: The proposed character range for the new selection.
//
// # Return Value
// 
// The actual character range for the new selection.
//
// # Discussion
// 
// This method is invoked before a text view finishes changing the
// selection—that is, when the last argument to a
// [SetSelectedRangeAffinityStillSelecting] message is [false].
// 
// Non-selectable text views do not process any mouse events. If for some
// reason it is necessary to disallow user selection change in a text view
// that handles mouse events, this can be achieved by making the text view
// selectable but implementing this delegate method to disallow selection
// changes.
// 
// # Special Considerations
// 
// In macOS 10.4 and later, if a delegate implements this delegate method and
// not its multiple-selection replacement,
// [TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges], then
// multiple selection is effectively disallowed; attempts to set the selected
// ranges call the old delegate method with the first subrange, and afterwards
// only a single selected range is set.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:willChangeSelectionFromCharacterRange:toCharacterRange:)
func (t NSTableView) TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView INSTextView, oldSelectedCharRange foundation.NSRange, newSelectedCharRange foundation.NSRange) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("textView:willChangeSelectionFromCharacterRange:toCharacterRange:"), textView, oldSelectedCharRange, newSelectedCharRange)
	return foundation.NSRange(rv)
}
// Returns the actual character ranges to select.
//
// textView: The text view sending the message. This is the first text view in a series
// shared by a layout manager, not necessarily the text view displaying the
// selected text.
//
// oldSelectedCharRanges: An array containing the original ranges of the selection. This must be a
// non-`nil`, non-empty array of objects responding to the [NSValue] method
// `rangeValue`, and in addition its elements must be sorted, non-overlapping,
// non-contiguous, and (except for the case of a single range) have
// non-zero-length.
//
// newSelectedCharRanges: An array containing the proposed character ranges for the new selection.
// This must be a non-`nil`, non-empty array of objects responding to the
// [NSValue] method `rangeValue`, and in addition its elements must be sorted,
// non-overlapping, non-contiguous, and (except for the case of a single
// range) have non-zero-length.
//
// # Return Value
// 
// An array containing the actual character ranges for the new selection.
//
// # Discussion
// 
// Invoked before an [NSTextView] object finishes changing the
// selection—that is, when the last argument to a
// [SetSelectedRangeAffinityStillSelecting] or
// [SetSelectedRangesAffinityStillSelecting] message is [false].
// 
// Non-selectable text views do not process any mouse events. If for some
// reason it is necessary to disallow user selection change in a text view
// that handles mouse events, this can be achieved by making the text view
// selectable but implementing this delegate method to disallow selection
// changes.
// 
// If a delegate implements both this method and
// [TextViewWillChangeSelectionFromCharacterRangeToCharacterRange], then the
// latter is ignored.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:willChangeSelectionFromCharacterRanges:toCharacterRanges:)
func (t NSTableView) TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView INSTextView, oldSelectedCharRanges []foundation.NSValue, newSelectedCharRanges []foundation.NSValue) []foundation.NSValue {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:"), textView, objectivec.IObjectSliceToNSArray(oldSelectedCharRanges), objectivec.IObjectSliceToNSArray(newSelectedCharRanges))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
// Invoked to allow the delegate to modify the text checking process before it
// occurs.
//
// view: The text view sending the message.
//
// range: The range to be checked.
//
// options: A dictionary of values used during the checking process to perform. See
// Spell Checking Option Dictionary Keys for the supported values.
//
// checkingTypes: The type of checking to be performed, passed by-reference. The possible
// constants are listed in [NSTextCheckingTypes] and can be combined using the
// C bit-wise [OR] operator to perform multiple checks at the same time.
// 
// You can change this parameter to alter the types of checking to be
// performed.
// //
// [NSTextCheckingTypes]: https://developer.apple.com/documentation/Foundation/NSTextCheckingTypes
//
// # Return Value
// 
// A dictionary containing an alternative to the options `dictionary`.
//
// # Discussion
// 
// Invoked by [CheckTextInRangeTypesOptions], this method allows control over
// text checking `options`s (via the return value) or types (by modifying the
// flags pointed to by the inout parameter `checkingTypes`)
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:willCheckTextIn:options:types:)
func (t NSTableView) TextViewWillCheckTextInRangeOptionsTypes(view INSTextView, range_ foundation.NSRange, options foundation.INSDictionary, checkingTypes unsafe.Pointer) foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textView:willCheckTextInRange:options:types:"), view, range_, options, checkingTypes)
	return foundation.NSDictionaryFromID(rv)
}
// Returns the actual tooltip to display.
//
// textView: The text view sending the message.
//
// tooltip: The proposed tooltip to display.
//
// characterIndex: The location in `textView`.
//
// # Return Value
// 
// The actual tooltip to display, or `nil` to suppress display of the tooltip.
//
// # Discussion
// 
// The tooltip string is the value of the [toolTip] attribute at
// `characterIndex`.
//
// [toolTip]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/toolTip
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:willDisplayToolTip:forCharacterAt:)
func (t NSTableView) TextViewWillDisplayToolTipForCharacterAtIndex(textView INSTextView, tooltip string, characterIndex uint) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textView:willDisplayToolTip:forCharacterAtIndex:"), textView, objc.String(tooltip), characterIndex)
	return foundation.NSStringFromID(rv).String()
}
// Returns a sharing service picker for the current selection.
//
// textView: The text view.
//
// servicePicker: The service picker.
//
// items: The ranges of the items to share.
//
// # Return Value
// 
// An [NSSharingServicePicker] instance. The original sharing picker or a new
// sharing picker instance can be returned.
//
// # Discussion
// 
// Returns a sharing service picker created for items right before shown to
// the screen when the `` method. Return `nil` to remove the Share item from
// the menu.
// 
// The delegate is specified as the delegate for the [NSSharingServicePicker]
// instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:willShow:forItems:)
func (t NSTableView) TextViewWillShowSharingServicePickerForItems(textView INSTextView, servicePicker INSSharingServicePicker, items foundation.INSArray) INSSharingServicePicker {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textView:willShowSharingServicePicker:forItems:"), textView, servicePicker, items)
	return NSSharingServicePickerFromID(rv)
}
// Returns the writable pasteboard types for a given cell.
//
// view: The text view sending the message.
//
// cell: The cell in question.
//
// charIndex: The character index in the text view that was clicked.
//
// # Return Value
// 
// An array of types that can be written to the pasteboard for `cell`.
//
// # Discussion
// 
// This method is invoked after the user clicks `cell` at the specified
// `charIndex` location in `aTextView`. If the
// [TextViewDraggedCellInRectEventAtIndex] is not used, this method and
// [TextViewWriteCellAtIndexToPasteboardType] allow `aTextView` to take care
// of attachment dragging and pasting, with the delegate responsible only for
// writing the attachment to the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:writablePasteboardTypesFor:at:)
func (t NSTableView) TextViewWritablePasteboardTypesForCellAtIndex(view INSTextView, cell NSTextAttachmentCell, charIndex uint) []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textView:writablePasteboardTypesForCell:atIndex:"), view, cell, charIndex)
	return objc.ConvertSliceToStrings(rv)
}
// Returns whether data of the specified type for the given cell could be
// written to the specified pasteboard.
//
// view: The text view sending the message.
//
// cell: The cell whose contents should be written to the pasteboard.
//
// charIndex: The index at which the cell was accessed.
//
// pboard: The pasteboard to which the cell’s contents should be written.
//
// type: The type of data that should be written.
//
// # Return Value
// 
// [true] if the write succeeded, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The receiver should attempt to write the `cell` to `pboard` with the given
// `type`, and return success or failure.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:write:at:to:type:)
func (t NSTableView) TextViewWriteCellAtIndexToPasteboardType(view INSTextView, cell NSTextAttachmentCell, charIndex uint, pboard INSPasteboard, type_ NSPasteboardType) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("textView:writeCell:atIndex:toPasteboard:type:"), view, cell, charIndex, pboard, objc.String(string(type_)))
	return rv
}
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textViewWritingToolsDidEnd(_:)
func (t NSTableView) TextViewWritingToolsDidEnd(textView INSTextView) {
	objc.Send[objc.ID](t.ID, objc.Sel("textViewWritingToolsDidEnd:"), textView)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:writingToolsIgnoredRangesInEnclosingRange:)
func (t NSTableView) TextViewWritingToolsIgnoredRangesInEnclosingRange(textView INSTextView, enclosingRange foundation.NSRange) []foundation.NSValue {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textView:writingToolsIgnoredRangesInEnclosingRange:"), textView, enclosingRange)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textViewWritingToolsWillBegin(_:)
func (t NSTableView) TextViewWritingToolsWillBegin(textView INSTextView) {
	objc.Send[objc.ID](t.ID, objc.Sel("textViewWritingToolsWillBegin:"), textView)
}
// Returns the undo manager for the specified text view.
//
// view: The text view whose undo manager should be returned.
//
// # Return Value
// 
// The undo manager for `view`.
//
// # Discussion
// 
// This method provides the flexibility to return a custom undo manager for
// the text view. Although [NSTextView] implements undo and redo for changes
// to text, applications may need a custom undo manager to handle interactions
// between changes to text and changes to other items in the application.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/undoManager(for:)
func (t NSTableView) UndoManagerForTextView(view INSTextView) foundation.NSUndoManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("undoManagerForTextView:"), view)
	return foundation.NSUndoManagerFromID(rv)
}
// Returns a Boolean value that indicates whether the sender should be
// enabled.
//
// item: The user interface item to validate. You can send `anItem` the [Action] and
// [Tag] messages.
//
// # Return Value
// 
// [true] if the user interface item should be enabled, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceValidations/validateUserInterfaceItem(_:)
func (t NSTableView) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}

// The object that provides the data displayed by the table view.
//
// # Discussion
// 
// The data source for the table view must implement the appropriate methods
// of the [NSTableViewDataSource] protocol. See [Populating a Table View
// Programmatically] and the [NSTableViewDataSource] `protocol` specification
// for more information. Note that in versions of macOS prior to v10.12, the
// table view did not retain the data source in a managed memory environment.
// 
// Setting the data source invokes [Tile].
// 
// If the delegate doesn’t respond to either [NumberOfRowsInTableView] or
// [TableViewObjectValueForTableColumnRow], [internalInconsistencyException]
// may be raised.
//
// [Populating a Table View Programmatically]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/TableView/PopulatingView-TablesProgrammatically/PopulatingView-TablesProgrammatically.html#//apple_ref/doc/uid/10000026i-CH14
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/dataSource
func (t NSTableView) DataSource() NSTableViewDataSource {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dataSource"))
	return NSTableViewDataSourceObjectFromID(rv)
}
func (t NSTableView) SetDataSource(value NSTableViewDataSource) {
	objc.Send[struct{}](t.ID, objc.Sel("setDataSource:"), value)
}
// A Boolean value indicating whether the table uses static data.
//
// # Discussion
// 
// A static table does not rely on a data source to provide the number of
// rows. A static table view’s contents are set at design time and can be
// changed programmatically as needed. Typically, you do not change the
// contents of a static table view after setting them.
// 
// In Xcode, any rows you add to a static table are saved in the corresponding
// nib or storyboard file and loaded with the rest of the table at runtime.
// You can add table rows programmatically to a static table view using the
// [InsertRowsAtIndexesWithAnimation] method. When adding rows
// programmatically, your table view delegate must implement the
// [TableViewViewForTableColumnRow] method to provide the corresponding view
// for any new rows. You can also remove rows at any time using the
// [RemoveRowsAtIndexesWithAnimation] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/usesStaticContents
func (t NSTableView) UsesStaticContents() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesStaticContents"))
	return rv
}
func (t NSTableView) SetUsesStaticContents(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesStaticContents:"), value)
}
// The dictionary of all registered nib files for view-based table view
// identifiers.
//
// # Discussion
// 
// Each key in the dictionary is the identifier string (given by
// [NSUserInterfaceItemIdentifier]) used to register the nib file in the
// [RegisterNibForIdentifier] method. The value of each key is the
// corresponding [NSNib] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/registeredNibsByIdentifier
func (t NSTableView) RegisteredNibsByIdentifier() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("registeredNibsByIdentifier"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The message sent to the table view’s target when the user double-clicks a
// cell or column header.
//
// # Discussion
// 
// This property stores a selector that corresponds to a method of the
// following form:
// 
// When the user double-clicks a cell or column header, the table calls the
// specified method of its [Target] object. The default value of this property
// is nil. If you do not specify a value for this property, the table view
// begins editing the cell.
// 
// The [ClickedRow] and [ClickedColumn] properties allow you to determine
// which row and column the double-click occurred in or if, rather than in a
// row, the double-click occurred in a column heading.
// 
// Note that if the table view uses Cocoa bindings and the Double Click Target
// binding is bound, both messages are invoked on their respective targets:
// First the Cocoa binding message is sent, then the `` message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/doubleAction
func (t NSTableView) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](t.ID, objc.Sel("doubleAction"))
	return rv
}
func (t NSTableView) SetDoubleAction(value objc.SEL) {
	objc.Send[struct{}](t.ID, objc.Sel("setDoubleAction:"), value)
}
// The index of the column the user clicked.
//
// # Discussion
// 
// This property contains the index in the [TableColumns] array of the column
// that the user clicked. The value is `-1` when the user clicks in an area of
// the table view that is not occupied by columns or when the user clicks a
// row that is a group separator.
// 
// The value of this property is meaningful in the target object’s
// implementation of the action and double-action methods. You can also use
// the value to determine which contextual menu to display when the user
// Control-clicks in a table. Note that the `clickedColumn` value remains
// valid when the menu item sends the action message. To see an example of
// using `clickedColumn` in the implementation of a contextual menu, download
// the [DragNDropOutlineView: implementing drag and drop in an NSOutlineView]
// sample project.
//
// [DragNDropOutlineView: implementing drag and drop in an NSOutlineView]: https://developer.apple.com/library/archive/samplecode/DragNDropOutlineView/Introduction/Intro.html#//apple_ref/doc/uid/DTS40008831
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/clickedColumn
func (t NSTableView) ClickedColumn() int {
	rv := objc.Send[int](t.ID, objc.Sel("clickedColumn"))
	return rv
}
// The index of the row the user clicked.
//
// # Return Value
// 
// The index of the row the user clicked to trigger an action message. Returns
// `–1` if the user clicked in an area of the table view not occupied by
// table rows.
// 
// # Discussion
// 
// This property contains the index of the row that the user clicked. The
// value is `-1` when the user clicks in an area of the table view that is not
// occupied by table rows.
// 
// The value of this property is meaningful in the target object’s
// implementation of the action and double-action methods. You can also use
// the value to determine which contextual menu to display when the user
// Control-clicks in a table. Note that you should check to see if
// `clickedRow` is one of the rows the user selected and if it is, perform the
// contextual menu operation on all of the selected rows. To see an example of
// using `clickedRow` in the implementation of a contextual menu, download the
// [DragNDropOutlineView: implementing drag and drop in an NSOutlineView]
// sample project.
//
// [DragNDropOutlineView: implementing drag and drop in an NSOutlineView]: https://developer.apple.com/library/archive/samplecode/DragNDropOutlineView/Introduction/Intro.html#//apple_ref/doc/uid/DTS40008831
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/clickedRow
func (t NSTableView) ClickedRow() int {
	rv := objc.Send[int](t.ID, objc.Sel("clickedRow"))
	return rv
}
// A Boolean value indicating whether the table view allows the user to
// rearrange columns by dragging their headers.
//
// # Discussion
// 
// The default value of this property is [true], which allows the user to
// rearrange the table view’s columns. You can rearrange columns
// programmatically regardless of this setting.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnReordering
func (t NSTableView) AllowsColumnReordering() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsColumnReordering"))
	return rv
}
func (t NSTableView) SetAllowsColumnReordering(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsColumnReordering:"), value)
}
// A Boolean value indicating whether the table view allows the user to resize
// columns by dragging between their headers.
//
// # Discussion
// 
// The default of this property is [true], which allows the user to resize the
// table view’s columns. You can resize columns programmatically regardless
// of this setting.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnResizing
func (t NSTableView) AllowsColumnResizing() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsColumnResizing"))
	return rv
}
func (t NSTableView) SetAllowsColumnResizing(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsColumnResizing:"), value)
}
// A Boolean value indicating whether the table view allows the user to select
// more than one column or row at a time.
//
// # Discussion
// 
// The default is [false], which allows the user to select only one column or
// row at a time. You can select multiple columns or rows programmatically
// regardless of this setting.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/allowsMultipleSelection
func (t NSTableView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}
func (t NSTableView) SetAllowsMultipleSelection(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}
// A Boolean value indicating whether the table view allows the user to select
// zero columns or rows.
//
// # Discussion
// 
// The default is [true], which allows the user to select zero columns or
// rows.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/allowsEmptySelection
func (t NSTableView) AllowsEmptySelection() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsEmptySelection"))
	return rv
}
func (t NSTableView) SetAllowsEmptySelection(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsEmptySelection:"), value)
}
// A Boolean value indicating whether the table view allows the user to select
// columns by clicking their headers.
//
// # Discussion
// 
// The default is [false], which prevents the user from selecting columns (if
// you create the table view in Interface Builder, the default value is
// [true]). You can select columns programmatically regardless of this
// setting.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/allowsColumnSelection
func (t NSTableView) AllowsColumnSelection() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsColumnSelection"))
	return rv
}
func (t NSTableView) SetAllowsColumnSelection(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsColumnSelection:"), value)
}
// A Boolean value that indicates whether the table view uses autolayout to
// calculate the height of rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/usesAutomaticRowHeights
func (t NSTableView) UsesAutomaticRowHeights() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesAutomaticRowHeights"))
	return rv
}
func (t NSTableView) SetUsesAutomaticRowHeights(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesAutomaticRowHeights:"), value)
}
// The horizontal and vertical spacing between cells.
//
// # Discussion
// 
// Changing the value of this property causes the table view to redisplay
// itself. Negative values aren’t supported. The default spacing varies
// based on the table’s style.
// 
// Table views normally have a 1-pixel separation between consecutively
// selected rows or columns. An intercell spacing of `(1.0, 1.0)` or greater
// is required if you want this separation. An intercell spacing of `(0.0,
// 0.0)` forces no separation between consecutive selections.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/intercellSpacing
func (t NSTableView) IntercellSpacing() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("intercellSpacing"))
	return corefoundation.CGSize(rv)
}
func (t NSTableView) SetIntercellSpacing(value corefoundation.CGSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setIntercellSpacing:"), value)
}
// The height of each row in the table.
//
// # Discussion
// 
// The default row height is `16.0`. The value in this property is used only
// if the table’s [RowSizeStyle] is set to
// [NSTableView.RowSizeStyle.custom].
// 
// When you change the value of this property, the table view calls the [Tile]
// method to redisplay the rows using the new value.
//
// [NSTableView.RowSizeStyle.custom]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/custom
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/rowHeight
func (t NSTableView) RowHeight() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("rowHeight"))
	return rv
}
func (t NSTableView) SetRowHeight(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setRowHeight:"), value)
}
// The color used to draw the background of the table.
//
// # Discussion
// 
// The default background color is light gray.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/backgroundColor
func (t NSTableView) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTableView) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setBackgroundColor:"), value)
}
// A Boolean value indicating whether the table view uses alternating row
// colors for its background.
//
// # Return Value
// 
// [true] if the table view uses standard alternating row colors for the
// background, [false] if it uses a solid color.
// 
// # Discussion
// 
// When the value of this property is [true], the table uses the standard
// alternating row colors for the background. When the value is [false], the
// table view uses a single solid color for the background.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/usesAlternatingRowBackgroundColors
func (t NSTableView) UsesAlternatingRowBackgroundColors() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesAlternatingRowBackgroundColors"))
	return rv
}
func (t NSTableView) SetUsesAlternatingRowBackgroundColors(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesAlternatingRowBackgroundColors:"), value)
}
// The style that the table view uses.
//
// # Discussion
// 
// The default value for this property is [NSTableView.Style.automatic] in
// macOS 11 and later. Apps that link to previous macOS versions default to
// [NSTableView.Style.plain].
//
// [NSTableView.Style.automatic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/automatic
// [NSTableView.Style.plain]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/plain
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/style-swift.property
func (t NSTableView) Style() NSTableViewStyle {
	rv := objc.Send[NSTableViewStyle](t.ID, objc.Sel("style"))
	return NSTableViewStyle(rv)
}
func (t NSTableView) SetStyle(value NSTableViewStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setStyle:"), value)
}
// The effective style that the table uses.
//
// # Discussion
// 
// If the [Style] property value is [NSTableView.Style.automatic], then this
// property contains the resolved style.
//
// [NSTableView.Style.automatic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/automatic
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/effectiveStyle
func (t NSTableView) EffectiveStyle() NSTableViewStyle {
	rv := objc.Send[NSTableViewStyle](t.ID, objc.Sel("effectiveStyle"))
	return NSTableViewStyle(rv)
}
// The selection highlight style used by the table view to indicate row and
// column selection.
//
// # Discussion
// 
// Setting the selection highlight style to
// [NSTableView.SelectionHighlightStyle.sourceList] causes the table view to
// draw its background using the source list style. It also sets the
// [DraggingDestinationFeedbackStyle] to
// [NSTableView.DraggingDestinationFeedbackStyle.sourceList].
//
// [NSTableView.DraggingDestinationFeedbackStyle.sourceList]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum/sourceList
// [NSTableView.SelectionHighlightStyle.sourceList]: https://developer.apple.com/documentation/AppKit/NSTableView/SelectionHighlightStyle-swift.enum/sourceList
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/selectionHighlightStyle-swift.property
func (t NSTableView) SelectionHighlightStyle() NSTableViewSelectionHighlightStyle {
	rv := objc.Send[NSTableViewSelectionHighlightStyle](t.ID, objc.Sel("selectionHighlightStyle"))
	return NSTableViewSelectionHighlightStyle(rv)
}
func (t NSTableView) SetSelectionHighlightStyle(value NSTableViewSelectionHighlightStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectionHighlightStyle:"), value)
}
// The color used to draw grid lines.
//
// # Discussion
// 
// The default color is gray.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/gridColor
func (t NSTableView) GridColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("gridColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTableView) SetGridColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setGridColor:"), value)
}
// The grid lines drawn by the table view.
//
// # Discussion
// 
// Use this property to specify whether lines should be drawn between rows and
// columns. When setting this property, you can specify multiple styles at
// once by adding the corresponding constants together. The default value of
// this property is [NSTableViewGridNone].
//
// [NSTableViewGridNone]: https://developer.apple.com/documentation/AppKit/NSTableViewGridLineStyle/NSTableViewGridNone
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/gridStyleMask
func (t NSTableView) GridStyleMask() NSTableViewGridLineStyle {
	rv := objc.Send[NSTableViewGridLineStyle](t.ID, objc.Sel("gridStyleMask"))
	return NSTableViewGridLineStyle(rv)
}
func (t NSTableView) SetGridStyleMask(value NSTableViewGridLineStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setGridStyleMask:"), value)
}
// The effective row size style for the table.
//
// # Discussion
// 
// If the value in the [RowSizeStyle] property is
// [NSTableView.RowSizeStyle.default], then this property contains the default
// size for this table. The default size is currently set in System
// Preferences by the user.
//
// [NSTableView.RowSizeStyle.default]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/default
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/effectiveRowSizeStyle
func (t NSTableView) EffectiveRowSizeStyle() NSTableViewRowSizeStyle {
	rv := objc.Send[NSTableViewRowSizeStyle](t.ID, objc.Sel("effectiveRowSizeStyle"))
	return NSTableViewRowSizeStyle(rv)
}
// The row size style (small, medium, large, or custom) used by the table
// view.
//
// # Discussion
// 
// To set the row size style on a row by row basis, set the value of this
// property to [NSTableView.RowSizeStyle.custom] and implement the
// [TableViewHeightOfRow] method in your table view delegate object.
// 
// The default value of this property is [NSTableView.RowSizeStyle.custom],
// which tells the table to use the [RowHeight] of the table instead of any
// pre-determined system values. Generally, `rowSizeStyle` should always be
// [NSTableView.RowSizeStyle.custom] except for “source lists”.
//
// [NSTableView.RowSizeStyle.custom]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/custom
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/rowSizeStyle-swift.property
func (t NSTableView) RowSizeStyle() NSTableViewRowSizeStyle {
	rv := objc.Send[NSTableViewRowSizeStyle](t.ID, objc.Sel("rowSizeStyle"))
	return NSTableViewRowSizeStyle(rv)
}
func (t NSTableView) SetRowSizeStyle(value NSTableViewRowSizeStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setRowSizeStyle:"), value)
}
// An array containing the current table column objects.
//
// # Discussion
// 
// This property contains an array of [NSTableColumn] objects corresponding to
// the columns in the table. This array contains all columns, including those
// that are currently hidden.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/tableColumns
func (t NSTableView) TableColumns() []NSTableColumn {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("tableColumns"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTableColumn {
		return NSTableColumnFromID(id)
	})
}
// The index of the last selected column (or the last column added to the
// selection).
//
// # Discussion
// 
// When multiple columns are selected, this property contains only the index
// of the last one in the selection. If no column is selected, the value of
// this property is `-1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/selectedColumn
func (t NSTableView) SelectedColumn() int {
	rv := objc.Send[int](t.ID, objc.Sel("selectedColumn"))
	return rv
}
// An index set containing the indexes of the selected columns.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/selectedColumnIndexes
func (t NSTableView) SelectedColumnIndexes() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("selectedColumnIndexes"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}
// The number of selected columns.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfSelectedColumns
func (t NSTableView) NumberOfSelectedColumns() int {
	rv := objc.Send[int](t.ID, objc.Sel("numberOfSelectedColumns"))
	return rv
}
// The index of the last selected row (or the last row added to the
// selection).
//
// # Discussion
// 
// When multiple rows are selected, this property contains only the index of
// the last one in the selection. If no row is selected, the value of this
// property is `-1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/selectedRow
func (t NSTableView) SelectedRow() int {
	rv := objc.Send[int](t.ID, objc.Sel("selectedRow"))
	return rv
}
// An index set containing the indexes of the selected rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/selectedRowIndexes
func (t NSTableView) SelectedRowIndexes() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("selectedRowIndexes"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}
// The number of selected rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfSelectedRows
func (t NSTableView) NumberOfSelectedRows() int {
	rv := objc.Send[int](t.ID, objc.Sel("numberOfSelectedRows"))
	return rv
}
// A Boolean value indicating whether the table view allows the user to type
// characters to select rows.
//
// # Discussion
// 
// The default value of this property is [true]. Set it to [false] if you want
// to disable selecting rows by typing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/allowsTypeSelect
func (t NSTableView) AllowsTypeSelect() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsTypeSelect"))
	return rv
}
func (t NSTableView) SetAllowsTypeSelect(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsTypeSelect:"), value)
}
// The number of columns in the table.
//
// # Discussion
// 
// The value in this property includes table columns that are currently
// hidden.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfColumns
func (t NSTableView) NumberOfColumns() int {
	rv := objc.Send[int](t.ID, objc.Sel("numberOfColumns"))
	return rv
}
// The number of rows in the table.
//
// # Discussion
// 
// Typically you should not ask the table view how many rows it has; instead,
// interrogate the table view’s data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/numberOfRows
func (t NSTableView) NumberOfRows() int {
	rv := objc.Send[int](t.ID, objc.Sel("numberOfRows"))
	return rv
}
// A Boolean value indicating whether the table view draws grouped rows as if
// they are floating.
//
// # Discussion
// 
// Group rows are rows for which the table view delegate’s
// [TableViewIsGroupRow] method returns YES. These rows can be displayed as if
// they are floating in a view-based table view.
// 
// The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/floatsGroupRows
func (t NSTableView) FloatsGroupRows() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("floatsGroupRows"))
	return rv
}
func (t NSTableView) SetFloatsGroupRows(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setFloatsGroupRows:"), value)
}
// The index of the column being edited.
//
// # Return Value
// 
// If sent during [EditColumnRowWithEventSelect], the index in the
// [TableColumns] array of the column being edited; otherwise `–1`.
// 
// # Discussion
// 
// This property does not apply to view-based table views. In a view-based
// table view, the views are responsible for their own editing behavior. For
// other tables, the value reflects the index of the column being edited or
// `–1` when there is no editing session in progress or when the currently
// edited row is a “full width” row.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/editedColumn
func (t NSTableView) EditedColumn() int {
	rv := objc.Send[int](t.ID, objc.Sel("editedColumn"))
	return rv
}
// The index of the row being edited.
//
// # Discussion
// 
// This property does not apply to view-based table views. In a view-based
// table view, the views are responsible for their own editing behavior. For
// other tables, the value reflects the index of the row being edited or
// `–1` when there is no editing session in progress.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/editedRow
func (t NSTableView) EditedRow() int {
	rv := objc.Send[int](t.ID, objc.Sel("editedRow"))
	return rv
}
// The view object used to draw headers over columns.
//
// # Discussion
// 
// To configure a table without a header view or to remove the table view’s
// current header view, set the value of this property to `nil`. For more
// information about header views, see [NSTableHeaderView].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/headerView
func (t NSTableView) HeaderView() INSTableHeaderView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("headerView"))
	return NSTableHeaderViewFromID(objc.ID(rv))
}
func (t NSTableView) SetHeaderView(value INSTableHeaderView) {
	objc.Send[struct{}](t.ID, objc.Sel("setHeaderView:"), value)
}
// The view used to draw the area to the right of the column headers and above
// the vertical scroller of the enclosing scroll view.
//
// # Return Value
// 
// The view used to draw the area to the right of the column headers and above
// the vertical scroller of the enclosing [NSScrollView] object.
// 
// # Discussion
// 
// The default corner view draws a bezeled rectangle using a blank
// [NSTableHeaderCell] object, but you can replace it with a custom view that
// displays an image, or with a control that can handle mouse events, such as
// a Select All button. Your custom corner view should be as wide as a
// vertical [NSScroller] object and as tall as the header view of the table
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/cornerView
func (t NSTableView) CornerView() INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("cornerView"))
	return NSViewFromID(objc.ID(rv))
}
func (t NSTableView) SetCornerView(value INSView) {
	objc.Send[struct{}](t.ID, objc.Sel("setCornerView:"), value)
}
// The table view’s column autoresizing style.
//
// # Discussion
// 
// This property determines how columns are resized when the table view size
// changes. The default value of this property is
// [NSTableView.ColumnAutoresizingStyle.lastColumnOnlyAutoresizingStyle].
//
// [NSTableView.ColumnAutoresizingStyle.lastColumnOnlyAutoresizingStyle]: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum/lastColumnOnlyAutoresizingStyle
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/columnAutoresizingStyle-swift.property
func (t NSTableView) ColumnAutoresizingStyle() NSTableViewColumnAutoresizingStyle {
	rv := objc.Send[NSTableViewColumnAutoresizingStyle](t.ID, objc.Sel("columnAutoresizingStyle"))
	return NSTableViewColumnAutoresizingStyle(rv)
}
func (t NSTableView) SetColumnAutoresizingStyle(value NSTableViewColumnAutoresizingStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setColumnAutoresizingStyle:"), value)
}
// A Boolean value indicating whether the order and width of the table
// view’s columns are automatically saved.
//
// # Discussion
// 
// When this property is set to [true], the table information is saved
// separately for each user and application under the name specified in the
// [AutosaveName] property. If you change the value of this property from
// [false] to [true], the table tries to read in any saved information and
// sets the order and width of this table view’s columns to match. If the
// [AutosaveName] property is `nil`, this setting is ignored and the table
// information is not read or saved.
// 
// When autosave is enabled, the table saves the table column width, the table
// column order, any applied sort descriptors, and the table column hidden
// state (in macOS 10.5 and later).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/autosaveTableColumns
func (t NSTableView) AutosaveTableColumns() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("autosaveTableColumns"))
	return rv
}
func (t NSTableView) SetAutosaveTableColumns(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutosaveTableColumns:"), value)
}
// The name under which table information is automatically saved.
//
// # Discussion
// 
// The table information is saved separately in user defaults for each user
// and for each application that user uses. If no name has been set, the value
// of this property is `nil`. Even when a table view has an autosave name, it
// only saves the table information when the [AutosaveTableColumns] property
// is [true].
// 
// If you change the value of this property to a new name, the table reads in
// any saved information and sets the order and width of this table view’s
// columns to match. Setting the name to `nil` removes any previously stored
// state from the user defaults.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/autosaveName-swift.property
func (t NSTableView) AutosaveName() NSTableViewAutosaveName {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("autosaveName"))
	return NSTableViewAutosaveName(foundation.NSStringFromID(rv).String())
}
func (t NSTableView) SetAutosaveName(value NSTableViewAutosaveName) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutosaveName:"), objc.String(string(value)))
}
// The table view’s delegate.
//
// # Discussion
// 
// The delegate must conform to the [NSTableViewDelegate] protocol. Setting
// the delegate will implicitly reload the table view. Note that in versions
// of macOS prior to v10.12, the table view did not retain the delegate in a
// managed memory environment.
// 
// # Special Considerations
// 
// When you set the table view’s delegate, it is automatically registered
// for the following notifications with the following delegate methods:
// 
// - The notification named [selectionDidChangeNotification] is configured to
// notify the delegate’s [TableViewSelectionDidChange]. - The notification
// named [columnDidMoveNotification] is configured to notify the delegate’s
// [TableViewColumnDidMove]. - The notification named
// [columnDidResizeNotification] is configured to notify the delegate’s
// [TableViewColumnDidResize]. - The notification named
// [selectionIsChangingNotification] is configured to notify the delegate’s
// [TableViewSelectionIsChanging].
// 
// Setting the delegate to `nil` causes these notifications to be
// disconnected. Rather than setting the delegate to `nil` and listening for
// notifications (and expecting [NSTableView] to still function correctly) you
// should instead implement the appropriate delegate method.
//
// [columnDidMoveNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/columnDidMoveNotification
// [columnDidResizeNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/columnDidResizeNotification
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionDidChangeNotification
// [selectionIsChangingNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionIsChangingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/delegate
func (t NSTableView) Delegate() NSTableViewDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSTableViewDelegateObjectFromID(rv)
}
func (t NSTableView) SetDelegate(value NSTableViewDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}
// The column highlighted in the table.
//
// # Discussion
// 
// Assigning a value to this property highlights the specified column. A
// highlightable column header can be used in conjunction with row selection
// to highlight a particular column of the table. An example of this is how
// the Mail application indicates the currently sorted column.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/highlightedTableColumn
func (t NSTableView) HighlightedTableColumn() INSTableColumn {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("highlightedTableColumn"))
	return NSTableColumnFromID(objc.ID(rv))
}
func (t NSTableView) SetHighlightedTableColumn(value INSTableColumn) {
	objc.Send[struct{}](t.ID, objc.Sel("setHighlightedTableColumn:"), value)
}
// A Boolean value indicating whether vertical motion is treated as a drag or
// selection change.
//
// # Discussion
// 
// The default value of this property is [true], which indicates that a
// vertical drag motion begins a drag. In this case a vertical drag will
// drag-select rows. Most often, you would want to disable vertical dragging
// when it’s expected that horizontal dragging is the natural motion.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/verticalMotionCanBeginDrag
func (t NSTableView) VerticalMotionCanBeginDrag() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("verticalMotionCanBeginDrag"))
	return rv
}
func (t NSTableView) SetVerticalMotionCanBeginDrag(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setVerticalMotionCanBeginDrag:"), value)
}
// The feedback style displayed when the user drags over the table view.
//
// # Discussion
// 
// The default value of this property is
// [NSTableView.DraggingDestinationFeedbackStyle.regular]. However, changing
// the selection highlight style to
// [NSTableView.SelectionHighlightStyle.sourceList] automatically changes the
// value of this property to
// [NSTableView.DraggingDestinationFeedbackStyle.sourceList].
//
// [NSTableView.DraggingDestinationFeedbackStyle.regular]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum/regular
// [NSTableView.DraggingDestinationFeedbackStyle.sourceList]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum/sourceList
// [NSTableView.SelectionHighlightStyle.sourceList]: https://developer.apple.com/documentation/AppKit/NSTableView/SelectionHighlightStyle-swift.enum/sourceList
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/draggingDestinationFeedbackStyle-swift.property
func (t NSTableView) DraggingDestinationFeedbackStyle() NSTableViewDraggingDestinationFeedbackStyle {
	rv := objc.Send[NSTableViewDraggingDestinationFeedbackStyle](t.ID, objc.Sel("draggingDestinationFeedbackStyle"))
	return NSTableViewDraggingDestinationFeedbackStyle(rv)
}
func (t NSTableView) SetDraggingDestinationFeedbackStyle(value NSTableViewDraggingDestinationFeedbackStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}
// The table view’s sort descriptors.
//
// # Discussion
// 
// This property contains an array of [NSSortDescriptor] objects. A table
// column is considered sortable if it has a sort descriptor that specifies
// the sorting direction, a key to sort by, and a selector defining how to
// sort. Changing the value of this property may have the side effect of
// calling the [TableViewSortDescriptorsDidChange] method on the table
// view’s data source.
// 
// The contents of this property are archived and persisted along with other
// column information if autosave is enabled for the table.
//
// [NSSortDescriptor]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/sortDescriptors
func (t NSTableView) SortDescriptors() []foundation.NSSortDescriptor {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("sortDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSSortDescriptor {
		return foundation.NSSortDescriptorFromID(id)
	})
}
func (t NSTableView) SetSortDescriptors(value []foundation.NSSortDescriptor) {
	objc.Send[struct{}](t.ID, objc.Sel("setSortDescriptors:"), objectivec.IObjectSliceToNSArray(value))
}
// A Boolean value indicating whether a table row’s actions are visible.
//
// # Discussion
// 
// This property contains a Boolean value indicating whether a table row’s
// actions are visible or not—the user has swiped the row to reveal the row
// actions. Set the value of this property to [false] to hide any visible row
// actions. Setting the value of this property to [true] is not supported, and
// will result in an exception.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/rowActionsVisible
func (t NSTableView) RowActionsVisible() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("rowActionsVisible"))
	return rv
}
func (t NSTableView) SetRowActionsVisible(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setRowActionsVisible:"), value)
}
// The indexes of all hidden table rows.
//
// # Discussion
// 
// The value of this property is an index set containing the indexes of any
// hidden table rows. Table rows may be hidden by invoking the
// [HideRowsAtIndexesWithAnimation] method. Some drag-and-drop operations also
// result in hidden rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/hiddenRowIndexes
func (t NSTableView) HiddenRowIndexes() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("hiddenRowIndexes"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// See: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (t NSTableView) IsEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("enabled"))
	return rv
}
func (t NSTableView) SetIsEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setEnabled:"), value)
}

// Sent when a control with editable cells begins an edit session.
//
// See: https://developer.apple.com/documentation/appkit/nscontrol/textdidbegineditingnotification
func (_NSTableViewClass NSTableViewClass) TextDidBeginEditingNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSTableViewClass.class), objc.Sel("NSControlTextDidBeginEditingNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}
// Sent when the text in the receiving control changes.
//
// See: https://developer.apple.com/documentation/appkit/nscontrol/textdidchangenotification
func (_NSTableViewClass NSTableViewClass) TextDidChangeNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSTableViewClass.class), objc.Sel("NSControlTextDidChangeNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}

			// Protocol methods for NSAccessibilityTable
			
// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
// 
// The element’s frame in screen coordinates.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
func (o NSTableView) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
	}
// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
// 
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityParent] property.
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
func (o NSTableView) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
	}
// Returns the accessibility element’s identity.
//
// # Return Value
// 
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIdentifier] property.
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
func (o NSTableView) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}
// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
// 
// [true] if this element has the keyboard focus; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
func (o NSTableView) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

			// Protocol methods for NSDraggingSource
			

			// Protocol methods for NSTextViewDelegate
			

			// Protocol methods for NSUserInterfaceValidations
			


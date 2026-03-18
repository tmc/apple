// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A set of optional methods you implement in a table view delegate to customize the behavior of the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate
type NSTableViewDelegate interface {
	objectivec.IObject
	NSControlTextEditingDelegate
}

// NSTableViewDelegateObject wraps an existing Objective-C object that conforms to the NSTableViewDelegate protocol.
type NSTableViewDelegateObject struct {
	objectivec.Object
}
func (o NSTableViewDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTableViewDelegateObjectFromID constructs a [NSTableViewDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTableViewDelegateObjectFromID(id objc.ID) NSTableViewDelegateObject {
	return NSTableViewDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate for a view to display the specified row and column.
//
// tableView: The table view that sent the message.
//
// tableColumn: The table column. (If the row is a group row, `tableColumn` is `nil`.)
//
// row: The row index.
//
// # Return Value
// 
// The view to display the specified column and row. If you return `nil`, a
// view will not be shown at that location.
//
// # Discussion
// 
// This method is required if you want to use [NSView] objects instead of
// [NSCell] objects for the cells within a table view. Cells and views can’t
// be mixed within the same table view.
// 
// It’s recommended that the implementation of this method first call the
// [NSTableView] method [ViewWithIdentifierOwner] passing, respectively, the
// `tableColumn` parameter’s identifier and `self` as the owner to attempt
// to reuse a view that is no longer visible or automatically unarchive an
// associated prototype view for that identifier. The `frame` of the view
// returned by this method is not important, and it will be automatically set
// by the table.
// 
// The view’s properties should be properly set up before returning the
// result.
// 
// When using Cocoa bindings, this method is optional if at least one
// identifier has been associated with the table view at design time. (Note
// that a view’s identifier must be the same as the identifier of its
// column. An easy way to achieve this is to use the Automatic identifier
// setting in Interface Builder.) If this method isn’t implemented, the
// table will automatically call the [NSTableView] method
// [ViewWithIdentifierOwner] with the `tableColumn` parameter’s identifier
// and the table view’s `delegate` as parameters, to attempt to reuse a
// previous view, or automatically unarchive a prototype associated with the
// table view. If this method is implemented, you can set up properties that
// aren’t using bindings.
// 
// The autoresizing mask of the returned view will automatically be set to
// [height] to resize properly on row height changes.
//
// [height]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct/height
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:viewFor:row:)

func (o NSTableViewDelegateObject) TableViewViewForTableColumnRow(tableView INSTableView, tableColumn INSTableColumn, row int) INSView {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tableView:viewForTableColumn:row:"), tableView, tableColumn, row)
	return NSViewFromID(rv)
	}

// Asks the delegate for a view to display the specified row.
//
// tableView: The table view that sent the message.
//
// row: The row index.
//
// # Return Value
// 
// An instance or subclass of [NSTableRowView]. If `nil` is returned, an
// [NSTableRowView] instance will be created and used.
//
// # Discussion
// 
// The delegate can implement this method to return a custom [NSTableRowView]
// for `row`.
// 
// The reuse queue can be used in the same way as documented in
// [TableViewViewForTableColumnRow]. The returned view will have attributes
// properly set to it before it’s added to the `tableView`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:rowViewForRow:)

func (o NSTableViewDelegateObject) TableViewRowViewForRow(tableView INSTableView, row int) INSTableRowView {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tableView:rowViewForRow:"), tableView, row)
	return NSTableRowViewFromID(rv)
	}

// Tells the delegate that a row view was added at the specified row.
//
// tableView: The table view that sent the message.
//
// rowView: The row view.
//
// row: The index of the row.
//
// # Discussion
// 
// At this point, the delegate can add extra views, or modify the properties
// of `rowView`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:didAdd:forRow:)

func (o NSTableViewDelegateObject) TableViewDidAddRowViewForRow(tableView INSTableView, rowView INSTableRowView, row int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:didAddRowView:forRow:"), tableView, rowView, row)
	}

// Tells the delegate that a row view was removed from the table at the
// specified row.
//
// tableView: The table view that sent the message.
//
// rowView: The row view.
//
// row: The index of the row.
//
// # Discussion
// 
// If `row` equals `-1`, the row is being deleted from the table and is no
// longer a valid row; otherwise `row` is a valid row that is being removed by
// being moved off screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:didRemove:forRow:)

func (o NSTableViewDelegateObject) TableViewDidRemoveRowViewForRow(tableView INSTableView, rowView INSTableRowView, row int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:didRemoveRowView:forRow:"), tableView, rowView, row)
	}

// Returns whether the specified row is a group row.
//
// tableView: The table view that sent the message.
//
// row: The row index.
//
// # Return Value
// 
// [true] if the specified row should have the group row style drawn, [false]
// otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the cell in `row` is an [NSTextFieldCell] object and contains only a
// string, the group row style attributes are automatically applied to the
// cell.
// 
// Group rows in [NSView]-based table views can be made to visually
// “float” by setting the table view method [FloatsGroupRows] to [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:isGroupRow:)

func (o NSTableViewDelegateObject) TableViewIsGroupRow(tableView INSTableView, row int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tableView:isGroupRow:"), tableView, row)
	return rv
	}

// Tells the delegate that the table view will display the specified cell at
// the specified row and column.
//
// tableView: The table view that sent the message.
//
// cell: The cell to be displayed.
//
// tableColumn: The table column.
//
// row: The row index.
//
// # Discussion
// 
// The delegate can modify the display attributes of `aCell` to alter the
// appearance of the cell.
// 
// Because `aCell` is reused for every row in `aTableColumn`, the delegate
// must set the display attributes both when drawing special cells and when
// drawing standard cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:willDisplayCell:for:row:)

func (o NSTableViewDelegateObject) TableViewWillDisplayCellForTableColumnRow(tableView INSTableView, cell objectivec.IObject, tableColumn INSTableColumn, row int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:willDisplayCell:forTableColumn:row:"), tableView, cell, tableColumn, row)
	}

// Asks the delegate for a custom data cell for the specified row and column.
//
// tableView: The table view that sent the message.
//
// tableColumn: The table column.
//
// row: The row index.
//
// # Return Value
// 
// An [NSCell] subclass that is used for the specified `row` and
// `tableColumn`. The returned cell must properly implement ``.
//
// # Discussion
// 
// A different data cell can be returned for any particular table column and
// row, or a cell that will be used for the entire row (that is, a full width
// cell).
// 
// If `tableColumn` is non-`nil`, you should return a cell (generally as the
// result of sending `tableColumn` a [DataCellForRow] message).
// 
// While each row is being drawn, this method is first called with a
// `tableColumn` value of `nil` to allow you to return a group cell—that is,
// a cell that will be used to draw the entire row. If you return a cell when
// `tableColumn` is `nil`, all implemented datasource and delegate methods
// must be prepared to handle a `nil` table column value. If you don’t
// return a cell, this method is called once for each `tableColumn` in
// `tableView`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:dataCellFor:row:)

func (o NSTableViewDelegateObject) TableViewDataCellForTableColumnRow(tableView INSTableView, tableColumn INSTableColumn, row int) INSCell {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tableView:dataCellForTableColumn:row:"), tableView, tableColumn, row)
	return NSCellFromID(rv)
	}

// Asks the delegate if an expansion tooltip should be displayed for a
// specific row and column.
//
// tableView: The table view that sent the message.
//
// tableColumn: The table column.
//
// row: The row index.
//
// # Return Value
// 
// [true] if an expansion tooltip should be displayed, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// An expansion tooltip can be displayed when the pointer hovers over a cell
// that contains truncated text. When this method returns [true], the cell’s
// full contents is shown in an expansion tooltip, which looks similar to a
// help tag.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:shouldShowCellExpansionFor:row:)

func (o NSTableViewDelegateObject) TableViewShouldShowCellExpansionForTableColumnRow(tableView INSTableView, tableColumn INSTableColumn, row int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tableView:shouldShowCellExpansionForTableColumn:row:"), tableView, tableColumn, row)
	return rv
	}

// Asks the delegate for a string to display in a tooltip for the specified
// cell in the column and row.
//
// tableView: The table view that sent the message.
//
// cell: The cell.
//
// rect: The proposed active area of the tooltip. You can modify `rect` to provide
// an alternative active area.
//
// tableColumn: The table column.
//
// row: The row index.
//
// mouseLocation: The mouse location.
//
// # Return Value
// 
// A string that should be displayed in the tooltip. Return `nil` or the empty
// string if no tooltip is desired.
//
// # Discussion
// 
// By default, `rect` is computed as
// 
// `[cell cellFrame]`. Note that tooltips are also known as help tags.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:toolTipFor:rect:tableColumn:row:mouseLocation:)

func (o NSTableViewDelegateObject) TableViewToolTipForCellRectTableColumnRowMouseLocation(tableView INSTableView, cell INSCell, rect foundation.NSRect, tableColumn INSTableColumn, row int, mouseLocation corefoundation.CGPoint) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tableView:toolTipForCell:rect:tableColumn:row:mouseLocation:"), tableView, cell, rect, tableColumn, row, mouseLocation)
	return foundation.NSStringFromID(rv).String()
	}

// Asks the delegate if the cell at the specified row and column can be
// edited.
//
// tableView: The table view that sent the message.
//
// tableColumn: The table column.
//
// row: The row index.
//
// # Return Value
// 
// [true] to allow editing the cell, [false] to deny editing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate can implement this method to disallow editing of specific
// cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:shouldEdit:row:)

func (o NSTableViewDelegateObject) TableViewShouldEditTableColumnRow(tableView INSTableView, tableColumn INSTableColumn, row int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tableView:shouldEditTableColumn:row:"), tableView, tableColumn, row)
	return rv
	}

// Asks the delegate for the height of the specified row.
//
// tableView: The table view that sent the message.
//
// row: The row index.
//
// # Return Value
// 
// The height of the row. The height doesn’t include intercell spacing and
// must be greater than zero.
//
// # Discussion
// 
// Implement this method if your table supports varying row heights.
// 
// Although table views may cache the returned values, you should ensure that
// this method is efficient. When you change a row’s height you must
// invalidate the existing row height by calling
// [NoteHeightOfRowsWithIndexesChanged]. [NSTableView] automatically
// invalidates its entire row height cache in response to calls to
// [ReloadData] or [NoteNumberOfRowsChanged].
// 
// If you call [ViewAtColumnRowMakeIfNecessary] or
// [RowViewAtRowMakeIfNecessary] within your implementation of this method,
// the table view throws an exception.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:heightOfRow:)

func (o NSTableViewDelegateObject) TableViewHeightOfRow(tableView INSTableView, row int) float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("tableView:heightOfRow:"), tableView, row)
	return rv
	}

// Asks the delegate to provide custom sizing behavior when a column’s
// resize divider is double clicked.
//
// tableView: The table view that sent the message.
//
// column: The index of the column.
//
// # Return Value
// 
// The width of the specified column.
//
// # Discussion
// 
// By default, [NSTableView] iterates every row in the table, accesses a cell
// via [preparedCell(atColumn:row:)], and requests the [CellSize] to find the
// appropriate largest width to use.
// 
// For accurate results and performance, it’s recommended that this method
// is implemented when using large tables. By default, large tables use a
// Monte Carlo simulation instead of iterating every row.
//
// [preparedCell(atColumn:row:)]: https://developer.apple.com/documentation/AppKit/NSTableView/preparedCell(atColumn:row:)
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:sizeToFitWidthOfColumn:)

func (o NSTableViewDelegateObject) TableViewSizeToFitWidthOfColumn(tableView INSTableView, column int) float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("tableView:sizeToFitWidthOfColumn:"), tableView, column)
	return rv
	}

// Asks the delegate if the user is allowed to change the selection.
//
// tableView: The table view that sent the message.
//
// # Return Value
// 
// [true] to allow the table view to change its selection (typically a row
// being edited), [false] to deny selection change.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The user can select and edit different cells within the same row, but
// can’t select another row unless the delegate approves. The delegate can
// implement this method for complex validation of edited rows based on the
// values of any of their cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/selectionShouldChange(in:)

func (o NSTableViewDelegateObject) SelectionShouldChangeInTableView(tableView INSTableView) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("selectionShouldChangeInTableView:"), tableView)
	return rv
	}

// Asks the delegate if the table view should allow selection of the specified
// row.
//
// tableView: The table view that sent the message.
//
// row: The row index.
//
// # Return Value
// 
// [true] to permit selection of the row, [false] to deny selection.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate can implement this method to disallow selection of particular
// rows.
// 
// For better performance and finer-grain control over the selection, use
// [TableViewSelectionIndexesForProposedSelection].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:shouldSelectRow:)

func (o NSTableViewDelegateObject) TableViewShouldSelectRow(tableView INSTableView, row int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tableView:shouldSelectRow:"), tableView, row)
	return rv
	}

// Asks the delegate to accept or reject the proposed selection.
//
// tableView: The table view that sent the message.
//
// proposedSelectionIndexes: An index set containing the indexes of the proposed selection.
//
// # Return Value
// 
// An [NSIndexSet] instance containing the indexes of the new selection.
// Return `proposedSelectionIndexes` if the proposed selection is acceptable,
// or the value of the table view’s existing selection to avoid changing the
// selection.
//
// [NSIndexSet]: https://developer.apple.com/documentation/Foundation/NSIndexSet
//
// # Discussion
// 
// This method may be called multiple times with one new index added to the
// existing selection to find out if a particular index can be selected when
// the user is extending the selection with the keyboard or mouse.
// 
// If implemented, this method will be called instead of
// [TableViewShouldSelectRow].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:selectionIndexesForProposedSelection:)

func (o NSTableViewDelegateObject) TableViewSelectionIndexesForProposedSelection(tableView INSTableView, proposedSelectionIndexes foundation.NSIndexSet) foundation.NSIndexSet {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tableView:selectionIndexesForProposedSelection:"), tableView, proposedSelectionIndexes)
	return foundation.NSIndexSetFromID(rv)
	}

// Asks the delegate whether the specified table column can be selected.
//
// tableView: The table view that sent the message.
//
// tableColumn: The table column.
//
// # Return Value
// 
// [true] to permit selection, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate can implement this method to disallow selection of particular
// columns.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:shouldSelect:)

func (o NSTableViewDelegateObject) TableViewShouldSelectTableColumn(tableView INSTableView, tableColumn INSTableColumn) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tableView:shouldSelectTableColumn:"), tableView, tableColumn)
	return rv
	}

// Tells the delegate that the table view’s selection is in the process of
// changing.
//
// notification: A notification named [selectionIsChangingNotification].
// //
// [selectionIsChangingNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionIsChangingNotification
//
// # Discussion
// 
// This method is called only when mouse events—and not keyboard
// events—are changing the selection. For example, this method is called
// when the user drags the mouse across multiple table rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableViewSelectionIsChanging(_:)

func (o NSTableViewDelegateObject) TableViewSelectionIsChanging(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableViewSelectionIsChanging:"), notification)
	}

// Tells the delegate that the table view’s selection has changed.
//
// notification: A notification named [selectionDidChangeNotification].
// //
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/selectionDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableViewSelectionDidChange(_:)

func (o NSTableViewDelegateObject) TableViewSelectionDidChange(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableViewSelectionDidChange:"), notification)
	}

// Asks the delegate to allow or deny type select for the specified event and
// current search string.
//
// tableView: The table view that sent the message.
//
// event: The event.
//
// searchString: The search string or `nil` if no type select has began.
//
// # Return Value
// 
// [true] to allow type select for event, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Typically, this is called from the [NSTableView] `` implementation and the
// event will be a key event.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:shouldTypeSelectFor:withCurrentSearch:)

func (o NSTableViewDelegateObject) TableViewShouldTypeSelectForEventWithCurrentSearchString(tableView INSTableView, event INSEvent, searchString string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tableView:shouldTypeSelectForEvent:withCurrentSearchString:"), tableView, event, objc.String(searchString))
	return rv
	}

// Asks the delegate to provide an alternative text value used for type
// selection for the specified row and column.
//
// tableView: The table view that sent the message.
//
// tableColumn: The table column.
//
// row: The row index.
//
// # Return Value
// 
// A string that is used in type select comparison for `row` and
// `tableColumn`. Return `nil` if `row` or `tableColumn` should not be
// searched.
//
// # Discussion
// 
// Implement this method to change the string value that is searched for based
// on what is displayed. By default, all cells with text in them are searched.
// 
// If this delegate method isn’t implemented, the default string value
// (which can also be returned from the delegate method) is:
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:typeSelectStringFor:row:)

func (o NSTableViewDelegateObject) TableViewTypeSelectStringForTableColumnRow(tableView INSTableView, tableColumn INSTableColumn, row int) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tableView:typeSelectStringForTableColumn:row:"), tableView, tableColumn, row)
	return foundation.NSStringFromID(rv).String()
	}

// Asks the delegate for the row within the specified search range that
// matches the specified string.
//
// tableView: The table view that sent the message.
//
// startRow: The starting row of the search range.
//
// endRow: The ending row of the search range.
//
// searchString: A string containing the typed selection.
//
// # Return Value
// 
// The first row in the range of `startRow` through `endRow` (excluding
// `endRow` itself) that matches `selectionString`. Return `-1` if no match is
// found.
//
// # Discussion
// 
// Use this method to control how type selection works in a table.
// (Implementation of this method isn’t required to support type selection.)
// Note that it’s possible for `endRow` to be less than `startRow` if the
// search will wrap.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:nextTypeSelectMatchFromRow:toRow:for:)

func (o NSTableViewDelegateObject) TableViewNextTypeSelectMatchFromRowToRowForString(tableView INSTableView, startRow int, endRow int, searchString string) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("tableView:nextTypeSelectMatchFromRow:toRow:forString:"), tableView, startRow, endRow, objc.String(searchString))
	return rv
	}

// Asks the delegate to allow or prohibit the specified column to be dragged
// to a new location.
//
// tableView: The table view that sent the message.
//
// columnIndex: The index of the column being dragged.
//
// newColumnIndex: The proposed target index of the column.
//
// # Return Value
// 
// [true] if the column reordering should be allowed, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// When a column is initially dragged by the user, the delegate is first
// called with a `newColumnIndex` value of `-1`. Returning [false] will
// disallow that column from being reordered at all. Returning [true] allows
// it to be reordered, and the delegate will be called again when the column
// reaches a new location.
// 
// The actual [NSTableColumn] instance can be retrieved from the
// [TableColumns] array.
// 
// If this method isn’t implemented, all columns are considered reorderable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:shouldReorderColumn:toColumn:)

func (o NSTableViewDelegateObject) TableViewShouldReorderColumnToColumn(tableView INSTableView, columnIndex int, newColumnIndex int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tableView:shouldReorderColumn:toColumn:"), tableView, columnIndex, newColumnIndex)
	return rv
	}

// Tells the delegate that the specified table column was dragged.
//
// tableView: The table view that sent the message.
//
// tableColumn: The table column.
//
// # Discussion
// 
// Specifically, this method is sent when the mouse button goes up in
// `tableView` and `tableColumn` has been dragged during the time the mouse
// button was down. In macOS 10.5 and later the dragged column is sent to the
// delegate. (In earlier versions of macOS the table column that’s currently
// located at the dragged column’s original index is sent.)
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:didDrag:)

func (o NSTableViewDelegateObject) TableViewDidDragTableColumn(tableView INSTableView, tableColumn INSTableColumn) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:didDragTableColumn:"), tableView, tableColumn)
	}

// Tells the delegate that a table column was moved by user action.
//
// notification: A notification named [columnDidMoveNotification].
// //
// [columnDidMoveNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/columnDidMoveNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableViewColumnDidMove(_:)

func (o NSTableViewDelegateObject) TableViewColumnDidMove(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableViewColumnDidMove:"), notification)
	}

// Tells the delegate that a table column was resized.
//
// notification: A notification named [columnDidResizeNotification].
// //
// [columnDidResizeNotification]: https://developer.apple.com/documentation/AppKit/NSTableView/columnDidResizeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableViewColumnDidResize(_:)

func (o NSTableViewDelegateObject) TableViewColumnDidResize(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableViewColumnDidResize:"), notification)
	}

// Tells the delegate that the mouse button was clicked in the specified table
// column, but the column was not dragged.
//
// tableView: The table view that sent the message.
//
// tableColumn: The table column.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:didClick:)

func (o NSTableViewDelegateObject) TableViewDidClickTableColumn(tableView INSTableView, tableColumn INSTableColumn) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:didClickTableColumn:"), tableView, tableColumn)
	}

// Tells the delegate that the mouse button was clicked in the specified table
// column’s header.
//
// tableView: The table view that sent the message.
//
// tableColumn: The table column.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:mouseDownInHeaderOf:)

func (o NSTableViewDelegateObject) TableViewMouseDownInHeaderOfTableColumn(tableView INSTableView, tableColumn INSTableColumn) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:mouseDownInHeaderOfTableColumn:"), tableView, tableColumn)
	}

// Asks the delegate whether the specified cell should be tracked.
//
// tableView: The table view that sent the message.
//
// cell: The cell to track.
//
// tableColumn: The table column.
//
// row: A row in `tableView`.
//
// # Return Value
// 
// [true] if the cell should be tracked, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// In general, only selectable or selected cells can be tracked. If you
// implement this method, cells that aren’t selectable or selected can be
// tracked; similarly, cells that are selectable or selected can be set as
// untracked.
// 
// For example, this allows you to have an [NSButtonCell] object in a table
// that doesn’t change the selection, but can still be clicked on and
// tracked.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:shouldTrackCell:for:row:)

func (o NSTableViewDelegateObject) TableViewShouldTrackCellForTableColumnRow(tableView INSTableView, cell INSCell, tableColumn INSTableColumn, row int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tableView:shouldTrackCell:forTableColumn:row:"), tableView, cell, tableColumn, row)
	return rv
	}

// Asks the delegate to provide an array of row actions to be attached to the
// specified edge of a table row and displayed when the user swipes
// horizontally across the row.
//
// tableView: The table view that sent the message.
//
// row: The index of the target row.
//
// edge: The edge (of class [NSTableView.RowActionEdge]) for which row actions are
// requested. This is based on the direction in which the user swiped on the
// row. Swiping to the right results in an edge value of
// [NSTableView.RowActionEdge.leading]. Swiping to the left results in an edge
// value of [NSTableView.RowActionEdge.trailing].
// //
// [NSTableView.RowActionEdge.leading]: https://developer.apple.com/documentation/AppKit/NSTableView/RowActionEdge/leading
// [NSTableView.RowActionEdge.trailing]: https://developer.apple.com/documentation/AppKit/NSTableView/RowActionEdge/trailing
// [NSTableView.RowActionEdge]: https://developer.apple.com/documentation/AppKit/NSTableView/RowActionEdge
//
// # Return Value
// 
// An array of row actions (of class [NSTableViewRowAction]) to be enabled on
// the specified edge of the table row.
//
// # Discussion
// 
// Implement this method if your table row supports actions that are displayed
// when the user swipes horizontally across the row. For example, your table
// view could use this method to implement a swipe left to delete function in
// your table rows. When called, this method receives the table view, the
// index of the row the user swiped, and an edge of type
// [NSTableRowActionEdge]. The method should return an array of any row
// actions of class [NSTableViewRowAction] that are supported for the
// specified edge. If no row actions are available, an empty array should be
// returned.
// 
// If this method isn’t implemented, then the table row displays no actions
// when the user swipes horizontally away from the specified edge.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:rowActionsForRow:edge:)

func (o NSTableViewDelegateObject) TableViewRowActionsForRowEdge(tableView INSTableView, row int, edge NSTableRowActionEdge) []NSTableViewRowAction {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("tableView:rowActionsForRow:edge:"), tableView, row, edge)
	return objc.ConvertSlice(rv, func(id objc.ID) NSTableViewRowAction {
		return NSTableViewRowActionFromID(id)
	})
	}

// Asks the delegate to verify that the user can change the given column’s
// visibility.
//
// tableView: The table view object requesting this information.
//
// column: The table column affected by the visibility change.
//
// # Return Value
// 
// [true] if the user can change the visibility of the column; otherwise,
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Implement this method to enable the table view to provide a menu that
// allows users to show or hide table columns.
// 
// To change a column’s visibility, ensure the column has a title. Further,
// setting the [Menu] property on the table view [HeaderView] or, subclassing
// [NSTableHeaderView] and overriding the [Menu] property also prevents the
// column visibility from changing.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:userCanChangeVisibilityOf:)

func (o NSTableViewDelegateObject) TableViewUserCanChangeVisibilityOfTableColumn(tableView INSTableView, column INSTableColumn) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tableView:userCanChangeVisibilityOfTableColumn:"), tableView, column)
	return rv
	}

// Tells the delegate that the user changed the visibility of one or more
// table columns.
//
// tableView: The table view object requesting this information.
//
// columns: The table columns affected by the visibility change.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDelegate/tableView(_:userDidChangeVisibilityOf:)

func (o NSTableViewDelegateObject) TableViewUserDidChangeVisibilityOfTableColumns(tableView INSTableView, columns []NSTableColumn) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:userDidChangeVisibilityOfTableColumns:"), tableView, objectivec.IObjectSliceToNSArray(columns))
	}

// Invoked when the insertion point leaves a cell belonging to the specified
// control, but before the value of the cell’s object is displayed.
//
// control: The control whose object value needs to be validated.
//
// obj: The object value to validate.
//
// # Return Value
// 
// [true] if you want to allow the control to display the specified value;
// otherwise, [false] to reject the value and return the cursor to the
// control’s cell.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method gives the delegate the opportunity to validate the contents of
// the control’s cell (or selected cell). In validating, the delegate should
// check the value in the `object` parameter and determine if it falls within
// a permissible range, has required attributes, accords with a given context,
// and so on. Examples of objects subject to such evaluations are an [NSDate]
// object that should not represent a future date or a monetary amount
// (represented by an [NSNumber] object) that exceeds a predetermined limit.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:isValidObject:)

func (o NSTableViewDelegateObject) ControlIsValidObject(control INSControl, obj objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("control:isValidObject:"), control, obj)
	return rv
	}

// Invoked when the formatter for the cell belonging to `control` (or selected
// cell) rejects a partial string a user is typing into the cell.
//
// control: The control whose cell rejected the string.
//
// string: The string that includes the character that caused the rejection.
//
// error: A localized, user-presentable string that explains why the string was
// rejected.
//
// # Discussion
// 
// You can implement this method to display a warning message or perform a
// similar action when the user enters improperly formatted text.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:didFailToValidatePartialString:errorDescription:)

func (o NSTableViewDelegateObject) ControlDidFailToValidatePartialStringErrorDescription(control INSControl, string_ string, error_ string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("control:didFailToValidatePartialString:errorDescription:"), control, objc.String(string_), objc.String(error_))
	}

// Invoked when the formatter for the cell belonging to the specified control
// cannot convert a string to an underlying object.
//
// control: The control whose cell could not convert the string.
//
// string: The string that could not be converted.
//
// error: A localized, user-presentable string that explains why the conversion
// failed.
//
// # Return Value
// 
// [true] if the value in the string parameter should be accepted as is;
// otherwise, [false] if the value in the parameter should be rejected.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Your implementation of this method should evaluate the error or query the
// user an appropriate value indicating whether the string should be accepted
// or rejected.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:didFailToFormatString:errorDescription:)

func (o NSTableViewDelegateObject) ControlDidFailToFormatStringErrorDescription(control INSControl, string_ string, error_ string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("control:didFailToFormatString:errorDescription:"), control, objc.String(string_), objc.String(error_))
	return rv
	}

// Invoked when the user tries to enter a character in a cell of a control
// that allows editing of text (such as a text field or form field).
//
// control: The control whose content is about to be edited.
//
// fieldEditor: The field editor of the control.
//
// # Return Value
// 
// [true] if the control’s field editor should be allowed to start editing
// the text; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You can use this method to allow or disallow editing in a control. This
// message is sent by the control directly to its delegate object.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:textShouldBeginEditing:)

func (o NSTableViewDelegateObject) ControlTextShouldBeginEditing(control INSControl, fieldEditor INSText) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("control:textShouldBeginEditing:"), control, fieldEditor)
	return rv
	}

// Invoked when the insertion point tries to leave a cell of the control that
// has been edited.
//
// control: The control for which editing is about to end.
//
// fieldEditor: The field editor of the control. You can use this parameter to get the
// edited text.
//
// # Return Value
// 
// [true] if the insertion point should be allowed to end the editing session;
// otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This message is sent only by controls that allow editing of text (such as a
// text field or a form field). This message is sent by the control directly
// to its delegate object.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:textShouldEndEditing:)

func (o NSTableViewDelegateObject) ControlTextShouldEndEditing(control INSControl, fieldEditor INSText) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("control:textShouldEndEditing:"), control, fieldEditor)
	return rv
	}

// Invoked to allow you to control the list of proposed text completions
// generated by text fields and other controls.
//
// control: The control whose cell initiated the message. If the control contains
// multiple cells, the one that initiated the message is usually the selected
// cell.
//
// textView: The field editor of the control. You can use this parameter to get the
// typed text.
//
// words: An array of [NSString] objects containing the potential completions. The
// completion strings are listed in their order of preference in the array.
//
// charRange: The range of characters the user has already typed.
//
// index: On input, an integer variable with the default value of `0`. On output, you
// can set this value to an index in the returned array indicating which item
// should be selected initially. Set the value to `-1` to indicate there
// should not be an initial selection.
//
// # Return Value
// 
// An array of [NSString] objects containing the list of completions to use in
// place of the array in the `words` parameter. The returned array should list
// the completions in their preferred order
//
// # Discussion
// 
// Each string you return should be a complete word that the user might be
// trying to type. The strings must be complete words rather than just the
// remainder of the word, in case completion requires some slight modification
// of what the user has already typed—for example, the addition of an
// accent, or a change in capitalization. You can also use this method to
// support abbreviations that complete into words that do not start with the
// characters of the abbreviation. The `index` argument allows you to return
// by reference an index specifying which of the completions should be
// selected initially.
// 
// The actual means of presentation of the potential completions is determined
// by the [Complete] method of [NSTextView].
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:textView:completions:forPartialWordRange:indexOfSelectedItem:)

func (o NSTableViewDelegateObject) ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control INSControl, textView INSTextView, words []string, charRange foundation.NSRange, index unsafe.Pointer) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), control, textView, objectivec.StringSliceToNSArray(words), charRange, index)
	return objc.ConvertSliceToStrings(rv)
	}

// Invoked when users press keys with predefined bindings in a cell of the
// specified control.
//
// control: The control whose cell initiated the message. If the control contains
// multiple cells, the one that initiated the message is usually the selected
// cell.
//
// textView: The field editor of the control.
//
// commandSelector: The selector that was associated with the binding.
//
// # Return Value
// 
// [true] if the delegate object handles the key binding; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// These bindings are usually implemented as methods (`command`) defined in
// the [NSResponder] class; examples of such key bindings are arrow keys (for
// directional movement) and the Escape key (for name completion). By
// implementing this method, the delegate can override the default
// implementation of `command` and supply its own behavior.
// 
// For example, the default method for completing partially typed pathnames or
// symbols (usually when users press the Escape key) is `complete(_:)`. The
// default implementation of the `complete(_:)` method (in [NSResponder]) does
// nothing. The delegate could evaluate the method in the `command` parameter
// and, if it’s `complete(_:)`, get the current string from the `textView`
// parameter and then expand it, or display a list of potential completions,
// or do whatever else is appropriate.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:textView:doCommandBy:)

func (o NSTableViewDelegateObject) ControlTextViewDoCommandBySelector(control INSControl, textView INSTextView, commandSelector objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("control:textView:doCommandBySelector:"), control, textView, commandSelector)
	return rv
	}

// Tells the delegate that the control started editing its text content.
//
// obj: A notification object that contains details about the editing
// configuration.
//
// # Discussion
// 
// Use the key `“NSFieldEditor”` to obtain the field editor from the
// notification object’s `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/controlTextDidBeginEditing(_:)

func (o NSTableViewDelegateObject) ControlTextDidBeginEditing(obj foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("controlTextDidBeginEditing:"), obj)
	}

// Tells the delegate that the control made changes to its text content.
//
// obj: A notification object that contains details about the editing
// configuration.
//
// # Discussion
// 
// Use the key `“NSFieldEditor”` to obtain the field editor from the
// notification object’s `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/controlTextDidChange(_:)

func (o NSTableViewDelegateObject) ControlTextDidChange(obj foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("controlTextDidChange:"), obj)
	}

// Tells the delegate that the control finished editing its text content and
// committed the changes.
//
// obj: A notification object that contains details about the editing
// configuration.
//
// # Discussion
// 
// Use the key `“NSFieldEditor”` to obtain the field editor from the
// notification object’s `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/controlTextDidEndEditing(_:)

func (o NSTableViewDelegateObject) ControlTextDidEndEditing(obj foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("controlTextDidEndEditing:"), obj)
	}

// NSTableViewDelegateConfig holds optional typed callbacks for [NSTableViewDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstableviewdelegate
type NSTableViewDelegateConfig struct {

	// Providing views for rows and columns
	// RowViewForRow — Asks the delegate for a view to display the specified row.
	RowViewForRow func(tableView NSTableView, row int) NSTableRowView

	// Grouping rows
	// IsGroupRow — Returns whether the specified row is a group row.
	IsGroupRow func(tableView NSTableView, row int) bool

	// Setting row and column size
	// HeightOfRow — Asks the delegate for the height of the specified row.
	HeightOfRow func(tableView NSTableView, row int) float64
	// SizeToFitWidthOfColumn — Asks the delegate to provide custom sizing behavior when a column’s resize divider is double clicked.
	SizeToFitWidthOfColumn func(tableView NSTableView, column int) float64

	// Selecting rows
	// ShouldSelectRow — Asks the delegate if the table view should allow selection of the specified row.
	ShouldSelectRow func(tableView NSTableView, row int) bool
	// SelectionIndexesForProposedSelection — Asks the delegate to accept or reject the proposed selection.
	SelectionIndexesForProposedSelection func(tableView NSTableView, proposedSelectionIndexes foundation.NSIndexSet) foundation.NSIndexSet
	// SelectionIsChanging — Tells the delegate that the table view’s selection is in the process of changing.
	SelectionIsChanging func(notification foundation.NSNotification)
	// SelectionDidChange — Tells the delegate that the table view’s selection has changed.
	SelectionDidChange func(notification foundation.NSNotification)

	// Moving and resizing columns
	// ShouldReorderColumnToColumn — Asks the delegate to allow or prohibit the specified column to be dragged to a new location.
	ShouldReorderColumnToColumn func(tableView NSTableView, columnIndex int, newColumnIndex int) bool
	// ColumnDidMove — Tells the delegate that a table column was moved by user action.
	ColumnDidMove func(notification foundation.NSNotification)
	// ColumnDidResize — Tells the delegate that a table column was resized.
	ColumnDidResize func(notification foundation.NSNotification)

	// Other Methods
	// ViewForTableColumnRow — Asks the delegate for a view to display the specified row and column.
	ViewForTableColumnRow func(tableView NSTableView, tableColumn NSTableColumn, row int) NSView
	// DidAddRowViewForRow — Tells the delegate that a row view was added at the specified row.
	DidAddRowViewForRow func(tableView NSTableView, rowView NSTableRowView, row int)
	// DidRemoveRowViewForRow — Tells the delegate that a row view was removed from the table at the specified row.
	DidRemoveRowViewForRow func(tableView NSTableView, rowView NSTableRowView, row int)
	// DataCellForTableColumnRow — Asks the delegate for a custom data cell for the specified row and column.
	DataCellForTableColumnRow func(tableView NSTableView, tableColumn NSTableColumn, row int) NSCell
	// ShouldShowCellExpansionForTableColumnRow — Asks the delegate if an expansion tooltip should be displayed for a specific row and column.
	ShouldShowCellExpansionForTableColumnRow func(tableView NSTableView, tableColumn NSTableColumn, row int) bool
	// ShouldEditTableColumnRow — Asks the delegate if the cell at the specified row and column can be edited.
	ShouldEditTableColumnRow func(tableView NSTableView, tableColumn NSTableColumn, row int) bool
	// SelectionShouldChangeInTableView — Asks the delegate if the user is allowed to change the selection.
	SelectionShouldChangeInTableView func(tableView NSTableView) bool
	// ShouldSelectTableColumn — Asks the delegate whether the specified table column can be selected.
	ShouldSelectTableColumn func(tableView NSTableView, tableColumn NSTableColumn) bool
	// DidDragTableColumn — Tells the delegate that the specified table column was dragged.
	DidDragTableColumn func(tableView NSTableView, tableColumn NSTableColumn)
	// DidClickTableColumn — Tells the delegate that the mouse button was clicked in the specified table column, but the column was not dragged.
	DidClickTableColumn func(tableView NSTableView, tableColumn NSTableColumn)
	// MouseDownInHeaderOfTableColumn — Tells the delegate that the mouse button was clicked in the specified table column’s header.
	MouseDownInHeaderOfTableColumn func(tableView NSTableView, tableColumn NSTableColumn)
	// ShouldTrackCellForTableColumnRow — Asks the delegate whether the specified cell should be tracked.
	ShouldTrackCellForTableColumnRow func(tableView NSTableView, cell NSCell, tableColumn NSTableColumn, row int) bool
	// UserCanChangeVisibilityOfTableColumn — Asks the delegate to verify that the user can change the given column’s visibility.
	UserCanChangeVisibilityOfTableColumn func(tableView NSTableView, column NSTableColumn) bool
}

// NewNSTableViewDelegate creates an Objective-C object implementing the [NSTableViewDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTableViewDelegateObject] satisfies the [NSTableViewDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstableviewdelegate
func NewNSTableViewDelegate(config NSTableViewDelegateConfig) NSTableViewDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTableViewDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ViewForTableColumnRow != nil {
		fn := config.ViewForTableColumnRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:viewForTableColumn:row:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, tableColumnID objc.ID, row int) objc.ID {
				tableView := NSTableViewFromID(tableViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				return fn(tableView, tableColumn, row).GetID()
			},
		})
	}

	if config.RowViewForRow != nil {
		fn := config.RowViewForRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:rowViewForRow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, row int) objc.ID {
				tableView := NSTableViewFromID(tableViewID)
				return fn(tableView, row).GetID()
			},
		})
	}

	if config.DidAddRowViewForRow != nil {
		fn := config.DidAddRowViewForRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:didAddRowView:forRow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, rowViewID objc.ID, row int) {
				tableView := NSTableViewFromID(tableViewID)
				rowView := NSTableRowViewFromID(rowViewID)
				fn(tableView, rowView, row)
			},
		})
	}

	if config.DidRemoveRowViewForRow != nil {
		fn := config.DidRemoveRowViewForRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:didRemoveRowView:forRow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, rowViewID objc.ID, row int) {
				tableView := NSTableViewFromID(tableViewID)
				rowView := NSTableRowViewFromID(rowViewID)
				fn(tableView, rowView, row)
			},
		})
	}

	if config.IsGroupRow != nil {
		fn := config.IsGroupRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:isGroupRow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, row int) bool {
				tableView := NSTableViewFromID(tableViewID)
				return fn(tableView, row)
			},
		})
	}

	if config.DataCellForTableColumnRow != nil {
		fn := config.DataCellForTableColumnRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:dataCellForTableColumn:row:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, tableColumnID objc.ID, row int) objc.ID {
				tableView := NSTableViewFromID(tableViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				return fn(tableView, tableColumn, row).GetID()
			},
		})
	}

	if config.ShouldShowCellExpansionForTableColumnRow != nil {
		fn := config.ShouldShowCellExpansionForTableColumnRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:shouldShowCellExpansionForTableColumn:row:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, tableColumnID objc.ID, row int) bool {
				tableView := NSTableViewFromID(tableViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				return fn(tableView, tableColumn, row)
			},
		})
	}

	if config.ShouldEditTableColumnRow != nil {
		fn := config.ShouldEditTableColumnRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:shouldEditTableColumn:row:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, tableColumnID objc.ID, row int) bool {
				tableView := NSTableViewFromID(tableViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				return fn(tableView, tableColumn, row)
			},
		})
	}

	if config.HeightOfRow != nil {
		fn := config.HeightOfRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:heightOfRow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, row int) float64 {
				tableView := NSTableViewFromID(tableViewID)
				return fn(tableView, row)
			},
		})
	}

	if config.SizeToFitWidthOfColumn != nil {
		fn := config.SizeToFitWidthOfColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:sizeToFitWidthOfColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, column int) float64 {
				tableView := NSTableViewFromID(tableViewID)
				return fn(tableView, column)
			},
		})
	}

	if config.SelectionShouldChangeInTableView != nil {
		fn := config.SelectionShouldChangeInTableView
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("selectionShouldChangeInTableView:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID) bool {
				tableView := NSTableViewFromID(tableViewID)
				return fn(tableView)
			},
		})
	}

	if config.ShouldSelectRow != nil {
		fn := config.ShouldSelectRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:shouldSelectRow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, row int) bool {
				tableView := NSTableViewFromID(tableViewID)
				return fn(tableView, row)
			},
		})
	}

	if config.SelectionIndexesForProposedSelection != nil {
		fn := config.SelectionIndexesForProposedSelection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:selectionIndexesForProposedSelection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, proposedSelectionIndexesID objc.ID) objc.ID {
				tableView := NSTableViewFromID(tableViewID)
				proposedSelectionIndexes := foundation.NSIndexSetFromID(proposedSelectionIndexesID)
				return fn(tableView, proposedSelectionIndexes).GetID()
			},
		})
	}

	if config.ShouldSelectTableColumn != nil {
		fn := config.ShouldSelectTableColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:shouldSelectTableColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, tableColumnID objc.ID) bool {
				tableView := NSTableViewFromID(tableViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				return fn(tableView, tableColumn)
			},
		})
	}

	if config.SelectionIsChanging != nil {
		fn := config.SelectionIsChanging
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableViewSelectionIsChanging:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.SelectionDidChange != nil {
		fn := config.SelectionDidChange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableViewSelectionDidChange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ShouldReorderColumnToColumn != nil {
		fn := config.ShouldReorderColumnToColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:shouldReorderColumn:toColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, columnIndex int, newColumnIndex int) bool {
				tableView := NSTableViewFromID(tableViewID)
				return fn(tableView, columnIndex, newColumnIndex)
			},
		})
	}

	if config.DidDragTableColumn != nil {
		fn := config.DidDragTableColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:didDragTableColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, tableColumnID objc.ID) {
				tableView := NSTableViewFromID(tableViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				fn(tableView, tableColumn)
			},
		})
	}

	if config.ColumnDidMove != nil {
		fn := config.ColumnDidMove
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableViewColumnDidMove:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ColumnDidResize != nil {
		fn := config.ColumnDidResize
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableViewColumnDidResize:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidClickTableColumn != nil {
		fn := config.DidClickTableColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:didClickTableColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, tableColumnID objc.ID) {
				tableView := NSTableViewFromID(tableViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				fn(tableView, tableColumn)
			},
		})
	}

	if config.MouseDownInHeaderOfTableColumn != nil {
		fn := config.MouseDownInHeaderOfTableColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:mouseDownInHeaderOfTableColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, tableColumnID objc.ID) {
				tableView := NSTableViewFromID(tableViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				fn(tableView, tableColumn)
			},
		})
	}

	if config.ShouldTrackCellForTableColumnRow != nil {
		fn := config.ShouldTrackCellForTableColumnRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:shouldTrackCell:forTableColumn:row:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, cellID objc.ID, tableColumnID objc.ID, row int) bool {
				tableView := NSTableViewFromID(tableViewID)
				cell := NSCellFromID(cellID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				return fn(tableView, cell, tableColumn, row)
			},
		})
	}

	if config.UserCanChangeVisibilityOfTableColumn != nil {
		fn := config.UserCanChangeVisibilityOfTableColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tableView:userCanChangeVisibilityOfTableColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tableViewID objc.ID, columnID objc.ID) bool {
				tableView := NSTableViewFromID(tableViewID)
				column := NSTableColumnFromID(columnID)
				return fn(tableView, column)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTableViewDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTableViewDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTableViewDelegateObjectFromID(instance)
}


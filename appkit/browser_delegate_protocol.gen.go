// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// A set of methods that a browser delegate implements to manage selection, scrolling, sizing, and other behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate
type NSBrowserDelegate interface {
	objectivec.IObject
}



// NSBrowserDelegateObject wraps an existing Objective-C object that conforms to the NSBrowserDelegate protocol.
type NSBrowserDelegateObject struct {
	objectivec.Object
}
func (o NSBrowserDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSBrowserDelegateObjectFromID constructs a [NSBrowserDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSBrowserDelegateObjectFromID(id objc.ID) NSBrowserDelegateObject {
	return NSBrowserDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Returns whether the contents of the specified column are valid.
//
// sender: The browser containing the column to validate.
//
// column: The index of the column to validate.
//
// # Return Value
// 
// [true] if the column’s contents are valid; otherwise, [false]. If [false]
// is returned, `sender` reloads the column.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is invoked in response to the [ValidateVisibleColumns]method of
// [NSBrowser] being sent to `sender`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:isColumnValid:)

func (o NSBrowserDelegateObject) BrowserIsColumnValid(sender INSBrowser, column int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("browser:isColumnValid:"), sender, column)
	return rv
	}

// Returns the number of rows of data in the specified column.
//
// sender: The browser.
//
// column: The index of the column.
//
// # Return Value
// 
// The number of rows of data.
//
// # Discussion
// 
// Either this method or [BrowserCreateRowsForColumnInMatrix] must be
// implemented, but not both.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:numberOfRowsInColumn:)

func (o NSBrowserDelegateObject) BrowserNumberOfRowsInColumn(sender INSBrowser, column int) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("browser:numberOfRowsInColumn:"), sender, column)
	return rv
	}

// Asks the delegate for the number of children the given item has.
//
// browser: The browser.
//
// item: The item that has some number of children.
//
// # Return Value
// 
// The number of children.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:numberOfChildrenOfItem:)

func (o NSBrowserDelegateObject) BrowserNumberOfChildrenOfItem(browser INSBrowser, item objectivec.IObject) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("browser:numberOfChildrenOfItem:"), browser, item)
	return rv
	}

// Asks the delegate for the title to display above the specified column.
//
// sender: The browser.
//
// column: The index the column to be titled.
//
// # Return Value
// 
// The title of the specified column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:titleOfColumn:)

func (o NSBrowserDelegateObject) BrowserTitleOfColumn(sender INSBrowser, column int) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("browser:titleOfColumn:"), sender, column)
	return foundation.NSStringFromID(rv).String()
	}

// Sent to the delegate to determine whether keyboard-based selection (type
// select) for a given event and search string should proceed.
//
// browser: The browser.
//
// event: The keyboard event being processed.
//
// searchString: The keyboard-based selection string. It is `nil` when no keyboard-based
// selection has begun.
//
// # Return Value
// 
// [true] to allow the selection; [false] to disallow it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:shouldTypeSelectFor:withCurrentSearch:)

func (o NSBrowserDelegateObject) BrowserShouldTypeSelectForEventWithCurrentSearchString(browser INSBrowser, event INSEvent, searchString string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("browser:shouldTypeSelectForEvent:withCurrentSearchString:"), browser, event, objc.String(searchString))
	return rv
	}

// Sent to the delegate to get the keyboard-based selection (type select)
// string for the specified row and column.
//
// browser: The browser.
//
// row: The row index.
//
// column: The column index.
//
// # Return Value
// 
// The keyboard-based selection string.
//
// # Discussion
// 
// Returning the empty string or `nil` (for example, when the cell does not
// contain text) specifies that the `[``column`, `row``]` cell has no text to
// search.
// 
// If the delegate does not implement this method, all cells with text are
// searched, and the browser determines the keyboard-based selection text by
// sending [StringValue] to the cell specified by `column` and `row`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:typeSelectStringForRow:inColumn:)

func (o NSBrowserDelegateObject) BrowserTypeSelectStringForRowInColumn(browser INSBrowser, row int, column int) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("browser:typeSelectStringForRow:inColumn:"), browser, row, column)
	return foundation.NSStringFromID(rv).String()
	}

// Sent to the delegate to customize a browser’s keyboard-based selection
// (type select) behavior.
//
// browser: The browser.
//
// startRow: The beginning of the row set to search.
//
// endRow: The end of the row set to search. This value can be less than
// `startRowIndex` when the search wraps around to the beginning.
//
// column: The column containing the rows being searched.
//
// searchString: The keyboard-based selection string. It is `nil` when no keyboard-based
// selection has begun.
//
// # Return Value
// 
// The index of the first row that matches `searchString` between
// `startRowIndex` and `endRowIndex` - 1, or `-1` if there is no match.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:nextTypeSelectMatchFromRow:toRow:inColumn:for:)

func (o NSBrowserDelegateObject) BrowserNextTypeSelectMatchFromRowToRowInColumnForString(browser INSBrowser, startRow int, endRow int, column int, searchString string) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("browser:nextTypeSelectMatchFromRow:toRow:inColumn:forString:"), browser, startRow, endRow, column, objc.String(searchString))
	return rv
	}

// Asks the delegate to select the cell with the given title in the specified
// column.
//
// sender: The browser.
//
// title: The title of the cell to select.
//
// column: The index of the column containing the cell to select.
//
// # Return Value
// 
// [true] if the cell was successfully selected; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Invoked in response to the [SetPath] method of [NSBrowser] being received
// by `sender`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:selectCellWith:inColumn:)

func (o NSBrowserDelegateObject) BrowserSelectCellWithStringInColumn(sender INSBrowser, title string, column int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("browser:selectCellWithString:inColumn:"), sender, objc.String(title), column)
	return rv
	}

// Asks the delegate to select the cell at the specified row and column
// location.
//
// sender: The browser.
//
// row: The index of the row containing the cell to select.
//
// column: The index of the column containing the cell to select.
//
// # Return Value
// 
// [true] if the cell was selected; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Invoked in response to [SelectRowInColumn] of [NSBrowser] being received by
// `sender`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:selectRow:inColumn:)

func (o NSBrowserDelegateObject) BrowserSelectRowInColumn(sender INSBrowser, row int, column int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("browser:selectRow:inColumn:"), sender, row, column)
	return rv
	}

// Asks the delegate for a set of indexes to select when the user changes the
// selection in the browser with the keyboard or mouse.
//
// browser: The browser.
//
// proposedSelectionIndexes: The set of indexes of the items in the proposed selection.
//
// column: The column index of the column containing the selection.
//
// # Return Value
// 
// The set of indexes of the items that should be selected.
//
// # Discussion
// 
// This method may be called multiple times, with one new index added to the
// previous selection, to see whether a particular index can be selected when
// the user is extending the selection with the keyboard or mouse. The
// `proposedSelectionIndexes` parameter contains the entire selection, and you
// can return the existing selection if you do not want to change it. This
// method works only for item-based browsers.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:selectionIndexesForProposedSelection:inColumn:)

func (o NSBrowserDelegateObject) BrowserSelectionIndexesForProposedSelectionInColumn(browser INSBrowser, proposedSelectionIndexes foundation.NSIndexSet, column int) foundation.NSIndexSet {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("browser:selectionIndexesForProposedSelection:inColumn:"), browser, proposedSelectionIndexes, column)
	return foundation.NSIndexSetFromID(rv)
	}

// Asks the delegate to return the child of the specified item at the
// specified index.
//
// browser: The browser.
//
// index: The child’s index.
//
// item: The item containing the child.
//
// # Return Value
// 
// The child at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:child:ofItem:)

func (o NSBrowserDelegateObject) BrowserChildOfItem(browser INSBrowser, index int, item objectivec.IObject) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("browser:child:ofItem:"), browser, index, item)
	return objectivec.Object{ID: rv}
	}

// Asks the delegate whether the specified item is a leaf item (an item that
// cannot be expanded).
//
// browser: The browser.
//
// item: The item to be checked.
//
// # Return Value
// 
// [true] if the specified item is a leaf item; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:isLeafItem:)

func (o NSBrowserDelegateObject) BrowserIsLeafItem(browser INSBrowser, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("browser:isLeafItem:"), browser, item)
	return rv
	}

// Asks the delegate whether the browser may start an editing session for the
// specified item.
//
// browser: The browser.
//
// item: The item to edit.
//
// # Return Value
// 
// [true] to allow the editing session to begin; [false] to disallow it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:shouldEditItem:)

func (o NSBrowserDelegateObject) BrowserShouldEditItem(browser INSBrowser, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("browser:shouldEditItem:"), browser, item)
	return rv
	}

// Returns the object that the specified item uses to draw its contents.
//
// browser: The browser.
//
// item: The item in question.
//
// # Return Value
// 
// The item’s object.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:objectValueForItem:)

func (o NSBrowserDelegateObject) BrowserObjectValueForItem(browser INSBrowser, item objectivec.IObject) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("browser:objectValueForItem:"), browser, item)
	return objectivec.Object{ID: rv}
	}

// Sets the object that the specified item uses to draw its contents to the
// specified object.
//
// browser: The browser.
//
// object: The object to set.
//
// item: The item whose object is set.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:setObjectValue:forItem:)

func (o NSBrowserDelegateObject) BrowserSetObjectValueForItem(browser INSBrowser, object objectivec.IObject, item objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("browser:setObjectValue:forItem:"), browser, object, item)
	}

// Asks the delegate to return the root item of the browser.
//
// browser: The browser.
//
// # Return Value
// 
// The browser’s root item.
//
// # Discussion
// 
// By default, `nil` identifies the root item. This method can specify a
// different root item. To reload the previously set root item, call
// [LoadColumnZero], and [RootItemForBrowser] will be called again.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/rootItem(for:)

func (o NSBrowserDelegateObject) RootItemForBrowser(browser INSBrowser) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("rootItemForBrowser:"), browser)
	return objectivec.Object{ID: rv}
	}

// Asks the delegate for a controller that provides a preview column for the
// specified leaf item.
//
// browser: The browser.
//
// item: The leaf item.
//
// # Return Value
// 
// A view controller that provides a preview column, or `nil` to suppress the
// preview column.
//
// # Discussion
// 
// The returned controller’s represented object is set to the specified leaf
// item. This method is called only if the delegate implements the item data
// source methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:previewViewControllerForLeafItem:)

func (o NSBrowserDelegateObject) BrowserPreviewViewControllerForLeafItem(browser INSBrowser, item objectivec.IObject) INSViewController {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("browser:previewViewControllerForLeafItem:"), browser, item)
	return NSViewControllerFromID(rv)
	}

// Asks the delegate for a controller that provides a header view for the
// specified column item.
//
// browser: The browser.
//
// item: The column item.
//
// # Return Value
// 
// A view controller that provides a header view, or `nil` to omit the header
// view.
//
// # Discussion
// 
// The returned controller’s represented object will be set to the column
// item. This method is called only if the delegate implements the item data
// source methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:headerViewControllerForItem:)

func (o NSBrowserDelegateObject) BrowserHeaderViewControllerForItem(browser INSBrowser, item objectivec.IObject) INSViewController {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("browser:headerViewControllerForItem:"), browser, item)
	return NSViewControllerFromID(rv)
	}

// Creates a row in the given matrix for each row of data in the specified
// column of the browser.
//
// sender: The browser.
//
// column: The index of the column in which the rows are located.
//
// matrix: The matrix in which the rows are created.
//
// # Discussion
// 
// Either this method or [BrowserNumberOfRowsInColumn] must be implemented,
// but not both, or an [NSBrowserIllegalDelegateException] will be raised.
//
// [NSBrowserIllegalDelegateException]: https://developer.apple.com/documentation/AppKit/NSBrowserIllegalDelegateException
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:createRowsForColumn:in:)

func (o NSBrowserDelegateObject) BrowserCreateRowsForColumnInMatrix(sender INSBrowser, column int, matrix INSMatrix) {
	
	objc.Send[struct{}](o.ID, objc.Sel("browser:createRowsForColumn:inMatrix:"), sender, column, matrix)
	}

// Gives the delegate the opportunity to modify the specified cell at the
// given row and column location before the browser displays it.
//
// sender: The browser.
//
// cell: The cell to be displayed.
//
// row: The row index of the cell to be displayed.
//
// column: The column index of the cell to be displayed.
//
// # Discussion
// 
// The delegate should set any state necessary for the correct display of the
// cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:willDisplayCell:atRow:column:)

func (o NSBrowserDelegateObject) BrowserWillDisplayCellAtRowColumn(sender INSBrowser, cell objectivec.IObject, row int, column int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("browser:willDisplayCell:atRow:column:"), sender, cell, row, column)
	}

// Tells the delegate that the browser’s last column changed.
//
// browser: The browser.
//
// oldLastColumn: The index of the old last column.
//
// column: The index of the new last column.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:didChangeLastColumn:toColumn:)

func (o NSBrowserDelegateObject) BrowserDidChangeLastColumnToColumn(browser INSBrowser, oldLastColumn int, column int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("browser:didChangeLastColumn:toColumn:"), browser, oldLastColumn, column)
	}

// Notifies the delegate when the browser will scroll.
//
// sender: The browser sending the message.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browserWillScroll(_:)

func (o NSBrowserDelegateObject) BrowserWillScroll(sender INSBrowser) {
	
	objc.Send[struct{}](o.ID, objc.Sel("browserWillScroll:"), sender)
	}

// Notifies the delegate when the browser has scrolled.
//
// sender: The browser sending the message.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browserDidScroll(_:)

func (o NSBrowserDelegateObject) BrowserDidScroll(sender INSBrowser) {
	
	objc.Send[struct{}](o.ID, objc.Sel("browserDidScroll:"), sender)
	}

// Sent to the delegate to determine whether the browser can attempt to
// initiate a drag of the specified rows for the specified event.
//
// browser: The browser.
//
// rowIndexes: The rows the user is dragging.
//
// column: The column containing the rows the user is dragging.
//
// event: The drag event.
//
// # Return Value
// 
// [true] to allow the drag operation; [false] to disallow it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:canDragRowsWith:inColumn:with:)

func (o NSBrowserDelegateObject) BrowserCanDragRowsWithIndexesInColumnWithEvent(browser INSBrowser, rowIndexes foundation.NSIndexSet, column int, event INSEvent) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("browser:canDragRowsWithIndexes:inColumn:withEvent:"), browser, rowIndexes, column, event)
	return rv
	}

// Sent to the delegate to obtain an image to represent dragged rows during a
// drag operation on a browser.
//
// browser: The browser.
//
// rowIndexes: The indexes of the rows the user is dragging.
//
// column: The column containing the rows the user is dragging.
//
// event: The drag event.
//
// dragImageOffset: The offset for the returned image:
// 
// - [NSZeroPoint]: Centers the image under the pointer.
// //
// [NSZeroPoint]: https://developer.apple.com/documentation/Foundation/NSZeroPoint
//
// # Return Value
// 
// An image representing the visible rows identified by `rowIndexes`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:draggingImageForRowsWith:inColumn:with:offset:)

func (o NSBrowserDelegateObject) BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset(browser INSBrowser, rowIndexes foundation.NSIndexSet, column int, event INSEvent, dragImageOffset foundation.NSPoint) INSImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("browser:draggingImageForRowsWithIndexes:inColumn:withEvent:offset:"), browser, rowIndexes, column, event, dragImageOffset)
	return NSImageFromID(rv)
	}

// Sent to the delegate during a dragging session to determine whether a drop
// should be accepted and to obtain the drop location. This method is required
// for a browser to be a drag destination.
//
// browser: The browser.
//
// info: The drag session information.
//
// row: On input, the proposed drop row. On output, the drop row.
//
// column: On input, the proposed drop column. On output, the drop column.
//
// dropOperation: On input, the proposed drop location. On output, the drop location.
//
// # Return Value
// 
// The drag operation that the data source is to perform. For the browser to
// accept the drop, it must not be [NSDragOperationNone].
//
// [NSDragOperationNone]: https://developer.apple.com/documentation/AppKit/NSDragOperation/NSDragOperationNone
//
// # Discussion
// 
// The browser proposes a drop column, row, and row-relative location for the
// drop based on the pointer position, as shown in this table:
// 
// [Table data omitted]
// 
// These are a few examples of how to specify a drop location:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:validateDrop:proposedRow:column:dropOperation:)

func (o NSBrowserDelegateObject) BrowserValidateDropProposedRowColumnDropOperation(browser INSBrowser, info NSDraggingInfo, row int, column int, dropOperation NSBrowserDropOperation) NSDragOperation {
	
	rv := objc.Send[NSDragOperation](o.ID, objc.Sel("browser:validateDrop:proposedRow:column:dropOperation:"), browser, info, row, column, dropOperation)
	return rv
	}

// Sent to the delegate during a dragging session to determine whether to
// accept the drop.
//
// browser: The browser.
//
// info: The drag session information.
//
// row: The drop row.
//
// column: The drop column.
//
// dropOperation: The drop location relative to `row`.
//
// # Return Value
// 
// [true] to accept the drop; [false] to decline it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is required for a browser to be a drag destination. It is
// invoked after the [BrowserValidateDropProposedRowColumnDropOperation]
// method allows the drop.
// 
// The delegate should incorporate the pasteboard data from the dragging
// session (`info``XCUIElementTypeDraggingPasteboard`).
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:acceptDrop:atRow:column:dropOperation:)

func (o NSBrowserDelegateObject) BrowserAcceptDropAtRowColumnDropOperation(browser INSBrowser, info NSDraggingInfo, row int, column int, dropOperation NSBrowserDropOperation) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("browser:acceptDrop:atRow:column:dropOperation:"), browser, info, row, column, dropOperation)
	return rv
	}

// Determines whether a drag operation can proceed. This method is required
// for a browser to be a drag source.
//
// browser: The browser.
//
// rowIndexes: The indexes of the rows the user is dragging.
//
// column: The index of the column containing the dragged rows.
//
// pasteboard: The pasteboard containing the content from the dragged rows.
//
// # Return Value
// 
// [true] to allow the dragging operation to proceed (see discussion for
// further details); [false] to disallow it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is called after a drag operation has been allowed to start
// ([BrowserCanDragRowsWithIndexesInColumnWithEvent] returns [true]), but
// before it actually begins.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:writeRowsWith:inColumn:to:)

func (o NSBrowserDelegateObject) BrowserWriteRowsWithIndexesInColumnToPasteboard(browser INSBrowser, rowIndexes foundation.NSIndexSet, column int, pasteboard INSPasteboard) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("browser:writeRowsWithIndexes:inColumn:toPasteboard:"), browser, rowIndexes, column, pasteboard)
	return rv
	}

// Used to determine a column’s initial size.
//
// browser: The browser.
//
// columnIndex: The index of the column to size.
//
// forUserResize: Currently, this is always set to [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// suggestedWidth: The suggested width for the column.
//
// # Return Value
// 
// The delegate’s desired initial width for a newly added column. If you
// want to accept the suggested width, return `suggestedWidth`. If you return
// `0` or a size too small to display the resize handle and a portion of the
// column, the actual size used will be larger than the size you requested.
//
// # Discussion
// 
// This method applies only to browsers with resize type
// [NSBrowser.ColumnResizingType.noColumnResizing] or
// [NSBrowser.ColumnResizingType.userColumnResizing] (see
// [NSBrowser.ColumnResizingType]).
//
// [NSBrowser.ColumnResizingType.noColumnResizing]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum/noColumnResizing
// [NSBrowser.ColumnResizingType.userColumnResizing]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum/userColumnResizing
// [NSBrowser.ColumnResizingType]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:shouldSizeColumn:forUserResize:toWidth:)

func (o NSBrowserDelegateObject) BrowserShouldSizeColumnForUserResizeToWidth(browser INSBrowser, columnIndex int, forUserResize bool, suggestedWidth float64) float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("browser:shouldSizeColumn:forUserResize:toWidth:"), browser, columnIndex, forUserResize, suggestedWidth)
	return rv
	}

// Returns the ideal width for a column.
//
// browser: The browser.
//
// columnIndex: The index of the column to size. If `-1`, the result is used to resize all
// columns.
//
// # Return Value
// 
// The ideal width of the column. This method is used when performing a
// “right-size” operation, that is, when sizing a column to the smallest
// width that contains all the content without clipping or truncating.
//
// # Discussion
// 
// If `columnIndex` is `–1`, you should return a size that can be uniformly
// applied to all columns (that is, every column will be set to this size).
// 
// Returning a value of `-1` allows you to opt-out of providing a width for
// the requested column.
// 
// # Discussion
// 
// This method applies only to browsers with resize type
// [NSBrowser.ColumnResizingType.userColumnResizing].
// 
// It is assumed that the implementation may be expensive, so it will be
// called only when necessary.
//
// [NSBrowser.ColumnResizingType.userColumnResizing]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum/userColumnResizing
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:sizeToFitWidthOfColumn:)

func (o NSBrowserDelegateObject) BrowserSizeToFitWidthOfColumn(browser INSBrowser, columnIndex int) float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("browser:sizeToFitWidthOfColumn:"), browser, columnIndex)
	return rv
	}

// Used by clients to implement their own column width persistence.
//
// notification: A notification named [columnConfigurationDidChangeNotification].
// //
// [columnConfigurationDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSBrowser/columnConfigurationDidChangeNotification
//
// # Discussion
// 
// This method applies only to browsers with resize type
// [NSBrowser.ColumnResizingType.userColumnResizing]. It is invoked when the
// [SetWidthOfColumn] method of [NSBrowser] is used to change the width of any
// browser columns or when the user resizes any columns. If the user resizes
// more than one column, a single notification is posted when the user is
// finished resizing.
//
// [NSBrowser.ColumnResizingType.userColumnResizing]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum/userColumnResizing
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browserColumnConfigurationDidChange(_:)

func (o NSBrowserDelegateObject) BrowserColumnConfigurationDidChange(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("browserColumnConfigurationDidChange:"), notification)
	}

// Specifies the height of the specified row in the specified column.
//
// browser: The browser.
//
// row: The index of the row.
//
// columnIndex: The index of the column.
//
// # Return Value
// 
// The height to set for the specified row, which must be greater than 0.
//
// # Discussion
// 
// The values returned for this method may be cached. Therefore, you should
// call [NoteHeightOfRowsWithIndexesChangedInColumn] to invalidate a row’s
// height before changing it.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:heightOfRow:inColumn:)

func (o NSBrowserDelegateObject) BrowserHeightOfRowInColumn(browser INSBrowser, row int, columnIndex int) float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("browser:heightOfRow:inColumn:"), browser, row, columnIndex)
	return rv
	}

// Invoked to allow the delegate to control cell expansion for a specific row
// and column.
//
// browser: The browser.
//
// row: The index of the row requesting an expansion tooltip.
//
// column: The index of the column containing the requesting row.
//
// # Return Value
// 
// [true] to allow the cell expansion tooltip; [false] to disallow it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Cell expansion can occur when the mouse hovers over the specified cell and
// the cell contents are unable to be fully displayed within the cell. If this
// method returns YES, the full cell contents will be shown in a special
// floating tool tip view, otherwise the content is truncated.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserDelegate/browser(_:shouldShowCellExpansionForRow:column:)

func (o NSBrowserDelegateObject) BrowserShouldShowCellExpansionForRowColumn(browser INSBrowser, row int, column int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("browser:shouldShowCellExpansionForRow:column:"), browser, row, column)
	return rv
	}





// NSBrowserDelegateConfig holds optional typed callbacks for [NSBrowserDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate
type NSBrowserDelegateConfig struct {

	// Getting Browser Information
	// IsColumnValid — Returns whether the contents of the specified column are valid.
	IsColumnValid func(sender NSBrowser, column int) bool
	// NumberOfRowsInColumn — Returns the number of rows of data in the specified column.
	NumberOfRowsInColumn func(sender NSBrowser, column int) int

	// Managing Selection
	// SelectRowInColumn — Asks the delegate to select the cell at the specified row and column location.
	SelectRowInColumn func(sender NSBrowser, row int, column int) bool
	// SelectionIndexesForProposedSelectionInColumn — Asks the delegate for a set of indexes to select when the user changes the selection in the browser with the keyboard or mouse.
	SelectionIndexesForProposedSelectionInColumn func(browser NSBrowser, proposedSelectionIndexes foundation.NSIndexSet, column int) foundation.NSIndexSet

	// Managing Columns
	// DidChangeLastColumnToColumn — Tells the delegate that the browser’s last column changed.
	DidChangeLastColumnToColumn func(browser NSBrowser, oldLastColumn int, column int)

	// Scrolling
	// WillScroll — Notifies the delegate when the browser will scroll.
	WillScroll func(sender NSBrowser)
	// DidScroll — Notifies the delegate when the browser has scrolled.
	DidScroll func(sender NSBrowser)

	// Sizing
	// ShouldSizeColumnForUserResizeToWidth — Used to determine a column’s initial size.
	ShouldSizeColumnForUserResizeToWidth func(browser NSBrowser, columnIndex int, forUserResize bool, suggestedWidth float64) float64
	// SizeToFitWidthOfColumn — Returns the ideal width for a column.
	SizeToFitWidthOfColumn func(browser NSBrowser, columnIndex int) float64
	// ColumnConfigurationDidChange — Used by clients to implement their own column width persistence.
	ColumnConfigurationDidChange func(notification foundation.NSNotification)
	// HeightOfRowInColumn — Specifies the height of the specified row in the specified column.
	HeightOfRowInColumn func(browser NSBrowser, row int, columnIndex int) float64

	// Displaying Cell Content
	// ShouldShowCellExpansionForRowColumn — Invoked to allow the delegate to control cell expansion for a specific row and column.
	ShouldShowCellExpansionForRowColumn func(browser NSBrowser, row int, column int) bool

	// Other Methods
	// CreateRowsForColumnInMatrix — Creates a row in the given matrix for each row of data in the specified column of the browser.
	CreateRowsForColumnInMatrix func(sender NSBrowser, column int, matrix NSMatrix)
	// CanDragRowsWithIndexesInColumnWithEvent — Sent to the delegate to determine whether the browser can attempt to initiate a drag of the specified rows for the specified event.
	CanDragRowsWithIndexesInColumnWithEvent func(browser NSBrowser, rowIndexes foundation.NSIndexSet, column int, event NSEvent) bool
	// DraggingImageForRowsWithIndexesInColumnWithEventOffset — Sent to the delegate to obtain an image to represent dragged rows during a drag operation on a browser.
	DraggingImageForRowsWithIndexesInColumnWithEventOffset func(browser NSBrowser, rowIndexes foundation.NSIndexSet, column int, event NSEvent, dragImageOffset foundation.NSPoint) NSImage
	// WriteRowsWithIndexesInColumnToPasteboard — Determines whether a drag operation can proceed. This method is required for a browser to be a drag source.
	WriteRowsWithIndexesInColumnToPasteboard func(browser NSBrowser, rowIndexes foundation.NSIndexSet, column int, pasteboard NSPasteboard) bool
}

// NewNSBrowserDelegate creates an Objective-C object implementing the [NSBrowserDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSBrowserDelegateObject] satisfies the [NSBrowserDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsbrowserdelegate
func NewNSBrowserDelegate(config NSBrowserDelegateConfig) NSBrowserDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSBrowserDelegate_%d", n)

	var methods []objc.MethodDef

	if config.IsColumnValid != nil {
		fn := config.IsColumnValid
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:isColumnValid:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, column int) bool {
				sender := NSBrowserFromID(senderID)
				return fn(sender, column)
			},
		})
	}

	if config.NumberOfRowsInColumn != nil {
		fn := config.NumberOfRowsInColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:numberOfRowsInColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, column int) int {
				sender := NSBrowserFromID(senderID)
				return fn(sender, column)
			},
		})
	}

	if config.SelectRowInColumn != nil {
		fn := config.SelectRowInColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:selectRow:inColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, row int, column int) bool {
				sender := NSBrowserFromID(senderID)
				return fn(sender, row, column)
			},
		})
	}

	if config.SelectionIndexesForProposedSelectionInColumn != nil {
		fn := config.SelectionIndexesForProposedSelectionInColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:selectionIndexesForProposedSelection:inColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, proposedSelectionIndexesID objc.ID, column int) objc.ID {
				browser := NSBrowserFromID(browserID)
				proposedSelectionIndexes := foundation.NSIndexSetFromID(proposedSelectionIndexesID)
				return fn(browser, proposedSelectionIndexes, column).GetID()
			},
		})
	}

	if config.CreateRowsForColumnInMatrix != nil {
		fn := config.CreateRowsForColumnInMatrix
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:createRowsForColumn:inMatrix:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, column int, matrixID objc.ID) {
				sender := NSBrowserFromID(senderID)
				matrix := NSMatrixFromID(matrixID)
				fn(sender, column, matrix)
			},
		})
	}

	if config.DidChangeLastColumnToColumn != nil {
		fn := config.DidChangeLastColumnToColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:didChangeLastColumn:toColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, oldLastColumn int, column int) {
				browser := NSBrowserFromID(browserID)
				fn(browser, oldLastColumn, column)
			},
		})
	}

	if config.WillScroll != nil {
		fn := config.WillScroll
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browserWillScroll:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := NSBrowserFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.DidScroll != nil {
		fn := config.DidScroll
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browserDidScroll:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := NSBrowserFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.CanDragRowsWithIndexesInColumnWithEvent != nil {
		fn := config.CanDragRowsWithIndexesInColumnWithEvent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:canDragRowsWithIndexes:inColumn:withEvent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, rowIndexesID objc.ID, column int, eventID objc.ID) bool {
				browser := NSBrowserFromID(browserID)
				rowIndexes := foundation.NSIndexSetFromID(rowIndexesID)
				event := NSEventFromID(eventID)
				return fn(browser, rowIndexes, column, event)
			},
		})
	}

	if config.DraggingImageForRowsWithIndexesInColumnWithEventOffset != nil {
		fn := config.DraggingImageForRowsWithIndexesInColumnWithEventOffset
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:draggingImageForRowsWithIndexes:inColumn:withEvent:offset:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, rowIndexesID objc.ID, column int, eventID objc.ID, dragImageOffset foundation.NSPoint) objc.ID {
				browser := NSBrowserFromID(browserID)
				rowIndexes := foundation.NSIndexSetFromID(rowIndexesID)
				event := NSEventFromID(eventID)
				return fn(browser, rowIndexes, column, event, dragImageOffset).GetID()
			},
		})
	}

	if config.WriteRowsWithIndexesInColumnToPasteboard != nil {
		fn := config.WriteRowsWithIndexesInColumnToPasteboard
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:writeRowsWithIndexes:inColumn:toPasteboard:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, rowIndexesID objc.ID, column int, pasteboardID objc.ID) bool {
				browser := NSBrowserFromID(browserID)
				rowIndexes := foundation.NSIndexSetFromID(rowIndexesID)
				pasteboard := NSPasteboardFromID(pasteboardID)
				return fn(browser, rowIndexes, column, pasteboard)
			},
		})
	}

	if config.ShouldSizeColumnForUserResizeToWidth != nil {
		fn := config.ShouldSizeColumnForUserResizeToWidth
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:shouldSizeColumn:forUserResize:toWidth:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, columnIndex int, forUserResize bool, suggestedWidth float64) float64 {
				browser := NSBrowserFromID(browserID)
				return fn(browser, columnIndex, forUserResize, suggestedWidth)
			},
		})
	}

	if config.SizeToFitWidthOfColumn != nil {
		fn := config.SizeToFitWidthOfColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:sizeToFitWidthOfColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, columnIndex int) float64 {
				browser := NSBrowserFromID(browserID)
				return fn(browser, columnIndex)
			},
		})
	}

	if config.ColumnConfigurationDidChange != nil {
		fn := config.ColumnConfigurationDidChange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browserColumnConfigurationDidChange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.HeightOfRowInColumn != nil {
		fn := config.HeightOfRowInColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:heightOfRow:inColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, row int, columnIndex int) float64 {
				browser := NSBrowserFromID(browserID)
				return fn(browser, row, columnIndex)
			},
		})
	}

	if config.ShouldShowCellExpansionForRowColumn != nil {
		fn := config.ShouldShowCellExpansionForRowColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("browser:shouldShowCellExpansionForRow:column:"),
			Fn: func(self objc.ID, _cmd objc.SEL, browserID objc.ID, row int, column int) bool {
				browser := NSBrowserFromID(browserID)
				return fn(browser, row, column)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSBrowserDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSBrowserDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSBrowserDelegateObjectFromID(instance)
}






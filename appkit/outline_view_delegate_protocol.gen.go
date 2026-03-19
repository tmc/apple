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

// A set of optional methods implemented by delegates of [NSOutlineView](<doc://com.apple.appkit/documentation/AppKit/NSOutlineView>) objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate
type NSOutlineViewDelegate interface {
	objectivec.IObject
	NSControlTextEditingDelegate
}

// NSOutlineViewDelegateObject wraps an existing Objective-C object that conforms to the NSOutlineViewDelegate protocol.
type NSOutlineViewDelegateObject struct {
	objectivec.Object
}
func (o NSOutlineViewDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSOutlineViewDelegateObjectFromID constructs a [NSOutlineViewDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSOutlineViewDelegateObjectFromID(id objc.ID) NSOutlineViewDelegateObject {
	return NSOutlineViewDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a Boolean value that indicates whether the outline view should
// expand a given item.
//
// outlineView: The outline view that sent the message.
//
// item: The item that should expand.
//
// # Return Value
// 
// [true] to permit `outlineView` to expand `item`, [false] to deny
// permission.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate can implement this method to disallow expanding of specific
// items.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:shouldExpandItem:)
func (o NSOutlineViewDelegateObject) OutlineViewShouldExpandItem(outlineView INSOutlineView, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:shouldExpandItem:"), outlineView, item)
	return rv
	}
// Returns a Boolean value that indicates whether the outline view should
// collapse a given item.
//
// outlineView: The outline view that sent the message.
//
// item: The item that should collapse.
//
// # Return Value
// 
// [true] to permit `outlineView` to collapse `item`, [false] to deny
// permission.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate can implement this method to disallow collapsing of specific
// items. For example, if the first row of your outline view should not be
// collapsed, your delegate method could contain this line of code:
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:shouldCollapseItem:)
func (o NSOutlineViewDelegateObject) OutlineViewShouldCollapseItem(outlineView INSOutlineView, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:shouldCollapseItem:"), outlineView, item)
	return rv
	}
// Returns the string that is used for type selection for a given column and
// item.
//
// outlineView: The outline view that sent the message.
//
// tableColumn: A table column in the outline view.
//
// item: An item in the outline view.
//
// # Return Value
// 
// The string that is used for type selection. You may want to change what is
// searched for based on what is displayed, or simply return nil for that row
// and/or column to not be searched
//
// # Discussion
// 
// Implement this method if you want to control the string that is used for
// type selection. You may want to change what is searched for based on what
// is displayed, or simply return `nil` to specify that the given row and/or
// column should not be searched. By default, all cells with text in them are
// searched.
// 
// The default value when this delegate method is not implemented is:
// 
// and you can return this value from the delegate method if you wish.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:typeSelectStringFor:item:)
func (o NSOutlineViewDelegateObject) OutlineViewTypeSelectStringForTableColumnItem(outlineView INSOutlineView, tableColumn INSTableColumn, item objectivec.IObject) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:typeSelectStringForTableColumn:item:"), outlineView, tableColumn, item)
	return foundation.NSStringFromID(rv).String()
	}
// Returns the first item that matches the searchString from within the range
// of startItem to endItem
//
// outlineView: The outline view that sent the message.
//
// startItem: The first item to search.
//
// endItem: The item before which to stop searching. It is possible for endItem to be
// less than startItem if the search will wrap.
//
// searchString: The string for which to search.
//
// # Return Value
// 
// The first item—from within the range of `startItem` to `endItem`—that
// matches the `searchString`, or `nil` if there is no match.
//
// # Discussion
// 
// Implement this method if you want to control how type selection works. You
// should include `startItem` as a possible match, but do not include
// `endItem`.
// 
// It is not necessary to implement this method in order to support type
// select.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:nextTypeSelectMatchFromItem:toItem:for:)
func (o NSOutlineViewDelegateObject) OutlineViewNextTypeSelectMatchFromItemToItemForString(outlineView INSOutlineView, startItem objectivec.IObject, endItem objectivec.IObject, searchString string) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:nextTypeSelectMatchFromItem:toItem:forString:"), outlineView, startItem, endItem, objc.String(searchString))
	return objectivec.Object{ID: rv}
	}
// Returns a Boolean value that indicates whether type select should proceed
// for a given event and search string.
//
// outlineView: The outline view that sent the message.
//
// event: The event that caused the message to be sent.
//
// searchString: The string for which searching is to proceed. The search string is `nil` if
// no type select has begun.
//
// # Return Value
// 
// [true] if type select should proceed, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Generally, this method will be called from [KeyDown] and the event will be
// a key event.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:shouldTypeSelectFor:withCurrentSearch:)
func (o NSOutlineViewDelegateObject) OutlineViewShouldTypeSelectForEventWithCurrentSearchString(outlineView INSOutlineView, event INSEvent, searchString string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:shouldTypeSelectForEvent:withCurrentSearchString:"), outlineView, event, objc.String(searchString))
	return rv
	}
// When the cursor pauses over a given cell, the value returned from this
// method is displayed in a tooltip.
//
// outlineView: The outline view that sent the message.
//
// cell: The cell for which to generate a tooltip.
//
// rect: The proposed active area of the tooltip. To control the default active
// area, you can modify the `rect` parameter. By default, `rect` is computed
// as `[cell cellFrame]`.
//
// tableColumn: The table column that contains `cell`.
//
// item: The item for which to display a tooltip.
//
// mouseLocation: The current mouse location in view coordinates.
//
// # Return Value
// 
// If you don’t want a tooltip at that location, return an empty string.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:toolTipFor:rect:tableColumn:item:mouseLocation:)
func (o NSOutlineViewDelegateObject) OutlineViewToolTipForCellRectTableColumnItemMouseLocation(outlineView INSOutlineView, cell INSCell, rect foundation.NSRect, tableColumn INSTableColumn, item objectivec.IObject, mouseLocation corefoundation.CGPoint) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:toolTipForCell:rect:tableColumn:item:mouseLocation:"), outlineView, cell, rect, tableColumn, item, mouseLocation)
	return foundation.NSStringFromID(rv).String()
	}
// Returns a Boolean value that indicates whether the outline view should
// select a given table column.
//
// outlineView: The outline view that sent the message.
//
// tableColumn: The table column.
//
// # Return Value
// 
// [true] to permit `outlineView` to select `tableColumn`, [false] to deny
// permission.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate can implement this method to disallow selection of specific
// columns.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:shouldSelect:)
func (o NSOutlineViewDelegateObject) OutlineViewShouldSelectTableColumn(outlineView INSOutlineView, tableColumn INSTableColumn) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:shouldSelectTableColumn:"), outlineView, tableColumn)
	return rv
	}
// Returns a Boolean value that indicates whether the outline view should
// select a given item.
//
// outlineView: The outline view that sent the message.
//
// item: The item.
//
// # Return Value
// 
// [true] to permit `outlineView` to select `item`, [false] to deny
// permission.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You implement this method to disallow selection of particular items.
// 
// For better performance and finer grain control over the selection, use
// [OutlineViewSelectionIndexesForProposedSelection].
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:shouldSelectItem:)
func (o NSOutlineViewDelegateObject) OutlineViewShouldSelectItem(outlineView INSOutlineView, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:shouldSelectItem:"), outlineView, item)
	return rv
	}
// Invoked to allow the delegate to modify the proposed selection.
//
// outlineView: The outline view that sent the message.
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
// Implementation of this method is optional. If implemented, this method will
// be called instead of [OutlineViewShouldSelectItem].
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:selectionIndexesForProposedSelection:)
func (o NSOutlineViewDelegateObject) OutlineViewSelectionIndexesForProposedSelection(outlineView INSOutlineView, proposedSelectionIndexes foundation.NSIndexSet) foundation.NSIndexSet {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:selectionIndexesForProposedSelection:"), outlineView, proposedSelectionIndexes)
	return foundation.NSIndexSetFromID(rv)
	}
// Returns a Boolean value that indicates whether the outline view should
// change its selection.
//
// outlineView: The outline view that sent the message.
//
// # Return Value
// 
// [true] to permit `outlineView` to change its selection (typically a row
// being edited), [false] to deny permission.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// For example, if the user is editing a cell and enters an improper value,
// the delegate can prevent the user from selecting or editing any other cells
// until a proper value has been entered into the original cell. The delegate
// can implement this method for complex validation of edited rows based on
// the values of any of their cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/selectionShouldChange(in:)
func (o NSOutlineViewDelegateObject) SelectionShouldChangeInOutlineView(outlineView INSOutlineView) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("selectionShouldChangeInOutlineView:"), outlineView)
	return rv
	}
// Invoked when `notification` is posted—that is, whenever the outline
// view’s selection changes.
//
// notification: The posted notification.
//
// # Discussion
// 
// This method is invoked as a result of posting an
// [selectionIsChangingNotification].
//
// [selectionIsChangingNotification]: https://developer.apple.com/documentation/AppKit/NSOutlineView/selectionIsChangingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineViewSelectionIsChanging(_:)
func (o NSOutlineViewDelegateObject) OutlineViewSelectionIsChanging(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineViewSelectionIsChanging:"), notification)
	}
// Invoked when the selection did change notification is posted—that is,
// immediately after the outline view’s selection has changed.
//
// notification: The posted notification.
//
// # Discussion
// 
// This method is invoked as a result of posting an
// [selectionDidChangeNotification].
//
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSOutlineView/selectionDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineViewSelectionDidChange(_:)
func (o NSOutlineViewDelegateObject) OutlineViewSelectionDidChange(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineViewSelectionDidChange:"), notification)
	}
// Informs the delegate that the cell specified by the column and item will be
// displayed.
//
// outlineView: The outline view that sent the message.
//
// cell: The cell.
//
// tableColumn: The table column.
//
// item: The item.
//
// # Discussion
// 
// The delegate can implement this method to modify `cell` to provide further
// setup for the `cell` in `tableColumn` and `item`. It is not safe to do
// drawing inside this method—you should only set up state for `cell`.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:willDisplayCell:for:item:)
func (o NSOutlineViewDelegateObject) OutlineViewWillDisplayCellForTableColumnItem(outlineView INSOutlineView, cell objectivec.IObject, tableColumn INSTableColumn, item objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:willDisplayCell:forTableColumn:item:"), outlineView, cell, tableColumn, item)
	}
// Informs the delegate that an outline view is about to display a cell used
// to draw the expansion symbol.
//
// outlineView: The outline view that sent the message.
//
// cell: The cell.
//
// tableColumn: The table column.
//
// item: The item.
//
// # Discussion
// 
// Informs the delegate that `outlineView` is about to display `cell`—an
// expandable cell (a cell that has the expansion symbol)—for the column and
// item specified by `tableColumn` and `item`. The delegate can modify cell to
// alter its display attributes.
// 
// This method is not invoked when `outlineView` is about to display a
// non-expandable cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:willDisplayOutlineCell:for:item:)
func (o NSOutlineViewDelegateObject) OutlineViewWillDisplayOutlineCellForTableColumnItem(outlineView INSOutlineView, cell objectivec.IObject, tableColumn INSTableColumn, item objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:willDisplayOutlineCell:forTableColumn:item:"), outlineView, cell, tableColumn, item)
	}
// Returns the cell to use in a given column for a given item.
//
// outlineView: The outline view that sent the message.
//
// tableColumn: The table column that requires the cell. This value can be `nil`.
//
// item: The item that requires the cell.
//
// # Return Value
// 
// The cell to use in column `tableColumn` for item `item`, or `nil`. Because
// the outline view might copy the cell, the cell must properly implement
// [copyWithZone:].
//
// [copyWithZone:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyWithZone:
//
// # Discussion
// 
// You can return a different data cell for any table column and item
// combination. Alternatively, you can return a full-width cell for the entire
// row.
// 
// If `tableColumn` is non-`nil`, you can return a cell. In most cases,
// however, you default to returning the result from
// [TableViewDataCellForTableColumnRow].
// 
// At the time of drawing, the outline view calls each row identified by
// `item` with a `nil` value for `tableColumn`. At this point, you can return
// a cell that the system can use to draw the entire row, acting as a group.
// 
// If you return a cell for the `nil` table column, prepare the other
// corresponding data source implementations and delegate methods to accept a
// `nil` value for `tableColumn`. If you don’t return a cell for the `nil`
// table column, the outline view calls this method once for each column, as
// usual.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:dataCellFor:item:)
func (o NSOutlineViewDelegateObject) OutlineViewDataCellForTableColumnItem(outlineView INSOutlineView, tableColumn INSTableColumn, item objectivec.IObject) INSCell {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:dataCellForTableColumn:item:"), outlineView, tableColumn, item)
	return NSCellFromID(rv)
	}
// Returns whether the specified item should display the outline cell (the
// disclosure triangle).
//
// outlineView: The outline view that sent the message.
//
// item: An item in the outline view.
//
// # Return Value
// 
// [true] if the outline cell should be displayed, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Returning [false] causes [FrameOfOutlineCellAtRow] to return [NSZeroRect],
// hiding the cell. In addition, the row will not be collapsible by keyboard
// shortcuts.
// 
// This method is called only for expandable rows.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:shouldShowOutlineCellForItem:)
func (o NSOutlineViewDelegateObject) OutlineViewShouldShowOutlineCellForItem(outlineView INSOutlineView, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:shouldShowOutlineCellForItem:"), outlineView, item)
	return rv
	}
// Invoked to allow the delegate to control cell expansion for a specific
// column and item.
//
// outlineView: The outline view that sent the message.
//
// tableColumn: A table column in the outline view.
//
// item: An item in the outline view.
//
// # Return Value
// 
// [true] to allow an expansion tooltip to appear in the column `tableColumn`
// for item `item`, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Cell expansion can occur when the mouse hovers over the specified cell and
// the cell contents are unable to be fully displayed within the cell. If this
// method returns [true], the full cell contents will be shown in a special
// floating tool tip view, otherwise the content is truncated.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:shouldShowCellExpansionFor:item:)
func (o NSOutlineViewDelegateObject) OutlineViewShouldShowCellExpansionForTableColumnItem(outlineView INSOutlineView, tableColumn INSTableColumn, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:shouldShowCellExpansionForTableColumn:item:"), outlineView, tableColumn, item)
	return rv
	}
// Sent to the delegate to allow or prohibit the specified column to be
// dragged to a new location.
//
// outlineView: The outline view that sent the message.
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
// If this method is not implemented, all columns are considered reorderable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:shouldReorderColumn:toColumn:)
func (o NSOutlineViewDelegateObject) OutlineViewShouldReorderColumnToColumn(outlineView INSOutlineView, columnIndex int, newColumnIndex int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:shouldReorderColumn:toColumn:"), outlineView, columnIndex, newColumnIndex)
	return rv
	}
// Invoked whenever the user moves a column in the outline view.
//
// notification: The posted notification.
//
// # Discussion
// 
// This method is invoked as a result of posting an
// [columnDidMoveNotification].
//
// [columnDidMoveNotification]: https://developer.apple.com/documentation/AppKit/NSOutlineView/columnDidMoveNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineViewColumnDidMove(_:)
func (o NSOutlineViewDelegateObject) OutlineViewColumnDidMove(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineViewColumnDidMove:"), notification)
	}
// Invoked whenever the user resizes a column in the outline view.
//
// notification: The posted notification.
//
// # Discussion
// 
// This method is invoked as a result of posting an
// [columnDidResizeNotification].
//
// [columnDidResizeNotification]: https://developer.apple.com/documentation/AppKit/NSOutlineView/columnDidResizeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineViewColumnDidResize(_:)
func (o NSOutlineViewDelegateObject) OutlineViewColumnDidResize(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineViewColumnDidResize:"), notification)
	}
// Invoked when `notification` is posted—that is, whenever the user is about
// to expand an item in the outline view.
//
// notification: The posted notification.
//
// # Discussion
// 
// This method is invoked as a result of posting an
// [itemWillExpandNotification].
//
// [itemWillExpandNotification]: https://developer.apple.com/documentation/AppKit/NSOutlineView/itemWillExpandNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineViewItemWillExpand(_:)
func (o NSOutlineViewDelegateObject) OutlineViewItemWillExpand(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineViewItemWillExpand:"), notification)
	}
// Invoked when `notification` is posted—that is, whenever the user expands
// an item in the outline view.
//
// notification: The posted notification.
//
// # Discussion
// 
// This method is invoked as a result of posting an
// [itemDidExpandNotification].
//
// [itemDidExpandNotification]: https://developer.apple.com/documentation/AppKit/NSOutlineView/itemDidExpandNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineViewItemDidExpand(_:)
func (o NSOutlineViewDelegateObject) OutlineViewItemDidExpand(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineViewItemDidExpand:"), notification)
	}
// Invoked when `notification` is posted—that is, whenever the user is about
// to collapse an item in the outline view.
//
// notification: The posted notification.
//
// # Discussion
// 
// This method is invoked as a result of posting an
// [itemWillCollapseNotification].
//
// [itemWillCollapseNotification]: https://developer.apple.com/documentation/AppKit/NSOutlineView/itemWillCollapseNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineViewItemWillCollapse(_:)
func (o NSOutlineViewDelegateObject) OutlineViewItemWillCollapse(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineViewItemWillCollapse:"), notification)
	}
// Invoked when the did collapse notification is posted—that is, whenever
// the user collapses an item in the outline view.
//
// notification: The posted notification.
//
// # Discussion
// 
// This method is invoked as a result of posting an
// [itemDidCollapseNotification].
//
// [itemDidCollapseNotification]: https://developer.apple.com/documentation/AppKit/NSOutlineView/itemDidCollapseNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineViewItemDidCollapse(_:)
func (o NSOutlineViewDelegateObject) OutlineViewItemDidCollapse(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineViewItemDidCollapse:"), notification)
	}
// Returns a Boolean value that indicates whether the outline view should
// allow editing of a given item in a given table column.
//
// outlineView: The outline view that sent the message.
//
// tableColumn: The table column.
//
// item: The item.
//
// # Return Value
// 
// [true] to permit `outlineView` to edit the cell specified by `tableColumn`
// and `item`, [false] to deny permission.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If this method returns [true], the cell may still not be editable—for
// example, if you have set up a custom [NSTextFieldCell] as a data cell, it
// must return [true] for `isEditable` to allow editing.
// 
// The delegate can implement this method to disallow editing of specific
// cells.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:shouldEdit:item:)
func (o NSOutlineViewDelegateObject) OutlineViewShouldEditTableColumnItem(outlineView INSOutlineView, tableColumn INSTableColumn, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:shouldEditTableColumn:item:"), outlineView, tableColumn, item)
	return rv
	}
// Sent to the delegate whenever the mouse button is clicked in `outlineView`
// while the cursor is in a column header `tableColumn`.
//
// outlineView: The outline view that sent the message.
//
// tableColumn: The table column.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:mouseDownInHeaderOf:)
func (o NSOutlineViewDelegateObject) OutlineViewMouseDownInHeaderOfTableColumn(outlineView INSOutlineView, tableColumn INSTableColumn) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:mouseDownInHeaderOfTableColumn:"), outlineView, tableColumn)
	}
// Sent at the time the mouse button subsequently goes up in `outlineView` and
// `tableColumn` has been “clicked” without having been dragged anywhere.
//
// outlineView: The outline view that sent the message.
//
// tableColumn: The table column.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:didClick:)
func (o NSOutlineViewDelegateObject) OutlineViewDidClickTableColumn(outlineView INSOutlineView, tableColumn INSTableColumn) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:didClickTableColumn:"), outlineView, tableColumn)
	}
// Sent at the time the mouse button goes up in `outlineView` and
// `tableColumn` has been dragged during the time the mouse button was down.
//
// outlineView: The outline view that sent the message.
//
// tableColumn: The table column.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:didDrag:)
func (o NSOutlineViewDelegateObject) OutlineViewDidDragTableColumn(outlineView INSOutlineView, tableColumn INSTableColumn) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:didDragTableColumn:"), outlineView, tableColumn)
	}
// Returns the height in points of the row containing `item`.
//
// outlineView: The outline view that sent the message.
//
// item: The row item.
//
// # Return Value
// 
// The height of the row.
//
// # Discussion
// 
// Values returned by this method should not include intercell spacing and
// must be greater than `0`.
// 
// Implement this method to support an outline view with varying row heights.
// 
// For large tables in particular, you should make sure that this method is
// efficient. [NSOutlineView] may cache the values this method returns, so if
// you would like to change a row’s height make sure to invalidate the row
// height by calling [NoteHeightOfRowsWithIndexesChanged]. [NSOutlineView]
// automatically invalidates its entire row height cache in [ReloadData] and
// [NoteNumberOfRowsChanged].
// 
// If you call [ViewAtColumnRowMakeIfNecessary] or
// [RowViewAtRowMakeIfNecessary] within your implementation of this method, an
// exception is thrown.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:heightOfRowByItem:)
func (o NSOutlineViewDelegateObject) OutlineViewHeightOfRowByItem(outlineView INSOutlineView, item objectivec.IObject) float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("outlineView:heightOfRowByItem:"), outlineView, item)
	return rv
	}
// Invoked to allow the delegate to provide custom sizing behavior when a
// column’s resize divider is double clicked.
//
// outlineView: The outline view that sent the message.
//
// column: The index of the column.
//
// # Return Value
// 
// The width of the specified column.
//
// # Discussion
// 
// By default, [NSOutlineView] iterates every row in the table, accesses a
// cell via [preparedCell(atColumn:row:)], and requests the [CellSize] to find
// the appropriate largest width to use.
// 
// For accurate results and performance, it is recommended that this method is
// implemented when using large tables. By default, large tables use a monte
// carlo simulation instead of iterating every row.
//
// [preparedCell(atColumn:row:)]: https://developer.apple.com/documentation/AppKit/NSTableView/preparedCell(atColumn:row:)
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:sizeToFitWidthOfColumn:)
func (o NSOutlineViewDelegateObject) OutlineViewSizeToFitWidthOfColumn(outlineView INSOutlineView, column int) float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("outlineView:sizeToFitWidthOfColumn:"), outlineView, column)
	return rv
	}
// Customizes an item’s tinting behavior.
//
// outlineView: The outline view to which you apply the tinting behavior.
//
// item: The item to which you apply the tinting behavior.
//
// # Return Value
// 
// Returns a new [NSTintConfiguration] object to create a particular tinting
// behavior for the item’s row, or `nil` to inherit the tinting behavior
// from the item’s parent.
//
// # Discussion
// 
// You typically use this method to customize the color for a sidebar.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:tintConfigurationForItem:)
func (o NSOutlineViewDelegateObject) OutlineViewTintConfigurationForItem(outlineView INSOutlineView, item objectivec.IObject) INSTintConfiguration {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:tintConfigurationForItem:"), outlineView, item)
	return NSTintConfigurationFromID(rv)
	}
// Returns a Boolean value that indicates whether a given cell should be
// tracked.
//
// outlineView: The outline view that sent the message.
//
// cell: The cell used to display item `item` in column `tableColumn`
//
// tableColumn: A table column in the outline view.
//
// item: An item in the outline view.
//
// # Return Value
// 
// [true] if the cell should be tracked for the item `item` in column
// `tableColumn`, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Normally, only selectable or selected cells can be tracked. If you
// implement this method, cells which are not selectable or selected can be
// tracked (and vice-versa). For example, this allows you to have a button
// cell in a table which does not change the selection, but can still be
// clicked on and tracked.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:shouldTrackCell:for:item:)
func (o NSOutlineViewDelegateObject) OutlineViewShouldTrackCellForTableColumnItem(outlineView INSOutlineView, cell INSCell, tableColumn INSTableColumn, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:shouldTrackCell:forTableColumn:item:"), outlineView, cell, tableColumn, item)
	return rv
	}
// Returns a Boolean that indicates whether a given row should be drawn in the
// “group row” style.
//
// outlineView: The outline view that sent the message.
//
// item: An item in the outline view.
//
// # Return Value
// 
// [true] to indicate a particular row should have the “group row” style
// drawn for that row, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the cell in that row is an instance of [NSTextFieldCell] and contains
// only a string value, the “group row” style attributes are automatically
// applied for that cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:isGroupItem:)
func (o NSOutlineViewDelegateObject) OutlineViewIsGroupItem(outlineView INSOutlineView, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:isGroupItem:"), outlineView, item)
	return rv
	}
// Implemented to know when a new row view is added to the table.
//
// outlineView: The outline view that sent the message.
//
// rowView: The new row view.
//
// row: The row to which the view was added.
//
// # Discussion
// 
// This delegate method is for NSView-based outline views. At this point, you
// can choose to add in extra views or modify any properties on `rowView`.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:didAdd:forRow:)
func (o NSOutlineViewDelegateObject) OutlineViewDidAddRowViewForRow(outlineView INSOutlineView, rowView INSTableRowView, row int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:didAddRowView:forRow:"), outlineView, rowView, row)
	}
// Implemented to know when a row view is removed from the table
//
// outlineView: The outline view that sent the message.
//
// rowView: The row view that was removed.
//
// row: The number of the row that was removed due to being moved offscreen, or
// `-1` if the row was removed from the table so it is no longer valid.
//
// # Discussion
// 
// The removed `rowView` may be reused by the table, so any additionally
// inserted views should be removed at this point.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:didRemove:forRow:)
func (o NSOutlineViewDelegateObject) OutlineViewDidRemoveRowViewForRow(outlineView INSOutlineView, rowView INSTableRowView, row int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:didRemoveRowView:forRow:"), outlineView, rowView, row)
	}
// implement this method to return a custom [NSTableRowView] for a particular
// item.
//
// outlineView: The outline view that sent the message.
//
// item: The item displayed by the returned table row view.
//
// # Return Value
// 
// An instance or subclass of [NSTableRowView]. If `nil` is returned, a
// [NSTableRowView] instance is created and used.
//
// # Discussion
// 
// This method, if implemented, is only invoked for NSView-based outline
// views.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:rowViewForItem:)
func (o NSOutlineViewDelegateObject) OutlineViewRowViewForItem(outlineView INSOutlineView, item objectivec.IObject) INSTableRowView {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:rowViewForItem:"), outlineView, item)
	return NSTableRowViewFromID(rv)
	}
// Implemented to return the view used to display the specified item and
// column.
//
// outlineView: The outline view that sent the message.
//
// tableColumn: The table column, or `nil` if the row is a group row.
//
// item: The item displayed by the returned view.
//
// # Return Value
// 
// The view to display the specified column and row. Returning `nil` is
// acceptable, in which case a view is not shown at that location.
//
// # Discussion
// 
// This method is required if you wish to use [NSView] objects instead of
// [NSCell] objects for the cells within an outline view. Cells and views
// cannot be mixed within the same outline view.
// 
// It is recommended that the implementation of this method first call the
// [NSTableView] method [ViewWithIdentifierOwner] passing, respectively, the
// `tableColumn` parameter’s identifier and `self` as the owner to attempt
// to reuse a view that is no longer visible. The frame of the view returned
// by this method is not important, and is automatically set by the outline
// view.
// 
// The view’s properties should be properly set up before returning the
// result.
// 
// When using Cocoa bindings, this method is optional if at least one
// identifier has been associated with the table view at design time. If this
// method is not implemented, the outline view automatically calls
// [ViewWithIdentifierOwner] with the `tableColumn` parameter’s identifier
// and the outline view’s delegate as parameters, to attempt to reuse a
// previous view or automatically unarchive a prototype associated with the
// table view.
// 
// The [autoresizingMask] of the returned view is automatically set to
// [height] to resize properly on row height changes.
//
// [autoresizingMask]: https://developer.apple.com/documentation/AppKit/NSView/autoresizingMask-swift.property
// [height]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct/height
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:viewFor:item:)
func (o NSOutlineViewDelegateObject) OutlineViewViewForTableColumnItem(outlineView INSOutlineView, tableColumn INSTableColumn, item objectivec.IObject) INSView {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:viewForTableColumn:item:"), outlineView, tableColumn, item)
	return NSViewFromID(rv)
	}
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:userCanChangeVisibilityOf:)
func (o NSOutlineViewDelegateObject) OutlineViewUserCanChangeVisibilityOfTableColumn(outlineView INSOutlineView, column INSTableColumn) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:userCanChangeVisibilityOfTableColumn:"), outlineView, column)
	return rv
	}
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDelegate/outlineView(_:userDidChangeVisibilityOf:)
func (o NSOutlineViewDelegateObject) OutlineViewUserDidChangeVisibilityOfTableColumns(outlineView INSOutlineView, columns []NSTableColumn) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:userDidChangeVisibilityOfTableColumns:"), outlineView, objectivec.IObjectSliceToNSArray(columns))
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
func (o NSOutlineViewDelegateObject) ControlIsValidObject(control INSControl, obj objectivec.IObject) bool {
	
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
func (o NSOutlineViewDelegateObject) ControlDidFailToValidatePartialStringErrorDescription(control INSControl, string_ string, error_ string) {
	
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
func (o NSOutlineViewDelegateObject) ControlDidFailToFormatStringErrorDescription(control INSControl, string_ string, error_ string) bool {
	
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
func (o NSOutlineViewDelegateObject) ControlTextShouldBeginEditing(control INSControl, fieldEditor INSText) bool {
	
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
func (o NSOutlineViewDelegateObject) ControlTextShouldEndEditing(control INSControl, fieldEditor INSText) bool {
	
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
func (o NSOutlineViewDelegateObject) ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control INSControl, textView INSTextView, words []string, charRange foundation.NSRange, index unsafe.Pointer) []string {
	
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
func (o NSOutlineViewDelegateObject) ControlTextViewDoCommandBySelector(control INSControl, textView INSTextView, commandSelector objc.SEL) bool {
	
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
func (o NSOutlineViewDelegateObject) ControlTextDidBeginEditing(obj foundation.NSNotification) {
	
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
func (o NSOutlineViewDelegateObject) ControlTextDidChange(obj foundation.NSNotification) {
	
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
func (o NSOutlineViewDelegateObject) ControlTextDidEndEditing(obj foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("controlTextDidEndEditing:"), obj)
	}

// NSOutlineViewDelegateConfig holds optional typed callbacks for [NSOutlineViewDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate
type NSOutlineViewDelegateConfig struct {

	// Handling Selection
	// SelectionIndexesForProposedSelection — Invoked to allow the delegate to modify the proposed selection.
	SelectionIndexesForProposedSelection func(outlineView NSOutlineView, proposedSelectionIndexes foundation.NSIndexSet) foundation.NSIndexSet
	// SelectionIsChanging — Invoked when `notification` is posted—that is, whenever the outline view’s selection changes.
	SelectionIsChanging func(notification foundation.NSNotification)
	// SelectionDidChange — Invoked when the selection did change notification is posted—that is, immediately after the outline view’s selection has changed.
	SelectionDidChange func(notification foundation.NSNotification)

	// Moving and Resizing Columns
	// ShouldReorderColumnToColumn — Sent to the delegate to allow or prohibit the specified column to be dragged to a new location.
	ShouldReorderColumnToColumn func(outlineView NSOutlineView, columnIndex int, newColumnIndex int) bool

	// Working with the Outline Column
	// ColumnDidMove — Invoked whenever the user moves a column in the outline view.
	ColumnDidMove func(notification foundation.NSNotification)
	// ColumnDidResize — Invoked whenever the user resizes a column in the outline view.
	ColumnDidResize func(notification foundation.NSNotification)
	// ItemWillExpand — Invoked when `notification` is posted—that is, whenever the user is about to expand an item in the outline view.
	ItemWillExpand func(notification foundation.NSNotification)
	// ItemDidExpand — Invoked when `notification` is posted—that is, whenever the user expands an item in the outline view.
	ItemDidExpand func(notification foundation.NSNotification)
	// ItemWillCollapse — Invoked when `notification` is posted—that is, whenever the user is about to collapse an item in the outline view.
	ItemWillCollapse func(notification foundation.NSNotification)
	// ItemDidCollapse — Invoked when the did collapse notification is posted—that is, whenever the user collapses an item in the outline view.
	ItemDidCollapse func(notification foundation.NSNotification)

	// Customizing Column and Row Sizes
	// SizeToFitWidthOfColumn — Invoked to allow the delegate to provide custom sizing behavior when a column’s resize divider is double clicked.
	SizeToFitWidthOfColumn func(outlineView NSOutlineView, column int) float64

	// Other Methods
	// ShouldSelectTableColumn — Returns a Boolean value that indicates whether the outline view should select a given table column.
	ShouldSelectTableColumn func(outlineView NSOutlineView, tableColumn NSTableColumn) bool
	// SelectionShouldChangeInOutlineView — Returns a Boolean value that indicates whether the outline view should change its selection.
	SelectionShouldChangeInOutlineView func(outlineView NSOutlineView) bool
	// MouseDownInHeaderOfTableColumn — Sent to the delegate whenever the mouse button is clicked in `outlineView` while the cursor is in a column header `tableColumn`.
	MouseDownInHeaderOfTableColumn func(outlineView NSOutlineView, tableColumn NSTableColumn)
	// DidClickTableColumn — Sent at the time the mouse button subsequently goes up in `outlineView` and `tableColumn` has been “clicked” without having been dragged anywhere.
	DidClickTableColumn func(outlineView NSOutlineView, tableColumn NSTableColumn)
	// DidDragTableColumn — Sent at the time the mouse button goes up in `outlineView` and `tableColumn` has been dragged during the time the mouse button was down.
	DidDragTableColumn func(outlineView NSOutlineView, tableColumn NSTableColumn)
	// DidAddRowViewForRow — Implemented to know when a new row view is added to the table.
	DidAddRowViewForRow func(outlineView NSOutlineView, rowView NSTableRowView, row int)
	// DidRemoveRowViewForRow — Implemented to know when a row view is removed from the table
	DidRemoveRowViewForRow func(outlineView NSOutlineView, rowView NSTableRowView, row int)
	UserCanChangeVisibilityOfTableColumn func(outlineView NSOutlineView, column NSTableColumn) bool
}

// NewNSOutlineViewDelegate creates an Objective-C object implementing the [NSOutlineViewDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSOutlineViewDelegateObject] satisfies the [NSOutlineViewDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate
func NewNSOutlineViewDelegate(config NSOutlineViewDelegateConfig) NSOutlineViewDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSOutlineViewDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ShouldSelectTableColumn != nil {
		fn := config.ShouldSelectTableColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineView:shouldSelectTableColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID, tableColumnID objc.ID) bool {
				outlineView := NSOutlineViewFromID(outlineViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				return fn(outlineView, tableColumn)
			},
		})
	}

	if config.SelectionIndexesForProposedSelection != nil {
		fn := config.SelectionIndexesForProposedSelection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineView:selectionIndexesForProposedSelection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID, proposedSelectionIndexesID objc.ID) objc.ID {
				outlineView := NSOutlineViewFromID(outlineViewID)
				proposedSelectionIndexes := foundation.NSIndexSetFromID(proposedSelectionIndexesID)
				return fn(outlineView, proposedSelectionIndexes).GetID()
			},
		})
	}

	if config.SelectionShouldChangeInOutlineView != nil {
		fn := config.SelectionShouldChangeInOutlineView
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("selectionShouldChangeInOutlineView:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID) bool {
				outlineView := NSOutlineViewFromID(outlineViewID)
				return fn(outlineView)
			},
		})
	}

	if config.SelectionIsChanging != nil {
		fn := config.SelectionIsChanging
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineViewSelectionIsChanging:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.SelectionDidChange != nil {
		fn := config.SelectionDidChange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineViewSelectionDidChange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ShouldReorderColumnToColumn != nil {
		fn := config.ShouldReorderColumnToColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineView:shouldReorderColumn:toColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID, columnIndex int, newColumnIndex int) bool {
				outlineView := NSOutlineViewFromID(outlineViewID)
				return fn(outlineView, columnIndex, newColumnIndex)
			},
		})
	}

	if config.ColumnDidMove != nil {
		fn := config.ColumnDidMove
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineViewColumnDidMove:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ColumnDidResize != nil {
		fn := config.ColumnDidResize
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineViewColumnDidResize:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ItemWillExpand != nil {
		fn := config.ItemWillExpand
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineViewItemWillExpand:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ItemDidExpand != nil {
		fn := config.ItemDidExpand
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineViewItemDidExpand:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ItemWillCollapse != nil {
		fn := config.ItemWillCollapse
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineViewItemWillCollapse:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ItemDidCollapse != nil {
		fn := config.ItemDidCollapse
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineViewItemDidCollapse:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.MouseDownInHeaderOfTableColumn != nil {
		fn := config.MouseDownInHeaderOfTableColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineView:mouseDownInHeaderOfTableColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID, tableColumnID objc.ID) {
				outlineView := NSOutlineViewFromID(outlineViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				fn(outlineView, tableColumn)
			},
		})
	}

	if config.DidClickTableColumn != nil {
		fn := config.DidClickTableColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineView:didClickTableColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID, tableColumnID objc.ID) {
				outlineView := NSOutlineViewFromID(outlineViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				fn(outlineView, tableColumn)
			},
		})
	}

	if config.DidDragTableColumn != nil {
		fn := config.DidDragTableColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineView:didDragTableColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID, tableColumnID objc.ID) {
				outlineView := NSOutlineViewFromID(outlineViewID)
				tableColumn := NSTableColumnFromID(tableColumnID)
				fn(outlineView, tableColumn)
			},
		})
	}

	if config.SizeToFitWidthOfColumn != nil {
		fn := config.SizeToFitWidthOfColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineView:sizeToFitWidthOfColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID, column int) float64 {
				outlineView := NSOutlineViewFromID(outlineViewID)
				return fn(outlineView, column)
			},
		})
	}

	if config.DidAddRowViewForRow != nil {
		fn := config.DidAddRowViewForRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineView:didAddRowView:forRow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID, rowViewID objc.ID, row int) {
				outlineView := NSOutlineViewFromID(outlineViewID)
				rowView := NSTableRowViewFromID(rowViewID)
				fn(outlineView, rowView, row)
			},
		})
	}

	if config.DidRemoveRowViewForRow != nil {
		fn := config.DidRemoveRowViewForRow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineView:didRemoveRowView:forRow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID, rowViewID objc.ID, row int) {
				outlineView := NSOutlineViewFromID(outlineViewID)
				rowView := NSTableRowViewFromID(rowViewID)
				fn(outlineView, rowView, row)
			},
		})
	}

	if config.UserCanChangeVisibilityOfTableColumn != nil {
		fn := config.UserCanChangeVisibilityOfTableColumn
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outlineView:userCanChangeVisibilityOfTableColumn:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outlineViewID objc.ID, columnID objc.ID) bool {
				outlineView := NSOutlineViewFromID(outlineViewID)
				column := NSTableColumnFromID(columnID)
				return fn(outlineView, column)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSOutlineViewDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSOutlineViewDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSOutlineViewDelegateObjectFromID(instance)
}


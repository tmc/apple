// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that a table view uses to provide data to a table view and to allow the editing of the table view’s data source object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource
type NSTableViewDataSource interface {
	objectivec.IObject
}



// NSTableViewDataSourceObject wraps an existing Objective-C object that conforms to the NSTableViewDataSource protocol.
type NSTableViewDataSourceObject struct {
	objectivec.Object
}
func (o NSTableViewDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSTableViewDataSourceObjectFromID constructs a [NSTableViewDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTableViewDataSourceObjectFromID(id objc.ID) NSTableViewDataSourceObject {
	return NSTableViewDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Returns the number of records managed for `aTableView` by the data source
// object.
//
// tableView: The table view that sent the message.
//
// # Return Value
// 
// The number of rows in `aTableView`.
//
// # Discussion
// 
// An instance of [NSTableView] uses this method to determine how many rows it
// should create and display. Your [NumberOfRowsInTableView] implementation is
// called very frequently, so it must be efficient.
// 
// Both view-based table views and cell-based table views must implement this
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/numberOfRows(in:)

func (o NSTableViewDataSourceObject) NumberOfRowsInTableView(tableView INSTableView) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("numberOfRowsInTableView:"), tableView)
	return rv
	}

// Called by the table view to return the data object associated with the
// specified row and column.
//
// tableView: The table view that sent the message.
//
// tableColumn: A column in `aTableView`.
//
// row: The row of the item in `aTableColumn`.
//
// # Return Value
// 
// An item in the data source in the specified table column of the view.
//
// # Discussion
// 
// [TableViewObjectValueForTableColumnRow] is called each time the table cell
// needs to be redisplayed, so it must be efficient.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:objectValueFor:row:)

func (o NSTableViewDataSourceObject) TableViewObjectValueForTableColumnRow(tableView INSTableView, tableColumn INSTableColumn, row int) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tableView:objectValueForTableColumn:row:"), tableView, tableColumn, row)
	return objectivec.Object{ID: rv}
	}

// Sets the data object for an item in the specified row and column.
//
// tableView: The table view that sent the message.
//
// object: The new value for the item.
//
// tableColumn: A column in `aTableView`.
//
// row: The row of the item in `aTableColumn`.
//
// # Discussion
// 
// This method is intended for use with cell-based table views, it must not be
// used with view-based table views. In view-based tables, use target/action
// to set each item in the view cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:setObjectValue:for:row:)

func (o NSTableViewDataSourceObject) TableViewSetObjectValueForTableColumnRow(tableView INSTableView, object objectivec.IObject, tableColumn INSTableColumn, row int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:setObjectValue:forTableColumn:row:"), tableView, object, tableColumn, row)
	}

// Called to allow the table to support multiple item dragging.
//
// tableView: The table view.
//
// row: The row.
//
// # Return Value
// 
// Returns an instance of [NSPasteboardItem] or a custom object that
// implements the [NSPasteboardWriting] protocol. Returning `nil` excludes the
// row from being dragged.
//
// # Discussion
// 
// This method is required for multi-image dragging.
// 
// If this method is implemented, then [tableView(_:writeRowsWith:to:)] will
// not be called.
//
// [tableView(_:writeRowsWith:to:)]: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:writeRowsWith:to:)
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:pasteboardWriterForRow:)

func (o NSTableViewDataSourceObject) TableViewPasteboardWriterForRow(tableView INSTableView, row int) NSPasteboardWriting {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tableView:pasteboardWriterForRow:"), tableView, row)
	return NSPasteboardWritingObjectFromID(rv)
	}

// Called by `aTableView` when the mouse button is released over a table view
// that previously decided to allow a drop.
//
// tableView: The table view that sent the message.
//
// info: An object that contains more information about this dragging operation.
//
// row: The index of the proposed target row.
//
// dropOperation: The type of dragging operation.
//
// # Return Value
// 
// [true] if the drop operation was successful, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The data source should incorporate the data from the dragging pasteboard in
// the implementation of this method. You can use the [DraggingPasteboard]
// method to get the data for the drop operation from `info`.
// 
// To accept a drop on the second row, `row` would be 2 and `operation` would
// be [NSTableViewDropOn]. To accept a drop below the last row, `row` would be
// `[aTableView numberOfRows]` and `operation` would be
// [NSTableViewDropAbove].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:acceptDrop:row:dropOperation:)

func (o NSTableViewDataSourceObject) TableViewAcceptDropRowDropOperation(tableView INSTableView, info NSDraggingInfo, row int, dropOperation NSTableViewDropOperation) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tableView:acceptDrop:row:dropOperation:"), tableView, info, row, dropOperation)
	return rv
	}

// Used by `aTableView` to determine a valid drop target.
//
// tableView: The table view that sent the message.
//
// info: An object that contains more information about this dragging operation.
//
// row: The index of the proposed target row.
//
// dropOperation: The type of dragging operation proposed.
//
// # Return Value
// 
// The dragging operation the data source will perform.
//
// # Discussion
// 
// The data source may retarget a drop by calling [SetDropRowDropOperation]
// and returning something other than [NSDragOperationNone]. A data source
// might retarget for various reasons, such as to provide better visual
// feedback when inserting into a sorted position.
// 
// To propose a drop on the second row, `row` would be 2 and `operation` would
// be [NSTableViewDropOn]. To propose a drop below the last row, `row` would
// be `[aTableView numberOfRows]` and `operation` would be
// [NSTableViewDropAbove].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:validateDrop:proposedRow:proposedDropOperation:)

func (o NSTableViewDataSourceObject) TableViewValidateDropProposedRowProposedDropOperation(tableView INSTableView, info NSDraggingInfo, row int, dropOperation NSTableViewDropOperation) NSDragOperation {
	
	rv := objc.Send[NSDragOperation](o.ID, objc.Sel("tableView:validateDrop:proposedRow:proposedDropOperation:"), tableView, info, row, dropOperation)
	return rv
	}

// Implement this method to determine when a dragging session will begin.
//
// tableView: The table view.
//
// session: The dragging session.
//
// screenPoint: The initial drag location in screen coordinates.
//
// rowIndexes: The indexes of the rows to be dragged, excluding rows that were not dragged
// due to [TableViewPasteboardWriterForRow] returning `nil`.
//
// # Discussion
// 
// Implement this method to know when the dragging session is about to begin
// and to potentially modify the dragging session.
// 
// The dragged item order will directly match the pasteboard writer array used
// to begin the dragging session with the [NSView] method
// [BeginDraggingSessionWithItemsEventSource]. Hence, the order is
// deterministic, and can be used in [TableViewAcceptDropRowDropOperation]
// when enumerating the [NSDraggingInfo] pasteboard classes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:draggingSession:willBeginAt:forRowIndexes:)

func (o NSTableViewDataSourceObject) TableViewDraggingSessionWillBeginAtPointForRowIndexes(tableView INSTableView, session INSDraggingSession, screenPoint corefoundation.CGPoint, rowIndexes foundation.NSIndexSet) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:draggingSession:willBeginAtPoint:forRowIndexes:"), tableView, session, screenPoint, rowIndexes)
	}

// Implement this method to allow the table to update dragging items as they
// are dragged over a view.
//
// tableView: The table view.
//
// draggingInfo: The dragging information.
//
// # Discussion
// 
// Required for multi-image dragging. Typically this will involve invoking
// [EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock] on
// the `draggingInfo` parameter value and setting the `draggingItem`
// object’s [ImageComponentsProvider] to a proper image based on the
// content.
// 
// For view-based table views, you can use the [NSTableCellView] method
// [DraggingImageComponents]. For cell-based tables, use the [NSCell] method
// [DraggingImageComponentsWithFrameInView].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:updateDraggingItemsForDrag:)

func (o NSTableViewDataSourceObject) TableViewUpdateDraggingItemsForDrag(tableView INSTableView, draggingInfo NSDraggingInfo) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:updateDraggingItemsForDrag:"), tableView, draggingInfo)
	}

// Implement this method to determine when a dragging session has ended.
//
// tableView: The table view.
//
// session: The dragging session.
//
// screenPoint: The ending drag location in screen coordinates.
//
// operation: The drag operation. See [NSDragOperation] for supported values.
// //
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// # Discussion
// 
// This delegate method can be used to determine when the dragging source
// operation ended at a specific location, such as the trash, by checking for
// an operation of [DragOperationDelete].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:draggingSession:endedAt:operation:)

func (o NSTableViewDataSourceObject) TableViewDraggingSessionEndedAtPointOperation(tableView INSTableView, session INSDraggingSession, screenPoint corefoundation.CGPoint, operation NSDragOperation) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:draggingSession:endedAtPoint:operation:"), tableView, session, screenPoint, operation)
	}

// Called by `aTableView` to indicate that sorting may need to be done.
//
// tableView: The table view that sent the message.
//
// oldDescriptors: An array that contains the previous descriptors.
//
// # Discussion
// 
// The data source typically sorts and reloads the data, and adjusts the
// selections accordingly. If you need to know the current sort descriptors
// and the data source doesn’t manage them itself, you can get the current
// sort descriptors by sending `aTableView` a [SortDescriptors] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:sortDescriptorsDidChange:)

func (o NSTableViewDataSourceObject) TableViewSortDescriptorsDidChange(tableView INSTableView, oldDescriptors []foundation.NSSortDescriptor) {
	
	objc.Send[struct{}](o.ID, objc.Sel("tableView:sortDescriptorsDidChange:"), tableView, objectivec.IObjectSliceToNSArray(oldDescriptors))
	}








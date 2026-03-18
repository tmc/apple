// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that an outline view calls to retrieve data and information about it from the data source delegate, and—optionally—to update data values.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource
type NSOutlineViewDataSource interface {
	objectivec.IObject
}

// NSOutlineViewDataSourceObject wraps an existing Objective-C object that conforms to the NSOutlineViewDataSource protocol.
type NSOutlineViewDataSourceObject struct {
	objectivec.Object
}
func (o NSOutlineViewDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSOutlineViewDataSourceObjectFromID constructs a [NSOutlineViewDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSOutlineViewDataSourceObjectFromID(id objc.ID) NSOutlineViewDataSourceObject {
	return NSOutlineViewDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a Boolean value that indicates whether a drop operation was
// successful.
//
// outlineView: The outline view that sent the message. `outlineView` must have previously
// allowed a drop.
//
// info: An object that contains more information about this dragging operation.
//
// item: The parent of the item over which the cursor was placed when the mouse
// button was released.
//
// index: The index of the child of `item` over which the cursor was placed when the
// mouse button was released.
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
// the implementation of this method. You can get the data for the drop
// operation from info using the [DraggingPasteboard] method.
// 
// The return value indicates success or failure of the drag operation to the
// system.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:acceptDrop:item:childIndex:)

func (o NSOutlineViewDataSourceObject) OutlineViewAcceptDropItemChildIndex(outlineView INSOutlineView, info NSDraggingInfo, item objectivec.IObject, index int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:acceptDrop:item:childIndex:"), outlineView, info, item, index)
	return rv
	}

// Returns the child item at the specified index of a given item.
//
// outlineView: The outline view that sent the message.
//
// index: The index of the child item from `item` to return.
//
// item: An item in the data source.
//
// # Return Value
// 
// The child item at `index` of `item`. If `item` is `nil`, returns the
// appropriate child item of the root object.
//
// # Discussion
// 
// Children of a given parent `item` are accessed sequentially. In order for
// the collapsed state of the outline view to remain consistent when it is
// reloaded you must always return the same object for a specified `child` and
// `item`.
// 
// # Special Considerations
// 
// The [OutlineViewChildOfItem] method is called very frequently, so it must
// be efficient.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:child:ofItem:)

func (o NSOutlineViewDataSourceObject) OutlineViewChildOfItem(outlineView INSOutlineView, index int, item objectivec.IObject) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:child:ofItem:"), outlineView, index, item)
	return objectivec.Object{ID: rv}
	}

// Implement this method to know when the given dragging session has ended.
//
// outlineView: The outline view in which the drag began.
//
// session: The dragging session that ended.
//
// screenPoint: The point onscreen at which the drag ended.
//
// operation: A mask specifying the types of drag operations permitted by the dragging
// source.
//
// # Discussion
// 
// You can implement this optional delegate method to know when the dragging
// source operation ended at a specific location, such as the trash (by
// checking for an operation of [DragOperationDelete]).
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:draggingSession:endedAt:operation:)

func (o NSOutlineViewDataSourceObject) OutlineViewDraggingSessionEndedAtPointOperation(outlineView INSOutlineView, session INSDraggingSession, screenPoint corefoundation.CGPoint, operation NSDragOperation) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:draggingSession:endedAtPoint:operation:"), outlineView, session, screenPoint, operation)
	}

// Implement this method know when the given dragging session is about to
// begin and potentially modify the dragging session.
//
// outlineView: The outline view in which the drag is about to begin.
//
// session: The dragging session that is about to begin.
//
// screenPoint: The point onscreen at which the drag is to begin.
//
// draggedItems: A array of items to be dragged, excluding items for which
// [OutlineViewPasteboardWriterForItem] returns `nil`.
//
// # Discussion
// 
// The `draggedItems` array directly matches the pasteboard writer array used
// to begin the dragging session with the [NSView] method
// [BeginDraggingSessionWithItemsEventSource]. Hence, the order is
// deterministic, and can be used in [OutlineViewAcceptDropItemChildIndex]
// when enumerating the [NSDraggingInfo] protocol’s pasteboard classes.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:draggingSession:willBeginAt:forItems:)

func (o NSOutlineViewDataSourceObject) OutlineViewDraggingSessionWillBeginAtPointForItems(outlineView INSOutlineView, session INSDraggingSession, screenPoint corefoundation.CGPoint, draggedItems foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:draggingSession:willBeginAtPoint:forItems:"), outlineView, session, screenPoint, draggedItems)
	}

// Returns a Boolean value that indicates whether the a given item is
// expandable.
//
// outlineView: The outline view that sent the message.
//
// item: An item in the data source.
//
// # Return Value
// 
// [true] if `item` can be expanded to display its children, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method may be called quite often, so it must be efficient.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:isItemExpandable:)

func (o NSOutlineViewDataSourceObject) OutlineViewIsItemExpandable(outlineView INSOutlineView, item objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("outlineView:isItemExpandable:"), outlineView, item)
	return rv
	}

// Invoked by `outlineView` to return the item for the archived `object`.
//
// outlineView: The outline view that sent the message.
//
// object: An archived representation of an item in `outlineView`’s data source.
//
// # Return Value
// 
// The unarchived item corresponding to `object`. If the item is an archived
// object, this method may return the object.
//
// # Discussion
// 
// When the outline view is restoring the saved expanded items, this method is
// called for each expanded item, to translate the archived object to an
// outline view item.
// 
// # Special Considerations
// 
// You must implement this method if you are automatically saving expanded
// items (that is, if [AutosaveExpandedItems] returns [true]).
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:itemForPersistentObject:)

func (o NSOutlineViewDataSourceObject) OutlineViewItemForPersistentObject(outlineView INSOutlineView, object objectivec.IObject) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:itemForPersistentObject:"), outlineView, object)
	return objectivec.Object{ID: rv}
	}

// Returns the number of child items encompassed by a given item.
//
// outlineView: The outline view that sent the message.
//
// item: An item in the data source.
//
// # Return Value
// 
// The number of child items encompassed by `item`. If `item` is `nil`, this
// method should return the number of children for the top-level item.
//
// # Discussion
// 
// The [OutlineViewNumberOfChildrenOfItem] method is called very frequently,
// so it must be efficient.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:numberOfChildrenOfItem:)

func (o NSOutlineViewDataSourceObject) OutlineViewNumberOfChildrenOfItem(outlineView INSOutlineView, item objectivec.IObject) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("outlineView:numberOfChildrenOfItem:"), outlineView, item)
	return rv
	}

// Invoked by `outlineView` to return the data object associated with the
// specified `item`.
//
// outlineView: The outline view that sent the message.
//
// tableColumn: A column in `outlineView`.
//
// item: An item in the data source in the specified `tableColumn` of the view.
//
// # Discussion
// 
// The item is located in the specified `tableColumn` of the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:objectValueFor:byItem:)

func (o NSOutlineViewDataSourceObject) OutlineViewObjectValueForTableColumnByItem(outlineView INSOutlineView, tableColumn INSTableColumn, item objectivec.IObject) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:objectValueForTableColumn:byItem:"), outlineView, tableColumn, item)
	return objectivec.Object{ID: rv}
	}

// Implement this method to enable the table to be an [NSDraggingSource] that
// supports dragging multiple items.
//
// outlineView: The outline view in which the drag begins.
//
// item: The item for which to return a pasteboard writer.
//
// # Return Value
// 
// A custom object that implements [NSPasteboardWriting] protocol (or simply
// use [NSPasteboardItem]).
//
// # Discussion
// 
// If this method is implemented, then [outlineView(_:writeItems:to:)] is not
// called.
//
// [outlineView(_:writeItems:to:)]: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:writeItems:to:)
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:pasteboardWriterForItem:)

func (o NSOutlineViewDataSourceObject) OutlineViewPasteboardWriterForItem(outlineView INSOutlineView, item objectivec.IObject) NSPasteboardWriting {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:pasteboardWriterForItem:"), outlineView, item)
	return NSPasteboardWritingObjectFromID(rv)
	}

// Invoked by `outlineView` to return an archived object for `item`.
//
// outlineView: The outline view that sent the message.
//
// item: The item for which to return an archived object.
//
// # Return Value
// 
// An archived representation of `item`. If the item is an archived object,
// this method may return the item.
//
// # Discussion
// 
// When the outline view is saving the expanded items, this method is called
// for each expanded item, to translate the outline view item to an archived
// object.
// 
// # Special Considerations
// 
// You must implement this method if you are automatically saving expanded
// items (that is, if [AutosaveExpandedItems] returns [true]).
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:persistentObjectForItem:)

func (o NSOutlineViewDataSourceObject) OutlineViewPersistentObjectForItem(outlineView INSOutlineView, item objectivec.IObject) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineView:persistentObjectForItem:"), outlineView, item)
	return objectivec.Object{ID: rv}
	}

// Set the data object for a given item in a given column.
//
// outlineView: The outline view that sent the message.
//
// object: The new value for the item.
//
// tableColumn: A column in `outlineView`.
//
// item: An item in the data source in the specified `tableColumn` of the view.
//
// # Discussion
// 
// The item is located in the specified `tableColumn` of the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:setObjectValue:for:byItem:)

func (o NSOutlineViewDataSourceObject) OutlineViewSetObjectValueForTableColumnByItem(outlineView INSOutlineView, object objectivec.IObject, tableColumn INSTableColumn, item objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:setObjectValue:forTableColumn:byItem:"), outlineView, object, tableColumn, item)
	}

// Invoked by an outline view to notify the data source that the descriptors
// changed and the data may need to be resorted.
//
// outlineView: The outline view that sent the message.
//
// oldDescriptors: An array that contains the previous descriptors.
//
// # Discussion
// 
// The data source typically sorts and reloads the data, and adjusts the
// selections accordingly. If you need to know the current sort descriptors
// and the data source does not itself manage them, you can get
// `outlineView`’s current sort descriptors by sending it a
// [SortDescriptors] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:sortDescriptorsDidChange:)

func (o NSOutlineViewDataSourceObject) OutlineViewSortDescriptorsDidChange(outlineView INSOutlineView, oldDescriptors []foundation.NSSortDescriptor) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:sortDescriptorsDidChange:"), outlineView, objectivec.IObjectSliceToNSArray(oldDescriptors))
	}

// Implement this method to enable the table to update dragging items as they
// are dragged over the view.
//
// outlineView: The outline view in which the drag occurs.
//
// draggingInfo: The dragging info object.
//
// # Discussion
// 
// Implementing this method is required for multi-image dragging. A typical
// implementation calls the passed-in dragging info object’s
// [EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock]
// method and sets the dragging item’s [ImageComponentsProvider] property to
// a proper image based on the content. For NSView-based table views, you can
// use the [NSTableCellView] method [DraggingImageComponents].
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:updateDraggingItemsForDrag:)

func (o NSOutlineViewDataSourceObject) OutlineViewUpdateDraggingItemsForDrag(outlineView INSOutlineView, draggingInfo NSDraggingInfo) {
	
	objc.Send[struct{}](o.ID, objc.Sel("outlineView:updateDraggingItemsForDrag:"), outlineView, draggingInfo)
	}

// Used by an outline view to determine a valid drop target.
//
// outlineView: The outline view that sent the message.
//
// info: An object that contains more information about this dragging operation.
//
// item: The proposed parent.
//
// index: The proposed child location.
//
// # Return Value
// 
// A value that indicates which dragging operation the data source will
// perform.
//
// # Discussion
// 
// Based on the mouse position, the outline view will suggest a proposed drop
// location. The data source may “retarget” a drop if desired by calling
// [SetDropItemDropChildIndex] and returning something other than
// [NSDragOperationNone]. You may choose to retarget for various reasons (for
// example, for better visual feedback when inserting into a sorted position).
// 
// Implementation of this method is optional.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineViewDataSource/outlineView(_:validateDrop:proposedItem:proposedChildIndex:)

func (o NSOutlineViewDataSourceObject) OutlineViewValidateDropProposedItemProposedChildIndex(outlineView INSOutlineView, info NSDraggingInfo, item objectivec.IObject, index int) NSDragOperation {
	
	rv := objc.Send[NSDragOperation](o.ID, objc.Sel("outlineView:validateDrop:proposedItem:proposedChildIndex:"), outlineView, info, item, index)
	return rv
	}


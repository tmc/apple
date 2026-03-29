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

// A set of methods that you use to manage the behavior of a collection view.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate
type NSCollectionViewDelegate interface {
	objectivec.IObject
}

// NSCollectionViewDelegateObject wraps an existing Objective-C object that conforms to the NSCollectionViewDelegate protocol.
type NSCollectionViewDelegateObject struct {
	objectivec.Object
}
func (o NSCollectionViewDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCollectionViewDelegateObjectFromID constructs a [NSCollectionViewDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCollectionViewDelegateObjectFromID(id objc.ID) NSCollectionViewDelegateObject {
	return NSCollectionViewDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate to approve the pending selection of items.
//
// collectionView: The collection view making the request.
//
// indexPaths: The set of [NSIndexPath] objects corresponding to the items selected by the
// user.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Return Value
// 
// The set of [NSIndexPath] objects corresponding to the items that you want
// to be selected. If you do not want any items selected, return an empty set.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// Use this method to approve or modify the items that the user tries to
// select. During interactive selection, the collection view calls this method
// whenever the user selects new items. Your implementation of the method can
// return the proposed set of index paths as-is or modify the set before
// returning it. You might modify the set to disallow the selection of
// specific items or specific combinations of items.
// 
// This method is not called when you set the selection programmatically using
// the methods of the [NSCollectionView] class. If you do not implement this
// method, the collection view selects the items specified by the `indexPaths`
// parameter.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:shouldSelectItemsAt:)
func (o NSCollectionViewDelegateObject) CollectionViewShouldSelectItemsAtIndexPaths(collectionView INSCollectionView, indexPaths foundation.INSSet) foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("collectionView:shouldSelectItemsAtIndexPaths:"), collectionView, indexPaths)
	return foundation.NSSetFromID(rv)
	}
// Notifies the delegate object that one or more items were selected.
//
// collectionView: The collection view notifying you of the selection change.
//
// indexPaths: The set of [NSIndexPath] objects corresponding to the items that are now
// selected.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// After the user successfully selects one or more items, the collection view
// calls this method to let you know that the selection has been made. Use
// this method to respond to the selection change and to make any necessary
// adjustments to your content or the collection view.
// 
// This method is not called when you set the selection programmatically using
// the methods of the [NSCollectionView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:didSelectItemsAt:)
func (o NSCollectionViewDelegateObject) CollectionViewDidSelectItemsAtIndexPaths(collectionView INSCollectionView, indexPaths foundation.INSSet) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:didSelectItemsAtIndexPaths:"), collectionView, indexPaths)
	}
// Asks the delegate object to approve the pending deselection of items.
//
// collectionView: The collection view making the request.
//
// indexPaths: The set of [NSIndexPath] objects corresponding to the items deselected by
// the user.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Return Value
// 
// The set of [NSIndexPath] objects corresponding to the items that you want
// to be deselected. If you do not want any items deselected, return an empty
// set.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// Use this method to approve or modify the items that the user tries to
// deselect. During interactive selection, the collection view calls this
// method whenever the user deselects items. Your implementation of the method
// can return the proposed set of index paths as-is or modify the set before
// returning it. You might modify the set to disallow the deselection of
// specific items.
// 
// This method is not called when you set the selection programmatically using
// the methods of the [NSCollectionView] class. If you do not implement this
// method, the collection view deselects the items specified by the
// `indexPaths` parameter.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:shouldDeselectItemsAt:)
func (o NSCollectionViewDelegateObject) CollectionViewShouldDeselectItemsAtIndexPaths(collectionView INSCollectionView, indexPaths foundation.INSSet) foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("collectionView:shouldDeselectItemsAtIndexPaths:"), collectionView, indexPaths)
	return foundation.NSSetFromID(rv)
	}
// Notifies the delegate object that one or more items were deselected.
//
// collectionView: The collection view notifying you of the selection change.
//
// indexPaths: The set of [NSIndexPath] objects corresponding to the items that were
// deselected.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// After the user successfully deselects one or more items, the collection
// view calls this method to let you know that the items are no longer
// selected. Use this method to respond to the selection change and to make
// any necessary adjustments to your content or the collection view.
// 
// This method is not called when you set the selection programmatically using
// the methods of the [NSCollectionView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:didDeselectItemsAt:)
func (o NSCollectionViewDelegateObject) CollectionViewDidDeselectItemsAtIndexPaths(collectionView INSCollectionView, indexPaths foundation.INSSet) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:didDeselectItemsAtIndexPaths:"), collectionView, indexPaths)
	}
// Asks the delegate to approve the pending highlighting of the specified
// items.
//
// collectionView: The collection view making the request.
//
// indexPaths: The set of [NSIndexPath] objects corresponding to the items being
// highlighted.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// highlightState: The new highlight state for the items.
//
// # Return Value
// 
// The set of [NSIndexPath] objects corresponding to the items that you want
// to receive the specified highlight. If you do not want any items to receive
// the specified highlight state, return an empty set.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// Use this method to approve or modify the set of items targeted to receive
// the specified highlight state. During interactive selection or dragging,
// the collection view calls this method when actions occur that would affect
// the highlight state of items. Your implementation of the method can return
// the proposed set of index paths as-is or modify the set and disallow the
// highlighting of some or all of the items. Removing items from the set
// suppresses the corresponding actions, such as selecting the item or showing
// its eligibility as a drop target.
// 
// If you do not implement this method, the collection view updates the
// highlight state for the items specified by the `indexPaths` parameter.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:shouldChangeItemsAt:to:)
func (o NSCollectionViewDelegateObject) CollectionViewShouldChangeItemsAtIndexPathsToHighlightState(collectionView INSCollectionView, indexPaths foundation.INSSet, highlightState NSCollectionViewItemHighlightState) foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:"), collectionView, indexPaths, highlightState)
	return foundation.NSSetFromID(rv)
	}
// Notifies the delegate that the highlight state of the specified items
// changed.
//
// collectionView: The collection view notifying you of the highlight change.
//
// indexPaths: The set of [NSIndexPath] objects corresponding to the items whose highlight
// state changed.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// highlightState: The new highlight state of the items.
//
// # Discussion
// 
// After the highlight state of one or more items changes successfully, the
// collection view calls this method to report the change. Use this method to
// respond to the change and to make any necessary adjustments to your content
// or the collection view.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:didChangeItemsAt:to:)
func (o NSCollectionViewDelegateObject) CollectionViewDidChangeItemsAtIndexPathsToHighlightState(collectionView INSCollectionView, indexPaths foundation.INSSet, highlightState NSCollectionViewItemHighlightState) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:didChangeItemsAtIndexPaths:toHighlightState:"), collectionView, indexPaths, highlightState)
	}
// Notifies the delegate that the specified item is about to be displayed by
// the collection view.
//
// collectionView: The collection view that is adding the item.
//
// item: The item being added.
//
// indexPath: The index path of the item.
//
// # Discussion
// 
// The collection view calls this method before adding new items to its
// content. You can use this method to track the addition of items and perform
// related tasks.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:willDisplay:forRepresentedObjectAt:)
func (o NSCollectionViewDelegateObject) CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath(collectionView INSCollectionView, item INSCollectionViewItem, indexPath foundation.INSIndexPath) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:"), collectionView, item, indexPath)
	}
// Notifies the delegate that the specified item was removed from the
// collection view.
//
// collectionView: The collection view that removed the item.
//
// item: The item that was removed.
//
// indexPath: The index path of the item.
//
// # Discussion
// 
// The collection view calls this method after removing an item from its
// content. You can use this method to track the removal of items and perform
// related tasks.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:didEndDisplaying:forRepresentedObjectAt:)
func (o NSCollectionViewDelegateObject) CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView INSCollectionView, item INSCollectionViewItem, indexPath foundation.INSIndexPath) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"), collectionView, item, indexPath)
	}
// Notifies the delegate that the specified supplementary view is about to be
// displayed by the collection view.
//
// collectionView: The collection view that is adding the supplementary view.
//
// view: The supplementary view being added.
//
// elementKind: The type of the supplementary view. Layouts are responsible for defining
// the types of supplementary views they support.
//
// indexPath: The index path associated with the supplementary view.
//
// # Discussion
// 
// The collection view calls this method before adding new supplementary views
// to its content. You can use this method to track the addition of those
// views and perform related tasks.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:willDisplaySupplementaryView:forElementKind:at:)
func (o NSCollectionViewDelegateObject) CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath(collectionView INSCollectionView, view INSView, elementKind NSCollectionViewSupplementaryElementKind, indexPath foundation.INSIndexPath) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:"), collectionView, view, objc.String(string(elementKind)), indexPath)
	}
// Notifies the delegate that the specified supplementary view was removed
// from the collection view.
//
// collectionView: The collection view that removed the view.
//
// view: The supplementary view that was removed.
//
// elementKind: The type of the supplementary view. Layouts are responsible for defining
// the types of supplementary views they support.
//
// indexPath: The index path associated with the supplementary view.
//
// # Discussion
// 
// The collection view calls this method after removing a supplementary view
// from its content. You can use this method to track the removal of views and
// perform related tasks.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:didEndDisplayingSupplementaryView:forElementOfKind:at:)
func (o NSCollectionViewDelegateObject) CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath(collectionView INSCollectionView, view INSView, elementKind NSCollectionViewSupplementaryElementKind, indexPath foundation.INSIndexPath) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:"), collectionView, view, objc.String(string(elementKind)), indexPath)
	}
// Returns the transition layout object to use when performing an animated
// change between different layouts.
//
// collectionView: The collection view whose layout object is changing.
//
// fromLayout: The current layout object of the collection view. This is the starting
// point for the transition.
//
// toLayout: The new layout for the collection view.
//
// # Return Value
// 
// The collection view transition layout object to use to perform the
// transition.
//
// # Discussion
// 
// When changing layouts for a collection view, you can use this method to
// customize the transition layout object used to make the change. A
// transition layout object lets you customize the behavior of collection view
// elements when transitioning from one layout to the next. Normally,
// transitioning between layouts causes the assorted items and views to
// animate from their current locations directly to their new locations. By
// returning a custom transition object, you could have those elements follow
// a nonlinear path, use a different timing algorithm, or move items in
// response to touch events.
// 
// If you do not implement this method in your delegate object, the collection
// view uses a standard [UICollectionViewTransitionLayout] object for the
// transition.
// 
// # Special Considerations
// 
// In OS X 10.11, this method is never called by the collection view.
//
// [UICollectionViewTransitionLayout]: https://developer.apple.com/documentation/UIKit/UICollectionViewTransitionLayout
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:transitionLayoutForOldLayout:newLayout:)
func (o NSCollectionViewDelegateObject) CollectionViewTransitionLayoutForOldLayoutNewLayout(collectionView INSCollectionView, fromLayout INSCollectionViewLayout, toLayout INSCollectionViewLayout) INSCollectionViewTransitionLayout {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("collectionView:transitionLayoutForOldLayout:newLayout:"), collectionView, fromLayout, toLayout)
	return NSCollectionViewTransitionLayoutFromID(rv)
	}
// Returns a Boolean indicating whether a drag operation involving the
// specified items can begin.
//
// collectionView: The collection view making the request.
//
// indexPaths: The index paths of the items about to be dragged.
//
// event: The mouse-down event that began the drag operation.
//
// # Return Value
// 
// [true] if the drag operation can begin or [false] if it cannot.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If you do not implement this method and your collection view has only one
// section, the collection view calls the legacy
// [CollectionViewCanDragItemsAtIndexesWithEvent] method. If you do not
// implement either method, the collection view assumes a return value of
// [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:canDragItemsAt:with:)-49wix
func (o NSCollectionViewDelegateObject) CollectionViewCanDragItemsAtIndexPathsWithEvent(collectionView INSCollectionView, indexPaths foundation.INSSet, event INSEvent) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("collectionView:canDragItemsAtIndexPaths:withEvent:"), collectionView, indexPaths, event)
	return rv
	}
// Provides the pasteboard writer for the item at the specified index path.
//
// collectionView: The collection view making the request.
//
// indexPath: The index path of the item requiring a pasteboard writer.
//
// # Return Value
// 
// The pasteboard writer object to use for managing the item data. Return
// `nil` to prevent the collection view from dragging the item.
//
// # Discussion
// 
// You must implement this method or the [collectionView(_:writeItemsAt:to:)]
// method to support drag operations. The collection view calls this method in
// preference over the [collectionView(_:writeItemsAt:to:)] method if both are
// implemented. If your app supports multi-image drag and drop, you must
// implement this method.
// 
// The collection view calls this method for each item involved in the drag
// operation after it has determined that a drag should begin but before the
// drag operation has started. Your implementation of this method should
// create and return the pasteboard writer—an object conforming to the
// [NSPasteboardWriting] protocol—to use for providing the item’s data.
// Using the object you provide, the collection view creates an
// [NSDraggingItem] object for you and configures its [DraggingFrame] and
// [ImageComponents] properties for you using information from the item at the
// specified index path.
// 
// If you implement this method, the collection view does not call the
// [CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset] of your
// delegate or the [DraggingImageForItemsAtIndexPathsWithEventOffset] method
// of [NSCollectionView].
//
// [collectionView(_:writeItemsAt:to:)]: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:writeItemsAt:to:)-23ozm
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:pasteboardWriterForItemAt:)-5eyyl
func (o NSCollectionViewDelegateObject) CollectionViewPasteboardWriterForItemAtIndexPath(collectionView INSCollectionView, indexPath foundation.INSIndexPath) NSPasteboardWriting {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("collectionView:pasteboardWriterForItemAtIndexPath:"), collectionView, indexPath)
	return NSPasteboardWritingObjectFromID(rv)
	}
// Creates and returns a drag image to represent the specified items during a
// drag.
//
// collectionView: The collection view making the request.
//
// indexPaths: The index paths of the items being dragged.
//
// event: The mouse-down event that began the drag operation. You can use the mouse
// location when determining what value to return in the `dragImageOffset`
// parameter.
//
// dragImageOffset: The offset value to use when positioning the image. On input, the point is
// [NSZeroPoint], which centers the returned image under the mouse. Your
// method can return a different point that repositions the drag image by the
// specified offset values.
// //
// [NSZeroPoint]: https://developer.apple.com/documentation/Foundation/NSZeroPoint
//
// # Return Value
// 
// The image to use for the dragged items.
//
// # Discussion
// 
// Your implementation of this method should create an appropriate image to
// use during the drag operation. The collection view places the center of
// your image at the current mouse location. Update the value in the
// `dragImageOffset` parameter to shift the position of your image by the
// specified amount.
// 
// If you do not implement this method, the collection view uses the drag
// image returned by the [DraggingImageForItemsAtIndexPathsWithEventOffset]
// method instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:draggingImageForItemsAt:with:offset:)-898js
func (o NSCollectionViewDelegateObject) CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset(collectionView INSCollectionView, indexPaths foundation.INSSet, event INSEvent, dragImageOffset foundation.NSPoint) INSImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:"), collectionView, indexPaths, event, dragImageOffset)
	return NSImageFromID(rv)
	}
// Notifies your delegate that a drag session is about to begin.
//
// collectionView: The collection view notifying your delegate object.
//
// session: The dragging session that is about to begin.
//
// screenPoint: The starting point (in screen coordinates) for the drag operation.
//
// indexPaths: The index paths of the items being dragged.
//
// # Discussion
// 
// You can use this method to modify the dragging session or to perform other
// tasks related to the beginning of a drag session.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:draggingSession:willBeginAt:forItemsAt:)-68x2y
func (o NSCollectionViewDelegateObject) CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths(collectionView INSCollectionView, session INSDraggingSession, screenPoint corefoundation.CGPoint, indexPaths foundation.INSSet) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:"), collectionView, session, screenPoint, indexPaths)
	}
// Notifies your delegate that a drag session ended.
//
// collectionView: The collection view notifying your delegate object.
//
// session: The dragging session that ended.
//
// screenPoint: The end point (in screen coordinates) for the drag operation.
//
// operation: The operation that was performed. Use this value to determine how the
// operation ended. For example, for content that was dragged to the trash,
// the operation type would be [DragOperationDelete].
//
// # Discussion
// 
// You can use this method to perform tasks related to the ending of a drag
// session.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:draggingSession:endedAt:dragOperation:)
func (o NSCollectionViewDelegateObject) CollectionViewDraggingSessionEndedAtPointDragOperation(collectionView INSCollectionView, session INSDraggingSession, screenPoint corefoundation.CGPoint, operation NSDragOperation) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:draggingSession:endedAtPoint:dragOperation:"), collectionView, session, screenPoint, operation)
	}
// Asks your delegate to update the dragging items during a drag operation.
//
// collectionView: The collection view asking you to update the dragging items.
//
// draggingInfo: The current information for the drag operation. Use this object to iterate
// over the dragging items.
//
// # Discussion
// 
// You can use this method to update the current drag items while a drag is in
// progress. Updating the drag items is optional, but you might use this
// method to change the image for an item. For example, you might change the
// image when the mouse hovers over a particular part of the collection view.
// Use the
// [EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock]
// method of the `draggingInfo` parameter to iterate over the drag items and
// update them as appropriate.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:updateDraggingItemsForDrag:)
func (o NSCollectionViewDelegateObject) CollectionViewUpdateDraggingItemsForDrag(collectionView INSCollectionView, draggingInfo NSDraggingInfo) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:updateDraggingItemsForDrag:"), collectionView, draggingInfo)
	}
// Validates whether a drop operation is possible at the specified location.
//
// collectionView: The collection view asking you to validate the drop operation.
//
// draggingInfo: The information about the drag operation.
//
// proposedDropIndexPath: The index path at which the drop would occur. This parameter is passed
// by-reference and can be modified to change the proposed index path.
//
// proposedDropOperation: The type of drop operation being proposed. This parameter is passed
// by-reference and can be modified to change the drop operation type.
//
// # Return Value
// 
// A value that indicates which dragging operation to perform. Return
// [NSDragOperationNone] to disallow a drop at the proposed location.
//
// [NSDragOperationNone]: https://developer.apple.com/documentation/AppKit/NSDragOperation/NSDragOperationNone
//
// # Discussion
// 
// Although implementation of this method is optional, you must implement it
// to support drops onto the associated collection view. You must also call
// the collection view’s [RegisterForDraggedTypes] method to register the
// types of drops it supports. If you do not perform both of these actions,
// the collection view does not accept drops.
// 
// When an interactive drag operation occurs, the collection view calls this
// method to determine whether the current mouse location is a valid place to
// drop the content. This method may be called many times during the course of
// the drag operation. Your implementation should look at the proposed
// location and return a constant that reflects how the drop would be handled.
// 
// While validating the drop location, you can suggest a better drop location
// by updating the values in the `proposedDropIndexPath` and
// `proposedDropOperation` parameters. For example, you might suggest dropping
// the content before the specified item instead of on it. The collection view
// sets the `proposedDropOperation` parameter to
// [NSCollectionView.DropOperation.on] when the mouse is closer to the middle
// of an item than to its edges; otherwise, it sets the parameter to
// [NSCollectionView.DropOperation.before].
//
// [NSCollectionView.DropOperation.before]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation/before
// [NSCollectionView.DropOperation.on]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation/on
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:validateDrop:proposedIndexPath:dropOperation:)
func (o NSCollectionViewDelegateObject) CollectionViewValidateDropProposedIndexPathDropOperation(collectionView INSCollectionView, draggingInfo NSDraggingInfo, proposedDropIndexPath objectivec.IObject, proposedDropOperation NSCollectionViewDropOperation) NSDragOperation {
	rv := objc.Send[NSDragOperation](o.ID, objc.Sel("collectionView:validateDrop:proposedIndexPath:dropOperation:"), collectionView, draggingInfo, proposedDropIndexPath, proposedDropOperation)
	return rv
	}
// Incorporates the dropped content into the collection view.
//
// collectionView: The collection view receiving the dropped content.
//
// draggingInfo: The information about the drag operation.
//
// indexPath: The index path at which the drop occurred. Use this location as the
// insertion point for the content.
//
// dropOperation: The type of drop operation to perform.
//
// # Return Value
// 
// [true] if the drop operation should be accepted or [false] if it should be
// rejected.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The collection view calls this method when the user releases the mouse
// button while it is over a valid drop target. This method is called after
// the [CollectionViewValidateDropProposedIndexPathDropOperation] method
// validates that dropping the content at the specified location is possible.
// You must implement this method to accept the dropped content and
// incorporate it into the collection view.
// 
// In your implementation, use the information in the `draggingInfo` parameter
// to retrieve the data, update your data source object, and insert the
// appropriate items into the collection view. The dropped data is stored in
// the [DraggingPasteboard] property of the dragging information object.
// 
// If the [AnimatesToDestination] property of the dragging information is
// [true], update the image and frame for each dragged item to its new
// location in the collection view. You can enumerate the list of dragged
// items using the
// [EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock]
// method of the dragging information object.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:acceptDrop:indexPath:dropOperation:)
func (o NSCollectionViewDelegateObject) CollectionViewAcceptDropIndexPathDropOperation(collectionView INSCollectionView, draggingInfo NSDraggingInfo, indexPath foundation.INSIndexPath, dropOperation NSCollectionViewDropOperation) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("collectionView:acceptDrop:indexPath:dropOperation:"), collectionView, draggingInfo, indexPath, dropOperation)
	return rv
	}
// Returns a Boolean indicating whether the collection view can begin dragging
// the specified items.
//
// collectionView: The collection view containing the items to be dragged.
//
// indexes: The indexes of the items to be dragged.
//
// event: The mouse event that initiated the drag action.
//
// # Return Value
// 
// [true] if the collection view can begin the drag operation for the
// specified items or [false] if it cannot.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Implement this method when you want selective control over the initiation
// of drag operations. In your implementation, use the provided information to
// determine whether the drag operation should occur and return the
// appropriate return value. For example, you might return [false] if your
// interface does not allow the user to drag the specified items.
// 
// If you do not implement this method in your delegate object, the collection
// view assumes a return value of [true] and begins the drag operation.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:canDragItemsAt:with:)-39rjh
func (o NSCollectionViewDelegateObject) CollectionViewCanDragItemsAtIndexesWithEvent(collectionView INSCollectionView, indexes foundation.NSIndexSet, event INSEvent) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("collectionView:canDragItemsAtIndexes:withEvent:"), collectionView, indexes, event)
	return rv
	}
// Provides the pasteboard writer for the item at the specified index
//
// collectionView: The collection view making the request.
//
// index: The index of the item requiring a pasteboard writer.
//
// # Return Value
// 
// The pasteboard writer object to use for managing the item data. Return
// `nil` to prevent the collection view from dragging the item.
//
// # Discussion
// 
// The collection view calls this method for each item involved in the drag
// operation after it has determined that a drag should begin but before the
// drag operation has started. Your implementation of this method should
// create and return the pasteboard writer—an object conforming to the
// [NSPasteboardWriting] protocol—to use for providing the item’s data.
// Using the object you provide, the collection view creates an
// [NSDraggingItem] object for you and configures its [DraggingFrame] and
// [ImageComponents] properties for you using information from the item at the
// specified index path.
// 
// If you implement this method, the collection view does not call the
// [CollectionViewDraggingImageForItemsAtIndexesWithEventOffset] of your
// delegate or the [DraggingImageForItemsAtIndexesWithEventOffset] method of
// [NSCollectionView].
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:pasteboardWriterForItemAt:)-7ldvs
func (o NSCollectionViewDelegateObject) CollectionViewPasteboardWriterForItemAtIndex(collectionView INSCollectionView, index uint) NSPasteboardWriting {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("collectionView:pasteboardWriterForItemAtIndex:"), collectionView, index)
	return NSPasteboardWritingObjectFromID(rv)
	}
// Creates and returns a drag image to represent the specified items during a
// drag.
//
// collectionView: The collection view making the request.
//
// indexes: The indexes of the items being dragged.
//
// event: The mouse-down event that initiated the drag.
//
// dragImageOffset: An in/out parameter that is initially set to [NSZeroPoint], which causes
// the image to be centered under the mouse. The value can be modified to
// reposition the returned image.
// //
// [NSZeroPoint]: https://developer.apple.com/documentation/Foundation/NSZeroPoint
//
// # Return Value
// 
// The image to use for the dragged items.
//
// # Discussion
// 
// If the delegate does not implement this method, the collection view uses
// the image returned by [DraggingImageForItemsAtIndexesWithEventOffset].
// 
// You do not need to implement this method for your collection view to be a
// drag source.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:draggingImageForItemsAt:with:offset:)-4yvk5
func (o NSCollectionViewDelegateObject) CollectionViewDraggingImageForItemsAtIndexesWithEventOffset(collectionView INSCollectionView, indexes foundation.NSIndexSet, event INSEvent, dragImageOffset foundation.NSPoint) INSImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("collectionView:draggingImageForItemsAtIndexes:withEvent:offset:"), collectionView, indexes, event, dragImageOffset)
	return NSImageFromID(rv)
	}
// Notifies your delegate that a drag session is about to begin.
//
// collectionView: The collection view notifying your delegate object.
//
// session: The dragging session that is about to begin.
//
// screenPoint: The starting point (in screen coordinates) for the drag operation.
//
// indexes: The indexes of the items being dragged.
//
// # Discussion
// 
// You can use this method to modify the dragging session or to perform other
// tasks related to the beginning of a drag session.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:draggingSession:willBeginAt:forItemsAt:)-cpuq
func (o NSCollectionViewDelegateObject) CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes(collectionView INSCollectionView, session INSDraggingSession, screenPoint corefoundation.CGPoint, indexes foundation.NSIndexSet) {
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:"), collectionView, session, screenPoint, indexes)
	}
// Validates the specified location to see if it is a valid drop target.
//
// collectionView: The collection view that send the message.
//
// draggingInfo: An object containing details about this dragging operation.
//
// proposedDropIndex: The proposed drop index. This parameter is passed by-reference and can be
// modified retarget the drop operation.
//
// proposedDropOperation: The proposed drop operation. This parameter is passed by-reference and can
// be modified to change the drop operation.
//
// # Return Value
// 
// A value that indicates which dragging operation to perform. Return
// [NSDragOperationNone] to disallow the drop.
//
// [NSDragOperationNone]: https://developer.apple.com/documentation/AppKit/NSDragOperation/NSDragOperationNone
//
// # Discussion
// 
// Based on the mouse position, the collection view will suggest a proposed
// index and drop operation. These values are in/out parameters and can be
// changed by the delegate to retarget the drop operation.
// 
// The collection view will propose [NSCollectionViewDropOn] when the dragging
// location is closer to the middle of the item than either of its edges.
// Otherwise, it will propose [NSCollectionViewDropBefore]. You may override
// this default behavior by changing `proposedDropOperation` or
// `proposedDropIndex`.
// 
// To receive drag messages, you must first send [RegisterForDraggedTypes] to
// the collection view with the drag types you want to support.
// 
// You must implement this method for your collection view to be a drag
// destination.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:validateDrop:proposedIndex:dropOperation:)
func (o NSCollectionViewDelegateObject) CollectionViewValidateDropProposedIndexDropOperation(collectionView INSCollectionView, draggingInfo NSDraggingInfo, proposedDropIndex unsafe.Pointer, proposedDropOperation NSCollectionViewDropOperation) NSDragOperation {
	rv := objc.Send[NSDragOperation](o.ID, objc.Sel("collectionView:validateDrop:proposedIndex:dropOperation:"), collectionView, draggingInfo, proposedDropIndex, proposedDropOperation)
	return rv
	}
// Invoked when the mouse is released over a collection view that previously
// allowed a drop.
//
// collectionView: The collection view that send the message.
//
// draggingInfo: An object that contains more information about this dragging operation.
//
// index: The index of the proposed drop item.
//
// dropOperation: The type of dragging operation.
//
// # Return Value
// 
// [true] if the drop operation should be accepted, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is called when the mouse is released over a collection view
// that previously decided to allow a drop via the
// [CollectionViewValidateDropProposedIndexDropOperation] method. At this
// time, the delegate should incorporate the data from the dragging pasteboard
// and update the collection view’s contents.
// 
// You must implement this method for your collection view to be a drag
// destination
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDelegate/collectionView(_:acceptDrop:index:dropOperation:)
func (o NSCollectionViewDelegateObject) CollectionViewAcceptDropIndexDropOperation(collectionView INSCollectionView, draggingInfo NSDraggingInfo, index int, dropOperation NSCollectionViewDropOperation) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("collectionView:acceptDrop:index:dropOperation:"), collectionView, draggingInfo, index, dropOperation)
	return rv
	}

// NSCollectionViewDelegateConfig holds optional typed callbacks for [NSCollectionViewDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nscollectionviewdelegate
type NSCollectionViewDelegateConfig struct {

	// Handling Layout Changes
	// TransitionLayoutForOldLayoutNewLayout — Returns the transition layout object to use when performing an animated change between different layouts.
	TransitionLayoutForOldLayoutNewLayout func(collectionView NSCollectionView, fromLayout NSCollectionViewLayout, toLayout NSCollectionViewLayout) NSCollectionViewTransitionLayout

	// Other Methods
	// ShouldSelectItemsAtIndexPaths — Asks the delegate to approve the pending selection of items.
	ShouldSelectItemsAtIndexPaths func(collectionView NSCollectionView, indexPaths foundation.INSSet) foundation.INSSet
	// DidSelectItemsAtIndexPaths — Notifies the delegate object that one or more items were selected.
	DidSelectItemsAtIndexPaths func(collectionView NSCollectionView, indexPaths foundation.INSSet)
	// ShouldDeselectItemsAtIndexPaths — Asks the delegate object to approve the pending deselection of items.
	ShouldDeselectItemsAtIndexPaths func(collectionView NSCollectionView, indexPaths foundation.INSSet) foundation.INSSet
	// DidDeselectItemsAtIndexPaths — Notifies the delegate object that one or more items were deselected.
	DidDeselectItemsAtIndexPaths func(collectionView NSCollectionView, indexPaths foundation.INSSet)
	// ShouldChangeItemsAtIndexPathsToHighlightState — Asks the delegate to approve the pending highlighting of the specified items.
	ShouldChangeItemsAtIndexPathsToHighlightState func(collectionView NSCollectionView, indexPaths foundation.INSSet, highlightState NSCollectionViewItemHighlightState) foundation.INSSet
	// DidChangeItemsAtIndexPathsToHighlightState — Notifies the delegate that the highlight state of the specified items changed.
	DidChangeItemsAtIndexPathsToHighlightState func(collectionView NSCollectionView, indexPaths foundation.INSSet, highlightState NSCollectionViewItemHighlightState)
	// WillDisplayItemForRepresentedObjectAtIndexPath — Notifies the delegate that the specified item is about to be displayed by the collection view.
	WillDisplayItemForRepresentedObjectAtIndexPath func(collectionView NSCollectionView, item NSCollectionViewItem, indexPath foundation.NSIndexPath)
	// DidEndDisplayingItemForRepresentedObjectAtIndexPath — Notifies the delegate that the specified item was removed from the collection view.
	DidEndDisplayingItemForRepresentedObjectAtIndexPath func(collectionView NSCollectionView, item NSCollectionViewItem, indexPath foundation.NSIndexPath)
	// WillDisplaySupplementaryViewForElementKindAtIndexPath — Notifies the delegate that the specified supplementary view is about to be displayed by the collection view.
	WillDisplaySupplementaryViewForElementKindAtIndexPath func(collectionView NSCollectionView, view NSView, elementKind NSCollectionViewSupplementaryElementKind, indexPath foundation.NSIndexPath)
	// DidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath — Notifies the delegate that the specified supplementary view was removed from the collection view.
	DidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath func(collectionView NSCollectionView, view NSView, elementKind NSCollectionViewSupplementaryElementKind, indexPath foundation.NSIndexPath)
	// CanDragItemsAtIndexPathsWithEvent — Returns a Boolean indicating whether a drag operation involving the specified items can begin.
	CanDragItemsAtIndexPathsWithEvent func(collectionView NSCollectionView, indexPaths foundation.INSSet, event NSEvent) bool
	// DraggingImageForItemsAtIndexPathsWithEventOffset — Creates and returns a drag image to represent the specified items during a drag.
	DraggingImageForItemsAtIndexPathsWithEventOffset func(collectionView NSCollectionView, indexPaths foundation.INSSet, event NSEvent, dragImageOffset foundation.NSPoint) NSImage
	// CanDragItemsAtIndexesWithEvent — Returns a Boolean indicating whether the collection view can begin dragging the specified items.
	CanDragItemsAtIndexesWithEvent func(collectionView NSCollectionView, indexes foundation.NSIndexSet, event NSEvent) bool
	// DraggingImageForItemsAtIndexesWithEventOffset — Creates and returns a drag image to represent the specified items during a drag.
	DraggingImageForItemsAtIndexesWithEventOffset func(collectionView NSCollectionView, indexes foundation.NSIndexSet, event NSEvent, dragImageOffset foundation.NSPoint) NSImage
}

// NewNSCollectionViewDelegate creates an Objective-C object implementing the [NSCollectionViewDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSCollectionViewDelegateObject] satisfies the [NSCollectionViewDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nscollectionviewdelegate
func NewNSCollectionViewDelegate(config NSCollectionViewDelegateConfig) NSCollectionViewDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSCollectionViewDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ShouldSelectItemsAtIndexPaths != nil {
		fn := config.ShouldSelectItemsAtIndexPaths
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:shouldSelectItemsAtIndexPaths:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, indexPathsID objc.ID) objc.ID {
				collectionView := NSCollectionViewFromID(collectionViewID)
				indexPaths := foundation.NSSetFromID(indexPathsID)
				return fn(collectionView, indexPaths).GetID()
			},
		})
	}

	if config.DidSelectItemsAtIndexPaths != nil {
		fn := config.DidSelectItemsAtIndexPaths
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:didSelectItemsAtIndexPaths:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, indexPathsID objc.ID) {
				collectionView := NSCollectionViewFromID(collectionViewID)
				indexPaths := foundation.NSSetFromID(indexPathsID)
				fn(collectionView, indexPaths)
			},
		})
	}

	if config.ShouldDeselectItemsAtIndexPaths != nil {
		fn := config.ShouldDeselectItemsAtIndexPaths
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:shouldDeselectItemsAtIndexPaths:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, indexPathsID objc.ID) objc.ID {
				collectionView := NSCollectionViewFromID(collectionViewID)
				indexPaths := foundation.NSSetFromID(indexPathsID)
				return fn(collectionView, indexPaths).GetID()
			},
		})
	}

	if config.DidDeselectItemsAtIndexPaths != nil {
		fn := config.DidDeselectItemsAtIndexPaths
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:didDeselectItemsAtIndexPaths:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, indexPathsID objc.ID) {
				collectionView := NSCollectionViewFromID(collectionViewID)
				indexPaths := foundation.NSSetFromID(indexPathsID)
				fn(collectionView, indexPaths)
			},
		})
	}

	if config.ShouldChangeItemsAtIndexPathsToHighlightState != nil {
		fn := config.ShouldChangeItemsAtIndexPathsToHighlightState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, indexPathsID objc.ID, highlightState NSCollectionViewItemHighlightState) objc.ID {
				collectionView := NSCollectionViewFromID(collectionViewID)
				indexPaths := foundation.NSSetFromID(indexPathsID)
				return fn(collectionView, indexPaths, highlightState).GetID()
			},
		})
	}

	if config.DidChangeItemsAtIndexPathsToHighlightState != nil {
		fn := config.DidChangeItemsAtIndexPathsToHighlightState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:didChangeItemsAtIndexPaths:toHighlightState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, indexPathsID objc.ID, highlightState NSCollectionViewItemHighlightState) {
				collectionView := NSCollectionViewFromID(collectionViewID)
				indexPaths := foundation.NSSetFromID(indexPathsID)
				fn(collectionView, indexPaths, highlightState)
			},
		})
	}

	if config.WillDisplayItemForRepresentedObjectAtIndexPath != nil {
		fn := config.WillDisplayItemForRepresentedObjectAtIndexPath
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, itemID objc.ID, indexPathID objc.ID) {
				collectionView := NSCollectionViewFromID(collectionViewID)
				item := NSCollectionViewItemFromID(itemID)
				indexPath := foundation.NSIndexPathFromID(indexPathID)
				fn(collectionView, item, indexPath)
			},
		})
	}

	if config.DidEndDisplayingItemForRepresentedObjectAtIndexPath != nil {
		fn := config.DidEndDisplayingItemForRepresentedObjectAtIndexPath
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, itemID objc.ID, indexPathID objc.ID) {
				collectionView := NSCollectionViewFromID(collectionViewID)
				item := NSCollectionViewItemFromID(itemID)
				indexPath := foundation.NSIndexPathFromID(indexPathID)
				fn(collectionView, item, indexPath)
			},
		})
	}

	if config.WillDisplaySupplementaryViewForElementKindAtIndexPath != nil {
		fn := config.WillDisplaySupplementaryViewForElementKindAtIndexPath
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, viewID objc.ID, elementKindID objc.ID, indexPathID objc.ID) {
				collectionView := NSCollectionViewFromID(collectionViewID)
				view := NSViewFromID(viewID)
				elementKind := NSCollectionViewSupplementaryElementKind(objc.GoString(objc.Send[*byte](elementKindID, objc.Sel("UTF8String"))))
				indexPath := foundation.NSIndexPathFromID(indexPathID)
				fn(collectionView, view, elementKind, indexPath)
			},
		})
	}

	if config.DidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath != nil {
		fn := config.DidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, viewID objc.ID, elementKindID objc.ID, indexPathID objc.ID) {
				collectionView := NSCollectionViewFromID(collectionViewID)
				view := NSViewFromID(viewID)
				elementKind := NSCollectionViewSupplementaryElementKind(objc.GoString(objc.Send[*byte](elementKindID, objc.Sel("UTF8String"))))
				indexPath := foundation.NSIndexPathFromID(indexPathID)
				fn(collectionView, view, elementKind, indexPath)
			},
		})
	}

	if config.TransitionLayoutForOldLayoutNewLayout != nil {
		fn := config.TransitionLayoutForOldLayoutNewLayout
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:transitionLayoutForOldLayout:newLayout:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, fromLayoutID objc.ID, toLayoutID objc.ID) objc.ID {
				collectionView := NSCollectionViewFromID(collectionViewID)
				fromLayout := NSCollectionViewLayoutFromID(fromLayoutID)
				toLayout := NSCollectionViewLayoutFromID(toLayoutID)
				return fn(collectionView, fromLayout, toLayout).GetID()
			},
		})
	}

	if config.CanDragItemsAtIndexPathsWithEvent != nil {
		fn := config.CanDragItemsAtIndexPathsWithEvent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:canDragItemsAtIndexPaths:withEvent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, indexPathsID objc.ID, eventID objc.ID) bool {
				collectionView := NSCollectionViewFromID(collectionViewID)
				indexPaths := foundation.NSSetFromID(indexPathsID)
				event := NSEventFromID(eventID)
				return fn(collectionView, indexPaths, event)
			},
		})
	}

	if config.DraggingImageForItemsAtIndexPathsWithEventOffset != nil {
		fn := config.DraggingImageForItemsAtIndexPathsWithEventOffset
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, indexPathsID objc.ID, eventID objc.ID, dragImageOffset foundation.NSPoint) objc.ID {
				collectionView := NSCollectionViewFromID(collectionViewID)
				indexPaths := foundation.NSSetFromID(indexPathsID)
				event := NSEventFromID(eventID)
				return fn(collectionView, indexPaths, event, dragImageOffset).GetID()
			},
		})
	}

	if config.CanDragItemsAtIndexesWithEvent != nil {
		fn := config.CanDragItemsAtIndexesWithEvent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:canDragItemsAtIndexes:withEvent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, indexesID objc.ID, eventID objc.ID) bool {
				collectionView := NSCollectionViewFromID(collectionViewID)
				indexes := foundation.NSIndexSetFromID(indexesID)
				event := NSEventFromID(eventID)
				return fn(collectionView, indexes, event)
			},
		})
	}

	if config.DraggingImageForItemsAtIndexesWithEventOffset != nil {
		fn := config.DraggingImageForItemsAtIndexesWithEventOffset
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("collectionView:draggingImageForItemsAtIndexes:withEvent:offset:"),
			Fn: func(self objc.ID, _cmd objc.SEL, collectionViewID objc.ID, indexesID objc.ID, eventID objc.ID, dragImageOffset foundation.NSPoint) objc.ID {
				collectionView := NSCollectionViewFromID(collectionViewID)
				indexes := foundation.NSIndexSetFromID(indexesID)
				event := NSEventFromID(eventID)
				return fn(collectionView, indexes, event, dragImageOffset).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSCollectionViewDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSCollectionViewDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSCollectionViewDelegateObjectFromID(instance)
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that the destination object (or recipient) of a dragged image must implement.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination
type NSDraggingDestination interface {
	objectivec.IObject
}

// NSDraggingDestinationObject wraps an existing Objective-C object that conforms to the NSDraggingDestination protocol.
type NSDraggingDestinationObject struct {
	objectivec.Object
}
func (o NSDraggingDestinationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSDraggingDestinationObjectFromID constructs a [NSDraggingDestinationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSDraggingDestinationObjectFromID(id objc.ID) NSDraggingDestinationObject {
	return NSDraggingDestinationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Invoked when the dragged image enters destination bounds or frame; delegate
// returns dragging operation to perform.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Return Value
// 
// One (and only one) of the dragging operation constants described in
// [NSDragOperation] in the [NSDraggingInfo] reference. The default return
// value (if this method is not implemented by the destination) is the value
// returned by the previous [DraggingEntered] message.
//
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// # Discussion
// 
// Invoked when a dragged image enters the destination but only if the
// destination has registered for the pasteboard data type involved in the
// drag operation. Specifically, this method is invoked when the mouse pointer
// enters the destination’s bounds rectangle (if it is a view object) or its
// frame rectangle (if it is a window object).
// 
// This method must return a value that indicates which dragging operation the
// destination will perform when the image is released. In deciding which
// dragging operation to return, the method should evaluate the overlap
// between both the dragging operations allowed by the source (obtained from
// `sender` with the [DraggingSourceOperationMask] method) and the dragging
// operations and pasteboard data types the destination itself supports.
// 
// If none of the operations is appropriate, this method should return
// [NSDragOperationNone] (this is the default response if the method is not
// implemented by the destination). A destination will still receive
// [DraggingUpdated] and [DraggingExited] even if [NSDragOperationNone] is
// returned by this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/draggingEntered(_:)
func (o NSDraggingDestinationObject) DraggingEntered(sender NSDraggingInfo) NSDragOperation {
	rv := objc.Send[NSDragOperation](o.ID, objc.Sel("draggingEntered:"), sender)
	return rv
	}
// Asks the destination object whether it wants to receive periodic
// [DraggingUpdated] messages.
//
// # Return Value
// 
// [true] if the destination wants to receive periodic [DraggingUpdated]
// messages, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the destination returns [false], these messages are sent only when the
// mouse moves or a modifier flag changes. Otherwise the destination gets the
// default behavior, where it receives periodic dragging-updated messages even
// if nothing changes.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/wantsPeriodicDraggingUpdates()
func (o NSDraggingDestinationObject) WantsPeriodicDraggingUpdates() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("wantsPeriodicDraggingUpdates"))
	return rv
	}
// Invoked periodically as the image is held within the destination area,
// allowing modification of the dragging operation or mouse-pointer position.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Return Value
// 
// One (and only one) of the dragging operation constants described in
// [NSDragOperation] in the [NSDraggingInfo] reference. The default return
// value (if this method is not implemented by the destination) is the value
// returned by the previous [DraggingEntered] message.
//
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// # Discussion
// 
// For this to be invoked, the destination must have registered for the
// pasteboard data type involved in the drag operation. The messages continue
// until the image is either released or dragged out of the window or view.
// 
// This method provides the destination with an opportunity to modify the
// dragging operation depending on the position of the mouse pointer inside of
// the destination view or window object. For example, you may have several
// graphics or areas of text contained within the same view and wish to tailor
// the dragging operation, or to ignore the drag event completely, depending
// upon which object is underneath the mouse pointer at the time when the user
// releases the dragged image and the [PerformDragOperation] method is
// invoked.
// 
// You typically examine the contents of the pasteboard in the
// [DraggingEntered] method, where this examination is performed only once,
// rather than in the [DraggingUpdated] method, which is invoked multiple
// times.
// 
// Only one destination at a time receives a sequence of [DraggingUpdated]
// messages. If the mouse pointer is within the bounds of two overlapping
// views that are both valid destinations, the uppermost view receives these
// messages until the image is either released or dragged out.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/draggingUpdated(_:)
func (o NSDraggingDestinationObject) DraggingUpdated(sender NSDraggingInfo) NSDragOperation {
	rv := objc.Send[NSDragOperation](o.ID, objc.Sel("draggingUpdated:"), sender)
	return rv
	}
// Invoked when the dragged image exits the destination’s bounds rectangle
// (in the case of a view object) or its frame rectangle (in the case of a
// window object).
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/draggingExited(_:)
func (o NSDraggingDestinationObject) DraggingExited(sender NSDraggingInfo) {
	objc.Send[struct{}](o.ID, objc.Sel("draggingExited:"), sender)
	}
// Called when a drag operation ends.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Discussion
// 
// Implement this method if you want to be notified when a drag operation
// ends, which can be useful if the drag ends in some other destination. For
// example, this method might be used by a destination doing auto-expansion in
// order to collapse any auto-expands.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/draggingEnded(_:)
func (o NSDraggingDestinationObject) DraggingEnded(sender NSDraggingInfo) {
	objc.Send[struct{}](o.ID, objc.Sel("draggingEnded:"), sender)
	}
// Invoked when the image is released, allowing the receiver to agree to or
// refuse drag operation.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Return Value
// 
// [true] if the receiver agrees to perform the drag operation and [false] if
// not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is invoked only if the most recent [DraggingEntered] or
// [DraggingUpdated] message returned an acceptable drag-operation value.
// 
// If you want the drag items to animate from their current location on screen
// to their final location in your view, set the sender object’s
// [AnimatesToDestination] property to [true] in your implementation of this
// method.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/prepareForDragOperation(_:)
func (o NSDraggingDestinationObject) PrepareForDragOperation(sender NSDraggingInfo) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("prepareForDragOperation:"), sender)
	return rv
	}
// Invoked after the released image has been removed from the screen,
// signaling the receiver to import the pasteboard data.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Return Value
// 
// If the destination accepts the data, it returns [true]; otherwise it
// returns [false]. The default is to return [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// For this method to be invoked, the previous [PrepareForDragOperation]
// message must have returned [true]. The destination should implement this
// method to do the real work of importing the pasteboard data represented by
// the image.
// 
// If the sender object’s [AnimatesToDestination] was set to [true] in
// [PrepareForDragOperation], then setup any animation to arrange space for
// the drag items to animate to. Also at this time, enumerate through the
// dragging items to set their destination frames and destination images.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/performDragOperation(_:)
func (o NSDraggingDestinationObject) PerformDragOperation(sender NSDraggingInfo) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("performDragOperation:"), sender)
	return rv
	}
// Invoked when the dragging operation is complete, signaling the receiver to
// perform any necessary clean-up.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Discussion
// 
// For this method to be invoked, the previous [PerformDragOperation] must
// have returned [true].
// 
// The destination implements this method to perform any tidying up that it
// needs to do, such as updating its visual representation now that it has
// incorporated the dragged data. This message is the last message sent from
// `sender` to the destination during a dragging session.
// 
// If the `sender` object’s [AnimatesToDestination] property was set to
// [true] in [PrepareForDragOperation], then the drag image is still visible.
// At this point you should draw the final visual representation in the view.
// When this method returns, the drag image is removed form the screen. If
// your final visual representation matches the visual representation in the
// drag, this is a seamless transition.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/concludeDragOperation(_:)
func (o NSDraggingDestinationObject) ConcludeDragOperation(sender NSDraggingInfo) {
	objc.Send[struct{}](o.ID, objc.Sel("concludeDragOperation:"), sender)
	}
// Invoked when the dragging images should be changed.
//
// sender: The object sending the message; use this object to get details about the
// dragging operation.
//
// # Discussion
// 
// While a destination may change the dragging images at any time, it is
// recommended to wait until this method is called before updating the
// dragging images.
// 
// This allows the system to delay changing the dragging images until it is
// likely that the user will drop on this destination. Otherwise, the dragging
// images will change too often during the drag which would be distracting to
// the user.
// 
// During `` you may set non-acceptable drag items images to `nil` to hide
// them or use the enumeration option of
// [DraggingItemEnumerationClearNonenumeratedImages] If there are items that
// you hide, then after enumeration, you need to set the
// [NumberOfValidItemsForDrop] to the number of non-hidden drag items.
// However, if the valid item count is `0`, then it is better to return
// [NSDragOperationNone] from your implementation of [DraggingEntered] and, or
// [DraggingUpdated] instead of hiding all drag items during enumeration.
//
// [NSDragOperationNone]: https://developer.apple.com/documentation/AppKit/NSDragOperation/NSDragOperationNone
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/updateDraggingItemsForDrag(_:)
func (o NSDraggingDestinationObject) UpdateDraggingItemsForDrag(sender NSDraggingInfo) {
	objc.Send[struct{}](o.ID, objc.Sel("updateDraggingItemsForDrag:"), sender)
	}


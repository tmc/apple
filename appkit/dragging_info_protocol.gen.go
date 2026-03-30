// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that supply information about a dragging session.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo
type NSDraggingInfo interface {
	objectivec.IObject

	// The pasteboard object that holds the dragged data.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingPasteboard
	DraggingPasteboard() INSPasteboard

	// A number that uniquely identifies the dragging session.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingSequenceNumber
	DraggingSequenceNumber() int

	// Information about the dragging operation and the data it contains.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingSourceOperationMask
	DraggingSourceOperationMask() NSDragOperation

	// The current location of the mouse pointer in the base coordinate system of the destination object’s window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingLocation
	DraggingLocation() corefoundation.CGPoint

	// The destination window for the dragging operation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingDestinationWindow
	DraggingDestinationWindow() INSWindow

	// The number of valid items for a drop operation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/numberOfValidItemsForDrop
	NumberOfValidItemsForDrop() int

	// The current location of the dragged image’s origin, in the base coordinate system of the destination object’s window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggedImageLocation
	DraggedImageLocation() corefoundation.CGPoint

	// The image that represents the dragging item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggedImage
	DraggedImage() INSImage

	// Slides the image to a specified location.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/slideDraggedImage(to:)
	SlideDraggedImageTo(screenPoint corefoundation.CGPoint)

	// A Boolean value that indicates whether the dragging formation animates while the drag is over the destination.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/animatesToDestination
	AnimatesToDestination() bool

	// The formation of the dragging items while the drag is over the destination.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingFormation
	DraggingFormation() NSDraggingFormation

	// Enumerates through each dragging item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/enumerateDraggingItems(options:for:classes:searchOptions:using:)
	EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts NSDraggingItemEnumerationOptions, view INSView, classArray []objc.Class, searchOptions foundation.INSDictionary, block DraggingItemHandler)

	// A highlighting style for your app’s user interface to display during a spring-loading operation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/springLoadingHighlight
	SpringLoadingHighlight() NSSpringLoadingHighlight

	// Resets a spring-loading operation to its initial state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/resetSpringLoading()
	ResetSpringLoading()

	// The number of valid items for a drop operation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/numberOfValidItemsForDrop
	SetNumberOfValidItemsForDrop(value int)

	// A Boolean value that indicates whether the dragging formation animates while the drag is over the destination.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/animatesToDestination
	SetAnimatesToDestination(value bool)

	// The formation of the dragging items while the drag is over the destination.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingFormation
	SetDraggingFormation(value NSDraggingFormation)
}

// NSDraggingInfoObject wraps an existing Objective-C object that conforms to the NSDraggingInfo protocol.
type NSDraggingInfoObject struct {
	objectivec.Object
}

func (o NSDraggingInfoObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSDraggingInfoObjectFromID constructs a [NSDraggingInfoObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSDraggingInfoObjectFromID(id objc.ID) NSDraggingInfoObject {
	return NSDraggingInfoObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The pasteboard object that holds the dragged data.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingPasteboard
func (o NSDraggingInfoObject) DraggingPasteboard() INSPasteboard {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("draggingPasteboard"))
	return NSPasteboardFromID(rv)
}

// A number that uniquely identifies the dragging session.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingSequenceNumber
func (o NSDraggingInfoObject) DraggingSequenceNumber() int {
	rv := objc.Send[int](o.ID, objc.Sel("draggingSequenceNumber"))
	return rv
}

// The source, or owner, of the dragged data.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingSource
func (o NSDraggingInfoObject) DraggingSource() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("draggingSource"))
	return objectivec.Object{ID: rv}
}

// Information about the dragging operation and the data it contains.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingSourceOperationMask
func (o NSDraggingInfoObject) DraggingSourceOperationMask() NSDragOperation {
	rv := objc.Send[NSDragOperation](o.ID, objc.Sel("draggingSourceOperationMask"))
	return rv
}

// The current location of the mouse pointer in the base coordinate system of
// the destination object’s window.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingLocation
func (o NSDraggingInfoObject) DraggingLocation() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("draggingLocation"))
	return rv
}

// The destination window for the dragging operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingDestinationWindow
func (o NSDraggingInfoObject) DraggingDestinationWindow() INSWindow {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("draggingDestinationWindow"))
	return NSWindowFromID(rv)
}

// The number of valid items for a drop operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/numberOfValidItemsForDrop
func (o NSDraggingInfoObject) NumberOfValidItemsForDrop() int {
	rv := objc.Send[int](o.ID, objc.Sel("numberOfValidItemsForDrop"))
	return rv
}

// The current location of the dragged image’s origin, in the base
// coordinate system of the destination object’s window.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggedImageLocation
func (o NSDraggingInfoObject) DraggedImageLocation() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("draggedImageLocation"))
	return rv
}

// The image that represents the dragging item.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggedImage
func (o NSDraggingInfoObject) DraggedImage() INSImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("draggedImage"))
	return NSImageFromID(rv)
}

// Slides the image to a specified location.
//
// screenPoint: A point that specifies a location in the screen coordinate system.
//
// # Discussion
//
// This method can be used to adjust the location to which the dragged image
// will slide back if the drag is rejected.
//
// It should only be invoked from within the destination’s implementation of
// prepareForDragOperation:, and will only have effect if the destination
// rejects the drag.
//
// This method is invoked after the user has released the image but before it
// is removed from the screen.
//
// # Special Considerations
//
// This method has been available since OS X v 10.0, however it was not
// implemented until OS X v 10.5. Previous to that version, it did nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/slideDraggedImage(to:)
func (o NSDraggingInfoObject) SlideDraggedImageTo(screenPoint corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("slideDraggedImageTo:"), screenPoint)
}

// A Boolean value that indicates whether the dragging formation animates
// while the drag is over the destination.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/animatesToDestination
func (o NSDraggingInfoObject) AnimatesToDestination() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("animatesToDestination"))
	return rv
}

// The formation of the dragging items while the drag is over the destination.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingFormation
func (o NSDraggingInfoObject) DraggingFormation() NSDraggingFormation {
	rv := objc.Send[NSDraggingFormation](o.ID, objc.Sel("draggingFormation"))
	return rv
}

// Enumerates through each dragging item.
//
// enumOpts: The enumeration options. See [NSDraggingItemEnumerationOptions] for the
// supported values.
//
// view: The view to use as the base coordinate system for the [NSDraggingItem]
// instances.
//
// classArray: An array of class objects.
//
// Arrange classes in the array in the preferred order of representation.
// Classes in the array must conform to the [NSPasteboardReading] protocol.
//
// searchOptions: A dictionary that specifies options to refine the search for pasteboard
// items, such as restricting the search to file URLs with particular content
// types. For valid dictionary keys, see [NSPasteboardReadingOptionKey].
//
// block: The block to execute for the enumeration. The block takes three arguments:
//
// `draggingItem`: A reference to the dragging item. The [DraggingFrame] of
// the dragging item is in the coordinate space of the view that `view`
// specifies. A `view` value of `nil` means the screen coordinate space.
// `idx`: The index of the element in the classes. `stop`: A reference to a
// Boolean value that the block can use to stop the enumeration by setting
// `*stop` to true.
//
// # Discussion
//
// Enumerate through dragging items to modify their properties, such as the
// drag image or size, to indicate that the user has dragged the items over a
// possible destination. Changes you make in this method on behalf of the
// dragging destination override changes from the source’s drag session.
//
// To get dragging items in a data type that you expect while enumerating,
// specify classes in the classesArray parameter that implement the
// [NSPasteboardReading] protocol, such as [NSImage], [NSString], [NSURL],
// [NSColor], [NSAttributedString], or [NSPasteboardItem]. For each item in
// the dragging pasteboard, the system performs the following steps:
//
// - The systems calls [ReadableTypesForPasteboard] on the item to determine
// the types of data the item conforms to. - It attempts to create an instance
// of a matching class from the dragging pasteboard data, using the class
// order you specify in the `classesArray` parameter. - If it can create an
// instance of a matching class, the system creates an instance of
// [NSDraggingItem] with the class instance and the dragging properties of
// that item. - The system passes the [NSDraggingItem] to the block you
// provide as the `draggingItem` parameter.
//
// If the system can’t create an instance of one of the classes you specify
// in `classesArray` with an item, the system skips the item without calling
// `block` and proceeds to the next item.
//
// When the system provides a `draggingItem` to your block, modify the
// item’s properties to change how the user sees the item while dragging.
// Provide a `view` to this method if you want to express each dragging
// item’s `draggingFrame` relative to that `view`.
//
// To refine the list of dragging items that this method provides, specify
// `enumOpts` and `searchOptions`.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/enumerateDraggingItems(options:for:classes:searchOptions:using:)
//
// [NSDraggingItemEnumerationOptions]: https://developer.apple.com/documentation/AppKit/NSDraggingItemEnumerationOptions
func (o NSDraggingInfoObject) EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts NSDraggingItemEnumerationOptions, view INSView, classArray []objc.Class, searchOptions foundation.INSDictionary, block DraggingItemHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"), enumOpts, view, objectivec.ClassSliceToNSArray(classArray), searchOptions, block)
}

// A highlighting style for your app’s user interface to display during a
// spring-loading operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/springLoadingHighlight
func (o NSDraggingInfoObject) SpringLoadingHighlight() NSSpringLoadingHighlight {
	rv := objc.Send[NSSpringLoadingHighlight](o.ID, objc.Sel("springLoadingHighlight"))
	return rv
}

// Resets a spring-loading operation to its initial state.
//
// # Discussion
//
// Call this method when a drag operation moves from one spring-loaded region
// into another within the same destination and the spring-loading options
// haven’t changed.
//
// For example, consider an [NSSegmentedControl] object. It’s a single
// spring-loading destination. However, each segment within the control can be
// individually spring-loaded. Each time the drag enters a new segment within
// the control, the spring-loading destination and options don’t change.
// [NSSegmentedControl] calls the [ResetSpringLoading] method to reset the
// spring-loading operation for the newly entered segment.
//
// When this method is called, the spring-loading operation is reset. Hover
// time required to activate spring-loading is reset and begins again. If a
// force click is currently active, the user must reduce pressure enough to
// disengage the force click. Then, the user must force click again to
// activate the newly entered spring-loaded region.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/resetSpringLoading()
func (o NSDraggingInfoObject) ResetSpringLoading() {
	objc.Send[struct{}](o.ID, objc.Sel("resetSpringLoading"))
}

// The number of valid items for a drop operation.
//
// # Discussion
//
// During draggingEntered: or draggingUpdated:, you are responsible for
// returning the drag operation. In some cases, you may accept some, but not
// all items on the dragging pasteboard. (For example, your application may
// only accept image files.)
//
// If you only accept some of the items, set this property to the number of
// items accepted so the drag manager can update the drag count badge.
//
// When [UpdateDraggingItemsForDrag] is called, you should set the image of
// non-valid dragging items to `nil`. If none of the drag items are valid then
// you should not “, simply return [NSDragOperationNone] from your
// implementation of draggingEntered: and, or draggingUpdated: and do not
// modify any drag item properties.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/numberOfValidItemsForDrop
func (o NSDraggingInfoObject) SetNumberOfValidItemsForDrop(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setNumberOfValidItemsForDrop:"), value)
}

// A Boolean value that indicates whether the dragging formation animates
// while the drag is over the destination.
//
// # Discussion
//
// During the conclusion of an accepted drag, if this property is set to true,
// the drag manager will animate each dragging image to their
// [NSDraggingFormationNone] locations. Otherwise, the drag images are removed
// without any animation.
//
// This property is inspected between prepareForDragOperation: and
// performDragOperation:. You should enumerate through the dragging items
// during performDragOperation: to set the item’s [DraggingFrame] to the
// correct destinations.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/animatesToDestination
func (o NSDraggingInfoObject) SetAnimatesToDestination(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAnimatesToDestination:"), value)
}

// The formation of the dragging items while the drag is over the destination.
//
// # Discussion
//
// Set this property to change the formation of the drag items. This is
// generally done during the [UpdateDraggingItemsForDrag] method or whenever
// you enumerate the dragging items.
//
// The default value is the current drag formation.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingInfo/draggingFormation
func (o NSDraggingInfoObject) SetDraggingFormation(value NSDraggingFormation) {
	objc.Send[struct{}](o.ID, objc.Sel("setDraggingFormation:"), value)
}

// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDraggingSession] class.
var (
	_NSDraggingSessionClass     NSDraggingSessionClass
	_NSDraggingSessionClassOnce sync.Once
)

func getNSDraggingSessionClass() NSDraggingSessionClass {
	_NSDraggingSessionClassOnce.Do(func() {
		_NSDraggingSessionClass = NSDraggingSessionClass{class: objc.GetClass("NSDraggingSession")}
	})
	return _NSDraggingSessionClass
}

// GetNSDraggingSessionClass returns the class object for NSDraggingSession.
func GetNSDraggingSessionClass() NSDraggingSessionClass {
	return getNSDraggingSessionClass()
}

type NSDraggingSessionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDraggingSessionClass) Alloc() NSDraggingSession {
	rv := objc.Send[NSDraggingSession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// The encapsulation of a drag-and-drop action that supports modification of
// the drag while in progress.
//
// # Overview
// 
// You start a new dragging session by calling the [NSView] method
// [BeginDraggingSessionWithItemsEventSource] method. This method immediately
// returns and you can further modify the properties of the dragging session.
// The actual drag begins at the next turn of the run loop.
//
// # Dragging Pasteboard
//
//   - [NSDraggingSession.DraggingPasteboard]: Returns the pasteboard object that contains the data being dragged.
//
// # Dragging Visual Representation
//
//   - [NSDraggingSession.AnimatesToStartingPositionsOnCancelOrFail]: Controls whether the dragging image animates back to its starting point on a cancelled or failed drag.
//   - [NSDraggingSession.SetAnimatesToStartingPositionsOnCancelOrFail]
//   - [NSDraggingSession.DraggingFormation]: Controls the dragging formation when the drag is not over the source or a valid destination.
//   - [NSDraggingSession.SetDraggingFormation]
//
// # Identifying the Dragging Session
//
//   - [NSDraggingSession.DraggingSequenceNumber]: Returns a number that uniquely identifies the dragging session.
//
// # Enumerating Dragging Items
//
//   - [NSDraggingSession.EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock]: Enumerates through each dragging item.
//
// # Dragging Session Location
//
//   - [NSDraggingSession.DraggingLocation]: The current cursor location of the drag in screen coordinates.
//
// # Dragging Item Location
//
//   - [NSDraggingSession.DraggingLeaderIndex]: The index of the dragging item under the cursor.
//   - [NSDraggingSession.SetDraggingLeaderIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSession
type NSDraggingSession struct {
	objectivec.Object
}

// NSDraggingSessionFromID constructs a [NSDraggingSession] from an objc.ID.
//
// The encapsulation of a drag-and-drop action that supports modification of
// the drag while in progress.
func NSDraggingSessionFromID(id objc.ID) NSDraggingSession {
	return NSDraggingSession{objectivec.Object{id}}
}
// NOTE: NSDraggingSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSDraggingSession] class.
//
// # Dragging Pasteboard
//
//   - [INSDraggingSession.DraggingPasteboard]: Returns the pasteboard object that contains the data being dragged.
//
// # Dragging Visual Representation
//
//   - [INSDraggingSession.AnimatesToStartingPositionsOnCancelOrFail]: Controls whether the dragging image animates back to its starting point on a cancelled or failed drag.
//   - [INSDraggingSession.SetAnimatesToStartingPositionsOnCancelOrFail]
//   - [INSDraggingSession.DraggingFormation]: Controls the dragging formation when the drag is not over the source or a valid destination.
//   - [INSDraggingSession.SetDraggingFormation]
//
// # Identifying the Dragging Session
//
//   - [INSDraggingSession.DraggingSequenceNumber]: Returns a number that uniquely identifies the dragging session.
//
// # Enumerating Dragging Items
//
//   - [INSDraggingSession.EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock]: Enumerates through each dragging item.
//
// # Dragging Session Location
//
//   - [INSDraggingSession.DraggingLocation]: The current cursor location of the drag in screen coordinates.
//
// # Dragging Item Location
//
//   - [INSDraggingSession.DraggingLeaderIndex]: The index of the dragging item under the cursor.
//   - [INSDraggingSession.SetDraggingLeaderIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSession
type INSDraggingSession interface {
	objectivec.IObject

	// Topic: Dragging Pasteboard

	// Returns the pasteboard object that contains the data being dragged.
	DraggingPasteboard() INSPasteboard

	// Topic: Dragging Visual Representation

	// Controls whether the dragging image animates back to its starting point on a cancelled or failed drag.
	AnimatesToStartingPositionsOnCancelOrFail() bool
	SetAnimatesToStartingPositionsOnCancelOrFail(value bool)
	// Controls the dragging formation when the drag is not over the source or a valid destination.
	DraggingFormation() NSDraggingFormation
	SetDraggingFormation(value NSDraggingFormation)

	// Topic: Identifying the Dragging Session

	// Returns a number that uniquely identifies the dragging session.
	DraggingSequenceNumber() int

	// Topic: Enumerating Dragging Items

	// Enumerates through each dragging item.
	EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts NSDraggingItemEnumerationOptions, view INSView, classArray []objc.Class, searchOptions foundation.INSDictionary, block bool)

	// Topic: Dragging Session Location

	// The current cursor location of the drag in screen coordinates.
	DraggingLocation() corefoundation.CGPoint

	// Topic: Dragging Item Location

	// The index of the dragging item under the cursor.
	DraggingLeaderIndex() int
	SetDraggingLeaderIndex(value int)
}





// Init initializes the instance.
func (d NSDraggingSession) Init() NSDraggingSession {
	rv := objc.Send[NSDraggingSession](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDraggingSession) Autorelease() NSDraggingSession {
	rv := objc.Send[NSDraggingSession](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDraggingSession creates a new NSDraggingSession instance.
func NewNSDraggingSession() NSDraggingSession {
	class := getNSDraggingSessionClass()
	rv := objc.Send[NSDraggingSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Enumerates through each dragging item.
//
// enumOpts: The enumeration options. See [NSDraggingItemEnumerationOptions] for the
// supported values.
// //
// [NSDraggingItemEnumerationOptions]: https://developer.apple.com/documentation/AppKit/NSDraggingItemEnumerationOptions
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
// `*stop` to [true].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Enumerate through dragging items to modify their properties, such as the
// drag image or size, while the user is dragging. When the dragging items are
// over the destination, changes that the destination’s [NSDraggingInfo]
// methods make may override changes you make during the drag session.
// 
// To get dragging items in a data type that you expect while enumerating,
// specify classes in the `classesArray` parameter that implement the
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
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSession/enumerateDraggingItems(options:for:classes:searchOptions:using:)
func (d NSDraggingSession) EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts NSDraggingItemEnumerationOptions, view INSView, classArray []objc.Class, searchOptions foundation.INSDictionary, block bool) {
	objc.Send[objc.ID](d.ID, objc.Sel("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"), enumOpts, view, objectivec.ClassSliceToNSArray(classArray), searchOptions, block)
}












// Returns the pasteboard object that contains the data being dragged.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingPasteboard
func (d NSDraggingSession) DraggingPasteboard() INSPasteboard {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("draggingPasteboard"))
	return NSPasteboardFromID(objc.ID(rv))
}



// Controls whether the dragging image animates back to its starting point on
// a cancelled or failed drag.
//
// # Discussion
// 
// This property should be set immediately after creating the dragging
// session.
// 
// The default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSession/animatesToStartingPositionsOnCancelOrFail
func (d NSDraggingSession) AnimatesToStartingPositionsOnCancelOrFail() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("animatesToStartingPositionsOnCancelOrFail"))
	return rv
}
func (d NSDraggingSession) SetAnimatesToStartingPositionsOnCancelOrFail(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setAnimatesToStartingPositionsOnCancelOrFail:"), value)
}



// Controls the dragging formation when the drag is not over the source or a
// valid destination.
//
// # Discussion
// 
// Setting this value causes the dragging formation to change immediately,
// provided a valid destination has not overriden the behavior. If the
// dragging session hasn’t started yet, the dragging items will animate into
// formation immediately upon start. It is recommended to never change the
// formation when starting a drag.
// 
// The default value is [NSDraggingFormation.none].
//
// [NSDraggingFormation.none]: https://developer.apple.com/documentation/AppKit/NSDraggingFormation/none
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingFormation
func (d NSDraggingSession) DraggingFormation() NSDraggingFormation {
	rv := objc.Send[NSDraggingFormation](d.ID, objc.Sel("draggingFormation"))
	return NSDraggingFormation(rv)
}
func (d NSDraggingSession) SetDraggingFormation(value NSDraggingFormation) {
	objc.Send[struct{}](d.ID, objc.Sel("setDraggingFormation:"), value)
}



// Returns a number that uniquely identifies the dragging session.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingSequenceNumber
func (d NSDraggingSession) DraggingSequenceNumber() int {
	rv := objc.Send[int](d.ID, objc.Sel("draggingSequenceNumber"))
	return rv
}



// The current cursor location of the drag in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingLocation
func (d NSDraggingSession) DraggingLocation() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](d.ID, objc.Sel("draggingLocation"))
	return corefoundation.CGPoint(rv)
}



// The index of the dragging item under the cursor.
//
// # Discussion
// 
// The index is to an element in the array passed as the first parameter to
// the [NSView] method [BeginDraggingSessionWithItemsEventSource].
// 
// The default is the [NSDraggingItem] closest to the `location` field in the
// event parameter that was passed to the
// [BeginDraggingSessionWithItemsEventSource] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSession/draggingLeaderIndex
func (d NSDraggingSession) DraggingLeaderIndex() int {
	rv := objc.Send[int](d.ID, objc.Sel("draggingLeaderIndex"))
	return rv
}
func (d NSDraggingSession) SetDraggingLeaderIndex(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setDraggingLeaderIndex:"), value)
}
























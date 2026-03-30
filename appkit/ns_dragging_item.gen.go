// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDraggingItem] class.
var (
	_NSDraggingItemClass     NSDraggingItemClass
	_NSDraggingItemClassOnce sync.Once
)

func getNSDraggingItemClass() NSDraggingItemClass {
	_NSDraggingItemClassOnce.Do(func() {
		_NSDraggingItemClass = NSDraggingItemClass{class: objc.GetClass("NSDraggingItem")}
	})
	return _NSDraggingItemClass
}

// GetNSDraggingItemClass returns the class object for NSDraggingItem.
func GetNSDraggingItemClass() NSDraggingItemClass {
	return getNSDraggingItemClass()
}

type NSDraggingItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSDraggingItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDraggingItemClass) Alloc() NSDraggingItem {
	rv := objc.Send[NSDraggingItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A single dragged item within a dragging session.
//
// # Overview
//
// [NSDraggingItem] objects have extremely limited lifetimes. Don’t retain
// these items because changing outside of the prescribed lifetimes has no
// impact on the drag.
//
// When you call the [NSDraggingSession] method
// [BeginDraggingSessionWithItemsEventSource], the system immediately consumes
// the dragging items that pass to the method, and doesn’t retain them. Any
// further changes to the dragging item associated with the returned
// [NSDraggingSession] must occur with the enumeration method
// [EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock].
// When enumerating, the system creates [NSDraggingItem] instances right
// before giving them to the enumeration block. After returning from the
// block, the dragging item is no longer valid.
//
// # Initializing a dragging item
//
//   - [NSDraggingItem.InitWithPasteboardWriter]: Creates and returns a dragging item using the specified content.
//
// # Dragging frame
//
//   - [NSDraggingItem.SetDraggingFrameContents]: Sets the item’s dragging frame and contents.
//   - [NSDraggingItem.DraggingFrame]: The frame of the dragging item.
//   - [NSDraggingItem.SetDraggingFrame]
//
// # Drag image components
//
//   - [NSDraggingItem.ImageComponents]: An array of dragging image components to use to create the drag image.
//   - [NSDraggingItem.ImageComponentsProvider]: An array of blocks that provide the dragging image components.
//   - [NSDraggingItem.SetImageComponentsProvider]
//   - [NSDraggingItem.Item]: The pasteboard reader or writer object dependent on the context where you use the dragging item.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem
type NSDraggingItem struct {
	objectivec.Object
}

// NSDraggingItemFromID constructs a [NSDraggingItem] from an objc.ID.
//
// A single dragged item within a dragging session.
func NSDraggingItemFromID(id objc.ID) NSDraggingItem {
	return NSDraggingItem{objectivec.Object{ID: id}}
}

// NOTE: NSDraggingItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDraggingItem] class.
//
// # Initializing a dragging item
//
//   - [INSDraggingItem.InitWithPasteboardWriter]: Creates and returns a dragging item using the specified content.
//
// # Dragging frame
//
//   - [INSDraggingItem.SetDraggingFrameContents]: Sets the item’s dragging frame and contents.
//   - [INSDraggingItem.DraggingFrame]: The frame of the dragging item.
//   - [INSDraggingItem.SetDraggingFrame]
//
// # Drag image components
//
//   - [INSDraggingItem.ImageComponents]: An array of dragging image components to use to create the drag image.
//   - [INSDraggingItem.ImageComponentsProvider]: An array of blocks that provide the dragging image components.
//   - [INSDraggingItem.SetImageComponentsProvider]
//   - [INSDraggingItem.Item]: The pasteboard reader or writer object dependent on the context where you use the dragging item.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem
type INSDraggingItem interface {
	objectivec.IObject

	// Topic: Initializing a dragging item

	// Creates and returns a dragging item using the specified content.
	InitWithPasteboardWriter(pasteboardWriter NSPasteboardWriting) NSDraggingItem

	// Topic: Dragging frame

	// Sets the item’s dragging frame and contents.
	SetDraggingFrameContents(frame corefoundation.CGRect, contents objectivec.IObject)
	// The frame of the dragging item.
	DraggingFrame() corefoundation.CGRect
	SetDraggingFrame(value corefoundation.CGRect)

	// Topic: Drag image components

	// An array of dragging image components to use to create the drag image.
	ImageComponents() []NSDraggingImageComponent
	// An array of blocks that provide the dragging image components.
	ImageComponentsProvider() VoidHandler
	SetImageComponentsProvider(value VoidHandler)
	// The pasteboard reader or writer object dependent on the context where you use the dragging item.
	Item() objectivec.IObject
}

// Init initializes the instance.
func (d NSDraggingItem) Init() NSDraggingItem {
	rv := objc.Send[NSDraggingItem](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDraggingItem) Autorelease() NSDraggingItem {
	rv := objc.Send[NSDraggingItem](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDraggingItem creates a new NSDraggingItem instance.
func NewNSDraggingItem() NSDraggingItem {
	class := getNSDraggingItemClass()
	rv := objc.Send[NSDraggingItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns a dragging item using the specified content.
//
// pasteboardWriter: The object that provides the dragging content. The object must implement
// the [NSPasteboardWriting] protocol.
//
// # Return Value
//
// An initialized NSDraggingItem instance with the specified dragging content.
//
// # Discussion
//
// When the developer creates an [NSDraggingItem] instance , it is for use
// with the view method [BeginDraggingSessionWithItemsEventSource] During the
// invocation of that method, the `pasteboardWriter` is placed onto the
// dragging pasteboard for the [NSDraggingSession] that contains the dragging
// item instance.
//
// The designated initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem/init(pasteboardWriter:)
func NewDraggingItemWithPasteboardWriter(pasteboardWriter NSPasteboardWriting) NSDraggingItem {
	instance := getNSDraggingItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPasteboardWriter:"), pasteboardWriter)
	return NSDraggingItemFromID(rv)
}

// Creates and returns a dragging item using the specified content.
//
// pasteboardWriter: The object that provides the dragging content. The object must implement
// the [NSPasteboardWriting] protocol.
//
// # Return Value
//
// An initialized NSDraggingItem instance with the specified dragging content.
//
// # Discussion
//
// When the developer creates an [NSDraggingItem] instance , it is for use
// with the view method [BeginDraggingSessionWithItemsEventSource] During the
// invocation of that method, the `pasteboardWriter` is placed onto the
// dragging pasteboard for the [NSDraggingSession] that contains the dragging
// item instance.
//
// The designated initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem/init(pasteboardWriter:)
func (d NSDraggingItem) InitWithPasteboardWriter(pasteboardWriter NSPasteboardWriting) NSDraggingItem {
	rv := objc.Send[NSDraggingItem](d.ID, objc.Sel("initWithPasteboardWriter:"), pasteboardWriter)
	return rv
}

// Sets the item’s dragging frame and contents.
//
// frame: The item content frame, which is in the same coordinate space as the value
// of [DraggingFrame].
//
// contents: The item contents to display when dragging. Typically this is an [NSImage],
// but a [CGImageRef] will also work.
//
// # Discussion
//
// Alternate single image component setter.
//
// This convenience method simplifies modifying the components of an
// [NSDraggingItem] when there is only one component. It sets the
// [DraggingFrame] and creates a single [NSDraggingImageComponent] instance
// with one image corresponding to the [icon] key. You should use this method
// only under the following conditions: the drag image for this item is
// composed of a single image, or there are a reasonable number of dragging
// item instances being created or enumerated.
//
// If your application requires the dragging of hundreds of items this method
// would create a instance for each item when it is called. Compare this to
// the [ImageComponentsProvider] block which is much faster to define and
// allows AppKit to create only a subset of the items using
// [ImageComponentsProvider].
//
// This method sets the [DraggingFrame] and [ImageComponents] properties.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem/setDraggingFrame(_:contents:)
//
// [icon]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/ImageComponentKey/icon
func (d NSDraggingItem) SetDraggingFrameContents(frame corefoundation.CGRect, contents objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("setDraggingFrame:contents:"), frame, contents)
}

// The frame of the dragging item.
//
// # Discussion
//
// The dragging frame provides the spatial relationship between
// [NSDraggingItem] instances when you set the dragging formation to
// [NSDraggingFormationNone].
//
// The exact coordinate space of this rectangle depends on where you use it.
// Examples are the view that initiates the drag using
// [BeginDraggingSessionWithItemsEventSource] or the view you pass to the
// [NSDraggingSession] implementation of
// [EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock].
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem/draggingFrame
func (d NSDraggingItem) DraggingFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](d.ID, objc.Sel("draggingFrame"))
	return corefoundation.CGRect(rv)
}
func (d NSDraggingItem) SetDraggingFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](d.ID, objc.Sel("setDraggingFrame:"), value)
}

// An array of dragging image components to use to create the drag image.
//
// # Discussion
//
// The array contains copies of the components. The drag does not reflect
// changes you make to these copies. If needed, the system calls the
// [ImageComponentsProvider] block to generate the image components.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem/imageComponents
func (d NSDraggingItem) ImageComponents() []NSDraggingImageComponent {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("imageComponents"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSDraggingImageComponent {
		return NSDraggingImageComponentFromID(id)
	})
}

// An array of blocks that provide the dragging image components.
//
// # Discussion
//
// The dragging image is the composite of an array of
// [NSDraggingImageComponent] objects.
//
// The dragging image components aren’t set directly. Instead, use a block
// to generate the components and the system calls the block if necessary.
//
// You can set the block to `nil`, meaning that the drag item has no image.
// Generally, only dragging destinations do this, and only if there’s at
// least one valid item in the drop, and the receiver isn’t that object.
//
// The system arranges the components in painting order. That is, the system
// paints each component in the array on top of the previous components in the
// array.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem/imageComponentsProvider
func (d NSDraggingItem) ImageComponentsProvider() VoidHandler {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("imageComponentsProvider"))
	_ = rv
	return nil
}
func (d NSDraggingItem) SetImageComponentsProvider(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](d.ID, objc.Sel("setImageComponentsProvider:"), block)
}

// The pasteboard reader or writer object dependent on the context where you
// use the dragging item.
//
// # Discussion
//
// When you create an [NSDraggingItem] instance, `item` is the
// `pasteboardWriter` passed to [InitWithPasteboardWriter].
//
// However, when enumerating dragging items using the [NSDraggingSession]
// method
// [EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock] or
// the [NSDraggingInfo] method
// [EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock],
// `item` is not the original pasteboard reader or writer instance. It is an
// instance of one of the classes provided to the enumeration method’s
// `classArray` parameter.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem/item
func (d NSDraggingItem) Item() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("item"))
	return objectivec.Object{ID: rv}
}

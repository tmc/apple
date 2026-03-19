// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutBoundarySupplementaryItem] class.
var (
	_NSCollectionLayoutBoundarySupplementaryItemClass     NSCollectionLayoutBoundarySupplementaryItemClass
	_NSCollectionLayoutBoundarySupplementaryItemClassOnce sync.Once
)

func getNSCollectionLayoutBoundarySupplementaryItemClass() NSCollectionLayoutBoundarySupplementaryItemClass {
	_NSCollectionLayoutBoundarySupplementaryItemClassOnce.Do(func() {
		_NSCollectionLayoutBoundarySupplementaryItemClass = NSCollectionLayoutBoundarySupplementaryItemClass{class: objc.GetClass("NSCollectionLayoutBoundarySupplementaryItem")}
	})
	return _NSCollectionLayoutBoundarySupplementaryItemClass
}

// GetNSCollectionLayoutBoundarySupplementaryItemClass returns the class object for NSCollectionLayoutBoundarySupplementaryItem.
func GetNSCollectionLayoutBoundarySupplementaryItemClass() NSCollectionLayoutBoundarySupplementaryItemClass {
	return getNSCollectionLayoutBoundarySupplementaryItemClass()
}

type NSCollectionLayoutBoundarySupplementaryItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutBoundarySupplementaryItemClass) Alloc() NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[NSCollectionLayoutBoundarySupplementaryItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object used to add headers or footers to a collection view.
//
// # Overview
// 
// A boundary supplementary item is a specialized type of supplementary item
// ([NSCollectionLayoutSupplementaryItem]). You use boundary supplementary
// items to add headers or footers to a section of a collection view or the
// entire collection view.
// 
// Each type of supplementary item must have a unique element kind. Consider
// tracking these strings together in a way that makes it straightforward to
// identify each element, for example:
// 
// Add boundary supplementary items to a section by setting that section’s
// [NSCollectionLayoutBoundarySupplementaryItem.BoundarySupplementaryItems] property:
//
// # Specifying scrolling behavior
//
//   - [NSCollectionLayoutBoundarySupplementaryItem.PinToVisibleBounds]: A Boolean value that indicates whether a header or footer is pinned to the top or bottom visible boundary of the section or layout it’s attached to.
//   - [NSCollectionLayoutBoundarySupplementaryItem.SetPinToVisibleBounds]
//
// # Specifying position
//
//   - [NSCollectionLayoutBoundarySupplementaryItem.Offset]: The floating-point value of the boundary supplementary item’s offset from the section or layout it’s attached to.
//   - [NSCollectionLayoutBoundarySupplementaryItem.Alignment]: The alignment of the boundary supplementary item relative to the section or layout it’s attached to.
//   - [NSCollectionLayoutBoundarySupplementaryItem.ExtendsBoundary]: A Boolean value that indicates whether a boundary supplementary item extends the content area of the section or layout it’s attached to.
//   - [NSCollectionLayoutBoundarySupplementaryItem.SetExtendsBoundary]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem
type NSCollectionLayoutBoundarySupplementaryItem struct {
	NSCollectionLayoutSupplementaryItem
}

// NSCollectionLayoutBoundarySupplementaryItemFromID constructs a [NSCollectionLayoutBoundarySupplementaryItem] from an objc.ID.
//
// An object used to add headers or footers to a collection view.
func NSCollectionLayoutBoundarySupplementaryItemFromID(id objc.ID) NSCollectionLayoutBoundarySupplementaryItem {
	return NSCollectionLayoutBoundarySupplementaryItem{NSCollectionLayoutSupplementaryItem: NSCollectionLayoutSupplementaryItemFromID(id)}
}
// NOTE: NSCollectionLayoutBoundarySupplementaryItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionLayoutBoundarySupplementaryItem] class.
//
// # Specifying scrolling behavior
//
//   - [INSCollectionLayoutBoundarySupplementaryItem.PinToVisibleBounds]: A Boolean value that indicates whether a header or footer is pinned to the top or bottom visible boundary of the section or layout it’s attached to.
//   - [INSCollectionLayoutBoundarySupplementaryItem.SetPinToVisibleBounds]
//
// # Specifying position
//
//   - [INSCollectionLayoutBoundarySupplementaryItem.Offset]: The floating-point value of the boundary supplementary item’s offset from the section or layout it’s attached to.
//   - [INSCollectionLayoutBoundarySupplementaryItem.Alignment]: The alignment of the boundary supplementary item relative to the section or layout it’s attached to.
//   - [INSCollectionLayoutBoundarySupplementaryItem.ExtendsBoundary]: A Boolean value that indicates whether a boundary supplementary item extends the content area of the section or layout it’s attached to.
//   - [INSCollectionLayoutBoundarySupplementaryItem.SetExtendsBoundary]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem
type INSCollectionLayoutBoundarySupplementaryItem interface {
	INSCollectionLayoutSupplementaryItem

	// Topic: Specifying scrolling behavior

	// A Boolean value that indicates whether a header or footer is pinned to the top or bottom visible boundary of the section or layout it’s attached to.
	PinToVisibleBounds() bool
	SetPinToVisibleBounds(value bool)

	// Topic: Specifying position

	// The floating-point value of the boundary supplementary item’s offset from the section or layout it’s attached to.
	Offset() corefoundation.CGPoint
	// The alignment of the boundary supplementary item relative to the section or layout it’s attached to.
	Alignment() NSRectAlignment
	// A Boolean value that indicates whether a boundary supplementary item extends the content area of the section or layout it’s attached to.
	ExtendsBoundary() bool
	SetExtendsBoundary(value bool)

	// An array of the supplementary items that are associated with the boundary edges of the entire layout, such as global headers and footers.
	BoundarySupplementaryItems() INSCollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value INSCollectionLayoutBoundarySupplementaryItem)
}

// Init initializes the instance.
func (c NSCollectionLayoutBoundarySupplementaryItem) Init() NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[NSCollectionLayoutBoundarySupplementaryItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutBoundarySupplementaryItem) Autorelease() NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[NSCollectionLayoutBoundarySupplementaryItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutBoundarySupplementaryItem creates a new NSCollectionLayoutBoundarySupplementaryItem instance.
func NewNSCollectionLayoutBoundarySupplementaryItem() NSCollectionLayoutBoundarySupplementaryItem {
	class := getNSCollectionLayoutBoundarySupplementaryItemClass()
	rv := objc.Send[NSCollectionLayoutBoundarySupplementaryItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an item of the specified size.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:)
func NewCollectionLayoutBoundarySupplementaryItemWithLayoutSize(layoutSize INSCollectionLayoutSize) NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutBoundarySupplementaryItemClass().class), objc.Sel("itemWithLayoutSize:"), layoutSize)
	return NSCollectionLayoutBoundarySupplementaryItemFromID(rv)
}

// Creates a boundary supplementary item of the specified size and element
// kind, with an alignment relative to a section or layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/init(layoutSize:elementKind:alignment:)
func NewCollectionLayoutBoundarySupplementaryItemWithLayoutSizeElementKindAlignment(layoutSize INSCollectionLayoutSize, elementKind string, alignment NSRectAlignment) NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutBoundarySupplementaryItemClass().class), objc.Sel("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:"), layoutSize, objc.String(elementKind), alignment)
	return NSCollectionLayoutBoundarySupplementaryItemFromID(rv)
}

// Creates a boundary supplementary item of the specified size and element
// kind, with an alignment relative to a section or layout at an absolute
// offset.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/init(layoutSize:elementKind:alignment:absoluteOffset:)
func NewCollectionLayoutBoundarySupplementaryItemWithLayoutSizeElementKindAlignmentAbsoluteOffset(layoutSize INSCollectionLayoutSize, elementKind string, alignment NSRectAlignment, absoluteOffset corefoundation.CGPoint) NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutBoundarySupplementaryItemClass().class), objc.Sel("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:absoluteOffset:"), layoutSize, objc.String(elementKind), alignment, absoluteOffset)
	return NSCollectionLayoutBoundarySupplementaryItemFromID(rv)
}

// Creates a supplementary item of the specified size and element kind, with
// an anchor relative to a container.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/init(layoutSize:elementKind:containerAnchor:)
func NewCollectionLayoutBoundarySupplementaryItemWithLayoutSizeElementKindContainerAnchor(layoutSize INSCollectionLayoutSize, elementKind string, containerAnchor INSCollectionLayoutAnchor) NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutBoundarySupplementaryItemClass().class), objc.Sel("supplementaryItemWithLayoutSize:elementKind:containerAnchor:"), layoutSize, objc.String(elementKind), containerAnchor)
	return NSCollectionLayoutBoundarySupplementaryItemFromID(rv)
}

// Creates a supplementary item of the specified size and element kind, an
// anchor relative to a container, and an anchor relative to an item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/init(layoutSize:elementKind:containerAnchor:itemAnchor:)
func NewCollectionLayoutBoundarySupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize INSCollectionLayoutSize, elementKind string, containerAnchor INSCollectionLayoutAnchor, itemAnchor INSCollectionLayoutAnchor) NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutBoundarySupplementaryItemClass().class), objc.Sel("supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:"), layoutSize, objc.String(elementKind), containerAnchor, itemAnchor)
	return NSCollectionLayoutBoundarySupplementaryItemFromID(rv)
}

// Creates an item of the specified size with an array of supplementary items
// to attach to the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:supplementaryItems:)
func NewCollectionLayoutBoundarySupplementaryItemWithLayoutSizeSupplementaryItems(layoutSize INSCollectionLayoutSize, supplementaryItems []NSCollectionLayoutSupplementaryItem) NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutBoundarySupplementaryItemClass().class), objc.Sel("itemWithLayoutSize:supplementaryItems:"), layoutSize, objectivec.IObjectSliceToNSArray(supplementaryItems))
	return NSCollectionLayoutBoundarySupplementaryItemFromID(rv)
}

// A Boolean value that indicates whether a header or footer is pinned to the
// top or bottom visible boundary of the section or layout it’s attached to.
//
// # Discussion
// 
// The default value of this property is [false], which means that the
// boundary supplementary item (header or footer) remains in its original
// position during scrolling, and moves offscreen as its section or layout
// scrolls. Set the value of this property to [true] to pin the boundary
// supplementary item to the visible bounds of the section or layout it’s
// attached to. This way, the boundary supplementary item is shown while any
// portion of the section or layout it’s attached to is visible.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/pinToVisibleBounds
func (c NSCollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("pinToVisibleBounds"))
	return rv
}
func (c NSCollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setPinToVisibleBounds:"), value)
}
// The floating-point value of the boundary supplementary item’s offset from
// the section or layout it’s attached to.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/offset
func (c NSCollectionLayoutBoundarySupplementaryItem) Offset() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("offset"))
	return corefoundation.CGPoint(rv)
}
// The alignment of the boundary supplementary item relative to the section or
// layout it’s attached to.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/alignment
func (c NSCollectionLayoutBoundarySupplementaryItem) Alignment() NSRectAlignment {
	rv := objc.Send[NSRectAlignment](c.ID, objc.Sel("alignment"))
	return NSRectAlignment(rv)
}
// A Boolean value that indicates whether a boundary supplementary item
// extends the content area of the section or layout it’s attached to.
//
// # Discussion
// 
// The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutBoundarySupplementaryItem/extendsBoundary
func (c NSCollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("extendsBoundary"))
	return rv
}
func (c NSCollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setExtendsBoundary:"), value)
}
// An array of the supplementary items that are associated with the boundary
// edges of the entire layout, such as global headers and footers.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayoutconfiguration/boundarysupplementaryitems
func (c NSCollectionLayoutBoundarySupplementaryItem) BoundarySupplementaryItems() INSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("boundarySupplementaryItems"))
	return NSCollectionLayoutBoundarySupplementaryItemFromID(objc.ID(rv))
}
func (c NSCollectionLayoutBoundarySupplementaryItem) SetBoundarySupplementaryItems(value INSCollectionLayoutBoundarySupplementaryItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setBoundarySupplementaryItems:"), value)
}


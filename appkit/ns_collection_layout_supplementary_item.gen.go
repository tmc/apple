// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutSupplementaryItem] class.
var (
	_NSCollectionLayoutSupplementaryItemClass     NSCollectionLayoutSupplementaryItemClass
	_NSCollectionLayoutSupplementaryItemClassOnce sync.Once
)

func getNSCollectionLayoutSupplementaryItemClass() NSCollectionLayoutSupplementaryItemClass {
	_NSCollectionLayoutSupplementaryItemClassOnce.Do(func() {
		_NSCollectionLayoutSupplementaryItemClass = NSCollectionLayoutSupplementaryItemClass{class: objc.GetClass("NSCollectionLayoutSupplementaryItem")}
	})
	return _NSCollectionLayoutSupplementaryItemClass
}

// GetNSCollectionLayoutSupplementaryItemClass returns the class object for NSCollectionLayoutSupplementaryItem.
func GetNSCollectionLayoutSupplementaryItemClass() NSCollectionLayoutSupplementaryItemClass {
	return getNSCollectionLayoutSupplementaryItemClass()
}

type NSCollectionLayoutSupplementaryItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutSupplementaryItemClass) Alloc() NSCollectionLayoutSupplementaryItem {
	rv := objc.Send[NSCollectionLayoutSupplementaryItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object used to add an extra visual decoration to an item in a collection
// view.
//
// # Overview
// 
// You use supplementary items to attach additional views to your content. For
// example, you might attach a badge to an item or a frame around a group. A
// supplementary item follows the index path of the item it’s attached to.
// 
// If you want to create a header or footer for your layout or its sections,
// use a boundary supplementary item () instead.
// 
// Each type of supplementary item must have a unique element kind. Consider
// tracking these strings together in a way that makes it straightforward to
// identify each element, for example:
// 
// Add supplementary items to an item by passing in an array of supplementary
// items when you construct the item:
//
// # Getting the anchors
//
//   - [NSCollectionLayoutSupplementaryItem.ItemAnchor]: The anchor between the supplementary item and the item it’s attached to.
//   - [NSCollectionLayoutSupplementaryItem.ContainerAnchor]: The anchor between the supplementary item and the container it’s attached to.
//
// # Getting the element kind
//
//   - [NSCollectionLayoutSupplementaryItem.ElementKind]: A string that identifies the type of supplementary item.
//
// # Specifying stacking order
//
//   - [NSCollectionLayoutSupplementaryItem.ZIndex]: The vertical stacking order of the supplementary item in relation to other items in the section.
//   - [NSCollectionLayoutSupplementaryItem.SetZIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem
type NSCollectionLayoutSupplementaryItem struct {
	NSCollectionLayoutItem
}

// NSCollectionLayoutSupplementaryItemFromID constructs a [NSCollectionLayoutSupplementaryItem] from an objc.ID.
//
// An object used to add an extra visual decoration to an item in a collection
// view.
func NSCollectionLayoutSupplementaryItemFromID(id objc.ID) NSCollectionLayoutSupplementaryItem {
	return NSCollectionLayoutSupplementaryItem{NSCollectionLayoutItem: NSCollectionLayoutItemFromID(id)}
}
// NOTE: NSCollectionLayoutSupplementaryItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionLayoutSupplementaryItem] class.
//
// # Getting the anchors
//
//   - [INSCollectionLayoutSupplementaryItem.ItemAnchor]: The anchor between the supplementary item and the item it’s attached to.
//   - [INSCollectionLayoutSupplementaryItem.ContainerAnchor]: The anchor between the supplementary item and the container it’s attached to.
//
// # Getting the element kind
//
//   - [INSCollectionLayoutSupplementaryItem.ElementKind]: A string that identifies the type of supplementary item.
//
// # Specifying stacking order
//
//   - [INSCollectionLayoutSupplementaryItem.ZIndex]: The vertical stacking order of the supplementary item in relation to other items in the section.
//   - [INSCollectionLayoutSupplementaryItem.SetZIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem
type INSCollectionLayoutSupplementaryItem interface {
	INSCollectionLayoutItem

	// Topic: Getting the anchors

	// The anchor between the supplementary item and the item it’s attached to.
	ItemAnchor() INSCollectionLayoutAnchor
	// The anchor between the supplementary item and the container it’s attached to.
	ContainerAnchor() INSCollectionLayoutAnchor

	// Topic: Getting the element kind

	// A string that identifies the type of supplementary item.
	ElementKind() string

	// Topic: Specifying stacking order

	// The vertical stacking order of the supplementary item in relation to other items in the section.
	ZIndex() int
	SetZIndex(value int)
}

// Init initializes the instance.
func (c NSCollectionLayoutSupplementaryItem) Init() NSCollectionLayoutSupplementaryItem {
	rv := objc.Send[NSCollectionLayoutSupplementaryItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutSupplementaryItem) Autorelease() NSCollectionLayoutSupplementaryItem {
	rv := objc.Send[NSCollectionLayoutSupplementaryItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutSupplementaryItem creates a new NSCollectionLayoutSupplementaryItem instance.
func NewNSCollectionLayoutSupplementaryItem() NSCollectionLayoutSupplementaryItem {
	class := getNSCollectionLayoutSupplementaryItemClass()
	rv := objc.Send[NSCollectionLayoutSupplementaryItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an item of the specified size.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:)
func NewCollectionLayoutSupplementaryItemWithLayoutSize(layoutSize INSCollectionLayoutSize) NSCollectionLayoutSupplementaryItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutSupplementaryItemClass().class), objc.Sel("itemWithLayoutSize:"), layoutSize)
	return NSCollectionLayoutSupplementaryItemFromID(rv)
}

// Creates a supplementary item of the specified size and element kind, with
// an anchor relative to a container.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/init(layoutSize:elementKind:containerAnchor:)
func NewCollectionLayoutSupplementaryItemWithLayoutSizeElementKindContainerAnchor(layoutSize INSCollectionLayoutSize, elementKind string, containerAnchor INSCollectionLayoutAnchor) NSCollectionLayoutSupplementaryItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutSupplementaryItemClass().class), objc.Sel("supplementaryItemWithLayoutSize:elementKind:containerAnchor:"), layoutSize, objc.String(elementKind), containerAnchor)
	return NSCollectionLayoutSupplementaryItemFromID(rv)
}

// Creates a supplementary item of the specified size and element kind, an
// anchor relative to a container, and an anchor relative to an item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/init(layoutSize:elementKind:containerAnchor:itemAnchor:)
func NewCollectionLayoutSupplementaryItemWithLayoutSizeElementKindContainerAnchorItemAnchor(layoutSize INSCollectionLayoutSize, elementKind string, containerAnchor INSCollectionLayoutAnchor, itemAnchor INSCollectionLayoutAnchor) NSCollectionLayoutSupplementaryItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutSupplementaryItemClass().class), objc.Sel("supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:"), layoutSize, objc.String(elementKind), containerAnchor, itemAnchor)
	return NSCollectionLayoutSupplementaryItemFromID(rv)
}

// Creates an item of the specified size with an array of supplementary items
// to attach to the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:supplementaryItems:)
func NewCollectionLayoutSupplementaryItemWithLayoutSizeSupplementaryItems(layoutSize INSCollectionLayoutSize, supplementaryItems []NSCollectionLayoutSupplementaryItem) NSCollectionLayoutSupplementaryItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutSupplementaryItemClass().class), objc.Sel("itemWithLayoutSize:supplementaryItems:"), layoutSize, objectivec.IObjectSliceToNSArray(supplementaryItems))
	return NSCollectionLayoutSupplementaryItemFromID(rv)
}

// The anchor between the supplementary item and the item it’s attached to.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/itemAnchor
func (c NSCollectionLayoutSupplementaryItem) ItemAnchor() INSCollectionLayoutAnchor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("itemAnchor"))
	return NSCollectionLayoutAnchorFromID(objc.ID(rv))
}
// The anchor between the supplementary item and the container it’s attached
// to.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/containerAnchor
func (c NSCollectionLayoutSupplementaryItem) ContainerAnchor() INSCollectionLayoutAnchor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("containerAnchor"))
	return NSCollectionLayoutAnchorFromID(objc.ID(rv))
}
// A string that identifies the type of supplementary item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/elementKind
func (c NSCollectionLayoutSupplementaryItem) ElementKind() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("elementKind"))
	return foundation.NSStringFromID(rv).String()
}
// The vertical stacking order of the supplementary item in relation to other
// items in the section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSupplementaryItem/zIndex
func (c NSCollectionLayoutSupplementaryItem) ZIndex() int {
	rv := objc.Send[int](c.ID, objc.Sel("zIndex"))
	return rv
}
func (c NSCollectionLayoutSupplementaryItem) SetZIndex(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setZIndex:"), value)
}


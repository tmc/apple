// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutItem] class.
var (
	_NSCollectionLayoutItemClass     NSCollectionLayoutItemClass
	_NSCollectionLayoutItemClassOnce sync.Once
)

func getNSCollectionLayoutItemClass() NSCollectionLayoutItemClass {
	_NSCollectionLayoutItemClassOnce.Do(func() {
		_NSCollectionLayoutItemClass = NSCollectionLayoutItemClass{class: objc.GetClass("NSCollectionLayoutItem")}
	})
	return _NSCollectionLayoutItemClass
}

// GetNSCollectionLayoutItemClass returns the class object for NSCollectionLayoutItem.
func GetNSCollectionLayoutItemClass() NSCollectionLayoutItemClass {
	return getNSCollectionLayoutItemClass()
}

type NSCollectionLayoutItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutItemClass) Alloc() NSCollectionLayoutItem {
	rv := objc.Send[NSCollectionLayoutItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The most basic component of a collection view’s layout.
//
// # Overview
// 
// An item is a blueprint for how to size, space, and arrange an individual
// piece of content in your collection view. An item represents a single view
// that’s rendered onscreen. Generally, an item is a cell, but items can be
// supplementary views like headers, footers, and other decorations.
// 
// For example, in the Photos app, an item might represent a single photo. In
// the App Store app, an item might be a cell displaying information about an
// individual app in a list of featured apps, such as the app icon, app name,
// tagline, and download button.
// 
// [media-3568665]
// 
// Each item specifies its own size in terms of a width dimension and a height
// dimension. Items can express their dimensions relative to their container,
// as an absolute value, or as an estimated value that might change at
// runtime, like in response to a change in system font size. For more
// information, see [NSCollectionLayoutDimension].
// 
// You combine items into groups that determine how those items are arranged
// in relation to each other. For more information, see
// [NSCollectionLayoutGroup].
//
// # Getting an item’s size
//
//   - [NSCollectionLayoutItem.LayoutSize]: The item’s size expressed in width and height dimensions.
//
// # Getting supplementary items
//
//   - [NSCollectionLayoutItem.SupplementaryItems]: An array of the supplementary items attached to the item.
//
// # Configuring spacing and insets
//
//   - [NSCollectionLayoutItem.EdgeSpacing]: The amount of space added around the boundaries of the item between other items and this item’s container.
//   - [NSCollectionLayoutItem.SetEdgeSpacing]
//   - [NSCollectionLayoutItem.ContentInsets]: The amount of space added around the content of the item to adjust its final size after its position is computed.
//   - [NSCollectionLayoutItem.SetContentInsets]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem
type NSCollectionLayoutItem struct {
	objectivec.Object
}

// NSCollectionLayoutItemFromID constructs a [NSCollectionLayoutItem] from an objc.ID.
//
// The most basic component of a collection view’s layout.
func NSCollectionLayoutItemFromID(id objc.ID) NSCollectionLayoutItem {
	return NSCollectionLayoutItem{objectivec.Object{ID: id}}
}
// NOTE: NSCollectionLayoutItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionLayoutItem] class.
//
// # Getting an item’s size
//
//   - [INSCollectionLayoutItem.LayoutSize]: The item’s size expressed in width and height dimensions.
//
// # Getting supplementary items
//
//   - [INSCollectionLayoutItem.SupplementaryItems]: An array of the supplementary items attached to the item.
//
// # Configuring spacing and insets
//
//   - [INSCollectionLayoutItem.EdgeSpacing]: The amount of space added around the boundaries of the item between other items and this item’s container.
//   - [INSCollectionLayoutItem.SetEdgeSpacing]
//   - [INSCollectionLayoutItem.ContentInsets]: The amount of space added around the content of the item to adjust its final size after its position is computed.
//   - [INSCollectionLayoutItem.SetContentInsets]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem
type INSCollectionLayoutItem interface {
	objectivec.IObject

	// Topic: Getting an item’s size

	// The item’s size expressed in width and height dimensions.
	LayoutSize() INSCollectionLayoutSize

	// Topic: Getting supplementary items

	// An array of the supplementary items attached to the item.
	SupplementaryItems() []NSCollectionLayoutSupplementaryItem

	// Topic: Configuring spacing and insets

	// The amount of space added around the boundaries of the item between other items and this item’s container.
	EdgeSpacing() INSCollectionLayoutEdgeSpacing
	SetEdgeSpacing(value INSCollectionLayoutEdgeSpacing)
	// The amount of space added around the content of the item to adjust its final size after its position is computed.
	ContentInsets() NSDirectionalEdgeInsets
	SetContentInsets(value NSDirectionalEdgeInsets)
}

// Init initializes the instance.
func (c NSCollectionLayoutItem) Init() NSCollectionLayoutItem {
	rv := objc.Send[NSCollectionLayoutItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutItem) Autorelease() NSCollectionLayoutItem {
	rv := objc.Send[NSCollectionLayoutItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutItem creates a new NSCollectionLayoutItem instance.
func NewNSCollectionLayoutItem() NSCollectionLayoutItem {
	class := getNSCollectionLayoutItemClass()
	rv := objc.Send[NSCollectionLayoutItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an item of the specified size.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:)
func NewCollectionLayoutItemWithLayoutSize(layoutSize INSCollectionLayoutSize) NSCollectionLayoutItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutItemClass().class), objc.Sel("itemWithLayoutSize:"), layoutSize)
	return NSCollectionLayoutItemFromID(rv)
}

// Creates an item of the specified size with an array of supplementary items
// to attach to the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:supplementaryItems:)
func NewCollectionLayoutItemWithLayoutSizeSupplementaryItems(layoutSize INSCollectionLayoutSize, supplementaryItems []NSCollectionLayoutSupplementaryItem) NSCollectionLayoutItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutItemClass().class), objc.Sel("itemWithLayoutSize:supplementaryItems:"), layoutSize, objectivec.IObjectSliceToNSArray(supplementaryItems))
	return NSCollectionLayoutItemFromID(rv)
}

// The item’s size expressed in width and height dimensions.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/layoutSize
func (c NSCollectionLayoutItem) LayoutSize() INSCollectionLayoutSize {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("layoutSize"))
	return NSCollectionLayoutSizeFromID(objc.ID(rv))
}
// An array of the supplementary items attached to the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/supplementaryItems
func (c NSCollectionLayoutItem) SupplementaryItems() []NSCollectionLayoutSupplementaryItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supplementaryItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSCollectionLayoutSupplementaryItem {
		return NSCollectionLayoutSupplementaryItemFromID(id)
	})
}
// The amount of space added around the boundaries of the item between other
// items and this item’s container.
//
// # Discussion
// 
// Use this property to adjust the position of the item in relation to its
// container and other items. For example, you can use this property to apply
// extra space to the trailing edge of each item. Edge spacing is applied
// before applying [ContentInsets].
// 
// The following diagram shows the result of applying 2 points of trailing
// edge spacing to the items in a group:
// 
// [media-3572326]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/edgeSpacing
func (c NSCollectionLayoutItem) EdgeSpacing() INSCollectionLayoutEdgeSpacing {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("edgeSpacing"))
	return NSCollectionLayoutEdgeSpacingFromID(objc.ID(rv))
}
func (c NSCollectionLayoutItem) SetEdgeSpacing(value INSCollectionLayoutEdgeSpacing) {
	objc.Send[struct{}](c.ID, objc.Sel("setEdgeSpacing:"), value)
}
// The amount of space added around the content of the item to adjust its
// final size after its position is computed.
//
// # Discussion
// 
// You can use this property within a grid layout to apply even spacing around
// each edge of each item. Content insets are applied after applying
// [EdgeSpacing].
// 
// The following diagram shows the result of applying 2 points of content
// insets to each edge of each item in a group.
// 
// [media-3570427]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/contentInsets
func (c NSCollectionLayoutItem) ContentInsets() NSDirectionalEdgeInsets {
	rv := objc.Send[NSDirectionalEdgeInsets](c.ID, objc.Sel("contentInsets"))
	return NSDirectionalEdgeInsets(rv)
}
func (c NSCollectionLayoutItem) SetContentInsets(value NSDirectionalEdgeInsets) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentInsets:"), value)
}


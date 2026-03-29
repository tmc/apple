// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutDecorationItem] class.
var (
	_NSCollectionLayoutDecorationItemClass     NSCollectionLayoutDecorationItemClass
	_NSCollectionLayoutDecorationItemClassOnce sync.Once
)

func getNSCollectionLayoutDecorationItemClass() NSCollectionLayoutDecorationItemClass {
	_NSCollectionLayoutDecorationItemClassOnce.Do(func() {
		_NSCollectionLayoutDecorationItemClass = NSCollectionLayoutDecorationItemClass{class: objc.GetClass("NSCollectionLayoutDecorationItem")}
	})
	return _NSCollectionLayoutDecorationItemClass
}

// GetNSCollectionLayoutDecorationItemClass returns the class object for NSCollectionLayoutDecorationItem.
func GetNSCollectionLayoutDecorationItemClass() NSCollectionLayoutDecorationItemClass {
	return getNSCollectionLayoutDecorationItemClass()
}

type NSCollectionLayoutDecorationItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCollectionLayoutDecorationItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutDecorationItemClass) Alloc() NSCollectionLayoutDecorationItem {
	rv := objc.Send[NSCollectionLayoutDecorationItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object used to add a background to a section of a collection view.
//
// # Overview
// 
// Each type of decoration item must have a unique element kind. Consider
// tracking these strings together in a way that makes it straightforward to
// identify each element, for example:
// 
// Add a background to a section by setting that section’s [NSCollectionLayoutDecorationItem.DecorationItems]
// property:
//
// # Getting the element kind
//
//   - [NSCollectionLayoutDecorationItem.ElementKind]: A string that identifies the type of decoration item.
//
// # Specifying stacking order
//
//   - [NSCollectionLayoutDecorationItem.ZIndex]: The vertical stacking order of the decoration item in relation to other items in the section.
//   - [NSCollectionLayoutDecorationItem.SetZIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDecorationItem
type NSCollectionLayoutDecorationItem struct {
	NSCollectionLayoutItem
}

// NSCollectionLayoutDecorationItemFromID constructs a [NSCollectionLayoutDecorationItem] from an objc.ID.
//
// An object used to add a background to a section of a collection view.
func NSCollectionLayoutDecorationItemFromID(id objc.ID) NSCollectionLayoutDecorationItem {
	return NSCollectionLayoutDecorationItem{NSCollectionLayoutItem: NSCollectionLayoutItemFromID(id)}
}
// NOTE: NSCollectionLayoutDecorationItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionLayoutDecorationItem] class.
//
// # Getting the element kind
//
//   - [INSCollectionLayoutDecorationItem.ElementKind]: A string that identifies the type of decoration item.
//
// # Specifying stacking order
//
//   - [INSCollectionLayoutDecorationItem.ZIndex]: The vertical stacking order of the decoration item in relation to other items in the section.
//   - [INSCollectionLayoutDecorationItem.SetZIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDecorationItem
type INSCollectionLayoutDecorationItem interface {
	INSCollectionLayoutItem

	// Topic: Getting the element kind

	// A string that identifies the type of decoration item.
	ElementKind() string

	// Topic: Specifying stacking order

	// The vertical stacking order of the decoration item in relation to other items in the section.
	ZIndex() int
	SetZIndex(value int)

	// An array of the decoration items that are anchored to the section, such as background decoration views.
	DecorationItems() INSCollectionLayoutDecorationItem
	SetDecorationItems(value INSCollectionLayoutDecorationItem)
}

// Init initializes the instance.
func (c NSCollectionLayoutDecorationItem) Init() NSCollectionLayoutDecorationItem {
	rv := objc.Send[NSCollectionLayoutDecorationItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutDecorationItem) Autorelease() NSCollectionLayoutDecorationItem {
	rv := objc.Send[NSCollectionLayoutDecorationItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutDecorationItem creates a new NSCollectionLayoutDecorationItem instance.
func NewNSCollectionLayoutDecorationItem() NSCollectionLayoutDecorationItem {
	class := getNSCollectionLayoutDecorationItemClass()
	rv := objc.Send[NSCollectionLayoutDecorationItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an item of the specified size.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:)
func NewCollectionLayoutDecorationItemWithLayoutSize(layoutSize INSCollectionLayoutSize) NSCollectionLayoutDecorationItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutDecorationItemClass().class), objc.Sel("itemWithLayoutSize:"), layoutSize)
	return NSCollectionLayoutDecorationItemFromID(rv)
}

// Creates an item of the specified size with an array of supplementary items
// to attach to the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:supplementaryItems:)
func NewCollectionLayoutDecorationItemWithLayoutSizeSupplementaryItems(layoutSize INSCollectionLayoutSize, supplementaryItems []NSCollectionLayoutSupplementaryItem) NSCollectionLayoutDecorationItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutDecorationItemClass().class), objc.Sel("itemWithLayoutSize:supplementaryItems:"), layoutSize, objectivec.IObjectSliceToNSArray(supplementaryItems))
	return NSCollectionLayoutDecorationItemFromID(rv)
}

// Creates a section background with a string to identify the element kind.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDecorationItem/background(elementKind:)
func (_NSCollectionLayoutDecorationItemClass NSCollectionLayoutDecorationItemClass) BackgroundDecorationItemWithElementKind(elementKind string) NSCollectionLayoutDecorationItem {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutDecorationItemClass.class), objc.Sel("backgroundDecorationItemWithElementKind:"), objc.String(elementKind))
	return NSCollectionLayoutDecorationItemFromID(rv)
}

// A string that identifies the type of decoration item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDecorationItem/elementKind
func (c NSCollectionLayoutDecorationItem) ElementKind() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("elementKind"))
	return foundation.NSStringFromID(rv).String()
}
// The vertical stacking order of the decoration item in relation to other
// items in the section.
//
// # Discussion
// 
// The default value of this property is `0`, which means the decoration item
// appears below all other items in the section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDecorationItem/zIndex
func (c NSCollectionLayoutDecorationItem) ZIndex() int {
	rv := objc.Send[int](c.ID, objc.Sel("zIndex"))
	return rv
}
func (c NSCollectionLayoutDecorationItem) SetZIndex(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setZIndex:"), value)
}
// An array of the decoration items that are anchored to the section, such as
// background decoration views.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionlayoutsection/decorationitems
func (c NSCollectionLayoutDecorationItem) DecorationItems() INSCollectionLayoutDecorationItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("decorationItems"))
	return NSCollectionLayoutDecorationItemFromID(objc.ID(rv))
}
func (c NSCollectionLayoutDecorationItem) SetDecorationItems(value INSCollectionLayoutDecorationItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setDecorationItems:"), value)
}


// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutGroup] class.
var (
	_NSCollectionLayoutGroupClass     NSCollectionLayoutGroupClass
	_NSCollectionLayoutGroupClassOnce sync.Once
)

func getNSCollectionLayoutGroupClass() NSCollectionLayoutGroupClass {
	_NSCollectionLayoutGroupClassOnce.Do(func() {
		_NSCollectionLayoutGroupClass = NSCollectionLayoutGroupClass{class: objc.GetClass("NSCollectionLayoutGroup")}
	})
	return _NSCollectionLayoutGroupClass
}

// GetNSCollectionLayoutGroupClass returns the class object for NSCollectionLayoutGroup.
func GetNSCollectionLayoutGroupClass() NSCollectionLayoutGroupClass {
	return getNSCollectionLayoutGroupClass()
}

type NSCollectionLayoutGroupClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutGroupClass) Alloc() NSCollectionLayoutGroup {
	rv := objc.Send[NSCollectionLayoutGroup](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A container for a set of items that lays out the items along a path.
//
// # Overview
// 
// Groups determine how the items in a collection view lay out in relation to
// each other. A group might lay out its items in a horizontal row, a vertical
// column, or a custom arrangement. A group determines the rules for how items
// are rendered in relation to each other, but in itself doesn’t render any
// content.
// 
// For example, in the Photos app, a group of items is a row of photos. In the
// App Store app, a group might be a single column of cells (items) arranged
// in a vertical column.
// 
// [media-3568663]
// 
// Each group specifies its own size in terms of a width dimension and a
// height dimension. Groups can express their dimensions relative to their
// container, as an absolute value, or as an estimated value that might change
// at runtime, like in response to a change in system font size. For more
// information, see [NSCollectionLayoutDimension].
// 
// Because a group is a subclass of [NSCollectionLayoutItem], it behaves like
// an item. You can combine a group with other items and groups into more
// complex layouts.
// 
// [media-3568666]
// 
// After you configure a group, you must initialize a section
// ([NSCollectionLayoutSection]) of your collection view layout with that
// group.
//
// # Getting the group’s items
//
//   - [NSCollectionLayoutGroup.Subitems]: An array of the items contained in the group.
//
// # Configuring group spacing
//
//   - [NSCollectionLayoutGroup.InterItemSpacing]: The amount of space between the items in the group.
//   - [NSCollectionLayoutGroup.SetInterItemSpacing]
//
// # Debugging group layout
//
//   - [NSCollectionLayoutGroup.VisualDescription]: Returns a string with an ASCII representation of the group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup
type NSCollectionLayoutGroup struct {
	NSCollectionLayoutItem
}

// NSCollectionLayoutGroupFromID constructs a [NSCollectionLayoutGroup] from an objc.ID.
//
// A container for a set of items that lays out the items along a path.
func NSCollectionLayoutGroupFromID(id objc.ID) NSCollectionLayoutGroup {
	return NSCollectionLayoutGroup{NSCollectionLayoutItem: NSCollectionLayoutItemFromID(id)}
}
// NOTE: NSCollectionLayoutGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionLayoutGroup] class.
//
// # Getting the group’s items
//
//   - [INSCollectionLayoutGroup.Subitems]: An array of the items contained in the group.
//
// # Configuring group spacing
//
//   - [INSCollectionLayoutGroup.InterItemSpacing]: The amount of space between the items in the group.
//   - [INSCollectionLayoutGroup.SetInterItemSpacing]
//
// # Debugging group layout
//
//   - [INSCollectionLayoutGroup.VisualDescription]: Returns a string with an ASCII representation of the group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup
type INSCollectionLayoutGroup interface {
	INSCollectionLayoutItem

	// Topic: Getting the group’s items

	// An array of the items contained in the group.
	Subitems() []NSCollectionLayoutItem

	// Topic: Configuring group spacing

	// The amount of space between the items in the group.
	InterItemSpacing() INSCollectionLayoutSpacing
	SetInterItemSpacing(value INSCollectionLayoutSpacing)

	// Topic: Debugging group layout

	// Returns a string with an ASCII representation of the group.
	VisualDescription() string
}

// Init initializes the instance.
func (c NSCollectionLayoutGroup) Init() NSCollectionLayoutGroup {
	rv := objc.Send[NSCollectionLayoutGroup](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutGroup) Autorelease() NSCollectionLayoutGroup {
	rv := objc.Send[NSCollectionLayoutGroup](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutGroup creates a new NSCollectionLayoutGroup instance.
func NewNSCollectionLayoutGroup() NSCollectionLayoutGroup {
	class := getNSCollectionLayoutGroupClass()
	rv := objc.Send[NSCollectionLayoutGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an item of the specified size.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:)
func NewCollectionLayoutGroupItemWithLayoutSize(layoutSize INSCollectionLayoutSize) NSCollectionLayoutGroup {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutGroupClass().class), objc.Sel("itemWithLayoutSize:"), layoutSize)
	return NSCollectionLayoutGroupFromID(rv)
}

// Creates an item of the specified size with an array of supplementary items
// to attach to the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutItem/init(layoutSize:supplementaryItems:)
func NewCollectionLayoutGroupItemWithLayoutSizeSupplementaryItems(layoutSize INSCollectionLayoutSize, supplementaryItems []NSCollectionLayoutSupplementaryItem) NSCollectionLayoutGroup {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutGroupClass().class), objc.Sel("itemWithLayoutSize:supplementaryItems:"), layoutSize, objectivec.IObjectSliceToNSArray(supplementaryItems))
	return NSCollectionLayoutGroupFromID(rv)
}

// Returns a string with an ASCII representation of the group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/visualDescription()
func (c NSCollectionLayoutGroup) VisualDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("visualDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Creates a group of the specified size, containing an array of items
// arranged in a horizontal line.
//
// layoutSize: The group’s size.
//
// subitems: The subitems to include.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/horizontal(layoutSize:subitems:)
func (_NSCollectionLayoutGroupClass NSCollectionLayoutGroupClass) HorizontalGroupWithLayoutSizeSubitems(layoutSize INSCollectionLayoutSize, subitems []NSCollectionLayoutItem) NSCollectionLayoutGroup {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutGroupClass.class), objc.Sel("horizontalGroupWithLayoutSize:subitems:"), layoutSize, objectivec.IObjectSliceToNSArray(subitems))
	return NSCollectionLayoutGroupFromID(rv)
}

// Creates a group of the specified size, containing an array of equally sized
// items arranged in a horizontal line up to the number specified by count.
//
// # Discussion
// 
// When you set a value for the [InterItemSpacing] property after using this
// initializer, the group keeps the same number of items and automatically
// resizes them to add the extra specified spacing between them.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/horizontal(layoutSize:subitem:count:)
func (_NSCollectionLayoutGroupClass NSCollectionLayoutGroupClass) HorizontalGroupWithLayoutSizeSubitemCount(layoutSize INSCollectionLayoutSize, subitem INSCollectionLayoutItem, count int) NSCollectionLayoutGroup {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutGroupClass.class), objc.Sel("horizontalGroupWithLayoutSize:subitem:count:"), layoutSize, subitem, count)
	return NSCollectionLayoutGroupFromID(rv)
}

// Creates a group of the specified size, containing an array of items
// arranged in a vertical line.
//
// layoutSize: The group’s size.
//
// subitems: The subitems to include.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/vertical(layoutSize:subitems:)
func (_NSCollectionLayoutGroupClass NSCollectionLayoutGroupClass) VerticalGroupWithLayoutSizeSubitems(layoutSize INSCollectionLayoutSize, subitems []NSCollectionLayoutItem) NSCollectionLayoutGroup {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutGroupClass.class), objc.Sel("verticalGroupWithLayoutSize:subitems:"), layoutSize, objectivec.IObjectSliceToNSArray(subitems))
	return NSCollectionLayoutGroupFromID(rv)
}

// Creates a group of the specified size, containing an array of equally sized
// items arranged in a vertical line up to the number specified by count.
//
// # Discussion
// 
// When you set a value for the [InterItemSpacing] property after using this
// initializer, the group keeps the same number of items and automatically
// resizes them to add the extra specified spacing between them.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/vertical(layoutSize:subitem:count:)
func (_NSCollectionLayoutGroupClass NSCollectionLayoutGroupClass) VerticalGroupWithLayoutSizeSubitemCount(layoutSize INSCollectionLayoutSize, subitem INSCollectionLayoutItem, count int) NSCollectionLayoutGroup {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutGroupClass.class), objc.Sel("verticalGroupWithLayoutSize:subitem:count:"), layoutSize, subitem, count)
	return NSCollectionLayoutGroupFromID(rv)
}

// Creates a group of the specified size, with an item provider that creates a
// custom arrangement for those items.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/custom(layoutSize:itemProvider:)
func (_NSCollectionLayoutGroupClass NSCollectionLayoutGroupClass) CustomGroupWithLayoutSizeItemProvider(layoutSize INSCollectionLayoutSize, itemProvider NSCollectionLayoutGroupCustomItemProvider) NSCollectionLayoutGroup {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutGroupClass.class), objc.Sel("customGroupWithLayoutSize:itemProvider:"), layoutSize, itemProvider)
	return NSCollectionLayoutGroupFromID(rv)
}

// An array of the items contained in the group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/subitems
func (c NSCollectionLayoutGroup) Subitems() []NSCollectionLayoutItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("subitems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSCollectionLayoutItem {
		return NSCollectionLayoutItemFromID(id)
	})
}

// The amount of space between the items in the group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroup/interItemSpacing
func (c NSCollectionLayoutGroup) InterItemSpacing() INSCollectionLayoutSpacing {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("interItemSpacing"))
	return NSCollectionLayoutSpacingFromID(objc.ID(rv))
}
func (c NSCollectionLayoutGroup) SetInterItemSpacing(value INSCollectionLayoutSpacing) {
	objc.Send[struct{}](c.ID, objc.Sel("setInterItemSpacing:"), value)
}


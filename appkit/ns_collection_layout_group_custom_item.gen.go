// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutGroupCustomItem] class.
var (
	_NSCollectionLayoutGroupCustomItemClass     NSCollectionLayoutGroupCustomItemClass
	_NSCollectionLayoutGroupCustomItemClassOnce sync.Once
)

func getNSCollectionLayoutGroupCustomItemClass() NSCollectionLayoutGroupCustomItemClass {
	_NSCollectionLayoutGroupCustomItemClassOnce.Do(func() {
		_NSCollectionLayoutGroupCustomItemClass = NSCollectionLayoutGroupCustomItemClass{class: objc.GetClass("NSCollectionLayoutGroupCustomItem")}
	})
	return _NSCollectionLayoutGroupCustomItemClass
}

// GetNSCollectionLayoutGroupCustomItemClass returns the class object for NSCollectionLayoutGroupCustomItem.
func GetNSCollectionLayoutGroupCustomItemClass() NSCollectionLayoutGroupCustomItemClass {
	return getNSCollectionLayoutGroupCustomItemClass()
}

type NSCollectionLayoutGroupCustomItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutGroupCustomItemClass) Alloc() NSCollectionLayoutGroupCustomItem {
	rv := objc.Send[NSCollectionLayoutGroupCustomItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An item used in a group with a custom layout arrangement.
//
// # Overview
// 
// You use a custom item if you want to specify a layout with a custom
// arrangement, like a radial or diagonal layout. You use custom items within
// a group that’s created with [CustomGroupWithLayoutSizeItemProvider].
// 
// Instead of providing a layout size for the custom item, like you do when
// you create an [NSCollectionLayoutItem], you provide a frame instead.
//
// # Getting the frame
//
//   - [NSCollectionLayoutGroupCustomItem.Frame]: The frame of the custom item.
//
// # Specifying stacking order
//
//   - [NSCollectionLayoutGroupCustomItem.ZIndex]: The vertical stacking order of the custom item in relation to other items in the group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem
type NSCollectionLayoutGroupCustomItem struct {
	objectivec.Object
}

// NSCollectionLayoutGroupCustomItemFromID constructs a [NSCollectionLayoutGroupCustomItem] from an objc.ID.
//
// An item used in a group with a custom layout arrangement.
func NSCollectionLayoutGroupCustomItemFromID(id objc.ID) NSCollectionLayoutGroupCustomItem {
	return NSCollectionLayoutGroupCustomItem{objectivec.Object{id}}
}
// NOTE: NSCollectionLayoutGroupCustomItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCollectionLayoutGroupCustomItem] class.
//
// # Getting the frame
//
//   - [INSCollectionLayoutGroupCustomItem.Frame]: The frame of the custom item.
//
// # Specifying stacking order
//
//   - [INSCollectionLayoutGroupCustomItem.ZIndex]: The vertical stacking order of the custom item in relation to other items in the group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem
type INSCollectionLayoutGroupCustomItem interface {
	objectivec.IObject

	// Topic: Getting the frame

	// The frame of the custom item.
	Frame() corefoundation.CGRect

	// Topic: Specifying stacking order

	// The vertical stacking order of the custom item in relation to other items in the group.
	ZIndex() int
}





// Init initializes the instance.
func (c NSCollectionLayoutGroupCustomItem) Init() NSCollectionLayoutGroupCustomItem {
	rv := objc.Send[NSCollectionLayoutGroupCustomItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutGroupCustomItem) Autorelease() NSCollectionLayoutGroupCustomItem {
	rv := objc.Send[NSCollectionLayoutGroupCustomItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutGroupCustomItem creates a new NSCollectionLayoutGroupCustomItem instance.
func NewNSCollectionLayoutGroupCustomItem() NSCollectionLayoutGroupCustomItem {
	class := getNSCollectionLayoutGroupCustomItemClass()
	rv := objc.Send[NSCollectionLayoutGroupCustomItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a custom item with the specified frame.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/init(frame:)
func NewCollectionLayoutGroupCustomItemWithFrame(frame corefoundation.CGRect) NSCollectionLayoutGroupCustomItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutGroupCustomItemClass().class), objc.Sel("customItemWithFrame:"), frame)
	return NSCollectionLayoutGroupCustomItemFromID(rv)
}


// Creates a custom item with the specified frame and vertical stacking order
// in relation to other items in the group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/init(frame:zIndex:)
func NewCollectionLayoutGroupCustomItemWithFrameZIndex(frame corefoundation.CGRect, zIndex int) NSCollectionLayoutGroupCustomItem {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutGroupCustomItemClass().class), objc.Sel("customItemWithFrame:zIndex:"), frame, zIndex)
	return NSCollectionLayoutGroupCustomItemFromID(rv)
}


















// The frame of the custom item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/frame
func (c NSCollectionLayoutGroupCustomItem) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}



// The vertical stacking order of the custom item in relation to other items
// in the group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/zIndex
func (c NSCollectionLayoutGroupCustomItem) ZIndex() int {
	rv := objc.Send[int](c.ID, objc.Sel("zIndex"))
	return rv
}


























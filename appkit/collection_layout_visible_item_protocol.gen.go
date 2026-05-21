// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An item that’s currently visible within the bounds of a section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem
type NSCollectionLayoutVisibleItem interface {
	objectivec.IObject

	// A Boolean value that determines whether the item is hidden.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/isHidden
	IsHidden() bool

	// The name of the item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/name
	Name() string

	// A string that identifies the type of item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/representedElementKind
	RepresentedElementKind() string

	// A category that identifies the item, such as decoration or supplementary view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/representedElementCategory
	RepresentedElementCategory() NSCollectionElementCategory

	// The index path of the item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/indexPath
	IndexPath() foundation.NSIndexPath

	// The transparency of the item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/alpha
	Alpha() float64
	SetAlpha(value float64)

	// A Boolean value that determines whether the item is hidden.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/isHidden
	Hidden() bool
	SetHidden(value bool)

	// The frame rectangle, which describes the item’s location and size in its section’s coordinate system.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/frame
	Frame() corefoundation.CGRect

	// The bounds rectangle, which describes the item’s location and size in its own coordinate system.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/bounds
	Bounds() corefoundation.CGRect

	// The center point of the item’s frame rectangle.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/center
	Center() corefoundation.CGPoint
	SetCenter(value corefoundation.CGPoint)

	// The vertical stacking order of the item in relation to other items in the section.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/zIndex
	ZIndex() int
	SetZIndex(value int)
}

// NSCollectionLayoutVisibleItemObject wraps an existing Objective-C object that conforms to the NSCollectionLayoutVisibleItem protocol.
type NSCollectionLayoutVisibleItemObject struct {
	objectivec.Object
}

func (o NSCollectionLayoutVisibleItemObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCollectionLayoutVisibleItemObjectFromID constructs a [NSCollectionLayoutVisibleItemObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCollectionLayoutVisibleItemObjectFromID(id objc.ID) NSCollectionLayoutVisibleItemObject {
	return NSCollectionLayoutVisibleItemObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value that determines whether the item is hidden.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/isHidden
func (o NSCollectionLayoutVisibleItemObject) IsHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isHidden"))
	return rv
}

// The name of the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/name
func (o NSCollectionLayoutVisibleItemObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// A string that identifies the type of item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/representedElementKind
func (o NSCollectionLayoutVisibleItemObject) RepresentedElementKind() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("representedElementKind"))
	return foundation.NSStringFromID(rv).String()
}

// A category that identifies the item, such as decoration or supplementary
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/representedElementCategory
func (o NSCollectionLayoutVisibleItemObject) RepresentedElementCategory() NSCollectionElementCategory {
	rv := objc.Send[NSCollectionElementCategory](o.ID, objc.Sel("representedElementCategory"))
	return NSCollectionElementCategory(rv)
}

// The index path of the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/indexPath
func (o NSCollectionLayoutVisibleItemObject) IndexPath() foundation.NSIndexPath {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("indexPath"))
	return foundation.NSIndexPathFromID(rv)
}

// The transparency of the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/alpha
func (o NSCollectionLayoutVisibleItemObject) Alpha() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("alpha"))
	return float64(rv)
}

func (o NSCollectionLayoutVisibleItemObject) SetAlpha(value float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setAlpha:"), value)
}

// A Boolean value that determines whether the item is hidden.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/isHidden
func (o NSCollectionLayoutVisibleItemObject) Hidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isHidden"))
	return bool(rv)
}

func (o NSCollectionLayoutVisibleItemObject) SetHidden(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setHidden:"), value)
}

// The frame rectangle, which describes the item’s location and size in its
// section’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/frame
func (o NSCollectionLayoutVisibleItemObject) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}

// The bounds rectangle, which describes the item’s location and size in its
// own coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/bounds
func (o NSCollectionLayoutVisibleItemObject) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}

// The center point of the item’s frame rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/center
func (o NSCollectionLayoutVisibleItemObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return corefoundation.CGPoint(rv)
}

func (o NSCollectionLayoutVisibleItemObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

// The vertical stacking order of the item in relation to other items in the
// section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutVisibleItem/zIndex
func (o NSCollectionLayoutVisibleItemObject) ZIndex() int {
	rv := objc.Send[int](o.ID, objc.Sel("zIndex"))
	return int(rv)
}

func (o NSCollectionLayoutVisibleItemObject) SetZIndex(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setZIndex:"), value)
}

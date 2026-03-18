// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutEdgeSpacing] class.
var (
	_NSCollectionLayoutEdgeSpacingClass     NSCollectionLayoutEdgeSpacingClass
	_NSCollectionLayoutEdgeSpacingClassOnce sync.Once
)

func getNSCollectionLayoutEdgeSpacingClass() NSCollectionLayoutEdgeSpacingClass {
	_NSCollectionLayoutEdgeSpacingClassOnce.Do(func() {
		_NSCollectionLayoutEdgeSpacingClass = NSCollectionLayoutEdgeSpacingClass{class: objc.GetClass("NSCollectionLayoutEdgeSpacing")}
	})
	return _NSCollectionLayoutEdgeSpacingClass
}

// GetNSCollectionLayoutEdgeSpacingClass returns the class object for NSCollectionLayoutEdgeSpacing.
func GetNSCollectionLayoutEdgeSpacingClass() NSCollectionLayoutEdgeSpacingClass {
	return getNSCollectionLayoutEdgeSpacingClass()
}

type NSCollectionLayoutEdgeSpacingClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutEdgeSpacingClass) Alloc() NSCollectionLayoutEdgeSpacing {
	rv := objc.Send[NSCollectionLayoutEdgeSpacing](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that defines the space around the edges of items in a collection
// view.
//
// # Overview
// 
// You use edge spacing to create additional spacing around the edges of an
// item to adjust the position of the item in relation to its container and
// other items.
// 
// The leading and trailing spaces within edge spacing differ in left-to-right
// versus right-to-left environments. In a left-to-right environment, the
// leading space is on the left, and the trailing space is on the right. In a
// right-to-left environment, the leading space is on the right, and the
// trailing space is on the left. This difference ensures that your collection
// view layout is built with support for right-to-left languages.
// 
// The following diagram shows the difference between adding 2 points of
// trailing edge spacing in a left-to-right versus a right-to-left
// environment.
// 
// [media-3570381]
//
// # Getting the edge spacing
//
//   - [NSCollectionLayoutEdgeSpacing.Leading]: The leading edge spacing value.
//   - [NSCollectionLayoutEdgeSpacing.Top]: The top edge spacing value.
//   - [NSCollectionLayoutEdgeSpacing.Trailing]: The trailing edge spacing value.
//   - [NSCollectionLayoutEdgeSpacing.Bottom]: The bottom edge spacing value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing
type NSCollectionLayoutEdgeSpacing struct {
	objectivec.Object
}

// NSCollectionLayoutEdgeSpacingFromID constructs a [NSCollectionLayoutEdgeSpacing] from an objc.ID.
//
// An object that defines the space around the edges of items in a collection
// view.
func NSCollectionLayoutEdgeSpacingFromID(id objc.ID) NSCollectionLayoutEdgeSpacing {
	return NSCollectionLayoutEdgeSpacing{objectivec.Object{ID: id}}
}
// NOTE: NSCollectionLayoutEdgeSpacing adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionLayoutEdgeSpacing] class.
//
// # Getting the edge spacing
//
//   - [INSCollectionLayoutEdgeSpacing.Leading]: The leading edge spacing value.
//   - [INSCollectionLayoutEdgeSpacing.Top]: The top edge spacing value.
//   - [INSCollectionLayoutEdgeSpacing.Trailing]: The trailing edge spacing value.
//   - [INSCollectionLayoutEdgeSpacing.Bottom]: The bottom edge spacing value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing
type INSCollectionLayoutEdgeSpacing interface {
	objectivec.IObject

	// Topic: Getting the edge spacing

	// The leading edge spacing value.
	Leading() INSCollectionLayoutSpacing
	// The top edge spacing value.
	Top() INSCollectionLayoutSpacing
	// The trailing edge spacing value.
	Trailing() INSCollectionLayoutSpacing
	// The bottom edge spacing value.
	Bottom() INSCollectionLayoutSpacing
}

// Init initializes the instance.
func (c NSCollectionLayoutEdgeSpacing) Init() NSCollectionLayoutEdgeSpacing {
	rv := objc.Send[NSCollectionLayoutEdgeSpacing](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutEdgeSpacing) Autorelease() NSCollectionLayoutEdgeSpacing {
	rv := objc.Send[NSCollectionLayoutEdgeSpacing](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutEdgeSpacing creates a new NSCollectionLayoutEdgeSpacing instance.
func NewNSCollectionLayoutEdgeSpacing() NSCollectionLayoutEdgeSpacing {
	class := getNSCollectionLayoutEdgeSpacingClass()
	rv := objc.Send[NSCollectionLayoutEdgeSpacing](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an edge spacing object with the specified leading, top, trailing,
// and bottom spacing.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/init(leading:top:trailing:bottom:)
func NewCollectionLayoutEdgeSpacingForLeadingTopTrailingBottom(leading INSCollectionLayoutSpacing, top INSCollectionLayoutSpacing, trailing INSCollectionLayoutSpacing, bottom INSCollectionLayoutSpacing) NSCollectionLayoutEdgeSpacing {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutEdgeSpacingClass().class), objc.Sel("spacingForLeading:top:trailing:bottom:"), leading, top, trailing, bottom)
	return NSCollectionLayoutEdgeSpacingFromID(rv)
}

// The leading edge spacing value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/leading
func (c NSCollectionLayoutEdgeSpacing) Leading() INSCollectionLayoutSpacing {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("leading"))
	return NSCollectionLayoutSpacingFromID(objc.ID(rv))
}

// The top edge spacing value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/top
func (c NSCollectionLayoutEdgeSpacing) Top() INSCollectionLayoutSpacing {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("top"))
	return NSCollectionLayoutSpacingFromID(objc.ID(rv))
}

// The trailing edge spacing value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/trailing
func (c NSCollectionLayoutEdgeSpacing) Trailing() INSCollectionLayoutSpacing {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("trailing"))
	return NSCollectionLayoutSpacingFromID(objc.ID(rv))
}

// The bottom edge spacing value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/bottom
func (c NSCollectionLayoutEdgeSpacing) Bottom() INSCollectionLayoutSpacing {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("bottom"))
	return NSCollectionLayoutSpacingFromID(objc.ID(rv))
}


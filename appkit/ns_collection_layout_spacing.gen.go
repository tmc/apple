// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutSpacing] class.
var (
	_NSCollectionLayoutSpacingClass     NSCollectionLayoutSpacingClass
	_NSCollectionLayoutSpacingClassOnce sync.Once
)

func getNSCollectionLayoutSpacingClass() NSCollectionLayoutSpacingClass {
	_NSCollectionLayoutSpacingClassOnce.Do(func() {
		_NSCollectionLayoutSpacingClass = NSCollectionLayoutSpacingClass{class: objc.GetClass("NSCollectionLayoutSpacing")}
	})
	return _NSCollectionLayoutSpacingClass
}

// GetNSCollectionLayoutSpacingClass returns the class object for NSCollectionLayoutSpacing.
func GetNSCollectionLayoutSpacingClass() NSCollectionLayoutSpacingClass {
	return getNSCollectionLayoutSpacingClass()
}

type NSCollectionLayoutSpacingClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutSpacingClass) Alloc() NSCollectionLayoutSpacing {
	rv := objc.Send[NSCollectionLayoutSpacing](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that defines the space between or around items in a collection
// view.
//
// # Overview
// 
// In a collection view layout, you use a spacing object to specify both the
// amount of space and the way in which it’s calculated.
// 
// You can express spacing using fixed or flexible spacing.
// 
// Use to provide an exact amount of space. For example, the following code
// creates exactly 200 points of space between the items in the group.
// 
// Use to provide a minimum amount of space that can grow as more space
// becomes available. For example, the following code creates at least 200
// points of space between the items in the group. As more space becomes
// available, items are respaced evenly in the additional space.
//
// # Getting the spacing value
//
//   - [NSCollectionLayoutSpacing.Spacing]: The floating-point value of the space.
//
// # Getting the spacing type
//
//   - [NSCollectionLayoutSpacing.IsFixedSpacing]: A Boolean value that indicates whether the space is fixed to a specific number of points.
//   - [NSCollectionLayoutSpacing.IsFlexibleSpacing]: A Boolean value that indicates whether the space is flexible.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing
type NSCollectionLayoutSpacing struct {
	objectivec.Object
}

// NSCollectionLayoutSpacingFromID constructs a [NSCollectionLayoutSpacing] from an objc.ID.
//
// An object that defines the space between or around items in a collection
// view.
func NSCollectionLayoutSpacingFromID(id objc.ID) NSCollectionLayoutSpacing {
	return NSCollectionLayoutSpacing{objectivec.Object{ID: id}}
}
// NOTE: NSCollectionLayoutSpacing adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionLayoutSpacing] class.
//
// # Getting the spacing value
//
//   - [INSCollectionLayoutSpacing.Spacing]: The floating-point value of the space.
//
// # Getting the spacing type
//
//   - [INSCollectionLayoutSpacing.IsFixedSpacing]: A Boolean value that indicates whether the space is fixed to a specific number of points.
//   - [INSCollectionLayoutSpacing.IsFlexibleSpacing]: A Boolean value that indicates whether the space is flexible.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing
type INSCollectionLayoutSpacing interface {
	objectivec.IObject

	// Topic: Getting the spacing value

	// The floating-point value of the space.
	Spacing() float64

	// Topic: Getting the spacing type

	// A Boolean value that indicates whether the space is fixed to a specific number of points.
	IsFixedSpacing() bool
	// A Boolean value that indicates whether the space is flexible.
	IsFlexibleSpacing() bool
}

// Init initializes the instance.
func (c NSCollectionLayoutSpacing) Init() NSCollectionLayoutSpacing {
	rv := objc.Send[NSCollectionLayoutSpacing](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutSpacing) Autorelease() NSCollectionLayoutSpacing {
	rv := objc.Send[NSCollectionLayoutSpacing](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutSpacing creates a new NSCollectionLayoutSpacing instance.
func NewNSCollectionLayoutSpacing() NSCollectionLayoutSpacing {
	class := getNSCollectionLayoutSpacingClass()
	rv := objc.Send[NSCollectionLayoutSpacing](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a space equivalent to the specified number of points.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/fixed(_:)
func (_NSCollectionLayoutSpacingClass NSCollectionLayoutSpacingClass) FixedSpacing(fixedSpacing float64) NSCollectionLayoutSpacing {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutSpacingClass.class), objc.Sel("fixedSpacing:"), fixedSpacing)
	return NSCollectionLayoutSpacingFromID(rv)
}

// Creates a space equivalent to or greater than the specified number of
// points, depending on the available space.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/flexible(_:)
func (_NSCollectionLayoutSpacingClass NSCollectionLayoutSpacingClass) FlexibleSpacing(flexibleSpacing float64) NSCollectionLayoutSpacing {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutSpacingClass.class), objc.Sel("flexibleSpacing:"), flexibleSpacing)
	return NSCollectionLayoutSpacingFromID(rv)
}

// The floating-point value of the space.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/spacing
func (c NSCollectionLayoutSpacing) Spacing() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("spacing"))
	return rv
}

// A Boolean value that indicates whether the space is fixed to a specific
// number of points.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/isFixed
func (c NSCollectionLayoutSpacing) IsFixedSpacing() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFixedSpacing"))
	return rv
}

// A Boolean value that indicates whether the space is flexible.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/isFlexible
func (c NSCollectionLayoutSpacing) IsFlexibleSpacing() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFlexibleSpacing"))
	return rv
}


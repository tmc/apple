// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutDimension] class.
var (
	_NSCollectionLayoutDimensionClass     NSCollectionLayoutDimensionClass
	_NSCollectionLayoutDimensionClassOnce sync.Once
)

func getNSCollectionLayoutDimensionClass() NSCollectionLayoutDimensionClass {
	_NSCollectionLayoutDimensionClassOnce.Do(func() {
		_NSCollectionLayoutDimensionClass = NSCollectionLayoutDimensionClass{class: objc.GetClass("NSCollectionLayoutDimension")}
	})
	return _NSCollectionLayoutDimensionClass
}

// GetNSCollectionLayoutDimensionClass returns the class object for NSCollectionLayoutDimension.
func GetNSCollectionLayoutDimensionClass() NSCollectionLayoutDimensionClass {
	return getNSCollectionLayoutDimensionClass()
}

type NSCollectionLayoutDimensionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCollectionLayoutDimensionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutDimensionClass) Alloc() NSCollectionLayoutDimension {
	rv := objc.Send[NSCollectionLayoutDimension](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An individual dimension representing an item’s width or height in a
// collection view.
//
// # Overview
// 
// Each item in a collection view has an explicit width dimension and height
// dimension, which combine to define the item’s size
// ([NSCollectionLayoutSize]).
// 
// You can express an item’s dimensions using an absolute, estimated, or
// fractional value.
// 
// Use an to specify exact dimensions, like a 44 x 44 point square:
// 
// Use an if the size of your content might change at runtime, such as when
// data is loaded or in response to a change in system font size. You provide
// an initial estimated size and the system computes the actual value later.
// 
// Use a to define a value that’s relative to a dimension of the item’s
// container. This option simplifies specifying aspect ratios. For example,
// the following item has a width and a height that are both equal to 20% of
// its container’s width, creating a square that grows and shrinks as the
// size of its container changes.
//
// # Getting the dimension value
//
//   - [NSCollectionLayoutDimension.Dimension]: The floating-point value of the dimension.
//
// # Getting the dimension type
//
//   - [NSCollectionLayoutDimension.IsAbsolute]: A Boolean value that indicates whether the dimension is expressed as an absolute value.
//   - [NSCollectionLayoutDimension.IsEstimated]: A Boolean value that indicates whether the dimension is expressed as an estimated value.
//   - [NSCollectionLayoutDimension.IsFractionalHeight]: A Boolean value that indicates whether the dimension is expressed as a fraction of its container’s height.
//   - [NSCollectionLayoutDimension.IsFractionalWidth]: A Boolean value that indicates whether the dimension is expressed as a fraction of its container’s width.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension
type NSCollectionLayoutDimension struct {
	objectivec.Object
}

// NSCollectionLayoutDimensionFromID constructs a [NSCollectionLayoutDimension] from an objc.ID.
//
// An individual dimension representing an item’s width or height in a
// collection view.
func NSCollectionLayoutDimensionFromID(id objc.ID) NSCollectionLayoutDimension {
	return NSCollectionLayoutDimension{objectivec.Object{ID: id}}
}
// NOTE: NSCollectionLayoutDimension adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionLayoutDimension] class.
//
// # Getting the dimension value
//
//   - [INSCollectionLayoutDimension.Dimension]: The floating-point value of the dimension.
//
// # Getting the dimension type
//
//   - [INSCollectionLayoutDimension.IsAbsolute]: A Boolean value that indicates whether the dimension is expressed as an absolute value.
//   - [INSCollectionLayoutDimension.IsEstimated]: A Boolean value that indicates whether the dimension is expressed as an estimated value.
//   - [INSCollectionLayoutDimension.IsFractionalHeight]: A Boolean value that indicates whether the dimension is expressed as a fraction of its container’s height.
//   - [INSCollectionLayoutDimension.IsFractionalWidth]: A Boolean value that indicates whether the dimension is expressed as a fraction of its container’s width.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension
type INSCollectionLayoutDimension interface {
	objectivec.IObject

	// Topic: Getting the dimension value

	// The floating-point value of the dimension.
	Dimension() float64

	// Topic: Getting the dimension type

	// A Boolean value that indicates whether the dimension is expressed as an absolute value.
	IsAbsolute() bool
	// A Boolean value that indicates whether the dimension is expressed as an estimated value.
	IsEstimated() bool
	// A Boolean value that indicates whether the dimension is expressed as a fraction of its container’s height.
	IsFractionalHeight() bool
	// A Boolean value that indicates whether the dimension is expressed as a fraction of its container’s width.
	IsFractionalWidth() bool
}

// Init initializes the instance.
func (c NSCollectionLayoutDimension) Init() NSCollectionLayoutDimension {
	rv := objc.Send[NSCollectionLayoutDimension](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutDimension) Autorelease() NSCollectionLayoutDimension {
	rv := objc.Send[NSCollectionLayoutDimension](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutDimension creates a new NSCollectionLayoutDimension instance.
func NewNSCollectionLayoutDimension() NSCollectionLayoutDimension {
	class := getNSCollectionLayoutDimensionClass()
	rv := objc.Send[NSCollectionLayoutDimension](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a dimension with an absolute point value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/absolute(_:)
func (_NSCollectionLayoutDimensionClass NSCollectionLayoutDimensionClass) AbsoluteDimension(absoluteDimension float64) NSCollectionLayoutDimension {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutDimensionClass.class), objc.Sel("absoluteDimension:"), absoluteDimension)
	return NSCollectionLayoutDimensionFromID(rv)
}
// Creates a dimension with an estimated point value.
//
// # Discussion
// 
// The final size of the dimension is determined when the content is rendered.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/estimated(_:)
func (_NSCollectionLayoutDimensionClass NSCollectionLayoutDimensionClass) EstimatedDimension(estimatedDimension float64) NSCollectionLayoutDimension {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutDimensionClass.class), objc.Sel("estimatedDimension:"), estimatedDimension)
	return NSCollectionLayoutDimensionFromID(rv)
}
// Creates a dimension that is computed as a fraction of the height of the
// containing group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/fractionalHeight(_:)
func (_NSCollectionLayoutDimensionClass NSCollectionLayoutDimensionClass) FractionalHeightDimension(fractionalHeight float64) NSCollectionLayoutDimension {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutDimensionClass.class), objc.Sel("fractionalHeightDimension:"), fractionalHeight)
	return NSCollectionLayoutDimensionFromID(rv)
}
// Creates a dimension that is computed as a fraction of the width of the
// containing group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/fractionalWidth(_:)
func (_NSCollectionLayoutDimensionClass NSCollectionLayoutDimensionClass) FractionalWidthDimension(fractionalWidth float64) NSCollectionLayoutDimension {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionLayoutDimensionClass.class), objc.Sel("fractionalWidthDimension:"), fractionalWidth)
	return NSCollectionLayoutDimensionFromID(rv)
}

// The floating-point value of the dimension.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/dimension
func (c NSCollectionLayoutDimension) Dimension() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("dimension"))
	return rv
}
// A Boolean value that indicates whether the dimension is expressed as an
// absolute value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/isAbsolute
func (c NSCollectionLayoutDimension) IsAbsolute() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAbsolute"))
	return rv
}
// A Boolean value that indicates whether the dimension is expressed as an
// estimated value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/isEstimated
func (c NSCollectionLayoutDimension) IsEstimated() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEstimated"))
	return rv
}
// A Boolean value that indicates whether the dimension is expressed as a
// fraction of its container’s height.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/isFractionalHeight
func (c NSCollectionLayoutDimension) IsFractionalHeight() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFractionalHeight"))
	return rv
}
// A Boolean value that indicates whether the dimension is expressed as a
// fraction of its container’s width.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/isFractionalWidth
func (c NSCollectionLayoutDimension) IsFractionalWidth() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFractionalWidth"))
	return rv
}


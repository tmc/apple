// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutSize] class.
var (
	_NSCollectionLayoutSizeClass     NSCollectionLayoutSizeClass
	_NSCollectionLayoutSizeClassOnce sync.Once
)

func getNSCollectionLayoutSizeClass() NSCollectionLayoutSizeClass {
	_NSCollectionLayoutSizeClassOnce.Do(func() {
		_NSCollectionLayoutSizeClass = NSCollectionLayoutSizeClass{class: objc.GetClass("NSCollectionLayoutSize")}
	})
	return _NSCollectionLayoutSizeClass
}

// GetNSCollectionLayoutSizeClass returns the class object for NSCollectionLayoutSize.
func GetNSCollectionLayoutSizeClass() NSCollectionLayoutSizeClass {
	return getNSCollectionLayoutSizeClass()
}

type NSCollectionLayoutSizeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCollectionLayoutSizeClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutSizeClass) Alloc() NSCollectionLayoutSize {
	rv := objc.Send[NSCollectionLayoutSize](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The width and the height of an item in a collection view.
//
// # Overview
// 
// A size is a pair of dimensions ([NSCollectionLayoutDimension]): a width
// dimension and a height dimension. Every component of a collection view
// layout has an explicit size.
//
// # Getting the width and height
//
//   - [NSCollectionLayoutSize.WidthDimension]: The width dimension of an item in a collection view layout.
//   - [NSCollectionLayoutSize.HeightDimension]: The height dimension of an item in a collection view layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize
type NSCollectionLayoutSize struct {
	objectivec.Object
}

// NSCollectionLayoutSizeFromID constructs a [NSCollectionLayoutSize] from an objc.ID.
//
// The width and the height of an item in a collection view.
func NSCollectionLayoutSizeFromID(id objc.ID) NSCollectionLayoutSize {
	return NSCollectionLayoutSize{objectivec.Object{ID: id}}
}
// NOTE: NSCollectionLayoutSize adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionLayoutSize] class.
//
// # Getting the width and height
//
//   - [INSCollectionLayoutSize.WidthDimension]: The width dimension of an item in a collection view layout.
//   - [INSCollectionLayoutSize.HeightDimension]: The height dimension of an item in a collection view layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize
type INSCollectionLayoutSize interface {
	objectivec.IObject

	// Topic: Getting the width and height

	// The width dimension of an item in a collection view layout.
	WidthDimension() INSCollectionLayoutDimension
	// The height dimension of an item in a collection view layout.
	HeightDimension() INSCollectionLayoutDimension
}

// Init initializes the instance.
func (c NSCollectionLayoutSize) Init() NSCollectionLayoutSize {
	rv := objc.Send[NSCollectionLayoutSize](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutSize) Autorelease() NSCollectionLayoutSize {
	rv := objc.Send[NSCollectionLayoutSize](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutSize creates a new NSCollectionLayoutSize instance.
func NewNSCollectionLayoutSize() NSCollectionLayoutSize {
	class := getNSCollectionLayoutSizeClass()
	rv := objc.Send[NSCollectionLayoutSize](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a size with the specified width and height dimensions.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/init(widthDimension:heightDimension:)
func NewCollectionLayoutSizeWithWidthDimensionHeightDimension(width INSCollectionLayoutDimension, height INSCollectionLayoutDimension) NSCollectionLayoutSize {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutSizeClass().class), objc.Sel("sizeWithWidthDimension:heightDimension:"), width, height)
	return NSCollectionLayoutSizeFromID(rv)
}

// The width dimension of an item in a collection view layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/widthDimension
func (c NSCollectionLayoutSize) WidthDimension() INSCollectionLayoutDimension {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("widthDimension"))
	return NSCollectionLayoutDimensionFromID(objc.ID(rv))
}
// The height dimension of an item in a collection view layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/heightDimension
func (c NSCollectionLayoutSize) HeightDimension() INSCollectionLayoutDimension {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("heightDimension"))
	return NSCollectionLayoutDimensionFromID(objc.ID(rv))
}


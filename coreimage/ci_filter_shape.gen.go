// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIFilterShape] class.
var (
	_CIFilterShapeClass     CIFilterShapeClass
	_CIFilterShapeClassOnce sync.Once
)

func getCIFilterShapeClass() CIFilterShapeClass {
	_CIFilterShapeClassOnce.Do(func() {
		_CIFilterShapeClass = CIFilterShapeClass{class: objc.GetClass("CIFilterShape")}
	})
	return _CIFilterShapeClass
}

// GetCIFilterShapeClass returns the class object for CIFilterShape.
func GetCIFilterShapeClass() CIFilterShapeClass {
	return getCIFilterShapeClass()
}

type CIFilterShapeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIFilterShapeClass) Alloc() CIFilterShape {
	rv := objc.Send[CIFilterShape](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A description of the bounding shape of a filter and the domain of
// definition for a filter operation.
//
// # Overview
// 
// You use [CIFilterShape] objects in conjunction with Core Image classes,
// such as [CIFilter], [CIKernel], and [CISampler], to create custom filters.
//
// # Initializing a Filter Shape
//
//   - [CIFilterShape.InitWithRect]: Initializes a filter shape object with a rectangle.
//
// # Inspecting a Filter Shape
//
//   - [CIFilterShape.Extent]: The extent of the filter shape.
//
// # Modifying a Filter Shape
//
//   - [CIFilterShape.InsetByXY]: Modifies a filter shape object so that it is inset by the specified x and y values.
//   - [CIFilterShape.IntersectWith]: Creates a filter shape object that represents the intersection of the current filter shape and the specified filter shape object.
//   - [CIFilterShape.IntersectWithRect]: Creates a filter shape that represents the intersection of the current filter shape and a rectangle.
//   - [CIFilterShape.TransformByInterior]: Creates a filter shape that results from applying a transform to the current filter shape.
//   - [CIFilterShape.UnionWith]: Creates a filter shape that results from the union of the current filter shape and another filter shape object.
//   - [CIFilterShape.UnionWithRect]: Creates a filter shape that results from the union of the current filter shape and a rectangle.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape
type CIFilterShape struct {
	objectivec.Object
}

// CIFilterShapeFromID constructs a [CIFilterShape] from an objc.ID.
//
// A description of the bounding shape of a filter and the domain of
// definition for a filter operation.
func CIFilterShapeFromID(id objc.ID) CIFilterShape {
	return CIFilterShape{objectivec.Object{ID: id}}
}
// NOTE: CIFilterShape adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIFilterShape] class.
//
// # Initializing a Filter Shape
//
//   - [ICIFilterShape.InitWithRect]: Initializes a filter shape object with a rectangle.
//
// # Inspecting a Filter Shape
//
//   - [ICIFilterShape.Extent]: The extent of the filter shape.
//
// # Modifying a Filter Shape
//
//   - [ICIFilterShape.InsetByXY]: Modifies a filter shape object so that it is inset by the specified x and y values.
//   - [ICIFilterShape.IntersectWith]: Creates a filter shape object that represents the intersection of the current filter shape and the specified filter shape object.
//   - [ICIFilterShape.IntersectWithRect]: Creates a filter shape that represents the intersection of the current filter shape and a rectangle.
//   - [ICIFilterShape.TransformByInterior]: Creates a filter shape that results from applying a transform to the current filter shape.
//   - [ICIFilterShape.UnionWith]: Creates a filter shape that results from the union of the current filter shape and another filter shape object.
//   - [ICIFilterShape.UnionWithRect]: Creates a filter shape that results from the union of the current filter shape and a rectangle.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape
type ICIFilterShape interface {
	objectivec.IObject

	// Topic: Initializing a Filter Shape

	// Initializes a filter shape object with a rectangle.
	InitWithRect(r corefoundation.CGRect) CIFilterShape

	// Topic: Inspecting a Filter Shape

	// The extent of the filter shape.
	Extent() corefoundation.CGRect

	// Topic: Modifying a Filter Shape

	// Modifies a filter shape object so that it is inset by the specified x and y values.
	InsetByXY(dx int, dy int) ICIFilterShape
	// Creates a filter shape object that represents the intersection of the current filter shape and the specified filter shape object.
	IntersectWith(s2 ICIFilterShape) ICIFilterShape
	// Creates a filter shape that represents the intersection of the current filter shape and a rectangle.
	IntersectWithRect(r corefoundation.CGRect) ICIFilterShape
	// Creates a filter shape that results from applying a transform to the current filter shape.
	TransformByInterior(m corefoundation.CGAffineTransform, flag bool) ICIFilterShape
	// Creates a filter shape that results from the union of the current filter shape and another filter shape object.
	UnionWith(s2 ICIFilterShape) ICIFilterShape
	// Creates a filter shape that results from the union of the current filter shape and a rectangle.
	UnionWithRect(r corefoundation.CGRect) ICIFilterShape
}

// Init initializes the instance.
func (f CIFilterShape) Init() CIFilterShape {
	rv := objc.Send[CIFilterShape](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f CIFilterShape) Autorelease() CIFilterShape {
	rv := objc.Send[CIFilterShape](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIFilterShape creates a new CIFilterShape instance.
func NewCIFilterShape() CIFilterShape {
	class := getCIFilterShapeClass()
	rv := objc.Send[CIFilterShape](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a filter shape object with a rectangle.
//
// r: A rectangle. Core Image uses the rectangle specified by integer parts of
// the values in the [CGRect] data structure.
//
// # Return Value
// 
// An initialized CIFilterShape object, or `nil` if the method fails.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape/init(rect:)
func NewFilterShapeWithRect(r corefoundation.CGRect) CIFilterShape {
	instance := getCIFilterShapeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRect:"), r)
	return CIFilterShapeFromID(rv)
}

// Initializes a filter shape object with a rectangle.
//
// r: A rectangle. Core Image uses the rectangle specified by integer parts of
// the values in the [CGRect] data structure.
//
// # Return Value
// 
// An initialized CIFilterShape object, or `nil` if the method fails.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape/init(rect:)
func (f CIFilterShape) InitWithRect(r corefoundation.CGRect) CIFilterShape {
	rv := objc.Send[CIFilterShape](f.ID, objc.Sel("initWithRect:"), r)
	return rv
}
// Modifies a filter shape object so that it is inset by the specified x and y
// values.
//
// dx: A value that specifies an inset in the x direction.
//
// dy: A value that specifies an inset in the y direction.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape/insetBy(x:y:)
func (f CIFilterShape) InsetByXY(dx int, dy int) ICIFilterShape {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("insetByX:Y:"), dx, dy)
	return CIFilterShapeFromID(rv)
}
// Creates a filter shape object that represents the intersection of the
// current filter shape and the specified filter shape object.
//
// s2: A filter shape object.
//
// # Return Value
// 
// The filter shape object that results from the intersection.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape/intersect(with:)-8iw
func (f CIFilterShape) IntersectWith(s2 ICIFilterShape) ICIFilterShape {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("intersectWith:"), s2)
	return CIFilterShapeFromID(rv)
}
// Creates a filter shape that represents the intersection of the current
// filter shape and a rectangle.
//
// r: A rectangle. Core Image uses the rectangle specified by integer parts of
// the width and height.
//
// # Return Value
// 
// The filter shape that results from the intersection
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape/intersect(with:)-2o2n8
func (f CIFilterShape) IntersectWithRect(r corefoundation.CGRect) ICIFilterShape {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("intersectWithRect:"), r)
	return CIFilterShapeFromID(rv)
}
// Creates a filter shape that results from applying a transform to the
// current filter shape.
//
// m: A transform.
//
// flag: [false] specifies that the new filter shape object can contain all the
// pixels in the transformed shape (and possibly some that are outside the
// transformed shape). [true] specifies that the new filter shape object can
// contain a subset of the pixels in the transformed shape (but none of those
// outside the transformed shape).
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The transformed filter shape object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape/transform(by:interior:)
func (f CIFilterShape) TransformByInterior(m corefoundation.CGAffineTransform, flag bool) ICIFilterShape {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("transformBy:interior:"), m, flag)
	return CIFilterShapeFromID(rv)
}
// Creates a filter shape that results from the union of the current filter
// shape and another filter shape object.
//
// s2: A filter shape object.
//
// # Return Value
// 
// The filter shape object that results from the union.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape/union(with:)-52mnd
func (f CIFilterShape) UnionWith(s2 ICIFilterShape) ICIFilterShape {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("unionWith:"), s2)
	return CIFilterShapeFromID(rv)
}
// Creates a filter shape that results from the union of the current filter
// shape and a rectangle.
//
// r: A rectangle. Core Image uses the rectangle specified by integer parts of
// the width and height.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape/union(with:)-75ebo
func (f CIFilterShape) UnionWithRect(r corefoundation.CGRect) ICIFilterShape {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("unionWithRect:"), r)
	return CIFilterShapeFromID(rv)
}

// Creates a filter shape object and initializes it with a rectangle.
//
// r: A rectangle. The filter shape object will contain the smallest integral
// rectangle specified by this argument.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape/shapeWithRect:
func (_CIFilterShapeClass CIFilterShapeClass) ShapeWithRect(r corefoundation.CGRect) CIFilterShape {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterShapeClass.class), objc.Sel("shapeWithRect:"), r)
	return CIFilterShapeFromID(rv)
}

// The extent of the filter shape.
//
// # Discussion
// 
// Extent is a rectangle that describes the filter shape in the working
// coordinate space with a fixed area.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterShape/extent
func (f CIFilterShape) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](f.ID, objc.Sel("extent"))
	return corefoundation.CGRect(rv)
}


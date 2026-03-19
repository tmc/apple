// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNPoint] class.
var (
	_VNPointClass     VNPointClass
	_VNPointClassOnce sync.Once
)

func getVNPointClass() VNPointClass {
	_VNPointClassOnce.Do(func() {
		_VNPointClass = VNPointClass{class: objc.GetClass("VNPoint")}
	})
	return _VNPointClass
}

// GetVNPointClass returns the class object for VNPoint.
func GetVNPointClass() VNPointClass {
	return getVNPointClass()
}

type VNPointClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNPointClass) Alloc() VNPoint {
	rv := objc.Send[VNPoint](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An immutable object that represents a single 2D point in an image.
//
// # Creating a Point
//
//   - [VNPoint.InitWithXY]: Creates a point object with the specified coordinates.
//   - [VNPoint.InitWithLocation]: Creates a point object from the specified Core Graphics point.
//
// # Inspecting a Point
//
//   - [VNPoint.X]: The x-coordinate.
//   - [VNPoint.Y]: The y-coordinate.
//   - [VNPoint.Location]: The Core Graphics point for this point.
//
// # Calculating Distance
//
//   - [VNPoint.DistanceToPoint]: Returns the distance to another point.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint
type VNPoint struct {
	objectivec.Object
}

// VNPointFromID constructs a [VNPoint] from an objc.ID.
//
// An immutable object that represents a single 2D point in an image.
func VNPointFromID(id objc.ID) VNPoint {
	return VNPoint{objectivec.Object{ID: id}}
}
// NOTE: VNPoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNPoint] class.
//
// # Creating a Point
//
//   - [IVNPoint.InitWithXY]: Creates a point object with the specified coordinates.
//   - [IVNPoint.InitWithLocation]: Creates a point object from the specified Core Graphics point.
//
// # Inspecting a Point
//
//   - [IVNPoint.X]: The x-coordinate.
//   - [IVNPoint.Y]: The y-coordinate.
//   - [IVNPoint.Location]: The Core Graphics point for this point.
//
// # Calculating Distance
//
//   - [IVNPoint.DistanceToPoint]: Returns the distance to another point.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint
type IVNPoint interface {
	objectivec.IObject

	// Topic: Creating a Point

	// Creates a point object with the specified coordinates.
	InitWithXY(x float64, y float64) VNPoint
	// Creates a point object from the specified Core Graphics point.
	InitWithLocation(location corefoundation.CGPoint) VNPoint

	// Topic: Inspecting a Point

	// The x-coordinate.
	X() float64
	// The y-coordinate.
	Y() float64
	// The Core Graphics point for this point.
	Location() corefoundation.CGPoint

	// Topic: Calculating Distance

	// Returns the distance to another point.
	DistanceToPoint(point IVNPoint) float64

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p VNPoint) Init() VNPoint {
	rv := objc.Send[VNPoint](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p VNPoint) Autorelease() VNPoint {
	rv := objc.Send[VNPoint](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNPoint creates a new VNPoint instance.
func NewVNPoint() VNPoint {
	class := getVNPointClass()
	rv := objc.Send[VNPoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a point object from the specified Core Graphics point.
//
// location: The Core Graphics point.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/init(location:)
func NewPointWithLocation(location corefoundation.CGPoint) VNPoint {
	instance := getVNPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocation:"), location)
	return VNPointFromID(rv)
}

// Creates a point object with the specified coordinates.
//
// x: The x-coordinate value.
//
// y: The y-coordinate value.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/init(x:y:)
func NewPointWithXY(x float64, y float64) VNPoint {
	instance := getVNPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithX:y:"), x, y)
	return VNPointFromID(rv)
}

// Creates a point object with the specified coordinates.
//
// x: The x-coordinate value.
//
// y: The y-coordinate value.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/init(x:y:)
func (p VNPoint) InitWithXY(x float64, y float64) VNPoint {
	rv := objc.Send[VNPoint](p.ID, objc.Sel("initWithX:y:"), x, y)
	return rv
}
// Creates a point object from the specified Core Graphics point.
//
// location: The Core Graphics point.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/init(location:)
func (p VNPoint) InitWithLocation(location corefoundation.CGPoint) VNPoint {
	rv := objc.Send[VNPoint](p.ID, objc.Sel("initWithLocation:"), location)
	return rv
}
// Returns the distance to another point.
//
// point: The point for which to calculate the distance.
//
// # Return Value
// 
// The calculated distance.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/distance(_:)
func (p VNPoint) DistanceToPoint(point IVNPoint) float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("distanceToPoint:"), point)
	return rv
}
func (p VNPoint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates a point object that’s shifted by the X and Y offsets of the
// specified vector.
//
// vector: The vector to apply to offset the point.
//
// point: The point to translate by the vector’s X and Y offsets.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/apply(_:to:)
func (_VNPointClass VNPointClass) PointByApplyingVectorToPoint(vector IVNVector, point IVNPoint) VNPoint {
	rv := objc.Send[objc.ID](objc.ID(_VNPointClass.class), objc.Sel("pointByApplyingVector:toPoint:"), vector, point)
	return VNPointFromID(rv)
}

// The x-coordinate.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/x
func (p VNPoint) X() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("x"))
	return rv
}
// The y-coordinate.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/y
func (p VNPoint) Y() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("y"))
	return rv
}
// The Core Graphics point for this point.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/location
func (p VNPoint) Location() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("location"))
	return corefoundation.CGPoint(rv)
}

// A point object that represents the origin.
//
// # Discussion
// 
// The origin point is (0.0, 0.0).
//
// See: https://developer.apple.com/documentation/Vision/VNPoint/zero
func (_VNPointClass VNPointClass) ZeroPoint() VNPoint {
	rv := objc.Send[objc.ID](objc.ID(_VNPointClass.class), objc.Sel("zeroPoint"))
	return VNPointFromID(objc.ID(rv))
}


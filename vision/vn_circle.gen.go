// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNCircle] class.
var (
	_VNCircleClass     VNCircleClass
	_VNCircleClassOnce sync.Once
)

func getVNCircleClass() VNCircleClass {
	_VNCircleClassOnce.Do(func() {
		_VNCircleClass = VNCircleClass{class: objc.GetClass("VNCircle")}
	})
	return _VNCircleClass
}

// GetVNCircleClass returns the class object for VNCircle.
func GetVNCircleClass() VNCircleClass {
	return getVNCircleClass()
}

type VNCircleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNCircleClass) Alloc() VNCircle {
	rv := objc.Send[VNCircle](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An immutable 2D circle represented by its center point and radius.
//
// # Creating a Circle
//
//   - [VNCircle.InitWithCenterRadius]: Creates a circle with the specified center and radius.
//   - [VNCircle.InitWithCenterDiameter]: Creates a circle with the specified center and diameter.
//
// # Inspecting a Circle
//
//   - [VNCircle.Center]: The circle’s center point.
//   - [VNCircle.Diameter]: The circle’s diameter.
//   - [VNCircle.Radius]: The circle’s radius.
//   - [VNCircle.ContainsPoint]: Determines if this circle, including its boundary, contains the specified point.
//   - [VNCircle.ContainsPointInCircumferentialRingOfWidth]: Determines if a ring around this circle’s circumference contains the specified point.
//
// See: https://developer.apple.com/documentation/Vision/VNCircle
type VNCircle struct {
	objectivec.Object
}

// VNCircleFromID constructs a [VNCircle] from an objc.ID.
//
// An immutable 2D circle represented by its center point and radius.
func VNCircleFromID(id objc.ID) VNCircle {
	return VNCircle{objectivec.Object{id}}
}
// NOTE: VNCircle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNCircle] class.
//
// # Creating a Circle
//
//   - [IVNCircle.InitWithCenterRadius]: Creates a circle with the specified center and radius.
//   - [IVNCircle.InitWithCenterDiameter]: Creates a circle with the specified center and diameter.
//
// # Inspecting a Circle
//
//   - [IVNCircle.Center]: The circle’s center point.
//   - [IVNCircle.Diameter]: The circle’s diameter.
//   - [IVNCircle.Radius]: The circle’s radius.
//   - [IVNCircle.ContainsPoint]: Determines if this circle, including its boundary, contains the specified point.
//   - [IVNCircle.ContainsPointInCircumferentialRingOfWidth]: Determines if a ring around this circle’s circumference contains the specified point.
//
// See: https://developer.apple.com/documentation/Vision/VNCircle
type IVNCircle interface {
	objectivec.IObject

	// Topic: Creating a Circle

	// Creates a circle with the specified center and radius.
	InitWithCenterRadius(center IVNPoint, radius float64) VNCircle
	// Creates a circle with the specified center and diameter.
	InitWithCenterDiameter(center IVNPoint, diameter float64) VNCircle

	// Topic: Inspecting a Circle

	// The circle’s center point.
	Center() IVNPoint
	// The circle’s diameter.
	Diameter() float64
	// The circle’s radius.
	Radius() float64
	// Determines if this circle, including its boundary, contains the specified point.
	ContainsPoint(point IVNPoint) bool
	// Determines if a ring around this circle’s circumference contains the specified point.
	ContainsPointInCircumferentialRingOfWidth(point IVNPoint, ringWidth float64) bool

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (c VNCircle) Init() VNCircle {
	rv := objc.Send[VNCircle](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VNCircle) Autorelease() VNCircle {
	rv := objc.Send[VNCircle](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNCircle creates a new VNCircle instance.
func NewVNCircle() VNCircle {
	class := getVNCircleClass()
	rv := objc.Send[VNCircle](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a circle with the specified center and diameter.
//
// center: The circle center.
//
// diameter: The circle diameter.
//
// See: https://developer.apple.com/documentation/Vision/VNCircle/init(center:diameter:)
func NewCircleWithCenterDiameter(center IVNPoint, diameter float64) VNCircle {
	instance := getVNCircleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCenter:diameter:"), center, diameter)
	return VNCircleFromID(rv)
}


// Creates a circle with the specified center and radius.
//
// center: The circle center.
//
// radius: The circle radius.
//
// See: https://developer.apple.com/documentation/Vision/VNCircle/init(center:radius:)
func NewCircleWithCenterRadius(center IVNPoint, radius float64) VNCircle {
	instance := getVNCircleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCenter:radius:"), center, radius)
	return VNCircleFromID(rv)
}







// Creates a circle with the specified center and radius.
//
// center: The circle center.
//
// radius: The circle radius.
//
// See: https://developer.apple.com/documentation/Vision/VNCircle/init(center:radius:)
func (c VNCircle) InitWithCenterRadius(center IVNPoint, radius float64) VNCircle {
	rv := objc.Send[VNCircle](c.ID, objc.Sel("initWithCenter:radius:"), center, radius)
	return rv
}

// Creates a circle with the specified center and diameter.
//
// center: The circle center.
//
// diameter: The circle diameter.
//
// See: https://developer.apple.com/documentation/Vision/VNCircle/init(center:diameter:)
func (c VNCircle) InitWithCenterDiameter(center IVNPoint, diameter float64) VNCircle {
	rv := objc.Send[VNCircle](c.ID, objc.Sel("initWithCenter:diameter:"), center, diameter)
	return rv
}

// Determines if this circle, including its boundary, contains the specified
// point.
//
// point: The point to test.
//
// # Return Value
// 
// [true] if the point is contained within this circle, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Foundation/NSExpression/true
//
// See: https://developer.apple.com/documentation/Vision/VNCircle/contains(_:)
func (c VNCircle) ContainsPoint(point IVNPoint) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("containsPoint:"), point)
	return rv
}

// Determines if a ring around this circle’s circumference contains the
// specified point.
//
// point: The point to test.
//
// ringWidth: The width of the ring around this circle’s circumference.
//
// # Return Value
// 
// [true] if the ring contains the specified point, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Vision/VNCircle/contains(_:inCircumferentialRingOfWidth:)
func (c VNCircle) ContainsPointInCircumferentialRingOfWidth(point IVNPoint, ringWidth float64) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("containsPoint:inCircumferentialRingOfWidth:"), point, ringWidth)
	return rv
}
func (c VNCircle) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The circle’s center point.
//
// See: https://developer.apple.com/documentation/Vision/VNCircle/center
func (c VNCircle) Center() IVNPoint {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("center"))
	return VNPointFromID(objc.ID(rv))
}



// The circle’s diameter.
//
// See: https://developer.apple.com/documentation/Vision/VNCircle/diameter
func (c VNCircle) Diameter() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("diameter"))
	return rv
}



// The circle’s radius.
//
// See: https://developer.apple.com/documentation/Vision/VNCircle/radius
func (c VNCircle) Radius() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("radius"))
	return rv
}







// A circle object centered at the origin, with a radius of zero.
//
// See: https://developer.apple.com/documentation/Vision/VNCircle/zero
func (_VNCircleClass VNCircleClass) ZeroCircle() VNCircle {
	rv := objc.Send[objc.ID](objc.ID(_VNCircleClass.class), objc.Sel("zeroCircle"))
	return VNCircleFromID(objc.ID(rv))
}























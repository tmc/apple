// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNVector] class.
var (
	_VNVectorClass     VNVectorClass
	_VNVectorClassOnce sync.Once
)

func getVNVectorClass() VNVectorClass {
	_VNVectorClassOnce.Do(func() {
		_VNVectorClass = VNVectorClass{class: objc.GetClass("VNVector")}
	})
	return _VNVectorClass
}

// GetVNVectorClass returns the class object for VNVector.
func GetVNVectorClass() VNVectorClass {
	return getVNVectorClass()
}

type VNVectorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNVectorClass) Alloc() VNVector {
	rv := objc.Send[VNVector](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An immutable 2D vector represented by its x-axis and y-axis projections.
//
// # Creating a Vector
//
//   - [VNVector.InitWithRTheta]: Creates a new vector in polar coordinate space.
//   - [VNVector.InitWithVectorHeadTail]: Creates a new vector in Cartesian coordinate space.
//   - [VNVector.InitWithXComponentYComponent]: Creates a new vector in Cartesian coordinate space, based on its x-axis and y-axis projections.
//
// # Inspecting a Vector
//
//   - [VNVector.Length]: The length, or absolute value, of the vector.
//   - [VNVector.R]: The radius, absolute value, or length of the vector.
//   - [VNVector.Theta]: The angle between the vector direction and the positive direction of the x-axis.
//   - [VNVector.SquaredLength]: The squared length of the vector.
//   - [VNVector.X]: A signed projection that indicates the vector’s direction on the x-axis.
//   - [VNVector.Y]: A signed projection that indicates the vector’s direction on the y-axis.
//
// See: https://developer.apple.com/documentation/Vision/VNVector
type VNVector struct {
	objectivec.Object
}

// VNVectorFromID constructs a [VNVector] from an objc.ID.
//
// An immutable 2D vector represented by its x-axis and y-axis projections.
func VNVectorFromID(id objc.ID) VNVector {
	return VNVector{objectivec.Object{ID: id}}
}
// NOTE: VNVector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNVector] class.
//
// # Creating a Vector
//
//   - [IVNVector.InitWithRTheta]: Creates a new vector in polar coordinate space.
//   - [IVNVector.InitWithVectorHeadTail]: Creates a new vector in Cartesian coordinate space.
//   - [IVNVector.InitWithXComponentYComponent]: Creates a new vector in Cartesian coordinate space, based on its x-axis and y-axis projections.
//
// # Inspecting a Vector
//
//   - [IVNVector.Length]: The length, or absolute value, of the vector.
//   - [IVNVector.R]: The radius, absolute value, or length of the vector.
//   - [IVNVector.Theta]: The angle between the vector direction and the positive direction of the x-axis.
//   - [IVNVector.SquaredLength]: The squared length of the vector.
//   - [IVNVector.X]: A signed projection that indicates the vector’s direction on the x-axis.
//   - [IVNVector.Y]: A signed projection that indicates the vector’s direction on the y-axis.
//
// See: https://developer.apple.com/documentation/Vision/VNVector
type IVNVector interface {
	objectivec.IObject

	// Topic: Creating a Vector

	// Creates a new vector in polar coordinate space.
	InitWithRTheta(r float64, theta float64) VNVector
	// Creates a new vector in Cartesian coordinate space.
	InitWithVectorHeadTail(head IVNPoint, tail IVNPoint) VNVector
	// Creates a new vector in Cartesian coordinate space, based on its x-axis and y-axis projections.
	InitWithXComponentYComponent(x float64, y float64) VNVector

	// Topic: Inspecting a Vector

	// The length, or absolute value, of the vector.
	Length() float64
	// The radius, absolute value, or length of the vector.
	R() float64
	// The angle between the vector direction and the positive direction of the x-axis.
	Theta() float64
	// The squared length of the vector.
	SquaredLength() float64
	// A signed projection that indicates the vector’s direction on the x-axis.
	X() float64
	// A signed projection that indicates the vector’s direction on the y-axis.
	Y() float64

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (v VNVector) Init() VNVector {
	rv := objc.Send[VNVector](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VNVector) Autorelease() VNVector {
	rv := objc.Send[VNVector](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNVector creates a new VNVector instance.
func NewVNVector() VNVector {
	class := getVNVectorClass()
	rv := objc.Send[VNVector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new vector by adding the specified vectors.
//
// v1: The first vector.
//
// v2: The second vector.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/init(byAdding:to:)
func NewVectorByAddingVectorToVector(v1 IVNVector, v2 IVNVector) VNVector {
	rv := objc.Send[objc.ID](objc.ID(getVNVectorClass().class), objc.Sel("vectorByAddingVector:toVector:"), v1, v2)
	return VNVectorFromID(rv)
}

// Creates a new vector by multiplying the specified vector’s x-axis and
// y-axis projections by the scalar value.
//
// vector: The vector.
//
// scalar: The scalar value by which to multiply the x-axis and y-axis projections.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/init(byMultiplying:byScalar:)
func NewVectorByMultiplyingVectorByScalar(vector IVNVector, scalar float64) VNVector {
	rv := objc.Send[objc.ID](objc.ID(getVNVectorClass().class), objc.Sel("vectorByMultiplyingVector:byScalar:"), vector, scalar)
	return VNVectorFromID(rv)
}

// Creates a new vector by subtracting the first vector from the second
// vector.
//
// v1: The first vector.
//
// v2: The second vector.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/init(bySubtracting:from:)
func NewVectorBySubtractingVectorFromVector(v1 IVNVector, v2 IVNVector) VNVector {
	rv := objc.Send[objc.ID](objc.ID(getVNVectorClass().class), objc.Sel("vectorBySubtractingVector:fromVector:"), v1, v2)
	return VNVectorFromID(rv)
}

// Creates a new vector in polar coordinate space.
//
// r: The length of the vector.
//
// theta: The angle that the vector forms with the positive direction of the x-axis.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/init(r:theta:)
func NewVectorWithRTheta(r float64, theta float64) VNVector {
	instance := getVNVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithR:theta:"), r, theta)
	return VNVectorFromID(rv)
}

// Creates a new vector in Cartesian coordinate space.
//
// head: The vector’s head point.
//
// tail: The vector’s tail point.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/init(vectorHead:tail:)
func NewVectorWithVectorHeadTail(head IVNPoint, tail IVNPoint) VNVector {
	instance := getVNVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVectorHead:tail:"), head, tail)
	return VNVectorFromID(rv)
}

// Creates a new vector in Cartesian coordinate space, based on its x-axis and
// y-axis projections.
//
// x: The x-component.
//
// y: The y-component.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/init(xComponent:yComponent:)
func NewVectorWithXComponentYComponent(x float64, y float64) VNVector {
	instance := getVNVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXComponent:yComponent:"), x, y)
	return VNVectorFromID(rv)
}

// Creates a new vector in polar coordinate space.
//
// r: The length of the vector.
//
// theta: The angle that the vector forms with the positive direction of the x-axis.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/init(r:theta:)
func (v VNVector) InitWithRTheta(r float64, theta float64) VNVector {
	rv := objc.Send[VNVector](v.ID, objc.Sel("initWithR:theta:"), r, theta)
	return rv
}

// Creates a new vector in Cartesian coordinate space.
//
// head: The vector’s head point.
//
// tail: The vector’s tail point.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/init(vectorHead:tail:)
func (v VNVector) InitWithVectorHeadTail(head IVNPoint, tail IVNPoint) VNVector {
	rv := objc.Send[VNVector](v.ID, objc.Sel("initWithVectorHead:tail:"), head, tail)
	return rv
}

// Creates a new vector in Cartesian coordinate space, based on its x-axis and
// y-axis projections.
//
// x: The x-component.
//
// y: The y-component.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/init(xComponent:yComponent:)
func (v VNVector) InitWithXComponentYComponent(x float64, y float64) VNVector {
	rv := objc.Send[VNVector](v.ID, objc.Sel("initWithXComponent:yComponent:"), x, y)
	return rv
}
func (v VNVector) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Caclulates the dot product of two vectors.
//
// v1: The first vector.
//
// v2: The second vector.
//
// # Return Value
// 
// The dot product value.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/dotProduct(of:vector:)
func (_VNVectorClass VNVectorClass) DotProductOfVectorVector(v1 IVNVector, v2 IVNVector) float64 {
	rv := objc.Send[float64](objc.ID(_VNVectorClass.class), objc.Sel("dotProductOfVector:vector:"), v1, v2)
	return rv
}

// Calculates a vector that’s normalized by preserving its direction, so
// that the vector length equals 1.0.
//
// vector: The vector whose unit vector you want to calculate.
//
// # Return Value
// 
// The unit vector.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/unitVector(for:)
func (_VNVectorClass VNVectorClass) UnitVectorForVector(vector IVNVector) VNVector {
	rv := objc.Send[objc.ID](objc.ID(_VNVectorClass.class), objc.Sel("unitVectorForVector:"), vector)
	return VNVectorFromID(rv)
}

// The length, or absolute value, of the vector.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/length
func (v VNVector) Length() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("length"))
	return rv
}

// The radius, absolute value, or length of the vector.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/r
func (v VNVector) R() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("r"))
	return rv
}

// The angle between the vector direction and the positive direction of the
// x-axis.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/theta
func (v VNVector) Theta() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("theta"))
	return rv
}

// The squared length of the vector.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/squaredLength
func (v VNVector) SquaredLength() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("squaredLength"))
	return rv
}

// A signed projection that indicates the vector’s direction on the x-axis.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/x
func (v VNVector) X() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("x"))
	return rv
}

// A signed projection that indicates the vector’s direction on the y-axis.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/y
func (v VNVector) Y() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("y"))
	return rv
}

// A vector object with zero length.
//
// See: https://developer.apple.com/documentation/Vision/VNVector/zero
func (_VNVectorClass VNVectorClass) ZeroVector() VNVector {
	rv := objc.Send[objc.ID](objc.ID(_VNVectorClass.class), objc.Sel("zeroVector"))
	return VNVectorFromID(objc.ID(rv))
}


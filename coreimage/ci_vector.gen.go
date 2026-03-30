// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIVector] class.
var (
	_CIVectorClass     CIVectorClass
	_CIVectorClassOnce sync.Once
)

func getCIVectorClass() CIVectorClass {
	_CIVectorClassOnce.Do(func() {
		_CIVectorClass = CIVectorClass{class: objc.GetClass("CIVector")}
	})
	return _CIVectorClass
}

// GetCIVectorClass returns the class object for CIVector.
func GetCIVectorClass() CIVectorClass {
	return getCIVectorClass()
}

type CIVectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIVectorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIVectorClass) Alloc() CIVector {
	rv := objc.Send[CIVector](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The Core Image class that defines a vector object.
//
// # Overview
//
// A [CIVector] can store one or more [CGFloat] in one object. They can store
// a group of float values for a variety of different uses such as coordinate
// points, direction vectors, geometric rectangles, transform matrices,
// convolution weights, or just a list a parameter values.
//
// You use [CIVector] objects in conjunction with other Core Image classes,
// such as [CIFilter] and [CIKernel]. Many of the built-in Core Image filters
// have one or more [CIVector] inputs that you can set to affect the
// filter’s behavior.
//
// # Initializing a Vector
//
//   - [CIVector.InitWithValuesCount]: Initialize a Core Image vector object with the specified the values.
//   - [CIVector.InitWithX]: Initialize a Core Image vector object with one value.
//   - [CIVector.InitWithXY]: Initialize a Core Image vector object with two values.
//   - [CIVector.InitWithXYZ]: Initialize a Core Image vector object with three values.
//   - [CIVector.InitWithXYZW]: Initialize a Core Image vector object with four values.
//   - [CIVector.InitWithString]: Initialize a Core Image vector object with values provided in a string representation.
//   - [CIVector.InitWithCGAffineTransform]: Initialize a Core Image vector object with six values provided by a [CGAffineTransform] structure.
//   - [CIVector.InitWithCGPoint]: Initialize a Core Image vector object with two values provided by a [CGPoint] structure.
//   - [CIVector.InitWithCGRect]: Initialize a Core Image vector object with four values provided by a [CGRect] structure.
//
// # Getting Values From a Vector
//
//   - [CIVector.ValueAtIndex]: Returns a value from a specific position in the vector.
//   - [CIVector.Count]: The number of items in the vector.
//   - [CIVector.X]: The value located in the first position in the vector.
//   - [CIVector.Y]: The value located in the second position in the vector.
//   - [CIVector.Z]: The value located in the third position in the vector.
//   - [CIVector.W]: The value located in the forth position in the vector.
//   - [CIVector.StringRepresentation]: Returns a formatted string with all the values of a [CIVector].
//   - [CIVector.CGAffineTransformValue]: Returns the values in the vector as a [CGAffineTransformValue] structure.
//   - [CIVector.CGPointValue]: Returns the values in the vector as a [CGPoint] structure.
//   - [CIVector.CGRectValue]: Returns the values in the vector as a [CGRect] structure.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector
type CIVector struct {
	objectivec.Object
}

// CIVectorFromID constructs a [CIVector] from an objc.ID.
//
// The Core Image class that defines a vector object.
func CIVectorFromID(id objc.ID) CIVector {
	return CIVector{objectivec.Object{ID: id}}
}

// NOTE: CIVector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIVector] class.
//
// # Initializing a Vector
//
//   - [ICIVector.InitWithValuesCount]: Initialize a Core Image vector object with the specified the values.
//   - [ICIVector.InitWithX]: Initialize a Core Image vector object with one value.
//   - [ICIVector.InitWithXY]: Initialize a Core Image vector object with two values.
//   - [ICIVector.InitWithXYZ]: Initialize a Core Image vector object with three values.
//   - [ICIVector.InitWithXYZW]: Initialize a Core Image vector object with four values.
//   - [ICIVector.InitWithString]: Initialize a Core Image vector object with values provided in a string representation.
//   - [ICIVector.InitWithCGAffineTransform]: Initialize a Core Image vector object with six values provided by a [CGAffineTransform] structure.
//   - [ICIVector.InitWithCGPoint]: Initialize a Core Image vector object with two values provided by a [CGPoint] structure.
//   - [ICIVector.InitWithCGRect]: Initialize a Core Image vector object with four values provided by a [CGRect] structure.
//
// # Getting Values From a Vector
//
//   - [ICIVector.ValueAtIndex]: Returns a value from a specific position in the vector.
//   - [ICIVector.Count]: The number of items in the vector.
//   - [ICIVector.X]: The value located in the first position in the vector.
//   - [ICIVector.Y]: The value located in the second position in the vector.
//   - [ICIVector.Z]: The value located in the third position in the vector.
//   - [ICIVector.W]: The value located in the forth position in the vector.
//   - [ICIVector.StringRepresentation]: Returns a formatted string with all the values of a [CIVector].
//   - [ICIVector.CGAffineTransformValue]: Returns the values in the vector as a [CGAffineTransformValue] structure.
//   - [ICIVector.CGPointValue]: Returns the values in the vector as a [CGPoint] structure.
//   - [ICIVector.CGRectValue]: Returns the values in the vector as a [CGRect] structure.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector
type ICIVector interface {
	objectivec.IObject

	// Topic: Initializing a Vector

	// Initialize a Core Image vector object with the specified the values.
	InitWithValuesCount(values []float64, count uintptr) CIVector
	// Initialize a Core Image vector object with one value.
	InitWithX(x float64) CIVector
	// Initialize a Core Image vector object with two values.
	InitWithXY(x float64, y float64) CIVector
	// Initialize a Core Image vector object with three values.
	InitWithXYZ(x float64, y float64, z float64) CIVector
	// Initialize a Core Image vector object with four values.
	InitWithXYZW(x float64, y float64, z float64, w float64) CIVector
	// Initialize a Core Image vector object with values provided in a string representation.
	InitWithString(representation string) CIVector
	// Initialize a Core Image vector object with six values provided by a [CGAffineTransform] structure.
	InitWithCGAffineTransform(t corefoundation.CGAffineTransform) CIVector
	// Initialize a Core Image vector object with two values provided by a [CGPoint] structure.
	InitWithCGPoint(p corefoundation.CGPoint) CIVector
	// Initialize a Core Image vector object with four values provided by a [CGRect] structure.
	InitWithCGRect(r corefoundation.CGRect) CIVector

	// Topic: Getting Values From a Vector

	// Returns a value from a specific position in the vector.
	ValueAtIndex(index uintptr) float64
	// The number of items in the vector.
	Count() uintptr
	// The value located in the first position in the vector.
	X() float64
	// The value located in the second position in the vector.
	Y() float64
	// The value located in the third position in the vector.
	Z() float64
	// The value located in the forth position in the vector.
	W() float64
	// Returns a formatted string with all the values of a [CIVector].
	StringRepresentation() string
	// Returns the values in the vector as a [CGAffineTransformValue] structure.
	CGAffineTransformValue() corefoundation.CGAffineTransform
	// Returns the values in the vector as a [CGPoint] structure.
	CGPointValue() corefoundation.CGPoint
	// Returns the values in the vector as a [CGRect] structure.
	CGRectValue() corefoundation.CGRect

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (v CIVector) Init() CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v CIVector) Autorelease() CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIVector creates a new CIVector instance.
func NewCIVector() CIVector {
	class := getCIVectorClass()
	rv := objc.Send[CIVector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize a Core Image vector object with six values provided by a
// [CGAffineTransform] structure.
//
// t: The [CGAffineTransform] structure.
//
// # Return Value
//
// An initialized [CIVector] object of length 6.
//
// # Discussion
//
// The [CGAffineTransform] structure’s `a`, `b`, `c`, `c`, `tx` and `ty`
// values are stored in the vector’s six values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgAffineTransform:)
func NewVectorWithCGAffineTransform(t corefoundation.CGAffineTransform) CIVector {
	instance := getCIVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGAffineTransform:"), t)
	return CIVectorFromID(rv)
}

// Initialize a Core Image vector object with two values provided by a
// [CGPoint] structure.
//
// p: The [CGPoint] structure.
//
// # Return Value
//
// An initialized [CIVector] object of length 2.
//
// # Discussion
//
// The [CGRect] structure’s `y` and `y` values are stored in the vector’s
// two values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgPoint:)
func NewVectorWithCGPoint(p corefoundation.CGPoint) CIVector {
	instance := getCIVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGPoint:"), p)
	return CIVectorFromID(rv)
}

// Initialize a Core Image vector object with four values provided by a
// [CGRect] structure.
//
// r: The [CGRect] structure.
//
// # Return Value
//
// An initialized [CIVector] object of length 4.
//
// # Discussion
//
// The [CGRect] structure’s `x`, `y`, `height` and `width` values are stored
// in the vector’s four values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgRect:)
func NewVectorWithCGRect(r corefoundation.CGRect) CIVector {
	instance := getCIVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGRect:"), r)
	return CIVectorFromID(rv)
}

// Initialize a Core Image vector object with values provided in a string
// representation.
//
// representation: A string that is in one of the formats returned by the
// `stringRepresentation` method.
//
// # Return Value
//
// An initialized [CIVector] object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(string:)
func NewVectorWithString(representation string) CIVector {
	instance := getCIVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(representation))
	return CIVectorFromID(rv)
}

// Initialize a Core Image vector object with the specified the values.
//
// values: A pointer [CGFloat] values for vector.
//
// count: The number of [CGFloats] specified by the `values` parameter.
//
// # Return Value
//
// An initialized [CIVector] object of length `count`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(values:count:)
func NewVectorWithValuesCount(values []float64, count uintptr) CIVector {
	instance := getCIVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValues:count:"), objc.CArray(values), count)
	return CIVectorFromID(rv)
}

// Initialize a Core Image vector object with one value.
//
// x: The value for the first position in the vector.
//
// # Return Value
//
// An initialized [CIVector] object of length 1.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:)
func NewVectorWithX(x float64) CIVector {
	instance := getCIVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithX:"), x)
	return CIVectorFromID(rv)
}

// Initialize a Core Image vector object with two values.
//
// x: The value for the first position in the vector.
//
// y: The value for the second position in the vector.
//
// # Return Value
//
// An initialized [CIVector] object of length 2.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:)
func NewVectorWithXY(x float64, y float64) CIVector {
	instance := getCIVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithX:Y:"), x, y)
	return CIVectorFromID(rv)
}

// Initialize a Core Image vector object with three values.
//
// x: The value for the first position in the vector.
//
// y: The value for the second position in the vector.
//
// z: The value for the third position in the vector.
//
// # Return Value
//
// An initialized [CIVector] object of length 3.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:z:)
func NewVectorWithXYZ(x float64, y float64, z float64) CIVector {
	instance := getCIVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithX:Y:Z:"), x, y, z)
	return CIVectorFromID(rv)
}

// Initialize a Core Image vector object with four values.
//
// x: The value for the first position in the vector.
//
// y: The value for the second position in the vector.
//
// z: The value for the third position in the vector.
//
// w: The value for the forth position in the vector.
//
// # Return Value
//
// An initialized [CIVector] object of length 4.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:z:w:)
func NewVectorWithXYZW(x float64, y float64, z float64, w float64) CIVector {
	instance := getCIVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithX:Y:Z:W:"), x, y, z, w)
	return CIVectorFromID(rv)
}

// Initialize a Core Image vector object with the specified the values.
//
// values: A pointer [CGFloat] values for vector.
//
// count: The number of [CGFloats] specified by the `values` parameter.
//
// # Return Value
//
// An initialized [CIVector] object of length `count`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(values:count:)
func (v CIVector) InitWithValuesCount(values []float64, count uintptr) CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("initWithValues:count:"), objc.CArray(values), count)
	return rv
}

// Initialize a Core Image vector object with one value.
//
// x: The value for the first position in the vector.
//
// # Return Value
//
// An initialized [CIVector] object of length 1.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:)
func (v CIVector) InitWithX(x float64) CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("initWithX:"), x)
	return rv
}

// Initialize a Core Image vector object with two values.
//
// x: The value for the first position in the vector.
//
// y: The value for the second position in the vector.
//
// # Return Value
//
// An initialized [CIVector] object of length 2.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:)
func (v CIVector) InitWithXY(x float64, y float64) CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("initWithX:Y:"), x, y)
	return rv
}

// Initialize a Core Image vector object with three values.
//
// x: The value for the first position in the vector.
//
// y: The value for the second position in the vector.
//
// z: The value for the third position in the vector.
//
// # Return Value
//
// An initialized [CIVector] object of length 3.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:z:)
func (v CIVector) InitWithXYZ(x float64, y float64, z float64) CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("initWithX:Y:Z:"), x, y, z)
	return rv
}

// Initialize a Core Image vector object with four values.
//
// x: The value for the first position in the vector.
//
// y: The value for the second position in the vector.
//
// z: The value for the third position in the vector.
//
// w: The value for the forth position in the vector.
//
// # Return Value
//
// An initialized [CIVector] object of length 4.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:z:w:)
func (v CIVector) InitWithXYZW(x float64, y float64, z float64, w float64) CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("initWithX:Y:Z:W:"), x, y, z, w)
	return rv
}

// Initialize a Core Image vector object with values provided in a string
// representation.
//
// representation: A string that is in one of the formats returned by the
// `stringRepresentation` method.
//
// # Return Value
//
// An initialized [CIVector] object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(string:)
func (v CIVector) InitWithString(representation string) CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("initWithString:"), objc.String(representation))
	return rv
}

// Initialize a Core Image vector object with six values provided by a
// [CGAffineTransform] structure.
//
// t: The [CGAffineTransform] structure.
//
// # Return Value
//
// An initialized [CIVector] object of length 6.
//
// # Discussion
//
// The [CGAffineTransform] structure’s `a`, `b`, `c`, `c`, `tx` and `ty`
// values are stored in the vector’s six values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgAffineTransform:)
func (v CIVector) InitWithCGAffineTransform(t corefoundation.CGAffineTransform) CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("initWithCGAffineTransform:"), t)
	return rv
}

// Initialize a Core Image vector object with two values provided by a
// [CGPoint] structure.
//
// p: The [CGPoint] structure.
//
// # Return Value
//
// An initialized [CIVector] object of length 2.
//
// # Discussion
//
// The [CGRect] structure’s `y` and `y` values are stored in the vector’s
// two values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgPoint:)
func (v CIVector) InitWithCGPoint(p corefoundation.CGPoint) CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("initWithCGPoint:"), p)
	return rv
}

// Initialize a Core Image vector object with four values provided by a
// [CGRect] structure.
//
// r: The [CGRect] structure.
//
// # Return Value
//
// An initialized [CIVector] object of length 4.
//
// # Discussion
//
// The [CGRect] structure’s `x`, `y`, `height` and `width` values are stored
// in the vector’s four values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgRect:)
func (v CIVector) InitWithCGRect(r corefoundation.CGRect) CIVector {
	rv := objc.Send[CIVector](v.ID, objc.Sel("initWithCGRect:"), r)
	return rv
}

// Returns a value from a specific position in the vector.
//
// index: The position in the vector of the value that you want to retrieve.
//
// # Return Value
//
// The value retrieved from the vector or `0` if the position is undefined.
//
// # Discussion
//
// The numbering of elements in a vector begins with zero.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/value(at:)
func (v CIVector) ValueAtIndex(index uintptr) float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("valueAtIndex:"), index)
	return rv
}
func (v CIVector) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Create a Core Image vector object that is initialized with six values
// provided by a [CGAffineTransform] structure.
//
// t: The [CGAffineTransform] structure.
//
// # Return Value
//
// An autoreleased [CIVector] object of length 6.
//
// # Discussion
//
// The [CGAffineTransform] structure’s `a`, `b`, `c`, `d`, `tx` and `ty`
// values are stored in the vector’s six values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGAffineTransform:
func (_CIVectorClass CIVectorClass) VectorWithCGAffineTransform(t corefoundation.CGAffineTransform) CIVector {
	rv := objc.Send[objc.ID](objc.ID(_CIVectorClass.class), objc.Sel("vectorWithCGAffineTransform:"), t)
	return CIVectorFromID(rv)
}

// Create a Core Image vector object that is initialized with two values
// provided by a [CGPoint] structure.
//
// p: The [CGPoint] structure.
//
// # Return Value
//
// An autoreleased [CIVector] object of length 2.
//
// # Discussion
//
// The [CGRect] structure’s `y` and `y` values are stored in the vector’s
// two values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGPoint:
func (_CIVectorClass CIVectorClass) VectorWithCGPoint(p corefoundation.CGPoint) CIVector {
	rv := objc.Send[objc.ID](objc.ID(_CIVectorClass.class), objc.Sel("vectorWithCGPoint:"), p)
	return CIVectorFromID(rv)
}

// Create a Core Image vector object that is initialized with four values
// provided by a [CGRect] structure.
//
// r: The [CGRect] structure.
//
// # Return Value
//
// An autoreleased [CIVector] object of length 4.
//
// # Discussion
//
// The [CGRect] structure’s `x`, `y`, `height` and `width` values are stored
// in the vector’s four values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGRect:
func (_CIVectorClass CIVectorClass) VectorWithCGRect(r corefoundation.CGRect) CIVector {
	rv := objc.Send[objc.ID](objc.ID(_CIVectorClass.class), objc.Sel("vectorWithCGRect:"), r)
	return CIVectorFromID(rv)
}

// Create a Core Image vector object with values provided in a string
// representation.
//
// representation: A string that is in one of the formats returned by the
// `stringRepresentation` method.
//
// # Return Value
//
// An autoreleased [CIVector] object.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithString:
func (_CIVectorClass CIVectorClass) VectorWithString(representation string) CIVector {
	rv := objc.Send[objc.ID](objc.ID(_CIVectorClass.class), objc.Sel("vectorWithString:"), objc.String(representation))
	return CIVectorFromID(rv)
}

// Create a Core Image vector object that is initialized with the specified
// values.
//
// values: The pointer [CGFloat] values to initialize the vector with.
//
// count: The number of [CGFloats] specified by the `values` parameter.
//
// # Return Value
//
// An autoreleased [CIVector] object of length `count`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithValues:count:
func (_CIVectorClass CIVectorClass) VectorWithValuesCount(values []float64, count uintptr) CIVector {
	rv := objc.Send[objc.ID](objc.ID(_CIVectorClass.class), objc.Sel("vectorWithValues:count:"), objc.CArray(values), count)
	return CIVectorFromID(rv)
}

// Create a Core Image vector object that is initialized with one value.
//
// x: The value for the first position in the vector.
//
// # Return Value
//
// An autoreleased [CIVector] object of length 1.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:
func (_CIVectorClass CIVectorClass) VectorWithX(x float64) CIVector {
	rv := objc.Send[objc.ID](objc.ID(_CIVectorClass.class), objc.Sel("vectorWithX:"), x)
	return CIVectorFromID(rv)
}

// Create a Core Image vector object that is initialized with two values.
//
// x: The value for the first position in the vector.
//
// y: The value for the second position in the vector.
//
// # Return Value
//
// An autoreleased [CIVector] object of length 2.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:
func (_CIVectorClass CIVectorClass) VectorWithXY(x float64, y float64) CIVector {
	rv := objc.Send[objc.ID](objc.ID(_CIVectorClass.class), objc.Sel("vectorWithX:Y:"), x, y)
	return CIVectorFromID(rv)
}

// Create a Core Image vector object that is initialized with three values.
//
// x: The value for the first position in the vector.
//
// y: The value for the second position in the vector.
//
// z: The value for the third position in the vector.
//
// # Return Value
//
// An autoreleased [CIVector] object of length 3.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:Z:
func (_CIVectorClass CIVectorClass) VectorWithXYZ(x float64, y float64, z float64) CIVector {
	rv := objc.Send[objc.ID](objc.ID(_CIVectorClass.class), objc.Sel("vectorWithX:Y:Z:"), x, y, z)
	return CIVectorFromID(rv)
}

// Create a Core Image vector object that is initialized with four values.
//
// x: The value for the first position in the vector.
//
// y: The value for the second position in the vector.
//
// z: The value for the third position in the vector.
//
// w: The value for the forth position in the vector.
//
// # Return Value
//
// An autoreleased [CIVector] object of length 4.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:Z:W:
func (_CIVectorClass CIVectorClass) VectorWithXYZW(x float64, y float64, z float64, w float64) CIVector {
	rv := objc.Send[objc.ID](objc.ID(_CIVectorClass.class), objc.Sel("vectorWithX:Y:Z:W:"), x, y, z, w)
	return CIVectorFromID(rv)
}

// The number of items in the vector.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/count
func (v CIVector) Count() uintptr {
	rv := objc.Send[uintptr](v.ID, objc.Sel("count"))
	return rv
}

// The value located in the first position in the vector.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/x
func (v CIVector) X() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("X"))
	return rv
}

// The value located in the second position in the vector.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/y
func (v CIVector) Y() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("Y"))
	return rv
}

// The value located in the third position in the vector.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/z
func (v CIVector) Z() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("Z"))
	return rv
}

// The value located in the forth position in the vector.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/w
func (v CIVector) W() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("W"))
	return rv
}

// Returns a formatted string with all the values of a [CIVector].
//
// # Discussion
//
// Some example string representations of vectors:
//
// [Table data omitted]
//
// To create a [CIVector] object from a string representation, use the
// [VectorWithString] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/stringRepresentation
func (v CIVector) StringRepresentation() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("stringRepresentation"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the values in the vector as a [CGAffineTransformValue] structure.
//
// # Return Value
//
// Reading this property creates a [CGAffineTransformValue] structure from the
// first six values in the vector.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/cgAffineTransformValue
func (v CIVector) CGAffineTransformValue() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](v.ID, objc.Sel("CGAffineTransformValue"))
	return corefoundation.CGAffineTransform(rv)
}

// Returns the values in the vector as a [CGPoint] structure.
//
// # Return Value
//
// Reading this property returns a [CGPoint] structure from the [X] and [Y]
// values from the vector.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/cgPointValue
func (v CIVector) CGPointValue() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("CGPointValue"))
	return corefoundation.CGPoint(rv)
}

// Returns the values in the vector as a [CGRect] structure.
//
// # Return Value
//
// Reading this property creates a [CGRect] structure whose origin is the [X],
// [Y], [Z] and [W] values from the vector.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/cgRectValue
func (v CIVector) CGRectValue() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("CGRectValue"))
	return corefoundation.CGRect(rv)
}

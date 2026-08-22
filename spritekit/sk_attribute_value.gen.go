// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKAttributeValue] class.
var (
	_SKAttributeValueClass     SKAttributeValueClass
	_SKAttributeValueClassOnce sync.Once
)

func getSKAttributeValueClass() SKAttributeValueClass {
	_SKAttributeValueClassOnce.Do(func() {
		_SKAttributeValueClass = SKAttributeValueClass{class: objc.GetClass("SKAttributeValue")}
	})
	return _SKAttributeValueClass
}

// GetSKAttributeValueClass returns the class object for SKAttributeValue.
func GetSKAttributeValueClass() SKAttributeValueClass {
	return getSKAttributeValueClass()
}

type SKAttributeValueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKAttributeValueClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKAttributeValueClass) Alloc() SKAttributeValue {
	rv := objc.Send[SKAttributeValue](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A container for dynamic shader data associated with a node.
//
// # Overview
//
// SpriteKit nodes that are rendered with a custom shader can use
// [SKAttributeValue] objects to pass dynamic values which can change without
// requiring that shader to be recompiled. An attribute value is passed to a
// shader using a node’s [setValue(_:forAttribute:)] method using the
// relevant attribute’s name. For example, given a shader with a
// [SKAttributeType.float] attribute named `a_radius`:
//
// Listing 1. Creating an attribute
//
// The following code sets the value of this attribute to `10` and passes it
// to a [SKSpriteNode] object’s shader:
//
// Listing 2. Setting an attribute value
//
// The attribute, `a_radius`, is available as a global floating-point variable
// within the shader code.
//
// Using this technique, a single shader can be shared across many nodes and
// each nodes can supply its own attributes. This approach is an alternative
// to using [SKUniform] objects which would require a recompilation for each
// distinct set of parameters.
//
// # Instance Properties
//
//   - [SKAttributeValue.FloatValue]: The receiver’s floating point value.
//   - [SKAttributeValue.SetFloatValue]
//   - [SKAttributeValue.VectorFloat2Value]: The receiver’s value as a vector of two floating-point numbers.
//   - [SKAttributeValue.SetVectorFloat2Value]
//   - [SKAttributeValue.VectorFloat3Value]: The receiver’s value as a vector of three floating point numbers.
//   - [SKAttributeValue.SetVectorFloat3Value]
//   - [SKAttributeValue.VectorFloat4Value]: The receiver’s value as a vector of four floating point numbers.
//   - [SKAttributeValue.SetVectorFloat4Value]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue
//
// [SKAttributeType.float]: https://developer.apple.com/documentation/SpriteKit/SKAttributeType/float
// [setValue(_:forAttribute:)]: https://developer.apple.com/documentation/SpriteKit/SKNode/setValue(_:forAttribute:)
type SKAttributeValue struct {
	objectivec.Object
}

// SKAttributeValueFromID constructs a [SKAttributeValue] from an objc.ID.
//
// A container for dynamic shader data associated with a node.
func SKAttributeValueFromID(id objc.ID) SKAttributeValue {
	return SKAttributeValue{objectivec.Object{ID: id}}
}

// NOTE: SKAttributeValue adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKAttributeValue] class.
//
// # Instance Properties
//
//   - [ISKAttributeValue.FloatValue]: The receiver’s floating point value.
//   - [ISKAttributeValue.SetFloatValue]
//   - [ISKAttributeValue.VectorFloat2Value]: The receiver’s value as a vector of two floating-point numbers.
//   - [ISKAttributeValue.SetVectorFloat2Value]
//   - [ISKAttributeValue.VectorFloat3Value]: The receiver’s value as a vector of three floating point numbers.
//   - [ISKAttributeValue.SetVectorFloat3Value]
//   - [ISKAttributeValue.VectorFloat4Value]: The receiver’s value as a vector of four floating point numbers.
//   - [ISKAttributeValue.SetVectorFloat4Value]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue
type ISKAttributeValue interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The receiver’s floating point value.
	FloatValue() float32
	SetFloatValue(value float32)
	// The receiver’s value as a vector of two floating-point numbers.
	VectorFloat2Value() [2]float32
	SetVectorFloat2Value(value [2]float32)
	// The receiver’s value as a vector of three floating point numbers.
	VectorFloat3Value() Vector_float3
	SetVectorFloat3Value(value Vector_float3)
	// The receiver’s value as a vector of four floating point numbers.
	VectorFloat4Value() [4]float32
	SetVectorFloat4Value(value [4]float32)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a SKAttributeValue) Init() SKAttributeValue {
	rv := objc.Send[SKAttributeValue](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a SKAttributeValue) Autorelease() SKAttributeValue {
	rv := objc.Send[SKAttributeValue](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKAttributeValue creates a new SKAttributeValue instance.
func NewSKAttributeValue() SKAttributeValue {
	class := getSKAttributeValueClass()
	rv := objc.Send[SKAttributeValue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and initializes a new attribute value object that holds a floating
// point number.
//
// value: The floating point value for the new attribute value.
//
// # Return Value
//
// A new attribute value object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(float:)
func NewAttributeValueWithFloat(value float32) SKAttributeValue {
	rv := objc.Send[objc.ID](objc.ID(getSKAttributeValueClass().class), objc.Sel("valueWithFloat:"), value)
	return SKAttributeValueFromID(rv)
}

// Creates and initializes a new attribute value object that holds a vector of
// two floating point numbers.
//
// value: The vector of two point values for the new attribute value.
//
// # Return Value
//
// A new attribute value object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(vectorFloat2:)
func NewAttributeValueWithVectorFloat2(value [2]float32) SKAttributeValue {
	rv := objc.Send[objc.ID](objc.ID(getSKAttributeValueClass().class), objc.Sel("valueWithVectorFloat2:"), value)
	return SKAttributeValueFromID(rv)
}

// Creates and initializes a new attribute value object that holds a vector of
// three floating point numbers.
//
// value: The vector of three point values for the new attribute value.
//
// # Return Value
//
// A new attribute value object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(vectorFloat3:)
func NewAttributeValueWithVectorFloat3(value Vector_float3) SKAttributeValue {
	rv := objc.Send[objc.ID](objc.ID(getSKAttributeValueClass().class), objc.Sel("valueWithVectorFloat3:"), value)
	return SKAttributeValueFromID(rv)
}

// Creates and initializes a new attribute value object that holds a vector of
// four floating point numbers.
//
// value: The vector of four point values for the new attribute value.
//
// # Return Value
//
// A new attribute value object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(vectorFloat4:)
func NewAttributeValueWithVectorFloat4(value [4]float32) SKAttributeValue {
	rv := objc.Send[objc.ID](objc.ID(getSKAttributeValueClass().class), objc.Sel("valueWithVectorFloat4:"), value)
	return SKAttributeValueFromID(rv)
}

func (a SKAttributeValue) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The receiver’s floating point value.
//
// # Discussion
//
// If the receiver’s original value is a vector,
// [SKAttributeValue.FloatValue] is the first element of the vector.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/floatValue
func (a SKAttributeValue) FloatValue() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("floatValue"))
	return rv
}
func (a SKAttributeValue) SetFloatValue(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setFloatValue:"), value)
}

// The receiver’s value as a vector of two floating-point numbers.
//
// # Discussion
//
// If the receiver’s original value is a floating-point number, the second
// item of the vector is set to 0. If the receiver’s original value is a
// vector of more than two items, the additional items are truncated.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/vectorFloat2Value
func (a SKAttributeValue) VectorFloat2Value() [2]float32 {
	rv := objc.Send[[2]float32](a.ID, objc.Sel("vectorFloat2Value"))
	return [2]float32(rv)
}
func (a SKAttributeValue) SetVectorFloat2Value(value [2]float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setVectorFloat2Value:"), value)
}

// The receiver’s value as a vector of three floating point numbers.
//
// # Discussion
//
// If the receiver’s original value is a floating-point number or a
// [vector_float2], the empty items in the vector are set to 0. If the
// receiver’s original value is a [vector_float4], the last item is
// truncated.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/vectorFloat3Value
//
// [vector_float2]: https://developer.apple.com/documentation/simd/vector_float2
// [vector_float4]: https://developer.apple.com/documentation/simd/vector_float4
func (a SKAttributeValue) VectorFloat3Value() Vector_float3 {
	rv := objc.Send[Vector_float3](a.ID, objc.Sel("vectorFloat3Value"))
	return Vector_float3(rv)
}
func (a SKAttributeValue) SetVectorFloat3Value(value Vector_float3) {
	objc.Send[struct{}](a.ID, objc.Sel("setVectorFloat3Value:"), value)
}

// The receiver’s value as a vector of four floating point numbers.
//
// # Discussion
//
// If the receiver’s original value is a floating-point number of a vector
// with a count less than four, the empty items in the vector are set to 0.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/vectorFloat4Value
func (a SKAttributeValue) VectorFloat4Value() [4]float32 {
	rv := objc.Send[[4]float32](a.ID, objc.Sel("vectorFloat4Value"))
	return [4]float32(rv)
}
func (a SKAttributeValue) SetVectorFloat4Value(value [4]float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setVectorFloat4Value:"), value)
}

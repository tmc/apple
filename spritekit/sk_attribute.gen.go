// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKAttribute] class.
var (
	_SKAttributeClass     SKAttributeClass
	_SKAttributeClassOnce sync.Once
)

func getSKAttributeClass() SKAttributeClass {
	_SKAttributeClassOnce.Do(func() {
		_SKAttributeClass = SKAttributeClass{class: objc.GetClass("SKAttribute")}
	})
	return _SKAttributeClass
}

// GetSKAttributeClass returns the class object for SKAttribute.
func GetSKAttributeClass() SKAttributeClass {
	return getSKAttributeClass()
}

type SKAttributeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKAttributeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKAttributeClass) Alloc() SKAttribute {
	rv := objc.Send[SKAttribute](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A specification for dynamic per-node data used with a custom shader.
//
// # Overview
//
// To define an attribute for your shader, you create an [SKAttribute] object
// with a unique name and data type, which is a [SKAttributeType] enum. After
// creating an [SKShader] object, custom attributes are added to its
// [SKShader.Attributes] array. Attribute values are set on the parent node
// with [setValue(_:forAttribute:)] and can change for each execution of a
// shader without the need for recompilation.
//
// The following listing shows how you can use an attribute to pass the size
// of a sprite into a shader using an attribute. In this example,
// `a_sprite_size` is available as a global `vec2` within the GLSL code.
//
// Listing 1. Passing an attribute to a shader.
//
// # Initializers
//
//   - [SKAttribute.InitWithNameType]: Creates and initializes a new attribute object of a specified type with a name that can be referenced within the shader.
//
// # Instance Properties
//
//   - [SKAttribute.Name]: The receiver’s name
//   - [SKAttribute.Type]: The data type of the attribute’s value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttribute
//
// [SKAttributeType]: https://developer.apple.com/documentation/SpriteKit/SKAttributeType
// [setValue(_:forAttribute:)]: https://developer.apple.com/documentation/SpriteKit/SKNode/setValue(_:forAttribute:)
type SKAttribute struct {
	objectivec.Object
}

// SKAttributeFromID constructs a [SKAttribute] from an objc.ID.
//
// A specification for dynamic per-node data used with a custom shader.
func SKAttributeFromID(id objc.ID) SKAttribute {
	return SKAttribute{objectivec.Object{ID: id}}
}

// NOTE: SKAttribute adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKAttribute] class.
//
// # Initializers
//
//   - [ISKAttribute.InitWithNameType]: Creates and initializes a new attribute object of a specified type with a name that can be referenced within the shader.
//
// # Instance Properties
//
//   - [ISKAttribute.Name]: The receiver’s name
//   - [ISKAttribute.Type]: The data type of the attribute’s value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttribute
type ISKAttribute interface {
	objectivec.IObject

	// Topic: Initializers

	// Creates and initializes a new attribute object of a specified type with a name that can be referenced within the shader.
	InitWithNameType(name string, type_ SKAttributeType) SKAttribute

	// Topic: Instance Properties

	// The receiver’s name
	Name() string
	// The data type of the attribute’s value.
	Type() SKAttributeType

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a SKAttribute) Init() SKAttribute {
	rv := objc.Send[SKAttribute](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a SKAttribute) Autorelease() SKAttribute {
	rv := objc.Send[SKAttribute](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKAttribute creates a new SKAttribute instance.
func NewSKAttribute() SKAttribute {
	class := getSKAttributeClass()
	rv := objc.Send[SKAttribute](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and initializes a new attribute object of a specified type with a
// name that can be referenced within the shader.
//
// name: The name of the attribute.
//
// type: The type of the attribute.
//
// # Return Value
//
// A new attribute object.
//
// # Discussion
//
// Attribute names are typically named with a preceding “a” and an
// underscore. The following code shows how to initialize an attribute named
// `a_frequency` which is of type [SKAttributeType.float].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttribute/init(name:type:)
//
// [SKAttributeType.float]: https://developer.apple.com/documentation/SpriteKit/SKAttributeType/float
func NewAttributeWithNameType(name string, type_ SKAttributeType) SKAttribute {
	instance := getSKAttributeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:type:"), objc.String(name), type_)
	return SKAttributeFromID(rv)
}

// Creates and initializes a new attribute object of a specified type with a
// name that can be referenced within the shader.
//
// name: The name of the attribute.
//
// type: The type of the attribute.
//
// # Return Value
//
// A new attribute object.
//
// # Discussion
//
// Attribute names are typically named with a preceding “a” and an
// underscore. The following code shows how to initialize an attribute named
// `a_frequency` which is of type [SKAttributeType.float].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttribute/init(name:type:)
//
// [SKAttributeType.float]: https://developer.apple.com/documentation/SpriteKit/SKAttributeType/float
func (a SKAttribute) InitWithNameType(name string, type_ SKAttributeType) SKAttribute {
	rv := objc.Send[SKAttribute](a.ID, objc.Sel("initWithName:type:"), objc.String(name), type_)
	return rv
}
func (a SKAttribute) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The receiver’s name
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttribute/name
func (a SKAttribute) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The data type of the attribute’s value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAttribute/type
func (a SKAttribute) Type() SKAttributeType {
	rv := objc.Send[SKAttributeType](a.ID, objc.Sel("type"))
	return SKAttributeType(rv)
}

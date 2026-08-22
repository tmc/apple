// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKUniform] class.
var (
	_SKUniformClass     SKUniformClass
	_SKUniformClassOnce sync.Once
)

func getSKUniformClass() SKUniformClass {
	_SKUniformClassOnce.Do(func() {
		_SKUniformClass = SKUniformClass{class: objc.GetClass("SKUniform")}
	})
	return _SKUniformClass
}

// GetSKUniformClass returns the class object for SKUniform.
func GetSKUniformClass() SKUniformClass {
	return getSKUniformClass()
}

type SKUniformClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKUniformClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKUniformClass) Alloc() SKUniform {
	rv := objc.Send[SKUniform](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A container for uniform shader data.
//
// # Overview
//
// An [SKUniform] object is used to hold uniform data for a custom OpenGL or
// OpenGL ES shader. The uniform data is accessible from all shaders that
// include the uniform.To use a uniform variable in your shader, create the
// [SKUniform] object and set its initial value. Once its value is specified,
// the [SKUniform.UniformType] property changes to match the type of the
// initial value you provided and can never change afterward. To use the
// uniform object, add it to an [SKShader] object that needs to access the
// uniform variable. To update the uniform variable’s value, choose the
// appropriate property on the uniform object based on the data type it
// encapsulates.
//
// # Creating and Initializing Uniform Objects
//
//   - [SKUniform.InitWithName]: Initializes a new uniform object.
//   - [SKUniform.InitWithNameFloat]: Initializes a new uniform object that holds a floating-point number.
//   - [SKUniform.InitWithNameTexture]: Initializes a new uniform object that holds a reference to a texture.
//
// # Reading Information About a Uniform
//
//   - [SKUniform.Name]: The uniform’s name.
//   - [SKUniform.UniformType]: The uniform object’s data type.
//
// # Reading and Writing an Uniform Object’s Value
//
//   - [SKUniform.FloatValue]: The receiver’s value as a floating-point value.
//   - [SKUniform.SetFloatValue]
//   - [SKUniform.TextureValue]: The receiver’s value as a SpriteKit texture.
//   - [SKUniform.SetTextureValue]
//
// # Initializers
//
//   - [SKUniform.InitWithNameMatrixFloat2x2]
//   - [SKUniform.InitWithNameMatrixFloat3x3]
//   - [SKUniform.InitWithNameMatrixFloat4x4]
//   - [SKUniform.InitWithNameVectorFloat2]
//   - [SKUniform.InitWithNameVectorFloat3]
//   - [SKUniform.InitWithNameVectorFloat4]
//
// # Instance Properties
//
//   - [SKUniform.MatrixFloat2x2Value]
//   - [SKUniform.SetMatrixFloat2x2Value]
//   - [SKUniform.MatrixFloat3x3Value]
//   - [SKUniform.SetMatrixFloat3x3Value]
//   - [SKUniform.MatrixFloat4x4Value]
//   - [SKUniform.SetMatrixFloat4x4Value]
//   - [SKUniform.VectorFloat2Value]
//   - [SKUniform.SetVectorFloat2Value]
//   - [SKUniform.VectorFloat3Value]
//   - [SKUniform.SetVectorFloat3Value]
//   - [SKUniform.VectorFloat4Value]
//   - [SKUniform.SetVectorFloat4Value]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform
type SKUniform struct {
	objectivec.Object
}

// SKUniformFromID constructs a [SKUniform] from an objc.ID.
//
// A container for uniform shader data.
func SKUniformFromID(id objc.ID) SKUniform {
	return SKUniform{objectivec.Object{ID: id}}
}

// NOTE: SKUniform adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKUniform] class.
//
// # Creating and Initializing Uniform Objects
//
//   - [ISKUniform.InitWithName]: Initializes a new uniform object.
//   - [ISKUniform.InitWithNameFloat]: Initializes a new uniform object that holds a floating-point number.
//   - [ISKUniform.InitWithNameTexture]: Initializes a new uniform object that holds a reference to a texture.
//
// # Reading Information About a Uniform
//
//   - [ISKUniform.Name]: The uniform’s name.
//   - [ISKUniform.UniformType]: The uniform object’s data type.
//
// # Reading and Writing an Uniform Object’s Value
//
//   - [ISKUniform.FloatValue]: The receiver’s value as a floating-point value.
//   - [ISKUniform.SetFloatValue]
//   - [ISKUniform.TextureValue]: The receiver’s value as a SpriteKit texture.
//   - [ISKUniform.SetTextureValue]
//
// # Initializers
//
//   - [ISKUniform.InitWithNameMatrixFloat2x2]
//   - [ISKUniform.InitWithNameMatrixFloat3x3]
//   - [ISKUniform.InitWithNameMatrixFloat4x4]
//   - [ISKUniform.InitWithNameVectorFloat2]
//   - [ISKUniform.InitWithNameVectorFloat3]
//   - [ISKUniform.InitWithNameVectorFloat4]
//
// # Instance Properties
//
//   - [ISKUniform.MatrixFloat2x2Value]
//   - [ISKUniform.SetMatrixFloat2x2Value]
//   - [ISKUniform.MatrixFloat3x3Value]
//   - [ISKUniform.SetMatrixFloat3x3Value]
//   - [ISKUniform.MatrixFloat4x4Value]
//   - [ISKUniform.SetMatrixFloat4x4Value]
//   - [ISKUniform.VectorFloat2Value]
//   - [ISKUniform.SetVectorFloat2Value]
//   - [ISKUniform.VectorFloat3Value]
//   - [ISKUniform.SetVectorFloat3Value]
//   - [ISKUniform.VectorFloat4Value]
//   - [ISKUniform.SetVectorFloat4Value]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform
type ISKUniform interface {
	objectivec.IObject

	// Topic: Creating and Initializing Uniform Objects

	// Initializes a new uniform object.
	InitWithName(name string) SKUniform
	// Initializes a new uniform object that holds a floating-point number.
	InitWithNameFloat(name string, value float32) SKUniform
	// Initializes a new uniform object that holds a reference to a texture.
	InitWithNameTexture(name string, texture ISKTexture) SKUniform

	// Topic: Reading Information About a Uniform

	// The uniform’s name.
	Name() string
	// The uniform object’s data type.
	UniformType() SKUniformType

	// Topic: Reading and Writing an Uniform Object’s Value

	// The receiver’s value as a floating-point value.
	FloatValue() float32
	SetFloatValue(value float32)
	// The receiver’s value as a SpriteKit texture.
	TextureValue() ISKTexture
	SetTextureValue(value ISKTexture)

	// Topic: Initializers

	InitWithNameMatrixFloat2x2(name string, value [2][2]float32) SKUniform
	InitWithNameMatrixFloat3x3(name string, value [3][4]float32) SKUniform
	InitWithNameMatrixFloat4x4(name string, value [4][4]float32) SKUniform
	InitWithNameVectorFloat2(name string, value [2]float32) SKUniform
	InitWithNameVectorFloat3(name string, value Vector_float3) SKUniform
	InitWithNameVectorFloat4(name string, value [4]float32) SKUniform

	// Topic: Instance Properties

	MatrixFloat2x2Value() [2][2]float32
	SetMatrixFloat2x2Value(value [2][2]float32)
	MatrixFloat3x3Value() [3][4]float32
	SetMatrixFloat3x3Value(value [3][4]float32)
	MatrixFloat4x4Value() [4][4]float32
	SetMatrixFloat4x4Value(value [4][4]float32)
	VectorFloat2Value() [2]float32
	SetVectorFloat2Value(value [2]float32)
	VectorFloat3Value() Vector_float3
	SetVectorFloat3Value(value Vector_float3)
	VectorFloat4Value() [4]float32
	SetVectorFloat4Value(value [4]float32)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u SKUniform) Init() SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u SKUniform) Autorelease() SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKUniform creates a new SKUniform instance.
func NewSKUniform() SKUniform {
	class := getSKUniformClass()
	rv := objc.Send[SKUniform](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new uniform object.
//
// name: The name used to identify the uniform variable; you use this name inside
// your shader to read the uniform variable’s value.
//
// # Return Value
//
// An initialized uniform object.
//
// # Discussion
//
// A uniform initialized with this method has no initial type and cannot be
// used in a shader until it is given an initial value. To set the initial
// value, use one of the properties defined in [SKUniform]. After its value is
// set, its [SKUniform.UniformType] property is set to match the uniform’s
// new type. Once set, the type may not be changed.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:)
func NewUniformWithName(name string) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), objc.String(name))
	return SKUniformFromID(rv)
}

// Initializes a new uniform object that holds a floating-point number.
//
// name: The name used to identify the uniform variable; you use this name inside
// your shader to read the uniform variable’s value.
//
// value: The initial floating-point value for the uniform variable.
//
// # Return Value
//
// An initialized uniform object whose type is set to [SKUniformType.float].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:float:)-48rln
//
// [SKUniformType.float]: https://developer.apple.com/documentation/SpriteKit/SKUniformType/float
func NewUniformWithNameFloat(name string, value float32) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:float:"), objc.String(name), value)
	return SKUniformFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:matrixFloat2x2:)
func NewUniformWithNameMatrixFloat2x2(name string, value [2][2]float32) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:matrixFloat2x2:"), objc.String(name), value)
	return SKUniformFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:matrixFloat3x3:)
func NewUniformWithNameMatrixFloat3x3(name string, value [3][4]float32) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:matrixFloat3x3:"), objc.String(name), value)
	return SKUniformFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:matrixFloat4x4:)
func NewUniformWithNameMatrixFloat4x4(name string, value [4][4]float32) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:matrixFloat4x4:"), objc.String(name), value)
	return SKUniformFromID(rv)
}

// Initializes a new uniform object that holds a reference to a texture.
//
// name: The name used to identify the uniform variable; you use this name inside
// your shader to read the uniform variable’s value.
//
// texture: The initial texture to use for the uniform variable.
//
// # Return Value
//
// An initialized uniform object whose type is set to [SKUniformType.texture].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:texture:)
//
// [SKUniformType.texture]: https://developer.apple.com/documentation/SpriteKit/SKUniformType/texture
func NewUniformWithNameTexture(name string, texture ISKTexture) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:texture:"), objc.String(name), texture)
	return SKUniformFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:vectorFloat2:)
func NewUniformWithNameVectorFloat2(name string, value [2]float32) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:vectorFloat2:"), objc.String(name), value)
	return SKUniformFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:vectorFloat3:)
func NewUniformWithNameVectorFloat3(name string, value Vector_float3) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:vectorFloat3:"), objc.String(name), value)
	return SKUniformFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:vectorFloat4:)
func NewUniformWithNameVectorFloat4(name string, value [4]float32) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:vectorFloat4:"), objc.String(name), value)
	return SKUniformFromID(rv)
}

// Initializes a new uniform object.
//
// name: The name used to identify the uniform variable; you use this name inside
// your shader to read the uniform variable’s value.
//
// # Return Value
//
// An initialized uniform object.
//
// # Discussion
//
// A uniform initialized with this method has no initial type and cannot be
// used in a shader until it is given an initial value. To set the initial
// value, use one of the properties defined in [SKUniform]. After its value is
// set, its [SKUniform.UniformType] property is set to match the uniform’s
// new type. Once set, the type may not be changed.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:)
func (u SKUniform) InitWithName(name string) SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("initWithName:"), objc.String(name))
	return rv
}

// Initializes a new uniform object that holds a floating-point number.
//
// name: The name used to identify the uniform variable; you use this name inside
// your shader to read the uniform variable’s value.
//
// value: The initial floating-point value for the uniform variable.
//
// # Return Value
//
// An initialized uniform object whose type is set to [SKUniformType.float].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:float:)-48rln
//
// [SKUniformType.float]: https://developer.apple.com/documentation/SpriteKit/SKUniformType/float
func (u SKUniform) InitWithNameFloat(name string, value float32) SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("initWithName:float:"), objc.String(name), value)
	return rv
}

// Initializes a new uniform object that holds a reference to a texture.
//
// name: The name used to identify the uniform variable; you use this name inside
// your shader to read the uniform variable’s value.
//
// texture: The initial texture to use for the uniform variable.
//
// # Return Value
//
// An initialized uniform object whose type is set to [SKUniformType.texture].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:texture:)
//
// [SKUniformType.texture]: https://developer.apple.com/documentation/SpriteKit/SKUniformType/texture
func (u SKUniform) InitWithNameTexture(name string, texture ISKTexture) SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("initWithName:texture:"), objc.String(name), texture)
	return rv
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:matrixFloat2x2:)
func (u SKUniform) InitWithNameMatrixFloat2x2(name string, value [2][2]float32) SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("initWithName:matrixFloat2x2:"), objc.String(name), value)
	return rv
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:matrixFloat3x3:)
func (u SKUniform) InitWithNameMatrixFloat3x3(name string, value [3][4]float32) SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("initWithName:matrixFloat3x3:"), objc.String(name), value)
	return rv
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:matrixFloat4x4:)
func (u SKUniform) InitWithNameMatrixFloat4x4(name string, value [4][4]float32) SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("initWithName:matrixFloat4x4:"), objc.String(name), value)
	return rv
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:vectorFloat2:)
func (u SKUniform) InitWithNameVectorFloat2(name string, value [2]float32) SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("initWithName:vectorFloat2:"), objc.String(name), value)
	return rv
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:vectorFloat3:)
func (u SKUniform) InitWithNameVectorFloat3(name string, value Vector_float3) SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("initWithName:vectorFloat3:"), objc.String(name), value)
	return rv
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:vectorFloat4:)
func (u SKUniform) InitWithNameVectorFloat4(name string, value [4]float32) SKUniform {
	rv := objc.Send[SKUniform](u.ID, objc.Sel("initWithName:vectorFloat4:"), objc.String(name), value)
	return rv
}
func (u SKUniform) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The uniform’s name.
//
// # Discussion
//
// Your custom fragment shader uses this name to identify the variable.
// SpriteKit automatically declares the uniform variable for your shader.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/name
func (u SKUniform) Name() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The uniform object’s data type.
//
// # Discussion
//
// A uniform object’s type is set to [SKUniformType.none] until the first
// time that the uniform variable’s value is set; this happens automatically
// if you use an initialization method that provides an initial type and
// value. Once the uniform object is given an initial value, its type changes
// to that value’s type and thereafter cannot be changed.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformType
//
// [SKUniformType.none]: https://developer.apple.com/documentation/SpriteKit/SKUniformType/none
func (u SKUniform) UniformType() SKUniformType {
	rv := objc.Send[SKUniformType](u.ID, objc.Sel("uniformType"))
	return SKUniformType(rv)
}

// The receiver’s value as a floating-point value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/floatValue
func (u SKUniform) FloatValue() float32 {
	rv := objc.Send[float32](u.ID, objc.Sel("floatValue"))
	return rv
}
func (u SKUniform) SetFloatValue(value float32) {
	objc.Send[struct{}](u.ID, objc.Sel("setFloatValue:"), value)
}

// The receiver’s value as a SpriteKit texture.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/textureValue
func (u SKUniform) TextureValue() ISKTexture {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("textureValue"))
	return SKTextureFromID(objc.ID(rv))
}
func (u SKUniform) SetTextureValue(value ISKTexture) {
	objc.Send[struct{}](u.ID, objc.Sel("setTextureValue:"), value)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/matrixFloat2x2Value
func (u SKUniform) MatrixFloat2x2Value() [2][2]float32 {
	rv := objc.Send[[2][2]float32](u.ID, objc.Sel("matrixFloat2x2Value"))
	return [2][2]float32(rv)
}
func (u SKUniform) SetMatrixFloat2x2Value(value [2][2]float32) {
	objc.Send[struct{}](u.ID, objc.Sel("setMatrixFloat2x2Value:"), value)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/matrixFloat3x3Value
func (u SKUniform) MatrixFloat3x3Value() [3][4]float32 {
	rv := objc.Send[[3][4]float32](u.ID, objc.Sel("matrixFloat3x3Value"))
	return [3][4]float32(rv)
}
func (u SKUniform) SetMatrixFloat3x3Value(value [3][4]float32) {
	objc.Send[struct{}](u.ID, objc.Sel("setMatrixFloat3x3Value:"), value)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/matrixFloat4x4Value
func (u SKUniform) MatrixFloat4x4Value() [4][4]float32 {
	rv := objc.Send[[4][4]float32](u.ID, objc.Sel("matrixFloat4x4Value"))
	return [4][4]float32(rv)
}
func (u SKUniform) SetMatrixFloat4x4Value(value [4][4]float32) {
	objc.Send[struct{}](u.ID, objc.Sel("setMatrixFloat4x4Value:"), value)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/vectorFloat2Value
func (u SKUniform) VectorFloat2Value() [2]float32 {
	rv := objc.Send[[2]float32](u.ID, objc.Sel("vectorFloat2Value"))
	return [2]float32(rv)
}
func (u SKUniform) SetVectorFloat2Value(value [2]float32) {
	objc.Send[struct{}](u.ID, objc.Sel("setVectorFloat2Value:"), value)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/vectorFloat3Value
func (u SKUniform) VectorFloat3Value() Vector_float3 {
	rv := objc.Send[Vector_float3](u.ID, objc.Sel("vectorFloat3Value"))
	return Vector_float3(rv)
}
func (u SKUniform) SetVectorFloat3Value(value Vector_float3) {
	objc.Send[struct{}](u.ID, objc.Sel("setVectorFloat3Value:"), value)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniform/vectorFloat4Value
func (u SKUniform) VectorFloat4Value() [4]float32 {
	rv := objc.Send[[4]float32](u.ID, objc.Sel("vectorFloat4Value"))
	return [4]float32(rv)
}
func (u SKUniform) SetVectorFloat4Value(value [4]float32) {
	objc.Send[struct{}](u.ID, objc.Sel("setVectorFloat4Value:"), value)
}

// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKShader] class.
var (
	_SKShaderClass     SKShaderClass
	_SKShaderClassOnce sync.Once
)

func getSKShaderClass() SKShaderClass {
	_SKShaderClassOnce.Do(func() {
		_SKShaderClass = SKShaderClass{class: objc.GetClass("SKShader")}
	})
	return _SKShaderClass
}

// GetSKShaderClass returns the class object for SKShader.
func GetSKShaderClass() SKShaderClass {
	return getSKShaderClass()
}

type SKShaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKShaderClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKShaderClass) Alloc() SKShader {
	rv := objc.Send[SKShader](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object that allows you to apply a custom fragment shader.
//
// # Overview
//
// An [SKShader] object holds a custom OpenGL ES fragment shader. Shader
// objects are used to customize the drawing behavior of many different kinds
// of nodes in SpriteKit.
//
// To use a custom shader, create an [SKShader] object and provide the source
// for the custom fragment shader. If your shader needs to provide uniform
// data to the shader, create one or more [SKUniform] objects and associate
// them with the shader object. If your shader needs to provide per-node data
// to the shader, create one or more [SKAttribute] objects and associate them
// with the relevant nodes. Then, assign the shader object to the
// [SKSpriteNode.Shader] property of any sprites that need the custom
// behavior.
//
// Compiling a shader and the uniform data associated with it can be
// expensive. Because of this, you should:
//
// - Initialize shader objects when your game launches, not while the game is
// running. - Avoid changing the shader’s source or changing the list of
// uniforms or attributes while your game is running. Either of these things
// recompiles the shader. - Share shader objects whenever possible. If
// multiple sprites need the same behavior, create one shader object and
// associate it with every sprite that needs that behavior. Do not create a
// separate shader for each sprite.
//
// # Creating a Shader
//
//   - [SKShader.InitWithSourceUniforms]: Initializes a new shader object using the specified source and uniform data.
//   - [SKShader.InitWithSource]: Initializes a new shader object using the specified source code.
//
// # Providing Uniform Data to a Shader
//
//   - [SKShader.AddUniform]: Adds a uniform to the shader.
//   - [SKShader.RemoveUniformNamed]: Removes a uniform from the shader.
//   - [SKShader.Uniforms]: The list of uniforms associated with the shader.
//   - [SKShader.SetUniforms]
//   - [SKShader.UniformNamed]: Returns the uniform object corresponding to a particular uniform variable.
//
// # Providing Attribute Data to a Shader
//
//   - [SKShader.Attributes]: The list of attributes associated with the shader.
//   - [SKShader.SetAttributes]
//
// # Accessing or Setting a Shader’s Source Code
//
//   - [SKShader.Source]: The source code for the shader.
//   - [SKShader.SetSource]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader
type SKShader struct {
	objectivec.Object
}

// SKShaderFromID constructs a [SKShader] from an objc.ID.
//
// An object that allows you to apply a custom fragment shader.
func SKShaderFromID(id objc.ID) SKShader {
	return SKShader{objectivec.Object{ID: id}}
}

// NOTE: SKShader adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKShader] class.
//
// # Creating a Shader
//
//   - [ISKShader.InitWithSourceUniforms]: Initializes a new shader object using the specified source and uniform data.
//   - [ISKShader.InitWithSource]: Initializes a new shader object using the specified source code.
//
// # Providing Uniform Data to a Shader
//
//   - [ISKShader.AddUniform]: Adds a uniform to the shader.
//   - [ISKShader.RemoveUniformNamed]: Removes a uniform from the shader.
//   - [ISKShader.Uniforms]: The list of uniforms associated with the shader.
//   - [ISKShader.SetUniforms]
//   - [ISKShader.UniformNamed]: Returns the uniform object corresponding to a particular uniform variable.
//
// # Providing Attribute Data to a Shader
//
//   - [ISKShader.Attributes]: The list of attributes associated with the shader.
//   - [ISKShader.SetAttributes]
//
// # Accessing or Setting a Shader’s Source Code
//
//   - [ISKShader.Source]: The source code for the shader.
//   - [ISKShader.SetSource]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader
type ISKShader interface {
	objectivec.IObject

	// Topic: Creating a Shader

	// Initializes a new shader object using the specified source and uniform data.
	InitWithSourceUniforms(source string, uniforms []SKUniform) SKShader
	// Initializes a new shader object using the specified source code.
	InitWithSource(source string) SKShader

	// Topic: Providing Uniform Data to a Shader

	// Adds a uniform to the shader.
	AddUniform(uniform ISKUniform)
	// Removes a uniform from the shader.
	RemoveUniformNamed(name string)
	// The list of uniforms associated with the shader.
	Uniforms() []SKUniform
	SetUniforms(value []SKUniform)
	// Returns the uniform object corresponding to a particular uniform variable.
	UniformNamed(name string) ISKUniform

	// Topic: Providing Attribute Data to a Shader

	// The list of attributes associated with the shader.
	Attributes() []SKAttribute
	SetAttributes(value []SKAttribute)

	// Topic: Accessing or Setting a Shader’s Source Code

	// The source code for the shader.
	Source() string
	SetSource(value string)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s SKShader) Init() SKShader {
	rv := objc.Send[SKShader](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SKShader) Autorelease() SKShader {
	rv := objc.Send[SKShader](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKShader creates a new SKShader instance.
func NewSKShader() SKShader {
	class := getSKShaderClass()
	rv := objc.Send[SKShader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new shader object by loading the source for a fragment shader
// from a file stored in the app’s bundle.
//
// name: The name of the fragment shader to load. The file must be present in your
// app bundle with the same name and a `XCUIElementTypeFsh` file extension.
//
// # Return Value
//
// A newly initialized shader object whose initial source is loaded from the
// shader file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/init(fileNamed:)
func NewShaderWithFileNamed(name string) SKShader {
	rv := objc.Send[objc.ID](objc.ID(getSKShaderClass().class), objc.Sel("shaderWithFileNamed:"), objc.String(name))
	return SKShaderFromID(rv)
}

// Initializes a new shader object using the specified source code.
//
// source: A string that holds the initial source for the shader.
//
// # Return Value
//
// An initialized shader object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/init(source:)
func NewShaderWithSource(source string) SKShader {
	instance := getSKShaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), objc.String(source))
	return SKShaderFromID(rv)
}

// Initializes a new shader object using the specified source and uniform
// data.
//
// source: A string that holds the initial source for the shader.
//
// uniforms: A list of uniforms to add to the shader object.
//
// # Return Value
//
// An initialized shader object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/init(source:uniforms:)
func NewShaderWithSourceUniforms(source string, uniforms []SKUniform) SKShader {
	instance := getSKShaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:uniforms:"), objc.String(source), objectivec.IObjectSliceToNSArray(uniforms))
	return SKShaderFromID(rv)
}

// Initializes a new shader object using the specified source and uniform
// data.
//
// source: A string that holds the initial source for the shader.
//
// uniforms: A list of uniforms to add to the shader object.
//
// # Return Value
//
// An initialized shader object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/init(source:uniforms:)
func (s SKShader) InitWithSourceUniforms(source string, uniforms []SKUniform) SKShader {
	rv := objc.Send[SKShader](s.ID, objc.Sel("initWithSource:uniforms:"), objc.String(source), objectivec.IObjectSliceToNSArray(uniforms))
	return rv
}

// Initializes a new shader object using the specified source code.
//
// source: A string that holds the initial source for the shader.
//
// # Return Value
//
// An initialized shader object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/init(source:)
func (s SKShader) InitWithSource(source string) SKShader {
	rv := objc.Send[SKShader](s.ID, objc.Sel("initWithSource:"), objc.String(source))
	return rv
}

// Adds a uniform to the shader.
//
// uniform: The new uniform object to add. The uniform object’s name must not already
// be in use by another uniform attached to the shader.
//
// # Discussion
//
// The uniform variable is automatically accessible inside your shader; do not
// add a declaration for it in your shader’s source code. The uniform must
// be accessed in the fragment shader.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/addUniform(_:)
func (s SKShader) AddUniform(uniform ISKUniform) {
	objc.Send[objc.ID](s.ID, objc.Sel("addUniform:"), uniform)
}

// Removes a uniform from the shader.
//
// name: The name of the uniform to remove.
//
// # Discussion
//
// If a uniform with that name does not exist in the shader, nothing happens.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/removeUniformNamed(_:)
func (s SKShader) RemoveUniformNamed(name string) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeUniformNamed:"), objc.String(name))
}

// Returns the uniform object corresponding to a particular uniform variable.
//
// name: The name of the uniform to search for.
//
// # Return Value
//
// The uniform object corresponding to the name, or `nil` if that uniform
// cannot be found.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/uniformNamed(_:)
func (s SKShader) UniformNamed(name string) ISKUniform {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("uniformNamed:"), objc.String(name))
	return SKUniformFromID(rv)
}
func (s SKShader) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The list of uniforms associated with the shader.
//
// # Discussion
//
// This property is not read-only, so you can also use it to provide all of
// the uniforms in a single operation. Each of the uniforms should be uniquely
// named.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/uniforms
func (s SKShader) Uniforms() []SKUniform {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("uniforms"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKUniform {
		return SKUniformFromID(id)
	})
}
func (s SKShader) SetUniforms(value []SKUniform) {
	objc.Send[struct{}](s.ID, objc.Sel("setUniforms:"), objectivec.IObjectSliceToNSArray(value))
}

// The list of attributes associated with the shader.
//
// # Discussion
//
// This property is not read-only, so you can also use it to provide all of
// the attributes in a single operation. Each of the attributes should be
// uniquely named.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/attributes
func (s SKShader) Attributes() []SKAttribute {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("attributes"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKAttribute {
		return SKAttributeFromID(id)
	})
}
func (s SKShader) SetAttributes(value []SKAttribute) {
	objc.Send[struct{}](s.ID, objc.Sel("setAttributes:"), objectivec.IObjectSliceToNSArray(value))
}

// The source code for the shader.
//
// # Discussion
//
// The source code for a shader object can be updated at runtime. However,
// recompiling the fragment shader can be an expensive operation.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShader/source
func (s SKShader) Source() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("source"))
	return foundation.NSStringFromID(rv).String()
}
func (s SKShader) SetSource(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setSource:"), objc.String(value))
}

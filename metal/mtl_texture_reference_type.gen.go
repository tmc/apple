// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTLTextureReferenceType] class.
var (
	_MTLTextureReferenceTypeClass     MTLTextureReferenceTypeClass
	_MTLTextureReferenceTypeClassOnce sync.Once
)

func getMTLTextureReferenceTypeClass() MTLTextureReferenceTypeClass {
	_MTLTextureReferenceTypeClassOnce.Do(func() {
		_MTLTextureReferenceTypeClass = MTLTextureReferenceTypeClass{class: objc.GetClass("MTLTextureReferenceType")}
	})
	return _MTLTextureReferenceTypeClass
}

// GetMTLTextureReferenceTypeClass returns the class object for MTLTextureReferenceType.
func GetMTLTextureReferenceTypeClass() MTLTextureReferenceTypeClass {
	return getMTLTextureReferenceTypeClass()
}

type MTLTextureReferenceTypeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLTextureReferenceTypeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTextureReferenceTypeClass) Alloc() MTLTextureReferenceType {
	rv := objc.Send[MTLTextureReferenceType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a texture.
//
// # Describing the texture
//
//   - [MTLTextureReferenceType.TextureType]: The texture type of the texture.
//   - [MTLTextureReferenceType.TextureDataType]: The data type of the texture.
//   - [MTLTextureReferenceType.Access]: The texture’s read/write access to the argument.
//   - [MTLTextureReferenceType.IsDepthTexture]: A Boolean value that indicates whether the texture is a depth texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType
type MTLTextureReferenceType struct {
	MTLType
}

// MTLTextureReferenceTypeFromID constructs a [MTLTextureReferenceType] from an objc.ID.
//
// A description of a texture.
func MTLTextureReferenceTypeFromID(id objc.ID) MTLTextureReferenceType {
	return MTLTextureReferenceType{MTLType: MTLTypeFromID(id)}
}
// NOTE: MTLTextureReferenceType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTextureReferenceType] class.
//
// # Describing the texture
//
//   - [IMTLTextureReferenceType.TextureType]: The texture type of the texture.
//   - [IMTLTextureReferenceType.TextureDataType]: The data type of the texture.
//   - [IMTLTextureReferenceType.Access]: The texture’s read/write access to the argument.
//   - [IMTLTextureReferenceType.IsDepthTexture]: A Boolean value that indicates whether the texture is a depth texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType
type IMTLTextureReferenceType interface {
	IMTLType

	// Topic: Describing the texture

	// The texture type of the texture.
	TextureType() MTLTextureType
	// The data type of the texture.
	TextureDataType() MTLDataType
	// The texture’s read/write access to the argument.
	Access() MTLBindingAccess
	// A Boolean value that indicates whether the texture is a depth texture.
	IsDepthTexture() bool
}

// Init initializes the instance.
func (t MTLTextureReferenceType) Init() MTLTextureReferenceType {
	rv := objc.Send[MTLTextureReferenceType](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTextureReferenceType) Autorelease() MTLTextureReferenceType {
	rv := objc.Send[MTLTextureReferenceType](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTextureReferenceType creates a new MTLTextureReferenceType instance.
func NewMTLTextureReferenceType() MTLTextureReferenceType {
	class := getMTLTextureReferenceTypeClass()
	rv := objc.Send[MTLTextureReferenceType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The texture type of the texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType/textureType
func (t MTLTextureReferenceType) TextureType() MTLTextureType {
	rv := objc.Send[MTLTextureType](t.ID, objc.Sel("textureType"))
	return MTLTextureType(rv)
}
// The data type of the texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType/textureDataType
func (t MTLTextureReferenceType) TextureDataType() MTLDataType {
	rv := objc.Send[MTLDataType](t.ID, objc.Sel("textureDataType"))
	return MTLDataType(rv)
}
// The texture’s read/write access to the argument.
//
// # Discussion
// 
// This property indicates the type of access qualifiers (read-only,
// write-only, or read-write) used in the Metal shading language code. For
// information on possible values, see [MTLArgumentAccess].
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType/access
func (t MTLTextureReferenceType) Access() MTLBindingAccess {
	rv := objc.Send[MTLBindingAccess](t.ID, objc.Sel("access"))
	return MTLBindingAccess(rv)
}
// A Boolean value that indicates whether the texture is a depth texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureReferenceType/isDepthTexture
func (t MTLTextureReferenceType) IsDepthTexture() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isDepthTexture"))
	return rv
}


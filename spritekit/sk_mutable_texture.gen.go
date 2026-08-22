// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKMutableTexture] class.
var (
	_SKMutableTextureClass     SKMutableTextureClass
	_SKMutableTextureClassOnce sync.Once
)

func getSKMutableTextureClass() SKMutableTextureClass {
	_SKMutableTextureClassOnce.Do(func() {
		_SKMutableTextureClass = SKMutableTextureClass{class: objc.GetClass("SKMutableTexture")}
	})
	return _SKMutableTextureClass
}

// GetSKMutableTextureClass returns the class object for SKMutableTexture.
func GetSKMutableTextureClass() SKMutableTextureClass {
	return getSKMutableTextureClass()
}

type SKMutableTextureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKMutableTextureClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKMutableTextureClass) Alloc() SKMutableTexture {
	rv := objc.Send[SKMutableTexture](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A texture whose contents can be dynamically updated.
//
// # Overview
//
// Normally, SpriteKit textures ([SKTexture] objects) are static, meaning that
// once created, their contents cannot be changed. This is important because a
// static image can be more efficiently managed inside the graphics hardware.
// However, sometimes you need to be able to update the contents of a texture
// dynamically. In this case, you should use a mutable texture. Because there
// is a performance penalty for updating the texture’s contents, consider
// other options first. For example, you can render a texture in hardware
// using the [SKView.TextureFromNode] method and a node tree.
//
// To use this class, create a mutable texture using either one of its
// creation methods or those of its superclass. Then, when you need to update
// the mutable texture object’s contents, call the
// [SKMutableTexture.ModifyPixelDataWithBlock] method. Your block is called
// with the location of the texture in memory. Your block should update this
// texture and then return.
//
// # Creating an Empty Mutable Texture
//
//   - [SKMutableTexture.InitWithSizePixelFormat]: Initializes an empty texture with a specific size and format.
//   - [SKMutableTexture.InitWithSize]: Initializes an empty texture with a specific size.
//
// # Modifying a Mutable Texture’s Contents
//
//   - [SKMutableTexture.ModifyPixelDataWithBlock]: Modifies the contents of a mutable texture.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture
type SKMutableTexture struct {
	SKTexture
}

// SKMutableTextureFromID constructs a [SKMutableTexture] from an objc.ID.
//
// A texture whose contents can be dynamically updated.
func SKMutableTextureFromID(id objc.ID) SKMutableTexture {
	return SKMutableTexture{SKTexture: SKTextureFromID(id)}
}

// NOTE: SKMutableTexture adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKMutableTexture] class.
//
// # Creating an Empty Mutable Texture
//
//   - [ISKMutableTexture.InitWithSizePixelFormat]: Initializes an empty texture with a specific size and format.
//   - [ISKMutableTexture.InitWithSize]: Initializes an empty texture with a specific size.
//
// # Modifying a Mutable Texture’s Contents
//
//   - [ISKMutableTexture.ModifyPixelDataWithBlock]: Modifies the contents of a mutable texture.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture
type ISKMutableTexture interface {
	ISKTexture

	// Topic: Creating an Empty Mutable Texture

	// Initializes an empty texture with a specific size and format.
	InitWithSizePixelFormat(size corefoundation.CGSize, format int32) SKMutableTexture
	// Initializes an empty texture with a specific size.
	InitWithSize(size corefoundation.CGSize) SKMutableTexture

	// Topic: Modifying a Mutable Texture’s Contents

	// Modifies the contents of a mutable texture.
	ModifyPixelDataWithBlock(block UnsafePointerUintptrHandler)
}

// Init initializes the instance.
func (m SKMutableTexture) Init() SKMutableTexture {
	rv := objc.Send[SKMutableTexture](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m SKMutableTexture) Autorelease() SKMutableTexture {
	rv := objc.Send[SKMutableTexture](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKMutableTexture creates a new SKMutableTexture instance.
func NewSKMutableTexture() SKMutableTexture {
	class := getSKMutableTextureClass()
	rv := objc.Send[SKMutableTexture](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an empty texture with a specific size.
//
// size: The size of the texture, in pixels.
//
// # Return Value
//
// An empty mutable texture.
//
// # Discussion
//
// You must call the [SKMutableTexture.ModifyPixelDataWithBlock] method at
// least once before using this texture.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture/init(size:)
func NewMutableTextureWithSize(size corefoundation.CGSize) SKMutableTexture {
	instance := getSKMutableTextureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSize:"), size)
	return SKMutableTextureFromID(rv)
}

// Initializes an empty texture with a specific size and format.
//
// size: The size of the texture, in pixels.
//
// format: A Core Video format code. Three codes are supported:
// [kCVPixelFormatType_32RGBA], [kCVPixelFormatType_64RGBAHalf], and
// [kCVPixelFormatType_128RGBAFloat] for byte, half-float, and float
// components respectively.
//
// # Return Value
//
// An empty mutable texture.
//
// # Discussion
//
// You must call the [SKMutableTexture.ModifyPixelDataWithBlock] method at
// least once before using this texture.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture/init(size:pixelFormat:)
//
// [kCVPixelFormatType_128RGBAFloat]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_128RGBAFloat
// [kCVPixelFormatType_32RGBA]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32RGBA
// [kCVPixelFormatType_64RGBAHalf]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_64RGBAHalf
func NewMutableTextureWithSizePixelFormat(size corefoundation.CGSize, format int32) SKMutableTexture {
	instance := getSKMutableTextureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSize:pixelFormat:"), size, format)
	return SKMutableTextureFromID(rv)
}

// Initializes an empty texture with a specific size and format.
//
// size: The size of the texture, in pixels.
//
// format: A Core Video format code. Three codes are supported:
// [kCVPixelFormatType_32RGBA], [kCVPixelFormatType_64RGBAHalf], and
// [kCVPixelFormatType_128RGBAFloat] for byte, half-float, and float
// components respectively.
//
// # Return Value
//
// An empty mutable texture.
//
// # Discussion
//
// You must call the [SKMutableTexture.ModifyPixelDataWithBlock] method at
// least once before using this texture.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture/init(size:pixelFormat:)
//
// [kCVPixelFormatType_128RGBAFloat]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_128RGBAFloat
// [kCVPixelFormatType_32RGBA]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32RGBA
// [kCVPixelFormatType_64RGBAHalf]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_64RGBAHalf
func (m SKMutableTexture) InitWithSizePixelFormat(size corefoundation.CGSize, format int32) SKMutableTexture {
	rv := objc.Send[SKMutableTexture](m.ID, objc.Sel("initWithSize:pixelFormat:"), size, format)
	return rv
}

// Initializes an empty texture with a specific size.
//
// size: The size of the texture, in pixels.
//
// # Return Value
//
// An empty mutable texture.
//
// # Discussion
//
// You must call the [SKMutableTexture.ModifyPixelDataWithBlock] method at
// least once before using this texture.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture/init(size:)
func (m SKMutableTexture) InitWithSize(size corefoundation.CGSize) SKMutableTexture {
	rv := objc.Send[SKMutableTexture](m.ID, objc.Sel("initWithSize:"), size)
	return rv
}

// Modifies the contents of a mutable texture.
//
// block: A block to be called when the texture can be safely modified. The block
// takes the following parameters:
//
// pixelData: A pointer to the start of the current texture data.
// lengthInBytes: The length of the texture data in bytes.
//
// # Discussion
//
// The contents of the texture can be modified only at specific times when the
// graphics hardware permits it. When this method is called, it schedules a
// new background task to update the texture and then returns. Your block is
// called when the texture can be modified. Your block is called on an
// arbitrary queue. Your block should modify the texture’s contents and then
// return.
//
// The texture bytes are assumed to be stored as tightly packed 32 bpp, 8bpc
// (unsigned integer) RGBA pixel data. The color components you provide should
// have already been multiplied by the alpha value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture/modifyPixelData(_:)
func (m SKMutableTexture) ModifyPixelDataWithBlock(block UnsafePointerUintptrHandler) {
	_block0, _ := NewUnsafePointerUintptrBlock(block)
	objc.Send[objc.ID](m.ID, objc.Sel("modifyPixelDataWithBlock:"), _block0)
}

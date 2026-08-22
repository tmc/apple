// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"context"
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKTexture] class.
var (
	_SKTextureClass     SKTextureClass
	_SKTextureClassOnce sync.Once
)

func getSKTextureClass() SKTextureClass {
	_SKTextureClassOnce.Do(func() {
		_SKTextureClass = SKTextureClass{class: objc.GetClass("SKTexture")}
	})
	return _SKTextureClass
}

// GetSKTextureClass returns the class object for SKTexture.
func GetSKTextureClass() SKTextureClass {
	return getSKTextureClass()
}

type SKTextureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKTextureClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKTextureClass) Alloc() SKTexture {
	rv := objc.Send[SKTexture](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An image, decoded on the GPU, that can be used to render various SpriteKit
// objects.
//
// # Overview
//
// An [SKTexture] object is an image that can be applied to [SKSpriteNode] and
// [SKShapeNode] objects, particles created by an [SKEmitterNode] object, or
// tiles used in an [SKTileMapNode]. A texture object manages the texture data
// and graphics resources that are needed to render the image. Most texture
// objects are created from source images stored in your app bundle—your
// game’s artwork. Once created, a texture object’s contents are
// immutable. Multiple sprites can share the same texture object, sharing a
// single resource.
//
// # Deallocating a Texture
//
// After a texture is loaded into the graphics hardware memory, it stays in
// memory until the referencing [SKTexture] object is deleted. This means that
// between levels (or in a dynamic game), you may need to make sure a texture
// object is deleted. Delete a [SKTexture] object by removing any strong
// references to it, including:
//
// - All texture references from [SKSpriteNode] and [SKEffectNode] objects in
// your game - Any strong references to the texture in your own code - An
// [SKTextureAtlas] object that was used to create the texture object
//
// # Reading a Texture’s Size and Optional Source Location
//
//   - [SKTexture.Size]: Gets the size of the texture.
//   - [SKTexture.TextureRect]: Gets a rectangle that defines the portion of the texture used to render its image.
//
// # Configuring a Texture’s Behavior for Scaling
//
//   - [SKTexture.FilteringMode]: The filtering mode used when the size of a sprite drawn with the texture is not drawn at the texture’s native size.
//   - [SKTexture.SetFilteringMode]
//   - [SKTexture.UsesMipmaps]: A Boolean value that indicates whether the texture attempts to generate mipmaps.
//   - [SKTexture.SetUsesMipmaps]
//
// # Getting a Texture’s Underlying Image
//
//   - [SKTexture.CGImage]: Returns the texture’s image data as a Quartz 2D image.
//
// # Preloading a Texture for Performance
//
//   - [SKTexture.PreloadWithCompletionHandler]: Load texture data into memory, calling a completion handler after the task completes.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTexture
type SKTexture struct {
	objectivec.Object
}

// SKTextureFromID constructs a [SKTexture] from an objc.ID.
//
// An image, decoded on the GPU, that can be used to render various SpriteKit
// objects.
func SKTextureFromID(id objc.ID) SKTexture {
	return SKTexture{objectivec.Object{ID: id}}
}

// NOTE: SKTexture adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKTexture] class.
//
// # Reading a Texture’s Size and Optional Source Location
//
//   - [ISKTexture.Size]: Gets the size of the texture.
//   - [ISKTexture.TextureRect]: Gets a rectangle that defines the portion of the texture used to render its image.
//
// # Configuring a Texture’s Behavior for Scaling
//
//   - [ISKTexture.FilteringMode]: The filtering mode used when the size of a sprite drawn with the texture is not drawn at the texture’s native size.
//   - [ISKTexture.SetFilteringMode]
//   - [ISKTexture.UsesMipmaps]: A Boolean value that indicates whether the texture attempts to generate mipmaps.
//   - [ISKTexture.SetUsesMipmaps]
//
// # Getting a Texture’s Underlying Image
//
//   - [ISKTexture.CGImage]: Returns the texture’s image data as a Quartz 2D image.
//
// # Preloading a Texture for Performance
//
//   - [ISKTexture.PreloadWithCompletionHandler]: Load texture data into memory, calling a completion handler after the task completes.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTexture
type ISKTexture interface {
	objectivec.IObject

	// Topic: Reading a Texture’s Size and Optional Source Location

	// Gets the size of the texture.
	Size() corefoundation.CGSize
	// Gets a rectangle that defines the portion of the texture used to render its image.
	TextureRect() corefoundation.CGRect

	// Topic: Configuring a Texture’s Behavior for Scaling

	// The filtering mode used when the size of a sprite drawn with the texture is not drawn at the texture’s native size.
	FilteringMode() SKTextureFilteringMode
	SetFilteringMode(value SKTextureFilteringMode)
	// A Boolean value that indicates whether the texture attempts to generate mipmaps.
	UsesMipmaps() bool
	SetUsesMipmaps(value bool)

	// Topic: Getting a Texture’s Underlying Image

	// Returns the texture’s image data as a Quartz 2D image.
	CGImage() coregraphics.CGImageRef

	// Topic: Preloading a Texture for Performance

	// Load texture data into memory, calling a completion handler after the task completes.
	PreloadWithCompletionHandler(completionHandler VoidHandler)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t SKTexture) Init() SKTexture {
	rv := objc.Send[SKTexture](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SKTexture) Autorelease() SKTexture {
	rv := objc.Send[SKTexture](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTexture creates a new SKTexture instance.
func NewSKTexture() SKTexture {
	class := getSKTextureClass()
	rv := objc.Send[SKTexture](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Gets the size of the texture.
//
// # Return Value
//
// The dimensions of the texture, measured in points.
//
// # Discussion
//
// If the texture was created using an image file and that image file hasn’t
// been loaded, calling this method forces the texture data to be loaded from
// the file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTexture/size()
func (t SKTexture) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}

// Gets a rectangle that defines the portion of the texture used to render its
// image.
//
// # Return Value
//
// A rectangle in the unit coordinate space.
//
// # Discussion
//
// The default value is a rectangle that covers the entire texture `(0,0)` -
// `(1,1)`. You cannot set this value directly; to use only a portion of a
// texture, use the [init(rect:in:)] method to create a new texture.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTexture/textureRect()
//
// [init(rect:in:)]: https://developer.apple.com/documentation/SpriteKit/SKTexture/init(rect:in:)
func (t SKTexture) TextureRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("textureRect"))
	return corefoundation.CGRect(rv)
}

// Returns the texture’s image data as a Quartz 2D image.
//
// # Discussion
//
// The [SKTexture.CGImage] property returns the contents of a texture as a
// Quartz Image.
//
// As an example use, you can create an image from a portion of your scene and
// save it to disk by doing the following:
//
// - Use the [SKView.TextureFromNode] method to render the scene’s contents
// to a texture. - Call [SKTexture.CGImage] on the result. - Use
// [CGImageDestination] to write the [CGImage] out to disk.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTexture/cgImage()
//
// [CGImageDestination]: https://developer.apple.com/documentation/ImageIO/CGImageDestination
func (t SKTexture) CGImage() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](t.ID, objc.Sel("CGImage"))
	return coregraphics.CGImageRef(rv)
}

// Load texture data into memory, calling a completion handler after the task
// completes.
//
// completionHandler: A block called after the texture data is loaded.
//
// # Discussion
//
// SpriteKit creates a background task to load the texture data from the
// associated file, then returns control to your game. After the texture data
// is loaded, your completion handler is called. Typically, you use this
// method when you want to guarantee that a particular texture is in memory
// before accessing it.
//
// If you need to preload multiple textures at once, use the
// [SKTextureClass.PreloadTexturesWithCompletionHandler] method instead.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTexture/preload(completionHandler:)
func (t SKTexture) PreloadWithCompletionHandler(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](t.ID, objc.Sel("preloadWithCompletionHandler:"), _block0)
}
func (t SKTexture) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Load the data of multiple textures into memory.
//
// textures: An array of [SKTexture] objects.
//
// completionHandler: A block called after all of the textures are loaded.
//
// # Discussion
//
// SpriteKit creates a background task that loads the texture data for all of
// the textures in the array, then returns control to your game. Your
// completion handler is called after all of the textures are loaded.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTexture/preload(_:withCompletionHandler:)
func (_SKTextureClass SKTextureClass) PreloadTexturesWithCompletionHandler(textures []SKTexture, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_SKTextureClass.class), objc.Sel("preloadTextures:withCompletionHandler:"), textures, _block1)
}

// The filtering mode used when the size of a sprite drawn with the texture is
// not drawn at the texture’s native size.
//
// # Discussion
//
// The possible values for this property are listed in
// [SKTextureFilteringMode]. The default value is
// [SKTextureFilteringMode.linear] where each pixel is drawn by using a linear
// filter of multiple texels in the texture. The other option is
// [SKTextureFilteringMode.nearest] where each pixel is drawn using the
// nearest point in the texture.
//
// The figure below shows the effect of different filtering modes. The rabbit
// texture (original on left) has been scaled up five times. Node 1 has been
// scaled using [SKTextureFilteringMode.nearest] and node 2 has been scaled
// with [SKTextureFilteringMode.linear].
//
// [media-2668828]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTexture/filteringMode
//
// [SKTextureFilteringMode.linear]: https://developer.apple.com/documentation/SpriteKit/SKTextureFilteringMode/linear
// [SKTextureFilteringMode.nearest]: https://developer.apple.com/documentation/SpriteKit/SKTextureFilteringMode/nearest
// [SKTextureFilteringMode]: https://developer.apple.com/documentation/SpriteKit/SKTextureFilteringMode
func (t SKTexture) FilteringMode() SKTextureFilteringMode {
	rv := objc.Send[SKTextureFilteringMode](t.ID, objc.Sel("filteringMode"))
	return SKTextureFilteringMode(rv)
}
func (t SKTexture) SetFilteringMode(value SKTextureFilteringMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setFilteringMode:"), value)
}

// A Boolean value that indicates whether the texture attempts to generate
// mipmaps.
//
// # Discussion
//
// The default value is false. If you set this to true, Sprite Kit creates
// mipmaps for the texture when it prepares the texture for rendering. Mipmaps
// take up additional memory (usually one-third more) but can improve
// rendering quality and performance when the texture is reduced in size (such
// as when you reduce the scale of a sprite rendered using the texture).
//
// You can only request mipmaps if both of the texture’s dimensions are a
// power of two.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTexture/usesMipmaps
func (t SKTexture) UsesMipmaps() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesMipmaps"))
	return rv
}
func (t SKTexture) SetUsesMipmaps(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesMipmaps:"), value)
}

// Preload is a synchronous wrapper around [SKTexture.PreloadWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t SKTexture) Preload(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.PreloadWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

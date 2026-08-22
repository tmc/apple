// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKTextureAtlas] class.
var (
	_SKTextureAtlasClass     SKTextureAtlasClass
	_SKTextureAtlasClassOnce sync.Once
)

func getSKTextureAtlasClass() SKTextureAtlasClass {
	_SKTextureAtlasClassOnce.Do(func() {
		_SKTextureAtlasClass = SKTextureAtlasClass{class: objc.GetClass("SKTextureAtlas")}
	})
	return _SKTextureAtlasClass
}

// GetSKTextureAtlasClass returns the class object for SKTextureAtlas.
func GetSKTextureAtlasClass() SKTextureAtlasClass {
	return getSKTextureAtlasClass()
}

type SKTextureAtlasClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKTextureAtlasClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKTextureAtlasClass) Alloc() SKTextureAtlas {
	rv := objc.Send[SKTextureAtlas](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A collection of textures optimized for storage and drawing performance.
//
// # Overview
//
// An [SKTextureAtlas] is a collection of textures that were either created
// from an `XCUIElementTypeAtlas` folder in the app bundle, or created at
// runtime. Texture atlases improve memory usage and rendering performance by
// reducing draw calls. Whenever you have textures that are always used
// together, store them in an atlas for best results.
//
// SpriteKit implicitly loads an atlas when one of the atlas’s textures is
// accessed. Use [SKTextureAtlas.TextureNamed] when you want to explicitly
// access a texture atlas’s contents.
//
// The preferred place to create a texture atlas is within an asset catalog
// (see [Creating a Sprite Atlas]), but you can also put your source textures
// in an `XCUIElementTypeAtlas` folder in the app bundle.
//
// # Accessing Textures
//
//   - [SKTextureAtlas.TextureNamed]: Creates a texture from data stored in the texture atlas.
//
// # Preloading Textures
//
//   - [SKTextureAtlas.PreloadWithCompletionHandler]: Loads an atlas object’s textures into memory, calling a completion handler after the task completes.
//
// # Reading Source Image Filenames
//
//   - [SKTextureAtlas.TextureNames]: The names of the texture images stored in the atlas.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas
//
// [Creating a Sprite Atlas]: https://developer.apple.com/documentation/SpriteKit/about-texture-atlases#Creating-a-Sprite-Atlas
type SKTextureAtlas struct {
	objectivec.Object
}

// SKTextureAtlasFromID constructs a [SKTextureAtlas] from an objc.ID.
//
// A collection of textures optimized for storage and drawing performance.
func SKTextureAtlasFromID(id objc.ID) SKTextureAtlas {
	return SKTextureAtlas{objectivec.Object{ID: id}}
}

// NOTE: SKTextureAtlas adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKTextureAtlas] class.
//
// # Accessing Textures
//
//   - [ISKTextureAtlas.TextureNamed]: Creates a texture from data stored in the texture atlas.
//
// # Preloading Textures
//
//   - [ISKTextureAtlas.PreloadWithCompletionHandler]: Loads an atlas object’s textures into memory, calling a completion handler after the task completes.
//
// # Reading Source Image Filenames
//
//   - [ISKTextureAtlas.TextureNames]: The names of the texture images stored in the atlas.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas
type ISKTextureAtlas interface {
	objectivec.IObject

	// Topic: Accessing Textures

	// Creates a texture from data stored in the texture atlas.
	TextureNamed(name string) ISKTexture

	// Topic: Preloading Textures

	// Loads an atlas object’s textures into memory, calling a completion handler after the task completes.
	PreloadWithCompletionHandler(completionHandler VoidHandler)

	// Topic: Reading Source Image Filenames

	// The names of the texture images stored in the atlas.
	TextureNames() []string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t SKTextureAtlas) Init() SKTextureAtlas {
	rv := objc.Send[SKTextureAtlas](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SKTextureAtlas) Autorelease() SKTextureAtlas {
	rv := objc.Send[SKTextureAtlas](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTextureAtlas creates a new SKTextureAtlas instance.
func NewSKTextureAtlas() SKTextureAtlas {
	class := getSKTextureAtlasClass()
	rv := objc.Send[SKTextureAtlas](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a texture atlas from data stored in the app bundle.
//
// name: The name of the texture atlas, without the `XCUIElementTypeAtlas`
// extension.
//
// # Return Value
//
// A new texture atlas object.
//
// # Discussion
//
// If the texture atlas cannot be found, an exception is thrown.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/init(named:)
func NewTextureAtlasNamed(name string) SKTextureAtlas {
	rv := objc.Send[objc.ID](objc.ID(getSKTextureAtlasClass().class), objc.Sel("atlasNamed:"), objc.String(name))
	return SKTextureAtlasFromID(rv)
}

// Creates a texture atlas from a set of image files.
//
// properties: A dictionary that defines which textures are to be merged into the atlas.
//
// # Return Value
//
// A new texture atlas object.
//
// # Discussion
//
// Normally, Xcode creates texture atlases at compile time from the image
// files included in your project. These atlases are compiled and installed
// inside the app bundle. However, sometimes the assets needed to create a
// texture atlas are not available at compile time. For example, those assets
// might be procedurally generated or downloaded from the network. However,
// you still want the benefit of texture atlases to reduce the number of state
// changes required in the hardware. You can use this method to generate an
// atlas object at runtime. This is a potentially expensive operation best
// performed when your game loop is not running.
//
// The keys in the dictionary represent the names of the individual textures.
// The associated object for each key can be:
//
// - An [NSString] object that contains a file system path to a file that
// contains the texture - An [NSURL] object that contains a file system path
// to a file that contains the texture - A [UIImage] object - An [NSImage]
// object
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/init(dictionary:)
//
// [NSImage]: https://developer.apple.com/documentation/AppKit/NSImage
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
func NewTextureAtlasWithDictionary(properties foundation.INSDictionary) SKTextureAtlas {
	rv := objc.Send[objc.ID](objc.ID(getSKTextureAtlasClass().class), objc.Sel("atlasWithDictionary:"), properties)
	return SKTextureAtlasFromID(rv)
}

// Creates a texture from data stored in the texture atlas.
//
// name: The name of a texture stored in the atlas object.
//
// # Return Value
//
// The SpriteKit texture associated with the name. If the specified image does
// not exist in the atlas object, SpriteKit returns a placeholder texture
// image.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/textureNamed(_:)
func (t SKTextureAtlas) TextureNamed(name string) ISKTexture {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textureNamed:"), objc.String(name))
	return SKTextureFromID(rv)
}

// Loads an atlas object’s textures into memory, calling a completion
// handler after the task completes.
//
// completionHandler: A block called after the texture atlas is loaded.
//
// # Discussion
//
// SpriteKit creates a background task that loads the texture data from the
// atlas object. Then, SpriteKit returns control to your game. After the
// texture atlas is loaded, your completion handler is called.
//
// If you need to preload multiple texture atlas objects immediately, use the
// [SKTextureAtlasClass.PreloadTextureAtlasesWithCompletionHandler] method
// instead.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/preload(completionHandler:)
func (t SKTextureAtlas) PreloadWithCompletionHandler(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](t.ID, objc.Sel("preloadWithCompletionHandler:"), _block0)
}
func (t SKTextureAtlas) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Loads the textures of multiple atlas objects into memory, calling a
// completion handler after the task completes.
//
// textureAtlases: An array of [SKTextureAtlas] objects.
//
// completionHandler: A block called after all of the texture atlases are loaded.
//
// # Discussion
//
// SpriteKit creates a background task that loads the texture data from all of
// the specified atlas objects. Then, SpriteKit returns control to your game.
// After the atlas objects are loaded, your completion handler is called.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/preloadTextureAtlases(_:withCompletionHandler:)
func (_SKTextureAtlasClass SKTextureAtlasClass) PreloadTextureAtlasesWithCompletionHandler(textureAtlases []SKTextureAtlas, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_SKTextureAtlasClass.class), objc.Sel("preloadTextureAtlases:withCompletionHandler:"), textureAtlases, _block1)
}

// Loads the textures of multiple atlases into memory, calling a completion
// handler after the task completes.
//
// atlasNames: An array containing the atlas names to preload.
//
// completionHandler: A block called after all of the texture atlases are loaded.
//
// # Discussion
//
// SpriteKit creates a background task that loads the texture data from all of
// the specified atlas objects. Then, SpriteKit returns control to your game.
// After the atlases are loaded, your completion handler is called.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/preloadTextureAtlasesNamed(_:withCompletionHandler:)
func (_SKTextureAtlasClass SKTextureAtlasClass) PreloadTextureAtlasesNamedWithCompletionHandler(atlasNames []string, completionHandler SKTextureAtlasArrayErrorHandler) {
	_block1, _ := NewSKTextureAtlasArrayErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_SKTextureAtlasClass.class), objc.Sel("preloadTextureAtlasesNamed:withCompletionHandler:"), atlasNames, _block1)
}

// The names of the texture images stored in the atlas.
//
// # Discussion
//
// The property holds an array of [NSString] objects. Each string is the name
// of a texture stored in the atlas. The count of the array is the number of
// textures stored in the atlas.
//
// If the atlas is not currently loaded into memory, this method forces it to
// be loaded from the app bundle. Your game blocks until the atlas is loaded.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/textureNames
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (t SKTextureAtlas) TextureNames() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textureNames"))
	return objc.ConvertSliceToStrings(rv)
}

// Preload is a synchronous wrapper around [SKTextureAtlas.PreloadWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t SKTextureAtlas) Preload(ctx context.Context) error {
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

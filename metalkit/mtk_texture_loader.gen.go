// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTKTextureLoader] class.
var (
	_MTKTextureLoaderClass     MTKTextureLoaderClass
	_MTKTextureLoaderClassOnce sync.Once
)

func getMTKTextureLoaderClass() MTKTextureLoaderClass {
	_MTKTextureLoaderClassOnce.Do(func() {
		_MTKTextureLoaderClass = MTKTextureLoaderClass{class: objc.GetClass("MTKTextureLoader")}
	})
	return _MTKTextureLoaderClass
}

// GetMTKTextureLoaderClass returns the class object for MTKTextureLoader.
func GetMTKTextureLoaderClass() MTKTextureLoaderClass {
	return getMTKTextureLoaderClass()
}

type MTKTextureLoaderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTKTextureLoaderClass) Alloc() MTKTextureLoader {
	rv := objc.Send[MTKTextureLoader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that creates textures from existing data in common image formats.
//
// # Overview
// 
// Use the [MTKTextureLoader] class to create a Metal texture from existing
// image data.
// 
// This class supports common file formats, like PNG, JPEG, and TIFF. It also
// loads image data from KTX and PVR files, asset catalogs, Core Graphics
// images, and other sources. It infers the output texture format and pixel
// format from the image data.
// 
// You create textures synchronously or asynchronously using
// [MTKTextureLoader] methods that return [MTLTexture] instances. Pass options
// to these methods that customize the image-loading and texture-creation
// process.
// 
// First create an [MTKTextureLoader] instance, passing the device that it
// uses to create textures. Then use one of the texture loader’s methods to
// create a texture. The code example below synchronously creates a texture
// from data at a URL, using the default options:
// 
// If you use custom data formats, or change the image data at runtime, use
// [MTLTexture] methods instead. For more information, see [Creating and
// sampling textures].
//
// [Creating and sampling textures]: https://developer.apple.com/documentation/Metal/creating-and-sampling-textures
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// # Creating a Texture Loader
//
//   - [MTKTextureLoader.InitWithDevice]: Initializes a new texture loader object.
//   - [MTKTextureLoader.Device]: The device object that the texture loader uses to create textures.
//
// # Loading Textures from URLs
//
//   - [MTKTextureLoader.NewTextureWithContentsOfURLOptionsError]: Synchronously loads image data and creates a new Metal texture from a given URL.
//   - [MTKTextureLoader.NewTextureWithContentsOfURLOptionsCompletionHandler]: Asynchronously loads image data and creates a new Metal texture from a given URL.
//   - [MTKTextureLoader.NewTexturesWithContentsOfURLsOptionsError]: Synchronously loads image data and creates new Metal textures from the specified list of URLs.
//   - [MTKTextureLoader.NewTexturesWithContentsOfURLsOptionsCompletionHandler]: Asynchronously loads image data and creates new Metal textures from the specified list of URLs.
//
// # Loading Textures from Asset Catalogs
//
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorBundleOptionsError]: Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorBundleOptionsCompletionHandler]: Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//   - [MTKTextureLoader.NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler]: Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError]: Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog, using a specified display gamut.
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler]: Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//   - [MTKTextureLoader.NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler]: Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
//
// # Loading Textures from Core Graphics Images
//
//   - [MTKTextureLoader.NewTextureWithCGImageOptionsError]: Synchronously loads image data and creates a new Metal texture from a given bitmap image.
//   - [MTKTextureLoader.NewTextureWithCGImageOptionsCompletionHandler]: Asynchronously loads image data and creates a new Metal texture from a given bitmap image.
//
// # Loading Textures from In-Memory Data Representations
//
//   - [MTKTextureLoader.NewTextureWithDataOptionsError]: Synchronously creates a new Metal texture from an in-memory representation of the texture’s data.
//   - [MTKTextureLoader.NewTextureWithDataOptionsCompletionHandler]: Asynchronously creates a new Metal texture from an in-memory representation of the texture’s data.
//
// # Loading Textures from Model I/O Representations
//
//   - [MTKTextureLoader.NewTextureWithMDLTextureOptionsError]: Synchronously loads image data and creates a Metal texture from the specified Model I/O texture.
//   - [MTKTextureLoader.NewTextureWithMDLTextureOptionsCompletionHandler]: Asynchronously loads image data and creates a Metal texture from the specified Model I/O texture.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader
type MTKTextureLoader struct {
	objectivec.Object
}

// MTKTextureLoaderFromID constructs a [MTKTextureLoader] from an objc.ID.
//
// An object that creates textures from existing data in common image formats.
func MTKTextureLoaderFromID(id objc.ID) MTKTextureLoader {
	return MTKTextureLoader{objectivec.Object{ID: id}}
}
// NOTE: MTKTextureLoader adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTKTextureLoader] class.
//
// # Creating a Texture Loader
//
//   - [IMTKTextureLoader.InitWithDevice]: Initializes a new texture loader object.
//   - [IMTKTextureLoader.Device]: The device object that the texture loader uses to create textures.
//
// # Loading Textures from URLs
//
//   - [IMTKTextureLoader.NewTextureWithContentsOfURLOptionsError]: Synchronously loads image data and creates a new Metal texture from a given URL.
//   - [IMTKTextureLoader.NewTextureWithContentsOfURLOptionsCompletionHandler]: Asynchronously loads image data and creates a new Metal texture from a given URL.
//   - [IMTKTextureLoader.NewTexturesWithContentsOfURLsOptionsError]: Synchronously loads image data and creates new Metal textures from the specified list of URLs.
//   - [IMTKTextureLoader.NewTexturesWithContentsOfURLsOptionsCompletionHandler]: Asynchronously loads image data and creates new Metal textures from the specified list of URLs.
//
// # Loading Textures from Asset Catalogs
//
//   - [IMTKTextureLoader.NewTextureWithNameScaleFactorBundleOptionsError]: Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//   - [IMTKTextureLoader.NewTextureWithNameScaleFactorBundleOptionsCompletionHandler]: Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//   - [IMTKTextureLoader.NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler]: Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
//   - [IMTKTextureLoader.NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError]: Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog, using a specified display gamut.
//   - [IMTKTextureLoader.NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler]: Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//   - [IMTKTextureLoader.NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler]: Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
//
// # Loading Textures from Core Graphics Images
//
//   - [IMTKTextureLoader.NewTextureWithCGImageOptionsError]: Synchronously loads image data and creates a new Metal texture from a given bitmap image.
//   - [IMTKTextureLoader.NewTextureWithCGImageOptionsCompletionHandler]: Asynchronously loads image data and creates a new Metal texture from a given bitmap image.
//
// # Loading Textures from In-Memory Data Representations
//
//   - [IMTKTextureLoader.NewTextureWithDataOptionsError]: Synchronously creates a new Metal texture from an in-memory representation of the texture’s data.
//   - [IMTKTextureLoader.NewTextureWithDataOptionsCompletionHandler]: Asynchronously creates a new Metal texture from an in-memory representation of the texture’s data.
//
// # Loading Textures from Model I/O Representations
//
//   - [IMTKTextureLoader.NewTextureWithMDLTextureOptionsError]: Synchronously loads image data and creates a Metal texture from the specified Model I/O texture.
//   - [IMTKTextureLoader.NewTextureWithMDLTextureOptionsCompletionHandler]: Asynchronously loads image data and creates a Metal texture from the specified Model I/O texture.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader
type IMTKTextureLoader interface {
	objectivec.IObject

	// Topic: Creating a Texture Loader

	// Initializes a new texture loader object.
	InitWithDevice(device metal.MTLDevice) MTKTextureLoader
	// The device object that the texture loader uses to create textures.
	Device() metal.MTLDevice

	// Topic: Loading Textures from URLs

	// Synchronously loads image data and creates a new Metal texture from a given URL.
	NewTextureWithContentsOfURLOptionsError(URL foundation.INSURL, options foundation.INSDictionary) (metal.MTLTexture, error)
	// Asynchronously loads image data and creates a new Metal texture from a given URL.
	NewTextureWithContentsOfURLOptionsCompletionHandler(URL foundation.INSURL, options foundation.INSDictionary, completionHandler ErrorHandler)
	// Synchronously loads image data and creates new Metal textures from the specified list of URLs.
	NewTexturesWithContentsOfURLsOptionsError(URLs []foundation.NSURL, options foundation.INSDictionary) ([]objectivec.IObject, error)
	// Asynchronously loads image data and creates new Metal textures from the specified list of URLs.
	NewTexturesWithContentsOfURLsOptionsCompletionHandler(URLs []foundation.NSURL, options foundation.INSDictionary, completionHandler ErrorHandler)

	// Topic: Loading Textures from Asset Catalogs

	// Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
	NewTextureWithNameScaleFactorBundleOptionsError(name string, scaleFactor float64, bundle foundation.NSBundle, options foundation.INSDictionary) (metal.MTLTexture, error)
	// Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
	NewTextureWithNameScaleFactorBundleOptionsCompletionHandler(name string, scaleFactor float64, bundle foundation.NSBundle, options foundation.INSDictionary, completionHandler ErrorHandler)
	// Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
	NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler(names []string, scaleFactor float64, bundle foundation.NSBundle, options foundation.INSDictionary, completionHandler ErrorHandler)
	// Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog, using a specified display gamut.
	NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError(name string, scaleFactor float64, displayGamut appkit.NSDisplayGamut, bundle foundation.NSBundle, options foundation.INSDictionary) (metal.MTLTexture, error)
	// Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
	NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler(name string, scaleFactor float64, displayGamut appkit.NSDisplayGamut, bundle foundation.NSBundle, options foundation.INSDictionary, completionHandler ErrorHandler)
	// Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
	NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler(names []string, scaleFactor float64, displayGamut appkit.NSDisplayGamut, bundle foundation.NSBundle, options foundation.INSDictionary, completionHandler ErrorHandler)

	// Topic: Loading Textures from Core Graphics Images

	// Synchronously loads image data and creates a new Metal texture from a given bitmap image.
	NewTextureWithCGImageOptionsError(cgImage coregraphics.CGImageRef, options foundation.INSDictionary) (metal.MTLTexture, error)
	// Asynchronously loads image data and creates a new Metal texture from a given bitmap image.
	NewTextureWithCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options foundation.INSDictionary, completionHandler ErrorHandler)

	// Topic: Loading Textures from In-Memory Data Representations

	// Synchronously creates a new Metal texture from an in-memory representation of the texture’s data.
	NewTextureWithDataOptionsError(data foundation.INSData, options foundation.INSDictionary) (metal.MTLTexture, error)
	// Asynchronously creates a new Metal texture from an in-memory representation of the texture’s data.
	NewTextureWithDataOptionsCompletionHandler(data foundation.INSData, options foundation.INSDictionary, completionHandler ErrorHandler)

	// Topic: Loading Textures from Model I/O Representations

	// Synchronously loads image data and creates a Metal texture from the specified Model I/O texture.
	NewTextureWithMDLTextureOptionsError(texture objectivec.IObject, options foundation.INSDictionary) (metal.MTLTexture, error)
	// Asynchronously loads image data and creates a Metal texture from the specified Model I/O texture.
	NewTextureWithMDLTextureOptionsCompletionHandler(texture objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler)
}

// Init initializes the instance.
func (t MTKTextureLoader) Init() MTKTextureLoader {
	rv := objc.Send[MTKTextureLoader](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTKTextureLoader) Autorelease() MTKTextureLoader {
	rv := objc.Send[MTKTextureLoader](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKTextureLoader creates a new MTKTextureLoader instance.
func NewMTKTextureLoader() MTKTextureLoader {
	class := getMTKTextureLoaderClass()
	rv := objc.Send[MTKTextureLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new texture loader object.
//
// device: The Metal device to create Metal textures with.
//
// # Return Value
// 
// An initialized texture loader object.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/init(device:)
func NewTextureLoaderWithDevice(device metal.MTLDevice) MTKTextureLoader {
	instance := getMTKTextureLoaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MTKTextureLoaderFromID(rv)
}

// Initializes a new texture loader object.
//
// device: The Metal device to create Metal textures with.
//
// # Return Value
// 
// An initialized texture loader object.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/init(device:)
func (t MTKTextureLoader) InitWithDevice(device metal.MTLDevice) MTKTextureLoader {
	rv := objc.Send[MTKTextureLoader](t.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

// Synchronously loads image data and creates a new Metal texture from a given
// URL.
//
// URL: The URL of the file to load.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
//
// # Return Value
// 
// A fully loaded and initialized Metal texture, or `nil` if an error
// occurred.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(URL:options:)
func (t MTKTextureLoader) NewTextureWithContentsOfURLOptionsError(URL foundation.INSURL, options foundation.INSDictionary) (metal.MTLTexture, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithContentsOfURL:options:error:"), URL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return metal.MTLTextureObjectFromID(rv), nil

}

// Asynchronously loads image data and creates a new Metal texture from a
// given URL.
//
// URL: The URL of the file to load.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
//
// completionHandler: A block called when the texture has been loaded and fully initialized.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(URL:options:completionHandler:)
func (t MTKTextureLoader) NewTextureWithContentsOfURLOptionsCompletionHandler(URL foundation.INSURL, options foundation.INSDictionary, completionHandler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithContentsOfURL:options:completionHandler:"), URL, options, _block2)
}

// Synchronously loads image data and creates new Metal textures from the
// specified list of URLs.
//
// URLs: An array of URLs referencing files to load.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
//
// error: If all textures were fully loaded and initialized, this pointer is `nil` on
// output. If an error occurs while loading any of the specified URLs, this
// pointer refers to an [NSError] object describing the failure. (Which
// element in the [URLs] array the error corresponds to is undefined.)
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Return Value
// 
// An array of Metal textures, each corresponding to a URL listed in the
// [URLs] parameter. If an error occurs while loading a texture, the
// corresponding array element is an [NSNull] object.
//
// [NSNull]: https://developer.apple.com/documentation/Foundation/NSNull
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(URLs:options:error:)
func (t MTKTextureLoader) NewTexturesWithContentsOfURLsOptionsError(URLs []foundation.NSURL, options foundation.INSDictionary) ([]objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("newTexturesWithContentsOfURLs:options:error:"), objectivec.IObjectSliceToNSArray(URLs), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	}), nil

}

// Asynchronously loads image data and creates new Metal textures from the
// specified list of URLs.
//
// URLs: An array of URLs referencing files to load.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
//
// completionHandler: A block called after all URLs have been processed. See the
// [MTKTextureLoaderArrayCallback] signature to determine whether each texture
// has successfully loaded.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(URLs:options:completionHandler:)
func (t MTKTextureLoader) NewTexturesWithContentsOfURLsOptionsCompletionHandler(URLs []foundation.NSURL, options foundation.INSDictionary, completionHandler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](t.ID, objc.Sel("newTexturesWithContentsOfURLs:options:completionHandler:"), URLs, options, _block2)
}

// Synchronously loads image data and creates a Metal texture from the named
// texture asset in an asset catalog.
//
// name: The name of a texture in an asset catalog.
//
// scaleFactor: The scale factor of texture to request.
// 
// In iOS and tvOS, pass the [contentsScale] value of the view where you plan
// to display texture content.
// 
// In macOS, pass the [backingScaleFactor] value of the window where you plan
// to display texture content.
// //
// [backingScaleFactor]: https://developer.apple.com/documentation/AppKit/NSWindow/backingScaleFactor
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// bundle: The resource bundle containing the asset catalog to load textures from.
//
// options: A dictionary describing any additional texture loading steps. See [Texture
// Loading Options].
// 
// When using this method, the texture loader ignores the [generateMipmaps],
// [SRGB], [cubeLayout], and [origin] options.
// //
// [SRGB]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/SRGB
// [Texture Loading Options]: https://developer.apple.com/documentation/GLKit/texture-loading-options
// [cubeLayout]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/cubeLayout
// [generateMipmaps]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/generateMipmaps
// [origin]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/origin
//
// # Return Value
// 
// A fully loaded and initialized Metal texture, or `nil` if an error
// occurred.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:bundle:options:)
func (t MTKTextureLoader) NewTextureWithNameScaleFactorBundleOptionsError(name string, scaleFactor float64, bundle foundation.NSBundle, options foundation.INSDictionary) (metal.MTLTexture, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithName:scaleFactor:bundle:options:error:"), objc.String(name), scaleFactor, bundle, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return metal.MTLTextureObjectFromID(rv), nil

}

// Asynchronously loads image data and creates a Metal texture from the named
// texture asset in an asset catalog.
//
// name: The name of a texture in an asset catalog.
//
// scaleFactor: The scale factor of texture to request.
// 
// In iOS and tvOS, pass the [contentsScale] value of the view where you plan
// to display texture content.
// 
// In macOS, pass the [backingScaleFactor] value of the window where you plan
// to display texture content.
// //
// [backingScaleFactor]: https://developer.apple.com/documentation/AppKit/NSWindow/backingScaleFactor
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// bundle: The resource bundle containing the asset catalog to load the texture from.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
// 
// When using this method, the texture loader ignores the [generateMipmaps],
// [SRGB], [cubeLayout], and [origin] options.
// //
// [SRGB]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/SRGB
// [cubeLayout]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/cubeLayout
// [generateMipmaps]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/generateMipmaps
// [origin]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/origin
//
// completionHandler: A block called when the texture has been loaded and fully initialized.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:bundle:options:completionHandler:)
func (t MTKTextureLoader) NewTextureWithNameScaleFactorBundleOptionsCompletionHandler(name string, scaleFactor float64, bundle foundation.NSBundle, options foundation.INSDictionary, completionHandler ErrorHandler) {
_block4, _cleanup4 := NewErrorBlock(completionHandler)
	defer _cleanup4()
	objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithName:scaleFactor:bundle:options:completionHandler:"), objc.String(name), scaleFactor, bundle, options, _block4)
}

// Asynchronously loads image data and creates Metal textures from the
// specified list of named texture assets in an asset catalog.
//
// names: An array of strings, each the name of a texture in an asset catalog.
//
// scaleFactor: The scale factor of texture to request.
// 
// In iOS and tvOS, pass the [contentsScale] value of the view where you plan
// to display texture content.
// 
// In macOS, pass the [backingScaleFactor] value of the window where you plan
// to display texture content.
// //
// [backingScaleFactor]: https://developer.apple.com/documentation/AppKit/NSWindow/backingScaleFactor
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// bundle: The resource bundle containing the asset catalog to load textures from.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
// 
// When using this method, the texture loader ignores the [generateMipmaps],
// [SRGB], [cubeLayout], and [origin] options.
// //
// [SRGB]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/SRGB
// [cubeLayout]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/cubeLayout
// [generateMipmaps]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/generateMipmaps
// [origin]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/origin
//
// completionHandler: A block called after all assets have been processed. See the
// [MTKTextureLoaderArrayCallback] signature to determine whether each texture
// has successfully loaded.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(names:scaleFactor:bundle:options:completionHandler:)
func (t MTKTextureLoader) NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler(names []string, scaleFactor float64, bundle foundation.NSBundle, options foundation.INSDictionary, completionHandler ErrorHandler) {
_block4, _cleanup4 := NewErrorBlock(completionHandler)
	defer _cleanup4()
	objc.Send[objc.ID](t.ID, objc.Sel("newTexturesWithNames:scaleFactor:bundle:options:completionHandler:"), names, scaleFactor, bundle, options, _block4)
}

// Synchronously loads image data and creates a Metal texture from the named
// texture asset in an asset catalog, using a specified display gamut.
//
// name: The name of a texture in an asset catalog.
//
// scaleFactor: The scale factor of texture to request.
// 
// In iOS and tvOS, pass the [contentsScale] value of the view where you plan
// to display texture content.
// 
// In macOS, pass the [backingScaleFactor] value of the window where you plan
// to display texture content.
// //
// [backingScaleFactor]: https://developer.apple.com/documentation/AppKit/NSWindow/backingScaleFactor
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// displayGamut: The version of the texture based on the trait in Xcode.
// 
// To determine the appropriate parameter value, pass the widest
// [NSDisplayGamut] value that returns [true] when queried against the ``
// method of [NSWindow].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// bundle: The resource bundle containing the asset catalog to load textures from.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
// 
// When using this method, the texture loader ignores the [generateMipmaps],
// [SRGB], [cubeLayout], and [origin] options.
// //
// [SRGB]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/SRGB
// [cubeLayout]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/cubeLayout
// [generateMipmaps]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/generateMipmaps
// [origin]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/origin
//
// # Return Value
// 
// A fully loaded and initialized Metal texture, or `nil` if an error
// occurred.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:displayGamut:bundle:options:)
func (t MTKTextureLoader) NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError(name string, scaleFactor float64, displayGamut appkit.NSDisplayGamut, bundle foundation.NSBundle, options foundation.INSDictionary) (metal.MTLTexture, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithName:scaleFactor:displayGamut:bundle:options:error:"), objc.String(name), scaleFactor, displayGamut, bundle, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return metal.MTLTextureObjectFromID(rv), nil

}

// Asynchronously loads image data and creates a Metal texture from the named
// texture asset in an asset catalog.
//
// name: The name of a texture in an asset catalog.
//
// scaleFactor: The scale factor of texture to request.
// 
// In iOS and tvOS, pass the [contentsScale] value of the view where you plan
// to display texture content.
// 
// In macOS, pass the [backingScaleFactor] value of the window where you plan
// to display texture content.
// //
// [backingScaleFactor]: https://developer.apple.com/documentation/AppKit/NSWindow/backingScaleFactor
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// displayGamut: The version of the texture based on the trait in Xcode.
// 
// To determine the appropriate parameter value, pass the widest
// [NSDisplayGamut] value that returns [true] when queried against the ``
// method of [NSWindow].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// bundle: The resource bundle containing the asset catalog to load the texture from.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
// 
// When using this method, the texture loader ignores the [generateMipmaps],
// [SRGB], [cubeLayout], and [origin] options.
// //
// [SRGB]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/SRGB
// [cubeLayout]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/cubeLayout
// [generateMipmaps]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/generateMipmaps
// [origin]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/origin
//
// completionHandler: A block called when the texture has been loaded and fully initialized.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:displayGamut:bundle:options:completionHandler:)
func (t MTKTextureLoader) NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler(name string, scaleFactor float64, displayGamut appkit.NSDisplayGamut, bundle foundation.NSBundle, options foundation.INSDictionary, completionHandler ErrorHandler) {
_block5, _cleanup5 := NewErrorBlock(completionHandler)
	defer _cleanup5()
	objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithName:scaleFactor:displayGamut:bundle:options:completionHandler:"), objc.String(name), scaleFactor, displayGamut, bundle, options, _block5)
}

// Asynchronously loads image data and creates Metal textures from the
// specified list of named texture assets in an asset catalog.
//
// names: An array of strings, each the name of a texture in an asset catalog.
//
// scaleFactor: The scale factor of texture to request.
// 
// In iOS and tvOS, pass the [contentsScale] value of the view where you plan
// to display texture content.
// 
// In macOS, pass the [backingScaleFactor] value of the window where you plan
// to display texture content.
// //
// [backingScaleFactor]: https://developer.apple.com/documentation/AppKit/NSWindow/backingScaleFactor
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// displayGamut: The version of the texture based on the trait in Xcode.
// 
// To determine the appropriate parameter value, pass the widest
// [NSDisplayGamut] value that returns [true] when queried against the ``
// method of [NSWindow].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// bundle: The resource bundle containing the asset catalog to load textures from.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
// 
// When using this method, the texture loader ignores the [generateMipmaps],
// [SRGB], [cubeLayout], and [origin] options.
// //
// [SRGB]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/SRGB
// [cubeLayout]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/cubeLayout
// [generateMipmaps]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/generateMipmaps
// [origin]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option/origin
//
// completionHandler: A block called after all assets have been processed. See the
// [MTKTextureLoaderArrayCallback] signature to determine whether each texture
// has successfully loaded.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(names:scaleFactor:displayGamut:bundle:options:completionHandler:)
func (t MTKTextureLoader) NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler(names []string, scaleFactor float64, displayGamut appkit.NSDisplayGamut, bundle foundation.NSBundle, options foundation.INSDictionary, completionHandler ErrorHandler) {
_block5, _cleanup5 := NewErrorBlock(completionHandler)
	defer _cleanup5()
	objc.Send[objc.ID](t.ID, objc.Sel("newTexturesWithNames:scaleFactor:displayGamut:bundle:options:completionHandler:"), names, scaleFactor, displayGamut, bundle, options, _block5)
}

// Synchronously loads image data and creates a new Metal texture from a given
// bitmap image.
//
// cgImage: The [CGImage] from which to load image data.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
//
// # Return Value
// 
// A fully loaded and initialized Metal texture, or `nil` if an error
// occurred.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(cgImage:options:)
func (t MTKTextureLoader) NewTextureWithCGImageOptionsError(cgImage coregraphics.CGImageRef, options foundation.INSDictionary) (metal.MTLTexture, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithCGImage:options:error:"), cgImage, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return metal.MTLTextureObjectFromID(rv), nil

}

// Asynchronously loads image data and creates a new Metal texture from a
// given bitmap image.
//
// cgImage: The [CGImage] from which to load image data.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
//
// completionHandler: A block called when the texture has been loaded and fully initialized.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(cgImage:options:completionHandler:)
func (t MTKTextureLoader) NewTextureWithCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options foundation.INSDictionary, completionHandler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithCGImage:options:completionHandler:"), cgImage, options, _block2)
}

// Synchronously creates a new Metal texture from an in-memory representation
// of the texture’s data.
//
// data: The [NSData] object containing image data.
// //
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
//
// # Return Value
// 
// A fully loaded and initialized Metal texture, or `nil` if an error
// occurred.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(data:options:)
func (t MTKTextureLoader) NewTextureWithDataOptionsError(data foundation.INSData, options foundation.INSDictionary) (metal.MTLTexture, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithData:options:error:"), data, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return metal.MTLTextureObjectFromID(rv), nil

}

// Asynchronously creates a new Metal texture from an in-memory representation
// of the texture’s data.
//
// data: The [NSData] object containing image data.
// //
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
//
// completionHandler: A block called when the texture has been loaded and fully initialized.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(data:options:completionHandler:)
func (t MTKTextureLoader) NewTextureWithDataOptionsCompletionHandler(data foundation.INSData, options foundation.INSDictionary, completionHandler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithData:options:completionHandler:"), data, options, _block2)
}

// Synchronously loads image data and creates a Metal texture from the
// specified Model I/O texture.
//
// texture: A Model I/O texture object containing image data from which to create the
// texture.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
//
// # Return Value
// 
// A fully loaded and initialized Metal texture, or `nil` if an error
// occurred.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(texture:options:)
func (t MTKTextureLoader) NewTextureWithMDLTextureOptionsError(texture objectivec.IObject, options foundation.INSDictionary) (metal.MTLTexture, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithMDLTexture:options:error:"), texture, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return metal.MTLTextureObjectFromID(rv), nil

}

// Asynchronously loads image data and creates a Metal texture from the
// specified Model I/O texture.
//
// texture: A Model I/O texture object containing image data from which to create the
// texture.
//
// options: A dictionary describing any additional texture loading steps. See `Texture
// Loading Options`.
//
// completionHandler: A block called when the texture has been loaded and fully initialized.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(texture:options:completionHandler:)
func (t MTKTextureLoader) NewTextureWithMDLTextureOptionsCompletionHandler(texture objectivec.IObject, options foundation.INSDictionary, completionHandler ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](t.ID, objc.Sel("newTextureWithMDLTexture:options:completionHandler:"), texture, options, _block2)
}

// The device object that the texture loader uses to create textures.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/device
func (t MTKTextureLoader) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}


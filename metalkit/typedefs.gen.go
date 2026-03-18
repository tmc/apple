// Code generated from Apple documentation. DO NOT EDIT.

package metalkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objectivec"
)

// MTKModelError is constants used to declare Model Errors.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKModelError
type MTKModelError = string

// MTKTextureLoaderArrayCallback is the signature for the block executed after an asynchronous loading operation for multiple textures has completed.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/ArrayCallback
type MTKTextureLoaderArrayCallback = func([]objectivec.IObject, foundation.NSError)

// MTKTextureLoaderCallback is the signature for the block executed after an asynchronous loading operation for a single texture has completed.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Callback
type MTKTextureLoaderCallback = func(metal.MTLTexture, foundation.NSError)

// MTKTextureLoaderCubeLayout is options for specifying how cube texture data is arranged in the source image.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/CubeLayout
type MTKTextureLoaderCubeLayout = string

// MTKTextureLoaderError is errors returned by the texture loader.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Error
type MTKTextureLoaderError = string

// MTKTextureLoaderOption is keys and values used to specify loading options.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Option
type MTKTextureLoaderOption = string

// MTKTextureLoaderOrigin is options for specifying when to flip the pixel coordinates of the texture.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/Origin
type MTKTextureLoaderOrigin = string


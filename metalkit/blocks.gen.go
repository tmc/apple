// Code generated from Apple documentation. DO NOT EDIT.

package metalkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ErrorHandler handles A block called when the texture has been loaded and fully initialized.
//
// Used by:
//   - [MTKTextureLoader.NewTextureWithCGImageOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithContentsOfURLOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithDataOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithMDLTextureOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorBundleOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTexturesWithContentsOfURLsOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler]
type ErrorHandler = func()

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTKTextureLoader.NewTextureWithCGImageOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithContentsOfURLOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithDataOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithMDLTextureOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorBundleOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTexturesWithContentsOfURLsOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// MTKTextureLoaderArrayCallback handles The signature for the block executed after an asynchronous loading operation for multiple textures has completed.

// NewMTKTextureLoaderArrayCallbackBlock wraps a Go [MTKTextureLoaderArrayCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTKTextureLoaderArrayCallbackBlock(handler MTKTextureLoaderArrayCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive []objectivec.IObject, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTKTextureLoaderCallback handles The signature for the block executed after an asynchronous loading operation for a single texture has completed.

// NewMTKTextureLoaderCallbackBlock wraps a Go [MTKTextureLoaderCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTKTextureLoaderCallbackBlock(handler MTKTextureLoaderCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive metal.MTLTexture, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// Code generated from Apple documentation. DO NOT EDIT.

package metalkit

import (
	"github.com/tmc/apple/objc"
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
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}


// Code generated from Apple documentation. DO NOT EDIT.

package metalkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// ErrorHandler handles A block called after all URLs have been processed.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTKTextureLoader.NewTexturesWithContentsOfURLsOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTKTextureLoader.NewTexturesWithContentsOfURLsOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// MTKTextureLoaderArrayCallback handles The signature for the block executed after an asynchronous loading operation for multiple textures has completed.

// MTKTextureLoaderCallback handles The signature for the block executed after an asynchronous loading operation for a single texture has completed.

// NewMTKTextureLoaderCallbackBlock wraps a Go [MTKTextureLoaderCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTKTextureLoaderCallbackBlock(handler MTKTextureLoaderCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 foundation.NSError) {
		var primitive metal.MTLTexture
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = metal.MTLTextureObjectFromID(primitiveID)
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLTextureErrorHandler handles A block called when the texture has been loaded and fully initialized.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTKTextureLoader.NewTextureWithCGImageOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithContentsOfURLOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithDataOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithMDLTextureOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorBundleOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler]
type MTLTextureErrorHandler = func(metal.MTLTexture, error)

// NewMTLTextureErrorBlock wraps a Go [MTLTextureErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTKTextureLoader.NewTextureWithCGImageOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithContentsOfURLOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithDataOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithMDLTextureOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorBundleOptionsCompletionHandler]
//   - [MTKTextureLoader.NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler]
func NewMTLTextureErrorBlock(handler MTLTextureErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result metal.MTLTexture
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = metal.MTLTextureObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

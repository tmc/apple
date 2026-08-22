// Code generated from Apple documentation. DO NOT EDIT.

package glkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// GLKTextureLoaderCallback handles Signature for the block executed after an asynchronous texture loading operation completes.

// NewGLKTextureLoaderCallbackBlock wraps a Go [GLKTextureLoaderCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGLKTextureLoaderCallbackBlock(handler GLKTextureLoaderCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *uintptr, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

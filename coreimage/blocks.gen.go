// Code generated from Apple documentation. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
)

// VoidHandler handles MTLTexture-rendering provider block to be called lazily when the destination is rendered to.
//
// Used by:
//   - [CIRenderDestination.InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CIRenderDestination.InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

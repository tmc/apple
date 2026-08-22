// Code generated from Apple documentation. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// CGRectInt32Handler handles A block or closure that computes the region of interest for a given rectangle of destination image pixels.
//
// Used by:
//   - [CIKernel.ApplyWithExtentRoiCallbackArguments]
//   - [CIWarpKernel.ApplyWithExtentRoiCallbackInputImageArguments]
type CGRectInt32Handler = func(int32, corefoundation.CGRect) corefoundation.CGRect

// NewCGRectInt32Block wraps a Go [CGRectInt32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CIKernel.ApplyWithExtentRoiCallbackArguments]
//   - [CIWarpKernel.ApplyWithExtentRoiCallbackInputImageArguments]
func NewCGRectInt32Block(handler CGRectInt32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int32, extra0 corefoundation.CGRect) corefoundation.CGRect {
		return handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLTextureVoidHandler handles MTLTexture-rendering provider block to be called lazily when the destination is rendered to.
//
// Used by:
//   - [CIRenderDestination.InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider]
type MTLTextureVoidHandler = func() metal.MTLTexture

// NewMTLTextureVoidBlock wraps a Go [MTLTextureVoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CIRenderDestination.InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider]
func NewMTLTextureVoidBlock(handler MTLTextureVoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) metal.MTLTexture {
		return handler()
	})
	return objc.ID(block), func() { block.Release() }
}

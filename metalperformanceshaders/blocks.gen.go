// Code generated from Apple documentation. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// MPSAccelerationStructureCompletionHandler handles A block of code that’s invoked when an operation on an acceleration structure has completed.

// NewMPSAccelerationStructureCompletionHandlerBlock wraps a Go [MPSAccelerationStructureCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMPSAccelerationStructureCompletionHandlerBlock(handler MPSAccelerationStructureCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal MPSAccelerationStructure) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// MPSAccelerationStructureHandler is the signature for a completion handler block.
//
// Used by:
//   - [MPSAccelerationStructure.RebuildWithCompletionHandler]
type MPSAccelerationStructureHandler = func(*MPSAccelerationStructure)

// NewMPSAccelerationStructureBlock wraps a Go [MPSAccelerationStructureHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MPSAccelerationStructure.RebuildWithCompletionHandler]
func NewMPSAccelerationStructureBlock(handler MPSAccelerationStructureHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *MPSAccelerationStructure
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MPSAccelerationStructureFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// MPSGradientNodeBlock handles completion with primitive and object results.

// NewMPSGradientNodeBlock wraps a Go [MPSGradientNodeBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMPSGradientNodeBlock(handler MPSGradientNodeBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MPSNNFilterNode, extra0 MPSNNFilterNode, extra1 MPSNNImageNode, extra2 MPSNNImageNode) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// MPSImageErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [MPSNNGraph.ExecuteAsyncWithSourceImagesCompletionHandler]
type MPSImageErrorHandler = func(*MPSImage, error)

// NewMPSImageErrorBlock wraps a Go [MPSImageErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MPSNNGraph.ExecuteAsyncWithSourceImagesCompletionHandler]
func NewMPSImageErrorBlock(handler MPSImageErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *MPSImage
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MPSImageFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MPSNNFilterNodeMPSNNFilterNodeMPSNNImageNodeMPSNNImageNodeHandler is the signature for a completion handler block.
//
// Used by:
//   - [MPSNNFilterNode.TrainingGraphWithSourceGradientNodeHandler]
type MPSNNFilterNodeMPSNNFilterNodeMPSNNImageNodeMPSNNImageNodeHandler = func(*MPSNNFilterNode, *MPSNNFilterNode, *MPSNNImageNode, *MPSNNImageNode)

// NewMPSNNFilterNodeMPSNNFilterNodeMPSNNImageNodeMPSNNImageNodeBlock wraps a Go [MPSNNFilterNodeMPSNNFilterNodeMPSNNImageNodeMPSNNImageNodeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MPSNNFilterNode.TrainingGraphWithSourceGradientNodeHandler]
func NewMPSNNFilterNodeMPSNNFilterNodeMPSNNImageNodeMPSNNImageNodeBlock(handler MPSNNFilterNodeMPSNNFilterNodeMPSNNImageNodeMPSNNImageNodeHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, extra1ID objc.ID, extra2ID objc.ID) {
		var result *MPSNNFilterNode
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MPSNNFilterNodeFromID(resultID)
			result = &v
		}
		var extra0 *MPSNNFilterNode
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := MPSNNFilterNodeFromID(extra0ID)
			extra0 = &v
		}
		var extra1 *MPSNNImageNode
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			v := MPSNNImageNodeFromID(extra1ID)
			extra1 = &v
		}
		var extra2 *MPSNNImageNode
		if extra2ID != 0 {
			objc.Send[objc.ID](extra2ID, objc.Sel("retain"))
			v := MPSNNImageNodeFromID(extra2ID)
			extra2 = &v
		}
		handler(result, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// MPSNNGraphCompletionHandler handles A notification when an asynchronous graph execution has finished.

// NewMPSNNGraphCompletionHandlerBlock wraps a Go [MPSNNGraphCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMPSNNGraphCompletionHandlerBlock(handler MPSNNGraphCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MPSImage, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler handles An optional block to allocate a new texture to hold the operation results, in case in-place operation is not possible.
//
// Used by:
//   - [MPSBinaryImageKernel.EncodeToCommandBufferInPlacePrimaryTextureSecondaryTextureFallbackCopyAllocator]
//   - [MPSBinaryImageKernel.EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator]
//   - [MPSUnaryImageKernel.EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator]
type MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler = func(*MPSKernel, metal.MTLCommandBuffer, metal.MTLTexture) metal.MTLTexture

// NewMTLTextureMPSKernelMTLCommandBufferMTLTextureBlock wraps a Go [MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MPSBinaryImageKernel.EncodeToCommandBufferInPlacePrimaryTextureSecondaryTextureFallbackCopyAllocator]
//   - [MPSBinaryImageKernel.EncodeToCommandBufferPrimaryTextureInPlaceSecondaryTextureFallbackCopyAllocator]
//   - [MPSUnaryImageKernel.EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator]
func NewMTLTextureMPSKernelMTLCommandBufferMTLTextureBlock(handler MTLTextureMPSKernelMTLCommandBufferMTLTextureHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, extra1ID objc.ID) metal.MTLTexture {
		var result *MPSKernel
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MPSKernelFromID(resultID)
			result = &v
		}
		var extra0 metal.MTLCommandBuffer
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			extra0 = metal.MTLCommandBufferObjectFromID(extra0ID)
		}
		var extra1 metal.MTLTexture
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			extra1 = metal.MTLTextureObjectFromID(extra1ID)
		}
		return handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

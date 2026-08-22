// Code generated from Apple documentation. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// Float32Float32Handler handles A custom block to be executed when a physics body is affected by the field.
//
// Used by:
//   - [SKFieldNode.CustomFieldWithEvaluationBlock]
type Float32Float32Handler = func(float32, float32, float32, float32, float64) float32

// NewFloat32Float32Block wraps a Go [Float32Float32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SKFieldNode.CustomFieldWithEvaluationBlock]
func NewFloat32Float32Block(handler Float32Float32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive float32, extra0 float32, extra1 float32, extra2 float32, extra3 float64) float32 {
		return handler(primitive, extra0, extra1, extra2, extra3)
	})
	return objc.ID(block), func() { block.Release() }
}

// SKActionTimingFunction handles The signature for the custom timing block.

// NewSKActionTimingFunctionBlock wraps a Go [SKActionTimingFunction] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewSKActionTimingFunctionBlock(handler SKActionTimingFunction) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal float32) float32 {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// SKNodeBoolHandler handles A block to execute on nodes that match the `name` parameter.
//
// Used by:
//   - [SKNode.EnumerateChildNodesWithNameUsingBlock]
type SKNodeBoolHandler = func(*SKNode, *bool)

// NewSKNodeBoolBlock wraps a Go [SKNodeBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SKNode.EnumerateChildNodesWithNameUsingBlock]
func NewSKNodeBoolBlock(handler SKNodeBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 *bool) {
		var result *SKNode
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := SKNodeFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// SKPhysicsBodyBoolHandler handles A block to be called for each physics body that contains the point.
//   - body: The physics body that the ray intersected.
//   - stop: A pointer to a Boolean variable. Your block can set this to [true](<https://developer.apple.com/documentation/Swift/true>) to terminate the enumeration.
//
// Used by:
//   - [SKPhysicsWorld.EnumerateBodiesAtPointUsingBlock]
//   - [SKPhysicsWorld.EnumerateBodiesInRectUsingBlock]
type SKPhysicsBodyBoolHandler = func(*SKPhysicsBody, *bool)

// NewSKPhysicsBodyBoolBlock wraps a Go [SKPhysicsBodyBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SKPhysicsWorld.EnumerateBodiesAtPointUsingBlock]
//   - [SKPhysicsWorld.EnumerateBodiesInRectUsingBlock]
func NewSKPhysicsBodyBoolBlock(handler SKPhysicsBodyBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 *bool) {
		var result *SKPhysicsBody
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := SKPhysicsBodyFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// SKPhysicsBodyCGPointCGVectorBoolHandler handles A block to be called for each physics body that the ray touches.
//   - body: The physics body that the ray intersected.
//   - point: The point in scene coordinates where the ray contacted the physics body.
//   - normal: The normal vector for the physics body at the point of contact.
//   - stop: A pointer to a Boolean variable. Your block can set this to [true](<https://developer.apple.com/documentation/Swift/true>) to terminate the enumeration.
//
// Used by:
//   - [SKPhysicsWorld.EnumerateBodiesAlongRayStartEndUsingBlock]
type SKPhysicsBodyCGPointCGVectorBoolHandler = func(*SKPhysicsBody, corefoundation.CGPoint, corefoundation.CGVector, *bool)

// NewSKPhysicsBodyCGPointCGVectorBoolBlock wraps a Go [SKPhysicsBodyCGPointCGVectorBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SKPhysicsWorld.EnumerateBodiesAlongRayStartEndUsingBlock]
func NewSKPhysicsBodyCGPointCGVectorBoolBlock(handler SKPhysicsBodyCGPointCGVectorBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 corefoundation.CGPoint, extra1 corefoundation.CGVector, extra2 *bool) {
		var result *SKPhysicsBody
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := SKPhysicsBodyFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// SKTextureAtlasArrayErrorHandler handles A block called after all of the texture atlases are loaded.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [SKTextureAtlas.PreloadTextureAtlasesNamedWithCompletionHandler]
type SKTextureAtlasArrayErrorHandler = func(*[]SKTextureAtlas, error)

// NewSKTextureAtlasArrayErrorBlock wraps a Go [SKTextureAtlasArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SKTextureAtlas.PreloadTextureAtlasesNamedWithCompletionHandler]
func NewSKTextureAtlasArrayErrorBlock(handler SKTextureAtlasArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]SKTextureAtlas
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]SKTextureAtlas, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = SKTextureAtlasFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// UnsafePointerUintptrHandler handles A block to be called when the texture can be safely modified.
//   - pixelData: A pointer to the start of the current texture data.
//   - lengthInBytes: The length of the texture data in bytes.
//
// Used by:
//   - [SKMutableTexture.ModifyPixelDataWithBlock]
type UnsafePointerUintptrHandler = func(unsafe.Pointer, uintptr)

// NewUnsafePointerUintptrBlock wraps a Go [UnsafePointerUintptrHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SKMutableTexture.ModifyPixelDataWithBlock]
func NewUnsafePointerUintptrBlock(handler UnsafePointerUintptrHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 uintptr) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles A completion block called when the action completes.
//
// Used by:
//   - [SKNode.RunActionCompletion]
//   - [SKTexture.PreloadTexturesWithCompletionHandler]
//   - [SKTexture.PreloadWithCompletionHandler]
//   - [SKTextureAtlas.PreloadTextureAtlasesWithCompletionHandler]
//   - [SKTextureAtlas.PreloadWithCompletionHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SKNode.RunActionCompletion]
//   - [SKTexture.PreloadTexturesWithCompletionHandler]
//   - [SKTexture.PreloadWithCompletionHandler]
//   - [SKTextureAtlas.PreloadTextureAtlasesWithCompletionHandler]
//   - [SKTextureAtlas.PreloadWithCompletionHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// Code generated from Apple documentation. DO NOT EDIT.

package coregraphics

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CGDisplayStreamFrameAvailableHandler handles A block called when a data stream has a new frame event to process.

// NewCGDisplayStreamFrameAvailableHandlerBlock wraps a Go [CGDisplayStreamFrameAvailableHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCGDisplayStreamFrameAvailableHandlerBlock(handler CGDisplayStreamFrameAvailableHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive CGDisplayStreamFrameStatus, extra0 uint64, extra1ID objc.ID, extra2 *CGDisplayStreamUpdateRef) {
		var extra1 objectivec.IObject
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			extra1 = objectivec.ObjectFromID(extra1ID)
		}
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// CGPathApplyBlock handles completion with a primitive value.

// NewCGPathApplyBlock wraps a Go [CGPathApplyBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCGPathApplyBlock(handler CGPathApplyBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *CGPathElement) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

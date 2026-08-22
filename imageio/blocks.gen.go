// Code generated from Apple documentation. DO NOT EDIT.

package imageio

import (
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objc"
)

// CGImageSourceAnimationBlock handles The block to execute for each frame of an image animation.

// NewCGImageSourceAnimationBlock wraps a Go [CGImageSourceAnimationBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCGImageSourceAnimationBlock(handler CGImageSourceAnimationBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint32, extra0 *coregraphics.CGImageRef, extra1 *bool) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// Code generated from Apple documentation. DO NOT EDIT.

package dispatch

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// DispatchIOHandler handles A handler block used to process operations on a dispatch I/O channel.

// NewDispatchIOHandlerBlock wraps a Go [DispatchIOHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewDispatchIOHandlerBlock(handler DispatchIOHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive bool, extra0 objectivec.Object, extra1 int32) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

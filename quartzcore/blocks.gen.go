// Code generated from Apple documentation. DO NOT EDIT.

package quartzcore

import (
	"github.com/tmc/apple/objc"
)

// VoidHandler handles A block object called when animations for this transaction group are completed.
//
// Used by:
//   - [CATransaction.SetCompletionBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CATransaction.SetCompletionBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}


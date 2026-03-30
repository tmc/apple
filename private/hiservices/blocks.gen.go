// Code generated from Apple documentation. DO NOT EDIT.

package hiservices

import (
	"github.com/tmc/apple/objc"
)

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [HIRunLoopSemaphore.InvokeLoopInModeForDurationWithBlock]
//   - [HIRunLoopSemaphore._observeWhilePerforming]
//   - [HIRunLoopUtilities.DeferActions]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [HIRunLoopSemaphore.InvokeLoopInModeForDurationWithBlock]
//   - [HIRunLoopSemaphore._observeWhilePerforming]
//   - [HIRunLoopUtilities.DeferActions]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

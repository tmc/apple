// Code generated from Apple documentation. DO NOT EDIT.

package os

import (
	"github.com/tmc/apple/objc"
)

// OSBlock handles A block that takes no arguments and returns no value.

// NewOSBlock wraps a Go [OSBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewOSBlock(handler OSBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

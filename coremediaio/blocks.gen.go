// Code generated from Apple documentation. DO NOT EDIT.

package coremediaio

import (
	"github.com/tmc/apple/objc"
)

// CMIOObjectPropertyListenerBlock handles completion with primitive and object results.

// NewCMIOObjectPropertyListenerBlock wraps a Go [CMIOObjectPropertyListenerBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCMIOObjectPropertyListenerBlock(handler CMIOObjectPropertyListenerBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint32, extra0 *CMIOObjectPropertyAddress) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

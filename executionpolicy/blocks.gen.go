// Code generated from Apple documentation. DO NOT EDIT.

package executionpolicy

import (
	"github.com/tmc/apple/objc"
)

// BoolHandler handles completion with a primitive value.
//
// Used by:
//   - [EPDeveloperTool.RequestDeveloperToolAccessWithCompletionHandler]
type BoolHandler = func(bool)

// NewBoolBlock wraps a Go [BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [EPDeveloperTool.RequestDeveloperToolAccessWithCompletionHandler]
func NewBoolBlock(handler BoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

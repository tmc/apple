// Code generated from Apple documentation. DO NOT EDIT.

package servicemanagement

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler handles A completion handler to call with the result of the unregistration operation.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [SMAppService.UnregisterWithCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SMAppService.UnregisterWithCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}


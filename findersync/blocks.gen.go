// Code generated from Apple documentation. DO NOT EDIT.

package findersync

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [FIFinderSyncController.SetLastUsedDateForItemWithURLCompletion]
//   - [FIFinderSyncController.SetTagDataForItemWithURLCompletion]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FIFinderSyncController.SetLastUsedDateForItemWithURLCompletion]
//   - [FIFinderSyncController.SetTagDataForItemWithURLCompletion]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// StringidDictionaryErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [FIFinderSync.ValuesForAttributesForItemWithURLCompletion]
type StringidDictionaryErrorHandler = func(*foundation.INSDictionary, error)

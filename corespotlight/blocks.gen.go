// Code generated from Apple documentation. DO NOT EDIT.

package corespotlight

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// CSSearchableItemArrayHandler handles A method the framework calls that provides an array of CSSearchableItem objects.
//
// Used by:
//   - [CSIndexExtensionRequestHandler.SearchableItemsForIdentifiersProtectionClassSearchableItemsHandler]
//   - [CSIndexExtensionRequestHandler.SearchableItemsForIdentifiersSearchableItemsHandler]
//   - [CSSearchableIndexDelegate.SearchableItemsForIdentifiersProtectionClassSearchableItemsHandler]
//   - [CSSearchableIndexDelegate.SearchableItemsForIdentifiersSearchableItemsHandler]
type CSSearchableItemArrayHandler = func(*[]CSSearchableItem)

// NewCSSearchableItemArrayBlock wraps a Go [CSSearchableItemArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CSIndexExtensionRequestHandler.SearchableItemsForIdentifiersProtectionClassSearchableItemsHandler]
//   - [CSIndexExtensionRequestHandler.SearchableItemsForIdentifiersSearchableItemsHandler]
//   - [CSSearchableIndexDelegate.SearchableItemsForIdentifiersProtectionClassSearchableItemsHandler]
//   - [CSSearchableIndexDelegate.SearchableItemsForIdentifiersSearchableItemsHandler]
func NewCSSearchableItemArrayBlock(handler CSSearchableItemArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]CSSearchableItem
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]CSSearchableItem, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = CSSearchableItemFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// CSSuggestionArrayHandler is the signature for a completion handler block.
type CSSuggestionArrayHandler = func(*[]CSSuggestion)

// NewCSSuggestionArrayBlock wraps a Go [CSSuggestionArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCSSuggestionArrayBlock(handler CSSuggestionArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]CSSuggestion
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]CSSuggestion, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = CSSuggestionFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// DataErrorHandler handles The block to call when the request has been journaled by the index, which means that the index makes a note that it has to perform this operation.
//   - error: If an error occurred, this parameter holds an error object that explains the error. Otherwise, the value of this parameter is `nil`.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CSSearchableIndex.FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler]
//   - [CSSearchableIndex.FetchLastClientStateWithCompletionHandler]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CSSearchableIndex.FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler]
//   - [CSSearchableIndex.FetchLastClientStateWithCompletionHandler]
func NewDataErrorBlock(handler DataErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles The block that’s called when the data has been journaled by the index, which means that the index makes a note that it has to perform this operation.
//   - error: If an error occurred, this parameter holds an error object that explains the error. Otherwise, the value of this parameter is `nil`.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CSSearchableIndex.DeleteAllSearchableItemsWithCompletionHandler]
//   - [CSSearchableIndex.DeleteSearchableItemsWithDomainIdentifiersCompletionHandler]
//   - [CSSearchableIndex.DeleteSearchableItemsWithIdentifiersCompletionHandler]
//   - [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler]
//   - [CSSearchableIndex.EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler]
//   - [CSSearchableIndex.IndexSearchableItemsCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CSSearchableIndex.DeleteAllSearchableItemsWithCompletionHandler]
//   - [CSSearchableIndex.DeleteSearchableItemsWithDomainIdentifiersCompletionHandler]
//   - [CSSearchableIndex.DeleteSearchableItemsWithIdentifiersCompletionHandler]
//   - [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler]
//   - [CSSearchableIndex.EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler]
//   - [CSSearchableIndex.IndexSearchableItemsCompletionHandler]
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

// VoidHandler handles The handler to call after all client state has been saved.
//
// Used by:
//   - [CSIndexExtensionRequestHandler.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler]
//   - [CSIndexExtensionRequestHandler.SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler]
//   - [CSSearchableIndexDelegate.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler]
//   - [CSSearchableIndexDelegate.SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CSIndexExtensionRequestHandler.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler]
//   - [CSIndexExtensionRequestHandler.SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler]
//   - [CSSearchableIndexDelegate.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler]
//   - [CSSearchableIndexDelegate.SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// Code generated from Apple documentation. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// DataHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSFileProviderEnumerator.CurrentSyncAnchorWithCompletionHandler]
type DataHandler = func(*foundation.NSData)

// NewDataBlock wraps a Go [DataHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderEnumerator.CurrentSyncAnchorWithCompletionHandler]
func NewDataBlock(handler DataHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles A block called by the system immediately after receiving the request.
//   - error: If an error occurs, this object contains information about the error; otherwise, it’s `nil`.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileProviderCustomAction.PerformActionWithIdentifierOnItemsWithIdentifiersCompletionHandler]
//   - [NSFileProviderExternalVolumeHandling.ShouldConnectExternalDomainWithCompletionHandler]
//   - [NSFileProviderManager.AddDomainCompletionHandler]
//   - [NSFileProviderManager.ClaimKnownFoldersLocalizedReasonCompletionHandler]
//   - [NSFileProviderManager.DisconnectWithReasonOptionsCompletionHandler]
//   - [NSFileProviderManager.EvictItemWithIdentifierCompletionHandler]
//   - [NSFileProviderManager.ImportDomainFromDirectoryAtURLCompletionHandler]
//   - [NSFileProviderManager.ReconnectWithCompletionHandler]
//   - [NSFileProviderManager.RegisterURLSessionTaskForItemWithIdentifierCompletionHandler]
//   - [NSFileProviderManager.ReimportItemsBelowItemWithIdentifierCompletionHandler]
//   - [NSFileProviderManager.ReleaseKnownFoldersLocalizedReasonCompletionHandler]
//   - [NSFileProviderManager.RemoveAllDomainsWithCompletionHandler]
//   - [NSFileProviderManager.RemoveDomainCompletionHandler]
//   - [NSFileProviderManager.RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler]
//   - [NSFileProviderManager.RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler]
//   - [NSFileProviderManager.SignalEnumeratorForContainerItemIdentifierCompletionHandler]
//   - [NSFileProviderManager.SignalErrorResolvedCompletionHandler]
//   - [NSFileProviderManager.WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler]
//   - [NSFileProviderManager.WaitForStabilizationWithCompletionHandler]
//   - [NSFileProviderReplicatedExtension.DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler]
//   - [NSFileProviderServicing.SupportedServiceSourcesForItemIdentifierCompletionHandler]
//   - [NSFileProviderThumbnailing.FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderCustomAction.PerformActionWithIdentifierOnItemsWithIdentifiersCompletionHandler]
//   - [NSFileProviderExternalVolumeHandling.ShouldConnectExternalDomainWithCompletionHandler]
//   - [NSFileProviderManager.AddDomainCompletionHandler]
//   - [NSFileProviderManager.ClaimKnownFoldersLocalizedReasonCompletionHandler]
//   - [NSFileProviderManager.DisconnectWithReasonOptionsCompletionHandler]
//   - [NSFileProviderManager.EvictItemWithIdentifierCompletionHandler]
//   - [NSFileProviderManager.ImportDomainFromDirectoryAtURLCompletionHandler]
//   - [NSFileProviderManager.ReconnectWithCompletionHandler]
//   - [NSFileProviderManager.RegisterURLSessionTaskForItemWithIdentifierCompletionHandler]
//   - [NSFileProviderManager.ReimportItemsBelowItemWithIdentifierCompletionHandler]
//   - [NSFileProviderManager.ReleaseKnownFoldersLocalizedReasonCompletionHandler]
//   - [NSFileProviderManager.RemoveAllDomainsWithCompletionHandler]
//   - [NSFileProviderManager.RemoveDomainCompletionHandler]
//   - [NSFileProviderManager.RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler]
//   - [NSFileProviderManager.RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler]
//   - [NSFileProviderManager.SignalEnumeratorForContainerItemIdentifierCompletionHandler]
//   - [NSFileProviderManager.SignalErrorResolvedCompletionHandler]
//   - [NSFileProviderManager.WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler]
//   - [NSFileProviderManager.WaitForStabilizationWithCompletionHandler]
//   - [NSFileProviderReplicatedExtension.DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler]
//   - [NSFileProviderServicing.SupportedServiceSourcesForItemIdentifierCompletionHandler]
//   - [NSFileProviderThumbnailing.FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler]
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

// FileProviderItemErrorHandler handles A block that you call after downloading the item’s metadata.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileProviderReplicatedExtension.ItemForIdentifierRequestCompletionHandler]
type FileProviderItemErrorHandler = func(NSFileProviderItem, error)

// NewFileProviderItemErrorBlock wraps a Go [FileProviderItemErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderReplicatedExtension.ItemForIdentifierRequestCompletionHandler]
func NewFileProviderItemErrorBlock(handler FileProviderItemErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result NSFileProviderItem
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = NSFileProviderItemObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FileProviderItemNSFileProviderItemFieldsBoolErrorHandler handles A block that you call after uploading the item to your remote storage.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileProviderReplicatedExtension.CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler]
//   - [NSFileProviderReplicatedExtension.ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler]
type FileProviderItemNSFileProviderItemFieldsBoolErrorHandler = func(NSFileProviderItem, NSFileProviderItemFields, bool, error)

// NewFileProviderItemNSFileProviderItemFieldsBoolErrorBlock wraps a Go [FileProviderItemNSFileProviderItemFieldsBoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderReplicatedExtension.CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler]
//   - [NSFileProviderReplicatedExtension.ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler]
func NewFileProviderItemNSFileProviderItemFieldsBoolErrorBlock(handler FileProviderItemNSFileProviderItemFieldsBoolErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 NSFileProviderItemFields, extra1 bool, errID objc.ID) {
		var result NSFileProviderItem
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = NSFileProviderItemObjectFromID(resultID)
		}
		handler(result, extra0, extra1, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FileProviderKnownFolderLocationsErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSFileProviderKnownFolderSupporting.GetKnownFolderLocationsCompletionHandler]
type FileProviderKnownFolderLocationsErrorHandler = func(*NSFileProviderKnownFolderLocations, error)

// NewFileProviderKnownFolderLocationsErrorBlock wraps a Go [FileProviderKnownFolderLocationsErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderKnownFolderSupporting.GetKnownFolderLocationsCompletionHandler]
func NewFileProviderKnownFolderLocationsErrorBlock(handler FileProviderKnownFolderLocationsErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSFileProviderKnownFolderLocations
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSFileProviderKnownFolderLocationsFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FileProviderServiceErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSFileProviderManager.GetServiceWithNameItemIdentifierCompletionHandler]
type FileProviderServiceErrorHandler = func(*foundation.NSFileProviderService, error)

// NewFileProviderServiceErrorBlock wraps a Go [FileProviderServiceErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderManager.GetServiceWithNameItemIdentifierCompletionHandler]
func NewFileProviderServiceErrorBlock(handler FileProviderServiceErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSFileProviderService
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSFileProviderServiceFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NSFileProviderDomainArrayErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSFileProviderManager.GetDomainsWithCompletionHandler]
type NSFileProviderDomainArrayErrorHandler = func(*[]NSFileProviderDomain, error)

// NewNSFileProviderDomainArrayErrorBlock wraps a Go [NSFileProviderDomainArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderManager.GetDomainsWithCompletionHandler]
func NewNSFileProviderDomainArrayErrorBlock(handler NSFileProviderDomainArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]NSFileProviderDomain
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NSFileProviderDomain, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NSFileProviderDomainFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// StringDataErrorHandler handles A block that you call once for each item in the `itemIdentifiers` array.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileProviderThumbnailing.FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler]
type StringDataErrorHandler = func(*string, *foundation.NSData, error)

// NewStringDataErrorBlock wraps a Go [StringDataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderThumbnailing.FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler]
func NewStringDataErrorBlock(handler StringDataErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		var extra0 *foundation.NSData
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := foundation.NSDataFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// StringStringErrorHandler handles A block that the system calls after it gets the items identifier.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileProviderManager.GetIdentifierForUserVisibleFileAtURLCompletionHandler]
type StringStringErrorHandler = func(*string, string, error)

// NewStringStringErrorBlock wraps a Go [StringStringErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderManager.GetIdentifierForUserVisibleFileAtURLCompletionHandler]
func NewStringStringErrorBlock(handler StringStringErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		var extra0 string = objc.IDToString(extra0ID)
		handler(result, extra0, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// URLErrorHandler handles A block that the system calls after determining the item’s URL.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileProviderManager.GetUserVisibleURLForItemIdentifierCompletionHandler]
//   - [NSFileProviderManager.RemoveDomainModeCompletionHandler]
type URLErrorHandler = func(*foundation.NSURL, error)

// NewURLErrorBlock wraps a Go [URLErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderManager.GetUserVisibleURLForItemIdentifierCompletionHandler]
//   - [NSFileProviderManager.RemoveDomainModeCompletionHandler]
func NewURLErrorBlock(handler URLErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSURLFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// URLFileProviderItemErrorHandler handles A block that you call after downloading the update.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileProviderIncrementalContentFetching.FetchContentsForItemWithIdentifierVersionUsingExistingContentsAtURLExistingVersionRequestCompletionHandler]
//   - [NSFileProviderReplicatedExtension.FetchContentsForItemWithIdentifierVersionRequestCompletionHandler]
type URLFileProviderItemErrorHandler = func(*foundation.NSURL, NSFileProviderItem, error)

// NewURLFileProviderItemErrorBlock wraps a Go [URLFileProviderItemErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderIncrementalContentFetching.FetchContentsForItemWithIdentifierVersionUsingExistingContentsAtURLExistingVersionRequestCompletionHandler]
//   - [NSFileProviderReplicatedExtension.FetchContentsForItemWithIdentifierVersionRequestCompletionHandler]
func NewURLFileProviderItemErrorBlock(handler URLFileProviderItemErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *foundation.NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSURLFromID(resultID)
			result = &v
		}
		var extra0 NSFileProviderItem
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			extra0 = NSFileProviderItemObjectFromID(extra0ID)
		}
		handler(result, extra0, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// URLFileProviderItemNSRangeNSFileProviderMaterializationFlagsErrorHandler handles A block that you call after downloading the item from your remote storage.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileProviderPartialContentFetching.FetchPartialContentsForItemWithIdentifierVersionRequestMinimalRangeAligningToOptionsCompletionHandler]
type URLFileProviderItemNSRangeNSFileProviderMaterializationFlagsErrorHandler = func(*foundation.NSURL, NSFileProviderItem, foundation.NSRange, NSFileProviderMaterializationFlags, error)

// NewURLFileProviderItemNSRangeNSFileProviderMaterializationFlagsErrorBlock wraps a Go [URLFileProviderItemNSRangeNSFileProviderMaterializationFlagsErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderPartialContentFetching.FetchPartialContentsForItemWithIdentifierVersionRequestMinimalRangeAligningToOptionsCompletionHandler]
func NewURLFileProviderItemNSRangeNSFileProviderMaterializationFlagsErrorBlock(handler URLFileProviderItemNSRangeNSFileProviderMaterializationFlagsErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, extra1 foundation.NSRange, extra2 NSFileProviderMaterializationFlags, errID objc.ID) {
		var result *foundation.NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSURLFromID(resultID)
			result = &v
		}
		var extra0 NSFileProviderItem
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			extra0 = NSFileProviderItemObjectFromID(extra0ID)
		}
		handler(result, extra0, extra1, extra2, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles A block that you call after you finish processing the changes.
//
// Used by:
//   - [NSFileProviderReplicatedExtension.ImportDidFinishWithCompletionHandler]
//   - [NSFileProviderReplicatedExtension.MaterializedItemsDidChangeWithCompletionHandler]
//   - [NSFileProviderReplicatedExtension.PendingItemsDidChangeWithCompletionHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderReplicatedExtension.ImportDidFinishWithCompletionHandler]
//   - [NSFileProviderReplicatedExtension.MaterializedItemsDidChangeWithCompletionHandler]
//   - [NSFileProviderReplicatedExtension.PendingItemsDidChangeWithCompletionHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// Code generated from Apple documentation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AppleEventDescriptorErrorHandler handles The completion handler Block that returns the result or an error.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSUserAppleScriptTask.ExecuteWithAppleEventCompletionHandler]
type AppleEventDescriptorErrorHandler = func(*NSAppleEventDescriptor, error)

// NewAppleEventDescriptorErrorBlock wraps a Go [AppleEventDescriptorErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSUserAppleScriptTask.ExecuteWithAppleEventCompletionHandler]
func NewAppleEventDescriptorErrorBlock(handler AppleEventDescriptorErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSAppleEventDescriptor
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSAppleEventDescriptorFromID(resultID)
			result = &v
		}
		handler(result, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// BackgroundActivityCompletionHandlerHandler handles A block of code to execute when the scheduler runs.
//
// Used by:
//   - [NSBackgroundActivityScheduler.ScheduleWithBlock]
type BackgroundActivityCompletionHandlerHandler = func(NSBackgroundActivityCompletionHandler)

// BoolHandler handles An optional block to be called when the request completes, performed as a background priority task.
//   - expired: A Boolean value that indicates whether the system is terminating a previous invocation of the `completionHandler` block.
//
// Used by:
//   - [NSExtensionContext.CompleteRequestReturningItemsCompletionHandler]
//   - [NSExtensionContext.OpenURLCompletionHandler]
//   - [NSProcessInfo.PerformExpiringActivityWithReasonUsingBlock]
type BoolHandler = func(bool)

// NewBoolBlock wraps a Go [BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSExtensionContext.CompleteRequestReturningItemsCompletionHandler]
//   - [NSExtensionContext.OpenURLCompletionHandler]
//   - [NSProcessInfo.PerformExpiringActivityWithReasonUsingBlock]
func NewBoolBlock(handler BoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolIObjectHandler handles The block to apply to elements in the array.
//   - obj: The element in the array.
//   - idx: The index of the element in the array.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further enumeration of the array. If a block stops further enumeration, that block continues to run until it’s finished. The `stop` argument is an out-only argument. You should only ever set this Boolean to [true](<https://developer.apple.com/documentation/Swift/true>) within the block.
//
// Used by:
//   - [NSArray.DifferenceFromArrayWithOptionsUsingEquivalenceTest]
//   - [NSArray.IndexOfObjectAtIndexesOptionsPassingTest]
//   - [NSArray.IndexOfObjectPassingTest]
//   - [NSArray.IndexOfObjectWithOptionsPassingTest]
//   - [NSArray.IndexesOfObjectsAtIndexesOptionsPassingTest]
//   - [NSArray.IndexesOfObjectsPassingTest]
//   - [NSArray.IndexesOfObjectsWithOptionsPassingTest]
//   - [NSComparisonPredicate.PredicateWithBlock]
//   - [NSCompoundPredicate.PredicateWithBlock]
//   - [NSDictionary.KeysOfEntriesPassingTest]
//   - [NSDictionary.KeysOfEntriesWithOptionsPassingTest]
//   - [NSOrderedSet.DifferenceFromOrderedSetWithOptionsUsingEquivalenceTest]
//   - [NSOrderedSet.IndexOfObjectAtIndexesOptionsPassingTest]
//   - [NSOrderedSet.IndexOfObjectPassingTest]
//   - [NSOrderedSet.IndexOfObjectWithOptionsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsAtIndexesOptionsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsWithOptionsPassingTest]
//   - [NSPredicate.PredicateWithBlock]
//   - [NSSet.ObjectsPassingTest]
//   - [NSSet.ObjectsWithOptionsPassingTest]
type BoolIObjectHandler = func(objectivec.IObject, uint, *bool) bool

// NewBoolIObjectBlock wraps a Go [BoolIObjectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSArray.DifferenceFromArrayWithOptionsUsingEquivalenceTest]
//   - [NSArray.IndexOfObjectAtIndexesOptionsPassingTest]
//   - [NSArray.IndexOfObjectPassingTest]
//   - [NSArray.IndexOfObjectWithOptionsPassingTest]
//   - [NSArray.IndexesOfObjectsAtIndexesOptionsPassingTest]
//   - [NSArray.IndexesOfObjectsPassingTest]
//   - [NSArray.IndexesOfObjectsWithOptionsPassingTest]
//   - [NSComparisonPredicate.PredicateWithBlock]
//   - [NSCompoundPredicate.PredicateWithBlock]
//   - [NSDictionary.KeysOfEntriesPassingTest]
//   - [NSDictionary.KeysOfEntriesWithOptionsPassingTest]
//   - [NSOrderedSet.DifferenceFromOrderedSetWithOptionsUsingEquivalenceTest]
//   - [NSOrderedSet.IndexOfObjectAtIndexesOptionsPassingTest]
//   - [NSOrderedSet.IndexOfObjectPassingTest]
//   - [NSOrderedSet.IndexOfObjectWithOptionsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsAtIndexesOptionsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsWithOptionsPassingTest]
//   - [NSPredicate.PredicateWithBlock]
//   - [NSSet.ObjectsPassingTest]
//   - [NSSet.ObjectsWithOptionsPassingTest]
func NewBoolIObjectBlock(handler BoolIObjectHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 uint, extra1 *bool) bool {
		var primitive objectivec.IObject
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = objectivec.ObjectFromID(primitiveID)
		}
		return handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolURLErrorHandler handles An optional error handler block for the file manager to call when an error occurs.
//   - url: An [NSURL](<https://developer.apple.com/documentation/Foundation/NSURL>) object that identifies the item for which the error occurred.
//   - error: An [NSError](<https://developer.apple.com/documentation/Foundation/NSError>) object that contains information about the error.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileManager.EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler]
type BoolURLErrorHandler = func(*NSURL, error) bool

// NewBoolURLErrorBlock wraps a Go [BoolURLErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileManager.EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler]
func NewBoolURLErrorBlock(handler BoolURLErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) bool {
		var result *NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSURLFromID(resultID)
			result = &v
		}
		return handler(result, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolUintHandler handles The Block to apply to elements in the set.
//   - idx: The index of the object.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further processing of the set. The `stop` argument is an out-only argument. You should only ever set this Boolean to YES within the Block.
//
// Used by:
//   - [NSIndexSet.IndexInRangeOptionsPassingTest]
//   - [NSIndexSet.IndexPassingTest]
//   - [NSIndexSet.IndexWithOptionsPassingTest]
//   - [NSIndexSet.IndexesInRangeOptionsPassingTest]
//   - [NSIndexSet.IndexesPassingTest]
//   - [NSIndexSet.IndexesWithOptionsPassingTest]
type BoolUintHandler = func(uint, *bool) bool

// NewBoolUintBlock wraps a Go [BoolUintHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSIndexSet.IndexInRangeOptionsPassingTest]
//   - [NSIndexSet.IndexPassingTest]
//   - [NSIndexSet.IndexWithOptionsPassingTest]
//   - [NSIndexSet.IndexesInRangeOptionsPassingTest]
//   - [NSIndexSet.IndexesPassingTest]
//   - [NSIndexSet.IndexesWithOptionsPassingTest]
func NewBoolUintBlock(handler BoolUintHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint, extra0 *bool) bool {
		return handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// CachedURLResponseHandler handles A completion handler that receives the cached URL response for the data task’s request, or `nil` if no response is found in the cache.
//
// Used by:
//   - [NSURLCache.GetCachedResponseForDataTaskCompletionHandler]
//   - [NSURLSessionDataDelegate.URLSessionDataTaskWillCacheResponseCompletionHandler]
type CachedURLResponseHandler = func(*NSCachedURLResponse)

// NewCachedURLResponseBlock wraps a Go [CachedURLResponseHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLCache.GetCachedResponseForDataTaskCompletionHandler]
//   - [NSURLSessionDataDelegate.URLSessionDataTaskWillCacheResponseCompletionHandler]
func NewCachedURLResponseBlock(handler CachedURLResponseHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSCachedURLResponse
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSCachedURLResponseFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// ClassErrorHandler handles A block capable of returning the data item as the specified type.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSItemProvider.RegisterItemForTypeIdentifierLoadHandler]
type ClassErrorHandler = func(objectivec.Class, error)

// NewClassErrorBlock wraps a Go [ClassErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSItemProvider.RegisterItemForTypeIdentifierLoadHandler]
func NewClassErrorBlock(handler ClassErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.Class, errID objc.ID) {
		handler(primitiveVal, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CompletionHandler handles A completion handler for getting an asynchronous attributed string.

// NewCompletionHandlerBlock wraps a Go [CompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCompletionHandlerBlock(handler CompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive NSAttributedString, extra0 INSDictionary, extra1 unsafe.Pointer) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// DataBoolErrorHandler handles The completion handler to call when all bytes are read, or an error occurs.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSURLSessionStreamTask.ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler]
type DataBoolErrorHandler = func(*NSData, bool, error)

// NewDataBoolErrorBlock wraps a Go [DataBoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSessionStreamTask.ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler]
func NewDataBoolErrorBlock(handler DataBoolErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 bool, errID objc.ID) {
		var result *NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSDataFromID(resultID)
			result = &v
		}
		handler(result, extra0, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DataErrorHandler handles The handler that’s called after the data is loaded.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSAttributedString.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
//   - [NSItemProvider.LoadDataRepresentationForContentTypeCompletionHandler]
//   - [NSItemProvider.LoadDataRepresentationForTypeIdentifierCompletionHandler]
//   - [NSItemProviderWriting.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
//   - [NSMutableString.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
//   - [NSString.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
//   - [NSURL.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
//   - [NSUserActivity.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
type DataErrorHandler = func(*NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSAttributedString.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
//   - [NSItemProvider.LoadDataRepresentationForContentTypeCompletionHandler]
//   - [NSItemProvider.LoadDataRepresentationForTypeIdentifierCompletionHandler]
//   - [NSItemProviderWriting.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
//   - [NSMutableString.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
//   - [NSString.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
//   - [NSURL.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
//   - [NSUserActivity.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
func NewDataErrorBlock(handler DataErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSDataFromID(resultID)
			result = &v
		}
		handler(result, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DataHandler handles A completion handler that is called when the download has been successfully canceled.
//
// Used by:
//   - [NSURLSessionDownloadTask.CancelByProducingResumeData]
//   - [NSURLSessionUploadTask.CancelByProducingResumeData]
type DataHandler = func(*NSData)

// NewDataBlock wraps a Go [DataHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSessionDownloadTask.CancelByProducingResumeData]
//   - [NSURLSessionUploadTask.CancelByProducingResumeData]
func NewDataBlock(handler DataHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSDataFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// DataURLResponseErrorHandler handles The completion handler to call when the load request is complete.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSURLSession.DataTaskWithRequestCompletionHandler]
//   - [NSURLSession.DataTaskWithURLCompletionHandler]
//   - [NSURLSession.UploadTaskWithRequestFromDataCompletionHandler]
//   - [NSURLSession.UploadTaskWithRequestFromFileCompletionHandler]
//   - [NSURLSession.UploadTaskWithResumeDataCompletionHandler]
type DataURLResponseErrorHandler = func(*NSData, *NSURLResponse, error)

// NewDataURLResponseErrorBlock wraps a Go [DataURLResponseErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSession.DataTaskWithRequestCompletionHandler]
//   - [NSURLSession.DataTaskWithURLCompletionHandler]
//   - [NSURLSession.UploadTaskWithRequestFromDataCompletionHandler]
//   - [NSURLSession.UploadTaskWithRequestFromFileCompletionHandler]
//   - [NSURLSession.UploadTaskWithResumeDataCompletionHandler]
func NewDataURLResponseErrorBlock(handler DataURLResponseErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSDataFromID(resultID)
			result = &v
		}
		var extra0 *NSURLResponse
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := NSURLResponseFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DateBoolBoolHandler handles The block to apply to each enumerated date.
//   - date: The enumerated date.
//   - idx: Whether `date` exactly matches the specified date components.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further processing of the array. The stop argument is an out-only argument. You should only ever set this Boolean to [true](<https://developer.apple.com/documentation/Swift/true>) within the Block.
//
// Used by:
//   - [NSCalendar.EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock]
type DateBoolBoolHandler = func(*NSDate, bool, *bool)

// NewDateBoolBoolBlock wraps a Go [DateBoolBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSCalendar.EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock]
func NewDateBoolBoolBlock(handler DateBoolBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 bool, extra1 *bool) {
		var result *NSDate
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSDateFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles A closure or block that the framework calls when the pause action completes.
//   - error: If an error occurs while waiting for access, this parameter contains an [NSError] object that describes the problem. If access is successfully granted, it is set to `nil`, and you may perform the intended file access.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSAttributedString.LoadFromHTMLWithDataOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithFileURLOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithRequestOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithStringOptionsCompletionHandler]
//   - [NSFileCoordinator.CoordinateAccessWithIntentsQueueByAccessor]
//   - [NSFileCoordinator.PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor]
//   - [NSFileManager.PauseSyncForUbiquitousItemAtURLCompletionHandler]
//   - [NSFileManager.ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler]
//   - [NSFileManager.UnmountVolumeAtURLOptionsCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedItemDeletionWithCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedItemEvictionWithCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedSubitemDeletionAtURLCompletionHandler]
//   - [NSFilePresenter.RelinquishPresentedItemToReader]
//   - [NSFilePresenter.RelinquishPresentedItemToWriter]
//   - [NSFilePresenter.SavePresentedItemChangesWithCompletionHandler]
//   - [NSItemProvider.RegisterCKShareWithContainerAllowedSharingOptionsPreparationHandler]
//   - [NSItemProvider.RegisterCloudKitShareWithPreparationHandler]
//   - [NSItemProvider.RegisterDataRepresentationForContentTypeVisibilityLoadHandler]
//   - [NSItemProvider.RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler]
//   - [NSItemProvider.RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandler]
//   - [NSItemProvider.RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler]
//   - [NSItemProvider.RegisterObjectOfClassVisibilityLoadHandler]
//   - [NSProgress.AddSubscriberForFileURLWithPublishingHandler]
//   - [NSURLSessionStreamTask.WriteDataTimeoutCompletionHandler]
//   - [NSURLSessionWebSocketTask.SendMessageCompletionHandler]
//   - [NSURLSessionWebSocketTask.SendPingWithPongReceiveHandler]
//   - [NSUserScriptTask.ExecuteWithCompletionHandler]
//   - [NSUserUnixTask.ExecuteWithArgumentsCompletionHandler]
//   - [NSXPCConnection.RemoteObjectProxyWithErrorHandler]
//   - [NSXPCConnection.SynchronousRemoteObjectProxyWithErrorHandler]
//   - [NSXPCProxyCreating.RemoteObjectProxyWithErrorHandler]
//   - [NSXPCProxyCreating.SynchronousRemoteObjectProxyWithErrorHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSAttributedString.LoadFromHTMLWithDataOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithFileURLOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithRequestOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithStringOptionsCompletionHandler]
//   - [NSFileCoordinator.CoordinateAccessWithIntentsQueueByAccessor]
//   - [NSFileCoordinator.PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor]
//   - [NSFileManager.PauseSyncForUbiquitousItemAtURLCompletionHandler]
//   - [NSFileManager.ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler]
//   - [NSFileManager.UnmountVolumeAtURLOptionsCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedItemDeletionWithCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedItemEvictionWithCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedSubitemDeletionAtURLCompletionHandler]
//   - [NSFilePresenter.RelinquishPresentedItemToReader]
//   - [NSFilePresenter.RelinquishPresentedItemToWriter]
//   - [NSFilePresenter.SavePresentedItemChangesWithCompletionHandler]
//   - [NSItemProvider.RegisterCKShareWithContainerAllowedSharingOptionsPreparationHandler]
//   - [NSItemProvider.RegisterCloudKitShareWithPreparationHandler]
//   - [NSItemProvider.RegisterDataRepresentationForContentTypeVisibilityLoadHandler]
//   - [NSItemProvider.RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler]
//   - [NSItemProvider.RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandler]
//   - [NSItemProvider.RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler]
//   - [NSItemProvider.RegisterObjectOfClassVisibilityLoadHandler]
//   - [NSProgress.AddSubscriberForFileURLWithPublishingHandler]
//   - [NSURLSessionStreamTask.WriteDataTimeoutCompletionHandler]
//   - [NSURLSessionWebSocketTask.SendMessageCompletionHandler]
//   - [NSURLSessionWebSocketTask.SendPingWithPongReceiveHandler]
//   - [NSUserScriptTask.ExecuteWithCompletionHandler]
//   - [NSUserUnixTask.ExecuteWithArgumentsCompletionHandler]
//   - [NSXPCConnection.RemoteObjectProxyWithErrorHandler]
//   - [NSXPCConnection.SynchronousRemoteObjectProxyWithErrorHandler]
//   - [NSXPCProxyCreating.RemoteObjectProxyWithErrorHandler]
//   - [NSXPCProxyCreating.SynchronousRemoteObjectProxyWithErrorHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// FileHandleHandler is the signature for a completion handler block.
type FileHandleHandler = func(*NSFileHandle)

// NewFileHandleBlock wraps a Go [FileHandleHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewFileHandleBlock(handler FileHandleHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSFileHandle
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSFileHandleFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// FileVersionErrorHandler handles A closure or block that the framework calls when the fetch action completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileManager.FetchLatestRemoteVersionOfItemAtURLCompletionHandler]
//   - [NSFileManager.UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler]
type FileVersionErrorHandler = func(*NSFileVersion, error)

// NewFileVersionErrorBlock wraps a Go [FileVersionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileManager.FetchLatestRemoteVersionOfItemAtURLCompletionHandler]
//   - [NSFileManager.UploadLocalVersionOfUbiquitousItemAtURLWithConflictResolutionPolicyCompletionHandler]
func NewFileVersionErrorBlock(handler FileVersionErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSFileVersion
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSFileVersionFromID(resultID)
			result = &v
		}
		handler(result, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// IObjectBoolHandler handles The block to apply to elements in the set.
//   - obj: The element in the set.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further processing of the set. The `stop` argument is an out-only argument. You should only ever set this Boolean to [true](<https://developer.apple.com/documentation/Swift/true>) within the block.
//
// Used by:
//   - [NSSet.EnumerateObjectsUsingBlock]
//   - [NSSet.EnumerateObjectsWithOptionsUsingBlock]
type IObjectBoolHandler = func(objectivec.IObject, *bool)

// NewIObjectBoolBlock wraps a Go [IObjectBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSSet.EnumerateObjectsUsingBlock]
//   - [NSSet.EnumerateObjectsWithOptionsUsingBlock]
func NewIObjectBoolBlock(handler IObjectBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 *bool) {
		var primitive objectivec.IObject
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = objectivec.ObjectFromID(primitiveID)
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// IObjectErrorHandler handles A block to be executed synchronously at the time a corresponding property is accessed.
//   - err: The error object that is being accessed.
//   - userInfoKey: The user info key corresponding to the accessed property.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSError.SetUserInfoValueProviderForDomainProvider]
type IObjectErrorHandler = func(error) objectivec.IObject

// NewIObjectErrorBlock wraps a Go [IObjectErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSError.SetUserInfoValueProviderForDomainProvider]
func NewIObjectErrorBlock(handler IObjectErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) objc.ID {
		return handler(SafeErrorFrom(errID)).GetID()
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// IObjectIObjectBoolHandler handles A block object to operate on entries in the dictionary.
//
// Used by:
//   - [NSDictionary.EnumerateKeysAndObjectsUsingBlock]
//   - [NSDictionary.EnumerateKeysAndObjectsWithOptionsUsingBlock]
type IObjectIObjectBoolHandler = func(objectivec.IObject, objectivec.IObject, *bool)

// NewIObjectIObjectBoolBlock wraps a Go [IObjectIObjectBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSDictionary.EnumerateKeysAndObjectsUsingBlock]
//   - [NSDictionary.EnumerateKeysAndObjectsWithOptionsUsingBlock]
func NewIObjectIObjectBoolBlock(handler IObjectIObjectBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0ID objc.ID, extra1 *bool) {
		var primitive objectivec.IObject
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = objectivec.ObjectFromID(primitiveID)
		}
		var extra0 objectivec.IObject
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			extra0 = objectivec.ObjectFromID(extra0ID)
		}
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// IObjectIObjectHandler handles The Block is applied to the object to be evaluated.
//   - evaluatedObject: The object to be evaluated.
//   - expressions: An array of predicate expressions that evaluates to a collection.
//   - context: A dictionary that the expression can use to store temporary state for one predicate evaluation.
//
// Used by:
//   - [NSExpression.ExpressionForBlockArguments]
type IObjectIObjectHandler = func(objectivec.IObject) objectivec.IObject

// NewIObjectIObjectBlock wraps a Go [IObjectIObjectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSExpression.ExpressionForBlockArguments]
func NewIObjectIObjectBlock(handler IObjectIObjectHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, valID objc.ID) objc.ID {
		var val objectivec.IObject
		if valID != 0 {
			objc.Send[objc.ID](valID, objc.Sel("retain"))
			obj := objectivec.ObjectFromID(valID)
			val = &obj
		}
		return handler(val).GetID()
	})
	return objc.ID(block), func() { block.Release() }
}

// IObjectNSRangeBoolHandler handles A closure or block to apply to ranges of the specified attribute in the attributed string, taking three arguments:
//
// Used by:
//   - [NSAttributedString.EnumerateAttributeInRangeOptionsUsingBlock]
type IObjectNSRangeBoolHandler = func(objectivec.IObject, NSRange, *bool)

// NewIObjectNSRangeBoolBlock wraps a Go [IObjectNSRangeBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSAttributedString.EnumerateAttributeInRangeOptionsUsingBlock]
func NewIObjectNSRangeBoolBlock(handler IObjectNSRangeBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 NSRange, extra1 *bool) {
		var primitive objectivec.IObject
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			obj := objectivec.ObjectFromID(primitiveID)
			primitive = &obj
		}
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// IObjectUintBoolHandler handles A closure or block to execute for each object in the array, taking three arguments:
//   - obj: The element in the array.
//   - idx: The index of the element in the array.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further enumeration of the array. If a block stops further enumeration, that block continues to run until it’s finished. When the [NSEnumerationConcurrent] enumeration option is specified, enumeration stops after all of the currently running blocks finish. The `stop` argument is an out-only argument. You should only ever set this Boolean to [true](<https://developer.apple.com/documentation/Swift/true>) within the block.
//
// Used by:
//   - [NSArray.EnumerateObjectsAtIndexesOptionsUsingBlock]
//   - [NSArray.EnumerateObjectsUsingBlock]
//   - [NSArray.EnumerateObjectsWithOptionsUsingBlock]
//   - [NSMetadataQuery.EnumerateResultsUsingBlock]
//   - [NSMetadataQuery.EnumerateResultsWithOptionsUsingBlock]
//   - [NSOrderedSet.EnumerateObjectsAtIndexesOptionsUsingBlock]
//   - [NSOrderedSet.EnumerateObjectsUsingBlock]
//   - [NSOrderedSet.EnumerateObjectsWithOptionsUsingBlock]
type IObjectUintBoolHandler = func(objectivec.IObject, uint, *bool)

// NewIObjectUintBoolBlock wraps a Go [IObjectUintBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSArray.EnumerateObjectsAtIndexesOptionsUsingBlock]
//   - [NSArray.EnumerateObjectsUsingBlock]
//   - [NSArray.EnumerateObjectsWithOptionsUsingBlock]
//   - [NSMetadataQuery.EnumerateResultsUsingBlock]
//   - [NSMetadataQuery.EnumerateResultsWithOptionsUsingBlock]
//   - [NSOrderedSet.EnumerateObjectsAtIndexesOptionsUsingBlock]
//   - [NSOrderedSet.EnumerateObjectsUsingBlock]
//   - [NSOrderedSet.EnumerateObjectsWithOptionsUsingBlock]
func NewIObjectUintBoolBlock(handler IObjectUintBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 uint, extra1 *bool) {
		var primitive objectivec.IObject
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = objectivec.ObjectFromID(primitiveID)
		}
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// IndexPathBoolHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSSet.EnumerateIndexPathsWithOptionsUsingBlock]
type IndexPathBoolHandler = func(*objectivec.Object, *bool)

// NewIndexPathBoolBlock wraps a Go [IndexPathBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSSet.EnumerateIndexPathsWithOptionsUsingBlock]
func NewIndexPathBoolBlock(handler IndexPathBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 *bool) {
		var result *objectivec.Object
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := objectivec.ObjectFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// InputStreamHandler handles A completion handler that your delegate method should call with the new body stream.
//
// Used by:
//   - [NSURLSessionTaskDelegate.URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler]
//   - [NSURLSessionTaskDelegate.URLSessionTaskNeedNewBodyStream]
type InputStreamHandler = func(*NSInputStream)

// NewInputStreamBlock wraps a Go [InputStreamHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSessionTaskDelegate.URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler]
//   - [NSURLSessionTaskDelegate.URLSessionTaskNeedNewBodyStream]
func NewInputStreamBlock(handler InputStreamHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSInputStream
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSInputStreamFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// InputStreamOutputStreamErrorHandler handles The completion handler block that returns streams.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSUserActivity.GetContinuationStreamsWithCompletionHandler]
type InputStreamOutputStreamErrorHandler = func(*NSInputStream, *NSOutputStream, error)

// NewInputStreamOutputStreamErrorBlock wraps a Go [InputStreamOutputStreamErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSUserActivity.GetContinuationStreamsWithCompletionHandler]
func NewInputStreamOutputStreamErrorBlock(handler InputStreamOutputStreamErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *NSInputStream
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSInputStreamFromID(resultID)
			result = &v
		}
		var extra0 *NSOutputStream
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := NSOutputStreamFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NSBackgroundActivityCompletionHandler handles completion with a primitive value.

// NewNSBackgroundActivityCompletionHandlerBlock wraps a Go [NSBackgroundActivityCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNSBackgroundActivityCompletionHandlerBlock(handler NSBackgroundActivityCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NSBackgroundActivityResult) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSComparisonResultIObjectHandler handles A comparator block used to compare the object `obj` with elements in the array.
//
// Used by:
//   - [NSArray.IndexOfObjectInSortedRangeOptionsUsingComparator]
//   - [NSArray.SortedArrayUsingComparator]
//   - [NSArray.SortedArrayWithOptionsUsingComparator]
//   - [NSDictionary.KeysSortedByValueUsingComparator]
//   - [NSDictionary.KeysSortedByValueWithOptionsUsingComparator]
//   - [NSMutableArray.SortUsingComparator]
//   - [NSMutableArray.SortWithOptionsUsingComparator]
//   - [NSMutableOrderedSet.SortRangeOptionsUsingComparator]
//   - [NSMutableOrderedSet.SortUsingComparator]
//   - [NSMutableOrderedSet.SortWithOptionsUsingComparator]
//   - [NSOrderedSet.IndexOfObjectInSortedRangeOptionsUsingComparator]
//   - [NSOrderedSet.SortedArrayUsingComparator]
//   - [NSOrderedSet.SortedArrayWithOptionsUsingComparator]
//   - [NSSortDescriptor.InitWithKeyAscendingComparator]
//   - [NSSortDescriptor.SortDescriptorWithKeyAscendingComparator]
type NSComparisonResultIObjectHandler = func(objectivec.IObject, objectivec.IObject) NSComparisonResult

// NewNSComparisonResultIObjectBlock wraps a Go [NSComparisonResultIObjectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSArray.IndexOfObjectInSortedRangeOptionsUsingComparator]
//   - [NSArray.SortedArrayUsingComparator]
//   - [NSArray.SortedArrayWithOptionsUsingComparator]
//   - [NSDictionary.KeysSortedByValueUsingComparator]
//   - [NSDictionary.KeysSortedByValueWithOptionsUsingComparator]
//   - [NSMutableArray.SortUsingComparator]
//   - [NSMutableArray.SortWithOptionsUsingComparator]
//   - [NSMutableOrderedSet.SortRangeOptionsUsingComparator]
//   - [NSMutableOrderedSet.SortUsingComparator]
//   - [NSMutableOrderedSet.SortWithOptionsUsingComparator]
//   - [NSOrderedSet.IndexOfObjectInSortedRangeOptionsUsingComparator]
//   - [NSOrderedSet.SortedArrayUsingComparator]
//   - [NSOrderedSet.SortedArrayWithOptionsUsingComparator]
//   - [NSSortDescriptor.InitWithKeyAscendingComparator]
//   - [NSSortDescriptor.SortDescriptorWithKeyAscendingComparator]
func NewNSComparisonResultIObjectBlock(handler NSComparisonResultIObjectHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0ID objc.ID) NSComparisonResult {
		var primitive objectivec.IObject
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			obj := objectivec.ObjectFromID(primitiveID)
			primitive = &obj
		}
		var extra0 objectivec.IObject
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			extra0 = objectivec.ObjectFromID(extra0ID)
		}
		return handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSFileVersionArrayErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSFileVersion.GetNonlocalVersionsOfItemAtURLCompletionHandler]
type NSFileVersionArrayErrorHandler = func(*[]NSFileVersion, error)

// NewNSFileVersionArrayErrorBlock wraps a Go [NSFileVersionArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileVersion.GetNonlocalVersionsOfItemAtURLCompletionHandler]
func NewNSFileVersionArrayErrorBlock(handler NSFileVersionArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]NSFileVersion
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NSFileVersion, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NSFileVersionFromID(item.GetID())
			}
			result = &res
		}
		handler(result, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NSHTTPCookieArrayHandler handles A completion handler that receives an array of cookies as its argument.
//
// Used by:
//   - [NSHTTPCookieStorage.GetCookiesForTaskCompletionHandler]
type NSHTTPCookieArrayHandler = func(*[]NSHTTPCookie)

// NewNSHTTPCookieArrayBlock wraps a Go [NSHTTPCookieArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSHTTPCookieStorage.GetCookiesForTaskCompletionHandler]
func NewNSHTTPCookieArrayBlock(handler NSHTTPCookieArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]NSHTTPCookie
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NSHTTPCookie, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NSHTTPCookieFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSItemProviderCompletionHandler handles A block that receives the item provider’s data.

// NewNSItemProviderCompletionHandlerBlock wraps a Go [NSItemProviderCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNSItemProviderCompletionHandlerBlock(handler NSItemProviderCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 NSError) {
		var primitive NSSecureCoding
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = NSSecureCodingObjectFromID(primitiveID)
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSItemProviderLoadHandler handles A block that loads the item provider’s data and coerces it to the specified type.

// NSItemProviderReadingErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSItemProvider.LoadObjectOfClassCompletionHandler]
type NSItemProviderReadingErrorHandler = func(NSItemProviderReading, error)

// NewNSItemProviderReadingErrorBlock wraps a Go [NSItemProviderReadingErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSItemProvider.LoadObjectOfClassCompletionHandler]
func NewNSItemProviderReadingErrorBlock(handler NSItemProviderReadingErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result NSItemProviderReading
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = NSItemProviderReadingObjectFromID(resultID)
		}
		handler(result, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NSOrderedCollectionChangeidOrderedCollectionChangeHandler handles A block receives an ordered collection change and returns an updated change.
//
// Used by:
//   - [NSOrderedCollectionDifference.DifferenceByTransformingChangesWithBlock]
type NSOrderedCollectionChangeidOrderedCollectionChangeHandler = func(*NSOrderedCollectionChange) NSOrderedCollectionChange

// NewNSOrderedCollectionChangeidOrderedCollectionChangeBlock wraps a Go [NSOrderedCollectionChangeidOrderedCollectionChangeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSOrderedCollectionDifference.DifferenceByTransformingChangesWithBlock]
func NewNSOrderedCollectionChangeidOrderedCollectionChangeBlock(handler NSOrderedCollectionChangeidOrderedCollectionChangeHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) objc.ID {
		var result *NSOrderedCollectionChange
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSOrderedCollectionChangeFromID(resultID)
			result = &v
		}
		return handler(result).ID
	})
	return objc.ID(block), func() { block.Release() }
}

// NSProgressUnpublishingHandler handles A block that the system calls when an observed progress object terminates the subscription.

// NewNSProgressUnpublishingHandlerBlock wraps a Go [NSProgressUnpublishingHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNSProgressUnpublishingHandlerBlock(handler NSProgressUnpublishingHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// NSRangeBoolHandler handles The block to apply to elements in the index set.
//   - range: The range of elements.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further processing of the array. The stop argument is an out-only argument. You should only ever set this Boolean to [true](<https://developer.apple.com/documentation/Swift/true>) within the Block.
//
// Used by:
//   - [NSIndexSet.EnumerateRangesInRangeOptionsUsingBlock]
//   - [NSIndexSet.EnumerateRangesUsingBlock]
//   - [NSIndexSet.EnumerateRangesWithOptionsUsingBlock]
type NSRangeBoolHandler = func(NSRange, *bool)

// NewNSRangeBoolBlock wraps a Go [NSRangeBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSIndexSet.EnumerateRangesInRangeOptionsUsingBlock]
//   - [NSIndexSet.EnumerateRangesUsingBlock]
//   - [NSIndexSet.EnumerateRangesWithOptionsUsingBlock]
func NewNSRangeBoolBlock(handler NSRangeBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive NSRange, extra0 *bool) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSURLSessionAuthChallengeDispositionURLCredentialHandler handles A handler that your delegate method must call.
//
// Used by:
//   - [NSURLSessionDelegate.URLSessionDidReceiveChallengeCompletionHandler]
//   - [NSURLSessionTaskDelegate.URLSessionTaskDidReceiveChallengeCompletionHandler]
type NSURLSessionAuthChallengeDispositionURLCredentialHandler = func(NSURLSessionAuthChallengeDisposition, *NSURLCredential)

// NewNSURLSessionAuthChallengeDispositionURLCredentialBlock wraps a Go [NSURLSessionAuthChallengeDispositionURLCredentialHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSessionDelegate.URLSessionDidReceiveChallengeCompletionHandler]
//   - [NSURLSessionTaskDelegate.URLSessionTaskDidReceiveChallengeCompletionHandler]
func NewNSURLSessionAuthChallengeDispositionURLCredentialBlock(handler NSURLSessionAuthChallengeDispositionURLCredentialHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive NSURLSessionAuthChallengeDisposition, extra0ID objc.ID) {
		var extra0 *NSURLCredential
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := NSURLCredentialFromID(extra0ID)
			extra0 = &v
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSURLSessionDataTaskArrayNSURLSessionUploadTaskArrayNSURLSessionDownloadTaskArrayHandler handles The completion handler to call with the list of tasks.
//
// Used by:
//   - [NSURLSession.GetTasksWithCompletionHandler]
type NSURLSessionDataTaskArrayNSURLSessionUploadTaskArrayNSURLSessionDownloadTaskArrayHandler = func(*[]NSURLSessionDataTask, *[]NSURLSessionUploadTask, *[]NSURLSessionDownloadTask)

// NewNSURLSessionDataTaskArrayNSURLSessionUploadTaskArrayNSURLSessionDownloadTaskArrayBlock wraps a Go [NSURLSessionDataTaskArrayNSURLSessionUploadTaskArrayNSURLSessionDownloadTaskArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSession.GetTasksWithCompletionHandler]
func NewNSURLSessionDataTaskArrayNSURLSessionUploadTaskArrayNSURLSessionDownloadTaskArrayBlock(handler NSURLSessionDataTaskArrayNSURLSessionUploadTaskArrayNSURLSessionDownloadTaskArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, extra1ID objc.ID) {
		var result *[]NSURLSessionDataTask
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NSURLSessionDataTask, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NSURLSessionDataTaskFromID(item.GetID())
			}
			result = &res
		}
		var extra0 *[]NSURLSessionUploadTask
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			obj := NSArrayFromID(extra0ID)
			count := obj.Count()
			res := make([]NSURLSessionUploadTask, count)
			for j := uint(0); j < count; j++ {
				item := obj.ObjectAtIndex(j)
				res[j] = NSURLSessionUploadTaskFromID(item.GetID())
			}
			extra0 = &res
		}
		var extra1 *[]NSURLSessionDownloadTask
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			obj := NSArrayFromID(extra1ID)
			count := obj.Count()
			res := make([]NSURLSessionDownloadTask, count)
			for j := uint(0); j < count; j++ {
				item := obj.ObjectAtIndex(j)
				res[j] = NSURLSessionDownloadTaskFromID(item.GetID())
			}
			extra1 = &res
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSURLSessionDelayedRequestDispositionURLRequestHandler handles A completion handler to perform the request.
//
// Used by:
//   - [NSURLSessionTaskDelegate.URLSessionTaskWillBeginDelayedRequestCompletionHandler]
type NSURLSessionDelayedRequestDispositionURLRequestHandler = func(NSURLSessionDelayedRequestDisposition, *NSURLRequest)

// NewNSURLSessionDelayedRequestDispositionURLRequestBlock wraps a Go [NSURLSessionDelayedRequestDispositionURLRequestHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSessionTaskDelegate.URLSessionTaskWillBeginDelayedRequestCompletionHandler]
func NewNSURLSessionDelayedRequestDispositionURLRequestBlock(handler NSURLSessionDelayedRequestDispositionURLRequestHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive NSURLSessionDelayedRequestDisposition, extra0ID objc.ID) {
		var extra0 *NSURLRequest
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := NSURLRequestFromID(extra0ID)
			extra0 = &v
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSURLSessionTaskArrayHandler handles The completion handler to call with the list of tasks.
//
// Used by:
//   - [NSURLSession.GetAllTasksWithCompletionHandler]
type NSURLSessionTaskArrayHandler = func(*[]NSURLSessionTask)

// NewNSURLSessionTaskArrayBlock wraps a Go [NSURLSessionTaskArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSession.GetAllTasksWithCompletionHandler]
func NewNSURLSessionTaskArrayBlock(handler NSURLSessionTaskArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]NSURLSessionTask
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NSURLSessionTask, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NSURLSessionTaskFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSUserAppleScriptTaskCompletionHandler handles Implement this block to retrieve the result of the AppleScript executed by [execute(withAppleEvent:completionHandler:)].

// NewNSUserAppleScriptTaskCompletionHandlerBlock wraps a Go [NSUserAppleScriptTaskCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNSUserAppleScriptTaskCompletionHandlerBlock(handler NSUserAppleScriptTaskCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive NSAppleEventDescriptor, extra0 NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSUserAutomatorTaskCompletionHandler handles Implement this block to retrieve the output of the Automator workflow executed by [execute(withInput:completionHandler:)].

// NewNSUserAutomatorTaskCompletionHandlerBlock wraps a Go [NSUserAutomatorTaskCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNSUserAutomatorTaskCompletionHandlerBlock(handler NSUserAutomatorTaskCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 NSError) {
		var primitive objectivec.IObject
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = objectivec.ObjectFromID(primitiveID)
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSUserScriptTaskCompletionHandler handles Implement this block to retrieve the error of the script executed by [execute(completionHandler:)].

// NewNSUserScriptTaskCompletionHandlerBlock wraps a Go [NSUserScriptTaskCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNSUserScriptTaskCompletionHandlerBlock(handler NSUserScriptTaskCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NSError) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSUserUnixTaskCompletionHandler handles Implement this block to retrieve an error from the Unix scripted executed by [execute(withArguments:completionHandler:)].

// NewNSUserUnixTaskCompletionHandlerBlock wraps a Go [NSUserUnixTaskCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNSUserUnixTaskCompletionHandlerBlock(handler NSUserUnixTaskCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NSError) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NotificationHandler handles The block that executes when receiving a notification.
//
// Used by:
//   - [NSNotificationCenter.AddObserverForNameObjectQueueUsingBlock]
type NotificationHandler = func(*NSNotification)

// NewNotificationBlock wraps a Go [NotificationHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSNotificationCenter.AddObserverForNameObjectQueueUsingBlock]
func NewNotificationBlock(handler NotificationHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSNotification
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSNotificationFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// ObjectErrorHandler handles The completion handler Block that returns the result or an error.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSUserAutomatorTask.ExecuteWithInputCompletionHandler]
type ObjectErrorHandler = func(objectivec.IObject, error)

// NewObjectErrorBlock wraps a Go [ObjectErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSUserAutomatorTask.ExecuteWithInputCompletionHandler]
func NewObjectErrorBlock(handler ObjectErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, valID objc.ID, errID objc.ID) {
		var val objectivec.IObject
		if valID != 0 {
			objc.Send[objc.ID](valID, objc.Sel("retain"))
			obj := objectivec.ObjectFromID(valID)
			val = &obj
		}
		handler(val, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// ObjectHandler handles A block to be executed when an operation is undone.
//
// Used by:
//   - [NSUndoManager.RegisterUndoWithTargetHandler]
type ObjectHandler = func(objectivec.IObject)

// NewObjectBlock wraps a Go [ObjectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSUndoManager.RegisterUndoWithTargetHandler]
func NewObjectBlock(handler ObjectHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, valID objc.ID) {
		var val objectivec.IObject
		if valID != 0 {
			objc.Send[objc.ID](valID, objc.Sel("retain"))
			obj := objectivec.ObjectFromID(valID)
			val = &obj
		}
		handler(val)
	})
	return objc.ID(block), func() { block.Release() }
}

// SecureCoding__kindofidErrorHandler handles A completion handler block to execute with the results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSItemProvider.LoadItemForTypeIdentifierOptionsCompletionHandler]
//   - [NSItemProvider.LoadPreviewImageWithOptionsCompletionHandler]
type SecureCoding__kindofidErrorHandler = func(NSSecureCoding, error)

// NewSecureCoding__kindofidErrorBlock wraps a Go [SecureCoding__kindofidErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSItemProvider.LoadItemForTypeIdentifierOptionsCompletionHandler]
//   - [NSItemProvider.LoadPreviewImageWithOptionsCompletionHandler]
func NewSecureCoding__kindofidErrorBlock(handler SecureCoding__kindofidErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, errID objc.ID) {
		var primitiveVal NSSecureCoding
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitiveVal = NSSecureCodingObjectFromID(primitiveID)
		}
		handler(primitiveVal, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// StringBoolHandler handles The block executed for the enumeration.
//   - line: The current line of the string being enumerated. The line contains just the contents of the line, without the line terminators. See [getLineStart(_:end:contentsEnd:for:)](<https://developer.apple.com/documentation/Foundation/NSString/getLineStart(_:end:contentsEnd:for:)>) for a discussion of line terminators.
//   - stop: A reference to a Boolean value that the block can use to stop the enumeration by setting `*stop = YES`; it should not touch `*stop` otherwise.
//
// Used by:
//   - [NSString.EnumerateLinesUsingBlock]
type StringBoolHandler = func(*string, *bool)

// NewStringBoolBlock wraps a Go [StringBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSString.EnumerateLinesUsingBlock]
func NewStringBoolBlock(handler StringBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 *bool) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// StringNSFileProviderServiceDictionaryErrorHandler handles A block that is called on an anonymous background queue.
//
// Used by:
//   - [NSFileManager.GetFileProviderServicesForItemAtURLCompletionHandler]
type StringNSFileProviderServiceDictionaryErrorHandler = func(*INSDictionary, error)

// StringNSRangeBoolHandler handles The block to apply to ranges of the string.
//   - tag: The located linguistic tag.
//   - tokenRange: The range of the linguistic tag.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further processing of the set. The `stop` argument is an out-only argument. You should only ever set this Boolean to [true](<https://developer.apple.com/documentation/Swift/true>) within the block.
//
// Used by:
//   - [NSLinguisticTagger.EnumerateTagsForStringRangeUnitSchemeOptionsOrthographyUsingBlock]
//   - [NSLinguisticTagger.EnumerateTagsInRangeUnitSchemeOptionsUsingBlock]
type StringNSRangeBoolHandler = func(*string, NSRange, *bool)

// NewStringNSRangeBoolBlock wraps a Go [StringNSRangeBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSLinguisticTagger.EnumerateTagsForStringRangeUnitSchemeOptionsOrthographyUsingBlock]
//   - [NSLinguisticTagger.EnumerateTagsInRangeUnitSchemeOptionsUsingBlock]
func NewStringNSRangeBoolBlock(handler StringNSRangeBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 NSRange, extra1 *bool) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// StringNSRangeNSRangeBoolHandler handles The block to apply to ranges of the string.
//   - tag: The located linguistic tag.
//   - tokenRange: The range of the linguistic tag.
//   - sentenceRange: The range of the sentence in which the tag occurs.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further processing of the set. The `stop` argument is an out-only argument. You should only ever set this Boolean to [true](<https://developer.apple.com/documentation/Swift/true>) within the block.
//
// Used by:
//   - [NSLinguisticTagger.EnumerateTagsInRangeSchemeOptionsUsingBlock]
//   - [NSString.EnumerateLinguisticTagsInRangeSchemeOptionsOrthographyUsingBlock]
//   - [NSString.EnumerateSubstringsInRangeOptionsUsingBlock]
type StringNSRangeNSRangeBoolHandler = func(*string, NSRange, NSRange, *bool)

// NewStringNSRangeNSRangeBoolBlock wraps a Go [StringNSRangeNSRangeBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSLinguisticTagger.EnumerateTagsInRangeSchemeOptionsUsingBlock]
//   - [NSString.EnumerateLinguisticTagsInRangeSchemeOptionsOrthographyUsingBlock]
//   - [NSString.EnumerateSubstringsInRangeOptionsUsingBlock]
func NewStringNSRangeNSRangeBoolBlock(handler StringNSRangeNSRangeBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 NSRange, extra1 NSRange, extra2 *bool) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		handler(result, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// StringNSURLCredentialDictionaryHandler handles A completion handler that receives a single argument with the credentials for the specified protection space and task.
//
// Used by:
//   - [NSURLCredentialStorage.GetCredentialsForProtectionSpaceTaskCompletionHandler]
type StringNSURLCredentialDictionaryHandler = func(*INSDictionary)

// StringStringImageHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSExtensionContext.LoadBroadcastingApplicationInfoWithCompletion]
type StringStringImageHandler = func(*string, string, *objectivec.Object)

// NewStringStringImageBlock wraps a Go [StringStringImageHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSExtensionContext.LoadBroadcastingApplicationInfoWithCompletion]
func NewStringStringImageBlock(handler StringStringImageHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, extra1ID objc.ID) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		var extra0 string = objc.IDToString(extra0ID)
		var extra1 *objectivec.Object
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			v := objectivec.ObjectFromID(extra1ID)
			extra1 = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// StringidDictionaryNSRangeBoolHandler handles The closure or block to apply to ranges of attributes in the attributed string, taking three arguments:
//
// Used by:
//   - [NSAttributedString.EnumerateAttributesInRangeOptionsUsingBlock]
type StringidDictionaryNSRangeBoolHandler = func(*INSDictionary, NSRange, *bool)

// TaskHandler handles The system invokes this completion block when the task has completed.
//
// Used by:
//   - [NSTask.LaunchedTaskWithExecutableURLArgumentsErrorTerminationHandler]
type TaskHandler = func(*NSTask)

// NewTaskBlock wraps a Go [TaskHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSTask.LaunchedTaskWithExecutableURLArgumentsErrorTerminationHandler]
func NewTaskBlock(handler TaskHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSTask
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSTaskFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// TextCheckingResultNSMatchingFlagsBoolHandler handles The Block enumerates the matches of the regular expression in the string.
//   - result: An [NSTextCheckingResult](<https://developer.apple.com/documentation/Foundation/NSTextCheckingResult>) specifying the match. This result gives the overall matched range via its [range](<https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/range>) property, and the range of each individual capture group via its [range(at:)](<https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/range(at:)>) method. The range {[NSNotFound], 0} is returned if one of the capture groups did not participate in this particular match.
//   - flags: The current state of the matching progress. See [NSRegularExpression.MatchingFlags](<https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags>) for the possible values.
//   - stop: A reference to a Boolean value. The Block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further processing of the array. The stop argument is an out-only argument. You should only ever set this Boolean to [true](<https://developer.apple.com/documentation/Swift/true>) within the Block.
//
// Used by:
//   - [NSRegularExpression.EnumerateMatchesInStringOptionsRangeUsingBlock]
type TextCheckingResultNSMatchingFlagsBoolHandler = func(*NSTextCheckingResult, NSMatchingFlags, *bool)

// NewTextCheckingResultNSMatchingFlagsBoolBlock wraps a Go [TextCheckingResultNSMatchingFlagsBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSRegularExpression.EnumerateMatchesInStringOptionsRangeUsingBlock]
func NewTextCheckingResultNSMatchingFlagsBoolBlock(handler TextCheckingResultNSMatchingFlagsBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 NSMatchingFlags, extra1 *bool) {
		var result *NSTextCheckingResult
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSTextCheckingResultFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// TimerHandler handles A block to be executed when the timer fires.
//
// Used by:
//   - [NSTimer.InitWithFireDateIntervalRepeatsBlock]
//   - [NSTimer.ScheduledTimerWithTimeIntervalRepeatsBlock]
//   - [NSTimer.TimerWithTimeIntervalRepeatsBlock]
type TimerHandler = func(*NSTimer)

// NewTimerBlock wraps a Go [TimerHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSTimer.InitWithFireDateIntervalRepeatsBlock]
//   - [NSTimer.ScheduledTimerWithTimeIntervalRepeatsBlock]
//   - [NSTimer.TimerWithTimeIntervalRepeatsBlock]
func NewTimerBlock(handler TimerHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSTimer
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSTimerFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLBoolErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSItemProvider.LoadFileRepresentationForContentTypeOpenInPlaceCompletionHandler]
//   - [NSItemProvider.LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler]
type URLBoolErrorHandler = func(*NSURL, bool, error)

// NewURLBoolErrorBlock wraps a Go [URLBoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSItemProvider.LoadFileRepresentationForContentTypeOpenInPlaceCompletionHandler]
//   - [NSItemProvider.LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler]
func NewURLBoolErrorBlock(handler URLBoolErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 bool, errID objc.ID) {
		var result *NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSURLFromID(resultID)
			result = &v
		}
		handler(result, extra0, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// URLCredentialHandler handles A completion handler that receives the default credential as its argument, or `nil` if there is no default credential for this combination of protection space and task.
//
// Used by:
//   - [NSURLCredentialStorage.GetDefaultCredentialForProtectionSpaceTaskCompletionHandler]
type URLCredentialHandler = func(*NSURLCredential)

// NewURLCredentialBlock wraps a Go [URLCredentialHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLCredentialStorage.GetDefaultCredentialForProtectionSpaceTaskCompletionHandler]
func NewURLCredentialBlock(handler URLCredentialHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSURLCredential
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSURLCredentialFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSItemProvider.LoadFileRepresentationForTypeIdentifierCompletionHandler]
type URLErrorHandler = func(*NSURL, error)

// NewURLErrorBlock wraps a Go [URLErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSItemProvider.LoadFileRepresentationForTypeIdentifierCompletionHandler]
func NewURLErrorBlock(handler URLErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSURLFromID(resultID)
			result = &v
		}
		handler(result, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// URLHandler handles A [Block object](<https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3>) containing the file operations you want to perform in a coordinated manner.
//
// Used by:
//   - [NSFileCoordinator.CoordinateReadingItemAtURLOptionsErrorByAccessor]
//   - [NSFileCoordinator.CoordinateWritingItemAtURLOptionsErrorByAccessor]
type URLHandler = func(*NSURL)

// NewURLBlock wraps a Go [URLHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileCoordinator.CoordinateReadingItemAtURLOptionsErrorByAccessor]
//   - [NSFileCoordinator.CoordinateWritingItemAtURLOptionsErrorByAccessor]
func NewURLBlock(handler URLHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSURLFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLRequestHandler handles A block that your handler should call with either the value of the `request` parameter, a modified URL request object, or [NULL] to refuse the redirect and return the body of the redirect response.
//
// Used by:
//   - [NSURLSessionTaskDelegate.URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler]
type URLRequestHandler = func(*NSURLRequest)

// NewURLRequestBlock wraps a Go [URLRequestHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSessionTaskDelegate.URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler]
func NewURLRequestBlock(handler URLRequestHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSURLRequest
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSURLRequestFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLSessionResponseDispositionHandler handles A completion handler that your code calls to continue a transfer, passing a URLSession.ResponseDisposition constant to indicate whether the transfer should continue as a data task or should become a download task.
//
// Used by:
//   - [NSURLSessionDataDelegate.URLSessionDataTaskDidReceiveResponseCompletionHandler]
type URLSessionResponseDispositionHandler = func(NSURLSessionResponseDisposition)

// NewURLSessionResponseDispositionBlock wraps a Go [URLSessionResponseDispositionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSessionDataDelegate.URLSessionDataTaskDidReceiveResponseCompletionHandler]
func NewURLSessionResponseDispositionBlock(handler URLSessionResponseDispositionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NSURLSessionResponseDisposition) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLSessionWebSocketMessageErrorHandler handles A closure that receives two parameters: the WebSocket message, and an NSError that indicates an error encountered while receiving the message.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSURLSessionWebSocketTask.ReceiveMessageWithCompletionHandler]
type URLSessionWebSocketMessageErrorHandler = func(*NSURLSessionWebSocketMessage, error)

// NewURLSessionWebSocketMessageErrorBlock wraps a Go [URLSessionWebSocketMessageErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSessionWebSocketTask.ReceiveMessageWithCompletionHandler]
func NewURLSessionWebSocketMessageErrorBlock(handler URLSessionWebSocketMessageErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSURLSessionWebSocketMessage
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSURLSessionWebSocketMessageFromID(resultID)
			result = &v
		}
		handler(result, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// URLURLHandler handles A [Block object](<https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3>) containing the read and write operations you want to perform in a coordinated manner.
//
// Used by:
//   - [NSFileCoordinator.CoordinateReadingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor]
//   - [NSFileCoordinator.CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor]
type URLURLHandler = func(*NSURL, *NSURL)

// NewURLURLBlock wraps a Go [URLURLHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileCoordinator.CoordinateReadingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor]
//   - [NSFileCoordinator.CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor]
func NewURLURLBlock(handler URLURLHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSURLFromID(resultID)
			result = &v
		}
		var extra0 *NSURL
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := NSURLFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLURLResponseErrorHandler handles The completion handler to call when the load request is complete.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSURLSession.DownloadTaskWithRequestCompletionHandler]
//   - [NSURLSession.DownloadTaskWithResumeDataCompletionHandler]
//   - [NSURLSession.DownloadTaskWithURLCompletionHandler]
type URLURLResponseErrorHandler = func(*NSURL, *NSURLResponse, error)

// NewURLURLResponseErrorBlock wraps a Go [URLURLResponseErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSession.DownloadTaskWithRequestCompletionHandler]
//   - [NSURLSession.DownloadTaskWithResumeDataCompletionHandler]
//   - [NSURLSession.DownloadTaskWithURLCompletionHandler]
func NewURLURLResponseErrorBlock(handler URLURLResponseErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSURLFromID(resultID)
			result = &v
		}
		var extra0 *NSURLResponse
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := NSURLResponseFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// UintBoolHandler handles The Block to apply to elements in the set.
//   - idx: The index of the object.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further processing of the set. The `stop` argument is an out-only argument. You should only ever set this Boolean to YES within the Block.
//
// Used by:
//   - [NSIndexSet.EnumerateIndexesInRangeOptionsUsingBlock]
//   - [NSIndexSet.EnumerateIndexesUsingBlock]
//   - [NSIndexSet.EnumerateIndexesWithOptionsUsingBlock]
type UintBoolHandler = func(uint, *bool)

// NewUintBoolBlock wraps a Go [UintBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSIndexSet.EnumerateIndexesInRangeOptionsUsingBlock]
//   - [NSIndexSet.EnumerateIndexesUsingBlock]
//   - [NSIndexSet.EnumerateIndexesWithOptionsUsingBlock]
func NewUintBoolBlock(handler UintBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint, extra0 *bool) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// UnsafePointerNSRangeBoolHandler handles The block to apply to byte ranges in the array.
//   - bytes: The bytes for the current range. This pointer is valid until the data object is deallocated.
//   - byteRange: The range of the current data bytes.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<https://developer.apple.com/documentation/Swift/true>) to stop further processing of the data. The stop argument is an out-only argument. You should only ever set this Boolean to [true](<https://developer.apple.com/documentation/Swift/true>) within the Block.
//
// Used by:
//   - [NSData.EnumerateByteRangesUsingBlock]
type UnsafePointerNSRangeBoolHandler = func(unsafe.Pointer, NSRange, *bool)

// NewUnsafePointerNSRangeBoolBlock wraps a Go [UnsafePointerNSRangeBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSData.EnumerateByteRangesUsingBlock]
func NewUnsafePointerNSRangeBoolBlock(handler UnsafePointerNSRangeBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 NSRange, extra1 *bool) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// UnsafePointerUintHandler handles A block to invoke when the resulting [NSData] object is deallocated.
//
// Used by:
//   - [NSConstantString.InitWithBytesNoCopyLengthEncodingDeallocator]
//   - [NSData.InitWithBytesNoCopyLengthDeallocator]
//   - [NSMutableData.InitWithBytesNoCopyLengthDeallocator]
//   - [NSMutableString.InitWithBytesNoCopyLengthEncodingDeallocator]
//   - [NSPurgeableData.InitWithBytesNoCopyLengthDeallocator]
//   - [NSSimpleCString.InitWithBytesNoCopyLengthEncodingDeallocator]
//   - [NSString.InitWithBytesNoCopyLengthEncodingDeallocator]
type UnsafePointerUintHandler = func(unsafe.Pointer, uint)

// NewUnsafePointerUintBlock wraps a Go [UnsafePointerUintHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSConstantString.InitWithBytesNoCopyLengthEncodingDeallocator]
//   - [NSData.InitWithBytesNoCopyLengthDeallocator]
//   - [NSMutableData.InitWithBytesNoCopyLengthDeallocator]
//   - [NSMutableString.InitWithBytesNoCopyLengthEncodingDeallocator]
//   - [NSPurgeableData.InitWithBytesNoCopyLengthDeallocator]
//   - [NSSimpleCString.InitWithBytesNoCopyLengthEncodingDeallocator]
//   - [NSString.InitWithBytesNoCopyLengthEncodingDeallocator]
func NewUnsafePointerUintBlock(handler UnsafePointerUintHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 uint) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// UnsafePointerVoidHandler is the signature for a completion handler block.
type UnsafePointerVoidHandler = func() unsafe.Pointer

// NewUnsafePointerVoidBlock wraps a Go [UnsafePointerVoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewUnsafePointerVoidBlock(handler UnsafePointerVoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) unsafe.Pointer {
		return handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles The block to add to the new block operation object’s list.
//
// Used by:
//   - [NSBlockOperation.AddExecutionBlock]
//   - [NSBlockOperation.BlockOperationWithBlock]
//   - [NSOperationQueue.AddBarrierBlock]
//   - [NSOperationQueue.AddOperationWithBlock]
//   - [NSProcessInfo.PerformActivityWithOptionsReasonUsingBlock]
//   - [NSProgress.PerformAsCurrentWithPendingUnitCountUsingBlock]
//   - [NSRunLoop.PerformBlock]
//   - [NSRunLoop.PerformInModesBlock]
//   - [NSThread.DetachNewThreadWithBlock]
//   - [NSThread.InitWithBlock]
//   - [NSURLSession.FlushWithCompletionHandler]
//   - [NSURLSession.ResetWithCompletionHandler]
//   - [NSUserActivity.DeleteAllSavedUserActivitiesWithCompletionHandler]
//   - [NSUserActivity.DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler]
//   - [NSXPCConnection.ScheduleSendBarrierBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSBlockOperation.AddExecutionBlock]
//   - [NSBlockOperation.BlockOperationWithBlock]
//   - [NSOperationQueue.AddBarrierBlock]
//   - [NSOperationQueue.AddOperationWithBlock]
//   - [NSProcessInfo.PerformActivityWithOptionsReasonUsingBlock]
//   - [NSProgress.PerformAsCurrentWithPendingUnitCountUsingBlock]
//   - [NSRunLoop.PerformBlock]
//   - [NSRunLoop.PerformInModesBlock]
//   - [NSThread.DetachNewThreadWithBlock]
//   - [NSThread.InitWithBlock]
//   - [NSURLSession.FlushWithCompletionHandler]
//   - [NSURLSession.ResetWithCompletionHandler]
//   - [NSUserActivity.DeleteAllSavedUserActivitiesWithCompletionHandler]
//   - [NSUserActivity.DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler]
//   - [NSXPCConnection.ScheduleSendBarrierBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// XPCConnectionErrorHandler handles A block that is called on an anonymous background queue.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileProviderService.GetFileProviderConnectionWithCompletionHandler]
type XPCConnectionErrorHandler = func(*NSXPCConnection, error)

// NewXPCConnectionErrorBlock wraps a Go [XPCConnectionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileProviderService.GetFileProviderConnectionWithCompletionHandler]
func NewXPCConnectionErrorBlock(handler XPCConnectionErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSXPCConnection
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSXPCConnectionFromID(resultID)
			result = &v
		}
		handler(result, SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// unsignedshortUintHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSConstantString.InitWithCharactersNoCopyLengthDeallocator]
//   - [NSMutableString.InitWithCharactersNoCopyLengthDeallocator]
//   - [NSSimpleCString.InitWithCharactersNoCopyLengthDeallocator]
//   - [NSString.InitWithCharactersNoCopyLengthDeallocator]
type unsignedshortUintHandler = func(uint16, uint)

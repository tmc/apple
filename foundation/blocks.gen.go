// Code generated from Apple documentation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ArrayErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSFileVersion.GetNonlocalVersionsOfItemAtURLCompletionHandler]
type ArrayErrorHandler = func(*[]NSFileVersion, error)

// ArrayHandler handles The Block is applied to the object to be evaluated.
//   - evaluatedObject: The object to be evaluated.
//   - expressions: An array of predicate expressions that evaluates to a collection.
//   - context: A dictionary that the expression can use to store temporary state for one predicate evaluation.
//
// Used by:
//   - [NSExpression.ExpressionForBlockArguments]
//   - [NSHTTPCookieStorage.GetCookiesForTaskCompletionHandler]
//   - [NSURLSession.GetAllTasksWithCompletionHandler]
type ArrayHandler = func(objectivec.IObject)

// NewArrayBlock wraps a Go [ArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSExpression.ExpressionForBlockArguments]
//   - [NSHTTPCookieStorage.GetCookiesForTaskCompletionHandler]
//   - [NSURLSession.GetAllTasksWithCompletionHandler]
func NewArrayBlock(handler ArrayHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, valID objc.ID) {
		var val objectivec.IObject
		if valID != 0 {
			obj := objectivec.ObjectFromID(valID)
			val = &obj
		}
		handler(val)
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
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSCachedURLResponse
		if resultID != 0 {
			v := NSCachedURLResponseFromID(resultID)
			result = &v
		}
		handler(result)
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
//   - [NSURLSessionStreamTask.ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler]
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
//   - [NSURLSessionStreamTask.ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler]
//   - [NSUserActivity.LoadDataWithTypeIdentifierForItemProviderCompletionHandler]
func NewDataErrorBlock(handler DataErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSData
		if resultID != 0 {
			v := NSDataFromID(resultID)
			result = &v
		}
		var nserr *NSError
		if errID != 0 {
			e := NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, NSErrorToError(nserr))
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSData
		if resultID != 0 {
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *NSData
		if resultID != 0 {
			v := NSDataFromID(resultID)
			result = &v
		}
		var extra0 *NSURLResponse
		if extra0ID != 0 {
			v := NSURLResponseFromID(extra0ID)
			extra0 = &v
		}
		var nserr *NSError
		if errID != 0 {
			e := NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, extra0, NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// DateHandler handles The block to apply to each enumerated date.
//   - date: The enumerated date.
//   - idx: Whether `date` exactly matches the specified date components.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<doc://com.apple.documentation/documentation/Swift/true>) to stop further processing of the array. The stop argument is an out-only argument. You should only ever set this Boolean to [true](<doc://com.apple.documentation/documentation/Swift/true>) within the Block.
//
// Used by:
//   - [NSCalendar.EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock]
type DateHandler = func(*NSDate)

// NewDateBlock wraps a Go [DateHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSCalendar.EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock]
func NewDateBlock(handler DateHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSDate
		if resultID != 0 {
			v := NSDateFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// DictionaryErrorHandler handles A block that is called on an anonymous background queue.
//
// Used by:
//   - [NSFileManager.GetFileProviderServicesForItemAtURLCompletionHandler]
type DictionaryErrorHandler = func(*INSDictionary, error)

// DictionaryHandler handles The block is applied to the object to be evaluated.
//   - evaluatedObject: The object to be evaluated.
//   - bindings: The substitution variables dictionary. The dictionary must contain key-value pairs for all variables in the receiver.
//
// Used by:
//   - [NSComparisonPredicate.PredicateWithBlock]
//   - [NSCompoundPredicate.PredicateWithBlock]
//   - [NSPredicate.PredicateWithBlock]
//   - [NSURLCredentialStorage.GetCredentialsForProtectionSpaceTaskCompletionHandler]
type DictionaryHandler = func(*INSDictionary)

// ErrorHandler handles A closure or block that the framework calls when the pause action completes.
//   - err: The error object that is being accessed.
//   - userInfoKey: The user info key corresponding to the accessed property.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSAppleEventManager.DispatchRawAppleEventWithRawReplyHandlerRefCon]
//   - [NSAppleEventManager.SetEventHandlerAndSelectorForEventClassAndEventID]
//   - [NSAttributedString.LoadFromHTMLWithDataOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithFileURLOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithRequestOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithStringOptionsCompletionHandler]
//   - [NSError.SetUserInfoValueProviderForDomainProvider]
//   - [NSFileCoordinator.CoordinateAccessWithIntentsQueueByAccessor]
//   - [NSFileManager.PauseSyncForUbiquitousItemAtURLCompletionHandler]
//   - [NSFileManager.ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler]
//   - [NSFileManager.UnmountVolumeAtURLOptionsCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedItemDeletionWithCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedItemEvictionWithCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedSubitemDeletionAtURLCompletionHandler]
//   - [NSFilePresenter.SavePresentedItemChangesWithCompletionHandler]
//   - [NSItemProvider.LoadItemForTypeIdentifierOptionsCompletionHandler]
//   - [NSItemProvider.LoadPreviewImageWithOptionsCompletionHandler]
//   - [NSItemProvider.RegisterCKShareWithContainerAllowedSharingOptionsPreparationHandler]
//   - [NSItemProvider.RegisterItemForTypeIdentifierLoadHandler]
//   - [NSProgress.AddSubscriberForFileURLWithPublishingHandler]
//   - [NSURLSessionStreamTask.WriteDataTimeoutCompletionHandler]
//   - [NSURLSessionWebSocketTask.SendMessageCompletionHandler]
//   - [NSURLSessionWebSocketTask.SendPingWithPongReceiveHandler]
//   - [NSUserAppleScriptTask.ExecuteWithAppleEventCompletionHandler]
//   - [NSUserAutomatorTask.ExecuteWithInputCompletionHandler]
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
//   - [NSAppleEventManager.DispatchRawAppleEventWithRawReplyHandlerRefCon]
//   - [NSAppleEventManager.SetEventHandlerAndSelectorForEventClassAndEventID]
//   - [NSAttributedString.LoadFromHTMLWithDataOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithFileURLOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithRequestOptionsCompletionHandler]
//   - [NSAttributedString.LoadFromHTMLWithStringOptionsCompletionHandler]
//   - [NSError.SetUserInfoValueProviderForDomainProvider]
//   - [NSFileCoordinator.CoordinateAccessWithIntentsQueueByAccessor]
//   - [NSFileManager.PauseSyncForUbiquitousItemAtURLCompletionHandler]
//   - [NSFileManager.ResumeSyncForUbiquitousItemAtURLWithBehaviorCompletionHandler]
//   - [NSFileManager.UnmountVolumeAtURLOptionsCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedItemDeletionWithCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedItemEvictionWithCompletionHandler]
//   - [NSFilePresenter.AccommodatePresentedSubitemDeletionAtURLCompletionHandler]
//   - [NSFilePresenter.SavePresentedItemChangesWithCompletionHandler]
//   - [NSItemProvider.LoadItemForTypeIdentifierOptionsCompletionHandler]
//   - [NSItemProvider.LoadPreviewImageWithOptionsCompletionHandler]
//   - [NSItemProvider.RegisterCKShareWithContainerAllowedSharingOptionsPreparationHandler]
//   - [NSItemProvider.RegisterItemForTypeIdentifierLoadHandler]
//   - [NSProgress.AddSubscriberForFileURLWithPublishingHandler]
//   - [NSURLSessionStreamTask.WriteDataTimeoutCompletionHandler]
//   - [NSURLSessionWebSocketTask.SendMessageCompletionHandler]
//   - [NSURLSessionWebSocketTask.SendPingWithPongReceiveHandler]
//   - [NSUserAppleScriptTask.ExecuteWithAppleEventCompletionHandler]
//   - [NSUserAutomatorTask.ExecuteWithInputCompletionHandler]
//   - [NSUserScriptTask.ExecuteWithCompletionHandler]
//   - [NSUserUnixTask.ExecuteWithArgumentsCompletionHandler]
//   - [NSXPCConnection.RemoteObjectProxyWithErrorHandler]
//   - [NSXPCConnection.SynchronousRemoteObjectProxyWithErrorHandler]
//   - [NSXPCProxyCreating.RemoteObjectProxyWithErrorHandler]
//   - [NSXPCProxyCreating.SynchronousRemoteObjectProxyWithErrorHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		var nserr *NSError
		if errID != 0 {
			e := NSErrorFromID(errID)
			nserr = &e
		}
		handler(NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// FileHandleHandler is the signature for a completion handler block.
type FileHandleHandler = func(*NSFileHandle)

// NewFileHandleBlock wraps a Go [FileHandleHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewFileHandleBlock(handler FileHandleHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSFileHandle
		if resultID != 0 {
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSFileVersion
		if resultID != 0 {
			v := NSFileVersionFromID(resultID)
			result = &v
		}
		var nserr *NSError
		if errID != 0 {
			e := NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// IndexPathHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSSet.EnumerateIndexPathsWithOptionsUsingBlock]
type IndexPathHandler = func(*objc.ID)

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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSInputStream
		if resultID != 0 {
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *NSInputStream
		if resultID != 0 {
			v := NSInputStreamFromID(resultID)
			result = &v
		}
		var extra0 *NSOutputStream
		if extra0ID != 0 {
			v := NSOutputStreamFromID(extra0ID)
			extra0 = &v
		}
		var nserr *NSError
		if errID != 0 {
			e := NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, extra0, NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// KeyTypeHandler handles A block object to operate on entries in the dictionary.
//
// Used by:
//   - [NSDictionary.EnumerateKeysAndObjectsUsingBlock]
//   - [NSDictionary.EnumerateKeysAndObjectsWithOptionsUsingBlock]
//   - [NSDictionary.KeysOfEntriesPassingTest]
//   - [NSDictionary.KeysOfEntriesWithOptionsPassingTest]
type KeyTypeHandler = func(objectivec.IObject)

// NewKeyTypeBlock wraps a Go [KeyTypeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSDictionary.EnumerateKeysAndObjectsUsingBlock]
//   - [NSDictionary.EnumerateKeysAndObjectsWithOptionsUsingBlock]
//   - [NSDictionary.KeysOfEntriesPassingTest]
//   - [NSDictionary.KeysOfEntriesWithOptionsPassingTest]
func NewKeyTypeBlock(handler KeyTypeHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.IObject) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result NSItemProviderReading
		if resultID != 0 {
			result = NSItemProviderReadingObjectFromID(resultID)
		}
		var nserr *NSError
		if errID != 0 {
			e := NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, NSErrorToError(nserr))
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSNotification
		if resultID != 0 {
			v := NSNotificationFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// ObjectHandler handles A closure or block to apply to ranges of the specified attribute in the attributed string, taking three arguments:
//
// Used by:
//   - [NSAttributedString.EnumerateAttributeInRangeOptionsUsingBlock]
//   - [NSMetadataQuery.EnumerateResultsUsingBlock]
//   - [NSMetadataQuery.EnumerateResultsWithOptionsUsingBlock]
//   - [NSUndoManager.RegisterUndoWithTargetHandler]
type ObjectHandler = func(objectivec.IObject)

// NewObjectBlock wraps a Go [ObjectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSAttributedString.EnumerateAttributeInRangeOptionsUsingBlock]
//   - [NSMetadataQuery.EnumerateResultsUsingBlock]
//   - [NSMetadataQuery.EnumerateResultsWithOptionsUsingBlock]
//   - [NSUndoManager.RegisterUndoWithTargetHandler]
func NewObjectBlock(handler ObjectHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, valID objc.ID) {
		var val objectivec.IObject
		if valID != 0 {
			obj := objectivec.ObjectFromID(valID)
			val = &obj
		}
		handler(val)
	})
	return objc.ID(block), func() { block.Release() }
}

// ObjectTypeHandler handles The block to apply to elements in the array.
//   - obj: The element in the array.
//   - idx: The index of the element in the array.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<doc://com.apple.documentation/documentation/Swift/true>) to stop further enumeration of the array. If a block stops further enumeration, that block continues to run until it’s finished. The `stop` argument is an out-only argument. You should only ever set this Boolean to [true](<doc://com.apple.documentation/documentation/Swift/true>) within the block.
//
// Used by:
//   - [NSArray.DifferenceFromArrayWithOptionsUsingEquivalenceTest]
//   - [NSArray.EnumerateObjectsAtIndexesOptionsUsingBlock]
//   - [NSArray.EnumerateObjectsUsingBlock]
//   - [NSArray.EnumerateObjectsWithOptionsUsingBlock]
//   - [NSArray.IndexOfObjectAtIndexesOptionsPassingTest]
//   - [NSArray.IndexOfObjectPassingTest]
//   - [NSArray.IndexOfObjectWithOptionsPassingTest]
//   - [NSArray.IndexesOfObjectsAtIndexesOptionsPassingTest]
//   - [NSArray.IndexesOfObjectsPassingTest]
//   - [NSArray.IndexesOfObjectsWithOptionsPassingTest]
//   - [NSOrderedSet.DifferenceFromOrderedSetWithOptionsUsingEquivalenceTest]
//   - [NSOrderedSet.EnumerateObjectsAtIndexesOptionsUsingBlock]
//   - [NSOrderedSet.EnumerateObjectsUsingBlock]
//   - [NSOrderedSet.EnumerateObjectsWithOptionsUsingBlock]
//   - [NSOrderedSet.IndexOfObjectAtIndexesOptionsPassingTest]
//   - [NSOrderedSet.IndexOfObjectPassingTest]
//   - [NSOrderedSet.IndexOfObjectWithOptionsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsAtIndexesOptionsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsWithOptionsPassingTest]
//   - [NSSet.EnumerateObjectsUsingBlock]
//   - [NSSet.EnumerateObjectsWithOptionsUsingBlock]
//   - [NSSet.ObjectsPassingTest]
//   - [NSSet.ObjectsWithOptionsPassingTest]
type ObjectTypeHandler = func(objectivec.IObject)

// NewObjectTypeBlock wraps a Go [ObjectTypeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSArray.DifferenceFromArrayWithOptionsUsingEquivalenceTest]
//   - [NSArray.EnumerateObjectsAtIndexesOptionsUsingBlock]
//   - [NSArray.EnumerateObjectsUsingBlock]
//   - [NSArray.EnumerateObjectsWithOptionsUsingBlock]
//   - [NSArray.IndexOfObjectAtIndexesOptionsPassingTest]
//   - [NSArray.IndexOfObjectPassingTest]
//   - [NSArray.IndexOfObjectWithOptionsPassingTest]
//   - [NSArray.IndexesOfObjectsAtIndexesOptionsPassingTest]
//   - [NSArray.IndexesOfObjectsPassingTest]
//   - [NSArray.IndexesOfObjectsWithOptionsPassingTest]
//   - [NSOrderedSet.DifferenceFromOrderedSetWithOptionsUsingEquivalenceTest]
//   - [NSOrderedSet.EnumerateObjectsAtIndexesOptionsUsingBlock]
//   - [NSOrderedSet.EnumerateObjectsUsingBlock]
//   - [NSOrderedSet.EnumerateObjectsWithOptionsUsingBlock]
//   - [NSOrderedSet.IndexOfObjectAtIndexesOptionsPassingTest]
//   - [NSOrderedSet.IndexOfObjectPassingTest]
//   - [NSOrderedSet.IndexOfObjectWithOptionsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsAtIndexesOptionsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsPassingTest]
//   - [NSOrderedSet.IndexesOfObjectsWithOptionsPassingTest]
//   - [NSSet.EnumerateObjectsUsingBlock]
//   - [NSSet.EnumerateObjectsWithOptionsUsingBlock]
//   - [NSSet.ObjectsPassingTest]
//   - [NSSet.ObjectsWithOptionsPassingTest]
func NewObjectTypeBlock(handler ObjectTypeHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.IObject) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// OrderedCollectionChangeHandler handles A block receives an ordered collection change and returns an updated change.
//
// Used by:
//   - [NSOrderedCollectionDifference.DifferenceByTransformingChangesWithBlock]
type OrderedCollectionChangeHandler = func(*NSOrderedCollectionChange)

// NewOrderedCollectionChangeBlock wraps a Go [OrderedCollectionChangeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSOrderedCollectionDifference.DifferenceByTransformingChangesWithBlock]
func NewOrderedCollectionChangeBlock(handler OrderedCollectionChangeHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSOrderedCollectionChange
		if resultID != 0 {
			v := NSOrderedCollectionChangeFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// RangeHandler handles The closure or block to apply to ranges of attributes in the attributed string, taking three arguments:
//   - tag: The located linguistic tag.
//   - tokenRange: The range of the linguistic tag.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<doc://com.apple.documentation/documentation/Swift/true>) to stop further processing of the set. The `stop` argument is an out-only argument. You should only ever set this Boolean to [true](<doc://com.apple.documentation/documentation/Swift/true>) within the block.
//
// Used by:
//   - [NSAttributedString.EnumerateAttributesInRangeOptionsUsingBlock]
//   - [NSLinguisticTagger.EnumerateTagsForStringRangeUnitSchemeOptionsOrthographyUsingBlock]
//   - [NSLinguisticTagger.EnumerateTagsInRangeSchemeOptionsUsingBlock]
//   - [NSLinguisticTagger.EnumerateTagsInRangeUnitSchemeOptionsUsingBlock]
//   - [NSString.EnumerateLinguisticTagsInRangeSchemeOptionsOrthographyUsingBlock]
//   - [NSString.EnumerateSubstringsInRangeOptionsUsingBlock]
type RangeHandler = func(*INSDictionary)

// StringHandler handles The block executed for the enumeration.
//   - line: The current line of the string being enumerated. The line contains just the contents of the line, without the line terminators. See [getLineStart(_:end:contentsEnd:for:)](<doc://com.apple.foundation/documentation/Foundation/NSString/getLineStart(_:end:contentsEnd:for:)>) for a discussion of line terminators.
//   - stop: A reference to a Boolean value that the block can use to stop the enumeration by setting `*stop = YES`; it should not touch `*stop` otherwise.
//
// Used by:
//   - [NSString.EnumerateLinesUsingBlock]
type StringHandler = func(*string)

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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSTask
		if resultID != 0 {
			v := NSTaskFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// TextCheckingResultMatchingFlagsHandler handles The Block enumerates the matches of the regular expression in the string.
//   - result: An [NSTextCheckingResult](<doc://com.apple.foundation/documentation/Foundation/NSTextCheckingResult>) specifying the match. This result gives the overall matched range via its [range](<doc://com.apple.foundation/documentation/Foundation/NSTextCheckingResult/range>) property, and the range of each individual capture group via its [range(at:)](<doc://com.apple.foundation/documentation/Foundation/NSTextCheckingResult/range(at:)>) method. The range {[NSNotFound], 0} is returned if one of the capture groups did not participate in this particular match.
//   - flags: The current state of the matching progress. See [NSRegularExpression.MatchingFlags](<doc://com.apple.foundation/documentation/Foundation/NSRegularExpression/MatchingFlags>) for the possible values.
//   - stop: A reference to a Boolean value. The Block can set the value to [true](<doc://com.apple.documentation/documentation/Swift/true>) to stop further processing of the array. The stop argument is an out-only argument. You should only ever set this Boolean to [true](<doc://com.apple.documentation/documentation/Swift/true>) within the Block.
//
// Used by:
//   - [NSRegularExpression.EnumerateMatchesInStringOptionsRangeUsingBlock]
type TextCheckingResultMatchingFlagsHandler = func(*NSTextCheckingResult, NSMatchingFlags)

// NewTextCheckingResultMatchingFlagsBlock wraps a Go [TextCheckingResultMatchingFlagsHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSRegularExpression.EnumerateMatchesInStringOptionsRangeUsingBlock]
func NewTextCheckingResultMatchingFlagsBlock(handler TextCheckingResultMatchingFlagsHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *NSTextCheckingResult
		if resultID != 0 {
			v := NSTextCheckingResultFromID(resultID)
			result = &v
		}
		var extra0 NSMatchingFlags = NSMatchingFlags(extra0ID)
		handler(result, extra0)
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSTimer
		if resultID != 0 {
			v := NSTimerFromID(resultID)
			result = &v
		}
		handler(result)
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSURLCredential
		if resultID != 0 {
			v := NSURLCredentialFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLErrorHandler handles An optional error handler block for the file manager to call when an error occurs.
//   - url: An [NSURL](<doc://com.apple.foundation/documentation/Foundation/NSURL>) object that identifies the item for which the error occurred.
//   - error: An [NSError](<doc://com.apple.foundation/documentation/Foundation/NSError>) object that contains information about the error.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFileManager.EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler]
//   - [NSItemProvider.LoadFileRepresentationForContentTypeOpenInPlaceCompletionHandler]
//   - [NSItemProvider.LoadFileRepresentationForTypeIdentifierCompletionHandler]
//   - [NSItemProvider.LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler]
type URLErrorHandler = func(*NSURL, error)

// NewURLErrorBlock wraps a Go [URLErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFileManager.EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler]
//   - [NSItemProvider.LoadFileRepresentationForContentTypeOpenInPlaceCompletionHandler]
//   - [NSItemProvider.LoadFileRepresentationForTypeIdentifierCompletionHandler]
//   - [NSItemProvider.LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler]
func NewURLErrorBlock(handler URLErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSURL
		if resultID != 0 {
			v := NSURLFromID(resultID)
			result = &v
		}
		var nserr *NSError
		if errID != 0 {
			e := NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, NSErrorToError(nserr))
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSURL
		if resultID != 0 {
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSURLRequest
		if resultID != 0 {
			v := NSURLRequestFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLSessionAuthChallengeDispositionURLCredentialHandler handles A handler that your delegate method must call.
//
// Used by:
//   - [NSURLSessionDelegate.URLSessionDidReceiveChallengeCompletionHandler]
//   - [NSURLSessionTaskDelegate.URLSessionTaskDidReceiveChallengeCompletionHandler]
type URLSessionAuthChallengeDispositionURLCredentialHandler = func(NSURLSessionAuthChallengeDisposition, *NSURLCredential)

// NewURLSessionAuthChallengeDispositionURLCredentialBlock wraps a Go [URLSessionAuthChallengeDispositionURLCredentialHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSessionDelegate.URLSessionDidReceiveChallengeCompletionHandler]
//   - [NSURLSessionTaskDelegate.URLSessionTaskDidReceiveChallengeCompletionHandler]
func NewURLSessionAuthChallengeDispositionURLCredentialBlock(handler URLSessionAuthChallengeDispositionURLCredentialHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result NSURLSessionAuthChallengeDisposition = NSURLSessionAuthChallengeDisposition(resultID)
		var extra0 *NSURLCredential
		if extra0ID != 0 {
			v := NSURLCredentialFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLSessionDelayedRequestDispositionURLRequestHandler handles A completion handler to perform the request.
//
// Used by:
//   - [NSURLSessionTaskDelegate.URLSessionTaskWillBeginDelayedRequestCompletionHandler]
type URLSessionDelayedRequestDispositionURLRequestHandler = func(NSURLSessionDelayedRequestDisposition, *NSURLRequest)

// NewURLSessionDelayedRequestDispositionURLRequestBlock wraps a Go [URLSessionDelayedRequestDispositionURLRequestHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSURLSessionTaskDelegate.URLSessionTaskWillBeginDelayedRequestCompletionHandler]
func NewURLSessionDelayedRequestDispositionURLRequestBlock(handler URLSessionDelayedRequestDispositionURLRequestHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result NSURLSessionDelayedRequestDisposition = NSURLSessionDelayedRequestDisposition(resultID)
		var extra0 *NSURLRequest
		if extra0ID != 0 {
			v := NSURLRequestFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result NSURLSessionResponseDisposition = NSURLSessionResponseDisposition(resultID)
		handler(result)
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSURLSessionWebSocketMessage
		if resultID != 0 {
			v := NSURLSessionWebSocketMessageFromID(resultID)
			result = &v
		}
		var nserr *NSError
		if errID != 0 {
			e := NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, NSErrorToError(nserr))
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *NSURL
		if resultID != 0 {
			v := NSURLFromID(resultID)
			result = &v
		}
		var extra0 *NSURL
		if extra0ID != 0 {
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *NSURL
		if resultID != 0 {
			v := NSURLFromID(resultID)
			result = &v
		}
		var extra0 *NSURLResponse
		if extra0ID != 0 {
			v := NSURLResponseFromID(extra0ID)
			extra0 = &v
		}
		var nserr *NSError
		if errID != 0 {
			e := NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, extra0, NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles The block to add to the new block operation object’s list.
//   - completionHandler: A completion handler block. The batch accessor must call the completion handler when it has finished its read and write calls.
//
// Used by:
//   - [NSBlockOperation.AddExecutionBlock]
//   - [NSBlockOperation.BlockOperationWithBlock]
//   - [NSExtensionContext.LoadBroadcastingApplicationInfoWithCompletion]
//   - [NSFileCoordinator.PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor]
//   - [NSFilePresenter.RelinquishPresentedItemToReader]
//   - [NSFilePresenter.RelinquishPresentedItemToWriter]
//   - [NSItemProvider.RegisterCloudKitShareWithPreparationHandler]
//   - [NSItemProvider.RegisterDataRepresentationForContentTypeVisibilityLoadHandler]
//   - [NSItemProvider.RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler]
//   - [NSItemProvider.RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandler]
//   - [NSItemProvider.RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler]
//   - [NSItemProvider.RegisterObjectOfClassVisibilityLoadHandler]
//   - [NSOperationQueue.AddBarrierBlock]
//   - [NSOperationQueue.AddOperationWithBlock]
//   - [NSProcessInfo.PerformActivityWithOptionsReasonUsingBlock]
//   - [NSProgress.PerformAsCurrentWithPendingUnitCountUsingBlock]
//   - [NSRunLoop.PerformBlock]
//   - [NSRunLoop.PerformInModesBlock]
//   - [NSThread.DetachNewThreadWithBlock]
//   - [NSThread.InitWithBlock]
//   - [NSURLSession.FlushWithCompletionHandler]
//   - [NSURLSession.GetTasksWithCompletionHandler]
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
//   - [NSExtensionContext.LoadBroadcastingApplicationInfoWithCompletion]
//   - [NSFileCoordinator.PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor]
//   - [NSFilePresenter.RelinquishPresentedItemToReader]
//   - [NSFilePresenter.RelinquishPresentedItemToWriter]
//   - [NSItemProvider.RegisterCloudKitShareWithPreparationHandler]
//   - [NSItemProvider.RegisterDataRepresentationForContentTypeVisibilityLoadHandler]
//   - [NSItemProvider.RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler]
//   - [NSItemProvider.RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandler]
//   - [NSItemProvider.RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler]
//   - [NSItemProvider.RegisterObjectOfClassVisibilityLoadHandler]
//   - [NSOperationQueue.AddBarrierBlock]
//   - [NSOperationQueue.AddOperationWithBlock]
//   - [NSProcessInfo.PerformActivityWithOptionsReasonUsingBlock]
//   - [NSProgress.PerformAsCurrentWithPendingUnitCountUsingBlock]
//   - [NSRunLoop.PerformBlock]
//   - [NSRunLoop.PerformInModesBlock]
//   - [NSThread.DetachNewThreadWithBlock]
//   - [NSThread.InitWithBlock]
//   - [NSURLSession.FlushWithCompletionHandler]
//   - [NSURLSession.GetTasksWithCompletionHandler]
//   - [NSURLSession.ResetWithCompletionHandler]
//   - [NSUserActivity.DeleteAllSavedUserActivitiesWithCompletionHandler]
//   - [NSUserActivity.DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler]
//   - [NSXPCConnection.ScheduleSendBarrierBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
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
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSXPCConnection
		if resultID != 0 {
			v := NSXPCConnectionFromID(resultID)
			result = &v
		}
		var nserr *NSError
		if errID != 0 {
			e := NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// unicharHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSConstantString.InitWithCharactersNoCopyLengthDeallocator]
//   - [NSMutableString.InitWithCharactersNoCopyLengthDeallocator]
//   - [NSSimpleCString.InitWithCharactersNoCopyLengthDeallocator]
//   - [NSString.InitWithCharactersNoCopyLengthDeallocator]
type unicharHandler = func(*uint16)


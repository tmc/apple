// Code generated from Apple documentation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

type CompletionHandler = func(NSAttributedString, objectivec.IObject, objectivec.IObject)







type NSAppleEventManagerSuspensionID = uintptr







type NSAttributedStringFormattingContextKey = string







type NSAttributedStringKey = string







type NSBackgroundActivityCompletionHandler = func(NSBackgroundActivityResult)







type NSCalendarIdentifier = string







type NSComparator = func(objectivec.IObject, objectivec.IObject) NSComparisonResult







type NSDistributedNotificationCenterType = string







type NSErrorDomain = string







type NSErrorUserInfoKey = string







type NSExceptionName = string







type NSFileAttributeKey = string







type NSFileAttributeType = string







type NSFileProtectionType = string







type NSFileProviderServiceName = string







type NSHTTPCookiePropertyKey = string







type NSHTTPCookieStringPolicy = string







type NSHashTableOptions = uint







type NSItemProviderCompletionHandler = func(NSSecureCoding, NSError)







type NSItemProviderLoadHandler = func(func(objc.ID, *NSError), objc.Class, *NSDictionary)







type NSKeyValueChangeKey = string







type NSKeyValueOperator = string







type NSLinguisticTag = string







type NSLinguisticTagScheme = string







type NSLocaleKey = string







type NSMapTableOptions = uint







type NSNotificationName = string







type NSPoint = corefoundation.CGPoint







type NSPointArray = *corefoundation.CGPoint







type NSPointPointer = *corefoundation.CGPoint







type NSProgressFileOperationKind = string







type NSProgressKind = string







type NSProgressUnpublishingHandler = func()







type NSProgressUserInfoKey = string







type NSPropertyListReadOptions = NSPropertyListMutabilityOptions







type NSPropertyListWriteOptions = uint








type NSRangePointer = *NSRange







type NSRect = corefoundation.CGRect







type NSRectArray = *corefoundation.CGRect







type NSRectPointer = *corefoundation.CGRect







type NSRunLoopMode = string







type NSSize = corefoundation.CGSize







type NSSizeArray = *corefoundation.CGSize







type NSSizePointer = *corefoundation.CGSize







type NSSocketNativeHandle = int







type NSStreamNetworkServiceTypeValue = string







type NSStreamPropertyKey = string







type NSStreamSOCKSProxyConfiguration = string







type NSStreamSOCKSProxyVersion = string







type NSStreamSocketSecurityLevel = string







type NSStringEncoding = uint







type NSStringEncodingDetectionOptionsKey = string







type NSStringTransform = string







type NSTextCheckingKey = string







type NSTextCheckingTypes = uint64







type NSTimeInterval = float64







type NSURLBookmarkFileCreationOptions = uint







type NSURLFileProtectionType = string







type NSURLFileResourceType = string







type NSURLResourceKey = string







type NSURLThumbnailDictionaryItem = string







type NSURLUbiquitousItemDownloadingStatus = string







type NSURLUbiquitousSharedItemPermissions = string







type NSURLUbiquitousSharedItemRole = string







type NSUncaughtExceptionHandler = objectivec.IObject







type NSUndoManagerUserInfoKey = string







type NSUserActivityPersistentIdentifier = string







type NSUserAppleScriptTaskCompletionHandler = func(NSAppleEventDescriptor, NSError)







type NSUserAutomatorTaskCompletionHandler = func(objectivec.IObject, NSError)







type NSUserScriptTaskCompletionHandler = func(NSError)







type NSUserUnixTaskCompletionHandler = func(NSError)







type NSValueTransformerName = string







type NSZone = unsafe.Pointer







type Unichar = uint16





// UIEdgeInsets is modeled as a CGFloat-based edge inset value for cross-platform NSValue helpers.

type UIEdgeInsets = struct{ Top, Left, Bottom, Right float64 }




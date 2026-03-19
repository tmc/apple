// Code generated from Apple documentation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CompletionHandler is a completion handler for getting an asynchronous attributed string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/CompletionHandler
type CompletionHandler = func(NSAttributedString, objectivec.IObject, objectivec.IObject)

// NSAppleEventManagerSuspensionID is identifies an Apple event whose handling has been suspended. Can be used to resume handling of the Apple event.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/SuspensionID
type NSAppleEventManagerSuspensionID = uintptr

// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringFormattingContextKey
type NSAttributedStringFormattingContextKey = string

// NSAttributedStringKey is the attributes you apply to ranges of characters in an attributed string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
type NSAttributedStringKey = string

// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/CompletionHandler
type NSBackgroundActivityCompletionHandler = func(NSBackgroundActivityResult)

// NSCalendarIdentifier is the supported calendar types.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier
type NSCalendarIdentifier = string

// NSComparator is defines the signature for a block object used for comparison operations.
//
// See: https://developer.apple.com/documentation/Foundation/Comparator
type NSComparator = func(objectivec.IObject, objectivec.IObject) NSComparisonResult

// NSDistributedNotificationCenterType is this constant specifies the notification center type.
//
// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/CenterType
type NSDistributedNotificationCenterType = string

// See: https://developer.apple.com/documentation/Foundation/NSErrorDomain
type NSErrorDomain = string

// NSErrorUserInfoKey is these keys may exist in the user info dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSError/UserInfoKey
type NSErrorUserInfoKey = string

// See: https://developer.apple.com/documentation/Foundation/NSExceptionName
type NSExceptionName = string

// NSFileAttributeKey is keys in dictionaries used to get and set file attributes.
//
// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey
type NSFileAttributeKey = string

// NSFileAttributeType is values representing a file’s type attribute.
//
// See: https://developer.apple.com/documentation/Foundation/FileAttributeType
type NSFileAttributeType = string

// NSFileProtectionType is protection level values that can be associated with a file attribute key.
//
// See: https://developer.apple.com/documentation/Foundation/FileProtectionType
type NSFileProtectionType = string

// NSFileProviderServiceName is the name used to identify a File Provider service.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileProviderServiceName
type NSFileProviderServiceName = string

// NSHTTPCookiePropertyKey is constants that define the supported keys in a cookie attributes dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey
type NSHTTPCookiePropertyKey = string

// NSHTTPCookieStringPolicy is values that indicate whether to restrict the cookie to requests sent back to the same site that created it.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStringPolicy
type NSHTTPCookieStringPolicy = string

// NSHashTableOptions is components in a bit-field to specify the behavior of elements in an [NSHashTable] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashTableOptions
type NSHashTableOptions = uint

// NSItemProviderCompletionHandler is a block that receives the item provider’s data.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/CompletionHandler
type NSItemProviderCompletionHandler = func(NSSecureCoding, NSError)

// NSItemProviderLoadHandler is a block that loads the item provider’s data and coerces it to the specified type.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/LoadHandler
type NSItemProviderLoadHandler = func()

// NSKeyValueChangeKey is the keys that can appear in the change dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyValueChangeKey
type NSKeyValueChangeKey = string

// NSKeyValueOperator is these constants define the available array operators. See [Using Collection Operators] for more information.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator
type NSKeyValueOperator = string

// NSLinguisticTag is a token, lexical class, name, lemma, language, or script returned by a linguistic tagger for natural language text.
//
// See: https://developer.apple.com/documentation/Foundation/NSLinguisticTag
type NSLinguisticTag = string

// NSLinguisticTagScheme is constants for the tag schemes specified when initializing a linguistic tagger.
//
// See: https://developer.apple.com/documentation/Foundation/NSLinguisticTagScheme
type NSLinguisticTagScheme = string

// NSLocaleKey is the keys used to access components of a locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key
type NSLocaleKey = string

// NSMapTableOptions is constants used as components in a bitfield to specify the behavior of elements (keys and values) in an [NSMapTable] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapTableOptions
type NSMapTableOptions = uint

// NSNotificationName is a structure that defines the name of a notification.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct
type NSNotificationName = string

// NSPoint is a point in a Cartesian coordinate system.
//
// See: https://developer.apple.com/documentation/Foundation/NSPoint
type NSPoint = corefoundation.CGPoint

// NSPointArray is type indicating a parameter is array of [NSPoint] structures.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointArray
type NSPointArray = *corefoundation.CGPoint

// NSPointPointer is type indicating a parameter is a pointer to an [NSPoint] structure.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointPointer
type NSPointPointer = *corefoundation.CGPoint

// NSProgressFileOperationKind is the kind of file operation.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/FileOperationKind-swift.struct
type NSProgressFileOperationKind = string

// NSProgressKind is an object that represents the kind of progress.
//
// See: https://developer.apple.com/documentation/Foundation/ProgressKind
type NSProgressKind = string

// NSProgressUnpublishingHandler is a block that the system calls when an observed progress object terminates the subscription.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/UnpublishingHandler
type NSProgressUnpublishingHandler = func()

// NSProgressUserInfoKey is keys for the user info dictionary that affect the autogenerated localized additional description string.
//
// See: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey
type NSProgressUserInfoKey = string

// NSPropertyListReadOptions is the only read options supported are described in [PropertyListSerialization.MutabilityOptions].
//
// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/ReadOptions
type NSPropertyListReadOptions = NSPropertyListMutabilityOptions

// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/WriteOptions
type NSPropertyListWriteOptions = uint

// NSRangePointer is type indicating a parameter is a pointer to an [NSRange] structure.
//
// See: https://developer.apple.com/documentation/Foundation/NSRangePointer
type NSRangePointer = *NSRange

// NSRect is a rectangle.
//
// See: https://developer.apple.com/documentation/Foundation/NSRect
type NSRect = corefoundation.CGRect

// NSRectArray is type indicating a parameter is array of [NSRect] structures.
//
// See: https://developer.apple.com/documentation/Foundation/NSRectArray
type NSRectArray = *corefoundation.CGRect

// NSRectPointer is type indicating a parameter is a pointer to an [NSRect] structure.
//
// See: https://developer.apple.com/documentation/Foundation/NSRectPointer
type NSRectPointer = *corefoundation.CGRect

// NSRunLoopMode is modes that a run loop operates in.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/Mode
type NSRunLoopMode = string

// NSSize is a two-dimensional size.
//
// See: https://developer.apple.com/documentation/Foundation/NSSize
type NSSize = corefoundation.CGSize

// NSSizeArray is type indicating a parameter is an array of [NSSize] structures.
//
// See: https://developer.apple.com/documentation/Foundation/NSSizeArray
type NSSizeArray = *corefoundation.CGSize

// NSSizePointer is type indicating parameter is a pointer to an [NSSize] structure.
//
// See: https://developer.apple.com/documentation/Foundation/NSSizePointer
type NSSizePointer = *corefoundation.CGSize

// NSSocketNativeHandle is type for the platform-specific native socket handle.
//
// See: https://developer.apple.com/documentation/Foundation/SocketNativeHandle
type NSSocketNativeHandle = int

// NSStreamNetworkServiceTypeValue is [NSStream] defines these string constants for specifying the service type of a stream.
//
// See: https://developer.apple.com/documentation/Foundation/StreamNetworkServiceTypeValue
type NSStreamNetworkServiceTypeValue = string

// NSStreamPropertyKey is [NSStream] defines these string constants as keys for accessing stream properties using [property(forKey:)] and setting properties with [setProperty(_:forKey:)]:
//
// See: https://developer.apple.com/documentation/Foundation/Stream/PropertyKey
type NSStreamPropertyKey = string

// See: https://developer.apple.com/documentation/Foundation/StreamSOCKSProxyConfiguration
type NSStreamSOCKSProxyConfiguration = string

// See: https://developer.apple.com/documentation/Foundation/StreamSOCKSProxyVersion
type NSStreamSOCKSProxyVersion = string

// NSStreamSocketSecurityLevel is [NSStream] defines these string constants for specifying the secure-socket layer (SSL) security level.
//
// See: https://developer.apple.com/documentation/Foundation/StreamSocketSecurityLevel
type NSStreamSocketSecurityLevel = string

// NSStringEncoding is the following constants are provided by [NSString] as possible string encodings.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringEncoding
type NSStringEncoding = uint

// See: https://developer.apple.com/documentation/Foundation/StringEncodingDetectionOptionsKey
type NSStringEncodingDetectionOptionsKey = string

// NSStringTransform is constants representing an ICU string transform.
//
// See: https://developer.apple.com/documentation/Foundation/StringTransform
type NSStringTransform = string

// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey
type NSTextCheckingKey = string

// NSTextCheckingTypes is defines the types of checking that are available. These values can be combined using the C-bitwise OR operator. The system supports its own internal types, and the user can extend those types by subclassing [NSTextCheckingResult] and adding their own custom types.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingTypes
type NSTextCheckingTypes = uint64

// NSTimeInterval is a number of seconds.
//
// See: https://developer.apple.com/documentation/Foundation/TimeInterval
type NSTimeInterval = float64

// NSURLBookmarkFileCreationOptions is options used when creating file bookmark data
//
// See: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkFileCreationOptions
type NSURLBookmarkFileCreationOptions = uint

// NSURLFileProtectionType is protection-level values for a URL resource key.
//
// See: https://developer.apple.com/documentation/Foundation/URLFileProtection
type NSURLFileProtectionType = string

// NSURLFileResourceType is possible values for the type of file resource.
//
// See: https://developer.apple.com/documentation/Foundation/URLFileResourceType
type NSURLFileResourceType = string

// NSURLResourceKey is keys that apply to file system URLs.
//
// See: https://developer.apple.com/documentation/Foundation/URLResourceKey
type NSURLResourceKey = string

// NSURLThumbnailDictionaryItem is possible keys for the [thumbnailDictionaryKey] dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/URLThumbnailDictionaryItem
type NSURLThumbnailDictionaryItem = string

// NSURLUbiquitousItemDownloadingStatus is values that describe the iCloud storage state of a file.
//
// See: https://developer.apple.com/documentation/Foundation/URLUbiquitousItemDownloadingStatus
type NSURLUbiquitousItemDownloadingStatus = string

// NSURLUbiquitousSharedItemPermissions is the key for the permissions of a shared item.
//
// See: https://developer.apple.com/documentation/Foundation/URLUbiquitousSharedItemPermissions
type NSURLUbiquitousSharedItemPermissions = string

// NSURLUbiquitousSharedItemRole is the key for the role of a shared item.
//
// See: https://developer.apple.com/documentation/Foundation/URLUbiquitousSharedItemRole
type NSURLUbiquitousSharedItemRole = string

// See: https://developer.apple.com/documentation/Foundation/NSUncaughtExceptionHandler
type NSUncaughtExceptionHandler = objectivec.IObject

// NSUndoManagerUserInfoKey is an extensible namespace for undo and redo user info keys.
//
// See: https://developer.apple.com/documentation/Foundation/UndoManager/UserInfoKey
type NSUndoManagerUserInfoKey = string

// NSUserActivityPersistentIdentifier is the type that defines a persistent identifier value for a user activity.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivityPersistentIdentifier
type NSUserActivityPersistentIdentifier = string

// NSUserAppleScriptTaskCompletionHandler is implement this block to retrieve the result of the AppleScript executed by [execute(withAppleEvent:completionHandler:)].
//
// See: https://developer.apple.com/documentation/Foundation/NSUserAppleScriptTask/CompletionHandler
type NSUserAppleScriptTaskCompletionHandler = func(NSAppleEventDescriptor, NSError)

// NSUserAutomatorTaskCompletionHandler is implement this block to retrieve the output of the Automator workflow executed by [execute(withInput:completionHandler:)].
//
// See: https://developer.apple.com/documentation/Foundation/NSUserAutomatorTask/CompletionHandler
type NSUserAutomatorTaskCompletionHandler = func(objectivec.IObject, NSError)

// NSUserScriptTaskCompletionHandler is implement this block to retrieve the error of the script executed by [execute(completionHandler:)].
//
// See: https://developer.apple.com/documentation/Foundation/NSUserScriptTask/CompletionHandler
type NSUserScriptTaskCompletionHandler = func(NSError)

// NSUserUnixTaskCompletionHandler is implement this block to retrieve an error from the Unix scripted executed by [execute(withArguments:completionHandler:)].
//
// See: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/CompletionHandler
type NSUserUnixTaskCompletionHandler = func(NSError)

// NSValueTransformerName is named value transformers defined by [NSValueTransformer].
//
// See: https://developer.apple.com/documentation/Foundation/NSValueTransformerName
type NSValueTransformerName = string

// NSZone is a type used to identify and manage memory zones.
//
// See: https://developer.apple.com/documentation/Foundation/NSZone
type NSZone = unsafe.Pointer

// Unichar is type for UTF-16 code units.
//
// See: https://developer.apple.com/documentation/Foundation/unichar
type Unichar = uint16

// UIEdgeInsets is modeled as a CGFloat-based edge inset value for cross-platform NSValue helpers.

type UIEdgeInsets = struct{ Top, Left, Bottom, Right float64 }


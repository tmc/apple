// Code generated from Apple documentation. DO NOT EDIT.

package corefoundation

import (
	"unsafe"
)

// CFAbsoluteTime is type used to represent a specific point in time relative to the absolute reference date of 1 Jan 2001 00:00:00 GMT.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTime
type CFAbsoluteTime = float64

// CFAllocatorAllocateCallBack is a prototype for a function callback that allocates memory of a requested size.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocateCallBack
type CFAllocatorAllocateCallBack = func(allocSize int, hint uint, info unsafe.Pointer) LPVOID

// CFAllocatorCopyDescriptionCallBack is a prototype for a function callback that provides a description of the specified data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorCopyDescriptionCallBack
type CFAllocatorCopyDescriptionCallBack = func(info unsafe.Pointer) CFStringRef

// CFAllocatorDeallocateCallBack is a prototype for a function callback that deallocates a block of memory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorDeallocateCallBack
type CFAllocatorDeallocateCallBack = func(ptr unsafe.Pointer, info unsafe.Pointer)

// CFAllocatorPreferredSizeCallBack is a prototype for a function callback that gives the size of memory likely to be allocated, given a certain request.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorPreferredSizeCallBack
type CFAllocatorPreferredSizeCallBack = func(size int, hint uint, info unsafe.Pointer) int

// CFAllocatorReallocateCallBack is a prototype for a function callback that reallocates memory of a requested size for an existing block of memory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocateCallBack
type CFAllocatorReallocateCallBack = func(ptr unsafe.Pointer, newsize int, hint uint, info unsafe.Pointer) LPVOID

// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocator
type CFAllocatorRef uintptr

// CFAllocatorReleaseCallBack is a prototype for a function callback that releases the given data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReleaseCallBack
type CFAllocatorReleaseCallBack = func(info unsafe.Pointer)

// CFAllocatorRetainCallBack is a prototype for a function callback that retains the given data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorRetainCallBack
type CFAllocatorRetainCallBack = func(info unsafe.Pointer) CFTypeRef

// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorTypeID
type CFAllocatorTypeID = uint64

// CFArrayApplierFunction is prototype of a callback function that may be applied to every value in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayApplierFunction
type CFArrayApplierFunction = func(value unsafe.Pointer, context unsafe.Pointer)

// CFArrayCopyDescriptionCallBack is prototype of a callback function used to get a description of a value in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayCopyDescriptionCallBack
type CFArrayCopyDescriptionCallBack = func(value unsafe.Pointer) CFStringRef

// CFArrayEqualCallBack is prototype of a callback function used to determine if two values in an array are equal.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayEqualCallBack
type CFArrayEqualCallBack = func(value1 unsafe.Pointer, value2 unsafe.Pointer) uint8

// See: https://developer.apple.com/documentation/CoreFoundation/CFArray
type CFArrayRef uintptr

// CFArrayReleaseCallBack is prototype of a callback function used to release a value before it’s removed from an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayReleaseCallBack
type CFArrayReleaseCallBack = func(allocator CFAllocatorRef, value unsafe.Pointer)

// CFArrayRetainCallBack is prototype of a callback function used to retain a value being added to an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayRetainCallBack
type CFArrayRetainCallBack = func(allocator CFAllocatorRef, value unsafe.Pointer) CFTypeRef

// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedString
type CFAttributedStringRef uintptr

// CFBagApplierFunction is prototype of a callback function that may be applied to every value in a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagApplierFunction
type CFBagApplierFunction = func(value unsafe.Pointer, context unsafe.Pointer)

// CFBagCopyDescriptionCallBack is prototype of a callback function used to get a description of a value in a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagCopyDescriptionCallBack
type CFBagCopyDescriptionCallBack = func(value unsafe.Pointer) CFStringRef

// CFBagEqualCallBack is prototype of a callback function used to determine if two values in a bag are equal.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagEqualCallBack
type CFBagEqualCallBack = func(value1 unsafe.Pointer, value2 unsafe.Pointer) uint8

// CFBagHashCallBack is prototype of a callback function invoked to compute a hash code for a value. Hash codes are used when values are accessed, added, or removed from a collection.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagHashCallBack
type CFBagHashCallBack = func(value unsafe.Pointer) uint

// See: https://developer.apple.com/documentation/CoreFoundation/CFBag
type CFBagRef uintptr

// CFBagReleaseCallBack is prototype of a callback function used to release a value before it’s removed from a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagReleaseCallBack
type CFBagReleaseCallBack = func(allocator CFAllocatorRef, value unsafe.Pointer)

// CFBagRetainCallBack is prototype of a callback function used to retain a value being added to a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagRetainCallBack
type CFBagRetainCallBack = func(allocator CFAllocatorRef, value unsafe.Pointer) CFTypeRef

// CFBinaryHeapApplierFunction is callback function used to apply a function to all members of a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapApplierFunction
type CFBinaryHeapApplierFunction = func(val unsafe.Pointer, context unsafe.Pointer)

// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeap
type CFBinaryHeapRef uintptr

// CFBit is a binary value of either `0` or `1`.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBit
type CFBit = uint32

// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVector
type CFBitVectorRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFBoolean
type CFBooleanRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFBundle
type CFBundleRef uintptr

// CFBundleRefNum is type that identifies a distinct reference number for a resource map.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleRefNum
type CFBundleRefNum = int32

// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier
type CFCalendarIdentifier = string

// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendar
type CFCalendarRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSet
type CFCharacterSetRef uintptr

// CFComparatorFunction is callback function that compares two values. You provide a pointer to this callback in certain Core Foundation sorting functions.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFComparatorFunction
type CFComparatorFunction = CFComparisonResult

// See: https://developer.apple.com/documentation/CoreFoundation/CFData
type CFDataRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey
type CFDateFormatterKey = string

// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatter
type CFDateFormatterRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFDate
type CFDateRef uintptr

// CFDictionaryApplierFunction is prototype of a callback function that may be applied to every key-value pair in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryApplierFunction
type CFDictionaryApplierFunction = func(key unsafe.Pointer, value unsafe.Pointer, context unsafe.Pointer)

// CFDictionaryCopyDescriptionCallBack is prototype of a callback function used to get a description of a value or key in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCopyDescriptionCallBack
type CFDictionaryCopyDescriptionCallBack = func(value unsafe.Pointer) CFStringRef

// CFDictionaryEqualCallBack is prototype of a callback function used to determine if two values or keys in a dictionary are equal.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryEqualCallBack
type CFDictionaryEqualCallBack = func(value1 unsafe.Pointer, value2 unsafe.Pointer) uint8

// CFDictionaryHashCallBack is prototype of a callback function invoked to compute a hash code for a key. Hash codes are used when key-value pairs are accessed, added, or removed from a collection.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryHashCallBack
type CFDictionaryHashCallBack = func(value unsafe.Pointer) uint

// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionary
type CFDictionaryRef uintptr

// CFDictionaryReleaseCallBack is prototype of a callback function used to release a key-value pair before it’s removed from a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryReleaseCallBack
type CFDictionaryReleaseCallBack = func(allocator CFAllocatorRef, value unsafe.Pointer)

// CFDictionaryRetainCallBack is prototype of a callback function used to retain a value or key being added to a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryRetainCallBack
type CFDictionaryRetainCallBack = func(allocator CFAllocatorRef, value unsafe.Pointer) CFTypeRef

// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorDomain
type CFErrorDomain = unsafe.Pointer

// See: https://developer.apple.com/documentation/CoreFoundation/CFError
type CFErrorRef uintptr

// CFFileDescriptorCallBack is defines a structure for a callback for a CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorCallBack
type CFFileDescriptorCallBack = func(CFFileDescriptorRef, uint, unsafe.Pointer)

// CFFileDescriptorNativeDescriptor is defines a type for the native file descriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorNativeDescriptor
type CFFileDescriptorNativeDescriptor = int32

// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptor
type CFFileDescriptorRef uintptr

// CFFileSecurityRef is encapsulates a file system object’s security information in a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurity
type CFFileSecurityRef uintptr

// CFHashCode is a type for hash codes returned by the [CFHash] function.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFHashCode
type CFHashCode = uint

// CFIndex is priority values used for kAXPriorityKey.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFIndex
type CFIndex = int

// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleIdentifier
type CFLocaleIdentifier = string

// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey
type CFLocaleKey = string

// See: https://developer.apple.com/documentation/CoreFoundation/CFLocale
type CFLocaleRef uintptr

// CFMachPortCallBack is callback invoked to process a message received on a CFMachPort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCallBack
type CFMachPortCallBack = func(port CFMachPortRef, msg unsafe.Pointer, size int, info unsafe.Pointer)

// CFMachPortInvalidationCallBack is callback invoked when a CFMachPort object is invalidated.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortInvalidationCallBack
type CFMachPortInvalidationCallBack = func(port CFMachPortRef, info unsafe.Pointer)

// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPort
type CFMachPortRef uintptr

// CFMessagePortCallBack is callback invoked to process a message received on a CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCallBack
type CFMessagePortCallBack = func(local CFMessagePortRef, msgid int32, data CFDataRef, info unsafe.Pointer) CFDataRef

// CFMessagePortInvalidationCallBack is callback invoked when a CFMessagePort object is invalidated.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortInvalidationCallBack
type CFMessagePortInvalidationCallBack = func(ms CFMessagePortRef, info unsafe.Pointer)

// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePort
type CFMessagePortRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFMutableArray
type CFMutableArrayRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFMutableAttributedString
type CFMutableAttributedStringRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFMutableBag
type CFMutableBagRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFMutableBitVector
type CFMutableBitVectorRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFMutableCharacterSet
type CFMutableCharacterSetRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFMutableData
type CFMutableDataRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFMutableDictionary
type CFMutableDictionaryRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFMutableSet
type CFMutableSetRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFMutableString
type CFMutableStringRef uintptr

// CFNotificationCallback is callback function invoked for each observer of a notification when the notification is posted.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCallback
type CFNotificationCallback = func(center CFNotificationCenterRef, observer unsafe.Pointer, name CFStringRef, object unsafe.Pointer, userInfo CFDictionaryRef)

// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenter
type CFNotificationCenterRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationName
type CFNotificationName = string

// See: https://developer.apple.com/documentation/CoreFoundation/CFNull
type CFNullRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey
type CFNumberFormatterKey = string

// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatter
type CFNumberFormatterRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFNumber
type CFNumberRef uintptr

// CFOptionFlags is a bitfield used for passing special allocation and other requests into Core Foundation functions.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFOptionFlags
type CFOptionFlags = uint

// CFPlugInDynamicRegisterFunction is a callback which provides a plug-in the opportunity to dynamically register its types with a host.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInDynamicRegisterFunction
type CFPlugInDynamicRegisterFunction = func(plugIn CFBundleRef)

// CFPlugInFactoryFunction is callback function that a plug-in author must implement to create a plug-in instance.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInFactoryFunction
type CFPlugInFactoryFunction = func(allocator CFAllocatorRef, typeUUID CFUUIDRef) LPVOID

// CFPlugInInstanceDeallocateInstanceDataFunction is not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceDeallocateInstanceDataFunction
type CFPlugInInstanceDeallocateInstanceDataFunction = func(unsafe.Pointer)

// CFPlugInInstanceGetInterfaceFunction is not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetInterfaceFunction
type CFPlugInInstanceGetInterfaceFunction = func(CFPlugInInstanceRef, CFStringRef, unsafe.Pointer) uint8

// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstance
type CFPlugInInstanceRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugIn
type CFPlugInRef uintptr

// CFPlugInUnloadFunction is callback function that is called, if present, just before a plug-in’s code is unloaded.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInUnloadFunction
type CFPlugInUnloadFunction = func(plugIn CFBundleRef)

// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyList
type CFPropertyListRef uintptr

// CFReadStreamClientCallBack is callback invoked when certain types of activity takes place on a readable stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamClientCallBack
type CFReadStreamClientCallBack = func(stream CFReadStreamRef, eventType CFStreamEventType, clientCallBackInfo unsafe.Pointer)

// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStream
type CFReadStreamRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopMode
type CFRunLoopMode = uint

// CFRunLoopObserverCallBack is callback invoked when a CFRunLoopObserver object is fired.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverCallBack
type CFRunLoopObserverCallBack = func(observer CFRunLoopObserverRef, activity CFRunLoopActivity, info unsafe.Pointer)

// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserver
type CFRunLoopObserverRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoop
type CFRunLoopRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSource
type CFRunLoopSourceRef uintptr

// CFRunLoopTimerCallBack is callback invoked when a CFRunLoopTimer object fires.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerCallBack
type CFRunLoopTimerCallBack = func(timer CFRunLoopTimerRef, info unsafe.Pointer)

// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimer
type CFRunLoopTimerRef uintptr

// CFSetApplierFunction is prototype of a callback function that may be applied to every value in a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetApplierFunction
type CFSetApplierFunction = func(value unsafe.Pointer, context unsafe.Pointer)

// CFSetCopyDescriptionCallBack is prototype of a callback function used to get a description of a value in a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetCopyDescriptionCallBack
type CFSetCopyDescriptionCallBack = func(value unsafe.Pointer) CFStringRef

// CFSetEqualCallBack is prototype of a callback function used to determine if two values in a set are equal.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetEqualCallBack
type CFSetEqualCallBack = func(value1 unsafe.Pointer, value2 unsafe.Pointer) uint8

// CFSetHashCallBack is prototype of a callback function called to compute a hash code for a value. Hash codes are used when values are accessed, added, or removed from a collection.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetHashCallBack
type CFSetHashCallBack = func(value unsafe.Pointer) uint

// See: https://developer.apple.com/documentation/CoreFoundation/CFSet
type CFSetRef uintptr

// CFSetReleaseCallBack is prototype of a callback function used to release a value before it’s removed from a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetReleaseCallBack
type CFSetReleaseCallBack = func(allocator CFAllocatorRef, value unsafe.Pointer)

// CFSetRetainCallBack is prototype of a callback function used to retain a value being added to a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetRetainCallBack
type CFSetRetainCallBack = func(allocator CFAllocatorRef, value unsafe.Pointer) CFTypeRef

// CFSocketCallBack is callback invoked when certain types of activity takes place on a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBack
type CFSocketCallBack = func(s CFSocketRef, callbackType CFSocketCallBackType, address CFDataRef, data unsafe.Pointer, info unsafe.Pointer)

// CFSocketNativeHandle is type for the platform-specific native socket handle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketNativeHandle
type CFSocketNativeHandle = int32

// See: https://developer.apple.com/documentation/CoreFoundation/CFSocket
type CFSocketRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamPropertyKey
type CFStreamPropertyKey = string

// CFStringEncoding is an integer type for constants used to specify supported string encodings in various CFString functions.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringEncoding
type CFStringEncoding = uint32

// See: https://developer.apple.com/documentation/CoreFoundation/CFString
type CFStringRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizer
type CFStringTokenizerRef uintptr

// CFTimeInterval is type used to represent elapsed time in seconds.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeInterval
type CFTimeInterval = float64

// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZone
type CFTimeZoneRef uintptr

// CFTreeApplierFunction is type of the callback function used by the CFTree apply function.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeApplierFunction
type CFTreeApplierFunction = func(value unsafe.Pointer, context unsafe.Pointer)

// CFTreeCopyDescriptionCallBack is callback function used to provide a description of the program-defined information pointer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeCopyDescriptionCallBack
type CFTreeCopyDescriptionCallBack = func(info unsafe.Pointer) CFStringRef

// See: https://developer.apple.com/documentation/CoreFoundation/CFTree
type CFTreeRef uintptr

// CFTreeReleaseCallBack is callback function used to release a previously retained program-defined information pointer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeReleaseCallBack
type CFTreeReleaseCallBack = func(info unsafe.Pointer)

// CFTreeRetainCallBack is callback function used to retain a program-defined information pointer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeRetainCallBack
type CFTreeRetainCallBack = func(info unsafe.Pointer) CFTypeRef

// CFTypeID is a type for unique, constant integer values that identify particular Core Foundation opaque types.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTypeID
type CFTypeID = uint

// CFTypeRef is an untyped “generic” reference to any Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTypeRef
type CFTypeRef = unsafe.Pointer

// CFURLBookmarkFileCreationOptions is type for bookmark file creation options.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkFileCreationOptions
type CFURLBookmarkFileCreationOptions = uint

// CFURLEnumeratorRef is a reference to a [CFURLEnumerator] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumerator
type CFURLEnumeratorRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFURL
type CFURLRef uintptr

// See: https://developer.apple.com/documentation/CoreFoundation/CFUUID
type CFUUIDRef uintptr

// CFUserNotificationCallBack is callback invoked when an asynchronous user notification dialog is dismissed.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCallBack
type CFUserNotificationCallBack = func(userNotification CFUserNotificationRef, responseFlags uint)

// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotification
type CFUserNotificationRef uintptr

// CFWriteStreamClientCallBack is callback invoked when certain types of activity takes place on a writable stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamClientCallBack
type CFWriteStreamClientCallBack = func(stream CFWriteStreamRef, eventType CFStreamEventType, clientCallBackInfo unsafe.Pointer)

// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStream
type CFWriteStreamRef uintptr

// CFXMLParserAddChildCallBack is callback function invoked by the parser to notify your application of parent/child relationships between XML structures.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserAddChildCallBack
type CFXMLParserAddChildCallBack = func(parser CFXMLParserRef, parent unsafe.Pointer, child unsafe.Pointer, info unsafe.Pointer)

// CFXMLParserCopyDescriptionCallBack is callback function invoked by the parser when handling the information pointer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCopyDescriptionCallBack
type CFXMLParserCopyDescriptionCallBack = func(info unsafe.Pointer) CFStringRef

// CFXMLParserCreateXMLStructureCallBack is callback function invoked when the parser encounters an XML open tag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCreateXMLStructureCallBack
type CFXMLParserCreateXMLStructureCallBack = func(parser CFXMLParserRef, nodeDesc uintptr, info unsafe.Pointer) LPVOID

// CFXMLParserEndXMLStructureCallBack is callback function invoked by the parser to notify your application that an XML structure (and all its children) have been completely parsed.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserEndXMLStructureCallBack
type CFXMLParserEndXMLStructureCallBack = func(parser CFXMLParserRef, xmlType unsafe.Pointer, info unsafe.Pointer)

// CFXMLParserHandleErrorCallBack is callback function invoked by the parser to notify your application that an error has occurred.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserHandleErrorCallBack
type CFXMLParserHandleErrorCallBack = func(parser CFXMLParserRef, error_ CFXMLParserStatusCode, info unsafe.Pointer) uint8

// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParser
type CFXMLParserRef uintptr

// CFXMLParserReleaseCallBack is callback function invoked by the parser when it wants to release a reference to the information pointer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserReleaseCallBack
type CFXMLParserReleaseCallBack = func(info unsafe.Pointer)

// CFXMLParserResolveExternalEntityCallBack is callback function invoked by the parser to notify your application that an external entity has been referenced.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserResolveExternalEntityCallBack
type CFXMLParserResolveExternalEntityCallBack = func(parser CFXMLParserRef, extID *CFXMLExternalID, info unsafe.Pointer) CFDataRef

// CFXMLParserRetainCallBack is callback function invoked by the parser when it needs another reference to the information pointer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserRetainCallBack
type CFXMLParserRetainCallBack = func(info unsafe.Pointer) CFTypeRef

// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTree
type CFXMLTreeRef uintptr

// CGFloat is the basic type for all floating-point values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CGFloat-c.typealias
type CGFloat = float64

// See: https://developer.apple.com/documentation/CoreFoundation/HRESULT
type HRESULT = int32

// See: https://developer.apple.com/documentation/CoreFoundation/LPVOID
type LPVOID = unsafe.Pointer

// See: https://developer.apple.com/documentation/CoreFoundation/REFIID
type REFIID = CFUUIDBytes

// See: https://developer.apple.com/documentation/CoreFoundation/ULONG
type ULONG = uint32

// CFAttributedString is a bare alias for CFAttributedStringRef.
type CFAttributedString = CFAttributedStringRef

// CFBundle is a bare alias for CFBundleRef.
type CFBundle = CFBundleRef

// CFCharacterSet is a bare alias for CFCharacterSetRef.
type CFCharacterSet = CFCharacterSetRef

// CFLocale is a bare alias for CFLocaleRef.
type CFLocale = CFLocaleRef

// CFMachPort is a bare alias for CFMachPortRef.
type CFMachPort = CFMachPortRef

// CFMutableDictionary is a bare alias for CFMutableDictionaryRef.
type CFMutableDictionary = CFMutableDictionaryRef

// CFSet is a bare alias for CFSetRef.
type CFSet = CFSetRef

// CFUUID is a bare alias for CFUUIDRef.
type CFUUID = CFUUIDRef

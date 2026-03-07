// Code generated from Apple documentation. DO NOT EDIT.

package corefoundation

import (
	"unsafe"
)

type CFAbsoluteTime = float64







type CFAllocatorAllocateCallBack = func(int, uint, unsafe.Pointer) unsafe.Pointer







type CFAllocatorCopyDescriptionCallBack = func(unsafe.Pointer) uintptr







type CFAllocatorDeallocateCallBack = func(unsafe.Pointer, unsafe.Pointer)







type CFAllocatorPreferredSizeCallBack = func(int, uint, unsafe.Pointer) int







type CFAllocatorReallocateCallBack = func(unsafe.Pointer, int, uint, unsafe.Pointer) unsafe.Pointer







type CFAllocatorRef uintptr







type CFAllocatorReleaseCallBack = func(unsafe.Pointer)







type CFAllocatorRetainCallBack = func(unsafe.Pointer) unsafe.Pointer







type CFAllocatorTypeID = uint64







type CFArrayApplierFunction = func(unsafe.Pointer, unsafe.Pointer)







type CFArrayCopyDescriptionCallBack = func(unsafe.Pointer) uintptr







type CFArrayEqualCallBack = func(unsafe.Pointer, unsafe.Pointer) uint8







type CFArrayRef uintptr







type CFArrayReleaseCallBack = func(uintptr, unsafe.Pointer)







type CFArrayRetainCallBack = func(uintptr, unsafe.Pointer) unsafe.Pointer







type CFAttributedStringRef uintptr







type CFBagApplierFunction = func(unsafe.Pointer, unsafe.Pointer)







type CFBagCopyDescriptionCallBack = func(unsafe.Pointer) uintptr







type CFBagEqualCallBack = func(unsafe.Pointer, unsafe.Pointer) uint8







type CFBagHashCallBack = func(unsafe.Pointer) uint







type CFBagRef uintptr







type CFBagReleaseCallBack = func(uintptr, unsafe.Pointer)







type CFBagRetainCallBack = func(uintptr, unsafe.Pointer) unsafe.Pointer







type CFBinaryHeapApplierFunction = func(unsafe.Pointer, unsafe.Pointer)







type CFBinaryHeapRef uintptr







type CFBit = uint32







type CFBitVectorRef uintptr







type CFBooleanRef uintptr







type CFBundleRef uintptr







type CFBundleRefNum = int







type CFCalendarIdentifier = string







type CFCalendarRef uintptr







type CFCharacterSetRef uintptr







type CFComparatorFunction = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) CFComparisonResult







type CFDataRef uintptr







type CFDateFormatterKey = string







type CFDateFormatterRef uintptr







type CFDateRef uintptr







type CFDictionaryApplierFunction = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)







type CFDictionaryCopyDescriptionCallBack = func(unsafe.Pointer) uintptr







type CFDictionaryEqualCallBack = func(unsafe.Pointer, unsafe.Pointer) uint8







type CFDictionaryHashCallBack = func(unsafe.Pointer) uint







type CFDictionaryRef uintptr







type CFDictionaryReleaseCallBack = func(uintptr, unsafe.Pointer)







type CFDictionaryRetainCallBack = func(uintptr, unsafe.Pointer) unsafe.Pointer







type CFErrorDomain = unsafe.Pointer







type CFErrorRef uintptr







type CFFileDescriptorCallBack = func(uintptr, uint, unsafe.Pointer)







type CFFileDescriptorNativeDescriptor = int







type CFFileDescriptorRef uintptr







type CFFileSecurityRef uintptr







type CFHashCode = uint







type CFIndex = int







type CFLocaleIdentifier = string







type CFLocaleKey = string







type CFLocaleRef uintptr







type CFMachPortCallBack = func(uintptr, unsafe.Pointer, int, unsafe.Pointer)







type CFMachPortInvalidationCallBack = func(uintptr, unsafe.Pointer)







type CFMachPortRef uintptr







type CFMessagePortCallBack = func(uintptr, int, uintptr, unsafe.Pointer) uintptr







type CFMessagePortInvalidationCallBack = func(uintptr, unsafe.Pointer)







type CFMessagePortRef uintptr







type CFMutableArrayRef uintptr







type CFMutableAttributedStringRef uintptr







type CFMutableBagRef uintptr







type CFMutableBitVectorRef uintptr







type CFMutableCharacterSetRef uintptr







type CFMutableDataRef uintptr







type CFMutableDictionaryRef uintptr







type CFMutableSetRef uintptr







type CFMutableStringRef uintptr







type CFNotificationCallback = func(uintptr, unsafe.Pointer, uintptr, unsafe.Pointer, uintptr)







type CFNotificationCenterRef uintptr







type CFNotificationName = string







type CFNullRef uintptr







type CFNumberFormatterKey = string







type CFNumberFormatterRef uintptr







type CFNumberRef uintptr







type CFOptionFlags = uint64







type CFPlugInDynamicRegisterFunction = func(uintptr)







type CFPlugInFactoryFunction = func(uintptr, uintptr) unsafe.Pointer







type CFPlugInInstanceDeallocateInstanceDataFunction = func(unsafe.Pointer)







type CFPlugInInstanceGetInterfaceFunction = func(uintptr, uintptr, unsafe.Pointer) uint8







type CFPlugInInstanceRef uintptr







type CFPlugInRef uintptr







type CFPlugInUnloadFunction = func(uintptr)







type CFPropertyListRef uintptr







type CFReadStreamClientCallBack = func(uintptr, CFStreamEventType, unsafe.Pointer)







type CFReadStreamRef uintptr







type CFRunLoopMode = uint







type CFRunLoopObserverCallBack = func(uintptr, CFRunLoopActivity, unsafe.Pointer)







type CFRunLoopObserverRef uintptr







type CFRunLoopRef uintptr







type CFRunLoopSourceRef uintptr







type CFRunLoopTimerCallBack = func(uintptr, unsafe.Pointer)







type CFRunLoopTimerRef uintptr







type CFSetApplierFunction = func(unsafe.Pointer, unsafe.Pointer)







type CFSetCopyDescriptionCallBack = func(unsafe.Pointer) uintptr







type CFSetEqualCallBack = func(unsafe.Pointer, unsafe.Pointer) uint8







type CFSetHashCallBack = func(unsafe.Pointer) uint







type CFSetRef uintptr







type CFSetReleaseCallBack = func(uintptr, unsafe.Pointer)







type CFSetRetainCallBack = func(uintptr, unsafe.Pointer) unsafe.Pointer







type CFSocketCallBack = func(uintptr, CFSocketCallBackType, uintptr, unsafe.Pointer, unsafe.Pointer)







type CFSocketNativeHandle = int







type CFSocketRef uintptr







type CFStreamPropertyKey = string







type CFStringEncoding = uint32







type CFStringRef uintptr







type CFStringTokenizerRef uintptr







type CFTimeInterval = float64







type CFTimeZoneRef uintptr







type CFTreeApplierFunction = func(unsafe.Pointer, unsafe.Pointer)







type CFTreeCopyDescriptionCallBack = func(unsafe.Pointer) uintptr







type CFTreeRef uintptr







type CFTreeReleaseCallBack = func(unsafe.Pointer)







type CFTreeRetainCallBack = func(unsafe.Pointer) unsafe.Pointer







type CFTypeID = uint







type CFTypeRef uintptr







type CFURLBookmarkFileCreationOptions = uint64







type CFURLEnumeratorRef uintptr







type CFURLRef uintptr







type CFUUIDRef uintptr







type CFUserNotificationCallBack = func(uintptr, uint)







type CFUserNotificationRef uintptr







type CFWriteStreamClientCallBack = func(uintptr, CFStreamEventType, unsafe.Pointer)







type CFWriteStreamRef uintptr







type CFXMLParserAddChildCallBack = func(uintptr, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)







type CFXMLParserCopyDescriptionCallBack = func(unsafe.Pointer) uintptr







type CFXMLParserCreateXMLStructureCallBack = func(uintptr, uintptr, unsafe.Pointer) unsafe.Pointer







type CFXMLParserEndXMLStructureCallBack = func(uintptr, unsafe.Pointer, unsafe.Pointer)







type CFXMLParserHandleErrorCallBack = func(uintptr, CFXMLParserStatusCode, unsafe.Pointer) uint8







type CFXMLParserRef uintptr







type CFXMLParserReleaseCallBack = func(unsafe.Pointer)







type CFXMLParserResolveExternalEntityCallBack = func(uintptr, *CFXMLExternalID, unsafe.Pointer) uintptr







type CFXMLParserRetainCallBack = func(unsafe.Pointer) unsafe.Pointer







type CFXMLTreeRef uintptr







type CGFloat = float64







type HRESULT = int32







type LPVOID = string







type REFIID = CFUUIDBytes







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





// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

// Package corefoundation provides Go bindings for the CoreFoundation framework.
//
// Access low-level functions, primitive data types, and various collection types that are bridged seamlessly with the Foundation framework.
//
// Core Foundation is a framework that provides fundamental software services useful to application services, application environments, and to applications themselves. Core Foundation also provides abstractions for common data types, facilitates internationalization with Unicode string storage, and offers a suite of utilities such as plug-in support, XML property lists, URL resource access, and preferences.
//
// # Utilities
//
//   - Base Utilities ([CFComparatorFunction], [CFIndex], [CFOptionFlags], [CFRange], [CFComparisonResult])
//   - Byte-Order Utilities ([CFSwappedFloat32], [CFSwappedFloat64], [CFByteOrder])
//   - Core Foundation URL Access Utilities ([CFURLError])
//   - Preferences Utilities
//   - Socket Name Server Utilities
//   - Time Utilities ([CFAbsoluteTime], [CFGregorianDate], [CFGregorianUnits], [CFTimeInterval], [CFGregorianUnitFlags])
//
// # Opaque Types
//
//   - CFAllocator ([CFAllocatorAllocateCallBack], [CFAllocatorCopyDescriptionCallBack], [CFAllocatorDeallocateCallBack], [CFAllocatorPreferredSizeCallBack], [CFAllocatorReallocateCallBack])
//   - CFArray ([CFArrayApplierFunction], [CFArrayCopyDescriptionCallBack], [CFArrayEqualCallBack], [CFArrayReleaseCallBack], [CFArrayRetainCallBack])
//   - CFAttributedString
//   - CFBag ([CFBagApplierFunction], [CFBagCopyDescriptionCallBack], [CFBagEqualCallBack], [CFBagHashCallBack], [CFBagReleaseCallBack])
//   - CFBinaryHeap ([CFBinaryHeapApplierFunction], [CFBinaryHeapCallBacks], [CFBinaryHeapCompareContext])
//   - CFBitVector ([CFBit])
//   - CFBoolean
//   - CFBundle ([CFBundleRefNum])
//   - CFCalendar ([CFCalendarUnit])
//   - CFCharacterSet ([CFCharacterSetPredefinedSet])
//   - CFData ([CFDataSearchFlags])
//   - CFDate
//   - CFDateFormatter ([CFDateFormatterStyle])
//   - CFDictionary ([CFDictionaryApplierFunction], [CFDictionaryCopyDescriptionCallBack], [CFDictionaryEqualCallBack], [CFDictionaryHashCallBack], [CFDictionaryReleaseCallBack])
//   - CFError
//   - CFFileDescriptor ([CFFileDescriptorNativeDescriptor], [CFFileDescriptorCallBack], [CFFileDescriptorContext])
//   - CFFileSecurity: Encapsulates a file system object’s security information in a Core Foundation object.
//   - CFLocale ([CFLocaleLanguageDirection])
//   - CFMachPort ([CFMachPortCallBack], [CFMachPortInvalidationCallBack], [CFMachPortContext])
//   - CFMessagePort ([CFMessagePortCallBack], [CFMessagePortInvalidationCallBack], [CFMessagePortContext])
//   - CFMutableArray
//   - CFMutableAttributedString
//   - CFMutableBag
//   - CFMutableBitVector
//   - CFMutableCharacterSet
//   - CFMutableData
//   - CFMutableDictionary
//   - CFMutableSet
//   - CFMutableString ([CFStringNormalizationForm])
//   - CFNotificationCenter ([CFNotificationCallback], [CFNotificationSuspensionBehavior])
//   - CFNull
//   - CFNumber ([CFNumberType])
//   - CFNumberFormatter ([CFNumberFormatterStyle], [CFNumberFormatterOptionFlags], [CFNumberFormatterPadPosition], [CFNumberFormatterRoundingMode])
//   - CFPlugIn ([CFPlugInDynamicRegisterFunction], [CFPlugInFactoryFunction], [CFPlugInUnloadFunction])
//   - CFPlugInInstance ([CFPlugInInstanceDeallocateInstanceDataFunction], [CFPlugInInstanceGetInterfaceFunction])
//   - CFPropertyList ([CFPropertyListMutabilityOptions], [CFPropertyListFormat])
//   - CFReadStream ([CFReadStreamClientCallBack], [CFStreamClientContext])
//   - CFRunLoop
//   - CFRunLoopObserver ([CFRunLoopObserverCallBack], [CFRunLoopObserverContext], [CFRunLoopActivity])
//   - CFRunLoopSource ([CFRunLoopSourceContext], [CFRunLoopSourceContext1])
//   - CFRunLoopTimer ([CFRunLoopTimerCallBack], [CFRunLoopTimerContext])
//   - CFSet ([CFSetApplierFunction], [CFSetCopyDescriptionCallBack], [CFSetEqualCallBack], [CFSetHashCallBack], [CFSetReleaseCallBack])
//   - CFSocket ([CFSocketCallBack], [CFSocketContext], [CFSocketNativeHandle], [CFSocketSignature], [CFSocketCallBackType])
//   - CFString ([CFStringEncoding], [CFStringEncodings], [CFStringCompareFlags], [CFStringInlineBuffer], [CFStringBuiltInEncodings])
//   - CFStringTokenizer ([CFStringTokenizerTokenType])
//   - CFTimeZone ([CFTimeZoneNameStyle])
//   - CFTree ([CFTreeApplierFunction], [CFTreeCopyDescriptionCallBack], [CFTreeReleaseCallBack], [CFTreeRetainCallBack], [CFTreeContext])
//   - CFURL ([CFURLBookmarkCreationOptions], [CFURLBookmarkFileCreationOptions], [CFURLBookmarkResolutionOptions], [CFURLComponentType], [CFURLPathStyle])
//   - CFUserNotification ([CFUserNotificationCallBack])
//   - CFURLEnumerator: A reference to a  object.
//   - CFUUID ([CFUUIDBytes])
//   - CFWriteStream ([CFWriteStreamClientCallBack])
//   - CFXMLNode ([CFXMLAttributeDeclarationInfo], [CFXMLAttributeListDeclarationInfo], [CFXMLDocumentInfo], [CFXMLDocumentTypeInfo], [CFXMLElementInfo])
//   - CFXMLParser ([CFXMLParserAddChildCallBack], [CFXMLParserCopyDescriptionCallBack], [CFXMLParserCreateXMLStructureCallBack], [CFXMLParserEndXMLStructureCallBack], [CFXMLParserHandleErrorCallBack])
//   - CFXMLTree
//
// # Variables
//
//   - kCFURLUbiquitousItemIsSyncPausedKey
//   - kCFURLUbiquitousItemSupportedSyncControlsKey
//
// # Functions
//
//   - CFAttributedStringGetStatisticalWritingDirections(_:_:_:_:_:)
//
// [CoreFoundation Documentation]: https://developer.apple.com/documentation/CoreFoundation
package corefoundation

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreFoundation: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}


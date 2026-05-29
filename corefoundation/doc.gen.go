// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

// Package corefoundation provides Go bindings for the CoreFoundation framework.
//
// Access low-level functions, primitive data types, and various collection
// types that are bridged seamlessly with the Foundation framework.
//
// Core Foundation is a framework that provides fundamental software services
// useful to application services, application environments, and to
// applications themselves. Core Foundation also provides abstractions for
// common data types, facilitates internationalization with Unicode string
// storage, and offers a suite of utilities such as plug-in support, XML
// property lists, URL resource access, and preferences.
//
// # Utilities
//
//   - [Base Utilities] ([CFComparatorFunction], [CFIndex], [CFOptionFlags], [CFRange], [CFComparisonResult])
//   - [Byte-Order Utilities] ([CFSwappedFloat32], [CFSwappedFloat64], [CFByteOrder])
//   - [Core Foundation URL Access Utilities] ([CFURLError])
//   - [Preferences Utilities]
//   - [Socket Name Server Utilities]
//   - [Time Utilities] ([CFAbsoluteTime], [CFGregorianDate], [CFGregorianUnits], [CFTimeInterval], [CFGregorianUnitFlags])
//
// # Opaque Types
//
//   - [CFAllocatorRef] ([CFAllocatorAllocateCallBack], [CFAllocatorCopyDescriptionCallBack], [CFAllocatorDeallocateCallBack], [CFAllocatorPreferredSizeCallBack], [CFAllocatorReallocateCallBack])
//   - [CFArrayRef] ([CFArrayApplierFunction], [CFArrayCopyDescriptionCallBack], [CFArrayEqualCallBack], [CFArrayReleaseCallBack], [CFArrayRetainCallBack])
//   - [CFAttributedStringRef]
//   - [CFBagRef] ([CFBagApplierFunction], [CFBagCopyDescriptionCallBack], [CFBagEqualCallBack], [CFBagHashCallBack], [CFBagReleaseCallBack])
//   - [CFBinaryHeapRef] ([CFBinaryHeapApplierFunction], [CFBinaryHeapCallBacks], [CFBinaryHeapCompareContext])
//   - [CFBitVectorRef] ([CFBit])
//   - [CFBooleanRef]
//   - [CFBundleRef] ([CFBundleRefNum])
//   - [CFCalendarRef] ([CFCalendarUnit])
//   - [CFCharacterSetRef] ([CFCharacterSetPredefinedSet])
//   - [CFDataRef] ([CFDataSearchFlags])
//   - [CFDateRef]
//   - [CFDateFormatterRef] ([CFDateFormatterStyle])
//   - [CFDictionaryRef] ([CFDictionaryApplierFunction], [CFDictionaryCopyDescriptionCallBack], [CFDictionaryEqualCallBack], [CFDictionaryHashCallBack], [CFDictionaryReleaseCallBack])
//   - [CFErrorRef]
//   - [CFFileDescriptorRef] ([CFFileDescriptorNativeDescriptor], [CFFileDescriptorCallBack], [CFFileDescriptorContext])
//   - [CFFileSecurityRef]: Encapsulates a file system object’s security information in a Core Foundation object.
//   - [CFLocaleRef] ([CFLocaleLanguageDirection])
//   - [CFMachPortRef] ([CFMachPortCallBack], [CFMachPortInvalidationCallBack], [CFMachPortContext])
//   - [CFMessagePortRef] ([CFMessagePortCallBack], [CFMessagePortInvalidationCallBack], [CFMessagePortContext])
//   - [CFMutableArrayRef]
//   - [CFMutableAttributedStringRef]
//   - [CFMutableBagRef]
//   - [CFMutableBitVectorRef]
//   - [CFMutableCharacterSetRef]
//   - [CFMutableDataRef]
//   - [CFMutableDictionaryRef]
//   - [CFMutableSetRef]
//   - [CFMutableStringRef] ([CFStringNormalizationForm])
//   - [CFNotificationCenterRef] ([CFNotificationCallback], [CFNotificationSuspensionBehavior])
//   - [CFNullRef]
//   - [CFNumberRef] ([CFNumberType])
//   - [CFNumberFormatterRef] ([CFNumberFormatterStyle], [CFNumberFormatterOptionFlags], [CFNumberFormatterPadPosition], [CFNumberFormatterRoundingMode])
//   - [CFPlugInRef] ([CFPlugInDynamicRegisterFunction], [CFPlugInFactoryFunction], [CFPlugInUnloadFunction])
//   - [CFPlugInInstanceRef] ([CFPlugInInstanceDeallocateInstanceDataFunction], [CFPlugInInstanceGetInterfaceFunction])
//   - [CFPropertyListRef] ([CFPropertyListMutabilityOptions], [CFPropertyListFormat])
//   - [CFReadStreamRef] ([CFReadStreamClientCallBack], [CFStreamClientContext])
//   - [CFRunLoopRef]
//   - [CFRunLoopObserverRef] ([CFRunLoopObserverCallBack], [CFRunLoopObserverContext], [CFRunLoopActivity])
//   - [CFRunLoopSourceRef] ([CFRunLoopSourceContext], [CFRunLoopSourceContext1])
//   - [CFRunLoopTimerRef] ([CFRunLoopTimerCallBack], [CFRunLoopTimerContext])
//   - [CFSetRef] ([CFSetApplierFunction], [CFSetCopyDescriptionCallBack], [CFSetEqualCallBack], [CFSetHashCallBack], [CFSetReleaseCallBack])
//   - [CFSocketRef] ([CFSocketCallBack], [CFSocketContext], [CFSocketNativeHandle], [CFSocketSignature], [CFSocketCallBackType])
//   - [CFStringRef] ([CFStringEncoding], [CFStringEncodings], [CFStringCompareFlags], [CFStringInlineBuffer], [CFStringBuiltInEncodings])
//   - [CFStringTokenizerRef] ([CFStringTokenizerTokenType])
//   - [CFTimeZoneRef] ([CFTimeZoneNameStyle])
//   - [CFTreeRef] ([CFTreeApplierFunction], [CFTreeCopyDescriptionCallBack], [CFTreeReleaseCallBack], [CFTreeRetainCallBack], [CFTreeContext])
//   - [CFURLRef] ([CFURLBookmarkCreationOptions], [CFURLBookmarkFileCreationOptions], [CFURLBookmarkResolutionOptions], [CFURLComponentType], [CFURLPathStyle])
//   - [CFUserNotificationRef] ([CFUserNotificationCallBack])
//   - [CFURLEnumeratorRef]: A reference to a object.
//   - [CFUUIDRef] ([CFUUIDBytes])
//   - [CFWriteStreamRef] ([CFWriteStreamClientCallBack])
//   - CFXMLNodeRef ([CFXMLAttributeDeclarationInfo], [CFXMLAttributeListDeclarationInfo], [CFXMLDocumentInfo], [CFXMLDocumentTypeInfo], [CFXMLElementInfo])
//   - [CFXMLParserRef] ([CFXMLParserAddChildCallBack], [CFXMLParserCopyDescriptionCallBack], [CFXMLParserCreateXMLStructureCallBack], [CFXMLParserEndXMLStructureCallBack], [CFXMLParserHandleErrorCallBack])
//   - [CFXMLTreeRef]
//
// # Variables
//
//   - [KCFBanglaCalendar]
//   - [KCFDangiCalendar]
//   - [KCFGujaratiCalendar]
//   - [KCFKannadaCalendar]
//   - [KCFMalayalamCalendar]
//   - [KCFMarathiCalendar]
//   - [KCFOdiaCalendar]
//   - [KCFTamilCalendar]
//   - [KCFTeluguCalendar]
//   - [KCFVietnameseCalendar]
//   - [KCFVikramCalendar]
//   - [KCFURLUbiquitousItemIsSyncPausedKey]
//   - [KCFURLUbiquitousItemSupportedSyncControlsKey]
//   - kCFUserNotificationAlertAccessibilityIdentifierKey
//   - kCFUserNotificationAlternateButtonAccessibilityIdentifierKey
//   - kCFUserNotificationDefaultButtonAccessibilityIdentifierKey
//   - kCFUserNotificationOtherButtonAccessibilityIdentifierKey
//
// # Functions
//
//   - [CFAllocatorCreateWithZone]
//   - [CFAttributedStringGetStatisticalWritingDirections]
//
// # Macros
//
//   - CF_HEADER_AUDIT_BEGIN
//   - CF_HEADER_AUDIT_END
//   - CF_SWIFT_MAIN_ACTOR
//   - CF_SWIFT_NONISOLATED
//   - CF_SWIFT_NONSENDABLE
//   - CF_SWIFT_SENDABLE//
//
// [Base Utilities]: https://developer.apple.com/documentation/corefoundation/base-utilities
// [Byte-Order Utilities]: https://developer.apple.com/documentation/corefoundation/byte-order-utilities
// [Core Foundation URL Access Utilities]: https://developer.apple.com/documentation/corefoundation/core-foundation-url-access-utilities
// [Preferences Utilities]: https://developer.apple.com/documentation/corefoundation/preferences-utilities
// [Socket Name Server Utilities]: https://developer.apple.com/documentation/corefoundation/socket-name-server-utilities
// [Time Utilities]: https://developer.apple.com/documentation/corefoundation/time-utilities
package corefoundation

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreFoundation library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation",
	"/usr/lib/libCoreFoundation.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: CoreFoundation: failed to load framework from any known path\n")
	}
}

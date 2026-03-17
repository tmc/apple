// Code generated from Apple documentation for hiservices. DO NOT EDIT.

package hiservices

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: hiservices: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: hiservices: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: hiservices: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _aXAPIEnabledSymbol uintptr

// AXAPIEnabled has an opaque C signature in discovered metadata.
// Call AXAPIEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXAPIEnabled
func AXAPIEnabled() {
	panic("hiservices: symbol AXAPIEnabled has opaque signature; use AXAPIEnabledSymbol() and a typed manual wrapper")
}

// AXAPIEnabledSymbol returns the raw symbol address for AXAPIEnabled.
func AXAPIEnabledSymbol() uintptr {
	if _aXAPIEnabledSymbol == 0 {
		return 0
	}
	return _aXAPIEnabledSymbol
}

var _aXIsProcessTrustedSymbol uintptr

// AXIsProcessTrusted has an opaque C signature in discovered metadata.
// Call AXIsProcessTrustedSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXIsProcessTrusted
func AXIsProcessTrusted() {
	panic("hiservices: symbol AXIsProcessTrusted has opaque signature; use AXIsProcessTrustedSymbol() and a typed manual wrapper")
}

// AXIsProcessTrustedSymbol returns the raw symbol address for AXIsProcessTrusted.
func AXIsProcessTrustedSymbol() uintptr {
	if _aXIsProcessTrustedSymbol == 0 {
		return 0
	}
	return _aXIsProcessTrustedSymbol
}

var _aXIsProcessTrustedWithOptionsSymbol uintptr

// AXIsProcessTrustedWithOptions has an opaque C signature in discovered metadata.
// Call AXIsProcessTrustedWithOptionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXIsProcessTrustedWithOptions
func AXIsProcessTrustedWithOptions() {
	panic("hiservices: symbol AXIsProcessTrustedWithOptions has opaque signature; use AXIsProcessTrustedWithOptionsSymbol() and a typed manual wrapper")
}

// AXIsProcessTrustedWithOptionsSymbol returns the raw symbol address for AXIsProcessTrustedWithOptions.
func AXIsProcessTrustedWithOptionsSymbol() uintptr {
	if _aXIsProcessTrustedWithOptionsSymbol == 0 {
		return 0
	}
	return _aXIsProcessTrustedWithOptionsSymbol
}

var _aXMakeProcessTrustedSymbol uintptr

// AXMakeProcessTrusted has an opaque C signature in discovered metadata.
// Call AXMakeProcessTrustedSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXMakeProcessTrusted
func AXMakeProcessTrusted() {
	panic("hiservices: symbol AXMakeProcessTrusted has opaque signature; use AXMakeProcessTrustedSymbol() and a typed manual wrapper")
}

// AXMakeProcessTrustedSymbol returns the raw symbol address for AXMakeProcessTrusted.
func AXMakeProcessTrustedSymbol() uintptr {
	if _aXMakeProcessTrustedSymbol == 0 {
		return 0
	}
	return _aXMakeProcessTrustedSymbol
}

var _aXObserverAddNotificationSymbol uintptr

// AXObserverAddNotification has an opaque C signature in discovered metadata.
// Call AXObserverAddNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXObserverAddNotification
func AXObserverAddNotification() {
	panic("hiservices: symbol AXObserverAddNotification has opaque signature; use AXObserverAddNotificationSymbol() and a typed manual wrapper")
}

// AXObserverAddNotificationSymbol returns the raw symbol address for AXObserverAddNotification.
func AXObserverAddNotificationSymbol() uintptr {
	if _aXObserverAddNotificationSymbol == 0 {
		return 0
	}
	return _aXObserverAddNotificationSymbol
}

var _aXObserverAddNotificationAsyncSymbol uintptr

// AXObserverAddNotificationAsync has an opaque C signature in discovered metadata.
// Call AXObserverAddNotificationAsyncSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXObserverAddNotificationAsync
func AXObserverAddNotificationAsync() {
	panic("hiservices: symbol AXObserverAddNotificationAsync has opaque signature; use AXObserverAddNotificationAsyncSymbol() and a typed manual wrapper")
}

// AXObserverAddNotificationAsyncSymbol returns the raw symbol address for AXObserverAddNotificationAsync.
func AXObserverAddNotificationAsyncSymbol() uintptr {
	if _aXObserverAddNotificationAsyncSymbol == 0 {
		return 0
	}
	return _aXObserverAddNotificationAsyncSymbol
}

var _aXObserverCreateSymbol uintptr

// AXObserverCreate has an opaque C signature in discovered metadata.
// Call AXObserverCreateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXObserverCreate
func AXObserverCreate() {
	panic("hiservices: symbol AXObserverCreate has opaque signature; use AXObserverCreateSymbol() and a typed manual wrapper")
}

// AXObserverCreateSymbol returns the raw symbol address for AXObserverCreate.
func AXObserverCreateSymbol() uintptr {
	if _aXObserverCreateSymbol == 0 {
		return 0
	}
	return _aXObserverCreateSymbol
}

var _aXObserverCreateWithInfoCallbackSymbol uintptr

// AXObserverCreateWithInfoCallback has an opaque C signature in discovered metadata.
// Call AXObserverCreateWithInfoCallbackSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXObserverCreateWithInfoCallback
func AXObserverCreateWithInfoCallback() {
	panic("hiservices: symbol AXObserverCreateWithInfoCallback has opaque signature; use AXObserverCreateWithInfoCallbackSymbol() and a typed manual wrapper")
}

// AXObserverCreateWithInfoCallbackSymbol returns the raw symbol address for AXObserverCreateWithInfoCallback.
func AXObserverCreateWithInfoCallbackSymbol() uintptr {
	if _aXObserverCreateWithInfoCallbackSymbol == 0 {
		return 0
	}
	return _aXObserverCreateWithInfoCallbackSymbol
}

var _aXObserverGetRunLoopSourceSymbol uintptr

// AXObserverGetRunLoopSource has an opaque C signature in discovered metadata.
// Call AXObserverGetRunLoopSourceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXObserverGetRunLoopSource
func AXObserverGetRunLoopSource() {
	panic("hiservices: symbol AXObserverGetRunLoopSource has opaque signature; use AXObserverGetRunLoopSourceSymbol() and a typed manual wrapper")
}

// AXObserverGetRunLoopSourceSymbol returns the raw symbol address for AXObserverGetRunLoopSource.
func AXObserverGetRunLoopSourceSymbol() uintptr {
	if _aXObserverGetRunLoopSourceSymbol == 0 {
		return 0
	}
	return _aXObserverGetRunLoopSourceSymbol
}

var _aXObserverGetTypeIDSymbol uintptr

// AXObserverGetTypeID has an opaque C signature in discovered metadata.
// Call AXObserverGetTypeIDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXObserverGetTypeID
func AXObserverGetTypeID() {
	panic("hiservices: symbol AXObserverGetTypeID has opaque signature; use AXObserverGetTypeIDSymbol() and a typed manual wrapper")
}

// AXObserverGetTypeIDSymbol returns the raw symbol address for AXObserverGetTypeID.
func AXObserverGetTypeIDSymbol() uintptr {
	if _aXObserverGetTypeIDSymbol == 0 {
		return 0
	}
	return _aXObserverGetTypeIDSymbol
}

var _aXObserverRemoveNotificationSymbol uintptr

// AXObserverRemoveNotification has an opaque C signature in discovered metadata.
// Call AXObserverRemoveNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXObserverRemoveNotification
func AXObserverRemoveNotification() {
	panic("hiservices: symbol AXObserverRemoveNotification has opaque signature; use AXObserverRemoveNotificationSymbol() and a typed manual wrapper")
}

// AXObserverRemoveNotificationSymbol returns the raw symbol address for AXObserverRemoveNotification.
func AXObserverRemoveNotificationSymbol() uintptr {
	if _aXObserverRemoveNotificationSymbol == 0 {
		return 0
	}
	return _aXObserverRemoveNotificationSymbol
}

var _aXObserverRemoveNotificationAsyncSymbol uintptr

// AXObserverRemoveNotificationAsync has an opaque C signature in discovered metadata.
// Call AXObserverRemoveNotificationAsyncSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXObserverRemoveNotificationAsync
func AXObserverRemoveNotificationAsync() {
	panic("hiservices: symbol AXObserverRemoveNotificationAsync has opaque signature; use AXObserverRemoveNotificationAsyncSymbol() and a typed manual wrapper")
}

// AXObserverRemoveNotificationAsyncSymbol returns the raw symbol address for AXObserverRemoveNotificationAsync.
func AXObserverRemoveNotificationAsyncSymbol() uintptr {
	if _aXObserverRemoveNotificationAsyncSymbol == 0 {
		return 0
	}
	return _aXObserverRemoveNotificationAsyncSymbol
}

var _aXSerializeCFTypeSymbol uintptr

// AXSerializeCFType has an opaque C signature in discovered metadata.
// Call AXSerializeCFTypeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXSerializeCFType
func AXSerializeCFType() {
	panic("hiservices: symbol AXSerializeCFType has opaque signature; use AXSerializeCFTypeSymbol() and a typed manual wrapper")
}

// AXSerializeCFTypeSymbol returns the raw symbol address for AXSerializeCFType.
func AXSerializeCFTypeSymbol() uintptr {
	if _aXSerializeCFTypeSymbol == 0 {
		return 0
	}
	return _aXSerializeCFTypeSymbol
}

var _aXTextMarkerCreateSymbol uintptr

// AXTextMarkerCreate has an opaque C signature in discovered metadata.
// Call AXTextMarkerCreateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXTextMarkerCreate
func AXTextMarkerCreate() {
	panic("hiservices: symbol AXTextMarkerCreate has opaque signature; use AXTextMarkerCreateSymbol() and a typed manual wrapper")
}

// AXTextMarkerCreateSymbol returns the raw symbol address for AXTextMarkerCreate.
func AXTextMarkerCreateSymbol() uintptr {
	if _aXTextMarkerCreateSymbol == 0 {
		return 0
	}
	return _aXTextMarkerCreateSymbol
}

var _aXTextMarkerGetBytePtrSymbol uintptr

// AXTextMarkerGetBytePtr has an opaque C signature in discovered metadata.
// Call AXTextMarkerGetBytePtrSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXTextMarkerGetBytePtr
func AXTextMarkerGetBytePtr() {
	panic("hiservices: symbol AXTextMarkerGetBytePtr has opaque signature; use AXTextMarkerGetBytePtrSymbol() and a typed manual wrapper")
}

// AXTextMarkerGetBytePtrSymbol returns the raw symbol address for AXTextMarkerGetBytePtr.
func AXTextMarkerGetBytePtrSymbol() uintptr {
	if _aXTextMarkerGetBytePtrSymbol == 0 {
		return 0
	}
	return _aXTextMarkerGetBytePtrSymbol
}

var _aXTextMarkerGetLengthSymbol uintptr

// AXTextMarkerGetLength has an opaque C signature in discovered metadata.
// Call AXTextMarkerGetLengthSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXTextMarkerGetLength
func AXTextMarkerGetLength() {
	panic("hiservices: symbol AXTextMarkerGetLength has opaque signature; use AXTextMarkerGetLengthSymbol() and a typed manual wrapper")
}

// AXTextMarkerGetLengthSymbol returns the raw symbol address for AXTextMarkerGetLength.
func AXTextMarkerGetLengthSymbol() uintptr {
	if _aXTextMarkerGetLengthSymbol == 0 {
		return 0
	}
	return _aXTextMarkerGetLengthSymbol
}

var _aXTextMarkerGetTypeIDSymbol uintptr

// AXTextMarkerGetTypeID has an opaque C signature in discovered metadata.
// Call AXTextMarkerGetTypeIDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXTextMarkerGetTypeID
func AXTextMarkerGetTypeID() {
	panic("hiservices: symbol AXTextMarkerGetTypeID has opaque signature; use AXTextMarkerGetTypeIDSymbol() and a typed manual wrapper")
}

// AXTextMarkerGetTypeIDSymbol returns the raw symbol address for AXTextMarkerGetTypeID.
func AXTextMarkerGetTypeIDSymbol() uintptr {
	if _aXTextMarkerGetTypeIDSymbol == 0 {
		return 0
	}
	return _aXTextMarkerGetTypeIDSymbol
}

var _aXTextMarkerRangeCopyEndMarkerSymbol uintptr

// AXTextMarkerRangeCopyEndMarker has an opaque C signature in discovered metadata.
// Call AXTextMarkerRangeCopyEndMarkerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXTextMarkerRangeCopyEndMarker
func AXTextMarkerRangeCopyEndMarker() {
	panic("hiservices: symbol AXTextMarkerRangeCopyEndMarker has opaque signature; use AXTextMarkerRangeCopyEndMarkerSymbol() and a typed manual wrapper")
}

// AXTextMarkerRangeCopyEndMarkerSymbol returns the raw symbol address for AXTextMarkerRangeCopyEndMarker.
func AXTextMarkerRangeCopyEndMarkerSymbol() uintptr {
	if _aXTextMarkerRangeCopyEndMarkerSymbol == 0 {
		return 0
	}
	return _aXTextMarkerRangeCopyEndMarkerSymbol
}

var _aXTextMarkerRangeCopyStartMarkerSymbol uintptr

// AXTextMarkerRangeCopyStartMarker has an opaque C signature in discovered metadata.
// Call AXTextMarkerRangeCopyStartMarkerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXTextMarkerRangeCopyStartMarker
func AXTextMarkerRangeCopyStartMarker() {
	panic("hiservices: symbol AXTextMarkerRangeCopyStartMarker has opaque signature; use AXTextMarkerRangeCopyStartMarkerSymbol() and a typed manual wrapper")
}

// AXTextMarkerRangeCopyStartMarkerSymbol returns the raw symbol address for AXTextMarkerRangeCopyStartMarker.
func AXTextMarkerRangeCopyStartMarkerSymbol() uintptr {
	if _aXTextMarkerRangeCopyStartMarkerSymbol == 0 {
		return 0
	}
	return _aXTextMarkerRangeCopyStartMarkerSymbol
}

var _aXTextMarkerRangeCreateSymbol uintptr

// AXTextMarkerRangeCreate has an opaque C signature in discovered metadata.
// Call AXTextMarkerRangeCreateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXTextMarkerRangeCreate
func AXTextMarkerRangeCreate() {
	panic("hiservices: symbol AXTextMarkerRangeCreate has opaque signature; use AXTextMarkerRangeCreateSymbol() and a typed manual wrapper")
}

// AXTextMarkerRangeCreateSymbol returns the raw symbol address for AXTextMarkerRangeCreate.
func AXTextMarkerRangeCreateSymbol() uintptr {
	if _aXTextMarkerRangeCreateSymbol == 0 {
		return 0
	}
	return _aXTextMarkerRangeCreateSymbol
}

var _aXTextMarkerRangeCreateWithBytesSymbol uintptr

// AXTextMarkerRangeCreateWithBytes has an opaque C signature in discovered metadata.
// Call AXTextMarkerRangeCreateWithBytesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXTextMarkerRangeCreateWithBytes
func AXTextMarkerRangeCreateWithBytes() {
	panic("hiservices: symbol AXTextMarkerRangeCreateWithBytes has opaque signature; use AXTextMarkerRangeCreateWithBytesSymbol() and a typed manual wrapper")
}

// AXTextMarkerRangeCreateWithBytesSymbol returns the raw symbol address for AXTextMarkerRangeCreateWithBytes.
func AXTextMarkerRangeCreateWithBytesSymbol() uintptr {
	if _aXTextMarkerRangeCreateWithBytesSymbol == 0 {
		return 0
	}
	return _aXTextMarkerRangeCreateWithBytesSymbol
}

var _aXTextMarkerRangeGetTypeIDSymbol uintptr

// AXTextMarkerRangeGetTypeID has an opaque C signature in discovered metadata.
// Call AXTextMarkerRangeGetTypeIDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXTextMarkerRangeGetTypeID
func AXTextMarkerRangeGetTypeID() {
	panic("hiservices: symbol AXTextMarkerRangeGetTypeID has opaque signature; use AXTextMarkerRangeGetTypeIDSymbol() and a typed manual wrapper")
}

// AXTextMarkerRangeGetTypeIDSymbol returns the raw symbol address for AXTextMarkerRangeGetTypeID.
func AXTextMarkerRangeGetTypeIDSymbol() uintptr {
	if _aXTextMarkerRangeGetTypeIDSymbol == 0 {
		return 0
	}
	return _aXTextMarkerRangeGetTypeIDSymbol
}

var _aXUIElementCopyActionDescriptionSymbol uintptr

// AXUIElementCopyActionDescription has an opaque C signature in discovered metadata.
// Call AXUIElementCopyActionDescriptionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCopyActionDescription
func AXUIElementCopyActionDescription() {
	panic("hiservices: symbol AXUIElementCopyActionDescription has opaque signature; use AXUIElementCopyActionDescriptionSymbol() and a typed manual wrapper")
}

// AXUIElementCopyActionDescriptionSymbol returns the raw symbol address for AXUIElementCopyActionDescription.
func AXUIElementCopyActionDescriptionSymbol() uintptr {
	if _aXUIElementCopyActionDescriptionSymbol == 0 {
		return 0
	}
	return _aXUIElementCopyActionDescriptionSymbol
}

var _aXUIElementCopyActionNamesSymbol uintptr

// AXUIElementCopyActionNames has an opaque C signature in discovered metadata.
// Call AXUIElementCopyActionNamesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCopyActionNames
func AXUIElementCopyActionNames() {
	panic("hiservices: symbol AXUIElementCopyActionNames has opaque signature; use AXUIElementCopyActionNamesSymbol() and a typed manual wrapper")
}

// AXUIElementCopyActionNamesSymbol returns the raw symbol address for AXUIElementCopyActionNames.
func AXUIElementCopyActionNamesSymbol() uintptr {
	if _aXUIElementCopyActionNamesSymbol == 0 {
		return 0
	}
	return _aXUIElementCopyActionNamesSymbol
}

var _aXUIElementCopyAttributeNamesSymbol uintptr

// AXUIElementCopyAttributeNames has an opaque C signature in discovered metadata.
// Call AXUIElementCopyAttributeNamesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCopyAttributeNames
func AXUIElementCopyAttributeNames() {
	panic("hiservices: symbol AXUIElementCopyAttributeNames has opaque signature; use AXUIElementCopyAttributeNamesSymbol() and a typed manual wrapper")
}

// AXUIElementCopyAttributeNamesSymbol returns the raw symbol address for AXUIElementCopyAttributeNames.
func AXUIElementCopyAttributeNamesSymbol() uintptr {
	if _aXUIElementCopyAttributeNamesSymbol == 0 {
		return 0
	}
	return _aXUIElementCopyAttributeNamesSymbol
}

var _aXUIElementCopyAttributeValueSymbol uintptr

// AXUIElementCopyAttributeValue has an opaque C signature in discovered metadata.
// Call AXUIElementCopyAttributeValueSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCopyAttributeValue
func AXUIElementCopyAttributeValue() {
	panic("hiservices: symbol AXUIElementCopyAttributeValue has opaque signature; use AXUIElementCopyAttributeValueSymbol() and a typed manual wrapper")
}

// AXUIElementCopyAttributeValueSymbol returns the raw symbol address for AXUIElementCopyAttributeValue.
func AXUIElementCopyAttributeValueSymbol() uintptr {
	if _aXUIElementCopyAttributeValueSymbol == 0 {
		return 0
	}
	return _aXUIElementCopyAttributeValueSymbol
}

var _aXUIElementCopyAttributeValuesSymbol uintptr

// AXUIElementCopyAttributeValues has an opaque C signature in discovered metadata.
// Call AXUIElementCopyAttributeValuesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCopyAttributeValues
func AXUIElementCopyAttributeValues() {
	panic("hiservices: symbol AXUIElementCopyAttributeValues has opaque signature; use AXUIElementCopyAttributeValuesSymbol() and a typed manual wrapper")
}

// AXUIElementCopyAttributeValuesSymbol returns the raw symbol address for AXUIElementCopyAttributeValues.
func AXUIElementCopyAttributeValuesSymbol() uintptr {
	if _aXUIElementCopyAttributeValuesSymbol == 0 {
		return 0
	}
	return _aXUIElementCopyAttributeValuesSymbol
}

var _aXUIElementCopyElementAtPositionSymbol uintptr

// AXUIElementCopyElementAtPosition has an opaque C signature in discovered metadata.
// Call AXUIElementCopyElementAtPositionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCopyElementAtPosition
func AXUIElementCopyElementAtPosition() {
	panic("hiservices: symbol AXUIElementCopyElementAtPosition has opaque signature; use AXUIElementCopyElementAtPositionSymbol() and a typed manual wrapper")
}

// AXUIElementCopyElementAtPositionSymbol returns the raw symbol address for AXUIElementCopyElementAtPosition.
func AXUIElementCopyElementAtPositionSymbol() uintptr {
	if _aXUIElementCopyElementAtPositionSymbol == 0 {
		return 0
	}
	return _aXUIElementCopyElementAtPositionSymbol
}

var _aXUIElementCopyHierarchySymbol uintptr

// AXUIElementCopyHierarchy has an opaque C signature in discovered metadata.
// Call AXUIElementCopyHierarchySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCopyHierarchy
func AXUIElementCopyHierarchy() {
	panic("hiservices: symbol AXUIElementCopyHierarchy has opaque signature; use AXUIElementCopyHierarchySymbol() and a typed manual wrapper")
}

// AXUIElementCopyHierarchySymbol returns the raw symbol address for AXUIElementCopyHierarchy.
func AXUIElementCopyHierarchySymbol() uintptr {
	if _aXUIElementCopyHierarchySymbol == 0 {
		return 0
	}
	return _aXUIElementCopyHierarchySymbol
}

var _aXUIElementCopyMultipleAttributeValuesSymbol uintptr

// AXUIElementCopyMultipleAttributeValues has an opaque C signature in discovered metadata.
// Call AXUIElementCopyMultipleAttributeValuesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCopyMultipleAttributeValues
func AXUIElementCopyMultipleAttributeValues() {
	panic("hiservices: symbol AXUIElementCopyMultipleAttributeValues has opaque signature; use AXUIElementCopyMultipleAttributeValuesSymbol() and a typed manual wrapper")
}

// AXUIElementCopyMultipleAttributeValuesSymbol returns the raw symbol address for AXUIElementCopyMultipleAttributeValues.
func AXUIElementCopyMultipleAttributeValuesSymbol() uintptr {
	if _aXUIElementCopyMultipleAttributeValuesSymbol == 0 {
		return 0
	}
	return _aXUIElementCopyMultipleAttributeValuesSymbol
}

var _aXUIElementCopyParameterizedAttributeNamesSymbol uintptr

// AXUIElementCopyParameterizedAttributeNames has an opaque C signature in discovered metadata.
// Call AXUIElementCopyParameterizedAttributeNamesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCopyParameterizedAttributeNames
func AXUIElementCopyParameterizedAttributeNames() {
	panic("hiservices: symbol AXUIElementCopyParameterizedAttributeNames has opaque signature; use AXUIElementCopyParameterizedAttributeNamesSymbol() and a typed manual wrapper")
}

// AXUIElementCopyParameterizedAttributeNamesSymbol returns the raw symbol address for AXUIElementCopyParameterizedAttributeNames.
func AXUIElementCopyParameterizedAttributeNamesSymbol() uintptr {
	if _aXUIElementCopyParameterizedAttributeNamesSymbol == 0 {
		return 0
	}
	return _aXUIElementCopyParameterizedAttributeNamesSymbol
}

var _aXUIElementCopyParameterizedAttributeValueSymbol uintptr

// AXUIElementCopyParameterizedAttributeValue has an opaque C signature in discovered metadata.
// Call AXUIElementCopyParameterizedAttributeValueSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCopyParameterizedAttributeValue
func AXUIElementCopyParameterizedAttributeValue() {
	panic("hiservices: symbol AXUIElementCopyParameterizedAttributeValue has opaque signature; use AXUIElementCopyParameterizedAttributeValueSymbol() and a typed manual wrapper")
}

// AXUIElementCopyParameterizedAttributeValueSymbol returns the raw symbol address for AXUIElementCopyParameterizedAttributeValue.
func AXUIElementCopyParameterizedAttributeValueSymbol() uintptr {
	if _aXUIElementCopyParameterizedAttributeValueSymbol == 0 {
		return 0
	}
	return _aXUIElementCopyParameterizedAttributeValueSymbol
}

var _aXUIElementCreateApplicationSymbol uintptr

// AXUIElementCreateApplication has an opaque C signature in discovered metadata.
// Call AXUIElementCreateApplicationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCreateApplication
func AXUIElementCreateApplication() {
	panic("hiservices: symbol AXUIElementCreateApplication has opaque signature; use AXUIElementCreateApplicationSymbol() and a typed manual wrapper")
}

// AXUIElementCreateApplicationSymbol returns the raw symbol address for AXUIElementCreateApplication.
func AXUIElementCreateApplicationSymbol() uintptr {
	if _aXUIElementCreateApplicationSymbol == 0 {
		return 0
	}
	return _aXUIElementCreateApplicationSymbol
}

var _aXUIElementCreateSystemWideSymbol uintptr

// AXUIElementCreateSystemWide has an opaque C signature in discovered metadata.
// Call AXUIElementCreateSystemWideSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementCreateSystemWide
func AXUIElementCreateSystemWide() {
	panic("hiservices: symbol AXUIElementCreateSystemWide has opaque signature; use AXUIElementCreateSystemWideSymbol() and a typed manual wrapper")
}

// AXUIElementCreateSystemWideSymbol returns the raw symbol address for AXUIElementCreateSystemWide.
func AXUIElementCreateSystemWideSymbol() uintptr {
	if _aXUIElementCreateSystemWideSymbol == 0 {
		return 0
	}
	return _aXUIElementCreateSystemWideSymbol
}

var _aXUIElementGetAttributeValueCountSymbol uintptr

// AXUIElementGetAttributeValueCount has an opaque C signature in discovered metadata.
// Call AXUIElementGetAttributeValueCountSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementGetAttributeValueCount
func AXUIElementGetAttributeValueCount() {
	panic("hiservices: symbol AXUIElementGetAttributeValueCount has opaque signature; use AXUIElementGetAttributeValueCountSymbol() and a typed manual wrapper")
}

// AXUIElementGetAttributeValueCountSymbol returns the raw symbol address for AXUIElementGetAttributeValueCount.
func AXUIElementGetAttributeValueCountSymbol() uintptr {
	if _aXUIElementGetAttributeValueCountSymbol == 0 {
		return 0
	}
	return _aXUIElementGetAttributeValueCountSymbol
}

var _aXUIElementGetPidSymbol uintptr

// AXUIElementGetPid has an opaque C signature in discovered metadata.
// Call AXUIElementGetPidSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementGetPid
func AXUIElementGetPid() {
	panic("hiservices: symbol AXUIElementGetPid has opaque signature; use AXUIElementGetPidSymbol() and a typed manual wrapper")
}

// AXUIElementGetPidSymbol returns the raw symbol address for AXUIElementGetPid.
func AXUIElementGetPidSymbol() uintptr {
	if _aXUIElementGetPidSymbol == 0 {
		return 0
	}
	return _aXUIElementGetPidSymbol
}

var _aXUIElementGetTypeIDSymbol uintptr

// AXUIElementGetTypeID has an opaque C signature in discovered metadata.
// Call AXUIElementGetTypeIDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementGetTypeID
func AXUIElementGetTypeID() {
	panic("hiservices: symbol AXUIElementGetTypeID has opaque signature; use AXUIElementGetTypeIDSymbol() and a typed manual wrapper")
}

// AXUIElementGetTypeIDSymbol returns the raw symbol address for AXUIElementGetTypeID.
func AXUIElementGetTypeIDSymbol() uintptr {
	if _aXUIElementGetTypeIDSymbol == 0 {
		return 0
	}
	return _aXUIElementGetTypeIDSymbol
}

var _aXUIElementIsAttributeSettableSymbol uintptr

// AXUIElementIsAttributeSettable has an opaque C signature in discovered metadata.
// Call AXUIElementIsAttributeSettableSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementIsAttributeSettable
func AXUIElementIsAttributeSettable() {
	panic("hiservices: symbol AXUIElementIsAttributeSettable has opaque signature; use AXUIElementIsAttributeSettableSymbol() and a typed manual wrapper")
}

// AXUIElementIsAttributeSettableSymbol returns the raw symbol address for AXUIElementIsAttributeSettable.
func AXUIElementIsAttributeSettableSymbol() uintptr {
	if _aXUIElementIsAttributeSettableSymbol == 0 {
		return 0
	}
	return _aXUIElementIsAttributeSettableSymbol
}

var _aXUIElementPerformActionSymbol uintptr

// AXUIElementPerformAction has an opaque C signature in discovered metadata.
// Call AXUIElementPerformActionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementPerformAction
func AXUIElementPerformAction() {
	panic("hiservices: symbol AXUIElementPerformAction has opaque signature; use AXUIElementPerformActionSymbol() and a typed manual wrapper")
}

// AXUIElementPerformActionSymbol returns the raw symbol address for AXUIElementPerformAction.
func AXUIElementPerformActionSymbol() uintptr {
	if _aXUIElementPerformActionSymbol == 0 {
		return 0
	}
	return _aXUIElementPerformActionSymbol
}

var _aXUIElementPostKeyboardEventSymbol uintptr

// AXUIElementPostKeyboardEvent has an opaque C signature in discovered metadata.
// Call AXUIElementPostKeyboardEventSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementPostKeyboardEvent
func AXUIElementPostKeyboardEvent() {
	panic("hiservices: symbol AXUIElementPostKeyboardEvent has opaque signature; use AXUIElementPostKeyboardEventSymbol() and a typed manual wrapper")
}

// AXUIElementPostKeyboardEventSymbol returns the raw symbol address for AXUIElementPostKeyboardEvent.
func AXUIElementPostKeyboardEventSymbol() uintptr {
	if _aXUIElementPostKeyboardEventSymbol == 0 {
		return 0
	}
	return _aXUIElementPostKeyboardEventSymbol
}

var _aXUIElementSetAttributeValueSymbol uintptr

// AXUIElementSetAttributeValue has an opaque C signature in discovered metadata.
// Call AXUIElementSetAttributeValueSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementSetAttributeValue
func AXUIElementSetAttributeValue() {
	panic("hiservices: symbol AXUIElementSetAttributeValue has opaque signature; use AXUIElementSetAttributeValueSymbol() and a typed manual wrapper")
}

// AXUIElementSetAttributeValueSymbol returns the raw symbol address for AXUIElementSetAttributeValue.
func AXUIElementSetAttributeValueSymbol() uintptr {
	if _aXUIElementSetAttributeValueSymbol == 0 {
		return 0
	}
	return _aXUIElementSetAttributeValueSymbol
}

var _aXUIElementSetMessagingTimeoutSymbol uintptr

// AXUIElementSetMessagingTimeout has an opaque C signature in discovered metadata.
// Call AXUIElementSetMessagingTimeoutSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUIElementSetMessagingTimeout
func AXUIElementSetMessagingTimeout() {
	panic("hiservices: symbol AXUIElementSetMessagingTimeout has opaque signature; use AXUIElementSetMessagingTimeoutSymbol() and a typed manual wrapper")
}

// AXUIElementSetMessagingTimeoutSymbol returns the raw symbol address for AXUIElementSetMessagingTimeout.
func AXUIElementSetMessagingTimeoutSymbol() uintptr {
	if _aXUIElementSetMessagingTimeoutSymbol == 0 {
		return 0
	}
	return _aXUIElementSetMessagingTimeoutSymbol
}

var _aXUnserializeCFTypeSymbol uintptr

// AXUnserializeCFType has an opaque C signature in discovered metadata.
// Call AXUnserializeCFTypeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXUnserializeCFType
func AXUnserializeCFType() {
	panic("hiservices: symbol AXUnserializeCFType has opaque signature; use AXUnserializeCFTypeSymbol() and a typed manual wrapper")
}

// AXUnserializeCFTypeSymbol returns the raw symbol address for AXUnserializeCFType.
func AXUnserializeCFTypeSymbol() uintptr {
	if _aXUnserializeCFTypeSymbol == 0 {
		return 0
	}
	return _aXUnserializeCFTypeSymbol
}

var _aXValueCreateSymbol uintptr

// AXValueCreate has an opaque C signature in discovered metadata.
// Call AXValueCreateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXValueCreate
func AXValueCreate() {
	panic("hiservices: symbol AXValueCreate has opaque signature; use AXValueCreateSymbol() and a typed manual wrapper")
}

// AXValueCreateSymbol returns the raw symbol address for AXValueCreate.
func AXValueCreateSymbol() uintptr {
	if _aXValueCreateSymbol == 0 {
		return 0
	}
	return _aXValueCreateSymbol
}

var _aXValueGetTypeSymbol uintptr

// AXValueGetType has an opaque C signature in discovered metadata.
// Call AXValueGetTypeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXValueGetType
func AXValueGetType() {
	panic("hiservices: symbol AXValueGetType has opaque signature; use AXValueGetTypeSymbol() and a typed manual wrapper")
}

// AXValueGetTypeSymbol returns the raw symbol address for AXValueGetType.
func AXValueGetTypeSymbol() uintptr {
	if _aXValueGetTypeSymbol == 0 {
		return 0
	}
	return _aXValueGetTypeSymbol
}

var _aXValueGetTypeIDSymbol uintptr

// AXValueGetTypeID has an opaque C signature in discovered metadata.
// Call AXValueGetTypeIDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXValueGetTypeID
func AXValueGetTypeID() {
	panic("hiservices: symbol AXValueGetTypeID has opaque signature; use AXValueGetTypeIDSymbol() and a typed manual wrapper")
}

// AXValueGetTypeIDSymbol returns the raw symbol address for AXValueGetTypeID.
func AXValueGetTypeIDSymbol() uintptr {
	if _aXValueGetTypeIDSymbol == 0 {
		return 0
	}
	return _aXValueGetTypeIDSymbol
}

var _aXValueGetValueSymbol uintptr

// AXValueGetValue has an opaque C signature in discovered metadata.
// Call AXValueGetValueSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/AXValueGetValue
func AXValueGetValue() {
	panic("hiservices: symbol AXValueGetValue has opaque signature; use AXValueGetValueSymbol() and a typed manual wrapper")
}

// AXValueGetValueSymbol returns the raw symbol address for AXValueGetValue.
func AXValueGetValueSymbol() uintptr {
	if _aXValueGetValueSymbol == 0 {
		return 0
	}
	return _aXValueGetValueSymbol
}

var _applicationTypeGetSymbol uintptr

// ApplicationTypeGet has an opaque C signature in discovered metadata.
// Call ApplicationTypeGetSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ApplicationTypeGet
func ApplicationTypeGet() {
	panic("hiservices: symbol ApplicationTypeGet has opaque signature; use ApplicationTypeGetSymbol() and a typed manual wrapper")
}

// ApplicationTypeGetSymbol returns the raw symbol address for ApplicationTypeGet.
func ApplicationTypeGetSymbol() uintptr {
	if _applicationTypeGetSymbol == 0 {
		return 0
	}
	return _applicationTypeGetSymbol
}

var _applicationTypeSetSymbol uintptr

// ApplicationTypeSet has an opaque C signature in discovered metadata.
// Call ApplicationTypeSetSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ApplicationTypeSet
func ApplicationTypeSet() {
	panic("hiservices: symbol ApplicationTypeSet has opaque signature; use ApplicationTypeSetSymbol() and a typed manual wrapper")
}

// ApplicationTypeSetSymbol returns the raw symbol address for ApplicationTypeSet.
func ApplicationTypeSetSymbol() uintptr {
	if _applicationTypeSetSymbol == 0 {
		return 0
	}
	return _applicationTypeSetSymbol
}

var _cGPointInIconRefSymbol uintptr

// CGPointInIconRef has an opaque C signature in discovered metadata.
// Call CGPointInIconRefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CGPointInIconRef
func CGPointInIconRef() {
	panic("hiservices: symbol CGPointInIconRef has opaque signature; use CGPointInIconRefSymbol() and a typed manual wrapper")
}

// CGPointInIconRefSymbol returns the raw symbol address for CGPointInIconRef.
func CGPointInIconRefSymbol() uintptr {
	if _cGPointInIconRefSymbol == 0 {
		return 0
	}
	return _cGPointInIconRefSymbol
}

var _cGRectInIconRefSymbol uintptr

// CGRectInIconRef has an opaque C signature in discovered metadata.
// Call CGRectInIconRefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CGRectInIconRef
func CGRectInIconRef() {
	panic("hiservices: symbol CGRectInIconRef has opaque signature; use CGRectInIconRefSymbol() and a typed manual wrapper")
}

// CGRectInIconRefSymbol returns the raw symbol address for CGRectInIconRef.
func CGRectInIconRefSymbol() uintptr {
	if _cGRectInIconRefSymbol == 0 {
		return 0
	}
	return _cGRectInIconRefSymbol
}

var _copyLabelColorAndNameSymbol uintptr

// CopyLabelColorAndName has an opaque C signature in discovered metadata.
// Call CopyLabelColorAndNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CopyLabelColorAndName
func CopyLabelColorAndName() {
	panic("hiservices: symbol CopyLabelColorAndName has opaque signature; use CopyLabelColorAndNameSymbol() and a typed manual wrapper")
}

// CopyLabelColorAndNameSymbol returns the raw symbol address for CopyLabelColorAndName.
func CopyLabelColorAndNameSymbol() uintptr {
	if _copyLabelColorAndNameSymbol == 0 {
		return 0
	}
	return _copyLabelColorAndNameSymbol
}

var _copyProcessNameSymbol uintptr

// CopyProcessName has an opaque C signature in discovered metadata.
// Call CopyProcessNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CopyProcessName
func CopyProcessName() {
	panic("hiservices: symbol CopyProcessName has opaque signature; use CopyProcessNameSymbol() and a typed manual wrapper")
}

// CopyProcessNameSymbol returns the raw symbol address for CopyProcessName.
func CopyProcessNameSymbol() uintptr {
	if _copyProcessNameSymbol == 0 {
		return 0
	}
	return _copyProcessNameSymbol
}

var _coreAppearanceGetFontCGStyleRefSymbol uintptr

// CoreAppearanceGetFontCGStyleRef has an opaque C signature in discovered metadata.
// Call CoreAppearanceGetFontCGStyleRefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreAppearanceGetFontCGStyleRef
func CoreAppearanceGetFontCGStyleRef() {
	panic("hiservices: symbol CoreAppearanceGetFontCGStyleRef has opaque signature; use CoreAppearanceGetFontCGStyleRefSymbol() and a typed manual wrapper")
}

// CoreAppearanceGetFontCGStyleRefSymbol returns the raw symbol address for CoreAppearanceGetFontCGStyleRef.
func CoreAppearanceGetFontCGStyleRefSymbol() uintptr {
	if _coreAppearanceGetFontCGStyleRefSymbol == 0 {
		return 0
	}
	return _coreAppearanceGetFontCGStyleRefSymbol
}

var _coreAppearanceGetFontShadowOutsetsSymbol uintptr

// CoreAppearanceGetFontShadowOutsets has an opaque C signature in discovered metadata.
// Call CoreAppearanceGetFontShadowOutsetsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreAppearanceGetFontShadowOutsets
func CoreAppearanceGetFontShadowOutsets() {
	panic("hiservices: symbol CoreAppearanceGetFontShadowOutsets has opaque signature; use CoreAppearanceGetFontShadowOutsetsSymbol() and a typed manual wrapper")
}

// CoreAppearanceGetFontShadowOutsetsSymbol returns the raw symbol address for CoreAppearanceGetFontShadowOutsets.
func CoreAppearanceGetFontShadowOutsetsSymbol() uintptr {
	if _coreAppearanceGetFontShadowOutsetsSymbol == 0 {
		return 0
	}
	return _coreAppearanceGetFontShadowOutsetsSymbol
}

var _coreAppearanceGetFontSizeSymbol uintptr

// CoreAppearanceGetFontSize has an opaque C signature in discovered metadata.
// Call CoreAppearanceGetFontSizeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreAppearanceGetFontSize
func CoreAppearanceGetFontSize() {
	panic("hiservices: symbol CoreAppearanceGetFontSize has opaque signature; use CoreAppearanceGetFontSizeSymbol() and a typed manual wrapper")
}

// CoreAppearanceGetFontSizeSymbol returns the raw symbol address for CoreAppearanceGetFontSize.
func CoreAppearanceGetFontSizeSymbol() uintptr {
	if _coreAppearanceGetFontSizeSymbol == 0 {
		return 0
	}
	return _coreAppearanceGetFontSizeSymbol
}

var _coreAppearanceGetQDFontForScriptSymbol uintptr

// CoreAppearanceGetQDFontForScript has an opaque C signature in discovered metadata.
// Call CoreAppearanceGetQDFontForScriptSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreAppearanceGetQDFontForScript
func CoreAppearanceGetQDFontForScript() {
	panic("hiservices: symbol CoreAppearanceGetQDFontForScript has opaque signature; use CoreAppearanceGetQDFontForScriptSymbol() and a typed manual wrapper")
}

// CoreAppearanceGetQDFontForScriptSymbol returns the raw symbol address for CoreAppearanceGetQDFontForScript.
func CoreAppearanceGetQDFontForScriptSymbol() uintptr {
	if _coreAppearanceGetQDFontForScriptSymbol == 0 {
		return 0
	}
	return _coreAppearanceGetQDFontForScriptSymbol
}

var _coreCursorCopyImagesSymbol uintptr

// CoreCursorCopyImages has an opaque C signature in discovered metadata.
// Call CoreCursorCopyImagesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreCursorCopyImages
func CoreCursorCopyImages() {
	panic("hiservices: symbol CoreCursorCopyImages has opaque signature; use CoreCursorCopyImagesSymbol() and a typed manual wrapper")
}

// CoreCursorCopyImagesSymbol returns the raw symbol address for CoreCursorCopyImages.
func CoreCursorCopyImagesSymbol() uintptr {
	if _coreCursorCopyImagesSymbol == 0 {
		return 0
	}
	return _coreCursorCopyImagesSymbol
}

var _coreCursorGetDataSymbol uintptr

// CoreCursorGetData has an opaque C signature in discovered metadata.
// Call CoreCursorGetDataSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreCursorGetData
func CoreCursorGetData() {
	panic("hiservices: symbol CoreCursorGetData has opaque signature; use CoreCursorGetDataSymbol() and a typed manual wrapper")
}

// CoreCursorGetDataSymbol returns the raw symbol address for CoreCursorGetData.
func CoreCursorGetDataSymbol() uintptr {
	if _coreCursorGetDataSymbol == 0 {
		return 0
	}
	return _coreCursorGetDataSymbol
}

var _coreCursorGetDataSizeSymbol uintptr

// CoreCursorGetDataSize has an opaque C signature in discovered metadata.
// Call CoreCursorGetDataSizeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreCursorGetDataSize
func CoreCursorGetDataSize() {
	panic("hiservices: symbol CoreCursorGetDataSize has opaque signature; use CoreCursorGetDataSizeSymbol() and a typed manual wrapper")
}

// CoreCursorGetDataSizeSymbol returns the raw symbol address for CoreCursorGetDataSize.
func CoreCursorGetDataSizeSymbol() uintptr {
	if _coreCursorGetDataSizeSymbol == 0 {
		return 0
	}
	return _coreCursorGetDataSizeSymbol
}

var _coreCursorSetSymbol uintptr

// CoreCursorSet has an opaque C signature in discovered metadata.
// Call CoreCursorSetSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreCursorSet
func CoreCursorSet() {
	panic("hiservices: symbol CoreCursorSet has opaque signature; use CoreCursorSetSymbol() and a typed manual wrapper")
}

// CoreCursorSetSymbol returns the raw symbol address for CoreCursorSet.
func CoreCursorSetSymbol() uintptr {
	if _coreCursorSetSymbol == 0 {
		return 0
	}
	return _coreCursorSetSymbol
}

var _coreCursorSetAndReturnSeedSymbol uintptr

// CoreCursorSetAndReturnSeed has an opaque C signature in discovered metadata.
// Call CoreCursorSetAndReturnSeedSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreCursorSetAndReturnSeed
func CoreCursorSetAndReturnSeed() {
	panic("hiservices: symbol CoreCursorSetAndReturnSeed has opaque signature; use CoreCursorSetAndReturnSeedSymbol() and a typed manual wrapper")
}

// CoreCursorSetAndReturnSeedSymbol returns the raw symbol address for CoreCursorSetAndReturnSeed.
func CoreCursorSetAndReturnSeedSymbol() uintptr {
	if _coreCursorSetAndReturnSeedSymbol == 0 {
		return 0
	}
	return _coreCursorSetAndReturnSeedSymbol
}

var _coreCursorUnregisterAllSymbol uintptr

// CoreCursorUnregisterAll has an opaque C signature in discovered metadata.
// Call CoreCursorUnregisterAllSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreCursorUnregisterAll
func CoreCursorUnregisterAll() {
	panic("hiservices: symbol CoreCursorUnregisterAll has opaque signature; use CoreCursorUnregisterAllSymbol() and a typed manual wrapper")
}

// CoreCursorUnregisterAllSymbol returns the raw symbol address for CoreCursorUnregisterAll.
func CoreCursorUnregisterAllSymbol() uintptr {
	if _coreCursorUnregisterAllSymbol == 0 {
		return 0
	}
	return _coreCursorUnregisterAllSymbol
}

var _coreDockAddFileToDockSymbol uintptr

// CoreDockAddFileToDock has an opaque C signature in discovered metadata.
// Call CoreDockAddFileToDockSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockAddFileToDock
func CoreDockAddFileToDock() {
	panic("hiservices: symbol CoreDockAddFileToDock has opaque signature; use CoreDockAddFileToDockSymbol() and a typed manual wrapper")
}

// CoreDockAddFileToDockSymbol returns the raw symbol address for CoreDockAddFileToDock.
func CoreDockAddFileToDockSymbol() uintptr {
	if _coreDockAddFileToDockSymbol == 0 {
		return 0
	}
	return _coreDockAddFileToDockSymbol
}

var _coreDockBounceAppTileSymbol uintptr

// CoreDockBounceAppTile has an opaque C signature in discovered metadata.
// Call CoreDockBounceAppTileSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockBounceAppTile
func CoreDockBounceAppTile() {
	panic("hiservices: symbol CoreDockBounceAppTile has opaque signature; use CoreDockBounceAppTileSymbol() and a typed manual wrapper")
}

// CoreDockBounceAppTileSymbol returns the raw symbol address for CoreDockBounceAppTile.
func CoreDockBounceAppTileSymbol() uintptr {
	if _coreDockBounceAppTileSymbol == 0 {
		return 0
	}
	return _coreDockBounceAppTileSymbol
}

var _coreDockCompositeProcessImageSymbol uintptr

// CoreDockCompositeProcessImage has an opaque C signature in discovered metadata.
// Call CoreDockCompositeProcessImageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockCompositeProcessImage
func CoreDockCompositeProcessImage() {
	panic("hiservices: symbol CoreDockCompositeProcessImage has opaque signature; use CoreDockCompositeProcessImageSymbol() and a typed manual wrapper")
}

// CoreDockCompositeProcessImageSymbol returns the raw symbol address for CoreDockCompositeProcessImage.
func CoreDockCompositeProcessImageSymbol() uintptr {
	if _coreDockCompositeProcessImageSymbol == 0 {
		return 0
	}
	return _coreDockCompositeProcessImageSymbol
}

var _coreDockCopyPreferencesSymbol uintptr

// CoreDockCopyPreferences has an opaque C signature in discovered metadata.
// Call CoreDockCopyPreferencesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockCopyPreferences
func CoreDockCopyPreferences() {
	panic("hiservices: symbol CoreDockCopyPreferences has opaque signature; use CoreDockCopyPreferencesSymbol() and a typed manual wrapper")
}

// CoreDockCopyPreferencesSymbol returns the raw symbol address for CoreDockCopyPreferences.
func CoreDockCopyPreferencesSymbol() uintptr {
	if _coreDockCopyPreferencesSymbol == 0 {
		return 0
	}
	return _coreDockCopyPreferencesSymbol
}

var _coreDockCopyWorkspacesAppBindingsSymbol uintptr

// CoreDockCopyWorkspacesAppBindings has an opaque C signature in discovered metadata.
// Call CoreDockCopyWorkspacesAppBindingsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockCopyWorkspacesAppBindings
func CoreDockCopyWorkspacesAppBindings() {
	panic("hiservices: symbol CoreDockCopyWorkspacesAppBindings has opaque signature; use CoreDockCopyWorkspacesAppBindingsSymbol() and a typed manual wrapper")
}

// CoreDockCopyWorkspacesAppBindingsSymbol returns the raw symbol address for CoreDockCopyWorkspacesAppBindings.
func CoreDockCopyWorkspacesAppBindingsSymbol() uintptr {
	if _coreDockCopyWorkspacesAppBindingsSymbol == 0 {
		return 0
	}
	return _coreDockCopyWorkspacesAppBindingsSymbol
}

var _coreDockCreateDragTrashContextSymbol uintptr

// CoreDockCreateDragTrashContext has an opaque C signature in discovered metadata.
// Call CoreDockCreateDragTrashContextSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockCreateDragTrashContext
func CoreDockCreateDragTrashContext() {
	panic("hiservices: symbol CoreDockCreateDragTrashContext has opaque signature; use CoreDockCreateDragTrashContextSymbol() and a typed manual wrapper")
}

// CoreDockCreateDragTrashContextSymbol returns the raw symbol address for CoreDockCreateDragTrashContext.
func CoreDockCreateDragTrashContextSymbol() uintptr {
	if _coreDockCreateDragTrashContextSymbol == 0 {
		return 0
	}
	return _coreDockCreateDragTrashContextSymbol
}

var _coreDockDisableExposeKeysIfNecessarySymbol uintptr

// CoreDockDisableExposeKeysIfNecessary has an opaque C signature in discovered metadata.
// Call CoreDockDisableExposeKeysIfNecessarySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockDisableExposeKeysIfNecessary
func CoreDockDisableExposeKeysIfNecessary() {
	panic("hiservices: symbol CoreDockDisableExposeKeysIfNecessary has opaque signature; use CoreDockDisableExposeKeysIfNecessarySymbol() and a typed manual wrapper")
}

// CoreDockDisableExposeKeysIfNecessarySymbol returns the raw symbol address for CoreDockDisableExposeKeysIfNecessary.
func CoreDockDisableExposeKeysIfNecessarySymbol() uintptr {
	if _coreDockDisableExposeKeysIfNecessarySymbol == 0 {
		return 0
	}
	return _coreDockDisableExposeKeysIfNecessarySymbol
}

var _coreDockGetAutoHideEnabledSymbol uintptr

// CoreDockGetAutoHideEnabled has an opaque C signature in discovered metadata.
// Call CoreDockGetAutoHideEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetAutoHideEnabled
func CoreDockGetAutoHideEnabled() {
	panic("hiservices: symbol CoreDockGetAutoHideEnabled has opaque signature; use CoreDockGetAutoHideEnabledSymbol() and a typed manual wrapper")
}

// CoreDockGetAutoHideEnabledSymbol returns the raw symbol address for CoreDockGetAutoHideEnabled.
func CoreDockGetAutoHideEnabledSymbol() uintptr {
	if _coreDockGetAutoHideEnabledSymbol == 0 {
		return 0
	}
	return _coreDockGetAutoHideEnabledSymbol
}

var _coreDockGetContainerRectSymbol uintptr

// CoreDockGetContainerRect has an opaque C signature in discovered metadata.
// Call CoreDockGetContainerRectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetContainerRect
func CoreDockGetContainerRect() {
	panic("hiservices: symbol CoreDockGetContainerRect has opaque signature; use CoreDockGetContainerRectSymbol() and a typed manual wrapper")
}

// CoreDockGetContainerRectSymbol returns the raw symbol address for CoreDockGetContainerRect.
func CoreDockGetContainerRectSymbol() uintptr {
	if _coreDockGetContainerRectSymbol == 0 {
		return 0
	}
	return _coreDockGetContainerRectSymbol
}

var _coreDockGetDashboardInDockSymbol uintptr

// CoreDockGetDashboardInDock has an opaque C signature in discovered metadata.
// Call CoreDockGetDashboardInDockSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetDashboardInDock
func CoreDockGetDashboardInDock() {
	panic("hiservices: symbol CoreDockGetDashboardInDock has opaque signature; use CoreDockGetDashboardInDockSymbol() and a typed manual wrapper")
}

// CoreDockGetDashboardInDockSymbol returns the raw symbol address for CoreDockGetDashboardInDock.
func CoreDockGetDashboardInDockSymbol() uintptr {
	if _coreDockGetDashboardInDockSymbol == 0 {
		return 0
	}
	return _coreDockGetDashboardInDockSymbol
}

var _coreDockGetEffectSymbol uintptr

// CoreDockGetEffect has an opaque C signature in discovered metadata.
// Call CoreDockGetEffectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetEffect
func CoreDockGetEffect() {
	panic("hiservices: symbol CoreDockGetEffect has opaque signature; use CoreDockGetEffectSymbol() and a typed manual wrapper")
}

// CoreDockGetEffectSymbol returns the raw symbol address for CoreDockGetEffect.
func CoreDockGetEffectSymbol() uintptr {
	if _coreDockGetEffectSymbol == 0 {
		return 0
	}
	return _coreDockGetEffectSymbol
}

var _coreDockGetExposeCornerActionsSymbol uintptr

// CoreDockGetExposeCornerActions has an opaque C signature in discovered metadata.
// Call CoreDockGetExposeCornerActionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetExposeCornerActions
func CoreDockGetExposeCornerActions() {
	panic("hiservices: symbol CoreDockGetExposeCornerActions has opaque signature; use CoreDockGetExposeCornerActionsSymbol() and a typed manual wrapper")
}

// CoreDockGetExposeCornerActionsSymbol returns the raw symbol address for CoreDockGetExposeCornerActions.
func CoreDockGetExposeCornerActionsSymbol() uintptr {
	if _coreDockGetExposeCornerActionsSymbol == 0 {
		return 0
	}
	return _coreDockGetExposeCornerActionsSymbol
}

var _coreDockGetExposeCornerActionsWithModifiersSymbol uintptr

// CoreDockGetExposeCornerActionsWithModifiers has an opaque C signature in discovered metadata.
// Call CoreDockGetExposeCornerActionsWithModifiersSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetExposeCornerActionsWithModifiers
func CoreDockGetExposeCornerActionsWithModifiers() {
	panic("hiservices: symbol CoreDockGetExposeCornerActionsWithModifiers has opaque signature; use CoreDockGetExposeCornerActionsWithModifiersSymbol() and a typed manual wrapper")
}

// CoreDockGetExposeCornerActionsWithModifiersSymbol returns the raw symbol address for CoreDockGetExposeCornerActionsWithModifiers.
func CoreDockGetExposeCornerActionsWithModifiersSymbol() uintptr {
	if _coreDockGetExposeCornerActionsWithModifiersSymbol == 0 {
		return 0
	}
	return _coreDockGetExposeCornerActionsWithModifiersSymbol
}

var _coreDockGetItemDockContextSymbol uintptr

// CoreDockGetItemDockContext has an opaque C signature in discovered metadata.
// Call CoreDockGetItemDockContextSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetItemDockContext
func CoreDockGetItemDockContext() {
	panic("hiservices: symbol CoreDockGetItemDockContext has opaque signature; use CoreDockGetItemDockContextSymbol() and a typed manual wrapper")
}

// CoreDockGetItemDockContextSymbol returns the raw symbol address for CoreDockGetItemDockContext.
func CoreDockGetItemDockContextSymbol() uintptr {
	if _coreDockGetItemDockContextSymbol == 0 {
		return 0
	}
	return _coreDockGetItemDockContextSymbol
}

var _coreDockGetItemDockWindowSymbol uintptr

// CoreDockGetItemDockWindow has an opaque C signature in discovered metadata.
// Call CoreDockGetItemDockWindowSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetItemDockWindow
func CoreDockGetItemDockWindow() {
	panic("hiservices: symbol CoreDockGetItemDockWindow has opaque signature; use CoreDockGetItemDockWindowSymbol() and a typed manual wrapper")
}

// CoreDockGetItemDockWindowSymbol returns the raw symbol address for CoreDockGetItemDockWindow.
func CoreDockGetItemDockWindowSymbol() uintptr {
	if _coreDockGetItemDockWindowSymbol == 0 {
		return 0
	}
	return _coreDockGetItemDockWindowSymbol
}

var _coreDockGetMagnificationSizeSymbol uintptr

// CoreDockGetMagnificationSize has an opaque C signature in discovered metadata.
// Call CoreDockGetMagnificationSizeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetMagnificationSize
func CoreDockGetMagnificationSize() {
	panic("hiservices: symbol CoreDockGetMagnificationSize has opaque signature; use CoreDockGetMagnificationSizeSymbol() and a typed manual wrapper")
}

// CoreDockGetMagnificationSizeSymbol returns the raw symbol address for CoreDockGetMagnificationSize.
func CoreDockGetMagnificationSizeSymbol() uintptr {
	if _coreDockGetMagnificationSizeSymbol == 0 {
		return 0
	}
	return _coreDockGetMagnificationSizeSymbol
}

var _coreDockGetMinimizeInPlaceSymbol uintptr

// CoreDockGetMinimizeInPlace has an opaque C signature in discovered metadata.
// Call CoreDockGetMinimizeInPlaceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetMinimizeInPlace
func CoreDockGetMinimizeInPlace() {
	panic("hiservices: symbol CoreDockGetMinimizeInPlace has opaque signature; use CoreDockGetMinimizeInPlaceSymbol() and a typed manual wrapper")
}

// CoreDockGetMinimizeInPlaceSymbol returns the raw symbol address for CoreDockGetMinimizeInPlace.
func CoreDockGetMinimizeInPlaceSymbol() uintptr {
	if _coreDockGetMinimizeInPlaceSymbol == 0 {
		return 0
	}
	return _coreDockGetMinimizeInPlaceSymbol
}

var _coreDockGetOrientationAndPinningSymbol uintptr

// CoreDockGetOrientationAndPinning has an opaque C signature in discovered metadata.
// Call CoreDockGetOrientationAndPinningSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetOrientationAndPinning
func CoreDockGetOrientationAndPinning() {
	panic("hiservices: symbol CoreDockGetOrientationAndPinning has opaque signature; use CoreDockGetOrientationAndPinningSymbol() and a typed manual wrapper")
}

// CoreDockGetOrientationAndPinningSymbol returns the raw symbol address for CoreDockGetOrientationAndPinning.
func CoreDockGetOrientationAndPinningSymbol() uintptr {
	if _coreDockGetOrientationAndPinningSymbol == 0 {
		return 0
	}
	return _coreDockGetOrientationAndPinningSymbol
}

var _coreDockGetProcessContextSymbol uintptr

// CoreDockGetProcessContext has an opaque C signature in discovered metadata.
// Call CoreDockGetProcessContextSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetProcessContext
func CoreDockGetProcessContext() {
	panic("hiservices: symbol CoreDockGetProcessContext has opaque signature; use CoreDockGetProcessContextSymbol() and a typed manual wrapper")
}

// CoreDockGetProcessContextSymbol returns the raw symbol address for CoreDockGetProcessContext.
func CoreDockGetProcessContextSymbol() uintptr {
	if _coreDockGetProcessContextSymbol == 0 {
		return 0
	}
	return _coreDockGetProcessContextSymbol
}

var _coreDockGetProcessWindowSymbol uintptr

// CoreDockGetProcessWindow has an opaque C signature in discovered metadata.
// Call CoreDockGetProcessWindowSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetProcessWindow
func CoreDockGetProcessWindow() {
	panic("hiservices: symbol CoreDockGetProcessWindow has opaque signature; use CoreDockGetProcessWindowSymbol() and a typed manual wrapper")
}

// CoreDockGetProcessWindowSymbol returns the raw symbol address for CoreDockGetProcessWindow.
func CoreDockGetProcessWindowSymbol() uintptr {
	if _coreDockGetProcessWindowSymbol == 0 {
		return 0
	}
	return _coreDockGetProcessWindowSymbol
}

var _coreDockGetRectSymbol uintptr

// CoreDockGetRect has an opaque C signature in discovered metadata.
// Call CoreDockGetRectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetRect
func CoreDockGetRect() {
	panic("hiservices: symbol CoreDockGetRect has opaque signature; use CoreDockGetRectSymbol() and a typed manual wrapper")
}

// CoreDockGetRectSymbol returns the raw symbol address for CoreDockGetRect.
func CoreDockGetRectSymbol() uintptr {
	if _coreDockGetRectSymbol == 0 {
		return 0
	}
	return _coreDockGetRectSymbol
}

var _coreDockGetRectAndOrientationSymbol uintptr

// CoreDockGetRectAndOrientation has an opaque C signature in discovered metadata.
// Call CoreDockGetRectAndOrientationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetRectAndOrientation
func CoreDockGetRectAndOrientation() {
	panic("hiservices: symbol CoreDockGetRectAndOrientation has opaque signature; use CoreDockGetRectAndOrientationSymbol() and a typed manual wrapper")
}

// CoreDockGetRectAndOrientationSymbol returns the raw symbol address for CoreDockGetRectAndOrientation.
func CoreDockGetRectAndOrientationSymbol() uintptr {
	if _coreDockGetRectAndOrientationSymbol == 0 {
		return 0
	}
	return _coreDockGetRectAndOrientationSymbol
}

var _coreDockGetRectAndReasonSymbol uintptr

// CoreDockGetRectAndReason has an opaque C signature in discovered metadata.
// Call CoreDockGetRectAndReasonSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetRectAndReason
func CoreDockGetRectAndReason() {
	panic("hiservices: symbol CoreDockGetRectAndReason has opaque signature; use CoreDockGetRectAndReasonSymbol() and a typed manual wrapper")
}

// CoreDockGetRectAndReasonSymbol returns the raw symbol address for CoreDockGetRectAndReason.
func CoreDockGetRectAndReasonSymbol() uintptr {
	if _coreDockGetRectAndReasonSymbol == 0 {
		return 0
	}
	return _coreDockGetRectAndReasonSymbol
}

var _coreDockGetSpringLoadingTimeSymbol uintptr

// CoreDockGetSpringLoadingTime has an opaque C signature in discovered metadata.
// Call CoreDockGetSpringLoadingTimeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetSpringLoadingTime
func CoreDockGetSpringLoadingTime() {
	panic("hiservices: symbol CoreDockGetSpringLoadingTime has opaque signature; use CoreDockGetSpringLoadingTimeSymbol() and a typed manual wrapper")
}

// CoreDockGetSpringLoadingTimeSymbol returns the raw symbol address for CoreDockGetSpringLoadingTime.
func CoreDockGetSpringLoadingTimeSymbol() uintptr {
	if _coreDockGetSpringLoadingTimeSymbol == 0 {
		return 0
	}
	return _coreDockGetSpringLoadingTimeSymbol
}

var _coreDockGetTileSizeSymbol uintptr

// CoreDockGetTileSize has an opaque C signature in discovered metadata.
// Call CoreDockGetTileSizeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetTileSize
func CoreDockGetTileSize() {
	panic("hiservices: symbol CoreDockGetTileSize has opaque signature; use CoreDockGetTileSizeSymbol() and a typed manual wrapper")
}

// CoreDockGetTileSizeSymbol returns the raw symbol address for CoreDockGetTileSize.
func CoreDockGetTileSizeSymbol() uintptr {
	if _coreDockGetTileSizeSymbol == 0 {
		return 0
	}
	return _coreDockGetTileSizeSymbol
}

var _coreDockGetTrashWindowSymbol uintptr

// CoreDockGetTrashWindow has an opaque C signature in discovered metadata.
// Call CoreDockGetTrashWindowSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetTrashWindow
func CoreDockGetTrashWindow() {
	panic("hiservices: symbol CoreDockGetTrashWindow has opaque signature; use CoreDockGetTrashWindowSymbol() and a typed manual wrapper")
}

// CoreDockGetTrashWindowSymbol returns the raw symbol address for CoreDockGetTrashWindow.
func CoreDockGetTrashWindowSymbol() uintptr {
	if _coreDockGetTrashWindowSymbol == 0 {
		return 0
	}
	return _coreDockGetTrashWindowSymbol
}

var _coreDockGetWorkspacesCountSymbol uintptr

// CoreDockGetWorkspacesCount has an opaque C signature in discovered metadata.
// Call CoreDockGetWorkspacesCountSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetWorkspacesCount
func CoreDockGetWorkspacesCount() {
	panic("hiservices: symbol CoreDockGetWorkspacesCount has opaque signature; use CoreDockGetWorkspacesCountSymbol() and a typed manual wrapper")
}

// CoreDockGetWorkspacesCountSymbol returns the raw symbol address for CoreDockGetWorkspacesCount.
func CoreDockGetWorkspacesCountSymbol() uintptr {
	if _coreDockGetWorkspacesCountSymbol == 0 {
		return 0
	}
	return _coreDockGetWorkspacesCountSymbol
}

var _coreDockGetWorkspacesEnabledSymbol uintptr

// CoreDockGetWorkspacesEnabled has an opaque C signature in discovered metadata.
// Call CoreDockGetWorkspacesEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetWorkspacesEnabled
func CoreDockGetWorkspacesEnabled() {
	panic("hiservices: symbol CoreDockGetWorkspacesEnabled has opaque signature; use CoreDockGetWorkspacesEnabledSymbol() and a typed manual wrapper")
}

// CoreDockGetWorkspacesEnabledSymbol returns the raw symbol address for CoreDockGetWorkspacesEnabled.
func CoreDockGetWorkspacesEnabledSymbol() uintptr {
	if _coreDockGetWorkspacesEnabledSymbol == 0 {
		return 0
	}
	return _coreDockGetWorkspacesEnabledSymbol
}

var _coreDockGetWorkspacesKeyBindingsSymbol uintptr

// CoreDockGetWorkspacesKeyBindings has an opaque C signature in discovered metadata.
// Call CoreDockGetWorkspacesKeyBindingsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockGetWorkspacesKeyBindings
func CoreDockGetWorkspacesKeyBindings() {
	panic("hiservices: symbol CoreDockGetWorkspacesKeyBindings has opaque signature; use CoreDockGetWorkspacesKeyBindingsSymbol() and a typed manual wrapper")
}

// CoreDockGetWorkspacesKeyBindingsSymbol returns the raw symbol address for CoreDockGetWorkspacesKeyBindings.
func CoreDockGetWorkspacesKeyBindingsSymbol() uintptr {
	if _coreDockGetWorkspacesKeyBindingsSymbol == 0 {
		return 0
	}
	return _coreDockGetWorkspacesKeyBindingsSymbol
}

var _coreDockIsDockRunningSymbol uintptr

// CoreDockIsDockRunning has an opaque C signature in discovered metadata.
// Call CoreDockIsDockRunningSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockIsDockRunning
func CoreDockIsDockRunning() {
	panic("hiservices: symbol CoreDockIsDockRunning has opaque signature; use CoreDockIsDockRunningSymbol() and a typed manual wrapper")
}

// CoreDockIsDockRunningSymbol returns the raw symbol address for CoreDockIsDockRunning.
func CoreDockIsDockRunningSymbol() uintptr {
	if _coreDockIsDockRunningSymbol == 0 {
		return 0
	}
	return _coreDockIsDockRunningSymbol
}

var _coreDockIsLaunchAnimationsEnabledSymbol uintptr

// CoreDockIsLaunchAnimationsEnabled has an opaque C signature in discovered metadata.
// Call CoreDockIsLaunchAnimationsEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockIsLaunchAnimationsEnabled
func CoreDockIsLaunchAnimationsEnabled() {
	panic("hiservices: symbol CoreDockIsLaunchAnimationsEnabled has opaque signature; use CoreDockIsLaunchAnimationsEnabledSymbol() and a typed manual wrapper")
}

// CoreDockIsLaunchAnimationsEnabledSymbol returns the raw symbol address for CoreDockIsLaunchAnimationsEnabled.
func CoreDockIsLaunchAnimationsEnabledSymbol() uintptr {
	if _coreDockIsLaunchAnimationsEnabledSymbol == 0 {
		return 0
	}
	return _coreDockIsLaunchAnimationsEnabledSymbol
}

var _coreDockIsMagnificationEnabledSymbol uintptr

// CoreDockIsMagnificationEnabled has an opaque C signature in discovered metadata.
// Call CoreDockIsMagnificationEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockIsMagnificationEnabled
func CoreDockIsMagnificationEnabled() {
	panic("hiservices: symbol CoreDockIsMagnificationEnabled has opaque signature; use CoreDockIsMagnificationEnabledSymbol() and a typed manual wrapper")
}

// CoreDockIsMagnificationEnabledSymbol returns the raw symbol address for CoreDockIsMagnificationEnabled.
func CoreDockIsMagnificationEnabledSymbol() uintptr {
	if _coreDockIsMagnificationEnabledSymbol == 0 {
		return 0
	}
	return _coreDockIsMagnificationEnabledSymbol
}

var _coreDockIsSpringLoadingEnabledSymbol uintptr

// CoreDockIsSpringLoadingEnabled has an opaque C signature in discovered metadata.
// Call CoreDockIsSpringLoadingEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockIsSpringLoadingEnabled
func CoreDockIsSpringLoadingEnabled() {
	panic("hiservices: symbol CoreDockIsSpringLoadingEnabled has opaque signature; use CoreDockIsSpringLoadingEnabledSymbol() and a typed manual wrapper")
}

// CoreDockIsSpringLoadingEnabledSymbol returns the raw symbol address for CoreDockIsSpringLoadingEnabled.
func CoreDockIsSpringLoadingEnabledSymbol() uintptr {
	if _coreDockIsSpringLoadingEnabledSymbol == 0 {
		return 0
	}
	return _coreDockIsSpringLoadingEnabledSymbol
}

var _coreDockMinimizeItemWithTitleSymbol uintptr

// CoreDockMinimizeItemWithTitle has an opaque C signature in discovered metadata.
// Call CoreDockMinimizeItemWithTitleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockMinimizeItemWithTitle
func CoreDockMinimizeItemWithTitle() {
	panic("hiservices: symbol CoreDockMinimizeItemWithTitle has opaque signature; use CoreDockMinimizeItemWithTitleSymbol() and a typed manual wrapper")
}

// CoreDockMinimizeItemWithTitleSymbol returns the raw symbol address for CoreDockMinimizeItemWithTitle.
func CoreDockMinimizeItemWithTitleSymbol() uintptr {
	if _coreDockMinimizeItemWithTitleSymbol == 0 {
		return 0
	}
	return _coreDockMinimizeItemWithTitleSymbol
}

var _coreDockMinimizeItemWithTitleAsyncSymbol uintptr

// CoreDockMinimizeItemWithTitleAsync has an opaque C signature in discovered metadata.
// Call CoreDockMinimizeItemWithTitleAsyncSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockMinimizeItemWithTitleAsync
func CoreDockMinimizeItemWithTitleAsync() {
	panic("hiservices: symbol CoreDockMinimizeItemWithTitleAsync has opaque signature; use CoreDockMinimizeItemWithTitleAsyncSymbol() and a typed manual wrapper")
}

// CoreDockMinimizeItemWithTitleAsyncSymbol returns the raw symbol address for CoreDockMinimizeItemWithTitleAsync.
func CoreDockMinimizeItemWithTitleAsyncSymbol() uintptr {
	if _coreDockMinimizeItemWithTitleAsyncSymbol == 0 {
		return 0
	}
	return _coreDockMinimizeItemWithTitleAsyncSymbol
}

var _coreDockMinimizeItemsWithTitleSymbol uintptr

// CoreDockMinimizeItemsWithTitle has an opaque C signature in discovered metadata.
// Call CoreDockMinimizeItemsWithTitleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockMinimizeItemsWithTitle
func CoreDockMinimizeItemsWithTitle() {
	panic("hiservices: symbol CoreDockMinimizeItemsWithTitle has opaque signature; use CoreDockMinimizeItemsWithTitleSymbol() and a typed manual wrapper")
}

// CoreDockMinimizeItemsWithTitleSymbol returns the raw symbol address for CoreDockMinimizeItemsWithTitle.
func CoreDockMinimizeItemsWithTitleSymbol() uintptr {
	if _coreDockMinimizeItemsWithTitleSymbol == 0 {
		return 0
	}
	return _coreDockMinimizeItemsWithTitleSymbol
}

var _coreDockMinimizeItemsWithTitleAsyncSymbol uintptr

// CoreDockMinimizeItemsWithTitleAsync has an opaque C signature in discovered metadata.
// Call CoreDockMinimizeItemsWithTitleAsyncSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockMinimizeItemsWithTitleAsync
func CoreDockMinimizeItemsWithTitleAsync() {
	panic("hiservices: symbol CoreDockMinimizeItemsWithTitleAsync has opaque signature; use CoreDockMinimizeItemsWithTitleAsyncSymbol() and a typed manual wrapper")
}

// CoreDockMinimizeItemsWithTitleAsyncSymbol returns the raw symbol address for CoreDockMinimizeItemsWithTitleAsync.
func CoreDockMinimizeItemsWithTitleAsyncSymbol() uintptr {
	if _coreDockMinimizeItemsWithTitleAsyncSymbol == 0 {
		return 0
	}
	return _coreDockMinimizeItemsWithTitleAsyncSymbol
}

var _coreDockPreventCommunicationWithDockSymbol uintptr

// CoreDockPreventCommunicationWithDock has an opaque C signature in discovered metadata.
// Call CoreDockPreventCommunicationWithDockSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockPreventCommunicationWithDock
func CoreDockPreventCommunicationWithDock() {
	panic("hiservices: symbol CoreDockPreventCommunicationWithDock has opaque signature; use CoreDockPreventCommunicationWithDockSymbol() and a typed manual wrapper")
}

// CoreDockPreventCommunicationWithDockSymbol returns the raw symbol address for CoreDockPreventCommunicationWithDock.
func CoreDockPreventCommunicationWithDockSymbol() uintptr {
	if _coreDockPreventCommunicationWithDockSymbol == 0 {
		return 0
	}
	return _coreDockPreventCommunicationWithDockSymbol
}

var _coreDockRegisterClientWithRunLoopSymbol uintptr

// CoreDockRegisterClientWithRunLoop has an opaque C signature in discovered metadata.
// Call CoreDockRegisterClientWithRunLoopSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockRegisterClientWithRunLoop
func CoreDockRegisterClientWithRunLoop() {
	panic("hiservices: symbol CoreDockRegisterClientWithRunLoop has opaque signature; use CoreDockRegisterClientWithRunLoopSymbol() and a typed manual wrapper")
}

// CoreDockRegisterClientWithRunLoopSymbol returns the raw symbol address for CoreDockRegisterClientWithRunLoop.
func CoreDockRegisterClientWithRunLoopSymbol() uintptr {
	if _coreDockRegisterClientWithRunLoopSymbol == 0 {
		return 0
	}
	return _coreDockRegisterClientWithRunLoopSymbol
}

var _coreDockReleaseDragTrashContextSymbol uintptr

// CoreDockReleaseDragTrashContext has an opaque C signature in discovered metadata.
// Call CoreDockReleaseDragTrashContextSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockReleaseDragTrashContext
func CoreDockReleaseDragTrashContext() {
	panic("hiservices: symbol CoreDockReleaseDragTrashContext has opaque signature; use CoreDockReleaseDragTrashContextSymbol() and a typed manual wrapper")
}

// CoreDockReleaseDragTrashContextSymbol returns the raw symbol address for CoreDockReleaseDragTrashContext.
func CoreDockReleaseDragTrashContextSymbol() uintptr {
	if _coreDockReleaseDragTrashContextSymbol == 0 {
		return 0
	}
	return _coreDockReleaseDragTrashContextSymbol
}

var _coreDockReleaseItemDockContextSymbol uintptr

// CoreDockReleaseItemDockContext has an opaque C signature in discovered metadata.
// Call CoreDockReleaseItemDockContextSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockReleaseItemDockContext
func CoreDockReleaseItemDockContext() {
	panic("hiservices: symbol CoreDockReleaseItemDockContext has opaque signature; use CoreDockReleaseItemDockContextSymbol() and a typed manual wrapper")
}

// CoreDockReleaseItemDockContextSymbol returns the raw symbol address for CoreDockReleaseItemDockContext.
func CoreDockReleaseItemDockContextSymbol() uintptr {
	if _coreDockReleaseItemDockContextSymbol == 0 {
		return 0
	}
	return _coreDockReleaseItemDockContextSymbol
}

var _coreDockReleaseItemDockWindowSymbol uintptr

// CoreDockReleaseItemDockWindow has an opaque C signature in discovered metadata.
// Call CoreDockReleaseItemDockWindowSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockReleaseItemDockWindow
func CoreDockReleaseItemDockWindow() {
	panic("hiservices: symbol CoreDockReleaseItemDockWindow has opaque signature; use CoreDockReleaseItemDockWindowSymbol() and a typed manual wrapper")
}

// CoreDockReleaseItemDockWindowSymbol returns the raw symbol address for CoreDockReleaseItemDockWindow.
func CoreDockReleaseItemDockWindowSymbol() uintptr {
	if _coreDockReleaseItemDockWindowSymbol == 0 {
		return 0
	}
	return _coreDockReleaseItemDockWindowSymbol
}

var _coreDockReleaseProcessContextSymbol uintptr

// CoreDockReleaseProcessContext has an opaque C signature in discovered metadata.
// Call CoreDockReleaseProcessContextSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockReleaseProcessContext
func CoreDockReleaseProcessContext() {
	panic("hiservices: symbol CoreDockReleaseProcessContext has opaque signature; use CoreDockReleaseProcessContextSymbol() and a typed manual wrapper")
}

// CoreDockReleaseProcessContextSymbol returns the raw symbol address for CoreDockReleaseProcessContext.
func CoreDockReleaseProcessContextSymbol() uintptr {
	if _coreDockReleaseProcessContextSymbol == 0 {
		return 0
	}
	return _coreDockReleaseProcessContextSymbol
}

var _coreDockReleaseProcessWindowSymbol uintptr

// CoreDockReleaseProcessWindow has an opaque C signature in discovered metadata.
// Call CoreDockReleaseProcessWindowSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockReleaseProcessWindow
func CoreDockReleaseProcessWindow() {
	panic("hiservices: symbol CoreDockReleaseProcessWindow has opaque signature; use CoreDockReleaseProcessWindowSymbol() and a typed manual wrapper")
}

// CoreDockReleaseProcessWindowSymbol returns the raw symbol address for CoreDockReleaseProcessWindow.
func CoreDockReleaseProcessWindowSymbol() uintptr {
	if _coreDockReleaseProcessWindowSymbol == 0 {
		return 0
	}
	return _coreDockReleaseProcessWindowSymbol
}

var _coreDockRemoveItemSymbol uintptr

// CoreDockRemoveItem has an opaque C signature in discovered metadata.
// Call CoreDockRemoveItemSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockRemoveItem
func CoreDockRemoveItem() {
	panic("hiservices: symbol CoreDockRemoveItem has opaque signature; use CoreDockRemoveItemSymbol() and a typed manual wrapper")
}

// CoreDockRemoveItemSymbol returns the raw symbol address for CoreDockRemoveItem.
func CoreDockRemoveItemSymbol() uintptr {
	if _coreDockRemoveItemSymbol == 0 {
		return 0
	}
	return _coreDockRemoveItemSymbol
}

var _coreDockRenderWindowIntoContextSymbol uintptr

// CoreDockRenderWindowIntoContext has an opaque C signature in discovered metadata.
// Call CoreDockRenderWindowIntoContextSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockRenderWindowIntoContext
func CoreDockRenderWindowIntoContext() {
	panic("hiservices: symbol CoreDockRenderWindowIntoContext has opaque signature; use CoreDockRenderWindowIntoContextSymbol() and a typed manual wrapper")
}

// CoreDockRenderWindowIntoContextSymbol returns the raw symbol address for CoreDockRenderWindowIntoContext.
func CoreDockRenderWindowIntoContextSymbol() uintptr {
	if _coreDockRenderWindowIntoContextSymbol == 0 {
		return 0
	}
	return _coreDockRenderWindowIntoContextSymbol
}

var _coreDockRestoreItemSymbol uintptr

// CoreDockRestoreItem has an opaque C signature in discovered metadata.
// Call CoreDockRestoreItemSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockRestoreItem
func CoreDockRestoreItem() {
	panic("hiservices: symbol CoreDockRestoreItem has opaque signature; use CoreDockRestoreItemSymbol() and a typed manual wrapper")
}

// CoreDockRestoreItemSymbol returns the raw symbol address for CoreDockRestoreItem.
func CoreDockRestoreItemSymbol() uintptr {
	if _coreDockRestoreItemSymbol == 0 {
		return 0
	}
	return _coreDockRestoreItemSymbol
}

var _coreDockRestoreItemAsyncSymbol uintptr

// CoreDockRestoreItemAsync has an opaque C signature in discovered metadata.
// Call CoreDockRestoreItemAsyncSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockRestoreItemAsync
func CoreDockRestoreItemAsync() {
	panic("hiservices: symbol CoreDockRestoreItemAsync has opaque signature; use CoreDockRestoreItemAsyncSymbol() and a typed manual wrapper")
}

// CoreDockRestoreItemAsyncSymbol returns the raw symbol address for CoreDockRestoreItemAsync.
func CoreDockRestoreItemAsyncSymbol() uintptr {
	if _coreDockRestoreItemAsyncSymbol == 0 {
		return 0
	}
	return _coreDockRestoreItemAsyncSymbol
}

var _coreDockRestoreItemWithOrderSymbol uintptr

// CoreDockRestoreItemWithOrder has an opaque C signature in discovered metadata.
// Call CoreDockRestoreItemWithOrderSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockRestoreItemWithOrder
func CoreDockRestoreItemWithOrder() {
	panic("hiservices: symbol CoreDockRestoreItemWithOrder has opaque signature; use CoreDockRestoreItemWithOrderSymbol() and a typed manual wrapper")
}

// CoreDockRestoreItemWithOrderSymbol returns the raw symbol address for CoreDockRestoreItemWithOrder.
func CoreDockRestoreItemWithOrderSymbol() uintptr {
	if _coreDockRestoreItemWithOrderSymbol == 0 {
		return 0
	}
	return _coreDockRestoreItemWithOrderSymbol
}

var _coreDockRestoreItemWithOrderAsyncSymbol uintptr

// CoreDockRestoreItemWithOrderAsync has an opaque C signature in discovered metadata.
// Call CoreDockRestoreItemWithOrderAsyncSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockRestoreItemWithOrderAsync
func CoreDockRestoreItemWithOrderAsync() {
	panic("hiservices: symbol CoreDockRestoreItemWithOrderAsync has opaque signature; use CoreDockRestoreItemWithOrderAsyncSymbol() and a typed manual wrapper")
}

// CoreDockRestoreItemWithOrderAsyncSymbol returns the raw symbol address for CoreDockRestoreItemWithOrderAsync.
func CoreDockRestoreItemWithOrderAsyncSymbol() uintptr {
	if _coreDockRestoreItemWithOrderAsyncSymbol == 0 {
		return 0
	}
	return _coreDockRestoreItemWithOrderAsyncSymbol
}

var _coreDockRestoreProcessImageSymbol uintptr

// CoreDockRestoreProcessImage has an opaque C signature in discovered metadata.
// Call CoreDockRestoreProcessImageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockRestoreProcessImage
func CoreDockRestoreProcessImage() {
	panic("hiservices: symbol CoreDockRestoreProcessImage has opaque signature; use CoreDockRestoreProcessImageSymbol() and a typed manual wrapper")
}

// CoreDockRestoreProcessImageSymbol returns the raw symbol address for CoreDockRestoreProcessImage.
func CoreDockRestoreProcessImageSymbol() uintptr {
	if _coreDockRestoreProcessImageSymbol == 0 {
		return 0
	}
	return _coreDockRestoreProcessImageSymbol
}

var _coreDockRevealWindowForShowDesktopSymbol uintptr

// CoreDockRevealWindowForShowDesktop has an opaque C signature in discovered metadata.
// Call CoreDockRevealWindowForShowDesktopSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockRevealWindowForShowDesktop
func CoreDockRevealWindowForShowDesktop() {
	panic("hiservices: symbol CoreDockRevealWindowForShowDesktop has opaque signature; use CoreDockRevealWindowForShowDesktopSymbol() and a typed manual wrapper")
}

// CoreDockRevealWindowForShowDesktopSymbol returns the raw symbol address for CoreDockRevealWindowForShowDesktop.
func CoreDockRevealWindowForShowDesktopSymbol() uintptr {
	if _coreDockRevealWindowForShowDesktopSymbol == 0 {
		return 0
	}
	return _coreDockRevealWindowForShowDesktopSymbol
}

var _coreDockSendDragWindowMessageSymbol uintptr

// CoreDockSendDragWindowMessage has an opaque C signature in discovered metadata.
// Call CoreDockSendDragWindowMessageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSendDragWindowMessage
func CoreDockSendDragWindowMessage() {
	panic("hiservices: symbol CoreDockSendDragWindowMessage has opaque signature; use CoreDockSendDragWindowMessageSymbol() and a typed manual wrapper")
}

// CoreDockSendDragWindowMessageSymbol returns the raw symbol address for CoreDockSendDragWindowMessage.
func CoreDockSendDragWindowMessageSymbol() uintptr {
	if _coreDockSendDragWindowMessageSymbol == 0 {
		return 0
	}
	return _coreDockSendDragWindowMessageSymbol
}

var _coreDockSendNotificationSymbol uintptr

// CoreDockSendNotification has an opaque C signature in discovered metadata.
// Call CoreDockSendNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSendNotification
func CoreDockSendNotification() {
	panic("hiservices: symbol CoreDockSendNotification has opaque signature; use CoreDockSendNotificationSymbol() and a typed manual wrapper")
}

// CoreDockSendNotificationSymbol returns the raw symbol address for CoreDockSendNotification.
func CoreDockSendNotificationSymbol() uintptr {
	if _coreDockSendNotificationSymbol == 0 {
		return 0
	}
	return _coreDockSendNotificationSymbol
}

var _coreDockSetAutoHideEnabledSymbol uintptr

// CoreDockSetAutoHideEnabled has an opaque C signature in discovered metadata.
// Call CoreDockSetAutoHideEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetAutoHideEnabled
func CoreDockSetAutoHideEnabled() {
	panic("hiservices: symbol CoreDockSetAutoHideEnabled has opaque signature; use CoreDockSetAutoHideEnabledSymbol() and a typed manual wrapper")
}

// CoreDockSetAutoHideEnabledSymbol returns the raw symbol address for CoreDockSetAutoHideEnabled.
func CoreDockSetAutoHideEnabledSymbol() uintptr {
	if _coreDockSetAutoHideEnabledSymbol == 0 {
		return 0
	}
	return _coreDockSetAutoHideEnabledSymbol
}

var _coreDockSetDashboardInDockSymbol uintptr

// CoreDockSetDashboardInDock has an opaque C signature in discovered metadata.
// Call CoreDockSetDashboardInDockSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetDashboardInDock
func CoreDockSetDashboardInDock() {
	panic("hiservices: symbol CoreDockSetDashboardInDock has opaque signature; use CoreDockSetDashboardInDockSymbol() and a typed manual wrapper")
}

// CoreDockSetDashboardInDockSymbol returns the raw symbol address for CoreDockSetDashboardInDock.
func CoreDockSetDashboardInDockSymbol() uintptr {
	if _coreDockSetDashboardInDockSymbol == 0 {
		return 0
	}
	return _coreDockSetDashboardInDockSymbol
}

var _coreDockSetDragStatusSymbol uintptr

// CoreDockSetDragStatus has an opaque C signature in discovered metadata.
// Call CoreDockSetDragStatusSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetDragStatus
func CoreDockSetDragStatus() {
	panic("hiservices: symbol CoreDockSetDragStatus has opaque signature; use CoreDockSetDragStatusSymbol() and a typed manual wrapper")
}

// CoreDockSetDragStatusSymbol returns the raw symbol address for CoreDockSetDragStatus.
func CoreDockSetDragStatusSymbol() uintptr {
	if _coreDockSetDragStatusSymbol == 0 {
		return 0
	}
	return _coreDockSetDragStatusSymbol
}

var _coreDockSetEffectSymbol uintptr

// CoreDockSetEffect has an opaque C signature in discovered metadata.
// Call CoreDockSetEffectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetEffect
func CoreDockSetEffect() {
	panic("hiservices: symbol CoreDockSetEffect has opaque signature; use CoreDockSetEffectSymbol() and a typed manual wrapper")
}

// CoreDockSetEffectSymbol returns the raw symbol address for CoreDockSetEffect.
func CoreDockSetEffectSymbol() uintptr {
	if _coreDockSetEffectSymbol == 0 {
		return 0
	}
	return _coreDockSetEffectSymbol
}

var _coreDockSetExposeCornerActionSymbol uintptr

// CoreDockSetExposeCornerAction has an opaque C signature in discovered metadata.
// Call CoreDockSetExposeCornerActionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetExposeCornerAction
func CoreDockSetExposeCornerAction() {
	panic("hiservices: symbol CoreDockSetExposeCornerAction has opaque signature; use CoreDockSetExposeCornerActionSymbol() and a typed manual wrapper")
}

// CoreDockSetExposeCornerActionSymbol returns the raw symbol address for CoreDockSetExposeCornerAction.
func CoreDockSetExposeCornerActionSymbol() uintptr {
	if _coreDockSetExposeCornerActionSymbol == 0 {
		return 0
	}
	return _coreDockSetExposeCornerActionSymbol
}

var _coreDockSetExposeCornerActionWithModifierSymbol uintptr

// CoreDockSetExposeCornerActionWithModifier has an opaque C signature in discovered metadata.
// Call CoreDockSetExposeCornerActionWithModifierSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetExposeCornerActionWithModifier
func CoreDockSetExposeCornerActionWithModifier() {
	panic("hiservices: symbol CoreDockSetExposeCornerActionWithModifier has opaque signature; use CoreDockSetExposeCornerActionWithModifierSymbol() and a typed manual wrapper")
}

// CoreDockSetExposeCornerActionWithModifierSymbol returns the raw symbol address for CoreDockSetExposeCornerActionWithModifier.
func CoreDockSetExposeCornerActionWithModifierSymbol() uintptr {
	if _coreDockSetExposeCornerActionWithModifierSymbol == 0 {
		return 0
	}
	return _coreDockSetExposeCornerActionWithModifierSymbol
}

var _coreDockSetItemTitleSymbol uintptr

// CoreDockSetItemTitle has an opaque C signature in discovered metadata.
// Call CoreDockSetItemTitleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetItemTitle
func CoreDockSetItemTitle() {
	panic("hiservices: symbol CoreDockSetItemTitle has opaque signature; use CoreDockSetItemTitleSymbol() and a typed manual wrapper")
}

// CoreDockSetItemTitleSymbol returns the raw symbol address for CoreDockSetItemTitle.
func CoreDockSetItemTitleSymbol() uintptr {
	if _coreDockSetItemTitleSymbol == 0 {
		return 0
	}
	return _coreDockSetItemTitleSymbol
}

var _coreDockSetLaunchAnimationsEnabledSymbol uintptr

// CoreDockSetLaunchAnimationsEnabled has an opaque C signature in discovered metadata.
// Call CoreDockSetLaunchAnimationsEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetLaunchAnimationsEnabled
func CoreDockSetLaunchAnimationsEnabled() {
	panic("hiservices: symbol CoreDockSetLaunchAnimationsEnabled has opaque signature; use CoreDockSetLaunchAnimationsEnabledSymbol() and a typed manual wrapper")
}

// CoreDockSetLaunchAnimationsEnabledSymbol returns the raw symbol address for CoreDockSetLaunchAnimationsEnabled.
func CoreDockSetLaunchAnimationsEnabledSymbol() uintptr {
	if _coreDockSetLaunchAnimationsEnabledSymbol == 0 {
		return 0
	}
	return _coreDockSetLaunchAnimationsEnabledSymbol
}

var _coreDockSetMagnificationEnabledSymbol uintptr

// CoreDockSetMagnificationEnabled has an opaque C signature in discovered metadata.
// Call CoreDockSetMagnificationEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetMagnificationEnabled
func CoreDockSetMagnificationEnabled() {
	panic("hiservices: symbol CoreDockSetMagnificationEnabled has opaque signature; use CoreDockSetMagnificationEnabledSymbol() and a typed manual wrapper")
}

// CoreDockSetMagnificationEnabledSymbol returns the raw symbol address for CoreDockSetMagnificationEnabled.
func CoreDockSetMagnificationEnabledSymbol() uintptr {
	if _coreDockSetMagnificationEnabledSymbol == 0 {
		return 0
	}
	return _coreDockSetMagnificationEnabledSymbol
}

var _coreDockSetMagnificationSizeSymbol uintptr

// CoreDockSetMagnificationSize has an opaque C signature in discovered metadata.
// Call CoreDockSetMagnificationSizeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetMagnificationSize
func CoreDockSetMagnificationSize() {
	panic("hiservices: symbol CoreDockSetMagnificationSize has opaque signature; use CoreDockSetMagnificationSizeSymbol() and a typed manual wrapper")
}

// CoreDockSetMagnificationSizeSymbol returns the raw symbol address for CoreDockSetMagnificationSize.
func CoreDockSetMagnificationSizeSymbol() uintptr {
	if _coreDockSetMagnificationSizeSymbol == 0 {
		return 0
	}
	return _coreDockSetMagnificationSizeSymbol
}

var _coreDockSetMiniViewSymbol uintptr

// CoreDockSetMiniView has an opaque C signature in discovered metadata.
// Call CoreDockSetMiniViewSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetMiniView
func CoreDockSetMiniView() {
	panic("hiservices: symbol CoreDockSetMiniView has opaque signature; use CoreDockSetMiniViewSymbol() and a typed manual wrapper")
}

// CoreDockSetMiniViewSymbol returns the raw symbol address for CoreDockSetMiniView.
func CoreDockSetMiniViewSymbol() uintptr {
	if _coreDockSetMiniViewSymbol == 0 {
		return 0
	}
	return _coreDockSetMiniViewSymbol
}

var _coreDockSetMinimizeInPlaceSymbol uintptr

// CoreDockSetMinimizeInPlace has an opaque C signature in discovered metadata.
// Call CoreDockSetMinimizeInPlaceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetMinimizeInPlace
func CoreDockSetMinimizeInPlace() {
	panic("hiservices: symbol CoreDockSetMinimizeInPlace has opaque signature; use CoreDockSetMinimizeInPlaceSymbol() and a typed manual wrapper")
}

// CoreDockSetMinimizeInPlaceSymbol returns the raw symbol address for CoreDockSetMinimizeInPlace.
func CoreDockSetMinimizeInPlaceSymbol() uintptr {
	if _coreDockSetMinimizeInPlaceSymbol == 0 {
		return 0
	}
	return _coreDockSetMinimizeInPlaceSymbol
}

var _coreDockSetOrientationAndPinningSymbol uintptr

// CoreDockSetOrientationAndPinning has an opaque C signature in discovered metadata.
// Call CoreDockSetOrientationAndPinningSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetOrientationAndPinning
func CoreDockSetOrientationAndPinning() {
	panic("hiservices: symbol CoreDockSetOrientationAndPinning has opaque signature; use CoreDockSetOrientationAndPinningSymbol() and a typed manual wrapper")
}

// CoreDockSetOrientationAndPinningSymbol returns the raw symbol address for CoreDockSetOrientationAndPinning.
func CoreDockSetOrientationAndPinningSymbol() uintptr {
	if _coreDockSetOrientationAndPinningSymbol == 0 {
		return 0
	}
	return _coreDockSetOrientationAndPinningSymbol
}

var _coreDockSetPreferencesSymbol uintptr

// CoreDockSetPreferences has an opaque C signature in discovered metadata.
// Call CoreDockSetPreferencesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetPreferences
func CoreDockSetPreferences() {
	panic("hiservices: symbol CoreDockSetPreferences has opaque signature; use CoreDockSetPreferencesSymbol() and a typed manual wrapper")
}

// CoreDockSetPreferencesSymbol returns the raw symbol address for CoreDockSetPreferences.
func CoreDockSetPreferencesSymbol() uintptr {
	if _coreDockSetPreferencesSymbol == 0 {
		return 0
	}
	return _coreDockSetPreferencesSymbol
}

var _coreDockSetProcessImageSymbol uintptr

// CoreDockSetProcessImage has an opaque C signature in discovered metadata.
// Call CoreDockSetProcessImageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetProcessImage
func CoreDockSetProcessImage() {
	panic("hiservices: symbol CoreDockSetProcessImage has opaque signature; use CoreDockSetProcessImageSymbol() and a typed manual wrapper")
}

// CoreDockSetProcessImageSymbol returns the raw symbol address for CoreDockSetProcessImage.
func CoreDockSetProcessImageSymbol() uintptr {
	if _coreDockSetProcessImageSymbol == 0 {
		return 0
	}
	return _coreDockSetProcessImageSymbol
}

var _coreDockSetProcessLabelSymbol uintptr

// CoreDockSetProcessLabel has an opaque C signature in discovered metadata.
// Call CoreDockSetProcessLabelSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetProcessLabel
func CoreDockSetProcessLabel() {
	panic("hiservices: symbol CoreDockSetProcessLabel has opaque signature; use CoreDockSetProcessLabelSymbol() and a typed manual wrapper")
}

// CoreDockSetProcessLabelSymbol returns the raw symbol address for CoreDockSetProcessLabel.
func CoreDockSetProcessLabelSymbol() uintptr {
	if _coreDockSetProcessLabelSymbol == 0 {
		return 0
	}
	return _coreDockSetProcessLabelSymbol
}

var _coreDockSetProcessOpenRecentsSymbol uintptr

// CoreDockSetProcessOpenRecents has an opaque C signature in discovered metadata.
// Call CoreDockSetProcessOpenRecentsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetProcessOpenRecents
func CoreDockSetProcessOpenRecents() {
	panic("hiservices: symbol CoreDockSetProcessOpenRecents has opaque signature; use CoreDockSetProcessOpenRecentsSymbol() and a typed manual wrapper")
}

// CoreDockSetProcessOpenRecentsSymbol returns the raw symbol address for CoreDockSetProcessOpenRecents.
func CoreDockSetProcessOpenRecentsSymbol() uintptr {
	if _coreDockSetProcessOpenRecentsSymbol == 0 {
		return 0
	}
	return _coreDockSetProcessOpenRecentsSymbol
}

var _coreDockSetShowDesktopCallbackSymbol uintptr

// CoreDockSetShowDesktopCallback has an opaque C signature in discovered metadata.
// Call CoreDockSetShowDesktopCallbackSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetShowDesktopCallback
func CoreDockSetShowDesktopCallback() {
	panic("hiservices: symbol CoreDockSetShowDesktopCallback has opaque signature; use CoreDockSetShowDesktopCallbackSymbol() and a typed manual wrapper")
}

// CoreDockSetShowDesktopCallbackSymbol returns the raw symbol address for CoreDockSetShowDesktopCallback.
func CoreDockSetShowDesktopCallbackSymbol() uintptr {
	if _coreDockSetShowDesktopCallbackSymbol == 0 {
		return 0
	}
	return _coreDockSetShowDesktopCallbackSymbol
}

var _coreDockSetSpringLoadingEnabledSymbol uintptr

// CoreDockSetSpringLoadingEnabled has an opaque C signature in discovered metadata.
// Call CoreDockSetSpringLoadingEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetSpringLoadingEnabled
func CoreDockSetSpringLoadingEnabled() {
	panic("hiservices: symbol CoreDockSetSpringLoadingEnabled has opaque signature; use CoreDockSetSpringLoadingEnabledSymbol() and a typed manual wrapper")
}

// CoreDockSetSpringLoadingEnabledSymbol returns the raw symbol address for CoreDockSetSpringLoadingEnabled.
func CoreDockSetSpringLoadingEnabledSymbol() uintptr {
	if _coreDockSetSpringLoadingEnabledSymbol == 0 {
		return 0
	}
	return _coreDockSetSpringLoadingEnabledSymbol
}

var _coreDockSetSpringLoadingTimeSymbol uintptr

// CoreDockSetSpringLoadingTime has an opaque C signature in discovered metadata.
// Call CoreDockSetSpringLoadingTimeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetSpringLoadingTime
func CoreDockSetSpringLoadingTime() {
	panic("hiservices: symbol CoreDockSetSpringLoadingTime has opaque signature; use CoreDockSetSpringLoadingTimeSymbol() and a typed manual wrapper")
}

// CoreDockSetSpringLoadingTimeSymbol returns the raw symbol address for CoreDockSetSpringLoadingTime.
func CoreDockSetSpringLoadingTimeSymbol() uintptr {
	if _coreDockSetSpringLoadingTimeSymbol == 0 {
		return 0
	}
	return _coreDockSetSpringLoadingTimeSymbol
}

var _coreDockSetSpringWindowCallbacksSymbol uintptr

// CoreDockSetSpringWindowCallbacks has an opaque C signature in discovered metadata.
// Call CoreDockSetSpringWindowCallbacksSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetSpringWindowCallbacks
func CoreDockSetSpringWindowCallbacks() {
	panic("hiservices: symbol CoreDockSetSpringWindowCallbacks has opaque signature; use CoreDockSetSpringWindowCallbacksSymbol() and a typed manual wrapper")
}

// CoreDockSetSpringWindowCallbacksSymbol returns the raw symbol address for CoreDockSetSpringWindowCallbacks.
func CoreDockSetSpringWindowCallbacksSymbol() uintptr {
	if _coreDockSetSpringWindowCallbacksSymbol == 0 {
		return 0
	}
	return _coreDockSetSpringWindowCallbacksSymbol
}

var _coreDockSetTileSizeSymbol uintptr

// CoreDockSetTileSize has an opaque C signature in discovered metadata.
// Call CoreDockSetTileSizeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetTileSize
func CoreDockSetTileSize() {
	panic("hiservices: symbol CoreDockSetTileSize has opaque signature; use CoreDockSetTileSizeSymbol() and a typed manual wrapper")
}

// CoreDockSetTileSizeSymbol returns the raw symbol address for CoreDockSetTileSize.
func CoreDockSetTileSizeSymbol() uintptr {
	if _coreDockSetTileSizeSymbol == 0 {
		return 0
	}
	return _coreDockSetTileSizeSymbol
}

var _coreDockSetTrashFullSymbol uintptr

// CoreDockSetTrashFull has an opaque C signature in discovered metadata.
// Call CoreDockSetTrashFullSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetTrashFull
func CoreDockSetTrashFull() {
	panic("hiservices: symbol CoreDockSetTrashFull has opaque signature; use CoreDockSetTrashFullSymbol() and a typed manual wrapper")
}

// CoreDockSetTrashFullSymbol returns the raw symbol address for CoreDockSetTrashFull.
func CoreDockSetTrashFullSymbol() uintptr {
	if _coreDockSetTrashFullSymbol == 0 {
		return 0
	}
	return _coreDockSetTrashFullSymbol
}

var _coreDockSetWindowLabelSymbol uintptr

// CoreDockSetWindowLabel has an opaque C signature in discovered metadata.
// Call CoreDockSetWindowLabelSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetWindowLabel
func CoreDockSetWindowLabel() {
	panic("hiservices: symbol CoreDockSetWindowLabel has opaque signature; use CoreDockSetWindowLabelSymbol() and a typed manual wrapper")
}

// CoreDockSetWindowLabelSymbol returns the raw symbol address for CoreDockSetWindowLabel.
func CoreDockSetWindowLabelSymbol() uintptr {
	if _coreDockSetWindowLabelSymbol == 0 {
		return 0
	}
	return _coreDockSetWindowLabelSymbol
}

var _coreDockSetWorkspacesAppBindingsSymbol uintptr

// CoreDockSetWorkspacesAppBindings has an opaque C signature in discovered metadata.
// Call CoreDockSetWorkspacesAppBindingsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetWorkspacesAppBindings
func CoreDockSetWorkspacesAppBindings() {
	panic("hiservices: symbol CoreDockSetWorkspacesAppBindings has opaque signature; use CoreDockSetWorkspacesAppBindingsSymbol() and a typed manual wrapper")
}

// CoreDockSetWorkspacesAppBindingsSymbol returns the raw symbol address for CoreDockSetWorkspacesAppBindings.
func CoreDockSetWorkspacesAppBindingsSymbol() uintptr {
	if _coreDockSetWorkspacesAppBindingsSymbol == 0 {
		return 0
	}
	return _coreDockSetWorkspacesAppBindingsSymbol
}

var _coreDockSetWorkspacesCountSymbol uintptr

// CoreDockSetWorkspacesCount has an opaque C signature in discovered metadata.
// Call CoreDockSetWorkspacesCountSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetWorkspacesCount
func CoreDockSetWorkspacesCount() {
	panic("hiservices: symbol CoreDockSetWorkspacesCount has opaque signature; use CoreDockSetWorkspacesCountSymbol() and a typed manual wrapper")
}

// CoreDockSetWorkspacesCountSymbol returns the raw symbol address for CoreDockSetWorkspacesCount.
func CoreDockSetWorkspacesCountSymbol() uintptr {
	if _coreDockSetWorkspacesCountSymbol == 0 {
		return 0
	}
	return _coreDockSetWorkspacesCountSymbol
}

var _coreDockSetWorkspacesEnabledSymbol uintptr

// CoreDockSetWorkspacesEnabled has an opaque C signature in discovered metadata.
// Call CoreDockSetWorkspacesEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetWorkspacesEnabled
func CoreDockSetWorkspacesEnabled() {
	panic("hiservices: symbol CoreDockSetWorkspacesEnabled has opaque signature; use CoreDockSetWorkspacesEnabledSymbol() and a typed manual wrapper")
}

// CoreDockSetWorkspacesEnabledSymbol returns the raw symbol address for CoreDockSetWorkspacesEnabled.
func CoreDockSetWorkspacesEnabledSymbol() uintptr {
	if _coreDockSetWorkspacesEnabledSymbol == 0 {
		return 0
	}
	return _coreDockSetWorkspacesEnabledSymbol
}

var _coreDockSetWorkspacesKeyBindingsSymbol uintptr

// CoreDockSetWorkspacesKeyBindings has an opaque C signature in discovered metadata.
// Call CoreDockSetWorkspacesKeyBindingsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockSetWorkspacesKeyBindings
func CoreDockSetWorkspacesKeyBindings() {
	panic("hiservices: symbol CoreDockSetWorkspacesKeyBindings has opaque signature; use CoreDockSetWorkspacesKeyBindingsSymbol() and a typed manual wrapper")
}

// CoreDockSetWorkspacesKeyBindingsSymbol returns the raw symbol address for CoreDockSetWorkspacesKeyBindings.
func CoreDockSetWorkspacesKeyBindingsSymbol() uintptr {
	if _coreDockSetWorkspacesKeyBindingsSymbol == 0 {
		return 0
	}
	return _coreDockSetWorkspacesKeyBindingsSymbol
}

var _coreDockUpdateWindowSymbol uintptr

// CoreDockUpdateWindow has an opaque C signature in discovered metadata.
// Call CoreDockUpdateWindowSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDockUpdateWindow
func CoreDockUpdateWindow() {
	panic("hiservices: symbol CoreDockUpdateWindow has opaque signature; use CoreDockUpdateWindowSymbol() and a typed manual wrapper")
}

// CoreDockUpdateWindowSymbol returns the raw symbol address for CoreDockUpdateWindow.
func CoreDockUpdateWindowSymbol() uintptr {
	if _coreDockUpdateWindowSymbol == 0 {
		return 0
	}
	return _coreDockUpdateWindowSymbol
}

var _coreDragAcceptDragSymbol uintptr

// CoreDragAcceptDrag has an opaque C signature in discovered metadata.
// Call CoreDragAcceptDragSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragAcceptDrag
func CoreDragAcceptDrag() {
	panic("hiservices: symbol CoreDragAcceptDrag has opaque signature; use CoreDragAcceptDragSymbol() and a typed manual wrapper")
}

// CoreDragAcceptDragSymbol returns the raw symbol address for CoreDragAcceptDrag.
func CoreDragAcceptDragSymbol() uintptr {
	if _coreDragAcceptDragSymbol == 0 {
		return 0
	}
	return _coreDragAcceptDragSymbol
}

var _coreDragCancelDragSymbol uintptr

// CoreDragCancelDrag has an opaque C signature in discovered metadata.
// Call CoreDragCancelDragSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragCancelDrag
func CoreDragCancelDrag() {
	panic("hiservices: symbol CoreDragCancelDrag has opaque signature; use CoreDragCancelDragSymbol() and a typed manual wrapper")
}

// CoreDragCancelDragSymbol returns the raw symbol address for CoreDragCancelDrag.
func CoreDragCancelDragSymbol() uintptr {
	if _coreDragCancelDragSymbol == 0 {
		return 0
	}
	return _coreDragCancelDragSymbol
}

var _coreDragChangeBehaviorsSymbol uintptr

// CoreDragChangeBehaviors has an opaque C signature in discovered metadata.
// Call CoreDragChangeBehaviorsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragChangeBehaviors
func CoreDragChangeBehaviors() {
	panic("hiservices: symbol CoreDragChangeBehaviors has opaque signature; use CoreDragChangeBehaviorsSymbol() and a typed manual wrapper")
}

// CoreDragChangeBehaviorsSymbol returns the raw symbol address for CoreDragChangeBehaviors.
func CoreDragChangeBehaviorsSymbol() uintptr {
	if _coreDragChangeBehaviorsSymbol == 0 {
		return 0
	}
	return _coreDragChangeBehaviorsSymbol
}

var _coreDragCleanDragStateSymbol uintptr

// CoreDragCleanDragState has an opaque C signature in discovered metadata.
// Call CoreDragCleanDragStateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragCleanDragState
func CoreDragCleanDragState() {
	panic("hiservices: symbol CoreDragCleanDragState has opaque signature; use CoreDragCleanDragStateSymbol() and a typed manual wrapper")
}

// CoreDragCleanDragStateSymbol returns the raw symbol address for CoreDragCleanDragState.
func CoreDragCleanDragStateSymbol() uintptr {
	if _coreDragCleanDragStateSymbol == 0 {
		return 0
	}
	return _coreDragCleanDragStateSymbol
}

var _coreDragClearAllImageOverridesSymbol uintptr

// CoreDragClearAllImageOverrides has an opaque C signature in discovered metadata.
// Call CoreDragClearAllImageOverridesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragClearAllImageOverrides
func CoreDragClearAllImageOverrides() {
	panic("hiservices: symbol CoreDragClearAllImageOverrides has opaque signature; use CoreDragClearAllImageOverridesSymbol() and a typed manual wrapper")
}

// CoreDragClearAllImageOverridesSymbol returns the raw symbol address for CoreDragClearAllImageOverrides.
func CoreDragClearAllImageOverridesSymbol() uintptr {
	if _coreDragClearAllImageOverridesSymbol == 0 {
		return 0
	}
	return _coreDragClearAllImageOverridesSymbol
}

var _coreDragCopyTrashLabelSymbol uintptr

// CoreDragCopyTrashLabel has an opaque C signature in discovered metadata.
// Call CoreDragCopyTrashLabelSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragCopyTrashLabel
func CoreDragCopyTrashLabel() {
	panic("hiservices: symbol CoreDragCopyTrashLabel has opaque signature; use CoreDragCopyTrashLabelSymbol() and a typed manual wrapper")
}

// CoreDragCopyTrashLabelSymbol returns the raw symbol address for CoreDragCopyTrashLabel.
func CoreDragCopyTrashLabelSymbol() uintptr {
	if _coreDragCopyTrashLabelSymbol == 0 {
		return 0
	}
	return _coreDragCopyTrashLabelSymbol
}

var _coreDragCreateSymbol uintptr

// CoreDragCreate has an opaque C signature in discovered metadata.
// Call CoreDragCreateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragCreate
func CoreDragCreate() {
	panic("hiservices: symbol CoreDragCreate has opaque signature; use CoreDragCreateSymbol() and a typed manual wrapper")
}

// CoreDragCreateSymbol returns the raw symbol address for CoreDragCreate.
func CoreDragCreateSymbol() uintptr {
	if _coreDragCreateSymbol == 0 {
		return 0
	}
	return _coreDragCreateSymbol
}

var _coreDragCreateWithCFPasteboardSymbol uintptr

// CoreDragCreateWithCFPasteboard has an opaque C signature in discovered metadata.
// Call CoreDragCreateWithCFPasteboardSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragCreateWithCFPasteboard
func CoreDragCreateWithCFPasteboard() {
	panic("hiservices: symbol CoreDragCreateWithCFPasteboard has opaque signature; use CoreDragCreateWithCFPasteboardSymbol() and a typed manual wrapper")
}

// CoreDragCreateWithCFPasteboardSymbol returns the raw symbol address for CoreDragCreateWithCFPasteboard.
func CoreDragCreateWithCFPasteboardSymbol() uintptr {
	if _coreDragCreateWithCFPasteboardSymbol == 0 {
		return 0
	}
	return _coreDragCreateWithCFPasteboardSymbol
}

var _coreDragCreateWithPasteboardSymbol uintptr

// CoreDragCreateWithPasteboard has an opaque C signature in discovered metadata.
// Call CoreDragCreateWithPasteboardSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragCreateWithPasteboard
func CoreDragCreateWithPasteboard() {
	panic("hiservices: symbol CoreDragCreateWithPasteboard has opaque signature; use CoreDragCreateWithPasteboardSymbol() and a typed manual wrapper")
}

// CoreDragCreateWithPasteboardSymbol returns the raw symbol address for CoreDragCreateWithPasteboard.
func CoreDragCreateWithPasteboardSymbol() uintptr {
	if _coreDragCreateWithPasteboardSymbol == 0 {
		return 0
	}
	return _coreDragCreateWithPasteboardSymbol
}

var _coreDragCreateWithPasteboardRefSymbol uintptr

// CoreDragCreateWithPasteboardRef has an opaque C signature in discovered metadata.
// Call CoreDragCreateWithPasteboardRefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragCreateWithPasteboardRef
func CoreDragCreateWithPasteboardRef() {
	panic("hiservices: symbol CoreDragCreateWithPasteboardRef has opaque signature; use CoreDragCreateWithPasteboardRefSymbol() and a typed manual wrapper")
}

// CoreDragCreateWithPasteboardRefSymbol returns the raw symbol address for CoreDragCreateWithPasteboardRef.
func CoreDragCreateWithPasteboardRefSymbol() uintptr {
	if _coreDragCreateWithPasteboardRefSymbol == 0 {
		return 0
	}
	return _coreDragCreateWithPasteboardRefSymbol
}

var _coreDragDisposeSymbol uintptr

// CoreDragDispose has an opaque C signature in discovered metadata.
// Call CoreDragDisposeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragDispose
func CoreDragDispose() {
	panic("hiservices: symbol CoreDragDispose has opaque signature; use CoreDragDisposeSymbol() and a typed manual wrapper")
}

// CoreDragDisposeSymbol returns the raw symbol address for CoreDragDispose.
func CoreDragDisposeSymbol() uintptr {
	if _coreDragDisposeSymbol == 0 {
		return 0
	}
	return _coreDragDisposeSymbol
}

var _coreDragEnableSpringLoadingSymbol uintptr

// CoreDragEnableSpringLoading has an opaque C signature in discovered metadata.
// Call CoreDragEnableSpringLoadingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragEnableSpringLoading
func CoreDragEnableSpringLoading() {
	panic("hiservices: symbol CoreDragEnableSpringLoading has opaque signature; use CoreDragEnableSpringLoadingSymbol() and a typed manual wrapper")
}

// CoreDragEnableSpringLoadingSymbol returns the raw symbol address for CoreDragEnableSpringLoading.
func CoreDragEnableSpringLoadingSymbol() uintptr {
	if _coreDragEnableSpringLoadingSymbol == 0 {
		return 0
	}
	return _coreDragEnableSpringLoadingSymbol
}

var _coreDragGetAllowableActionsSymbol uintptr

// CoreDragGetAllowableActions has an opaque C signature in discovered metadata.
// Call CoreDragGetAllowableActionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetAllowableActions
func CoreDragGetAllowableActions() {
	panic("hiservices: symbol CoreDragGetAllowableActions has opaque signature; use CoreDragGetAllowableActionsSymbol() and a typed manual wrapper")
}

// CoreDragGetAllowableActionsSymbol returns the raw symbol address for CoreDragGetAllowableActions.
func CoreDragGetAllowableActionsSymbol() uintptr {
	if _coreDragGetAllowableActionsSymbol == 0 {
		return 0
	}
	return _coreDragGetAllowableActionsSymbol
}

var _coreDragGetAttributesSymbol uintptr

// CoreDragGetAttributes has an opaque C signature in discovered metadata.
// Call CoreDragGetAttributesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetAttributes
func CoreDragGetAttributes() {
	panic("hiservices: symbol CoreDragGetAttributes has opaque signature; use CoreDragGetAttributesSymbol() and a typed manual wrapper")
}

// CoreDragGetAttributesSymbol returns the raw symbol address for CoreDragGetAttributes.
func CoreDragGetAttributesSymbol() uintptr {
	if _coreDragGetAttributesSymbol == 0 {
		return 0
	}
	return _coreDragGetAttributesSymbol
}

var _coreDragGetCurrentDragSymbol uintptr

// CoreDragGetCurrentDrag has an opaque C signature in discovered metadata.
// Call CoreDragGetCurrentDragSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetCurrentDrag
func CoreDragGetCurrentDrag() {
	panic("hiservices: symbol CoreDragGetCurrentDrag has opaque signature; use CoreDragGetCurrentDragSymbol() and a typed manual wrapper")
}

// CoreDragGetCurrentDragSymbol returns the raw symbol address for CoreDragGetCurrentDrag.
func CoreDragGetCurrentDragSymbol() uintptr {
	if _coreDragGetCurrentDragSymbol == 0 {
		return 0
	}
	return _coreDragGetCurrentDragSymbol
}

var _coreDragGetDragWindowSymbol uintptr

// CoreDragGetDragWindow has an opaque C signature in discovered metadata.
// Call CoreDragGetDragWindowSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetDragWindow
func CoreDragGetDragWindow() {
	panic("hiservices: symbol CoreDragGetDragWindow has opaque signature; use CoreDragGetDragWindowSymbol() and a typed manual wrapper")
}

// CoreDragGetDragWindowSymbol returns the raw symbol address for CoreDragGetDragWindow.
func CoreDragGetDragWindowSymbol() uintptr {
	if _coreDragGetDragWindowSymbol == 0 {
		return 0
	}
	return _coreDragGetDragWindowSymbol
}

var _coreDragGetDropActionsSymbol uintptr

// CoreDragGetDropActions has an opaque C signature in discovered metadata.
// Call CoreDragGetDropActionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetDropActions
func CoreDragGetDropActions() {
	panic("hiservices: symbol CoreDragGetDropActions has opaque signature; use CoreDragGetDropActionsSymbol() and a typed manual wrapper")
}

// CoreDragGetDropActionsSymbol returns the raw symbol address for CoreDragGetDropActions.
func CoreDragGetDropActionsSymbol() uintptr {
	if _coreDragGetDropActionsSymbol == 0 {
		return 0
	}
	return _coreDragGetDropActionsSymbol
}

var _coreDragGetDropLocationSymbol uintptr

// CoreDragGetDropLocation has an opaque C signature in discovered metadata.
// Call CoreDragGetDropLocationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetDropLocation
func CoreDragGetDropLocation() {
	panic("hiservices: symbol CoreDragGetDropLocation has opaque signature; use CoreDragGetDropLocationSymbol() and a typed manual wrapper")
}

// CoreDragGetDropLocationSymbol returns the raw symbol address for CoreDragGetDropLocation.
func CoreDragGetDropLocationSymbol() uintptr {
	if _coreDragGetDropLocationSymbol == 0 {
		return 0
	}
	return _coreDragGetDropLocationSymbol
}

var _coreDragGetForceSymbol uintptr

// CoreDragGetForce has an opaque C signature in discovered metadata.
// Call CoreDragGetForceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetForce
func CoreDragGetForce() {
	panic("hiservices: symbol CoreDragGetForce has opaque signature; use CoreDragGetForceSymbol() and a typed manual wrapper")
}

// CoreDragGetForceSymbol returns the raw symbol address for CoreDragGetForce.
func CoreDragGetForceSymbol() uintptr {
	if _coreDragGetForceSymbol == 0 {
		return 0
	}
	return _coreDragGetForceSymbol
}

var _coreDragGetForceStageSymbol uintptr

// CoreDragGetForceStage has an opaque C signature in discovered metadata.
// Call CoreDragGetForceStageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetForceStage
func CoreDragGetForceStage() {
	panic("hiservices: symbol CoreDragGetForceStage has opaque signature; use CoreDragGetForceStageSymbol() and a typed manual wrapper")
}

// CoreDragGetForceStageSymbol returns the raw symbol address for CoreDragGetForceStage.
func CoreDragGetForceStageSymbol() uintptr {
	if _coreDragGetForceStageSymbol == 0 {
		return 0
	}
	return _coreDragGetForceStageSymbol
}

var _coreDragGetItemBoundsSymbol uintptr

// CoreDragGetItemBounds has an opaque C signature in discovered metadata.
// Call CoreDragGetItemBoundsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetItemBounds
func CoreDragGetItemBounds() {
	panic("hiservices: symbol CoreDragGetItemBounds has opaque signature; use CoreDragGetItemBoundsSymbol() and a typed manual wrapper")
}

// CoreDragGetItemBoundsSymbol returns the raw symbol address for CoreDragGetItemBounds.
func CoreDragGetItemBoundsSymbol() uintptr {
	if _coreDragGetItemBoundsSymbol == 0 {
		return 0
	}
	return _coreDragGetItemBoundsSymbol
}

var _coreDragGetModifiersSymbol uintptr

// CoreDragGetModifiers has an opaque C signature in discovered metadata.
// Call CoreDragGetModifiersSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetModifiers
func CoreDragGetModifiers() {
	panic("hiservices: symbol CoreDragGetModifiers has opaque signature; use CoreDragGetModifiersSymbol() and a typed manual wrapper")
}

// CoreDragGetModifiersSymbol returns the raw symbol address for CoreDragGetModifiers.
func CoreDragGetModifiersSymbol() uintptr {
	if _coreDragGetModifiersSymbol == 0 {
		return 0
	}
	return _coreDragGetModifiersSymbol
}

var _coreDragGetMouseLocationSymbol uintptr

// CoreDragGetMouseLocation has an opaque C signature in discovered metadata.
// Call CoreDragGetMouseLocationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetMouseLocation
func CoreDragGetMouseLocation() {
	panic("hiservices: symbol CoreDragGetMouseLocation has opaque signature; use CoreDragGetMouseLocationSymbol() and a typed manual wrapper")
}

// CoreDragGetMouseLocationSymbol returns the raw symbol address for CoreDragGetMouseLocation.
func CoreDragGetMouseLocationSymbol() uintptr {
	if _coreDragGetMouseLocationSymbol == 0 {
		return 0
	}
	return _coreDragGetMouseLocationSymbol
}

var _coreDragGetOriginSymbol uintptr

// CoreDragGetOrigin has an opaque C signature in discovered metadata.
// Call CoreDragGetOriginSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetOrigin
func CoreDragGetOrigin() {
	panic("hiservices: symbol CoreDragGetOrigin has opaque signature; use CoreDragGetOriginSymbol() and a typed manual wrapper")
}

// CoreDragGetOriginSymbol returns the raw symbol address for CoreDragGetOrigin.
func CoreDragGetOriginSymbol() uintptr {
	if _coreDragGetOriginSymbol == 0 {
		return 0
	}
	return _coreDragGetOriginSymbol
}

var _coreDragGetPasteboardSymbol uintptr

// CoreDragGetPasteboard has an opaque C signature in discovered metadata.
// Call CoreDragGetPasteboardSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetPasteboard
func CoreDragGetPasteboard() {
	panic("hiservices: symbol CoreDragGetPasteboard has opaque signature; use CoreDragGetPasteboardSymbol() and a typed manual wrapper")
}

// CoreDragGetPasteboardSymbol returns the raw symbol address for CoreDragGetPasteboard.
func CoreDragGetPasteboardSymbol() uintptr {
	if _coreDragGetPasteboardSymbol == 0 {
		return 0
	}
	return _coreDragGetPasteboardSymbol
}

var _coreDragGetPasteboardRefSymbol uintptr

// CoreDragGetPasteboardRef has an opaque C signature in discovered metadata.
// Call CoreDragGetPasteboardRefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetPasteboardRef
func CoreDragGetPasteboardRef() {
	panic("hiservices: symbol CoreDragGetPasteboardRef has opaque signature; use CoreDragGetPasteboardRefSymbol() and a typed manual wrapper")
}

// CoreDragGetPasteboardRefSymbol returns the raw symbol address for CoreDragGetPasteboardRef.
func CoreDragGetPasteboardRefSymbol() uintptr {
	if _coreDragGetPasteboardRefSymbol == 0 {
		return 0
	}
	return _coreDragGetPasteboardRefSymbol
}

var _coreDragGetSpringLoadingAttributesSymbol uintptr

// CoreDragGetSpringLoadingAttributes has an opaque C signature in discovered metadata.
// Call CoreDragGetSpringLoadingAttributesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetSpringLoadingAttributes
func CoreDragGetSpringLoadingAttributes() {
	panic("hiservices: symbol CoreDragGetSpringLoadingAttributes has opaque signature; use CoreDragGetSpringLoadingAttributesSymbol() and a typed manual wrapper")
}

// CoreDragGetSpringLoadingAttributesSymbol returns the raw symbol address for CoreDragGetSpringLoadingAttributes.
func CoreDragGetSpringLoadingAttributesSymbol() uintptr {
	if _coreDragGetSpringLoadingAttributesSymbol == 0 {
		return 0
	}
	return _coreDragGetSpringLoadingAttributesSymbol
}

var _coreDragGetStandardDropLocationSymbol uintptr

// CoreDragGetStandardDropLocation has an opaque C signature in discovered metadata.
// Call CoreDragGetStandardDropLocationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetStandardDropLocation
func CoreDragGetStandardDropLocation() {
	panic("hiservices: symbol CoreDragGetStandardDropLocation has opaque signature; use CoreDragGetStandardDropLocationSymbol() and a typed manual wrapper")
}

// CoreDragGetStandardDropLocationSymbol returns the raw symbol address for CoreDragGetStandardDropLocation.
func CoreDragGetStandardDropLocationSymbol() uintptr {
	if _coreDragGetStandardDropLocationSymbol == 0 {
		return 0
	}
	return _coreDragGetStandardDropLocationSymbol
}

var _coreDragGetValueForKeySymbol uintptr

// CoreDragGetValueForKey has an opaque C signature in discovered metadata.
// Call CoreDragGetValueForKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragGetValueForKey
func CoreDragGetValueForKey() {
	panic("hiservices: symbol CoreDragGetValueForKey has opaque signature; use CoreDragGetValueForKeySymbol() and a typed manual wrapper")
}

// CoreDragGetValueForKeySymbol returns the raw symbol address for CoreDragGetValueForKey.
func CoreDragGetValueForKeySymbol() uintptr {
	if _coreDragGetValueForKeySymbol == 0 {
		return 0
	}
	return _coreDragGetValueForKeySymbol
}

var _coreDragHasImageOverridesSymbol uintptr

// CoreDragHasImageOverrides has an opaque C signature in discovered metadata.
// Call CoreDragHasImageOverridesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragHasImageOverrides
func CoreDragHasImageOverrides() {
	panic("hiservices: symbol CoreDragHasImageOverrides has opaque signature; use CoreDragHasImageOverridesSymbol() and a typed manual wrapper")
}

// CoreDragHasImageOverridesSymbol returns the raw symbol address for CoreDragHasImageOverrides.
func CoreDragHasImageOverridesSymbol() uintptr {
	if _coreDragHasImageOverridesSymbol == 0 {
		return 0
	}
	return _coreDragHasImageOverridesSymbol
}

var _coreDragInstallContextReceiveMessageHandlerOnConnectionSymbol uintptr

// CoreDragInstallContextReceiveMessageHandlerOnConnection has an opaque C signature in discovered metadata.
// Call CoreDragInstallContextReceiveMessageHandlerOnConnectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragInstallContextReceiveMessageHandlerOnConnection
func CoreDragInstallContextReceiveMessageHandlerOnConnection() {
	panic("hiservices: symbol CoreDragInstallContextReceiveMessageHandlerOnConnection has opaque signature; use CoreDragInstallContextReceiveMessageHandlerOnConnectionSymbol() and a typed manual wrapper")
}

// CoreDragInstallContextReceiveMessageHandlerOnConnectionSymbol returns the raw symbol address for CoreDragInstallContextReceiveMessageHandlerOnConnection.
func CoreDragInstallContextReceiveMessageHandlerOnConnectionSymbol() uintptr {
	if _coreDragInstallContextReceiveMessageHandlerOnConnectionSymbol == 0 {
		return 0
	}
	return _coreDragInstallContextReceiveMessageHandlerOnConnectionSymbol
}

var _coreDragInstallContextTrackingHandlerOnConnectionSymbol uintptr

// CoreDragInstallContextTrackingHandlerOnConnection has an opaque C signature in discovered metadata.
// Call CoreDragInstallContextTrackingHandlerOnConnectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragInstallContextTrackingHandlerOnConnection
func CoreDragInstallContextTrackingHandlerOnConnection() {
	panic("hiservices: symbol CoreDragInstallContextTrackingHandlerOnConnection has opaque signature; use CoreDragInstallContextTrackingHandlerOnConnectionSymbol() and a typed manual wrapper")
}

// CoreDragInstallContextTrackingHandlerOnConnectionSymbol returns the raw symbol address for CoreDragInstallContextTrackingHandlerOnConnection.
func CoreDragInstallContextTrackingHandlerOnConnectionSymbol() uintptr {
	if _coreDragInstallContextTrackingHandlerOnConnectionSymbol == 0 {
		return 0
	}
	return _coreDragInstallContextTrackingHandlerOnConnectionSymbol
}

var _coreDragInstallReceiveHandlerSymbol uintptr

// CoreDragInstallReceiveHandler has an opaque C signature in discovered metadata.
// Call CoreDragInstallReceiveHandlerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragInstallReceiveHandler
func CoreDragInstallReceiveHandler() {
	panic("hiservices: symbol CoreDragInstallReceiveHandler has opaque signature; use CoreDragInstallReceiveHandlerSymbol() and a typed manual wrapper")
}

// CoreDragInstallReceiveHandlerSymbol returns the raw symbol address for CoreDragInstallReceiveHandler.
func CoreDragInstallReceiveHandlerSymbol() uintptr {
	if _coreDragInstallReceiveHandlerSymbol == 0 {
		return 0
	}
	return _coreDragInstallReceiveHandlerSymbol
}

var _coreDragInstallReceiveHandlerOnConnectionSymbol uintptr

// CoreDragInstallReceiveHandlerOnConnection has an opaque C signature in discovered metadata.
// Call CoreDragInstallReceiveHandlerOnConnectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragInstallReceiveHandlerOnConnection
func CoreDragInstallReceiveHandlerOnConnection() {
	panic("hiservices: symbol CoreDragInstallReceiveHandlerOnConnection has opaque signature; use CoreDragInstallReceiveHandlerOnConnectionSymbol() and a typed manual wrapper")
}

// CoreDragInstallReceiveHandlerOnConnectionSymbol returns the raw symbol address for CoreDragInstallReceiveHandlerOnConnection.
func CoreDragInstallReceiveHandlerOnConnectionSymbol() uintptr {
	if _coreDragInstallReceiveHandlerOnConnectionSymbol == 0 {
		return 0
	}
	return _coreDragInstallReceiveHandlerOnConnectionSymbol
}

var _coreDragInstallReceiveMessageHandlerSymbol uintptr

// CoreDragInstallReceiveMessageHandler has an opaque C signature in discovered metadata.
// Call CoreDragInstallReceiveMessageHandlerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragInstallReceiveMessageHandler
func CoreDragInstallReceiveMessageHandler() {
	panic("hiservices: symbol CoreDragInstallReceiveMessageHandler has opaque signature; use CoreDragInstallReceiveMessageHandlerSymbol() and a typed manual wrapper")
}

// CoreDragInstallReceiveMessageHandlerSymbol returns the raw symbol address for CoreDragInstallReceiveMessageHandler.
func CoreDragInstallReceiveMessageHandlerSymbol() uintptr {
	if _coreDragInstallReceiveMessageHandlerSymbol == 0 {
		return 0
	}
	return _coreDragInstallReceiveMessageHandlerSymbol
}

var _coreDragInstallReceiveMessageHandlerOnConnectionSymbol uintptr

// CoreDragInstallReceiveMessageHandlerOnConnection has an opaque C signature in discovered metadata.
// Call CoreDragInstallReceiveMessageHandlerOnConnectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragInstallReceiveMessageHandlerOnConnection
func CoreDragInstallReceiveMessageHandlerOnConnection() {
	panic("hiservices: symbol CoreDragInstallReceiveMessageHandlerOnConnection has opaque signature; use CoreDragInstallReceiveMessageHandlerOnConnectionSymbol() and a typed manual wrapper")
}

// CoreDragInstallReceiveMessageHandlerOnConnectionSymbol returns the raw symbol address for CoreDragInstallReceiveMessageHandlerOnConnection.
func CoreDragInstallReceiveMessageHandlerOnConnectionSymbol() uintptr {
	if _coreDragInstallReceiveMessageHandlerOnConnectionSymbol == 0 {
		return 0
	}
	return _coreDragInstallReceiveMessageHandlerOnConnectionSymbol
}

var _coreDragInstallTrackingHandlerSymbol uintptr

// CoreDragInstallTrackingHandler has an opaque C signature in discovered metadata.
// Call CoreDragInstallTrackingHandlerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragInstallTrackingHandler
func CoreDragInstallTrackingHandler() {
	panic("hiservices: symbol CoreDragInstallTrackingHandler has opaque signature; use CoreDragInstallTrackingHandlerSymbol() and a typed manual wrapper")
}

// CoreDragInstallTrackingHandlerSymbol returns the raw symbol address for CoreDragInstallTrackingHandler.
func CoreDragInstallTrackingHandlerSymbol() uintptr {
	if _coreDragInstallTrackingHandlerSymbol == 0 {
		return 0
	}
	return _coreDragInstallTrackingHandlerSymbol
}

var _coreDragInstallTrackingHandlerOnConnectionSymbol uintptr

// CoreDragInstallTrackingHandlerOnConnection has an opaque C signature in discovered metadata.
// Call CoreDragInstallTrackingHandlerOnConnectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragInstallTrackingHandlerOnConnection
func CoreDragInstallTrackingHandlerOnConnection() {
	panic("hiservices: symbol CoreDragInstallTrackingHandlerOnConnection has opaque signature; use CoreDragInstallTrackingHandlerOnConnectionSymbol() and a typed manual wrapper")
}

// CoreDragInstallTrackingHandlerOnConnectionSymbol returns the raw symbol address for CoreDragInstallTrackingHandlerOnConnection.
func CoreDragInstallTrackingHandlerOnConnectionSymbol() uintptr {
	if _coreDragInstallTrackingHandlerOnConnectionSymbol == 0 {
		return 0
	}
	return _coreDragInstallTrackingHandlerOnConnectionSymbol
}

var _coreDragIsTrashLabelSetSymbol uintptr

// CoreDragIsTrashLabelSet has an opaque C signature in discovered metadata.
// Call CoreDragIsTrashLabelSetSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragIsTrashLabelSet
func CoreDragIsTrashLabelSet() {
	panic("hiservices: symbol CoreDragIsTrashLabelSet has opaque signature; use CoreDragIsTrashLabelSetSymbol() and a typed manual wrapper")
}

// CoreDragIsTrashLabelSetSymbol returns the raw symbol address for CoreDragIsTrashLabelSet.
func CoreDragIsTrashLabelSetSymbol() uintptr {
	if _coreDragIsTrashLabelSetSymbol == 0 {
		return 0
	}
	return _coreDragIsTrashLabelSetSymbol
}

var _coreDragItemGetImageComponentsSymbol uintptr

// CoreDragItemGetImageComponents has an opaque C signature in discovered metadata.
// Call CoreDragItemGetImageComponentsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragItemGetImageComponents
func CoreDragItemGetImageComponents() {
	panic("hiservices: symbol CoreDragItemGetImageComponents has opaque signature; use CoreDragItemGetImageComponentsSymbol() and a typed manual wrapper")
}

// CoreDragItemGetImageComponentsSymbol returns the raw symbol address for CoreDragItemGetImageComponents.
func CoreDragItemGetImageComponentsSymbol() uintptr {
	if _coreDragItemGetImageComponentsSymbol == 0 {
		return 0
	}
	return _coreDragItemGetImageComponentsSymbol
}

var _coreDragItemGetScreenFrameSymbol uintptr

// CoreDragItemGetScreenFrame has an opaque C signature in discovered metadata.
// Call CoreDragItemGetScreenFrameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragItemGetScreenFrame
func CoreDragItemGetScreenFrame() {
	panic("hiservices: symbol CoreDragItemGetScreenFrame has opaque signature; use CoreDragItemGetScreenFrameSymbol() and a typed manual wrapper")
}

// CoreDragItemGetScreenFrameSymbol returns the raw symbol address for CoreDragItemGetScreenFrame.
func CoreDragItemGetScreenFrameSymbol() uintptr {
	if _coreDragItemGetScreenFrameSymbol == 0 {
		return 0
	}
	return _coreDragItemGetScreenFrameSymbol
}

var _coreDragItemHasAlternateSymbol uintptr

// CoreDragItemHasAlternate has an opaque C signature in discovered metadata.
// Call CoreDragItemHasAlternateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragItemHasAlternate
func CoreDragItemHasAlternate() {
	panic("hiservices: symbol CoreDragItemHasAlternate has opaque signature; use CoreDragItemHasAlternateSymbol() and a typed manual wrapper")
}

// CoreDragItemHasAlternateSymbol returns the raw symbol address for CoreDragItemHasAlternate.
func CoreDragItemHasAlternateSymbol() uintptr {
	if _coreDragItemHasAlternateSymbol == 0 {
		return 0
	}
	return _coreDragItemHasAlternateSymbol
}

var _coreDragItemSetImageComponentsSymbol uintptr

// CoreDragItemSetImageComponents has an opaque C signature in discovered metadata.
// Call CoreDragItemSetImageComponentsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragItemSetImageComponents
func CoreDragItemSetImageComponents() {
	panic("hiservices: symbol CoreDragItemSetImageComponents has opaque signature; use CoreDragItemSetImageComponentsSymbol() and a typed manual wrapper")
}

// CoreDragItemSetImageComponentsSymbol returns the raw symbol address for CoreDragItemSetImageComponents.
func CoreDragItemSetImageComponentsSymbol() uintptr {
	if _coreDragItemSetImageComponentsSymbol == 0 {
		return 0
	}
	return _coreDragItemSetImageComponentsSymbol
}

var _coreDragMessageNameSymbol uintptr

// CoreDragMessageName has an opaque C signature in discovered metadata.
// Call CoreDragMessageNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragMessageName
func CoreDragMessageName() {
	panic("hiservices: symbol CoreDragMessageName has opaque signature; use CoreDragMessageNameSymbol() and a typed manual wrapper")
}

// CoreDragMessageNameSymbol returns the raw symbol address for CoreDragMessageName.
func CoreDragMessageNameSymbol() uintptr {
	if _coreDragMessageNameSymbol == 0 {
		return 0
	}
	return _coreDragMessageNameSymbol
}

var _coreDragRefIsValidSymbol uintptr

// CoreDragRefIsValid has an opaque C signature in discovered metadata.
// Call CoreDragRefIsValidSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRefIsValid
func CoreDragRefIsValid() {
	panic("hiservices: symbol CoreDragRefIsValid has opaque signature; use CoreDragRefIsValidSymbol() and a typed manual wrapper")
}

// CoreDragRefIsValidSymbol returns the raw symbol address for CoreDragRefIsValid.
func CoreDragRefIsValidSymbol() uintptr {
	if _coreDragRefIsValidSymbol == 0 {
		return 0
	}
	return _coreDragRefIsValidSymbol
}

var _coreDragRefSetImageDataForItemSymbol uintptr

// CoreDragRefSetImageDataForItem has an opaque C signature in discovered metadata.
// Call CoreDragRefSetImageDataForItemSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRefSetImageDataForItem
func CoreDragRefSetImageDataForItem() {
	panic("hiservices: symbol CoreDragRefSetImageDataForItem has opaque signature; use CoreDragRefSetImageDataForItemSymbol() and a typed manual wrapper")
}

// CoreDragRefSetImageDataForItemSymbol returns the raw symbol address for CoreDragRefSetImageDataForItem.
func CoreDragRefSetImageDataForItemSymbol() uintptr {
	if _coreDragRefSetImageDataForItemSymbol == 0 {
		return 0
	}
	return _coreDragRefSetImageDataForItemSymbol
}

var _coreDragRegisterClientSymbol uintptr

// CoreDragRegisterClient has an opaque C signature in discovered metadata.
// Call CoreDragRegisterClientSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRegisterClient
func CoreDragRegisterClient() {
	panic("hiservices: symbol CoreDragRegisterClient has opaque signature; use CoreDragRegisterClientSymbol() and a typed manual wrapper")
}

// CoreDragRegisterClientSymbol returns the raw symbol address for CoreDragRegisterClient.
func CoreDragRegisterClientSymbol() uintptr {
	if _coreDragRegisterClientSymbol == 0 {
		return 0
	}
	return _coreDragRegisterClientSymbol
}

var _coreDragRegisterClientInModesSymbol uintptr

// CoreDragRegisterClientInModes has an opaque C signature in discovered metadata.
// Call CoreDragRegisterClientInModesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRegisterClientInModes
func CoreDragRegisterClientInModes() {
	panic("hiservices: symbol CoreDragRegisterClientInModes has opaque signature; use CoreDragRegisterClientInModesSymbol() and a typed manual wrapper")
}

// CoreDragRegisterClientInModesSymbol returns the raw symbol address for CoreDragRegisterClientInModes.
func CoreDragRegisterClientInModesSymbol() uintptr {
	if _coreDragRegisterClientInModesSymbol == 0 {
		return 0
	}
	return _coreDragRegisterClientInModesSymbol
}

var _coreDragRegisterClientWithOptionsSymbol uintptr

// CoreDragRegisterClientWithOptions has an opaque C signature in discovered metadata.
// Call CoreDragRegisterClientWithOptionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRegisterClientWithOptions
func CoreDragRegisterClientWithOptions() {
	panic("hiservices: symbol CoreDragRegisterClientWithOptions has opaque signature; use CoreDragRegisterClientWithOptionsSymbol() and a typed manual wrapper")
}

// CoreDragRegisterClientWithOptionsSymbol returns the raw symbol address for CoreDragRegisterClientWithOptions.
func CoreDragRegisterClientWithOptionsSymbol() uintptr {
	if _coreDragRegisterClientWithOptionsSymbol == 0 {
		return 0
	}
	return _coreDragRegisterClientWithOptionsSymbol
}

var _coreDragReleaseImageDataSymbol uintptr

// CoreDragReleaseImageData has an opaque C signature in discovered metadata.
// Call CoreDragReleaseImageDataSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragReleaseImageData
func CoreDragReleaseImageData() {
	panic("hiservices: symbol CoreDragReleaseImageData has opaque signature; use CoreDragReleaseImageDataSymbol() and a typed manual wrapper")
}

// CoreDragReleaseImageDataSymbol returns the raw symbol address for CoreDragReleaseImageData.
func CoreDragReleaseImageDataSymbol() uintptr {
	if _coreDragReleaseImageDataSymbol == 0 {
		return 0
	}
	return _coreDragReleaseImageDataSymbol
}

var _coreDragRemoveReceiveHandlerSymbol uintptr

// CoreDragRemoveReceiveHandler has an opaque C signature in discovered metadata.
// Call CoreDragRemoveReceiveHandlerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRemoveReceiveHandler
func CoreDragRemoveReceiveHandler() {
	panic("hiservices: symbol CoreDragRemoveReceiveHandler has opaque signature; use CoreDragRemoveReceiveHandlerSymbol() and a typed manual wrapper")
}

// CoreDragRemoveReceiveHandlerSymbol returns the raw symbol address for CoreDragRemoveReceiveHandler.
func CoreDragRemoveReceiveHandlerSymbol() uintptr {
	if _coreDragRemoveReceiveHandlerSymbol == 0 {
		return 0
	}
	return _coreDragRemoveReceiveHandlerSymbol
}

var _coreDragRemoveReceiveHandlerOnConnectionSymbol uintptr

// CoreDragRemoveReceiveHandlerOnConnection has an opaque C signature in discovered metadata.
// Call CoreDragRemoveReceiveHandlerOnConnectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRemoveReceiveHandlerOnConnection
func CoreDragRemoveReceiveHandlerOnConnection() {
	panic("hiservices: symbol CoreDragRemoveReceiveHandlerOnConnection has opaque signature; use CoreDragRemoveReceiveHandlerOnConnectionSymbol() and a typed manual wrapper")
}

// CoreDragRemoveReceiveHandlerOnConnectionSymbol returns the raw symbol address for CoreDragRemoveReceiveHandlerOnConnection.
func CoreDragRemoveReceiveHandlerOnConnectionSymbol() uintptr {
	if _coreDragRemoveReceiveHandlerOnConnectionSymbol == 0 {
		return 0
	}
	return _coreDragRemoveReceiveHandlerOnConnectionSymbol
}

var _coreDragRemoveReceiveMessageHandlerSymbol uintptr

// CoreDragRemoveReceiveMessageHandler has an opaque C signature in discovered metadata.
// Call CoreDragRemoveReceiveMessageHandlerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRemoveReceiveMessageHandler
func CoreDragRemoveReceiveMessageHandler() {
	panic("hiservices: symbol CoreDragRemoveReceiveMessageHandler has opaque signature; use CoreDragRemoveReceiveMessageHandlerSymbol() and a typed manual wrapper")
}

// CoreDragRemoveReceiveMessageHandlerSymbol returns the raw symbol address for CoreDragRemoveReceiveMessageHandler.
func CoreDragRemoveReceiveMessageHandlerSymbol() uintptr {
	if _coreDragRemoveReceiveMessageHandlerSymbol == 0 {
		return 0
	}
	return _coreDragRemoveReceiveMessageHandlerSymbol
}

var _coreDragRemoveReceiveMessageHandlerOnConnectionSymbol uintptr

// CoreDragRemoveReceiveMessageHandlerOnConnection has an opaque C signature in discovered metadata.
// Call CoreDragRemoveReceiveMessageHandlerOnConnectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRemoveReceiveMessageHandlerOnConnection
func CoreDragRemoveReceiveMessageHandlerOnConnection() {
	panic("hiservices: symbol CoreDragRemoveReceiveMessageHandlerOnConnection has opaque signature; use CoreDragRemoveReceiveMessageHandlerOnConnectionSymbol() and a typed manual wrapper")
}

// CoreDragRemoveReceiveMessageHandlerOnConnectionSymbol returns the raw symbol address for CoreDragRemoveReceiveMessageHandlerOnConnection.
func CoreDragRemoveReceiveMessageHandlerOnConnectionSymbol() uintptr {
	if _coreDragRemoveReceiveMessageHandlerOnConnectionSymbol == 0 {
		return 0
	}
	return _coreDragRemoveReceiveMessageHandlerOnConnectionSymbol
}

var _coreDragRemoveTrackingHandlerSymbol uintptr

// CoreDragRemoveTrackingHandler has an opaque C signature in discovered metadata.
// Call CoreDragRemoveTrackingHandlerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRemoveTrackingHandler
func CoreDragRemoveTrackingHandler() {
	panic("hiservices: symbol CoreDragRemoveTrackingHandler has opaque signature; use CoreDragRemoveTrackingHandlerSymbol() and a typed manual wrapper")
}

// CoreDragRemoveTrackingHandlerSymbol returns the raw symbol address for CoreDragRemoveTrackingHandler.
func CoreDragRemoveTrackingHandlerSymbol() uintptr {
	if _coreDragRemoveTrackingHandlerSymbol == 0 {
		return 0
	}
	return _coreDragRemoveTrackingHandlerSymbol
}

var _coreDragRemoveTrackingHandlerOnConnectionSymbol uintptr

// CoreDragRemoveTrackingHandlerOnConnection has an opaque C signature in discovered metadata.
// Call CoreDragRemoveTrackingHandlerOnConnectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRemoveTrackingHandlerOnConnection
func CoreDragRemoveTrackingHandlerOnConnection() {
	panic("hiservices: symbol CoreDragRemoveTrackingHandlerOnConnection has opaque signature; use CoreDragRemoveTrackingHandlerOnConnectionSymbol() and a typed manual wrapper")
}

// CoreDragRemoveTrackingHandlerOnConnectionSymbol returns the raw symbol address for CoreDragRemoveTrackingHandlerOnConnection.
func CoreDragRemoveTrackingHandlerOnConnectionSymbol() uintptr {
	if _coreDragRemoveTrackingHandlerOnConnectionSymbol == 0 {
		return 0
	}
	return _coreDragRemoveTrackingHandlerOnConnectionSymbol
}

var _coreDragRequestDragCompleteMessageSymbol uintptr

// CoreDragRequestDragCompleteMessage has an opaque C signature in discovered metadata.
// Call CoreDragRequestDragCompleteMessageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragRequestDragCompleteMessage
func CoreDragRequestDragCompleteMessage() {
	panic("hiservices: symbol CoreDragRequestDragCompleteMessage has opaque signature; use CoreDragRequestDragCompleteMessageSymbol() and a typed manual wrapper")
}

// CoreDragRequestDragCompleteMessageSymbol returns the raw symbol address for CoreDragRequestDragCompleteMessage.
func CoreDragRequestDragCompleteMessageSymbol() uintptr {
	if _coreDragRequestDragCompleteMessageSymbol == 0 {
		return 0
	}
	return _coreDragRequestDragCompleteMessageSymbol
}

var _coreDragSetAllowableActionsSymbol uintptr

// CoreDragSetAllowableActions has an opaque C signature in discovered metadata.
// Call CoreDragSetAllowableActionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetAllowableActions
func CoreDragSetAllowableActions() {
	panic("hiservices: symbol CoreDragSetAllowableActions has opaque signature; use CoreDragSetAllowableActionsSymbol() and a typed manual wrapper")
}

// CoreDragSetAllowableActionsSymbol returns the raw symbol address for CoreDragSetAllowableActions.
func CoreDragSetAllowableActionsSymbol() uintptr {
	if _coreDragSetAllowableActionsSymbol == 0 {
		return 0
	}
	return _coreDragSetAllowableActionsSymbol
}

var _coreDragSetAttributeSymbol uintptr

// CoreDragSetAttribute has an opaque C signature in discovered metadata.
// Call CoreDragSetAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetAttribute
func CoreDragSetAttribute() {
	panic("hiservices: symbol CoreDragSetAttribute has opaque signature; use CoreDragSetAttributeSymbol() and a typed manual wrapper")
}

// CoreDragSetAttributeSymbol returns the raw symbol address for CoreDragSetAttribute.
func CoreDragSetAttributeSymbol() uintptr {
	if _coreDragSetAttributeSymbol == 0 {
		return 0
	}
	return _coreDragSetAttributeSymbol
}

var _coreDragSetCGEventInputProcSymbol uintptr

// CoreDragSetCGEventInputProc has an opaque C signature in discovered metadata.
// Call CoreDragSetCGEventInputProcSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetCGEventInputProc
func CoreDragSetCGEventInputProc() {
	panic("hiservices: symbol CoreDragSetCGEventInputProc has opaque signature; use CoreDragSetCGEventInputProcSymbol() and a typed manual wrapper")
}

// CoreDragSetCGEventInputProcSymbol returns the raw symbol address for CoreDragSetCGEventInputProc.
func CoreDragSetCGEventInputProcSymbol() uintptr {
	if _coreDragSetCGEventInputProcSymbol == 0 {
		return 0
	}
	return _coreDragSetCGEventInputProcSymbol
}

var _coreDragSetCGEventProcsSymbol uintptr

// CoreDragSetCGEventProcs has an opaque C signature in discovered metadata.
// Call CoreDragSetCGEventProcsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetCGEventProcs
func CoreDragSetCGEventProcs() {
	panic("hiservices: symbol CoreDragSetCGEventProcs has opaque signature; use CoreDragSetCGEventProcsSymbol() and a typed manual wrapper")
}

// CoreDragSetCGEventProcsSymbol returns the raw symbol address for CoreDragSetCGEventProcs.
func CoreDragSetCGEventProcsSymbol() uintptr {
	if _coreDragSetCGEventProcsSymbol == 0 {
		return 0
	}
	return _coreDragSetCGEventProcsSymbol
}

var _coreDragSetCGImageSymbol uintptr

// CoreDragSetCGImage has an opaque C signature in discovered metadata.
// Call CoreDragSetCGImageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetCGImage
func CoreDragSetCGImage() {
	panic("hiservices: symbol CoreDragSetCGImage has opaque signature; use CoreDragSetCGImageSymbol() and a typed manual wrapper")
}

// CoreDragSetCGImageSymbol returns the raw symbol address for CoreDragSetCGImage.
func CoreDragSetCGImageSymbol() uintptr {
	if _coreDragSetCGImageSymbol == 0 {
		return 0
	}
	return _coreDragSetCGImageSymbol
}

var _coreDragSetCGImageWithScaleSymbol uintptr

// CoreDragSetCGImageWithScale has an opaque C signature in discovered metadata.
// Call CoreDragSetCGImageWithScaleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetCGImageWithScale
func CoreDragSetCGImageWithScale() {
	panic("hiservices: symbol CoreDragSetCGImageWithScale has opaque signature; use CoreDragSetCGImageWithScaleSymbol() and a typed manual wrapper")
}

// CoreDragSetCGImageWithScaleSymbol returns the raw symbol address for CoreDragSetCGImageWithScale.
func CoreDragSetCGImageWithScaleSymbol() uintptr {
	if _coreDragSetCGImageWithScaleSymbol == 0 {
		return 0
	}
	return _coreDragSetCGImageWithScaleSymbol
}

var _coreDragSetDestClippingRectSymbol uintptr

// CoreDragSetDestClippingRect has an opaque C signature in discovered metadata.
// Call CoreDragSetDestClippingRectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetDestClippingRect
func CoreDragSetDestClippingRect() {
	panic("hiservices: symbol CoreDragSetDestClippingRect has opaque signature; use CoreDragSetDestClippingRectSymbol() and a typed manual wrapper")
}

// CoreDragSetDestClippingRectSymbol returns the raw symbol address for CoreDragSetDestClippingRect.
func CoreDragSetDestClippingRectSymbol() uintptr {
	if _coreDragSetDestClippingRectSymbol == 0 {
		return 0
	}
	return _coreDragSetDestClippingRectSymbol
}

var _coreDragSetDragRegionSymbol uintptr

// CoreDragSetDragRegion has an opaque C signature in discovered metadata.
// Call CoreDragSetDragRegionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetDragRegion
func CoreDragSetDragRegion() {
	panic("hiservices: symbol CoreDragSetDragRegion has opaque signature; use CoreDragSetDragRegionSymbol() and a typed manual wrapper")
}

// CoreDragSetDragRegionSymbol returns the raw symbol address for CoreDragSetDragRegion.
func CoreDragSetDragRegionSymbol() uintptr {
	if _coreDragSetDragRegionSymbol == 0 {
		return 0
	}
	return _coreDragSetDragRegionSymbol
}

var _coreDragSetDragRegionWithScaleSymbol uintptr

// CoreDragSetDragRegionWithScale has an opaque C signature in discovered metadata.
// Call CoreDragSetDragRegionWithScaleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetDragRegionWithScale
func CoreDragSetDragRegionWithScale() {
	panic("hiservices: symbol CoreDragSetDragRegionWithScale has opaque signature; use CoreDragSetDragRegionWithScaleSymbol() and a typed manual wrapper")
}

// CoreDragSetDragRegionWithScaleSymbol returns the raw symbol address for CoreDragSetDragRegionWithScale.
func CoreDragSetDragRegionWithScaleSymbol() uintptr {
	if _coreDragSetDragRegionWithScaleSymbol == 0 {
		return 0
	}
	return _coreDragSetDragRegionWithScaleSymbol
}

var _coreDragSetDragWindowSymbol uintptr

// CoreDragSetDragWindow has an opaque C signature in discovered metadata.
// Call CoreDragSetDragWindowSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetDragWindow
func CoreDragSetDragWindow() {
	panic("hiservices: symbol CoreDragSetDragWindow has opaque signature; use CoreDragSetDragWindowSymbol() and a typed manual wrapper")
}

// CoreDragSetDragWindowSymbol returns the raw symbol address for CoreDragSetDragWindow.
func CoreDragSetDragWindowSymbol() uintptr {
	if _coreDragSetDragWindowSymbol == 0 {
		return 0
	}
	return _coreDragSetDragWindowSymbol
}

var _coreDragSetDropActionsSymbol uintptr

// CoreDragSetDropActions has an opaque C signature in discovered metadata.
// Call CoreDragSetDropActionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetDropActions
func CoreDragSetDropActions() {
	panic("hiservices: symbol CoreDragSetDropActions has opaque signature; use CoreDragSetDropActionsSymbol() and a typed manual wrapper")
}

// CoreDragSetDropActionsSymbol returns the raw symbol address for CoreDragSetDropActions.
func CoreDragSetDropActionsSymbol() uintptr {
	if _coreDragSetDropActionsSymbol == 0 {
		return 0
	}
	return _coreDragSetDropActionsSymbol
}

var _coreDragSetDropLocationSymbol uintptr

// CoreDragSetDropLocation has an opaque C signature in discovered metadata.
// Call CoreDragSetDropLocationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetDropLocation
func CoreDragSetDropLocation() {
	panic("hiservices: symbol CoreDragSetDropLocation has opaque signature; use CoreDragSetDropLocationSymbol() and a typed manual wrapper")
}

// CoreDragSetDropLocationSymbol returns the raw symbol address for CoreDragSetDropLocation.
func CoreDragSetDropLocationSymbol() uintptr {
	if _coreDragSetDropLocationSymbol == 0 {
		return 0
	}
	return _coreDragSetDropLocationSymbol
}

var _coreDragSetDropProcSymbol uintptr

// CoreDragSetDropProc has an opaque C signature in discovered metadata.
// Call CoreDragSetDropProcSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetDropProc
func CoreDragSetDropProc() {
	panic("hiservices: symbol CoreDragSetDropProc has opaque signature; use CoreDragSetDropProcSymbol() and a typed manual wrapper")
}

// CoreDragSetDropProcSymbol returns the raw symbol address for CoreDragSetDropProc.
func CoreDragSetDropProcSymbol() uintptr {
	if _coreDragSetDropProcSymbol == 0 {
		return 0
	}
	return _coreDragSetDropProcSymbol
}

var _coreDragSetEventProcSymbol uintptr

// CoreDragSetEventProc has an opaque C signature in discovered metadata.
// Call CoreDragSetEventProcSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetEventProc
func CoreDragSetEventProc() {
	panic("hiservices: symbol CoreDragSetEventProc has opaque signature; use CoreDragSetEventProcSymbol() and a typed manual wrapper")
}

// CoreDragSetEventProcSymbol returns the raw symbol address for CoreDragSetEventProc.
func CoreDragSetEventProcSymbol() uintptr {
	if _coreDragSetEventProcSymbol == 0 {
		return 0
	}
	return _coreDragSetEventProcSymbol
}

var _coreDragSetExtendedEventProcSymbol uintptr

// CoreDragSetExtendedEventProc has an opaque C signature in discovered metadata.
// Call CoreDragSetExtendedEventProcSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetExtendedEventProc
func CoreDragSetExtendedEventProc() {
	panic("hiservices: symbol CoreDragSetExtendedEventProc has opaque signature; use CoreDragSetExtendedEventProcSymbol() and a typed manual wrapper")
}

// CoreDragSetExtendedEventProcSymbol returns the raw symbol address for CoreDragSetExtendedEventProc.
func CoreDragSetExtendedEventProcSymbol() uintptr {
	if _coreDragSetExtendedEventProcSymbol == 0 {
		return 0
	}
	return _coreDragSetExtendedEventProcSymbol
}

var _coreDragSetImageSymbol uintptr

// CoreDragSetImage has an opaque C signature in discovered metadata.
// Call CoreDragSetImageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetImage
func CoreDragSetImage() {
	panic("hiservices: symbol CoreDragSetImage has opaque signature; use CoreDragSetImageSymbol() and a typed manual wrapper")
}

// CoreDragSetImageSymbol returns the raw symbol address for CoreDragSetImage.
func CoreDragSetImageSymbol() uintptr {
	if _coreDragSetImageSymbol == 0 {
		return 0
	}
	return _coreDragSetImageSymbol
}

var _coreDragSetImageOptionsSymbol uintptr

// CoreDragSetImageOptions has an opaque C signature in discovered metadata.
// Call CoreDragSetImageOptionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetImageOptions
func CoreDragSetImageOptions() {
	panic("hiservices: symbol CoreDragSetImageOptions has opaque signature; use CoreDragSetImageOptionsSymbol() and a typed manual wrapper")
}

// CoreDragSetImageOptionsSymbol returns the raw symbol address for CoreDragSetImageOptions.
func CoreDragSetImageOptionsSymbol() uintptr {
	if _coreDragSetImageOptionsSymbol == 0 {
		return 0
	}
	return _coreDragSetImageOptionsSymbol
}

var _coreDragSetInputProcSymbol uintptr

// CoreDragSetInputProc has an opaque C signature in discovered metadata.
// Call CoreDragSetInputProcSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetInputProc
func CoreDragSetInputProc() {
	panic("hiservices: symbol CoreDragSetInputProc has opaque signature; use CoreDragSetInputProcSymbol() and a typed manual wrapper")
}

// CoreDragSetInputProcSymbol returns the raw symbol address for CoreDragSetInputProc.
func CoreDragSetInputProcSymbol() uintptr {
	if _coreDragSetInputProcSymbol == 0 {
		return 0
	}
	return _coreDragSetInputProcSymbol
}

var _coreDragSetItemBoundsSymbol uintptr

// CoreDragSetItemBounds has an opaque C signature in discovered metadata.
// Call CoreDragSetItemBoundsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetItemBounds
func CoreDragSetItemBounds() {
	panic("hiservices: symbol CoreDragSetItemBounds has opaque signature; use CoreDragSetItemBoundsSymbol() and a typed manual wrapper")
}

// CoreDragSetItemBoundsSymbol returns the raw symbol address for CoreDragSetItemBounds.
func CoreDragSetItemBoundsSymbol() uintptr {
	if _coreDragSetItemBoundsSymbol == 0 {
		return 0
	}
	return _coreDragSetItemBoundsSymbol
}

var _coreDragSetMouseLocationSymbol uintptr

// CoreDragSetMouseLocation has an opaque C signature in discovered metadata.
// Call CoreDragSetMouseLocationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetMouseLocation
func CoreDragSetMouseLocation() {
	panic("hiservices: symbol CoreDragSetMouseLocation has opaque signature; use CoreDragSetMouseLocationSymbol() and a typed manual wrapper")
}

// CoreDragSetMouseLocationSymbol returns the raw symbol address for CoreDragSetMouseLocation.
func CoreDragSetMouseLocationSymbol() uintptr {
	if _coreDragSetMouseLocationSymbol == 0 {
		return 0
	}
	return _coreDragSetMouseLocationSymbol
}

var _coreDragSetOriginSymbol uintptr

// CoreDragSetOrigin has an opaque C signature in discovered metadata.
// Call CoreDragSetOriginSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetOrigin
func CoreDragSetOrigin() {
	panic("hiservices: symbol CoreDragSetOrigin has opaque signature; use CoreDragSetOriginSymbol() and a typed manual wrapper")
}

// CoreDragSetOriginSymbol returns the raw symbol address for CoreDragSetOrigin.
func CoreDragSetOriginSymbol() uintptr {
	if _coreDragSetOriginSymbol == 0 {
		return 0
	}
	return _coreDragSetOriginSymbol
}

var _coreDragSetRootCALayerSymbol uintptr

// CoreDragSetRootCALayer has an opaque C signature in discovered metadata.
// Call CoreDragSetRootCALayerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetRootCALayer
func CoreDragSetRootCALayer() {
	panic("hiservices: symbol CoreDragSetRootCALayer has opaque signature; use CoreDragSetRootCALayerSymbol() and a typed manual wrapper")
}

// CoreDragSetRootCALayerSymbol returns the raw symbol address for CoreDragSetRootCALayer.
func CoreDragSetRootCALayerSymbol() uintptr {
	if _coreDragSetRootCALayerSymbol == 0 {
		return 0
	}
	return _coreDragSetRootCALayerSymbol
}

var _coreDragSetSourceClippingRectSymbol uintptr

// CoreDragSetSourceClippingRect has an opaque C signature in discovered metadata.
// Call CoreDragSetSourceClippingRectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetSourceClippingRect
func CoreDragSetSourceClippingRect() {
	panic("hiservices: symbol CoreDragSetSourceClippingRect has opaque signature; use CoreDragSetSourceClippingRectSymbol() and a typed manual wrapper")
}

// CoreDragSetSourceClippingRectSymbol returns the raw symbol address for CoreDragSetSourceClippingRect.
func CoreDragSetSourceClippingRectSymbol() uintptr {
	if _coreDragSetSourceClippingRectSymbol == 0 {
		return 0
	}
	return _coreDragSetSourceClippingRectSymbol
}

var _coreDragSetStandardDropLocationSymbol uintptr

// CoreDragSetStandardDropLocation has an opaque C signature in discovered metadata.
// Call CoreDragSetStandardDropLocationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetStandardDropLocation
func CoreDragSetStandardDropLocation() {
	panic("hiservices: symbol CoreDragSetStandardDropLocation has opaque signature; use CoreDragSetStandardDropLocationSymbol() and a typed manual wrapper")
}

// CoreDragSetStandardDropLocationSymbol returns the raw symbol address for CoreDragSetStandardDropLocation.
func CoreDragSetStandardDropLocationSymbol() uintptr {
	if _coreDragSetStandardDropLocationSymbol == 0 {
		return 0
	}
	return _coreDragSetStandardDropLocationSymbol
}

var _coreDragSetTrashDrawProcSymbol uintptr

// CoreDragSetTrashDrawProc has an opaque C signature in discovered metadata.
// Call CoreDragSetTrashDrawProcSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetTrashDrawProc
func CoreDragSetTrashDrawProc() {
	panic("hiservices: symbol CoreDragSetTrashDrawProc has opaque signature; use CoreDragSetTrashDrawProcSymbol() and a typed manual wrapper")
}

// CoreDragSetTrashDrawProcSymbol returns the raw symbol address for CoreDragSetTrashDrawProc.
func CoreDragSetTrashDrawProcSymbol() uintptr {
	if _coreDragSetTrashDrawProcSymbol == 0 {
		return 0
	}
	return _coreDragSetTrashDrawProcSymbol
}

var _coreDragSetTrashLabelSymbol uintptr

// CoreDragSetTrashLabel has an opaque C signature in discovered metadata.
// Call CoreDragSetTrashLabelSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetTrashLabel
func CoreDragSetTrashLabel() {
	panic("hiservices: symbol CoreDragSetTrashLabel has opaque signature; use CoreDragSetTrashLabelSymbol() and a typed manual wrapper")
}

// CoreDragSetTrashLabelSymbol returns the raw symbol address for CoreDragSetTrashLabel.
func CoreDragSetTrashLabelSymbol() uintptr {
	if _coreDragSetTrashLabelSymbol == 0 {
		return 0
	}
	return _coreDragSetTrashLabelSymbol
}

var _coreDragSetValueForKeySymbol uintptr

// CoreDragSetValueForKey has an opaque C signature in discovered metadata.
// Call CoreDragSetValueForKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragSetValueForKey
func CoreDragSetValueForKey() {
	panic("hiservices: symbol CoreDragSetValueForKey has opaque signature; use CoreDragSetValueForKeySymbol() and a typed manual wrapper")
}

// CoreDragSetValueForKeySymbol returns the raw symbol address for CoreDragSetValueForKey.
func CoreDragSetValueForKeySymbol() uintptr {
	if _coreDragSetValueForKeySymbol == 0 {
		return 0
	}
	return _coreDragSetValueForKeySymbol
}

var _coreDragStartDraggingSymbol uintptr

// CoreDragStartDragging has an opaque C signature in discovered metadata.
// Call CoreDragStartDraggingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragStartDragging
func CoreDragStartDragging() {
	panic("hiservices: symbol CoreDragStartDragging has opaque signature; use CoreDragStartDraggingSymbol() and a typed manual wrapper")
}

// CoreDragStartDraggingSymbol returns the raw symbol address for CoreDragStartDragging.
func CoreDragStartDraggingSymbol() uintptr {
	if _coreDragStartDraggingSymbol == 0 {
		return 0
	}
	return _coreDragStartDraggingSymbol
}

var _coreDragStartDraggingAsyncSymbol uintptr

// CoreDragStartDraggingAsync has an opaque C signature in discovered metadata.
// Call CoreDragStartDraggingAsyncSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragStartDraggingAsync
func CoreDragStartDraggingAsync() {
	panic("hiservices: symbol CoreDragStartDraggingAsync has opaque signature; use CoreDragStartDraggingAsyncSymbol() and a typed manual wrapper")
}

// CoreDragStartDraggingAsyncSymbol returns the raw symbol address for CoreDragStartDraggingAsync.
func CoreDragStartDraggingAsyncSymbol() uintptr {
	if _coreDragStartDraggingAsyncSymbol == 0 {
		return 0
	}
	return _coreDragStartDraggingAsyncSymbol
}

var _coreDragUpdatesBeginSymbol uintptr

// CoreDragUpdatesBegin has an opaque C signature in discovered metadata.
// Call CoreDragUpdatesBeginSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragUpdatesBegin
func CoreDragUpdatesBegin() {
	panic("hiservices: symbol CoreDragUpdatesBegin has opaque signature; use CoreDragUpdatesBeginSymbol() and a typed manual wrapper")
}

// CoreDragUpdatesBeginSymbol returns the raw symbol address for CoreDragUpdatesBegin.
func CoreDragUpdatesBeginSymbol() uintptr {
	if _coreDragUpdatesBeginSymbol == 0 {
		return 0
	}
	return _coreDragUpdatesBeginSymbol
}

var _coreDragUpdatesCommitSymbol uintptr

// CoreDragUpdatesCommit has an opaque C signature in discovered metadata.
// Call CoreDragUpdatesCommitSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreDragUpdatesCommit
func CoreDragUpdatesCommit() {
	panic("hiservices: symbol CoreDragUpdatesCommit has opaque signature; use CoreDragUpdatesCommitSymbol() and a typed manual wrapper")
}

// CoreDragUpdatesCommitSymbol returns the raw symbol address for CoreDragUpdatesCommit.
func CoreDragUpdatesCommitSymbol() uintptr {
	if _coreDragUpdatesCommitSymbol == 0 {
		return 0
	}
	return _coreDragUpdatesCommitSymbol
}

var _coreGetDragInfoSymbol uintptr

// CoreGetDragInfo has an opaque C signature in discovered metadata.
// Call CoreGetDragInfoSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreGetDragInfo
func CoreGetDragInfo() {
	panic("hiservices: symbol CoreGetDragInfo has opaque signature; use CoreGetDragInfoSymbol() and a typed manual wrapper")
}

// CoreGetDragInfoSymbol returns the raw symbol address for CoreGetDragInfo.
func CoreGetDragInfoSymbol() uintptr {
	if _coreGetDragInfoSymbol == 0 {
		return 0
	}
	return _coreGetDragInfoSymbol
}

var _coreImagingCreateImageSymbol uintptr

// CoreImagingCreateImage has an opaque C signature in discovered metadata.
// Call CoreImagingCreateImageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreImagingCreateImage
func CoreImagingCreateImage() {
	panic("hiservices: symbol CoreImagingCreateImage has opaque signature; use CoreImagingCreateImageSymbol() and a typed manual wrapper")
}

// CoreImagingCreateImageSymbol returns the raw symbol address for CoreImagingCreateImage.
func CoreImagingCreateImageSymbol() uintptr {
	if _coreImagingCreateImageSymbol == 0 {
		return 0
	}
	return _coreImagingCreateImageSymbol
}

var _coreMenuCreateKeyEquivalentStringSymbol uintptr

// CoreMenuCreateKeyEquivalentString has an opaque C signature in discovered metadata.
// Call CoreMenuCreateKeyEquivalentStringSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreMenuCreateKeyEquivalentString
func CoreMenuCreateKeyEquivalentString() {
	panic("hiservices: symbol CoreMenuCreateKeyEquivalentString has opaque signature; use CoreMenuCreateKeyEquivalentStringSymbol() and a typed manual wrapper")
}

// CoreMenuCreateKeyEquivalentStringSymbol returns the raw symbol address for CoreMenuCreateKeyEquivalentString.
func CoreMenuCreateKeyEquivalentStringSymbol() uintptr {
	if _coreMenuCreateKeyEquivalentStringSymbol == 0 {
		return 0
	}
	return _coreMenuCreateKeyEquivalentStringSymbol
}

var _coreMenuCreateVirtualKeyStringSymbol uintptr

// CoreMenuCreateVirtualKeyString has an opaque C signature in discovered metadata.
// Call CoreMenuCreateVirtualKeyStringSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreMenuCreateVirtualKeyString
func CoreMenuCreateVirtualKeyString() {
	panic("hiservices: symbol CoreMenuCreateVirtualKeyString has opaque signature; use CoreMenuCreateVirtualKeyStringSymbol() and a typed manual wrapper")
}

// CoreMenuCreateVirtualKeyStringSymbol returns the raw symbol address for CoreMenuCreateVirtualKeyString.
func CoreMenuCreateVirtualKeyStringSymbol() uintptr {
	if _coreMenuCreateVirtualKeyStringSymbol == 0 {
		return 0
	}
	return _coreMenuCreateVirtualKeyStringSymbol
}

var _coreMenuExtraAddMenuExtraSymbol uintptr

// CoreMenuExtraAddMenuExtra has an opaque C signature in discovered metadata.
// Call CoreMenuExtraAddMenuExtraSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreMenuExtraAddMenuExtra
func CoreMenuExtraAddMenuExtra() {
	panic("hiservices: symbol CoreMenuExtraAddMenuExtra has opaque signature; use CoreMenuExtraAddMenuExtraSymbol() and a typed manual wrapper")
}

// CoreMenuExtraAddMenuExtraSymbol returns the raw symbol address for CoreMenuExtraAddMenuExtra.
func CoreMenuExtraAddMenuExtraSymbol() uintptr {
	if _coreMenuExtraAddMenuExtraSymbol == 0 {
		return 0
	}
	return _coreMenuExtraAddMenuExtraSymbol
}

var _coreMenuExtraGetMenuExtraSymbol uintptr

// CoreMenuExtraGetMenuExtra has an opaque C signature in discovered metadata.
// Call CoreMenuExtraGetMenuExtraSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreMenuExtraGetMenuExtra
func CoreMenuExtraGetMenuExtra() {
	panic("hiservices: symbol CoreMenuExtraGetMenuExtra has opaque signature; use CoreMenuExtraGetMenuExtraSymbol() and a typed manual wrapper")
}

// CoreMenuExtraGetMenuExtraSymbol returns the raw symbol address for CoreMenuExtraGetMenuExtra.
func CoreMenuExtraGetMenuExtraSymbol() uintptr {
	if _coreMenuExtraGetMenuExtraSymbol == 0 {
		return 0
	}
	return _coreMenuExtraGetMenuExtraSymbol
}

var _coreMenuExtraGetOrderSymbol uintptr

// CoreMenuExtraGetOrder has an opaque C signature in discovered metadata.
// Call CoreMenuExtraGetOrderSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreMenuExtraGetOrder
func CoreMenuExtraGetOrder() {
	panic("hiservices: symbol CoreMenuExtraGetOrder has opaque signature; use CoreMenuExtraGetOrderSymbol() and a typed manual wrapper")
}

// CoreMenuExtraGetOrderSymbol returns the raw symbol address for CoreMenuExtraGetOrder.
func CoreMenuExtraGetOrderSymbol() uintptr {
	if _coreMenuExtraGetOrderSymbol == 0 {
		return 0
	}
	return _coreMenuExtraGetOrderSymbol
}

var _coreMenuExtraInsertMenuExtraSymbol uintptr

// CoreMenuExtraInsertMenuExtra has an opaque C signature in discovered metadata.
// Call CoreMenuExtraInsertMenuExtraSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreMenuExtraInsertMenuExtra
func CoreMenuExtraInsertMenuExtra() {
	panic("hiservices: symbol CoreMenuExtraInsertMenuExtra has opaque signature; use CoreMenuExtraInsertMenuExtraSymbol() and a typed manual wrapper")
}

// CoreMenuExtraInsertMenuExtraSymbol returns the raw symbol address for CoreMenuExtraInsertMenuExtra.
func CoreMenuExtraInsertMenuExtraSymbol() uintptr {
	if _coreMenuExtraInsertMenuExtraSymbol == 0 {
		return 0
	}
	return _coreMenuExtraInsertMenuExtraSymbol
}

var _coreMenuExtraRemoveMenuExtraSymbol uintptr

// CoreMenuExtraRemoveMenuExtra has an opaque C signature in discovered metadata.
// Call CoreMenuExtraRemoveMenuExtraSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreMenuExtraRemoveMenuExtra
func CoreMenuExtraRemoveMenuExtra() {
	panic("hiservices: symbol CoreMenuExtraRemoveMenuExtra has opaque signature; use CoreMenuExtraRemoveMenuExtraSymbol() and a typed manual wrapper")
}

// CoreMenuExtraRemoveMenuExtraSymbol returns the raw symbol address for CoreMenuExtraRemoveMenuExtra.
func CoreMenuExtraRemoveMenuExtraSymbol() uintptr {
	if _coreMenuExtraRemoveMenuExtraSymbol == 0 {
		return 0
	}
	return _coreMenuExtraRemoveMenuExtraSymbol
}

var _coreMenuGetVirtualKeyMapSymbol uintptr

// CoreMenuGetVirtualKeyMap has an opaque C signature in discovered metadata.
// Call CoreMenuGetVirtualKeyMapSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CoreMenuGetVirtualKeyMap
func CoreMenuGetVirtualKeyMap() {
	panic("hiservices: symbol CoreMenuGetVirtualKeyMap has opaque signature; use CoreMenuGetVirtualKeyMapSymbol() and a typed manual wrapper")
}

// CoreMenuGetVirtualKeyMapSymbol returns the raw symbol address for CoreMenuGetVirtualKeyMap.
func CoreMenuGetVirtualKeyMapSymbol() uintptr {
	if _coreMenuGetVirtualKeyMapSymbol == 0 {
		return 0
	}
	return _coreMenuGetVirtualKeyMapSymbol
}

var _createPasteboardFlavorTypeNameSymbol uintptr

// CreatePasteboardFlavorTypeName has an opaque C signature in discovered metadata.
// Call CreatePasteboardFlavorTypeNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/CreatePasteboardFlavorTypeName
func CreatePasteboardFlavorTypeName() {
	panic("hiservices: symbol CreatePasteboardFlavorTypeName has opaque signature; use CreatePasteboardFlavorTypeNameSymbol() and a typed manual wrapper")
}

// CreatePasteboardFlavorTypeNameSymbol returns the raw symbol address for CreatePasteboardFlavorTypeName.
func CreatePasteboardFlavorTypeNameSymbol() uintptr {
	if _createPasteboardFlavorTypeNameSymbol == 0 {
		return 0
	}
	return _createPasteboardFlavorTypeNameSymbol
}

var _desktopPictureCopyDisplaySymbol uintptr

// DesktopPictureCopyDisplay has an opaque C signature in discovered metadata.
// Call DesktopPictureCopyDisplaySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/DesktopPictureCopyDisplay
func DesktopPictureCopyDisplay() {
	panic("hiservices: symbol DesktopPictureCopyDisplay has opaque signature; use DesktopPictureCopyDisplaySymbol() and a typed manual wrapper")
}

// DesktopPictureCopyDisplaySymbol returns the raw symbol address for DesktopPictureCopyDisplay.
func DesktopPictureCopyDisplaySymbol() uintptr {
	if _desktopPictureCopyDisplaySymbol == 0 {
		return 0
	}
	return _desktopPictureCopyDisplaySymbol
}

var _desktopPictureCopyDisplayForSpaceSymbol uintptr

// DesktopPictureCopyDisplayForSpace has an opaque C signature in discovered metadata.
// Call DesktopPictureCopyDisplayForSpaceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/DesktopPictureCopyDisplayForSpace
func DesktopPictureCopyDisplayForSpace() {
	panic("hiservices: symbol DesktopPictureCopyDisplayForSpace has opaque signature; use DesktopPictureCopyDisplayForSpaceSymbol() and a typed manual wrapper")
}

// DesktopPictureCopyDisplayForSpaceSymbol returns the raw symbol address for DesktopPictureCopyDisplayForSpace.
func DesktopPictureCopyDisplayForSpaceSymbol() uintptr {
	if _desktopPictureCopyDisplayForSpaceSymbol == 0 {
		return 0
	}
	return _desktopPictureCopyDisplayForSpaceSymbol
}

var _desktopPictureCopyValueSymbol uintptr

// DesktopPictureCopyValue has an opaque C signature in discovered metadata.
// Call DesktopPictureCopyValueSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/DesktopPictureCopyValue
func DesktopPictureCopyValue() {
	panic("hiservices: symbol DesktopPictureCopyValue has opaque signature; use DesktopPictureCopyValueSymbol() and a typed manual wrapper")
}

// DesktopPictureCopyValueSymbol returns the raw symbol address for DesktopPictureCopyValue.
func DesktopPictureCopyValueSymbol() uintptr {
	if _desktopPictureCopyValueSymbol == 0 {
		return 0
	}
	return _desktopPictureCopyValueSymbol
}

var _desktopPictureSetDisplaySymbol uintptr

// DesktopPictureSetDisplay has an opaque C signature in discovered metadata.
// Call DesktopPictureSetDisplaySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/DesktopPictureSetDisplay
func DesktopPictureSetDisplay() {
	panic("hiservices: symbol DesktopPictureSetDisplay has opaque signature; use DesktopPictureSetDisplaySymbol() and a typed manual wrapper")
}

// DesktopPictureSetDisplaySymbol returns the raw symbol address for DesktopPictureSetDisplay.
func DesktopPictureSetDisplaySymbol() uintptr {
	if _desktopPictureSetDisplaySymbol == 0 {
		return 0
	}
	return _desktopPictureSetDisplaySymbol
}

var _desktopPictureSetDisplayForSpaceSymbol uintptr

// DesktopPictureSetDisplayForSpace has an opaque C signature in discovered metadata.
// Call DesktopPictureSetDisplayForSpaceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/DesktopPictureSetDisplayForSpace
func DesktopPictureSetDisplayForSpace() {
	panic("hiservices: symbol DesktopPictureSetDisplayForSpace has opaque signature; use DesktopPictureSetDisplayForSpaceSymbol() and a typed manual wrapper")
}

// DesktopPictureSetDisplayForSpaceSymbol returns the raw symbol address for DesktopPictureSetDisplayForSpace.
func DesktopPictureSetDisplayForSpaceSymbol() uintptr {
	if _desktopPictureSetDisplayForSpaceSymbol == 0 {
		return 0
	}
	return _desktopPictureSetDisplayForSpaceSymbol
}

var _desktopPictureSetValueSymbol uintptr

// DesktopPictureSetValue has an opaque C signature in discovered metadata.
// Call DesktopPictureSetValueSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/DesktopPictureSetValue
func DesktopPictureSetValue() {
	panic("hiservices: symbol DesktopPictureSetValue has opaque signature; use DesktopPictureSetValueSymbol() and a typed manual wrapper")
}

// DesktopPictureSetValueSymbol returns the raw symbol address for DesktopPictureSetValue.
func DesktopPictureSetValueSymbol() uintptr {
	if _desktopPictureSetValueSymbol == 0 {
		return 0
	}
	return _desktopPictureSetValueSymbol
}

var _disposeIconActionUPPSymbol uintptr

// DisposeIconActionUPP has an opaque C signature in discovered metadata.
// Call DisposeIconActionUPPSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/DisposeIconActionUPP
func DisposeIconActionUPP() {
	panic("hiservices: symbol DisposeIconActionUPP has opaque signature; use DisposeIconActionUPPSymbol() and a typed manual wrapper")
}

// DisposeIconActionUPPSymbol returns the raw symbol address for DisposeIconActionUPP.
func DisposeIconActionUPPSymbol() uintptr {
	if _disposeIconActionUPPSymbol == 0 {
		return 0
	}
	return _disposeIconActionUPPSymbol
}

var _disposeIconGetterUPPSymbol uintptr

// DisposeIconGetterUPP has an opaque C signature in discovered metadata.
// Call DisposeIconGetterUPPSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/DisposeIconGetterUPP
func DisposeIconGetterUPP() {
	panic("hiservices: symbol DisposeIconGetterUPP has opaque signature; use DisposeIconGetterUPPSymbol() and a typed manual wrapper")
}

// DisposeIconGetterUPPSymbol returns the raw symbol address for DisposeIconGetterUPP.
func DisposeIconGetterUPPSymbol() uintptr {
	if _disposeIconGetterUPPSymbol == 0 {
		return 0
	}
	return _disposeIconGetterUPPSymbol
}

var _exitToShellSymbol uintptr

// ExitToShell has an opaque C signature in discovered metadata.
// Call ExitToShellSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ExitToShell
func ExitToShell() {
	panic("hiservices: symbol ExitToShell has opaque signature; use ExitToShellSymbol() and a typed manual wrapper")
}

// ExitToShellSymbol returns the raw symbol address for ExitToShell.
func ExitToShellSymbol() uintptr {
	if _exitToShellSymbol == 0 {
		return 0
	}
	return _exitToShellSymbol
}

var _getApplicationIsDaemonSymbol uintptr

// GetApplicationIsDaemon has an opaque C signature in discovered metadata.
// Call GetApplicationIsDaemonSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetApplicationIsDaemon
func GetApplicationIsDaemon() {
	panic("hiservices: symbol GetApplicationIsDaemon has opaque signature; use GetApplicationIsDaemonSymbol() and a typed manual wrapper")
}

// GetApplicationIsDaemonSymbol returns the raw symbol address for GetApplicationIsDaemon.
func GetApplicationIsDaemonSymbol() uintptr {
	if _getApplicationIsDaemonSymbol == 0 {
		return 0
	}
	return _getApplicationIsDaemonSymbol
}

var _getCachedFlavorNameSymbol uintptr

// GetCachedFlavorName has an opaque C signature in discovered metadata.
// Call GetCachedFlavorNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetCachedFlavorName
func GetCachedFlavorName() {
	panic("hiservices: symbol GetCachedFlavorName has opaque signature; use GetCachedFlavorNameSymbol() and a typed manual wrapper")
}

// GetCachedFlavorNameSymbol returns the raw symbol address for GetCachedFlavorName.
func GetCachedFlavorNameSymbol() uintptr {
	if _getCachedFlavorNameSymbol == 0 {
		return 0
	}
	return _getCachedFlavorNameSymbol
}

var _getCurrentProcessSymbol uintptr

// GetCurrentProcess has an opaque C signature in discovered metadata.
// Call GetCurrentProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetCurrentProcess
func GetCurrentProcess() {
	panic("hiservices: symbol GetCurrentProcess has opaque signature; use GetCurrentProcessSymbol() and a typed manual wrapper")
}

// GetCurrentProcessSymbol returns the raw symbol address for GetCurrentProcess.
func GetCurrentProcessSymbol() uintptr {
	if _getCurrentProcessSymbol == 0 {
		return 0
	}
	return _getCurrentProcessSymbol
}

var _getFrontProcessSymbol uintptr

// GetFrontProcess has an opaque C signature in discovered metadata.
// Call GetFrontProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetFrontProcess
func GetFrontProcess() {
	panic("hiservices: symbol GetFrontProcess has opaque signature; use GetFrontProcessSymbol() and a typed manual wrapper")
}

// GetFrontProcessSymbol returns the raw symbol address for GetFrontProcess.
func GetFrontProcessSymbol() uintptr {
	if _getFrontProcessSymbol == 0 {
		return 0
	}
	return _getFrontProcessSymbol
}

var _getGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol uintptr

// GetGlobalIconImagesCacheMaxEntriesAndMaxDataSize has an opaque C signature in discovered metadata.
// Call GetGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetGlobalIconImagesCacheMaxEntriesAndMaxDataSize
func GetGlobalIconImagesCacheMaxEntriesAndMaxDataSize() {
	panic("hiservices: symbol GetGlobalIconImagesCacheMaxEntriesAndMaxDataSize has opaque signature; use GetGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol() and a typed manual wrapper")
}

// GetGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol returns the raw symbol address for GetGlobalIconImagesCacheMaxEntriesAndMaxDataSize.
func GetGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol() uintptr {
	if _getGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol == 0 {
		return 0
	}
	return _getGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol
}

var _getIconFamilyDataSymbol uintptr

// GetIconFamilyData has an opaque C signature in discovered metadata.
// Call GetIconFamilyDataSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetIconFamilyData
func GetIconFamilyData() {
	panic("hiservices: symbol GetIconFamilyData has opaque signature; use GetIconFamilyDataSymbol() and a typed manual wrapper")
}

// GetIconFamilyDataSymbol returns the raw symbol address for GetIconFamilyData.
func GetIconFamilyDataSymbol() uintptr {
	if _getIconFamilyDataSymbol == 0 {
		return 0
	}
	return _getIconFamilyDataSymbol
}

var _getIconRefVariantSymbol uintptr

// GetIconRefVariant has an opaque C signature in discovered metadata.
// Call GetIconRefVariantSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetIconRefVariant
func GetIconRefVariant() {
	panic("hiservices: symbol GetIconRefVariant has opaque signature; use GetIconRefVariantSymbol() and a typed manual wrapper")
}

// GetIconRefVariantSymbol returns the raw symbol address for GetIconRefVariant.
func GetIconRefVariantSymbol() uintptr {
	if _getIconRefVariantSymbol == 0 {
		return 0
	}
	return _getIconRefVariantSymbol
}

var _getNextProcessSymbol uintptr

// GetNextProcess has an opaque C signature in discovered metadata.
// Call GetNextProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetNextProcess
func GetNextProcess() {
	panic("hiservices: symbol GetNextProcess has opaque signature; use GetNextProcessSymbol() and a typed manual wrapper")
}

// GetNextProcessSymbol returns the raw symbol address for GetNextProcess.
func GetNextProcessSymbol() uintptr {
	if _getNextProcessSymbol == 0 {
		return 0
	}
	return _getNextProcessSymbol
}

var _getProcessBundleLocationSymbol uintptr

// GetProcessBundleLocation has an opaque C signature in discovered metadata.
// Call GetProcessBundleLocationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetProcessBundleLocation
func GetProcessBundleLocation() {
	panic("hiservices: symbol GetProcessBundleLocation has opaque signature; use GetProcessBundleLocationSymbol() and a typed manual wrapper")
}

// GetProcessBundleLocationSymbol returns the raw symbol address for GetProcessBundleLocation.
func GetProcessBundleLocationSymbol() uintptr {
	if _getProcessBundleLocationSymbol == 0 {
		return 0
	}
	return _getProcessBundleLocationSymbol
}

var _getProcessForPIDSymbol uintptr

// GetProcessForPID has an opaque C signature in discovered metadata.
// Call GetProcessForPIDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetProcessForPID
func GetProcessForPID() {
	panic("hiservices: symbol GetProcessForPID has opaque signature; use GetProcessForPIDSymbol() and a typed manual wrapper")
}

// GetProcessForPIDSymbol returns the raw symbol address for GetProcessForPID.
func GetProcessForPIDSymbol() uintptr {
	if _getProcessForPIDSymbol == 0 {
		return 0
	}
	return _getProcessForPIDSymbol
}

var _getProcessInformationSymbol uintptr

// GetProcessInformation has an opaque C signature in discovered metadata.
// Call GetProcessInformationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetProcessInformation
func GetProcessInformation() {
	panic("hiservices: symbol GetProcessInformation has opaque signature; use GetProcessInformationSymbol() and a typed manual wrapper")
}

// GetProcessInformationSymbol returns the raw symbol address for GetProcessInformation.
func GetProcessInformationSymbol() uintptr {
	if _getProcessInformationSymbol == 0 {
		return 0
	}
	return _getProcessInformationSymbol
}

var _getProcessPIDSymbol uintptr

// GetProcessPID has an opaque C signature in discovered metadata.
// Call GetProcessPIDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/GetProcessPID
func GetProcessPID() {
	panic("hiservices: symbol GetProcessPID has opaque signature; use GetProcessPIDSymbol() and a typed manual wrapper")
}

// GetProcessPIDSymbol returns the raw symbol address for GetProcessPID.
func GetProcessPIDSymbol() uintptr {
	if _getProcessPIDSymbol == 0 {
		return 0
	}
	return _getProcessPIDSymbol
}

var _hIS_XPC_CFNotificationCenterPostNotificationSymbol uintptr

// HIS_XPC_CFNotificationCenterPostNotification has an opaque C signature in discovered metadata.
// Call HIS_XPC_CFNotificationCenterPostNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_CFNotificationCenterPostNotification
func HIS_XPC_CFNotificationCenterPostNotification() {
	panic("hiservices: symbol HIS_XPC_CFNotificationCenterPostNotification has opaque signature; use HIS_XPC_CFNotificationCenterPostNotificationSymbol() and a typed manual wrapper")
}

// HIS_XPC_CFNotificationCenterPostNotificationSymbol returns the raw symbol address for HIS_XPC_CFNotificationCenterPostNotification.
func HIS_XPC_CFNotificationCenterPostNotificationSymbol() uintptr {
	if _hIS_XPC_CFNotificationCenterPostNotificationSymbol == 0 {
		return 0
	}
	return _hIS_XPC_CFNotificationCenterPostNotificationSymbol
}

var _hIS_XPC_CFPreferencesCopyValueSymbol uintptr

// HIS_XPC_CFPreferencesCopyValue has an opaque C signature in discovered metadata.
// Call HIS_XPC_CFPreferencesCopyValueSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_CFPreferencesCopyValue
func HIS_XPC_CFPreferencesCopyValue() {
	panic("hiservices: symbol HIS_XPC_CFPreferencesCopyValue has opaque signature; use HIS_XPC_CFPreferencesCopyValueSymbol() and a typed manual wrapper")
}

// HIS_XPC_CFPreferencesCopyValueSymbol returns the raw symbol address for HIS_XPC_CFPreferencesCopyValue.
func HIS_XPC_CFPreferencesCopyValueSymbol() uintptr {
	if _hIS_XPC_CFPreferencesCopyValueSymbol == 0 {
		return 0
	}
	return _hIS_XPC_CFPreferencesCopyValueSymbol
}

var _hIS_XPC_CFPreferencesSetValueSymbol uintptr

// HIS_XPC_CFPreferencesSetValue has an opaque C signature in discovered metadata.
// Call HIS_XPC_CFPreferencesSetValueSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_CFPreferencesSetValue
func HIS_XPC_CFPreferencesSetValue() {
	panic("hiservices: symbol HIS_XPC_CFPreferencesSetValue has opaque signature; use HIS_XPC_CFPreferencesSetValueSymbol() and a typed manual wrapper")
}

// HIS_XPC_CFPreferencesSetValueSymbol returns the raw symbol address for HIS_XPC_CFPreferencesSetValue.
func HIS_XPC_CFPreferencesSetValueSymbol() uintptr {
	if _hIS_XPC_CFPreferencesSetValueSymbol == 0 {
		return 0
	}
	return _hIS_XPC_CFPreferencesSetValueSymbol
}

var _hIS_XPC_CFPreferencesSynchronizeSymbol uintptr

// HIS_XPC_CFPreferencesSynchronize has an opaque C signature in discovered metadata.
// Call HIS_XPC_CFPreferencesSynchronizeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_CFPreferencesSynchronize
func HIS_XPC_CFPreferencesSynchronize() {
	panic("hiservices: symbol HIS_XPC_CFPreferencesSynchronize has opaque signature; use HIS_XPC_CFPreferencesSynchronizeSymbol() and a typed manual wrapper")
}

// HIS_XPC_CFPreferencesSynchronizeSymbol returns the raw symbol address for HIS_XPC_CFPreferencesSynchronize.
func HIS_XPC_CFPreferencesSynchronizeSymbol() uintptr {
	if _hIS_XPC_CFPreferencesSynchronizeSymbol == 0 {
		return 0
	}
	return _hIS_XPC_CFPreferencesSynchronizeSymbol
}

var _hIS_XPC_CopyCapsLockKeyLabelSymbol uintptr

// HIS_XPC_CopyCapsLockKeyLabel has an opaque C signature in discovered metadata.
// Call HIS_XPC_CopyCapsLockKeyLabelSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_CopyCapsLockKeyLabel
func HIS_XPC_CopyCapsLockKeyLabel() {
	panic("hiservices: symbol HIS_XPC_CopyCapsLockKeyLabel has opaque signature; use HIS_XPC_CopyCapsLockKeyLabelSymbol() and a typed manual wrapper")
}

// HIS_XPC_CopyCapsLockKeyLabelSymbol returns the raw symbol address for HIS_XPC_CopyCapsLockKeyLabel.
func HIS_XPC_CopyCapsLockKeyLabelSymbol() uintptr {
	if _hIS_XPC_CopyCapsLockKeyLabelSymbol == 0 {
		return 0
	}
	return _hIS_XPC_CopyCapsLockKeyLabelSymbol
}

var _hIS_XPC_CopyMacManagerPrefsSymbol uintptr

// HIS_XPC_CopyMacManagerPrefs has an opaque C signature in discovered metadata.
// Call HIS_XPC_CopyMacManagerPrefsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_CopyMacManagerPrefs
func HIS_XPC_CopyMacManagerPrefs() {
	panic("hiservices: symbol HIS_XPC_CopyMacManagerPrefs has opaque signature; use HIS_XPC_CopyMacManagerPrefsSymbol() and a typed manual wrapper")
}

// HIS_XPC_CopyMacManagerPrefsSymbol returns the raw symbol address for HIS_XPC_CopyMacManagerPrefs.
func HIS_XPC_CopyMacManagerPrefsSymbol() uintptr {
	if _hIS_XPC_CopyMacManagerPrefsSymbol == 0 {
		return 0
	}
	return _hIS_XPC_CopyMacManagerPrefsSymbol
}

var _hIS_XPC_CopyPrimaryKeyboardLanguageSymbol uintptr

// HIS_XPC_CopyPrimaryKeyboardLanguage has an opaque C signature in discovered metadata.
// Call HIS_XPC_CopyPrimaryKeyboardLanguageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_CopyPrimaryKeyboardLanguage
func HIS_XPC_CopyPrimaryKeyboardLanguage() {
	panic("hiservices: symbol HIS_XPC_CopyPrimaryKeyboardLanguage has opaque signature; use HIS_XPC_CopyPrimaryKeyboardLanguageSymbol() and a typed manual wrapper")
}

// HIS_XPC_CopyPrimaryKeyboardLanguageSymbol returns the raw symbol address for HIS_XPC_CopyPrimaryKeyboardLanguage.
func HIS_XPC_CopyPrimaryKeyboardLanguageSymbol() uintptr {
	if _hIS_XPC_CopyPrimaryKeyboardLanguageSymbol == 0 {
		return 0
	}
	return _hIS_XPC_CopyPrimaryKeyboardLanguageSymbol
}

var _hIS_XPC_GetApplicationPolicyForURLsSymbol uintptr

// HIS_XPC_GetApplicationPolicyForURLs has an opaque C signature in discovered metadata.
// Call HIS_XPC_GetApplicationPolicyForURLsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_GetApplicationPolicyForURLs
func HIS_XPC_GetApplicationPolicyForURLs() {
	panic("hiservices: symbol HIS_XPC_GetApplicationPolicyForURLs has opaque signature; use HIS_XPC_GetApplicationPolicyForURLsSymbol() and a typed manual wrapper")
}

// HIS_XPC_GetApplicationPolicyForURLsSymbol returns the raw symbol address for HIS_XPC_GetApplicationPolicyForURLs.
func HIS_XPC_GetApplicationPolicyForURLsSymbol() uintptr {
	if _hIS_XPC_GetApplicationPolicyForURLsSymbol == 0 {
		return 0
	}
	return _hIS_XPC_GetApplicationPolicyForURLsSymbol
}

var _hIS_XPC_GetCapsLockLanguageSwitchSymbol uintptr

// HIS_XPC_GetCapsLockLanguageSwitch has an opaque C signature in discovered metadata.
// Call HIS_XPC_GetCapsLockLanguageSwitchSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_GetCapsLockLanguageSwitch
func HIS_XPC_GetCapsLockLanguageSwitch() {
	panic("hiservices: symbol HIS_XPC_GetCapsLockLanguageSwitch has opaque signature; use HIS_XPC_GetCapsLockLanguageSwitchSymbol() and a typed manual wrapper")
}

// HIS_XPC_GetCapsLockLanguageSwitchSymbol returns the raw symbol address for HIS_XPC_GetCapsLockLanguageSwitch.
func HIS_XPC_GetCapsLockLanguageSwitchSymbol() uintptr {
	if _hIS_XPC_GetCapsLockLanguageSwitchSymbol == 0 {
		return 0
	}
	return _hIS_XPC_GetCapsLockLanguageSwitchSymbol
}

var _hIS_XPC_GetCapsLockModifierStateSymbol uintptr

// HIS_XPC_GetCapsLockModifierState has an opaque C signature in discovered metadata.
// Call HIS_XPC_GetCapsLockModifierStateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_GetCapsLockModifierState
func HIS_XPC_GetCapsLockModifierState() {
	panic("hiservices: symbol HIS_XPC_GetCapsLockModifierState has opaque signature; use HIS_XPC_GetCapsLockModifierStateSymbol() and a typed manual wrapper")
}

// HIS_XPC_GetCapsLockModifierStateSymbol returns the raw symbol address for HIS_XPC_GetCapsLockModifierState.
func HIS_XPC_GetCapsLockModifierStateSymbol() uintptr {
	if _hIS_XPC_GetCapsLockModifierStateSymbol == 0 {
		return 0
	}
	return _hIS_XPC_GetCapsLockModifierStateSymbol
}

var _hIS_XPC_GetGlobeKeyAvailabilitySymbol uintptr

// HIS_XPC_GetGlobeKeyAvailability has an opaque C signature in discovered metadata.
// Call HIS_XPC_GetGlobeKeyAvailabilitySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_GetGlobeKeyAvailability
func HIS_XPC_GetGlobeKeyAvailability() {
	panic("hiservices: symbol HIS_XPC_GetGlobeKeyAvailability has opaque signature; use HIS_XPC_GetGlobeKeyAvailabilitySymbol() and a typed manual wrapper")
}

// HIS_XPC_GetGlobeKeyAvailabilitySymbol returns the raw symbol address for HIS_XPC_GetGlobeKeyAvailability.
func HIS_XPC_GetGlobeKeyAvailabilitySymbol() uintptr {
	if _hIS_XPC_GetGlobeKeyAvailabilitySymbol == 0 {
		return 0
	}
	return _hIS_XPC_GetGlobeKeyAvailabilitySymbol
}

var _hIS_XPC_GetMicKeyAvailabilitySymbol uintptr

// HIS_XPC_GetMicKeyAvailability has an opaque C signature in discovered metadata.
// Call HIS_XPC_GetMicKeyAvailabilitySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_GetMicKeyAvailability
func HIS_XPC_GetMicKeyAvailability() {
	panic("hiservices: symbol HIS_XPC_GetMicKeyAvailability has opaque signature; use HIS_XPC_GetMicKeyAvailabilitySymbol() and a typed manual wrapper")
}

// HIS_XPC_GetMicKeyAvailabilitySymbol returns the raw symbol address for HIS_XPC_GetMicKeyAvailability.
func HIS_XPC_GetMicKeyAvailabilitySymbol() uintptr {
	if _hIS_XPC_GetMicKeyAvailabilitySymbol == 0 {
		return 0
	}
	return _hIS_XPC_GetMicKeyAvailabilitySymbol
}

var _hIS_XPC_ResetMessageConnectionSymbol uintptr

// HIS_XPC_ResetMessageConnection has an opaque C signature in discovered metadata.
// Call HIS_XPC_ResetMessageConnectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_ResetMessageConnection
func HIS_XPC_ResetMessageConnection() {
	panic("hiservices: symbol HIS_XPC_ResetMessageConnection has opaque signature; use HIS_XPC_ResetMessageConnectionSymbol() and a typed manual wrapper")
}

// HIS_XPC_ResetMessageConnectionSymbol returns the raw symbol address for HIS_XPC_ResetMessageConnection.
func HIS_XPC_ResetMessageConnectionSymbol() uintptr {
	if _hIS_XPC_ResetMessageConnectionSymbol == 0 {
		return 0
	}
	return _hIS_XPC_ResetMessageConnectionSymbol
}

var _hIS_XPC_RevealFileInFinderSymbol uintptr

// HIS_XPC_RevealFileInFinder has an opaque C signature in discovered metadata.
// Call HIS_XPC_RevealFileInFinderSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_RevealFileInFinder
func HIS_XPC_RevealFileInFinder() {
	panic("hiservices: symbol HIS_XPC_RevealFileInFinder has opaque signature; use HIS_XPC_RevealFileInFinderSymbol() and a typed manual wrapper")
}

// HIS_XPC_RevealFileInFinderSymbol returns the raw symbol address for HIS_XPC_RevealFileInFinder.
func HIS_XPC_RevealFileInFinderSymbol() uintptr {
	if _hIS_XPC_RevealFileInFinderSymbol == 0 {
		return 0
	}
	return _hIS_XPC_RevealFileInFinderSymbol
}

var _hIS_XPC_SendAppleEventToSystemProcessSymbol uintptr

// HIS_XPC_SendAppleEventToSystemProcess has an opaque C signature in discovered metadata.
// Call HIS_XPC_SendAppleEventToSystemProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_SendAppleEventToSystemProcess
func HIS_XPC_SendAppleEventToSystemProcess() {
	panic("hiservices: symbol HIS_XPC_SendAppleEventToSystemProcess has opaque signature; use HIS_XPC_SendAppleEventToSystemProcessSymbol() and a typed manual wrapper")
}

// HIS_XPC_SendAppleEventToSystemProcessSymbol returns the raw symbol address for HIS_XPC_SendAppleEventToSystemProcess.
func HIS_XPC_SendAppleEventToSystemProcessSymbol() uintptr {
	if _hIS_XPC_SendAppleEventToSystemProcessSymbol == 0 {
		return 0
	}
	return _hIS_XPC_SendAppleEventToSystemProcessSymbol
}

var _hIS_XPC_SetCapsLockDelayOverrideSymbol uintptr

// HIS_XPC_SetCapsLockDelayOverride has an opaque C signature in discovered metadata.
// Call HIS_XPC_SetCapsLockDelayOverrideSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_SetCapsLockDelayOverride
func HIS_XPC_SetCapsLockDelayOverride() {
	panic("hiservices: symbol HIS_XPC_SetCapsLockDelayOverride has opaque signature; use HIS_XPC_SetCapsLockDelayOverrideSymbol() and a typed manual wrapper")
}

// HIS_XPC_SetCapsLockDelayOverrideSymbol returns the raw symbol address for HIS_XPC_SetCapsLockDelayOverride.
func HIS_XPC_SetCapsLockDelayOverrideSymbol() uintptr {
	if _hIS_XPC_SetCapsLockDelayOverrideSymbol == 0 {
		return 0
	}
	return _hIS_XPC_SetCapsLockDelayOverrideSymbol
}

var _hIS_XPC_SetCapsLockLEDSymbol uintptr

// HIS_XPC_SetCapsLockLED has an opaque C signature in discovered metadata.
// Call HIS_XPC_SetCapsLockLEDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_SetCapsLockLED
func HIS_XPC_SetCapsLockLED() {
	panic("hiservices: symbol HIS_XPC_SetCapsLockLED has opaque signature; use HIS_XPC_SetCapsLockLEDSymbol() and a typed manual wrapper")
}

// HIS_XPC_SetCapsLockLEDSymbol returns the raw symbol address for HIS_XPC_SetCapsLockLED.
func HIS_XPC_SetCapsLockLEDSymbol() uintptr {
	if _hIS_XPC_SetCapsLockLEDSymbol == 0 {
		return 0
	}
	return _hIS_XPC_SetCapsLockLEDSymbol
}

var _hIS_XPC_SetCapsLockLEDInhibitSymbol uintptr

// HIS_XPC_SetCapsLockLEDInhibit has an opaque C signature in discovered metadata.
// Call HIS_XPC_SetCapsLockLEDInhibitSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_SetCapsLockLEDInhibit
func HIS_XPC_SetCapsLockLEDInhibit() {
	panic("hiservices: symbol HIS_XPC_SetCapsLockLEDInhibit has opaque signature; use HIS_XPC_SetCapsLockLEDInhibitSymbol() and a typed manual wrapper")
}

// HIS_XPC_SetCapsLockLEDInhibitSymbol returns the raw symbol address for HIS_XPC_SetCapsLockLEDInhibit.
func HIS_XPC_SetCapsLockLEDInhibitSymbol() uintptr {
	if _hIS_XPC_SetCapsLockLEDInhibitSymbol == 0 {
		return 0
	}
	return _hIS_XPC_SetCapsLockLEDInhibitSymbol
}

var _hIS_XPC_SetCapsLockModifierStateSymbol uintptr

// HIS_XPC_SetCapsLockModifierState has an opaque C signature in discovered metadata.
// Call HIS_XPC_SetCapsLockModifierStateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_SetCapsLockModifierState
func HIS_XPC_SetCapsLockModifierState() {
	panic("hiservices: symbol HIS_XPC_SetCapsLockModifierState has opaque signature; use HIS_XPC_SetCapsLockModifierStateSymbol() and a typed manual wrapper")
}

// HIS_XPC_SetCapsLockModifierStateSymbol returns the raw symbol address for HIS_XPC_SetCapsLockModifierState.
func HIS_XPC_SetCapsLockModifierStateSymbol() uintptr {
	if _hIS_XPC_SetCapsLockModifierStateSymbol == 0 {
		return 0
	}
	return _hIS_XPC_SetCapsLockModifierStateSymbol
}

var _hIS_XPC_SetNetworkLocationSymbol uintptr

// HIS_XPC_SetNetworkLocation has an opaque C signature in discovered metadata.
// Call HIS_XPC_SetNetworkLocationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIS_XPC_SetNetworkLocation
func HIS_XPC_SetNetworkLocation() {
	panic("hiservices: symbol HIS_XPC_SetNetworkLocation has opaque signature; use HIS_XPC_SetNetworkLocationSymbol() and a typed manual wrapper")
}

// HIS_XPC_SetNetworkLocationSymbol returns the raw symbol address for HIS_XPC_SetNetworkLocation.
func HIS_XPC_SetNetworkLocationSymbol() uintptr {
	if _hIS_XPC_SetNetworkLocationSymbol == 0 {
		return 0
	}
	return _hIS_XPC_SetNetworkLocationSymbol
}

var _hIShapeContainsPointSymbol uintptr

// HIShapeContainsPoint has an opaque C signature in discovered metadata.
// Call HIShapeContainsPointSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeContainsPoint
func HIShapeContainsPoint() {
	panic("hiservices: symbol HIShapeContainsPoint has opaque signature; use HIShapeContainsPointSymbol() and a typed manual wrapper")
}

// HIShapeContainsPointSymbol returns the raw symbol address for HIShapeContainsPoint.
func HIShapeContainsPointSymbol() uintptr {
	if _hIShapeContainsPointSymbol == 0 {
		return 0
	}
	return _hIShapeContainsPointSymbol
}

var _hIShapeCreateCopySymbol uintptr

// HIShapeCreateCopy has an opaque C signature in discovered metadata.
// Call HIShapeCreateCopySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateCopy
func HIShapeCreateCopy() {
	panic("hiservices: symbol HIShapeCreateCopy has opaque signature; use HIShapeCreateCopySymbol() and a typed manual wrapper")
}

// HIShapeCreateCopySymbol returns the raw symbol address for HIShapeCreateCopy.
func HIShapeCreateCopySymbol() uintptr {
	if _hIShapeCreateCopySymbol == 0 {
		return 0
	}
	return _hIShapeCreateCopySymbol
}

var _hIShapeCreateCopyAsQDRgnSymbol uintptr

// HIShapeCreateCopyAsQDRgn has an opaque C signature in discovered metadata.
// Call HIShapeCreateCopyAsQDRgnSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateCopyAsQDRgn
func HIShapeCreateCopyAsQDRgn() {
	panic("hiservices: symbol HIShapeCreateCopyAsQDRgn has opaque signature; use HIShapeCreateCopyAsQDRgnSymbol() and a typed manual wrapper")
}

// HIShapeCreateCopyAsQDRgnSymbol returns the raw symbol address for HIShapeCreateCopyAsQDRgn.
func HIShapeCreateCopyAsQDRgnSymbol() uintptr {
	if _hIShapeCreateCopyAsQDRgnSymbol == 0 {
		return 0
	}
	return _hIShapeCreateCopyAsQDRgnSymbol
}

var _hIShapeCreateDifferenceSymbol uintptr

// HIShapeCreateDifference has an opaque C signature in discovered metadata.
// Call HIShapeCreateDifferenceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateDifference
func HIShapeCreateDifference() {
	panic("hiservices: symbol HIShapeCreateDifference has opaque signature; use HIShapeCreateDifferenceSymbol() and a typed manual wrapper")
}

// HIShapeCreateDifferenceSymbol returns the raw symbol address for HIShapeCreateDifference.
func HIShapeCreateDifferenceSymbol() uintptr {
	if _hIShapeCreateDifferenceSymbol == 0 {
		return 0
	}
	return _hIShapeCreateDifferenceSymbol
}

var _hIShapeCreateEmptySymbol uintptr

// HIShapeCreateEmpty has an opaque C signature in discovered metadata.
// Call HIShapeCreateEmptySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateEmpty
func HIShapeCreateEmpty() {
	panic("hiservices: symbol HIShapeCreateEmpty has opaque signature; use HIShapeCreateEmptySymbol() and a typed manual wrapper")
}

// HIShapeCreateEmptySymbol returns the raw symbol address for HIShapeCreateEmpty.
func HIShapeCreateEmptySymbol() uintptr {
	if _hIShapeCreateEmptySymbol == 0 {
		return 0
	}
	return _hIShapeCreateEmptySymbol
}

var _hIShapeCreateIntersectionSymbol uintptr

// HIShapeCreateIntersection has an opaque C signature in discovered metadata.
// Call HIShapeCreateIntersectionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateIntersection
func HIShapeCreateIntersection() {
	panic("hiservices: symbol HIShapeCreateIntersection has opaque signature; use HIShapeCreateIntersectionSymbol() and a typed manual wrapper")
}

// HIShapeCreateIntersectionSymbol returns the raw symbol address for HIShapeCreateIntersection.
func HIShapeCreateIntersectionSymbol() uintptr {
	if _hIShapeCreateIntersectionSymbol == 0 {
		return 0
	}
	return _hIShapeCreateIntersectionSymbol
}

var _hIShapeCreateMutableSymbol uintptr

// HIShapeCreateMutable has an opaque C signature in discovered metadata.
// Call HIShapeCreateMutableSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateMutable
func HIShapeCreateMutable() {
	panic("hiservices: symbol HIShapeCreateMutable has opaque signature; use HIShapeCreateMutableSymbol() and a typed manual wrapper")
}

// HIShapeCreateMutableSymbol returns the raw symbol address for HIShapeCreateMutable.
func HIShapeCreateMutableSymbol() uintptr {
	if _hIShapeCreateMutableSymbol == 0 {
		return 0
	}
	return _hIShapeCreateMutableSymbol
}

var _hIShapeCreateMutableCopySymbol uintptr

// HIShapeCreateMutableCopy has an opaque C signature in discovered metadata.
// Call HIShapeCreateMutableCopySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateMutableCopy
func HIShapeCreateMutableCopy() {
	panic("hiservices: symbol HIShapeCreateMutableCopy has opaque signature; use HIShapeCreateMutableCopySymbol() and a typed manual wrapper")
}

// HIShapeCreateMutableCopySymbol returns the raw symbol address for HIShapeCreateMutableCopy.
func HIShapeCreateMutableCopySymbol() uintptr {
	if _hIShapeCreateMutableCopySymbol == 0 {
		return 0
	}
	return _hIShapeCreateMutableCopySymbol
}

var _hIShapeCreateMutableWithRectSymbol uintptr

// HIShapeCreateMutableWithRect has an opaque C signature in discovered metadata.
// Call HIShapeCreateMutableWithRectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateMutableWithRect
func HIShapeCreateMutableWithRect() {
	panic("hiservices: symbol HIShapeCreateMutableWithRect has opaque signature; use HIShapeCreateMutableWithRectSymbol() and a typed manual wrapper")
}

// HIShapeCreateMutableWithRectSymbol returns the raw symbol address for HIShapeCreateMutableWithRect.
func HIShapeCreateMutableWithRectSymbol() uintptr {
	if _hIShapeCreateMutableWithRectSymbol == 0 {
		return 0
	}
	return _hIShapeCreateMutableWithRectSymbol
}

var _hIShapeCreateUnionSymbol uintptr

// HIShapeCreateUnion has an opaque C signature in discovered metadata.
// Call HIShapeCreateUnionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateUnion
func HIShapeCreateUnion() {
	panic("hiservices: symbol HIShapeCreateUnion has opaque signature; use HIShapeCreateUnionSymbol() and a typed manual wrapper")
}

// HIShapeCreateUnionSymbol returns the raw symbol address for HIShapeCreateUnion.
func HIShapeCreateUnionSymbol() uintptr {
	if _hIShapeCreateUnionSymbol == 0 {
		return 0
	}
	return _hIShapeCreateUnionSymbol
}

var _hIShapeCreateWithQDRgnSymbol uintptr

// HIShapeCreateWithQDRgn has an opaque C signature in discovered metadata.
// Call HIShapeCreateWithQDRgnSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateWithQDRgn
func HIShapeCreateWithQDRgn() {
	panic("hiservices: symbol HIShapeCreateWithQDRgn has opaque signature; use HIShapeCreateWithQDRgnSymbol() and a typed manual wrapper")
}

// HIShapeCreateWithQDRgnSymbol returns the raw symbol address for HIShapeCreateWithQDRgn.
func HIShapeCreateWithQDRgnSymbol() uintptr {
	if _hIShapeCreateWithQDRgnSymbol == 0 {
		return 0
	}
	return _hIShapeCreateWithQDRgnSymbol
}

var _hIShapeCreateWithRectSymbol uintptr

// HIShapeCreateWithRect has an opaque C signature in discovered metadata.
// Call HIShapeCreateWithRectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateWithRect
func HIShapeCreateWithRect() {
	panic("hiservices: symbol HIShapeCreateWithRect has opaque signature; use HIShapeCreateWithRectSymbol() and a typed manual wrapper")
}

// HIShapeCreateWithRectSymbol returns the raw symbol address for HIShapeCreateWithRect.
func HIShapeCreateWithRectSymbol() uintptr {
	if _hIShapeCreateWithRectSymbol == 0 {
		return 0
	}
	return _hIShapeCreateWithRectSymbol
}

var _hIShapeCreateXorSymbol uintptr

// HIShapeCreateXor has an opaque C signature in discovered metadata.
// Call HIShapeCreateXorSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeCreateXor
func HIShapeCreateXor() {
	panic("hiservices: symbol HIShapeCreateXor has opaque signature; use HIShapeCreateXorSymbol() and a typed manual wrapper")
}

// HIShapeCreateXorSymbol returns the raw symbol address for HIShapeCreateXor.
func HIShapeCreateXorSymbol() uintptr {
	if _hIShapeCreateXorSymbol == 0 {
		return 0
	}
	return _hIShapeCreateXorSymbol
}

var _hIShapeDifferenceSymbol uintptr

// HIShapeDifference has an opaque C signature in discovered metadata.
// Call HIShapeDifferenceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeDifference
func HIShapeDifference() {
	panic("hiservices: symbol HIShapeDifference has opaque signature; use HIShapeDifferenceSymbol() and a typed manual wrapper")
}

// HIShapeDifferenceSymbol returns the raw symbol address for HIShapeDifference.
func HIShapeDifferenceSymbol() uintptr {
	if _hIShapeDifferenceSymbol == 0 {
		return 0
	}
	return _hIShapeDifferenceSymbol
}

var _hIShapeEnumerateSymbol uintptr

// HIShapeEnumerate has an opaque C signature in discovered metadata.
// Call HIShapeEnumerateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeEnumerate
func HIShapeEnumerate() {
	panic("hiservices: symbol HIShapeEnumerate has opaque signature; use HIShapeEnumerateSymbol() and a typed manual wrapper")
}

// HIShapeEnumerateSymbol returns the raw symbol address for HIShapeEnumerate.
func HIShapeEnumerateSymbol() uintptr {
	if _hIShapeEnumerateSymbol == 0 {
		return 0
	}
	return _hIShapeEnumerateSymbol
}

var _hIShapeGetAsQDRgnSymbol uintptr

// HIShapeGetAsQDRgn has an opaque C signature in discovered metadata.
// Call HIShapeGetAsQDRgnSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeGetAsQDRgn
func HIShapeGetAsQDRgn() {
	panic("hiservices: symbol HIShapeGetAsQDRgn has opaque signature; use HIShapeGetAsQDRgnSymbol() and a typed manual wrapper")
}

// HIShapeGetAsQDRgnSymbol returns the raw symbol address for HIShapeGetAsQDRgn.
func HIShapeGetAsQDRgnSymbol() uintptr {
	if _hIShapeGetAsQDRgnSymbol == 0 {
		return 0
	}
	return _hIShapeGetAsQDRgnSymbol
}

var _hIShapeGetBoundsSymbol uintptr

// HIShapeGetBounds has an opaque C signature in discovered metadata.
// Call HIShapeGetBoundsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeGetBounds
func HIShapeGetBounds() {
	panic("hiservices: symbol HIShapeGetBounds has opaque signature; use HIShapeGetBoundsSymbol() and a typed manual wrapper")
}

// HIShapeGetBoundsSymbol returns the raw symbol address for HIShapeGetBounds.
func HIShapeGetBoundsSymbol() uintptr {
	if _hIShapeGetBoundsSymbol == 0 {
		return 0
	}
	return _hIShapeGetBoundsSymbol
}

var _hIShapeGetTypeIDSymbol uintptr

// HIShapeGetTypeID has an opaque C signature in discovered metadata.
// Call HIShapeGetTypeIDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeGetTypeID
func HIShapeGetTypeID() {
	panic("hiservices: symbol HIShapeGetTypeID has opaque signature; use HIShapeGetTypeIDSymbol() and a typed manual wrapper")
}

// HIShapeGetTypeIDSymbol returns the raw symbol address for HIShapeGetTypeID.
func HIShapeGetTypeIDSymbol() uintptr {
	if _hIShapeGetTypeIDSymbol == 0 {
		return 0
	}
	return _hIShapeGetTypeIDSymbol
}

var _hIShapeInsetSymbol uintptr

// HIShapeInset has an opaque C signature in discovered metadata.
// Call HIShapeInsetSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeInset
func HIShapeInset() {
	panic("hiservices: symbol HIShapeInset has opaque signature; use HIShapeInsetSymbol() and a typed manual wrapper")
}

// HIShapeInsetSymbol returns the raw symbol address for HIShapeInset.
func HIShapeInsetSymbol() uintptr {
	if _hIShapeInsetSymbol == 0 {
		return 0
	}
	return _hIShapeInsetSymbol
}

var _hIShapeIntersectSymbol uintptr

// HIShapeIntersect has an opaque C signature in discovered metadata.
// Call HIShapeIntersectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeIntersect
func HIShapeIntersect() {
	panic("hiservices: symbol HIShapeIntersect has opaque signature; use HIShapeIntersectSymbol() and a typed manual wrapper")
}

// HIShapeIntersectSymbol returns the raw symbol address for HIShapeIntersect.
func HIShapeIntersectSymbol() uintptr {
	if _hIShapeIntersectSymbol == 0 {
		return 0
	}
	return _hIShapeIntersectSymbol
}

var _hIShapeIntersectsRectSymbol uintptr

// HIShapeIntersectsRect has an opaque C signature in discovered metadata.
// Call HIShapeIntersectsRectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeIntersectsRect
func HIShapeIntersectsRect() {
	panic("hiservices: symbol HIShapeIntersectsRect has opaque signature; use HIShapeIntersectsRectSymbol() and a typed manual wrapper")
}

// HIShapeIntersectsRectSymbol returns the raw symbol address for HIShapeIntersectsRect.
func HIShapeIntersectsRectSymbol() uintptr {
	if _hIShapeIntersectsRectSymbol == 0 {
		return 0
	}
	return _hIShapeIntersectsRectSymbol
}

var _hIShapeIsEmptySymbol uintptr

// HIShapeIsEmpty has an opaque C signature in discovered metadata.
// Call HIShapeIsEmptySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeIsEmpty
func HIShapeIsEmpty() {
	panic("hiservices: symbol HIShapeIsEmpty has opaque signature; use HIShapeIsEmptySymbol() and a typed manual wrapper")
}

// HIShapeIsEmptySymbol returns the raw symbol address for HIShapeIsEmpty.
func HIShapeIsEmptySymbol() uintptr {
	if _hIShapeIsEmptySymbol == 0 {
		return 0
	}
	return _hIShapeIsEmptySymbol
}

var _hIShapeIsRectangularSymbol uintptr

// HIShapeIsRectangular has an opaque C signature in discovered metadata.
// Call HIShapeIsRectangularSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeIsRectangular
func HIShapeIsRectangular() {
	panic("hiservices: symbol HIShapeIsRectangular has opaque signature; use HIShapeIsRectangularSymbol() and a typed manual wrapper")
}

// HIShapeIsRectangularSymbol returns the raw symbol address for HIShapeIsRectangular.
func HIShapeIsRectangularSymbol() uintptr {
	if _hIShapeIsRectangularSymbol == 0 {
		return 0
	}
	return _hIShapeIsRectangularSymbol
}

var _hIShapeOffsetSymbol uintptr

// HIShapeOffset has an opaque C signature in discovered metadata.
// Call HIShapeOffsetSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeOffset
func HIShapeOffset() {
	panic("hiservices: symbol HIShapeOffset has opaque signature; use HIShapeOffsetSymbol() and a typed manual wrapper")
}

// HIShapeOffsetSymbol returns the raw symbol address for HIShapeOffset.
func HIShapeOffsetSymbol() uintptr {
	if _hIShapeOffsetSymbol == 0 {
		return 0
	}
	return _hIShapeOffsetSymbol
}

var _hIShapeReplacePathInCGContextSymbol uintptr

// HIShapeReplacePathInCGContext has an opaque C signature in discovered metadata.
// Call HIShapeReplacePathInCGContextSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeReplacePathInCGContext
func HIShapeReplacePathInCGContext() {
	panic("hiservices: symbol HIShapeReplacePathInCGContext has opaque signature; use HIShapeReplacePathInCGContextSymbol() and a typed manual wrapper")
}

// HIShapeReplacePathInCGContextSymbol returns the raw symbol address for HIShapeReplacePathInCGContext.
func HIShapeReplacePathInCGContextSymbol() uintptr {
	if _hIShapeReplacePathInCGContextSymbol == 0 {
		return 0
	}
	return _hIShapeReplacePathInCGContextSymbol
}

var _hIShapeSetEmptySymbol uintptr

// HIShapeSetEmpty has an opaque C signature in discovered metadata.
// Call HIShapeSetEmptySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeSetEmpty
func HIShapeSetEmpty() {
	panic("hiservices: symbol HIShapeSetEmpty has opaque signature; use HIShapeSetEmptySymbol() and a typed manual wrapper")
}

// HIShapeSetEmptySymbol returns the raw symbol address for HIShapeSetEmpty.
func HIShapeSetEmptySymbol() uintptr {
	if _hIShapeSetEmptySymbol == 0 {
		return 0
	}
	return _hIShapeSetEmptySymbol
}

var _hIShapeSetWithShapeSymbol uintptr

// HIShapeSetWithShape has an opaque C signature in discovered metadata.
// Call HIShapeSetWithShapeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeSetWithShape
func HIShapeSetWithShape() {
	panic("hiservices: symbol HIShapeSetWithShape has opaque signature; use HIShapeSetWithShapeSymbol() and a typed manual wrapper")
}

// HIShapeSetWithShapeSymbol returns the raw symbol address for HIShapeSetWithShape.
func HIShapeSetWithShapeSymbol() uintptr {
	if _hIShapeSetWithShapeSymbol == 0 {
		return 0
	}
	return _hIShapeSetWithShapeSymbol
}

var _hIShapeUnionSymbol uintptr

// HIShapeUnion has an opaque C signature in discovered metadata.
// Call HIShapeUnionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeUnion
func HIShapeUnion() {
	panic("hiservices: symbol HIShapeUnion has opaque signature; use HIShapeUnionSymbol() and a typed manual wrapper")
}

// HIShapeUnionSymbol returns the raw symbol address for HIShapeUnion.
func HIShapeUnionSymbol() uintptr {
	if _hIShapeUnionSymbol == 0 {
		return 0
	}
	return _hIShapeUnionSymbol
}

var _hIShapeUnionWithRectSymbol uintptr

// HIShapeUnionWithRect has an opaque C signature in discovered metadata.
// Call HIShapeUnionWithRectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeUnionWithRect
func HIShapeUnionWithRect() {
	panic("hiservices: symbol HIShapeUnionWithRect has opaque signature; use HIShapeUnionWithRectSymbol() and a typed manual wrapper")
}

// HIShapeUnionWithRectSymbol returns the raw symbol address for HIShapeUnionWithRect.
func HIShapeUnionWithRectSymbol() uintptr {
	if _hIShapeUnionWithRectSymbol == 0 {
		return 0
	}
	return _hIShapeUnionWithRectSymbol
}

var _hIShapeXorSymbol uintptr

// HIShapeXor has an opaque C signature in discovered metadata.
// Call HIShapeXorSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/HIShapeXor
func HIShapeXor() {
	panic("hiservices: symbol HIShapeXor has opaque signature; use HIShapeXorSymbol() and a typed manual wrapper")
}

// HIShapeXorSymbol returns the raw symbol address for HIShapeXor.
func HIShapeXorSymbol() uintptr {
	if _hIShapeXorSymbol == 0 {
		return 0
	}
	return _hIShapeXorSymbol
}

var _iCAddMapEntrySymbol uintptr

// ICAddMapEntry has an opaque C signature in discovered metadata.
// Call ICAddMapEntrySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICAddMapEntry
func ICAddMapEntry() {
	panic("hiservices: symbol ICAddMapEntry has opaque signature; use ICAddMapEntrySymbol() and a typed manual wrapper")
}

// ICAddMapEntrySymbol returns the raw symbol address for ICAddMapEntry.
func ICAddMapEntrySymbol() uintptr {
	if _iCAddMapEntrySymbol == 0 {
		return 0
	}
	return _iCAddMapEntrySymbol
}

var _iCAddProfileSymbol uintptr

// ICAddProfile has an opaque C signature in discovered metadata.
// Call ICAddProfileSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICAddProfile
func ICAddProfile() {
	panic("hiservices: symbol ICAddProfile has opaque signature; use ICAddProfileSymbol() and a typed manual wrapper")
}

// ICAddProfileSymbol returns the raw symbol address for ICAddProfile.
func ICAddProfileSymbol() uintptr {
	if _iCAddProfileSymbol == 0 {
		return 0
	}
	return _iCAddProfileSymbol
}

var _iCBeginSymbol uintptr

// ICBegin has an opaque C signature in discovered metadata.
// Call ICBeginSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICBegin
func ICBegin() {
	panic("hiservices: symbol ICBegin has opaque signature; use ICBeginSymbol() and a typed manual wrapper")
}

// ICBeginSymbol returns the raw symbol address for ICBegin.
func ICBeginSymbol() uintptr {
	if _iCBeginSymbol == 0 {
		return 0
	}
	return _iCBeginSymbol
}

var _iCCountMapEntriesSymbol uintptr

// ICCountMapEntries has an opaque C signature in discovered metadata.
// Call ICCountMapEntriesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICCountMapEntries
func ICCountMapEntries() {
	panic("hiservices: symbol ICCountMapEntries has opaque signature; use ICCountMapEntriesSymbol() and a typed manual wrapper")
}

// ICCountMapEntriesSymbol returns the raw symbol address for ICCountMapEntries.
func ICCountMapEntriesSymbol() uintptr {
	if _iCCountMapEntriesSymbol == 0 {
		return 0
	}
	return _iCCountMapEntriesSymbol
}

var _iCCountPrefSymbol uintptr

// ICCountPref has an opaque C signature in discovered metadata.
// Call ICCountPrefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICCountPref
func ICCountPref() {
	panic("hiservices: symbol ICCountPref has opaque signature; use ICCountPrefSymbol() and a typed manual wrapper")
}

// ICCountPrefSymbol returns the raw symbol address for ICCountPref.
func ICCountPrefSymbol() uintptr {
	if _iCCountPrefSymbol == 0 {
		return 0
	}
	return _iCCountPrefSymbol
}

var _iCCountProfilesSymbol uintptr

// ICCountProfiles has an opaque C signature in discovered metadata.
// Call ICCountProfilesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICCountProfiles
func ICCountProfiles() {
	panic("hiservices: symbol ICCountProfiles has opaque signature; use ICCountProfilesSymbol() and a typed manual wrapper")
}

// ICCountProfilesSymbol returns the raw symbol address for ICCountProfiles.
func ICCountProfilesSymbol() uintptr {
	if _iCCountProfilesSymbol == 0 {
		return 0
	}
	return _iCCountProfilesSymbol
}

var _iCCreateGURLEventSymbol uintptr

// ICCreateGURLEvent has an opaque C signature in discovered metadata.
// Call ICCreateGURLEventSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICCreateGURLEvent
func ICCreateGURLEvent() {
	panic("hiservices: symbol ICCreateGURLEvent has opaque signature; use ICCreateGURLEventSymbol() and a typed manual wrapper")
}

// ICCreateGURLEventSymbol returns the raw symbol address for ICCreateGURLEvent.
func ICCreateGURLEventSymbol() uintptr {
	if _iCCreateGURLEventSymbol == 0 {
		return 0
	}
	return _iCCreateGURLEventSymbol
}

var _iCDeleteMapEntrySymbol uintptr

// ICDeleteMapEntry has an opaque C signature in discovered metadata.
// Call ICDeleteMapEntrySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICDeleteMapEntry
func ICDeleteMapEntry() {
	panic("hiservices: symbol ICDeleteMapEntry has opaque signature; use ICDeleteMapEntrySymbol() and a typed manual wrapper")
}

// ICDeleteMapEntrySymbol returns the raw symbol address for ICDeleteMapEntry.
func ICDeleteMapEntrySymbol() uintptr {
	if _iCDeleteMapEntrySymbol == 0 {
		return 0
	}
	return _iCDeleteMapEntrySymbol
}

var _iCDeletePrefSymbol uintptr

// ICDeletePref has an opaque C signature in discovered metadata.
// Call ICDeletePrefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICDeletePref
func ICDeletePref() {
	panic("hiservices: symbol ICDeletePref has opaque signature; use ICDeletePrefSymbol() and a typed manual wrapper")
}

// ICDeletePrefSymbol returns the raw symbol address for ICDeletePref.
func ICDeletePrefSymbol() uintptr {
	if _iCDeletePrefSymbol == 0 {
		return 0
	}
	return _iCDeletePrefSymbol
}

var _iCDeleteProfileSymbol uintptr

// ICDeleteProfile has an opaque C signature in discovered metadata.
// Call ICDeleteProfileSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICDeleteProfile
func ICDeleteProfile() {
	panic("hiservices: symbol ICDeleteProfile has opaque signature; use ICDeleteProfileSymbol() and a typed manual wrapper")
}

// ICDeleteProfileSymbol returns the raw symbol address for ICDeleteProfile.
func ICDeleteProfileSymbol() uintptr {
	if _iCDeleteProfileSymbol == 0 {
		return 0
	}
	return _iCDeleteProfileSymbol
}

var _iCEditPreferencesSymbol uintptr

// ICEditPreferences has an opaque C signature in discovered metadata.
// Call ICEditPreferencesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICEditPreferences
func ICEditPreferences() {
	panic("hiservices: symbol ICEditPreferences has opaque signature; use ICEditPreferencesSymbol() and a typed manual wrapper")
}

// ICEditPreferencesSymbol returns the raw symbol address for ICEditPreferences.
func ICEditPreferencesSymbol() uintptr {
	if _iCEditPreferencesSymbol == 0 {
		return 0
	}
	return _iCEditPreferencesSymbol
}

var _iCEndSymbol uintptr

// ICEnd has an opaque C signature in discovered metadata.
// Call ICEndSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICEnd
func ICEnd() {
	panic("hiservices: symbol ICEnd has opaque signature; use ICEndSymbol() and a typed manual wrapper")
}

// ICEndSymbol returns the raw symbol address for ICEnd.
func ICEndSymbol() uintptr {
	if _iCEndSymbol == 0 {
		return 0
	}
	return _iCEndSymbol
}

var _iCFindPrefHandleSymbol uintptr

// ICFindPrefHandle has an opaque C signature in discovered metadata.
// Call ICFindPrefHandleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICFindPrefHandle
func ICFindPrefHandle() {
	panic("hiservices: symbol ICFindPrefHandle has opaque signature; use ICFindPrefHandleSymbol() and a typed manual wrapper")
}

// ICFindPrefHandleSymbol returns the raw symbol address for ICFindPrefHandle.
func ICFindPrefHandleSymbol() uintptr {
	if _iCFindPrefHandleSymbol == 0 {
		return 0
	}
	return _iCFindPrefHandleSymbol
}

var _iCGetConfigNameSymbol uintptr

// ICGetConfigName has an opaque C signature in discovered metadata.
// Call ICGetConfigNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetConfigName
func ICGetConfigName() {
	panic("hiservices: symbol ICGetConfigName has opaque signature; use ICGetConfigNameSymbol() and a typed manual wrapper")
}

// ICGetConfigNameSymbol returns the raw symbol address for ICGetConfigName.
func ICGetConfigNameSymbol() uintptr {
	if _iCGetConfigNameSymbol == 0 {
		return 0
	}
	return _iCGetConfigNameSymbol
}

var _iCGetCurrentProfileSymbol uintptr

// ICGetCurrentProfile has an opaque C signature in discovered metadata.
// Call ICGetCurrentProfileSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetCurrentProfile
func ICGetCurrentProfile() {
	panic("hiservices: symbol ICGetCurrentProfile has opaque signature; use ICGetCurrentProfileSymbol() and a typed manual wrapper")
}

// ICGetCurrentProfileSymbol returns the raw symbol address for ICGetCurrentProfile.
func ICGetCurrentProfileSymbol() uintptr {
	if _iCGetCurrentProfileSymbol == 0 {
		return 0
	}
	return _iCGetCurrentProfileSymbol
}

var _iCGetDefaultPrefSymbol uintptr

// ICGetDefaultPref has an opaque C signature in discovered metadata.
// Call ICGetDefaultPrefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetDefaultPref
func ICGetDefaultPref() {
	panic("hiservices: symbol ICGetDefaultPref has opaque signature; use ICGetDefaultPrefSymbol() and a typed manual wrapper")
}

// ICGetDefaultPrefSymbol returns the raw symbol address for ICGetDefaultPref.
func ICGetDefaultPrefSymbol() uintptr {
	if _iCGetDefaultPrefSymbol == 0 {
		return 0
	}
	return _iCGetDefaultPrefSymbol
}

var _iCGetIndMapEntrySymbol uintptr

// ICGetIndMapEntry has an opaque C signature in discovered metadata.
// Call ICGetIndMapEntrySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetIndMapEntry
func ICGetIndMapEntry() {
	panic("hiservices: symbol ICGetIndMapEntry has opaque signature; use ICGetIndMapEntrySymbol() and a typed manual wrapper")
}

// ICGetIndMapEntrySymbol returns the raw symbol address for ICGetIndMapEntry.
func ICGetIndMapEntrySymbol() uintptr {
	if _iCGetIndMapEntrySymbol == 0 {
		return 0
	}
	return _iCGetIndMapEntrySymbol
}

var _iCGetIndPrefSymbol uintptr

// ICGetIndPref has an opaque C signature in discovered metadata.
// Call ICGetIndPrefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetIndPref
func ICGetIndPref() {
	panic("hiservices: symbol ICGetIndPref has opaque signature; use ICGetIndPrefSymbol() and a typed manual wrapper")
}

// ICGetIndPrefSymbol returns the raw symbol address for ICGetIndPref.
func ICGetIndPrefSymbol() uintptr {
	if _iCGetIndPrefSymbol == 0 {
		return 0
	}
	return _iCGetIndPrefSymbol
}

var _iCGetIndProfileSymbol uintptr

// ICGetIndProfile has an opaque C signature in discovered metadata.
// Call ICGetIndProfileSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetIndProfile
func ICGetIndProfile() {
	panic("hiservices: symbol ICGetIndProfile has opaque signature; use ICGetIndProfileSymbol() and a typed manual wrapper")
}

// ICGetIndProfileSymbol returns the raw symbol address for ICGetIndProfile.
func ICGetIndProfileSymbol() uintptr {
	if _iCGetIndProfileSymbol == 0 {
		return 0
	}
	return _iCGetIndProfileSymbol
}

var _iCGetMapEntrySymbol uintptr

// ICGetMapEntry has an opaque C signature in discovered metadata.
// Call ICGetMapEntrySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetMapEntry
func ICGetMapEntry() {
	panic("hiservices: symbol ICGetMapEntry has opaque signature; use ICGetMapEntrySymbol() and a typed manual wrapper")
}

// ICGetMapEntrySymbol returns the raw symbol address for ICGetMapEntry.
func ICGetMapEntrySymbol() uintptr {
	if _iCGetMapEntrySymbol == 0 {
		return 0
	}
	return _iCGetMapEntrySymbol
}

var _iCGetPermSymbol uintptr

// ICGetPerm has an opaque C signature in discovered metadata.
// Call ICGetPermSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetPerm
func ICGetPerm() {
	panic("hiservices: symbol ICGetPerm has opaque signature; use ICGetPermSymbol() and a typed manual wrapper")
}

// ICGetPermSymbol returns the raw symbol address for ICGetPerm.
func ICGetPermSymbol() uintptr {
	if _iCGetPermSymbol == 0 {
		return 0
	}
	return _iCGetPermSymbol
}

var _iCGetPrefSymbol uintptr

// ICGetPref has an opaque C signature in discovered metadata.
// Call ICGetPrefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetPref
func ICGetPref() {
	panic("hiservices: symbol ICGetPref has opaque signature; use ICGetPrefSymbol() and a typed manual wrapper")
}

// ICGetPrefSymbol returns the raw symbol address for ICGetPref.
func ICGetPrefSymbol() uintptr {
	if _iCGetPrefSymbol == 0 {
		return 0
	}
	return _iCGetPrefSymbol
}

var _iCGetPrefHandleSymbol uintptr

// ICGetPrefHandle has an opaque C signature in discovered metadata.
// Call ICGetPrefHandleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetPrefHandle
func ICGetPrefHandle() {
	panic("hiservices: symbol ICGetPrefHandle has opaque signature; use ICGetPrefHandleSymbol() and a typed manual wrapper")
}

// ICGetPrefHandleSymbol returns the raw symbol address for ICGetPrefHandle.
func ICGetPrefHandleSymbol() uintptr {
	if _iCGetPrefHandleSymbol == 0 {
		return 0
	}
	return _iCGetPrefHandleSymbol
}

var _iCGetProfileNameSymbol uintptr

// ICGetProfileName has an opaque C signature in discovered metadata.
// Call ICGetProfileNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetProfileName
func ICGetProfileName() {
	panic("hiservices: symbol ICGetProfileName has opaque signature; use ICGetProfileNameSymbol() and a typed manual wrapper")
}

// ICGetProfileNameSymbol returns the raw symbol address for ICGetProfileName.
func ICGetProfileNameSymbol() uintptr {
	if _iCGetProfileNameSymbol == 0 {
		return 0
	}
	return _iCGetProfileNameSymbol
}

var _iCGetSeedSymbol uintptr

// ICGetSeed has an opaque C signature in discovered metadata.
// Call ICGetSeedSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetSeed
func ICGetSeed() {
	panic("hiservices: symbol ICGetSeed has opaque signature; use ICGetSeedSymbol() and a typed manual wrapper")
}

// ICGetSeedSymbol returns the raw symbol address for ICGetSeed.
func ICGetSeedSymbol() uintptr {
	if _iCGetSeedSymbol == 0 {
		return 0
	}
	return _iCGetSeedSymbol
}

var _iCGetVersionSymbol uintptr

// ICGetVersion has an opaque C signature in discovered metadata.
// Call ICGetVersionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICGetVersion
func ICGetVersion() {
	panic("hiservices: symbol ICGetVersion has opaque signature; use ICGetVersionSymbol() and a typed manual wrapper")
}

// ICGetVersionSymbol returns the raw symbol address for ICGetVersion.
func ICGetVersionSymbol() uintptr {
	if _iCGetVersionSymbol == 0 {
		return 0
	}
	return _iCGetVersionSymbol
}

var _iCLaunchURLSymbol uintptr

// ICLaunchURL has an opaque C signature in discovered metadata.
// Call ICLaunchURLSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICLaunchURL
func ICLaunchURL() {
	panic("hiservices: symbol ICLaunchURL has opaque signature; use ICLaunchURLSymbol() and a typed manual wrapper")
}

// ICLaunchURLSymbol returns the raw symbol address for ICLaunchURL.
func ICLaunchURLSymbol() uintptr {
	if _iCLaunchURLSymbol == 0 {
		return 0
	}
	return _iCLaunchURLSymbol
}

var _iCMapEntriesFilenameSymbol uintptr

// ICMapEntriesFilename has an opaque C signature in discovered metadata.
// Call ICMapEntriesFilenameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICMapEntriesFilename
func ICMapEntriesFilename() {
	panic("hiservices: symbol ICMapEntriesFilename has opaque signature; use ICMapEntriesFilenameSymbol() and a typed manual wrapper")
}

// ICMapEntriesFilenameSymbol returns the raw symbol address for ICMapEntriesFilename.
func ICMapEntriesFilenameSymbol() uintptr {
	if _iCMapEntriesFilenameSymbol == 0 {
		return 0
	}
	return _iCMapEntriesFilenameSymbol
}

var _iCMapEntriesTypeCreatorSymbol uintptr

// ICMapEntriesTypeCreator has an opaque C signature in discovered metadata.
// Call ICMapEntriesTypeCreatorSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICMapEntriesTypeCreator
func ICMapEntriesTypeCreator() {
	panic("hiservices: symbol ICMapEntriesTypeCreator has opaque signature; use ICMapEntriesTypeCreatorSymbol() and a typed manual wrapper")
}

// ICMapEntriesTypeCreatorSymbol returns the raw symbol address for ICMapEntriesTypeCreator.
func ICMapEntriesTypeCreatorSymbol() uintptr {
	if _iCMapEntriesTypeCreatorSymbol == 0 {
		return 0
	}
	return _iCMapEntriesTypeCreatorSymbol
}

var _iCMapFilenameSymbol uintptr

// ICMapFilename has an opaque C signature in discovered metadata.
// Call ICMapFilenameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICMapFilename
func ICMapFilename() {
	panic("hiservices: symbol ICMapFilename has opaque signature; use ICMapFilenameSymbol() and a typed manual wrapper")
}

// ICMapFilenameSymbol returns the raw symbol address for ICMapFilename.
func ICMapFilenameSymbol() uintptr {
	if _iCMapFilenameSymbol == 0 {
		return 0
	}
	return _iCMapFilenameSymbol
}

var _iCMapTypeCreatorSymbol uintptr

// ICMapTypeCreator has an opaque C signature in discovered metadata.
// Call ICMapTypeCreatorSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICMapTypeCreator
func ICMapTypeCreator() {
	panic("hiservices: symbol ICMapTypeCreator has opaque signature; use ICMapTypeCreatorSymbol() and a typed manual wrapper")
}

// ICMapTypeCreatorSymbol returns the raw symbol address for ICMapTypeCreator.
func ICMapTypeCreatorSymbol() uintptr {
	if _iCMapTypeCreatorSymbol == 0 {
		return 0
	}
	return _iCMapTypeCreatorSymbol
}

var _iCParseURLSymbol uintptr

// ICParseURL has an opaque C signature in discovered metadata.
// Call ICParseURLSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICParseURL
func ICParseURL() {
	panic("hiservices: symbol ICParseURL has opaque signature; use ICParseURLSymbol() and a typed manual wrapper")
}

// ICParseURLSymbol returns the raw symbol address for ICParseURL.
func ICParseURLSymbol() uintptr {
	if _iCParseURLSymbol == 0 {
		return 0
	}
	return _iCParseURLSymbol
}

var _iCSendGURLEventSymbol uintptr

// ICSendGURLEvent has an opaque C signature in discovered metadata.
// Call ICSendGURLEventSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICSendGURLEvent
func ICSendGURLEvent() {
	panic("hiservices: symbol ICSendGURLEvent has opaque signature; use ICSendGURLEventSymbol() and a typed manual wrapper")
}

// ICSendGURLEventSymbol returns the raw symbol address for ICSendGURLEvent.
func ICSendGURLEventSymbol() uintptr {
	if _iCSendGURLEventSymbol == 0 {
		return 0
	}
	return _iCSendGURLEventSymbol
}

var _iCSetCurrentProfileSymbol uintptr

// ICSetCurrentProfile has an opaque C signature in discovered metadata.
// Call ICSetCurrentProfileSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICSetCurrentProfile
func ICSetCurrentProfile() {
	panic("hiservices: symbol ICSetCurrentProfile has opaque signature; use ICSetCurrentProfileSymbol() and a typed manual wrapper")
}

// ICSetCurrentProfileSymbol returns the raw symbol address for ICSetCurrentProfile.
func ICSetCurrentProfileSymbol() uintptr {
	if _iCSetCurrentProfileSymbol == 0 {
		return 0
	}
	return _iCSetCurrentProfileSymbol
}

var _iCSetMapEntrySymbol uintptr

// ICSetMapEntry has an opaque C signature in discovered metadata.
// Call ICSetMapEntrySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICSetMapEntry
func ICSetMapEntry() {
	panic("hiservices: symbol ICSetMapEntry has opaque signature; use ICSetMapEntrySymbol() and a typed manual wrapper")
}

// ICSetMapEntrySymbol returns the raw symbol address for ICSetMapEntry.
func ICSetMapEntrySymbol() uintptr {
	if _iCSetMapEntrySymbol == 0 {
		return 0
	}
	return _iCSetMapEntrySymbol
}

var _iCSetPrefSymbol uintptr

// ICSetPref has an opaque C signature in discovered metadata.
// Call ICSetPrefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICSetPref
func ICSetPref() {
	panic("hiservices: symbol ICSetPref has opaque signature; use ICSetPrefSymbol() and a typed manual wrapper")
}

// ICSetPrefSymbol returns the raw symbol address for ICSetPref.
func ICSetPrefSymbol() uintptr {
	if _iCSetPrefSymbol == 0 {
		return 0
	}
	return _iCSetPrefSymbol
}

var _iCSetPrefHandleSymbol uintptr

// ICSetPrefHandle has an opaque C signature in discovered metadata.
// Call ICSetPrefHandleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICSetPrefHandle
func ICSetPrefHandle() {
	panic("hiservices: symbol ICSetPrefHandle has opaque signature; use ICSetPrefHandleSymbol() and a typed manual wrapper")
}

// ICSetPrefHandleSymbol returns the raw symbol address for ICSetPrefHandle.
func ICSetPrefHandleSymbol() uintptr {
	if _iCSetPrefHandleSymbol == 0 {
		return 0
	}
	return _iCSetPrefHandleSymbol
}

var _iCSetProfileNameSymbol uintptr

// ICSetProfileName has an opaque C signature in discovered metadata.
// Call ICSetProfileNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICSetProfileName
func ICSetProfileName() {
	panic("hiservices: symbol ICSetProfileName has opaque signature; use ICSetProfileNameSymbol() and a typed manual wrapper")
}

// ICSetProfileNameSymbol returns the raw symbol address for ICSetProfileName.
func ICSetProfileNameSymbol() uintptr {
	if _iCSetProfileNameSymbol == 0 {
		return 0
	}
	return _iCSetProfileNameSymbol
}

var _iCStartSymbol uintptr

// ICStart has an opaque C signature in discovered metadata.
// Call ICStartSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICStart
func ICStart() {
	panic("hiservices: symbol ICStart has opaque signature; use ICStartSymbol() and a typed manual wrapper")
}

// ICStartSymbol returns the raw symbol address for ICStart.
func ICStartSymbol() uintptr {
	if _iCStartSymbol == 0 {
		return 0
	}
	return _iCStartSymbol
}

var _iCStopSymbol uintptr

// ICStop has an opaque C signature in discovered metadata.
// Call ICStopSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ICStop
func ICStop() {
	panic("hiservices: symbol ICStop has opaque signature; use ICStopSymbol() and a typed manual wrapper")
}

// ICStopSymbol returns the raw symbol address for ICStop.
func ICStopSymbol() uintptr {
	if _iCStopSymbol == 0 {
		return 0
	}
	return _iCStopSymbol
}

var _iconRefContainsCGPointSymbol uintptr

// IconRefContainsCGPoint has an opaque C signature in discovered metadata.
// Call IconRefContainsCGPointSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/IconRefContainsCGPoint
func IconRefContainsCGPoint() {
	panic("hiservices: symbol IconRefContainsCGPoint has opaque signature; use IconRefContainsCGPointSymbol() and a typed manual wrapper")
}

// IconRefContainsCGPointSymbol returns the raw symbol address for IconRefContainsCGPoint.
func IconRefContainsCGPointSymbol() uintptr {
	if _iconRefContainsCGPointSymbol == 0 {
		return 0
	}
	return _iconRefContainsCGPointSymbol
}

var _iconRefIntersectsCGRectSymbol uintptr

// IconRefIntersectsCGRect has an opaque C signature in discovered metadata.
// Call IconRefIntersectsCGRectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/IconRefIntersectsCGRect
func IconRefIntersectsCGRect() {
	panic("hiservices: symbol IconRefIntersectsCGRect has opaque signature; use IconRefIntersectsCGRectSymbol() and a typed manual wrapper")
}

// IconRefIntersectsCGRectSymbol returns the raw symbol address for IconRefIntersectsCGRect.
func IconRefIntersectsCGRectSymbol() uintptr {
	if _iconRefIntersectsCGRectSymbol == 0 {
		return 0
	}
	return _iconRefIntersectsCGRectSymbol
}

var _iconRefToHIShapeSymbol uintptr

// IconRefToHIShape has an opaque C signature in discovered metadata.
// Call IconRefToHIShapeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/IconRefToHIShape
func IconRefToHIShape() {
	panic("hiservices: symbol IconRefToHIShape has opaque signature; use IconRefToHIShapeSymbol() and a typed manual wrapper")
}

// IconRefToHIShapeSymbol returns the raw symbol address for IconRefToHIShape.
func IconRefToHIShapeSymbol() uintptr {
	if _iconRefToHIShapeSymbol == 0 {
		return 0
	}
	return _iconRefToHIShapeSymbol
}

var _iconRefToIconFamilySymbol uintptr

// IconRefToIconFamily has an opaque C signature in discovered metadata.
// Call IconRefToIconFamilySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/IconRefToIconFamily
func IconRefToIconFamily() {
	panic("hiservices: symbol IconRefToIconFamily has opaque signature; use IconRefToIconFamilySymbol() and a typed manual wrapper")
}

// IconRefToIconFamilySymbol returns the raw symbol address for IconRefToIconFamily.
func IconRefToIconFamilySymbol() uintptr {
	if _iconRefToIconFamilySymbol == 0 {
		return 0
	}
	return _iconRefToIconFamilySymbol
}

var _invokeIconActionUPPSymbol uintptr

// InvokeIconActionUPP has an opaque C signature in discovered metadata.
// Call InvokeIconActionUPPSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/InvokeIconActionUPP
func InvokeIconActionUPP() {
	panic("hiservices: symbol InvokeIconActionUPP has opaque signature; use InvokeIconActionUPPSymbol() and a typed manual wrapper")
}

// InvokeIconActionUPPSymbol returns the raw symbol address for InvokeIconActionUPP.
func InvokeIconActionUPPSymbol() uintptr {
	if _invokeIconActionUPPSymbol == 0 {
		return 0
	}
	return _invokeIconActionUPPSymbol
}

var _invokeIconGetterUPPSymbol uintptr

// InvokeIconGetterUPP has an opaque C signature in discovered metadata.
// Call InvokeIconGetterUPPSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/InvokeIconGetterUPP
func InvokeIconGetterUPP() {
	panic("hiservices: symbol InvokeIconGetterUPP has opaque signature; use InvokeIconGetterUPPSymbol() and a typed manual wrapper")
}

// InvokeIconGetterUPPSymbol returns the raw symbol address for InvokeIconGetterUPP.
func InvokeIconGetterUPPSymbol() uintptr {
	if _invokeIconGetterUPPSymbol == 0 {
		return 0
	}
	return _invokeIconGetterUPPSymbol
}

var _isIconRefMaskEmptySymbol uintptr

// IsIconRefMaskEmpty has an opaque C signature in discovered metadata.
// Call IsIconRefMaskEmptySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/IsIconRefMaskEmpty
func IsIconRefMaskEmpty() {
	panic("hiservices: symbol IsIconRefMaskEmpty has opaque signature; use IsIconRefMaskEmptySymbol() and a typed manual wrapper")
}

// IsIconRefMaskEmptySymbol returns the raw symbol address for IsIconRefMaskEmpty.
func IsIconRefMaskEmptySymbol() uintptr {
	if _isIconRefMaskEmptySymbol == 0 {
		return 0
	}
	return _isIconRefMaskEmptySymbol
}

var _isProcessManagerInitializedSymbol uintptr

// IsProcessManagerInitialized has an opaque C signature in discovered metadata.
// Call IsProcessManagerInitializedSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/IsProcessManagerInitialized
func IsProcessManagerInitialized() {
	panic("hiservices: symbol IsProcessManagerInitialized has opaque signature; use IsProcessManagerInitializedSymbol() and a typed manual wrapper")
}

// IsProcessManagerInitializedSymbol returns the raw symbol address for IsProcessManagerInitialized.
func IsProcessManagerInitializedSymbol() uintptr {
	if _isProcessManagerInitializedSymbol == 0 {
		return 0
	}
	return _isProcessManagerInitializedSymbol
}

var _isProcessVisibleSymbol uintptr

// IsProcessVisible has an opaque C signature in discovered metadata.
// Call IsProcessVisibleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/IsProcessVisible
func IsProcessVisible() {
	panic("hiservices: symbol IsProcessVisible has opaque signature; use IsProcessVisibleSymbol() and a typed manual wrapper")
}

// IsProcessVisibleSymbol returns the raw symbol address for IsProcessVisible.
func IsProcessVisibleSymbol() uintptr {
	if _isProcessVisibleSymbol == 0 {
		return 0
	}
	return _isProcessVisibleSymbol
}

var _killProcessSymbol uintptr

// KillProcess has an opaque C signature in discovered metadata.
// Call KillProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/KillProcess
func KillProcess() {
	panic("hiservices: symbol KillProcess has opaque signature; use KillProcessSymbol() and a typed manual wrapper")
}

// KillProcessSymbol returns the raw symbol address for KillProcess.
func KillProcessSymbol() uintptr {
	if _killProcessSymbol == 0 {
		return 0
	}
	return _killProcessSymbol
}

var _launchApplicationSymbol uintptr

// LaunchApplication has an opaque C signature in discovered metadata.
// Call LaunchApplicationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/LaunchApplication
func LaunchApplication() {
	panic("hiservices: symbol LaunchApplication has opaque signature; use LaunchApplicationSymbol() and a typed manual wrapper")
}

// LaunchApplicationSymbol returns the raw symbol address for LaunchApplication.
func LaunchApplicationSymbol() uintptr {
	if _launchApplicationSymbol == 0 {
		return 0
	}
	return _launchApplicationSymbol
}

var _launchProcessSymbol uintptr

// LaunchProcess has an opaque C signature in discovered metadata.
// Call LaunchProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/LaunchProcess
func LaunchProcess() {
	panic("hiservices: symbol LaunchProcess has opaque signature; use LaunchProcessSymbol() and a typed manual wrapper")
}

// LaunchProcessSymbol returns the raw symbol address for LaunchProcess.
func LaunchProcessSymbol() uintptr {
	if _launchProcessSymbol == 0 {
		return 0
	}
	return _launchProcessSymbol
}

var _launchProcessAsyncSymbol uintptr

// LaunchProcessAsync has an opaque C signature in discovered metadata.
// Call LaunchProcessAsyncSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/LaunchProcessAsync
func LaunchProcessAsync() {
	panic("hiservices: symbol LaunchProcessAsync has opaque signature; use LaunchProcessAsyncSymbol() and a typed manual wrapper")
}

// LaunchProcessAsyncSymbol returns the raw symbol address for LaunchProcessAsync.
func LaunchProcessAsyncSymbol() uintptr {
	if _launchProcessAsyncSymbol == 0 {
		return 0
	}
	return _launchProcessAsyncSymbol
}

var _mSHCreateMIGServerSourceSymbol uintptr

// MSHCreateMIGServerSource has an opaque C signature in discovered metadata.
// Call MSHCreateMIGServerSourceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/MSHCreateMIGServerSource
func MSHCreateMIGServerSource() {
	panic("hiservices: symbol MSHCreateMIGServerSource has opaque signature; use MSHCreateMIGServerSourceSymbol() and a typed manual wrapper")
}

// MSHCreateMIGServerSourceSymbol returns the raw symbol address for MSHCreateMIGServerSource.
func MSHCreateMIGServerSourceSymbol() uintptr {
	if _mSHCreateMIGServerSourceSymbol == 0 {
		return 0
	}
	return _mSHCreateMIGServerSourceSymbol
}

var _mSHCreateMachServerSourceSymbol uintptr

// MSHCreateMachServerSource has an opaque C signature in discovered metadata.
// Call MSHCreateMachServerSourceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/MSHCreateMachServerSource
func MSHCreateMachServerSource() {
	panic("hiservices: symbol MSHCreateMachServerSource has opaque signature; use MSHCreateMachServerSourceSymbol() and a typed manual wrapper")
}

// MSHCreateMachServerSourceSymbol returns the raw symbol address for MSHCreateMachServerSource.
func MSHCreateMachServerSourceSymbol() uintptr {
	if _mSHCreateMachServerSourceSymbol == 0 {
		return 0
	}
	return _mSHCreateMachServerSourceSymbol
}

var _mSHGetMachPortFromSourceSymbol uintptr

// MSHGetMachPortFromSource has an opaque C signature in discovered metadata.
// Call MSHGetMachPortFromSourceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/MSHGetMachPortFromSource
func MSHGetMachPortFromSource() {
	panic("hiservices: symbol MSHGetMachPortFromSource has opaque signature; use MSHGetMachPortFromSourceSymbol() and a typed manual wrapper")
}

// MSHGetMachPortFromSourceSymbol returns the raw symbol address for MSHGetMachPortFromSource.
func MSHGetMachPortFromSourceSymbol() uintptr {
	if _mSHGetMachPortFromSourceSymbol == 0 {
		return 0
	}
	return _mSHGetMachPortFromSourceSymbol
}

var _mSHMIGSourceSetNoSendersCallbackSymbol uintptr

// MSHMIGSourceSetNoSendersCallback has an opaque C signature in discovered metadata.
// Call MSHMIGSourceSetNoSendersCallbackSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/MSHMIGSourceSetNoSendersCallback
func MSHMIGSourceSetNoSendersCallback() {
	panic("hiservices: symbol MSHMIGSourceSetNoSendersCallback has opaque signature; use MSHMIGSourceSetNoSendersCallbackSymbol() and a typed manual wrapper")
}

// MSHMIGSourceSetNoSendersCallbackSymbol returns the raw symbol address for MSHMIGSourceSetNoSendersCallback.
func MSHMIGSourceSetNoSendersCallbackSymbol() uintptr {
	if _mSHMIGSourceSetNoSendersCallbackSymbol == 0 {
		return 0
	}
	return _mSHMIGSourceSetNoSendersCallbackSymbol
}

var _mSHMIGSourceSetSendOnceCallbackSymbol uintptr

// MSHMIGSourceSetSendOnceCallback has an opaque C signature in discovered metadata.
// Call MSHMIGSourceSetSendOnceCallbackSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/MSHMIGSourceSetSendOnceCallback
func MSHMIGSourceSetSendOnceCallback() {
	panic("hiservices: symbol MSHMIGSourceSetSendOnceCallback has opaque signature; use MSHMIGSourceSetSendOnceCallbackSymbol() and a typed manual wrapper")
}

// MSHMIGSourceSetSendOnceCallbackSymbol returns the raw symbol address for MSHMIGSourceSetSendOnceCallback.
func MSHMIGSourceSetSendOnceCallbackSymbol() uintptr {
	if _mSHMIGSourceSetSendOnceCallbackSymbol == 0 {
		return 0
	}
	return _mSHMIGSourceSetSendOnceCallbackSymbol
}

var _newIconActionUPPSymbol uintptr

// NewIconActionUPP has an opaque C signature in discovered metadata.
// Call NewIconActionUPPSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/NewIconActionUPP
func NewIconActionUPP() {
	panic("hiservices: symbol NewIconActionUPP has opaque signature; use NewIconActionUPPSymbol() and a typed manual wrapper")
}

// NewIconActionUPPSymbol returns the raw symbol address for NewIconActionUPP.
func NewIconActionUPPSymbol() uintptr {
	if _newIconActionUPPSymbol == 0 {
		return 0
	}
	return _newIconActionUPPSymbol
}

var _newIconGetterUPPSymbol uintptr

// NewIconGetterUPP has an opaque C signature in discovered metadata.
// Call NewIconGetterUPPSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/NewIconGetterUPP
func NewIconGetterUPP() {
	panic("hiservices: symbol NewIconGetterUPP has opaque signature; use NewIconGetterUPPSymbol() and a typed manual wrapper")
}

// NewIconGetterUPPSymbol returns the raw symbol address for NewIconGetterUPP.
func NewIconGetterUPPSymbol() uintptr {
	if _newIconGetterUPPSymbol == 0 {
		return 0
	}
	return _newIconGetterUPPSymbol
}

var _pasteboardClearSymbol uintptr

// PasteboardClear has an opaque C signature in discovered metadata.
// Call PasteboardClearSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardClear
func PasteboardClear() {
	panic("hiservices: symbol PasteboardClear has opaque signature; use PasteboardClearSymbol() and a typed manual wrapper")
}

// PasteboardClearSymbol returns the raw symbol address for PasteboardClear.
func PasteboardClearSymbol() uintptr {
	if _pasteboardClearSymbol == 0 {
		return 0
	}
	return _pasteboardClearSymbol
}

var _pasteboardCopyItemFlavorDataSymbol uintptr

// PasteboardCopyItemFlavorData has an opaque C signature in discovered metadata.
// Call PasteboardCopyItemFlavorDataSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardCopyItemFlavorData
func PasteboardCopyItemFlavorData() {
	panic("hiservices: symbol PasteboardCopyItemFlavorData has opaque signature; use PasteboardCopyItemFlavorDataSymbol() and a typed manual wrapper")
}

// PasteboardCopyItemFlavorDataSymbol returns the raw symbol address for PasteboardCopyItemFlavorData.
func PasteboardCopyItemFlavorDataSymbol() uintptr {
	if _pasteboardCopyItemFlavorDataSymbol == 0 {
		return 0
	}
	return _pasteboardCopyItemFlavorDataSymbol
}

var _pasteboardCopyItemFlavorsSymbol uintptr

// PasteboardCopyItemFlavors has an opaque C signature in discovered metadata.
// Call PasteboardCopyItemFlavorsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardCopyItemFlavors
func PasteboardCopyItemFlavors() {
	panic("hiservices: symbol PasteboardCopyItemFlavors has opaque signature; use PasteboardCopyItemFlavorsSymbol() and a typed manual wrapper")
}

// PasteboardCopyItemFlavorsSymbol returns the raw symbol address for PasteboardCopyItemFlavors.
func PasteboardCopyItemFlavorsSymbol() uintptr {
	if _pasteboardCopyItemFlavorsSymbol == 0 {
		return 0
	}
	return _pasteboardCopyItemFlavorsSymbol
}

var _pasteboardCopyNameSymbol uintptr

// PasteboardCopyName has an opaque C signature in discovered metadata.
// Call PasteboardCopyNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardCopyName
func PasteboardCopyName() {
	panic("hiservices: symbol PasteboardCopyName has opaque signature; use PasteboardCopyNameSymbol() and a typed manual wrapper")
}

// PasteboardCopyNameSymbol returns the raw symbol address for PasteboardCopyName.
func PasteboardCopyNameSymbol() uintptr {
	if _pasteboardCopyNameSymbol == 0 {
		return 0
	}
	return _pasteboardCopyNameSymbol
}

var _pasteboardCopyPasteLocationSymbol uintptr

// PasteboardCopyPasteLocation has an opaque C signature in discovered metadata.
// Call PasteboardCopyPasteLocationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardCopyPasteLocation
func PasteboardCopyPasteLocation() {
	panic("hiservices: symbol PasteboardCopyPasteLocation has opaque signature; use PasteboardCopyPasteLocationSymbol() and a typed manual wrapper")
}

// PasteboardCopyPasteLocationSymbol returns the raw symbol address for PasteboardCopyPasteLocation.
func PasteboardCopyPasteLocationSymbol() uintptr {
	if _pasteboardCopyPasteLocationSymbol == 0 {
		return 0
	}
	return _pasteboardCopyPasteLocationSymbol
}

var _pasteboardCreateSymbol uintptr

// PasteboardCreate has an opaque C signature in discovered metadata.
// Call PasteboardCreateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardCreate
func PasteboardCreate() {
	panic("hiservices: symbol PasteboardCreate has opaque signature; use PasteboardCreateSymbol() and a typed manual wrapper")
}

// PasteboardCreateSymbol returns the raw symbol address for PasteboardCreate.
func PasteboardCreateSymbol() uintptr {
	if _pasteboardCreateSymbol == 0 {
		return 0
	}
	return _pasteboardCreateSymbol
}

var _pasteboardCreateWithCFPasteboardSymbol uintptr

// PasteboardCreateWithCFPasteboard has an opaque C signature in discovered metadata.
// Call PasteboardCreateWithCFPasteboardSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardCreateWithCFPasteboard
func PasteboardCreateWithCFPasteboard() {
	panic("hiservices: symbol PasteboardCreateWithCFPasteboard has opaque signature; use PasteboardCreateWithCFPasteboardSymbol() and a typed manual wrapper")
}

// PasteboardCreateWithCFPasteboardSymbol returns the raw symbol address for PasteboardCreateWithCFPasteboard.
func PasteboardCreateWithCFPasteboardSymbol() uintptr {
	if _pasteboardCreateWithCFPasteboardSymbol == 0 {
		return 0
	}
	return _pasteboardCreateWithCFPasteboardSymbol
}

var _pasteboardDontHonorPromisesSymbol uintptr

// PasteboardDontHonorPromises has an opaque C signature in discovered metadata.
// Call PasteboardDontHonorPromisesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardDontHonorPromises
func PasteboardDontHonorPromises() {
	panic("hiservices: symbol PasteboardDontHonorPromises has opaque signature; use PasteboardDontHonorPromisesSymbol() and a typed manual wrapper")
}

// PasteboardDontHonorPromisesSymbol returns the raw symbol address for PasteboardDontHonorPromises.
func PasteboardDontHonorPromisesSymbol() uintptr {
	if _pasteboardDontHonorPromisesSymbol == 0 {
		return 0
	}
	return _pasteboardDontHonorPromisesSymbol
}

var _pasteboardGetCFPasteboardSymbol uintptr

// PasteboardGetCFPasteboard has an opaque C signature in discovered metadata.
// Call PasteboardGetCFPasteboardSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardGetCFPasteboard
func PasteboardGetCFPasteboard() {
	panic("hiservices: symbol PasteboardGetCFPasteboard has opaque signature; use PasteboardGetCFPasteboardSymbol() and a typed manual wrapper")
}

// PasteboardGetCFPasteboardSymbol returns the raw symbol address for PasteboardGetCFPasteboard.
func PasteboardGetCFPasteboardSymbol() uintptr {
	if _pasteboardGetCFPasteboardSymbol == 0 {
		return 0
	}
	return _pasteboardGetCFPasteboardSymbol
}

var _pasteboardGetItemCountSymbol uintptr

// PasteboardGetItemCount has an opaque C signature in discovered metadata.
// Call PasteboardGetItemCountSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardGetItemCount
func PasteboardGetItemCount() {
	panic("hiservices: symbol PasteboardGetItemCount has opaque signature; use PasteboardGetItemCountSymbol() and a typed manual wrapper")
}

// PasteboardGetItemCountSymbol returns the raw symbol address for PasteboardGetItemCount.
func PasteboardGetItemCountSymbol() uintptr {
	if _pasteboardGetItemCountSymbol == 0 {
		return 0
	}
	return _pasteboardGetItemCountSymbol
}

var _pasteboardGetItemFlavorFlagsSymbol uintptr

// PasteboardGetItemFlavorFlags has an opaque C signature in discovered metadata.
// Call PasteboardGetItemFlavorFlagsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardGetItemFlavorFlags
func PasteboardGetItemFlavorFlags() {
	panic("hiservices: symbol PasteboardGetItemFlavorFlags has opaque signature; use PasteboardGetItemFlavorFlagsSymbol() and a typed manual wrapper")
}

// PasteboardGetItemFlavorFlagsSymbol returns the raw symbol address for PasteboardGetItemFlavorFlags.
func PasteboardGetItemFlavorFlagsSymbol() uintptr {
	if _pasteboardGetItemFlavorFlagsSymbol == 0 {
		return 0
	}
	return _pasteboardGetItemFlavorFlagsSymbol
}

var _pasteboardGetItemIdentifierSymbol uintptr

// PasteboardGetItemIdentifier has an opaque C signature in discovered metadata.
// Call PasteboardGetItemIdentifierSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardGetItemIdentifier
func PasteboardGetItemIdentifier() {
	panic("hiservices: symbol PasteboardGetItemIdentifier has opaque signature; use PasteboardGetItemIdentifierSymbol() and a typed manual wrapper")
}

// PasteboardGetItemIdentifierSymbol returns the raw symbol address for PasteboardGetItemIdentifier.
func PasteboardGetItemIdentifierSymbol() uintptr {
	if _pasteboardGetItemIdentifierSymbol == 0 {
		return 0
	}
	return _pasteboardGetItemIdentifierSymbol
}

var _pasteboardGetTypeIDSymbol uintptr

// PasteboardGetTypeID has an opaque C signature in discovered metadata.
// Call PasteboardGetTypeIDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardGetTypeID
func PasteboardGetTypeID() {
	panic("hiservices: symbol PasteboardGetTypeID has opaque signature; use PasteboardGetTypeIDSymbol() and a typed manual wrapper")
}

// PasteboardGetTypeIDSymbol returns the raw symbol address for PasteboardGetTypeID.
func PasteboardGetTypeIDSymbol() uintptr {
	if _pasteboardGetTypeIDSymbol == 0 {
		return 0
	}
	return _pasteboardGetTypeIDSymbol
}

var _pasteboardPutItemFlavorSymbol uintptr

// PasteboardPutItemFlavor has an opaque C signature in discovered metadata.
// Call PasteboardPutItemFlavorSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardPutItemFlavor
func PasteboardPutItemFlavor() {
	panic("hiservices: symbol PasteboardPutItemFlavor has opaque signature; use PasteboardPutItemFlavorSymbol() and a typed manual wrapper")
}

// PasteboardPutItemFlavorSymbol returns the raw symbol address for PasteboardPutItemFlavor.
func PasteboardPutItemFlavorSymbol() uintptr {
	if _pasteboardPutItemFlavorSymbol == 0 {
		return 0
	}
	return _pasteboardPutItemFlavorSymbol
}

var _pasteboardResolvePromisesSymbol uintptr

// PasteboardResolvePromises has an opaque C signature in discovered metadata.
// Call PasteboardResolvePromisesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardResolvePromises
func PasteboardResolvePromises() {
	panic("hiservices: symbol PasteboardResolvePromises has opaque signature; use PasteboardResolvePromisesSymbol() and a typed manual wrapper")
}

// PasteboardResolvePromisesSymbol returns the raw symbol address for PasteboardResolvePromises.
func PasteboardResolvePromisesSymbol() uintptr {
	if _pasteboardResolvePromisesSymbol == 0 {
		return 0
	}
	return _pasteboardResolvePromisesSymbol
}

var _pasteboardSetPasteLocationSymbol uintptr

// PasteboardSetPasteLocation has an opaque C signature in discovered metadata.
// Call PasteboardSetPasteLocationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardSetPasteLocation
func PasteboardSetPasteLocation() {
	panic("hiservices: symbol PasteboardSetPasteLocation has opaque signature; use PasteboardSetPasteLocationSymbol() and a typed manual wrapper")
}

// PasteboardSetPasteLocationSymbol returns the raw symbol address for PasteboardSetPasteLocation.
func PasteboardSetPasteLocationSymbol() uintptr {
	if _pasteboardSetPasteLocationSymbol == 0 {
		return 0
	}
	return _pasteboardSetPasteLocationSymbol
}

var _pasteboardSetPromiseKeeperSymbol uintptr

// PasteboardSetPromiseKeeper has an opaque C signature in discovered metadata.
// Call PasteboardSetPromiseKeeperSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardSetPromiseKeeper
func PasteboardSetPromiseKeeper() {
	panic("hiservices: symbol PasteboardSetPromiseKeeper has opaque signature; use PasteboardSetPromiseKeeperSymbol() and a typed manual wrapper")
}

// PasteboardSetPromiseKeeperSymbol returns the raw symbol address for PasteboardSetPromiseKeeper.
func PasteboardSetPromiseKeeperSymbol() uintptr {
	if _pasteboardSetPromiseKeeperSymbol == 0 {
		return 0
	}
	return _pasteboardSetPromiseKeeperSymbol
}

var _pasteboardSynchronizeSymbol uintptr

// PasteboardSynchronize has an opaque C signature in discovered metadata.
// Call PasteboardSynchronizeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardSynchronize
func PasteboardSynchronize() {
	panic("hiservices: symbol PasteboardSynchronize has opaque signature; use PasteboardSynchronizeSymbol() and a typed manual wrapper")
}

// PasteboardSynchronizeSymbol returns the raw symbol address for PasteboardSynchronize.
func PasteboardSynchronizeSymbol() uintptr {
	if _pasteboardSynchronizeSymbol == 0 {
		return 0
	}
	return _pasteboardSynchronizeSymbol
}

var _pasteboardToggleDuplicateFlavorCheckSymbol uintptr

// PasteboardToggleDuplicateFlavorCheck has an opaque C signature in discovered metadata.
// Call PasteboardToggleDuplicateFlavorCheckSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PasteboardToggleDuplicateFlavorCheck
func PasteboardToggleDuplicateFlavorCheck() {
	panic("hiservices: symbol PasteboardToggleDuplicateFlavorCheck has opaque signature; use PasteboardToggleDuplicateFlavorCheckSymbol() and a typed manual wrapper")
}

// PasteboardToggleDuplicateFlavorCheckSymbol returns the raw symbol address for PasteboardToggleDuplicateFlavorCheck.
func PasteboardToggleDuplicateFlavorCheckSymbol() uintptr {
	if _pasteboardToggleDuplicateFlavorCheckSymbol == 0 {
		return 0
	}
	return _pasteboardToggleDuplicateFlavorCheckSymbol
}

var _plotIconRefInContextSymbol uintptr

// PlotIconRefInContext has an opaque C signature in discovered metadata.
// Call PlotIconRefInContextSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/PlotIconRefInContext
func PlotIconRefInContext() {
	panic("hiservices: symbol PlotIconRefInContext has opaque signature; use PlotIconRefInContextSymbol() and a typed manual wrapper")
}

// PlotIconRefInContextSymbol returns the raw symbol address for PlotIconRefInContext.
func PlotIconRefInContextSymbol() uintptr {
	if _plotIconRefInContextSymbol == 0 {
		return 0
	}
	return _plotIconRefInContextSymbol
}

var _processInformationCopyDictionarySymbol uintptr

// ProcessInformationCopyDictionary has an opaque C signature in discovered metadata.
// Call ProcessInformationCopyDictionarySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ProcessInformationCopyDictionary
func ProcessInformationCopyDictionary() {
	panic("hiservices: symbol ProcessInformationCopyDictionary has opaque signature; use ProcessInformationCopyDictionarySymbol() and a typed manual wrapper")
}

// ProcessInformationCopyDictionarySymbol returns the raw symbol address for ProcessInformationCopyDictionary.
func ProcessInformationCopyDictionarySymbol() uintptr {
	if _processInformationCopyDictionarySymbol == 0 {
		return 0
	}
	return _processInformationCopyDictionarySymbol
}

var _sXArbitrationAddQueuedOutputRequestSymbol uintptr

// SXArbitrationAddQueuedOutputRequest has an opaque C signature in discovered metadata.
// Call SXArbitrationAddQueuedOutputRequestSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SXArbitrationAddQueuedOutputRequest
func SXArbitrationAddQueuedOutputRequest() {
	panic("hiservices: symbol SXArbitrationAddQueuedOutputRequest has opaque signature; use SXArbitrationAddQueuedOutputRequestSymbol() and a typed manual wrapper")
}

// SXArbitrationAddQueuedOutputRequestSymbol returns the raw symbol address for SXArbitrationAddQueuedOutputRequest.
func SXArbitrationAddQueuedOutputRequestSymbol() uintptr {
	if _sXArbitrationAddQueuedOutputRequestSymbol == 0 {
		return 0
	}
	return _sXArbitrationAddQueuedOutputRequestSymbol
}

var _sXArbitrationCancelQueuedRequestSymbol uintptr

// SXArbitrationCancelQueuedRequest has an opaque C signature in discovered metadata.
// Call SXArbitrationCancelQueuedRequestSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SXArbitrationCancelQueuedRequest
func SXArbitrationCancelQueuedRequest() {
	panic("hiservices: symbol SXArbitrationCancelQueuedRequest has opaque signature; use SXArbitrationCancelQueuedRequestSymbol() and a typed manual wrapper")
}

// SXArbitrationCancelQueuedRequestSymbol returns the raw symbol address for SXArbitrationCancelQueuedRequest.
func SXArbitrationCancelQueuedRequestSymbol() uintptr {
	if _sXArbitrationCancelQueuedRequestSymbol == 0 {
		return 0
	}
	return _sXArbitrationCancelQueuedRequestSymbol
}

var _sXArbitrationCreateServerSourceSymbol uintptr

// SXArbitrationCreateServerSource has an opaque C signature in discovered metadata.
// Call SXArbitrationCreateServerSourceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SXArbitrationCreateServerSource
func SXArbitrationCreateServerSource() {
	panic("hiservices: symbol SXArbitrationCreateServerSource has opaque signature; use SXArbitrationCreateServerSourceSymbol() and a typed manual wrapper")
}

// SXArbitrationCreateServerSourceSymbol returns the raw symbol address for SXArbitrationCreateServerSource.
func SXArbitrationCreateServerSourceSymbol() uintptr {
	if _sXArbitrationCreateServerSourceSymbol == 0 {
		return 0
	}
	return _sXArbitrationCreateServerSourceSymbol
}

var _sXArbitrationIsQueuedRequestPendingSymbol uintptr

// SXArbitrationIsQueuedRequestPending has an opaque C signature in discovered metadata.
// Call SXArbitrationIsQueuedRequestPendingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SXArbitrationIsQueuedRequestPending
func SXArbitrationIsQueuedRequestPending() {
	panic("hiservices: symbol SXArbitrationIsQueuedRequestPending has opaque signature; use SXArbitrationIsQueuedRequestPendingSymbol() and a typed manual wrapper")
}

// SXArbitrationIsQueuedRequestPendingSymbol returns the raw symbol address for SXArbitrationIsQueuedRequestPending.
func SXArbitrationIsQueuedRequestPendingSymbol() uintptr {
	if _sXArbitrationIsQueuedRequestPendingSymbol == 0 {
		return 0
	}
	return _sXArbitrationIsQueuedRequestPendingSymbol
}

var _sXArbitrationRegisterOutputStartingSymbol uintptr

// SXArbitrationRegisterOutputStarting has an opaque C signature in discovered metadata.
// Call SXArbitrationRegisterOutputStartingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SXArbitrationRegisterOutputStarting
func SXArbitrationRegisterOutputStarting() {
	panic("hiservices: symbol SXArbitrationRegisterOutputStarting has opaque signature; use SXArbitrationRegisterOutputStartingSymbol() and a typed manual wrapper")
}

// SXArbitrationRegisterOutputStartingSymbol returns the raw symbol address for SXArbitrationRegisterOutputStarting.
func SXArbitrationRegisterOutputStartingSymbol() uintptr {
	if _sXArbitrationRegisterOutputStartingSymbol == 0 {
		return 0
	}
	return _sXArbitrationRegisterOutputStartingSymbol
}

var _sXArbitrationRegisterOutputStoppedSymbol uintptr

// SXArbitrationRegisterOutputStopped has an opaque C signature in discovered metadata.
// Call SXArbitrationRegisterOutputStoppedSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SXArbitrationRegisterOutputStopped
func SXArbitrationRegisterOutputStopped() {
	panic("hiservices: symbol SXArbitrationRegisterOutputStopped has opaque signature; use SXArbitrationRegisterOutputStoppedSymbol() and a typed manual wrapper")
}

// SXArbitrationRegisterOutputStoppedSymbol returns the raw symbol address for SXArbitrationRegisterOutputStopped.
func SXArbitrationRegisterOutputStoppedSymbol() uintptr {
	if _sXArbitrationRegisterOutputStoppedSymbol == 0 {
		return 0
	}
	return _sXArbitrationRegisterOutputStoppedSymbol
}

var _sameProcessSymbol uintptr

// SameProcess has an opaque C signature in discovered metadata.
// Call SameProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SameProcess
func SameProcess() {
	panic("hiservices: symbol SameProcess has opaque signature; use SameProcessSymbol() and a typed manual wrapper")
}

// SameProcessSymbol returns the raw symbol address for SameProcess.
func SameProcessSymbol() uintptr {
	if _sameProcessSymbol == 0 {
		return 0
	}
	return _sameProcessSymbol
}

var _serializeCFTypeSymbol uintptr

// SerializeCFType has an opaque C signature in discovered metadata.
// Call SerializeCFTypeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SerializeCFType
func SerializeCFType() {
	panic("hiservices: symbol SerializeCFType has opaque signature; use SerializeCFTypeSymbol() and a typed manual wrapper")
}

// SerializeCFTypeSymbol returns the raw symbol address for SerializeCFType.
func SerializeCFTypeSymbol() uintptr {
	if _serializeCFTypeSymbol == 0 {
		return 0
	}
	return _serializeCFTypeSymbol
}

var _setApplicationIsDaemonSymbol uintptr

// SetApplicationIsDaemon has an opaque C signature in discovered metadata.
// Call SetApplicationIsDaemonSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SetApplicationIsDaemon
func SetApplicationIsDaemon() {
	panic("hiservices: symbol SetApplicationIsDaemon has opaque signature; use SetApplicationIsDaemonSymbol() and a typed manual wrapper")
}

// SetApplicationIsDaemonSymbol returns the raw symbol address for SetApplicationIsDaemon.
func SetApplicationIsDaemonSymbol() uintptr {
	if _setApplicationIsDaemonSymbol == 0 {
		return 0
	}
	return _setApplicationIsDaemonSymbol
}

var _setFrontProcessSymbol uintptr

// SetFrontProcess has an opaque C signature in discovered metadata.
// Call SetFrontProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SetFrontProcess
func SetFrontProcess() {
	panic("hiservices: symbol SetFrontProcess has opaque signature; use SetFrontProcessSymbol() and a typed manual wrapper")
}

// SetFrontProcessSymbol returns the raw symbol address for SetFrontProcess.
func SetFrontProcessSymbol() uintptr {
	if _setFrontProcessSymbol == 0 {
		return 0
	}
	return _setFrontProcessSymbol
}

var _setFrontProcessWithOptionsSymbol uintptr

// SetFrontProcessWithOptions has an opaque C signature in discovered metadata.
// Call SetFrontProcessWithOptionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SetFrontProcessWithOptions
func SetFrontProcessWithOptions() {
	panic("hiservices: symbol SetFrontProcessWithOptions has opaque signature; use SetFrontProcessWithOptionsSymbol() and a typed manual wrapper")
}

// SetFrontProcessWithOptionsSymbol returns the raw symbol address for SetFrontProcessWithOptions.
func SetFrontProcessWithOptionsSymbol() uintptr {
	if _setFrontProcessWithOptionsSymbol == 0 {
		return 0
	}
	return _setFrontProcessWithOptionsSymbol
}

var _setGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol uintptr

// SetGlobalIconImagesCacheMaxEntriesAndMaxDataSize has an opaque C signature in discovered metadata.
// Call SetGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SetGlobalIconImagesCacheMaxEntriesAndMaxDataSize
func SetGlobalIconImagesCacheMaxEntriesAndMaxDataSize() {
	panic("hiservices: symbol SetGlobalIconImagesCacheMaxEntriesAndMaxDataSize has opaque signature; use SetGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol() and a typed manual wrapper")
}

// SetGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol returns the raw symbol address for SetGlobalIconImagesCacheMaxEntriesAndMaxDataSize.
func SetGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol() uintptr {
	if _setGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol == 0 {
		return 0
	}
	return _setGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol
}

var _setIconFamilyDataSymbol uintptr

// SetIconFamilyData has an opaque C signature in discovered metadata.
// Call SetIconFamilyDataSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SetIconFamilyData
func SetIconFamilyData() {
	panic("hiservices: symbol SetIconFamilyData has opaque signature; use SetIconFamilyDataSymbol() and a typed manual wrapper")
}

// SetIconFamilyDataSymbol returns the raw symbol address for SetIconFamilyData.
func SetIconFamilyDataSymbol() uintptr {
	if _setIconFamilyDataSymbol == 0 {
		return 0
	}
	return _setIconFamilyDataSymbol
}

var _setLabelColorAndNameSymbol uintptr

// SetLabelColorAndName has an opaque C signature in discovered metadata.
// Call SetLabelColorAndNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/SetLabelColorAndName
func SetLabelColorAndName() {
	panic("hiservices: symbol SetLabelColorAndName has opaque signature; use SetLabelColorAndNameSymbol() and a typed manual wrapper")
}

// SetLabelColorAndNameSymbol returns the raw symbol address for SetLabelColorAndName.
func SetLabelColorAndNameSymbol() uintptr {
	if _setLabelColorAndNameSymbol == 0 {
		return 0
	}
	return _setLabelColorAndNameSymbol
}

var _showHideDragSymbol uintptr

// ShowHideDrag has an opaque C signature in discovered metadata.
// Call ShowHideDragSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ShowHideDrag
func ShowHideDrag() {
	panic("hiservices: symbol ShowHideDrag has opaque signature; use ShowHideDragSymbol() and a typed manual wrapper")
}

// ShowHideDragSymbol returns the raw symbol address for ShowHideDrag.
func ShowHideDragSymbol() uintptr {
	if _showHideDragSymbol == 0 {
		return 0
	}
	return _showHideDragSymbol
}

var _showHideProcessSymbol uintptr

// ShowHideProcess has an opaque C signature in discovered metadata.
// Call ShowHideProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/ShowHideProcess
func ShowHideProcess() {
	panic("hiservices: symbol ShowHideProcess has opaque signature; use ShowHideProcessSymbol() and a typed manual wrapper")
}

// ShowHideProcessSymbol returns the raw symbol address for ShowHideProcess.
func ShowHideProcessSymbol() uintptr {
	if _showHideProcessSymbol == 0 {
		return 0
	}
	return _showHideProcessSymbol
}

var _startIPCFlockingPingSymbol uintptr

// StartIPCFlockingPing has an opaque C signature in discovered metadata.
// Call StartIPCFlockingPingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/StartIPCFlockingPing
func StartIPCFlockingPing() {
	panic("hiservices: symbol StartIPCFlockingPing has opaque signature; use StartIPCFlockingPingSymbol() and a typed manual wrapper")
}

// StartIPCFlockingPingSymbol returns the raw symbol address for StartIPCFlockingPing.
func StartIPCFlockingPingSymbol() uintptr {
	if _startIPCFlockingPingSymbol == 0 {
		return 0
	}
	return _startIPCFlockingPingSymbol
}

var _startIPCPingSymbol uintptr

// StartIPCPing has an opaque C signature in discovered metadata.
// Call StartIPCPingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/StartIPCPing
func StartIPCPing() {
	panic("hiservices: symbol StartIPCPing has opaque signature; use StartIPCPingSymbol() and a typed manual wrapper")
}

// StartIPCPingSymbol returns the raw symbol address for StartIPCPing.
func StartIPCPingSymbol() uintptr {
	if _startIPCPingSymbol == 0 {
		return 0
	}
	return _startIPCPingSymbol
}

var _transformProcessTypeSymbol uintptr

// TransformProcessType has an opaque C signature in discovered metadata.
// Call TransformProcessTypeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/TransformProcessType
func TransformProcessType() {
	panic("hiservices: symbol TransformProcessType has opaque signature; use TransformProcessTypeSymbol() and a typed manual wrapper")
}

// TransformProcessTypeSymbol returns the raw symbol address for TransformProcessType.
func TransformProcessTypeSymbol() uintptr {
	if _transformProcessTypeSymbol == 0 {
		return 0
	}
	return _transformProcessTypeSymbol
}

var _translationCopyDestinationTypeSymbol uintptr

// TranslationCopyDestinationType has an opaque C signature in discovered metadata.
// Call TranslationCopyDestinationTypeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/TranslationCopyDestinationType
func TranslationCopyDestinationType() {
	panic("hiservices: symbol TranslationCopyDestinationType has opaque signature; use TranslationCopyDestinationTypeSymbol() and a typed manual wrapper")
}

// TranslationCopyDestinationTypeSymbol returns the raw symbol address for TranslationCopyDestinationType.
func TranslationCopyDestinationTypeSymbol() uintptr {
	if _translationCopyDestinationTypeSymbol == 0 {
		return 0
	}
	return _translationCopyDestinationTypeSymbol
}

var _translationCopySourceTypeSymbol uintptr

// TranslationCopySourceType has an opaque C signature in discovered metadata.
// Call TranslationCopySourceTypeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/TranslationCopySourceType
func TranslationCopySourceType() {
	panic("hiservices: symbol TranslationCopySourceType has opaque signature; use TranslationCopySourceTypeSymbol() and a typed manual wrapper")
}

// TranslationCopySourceTypeSymbol returns the raw symbol address for TranslationCopySourceType.
func TranslationCopySourceTypeSymbol() uintptr {
	if _translationCopySourceTypeSymbol == 0 {
		return 0
	}
	return _translationCopySourceTypeSymbol
}

var _translationCreateSymbol uintptr

// TranslationCreate has an opaque C signature in discovered metadata.
// Call TranslationCreateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/TranslationCreate
func TranslationCreate() {
	panic("hiservices: symbol TranslationCreate has opaque signature; use TranslationCreateSymbol() and a typed manual wrapper")
}

// TranslationCreateSymbol returns the raw symbol address for TranslationCreate.
func TranslationCreateSymbol() uintptr {
	if _translationCreateSymbol == 0 {
		return 0
	}
	return _translationCreateSymbol
}

var _translationCreateWithSourceArraySymbol uintptr

// TranslationCreateWithSourceArray has an opaque C signature in discovered metadata.
// Call TranslationCreateWithSourceArraySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/TranslationCreateWithSourceArray
func TranslationCreateWithSourceArray() {
	panic("hiservices: symbol TranslationCreateWithSourceArray has opaque signature; use TranslationCreateWithSourceArraySymbol() and a typed manual wrapper")
}

// TranslationCreateWithSourceArraySymbol returns the raw symbol address for TranslationCreateWithSourceArray.
func TranslationCreateWithSourceArraySymbol() uintptr {
	if _translationCreateWithSourceArraySymbol == 0 {
		return 0
	}
	return _translationCreateWithSourceArraySymbol
}

var _translationGetTranslationFlagsSymbol uintptr

// TranslationGetTranslationFlags has an opaque C signature in discovered metadata.
// Call TranslationGetTranslationFlagsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/TranslationGetTranslationFlags
func TranslationGetTranslationFlags() {
	panic("hiservices: symbol TranslationGetTranslationFlags has opaque signature; use TranslationGetTranslationFlagsSymbol() and a typed manual wrapper")
}

// TranslationGetTranslationFlagsSymbol returns the raw symbol address for TranslationGetTranslationFlags.
func TranslationGetTranslationFlagsSymbol() uintptr {
	if _translationGetTranslationFlagsSymbol == 0 {
		return 0
	}
	return _translationGetTranslationFlagsSymbol
}

var _translationGetTypeIDSymbol uintptr

// TranslationGetTypeID has an opaque C signature in discovered metadata.
// Call TranslationGetTypeIDSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/TranslationGetTypeID
func TranslationGetTypeID() {
	panic("hiservices: symbol TranslationGetTypeID has opaque signature; use TranslationGetTypeIDSymbol() and a typed manual wrapper")
}

// TranslationGetTypeIDSymbol returns the raw symbol address for TranslationGetTypeID.
func TranslationGetTypeIDSymbol() uintptr {
	if _translationGetTypeIDSymbol == 0 {
		return 0
	}
	return _translationGetTypeIDSymbol
}

var _translationPerformForDataSymbol uintptr

// TranslationPerformForData has an opaque C signature in discovered metadata.
// Call TranslationPerformForDataSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/TranslationPerformForData
func TranslationPerformForData() {
	panic("hiservices: symbol TranslationPerformForData has opaque signature; use TranslationPerformForDataSymbol() and a typed manual wrapper")
}

// TranslationPerformForDataSymbol returns the raw symbol address for TranslationPerformForData.
func TranslationPerformForDataSymbol() uintptr {
	if _translationPerformForDataSymbol == 0 {
		return 0
	}
	return _translationPerformForDataSymbol
}

var _translationPerformForFileSymbol uintptr

// TranslationPerformForFile has an opaque C signature in discovered metadata.
// Call TranslationPerformForFileSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/TranslationPerformForFile
func TranslationPerformForFile() {
	panic("hiservices: symbol TranslationPerformForFile has opaque signature; use TranslationPerformForFileSymbol() and a typed manual wrapper")
}

// TranslationPerformForFileSymbol returns the raw symbol address for TranslationPerformForFile.
func TranslationPerformForFileSymbol() uintptr {
	if _translationPerformForFileSymbol == 0 {
		return 0
	}
	return _translationPerformForFileSymbol
}

var _translationPerformForURLSymbol uintptr

// TranslationPerformForURL has an opaque C signature in discovered metadata.
// Call TranslationPerformForURLSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/TranslationPerformForURL
func TranslationPerformForURL() {
	panic("hiservices: symbol TranslationPerformForURL has opaque signature; use TranslationPerformForURLSymbol() and a typed manual wrapper")
}

// TranslationPerformForURLSymbol returns the raw symbol address for TranslationPerformForURL.
func TranslationPerformForURLSymbol() uintptr {
	if _translationPerformForURLSymbol == 0 {
		return 0
	}
	return _translationPerformForURLSymbol
}

var _uAZoomChangeFocusSymbol uintptr

// UAZoomChangeFocus has an opaque C signature in discovered metadata.
// Call UAZoomChangeFocusSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/UAZoomChangeFocus
func UAZoomChangeFocus() {
	panic("hiservices: symbol UAZoomChangeFocus has opaque signature; use UAZoomChangeFocusSymbol() and a typed manual wrapper")
}

// UAZoomChangeFocusSymbol returns the raw symbol address for UAZoomChangeFocus.
func UAZoomChangeFocusSymbol() uintptr {
	if _uAZoomChangeFocusSymbol == 0 {
		return 0
	}
	return _uAZoomChangeFocusSymbol
}

var _uAZoomEnabledSymbol uintptr

// UAZoomEnabled has an opaque C signature in discovered metadata.
// Call UAZoomEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/UAZoomEnabled
func UAZoomEnabled() {
	panic("hiservices: symbol UAZoomEnabled has opaque signature; use UAZoomEnabledSymbol() and a typed manual wrapper")
}

// UAZoomEnabledSymbol returns the raw symbol address for UAZoomEnabled.
func UAZoomEnabledSymbol() uintptr {
	if _uAZoomEnabledSymbol == 0 {
		return 0
	}
	return _uAZoomEnabledSymbol
}

var _unserializeCFTypeSymbol uintptr

// UnserializeCFType has an opaque C signature in discovered metadata.
// Call UnserializeCFTypeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/UnserializeCFType
func UnserializeCFType() {
	panic("hiservices: symbol UnserializeCFType has opaque signature; use UnserializeCFTypeSymbol() and a typed manual wrapper")
}

// UnserializeCFTypeSymbol returns the raw symbol address for UnserializeCFType.
func UnserializeCFTypeSymbol() uintptr {
	if _unserializeCFTypeSymbol == 0 {
		return 0
	}
	return _unserializeCFTypeSymbol
}

var _wakeUpProcessSymbol uintptr

// WakeUpProcess has an opaque C signature in discovered metadata.
// Call WakeUpProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/WakeUpProcess
func WakeUpProcess() {
	panic("hiservices: symbol WakeUpProcess has opaque signature; use WakeUpProcessSymbol() and a typed manual wrapper")
}

// WakeUpProcessSymbol returns the raw symbol address for WakeUpProcess.
func WakeUpProcessSymbol() uintptr {
	if _wakeUpProcessSymbol == 0 {
		return 0
	}
	return _wakeUpProcessSymbol
}

var __AXCopyActionDescriptionSymbol uintptr

// _AXCopyActionDescription has an opaque C signature in discovered metadata.
// Call _AXCopyActionDescriptionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXCopyActionDescription
func _AXCopyActionDescription() {
	panic("hiservices: symbol _AXCopyActionDescription has opaque signature; use _AXCopyActionDescriptionSymbol() and a typed manual wrapper")
}

// _AXCopyActionDescriptionSymbol returns the raw symbol address for _AXCopyActionDescription.
func _AXCopyActionDescriptionSymbol() uintptr {
	if __AXCopyActionDescriptionSymbol == 0 {
		return 0
	}
	return __AXCopyActionDescriptionSymbol
}

var __AXCopyChildrenHashSymbol uintptr

// _AXCopyChildrenHash has an opaque C signature in discovered metadata.
// Call _AXCopyChildrenHashSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXCopyChildrenHash
func _AXCopyChildrenHash() {
	panic("hiservices: symbol _AXCopyChildrenHash has opaque signature; use _AXCopyChildrenHashSymbol() and a typed manual wrapper")
}

// _AXCopyChildrenHashSymbol returns the raw symbol address for _AXCopyChildrenHash.
func _AXCopyChildrenHashSymbol() uintptr {
	if __AXCopyChildrenHashSymbol == 0 {
		return 0
	}
	return __AXCopyChildrenHashSymbol
}

var __AXCopyChildrenHashWithRelativeFrameSymbol uintptr

// _AXCopyChildrenHashWithRelativeFrame has an opaque C signature in discovered metadata.
// Call _AXCopyChildrenHashWithRelativeFrameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXCopyChildrenHashWithRelativeFrame
func _AXCopyChildrenHashWithRelativeFrame() {
	panic("hiservices: symbol _AXCopyChildrenHashWithRelativeFrame has opaque signature; use _AXCopyChildrenHashWithRelativeFrameSymbol() and a typed manual wrapper")
}

// _AXCopyChildrenHashWithRelativeFrameSymbol returns the raw symbol address for _AXCopyChildrenHashWithRelativeFrame.
func _AXCopyChildrenHashWithRelativeFrameSymbol() uintptr {
	if __AXCopyChildrenHashWithRelativeFrameSymbol == 0 {
		return 0
	}
	return __AXCopyChildrenHashWithRelativeFrameSymbol
}

var __AXCopyRoleDescriptionSymbol uintptr

// _AXCopyRoleDescription has an opaque C signature in discovered metadata.
// Call _AXCopyRoleDescriptionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXCopyRoleDescription
func _AXCopyRoleDescription() {
	panic("hiservices: symbol _AXCopyRoleDescription has opaque signature; use _AXCopyRoleDescriptionSymbol() and a typed manual wrapper")
}

// _AXCopyRoleDescriptionSymbol returns the raw symbol address for _AXCopyRoleDescription.
func _AXCopyRoleDescriptionSymbol() uintptr {
	if __AXCopyRoleDescriptionSymbol == 0 {
		return 0
	}
	return __AXCopyRoleDescriptionSymbol
}

var __AXCopyRoleDescriptionWithSubroleSymbol uintptr

// _AXCopyRoleDescriptionWithSubrole has an opaque C signature in discovered metadata.
// Call _AXCopyRoleDescriptionWithSubroleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXCopyRoleDescriptionWithSubrole
func _AXCopyRoleDescriptionWithSubrole() {
	panic("hiservices: symbol _AXCopyRoleDescriptionWithSubrole has opaque signature; use _AXCopyRoleDescriptionWithSubroleSymbol() and a typed manual wrapper")
}

// _AXCopyRoleDescriptionWithSubroleSymbol returns the raw symbol address for _AXCopyRoleDescriptionWithSubrole.
func _AXCopyRoleDescriptionWithSubroleSymbol() uintptr {
	if __AXCopyRoleDescriptionWithSubroleSymbol == 0 {
		return 0
	}
	return __AXCopyRoleDescriptionWithSubroleSymbol
}

var __AXCopyTitleSymbol uintptr

// _AXCopyTitle has an opaque C signature in discovered metadata.
// Call _AXCopyTitleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXCopyTitle
func _AXCopyTitle() {
	panic("hiservices: symbol _AXCopyTitle has opaque signature; use _AXCopyTitleSymbol() and a typed manual wrapper")
}

// _AXCopyTitleSymbol returns the raw symbol address for _AXCopyTitle.
func _AXCopyTitleSymbol() uintptr {
	if __AXCopyTitleSymbol == 0 {
		return 0
	}
	return __AXCopyTitleSymbol
}

var __AXCreateElementOrderingSymbol uintptr

// _AXCreateElementOrdering has an opaque C signature in discovered metadata.
// Call _AXCreateElementOrderingSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXCreateElementOrdering
func _AXCreateElementOrdering() {
	panic("hiservices: symbol _AXCreateElementOrdering has opaque signature; use _AXCreateElementOrderingSymbol() and a typed manual wrapper")
}

// _AXCreateElementOrderingSymbol returns the raw symbol address for _AXCreateElementOrdering.
func _AXCreateElementOrderingSymbol() uintptr {
	if __AXCreateElementOrderingSymbol == 0 {
		return 0
	}
	return __AXCreateElementOrderingSymbol
}

var __AXCurrentRequestCanAccessRemoteDeviceContentSymbol uintptr

// _AXCurrentRequestCanAccessRemoteDeviceContent has an opaque C signature in discovered metadata.
// Call _AXCurrentRequestCanAccessRemoteDeviceContentSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXCurrentRequestCanAccessRemoteDeviceContent
func _AXCurrentRequestCanAccessRemoteDeviceContent() {
	panic("hiservices: symbol _AXCurrentRequestCanAccessRemoteDeviceContent has opaque signature; use _AXCurrentRequestCanAccessRemoteDeviceContentSymbol() and a typed manual wrapper")
}

// _AXCurrentRequestCanAccessRemoteDeviceContentSymbol returns the raw symbol address for _AXCurrentRequestCanAccessRemoteDeviceContent.
func _AXCurrentRequestCanAccessRemoteDeviceContentSymbol() uintptr {
	if __AXCurrentRequestCanAccessRemoteDeviceContentSymbol == 0 {
		return 0
	}
	return __AXCurrentRequestCanAccessRemoteDeviceContentSymbol
}

var __AXCurrentRequestCanReturnInspectionContentSymbol uintptr

// _AXCurrentRequestCanReturnInspectionContent has an opaque C signature in discovered metadata.
// Call _AXCurrentRequestCanReturnInspectionContentSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXCurrentRequestCanReturnInspectionContent
func _AXCurrentRequestCanReturnInspectionContent() {
	panic("hiservices: symbol _AXCurrentRequestCanReturnInspectionContent has opaque signature; use _AXCurrentRequestCanReturnInspectionContentSymbol() and a typed manual wrapper")
}

// _AXCurrentRequestCanReturnInspectionContentSymbol returns the raw symbol address for _AXCurrentRequestCanReturnInspectionContent.
func _AXCurrentRequestCanReturnInspectionContentSymbol() uintptr {
	if __AXCurrentRequestCanReturnInspectionContentSymbol == 0 {
		return 0
	}
	return __AXCurrentRequestCanReturnInspectionContentSymbol
}

var __AXCurrentRequestCanReturnProtectedContentSymbol uintptr

// _AXCurrentRequestCanReturnProtectedContent has an opaque C signature in discovered metadata.
// Call _AXCurrentRequestCanReturnProtectedContentSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXCurrentRequestCanReturnProtectedContent
func _AXCurrentRequestCanReturnProtectedContent() {
	panic("hiservices: symbol _AXCurrentRequestCanReturnProtectedContent has opaque signature; use _AXCurrentRequestCanReturnProtectedContentSymbol() and a typed manual wrapper")
}

// _AXCurrentRequestCanReturnProtectedContentSymbol returns the raw symbol address for _AXCurrentRequestCanReturnProtectedContent.
func _AXCurrentRequestCanReturnProtectedContentSymbol() uintptr {
	if __AXCurrentRequestCanReturnProtectedContentSymbol == 0 {
		return 0
	}
	return __AXCurrentRequestCanReturnProtectedContentSymbol
}

var __AXGetClientForCurrentRequestUntrustedSymbol uintptr

// _AXGetClientForCurrentRequestUntrusted has an opaque C signature in discovered metadata.
// Call _AXGetClientForCurrentRequestUntrustedSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXGetClientForCurrentRequestUntrusted
func _AXGetClientForCurrentRequestUntrusted() {
	panic("hiservices: symbol _AXGetClientForCurrentRequestUntrusted has opaque signature; use _AXGetClientForCurrentRequestUntrustedSymbol() and a typed manual wrapper")
}

// _AXGetClientForCurrentRequestUntrustedSymbol returns the raw symbol address for _AXGetClientForCurrentRequestUntrusted.
func _AXGetClientForCurrentRequestUntrustedSymbol() uintptr {
	if __AXGetClientForCurrentRequestUntrustedSymbol == 0 {
		return 0
	}
	return __AXGetClientForCurrentRequestUntrustedSymbol
}

var __AXHasClientsWithAccessRemoteDeviceContentSymbol uintptr

// _AXHasClientsWithAccessRemoteDeviceContent has an opaque C signature in discovered metadata.
// Call _AXHasClientsWithAccessRemoteDeviceContentSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXHasClientsWithAccessRemoteDeviceContent
func _AXHasClientsWithAccessRemoteDeviceContent() {
	panic("hiservices: symbol _AXHasClientsWithAccessRemoteDeviceContent has opaque signature; use _AXHasClientsWithAccessRemoteDeviceContentSymbol() and a typed manual wrapper")
}

// _AXHasClientsWithAccessRemoteDeviceContentSymbol returns the raw symbol address for _AXHasClientsWithAccessRemoteDeviceContent.
func _AXHasClientsWithAccessRemoteDeviceContentSymbol() uintptr {
	if __AXHasClientsWithAccessRemoteDeviceContentSymbol == 0 {
		return 0
	}
	return __AXHasClientsWithAccessRemoteDeviceContentSymbol
}

var __AXInterfaceCopyCursorColorFillSymbol uintptr

// _AXInterfaceCopyCursorColorFill has an opaque C signature in discovered metadata.
// Call _AXInterfaceCopyCursorColorFillSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceCopyCursorColorFill
func _AXInterfaceCopyCursorColorFill() {
	panic("hiservices: symbol _AXInterfaceCopyCursorColorFill has opaque signature; use _AXInterfaceCopyCursorColorFillSymbol() and a typed manual wrapper")
}

// _AXInterfaceCopyCursorColorFillSymbol returns the raw symbol address for _AXInterfaceCopyCursorColorFill.
func _AXInterfaceCopyCursorColorFillSymbol() uintptr {
	if __AXInterfaceCopyCursorColorFillSymbol == 0 {
		return 0
	}
	return __AXInterfaceCopyCursorColorFillSymbol
}

var __AXInterfaceCopyCursorColorOutlineSymbol uintptr

// _AXInterfaceCopyCursorColorOutline has an opaque C signature in discovered metadata.
// Call _AXInterfaceCopyCursorColorOutlineSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceCopyCursorColorOutline
func _AXInterfaceCopyCursorColorOutline() {
	panic("hiservices: symbol _AXInterfaceCopyCursorColorOutline has opaque signature; use _AXInterfaceCopyCursorColorOutlineSymbol() and a typed manual wrapper")
}

// _AXInterfaceCopyCursorColorOutlineSymbol returns the raw symbol address for _AXInterfaceCopyCursorColorOutline.
func _AXInterfaceCopyCursorColorOutlineSymbol() uintptr {
	if __AXInterfaceCopyCursorColorOutlineSymbol == 0 {
		return 0
	}
	return __AXInterfaceCopyCursorColorOutlineSymbol
}

var __AXInterfaceCursorIsOverriddenSymbol uintptr

// _AXInterfaceCursorIsOverridden has an opaque C signature in discovered metadata.
// Call _AXInterfaceCursorIsOverriddenSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceCursorIsOverridden
func _AXInterfaceCursorIsOverridden() {
	panic("hiservices: symbol _AXInterfaceCursorIsOverridden has opaque signature; use _AXInterfaceCursorIsOverriddenSymbol() and a typed manual wrapper")
}

// _AXInterfaceCursorIsOverriddenSymbol returns the raw symbol address for _AXInterfaceCursorIsOverridden.
func _AXInterfaceCursorIsOverriddenSymbol() uintptr {
	if __AXInterfaceCursorIsOverriddenSymbol == 0 {
		return 0
	}
	return __AXInterfaceCursorIsOverriddenSymbol
}

var __AXInterfaceCursorSetAndReturnSeedSymbol uintptr

// _AXInterfaceCursorSetAndReturnSeed has an opaque C signature in discovered metadata.
// Call _AXInterfaceCursorSetAndReturnSeedSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceCursorSetAndReturnSeed
func _AXInterfaceCursorSetAndReturnSeed() {
	panic("hiservices: symbol _AXInterfaceCursorSetAndReturnSeed has opaque signature; use _AXInterfaceCursorSetAndReturnSeedSymbol() and a typed manual wrapper")
}

// _AXInterfaceCursorSetAndReturnSeedSymbol returns the raw symbol address for _AXInterfaceCursorSetAndReturnSeed.
func _AXInterfaceCursorSetAndReturnSeedSymbol() uintptr {
	if __AXInterfaceCursorSetAndReturnSeedSymbol == 0 {
		return 0
	}
	return __AXInterfaceCursorSetAndReturnSeedSymbol
}

var __AXInterfaceGetBristolEnabledSymbol uintptr

// _AXInterfaceGetBristolEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceGetBristolEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceGetBristolEnabled
func _AXInterfaceGetBristolEnabled() {
	panic("hiservices: symbol _AXInterfaceGetBristolEnabled has opaque signature; use _AXInterfaceGetBristolEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceGetBristolEnabledSymbol returns the raw symbol address for _AXInterfaceGetBristolEnabled.
func _AXInterfaceGetBristolEnabledSymbol() uintptr {
	if __AXInterfaceGetBristolEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceGetBristolEnabledSymbol
}

var __AXInterfaceGetClassicInvertColorEnabledSymbol uintptr

// _AXInterfaceGetClassicInvertColorEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceGetClassicInvertColorEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceGetClassicInvertColorEnabled
func _AXInterfaceGetClassicInvertColorEnabled() {
	panic("hiservices: symbol _AXInterfaceGetClassicInvertColorEnabled has opaque signature; use _AXInterfaceGetClassicInvertColorEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceGetClassicInvertColorEnabledSymbol returns the raw symbol address for _AXInterfaceGetClassicInvertColorEnabled.
func _AXInterfaceGetClassicInvertColorEnabledSymbol() uintptr {
	if __AXInterfaceGetClassicInvertColorEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceGetClassicInvertColorEnabledSymbol
}

var __AXInterfaceGetDifferentiateWithoutColorEnabledSymbol uintptr

// _AXInterfaceGetDifferentiateWithoutColorEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceGetDifferentiateWithoutColorEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceGetDifferentiateWithoutColorEnabled
func _AXInterfaceGetDifferentiateWithoutColorEnabled() {
	panic("hiservices: symbol _AXInterfaceGetDifferentiateWithoutColorEnabled has opaque signature; use _AXInterfaceGetDifferentiateWithoutColorEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceGetDifferentiateWithoutColorEnabledSymbol returns the raw symbol address for _AXInterfaceGetDifferentiateWithoutColorEnabled.
func _AXInterfaceGetDifferentiateWithoutColorEnabledSymbol() uintptr {
	if __AXInterfaceGetDifferentiateWithoutColorEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceGetDifferentiateWithoutColorEnabledSymbol
}

var __AXInterfaceGetIncreaseContrastEnabledSymbol uintptr

// _AXInterfaceGetIncreaseContrastEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceGetIncreaseContrastEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceGetIncreaseContrastEnabled
func _AXInterfaceGetIncreaseContrastEnabled() {
	panic("hiservices: symbol _AXInterfaceGetIncreaseContrastEnabled has opaque signature; use _AXInterfaceGetIncreaseContrastEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceGetIncreaseContrastEnabledSymbol returns the raw symbol address for _AXInterfaceGetIncreaseContrastEnabled.
func _AXInterfaceGetIncreaseContrastEnabledSymbol() uintptr {
	if __AXInterfaceGetIncreaseContrastEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceGetIncreaseContrastEnabledSymbol
}

var __AXInterfaceGetReduceMotionEnabledSymbol uintptr

// _AXInterfaceGetReduceMotionEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceGetReduceMotionEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceGetReduceMotionEnabled
func _AXInterfaceGetReduceMotionEnabled() {
	panic("hiservices: symbol _AXInterfaceGetReduceMotionEnabled has opaque signature; use _AXInterfaceGetReduceMotionEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceGetReduceMotionEnabledSymbol returns the raw symbol address for _AXInterfaceGetReduceMotionEnabled.
func _AXInterfaceGetReduceMotionEnabledSymbol() uintptr {
	if __AXInterfaceGetReduceMotionEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceGetReduceMotionEnabledSymbol
}

var __AXInterfaceGetReduceTextInsertionPointModulationEnabledSymbol uintptr

// _AXInterfaceGetReduceTextInsertionPointModulationEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceGetReduceTextInsertionPointModulationEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceGetReduceTextInsertionPointModulationEnabled
func _AXInterfaceGetReduceTextInsertionPointModulationEnabled() {
	panic("hiservices: symbol _AXInterfaceGetReduceTextInsertionPointModulationEnabled has opaque signature; use _AXInterfaceGetReduceTextInsertionPointModulationEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceGetReduceTextInsertionPointModulationEnabledSymbol returns the raw symbol address for _AXInterfaceGetReduceTextInsertionPointModulationEnabled.
func _AXInterfaceGetReduceTextInsertionPointModulationEnabledSymbol() uintptr {
	if __AXInterfaceGetReduceTextInsertionPointModulationEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceGetReduceTextInsertionPointModulationEnabledSymbol
}

var __AXInterfaceGetReduceTransparencyEnabledSymbol uintptr

// _AXInterfaceGetReduceTransparencyEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceGetReduceTransparencyEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceGetReduceTransparencyEnabled
func _AXInterfaceGetReduceTransparencyEnabled() {
	panic("hiservices: symbol _AXInterfaceGetReduceTransparencyEnabled has opaque signature; use _AXInterfaceGetReduceTransparencyEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceGetReduceTransparencyEnabledSymbol returns the raw symbol address for _AXInterfaceGetReduceTransparencyEnabled.
func _AXInterfaceGetReduceTransparencyEnabledSymbol() uintptr {
	if __AXInterfaceGetReduceTransparencyEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceGetReduceTransparencyEnabledSymbol
}

var __AXInterfaceGetRichmondEnabledSymbol uintptr

// _AXInterfaceGetRichmondEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceGetRichmondEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceGetRichmondEnabled
func _AXInterfaceGetRichmondEnabled() {
	panic("hiservices: symbol _AXInterfaceGetRichmondEnabled has opaque signature; use _AXInterfaceGetRichmondEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceGetRichmondEnabledSymbol returns the raw symbol address for _AXInterfaceGetRichmondEnabled.
func _AXInterfaceGetRichmondEnabledSymbol() uintptr {
	if __AXInterfaceGetRichmondEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceGetRichmondEnabledSymbol
}

var __AXInterfaceGetShowToolbarButtonShapesEnabledSymbol uintptr

// _AXInterfaceGetShowToolbarButtonShapesEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceGetShowToolbarButtonShapesEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceGetShowToolbarButtonShapesEnabled
func _AXInterfaceGetShowToolbarButtonShapesEnabled() {
	panic("hiservices: symbol _AXInterfaceGetShowToolbarButtonShapesEnabled has opaque signature; use _AXInterfaceGetShowToolbarButtonShapesEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceGetShowToolbarButtonShapesEnabledSymbol returns the raw symbol address for _AXInterfaceGetShowToolbarButtonShapesEnabled.
func _AXInterfaceGetShowToolbarButtonShapesEnabledSymbol() uintptr {
	if __AXInterfaceGetShowToolbarButtonShapesEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceGetShowToolbarButtonShapesEnabledSymbol
}

var __AXInterfaceGetShowWindowTitlebarIconsEnabledSymbol uintptr

// _AXInterfaceGetShowWindowTitlebarIconsEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceGetShowWindowTitlebarIconsEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceGetShowWindowTitlebarIconsEnabled
func _AXInterfaceGetShowWindowTitlebarIconsEnabled() {
	panic("hiservices: symbol _AXInterfaceGetShowWindowTitlebarIconsEnabled has opaque signature; use _AXInterfaceGetShowWindowTitlebarIconsEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceGetShowWindowTitlebarIconsEnabledSymbol returns the raw symbol address for _AXInterfaceGetShowWindowTitlebarIconsEnabled.
func _AXInterfaceGetShowWindowTitlebarIconsEnabledSymbol() uintptr {
	if __AXInterfaceGetShowWindowTitlebarIconsEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceGetShowWindowTitlebarIconsEnabledSymbol
}

var __AXInterfaceSetClassicInvertColorEnabledSymbol uintptr

// _AXInterfaceSetClassicInvertColorEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetClassicInvertColorEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetClassicInvertColorEnabled
func _AXInterfaceSetClassicInvertColorEnabled() {
	panic("hiservices: symbol _AXInterfaceSetClassicInvertColorEnabled has opaque signature; use _AXInterfaceSetClassicInvertColorEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetClassicInvertColorEnabledSymbol returns the raw symbol address for _AXInterfaceSetClassicInvertColorEnabled.
func _AXInterfaceSetClassicInvertColorEnabledSymbol() uintptr {
	if __AXInterfaceSetClassicInvertColorEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetClassicInvertColorEnabledSymbol
}

var __AXInterfaceSetCursorColorFillSymbol uintptr

// _AXInterfaceSetCursorColorFill has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetCursorColorFillSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetCursorColorFill
func _AXInterfaceSetCursorColorFill() {
	panic("hiservices: symbol _AXInterfaceSetCursorColorFill has opaque signature; use _AXInterfaceSetCursorColorFillSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetCursorColorFillSymbol returns the raw symbol address for _AXInterfaceSetCursorColorFill.
func _AXInterfaceSetCursorColorFillSymbol() uintptr {
	if __AXInterfaceSetCursorColorFillSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetCursorColorFillSymbol
}

var __AXInterfaceSetCursorColorOutlineSymbol uintptr

// _AXInterfaceSetCursorColorOutline has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetCursorColorOutlineSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetCursorColorOutline
func _AXInterfaceSetCursorColorOutline() {
	panic("hiservices: symbol _AXInterfaceSetCursorColorOutline has opaque signature; use _AXInterfaceSetCursorColorOutlineSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetCursorColorOutlineSymbol returns the raw symbol address for _AXInterfaceSetCursorColorOutline.
func _AXInterfaceSetCursorColorOutlineSymbol() uintptr {
	if __AXInterfaceSetCursorColorOutlineSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetCursorColorOutlineSymbol
}

var __AXInterfaceSetCursorIsOverriddenSymbol uintptr

// _AXInterfaceSetCursorIsOverridden has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetCursorIsOverriddenSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetCursorIsOverridden
func _AXInterfaceSetCursorIsOverridden() {
	panic("hiservices: symbol _AXInterfaceSetCursorIsOverridden has opaque signature; use _AXInterfaceSetCursorIsOverriddenSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetCursorIsOverriddenSymbol returns the raw symbol address for _AXInterfaceSetCursorIsOverridden.
func _AXInterfaceSetCursorIsOverriddenSymbol() uintptr {
	if __AXInterfaceSetCursorIsOverriddenSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetCursorIsOverriddenSymbol
}

var __AXInterfaceSetDifferentiateWithoutColorEnabledSymbol uintptr

// _AXInterfaceSetDifferentiateWithoutColorEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetDifferentiateWithoutColorEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetDifferentiateWithoutColorEnabled
func _AXInterfaceSetDifferentiateWithoutColorEnabled() {
	panic("hiservices: symbol _AXInterfaceSetDifferentiateWithoutColorEnabled has opaque signature; use _AXInterfaceSetDifferentiateWithoutColorEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetDifferentiateWithoutColorEnabledSymbol returns the raw symbol address for _AXInterfaceSetDifferentiateWithoutColorEnabled.
func _AXInterfaceSetDifferentiateWithoutColorEnabledSymbol() uintptr {
	if __AXInterfaceSetDifferentiateWithoutColorEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetDifferentiateWithoutColorEnabledSymbol
}

var __AXInterfaceSetDifferentiateWithoutColorEnabledOverrideSymbol uintptr

// _AXInterfaceSetDifferentiateWithoutColorEnabledOverride has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetDifferentiateWithoutColorEnabledOverrideSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetDifferentiateWithoutColorEnabledOverride
func _AXInterfaceSetDifferentiateWithoutColorEnabledOverride() {
	panic("hiservices: symbol _AXInterfaceSetDifferentiateWithoutColorEnabledOverride has opaque signature; use _AXInterfaceSetDifferentiateWithoutColorEnabledOverrideSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetDifferentiateWithoutColorEnabledOverrideSymbol returns the raw symbol address for _AXInterfaceSetDifferentiateWithoutColorEnabledOverride.
func _AXInterfaceSetDifferentiateWithoutColorEnabledOverrideSymbol() uintptr {
	if __AXInterfaceSetDifferentiateWithoutColorEnabledOverrideSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetDifferentiateWithoutColorEnabledOverrideSymbol
}

var __AXInterfaceSetIncreaseContrastEnabledSymbol uintptr

// _AXInterfaceSetIncreaseContrastEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetIncreaseContrastEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetIncreaseContrastEnabled
func _AXInterfaceSetIncreaseContrastEnabled() {
	panic("hiservices: symbol _AXInterfaceSetIncreaseContrastEnabled has opaque signature; use _AXInterfaceSetIncreaseContrastEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetIncreaseContrastEnabledSymbol returns the raw symbol address for _AXInterfaceSetIncreaseContrastEnabled.
func _AXInterfaceSetIncreaseContrastEnabledSymbol() uintptr {
	if __AXInterfaceSetIncreaseContrastEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetIncreaseContrastEnabledSymbol
}

var __AXInterfaceSetIncreaseContrastEnabledOverrideSymbol uintptr

// _AXInterfaceSetIncreaseContrastEnabledOverride has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetIncreaseContrastEnabledOverrideSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetIncreaseContrastEnabledOverride
func _AXInterfaceSetIncreaseContrastEnabledOverride() {
	panic("hiservices: symbol _AXInterfaceSetIncreaseContrastEnabledOverride has opaque signature; use _AXInterfaceSetIncreaseContrastEnabledOverrideSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetIncreaseContrastEnabledOverrideSymbol returns the raw symbol address for _AXInterfaceSetIncreaseContrastEnabledOverride.
func _AXInterfaceSetIncreaseContrastEnabledOverrideSymbol() uintptr {
	if __AXInterfaceSetIncreaseContrastEnabledOverrideSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetIncreaseContrastEnabledOverrideSymbol
}

var __AXInterfaceSetReduceMotionEnabledSymbol uintptr

// _AXInterfaceSetReduceMotionEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetReduceMotionEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetReduceMotionEnabled
func _AXInterfaceSetReduceMotionEnabled() {
	panic("hiservices: symbol _AXInterfaceSetReduceMotionEnabled has opaque signature; use _AXInterfaceSetReduceMotionEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetReduceMotionEnabledSymbol returns the raw symbol address for _AXInterfaceSetReduceMotionEnabled.
func _AXInterfaceSetReduceMotionEnabledSymbol() uintptr {
	if __AXInterfaceSetReduceMotionEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetReduceMotionEnabledSymbol
}

var __AXInterfaceSetReduceMotionEnabledOverrideSymbol uintptr

// _AXInterfaceSetReduceMotionEnabledOverride has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetReduceMotionEnabledOverrideSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetReduceMotionEnabledOverride
func _AXInterfaceSetReduceMotionEnabledOverride() {
	panic("hiservices: symbol _AXInterfaceSetReduceMotionEnabledOverride has opaque signature; use _AXInterfaceSetReduceMotionEnabledOverrideSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetReduceMotionEnabledOverrideSymbol returns the raw symbol address for _AXInterfaceSetReduceMotionEnabledOverride.
func _AXInterfaceSetReduceMotionEnabledOverrideSymbol() uintptr {
	if __AXInterfaceSetReduceMotionEnabledOverrideSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetReduceMotionEnabledOverrideSymbol
}

var __AXInterfaceSetReduceTextInsertionPointModulationEnabledSymbol uintptr

// _AXInterfaceSetReduceTextInsertionPointModulationEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetReduceTextInsertionPointModulationEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetReduceTextInsertionPointModulationEnabled
func _AXInterfaceSetReduceTextInsertionPointModulationEnabled() {
	panic("hiservices: symbol _AXInterfaceSetReduceTextInsertionPointModulationEnabled has opaque signature; use _AXInterfaceSetReduceTextInsertionPointModulationEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetReduceTextInsertionPointModulationEnabledSymbol returns the raw symbol address for _AXInterfaceSetReduceTextInsertionPointModulationEnabled.
func _AXInterfaceSetReduceTextInsertionPointModulationEnabledSymbol() uintptr {
	if __AXInterfaceSetReduceTextInsertionPointModulationEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetReduceTextInsertionPointModulationEnabledSymbol
}

var __AXInterfaceSetReduceTextInsertionPointModulationEnabledOverrideSymbol uintptr

// _AXInterfaceSetReduceTextInsertionPointModulationEnabledOverride has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetReduceTextInsertionPointModulationEnabledOverrideSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetReduceTextInsertionPointModulationEnabledOverride
func _AXInterfaceSetReduceTextInsertionPointModulationEnabledOverride() {
	panic("hiservices: symbol _AXInterfaceSetReduceTextInsertionPointModulationEnabledOverride has opaque signature; use _AXInterfaceSetReduceTextInsertionPointModulationEnabledOverrideSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetReduceTextInsertionPointModulationEnabledOverrideSymbol returns the raw symbol address for _AXInterfaceSetReduceTextInsertionPointModulationEnabledOverride.
func _AXInterfaceSetReduceTextInsertionPointModulationEnabledOverrideSymbol() uintptr {
	if __AXInterfaceSetReduceTextInsertionPointModulationEnabledOverrideSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetReduceTextInsertionPointModulationEnabledOverrideSymbol
}

var __AXInterfaceSetReduceTransparencyEnabledSymbol uintptr

// _AXInterfaceSetReduceTransparencyEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetReduceTransparencyEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetReduceTransparencyEnabled
func _AXInterfaceSetReduceTransparencyEnabled() {
	panic("hiservices: symbol _AXInterfaceSetReduceTransparencyEnabled has opaque signature; use _AXInterfaceSetReduceTransparencyEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetReduceTransparencyEnabledSymbol returns the raw symbol address for _AXInterfaceSetReduceTransparencyEnabled.
func _AXInterfaceSetReduceTransparencyEnabledSymbol() uintptr {
	if __AXInterfaceSetReduceTransparencyEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetReduceTransparencyEnabledSymbol
}

var __AXInterfaceSetReduceTransparencyEnabledOverrideSymbol uintptr

// _AXInterfaceSetReduceTransparencyEnabledOverride has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetReduceTransparencyEnabledOverrideSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetReduceTransparencyEnabledOverride
func _AXInterfaceSetReduceTransparencyEnabledOverride() {
	panic("hiservices: symbol _AXInterfaceSetReduceTransparencyEnabledOverride has opaque signature; use _AXInterfaceSetReduceTransparencyEnabledOverrideSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetReduceTransparencyEnabledOverrideSymbol returns the raw symbol address for _AXInterfaceSetReduceTransparencyEnabledOverride.
func _AXInterfaceSetReduceTransparencyEnabledOverrideSymbol() uintptr {
	if __AXInterfaceSetReduceTransparencyEnabledOverrideSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetReduceTransparencyEnabledOverrideSymbol
}

var __AXInterfaceSetShowToolbarButtonShapesEnabledSymbol uintptr

// _AXInterfaceSetShowToolbarButtonShapesEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetShowToolbarButtonShapesEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetShowToolbarButtonShapesEnabled
func _AXInterfaceSetShowToolbarButtonShapesEnabled() {
	panic("hiservices: symbol _AXInterfaceSetShowToolbarButtonShapesEnabled has opaque signature; use _AXInterfaceSetShowToolbarButtonShapesEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetShowToolbarButtonShapesEnabledSymbol returns the raw symbol address for _AXInterfaceSetShowToolbarButtonShapesEnabled.
func _AXInterfaceSetShowToolbarButtonShapesEnabledSymbol() uintptr {
	if __AXInterfaceSetShowToolbarButtonShapesEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetShowToolbarButtonShapesEnabledSymbol
}

var __AXInterfaceSetShowToolbarButtonShapesEnabledOverrideSymbol uintptr

// _AXInterfaceSetShowToolbarButtonShapesEnabledOverride has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetShowToolbarButtonShapesEnabledOverrideSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetShowToolbarButtonShapesEnabledOverride
func _AXInterfaceSetShowToolbarButtonShapesEnabledOverride() {
	panic("hiservices: symbol _AXInterfaceSetShowToolbarButtonShapesEnabledOverride has opaque signature; use _AXInterfaceSetShowToolbarButtonShapesEnabledOverrideSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetShowToolbarButtonShapesEnabledOverrideSymbol returns the raw symbol address for _AXInterfaceSetShowToolbarButtonShapesEnabledOverride.
func _AXInterfaceSetShowToolbarButtonShapesEnabledOverrideSymbol() uintptr {
	if __AXInterfaceSetShowToolbarButtonShapesEnabledOverrideSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetShowToolbarButtonShapesEnabledOverrideSymbol
}

var __AXInterfaceSetShowWindowTitlebarIconsEnabledSymbol uintptr

// _AXInterfaceSetShowWindowTitlebarIconsEnabled has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetShowWindowTitlebarIconsEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetShowWindowTitlebarIconsEnabled
func _AXInterfaceSetShowWindowTitlebarIconsEnabled() {
	panic("hiservices: symbol _AXInterfaceSetShowWindowTitlebarIconsEnabled has opaque signature; use _AXInterfaceSetShowWindowTitlebarIconsEnabledSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetShowWindowTitlebarIconsEnabledSymbol returns the raw symbol address for _AXInterfaceSetShowWindowTitlebarIconsEnabled.
func _AXInterfaceSetShowWindowTitlebarIconsEnabledSymbol() uintptr {
	if __AXInterfaceSetShowWindowTitlebarIconsEnabledSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetShowWindowTitlebarIconsEnabledSymbol
}

var __AXInterfaceSetShowWindowTitlebarIconsEnabledOverrideSymbol uintptr

// _AXInterfaceSetShowWindowTitlebarIconsEnabledOverride has an opaque C signature in discovered metadata.
// Call _AXInterfaceSetShowWindowTitlebarIconsEnabledOverrideSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXInterfaceSetShowWindowTitlebarIconsEnabledOverride
func _AXInterfaceSetShowWindowTitlebarIconsEnabledOverride() {
	panic("hiservices: symbol _AXInterfaceSetShowWindowTitlebarIconsEnabledOverride has opaque signature; use _AXInterfaceSetShowWindowTitlebarIconsEnabledOverrideSymbol() and a typed manual wrapper")
}

// _AXInterfaceSetShowWindowTitlebarIconsEnabledOverrideSymbol returns the raw symbol address for _AXInterfaceSetShowWindowTitlebarIconsEnabledOverride.
func _AXInterfaceSetShowWindowTitlebarIconsEnabledOverrideSymbol() uintptr {
	if __AXInterfaceSetShowWindowTitlebarIconsEnabledOverrideSymbol == 0 {
		return 0
	}
	return __AXInterfaceSetShowWindowTitlebarIconsEnabledOverrideSymbol
}

var __AXIsAppleClientForCurrentRequestUntrustedSymbol uintptr

// _AXIsAppleClientForCurrentRequestUntrusted has an opaque C signature in discovered metadata.
// Call _AXIsAppleClientForCurrentRequestUntrustedSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXIsAppleClientForCurrentRequestUntrusted
func _AXIsAppleClientForCurrentRequestUntrusted() {
	panic("hiservices: symbol _AXIsAppleClientForCurrentRequestUntrusted has opaque signature; use _AXIsAppleClientForCurrentRequestUntrustedSymbol() and a typed manual wrapper")
}

// _AXIsAppleClientForCurrentRequestUntrustedSymbol returns the raw symbol address for _AXIsAppleClientForCurrentRequestUntrusted.
func _AXIsAppleClientForCurrentRequestUntrustedSymbol() uintptr {
	if __AXIsAppleClientForCurrentRequestUntrustedSymbol == 0 {
		return 0
	}
	return __AXIsAppleClientForCurrentRequestUntrustedSymbol
}

var __AXNotificationHandlerCreateWithCallbackSymbol uintptr

// _AXNotificationHandlerCreateWithCallback has an opaque C signature in discovered metadata.
// Call _AXNotificationHandlerCreateWithCallbackSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXNotificationHandlerCreateWithCallback
func _AXNotificationHandlerCreateWithCallback() {
	panic("hiservices: symbol _AXNotificationHandlerCreateWithCallback has opaque signature; use _AXNotificationHandlerCreateWithCallbackSymbol() and a typed manual wrapper")
}

// _AXNotificationHandlerCreateWithCallbackSymbol returns the raw symbol address for _AXNotificationHandlerCreateWithCallback.
func _AXNotificationHandlerCreateWithCallbackSymbol() uintptr {
	if __AXNotificationHandlerCreateWithCallbackSymbol == 0 {
		return 0
	}
	return __AXNotificationHandlerCreateWithCallbackSymbol
}

var __AXRegisterControlComputerAccessSymbol uintptr

// _AXRegisterControlComputerAccess has an opaque C signature in discovered metadata.
// Call _AXRegisterControlComputerAccessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXRegisterControlComputerAccess
func _AXRegisterControlComputerAccess() {
	panic("hiservices: symbol _AXRegisterControlComputerAccess has opaque signature; use _AXRegisterControlComputerAccessSymbol() and a typed manual wrapper")
}

// _AXRegisterControlComputerAccessSymbol returns the raw symbol address for _AXRegisterControlComputerAccess.
func _AXRegisterControlComputerAccessSymbol() uintptr {
	if __AXRegisterControlComputerAccessSymbol == 0 {
		return 0
	}
	return __AXRegisterControlComputerAccessSymbol
}

var __AXSetAuditTokenIsAuthenticatedCallbackSymbol uintptr

// _AXSetAuditTokenIsAuthenticatedCallback has an opaque C signature in discovered metadata.
// Call _AXSetAuditTokenIsAuthenticatedCallbackSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXSetAuditTokenIsAuthenticatedCallback
func _AXSetAuditTokenIsAuthenticatedCallback() {
	panic("hiservices: symbol _AXSetAuditTokenIsAuthenticatedCallback has opaque signature; use _AXSetAuditTokenIsAuthenticatedCallbackSymbol() and a typed manual wrapper")
}

// _AXSetAuditTokenIsAuthenticatedCallbackSymbol returns the raw symbol address for _AXSetAuditTokenIsAuthenticatedCallback.
func _AXSetAuditTokenIsAuthenticatedCallbackSymbol() uintptr {
	if __AXSetAuditTokenIsAuthenticatedCallbackSymbol == 0 {
		return 0
	}
	return __AXSetAuditTokenIsAuthenticatedCallbackSymbol
}

var __AXSetClientIdentificationOverrideSymbol uintptr

// _AXSetClientIdentificationOverride has an opaque C signature in discovered metadata.
// Call _AXSetClientIdentificationOverrideSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXSetClientIdentificationOverride
func _AXSetClientIdentificationOverride() {
	panic("hiservices: symbol _AXSetClientIdentificationOverride has opaque signature; use _AXSetClientIdentificationOverrideSymbol() and a typed manual wrapper")
}

// _AXSetClientIdentificationOverrideSymbol returns the raw symbol address for _AXSetClientIdentificationOverride.
func _AXSetClientIdentificationOverrideSymbol() uintptr {
	if __AXSetClientIdentificationOverrideSymbol == 0 {
		return 0
	}
	return __AXSetClientIdentificationOverrideSymbol
}

var __AXShouldElementBeIgnoredForNavigationSymbol uintptr

// _AXShouldElementBeIgnoredForNavigation has an opaque C signature in discovered metadata.
// Call _AXShouldElementBeIgnoredForNavigationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXShouldElementBeIgnoredForNavigation
func _AXShouldElementBeIgnoredForNavigation() {
	panic("hiservices: symbol _AXShouldElementBeIgnoredForNavigation has opaque signature; use _AXShouldElementBeIgnoredForNavigationSymbol() and a typed manual wrapper")
}

// _AXShouldElementBeIgnoredForNavigationSymbol returns the raw symbol address for _AXShouldElementBeIgnoredForNavigation.
func _AXShouldElementBeIgnoredForNavigationSymbol() uintptr {
	if __AXShouldElementBeIgnoredForNavigationSymbol == 0 {
		return 0
	}
	return __AXShouldElementBeIgnoredForNavigationSymbol
}

var __AXUIElementCopyElementAtPositionIncludeIgnoredSymbol uintptr

// _AXUIElementCopyElementAtPositionIncludeIgnored has an opaque C signature in discovered metadata.
// Call _AXUIElementCopyElementAtPositionIncludeIgnoredSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementCopyElementAtPositionIncludeIgnored
func _AXUIElementCopyElementAtPositionIncludeIgnored() {
	panic("hiservices: symbol _AXUIElementCopyElementAtPositionIncludeIgnored has opaque signature; use _AXUIElementCopyElementAtPositionIncludeIgnoredSymbol() and a typed manual wrapper")
}

// _AXUIElementCopyElementAtPositionIncludeIgnoredSymbol returns the raw symbol address for _AXUIElementCopyElementAtPositionIncludeIgnored.
func _AXUIElementCopyElementAtPositionIncludeIgnoredSymbol() uintptr {
	if __AXUIElementCopyElementAtPositionIncludeIgnoredSymbol == 0 {
		return 0
	}
	return __AXUIElementCopyElementAtPositionIncludeIgnoredSymbol
}

var __AXUIElementCreateApplicationWithPresenterPidSymbol uintptr

// _AXUIElementCreateApplicationWithPresenterPid has an opaque C signature in discovered metadata.
// Call _AXUIElementCreateApplicationWithPresenterPidSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementCreateApplicationWithPresenterPid
func _AXUIElementCreateApplicationWithPresenterPid() {
	panic("hiservices: symbol _AXUIElementCreateApplicationWithPresenterPid has opaque signature; use _AXUIElementCreateApplicationWithPresenterPidSymbol() and a typed manual wrapper")
}

// _AXUIElementCreateApplicationWithPresenterPidSymbol returns the raw symbol address for _AXUIElementCreateApplicationWithPresenterPid.
func _AXUIElementCreateApplicationWithPresenterPidSymbol() uintptr {
	if __AXUIElementCreateApplicationWithPresenterPidSymbol == 0 {
		return 0
	}
	return __AXUIElementCreateApplicationWithPresenterPidSymbol
}

var __AXUIElementCreateWithDataSymbol uintptr

// _AXUIElementCreateWithData has an opaque C signature in discovered metadata.
// Call _AXUIElementCreateWithDataSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementCreateWithData
func _AXUIElementCreateWithData() {
	panic("hiservices: symbol _AXUIElementCreateWithData has opaque signature; use _AXUIElementCreateWithDataSymbol() and a typed manual wrapper")
}

// _AXUIElementCreateWithDataSymbol returns the raw symbol address for _AXUIElementCreateWithData.
func _AXUIElementCreateWithDataSymbol() uintptr {
	if __AXUIElementCreateWithDataSymbol == 0 {
		return 0
	}
	return __AXUIElementCreateWithDataSymbol
}

var __AXUIElementCreateWithDataAndPidSymbol uintptr

// _AXUIElementCreateWithDataAndPid has an opaque C signature in discovered metadata.
// Call _AXUIElementCreateWithDataAndPidSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementCreateWithDataAndPid
func _AXUIElementCreateWithDataAndPid() {
	panic("hiservices: symbol _AXUIElementCreateWithDataAndPid has opaque signature; use _AXUIElementCreateWithDataAndPidSymbol() and a typed manual wrapper")
}

// _AXUIElementCreateWithDataAndPidSymbol returns the raw symbol address for _AXUIElementCreateWithDataAndPid.
func _AXUIElementCreateWithDataAndPidSymbol() uintptr {
	if __AXUIElementCreateWithDataAndPidSymbol == 0 {
		return 0
	}
	return __AXUIElementCreateWithDataAndPidSymbol
}

var __AXUIElementCreateWithDataAndPresenterPidSymbol uintptr

// _AXUIElementCreateWithDataAndPresenterPid has an opaque C signature in discovered metadata.
// Call _AXUIElementCreateWithDataAndPresenterPidSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementCreateWithDataAndPresenterPid
func _AXUIElementCreateWithDataAndPresenterPid() {
	panic("hiservices: symbol _AXUIElementCreateWithDataAndPresenterPid has opaque signature; use _AXUIElementCreateWithDataAndPresenterPidSymbol() and a typed manual wrapper")
}

// _AXUIElementCreateWithDataAndPresenterPidSymbol returns the raw symbol address for _AXUIElementCreateWithDataAndPresenterPid.
func _AXUIElementCreateWithDataAndPresenterPidSymbol() uintptr {
	if __AXUIElementCreateWithDataAndPresenterPidSymbol == 0 {
		return 0
	}
	return __AXUIElementCreateWithDataAndPresenterPidSymbol
}

var __AXUIElementCreateWithPtrSymbol uintptr

// _AXUIElementCreateWithPtr has an opaque C signature in discovered metadata.
// Call _AXUIElementCreateWithPtrSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementCreateWithPtr
func _AXUIElementCreateWithPtr() {
	panic("hiservices: symbol _AXUIElementCreateWithPtr has opaque signature; use _AXUIElementCreateWithPtrSymbol() and a typed manual wrapper")
}

// _AXUIElementCreateWithPtrSymbol returns the raw symbol address for _AXUIElementCreateWithPtr.
func _AXUIElementCreateWithPtrSymbol() uintptr {
	if __AXUIElementCreateWithPtrSymbol == 0 {
		return 0
	}
	return __AXUIElementCreateWithPtrSymbol
}

var __AXUIElementCreateWithRemoteTokenSymbol uintptr

// _AXUIElementCreateWithRemoteToken has an opaque C signature in discovered metadata.
// Call _AXUIElementCreateWithRemoteTokenSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementCreateWithRemoteToken
func _AXUIElementCreateWithRemoteToken() {
	panic("hiservices: symbol _AXUIElementCreateWithRemoteToken has opaque signature; use _AXUIElementCreateWithRemoteTokenSymbol() and a typed manual wrapper")
}

// _AXUIElementCreateWithRemoteTokenSymbol returns the raw symbol address for _AXUIElementCreateWithRemoteToken.
func _AXUIElementCreateWithRemoteTokenSymbol() uintptr {
	if __AXUIElementCreateWithRemoteTokenSymbol == 0 {
		return 0
	}
	return __AXUIElementCreateWithRemoteTokenSymbol
}

var __AXUIElementGetActualPidSymbol uintptr

// _AXUIElementGetActualPid has an opaque C signature in discovered metadata.
// Call _AXUIElementGetActualPidSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementGetActualPid
func _AXUIElementGetActualPid() {
	panic("hiservices: symbol _AXUIElementGetActualPid has opaque signature; use _AXUIElementGetActualPidSymbol() and a typed manual wrapper")
}

// _AXUIElementGetActualPidSymbol returns the raw symbol address for _AXUIElementGetActualPid.
func _AXUIElementGetActualPidSymbol() uintptr {
	if __AXUIElementGetActualPidSymbol == 0 {
		return 0
	}
	return __AXUIElementGetActualPidSymbol
}

var __AXUIElementGetDataSymbol uintptr

// _AXUIElementGetData has an opaque C signature in discovered metadata.
// Call _AXUIElementGetDataSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementGetData
func _AXUIElementGetData() {
	panic("hiservices: symbol _AXUIElementGetData has opaque signature; use _AXUIElementGetDataSymbol() and a typed manual wrapper")
}

// _AXUIElementGetDataSymbol returns the raw symbol address for _AXUIElementGetData.
func _AXUIElementGetDataSymbol() uintptr {
	if __AXUIElementGetDataSymbol == 0 {
		return 0
	}
	return __AXUIElementGetDataSymbol
}

var __AXUIElementGetIsProcessSuspendedSymbol uintptr

// _AXUIElementGetIsProcessSuspended has an opaque C signature in discovered metadata.
// Call _AXUIElementGetIsProcessSuspendedSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementGetIsProcessSuspended
func _AXUIElementGetIsProcessSuspended() {
	panic("hiservices: symbol _AXUIElementGetIsProcessSuspended has opaque signature; use _AXUIElementGetIsProcessSuspendedSymbol() and a typed manual wrapper")
}

// _AXUIElementGetIsProcessSuspendedSymbol returns the raw symbol address for _AXUIElementGetIsProcessSuspended.
func _AXUIElementGetIsProcessSuspendedSymbol() uintptr {
	if __AXUIElementGetIsProcessSuspendedSymbol == 0 {
		return 0
	}
	return __AXUIElementGetIsProcessSuspendedSymbol
}

var __AXUIElementGetWindowSymbol uintptr

// _AXUIElementGetWindow has an opaque C signature in discovered metadata.
// Call _AXUIElementGetWindowSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementGetWindow
func _AXUIElementGetWindow() {
	panic("hiservices: symbol _AXUIElementGetWindow has opaque signature; use _AXUIElementGetWindowSymbol() and a typed manual wrapper")
}

// _AXUIElementGetWindowSymbol returns the raw symbol address for _AXUIElementGetWindow.
func _AXUIElementGetWindowSymbol() uintptr {
	if __AXUIElementGetWindowSymbol == 0 {
		return 0
	}
	return __AXUIElementGetWindowSymbol
}

var __AXUIElementLoadAccessibilityBundlesSymbol uintptr

// _AXUIElementLoadAccessibilityBundles has an opaque C signature in discovered metadata.
// Call _AXUIElementLoadAccessibilityBundlesSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementLoadAccessibilityBundles
func _AXUIElementLoadAccessibilityBundles() {
	panic("hiservices: symbol _AXUIElementLoadAccessibilityBundles has opaque signature; use _AXUIElementLoadAccessibilityBundlesSymbol() and a typed manual wrapper")
}

// _AXUIElementLoadAccessibilityBundlesSymbol returns the raw symbol address for _AXUIElementLoadAccessibilityBundles.
func _AXUIElementLoadAccessibilityBundlesSymbol() uintptr {
	if __AXUIElementLoadAccessibilityBundlesSymbol == 0 {
		return 0
	}
	return __AXUIElementLoadAccessibilityBundlesSymbol
}

var __AXUIElementNotifyProcessSuspendStatusSymbol uintptr

// _AXUIElementNotifyProcessSuspendStatus has an opaque C signature in discovered metadata.
// Call _AXUIElementNotifyProcessSuspendStatusSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementNotifyProcessSuspendStatus
func _AXUIElementNotifyProcessSuspendStatus() {
	panic("hiservices: symbol _AXUIElementNotifyProcessSuspendStatus has opaque signature; use _AXUIElementNotifyProcessSuspendStatusSymbol() and a typed manual wrapper")
}

// _AXUIElementNotifyProcessSuspendStatusSymbol returns the raw symbol address for _AXUIElementNotifyProcessSuspendStatus.
func _AXUIElementNotifyProcessSuspendStatusSymbol() uintptr {
	if __AXUIElementNotifyProcessSuspendStatusSymbol == 0 {
		return 0
	}
	return __AXUIElementNotifyProcessSuspendStatusSymbol
}

var __AXUIElementPostNotificationSymbol uintptr

// _AXUIElementPostNotification has an opaque C signature in discovered metadata.
// Call _AXUIElementPostNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementPostNotification
func _AXUIElementPostNotification() {
	panic("hiservices: symbol _AXUIElementPostNotification has opaque signature; use _AXUIElementPostNotificationSymbol() and a typed manual wrapper")
}

// _AXUIElementPostNotificationSymbol returns the raw symbol address for _AXUIElementPostNotification.
func _AXUIElementPostNotificationSymbol() uintptr {
	if __AXUIElementPostNotificationSymbol == 0 {
		return 0
	}
	return __AXUIElementPostNotificationSymbol
}

var __AXUIElementPostNotificationForObservedElementSymbol uintptr

// _AXUIElementPostNotificationForObservedElement has an opaque C signature in discovered metadata.
// Call _AXUIElementPostNotificationForObservedElementSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementPostNotificationForObservedElement
func _AXUIElementPostNotificationForObservedElement() {
	panic("hiservices: symbol _AXUIElementPostNotificationForObservedElement has opaque signature; use _AXUIElementPostNotificationForObservedElementSymbol() and a typed manual wrapper")
}

// _AXUIElementPostNotificationForObservedElementSymbol returns the raw symbol address for _AXUIElementPostNotificationForObservedElement.
func _AXUIElementPostNotificationForObservedElementSymbol() uintptr {
	if __AXUIElementPostNotificationForObservedElementSymbol == 0 {
		return 0
	}
	return __AXUIElementPostNotificationForObservedElementSymbol
}

var __AXUIElementPostNotificationWithInfoSymbol uintptr

// _AXUIElementPostNotificationWithInfo has an opaque C signature in discovered metadata.
// Call _AXUIElementPostNotificationWithInfoSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementPostNotificationWithInfo
func _AXUIElementPostNotificationWithInfo() {
	panic("hiservices: symbol _AXUIElementPostNotificationWithInfo has opaque signature; use _AXUIElementPostNotificationWithInfoSymbol() and a typed manual wrapper")
}

// _AXUIElementPostNotificationWithInfoSymbol returns the raw symbol address for _AXUIElementPostNotificationWithInfo.
func _AXUIElementPostNotificationWithInfoSymbol() uintptr {
	if __AXUIElementPostNotificationWithInfoSymbol == 0 {
		return 0
	}
	return __AXUIElementPostNotificationWithInfoSymbol
}

var __AXUIElementRegisterServerWithRunLoopSymbol uintptr

// _AXUIElementRegisterServerWithRunLoop has an opaque C signature in discovered metadata.
// Call _AXUIElementRegisterServerWithRunLoopSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementRegisterServerWithRunLoop
func _AXUIElementRegisterServerWithRunLoop() {
	panic("hiservices: symbol _AXUIElementRegisterServerWithRunLoop has opaque signature; use _AXUIElementRegisterServerWithRunLoopSymbol() and a typed manual wrapper")
}

// _AXUIElementRegisterServerWithRunLoopSymbol returns the raw symbol address for _AXUIElementRegisterServerWithRunLoop.
func _AXUIElementRegisterServerWithRunLoopSymbol() uintptr {
	if __AXUIElementRegisterServerWithRunLoopSymbol == 0 {
		return 0
	}
	return __AXUIElementRegisterServerWithRunLoopSymbol
}

var __AXUIElementRemoteTokenCreateSymbol uintptr

// _AXUIElementRemoteTokenCreate has an opaque C signature in discovered metadata.
// Call _AXUIElementRemoteTokenCreateSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementRemoteTokenCreate
func _AXUIElementRemoteTokenCreate() {
	panic("hiservices: symbol _AXUIElementRemoteTokenCreate has opaque signature; use _AXUIElementRemoteTokenCreateSymbol() and a typed manual wrapper")
}

// _AXUIElementRemoteTokenCreateSymbol returns the raw symbol address for _AXUIElementRemoteTokenCreate.
func _AXUIElementRemoteTokenCreateSymbol() uintptr {
	if __AXUIElementRemoteTokenCreateSymbol == 0 {
		return 0
	}
	return __AXUIElementRemoteTokenCreateSymbol
}

var __AXUIElementRequestServicedBySecondaryAXThreadSymbol uintptr

// _AXUIElementRequestServicedBySecondaryAXThread has an opaque C signature in discovered metadata.
// Call _AXUIElementRequestServicedBySecondaryAXThreadSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementRequestServicedBySecondaryAXThread
func _AXUIElementRequestServicedBySecondaryAXThread() {
	panic("hiservices: symbol _AXUIElementRequestServicedBySecondaryAXThread has opaque signature; use _AXUIElementRequestServicedBySecondaryAXThreadSymbol() and a typed manual wrapper")
}

// _AXUIElementRequestServicedBySecondaryAXThreadSymbol returns the raw symbol address for _AXUIElementRequestServicedBySecondaryAXThread.
func _AXUIElementRequestServicedBySecondaryAXThreadSymbol() uintptr {
	if __AXUIElementRequestServicedBySecondaryAXThreadSymbol == 0 {
		return 0
	}
	return __AXUIElementRequestServicedBySecondaryAXThreadSymbol
}

var __AXUIElementUnregisterServerSymbol uintptr

// _AXUIElementUnregisterServer has an opaque C signature in discovered metadata.
// Call _AXUIElementUnregisterServerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementUnregisterServer
func _AXUIElementUnregisterServer() {
	panic("hiservices: symbol _AXUIElementUnregisterServer has opaque signature; use _AXUIElementUnregisterServerSymbol() and a typed manual wrapper")
}

// _AXUIElementUnregisterServerSymbol returns the raw symbol address for _AXUIElementUnregisterServer.
func _AXUIElementUnregisterServerSymbol() uintptr {
	if __AXUIElementUnregisterServerSymbol == 0 {
		return 0
	}
	return __AXUIElementUnregisterServerSymbol
}

var __AXUIElementUpdateServerSymbol uintptr

// _AXUIElementUpdateServer has an opaque C signature in discovered metadata.
// Call _AXUIElementUpdateServerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementUpdateServer
func _AXUIElementUpdateServer() {
	panic("hiservices: symbol _AXUIElementUpdateServer has opaque signature; use _AXUIElementUpdateServerSymbol() and a typed manual wrapper")
}

// _AXUIElementUpdateServerSymbol returns the raw symbol address for _AXUIElementUpdateServer.
func _AXUIElementUpdateServerSymbol() uintptr {
	if __AXUIElementUpdateServerSymbol == 0 {
		return 0
	}
	return __AXUIElementUpdateServerSymbol
}

var __AXUIElementUseSecondaryAXThreadSymbol uintptr

// _AXUIElementUseSecondaryAXThread has an opaque C signature in discovered metadata.
// Call _AXUIElementUseSecondaryAXThreadSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AXUIElementUseSecondaryAXThread
func _AXUIElementUseSecondaryAXThread() {
	panic("hiservices: symbol _AXUIElementUseSecondaryAXThread has opaque signature; use _AXUIElementUseSecondaryAXThreadSymbol() and a typed manual wrapper")
}

// _AXUIElementUseSecondaryAXThreadSymbol returns the raw symbol address for _AXUIElementUseSecondaryAXThread.
func _AXUIElementUseSecondaryAXThreadSymbol() uintptr {
	if __AXUIElementUseSecondaryAXThreadSymbol == 0 {
		return 0
	}
	return __AXUIElementUseSecondaryAXThreadSymbol
}

var __AddLabelsChangedCallbackSymbol uintptr

// _AddLabelsChangedCallback has an opaque C signature in discovered metadata.
// Call _AddLabelsChangedCallbackSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_AddLabelsChangedCallback
func _AddLabelsChangedCallback() {
	panic("hiservices: symbol _AddLabelsChangedCallback has opaque signature; use _AddLabelsChangedCallbackSymbol() and a typed manual wrapper")
}

// _AddLabelsChangedCallbackSymbol returns the raw symbol address for _AddLabelsChangedCallback.
func _AddLabelsChangedCallbackSymbol() uintptr {
	if __AddLabelsChangedCallbackSymbol == 0 {
		return 0
	}
	return __AddLabelsChangedCallbackSymbol
}

var __CopyProcessBundleLocationURLSymbol uintptr

// _CopyProcessBundleLocationURL has an opaque C signature in discovered metadata.
// Call _CopyProcessBundleLocationURLSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_CopyProcessBundleLocationURL
func _CopyProcessBundleLocationURL() {
	panic("hiservices: symbol _CopyProcessBundleLocationURL has opaque signature; use _CopyProcessBundleLocationURLSymbol() and a typed manual wrapper")
}

// _CopyProcessBundleLocationURLSymbol returns the raw symbol address for _CopyProcessBundleLocationURL.
func _CopyProcessBundleLocationURLSymbol() uintptr {
	if __CopyProcessBundleLocationURLSymbol == 0 {
		return 0
	}
	return __CopyProcessBundleLocationURLSymbol
}

var __GDBIconsCGCacheListSymbol uintptr

// _GDBIconsCGCacheList has an opaque C signature in discovered metadata.
// Call _GDBIconsCGCacheListSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_GDBIconsCGCacheList
func _GDBIconsCGCacheList() {
	panic("hiservices: symbol _GDBIconsCGCacheList has opaque signature; use _GDBIconsCGCacheListSymbol() and a typed manual wrapper")
}

// _GDBIconsCGCacheListSymbol returns the raw symbol address for _GDBIconsCGCacheList.
func _GDBIconsCGCacheListSymbol() uintptr {
	if __GDBIconsCGCacheListSymbol == 0 {
		return 0
	}
	return __GDBIconsCGCacheListSymbol
}

var __GetApplicationDesiresAttentionSymbol uintptr

// _GetApplicationDesiresAttention has an opaque C signature in discovered metadata.
// Call _GetApplicationDesiresAttentionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_GetApplicationDesiresAttention
func _GetApplicationDesiresAttention() {
	panic("hiservices: symbol _GetApplicationDesiresAttention has opaque signature; use _GetApplicationDesiresAttentionSymbol() and a typed manual wrapper")
}

// _GetApplicationDesiresAttentionSymbol returns the raw symbol address for _GetApplicationDesiresAttention.
func _GetApplicationDesiresAttentionSymbol() uintptr {
	if __GetApplicationDesiresAttentionSymbol == 0 {
		return 0
	}
	return __GetApplicationDesiresAttentionSymbol
}

var __GetFrontUIProcessSymbol uintptr

// _GetFrontUIProcess has an opaque C signature in discovered metadata.
// Call _GetFrontUIProcessSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_GetFrontUIProcess
func _GetFrontUIProcess() {
	panic("hiservices: symbol _GetFrontUIProcess has opaque signature; use _GetFrontUIProcessSymbol() and a typed manual wrapper")
}

// _GetFrontUIProcessSymbol returns the raw symbol address for _GetFrontUIProcess.
func _GetFrontUIProcessSymbol() uintptr {
	if __GetFrontUIProcessSymbol == 0 {
		return 0
	}
	return __GetFrontUIProcessSymbol
}

var __HIE_CaptureExceptionTelltaleOnceSymbol uintptr

// _HIE_CaptureExceptionTelltaleOnce has an opaque C signature in discovered metadata.
// Call _HIE_CaptureExceptionTelltaleOnceSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIE_CaptureExceptionTelltaleOnce
func _HIE_CaptureExceptionTelltaleOnce() {
	panic("hiservices: symbol _HIE_CaptureExceptionTelltaleOnce has opaque signature; use _HIE_CaptureExceptionTelltaleOnceSymbol() and a typed manual wrapper")
}

// _HIE_CaptureExceptionTelltaleOnceSymbol returns the raw symbol address for _HIE_CaptureExceptionTelltaleOnce.
func _HIE_CaptureExceptionTelltaleOnceSymbol() uintptr {
	if __HIE_CaptureExceptionTelltaleOnceSymbol == 0 {
		return 0
	}
	return __HIE_CaptureExceptionTelltaleOnceSymbol
}

var __HIE_CrashSymbol uintptr

// _HIE_Crash has an opaque C signature in discovered metadata.
// Call _HIE_CrashSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIE_Crash
func _HIE_Crash() {
	panic("hiservices: symbol _HIE_Crash has opaque signature; use _HIE_CrashSymbol() and a typed manual wrapper")
}

// _HIE_CrashSymbol returns the raw symbol address for _HIE_Crash.
func _HIE_CrashSymbol() uintptr {
	if __HIE_CrashSymbol == 0 {
		return 0
	}
	return __HIE_CrashSymbol
}

var __HIE_WillSwallowExceptionSymbol uintptr

// _HIE_WillSwallowException has an opaque C signature in discovered metadata.
// Call _HIE_WillSwallowExceptionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIE_WillSwallowException
func _HIE_WillSwallowException() {
	panic("hiservices: symbol _HIE_WillSwallowException has opaque signature; use _HIE_WillSwallowExceptionSymbol() and a typed manual wrapper")
}

// _HIE_WillSwallowExceptionSymbol returns the raw symbol address for _HIE_WillSwallowException.
func _HIE_WillSwallowExceptionSymbol() uintptr {
	if __HIE_WillSwallowExceptionSymbol == 0 {
		return 0
	}
	return __HIE_WillSwallowExceptionSymbol
}

var __HIRLU_AddRunLoopModeForDeferredActionsSymbol uintptr

// _HIRLU_AddRunLoopModeForDeferredActions has an opaque C signature in discovered metadata.
// Call _HIRLU_AddRunLoopModeForDeferredActionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIRLU_AddRunLoopModeForDeferredActions
func _HIRLU_AddRunLoopModeForDeferredActions() {
	panic("hiservices: symbol _HIRLU_AddRunLoopModeForDeferredActions has opaque signature; use _HIRLU_AddRunLoopModeForDeferredActionsSymbol() and a typed manual wrapper")
}

// _HIRLU_AddRunLoopModeForDeferredActionsSymbol returns the raw symbol address for _HIRLU_AddRunLoopModeForDeferredActions.
func _HIRLU_AddRunLoopModeForDeferredActionsSymbol() uintptr {
	if __HIRLU_AddRunLoopModeForDeferredActionsSymbol == 0 {
		return 0
	}
	return __HIRLU_AddRunLoopModeForDeferredActionsSymbol
}

var __HIRLU_DeferActionsSymbol uintptr

// _HIRLU_DeferActions has an opaque C signature in discovered metadata.
// Call _HIRLU_DeferActionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIRLU_DeferActions
func _HIRLU_DeferActions() {
	panic("hiservices: symbol _HIRLU_DeferActions has opaque signature; use _HIRLU_DeferActionsSymbol() and a typed manual wrapper")
}

// _HIRLU_DeferActionsSymbol returns the raw symbol address for _HIRLU_DeferActions.
func _HIRLU_DeferActionsSymbol() uintptr {
	if __HIRLU_DeferActionsSymbol == 0 {
		return 0
	}
	return __HIRLU_DeferActionsSymbol
}

var __HIRunLoopSemaphoreCreateWithRunLoopModeSymbol uintptr

// _HIRunLoopSemaphoreCreateWithRunLoopMode has an opaque C signature in discovered metadata.
// Call _HIRunLoopSemaphoreCreateWithRunLoopModeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIRunLoopSemaphoreCreateWithRunLoopMode
func _HIRunLoopSemaphoreCreateWithRunLoopMode() {
	panic("hiservices: symbol _HIRunLoopSemaphoreCreateWithRunLoopMode has opaque signature; use _HIRunLoopSemaphoreCreateWithRunLoopModeSymbol() and a typed manual wrapper")
}

// _HIRunLoopSemaphoreCreateWithRunLoopModeSymbol returns the raw symbol address for _HIRunLoopSemaphoreCreateWithRunLoopMode.
func _HIRunLoopSemaphoreCreateWithRunLoopModeSymbol() uintptr {
	if __HIRunLoopSemaphoreCreateWithRunLoopModeSymbol == 0 {
		return 0
	}
	return __HIRunLoopSemaphoreCreateWithRunLoopModeSymbol
}

var __HIRunLoopSemaphoreSetLegendSymbol uintptr

// _HIRunLoopSemaphoreSetLegend has an opaque C signature in discovered metadata.
// Call _HIRunLoopSemaphoreSetLegendSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIRunLoopSemaphoreSetLegend
func _HIRunLoopSemaphoreSetLegend() {
	panic("hiservices: symbol _HIRunLoopSemaphoreSetLegend has opaque signature; use _HIRunLoopSemaphoreSetLegendSymbol() and a typed manual wrapper")
}

// _HIRunLoopSemaphoreSetLegendSymbol returns the raw symbol address for _HIRunLoopSemaphoreSetLegend.
func _HIRunLoopSemaphoreSetLegendSymbol() uintptr {
	if __HIRunLoopSemaphoreSetLegendSymbol == 0 {
		return 0
	}
	return __HIRunLoopSemaphoreSetLegendSymbol
}

var __HIRunLoopSemaphoreSignalSymbol uintptr

// _HIRunLoopSemaphoreSignal has an opaque C signature in discovered metadata.
// Call _HIRunLoopSemaphoreSignalSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIRunLoopSemaphoreSignal
func _HIRunLoopSemaphoreSignal() {
	panic("hiservices: symbol _HIRunLoopSemaphoreSignal has opaque signature; use _HIRunLoopSemaphoreSignalSymbol() and a typed manual wrapper")
}

// _HIRunLoopSemaphoreSignalSymbol returns the raw symbol address for _HIRunLoopSemaphoreSignal.
func _HIRunLoopSemaphoreSignalSymbol() uintptr {
	if __HIRunLoopSemaphoreSignalSymbol == 0 {
		return 0
	}
	return __HIRunLoopSemaphoreSignalSymbol
}

var __HIRunLoopSemaphoreWaitSymbol uintptr

// _HIRunLoopSemaphoreWait has an opaque C signature in discovered metadata.
// Call _HIRunLoopSemaphoreWaitSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIRunLoopSemaphoreWait
func _HIRunLoopSemaphoreWait() {
	panic("hiservices: symbol _HIRunLoopSemaphoreWait has opaque signature; use _HIRunLoopSemaphoreWaitSymbol() and a typed manual wrapper")
}

// _HIRunLoopSemaphoreWaitSymbol returns the raw symbol address for _HIRunLoopSemaphoreWait.
func _HIRunLoopSemaphoreWaitSymbol() uintptr {
	if __HIRunLoopSemaphoreWaitSymbol == 0 {
		return 0
	}
	return __HIRunLoopSemaphoreWaitSymbol
}

var __HIShapeCreateWithCGImageSymbol uintptr

// _HIShapeCreateWithCGImage has an opaque C signature in discovered metadata.
// Call _HIShapeCreateWithCGImageSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIShapeCreateWithCGImage
func _HIShapeCreateWithCGImage() {
	panic("hiservices: symbol _HIShapeCreateWithCGImage has opaque signature; use _HIShapeCreateWithCGImageSymbol() and a typed manual wrapper")
}

// _HIShapeCreateWithCGImageSymbol returns the raw symbol address for _HIShapeCreateWithCGImage.
func _HIShapeCreateWithCGImageSymbol() uintptr {
	if __HIShapeCreateWithCGImageSymbol == 0 {
		return 0
	}
	return __HIShapeCreateWithCGImageSymbol
}

var __HIShapeCreateWithCGSRegionObjSymbol uintptr

// _HIShapeCreateWithCGSRegionObj has an opaque C signature in discovered metadata.
// Call _HIShapeCreateWithCGSRegionObjSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIShapeCreateWithCGSRegionObj
func _HIShapeCreateWithCGSRegionObj() {
	panic("hiservices: symbol _HIShapeCreateWithCGSRegionObj has opaque signature; use _HIShapeCreateWithCGSRegionObjSymbol() and a typed manual wrapper")
}

// _HIShapeCreateWithCGSRegionObjSymbol returns the raw symbol address for _HIShapeCreateWithCGSRegionObj.
func _HIShapeCreateWithCGSRegionObjSymbol() uintptr {
	if __HIShapeCreateWithCGSRegionObjSymbol == 0 {
		return 0
	}
	return __HIShapeCreateWithCGSRegionObjSymbol
}

var __HIShapeGetNativeSymbol uintptr

// _HIShapeGetNative has an opaque C signature in discovered metadata.
// Call _HIShapeGetNativeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIShapeGetNative
func _HIShapeGetNative() {
	panic("hiservices: symbol _HIShapeGetNative has opaque signature; use _HIShapeGetNativeSymbol() and a typed manual wrapper")
}

// _HIShapeGetNativeSymbol returns the raw symbol address for _HIShapeGetNative.
func _HIShapeGetNativeSymbol() uintptr {
	if __HIShapeGetNativeSymbol == 0 {
		return 0
	}
	return __HIShapeGetNativeSymbol
}

var __HIShapeOutsetToPixelBoundarySymbol uintptr

// _HIShapeOutsetToPixelBoundary has an opaque C signature in discovered metadata.
// Call _HIShapeOutsetToPixelBoundarySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIShapeOutsetToPixelBoundary
func _HIShapeOutsetToPixelBoundary() {
	panic("hiservices: symbol _HIShapeOutsetToPixelBoundary has opaque signature; use _HIShapeOutsetToPixelBoundarySymbol() and a typed manual wrapper")
}

// _HIShapeOutsetToPixelBoundarySymbol returns the raw symbol address for _HIShapeOutsetToPixelBoundary.
func _HIShapeOutsetToPixelBoundarySymbol() uintptr {
	if __HIShapeOutsetToPixelBoundarySymbol == 0 {
		return 0
	}
	return __HIShapeOutsetToPixelBoundarySymbol
}

var __HIShapeSetImmutableSymbol uintptr

// _HIShapeSetImmutable has an opaque C signature in discovered metadata.
// Call _HIShapeSetImmutableSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIShapeSetImmutable
func _HIShapeSetImmutable() {
	panic("hiservices: symbol _HIShapeSetImmutable has opaque signature; use _HIShapeSetImmutableSymbol() and a typed manual wrapper")
}

// _HIShapeSetImmutableSymbol returns the raw symbol address for _HIShapeSetImmutable.
func _HIShapeSetImmutableSymbol() uintptr {
	if __HIShapeSetImmutableSymbol == 0 {
		return 0
	}
	return __HIShapeSetImmutableSymbol
}

var __HIShapeSetShapeWithOffsetSymbol uintptr

// _HIShapeSetShapeWithOffset has an opaque C signature in discovered metadata.
// Call _HIShapeSetShapeWithOffsetSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HIShapeSetShapeWithOffset
func _HIShapeSetShapeWithOffset() {
	panic("hiservices: symbol _HIShapeSetShapeWithOffset has opaque signature; use _HIShapeSetShapeWithOffsetSymbol() and a typed manual wrapper")
}

// _HIShapeSetShapeWithOffsetSymbol returns the raw symbol address for _HIShapeSetShapeWithOffset.
func _HIShapeSetShapeWithOffsetSymbol() uintptr {
	if __HIShapeSetShapeWithOffsetSymbol == 0 {
		return 0
	}
	return __HIShapeSetShapeWithOffsetSymbol
}

var __HideOtherApplicationsSymbol uintptr

// _HideOtherApplications has an opaque C signature in discovered metadata.
// Call _HideOtherApplicationsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_HideOtherApplications
func _HideOtherApplications() {
	panic("hiservices: symbol _HideOtherApplications has opaque signature; use _HideOtherApplicationsSymbol() and a typed manual wrapper")
}

// _HideOtherApplicationsSymbol returns the raw symbol address for _HideOtherApplications.
func _HideOtherApplicationsSymbol() uintptr {
	if __HideOtherApplicationsSymbol == 0 {
		return 0
	}
	return __HideOtherApplicationsSymbol
}

var __ICCopyMailHostNameSymbol uintptr

// _ICCopyMailHostName has an opaque C signature in discovered metadata.
// Call _ICCopyMailHostNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_ICCopyMailHostName
func _ICCopyMailHostName() {
	panic("hiservices: symbol _ICCopyMailHostName has opaque signature; use _ICCopyMailHostNameSymbol() and a typed manual wrapper")
}

// _ICCopyMailHostNameSymbol returns the raw symbol address for _ICCopyMailHostName.
func _ICCopyMailHostNameSymbol() uintptr {
	if __ICCopyMailHostNameSymbol == 0 {
		return 0
	}
	return __ICCopyMailHostNameSymbol
}

var __ICCopyMailUserNameSymbol uintptr

// _ICCopyMailUserName has an opaque C signature in discovered metadata.
// Call _ICCopyMailUserNameSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_ICCopyMailUserName
func _ICCopyMailUserName() {
	panic("hiservices: symbol _ICCopyMailUserName has opaque signature; use _ICCopyMailUserNameSymbol() and a typed manual wrapper")
}

// _ICCopyMailUserNameSymbol returns the raw symbol address for _ICCopyMailUserName.
func _ICCopyMailUserNameSymbol() uintptr {
	if __ICCopyMailUserNameSymbol == 0 {
		return 0
	}
	return __ICCopyMailUserNameSymbol
}

var __ISCreateCGImageForTypeSymbol uintptr

// _ISCreateCGImageForType has an opaque C signature in discovered metadata.
// Call _ISCreateCGImageForTypeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_ISCreateCGImageForType
func _ISCreateCGImageForType() {
	panic("hiservices: symbol _ISCreateCGImageForType has opaque signature; use _ISCreateCGImageForTypeSymbol() and a typed manual wrapper")
}

// _ISCreateCGImageForTypeSymbol returns the raw symbol address for _ISCreateCGImageForType.
func _ISCreateCGImageForTypeSymbol() uintptr {
	if __ISCreateCGImageForTypeSymbol == 0 {
		return 0
	}
	return __ISCreateCGImageForTypeSymbol
}

var __ISCreateCGImageForTypeAtScaleSymbol uintptr

// _ISCreateCGImageForTypeAtScale has an opaque C signature in discovered metadata.
// Call _ISCreateCGImageForTypeAtScaleSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_ISCreateCGImageForTypeAtScale
func _ISCreateCGImageForTypeAtScale() {
	panic("hiservices: symbol _ISCreateCGImageForTypeAtScale has opaque signature; use _ISCreateCGImageForTypeAtScaleSymbol() and a typed manual wrapper")
}

// _ISCreateCGImageForTypeAtScaleSymbol returns the raw symbol address for _ISCreateCGImageForTypeAtScale.
func _ISCreateCGImageForTypeAtScaleSymbol() uintptr {
	if __ISCreateCGImageForTypeAtScaleSymbol == 0 {
		return 0
	}
	return __ISCreateCGImageForTypeAtScaleSymbol
}

var __IconServicesGetCGImageRefFromIconRefSymbol uintptr

// _IconServicesGetCGImageRefFromIconRef has an opaque C signature in discovered metadata.
// Call _IconServicesGetCGImageRefFromIconRefSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_IconServicesGetCGImageRefFromIconRef
func _IconServicesGetCGImageRefFromIconRef() {
	panic("hiservices: symbol _IconServicesGetCGImageRefFromIconRef has opaque signature; use _IconServicesGetCGImageRefFromIconRefSymbol() and a typed manual wrapper")
}

// _IconServicesGetCGImageRefFromIconRefSymbol returns the raw symbol address for _IconServicesGetCGImageRefFromIconRef.
func _IconServicesGetCGImageRefFromIconRefSymbol() uintptr {
	if __IconServicesGetCGImageRefFromIconRefSymbol == 0 {
		return 0
	}
	return __IconServicesGetCGImageRefFromIconRefSymbol
}

var __IconServicesGetCGImageRefFromURLSymbol uintptr

// _IconServicesGetCGImageRefFromURL has an opaque C signature in discovered metadata.
// Call _IconServicesGetCGImageRefFromURLSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_IconServicesGetCGImageRefFromURL
func _IconServicesGetCGImageRefFromURL() {
	panic("hiservices: symbol _IconServicesGetCGImageRefFromURL has opaque signature; use _IconServicesGetCGImageRefFromURLSymbol() and a typed manual wrapper")
}

// _IconServicesGetCGImageRefFromURLSymbol returns the raw symbol address for _IconServicesGetCGImageRefFromURL.
func _IconServicesGetCGImageRefFromURLSymbol() uintptr {
	if __IconServicesGetCGImageRefFromURLSymbol == 0 {
		return 0
	}
	return __IconServicesGetCGImageRefFromURLSymbol
}

var __InstallGURLEventHandlerSymbol uintptr

// _InstallGURLEventHandler has an opaque C signature in discovered metadata.
// Call _InstallGURLEventHandlerSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_InstallGURLEventHandler
func _InstallGURLEventHandler() {
	panic("hiservices: symbol _InstallGURLEventHandler has opaque signature; use _InstallGURLEventHandlerSymbol() and a typed manual wrapper")
}

// _InstallGURLEventHandlerSymbol returns the raw symbol address for _InstallGURLEventHandler.
func _InstallGURLEventHandlerSymbol() uintptr {
	if __InstallGURLEventHandlerSymbol == 0 {
		return 0
	}
	return __InstallGURLEventHandlerSymbol
}

var __PIPZoomingEnabledSymbol uintptr

// _PIPZoomingEnabled has an opaque C signature in discovered metadata.
// Call _PIPZoomingEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_PIPZoomingEnabled
func _PIPZoomingEnabled() {
	panic("hiservices: symbol _PIPZoomingEnabled has opaque signature; use _PIPZoomingEnabledSymbol() and a typed manual wrapper")
}

// _PIPZoomingEnabledSymbol returns the raw symbol address for _PIPZoomingEnabled.
func _PIPZoomingEnabledSymbol() uintptr {
	if __PIPZoomingEnabledSymbol == 0 {
		return 0
	}
	return __PIPZoomingEnabledSymbol
}

var __RegisterApplicationSymbol uintptr

// _RegisterApplication has an opaque C signature in discovered metadata.
// Call _RegisterApplicationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_RegisterApplication
func _RegisterApplication() {
	panic("hiservices: symbol _RegisterApplication has opaque signature; use _RegisterApplicationSymbol() and a typed manual wrapper")
}

// _RegisterApplicationSymbol returns the raw symbol address for _RegisterApplication.
func _RegisterApplicationSymbol() uintptr {
	if __RegisterApplicationSymbol == 0 {
		return 0
	}
	return __RegisterApplicationSymbol
}

var __RegisterAsSessionLauncherApplicationSymbol uintptr

// _RegisterAsSessionLauncherApplication has an opaque C signature in discovered metadata.
// Call _RegisterAsSessionLauncherApplicationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_RegisterAsSessionLauncherApplication
func _RegisterAsSessionLauncherApplication() {
	panic("hiservices: symbol _RegisterAsSessionLauncherApplication has opaque signature; use _RegisterAsSessionLauncherApplicationSymbol() and a typed manual wrapper")
}

// _RegisterAsSessionLauncherApplicationSymbol returns the raw symbol address for _RegisterAsSessionLauncherApplication.
func _RegisterAsSessionLauncherApplicationSymbol() uintptr {
	if __RegisterAsSessionLauncherApplicationSymbol == 0 {
		return 0
	}
	return __RegisterAsSessionLauncherApplicationSymbol
}

var __RemoveLabelsChangedCallbackSymbol uintptr

// _RemoveLabelsChangedCallback has an opaque C signature in discovered metadata.
// Call _RemoveLabelsChangedCallbackSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_RemoveLabelsChangedCallback
func _RemoveLabelsChangedCallback() {
	panic("hiservices: symbol _RemoveLabelsChangedCallback has opaque signature; use _RemoveLabelsChangedCallbackSymbol() and a typed manual wrapper")
}

// _RemoveLabelsChangedCallbackSymbol returns the raw symbol address for _RemoveLabelsChangedCallback.
func _RemoveLabelsChangedCallbackSymbol() uintptr {
	if __RemoveLabelsChangedCallbackSymbol == 0 {
		return 0
	}
	return __RemoveLabelsChangedCallbackSymbol
}

var __SetApplicationDesiresAttentionSymbol uintptr

// _SetApplicationDesiresAttention has an opaque C signature in discovered metadata.
// Call _SetApplicationDesiresAttentionSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_SetApplicationDesiresAttention
func _SetApplicationDesiresAttention() {
	panic("hiservices: symbol _SetApplicationDesiresAttention has opaque signature; use _SetApplicationDesiresAttentionSymbol() and a typed manual wrapper")
}

// _SetApplicationDesiresAttentionSymbol returns the raw symbol address for _SetApplicationDesiresAttention.
func _SetApplicationDesiresAttentionSymbol() uintptr {
	if __SetApplicationDesiresAttentionSymbol == 0 {
		return 0
	}
	return __SetApplicationDesiresAttentionSymbol
}

var __SetFrontProcessWithOptionsSymbol uintptr

// _SetFrontProcessWithOptions has an opaque C signature in discovered metadata.
// Call _SetFrontProcessWithOptionsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_SetFrontProcessWithOptions
func _SetFrontProcessWithOptions() {
	panic("hiservices: symbol _SetFrontProcessWithOptions has opaque signature; use _SetFrontProcessWithOptionsSymbol() and a typed manual wrapper")
}

// _SetFrontProcessWithOptionsSymbol returns the raw symbol address for _SetFrontProcessWithOptions.
func _SetFrontProcessWithOptionsSymbol() uintptr {
	if __SetFrontProcessWithOptionsSymbol == 0 {
		return 0
	}
	return __SetFrontProcessWithOptionsSymbol
}

var __SetHLTBWakeUpHookSymbol uintptr

// _SetHLTBWakeUpHook has an opaque C signature in discovered metadata.
// Call _SetHLTBWakeUpHookSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_SetHLTBWakeUpHook
func _SetHLTBWakeUpHook() {
	panic("hiservices: symbol _SetHLTBWakeUpHook has opaque signature; use _SetHLTBWakeUpHookSymbol() and a typed manual wrapper")
}

// _SetHLTBWakeUpHookSymbol returns the raw symbol address for _SetHLTBWakeUpHook.
func _SetHLTBWakeUpHookSymbol() uintptr {
	if __SetHLTBWakeUpHookSymbol == 0 {
		return 0
	}
	return __SetHLTBWakeUpHookSymbol
}

var __ShowAllApplicationsSymbol uintptr

// _ShowAllApplications has an opaque C signature in discovered metadata.
// Call _ShowAllApplicationsSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_ShowAllApplications
func _ShowAllApplications() {
	panic("hiservices: symbol _ShowAllApplications has opaque signature; use _ShowAllApplicationsSymbol() and a typed manual wrapper")
}

// _ShowAllApplicationsSymbol returns the raw symbol address for _ShowAllApplications.
func _ShowAllApplicationsSymbol() uintptr {
	if __ShowAllApplicationsSymbol == 0 {
		return 0
	}
	return __ShowAllApplicationsSymbol
}

var __SignalApplicationReadySymbol uintptr

// _SignalApplicationReady has an opaque C signature in discovered metadata.
// Call _SignalApplicationReadySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_SignalApplicationReady
func _SignalApplicationReady() {
	panic("hiservices: symbol _SignalApplicationReady has opaque signature; use _SignalApplicationReadySymbol() and a typed manual wrapper")
}

// _SignalApplicationReadySymbol returns the raw symbol address for _SignalApplicationReady.
func _SignalApplicationReadySymbol() uintptr {
	if __SignalApplicationReadySymbol == 0 {
		return 0
	}
	return __SignalApplicationReadySymbol
}

var __UAZoomFocusChangeSymbol uintptr

// _UAZoomFocusChange has an opaque C signature in discovered metadata.
// Call _UAZoomFocusChangeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_UAZoomFocusChange
func _UAZoomFocusChange() {
	panic("hiservices: symbol _UAZoomFocusChange has opaque signature; use _UAZoomFocusChangeSymbol() and a typed manual wrapper")
}

// _UAZoomFocusChangeSymbol returns the raw symbol address for _UAZoomFocusChange.
func _UAZoomFocusChangeSymbol() uintptr {
	if __UAZoomFocusChangeSymbol == 0 {
		return 0
	}
	return __UAZoomFocusChangeSymbol
}

var __UAZoomFocusChangeAnchoredSymbol uintptr

// _UAZoomFocusChangeAnchored has an opaque C signature in discovered metadata.
// Call _UAZoomFocusChangeAnchoredSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_UAZoomFocusChangeAnchored
func _UAZoomFocusChangeAnchored() {
	panic("hiservices: symbol _UAZoomFocusChangeAnchored has opaque signature; use _UAZoomFocusChangeAnchoredSymbol() and a typed manual wrapper")
}

// _UAZoomFocusChangeAnchoredSymbol returns the raw symbol address for _UAZoomFocusChangeAnchored.
func _UAZoomFocusChangeAnchoredSymbol() uintptr {
	if __UAZoomFocusChangeAnchoredSymbol == 0 {
		return 0
	}
	return __UAZoomFocusChangeAnchoredSymbol
}

var __UAZoomFocusChangeHighlightRectSymbol uintptr

// _UAZoomFocusChangeHighlightRect has an opaque C signature in discovered metadata.
// Call _UAZoomFocusChangeHighlightRectSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_UAZoomFocusChangeHighlightRect
func _UAZoomFocusChangeHighlightRect() {
	panic("hiservices: symbol _UAZoomFocusChangeHighlightRect has opaque signature; use _UAZoomFocusChangeHighlightRectSymbol() and a typed manual wrapper")
}

// _UAZoomFocusChangeHighlightRectSymbol returns the raw symbol address for _UAZoomFocusChangeHighlightRect.
func _UAZoomFocusChangeHighlightRectSymbol() uintptr {
	if __UAZoomFocusChangeHighlightRectSymbol == 0 {
		return 0
	}
	return __UAZoomFocusChangeHighlightRectSymbol
}

var __UAZoomFocusChangeHighlightRectAnchoredSymbol uintptr

// _UAZoomFocusChangeHighlightRectAnchored has an opaque C signature in discovered metadata.
// Call _UAZoomFocusChangeHighlightRectAnchoredSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_UAZoomFocusChangeHighlightRectAnchored
func _UAZoomFocusChangeHighlightRectAnchored() {
	panic("hiservices: symbol _UAZoomFocusChangeHighlightRectAnchored has opaque signature; use _UAZoomFocusChangeHighlightRectAnchoredSymbol() and a typed manual wrapper")
}

// _UAZoomFocusChangeHighlightRectAnchoredSymbol returns the raw symbol address for _UAZoomFocusChangeHighlightRectAnchored.
func _UAZoomFocusChangeHighlightRectAnchoredSymbol() uintptr {
	if __UAZoomFocusChangeHighlightRectAnchoredSymbol == 0 {
		return 0
	}
	return __UAZoomFocusChangeHighlightRectAnchoredSymbol
}

var __UAZoomingEnabledSymbol uintptr

// _UAZoomingEnabled has an opaque C signature in discovered metadata.
// Call _UAZoomingEnabledSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_UAZoomingEnabled
func _UAZoomingEnabled() {
	panic("hiservices: symbol _UAZoomingEnabled has opaque signature; use _UAZoomingEnabledSymbol() and a typed manual wrapper")
}

// _UAZoomingEnabledSymbol returns the raw symbol address for _UAZoomingEnabled.
func _UAZoomingEnabledSymbol() uintptr {
	if __UAZoomingEnabledSymbol == 0 {
		return 0
	}
	return __UAZoomingEnabledSymbol
}

var __UnregisterAsSessionLauncherApplicationSymbol uintptr

// _UnregisterAsSessionLauncherApplication has an opaque C signature in discovered metadata.
// Call _UnregisterAsSessionLauncherApplicationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/_UnregisterAsSessionLauncherApplication
func _UnregisterAsSessionLauncherApplication() {
	panic("hiservices: symbol _UnregisterAsSessionLauncherApplication has opaque signature; use _UnregisterAsSessionLauncherApplicationSymbol() and a typed manual wrapper")
}

// _UnregisterAsSessionLauncherApplicationSymbol returns the raw symbol address for _UnregisterAsSessionLauncherApplication.
func _UnregisterAsSessionLauncherApplicationSymbol() uintptr {
	if __UnregisterAsSessionLauncherApplicationSymbol == 0 {
		return 0
	}
	return __UnregisterAsSessionLauncherApplicationSymbol
}

var _gDockDragCallbackSymbol uintptr

// GDockDragCallback has an opaque C signature in discovered metadata.
// Call GDockDragCallbackSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/gDockDragCallback
func GDockDragCallback() {
	panic("hiservices: symbol gDockDragCallback has opaque signature; use GDockDragCallbackSymbol() and a typed manual wrapper")
}

// GDockDragCallbackSymbol returns the raw symbol address for gDockDragCallback.
func GDockDragCallbackSymbol() uintptr {
	if _gDockDragCallbackSymbol == 0 {
		return 0
	}
	return _gDockDragCallbackSymbol
}

var _gDockDragWindowCallbackSymbol uintptr

// GDockDragWindowCallback has an opaque C signature in discovered metadata.
// Call GDockDragWindowCallbackSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/gDockDragWindowCallback
func GDockDragWindowCallback() {
	panic("hiservices: symbol gDockDragWindowCallback has opaque signature; use GDockDragWindowCallbackSymbol() and a typed manual wrapper")
}

// GDockDragWindowCallbackSymbol returns the raw symbol address for gDockDragWindowCallback.
func GDockDragWindowCallbackSymbol() uintptr {
	if _gDockDragWindowCallbackSymbol == 0 {
		return 0
	}
	return _gDockDragWindowCallbackSymbol
}

var _kAXAttachmentTextAttributeSymbol uintptr

// KAXAttachmentTextAttribute has an opaque C signature in discovered metadata.
// Call KAXAttachmentTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXAttachmentTextAttribute
func KAXAttachmentTextAttribute() {
	panic("hiservices: symbol kAXAttachmentTextAttribute has opaque signature; use KAXAttachmentTextAttributeSymbol() and a typed manual wrapper")
}

// KAXAttachmentTextAttributeSymbol returns the raw symbol address for kAXAttachmentTextAttribute.
func KAXAttachmentTextAttributeSymbol() uintptr {
	if _kAXAttachmentTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXAttachmentTextAttributeSymbol
}

var _kAXAutocorrectedTextAttributeSymbol uintptr

// KAXAutocorrectedTextAttribute has an opaque C signature in discovered metadata.
// Call KAXAutocorrectedTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXAutocorrectedTextAttribute
func KAXAutocorrectedTextAttribute() {
	panic("hiservices: symbol kAXAutocorrectedTextAttribute has opaque signature; use KAXAutocorrectedTextAttributeSymbol() and a typed manual wrapper")
}

// KAXAutocorrectedTextAttributeSymbol returns the raw symbol address for kAXAutocorrectedTextAttribute.
func KAXAutocorrectedTextAttributeSymbol() uintptr {
	if _kAXAutocorrectedTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXAutocorrectedTextAttributeSymbol
}

var _kAXBackgroundColorTextAttributeSymbol uintptr

// KAXBackgroundColorTextAttribute has an opaque C signature in discovered metadata.
// Call KAXBackgroundColorTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXBackgroundColorTextAttribute
func KAXBackgroundColorTextAttribute() {
	panic("hiservices: symbol kAXBackgroundColorTextAttribute has opaque signature; use KAXBackgroundColorTextAttributeSymbol() and a typed manual wrapper")
}

// KAXBackgroundColorTextAttributeSymbol returns the raw symbol address for kAXBackgroundColorTextAttribute.
func KAXBackgroundColorTextAttributeSymbol() uintptr {
	if _kAXBackgroundColorTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXBackgroundColorTextAttributeSymbol
}

var _kAXElementOrderHorizontalKeySymbol uintptr

// KAXElementOrderHorizontalKey has an opaque C signature in discovered metadata.
// Call KAXElementOrderHorizontalKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXElementOrderHorizontalKey
func KAXElementOrderHorizontalKey() {
	panic("hiservices: symbol kAXElementOrderHorizontalKey has opaque signature; use KAXElementOrderHorizontalKeySymbol() and a typed manual wrapper")
}

// KAXElementOrderHorizontalKeySymbol returns the raw symbol address for kAXElementOrderHorizontalKey.
func KAXElementOrderHorizontalKeySymbol() uintptr {
	if _kAXElementOrderHorizontalKeySymbol == 0 {
		return 0
	}
	return _kAXElementOrderHorizontalKeySymbol
}

var _kAXElementOrderVerticalKeySymbol uintptr

// KAXElementOrderVerticalKey has an opaque C signature in discovered metadata.
// Call KAXElementOrderVerticalKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXElementOrderVerticalKey
func KAXElementOrderVerticalKey() {
	panic("hiservices: symbol kAXElementOrderVerticalKey has opaque signature; use KAXElementOrderVerticalKeySymbol() and a typed manual wrapper")
}

// KAXElementOrderVerticalKeySymbol returns the raw symbol address for kAXElementOrderVerticalKey.
func KAXElementOrderVerticalKeySymbol() uintptr {
	if _kAXElementOrderVerticalKeySymbol == 0 {
		return 0
	}
	return _kAXElementOrderVerticalKeySymbol
}

var _kAXElementToFocusForLayoutChangeKeySymbol uintptr

// KAXElementToFocusForLayoutChangeKey has an opaque C signature in discovered metadata.
// Call KAXElementToFocusForLayoutChangeKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXElementToFocusForLayoutChangeKey
func KAXElementToFocusForLayoutChangeKey() {
	panic("hiservices: symbol kAXElementToFocusForLayoutChangeKey has opaque signature; use KAXElementToFocusForLayoutChangeKeySymbol() and a typed manual wrapper")
}

// KAXElementToFocusForLayoutChangeKeySymbol returns the raw symbol address for kAXElementToFocusForLayoutChangeKey.
func KAXElementToFocusForLayoutChangeKeySymbol() uintptr {
	if _kAXElementToFocusForLayoutChangeKeySymbol == 0 {
		return 0
	}
	return _kAXElementToFocusForLayoutChangeKeySymbol
}

var _kAXFontFamilyKeySymbol uintptr

// KAXFontFamilyKey has an opaque C signature in discovered metadata.
// Call KAXFontFamilyKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXFontFamilyKey
func KAXFontFamilyKey() {
	panic("hiservices: symbol kAXFontFamilyKey has opaque signature; use KAXFontFamilyKeySymbol() and a typed manual wrapper")
}

// KAXFontFamilyKeySymbol returns the raw symbol address for kAXFontFamilyKey.
func KAXFontFamilyKeySymbol() uintptr {
	if _kAXFontFamilyKeySymbol == 0 {
		return 0
	}
	return _kAXFontFamilyKeySymbol
}

var _kAXFontNameKeySymbol uintptr

// KAXFontNameKey has an opaque C signature in discovered metadata.
// Call KAXFontNameKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXFontNameKey
func KAXFontNameKey() {
	panic("hiservices: symbol kAXFontNameKey has opaque signature; use KAXFontNameKeySymbol() and a typed manual wrapper")
}

// KAXFontNameKeySymbol returns the raw symbol address for kAXFontNameKey.
func KAXFontNameKeySymbol() uintptr {
	if _kAXFontNameKeySymbol == 0 {
		return 0
	}
	return _kAXFontNameKeySymbol
}

var _kAXFontSizeKeySymbol uintptr

// KAXFontSizeKey has an opaque C signature in discovered metadata.
// Call KAXFontSizeKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXFontSizeKey
func KAXFontSizeKey() {
	panic("hiservices: symbol kAXFontSizeKey has opaque signature; use KAXFontSizeKeySymbol() and a typed manual wrapper")
}

// KAXFontSizeKeySymbol returns the raw symbol address for kAXFontSizeKey.
func KAXFontSizeKeySymbol() uintptr {
	if _kAXFontSizeKeySymbol == 0 {
		return 0
	}
	return _kAXFontSizeKeySymbol
}

var _kAXFontTextAttributeSymbol uintptr

// KAXFontTextAttribute has an opaque C signature in discovered metadata.
// Call KAXFontTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXFontTextAttribute
func KAXFontTextAttribute() {
	panic("hiservices: symbol kAXFontTextAttribute has opaque signature; use KAXFontTextAttributeSymbol() and a typed manual wrapper")
}

// KAXFontTextAttributeSymbol returns the raw symbol address for kAXFontTextAttribute.
func KAXFontTextAttributeSymbol() uintptr {
	if _kAXFontTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXFontTextAttributeSymbol
}

var _kAXForegoundColorTextAttributeSymbol uintptr

// KAXForegoundColorTextAttribute has an opaque C signature in discovered metadata.
// Call KAXForegoundColorTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXForegoundColorTextAttribute
func KAXForegoundColorTextAttribute() {
	panic("hiservices: symbol kAXForegoundColorTextAttribute has opaque signature; use KAXForegoundColorTextAttributeSymbol() and a typed manual wrapper")
}

// KAXForegoundColorTextAttributeSymbol returns the raw symbol address for kAXForegoundColorTextAttribute.
func KAXForegoundColorTextAttributeSymbol() uintptr {
	if _kAXForegoundColorTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXForegoundColorTextAttributeSymbol
}

var _kAXForegroundColorTextAttributeSymbol uintptr

// KAXForegroundColorTextAttribute has an opaque C signature in discovered metadata.
// Call KAXForegroundColorTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXForegroundColorTextAttribute
func KAXForegroundColorTextAttribute() {
	panic("hiservices: symbol kAXForegroundColorTextAttribute has opaque signature; use KAXForegroundColorTextAttributeSymbol() and a typed manual wrapper")
}

// KAXForegroundColorTextAttributeSymbol returns the raw symbol address for kAXForegroundColorTextAttribute.
func KAXForegroundColorTextAttributeSymbol() uintptr {
	if _kAXForegroundColorTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXForegroundColorTextAttributeSymbol
}

var _kAXHasClientsWithAccessRemoteDeviceContentDidChangeSymbol uintptr

// KAXHasClientsWithAccessRemoteDeviceContentDidChange has an opaque C signature in discovered metadata.
// Call KAXHasClientsWithAccessRemoteDeviceContentDidChangeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXHasClientsWithAccessRemoteDeviceContentDidChange
func KAXHasClientsWithAccessRemoteDeviceContentDidChange() {
	panic("hiservices: symbol kAXHasClientsWithAccessRemoteDeviceContentDidChange has opaque signature; use KAXHasClientsWithAccessRemoteDeviceContentDidChangeSymbol() and a typed manual wrapper")
}

// KAXHasClientsWithAccessRemoteDeviceContentDidChangeSymbol returns the raw symbol address for kAXHasClientsWithAccessRemoteDeviceContentDidChange.
func KAXHasClientsWithAccessRemoteDeviceContentDidChangeSymbol() uintptr {
	if _kAXHasClientsWithAccessRemoteDeviceContentDidChangeSymbol == 0 {
		return 0
	}
	return _kAXHasClientsWithAccessRemoteDeviceContentDidChangeSymbol
}

var _kAXInterfaceBristolKeySymbol uintptr

// KAXInterfaceBristolKey has an opaque C signature in discovered metadata.
// Call KAXInterfaceBristolKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceBristolKey
func KAXInterfaceBristolKey() {
	panic("hiservices: symbol kAXInterfaceBristolKey has opaque signature; use KAXInterfaceBristolKeySymbol() and a typed manual wrapper")
}

// KAXInterfaceBristolKeySymbol returns the raw symbol address for kAXInterfaceBristolKey.
func KAXInterfaceBristolKeySymbol() uintptr {
	if _kAXInterfaceBristolKeySymbol == 0 {
		return 0
	}
	return _kAXInterfaceBristolKeySymbol
}

var _kAXInterfaceClassicInvertColorKeySymbol uintptr

// KAXInterfaceClassicInvertColorKey has an opaque C signature in discovered metadata.
// Call KAXInterfaceClassicInvertColorKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceClassicInvertColorKey
func KAXInterfaceClassicInvertColorKey() {
	panic("hiservices: symbol kAXInterfaceClassicInvertColorKey has opaque signature; use KAXInterfaceClassicInvertColorKeySymbol() and a typed manual wrapper")
}

// KAXInterfaceClassicInvertColorKeySymbol returns the raw symbol address for kAXInterfaceClassicInvertColorKey.
func KAXInterfaceClassicInvertColorKeySymbol() uintptr {
	if _kAXInterfaceClassicInvertColorKeySymbol == 0 {
		return 0
	}
	return _kAXInterfaceClassicInvertColorKeySymbol
}

var _kAXInterfaceClassicInvertColorStatusDidChangeNotificationSymbol uintptr

// KAXInterfaceClassicInvertColorStatusDidChangeNotification has an opaque C signature in discovered metadata.
// Call KAXInterfaceClassicInvertColorStatusDidChangeNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceClassicInvertColorStatusDidChangeNotification
func KAXInterfaceClassicInvertColorStatusDidChangeNotification() {
	panic("hiservices: symbol kAXInterfaceClassicInvertColorStatusDidChangeNotification has opaque signature; use KAXInterfaceClassicInvertColorStatusDidChangeNotificationSymbol() and a typed manual wrapper")
}

// KAXInterfaceClassicInvertColorStatusDidChangeNotificationSymbol returns the raw symbol address for kAXInterfaceClassicInvertColorStatusDidChangeNotification.
func KAXInterfaceClassicInvertColorStatusDidChangeNotificationSymbol() uintptr {
	if _kAXInterfaceClassicInvertColorStatusDidChangeNotificationSymbol == 0 {
		return 0
	}
	return _kAXInterfaceClassicInvertColorStatusDidChangeNotificationSymbol
}

var _kAXInterfaceCursorStatusDidChangeNotificationSymbol uintptr

// KAXInterfaceCursorStatusDidChangeNotification has an opaque C signature in discovered metadata.
// Call KAXInterfaceCursorStatusDidChangeNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceCursorStatusDidChangeNotification
func KAXInterfaceCursorStatusDidChangeNotification() {
	panic("hiservices: symbol kAXInterfaceCursorStatusDidChangeNotification has opaque signature; use KAXInterfaceCursorStatusDidChangeNotificationSymbol() and a typed manual wrapper")
}

// KAXInterfaceCursorStatusDidChangeNotificationSymbol returns the raw symbol address for kAXInterfaceCursorStatusDidChangeNotification.
func KAXInterfaceCursorStatusDidChangeNotificationSymbol() uintptr {
	if _kAXInterfaceCursorStatusDidChangeNotificationSymbol == 0 {
		return 0
	}
	return _kAXInterfaceCursorStatusDidChangeNotificationSymbol
}

var _kAXInterfaceDifferentiateWithoutColorKeySymbol uintptr

// KAXInterfaceDifferentiateWithoutColorKey has an opaque C signature in discovered metadata.
// Call KAXInterfaceDifferentiateWithoutColorKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceDifferentiateWithoutColorKey
func KAXInterfaceDifferentiateWithoutColorKey() {
	panic("hiservices: symbol kAXInterfaceDifferentiateWithoutColorKey has opaque signature; use KAXInterfaceDifferentiateWithoutColorKeySymbol() and a typed manual wrapper")
}

// KAXInterfaceDifferentiateWithoutColorKeySymbol returns the raw symbol address for kAXInterfaceDifferentiateWithoutColorKey.
func KAXInterfaceDifferentiateWithoutColorKeySymbol() uintptr {
	if _kAXInterfaceDifferentiateWithoutColorKeySymbol == 0 {
		return 0
	}
	return _kAXInterfaceDifferentiateWithoutColorKeySymbol
}

var _kAXInterfaceDifferentiateWithoutColorStatusDidChangeNotificationSymbol uintptr

// KAXInterfaceDifferentiateWithoutColorStatusDidChangeNotification has an opaque C signature in discovered metadata.
// Call KAXInterfaceDifferentiateWithoutColorStatusDidChangeNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceDifferentiateWithoutColorStatusDidChangeNotification
func KAXInterfaceDifferentiateWithoutColorStatusDidChangeNotification() {
	panic("hiservices: symbol kAXInterfaceDifferentiateWithoutColorStatusDidChangeNotification has opaque signature; use KAXInterfaceDifferentiateWithoutColorStatusDidChangeNotificationSymbol() and a typed manual wrapper")
}

// KAXInterfaceDifferentiateWithoutColorStatusDidChangeNotificationSymbol returns the raw symbol address for kAXInterfaceDifferentiateWithoutColorStatusDidChangeNotification.
func KAXInterfaceDifferentiateWithoutColorStatusDidChangeNotificationSymbol() uintptr {
	if _kAXInterfaceDifferentiateWithoutColorStatusDidChangeNotificationSymbol == 0 {
		return 0
	}
	return _kAXInterfaceDifferentiateWithoutColorStatusDidChangeNotificationSymbol
}

var _kAXInterfaceIncreaseContrastKeySymbol uintptr

// KAXInterfaceIncreaseContrastKey has an opaque C signature in discovered metadata.
// Call KAXInterfaceIncreaseContrastKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceIncreaseContrastKey
func KAXInterfaceIncreaseContrastKey() {
	panic("hiservices: symbol kAXInterfaceIncreaseContrastKey has opaque signature; use KAXInterfaceIncreaseContrastKeySymbol() and a typed manual wrapper")
}

// KAXInterfaceIncreaseContrastKeySymbol returns the raw symbol address for kAXInterfaceIncreaseContrastKey.
func KAXInterfaceIncreaseContrastKeySymbol() uintptr {
	if _kAXInterfaceIncreaseContrastKeySymbol == 0 {
		return 0
	}
	return _kAXInterfaceIncreaseContrastKeySymbol
}

var _kAXInterfaceIncreaseContrastStatusDidChangeNotificationSymbol uintptr

// KAXInterfaceIncreaseContrastStatusDidChangeNotification has an opaque C signature in discovered metadata.
// Call KAXInterfaceIncreaseContrastStatusDidChangeNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceIncreaseContrastStatusDidChangeNotification
func KAXInterfaceIncreaseContrastStatusDidChangeNotification() {
	panic("hiservices: symbol kAXInterfaceIncreaseContrastStatusDidChangeNotification has opaque signature; use KAXInterfaceIncreaseContrastStatusDidChangeNotificationSymbol() and a typed manual wrapper")
}

// KAXInterfaceIncreaseContrastStatusDidChangeNotificationSymbol returns the raw symbol address for kAXInterfaceIncreaseContrastStatusDidChangeNotification.
func KAXInterfaceIncreaseContrastStatusDidChangeNotificationSymbol() uintptr {
	if _kAXInterfaceIncreaseContrastStatusDidChangeNotificationSymbol == 0 {
		return 0
	}
	return _kAXInterfaceIncreaseContrastStatusDidChangeNotificationSymbol
}

var _kAXInterfaceReduceMotionKeySymbol uintptr

// KAXInterfaceReduceMotionKey has an opaque C signature in discovered metadata.
// Call KAXInterfaceReduceMotionKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceReduceMotionKey
func KAXInterfaceReduceMotionKey() {
	panic("hiservices: symbol kAXInterfaceReduceMotionKey has opaque signature; use KAXInterfaceReduceMotionKeySymbol() and a typed manual wrapper")
}

// KAXInterfaceReduceMotionKeySymbol returns the raw symbol address for kAXInterfaceReduceMotionKey.
func KAXInterfaceReduceMotionKeySymbol() uintptr {
	if _kAXInterfaceReduceMotionKeySymbol == 0 {
		return 0
	}
	return _kAXInterfaceReduceMotionKeySymbol
}

var _kAXInterfaceReduceMotionStatusDidChangeNotificationSymbol uintptr

// KAXInterfaceReduceMotionStatusDidChangeNotification has an opaque C signature in discovered metadata.
// Call KAXInterfaceReduceMotionStatusDidChangeNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceReduceMotionStatusDidChangeNotification
func KAXInterfaceReduceMotionStatusDidChangeNotification() {
	panic("hiservices: symbol kAXInterfaceReduceMotionStatusDidChangeNotification has opaque signature; use KAXInterfaceReduceMotionStatusDidChangeNotificationSymbol() and a typed manual wrapper")
}

// KAXInterfaceReduceMotionStatusDidChangeNotificationSymbol returns the raw symbol address for kAXInterfaceReduceMotionStatusDidChangeNotification.
func KAXInterfaceReduceMotionStatusDidChangeNotificationSymbol() uintptr {
	if _kAXInterfaceReduceMotionStatusDidChangeNotificationSymbol == 0 {
		return 0
	}
	return _kAXInterfaceReduceMotionStatusDidChangeNotificationSymbol
}

var _kAXInterfaceReduceTextInsertionPointModulationDidChangeNotificationSymbol uintptr

// KAXInterfaceReduceTextInsertionPointModulationDidChangeNotification has an opaque C signature in discovered metadata.
// Call KAXInterfaceReduceTextInsertionPointModulationDidChangeNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceReduceTextInsertionPointModulationDidChangeNotification
func KAXInterfaceReduceTextInsertionPointModulationDidChangeNotification() {
	panic("hiservices: symbol kAXInterfaceReduceTextInsertionPointModulationDidChangeNotification has opaque signature; use KAXInterfaceReduceTextInsertionPointModulationDidChangeNotificationSymbol() and a typed manual wrapper")
}

// KAXInterfaceReduceTextInsertionPointModulationDidChangeNotificationSymbol returns the raw symbol address for kAXInterfaceReduceTextInsertionPointModulationDidChangeNotification.
func KAXInterfaceReduceTextInsertionPointModulationDidChangeNotificationSymbol() uintptr {
	if _kAXInterfaceReduceTextInsertionPointModulationDidChangeNotificationSymbol == 0 {
		return 0
	}
	return _kAXInterfaceReduceTextInsertionPointModulationDidChangeNotificationSymbol
}

var _kAXInterfaceReduceTextInsertionPointModulationKeySymbol uintptr

// KAXInterfaceReduceTextInsertionPointModulationKey has an opaque C signature in discovered metadata.
// Call KAXInterfaceReduceTextInsertionPointModulationKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceReduceTextInsertionPointModulationKey
func KAXInterfaceReduceTextInsertionPointModulationKey() {
	panic("hiservices: symbol kAXInterfaceReduceTextInsertionPointModulationKey has opaque signature; use KAXInterfaceReduceTextInsertionPointModulationKeySymbol() and a typed manual wrapper")
}

// KAXInterfaceReduceTextInsertionPointModulationKeySymbol returns the raw symbol address for kAXInterfaceReduceTextInsertionPointModulationKey.
func KAXInterfaceReduceTextInsertionPointModulationKeySymbol() uintptr {
	if _kAXInterfaceReduceTextInsertionPointModulationKeySymbol == 0 {
		return 0
	}
	return _kAXInterfaceReduceTextInsertionPointModulationKeySymbol
}

var _kAXInterfaceReduceTransparencyKeySymbol uintptr

// KAXInterfaceReduceTransparencyKey has an opaque C signature in discovered metadata.
// Call KAXInterfaceReduceTransparencyKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceReduceTransparencyKey
func KAXInterfaceReduceTransparencyKey() {
	panic("hiservices: symbol kAXInterfaceReduceTransparencyKey has opaque signature; use KAXInterfaceReduceTransparencyKeySymbol() and a typed manual wrapper")
}

// KAXInterfaceReduceTransparencyKeySymbol returns the raw symbol address for kAXInterfaceReduceTransparencyKey.
func KAXInterfaceReduceTransparencyKeySymbol() uintptr {
	if _kAXInterfaceReduceTransparencyKeySymbol == 0 {
		return 0
	}
	return _kAXInterfaceReduceTransparencyKeySymbol
}

var _kAXInterfaceReduceTransparencyStatusDidChangeNotificationSymbol uintptr

// KAXInterfaceReduceTransparencyStatusDidChangeNotification has an opaque C signature in discovered metadata.
// Call KAXInterfaceReduceTransparencyStatusDidChangeNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceReduceTransparencyStatusDidChangeNotification
func KAXInterfaceReduceTransparencyStatusDidChangeNotification() {
	panic("hiservices: symbol kAXInterfaceReduceTransparencyStatusDidChangeNotification has opaque signature; use KAXInterfaceReduceTransparencyStatusDidChangeNotificationSymbol() and a typed manual wrapper")
}

// KAXInterfaceReduceTransparencyStatusDidChangeNotificationSymbol returns the raw symbol address for kAXInterfaceReduceTransparencyStatusDidChangeNotification.
func KAXInterfaceReduceTransparencyStatusDidChangeNotificationSymbol() uintptr {
	if _kAXInterfaceReduceTransparencyStatusDidChangeNotificationSymbol == 0 {
		return 0
	}
	return _kAXInterfaceReduceTransparencyStatusDidChangeNotificationSymbol
}

var _kAXInterfaceShowToolbarButtonShapesKeySymbol uintptr

// KAXInterfaceShowToolbarButtonShapesKey has an opaque C signature in discovered metadata.
// Call KAXInterfaceShowToolbarButtonShapesKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceShowToolbarButtonShapesKey
func KAXInterfaceShowToolbarButtonShapesKey() {
	panic("hiservices: symbol kAXInterfaceShowToolbarButtonShapesKey has opaque signature; use KAXInterfaceShowToolbarButtonShapesKeySymbol() and a typed manual wrapper")
}

// KAXInterfaceShowToolbarButtonShapesKeySymbol returns the raw symbol address for kAXInterfaceShowToolbarButtonShapesKey.
func KAXInterfaceShowToolbarButtonShapesKeySymbol() uintptr {
	if _kAXInterfaceShowToolbarButtonShapesKeySymbol == 0 {
		return 0
	}
	return _kAXInterfaceShowToolbarButtonShapesKeySymbol
}

var _kAXInterfaceShowToolbarButtonShapesStatusDidChangeNotificationSymbol uintptr

// KAXInterfaceShowToolbarButtonShapesStatusDidChangeNotification has an opaque C signature in discovered metadata.
// Call KAXInterfaceShowToolbarButtonShapesStatusDidChangeNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceShowToolbarButtonShapesStatusDidChangeNotification
func KAXInterfaceShowToolbarButtonShapesStatusDidChangeNotification() {
	panic("hiservices: symbol kAXInterfaceShowToolbarButtonShapesStatusDidChangeNotification has opaque signature; use KAXInterfaceShowToolbarButtonShapesStatusDidChangeNotificationSymbol() and a typed manual wrapper")
}

// KAXInterfaceShowToolbarButtonShapesStatusDidChangeNotificationSymbol returns the raw symbol address for kAXInterfaceShowToolbarButtonShapesStatusDidChangeNotification.
func KAXInterfaceShowToolbarButtonShapesStatusDidChangeNotificationSymbol() uintptr {
	if _kAXInterfaceShowToolbarButtonShapesStatusDidChangeNotificationSymbol == 0 {
		return 0
	}
	return _kAXInterfaceShowToolbarButtonShapesStatusDidChangeNotificationSymbol
}

var _kAXInterfaceShowWindowTitlebarIconsKeySymbol uintptr

// KAXInterfaceShowWindowTitlebarIconsKey has an opaque C signature in discovered metadata.
// Call KAXInterfaceShowWindowTitlebarIconsKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceShowWindowTitlebarIconsKey
func KAXInterfaceShowWindowTitlebarIconsKey() {
	panic("hiservices: symbol kAXInterfaceShowWindowTitlebarIconsKey has opaque signature; use KAXInterfaceShowWindowTitlebarIconsKeySymbol() and a typed manual wrapper")
}

// KAXInterfaceShowWindowTitlebarIconsKeySymbol returns the raw symbol address for kAXInterfaceShowWindowTitlebarIconsKey.
func KAXInterfaceShowWindowTitlebarIconsKeySymbol() uintptr {
	if _kAXInterfaceShowWindowTitlebarIconsKeySymbol == 0 {
		return 0
	}
	return _kAXInterfaceShowWindowTitlebarIconsKeySymbol
}

var _kAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotificationSymbol uintptr

// KAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotification has an opaque C signature in discovered metadata.
// Call KAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotificationSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotification
func KAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotification() {
	panic("hiservices: symbol kAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotification has opaque signature; use KAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotificationSymbol() and a typed manual wrapper")
}

// KAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotificationSymbol returns the raw symbol address for kAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotification.
func KAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotificationSymbol() uintptr {
	if _kAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotificationSymbol == 0 {
		return 0
	}
	return _kAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotificationSymbol
}

var _kAXInterfaceWhiteOnBlackKeySymbol uintptr

// KAXInterfaceWhiteOnBlackKey has an opaque C signature in discovered metadata.
// Call KAXInterfaceWhiteOnBlackKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXInterfaceWhiteOnBlackKey
func KAXInterfaceWhiteOnBlackKey() {
	panic("hiservices: symbol kAXInterfaceWhiteOnBlackKey has opaque signature; use KAXInterfaceWhiteOnBlackKeySymbol() and a typed manual wrapper")
}

// KAXInterfaceWhiteOnBlackKeySymbol returns the raw symbol address for kAXInterfaceWhiteOnBlackKey.
func KAXInterfaceWhiteOnBlackKeySymbol() uintptr {
	if _kAXInterfaceWhiteOnBlackKeySymbol == 0 {
		return 0
	}
	return _kAXInterfaceWhiteOnBlackKeySymbol
}

var _kAXLinkTextAttributeSymbol uintptr

// KAXLinkTextAttribute has an opaque C signature in discovered metadata.
// Call KAXLinkTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXLinkTextAttribute
func KAXLinkTextAttribute() {
	panic("hiservices: symbol kAXLinkTextAttribute has opaque signature; use KAXLinkTextAttributeSymbol() and a typed manual wrapper")
}

// KAXLinkTextAttributeSymbol returns the raw symbol address for kAXLinkTextAttribute.
func KAXLinkTextAttributeSymbol() uintptr {
	if _kAXLinkTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXLinkTextAttributeSymbol
}

var _kAXListItemIndexTextAttributeSymbol uintptr

// KAXListItemIndexTextAttribute has an opaque C signature in discovered metadata.
// Call KAXListItemIndexTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXListItemIndexTextAttribute
func KAXListItemIndexTextAttribute() {
	panic("hiservices: symbol kAXListItemIndexTextAttribute has opaque signature; use KAXListItemIndexTextAttributeSymbol() and a typed manual wrapper")
}

// KAXListItemIndexTextAttributeSymbol returns the raw symbol address for kAXListItemIndexTextAttribute.
func KAXListItemIndexTextAttributeSymbol() uintptr {
	if _kAXListItemIndexTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXListItemIndexTextAttributeSymbol
}

var _kAXListItemLevelTextAttributeSymbol uintptr

// KAXListItemLevelTextAttribute has an opaque C signature in discovered metadata.
// Call KAXListItemLevelTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXListItemLevelTextAttribute
func KAXListItemLevelTextAttribute() {
	panic("hiservices: symbol kAXListItemLevelTextAttribute has opaque signature; use KAXListItemLevelTextAttributeSymbol() and a typed manual wrapper")
}

// KAXListItemLevelTextAttributeSymbol returns the raw symbol address for kAXListItemLevelTextAttribute.
func KAXListItemLevelTextAttributeSymbol() uintptr {
	if _kAXListItemLevelTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXListItemLevelTextAttributeSymbol
}

var _kAXListItemPrefixTextAttributeSymbol uintptr

// KAXListItemPrefixTextAttribute has an opaque C signature in discovered metadata.
// Call KAXListItemPrefixTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXListItemPrefixTextAttribute
func KAXListItemPrefixTextAttribute() {
	panic("hiservices: symbol kAXListItemPrefixTextAttribute has opaque signature; use KAXListItemPrefixTextAttributeSymbol() and a typed manual wrapper")
}

// KAXListItemPrefixTextAttributeSymbol returns the raw symbol address for kAXListItemPrefixTextAttribute.
func KAXListItemPrefixTextAttributeSymbol() uintptr {
	if _kAXListItemPrefixTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXListItemPrefixTextAttributeSymbol
}

var _kAXMarkedMisspelledTextAttributeSymbol uintptr

// KAXMarkedMisspelledTextAttribute has an opaque C signature in discovered metadata.
// Call KAXMarkedMisspelledTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXMarkedMisspelledTextAttribute
func KAXMarkedMisspelledTextAttribute() {
	panic("hiservices: symbol kAXMarkedMisspelledTextAttribute has opaque signature; use KAXMarkedMisspelledTextAttributeSymbol() and a typed manual wrapper")
}

// KAXMarkedMisspelledTextAttributeSymbol returns the raw symbol address for kAXMarkedMisspelledTextAttribute.
func KAXMarkedMisspelledTextAttributeSymbol() uintptr {
	if _kAXMarkedMisspelledTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXMarkedMisspelledTextAttributeSymbol
}

var _kAXMisspelledTextAttributeSymbol uintptr

// KAXMisspelledTextAttribute has an opaque C signature in discovered metadata.
// Call KAXMisspelledTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXMisspelledTextAttribute
func KAXMisspelledTextAttribute() {
	panic("hiservices: symbol kAXMisspelledTextAttribute has opaque signature; use KAXMisspelledTextAttributeSymbol() and a typed manual wrapper")
}

// KAXMisspelledTextAttributeSymbol returns the raw symbol address for kAXMisspelledTextAttribute.
func KAXMisspelledTextAttributeSymbol() uintptr {
	if _kAXMisspelledTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXMisspelledTextAttributeSymbol
}

var _kAXNaturalLanguageTextAttributeSymbol uintptr

// KAXNaturalLanguageTextAttribute has an opaque C signature in discovered metadata.
// Call KAXNaturalLanguageTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXNaturalLanguageTextAttribute
func KAXNaturalLanguageTextAttribute() {
	panic("hiservices: symbol kAXNaturalLanguageTextAttribute has opaque signature; use KAXNaturalLanguageTextAttributeSymbol() and a typed manual wrapper")
}

// KAXNaturalLanguageTextAttributeSymbol returns the raw symbol address for kAXNaturalLanguageTextAttribute.
func KAXNaturalLanguageTextAttributeSymbol() uintptr {
	if _kAXNaturalLanguageTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXNaturalLanguageTextAttributeSymbol
}

var _kAXParagraphStyleTextAttributeSymbol uintptr

// KAXParagraphStyleTextAttribute has an opaque C signature in discovered metadata.
// Call KAXParagraphStyleTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXParagraphStyleTextAttribute
func KAXParagraphStyleTextAttribute() {
	panic("hiservices: symbol kAXParagraphStyleTextAttribute has opaque signature; use KAXParagraphStyleTextAttributeSymbol() and a typed manual wrapper")
}

// KAXParagraphStyleTextAttributeSymbol returns the raw symbol address for kAXParagraphStyleTextAttribute.
func KAXParagraphStyleTextAttributeSymbol() uintptr {
	if _kAXParagraphStyleTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXParagraphStyleTextAttributeSymbol
}

var _kAXReplacementStringTextAttributeSymbol uintptr

// KAXReplacementStringTextAttribute has an opaque C signature in discovered metadata.
// Call KAXReplacementStringTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXReplacementStringTextAttribute
func KAXReplacementStringTextAttribute() {
	panic("hiservices: symbol kAXReplacementStringTextAttribute has opaque signature; use KAXReplacementStringTextAttributeSymbol() and a typed manual wrapper")
}

// KAXReplacementStringTextAttributeSymbol returns the raw symbol address for kAXReplacementStringTextAttribute.
func KAXReplacementStringTextAttributeSymbol() uintptr {
	if _kAXReplacementStringTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXReplacementStringTextAttributeSymbol
}

var _kAXShadowTextAttributeSymbol uintptr

// KAXShadowTextAttribute has an opaque C signature in discovered metadata.
// Call KAXShadowTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXShadowTextAttribute
func KAXShadowTextAttribute() {
	panic("hiservices: symbol kAXShadowTextAttribute has opaque signature; use KAXShadowTextAttributeSymbol() and a typed manual wrapper")
}

// KAXShadowTextAttributeSymbol returns the raw symbol address for kAXShadowTextAttribute.
func KAXShadowTextAttributeSymbol() uintptr {
	if _kAXShadowTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXShadowTextAttributeSymbol
}

var _kAXStrikethroughColorTextAttributeSymbol uintptr

// KAXStrikethroughColorTextAttribute has an opaque C signature in discovered metadata.
// Call KAXStrikethroughColorTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXStrikethroughColorTextAttribute
func KAXStrikethroughColorTextAttribute() {
	panic("hiservices: symbol kAXStrikethroughColorTextAttribute has opaque signature; use KAXStrikethroughColorTextAttributeSymbol() and a typed manual wrapper")
}

// KAXStrikethroughColorTextAttributeSymbol returns the raw symbol address for kAXStrikethroughColorTextAttribute.
func KAXStrikethroughColorTextAttributeSymbol() uintptr {
	if _kAXStrikethroughColorTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXStrikethroughColorTextAttributeSymbol
}

var _kAXStrikethroughTextAttributeSymbol uintptr

// KAXStrikethroughTextAttribute has an opaque C signature in discovered metadata.
// Call KAXStrikethroughTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXStrikethroughTextAttribute
func KAXStrikethroughTextAttribute() {
	panic("hiservices: symbol kAXStrikethroughTextAttribute has opaque signature; use KAXStrikethroughTextAttributeSymbol() and a typed manual wrapper")
}

// KAXStrikethroughTextAttributeSymbol returns the raw symbol address for kAXStrikethroughTextAttribute.
func KAXStrikethroughTextAttributeSymbol() uintptr {
	if _kAXStrikethroughTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXStrikethroughTextAttributeSymbol
}

var _kAXSuperscriptTextAttributeSymbol uintptr

// KAXSuperscriptTextAttribute has an opaque C signature in discovered metadata.
// Call KAXSuperscriptTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXSuperscriptTextAttribute
func KAXSuperscriptTextAttribute() {
	panic("hiservices: symbol kAXSuperscriptTextAttribute has opaque signature; use KAXSuperscriptTextAttributeSymbol() and a typed manual wrapper")
}

// KAXSuperscriptTextAttributeSymbol returns the raw symbol address for kAXSuperscriptTextAttribute.
func KAXSuperscriptTextAttributeSymbol() uintptr {
	if _kAXSuperscriptTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXSuperscriptTextAttributeSymbol
}

var _kAXTextAlignmentAttributeSymbol uintptr

// KAXTextAlignmentAttribute has an opaque C signature in discovered metadata.
// Call KAXTextAlignmentAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXTextAlignmentAttribute
func KAXTextAlignmentAttribute() {
	panic("hiservices: symbol kAXTextAlignmentAttribute has opaque signature; use KAXTextAlignmentAttributeSymbol() and a typed manual wrapper")
}

// KAXTextAlignmentAttributeSymbol returns the raw symbol address for kAXTextAlignmentAttribute.
func KAXTextAlignmentAttributeSymbol() uintptr {
	if _kAXTextAlignmentAttributeSymbol == 0 {
		return 0
	}
	return _kAXTextAlignmentAttributeSymbol
}

var _kAXTextAlignmentKeySymbol uintptr

// KAXTextAlignmentKey has an opaque C signature in discovered metadata.
// Call KAXTextAlignmentKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXTextAlignmentKey
func KAXTextAlignmentKey() {
	panic("hiservices: symbol kAXTextAlignmentKey has opaque signature; use KAXTextAlignmentKeySymbol() and a typed manual wrapper")
}

// KAXTextAlignmentKeySymbol returns the raw symbol address for kAXTextAlignmentKey.
func KAXTextAlignmentKeySymbol() uintptr {
	if _kAXTextAlignmentKeySymbol == 0 {
		return 0
	}
	return _kAXTextAlignmentKeySymbol
}

var _kAXTrustedCheckOptionPromptSymbol uintptr

// KAXTrustedCheckOptionPrompt has an opaque C signature in discovered metadata.
// Call KAXTrustedCheckOptionPromptSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXTrustedCheckOptionPrompt
func KAXTrustedCheckOptionPrompt() {
	panic("hiservices: symbol kAXTrustedCheckOptionPrompt has opaque signature; use KAXTrustedCheckOptionPromptSymbol() and a typed manual wrapper")
}

// KAXTrustedCheckOptionPromptSymbol returns the raw symbol address for kAXTrustedCheckOptionPrompt.
func KAXTrustedCheckOptionPromptSymbol() uintptr {
	if _kAXTrustedCheckOptionPromptSymbol == 0 {
		return 0
	}
	return _kAXTrustedCheckOptionPromptSymbol
}

var _kAXUIElementCopyHierarchyArrayAttributesKeySymbol uintptr

// KAXUIElementCopyHierarchyArrayAttributesKey has an opaque C signature in discovered metadata.
// Call KAXUIElementCopyHierarchyArrayAttributesKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUIElementCopyHierarchyArrayAttributesKey
func KAXUIElementCopyHierarchyArrayAttributesKey() {
	panic("hiservices: symbol kAXUIElementCopyHierarchyArrayAttributesKey has opaque signature; use KAXUIElementCopyHierarchyArrayAttributesKeySymbol() and a typed manual wrapper")
}

// KAXUIElementCopyHierarchyArrayAttributesKeySymbol returns the raw symbol address for kAXUIElementCopyHierarchyArrayAttributesKey.
func KAXUIElementCopyHierarchyArrayAttributesKeySymbol() uintptr {
	if _kAXUIElementCopyHierarchyArrayAttributesKeySymbol == 0 {
		return 0
	}
	return _kAXUIElementCopyHierarchyArrayAttributesKeySymbol
}

var _kAXUIElementCopyHierarchyIncompleteResultKeySymbol uintptr

// KAXUIElementCopyHierarchyIncompleteResultKey has an opaque C signature in discovered metadata.
// Call KAXUIElementCopyHierarchyIncompleteResultKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUIElementCopyHierarchyIncompleteResultKey
func KAXUIElementCopyHierarchyIncompleteResultKey() {
	panic("hiservices: symbol kAXUIElementCopyHierarchyIncompleteResultKey has opaque signature; use KAXUIElementCopyHierarchyIncompleteResultKeySymbol() and a typed manual wrapper")
}

// KAXUIElementCopyHierarchyIncompleteResultKeySymbol returns the raw symbol address for kAXUIElementCopyHierarchyIncompleteResultKey.
func KAXUIElementCopyHierarchyIncompleteResultKeySymbol() uintptr {
	if _kAXUIElementCopyHierarchyIncompleteResultKeySymbol == 0 {
		return 0
	}
	return _kAXUIElementCopyHierarchyIncompleteResultKeySymbol
}

var _kAXUIElementCopyHierarchyMaxArrayCountKeySymbol uintptr

// KAXUIElementCopyHierarchyMaxArrayCountKey has an opaque C signature in discovered metadata.
// Call KAXUIElementCopyHierarchyMaxArrayCountKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUIElementCopyHierarchyMaxArrayCountKey
func KAXUIElementCopyHierarchyMaxArrayCountKey() {
	panic("hiservices: symbol kAXUIElementCopyHierarchyMaxArrayCountKey has opaque signature; use KAXUIElementCopyHierarchyMaxArrayCountKeySymbol() and a typed manual wrapper")
}

// KAXUIElementCopyHierarchyMaxArrayCountKeySymbol returns the raw symbol address for kAXUIElementCopyHierarchyMaxArrayCountKey.
func KAXUIElementCopyHierarchyMaxArrayCountKeySymbol() uintptr {
	if _kAXUIElementCopyHierarchyMaxArrayCountKeySymbol == 0 {
		return 0
	}
	return _kAXUIElementCopyHierarchyMaxArrayCountKeySymbol
}

var _kAXUIElementCopyHierarchyMaxDepthKeySymbol uintptr

// KAXUIElementCopyHierarchyMaxDepthKey has an opaque C signature in discovered metadata.
// Call KAXUIElementCopyHierarchyMaxDepthKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUIElementCopyHierarchyMaxDepthKey
func KAXUIElementCopyHierarchyMaxDepthKey() {
	panic("hiservices: symbol kAXUIElementCopyHierarchyMaxDepthKey has opaque signature; use KAXUIElementCopyHierarchyMaxDepthKeySymbol() and a typed manual wrapper")
}

// KAXUIElementCopyHierarchyMaxDepthKeySymbol returns the raw symbol address for kAXUIElementCopyHierarchyMaxDepthKey.
func KAXUIElementCopyHierarchyMaxDepthKeySymbol() uintptr {
	if _kAXUIElementCopyHierarchyMaxDepthKeySymbol == 0 {
		return 0
	}
	return _kAXUIElementCopyHierarchyMaxDepthKeySymbol
}

var _kAXUIElementCopyHierarchyResultCountKeySymbol uintptr

// KAXUIElementCopyHierarchyResultCountKey has an opaque C signature in discovered metadata.
// Call KAXUIElementCopyHierarchyResultCountKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUIElementCopyHierarchyResultCountKey
func KAXUIElementCopyHierarchyResultCountKey() {
	panic("hiservices: symbol kAXUIElementCopyHierarchyResultCountKey has opaque signature; use KAXUIElementCopyHierarchyResultCountKeySymbol() and a typed manual wrapper")
}

// KAXUIElementCopyHierarchyResultCountKeySymbol returns the raw symbol address for kAXUIElementCopyHierarchyResultCountKey.
func KAXUIElementCopyHierarchyResultCountKeySymbol() uintptr {
	if _kAXUIElementCopyHierarchyResultCountKeySymbol == 0 {
		return 0
	}
	return _kAXUIElementCopyHierarchyResultCountKeySymbol
}

var _kAXUIElementCopyHierarchyResultErrorKeySymbol uintptr

// KAXUIElementCopyHierarchyResultErrorKey has an opaque C signature in discovered metadata.
// Call KAXUIElementCopyHierarchyResultErrorKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUIElementCopyHierarchyResultErrorKey
func KAXUIElementCopyHierarchyResultErrorKey() {
	panic("hiservices: symbol kAXUIElementCopyHierarchyResultErrorKey has opaque signature; use KAXUIElementCopyHierarchyResultErrorKeySymbol() and a typed manual wrapper")
}

// KAXUIElementCopyHierarchyResultErrorKeySymbol returns the raw symbol address for kAXUIElementCopyHierarchyResultErrorKey.
func KAXUIElementCopyHierarchyResultErrorKeySymbol() uintptr {
	if _kAXUIElementCopyHierarchyResultErrorKeySymbol == 0 {
		return 0
	}
	return _kAXUIElementCopyHierarchyResultErrorKeySymbol
}

var _kAXUIElementCopyHierarchyResultValueKeySymbol uintptr

// KAXUIElementCopyHierarchyResultValueKey has an opaque C signature in discovered metadata.
// Call KAXUIElementCopyHierarchyResultValueKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUIElementCopyHierarchyResultValueKey
func KAXUIElementCopyHierarchyResultValueKey() {
	panic("hiservices: symbol kAXUIElementCopyHierarchyResultValueKey has opaque signature; use KAXUIElementCopyHierarchyResultValueKeySymbol() and a typed manual wrapper")
}

// KAXUIElementCopyHierarchyResultValueKeySymbol returns the raw symbol address for kAXUIElementCopyHierarchyResultValueKey.
func KAXUIElementCopyHierarchyResultValueKeySymbol() uintptr {
	if _kAXUIElementCopyHierarchyResultValueKeySymbol == 0 {
		return 0
	}
	return _kAXUIElementCopyHierarchyResultValueKeySymbol
}

var _kAXUIElementCopyHierarchyReturnAttributeErrorsKeySymbol uintptr

// KAXUIElementCopyHierarchyReturnAttributeErrorsKey has an opaque C signature in discovered metadata.
// Call KAXUIElementCopyHierarchyReturnAttributeErrorsKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUIElementCopyHierarchyReturnAttributeErrorsKey
func KAXUIElementCopyHierarchyReturnAttributeErrorsKey() {
	panic("hiservices: symbol kAXUIElementCopyHierarchyReturnAttributeErrorsKey has opaque signature; use KAXUIElementCopyHierarchyReturnAttributeErrorsKeySymbol() and a typed manual wrapper")
}

// KAXUIElementCopyHierarchyReturnAttributeErrorsKeySymbol returns the raw symbol address for kAXUIElementCopyHierarchyReturnAttributeErrorsKey.
func KAXUIElementCopyHierarchyReturnAttributeErrorsKeySymbol() uintptr {
	if _kAXUIElementCopyHierarchyReturnAttributeErrorsKeySymbol == 0 {
		return 0
	}
	return _kAXUIElementCopyHierarchyReturnAttributeErrorsKeySymbol
}

var _kAXUIElementCopyHierarchySkipInspectionForAttributesKeySymbol uintptr

// KAXUIElementCopyHierarchySkipInspectionForAttributesKey has an opaque C signature in discovered metadata.
// Call KAXUIElementCopyHierarchySkipInspectionForAttributesKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUIElementCopyHierarchySkipInspectionForAttributesKey
func KAXUIElementCopyHierarchySkipInspectionForAttributesKey() {
	panic("hiservices: symbol kAXUIElementCopyHierarchySkipInspectionForAttributesKey has opaque signature; use KAXUIElementCopyHierarchySkipInspectionForAttributesKeySymbol() and a typed manual wrapper")
}

// KAXUIElementCopyHierarchySkipInspectionForAttributesKeySymbol returns the raw symbol address for kAXUIElementCopyHierarchySkipInspectionForAttributesKey.
func KAXUIElementCopyHierarchySkipInspectionForAttributesKeySymbol() uintptr {
	if _kAXUIElementCopyHierarchySkipInspectionForAttributesKeySymbol == 0 {
		return 0
	}
	return _kAXUIElementCopyHierarchySkipInspectionForAttributesKeySymbol
}

var _kAXUIElementCopyHierarchyTruncateStringsKeySymbol uintptr

// KAXUIElementCopyHierarchyTruncateStringsKey has an opaque C signature in discovered metadata.
// Call KAXUIElementCopyHierarchyTruncateStringsKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUIElementCopyHierarchyTruncateStringsKey
func KAXUIElementCopyHierarchyTruncateStringsKey() {
	panic("hiservices: symbol kAXUIElementCopyHierarchyTruncateStringsKey has opaque signature; use KAXUIElementCopyHierarchyTruncateStringsKeySymbol() and a typed manual wrapper")
}

// KAXUIElementCopyHierarchyTruncateStringsKeySymbol returns the raw symbol address for kAXUIElementCopyHierarchyTruncateStringsKey.
func KAXUIElementCopyHierarchyTruncateStringsKeySymbol() uintptr {
	if _kAXUIElementCopyHierarchyTruncateStringsKeySymbol == 0 {
		return 0
	}
	return _kAXUIElementCopyHierarchyTruncateStringsKeySymbol
}

var _kAXUnderlineColorTextAttributeSymbol uintptr

// KAXUnderlineColorTextAttribute has an opaque C signature in discovered metadata.
// Call KAXUnderlineColorTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUnderlineColorTextAttribute
func KAXUnderlineColorTextAttribute() {
	panic("hiservices: symbol kAXUnderlineColorTextAttribute has opaque signature; use KAXUnderlineColorTextAttributeSymbol() and a typed manual wrapper")
}

// KAXUnderlineColorTextAttributeSymbol returns the raw symbol address for kAXUnderlineColorTextAttribute.
func KAXUnderlineColorTextAttributeSymbol() uintptr {
	if _kAXUnderlineColorTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXUnderlineColorTextAttributeSymbol
}

var _kAXUnderlineTextAttributeSymbol uintptr

// KAXUnderlineTextAttribute has an opaque C signature in discovered metadata.
// Call KAXUnderlineTextAttributeSymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXUnderlineTextAttribute
func KAXUnderlineTextAttribute() {
	panic("hiservices: symbol kAXUnderlineTextAttribute has opaque signature; use KAXUnderlineTextAttributeSymbol() and a typed manual wrapper")
}

// KAXUnderlineTextAttributeSymbol returns the raw symbol address for kAXUnderlineTextAttribute.
func KAXUnderlineTextAttributeSymbol() uintptr {
	if _kAXUnderlineTextAttributeSymbol == 0 {
		return 0
	}
	return _kAXUnderlineTextAttributeSymbol
}

var _kAXVisibleNameKeySymbol uintptr

// KAXVisibleNameKey has an opaque C signature in discovered metadata.
// Call KAXVisibleNameKeySymbol to get the raw symbol address.
//
// See: https://developer.apple.com/documentation/HIServices/kAXVisibleNameKey
func KAXVisibleNameKey() {
	panic("hiservices: symbol kAXVisibleNameKey has opaque signature; use KAXVisibleNameKeySymbol() and a typed manual wrapper")
}

// KAXVisibleNameKeySymbol returns the raw symbol address for kAXVisibleNameKey.
func KAXVisibleNameKeySymbol() uintptr {
	if _kAXVisibleNameKeySymbol == 0 {
		return 0
	}
	return _kAXVisibleNameKeySymbol
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerSymbol(&_aXAPIEnabledSymbol, frameworkHandle, "AXAPIEnabled")
		registerSymbol(&_aXIsProcessTrustedSymbol, frameworkHandle, "AXIsProcessTrusted")
		registerSymbol(&_aXIsProcessTrustedWithOptionsSymbol, frameworkHandle, "AXIsProcessTrustedWithOptions")
		registerSymbol(&_aXMakeProcessTrustedSymbol, frameworkHandle, "AXMakeProcessTrusted")
		registerSymbol(&_aXObserverAddNotificationSymbol, frameworkHandle, "AXObserverAddNotification")
		registerSymbol(&_aXObserverAddNotificationAsyncSymbol, frameworkHandle, "AXObserverAddNotificationAsync")
		registerSymbol(&_aXObserverCreateSymbol, frameworkHandle, "AXObserverCreate")
		registerSymbol(&_aXObserverCreateWithInfoCallbackSymbol, frameworkHandle, "AXObserverCreateWithInfoCallback")
		registerSymbol(&_aXObserverGetRunLoopSourceSymbol, frameworkHandle, "AXObserverGetRunLoopSource")
		registerSymbol(&_aXObserverGetTypeIDSymbol, frameworkHandle, "AXObserverGetTypeID")
		registerSymbol(&_aXObserverRemoveNotificationSymbol, frameworkHandle, "AXObserverRemoveNotification")
		registerSymbol(&_aXObserverRemoveNotificationAsyncSymbol, frameworkHandle, "AXObserverRemoveNotificationAsync")
		registerSymbol(&_aXSerializeCFTypeSymbol, frameworkHandle, "AXSerializeCFType")
		registerSymbol(&_aXTextMarkerCreateSymbol, frameworkHandle, "AXTextMarkerCreate")
		registerSymbol(&_aXTextMarkerGetBytePtrSymbol, frameworkHandle, "AXTextMarkerGetBytePtr")
		registerSymbol(&_aXTextMarkerGetLengthSymbol, frameworkHandle, "AXTextMarkerGetLength")
		registerSymbol(&_aXTextMarkerGetTypeIDSymbol, frameworkHandle, "AXTextMarkerGetTypeID")
		registerSymbol(&_aXTextMarkerRangeCopyEndMarkerSymbol, frameworkHandle, "AXTextMarkerRangeCopyEndMarker")
		registerSymbol(&_aXTextMarkerRangeCopyStartMarkerSymbol, frameworkHandle, "AXTextMarkerRangeCopyStartMarker")
		registerSymbol(&_aXTextMarkerRangeCreateSymbol, frameworkHandle, "AXTextMarkerRangeCreate")
		registerSymbol(&_aXTextMarkerRangeCreateWithBytesSymbol, frameworkHandle, "AXTextMarkerRangeCreateWithBytes")
		registerSymbol(&_aXTextMarkerRangeGetTypeIDSymbol, frameworkHandle, "AXTextMarkerRangeGetTypeID")
		registerSymbol(&_aXUIElementCopyActionDescriptionSymbol, frameworkHandle, "AXUIElementCopyActionDescription")
		registerSymbol(&_aXUIElementCopyActionNamesSymbol, frameworkHandle, "AXUIElementCopyActionNames")
		registerSymbol(&_aXUIElementCopyAttributeNamesSymbol, frameworkHandle, "AXUIElementCopyAttributeNames")
		registerSymbol(&_aXUIElementCopyAttributeValueSymbol, frameworkHandle, "AXUIElementCopyAttributeValue")
		registerSymbol(&_aXUIElementCopyAttributeValuesSymbol, frameworkHandle, "AXUIElementCopyAttributeValues")
		registerSymbol(&_aXUIElementCopyElementAtPositionSymbol, frameworkHandle, "AXUIElementCopyElementAtPosition")
		registerSymbol(&_aXUIElementCopyHierarchySymbol, frameworkHandle, "AXUIElementCopyHierarchy")
		registerSymbol(&_aXUIElementCopyMultipleAttributeValuesSymbol, frameworkHandle, "AXUIElementCopyMultipleAttributeValues")
		registerSymbol(&_aXUIElementCopyParameterizedAttributeNamesSymbol, frameworkHandle, "AXUIElementCopyParameterizedAttributeNames")
		registerSymbol(&_aXUIElementCopyParameterizedAttributeValueSymbol, frameworkHandle, "AXUIElementCopyParameterizedAttributeValue")
		registerSymbol(&_aXUIElementCreateApplicationSymbol, frameworkHandle, "AXUIElementCreateApplication")
		registerSymbol(&_aXUIElementCreateSystemWideSymbol, frameworkHandle, "AXUIElementCreateSystemWide")
		registerSymbol(&_aXUIElementGetAttributeValueCountSymbol, frameworkHandle, "AXUIElementGetAttributeValueCount")
		registerSymbol(&_aXUIElementGetPidSymbol, frameworkHandle, "AXUIElementGetPid")
		registerSymbol(&_aXUIElementGetTypeIDSymbol, frameworkHandle, "AXUIElementGetTypeID")
		registerSymbol(&_aXUIElementIsAttributeSettableSymbol, frameworkHandle, "AXUIElementIsAttributeSettable")
		registerSymbol(&_aXUIElementPerformActionSymbol, frameworkHandle, "AXUIElementPerformAction")
		registerSymbol(&_aXUIElementPostKeyboardEventSymbol, frameworkHandle, "AXUIElementPostKeyboardEvent")
		registerSymbol(&_aXUIElementSetAttributeValueSymbol, frameworkHandle, "AXUIElementSetAttributeValue")
		registerSymbol(&_aXUIElementSetMessagingTimeoutSymbol, frameworkHandle, "AXUIElementSetMessagingTimeout")
		registerSymbol(&_aXUnserializeCFTypeSymbol, frameworkHandle, "AXUnserializeCFType")
		registerSymbol(&_aXValueCreateSymbol, frameworkHandle, "AXValueCreate")
		registerSymbol(&_aXValueGetTypeSymbol, frameworkHandle, "AXValueGetType")
		registerSymbol(&_aXValueGetTypeIDSymbol, frameworkHandle, "AXValueGetTypeID")
		registerSymbol(&_aXValueGetValueSymbol, frameworkHandle, "AXValueGetValue")
		registerSymbol(&_applicationTypeGetSymbol, frameworkHandle, "ApplicationTypeGet")
		registerSymbol(&_applicationTypeSetSymbol, frameworkHandle, "ApplicationTypeSet")
		registerSymbol(&_cGPointInIconRefSymbol, frameworkHandle, "CGPointInIconRef")
		registerSymbol(&_cGRectInIconRefSymbol, frameworkHandle, "CGRectInIconRef")
		registerSymbol(&_copyLabelColorAndNameSymbol, frameworkHandle, "CopyLabelColorAndName")
		registerSymbol(&_copyProcessNameSymbol, frameworkHandle, "CopyProcessName")
		registerSymbol(&_coreAppearanceGetFontCGStyleRefSymbol, frameworkHandle, "CoreAppearanceGetFontCGStyleRef")
		registerSymbol(&_coreAppearanceGetFontShadowOutsetsSymbol, frameworkHandle, "CoreAppearanceGetFontShadowOutsets")
		registerSymbol(&_coreAppearanceGetFontSizeSymbol, frameworkHandle, "CoreAppearanceGetFontSize")
		registerSymbol(&_coreAppearanceGetQDFontForScriptSymbol, frameworkHandle, "CoreAppearanceGetQDFontForScript")
		registerSymbol(&_coreCursorCopyImagesSymbol, frameworkHandle, "CoreCursorCopyImages")
		registerSymbol(&_coreCursorGetDataSymbol, frameworkHandle, "CoreCursorGetData")
		registerSymbol(&_coreCursorGetDataSizeSymbol, frameworkHandle, "CoreCursorGetDataSize")
		registerSymbol(&_coreCursorSetSymbol, frameworkHandle, "CoreCursorSet")
		registerSymbol(&_coreCursorSetAndReturnSeedSymbol, frameworkHandle, "CoreCursorSetAndReturnSeed")
		registerSymbol(&_coreCursorUnregisterAllSymbol, frameworkHandle, "CoreCursorUnregisterAll")
		registerSymbol(&_coreDockAddFileToDockSymbol, frameworkHandle, "CoreDockAddFileToDock")
		registerSymbol(&_coreDockBounceAppTileSymbol, frameworkHandle, "CoreDockBounceAppTile")
		registerSymbol(&_coreDockCompositeProcessImageSymbol, frameworkHandle, "CoreDockCompositeProcessImage")
		registerSymbol(&_coreDockCopyPreferencesSymbol, frameworkHandle, "CoreDockCopyPreferences")
		registerSymbol(&_coreDockCopyWorkspacesAppBindingsSymbol, frameworkHandle, "CoreDockCopyWorkspacesAppBindings")
		registerSymbol(&_coreDockCreateDragTrashContextSymbol, frameworkHandle, "CoreDockCreateDragTrashContext")
		registerSymbol(&_coreDockDisableExposeKeysIfNecessarySymbol, frameworkHandle, "CoreDockDisableExposeKeysIfNecessary")
		registerSymbol(&_coreDockGetAutoHideEnabledSymbol, frameworkHandle, "CoreDockGetAutoHideEnabled")
		registerSymbol(&_coreDockGetContainerRectSymbol, frameworkHandle, "CoreDockGetContainerRect")
		registerSymbol(&_coreDockGetDashboardInDockSymbol, frameworkHandle, "CoreDockGetDashboardInDock")
		registerSymbol(&_coreDockGetEffectSymbol, frameworkHandle, "CoreDockGetEffect")
		registerSymbol(&_coreDockGetExposeCornerActionsSymbol, frameworkHandle, "CoreDockGetExposeCornerActions")
		registerSymbol(&_coreDockGetExposeCornerActionsWithModifiersSymbol, frameworkHandle, "CoreDockGetExposeCornerActionsWithModifiers")
		registerSymbol(&_coreDockGetItemDockContextSymbol, frameworkHandle, "CoreDockGetItemDockContext")
		registerSymbol(&_coreDockGetItemDockWindowSymbol, frameworkHandle, "CoreDockGetItemDockWindow")
		registerSymbol(&_coreDockGetMagnificationSizeSymbol, frameworkHandle, "CoreDockGetMagnificationSize")
		registerSymbol(&_coreDockGetMinimizeInPlaceSymbol, frameworkHandle, "CoreDockGetMinimizeInPlace")
		registerSymbol(&_coreDockGetOrientationAndPinningSymbol, frameworkHandle, "CoreDockGetOrientationAndPinning")
		registerSymbol(&_coreDockGetProcessContextSymbol, frameworkHandle, "CoreDockGetProcessContext")
		registerSymbol(&_coreDockGetProcessWindowSymbol, frameworkHandle, "CoreDockGetProcessWindow")
		registerSymbol(&_coreDockGetRectSymbol, frameworkHandle, "CoreDockGetRect")
		registerSymbol(&_coreDockGetRectAndOrientationSymbol, frameworkHandle, "CoreDockGetRectAndOrientation")
		registerSymbol(&_coreDockGetRectAndReasonSymbol, frameworkHandle, "CoreDockGetRectAndReason")
		registerSymbol(&_coreDockGetSpringLoadingTimeSymbol, frameworkHandle, "CoreDockGetSpringLoadingTime")
		registerSymbol(&_coreDockGetTileSizeSymbol, frameworkHandle, "CoreDockGetTileSize")
		registerSymbol(&_coreDockGetTrashWindowSymbol, frameworkHandle, "CoreDockGetTrashWindow")
		registerSymbol(&_coreDockGetWorkspacesCountSymbol, frameworkHandle, "CoreDockGetWorkspacesCount")
		registerSymbol(&_coreDockGetWorkspacesEnabledSymbol, frameworkHandle, "CoreDockGetWorkspacesEnabled")
		registerSymbol(&_coreDockGetWorkspacesKeyBindingsSymbol, frameworkHandle, "CoreDockGetWorkspacesKeyBindings")
		registerSymbol(&_coreDockIsDockRunningSymbol, frameworkHandle, "CoreDockIsDockRunning")
		registerSymbol(&_coreDockIsLaunchAnimationsEnabledSymbol, frameworkHandle, "CoreDockIsLaunchAnimationsEnabled")
		registerSymbol(&_coreDockIsMagnificationEnabledSymbol, frameworkHandle, "CoreDockIsMagnificationEnabled")
		registerSymbol(&_coreDockIsSpringLoadingEnabledSymbol, frameworkHandle, "CoreDockIsSpringLoadingEnabled")
		registerSymbol(&_coreDockMinimizeItemWithTitleSymbol, frameworkHandle, "CoreDockMinimizeItemWithTitle")
		registerSymbol(&_coreDockMinimizeItemWithTitleAsyncSymbol, frameworkHandle, "CoreDockMinimizeItemWithTitleAsync")
		registerSymbol(&_coreDockMinimizeItemsWithTitleSymbol, frameworkHandle, "CoreDockMinimizeItemsWithTitle")
		registerSymbol(&_coreDockMinimizeItemsWithTitleAsyncSymbol, frameworkHandle, "CoreDockMinimizeItemsWithTitleAsync")
		registerSymbol(&_coreDockPreventCommunicationWithDockSymbol, frameworkHandle, "CoreDockPreventCommunicationWithDock")
		registerSymbol(&_coreDockRegisterClientWithRunLoopSymbol, frameworkHandle, "CoreDockRegisterClientWithRunLoop")
		registerSymbol(&_coreDockReleaseDragTrashContextSymbol, frameworkHandle, "CoreDockReleaseDragTrashContext")
		registerSymbol(&_coreDockReleaseItemDockContextSymbol, frameworkHandle, "CoreDockReleaseItemDockContext")
		registerSymbol(&_coreDockReleaseItemDockWindowSymbol, frameworkHandle, "CoreDockReleaseItemDockWindow")
		registerSymbol(&_coreDockReleaseProcessContextSymbol, frameworkHandle, "CoreDockReleaseProcessContext")
		registerSymbol(&_coreDockReleaseProcessWindowSymbol, frameworkHandle, "CoreDockReleaseProcessWindow")
		registerSymbol(&_coreDockRemoveItemSymbol, frameworkHandle, "CoreDockRemoveItem")
		registerSymbol(&_coreDockRenderWindowIntoContextSymbol, frameworkHandle, "CoreDockRenderWindowIntoContext")
		registerSymbol(&_coreDockRestoreItemSymbol, frameworkHandle, "CoreDockRestoreItem")
		registerSymbol(&_coreDockRestoreItemAsyncSymbol, frameworkHandle, "CoreDockRestoreItemAsync")
		registerSymbol(&_coreDockRestoreItemWithOrderSymbol, frameworkHandle, "CoreDockRestoreItemWithOrder")
		registerSymbol(&_coreDockRestoreItemWithOrderAsyncSymbol, frameworkHandle, "CoreDockRestoreItemWithOrderAsync")
		registerSymbol(&_coreDockRestoreProcessImageSymbol, frameworkHandle, "CoreDockRestoreProcessImage")
		registerSymbol(&_coreDockRevealWindowForShowDesktopSymbol, frameworkHandle, "CoreDockRevealWindowForShowDesktop")
		registerSymbol(&_coreDockSendDragWindowMessageSymbol, frameworkHandle, "CoreDockSendDragWindowMessage")
		registerSymbol(&_coreDockSendNotificationSymbol, frameworkHandle, "CoreDockSendNotification")
		registerSymbol(&_coreDockSetAutoHideEnabledSymbol, frameworkHandle, "CoreDockSetAutoHideEnabled")
		registerSymbol(&_coreDockSetDashboardInDockSymbol, frameworkHandle, "CoreDockSetDashboardInDock")
		registerSymbol(&_coreDockSetDragStatusSymbol, frameworkHandle, "CoreDockSetDragStatus")
		registerSymbol(&_coreDockSetEffectSymbol, frameworkHandle, "CoreDockSetEffect")
		registerSymbol(&_coreDockSetExposeCornerActionSymbol, frameworkHandle, "CoreDockSetExposeCornerAction")
		registerSymbol(&_coreDockSetExposeCornerActionWithModifierSymbol, frameworkHandle, "CoreDockSetExposeCornerActionWithModifier")
		registerSymbol(&_coreDockSetItemTitleSymbol, frameworkHandle, "CoreDockSetItemTitle")
		registerSymbol(&_coreDockSetLaunchAnimationsEnabledSymbol, frameworkHandle, "CoreDockSetLaunchAnimationsEnabled")
		registerSymbol(&_coreDockSetMagnificationEnabledSymbol, frameworkHandle, "CoreDockSetMagnificationEnabled")
		registerSymbol(&_coreDockSetMagnificationSizeSymbol, frameworkHandle, "CoreDockSetMagnificationSize")
		registerSymbol(&_coreDockSetMiniViewSymbol, frameworkHandle, "CoreDockSetMiniView")
		registerSymbol(&_coreDockSetMinimizeInPlaceSymbol, frameworkHandle, "CoreDockSetMinimizeInPlace")
		registerSymbol(&_coreDockSetOrientationAndPinningSymbol, frameworkHandle, "CoreDockSetOrientationAndPinning")
		registerSymbol(&_coreDockSetPreferencesSymbol, frameworkHandle, "CoreDockSetPreferences")
		registerSymbol(&_coreDockSetProcessImageSymbol, frameworkHandle, "CoreDockSetProcessImage")
		registerSymbol(&_coreDockSetProcessLabelSymbol, frameworkHandle, "CoreDockSetProcessLabel")
		registerSymbol(&_coreDockSetProcessOpenRecentsSymbol, frameworkHandle, "CoreDockSetProcessOpenRecents")
		registerSymbol(&_coreDockSetShowDesktopCallbackSymbol, frameworkHandle, "CoreDockSetShowDesktopCallback")
		registerSymbol(&_coreDockSetSpringLoadingEnabledSymbol, frameworkHandle, "CoreDockSetSpringLoadingEnabled")
		registerSymbol(&_coreDockSetSpringLoadingTimeSymbol, frameworkHandle, "CoreDockSetSpringLoadingTime")
		registerSymbol(&_coreDockSetSpringWindowCallbacksSymbol, frameworkHandle, "CoreDockSetSpringWindowCallbacks")
		registerSymbol(&_coreDockSetTileSizeSymbol, frameworkHandle, "CoreDockSetTileSize")
		registerSymbol(&_coreDockSetTrashFullSymbol, frameworkHandle, "CoreDockSetTrashFull")
		registerSymbol(&_coreDockSetWindowLabelSymbol, frameworkHandle, "CoreDockSetWindowLabel")
		registerSymbol(&_coreDockSetWorkspacesAppBindingsSymbol, frameworkHandle, "CoreDockSetWorkspacesAppBindings")
		registerSymbol(&_coreDockSetWorkspacesCountSymbol, frameworkHandle, "CoreDockSetWorkspacesCount")
		registerSymbol(&_coreDockSetWorkspacesEnabledSymbol, frameworkHandle, "CoreDockSetWorkspacesEnabled")
		registerSymbol(&_coreDockSetWorkspacesKeyBindingsSymbol, frameworkHandle, "CoreDockSetWorkspacesKeyBindings")
		registerSymbol(&_coreDockUpdateWindowSymbol, frameworkHandle, "CoreDockUpdateWindow")
		registerSymbol(&_coreDragAcceptDragSymbol, frameworkHandle, "CoreDragAcceptDrag")
		registerSymbol(&_coreDragCancelDragSymbol, frameworkHandle, "CoreDragCancelDrag")
		registerSymbol(&_coreDragChangeBehaviorsSymbol, frameworkHandle, "CoreDragChangeBehaviors")
		registerSymbol(&_coreDragCleanDragStateSymbol, frameworkHandle, "CoreDragCleanDragState")
		registerSymbol(&_coreDragClearAllImageOverridesSymbol, frameworkHandle, "CoreDragClearAllImageOverrides")
		registerSymbol(&_coreDragCopyTrashLabelSymbol, frameworkHandle, "CoreDragCopyTrashLabel")
		registerSymbol(&_coreDragCreateSymbol, frameworkHandle, "CoreDragCreate")
		registerSymbol(&_coreDragCreateWithCFPasteboardSymbol, frameworkHandle, "CoreDragCreateWithCFPasteboard")
		registerSymbol(&_coreDragCreateWithPasteboardSymbol, frameworkHandle, "CoreDragCreateWithPasteboard")
		registerSymbol(&_coreDragCreateWithPasteboardRefSymbol, frameworkHandle, "CoreDragCreateWithPasteboardRef")
		registerSymbol(&_coreDragDisposeSymbol, frameworkHandle, "CoreDragDispose")
		registerSymbol(&_coreDragEnableSpringLoadingSymbol, frameworkHandle, "CoreDragEnableSpringLoading")
		registerSymbol(&_coreDragGetAllowableActionsSymbol, frameworkHandle, "CoreDragGetAllowableActions")
		registerSymbol(&_coreDragGetAttributesSymbol, frameworkHandle, "CoreDragGetAttributes")
		registerSymbol(&_coreDragGetCurrentDragSymbol, frameworkHandle, "CoreDragGetCurrentDrag")
		registerSymbol(&_coreDragGetDragWindowSymbol, frameworkHandle, "CoreDragGetDragWindow")
		registerSymbol(&_coreDragGetDropActionsSymbol, frameworkHandle, "CoreDragGetDropActions")
		registerSymbol(&_coreDragGetDropLocationSymbol, frameworkHandle, "CoreDragGetDropLocation")
		registerSymbol(&_coreDragGetForceSymbol, frameworkHandle, "CoreDragGetForce")
		registerSymbol(&_coreDragGetForceStageSymbol, frameworkHandle, "CoreDragGetForceStage")
		registerSymbol(&_coreDragGetItemBoundsSymbol, frameworkHandle, "CoreDragGetItemBounds")
		registerSymbol(&_coreDragGetModifiersSymbol, frameworkHandle, "CoreDragGetModifiers")
		registerSymbol(&_coreDragGetMouseLocationSymbol, frameworkHandle, "CoreDragGetMouseLocation")
		registerSymbol(&_coreDragGetOriginSymbol, frameworkHandle, "CoreDragGetOrigin")
		registerSymbol(&_coreDragGetPasteboardSymbol, frameworkHandle, "CoreDragGetPasteboard")
		registerSymbol(&_coreDragGetPasteboardRefSymbol, frameworkHandle, "CoreDragGetPasteboardRef")
		registerSymbol(&_coreDragGetSpringLoadingAttributesSymbol, frameworkHandle, "CoreDragGetSpringLoadingAttributes")
		registerSymbol(&_coreDragGetStandardDropLocationSymbol, frameworkHandle, "CoreDragGetStandardDropLocation")
		registerSymbol(&_coreDragGetValueForKeySymbol, frameworkHandle, "CoreDragGetValueForKey")
		registerSymbol(&_coreDragHasImageOverridesSymbol, frameworkHandle, "CoreDragHasImageOverrides")
		registerSymbol(&_coreDragInstallContextReceiveMessageHandlerOnConnectionSymbol, frameworkHandle, "CoreDragInstallContextReceiveMessageHandlerOnConnection")
		registerSymbol(&_coreDragInstallContextTrackingHandlerOnConnectionSymbol, frameworkHandle, "CoreDragInstallContextTrackingHandlerOnConnection")
		registerSymbol(&_coreDragInstallReceiveHandlerSymbol, frameworkHandle, "CoreDragInstallReceiveHandler")
		registerSymbol(&_coreDragInstallReceiveHandlerOnConnectionSymbol, frameworkHandle, "CoreDragInstallReceiveHandlerOnConnection")
		registerSymbol(&_coreDragInstallReceiveMessageHandlerSymbol, frameworkHandle, "CoreDragInstallReceiveMessageHandler")
		registerSymbol(&_coreDragInstallReceiveMessageHandlerOnConnectionSymbol, frameworkHandle, "CoreDragInstallReceiveMessageHandlerOnConnection")
		registerSymbol(&_coreDragInstallTrackingHandlerSymbol, frameworkHandle, "CoreDragInstallTrackingHandler")
		registerSymbol(&_coreDragInstallTrackingHandlerOnConnectionSymbol, frameworkHandle, "CoreDragInstallTrackingHandlerOnConnection")
		registerSymbol(&_coreDragIsTrashLabelSetSymbol, frameworkHandle, "CoreDragIsTrashLabelSet")
		registerSymbol(&_coreDragItemGetImageComponentsSymbol, frameworkHandle, "CoreDragItemGetImageComponents")
		registerSymbol(&_coreDragItemGetScreenFrameSymbol, frameworkHandle, "CoreDragItemGetScreenFrame")
		registerSymbol(&_coreDragItemHasAlternateSymbol, frameworkHandle, "CoreDragItemHasAlternate")
		registerSymbol(&_coreDragItemSetImageComponentsSymbol, frameworkHandle, "CoreDragItemSetImageComponents")
		registerSymbol(&_coreDragMessageNameSymbol, frameworkHandle, "CoreDragMessageName")
		registerSymbol(&_coreDragRefIsValidSymbol, frameworkHandle, "CoreDragRefIsValid")
		registerSymbol(&_coreDragRefSetImageDataForItemSymbol, frameworkHandle, "CoreDragRefSetImageDataForItem")
		registerSymbol(&_coreDragRegisterClientSymbol, frameworkHandle, "CoreDragRegisterClient")
		registerSymbol(&_coreDragRegisterClientInModesSymbol, frameworkHandle, "CoreDragRegisterClientInModes")
		registerSymbol(&_coreDragRegisterClientWithOptionsSymbol, frameworkHandle, "CoreDragRegisterClientWithOptions")
		registerSymbol(&_coreDragReleaseImageDataSymbol, frameworkHandle, "CoreDragReleaseImageData")
		registerSymbol(&_coreDragRemoveReceiveHandlerSymbol, frameworkHandle, "CoreDragRemoveReceiveHandler")
		registerSymbol(&_coreDragRemoveReceiveHandlerOnConnectionSymbol, frameworkHandle, "CoreDragRemoveReceiveHandlerOnConnection")
		registerSymbol(&_coreDragRemoveReceiveMessageHandlerSymbol, frameworkHandle, "CoreDragRemoveReceiveMessageHandler")
		registerSymbol(&_coreDragRemoveReceiveMessageHandlerOnConnectionSymbol, frameworkHandle, "CoreDragRemoveReceiveMessageHandlerOnConnection")
		registerSymbol(&_coreDragRemoveTrackingHandlerSymbol, frameworkHandle, "CoreDragRemoveTrackingHandler")
		registerSymbol(&_coreDragRemoveTrackingHandlerOnConnectionSymbol, frameworkHandle, "CoreDragRemoveTrackingHandlerOnConnection")
		registerSymbol(&_coreDragRequestDragCompleteMessageSymbol, frameworkHandle, "CoreDragRequestDragCompleteMessage")
		registerSymbol(&_coreDragSetAllowableActionsSymbol, frameworkHandle, "CoreDragSetAllowableActions")
		registerSymbol(&_coreDragSetAttributeSymbol, frameworkHandle, "CoreDragSetAttribute")
		registerSymbol(&_coreDragSetCGEventInputProcSymbol, frameworkHandle, "CoreDragSetCGEventInputProc")
		registerSymbol(&_coreDragSetCGEventProcsSymbol, frameworkHandle, "CoreDragSetCGEventProcs")
		registerSymbol(&_coreDragSetCGImageSymbol, frameworkHandle, "CoreDragSetCGImage")
		registerSymbol(&_coreDragSetCGImageWithScaleSymbol, frameworkHandle, "CoreDragSetCGImageWithScale")
		registerSymbol(&_coreDragSetDestClippingRectSymbol, frameworkHandle, "CoreDragSetDestClippingRect")
		registerSymbol(&_coreDragSetDragRegionSymbol, frameworkHandle, "CoreDragSetDragRegion")
		registerSymbol(&_coreDragSetDragRegionWithScaleSymbol, frameworkHandle, "CoreDragSetDragRegionWithScale")
		registerSymbol(&_coreDragSetDragWindowSymbol, frameworkHandle, "CoreDragSetDragWindow")
		registerSymbol(&_coreDragSetDropActionsSymbol, frameworkHandle, "CoreDragSetDropActions")
		registerSymbol(&_coreDragSetDropLocationSymbol, frameworkHandle, "CoreDragSetDropLocation")
		registerSymbol(&_coreDragSetDropProcSymbol, frameworkHandle, "CoreDragSetDropProc")
		registerSymbol(&_coreDragSetEventProcSymbol, frameworkHandle, "CoreDragSetEventProc")
		registerSymbol(&_coreDragSetExtendedEventProcSymbol, frameworkHandle, "CoreDragSetExtendedEventProc")
		registerSymbol(&_coreDragSetImageSymbol, frameworkHandle, "CoreDragSetImage")
		registerSymbol(&_coreDragSetImageOptionsSymbol, frameworkHandle, "CoreDragSetImageOptions")
		registerSymbol(&_coreDragSetInputProcSymbol, frameworkHandle, "CoreDragSetInputProc")
		registerSymbol(&_coreDragSetItemBoundsSymbol, frameworkHandle, "CoreDragSetItemBounds")
		registerSymbol(&_coreDragSetMouseLocationSymbol, frameworkHandle, "CoreDragSetMouseLocation")
		registerSymbol(&_coreDragSetOriginSymbol, frameworkHandle, "CoreDragSetOrigin")
		registerSymbol(&_coreDragSetRootCALayerSymbol, frameworkHandle, "CoreDragSetRootCALayer")
		registerSymbol(&_coreDragSetSourceClippingRectSymbol, frameworkHandle, "CoreDragSetSourceClippingRect")
		registerSymbol(&_coreDragSetStandardDropLocationSymbol, frameworkHandle, "CoreDragSetStandardDropLocation")
		registerSymbol(&_coreDragSetTrashDrawProcSymbol, frameworkHandle, "CoreDragSetTrashDrawProc")
		registerSymbol(&_coreDragSetTrashLabelSymbol, frameworkHandle, "CoreDragSetTrashLabel")
		registerSymbol(&_coreDragSetValueForKeySymbol, frameworkHandle, "CoreDragSetValueForKey")
		registerSymbol(&_coreDragStartDraggingSymbol, frameworkHandle, "CoreDragStartDragging")
		registerSymbol(&_coreDragStartDraggingAsyncSymbol, frameworkHandle, "CoreDragStartDraggingAsync")
		registerSymbol(&_coreDragUpdatesBeginSymbol, frameworkHandle, "CoreDragUpdatesBegin")
		registerSymbol(&_coreDragUpdatesCommitSymbol, frameworkHandle, "CoreDragUpdatesCommit")
		registerSymbol(&_coreGetDragInfoSymbol, frameworkHandle, "CoreGetDragInfo")
		registerSymbol(&_coreImagingCreateImageSymbol, frameworkHandle, "CoreImagingCreateImage")
		registerSymbol(&_coreMenuCreateKeyEquivalentStringSymbol, frameworkHandle, "CoreMenuCreateKeyEquivalentString")
		registerSymbol(&_coreMenuCreateVirtualKeyStringSymbol, frameworkHandle, "CoreMenuCreateVirtualKeyString")
		registerSymbol(&_coreMenuExtraAddMenuExtraSymbol, frameworkHandle, "CoreMenuExtraAddMenuExtra")
		registerSymbol(&_coreMenuExtraGetMenuExtraSymbol, frameworkHandle, "CoreMenuExtraGetMenuExtra")
		registerSymbol(&_coreMenuExtraGetOrderSymbol, frameworkHandle, "CoreMenuExtraGetOrder")
		registerSymbol(&_coreMenuExtraInsertMenuExtraSymbol, frameworkHandle, "CoreMenuExtraInsertMenuExtra")
		registerSymbol(&_coreMenuExtraRemoveMenuExtraSymbol, frameworkHandle, "CoreMenuExtraRemoveMenuExtra")
		registerSymbol(&_coreMenuGetVirtualKeyMapSymbol, frameworkHandle, "CoreMenuGetVirtualKeyMap")
		registerSymbol(&_createPasteboardFlavorTypeNameSymbol, frameworkHandle, "CreatePasteboardFlavorTypeName")
		registerSymbol(&_desktopPictureCopyDisplaySymbol, frameworkHandle, "DesktopPictureCopyDisplay")
		registerSymbol(&_desktopPictureCopyDisplayForSpaceSymbol, frameworkHandle, "DesktopPictureCopyDisplayForSpace")
		registerSymbol(&_desktopPictureCopyValueSymbol, frameworkHandle, "DesktopPictureCopyValue")
		registerSymbol(&_desktopPictureSetDisplaySymbol, frameworkHandle, "DesktopPictureSetDisplay")
		registerSymbol(&_desktopPictureSetDisplayForSpaceSymbol, frameworkHandle, "DesktopPictureSetDisplayForSpace")
		registerSymbol(&_desktopPictureSetValueSymbol, frameworkHandle, "DesktopPictureSetValue")
		registerSymbol(&_disposeIconActionUPPSymbol, frameworkHandle, "DisposeIconActionUPP")
		registerSymbol(&_disposeIconGetterUPPSymbol, frameworkHandle, "DisposeIconGetterUPP")
		registerSymbol(&_exitToShellSymbol, frameworkHandle, "ExitToShell")
		registerSymbol(&_getApplicationIsDaemonSymbol, frameworkHandle, "GetApplicationIsDaemon")
		registerSymbol(&_getCachedFlavorNameSymbol, frameworkHandle, "GetCachedFlavorName")
		registerSymbol(&_getCurrentProcessSymbol, frameworkHandle, "GetCurrentProcess")
		registerSymbol(&_getFrontProcessSymbol, frameworkHandle, "GetFrontProcess")
		registerSymbol(&_getGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol, frameworkHandle, "GetGlobalIconImagesCacheMaxEntriesAndMaxDataSize")
		registerSymbol(&_getIconFamilyDataSymbol, frameworkHandle, "GetIconFamilyData")
		registerSymbol(&_getIconRefVariantSymbol, frameworkHandle, "GetIconRefVariant")
		registerSymbol(&_getNextProcessSymbol, frameworkHandle, "GetNextProcess")
		registerSymbol(&_getProcessBundleLocationSymbol, frameworkHandle, "GetProcessBundleLocation")
		registerSymbol(&_getProcessForPIDSymbol, frameworkHandle, "GetProcessForPID")
		registerSymbol(&_getProcessInformationSymbol, frameworkHandle, "GetProcessInformation")
		registerSymbol(&_getProcessPIDSymbol, frameworkHandle, "GetProcessPID")
		registerSymbol(&_hIS_XPC_CFNotificationCenterPostNotificationSymbol, frameworkHandle, "HIS_XPC_CFNotificationCenterPostNotification")
		registerSymbol(&_hIS_XPC_CFPreferencesCopyValueSymbol, frameworkHandle, "HIS_XPC_CFPreferencesCopyValue")
		registerSymbol(&_hIS_XPC_CFPreferencesSetValueSymbol, frameworkHandle, "HIS_XPC_CFPreferencesSetValue")
		registerSymbol(&_hIS_XPC_CFPreferencesSynchronizeSymbol, frameworkHandle, "HIS_XPC_CFPreferencesSynchronize")
		registerSymbol(&_hIS_XPC_CopyCapsLockKeyLabelSymbol, frameworkHandle, "HIS_XPC_CopyCapsLockKeyLabel")
		registerSymbol(&_hIS_XPC_CopyMacManagerPrefsSymbol, frameworkHandle, "HIS_XPC_CopyMacManagerPrefs")
		registerSymbol(&_hIS_XPC_CopyPrimaryKeyboardLanguageSymbol, frameworkHandle, "HIS_XPC_CopyPrimaryKeyboardLanguage")
		registerSymbol(&_hIS_XPC_GetApplicationPolicyForURLsSymbol, frameworkHandle, "HIS_XPC_GetApplicationPolicyForURLs")
		registerSymbol(&_hIS_XPC_GetCapsLockLanguageSwitchSymbol, frameworkHandle, "HIS_XPC_GetCapsLockLanguageSwitch")
		registerSymbol(&_hIS_XPC_GetCapsLockModifierStateSymbol, frameworkHandle, "HIS_XPC_GetCapsLockModifierState")
		registerSymbol(&_hIS_XPC_GetGlobeKeyAvailabilitySymbol, frameworkHandle, "HIS_XPC_GetGlobeKeyAvailability")
		registerSymbol(&_hIS_XPC_GetMicKeyAvailabilitySymbol, frameworkHandle, "HIS_XPC_GetMicKeyAvailability")
		registerSymbol(&_hIS_XPC_ResetMessageConnectionSymbol, frameworkHandle, "HIS_XPC_ResetMessageConnection")
		registerSymbol(&_hIS_XPC_RevealFileInFinderSymbol, frameworkHandle, "HIS_XPC_RevealFileInFinder")
		registerSymbol(&_hIS_XPC_SendAppleEventToSystemProcessSymbol, frameworkHandle, "HIS_XPC_SendAppleEventToSystemProcess")
		registerSymbol(&_hIS_XPC_SetCapsLockDelayOverrideSymbol, frameworkHandle, "HIS_XPC_SetCapsLockDelayOverride")
		registerSymbol(&_hIS_XPC_SetCapsLockLEDSymbol, frameworkHandle, "HIS_XPC_SetCapsLockLED")
		registerSymbol(&_hIS_XPC_SetCapsLockLEDInhibitSymbol, frameworkHandle, "HIS_XPC_SetCapsLockLEDInhibit")
		registerSymbol(&_hIS_XPC_SetCapsLockModifierStateSymbol, frameworkHandle, "HIS_XPC_SetCapsLockModifierState")
		registerSymbol(&_hIS_XPC_SetNetworkLocationSymbol, frameworkHandle, "HIS_XPC_SetNetworkLocation")
		registerSymbol(&_hIShapeContainsPointSymbol, frameworkHandle, "HIShapeContainsPoint")
		registerSymbol(&_hIShapeCreateCopySymbol, frameworkHandle, "HIShapeCreateCopy")
		registerSymbol(&_hIShapeCreateCopyAsQDRgnSymbol, frameworkHandle, "HIShapeCreateCopyAsQDRgn")
		registerSymbol(&_hIShapeCreateDifferenceSymbol, frameworkHandle, "HIShapeCreateDifference")
		registerSymbol(&_hIShapeCreateEmptySymbol, frameworkHandle, "HIShapeCreateEmpty")
		registerSymbol(&_hIShapeCreateIntersectionSymbol, frameworkHandle, "HIShapeCreateIntersection")
		registerSymbol(&_hIShapeCreateMutableSymbol, frameworkHandle, "HIShapeCreateMutable")
		registerSymbol(&_hIShapeCreateMutableCopySymbol, frameworkHandle, "HIShapeCreateMutableCopy")
		registerSymbol(&_hIShapeCreateMutableWithRectSymbol, frameworkHandle, "HIShapeCreateMutableWithRect")
		registerSymbol(&_hIShapeCreateUnionSymbol, frameworkHandle, "HIShapeCreateUnion")
		registerSymbol(&_hIShapeCreateWithQDRgnSymbol, frameworkHandle, "HIShapeCreateWithQDRgn")
		registerSymbol(&_hIShapeCreateWithRectSymbol, frameworkHandle, "HIShapeCreateWithRect")
		registerSymbol(&_hIShapeCreateXorSymbol, frameworkHandle, "HIShapeCreateXor")
		registerSymbol(&_hIShapeDifferenceSymbol, frameworkHandle, "HIShapeDifference")
		registerSymbol(&_hIShapeEnumerateSymbol, frameworkHandle, "HIShapeEnumerate")
		registerSymbol(&_hIShapeGetAsQDRgnSymbol, frameworkHandle, "HIShapeGetAsQDRgn")
		registerSymbol(&_hIShapeGetBoundsSymbol, frameworkHandle, "HIShapeGetBounds")
		registerSymbol(&_hIShapeGetTypeIDSymbol, frameworkHandle, "HIShapeGetTypeID")
		registerSymbol(&_hIShapeInsetSymbol, frameworkHandle, "HIShapeInset")
		registerSymbol(&_hIShapeIntersectSymbol, frameworkHandle, "HIShapeIntersect")
		registerSymbol(&_hIShapeIntersectsRectSymbol, frameworkHandle, "HIShapeIntersectsRect")
		registerSymbol(&_hIShapeIsEmptySymbol, frameworkHandle, "HIShapeIsEmpty")
		registerSymbol(&_hIShapeIsRectangularSymbol, frameworkHandle, "HIShapeIsRectangular")
		registerSymbol(&_hIShapeOffsetSymbol, frameworkHandle, "HIShapeOffset")
		registerSymbol(&_hIShapeReplacePathInCGContextSymbol, frameworkHandle, "HIShapeReplacePathInCGContext")
		registerSymbol(&_hIShapeSetEmptySymbol, frameworkHandle, "HIShapeSetEmpty")
		registerSymbol(&_hIShapeSetWithShapeSymbol, frameworkHandle, "HIShapeSetWithShape")
		registerSymbol(&_hIShapeUnionSymbol, frameworkHandle, "HIShapeUnion")
		registerSymbol(&_hIShapeUnionWithRectSymbol, frameworkHandle, "HIShapeUnionWithRect")
		registerSymbol(&_hIShapeXorSymbol, frameworkHandle, "HIShapeXor")
		registerSymbol(&_iCAddMapEntrySymbol, frameworkHandle, "ICAddMapEntry")
		registerSymbol(&_iCAddProfileSymbol, frameworkHandle, "ICAddProfile")
		registerSymbol(&_iCBeginSymbol, frameworkHandle, "ICBegin")
		registerSymbol(&_iCCountMapEntriesSymbol, frameworkHandle, "ICCountMapEntries")
		registerSymbol(&_iCCountPrefSymbol, frameworkHandle, "ICCountPref")
		registerSymbol(&_iCCountProfilesSymbol, frameworkHandle, "ICCountProfiles")
		registerSymbol(&_iCCreateGURLEventSymbol, frameworkHandle, "ICCreateGURLEvent")
		registerSymbol(&_iCDeleteMapEntrySymbol, frameworkHandle, "ICDeleteMapEntry")
		registerSymbol(&_iCDeletePrefSymbol, frameworkHandle, "ICDeletePref")
		registerSymbol(&_iCDeleteProfileSymbol, frameworkHandle, "ICDeleteProfile")
		registerSymbol(&_iCEditPreferencesSymbol, frameworkHandle, "ICEditPreferences")
		registerSymbol(&_iCEndSymbol, frameworkHandle, "ICEnd")
		registerSymbol(&_iCFindPrefHandleSymbol, frameworkHandle, "ICFindPrefHandle")
		registerSymbol(&_iCGetConfigNameSymbol, frameworkHandle, "ICGetConfigName")
		registerSymbol(&_iCGetCurrentProfileSymbol, frameworkHandle, "ICGetCurrentProfile")
		registerSymbol(&_iCGetDefaultPrefSymbol, frameworkHandle, "ICGetDefaultPref")
		registerSymbol(&_iCGetIndMapEntrySymbol, frameworkHandle, "ICGetIndMapEntry")
		registerSymbol(&_iCGetIndPrefSymbol, frameworkHandle, "ICGetIndPref")
		registerSymbol(&_iCGetIndProfileSymbol, frameworkHandle, "ICGetIndProfile")
		registerSymbol(&_iCGetMapEntrySymbol, frameworkHandle, "ICGetMapEntry")
		registerSymbol(&_iCGetPermSymbol, frameworkHandle, "ICGetPerm")
		registerSymbol(&_iCGetPrefSymbol, frameworkHandle, "ICGetPref")
		registerSymbol(&_iCGetPrefHandleSymbol, frameworkHandle, "ICGetPrefHandle")
		registerSymbol(&_iCGetProfileNameSymbol, frameworkHandle, "ICGetProfileName")
		registerSymbol(&_iCGetSeedSymbol, frameworkHandle, "ICGetSeed")
		registerSymbol(&_iCGetVersionSymbol, frameworkHandle, "ICGetVersion")
		registerSymbol(&_iCLaunchURLSymbol, frameworkHandle, "ICLaunchURL")
		registerSymbol(&_iCMapEntriesFilenameSymbol, frameworkHandle, "ICMapEntriesFilename")
		registerSymbol(&_iCMapEntriesTypeCreatorSymbol, frameworkHandle, "ICMapEntriesTypeCreator")
		registerSymbol(&_iCMapFilenameSymbol, frameworkHandle, "ICMapFilename")
		registerSymbol(&_iCMapTypeCreatorSymbol, frameworkHandle, "ICMapTypeCreator")
		registerSymbol(&_iCParseURLSymbol, frameworkHandle, "ICParseURL")
		registerSymbol(&_iCSendGURLEventSymbol, frameworkHandle, "ICSendGURLEvent")
		registerSymbol(&_iCSetCurrentProfileSymbol, frameworkHandle, "ICSetCurrentProfile")
		registerSymbol(&_iCSetMapEntrySymbol, frameworkHandle, "ICSetMapEntry")
		registerSymbol(&_iCSetPrefSymbol, frameworkHandle, "ICSetPref")
		registerSymbol(&_iCSetPrefHandleSymbol, frameworkHandle, "ICSetPrefHandle")
		registerSymbol(&_iCSetProfileNameSymbol, frameworkHandle, "ICSetProfileName")
		registerSymbol(&_iCStartSymbol, frameworkHandle, "ICStart")
		registerSymbol(&_iCStopSymbol, frameworkHandle, "ICStop")
		registerSymbol(&_iconRefContainsCGPointSymbol, frameworkHandle, "IconRefContainsCGPoint")
		registerSymbol(&_iconRefIntersectsCGRectSymbol, frameworkHandle, "IconRefIntersectsCGRect")
		registerSymbol(&_iconRefToHIShapeSymbol, frameworkHandle, "IconRefToHIShape")
		registerSymbol(&_iconRefToIconFamilySymbol, frameworkHandle, "IconRefToIconFamily")
		registerSymbol(&_invokeIconActionUPPSymbol, frameworkHandle, "InvokeIconActionUPP")
		registerSymbol(&_invokeIconGetterUPPSymbol, frameworkHandle, "InvokeIconGetterUPP")
		registerSymbol(&_isIconRefMaskEmptySymbol, frameworkHandle, "IsIconRefMaskEmpty")
		registerSymbol(&_isProcessManagerInitializedSymbol, frameworkHandle, "IsProcessManagerInitialized")
		registerSymbol(&_isProcessVisibleSymbol, frameworkHandle, "IsProcessVisible")
		registerSymbol(&_killProcessSymbol, frameworkHandle, "KillProcess")
		registerSymbol(&_launchApplicationSymbol, frameworkHandle, "LaunchApplication")
		registerSymbol(&_launchProcessSymbol, frameworkHandle, "LaunchProcess")
		registerSymbol(&_launchProcessAsyncSymbol, frameworkHandle, "LaunchProcessAsync")
		registerSymbol(&_mSHCreateMIGServerSourceSymbol, frameworkHandle, "MSHCreateMIGServerSource")
		registerSymbol(&_mSHCreateMachServerSourceSymbol, frameworkHandle, "MSHCreateMachServerSource")
		registerSymbol(&_mSHGetMachPortFromSourceSymbol, frameworkHandle, "MSHGetMachPortFromSource")
		registerSymbol(&_mSHMIGSourceSetNoSendersCallbackSymbol, frameworkHandle, "MSHMIGSourceSetNoSendersCallback")
		registerSymbol(&_mSHMIGSourceSetSendOnceCallbackSymbol, frameworkHandle, "MSHMIGSourceSetSendOnceCallback")
		registerSymbol(&_newIconActionUPPSymbol, frameworkHandle, "NewIconActionUPP")
		registerSymbol(&_newIconGetterUPPSymbol, frameworkHandle, "NewIconGetterUPP")
		registerSymbol(&_pasteboardClearSymbol, frameworkHandle, "PasteboardClear")
		registerSymbol(&_pasteboardCopyItemFlavorDataSymbol, frameworkHandle, "PasteboardCopyItemFlavorData")
		registerSymbol(&_pasteboardCopyItemFlavorsSymbol, frameworkHandle, "PasteboardCopyItemFlavors")
		registerSymbol(&_pasteboardCopyNameSymbol, frameworkHandle, "PasteboardCopyName")
		registerSymbol(&_pasteboardCopyPasteLocationSymbol, frameworkHandle, "PasteboardCopyPasteLocation")
		registerSymbol(&_pasteboardCreateSymbol, frameworkHandle, "PasteboardCreate")
		registerSymbol(&_pasteboardCreateWithCFPasteboardSymbol, frameworkHandle, "PasteboardCreateWithCFPasteboard")
		registerSymbol(&_pasteboardDontHonorPromisesSymbol, frameworkHandle, "PasteboardDontHonorPromises")
		registerSymbol(&_pasteboardGetCFPasteboardSymbol, frameworkHandle, "PasteboardGetCFPasteboard")
		registerSymbol(&_pasteboardGetItemCountSymbol, frameworkHandle, "PasteboardGetItemCount")
		registerSymbol(&_pasteboardGetItemFlavorFlagsSymbol, frameworkHandle, "PasteboardGetItemFlavorFlags")
		registerSymbol(&_pasteboardGetItemIdentifierSymbol, frameworkHandle, "PasteboardGetItemIdentifier")
		registerSymbol(&_pasteboardGetTypeIDSymbol, frameworkHandle, "PasteboardGetTypeID")
		registerSymbol(&_pasteboardPutItemFlavorSymbol, frameworkHandle, "PasteboardPutItemFlavor")
		registerSymbol(&_pasteboardResolvePromisesSymbol, frameworkHandle, "PasteboardResolvePromises")
		registerSymbol(&_pasteboardSetPasteLocationSymbol, frameworkHandle, "PasteboardSetPasteLocation")
		registerSymbol(&_pasteboardSetPromiseKeeperSymbol, frameworkHandle, "PasteboardSetPromiseKeeper")
		registerSymbol(&_pasteboardSynchronizeSymbol, frameworkHandle, "PasteboardSynchronize")
		registerSymbol(&_pasteboardToggleDuplicateFlavorCheckSymbol, frameworkHandle, "PasteboardToggleDuplicateFlavorCheck")
		registerSymbol(&_plotIconRefInContextSymbol, frameworkHandle, "PlotIconRefInContext")
		registerSymbol(&_processInformationCopyDictionarySymbol, frameworkHandle, "ProcessInformationCopyDictionary")
		registerSymbol(&_sXArbitrationAddQueuedOutputRequestSymbol, frameworkHandle, "SXArbitrationAddQueuedOutputRequest")
		registerSymbol(&_sXArbitrationCancelQueuedRequestSymbol, frameworkHandle, "SXArbitrationCancelQueuedRequest")
		registerSymbol(&_sXArbitrationCreateServerSourceSymbol, frameworkHandle, "SXArbitrationCreateServerSource")
		registerSymbol(&_sXArbitrationIsQueuedRequestPendingSymbol, frameworkHandle, "SXArbitrationIsQueuedRequestPending")
		registerSymbol(&_sXArbitrationRegisterOutputStartingSymbol, frameworkHandle, "SXArbitrationRegisterOutputStarting")
		registerSymbol(&_sXArbitrationRegisterOutputStoppedSymbol, frameworkHandle, "SXArbitrationRegisterOutputStopped")
		registerSymbol(&_sameProcessSymbol, frameworkHandle, "SameProcess")
		registerSymbol(&_serializeCFTypeSymbol, frameworkHandle, "SerializeCFType")
		registerSymbol(&_setApplicationIsDaemonSymbol, frameworkHandle, "SetApplicationIsDaemon")
		registerSymbol(&_setFrontProcessSymbol, frameworkHandle, "SetFrontProcess")
		registerSymbol(&_setFrontProcessWithOptionsSymbol, frameworkHandle, "SetFrontProcessWithOptions")
		registerSymbol(&_setGlobalIconImagesCacheMaxEntriesAndMaxDataSizeSymbol, frameworkHandle, "SetGlobalIconImagesCacheMaxEntriesAndMaxDataSize")
		registerSymbol(&_setIconFamilyDataSymbol, frameworkHandle, "SetIconFamilyData")
		registerSymbol(&_setLabelColorAndNameSymbol, frameworkHandle, "SetLabelColorAndName")
		registerSymbol(&_showHideDragSymbol, frameworkHandle, "ShowHideDrag")
		registerSymbol(&_showHideProcessSymbol, frameworkHandle, "ShowHideProcess")
		registerSymbol(&_startIPCFlockingPingSymbol, frameworkHandle, "StartIPCFlockingPing")
		registerSymbol(&_startIPCPingSymbol, frameworkHandle, "StartIPCPing")
		registerSymbol(&_transformProcessTypeSymbol, frameworkHandle, "TransformProcessType")
		registerSymbol(&_translationCopyDestinationTypeSymbol, frameworkHandle, "TranslationCopyDestinationType")
		registerSymbol(&_translationCopySourceTypeSymbol, frameworkHandle, "TranslationCopySourceType")
		registerSymbol(&_translationCreateSymbol, frameworkHandle, "TranslationCreate")
		registerSymbol(&_translationCreateWithSourceArraySymbol, frameworkHandle, "TranslationCreateWithSourceArray")
		registerSymbol(&_translationGetTranslationFlagsSymbol, frameworkHandle, "TranslationGetTranslationFlags")
		registerSymbol(&_translationGetTypeIDSymbol, frameworkHandle, "TranslationGetTypeID")
		registerSymbol(&_translationPerformForDataSymbol, frameworkHandle, "TranslationPerformForData")
		registerSymbol(&_translationPerformForFileSymbol, frameworkHandle, "TranslationPerformForFile")
		registerSymbol(&_translationPerformForURLSymbol, frameworkHandle, "TranslationPerformForURL")
		registerSymbol(&_uAZoomChangeFocusSymbol, frameworkHandle, "UAZoomChangeFocus")
		registerSymbol(&_uAZoomEnabledSymbol, frameworkHandle, "UAZoomEnabled")
		registerSymbol(&_unserializeCFTypeSymbol, frameworkHandle, "UnserializeCFType")
		registerSymbol(&_wakeUpProcessSymbol, frameworkHandle, "WakeUpProcess")
		registerSymbol(&__AXCopyActionDescriptionSymbol, frameworkHandle, "_AXCopyActionDescription")
		registerSymbol(&__AXCopyChildrenHashSymbol, frameworkHandle, "_AXCopyChildrenHash")
		registerSymbol(&__AXCopyChildrenHashWithRelativeFrameSymbol, frameworkHandle, "_AXCopyChildrenHashWithRelativeFrame")
		registerSymbol(&__AXCopyRoleDescriptionSymbol, frameworkHandle, "_AXCopyRoleDescription")
		registerSymbol(&__AXCopyRoleDescriptionWithSubroleSymbol, frameworkHandle, "_AXCopyRoleDescriptionWithSubrole")
		registerSymbol(&__AXCopyTitleSymbol, frameworkHandle, "_AXCopyTitle")
		registerSymbol(&__AXCreateElementOrderingSymbol, frameworkHandle, "_AXCreateElementOrdering")
		registerSymbol(&__AXCurrentRequestCanAccessRemoteDeviceContentSymbol, frameworkHandle, "_AXCurrentRequestCanAccessRemoteDeviceContent")
		registerSymbol(&__AXCurrentRequestCanReturnInspectionContentSymbol, frameworkHandle, "_AXCurrentRequestCanReturnInspectionContent")
		registerSymbol(&__AXCurrentRequestCanReturnProtectedContentSymbol, frameworkHandle, "_AXCurrentRequestCanReturnProtectedContent")
		registerSymbol(&__AXGetClientForCurrentRequestUntrustedSymbol, frameworkHandle, "_AXGetClientForCurrentRequestUntrusted")
		registerSymbol(&__AXHasClientsWithAccessRemoteDeviceContentSymbol, frameworkHandle, "_AXHasClientsWithAccessRemoteDeviceContent")
		registerSymbol(&__AXInterfaceCopyCursorColorFillSymbol, frameworkHandle, "_AXInterfaceCopyCursorColorFill")
		registerSymbol(&__AXInterfaceCopyCursorColorOutlineSymbol, frameworkHandle, "_AXInterfaceCopyCursorColorOutline")
		registerSymbol(&__AXInterfaceCursorIsOverriddenSymbol, frameworkHandle, "_AXInterfaceCursorIsOverridden")
		registerSymbol(&__AXInterfaceCursorSetAndReturnSeedSymbol, frameworkHandle, "_AXInterfaceCursorSetAndReturnSeed")
		registerSymbol(&__AXInterfaceGetBristolEnabledSymbol, frameworkHandle, "_AXInterfaceGetBristolEnabled")
		registerSymbol(&__AXInterfaceGetClassicInvertColorEnabledSymbol, frameworkHandle, "_AXInterfaceGetClassicInvertColorEnabled")
		registerSymbol(&__AXInterfaceGetDifferentiateWithoutColorEnabledSymbol, frameworkHandle, "_AXInterfaceGetDifferentiateWithoutColorEnabled")
		registerSymbol(&__AXInterfaceGetIncreaseContrastEnabledSymbol, frameworkHandle, "_AXInterfaceGetIncreaseContrastEnabled")
		registerSymbol(&__AXInterfaceGetReduceMotionEnabledSymbol, frameworkHandle, "_AXInterfaceGetReduceMotionEnabled")
		registerSymbol(&__AXInterfaceGetReduceTextInsertionPointModulationEnabledSymbol, frameworkHandle, "_AXInterfaceGetReduceTextInsertionPointModulationEnabled")
		registerSymbol(&__AXInterfaceGetReduceTransparencyEnabledSymbol, frameworkHandle, "_AXInterfaceGetReduceTransparencyEnabled")
		registerSymbol(&__AXInterfaceGetRichmondEnabledSymbol, frameworkHandle, "_AXInterfaceGetRichmondEnabled")
		registerSymbol(&__AXInterfaceGetShowToolbarButtonShapesEnabledSymbol, frameworkHandle, "_AXInterfaceGetShowToolbarButtonShapesEnabled")
		registerSymbol(&__AXInterfaceGetShowWindowTitlebarIconsEnabledSymbol, frameworkHandle, "_AXInterfaceGetShowWindowTitlebarIconsEnabled")
		registerSymbol(&__AXInterfaceSetClassicInvertColorEnabledSymbol, frameworkHandle, "_AXInterfaceSetClassicInvertColorEnabled")
		registerSymbol(&__AXInterfaceSetCursorColorFillSymbol, frameworkHandle, "_AXInterfaceSetCursorColorFill")
		registerSymbol(&__AXInterfaceSetCursorColorOutlineSymbol, frameworkHandle, "_AXInterfaceSetCursorColorOutline")
		registerSymbol(&__AXInterfaceSetCursorIsOverriddenSymbol, frameworkHandle, "_AXInterfaceSetCursorIsOverridden")
		registerSymbol(&__AXInterfaceSetDifferentiateWithoutColorEnabledSymbol, frameworkHandle, "_AXInterfaceSetDifferentiateWithoutColorEnabled")
		registerSymbol(&__AXInterfaceSetDifferentiateWithoutColorEnabledOverrideSymbol, frameworkHandle, "_AXInterfaceSetDifferentiateWithoutColorEnabledOverride")
		registerSymbol(&__AXInterfaceSetIncreaseContrastEnabledSymbol, frameworkHandle, "_AXInterfaceSetIncreaseContrastEnabled")
		registerSymbol(&__AXInterfaceSetIncreaseContrastEnabledOverrideSymbol, frameworkHandle, "_AXInterfaceSetIncreaseContrastEnabledOverride")
		registerSymbol(&__AXInterfaceSetReduceMotionEnabledSymbol, frameworkHandle, "_AXInterfaceSetReduceMotionEnabled")
		registerSymbol(&__AXInterfaceSetReduceMotionEnabledOverrideSymbol, frameworkHandle, "_AXInterfaceSetReduceMotionEnabledOverride")
		registerSymbol(&__AXInterfaceSetReduceTextInsertionPointModulationEnabledSymbol, frameworkHandle, "_AXInterfaceSetReduceTextInsertionPointModulationEnabled")
		registerSymbol(&__AXInterfaceSetReduceTextInsertionPointModulationEnabledOverrideSymbol, frameworkHandle, "_AXInterfaceSetReduceTextInsertionPointModulationEnabledOverride")
		registerSymbol(&__AXInterfaceSetReduceTransparencyEnabledSymbol, frameworkHandle, "_AXInterfaceSetReduceTransparencyEnabled")
		registerSymbol(&__AXInterfaceSetReduceTransparencyEnabledOverrideSymbol, frameworkHandle, "_AXInterfaceSetReduceTransparencyEnabledOverride")
		registerSymbol(&__AXInterfaceSetShowToolbarButtonShapesEnabledSymbol, frameworkHandle, "_AXInterfaceSetShowToolbarButtonShapesEnabled")
		registerSymbol(&__AXInterfaceSetShowToolbarButtonShapesEnabledOverrideSymbol, frameworkHandle, "_AXInterfaceSetShowToolbarButtonShapesEnabledOverride")
		registerSymbol(&__AXInterfaceSetShowWindowTitlebarIconsEnabledSymbol, frameworkHandle, "_AXInterfaceSetShowWindowTitlebarIconsEnabled")
		registerSymbol(&__AXInterfaceSetShowWindowTitlebarIconsEnabledOverrideSymbol, frameworkHandle, "_AXInterfaceSetShowWindowTitlebarIconsEnabledOverride")
		registerSymbol(&__AXIsAppleClientForCurrentRequestUntrustedSymbol, frameworkHandle, "_AXIsAppleClientForCurrentRequestUntrusted")
		registerSymbol(&__AXNotificationHandlerCreateWithCallbackSymbol, frameworkHandle, "_AXNotificationHandlerCreateWithCallback")
		registerSymbol(&__AXRegisterControlComputerAccessSymbol, frameworkHandle, "_AXRegisterControlComputerAccess")
		registerSymbol(&__AXSetAuditTokenIsAuthenticatedCallbackSymbol, frameworkHandle, "_AXSetAuditTokenIsAuthenticatedCallback")
		registerSymbol(&__AXSetClientIdentificationOverrideSymbol, frameworkHandle, "_AXSetClientIdentificationOverride")
		registerSymbol(&__AXShouldElementBeIgnoredForNavigationSymbol, frameworkHandle, "_AXShouldElementBeIgnoredForNavigation")
		registerSymbol(&__AXUIElementCopyElementAtPositionIncludeIgnoredSymbol, frameworkHandle, "_AXUIElementCopyElementAtPositionIncludeIgnored")
		registerSymbol(&__AXUIElementCreateApplicationWithPresenterPidSymbol, frameworkHandle, "_AXUIElementCreateApplicationWithPresenterPid")
		registerSymbol(&__AXUIElementCreateWithDataSymbol, frameworkHandle, "_AXUIElementCreateWithData")
		registerSymbol(&__AXUIElementCreateWithDataAndPidSymbol, frameworkHandle, "_AXUIElementCreateWithDataAndPid")
		registerSymbol(&__AXUIElementCreateWithDataAndPresenterPidSymbol, frameworkHandle, "_AXUIElementCreateWithDataAndPresenterPid")
		registerSymbol(&__AXUIElementCreateWithPtrSymbol, frameworkHandle, "_AXUIElementCreateWithPtr")
		registerSymbol(&__AXUIElementCreateWithRemoteTokenSymbol, frameworkHandle, "_AXUIElementCreateWithRemoteToken")
		registerSymbol(&__AXUIElementGetActualPidSymbol, frameworkHandle, "_AXUIElementGetActualPid")
		registerSymbol(&__AXUIElementGetDataSymbol, frameworkHandle, "_AXUIElementGetData")
		registerSymbol(&__AXUIElementGetIsProcessSuspendedSymbol, frameworkHandle, "_AXUIElementGetIsProcessSuspended")
		registerSymbol(&__AXUIElementGetWindowSymbol, frameworkHandle, "_AXUIElementGetWindow")
		registerSymbol(&__AXUIElementLoadAccessibilityBundlesSymbol, frameworkHandle, "_AXUIElementLoadAccessibilityBundles")
		registerSymbol(&__AXUIElementNotifyProcessSuspendStatusSymbol, frameworkHandle, "_AXUIElementNotifyProcessSuspendStatus")
		registerSymbol(&__AXUIElementPostNotificationSymbol, frameworkHandle, "_AXUIElementPostNotification")
		registerSymbol(&__AXUIElementPostNotificationForObservedElementSymbol, frameworkHandle, "_AXUIElementPostNotificationForObservedElement")
		registerSymbol(&__AXUIElementPostNotificationWithInfoSymbol, frameworkHandle, "_AXUIElementPostNotificationWithInfo")
		registerSymbol(&__AXUIElementRegisterServerWithRunLoopSymbol, frameworkHandle, "_AXUIElementRegisterServerWithRunLoop")
		registerSymbol(&__AXUIElementRemoteTokenCreateSymbol, frameworkHandle, "_AXUIElementRemoteTokenCreate")
		registerSymbol(&__AXUIElementRequestServicedBySecondaryAXThreadSymbol, frameworkHandle, "_AXUIElementRequestServicedBySecondaryAXThread")
		registerSymbol(&__AXUIElementUnregisterServerSymbol, frameworkHandle, "_AXUIElementUnregisterServer")
		registerSymbol(&__AXUIElementUpdateServerSymbol, frameworkHandle, "_AXUIElementUpdateServer")
		registerSymbol(&__AXUIElementUseSecondaryAXThreadSymbol, frameworkHandle, "_AXUIElementUseSecondaryAXThread")
		registerSymbol(&__AddLabelsChangedCallbackSymbol, frameworkHandle, "_AddLabelsChangedCallback")
		registerSymbol(&__CopyProcessBundleLocationURLSymbol, frameworkHandle, "_CopyProcessBundleLocationURL")
		registerSymbol(&__GDBIconsCGCacheListSymbol, frameworkHandle, "_GDBIconsCGCacheList")
		registerSymbol(&__GetApplicationDesiresAttentionSymbol, frameworkHandle, "_GetApplicationDesiresAttention")
		registerSymbol(&__GetFrontUIProcessSymbol, frameworkHandle, "_GetFrontUIProcess")
		registerSymbol(&__HIE_CaptureExceptionTelltaleOnceSymbol, frameworkHandle, "_HIE_CaptureExceptionTelltaleOnce")
		registerSymbol(&__HIE_CrashSymbol, frameworkHandle, "_HIE_Crash")
		registerSymbol(&__HIE_WillSwallowExceptionSymbol, frameworkHandle, "_HIE_WillSwallowException")
		registerSymbol(&__HIRLU_AddRunLoopModeForDeferredActionsSymbol, frameworkHandle, "_HIRLU_AddRunLoopModeForDeferredActions")
		registerSymbol(&__HIRLU_DeferActionsSymbol, frameworkHandle, "_HIRLU_DeferActions")
		registerSymbol(&__HIRunLoopSemaphoreCreateWithRunLoopModeSymbol, frameworkHandle, "_HIRunLoopSemaphoreCreateWithRunLoopMode")
		registerSymbol(&__HIRunLoopSemaphoreSetLegendSymbol, frameworkHandle, "_HIRunLoopSemaphoreSetLegend")
		registerSymbol(&__HIRunLoopSemaphoreSignalSymbol, frameworkHandle, "_HIRunLoopSemaphoreSignal")
		registerSymbol(&__HIRunLoopSemaphoreWaitSymbol, frameworkHandle, "_HIRunLoopSemaphoreWait")
		registerSymbol(&__HIShapeCreateWithCGImageSymbol, frameworkHandle, "_HIShapeCreateWithCGImage")
		registerSymbol(&__HIShapeCreateWithCGSRegionObjSymbol, frameworkHandle, "_HIShapeCreateWithCGSRegionObj")
		registerSymbol(&__HIShapeGetNativeSymbol, frameworkHandle, "_HIShapeGetNative")
		registerSymbol(&__HIShapeOutsetToPixelBoundarySymbol, frameworkHandle, "_HIShapeOutsetToPixelBoundary")
		registerSymbol(&__HIShapeSetImmutableSymbol, frameworkHandle, "_HIShapeSetImmutable")
		registerSymbol(&__HIShapeSetShapeWithOffsetSymbol, frameworkHandle, "_HIShapeSetShapeWithOffset")
		registerSymbol(&__HideOtherApplicationsSymbol, frameworkHandle, "_HideOtherApplications")
		registerSymbol(&__ICCopyMailHostNameSymbol, frameworkHandle, "_ICCopyMailHostName")
		registerSymbol(&__ICCopyMailUserNameSymbol, frameworkHandle, "_ICCopyMailUserName")
		registerSymbol(&__ISCreateCGImageForTypeSymbol, frameworkHandle, "_ISCreateCGImageForType")
		registerSymbol(&__ISCreateCGImageForTypeAtScaleSymbol, frameworkHandle, "_ISCreateCGImageForTypeAtScale")
		registerSymbol(&__IconServicesGetCGImageRefFromIconRefSymbol, frameworkHandle, "_IconServicesGetCGImageRefFromIconRef")
		registerSymbol(&__IconServicesGetCGImageRefFromURLSymbol, frameworkHandle, "_IconServicesGetCGImageRefFromURL")
		registerSymbol(&__InstallGURLEventHandlerSymbol, frameworkHandle, "_InstallGURLEventHandler")
		registerSymbol(&__PIPZoomingEnabledSymbol, frameworkHandle, "_PIPZoomingEnabled")
		registerSymbol(&__RegisterApplicationSymbol, frameworkHandle, "_RegisterApplication")
		registerSymbol(&__RegisterAsSessionLauncherApplicationSymbol, frameworkHandle, "_RegisterAsSessionLauncherApplication")
		registerSymbol(&__RemoveLabelsChangedCallbackSymbol, frameworkHandle, "_RemoveLabelsChangedCallback")
		registerSymbol(&__SetApplicationDesiresAttentionSymbol, frameworkHandle, "_SetApplicationDesiresAttention")
		registerSymbol(&__SetFrontProcessWithOptionsSymbol, frameworkHandle, "_SetFrontProcessWithOptions")
		registerSymbol(&__SetHLTBWakeUpHookSymbol, frameworkHandle, "_SetHLTBWakeUpHook")
		registerSymbol(&__ShowAllApplicationsSymbol, frameworkHandle, "_ShowAllApplications")
		registerSymbol(&__SignalApplicationReadySymbol, frameworkHandle, "_SignalApplicationReady")
		registerSymbol(&__UAZoomFocusChangeSymbol, frameworkHandle, "_UAZoomFocusChange")
		registerSymbol(&__UAZoomFocusChangeAnchoredSymbol, frameworkHandle, "_UAZoomFocusChangeAnchored")
		registerSymbol(&__UAZoomFocusChangeHighlightRectSymbol, frameworkHandle, "_UAZoomFocusChangeHighlightRect")
		registerSymbol(&__UAZoomFocusChangeHighlightRectAnchoredSymbol, frameworkHandle, "_UAZoomFocusChangeHighlightRectAnchored")
		registerSymbol(&__UAZoomingEnabledSymbol, frameworkHandle, "_UAZoomingEnabled")
		registerSymbol(&__UnregisterAsSessionLauncherApplicationSymbol, frameworkHandle, "_UnregisterAsSessionLauncherApplication")
		registerSymbol(&_gDockDragCallbackSymbol, frameworkHandle, "gDockDragCallback")
		registerSymbol(&_gDockDragWindowCallbackSymbol, frameworkHandle, "gDockDragWindowCallback")
		registerSymbol(&_kAXAttachmentTextAttributeSymbol, frameworkHandle, "kAXAttachmentTextAttribute")
		registerSymbol(&_kAXAutocorrectedTextAttributeSymbol, frameworkHandle, "kAXAutocorrectedTextAttribute")
		registerSymbol(&_kAXBackgroundColorTextAttributeSymbol, frameworkHandle, "kAXBackgroundColorTextAttribute")
		registerSymbol(&_kAXElementOrderHorizontalKeySymbol, frameworkHandle, "kAXElementOrderHorizontalKey")
		registerSymbol(&_kAXElementOrderVerticalKeySymbol, frameworkHandle, "kAXElementOrderVerticalKey")
		registerSymbol(&_kAXElementToFocusForLayoutChangeKeySymbol, frameworkHandle, "kAXElementToFocusForLayoutChangeKey")
		registerSymbol(&_kAXFontFamilyKeySymbol, frameworkHandle, "kAXFontFamilyKey")
		registerSymbol(&_kAXFontNameKeySymbol, frameworkHandle, "kAXFontNameKey")
		registerSymbol(&_kAXFontSizeKeySymbol, frameworkHandle, "kAXFontSizeKey")
		registerSymbol(&_kAXFontTextAttributeSymbol, frameworkHandle, "kAXFontTextAttribute")
		registerSymbol(&_kAXForegoundColorTextAttributeSymbol, frameworkHandle, "kAXForegoundColorTextAttribute")
		registerSymbol(&_kAXForegroundColorTextAttributeSymbol, frameworkHandle, "kAXForegroundColorTextAttribute")
		registerSymbol(&_kAXHasClientsWithAccessRemoteDeviceContentDidChangeSymbol, frameworkHandle, "kAXHasClientsWithAccessRemoteDeviceContentDidChange")
		registerSymbol(&_kAXInterfaceBristolKeySymbol, frameworkHandle, "kAXInterfaceBristolKey")
		registerSymbol(&_kAXInterfaceClassicInvertColorKeySymbol, frameworkHandle, "kAXInterfaceClassicInvertColorKey")
		registerSymbol(&_kAXInterfaceClassicInvertColorStatusDidChangeNotificationSymbol, frameworkHandle, "kAXInterfaceClassicInvertColorStatusDidChangeNotification")
		registerSymbol(&_kAXInterfaceCursorStatusDidChangeNotificationSymbol, frameworkHandle, "kAXInterfaceCursorStatusDidChangeNotification")
		registerSymbol(&_kAXInterfaceDifferentiateWithoutColorKeySymbol, frameworkHandle, "kAXInterfaceDifferentiateWithoutColorKey")
		registerSymbol(&_kAXInterfaceDifferentiateWithoutColorStatusDidChangeNotificationSymbol, frameworkHandle, "kAXInterfaceDifferentiateWithoutColorStatusDidChangeNotification")
		registerSymbol(&_kAXInterfaceIncreaseContrastKeySymbol, frameworkHandle, "kAXInterfaceIncreaseContrastKey")
		registerSymbol(&_kAXInterfaceIncreaseContrastStatusDidChangeNotificationSymbol, frameworkHandle, "kAXInterfaceIncreaseContrastStatusDidChangeNotification")
		registerSymbol(&_kAXInterfaceReduceMotionKeySymbol, frameworkHandle, "kAXInterfaceReduceMotionKey")
		registerSymbol(&_kAXInterfaceReduceMotionStatusDidChangeNotificationSymbol, frameworkHandle, "kAXInterfaceReduceMotionStatusDidChangeNotification")
		registerSymbol(&_kAXInterfaceReduceTextInsertionPointModulationDidChangeNotificationSymbol, frameworkHandle, "kAXInterfaceReduceTextInsertionPointModulationDidChangeNotification")
		registerSymbol(&_kAXInterfaceReduceTextInsertionPointModulationKeySymbol, frameworkHandle, "kAXInterfaceReduceTextInsertionPointModulationKey")
		registerSymbol(&_kAXInterfaceReduceTransparencyKeySymbol, frameworkHandle, "kAXInterfaceReduceTransparencyKey")
		registerSymbol(&_kAXInterfaceReduceTransparencyStatusDidChangeNotificationSymbol, frameworkHandle, "kAXInterfaceReduceTransparencyStatusDidChangeNotification")
		registerSymbol(&_kAXInterfaceShowToolbarButtonShapesKeySymbol, frameworkHandle, "kAXInterfaceShowToolbarButtonShapesKey")
		registerSymbol(&_kAXInterfaceShowToolbarButtonShapesStatusDidChangeNotificationSymbol, frameworkHandle, "kAXInterfaceShowToolbarButtonShapesStatusDidChangeNotification")
		registerSymbol(&_kAXInterfaceShowWindowTitlebarIconsKeySymbol, frameworkHandle, "kAXInterfaceShowWindowTitlebarIconsKey")
		registerSymbol(&_kAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotificationSymbol, frameworkHandle, "kAXInterfaceShowWindowTitlebarIconsStatusDidChangeNotification")
		registerSymbol(&_kAXInterfaceWhiteOnBlackKeySymbol, frameworkHandle, "kAXInterfaceWhiteOnBlackKey")
		registerSymbol(&_kAXLinkTextAttributeSymbol, frameworkHandle, "kAXLinkTextAttribute")
		registerSymbol(&_kAXListItemIndexTextAttributeSymbol, frameworkHandle, "kAXListItemIndexTextAttribute")
		registerSymbol(&_kAXListItemLevelTextAttributeSymbol, frameworkHandle, "kAXListItemLevelTextAttribute")
		registerSymbol(&_kAXListItemPrefixTextAttributeSymbol, frameworkHandle, "kAXListItemPrefixTextAttribute")
		registerSymbol(&_kAXMarkedMisspelledTextAttributeSymbol, frameworkHandle, "kAXMarkedMisspelledTextAttribute")
		registerSymbol(&_kAXMisspelledTextAttributeSymbol, frameworkHandle, "kAXMisspelledTextAttribute")
		registerSymbol(&_kAXNaturalLanguageTextAttributeSymbol, frameworkHandle, "kAXNaturalLanguageTextAttribute")
		registerSymbol(&_kAXParagraphStyleTextAttributeSymbol, frameworkHandle, "kAXParagraphStyleTextAttribute")
		registerSymbol(&_kAXReplacementStringTextAttributeSymbol, frameworkHandle, "kAXReplacementStringTextAttribute")
		registerSymbol(&_kAXShadowTextAttributeSymbol, frameworkHandle, "kAXShadowTextAttribute")
		registerSymbol(&_kAXStrikethroughColorTextAttributeSymbol, frameworkHandle, "kAXStrikethroughColorTextAttribute")
		registerSymbol(&_kAXStrikethroughTextAttributeSymbol, frameworkHandle, "kAXStrikethroughTextAttribute")
		registerSymbol(&_kAXSuperscriptTextAttributeSymbol, frameworkHandle, "kAXSuperscriptTextAttribute")
		registerSymbol(&_kAXTextAlignmentAttributeSymbol, frameworkHandle, "kAXTextAlignmentAttribute")
		registerSymbol(&_kAXTextAlignmentKeySymbol, frameworkHandle, "kAXTextAlignmentKey")
		registerSymbol(&_kAXTrustedCheckOptionPromptSymbol, frameworkHandle, "kAXTrustedCheckOptionPrompt")
		registerSymbol(&_kAXUIElementCopyHierarchyArrayAttributesKeySymbol, frameworkHandle, "kAXUIElementCopyHierarchyArrayAttributesKey")
		registerSymbol(&_kAXUIElementCopyHierarchyIncompleteResultKeySymbol, frameworkHandle, "kAXUIElementCopyHierarchyIncompleteResultKey")
		registerSymbol(&_kAXUIElementCopyHierarchyMaxArrayCountKeySymbol, frameworkHandle, "kAXUIElementCopyHierarchyMaxArrayCountKey")
		registerSymbol(&_kAXUIElementCopyHierarchyMaxDepthKeySymbol, frameworkHandle, "kAXUIElementCopyHierarchyMaxDepthKey")
		registerSymbol(&_kAXUIElementCopyHierarchyResultCountKeySymbol, frameworkHandle, "kAXUIElementCopyHierarchyResultCountKey")
		registerSymbol(&_kAXUIElementCopyHierarchyResultErrorKeySymbol, frameworkHandle, "kAXUIElementCopyHierarchyResultErrorKey")
		registerSymbol(&_kAXUIElementCopyHierarchyResultValueKeySymbol, frameworkHandle, "kAXUIElementCopyHierarchyResultValueKey")
		registerSymbol(&_kAXUIElementCopyHierarchyReturnAttributeErrorsKeySymbol, frameworkHandle, "kAXUIElementCopyHierarchyReturnAttributeErrorsKey")
		registerSymbol(&_kAXUIElementCopyHierarchySkipInspectionForAttributesKeySymbol, frameworkHandle, "kAXUIElementCopyHierarchySkipInspectionForAttributesKey")
		registerSymbol(&_kAXUIElementCopyHierarchyTruncateStringsKeySymbol, frameworkHandle, "kAXUIElementCopyHierarchyTruncateStringsKey")
		registerSymbol(&_kAXUnderlineColorTextAttributeSymbol, frameworkHandle, "kAXUnderlineColorTextAttribute")
		registerSymbol(&_kAXUnderlineTextAttributeSymbol, frameworkHandle, "kAXUnderlineTextAttribute")
		registerSymbol(&_kAXVisibleNameKeySymbol, frameworkHandle, "kAXVisibleNameKey")
	}


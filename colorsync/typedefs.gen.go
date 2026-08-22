// Code generated from Apple documentation. DO NOT EDIT.

package colorsync

import (
	"unsafe"
)

// CMMApplyTransformProc is a function a CMM provider implements to apply a color transform to image data.
//
// See: https://developer.apple.com/documentation/ColorSync/CMMApplyTransformProc
type CMMApplyTransformProc = func(uintptr, uint, uint, uint, unsafe.Pointer, ColorSyncDataDepth, uint32, uint, uint, unsafe.Pointer, ColorSyncDataDepth, uint32, uint, uintptr) bool

// CMMCreateTransformPropertyProc is a function a CMM provider implements to create a transform property for a given key.
//
// See: https://developer.apple.com/documentation/ColorSync/CMMCreateTransformPropertyProc
type CMMCreateTransformPropertyProc = func(uintptr, unsafe.Pointer, uintptr) unsafe.Pointer

// CMMInitializeLinkProfileProc is a function a CMM provider implements to initialize a device-link profile.
//
// See: https://developer.apple.com/documentation/ColorSync/CMMInitializeLinkProfileProc
type CMMInitializeLinkProfileProc = func(uintptr, uintptr, uintptr) bool

// CMMInitializeTransformProc is a function a CMM provider implements to initialize a color transform.
//
// See: https://developer.apple.com/documentation/ColorSync/CMMInitializeTransformProc
type CMMInitializeTransformProc = func(uintptr, uintptr, uintptr) bool

// ColorSyncCMMIterateCallback is a callback that the framework invokes for each installed CMM during iteration.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMIterateCallback
type ColorSyncCMMIterateCallback = func(cmm uintptr, userInfo unsafe.Pointer) bool

// ColorSyncDeviceProfileIterateCallback is a callback that ColorSync invokes for each device profile during iteration.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceProfileIterateCallback
type ColorSyncDeviceProfileIterateCallback = func(colorSyncDeviceProfileInfo uintptr, userInfo unsafe.Pointer) bool

// ColorSyncProfileIterateCallback is a callback that the framework invokes for each installed profile during iteration.
//
// See: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIterateCallback
type ColorSyncProfileIterateCallback = func(profileInfo uintptr, userInfo unsafe.Pointer) bool

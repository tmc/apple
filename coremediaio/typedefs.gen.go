// Code generated from Apple documentation. DO NOT EDIT.

package coremediaio

import (
	"unsafe"
)

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOClassID
type CMIOClassID = uint32

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOControlID
type CMIOControlID = string

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceGetSMPTETimeProc
type CMIODeviceGetSMPTETimeProc = func(unsafe.Pointer, *uint64, *byte, *uint) int

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceID
type CMIODeviceID = string

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODevicePropertyID
type CMIODevicePropertyID = string

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStreamQueueAlteredProc
type CMIODeviceStreamQueueAlteredProc = func(uint, unsafe.Pointer, unsafe.Pointer)

// CMIOExtensionProperty is a structure that defines the properties that providers, devices, and streams support.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionProperty
type CMIOExtensionProperty = string

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOHardwarePlugInRef
type CMIOHardwarePlugInRef = *CMIOHardwarePlugInInterface

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOHardwarePropertyID
type CMIOHardwarePropertyID = string

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectID
type CMIOObjectID = uint32

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertyElement
type CMIOObjectPropertyElement = uint32

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertyListenerBlock
type CMIOObjectPropertyListenerBlock = func(uint32, *CMIOObjectPropertyAddress)

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertyListenerProc
type CMIOObjectPropertyListenerProc = func(uint, uint, uintptr, unsafe.Pointer) int

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertyScope
type CMIOObjectPropertyScope = uint32

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertySelector
type CMIOObjectPropertySelector = uint32

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamID
type CMIOStreamID = string

// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamScheduledOutputNotificationProc
type CMIOStreamScheduledOutputNotificationProc = func(uint64, uint64, unsafe.Pointer)


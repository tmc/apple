// Code generated from Apple documentation. DO NOT EDIT.

package iousbhost

import (
	"unsafe"
)

// IOUSBHostCompletionHandler is the completion handler for asynchronous control, bulk, and interrupt transfers.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCompletionHandler
type IOUSBHostCompletionHandler = func(status int32, bytesTransferred uint32)

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterfaceCommandHandler
type IOUSBHostControllerInterfaceCommandHandler = func(IOUSBHostControllerInterface, IOUSBHostCIMessage)

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterfaceDoorbellHandler
type IOUSBHostControllerInterfaceDoorbellHandler = func(IOUSBHostControllerInterface, *uint32, uint32)

// IOUSBHostDevicePropertyKey is properties of a USB device that describe its state.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevicePropertyKey
type IOUSBHostDevicePropertyKey = string

// IOUSBHostInterestHandler is the callback that handles underlying service-state changes.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterestHandler
type IOUSBHostInterestHandler = func(hostObject IOUSBHostObject, messageType uint32, messageArgument unsafe.Pointer)

// IOUSBHostInterfacePropertyKey is properties of a USB interface that describe its state.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterfacePropertyKey
type IOUSBHostInterfacePropertyKey = string

// IOUSBHostIsochronousCompletionHandler is a completion handler for asynchronous isochronous transfers.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousCompletionHandler
type IOUSBHostIsochronousCompletionHandler = func(status int32, frameList *IOUSBHostIsochronousFrame)

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousTransactionCompletionHandler
type IOUSBHostIsochronousTransactionCompletionHandler = func(int32, *IOUSBHostIsochronousTransaction)

// IOUSBHostMatchingPropertyKey is properties for implementing the matching service.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostMatchingPropertyKey
type IOUSBHostMatchingPropertyKey = string

// IOUSBHostPropertyKey is properties that the USB host device and interface classes share.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPropertyKey
type IOUSBHostPropertyKey = string

// IOUSBHostTime is the absolute time.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostTime
type IOUSBHostTime = uint64

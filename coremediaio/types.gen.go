// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

// C struct types

// CMIODeviceAVCCommand
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceAVCCommand
type CMIODeviceAVCCommand struct {
	MCommand        *uint8
	MCommandLength  uint32
	MResponse       *uint8
	MResponseLength uint32
	MResponseUsed   uint32
}

// CMIODeviceRS422Command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceRS422Command
type CMIODeviceRS422Command struct {
	MCommand        *uint8
	MCommandLength  uint32
	MResponse       *uint8
	MResponseLength uint32
	MResponseUsed   uint32
}

// CMIODeviceSMPTETimeCallback
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceSMPTETimeCallback
type CMIODeviceSMPTETimeCallback struct {
	MGetSMPTETimeProc CMIODeviceGetSMPTETimeProc
	MRefCon           unsafe.Pointer
}

// CMIODeviceStreamConfiguration
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStreamConfiguration
type CMIODeviceStreamConfiguration struct {
	MNumberStreams  uint32
	MNumberChannels uint32
}

// CMIOHardwarePlugInInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOHardwarePlugInInterface
type CMIOHardwarePlugInInterface struct {
	AddRef                    func(unsafe.Pointer) uint
	DeviceProcessAVCCommand   func(uintptr, uint, uintptr) int
	DeviceProcessRS422Command func(uintptr, uint, uintptr) int
	DeviceResume              func(uintptr, uint) int
	DeviceStartStream         func(uintptr, uint, uint) int
	DeviceStopStream          func(uintptr, uint, uint) int
	DeviceSuspend             func(uintptr, uint) int
	Initialize                func(uintptr) int
	InitializeWithObjectID    func(uintptr, uint) int
	ObjectGetPropertyData     func(uintptr, uint, uintptr, uint, unsafe.Pointer, uint, *uint, unsafe.Pointer) int
	ObjectGetPropertyDataSize func(uintptr, uint, uintptr, uint, unsafe.Pointer, *uint) int
	ObjectHasProperty         func(uintptr, uint, uintptr) uint8
	ObjectIsPropertySettable  func(uintptr, uint, uintptr, *byte) int
	ObjectSetPropertyData     func(uintptr, uint, uintptr, uint, unsafe.Pointer, uint, unsafe.Pointer) int
	ObjectShow                func(uintptr, uint)
	QueryInterface            func(unsafe.Pointer, corefoundation.CFUUIDBytes, unsafe.Pointer) int
	Release                   func(unsafe.Pointer) uint
	StreamCopyBufferQueue     func(uintptr, uint, func(uint, unsafe.Pointer, unsafe.Pointer), unsafe.Pointer, uintptr) int
	StreamDeckCueTo           func(uintptr, uint, float64, uint8) int
	StreamDeckJog             func(uintptr, uint, int) int
	StreamDeckPlay            func(uintptr, uint) int
	StreamDeckStop            func(uintptr, uint) int
	Teardown                  func(uintptr) int
}

// CMIOObjectPropertyAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertyAddress
type CMIOObjectPropertyAddress struct {
	MSelector CMIOObjectPropertySelector
	MScope    CMIOObjectPropertyScope
	MElement  CMIOObjectPropertyElement
}

// CMIOStreamDeck
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeck
type CMIOStreamDeck struct {
	MStatus uint32
	MState  uint32
	MState2 uint32
}

// CMIOStreamScheduledOutputNotificationProcAndRefCon
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamScheduledOutputNotificationProcAndRefCon
type CMIOStreamScheduledOutputNotificationProcAndRefCon struct {
	ScheduledOutputNotificationProc   CMIOStreamScheduledOutputNotificationProc
	ScheduledOutputNotificationRefCon unsafe.Pointer
}

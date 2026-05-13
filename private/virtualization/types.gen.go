// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"

	"github.com/tmc/apple/private/coreml"
)

// C struct types

// AvpBreadcrumb
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/AvpBreadcrumb
type AvpBreadcrumb struct {
	Status uint
	Data   unsafe.Pointer
}

// AvpHidGenericDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/AvpHidGenericDevice
type AvpHidGenericDevice struct {
}

// AvpSerialNumber
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/AvpSerialNumber
type AvpSerialNumber struct {
	_serial_number unsafe.Pointer
}

// CGContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/CGContext
type CGContext struct {
}

// CpuExitInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/CpuExitInfo
type CpuExitInfo struct {
	Cpu_index uint
	Cpu_exit  unsafe.Pointer
}

// DebugStub
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/DebugStub
type DebugStub struct {
}

// Descriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/Descriptor
type Descriptor struct {
	Format           DiskImageFormatRef
	Cache_mode       int
	Parameters       unsafe.Pointer
	Per_io_encrypted bool
}

// DiskImageFormat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/DiskImageFormat
type DiskImageFormat struct {
}

// DispatchGroupSession
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/DispatchGroupSession
type DispatchGroupSession struct {
}

// DispatchQueue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/DispatchQueue
type DispatchQueue struct {
	_object OpaqueIdRef
}

// DispatchSource
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/DispatchSource
type DispatchSource struct {
	_object OpaqueIdRef
}

// Element
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/Element
type Element struct {
	Bytes_written       uint
	Guest_read_buffers  unsafe.Pointer
	Guest_write_buffers unsafe.Pointer
	Read_buffers        unsafe.Pointer
	Write_buffers       unsafe.Pointer
}

// FileDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/FileDescriptor
type FileDescriptor struct {
}

// IoService
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/IoService
type IoService struct {
	_object uint
}

// MailboxHandle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/MailboxHandle
type MailboxHandle struct {
}

// Mutex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/Mutex
type Mutex struct {
	_unfair_lock coreml.Os_unfair_lock_s
}

// PciDeviceLocation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/PciDeviceLocation
type PciDeviceLocation struct {
	_bus_number      uint
	_device_number   uint
	_function_number uint
}

// PipePair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/PipePair
type PipePair struct {
	For_reading FileDescriptor
	For_writing FileDescriptor
}

// PluginInstance
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/PluginInstance
type PluginInstance struct {
}

// SharedRamManager
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/SharedRamManager
type SharedRamManager struct {
	_shared_memory_map unsafe.Pointer
}

// VZVirtualMachineProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineProperties
type VZVirtualMachineProperties struct {
	Keyboard_properties                   unsafe.Pointer
	Multi_touch_device_id                 unsafe.Pointer
	Pointing_device_properties            unsafe.Pointer
	State                                 int64
	Usb_passthrough_pointing_device_count uint
	Graphics_devices                      unsafe.Pointer
}

// _IOHIDEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/__IOHIDEvent
type _IOHIDEvent struct {
}

// _SecKey
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/__SecKey
type _SecKey struct {
}

// Sockaddr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/sockaddr
type Sockaddr struct {
}

// Sockaddr_storage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/sockaddr_storage
type Sockaddr_storage struct {
	Ss_len     uint8
	Ss_family  uint8
	__ss_pad1  unsafe.Pointer
	__ss_align int64
	__ss_pad2  unsafe.Pointer
}

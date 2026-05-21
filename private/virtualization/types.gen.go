// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
)

// C struct types

// AvpBreadcrumb
type AvpBreadcrumb struct {
	Status uint
	Data   unsafe.Pointer
}

// AvpHidGenericDevice
type AvpHidGenericDevice struct {
}

// AvpSerialNumber
type AvpSerialNumber struct {
	_serial_number unsafe.Pointer
}

// CGContext
type CGContext struct {
}

// CpuExitInfo
type CpuExitInfo struct {
	Cpu_index uint
	Cpu_exit  unsafe.Pointer
}

// DebugStub
type DebugStub struct {
}

// Descriptor
type Descriptor struct {
	Format           DiskImageFormatRef
	Cache_mode       int
	Parameters       unsafe.Pointer
	Per_io_encrypted bool
}

// DiskImageFormat
type DiskImageFormat struct {
}

// DispatchGroupSession
type DispatchGroupSession struct {
}

// DispatchQueue
type DispatchQueue struct {
	_object OpaqueIdRef
}

// DispatchSource
type DispatchSource struct {
	_object OpaqueIdRef
}

// Element
type Element struct {
	Bytes_written       uint
	Guest_read_buffers  unsafe.Pointer
	Guest_write_buffers unsafe.Pointer
	Read_buffers        unsafe.Pointer
	Write_buffers       unsafe.Pointer
}

// FileDescriptor
type FileDescriptor struct {
}

// IoService
type IoService struct {
	_object uint
}

// MailboxHandle
type MailboxHandle struct {
}

// Mutex
type Mutex struct {
	_unfair_lock [4]byte
}

// PciDeviceLocation
type PciDeviceLocation struct {
	_bus_number      uint
	_device_number   uint
	_function_number uint
}

// PipePair
type PipePair struct {
	For_reading FileDescriptor
	For_writing FileDescriptor
}

// PluginInstance
type PluginInstance struct {
}

// SharedRamManager
type SharedRamManager struct {
	_shared_memory_map unsafe.Pointer
}

// VZVirtualMachineProperties
type VZVirtualMachineProperties struct {
	Keyboard_properties                   unsafe.Pointer
	Multi_touch_device_id                 unsafe.Pointer
	Pointing_device_properties            unsafe.Pointer
	State                                 int64
	Usb_passthrough_pointing_device_count uint
	Graphics_devices                      unsafe.Pointer
}

// IOHIDEvent
type IOHIDEvent struct {
}

// SecKey
type SecKey struct {
}

// Sockaddr
type Sockaddr struct {
}

// SockaddrStorage
type SockaddrStorage struct {
	Ss_len     uint8
	Ss_family  uint8
	__ss_pad1  unsafe.Pointer
	__ss_align int64
	__ss_pad2  unsafe.Pointer
}

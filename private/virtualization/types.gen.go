// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"

	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/xpc"
)

// C struct types

// AccessorProperties
type AccessorProperties struct {
}

// ApConfiguration
type ApConfiguration struct {
	Field1 uint64
	Field2 uint32
	Field3 bool
}

// Avp
type Avp struct {
	Field1 uint32
	Field2 bool
	Field3 bool
}

// AvpBreadcrumb
type AvpBreadcrumb struct {
	Status uint32
	Data   [3]uint32
}

// AvpHidGenericDevice
type AvpHidGenericDevice struct {
	Field1 uint16
	Field2 uint16
	Field3 uint32
	Field4 uint32
	Field5 [1]uint64
}

// AvpSerialNumber
type AvpSerialNumber struct {
	_serial_number [10]uint8
}

// AvpTouchScreen
type AvpTouchScreen struct {
	Field1 [1]uint64
}

// Battery
type Battery struct {
	Field1 unsafe.Pointer
}

// BifrostDevice
type BifrostDevice struct {
	Field1 unsafe.Pointer
	Field2 uint64
}

// Binary
type Binary struct {
	Field1 [3]uint64
	Field2 uint64
}

// Breadcrumb
type Breadcrumb struct {
	Field1 uint64
}

// Bridged
type Bridged struct {
	Field1 unsafe.Pointer
	Field2 bool
}

// CGColorSpace
type CGColorSpace struct {
}

// CGContext
type CGContext struct {
}

// Connection
type Connection struct {
	_object OpaqueIdRef
}

// Coprocessor
type Coprocessor struct {
	Field1  int32
	Field2  [2]uint32
	Field3  [2]uint32
	Field4  [2]uint32
	Field5  [2]uint32
	Field6  FileDescriptor
	Field7  [3]uint32
	Field8  [3]uint64
	Field9  [3]uint64
	Field10 ApConfiguration
}

// CoprocessorMailbox
type CoprocessorMailbox struct {
	Field1 unsafe.Pointer
	Field2 [2]uint32
}

// CpuEmulator
type CpuEmulator struct {
	Field1 int32
	Field2 [2]uint32
	Field3 unsafe.Pointer
	Field4 [2]uint64
}

// CpuExitContextMessenger
type CpuExitContextMessenger struct {
}

// CpuExitInfo
type CpuExitInfo struct {
	Cpu_index uint32
	Cpu_exit  [4]uint64
}

// CursorUpdate
type CursorUpdate struct {
}

// CustomMmioDeviceMessenger
type CustomMmioDeviceMessenger struct {
}

// CustomVirtioDeviceManager
type CustomVirtioDeviceManager struct {
}

// Darwin
type Darwin struct {
}

// DebugStub
type DebugStub struct {
	Field1 int32
	Field2 FileDescriptor
}

// Descriptor
type Descriptor struct {
	Format           DiskImageFormatRef
	Cache_mode       int32
	Parameters       [3]uint64
	Per_io_encrypted bool
}

// DeviceSpecificConfiguration
type DeviceSpecificConfiguration struct {
	Field1 [2]uint64
}

// DiskImageFormat
type DiskImageFormat struct {
	Field1 unsafe.Pointer
}

// DispatchGroupSession
type DispatchGroupSession struct {
	Field1 objectivec.Object
	Field2 PendingOperation
}

// DispatchQueue
type DispatchQueue struct {
	_object OpaqueIdRef
}

// DispatchSource
type DispatchSource struct {
	_object OpaqueIdRef
}

// Display
type Display struct {
	Field1 [2]uint64
	Field2 uint64
	Field3 int32
	Field4 int32
	Field5 unsafe.Pointer
	Field6 Uuid
}

// Element
type Element struct {
	Field1              unsafe.Pointer
	Bytes_written       uint32
	Guest_read_buffers  GuestDescriptors
	Guest_write_buffers GuestDescriptors
	Read_buffers        IoVector
	Write_buffers       IoVector
}

// FileDescriptor
type FileDescriptor struct {
	Field1 [1]uint32
	Field2 uint64
}

// FileDescriptors
type FileDescriptors struct {
	Field1 [2]uint32
	Field2 [2]uint32
}

// FrameUpdate
type FrameUpdate struct {
}

// FramebufferObserver
type FramebufferObserver struct {
}

// GuestDescriptors
type GuestDescriptors struct {
	_elements [3]uint64
}

// Handle
type Handle struct {
	Field1 *byte
	Field2 uint64
	Field3 unsafe.Pointer
}

// HostOnly
type HostOnly struct {
}

// IoService
type IoService struct {
	_object uint32
}

// IoVector
type IoVector struct {
	_size    uint64
	_buffers [3]uint64
}

// KeyboardEventTapMessenger
type KeyboardEventTapMessenger struct {
}

// KeyboardProperties
type KeyboardProperties struct {
	Device_id uint32
	Type      int64
}

// LinearFramebuffer
type LinearFramebuffer struct {
	Field1 [2]uint32
	Field2 uint32
	Field3 Handle
}

// Linux
type Linux struct {
	Field1 FileDescriptor
	Field2 unsafe.Pointer
	Field3 FileDescriptor
}

// Location
type Location struct {
	X float64
	Y float64
}

// MacM2ScalerAcceleratorDevice
type MacM2ScalerAcceleratorDevice struct {
}

// MacNeuralEngineAcceleratorDevice
type MacNeuralEngineAcceleratorDevice struct {
	Field1 Path
	Field2 bool
}

// MacOs
type MacOs struct {
	Field1 unsafe.Pointer
}

// MacTouchIDBiometricDevice
type MacTouchIDBiometricDevice struct {
}

// MacVideoToolboxAcceleratorDevice
type MacVideoToolboxAcceleratorDevice struct {
}

// MachO
type MachO struct {
	Field1 FileDescriptor
}

// MailboxHandle
type MailboxHandle struct {
	Field1 xpc.Endpoint
}

// MmapedMemory
type MmapedMemory struct {
	Field1 *byte
	Field2 uint64
}

// Mutex
type Mutex struct {
	_unfair_lock OSUnfairLockS
}

// Nat
type Nat struct {
	Field1 bool
}

// NoSecurity
type NoSecurity struct {
}

// OpaqueId
type OpaqueId struct {
}

// OutOfLineBuffer
type OutOfLineBuffer struct {
	Field1 *byte
}

// PanicDevice
type PanicDevice struct {
	Field1 int32
}

// ParavirtualizedGraphics
type ParavirtualizedGraphics struct {
	Field1 bool
	Field2 bool
	Field3 bool
	Field4 uint32
	Field5 [3]uint64
	Field6 int32
}

// PassthroughDevice
type PassthroughDevice struct {
	Field1 uint32
	Field2 PciDeviceLocation
}

// PciDeviceLocation
type PciDeviceLocation struct {
	Field1 uint32
	Field2 uint32
	Field3 uint32
}

// PendingOperation
type PendingOperation struct {
	Field1 Breadcrumb
}

// PipePair
type PipePair struct {
	For_reading FileDescriptor
	For_writing FileDescriptor
}

// PluginIdentifier
type PluginIdentifier struct {
	Field1 unsafe.Pointer
	Field2 unsafe.Pointer
}

// PluginInstance
type PluginInstance struct {
}

// PointingDeviceProperties
type PointingDeviceProperties struct {
	Device_id uint32
	Type      int64
}

// PowerSourceDevice
type PowerSourceDevice struct {
	Field1 unsafe.Pointer
}

// PropertyBag
type PropertyBag struct {
}

// Scanout
type Scanout struct {
	Field1 [4]uint32
	Field2 Uuid
}

// Segment
type Segment struct {
}

// SerialPort
type SerialPort struct {
	Field1 int32
	Field2 [2]uint32
	Field3 unsafe.Pointer
}

// Server
type Server struct {
}

// ServerDelegate
type ServerDelegate struct {
}

// SharedMemoryMap
type SharedMemoryMap struct {
}

// SharedRamEntry
type SharedRamEntry struct {
}

// SharedRamManager
type SharedRamManager struct {
	_shared_memory_map [2]uint64
}

// Socket
type Socket struct {
	Field1 int32
	Field2 FileDescriptor
}

// SocketDeviceObserverMessenger
type SocketDeviceObserverMessenger struct {
}

// Synthetic
type Synthetic struct {
	Field1 float64
	Field2 int32
}

// Touch
type Touch struct {
	Transducer_index objectivec.Object
	Phase            int32
	Location         Location
	Swipe_aim        [1]uint32
	Timestamp        float64
}

// UnixSocket
type UnixSocket struct {
	Field1 Path
}

// Usb
type Usb struct {
	Field1 uint32
}

// UsbTouchScreen
type UsbTouchScreen struct {
}

// Uuid
type Uuid struct {
	Field1 [16]uint8
}

// VZDisplayPresenterMessenger
type VZDisplayPresenterMessenger struct {
}

// VZVirtualMachineAccessorManagerMessenger
type VZVirtualMachineAccessorManagerMessenger struct {
}

// VZVirtualMachineAccessorMessenger
type VZVirtualMachineAccessorMessenger struct {
}

// VZVirtualMachineProperties
type VZVirtualMachineProperties struct {
	Keyboard_properties                   [3]uint64
	Multi_touch_device_id                 [2]uint32
	Pointing_device_properties            [3]uint64
	State                                 int64
	Usb_passthrough_pointing_device_count uint32
	Graphics_devices                      [3]uint64
}

// VariableStore
type VariableStore struct {
	Field1 [2]uint64
}

// VhostUser
type VhostUser struct {
	Field1 unsafe.Pointer
	Field2 int32
	Field3 uint64
	Field4 [2]uint8
	Field5 [2]uint8
	Field6 [2]uint8
	Field7 [2]uint8
}

// Virtio
type Virtio struct {
	Field1 [3]uint64
	Field2 int32
}

// VirtioQueue
type VirtioQueue struct {
}

// Vmnet
type Vmnet struct {
	Field1 [1]uint64
}

// VncAuthentication
type VncAuthentication struct {
	Password unsafe.Pointer
}

// Wall
type Wall struct {
}

// WrappingKey
type WrappingKey struct {
}

// XPC
type XPC struct {
}

// Xpc is a type alias for XPC for use in objc.Send[T] calls.
type Xpc = XPC

// XPCClient
type XPCClient struct {
	Field1 xpc.Endpoint
}

// XpcClient is a type alias for XPCClient for use in objc.Send[T] calls.
type XpcClient = XPCClient

// XPCServer
type XPCServer struct {
	Field1 Connection
}

// XpcServer is a type alias for XPCServer for use in objc.Send[T] calls.
type XpcServer = XPCServer

// CFDictionary
type CFDictionary struct {
}

// IOHIDEvent
type IOHIDEvent struct {
}

// IOHIDEventSystemClient
type IOHIDEventSystemClient struct {
}

// SecKey
type SecKey struct {
}

// Empty
type Empty struct {
}

// Long
type Long struct {
	Field1 *byte
	Field2 uint64
	Field3 objectivec.Object
	Field4 objectivec.Object
}

// Rep
type Rep struct {
	Field1 int16
	Field2 int
}

// Repr
type Repr struct {
	Field1 unsafe.Pointer
	Field2 bool
}

// SharedWeakCount
type SharedWeakCount struct {
}

// Short
type Short struct {
	Field1 [23]int8
	Field2 objectivec.Object
	Field3 objectivec.Object
}

// Union
type Union struct {
	Field1 SerialPort
	Field2 objectivec.Object
}

// OSUnfairLockS
type OSUnfairLockS struct {
	_os_unfair_lock_opaque uint32
}

// Os_unfair_lock_s is a type alias for OSUnfairLockS for use in objc.Send[T] calls.
type Os_unfair_lock_s = OSUnfairLockS

// Path
type Path struct {
	Field1 unsafe.Pointer
}

// Sockaddr
type Sockaddr struct {
	Field1 uint8
	Field2 uint8
	Field3 [14]int8
}

// SockaddrStorage
type SockaddrStorage struct {
	Ss_len     uint8
	Ss_family  uint8
	__ss_pad1  [6]int8
	__ss_align int64
	__ss_pad2  [112]int8
}

// Sockaddr_storage is a type alias for SockaddrStorage for use in objc.Send[T] calls.
type Sockaddr_storage = SockaddrStorage

// SockaddrUn
type SockaddrUn struct {
	Field1 uint8
	Field2 uint8
	Field3 [104]int8
}

// Sockaddr_un is a type alias for SockaddrUn for use in objc.Send[T] calls.
type Sockaddr_un = SockaddrUn

// VmnetNetwork
type VmnetNetwork struct {
}

// Vmnet_network is a type alias for VmnetNetwork for use in objc.Send[T] calls.
type Vmnet_network = VmnetNetwork

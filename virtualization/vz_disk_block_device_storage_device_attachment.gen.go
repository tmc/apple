// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZDiskBlockDeviceStorageDeviceAttachment] class.
var (
	_VZDiskBlockDeviceStorageDeviceAttachmentClass     VZDiskBlockDeviceStorageDeviceAttachmentClass
	_VZDiskBlockDeviceStorageDeviceAttachmentClassOnce sync.Once
)

func getVZDiskBlockDeviceStorageDeviceAttachmentClass() VZDiskBlockDeviceStorageDeviceAttachmentClass {
	_VZDiskBlockDeviceStorageDeviceAttachmentClassOnce.Do(func() {
		_VZDiskBlockDeviceStorageDeviceAttachmentClass = VZDiskBlockDeviceStorageDeviceAttachmentClass{class: objc.GetClass("VZDiskBlockDeviceStorageDeviceAttachment")}
	})
	return _VZDiskBlockDeviceStorageDeviceAttachmentClass
}

// GetVZDiskBlockDeviceStorageDeviceAttachmentClass returns the class object for VZDiskBlockDeviceStorageDeviceAttachment.
func GetVZDiskBlockDeviceStorageDeviceAttachmentClass() VZDiskBlockDeviceStorageDeviceAttachmentClass {
	return getVZDiskBlockDeviceStorageDeviceAttachmentClass()
}

type VZDiskBlockDeviceStorageDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDiskBlockDeviceStorageDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDiskBlockDeviceStorageDeviceAttachmentClass) Alloc() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A storage device attachment that uses a disk to store data.
//
// # Overview
//
// The disk block device implements a storage attachment by using an actual
// disk rather than a disk image on a file system.
//
// In the following example, a disk device at `/dev/rdisk42` executes the I/O
// operations directly on that disk rather than through a file system:
//
// By default, only the `root` user can access the disk file handle. Running
// virtual machines as `root` isn’t recommended. The best practice is to
// open the file in a separate process that has `root` privileges, then pass
// the open file descriptor using XPC or a Unix socket to a non-`root` process
// running Virtualization. For more information about Unix sockets, see
// [Streams, Sockets, and Ports]; for more information on XPC services, see
// the [XPC] framework documentation.
//
// # Initializers
//
//   - [VZDiskBlockDeviceStorageDeviceAttachment.InitWithFileHandleReadOnlySynchronizationModeError]: Creates a new block storage device attachment from a file handle and with the specified access mode, synchronization mode, and error object that you provide.
//
// # Getting the block storage device details
//
//   - [VZDiskBlockDeviceStorageDeviceAttachment.FileHandle]: A file handle to a block device.
//   - [VZDiskBlockDeviceStorageDeviceAttachment.ReadOnly]: A Boolean value that indicates whether this disk attachment is read-only; otherwise, if the file handle allows writes, the device can write data into it.
//   - [VZDiskBlockDeviceStorageDeviceAttachment.SynchronizationMode]: The value that defines how the disk synchronizes with the underlying storage when the guest operating system flushes data.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment
//
// [Streams, Sockets, and Ports]: https://developer.apple.com/documentation/Foundation/streams-sockets-and-ports
// [XPC]: https://developer.apple.com/documentation/XPC
type VZDiskBlockDeviceStorageDeviceAttachment struct {
	VZStorageDeviceAttachment
}

// VZDiskBlockDeviceStorageDeviceAttachmentFromID constructs a [VZDiskBlockDeviceStorageDeviceAttachment] from an objc.ID.
//
// A storage device attachment that uses a disk to store data.
func VZDiskBlockDeviceStorageDeviceAttachmentFromID(id objc.ID) VZDiskBlockDeviceStorageDeviceAttachment {
	return VZDiskBlockDeviceStorageDeviceAttachment{VZStorageDeviceAttachment: VZStorageDeviceAttachmentFromID(id)}
}

// NOTE: VZDiskBlockDeviceStorageDeviceAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZDiskBlockDeviceStorageDeviceAttachment] class.
//
// # Initializers
//
//   - [IVZDiskBlockDeviceStorageDeviceAttachment.InitWithFileHandleReadOnlySynchronizationModeError]: Creates a new block storage device attachment from a file handle and with the specified access mode, synchronization mode, and error object that you provide.
//
// # Getting the block storage device details
//
//   - [IVZDiskBlockDeviceStorageDeviceAttachment.FileHandle]: A file handle to a block device.
//   - [IVZDiskBlockDeviceStorageDeviceAttachment.ReadOnly]: A Boolean value that indicates whether this disk attachment is read-only; otherwise, if the file handle allows writes, the device can write data into it.
//   - [IVZDiskBlockDeviceStorageDeviceAttachment.SynchronizationMode]: The value that defines how the disk synchronizes with the underlying storage when the guest operating system flushes data.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment
type IVZDiskBlockDeviceStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment

	// Topic: Initializers

	// Creates a new block storage device attachment from a file handle and with the specified access mode, synchronization mode, and error object that you provide.
	InitWithFileHandleReadOnlySynchronizationModeError(fileHandle foundation.NSFileHandle, readOnly bool, synchronizationMode VZDiskSynchronizationMode) (VZDiskBlockDeviceStorageDeviceAttachment, error)

	// Topic: Getting the block storage device details

	// A file handle to a block device.
	FileHandle() foundation.NSFileHandle
	// A Boolean value that indicates whether this disk attachment is read-only; otherwise, if the file handle allows writes, the device can write data into it.
	ReadOnly() bool
	// The value that defines how the disk synchronizes with the underlying storage when the guest operating system flushes data.
	SynchronizationMode() VZDiskSynchronizationMode
}

// Init initializes the instance.
func (d VZDiskBlockDeviceStorageDeviceAttachment) Init() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VZDiskBlockDeviceStorageDeviceAttachment) Autorelease() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDiskBlockDeviceStorageDeviceAttachment creates a new VZDiskBlockDeviceStorageDeviceAttachment instance.
func NewVZDiskBlockDeviceStorageDeviceAttachment() VZDiskBlockDeviceStorageDeviceAttachment {
	class := getVZDiskBlockDeviceStorageDeviceAttachmentClass()
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new block storage device attachment from a file handle and with
// the specified access mode, synchronization mode, and error object that you
// provide.
//
// fileHandle: The [FileHandle] to a block device to attach to this VM.
//
// readOnly: A Boolean value that indicates whether this disk attachment is read-only;
// otherwise, if the file handle allows writes, the device can write data into
// it.
//
// synchronizationMode: The [VZDiskSynchronizationMode] value that defines how the disk
// synchronizes with the underlying storage when the guest operating system
// flushes data.
//
// # Discussion
//
// The `readOnly` parameter affects how the Virtualization framework exposes
// the disk to the guest operating system by the storage controller. If you
// intend to use the disk in read-only mode, it’s also a best practice to
// open the file handle as read-only.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/init(fileHandle:readOnly:synchronizationMode:)
//
// [FileHandle]: https://developer.apple.com/documentation/Foundation/FileHandle
// [VZDiskSynchronizationMode]: https://developer.apple.com/documentation/Virtualization/VZDiskSynchronizationMode
func NewDiskBlockDeviceStorageDeviceAttachmentWithFileHandleReadOnlySynchronizationModeError(fileHandle foundation.NSFileHandle, readOnly bool, synchronizationMode VZDiskSynchronizationMode) (VZDiskBlockDeviceStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	instance := getVZDiskBlockDeviceStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileHandle:readOnly:synchronizationMode:error:"), fileHandle, readOnly, synchronizationMode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZDiskBlockDeviceStorageDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZDiskBlockDeviceStorageDeviceAttachmentFromID(rv), nil
}

// Creates a new block storage device attachment from a file handle and with
// the specified access mode, synchronization mode, and error object that you
// provide.
//
// fileHandle: The [FileHandle] to a block device to attach to this VM.
//
// readOnly: A Boolean value that indicates whether this disk attachment is read-only;
// otherwise, if the file handle allows writes, the device can write data into
// it.
//
// synchronizationMode: The [VZDiskSynchronizationMode] value that defines how the disk
// synchronizes with the underlying storage when the guest operating system
// flushes data.
//
// # Discussion
//
// The `readOnly` parameter affects how the Virtualization framework exposes
// the disk to the guest operating system by the storage controller. If you
// intend to use the disk in read-only mode, it’s also a best practice to
// open the file handle as read-only.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/init(fileHandle:readOnly:synchronizationMode:)
//
// [FileHandle]: https://developer.apple.com/documentation/Foundation/FileHandle
// [VZDiskSynchronizationMode]: https://developer.apple.com/documentation/Virtualization/VZDiskSynchronizationMode
func (d VZDiskBlockDeviceStorageDeviceAttachment) InitWithFileHandleReadOnlySynchronizationModeError(fileHandle foundation.NSFileHandle, readOnly bool, synchronizationMode VZDiskSynchronizationMode) (VZDiskBlockDeviceStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithFileHandle:readOnly:synchronizationMode:error:"), fileHandle, readOnly, synchronizationMode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZDiskBlockDeviceStorageDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZDiskBlockDeviceStorageDeviceAttachmentFromID(rv), nil

}

// A file handle to a block device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/fileHandle
func (d VZDiskBlockDeviceStorageDeviceAttachment) FileHandle() foundation.NSFileHandle {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileHandle"))
	return foundation.NSFileHandleFromID(objc.ID(rv))
}

// A Boolean value that indicates whether this disk attachment is read-only;
// otherwise, if the file handle allows writes, the device can write data into
// it.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/isReadOnly
func (d VZDiskBlockDeviceStorageDeviceAttachment) ReadOnly() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isReadOnly"))
	return rv
}

// The value that defines how the disk synchronizes with the underlying
// storage when the guest operating system flushes data.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/synchronizationMode
func (d VZDiskBlockDeviceStorageDeviceAttachment) SynchronizationMode() VZDiskSynchronizationMode {
	rv := objc.Send[VZDiskSynchronizationMode](d.ID, objc.Sel("synchronizationMode"))
	return VZDiskSynchronizationMode(rv)
}

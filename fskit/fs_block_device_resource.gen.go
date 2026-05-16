// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSBlockDeviceResource] class.
var (
	_FSBlockDeviceResourceClass     FSBlockDeviceResourceClass
	_FSBlockDeviceResourceClassOnce sync.Once
)

func getFSBlockDeviceResourceClass() FSBlockDeviceResourceClass {
	_FSBlockDeviceResourceClassOnce.Do(func() {
		_FSBlockDeviceResourceClass = FSBlockDeviceResourceClass{class: objc.GetClass("FSBlockDeviceResource")}
	})
	return _FSBlockDeviceResourceClass
}

// GetFSBlockDeviceResourceClass returns the class object for FSBlockDeviceResource.
func GetFSBlockDeviceResourceClass() FSBlockDeviceResourceClass {
	return getFSBlockDeviceResourceClass()
}

type FSBlockDeviceResourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSBlockDeviceResourceClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSBlockDeviceResourceClass) Alloc() FSBlockDeviceResource {
	rv := objc.Send[FSBlockDeviceResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A resource that represents a block storage disk partition.
//
// # Overview
//
// A [FSBlockDeviceResource] can exist in either a proxied or nonproxied
// version. Only the `fskitd` daemon creates “real” (nonproxied) instances
// of this class. Client applications and daemons create proxy objects for
// requests, and `fskitd` opens the underlying device during the processing of
// the request.
//
// This class wraps a file descriptor for a disk device or partition. Its
// fundamental identifier is the BSD disk name ([FSBlockDeviceResource.BSDName]) for the underlying
// IOMedia object. However, [FSBlockDeviceResource] doesn’t expose the
// underlying file descriptor. Instead, it provides accessor methods that can
// read from and write to the partition, either directly or using the kernel
// buffer cache.
//
// When you use a [FSBlockDeviceResource], your file system implementation
// also conforms to a maintenance operation protocol. These protocols add
// support for checking, repairing, and optionally formatting file systems.
// The system doesn’t mount block device file systems until they pass a file
// system check. For an [FSUnaryFileSystem] that uses [FSBlockDeviceResource],
// conform to [FSManageableResourceMaintenanceOperations].
//
// # Accessing resource properties
//
//   - [FSBlockDeviceResource.BSDName]: The device name of the resource.
//   - [FSBlockDeviceResource.Writable]: A Boolean property that indicates whether the resource can write data to the device.
//   - [FSBlockDeviceResource.BlockCount]: The block count on this resource.
//   - [FSBlockDeviceResource.BlockSize]: The logical block size, the size of data blocks used by the file system.
//   - [FSBlockDeviceResource.PhysicalBlockSize]: The sector size of the device.
//
// # Reading and writing data with kernel buffer cache
//
//   - [FSBlockDeviceResource.MetadataFlushWithError]: Synchronously flushes the resource’s buffer cache.
//   - [FSBlockDeviceResource.AsynchronousMetadataFlushWithError]: Asynchronously flushes the resource’s buffer cache.
//   - [FSBlockDeviceResource.MetadataClearWithDelayedWritesError]: Clears the given ranges within the buffer cache.
//   - [FSBlockDeviceResource.MetadataPurgeError]: Synchronously purges the given ranges from the buffer cache.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource
type FSBlockDeviceResource struct {
	FSResource
}

// FSBlockDeviceResourceFromID constructs a [FSBlockDeviceResource] from an objc.ID.
//
// A resource that represents a block storage disk partition.
func FSBlockDeviceResourceFromID(id objc.ID) FSBlockDeviceResource {
	return FSBlockDeviceResource{FSResource: FSResourceFromID(id)}
}

// NOTE: FSBlockDeviceResource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSBlockDeviceResource] class.
//
// # Accessing resource properties
//
//   - [IFSBlockDeviceResource.BSDName]: The device name of the resource.
//   - [IFSBlockDeviceResource.Writable]: A Boolean property that indicates whether the resource can write data to the device.
//   - [IFSBlockDeviceResource.BlockCount]: The block count on this resource.
//   - [IFSBlockDeviceResource.BlockSize]: The logical block size, the size of data blocks used by the file system.
//   - [IFSBlockDeviceResource.PhysicalBlockSize]: The sector size of the device.
//
// # Reading and writing data with kernel buffer cache
//
//   - [IFSBlockDeviceResource.MetadataFlushWithError]: Synchronously flushes the resource’s buffer cache.
//   - [IFSBlockDeviceResource.AsynchronousMetadataFlushWithError]: Asynchronously flushes the resource’s buffer cache.
//   - [IFSBlockDeviceResource.MetadataClearWithDelayedWritesError]: Clears the given ranges within the buffer cache.
//   - [IFSBlockDeviceResource.MetadataPurgeError]: Synchronously purges the given ranges from the buffer cache.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource
type IFSBlockDeviceResource interface {
	IFSResource

	// Topic: Accessing resource properties

	// The device name of the resource.
	BSDName() string
	// A Boolean property that indicates whether the resource can write data to the device.
	Writable() bool
	// The block count on this resource.
	BlockCount() uint64
	// The logical block size, the size of data blocks used by the file system.
	BlockSize() uint64
	// The sector size of the device.
	PhysicalBlockSize() uint64

	// Topic: Reading and writing data with kernel buffer cache

	// Synchronously flushes the resource’s buffer cache.
	MetadataFlushWithError() (bool, error)
	// Asynchronously flushes the resource’s buffer cache.
	AsynchronousMetadataFlushWithError() (bool, error)
	// Clears the given ranges within the buffer cache.
	MetadataClearWithDelayedWritesError(rangesToClear []FSMetadataRange, withDelayedWrites bool) (bool, error)
	// Synchronously purges the given ranges from the buffer cache.
	MetadataPurgeError(rangesToPurge []FSMetadataRange) (bool, error)
}

// Init initializes the instance.
func (b FSBlockDeviceResource) Init() FSBlockDeviceResource {
	rv := objc.Send[FSBlockDeviceResource](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b FSBlockDeviceResource) Autorelease() FSBlockDeviceResource {
	rv := objc.Send[FSBlockDeviceResource](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSBlockDeviceResource creates a new FSBlockDeviceResource instance.
func NewFSBlockDeviceResource() FSBlockDeviceResource {
	class := getFSBlockDeviceResourceClass()
	rv := objc.Send[FSBlockDeviceResource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Synchronously flushes the resource’s buffer cache.
//
// # Discussion
//
// This method flushes data previously written with
// [delayedMetadataWriteFrom:startingAt:length:error:] to the resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataFlush()
//
// [delayedMetadataWriteFrom:startingAt:length:error:]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/delayedMetadataWriteFrom:startingAt:length:error:
func (b FSBlockDeviceResource) MetadataFlushWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("metadataFlushWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("metadataFlushWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// Asynchronously flushes the resource’s buffer cache.
//
// # Discussion
//
// This method schedules a flush of data previously written with
// [delayedMetadataWriteFrom:startingAt:length:error:] to the resource and
// returns immediately without blocking. This method doesn’t wait to check
// the flush’s status. If an error prevents the flush from being scheduled,
// the error is indicated by the in-out `error` parameter.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/asynchronousMetadataFlush()
//
// [delayedMetadataWriteFrom:startingAt:length:error:]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/delayedMetadataWriteFrom:startingAt:length:error:
func (b FSBlockDeviceResource) AsynchronousMetadataFlushWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("asynchronousMetadataFlushWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("asynchronousMetadataFlushWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// Clears the given ranges within the buffer cache.
//
// rangesToClear: The metadata ranges to clear.
//
// withDelayedWrites: A Boolean value that determines whether to perform the clear operation with
// delayed writes. The delay works in the same manner as
// [delayedMetadataWriteFrom:startingAt:length:error:]. When using delayed
// writes, the client can flush the metadata with [MetadataFlushWithError] or
// [AsynchronousMetadataFlushWithError]. The system also flushes stale data in
// the buffer cache periodically.
//
// # Discussion
//
// This method clears the specified ranges in the resource’s buffer cache by
// writing zeroes into them.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataClear(_:withDelayedWrites:)
//
// [delayedMetadataWriteFrom:startingAt:length:error:]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/delayedMetadataWriteFrom:startingAt:length:error:
func (b FSBlockDeviceResource) MetadataClearWithDelayedWritesError(rangesToClear []FSMetadataRange, withDelayedWrites bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("metadataClear:withDelayedWrites:error:"), objectivec.IObjectSliceToNSArray(rangesToClear), withDelayedWrites, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("metadataClear:withDelayedWrites:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Synchronously purges the given ranges from the buffer cache.
//
// rangesToPurge: The metadata ranges to purge.
//
// # Discussion
//
// This method removes the given ranges from the resource’s buffer cache.
// This process drops any dirty data in the cache, preventing the data from
// reaching the device.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataPurge(_:)
func (b FSBlockDeviceResource) MetadataPurgeError(rangesToPurge []FSMetadataRange) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("metadataPurge:error:"), objectivec.IObjectSliceToNSArray(rangesToPurge), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("metadataPurge:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The device name of the resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/bsdName
func (b FSBlockDeviceResource) BSDName() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("BSDName"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean property that indicates whether the resource can write data to
// the device.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/isWritable
func (b FSBlockDeviceResource) Writable() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isWritable"))
	return rv
}

// The block count on this resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/blockCount
func (b FSBlockDeviceResource) BlockCount() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("blockCount"))
	return rv
}

// The logical block size, the size of data blocks used by the file system.
//
// # Discussion
//
// This is equivalent to the [DKIOCGETBLOCKSIZE] device parameter.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/blockSize
func (b FSBlockDeviceResource) BlockSize() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("blockSize"))
	return rv
}

// The sector size of the device.
//
// # Discussion
//
// This is equivalent to the [DKIOCGETPHYSICALBLOCKSIZE] device parameter.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/physicalBlockSize
func (b FSBlockDeviceResource) PhysicalBlockSize() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("physicalBlockSize"))
	return rv
}

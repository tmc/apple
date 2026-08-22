// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"context"
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
// fundamental identifier is the BSD disk name
// ([FSBlockDeviceResource.BSDName]) for the underlying IOMedia object.
// However, [FSBlockDeviceResource] doesn’t expose the underlying file
// descriptor. Instead, it provides accessor methods that can read from and
// write to the partition, either directly or using the kernel buffer cache.
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
//   - [FSBlockDeviceResource.IsWritable]: A Boolean property that indicates whether the resource can write data to the device.
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
//   - [IFSBlockDeviceResource.IsWritable]: A Boolean property that indicates whether the resource can write data to the device.
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
	IsWritable() bool
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

	// Writes file system metadata from a buffer to a cache, prior to flushing it to the resource.
	DelayedMetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset int64, length uintptr) (bool, error)
	// Synchronously reads file system metadata from the resource into a buffer.
	MetadataReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset int64, length uintptr) (bool, error)
	// Synchronously writes file system metadata from a buffer to the resource.
	MetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset int64, length uintptr) (bool, error)
	// Reads data from the resource into a buffer and executes a block afterwards.
	ReadIntoStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset int64, length uintptr, completionHandler size_tErrorHandler)
	// Synchronously reads data from the resource into a buffer.
	ReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset int64, length uintptr) (uintptr, error)
	// Writes data from from a buffer to the resource and executes a block afterwards.
	WriteFromStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset int64, length uintptr, completionHandler size_tErrorHandler)
	// Synchronously writes data from from a buffer to the resource and executes a block afterwards.
	WriteFromStartingAtLengthError(buffer unsafe.Pointer, offset int64, length uintptr) (uintptr, error)
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

// See: https://developer.apple.com/documentation/FSKit/FSResource/init(coder:)
func NewBlockDeviceResourceWithCoder(coder foundation.INSCoder) FSBlockDeviceResource {
	instance := getFSBlockDeviceResourceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return FSBlockDeviceResourceFromID(rv)
}

// Synchronously flushes the resource’s buffer cache.
//
// # Discussion
//
// This method flushes data previously written with
// [FSBlockDeviceResource.DelayedMetadataWriteFromStartingAtLengthError] to
// the resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataFlush()
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
// [FSBlockDeviceResource.DelayedMetadataWriteFromStartingAtLengthError] to
// the resource and returns immediately without blocking. This method
// doesn’t wait to check the flush’s status. If an error prevents the
// flush from being scheduled, the error is indicated by the in-out `error`
// parameter.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/asynchronousMetadataFlush()
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
// [FSBlockDeviceResource.DelayedMetadataWriteFromStartingAtLengthError]. When
// using delayed writes, the client can flush the metadata with
// [FSBlockDeviceResource.MetadataFlushWithError] or
// [FSBlockDeviceResource.AsynchronousMetadataFlushWithError]. The system also
// flushes stale data in the buffer cache periodically.
//
// # Discussion
//
// This method clears the specified ranges in the resource’s buffer cache by
// writing zeroes into them.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataClear(_:withDelayedWrites:)
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

// Writes file system metadata from a buffer to a cache, prior to flushing it
// to the resource.
//
// buffer: A buffer to provide the data.
//
// offset: The offset into the resource from which to start writing.
//
// length: The number of bytes to writing.
//
// error: On return, any error encountered while writing data, or `nil` if no error
// occurred.
//
// # Return Value
//
// A Boolean value indicating whether the metadata write succeeded.
//
// # Discussion
//
// This method provides access to the Kernel Buffer Cache, which is the
// primary system cache for file system metadata. Unlike equivalent kernel
// APIs, this method doesn’t hold any kernel-level claim to the underlying
// buffers.
//
// This method is equivalent to
// [FSBlockDeviceResource.MetadataWriteFromStartingAtLengthError], except that
// it writes data to the resource’s buffer cache instead of writing to disk
// immediately. To ensure writing data to disk, the client must flush the
// metadata by calling [FSBlockDeviceResource.MetadataFlushWithError] or
// [FSBlockDeviceResource.AsynchronousMetadataFlushWithError].
//
// Delayed writes offer two significant advantages:
//
// - Delayed writes are more performant, since the file system can avoid
// waiting for the actual write, reducing I/O latency. - When writing to a
// specific range repeatedly, delayed writes allow the file system to flush
// data to the disk only when necessary. This reduces disk usage by
// eliminating unnecessary writes.
//
// For the write to succeed, requests must conform to any transfer
// requirements of the underlying resource. Disk drives typically require
// sector (`physicalBlockSize`) addressed operations of one or more
// sector-aligned offsets.
//
// This method doesn’t support partial writing of metadata.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/delayedMetadataWriteFrom:startingAt:length:error:
func (b FSBlockDeviceResource) DelayedMetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset int64, length uintptr) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("delayedMetadataWriteFrom:startingAt:length:error:"), buffer, offset, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("delayedMetadataWriteFrom:startingAt:length:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Synchronously reads file system metadata from the resource into a buffer.
//
// buffer: A buffer to receive the data.
//
// offset: The offset into the resource from which to start reading.
//
// length: The number of bytes to read.
//
// error: On return, any error encountered while reading data, or `nil` if no error
// occurred.
//
// # Return Value
//
// A Boolean value indicating whether the metadata read succeeded.
//
// # Discussion
//
// This method provides access to the Kernel Buffer Cache, which is the
// primary system cache for file system metadata. Unlike equivalent kernel
// APIs, this method doesn’t hold any kernel-level claim to the underlying
// buffers.
//
// For the read to succeed, requests must conform to any transfer requirements
// of the underlying resource. Disk drives typically require sector
// (`physicalBlockSize`) addressed operations of one or more sector-aligned
// offsets.
//
// This method doesn’t support partial reading of metadata.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataReadInto:startingAt:length:error:
func (b FSBlockDeviceResource) MetadataReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset int64, length uintptr) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("metadataReadInto:startingAt:length:error:"), buffer, offset, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("metadataReadInto:startingAt:length:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Synchronously writes file system metadata from a buffer to the resource.
//
// buffer: A buffer to provide the data.
//
// offset: The offset into the resource from which to start writing.
//
// length: The number of bytes to writing.
//
// error: On return, any error encountered while writing data, or `nil` if no error
// occurred.
//
// # Return Value
//
// A Boolean value indicating whether the metadata write succeeded.
//
// # Discussion
//
// This method provides access to the Kernel Buffer Cache, which is the
// primary system cache for file system metadata. Unlike equivalent kernel
// APIs, this method doesn’t hold any kernel-level claim to the underlying
// buffers.
//
// For the write to succeed, requests must conform to any transfer
// requirements of the underlying resource. Disk drives typically require
// sector (`physicalBlockSize`) addressed operations of one or more
// sector-aligned offsets.
//
// This method doesn’t support partial writing of metadata.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataWriteFrom:startingAt:length:error:
func (b FSBlockDeviceResource) MetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset int64, length uintptr) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("metadataWriteFrom:startingAt:length:error:"), buffer, offset, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("metadataWriteFrom:startingAt:length:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Reads data from the resource into a buffer and executes a block afterwards.
//
// buffer: A buffer to receive the data.
//
// offset: The offset into the resource from which to start reading.
//
// length: A maximum number of bytes to read. The completion handler receives a
// parameter with the actual number of bytes read.
//
// completionHandler: A block that executes after the read operation completes. If successful,
// the first parameter contains the number of bytes actually read. In the case
// of an error, the second parameter contains a non-`nil` error. This value is
// [EFAULT] if `buffer` is [NULL], or `errno` if reading from the resource
// failed.
//
// # Discussion
//
// For the read to succeed, requests must conform to any transfer requirements
// of the underlying resource. Disk drives typically require sector
// (`physicalBlockSize`) addressed operations of one or more sector-aligned
// offsets.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/readInto:startingAt:length:completionHandler:
func (b FSBlockDeviceResource) ReadIntoStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset int64, length uintptr, completionHandler size_tErrorHandler) {
	_block3, _ := Newsize_tErrorBlock(completionHandler)
	objc.Send[objc.ID](b.ID, objc.Sel("readInto:startingAt:length:completionHandler:"), buffer, offset, length, _block3)
}

// Synchronously reads data from the resource into a buffer.
//
// buffer: A buffer to receive the data.
//
// offset: The offset into the resource from which to start reading.
//
// length: A maximum number of bytes to read. The method’s return value contains the
// actual number of bytes read.
//
// error: On return, any error encountered while reading data, or `nil` if no error
// occurred.
//
// # Return Value
//
// The actual number of bytes read.
//
// # Discussion
//
// This is a synchronous version of
// [FSBlockDeviceResource.ReadIntoStartingAtLengthCompletionHandler].
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/readInto:startingAt:length:error:
func (b FSBlockDeviceResource) ReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset int64, length uintptr) (uintptr, error) {
	var errorPtr objc.ID
	rv := objc.Send[uintptr](b.ID, objc.Sel("readInto:startingAt:length:error:"), buffer, offset, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// Writes data from from a buffer to the resource and executes a block
// afterwards.
//
// buffer: A buffer to provide the data.
//
// offset: The offset into the resource from which to start writing.
//
// length: A maximum number of bytes to write. The completion handler receives a
// parameter with the actual number of bytes write.
//
// completionHandler: A block that executes after the write operation completes. If successful,
// the first parameter contains the number of bytes actually written. In the
// case of an error, the second parameter contains a non-`nil` error. This
// value is [EFAULT] if `buffer` is [NULL], or `errno` if writing to the
// resource failed.
//
// # Discussion
//
// For the write to succeed, requests must conform to any transfer
// requirements of the underlying resource. Disk drives typically require
// sector (`physicalBlockSize`) addressed operations of one or more
// sector-aligned offsets.
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/writeFrom:startingAt:length:completionHandler:
func (b FSBlockDeviceResource) WriteFromStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset int64, length uintptr, completionHandler size_tErrorHandler) {
	_block3, _ := Newsize_tErrorBlock(completionHandler)
	objc.Send[objc.ID](b.ID, objc.Sel("writeFrom:startingAt:length:completionHandler:"), buffer, offset, length, _block3)
}

// Synchronously writes data from from a buffer to the resource and executes a
// block afterwards.
//
// buffer: A buffer to provide the data.
//
// offset: The offset into the resource from which to start writing.
//
// length: A maximum number of bytes to write. The completion handler receives a
// parameter with the actual number of bytes write.
//
// error: On return, any error encountered while writing data, or `nil` if no error
// occurred.
//
// # Return Value
//
// The actual number of bytes written.
//
// # Discussion
//
// This is a synchronous version of
// [FSBlockDeviceResource.WriteFromStartingAtLengthCompletionHandler].
//
// See: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/writeFrom:startingAt:length:error:
func (b FSBlockDeviceResource) WriteFromStartingAtLengthError(buffer unsafe.Pointer, offset int64, length uintptr) (uintptr, error) {
	var errorPtr objc.ID
	rv := objc.Send[uintptr](b.ID, objc.Sel("writeFrom:startingAt:length:error:"), buffer, offset, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
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
func (b FSBlockDeviceResource) IsWritable() bool {
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

// ReadIntoStartingAtLength is a synchronous wrapper around [FSBlockDeviceResource.ReadIntoStartingAtLengthCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (b FSBlockDeviceResource) ReadIntoStartingAtLength(ctx context.Context, buffer unsafe.Pointer, offset int64, length uintptr) (uintptr, error) {
	type result struct {
		val uintptr
		err error
	}
	done := make(chan result, 1)
	b.ReadIntoStartingAtLengthCompletionHandler(buffer, offset, length, func(val uintptr, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// WriteFromStartingAtLength is a synchronous wrapper around [FSBlockDeviceResource.WriteFromStartingAtLengthCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (b FSBlockDeviceResource) WriteFromStartingAtLength(ctx context.Context, buffer unsafe.Pointer, offset int64, length uintptr) (uintptr, error) {
	type result struct {
		val uintptr
		err error
	}
	done := make(chan result, 1)
	b.WriteFromStartingAtLengthCompletionHandler(buffer, offset, length, func(val uintptr, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

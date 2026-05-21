// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSStatFSResult] class.
var (
	_FSStatFSResultClass     FSStatFSResultClass
	_FSStatFSResultClassOnce sync.Once
)

func getFSStatFSResultClass() FSStatFSResultClass {
	_FSStatFSResultClassOnce.Do(func() {
		_FSStatFSResultClass = FSStatFSResultClass{class: objc.GetClass("FSStatFSResult")}
	})
	return _FSStatFSResultClass
}

// GetFSStatFSResultClass returns the class object for FSStatFSResult.
func GetFSStatFSResultClass() FSStatFSResultClass {
	return getFSStatFSResultClass()
}

type FSStatFSResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSStatFSResultClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSStatFSResultClass) Alloc() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A type used to report a volume’s statistics.
//
// # Overview
//
// The names of this type’s properties match those in the `statfs` structure
// in `statfs(2)`, which reports these values for an FSKit file system. All
// numeric properties default to `0`. Override these values, unless a given
// property has no meaningful value to provide.
//
// For the read-only [FSStatFSResult.FileSystemTypeName], set this value with
// the designated initializer.
//
// # Initializers
//
//   - [FSStatFSResult.InitWithCoder]
//   - [FSStatFSResult.InitWithFileSystemTypeName]: Creates an statistics result instance, using the given file system type name.
//
// # Instance Properties
//
//   - [FSStatFSResult.AvailableBlocks]: A property for the number of free blocks available to a non-superuser on the volume.
//   - [FSStatFSResult.SetAvailableBlocks]
//   - [FSStatFSResult.AvailableBytes]: A property for the amount of space available to users, in bytes, in the volume.
//   - [FSStatFSResult.SetAvailableBytes]
//   - [FSStatFSResult.BlockSize]: A property for the volume’s block size, in bytes.
//   - [FSStatFSResult.SetBlockSize]
//   - [FSStatFSResult.FileSystemSubType]: A property for the file system’s subtype or flavor.
//   - [FSStatFSResult.SetFileSystemSubType]
//   - [FSStatFSResult.FileSystemTypeName]: A property for the file system type name.
//   - [FSStatFSResult.FreeBlocks]: A property for the number of free blocks in the volume.
//   - [FSStatFSResult.SetFreeBlocks]
//   - [FSStatFSResult.FreeBytes]: A property for the amount of free space, in bytes, in the volume.
//   - [FSStatFSResult.SetFreeBytes]
//   - [FSStatFSResult.FreeFiles]: A property for the total number of free file slots in the volume.
//   - [FSStatFSResult.SetFreeFiles]
//   - [FSStatFSResult.IoSize]: A property for the optimal block size with which to perform I/O.
//   - [FSStatFSResult.SetIoSize]
//   - [FSStatFSResult.TotalBlocks]: A property for the volume’s total data block count.
//   - [FSStatFSResult.SetTotalBlocks]
//   - [FSStatFSResult.TotalBytes]: A property for the total size, in bytes, of the volume.
//   - [FSStatFSResult.SetTotalBytes]
//   - [FSStatFSResult.TotalFiles]: A property for the total number of file slots in the volume,
//   - [FSStatFSResult.SetTotalFiles]
//   - [FSStatFSResult.UsedBlocks]: A property for the number of used blocks in the volume.
//   - [FSStatFSResult.SetUsedBlocks]
//   - [FSStatFSResult.UsedBytes]: A property for the amount of used space, in bytes, in the volume.
//   - [FSStatFSResult.SetUsedBytes]
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult
type FSStatFSResult struct {
	objectivec.Object
}

// FSStatFSResultFromID constructs a [FSStatFSResult] from an objc.ID.
//
// A type used to report a volume’s statistics.
func FSStatFSResultFromID(id objc.ID) FSStatFSResult {
	return FSStatFSResult{objectivec.Object{ID: id}}
}

// NOTE: FSStatFSResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSStatFSResult] class.
//
// # Initializers
//
//   - [IFSStatFSResult.InitWithCoder]
//   - [IFSStatFSResult.InitWithFileSystemTypeName]: Creates an statistics result instance, using the given file system type name.
//
// # Instance Properties
//
//   - [IFSStatFSResult.AvailableBlocks]: A property for the number of free blocks available to a non-superuser on the volume.
//   - [IFSStatFSResult.SetAvailableBlocks]
//   - [IFSStatFSResult.AvailableBytes]: A property for the amount of space available to users, in bytes, in the volume.
//   - [IFSStatFSResult.SetAvailableBytes]
//   - [IFSStatFSResult.BlockSize]: A property for the volume’s block size, in bytes.
//   - [IFSStatFSResult.SetBlockSize]
//   - [IFSStatFSResult.FileSystemSubType]: A property for the file system’s subtype or flavor.
//   - [IFSStatFSResult.SetFileSystemSubType]
//   - [IFSStatFSResult.FileSystemTypeName]: A property for the file system type name.
//   - [IFSStatFSResult.FreeBlocks]: A property for the number of free blocks in the volume.
//   - [IFSStatFSResult.SetFreeBlocks]
//   - [IFSStatFSResult.FreeBytes]: A property for the amount of free space, in bytes, in the volume.
//   - [IFSStatFSResult.SetFreeBytes]
//   - [IFSStatFSResult.FreeFiles]: A property for the total number of free file slots in the volume.
//   - [IFSStatFSResult.SetFreeFiles]
//   - [IFSStatFSResult.IoSize]: A property for the optimal block size with which to perform I/O.
//   - [IFSStatFSResult.SetIoSize]
//   - [IFSStatFSResult.TotalBlocks]: A property for the volume’s total data block count.
//   - [IFSStatFSResult.SetTotalBlocks]
//   - [IFSStatFSResult.TotalBytes]: A property for the total size, in bytes, of the volume.
//   - [IFSStatFSResult.SetTotalBytes]
//   - [IFSStatFSResult.TotalFiles]: A property for the total number of file slots in the volume,
//   - [IFSStatFSResult.SetTotalFiles]
//   - [IFSStatFSResult.UsedBlocks]: A property for the number of used blocks in the volume.
//   - [IFSStatFSResult.SetUsedBlocks]
//   - [IFSStatFSResult.UsedBytes]: A property for the amount of used space, in bytes, in the volume.
//   - [IFSStatFSResult.SetUsedBytes]
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult
type IFSStatFSResult interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) FSStatFSResult
	// Creates an statistics result instance, using the given file system type name.
	InitWithFileSystemTypeName(fileSystemTypeName string) FSStatFSResult

	// Topic: Instance Properties

	// A property for the number of free blocks available to a non-superuser on the volume.
	AvailableBlocks() uint64
	SetAvailableBlocks(value uint64)
	// A property for the amount of space available to users, in bytes, in the volume.
	AvailableBytes() uint64
	SetAvailableBytes(value uint64)
	// A property for the volume’s block size, in bytes.
	BlockSize() int
	SetBlockSize(value int)
	// A property for the file system’s subtype or flavor.
	FileSystemSubType() int
	SetFileSystemSubType(value int)
	// A property for the file system type name.
	FileSystemTypeName() string
	// A property for the number of free blocks in the volume.
	FreeBlocks() uint64
	SetFreeBlocks(value uint64)
	// A property for the amount of free space, in bytes, in the volume.
	FreeBytes() uint64
	SetFreeBytes(value uint64)
	// A property for the total number of free file slots in the volume.
	FreeFiles() uint64
	SetFreeFiles(value uint64)
	// A property for the optimal block size with which to perform I/O.
	IoSize() int
	SetIoSize(value int)
	// A property for the volume’s total data block count.
	TotalBlocks() uint64
	SetTotalBlocks(value uint64)
	// A property for the total size, in bytes, of the volume.
	TotalBytes() uint64
	SetTotalBytes(value uint64)
	// A property for the total number of file slots in the volume,
	TotalFiles() uint64
	SetTotalFiles(value uint64)
	// A property for the number of used blocks in the volume.
	UsedBlocks() uint64
	SetUsedBlocks(value uint64)
	// A property for the amount of used space, in bytes, in the volume.
	UsedBytes() uint64
	SetUsedBytes(value uint64)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s FSStatFSResult) Init() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s FSStatFSResult) Autorelease() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSStatFSResult creates a new FSStatFSResult instance.
func NewFSStatFSResult() FSStatFSResult {
	class := getFSStatFSResultClass()
	rv := objc.Send[FSStatFSResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/init(coder:)
func NewStatFSResultWithCoder(coder foundation.INSCoder) FSStatFSResult {
	instance := getFSStatFSResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return FSStatFSResultFromID(rv)
}

// Creates an statistics result instance, using the given file system type
// name.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/init(fileSystemTypeName:)
func NewStatFSResultWithFileSystemTypeName(fileSystemTypeName string) FSStatFSResult {
	instance := getFSStatFSResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileSystemTypeName:"), objc.String(fileSystemTypeName))
	return FSStatFSResultFromID(rv)
}

// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/init(coder:)
func (s FSStatFSResult) InitWithCoder(coder foundation.INSCoder) FSStatFSResult {
	rv := objc.Send[FSStatFSResult](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Creates an statistics result instance, using the given file system type
// name.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/init(fileSystemTypeName:)
func (s FSStatFSResult) InitWithFileSystemTypeName(fileSystemTypeName string) FSStatFSResult {
	rv := objc.Send[FSStatFSResult](s.ID, objc.Sel("initWithFileSystemTypeName:"), objc.String(fileSystemTypeName))
	return rv
}
func (s FSStatFSResult) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A property for the number of free blocks available to a non-superuser on
// the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/availableBlocks
func (s FSStatFSResult) AvailableBlocks() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("availableBlocks"))
	return rv
}
func (s FSStatFSResult) SetAvailableBlocks(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setAvailableBlocks:"), value)
}

// A property for the amount of space available to users, in bytes, in the
// volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/availableBytes
func (s FSStatFSResult) AvailableBytes() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("availableBytes"))
	return rv
}
func (s FSStatFSResult) SetAvailableBytes(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setAvailableBytes:"), value)
}

// A property for the volume’s block size, in bytes.
//
// # Discussion
//
// This value defaults to `4096`. Zero isn’t a valid block size.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/blockSize
func (s FSStatFSResult) BlockSize() int {
	rv := objc.Send[int](s.ID, objc.Sel("blockSize"))
	return rv
}
func (s FSStatFSResult) SetBlockSize(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setBlockSize:"), value)
}

// A property for the file system’s subtype or flavor.
//
// # Discussion
//
// Match this value to the [FSPersonalities]‘s [FSSubType] attribute, if it
// exists within the [EXAppExtensionAttributes] dictionary of the module’s
// `Info.Plist()`.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/fileSystemSubType
func (s FSStatFSResult) FileSystemSubType() int {
	rv := objc.Send[int](s.ID, objc.Sel("fileSystemSubType"))
	return rv
}
func (s FSStatFSResult) SetFileSystemSubType(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setFileSystemSubType:"), value)
}

// A property for the file system type name.
//
// # Discussion
//
// Match this value to the [FSShortName] attribute within the
// [EXAppExtensionAttributes] dictionary of the module’s `Info.Plist()`. The
// maximum allowed length is [MFSTYPENAMELEN], including the terminating [NUL]
// character.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/fileSystemTypeName
func (s FSStatFSResult) FileSystemTypeName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fileSystemTypeName"))
	return foundation.NSStringFromID(rv).String()
}

// A property for the number of free blocks in the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/freeBlocks
func (s FSStatFSResult) FreeBlocks() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("freeBlocks"))
	return rv
}
func (s FSStatFSResult) SetFreeBlocks(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setFreeBlocks:"), value)
}

// A property for the amount of free space, in bytes, in the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/freeBytes
func (s FSStatFSResult) FreeBytes() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("freeBytes"))
	return rv
}
func (s FSStatFSResult) SetFreeBytes(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setFreeBytes:"), value)
}

// A property for the total number of free file slots in the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/freeFiles
func (s FSStatFSResult) FreeFiles() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("freeFiles"))
	return rv
}
func (s FSStatFSResult) SetFreeFiles(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setFreeFiles:"), value)
}

// A property for the optimal block size with which to perform I/O.
//
// # Discussion
//
// For best performance, specify an `ioSize` that’s an even multiple of
// [FSStatFSResult.BlockSize].
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/ioSize
func (s FSStatFSResult) IoSize() int {
	rv := objc.Send[int](s.ID, objc.Sel("ioSize"))
	return rv
}
func (s FSStatFSResult) SetIoSize(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setIoSize:"), value)
}

// A property for the volume’s total data block count.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/totalBlocks
func (s FSStatFSResult) TotalBlocks() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("totalBlocks"))
	return rv
}
func (s FSStatFSResult) SetTotalBlocks(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setTotalBlocks:"), value)
}

// A property for the total size, in bytes, of the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/totalBytes
func (s FSStatFSResult) TotalBytes() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("totalBytes"))
	return rv
}
func (s FSStatFSResult) SetTotalBytes(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setTotalBytes:"), value)
}

// A property for the total number of file slots in the volume,
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/totalFiles
func (s FSStatFSResult) TotalFiles() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("totalFiles"))
	return rv
}
func (s FSStatFSResult) SetTotalFiles(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setTotalFiles:"), value)
}

// A property for the number of used blocks in the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/usedBlocks
func (s FSStatFSResult) UsedBlocks() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("usedBlocks"))
	return rv
}
func (s FSStatFSResult) SetUsedBlocks(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setUsedBlocks:"), value)
}

// A property for the amount of used space, in bytes, in the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSStatFSResult/usedBytes
func (s FSStatFSResult) UsedBytes() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("usedBytes"))
	return rv
}
func (s FSStatFSResult) SetUsedBytes(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setUsedBytes:"), value)
}

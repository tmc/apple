// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"syscall"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSItemAttributes] class.
var (
	_FSItemAttributesClass     FSItemAttributesClass
	_FSItemAttributesClassOnce sync.Once
)

func getFSItemAttributesClass() FSItemAttributesClass {
	_FSItemAttributesClassOnce.Do(func() {
		_FSItemAttributesClass = FSItemAttributesClass{class: objc.GetClass("FSItemAttributes")}
	})
	return _FSItemAttributesClass
}

// GetFSItemAttributesClass returns the class object for FSItemAttributes.
func GetFSItemAttributesClass() FSItemAttributesClass {
	return getFSItemAttributesClass()
}

type FSItemAttributesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSItemAttributesClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSItemAttributesClass) Alloc() FSItemAttributes {
	rv := objc.Send[FSItemAttributes](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// Attributes of an item, such as size, creation and modification times, and
// user and group identifiers.
//
// # Validating and invalidating attributes
//
//   - [FSItemAttributes.IsValid]: Returns a Boolean value that indicates whether the attribute is valid.
//   - [FSItemAttributes.InvalidateAllProperties]: Marks all attributes inactive.
//
// # Working with identifier attributes
//
//   - [FSItemAttributes.FileID]: The item’s file identifier.
//   - [FSItemAttributes.SetFileID]
//   - [FSItemAttributes.ParentID]: The identifier of the item’s parent.
//   - [FSItemAttributes.SetParentID]
//
// # Working with metadata attributes
//
//   - [FSItemAttributes.Type]: The item type, such as a regular file, directory, or symbolic link.
//   - [FSItemAttributes.SetType]
//   - [FSItemAttributes.Mode]: The mode of the item.
//   - [FSItemAttributes.SetMode]
//   - [FSItemAttributes.LinkCount]: The number of hard links to the item.
//   - [FSItemAttributes.SetLinkCount]
//   - [FSItemAttributes.Uid]: The user identifier.
//   - [FSItemAttributes.SetUid]
//   - [FSItemAttributes.Gid]: The group identifier.
//   - [FSItemAttributes.SetGid]
//   - [FSItemAttributes.Flags]: The item’s behavior flags.
//   - [FSItemAttributes.SetFlags]
//   - [FSItemAttributes.Size]: The item’s size.
//   - [FSItemAttributes.SetSize]
//   - [FSItemAttributes.AllocSize]: The item’s allocated size.
//   - [FSItemAttributes.SetAllocSize]
//   - [FSItemAttributes.SupportsLimitedXAttrs]: A Boolean value that indicates whether the item supports a limited set of extended attributes.
//   - [FSItemAttributes.SetSupportsLimitedXAttrs]
//   - [FSItemAttributes.InhibitKernelOffloadedIO]: A Boolean value that indicates whether the file system overrides the per-volume settings for kernel offloaded I/O for a specific file.
//   - [FSItemAttributes.SetInhibitKernelOffloadedIO]
//
// # Working with time attributes
//
//   - [FSItemAttributes.AccessTime]: The item’s last-accessed time.
//   - [FSItemAttributes.SetAccessTime]
//   - [FSItemAttributes.ModifyTime]: The item’s last-modified time.
//   - [FSItemAttributes.SetModifyTime]
//   - [FSItemAttributes.ChangeTime]: The item’s last-changed time.
//   - [FSItemAttributes.SetChangeTime]
//   - [FSItemAttributes.BirthTime]: The item’s creation time.
//   - [FSItemAttributes.SetBirthTime]
//   - [FSItemAttributes.BackupTime]: The item’s last-backup time.
//   - [FSItemAttributes.SetBackupTime]
//   - [FSItemAttributes.AddedTime]: The item’s added time.
//   - [FSItemAttributes.SetAddedTime]
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes
type FSItemAttributes struct {
	objectivec.Object
}

// FSItemAttributesFromID constructs a [FSItemAttributes] from an objc.ID.
//
// Attributes of an item, such as size, creation and modification times, and
// user and group identifiers.
func FSItemAttributesFromID(id objc.ID) FSItemAttributes {
	return FSItemAttributes{objectivec.Object{ID: id}}
}

// NOTE: FSItemAttributes adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSItemAttributes] class.
//
// # Validating and invalidating attributes
//
//   - [IFSItemAttributes.IsValid]: Returns a Boolean value that indicates whether the attribute is valid.
//   - [IFSItemAttributes.InvalidateAllProperties]: Marks all attributes inactive.
//
// # Working with identifier attributes
//
//   - [IFSItemAttributes.FileID]: The item’s file identifier.
//   - [IFSItemAttributes.SetFileID]
//   - [IFSItemAttributes.ParentID]: The identifier of the item’s parent.
//   - [IFSItemAttributes.SetParentID]
//
// # Working with metadata attributes
//
//   - [IFSItemAttributes.Type]: The item type, such as a regular file, directory, or symbolic link.
//   - [IFSItemAttributes.SetType]
//   - [IFSItemAttributes.Mode]: The mode of the item.
//   - [IFSItemAttributes.SetMode]
//   - [IFSItemAttributes.LinkCount]: The number of hard links to the item.
//   - [IFSItemAttributes.SetLinkCount]
//   - [IFSItemAttributes.Uid]: The user identifier.
//   - [IFSItemAttributes.SetUid]
//   - [IFSItemAttributes.Gid]: The group identifier.
//   - [IFSItemAttributes.SetGid]
//   - [IFSItemAttributes.Flags]: The item’s behavior flags.
//   - [IFSItemAttributes.SetFlags]
//   - [IFSItemAttributes.Size]: The item’s size.
//   - [IFSItemAttributes.SetSize]
//   - [IFSItemAttributes.AllocSize]: The item’s allocated size.
//   - [IFSItemAttributes.SetAllocSize]
//   - [IFSItemAttributes.SupportsLimitedXAttrs]: A Boolean value that indicates whether the item supports a limited set of extended attributes.
//   - [IFSItemAttributes.SetSupportsLimitedXAttrs]
//   - [IFSItemAttributes.InhibitKernelOffloadedIO]: A Boolean value that indicates whether the file system overrides the per-volume settings for kernel offloaded I/O for a specific file.
//   - [IFSItemAttributes.SetInhibitKernelOffloadedIO]
//
// # Working with time attributes
//
//   - [IFSItemAttributes.AccessTime]: The item’s last-accessed time.
//   - [IFSItemAttributes.SetAccessTime]
//   - [IFSItemAttributes.ModifyTime]: The item’s last-modified time.
//   - [IFSItemAttributes.SetModifyTime]
//   - [IFSItemAttributes.ChangeTime]: The item’s last-changed time.
//   - [IFSItemAttributes.SetChangeTime]
//   - [IFSItemAttributes.BirthTime]: The item’s creation time.
//   - [IFSItemAttributes.SetBirthTime]
//   - [IFSItemAttributes.BackupTime]: The item’s last-backup time.
//   - [IFSItemAttributes.SetBackupTime]
//   - [IFSItemAttributes.AddedTime]: The item’s added time.
//   - [IFSItemAttributes.SetAddedTime]
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes
type IFSItemAttributes interface {
	objectivec.IObject

	// Topic: Validating and invalidating attributes

	// Returns a Boolean value that indicates whether the attribute is valid.
	IsValid(attribute FSItemAttribute) bool
	// Marks all attributes inactive.
	InvalidateAllProperties()

	// Topic: Working with identifier attributes

	// The item’s file identifier.
	FileID() FSItemID
	SetFileID(value FSItemID)
	// The identifier of the item’s parent.
	ParentID() FSItemID
	SetParentID(value FSItemID)

	// Topic: Working with metadata attributes

	// The item type, such as a regular file, directory, or symbolic link.
	Type() FSItemType
	SetType(value FSItemType)
	// The mode of the item.
	Mode() uint32
	SetMode(value uint32)
	// The number of hard links to the item.
	LinkCount() uint32
	SetLinkCount(value uint32)
	// The user identifier.
	Uid() uint32
	SetUid(value uint32)
	// The group identifier.
	Gid() uint32
	SetGid(value uint32)
	// The item’s behavior flags.
	Flags() uint32
	SetFlags(value uint32)
	// The item’s size.
	Size() uint64
	SetSize(value uint64)
	// The item’s allocated size.
	AllocSize() uint64
	SetAllocSize(value uint64)
	// A Boolean value that indicates whether the item supports a limited set of extended attributes.
	SupportsLimitedXAttrs() bool
	SetSupportsLimitedXAttrs(value bool)
	// A Boolean value that indicates whether the file system overrides the per-volume settings for kernel offloaded I/O for a specific file.
	InhibitKernelOffloadedIO() bool
	SetInhibitKernelOffloadedIO(value bool)

	// Topic: Working with time attributes

	// The item’s last-accessed time.
	AccessTime() syscall.Timespec
	SetAccessTime(value syscall.Timespec)
	// The item’s last-modified time.
	ModifyTime() syscall.Timespec
	SetModifyTime(value syscall.Timespec)
	// The item’s last-changed time.
	ChangeTime() syscall.Timespec
	SetChangeTime(value syscall.Timespec)
	// The item’s creation time.
	BirthTime() syscall.Timespec
	SetBirthTime(value syscall.Timespec)
	// The item’s last-backup time.
	BackupTime() syscall.Timespec
	SetBackupTime(value syscall.Timespec)
	// The item’s added time.
	AddedTime() syscall.Timespec
	SetAddedTime(value syscall.Timespec)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (i FSItemAttributes) Init() FSItemAttributes {
	rv := objc.Send[FSItemAttributes](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i FSItemAttributes) Autorelease() FSItemAttributes {
	rv := objc.Send[FSItemAttributes](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSItemAttributes creates a new FSItemAttributes instance.
func NewFSItemAttributes() FSItemAttributes {
	class := getFSItemAttributesClass()
	rv := objc.Send[FSItemAttributes](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a Boolean value that indicates whether the attribute is valid.
//
// # Discussion
//
// If the value returned by this method is [YES] (Objective-C) or `true`
// (Swift), a caller can safely use the given attribute.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/isValid(_:)
func (i FSItemAttributes) IsValid(attribute FSItemAttribute) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isValid:"), attribute)
	return rv
}

// Marks all attributes inactive.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/invalidateAllProperties()
func (i FSItemAttributes) InvalidateAllProperties() {
	objc.Send[objc.ID](i.ID, objc.Sel("invalidateAllProperties"))
}
func (i FSItemAttributes) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The item’s file identifier.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/fileID
func (i FSItemAttributes) FileID() FSItemID {
	rv := objc.Send[FSItemID](i.ID, objc.Sel("fileID"))
	return FSItemID(rv)
}
func (i FSItemAttributes) SetFileID(value FSItemID) {
	objc.Send[struct{}](i.ID, objc.Sel("setFileID:"), value)
}

// The identifier of the item’s parent.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/parentID
func (i FSItemAttributes) ParentID() FSItemID {
	rv := objc.Send[FSItemID](i.ID, objc.Sel("parentID"))
	return FSItemID(rv)
}
func (i FSItemAttributes) SetParentID(value FSItemID) {
	objc.Send[struct{}](i.ID, objc.Sel("setParentID:"), value)
}

// The item type, such as a regular file, directory, or symbolic link.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/type
func (i FSItemAttributes) Type() FSItemType {
	rv := objc.Send[FSItemType](i.ID, objc.Sel("type"))
	return FSItemType(rv)
}
func (i FSItemAttributes) SetType(value FSItemType) {
	objc.Send[struct{}](i.ID, objc.Sel("setType:"), value)
}

// The mode of the item.
//
// # Discussion
//
// The mode is often used for `setuid`, `setgid`, and `sticky` bits.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/mode
func (i FSItemAttributes) Mode() uint32 {
	rv := objc.Send[uint32](i.ID, objc.Sel("mode"))
	return rv
}
func (i FSItemAttributes) SetMode(value uint32) {
	objc.Send[struct{}](i.ID, objc.Sel("setMode:"), value)
}

// The number of hard links to the item.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/linkCount
func (i FSItemAttributes) LinkCount() uint32 {
	rv := objc.Send[uint32](i.ID, objc.Sel("linkCount"))
	return rv
}
func (i FSItemAttributes) SetLinkCount(value uint32) {
	objc.Send[struct{}](i.ID, objc.Sel("setLinkCount:"), value)
}

// The user identifier.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/uid
func (i FSItemAttributes) Uid() uint32 {
	rv := objc.Send[uint32](i.ID, objc.Sel("uid"))
	return rv
}
func (i FSItemAttributes) SetUid(value uint32) {
	objc.Send[struct{}](i.ID, objc.Sel("setUid:"), value)
}

// The group identifier.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/gid
func (i FSItemAttributes) Gid() uint32 {
	rv := objc.Send[uint32](i.ID, objc.Sel("gid"))
	return rv
}
func (i FSItemAttributes) SetGid(value uint32) {
	objc.Send[struct{}](i.ID, objc.Sel("setGid:"), value)
}

// The item’s behavior flags.
//
// # Discussion
//
// See `st_flags` in `stat.H()` for flag definitions.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/flags
func (i FSItemAttributes) Flags() uint32 {
	rv := objc.Send[uint32](i.ID, objc.Sel("flags"))
	return rv
}
func (i FSItemAttributes) SetFlags(value uint32) {
	objc.Send[struct{}](i.ID, objc.Sel("setFlags:"), value)
}

// The item’s size.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/size
func (i FSItemAttributes) Size() uint64 {
	rv := objc.Send[uint64](i.ID, objc.Sel("size"))
	return rv
}
func (i FSItemAttributes) SetSize(value uint64) {
	objc.Send[struct{}](i.ID, objc.Sel("setSize:"), value)
}

// The item’s allocated size.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/allocSize
func (i FSItemAttributes) AllocSize() uint64 {
	rv := objc.Send[uint64](i.ID, objc.Sel("allocSize"))
	return rv
}
func (i FSItemAttributes) SetAllocSize(value uint64) {
	objc.Send[struct{}](i.ID, objc.Sel("setAllocSize:"), value)
}

// A Boolean value that indicates whether the item supports a limited set of
// extended attributes.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/supportsLimitedXAttrs
func (i FSItemAttributes) SupportsLimitedXAttrs() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("supportsLimitedXAttrs"))
	return rv
}
func (i FSItemAttributes) SetSupportsLimitedXAttrs(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setSupportsLimitedXAttrs:"), value)
}

// A Boolean value that indicates whether the file system overrides the
// per-volume settings for kernel offloaded I/O for a specific file.
//
// # Discussion
//
// This property has no meaning if the volume doesn’t conform to
// [FSVolumeKernelOffloadedIOOperations].
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/inhibitKernelOffloadedIO
func (i FSItemAttributes) InhibitKernelOffloadedIO() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("inhibitKernelOffloadedIO"))
	return rv
}
func (i FSItemAttributes) SetInhibitKernelOffloadedIO(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setInhibitKernelOffloadedIO:"), value)
}

// The item’s last-accessed time.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/accessTime
func (i FSItemAttributes) AccessTime() syscall.Timespec {
	rv := objc.Send[syscall.Timespec](i.ID, objc.Sel("accessTime"))
	return syscall.Timespec(rv)
}
func (i FSItemAttributes) SetAccessTime(value syscall.Timespec) {
	objc.Send[struct{}](i.ID, objc.Sel("setAccessTime:"), value)
}

// The item’s last-modified time.
//
// # Discussion
//
// This property represents `mtime`, the last time the item’s contents
// changed.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/modifyTime
func (i FSItemAttributes) ModifyTime() syscall.Timespec {
	rv := objc.Send[syscall.Timespec](i.ID, objc.Sel("modifyTime"))
	return syscall.Timespec(rv)
}
func (i FSItemAttributes) SetModifyTime(value syscall.Timespec) {
	objc.Send[struct{}](i.ID, objc.Sel("setModifyTime:"), value)
}

// The item’s last-changed time.
//
// # Discussion
//
// This property represents `ctime`, the last time the item’s metadata
// changed.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/changeTime
func (i FSItemAttributes) ChangeTime() syscall.Timespec {
	rv := objc.Send[syscall.Timespec](i.ID, objc.Sel("changeTime"))
	return syscall.Timespec(rv)
}
func (i FSItemAttributes) SetChangeTime(value syscall.Timespec) {
	objc.Send[struct{}](i.ID, objc.Sel("setChangeTime:"), value)
}

// The item’s creation time.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/birthTime
func (i FSItemAttributes) BirthTime() syscall.Timespec {
	rv := objc.Send[syscall.Timespec](i.ID, objc.Sel("birthTime"))
	return syscall.Timespec(rv)
}
func (i FSItemAttributes) SetBirthTime(value syscall.Timespec) {
	objc.Send[struct{}](i.ID, objc.Sel("setBirthTime:"), value)
}

// The item’s last-backup time.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/backupTime
func (i FSItemAttributes) BackupTime() syscall.Timespec {
	rv := objc.Send[syscall.Timespec](i.ID, objc.Sel("backupTime"))
	return syscall.Timespec(rv)
}
func (i FSItemAttributes) SetBackupTime(value syscall.Timespec) {
	objc.Send[struct{}](i.ID, objc.Sel("setBackupTime:"), value)
}

// The item’s added time.
//
// # Discussion
//
// This property represents the time the file system added the item to its
// parent directory.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem/Attributes/addedTime
func (i FSItemAttributes) AddedTime() syscall.Timespec {
	rv := objc.Send[syscall.Timespec](i.ID, objc.Sel("addedTime"))
	return syscall.Timespec(rv)
}
func (i FSItemAttributes) SetAddedTime(value syscall.Timespec) {
	objc.Send[struct{}](i.ID, objc.Sel("setAddedTime:"), value)
}

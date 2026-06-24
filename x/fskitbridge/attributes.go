package fskitbridge

import (
	"syscall"

	"github.com/tmc/apple/fskit"
)

// ItemAttributesBuilder builds FSItemAttributes with the common fields used
// by FSKit file systems.
type ItemAttributesBuilder struct {
	attrs fskit.FSItemAttributes
}

// NewItemAttributes returns a builder for FSItemAttributes.
func NewItemAttributes() ItemAttributesBuilder {
	return ItemAttributesBuilder{attrs: fskit.NewFSItemAttributes()}
}

// ID sets the file ID.
func (b ItemAttributesBuilder) ID(id fskit.FSItemID) ItemAttributesBuilder {
	b.attrs.SetFileID(id)
	return b
}

// ParentID sets the parent directory's file ID.
func (b ItemAttributesBuilder) ParentID(id fskit.FSItemID) ItemAttributesBuilder {
	b.attrs.SetParentID(id)
	return b
}

// Type sets the item type.
func (b ItemAttributesBuilder) Type(t fskit.FSItemType) ItemAttributesBuilder {
	b.attrs.SetType(t)
	return b
}

// Mode sets the POSIX mode bits.
func (b ItemAttributesBuilder) Mode(mode uint32) ItemAttributesBuilder {
	b.attrs.SetMode(mode)
	return b
}

// LinkCount sets the hard link count.
func (b ItemAttributesBuilder) LinkCount(n uint32) ItemAttributesBuilder {
	b.attrs.SetLinkCount(n)
	return b
}

// UID sets the owning user ID.
func (b ItemAttributesBuilder) UID(uid uint32) ItemAttributesBuilder {
	b.attrs.SetUid(uid)
	return b
}

// GID sets the owning group ID.
func (b ItemAttributesBuilder) GID(gid uint32) ItemAttributesBuilder {
	b.attrs.SetGid(gid)
	return b
}

// Flags sets the BSD file flags.
func (b ItemAttributesBuilder) Flags(flags uint32) ItemAttributesBuilder {
	b.attrs.SetFlags(flags)
	return b
}

// Size sets the file size in bytes.
func (b ItemAttributesBuilder) Size(size uint64) ItemAttributesBuilder {
	b.attrs.SetSize(size)
	return b
}

// AllocSize sets the allocated size in bytes.
func (b ItemAttributesBuilder) AllocSize(size uint64) ItemAttributesBuilder {
	b.attrs.SetAllocSize(size)
	return b
}

// AccessTime sets the access time.
func (b ItemAttributesBuilder) AccessTime(ts syscall.Timespec) ItemAttributesBuilder {
	b.attrs.SetAccessTime(ts)
	return b
}

// ModifyTime sets the content modification time.
func (b ItemAttributesBuilder) ModifyTime(ts syscall.Timespec) ItemAttributesBuilder {
	b.attrs.SetModifyTime(ts)
	return b
}

// ChangeTime sets the attribute change time.
func (b ItemAttributesBuilder) ChangeTime(ts syscall.Timespec) ItemAttributesBuilder {
	b.attrs.SetChangeTime(ts)
	return b
}

// BirthTime sets the creation time.
func (b ItemAttributesBuilder) BirthTime(ts syscall.Timespec) ItemAttributesBuilder {
	b.attrs.SetBirthTime(ts)
	return b
}

// Build returns the assembled attributes.
func (b ItemAttributesBuilder) Build() fskit.FSItemAttributes {
	return b.attrs
}

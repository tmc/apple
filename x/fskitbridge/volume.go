package fskitbridge

import (
	"syscall"

	"github.com/tmc/apple/fskit"
)

// An Item identifies a file system object. Implementations define their own
// item values; the Server associates each value with the FSItem object it
// hands to FSKit and passes the value back to later operations.
type Item any

// A Volume implements the core, read-only operations of an FSKit volume.
// The Server adapts a Volume to the FSVolumeOperations protocol and answers
// the mutating operations with EROFS unless the Volume also implements
// [MutableVolume].
//
// Optional interfaces extend a Volume with additional FSKit operations:
// [MutableVolume], [SymlinkVolume], [LinkVolume], [XattrVolume],
// [PreallocateVolume], [OpenCloseVolume], [AccessCheckVolume],
// [CapabilitiesVolume], [StatisticsVolume], and [PathConfVolume].
//
// Operations report failure by returning an error. An error that is or
// wraps a syscall.Errno reports that errno to FSKit; fs.ErrNotExist,
// fs.ErrExist, fs.ErrPermission, fs.ErrInvalid, and errors.ErrUnsupported
// report ENOENT, EEXIST, EACCES, EINVAL, and ENOTSUP; any other error
// reports EIO. In xattr operations fs.ErrNotExist reports ENOATTR.
type Volume interface {
	// VolumeName returns the volume's display name.
	VolumeName() string

	// Root returns the item for the root directory.
	Root() (Item, error)

	// Lookup resolves name within the directory dir.
	Lookup(dir Item, name string) (Item, error)

	// Reclaim releases any state held for item. The system is done with
	// the item; it may be reclaimed more than once.
	Reclaim(item Item)

	// Attributes returns the item's attributes, including its file ID.
	Attributes(item Item) (fskit.FSItemAttributes, error)

	// ReadDir lists the directory dir.
	ReadDir(dir Item) ([]DirEntry, error)

	// Read reads from file at offset into buf, returning the number of
	// bytes read. Reading past the end of the file returns 0, and io.EOF
	// is not an error.
	Read(file Item, offset int64, buf []byte) (int, error)
}

// A DirEntry is one entry of a directory listing.
type DirEntry struct {
	Name       string
	Type       fskit.FSItemType
	Attributes fskit.FSItemAttributes // must include the entry's file ID
}

// SetAttributes holds the attributes of a set-attributes request.
// A nil field was not requested.
//
// The Server reports back to FSKit only the fields a Volume returns from
// SetAttributes, so a Volume that cannot change ownership should omit UID
// and GID from what it returns rather than failing the whole request.
type SetAttributes struct {
	Mode       *uint32
	UID        *uint32
	GID        *uint32
	Flags      *uint32
	Size       *uint64
	AccessTime *syscall.Timespec
	ModifyTime *syscall.Timespec
}

// A MutableVolume is a Volume that supports mutation. Without it the
// Server reports EROFS for all mutating operations.
type MutableVolume interface {
	Volume

	// Create creates a file (FSItemTypeFile) or directory
	// (FSItemTypeDirectory) named name in dir.
	Create(dir Item, name string, typ fskit.FSItemType, mode uint32) (Item, error)

	// Remove removes the item named name from dir.
	Remove(dir Item, name string, item Item) error

	// Rename moves item, named srcName in srcDir, to dstName in dstDir.
	// If over is non-nil the destination exists and is replaced; the
	// Server reclaims over after a successful rename.
	Rename(item Item, srcDir Item, srcName string, dstDir Item, dstName string, over Item) error

	// SetAttributes applies the requested attribute changes to item and
	// returns the subset it actually applied, as the fields left non-nil.
	// A Volume that applied everything returns set unchanged.
	//
	// The Server reports the returned subset to FSKit as the consumed
	// attributes, and FSKit takes an unconsumed attribute as one the file
	// system does not support: that is how a caller learns a chown did not
	// happen instead of believing a silent success. Returning the applied
	// set rather than clearing an in-out parameter keeps the mistake of
	// forgetting to report cheap -- the zero value claims nothing, so an
	// implementer who gets it wrong under-claims.
	SetAttributes(item Item, set SetAttributes) (applied SetAttributes, err error)

	// Write writes data to file at offset, returning the number of bytes
	// written.
	Write(file Item, offset int64, data []byte) (int, error)
}

// A SymlinkVolume is a Volume with symbolic links.
type SymlinkVolume interface {
	Volume

	// Readlink returns the target of the symbolic link item.
	Readlink(item Item) (string, error)

	// Symlink creates a symbolic link named name to target in dir.
	Symlink(dir Item, name, target string) (Item, error)
}

// A LinkVolume is a Volume with hard links.
type LinkVolume interface {
	Volume

	// Link links item under a new name in dir.
	Link(item Item, dir Item, name string) error
}

// An XattrVolume is a Volume with extended attributes. The Server
// implements the FSKit set-xattr policies (must-create, must-replace,
// delete) in terms of these methods.
type XattrVolume interface {
	Volume

	// GetXattr returns the value of the extended attribute name on item.
	GetXattr(item Item, name string) ([]byte, error)

	// SetXattr sets the extended attribute name on item to data.
	SetXattr(item Item, name string, data []byte) error

	// RemoveXattr removes the extended attribute name from item.
	RemoveXattr(item Item, name string) error

	// ListXattr lists the extended attribute names on item.
	ListXattr(item Item) ([]string, error)
}

// A PreallocateVolume is a Volume that supports preallocating space.
type PreallocateVolume interface {
	Volume

	// Preallocate reserves length bytes at offset in file, returning the
	// number of bytes reserved.
	Preallocate(file Item, offset int64, length uint64) (uint64, error)
}

// An OpenCloseVolume is a Volume that tracks open and close of its items.
type OpenCloseVolume interface {
	Volume

	// Open notes that item was opened with the given modes.
	Open(item Item, modes fskit.FSVolumeOpenModes) error

	// Close notes that item was closed, keeping the given modes.
	Close(item Item, modes fskit.FSVolumeOpenModes) error
}

// An AccessCheckVolume is a Volume that answers access checks itself.
// Without it the kernel checks access against the item mode bits.
type AccessCheckVolume interface {
	Volume

	// CheckAccess reports whether the requested access to item is allowed.
	CheckAccess(item Item, access fskit.FSAccessMask) (bool, error)
}

// A CapabilitiesVolume is a Volume that reports its own supported
// capabilities. Without it the Server reports 64-bit object IDs, hidden
// files, case sensitivity, and symbolic or hard link support according to
// the [SymlinkVolume] and [LinkVolume] interfaces.
type CapabilitiesVolume interface {
	Volume

	// SupportedCapabilities returns the volume's capabilities.
	SupportedCapabilities() fskit.FSVolumeSupportedCapabilities
}

// A StatisticsVolume is a Volume that reports its own statfs statistics.
// Without it the Server reports single-block placeholder statistics under
// the volume name.
type StatisticsVolume interface {
	Volume

	// Statistics returns the volume's statfs statistics.
	Statistics() fskit.FSStatFSResult
}

// A PathConfVolume is a Volume that reports its own path configuration
// limits. Without it the Server reports [DefaultPathConf].
type PathConfVolume interface {
	Volume

	// PathConf returns the volume's path configuration limits.
	PathConf() PathConf
}

// PathConf holds the path configuration limits and behavior the Server
// reports to FSKit. The Server reports the values as given; most
// implementations start from [DefaultPathConf].
type PathConf struct {
	MaximumLinkCount          int
	MaximumNameLength         int
	MaximumFileSize           uint64
	MaximumXattrSize          int
	RestrictsOwnershipChanges bool
	TruncatesLongNames        bool

	// OpenUnlinkEmulation asks FSKit to emulate open-unlink semantics,
	// for volumes that cannot keep unlinked-but-open files alive.
	OpenUnlinkEmulation bool
}

// DefaultPathConf returns the path configuration the Server reports for
// volumes that do not implement [PathConfVolume]: no hard links, 255-byte
// names, 63-bit file sizes, 64 KiB extended attributes, and open-unlink
// emulation.
func DefaultPathConf() PathConf {
	return PathConf{
		MaximumLinkCount:    1,
		MaximumNameLength:   255,
		MaximumFileSize:     1<<63 - 1,
		MaximumXattrSize:    64 << 10,
		OpenUnlinkEmulation: true,
	}
}

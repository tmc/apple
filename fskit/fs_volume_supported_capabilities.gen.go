// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSVolumeSupportedCapabilities] class.
var (
	_FSVolumeSupportedCapabilitiesClass     FSVolumeSupportedCapabilitiesClass
	_FSVolumeSupportedCapabilitiesClassOnce sync.Once
)

func getFSVolumeSupportedCapabilitiesClass() FSVolumeSupportedCapabilitiesClass {
	_FSVolumeSupportedCapabilitiesClassOnce.Do(func() {
		_FSVolumeSupportedCapabilitiesClass = FSVolumeSupportedCapabilitiesClass{class: objc.GetClass("FSVolumeSupportedCapabilities")}
	})
	return _FSVolumeSupportedCapabilitiesClass
}

// GetFSVolumeSupportedCapabilitiesClass returns the class object for FSVolumeSupportedCapabilities.
func GetFSVolumeSupportedCapabilitiesClass() FSVolumeSupportedCapabilitiesClass {
	return getFSVolumeSupportedCapabilitiesClass()
}

type FSVolumeSupportedCapabilitiesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSVolumeSupportedCapabilitiesClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSVolumeSupportedCapabilitiesClass) Alloc() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A type that represents capabillities supported by a volume, such as hard
// and symbolic links, journaling, and large file sizes.
//
// # Declaring identifier capabilities
//
//   - [FSVolumeSupportedCapabilities.SupportsPersistentObjectIDs]: A Boolean property that indicates whether the volume supports persistent object identifiers and can look up file system objects by their IDs.
//   - [FSVolumeSupportedCapabilities.SetSupportsPersistentObjectIDs]
//   - [FSVolumeSupportedCapabilities.Supports64BitObjectIDs]: A Boolean property that indicates whether the volume supports 64-bit object IDs.
//   - [FSVolumeSupportedCapabilities.SetSupports64BitObjectIDs]
//   - [FSVolumeSupportedCapabilities.SupportsDocumentID]: A Boolean property that indicates whether the volume supports document IDs for document revisions.
//   - [FSVolumeSupportedCapabilities.SetSupportsDocumentID]
//
// # Declaring linking capabilities
//
//   - [FSVolumeSupportedCapabilities.SupportsSymbolicLinks]: A Boolean property that indicates whether the volume supports symbolic links.
//   - [FSVolumeSupportedCapabilities.SetSupportsSymbolicLinks]
//   - [FSVolumeSupportedCapabilities.SupportsHardLinks]: A Boolean property that indicates whether the volume supports hard links.
//   - [FSVolumeSupportedCapabilities.SetSupportsHardLinks]
//
// # Declaring journaling capabilities
//
//   - [FSVolumeSupportedCapabilities.SupportsJournal]: A Boolean property that indicates whether the volume supports a journal used to speed recovery in case of unplanned restart, such as a power outage or crash.
//   - [FSVolumeSupportedCapabilities.SetSupportsJournal]
//   - [FSVolumeSupportedCapabilities.SupportsActiveJournal]: A Boolean property that indicates whether the volume currently uses a journal for speeding recovery after an unplanned shutdown.
//   - [FSVolumeSupportedCapabilities.SetSupportsActiveJournal]
//
// # Declaring root capabilites
//
//   - [FSVolumeSupportedCapabilities.DoesNotSupportRootTimes]: A Boolan property that indicates the volume doesn’t store reliable times for the root directory.
//   - [FSVolumeSupportedCapabilities.SetDoesNotSupportRootTimes]
//
// # Declaring file capabilities
//
//   - [FSVolumeSupportedCapabilities.SupportsSparseFiles]: A Boolean property that indicates whether the volume supports sparse files.
//   - [FSVolumeSupportedCapabilities.SetSupportsSparseFiles]
//   - [FSVolumeSupportedCapabilities.SupportsZeroRuns]: A Boolean property that indicates whether the volume supports zero runs
//   - [FSVolumeSupportedCapabilities.SetSupportsZeroRuns]
//   - [FSVolumeSupportedCapabilities.SupportsFastStatFS]: A Boolean property that indicates whether the volume supports fast results when fetching file system statistics.
//   - [FSVolumeSupportedCapabilities.SetSupportsFastStatFS]
//   - [FSVolumeSupportedCapabilities.Supports2TBFiles]: A Boolean property that indicates whether the volume supports file sizes larger than 4GB, and potentially up to 2TB.
//   - [FSVolumeSupportedCapabilities.SetSupports2TBFiles]
//   - [FSVolumeSupportedCapabilities.SupportsOpenDenyModes]: A Boolean property that indicates whether the volume supports open deny modes.
//   - [FSVolumeSupportedCapabilities.SetSupportsOpenDenyModes]
//   - [FSVolumeSupportedCapabilities.SupportsHiddenFiles]: A Boolean property that indicates whether the volume supports hidden files.
//   - [FSVolumeSupportedCapabilities.SetSupportsHiddenFiles]
//   - [FSVolumeSupportedCapabilities.DoesNotSupportImmutableFiles]: A Boolean property that indicates the volume doesn’t support immutable files.
//   - [FSVolumeSupportedCapabilities.SetDoesNotSupportImmutableFiles]
//   - [FSVolumeSupportedCapabilities.DoesNotSupportSettingFilePermissions]: A Boolean property that indicates the volume doesn’t set file permissions.
//   - [FSVolumeSupportedCapabilities.SetDoesNotSupportSettingFilePermissions]
//
// # Declaring volume capabilities
//
//   - [FSVolumeSupportedCapabilities.SupportsSharedSpace]: A Boolean property that indicates whether the volume supports multiple logical file systems that share space in a single “partition.”
//   - [FSVolumeSupportedCapabilities.SetSupportsSharedSpace]
//   - [FSVolumeSupportedCapabilities.SupportsVolumeGroups]: A Boolean property that indicates whether the volume supports volume groups.
//   - [FSVolumeSupportedCapabilities.SetSupportsVolumeGroups]
//   - [FSVolumeSupportedCapabilities.DoesNotSupportVolumeSizes]: A Boolean property that indicates the volume doesn’t support certain volume size reports.
//   - [FSVolumeSupportedCapabilities.SetDoesNotSupportVolumeSizes]
//
// # Working with case sensitivity
//
//   - [FSVolumeSupportedCapabilities.CaseFormat]: A value that indicates the volume’s support for case sensitivity.
//   - [FSVolumeSupportedCapabilities.SetCaseFormat]
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities
type FSVolumeSupportedCapabilities struct {
	objectivec.Object
}

// FSVolumeSupportedCapabilitiesFromID constructs a [FSVolumeSupportedCapabilities] from an objc.ID.
//
// A type that represents capabillities supported by a volume, such as hard
// and symbolic links, journaling, and large file sizes.
func FSVolumeSupportedCapabilitiesFromID(id objc.ID) FSVolumeSupportedCapabilities {
	return FSVolumeSupportedCapabilities{objectivec.Object{ID: id}}
}

// NOTE: FSVolumeSupportedCapabilities adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSVolumeSupportedCapabilities] class.
//
// # Declaring identifier capabilities
//
//   - [IFSVolumeSupportedCapabilities.SupportsPersistentObjectIDs]: A Boolean property that indicates whether the volume supports persistent object identifiers and can look up file system objects by their IDs.
//   - [IFSVolumeSupportedCapabilities.SetSupportsPersistentObjectIDs]
//   - [IFSVolumeSupportedCapabilities.Supports64BitObjectIDs]: A Boolean property that indicates whether the volume supports 64-bit object IDs.
//   - [IFSVolumeSupportedCapabilities.SetSupports64BitObjectIDs]
//   - [IFSVolumeSupportedCapabilities.SupportsDocumentID]: A Boolean property that indicates whether the volume supports document IDs for document revisions.
//   - [IFSVolumeSupportedCapabilities.SetSupportsDocumentID]
//
// # Declaring linking capabilities
//
//   - [IFSVolumeSupportedCapabilities.SupportsSymbolicLinks]: A Boolean property that indicates whether the volume supports symbolic links.
//   - [IFSVolumeSupportedCapabilities.SetSupportsSymbolicLinks]
//   - [IFSVolumeSupportedCapabilities.SupportsHardLinks]: A Boolean property that indicates whether the volume supports hard links.
//   - [IFSVolumeSupportedCapabilities.SetSupportsHardLinks]
//
// # Declaring journaling capabilities
//
//   - [IFSVolumeSupportedCapabilities.SupportsJournal]: A Boolean property that indicates whether the volume supports a journal used to speed recovery in case of unplanned restart, such as a power outage or crash.
//   - [IFSVolumeSupportedCapabilities.SetSupportsJournal]
//   - [IFSVolumeSupportedCapabilities.SupportsActiveJournal]: A Boolean property that indicates whether the volume currently uses a journal for speeding recovery after an unplanned shutdown.
//   - [IFSVolumeSupportedCapabilities.SetSupportsActiveJournal]
//
// # Declaring root capabilites
//
//   - [IFSVolumeSupportedCapabilities.DoesNotSupportRootTimes]: A Boolan property that indicates the volume doesn’t store reliable times for the root directory.
//   - [IFSVolumeSupportedCapabilities.SetDoesNotSupportRootTimes]
//
// # Declaring file capabilities
//
//   - [IFSVolumeSupportedCapabilities.SupportsSparseFiles]: A Boolean property that indicates whether the volume supports sparse files.
//   - [IFSVolumeSupportedCapabilities.SetSupportsSparseFiles]
//   - [IFSVolumeSupportedCapabilities.SupportsZeroRuns]: A Boolean property that indicates whether the volume supports zero runs
//   - [IFSVolumeSupportedCapabilities.SetSupportsZeroRuns]
//   - [IFSVolumeSupportedCapabilities.SupportsFastStatFS]: A Boolean property that indicates whether the volume supports fast results when fetching file system statistics.
//   - [IFSVolumeSupportedCapabilities.SetSupportsFastStatFS]
//   - [IFSVolumeSupportedCapabilities.Supports2TBFiles]: A Boolean property that indicates whether the volume supports file sizes larger than 4GB, and potentially up to 2TB.
//   - [IFSVolumeSupportedCapabilities.SetSupports2TBFiles]
//   - [IFSVolumeSupportedCapabilities.SupportsOpenDenyModes]: A Boolean property that indicates whether the volume supports open deny modes.
//   - [IFSVolumeSupportedCapabilities.SetSupportsOpenDenyModes]
//   - [IFSVolumeSupportedCapabilities.SupportsHiddenFiles]: A Boolean property that indicates whether the volume supports hidden files.
//   - [IFSVolumeSupportedCapabilities.SetSupportsHiddenFiles]
//   - [IFSVolumeSupportedCapabilities.DoesNotSupportImmutableFiles]: A Boolean property that indicates the volume doesn’t support immutable files.
//   - [IFSVolumeSupportedCapabilities.SetDoesNotSupportImmutableFiles]
//   - [IFSVolumeSupportedCapabilities.DoesNotSupportSettingFilePermissions]: A Boolean property that indicates the volume doesn’t set file permissions.
//   - [IFSVolumeSupportedCapabilities.SetDoesNotSupportSettingFilePermissions]
//
// # Declaring volume capabilities
//
//   - [IFSVolumeSupportedCapabilities.SupportsSharedSpace]: A Boolean property that indicates whether the volume supports multiple logical file systems that share space in a single “partition.”
//   - [IFSVolumeSupportedCapabilities.SetSupportsSharedSpace]
//   - [IFSVolumeSupportedCapabilities.SupportsVolumeGroups]: A Boolean property that indicates whether the volume supports volume groups.
//   - [IFSVolumeSupportedCapabilities.SetSupportsVolumeGroups]
//   - [IFSVolumeSupportedCapabilities.DoesNotSupportVolumeSizes]: A Boolean property that indicates the volume doesn’t support certain volume size reports.
//   - [IFSVolumeSupportedCapabilities.SetDoesNotSupportVolumeSizes]
//
// # Working with case sensitivity
//
//   - [IFSVolumeSupportedCapabilities.CaseFormat]: A value that indicates the volume’s support for case sensitivity.
//   - [IFSVolumeSupportedCapabilities.SetCaseFormat]
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities
type IFSVolumeSupportedCapabilities interface {
	objectivec.IObject

	// Topic: Declaring identifier capabilities

	// A Boolean property that indicates whether the volume supports persistent object identifiers and can look up file system objects by their IDs.
	SupportsPersistentObjectIDs() bool
	SetSupportsPersistentObjectIDs(value bool)
	// A Boolean property that indicates whether the volume supports 64-bit object IDs.
	Supports64BitObjectIDs() bool
	SetSupports64BitObjectIDs(value bool)
	// A Boolean property that indicates whether the volume supports document IDs for document revisions.
	SupportsDocumentID() bool
	SetSupportsDocumentID(value bool)

	// Topic: Declaring linking capabilities

	// A Boolean property that indicates whether the volume supports symbolic links.
	SupportsSymbolicLinks() bool
	SetSupportsSymbolicLinks(value bool)
	// A Boolean property that indicates whether the volume supports hard links.
	SupportsHardLinks() bool
	SetSupportsHardLinks(value bool)

	// Topic: Declaring journaling capabilities

	// A Boolean property that indicates whether the volume supports a journal used to speed recovery in case of unplanned restart, such as a power outage or crash.
	SupportsJournal() bool
	SetSupportsJournal(value bool)
	// A Boolean property that indicates whether the volume currently uses a journal for speeding recovery after an unplanned shutdown.
	SupportsActiveJournal() bool
	SetSupportsActiveJournal(value bool)

	// Topic: Declaring root capabilites

	// A Boolan property that indicates the volume doesn’t store reliable times for the root directory.
	DoesNotSupportRootTimes() bool
	SetDoesNotSupportRootTimes(value bool)

	// Topic: Declaring file capabilities

	// A Boolean property that indicates whether the volume supports sparse files.
	SupportsSparseFiles() bool
	SetSupportsSparseFiles(value bool)
	// A Boolean property that indicates whether the volume supports zero runs
	SupportsZeroRuns() bool
	SetSupportsZeroRuns(value bool)
	// A Boolean property that indicates whether the volume supports fast results when fetching file system statistics.
	SupportsFastStatFS() bool
	SetSupportsFastStatFS(value bool)
	// A Boolean property that indicates whether the volume supports file sizes larger than 4GB, and potentially up to 2TB.
	Supports2TBFiles() bool
	SetSupports2TBFiles(value bool)
	// A Boolean property that indicates whether the volume supports open deny modes.
	SupportsOpenDenyModes() bool
	SetSupportsOpenDenyModes(value bool)
	// A Boolean property that indicates whether the volume supports hidden files.
	SupportsHiddenFiles() bool
	SetSupportsHiddenFiles(value bool)
	// A Boolean property that indicates the volume doesn’t support immutable files.
	DoesNotSupportImmutableFiles() bool
	SetDoesNotSupportImmutableFiles(value bool)
	// A Boolean property that indicates the volume doesn’t set file permissions.
	DoesNotSupportSettingFilePermissions() bool
	SetDoesNotSupportSettingFilePermissions(value bool)

	// Topic: Declaring volume capabilities

	// A Boolean property that indicates whether the volume supports multiple logical file systems that share space in a single “partition.”
	SupportsSharedSpace() bool
	SetSupportsSharedSpace(value bool)
	// A Boolean property that indicates whether the volume supports volume groups.
	SupportsVolumeGroups() bool
	SetSupportsVolumeGroups(value bool)
	// A Boolean property that indicates the volume doesn’t support certain volume size reports.
	DoesNotSupportVolumeSizes() bool
	SetDoesNotSupportVolumeSizes(value bool)

	// Topic: Working with case sensitivity

	// A value that indicates the volume’s support for case sensitivity.
	CaseFormat() FSVolumeCaseFormat
	SetCaseFormat(value FSVolumeCaseFormat)

	// A property that provides the supported capabilities of the volume.
	SupportedVolumeCapabilities() IFSVolumeSupportedCapabilities
	SetSupportedVolumeCapabilities(value IFSVolumeSupportedCapabilities)
	// A property that provides up-to-date statistics of the volume.
	VolumeStatistics() IFSStatFSResult
	SetVolumeStatistics(value IFSStatFSResult)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (v FSVolumeSupportedCapabilities) Init() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v FSVolumeSupportedCapabilities) Autorelease() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSVolumeSupportedCapabilities creates a new FSVolumeSupportedCapabilities instance.
func NewFSVolumeSupportedCapabilities() FSVolumeSupportedCapabilities {
	class := getFSVolumeSupportedCapabilitiesClass()
	rv := objc.Send[FSVolumeSupportedCapabilities](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v FSVolumeSupportedCapabilities) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean property that indicates whether the volume supports persistent
// object identifiers and can look up file system objects by their IDs.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsPersistentObjectIDs
func (v FSVolumeSupportedCapabilities) SupportsPersistentObjectIDs() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsPersistentObjectIDs"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsPersistentObjectIDs(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsPersistentObjectIDs:"), value)
}

// A Boolean property that indicates whether the volume supports 64-bit object
// IDs.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supports64BitObjectIDs
func (v FSVolumeSupportedCapabilities) Supports64BitObjectIDs() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supports64BitObjectIDs"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupports64BitObjectIDs(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupports64BitObjectIDs:"), value)
}

// A Boolean property that indicates whether the volume supports document IDs
// for document revisions.
//
// # Discussion
//
// A document ID is an identifier that persists across object ID changes.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsDocumentID
func (v FSVolumeSupportedCapabilities) SupportsDocumentID() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsDocumentID"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsDocumentID(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsDocumentID:"), value)
}

// A Boolean property that indicates whether the volume supports symbolic
// links.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsSymbolicLinks
func (v FSVolumeSupportedCapabilities) SupportsSymbolicLinks() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsSymbolicLinks"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsSymbolicLinks(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsSymbolicLinks:"), value)
}

// A Boolean property that indicates whether the volume supports hard links.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsHardLinks
func (v FSVolumeSupportedCapabilities) SupportsHardLinks() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsHardLinks"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsHardLinks(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsHardLinks:"), value)
}

// A Boolean property that indicates whether the volume supports a journal
// used to speed recovery in case of unplanned restart, such as a power outage
// or crash.
//
// # Discussion
//
// This property doesn’t necessarily mean the volume is actively using a
// journal.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsJournal
func (v FSVolumeSupportedCapabilities) SupportsJournal() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsJournal"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsJournal(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsJournal:"), value)
}

// A Boolean property that indicates whether the volume currently uses a
// journal for speeding recovery after an unplanned shutdown.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsActiveJournal
func (v FSVolumeSupportedCapabilities) SupportsActiveJournal() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsActiveJournal"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsActiveJournal(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsActiveJournal:"), value)
}

// A Boolan property that indicates the volume doesn’t store reliable times
// for the root directory.
//
// # Discussion
//
// If this value is `true` (Swift) or [YES] (Objective-C), the volume
// doesn’t store reliable times for the root directory.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportRootTimes
func (v FSVolumeSupportedCapabilities) DoesNotSupportRootTimes() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("doesNotSupportRootTimes"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetDoesNotSupportRootTimes(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setDoesNotSupportRootTimes:"), value)
}

// A Boolean property that indicates whether the volume supports sparse files.
//
// # Discussion
//
// A sparse file is a file that can have “holes” that the file system has
// never written to, and as a result don’t consume space on disk.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsSparseFiles
func (v FSVolumeSupportedCapabilities) SupportsSparseFiles() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsSparseFiles"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsSparseFiles(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsSparseFiles:"), value)
}

// A Boolean property that indicates whether the volume supports zero runs
//
// # Discussion
//
// If this value is true, the volume keeps track of allocated but unwritten
// runs of a file so that it can substitute zeroes without actually writing
// zeroes to the media.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsZeroRuns
func (v FSVolumeSupportedCapabilities) SupportsZeroRuns() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsZeroRuns"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsZeroRuns(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsZeroRuns:"), value)
}

// A Boolean property that indicates whether the volume supports fast results
// when fetching file system statistics.
//
// # Discussion
//
// A true value means this volume hints to upper layers to indicate that
// `statfs(2)` is fast enough that its results need not be cached by the
// caller.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsFastStatFS
func (v FSVolumeSupportedCapabilities) SupportsFastStatFS() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsFastStatFS"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsFastStatFS(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsFastStatFS:"), value)
}

// A Boolean property that indicates whether the volume supports file sizes
// larger than 4GB, and potentially up to 2TB.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supports2TBFiles
func (v FSVolumeSupportedCapabilities) Supports2TBFiles() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supports2TBFiles"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupports2TBFiles(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupports2TBFiles:"), value)
}

// A Boolean property that indicates whether the volume supports open deny
// modes.
//
// # Discussion
//
// These are modes such as “open for read write, deny write”.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsOpenDenyModes
func (v FSVolumeSupportedCapabilities) SupportsOpenDenyModes() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsOpenDenyModes"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsOpenDenyModes(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsOpenDenyModes:"), value)
}

// A Boolean property that indicates whether the volume supports hidden files.
//
// # Discussion
//
// A `true` value means the volume supports the `UF_HIDDEN` file flag.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsHiddenFiles
func (v FSVolumeSupportedCapabilities) SupportsHiddenFiles() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsHiddenFiles"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsHiddenFiles(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsHiddenFiles:"), value)
}

// A Boolean property that indicates the volume doesn’t support immutable
// files.
//
// # Discussion
//
// A `true` value means this volume doesn’t support setting the
// `UF_IMMUTABLE` flag.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportImmutableFiles
func (v FSVolumeSupportedCapabilities) DoesNotSupportImmutableFiles() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("doesNotSupportImmutableFiles"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetDoesNotSupportImmutableFiles(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setDoesNotSupportImmutableFiles:"), value)
}

// A Boolean property that indicates the volume doesn’t set file
// permissions.
//
// # Discussion
//
// If this value is `true` (Swift) or [YES] (Objective-C), the volume
// doesn’t support setting file permissions.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportSettingFilePermissions
func (v FSVolumeSupportedCapabilities) DoesNotSupportSettingFilePermissions() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("doesNotSupportSettingFilePermissions"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetDoesNotSupportSettingFilePermissions(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setDoesNotSupportSettingFilePermissions:"), value)
}

// A Boolean property that indicates whether the volume supports multiple
// logical file systems that share space in a single “partition.”
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsSharedSpace
func (v FSVolumeSupportedCapabilities) SupportsSharedSpace() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsSharedSpace"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsSharedSpace(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsSharedSpace:"), value)
}

// A Boolean property that indicates whether the volume supports volume
// groups.
//
// # Discussion
//
// Volume groups involve multiple logical file systems that the system can
// mount and unmount together, and for which the system can present common
// file system identifier information.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/supportsVolumeGroups
func (v FSVolumeSupportedCapabilities) SupportsVolumeGroups() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsVolumeGroups"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetSupportsVolumeGroups(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsVolumeGroups:"), value)
}

// A Boolean property that indicates the volume doesn’t support certain
// volume size reports.
//
// # Discussion
//
// A true value means the volume doesn’t support determining values for
// total data blocks, available blocks, or free blocks, as in `f_blocks`,
// `f_bavail`, and `f_bfree` in the struct `statFS` returned by `statfs(2)`.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/doesNotSupportVolumeSizes
func (v FSVolumeSupportedCapabilities) DoesNotSupportVolumeSizes() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("doesNotSupportVolumeSizes"))
	return rv
}
func (v FSVolumeSupportedCapabilities) SetDoesNotSupportVolumeSizes(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setDoesNotSupportVolumeSizes:"), value)
}

// A value that indicates the volume’s support for case sensitivity.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities/caseFormat
func (v FSVolumeSupportedCapabilities) CaseFormat() FSVolumeCaseFormat {
	rv := objc.Send[FSVolumeCaseFormat](v.ID, objc.Sel("caseFormat"))
	return FSVolumeCaseFormat(rv)
}
func (v FSVolumeSupportedCapabilities) SetCaseFormat(value FSVolumeCaseFormat) {
	objc.Send[struct{}](v.ID, objc.Sel("setCaseFormat:"), value)
}

// A property that provides the supported capabilities of the volume.
//
// See: https://developer.apple.com/documentation/fskit/fsvolume/operations/supportedvolumecapabilities
func (v FSVolumeSupportedCapabilities) SupportedVolumeCapabilities() IFSVolumeSupportedCapabilities {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("supportedVolumeCapabilities"))
	return FSVolumeSupportedCapabilitiesFromID(objc.ID(rv))
}
func (v FSVolumeSupportedCapabilities) SetSupportedVolumeCapabilities(value IFSVolumeSupportedCapabilities) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportedVolumeCapabilities:"), value)
}

// A property that provides up-to-date statistics of the volume.
//
// See: https://developer.apple.com/documentation/fskit/fsvolume/operations/volumestatistics
func (v FSVolumeSupportedCapabilities) VolumeStatistics() IFSStatFSResult {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("volumeStatistics"))
	return FSStatFSResultFromID(objc.ID(rv))
}
func (v FSVolumeSupportedCapabilities) SetVolumeStatistics(value IFSStatFSResult) {
	objc.Send[struct{}](v.ID, objc.Sel("setVolumeStatistics:"), value)
}

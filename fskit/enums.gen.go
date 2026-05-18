// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask
type FSAccessMask uint

const (
	// FSAccessAddFile: The file system allows adding files.
	FSAccessAddFile FSAccessMask = 4
	// FSAccessAddSubdirectory: The file system allows adding subdirectories.
	FSAccessAddSubdirectory FSAccessMask = 32
	// FSAccessAppendData: The file system allows appending data to a file.
	FSAccessAppendData FSAccessMask = 32
	// FSAccessDelete: The file system allows deleting a file.
	FSAccessDelete FSAccessMask = 16
	// FSAccessDeleteChild: The file system allows deleting subdirectories.
	FSAccessDeleteChild FSAccessMask = 64
	// FSAccessExecute: The file system allows file executuion.
	FSAccessExecute FSAccessMask = 8
	// FSAccessListDirectory: The file system allows listing directory contents.
	FSAccessListDirectory FSAccessMask = 2
	// FSAccessReadAttributes: The file system allows reading file attributes.
	FSAccessReadAttributes FSAccessMask = 128
	// FSAccessReadData: The file system allows reading data.
	FSAccessReadData FSAccessMask = 2
	// FSAccessReadSecurity: The file system allows reading a file’s security descriptors.
	FSAccessReadSecurity FSAccessMask = 2048
	// FSAccessReadXattr: The file system allows reading extended file attributes.
	FSAccessReadXattr FSAccessMask = 512
	// FSAccessSearch: The file system allows searching files.
	FSAccessSearch FSAccessMask = 8
	// FSAccessTakeOwnership: The file system allows taking ownership of a file.
	FSAccessTakeOwnership FSAccessMask = 8192
	// FSAccessWriteAttributes: The file system allows writing file attributes.
	FSAccessWriteAttributes FSAccessMask = 256
	// FSAccessWriteData: The file system allows writing data.
	FSAccessWriteData FSAccessMask = 4
	// FSAccessWriteSecurity: The file system allows writing a file’s security descriptors.
	FSAccessWriteSecurity FSAccessMask = 4096
	// FSAccessWriteXattr: The file system allows writing extended file attributes.
	FSAccessWriteXattr FSAccessMask = 1024
)

func (e FSAccessMask) String() string {
	switch e {
	case FSAccessAddFile:
		return "FSAccessAddFile"
	case FSAccessAddSubdirectory:
		return "FSAccessAddSubdirectory"
	case FSAccessDelete:
		return "FSAccessDelete"
	case FSAccessDeleteChild:
		return "FSAccessDeleteChild"
	case FSAccessExecute:
		return "FSAccessExecute"
	case FSAccessListDirectory:
		return "FSAccessListDirectory"
	case FSAccessReadAttributes:
		return "FSAccessReadAttributes"
	case FSAccessReadSecurity:
		return "FSAccessReadSecurity"
	case FSAccessReadXattr:
		return "FSAccessReadXattr"
	case FSAccessTakeOwnership:
		return "FSAccessTakeOwnership"
	case FSAccessWriteAttributes:
		return "FSAccessWriteAttributes"
	case FSAccessWriteSecurity:
		return "FSAccessWriteSecurity"
	case FSAccessWriteXattr:
		return "FSAccessWriteXattr"
	default:
		return fmt.Sprintf("FSAccessMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSBlockmapFlags
type FSBlockmapFlags uint

const (
	// FSBlockmapFlagsRead: A flag that describes a read operation.
	FSBlockmapFlagsRead FSBlockmapFlags = 0x100
	// FSBlockmapFlagsWrite: A flag that describes a write operation.
	FSBlockmapFlagsWrite FSBlockmapFlags = 0x200
)

func (e FSBlockmapFlags) String() string {
	switch e {
	case FSBlockmapFlagsRead:
		return "FSBlockmapFlagsRead"
	case FSBlockmapFlagsWrite:
		return "FSBlockmapFlagsWrite"
	default:
		return fmt.Sprintf("FSBlockmapFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSCompleteIOFlags
type FSCompleteIOFlags uint

const (
	// FSCompleteIOFlagsAsync: A flag that requests that the file system module flush metadata I/O asynchronously.
	FSCompleteIOFlagsAsync FSCompleteIOFlags = 0x400
	// FSCompleteIOFlagsRead: A flag that describes a read operation.
	FSCompleteIOFlagsRead FSCompleteIOFlags = 256
	// FSCompleteIOFlagsWrite: A flag that describes a write operation.
	FSCompleteIOFlagsWrite FSCompleteIOFlags = 512
)

func (e FSCompleteIOFlags) String() string {
	switch e {
	case FSCompleteIOFlagsAsync:
		return "FSCompleteIOFlagsAsync"
	case FSCompleteIOFlagsRead:
		return "FSCompleteIOFlagsRead"
	case FSCompleteIOFlagsWrite:
		return "FSCompleteIOFlagsWrite"
	default:
		return fmt.Sprintf("FSCompleteIOFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSContainerState
type FSContainerState int

const (
	// FSContainerStateActive: The container is active, and one or more volumes are active.
	FSContainerStateActive FSContainerState = 3
	// FSContainerStateBlocked: The container is blocked from transitioning from the not-ready state to the ready state by a potentially-recoverable error.
	FSContainerStateBlocked FSContainerState = 1
	// FSContainerStateNotReady: The container isn’t ready.
	FSContainerStateNotReady FSContainerState = 0
	// FSContainerStateReady: The container is ready, but inactive.
	FSContainerStateReady FSContainerState = 2
)

func (e FSContainerState) String() string {
	switch e {
	case FSContainerStateActive:
		return "FSContainerStateActive"
	case FSContainerStateBlocked:
		return "FSContainerStateBlocked"
	case FSContainerStateNotReady:
		return "FSContainerStateNotReady"
	case FSContainerStateReady:
		return "FSContainerStateReady"
	default:
		return fmt.Sprintf("FSContainerState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSDeactivateOptions
type FSDeactivateOptions int

const (
	// FSDeactivateOptionsForce: An option to force deactivation.
	FSDeactivateOptionsForce FSDeactivateOptions = 1
)

func (e FSDeactivateOptions) String() string {
	switch e {
	case FSDeactivateOptionsForce:
		return "FSDeactivateOptionsForce"
	default:
		return fmt.Sprintf("FSDeactivateOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSError/Code
type FSErrorCode int

const (
	// FSErrorInvalidDirectoryCookie: While enumerating a directory, the given cookie didn’t resolve to a valid directory entry.
	FSErrorInvalidDirectoryCookie FSErrorCode = 4506
	// FSErrorModuleLoadFailed: The module failed to load.
	FSErrorModuleLoadFailed FSErrorCode = 4500
	// FSErrorResourceDamaged: The resource is damaged.
	FSErrorResourceDamaged FSErrorCode = 4502
	// FSErrorResourceUnrecognized: FSKit didn’t recognize the resource, and probing failed to find a match.
	FSErrorResourceUnrecognized FSErrorCode = 4501
	// FSErrorResourceUnusable: FSKit recognizes the resource, but the resource isn’t usable.
	FSErrorResourceUnusable FSErrorCode = 4503
	// FSErrorStatusOperationInProgress: An operation is in progress.
	FSErrorStatusOperationInProgress FSErrorCode = 4504
	// FSErrorStatusOperationPaused: An operation is paused.
	FSErrorStatusOperationPaused FSErrorCode = 4505
)

func (e FSErrorCode) String() string {
	switch e {
	case FSErrorInvalidDirectoryCookie:
		return "FSErrorInvalidDirectoryCookie"
	case FSErrorModuleLoadFailed:
		return "FSErrorModuleLoadFailed"
	case FSErrorResourceDamaged:
		return "FSErrorResourceDamaged"
	case FSErrorResourceUnrecognized:
		return "FSErrorResourceUnrecognized"
	case FSErrorResourceUnusable:
		return "FSErrorResourceUnusable"
	case FSErrorStatusOperationInProgress:
		return "FSErrorStatusOperationInProgress"
	case FSErrorStatusOperationPaused:
		return "FSErrorStatusOperationPaused"
	default:
		return fmt.Sprintf("FSErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSExtentType
type FSExtentType int

const (
	// FSExtentTypeData: An extent type to indicate valid data.
	FSExtentTypeData FSExtentType = 0
	// FSExtentTypeZeroFill: An extent type to indicate uninitialized data.
	FSExtentTypeZeroFill FSExtentType = 1
)

func (e FSExtentType) String() string {
	switch e {
	case FSExtentTypeData:
		return "FSExtentTypeData"
	case FSExtentTypeZeroFill:
		return "FSExtentTypeZeroFill"
	default:
		return fmt.Sprintf("FSExtentType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSItem/Attribute
type FSItemAttribute int

const (
	// FSItemAttributeAccessTime: The last-accessed time attribute.
	FSItemAttributeAccessTime FSItemAttribute = 1024
	// FSItemAttributeAddedTime: The time added attribute.
	FSItemAttributeAddedTime FSItemAttribute = 32768
	// FSItemAttributeAllocSize: The allocated size attribute.
	FSItemAttributeAllocSize FSItemAttribute = 128
	// FSItemAttributeBackupTime: The backup time attribute.
	FSItemAttributeBackupTime FSItemAttribute = 16384
	// FSItemAttributeBirthTime: The creation time attribute.
	FSItemAttributeBirthTime FSItemAttribute = 8192
	// FSItemAttributeChangeTime: The last-changed time attribute.
	FSItemAttributeChangeTime FSItemAttribute = 4096
	// FSItemAttributeFileID: The file ID attribute.
	FSItemAttributeFileID FSItemAttribute = 256
	// FSItemAttributeFlags: The flags attribute.
	FSItemAttributeFlags FSItemAttribute = 32
	// FSItemAttributeGID: The group ID (gid) attribute.
	FSItemAttributeGID FSItemAttribute = 16
	// FSItemAttributeInhibitKernelOffloadedIO: The inhibit kernel offloaded I/O attribute.
	FSItemAttributeInhibitKernelOffloadedIO FSItemAttribute = 131072
	// FSItemAttributeLinkCount: The link count attribute.
	FSItemAttributeLinkCount FSItemAttribute = 4
	// FSItemAttributeMode: The mode attribute.
	FSItemAttributeMode FSItemAttribute = 2
	// FSItemAttributeModifyTime: The last-modified time attribute.
	FSItemAttributeModifyTime FSItemAttribute = 2048
	// FSItemAttributeParentID: The parent ID attribute.
	FSItemAttributeParentID FSItemAttribute = 512
	// FSItemAttributeSize: The size attribute.
	FSItemAttributeSize FSItemAttribute = 64
	// FSItemAttributeSupportsLimitedXAttrs: The supports limited extended attributes attribute.
	FSItemAttributeSupportsLimitedXAttrs FSItemAttribute = 65536
	// FSItemAttributeType: The type attribute.
	FSItemAttributeType FSItemAttribute = 1
	// FSItemAttributeUID: The user ID (uid) attribute.
	FSItemAttributeUID FSItemAttribute = 8
)

func (e FSItemAttribute) String() string {
	switch e {
	case FSItemAttributeAccessTime:
		return "FSItemAttributeAccessTime"
	case FSItemAttributeAddedTime:
		return "FSItemAttributeAddedTime"
	case FSItemAttributeAllocSize:
		return "FSItemAttributeAllocSize"
	case FSItemAttributeBackupTime:
		return "FSItemAttributeBackupTime"
	case FSItemAttributeBirthTime:
		return "FSItemAttributeBirthTime"
	case FSItemAttributeChangeTime:
		return "FSItemAttributeChangeTime"
	case FSItemAttributeFileID:
		return "FSItemAttributeFileID"
	case FSItemAttributeFlags:
		return "FSItemAttributeFlags"
	case FSItemAttributeGID:
		return "FSItemAttributeGID"
	case FSItemAttributeInhibitKernelOffloadedIO:
		return "FSItemAttributeInhibitKernelOffloadedIO"
	case FSItemAttributeLinkCount:
		return "FSItemAttributeLinkCount"
	case FSItemAttributeMode:
		return "FSItemAttributeMode"
	case FSItemAttributeModifyTime:
		return "FSItemAttributeModifyTime"
	case FSItemAttributeParentID:
		return "FSItemAttributeParentID"
	case FSItemAttributeSize:
		return "FSItemAttributeSize"
	case FSItemAttributeSupportsLimitedXAttrs:
		return "FSItemAttributeSupportsLimitedXAttrs"
	case FSItemAttributeType:
		return "FSItemAttributeType"
	case FSItemAttributeUID:
		return "FSItemAttributeUID"
	default:
		return fmt.Sprintf("FSItemAttribute(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivationOptions
type FSItemDeactivationOptions int

const (
	// FSItemDeactivationAlways: An option to always perform deactivation calls.
	FSItemDeactivationAlways FSItemDeactivationOptions = -1
	// FSItemDeactivationForPreallocatedItems: An option to process deactivation for for files with preallocated space.
	FSItemDeactivationForPreallocatedItems FSItemDeactivationOptions = 2
	// FSItemDeactivationForRemovedItems: An option to process deactivation for open-unlinked items at the moment of last close.
	FSItemDeactivationForRemovedItems FSItemDeactivationOptions = 1
	// FSItemDeactivationNever: An option to never perform deactivation.
	FSItemDeactivationNever FSItemDeactivationOptions = 0
)

func (e FSItemDeactivationOptions) String() string {
	switch e {
	case FSItemDeactivationAlways:
		return "FSItemDeactivationAlways"
	case FSItemDeactivationForPreallocatedItems:
		return "FSItemDeactivationForPreallocatedItems"
	case FSItemDeactivationForRemovedItems:
		return "FSItemDeactivationForRemovedItems"
	case FSItemDeactivationNever:
		return "FSItemDeactivationNever"
	default:
		return fmt.Sprintf("FSItemDeactivationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSItem/Identifier
type FSItemID uint64

const (
	// FSItemIDInvalid: The identifier for an invalid item.
	FSItemIDInvalid FSItemID = 0
	// FSItemIDParentOfRoot: The identifier for an item that serves as the parent of the root directory.
	FSItemIDParentOfRoot FSItemID = 1
	// FSItemIDRootDirectory: The item identifier for the root directory.
	FSItemIDRootDirectory FSItemID = 2
)

func (e FSItemID) String() string {
	switch e {
	case FSItemIDInvalid:
		return "FSItemIDInvalid"
	case FSItemIDParentOfRoot:
		return "FSItemIDParentOfRoot"
	case FSItemIDRootDirectory:
		return "FSItemIDRootDirectory"
	default:
		return fmt.Sprintf("FSItemID(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSItem/ItemType
type FSItemType int

const (
	// FSItemTypeBlockDevice: The item type of a block device.
	FSItemTypeBlockDevice FSItemType = 6
	// FSItemTypeCharDevice: The item type of a character device.
	FSItemTypeCharDevice FSItemType = 5
	// FSItemTypeDirectory: The item type of a directory.
	FSItemTypeDirectory FSItemType = 2
	// FSItemTypeFIFO: The item type of a first-in/first-out named pipe.
	FSItemTypeFIFO FSItemType = 4
	// FSItemTypeFile: The item type of a regular file.
	FSItemTypeFile FSItemType = 1
	// FSItemTypeSocket: The item type of a socket.
	FSItemTypeSocket FSItemType = 7
	// FSItemTypeSymlink: The item type of a symbolic link.
	FSItemTypeSymlink FSItemType = 3
	// FSItemTypeUnknown: The item type of an unknown item.
	FSItemTypeUnknown FSItemType = 0
)

func (e FSItemType) String() string {
	switch e {
	case FSItemTypeBlockDevice:
		return "FSItemTypeBlockDevice"
	case FSItemTypeCharDevice:
		return "FSItemTypeCharDevice"
	case FSItemTypeDirectory:
		return "FSItemTypeDirectory"
	case FSItemTypeFIFO:
		return "FSItemTypeFIFO"
	case FSItemTypeFile:
		return "FSItemTypeFile"
	case FSItemTypeSocket:
		return "FSItemTypeSocket"
	case FSItemTypeSymlink:
		return "FSItemTypeSymlink"
	case FSItemTypeUnknown:
		return "FSItemTypeUnknown"
	default:
		return fmt.Sprintf("FSItemType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSMatchResult
type FSMatchResult int

const (
	// FSMatchResultNotRecognized: The probe doesn’t recognize the resource.
	FSMatchResultNotRecognized FSMatchResult = 0
	// FSMatchResultRecognized: The probe recognizes the resource but can’t use it.
	FSMatchResultRecognized FSMatchResult = 1
	// FSMatchResultUsable: The probe recognizes the resource and is ready to use it.
	FSMatchResultUsable FSMatchResult = 3
	// FSMatchResultUsableButLimited: The probe recognizes the resource and is ready to use it, but only in a limited capacity.
	FSMatchResultUsableButLimited FSMatchResult = 2
)

func (e FSMatchResult) String() string {
	switch e {
	case FSMatchResultNotRecognized:
		return "FSMatchResultNotRecognized"
	case FSMatchResultRecognized:
		return "FSMatchResultRecognized"
	case FSMatchResultUsable:
		return "FSMatchResultUsable"
	case FSMatchResultUsableButLimited:
		return "FSMatchResultUsableButLimited"
	default:
		return fmt.Sprintf("FSMatchResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSVolume/MountOptions
type FSMountOptions uint

const (
	// FSMountOptionsReadOnly: An option to request a read-only mount.
	FSMountOptionsReadOnly FSMountOptions = 1
)

func (e FSMountOptions) String() string {
	switch e {
	case FSMountOptionsReadOnly:
		return "FSMountOptionsReadOnly"
	default:
		return fmt.Sprintf("FSMountOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateFlags
type FSPreallocateFlags uint

const (
	// FSPreallocateFlagsAll: Allocates all requested space or no space at all.
	FSPreallocateFlagsAll FSPreallocateFlags = 0x4
	// FSPreallocateFlagsContiguous: Allocates contiguous space.
	FSPreallocateFlagsContiguous FSPreallocateFlags = 0x2
	// FSPreallocateFlagsFromEOF: Allocates space from the physical end of file.
	FSPreallocateFlagsFromEOF FSPreallocateFlags = 0x10
	// FSPreallocateFlagsPersist: Allocates space that isn’t freed when deleting the descriptor.
	FSPreallocateFlagsPersist FSPreallocateFlags = 0x8
)

func (e FSPreallocateFlags) String() string {
	switch e {
	case FSPreallocateFlagsAll:
		return "FSPreallocateFlagsAll"
	case FSPreallocateFlagsContiguous:
		return "FSPreallocateFlagsContiguous"
	case FSPreallocateFlagsFromEOF:
		return "FSPreallocateFlagsFromEOF"
	case FSPreallocateFlagsPersist:
		return "FSPreallocateFlagsPersist"
	default:
		return fmt.Sprintf("FSPreallocateFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSVolume/SetXattrPolicy
type FSSetXattrPolicy uint

const (
	// FSSetXattrPolicyAlwaysSet: Set the value, regardless of previous state.
	FSSetXattrPolicyAlwaysSet FSSetXattrPolicy = 0
	// FSSetXattrPolicyDelete: Delete the value, failing if the extended attribute doesn’t exist.
	FSSetXattrPolicyDelete FSSetXattrPolicy = 3
	// FSSetXattrPolicyMustCreate: Set the value, but fail if the extended attribute already exists.
	FSSetXattrPolicyMustCreate FSSetXattrPolicy = 1
	// FSSetXattrPolicyMustReplace: Set the value, but fail if the extended attribute doesn’t already exist.
	FSSetXattrPolicyMustReplace FSSetXattrPolicy = 2
)

func (e FSSetXattrPolicy) String() string {
	switch e {
	case FSSetXattrPolicyAlwaysSet:
		return "FSSetXattrPolicyAlwaysSet"
	case FSSetXattrPolicyDelete:
		return "FSSetXattrPolicyDelete"
	case FSSetXattrPolicyMustCreate:
		return "FSSetXattrPolicyMustCreate"
	case FSSetXattrPolicyMustReplace:
		return "FSSetXattrPolicyMustReplace"
	default:
		return fmt.Sprintf("FSSetXattrPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSSyncFlags
type FSSyncFlags int

const (
	// FSSyncFlagsDWait: A flag for synchronized I/O with data-integrity completion.
	FSSyncFlagsDWait FSSyncFlags = 4
	// FSSyncFlagsNoWait: A flag for synchronized I/O that starts I/O but doesn’t wait for it.
	FSSyncFlagsNoWait FSSyncFlags = 2
	// FSSyncFlagsWait: A flag for synchronized I/O with file-integrity completion.
	FSSyncFlagsWait FSSyncFlags = 1
)

func (e FSSyncFlags) String() string {
	switch e {
	case FSSyncFlagsDWait:
		return "FSSyncFlagsDWait"
	case FSSyncFlagsNoWait:
		return "FSSyncFlagsNoWait"
	case FSSyncFlagsWait:
		return "FSSyncFlagsWait"
	default:
		return fmt.Sprintf("FSSyncFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSVolume/CaseFormat
type FSVolumeCaseFormat int

const (
	// FSVolumeCaseFormatInsensitive: The volume isn’t case sensitive.
	FSVolumeCaseFormatInsensitive FSVolumeCaseFormat = 1
	// FSVolumeCaseFormatInsensitiveCasePreserving: The volume isn’t case sensitive, but supports preserving the case of file and directory names.
	FSVolumeCaseFormatInsensitiveCasePreserving FSVolumeCaseFormat = 2
	// FSVolumeCaseFormatSensitive: The volume is case sensitive.
	FSVolumeCaseFormatSensitive FSVolumeCaseFormat = 0
)

func (e FSVolumeCaseFormat) String() string {
	switch e {
	case FSVolumeCaseFormatInsensitive:
		return "FSVolumeCaseFormatInsensitive"
	case FSVolumeCaseFormatInsensitiveCasePreserving:
		return "FSVolumeCaseFormatInsensitiveCasePreserving"
	case FSVolumeCaseFormatSensitive:
		return "FSVolumeCaseFormatSensitive"
	default:
		return fmt.Sprintf("FSVolumeCaseFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/FSKit/FSVolume/OpenModes
type FSVolumeOpenModes uint

const (
	// FSVolumeOpenModesRead: The read mode.
	FSVolumeOpenModesRead FSVolumeOpenModes = 0x1
	// FSVolumeOpenModesWrite: The write mode.
	FSVolumeOpenModesWrite FSVolumeOpenModes = 0x2
)

func (e FSVolumeOpenModes) String() string {
	switch e {
	case FSVolumeOpenModesRead:
		return "FSVolumeOpenModesRead"
	case FSVolumeOpenModesWrite:
		return "FSVolumeOpenModesWrite"
	default:
		return fmt.Sprintf("FSVolumeOpenModes(%d)", e)
	}
}

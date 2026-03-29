// Code generated from Apple documentation. DO NOT EDIT.

package diskarbitration

import (
	"unsafe"
	"github.com/tmc/apple/objectivec"
)

// See: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionRef
type DAApprovalSessionRef uintptr

// DADiskAppearedCallback is type of the callback function used by DARegisterDiskAppearedCallback().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskAppearedCallback
type DADiskAppearedCallback = func(uintptr, unsafe.Pointer)

// DADiskClaimCallback is type of the callback function used by DADiskClaim().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskClaimCallback
type DADiskClaimCallback = func(uintptr, uintptr, unsafe.Pointer)

// DADiskClaimReleaseCallback is type of the callback function used by DADiskClaim().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskClaimReleaseCallback
type DADiskClaimReleaseCallback = func(uintptr, unsafe.Pointer) uintptr

// DADiskDescriptionChangedCallback is type of the callback function used by DARegisterDiskDescriptionChangedCallback().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskDescriptionChangedCallback
type DADiskDescriptionChangedCallback = func(uintptr, uintptr, unsafe.Pointer)

// DADiskDisappearedCallback is type of the callback function used by DARegisterDiskDisappearedCallback().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskDisappearedCallback
type DADiskDisappearedCallback = func(uintptr, unsafe.Pointer)

// DADiskEjectApprovalCallback is type of the callback function used by DARegisterDiskEjectApprovalCallback().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskEjectApprovalCallback
type DADiskEjectApprovalCallback = func(uintptr, unsafe.Pointer) uintptr

// DADiskEjectCallback is type of the callback function used by DADiskEject().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskEjectCallback
type DADiskEjectCallback = func(uintptr, uintptr, unsafe.Pointer)

// DADiskEjectOptions is options for DADiskEject().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskEjectOptions
type DADiskEjectOptions = uint32

// DADiskMountApprovalCallback is type of the callback function used by DARegisterDiskMountApprovalCallback().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskMountApprovalCallback
type DADiskMountApprovalCallback = func(uintptr, unsafe.Pointer) uintptr

// DADiskMountCallback is type of the callback function used by DADiskMount().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskMountCallback
type DADiskMountCallback = func(uintptr, uintptr, unsafe.Pointer)

// DADiskMountOptions is options for DADiskMount().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskMountOptions
type DADiskMountOptions = uint32

// DADiskOptions is options for DADiskGetOptions() and DADiskSetOptions().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskOptions
type DADiskOptions = uint32

// DADiskPeekCallback is type of the callback function used by DARegisterDiskPeekCallback().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskPeekCallback
type DADiskPeekCallback = func(uintptr, unsafe.Pointer)

// DADiskRef is type of a reference to DADisk instances.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADisk
type DADiskRef uintptr

// DADiskRenameCallback is type of the callback function used by DADiskRename().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskRenameCallback
type DADiskRenameCallback = func(uintptr, uintptr, unsafe.Pointer)

// DADiskRenameOptions is options for DADiskRename().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskRenameOptions
type DADiskRenameOptions = uint32

// DADiskUnmountApprovalCallback is type of the callback function used by DARegisterDiskUnmountApprovalCallback().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmountApprovalCallback
type DADiskUnmountApprovalCallback = func(uintptr, unsafe.Pointer) uintptr

// DADiskUnmountCallback is type of the callback function used by DADiskUnmount().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmountCallback
type DADiskUnmountCallback = func(uintptr, uintptr, unsafe.Pointer)

// DADiskUnmountOptions is options for DADiskUnmount().
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmountOptions
type DADiskUnmountOptions = uint32

// DADissenterRef is type of a reference to DADissenter instances.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADissenter
type DADissenterRef uintptr

// DAReturn is a return code.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAReturn
type DAReturn = objectivec.IObject

// DASessionRef is type of a reference to DASession instances.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASession
type DASessionRef uintptr


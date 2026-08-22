// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration

import (
	"fmt"
)

type KDADiskClaimOption uint32

const (
	KDADiskClaimOptionDefault KDADiskClaimOption = 0
)

func (e KDADiskClaimOption) String() string {
	switch e {
	case KDADiskClaimOptionDefault:
		return "KDADiskClaimOptionDefault"
	default:
		return fmt.Sprintf("KDADiskClaimOption(%d)", e)
	}
}

type KDADiskEjectOption uint32

const (
	KDADiskEjectOptionDefault KDADiskEjectOption = 0
)

func (e KDADiskEjectOption) String() string {
	switch e {
	case KDADiskEjectOptionDefault:
		return "KDADiskEjectOptionDefault"
	default:
		return fmt.Sprintf("KDADiskEjectOption(%d)", e)
	}
}

type KDADiskMountOption uint32

const (
	KDADiskMountOptionDefault  KDADiskMountOption = 0
	KDADiskMountOptionNoFollow KDADiskMountOption = 0x2
	// KDADiskMountOptionWhole: Mount the volumes tied to the whole disk object.
	KDADiskMountOptionWhole KDADiskMountOption = 0x1
)

func (e KDADiskMountOption) String() string {
	switch e {
	case KDADiskMountOptionDefault:
		return "KDADiskMountOptionDefault"
	case KDADiskMountOptionNoFollow:
		return "KDADiskMountOptionNoFollow"
	case KDADiskMountOptionWhole:
		return "KDADiskMountOptionWhole"
	default:
		return fmt.Sprintf("KDADiskMountOption(%d)", e)
	}
}

type KDADiskOption uint32

const (
	KDADiskOptionDefault KDADiskOption = 0
)

func (e KDADiskOption) String() string {
	switch e {
	case KDADiskOptionDefault:
		return "KDADiskOptionDefault"
	default:
		return fmt.Sprintf("KDADiskOption(%d)", e)
	}
}

type KDADiskRenameOption uint32

const (
	KDADiskRenameOptionDefault KDADiskRenameOption = 0
)

func (e KDADiskRenameOption) String() string {
	switch e {
	case KDADiskRenameOptionDefault:
		return "KDADiskRenameOptionDefault"
	default:
		return fmt.Sprintf("KDADiskRenameOption(%d)", e)
	}
}

type KDADiskUnmountOption uint32

const (
	KDADiskUnmountOptionDefault KDADiskUnmountOption = 0
	// KDADiskUnmountOptionForce: Unmount the volume even if files are still active.
	KDADiskUnmountOptionForce KDADiskUnmountOption = 0x80000
	// KDADiskUnmountOptionWhole: Unmount the volumes tied to the whole disk object.
	KDADiskUnmountOptionWhole KDADiskUnmountOption = 0x1
)

func (e KDADiskUnmountOption) String() string {
	switch e {
	case KDADiskUnmountOptionDefault:
		return "KDADiskUnmountOptionDefault"
	case KDADiskUnmountOptionForce:
		return "KDADiskUnmountOptionForce"
	case KDADiskUnmountOptionWhole:
		return "KDADiskUnmountOptionWhole"
	default:
		return fmt.Sprintf("KDADiskUnmountOption(%d)", e)
	}
}

type KDAReturn int32

const (
	KDAReturnBadArgument     KDAReturn = -119930877
	KDAReturnBusy            KDAReturn = -119930878
	KDAReturnError           KDAReturn = -119930879
	KDAReturnExclusiveAccess KDAReturn = -119930876
	KDAReturnNoResources     KDAReturn = -119930875
	KDAReturnNotFound        KDAReturn = -119930874
	KDAReturnNotMounted      KDAReturn = -119930873
	KDAReturnNotPermitted    KDAReturn = -119930872
	KDAReturnNotPrivileged   KDAReturn = -119930871
	KDAReturnNotReady        KDAReturn = -119930870
	KDAReturnNotWritable     KDAReturn = -119930869
	KDAReturnSuccess         KDAReturn = 0
	KDAReturnUnsupported     KDAReturn = -119930868
)

func (e KDAReturn) String() string {
	switch e {
	case KDAReturnBadArgument:
		return "KDAReturnBadArgument"
	case KDAReturnBusy:
		return "KDAReturnBusy"
	case KDAReturnError:
		return "KDAReturnError"
	case KDAReturnExclusiveAccess:
		return "KDAReturnExclusiveAccess"
	case KDAReturnNoResources:
		return "KDAReturnNoResources"
	case KDAReturnNotFound:
		return "KDAReturnNotFound"
	case KDAReturnNotMounted:
		return "KDAReturnNotMounted"
	case KDAReturnNotPermitted:
		return "KDAReturnNotPermitted"
	case KDAReturnNotPrivileged:
		return "KDAReturnNotPrivileged"
	case KDAReturnNotReady:
		return "KDAReturnNotReady"
	case KDAReturnNotWritable:
		return "KDAReturnNotWritable"
	case KDAReturnSuccess:
		return "KDAReturnSuccess"
	case KDAReturnUnsupported:
		return "KDAReturnUnsupported"
	default:
		return fmt.Sprintf("KDAReturn(%d)", e)
	}
}

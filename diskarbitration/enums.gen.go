// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration

import (
	"fmt"
)

type KDADiskClaimOption uint

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

type KDADiskMountOption uint

const (
	KDADiskMountOptionDefault  KDADiskMountOption = 0
	KDADiskMountOptionNoFollow KDADiskMountOption = 0
	// KDADiskMountOptionWhole: Mount the volumes tied to the whole disk object.
	KDADiskMountOptionWhole KDADiskMountOption = 0
)

func (e KDADiskMountOption) String() string {
	switch e {
	case KDADiskMountOptionDefault:
		return "KDADiskMountOptionDefault"
	default:
		return fmt.Sprintf("KDADiskMountOption(%d)", e)
	}
}

type KDADiskOption uint

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

type KDADiskRenameOption uint

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

type KDADiskUnmountOption uint

const (
	KDADiskUnmountOptionDefault KDADiskUnmountOption = 0
	// KDADiskUnmountOptionForce: Unmount the volume even if files are still active.
	KDADiskUnmountOptionForce KDADiskUnmountOption = 0
	// KDADiskUnmountOptionWhole: Unmount the volumes tied to the whole disk object.
	KDADiskUnmountOptionWhole KDADiskUnmountOption = 0
)

func (e KDADiskUnmountOption) String() string {
	switch e {
	case KDADiskUnmountOptionDefault:
		return "KDADiskUnmountOptionDefault"
	default:
		return fmt.Sprintf("KDADiskUnmountOption(%d)", e)
	}
}

type KDAReturn uint

const (
	KDAReturnBadArgument     KDAReturn = 0
	KDAReturnBusy            KDAReturn = 0
	KDAReturnError           KDAReturn = 0
	KDAReturnExclusiveAccess KDAReturn = 0
	KDAReturnNoResources     KDAReturn = 0
	KDAReturnNotFound        KDAReturn = 0
	KDAReturnNotMounted      KDAReturn = 0
	KDAReturnNotPermitted    KDAReturn = 0
	KDAReturnNotPrivileged   KDAReturn = 0
	KDAReturnNotReady        KDAReturn = 0
	KDAReturnNotWritable     KDAReturn = 0
	KDAReturnSuccess         KDAReturn = 0
	KDAReturnUnsupported     KDAReturn = 0
)

func (e KDAReturn) String() string {
	switch e {
	case KDAReturnBadArgument:
		return "KDAReturnBadArgument"
	default:
		return fmt.Sprintf("KDAReturn(%d)", e)
	}
}

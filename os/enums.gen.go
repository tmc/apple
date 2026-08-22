// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"fmt"
)

type OSActivityFlag uint32

const (
	// OSActivityFlagDefault: Creates a new activity and associates it as a child of any provided parent activity.
	OSActivityFlagDefault OSActivityFlag = 0
	// OSActivityFlagDetached: Creates a new activity that is independent of any provided parent activity.
	OSActivityFlagDetached OSActivityFlag = 0x1
	// OSActivityFlagIfNonePresent: Creates a new activity only if one is not already present.
	OSActivityFlagIfNonePresent OSActivityFlag = 0x2
)

func (e OSActivityFlag) String() string {
	switch e {
	case OSActivityFlagDefault:
		return "OSActivityFlagDefault"
	case OSActivityFlagDetached:
		return "OSActivityFlagDetached"
	case OSActivityFlagIfNonePresent:
		return "OSActivityFlagIfNonePresent"
	default:
		return fmt.Sprintf("OSActivityFlag(%d)", e)
	}
}

type OSClockid uint32

const (
	// OSClockMachAbsoluteTime: Units relative to Mach absolute time value.
	OSClockMachAbsoluteTime OSClockid = 32
)

func (e OSClockid) String() string {
	switch e {
	case OSClockMachAbsoluteTime:
		return "OSClockMachAbsoluteTime"
	default:
		return fmt.Sprintf("OSClockid(%d)", e)
	}
}

type OSSecurityConfig uint64

const (
	// OSSecurityConfigGuardObjects: # Discussion
	OSSecurityConfigGuardObjects OSSecurityConfig = 0x100
	// OSSecurityConfigHardenedHeap: # Discussion
	OSSecurityConfigHardenedHeap OSSecurityConfig = 0x1
	// OSSecurityConfigMte: # Discussion
	OSSecurityConfigMte OSSecurityConfig = 0x4
	// OSSecurityConfigNone: # Discussion
	OSSecurityConfigNone OSSecurityConfig = 0
	// OSSecurityConfigTpro: # Discussion
	OSSecurityConfigTpro OSSecurityConfig = 0x2
)

func (e OSSecurityConfig) String() string {
	switch e {
	case OSSecurityConfigGuardObjects:
		return "OSSecurityConfigGuardObjects"
	case OSSecurityConfigHardenedHeap:
		return "OSSecurityConfigHardenedHeap"
	case OSSecurityConfigMte:
		return "OSSecurityConfigMte"
	case OSSecurityConfigNone:
		return "OSSecurityConfigNone"
	case OSSecurityConfigTpro:
		return "OSSecurityConfigTpro"
	default:
		return fmt.Sprintf("OSSecurityConfig(%d)", e)
	}
}

type OSSyncWaitOnAddressFlags uint32

const (
	// OSSyncWaitOnAddressNone: Default behavior for futex functions that block a thread.
	OSSyncWaitOnAddressNone OSSyncWaitOnAddressFlags = 0
	// OSSyncWaitOnAddressShared: Flag to indicate an address is in a shared memory region, allowing for a futex wake from another process.
	OSSyncWaitOnAddressShared OSSyncWaitOnAddressFlags = 0x1
)

func (e OSSyncWaitOnAddressFlags) String() string {
	switch e {
	case OSSyncWaitOnAddressNone:
		return "OSSyncWaitOnAddressNone"
	case OSSyncWaitOnAddressShared:
		return "OSSyncWaitOnAddressShared"
	default:
		return fmt.Sprintf("OSSyncWaitOnAddressFlags(%d)", e)
	}
}

type OSSyncWakeByAddressFlags uint32

const (
	// OSSyncWakeByAddressNone: The default behavior for futex functions that wake a thread.
	OSSyncWakeByAddressNone OSSyncWakeByAddressFlags = 0
	// OSSyncWakeByAddressShared: A flag to indicate an address is in a shared memory region, allowing you to wake another waiting process.
	OSSyncWakeByAddressShared OSSyncWakeByAddressFlags = 0x1
)

func (e OSSyncWakeByAddressFlags) String() string {
	switch e {
	case OSSyncWakeByAddressNone:
		return "OSSyncWakeByAddressNone"
	case OSSyncWakeByAddressShared:
		return "OSSyncWakeByAddressShared"
	default:
		return fmt.Sprintf("OSSyncWakeByAddressFlags(%d)", e)
	}
}

type OSUnfairLockFlags uint32

const (
	OSUnfairLockFlagAdaptiveSpin OSUnfairLockFlags = 0x40000
	OSUnfairLockFlagNone         OSUnfairLockFlags = 0
)

func (e OSUnfairLockFlags) String() string {
	switch e {
	case OSUnfairLockFlagAdaptiveSpin:
		return "OSUnfairLockFlagAdaptiveSpin"
	case OSUnfairLockFlagNone:
		return "OSUnfairLockFlagNone"
	default:
		return fmt.Sprintf("OSUnfairLockFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/os/OSLogType
type OSLogType uint8

const (
	// OSLogTypeDebug: The debug log level.
	OSLogTypeDebug OSLogType = 0x2
	// OSLogTypeDefault: The default log level.
	OSLogTypeDefault OSLogType = 0
	// OSLogTypeError: The error log level.
	OSLogTypeError OSLogType = 0x10
	// OSLogTypeFault: The fault log level.
	OSLogTypeFault OSLogType = 0x11
	// OSLogTypeInfo: The informational log level.
	OSLogTypeInfo OSLogType = 0x1
)

func (e OSLogType) String() string {
	switch e {
	case OSLogTypeDebug:
		return "OSLogTypeDebug"
	case OSLogTypeDefault:
		return "OSLogTypeDefault"
	case OSLogTypeError:
		return "OSLogTypeError"
	case OSLogTypeFault:
		return "OSLogTypeFault"
	case OSLogTypeInfo:
		return "OSLogTypeInfo"
	default:
		return fmt.Sprintf("OSLogType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/os/OSSignpostType
type OSSignpostType uint8

const (
	// OSSignpostEvent: A signpost that marks an event in your code.
	OSSignpostEvent OSSignpostType = 0
	// OSSignpostIntervalBegin: A signpost that marks the start of a time interval of interest in your code.
	OSSignpostIntervalBegin OSSignpostType = 0x1
	// OSSignpostIntervalEnd: A signpost that marks the end of a time interval of interest in your code.
	OSSignpostIntervalEnd OSSignpostType = 0x2
)

func (e OSSignpostType) String() string {
	switch e {
	case OSSignpostEvent:
		return "OSSignpostEvent"
	case OSSignpostIntervalBegin:
		return "OSSignpostIntervalBegin"
	case OSSignpostIntervalEnd:
		return "OSSignpostIntervalEnd"
	default:
		return fmt.Sprintf("OSSignpostType(%d)", e)
	}
}

// Os_log_type_t is a C-name alias for OSLogType.
type Os_log_type_t = OSLogType

// Os_signpost_type_t is a C-name alias for OSSignpostType.
type Os_signpost_type_t = OSSignpostType

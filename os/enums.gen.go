// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"fmt"
)

type OsActivityFlag uint

const (
	// OS_ACTIVITY_FLAG_DEFAULT: Creates a new activity and associates it as a child of any provided parent activity.
	OS_ACTIVITY_FLAG_DEFAULT OsActivityFlag = 0
	// OS_ACTIVITY_FLAG_DETACHED: Creates a new activity that is independent of any provided parent activity.
	OS_ACTIVITY_FLAG_DETACHED OsActivityFlag = 0x1
	// OS_ACTIVITY_FLAG_IF_NONE_PRESENT: Creates a new activity only if one is not already present.
	OS_ACTIVITY_FLAG_IF_NONE_PRESENT OsActivityFlag = 0x2
)

func (e OsActivityFlag) String() string {
	switch e {
	case OS_ACTIVITY_FLAG_DEFAULT:
		return "OS_ACTIVITY_FLAG_DEFAULT"
	case OS_ACTIVITY_FLAG_DETACHED:
		return "OS_ACTIVITY_FLAG_DETACHED"
	case OS_ACTIVITY_FLAG_IF_NONE_PRESENT:
		return "OS_ACTIVITY_FLAG_IF_NONE_PRESENT"
	default:
		return fmt.Sprintf("OsActivityFlag(%d)", e)
	}
}

type OsClockMachAbsolute uint

const (
	// OS_CLOCK_MACH_ABSOLUTE_TIME: Units relative to Mach absolute time value.
	OS_CLOCK_MACH_ABSOLUTE_TIME OsClockMachAbsolute = 32
)

func (e OsClockMachAbsolute) String() string {
	switch e {
	case OS_CLOCK_MACH_ABSOLUTE_TIME:
		return "OS_CLOCK_MACH_ABSOLUTE_TIME"
	default:
		return fmt.Sprintf("OsClockMachAbsolute(%d)", e)
	}
}

type OsLogType uint

const (
	// OS_LOG_TYPE_DEBUG: The debug log level.
	OS_LOG_TYPE_DEBUG OsLogType = 0x2
	// OS_LOG_TYPE_DEFAULT: The default log level.
	OS_LOG_TYPE_DEFAULT OsLogType = 0
	// OS_LOG_TYPE_ERROR: The error log level.
	OS_LOG_TYPE_ERROR OsLogType = 0x10
	// OS_LOG_TYPE_FAULT: The fault log level.
	OS_LOG_TYPE_FAULT OsLogType = 0x11
	// OS_LOG_TYPE_INFO: The informational log level.
	OS_LOG_TYPE_INFO OsLogType = 0x1
)

func (e OsLogType) String() string {
	switch e {
	case OS_LOG_TYPE_DEBUG:
		return "OS_LOG_TYPE_DEBUG"
	case OS_LOG_TYPE_DEFAULT:
		return "OS_LOG_TYPE_DEFAULT"
	case OS_LOG_TYPE_ERROR:
		return "OS_LOG_TYPE_ERROR"
	case OS_LOG_TYPE_FAULT:
		return "OS_LOG_TYPE_FAULT"
	case OS_LOG_TYPE_INFO:
		return "OS_LOG_TYPE_INFO"
	default:
		return fmt.Sprintf("OsLogType(%d)", e)
	}
}

type OsSecurityConfig uint

const (
	// OS_SECURITY_CONFIG_GUARD_OBJECTS: # Discussion
	OS_SECURITY_CONFIG_GUARD_OBJECTS OsSecurityConfig = 0
	// OS_SECURITY_CONFIG_HARDENED_HEAP: # Discussion
	OS_SECURITY_CONFIG_HARDENED_HEAP OsSecurityConfig = 0
	// OS_SECURITY_CONFIG_MTE: # Discussion
	OS_SECURITY_CONFIG_MTE OsSecurityConfig = 0
	// OS_SECURITY_CONFIG_NONE: # Discussion
	OS_SECURITY_CONFIG_NONE OsSecurityConfig = 0
	// OS_SECURITY_CONFIG_TPRO: # Discussion
	OS_SECURITY_CONFIG_TPRO OsSecurityConfig = 0
)

func (e OsSecurityConfig) String() string {
	switch e {
	case OS_SECURITY_CONFIG_GUARD_OBJECTS:
		return "OS_SECURITY_CONFIG_GUARD_OBJECTS"
	default:
		return fmt.Sprintf("OsSecurityConfig(%d)", e)
	}
}

type OsSignpost uint

const (
	// OS_SIGNPOST_EVENT: A signpost that marks an event in your code.
	OS_SIGNPOST_EVENT OsSignpost = 0
	// OS_SIGNPOST_INTERVAL_BEGIN: A signpost that marks the start of a time interval of interest in your code.
	OS_SIGNPOST_INTERVAL_BEGIN OsSignpost = 0x1
	// OS_SIGNPOST_INTERVAL_END: A signpost that marks the end of a time interval of interest in your code.
	OS_SIGNPOST_INTERVAL_END OsSignpost = 0x2
)

func (e OsSignpost) String() string {
	switch e {
	case OS_SIGNPOST_EVENT:
		return "OS_SIGNPOST_EVENT"
	case OS_SIGNPOST_INTERVAL_BEGIN:
		return "OS_SIGNPOST_INTERVAL_BEGIN"
	case OS_SIGNPOST_INTERVAL_END:
		return "OS_SIGNPOST_INTERVAL_END"
	default:
		return fmt.Sprintf("OsSignpost(%d)", e)
	}
}

type OsSyncWaitOnAddress uint

const (
	// OS_SYNC_WAIT_ON_ADDRESS_NONE: Default behavior for futex functions that block a thread.
	OS_SYNC_WAIT_ON_ADDRESS_NONE OsSyncWaitOnAddress = 0
	// OS_SYNC_WAIT_ON_ADDRESS_SHARED: Flag to indicate an address is in a shared memory region, allowing for a futex wake from another process.
	OS_SYNC_WAIT_ON_ADDRESS_SHARED OsSyncWaitOnAddress = 0
)

func (e OsSyncWaitOnAddress) String() string {
	switch e {
	case OS_SYNC_WAIT_ON_ADDRESS_NONE:
		return "OS_SYNC_WAIT_ON_ADDRESS_NONE"
	default:
		return fmt.Sprintf("OsSyncWaitOnAddress(%d)", e)
	}
}

type OsSyncWakeByAddress uint

const (
	// OS_SYNC_WAKE_BY_ADDRESS_NONE: The default behavior for futex functions that wake a thread.
	OS_SYNC_WAKE_BY_ADDRESS_NONE OsSyncWakeByAddress = 0
	// OS_SYNC_WAKE_BY_ADDRESS_SHARED: A flag to indicate an address is in a shared memory region, allowing you to wake another waiting process.
	OS_SYNC_WAKE_BY_ADDRESS_SHARED OsSyncWakeByAddress = 0
)

func (e OsSyncWakeByAddress) String() string {
	switch e {
	case OS_SYNC_WAKE_BY_ADDRESS_NONE:
		return "OS_SYNC_WAKE_BY_ADDRESS_NONE"
	default:
		return fmt.Sprintf("OsSyncWakeByAddress(%d)", e)
	}
}

type OsUnfairLockFlag uint

const (
	OS_UNFAIR_LOCK_FLAG_ADAPTIVE_SPIN OsUnfairLockFlag = 0x40000
	OS_UNFAIR_LOCK_FLAG_NONE          OsUnfairLockFlag = 0
)

func (e OsUnfairLockFlag) String() string {
	switch e {
	case OS_UNFAIR_LOCK_FLAG_ADAPTIVE_SPIN:
		return "OS_UNFAIR_LOCK_FLAG_ADAPTIVE_SPIN"
	case OS_UNFAIR_LOCK_FLAG_NONE:
		return "OS_UNFAIR_LOCK_FLAG_NONE"
	default:
		return fmt.Sprintf("OsUnfairLockFlag(%d)", e)
	}
}

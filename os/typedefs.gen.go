// Code generated from Apple documentation. DO NOT EDIT.

package os

import (
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objectivec"
)

type Os_activity_flag_t = OsActivityFlag

// Os_activity_id_t is a number that uniquely identifies an activity.
//
// See: https://developer.apple.com/documentation/os/os_activity_id_t
type Os_activity_id_t = uint64

// Os_activity_scope_state_t is an opaque structure that contains a saved activity-execution context.
//
// See: https://developer.apple.com/documentation/os/os_activity_scope_state_t
type Os_activity_scope_state_t = uintptr

// Os_activity_t is an object that represents an activity triggered by the user.
//
// See: https://developer.apple.com/documentation/os/os_activity_t
type Os_activity_t = objectivec.Object

// Os_block_t is a block that takes no arguments and returns no value.
//
// See: https://developer.apple.com/documentation/os/os_block_t
type Os_block_t = func()

type Os_clockid_t = OsClockMachAbsolute

// Os_function_t is a pointer to a function.
//
// See: https://developer.apple.com/documentation/os/os_function_t
type Os_function_t = func(kernel.Pointer)

// Os_log_t is a log object that you pass to logging functions to send messages to that log.
//
// See: https://developer.apple.com/documentation/os/os_log_t
type Os_log_t = objectivec.Object

type Os_log_type_t = OsLogType

type Os_security_config_t = OsSecurityConfig

// Os_signpost_id_t is an identifier you use to distinguish between signposts that have the same name and destination log.
//
// See: https://developer.apple.com/documentation/os/os_signpost_id_t
type Os_signpost_id_t = uint64

type Os_signpost_type_t = OsSignpost

type Os_sync_wait_on_address_flags_t = OsSyncWaitOnAddress

type Os_sync_wake_by_address_flags_t = OsSyncWakeByAddress

type Os_unfair_lock_flags_t = OsUnfairLockFlag

// Os_unfair_lock_t is a pointer to an unfair lock structure.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_t
type Os_unfair_lock_t = uintptr

// See: https://developer.apple.com/documentation/os/os_workgroup_attr_s
type Os_workgroup_attr_s = Os_workgroup_attr_opaque_s

// Os_workgroup_attr_t is an opaque structure for storing workgroup-related attributes.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_attr_t
type Os_workgroup_attr_t = uintptr

// Os_workgroup_index is a unique index that the workgroup assigns to its joined threads.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_index
type Os_workgroup_index = uint32

// See: https://developer.apple.com/documentation/os/os_workgroup_interval_data_s
type Os_workgroup_interval_data_s = Os_workgroup_interval_data_opaque_s

// Os_workgroup_interval_data_t is an opaque structure that contains additional configuration data for the interval workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_interval_data_t
type Os_workgroup_interval_data_t = uintptr

// Os_workgroup_interval_t is a workgroup object that supports the scheduling of threads on a repeating cadence.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_interval_t
type Os_workgroup_interval_t = kernel.Pointer

// See: https://developer.apple.com/documentation/os/os_workgroup_join_token_s
type Os_workgroup_join_token_s = Os_workgroup_join_token_opaque_s

// Os_workgroup_join_token_t is an opaque token that represents a connection between a thread and a workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_join_token_t
type Os_workgroup_join_token_t = uintptr

// See: https://developer.apple.com/documentation/os/os_workgroup_mpt_attr_s
type Os_workgroup_mpt_attr_s = kernel.Pointer

// Os_workgroup_mpt_attr_t is an opaque structure containing attributes related to a request for the maximum number of parallel threads.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_mpt_attr_t
type Os_workgroup_mpt_attr_t = uintptr

// Os_workgroup_parallel_t is a workgroup object that supports the scheduling of threads that work in parallel to complete a task.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_parallel_t
type Os_workgroup_parallel_t = kernel.Pointer

// Os_workgroup_t is an opaque object representing a default workgroup in the current process.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_t
type Os_workgroup_t = *OS_os_workgroup

// Os_workgroup_working_arena_destructor_t is a function that deallocates a workgroup’s currently assigned shared memory.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_working_arena_destructor_t
type Os_workgroup_working_arena_destructor_t = func(kernel.Pointer)

// OSActivityFlag is a Go-name alias for Os_activity_flag_t.
type OSActivityFlag = Os_activity_flag_t

// OSActivityID is a Go-name alias for Os_activity_id_t.
type OSActivityID = Os_activity_id_t

// OSActivityScopeState is a Go-name alias for Os_activity_scope_state_t.
type OSActivityScopeState = Os_activity_scope_state_t

// OSActivity is a Go-name alias for Os_activity_t.
type OSActivity = Os_activity_t

// OSBlock is a Go-name alias for Os_block_t.
type OSBlock = Os_block_t

// OSClockid is a Go-name alias for Os_clockid_t.
type OSClockid = Os_clockid_t

// OSFunction is a Go-name alias for Os_function_t.
type OSFunction = Os_function_t

// OSLogType is a Go-name alias for Os_log_type_t.
type OSLogType = Os_log_type_t

// OSSecurityConfig is a Go-name alias for Os_security_config_t.
type OSSecurityConfig = Os_security_config_t

// OSSignpostID is a Go-name alias for Os_signpost_id_t.
type OSSignpostID = Os_signpost_id_t

// OSSignpostType is a Go-name alias for Os_signpost_type_t.
type OSSignpostType = Os_signpost_type_t

// OSSyncWaitOnAddressFlags is a Go-name alias for Os_sync_wait_on_address_flags_t.
type OSSyncWaitOnAddressFlags = Os_sync_wait_on_address_flags_t

// OSSyncWakeByAddressFlags is a Go-name alias for Os_sync_wake_by_address_flags_t.
type OSSyncWakeByAddressFlags = Os_sync_wake_by_address_flags_t

// OSUnfairLockFlags is a Go-name alias for Os_unfair_lock_flags_t.
type OSUnfairLockFlags = Os_unfair_lock_flags_t

// OSUnfairLock is a Go-name alias for Os_unfair_lock_t.
type OSUnfairLock = Os_unfair_lock_t

// OSWorkgroupAttrS is a Go-name alias for Os_workgroup_attr_s.
type OSWorkgroupAttrS = Os_workgroup_attr_s

// OSWorkgroupAttr is a Go-name alias for Os_workgroup_attr_t.
type OSWorkgroupAttr = Os_workgroup_attr_t

// OSWorkgroupIndex is a Go-name alias for Os_workgroup_index.
type OSWorkgroupIndex = Os_workgroup_index

// OSWorkgroupIntervalDataS is a Go-name alias for Os_workgroup_interval_data_s.
type OSWorkgroupIntervalDataS = Os_workgroup_interval_data_s

// OSWorkgroupIntervalData is a Go-name alias for Os_workgroup_interval_data_t.
type OSWorkgroupIntervalData = Os_workgroup_interval_data_t

// OSWorkgroupInterval is a Go-name alias for Os_workgroup_interval_t.
type OSWorkgroupInterval = Os_workgroup_interval_t

// OSWorkgroupJoinTokenS is a Go-name alias for Os_workgroup_join_token_s.
type OSWorkgroupJoinTokenS = Os_workgroup_join_token_s

// OSWorkgroupJoinToken is a Go-name alias for Os_workgroup_join_token_t.
type OSWorkgroupJoinToken = Os_workgroup_join_token_t

// OSWorkgroupMptAttrS is a Go-name alias for Os_workgroup_mpt_attr_s.
type OSWorkgroupMptAttrS = Os_workgroup_mpt_attr_s

// OSWorkgroupMptAttr is a Go-name alias for Os_workgroup_mpt_attr_t.
type OSWorkgroupMptAttr = Os_workgroup_mpt_attr_t

// OSWorkgroupParallel is a Go-name alias for Os_workgroup_parallel_t.
type OSWorkgroupParallel = Os_workgroup_parallel_t

// OSWorkgroup is a Go-name alias for Os_workgroup_t.
type OSWorkgroup = Os_workgroup_t

// OSWorkgroupWorkingArenaDestructor is a Go-name alias for Os_workgroup_working_arena_destructor_t.
type OSWorkgroupWorkingArenaDestructor = Os_workgroup_working_arena_destructor_t

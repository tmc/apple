// Code generated from Apple documentation. DO NOT EDIT.

package os

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OSActivityID is a number that uniquely identifies an activity.
//
// See: https://developer.apple.com/documentation/os/os_activity_id_t
type OSActivityID = uint64

// OSActivityScopeState is an opaque structure that contains a saved activity-execution context.
//
// See: https://developer.apple.com/documentation/os/os_activity_scope_state_t
type OSActivityScopeState = uintptr

// OSActivity is an object that represents an activity triggered by the user.
//
// See: https://developer.apple.com/documentation/os/os_activity_t
type OSActivity = objectivec.Object

// OSActivityFromID constructs a [OSActivity] from an objc.ID.
func OSActivityFromID(id objc.ID) OSActivity {
	return OSActivity{ID: id}
}

// OSBlock is a block that takes no arguments and returns no value.
//
// See: https://developer.apple.com/documentation/os/os_block_t
type OSBlock = func()

// OSFunction is a pointer to a function.
//
// See: https://developer.apple.com/documentation/os/os_function_t
type OSFunction = func(unsafe.Pointer)

// OSSignpostID is an identifier you use to distinguish between signposts that have the same name and destination log.
//
// See: https://developer.apple.com/documentation/os/os_signpost_id_t
type OSSignpostID = uint64

// OSUnfairLock is a structure that contains the data for an unfair lock.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock
type OSUnfairLock = [4]byte

// See: https://developer.apple.com/documentation/os/os_workgroup_attr_s
type OSWorkgroupAttrS = Os_workgroup_attr_opaque_s

// OSWorkgroupAttr is an opaque structure for storing workgroup-related attributes.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_attr_t
type OSWorkgroupAttr = uintptr

// OSWorkgroupIndex is a unique index that the workgroup assigns to its joined threads.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_index
type OSWorkgroupIndex = uint32

// See: https://developer.apple.com/documentation/os/os_workgroup_interval_data_s
type OSWorkgroupIntervalDataS = Os_workgroup_interval_data_opaque_s

// OSWorkgroupIntervalData is an opaque structure that contains additional configuration data for the interval workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_interval_data_t
type OSWorkgroupIntervalData = uintptr

// OSWorkgroupInterval is a workgroup object that supports the scheduling of threads on a repeating cadence.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_interval_t
type OSWorkgroupInterval = unsafe.Pointer

// See: https://developer.apple.com/documentation/os/os_workgroup_join_token_s
type OSWorkgroupJoinTokenS = Os_workgroup_join_token_opaque_s

// OSWorkgroupJoinToken is an opaque token that represents a connection between a thread and a workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_join_token_t
type OSWorkgroupJoinToken = uintptr

// See: https://developer.apple.com/documentation/os/os_workgroup_mpt_attr_s
// OSWorkgroupMptAttrS is an unresolved C aggregate typedef.
type OSWorkgroupMptAttrS unsafe.Pointer

// OSWorkgroupMptAttr is an opaque structure containing attributes related to a request for the maximum number of parallel threads.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_mpt_attr_t
type OSWorkgroupMptAttr = uintptr

// OSWorkgroupParallel is a workgroup object that supports the scheduling of threads that work in parallel to complete a task.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_parallel_t
type OSWorkgroupParallel = unsafe.Pointer

// OSWorkgroup is an opaque object representing a default workgroup in the current process.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_t
type OSWorkgroup = *OSOSWorkgroup

// OSWorkgroupWorkingArenaDestructor is a function that deallocates a workgroup’s currently assigned shared memory.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_working_arena_destructor_t
type OSWorkgroupWorkingArenaDestructor = func(unsafe.Pointer)

// Os_activity_flag_t is a C-name alias for OSActivityFlag.
type Os_activity_flag_t = OSActivityFlag

// Os_activity_id_t is a C-name alias for OSActivityID.
type Os_activity_id_t = OSActivityID

// Os_activity_scope_state_t is a C-name alias for OSActivityScopeState.
type Os_activity_scope_state_t = OSActivityScopeState

// Os_activity_t is a C-name alias for OSActivity.
type Os_activity_t = OSActivity

// Os_block_t is a C-name alias for OSBlock.
type Os_block_t = OSBlock

// Os_clockid_t is a C-name alias for OSClockid.
type Os_clockid_t = OSClockid

// Os_function_t is a C-name alias for OSFunction.
type Os_function_t = OSFunction

// Os_security_config_t is a C-name alias for OSSecurityConfig.
type Os_security_config_t = OSSecurityConfig

// Os_signpost_id_t is a C-name alias for OSSignpostID.
type Os_signpost_id_t = OSSignpostID

// Os_sync_wait_on_address_flags_t is a C-name alias for OSSyncWaitOnAddressFlags.
type Os_sync_wait_on_address_flags_t = OSSyncWaitOnAddressFlags

// Os_sync_wake_by_address_flags_t is a C-name alias for OSSyncWakeByAddressFlags.
type Os_sync_wake_by_address_flags_t = OSSyncWakeByAddressFlags

// Os_unfair_lock is a C-name alias for OSUnfairLock.
type Os_unfair_lock = OSUnfairLock

// Os_unfair_lock_flags_t is a C-name alias for OSUnfairLockFlags.
type Os_unfair_lock_flags_t = OSUnfairLockFlags

// Os_workgroup_attr_s is a C-name alias for OSWorkgroupAttrS.
type Os_workgroup_attr_s = OSWorkgroupAttrS

// Os_workgroup_attr_t is a C-name alias for OSWorkgroupAttr.
type Os_workgroup_attr_t = OSWorkgroupAttr

// Os_workgroup_index is a C-name alias for OSWorkgroupIndex.
type Os_workgroup_index = OSWorkgroupIndex

// Os_workgroup_interval_data_s is a C-name alias for OSWorkgroupIntervalDataS.
type Os_workgroup_interval_data_s = OSWorkgroupIntervalDataS

// Os_workgroup_interval_data_t is a C-name alias for OSWorkgroupIntervalData.
type Os_workgroup_interval_data_t = OSWorkgroupIntervalData

// Os_workgroup_interval_t is a C-name alias for OSWorkgroupInterval.
type Os_workgroup_interval_t = OSWorkgroupInterval

// Os_workgroup_join_token_s is a C-name alias for OSWorkgroupJoinTokenS.
type Os_workgroup_join_token_s = OSWorkgroupJoinTokenS

// Os_workgroup_join_token_t is a C-name alias for OSWorkgroupJoinToken.
type Os_workgroup_join_token_t = OSWorkgroupJoinToken

// Os_workgroup_mpt_attr_s is a C-name alias for OSWorkgroupMptAttrS.
type Os_workgroup_mpt_attr_s = OSWorkgroupMptAttrS

// Os_workgroup_mpt_attr_t is a C-name alias for OSWorkgroupMptAttr.
type Os_workgroup_mpt_attr_t = OSWorkgroupMptAttr

// Os_workgroup_parallel_t is a C-name alias for OSWorkgroupParallel.
type Os_workgroup_parallel_t = OSWorkgroupParallel

// Os_workgroup_t is a C-name alias for OSWorkgroup.
type Os_workgroup_t = OSWorkgroup

// Os_workgroup_working_arena_destructor_t is a C-name alias for OSWorkgroupWorkingArenaDestructor.
type Os_workgroup_working_arena_destructor_t = OSWorkgroupWorkingArenaDestructor

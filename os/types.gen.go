// Code generated from Apple documentation for os. DO NOT EDIT.

package os

// C struct types

// OSActivityScopeStateS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/os/os_activity_scope_state_s
type OSActivityScopeStateS struct {
	Opaque [2]uint64 // Opaque data that the system uses to store the execution state.

}

// Os_activity_scope_state_s is a type alias for OSActivityScopeStateS for use in objc.Send[T] calls.
type Os_activity_scope_state_s = OSActivityScopeStateS

// OSUnfairLockS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/os/os_unfair_lock_s
type OSUnfairLockS struct {
	_os_unfair_lock_opaque uint32
}

// Os_unfair_lock_s is a type alias for OSUnfairLockS for use in objc.Send[T] calls.
type Os_unfair_lock_s = OSUnfairLockS

// OSWorkgroupAttrOpaqueS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/os/os_workgroup_attr_opaque_s
type OSWorkgroupAttrOpaqueS struct {
	Sig    uint32
	Opaque [60]int8
}

// Os_workgroup_attr_opaque_s is a type alias for OSWorkgroupAttrOpaqueS for use in objc.Send[T] calls.
type Os_workgroup_attr_opaque_s = OSWorkgroupAttrOpaqueS

// OSWorkgroupIntervalDataOpaqueS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/os/os_workgroup_interval_data_opaque_s
type OSWorkgroupIntervalDataOpaqueS struct {
	Sig    uint32
	Opaque [56]int8
}

// Os_workgroup_interval_data_opaque_s is a type alias for OSWorkgroupIntervalDataOpaqueS for use in objc.Send[T] calls.
type Os_workgroup_interval_data_opaque_s = OSWorkgroupIntervalDataOpaqueS

// OSWorkgroupJoinTokenOpaqueS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/os/os_workgroup_join_token_opaque_s
type OSWorkgroupJoinTokenOpaqueS struct {
	Sig    uint32
	Opaque [36]int8
}

// Os_workgroup_join_token_opaque_s is a type alias for OSWorkgroupJoinTokenOpaqueS for use in objc.Send[T] calls.
type Os_workgroup_join_token_opaque_s = OSWorkgroupJoinTokenOpaqueS

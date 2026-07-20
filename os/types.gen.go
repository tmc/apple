// Code generated from Apple documentation for os. DO NOT EDIT.

package os

// C struct types

// Os_activity_scope_state_s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/os/os_activity_scope_state_s
type Os_activity_scope_state_s struct {
	Opaque uint64 // Opaque data that the system uses to store the execution state.

}

// Os_workgroup_attr_opaque_s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/os/os_workgroup_attr_opaque_s
type Os_workgroup_attr_opaque_s struct {
	Opaque int8
	Sig    uint32
}

// Os_workgroup_interval_data_opaque_s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/os/os_workgroup_interval_data_opaque_s
type Os_workgroup_interval_data_opaque_s struct {
	Opaque int8
	Sig    uint32
}

// Os_workgroup_join_token_opaque_s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/os/os_workgroup_join_token_opaque_s
type Os_workgroup_join_token_opaque_s struct {
	Opaque int8
	Sig    uint32
}

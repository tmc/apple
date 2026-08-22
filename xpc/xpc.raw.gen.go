// Code generated from Apple documentation for XPC. DO NOT EDIT.

package xpc

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

var rawSymbolLookupErrors = map[string]error{}

func registerRawFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		rawSymbolLookupErrors[name] = err
		return
	}
	defer func() {
		if r := recover(); r != nil {
			rawSymbolLookupErrors[name] = errRawRegisterPanic{name: name}
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

type errRawRegisterPanic struct {
	name string
}

func (e errRawRegisterPanic) Error() string {
	return "xpc: failed to register symbol " + e.name
}

func rawSymbolError(name string) error {
	return rawSymbolLookupErrors[name]
}

var rawfn_launch_activate_socket func(name string, fds **int32, cnt *uintptr) int32

func raw_launch_activate_socket(name string, fds **int32, cnt *uintptr) int32 {
	if rawfn_launch_activate_socket == nil {
		var zero int32
		return zero
	}
	return rawfn_launch_activate_socket(name, fds, cnt)
}

var rawfn_xpc_activity_copy_criteria func(activity unsafe.Pointer) unsafe.Pointer

func raw_xpc_activity_copy_criteria(activity unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_activity_copy_criteria == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_activity_copy_criteria(activity)
}

var rawfn_xpc_activity_get_state func(activity unsafe.Pointer) unsafe.Pointer

func raw_xpc_activity_get_state(activity unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_activity_get_state == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_activity_get_state(activity)
}

var rawfn_xpc_activity_register func(identifier string, criteria unsafe.Pointer, handler unsafe.Pointer)

func raw_xpc_activity_register(identifier string, criteria unsafe.Pointer, handler unsafe.Pointer) {
	if rawfn_xpc_activity_register == nil {
		return
	}
	rawfn_xpc_activity_register(identifier, criteria, handler)
}

var rawfn_xpc_activity_set_criteria func(activity unsafe.Pointer, criteria unsafe.Pointer)

func raw_xpc_activity_set_criteria(activity unsafe.Pointer, criteria unsafe.Pointer) {
	if rawfn_xpc_activity_set_criteria == nil {
		return
	}
	rawfn_xpc_activity_set_criteria(activity, criteria)
}

var rawfn_xpc_activity_set_state func(activity unsafe.Pointer, state unsafe.Pointer) bool

func raw_xpc_activity_set_state(activity unsafe.Pointer, state unsafe.Pointer) bool {
	if rawfn_xpc_activity_set_state == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_activity_set_state(activity, state)
}

var rawfn_xpc_activity_should_defer func(activity unsafe.Pointer) bool

func raw_xpc_activity_should_defer(activity unsafe.Pointer) bool {
	if rawfn_xpc_activity_should_defer == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_activity_should_defer(activity)
}

var rawfn_xpc_activity_unregister func(identifier string)

func raw_xpc_activity_unregister(identifier string) {
	if rawfn_xpc_activity_unregister == nil {
		return
	}
	rawfn_xpc_activity_unregister(identifier)
}

var rawfn_xpc_array_append_value func(xarray unsafe.Pointer, value unsafe.Pointer)

func raw_xpc_array_append_value(xarray unsafe.Pointer, value unsafe.Pointer) {
	if rawfn_xpc_array_append_value == nil {
		return
	}
	rawfn_xpc_array_append_value(xarray, value)
}

var rawfn_xpc_array_apply func(xarray unsafe.Pointer, applier unsafe.Pointer) bool

func raw_xpc_array_apply(xarray unsafe.Pointer, applier unsafe.Pointer) bool {
	if rawfn_xpc_array_apply == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_array_apply(xarray, applier)
}

var rawfn_xpc_array_create func(objects *unsafe.Pointer, count uintptr) unsafe.Pointer

func raw_xpc_array_create(objects *unsafe.Pointer, count uintptr) unsafe.Pointer {
	if rawfn_xpc_array_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_array_create(objects, count)
}

var rawfn_xpc_array_create_connection func(xarray unsafe.Pointer, index uintptr) unsafe.Pointer

func raw_xpc_array_create_connection(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	if rawfn_xpc_array_create_connection == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_array_create_connection(xarray, index)
}

var rawfn_xpc_array_create_empty func() unsafe.Pointer

func raw_xpc_array_create_empty() unsafe.Pointer {
	if rawfn_xpc_array_create_empty == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_array_create_empty()
}

var rawfn_xpc_array_dup_fd func(xarray unsafe.Pointer, index uintptr) int32

func raw_xpc_array_dup_fd(xarray unsafe.Pointer, index uintptr) int32 {
	if rawfn_xpc_array_dup_fd == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_array_dup_fd(xarray, index)
}

var rawfn_xpc_array_get_array func(xarray unsafe.Pointer, index uintptr) unsafe.Pointer

func raw_xpc_array_get_array(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	if rawfn_xpc_array_get_array == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_array_get_array(xarray, index)
}

var rawfn_xpc_array_get_bool func(xarray unsafe.Pointer, index uintptr) bool

func raw_xpc_array_get_bool(xarray unsafe.Pointer, index uintptr) bool {
	if rawfn_xpc_array_get_bool == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_array_get_bool(xarray, index)
}

var rawfn_xpc_array_get_count func(xarray unsafe.Pointer) uintptr

func raw_xpc_array_get_count(xarray unsafe.Pointer) uintptr {
	if rawfn_xpc_array_get_count == nil {
		var zero uintptr
		return zero
	}
	return rawfn_xpc_array_get_count(xarray)
}

var rawfn_xpc_array_get_data func(xarray unsafe.Pointer, index uintptr, length *uintptr) unsafe.Pointer

func raw_xpc_array_get_data(xarray unsafe.Pointer, index uintptr, length *uintptr) unsafe.Pointer {
	if rawfn_xpc_array_get_data == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_array_get_data(xarray, index, length)
}

var rawfn_xpc_array_get_date func(xarray unsafe.Pointer, index uintptr) int64

func raw_xpc_array_get_date(xarray unsafe.Pointer, index uintptr) int64 {
	if rawfn_xpc_array_get_date == nil {
		var zero int64
		return zero
	}
	return rawfn_xpc_array_get_date(xarray, index)
}

var rawfn_xpc_array_get_dictionary func(xarray unsafe.Pointer, index uintptr) unsafe.Pointer

func raw_xpc_array_get_dictionary(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	if rawfn_xpc_array_get_dictionary == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_array_get_dictionary(xarray, index)
}

var rawfn_xpc_array_get_double func(xarray unsafe.Pointer, index uintptr) float64

func raw_xpc_array_get_double(xarray unsafe.Pointer, index uintptr) float64 {
	if rawfn_xpc_array_get_double == nil {
		var zero float64
		return zero
	}
	return rawfn_xpc_array_get_double(xarray, index)
}

var rawfn_xpc_array_get_int64 func(xarray unsafe.Pointer, index uintptr) int64

func raw_xpc_array_get_int64(xarray unsafe.Pointer, index uintptr) int64 {
	if rawfn_xpc_array_get_int64 == nil {
		var zero int64
		return zero
	}
	return rawfn_xpc_array_get_int64(xarray, index)
}

var rawfn_xpc_array_get_string func(xarray unsafe.Pointer, index uintptr) *byte

func raw_xpc_array_get_string(xarray unsafe.Pointer, index uintptr) *byte {
	if rawfn_xpc_array_get_string == nil {
		var zero *byte
		return zero
	}
	return rawfn_xpc_array_get_string(xarray, index)
}

var rawfn_xpc_array_get_uint64 func(xarray unsafe.Pointer, index uintptr) uint64

func raw_xpc_array_get_uint64(xarray unsafe.Pointer, index uintptr) uint64 {
	if rawfn_xpc_array_get_uint64 == nil {
		var zero uint64
		return zero
	}
	return rawfn_xpc_array_get_uint64(xarray, index)
}

var rawfn_xpc_array_get_uuid func(xarray unsafe.Pointer, index uintptr) *uint8

func raw_xpc_array_get_uuid(xarray unsafe.Pointer, index uintptr) *uint8 {
	if rawfn_xpc_array_get_uuid == nil {
		var zero *uint8
		return zero
	}
	return rawfn_xpc_array_get_uuid(xarray, index)
}

var rawfn_xpc_array_get_value func(xarray unsafe.Pointer, index uintptr) unsafe.Pointer

func raw_xpc_array_get_value(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	if rawfn_xpc_array_get_value == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_array_get_value(xarray, index)
}

var rawfn_xpc_array_set_bool func(xarray unsafe.Pointer, index uintptr, value bool)

func raw_xpc_array_set_bool(xarray unsafe.Pointer, index uintptr, value bool) {
	if rawfn_xpc_array_set_bool == nil {
		return
	}
	rawfn_xpc_array_set_bool(xarray, index, value)
}

var rawfn_xpc_array_set_connection func(xarray unsafe.Pointer, index uintptr, connection unsafe.Pointer)

func raw_xpc_array_set_connection(xarray unsafe.Pointer, index uintptr, connection unsafe.Pointer) {
	if rawfn_xpc_array_set_connection == nil {
		return
	}
	rawfn_xpc_array_set_connection(xarray, index, connection)
}

var rawfn_xpc_array_set_data func(xarray unsafe.Pointer, index uintptr, bytes unsafe.Pointer, length uintptr)

func raw_xpc_array_set_data(xarray unsafe.Pointer, index uintptr, bytes unsafe.Pointer, length uintptr) {
	if rawfn_xpc_array_set_data == nil {
		return
	}
	rawfn_xpc_array_set_data(xarray, index, bytes, length)
}

var rawfn_xpc_array_set_date func(xarray unsafe.Pointer, index uintptr, value int64)

func raw_xpc_array_set_date(xarray unsafe.Pointer, index uintptr, value int64) {
	if rawfn_xpc_array_set_date == nil {
		return
	}
	rawfn_xpc_array_set_date(xarray, index, value)
}

var rawfn_xpc_array_set_double func(xarray unsafe.Pointer, index uintptr, value float64)

func raw_xpc_array_set_double(xarray unsafe.Pointer, index uintptr, value float64) {
	if rawfn_xpc_array_set_double == nil {
		return
	}
	rawfn_xpc_array_set_double(xarray, index, value)
}

var rawfn_xpc_array_set_fd func(xarray unsafe.Pointer, index uintptr, fd int32)

func raw_xpc_array_set_fd(xarray unsafe.Pointer, index uintptr, fd int32) {
	if rawfn_xpc_array_set_fd == nil {
		return
	}
	rawfn_xpc_array_set_fd(xarray, index, fd)
}

var rawfn_xpc_array_set_int64 func(xarray unsafe.Pointer, index uintptr, value int64)

func raw_xpc_array_set_int64(xarray unsafe.Pointer, index uintptr, value int64) {
	if rawfn_xpc_array_set_int64 == nil {
		return
	}
	rawfn_xpc_array_set_int64(xarray, index, value)
}

var rawfn_xpc_array_set_string func(xarray unsafe.Pointer, index uintptr, string_ string)

func raw_xpc_array_set_string(xarray unsafe.Pointer, index uintptr, string_ string) {
	if rawfn_xpc_array_set_string == nil {
		return
	}
	rawfn_xpc_array_set_string(xarray, index, string_)
}

var rawfn_xpc_array_set_uint64 func(xarray unsafe.Pointer, index uintptr, value uint64)

func raw_xpc_array_set_uint64(xarray unsafe.Pointer, index uintptr, value uint64) {
	if rawfn_xpc_array_set_uint64 == nil {
		return
	}
	rawfn_xpc_array_set_uint64(xarray, index, value)
}

var rawfn_xpc_array_set_uuid func(xarray unsafe.Pointer, index uintptr, uuid *[16]byte)

func raw_xpc_array_set_uuid(xarray unsafe.Pointer, index uintptr, uuid [16]byte) {
	if rawfn_xpc_array_set_uuid == nil {
		return
	}
	rawfn_xpc_array_set_uuid(xarray, index, &uuid)
}

var rawfn_xpc_array_set_value func(xarray unsafe.Pointer, index uintptr, value unsafe.Pointer)

func raw_xpc_array_set_value(xarray unsafe.Pointer, index uintptr, value unsafe.Pointer) {
	if rawfn_xpc_array_set_value == nil {
		return
	}
	rawfn_xpc_array_set_value(xarray, index, value)
}

var rawfn_xpc_bool_create func(value bool) unsafe.Pointer

func raw_xpc_bool_create(value bool) unsafe.Pointer {
	if rawfn_xpc_bool_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_bool_create(value)
}

var rawfn_xpc_bool_get_value func(xbool unsafe.Pointer) bool

func raw_xpc_bool_get_value(xbool unsafe.Pointer) bool {
	if rawfn_xpc_bool_get_value == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_bool_get_value(xbool)
}

var rawfn_xpc_connection_activate func(connection unsafe.Pointer)

func raw_xpc_connection_activate(connection unsafe.Pointer) {
	if rawfn_xpc_connection_activate == nil {
		return
	}
	rawfn_xpc_connection_activate(connection)
}

var rawfn_xpc_connection_cancel func(connection unsafe.Pointer)

func raw_xpc_connection_cancel(connection unsafe.Pointer) {
	if rawfn_xpc_connection_cancel == nil {
		return
	}
	rawfn_xpc_connection_cancel(connection)
}

var rawfn_xpc_connection_copy_invalidation_reason func(connection unsafe.Pointer) *byte

func raw_xpc_connection_copy_invalidation_reason(connection unsafe.Pointer) *byte {
	if rawfn_xpc_connection_copy_invalidation_reason == nil {
		var zero *byte
		return zero
	}
	return rawfn_xpc_connection_copy_invalidation_reason(connection)
}

var rawfn_xpc_connection_create func(name string, targetq unsafe.Pointer) unsafe.Pointer

func raw_xpc_connection_create(name string, targetq unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_connection_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_connection_create(name, targetq)
}

var rawfn_xpc_connection_create_from_endpoint func(endpoint unsafe.Pointer) unsafe.Pointer

func raw_xpc_connection_create_from_endpoint(endpoint unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_connection_create_from_endpoint == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_connection_create_from_endpoint(endpoint)
}

var rawfn_xpc_connection_create_mach_service func(name string, targetq unsafe.Pointer, flags uint64) unsafe.Pointer

func raw_xpc_connection_create_mach_service(name string, targetq unsafe.Pointer, flags uint64) unsafe.Pointer {
	if rawfn_xpc_connection_create_mach_service == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_connection_create_mach_service(name, targetq, flags)
}

var rawfn_xpc_connection_get_asid func(connection unsafe.Pointer) int32

func raw_xpc_connection_get_asid(connection unsafe.Pointer) int32 {
	if rawfn_xpc_connection_get_asid == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_connection_get_asid(connection)
}

var rawfn_xpc_connection_get_context func(connection unsafe.Pointer) unsafe.Pointer

func raw_xpc_connection_get_context(connection unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_connection_get_context == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_connection_get_context(connection)
}

var rawfn_xpc_connection_get_egid func(connection unsafe.Pointer) uint32

func raw_xpc_connection_get_egid(connection unsafe.Pointer) uint32 {
	if rawfn_xpc_connection_get_egid == nil {
		var zero uint32
		return zero
	}
	return rawfn_xpc_connection_get_egid(connection)
}

var rawfn_xpc_connection_get_euid func(connection unsafe.Pointer) uint32

func raw_xpc_connection_get_euid(connection unsafe.Pointer) uint32 {
	if rawfn_xpc_connection_get_euid == nil {
		var zero uint32
		return zero
	}
	return rawfn_xpc_connection_get_euid(connection)
}

var rawfn_xpc_connection_get_name func(connection unsafe.Pointer) *byte

func raw_xpc_connection_get_name(connection unsafe.Pointer) *byte {
	if rawfn_xpc_connection_get_name == nil {
		var zero *byte
		return zero
	}
	return rawfn_xpc_connection_get_name(connection)
}

var rawfn_xpc_connection_get_pid func(connection unsafe.Pointer) int32

func raw_xpc_connection_get_pid(connection unsafe.Pointer) int32 {
	if rawfn_xpc_connection_get_pid == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_connection_get_pid(connection)
}

var rawfn_xpc_connection_resume func(connection unsafe.Pointer)

func raw_xpc_connection_resume(connection unsafe.Pointer) {
	if rawfn_xpc_connection_resume == nil {
		return
	}
	rawfn_xpc_connection_resume(connection)
}

var rawfn_xpc_connection_send_barrier func(connection unsafe.Pointer, barrier unsafe.Pointer)

func raw_xpc_connection_send_barrier(connection unsafe.Pointer, barrier unsafe.Pointer) {
	if rawfn_xpc_connection_send_barrier == nil {
		return
	}
	rawfn_xpc_connection_send_barrier(connection, barrier)
}

var rawfn_xpc_connection_send_message func(connection unsafe.Pointer, message unsafe.Pointer)

func raw_xpc_connection_send_message(connection unsafe.Pointer, message unsafe.Pointer) {
	if rawfn_xpc_connection_send_message == nil {
		return
	}
	rawfn_xpc_connection_send_message(connection, message)
}

var rawfn_xpc_connection_send_message_with_reply func(connection unsafe.Pointer, message unsafe.Pointer, replyq unsafe.Pointer, handler unsafe.Pointer)

func raw_xpc_connection_send_message_with_reply(connection unsafe.Pointer, message unsafe.Pointer, replyq unsafe.Pointer, handler unsafe.Pointer) {
	if rawfn_xpc_connection_send_message_with_reply == nil {
		return
	}
	rawfn_xpc_connection_send_message_with_reply(connection, message, replyq, handler)
}

var rawfn_xpc_connection_send_message_with_reply_sync func(connection unsafe.Pointer, message unsafe.Pointer) unsafe.Pointer

func raw_xpc_connection_send_message_with_reply_sync(connection unsafe.Pointer, message unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_connection_send_message_with_reply_sync == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_connection_send_message_with_reply_sync(connection, message)
}

var rawfn_xpc_connection_set_context func(connection unsafe.Pointer, context unsafe.Pointer)

func raw_xpc_connection_set_context(connection unsafe.Pointer, context unsafe.Pointer) {
	if rawfn_xpc_connection_set_context == nil {
		return
	}
	rawfn_xpc_connection_set_context(connection, context)
}

var rawfn_xpc_connection_set_event_handler func(connection unsafe.Pointer, handler unsafe.Pointer)

func raw_xpc_connection_set_event_handler(connection unsafe.Pointer, handler unsafe.Pointer) {
	if rawfn_xpc_connection_set_event_handler == nil {
		return
	}
	rawfn_xpc_connection_set_event_handler(connection, handler)
}

var rawfn_xpc_connection_set_finalizer_f func(connection unsafe.Pointer, finalizer unsafe.Pointer)

func raw_xpc_connection_set_finalizer_f(connection unsafe.Pointer, finalizer unsafe.Pointer) {
	if rawfn_xpc_connection_set_finalizer_f == nil {
		return
	}
	rawfn_xpc_connection_set_finalizer_f(connection, finalizer)
}

var rawfn_xpc_connection_set_peer_code_signing_requirement func(connection unsafe.Pointer, requirement string) int32

func raw_xpc_connection_set_peer_code_signing_requirement(connection unsafe.Pointer, requirement string) int32 {
	if rawfn_xpc_connection_set_peer_code_signing_requirement == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_connection_set_peer_code_signing_requirement(connection, requirement)
}

var rawfn_xpc_connection_set_peer_entitlement_exists_requirement func(connection unsafe.Pointer, entitlement string) int32

func raw_xpc_connection_set_peer_entitlement_exists_requirement(connection unsafe.Pointer, entitlement string) int32 {
	if rawfn_xpc_connection_set_peer_entitlement_exists_requirement == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_connection_set_peer_entitlement_exists_requirement(connection, entitlement)
}

var rawfn_xpc_connection_set_peer_entitlement_matches_value_requirement func(connection unsafe.Pointer, entitlement string, value unsafe.Pointer) int32

func raw_xpc_connection_set_peer_entitlement_matches_value_requirement(connection unsafe.Pointer, entitlement string, value unsafe.Pointer) int32 {
	if rawfn_xpc_connection_set_peer_entitlement_matches_value_requirement == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_connection_set_peer_entitlement_matches_value_requirement(connection, entitlement, value)
}

var rawfn_xpc_connection_set_peer_lightweight_code_requirement func(connection unsafe.Pointer, lwcr unsafe.Pointer) int32

func raw_xpc_connection_set_peer_lightweight_code_requirement(connection unsafe.Pointer, lwcr unsafe.Pointer) int32 {
	if rawfn_xpc_connection_set_peer_lightweight_code_requirement == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_connection_set_peer_lightweight_code_requirement(connection, lwcr)
}

var rawfn_xpc_connection_set_peer_platform_identity_requirement func(connection unsafe.Pointer, signing_identifier string) int32

func raw_xpc_connection_set_peer_platform_identity_requirement(connection unsafe.Pointer, signing_identifier string) int32 {
	if rawfn_xpc_connection_set_peer_platform_identity_requirement == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_connection_set_peer_platform_identity_requirement(connection, signing_identifier)
}

var rawfn_xpc_connection_set_peer_requirement func(connection unsafe.Pointer, peer_requirement unsafe.Pointer)

func raw_xpc_connection_set_peer_requirement(connection unsafe.Pointer, peer_requirement unsafe.Pointer) {
	if rawfn_xpc_connection_set_peer_requirement == nil {
		return
	}
	rawfn_xpc_connection_set_peer_requirement(connection, peer_requirement)
}

var rawfn_xpc_connection_set_peer_team_identity_requirement func(connection unsafe.Pointer, signing_identifier string) int32

func raw_xpc_connection_set_peer_team_identity_requirement(connection unsafe.Pointer, signing_identifier string) int32 {
	if rawfn_xpc_connection_set_peer_team_identity_requirement == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_connection_set_peer_team_identity_requirement(connection, signing_identifier)
}

var rawfn_xpc_connection_set_target_queue func(connection unsafe.Pointer, targetq unsafe.Pointer)

func raw_xpc_connection_set_target_queue(connection unsafe.Pointer, targetq unsafe.Pointer) {
	if rawfn_xpc_connection_set_target_queue == nil {
		return
	}
	rawfn_xpc_connection_set_target_queue(connection, targetq)
}

var rawfn_xpc_connection_suspend func(connection unsafe.Pointer)

func raw_xpc_connection_suspend(connection unsafe.Pointer) {
	if rawfn_xpc_connection_suspend == nil {
		return
	}
	rawfn_xpc_connection_suspend(connection)
}

var rawfn_xpc_copy func(object unsafe.Pointer) unsafe.Pointer

func raw_xpc_copy(object unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_copy == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_copy(object)
}

var rawfn_xpc_copy_description func(object unsafe.Pointer) *byte

func raw_xpc_copy_description(object unsafe.Pointer) *byte {
	if rawfn_xpc_copy_description == nil {
		var zero *byte
		return zero
	}
	return rawfn_xpc_copy_description(object)
}

var rawfn_xpc_data_create func(bytes unsafe.Pointer, length uintptr) unsafe.Pointer

func raw_xpc_data_create(bytes unsafe.Pointer, length uintptr) unsafe.Pointer {
	if rawfn_xpc_data_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_data_create(bytes, length)
}

var rawfn_xpc_data_create_with_dispatch_data func(ddata unsafe.Pointer) unsafe.Pointer

func raw_xpc_data_create_with_dispatch_data(ddata unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_data_create_with_dispatch_data == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_data_create_with_dispatch_data(ddata)
}

var rawfn_xpc_data_get_bytes func(xdata unsafe.Pointer, buffer unsafe.Pointer, off uintptr, length uintptr) uintptr

func raw_xpc_data_get_bytes(xdata unsafe.Pointer, buffer unsafe.Pointer, off uintptr, length uintptr) uintptr {
	if rawfn_xpc_data_get_bytes == nil {
		var zero uintptr
		return zero
	}
	return rawfn_xpc_data_get_bytes(xdata, buffer, off, length)
}

var rawfn_xpc_data_get_bytes_ptr func(xdata unsafe.Pointer) unsafe.Pointer

func raw_xpc_data_get_bytes_ptr(xdata unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_data_get_bytes_ptr == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_data_get_bytes_ptr(xdata)
}

var rawfn_xpc_data_get_length func(xdata unsafe.Pointer) uintptr

func raw_xpc_data_get_length(xdata unsafe.Pointer) uintptr {
	if rawfn_xpc_data_get_length == nil {
		var zero uintptr
		return zero
	}
	return rawfn_xpc_data_get_length(xdata)
}

var rawfn_xpc_date_create func(interval int64) unsafe.Pointer

func raw_xpc_date_create(interval int64) unsafe.Pointer {
	if rawfn_xpc_date_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_date_create(interval)
}

var rawfn_xpc_date_create_from_current func() unsafe.Pointer

func raw_xpc_date_create_from_current() unsafe.Pointer {
	if rawfn_xpc_date_create_from_current == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_date_create_from_current()
}

var rawfn_xpc_date_get_value func(xdate unsafe.Pointer) int64

func raw_xpc_date_get_value(xdate unsafe.Pointer) int64 {
	if rawfn_xpc_date_get_value == nil {
		var zero int64
		return zero
	}
	return rawfn_xpc_date_get_value(xdate)
}

var rawfn_xpc_dictionary_apply func(xdict unsafe.Pointer, applier unsafe.Pointer) bool

func raw_xpc_dictionary_apply(xdict unsafe.Pointer, applier unsafe.Pointer) bool {
	if rawfn_xpc_dictionary_apply == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_dictionary_apply(xdict, applier)
}

var rawfn_xpc_dictionary_copy_mach_send func(xdict unsafe.Pointer, key string) uint32

func raw_xpc_dictionary_copy_mach_send(xdict unsafe.Pointer, key string) uint32 {
	if rawfn_xpc_dictionary_copy_mach_send == nil {
		var zero uint32
		return zero
	}
	return rawfn_xpc_dictionary_copy_mach_send(xdict, key)
}

var rawfn_xpc_dictionary_create func(keys string, values *unsafe.Pointer, count uintptr) unsafe.Pointer

func raw_xpc_dictionary_create(keys string, values *unsafe.Pointer, count uintptr) unsafe.Pointer {
	if rawfn_xpc_dictionary_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_dictionary_create(keys, values, count)
}

var rawfn_xpc_dictionary_create_connection func(xdict unsafe.Pointer, key string) unsafe.Pointer

func raw_xpc_dictionary_create_connection(xdict unsafe.Pointer, key string) unsafe.Pointer {
	if rawfn_xpc_dictionary_create_connection == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_dictionary_create_connection(xdict, key)
}

var rawfn_xpc_dictionary_create_empty func() unsafe.Pointer

func raw_xpc_dictionary_create_empty() unsafe.Pointer {
	if rawfn_xpc_dictionary_create_empty == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_dictionary_create_empty()
}

var rawfn_xpc_dictionary_create_reply func(original unsafe.Pointer) unsafe.Pointer

func raw_xpc_dictionary_create_reply(original unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_dictionary_create_reply == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_dictionary_create_reply(original)
}

var rawfn_xpc_dictionary_dup_fd func(xdict unsafe.Pointer, key string) int32

func raw_xpc_dictionary_dup_fd(xdict unsafe.Pointer, key string) int32 {
	if rawfn_xpc_dictionary_dup_fd == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_dictionary_dup_fd(xdict, key)
}

var rawfn_xpc_dictionary_get_array func(xdict unsafe.Pointer, key string) unsafe.Pointer

func raw_xpc_dictionary_get_array(xdict unsafe.Pointer, key string) unsafe.Pointer {
	if rawfn_xpc_dictionary_get_array == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_dictionary_get_array(xdict, key)
}

var rawfn_xpc_dictionary_get_bool func(xdict unsafe.Pointer, key string) bool

func raw_xpc_dictionary_get_bool(xdict unsafe.Pointer, key string) bool {
	if rawfn_xpc_dictionary_get_bool == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_dictionary_get_bool(xdict, key)
}

var rawfn_xpc_dictionary_get_count func(xdict unsafe.Pointer) uintptr

func raw_xpc_dictionary_get_count(xdict unsafe.Pointer) uintptr {
	if rawfn_xpc_dictionary_get_count == nil {
		var zero uintptr
		return zero
	}
	return rawfn_xpc_dictionary_get_count(xdict)
}

var rawfn_xpc_dictionary_get_data func(xdict unsafe.Pointer, key string, length *uintptr) unsafe.Pointer

func raw_xpc_dictionary_get_data(xdict unsafe.Pointer, key string, length *uintptr) unsafe.Pointer {
	if rawfn_xpc_dictionary_get_data == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_dictionary_get_data(xdict, key, length)
}

var rawfn_xpc_dictionary_get_date func(xdict unsafe.Pointer, key string) int64

func raw_xpc_dictionary_get_date(xdict unsafe.Pointer, key string) int64 {
	if rawfn_xpc_dictionary_get_date == nil {
		var zero int64
		return zero
	}
	return rawfn_xpc_dictionary_get_date(xdict, key)
}

var rawfn_xpc_dictionary_get_dictionary func(xdict unsafe.Pointer, key string) unsafe.Pointer

func raw_xpc_dictionary_get_dictionary(xdict unsafe.Pointer, key string) unsafe.Pointer {
	if rawfn_xpc_dictionary_get_dictionary == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_dictionary_get_dictionary(xdict, key)
}

var rawfn_xpc_dictionary_get_double func(xdict unsafe.Pointer, key string) float64

func raw_xpc_dictionary_get_double(xdict unsafe.Pointer, key string) float64 {
	if rawfn_xpc_dictionary_get_double == nil {
		var zero float64
		return zero
	}
	return rawfn_xpc_dictionary_get_double(xdict, key)
}

var rawfn_xpc_dictionary_get_int64 func(xdict unsafe.Pointer, key string) int64

func raw_xpc_dictionary_get_int64(xdict unsafe.Pointer, key string) int64 {
	if rawfn_xpc_dictionary_get_int64 == nil {
		var zero int64
		return zero
	}
	return rawfn_xpc_dictionary_get_int64(xdict, key)
}

var rawfn_xpc_dictionary_get_remote_connection func(xdict unsafe.Pointer) unsafe.Pointer

func raw_xpc_dictionary_get_remote_connection(xdict unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_dictionary_get_remote_connection == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_dictionary_get_remote_connection(xdict)
}

var rawfn_xpc_dictionary_get_string func(xdict unsafe.Pointer, key string) *byte

func raw_xpc_dictionary_get_string(xdict unsafe.Pointer, key string) *byte {
	if rawfn_xpc_dictionary_get_string == nil {
		var zero *byte
		return zero
	}
	return rawfn_xpc_dictionary_get_string(xdict, key)
}

var rawfn_xpc_dictionary_get_uint64 func(xdict unsafe.Pointer, key string) uint64

func raw_xpc_dictionary_get_uint64(xdict unsafe.Pointer, key string) uint64 {
	if rawfn_xpc_dictionary_get_uint64 == nil {
		var zero uint64
		return zero
	}
	return rawfn_xpc_dictionary_get_uint64(xdict, key)
}

var rawfn_xpc_dictionary_get_uuid func(xdict unsafe.Pointer, key string) *uint8

func raw_xpc_dictionary_get_uuid(xdict unsafe.Pointer, key string) *uint8 {
	if rawfn_xpc_dictionary_get_uuid == nil {
		var zero *uint8
		return zero
	}
	return rawfn_xpc_dictionary_get_uuid(xdict, key)
}

var rawfn_xpc_dictionary_get_value func(xdict unsafe.Pointer, key string) unsafe.Pointer

func raw_xpc_dictionary_get_value(xdict unsafe.Pointer, key string) unsafe.Pointer {
	if rawfn_xpc_dictionary_get_value == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_dictionary_get_value(xdict, key)
}

var rawfn_xpc_dictionary_set_bool func(xdict unsafe.Pointer, key string, value bool)

func raw_xpc_dictionary_set_bool(xdict unsafe.Pointer, key string, value bool) {
	if rawfn_xpc_dictionary_set_bool == nil {
		return
	}
	rawfn_xpc_dictionary_set_bool(xdict, key, value)
}

var rawfn_xpc_dictionary_set_connection func(xdict unsafe.Pointer, key string, connection unsafe.Pointer)

func raw_xpc_dictionary_set_connection(xdict unsafe.Pointer, key string, connection unsafe.Pointer) {
	if rawfn_xpc_dictionary_set_connection == nil {
		return
	}
	rawfn_xpc_dictionary_set_connection(xdict, key, connection)
}

var rawfn_xpc_dictionary_set_data func(xdict unsafe.Pointer, key string, bytes unsafe.Pointer, length uintptr)

func raw_xpc_dictionary_set_data(xdict unsafe.Pointer, key string, bytes unsafe.Pointer, length uintptr) {
	if rawfn_xpc_dictionary_set_data == nil {
		return
	}
	rawfn_xpc_dictionary_set_data(xdict, key, bytes, length)
}

var rawfn_xpc_dictionary_set_date func(xdict unsafe.Pointer, key string, value int64)

func raw_xpc_dictionary_set_date(xdict unsafe.Pointer, key string, value int64) {
	if rawfn_xpc_dictionary_set_date == nil {
		return
	}
	rawfn_xpc_dictionary_set_date(xdict, key, value)
}

var rawfn_xpc_dictionary_set_double func(xdict unsafe.Pointer, key string, value float64)

func raw_xpc_dictionary_set_double(xdict unsafe.Pointer, key string, value float64) {
	if rawfn_xpc_dictionary_set_double == nil {
		return
	}
	rawfn_xpc_dictionary_set_double(xdict, key, value)
}

var rawfn_xpc_dictionary_set_fd func(xdict unsafe.Pointer, key string, fd int32)

func raw_xpc_dictionary_set_fd(xdict unsafe.Pointer, key string, fd int32) {
	if rawfn_xpc_dictionary_set_fd == nil {
		return
	}
	rawfn_xpc_dictionary_set_fd(xdict, key, fd)
}

var rawfn_xpc_dictionary_set_int64 func(xdict unsafe.Pointer, key string, value int64)

func raw_xpc_dictionary_set_int64(xdict unsafe.Pointer, key string, value int64) {
	if rawfn_xpc_dictionary_set_int64 == nil {
		return
	}
	rawfn_xpc_dictionary_set_int64(xdict, key, value)
}

var rawfn_xpc_dictionary_set_mach_send func(xdict unsafe.Pointer, key string, p uint32)

func raw_xpc_dictionary_set_mach_send(xdict unsafe.Pointer, key string, p uint32) {
	if rawfn_xpc_dictionary_set_mach_send == nil {
		return
	}
	rawfn_xpc_dictionary_set_mach_send(xdict, key, p)
}

var rawfn_xpc_dictionary_set_string func(xdict unsafe.Pointer, key string, string_ string)

func raw_xpc_dictionary_set_string(xdict unsafe.Pointer, key string, string_ string) {
	if rawfn_xpc_dictionary_set_string == nil {
		return
	}
	rawfn_xpc_dictionary_set_string(xdict, key, string_)
}

var rawfn_xpc_dictionary_set_uint64 func(xdict unsafe.Pointer, key string, value uint64)

func raw_xpc_dictionary_set_uint64(xdict unsafe.Pointer, key string, value uint64) {
	if rawfn_xpc_dictionary_set_uint64 == nil {
		return
	}
	rawfn_xpc_dictionary_set_uint64(xdict, key, value)
}

var rawfn_xpc_dictionary_set_uuid func(xdict unsafe.Pointer, key string, uuid *[16]byte)

func raw_xpc_dictionary_set_uuid(xdict unsafe.Pointer, key string, uuid [16]byte) {
	if rawfn_xpc_dictionary_set_uuid == nil {
		return
	}
	rawfn_xpc_dictionary_set_uuid(xdict, key, &uuid)
}

var rawfn_xpc_dictionary_set_value func(xdict unsafe.Pointer, key string, value unsafe.Pointer)

func raw_xpc_dictionary_set_value(xdict unsafe.Pointer, key string, value unsafe.Pointer) {
	if rawfn_xpc_dictionary_set_value == nil {
		return
	}
	rawfn_xpc_dictionary_set_value(xdict, key, value)
}

var rawfn_xpc_double_create func(value float64) unsafe.Pointer

func raw_xpc_double_create(value float64) unsafe.Pointer {
	if rawfn_xpc_double_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_double_create(value)
}

var rawfn_xpc_double_get_value func(xdouble unsafe.Pointer) float64

func raw_xpc_double_get_value(xdouble unsafe.Pointer) float64 {
	if rawfn_xpc_double_get_value == nil {
		var zero float64
		return zero
	}
	return rawfn_xpc_double_get_value(xdouble)
}

var rawfn_xpc_endpoint_create func(connection unsafe.Pointer) unsafe.Pointer

func raw_xpc_endpoint_create(connection unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_endpoint_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_endpoint_create(connection)
}

var rawfn_xpc_equal func(object1 unsafe.Pointer, object2 unsafe.Pointer) bool

func raw_xpc_equal(object1 unsafe.Pointer, object2 unsafe.Pointer) bool {
	if rawfn_xpc_equal == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_equal(object1, object2)
}

var rawfn_xpc_fd_create func(fd int32) unsafe.Pointer

func raw_xpc_fd_create(fd int32) unsafe.Pointer {
	if rawfn_xpc_fd_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_fd_create(fd)
}

var rawfn_xpc_fd_dup func(xfd unsafe.Pointer) int32

func raw_xpc_fd_dup(xfd unsafe.Pointer) int32 {
	if rawfn_xpc_fd_dup == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_fd_dup(xfd)
}

var rawfn_xpc_get_type func(object unsafe.Pointer) unsafe.Pointer

func raw_xpc_get_type(object unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_get_type == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_get_type(object)
}

var rawfn_xpc_hash func(object unsafe.Pointer) uintptr

func raw_xpc_hash(object unsafe.Pointer) uintptr {
	if rawfn_xpc_hash == nil {
		var zero uintptr
		return zero
	}
	return rawfn_xpc_hash(object)
}

var rawfn_xpc_int64_create func(value int64) unsafe.Pointer

func raw_xpc_int64_create(value int64) unsafe.Pointer {
	if rawfn_xpc_int64_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_int64_create(value)
}

var rawfn_xpc_int64_get_value func(xint unsafe.Pointer) int64

func raw_xpc_int64_get_value(xint unsafe.Pointer) int64 {
	if rawfn_xpc_int64_get_value == nil {
		var zero int64
		return zero
	}
	return rawfn_xpc_int64_get_value(xint)
}

var rawfn_xpc_listener_activate func(listener unsafe.Pointer, error_out *unsafe.Pointer) bool

func raw_xpc_listener_activate(listener unsafe.Pointer, error_out *unsafe.Pointer) bool {
	if rawfn_xpc_listener_activate == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_listener_activate(listener, error_out)
}

var rawfn_xpc_listener_cancel func(listener unsafe.Pointer)

func raw_xpc_listener_cancel(listener unsafe.Pointer) {
	if rawfn_xpc_listener_cancel == nil {
		return
	}
	rawfn_xpc_listener_cancel(listener)
}

var rawfn_xpc_listener_copy_description func(listener unsafe.Pointer) *byte

func raw_xpc_listener_copy_description(listener unsafe.Pointer) *byte {
	if rawfn_xpc_listener_copy_description == nil {
		var zero *byte
		return zero
	}
	return rawfn_xpc_listener_copy_description(listener)
}

var rawfn_xpc_listener_create func(service string, target_queue unsafe.Pointer, flags uint64, incoming_session_handler unsafe.Pointer, error_out *unsafe.Pointer) unsafe.Pointer

func raw_xpc_listener_create(service string, target_queue unsafe.Pointer, flags uint64, incoming_session_handler unsafe.Pointer, error_out *unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_listener_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_listener_create(service, target_queue, flags, incoming_session_handler, error_out)
}

var rawfn_xpc_listener_reject_peer func(peer unsafe.Pointer, reason string)

func raw_xpc_listener_reject_peer(peer unsafe.Pointer, reason string) {
	if rawfn_xpc_listener_reject_peer == nil {
		return
	}
	rawfn_xpc_listener_reject_peer(peer, reason)
}

var rawfn_xpc_listener_set_peer_code_signing_requirement func(listener unsafe.Pointer, requirement string) int32

func raw_xpc_listener_set_peer_code_signing_requirement(listener unsafe.Pointer, requirement string) int32 {
	if rawfn_xpc_listener_set_peer_code_signing_requirement == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_listener_set_peer_code_signing_requirement(listener, requirement)
}

var rawfn_xpc_listener_set_peer_requirement func(listener unsafe.Pointer, requirement unsafe.Pointer)

func raw_xpc_listener_set_peer_requirement(listener unsafe.Pointer, requirement unsafe.Pointer) {
	if rawfn_xpc_listener_set_peer_requirement == nil {
		return
	}
	rawfn_xpc_listener_set_peer_requirement(listener, requirement)
}

var rawfn_xpc_main func(handler unsafe.Pointer)

func raw_xpc_main(handler unsafe.Pointer) {
	if rawfn_xpc_main == nil {
		return
	}
	rawfn_xpc_main(handler)
}

var rawfn_xpc_null_create func() unsafe.Pointer

func raw_xpc_null_create() unsafe.Pointer {
	if rawfn_xpc_null_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_null_create()
}

var rawfn_xpc_peer_requirement_create_entitlement_exists func(entitlement string, error_out *unsafe.Pointer) unsafe.Pointer

func raw_xpc_peer_requirement_create_entitlement_exists(entitlement string, error_out *unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_peer_requirement_create_entitlement_exists == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_peer_requirement_create_entitlement_exists(entitlement, error_out)
}

var rawfn_xpc_peer_requirement_create_entitlement_matches_value func(entitlement string, value unsafe.Pointer, error_out *unsafe.Pointer) unsafe.Pointer

func raw_xpc_peer_requirement_create_entitlement_matches_value(entitlement string, value unsafe.Pointer, error_out *unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_peer_requirement_create_entitlement_matches_value == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_peer_requirement_create_entitlement_matches_value(entitlement, value, error_out)
}

var rawfn_xpc_peer_requirement_create_lwcr func(lwcr unsafe.Pointer, error_out *unsafe.Pointer) unsafe.Pointer

func raw_xpc_peer_requirement_create_lwcr(lwcr unsafe.Pointer, error_out *unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_peer_requirement_create_lwcr == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_peer_requirement_create_lwcr(lwcr, error_out)
}

var rawfn_xpc_peer_requirement_create_platform_identity func(signing_identifier *byte, error_out *unsafe.Pointer) unsafe.Pointer

func raw_xpc_peer_requirement_create_platform_identity(signing_identifier *byte, error_out *unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_peer_requirement_create_platform_identity == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_peer_requirement_create_platform_identity(signing_identifier, error_out)
}

var rawfn_xpc_peer_requirement_create_team_identity func(signing_identifier *byte, error_out *unsafe.Pointer) unsafe.Pointer

func raw_xpc_peer_requirement_create_team_identity(signing_identifier *byte, error_out *unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_peer_requirement_create_team_identity == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_peer_requirement_create_team_identity(signing_identifier, error_out)
}

var rawfn_xpc_peer_requirement_match_received_message func(peer_requirement unsafe.Pointer, message unsafe.Pointer, error_out *unsafe.Pointer) bool

func raw_xpc_peer_requirement_match_received_message(peer_requirement unsafe.Pointer, message unsafe.Pointer, error_out *unsafe.Pointer) bool {
	if rawfn_xpc_peer_requirement_match_received_message == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_peer_requirement_match_received_message(peer_requirement, message, error_out)
}

var rawfn_xpc_release func(object unsafe.Pointer)

func raw_xpc_release(object unsafe.Pointer) {
	if rawfn_xpc_release == nil {
		return
	}
	rawfn_xpc_release(object)
}

var rawfn_xpc_retain func(object unsafe.Pointer) unsafe.Pointer

func raw_xpc_retain(object unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_retain == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_retain(object)
}

var rawfn_xpc_rich_error_can_retry func(err unsafe.Pointer) bool

func raw_xpc_rich_error_can_retry(err unsafe.Pointer) bool {
	if rawfn_xpc_rich_error_can_retry == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_rich_error_can_retry(err)
}

var rawfn_xpc_rich_error_copy_description func(err unsafe.Pointer) *byte

func raw_xpc_rich_error_copy_description(err unsafe.Pointer) *byte {
	if rawfn_xpc_rich_error_copy_description == nil {
		var zero *byte
		return zero
	}
	return rawfn_xpc_rich_error_copy_description(err)
}

var rawfn_xpc_session_activate func(session unsafe.Pointer, error_out *unsafe.Pointer) bool

func raw_xpc_session_activate(session unsafe.Pointer, error_out *unsafe.Pointer) bool {
	if rawfn_xpc_session_activate == nil {
		var zero bool
		return zero
	}
	return rawfn_xpc_session_activate(session, error_out)
}

var rawfn_xpc_session_cancel func(session unsafe.Pointer)

func raw_xpc_session_cancel(session unsafe.Pointer) {
	if rawfn_xpc_session_cancel == nil {
		return
	}
	rawfn_xpc_session_cancel(session)
}

var rawfn_xpc_session_copy_description func(session unsafe.Pointer) *byte

func raw_xpc_session_copy_description(session unsafe.Pointer) *byte {
	if rawfn_xpc_session_copy_description == nil {
		var zero *byte
		return zero
	}
	return rawfn_xpc_session_copy_description(session)
}

var rawfn_xpc_session_create_mach_service func(mach_service string, target_queue unsafe.Pointer, flags uint64, error_out *unsafe.Pointer) unsafe.Pointer

func raw_xpc_session_create_mach_service(mach_service string, target_queue unsafe.Pointer, flags uint64, error_out *unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_session_create_mach_service == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_session_create_mach_service(mach_service, target_queue, flags, error_out)
}

var rawfn_xpc_session_create_xpc_service func(name string, target_queue unsafe.Pointer, flags uint64, error_out *unsafe.Pointer) unsafe.Pointer

func raw_xpc_session_create_xpc_service(name string, target_queue unsafe.Pointer, flags uint64, error_out *unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_session_create_xpc_service == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_session_create_xpc_service(name, target_queue, flags, error_out)
}

var rawfn_xpc_session_send_message func(session unsafe.Pointer, message unsafe.Pointer) unsafe.Pointer

func raw_xpc_session_send_message(session unsafe.Pointer, message unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_session_send_message == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_session_send_message(session, message)
}

var rawfn_xpc_session_send_message_with_reply_async func(session unsafe.Pointer, message unsafe.Pointer, reply_handler unsafe.Pointer)

func raw_xpc_session_send_message_with_reply_async(session unsafe.Pointer, message unsafe.Pointer, reply_handler unsafe.Pointer) {
	if rawfn_xpc_session_send_message_with_reply_async == nil {
		return
	}
	rawfn_xpc_session_send_message_with_reply_async(session, message, reply_handler)
}

var rawfn_xpc_session_send_message_with_reply_sync func(session unsafe.Pointer, message unsafe.Pointer, error_out *unsafe.Pointer) unsafe.Pointer

func raw_xpc_session_send_message_with_reply_sync(session unsafe.Pointer, message unsafe.Pointer, error_out *unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_session_send_message_with_reply_sync == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_session_send_message_with_reply_sync(session, message, error_out)
}

var rawfn_xpc_session_set_cancel_handler func(session unsafe.Pointer, cancel_handler unsafe.Pointer)

func raw_xpc_session_set_cancel_handler(session unsafe.Pointer, cancel_handler unsafe.Pointer) {
	if rawfn_xpc_session_set_cancel_handler == nil {
		return
	}
	rawfn_xpc_session_set_cancel_handler(session, cancel_handler)
}

var rawfn_xpc_session_set_incoming_message_handler func(session unsafe.Pointer, handler unsafe.Pointer)

func raw_xpc_session_set_incoming_message_handler(session unsafe.Pointer, handler unsafe.Pointer) {
	if rawfn_xpc_session_set_incoming_message_handler == nil {
		return
	}
	rawfn_xpc_session_set_incoming_message_handler(session, handler)
}

var rawfn_xpc_session_set_peer_code_signing_requirement func(session unsafe.Pointer, requirement string) int32

func raw_xpc_session_set_peer_code_signing_requirement(session unsafe.Pointer, requirement string) int32 {
	if rawfn_xpc_session_set_peer_code_signing_requirement == nil {
		var zero int32
		return zero
	}
	return rawfn_xpc_session_set_peer_code_signing_requirement(session, requirement)
}

var rawfn_xpc_session_set_peer_requirement func(session unsafe.Pointer, requirement unsafe.Pointer)

func raw_xpc_session_set_peer_requirement(session unsafe.Pointer, requirement unsafe.Pointer) {
	if rawfn_xpc_session_set_peer_requirement == nil {
		return
	}
	rawfn_xpc_session_set_peer_requirement(session, requirement)
}

var rawfn_xpc_session_set_target_queue func(session unsafe.Pointer, target_queue unsafe.Pointer)

func raw_xpc_session_set_target_queue(session unsafe.Pointer, target_queue unsafe.Pointer) {
	if rawfn_xpc_session_set_target_queue == nil {
		return
	}
	rawfn_xpc_session_set_target_queue(session, target_queue)
}

var rawfn_xpc_set_event_stream_handler func(stream string, targetq unsafe.Pointer, handler unsafe.Pointer)

func raw_xpc_set_event_stream_handler(stream string, targetq unsafe.Pointer, handler unsafe.Pointer) {
	if rawfn_xpc_set_event_stream_handler == nil {
		return
	}
	rawfn_xpc_set_event_stream_handler(stream, targetq, handler)
}

var rawfn_xpc_shmem_create func(region unsafe.Pointer, length uintptr) unsafe.Pointer

func raw_xpc_shmem_create(region unsafe.Pointer, length uintptr) unsafe.Pointer {
	if rawfn_xpc_shmem_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_shmem_create(region, length)
}

var rawfn_xpc_shmem_map func(xshmem unsafe.Pointer, region unsafe.Pointer) uintptr

func raw_xpc_shmem_map(xshmem unsafe.Pointer, region unsafe.Pointer) uintptr {
	if rawfn_xpc_shmem_map == nil {
		var zero uintptr
		return zero
	}
	return rawfn_xpc_shmem_map(xshmem, region)
}

var rawfn_xpc_string_create func(string_ string) unsafe.Pointer

func raw_xpc_string_create(string_ string) unsafe.Pointer {
	if rawfn_xpc_string_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_string_create(string_)
}

var rawfn_xpc_string_create_with_format func(fmt string) unsafe.Pointer

func raw_xpc_string_create_with_format(fmt string) unsafe.Pointer {
	if rawfn_xpc_string_create_with_format == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_string_create_with_format(fmt)
}

var rawfn_xpc_string_create_with_format_and_arguments func(fmt string, ap unsafe.Pointer) unsafe.Pointer

func raw_xpc_string_create_with_format_and_arguments(fmt string, ap unsafe.Pointer) unsafe.Pointer {
	if rawfn_xpc_string_create_with_format_and_arguments == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_string_create_with_format_and_arguments(fmt, ap)
}

var rawfn_xpc_string_get_length func(xstring unsafe.Pointer) uintptr

func raw_xpc_string_get_length(xstring unsafe.Pointer) uintptr {
	if rawfn_xpc_string_get_length == nil {
		var zero uintptr
		return zero
	}
	return rawfn_xpc_string_get_length(xstring)
}

var rawfn_xpc_string_get_string_ptr func(xstring unsafe.Pointer) *byte

func raw_xpc_string_get_string_ptr(xstring unsafe.Pointer) *byte {
	if rawfn_xpc_string_get_string_ptr == nil {
		var zero *byte
		return zero
	}
	return rawfn_xpc_string_get_string_ptr(xstring)
}

var rawfn_xpc_transaction_begin func()

func raw_xpc_transaction_begin() {
	if rawfn_xpc_transaction_begin == nil {
		return
	}
	rawfn_xpc_transaction_begin()
}

var rawfn_xpc_transaction_end func()

func raw_xpc_transaction_end() {
	if rawfn_xpc_transaction_end == nil {
		return
	}
	rawfn_xpc_transaction_end()
}

var rawfn_xpc_type_get_name func(type_ unsafe.Pointer) *byte

func raw_xpc_type_get_name(type_ unsafe.Pointer) *byte {
	if rawfn_xpc_type_get_name == nil {
		var zero *byte
		return zero
	}
	return rawfn_xpc_type_get_name(type_)
}

var rawfn_xpc_uint64_create func(value uint64) unsafe.Pointer

func raw_xpc_uint64_create(value uint64) unsafe.Pointer {
	if rawfn_xpc_uint64_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_uint64_create(value)
}

var rawfn_xpc_uint64_get_value func(xuint unsafe.Pointer) uint64

func raw_xpc_uint64_get_value(xuint unsafe.Pointer) uint64 {
	if rawfn_xpc_uint64_get_value == nil {
		var zero uint64
		return zero
	}
	return rawfn_xpc_uint64_get_value(xuint)
}

var rawfn_xpc_uuid_create func(uuid *[16]byte) unsafe.Pointer

func raw_xpc_uuid_create(uuid [16]byte) unsafe.Pointer {
	if rawfn_xpc_uuid_create == nil {
		var zero unsafe.Pointer
		return zero
	}
	return rawfn_xpc_uuid_create(&uuid)
}

var rawfn_xpc_uuid_get_bytes func(xuuid unsafe.Pointer) *uint8

func raw_xpc_uuid_get_bytes(xuuid unsafe.Pointer) *uint8 {
	if rawfn_xpc_uuid_get_bytes == nil {
		var zero *uint8
		return zero
	}
	return rawfn_xpc_uuid_get_bytes(xuuid)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerRawFunc(&rawfn_launch_activate_socket, frameworkHandle, "launch_activate_socket")
	registerRawFunc(&rawfn_xpc_activity_copy_criteria, frameworkHandle, "xpc_activity_copy_criteria")
	registerRawFunc(&rawfn_xpc_activity_get_state, frameworkHandle, "xpc_activity_get_state")
	registerRawFunc(&rawfn_xpc_activity_register, frameworkHandle, "xpc_activity_register")
	registerRawFunc(&rawfn_xpc_activity_set_criteria, frameworkHandle, "xpc_activity_set_criteria")
	registerRawFunc(&rawfn_xpc_activity_set_state, frameworkHandle, "xpc_activity_set_state")
	registerRawFunc(&rawfn_xpc_activity_should_defer, frameworkHandle, "xpc_activity_should_defer")
	registerRawFunc(&rawfn_xpc_activity_unregister, frameworkHandle, "xpc_activity_unregister")
	registerRawFunc(&rawfn_xpc_array_append_value, frameworkHandle, "xpc_array_append_value")
	registerRawFunc(&rawfn_xpc_array_apply, frameworkHandle, "xpc_array_apply")
	registerRawFunc(&rawfn_xpc_array_create, frameworkHandle, "xpc_array_create")
	registerRawFunc(&rawfn_xpc_array_create_connection, frameworkHandle, "xpc_array_create_connection")
	registerRawFunc(&rawfn_xpc_array_create_empty, frameworkHandle, "xpc_array_create_empty")
	registerRawFunc(&rawfn_xpc_array_dup_fd, frameworkHandle, "xpc_array_dup_fd")
	registerRawFunc(&rawfn_xpc_array_get_array, frameworkHandle, "xpc_array_get_array")
	registerRawFunc(&rawfn_xpc_array_get_bool, frameworkHandle, "xpc_array_get_bool")
	registerRawFunc(&rawfn_xpc_array_get_count, frameworkHandle, "xpc_array_get_count")
	registerRawFunc(&rawfn_xpc_array_get_data, frameworkHandle, "xpc_array_get_data")
	registerRawFunc(&rawfn_xpc_array_get_date, frameworkHandle, "xpc_array_get_date")
	registerRawFunc(&rawfn_xpc_array_get_dictionary, frameworkHandle, "xpc_array_get_dictionary")
	registerRawFunc(&rawfn_xpc_array_get_double, frameworkHandle, "xpc_array_get_double")
	registerRawFunc(&rawfn_xpc_array_get_int64, frameworkHandle, "xpc_array_get_int64")
	registerRawFunc(&rawfn_xpc_array_get_string, frameworkHandle, "xpc_array_get_string")
	registerRawFunc(&rawfn_xpc_array_get_uint64, frameworkHandle, "xpc_array_get_uint64")
	registerRawFunc(&rawfn_xpc_array_get_uuid, frameworkHandle, "xpc_array_get_uuid")
	registerRawFunc(&rawfn_xpc_array_get_value, frameworkHandle, "xpc_array_get_value")
	registerRawFunc(&rawfn_xpc_array_set_bool, frameworkHandle, "xpc_array_set_bool")
	registerRawFunc(&rawfn_xpc_array_set_connection, frameworkHandle, "xpc_array_set_connection")
	registerRawFunc(&rawfn_xpc_array_set_data, frameworkHandle, "xpc_array_set_data")
	registerRawFunc(&rawfn_xpc_array_set_date, frameworkHandle, "xpc_array_set_date")
	registerRawFunc(&rawfn_xpc_array_set_double, frameworkHandle, "xpc_array_set_double")
	registerRawFunc(&rawfn_xpc_array_set_fd, frameworkHandle, "xpc_array_set_fd")
	registerRawFunc(&rawfn_xpc_array_set_int64, frameworkHandle, "xpc_array_set_int64")
	registerRawFunc(&rawfn_xpc_array_set_string, frameworkHandle, "xpc_array_set_string")
	registerRawFunc(&rawfn_xpc_array_set_uint64, frameworkHandle, "xpc_array_set_uint64")
	registerRawFunc(&rawfn_xpc_array_set_uuid, frameworkHandle, "xpc_array_set_uuid")
	registerRawFunc(&rawfn_xpc_array_set_value, frameworkHandle, "xpc_array_set_value")
	registerRawFunc(&rawfn_xpc_bool_create, frameworkHandle, "xpc_bool_create")
	registerRawFunc(&rawfn_xpc_bool_get_value, frameworkHandle, "xpc_bool_get_value")
	registerRawFunc(&rawfn_xpc_connection_activate, frameworkHandle, "xpc_connection_activate")
	registerRawFunc(&rawfn_xpc_connection_cancel, frameworkHandle, "xpc_connection_cancel")
	registerRawFunc(&rawfn_xpc_connection_copy_invalidation_reason, frameworkHandle, "xpc_connection_copy_invalidation_reason")
	registerRawFunc(&rawfn_xpc_connection_create, frameworkHandle, "xpc_connection_create")
	registerRawFunc(&rawfn_xpc_connection_create_from_endpoint, frameworkHandle, "xpc_connection_create_from_endpoint")
	registerRawFunc(&rawfn_xpc_connection_create_mach_service, frameworkHandle, "xpc_connection_create_mach_service")
	registerRawFunc(&rawfn_xpc_connection_get_asid, frameworkHandle, "xpc_connection_get_asid")
	registerRawFunc(&rawfn_xpc_connection_get_context, frameworkHandle, "xpc_connection_get_context")
	registerRawFunc(&rawfn_xpc_connection_get_egid, frameworkHandle, "xpc_connection_get_egid")
	registerRawFunc(&rawfn_xpc_connection_get_euid, frameworkHandle, "xpc_connection_get_euid")
	registerRawFunc(&rawfn_xpc_connection_get_name, frameworkHandle, "xpc_connection_get_name")
	registerRawFunc(&rawfn_xpc_connection_get_pid, frameworkHandle, "xpc_connection_get_pid")
	registerRawFunc(&rawfn_xpc_connection_resume, frameworkHandle, "xpc_connection_resume")
	registerRawFunc(&rawfn_xpc_connection_send_barrier, frameworkHandle, "xpc_connection_send_barrier")
	registerRawFunc(&rawfn_xpc_connection_send_message, frameworkHandle, "xpc_connection_send_message")
	registerRawFunc(&rawfn_xpc_connection_send_message_with_reply, frameworkHandle, "xpc_connection_send_message_with_reply")
	registerRawFunc(&rawfn_xpc_connection_send_message_with_reply_sync, frameworkHandle, "xpc_connection_send_message_with_reply_sync")
	registerRawFunc(&rawfn_xpc_connection_set_context, frameworkHandle, "xpc_connection_set_context")
	registerRawFunc(&rawfn_xpc_connection_set_event_handler, frameworkHandle, "xpc_connection_set_event_handler")
	registerRawFunc(&rawfn_xpc_connection_set_finalizer_f, frameworkHandle, "xpc_connection_set_finalizer_f")
	registerRawFunc(&rawfn_xpc_connection_set_peer_code_signing_requirement, frameworkHandle, "xpc_connection_set_peer_code_signing_requirement")
	registerRawFunc(&rawfn_xpc_connection_set_peer_entitlement_exists_requirement, frameworkHandle, "xpc_connection_set_peer_entitlement_exists_requirement")
	registerRawFunc(&rawfn_xpc_connection_set_peer_entitlement_matches_value_requirement, frameworkHandle, "xpc_connection_set_peer_entitlement_matches_value_requirement")
	registerRawFunc(&rawfn_xpc_connection_set_peer_lightweight_code_requirement, frameworkHandle, "xpc_connection_set_peer_lightweight_code_requirement")
	registerRawFunc(&rawfn_xpc_connection_set_peer_platform_identity_requirement, frameworkHandle, "xpc_connection_set_peer_platform_identity_requirement")
	registerRawFunc(&rawfn_xpc_connection_set_peer_requirement, frameworkHandle, "xpc_connection_set_peer_requirement")
	registerRawFunc(&rawfn_xpc_connection_set_peer_team_identity_requirement, frameworkHandle, "xpc_connection_set_peer_team_identity_requirement")
	registerRawFunc(&rawfn_xpc_connection_set_target_queue, frameworkHandle, "xpc_connection_set_target_queue")
	registerRawFunc(&rawfn_xpc_connection_suspend, frameworkHandle, "xpc_connection_suspend")
	registerRawFunc(&rawfn_xpc_copy, frameworkHandle, "xpc_copy")
	registerRawFunc(&rawfn_xpc_copy_description, frameworkHandle, "xpc_copy_description")
	registerRawFunc(&rawfn_xpc_data_create, frameworkHandle, "xpc_data_create")
	registerRawFunc(&rawfn_xpc_data_create_with_dispatch_data, frameworkHandle, "xpc_data_create_with_dispatch_data")
	registerRawFunc(&rawfn_xpc_data_get_bytes, frameworkHandle, "xpc_data_get_bytes")
	registerRawFunc(&rawfn_xpc_data_get_bytes_ptr, frameworkHandle, "xpc_data_get_bytes_ptr")
	registerRawFunc(&rawfn_xpc_data_get_length, frameworkHandle, "xpc_data_get_length")
	registerRawFunc(&rawfn_xpc_date_create, frameworkHandle, "xpc_date_create")
	registerRawFunc(&rawfn_xpc_date_create_from_current, frameworkHandle, "xpc_date_create_from_current")
	registerRawFunc(&rawfn_xpc_date_get_value, frameworkHandle, "xpc_date_get_value")
	registerRawFunc(&rawfn_xpc_dictionary_apply, frameworkHandle, "xpc_dictionary_apply")
	registerRawFunc(&rawfn_xpc_dictionary_copy_mach_send, frameworkHandle, "xpc_dictionary_copy_mach_send")
	registerRawFunc(&rawfn_xpc_dictionary_create, frameworkHandle, "xpc_dictionary_create")
	registerRawFunc(&rawfn_xpc_dictionary_create_connection, frameworkHandle, "xpc_dictionary_create_connection")
	registerRawFunc(&rawfn_xpc_dictionary_create_empty, frameworkHandle, "xpc_dictionary_create_empty")
	registerRawFunc(&rawfn_xpc_dictionary_create_reply, frameworkHandle, "xpc_dictionary_create_reply")
	registerRawFunc(&rawfn_xpc_dictionary_dup_fd, frameworkHandle, "xpc_dictionary_dup_fd")
	registerRawFunc(&rawfn_xpc_dictionary_get_array, frameworkHandle, "xpc_dictionary_get_array")
	registerRawFunc(&rawfn_xpc_dictionary_get_bool, frameworkHandle, "xpc_dictionary_get_bool")
	registerRawFunc(&rawfn_xpc_dictionary_get_count, frameworkHandle, "xpc_dictionary_get_count")
	registerRawFunc(&rawfn_xpc_dictionary_get_data, frameworkHandle, "xpc_dictionary_get_data")
	registerRawFunc(&rawfn_xpc_dictionary_get_date, frameworkHandle, "xpc_dictionary_get_date")
	registerRawFunc(&rawfn_xpc_dictionary_get_dictionary, frameworkHandle, "xpc_dictionary_get_dictionary")
	registerRawFunc(&rawfn_xpc_dictionary_get_double, frameworkHandle, "xpc_dictionary_get_double")
	registerRawFunc(&rawfn_xpc_dictionary_get_int64, frameworkHandle, "xpc_dictionary_get_int64")
	registerRawFunc(&rawfn_xpc_dictionary_get_remote_connection, frameworkHandle, "xpc_dictionary_get_remote_connection")
	registerRawFunc(&rawfn_xpc_dictionary_get_string, frameworkHandle, "xpc_dictionary_get_string")
	registerRawFunc(&rawfn_xpc_dictionary_get_uint64, frameworkHandle, "xpc_dictionary_get_uint64")
	registerRawFunc(&rawfn_xpc_dictionary_get_uuid, frameworkHandle, "xpc_dictionary_get_uuid")
	registerRawFunc(&rawfn_xpc_dictionary_get_value, frameworkHandle, "xpc_dictionary_get_value")
	registerRawFunc(&rawfn_xpc_dictionary_set_bool, frameworkHandle, "xpc_dictionary_set_bool")
	registerRawFunc(&rawfn_xpc_dictionary_set_connection, frameworkHandle, "xpc_dictionary_set_connection")
	registerRawFunc(&rawfn_xpc_dictionary_set_data, frameworkHandle, "xpc_dictionary_set_data")
	registerRawFunc(&rawfn_xpc_dictionary_set_date, frameworkHandle, "xpc_dictionary_set_date")
	registerRawFunc(&rawfn_xpc_dictionary_set_double, frameworkHandle, "xpc_dictionary_set_double")
	registerRawFunc(&rawfn_xpc_dictionary_set_fd, frameworkHandle, "xpc_dictionary_set_fd")
	registerRawFunc(&rawfn_xpc_dictionary_set_int64, frameworkHandle, "xpc_dictionary_set_int64")
	registerRawFunc(&rawfn_xpc_dictionary_set_mach_send, frameworkHandle, "xpc_dictionary_set_mach_send")
	registerRawFunc(&rawfn_xpc_dictionary_set_string, frameworkHandle, "xpc_dictionary_set_string")
	registerRawFunc(&rawfn_xpc_dictionary_set_uint64, frameworkHandle, "xpc_dictionary_set_uint64")
	registerRawFunc(&rawfn_xpc_dictionary_set_uuid, frameworkHandle, "xpc_dictionary_set_uuid")
	registerRawFunc(&rawfn_xpc_dictionary_set_value, frameworkHandle, "xpc_dictionary_set_value")
	registerRawFunc(&rawfn_xpc_double_create, frameworkHandle, "xpc_double_create")
	registerRawFunc(&rawfn_xpc_double_get_value, frameworkHandle, "xpc_double_get_value")
	registerRawFunc(&rawfn_xpc_endpoint_create, frameworkHandle, "xpc_endpoint_create")
	registerRawFunc(&rawfn_xpc_equal, frameworkHandle, "xpc_equal")
	registerRawFunc(&rawfn_xpc_fd_create, frameworkHandle, "xpc_fd_create")
	registerRawFunc(&rawfn_xpc_fd_dup, frameworkHandle, "xpc_fd_dup")
	registerRawFunc(&rawfn_xpc_get_type, frameworkHandle, "xpc_get_type")
	registerRawFunc(&rawfn_xpc_hash, frameworkHandle, "xpc_hash")
	registerRawFunc(&rawfn_xpc_int64_create, frameworkHandle, "xpc_int64_create")
	registerRawFunc(&rawfn_xpc_int64_get_value, frameworkHandle, "xpc_int64_get_value")
	registerRawFunc(&rawfn_xpc_listener_activate, frameworkHandle, "xpc_listener_activate")
	registerRawFunc(&rawfn_xpc_listener_cancel, frameworkHandle, "xpc_listener_cancel")
	registerRawFunc(&rawfn_xpc_listener_copy_description, frameworkHandle, "xpc_listener_copy_description")
	registerRawFunc(&rawfn_xpc_listener_create, frameworkHandle, "xpc_listener_create")
	registerRawFunc(&rawfn_xpc_listener_reject_peer, frameworkHandle, "xpc_listener_reject_peer")
	registerRawFunc(&rawfn_xpc_listener_set_peer_code_signing_requirement, frameworkHandle, "xpc_listener_set_peer_code_signing_requirement")
	registerRawFunc(&rawfn_xpc_listener_set_peer_requirement, frameworkHandle, "xpc_listener_set_peer_requirement")
	registerRawFunc(&rawfn_xpc_main, frameworkHandle, "xpc_main")
	registerRawFunc(&rawfn_xpc_null_create, frameworkHandle, "xpc_null_create")
	registerRawFunc(&rawfn_xpc_peer_requirement_create_entitlement_exists, frameworkHandle, "xpc_peer_requirement_create_entitlement_exists")
	registerRawFunc(&rawfn_xpc_peer_requirement_create_entitlement_matches_value, frameworkHandle, "xpc_peer_requirement_create_entitlement_matches_value")
	registerRawFunc(&rawfn_xpc_peer_requirement_create_lwcr, frameworkHandle, "xpc_peer_requirement_create_lwcr")
	registerRawFunc(&rawfn_xpc_peer_requirement_create_platform_identity, frameworkHandle, "xpc_peer_requirement_create_platform_identity")
	registerRawFunc(&rawfn_xpc_peer_requirement_create_team_identity, frameworkHandle, "xpc_peer_requirement_create_team_identity")
	registerRawFunc(&rawfn_xpc_peer_requirement_match_received_message, frameworkHandle, "xpc_peer_requirement_match_received_message")
	registerRawFunc(&rawfn_xpc_release, frameworkHandle, "xpc_release")
	registerRawFunc(&rawfn_xpc_retain, frameworkHandle, "xpc_retain")
	registerRawFunc(&rawfn_xpc_rich_error_can_retry, frameworkHandle, "xpc_rich_error_can_retry")
	registerRawFunc(&rawfn_xpc_rich_error_copy_description, frameworkHandle, "xpc_rich_error_copy_description")
	registerRawFunc(&rawfn_xpc_session_activate, frameworkHandle, "xpc_session_activate")
	registerRawFunc(&rawfn_xpc_session_cancel, frameworkHandle, "xpc_session_cancel")
	registerRawFunc(&rawfn_xpc_session_copy_description, frameworkHandle, "xpc_session_copy_description")
	registerRawFunc(&rawfn_xpc_session_create_mach_service, frameworkHandle, "xpc_session_create_mach_service")
	registerRawFunc(&rawfn_xpc_session_create_xpc_service, frameworkHandle, "xpc_session_create_xpc_service")
	registerRawFunc(&rawfn_xpc_session_send_message, frameworkHandle, "xpc_session_send_message")
	registerRawFunc(&rawfn_xpc_session_send_message_with_reply_async, frameworkHandle, "xpc_session_send_message_with_reply_async")
	registerRawFunc(&rawfn_xpc_session_send_message_with_reply_sync, frameworkHandle, "xpc_session_send_message_with_reply_sync")
	registerRawFunc(&rawfn_xpc_session_set_cancel_handler, frameworkHandle, "xpc_session_set_cancel_handler")
	registerRawFunc(&rawfn_xpc_session_set_incoming_message_handler, frameworkHandle, "xpc_session_set_incoming_message_handler")
	registerRawFunc(&rawfn_xpc_session_set_peer_code_signing_requirement, frameworkHandle, "xpc_session_set_peer_code_signing_requirement")
	registerRawFunc(&rawfn_xpc_session_set_peer_requirement, frameworkHandle, "xpc_session_set_peer_requirement")
	registerRawFunc(&rawfn_xpc_session_set_target_queue, frameworkHandle, "xpc_session_set_target_queue")
	registerRawFunc(&rawfn_xpc_set_event_stream_handler, frameworkHandle, "xpc_set_event_stream_handler")
	registerRawFunc(&rawfn_xpc_shmem_create, frameworkHandle, "xpc_shmem_create")
	registerRawFunc(&rawfn_xpc_shmem_map, frameworkHandle, "xpc_shmem_map")
	registerRawFunc(&rawfn_xpc_string_create, frameworkHandle, "xpc_string_create")
	registerRawFunc(&rawfn_xpc_string_create_with_format, frameworkHandle, "xpc_string_create_with_format")
	registerRawFunc(&rawfn_xpc_string_create_with_format_and_arguments, frameworkHandle, "xpc_string_create_with_format_and_arguments")
	registerRawFunc(&rawfn_xpc_string_get_length, frameworkHandle, "xpc_string_get_length")
	registerRawFunc(&rawfn_xpc_string_get_string_ptr, frameworkHandle, "xpc_string_get_string_ptr")
	registerRawFunc(&rawfn_xpc_transaction_begin, frameworkHandle, "xpc_transaction_begin")
	registerRawFunc(&rawfn_xpc_transaction_end, frameworkHandle, "xpc_transaction_end")
	registerRawFunc(&rawfn_xpc_type_get_name, frameworkHandle, "xpc_type_get_name")
	registerRawFunc(&rawfn_xpc_uint64_create, frameworkHandle, "xpc_uint64_create")
	registerRawFunc(&rawfn_xpc_uint64_get_value, frameworkHandle, "xpc_uint64_get_value")
	registerRawFunc(&rawfn_xpc_uuid_create, frameworkHandle, "xpc_uuid_create")
	registerRawFunc(&rawfn_xpc_uuid_get_bytes, frameworkHandle, "xpc_uuid_get_bytes")
}

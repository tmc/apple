// Code generated from Apple documentation for EndpointSecurity. DO NOT EDIT.

package endpointsecurity

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: EndpointSecurity: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: EndpointSecurity: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: EndpointSecurity: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _es_clear_cache func(client *Es_client_t) unsafe.Pointer

// Es_clear_cache clears all cached results for all clients.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_clear_cache(_:)
func Es_clear_cache(client *Es_client_t) unsafe.Pointer {
	if _es_clear_cache == nil {
		panic("EndpointSecurity: symbol es_clear_cache not loaded")
	}
	return _es_clear_cache(client)
}

var _es_delete_client func(client *Es_client_t) unsafe.Pointer

// Es_delete_client destroys and disconnects a client instance from the Endpoint Security system.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_delete_client(_:)
func Es_delete_client(client *Es_client_t) unsafe.Pointer {
	if _es_delete_client == nil {
		panic("EndpointSecurity: symbol es_delete_client not loaded")
	}
	return _es_delete_client(client)
}

var _es_exec_arg func(event *Es_event_exec_t, index uint32) Es_string_token_t

// Es_exec_arg gets the argument at the specified position from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_arg(_:_:)
func Es_exec_arg(event *Es_event_exec_t, index uint32) Es_string_token_t {
	if _es_exec_arg == nil {
		panic("EndpointSecurity: symbol es_exec_arg not loaded")
	}
	return _es_exec_arg(event, index)
}

var _es_exec_arg_count func(event *Es_event_exec_t) uint32

// Es_exec_arg_count gets the number of arguments from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_arg_count(_:)
func Es_exec_arg_count(event *Es_event_exec_t) uint32 {
	if _es_exec_arg_count == nil {
		panic("EndpointSecurity: symbol es_exec_arg_count not loaded")
	}
	return _es_exec_arg_count(event)
}

var _es_exec_env func(event *Es_event_exec_t, index uint32) Es_string_token_t

// Es_exec_env gets the environment variable at the specified position from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_env(_:_:)
func Es_exec_env(event *Es_event_exec_t, index uint32) Es_string_token_t {
	if _es_exec_env == nil {
		panic("EndpointSecurity: symbol es_exec_env not loaded")
	}
	return _es_exec_env(event, index)
}

var _es_exec_env_count func(event *Es_event_exec_t) uint32

// Es_exec_env_count gets the number of environment variables from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_env_count(_:)
func Es_exec_env_count(event *Es_event_exec_t) uint32 {
	if _es_exec_env_count == nil {
		panic("EndpointSecurity: symbol es_exec_env_count not loaded")
	}
	return _es_exec_env_count(event)
}

var _es_exec_fd func(event *Es_event_exec_t, index uint32) *Es_fd_t

// Es_exec_fd gets the file descriptor at the specified position from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_fd(_:_:)
func Es_exec_fd(event *Es_event_exec_t, index uint32) *Es_fd_t {
	if _es_exec_fd == nil {
		panic("EndpointSecurity: symbol es_exec_fd not loaded")
	}
	return _es_exec_fd(event, index)
}

var _es_exec_fd_count func(event *Es_event_exec_t) uint32

// Es_exec_fd_count gets the number of file descriptors from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_fd_count(_:)
func Es_exec_fd_count(event *Es_event_exec_t) uint32 {
	if _es_exec_fd_count == nil {
		panic("EndpointSecurity: symbol es_exec_fd_count not loaded")
	}
	return _es_exec_fd_count(event)
}

var _es_invert_muting func(client *Es_client_t, mute_type unsafe.Pointer) unsafe.Pointer

// Es_invert_muting.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_invert_muting(_:_:)
func Es_invert_muting(client *Es_client_t, mute_type unsafe.Pointer) unsafe.Pointer {
	if _es_invert_muting == nil {
		panic("EndpointSecurity: symbol es_invert_muting not loaded")
	}
	return _es_invert_muting(client, mute_type)
}

var _es_mute_path func(client *Es_client_t, path string, type_ unsafe.Pointer) unsafe.Pointer

// Es_mute_path suppresses events from executables that match a given path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path(_:_:_:)
func Es_mute_path(client *Es_client_t, path string, type_ unsafe.Pointer) unsafe.Pointer {
	if _es_mute_path == nil {
		panic("EndpointSecurity: symbol es_mute_path not loaded")
	}
	return _es_mute_path(client, path, type_)
}

var _es_mute_path_events func(client *Es_client_t, path string, type_ unsafe.Pointer, events uintptr, event_count uintptr) unsafe.Pointer

// Es_mute_path_events suppresses a subset of events from executables that match a given path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path_events(_:_:_:_:_:)
func Es_mute_path_events(client *Es_client_t, path string, type_ unsafe.Pointer, events uintptr, event_count uintptr) unsafe.Pointer {
	if _es_mute_path_events == nil {
		panic("EndpointSecurity: symbol es_mute_path_events not loaded")
	}
	return _es_mute_path_events(client, path, type_, events, event_count)
}

var _es_mute_process func(client *Es_client_t, audit_token uintptr) unsafe.Pointer

// Es_mute_process suppresses events from a given process.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_process(_:_:)
func Es_mute_process(client *Es_client_t, audit_token uintptr) unsafe.Pointer {
	if _es_mute_process == nil {
		panic("EndpointSecurity: symbol es_mute_process not loaded")
	}
	return _es_mute_process(client, audit_token)
}

var _es_mute_process_events func(client *Es_client_t, audit_token uintptr, events uintptr, event_count uintptr) unsafe.Pointer

// Es_mute_process_events suppresses a subset of events from a given process.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_process_events(_:_:_:_:)
func Es_mute_process_events(client *Es_client_t, audit_token uintptr, events uintptr, event_count uintptr) unsafe.Pointer {
	if _es_mute_process_events == nil {
		panic("EndpointSecurity: symbol es_mute_process_events not loaded")
	}
	return _es_mute_process_events(client, audit_token, events, event_count)
}

var _es_muted_paths_events func(client *Es_client_t, muted_paths *Es_muted_paths_t) unsafe.Pointer

// Es_muted_paths_events retrieve a list of all muted paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_muted_paths_events(_:_:)
func Es_muted_paths_events(client *Es_client_t, muted_paths *Es_muted_paths_t) unsafe.Pointer {
	if _es_muted_paths_events == nil {
		panic("EndpointSecurity: symbol es_muted_paths_events not loaded")
	}
	return _es_muted_paths_events(client, muted_paths)
}

var _es_muted_processes_events func(client *Es_client_t, muted_processes *Es_muted_processes_t) unsafe.Pointer

// Es_muted_processes_events retrieve a list of all muted processes.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_muted_processes_events(_:_:)
func Es_muted_processes_events(client *Es_client_t, muted_processes *Es_muted_processes_t) unsafe.Pointer {
	if _es_muted_processes_events == nil {
		panic("EndpointSecurity: symbol es_muted_processes_events not loaded")
	}
	return _es_muted_processes_events(client, muted_processes)
}

var _es_muting_inverted func(client *Es_client_t, mute_type unsafe.Pointer) unsafe.Pointer

// Es_muting_inverted.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_muting_inverted(_:_:)
func Es_muting_inverted(client *Es_client_t, mute_type unsafe.Pointer) unsafe.Pointer {
	if _es_muting_inverted == nil {
		panic("EndpointSecurity: symbol es_muting_inverted not loaded")
	}
	return _es_muting_inverted(client, mute_type)
}

var _es_new_client func(client *Es_client_t, handler Es_handler_block_t) unsafe.Pointer

// Es_new_client creates a new client instance and connects it to the Endpoint Security system.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_new_client(_:_:)
func Es_new_client(client *Es_client_t, handler Es_handler_block_t) unsafe.Pointer {
	if _es_new_client == nil {
		panic("EndpointSecurity: symbol es_new_client not loaded")
	}
	return _es_new_client(client, handler)
}

var _es_release_message func(msg *Es_message_t)

// Es_release_message releases a previously-retained message.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_release_message(_:)
func Es_release_message(msg *Es_message_t) {
	if _es_release_message == nil {
		panic("EndpointSecurity: symbol es_release_message not loaded")
	}
	_es_release_message(msg)
}

var _es_release_muted_paths func(muted_paths *Es_muted_paths_t)

// Es_release_muted_paths frees resources associated with a set of previously-retrieved muted paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_release_muted_paths(_:)
func Es_release_muted_paths(muted_paths *Es_muted_paths_t) {
	if _es_release_muted_paths == nil {
		panic("EndpointSecurity: symbol es_release_muted_paths not loaded")
	}
	_es_release_muted_paths(muted_paths)
}

var _es_release_muted_processes func(muted_processes *Es_muted_processes_t)

// Es_release_muted_processes frees resources associated with a set of previously-retrieved muted processes.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_release_muted_processes(_:)
func Es_release_muted_processes(muted_processes *Es_muted_processes_t) {
	if _es_release_muted_processes == nil {
		panic("EndpointSecurity: symbol es_release_muted_processes not loaded")
	}
	_es_release_muted_processes(muted_processes)
}

var _es_respond_auth_result func(client *Es_client_t, message *Es_message_t, result unsafe.Pointer, cache bool) unsafe.Pointer

// Es_respond_auth_result responds to an event that requires an authorization response.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_respond_auth_result(_:_:_:_:)
func Es_respond_auth_result(client *Es_client_t, message *Es_message_t, result unsafe.Pointer, cache bool) unsafe.Pointer {
	if _es_respond_auth_result == nil {
		panic("EndpointSecurity: symbol es_respond_auth_result not loaded")
	}
	return _es_respond_auth_result(client, message, result, cache)
}

var _es_respond_flags_result func(client *Es_client_t, message *Es_message_t, authorized_flags uint32, cache bool) unsafe.Pointer

// Es_respond_flags_result responds to an event that requires authorization flags as a response.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_respond_flags_result(_:_:_:_:)
func Es_respond_flags_result(client *Es_client_t, message *Es_message_t, authorized_flags uint32, cache bool) unsafe.Pointer {
	if _es_respond_flags_result == nil {
		panic("EndpointSecurity: symbol es_respond_flags_result not loaded")
	}
	return _es_respond_flags_result(client, message, authorized_flags, cache)
}

var _es_retain_message func(msg *Es_message_t)

// Es_retain_message retains the given message, extending its lifetime until released.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_retain_message(_:)
func Es_retain_message(msg *Es_message_t) {
	if _es_retain_message == nil {
		panic("EndpointSecurity: symbol es_retain_message not loaded")
	}
	_es_retain_message(msg)
}

var _es_subscribe func(client *Es_client_t, events uintptr, event_count uint32) unsafe.Pointer

// Es_subscribe subscribes a client to a set of events.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_subscribe(_:_:_:)
func Es_subscribe(client *Es_client_t, events uintptr, event_count uint32) unsafe.Pointer {
	if _es_subscribe == nil {
		panic("EndpointSecurity: symbol es_subscribe not loaded")
	}
	return _es_subscribe(client, events, event_count)
}

var _es_subscriptions func(client *Es_client_t, count *uintptr, subscriptions uintptr) unsafe.Pointer

// Es_subscriptions returns a list of the client’s subscriptions.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_subscriptions(_:_:_:)
func Es_subscriptions(client *Es_client_t, count *uintptr, subscriptions uintptr) unsafe.Pointer {
	if _es_subscriptions == nil {
		panic("EndpointSecurity: symbol es_subscriptions not loaded")
	}
	return _es_subscriptions(client, count, subscriptions)
}

var _es_unmute_all_paths func(client *Es_client_t) unsafe.Pointer

// Es_unmute_all_paths restores event delivery from previously-muted paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_all_paths(_:)
func Es_unmute_all_paths(client *Es_client_t) unsafe.Pointer {
	if _es_unmute_all_paths == nil {
		panic("EndpointSecurity: symbol es_unmute_all_paths not loaded")
	}
	return _es_unmute_all_paths(client)
}

var _es_unmute_all_target_paths func(client *Es_client_t) unsafe.Pointer

// Es_unmute_all_target_paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_all_target_paths(_:)
func Es_unmute_all_target_paths(client *Es_client_t) unsafe.Pointer {
	if _es_unmute_all_target_paths == nil {
		panic("EndpointSecurity: symbol es_unmute_all_target_paths not loaded")
	}
	return _es_unmute_all_target_paths(client)
}

var _es_unmute_path func(client *Es_client_t, path string, type_ unsafe.Pointer) unsafe.Pointer

// Es_unmute_path restores event delivery from a previously-muted path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_path(_:_:_:)
func Es_unmute_path(client *Es_client_t, path string, type_ unsafe.Pointer) unsafe.Pointer {
	if _es_unmute_path == nil {
		panic("EndpointSecurity: symbol es_unmute_path not loaded")
	}
	return _es_unmute_path(client, path, type_)
}

var _es_unmute_path_events func(client *Es_client_t, path string, type_ unsafe.Pointer, events uintptr, event_count uintptr) unsafe.Pointer

// Es_unmute_path_events restores event delivery of a subset of events from a previously-muted path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_path_events(_:_:_:_:_:)
func Es_unmute_path_events(client *Es_client_t, path string, type_ unsafe.Pointer, events uintptr, event_count uintptr) unsafe.Pointer {
	if _es_unmute_path_events == nil {
		panic("EndpointSecurity: symbol es_unmute_path_events not loaded")
	}
	return _es_unmute_path_events(client, path, type_, events, event_count)
}

var _es_unmute_process func(client *Es_client_t, audit_token uintptr) unsafe.Pointer

// Es_unmute_process restores event delivery from a previously-muted process.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_process(_:_:)
func Es_unmute_process(client *Es_client_t, audit_token uintptr) unsafe.Pointer {
	if _es_unmute_process == nil {
		panic("EndpointSecurity: symbol es_unmute_process not loaded")
	}
	return _es_unmute_process(client, audit_token)
}

var _es_unmute_process_events func(client *Es_client_t, audit_token uintptr, events uintptr, event_count uintptr) unsafe.Pointer

// Es_unmute_process_events restores event delivery of a subset of events from a previously-muted process.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_process_events(_:_:_:_:)
func Es_unmute_process_events(client *Es_client_t, audit_token uintptr, events uintptr, event_count uintptr) unsafe.Pointer {
	if _es_unmute_process_events == nil {
		panic("EndpointSecurity: symbol es_unmute_process_events not loaded")
	}
	return _es_unmute_process_events(client, audit_token, events, event_count)
}

var _es_unsubscribe func(client *Es_client_t, events uintptr, event_count uint32) unsafe.Pointer

// Es_unsubscribe unsubscribes the provided client from a set of events.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unsubscribe(_:_:_:)
func Es_unsubscribe(client *Es_client_t, events uintptr, event_count uint32) unsafe.Pointer {
	if _es_unsubscribe == nil {
		panic("EndpointSecurity: symbol es_unsubscribe not loaded")
	}
	return _es_unsubscribe(client, events, event_count)
}

var _es_unsubscribe_all func(client *Es_client_t) unsafe.Pointer

// Es_unsubscribe_all unsubscribes a client from all events.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unsubscribe_all(_:)
func Es_unsubscribe_all(client *Es_client_t) unsafe.Pointer {
	if _es_unsubscribe_all == nil {
		panic("EndpointSecurity: symbol es_unsubscribe_all not loaded")
	}
	return _es_unsubscribe_all(client)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_es_clear_cache, frameworkHandle, "es_clear_cache")
		registerFunc(&_es_delete_client, frameworkHandle, "es_delete_client")
		registerFunc(&_es_exec_arg, frameworkHandle, "es_exec_arg")
		registerFunc(&_es_exec_arg_count, frameworkHandle, "es_exec_arg_count")
		registerFunc(&_es_exec_env, frameworkHandle, "es_exec_env")
		registerFunc(&_es_exec_env_count, frameworkHandle, "es_exec_env_count")
		registerFunc(&_es_exec_fd, frameworkHandle, "es_exec_fd")
		registerFunc(&_es_exec_fd_count, frameworkHandle, "es_exec_fd_count")
		registerFunc(&_es_invert_muting, frameworkHandle, "es_invert_muting")
		registerFunc(&_es_mute_path, frameworkHandle, "es_mute_path")
		registerFunc(&_es_mute_path_events, frameworkHandle, "es_mute_path_events")
		registerFunc(&_es_mute_process, frameworkHandle, "es_mute_process")
		registerFunc(&_es_mute_process_events, frameworkHandle, "es_mute_process_events")
		registerFunc(&_es_muted_paths_events, frameworkHandle, "es_muted_paths_events")
		registerFunc(&_es_muted_processes_events, frameworkHandle, "es_muted_processes_events")
		registerFunc(&_es_muting_inverted, frameworkHandle, "es_muting_inverted")
		registerFunc(&_es_new_client, frameworkHandle, "es_new_client")
		registerFunc(&_es_release_message, frameworkHandle, "es_release_message")
		registerFunc(&_es_release_muted_paths, frameworkHandle, "es_release_muted_paths")
		registerFunc(&_es_release_muted_processes, frameworkHandle, "es_release_muted_processes")
		registerFunc(&_es_respond_auth_result, frameworkHandle, "es_respond_auth_result")
		registerFunc(&_es_respond_flags_result, frameworkHandle, "es_respond_flags_result")
		registerFunc(&_es_retain_message, frameworkHandle, "es_retain_message")
		registerFunc(&_es_subscribe, frameworkHandle, "es_subscribe")
		registerFunc(&_es_subscriptions, frameworkHandle, "es_subscriptions")
		registerFunc(&_es_unmute_all_paths, frameworkHandle, "es_unmute_all_paths")
		registerFunc(&_es_unmute_all_target_paths, frameworkHandle, "es_unmute_all_target_paths")
		registerFunc(&_es_unmute_path, frameworkHandle, "es_unmute_path")
		registerFunc(&_es_unmute_path_events, frameworkHandle, "es_unmute_path_events")
		registerFunc(&_es_unmute_process, frameworkHandle, "es_unmute_process")
		registerFunc(&_es_unmute_process_events, frameworkHandle, "es_unmute_process_events")
		registerFunc(&_es_unsubscribe, frameworkHandle, "es_unsubscribe")
		registerFunc(&_es_unsubscribe_all, frameworkHandle, "es_unsubscribe_all")
	}


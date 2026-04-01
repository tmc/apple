// Code generated from Apple documentation for EndpointSecurity. DO NOT EDIT.

package endpointsecurity

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("EndpointSecurity: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("EndpointSecurity: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("EndpointSecurity: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("EndpointSecurity: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _es_clear_cache func(client *Es_client_t) EsClearCacheResult
var _es_clear_cacheErr error

func tryEs_clear_cache(client *Es_client_t) (EsClearCacheResult, error) {
	if _es_clear_cache == nil {
		return *new(EsClearCacheResult), symbolCallError("es_clear_cache", "10.15", _es_clear_cacheErr)
	}
	return _es_clear_cache(client), nil
}

// Es_clear_cache clears all cached results for all clients.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_clear_cache(_:)
func Es_clear_cache(client *Es_client_t) EsClearCacheResult {
	result, callErr := tryEs_clear_cache(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_delete_client func(client *Es_client_t) EsReturn
var _es_delete_clientErr error

func tryEs_delete_client(client *Es_client_t) (EsReturn, error) {
	if _es_delete_client == nil {
		return *new(EsReturn), symbolCallError("es_delete_client", "10.15", _es_delete_clientErr)
	}
	return _es_delete_client(client), nil
}

// Es_delete_client destroys and disconnects a client instance from the Endpoint Security system.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_delete_client(_:)
func Es_delete_client(client *Es_client_t) EsReturn {
	result, callErr := tryEs_delete_client(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_exec_arg func(event *Es_event_exec_t, index uint32) Es_string_token_t
var _es_exec_argErr error

func tryEs_exec_arg(event *Es_event_exec_t, index uint32) (Es_string_token_t, error) {
	if _es_exec_arg == nil {
		return Es_string_token_t{}, symbolCallError("es_exec_arg", "10.15", _es_exec_argErr)
	}
	return _es_exec_arg(event, index), nil
}

// Es_exec_arg gets the argument at the specified position from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_arg(_:_:)
func Es_exec_arg(event *Es_event_exec_t, index uint32) Es_string_token_t {
	result, callErr := tryEs_exec_arg(event, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_exec_arg_count func(event *Es_event_exec_t) uint32
var _es_exec_arg_countErr error

func tryEs_exec_arg_count(event *Es_event_exec_t) (uint32, error) {
	if _es_exec_arg_count == nil {
		return 0, symbolCallError("es_exec_arg_count", "10.15", _es_exec_arg_countErr)
	}
	return _es_exec_arg_count(event), nil
}

// Es_exec_arg_count gets the number of arguments from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_arg_count(_:)
func Es_exec_arg_count(event *Es_event_exec_t) uint32 {
	result, callErr := tryEs_exec_arg_count(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_exec_env func(event *Es_event_exec_t, index uint32) Es_string_token_t
var _es_exec_envErr error

func tryEs_exec_env(event *Es_event_exec_t, index uint32) (Es_string_token_t, error) {
	if _es_exec_env == nil {
		return Es_string_token_t{}, symbolCallError("es_exec_env", "10.15", _es_exec_envErr)
	}
	return _es_exec_env(event, index), nil
}

// Es_exec_env gets the environment variable at the specified position from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_env(_:_:)
func Es_exec_env(event *Es_event_exec_t, index uint32) Es_string_token_t {
	result, callErr := tryEs_exec_env(event, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_exec_env_count func(event *Es_event_exec_t) uint32
var _es_exec_env_countErr error

func tryEs_exec_env_count(event *Es_event_exec_t) (uint32, error) {
	if _es_exec_env_count == nil {
		return 0, symbolCallError("es_exec_env_count", "10.15", _es_exec_env_countErr)
	}
	return _es_exec_env_count(event), nil
}

// Es_exec_env_count gets the number of environment variables from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_env_count(_:)
func Es_exec_env_count(event *Es_event_exec_t) uint32 {
	result, callErr := tryEs_exec_env_count(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_exec_fd func(event *Es_event_exec_t, index uint32) *Es_fd_t
var _es_exec_fdErr error

func tryEs_exec_fd(event *Es_event_exec_t, index uint32) (*Es_fd_t, error) {
	if _es_exec_fd == nil {
		return nil, symbolCallError("es_exec_fd", "11.0", _es_exec_fdErr)
	}
	return _es_exec_fd(event, index), nil
}

// Es_exec_fd gets the file descriptor at the specified position from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_fd(_:_:)
func Es_exec_fd(event *Es_event_exec_t, index uint32) *Es_fd_t {
	result, callErr := tryEs_exec_fd(event, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_exec_fd_count func(event *Es_event_exec_t) uint32
var _es_exec_fd_countErr error

func tryEs_exec_fd_count(event *Es_event_exec_t) (uint32, error) {
	if _es_exec_fd_count == nil {
		return 0, symbolCallError("es_exec_fd_count", "11.0", _es_exec_fd_countErr)
	}
	return _es_exec_fd_count(event), nil
}

// Es_exec_fd_count gets the number of file descriptors from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_fd_count(_:)
func Es_exec_fd_count(event *Es_event_exec_t) uint32 {
	result, callErr := tryEs_exec_fd_count(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_invert_muting func(client *Es_client_t, mute_type EsMuteInversionType) EsReturn
var _es_invert_mutingErr error

func tryEs_invert_muting(client *Es_client_t, mute_type EsMuteInversionType) (EsReturn, error) {
	if _es_invert_muting == nil {
		return *new(EsReturn), symbolCallError("es_invert_muting", "13.0", _es_invert_mutingErr)
	}
	return _es_invert_muting(client, mute_type), nil
}

// Es_invert_muting.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_invert_muting(_:_:)
func Es_invert_muting(client *Es_client_t, mute_type EsMuteInversionType) EsReturn {
	result, callErr := tryEs_invert_muting(client, mute_type)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_mute_path func(client *Es_client_t, path string, type_ EsMutePathType) EsReturn
var _es_mute_pathErr error

func tryEs_mute_path(client *Es_client_t, path string, type_ EsMutePathType) (EsReturn, error) {
	if _es_mute_path == nil {
		return *new(EsReturn), symbolCallError("es_mute_path", "12.0", _es_mute_pathErr)
	}
	return _es_mute_path(client, path, type_), nil
}

// Es_mute_path suppresses events from executables that match a given path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path(_:_:_:)
func Es_mute_path(client *Es_client_t, path string, type_ EsMutePathType) EsReturn {
	result, callErr := tryEs_mute_path(client, path, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_mute_path_events func(client *Es_client_t, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) EsReturn
var _es_mute_path_eventsErr error

func tryEs_mute_path_events(client *Es_client_t, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) (EsReturn, error) {
	if _es_mute_path_events == nil {
		return *new(EsReturn), symbolCallError("es_mute_path_events", "12.0", _es_mute_path_eventsErr)
	}
	return _es_mute_path_events(client, path, type_, events, event_count), nil
}

// Es_mute_path_events suppresses a subset of events from executables that match a given path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path_events(_:_:_:_:_:)
func Es_mute_path_events(client *Es_client_t, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) EsReturn {
	result, callErr := tryEs_mute_path_events(client, path, type_, events, event_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_muted_paths_events func(client *Es_client_t, muted_paths *Es_muted_paths_t) EsReturn
var _es_muted_paths_eventsErr error

func tryEs_muted_paths_events(client *Es_client_t, muted_paths *Es_muted_paths_t) (EsReturn, error) {
	if _es_muted_paths_events == nil {
		return *new(EsReturn), symbolCallError("es_muted_paths_events", "12.0", _es_muted_paths_eventsErr)
	}
	return _es_muted_paths_events(client, muted_paths), nil
}

// Es_muted_paths_events retrieve a list of all muted paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_muted_paths_events(_:_:)
func Es_muted_paths_events(client *Es_client_t, muted_paths *Es_muted_paths_t) EsReturn {
	result, callErr := tryEs_muted_paths_events(client, muted_paths)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_muted_processes_events func(client *Es_client_t, muted_processes *Es_muted_processes_t) EsReturn
var _es_muted_processes_eventsErr error

func tryEs_muted_processes_events(client *Es_client_t, muted_processes *Es_muted_processes_t) (EsReturn, error) {
	if _es_muted_processes_events == nil {
		return *new(EsReturn), symbolCallError("es_muted_processes_events", "12.0", _es_muted_processes_eventsErr)
	}
	return _es_muted_processes_events(client, muted_processes), nil
}

// Es_muted_processes_events retrieve a list of all muted processes.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_muted_processes_events(_:_:)
func Es_muted_processes_events(client *Es_client_t, muted_processes *Es_muted_processes_t) EsReturn {
	result, callErr := tryEs_muted_processes_events(client, muted_processes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_muting_inverted func(client *Es_client_t, mute_type EsMuteInversionType) Es_mute_inverted_return_t
var _es_muting_invertedErr error

func tryEs_muting_inverted(client *Es_client_t, mute_type EsMuteInversionType) (Es_mute_inverted_return_t, error) {
	if _es_muting_inverted == nil {
		return *new(Es_mute_inverted_return_t), symbolCallError("es_muting_inverted", "13.0", _es_muting_invertedErr)
	}
	return _es_muting_inverted(client, mute_type), nil
}

// Es_muting_inverted.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_muting_inverted(_:_:)
func Es_muting_inverted(client *Es_client_t, mute_type EsMuteInversionType) Es_mute_inverted_return_t {
	result, callErr := tryEs_muting_inverted(client, mute_type)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_new_client func(client **Es_client_t, handler unsafe.Pointer) EsNewClientResult
var _es_new_clientErr error

func tryEs_new_client(client **Es_client_t, handler Es_handler_block_t) (EsNewClientResult, error) {
	if _es_new_client == nil {
		return *new(EsNewClientResult), symbolCallError("es_new_client", "10.15", _es_new_clientErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *Es_client_t, blockArg1 *Es_message_t) { handler(blockArg0, blockArg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _es_new_client(client, _block0), nil
}

// Es_new_client creates a new client instance and connects it to the Endpoint Security system.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_new_client(_:_:)
func Es_new_client(client **Es_client_t, handler Es_handler_block_t) EsNewClientResult {
	result, callErr := tryEs_new_client(client, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_release_message func(msg *Es_message_t)
var _es_release_messageErr error

func tryEs_release_message(msg *Es_message_t) error {
	if _es_release_message == nil {
		return symbolCallError("es_release_message", "11.0", _es_release_messageErr)
	}
	_es_release_message(msg)
	return nil
}

// Es_release_message releases a previously-retained message.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_release_message(_:)
func Es_release_message(msg *Es_message_t) {
	if callErr := tryEs_release_message(msg); callErr != nil {
		panic(callErr)
	}
}

var _es_release_muted_paths func(muted_paths *Es_muted_paths_t)
var _es_release_muted_pathsErr error

func tryEs_release_muted_paths(muted_paths *Es_muted_paths_t) error {
	if _es_release_muted_paths == nil {
		return symbolCallError("es_release_muted_paths", "12.0", _es_release_muted_pathsErr)
	}
	_es_release_muted_paths(muted_paths)
	return nil
}

// Es_release_muted_paths frees resources associated with a set of previously-retrieved muted paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_release_muted_paths(_:)
func Es_release_muted_paths(muted_paths *Es_muted_paths_t) {
	if callErr := tryEs_release_muted_paths(muted_paths); callErr != nil {
		panic(callErr)
	}
}

var _es_release_muted_processes func(muted_processes *Es_muted_processes_t)
var _es_release_muted_processesErr error

func tryEs_release_muted_processes(muted_processes *Es_muted_processes_t) error {
	if _es_release_muted_processes == nil {
		return symbolCallError("es_release_muted_processes", "12.0", _es_release_muted_processesErr)
	}
	_es_release_muted_processes(muted_processes)
	return nil
}

// Es_release_muted_processes frees resources associated with a set of previously-retrieved muted processes.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_release_muted_processes(_:)
func Es_release_muted_processes(muted_processes *Es_muted_processes_t) {
	if callErr := tryEs_release_muted_processes(muted_processes); callErr != nil {
		panic(callErr)
	}
}

var _es_respond_auth_result func(client *Es_client_t, message *Es_message_t, result EsAuthResult, cache bool) EsRespondResult
var _es_respond_auth_resultErr error

func tryEs_respond_auth_result(client *Es_client_t, message *Es_message_t, result EsAuthResult, cache bool) (EsRespondResult, error) {
	if _es_respond_auth_result == nil {
		return *new(EsRespondResult), symbolCallError("es_respond_auth_result", "10.15", _es_respond_auth_resultErr)
	}
	return _es_respond_auth_result(client, message, result, cache), nil
}

// Es_respond_auth_result responds to an event that requires an authorization response.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_respond_auth_result(_:_:_:_:)
func Es_respond_auth_result(client *Es_client_t, message *Es_message_t, result EsAuthResult, cache bool) EsRespondResult {
	result0, callErr := tryEs_respond_auth_result(client, message, result, cache)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _es_respond_flags_result func(client *Es_client_t, message *Es_message_t, authorized_flags uint32, cache bool) EsRespondResult
var _es_respond_flags_resultErr error

func tryEs_respond_flags_result(client *Es_client_t, message *Es_message_t, authorized_flags uint32, cache bool) (EsRespondResult, error) {
	if _es_respond_flags_result == nil {
		return *new(EsRespondResult), symbolCallError("es_respond_flags_result", "10.15", _es_respond_flags_resultErr)
	}
	return _es_respond_flags_result(client, message, authorized_flags, cache), nil
}

// Es_respond_flags_result responds to an event that requires authorization flags as a response.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_respond_flags_result(_:_:_:_:)
func Es_respond_flags_result(client *Es_client_t, message *Es_message_t, authorized_flags uint32, cache bool) EsRespondResult {
	result, callErr := tryEs_respond_flags_result(client, message, authorized_flags, cache)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_retain_message func(msg *Es_message_t)
var _es_retain_messageErr error

func tryEs_retain_message(msg *Es_message_t) error {
	if _es_retain_message == nil {
		return symbolCallError("es_retain_message", "11.0", _es_retain_messageErr)
	}
	_es_retain_message(msg)
	return nil
}

// Es_retain_message retains the given message, extending its lifetime until released.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_retain_message(_:)
func Es_retain_message(msg *Es_message_t) {
	if callErr := tryEs_retain_message(msg); callErr != nil {
		panic(callErr)
	}
}

var _es_subscribe func(client *Es_client_t, events *EsEventType, event_count uint32) EsReturn
var _es_subscribeErr error

func tryEs_subscribe(client *Es_client_t, events *EsEventType, event_count uint32) (EsReturn, error) {
	if _es_subscribe == nil {
		return *new(EsReturn), symbolCallError("es_subscribe", "10.15", _es_subscribeErr)
	}
	return _es_subscribe(client, events, event_count), nil
}

// Es_subscribe subscribes a client to a set of events.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_subscribe(_:_:_:)
func Es_subscribe(client *Es_client_t, events *EsEventType, event_count uint32) EsReturn {
	result, callErr := tryEs_subscribe(client, events, event_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_subscriptions func(client *Es_client_t, count *uintptr, subscriptions *EsEventType) EsReturn
var _es_subscriptionsErr error

func tryEs_subscriptions(client *Es_client_t, count *uintptr, subscriptions *EsEventType) (EsReturn, error) {
	if _es_subscriptions == nil {
		return *new(EsReturn), symbolCallError("es_subscriptions", "10.15", _es_subscriptionsErr)
	}
	return _es_subscriptions(client, count, subscriptions), nil
}

// Es_subscriptions returns a list of the client’s subscriptions.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_subscriptions(_:_:_:)
func Es_subscriptions(client *Es_client_t, count *uintptr, subscriptions *EsEventType) EsReturn {
	result, callErr := tryEs_subscriptions(client, count, subscriptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_unmute_all_paths func(client *Es_client_t) EsReturn
var _es_unmute_all_pathsErr error

func tryEs_unmute_all_paths(client *Es_client_t) (EsReturn, error) {
	if _es_unmute_all_paths == nil {
		return *new(EsReturn), symbolCallError("es_unmute_all_paths", "10.15", _es_unmute_all_pathsErr)
	}
	return _es_unmute_all_paths(client), nil
}

// Es_unmute_all_paths restores event delivery from previously-muted paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_all_paths(_:)
func Es_unmute_all_paths(client *Es_client_t) EsReturn {
	result, callErr := tryEs_unmute_all_paths(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_unmute_all_target_paths func(client *Es_client_t) EsReturn
var _es_unmute_all_target_pathsErr error

func tryEs_unmute_all_target_paths(client *Es_client_t) (EsReturn, error) {
	if _es_unmute_all_target_paths == nil {
		return *new(EsReturn), symbolCallError("es_unmute_all_target_paths", "13.0", _es_unmute_all_target_pathsErr)
	}
	return _es_unmute_all_target_paths(client), nil
}

// Es_unmute_all_target_paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_all_target_paths(_:)
func Es_unmute_all_target_paths(client *Es_client_t) EsReturn {
	result, callErr := tryEs_unmute_all_target_paths(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_unmute_path func(client *Es_client_t, path string, type_ EsMutePathType) EsReturn
var _es_unmute_pathErr error

func tryEs_unmute_path(client *Es_client_t, path string, type_ EsMutePathType) (EsReturn, error) {
	if _es_unmute_path == nil {
		return *new(EsReturn), symbolCallError("es_unmute_path", "12.0", _es_unmute_pathErr)
	}
	return _es_unmute_path(client, path, type_), nil
}

// Es_unmute_path restores event delivery from a previously-muted path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_path(_:_:_:)
func Es_unmute_path(client *Es_client_t, path string, type_ EsMutePathType) EsReturn {
	result, callErr := tryEs_unmute_path(client, path, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_unmute_path_events func(client *Es_client_t, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) EsReturn
var _es_unmute_path_eventsErr error

func tryEs_unmute_path_events(client *Es_client_t, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) (EsReturn, error) {
	if _es_unmute_path_events == nil {
		return *new(EsReturn), symbolCallError("es_unmute_path_events", "12.0", _es_unmute_path_eventsErr)
	}
	return _es_unmute_path_events(client, path, type_, events, event_count), nil
}

// Es_unmute_path_events restores event delivery of a subset of events from a previously-muted path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_path_events(_:_:_:_:_:)
func Es_unmute_path_events(client *Es_client_t, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) EsReturn {
	result, callErr := tryEs_unmute_path_events(client, path, type_, events, event_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_unsubscribe func(client *Es_client_t, events *EsEventType, event_count uint32) EsReturn
var _es_unsubscribeErr error

func tryEs_unsubscribe(client *Es_client_t, events *EsEventType, event_count uint32) (EsReturn, error) {
	if _es_unsubscribe == nil {
		return *new(EsReturn), symbolCallError("es_unsubscribe", "10.15", _es_unsubscribeErr)
	}
	return _es_unsubscribe(client, events, event_count), nil
}

// Es_unsubscribe unsubscribes the provided client from a set of events.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unsubscribe(_:_:_:)
func Es_unsubscribe(client *Es_client_t, events *EsEventType, event_count uint32) EsReturn {
	result, callErr := tryEs_unsubscribe(client, events, event_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _es_unsubscribe_all func(client *Es_client_t) EsReturn
var _es_unsubscribe_allErr error

func tryEs_unsubscribe_all(client *Es_client_t) (EsReturn, error) {
	if _es_unsubscribe_all == nil {
		return *new(EsReturn), symbolCallError("es_unsubscribe_all", "10.15", _es_unsubscribe_allErr)
	}
	return _es_unsubscribe_all(client), nil
}

// Es_unsubscribe_all unsubscribes a client from all events.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unsubscribe_all(_:)
func Es_unsubscribe_all(client *Es_client_t) EsReturn {
	result, callErr := tryEs_unsubscribe_all(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_es_clear_cache, &_es_clear_cacheErr, frameworkHandle, "es_clear_cache", "10.15")
	registerFunc(&_es_delete_client, &_es_delete_clientErr, frameworkHandle, "es_delete_client", "10.15")
	registerFunc(&_es_exec_arg, &_es_exec_argErr, frameworkHandle, "es_exec_arg", "10.15")
	registerFunc(&_es_exec_arg_count, &_es_exec_arg_countErr, frameworkHandle, "es_exec_arg_count", "10.15")
	registerFunc(&_es_exec_env, &_es_exec_envErr, frameworkHandle, "es_exec_env", "10.15")
	registerFunc(&_es_exec_env_count, &_es_exec_env_countErr, frameworkHandle, "es_exec_env_count", "10.15")
	registerFunc(&_es_exec_fd, &_es_exec_fdErr, frameworkHandle, "es_exec_fd", "11.0")
	registerFunc(&_es_exec_fd_count, &_es_exec_fd_countErr, frameworkHandle, "es_exec_fd_count", "11.0")
	registerFunc(&_es_invert_muting, &_es_invert_mutingErr, frameworkHandle, "es_invert_muting", "13.0")
	registerFunc(&_es_mute_path, &_es_mute_pathErr, frameworkHandle, "es_mute_path", "12.0")
	registerFunc(&_es_mute_path_events, &_es_mute_path_eventsErr, frameworkHandle, "es_mute_path_events", "12.0")
	registerFunc(&_es_muted_paths_events, &_es_muted_paths_eventsErr, frameworkHandle, "es_muted_paths_events", "12.0")
	registerFunc(&_es_muted_processes_events, &_es_muted_processes_eventsErr, frameworkHandle, "es_muted_processes_events", "12.0")
	registerFunc(&_es_muting_inverted, &_es_muting_invertedErr, frameworkHandle, "es_muting_inverted", "13.0")
	registerFunc(&_es_new_client, &_es_new_clientErr, frameworkHandle, "es_new_client", "10.15")
	registerFunc(&_es_release_message, &_es_release_messageErr, frameworkHandle, "es_release_message", "11.0")
	registerFunc(&_es_release_muted_paths, &_es_release_muted_pathsErr, frameworkHandle, "es_release_muted_paths", "12.0")
	registerFunc(&_es_release_muted_processes, &_es_release_muted_processesErr, frameworkHandle, "es_release_muted_processes", "12.0")
	registerFunc(&_es_respond_auth_result, &_es_respond_auth_resultErr, frameworkHandle, "es_respond_auth_result", "10.15")
	registerFunc(&_es_respond_flags_result, &_es_respond_flags_resultErr, frameworkHandle, "es_respond_flags_result", "10.15")
	registerFunc(&_es_retain_message, &_es_retain_messageErr, frameworkHandle, "es_retain_message", "11.0")
	registerFunc(&_es_subscribe, &_es_subscribeErr, frameworkHandle, "es_subscribe", "10.15")
	registerFunc(&_es_subscriptions, &_es_subscriptionsErr, frameworkHandle, "es_subscriptions", "10.15")
	registerFunc(&_es_unmute_all_paths, &_es_unmute_all_pathsErr, frameworkHandle, "es_unmute_all_paths", "10.15")
	registerFunc(&_es_unmute_all_target_paths, &_es_unmute_all_target_pathsErr, frameworkHandle, "es_unmute_all_target_paths", "13.0")
	registerFunc(&_es_unmute_path, &_es_unmute_pathErr, frameworkHandle, "es_unmute_path", "12.0")
	registerFunc(&_es_unmute_path_events, &_es_unmute_path_eventsErr, frameworkHandle, "es_unmute_path_events", "12.0")
	registerFunc(&_es_unsubscribe, &_es_unsubscribeErr, frameworkHandle, "es_unsubscribe", "10.15")
	registerFunc(&_es_unsubscribe_all, &_es_unsubscribe_allErr, frameworkHandle, "es_unsubscribe_all", "10.15")
}

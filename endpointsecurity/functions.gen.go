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

var _esClearCache func(client *EsClient) EsClearCacheResult
var _esClearCacheErr error

func tryEsClearCache(client *EsClient) (EsClearCacheResult, error) {
	if _esClearCache == nil {
		return *new(EsClearCacheResult), symbolCallError("es_clear_cache", "10.15", _esClearCacheErr)
	}
	return _esClearCache(client), nil
}

// EsClearCache clears all cached results for all clients.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_clear_cache(_:)
func EsClearCache(client *EsClient) EsClearCacheResult {
	result, callErr := tryEsClearCache(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esDeleteClient func(client *EsClient) EsReturn
var _esDeleteClientErr error

func tryEsDeleteClient(client *EsClient) (EsReturn, error) {
	if _esDeleteClient == nil {
		return *new(EsReturn), symbolCallError("es_delete_client", "10.15", _esDeleteClientErr)
	}
	return _esDeleteClient(client), nil
}

// EsDeleteClient destroys and disconnects a client instance from the Endpoint Security system.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_delete_client(_:)
func EsDeleteClient(client *EsClient) EsReturn {
	result, callErr := tryEsDeleteClient(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esExecArg func(event *EsEventExec, index uint32) EsStringToken
var _esExecArgErr error

func tryEsExecArg(event *EsEventExec, index uint32) (EsStringToken, error) {
	if _esExecArg == nil {
		return EsStringToken{}, symbolCallError("es_exec_arg", "10.15", _esExecArgErr)
	}
	return _esExecArg(event, index), nil
}

// EsExecArg gets the argument at the specified position from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_arg(_:_:)
func EsExecArg(event *EsEventExec, index uint32) EsStringToken {
	result, callErr := tryEsExecArg(event, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esExecArgCount func(event *EsEventExec) uint32
var _esExecArgCountErr error

func tryEsExecArgCount(event *EsEventExec) (uint32, error) {
	if _esExecArgCount == nil {
		return 0, symbolCallError("es_exec_arg_count", "10.15", _esExecArgCountErr)
	}
	return _esExecArgCount(event), nil
}

// EsExecArgCount gets the number of arguments from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_arg_count(_:)
func EsExecArgCount(event *EsEventExec) uint32 {
	result, callErr := tryEsExecArgCount(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esExecEnv func(event *EsEventExec, index uint32) EsStringToken
var _esExecEnvErr error

func tryEsExecEnv(event *EsEventExec, index uint32) (EsStringToken, error) {
	if _esExecEnv == nil {
		return EsStringToken{}, symbolCallError("es_exec_env", "10.15", _esExecEnvErr)
	}
	return _esExecEnv(event, index), nil
}

// EsExecEnv gets the environment variable at the specified position from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_env(_:_:)
func EsExecEnv(event *EsEventExec, index uint32) EsStringToken {
	result, callErr := tryEsExecEnv(event, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esExecEnvCount func(event *EsEventExec) uint32
var _esExecEnvCountErr error

func tryEsExecEnvCount(event *EsEventExec) (uint32, error) {
	if _esExecEnvCount == nil {
		return 0, symbolCallError("es_exec_env_count", "10.15", _esExecEnvCountErr)
	}
	return _esExecEnvCount(event), nil
}

// EsExecEnvCount gets the number of environment variables from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_env_count(_:)
func EsExecEnvCount(event *EsEventExec) uint32 {
	result, callErr := tryEsExecEnvCount(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esExecFd func(event *EsEventExec, index uint32) *EsFd
var _esExecFdErr error

func tryEsExecFd(event *EsEventExec, index uint32) (*EsFd, error) {
	if _esExecFd == nil {
		return nil, symbolCallError("es_exec_fd", "11.0", _esExecFdErr)
	}
	return _esExecFd(event, index), nil
}

// EsExecFd gets the file descriptor at the specified position from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_fd(_:_:)
func EsExecFd(event *EsEventExec, index uint32) *EsFd {
	result, callErr := tryEsExecFd(event, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esExecFdCount func(event *EsEventExec) uint32
var _esExecFdCountErr error

func tryEsExecFdCount(event *EsEventExec) (uint32, error) {
	if _esExecFdCount == nil {
		return 0, symbolCallError("es_exec_fd_count", "11.0", _esExecFdCountErr)
	}
	return _esExecFdCount(event), nil
}

// EsExecFdCount gets the number of file descriptors from a process execution event.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_exec_fd_count(_:)
func EsExecFdCount(event *EsEventExec) uint32 {
	result, callErr := tryEsExecFdCount(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esInvertMuting func(client *EsClient, mute_type EsMuteInversionType) EsReturn
var _esInvertMutingErr error

func tryEsInvertMuting(client *EsClient, mute_type EsMuteInversionType) (EsReturn, error) {
	if _esInvertMuting == nil {
		return *new(EsReturn), symbolCallError("es_invert_muting", "13.0", _esInvertMutingErr)
	}
	return _esInvertMuting(client, mute_type), nil
}

// EsInvertMuting.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_invert_muting(_:_:)
func EsInvertMuting(client *EsClient, mute_type EsMuteInversionType) EsReturn {
	result, callErr := tryEsInvertMuting(client, mute_type)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esMutePath func(client *EsClient, path string, type_ EsMutePathType) EsReturn
var _esMutePathErr error

func tryEsMutePath(client *EsClient, path string, type_ EsMutePathType) (EsReturn, error) {
	if _esMutePath == nil {
		return *new(EsReturn), symbolCallError("es_mute_path", "12.0", _esMutePathErr)
	}
	return _esMutePath(client, path, type_), nil
}

// EsMutePath suppresses events from executables that match a given path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path(_:_:_:)
func EsMutePath(client *EsClient, path string, type_ EsMutePathType) EsReturn {
	result, callErr := tryEsMutePath(client, path, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esMutePathEvents func(client *EsClient, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) EsReturn
var _esMutePathEventsErr error

func tryEsMutePathEvents(client *EsClient, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) (EsReturn, error) {
	if _esMutePathEvents == nil {
		return *new(EsReturn), symbolCallError("es_mute_path_events", "12.0", _esMutePathEventsErr)
	}
	return _esMutePathEvents(client, path, type_, events, event_count), nil
}

// EsMutePathEvents suppresses a subset of events from executables that match a given path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path_events(_:_:_:_:_:)
func EsMutePathEvents(client *EsClient, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) EsReturn {
	result, callErr := tryEsMutePathEvents(client, path, type_, events, event_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esMuteProcess func(client *EsClient, audit_token *[32]byte) EsReturn
var _esMuteProcessErr error

func tryEsMuteProcess(client *EsClient, audit_token *[32]byte) (EsReturn, error) {
	if _esMuteProcess == nil {
		return *new(EsReturn), symbolCallError("es_mute_process", "10.15", _esMuteProcessErr)
	}
	return _esMuteProcess(client, audit_token), nil
}

// EsMuteProcess suppresses events from a given process.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_process(_:_:)
func EsMuteProcess(client *EsClient, audit_token *[32]byte) EsReturn {
	result, callErr := tryEsMuteProcess(client, audit_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esMuteProcessEvents func(client *EsClient, audit_token *[32]byte, events *EsEventType, event_count uintptr) EsReturn
var _esMuteProcessEventsErr error

func tryEsMuteProcessEvents(client *EsClient, audit_token *[32]byte, events *EsEventType, event_count uintptr) (EsReturn, error) {
	if _esMuteProcessEvents == nil {
		return *new(EsReturn), symbolCallError("es_mute_process_events", "12.0", _esMuteProcessEventsErr)
	}
	return _esMuteProcessEvents(client, audit_token, events, event_count), nil
}

// EsMuteProcessEvents suppresses a subset of events from a given process.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_process_events(_:_:_:_:)
func EsMuteProcessEvents(client *EsClient, audit_token *[32]byte, events *EsEventType, event_count uintptr) EsReturn {
	result, callErr := tryEsMuteProcessEvents(client, audit_token, events, event_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esMutedPathsEvents func(client *EsClient, muted_paths **EsMutedPaths) EsReturn
var _esMutedPathsEventsErr error

func tryEsMutedPathsEvents(client *EsClient, muted_paths **EsMutedPaths) (EsReturn, error) {
	if _esMutedPathsEvents == nil {
		return *new(EsReturn), symbolCallError("es_muted_paths_events", "12.0", _esMutedPathsEventsErr)
	}
	return _esMutedPathsEvents(client, muted_paths), nil
}

// EsMutedPathsEvents retrieve a list of all muted paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_muted_paths_events(_:_:)
func EsMutedPathsEvents(client *EsClient, muted_paths **EsMutedPaths) EsReturn {
	result, callErr := tryEsMutedPathsEvents(client, muted_paths)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esMutedProcessesEvents func(client *EsClient, muted_processes **EsMutedProcesses) EsReturn
var _esMutedProcessesEventsErr error

func tryEsMutedProcessesEvents(client *EsClient, muted_processes **EsMutedProcesses) (EsReturn, error) {
	if _esMutedProcessesEvents == nil {
		return *new(EsReturn), symbolCallError("es_muted_processes_events", "12.0", _esMutedProcessesEventsErr)
	}
	return _esMutedProcessesEvents(client, muted_processes), nil
}

// EsMutedProcessesEvents retrieve a list of all muted processes.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_muted_processes_events(_:_:)
func EsMutedProcessesEvents(client *EsClient, muted_processes **EsMutedProcesses) EsReturn {
	result, callErr := tryEsMutedProcessesEvents(client, muted_processes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esMutingInverted func(client *EsClient, mute_type EsMuteInversionType) EsMuteInvertedReturn
var _esMutingInvertedErr error

func tryEsMutingInverted(client *EsClient, mute_type EsMuteInversionType) (EsMuteInvertedReturn, error) {
	if _esMutingInverted == nil {
		return *new(EsMuteInvertedReturn), symbolCallError("es_muting_inverted", "13.0", _esMutingInvertedErr)
	}
	return _esMutingInverted(client, mute_type), nil
}

// EsMutingInverted.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_muting_inverted(_:_:)
func EsMutingInverted(client *EsClient, mute_type EsMuteInversionType) EsMuteInvertedReturn {
	result, callErr := tryEsMutingInverted(client, mute_type)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esNewClient func(client **Es_client_t, handler unsafe.Pointer) EsNewClientResult
var _esNewClientErr error

func tryEsNewClient(client **Es_client_t, handler func(*Es_client_t, *Es_message_t)) (EsNewClientResult, error) {
	if _esNewClient == nil {
		return *new(EsNewClientResult), symbolCallError("es_new_client", "10.15", _esNewClientErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *Es_client_t, blockArg1 *Es_message_t) { handler(blockArg0, blockArg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _esNewClient(client, _block0), nil
}

// EsNewClient creates a new client instance and connects it to the Endpoint Security system.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_new_client(_:_:)
func EsNewClient(client **Es_client_t, handler func(*Es_client_t, *Es_message_t)) EsNewClientResult {
	result, callErr := tryEsNewClient(client, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esReleaseMessage func(msg *EsMessage)
var _esReleaseMessageErr error

func tryEsReleaseMessage(msg *EsMessage) error {
	if _esReleaseMessage == nil {
		return symbolCallError("es_release_message", "11.0", _esReleaseMessageErr)
	}
	_esReleaseMessage(msg)
	return nil
}

// EsReleaseMessage releases a previously-retained message.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_release_message(_:)
func EsReleaseMessage(msg *EsMessage) {
	if callErr := tryEsReleaseMessage(msg); callErr != nil {
		panic(callErr)
	}
}

var _esReleaseMutedPaths func(muted_paths *EsMutedPaths)
var _esReleaseMutedPathsErr error

func tryEsReleaseMutedPaths(muted_paths *EsMutedPaths) error {
	if _esReleaseMutedPaths == nil {
		return symbolCallError("es_release_muted_paths", "12.0", _esReleaseMutedPathsErr)
	}
	_esReleaseMutedPaths(muted_paths)
	return nil
}

// EsReleaseMutedPaths frees resources associated with a set of previously-retrieved muted paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_release_muted_paths(_:)
func EsReleaseMutedPaths(muted_paths *EsMutedPaths) {
	if callErr := tryEsReleaseMutedPaths(muted_paths); callErr != nil {
		panic(callErr)
	}
}

var _esReleaseMutedProcesses func(muted_processes *EsMutedProcesses)
var _esReleaseMutedProcessesErr error

func tryEsReleaseMutedProcesses(muted_processes *EsMutedProcesses) error {
	if _esReleaseMutedProcesses == nil {
		return symbolCallError("es_release_muted_processes", "12.0", _esReleaseMutedProcessesErr)
	}
	_esReleaseMutedProcesses(muted_processes)
	return nil
}

// EsReleaseMutedProcesses frees resources associated with a set of previously-retrieved muted processes.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_release_muted_processes(_:)
func EsReleaseMutedProcesses(muted_processes *EsMutedProcesses) {
	if callErr := tryEsReleaseMutedProcesses(muted_processes); callErr != nil {
		panic(callErr)
	}
}

var _esRespondAuthResult func(client *EsClient, message *EsMessage, result EsAuthResult, cache bool) EsRespondResult
var _esRespondAuthResultErr error

func tryEsRespondAuthResult(client *EsClient, message *EsMessage, result EsAuthResult, cache bool) (EsRespondResult, error) {
	if _esRespondAuthResult == nil {
		return *new(EsRespondResult), symbolCallError("es_respond_auth_result", "10.15", _esRespondAuthResultErr)
	}
	return _esRespondAuthResult(client, message, result, cache), nil
}

// EsRespondAuthResult responds to an event that requires an authorization response.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_respond_auth_result(_:_:_:_:)
func EsRespondAuthResult(client *EsClient, message *EsMessage, result EsAuthResult, cache bool) EsRespondResult {
	result0, callErr := tryEsRespondAuthResult(client, message, result, cache)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _esRespondFlagsResult func(client *EsClient, message *EsMessage, authorized_flags uint32, cache bool) EsRespondResult
var _esRespondFlagsResultErr error

func tryEsRespondFlagsResult(client *EsClient, message *EsMessage, authorized_flags uint32, cache bool) (EsRespondResult, error) {
	if _esRespondFlagsResult == nil {
		return *new(EsRespondResult), symbolCallError("es_respond_flags_result", "10.15", _esRespondFlagsResultErr)
	}
	return _esRespondFlagsResult(client, message, authorized_flags, cache), nil
}

// EsRespondFlagsResult responds to an event that requires authorization flags as a response.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_respond_flags_result(_:_:_:_:)
func EsRespondFlagsResult(client *EsClient, message *EsMessage, authorized_flags uint32, cache bool) EsRespondResult {
	result, callErr := tryEsRespondFlagsResult(client, message, authorized_flags, cache)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esRetainMessage func(msg *EsMessage)
var _esRetainMessageErr error

func tryEsRetainMessage(msg *EsMessage) error {
	if _esRetainMessage == nil {
		return symbolCallError("es_retain_message", "11.0", _esRetainMessageErr)
	}
	_esRetainMessage(msg)
	return nil
}

// EsRetainMessage retains the given message, extending its lifetime until released.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_retain_message(_:)
func EsRetainMessage(msg *EsMessage) {
	if callErr := tryEsRetainMessage(msg); callErr != nil {
		panic(callErr)
	}
}

var _esSubscribe func(client *EsClient, events *EsEventType, event_count uint32) EsReturn
var _esSubscribeErr error

func tryEsSubscribe(client *EsClient, events *EsEventType, event_count uint32) (EsReturn, error) {
	if _esSubscribe == nil {
		return *new(EsReturn), symbolCallError("es_subscribe", "10.15", _esSubscribeErr)
	}
	return _esSubscribe(client, events, event_count), nil
}

// EsSubscribe subscribes a client to a set of events.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_subscribe(_:_:_:)
func EsSubscribe(client *EsClient, events *EsEventType, event_count uint32) EsReturn {
	result, callErr := tryEsSubscribe(client, events, event_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esSubscriptions func(client *EsClient, count *uintptr, subscriptions *EsEventType) EsReturn
var _esSubscriptionsErr error

func tryEsSubscriptions(client *EsClient, count *uintptr, subscriptions *EsEventType) (EsReturn, error) {
	if _esSubscriptions == nil {
		return *new(EsReturn), symbolCallError("es_subscriptions", "10.15", _esSubscriptionsErr)
	}
	return _esSubscriptions(client, count, subscriptions), nil
}

// EsSubscriptions returns a list of the client’s subscriptions.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_subscriptions(_:_:_:)
func EsSubscriptions(client *EsClient, count *uintptr, subscriptions *EsEventType) EsReturn {
	result, callErr := tryEsSubscriptions(client, count, subscriptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esUnmuteAllPaths func(client *EsClient) EsReturn
var _esUnmuteAllPathsErr error

func tryEsUnmuteAllPaths(client *EsClient) (EsReturn, error) {
	if _esUnmuteAllPaths == nil {
		return *new(EsReturn), symbolCallError("es_unmute_all_paths", "10.15", _esUnmuteAllPathsErr)
	}
	return _esUnmuteAllPaths(client), nil
}

// EsUnmuteAllPaths restores event delivery from previously-muted paths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_all_paths(_:)
func EsUnmuteAllPaths(client *EsClient) EsReturn {
	result, callErr := tryEsUnmuteAllPaths(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esUnmuteAllTargetPaths func(client *EsClient) EsReturn
var _esUnmuteAllTargetPathsErr error

func tryEsUnmuteAllTargetPaths(client *EsClient) (EsReturn, error) {
	if _esUnmuteAllTargetPaths == nil {
		return *new(EsReturn), symbolCallError("es_unmute_all_target_paths", "13.0", _esUnmuteAllTargetPathsErr)
	}
	return _esUnmuteAllTargetPaths(client), nil
}

// EsUnmuteAllTargetPaths.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_all_target_paths(_:)
func EsUnmuteAllTargetPaths(client *EsClient) EsReturn {
	result, callErr := tryEsUnmuteAllTargetPaths(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esUnmutePath func(client *EsClient, path string, type_ EsMutePathType) EsReturn
var _esUnmutePathErr error

func tryEsUnmutePath(client *EsClient, path string, type_ EsMutePathType) (EsReturn, error) {
	if _esUnmutePath == nil {
		return *new(EsReturn), symbolCallError("es_unmute_path", "12.0", _esUnmutePathErr)
	}
	return _esUnmutePath(client, path, type_), nil
}

// EsUnmutePath restores event delivery from a previously-muted path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_path(_:_:_:)
func EsUnmutePath(client *EsClient, path string, type_ EsMutePathType) EsReturn {
	result, callErr := tryEsUnmutePath(client, path, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esUnmutePathEvents func(client *EsClient, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) EsReturn
var _esUnmutePathEventsErr error

func tryEsUnmutePathEvents(client *EsClient, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) (EsReturn, error) {
	if _esUnmutePathEvents == nil {
		return *new(EsReturn), symbolCallError("es_unmute_path_events", "12.0", _esUnmutePathEventsErr)
	}
	return _esUnmutePathEvents(client, path, type_, events, event_count), nil
}

// EsUnmutePathEvents restores event delivery of a subset of events from a previously-muted path.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_path_events(_:_:_:_:_:)
func EsUnmutePathEvents(client *EsClient, path string, type_ EsMutePathType, events *EsEventType, event_count uintptr) EsReturn {
	result, callErr := tryEsUnmutePathEvents(client, path, type_, events, event_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esUnmuteProcess func(client *EsClient, audit_token *[32]byte) EsReturn
var _esUnmuteProcessErr error

func tryEsUnmuteProcess(client *EsClient, audit_token *[32]byte) (EsReturn, error) {
	if _esUnmuteProcess == nil {
		return *new(EsReturn), symbolCallError("es_unmute_process", "10.15", _esUnmuteProcessErr)
	}
	return _esUnmuteProcess(client, audit_token), nil
}

// EsUnmuteProcess restores event delivery from a previously-muted process.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_process(_:_:)
func EsUnmuteProcess(client *EsClient, audit_token *[32]byte) EsReturn {
	result, callErr := tryEsUnmuteProcess(client, audit_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esUnmuteProcessEvents func(client *EsClient, audit_token *[32]byte, events *EsEventType, event_count uintptr) EsReturn
var _esUnmuteProcessEventsErr error

func tryEsUnmuteProcessEvents(client *EsClient, audit_token *[32]byte, events *EsEventType, event_count uintptr) (EsReturn, error) {
	if _esUnmuteProcessEvents == nil {
		return *new(EsReturn), symbolCallError("es_unmute_process_events", "12.0", _esUnmuteProcessEventsErr)
	}
	return _esUnmuteProcessEvents(client, audit_token, events, event_count), nil
}

// EsUnmuteProcessEvents restores event delivery of a subset of events from a previously-muted process.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_process_events(_:_:_:_:)
func EsUnmuteProcessEvents(client *EsClient, audit_token *[32]byte, events *EsEventType, event_count uintptr) EsReturn {
	result, callErr := tryEsUnmuteProcessEvents(client, audit_token, events, event_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esUnsubscribe func(client *EsClient, events *EsEventType, event_count uint32) EsReturn
var _esUnsubscribeErr error

func tryEsUnsubscribe(client *EsClient, events *EsEventType, event_count uint32) (EsReturn, error) {
	if _esUnsubscribe == nil {
		return *new(EsReturn), symbolCallError("es_unsubscribe", "10.15", _esUnsubscribeErr)
	}
	return _esUnsubscribe(client, events, event_count), nil
}

// EsUnsubscribe unsubscribes the provided client from a set of events.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unsubscribe(_:_:_:)
func EsUnsubscribe(client *EsClient, events *EsEventType, event_count uint32) EsReturn {
	result, callErr := tryEsUnsubscribe(client, events, event_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _esUnsubscribeAll func(client *EsClient) EsReturn
var _esUnsubscribeAllErr error

func tryEsUnsubscribeAll(client *EsClient) (EsReturn, error) {
	if _esUnsubscribeAll == nil {
		return *new(EsReturn), symbolCallError("es_unsubscribe_all", "10.15", _esUnsubscribeAllErr)
	}
	return _esUnsubscribeAll(client), nil
}

// EsUnsubscribeAll unsubscribes a client from all events.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_unsubscribe_all(_:)
func EsUnsubscribeAll(client *EsClient) EsReturn {
	result, callErr := tryEsUnsubscribeAll(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_esClearCache, &_esClearCacheErr, frameworkHandle, "es_clear_cache", "10.15")
	registerFunc(&_esDeleteClient, &_esDeleteClientErr, frameworkHandle, "es_delete_client", "10.15")
	registerFunc(&_esExecArg, &_esExecArgErr, frameworkHandle, "es_exec_arg", "10.15")
	registerFunc(&_esExecArgCount, &_esExecArgCountErr, frameworkHandle, "es_exec_arg_count", "10.15")
	registerFunc(&_esExecEnv, &_esExecEnvErr, frameworkHandle, "es_exec_env", "10.15")
	registerFunc(&_esExecEnvCount, &_esExecEnvCountErr, frameworkHandle, "es_exec_env_count", "10.15")
	registerFunc(&_esExecFd, &_esExecFdErr, frameworkHandle, "es_exec_fd", "11.0")
	registerFunc(&_esExecFdCount, &_esExecFdCountErr, frameworkHandle, "es_exec_fd_count", "11.0")
	registerFunc(&_esInvertMuting, &_esInvertMutingErr, frameworkHandle, "es_invert_muting", "13.0")
	registerFunc(&_esMutePath, &_esMutePathErr, frameworkHandle, "es_mute_path", "12.0")
	registerFunc(&_esMutePathEvents, &_esMutePathEventsErr, frameworkHandle, "es_mute_path_events", "12.0")
	registerFunc(&_esMuteProcess, &_esMuteProcessErr, frameworkHandle, "es_mute_process", "10.15")
	registerFunc(&_esMuteProcessEvents, &_esMuteProcessEventsErr, frameworkHandle, "es_mute_process_events", "12.0")
	registerFunc(&_esMutedPathsEvents, &_esMutedPathsEventsErr, frameworkHandle, "es_muted_paths_events", "12.0")
	registerFunc(&_esMutedProcessesEvents, &_esMutedProcessesEventsErr, frameworkHandle, "es_muted_processes_events", "12.0")
	registerFunc(&_esMutingInverted, &_esMutingInvertedErr, frameworkHandle, "es_muting_inverted", "13.0")
	registerFunc(&_esNewClient, &_esNewClientErr, frameworkHandle, "es_new_client", "10.15")
	registerFunc(&_esReleaseMessage, &_esReleaseMessageErr, frameworkHandle, "es_release_message", "11.0")
	registerFunc(&_esReleaseMutedPaths, &_esReleaseMutedPathsErr, frameworkHandle, "es_release_muted_paths", "12.0")
	registerFunc(&_esReleaseMutedProcesses, &_esReleaseMutedProcessesErr, frameworkHandle, "es_release_muted_processes", "12.0")
	registerFunc(&_esRespondAuthResult, &_esRespondAuthResultErr, frameworkHandle, "es_respond_auth_result", "10.15")
	registerFunc(&_esRespondFlagsResult, &_esRespondFlagsResultErr, frameworkHandle, "es_respond_flags_result", "10.15")
	registerFunc(&_esRetainMessage, &_esRetainMessageErr, frameworkHandle, "es_retain_message", "11.0")
	registerFunc(&_esSubscribe, &_esSubscribeErr, frameworkHandle, "es_subscribe", "10.15")
	registerFunc(&_esSubscriptions, &_esSubscriptionsErr, frameworkHandle, "es_subscriptions", "10.15")
	registerFunc(&_esUnmuteAllPaths, &_esUnmuteAllPathsErr, frameworkHandle, "es_unmute_all_paths", "10.15")
	registerFunc(&_esUnmuteAllTargetPaths, &_esUnmuteAllTargetPathsErr, frameworkHandle, "es_unmute_all_target_paths", "13.0")
	registerFunc(&_esUnmutePath, &_esUnmutePathErr, frameworkHandle, "es_unmute_path", "12.0")
	registerFunc(&_esUnmutePathEvents, &_esUnmutePathEventsErr, frameworkHandle, "es_unmute_path_events", "12.0")
	registerFunc(&_esUnmuteProcess, &_esUnmuteProcessErr, frameworkHandle, "es_unmute_process", "10.15")
	registerFunc(&_esUnmuteProcessEvents, &_esUnmuteProcessEventsErr, frameworkHandle, "es_unmute_process_events", "12.0")
	registerFunc(&_esUnsubscribe, &_esUnsubscribeErr, frameworkHandle, "es_unsubscribe", "10.15")
	registerFunc(&_esUnsubscribeAll, &_esUnsubscribeAllErr, frameworkHandle, "es_unsubscribe_all", "10.15")
}

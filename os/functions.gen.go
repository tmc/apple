// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/kernel"
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
		return fmt.Sprintf("os: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("os: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("os: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("os: register symbol %s: %v", name, r)
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

var _os_activity_apply func(activity Os_activity_t, block unsafe.Pointer)
var _os_activity_applyErr error

func tryOs_activity_apply(activity Os_activity_t, block Os_block_t) error {
	if _os_activity_apply == nil {
		return symbolCallError("os_activity_apply", "10.12", _os_activity_applyErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block) { block() })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_os_activity_apply(activity, _block0)
	return nil
}

// Os_activity_apply execute a block using a given activity object.
//
// See: https://developer.apple.com/documentation/os/os_activity_apply
func Os_activity_apply(activity Os_activity_t, block Os_block_t) {
	if callErr := tryOs_activity_apply(activity, block); callErr != nil {
		panic(callErr)
	}
}

var _os_activity_apply_f func(activity Os_activity_t, context unsafe.Pointer, function Os_function_t)
var _os_activity_apply_fErr error

func tryOs_activity_apply_f(activity Os_activity_t, context unsafe.Pointer, function Os_function_t) error {
	if _os_activity_apply_f == nil {
		return symbolCallError("os_activity_apply_f", "10.12", _os_activity_apply_fErr)
	}
	_os_activity_apply_f(activity, context, function)
	return nil
}

// Os_activity_apply_f execute a function using a given activity object.
//
// See: https://developer.apple.com/documentation/os/os_activity_apply_f
func Os_activity_apply_f(activity Os_activity_t, context unsafe.Pointer, function Os_function_t) {
	if callErr := tryOs_activity_apply_f(activity, context, function); callErr != nil {
		panic(callErr)
	}
}

var _os_activity_get_identifier func(activity Os_activity_t, parent_id *Os_activity_id_t) Os_activity_id_t
var _os_activity_get_identifierErr error

func tryOs_activity_get_identifier(activity Os_activity_t, parent_id *Os_activity_id_t) (Os_activity_id_t, error) {
	if _os_activity_get_identifier == nil {
		return *new(Os_activity_id_t), symbolCallError("os_activity_get_identifier", "10.12", _os_activity_get_identifierErr)
	}
	return _os_activity_get_identifier(activity, parent_id), nil
}

// Os_activity_get_identifier retrieves the identifier for a given activity object.
//
// See: https://developer.apple.com/documentation/os/os_activity_get_identifier
func Os_activity_get_identifier(activity Os_activity_t, parent_id *Os_activity_id_t) Os_activity_id_t {
	result, callErr := tryOs_activity_get_identifier(activity, parent_id)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_activity_scope_enter func(activity Os_activity_t, state Os_activity_scope_state_t)
var _os_activity_scope_enterErr error

func tryOs_activity_scope_enter(activity Os_activity_t, state Os_activity_scope_state_t) error {
	if _os_activity_scope_enter == nil {
		return symbolCallError("os_activity_scope_enter", "10.12", _os_activity_scope_enterErr)
	}
	_os_activity_scope_enter(activity, state)
	return nil
}

// Os_activity_scope_enter switches the current activity, saving the existing execution context.
//
// See: https://developer.apple.com/documentation/os/os_activity_scope_enter
func Os_activity_scope_enter(activity Os_activity_t, state Os_activity_scope_state_t) {
	if callErr := tryOs_activity_scope_enter(activity, state); callErr != nil {
		panic(callErr)
	}
}

var _os_activity_scope_leave func(state Os_activity_scope_state_t)
var _os_activity_scope_leaveErr error

func tryOs_activity_scope_leave(state Os_activity_scope_state_t) error {
	if _os_activity_scope_leave == nil {
		return symbolCallError("os_activity_scope_leave", "10.12", _os_activity_scope_leaveErr)
	}
	_os_activity_scope_leave(state)
	return nil
}

// Os_activity_scope_leave restores the current activity to a previously saved state.
//
// See: https://developer.apple.com/documentation/os/os_activity_scope_leave
func Os_activity_scope_leave(state Os_activity_scope_state_t) {
	if callErr := tryOs_activity_scope_leave(state); callErr != nil {
		panic(callErr)
	}
}

var _os_log_create func(subsystem string, category string) Os_log_t
var _os_log_createErr error

func tryOs_log_create(subsystem string, category string) (Os_log_t, error) {
	if _os_log_create == nil {
		return *new(Os_log_t), symbolCallError("os_log_create", "10.12", _os_log_createErr)
	}
	return _os_log_create(subsystem, category), nil
}

// Os_log_create creates a custom log object.
//
// See: https://developer.apple.com/documentation/os/os_log_create
func Os_log_create(subsystem string, category string) Os_log_t {
	result, callErr := tryOs_log_create(subsystem, category)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_log_type_enabled func(oslog Os_log_t, type_ Os_log_type_t) bool
var _os_log_type_enabledErr error

func tryOs_log_type_enabled(oslog Os_log_t, type_ Os_log_type_t) (bool, error) {
	if _os_log_type_enabled == nil {
		return false, symbolCallError("os_log_type_enabled", "10.12", _os_log_type_enabledErr)
	}
	return _os_log_type_enabled(oslog, type_), nil
}

// Os_log_type_enabled returns a Boolean value that indicates whether the log can write messages with the specified log type.
//
// See: https://developer.apple.com/documentation/os/OSLog/isEnabled(type:)
func Os_log_type_enabled(oslog Os_log_t, type_ Os_log_type_t) bool {
	result, callErr := tryOs_log_type_enabled(oslog, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_release func(object unsafe.Pointer)
var _os_releaseErr error

func tryOs_release(object unsafe.Pointer) error {
	if _os_release == nil {
		return symbolCallError("os_release", "10.10", _os_releaseErr)
	}
	_os_release(object)
	return nil
}

// Os_release.
//
// See: https://developer.apple.com/documentation/os/os_release-c.func
func Os_release(object unsafe.Pointer) {
	if callErr := tryOs_release(object); callErr != nil {
		panic(callErr)
	}
}

var _os_retain func(object unsafe.Pointer) unsafe.Pointer
var _os_retainErr error

func tryOs_retain(object unsafe.Pointer) (unsafe.Pointer, error) {
	if _os_retain == nil {
		return nil, symbolCallError("os_retain", "10.10", _os_retainErr)
	}
	return _os_retain(object), nil
}

// Os_retain.
//
// See: https://developer.apple.com/documentation/os/os_retain-c.func
func Os_retain(object unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryOs_retain(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_security_config_get func() OsSecurityConfig
var _os_security_config_getErr error

func tryOs_security_config_get() (OsSecurityConfig, error) {
	if _os_security_config_get == nil {
		return *new(OsSecurityConfig), symbolCallError("os_security_config_get", "26.0", _os_security_config_getErr)
	}
	return _os_security_config_get(), nil
}

// Os_security_config_get.
//
// See: https://developer.apple.com/documentation/os/os_security_config_get
func Os_security_config_get() OsSecurityConfig {
	result, callErr := tryOs_security_config_get()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_security_config_get_for_proc func(pid int32, config *OsSecurityConfig) int
var _os_security_config_get_for_procErr error

func tryOs_security_config_get_for_proc(pid int32, config *OsSecurityConfig) (int, error) {
	if _os_security_config_get_for_proc == nil {
		return 0, symbolCallError("os_security_config_get_for_proc", "26.0", _os_security_config_get_for_procErr)
	}
	return _os_security_config_get_for_proc(pid, config), nil
}

// Os_security_config_get_for_proc.
//
// See: https://developer.apple.com/documentation/os/os_security_config_get_for_proc
func Os_security_config_get_for_proc(pid int32, config *OsSecurityConfig) int {
	result, callErr := tryOs_security_config_get_for_proc(pid, config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_security_config_get_for_task func(task kernel.Task_t, config *OsSecurityConfig) int
var _os_security_config_get_for_taskErr error

func tryOs_security_config_get_for_task(task kernel.Task_t, config *OsSecurityConfig) (int, error) {
	if _os_security_config_get_for_task == nil {
		return 0, symbolCallError("os_security_config_get_for_task", "26.0", _os_security_config_get_for_taskErr)
	}
	return _os_security_config_get_for_task(task, config), nil
}

// Os_security_config_get_for_task.
//
// See: https://developer.apple.com/documentation/os/os_security_config_get_for_task
func Os_security_config_get_for_task(task kernel.Task_t, config *OsSecurityConfig) int {
	result, callErr := tryOs_security_config_get_for_task(task, config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_signpost_enabled func(log Os_log_t) bool
var _os_signpost_enabledErr error

func tryOs_signpost_enabled(log Os_log_t) (bool, error) {
	if _os_signpost_enabled == nil {
		return false, symbolCallError("os_signpost_enabled", "10.14", _os_signpost_enabledErr)
	}
	return _os_signpost_enabled(log), nil
}

// Os_signpost_enabled returns a Boolean value that indicates whether signposts are in an enabled state for the specified log.
//
// See: https://developer.apple.com/documentation/os/os_signpost_enabled
func Os_signpost_enabled(log Os_log_t) bool {
	result, callErr := tryOs_signpost_enabled(log)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_signpost_id_generate func(log Os_log_t) Os_signpost_id_t
var _os_signpost_id_generateErr error

func tryOs_signpost_id_generate(log Os_log_t) (Os_signpost_id_t, error) {
	if _os_signpost_id_generate == nil {
		return *new(Os_signpost_id_t), symbolCallError("os_signpost_id_generate", "10.14", _os_signpost_id_generateErr)
	}
	return _os_signpost_id_generate(log), nil
}

// Os_signpost_id_generate creates a signpost identifier that’s unique among signposts logged to a specified log.
//
// See: https://developer.apple.com/documentation/os/os_signpost_id_generate
func Os_signpost_id_generate(log Os_log_t) Os_signpost_id_t {
	result, callErr := tryOs_signpost_id_generate(log)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_signpost_id_make_with_pointer func(log Os_log_t, ptr unsafe.Pointer) Os_signpost_id_t
var _os_signpost_id_make_with_pointerErr error

func tryOs_signpost_id_make_with_pointer(log Os_log_t, ptr unsafe.Pointer) (Os_signpost_id_t, error) {
	if _os_signpost_id_make_with_pointer == nil {
		return *new(Os_signpost_id_t), symbolCallError("os_signpost_id_make_with_pointer", "10.14", _os_signpost_id_make_with_pointerErr)
	}
	return _os_signpost_id_make_with_pointer(log, ptr), nil
}

// Os_signpost_id_make_with_pointer creates a signpost identifier that’s unique among signposts logging to the specified log, using a pointer value to generate the unique value.
//
// See: https://developer.apple.com/documentation/os/os_signpost_id_make_with_pointer
func Os_signpost_id_make_with_pointer(log Os_log_t, ptr unsafe.Pointer) Os_signpost_id_t {
	result, callErr := tryOs_signpost_id_make_with_pointer(log, ptr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_sync_wait_on_address func(addr unsafe.Pointer, value uint64, size uintptr, flags OsSyncWaitOnAddress) int
var _os_sync_wait_on_addressErr error

func tryOs_sync_wait_on_address(addr unsafe.Pointer, value uint64, size uintptr, flags OsSyncWaitOnAddress) (int, error) {
	if _os_sync_wait_on_address == nil {
		return 0, symbolCallError("os_sync_wait_on_address", "14.4", _os_sync_wait_on_addressErr)
	}
	return _os_sync_wait_on_address(addr, value, size, flags), nil
}

// Os_sync_wait_on_address an atomic compare-and-wait operation, used to implement higher-level synchronization primitives.
//
// See: https://developer.apple.com/documentation/os/os_sync_wait_on_address
func Os_sync_wait_on_address(addr unsafe.Pointer, value uint64, size uintptr, flags OsSyncWaitOnAddress) int {
	result, callErr := tryOs_sync_wait_on_address(addr, value, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_sync_wait_on_address_with_deadline func(addr unsafe.Pointer, value uint64, size uintptr, flags OsSyncWaitOnAddress, clockid OsClockMachAbsolute, deadline uint64) int
var _os_sync_wait_on_address_with_deadlineErr error

func tryOs_sync_wait_on_address_with_deadline(addr unsafe.Pointer, value uint64, size uintptr, flags OsSyncWaitOnAddress, clockid OsClockMachAbsolute, deadline uint64) (int, error) {
	if _os_sync_wait_on_address_with_deadline == nil {
		return 0, symbolCallError("os_sync_wait_on_address_with_deadline", "14.4", _os_sync_wait_on_address_with_deadlineErr)
	}
	return _os_sync_wait_on_address_with_deadline(addr, value, size, flags, clockid, deadline), nil
}

// Os_sync_wait_on_address_with_deadline an atomic compare-and-wait operation with a deadline, used to implement higher-level synchronization primitives.
//
// See: https://developer.apple.com/documentation/os/os_sync_wait_on_address_with_deadline
func Os_sync_wait_on_address_with_deadline(addr unsafe.Pointer, value uint64, size uintptr, flags OsSyncWaitOnAddress, clockid OsClockMachAbsolute, deadline uint64) int {
	result, callErr := tryOs_sync_wait_on_address_with_deadline(addr, value, size, flags, clockid, deadline)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_sync_wait_on_address_with_timeout func(addr unsafe.Pointer, value uint64, size uintptr, flags OsSyncWaitOnAddress, clockid OsClockMachAbsolute, timeout_ns uint64) int
var _os_sync_wait_on_address_with_timeoutErr error

func tryOs_sync_wait_on_address_with_timeout(addr unsafe.Pointer, value uint64, size uintptr, flags OsSyncWaitOnAddress, clockid OsClockMachAbsolute, timeout_ns uint64) (int, error) {
	if _os_sync_wait_on_address_with_timeout == nil {
		return 0, symbolCallError("os_sync_wait_on_address_with_timeout", "14.4", _os_sync_wait_on_address_with_timeoutErr)
	}
	return _os_sync_wait_on_address_with_timeout(addr, value, size, flags, clockid, timeout_ns), nil
}

// Os_sync_wait_on_address_with_timeout an atomic compare-and-wait operation with a timeout, used to implement higher-level synchronization primitives.
//
// See: https://developer.apple.com/documentation/os/os_sync_wait_on_address_with_timeout
func Os_sync_wait_on_address_with_timeout(addr unsafe.Pointer, value uint64, size uintptr, flags OsSyncWaitOnAddress, clockid OsClockMachAbsolute, timeout_ns uint64) int {
	result, callErr := tryOs_sync_wait_on_address_with_timeout(addr, value, size, flags, clockid, timeout_ns)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_sync_wake_by_address_all func(addr unsafe.Pointer, size uintptr, flags OsSyncWakeByAddress) int
var _os_sync_wake_by_address_allErr error

func tryOs_sync_wake_by_address_all(addr unsafe.Pointer, size uintptr, flags OsSyncWakeByAddress) (int, error) {
	if _os_sync_wake_by_address_all == nil {
		return 0, symbolCallError("os_sync_wake_by_address_all", "14.4", _os_sync_wake_by_address_allErr)
	}
	return _os_sync_wake_by_address_all(addr, size, flags), nil
}

// Os_sync_wake_by_address_all an atomic operation that wakes all threads blocked on a futex wait, used to implement higher-level synchronization primitives.
//
// See: https://developer.apple.com/documentation/os/os_sync_wake_by_address_all
func Os_sync_wake_by_address_all(addr unsafe.Pointer, size uintptr, flags OsSyncWakeByAddress) int {
	result, callErr := tryOs_sync_wake_by_address_all(addr, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_sync_wake_by_address_any func(addr unsafe.Pointer, size uintptr, flags OsSyncWakeByAddress) int
var _os_sync_wake_by_address_anyErr error

func tryOs_sync_wake_by_address_any(addr unsafe.Pointer, size uintptr, flags OsSyncWakeByAddress) (int, error) {
	if _os_sync_wake_by_address_any == nil {
		return 0, symbolCallError("os_sync_wake_by_address_any", "14.4", _os_sync_wake_by_address_anyErr)
	}
	return _os_sync_wake_by_address_any(addr, size, flags), nil
}

// Os_sync_wake_by_address_any an atomic operation that wakes one thread blocked on a futex wait, used to implement higher-level synchronization primitives.
//
// See: https://developer.apple.com/documentation/os/os_sync_wake_by_address_any
func Os_sync_wake_by_address_any(addr unsafe.Pointer, size uintptr, flags OsSyncWakeByAddress) int {
	result, callErr := tryOs_sync_wake_by_address_any(addr, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_unfair_lock_assert_not_owner func(lock *[4]byte)
var _os_unfair_lock_assert_not_ownerErr error

func tryOs_unfair_lock_assert_not_owner(lock *[4]byte) error {
	if _os_unfair_lock_assert_not_owner == nil {
		return symbolCallError("os_unfair_lock_assert_not_owner", "10.12", _os_unfair_lock_assert_not_ownerErr)
	}
	_os_unfair_lock_assert_not_owner(lock)
	return nil
}

// Os_unfair_lock_assert_not_owner triggers an assertion if the calling thread owns the specified unfair lock.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_assert_not_owner
func Os_unfair_lock_assert_not_owner(lock *[4]byte) {
	if callErr := tryOs_unfair_lock_assert_not_owner(lock); callErr != nil {
		panic(callErr)
	}
}

var _os_unfair_lock_assert_owner func(lock *[4]byte)
var _os_unfair_lock_assert_ownerErr error

func tryOs_unfair_lock_assert_owner(lock *[4]byte) error {
	if _os_unfair_lock_assert_owner == nil {
		return symbolCallError("os_unfair_lock_assert_owner", "10.12", _os_unfair_lock_assert_ownerErr)
	}
	_os_unfair_lock_assert_owner(lock)
	return nil
}

// Os_unfair_lock_assert_owner triggers an assertion if the calling thread doesn’t own the specified unfair lock.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_assert_owner
func Os_unfair_lock_assert_owner(lock *[4]byte) {
	if callErr := tryOs_unfair_lock_assert_owner(lock); callErr != nil {
		panic(callErr)
	}
}

var _os_unfair_lock_lock func(lock Os_unfair_lock_t)
var _os_unfair_lock_lockErr error

func tryOs_unfair_lock_lock(lock Os_unfair_lock_t) error {
	if _os_unfair_lock_lock == nil {
		return symbolCallError("os_unfair_lock_lock", "10.12", _os_unfair_lock_lockErr)
	}
	_os_unfair_lock_lock(lock)
	return nil
}

// Os_unfair_lock_lock a low-level lock that allows waiters to block efficiently on contention.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_lock
func Os_unfair_lock_lock(lock Os_unfair_lock_t) {
	if callErr := tryOs_unfair_lock_lock(lock); callErr != nil {
		panic(callErr)
	}
}

var _os_unfair_lock_lock_with_flags func(lock Os_unfair_lock_t, flags OsUnfairLockFlag)
var _os_unfair_lock_lock_with_flagsErr error

func tryOs_unfair_lock_lock_with_flags(lock Os_unfair_lock_t, flags OsUnfairLockFlag) error {
	if _os_unfair_lock_lock_with_flags == nil {
		return symbolCallError("os_unfair_lock_lock_with_flags", "15.0", _os_unfair_lock_lock_with_flagsErr)
	}
	_os_unfair_lock_lock_with_flags(lock, flags)
	return nil
}

// Os_unfair_lock_lock_with_flags.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_lock_with_flags
func Os_unfair_lock_lock_with_flags(lock Os_unfair_lock_t, flags OsUnfairLockFlag) {
	if callErr := tryOs_unfair_lock_lock_with_flags(lock, flags); callErr != nil {
		panic(callErr)
	}
}

var _os_unfair_lock_trylock func(lock Os_unfair_lock_t) bool
var _os_unfair_lock_trylockErr error

func tryOs_unfair_lock_trylock(lock Os_unfair_lock_t) (bool, error) {
	if _os_unfair_lock_trylock == nil {
		return false, symbolCallError("os_unfair_lock_trylock", "10.12", _os_unfair_lock_trylockErr)
	}
	return _os_unfair_lock_trylock(lock), nil
}

// Os_unfair_lock_trylock locks an unfair lock if it is not already locked.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_trylock
func Os_unfair_lock_trylock(lock Os_unfair_lock_t) bool {
	result, callErr := tryOs_unfair_lock_trylock(lock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_unfair_lock_unlock func(lock Os_unfair_lock_t)
var _os_unfair_lock_unlockErr error

func tryOs_unfair_lock_unlock(lock Os_unfair_lock_t) error {
	if _os_unfair_lock_unlock == nil {
		return symbolCallError("os_unfair_lock_unlock", "10.12", _os_unfair_lock_unlockErr)
	}
	_os_unfair_lock_unlock(lock)
	return nil
}

// Os_unfair_lock_unlock unlocks an unfair lock.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_unlock
func Os_unfair_lock_unlock(lock Os_unfair_lock_t) {
	if callErr := tryOs_unfair_lock_unlock(lock); callErr != nil {
		panic(callErr)
	}
}

var _os_workgroup_cancel func(wg Os_workgroup_t)
var _os_workgroup_cancelErr error

func tryOs_workgroup_cancel(wg Os_workgroup_t) error {
	if _os_workgroup_cancel == nil {
		return symbolCallError("os_workgroup_cancel", "11.0", _os_workgroup_cancelErr)
	}
	_os_workgroup_cancel(wg)
	return nil
}

// Os_workgroup_cancel cancels and invalidates the specified workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_cancel
func Os_workgroup_cancel(wg Os_workgroup_t) {
	if callErr := tryOs_workgroup_cancel(wg); callErr != nil {
		panic(callErr)
	}
}

var _os_workgroup_copy_port func(wg Os_workgroup_t, mach_port_out *uint32) int
var _os_workgroup_copy_portErr error

func tryOs_workgroup_copy_port(wg Os_workgroup_t, mach_port_out *uint32) (int, error) {
	if _os_workgroup_copy_port == nil {
		return 0, symbolCallError("os_workgroup_copy_port", "11.0", _os_workgroup_copy_portErr)
	}
	return _os_workgroup_copy_port(wg, mach_port_out), nil
}

// Os_workgroup_copy_port returns the Mach port associated with the workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_copy_port
func Os_workgroup_copy_port(wg Os_workgroup_t, mach_port_out *uint32) int {
	result, callErr := tryOs_workgroup_copy_port(wg, mach_port_out)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_create_with_port func(name string, mach_port uint32) Os_workgroup_t
var _os_workgroup_create_with_portErr error

func tryOs_workgroup_create_with_port(name string, mach_port uint32) (Os_workgroup_t, error) {
	if _os_workgroup_create_with_port == nil {
		return *new(Os_workgroup_t), symbolCallError("os_workgroup_create_with_port", "11.0", _os_workgroup_create_with_portErr)
	}
	return _os_workgroup_create_with_port(name, mach_port), nil
}

// Os_workgroup_create_with_port creates a new workgroup that is bound to the workgroup with the specified Mach port.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_create_with_port
func Os_workgroup_create_with_port(name string, mach_port uint32) Os_workgroup_t {
	result, callErr := tryOs_workgroup_create_with_port(name, mach_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_create_with_workgroup func(name string, wg Os_workgroup_t) Os_workgroup_t
var _os_workgroup_create_with_workgroupErr error

func tryOs_workgroup_create_with_workgroup(name string, wg Os_workgroup_t) (Os_workgroup_t, error) {
	if _os_workgroup_create_with_workgroup == nil {
		return *new(Os_workgroup_t), symbolCallError("os_workgroup_create_with_workgroup", "11.0", _os_workgroup_create_with_workgroupErr)
	}
	return _os_workgroup_create_with_workgroup(name, wg), nil
}

// Os_workgroup_create_with_workgroup create a new workgroup that is bound to the specified workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_create_with_workgroup
func Os_workgroup_create_with_workgroup(name string, wg Os_workgroup_t) Os_workgroup_t {
	result, callErr := tryOs_workgroup_create_with_workgroup(name, wg)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_get_working_arena func(wg Os_workgroup_t, index_out *Os_workgroup_index) unsafe.Pointer
var _os_workgroup_get_working_arenaErr error

func tryOs_workgroup_get_working_arena(wg Os_workgroup_t, index_out *Os_workgroup_index) (unsafe.Pointer, error) {
	if _os_workgroup_get_working_arena == nil {
		return nil, symbolCallError("os_workgroup_get_working_arena", "11.0", _os_workgroup_get_working_arenaErr)
	}
	return _os_workgroup_get_working_arena(wg, index_out), nil
}

// Os_workgroup_get_working_arena retrieves the workgroup’s shared data, and the thread-specific index into that data.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_get_working_arena
func Os_workgroup_get_working_arena(wg Os_workgroup_t, index_out *Os_workgroup_index) unsafe.Pointer {
	result, callErr := tryOs_workgroup_get_working_arena(wg, index_out)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_interval_finish func(wg Os_workgroup_interval_t, data Os_workgroup_interval_data_t) int
var _os_workgroup_interval_finishErr error

func tryOs_workgroup_interval_finish(wg Os_workgroup_interval_t, data Os_workgroup_interval_data_t) (int, error) {
	if _os_workgroup_interval_finish == nil {
		return 0, symbolCallError("os_workgroup_interval_finish", "11.0", _os_workgroup_interval_finishErr)
	}
	return _os_workgroup_interval_finish(wg, data), nil
}

// Os_workgroup_interval_finish stops the current interval-based execution of the workgroup’s threads.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_interval_finish
func Os_workgroup_interval_finish(wg Os_workgroup_interval_t, data Os_workgroup_interval_data_t) int {
	result, callErr := tryOs_workgroup_interval_finish(wg, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_interval_start func(wg Os_workgroup_interval_t, start uint64, deadline uint64, data Os_workgroup_interval_data_t) int
var _os_workgroup_interval_startErr error

func tryOs_workgroup_interval_start(wg Os_workgroup_interval_t, start uint64, deadline uint64, data Os_workgroup_interval_data_t) (int, error) {
	if _os_workgroup_interval_start == nil {
		return 0, symbolCallError("os_workgroup_interval_start", "11.0", _os_workgroup_interval_startErr)
	}
	return _os_workgroup_interval_start(wg, start, deadline, data), nil
}

// Os_workgroup_interval_start starts the regular execution of the workgroup’s threads at the specified time.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_interval_start
func Os_workgroup_interval_start(wg Os_workgroup_interval_t, start uint64, deadline uint64, data Os_workgroup_interval_data_t) int {
	result, callErr := tryOs_workgroup_interval_start(wg, start, deadline, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_interval_update func(wg Os_workgroup_interval_t, deadline uint64, data Os_workgroup_interval_data_t) int
var _os_workgroup_interval_updateErr error

func tryOs_workgroup_interval_update(wg Os_workgroup_interval_t, deadline uint64, data Os_workgroup_interval_data_t) (int, error) {
	if _os_workgroup_interval_update == nil {
		return 0, symbolCallError("os_workgroup_interval_update", "11.0", _os_workgroup_interval_updateErr)
	}
	return _os_workgroup_interval_update(wg, deadline, data), nil
}

// Os_workgroup_interval_update schedules a new deadline for workgroup threads that run at regular intervals.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_interval_update
func Os_workgroup_interval_update(wg Os_workgroup_interval_t, deadline uint64, data Os_workgroup_interval_data_t) int {
	result, callErr := tryOs_workgroup_interval_update(wg, deadline, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_join func(wg Os_workgroup_t, token_out Os_workgroup_join_token_t) int
var _os_workgroup_joinErr error

func tryOs_workgroup_join(wg Os_workgroup_t, token_out Os_workgroup_join_token_t) (int, error) {
	if _os_workgroup_join == nil {
		return 0, symbolCallError("os_workgroup_join", "11.0", _os_workgroup_joinErr)
	}
	return _os_workgroup_join(wg, token_out), nil
}

// Os_workgroup_join adds the current thread to the specified workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_join
func Os_workgroup_join(wg Os_workgroup_t, token_out Os_workgroup_join_token_t) int {
	result, callErr := tryOs_workgroup_join(wg, token_out)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_leave func(wg Os_workgroup_t, token Os_workgroup_join_token_t)
var _os_workgroup_leaveErr error

func tryOs_workgroup_leave(wg Os_workgroup_t, token Os_workgroup_join_token_t) error {
	if _os_workgroup_leave == nil {
		return symbolCallError("os_workgroup_leave", "11.0", _os_workgroup_leaveErr)
	}
	_os_workgroup_leave(wg, token)
	return nil
}

// Os_workgroup_leave removes the current thread from the workgroup it previously joined.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_leave
func Os_workgroup_leave(wg Os_workgroup_t, token Os_workgroup_join_token_t) {
	if callErr := tryOs_workgroup_leave(wg, token); callErr != nil {
		panic(callErr)
	}
}

var _os_workgroup_max_parallel_threads func(wg Os_workgroup_t, attr Os_workgroup_mpt_attr_t) int
var _os_workgroup_max_parallel_threadsErr error

func tryOs_workgroup_max_parallel_threads(wg Os_workgroup_t, attr Os_workgroup_mpt_attr_t) (int, error) {
	if _os_workgroup_max_parallel_threads == nil {
		return 0, symbolCallError("os_workgroup_max_parallel_threads", "11.0", _os_workgroup_max_parallel_threadsErr)
	}
	return _os_workgroup_max_parallel_threads(wg, attr), nil
}

// Os_workgroup_max_parallel_threads returns the maximum number of threads that the system recommends you add to the specified workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_max_parallel_threads
func Os_workgroup_max_parallel_threads(wg Os_workgroup_t, attr Os_workgroup_mpt_attr_t) int {
	result, callErr := tryOs_workgroup_max_parallel_threads(wg, attr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_parallel_create func(name string, attr Os_workgroup_attr_t) Os_workgroup_parallel_t
var _os_workgroup_parallel_createErr error

func tryOs_workgroup_parallel_create(name string, attr Os_workgroup_attr_t) (Os_workgroup_parallel_t, error) {
	if _os_workgroup_parallel_create == nil {
		return *new(Os_workgroup_parallel_t), symbolCallError("os_workgroup_parallel_create", "11.0", _os_workgroup_parallel_createErr)
	}
	return _os_workgroup_parallel_create(name, attr), nil
}

// Os_workgroup_parallel_create creates a new workgroup that manges threads working on a single task in parallel.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_parallel_create
func Os_workgroup_parallel_create(name string, attr Os_workgroup_attr_t) Os_workgroup_parallel_t {
	result, callErr := tryOs_workgroup_parallel_create(name, attr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_set_working_arena func(wg Os_workgroup_t, arena unsafe.Pointer, max_workers uint32, destructor Os_workgroup_working_arena_destructor_t) int
var _os_workgroup_set_working_arenaErr error

func tryOs_workgroup_set_working_arena(wg Os_workgroup_t, arena unsafe.Pointer, max_workers uint32, destructor Os_workgroup_working_arena_destructor_t) (int, error) {
	if _os_workgroup_set_working_arena == nil {
		return 0, symbolCallError("os_workgroup_set_working_arena", "11.0", _os_workgroup_set_working_arenaErr)
	}
	return _os_workgroup_set_working_arena(wg, arena, max_workers, destructor), nil
}

// Os_workgroup_set_working_arena distributes a block of managed memory to the threads of a workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_set_working_arena
func Os_workgroup_set_working_arena(wg Os_workgroup_t, arena unsafe.Pointer, max_workers uint32, destructor Os_workgroup_working_arena_destructor_t) int {
	result, callErr := tryOs_workgroup_set_working_arena(wg, arena, max_workers, destructor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_workgroup_testcancel func(wg Os_workgroup_t) bool
var _os_workgroup_testcancelErr error

func tryOs_workgroup_testcancel(wg Os_workgroup_t) (bool, error) {
	if _os_workgroup_testcancel == nil {
		return false, symbolCallError("os_workgroup_testcancel", "11.0", _os_workgroup_testcancelErr)
	}
	return _os_workgroup_testcancel(wg), nil
}

// Os_workgroup_testcancel returns a Boolean value that indicates whether the workgroup is canceled.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_testcancel
func Os_workgroup_testcancel(wg Os_workgroup_t) bool {
	result, callErr := tryOs_workgroup_testcancel(wg)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_os_activity_apply, &_os_activity_applyErr, frameworkHandle, "os_activity_apply", "10.12")
	registerFunc(&_os_activity_apply_f, &_os_activity_apply_fErr, frameworkHandle, "os_activity_apply_f", "10.12")
	registerFunc(&_os_activity_get_identifier, &_os_activity_get_identifierErr, frameworkHandle, "os_activity_get_identifier", "10.12")
	registerFunc(&_os_activity_scope_enter, &_os_activity_scope_enterErr, frameworkHandle, "os_activity_scope_enter", "10.12")
	registerFunc(&_os_activity_scope_leave, &_os_activity_scope_leaveErr, frameworkHandle, "os_activity_scope_leave", "10.12")
	registerFunc(&_os_log_create, &_os_log_createErr, frameworkHandle, "os_log_create", "10.12")
	registerFunc(&_os_log_type_enabled, &_os_log_type_enabledErr, frameworkHandle, "os_log_type_enabled", "10.12")
	registerFunc(&_os_release, &_os_releaseErr, frameworkHandle, "os_release", "10.10")
	registerFunc(&_os_retain, &_os_retainErr, frameworkHandle, "os_retain", "10.10")
	registerFunc(&_os_security_config_get, &_os_security_config_getErr, frameworkHandle, "os_security_config_get", "26.0")
	registerFunc(&_os_security_config_get_for_proc, &_os_security_config_get_for_procErr, frameworkHandle, "os_security_config_get_for_proc", "26.0")
	registerFunc(&_os_security_config_get_for_task, &_os_security_config_get_for_taskErr, frameworkHandle, "os_security_config_get_for_task", "26.0")
	registerFunc(&_os_signpost_enabled, &_os_signpost_enabledErr, frameworkHandle, "os_signpost_enabled", "10.14")
	registerFunc(&_os_signpost_id_generate, &_os_signpost_id_generateErr, frameworkHandle, "os_signpost_id_generate", "10.14")
	registerFunc(&_os_signpost_id_make_with_pointer, &_os_signpost_id_make_with_pointerErr, frameworkHandle, "os_signpost_id_make_with_pointer", "10.14")
	registerFunc(&_os_sync_wait_on_address, &_os_sync_wait_on_addressErr, frameworkHandle, "os_sync_wait_on_address", "14.4")
	registerFunc(&_os_sync_wait_on_address_with_deadline, &_os_sync_wait_on_address_with_deadlineErr, frameworkHandle, "os_sync_wait_on_address_with_deadline", "14.4")
	registerFunc(&_os_sync_wait_on_address_with_timeout, &_os_sync_wait_on_address_with_timeoutErr, frameworkHandle, "os_sync_wait_on_address_with_timeout", "14.4")
	registerFunc(&_os_sync_wake_by_address_all, &_os_sync_wake_by_address_allErr, frameworkHandle, "os_sync_wake_by_address_all", "14.4")
	registerFunc(&_os_sync_wake_by_address_any, &_os_sync_wake_by_address_anyErr, frameworkHandle, "os_sync_wake_by_address_any", "14.4")
	registerFunc(&_os_unfair_lock_assert_not_owner, &_os_unfair_lock_assert_not_ownerErr, frameworkHandle, "os_unfair_lock_assert_not_owner", "10.12")
	registerFunc(&_os_unfair_lock_assert_owner, &_os_unfair_lock_assert_ownerErr, frameworkHandle, "os_unfair_lock_assert_owner", "10.12")
	registerFunc(&_os_unfair_lock_lock, &_os_unfair_lock_lockErr, frameworkHandle, "os_unfair_lock_lock", "10.12")
	registerFunc(&_os_unfair_lock_lock_with_flags, &_os_unfair_lock_lock_with_flagsErr, frameworkHandle, "os_unfair_lock_lock_with_flags", "15.0")
	registerFunc(&_os_unfair_lock_trylock, &_os_unfair_lock_trylockErr, frameworkHandle, "os_unfair_lock_trylock", "10.12")
	registerFunc(&_os_unfair_lock_unlock, &_os_unfair_lock_unlockErr, frameworkHandle, "os_unfair_lock_unlock", "10.12")
	registerFunc(&_os_workgroup_cancel, &_os_workgroup_cancelErr, frameworkHandle, "os_workgroup_cancel", "11.0")
	registerFunc(&_os_workgroup_copy_port, &_os_workgroup_copy_portErr, frameworkHandle, "os_workgroup_copy_port", "11.0")
	registerFunc(&_os_workgroup_create_with_port, &_os_workgroup_create_with_portErr, frameworkHandle, "os_workgroup_create_with_port", "11.0")
	registerFunc(&_os_workgroup_create_with_workgroup, &_os_workgroup_create_with_workgroupErr, frameworkHandle, "os_workgroup_create_with_workgroup", "11.0")
	registerFunc(&_os_workgroup_get_working_arena, &_os_workgroup_get_working_arenaErr, frameworkHandle, "os_workgroup_get_working_arena", "11.0")
	registerFunc(&_os_workgroup_interval_finish, &_os_workgroup_interval_finishErr, frameworkHandle, "os_workgroup_interval_finish", "11.0")
	registerFunc(&_os_workgroup_interval_start, &_os_workgroup_interval_startErr, frameworkHandle, "os_workgroup_interval_start", "11.0")
	registerFunc(&_os_workgroup_interval_update, &_os_workgroup_interval_updateErr, frameworkHandle, "os_workgroup_interval_update", "11.0")
	registerFunc(&_os_workgroup_join, &_os_workgroup_joinErr, frameworkHandle, "os_workgroup_join", "11.0")
	registerFunc(&_os_workgroup_leave, &_os_workgroup_leaveErr, frameworkHandle, "os_workgroup_leave", "11.0")
	registerFunc(&_os_workgroup_max_parallel_threads, &_os_workgroup_max_parallel_threadsErr, frameworkHandle, "os_workgroup_max_parallel_threads", "11.0")
	registerFunc(&_os_workgroup_parallel_create, &_os_workgroup_parallel_createErr, frameworkHandle, "os_workgroup_parallel_create", "11.0")
	registerFunc(&_os_workgroup_set_working_arena, &_os_workgroup_set_working_arenaErr, frameworkHandle, "os_workgroup_set_working_arena", "11.0")
	registerFunc(&_os_workgroup_testcancel, &_os_workgroup_testcancelErr, frameworkHandle, "os_workgroup_testcancel", "11.0")
}

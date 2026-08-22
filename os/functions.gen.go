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

var _oSActivityApply func(activity OSActivity, block unsafe.Pointer)
var _oSActivityApplyErr error

func tryOSActivityApply(activity OSActivity, block OSBlock) error {
	if _oSActivityApply == nil {
		return symbolCallError("os_activity_apply", "10.12", _oSActivityApplyErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block) { block() })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_oSActivityApply(activity, _block0)
	return nil
}

// OSActivityApply execute a block using a given activity object.
//
// See: https://developer.apple.com/documentation/os/os_activity_apply
func OSActivityApply(activity OSActivity, block OSBlock) {
	if callErr := tryOSActivityApply(activity, block); callErr != nil {
		panic(callErr)
	}
}

var _oSActivityApplyF func(activity OSActivity, context unsafe.Pointer, function OSFunction)
var _oSActivityApplyFErr error

func tryOSActivityApplyF(activity OSActivity, context unsafe.Pointer, function OSFunction) error {
	if _oSActivityApplyF == nil {
		return symbolCallError("os_activity_apply_f", "10.12", _oSActivityApplyFErr)
	}
	_oSActivityApplyF(activity, context, function)
	return nil
}

// OSActivityApplyF execute a function using a given activity object.
//
// See: https://developer.apple.com/documentation/os/os_activity_apply_f
func OSActivityApplyF(activity OSActivity, context unsafe.Pointer, function OSFunction) {
	if callErr := tryOSActivityApplyF(activity, context, function); callErr != nil {
		panic(callErr)
	}
}

var _oSActivityGetIdentifier func(activity OSActivity, parent_id *OSActivityID) OSActivityID
var _oSActivityGetIdentifierErr error

func tryOSActivityGetIdentifier(activity OSActivity, parent_id *OSActivityID) (OSActivityID, error) {
	if _oSActivityGetIdentifier == nil {
		return *new(OSActivityID), symbolCallError("os_activity_get_identifier", "10.12", _oSActivityGetIdentifierErr)
	}
	return _oSActivityGetIdentifier(activity, parent_id), nil
}

// OSActivityGetIdentifier retrieves the identifier for a given activity object.
//
// See: https://developer.apple.com/documentation/os/os_activity_get_identifier
func OSActivityGetIdentifier(activity OSActivity, parent_id *OSActivityID) OSActivityID {
	result, callErr := tryOSActivityGetIdentifier(activity, parent_id)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSActivityScopeEnter func(activity OSActivity, state OSActivityScopeState)
var _oSActivityScopeEnterErr error

func tryOSActivityScopeEnter(activity OSActivity, state OSActivityScopeState) error {
	if _oSActivityScopeEnter == nil {
		return symbolCallError("os_activity_scope_enter", "10.12", _oSActivityScopeEnterErr)
	}
	_oSActivityScopeEnter(activity, state)
	return nil
}

// OSActivityScopeEnter switches the current activity, saving the existing execution context.
//
// See: https://developer.apple.com/documentation/os/os_activity_scope_enter
func OSActivityScopeEnter(activity OSActivity, state OSActivityScopeState) {
	if callErr := tryOSActivityScopeEnter(activity, state); callErr != nil {
		panic(callErr)
	}
}

var _oSActivityScopeLeave func(state OSActivityScopeState)
var _oSActivityScopeLeaveErr error

func tryOSActivityScopeLeave(state OSActivityScopeState) error {
	if _oSActivityScopeLeave == nil {
		return symbolCallError("os_activity_scope_leave", "10.12", _oSActivityScopeLeaveErr)
	}
	_oSActivityScopeLeave(state)
	return nil
}

// OSActivityScopeLeave restores the current activity to a previously saved state.
//
// See: https://developer.apple.com/documentation/os/os_activity_scope_leave
func OSActivityScopeLeave(state OSActivityScopeState) {
	if callErr := tryOSActivityScopeLeave(state); callErr != nil {
		panic(callErr)
	}
}

var _oSLogCreate func(subsystem string, category string) kernel.Os_log_t
var _oSLogCreateErr error

func tryOSLogCreate(subsystem string, category string) (kernel.Os_log_t, error) {
	if _oSLogCreate == nil {
		return *new(kernel.Os_log_t), symbolCallError("os_log_create", "10.12", _oSLogCreateErr)
	}
	return _oSLogCreate(subsystem, category), nil
}

// OSLogCreate creates a custom log object.
//
// See: https://developer.apple.com/documentation/os/os_log_create
func OSLogCreate(subsystem string, category string) kernel.Os_log_t {
	result, callErr := tryOSLogCreate(subsystem, category)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSLogTypeEnabled func(oslog kernel.Os_log_t, type_ OSLogType) bool
var _oSLogTypeEnabledErr error

func tryOSLogTypeEnabled(oslog kernel.Os_log_t, type_ OSLogType) (bool, error) {
	if _oSLogTypeEnabled == nil {
		return false, symbolCallError("os_log_type_enabled", "10.12", _oSLogTypeEnabledErr)
	}
	return _oSLogTypeEnabled(oslog, type_), nil
}

// OSLogTypeEnabled returns a Boolean value that indicates whether the log can write messages with the specified log type.
//
// See: https://developer.apple.com/documentation/os/OSLog/isEnabled(type:)
func OSLogTypeEnabled(oslog kernel.Os_log_t, type_ OSLogType) bool {
	result, callErr := tryOSLogTypeEnabled(oslog, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSRelease func(object unsafe.Pointer)
var _oSReleaseErr error

func tryOSRelease(object unsafe.Pointer) error {
	if _oSRelease == nil {
		return symbolCallError("os_release", "10.10", _oSReleaseErr)
	}
	_oSRelease(object)
	return nil
}

// OSRelease.
//
// See: https://developer.apple.com/documentation/os/os_release-c.func
func OSRelease(object unsafe.Pointer) {
	if callErr := tryOSRelease(object); callErr != nil {
		panic(callErr)
	}
}

var _oSRetain func(object unsafe.Pointer) unsafe.Pointer
var _oSRetainErr error

func tryOSRetain(object unsafe.Pointer) (unsafe.Pointer, error) {
	if _oSRetain == nil {
		return nil, symbolCallError("os_retain", "10.10", _oSRetainErr)
	}
	return _oSRetain(object), nil
}

// OSRetain.
//
// See: https://developer.apple.com/documentation/os/os_retain-c.func
func OSRetain(object unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryOSRetain(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSecurityConfigGet func() OSSecurityConfig
var _oSSecurityConfigGetErr error

func tryOSSecurityConfigGet() (OSSecurityConfig, error) {
	if _oSSecurityConfigGet == nil {
		return *new(OSSecurityConfig), symbolCallError("os_security_config_get", "26.0", _oSSecurityConfigGetErr)
	}
	return _oSSecurityConfigGet(), nil
}

// OSSecurityConfigGet.
//
// See: https://developer.apple.com/documentation/os/os_security_config_get
func OSSecurityConfigGet() OSSecurityConfig {
	result, callErr := tryOSSecurityConfigGet()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSecurityConfigGetForProc func(pid int32, config *OSSecurityConfig) int32
var _oSSecurityConfigGetForProcErr error

func tryOSSecurityConfigGetForProc(pid int32, config *OSSecurityConfig) (int32, error) {
	if _oSSecurityConfigGetForProc == nil {
		return 0, symbolCallError("os_security_config_get_for_proc", "26.0", _oSSecurityConfigGetForProcErr)
	}
	return _oSSecurityConfigGetForProc(pid, config), nil
}

// OSSecurityConfigGetForProc.
//
// See: https://developer.apple.com/documentation/os/os_security_config_get_for_proc
func OSSecurityConfigGetForProc(pid int32, config *OSSecurityConfig) int32 {
	result, callErr := tryOSSecurityConfigGetForProc(pid, config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSecurityConfigGetForTask func(task uint32, config *OSSecurityConfig) int32
var _oSSecurityConfigGetForTaskErr error

func tryOSSecurityConfigGetForTask(task uint32, config *OSSecurityConfig) (int32, error) {
	if _oSSecurityConfigGetForTask == nil {
		return 0, symbolCallError("os_security_config_get_for_task", "26.0", _oSSecurityConfigGetForTaskErr)
	}
	return _oSSecurityConfigGetForTask(task, config), nil
}

// OSSecurityConfigGetForTask.
//
// See: https://developer.apple.com/documentation/os/os_security_config_get_for_task
func OSSecurityConfigGetForTask(task uint32, config *OSSecurityConfig) int32 {
	result, callErr := tryOSSecurityConfigGetForTask(task, config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSignpostEnabled func(log kernel.Os_log_t) bool
var _oSSignpostEnabledErr error

func tryOSSignpostEnabled(log kernel.Os_log_t) (bool, error) {
	if _oSSignpostEnabled == nil {
		return false, symbolCallError("os_signpost_enabled", "10.14", _oSSignpostEnabledErr)
	}
	return _oSSignpostEnabled(log), nil
}

// OSSignpostEnabled returns a Boolean value that indicates whether signposts are in an enabled state for the specified log.
//
// See: https://developer.apple.com/documentation/os/os_signpost_enabled
func OSSignpostEnabled(log kernel.Os_log_t) bool {
	result, callErr := tryOSSignpostEnabled(log)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSignpostIDGenerate func(log kernel.Os_log_t) OSSignpostID
var _oSSignpostIDGenerateErr error

func tryOSSignpostIDGenerate(log kernel.Os_log_t) (OSSignpostID, error) {
	if _oSSignpostIDGenerate == nil {
		return *new(OSSignpostID), symbolCallError("os_signpost_id_generate", "10.14", _oSSignpostIDGenerateErr)
	}
	return _oSSignpostIDGenerate(log), nil
}

// OSSignpostIDGenerate creates a signpost identifier that’s unique among signposts logged to a specified log.
//
// See: https://developer.apple.com/documentation/os/os_signpost_id_generate
func OSSignpostIDGenerate(log kernel.Os_log_t) OSSignpostID {
	result, callErr := tryOSSignpostIDGenerate(log)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSignpostIDMakeWithPointer func(log kernel.Os_log_t, ptr unsafe.Pointer) OSSignpostID
var _oSSignpostIDMakeWithPointerErr error

func tryOSSignpostIDMakeWithPointer(log kernel.Os_log_t, ptr unsafe.Pointer) (OSSignpostID, error) {
	if _oSSignpostIDMakeWithPointer == nil {
		return *new(OSSignpostID), symbolCallError("os_signpost_id_make_with_pointer", "10.14", _oSSignpostIDMakeWithPointerErr)
	}
	return _oSSignpostIDMakeWithPointer(log, ptr), nil
}

// OSSignpostIDMakeWithPointer creates a signpost identifier that’s unique among signposts logging to the specified log, using a pointer value to generate the unique value.
//
// See: https://developer.apple.com/documentation/os/os_signpost_id_make_with_pointer
func OSSignpostIDMakeWithPointer(log kernel.Os_log_t, ptr unsafe.Pointer) OSSignpostID {
	result, callErr := tryOSSignpostIDMakeWithPointer(log, ptr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSyncWaitOnAddress func(addr unsafe.Pointer, value uint64, size uintptr, flags OSSyncWaitOnAddressFlags) int32
var _oSSyncWaitOnAddressErr error

func tryOSSyncWaitOnAddress(addr unsafe.Pointer, value uint64, size uintptr, flags OSSyncWaitOnAddressFlags) (int32, error) {
	if _oSSyncWaitOnAddress == nil {
		return 0, symbolCallError("os_sync_wait_on_address", "14.4", _oSSyncWaitOnAddressErr)
	}
	return _oSSyncWaitOnAddress(addr, value, size, flags), nil
}

// OSSyncWaitOnAddress an atomic compare-and-wait operation, used to implement higher-level synchronization primitives.
//
// See: https://developer.apple.com/documentation/os/os_sync_wait_on_address
func OSSyncWaitOnAddress(addr unsafe.Pointer, value uint64, size uintptr, flags OSSyncWaitOnAddressFlags) int32 {
	result, callErr := tryOSSyncWaitOnAddress(addr, value, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSyncWaitOnAddressWithDeadline func(addr unsafe.Pointer, value uint64, size uintptr, flags OSSyncWaitOnAddressFlags, clockid OSClockid, deadline uint64) int32
var _oSSyncWaitOnAddressWithDeadlineErr error

func tryOSSyncWaitOnAddressWithDeadline(addr unsafe.Pointer, value uint64, size uintptr, flags OSSyncWaitOnAddressFlags, clockid OSClockid, deadline uint64) (int32, error) {
	if _oSSyncWaitOnAddressWithDeadline == nil {
		return 0, symbolCallError("os_sync_wait_on_address_with_deadline", "14.4", _oSSyncWaitOnAddressWithDeadlineErr)
	}
	return _oSSyncWaitOnAddressWithDeadline(addr, value, size, flags, clockid, deadline), nil
}

// OSSyncWaitOnAddressWithDeadline an atomic compare-and-wait operation with a deadline, used to implement higher-level synchronization primitives.
//
// See: https://developer.apple.com/documentation/os/os_sync_wait_on_address_with_deadline
func OSSyncWaitOnAddressWithDeadline(addr unsafe.Pointer, value uint64, size uintptr, flags OSSyncWaitOnAddressFlags, clockid OSClockid, deadline uint64) int32 {
	result, callErr := tryOSSyncWaitOnAddressWithDeadline(addr, value, size, flags, clockid, deadline)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSyncWaitOnAddressWithTimeout func(addr unsafe.Pointer, value uint64, size uintptr, flags OSSyncWaitOnAddressFlags, clockid OSClockid, timeout_ns uint64) int32
var _oSSyncWaitOnAddressWithTimeoutErr error

func tryOSSyncWaitOnAddressWithTimeout(addr unsafe.Pointer, value uint64, size uintptr, flags OSSyncWaitOnAddressFlags, clockid OSClockid, timeout_ns uint64) (int32, error) {
	if _oSSyncWaitOnAddressWithTimeout == nil {
		return 0, symbolCallError("os_sync_wait_on_address_with_timeout", "14.4", _oSSyncWaitOnAddressWithTimeoutErr)
	}
	return _oSSyncWaitOnAddressWithTimeout(addr, value, size, flags, clockid, timeout_ns), nil
}

// OSSyncWaitOnAddressWithTimeout an atomic compare-and-wait operation with a timeout, used to implement higher-level synchronization primitives.
//
// See: https://developer.apple.com/documentation/os/os_sync_wait_on_address_with_timeout
func OSSyncWaitOnAddressWithTimeout(addr unsafe.Pointer, value uint64, size uintptr, flags OSSyncWaitOnAddressFlags, clockid OSClockid, timeout_ns uint64) int32 {
	result, callErr := tryOSSyncWaitOnAddressWithTimeout(addr, value, size, flags, clockid, timeout_ns)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSyncWakeByAddressAll func(addr unsafe.Pointer, size uintptr, flags OSSyncWakeByAddressFlags) int32
var _oSSyncWakeByAddressAllErr error

func tryOSSyncWakeByAddressAll(addr unsafe.Pointer, size uintptr, flags OSSyncWakeByAddressFlags) (int32, error) {
	if _oSSyncWakeByAddressAll == nil {
		return 0, symbolCallError("os_sync_wake_by_address_all", "14.4", _oSSyncWakeByAddressAllErr)
	}
	return _oSSyncWakeByAddressAll(addr, size, flags), nil
}

// OSSyncWakeByAddressAll an atomic operation that wakes all threads blocked on a futex wait, used to implement higher-level synchronization primitives.
//
// See: https://developer.apple.com/documentation/os/os_sync_wake_by_address_all
func OSSyncWakeByAddressAll(addr unsafe.Pointer, size uintptr, flags OSSyncWakeByAddressFlags) int32 {
	result, callErr := tryOSSyncWakeByAddressAll(addr, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSSyncWakeByAddressAny func(addr unsafe.Pointer, size uintptr, flags OSSyncWakeByAddressFlags) int32
var _oSSyncWakeByAddressAnyErr error

func tryOSSyncWakeByAddressAny(addr unsafe.Pointer, size uintptr, flags OSSyncWakeByAddressFlags) (int32, error) {
	if _oSSyncWakeByAddressAny == nil {
		return 0, symbolCallError("os_sync_wake_by_address_any", "14.4", _oSSyncWakeByAddressAnyErr)
	}
	return _oSSyncWakeByAddressAny(addr, size, flags), nil
}

// OSSyncWakeByAddressAny an atomic operation that wakes one thread blocked on a futex wait, used to implement higher-level synchronization primitives.
//
// See: https://developer.apple.com/documentation/os/os_sync_wake_by_address_any
func OSSyncWakeByAddressAny(addr unsafe.Pointer, size uintptr, flags OSSyncWakeByAddressFlags) int32 {
	result, callErr := tryOSSyncWakeByAddressAny(addr, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSUnfairLockAssertNotOwner func(lock *[4]byte)
var _oSUnfairLockAssertNotOwnerErr error

func tryOSUnfairLockAssertNotOwner(lock *[4]byte) error {
	if _oSUnfairLockAssertNotOwner == nil {
		return symbolCallError("os_unfair_lock_assert_not_owner", "10.12", _oSUnfairLockAssertNotOwnerErr)
	}
	_oSUnfairLockAssertNotOwner(lock)
	return nil
}

// OSUnfairLockAssertNotOwner triggers an assertion if the calling thread owns the specified unfair lock.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_assert_not_owner
func OSUnfairLockAssertNotOwner(lock *[4]byte) {
	if callErr := tryOSUnfairLockAssertNotOwner(lock); callErr != nil {
		panic(callErr)
	}
}

var _oSUnfairLockAssertOwner func(lock *[4]byte)
var _oSUnfairLockAssertOwnerErr error

func tryOSUnfairLockAssertOwner(lock *[4]byte) error {
	if _oSUnfairLockAssertOwner == nil {
		return symbolCallError("os_unfair_lock_assert_owner", "10.12", _oSUnfairLockAssertOwnerErr)
	}
	_oSUnfairLockAssertOwner(lock)
	return nil
}

// OSUnfairLockAssertOwner triggers an assertion if the calling thread doesn’t own the specified unfair lock.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_assert_owner
func OSUnfairLockAssertOwner(lock *[4]byte) {
	if callErr := tryOSUnfairLockAssertOwner(lock); callErr != nil {
		panic(callErr)
	}
}

var _oSUnfairLockLock func(lock unsafe.Pointer)
var _oSUnfairLockLockErr error

func tryOSUnfairLockLock(lock unsafe.Pointer) error {
	if _oSUnfairLockLock == nil {
		return symbolCallError("os_unfair_lock_lock", "10.12", _oSUnfairLockLockErr)
	}
	_oSUnfairLockLock(lock)
	return nil
}

// OSUnfairLockLock a low-level lock that allows waiters to block efficiently on contention.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_lock
func OSUnfairLockLock(lock unsafe.Pointer) {
	if callErr := tryOSUnfairLockLock(lock); callErr != nil {
		panic(callErr)
	}
}

var _oSUnfairLockLockWithFlags func(lock unsafe.Pointer, flags OSUnfairLockFlags)
var _oSUnfairLockLockWithFlagsErr error

func tryOSUnfairLockLockWithFlags(lock unsafe.Pointer, flags OSUnfairLockFlags) error {
	if _oSUnfairLockLockWithFlags == nil {
		return symbolCallError("os_unfair_lock_lock_with_flags", "15.0", _oSUnfairLockLockWithFlagsErr)
	}
	_oSUnfairLockLockWithFlags(lock, flags)
	return nil
}

// OSUnfairLockLockWithFlags.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_lock_with_flags
func OSUnfairLockLockWithFlags(lock unsafe.Pointer, flags OSUnfairLockFlags) {
	if callErr := tryOSUnfairLockLockWithFlags(lock, flags); callErr != nil {
		panic(callErr)
	}
}

var _oSUnfairLockTrylock func(lock unsafe.Pointer) bool
var _oSUnfairLockTrylockErr error

func tryOSUnfairLockTrylock(lock unsafe.Pointer) (bool, error) {
	if _oSUnfairLockTrylock == nil {
		return false, symbolCallError("os_unfair_lock_trylock", "10.12", _oSUnfairLockTrylockErr)
	}
	return _oSUnfairLockTrylock(lock), nil
}

// OSUnfairLockTrylock locks an unfair lock if it is not already locked.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_trylock
func OSUnfairLockTrylock(lock unsafe.Pointer) bool {
	result, callErr := tryOSUnfairLockTrylock(lock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSUnfairLockUnlock func(lock unsafe.Pointer)
var _oSUnfairLockUnlockErr error

func tryOSUnfairLockUnlock(lock unsafe.Pointer) error {
	if _oSUnfairLockUnlock == nil {
		return symbolCallError("os_unfair_lock_unlock", "10.12", _oSUnfairLockUnlockErr)
	}
	_oSUnfairLockUnlock(lock)
	return nil
}

// OSUnfairLockUnlock unlocks an unfair lock.
//
// See: https://developer.apple.com/documentation/os/os_unfair_lock_unlock
func OSUnfairLockUnlock(lock unsafe.Pointer) {
	if callErr := tryOSUnfairLockUnlock(lock); callErr != nil {
		panic(callErr)
	}
}

var _oSWorkgroupCancel func(wg OSWorkgroup)
var _oSWorkgroupCancelErr error

func tryOSWorkgroupCancel(wg OSWorkgroup) error {
	if _oSWorkgroupCancel == nil {
		return symbolCallError("os_workgroup_cancel", "11.0", _oSWorkgroupCancelErr)
	}
	_oSWorkgroupCancel(wg)
	return nil
}

// OSWorkgroupCancel cancels and invalidates the specified workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_cancel
func OSWorkgroupCancel(wg OSWorkgroup) {
	if callErr := tryOSWorkgroupCancel(wg); callErr != nil {
		panic(callErr)
	}
}

var _oSWorkgroupCopyPort func(wg OSWorkgroup, mach_port_out *uint32) int32
var _oSWorkgroupCopyPortErr error

func tryOSWorkgroupCopyPort(wg OSWorkgroup, mach_port_out *uint32) (int32, error) {
	if _oSWorkgroupCopyPort == nil {
		return 0, symbolCallError("os_workgroup_copy_port", "11.0", _oSWorkgroupCopyPortErr)
	}
	return _oSWorkgroupCopyPort(wg, mach_port_out), nil
}

// OSWorkgroupCopyPort returns the Mach port associated with the workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_copy_port
func OSWorkgroupCopyPort(wg OSWorkgroup, mach_port_out *uint32) int32 {
	result, callErr := tryOSWorkgroupCopyPort(wg, mach_port_out)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupCreateWithPort func(name string, mach_port uint32) OSWorkgroup
var _oSWorkgroupCreateWithPortErr error

func tryOSWorkgroupCreateWithPort(name string, mach_port uint32) (OSWorkgroup, error) {
	if _oSWorkgroupCreateWithPort == nil {
		return *new(OSWorkgroup), symbolCallError("os_workgroup_create_with_port", "11.0", _oSWorkgroupCreateWithPortErr)
	}
	return _oSWorkgroupCreateWithPort(name, mach_port), nil
}

// OSWorkgroupCreateWithPort creates a new workgroup that is bound to the workgroup with the specified Mach port.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_create_with_port
func OSWorkgroupCreateWithPort(name string, mach_port uint32) OSWorkgroup {
	result, callErr := tryOSWorkgroupCreateWithPort(name, mach_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupCreateWithWorkgroup func(name string, wg OSWorkgroup) OSWorkgroup
var _oSWorkgroupCreateWithWorkgroupErr error

func tryOSWorkgroupCreateWithWorkgroup(name string, wg OSWorkgroup) (OSWorkgroup, error) {
	if _oSWorkgroupCreateWithWorkgroup == nil {
		return *new(OSWorkgroup), symbolCallError("os_workgroup_create_with_workgroup", "11.0", _oSWorkgroupCreateWithWorkgroupErr)
	}
	return _oSWorkgroupCreateWithWorkgroup(name, wg), nil
}

// OSWorkgroupCreateWithWorkgroup create a new workgroup that is bound to the specified workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_create_with_workgroup
func OSWorkgroupCreateWithWorkgroup(name string, wg OSWorkgroup) OSWorkgroup {
	result, callErr := tryOSWorkgroupCreateWithWorkgroup(name, wg)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupGetWorkingArena func(wg OSWorkgroup, index_out *OSWorkgroupIndex) unsafe.Pointer
var _oSWorkgroupGetWorkingArenaErr error

func tryOSWorkgroupGetWorkingArena(wg OSWorkgroup, index_out *OSWorkgroupIndex) (unsafe.Pointer, error) {
	if _oSWorkgroupGetWorkingArena == nil {
		return nil, symbolCallError("os_workgroup_get_working_arena", "11.0", _oSWorkgroupGetWorkingArenaErr)
	}
	return _oSWorkgroupGetWorkingArena(wg, index_out), nil
}

// OSWorkgroupGetWorkingArena retrieves the workgroup’s shared data, and the thread-specific index into that data.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_get_working_arena
func OSWorkgroupGetWorkingArena(wg OSWorkgroup, index_out *OSWorkgroupIndex) unsafe.Pointer {
	result, callErr := tryOSWorkgroupGetWorkingArena(wg, index_out)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupIntervalFinish func(wg OSWorkgroupInterval, data OSWorkgroupIntervalData) int32
var _oSWorkgroupIntervalFinishErr error

func tryOSWorkgroupIntervalFinish(wg OSWorkgroupInterval, data OSWorkgroupIntervalData) (int32, error) {
	if _oSWorkgroupIntervalFinish == nil {
		return 0, symbolCallError("os_workgroup_interval_finish", "11.0", _oSWorkgroupIntervalFinishErr)
	}
	return _oSWorkgroupIntervalFinish(wg, data), nil
}

// OSWorkgroupIntervalFinish stops the current interval-based execution of the workgroup’s threads.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_interval_finish
func OSWorkgroupIntervalFinish(wg OSWorkgroupInterval, data OSWorkgroupIntervalData) int32 {
	result, callErr := tryOSWorkgroupIntervalFinish(wg, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupIntervalStart func(wg OSWorkgroupInterval, start uint64, deadline uint64, data OSWorkgroupIntervalData) int32
var _oSWorkgroupIntervalStartErr error

func tryOSWorkgroupIntervalStart(wg OSWorkgroupInterval, start uint64, deadline uint64, data OSWorkgroupIntervalData) (int32, error) {
	if _oSWorkgroupIntervalStart == nil {
		return 0, symbolCallError("os_workgroup_interval_start", "11.0", _oSWorkgroupIntervalStartErr)
	}
	return _oSWorkgroupIntervalStart(wg, start, deadline, data), nil
}

// OSWorkgroupIntervalStart starts the regular execution of the workgroup’s threads at the specified time.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_interval_start
func OSWorkgroupIntervalStart(wg OSWorkgroupInterval, start uint64, deadline uint64, data OSWorkgroupIntervalData) int32 {
	result, callErr := tryOSWorkgroupIntervalStart(wg, start, deadline, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupIntervalUpdate func(wg OSWorkgroupInterval, deadline uint64, data OSWorkgroupIntervalData) int32
var _oSWorkgroupIntervalUpdateErr error

func tryOSWorkgroupIntervalUpdate(wg OSWorkgroupInterval, deadline uint64, data OSWorkgroupIntervalData) (int32, error) {
	if _oSWorkgroupIntervalUpdate == nil {
		return 0, symbolCallError("os_workgroup_interval_update", "11.0", _oSWorkgroupIntervalUpdateErr)
	}
	return _oSWorkgroupIntervalUpdate(wg, deadline, data), nil
}

// OSWorkgroupIntervalUpdate schedules a new deadline for workgroup threads that run at regular intervals.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_interval_update
func OSWorkgroupIntervalUpdate(wg OSWorkgroupInterval, deadline uint64, data OSWorkgroupIntervalData) int32 {
	result, callErr := tryOSWorkgroupIntervalUpdate(wg, deadline, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupJoin func(wg OSWorkgroup, token_out OSWorkgroupJoinToken) int32
var _oSWorkgroupJoinErr error

func tryOSWorkgroupJoin(wg OSWorkgroup, token_out OSWorkgroupJoinToken) (int32, error) {
	if _oSWorkgroupJoin == nil {
		return 0, symbolCallError("os_workgroup_join", "11.0", _oSWorkgroupJoinErr)
	}
	return _oSWorkgroupJoin(wg, token_out), nil
}

// OSWorkgroupJoin adds the current thread to the specified workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_join
func OSWorkgroupJoin(wg OSWorkgroup, token_out OSWorkgroupJoinToken) int32 {
	result, callErr := tryOSWorkgroupJoin(wg, token_out)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupLeave func(wg OSWorkgroup, token OSWorkgroupJoinToken)
var _oSWorkgroupLeaveErr error

func tryOSWorkgroupLeave(wg OSWorkgroup, token OSWorkgroupJoinToken) error {
	if _oSWorkgroupLeave == nil {
		return symbolCallError("os_workgroup_leave", "11.0", _oSWorkgroupLeaveErr)
	}
	_oSWorkgroupLeave(wg, token)
	return nil
}

// OSWorkgroupLeave removes the current thread from the workgroup it previously joined.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_leave
func OSWorkgroupLeave(wg OSWorkgroup, token OSWorkgroupJoinToken) {
	if callErr := tryOSWorkgroupLeave(wg, token); callErr != nil {
		panic(callErr)
	}
}

var _oSWorkgroupMaxParallelThreads func(wg OSWorkgroup, attr OSWorkgroupMptAttr) int32
var _oSWorkgroupMaxParallelThreadsErr error

func tryOSWorkgroupMaxParallelThreads(wg OSWorkgroup, attr OSWorkgroupMptAttr) (int32, error) {
	if _oSWorkgroupMaxParallelThreads == nil {
		return 0, symbolCallError("os_workgroup_max_parallel_threads", "11.0", _oSWorkgroupMaxParallelThreadsErr)
	}
	return _oSWorkgroupMaxParallelThreads(wg, attr), nil
}

// OSWorkgroupMaxParallelThreads returns the maximum number of threads that the system recommends you add to the specified workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_max_parallel_threads
func OSWorkgroupMaxParallelThreads(wg OSWorkgroup, attr OSWorkgroupMptAttr) int32 {
	result, callErr := tryOSWorkgroupMaxParallelThreads(wg, attr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupParallelCreate func(name string, attr OSWorkgroupAttr) OSWorkgroupParallel
var _oSWorkgroupParallelCreateErr error

func tryOSWorkgroupParallelCreate(name string, attr OSWorkgroupAttr) (OSWorkgroupParallel, error) {
	if _oSWorkgroupParallelCreate == nil {
		return *new(OSWorkgroupParallel), symbolCallError("os_workgroup_parallel_create", "11.0", _oSWorkgroupParallelCreateErr)
	}
	return _oSWorkgroupParallelCreate(name, attr), nil
}

// OSWorkgroupParallelCreate creates a new workgroup that manges threads working on a single task in parallel.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_parallel_create
func OSWorkgroupParallelCreate(name string, attr OSWorkgroupAttr) OSWorkgroupParallel {
	result, callErr := tryOSWorkgroupParallelCreate(name, attr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupSetWorkingArena func(wg OSWorkgroup, arena unsafe.Pointer, max_workers uint32, destructor OSWorkgroupWorkingArenaDestructor) int32
var _oSWorkgroupSetWorkingArenaErr error

func tryOSWorkgroupSetWorkingArena(wg OSWorkgroup, arena unsafe.Pointer, max_workers uint32, destructor OSWorkgroupWorkingArenaDestructor) (int32, error) {
	if _oSWorkgroupSetWorkingArena == nil {
		return 0, symbolCallError("os_workgroup_set_working_arena", "11.0", _oSWorkgroupSetWorkingArenaErr)
	}
	return _oSWorkgroupSetWorkingArena(wg, arena, max_workers, destructor), nil
}

// OSWorkgroupSetWorkingArena distributes a block of managed memory to the threads of a workgroup.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_set_working_arena
func OSWorkgroupSetWorkingArena(wg OSWorkgroup, arena unsafe.Pointer, max_workers uint32, destructor OSWorkgroupWorkingArenaDestructor) int32 {
	result, callErr := tryOSWorkgroupSetWorkingArena(wg, arena, max_workers, destructor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oSWorkgroupTestcancel func(wg OSWorkgroup) bool
var _oSWorkgroupTestcancelErr error

func tryOSWorkgroupTestcancel(wg OSWorkgroup) (bool, error) {
	if _oSWorkgroupTestcancel == nil {
		return false, symbolCallError("os_workgroup_testcancel", "11.0", _oSWorkgroupTestcancelErr)
	}
	return _oSWorkgroupTestcancel(wg), nil
}

// OSWorkgroupTestcancel returns a Boolean value that indicates whether the workgroup is canceled.
//
// See: https://developer.apple.com/documentation/os/os_workgroup_testcancel
func OSWorkgroupTestcancel(wg OSWorkgroup) bool {
	result, callErr := tryOSWorkgroupTestcancel(wg)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_oSActivityApply, &_oSActivityApplyErr, frameworkHandle, "os_activity_apply", "10.12")
	registerFunc(&_oSActivityApplyF, &_oSActivityApplyFErr, frameworkHandle, "os_activity_apply_f", "10.12")
	registerFunc(&_oSActivityGetIdentifier, &_oSActivityGetIdentifierErr, frameworkHandle, "os_activity_get_identifier", "10.12")
	registerFunc(&_oSActivityScopeEnter, &_oSActivityScopeEnterErr, frameworkHandle, "os_activity_scope_enter", "10.12")
	registerFunc(&_oSActivityScopeLeave, &_oSActivityScopeLeaveErr, frameworkHandle, "os_activity_scope_leave", "10.12")
	registerFunc(&_oSLogCreate, &_oSLogCreateErr, frameworkHandle, "os_log_create", "10.12")
	registerFunc(&_oSLogTypeEnabled, &_oSLogTypeEnabledErr, frameworkHandle, "os_log_type_enabled", "10.12")
	registerFunc(&_oSRelease, &_oSReleaseErr, frameworkHandle, "os_release", "10.10")
	registerFunc(&_oSRetain, &_oSRetainErr, frameworkHandle, "os_retain", "10.10")
	registerFunc(&_oSSecurityConfigGet, &_oSSecurityConfigGetErr, frameworkHandle, "os_security_config_get", "26.0")
	registerFunc(&_oSSecurityConfigGetForProc, &_oSSecurityConfigGetForProcErr, frameworkHandle, "os_security_config_get_for_proc", "26.0")
	registerFunc(&_oSSecurityConfigGetForTask, &_oSSecurityConfigGetForTaskErr, frameworkHandle, "os_security_config_get_for_task", "26.0")
	registerFunc(&_oSSignpostEnabled, &_oSSignpostEnabledErr, frameworkHandle, "os_signpost_enabled", "10.14")
	registerFunc(&_oSSignpostIDGenerate, &_oSSignpostIDGenerateErr, frameworkHandle, "os_signpost_id_generate", "10.14")
	registerFunc(&_oSSignpostIDMakeWithPointer, &_oSSignpostIDMakeWithPointerErr, frameworkHandle, "os_signpost_id_make_with_pointer", "10.14")
	registerFunc(&_oSSyncWaitOnAddress, &_oSSyncWaitOnAddressErr, frameworkHandle, "os_sync_wait_on_address", "14.4")
	registerFunc(&_oSSyncWaitOnAddressWithDeadline, &_oSSyncWaitOnAddressWithDeadlineErr, frameworkHandle, "os_sync_wait_on_address_with_deadline", "14.4")
	registerFunc(&_oSSyncWaitOnAddressWithTimeout, &_oSSyncWaitOnAddressWithTimeoutErr, frameworkHandle, "os_sync_wait_on_address_with_timeout", "14.4")
	registerFunc(&_oSSyncWakeByAddressAll, &_oSSyncWakeByAddressAllErr, frameworkHandle, "os_sync_wake_by_address_all", "14.4")
	registerFunc(&_oSSyncWakeByAddressAny, &_oSSyncWakeByAddressAnyErr, frameworkHandle, "os_sync_wake_by_address_any", "14.4")
	registerFunc(&_oSUnfairLockAssertNotOwner, &_oSUnfairLockAssertNotOwnerErr, frameworkHandle, "os_unfair_lock_assert_not_owner", "10.12")
	registerFunc(&_oSUnfairLockAssertOwner, &_oSUnfairLockAssertOwnerErr, frameworkHandle, "os_unfair_lock_assert_owner", "10.12")
	registerFunc(&_oSUnfairLockLock, &_oSUnfairLockLockErr, frameworkHandle, "os_unfair_lock_lock", "10.12")
	registerFunc(&_oSUnfairLockLockWithFlags, &_oSUnfairLockLockWithFlagsErr, frameworkHandle, "os_unfair_lock_lock_with_flags", "15.0")
	registerFunc(&_oSUnfairLockTrylock, &_oSUnfairLockTrylockErr, frameworkHandle, "os_unfair_lock_trylock", "10.12")
	registerFunc(&_oSUnfairLockUnlock, &_oSUnfairLockUnlockErr, frameworkHandle, "os_unfair_lock_unlock", "10.12")
	registerFunc(&_oSWorkgroupCancel, &_oSWorkgroupCancelErr, frameworkHandle, "os_workgroup_cancel", "11.0")
	registerFunc(&_oSWorkgroupCopyPort, &_oSWorkgroupCopyPortErr, frameworkHandle, "os_workgroup_copy_port", "11.0")
	registerFunc(&_oSWorkgroupCreateWithPort, &_oSWorkgroupCreateWithPortErr, frameworkHandle, "os_workgroup_create_with_port", "11.0")
	registerFunc(&_oSWorkgroupCreateWithWorkgroup, &_oSWorkgroupCreateWithWorkgroupErr, frameworkHandle, "os_workgroup_create_with_workgroup", "11.0")
	registerFunc(&_oSWorkgroupGetWorkingArena, &_oSWorkgroupGetWorkingArenaErr, frameworkHandle, "os_workgroup_get_working_arena", "11.0")
	registerFunc(&_oSWorkgroupIntervalFinish, &_oSWorkgroupIntervalFinishErr, frameworkHandle, "os_workgroup_interval_finish", "11.0")
	registerFunc(&_oSWorkgroupIntervalStart, &_oSWorkgroupIntervalStartErr, frameworkHandle, "os_workgroup_interval_start", "11.0")
	registerFunc(&_oSWorkgroupIntervalUpdate, &_oSWorkgroupIntervalUpdateErr, frameworkHandle, "os_workgroup_interval_update", "11.0")
	registerFunc(&_oSWorkgroupJoin, &_oSWorkgroupJoinErr, frameworkHandle, "os_workgroup_join", "11.0")
	registerFunc(&_oSWorkgroupLeave, &_oSWorkgroupLeaveErr, frameworkHandle, "os_workgroup_leave", "11.0")
	registerFunc(&_oSWorkgroupMaxParallelThreads, &_oSWorkgroupMaxParallelThreadsErr, frameworkHandle, "os_workgroup_max_parallel_threads", "11.0")
	registerFunc(&_oSWorkgroupParallelCreate, &_oSWorkgroupParallelCreateErr, frameworkHandle, "os_workgroup_parallel_create", "11.0")
	registerFunc(&_oSWorkgroupSetWorkingArena, &_oSWorkgroupSetWorkingArenaErr, frameworkHandle, "os_workgroup_set_working_arena", "11.0")
	registerFunc(&_oSWorkgroupTestcancel, &_oSWorkgroupTestcancelErr, frameworkHandle, "os_workgroup_testcancel", "11.0")
}

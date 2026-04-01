// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/ebitengine/purego"
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
		return fmt.Sprintf("Dispatch: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("Dispatch: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("Dispatch: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("Dispatch: register symbol %s: %v", name, r)
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

var _dispatch_activate func(object uintptr)
var _dispatch_activateErr error

var _dispatch_after func(when uint64, queue uintptr, block unsafe.Pointer)
var _dispatch_afterErr error

var _dispatch_after_f func(when uint64, queue uintptr, context unsafe.Pointer, work unsafe.Pointer)
var _dispatch_after_fErr error

var _dispatch_allow_send_signals func(preserve_signum int) int
var _dispatch_allow_send_signalsErr error

var _dispatch_apply func(iterations uintptr, queue uintptr, work unsafe.Pointer)
var _dispatch_applyErr error

var _dispatch_apply_f func(iterations uintptr, queue uintptr, context unsafe.Pointer, work unsafe.Pointer)
var _dispatch_apply_fErr error

var _dispatch_assert_queue func(queue uintptr)
var _dispatch_assert_queueErr error

var _dispatch_assert_queue_barrier func(queue uintptr)
var _dispatch_assert_queue_barrierErr error

var _dispatch_assert_queue_not func(queue uintptr)
var _dispatch_assert_queue_notErr error

var _dispatch_async func(queue uintptr, block unsafe.Pointer)
var _dispatch_asyncErr error

var _dispatch_async_and_wait func(queue uintptr, block unsafe.Pointer)
var _dispatch_async_and_waitErr error

var _dispatch_async_and_wait_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)
var _dispatch_async_and_wait_fErr error

var _dispatch_async_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)
var _dispatch_async_fErr error

var _dispatch_barrier_async func(queue uintptr, block unsafe.Pointer)
var _dispatch_barrier_asyncErr error

var _dispatch_barrier_async_and_wait func(queue uintptr, block unsafe.Pointer)
var _dispatch_barrier_async_and_waitErr error

var _dispatch_barrier_async_and_wait_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)
var _dispatch_barrier_async_and_wait_fErr error

var _dispatch_barrier_async_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)
var _dispatch_barrier_async_fErr error

var _dispatch_barrier_sync func(queue uintptr, block unsafe.Pointer)
var _dispatch_barrier_syncErr error

var _dispatch_barrier_sync_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)
var _dispatch_barrier_sync_fErr error

var _dispatch_block_cancel func(block unsafe.Pointer)
var _dispatch_block_cancelErr error

var _dispatch_block_create func(flags uintptr, block unsafe.Pointer) unsafe.Pointer
var _dispatch_block_createErr error

var _dispatch_block_create_with_qos_class func(flags uintptr, qos_class uint32, relative_priority int, block unsafe.Pointer) unsafe.Pointer
var _dispatch_block_create_with_qos_classErr error

var _dispatch_block_notify func(block unsafe.Pointer, queue uintptr, notification_block unsafe.Pointer)
var _dispatch_block_notifyErr error

var _dispatch_block_perform func(flags uintptr, block unsafe.Pointer)
var _dispatch_block_performErr error

var _dispatch_block_testcancel func(block unsafe.Pointer) int
var _dispatch_block_testcancelErr error

var _dispatch_block_wait func(block unsafe.Pointer, timeout uint64) int
var _dispatch_block_waitErr error

var _dispatch_data_apply func(data uintptr, applier unsafe.Pointer) bool
var _dispatch_data_applyErr error

var _dispatch_data_copy_region func(data uintptr, location uintptr, offset_ptr *uintptr) uintptr
var _dispatch_data_copy_regionErr error

var _dispatch_data_create func(buffer unsafe.Pointer, size uintptr, queue uintptr, destructor unsafe.Pointer) uintptr
var _dispatch_data_createErr error

var _dispatch_data_create_concat func(data1 uintptr, data2 uintptr) uintptr
var _dispatch_data_create_concatErr error

var _dispatch_data_create_map func(data uintptr, buffer_ptr *unsafe.Pointer, size_ptr *uintptr) uintptr
var _dispatch_data_create_mapErr error

var _dispatch_data_create_subrange func(data uintptr, offset uintptr, length uintptr) uintptr
var _dispatch_data_create_subrangeErr error

var _dispatch_data_get_size func(data uintptr) uintptr
var _dispatch_data_get_sizeErr error

var _dispatch_debug func(object uintptr, message string)
var _dispatch_debugErr error

var _dispatch_get_context func(object uintptr) uintptr
var _dispatch_get_contextErr error

var _dispatch_get_global_queue func(identifier int, flags uintptr) uintptr
var _dispatch_get_global_queueErr error

var _dispatch_get_specific func(key unsafe.Pointer) unsafe.Pointer
var _dispatch_get_specificErr error

var _dispatch_group_async func(group uintptr, queue uintptr, block unsafe.Pointer)
var _dispatch_group_asyncErr error

var _dispatch_group_async_f func(group uintptr, queue uintptr, context unsafe.Pointer, work unsafe.Pointer)
var _dispatch_group_async_fErr error

var _dispatch_group_create func() uintptr
var _dispatch_group_createErr error

var _dispatch_group_enter func(group uintptr)
var _dispatch_group_enterErr error

var _dispatch_group_leave func(group uintptr)
var _dispatch_group_leaveErr error

var _dispatch_group_notify func(group uintptr, queue uintptr, block unsafe.Pointer)
var _dispatch_group_notifyErr error

var _dispatch_group_notify_f func(group uintptr, queue uintptr, context unsafe.Pointer, work unsafe.Pointer)
var _dispatch_group_notify_fErr error

var _dispatch_group_wait func(group uintptr, timeout uint64) int
var _dispatch_group_waitErr error

var _dispatch_io_barrier func(channel uintptr, barrier unsafe.Pointer)
var _dispatch_io_barrierErr error

var _dispatch_io_close func(channel uintptr, flags uint)
var _dispatch_io_closeErr error

var _dispatch_io_create func(type_ uint, fd int32, queue uintptr, cleanup_handler uintptr) uintptr
var _dispatch_io_createErr error

var _dispatch_io_create_with_io func(type_ uint, io uintptr, queue uintptr, cleanup_handler uintptr) uintptr
var _dispatch_io_create_with_ioErr error

var _dispatch_io_create_with_path func(type_ uint, path string, oflag int, mode uint16, queue uintptr, cleanup_handler uintptr) uintptr
var _dispatch_io_create_with_pathErr error

var _dispatch_io_get_descriptor func(channel uintptr) int32
var _dispatch_io_get_descriptorErr error

var _dispatch_io_read func(channel uintptr, offset int64, length uintptr, queue uintptr, io_handler unsafe.Pointer)
var _dispatch_io_readErr error

var _dispatch_io_set_high_water func(channel uintptr, high_water uintptr)
var _dispatch_io_set_high_waterErr error

var _dispatch_io_set_interval func(channel uintptr, interval uint64, flags uint)
var _dispatch_io_set_intervalErr error

var _dispatch_io_set_low_water func(channel uintptr, low_water uintptr)
var _dispatch_io_set_low_waterErr error

var _dispatch_io_write func(channel uintptr, offset int64, data uintptr, queue uintptr, io_handler unsafe.Pointer)
var _dispatch_io_writeErr error

var _dispatch_main func()
var _dispatch_mainErr error

var _dispatch_once func(predicate *int, block unsafe.Pointer)
var _dispatch_onceErr error

var _dispatch_once_f func(predicate *int, context unsafe.Pointer, function unsafe.Pointer)
var _dispatch_once_fErr error

var _dispatch_queue_attr_make_initially_inactive func(attr unsafe.Pointer) unsafe.Pointer
var _dispatch_queue_attr_make_initially_inactiveErr error

var _dispatch_queue_attr_make_with_autorelease_frequency func(attr unsafe.Pointer, frequency uintptr) unsafe.Pointer
var _dispatch_queue_attr_make_with_autorelease_frequencyErr error

var _dispatch_queue_attr_make_with_qos_class func(attr unsafe.Pointer, qos_class uint32, relative_priority int) unsafe.Pointer
var _dispatch_queue_attr_make_with_qos_classErr error

var _dispatch_queue_create func(label string, attr dispatch_queue_attr_t) uintptr
var _dispatch_queue_createErr error

var _dispatch_queue_create_with_target func(label string, attr unsafe.Pointer, target uintptr) uintptr
var _dispatch_queue_create_with_targetErr error

var _dispatch_queue_get_label func(queue uintptr) *byte
var _dispatch_queue_get_labelErr error

var _dispatch_queue_get_qos_class func(queue uintptr, relative_priority_ptr *int32) uint32
var _dispatch_queue_get_qos_classErr error

var _dispatch_queue_get_specific func(queue uintptr, key unsafe.Pointer) unsafe.Pointer
var _dispatch_queue_get_specificErr error

var _dispatch_queue_set_specific func(queue uintptr, key unsafe.Pointer, context unsafe.Pointer, destructor unsafe.Pointer)
var _dispatch_queue_set_specificErr error

var _dispatch_read func(fd int32, length uintptr, queue uintptr, io_handler unsafe.Pointer)
var _dispatch_readErr error

var _dispatch_release func(object uintptr)
var _dispatch_releaseErr error

var _dispatch_resume func(object uintptr)
var _dispatch_resumeErr error

var _dispatch_retain func(object uintptr)
var _dispatch_retainErr error

var _dispatch_semaphore_create func(value int) uintptr
var _dispatch_semaphore_createErr error

var _dispatch_semaphore_signal func(dsema uintptr) int
var _dispatch_semaphore_signalErr error

var _dispatch_semaphore_wait func(dsema uintptr, timeout uint64) int
var _dispatch_semaphore_waitErr error

var _dispatch_set_context func(object uintptr, context uintptr)
var _dispatch_set_contextErr error

var _dispatch_set_finalizer_f func(object uintptr, finalizer unsafe.Pointer)
var _dispatch_set_finalizer_fErr error

var _dispatch_set_qos_class_floor func(object uintptr, qos_class uint32, relative_priority int)
var _dispatch_set_qos_class_floorErr error

var _dispatch_set_target_queue func(object uintptr, queue uintptr)
var _dispatch_set_target_queueErr error

var _dispatch_source_cancel func(source uintptr)
var _dispatch_source_cancelErr error

var _dispatch_source_create func(type_ Dispatch_source_type_t, handle uintptr, mask uintptr, queue uintptr) uintptr
var _dispatch_source_createErr error

var _dispatch_source_get_data func(source uintptr) uintptr
var _dispatch_source_get_dataErr error

var _dispatch_source_get_handle func(source uintptr) uintptr
var _dispatch_source_get_handleErr error

var _dispatch_source_get_mask func(source uintptr) uintptr
var _dispatch_source_get_maskErr error

var _dispatch_source_merge_data func(source uintptr, value uintptr)
var _dispatch_source_merge_dataErr error

var _dispatch_source_set_cancel_handler func(source uintptr, handler unsafe.Pointer)
var _dispatch_source_set_cancel_handlerErr error

var _dispatch_source_set_cancel_handler_f func(source uintptr, handler unsafe.Pointer)
var _dispatch_source_set_cancel_handler_fErr error

var _dispatch_source_set_event_handler func(source uintptr, handler unsafe.Pointer)
var _dispatch_source_set_event_handlerErr error

var _dispatch_source_set_event_handler_f func(source uintptr, handler unsafe.Pointer)
var _dispatch_source_set_event_handler_fErr error

var _dispatch_source_set_registration_handler func(source uintptr, handler unsafe.Pointer)
var _dispatch_source_set_registration_handlerErr error

var _dispatch_source_set_registration_handler_f func(source uintptr, handler unsafe.Pointer)
var _dispatch_source_set_registration_handler_fErr error

var _dispatch_source_set_timer func(source uintptr, start uint64, interval uint64, leeway uint64)
var _dispatch_source_set_timerErr error

var _dispatch_source_testcancel func(source uintptr) int
var _dispatch_source_testcancelErr error

var _dispatch_suspend func(object uintptr)
var _dispatch_suspendErr error

var _dispatch_sync func(queue uintptr, block unsafe.Pointer)
var _dispatch_syncErr error

var _dispatch_sync_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)
var _dispatch_sync_fErr error

var _dispatch_time func(when uint64, delta int64) uint64
var _dispatch_timeErr error

var _dispatch_walltime func(when *syscall.Timespec, delta int64) uint64
var _dispatch_walltimeErr error

var _dispatch_workloop_create func(label string) uintptr
var _dispatch_workloop_createErr error

var _dispatch_workloop_create_inactive func(label string) uintptr
var _dispatch_workloop_create_inactiveErr error

var _dispatch_workloop_set_autorelease_frequency func(workloop uintptr, frequency uintptr)
var _dispatch_workloop_set_autorelease_frequencyErr error

var _dispatch_workloop_set_os_workgroup func(workloop uintptr, workgroup unsafe.Pointer)
var _dispatch_workloop_set_os_workgroupErr error

var _dispatch_write func(fd int32, data uintptr, queue uintptr, io_handler unsafe.Pointer)
var _dispatch_writeErr error

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_dispatch_activate, &_dispatch_activateErr, frameworkHandle, "dispatch_activate", "10.12")
	registerFunc(&_dispatch_after, &_dispatch_afterErr, frameworkHandle, "dispatch_after", "10.6")
	registerFunc(&_dispatch_after_f, &_dispatch_after_fErr, frameworkHandle, "dispatch_after_f", "10.6")
	registerFunc(&_dispatch_allow_send_signals, &_dispatch_allow_send_signalsErr, frameworkHandle, "dispatch_allow_send_signals", "14.4")
	registerFunc(&_dispatch_apply, &_dispatch_applyErr, frameworkHandle, "dispatch_apply", "10.6")
	registerFunc(&_dispatch_apply_f, &_dispatch_apply_fErr, frameworkHandle, "dispatch_apply_f", "10.6")
	registerFunc(&_dispatch_assert_queue, &_dispatch_assert_queueErr, frameworkHandle, "dispatch_assert_queue", "10.12")
	registerFunc(&_dispatch_assert_queue_barrier, &_dispatch_assert_queue_barrierErr, frameworkHandle, "dispatch_assert_queue_barrier", "10.12")
	registerFunc(&_dispatch_assert_queue_not, &_dispatch_assert_queue_notErr, frameworkHandle, "dispatch_assert_queue_not", "10.12")
	registerFunc(&_dispatch_async, &_dispatch_asyncErr, frameworkHandle, "dispatch_async", "10.6")
	registerFunc(&_dispatch_async_and_wait, &_dispatch_async_and_waitErr, frameworkHandle, "dispatch_async_and_wait", "10.14")
	registerFunc(&_dispatch_async_and_wait_f, &_dispatch_async_and_wait_fErr, frameworkHandle, "dispatch_async_and_wait_f", "10.14")
	registerFunc(&_dispatch_async_f, &_dispatch_async_fErr, frameworkHandle, "dispatch_async_f", "10.6")
	registerFunc(&_dispatch_barrier_async, &_dispatch_barrier_asyncErr, frameworkHandle, "dispatch_barrier_async", "10.7")
	registerFunc(&_dispatch_barrier_async_and_wait, &_dispatch_barrier_async_and_waitErr, frameworkHandle, "dispatch_barrier_async_and_wait", "10.14")
	registerFunc(&_dispatch_barrier_async_and_wait_f, &_dispatch_barrier_async_and_wait_fErr, frameworkHandle, "dispatch_barrier_async_and_wait_f", "10.14")
	registerFunc(&_dispatch_barrier_async_f, &_dispatch_barrier_async_fErr, frameworkHandle, "dispatch_barrier_async_f", "10.7")
	registerFunc(&_dispatch_barrier_sync, &_dispatch_barrier_syncErr, frameworkHandle, "dispatch_barrier_sync", "10.7")
	registerFunc(&_dispatch_barrier_sync_f, &_dispatch_barrier_sync_fErr, frameworkHandle, "dispatch_barrier_sync_f", "10.7")
	registerFunc(&_dispatch_block_cancel, &_dispatch_block_cancelErr, frameworkHandle, "dispatch_block_cancel", "10.10")
	registerFunc(&_dispatch_block_create, &_dispatch_block_createErr, frameworkHandle, "dispatch_block_create", "10.10")
	registerFunc(&_dispatch_block_create_with_qos_class, &_dispatch_block_create_with_qos_classErr, frameworkHandle, "dispatch_block_create_with_qos_class", "10.10")
	registerFunc(&_dispatch_block_notify, &_dispatch_block_notifyErr, frameworkHandle, "dispatch_block_notify", "10.10")
	registerFunc(&_dispatch_block_perform, &_dispatch_block_performErr, frameworkHandle, "dispatch_block_perform", "10.10")
	registerFunc(&_dispatch_block_testcancel, &_dispatch_block_testcancelErr, frameworkHandle, "dispatch_block_testcancel", "10.10")
	registerFunc(&_dispatch_block_wait, &_dispatch_block_waitErr, frameworkHandle, "dispatch_block_wait", "10.10")
	registerFunc(&_dispatch_data_apply, &_dispatch_data_applyErr, frameworkHandle, "dispatch_data_apply", "10.7")
	registerFunc(&_dispatch_data_copy_region, &_dispatch_data_copy_regionErr, frameworkHandle, "dispatch_data_copy_region", "10.7")
	registerFunc(&_dispatch_data_create, &_dispatch_data_createErr, frameworkHandle, "dispatch_data_create", "10.7")
	registerFunc(&_dispatch_data_create_concat, &_dispatch_data_create_concatErr, frameworkHandle, "dispatch_data_create_concat", "10.7")
	registerFunc(&_dispatch_data_create_map, &_dispatch_data_create_mapErr, frameworkHandle, "dispatch_data_create_map", "10.7")
	registerFunc(&_dispatch_data_create_subrange, &_dispatch_data_create_subrangeErr, frameworkHandle, "dispatch_data_create_subrange", "10.7")
	registerFunc(&_dispatch_data_get_size, &_dispatch_data_get_sizeErr, frameworkHandle, "dispatch_data_get_size", "10.7")
	registerFunc(&_dispatch_debug, &_dispatch_debugErr, frameworkHandle, "dispatch_debug", "10.6")
	registerFunc(&_dispatch_get_context, &_dispatch_get_contextErr, frameworkHandle, "dispatch_get_context", "10.6")
	registerFunc(&_dispatch_get_global_queue, &_dispatch_get_global_queueErr, frameworkHandle, "dispatch_get_global_queue", "10.6")
	registerFunc(&_dispatch_get_specific, &_dispatch_get_specificErr, frameworkHandle, "dispatch_get_specific", "10.7")
	registerFunc(&_dispatch_group_async, &_dispatch_group_asyncErr, frameworkHandle, "dispatch_group_async", "10.6")
	registerFunc(&_dispatch_group_async_f, &_dispatch_group_async_fErr, frameworkHandle, "dispatch_group_async_f", "10.6")
	registerFunc(&_dispatch_group_create, &_dispatch_group_createErr, frameworkHandle, "dispatch_group_create", "10.6")
	registerFunc(&_dispatch_group_enter, &_dispatch_group_enterErr, frameworkHandle, "dispatch_group_enter", "10.6")
	registerFunc(&_dispatch_group_leave, &_dispatch_group_leaveErr, frameworkHandle, "dispatch_group_leave", "10.6")
	registerFunc(&_dispatch_group_notify, &_dispatch_group_notifyErr, frameworkHandle, "dispatch_group_notify", "10.6")
	registerFunc(&_dispatch_group_notify_f, &_dispatch_group_notify_fErr, frameworkHandle, "dispatch_group_notify_f", "10.6")
	registerFunc(&_dispatch_group_wait, &_dispatch_group_waitErr, frameworkHandle, "dispatch_group_wait", "10.6")
	registerFunc(&_dispatch_io_barrier, &_dispatch_io_barrierErr, frameworkHandle, "dispatch_io_barrier", "10.7")
	registerFunc(&_dispatch_io_close, &_dispatch_io_closeErr, frameworkHandle, "dispatch_io_close", "10.7")
	registerFunc(&_dispatch_io_create, &_dispatch_io_createErr, frameworkHandle, "dispatch_io_create", "10.7")
	registerFunc(&_dispatch_io_create_with_io, &_dispatch_io_create_with_ioErr, frameworkHandle, "dispatch_io_create_with_io", "10.7")
	registerFunc(&_dispatch_io_create_with_path, &_dispatch_io_create_with_pathErr, frameworkHandle, "dispatch_io_create_with_path", "10.7")
	registerFunc(&_dispatch_io_get_descriptor, &_dispatch_io_get_descriptorErr, frameworkHandle, "dispatch_io_get_descriptor", "10.7")
	registerFunc(&_dispatch_io_read, &_dispatch_io_readErr, frameworkHandle, "dispatch_io_read", "10.7")
	registerFunc(&_dispatch_io_set_high_water, &_dispatch_io_set_high_waterErr, frameworkHandle, "dispatch_io_set_high_water", "10.7")
	registerFunc(&_dispatch_io_set_interval, &_dispatch_io_set_intervalErr, frameworkHandle, "dispatch_io_set_interval", "10.7")
	registerFunc(&_dispatch_io_set_low_water, &_dispatch_io_set_low_waterErr, frameworkHandle, "dispatch_io_set_low_water", "10.7")
	registerFunc(&_dispatch_io_write, &_dispatch_io_writeErr, frameworkHandle, "dispatch_io_write", "10.7")
	registerFunc(&_dispatch_main, &_dispatch_mainErr, frameworkHandle, "dispatch_main", "10.6")
	registerFunc(&_dispatch_once, &_dispatch_onceErr, frameworkHandle, "dispatch_once", "10.6")
	registerFunc(&_dispatch_once_f, &_dispatch_once_fErr, frameworkHandle, "dispatch_once_f", "10.6")
	registerFunc(&_dispatch_queue_attr_make_initially_inactive, &_dispatch_queue_attr_make_initially_inactiveErr, frameworkHandle, "dispatch_queue_attr_make_initially_inactive", "10.12")
	registerFunc(&_dispatch_queue_attr_make_with_autorelease_frequency, &_dispatch_queue_attr_make_with_autorelease_frequencyErr, frameworkHandle, "dispatch_queue_attr_make_with_autorelease_frequency", "10.12")
	registerFunc(&_dispatch_queue_attr_make_with_qos_class, &_dispatch_queue_attr_make_with_qos_classErr, frameworkHandle, "dispatch_queue_attr_make_with_qos_class", "10.10")
	registerFunc(&_dispatch_queue_create, &_dispatch_queue_createErr, frameworkHandle, "dispatch_queue_create", "10.6")
	registerFunc(&_dispatch_queue_create_with_target, &_dispatch_queue_create_with_targetErr, frameworkHandle, "dispatch_queue_create_with_target", "10.12")
	registerFunc(&_dispatch_queue_get_label, &_dispatch_queue_get_labelErr, frameworkHandle, "dispatch_queue_get_label", "10.6")
	registerFunc(&_dispatch_queue_get_qos_class, &_dispatch_queue_get_qos_classErr, frameworkHandle, "dispatch_queue_get_qos_class", "10.10")
	registerFunc(&_dispatch_queue_get_specific, &_dispatch_queue_get_specificErr, frameworkHandle, "dispatch_queue_get_specific", "10.7")
	registerFunc(&_dispatch_queue_set_specific, &_dispatch_queue_set_specificErr, frameworkHandle, "dispatch_queue_set_specific", "10.7")
	registerFunc(&_dispatch_read, &_dispatch_readErr, frameworkHandle, "dispatch_read", "10.7")
	registerFunc(&_dispatch_release, &_dispatch_releaseErr, frameworkHandle, "dispatch_release", "10.6")
	registerFunc(&_dispatch_resume, &_dispatch_resumeErr, frameworkHandle, "dispatch_resume", "10.6")
	registerFunc(&_dispatch_retain, &_dispatch_retainErr, frameworkHandle, "dispatch_retain", "10.6")
	registerFunc(&_dispatch_semaphore_create, &_dispatch_semaphore_createErr, frameworkHandle, "dispatch_semaphore_create", "10.6")
	registerFunc(&_dispatch_semaphore_signal, &_dispatch_semaphore_signalErr, frameworkHandle, "dispatch_semaphore_signal", "10.6")
	registerFunc(&_dispatch_semaphore_wait, &_dispatch_semaphore_waitErr, frameworkHandle, "dispatch_semaphore_wait", "10.6")
	registerFunc(&_dispatch_set_context, &_dispatch_set_contextErr, frameworkHandle, "dispatch_set_context", "10.6")
	registerFunc(&_dispatch_set_finalizer_f, &_dispatch_set_finalizer_fErr, frameworkHandle, "dispatch_set_finalizer_f", "10.6")
	registerFunc(&_dispatch_set_qos_class_floor, &_dispatch_set_qos_class_floorErr, frameworkHandle, "dispatch_set_qos_class_floor", "10.14")
	registerFunc(&_dispatch_set_target_queue, &_dispatch_set_target_queueErr, frameworkHandle, "dispatch_set_target_queue", "10.6")
	registerFunc(&_dispatch_source_cancel, &_dispatch_source_cancelErr, frameworkHandle, "dispatch_source_cancel", "10.6")
	registerFunc(&_dispatch_source_create, &_dispatch_source_createErr, frameworkHandle, "dispatch_source_create", "10.6")
	registerFunc(&_dispatch_source_get_data, &_dispatch_source_get_dataErr, frameworkHandle, "dispatch_source_get_data", "10.6")
	registerFunc(&_dispatch_source_get_handle, &_dispatch_source_get_handleErr, frameworkHandle, "dispatch_source_get_handle", "10.6")
	registerFunc(&_dispatch_source_get_mask, &_dispatch_source_get_maskErr, frameworkHandle, "dispatch_source_get_mask", "10.6")
	registerFunc(&_dispatch_source_merge_data, &_dispatch_source_merge_dataErr, frameworkHandle, "dispatch_source_merge_data", "10.6")
	registerFunc(&_dispatch_source_set_cancel_handler, &_dispatch_source_set_cancel_handlerErr, frameworkHandle, "dispatch_source_set_cancel_handler", "10.6")
	registerFunc(&_dispatch_source_set_cancel_handler_f, &_dispatch_source_set_cancel_handler_fErr, frameworkHandle, "dispatch_source_set_cancel_handler_f", "10.6")
	registerFunc(&_dispatch_source_set_event_handler, &_dispatch_source_set_event_handlerErr, frameworkHandle, "dispatch_source_set_event_handler", "10.6")
	registerFunc(&_dispatch_source_set_event_handler_f, &_dispatch_source_set_event_handler_fErr, frameworkHandle, "dispatch_source_set_event_handler_f", "10.6")
	registerFunc(&_dispatch_source_set_registration_handler, &_dispatch_source_set_registration_handlerErr, frameworkHandle, "dispatch_source_set_registration_handler", "10.7")
	registerFunc(&_dispatch_source_set_registration_handler_f, &_dispatch_source_set_registration_handler_fErr, frameworkHandle, "dispatch_source_set_registration_handler_f", "10.7")
	registerFunc(&_dispatch_source_set_timer, &_dispatch_source_set_timerErr, frameworkHandle, "dispatch_source_set_timer", "10.6")
	registerFunc(&_dispatch_source_testcancel, &_dispatch_source_testcancelErr, frameworkHandle, "dispatch_source_testcancel", "10.6")
	registerFunc(&_dispatch_suspend, &_dispatch_suspendErr, frameworkHandle, "dispatch_suspend", "10.6")
	registerFunc(&_dispatch_sync, &_dispatch_syncErr, frameworkHandle, "dispatch_sync", "10.6")
	registerFunc(&_dispatch_sync_f, &_dispatch_sync_fErr, frameworkHandle, "dispatch_sync_f", "10.6")
	registerFunc(&_dispatch_time, &_dispatch_timeErr, frameworkHandle, "dispatch_time", "10.6")
	registerFunc(&_dispatch_walltime, &_dispatch_walltimeErr, frameworkHandle, "dispatch_walltime", "10.6")
	registerFunc(&_dispatch_workloop_create, &_dispatch_workloop_createErr, frameworkHandle, "dispatch_workloop_create", "10.14")
	registerFunc(&_dispatch_workloop_create_inactive, &_dispatch_workloop_create_inactiveErr, frameworkHandle, "dispatch_workloop_create_inactive", "10.14")
	registerFunc(&_dispatch_workloop_set_autorelease_frequency, &_dispatch_workloop_set_autorelease_frequencyErr, frameworkHandle, "dispatch_workloop_set_autorelease_frequency", "10.14")
	registerFunc(&_dispatch_workloop_set_os_workgroup, &_dispatch_workloop_set_os_workgroupErr, frameworkHandle, "dispatch_workloop_set_os_workgroup", "11.0")
	registerFunc(&_dispatch_write, &_dispatch_writeErr, frameworkHandle, "dispatch_write", "10.7")
}

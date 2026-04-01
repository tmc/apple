// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/ebitengine/purego"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Dispatch: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: Dispatch: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Dispatch: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _dispatch_activate func(object uintptr)

var _dispatch_after func(when uint64, queue uintptr, block unsafe.Pointer)

var _dispatch_after_f func(when uint64, queue uintptr, context unsafe.Pointer, work unsafe.Pointer)

var _dispatch_allow_send_signals func(preserve_signum int) int

var _dispatch_apply func(iterations uintptr, queue uintptr, work unsafe.Pointer)

var _dispatch_apply_f func(iterations uintptr, queue uintptr, context unsafe.Pointer, work unsafe.Pointer)

var _dispatch_assert_queue func(queue uintptr)

var _dispatch_assert_queue_barrier func(queue uintptr)

var _dispatch_assert_queue_not func(queue uintptr)

var _dispatch_async func(queue uintptr, block unsafe.Pointer)

var _dispatch_async_and_wait func(queue uintptr, block unsafe.Pointer)

var _dispatch_async_and_wait_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)

var _dispatch_async_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)

var _dispatch_barrier_async func(queue uintptr, block unsafe.Pointer)

var _dispatch_barrier_async_and_wait func(queue uintptr, block unsafe.Pointer)

var _dispatch_barrier_async_and_wait_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)

var _dispatch_barrier_async_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)

var _dispatch_barrier_sync func(queue uintptr, block unsafe.Pointer)

var _dispatch_barrier_sync_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)

var _dispatch_block_cancel func(block unsafe.Pointer)

var _dispatch_block_create func(flags uintptr, block unsafe.Pointer) unsafe.Pointer

var _dispatch_block_create_with_qos_class func(flags uintptr, qos_class uint32, relative_priority int, block unsafe.Pointer) unsafe.Pointer

var _dispatch_block_notify func(block unsafe.Pointer, queue uintptr, notification_block unsafe.Pointer)

var _dispatch_block_perform func(flags uintptr, block unsafe.Pointer)

var _dispatch_block_testcancel func(block unsafe.Pointer) int

var _dispatch_block_wait func(block unsafe.Pointer, timeout uint64) int

var _dispatch_data_apply func(data uintptr, applier unsafe.Pointer) bool

var _dispatch_data_copy_region func(data uintptr, location uintptr, offset_ptr *uintptr) uintptr

var _dispatch_data_create func(buffer unsafe.Pointer, size uintptr, queue uintptr, destructor unsafe.Pointer) uintptr

var _dispatch_data_create_concat func(data1 uintptr, data2 uintptr) uintptr

var _dispatch_data_create_map func(data uintptr, buffer_ptr *unsafe.Pointer, size_ptr *uintptr) uintptr

var _dispatch_data_create_subrange func(data uintptr, offset uintptr, length uintptr) uintptr

var _dispatch_data_get_size func(data uintptr) uintptr

var _dispatch_debug func(object uintptr, message string)

var _dispatch_get_context func(object uintptr) uintptr

var _dispatch_get_global_queue func(identifier int, flags uintptr) uintptr

var _dispatch_get_specific func(key unsafe.Pointer) unsafe.Pointer

var _dispatch_group_async func(group uintptr, queue uintptr, block unsafe.Pointer)

var _dispatch_group_async_f func(group uintptr, queue uintptr, context unsafe.Pointer, work unsafe.Pointer)

var _dispatch_group_create func() uintptr

var _dispatch_group_enter func(group uintptr)

var _dispatch_group_leave func(group uintptr)

var _dispatch_group_notify func(group uintptr, queue uintptr, block unsafe.Pointer)

var _dispatch_group_notify_f func(group uintptr, queue uintptr, context unsafe.Pointer, work unsafe.Pointer)

var _dispatch_group_wait func(group uintptr, timeout uint64) int

var _dispatch_io_barrier func(channel uintptr, barrier unsafe.Pointer)

var _dispatch_io_close func(channel uintptr, flags uint)

var _dispatch_io_create func(type_ uint, fd int32, queue uintptr, cleanup_handler uintptr) uintptr

var _dispatch_io_create_with_io func(type_ uint, io uintptr, queue uintptr, cleanup_handler uintptr) uintptr

var _dispatch_io_create_with_path func(type_ uint, path string, oflag int, mode uint16, queue uintptr, cleanup_handler uintptr) uintptr

var _dispatch_io_get_descriptor func(channel uintptr) int32

var _dispatch_io_read func(channel uintptr, offset int64, length uintptr, queue uintptr, io_handler unsafe.Pointer)

var _dispatch_io_set_high_water func(channel uintptr, high_water uintptr)

var _dispatch_io_set_interval func(channel uintptr, interval uint64, flags uint)

var _dispatch_io_set_low_water func(channel uintptr, low_water uintptr)

var _dispatch_io_write func(channel uintptr, offset int64, data uintptr, queue uintptr, io_handler unsafe.Pointer)

var _dispatch_main func()

var _dispatch_once func(predicate *int, block unsafe.Pointer)

var _dispatch_once_f func(predicate *int, context unsafe.Pointer, function unsafe.Pointer)

var _dispatch_queue_attr_make_initially_inactive func(attr unsafe.Pointer) unsafe.Pointer

var _dispatch_queue_attr_make_with_autorelease_frequency func(attr unsafe.Pointer, frequency uintptr) unsafe.Pointer

var _dispatch_queue_attr_make_with_qos_class func(attr unsafe.Pointer, qos_class uint32, relative_priority int) unsafe.Pointer

var _dispatch_queue_create func(label string, attr dispatch_queue_attr_t) uintptr

var _dispatch_queue_create_with_target func(label string, attr unsafe.Pointer, target uintptr) uintptr

var _dispatch_queue_get_label func(queue uintptr) *byte

var _dispatch_queue_get_qos_class func(queue uintptr, relative_priority_ptr *int32) uint32

var _dispatch_queue_get_specific func(queue uintptr, key unsafe.Pointer) unsafe.Pointer

var _dispatch_queue_set_specific func(queue uintptr, key unsafe.Pointer, context unsafe.Pointer, destructor unsafe.Pointer)

var _dispatch_read func(fd int32, length uintptr, queue uintptr, io_handler unsafe.Pointer)

var _dispatch_release func(object uintptr)

var _dispatch_resume func(object uintptr)

var _dispatch_retain func(object uintptr)

var _dispatch_semaphore_create func(value int) uintptr

var _dispatch_semaphore_signal func(dsema uintptr) int

var _dispatch_semaphore_wait func(dsema uintptr, timeout uint64) int

var _dispatch_set_context func(object uintptr, context uintptr)

var _dispatch_set_finalizer_f func(object uintptr, finalizer unsafe.Pointer)

var _dispatch_set_qos_class_floor func(object uintptr, qos_class uint32, relative_priority int)

var _dispatch_set_target_queue func(object uintptr, queue uintptr)

var _dispatch_source_cancel func(source uintptr)

var _dispatch_source_create func(type_ Dispatch_source_type_t, handle uintptr, mask uintptr, queue uintptr) uintptr

var _dispatch_source_get_data func(source uintptr) uintptr

var _dispatch_source_get_handle func(source uintptr) uintptr

var _dispatch_source_get_mask func(source uintptr) uintptr

var _dispatch_source_merge_data func(source uintptr, value uintptr)

var _dispatch_source_set_cancel_handler func(source uintptr, handler unsafe.Pointer)

var _dispatch_source_set_cancel_handler_f func(source uintptr, handler unsafe.Pointer)

var _dispatch_source_set_event_handler func(source uintptr, handler unsafe.Pointer)

var _dispatch_source_set_event_handler_f func(source uintptr, handler unsafe.Pointer)

var _dispatch_source_set_registration_handler func(source uintptr, handler unsafe.Pointer)

var _dispatch_source_set_registration_handler_f func(source uintptr, handler unsafe.Pointer)

var _dispatch_source_set_timer func(source uintptr, start uint64, interval uint64, leeway uint64)

var _dispatch_source_testcancel func(source uintptr) int

var _dispatch_suspend func(object uintptr)

var _dispatch_sync func(queue uintptr, block unsafe.Pointer)

var _dispatch_sync_f func(queue uintptr, context unsafe.Pointer, work unsafe.Pointer)

var _dispatch_time func(when uint64, delta int64) uint64

var _dispatch_walltime func(when *syscall.Timespec, delta int64) uint64

var _dispatch_workloop_create func(label string) uintptr

var _dispatch_workloop_create_inactive func(label string) uintptr

var _dispatch_workloop_set_autorelease_frequency func(workloop uintptr, frequency uintptr)

var _dispatch_workloop_set_os_workgroup func(workloop uintptr, workgroup unsafe.Pointer)

var _dispatch_write func(fd int32, data uintptr, queue uintptr, io_handler unsafe.Pointer)

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_dispatch_activate, frameworkHandle, "dispatch_activate")
	registerFunc(&_dispatch_after, frameworkHandle, "dispatch_after")
	registerFunc(&_dispatch_after_f, frameworkHandle, "dispatch_after_f")
	registerFunc(&_dispatch_allow_send_signals, frameworkHandle, "dispatch_allow_send_signals")
	registerFunc(&_dispatch_apply, frameworkHandle, "dispatch_apply")
	registerFunc(&_dispatch_apply_f, frameworkHandle, "dispatch_apply_f")
	registerFunc(&_dispatch_assert_queue, frameworkHandle, "dispatch_assert_queue")
	registerFunc(&_dispatch_assert_queue_barrier, frameworkHandle, "dispatch_assert_queue_barrier")
	registerFunc(&_dispatch_assert_queue_not, frameworkHandle, "dispatch_assert_queue_not")
	registerFunc(&_dispatch_async, frameworkHandle, "dispatch_async")
	registerFunc(&_dispatch_async_and_wait, frameworkHandle, "dispatch_async_and_wait")
	registerFunc(&_dispatch_async_and_wait_f, frameworkHandle, "dispatch_async_and_wait_f")
	registerFunc(&_dispatch_async_f, frameworkHandle, "dispatch_async_f")
	registerFunc(&_dispatch_barrier_async, frameworkHandle, "dispatch_barrier_async")
	registerFunc(&_dispatch_barrier_async_and_wait, frameworkHandle, "dispatch_barrier_async_and_wait")
	registerFunc(&_dispatch_barrier_async_and_wait_f, frameworkHandle, "dispatch_barrier_async_and_wait_f")
	registerFunc(&_dispatch_barrier_async_f, frameworkHandle, "dispatch_barrier_async_f")
	registerFunc(&_dispatch_barrier_sync, frameworkHandle, "dispatch_barrier_sync")
	registerFunc(&_dispatch_barrier_sync_f, frameworkHandle, "dispatch_barrier_sync_f")
	registerFunc(&_dispatch_block_cancel, frameworkHandle, "dispatch_block_cancel")
	registerFunc(&_dispatch_block_create, frameworkHandle, "dispatch_block_create")
	registerFunc(&_dispatch_block_create_with_qos_class, frameworkHandle, "dispatch_block_create_with_qos_class")
	registerFunc(&_dispatch_block_notify, frameworkHandle, "dispatch_block_notify")
	registerFunc(&_dispatch_block_perform, frameworkHandle, "dispatch_block_perform")
	registerFunc(&_dispatch_block_testcancel, frameworkHandle, "dispatch_block_testcancel")
	registerFunc(&_dispatch_block_wait, frameworkHandle, "dispatch_block_wait")
	registerFunc(&_dispatch_data_apply, frameworkHandle, "dispatch_data_apply")
	registerFunc(&_dispatch_data_copy_region, frameworkHandle, "dispatch_data_copy_region")
	registerFunc(&_dispatch_data_create, frameworkHandle, "dispatch_data_create")
	registerFunc(&_dispatch_data_create_concat, frameworkHandle, "dispatch_data_create_concat")
	registerFunc(&_dispatch_data_create_map, frameworkHandle, "dispatch_data_create_map")
	registerFunc(&_dispatch_data_create_subrange, frameworkHandle, "dispatch_data_create_subrange")
	registerFunc(&_dispatch_data_get_size, frameworkHandle, "dispatch_data_get_size")
	registerFunc(&_dispatch_debug, frameworkHandle, "dispatch_debug")
	registerFunc(&_dispatch_get_context, frameworkHandle, "dispatch_get_context")
	registerFunc(&_dispatch_get_global_queue, frameworkHandle, "dispatch_get_global_queue")
	registerFunc(&_dispatch_get_specific, frameworkHandle, "dispatch_get_specific")
	registerFunc(&_dispatch_group_async, frameworkHandle, "dispatch_group_async")
	registerFunc(&_dispatch_group_async_f, frameworkHandle, "dispatch_group_async_f")
	registerFunc(&_dispatch_group_create, frameworkHandle, "dispatch_group_create")
	registerFunc(&_dispatch_group_enter, frameworkHandle, "dispatch_group_enter")
	registerFunc(&_dispatch_group_leave, frameworkHandle, "dispatch_group_leave")
	registerFunc(&_dispatch_group_notify, frameworkHandle, "dispatch_group_notify")
	registerFunc(&_dispatch_group_notify_f, frameworkHandle, "dispatch_group_notify_f")
	registerFunc(&_dispatch_group_wait, frameworkHandle, "dispatch_group_wait")
	registerFunc(&_dispatch_io_barrier, frameworkHandle, "dispatch_io_barrier")
	registerFunc(&_dispatch_io_close, frameworkHandle, "dispatch_io_close")
	registerFunc(&_dispatch_io_create, frameworkHandle, "dispatch_io_create")
	registerFunc(&_dispatch_io_create_with_io, frameworkHandle, "dispatch_io_create_with_io")
	registerFunc(&_dispatch_io_create_with_path, frameworkHandle, "dispatch_io_create_with_path")
	registerFunc(&_dispatch_io_get_descriptor, frameworkHandle, "dispatch_io_get_descriptor")
	registerFunc(&_dispatch_io_read, frameworkHandle, "dispatch_io_read")
	registerFunc(&_dispatch_io_set_high_water, frameworkHandle, "dispatch_io_set_high_water")
	registerFunc(&_dispatch_io_set_interval, frameworkHandle, "dispatch_io_set_interval")
	registerFunc(&_dispatch_io_set_low_water, frameworkHandle, "dispatch_io_set_low_water")
	registerFunc(&_dispatch_io_write, frameworkHandle, "dispatch_io_write")
	registerFunc(&_dispatch_main, frameworkHandle, "dispatch_main")
	registerFunc(&_dispatch_once, frameworkHandle, "dispatch_once")
	registerFunc(&_dispatch_once_f, frameworkHandle, "dispatch_once_f")
	registerFunc(&_dispatch_queue_attr_make_initially_inactive, frameworkHandle, "dispatch_queue_attr_make_initially_inactive")
	registerFunc(&_dispatch_queue_attr_make_with_autorelease_frequency, frameworkHandle, "dispatch_queue_attr_make_with_autorelease_frequency")
	registerFunc(&_dispatch_queue_attr_make_with_qos_class, frameworkHandle, "dispatch_queue_attr_make_with_qos_class")
	registerFunc(&_dispatch_queue_create, frameworkHandle, "dispatch_queue_create")
	registerFunc(&_dispatch_queue_create_with_target, frameworkHandle, "dispatch_queue_create_with_target")
	registerFunc(&_dispatch_queue_get_label, frameworkHandle, "dispatch_queue_get_label")
	registerFunc(&_dispatch_queue_get_qos_class, frameworkHandle, "dispatch_queue_get_qos_class")
	registerFunc(&_dispatch_queue_get_specific, frameworkHandle, "dispatch_queue_get_specific")
	registerFunc(&_dispatch_queue_set_specific, frameworkHandle, "dispatch_queue_set_specific")
	registerFunc(&_dispatch_read, frameworkHandle, "dispatch_read")
	registerFunc(&_dispatch_release, frameworkHandle, "dispatch_release")
	registerFunc(&_dispatch_resume, frameworkHandle, "dispatch_resume")
	registerFunc(&_dispatch_retain, frameworkHandle, "dispatch_retain")
	registerFunc(&_dispatch_semaphore_create, frameworkHandle, "dispatch_semaphore_create")
	registerFunc(&_dispatch_semaphore_signal, frameworkHandle, "dispatch_semaphore_signal")
	registerFunc(&_dispatch_semaphore_wait, frameworkHandle, "dispatch_semaphore_wait")
	registerFunc(&_dispatch_set_context, frameworkHandle, "dispatch_set_context")
	registerFunc(&_dispatch_set_finalizer_f, frameworkHandle, "dispatch_set_finalizer_f")
	registerFunc(&_dispatch_set_qos_class_floor, frameworkHandle, "dispatch_set_qos_class_floor")
	registerFunc(&_dispatch_set_target_queue, frameworkHandle, "dispatch_set_target_queue")
	registerFunc(&_dispatch_source_cancel, frameworkHandle, "dispatch_source_cancel")
	registerFunc(&_dispatch_source_create, frameworkHandle, "dispatch_source_create")
	registerFunc(&_dispatch_source_get_data, frameworkHandle, "dispatch_source_get_data")
	registerFunc(&_dispatch_source_get_handle, frameworkHandle, "dispatch_source_get_handle")
	registerFunc(&_dispatch_source_get_mask, frameworkHandle, "dispatch_source_get_mask")
	registerFunc(&_dispatch_source_merge_data, frameworkHandle, "dispatch_source_merge_data")
	registerFunc(&_dispatch_source_set_cancel_handler, frameworkHandle, "dispatch_source_set_cancel_handler")
	registerFunc(&_dispatch_source_set_cancel_handler_f, frameworkHandle, "dispatch_source_set_cancel_handler_f")
	registerFunc(&_dispatch_source_set_event_handler, frameworkHandle, "dispatch_source_set_event_handler")
	registerFunc(&_dispatch_source_set_event_handler_f, frameworkHandle, "dispatch_source_set_event_handler_f")
	registerFunc(&_dispatch_source_set_registration_handler, frameworkHandle, "dispatch_source_set_registration_handler")
	registerFunc(&_dispatch_source_set_registration_handler_f, frameworkHandle, "dispatch_source_set_registration_handler_f")
	registerFunc(&_dispatch_source_set_timer, frameworkHandle, "dispatch_source_set_timer")
	registerFunc(&_dispatch_source_testcancel, frameworkHandle, "dispatch_source_testcancel")
	registerFunc(&_dispatch_suspend, frameworkHandle, "dispatch_suspend")
	registerFunc(&_dispatch_sync, frameworkHandle, "dispatch_sync")
	registerFunc(&_dispatch_sync_f, frameworkHandle, "dispatch_sync_f")
	registerFunc(&_dispatch_time, frameworkHandle, "dispatch_time")
	registerFunc(&_dispatch_walltime, frameworkHandle, "dispatch_walltime")
	registerFunc(&_dispatch_workloop_create, frameworkHandle, "dispatch_workloop_create")
	registerFunc(&_dispatch_workloop_create_inactive, frameworkHandle, "dispatch_workloop_create_inactive")
	registerFunc(&_dispatch_workloop_set_autorelease_frequency, frameworkHandle, "dispatch_workloop_set_autorelease_frequency")
	registerFunc(&_dispatch_workloop_set_os_workgroup, frameworkHandle, "dispatch_workloop_set_os_workgroup")
	registerFunc(&_dispatch_write, frameworkHandle, "dispatch_write")
}

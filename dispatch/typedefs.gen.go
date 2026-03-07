// Code generated from Apple documentation. DO NOT EDIT.

package dispatch

import (
	"unsafe"
	"github.com/tmc/apple/objectivec"
)

type Dispatch_block_t = func()







type Dispatch_data_applier_t = func(objectivec.Object, uint32, unsafe.Pointer, uint32) bool







type dispatch_data_t uintptr







type Dispatch_fd_t = int







type Dispatch_function_t = unsafe.Pointer







type dispatch_group_t uintptr







type Dispatch_io_close_flags_t = uint







type Dispatch_io_handler_t = func(bool, objectivec.Object, int)







type Dispatch_io_interval_flags_t = uint







type dispatch_io_t uintptr







type dispatch_io_type_t uintptr







type dispatch_object_t uintptr







type Dispatch_once_t = int







type Dispatch_qos_class_t = uint32







type dispatch_queue_attr_t uintptr







type Dispatch_queue_concurrent_t = dispatch_queue_t







type Dispatch_queue_global_t = dispatch_queue_t







type Dispatch_queue_main_t = dispatch_queue_t







type Dispatch_queue_priority_t = int







type Dispatch_queue_serial_executor_t = dispatch_queue_t







type Dispatch_queue_serial_t = dispatch_queue_t







type dispatch_queue_t uintptr







type dispatch_semaphore_t uintptr







type Dispatch_source_mach_recv_flags_t = uint







type Dispatch_source_mach_send_flags_t = uint







type Dispatch_source_memorypressure_flags_t = uint







type Dispatch_source_proc_flags_t = uint







type dispatch_source_t uintptr







type Dispatch_source_timer_flags_t = uint







type dispatch_source_type_t uintptr







type Dispatch_source_vnode_flags_t = uint







type Dispatch_time_t = uint64







type Dispatch_workloop_t = dispatch_queue_t






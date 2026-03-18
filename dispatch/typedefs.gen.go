// Code generated from Apple documentation. DO NOT EDIT.

package dispatch

import (
	"unsafe"
	"github.com/tmc/apple/objectivec"
)

// Dispatch_block_t is the prototype of blocks submitted to dispatch queues, which take no arguments and have no return value.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_block_t
type Dispatch_block_t = func()

// Dispatch_data_applier_t is a block to invoke for every contiguous memory region in a data object.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_data_applier_t
type Dispatch_data_applier_t = func(objectivec.Object, uint32, unsafe.Pointer, uint32) bool

// Dispatch_data_t is an immutable object representing a contiguous or sparse region of memory.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_data_t
type dispatch_data_t uintptr

// Dispatch_fd_t is a file descriptor used for I/O operations.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_fd_t
type Dispatch_fd_t = int

// Dispatch_function_t is the prototype of functions submitted to dispatch queues.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_function_t
type Dispatch_function_t = unsafe.Pointer

// Dispatch_group_t is a group of block objects submitted to a queue for asynchronous invocation.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_group_t
type dispatch_group_t uintptr

// Dispatch_io_close_flags_t is additional flags to use when closing an I/O channel.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_io_close_flags_t
type Dispatch_io_close_flags_t = uint

// Dispatch_io_handler_t is a handler block used to process operations on a dispatch I/O channel.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_io_handler_t
type Dispatch_io_handler_t = func(bool, objectivec.Object, int)

// Dispatch_io_interval_flags_t is the desired delivery behavior for interval events.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_io_interval_flags_t
type Dispatch_io_interval_flags_t = uint

// Dispatch_io_t is a dispatch I/O channel.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_io_t
type dispatch_io_t uintptr

// Dispatch_io_type_t is the type of a dispatch I/O channel.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_io_type_t
type dispatch_io_type_t uintptr

// Dispatch_object_t is a dispatch object.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_object_t
type dispatch_object_t uintptr

// Dispatch_once_t is a predicate for use with the `dispatch_once` function.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_once_t
type Dispatch_once_t = int

// Dispatch_qos_class_t is quality-of-service classes that specify the priorities for executing tasks.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_qos_class_t
type Dispatch_qos_class_t = uint32

// Dispatch_queue_attr_t is attributes describing the behaviors of a dispatch queue.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_queue_attr_t
type dispatch_queue_attr_t uintptr

// Dispatch_queue_concurrent_t is a dispatch queue that executes tasks concurrently and in any order, respecting any barriers that may be in place.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_queue_concurrent_t
type Dispatch_queue_concurrent_t = dispatch_queue_t

// Dispatch_queue_global_t is a dispatch queue that executes tasks concurrently using threads from the global thread pool.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_queue_global_t
type Dispatch_queue_global_t = dispatch_queue_t

// Dispatch_queue_main_t is a dispatch queue that is bound to the app’s main thread and executes tasks serially on that thread.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_queue_main_t
type Dispatch_queue_main_t = dispatch_queue_t

// Dispatch_queue_priority_t is the execution priority for tasks in a global concurrent queue.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_queue_priority_t
type Dispatch_queue_priority_t = int

// See: https://developer.apple.com/documentation/Dispatch/dispatch_queue_serial_executor_t
type Dispatch_queue_serial_executor_t = dispatch_queue_t

// Dispatch_queue_serial_t is a dispatch queue that executes tasks serially in first-in, first-out (FIFO) order.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_queue_serial_t
type Dispatch_queue_serial_t = dispatch_queue_t

// Dispatch_queue_t is a lightweight object to which your application submits blocks for subsequent execution.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_queue_t
type dispatch_queue_t uintptr

// Dispatch_semaphore_t is a dispatch semaphore object.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_semaphore_t
type dispatch_semaphore_t uintptr

// Dispatch_source_mach_recv_flags_t is mach receive-right flags.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_source_mach_recv_flags_t
type Dispatch_source_mach_recv_flags_t = uint

// Dispatch_source_mach_send_flags_t is mach send-right flags.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_source_mach_send_flags_t
type Dispatch_source_mach_send_flags_t = uint

// Dispatch_source_memorypressure_flags_t is memory pressure events.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_source_memorypressure_flags_t
type Dispatch_source_memorypressure_flags_t = uint

// Dispatch_source_proc_flags_t is events related to a process.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_source_proc_flags_t
type Dispatch_source_proc_flags_t = uint

// Dispatch_source_t is an object that coordinates the processing of specific low-level system events, such as file-system events, timers, and UNIX signals.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_source_t
type dispatch_source_t uintptr

// Dispatch_source_timer_flags_t is flags to use when configuring a timer dispatch source.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_source_timer_flags_t
type Dispatch_source_timer_flags_t = uint

// Dispatch_source_type_t is an identifier for the type of system object being monitored by a dispatch source.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_source_type_t
type dispatch_source_type_t uintptr

// Dispatch_source_vnode_flags_t is events involving a change to a file system object.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_source_vnode_flags_t
type Dispatch_source_vnode_flags_t = uint

// Dispatch_time_t is an abstract representation of time.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_time_t
type Dispatch_time_t = uint64

// Dispatch_workloop_t is a dispatch queue that prioritizes the execution of tasks based on their quality-of-service level.
//
// See: https://developer.apple.com/documentation/Dispatch/dispatch_workloop_t
type Dispatch_workloop_t = dispatch_queue_t


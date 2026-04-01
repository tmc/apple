// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

// Package dispatch provides Go bindings for Apple's Grand Central Dispatch (GCD).
//
// GCD is a low-level API for managing concurrent operations. This package
// exposes Go-idiomatic wrappers (exported, like QueueCreate). Raw generated
// C bindings are package-private and used internally by these wrappers.
//
// # Go Runtime Compatibility
//
// GCD and Go's runtime both manage threads. Key considerations:
//
//   - Use dispatch_async_f (not dispatch_async) for Go function callbacks
//   - Go closures passed to GCD will execute on GCD-managed threads
//   - Use sync.WaitGroup or channels to coordinate with Go code
//   - For main thread work on macOS, use the main dispatch queue
//
// # Basic Usage
//
//	queue := dispatch.GetGlobalQueue(dispatch.QOSDefault)
//	queue.Async(func() {
//	    // Work executes on a GCD thread
//	})
//
// # Groups for Coordination
//
//	group := dispatch.GroupCreate()
//	queue := dispatch.GetGlobalQueue(dispatch.QOSDefault)
//	for i := 0; i < 10; i++ {
//	    group.Async(queue, func() {
//	        // Parallel work
//	    })
//	}
//	group.Wait(dispatch.TimeForever)
//
// # Callback Memory
//
// Each unique dispatch function (Async, Sync, BarrierAsync, After,
// Group.Async, Group.Notify) uses a single permanent purego callback.
// Work closures are passed via a context map, not via new callbacks.
// This means an unlimited number of dispatches can be performed without
// exhausting purego's 2000-callback limit.
//
// # Known Limitations with Apple Frameworks
//
// When passing dispatch queues to Apple frameworks (like Virtualization, AVFoundation),
// some frameworks internally use ObjC dispatch_sync blocks which require cgo coordination.
// This package uses purego (cgo-free), which may cause crashes when:
//
//   - Framework wraps callbacks in ObjC dispatch_sync blocks
//   - Framework schedules work that calls back into Go
//   - Go callbacks execute on framework-managed threads
//
// Workarounds:
//
//   - Use nil/default queue when passing to frameworks (let framework manage its queue)
//   - Use the main dispatch queue for framework interactions
//   - For Virtualization framework specifically, consider using cgo-based wrappers
//
// For Go-to-Go dispatch operations (queue.Async, queue.Sync, groups, semaphores),
// this package works correctly with race detection and GC.
package dispatch

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

// Properly typed functions for dispatch functions that take integer values.
// The generated versions used pointer-typed params that fail checkptr when passed Go pointers.
var (
	// dispatch_queue_global_t Dispatch_get_global_queue(intptr_t identifier, unsigned long flags)
	_dispatch_get_global_queue_typed func(identifier uintptr, flags uintptr) uintptr
	// dispatch_semaphore_t Dispatch_semaphore_create(intptr_t value)
	_dispatch_semaphore_create_typed func(value uintptr) uintptr
	_dispatch_main                   func()
	_dispatch_queue_create           func(label *byte, attr dispatch_queue_attr_t) uintptr
	_dispatch_async_f                func(queue uintptr, context unsafe.Pointer, work Dispatch_function_t)
	_dispatch_sync_f                 func(queue uintptr, context unsafe.Pointer, work Dispatch_function_t)
	_dispatch_barrier_async_f        func(queue uintptr, context unsafe.Pointer, work Dispatch_function_t)
	_dispatch_queue_get_label        func(queue uintptr) *byte
	_dispatch_after_f                func(when uint64, queue uintptr, context unsafe.Pointer, work Dispatch_function_t)
	_dispatch_time                   func(when uint64, delta int64) uint64
	_dispatch_group_create           func() uintptr
	_dispatch_group_async_f          func(group uintptr, queue uintptr, context unsafe.Pointer, work Dispatch_function_t)
	_dispatch_group_wait             func(group uintptr, timeout uint64) int
	_dispatch_group_notify_f         func(group uintptr, queue uintptr, context unsafe.Pointer, work Dispatch_function_t)
	_dispatch_group_enter            func(group uintptr)
	_dispatch_group_leave            func(group uintptr)
	_dispatch_semaphore_signal       func(dsema uintptr) int
	_dispatch_semaphore_wait         func(dsema uintptr, timeout uint64) int

	// dispatch_data functions
	_dispatch_data_create        func(buffer, size, queue, destructor uintptr) uintptr
	_dispatch_data_get_size      func(data uintptr) uint64
	_dispatch_data_create_concat func(data1, data2 uintptr) uintptr
	_dispatch_data_create_map    func(data, bufferPtr, sizePtr uintptr) uintptr
	_dispatch_release            func(obj uintptr)
)

// _dispatch_queue_attr_concurrent is the DISPATCH_QUEUE_CONCURRENT attribute
// loaded via dlsym (it's a global symbol, not a function).
var _dispatch_queue_attr_concurrent dispatch_queue_attr_t

// _dispatch_main_q is the main dispatch queue (Dispatch_get_main_queue()
// is a macro that expands to &_dispatch_main_q).
var _dispatch_main_q uintptr

// _dispatch_data_empty_val is the global empty dispatch data singleton.
var _dispatch_data_empty_val dispatch_data_t

// workMap stores pending work closures keyed by a monotonic ID.
// dispatch_*_f functions receive the ID as the context pointer,
// and the single shared callback looks up and executes the closure.
// This avoids calling purego.NewCallback per dispatch, which would
// exhaust the 2000-callback limit.
var workMap struct {
	mu    sync.Mutex
	items map[uintptr]workItem
}

var workID atomic.Uintptr

type workItem struct {
	fn    func()
	token *uintptr
}

// storeWork saves a closure and returns a context pointer for dispatch_*_f.
func storeWork(fn func()) unsafe.Pointer {
	token := new(uintptr)
	*token = workID.Add(1)
	key := uintptr(unsafe.Pointer(token))
	workMap.mu.Lock()
	workMap.items[key] = workItem{fn: fn, token: token}
	workMap.mu.Unlock()
	return unsafe.Pointer(token)
}

// sharedTrampoline is the single C callback used by all dispatch operations.
// It receives the work ID as its context argument, looks up the closure,
// and executes it.
var sharedTrampoline unsafe.Pointer

func init() {
	workMap.items = make(map[uintptr]workItem)

	sharedTrampoline = unsafe.Pointer(purego.NewCallback(func(ctx uintptr) {
		workMap.mu.Lock()
		item := workMap.items[ctx]
		delete(workMap.items, ctx)
		workMap.mu.Unlock()
		if item.fn != nil {
			item.fn()
		}
	}))

	lib := frameworkHandle
	if lib == 0 {
		for _, path := range frameworkPaths {
			lib, _ = purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
			if lib != 0 {
				break
			}
		}
	}
	if lib == 0 {
		fmt.Fprintf(os.Stderr, "warning: dispatch: failed to load framework from any known path\n")
		_dispatch_get_global_queue_typed = func(identifier uintptr, flags uintptr) uintptr {
			panic("dispatch: dispatch_get_global_queue unavailable")
		}
		_dispatch_semaphore_create_typed = func(value uintptr) uintptr {
			panic("dispatch: dispatch_semaphore_create unavailable")
		}
		_dispatch_main = func() { panic("dispatch: dispatch_main unavailable") }
		_dispatch_queue_create = func(*byte, dispatch_queue_attr_t) uintptr {
			panic("dispatch: dispatch_queue_create unavailable")
		}
		_dispatch_async_f = func(uintptr, unsafe.Pointer, Dispatch_function_t) {
			panic("dispatch: dispatch_async_f unavailable")
		}
		_dispatch_sync_f = func(uintptr, unsafe.Pointer, Dispatch_function_t) {
			panic("dispatch: dispatch_sync_f unavailable")
		}
		_dispatch_barrier_async_f = func(uintptr, unsafe.Pointer, Dispatch_function_t) {
			panic("dispatch: dispatch_barrier_async_f unavailable")
		}
		_dispatch_queue_get_label = func(uintptr) *byte {
			panic("dispatch: dispatch_queue_get_label unavailable")
		}
		_dispatch_after_f = func(uint64, uintptr, unsafe.Pointer, Dispatch_function_t) {
			panic("dispatch: dispatch_after_f unavailable")
		}
		_dispatch_time = func(uint64, int64) uint64 {
			panic("dispatch: dispatch_time unavailable")
		}
		_dispatch_group_create = func() uintptr {
			panic("dispatch: dispatch_group_create unavailable")
		}
		_dispatch_group_async_f = func(uintptr, uintptr, unsafe.Pointer, Dispatch_function_t) {
			panic("dispatch: dispatch_group_async_f unavailable")
		}
		_dispatch_group_wait = func(uintptr, uint64) int {
			panic("dispatch: dispatch_group_wait unavailable")
		}
		_dispatch_group_notify_f = func(uintptr, uintptr, unsafe.Pointer, Dispatch_function_t) {
			panic("dispatch: dispatch_group_notify_f unavailable")
		}
		_dispatch_group_enter = func(uintptr) { panic("dispatch: dispatch_group_enter unavailable") }
		_dispatch_group_leave = func(uintptr) { panic("dispatch: dispatch_group_leave unavailable") }
		_dispatch_semaphore_signal = func(uintptr) int {
			panic("dispatch: dispatch_semaphore_signal unavailable")
		}
		_dispatch_semaphore_wait = func(uintptr, uint64) int {
			panic("dispatch: dispatch_semaphore_wait unavailable")
		}
		return
	}
	purego.RegisterLibFunc(&_dispatch_get_global_queue_typed, lib, "dispatch_get_global_queue")
	purego.RegisterLibFunc(&_dispatch_semaphore_create_typed, lib, "dispatch_semaphore_create")
	purego.RegisterLibFunc(&_dispatch_main, lib, "dispatch_main")
	purego.RegisterLibFunc(&_dispatch_queue_create, lib, "dispatch_queue_create")
	purego.RegisterLibFunc(&_dispatch_async_f, lib, "dispatch_async_f")
	purego.RegisterLibFunc(&_dispatch_sync_f, lib, "dispatch_sync_f")
	purego.RegisterLibFunc(&_dispatch_barrier_async_f, lib, "dispatch_barrier_async_f")
	purego.RegisterLibFunc(&_dispatch_queue_get_label, lib, "dispatch_queue_get_label")
	purego.RegisterLibFunc(&_dispatch_after_f, lib, "dispatch_after_f")
	purego.RegisterLibFunc(&_dispatch_time, lib, "dispatch_time")
	purego.RegisterLibFunc(&_dispatch_group_create, lib, "dispatch_group_create")
	purego.RegisterLibFunc(&_dispatch_group_async_f, lib, "dispatch_group_async_f")
	purego.RegisterLibFunc(&_dispatch_group_wait, lib, "dispatch_group_wait")
	purego.RegisterLibFunc(&_dispatch_group_notify_f, lib, "dispatch_group_notify_f")
	purego.RegisterLibFunc(&_dispatch_group_enter, lib, "dispatch_group_enter")
	purego.RegisterLibFunc(&_dispatch_group_leave, lib, "dispatch_group_leave")
	purego.RegisterLibFunc(&_dispatch_semaphore_signal, lib, "dispatch_semaphore_signal")
	purego.RegisterLibFunc(&_dispatch_semaphore_wait, lib, "dispatch_semaphore_wait")
	purego.RegisterLibFunc(&_dispatch_data_create, lib, "dispatch_data_create")
	purego.RegisterLibFunc(&_dispatch_data_get_size, lib, "dispatch_data_get_size")
	purego.RegisterLibFunc(&_dispatch_data_create_concat, lib, "dispatch_data_create_concat")
	purego.RegisterLibFunc(&_dispatch_data_create_map, lib, "dispatch_data_create_map")
	purego.RegisterLibFunc(&_dispatch_release, lib, "dispatch_release")

	if sym, err := purego.Dlsym(lib, "_dispatch_queue_attr_concurrent"); err == nil {
		_dispatch_queue_attr_concurrent = dispatch_queue_attr_t(sym)
	}
	if sym, err := purego.Dlsym(lib, "_dispatch_main_q"); err == nil {
		_dispatch_main_q = sym
	}
	if sym, err := purego.Dlsym(lib, "_dispatch_data_empty"); err == nil {
		_dispatch_data_empty_val = dispatch_data_t(sym)
	}
}

// QoS (Quality of Service) classes for dispatch queues.
// Higher QoS classes get more system resources.
type QOS uint32

const (
	// QOSUserInteractive is for work that is interacting with the user,
	// such as operating on the main thread, refreshing the UI, or performing animations.
	QOSUserInteractive QOS = 0x21

	// QOSUserInitiated is for work that the user has initiated and is waiting for.
	QOSUserInitiated QOS = 0x19

	// QOSDefault is the default QoS class used when no explicit class is specified.
	QOSDefault QOS = 0x15

	// QOSUtility is for work that the user is unlikely to be immediately waiting for.
	QOSUtility QOS = 0x11

	// QOSBackground is for work that is not visible to the user, such as indexing or syncing.
	QOSBackground QOS = 0x09
)

// Time constants for dispatch operations.
const (
	// TimeNow represents the current time (no delay).
	TimeNow Dispatch_time_t = 0

	// TimeForever represents an infinite timeout.
	TimeForever Dispatch_time_t = ^Dispatch_time_t(0)
)

// Queue wraps a dispatch queue with Go-friendly methods.
type Queue struct {
	queue dispatch_queue_t
}

// ID returns the underlying queue handle as an objc.ID for framework compatibility.
func (q Queue) ID() objc.ID {
	return objc.ID(q.queue)
}

// GetID implements the objectivec.IObject interface.
func (q Queue) GetID() objc.ID {
	return q.ID()
}

// Group wraps a dispatch group for coordinating multiple async operations.
type Group struct {
	group dispatch_group_t
}

// Semaphore wraps a dispatch semaphore for resource counting.
type Semaphore struct {
	sema dispatch_semaphore_t
}

// QueueCreate creates a new serial dispatch queue with the given label.
// Serial queues execute tasks one at a time in FIFO order.
func QueueCreate(label string) Queue {
	b := append([]byte(label), 0)
	return Queue{queue: dispatch_queue_t(_dispatch_queue_create(&b[0], 0))}
}

// QueueCreateConcurrent creates a new concurrent dispatch queue with the given label.
// Concurrent queues execute tasks in parallel, subject to system resources.
func QueueCreateConcurrent(label string) Queue {
	b := append([]byte(label), 0)
	return Queue{queue: dispatch_queue_t(_dispatch_queue_create(&b[0], _dispatch_queue_attr_concurrent))}
}

// GetGlobalQueue returns a global concurrent queue with the specified QoS class.
// Global queues are shared system resources and should not be released.
func GetGlobalQueue(qos QOS) Queue {
	return Queue{queue: dispatch_queue_t(_dispatch_get_global_queue_typed(uintptr(qos), 0))}
}

// MainQueue returns the main dispatch queue.
// Work submitted to this queue executes on the application's main thread.
func MainQueue() Queue {
	return Queue{queue: dispatch_queue_t(_dispatch_main_q)}
}

// Main enters the main dispatch event loop and does not return.
func Main() {
	_dispatch_main()
}

// Async submits a function for asynchronous execution on the queue.
// The function will execute on a GCD-managed thread.
func (q Queue) Async(work func()) {
	_dispatch_async_f(uintptr(q.queue), storeWork(work), Dispatch_function_t(sharedTrampoline))
}

// Sync submits a function for synchronous execution on the queue.
// This blocks until the work completes.
//
// Warning: Do not call Sync on a serial queue from within that same queue,
// as this will deadlock.
func (q Queue) Sync(work func()) {
	_dispatch_sync_f(uintptr(q.queue), storeWork(work), Dispatch_function_t(sharedTrampoline))
}

// BarrierAsync submits a barrier function for asynchronous execution.
// Only meaningful on concurrent queues - waits for all prior tasks
// to complete, executes exclusively, then resumes concurrent execution.
func (q Queue) BarrierAsync(work func()) {
	_dispatch_barrier_async_f(uintptr(q.queue), storeWork(work), Dispatch_function_t(sharedTrampoline))
}

// Label returns the label assigned to the queue at creation time.
func (q Queue) Label() string {
	ptr := _dispatch_queue_get_label(uintptr(q.queue))
	if ptr == nil {
		return ""
	}
	return objc.GoString(ptr)
}

// Handle returns the underlying dispatch queue handle for advanced usage.
func (q Queue) Handle() uintptr {
	return uintptr(q.queue)
}

// QueueFromHandle creates a Queue from a raw handle.
func QueueFromHandle(handle uintptr) Queue {
	return Queue{queue: dispatch_queue_t(handle)}
}

// After schedules work to execute after a delay.
func After(when Dispatch_time_t, queue Queue, work func()) {
	_dispatch_after_f(uint64(when), uintptr(queue.queue), storeWork(work), Dispatch_function_t(sharedTrampoline))
}

// TimeFromNow returns a dispatch time representing now + the given nanoseconds.
func TimeFromNow(nanoseconds int64) Dispatch_time_t {
	return Dispatch_time_t(_dispatch_time(uint64(TimeNow), nanoseconds))
}

// GroupCreate creates a new dispatch group for coordinating async operations.
func GroupCreate() Group {
	return Group{group: dispatch_group_t(_dispatch_group_create())}
}

// Async submits work to a queue as part of the group.
func (g Group) Async(queue Queue, work func()) {
	_dispatch_group_async_f(uintptr(g.group), uintptr(queue.queue), storeWork(work), Dispatch_function_t(sharedTrampoline))
}

// Wait blocks until all work in the group completes or timeout expires.
// Returns true if all work completed, false if timeout expired.
func (g Group) Wait(timeout Dispatch_time_t) bool {
	result := _dispatch_group_wait(uintptr(g.group), uint64(timeout))
	return result == 0
}

// Notify schedules work to execute when all current group work completes.
func (g Group) Notify(queue Queue, work func()) {
	_dispatch_group_notify_f(uintptr(g.group), uintptr(queue.queue), storeWork(work), Dispatch_function_t(sharedTrampoline))
}

// Enter manually increments the group's task count.
// Each Enter must be balanced with a Leave.
func (g Group) Enter() {
	_dispatch_group_enter(uintptr(g.group))
}

// Leave manually decrements the group's task count.
func (g Group) Leave() {
	_dispatch_group_leave(uintptr(g.group))
}

// Handle returns the underlying dispatch group handle for advanced usage.
func (g Group) Handle() uintptr {
	return uintptr(g.group)
}

// SemaphoreCreate creates a counting semaphore with the given initial value.
// Use value=0 for signaling, value>0 for resource counting.
func SemaphoreCreate(value int) Semaphore {
	return Semaphore{sema: dispatch_semaphore_t(_dispatch_semaphore_create_typed(uintptr(value)))}
}

// Signal increments the semaphore count.
// If threads are waiting, one will be woken.
// Returns true if a thread was woken.
func (s Semaphore) Signal() bool {
	result := _dispatch_semaphore_signal(uintptr(s.sema))
	return result != 0
}

// Wait decrements the semaphore count, blocking if count is zero.
// Returns true if the semaphore was acquired, false on timeout.
func (s Semaphore) Wait(timeout Dispatch_time_t) bool {
	result := _dispatch_semaphore_wait(uintptr(s.sema), uint64(timeout))
	return result == 0
}

// Handle returns the underlying dispatch semaphore handle for advanced usage.
func (s Semaphore) Handle() uintptr {
	return uintptr(s.sema)
}

// SemaphoreFromHandle creates a Semaphore from a raw handle.
func SemaphoreFromHandle(handle uintptr) Semaphore {
	return Semaphore{sema: dispatch_semaphore_t(handle)}
}

// Data wraps a dispatch data object.
type Data struct {
	data dispatch_data_t
}

// DataFromHandle creates a Data from a raw handle.
func DataFromHandle(handle uintptr) Data {
	return Data{data: dispatch_data_t(handle)}
}

// Handle returns the underlying dispatch data handle.
func (d Data) Handle() uintptr {
	return uintptr(d.data)
}

// DataCreate creates a dispatch data object from a Go byte slice.
// The slice contents are copied; the caller may freely modify or
// discard buf after this call returns.
func DataCreate(buf []byte) Data {
	if len(buf) == 0 {
		return DataEmpty()
	}
	var pinner runtime.Pinner
	pinner.Pin(&buf[0])
	defer pinner.Unpin()
	ptr := unsafe.Pointer(&buf[0])
	handle := _dispatch_data_create(uintptr(ptr), uintptr(len(buf)), 0, 0)
	if handle == 0 {
		return DataEmpty()
	}
	return Data{data: dispatch_data_t(handle)}
}

// DataEmpty returns the global empty dispatch data object.
func DataEmpty() Data {
	return Data{data: _dispatch_data_empty_val}
}

// DataGetSize returns the number of bytes in the dispatch data object.
func DataGetSize(d Data) int {
	return int(_dispatch_data_get_size(uintptr(d.data)))
}

// DataCreateConcat creates a new dispatch data object by concatenating two
// existing data objects.
func DataCreateConcat(a, b Data) Data {
	handle := _dispatch_data_create_concat(uintptr(a.data), uintptr(b.data))
	if handle == 0 {
		return DataEmpty()
	}
	return Data{data: dispatch_data_t(handle)}
}

// DataToBytes extracts the contents of a dispatch data object into a Go byte
// slice. The returned slice is a copy; the caller owns it.
func DataToBytes(d Data) []byte {
	if d.Handle() == 0 {
		return nil
	}
	var ptr uintptr
	var size uint64
	mapped := _dispatch_data_create_map(uintptr(d.data),
		uintptr(unsafe.Pointer(&ptr)),
		uintptr(unsafe.Pointer(&size)))
	defer _dispatch_release(mapped)
	if size == 0 {
		return nil
	}
	out := make([]byte, size)
	copy(out, unsafe.Slice((*byte)(unsafe.Pointer(ptr)), size))
	return out
}

// Len returns the number of bytes in the data object.
func (d Data) Len() int {
	return DataGetSize(d)
}

// Bytes extracts the contents into a Go byte slice. The returned slice is a
// copy; the caller owns it.
func (d Data) Bytes() []byte {
	return DataToBytes(d)
}

// IO wraps a dispatch I/O channel.
type IO struct {
	io dispatch_io_t
}

// IOFromHandle creates an IO from a raw handle.
func IOFromHandle(handle uintptr) IO {
	return IO{io: dispatch_io_t(handle)}
}

// Handle returns the underlying dispatch I/O handle.
func (i IO) Handle() uintptr {
	return uintptr(i.io)
}

// Source wraps a dispatch event source.
type Source struct {
	source dispatch_source_t
}

// SourceFromHandle creates a Source from a raw handle.
func SourceFromHandle(handle uintptr) Source {
	return Source{source: dispatch_source_t(handle)}
}

// Handle returns the underlying dispatch source handle.
func (s Source) Handle() uintptr {
	return uintptr(s.source)
}

// GroupFromHandle creates a Group from a raw handle.
func GroupFromHandle(handle uintptr) Group {
	return Group{group: dispatch_group_t(handle)}
}

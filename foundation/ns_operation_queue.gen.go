// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OperationQueue] class.
var (
	_OperationQueueClass     OperationQueueClass
	_OperationQueueClassOnce sync.Once
)

func getOperationQueueClass() OperationQueueClass {
	_OperationQueueClassOnce.Do(func() {
		_OperationQueueClass = OperationQueueClass{class: objc.GetClass("NSOperationQueue")}
	})
	return _OperationQueueClass
}

// GetOperationQueueClass returns the class object for NSOperationQueue.
func GetOperationQueueClass() OperationQueueClass {
	return getOperationQueueClass()
}

type OperationQueueClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (oc OperationQueueClass) Alloc() OperationQueue {
	rv := objc.Send[OperationQueue](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// A queue that regulates the execution of operations.
//
// # Overview
// 
// An operation queue invokes its queued [NSOperation] objects based on their
// priority and readiness. After you add an operation to a queue, it remains
// in the queue until the operation finishes its task. You can’t directly
// remove an operation from a queue after you add it.
// 
// For more information about using operation queues, see the [Concurrency
// Programming Guide].
// 
// # Determine the Execution Order
// 
// An operation queue organizes and invokes its operations according to their
// readiness, priority level, and interoperation dependencies. If all of the
// queued operations have the same [QueuePriority] and the [Ready] property
// returns [true], the queue invokes them in the order you added them.
// Otherwise, the operation queue always invokes the operation with the
// highest priority relative to the other ready operations.
// 
// However, don’t rely on queue semantics to ensure a specific execution
// order of operations because changes in the readiness of an operation can
// change the resulting execution order. Interoperation dependencies provide
// an absolute execution order for operations, even if those operations are
// located in different operation queues. An operation object isn’t ready to
// run until all of its dependent operations have finished running.
// 
// For details on how to set priority levels and dependencies, see Managing
// Dependencies in [NSOperation].
// 
// # Respond to Operation Cancelation
// 
// Finishing its task doesn’t necessarily mean that the operation performed
// that task to completion; an operation can also be canceled. Canceling an
// operation object leaves the object in the queue but notifies the object
// that it should stop its task as quickly as possible. For currently
// executing operations, this means that the operation object’s work code
// must check the cancellation state, stop what it is doing, and mark itself
// as finished. For operations that are queued but not yet executing, the
// queue must still call the operation object’s [Start] method so that it
// can processes the cancellation event and mark itself as finished.
// 
// For more information about operation cancellation, see [NSOperation] in
// [NSOperation].
// 
// # Observe Operations Using Key-Value Observing
// 
// The [NSOperationQueue] class is key-value coding (KVC) and key-value
// observing (KVO) compliant. You can observe these properties to control
// other parts of your application. To observe the properties, use the
// following key paths:
// 
// - [Operations] — Read-only - [OperationCount] — Read-only -
// [MaxConcurrentOperationCount] — Readable and writable - [Suspended] —
// Readable and writable - [Name] — Readable and writable
// 
// Although you can attach observers to these properties, don’t use Cocoa
// bindings to bind these properties to elements of your application’s user
// interface. Code associated with your user interface typically must run only
// in your app’s main thread. However, KVO notifications associated with an
// operation queue may occur in any thread.
// 
// For more information about KVO and how to attach observers to an object,
// see the [Key-Value Observing Programming Guide].
// 
// # Plan for Thread Safety
// 
// You can safely use a single [NSOperationQueue] object from multiple threads
// without creating additional locks to synchronize access to that object.
// 
// Operation queues use the [Dispatch] framework to initiate the execution of
// their operations. As a result, queues always invoke operations on a
// separate thread, regardless of whether the operation is synchronous or
// asynchronous.
//
// [Concurrency Programming Guide]: https://developer.apple.com/library/archive/documentation/General/Conceptual/ConcurrencyProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40008091
// [Dispatch]: https://developer.apple.com/documentation/Dispatch
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Managing Operations in the Queue
//
//   - [OperationQueue.AddOperation]: Adds the specified operation to the receiver.
//   - [OperationQueue.AddOperationsWaitUntilFinished]: Adds the specified operations to the queue.
//   - [OperationQueue.AddOperationWithBlock]: Wraps the specified block in an operation and adds it to the receiver.
//   - [OperationQueue.AddBarrierBlock]: Invokes a block when the queue finishes all enqueued operations, and prevents subsequent operations from starting until the block has completed.
//   - [OperationQueue.CancelAllOperations]: Cancels all queued and executing operations.
//   - [OperationQueue.WaitUntilAllOperationsAreFinished]: Blocks the current thread until all the receiver’s queued and executing operations finish executing.
//   - [OperationQueue.Operations]: The operations currently in the queue.
//   - [OperationQueue.OperationCount]: The number of operations currently in the queue.
//
// # Managing the Execution of Operations
//
//   - [OperationQueue.QualityOfService]: The default service level to apply to operations that the queue invokes.
//   - [OperationQueue.SetQualityOfService]
//   - [OperationQueue.MaxConcurrentOperationCount]: The maximum number of queued operations that can run at the same time.
//   - [OperationQueue.SetMaxConcurrentOperationCount]
//
// # Suspending Execution
//
//   - [OperationQueue.Suspended]: A Boolean value indicating whether the queue is actively scheduling operations for execution.
//   - [OperationQueue.SetSuspended]
//
// # Configuring the Queue
//
//   - [OperationQueue.Name]: The name of the operation queue.
//   - [OperationQueue.SetName]
//   - [OperationQueue.UnderlyingQueue]: The dispatch queue that the operation queue uses to invoke operations.
//   - [OperationQueue.SetUnderlyingQueue]
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue
type OperationQueue struct {
	objectivec.Object
}

// OperationQueueFromID constructs a [OperationQueue] from an objc.ID.
//
// A queue that regulates the execution of operations.
func OperationQueueFromID(id objc.ID) OperationQueue {
	return NSOperationQueue{objectivec.Object{ID: id}}
}

// NSOperationQueueFromID is an alias for [OperationQueueFromID] for cross-framework compatibility.
func NSOperationQueueFromID(id objc.ID) OperationQueue { return OperationQueueFromID(id) }
// NOTE: OperationQueue adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OperationQueue] class.
//
// # Managing Operations in the Queue
//
//   - [IOperationQueue.AddOperation]: Adds the specified operation to the receiver.
//   - [IOperationQueue.AddOperationsWaitUntilFinished]: Adds the specified operations to the queue.
//   - [IOperationQueue.AddOperationWithBlock]: Wraps the specified block in an operation and adds it to the receiver.
//   - [IOperationQueue.AddBarrierBlock]: Invokes a block when the queue finishes all enqueued operations, and prevents subsequent operations from starting until the block has completed.
//   - [IOperationQueue.CancelAllOperations]: Cancels all queued and executing operations.
//   - [IOperationQueue.WaitUntilAllOperationsAreFinished]: Blocks the current thread until all the receiver’s queued and executing operations finish executing.
//   - [IOperationQueue.Operations]: The operations currently in the queue.
//   - [IOperationQueue.OperationCount]: The number of operations currently in the queue.
//
// # Managing the Execution of Operations
//
//   - [IOperationQueue.QualityOfService]: The default service level to apply to operations that the queue invokes.
//   - [IOperationQueue.SetQualityOfService]
//   - [IOperationQueue.MaxConcurrentOperationCount]: The maximum number of queued operations that can run at the same time.
//   - [IOperationQueue.SetMaxConcurrentOperationCount]
//
// # Suspending Execution
//
//   - [IOperationQueue.Suspended]: A Boolean value indicating whether the queue is actively scheduling operations for execution.
//   - [IOperationQueue.SetSuspended]
//
// # Configuring the Queue
//
//   - [IOperationQueue.Name]: The name of the operation queue.
//   - [IOperationQueue.SetName]
//   - [IOperationQueue.UnderlyingQueue]: The dispatch queue that the operation queue uses to invoke operations.
//   - [IOperationQueue.SetUnderlyingQueue]
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue
type IOperationQueue interface {
	objectivec.IObject

	// Topic: Managing Operations in the Queue

	// Adds the specified operation to the receiver.
	AddOperation(op INSOperation)
	// Adds the specified operations to the queue.
	AddOperationsWaitUntilFinished(ops []NSOperation, wait bool)
	// Wraps the specified block in an operation and adds it to the receiver.
	AddOperationWithBlock(block VoidHandler)
	// Invokes a block when the queue finishes all enqueued operations, and prevents subsequent operations from starting until the block has completed.
	AddBarrierBlock(barrier VoidHandler)
	// Cancels all queued and executing operations.
	CancelAllOperations()
	// Blocks the current thread until all the receiver’s queued and executing operations finish executing.
	WaitUntilAllOperationsAreFinished()
	// The operations currently in the queue.
	Operations() []NSOperation
	// The number of operations currently in the queue.
	OperationCount() uint

	// Topic: Managing the Execution of Operations

	// The default service level to apply to operations that the queue invokes.
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	// The maximum number of queued operations that can run at the same time.
	MaxConcurrentOperationCount() int
	SetMaxConcurrentOperationCount(value int)

	// Topic: Suspending Execution

	// A Boolean value indicating whether the queue is actively scheduling operations for execution.
	Suspended() bool
	SetSuspended(value bool)

	// Topic: Configuring the Queue

	// The name of the operation queue.
	Name() string
	SetName(value string)
	// The dispatch queue that the operation queue uses to invoke operations.
	UnderlyingQueue() dispatch.Queue
	SetUnderlyingQueue(value dispatch.Queue)

	// A Boolean value indicating whether the operation can be performed now.
	IsReady() bool
	SetIsReady(value bool)
	// The execution priority of the operation in an operation queue.
	QueuePriority() NSOperationQueuePriority
	SetQueuePriority(value NSOperationQueuePriority)
}

// Init initializes the instance.
func (o OperationQueue) Init() OperationQueue {
	rv := objc.Send[OperationQueue](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OperationQueue) Autorelease() OperationQueue {
	rv := objc.Send[OperationQueue](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOperationQueue creates a new OperationQueue instance.
func NewOperationQueue() OperationQueue {
	class := getOperationQueueClass()
	rv := objc.Send[OperationQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds the specified operation to the receiver.
//
// op: The operation to be added to the queue.
//
// # Discussion
// 
// Once added, the specified operation remains in the queue until it finishes
// executing.
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/addOperation(_:)-64o8a
func (o OperationQueue) AddOperation(op INSOperation) {
	objc.Send[objc.ID](o.ID, objc.Sel("addOperation:"), op)
}
// Adds the specified operations to the queue.
//
// ops: The operations to be added to the queue.
//
// wait: If [true], the current thread is blocked until all of the specified
// operations finish executing. If [false], the operations are added to the
// queue and control returns immediately to the caller.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// An operation object can be in at most one operation queue at a time and
// cannot be added if it is currently executing or finished. This method
// throws an [NSInvalidArgumentException] exception if any of those error
// conditions are true for any of the operations in the `ops` parameter.
// 
// Once added, the specified `operation` remains in the queue until its
// [Finished] method returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/addOperations(_:waitUntilFinished:)
func (o OperationQueue) AddOperationsWaitUntilFinished(ops []NSOperation, wait bool) {
	objc.Send[objc.ID](o.ID, objc.Sel("addOperations:waitUntilFinished:"), objectivec.IObjectSliceToNSArray(ops), wait)
}
// Wraps the specified block in an operation and adds it to the receiver.
//
// block: The block to execute from the operation. The block takes no parameters and
// has no return value.
//
// # Discussion
// 
// This method adds a single block to the receiver by first wrapping it in an
// operation object. You should not attempt to get a reference to the newly
// created operation object or determine its type information.
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/addOperation(_:)-5s294
func (o OperationQueue) AddOperationWithBlock(block VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
	objc.Send[objc.ID](o.ID, objc.Sel("addOperationWithBlock:"), _block0)
}
// Invokes a block when the queue finishes all enqueued operations, and
// prevents subsequent operations from starting until the block has completed.
//
// barrier: The block to invoke after all currently enqueued operations have finished.
// Operations you add after the barrier block don’t start until the block
// has completed.
//
// # Discussion
// 
// This method is similar to [dispatch_barrier_async].
//
// [dispatch_barrier_async]: https://developer.apple.com/documentation/Dispatch/dispatch_barrier_async
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/addBarrierBlock(_:)
func (o OperationQueue) AddBarrierBlock(barrier VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(barrier)
	defer _cleanup0()
	objc.Send[objc.ID](o.ID, objc.Sel("addBarrierBlock:"), _block0)
}
// Cancels all queued and executing operations.
//
// # Discussion
// 
// This method calls the [Cancel] method on all operations currently in the
// queue.
// 
// Canceling the operations does not automatically remove them from the queue
// or stop those that are currently executing. For operations that are queued
// and waiting execution, the queue must still attempt to execute the
// operation before recognizing that it is canceled and moving it to the
// finished state. For operations that are already executing, the operation
// object itself must check for cancellation and stop what it is doing so that
// it can move to the finished state. In both cases, a finished (or canceled)
// operation is still given a chance to execute its completion block before it
// is removed from the queue.
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/cancelAllOperations()
func (o OperationQueue) CancelAllOperations() {
	objc.Send[objc.ID](o.ID, objc.Sel("cancelAllOperations"))
}
// Blocks the current thread until all the receiver’s queued and executing
// operations finish executing.
//
// # Discussion
// 
// When called, this method blocks the current thread and waits for the
// receiver’s current and queued operations to finish executing. While the
// current thread is blocked, the receiver continues to launch already queued
// operations and monitor those that are executing. During this time, the
// current thread cannot add operations to the queue, but other threads may.
// Once all of the pending operations are finished, this method returns.
// 
// If there are no operations in the queue, this method returns immediately.
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/waitUntilAllOperationsAreFinished()
func (o OperationQueue) WaitUntilAllOperationsAreFinished() {
	objc.Send[objc.ID](o.ID, objc.Sel("waitUntilAllOperationsAreFinished"))
}

// The operations currently in the queue.
//
// # Discussion
// 
// The array in this property contains zero or more [NSOperation] objects in
// the order you added them to the queue. This order doesn’t necessarily
// reflect the order in which the queue invokes those operations.
// 
// You can use this property to access the operations queued at any given
// moment. Operations remain queued until they finish their task. Therefore,
// the array may contain operations that are currently running or waiting to
// run. The list may also contain operations that were running when you
// retrieved the array but have subsequently finished.
// 
// You can monitor changes to the value of this property using [Key-value
// observing] (KVO). Configure an observer to monitor the [Operations] key
// path of the operation queue.
//
// [Key-value observing]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/KVO.html#//apple_ref/doc/uid/TP40008195-CH16
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/operations
func (o OperationQueue) Operations() []NSOperation {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("operations"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSOperation {
		return NSOperationFromID(id)
	})
}
// The number of operations currently in the queue.
//
// # Discussion
// 
// Because the number of operations in the queue changes as those operations
// finish executing, the value returned by this property reflects the
// instantaneous number of operations at the time the property was accessed.
// By the time you use the value, the actual number of operations may be
// different. As a result, do not use this value for object enumerations or
// other precise calculations.
// 
// You may monitor changes to the value of this property using [Key-value
// observing]. Configure an observer to monitor the [OperationCount] key path
// of the operation queue.
//
// [Key-value observing]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/KVO.html#//apple_ref/doc/uid/TP40008195-CH16
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/operationCount
func (o OperationQueue) OperationCount() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("operationCount"))
	return rv
}
// The default service level to apply to operations that the queue invokes.
//
// # Discussion
// 
// This property specifies the service level applied to operation objects
// added to the queue. If the operation object has an explicit service level
// set, that value is used instead. The default value of this property depends
// on how you created the queue. For queues you create yourself, the default
// value is [NSOperationQualityOfServiceBackground]. For the queue returned by
// the [MainQueue] method, the default value is
// [NSOperationQualityOfServiceUserInteractive] and cannot be changed.
// 
// Service levels affect the priority with which operation objects are given
// access to system resources such as CPU time, network resources, disk
// resources, and so on. Operations with a higher quality of service level are
// given greater priority over system resources so that they may perform their
// task more quickly. You use service levels to ensure that operations
// responding to explicit user requests are given priority over less critical
// work.
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/qualityOfService
func (o OperationQueue) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](o.ID, objc.Sel("qualityOfService"))
	return QualityOfService(rv)
}
func (o OperationQueue) SetQualityOfService(value QualityOfService) {
	objc.Send[struct{}](o.ID, objc.Sel("setQualityOfService:"), value)
}
// The maximum number of queued operations that can run at the same time.
//
// # Discussion
// 
// The value in this property affects only the operations that the current
// queue has executing at the same time. Other operation queues can also
// execute their maximum number of operations in parallel.
// 
// Reducing the number of concurrent operations does not affect any operations
// that are currently executing. Specifying the value
// [NSOperationQueueDefaultMaxConcurrentOperationCount] (which is recommended)
// causes the system to set the maximum number of operations based on system
// conditions.
// 
// The default value of this property is [defaultMaxConcurrentOperationCount].
// You may monitor changes to the value of this property using [Key-value
// observing]. Configure an observer to monitor the
// [MaxConcurrentOperationCount] key path of the operation queue.
//
// [Key-value observing]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/KVO.html#//apple_ref/doc/uid/TP40008195-CH16
// [defaultMaxConcurrentOperationCount]: https://developer.apple.com/documentation/Foundation/OperationQueue/defaultMaxConcurrentOperationCount
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/maxConcurrentOperationCount
func (o OperationQueue) MaxConcurrentOperationCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("maxConcurrentOperationCount"))
	return rv
}
func (o OperationQueue) SetMaxConcurrentOperationCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaxConcurrentOperationCount:"), value)
}
// An object that represents the total progress of the operations executing in
// the queue.
//
// # Discussion
// 
// By default, [NSOperationQueue] doesn’t report progress until
// [TotalUnitCount] is set. When [TotalUnitCount] is set, the queue begins
// reporting progress. Each operation in the queue contributes one unit of
// completion to the overall progress of the queue for operations that are
// finished by the end of [Main]. Operations that override [Start] and don’t
// invoke super don’t contribute to the queue’s progress.
// 
// To update [TotalUnitCount] in a thread-safe manner, use the
// [AddBarrierBlock] method. This method ensures that the operation queue
// completes the operations in the queue, preventing an inadvertent backward
// jump in progress.
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/progress
func (o OperationQueue) Progress() INSProgress {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("progress"))
	return NSProgressFromID(objc.ID(rv))
}
// A Boolean value indicating whether the queue is actively scheduling
// operations for execution.
//
// # Discussion
// 
// When the value of this property is [false], the queue actively starts
// operations that are in the queue and ready to execute. Setting this
// property to [true] prevents the queue from starting any queued operations,
// but already executing operations continue to execute. You may continue to
// add operations to a queue that is suspended but those operations are not
// scheduled for execution until you change this property to [false].
// 
// Operations are removed from the queue only when they finish executing.
// However, in order to finish executing, an operation must first be started.
// Because a suspended queue does not start any new operations, it does not
// remove any operations (including cancelled operations) that are currently
// queued and not executing.
// 
// You may monitor changes to the value of this property using [Key-value
// observing]. Configure an observer to monitor the [Suspended] key path of
// the operation queue.
// 
// The default value of this property is [false].
//
// [Key-value observing]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/KVO.html#//apple_ref/doc/uid/TP40008195-CH16
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/isSuspended
func (o OperationQueue) Suspended() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isSuspended"))
	return rv
}
func (o OperationQueue) SetSuspended(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setSuspended:"), value)
}
// The name of the operation queue.
//
// # Discussion
// 
// Names provide a way for you to identify your operation queues at run time.
// Tools may also use this name to provide additional context during debugging
// or analysis of your code.
// 
// The default value of this property is a string containing the memory
// address of the operation queue. You may monitor changes to the value of
// this property using [Key-value observing]. Configure an observer to monitor
// the [Name] key path of the operation queue.
//
// [Key-value observing]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/KVO.html#//apple_ref/doc/uid/TP40008195-CH16
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/name
func (o OperationQueue) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
func (o OperationQueue) SetName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setName:"), objc.String(value))
}
// The dispatch queue that the operation queue uses to invoke operations.
//
// # Discussion
// 
// The default value of this property is `nil`. You can set the value of this
// property to an existing dispatch queue to have enqueued operations
// interspersed with blocks submitted to that dispatch queue.
// 
// The value of this property should only be set if there are no operations in
// the queue; setting the value of this property when [OperationCount] is not
// equal to `0` raises an [invalidArgumentException]. The value of this
// property must not be the value returned by [dispatch_get_main_queue]. The
// quality-of-service level set for the underlying dispatch queue overrides
// any value set for the operation queue’s [QualityOfService] property.
//
// [dispatch_get_main_queue]: https://developer.apple.com/documentation/Dispatch/dispatch_get_main_queue
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/underlyingQueue
func (o OperationQueue) UnderlyingQueue() dispatch.Queue {
	rv := objc.Send[uintptr](o.ID, objc.Sel("underlyingQueue"))
	return dispatch.QueueFromHandle(rv)
}
func (o OperationQueue) SetUnderlyingQueue(value dispatch.Queue) {
	objc.Send[struct{}](o.ID, objc.Sel("setUnderlyingQueue:"), uintptr(value.Handle()))
}
// A Boolean value indicating whether the operation can be performed now.
//
// See: https://developer.apple.com/documentation/foundation/operation/isready
func (o OperationQueue) IsReady() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("ready"))
	return rv
}
func (o OperationQueue) SetIsReady(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setReady:"), value)
}
// The execution priority of the operation in an operation queue.
//
// See: https://developer.apple.com/documentation/foundation/operation/queuepriority-swift.property
func (o OperationQueue) QueuePriority() NSOperationQueuePriority {
	rv := objc.Send[NSOperationQueuePriority](o.ID, objc.Sel("queuePriority"))
	return NSOperationQueuePriority(rv)
}
func (o OperationQueue) SetQueuePriority(value NSOperationQueuePriority) {
	objc.Send[struct{}](o.ID, objc.Sel("setQueuePriority:"), value)
}

// Returns the operation queue associated with the main thread.
//
// # Return Value
// 
// The default operation queue bound to the main thread.
// 
// # Discussion
// 
// The returned queue executes one operation at a time on the app’s main
// thread. The execution of operations on the main thread is interleaved with
// the other tasks that must execute on the main thread, such as the servicing
// of events and the updating of an app’s user interface. The queue executes
// those operations in the run loop common modes, as represented by the
// [common] constant. The value of the [UnderlyingQueue] property for the
// queue is the dispatch queue for the main thread; this property cannot be
// set to another value.
//
// [common]: https://developer.apple.com/documentation/Foundation/RunLoop/Mode/common
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/main
func (_OperationQueueClass OperationQueueClass) MainQueue() OperationQueue {
	rv := objc.Send[objc.ID](objc.ID(_OperationQueueClass.class), objc.Sel("mainQueue"))
	return NSOperationQueueFromID(objc.ID(rv))
}
// Returns the operation queue that launched the current operation.
//
// # Return Value
// 
// The operation queue that started the operation or `nil` if the queue could
// not be determined.
// 
// # Discussion
// 
// You can use this method from within a running operation object to get a
// reference to the operation queue that started it. Calling this method from
// outside the context of a running operation typically results in `nil` being
// returned.
//
// See: https://developer.apple.com/documentation/Foundation/OperationQueue/current
func (_OperationQueueClass OperationQueueClass) CurrentQueue() OperationQueue {
	rv := objc.Send[objc.ID](objc.ID(_OperationQueueClass.class), objc.Sel("currentQueue"))
	return NSOperationQueueFromID(objc.ID(rv))
}

// AddOperationWithBlockSync is a synchronous wrapper around [OperationQueue.AddOperationWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (o OperationQueue) AddOperationWithBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	o.AddOperationWithBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// AddBarrierBlockSync is a synchronous wrapper around [OperationQueue.AddBarrierBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (o OperationQueue) AddBarrierBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	o.AddBarrierBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}


// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Operation] class.
var (
	_OperationClass     OperationClass
	_OperationClassOnce sync.Once
)

func getOperationClass() OperationClass {
	_OperationClassOnce.Do(func() {
		_OperationClass = OperationClass{class: objc.GetClass("NSOperation")}
	})
	return _OperationClass
}

// GetOperationClass returns the class object for NSOperation.
func GetOperationClass() OperationClass {
	return getOperationClass()
}

type OperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (oc OperationClass) Alloc() Operation {
	rv := objc.Send[Operation](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that represents the code and data associated with a
// single task.
//
// # Overview
// 
// Because the [NSOperation] class is an abstract class, you do not use it
// directly but instead subclass or use one of the system-defined subclasses
// ([NSInvocationOperation] or [NSBlockOperation]) to perform the actual task.
// Despite being abstract, the base implementation of [NSOperation] does
// include significant logic to coordinate the safe execution of your task.
// The presence of this built-in logic allows you to focus on the actual
// implementation of your task, rather than on the glue code needed to ensure
// it works correctly with other system objects.
// 
// An operation object is a single-shot object—that is, it executes its task
// once and cannot be used to execute it again. You typically execute
// operations by adding them to an operation queue (an instance of the
// [NSOperationQueue] class). An operation queue executes its operations
// either directly, by running them on secondary threads, or indirectly using
// the `libdispatch` library (also known as Grand Central Dispatch). For more
// information about how queues execute operations, see [NSOperationQueue].
// 
// If you do not want to use an operation queue, you can execute an operation
// yourself by calling its [Start] method directly from your code. Executing
// operations manually does put more of a burden on your code, because
// starting an operation that is not in the ready state triggers an exception.
// The [Ready] property reports on the operation’s readiness.
// 
// # Operation Dependencies
// 
// Dependencies are a convenient way to execute operations in a specific
// order. You can add and remove dependencies for an operation using the
// [AddDependency] and [RemoveDependency] methods. By default, an operation
// object that has dependencies is not considered ready until all of its
// dependent operation objects have finished executing. Once the last
// dependent operation finishes, however, the operation object becomes ready
// and able to execute.
// 
// The dependencies supported by [NSOperation] make no distinction about
// whether a dependent operation finished successfully or unsuccessfully. (In
// other words, canceling an operation similarly marks it as finished.) It is
// up to you to determine whether an operation with dependencies should
// proceed in cases where its dependent operations were cancelled or did not
// complete their task successfully. This may require you to incorporate some
// additional error tracking capabilities into your operation objects.
// 
// # KVO-Compliant Properties
// 
// The [NSOperation] class is key-value coding (KVC) and key-value observing
// (KVO) compliant for several of its properties. As needed, you can observe
// these properties to control other parts of your application. To observe the
// properties, use the following key paths:
// 
// - `isCancelled` - read-only - `isAsynchronous` - read-only - `isExecuting`
// - read-only - `isFinished` - read-only - `isReady` - read-only -
// `dependencies` - read-only - `queuePriority` - readable and writable -
// `completionBlock` - readable and writable
// 
// Although you can attach observers to these properties, you should not use
// Cocoa bindings to bind them to elements of your application’s user
// interface. Code associated with your user interface typically must execute
// only in your application’s main thread. Because an operation may execute
// in any thread, KVO notifications associated with that operation may
// similarly occur in any thread.
// 
// If you provide custom implementations for any of the preceding properties,
// your implementations must maintain KVC and KVO compliance. If you define
// additional properties for your [NSOperation] objects, it is recommended
// that you make those properties KVC and KVO compliant as well. For
// information on how to support key-value coding, see [Key-Value Coding
// Programming Guide]. For information on how to support key-value observing,
// see [Key-Value Observing Programming Guide].
// 
// # Multicore Considerations
// 
// The [NSOperation] class is itself multicore aware. It is therefore safe to
// call the methods of an [NSOperation] object from multiple threads without
// creating additional locks to synchronize access to the object. This
// behavior is necessary because an operation typically runs in a separate
// thread from the one that created and is monitoring it.
// 
// When you subclass [NSOperation], you must make sure that any overridden
// methods remain safe to call from multiple threads. If you implement custom
// methods in your subclass, such as custom data accessors, you must also make
// sure those methods are thread-safe. Thus, access to any data variables in
// the operation must be synchronized to prevent potential data corruption.
// For more information about synchronization, see [Threading Programming
// Guide].
// 
// # Asynchronous Versus Synchronous Operations
// 
// If you plan on executing an operation object manually, instead of adding it
// to a queue, you can design your operation to execute in a synchronous or
// asynchronous manner. Operation objects are synchronous by default. In a
// synchronous operation, the operation object does not create a separate
// thread on which to run its task. When you call the [Start] method of a
// synchronous operation directly from your code, the operation executes
// immediately in the current thread. By the time the [Start] method of such
// an object returns control to the caller, the task itself is complete.
// 
// When you call the [Start] method of an asynchronous operation, that method
// may return before the corresponding task is completed. An asynchronous
// operation object is responsible for scheduling its task on a separate
// thread. The operation could do that by starting a new thread directly, by
// calling an asynchronous method, or by submitting a block to a dispatch
// queue for execution. It does not actually matter if the operation is
// ongoing when control returns to the caller, only that it could be ongoing.
// 
// If you always plan to use queues to execute your operations, it is simpler
// to define them as synchronous. If you execute operations manually, though,
// you might want to define your operation objects as asynchronous. Defining
// an asynchronous operation requires more work, because you have to monitor
// the ongoing state of your task and report changes in that state using KVO
// notifications. But defining asynchronous operations is useful in cases
// where you want to ensure that a manually executed operation does not block
// the calling thread.
// 
// When you add an operation to an operation queue, the queue ignores the
// value of the [Asynchronous] property and always calls the [Start] method
// from a separate thread. Therefore, if you always run operations by adding
// them to an operation queue, there is no reason to make them asynchronous.
// 
// For information on how to define both synchronous and asynchronous
// operations, see the subclassing notes.
// 
// # Subclassing Notes
// 
// The [NSOperation] class provides the basic logic to track the execution
// state of your operation but otherwise must be subclassed to do any real
// work. How you create your subclass depends on whether your operation is
// designed to execute concurrently or non-concurrently.
// 
// # Methods to Override
// 
// For non-concurrent operations, you typically override only one method:
// 
// - [Main]
// 
// Into this method, you place the code needed to perform the given task. Of
// course, you should also define a custom initialization method to make it
// easier to create instances of your custom class. You might also want to
// define getter and setter methods to access the data from the operation.
// However, if you do define custom getter and setter methods, you must make
// sure those methods can be called safely from multiple threads.
// 
// If you are creating a concurrent operation, you need to override the
// following methods and properties at a minimum:
// 
// - [Start]
// - [Asynchronous]
// - [Executing]
// - [Finished]
// 
// In a concurrent operation, your [Start] method is responsible for starting
// the operation in an asynchronous manner. Whether you spawn a thread or call
// an asynchronous function, you do it from this method. Upon starting the
// operation, your [Start] method should also update the execution state of
// the operation as reported by the [Executing] property. You do this by
// sending out KVO notifications for the [Executing] key path, which lets
// interested clients know that the operation is now running. Your [Executing]
// property must also provide the status in a thread-safe manner.
// 
// Upon completion or cancellation of its task, your concurrent operation
// object must generate KVO notifications for both the `isExecuting` and
// `isFinished` key paths to mark the final change of state for your
// operation. (In the case of cancellation, it is still important to update
// the `isFinished` key path, even if the operation did not completely finish
// its task. Queued operations must report that they are finished before they
// can be removed from a queue.) In addition to generating KVO notifications,
// your overrides of the [Executing] and [Finished] properties should also
// continue to report accurate values based on the state of your operation.
// 
// For additional information and guidance on how to define concurrent
// operations, see [Concurrency Programming Guide].
// 
// Even for concurrent operations, there should be little need to override
// methods other than those described above. However, if you customize the
// dependency features of operations, you might have to override additional
// methods and provide additional KVO notifications. In the case of
// dependencies, this would likely only require providing notifications for
// the `isReady` key path. Because the [Dependencies] property contains the
// list of dependent operations, changes to it are already handled by the
// default [NSOperation] class.
// 
// # Maintaining Operation Object States
// 
// Operation objects maintain state information internally to determine when
// it is safe to execute and also to notify external clients of the
// progression through the operation’s life cycle. Your custom subclasses
// maintains this state information to ensure the correct execution of
// operations in your code. The key paths associated with an operation’s
// states are:
// 
// In most cases, you do not have to manage the state of this key path
// yourself. If the readiness of your operations is determined by factors
// other than dependent operations, however—such as by some external
// condition in your program—you can provide your own implementation of the
// [Ready] property and track your operation’s readiness yourself. It is
// often simpler though just to create operation objects only when your
// external state allows it.
// 
// In macOS 10.6 and later, if you cancel an operation while it is waiting on
// the completion of one or more dependent operations, those dependencies are
// thereafter ignored and the value of this property is updated to reflect
// that it is now ready to run. This behavior gives an operation queue the
// chance to flush cancelled operations out of its queue more quickly.
// 
// If you replace the [Start] method of your operation object, you must also
// replace the [Executing] property and generate KVO notifications when the
// execution state of your operation changes.
// 
// If you replace the [Start] method or your operation object, you must also
// replace the [Finished] property and generate KVO notifications when the
// operation finishes executing or is cancelled.
// 
// # Responding to the Cancel Command
// 
// Once you add an operation to a queue, the operation is out of your hands.
// The queue takes over and handles the scheduling of that task. However, if
// you decide later that you do not want to execute the operation after
// all—because the user pressed a cancel button in a progress panel or quit
// the application, for example—you can cancel the operation to prevent it
// from consuming CPU time needlessly. You do this by calling the [Cancel]
// method of the operation object itself or by calling the
// [CancelAllOperations] method of the [NSOperationQueue] class.
// 
// Canceling an operation does not immediately force it to stop what it is
// doing. Although respecting the value in the [Cancelled] property is
// expected of all operations, your code must explicitly check the value in
// this property and abort as needed. The default implementation of
// [NSOperation] includes checks for cancellation. For example, if you cancel
// an operation before its [Start] method is called, the [Start] method exits
// without starting the task.
// 
// You should always support cancellation semantics in any custom code you
// write. In particular, your main task code should periodically check the
// value of the [Cancelled] property. If the property reports the value
// [true], your operation object should clean up and exit as quickly as
// possible. If you implement a custom [Start] method, that method should
// include early checks for cancellation and behave appropriately. Your custom
// [Start] method must be prepared to handle this type of early cancellation.
// 
// In addition to simply exiting when an operation is cancelled, it is also
// important that you move a cancelled operation to the appropriate final
// state. Specifically, if you manage the values for the [Finished] and
// [Executing] properties yourself (perhaps because you are implementing a
// concurrent operation), you must update those properties accordingly.
// Specifically, you must change the value returned by [Finished] to [true]
// and the value returned by [Executing] to [false]. You must make these
// changes even if the operation was cancelled before it started executing.
//
// [Concurrency Programming Guide]: https://developer.apple.com/library/archive/documentation/General/Conceptual/ConcurrencyProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40008091
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
// [Threading Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Multithreading/Introduction/Introduction.html#//apple_ref/doc/uid/10000057i
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Executing the Operation
//
//   - [Operation.Start]: Begins the execution of the operation.
//   - [Operation.Main]: Performs the receiver’s non-concurrent task.
//   - [Operation.CompletionBlock]: The block to execute after the operation’s main task is completed.
//   - [Operation.SetCompletionBlock]
//
// # Canceling Operations
//
//   - [Operation.Cancel]: Advises the operation object that it should stop executing its task.
//
// # Getting the Operation Status
//
//   - [Operation.Cancelled]: A Boolean value indicating whether the operation has been cancelled
//   - [Operation.Executing]: A Boolean value indicating whether the operation is currently executing.
//   - [Operation.Finished]: A Boolean value indicating whether the operation has finished executing its task.
//   - [Operation.Concurrent]: A Boolean value indicating whether the operation executes its task asynchronously.
//   - [Operation.Asynchronous]: A Boolean value indicating whether the operation executes its task asynchronously.
//   - [Operation.Ready]: A Boolean value indicating whether the operation can be performed now.
//   - [Operation.Name]: The name of the operation.
//   - [Operation.SetName]
//
// # Managing Dependencies
//
//   - [Operation.AddDependency]: Makes the receiver dependent on the completion of the specified operation.
//   - [Operation.RemoveDependency]: Removes the receiver’s dependence on the specified operation.
//   - [Operation.Dependencies]: An array of the operation objects that must finish executing before the current object can begin executing.
//
// # Configuring the Execution Priority
//
//   - [Operation.QualityOfService]: The relative amount of importance for granting system resources to the operation.
//   - [Operation.SetQualityOfService]
//   - [Operation.QueuePriority]: The execution priority of the operation in an operation queue.
//   - [Operation.SetQueuePriority]
//
// # Waiting on an Operation Object
//
//   - [Operation.WaitUntilFinished]: Blocks execution of the current thread until the operation object finishes its task.
//
// See: https://developer.apple.com/documentation/Foundation/Operation
type Operation struct {
	objectivec.Object
}

// OperationFromID constructs a [Operation] from an objc.ID.
//
// An abstract class that represents the code and data associated with a
// single task.
func OperationFromID(id objc.ID) Operation {
	return NSOperation{objectivec.Object{ID: id}}
}

// NSOperationFromID is an alias for [OperationFromID] for cross-framework compatibility.
func NSOperationFromID(id objc.ID) Operation { return OperationFromID(id) }
// NOTE: Operation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [Operation] class.
//
// # Executing the Operation
//
//   - [IOperation.Start]: Begins the execution of the operation.
//   - [IOperation.Main]: Performs the receiver’s non-concurrent task.
//   - [IOperation.CompletionBlock]: The block to execute after the operation’s main task is completed.
//   - [IOperation.SetCompletionBlock]
//
// # Canceling Operations
//
//   - [IOperation.Cancel]: Advises the operation object that it should stop executing its task.
//
// # Getting the Operation Status
//
//   - [IOperation.Cancelled]: A Boolean value indicating whether the operation has been cancelled
//   - [IOperation.Executing]: A Boolean value indicating whether the operation is currently executing.
//   - [IOperation.Finished]: A Boolean value indicating whether the operation has finished executing its task.
//   - [IOperation.Concurrent]: A Boolean value indicating whether the operation executes its task asynchronously.
//   - [IOperation.Asynchronous]: A Boolean value indicating whether the operation executes its task asynchronously.
//   - [IOperation.Ready]: A Boolean value indicating whether the operation can be performed now.
//   - [IOperation.Name]: The name of the operation.
//   - [IOperation.SetName]
//
// # Managing Dependencies
//
//   - [IOperation.AddDependency]: Makes the receiver dependent on the completion of the specified operation.
//   - [IOperation.RemoveDependency]: Removes the receiver’s dependence on the specified operation.
//   - [IOperation.Dependencies]: An array of the operation objects that must finish executing before the current object can begin executing.
//
// # Configuring the Execution Priority
//
//   - [IOperation.QualityOfService]: The relative amount of importance for granting system resources to the operation.
//   - [IOperation.SetQualityOfService]
//   - [IOperation.QueuePriority]: The execution priority of the operation in an operation queue.
//   - [IOperation.SetQueuePriority]
//
// # Waiting on an Operation Object
//
//   - [IOperation.WaitUntilFinished]: Blocks execution of the current thread until the operation object finishes its task.
//
// See: https://developer.apple.com/documentation/Foundation/Operation
type IOperation interface {
	objectivec.IObject

	// Topic: Executing the Operation

	// Begins the execution of the operation.
	Start()
	// Performs the receiver’s non-concurrent task.
	Main()
	// The block to execute after the operation’s main task is completed.
	CompletionBlock() VoidHandler
	SetCompletionBlock(value VoidHandler)

	// Topic: Canceling Operations

	// Advises the operation object that it should stop executing its task.
	Cancel()

	// Topic: Getting the Operation Status

	// A Boolean value indicating whether the operation has been cancelled
	Cancelled() bool
	// A Boolean value indicating whether the operation is currently executing.
	Executing() bool
	// A Boolean value indicating whether the operation has finished executing its task.
	Finished() bool
	// A Boolean value indicating whether the operation executes its task asynchronously.
	Concurrent() bool
	// A Boolean value indicating whether the operation executes its task asynchronously.
	Asynchronous() bool
	// A Boolean value indicating whether the operation can be performed now.
	Ready() bool
	// The name of the operation.
	Name() string
	SetName(value string)

	// Topic: Managing Dependencies

	// Makes the receiver dependent on the completion of the specified operation.
	AddDependency(op INSOperation)
	// Removes the receiver’s dependence on the specified operation.
	RemoveDependency(op INSOperation)
	// An array of the operation objects that must finish executing before the current object can begin executing.
	Dependencies() []NSOperation

	// Topic: Configuring the Execution Priority

	// The relative amount of importance for granting system resources to the operation.
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	// The execution priority of the operation in an operation queue.
	QueuePriority() NSOperationQueuePriority
	SetQueuePriority(value NSOperationQueuePriority)

	// Topic: Waiting on an Operation Object

	// Blocks execution of the current thread until the operation object finishes its task.
	WaitUntilFinished()
}

// Init initializes the instance.
func (o Operation) Init() Operation {
	rv := objc.Send[Operation](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o Operation) Autorelease() Operation {
	rv := objc.Send[Operation](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOperation creates a new Operation instance.
func NewOperation() Operation {
	class := getOperationClass()
	rv := objc.Send[Operation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Begins the execution of the operation.
//
// # Discussion
// 
// The default implementation of this method updates the execution state of
// the operation and calls the receiver’s [Main] method. This method also
// performs several checks to ensure that the operation can actually run. For
// example, if the receiver was cancelled or is already finished, this method
// simply returns without calling [Main]. (In OS X v10.5, this method throws
// an exception if the operation is already finished.) If the operation is
// currently executing or is not ready to execute, this method throws an
// [NSInvalidArgumentException] exception. In OS X v10.5, this method catches
// and ignores any exceptions thrown by your [Main] method automatically. In
// macOS 10.6 and later, exceptions are allowed to propagate beyond this
// method. You should never allow exceptions to propagate out of your [Main]
// method.
// 
// If you are implementing a concurrent operation, you must override this
// method and use it to initiate your operation. Your custom implementation
// must not call `super` at any time. In addition to configuring the execution
// environment for your task, your implementation of this method must also
// track the state of the operation and provide appropriate state transitions.
// When the operation executes and subsequently finishes its work, it should
// generate KVO notifications for the `isExecuting` and `isFinished` key paths
// respectively. For more information about manually generating KVO
// notifications, see [Key-Value Observing Programming Guide].
// 
// You can call this method explicitly if you want to execute your operations
// manually. However, it is a programmer error to call this method on an
// operation object that is already in an operation queue or to queue the
// operation after calling this method. Once you add an operation object to a
// queue, the queue assumes all responsibility for it.
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
//
// See: https://developer.apple.com/documentation/Foundation/Operation/start()
func (o Operation) Start() {
	objc.Send[objc.ID](o.ID, objc.Sel("start"))
}

// Performs the receiver’s non-concurrent task.
//
// # Discussion
// 
// The default implementation of this method does nothing. You should override
// this method to perform the desired task. In your implementation, do not
// invoke `super`. This method will automatically execute within an
// autorelease pool provided by [NSOperation], so you do not need to create
// your own autorelease pool block in your implementation.
// 
// If you are implementing a concurrent operation, you are not required to
// override this method but may do so if you plan to call it from your custom
// [Start] method.
//
// See: https://developer.apple.com/documentation/Foundation/Operation/main()
func (o Operation) Main() {
	objc.Send[objc.ID](o.ID, objc.Sel("main"))
}

// Advises the operation object that it should stop executing its task.
//
// # Discussion
// 
// This method does not force your operation code to stop. Instead, it updates
// the object’s internal flags to reflect the change in state. If the
// operation has already finished executing, this method has no effect.
// Canceling an operation that is currently in an operation queue, but not yet
// executing, makes it possible to remove the operation from the queue sooner
// than usual.
// 
// In macOS 10.6 and later, if an operation is in a queue but waiting on
// unfinished dependent operations, those operations are subsequently ignored.
// Because it is already cancelled, this behavior allows the operation queue
// to call the operation’s [Start] method sooner and clear the object out of
// the queue. If you cancel an operation that is not in a queue, this method
// immediately marks the object as finished. In each case, marking the object
// as ready or finished results in the generation of the appropriate KVO
// notifications.
// 
// In versions of macOS prior to 10.6, an operation object remains in the
// queue until all of its dependencies are removed through the normal
// processes. Thus, the operation must wait until all of its dependent
// operations finish executing or are themselves cancelled and have their
// [Start] method called.
// 
// For more information on what you must do in your operation objects to
// support cancellation, see [NSOperation].
//
// See: https://developer.apple.com/documentation/Foundation/Operation/cancel()
func (o Operation) Cancel() {
	objc.Send[objc.ID](o.ID, objc.Sel("cancel"))
}

// Makes the receiver dependent on the completion of the specified operation.
//
// op: The operation on which the receiver should depend. The same dependency
// should not be added more than once to the receiver, and the results of
// doing so are undefined.
//
// # Discussion
// 
// The receiver is not considered ready to execute until all of its dependent
// operations have finished executing. If the receiver is already executing
// its task, adding dependencies has no practical effect. This method may
// change the `isReady` and `dependencies` properties of the receiver.
// 
// It is a programmer error to create any circular dependencies among a set of
// operations. Doing so can cause a deadlock among the operations and may
// freeze your program.
//
// See: https://developer.apple.com/documentation/Foundation/Operation/addDependency(_:)
func (o Operation) AddDependency(op INSOperation) {
	objc.Send[objc.ID](o.ID, objc.Sel("addDependency:"), op)
}

// Removes the receiver’s dependence on the specified operation.
//
// op: The dependent operation to be removed from the receiver.
//
// # Discussion
// 
// This method may change the `isReady` and `dependencies` properties of the
// receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Operation/removeDependency(_:)
func (o Operation) RemoveDependency(op INSOperation) {
	objc.Send[objc.ID](o.ID, objc.Sel("removeDependency:"), op)
}

// Blocks execution of the current thread until the operation object finishes
// its task.
//
// # Discussion
// 
// An operation object must never call this method on itself and should avoid
// calling it on any operations submitted to the same operation queue as
// itself. Doing so can cause the operation to deadlock. Instead, other parts
// of your app may call this method as needed to prevent other tasks from
// completing until the target operation object finishes. It is generally safe
// to call this method on an operation that is in a different operation queue,
// although it is still possible to create deadlocks if each operation waits
// on the other.
// 
// A typical use for this method would be to call it from the code that
// created the operation in the first place. After submitting the operation to
// a queue, you would call this method to wait until that operation finished
// executing.
//
// See: https://developer.apple.com/documentation/Foundation/Operation/waitUntilFinished()
func (o Operation) WaitUntilFinished() {
	objc.Send[objc.ID](o.ID, objc.Sel("waitUntilFinished"))
}

// The block to execute after the operation’s main task is completed.
//
// # Discussion
// 
// The completion block takes no parameters and has no return value.
// 
// The exact execution context for your completion block is not guaranteed but
// is typically a secondary thread. Therefore, you should not use this block
// to do any work that requires a very specific execution context. Instead,
// you should shunt that work to your application’s main thread or to the
// specific thread that is capable of doing it. For example, if you have a
// custom thread for coordinating the completion of the operation, you could
// use the completion block to ping that thread.
// 
// The completion block you provide is executed when the value in the
// [Finished] property changes to [true]. Because the completion block
// executes after the operation indicates it has finished its task, you must
// not use a completion block to queue additional work considered to be part
// of that task. An operation object whose [Finished] property contains the
// value [true] must be done with all of its task-related work by definition.
// The completion block should be used to notify interested objects that the
// work is complete or perform other tasks that might be related to, but not
// part of, the operation’s actual task.
// 
// A finished operation may finish either because it was cancelled or because
// it successfully completed its task. You should take that fact into account
// when writing your block code. Similarly, you should not make any
// assumptions about the successful completion of dependent operations, which
// may themselves have been cancelled.
// 
// In iOS 8 and later and macOS 10.10 and later, this property is set to `nil`
// after the completion block begins executing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (o Operation) CompletionBlock() VoidHandler {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("completionBlock"))
	_ = rv
	return nil
}
func (o Operation) SetCompletionBlock(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](o.ID, objc.Sel("setCompletionBlock:"), block)
}

// A Boolean value indicating whether the operation has been cancelled
//
// # Discussion
// 
// The default value of this property is [false]. Calling the [Cancel] method
// of this object sets the value of this property to [true]. Once canceled, an
// operation must move to the finished state.
// 
// Canceling an operation does not actively stop the receiver’s code from
// executing. An operation object is responsible for calling this method
// periodically and stopping itself if the method returns [true].
// 
// You should always check the value of this property before doing any work
// towards accomplishing the operation’s task, which typically means
// checking it at the beginning of your custom [Main] method. It is possible
// for an operation to be cancelled before it begins executing or at any time
// while it is executing. Therefore, checking the value at the beginning of
// your [Main] method (and periodically throughout that method) lets you exit
// as quickly as possible when an operation is cancelled.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Operation/isCancelled
func (o Operation) Cancelled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isCancelled"))
	return rv
}

// A Boolean value indicating whether the operation is currently executing.
//
// # Discussion
// 
// The value of this property is [true] if the operation is currently
// executing its main task or [false] if it is not.
// 
// When implementing a concurrent operation object, you must override the
// implementation of this property so that you can return the execution state
// of your operation. In your custom implementation, you must generate KVO
// notifications for the `isExecuting` key path whenever the execution state
// of your operation object changes. For more information about manually
// generating KVO notifications, see [Key-Value Observing Programming Guide].
// 
// You do not need to reimplement this property for nonconcurrent operations.
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Operation/isExecuting
func (o Operation) Executing() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isExecuting"))
	return rv
}

// A Boolean value indicating whether the operation has finished executing its
// task.
//
// # Discussion
// 
// The value of this property is [true] if the operation has finished its main
// task or [false] if it is executing that task or has not yet started it.
// 
// When implementing a concurrent operation object, you must override the
// implementation of this property so that you can return the finished state
// of your operation. In your custom implementation, you must generate KVO
// notifications for the `isFinished` key path whenever the finished state of
// your operation object changes. For more information about manually
// generating KVO notifications, see [Key-Value Observing Programming Guide].
// 
// You do not need to reimplement this property for nonconcurrent operations.
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Operation/isFinished
func (o Operation) Finished() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isFinished"))
	return rv
}

// A Boolean value indicating whether the operation executes its task
// asynchronously.
//
// # Discussion
// 
// Use the [Asynchronous] property instead.
// 
// The value of this property is [true] for operations that run asynchronously
// with respect to the current thread or [false] for operations that run
// synchronously on the current thread. The default value of this property is
// [false].
// 
// In macOS 10.6 and later, operation queues ignore the value in this property
// and always start operations on a separate thread.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Operation/isConcurrent
func (o Operation) Concurrent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isConcurrent"))
	return rv
}

// A Boolean value indicating whether the operation executes its task
// asynchronously.
//
// # Discussion
// 
// The value of this property is [true] for operations that run asynchronously
// with respect to the current thread or [false] for operations that run
// synchronously on the current thread. The default value of this property is
// [false].
// 
// When implementing an asynchronous operation object, you must implement this
// property and return [true]. For more information about how to implement an
// asynchronous operation, see [NSOperation].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Operation/isAsynchronous
func (o Operation) Asynchronous() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAsynchronous"))
	return rv
}

// A Boolean value indicating whether the operation can be performed now.
//
// # Discussion
// 
// The readiness of operations is determined by their dependencies on other
// operations and potentially by custom conditions that you define. The
// [NSOperation] class manages dependencies on other operations and reports
// the readiness of the receiver based on those dependencies.
// 
// If you want to use custom conditions to define the readiness of your
// operation object, reimplement this property and return a value that
// accurately reflects the readiness of the receiver. If you do so, your
// custom implementation must get the default property value from `super` and
// incorporate that readiness value into the new value of the property. In
// your custom implementation, you must generate KVO notifications for the
// `isReady` key path whenever the ready state of your operation object
// changes. For more information about generating KVO notifications, see
// [Key-Value Observing Programming Guide].
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
//
// See: https://developer.apple.com/documentation/Foundation/Operation/isReady
func (o Operation) Ready() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isReady"))
	return rv
}

// The name of the operation.
//
// # Discussion
// 
// Assign a name to the operation object to help identify it during debugging.
//
// See: https://developer.apple.com/documentation/Foundation/Operation/name
func (o Operation) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
func (o Operation) SetName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setName:"), objc.String(value))
}

// An array of the operation objects that must finish executing before the
// current object can begin executing.
//
// # Discussion
// 
// This property contains an array of [NSOperation] objects. To add an object
// to this array, use the [AddDependency] method.
// 
// An operation object must not execute until all of its dependent operations
// finish executing. Operations are not removed from this dependency list as
// they finish executing. You can use this list to track all dependent
// operations, including those that have already finished executing. The only
// way to remove an operation from this list is to use the [RemoveDependency]
// method.
//
// See: https://developer.apple.com/documentation/Foundation/Operation/dependencies
func (o Operation) Dependencies() []NSOperation {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("dependencies"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSOperation {
		return NSOperationFromID(id)
	})
}

// The relative amount of importance for granting system resources to the
// operation.
//
// # Discussion
// 
// Service levels affect the priority with which an operation object is given
// access to system resources such as CPU time, network resources, disk
// resources, and so on. Operations with a higher quality of service level are
// given greater priority over system resources so that they may perform their
// task more quickly. You use service levels to ensure that operations
// responding to explicit user requests are given priority over less critical
// work.
// 
// This property reflects the minimum service level needed to execute the
// operation effectively. The default value of this property is
// [QualityOfServiceDefault] and you should leave that value in place whenever
// possible. When changing the service level, use the minimum level that is
// appropriate for executing the corresponding task. For example, if the user
// initiates a task and is waiting for it to finish, assign the value
// [QualityOfServiceUserInteractive] to this property. The system may give the
// operation a higher service level to the operation if the resources are
// available to do so. For additional information, see [Prioritize Work with
// Quality of Service Classes] in [Energy Efficiency Guide for iOS Apps] and
// [Prioritize Work at the Task Level] in [Energy Efficiency Guide for Mac
// Apps].
//
// [Energy Efficiency Guide for Mac Apps]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/power_efficiency_guidelines_osx/index.html#//apple_ref/doc/uid/TP40013929
// [Energy Efficiency Guide for iOS Apps]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/EnergyGuide-iOS/index.html#//apple_ref/doc/uid/TP40015243
// [Prioritize Work at the Task Level]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/power_efficiency_guidelines_osx/PrioritizeWorkAtTheTaskLevel.html#//apple_ref/doc/uid/TP40013929-CH35
// [Prioritize Work with Quality of Service Classes]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/EnergyGuide-iOS/PrioritizeWorkWithQoS.html#//apple_ref/doc/uid/TP40015243-CH39
//
// See: https://developer.apple.com/documentation/Foundation/Operation/qualityOfService
func (o Operation) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](o.ID, objc.Sel("qualityOfService"))
	return QualityOfService(rv)
}
func (o Operation) SetQualityOfService(value QualityOfService) {
	objc.Send[struct{}](o.ID, objc.Sel("setQualityOfService:"), value)
}

// The execution priority of the operation in an operation queue.
//
// # Discussion
// 
// This property contains the relative priority of the operation. This value
// is used to influence the order in which operations are dequeued and
// executed. The returned value always corresponds to one of the predefined
// constants. (For a list of valid values, see [Operation.QueuePriority].) If
// no priority is explicitly set, this method returns
// [NSOperationQueuePriorityNormal].
// 
// You should use priority values only as needed to classify the relative
// priority of non-dependent operations. Priority values should not be used to
// implement dependency management among different operation objects. If you
// need to establish dependencies between operations, use the [AddDependency]
// method instead.
// 
// If you attempt to specify a priority value that does not match one of the
// defined constants, the operation object automatically adjusts the value you
// specify towards the [NSOperationQueuePriorityNormal] priority, stopping at
// the first valid constant value. For example, if you specified the value
// -10, the operation would adjust that value to match the
// [NSOperationQueuePriorityVeryLow] constant. Similarly, if you specified
// +10, this operation would adjust the value to match the
// [NSOperationQueuePriorityVeryHigh] constant.
//
// [Operation.QueuePriority]: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum
//
// See: https://developer.apple.com/documentation/Foundation/Operation/queuePriority-swift.property
func (o Operation) QueuePriority() NSOperationQueuePriority {
	rv := objc.Send[NSOperationQueuePriority](o.ID, objc.Sel("queuePriority"))
	return NSOperationQueuePriority(rv)
}
func (o Operation) SetQueuePriority(value NSOperationQueuePriority) {
	objc.Send[struct{}](o.ID, objc.Sel("setQueuePriority:"), value)
}


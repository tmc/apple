// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSTask] class.
var (
	_FSTaskClass     FSTaskClass
	_FSTaskClassOnce sync.Once
)

func getFSTaskClass() FSTaskClass {
	_FSTaskClassOnce.Do(func() {
		_FSTaskClass = FSTaskClass{class: objc.GetClass("FSTask")}
	})
	return _FSTaskClass
}

// GetFSTaskClass returns the class object for FSTask.
func GetFSTaskClass() FSTaskClass {
	return getFSTaskClass()
}

type FSTaskClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSTaskClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSTaskClass) Alloc() FSTask {
	rv := objc.Send[FSTask](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A class that enables a file system module to pass log messages and
// completion notifications to clients.
//
// # Overview
//
// FSKit creates an instance of this class for each long-running operations.
//
// # Logging
//
//   - [FSTask.LogMessage]: Logs the given string to the initiating client.
//
// # Sending completion messages
//
//   - [FSTask.DidCompleteWithError]: Informs the client that the task completed.
//
// # Handling task cancellation
//
//   - [FSTask.CancellationHandler]: A handler called by FSKit upon canceling the task.
//   - [FSTask.SetCancellationHandler]
//
// See: https://developer.apple.com/documentation/FSKit/FSTask
type FSTask struct {
	objectivec.Object
}

// FSTaskFromID constructs a [FSTask] from an objc.ID.
//
// A class that enables a file system module to pass log messages and
// completion notifications to clients.
func FSTaskFromID(id objc.ID) FSTask {
	return FSTask{objectivec.Object{ID: id}}
}

// NOTE: FSTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSTask] class.
//
// # Logging
//
//   - [IFSTask.LogMessage]: Logs the given string to the initiating client.
//
// # Sending completion messages
//
//   - [IFSTask.DidCompleteWithError]: Informs the client that the task completed.
//
// # Handling task cancellation
//
//   - [IFSTask.CancellationHandler]: A handler called by FSKit upon canceling the task.
//   - [IFSTask.SetCancellationHandler]
//
// See: https://developer.apple.com/documentation/FSKit/FSTask
type IFSTask interface {
	objectivec.IObject

	// Topic: Logging

	// Logs the given string to the initiating client.
	LogMessage(str string)

	// Topic: Sending completion messages

	// Informs the client that the task completed.
	DidCompleteWithError(error_ foundation.NSError)

	// Topic: Handling task cancellation

	// A handler called by FSKit upon canceling the task.
	CancellationHandler() VoidHandler
	SetCancellationHandler(value VoidHandler)
}

// Init initializes the instance.
func (t FSTask) Init() FSTask {
	rv := objc.Send[FSTask](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t FSTask) Autorelease() FSTask {
	rv := objc.Send[FSTask](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSTask creates a new FSTask instance.
func NewFSTask() FSTask {
	class := getFSTaskClass()
	rv := objc.Send[FSTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Logs the given string to the initiating client.
//
// str: The string to log.
//
// See: https://developer.apple.com/documentation/FSKit/FSTask/logMessage(_:)
func (t FSTask) LogMessage(str string) {
	objc.Send[objc.ID](t.ID, objc.Sel("logMessage:"), objc.String(str))
}

// Informs the client that the task completed.
//
// error: `nil` if the task completed successfully; otherwise, an error that caused
// the task to fail.
//
// See: https://developer.apple.com/documentation/FSKit/FSTask/didComplete(error:)
func (t FSTask) DidCompleteWithError(error_ foundation.NSError) {
	objc.Send[objc.ID](t.ID, objc.Sel("didCompleteWithError:"), error_)
}

// A handler called by FSKit upon canceling the task.
//
// # Discussion
//
// FSKit calls the cancellation handler within an independent execution
// context.
//
// If the handler can’t complete its work successfully, it can return an
// error from the block or closure. FSKit logs any returned error and then
// terminates all activity in the container.
//
// The task object clears its `cancellationHandler` property after the
// task’s cancellation or completion. This helps accelerate the cleanup of
// retained state.
//
// The exact structuring of the completion handler depends on the structuring
// of the code imlementing the task. As a concrete example, consider a check
// operation with the following class:
//
// and a `startCheckWithTask` method with a helper method `performCheck` like
// the following:
//
// When canceled, the handler block in this example sets the checker’s
// `interrupted` property, and then calls the [DispatchGroup] method [wait()]
// (Swift) or the function [dispatch_group_wait] (Objective-C) on the
// checker’s work group. Because neither of these operations can fail, the
// handler returns `nil` to indicate it didn’t encounter an error.
//
// For simplicity, this example doesn’t account for errors, whereas
// production code must do so. Furthermore, when fully implemented, the
// `performCheck` method should perform a check operation. Specifically, it
// should periodically update the progress object and check its `interrupted`
// variable. The check can either complete successfully, complete with an
// error, or enter the interrupted state. It should then call
// [FSTask.DidCompleteWithError] wtih the appropriate error value or `nil`.
// Finally it should call `context.Work_groupXCUIElementTypeLeave()` (Swift)
// or `dispatch_group_leave(context.Work_group())` (Objective-C) to remove
// itself from its dispatch group.
//
// See: https://developer.apple.com/documentation/FSKit/FSTask/cancellationHandler
//
// [DispatchGroup]: https://developer.apple.com/documentation/Dispatch/DispatchGroup
// [dispatch_group_wait]: https://developer.apple.com/documentation/Dispatch/dispatch_group_wait
// [wait()]: https://developer.apple.com/documentation/Dispatch/DispatchGroup/wait()
func (t FSTask) CancellationHandler() VoidHandler {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("cancellationHandler"))
	_ = rv
	return nil
}
func (t FSTask) SetCancellationHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](t.ID, objc.Sel("setCancellationHandler:"), block)
}

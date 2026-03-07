// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLTask] class.
var (
	_MLTaskClass     MLTaskClass
	_MLTaskClassOnce sync.Once
)

func getMLTaskClass() MLTaskClass {
	_MLTaskClassOnce.Do(func() {
		_MLTaskClass = MLTaskClass{class: objc.GetClass("MLTask")}
	})
	return _MLTaskClass
}

// GetMLTaskClass returns the class object for MLTask.
func GetMLTaskClass() MLTaskClass {
	return getMLTaskClass()
}

type MLTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLTaskClass) Alloc() MLTask {
	rv := objc.Send[MLTask](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// An abstract base class for machine learning tasks.
//
// # Overview
// 
// You don’t create use this class directly. Instead, use a class that
// inherits from this one, such as [MLUpdateTask].
//
// # Identifying a task
//
//   - [MLTask.TaskIdentifier]: A unique name of the task to distinguish it from all other tasks at runtime.
//
// # Starting and stopping a task
//
//   - [MLTask.Resume]: Begins or resumes a machine learning task.
//   - [MLTask.Cancel]: Cancels a machine learning task before it completes.
//
// # Checking the state of a task
//
//   - [MLTask.State]: The current state of the machine learning task.
//   - [MLTask.Error]: The underlying error if the task is in a failed state.
//
// See: https://developer.apple.com/documentation/CoreML/MLTask
type MLTask struct {
	objectivec.Object
}

// MLTaskFromID constructs a [MLTask] from an objc.ID.
//
// An abstract base class for machine learning tasks.
func MLTaskFromID(id objc.ID) MLTask {
	return MLTask{objectivec.Object{id}}
}
// NOTE: MLTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MLTask] class.
//
// # Identifying a task
//
//   - [IMLTask.TaskIdentifier]: A unique name of the task to distinguish it from all other tasks at runtime.
//
// # Starting and stopping a task
//
//   - [IMLTask.Resume]: Begins or resumes a machine learning task.
//   - [IMLTask.Cancel]: Cancels a machine learning task before it completes.
//
// # Checking the state of a task
//
//   - [IMLTask.State]: The current state of the machine learning task.
//   - [IMLTask.Error]: The underlying error if the task is in a failed state.
//
// See: https://developer.apple.com/documentation/CoreML/MLTask
type IMLTask interface {
	objectivec.IObject

	// Topic: Identifying a task

	// A unique name of the task to distinguish it from all other tasks at runtime.
	TaskIdentifier() string

	// Topic: Starting and stopping a task

	// Begins or resumes a machine learning task.
	Resume()
	// Cancels a machine learning task before it completes.
	Cancel()

	// Topic: Checking the state of a task

	// The current state of the machine learning task.
	State() MLTaskState
	// The underlying error if the task is in a failed state.
	Error() foundation.INSError
}





// Init initializes the instance.
func (t MLTask) Init() MLTask {
	rv := objc.Send[MLTask](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MLTask) Autorelease() MLTask {
	rv := objc.Send[MLTask](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTask creates a new MLTask instance.
func NewMLTask() MLTask {
	class := getMLTaskClass()
	rv := objc.Send[MLTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Begins or resumes a machine learning task.
//
// # Discussion
// 
// Use this method to start a task for the first time or resumes a task that
// has paused. Tasks pause when they notify your app’s progress handlers,
// such as those you provide to an [MLUpdateProgressHandlers] instance.
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/resume()
func (t MLTask) Resume() {
	objc.Send[objc.ID](t.ID, objc.Sel("resume"))
}

// Cancels a machine learning task before it completes.
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/cancel()
func (t MLTask) Cancel() {
	objc.Send[objc.ID](t.ID, objc.Sel("cancel"))
}












// A unique name of the task to distinguish it from all other tasks at
// runtime.
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/taskIdentifier
func (t MLTask) TaskIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("taskIdentifier"))
	return foundation.NSStringFromID(rv).String()
}



// The current state of the machine learning task.
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/state
func (t MLTask) State() MLTaskState {
	rv := objc.Send[MLTaskState](t.ID, objc.Sel("state"))
	return MLTaskState(rv)
}



// The underlying error if the task is in a failed state.
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/error
func (t MLTask) Error() foundation.INSError {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
























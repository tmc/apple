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

// Class returns the underlying Objective-C class pointer.
func (mc MLTaskClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLTaskClass) Alloc() MLTask {
	rv := objc.Send[MLTask](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLTask._canCancel]
//   - [MLTask._canComplete]
//   - [MLTask._canFail]
//   - [MLTask._canResume]
//   - [MLTask._canSuspend]
//   - [MLTask._resumeWithTaskContext]
//   - [MLTask.CompleteWithTaskContext]
//   - [MLTask.FailWithErrorTaskContext]
//   - [MLTask.ResumeWithTaskContext]
//   - [MLTask.SuspendWithTaskContext]
//   - [MLTask.SyncQueue]
//   - [MLTask.TaskStatesToString]
//   - [MLTask.InitWithState]
//   - [MLTask.Error]
//   - [MLTask.SetError]
//   - [MLTask.State]
//   - [MLTask.SetState]
// See: https://developer.apple.com/documentation/CoreML/MLTask
type MLTask struct {
	objectivec.Object
}

// MLTaskFromID constructs a [MLTask] from an objc.ID.
func MLTaskFromID(id objc.ID) MLTask {
	return MLTask{objectivec.Object{ID: id}}
}
// Ensure MLTask implements IMLTask.
var _ IMLTask = MLTask{}

// An interface definition for the [MLTask] class.
//
// # Methods
//
//   - [IMLTask._canCancel]
//   - [IMLTask._canComplete]
//   - [IMLTask._canFail]
//   - [IMLTask._canResume]
//   - [IMLTask._canSuspend]
//   - [IMLTask._resumeWithTaskContext]
//   - [IMLTask.CompleteWithTaskContext]
//   - [IMLTask.FailWithErrorTaskContext]
//   - [IMLTask.ResumeWithTaskContext]
//   - [IMLTask.SuspendWithTaskContext]
//   - [IMLTask.SyncQueue]
//   - [IMLTask.TaskStatesToString]
//   - [IMLTask.InitWithState]
//   - [IMLTask.Error]
//   - [IMLTask.SetError]
//   - [IMLTask.State]
//   - [IMLTask.SetState]
//
// See: https://developer.apple.com/documentation/CoreML/MLTask
type IMLTask interface {
	objectivec.IObject

	// Topic: Methods

	_canCancel() bool
	_canComplete() bool
	_canFail() bool
	_canResume() bool
	_canSuspend() bool
	_resumeWithTaskContext(context objectivec.IObject)
	CompleteWithTaskContext(context objectivec.IObject)
	FailWithErrorTaskContext(error_ objectivec.IObject, context objectivec.IObject)
	ResumeWithTaskContext(context objectivec.IObject)
	SuspendWithTaskContext(context objectivec.IObject)
	SyncQueue() objectivec.Object
	TaskStatesToString(string_ int64) objectivec.IObject
	InitWithState(state int64) MLTask
	Error() foundation.INSError
	SetError(value foundation.INSError)
	State() int64
	SetState(value int64)
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

//
// See: https://developer.apple.com/documentation/CoreML/MLTask/initWithState:
func NewTaskWithState(state int64) MLTask {
	instance := getMLTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithState:"), state)
	return MLTaskFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLTask/_canCancel
func (t MLTask) _canCancel() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_canCancel"))
	return rv
}

// CanCancel is an exported wrapper for the private method _canCancel.
func (t MLTask) CanCancel() bool {
	return t._canCancel()
}
// See: https://developer.apple.com/documentation/CoreML/MLTask/_canComplete
func (t MLTask) _canComplete() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_canComplete"))
	return rv
}

// CanComplete is an exported wrapper for the private method _canComplete.
func (t MLTask) CanComplete() bool {
	return t._canComplete()
}
// See: https://developer.apple.com/documentation/CoreML/MLTask/_canFail
func (t MLTask) _canFail() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_canFail"))
	return rv
}

// CanFail is an exported wrapper for the private method _canFail.
func (t MLTask) CanFail() bool {
	return t._canFail()
}
// See: https://developer.apple.com/documentation/CoreML/MLTask/_canResume
func (t MLTask) _canResume() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_canResume"))
	return rv
}

// CanResume is an exported wrapper for the private method _canResume.
func (t MLTask) CanResume() bool {
	return t._canResume()
}
// See: https://developer.apple.com/documentation/CoreML/MLTask/_canSuspend
func (t MLTask) _canSuspend() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_canSuspend"))
	return rv
}

// CanSuspend is an exported wrapper for the private method _canSuspend.
func (t MLTask) CanSuspend() bool {
	return t._canSuspend()
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/_resumeWithTaskContext:
func (t MLTask) _resumeWithTaskContext(context objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_resumeWithTaskContext:"), context)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/completeWithTaskContext:
func (t MLTask) CompleteWithTaskContext(context objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("completeWithTaskContext:"), context)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/failWithError:taskContext:
func (t MLTask) FailWithErrorTaskContext(error_ objectivec.IObject, context objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("failWithError:taskContext:"), error_, context)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/resumeWithTaskContext:
func (t MLTask) ResumeWithTaskContext(context objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("resumeWithTaskContext:"), context)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/suspendWithTaskContext:
func (t MLTask) SuspendWithTaskContext(context objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("suspendWithTaskContext:"), context)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/taskStatesToString:
func (t MLTask) TaskStatesToString(string_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("taskStatesToString:"), string_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTask/initWithState:
func (t MLTask) InitWithState(state int64) MLTask {
	rv := objc.Send[MLTask](t.ID, objc.Sel("initWithState:"), state)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTask/error
func (t MLTask) Error() foundation.INSError {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (t MLTask) SetError(value foundation.INSError) {
	objc.Send[struct{}](t.ID, objc.Sel("setError:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLTask/state
func (t MLTask) State() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("state"))
	return rv
}
func (t MLTask) SetState(value int64) {
	objc.Send[struct{}](t.ID, objc.Sel("setState:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLTask/syncQueue
func (t MLTask) SyncQueue() objectivec.Object {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("syncQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}


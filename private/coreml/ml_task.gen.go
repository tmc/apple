// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLTask](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

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
	Error() foundation.NSError
	SetError(value foundation.NSError)
	State() int64
	SetState(value int64)
}

// Init initializes the instance.
func (m MLTask) Init() MLTask {
	rv := objc.SendIfResponds[MLTask](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLTask) Autorelease() MLTask {
	rv := objc.SendIfResponds[MLTask](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTask creates a new MLTask instance.
func NewMLTask() MLTask {
	class := getMLTaskClass()
	rv := objc.SendIfResponds[MLTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTaskWithState(state int64) MLTask {
	instance := getMLTaskClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithState:"), state)
	return MLTaskFromID(rv)
}

func (m MLTask) _canCancel() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("_canCancel"))
	return rv
}

// CanCancel is an exported wrapper for the private method _canCancel.
func (m MLTask) CanCancel() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_canCancel")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_canCancel"}
		return false, err
	}
	return m._canCancel(), nil
}

// CanCanCancel reports whether the receiver responds to the private selector _canCancel.
func (m MLTask) CanCanCancel() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_canCancel"))
}
func (m MLTask) _canComplete() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("_canComplete"))
	return rv
}

// CanComplete is an exported wrapper for the private method _canComplete.
func (m MLTask) CanComplete() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_canComplete")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_canComplete"}
		return false, err
	}
	return m._canComplete(), nil
}

// CanCanComplete reports whether the receiver responds to the private selector _canComplete.
func (m MLTask) CanCanComplete() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_canComplete"))
}
func (m MLTask) _canFail() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("_canFail"))
	return rv
}

// CanFail is an exported wrapper for the private method _canFail.
func (m MLTask) CanFail() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_canFail")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_canFail"}
		return false, err
	}
	return m._canFail(), nil
}

// CanCanFail reports whether the receiver responds to the private selector _canFail.
func (m MLTask) CanCanFail() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_canFail"))
}
func (m MLTask) _canResume() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("_canResume"))
	return rv
}

// CanResume is an exported wrapper for the private method _canResume.
func (m MLTask) CanResume() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_canResume")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_canResume"}
		return false, err
	}
	return m._canResume(), nil
}

// CanCanResume reports whether the receiver responds to the private selector _canResume.
func (m MLTask) CanCanResume() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_canResume"))
}
func (m MLTask) _canSuspend() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("_canSuspend"))
	return rv
}

// CanSuspend is an exported wrapper for the private method _canSuspend.
func (m MLTask) CanSuspend() (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_canSuspend")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_canSuspend"}
		return false, err
	}
	return m._canSuspend(), nil
}

// CanCanSuspend reports whether the receiver responds to the private selector _canSuspend.
func (m MLTask) CanCanSuspend() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_canSuspend"))
}
func (m MLTask) _resumeWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_resumeWithTaskContext:"), context)
}
func (m MLTask) CompleteWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("completeWithTaskContext:"), context)
}
func (m MLTask) FailWithErrorTaskContext(error_ objectivec.IObject, context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("failWithError:taskContext:"), error_, context)
}
func (m MLTask) ResumeWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("resumeWithTaskContext:"), context)
}
func (m MLTask) SuspendWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("suspendWithTaskContext:"), context)
}
func (m MLTask) TaskStatesToString(string_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("taskStatesToString:"), string_)
	return objectivec.Object{ID: rv}
}
func (m MLTask) InitWithState(state int64) MLTask {
	rv := objc.SendIfResponds[MLTask](m.ID, objc.Sel("initWithState:"), state)
	return rv
}

func (m MLTask) Error() foundation.NSError {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (m MLTask) SetError(value foundation.NSError) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setError:"), value)
}
func (m MLTask) State() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("state"))
	return rv
}
func (m MLTask) SetState(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setState:"), value)
}
func (m MLTask) SyncQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("syncQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

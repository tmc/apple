// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLBackgroundTask] class.
var (
	_MLBackgroundTaskClass     MLBackgroundTaskClass
	_MLBackgroundTaskClassOnce sync.Once
)

func getMLBackgroundTaskClass() MLBackgroundTaskClass {
	_MLBackgroundTaskClassOnce.Do(func() {
		_MLBackgroundTaskClass = MLBackgroundTaskClass{class: objc.GetClass("MLBackgroundTask")}
	})
	return _MLBackgroundTaskClass
}

// GetMLBackgroundTaskClass returns the class object for MLBackgroundTask.
func GetMLBackgroundTaskClass() MLBackgroundTaskClass {
	return getMLBackgroundTaskClass()
}

type MLBackgroundTaskClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLBackgroundTaskClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLBackgroundTaskClass) Alloc() MLBackgroundTask {
	rv := objc.SendIfResponds[MLBackgroundTask](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLBackgroundTask.ActivityForScheduling]
//   - [MLBackgroundTask.EncodeWithCoder]
//   - [MLBackgroundTask.TaskIdentifier]
//   - [MLBackgroundTask.SetTaskIdentifier]
//   - [MLBackgroundTask.InitWithCoder]
type MLBackgroundTask struct {
	objectivec.Object
}

// MLBackgroundTaskFromID constructs a [MLBackgroundTask] from an objc.ID.
func MLBackgroundTaskFromID(id objc.ID) MLBackgroundTask {
	return MLBackgroundTask{objectivec.Object{ID: id}}
}

// Ensure MLBackgroundTask implements IMLBackgroundTask.
var _ IMLBackgroundTask = MLBackgroundTask{}

// An interface definition for the [MLBackgroundTask] class.
//
// # Methods
//
//   - [IMLBackgroundTask.ActivityForScheduling]
//   - [IMLBackgroundTask.EncodeWithCoder]
//   - [IMLBackgroundTask.TaskIdentifier]
//   - [IMLBackgroundTask.SetTaskIdentifier]
//   - [IMLBackgroundTask.InitWithCoder]
type IMLBackgroundTask interface {
	objectivec.IObject

	// Topic: Methods

	ActivityForScheduling() objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	TaskIdentifier() string
	SetTaskIdentifier(value string)
	InitWithCoder(coder foundation.INSCoder) MLBackgroundTask
}

// Init initializes the instance.
func (m MLBackgroundTask) Init() MLBackgroundTask {
	rv := objc.SendIfResponds[MLBackgroundTask](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLBackgroundTask) Autorelease() MLBackgroundTask {
	rv := objc.SendIfResponds[MLBackgroundTask](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBackgroundTask creates a new MLBackgroundTask instance.
func NewMLBackgroundTask() MLBackgroundTask {
	class := getMLBackgroundTaskClass()
	rv := objc.SendIfResponds[MLBackgroundTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBackgroundTaskWithCoder(coder objectivec.IObject) MLBackgroundTask {
	instance := getMLBackgroundTaskClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLBackgroundTaskFromID(rv)
}

func (m MLBackgroundTask) ActivityForScheduling() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("activityForScheduling"))
	return objectivec.Object{ID: rv}
}
func (m MLBackgroundTask) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (m MLBackgroundTask) InitWithCoder(coder foundation.INSCoder) MLBackgroundTask {
	rv := objc.SendIfResponds[MLBackgroundTask](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_MLBackgroundTaskClass MLBackgroundTaskClass) CancelAllTasks() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLBackgroundTaskClass.class), objc.Sel("cancelAllTasks"))
	return rv
}
func (_MLBackgroundTaskClass MLBackgroundTaskClass) CancelTaskWithIdentifier(identifier objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLBackgroundTaskClass.class), objc.Sel("cancelTaskWithIdentifier:"), identifier)
	return rv
}
func (_MLBackgroundTaskClass MLBackgroundTaskClass) ScheduleTask(task objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLBackgroundTaskClass.class), objc.Sel("scheduleTask:"), task)
	return rv
}
func (_MLBackgroundTaskClass MLBackgroundTaskClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLBackgroundTaskClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
func (_MLBackgroundTaskClass MLBackgroundTaskClass) TaskIsScheduledWithIdentifier(identifier objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLBackgroundTaskClass.class), objc.Sel("taskIsScheduledWithIdentifier:"), identifier)
	return rv
}

func (m MLBackgroundTask) TaskIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("taskIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLBackgroundTask) SetTaskIdentifier(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setTaskIdentifier:"), objc.String(value))
}

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
	rv := objc.Send[MLBackgroundTask](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLBackgroundTask.ActivityForScheduling]
//   - [MLBackgroundTask.EncodeWithCoder]
//   - [MLBackgroundTask.TaskIdentifier]
//   - [MLBackgroundTask.SetTaskIdentifier]
//   - [MLBackgroundTask.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask
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
//
// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask
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
func (b MLBackgroundTask) Init() MLBackgroundTask {
	rv := objc.Send[MLBackgroundTask](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MLBackgroundTask) Autorelease() MLBackgroundTask {
	rv := objc.Send[MLBackgroundTask](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBackgroundTask creates a new MLBackgroundTask instance.
func NewMLBackgroundTask() MLBackgroundTask {
	class := getMLBackgroundTaskClass()
	rv := objc.Send[MLBackgroundTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask/initWithCoder:
func NewBackgroundTaskWithCoder(coder objectivec.IObject) MLBackgroundTask {
	instance := getMLBackgroundTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLBackgroundTaskFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask/activityForScheduling
func (b MLBackgroundTask) ActivityForScheduling() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("activityForScheduling"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask/encodeWithCoder:
func (b MLBackgroundTask) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask/initWithCoder:
func (b MLBackgroundTask) InitWithCoder(coder foundation.INSCoder) MLBackgroundTask {
	rv := objc.Send[MLBackgroundTask](b.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask/cancelAllTasks
func (_MLBackgroundTaskClass MLBackgroundTaskClass) CancelAllTasks() bool {
	rv := objc.Send[bool](objc.ID(_MLBackgroundTaskClass.class), objc.Sel("cancelAllTasks"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask/cancelTaskWithIdentifier:
func (_MLBackgroundTaskClass MLBackgroundTaskClass) CancelTaskWithIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLBackgroundTaskClass.class), objc.Sel("cancelTaskWithIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask/scheduleTask:
func (_MLBackgroundTaskClass MLBackgroundTaskClass) ScheduleTask(task objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLBackgroundTaskClass.class), objc.Sel("scheduleTask:"), task)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask/supportsSecureCoding
func (_MLBackgroundTaskClass MLBackgroundTaskClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLBackgroundTaskClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask/taskIsScheduledWithIdentifier:
func (_MLBackgroundTaskClass MLBackgroundTaskClass) TaskIsScheduledWithIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLBackgroundTaskClass.class), objc.Sel("taskIsScheduledWithIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundTask/taskIdentifier
func (b MLBackgroundTask) TaskIdentifier() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("taskIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (b MLBackgroundTask) SetTaskIdentifier(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setTaskIdentifier:"), objc.String(value))
}

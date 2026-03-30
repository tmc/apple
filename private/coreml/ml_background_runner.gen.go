// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLBackgroundRunner] class.
var (
	_MLBackgroundRunnerClass     MLBackgroundRunnerClass
	_MLBackgroundRunnerClassOnce sync.Once
)

func getMLBackgroundRunnerClass() MLBackgroundRunnerClass {
	_MLBackgroundRunnerClassOnce.Do(func() {
		_MLBackgroundRunnerClass = MLBackgroundRunnerClass{class: objc.GetClass("MLBackgroundRunner")}
	})
	return _MLBackgroundRunnerClass
}

// GetMLBackgroundRunnerClass returns the class object for MLBackgroundRunner.
func GetMLBackgroundRunnerClass() MLBackgroundRunnerClass {
	return getMLBackgroundRunnerClass()
}

type MLBackgroundRunnerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLBackgroundRunnerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLBackgroundRunnerClass) Alloc() MLBackgroundRunner {
	rv := objc.Send[MLBackgroundRunner](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLBackgroundRunner.Activity]
//   - [MLBackgroundRunner.SetActivity]
//   - [MLBackgroundRunner.CreateExtensionDataSourceWithInfoKeyConformingToProtocol]
//   - [MLBackgroundRunner.DataSource]
//   - [MLBackgroundRunner.SetDataSource]
//   - [MLBackgroundRunner.DelegateQueue]
//   - [MLBackgroundRunner.SetDelegateQueue]
//   - [MLBackgroundRunner.PrepareForActivity]
//   - [MLBackgroundRunner.ShouldStop]
//   - [MLBackgroundRunner.SetShouldStop]
//   - [MLBackgroundRunner.Start]
//   - [MLBackgroundRunner.Stop]
//   - [MLBackgroundRunner.Task]
//   - [MLBackgroundRunner.SetTask]
//   - [MLBackgroundRunner.WatchdogQueue]
//   - [MLBackgroundRunner.SetWatchdogQueue]
//   - [MLBackgroundRunner.DebugDescription]
//   - [MLBackgroundRunner.Description]
//   - [MLBackgroundRunner.Hash]
//   - [MLBackgroundRunner.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner
type MLBackgroundRunner struct {
	objectivec.Object
}

// MLBackgroundRunnerFromID constructs a [MLBackgroundRunner] from an objc.ID.
func MLBackgroundRunnerFromID(id objc.ID) MLBackgroundRunner {
	return MLBackgroundRunner{objectivec.Object{ID: id}}
}

// Ensure MLBackgroundRunner implements IMLBackgroundRunner.
var _ IMLBackgroundRunner = MLBackgroundRunner{}

// An interface definition for the [MLBackgroundRunner] class.
//
// # Methods
//
//   - [IMLBackgroundRunner.Activity]
//   - [IMLBackgroundRunner.SetActivity]
//   - [IMLBackgroundRunner.CreateExtensionDataSourceWithInfoKeyConformingToProtocol]
//   - [IMLBackgroundRunner.DataSource]
//   - [IMLBackgroundRunner.SetDataSource]
//   - [IMLBackgroundRunner.DelegateQueue]
//   - [IMLBackgroundRunner.SetDelegateQueue]
//   - [IMLBackgroundRunner.PrepareForActivity]
//   - [IMLBackgroundRunner.ShouldStop]
//   - [IMLBackgroundRunner.SetShouldStop]
//   - [IMLBackgroundRunner.Start]
//   - [IMLBackgroundRunner.Stop]
//   - [IMLBackgroundRunner.Task]
//   - [IMLBackgroundRunner.SetTask]
//   - [IMLBackgroundRunner.WatchdogQueue]
//   - [IMLBackgroundRunner.SetWatchdogQueue]
//   - [IMLBackgroundRunner.DebugDescription]
//   - [IMLBackgroundRunner.Description]
//   - [IMLBackgroundRunner.Hash]
//   - [IMLBackgroundRunner.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner
type IMLBackgroundRunner interface {
	objectivec.IObject

	// Topic: Methods

	Activity() unsafe.Pointer
	SetActivity(value unsafe.Pointer)
	CreateExtensionDataSourceWithInfoKeyConformingToProtocol(key objectivec.IObject, protocol_ objectivec.IObject) bool
	DataSource() objectivec.IObject
	SetDataSource(value objectivec.IObject)
	DelegateQueue() objectivec.Object
	SetDelegateQueue(value objectivec.Object)
	PrepareForActivity(activity objectivec.IObject) bool
	ShouldStop() bool
	SetShouldStop(value bool)
	Start() byte
	Stop()
	Task() IMLBackgroundTask
	SetTask(value IMLBackgroundTask)
	WatchdogQueue() objectivec.Object
	SetWatchdogQueue(value objectivec.Object)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (b MLBackgroundRunner) Init() MLBackgroundRunner {
	rv := objc.Send[MLBackgroundRunner](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MLBackgroundRunner) Autorelease() MLBackgroundRunner {
	rv := objc.Send[MLBackgroundRunner](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBackgroundRunner creates a new MLBackgroundRunner instance.
func NewMLBackgroundRunner() MLBackgroundRunner {
	class := getMLBackgroundRunnerClass()
	rv := objc.Send[MLBackgroundRunner](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/createExtensionDataSourceWithInfoKey:conformingToProtocol:
func (b MLBackgroundRunner) CreateExtensionDataSourceWithInfoKeyConformingToProtocol(key objectivec.IObject, protocol_ objectivec.IObject) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("createExtensionDataSourceWithInfoKey:conformingToProtocol:"), key, protocol_)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/prepareForActivity:
func (b MLBackgroundRunner) PrepareForActivity(activity objectivec.IObject) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("prepareForActivity:"), activity)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/start
func (b MLBackgroundRunner) Start() byte {
	rv := objc.Send[byte](b.ID, objc.Sel("start"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/stop
func (b MLBackgroundRunner) Stop() {
	objc.Send[objc.ID](b.ID, objc.Sel("stop"))
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/activity
func (b MLBackgroundRunner) Activity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b.ID, objc.Sel("activity"))
	return rv
}
func (b MLBackgroundRunner) SetActivity(value unsafe.Pointer) {
	objc.Send[struct{}](b.ID, objc.Sel("setActivity:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/dataSource
func (b MLBackgroundRunner) DataSource() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("dataSource"))
	return objectivec.Object{ID: rv}
}
func (b MLBackgroundRunner) SetDataSource(value objectivec.IObject) {
	objc.Send[struct{}](b.ID, objc.Sel("setDataSource:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/debugDescription
func (b MLBackgroundRunner) DebugDescription() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/delegateQueue
func (b MLBackgroundRunner) DelegateQueue() objectivec.Object {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("delegateQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (b MLBackgroundRunner) SetDelegateQueue(value objectivec.Object) {
	objc.Send[struct{}](b.ID, objc.Sel("setDelegateQueue:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/description
func (b MLBackgroundRunner) Description() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/hash
func (b MLBackgroundRunner) Hash() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/shouldStop
func (b MLBackgroundRunner) ShouldStop() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("shouldStop"))
	return rv
}
func (b MLBackgroundRunner) SetShouldStop(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setShouldStop:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/superclass
func (b MLBackgroundRunner) Superclass() objc.Class {
	rv := objc.Send[objc.Class](b.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/task
func (b MLBackgroundRunner) Task() IMLBackgroundTask {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("task"))
	return MLBackgroundTaskFromID(objc.ID(rv))
}
func (b MLBackgroundRunner) SetTask(value IMLBackgroundTask) {
	objc.Send[struct{}](b.ID, objc.Sel("setTask:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundRunner/watchdogQueue
func (b MLBackgroundRunner) WatchdogQueue() objectivec.Object {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("watchdogQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (b MLBackgroundRunner) SetWatchdogQueue(value objectivec.Object) {
	objc.Send[struct{}](b.ID, objc.Sel("setWatchdogQueue:"), value)
}

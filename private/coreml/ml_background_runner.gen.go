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
	rv := objc.SendIfResponds[MLBackgroundRunner](objc.ID(mc.class), objc.Sel("alloc"))
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
type IMLBackgroundRunner interface {
	objectivec.IObject

	// Topic: Methods

	Activity() unsafe.Pointer
	SetActivity(value unsafe.Pointer)
	CreateExtensionDataSourceWithInfoKeyConformingToProtocol(key objectivec.IObject, protocol_ objectivec.IObject) bool
	DataSource() unsafe.Pointer
	SetDataSource(value unsafe.Pointer)
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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLBackgroundRunner) Init() MLBackgroundRunner {
	rv := objc.SendIfResponds[MLBackgroundRunner](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLBackgroundRunner) Autorelease() MLBackgroundRunner {
	rv := objc.SendIfResponds[MLBackgroundRunner](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBackgroundRunner creates a new MLBackgroundRunner instance.
func NewMLBackgroundRunner() MLBackgroundRunner {
	class := getMLBackgroundRunnerClass()
	rv := objc.SendIfResponds[MLBackgroundRunner](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLBackgroundRunner) CreateExtensionDataSourceWithInfoKeyConformingToProtocol(key objectivec.IObject, protocol_ objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("createExtensionDataSourceWithInfoKey:conformingToProtocol:"), key, protocol_)
	return rv
}
func (m MLBackgroundRunner) PrepareForActivity(activity objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("prepareForActivity:"), activity)
	return rv
}
func (m MLBackgroundRunner) Start() byte {
	rv := objc.SendIfResponds[byte](m.ID, objc.Sel("start"))
	return rv
}
func (m MLBackgroundRunner) Stop() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("stop"))
}

func (m MLBackgroundRunner) Activity() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("activity"))
	return rv
}
func (m MLBackgroundRunner) SetActivity(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setActivity:"), value)
}
func (m MLBackgroundRunner) DataSource() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("dataSource"))
	return rv
}
func (m MLBackgroundRunner) SetDataSource(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setDataSource:"), value)
}
func (m MLBackgroundRunner) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLBackgroundRunner) DelegateQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("delegateQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLBackgroundRunner) SetDelegateQueue(value objectivec.Object) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setDelegateQueue:"), value)
}
func (m MLBackgroundRunner) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLBackgroundRunner) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLBackgroundRunner) ShouldStop() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shouldStop"))
	return rv
}
func (m MLBackgroundRunner) SetShouldStop(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setShouldStop:"), value)
}
func (m MLBackgroundRunner) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (m MLBackgroundRunner) Task() IMLBackgroundTask {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("task"))
	return MLBackgroundTaskFromID(objc.ID(rv))
}
func (m MLBackgroundRunner) SetTask(value IMLBackgroundTask) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setTask:"), value)
}
func (m MLBackgroundRunner) WatchdogQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("watchdogQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLBackgroundRunner) SetWatchdogQueue(value objectivec.Object) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setWatchdogQueue:"), value)
}

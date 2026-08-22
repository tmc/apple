// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXSpecialKeyEventProcessor] class.
var (
	_CPXSpecialKeyEventProcessorClass     CPXSpecialKeyEventProcessorClass
	_CPXSpecialKeyEventProcessorClassOnce sync.Once
)

func getCPXSpecialKeyEventProcessorClass() CPXSpecialKeyEventProcessorClass {
	_CPXSpecialKeyEventProcessorClassOnce.Do(func() {
		_CPXSpecialKeyEventProcessorClass = CPXSpecialKeyEventProcessorClass{class: objc.GetClass("CPXSpecialKeyEventProcessor")}
	})
	return _CPXSpecialKeyEventProcessorClass
}

// GetCPXSpecialKeyEventProcessorClass returns the class object for CPXSpecialKeyEventProcessor.
func GetCPXSpecialKeyEventProcessorClass() CPXSpecialKeyEventProcessorClass {
	return getCPXSpecialKeyEventProcessorClass()
}

type CPXSpecialKeyEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXSpecialKeyEventProcessorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXSpecialKeyEventProcessorClass) Alloc() CPXSpecialKeyEventProcessor {
	rv := objc.SendIfResponds[CPXSpecialKeyEventProcessor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXSpecialKeyEventProcessor.CleanupForProcessDeath]
//   - [CPXSpecialKeyEventProcessor.ClearEventState]
//   - [CPXSpecialKeyEventProcessor.ExitSpecialKeyModeForProcess]
//   - [CPXSpecialKeyEventProcessor.HotKeyChanged]
//   - [CPXSpecialKeyEventProcessor.ProcessEventContextDispatcher]
//   - [CPXSpecialKeyEventProcessor.ProcessHotKeyEventHotKeyIDIsDownContextDispatcher]
//   - [CPXSpecialKeyEventProcessor.RegisterSpecialKeyConnectionForProcess]
//   - [CPXSpecialKeyEventProcessor.UnregisterSpecialKeyForProcess]
//   - [CPXSpecialKeyEventProcessor.InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter]
//   - [CPXSpecialKeyEventProcessor.InitWithSession]
//   - [CPXSpecialKeyEventProcessor.DebugDescription]
//   - [CPXSpecialKeyEventProcessor.Description]
//   - [CPXSpecialKeyEventProcessor.Hash]
//   - [CPXSpecialKeyEventProcessor.Superclass]
type CPXSpecialKeyEventProcessor struct {
	objectivec.Object
}

// CPXSpecialKeyEventProcessorFromID constructs a [CPXSpecialKeyEventProcessor] from an objc.ID.
func CPXSpecialKeyEventProcessorFromID(id objc.ID) CPXSpecialKeyEventProcessor {
	return CPXSpecialKeyEventProcessor{objectivec.Object{ID: id}}
}

// Ensure CPXSpecialKeyEventProcessor implements ICPXSpecialKeyEventProcessor.
var _ ICPXSpecialKeyEventProcessor = CPXSpecialKeyEventProcessor{}

// An interface definition for the [CPXSpecialKeyEventProcessor] class.
//
// # Methods
//
//   - [ICPXSpecialKeyEventProcessor.CleanupForProcessDeath]
//   - [ICPXSpecialKeyEventProcessor.ClearEventState]
//   - [ICPXSpecialKeyEventProcessor.ExitSpecialKeyModeForProcess]
//   - [ICPXSpecialKeyEventProcessor.HotKeyChanged]
//   - [ICPXSpecialKeyEventProcessor.ProcessEventContextDispatcher]
//   - [ICPXSpecialKeyEventProcessor.ProcessHotKeyEventHotKeyIDIsDownContextDispatcher]
//   - [ICPXSpecialKeyEventProcessor.RegisterSpecialKeyConnectionForProcess]
//   - [ICPXSpecialKeyEventProcessor.UnregisterSpecialKeyForProcess]
//   - [ICPXSpecialKeyEventProcessor.InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter]
//   - [ICPXSpecialKeyEventProcessor.InitWithSession]
//   - [ICPXSpecialKeyEventProcessor.DebugDescription]
//   - [ICPXSpecialKeyEventProcessor.Description]
//   - [ICPXSpecialKeyEventProcessor.Hash]
//   - [ICPXSpecialKeyEventProcessor.Superclass]
type ICPXSpecialKeyEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	CleanupForProcessDeath(death *CPSProcessRec)
	ClearEventState()
	ExitSpecialKeyModeForProcess(mode uint32, process *CPSProcessRec)
	HotKeyChanged(changed unsafe.Pointer)
	ProcessEventContextDispatcher(event *SLSEventRecord, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64
	ProcessHotKeyEventHotKeyIDIsDownContextDispatcher(event *SLSEventRecord, id uint64, down bool, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64
	RegisterSpecialKeyConnectionForProcess(key uint32, connection *CGXConnection, process *CPSProcessRec) int
	UnregisterSpecialKeyForProcess(key uint32, process *CPSProcessRec) int
	InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXSpecialKeyEventProcessor
	InitWithSession(session *CGXSession) CPXSpecialKeyEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (c CPXSpecialKeyEventProcessor) Init() CPXSpecialKeyEventProcessor {
	rv := objc.SendIfResponds[CPXSpecialKeyEventProcessor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXSpecialKeyEventProcessor) Autorelease() CPXSpecialKeyEventProcessor {
	rv := objc.SendIfResponds[CPXSpecialKeyEventProcessor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXSpecialKeyEventProcessor creates a new CPXSpecialKeyEventProcessor instance.
func NewCPXSpecialKeyEventProcessor() CPXSpecialKeyEventProcessor {
	class := getCPXSpecialKeyEventProcessorClass()
	rv := objc.SendIfResponds[CPXSpecialKeyEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCPXSpecialKeyEventProcessorWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXSpecialKeyEventProcessor {
	instance := getCPXSpecialKeyEventProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:"), manager, manager2, registry, scheduler, center)
	return CPXSpecialKeyEventProcessorFromID(rv)
}

func NewCPXSpecialKeyEventProcessorWithSession(session *CGXSession) CPXSpecialKeyEventProcessor {
	instance := getCPXSpecialKeyEventProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSession:"), unsafe.Pointer(session))
	return CPXSpecialKeyEventProcessorFromID(rv)
}

func (c CPXSpecialKeyEventProcessor) CleanupForProcessDeath(death *CPSProcessRec) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("cleanupForProcessDeath:"), unsafe.Pointer(death))
}
func (c CPXSpecialKeyEventProcessor) ClearEventState() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("clearEventState"))
}
func (c CPXSpecialKeyEventProcessor) ExitSpecialKeyModeForProcess(mode uint32, process *CPSProcessRec) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("exitSpecialKeyMode:forProcess:"), mode, unsafe.Pointer(process))
}
func (c CPXSpecialKeyEventProcessor) HotKeyChanged(changed unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("hotKeyChanged:"), changed)
}
func (c CPXSpecialKeyEventProcessor) ProcessEventContextDispatcher(event *SLSEventRecord, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](c.ID, objc.Sel("processEvent:context:dispatcher:"), unsafe.Pointer(event), unsafe.Pointer(context), dispatcher)
	return rv
}
func (c CPXSpecialKeyEventProcessor) ProcessHotKeyEventHotKeyIDIsDownContextDispatcher(event *SLSEventRecord, id uint64, down bool, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](c.ID, objc.Sel("processHotKeyEvent:hotKeyID:isDown:context:dispatcher:"), unsafe.Pointer(event), id, down, unsafe.Pointer(context), dispatcher)
	return rv
}
func (c CPXSpecialKeyEventProcessor) RegisterSpecialKeyConnectionForProcess(key uint32, connection *CGXConnection, process *CPSProcessRec) int {
	rv := objc.SendIfResponds[int](c.ID, objc.Sel("registerSpecialKey:connection:forProcess:"), key, unsafe.Pointer(connection), unsafe.Pointer(process))
	return rv
}
func (c CPXSpecialKeyEventProcessor) UnregisterSpecialKeyForProcess(key uint32, process *CPSProcessRec) int {
	rv := objc.SendIfResponds[int](c.ID, objc.Sel("unregisterSpecialKey:forProcess:"), key, unsafe.Pointer(process))
	return rv
}
func (c CPXSpecialKeyEventProcessor) InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXSpecialKeyEventProcessor {
	rv := objc.SendIfResponds[CPXSpecialKeyEventProcessor](c.ID, objc.Sel("initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:"), manager, manager2, registry, scheduler, center)
	return rv
}
func (c CPXSpecialKeyEventProcessor) InitWithSession(session *CGXSession) CPXSpecialKeyEventProcessor {
	rv := objc.SendIfResponds[CPXSpecialKeyEventProcessor](c.ID, objc.Sel("initWithSession:"), unsafe.Pointer(session))
	return rv
}

func (c CPXSpecialKeyEventProcessor) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXSpecialKeyEventProcessor) Description() string {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXSpecialKeyEventProcessor) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](c.ID, objc.Sel("hash"))
	return rv
}
func (c CPXSpecialKeyEventProcessor) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](c.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

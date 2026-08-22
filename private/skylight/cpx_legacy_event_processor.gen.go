// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXLegacyEventProcessor] class.
var (
	_CPXLegacyEventProcessorClass     CPXLegacyEventProcessorClass
	_CPXLegacyEventProcessorClassOnce sync.Once
)

func getCPXLegacyEventProcessorClass() CPXLegacyEventProcessorClass {
	_CPXLegacyEventProcessorClassOnce.Do(func() {
		_CPXLegacyEventProcessorClass = CPXLegacyEventProcessorClass{class: objc.GetClass("CPXLegacyEventProcessor")}
	})
	return _CPXLegacyEventProcessorClass
}

// GetCPXLegacyEventProcessorClass returns the class object for CPXLegacyEventProcessor.
func GetCPXLegacyEventProcessorClass() CPXLegacyEventProcessorClass {
	return getCPXLegacyEventProcessorClass()
}

type CPXLegacyEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXLegacyEventProcessorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXLegacyEventProcessorClass) Alloc() CPXLegacyEventProcessor {
	rv := objc.SendIfResponds[CPXLegacyEventProcessor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXLegacyEventProcessor.CleanupForProcessDeath]
//   - [CPXLegacyEventProcessor.ClearEventState]
//   - [CPXLegacyEventProcessor.ExitSpecialKeyModeForProcess]
//   - [CPXLegacyEventProcessor.HotKeyChanged]
//   - [CPXLegacyEventProcessor.ProcessEventContextDispatcher]
//   - [CPXLegacyEventProcessor.ProcessHotKeyEventHotKeyIDIsDownContextDispatcher]
//   - [CPXLegacyEventProcessor.RegisterSpecialKeyConnectionForProcess]
//   - [CPXLegacyEventProcessor.UnregisterSpecialKeyForProcess]
//   - [CPXLegacyEventProcessor.InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter]
//   - [CPXLegacyEventProcessor.InitWithSession]
//   - [CPXLegacyEventProcessor.DebugDescription]
//   - [CPXLegacyEventProcessor.Description]
//   - [CPXLegacyEventProcessor.Hash]
//   - [CPXLegacyEventProcessor.Superclass]
type CPXLegacyEventProcessor struct {
	objectivec.Object
}

// CPXLegacyEventProcessorFromID constructs a [CPXLegacyEventProcessor] from an objc.ID.
func CPXLegacyEventProcessorFromID(id objc.ID) CPXLegacyEventProcessor {
	return CPXLegacyEventProcessor{objectivec.Object{ID: id}}
}

// Ensure CPXLegacyEventProcessor implements ICPXLegacyEventProcessor.
var _ ICPXLegacyEventProcessor = CPXLegacyEventProcessor{}

// An interface definition for the [CPXLegacyEventProcessor] class.
//
// # Methods
//
//   - [ICPXLegacyEventProcessor.CleanupForProcessDeath]
//   - [ICPXLegacyEventProcessor.ClearEventState]
//   - [ICPXLegacyEventProcessor.ExitSpecialKeyModeForProcess]
//   - [ICPXLegacyEventProcessor.HotKeyChanged]
//   - [ICPXLegacyEventProcessor.ProcessEventContextDispatcher]
//   - [ICPXLegacyEventProcessor.ProcessHotKeyEventHotKeyIDIsDownContextDispatcher]
//   - [ICPXLegacyEventProcessor.RegisterSpecialKeyConnectionForProcess]
//   - [ICPXLegacyEventProcessor.UnregisterSpecialKeyForProcess]
//   - [ICPXLegacyEventProcessor.InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter]
//   - [ICPXLegacyEventProcessor.InitWithSession]
//   - [ICPXLegacyEventProcessor.DebugDescription]
//   - [ICPXLegacyEventProcessor.Description]
//   - [ICPXLegacyEventProcessor.Hash]
//   - [ICPXLegacyEventProcessor.Superclass]
type ICPXLegacyEventProcessor interface {
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
	InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXLegacyEventProcessor
	InitWithSession(session *CGXSession) CPXLegacyEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (c CPXLegacyEventProcessor) Init() CPXLegacyEventProcessor {
	rv := objc.SendIfResponds[CPXLegacyEventProcessor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXLegacyEventProcessor) Autorelease() CPXLegacyEventProcessor {
	rv := objc.SendIfResponds[CPXLegacyEventProcessor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXLegacyEventProcessor creates a new CPXLegacyEventProcessor instance.
func NewCPXLegacyEventProcessor() CPXLegacyEventProcessor {
	class := getCPXLegacyEventProcessorClass()
	rv := objc.SendIfResponds[CPXLegacyEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCPXLegacyEventProcessorWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXLegacyEventProcessor {
	instance := getCPXLegacyEventProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:"), manager, manager2, registry, scheduler, center)
	return CPXLegacyEventProcessorFromID(rv)
}

func NewCPXLegacyEventProcessorWithSession(session *CGXSession) CPXLegacyEventProcessor {
	instance := getCPXLegacyEventProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSession:"), unsafe.Pointer(session))
	return CPXLegacyEventProcessorFromID(rv)
}

func (c CPXLegacyEventProcessor) CleanupForProcessDeath(death *CPSProcessRec) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("cleanupForProcessDeath:"), unsafe.Pointer(death))
}
func (c CPXLegacyEventProcessor) ClearEventState() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("clearEventState"))
}
func (c CPXLegacyEventProcessor) ExitSpecialKeyModeForProcess(mode uint32, process *CPSProcessRec) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("exitSpecialKeyMode:forProcess:"), mode, unsafe.Pointer(process))
}
func (c CPXLegacyEventProcessor) HotKeyChanged(changed unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("hotKeyChanged:"), changed)
}
func (c CPXLegacyEventProcessor) ProcessEventContextDispatcher(event *SLSEventRecord, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](c.ID, objc.Sel("processEvent:context:dispatcher:"), unsafe.Pointer(event), unsafe.Pointer(context), dispatcher)
	return rv
}
func (c CPXLegacyEventProcessor) ProcessHotKeyEventHotKeyIDIsDownContextDispatcher(event *SLSEventRecord, id uint64, down bool, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](c.ID, objc.Sel("processHotKeyEvent:hotKeyID:isDown:context:dispatcher:"), unsafe.Pointer(event), id, down, unsafe.Pointer(context), dispatcher)
	return rv
}
func (c CPXLegacyEventProcessor) RegisterSpecialKeyConnectionForProcess(key uint32, connection *CGXConnection, process *CPSProcessRec) int {
	rv := objc.SendIfResponds[int](c.ID, objc.Sel("registerSpecialKey:connection:forProcess:"), key, unsafe.Pointer(connection), unsafe.Pointer(process))
	return rv
}
func (c CPXLegacyEventProcessor) UnregisterSpecialKeyForProcess(key uint32, process *CPSProcessRec) int {
	rv := objc.SendIfResponds[int](c.ID, objc.Sel("unregisterSpecialKey:forProcess:"), key, unsafe.Pointer(process))
	return rv
}
func (c CPXLegacyEventProcessor) InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXLegacyEventProcessor {
	rv := objc.SendIfResponds[CPXLegacyEventProcessor](c.ID, objc.Sel("initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:"), manager, manager2, registry, scheduler, center)
	return rv
}
func (c CPXLegacyEventProcessor) InitWithSession(session *CGXSession) CPXLegacyEventProcessor {
	rv := objc.SendIfResponds[CPXLegacyEventProcessor](c.ID, objc.Sel("initWithSession:"), unsafe.Pointer(session))
	return rv
}

func (c CPXLegacyEventProcessor) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXLegacyEventProcessor) Description() string {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXLegacyEventProcessor) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](c.ID, objc.Sel("hash"))
	return rv
}
func (c CPXLegacyEventProcessor) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](c.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

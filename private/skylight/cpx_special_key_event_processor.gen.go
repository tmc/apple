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
	rv := objc.Send[CPXSpecialKeyEventProcessor](objc.ID(cc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor
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
//
// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor
type ICPXSpecialKeyEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	CleanupForProcessDeath(death *CPSProcessRecRef)
	ClearEventState()
	ExitSpecialKeyModeForProcess(mode uint32, process *CPSProcessRecRef)
	HotKeyChanged(changed objectivec.IObject)
	ProcessEventContextDispatcher(event *SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64
	ProcessHotKeyEventHotKeyIDIsDownContextDispatcher(event *SLSEventRecordRef, id uint64, down bool, context unsafe.Pointer, dispatcher objectivec.IObject) int64
	RegisterSpecialKeyConnectionForProcess(key uint32, connection unsafe.Pointer, process *CPSProcessRecRef) int
	UnregisterSpecialKeyForProcess(key uint32, process *CPSProcessRecRef) int
	InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXSpecialKeyEventProcessor
	InitWithSession(session unsafe.Pointer) CPXSpecialKeyEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXSpecialKeyEventProcessor) Init() CPXSpecialKeyEventProcessor {
	rv := objc.Send[CPXSpecialKeyEventProcessor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXSpecialKeyEventProcessor) Autorelease() CPXSpecialKeyEventProcessor {
	rv := objc.Send[CPXSpecialKeyEventProcessor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXSpecialKeyEventProcessor creates a new CPXSpecialKeyEventProcessor instance.
func NewCPXSpecialKeyEventProcessor() CPXSpecialKeyEventProcessor {
	class := getCPXSpecialKeyEventProcessorClass()
	rv := objc.Send[CPXSpecialKeyEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:
func NewCPXSpecialKeyEventProcessorWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXSpecialKeyEventProcessor {
	instance := getCPXSpecialKeyEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:"), manager, manager2, registry, scheduler, center)
	return CPXSpecialKeyEventProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/initWithSession:
func NewCPXSpecialKeyEventProcessorWithSession(session unsafe.Pointer) CPXSpecialKeyEventProcessor {
	instance := getCPXSpecialKeyEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return CPXSpecialKeyEventProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/cleanupForProcessDeath:
func (c CPXSpecialKeyEventProcessor) CleanupForProcessDeath(death *CPSProcessRecRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("cleanupForProcessDeath:"), death)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/clearEventState
func (c CPXSpecialKeyEventProcessor) ClearEventState() {
	objc.Send[objc.ID](c.ID, objc.Sel("clearEventState"))
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/exitSpecialKeyMode:forProcess:
func (c CPXSpecialKeyEventProcessor) ExitSpecialKeyModeForProcess(mode uint32, process *CPSProcessRecRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("exitSpecialKeyMode:forProcess:"), mode, process)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/hotKeyChanged:
func (c CPXSpecialKeyEventProcessor) HotKeyChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("hotKeyChanged:"), changed)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/processEvent:context:dispatcher:
func (c CPXSpecialKeyEventProcessor) ProcessEventContextDispatcher(event *SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("processEvent:context:dispatcher:"), event, context, dispatcher)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/processHotKeyEvent:hotKeyID:isDown:context:dispatcher:
func (c CPXSpecialKeyEventProcessor) ProcessHotKeyEventHotKeyIDIsDownContextDispatcher(event *SLSEventRecordRef, id uint64, down bool, context unsafe.Pointer, dispatcher objectivec.IObject) int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("processHotKeyEvent:hotKeyID:isDown:context:dispatcher:"), event, id, down, context, dispatcher)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/registerSpecialKey:connection:forProcess:
func (c CPXSpecialKeyEventProcessor) RegisterSpecialKeyConnectionForProcess(key uint32, connection unsafe.Pointer, process *CPSProcessRecRef) int {
	rv := objc.Send[int](c.ID, objc.Sel("registerSpecialKey:connection:forProcess:"), key, connection, process)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/unregisterSpecialKey:forProcess:
func (c CPXSpecialKeyEventProcessor) UnregisterSpecialKeyForProcess(key uint32, process *CPSProcessRecRef) int {
	rv := objc.Send[int](c.ID, objc.Sel("unregisterSpecialKey:forProcess:"), key, process)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:
func (c CPXSpecialKeyEventProcessor) InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXSpecialKeyEventProcessor {
	rv := objc.Send[CPXSpecialKeyEventProcessor](c.ID, objc.Sel("initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:"), manager, manager2, registry, scheduler, center)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/initWithSession:
func (c CPXSpecialKeyEventProcessor) InitWithSession(session unsafe.Pointer) CPXSpecialKeyEventProcessor {
	rv := objc.Send[CPXSpecialKeyEventProcessor](c.ID, objc.Sel("initWithSession:"), session)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/debugDescription
func (c CPXSpecialKeyEventProcessor) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/description
func (c CPXSpecialKeyEventProcessor) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/hash
func (c CPXSpecialKeyEventProcessor) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyEventProcessor/superclass
func (c CPXSpecialKeyEventProcessor) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

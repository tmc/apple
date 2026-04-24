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
	rv := objc.Send[CPXLegacyEventProcessor](objc.ID(cc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor
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
//
// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor
type ICPXLegacyEventProcessor interface {
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
	InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXLegacyEventProcessor
	InitWithSession(session unsafe.Pointer) CPXLegacyEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXLegacyEventProcessor) Init() CPXLegacyEventProcessor {
	rv := objc.Send[CPXLegacyEventProcessor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXLegacyEventProcessor) Autorelease() CPXLegacyEventProcessor {
	rv := objc.Send[CPXLegacyEventProcessor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXLegacyEventProcessor creates a new CPXLegacyEventProcessor instance.
func NewCPXLegacyEventProcessor() CPXLegacyEventProcessor {
	class := getCPXLegacyEventProcessorClass()
	rv := objc.Send[CPXLegacyEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:
func NewCPXLegacyEventProcessorWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXLegacyEventProcessor {
	instance := getCPXLegacyEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:"), manager, manager2, registry, scheduler, center)
	return CPXLegacyEventProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/initWithSession:
func NewCPXLegacyEventProcessorWithSession(session unsafe.Pointer) CPXLegacyEventProcessor {
	instance := getCPXLegacyEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return CPXLegacyEventProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/cleanupForProcessDeath:
func (c CPXLegacyEventProcessor) CleanupForProcessDeath(death *CPSProcessRecRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("cleanupForProcessDeath:"), death)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/clearEventState
func (c CPXLegacyEventProcessor) ClearEventState() {
	objc.Send[objc.ID](c.ID, objc.Sel("clearEventState"))
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/exitSpecialKeyMode:forProcess:
func (c CPXLegacyEventProcessor) ExitSpecialKeyModeForProcess(mode uint32, process *CPSProcessRecRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("exitSpecialKeyMode:forProcess:"), mode, process)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/hotKeyChanged:
func (c CPXLegacyEventProcessor) HotKeyChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("hotKeyChanged:"), changed)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/processEvent:context:dispatcher:
func (c CPXLegacyEventProcessor) ProcessEventContextDispatcher(event *SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("processEvent:context:dispatcher:"), event, context, dispatcher)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/processHotKeyEvent:hotKeyID:isDown:context:dispatcher:
func (c CPXLegacyEventProcessor) ProcessHotKeyEventHotKeyIDIsDownContextDispatcher(event *SLSEventRecordRef, id uint64, down bool, context unsafe.Pointer, dispatcher objectivec.IObject) int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("processHotKeyEvent:hotKeyID:isDown:context:dispatcher:"), event, id, down, context, dispatcher)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/registerSpecialKey:connection:forProcess:
func (c CPXLegacyEventProcessor) RegisterSpecialKeyConnectionForProcess(key uint32, connection unsafe.Pointer, process *CPSProcessRecRef) int {
	rv := objc.Send[int](c.ID, objc.Sel("registerSpecialKey:connection:forProcess:"), key, connection, process)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/unregisterSpecialKey:forProcess:
func (c CPXLegacyEventProcessor) UnregisterSpecialKeyForProcess(key uint32, process *CPSProcessRecRef) int {
	rv := objc.Send[int](c.ID, objc.Sel("unregisterSpecialKey:forProcess:"), key, process)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:
func (c CPXLegacyEventProcessor) InitWithProcessManagerFocusManagerSymbolicHotKeyRegistryCallbackSchedulerNotificationCenter(manager objectivec.IObject, manager2 objectivec.IObject, registry objectivec.IObject, scheduler objectivec.IObject, center objectivec.IObject) CPXLegacyEventProcessor {
	rv := objc.Send[CPXLegacyEventProcessor](c.ID, objc.Sel("initWithProcessManager:focusManager:symbolicHotKeyRegistry:callbackScheduler:notificationCenter:"), manager, manager2, registry, scheduler, center)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/initWithSession:
func (c CPXLegacyEventProcessor) InitWithSession(session unsafe.Pointer) CPXLegacyEventProcessor {
	rv := objc.Send[CPXLegacyEventProcessor](c.ID, objc.Sel("initWithSession:"), session)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/debugDescription
func (c CPXLegacyEventProcessor) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/description
func (c CPXLegacyEventProcessor) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/hash
func (c CPXLegacyEventProcessor) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXLegacyEventProcessor/superclass
func (c CPXLegacyEventProcessor) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

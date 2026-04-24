// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXKeyboardEventProcessor] class.
var (
	_CPXKeyboardEventProcessorClass     CPXKeyboardEventProcessorClass
	_CPXKeyboardEventProcessorClassOnce sync.Once
)

func getCPXKeyboardEventProcessorClass() CPXKeyboardEventProcessorClass {
	_CPXKeyboardEventProcessorClassOnce.Do(func() {
		_CPXKeyboardEventProcessorClass = CPXKeyboardEventProcessorClass{class: objc.GetClass("CPXKeyboardEventProcessor")}
	})
	return _CPXKeyboardEventProcessorClass
}

// GetCPXKeyboardEventProcessorClass returns the class object for CPXKeyboardEventProcessor.
func GetCPXKeyboardEventProcessorClass() CPXKeyboardEventProcessorClass {
	return getCPXKeyboardEventProcessorClass()
}

type CPXKeyboardEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXKeyboardEventProcessorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXKeyboardEventProcessorClass) Alloc() CPXKeyboardEventProcessor {
	rv := objc.Send[CPXKeyboardEventProcessor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXKeyboardEventProcessor.ClearEventState]
//   - [CPXKeyboardEventProcessor.ProcessEventContextDispatcher]
//   - [CPXKeyboardEventProcessor.InitWithDeliveryManagerSpecialKeyEventProcessorProcessManagerDestinationGeneratorNotificationCenterKeyEventTracker]
//   - [CPXKeyboardEventProcessor.InitWithSessionSpecialKeyEventProcessor]
//   - [CPXKeyboardEventProcessor.DebugDescription]
//   - [CPXKeyboardEventProcessor.Description]
//   - [CPXKeyboardEventProcessor.Hash]
//   - [CPXKeyboardEventProcessor.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor
type CPXKeyboardEventProcessor struct {
	objectivec.Object
}

// CPXKeyboardEventProcessorFromID constructs a [CPXKeyboardEventProcessor] from an objc.ID.
func CPXKeyboardEventProcessorFromID(id objc.ID) CPXKeyboardEventProcessor {
	return CPXKeyboardEventProcessor{objectivec.Object{ID: id}}
}

// Ensure CPXKeyboardEventProcessor implements ICPXKeyboardEventProcessor.
var _ ICPXKeyboardEventProcessor = CPXKeyboardEventProcessor{}

// An interface definition for the [CPXKeyboardEventProcessor] class.
//
// # Methods
//
//   - [ICPXKeyboardEventProcessor.ClearEventState]
//   - [ICPXKeyboardEventProcessor.ProcessEventContextDispatcher]
//   - [ICPXKeyboardEventProcessor.InitWithDeliveryManagerSpecialKeyEventProcessorProcessManagerDestinationGeneratorNotificationCenterKeyEventTracker]
//   - [ICPXKeyboardEventProcessor.InitWithSessionSpecialKeyEventProcessor]
//   - [ICPXKeyboardEventProcessor.DebugDescription]
//   - [ICPXKeyboardEventProcessor.Description]
//   - [ICPXKeyboardEventProcessor.Hash]
//   - [ICPXKeyboardEventProcessor.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor
type ICPXKeyboardEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	ClearEventState()
	ProcessEventContextDispatcher(event *SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64
	InitWithDeliveryManagerSpecialKeyEventProcessorProcessManagerDestinationGeneratorNotificationCenterKeyEventTracker(manager objectivec.IObject, processor objectivec.IObject, manager2 objectivec.IObject, generator objectivec.IObject, center objectivec.IObject, tracker objectivec.IObject) CPXKeyboardEventProcessor
	InitWithSessionSpecialKeyEventProcessor(session unsafe.Pointer, processor objectivec.IObject) CPXKeyboardEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXKeyboardEventProcessor) Init() CPXKeyboardEventProcessor {
	rv := objc.Send[CPXKeyboardEventProcessor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXKeyboardEventProcessor) Autorelease() CPXKeyboardEventProcessor {
	rv := objc.Send[CPXKeyboardEventProcessor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXKeyboardEventProcessor creates a new CPXKeyboardEventProcessor instance.
func NewCPXKeyboardEventProcessor() CPXKeyboardEventProcessor {
	class := getCPXKeyboardEventProcessorClass()
	rv := objc.Send[CPXKeyboardEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor/initWithDeliveryManager:specialKeyEventProcessor:processManager:destinationGenerator:notificationCenter:keyEventTracker:
func NewCPXKeyboardEventProcessorWithDeliveryManagerSpecialKeyEventProcessorProcessManagerDestinationGeneratorNotificationCenterKeyEventTracker(manager objectivec.IObject, processor objectivec.IObject, manager2 objectivec.IObject, generator objectivec.IObject, center objectivec.IObject, tracker objectivec.IObject) CPXKeyboardEventProcessor {
	instance := getCPXKeyboardEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDeliveryManager:specialKeyEventProcessor:processManager:destinationGenerator:notificationCenter:keyEventTracker:"), manager, processor, manager2, generator, center, tracker)
	return CPXKeyboardEventProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor/initWithSession:specialKeyEventProcessor:
func NewCPXKeyboardEventProcessorWithSessionSpecialKeyEventProcessor(session unsafe.Pointer, processor objectivec.IObject) CPXKeyboardEventProcessor {
	instance := getCPXKeyboardEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:specialKeyEventProcessor:"), session, processor)
	return CPXKeyboardEventProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor/clearEventState
func (c CPXKeyboardEventProcessor) ClearEventState() {
	objc.Send[objc.ID](c.ID, objc.Sel("clearEventState"))
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor/processEvent:context:dispatcher:
func (c CPXKeyboardEventProcessor) ProcessEventContextDispatcher(event *SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("processEvent:context:dispatcher:"), event, context, dispatcher)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor/initWithDeliveryManager:specialKeyEventProcessor:processManager:destinationGenerator:notificationCenter:keyEventTracker:
func (c CPXKeyboardEventProcessor) InitWithDeliveryManagerSpecialKeyEventProcessorProcessManagerDestinationGeneratorNotificationCenterKeyEventTracker(manager objectivec.IObject, processor objectivec.IObject, manager2 objectivec.IObject, generator objectivec.IObject, center objectivec.IObject, tracker objectivec.IObject) CPXKeyboardEventProcessor {
	rv := objc.Send[CPXKeyboardEventProcessor](c.ID, objc.Sel("initWithDeliveryManager:specialKeyEventProcessor:processManager:destinationGenerator:notificationCenter:keyEventTracker:"), manager, processor, manager2, generator, center, tracker)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor/initWithSession:specialKeyEventProcessor:
func (c CPXKeyboardEventProcessor) InitWithSessionSpecialKeyEventProcessor(session unsafe.Pointer, processor objectivec.IObject) CPXKeyboardEventProcessor {
	rv := objc.Send[CPXKeyboardEventProcessor](c.ID, objc.Sel("initWithSession:specialKeyEventProcessor:"), session, processor)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor/debugDescription
func (c CPXKeyboardEventProcessor) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor/description
func (c CPXKeyboardEventProcessor) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor/hash
func (c CPXKeyboardEventProcessor) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessor/superclass
func (c CPXKeyboardEventProcessor) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

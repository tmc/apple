// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

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
type ICPXKeyboardEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	ClearEventState()
	ProcessEventContextDispatcher(event *SLSEventRecord, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64
	InitWithDeliveryManagerSpecialKeyEventProcessorProcessManagerDestinationGeneratorNotificationCenterKeyEventTracker(manager objectivec.IObject, processor objectivec.IObject, manager2 objectivec.IObject, generator objectivec.IObject, center objectivec.IObject, tracker objectivec.IObject) CPXKeyboardEventProcessor
	InitWithSessionSpecialKeyEventProcessor(session *CGXSession, processor objectivec.IObject) CPXKeyboardEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
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

func NewCPXKeyboardEventProcessorWithDeliveryManagerSpecialKeyEventProcessorProcessManagerDestinationGeneratorNotificationCenterKeyEventTracker(manager objectivec.IObject, processor objectivec.IObject, manager2 objectivec.IObject, generator objectivec.IObject, center objectivec.IObject, tracker objectivec.IObject) CPXKeyboardEventProcessor {
	instance := getCPXKeyboardEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDeliveryManager:specialKeyEventProcessor:processManager:destinationGenerator:notificationCenter:keyEventTracker:"), manager, processor, manager2, generator, center, tracker)
	return CPXKeyboardEventProcessorFromID(rv)
}

func NewCPXKeyboardEventProcessorWithSessionSpecialKeyEventProcessor(session *CGXSession, processor objectivec.IObject) CPXKeyboardEventProcessor {
	instance := getCPXKeyboardEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:specialKeyEventProcessor:"), session, processor)
	return CPXKeyboardEventProcessorFromID(rv)
}

func (c CPXKeyboardEventProcessor) ClearEventState() {
	objc.Send[objc.ID](c.ID, objc.Sel("clearEventState"))
}
func (c CPXKeyboardEventProcessor) ProcessEventContextDispatcher(event *SLSEventRecord, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("processEvent:context:dispatcher:"), event, context, dispatcher)
	return rv
}
func (c CPXKeyboardEventProcessor) InitWithDeliveryManagerSpecialKeyEventProcessorProcessManagerDestinationGeneratorNotificationCenterKeyEventTracker(manager objectivec.IObject, processor objectivec.IObject, manager2 objectivec.IObject, generator objectivec.IObject, center objectivec.IObject, tracker objectivec.IObject) CPXKeyboardEventProcessor {
	rv := objc.Send[CPXKeyboardEventProcessor](c.ID, objc.Sel("initWithDeliveryManager:specialKeyEventProcessor:processManager:destinationGenerator:notificationCenter:keyEventTracker:"), manager, processor, manager2, generator, center, tracker)
	return rv
}
func (c CPXKeyboardEventProcessor) InitWithSessionSpecialKeyEventProcessor(session *CGXSession, processor objectivec.IObject) CPXKeyboardEventProcessor {
	rv := objc.Send[CPXKeyboardEventProcessor](c.ID, objc.Sel("initWithSession:specialKeyEventProcessor:"), session, processor)
	return rv
}

func (c CPXKeyboardEventProcessor) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXKeyboardEventProcessor) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXKeyboardEventProcessor) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}
func (c CPXKeyboardEventProcessor) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](c.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

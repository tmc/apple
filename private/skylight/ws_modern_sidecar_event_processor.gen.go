// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSModernSidecarEventProcessor] class.
var (
	_WSModernSidecarEventProcessorClass     WSModernSidecarEventProcessorClass
	_WSModernSidecarEventProcessorClassOnce sync.Once
)

func getWSModernSidecarEventProcessorClass() WSModernSidecarEventProcessorClass {
	_WSModernSidecarEventProcessorClassOnce.Do(func() {
		_WSModernSidecarEventProcessorClass = WSModernSidecarEventProcessorClass{class: objc.GetClass("WSModernSidecarEventProcessor")}
	})
	return _WSModernSidecarEventProcessorClass
}

// GetWSModernSidecarEventProcessorClass returns the class object for WSModernSidecarEventProcessor.
func GetWSModernSidecarEventProcessorClass() WSModernSidecarEventProcessorClass {
	return getWSModernSidecarEventProcessorClass()
}

type WSModernSidecarEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSModernSidecarEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSModernSidecarEventProcessorClass) Alloc() WSModernSidecarEventProcessor {
	rv := objc.SendIfResponds[WSModernSidecarEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSModernSidecarEventProcessor._fixupSideCar2Event]
//   - [WSModernSidecarEventProcessor.BkEventProcessor]
//   - [WSModernSidecarEventProcessor.ClearEventState]
//   - [WSModernSidecarEventProcessor.DeliveryObservationManager]
//   - [WSModernSidecarEventProcessor.DisplayRenderSpace]
//   - [WSModernSidecarEventProcessor.IndirectSystemGestureEventProcessor]
//   - [WSModernSidecarEventProcessor.InvalidateAndClearTouchLifecycleObservers]
//   - [WSModernSidecarEventProcessor.ProcessEventDispatcher]
//   - [WSModernSidecarEventProcessor.SystemGestureEventProcessor]
//   - [WSModernSidecarEventProcessor.TouchLifecycleObserverTokens]
//   - [WSModernSidecarEventProcessor.ValidateAndUpdateTouchLifecycleObservers]
//   - [WSModernSidecarEventProcessor.InitWithDeliveryManagerDispatcherHidSystem]
//   - [WSModernSidecarEventProcessor.DebugDescription]
//   - [WSModernSidecarEventProcessor.Description]
//   - [WSModernSidecarEventProcessor.Hash]
//   - [WSModernSidecarEventProcessor.Superclass]
type WSModernSidecarEventProcessor struct {
	objectivec.Object
}

// WSModernSidecarEventProcessorFromID constructs a [WSModernSidecarEventProcessor] from an objc.ID.
func WSModernSidecarEventProcessorFromID(id objc.ID) WSModernSidecarEventProcessor {
	return WSModernSidecarEventProcessor{objectivec.Object{ID: id}}
}

// Ensure WSModernSidecarEventProcessor implements IWSModernSidecarEventProcessor.
var _ IWSModernSidecarEventProcessor = WSModernSidecarEventProcessor{}

// An interface definition for the [WSModernSidecarEventProcessor] class.
//
// # Methods
//
//   - [IWSModernSidecarEventProcessor._fixupSideCar2Event]
//   - [IWSModernSidecarEventProcessor.BkEventProcessor]
//   - [IWSModernSidecarEventProcessor.ClearEventState]
//   - [IWSModernSidecarEventProcessor.DeliveryObservationManager]
//   - [IWSModernSidecarEventProcessor.DisplayRenderSpace]
//   - [IWSModernSidecarEventProcessor.IndirectSystemGestureEventProcessor]
//   - [IWSModernSidecarEventProcessor.InvalidateAndClearTouchLifecycleObservers]
//   - [IWSModernSidecarEventProcessor.ProcessEventDispatcher]
//   - [IWSModernSidecarEventProcessor.SystemGestureEventProcessor]
//   - [IWSModernSidecarEventProcessor.TouchLifecycleObserverTokens]
//   - [IWSModernSidecarEventProcessor.ValidateAndUpdateTouchLifecycleObservers]
//   - [IWSModernSidecarEventProcessor.InitWithDeliveryManagerDispatcherHidSystem]
//   - [IWSModernSidecarEventProcessor.DebugDescription]
//   - [IWSModernSidecarEventProcessor.Description]
//   - [IWSModernSidecarEventProcessor.Hash]
//   - [IWSModernSidecarEventProcessor.Superclass]
type IWSModernSidecarEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	_fixupSideCar2Event(car2Event uintptr)
	BkEventProcessor() objectivec.IObject
	ClearEventState()
	DeliveryObservationManager() objectivec.IObject
	DisplayRenderSpace() IWSDisplayRenderSpace
	IndirectSystemGestureEventProcessor() IWSIndirectSystemGestureEventProcessor
	InvalidateAndClearTouchLifecycleObservers()
	ProcessEventDispatcher(event *SLSEventRecord, dispatcher objectivec.IObject) int64
	SystemGestureEventProcessor() IWSSystemGestureEventProcessor
	TouchLifecycleObserverTokens() foundation.INSArray
	ValidateAndUpdateTouchLifecycleObservers()
	InitWithDeliveryManagerDispatcherHidSystem(manager objectivec.IObject, dispatcher objectivec.IObject, system objectivec.IObject) WSModernSidecarEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSModernSidecarEventProcessor) Init() WSModernSidecarEventProcessor {
	rv := objc.SendIfResponds[WSModernSidecarEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSModernSidecarEventProcessor) Autorelease() WSModernSidecarEventProcessor {
	rv := objc.SendIfResponds[WSModernSidecarEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSModernSidecarEventProcessor creates a new WSModernSidecarEventProcessor instance.
func NewWSModernSidecarEventProcessor() WSModernSidecarEventProcessor {
	class := getWSModernSidecarEventProcessorClass()
	rv := objc.SendIfResponds[WSModernSidecarEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSModernSidecarEventProcessorWithDeliveryManagerDispatcherHidSystem(manager objectivec.IObject, dispatcher objectivec.IObject, system objectivec.IObject) WSModernSidecarEventProcessor {
	instance := getWSModernSidecarEventProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDeliveryManager:dispatcher:hidSystem:"), manager, dispatcher, system)
	return WSModernSidecarEventProcessorFromID(rv)
}

func (w WSModernSidecarEventProcessor) _fixupSideCar2Event(car2Event uintptr) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_fixupSideCar2Event:"), car2Event)
}

// FixupSideCar2Event is an exported wrapper for the private method _fixupSideCar2Event.
func (w WSModernSidecarEventProcessor) FixupSideCar2Event(car2Event uintptr) error {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_fixupSideCar2Event:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_fixupSideCar2Event:"}
		return err
	}
	w._fixupSideCar2Event(car2Event)
	return nil
}

// CanFixupSideCar2Event reports whether the receiver responds to the private selector _fixupSideCar2Event:.
func (w WSModernSidecarEventProcessor) CanFixupSideCar2Event() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_fixupSideCar2Event:"))
}
func (w WSModernSidecarEventProcessor) ClearEventState() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("clearEventState"))
}
func (w WSModernSidecarEventProcessor) InvalidateAndClearTouchLifecycleObservers() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("invalidateAndClearTouchLifecycleObservers"))
}
func (w WSModernSidecarEventProcessor) ProcessEventDispatcher(event *SLSEventRecord, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](w.ID, objc.Sel("processEvent:dispatcher:"), unsafe.Pointer(event), dispatcher)
	return rv
}
func (w WSModernSidecarEventProcessor) ValidateAndUpdateTouchLifecycleObservers() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("validateAndUpdateTouchLifecycleObservers"))
}
func (w WSModernSidecarEventProcessor) InitWithDeliveryManagerDispatcherHidSystem(manager objectivec.IObject, dispatcher objectivec.IObject, system objectivec.IObject) WSModernSidecarEventProcessor {
	rv := objc.SendIfResponds[WSModernSidecarEventProcessor](w.ID, objc.Sel("initWithDeliveryManager:dispatcher:hidSystem:"), manager, dispatcher, system)
	return rv
}

func (_WSModernSidecarEventProcessorClass WSModernSidecarEventProcessorClass) IsAvailable() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_WSModernSidecarEventProcessorClass.class), objc.Sel("isAvailable"))
	return rv
}

func (w WSModernSidecarEventProcessor) BkEventProcessor() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("bkEventProcessor"))
	return objectivec.Object{ID: rv}
}
func (w WSModernSidecarEventProcessor) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSModernSidecarEventProcessor) DeliveryObservationManager() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("deliveryObservationManager"))
	return objectivec.Object{ID: rv}
}
func (w WSModernSidecarEventProcessor) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSModernSidecarEventProcessor) DisplayRenderSpace() IWSDisplayRenderSpace {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("displayRenderSpace"))
	return WSDisplayRenderSpaceFromID(objc.ID(rv))
}
func (w WSModernSidecarEventProcessor) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSModernSidecarEventProcessor) IndirectSystemGestureEventProcessor() IWSIndirectSystemGestureEventProcessor {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("indirectSystemGestureEventProcessor"))
	return WSIndirectSystemGestureEventProcessorFromID(objc.ID(rv))
}
func (w WSModernSidecarEventProcessor) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (w WSModernSidecarEventProcessor) SystemGestureEventProcessor() IWSSystemGestureEventProcessor {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("systemGestureEventProcessor"))
	return WSSystemGestureEventProcessorFromID(objc.ID(rv))
}
func (w WSModernSidecarEventProcessor) TouchLifecycleObserverTokens() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("touchLifecycleObserverTokens"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

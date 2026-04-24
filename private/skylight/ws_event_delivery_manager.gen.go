// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSEventDeliveryManager] class.
var (
	_WSEventDeliveryManagerClass     WSEventDeliveryManagerClass
	_WSEventDeliveryManagerClassOnce sync.Once
)

func getWSEventDeliveryManagerClass() WSEventDeliveryManagerClass {
	_WSEventDeliveryManagerClassOnce.Do(func() {
		_WSEventDeliveryManagerClass = WSEventDeliveryManagerClass{class: objc.GetClass("WSEventDeliveryManager")}
	})
	return _WSEventDeliveryManagerClass
}

// GetWSEventDeliveryManagerClass returns the class object for WSEventDeliveryManager.
func GetWSEventDeliveryManagerClass() WSEventDeliveryManagerClass {
	return getWSEventDeliveryManagerClass()
}

type WSEventDeliveryManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSEventDeliveryManagerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSEventDeliveryManagerClass) Alloc() WSEventDeliveryManager {
	rv := objc.Send[WSEventDeliveryManager](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSEventDeliveryManager.BkDeliveryManager]
//   - [WSEventDeliveryManager.BkEventDeliveryObserverService]
//   - [WSEventDeliveryManager.DeferEventsMatchingPredicateToTargetWithReason]
//   - [WSEventDeliveryManager.DeliveryChainsForDeferringTargetEvent]
//   - [WSEventDeliveryManager.DeliveryGraphDescription]
//   - [WSEventDeliveryManager.DescriptionOfResolutionPathForEventDescriptorSenderDescriptor]
//   - [WSEventDeliveryManager.DestinationsForEvent]
//   - [WSEventDeliveryManager.DispatchDiscreteEventsForReasonWithRules]
//   - [WSEventDeliveryManager.ResolveDestinationsForEventStartingFromPID]
//   - [WSEventDeliveryManager.TransactionAssertionWithReason]
//   - [WSEventDeliveryManager.ValidateTokenAndEnvironmentForEventProcessManagerOutReason]
//   - [WSEventDeliveryManager.InitWithObserverService]
//
// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager
type WSEventDeliveryManager struct {
	objectivec.Object
}

// WSEventDeliveryManagerFromID constructs a [WSEventDeliveryManager] from an objc.ID.
func WSEventDeliveryManagerFromID(id objc.ID) WSEventDeliveryManager {
	return WSEventDeliveryManager{objectivec.Object{ID: id}}
}

// Ensure WSEventDeliveryManager implements IWSEventDeliveryManager.
var _ IWSEventDeliveryManager = WSEventDeliveryManager{}

// An interface definition for the [WSEventDeliveryManager] class.
//
// # Methods
//
//   - [IWSEventDeliveryManager.BkDeliveryManager]
//   - [IWSEventDeliveryManager.BkEventDeliveryObserverService]
//   - [IWSEventDeliveryManager.DeferEventsMatchingPredicateToTargetWithReason]
//   - [IWSEventDeliveryManager.DeliveryChainsForDeferringTargetEvent]
//   - [IWSEventDeliveryManager.DeliveryGraphDescription]
//   - [IWSEventDeliveryManager.DescriptionOfResolutionPathForEventDescriptorSenderDescriptor]
//   - [IWSEventDeliveryManager.DestinationsForEvent]
//   - [IWSEventDeliveryManager.DispatchDiscreteEventsForReasonWithRules]
//   - [IWSEventDeliveryManager.ResolveDestinationsForEventStartingFromPID]
//   - [IWSEventDeliveryManager.TransactionAssertionWithReason]
//   - [IWSEventDeliveryManager.ValidateTokenAndEnvironmentForEventProcessManagerOutReason]
//   - [IWSEventDeliveryManager.InitWithObserverService]
//
// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager
type IWSEventDeliveryManager interface {
	objectivec.IObject

	// Topic: Methods

	BkDeliveryManager() unsafe.Pointer
	BkEventDeliveryObserverService() unsafe.Pointer
	DeferEventsMatchingPredicateToTargetWithReason(predicate objectivec.IObject, target objectivec.IObject, reason objectivec.IObject) objectivec.IObject
	DeliveryChainsForDeferringTargetEvent(target objectivec.IObject, event objectivec.IObject) objectivec.IObject
	DeliveryGraphDescription() objectivec.IObject
	DescriptionOfResolutionPathForEventDescriptorSenderDescriptor(descriptor objectivec.IObject, descriptor2 objectivec.IObject) objectivec.IObject
	DestinationsForEvent(event *SLSEventRecordRef) objectivec.IObject
	DispatchDiscreteEventsForReasonWithRules(reason objectivec.IObject, rules objectivec.IObject) objectivec.IObject
	ResolveDestinationsForEventStartingFromPID(event *SLSEventRecordRef, pid int) objectivec.IObject
	TransactionAssertionWithReason(reason objectivec.IObject) objectivec.IObject
	ValidateTokenAndEnvironmentForEventProcessManagerOutReason(event *SLSEventRecordRef, manager objectivec.IObject, reason []objectivec.IObject) bool
	InitWithObserverService(service objectivec.IObject) WSEventDeliveryManager
}

// Init initializes the instance.
func (w WSEventDeliveryManager) Init() WSEventDeliveryManager {
	rv := objc.Send[WSEventDeliveryManager](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSEventDeliveryManager) Autorelease() WSEventDeliveryManager {
	rv := objc.Send[WSEventDeliveryManager](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSEventDeliveryManager creates a new WSEventDeliveryManager instance.
func NewWSEventDeliveryManager() WSEventDeliveryManager {
	class := getWSEventDeliveryManagerClass()
	rv := objc.Send[WSEventDeliveryManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/initWithObserverService:
func NewWSEventDeliveryManagerWithObserverService(service objectivec.IObject) WSEventDeliveryManager {
	instance := getWSEventDeliveryManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObserverService:"), service)
	return WSEventDeliveryManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/deferEventsMatchingPredicate:toTarget:withReason:
func (w WSEventDeliveryManager) DeferEventsMatchingPredicateToTargetWithReason(predicate objectivec.IObject, target objectivec.IObject, reason objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("deferEventsMatchingPredicate:toTarget:withReason:"), predicate, target, reason)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/deliveryChainsForDeferringTarget:event:
func (w WSEventDeliveryManager) DeliveryChainsForDeferringTargetEvent(target objectivec.IObject, event objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("deliveryChainsForDeferringTarget:event:"), target, event)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/deliveryGraphDescription
func (w WSEventDeliveryManager) DeliveryGraphDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("deliveryGraphDescription"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/descriptionOfResolutionPathForEventDescriptor:senderDescriptor:
func (w WSEventDeliveryManager) DescriptionOfResolutionPathForEventDescriptorSenderDescriptor(descriptor objectivec.IObject, descriptor2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("descriptionOfResolutionPathForEventDescriptor:senderDescriptor:"), descriptor, descriptor2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/destinationsForEvent:
func (w WSEventDeliveryManager) DestinationsForEvent(event *SLSEventRecordRef) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("destinationsForEvent:"), event)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/dispatchDiscreteEventsForReason:withRules:
func (w WSEventDeliveryManager) DispatchDiscreteEventsForReasonWithRules(reason objectivec.IObject, rules objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("dispatchDiscreteEventsForReason:withRules:"), reason, rules)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/resolveDestinationsForEvent:startingFromPID:
func (w WSEventDeliveryManager) ResolveDestinationsForEventStartingFromPID(event *SLSEventRecordRef, pid int) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("resolveDestinationsForEvent:startingFromPID:"), event, pid)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/transactionAssertionWithReason:
func (w WSEventDeliveryManager) TransactionAssertionWithReason(reason objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("transactionAssertionWithReason:"), reason)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/validateTokenAndEnvironmentForEvent:processManager:outReason:
func (w WSEventDeliveryManager) ValidateTokenAndEnvironmentForEventProcessManagerOutReason(event *SLSEventRecordRef, manager objectivec.IObject, reason []objectivec.IObject) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("validateTokenAndEnvironmentForEvent:processManager:outReason:"), event, manager, objectivec.IObjectSliceToNSArray(reason))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/initWithObserverService:
func (w WSEventDeliveryManager) InitWithObserverService(service objectivec.IObject) WSEventDeliveryManager {
	rv := objc.Send[WSEventDeliveryManager](w.ID, objc.Sel("initWithObserverService:"), service)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/possibleIOHIDKeyboardEventUsageForVirtualKeyCode:
func (_WSEventDeliveryManagerClass WSEventDeliveryManagerClass) PossibleIOHIDKeyboardEventUsageForVirtualKeyCode(code uint32) uint32 {
	rv := objc.Send[uint32](objc.ID(_WSEventDeliveryManagerClass.class), objc.Sel("possibleIOHIDKeyboardEventUsageForVirtualKeyCode:"), code)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/bkDeliveryManager
func (w WSEventDeliveryManager) BkDeliveryManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w.ID, objc.Sel("bkDeliveryManager"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventDeliveryManager/bkEventDeliveryObserverService
func (w WSEventDeliveryManager) BkEventDeliveryObserverService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w.ID, objc.Sel("bkEventDeliveryObserverService"))
	return rv
}

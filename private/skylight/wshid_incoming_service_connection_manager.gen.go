// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSHIDIncomingServiceConnectionManager] class.
var (
	_WSHIDIncomingServiceConnectionManagerClass     WSHIDIncomingServiceConnectionManagerClass
	_WSHIDIncomingServiceConnectionManagerClassOnce sync.Once
)

func getWSHIDIncomingServiceConnectionManagerClass() WSHIDIncomingServiceConnectionManagerClass {
	_WSHIDIncomingServiceConnectionManagerClassOnce.Do(func() {
		_WSHIDIncomingServiceConnectionManagerClass = WSHIDIncomingServiceConnectionManagerClass{class: objc.GetClass("WSHIDIncomingServiceConnectionManager")}
	})
	return _WSHIDIncomingServiceConnectionManagerClass
}

// GetWSHIDIncomingServiceConnectionManagerClass returns the class object for WSHIDIncomingServiceConnectionManager.
func GetWSHIDIncomingServiceConnectionManagerClass() WSHIDIncomingServiceConnectionManagerClass {
	return getWSHIDIncomingServiceConnectionManagerClass()
}

type WSHIDIncomingServiceConnectionManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSHIDIncomingServiceConnectionManagerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSHIDIncomingServiceConnectionManagerClass) Alloc() WSHIDIncomingServiceConnectionManager {
	rv := objc.SendIfResponds[WSHIDIncomingServiceConnectionManager](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSHIDIncomingServiceConnectionManager._init]
//   - [WSHIDIncomingServiceConnectionManager._queue_appendDescriptionToStream]
//   - [WSHIDIncomingServiceConnectionManager._queue_deliveryManagerForAuditToken]
//   - [WSHIDIncomingServiceConnectionManager._queue_description]
//   - [WSHIDIncomingServiceConnectionManager._queue_eventDeliveryObserverServiceForAuditToken]
//   - [WSHIDIncomingServiceConnectionManager.AppendDescriptionToStream]
//   - [WSHIDIncomingServiceConnectionManager.DidUpdateEventDeliveryManagerForSession]
//   - [WSHIDIncomingServiceConnectionManager.HandleIncomingDeliveryManagerConnection]
//   - [WSHIDIncomingServiceConnectionManager.HandleIncomingDeliveryObserverConnection]
//   - [WSHIDIncomingServiceConnectionManager.IncomingServiceConnectionDidRevoke]
type WSHIDIncomingServiceConnectionManager struct {
	objectivec.Object
}

// WSHIDIncomingServiceConnectionManagerFromID constructs a [WSHIDIncomingServiceConnectionManager] from an objc.ID.
func WSHIDIncomingServiceConnectionManagerFromID(id objc.ID) WSHIDIncomingServiceConnectionManager {
	return WSHIDIncomingServiceConnectionManager{objectivec.Object{ID: id}}
}

// Ensure WSHIDIncomingServiceConnectionManager implements IWSHIDIncomingServiceConnectionManager.
var _ IWSHIDIncomingServiceConnectionManager = WSHIDIncomingServiceConnectionManager{}

// An interface definition for the [WSHIDIncomingServiceConnectionManager] class.
//
// # Methods
//
//   - [IWSHIDIncomingServiceConnectionManager._init]
//   - [IWSHIDIncomingServiceConnectionManager._queue_appendDescriptionToStream]
//   - [IWSHIDIncomingServiceConnectionManager._queue_deliveryManagerForAuditToken]
//   - [IWSHIDIncomingServiceConnectionManager._queue_description]
//   - [IWSHIDIncomingServiceConnectionManager._queue_eventDeliveryObserverServiceForAuditToken]
//   - [IWSHIDIncomingServiceConnectionManager.AppendDescriptionToStream]
//   - [IWSHIDIncomingServiceConnectionManager.DidUpdateEventDeliveryManagerForSession]
//   - [IWSHIDIncomingServiceConnectionManager.HandleIncomingDeliveryManagerConnection]
//   - [IWSHIDIncomingServiceConnectionManager.HandleIncomingDeliveryObserverConnection]
//   - [IWSHIDIncomingServiceConnectionManager.IncomingServiceConnectionDidRevoke]
type IWSHIDIncomingServiceConnectionManager interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_queue_appendDescriptionToStream(stream objectivec.IObject)
	_queue_deliveryManagerForAuditToken(token objectivec.IObject) objectivec.IObject
	_queue_description() objectivec.IObject
	_queue_eventDeliveryObserverServiceForAuditToken(token objectivec.IObject) objectivec.IObject
	AppendDescriptionToStream(stream objectivec.IObject)
	DidUpdateEventDeliveryManagerForSession()
	HandleIncomingDeliveryManagerConnection(connection objectivec.IObject)
	HandleIncomingDeliveryObserverConnection(connection objectivec.IObject)
	IncomingServiceConnectionDidRevoke(revoke objectivec.IObject)
}

// Init initializes the instance.
func (w WSHIDIncomingServiceConnectionManager) Init() WSHIDIncomingServiceConnectionManager {
	rv := objc.SendIfResponds[WSHIDIncomingServiceConnectionManager](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSHIDIncomingServiceConnectionManager) Autorelease() WSHIDIncomingServiceConnectionManager {
	rv := objc.SendIfResponds[WSHIDIncomingServiceConnectionManager](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSHIDIncomingServiceConnectionManager creates a new WSHIDIncomingServiceConnectionManager instance.
func NewWSHIDIncomingServiceConnectionManager() WSHIDIncomingServiceConnectionManager {
	class := getWSHIDIncomingServiceConnectionManagerClass()
	rv := objc.SendIfResponds[WSHIDIncomingServiceConnectionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSHIDIncomingServiceConnectionManager) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (w WSHIDIncomingServiceConnectionManager) _queue_appendDescriptionToStream(stream objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_appendDescriptionToStream:"), stream)
}

// Queue_appendDescriptionToStream is an exported wrapper for the private method _queue_appendDescriptionToStream.
func (w WSHIDIncomingServiceConnectionManager) Queue_appendDescriptionToStream(stream objectivec.IObject) error {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_appendDescriptionToStream:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_appendDescriptionToStream:"}
		return err
	}
	w._queue_appendDescriptionToStream(stream)
	return nil
}

// CanQueue_appendDescriptionToStream reports whether the receiver responds to the private selector _queue_appendDescriptionToStream:.
func (w WSHIDIncomingServiceConnectionManager) CanQueue_appendDescriptionToStream() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_appendDescriptionToStream:"))
}
func (w WSHIDIncomingServiceConnectionManager) _queue_deliveryManagerForAuditToken(token objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_deliveryManagerForAuditToken:"), token)
	return objectivec.Object{ID: rv}
}

// Queue_deliveryManagerForAuditToken is an exported wrapper for the private method _queue_deliveryManagerForAuditToken.
func (w WSHIDIncomingServiceConnectionManager) Queue_deliveryManagerForAuditToken(token objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_deliveryManagerForAuditToken:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_deliveryManagerForAuditToken:"}
		return nil, err
	}
	return w._queue_deliveryManagerForAuditToken(token), nil
}

// CanQueue_deliveryManagerForAuditToken reports whether the receiver responds to the private selector _queue_deliveryManagerForAuditToken:.
func (w WSHIDIncomingServiceConnectionManager) CanQueue_deliveryManagerForAuditToken() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_deliveryManagerForAuditToken:"))
}
func (w WSHIDIncomingServiceConnectionManager) _queue_description() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_description"))
	return objectivec.Object{ID: rv}
}

// Queue_description is an exported wrapper for the private method _queue_description.
func (w WSHIDIncomingServiceConnectionManager) Queue_description() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_description")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_description"}
		return nil, err
	}
	return w._queue_description(), nil
}

// CanQueue_description reports whether the receiver responds to the private selector _queue_description.
func (w WSHIDIncomingServiceConnectionManager) CanQueue_description() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_description"))
}
func (w WSHIDIncomingServiceConnectionManager) _queue_eventDeliveryObserverServiceForAuditToken(token objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_eventDeliveryObserverServiceForAuditToken:"), token)
	return objectivec.Object{ID: rv}
}

// Queue_eventDeliveryObserverServiceForAuditToken is an exported wrapper for the private method _queue_eventDeliveryObserverServiceForAuditToken.
func (w WSHIDIncomingServiceConnectionManager) Queue_eventDeliveryObserverServiceForAuditToken(token objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_eventDeliveryObserverServiceForAuditToken:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_eventDeliveryObserverServiceForAuditToken:"}
		return nil, err
	}
	return w._queue_eventDeliveryObserverServiceForAuditToken(token), nil
}

// CanQueue_eventDeliveryObserverServiceForAuditToken reports whether the receiver responds to the private selector _queue_eventDeliveryObserverServiceForAuditToken:.
func (w WSHIDIncomingServiceConnectionManager) CanQueue_eventDeliveryObserverServiceForAuditToken() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_eventDeliveryObserverServiceForAuditToken:"))
}
func (w WSHIDIncomingServiceConnectionManager) AppendDescriptionToStream(stream objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("appendDescriptionToStream:"), stream)
}
func (w WSHIDIncomingServiceConnectionManager) DidUpdateEventDeliveryManagerForSession() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("didUpdateEventDeliveryManagerForSession"))
}
func (w WSHIDIncomingServiceConnectionManager) HandleIncomingDeliveryManagerConnection(connection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("handleIncomingDeliveryManagerConnection:"), connection)
}
func (w WSHIDIncomingServiceConnectionManager) HandleIncomingDeliveryObserverConnection(connection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("handleIncomingDeliveryObserverConnection:"), connection)
}
func (w WSHIDIncomingServiceConnectionManager) IncomingServiceConnectionDidRevoke(revoke objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("incomingServiceConnectionDidRevoke:"), revoke)
}

func (_WSHIDIncomingServiceConnectionManagerClass WSHIDIncomingServiceConnectionManagerClass) SharedInstance() WSHIDIncomingServiceConnectionManager {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_WSHIDIncomingServiceConnectionManagerClass.class), objc.Sel("sharedInstance"))
	return WSHIDIncomingServiceConnectionManagerFromID(rv)
}

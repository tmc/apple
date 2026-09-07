// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSHIDIncomingServiceConnectionManagerLegacy] class.
var (
	_WSHIDIncomingServiceConnectionManagerLegacyClass     WSHIDIncomingServiceConnectionManagerLegacyClass
	_WSHIDIncomingServiceConnectionManagerLegacyClassOnce sync.Once
)

func getWSHIDIncomingServiceConnectionManagerLegacyClass() WSHIDIncomingServiceConnectionManagerLegacyClass {
	_WSHIDIncomingServiceConnectionManagerLegacyClassOnce.Do(func() {
		_WSHIDIncomingServiceConnectionManagerLegacyClass = WSHIDIncomingServiceConnectionManagerLegacyClass{class: objc.GetClass("_WSHIDIncomingServiceConnectionManager_Legacy")}
	})
	return _WSHIDIncomingServiceConnectionManagerLegacyClass
}

// GetWSHIDIncomingServiceConnectionManagerLegacyClass returns the class object for _WSHIDIncomingServiceConnectionManager_Legacy.
func GetWSHIDIncomingServiceConnectionManagerLegacyClass() WSHIDIncomingServiceConnectionManagerLegacyClass {
	return getWSHIDIncomingServiceConnectionManagerLegacyClass()
}

type WSHIDIncomingServiceConnectionManagerLegacyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSHIDIncomingServiceConnectionManagerLegacyClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSHIDIncomingServiceConnectionManagerLegacyClass) Alloc() WSHIDIncomingServiceConnectionManagerLegacy {
	rv := objc.SendIfResponds[WSHIDIncomingServiceConnectionManagerLegacy](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSHIDIncomingServiceConnectionManagerLegacy._queue_acceptIncomingConnectionsMappedObjectFetcher]
//   - [WSHIDIncomingServiceConnectionManagerLegacy._queue_appendDescriptionToStream]
//   - [WSHIDIncomingServiceConnectionManagerLegacy._queue_deliveryManagerForAuditToken]
//   - [WSHIDIncomingServiceConnectionManagerLegacy._queue_description]
//   - [WSHIDIncomingServiceConnectionManagerLegacy._queue_eventDeliveryObserverServiceForAuditToken]
//   - [WSHIDIncomingServiceConnectionManagerLegacy._queue_modernSidecarEventProcessorForAuditToken]
//   - [WSHIDIncomingServiceConnectionManagerLegacy._queue_touchDeliveryObserverForAuditToken]
//   - [WSHIDIncomingServiceConnectionManagerLegacy.AppendDescriptionToStream]
//   - [WSHIDIncomingServiceConnectionManagerLegacy.HandleIncomingDeliveryManagerConnection]
//   - [WSHIDIncomingServiceConnectionManagerLegacy.HandleIncomingDeliveryObserverConnection]
//   - [WSHIDIncomingServiceConnectionManagerLegacy.IncomingServiceConnectionDidRevoke]
//   - [WSHIDIncomingServiceConnectionManagerLegacy.DebugDescription]
//   - [WSHIDIncomingServiceConnectionManagerLegacy.Description]
//   - [WSHIDIncomingServiceConnectionManagerLegacy.Hash]
//   - [WSHIDIncomingServiceConnectionManagerLegacy.Superclass]
type WSHIDIncomingServiceConnectionManagerLegacy struct {
	objectivec.Object
}

// WSHIDIncomingServiceConnectionManagerLegacyFromID constructs a [WSHIDIncomingServiceConnectionManagerLegacy] from an objc.ID.
func WSHIDIncomingServiceConnectionManagerLegacyFromID(id objc.ID) WSHIDIncomingServiceConnectionManagerLegacy {
	return WSHIDIncomingServiceConnectionManagerLegacy{objectivec.Object{ID: id}}
}

// Ensure WSHIDIncomingServiceConnectionManagerLegacy implements IWSHIDIncomingServiceConnectionManagerLegacy.
var _ IWSHIDIncomingServiceConnectionManagerLegacy = WSHIDIncomingServiceConnectionManagerLegacy{}

// An interface definition for the [WSHIDIncomingServiceConnectionManagerLegacy] class.
//
// # Methods
//
//   - [IWSHIDIncomingServiceConnectionManagerLegacy._queue_acceptIncomingConnectionsMappedObjectFetcher]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy._queue_appendDescriptionToStream]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy._queue_deliveryManagerForAuditToken]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy._queue_description]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy._queue_eventDeliveryObserverServiceForAuditToken]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy._queue_modernSidecarEventProcessorForAuditToken]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy._queue_touchDeliveryObserverForAuditToken]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy.AppendDescriptionToStream]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy.HandleIncomingDeliveryManagerConnection]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy.HandleIncomingDeliveryObserverConnection]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy.IncomingServiceConnectionDidRevoke]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy.DebugDescription]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy.Description]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy.Hash]
//   - [IWSHIDIncomingServiceConnectionManagerLegacy.Superclass]
type IWSHIDIncomingServiceConnectionManagerLegacy interface {
	objectivec.IObject

	// Topic: Methods

	_queue_acceptIncomingConnectionsMappedObjectFetcher(connections objectivec.IObject, fetcher VoidHandler)
	_queue_appendDescriptionToStream(stream objectivec.IObject)
	_queue_deliveryManagerForAuditToken(token objectivec.IObject) objectivec.IObject
	_queue_description() objectivec.IObject
	_queue_eventDeliveryObserverServiceForAuditToken(token objectivec.IObject) objectivec.IObject
	_queue_modernSidecarEventProcessorForAuditToken(token objectivec.IObject) objectivec.IObject
	_queue_touchDeliveryObserverForAuditToken(token objectivec.IObject) objectivec.IObject
	AppendDescriptionToStream(stream objectivec.IObject)
	HandleIncomingDeliveryManagerConnection(connection objectivec.IObject)
	HandleIncomingDeliveryObserverConnection(connection objectivec.IObject)
	IncomingServiceConnectionDidRevoke(revoke objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSHIDIncomingServiceConnectionManagerLegacy) Init() WSHIDIncomingServiceConnectionManagerLegacy {
	rv := objc.SendIfResponds[WSHIDIncomingServiceConnectionManagerLegacy](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSHIDIncomingServiceConnectionManagerLegacy) Autorelease() WSHIDIncomingServiceConnectionManagerLegacy {
	rv := objc.SendIfResponds[WSHIDIncomingServiceConnectionManagerLegacy](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSHIDIncomingServiceConnectionManagerLegacy creates a new WSHIDIncomingServiceConnectionManagerLegacy instance.
func NewWSHIDIncomingServiceConnectionManagerLegacy() WSHIDIncomingServiceConnectionManagerLegacy {
	class := getWSHIDIncomingServiceConnectionManagerLegacyClass()
	rv := objc.SendIfResponds[WSHIDIncomingServiceConnectionManagerLegacy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSHIDIncomingServiceConnectionManagerLegacy) _queue_acceptIncomingConnectionsMappedObjectFetcher(connections objectivec.IObject, fetcher VoidHandler) {
	_block1, _ := NewVoidBlock(fetcher)
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_acceptIncomingConnections:mappedObjectFetcher:"), connections, _block1)
}

// Queue_acceptIncomingConnectionsMappedObjectFetcher is an exported wrapper for the private method _queue_acceptIncomingConnectionsMappedObjectFetcher.
func (w WSHIDIncomingServiceConnectionManagerLegacy) Queue_acceptIncomingConnectionsMappedObjectFetcher(connections objectivec.IObject, fetcher VoidHandler) error {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_acceptIncomingConnections:mappedObjectFetcher:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_acceptIncomingConnections:mappedObjectFetcher:"}
		return err
	}
	w._queue_acceptIncomingConnectionsMappedObjectFetcher(connections, fetcher)
	return nil
}

// CanQueue_acceptIncomingConnectionsMappedObjectFetcher reports whether the receiver responds to the private selector _queue_acceptIncomingConnections:mappedObjectFetcher:.
func (w WSHIDIncomingServiceConnectionManagerLegacy) CanQueue_acceptIncomingConnectionsMappedObjectFetcher() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_acceptIncomingConnections:mappedObjectFetcher:"))
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) _queue_appendDescriptionToStream(stream objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_appendDescriptionToStream:"), stream)
}

// Queue_appendDescriptionToStream is an exported wrapper for the private method _queue_appendDescriptionToStream.
func (w WSHIDIncomingServiceConnectionManagerLegacy) Queue_appendDescriptionToStream(stream objectivec.IObject) error {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_appendDescriptionToStream:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_appendDescriptionToStream:"}
		return err
	}
	w._queue_appendDescriptionToStream(stream)
	return nil
}

// CanQueue_appendDescriptionToStream reports whether the receiver responds to the private selector _queue_appendDescriptionToStream:.
func (w WSHIDIncomingServiceConnectionManagerLegacy) CanQueue_appendDescriptionToStream() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_appendDescriptionToStream:"))
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) _queue_deliveryManagerForAuditToken(token objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_deliveryManagerForAuditToken:"), token)
	return objectivec.Object{ID: rv}
}

// Queue_deliveryManagerForAuditToken is an exported wrapper for the private method _queue_deliveryManagerForAuditToken.
func (w WSHIDIncomingServiceConnectionManagerLegacy) Queue_deliveryManagerForAuditToken(token objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_deliveryManagerForAuditToken:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_deliveryManagerForAuditToken:"}
		return nil, err
	}
	return w._queue_deliveryManagerForAuditToken(token), nil
}

// CanQueue_deliveryManagerForAuditToken reports whether the receiver responds to the private selector _queue_deliveryManagerForAuditToken:.
func (w WSHIDIncomingServiceConnectionManagerLegacy) CanQueue_deliveryManagerForAuditToken() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_deliveryManagerForAuditToken:"))
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) _queue_description() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_description"))
	return objectivec.Object{ID: rv}
}

// Queue_description is an exported wrapper for the private method _queue_description.
func (w WSHIDIncomingServiceConnectionManagerLegacy) Queue_description() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_description")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_description"}
		return nil, err
	}
	return w._queue_description(), nil
}

// CanQueue_description reports whether the receiver responds to the private selector _queue_description.
func (w WSHIDIncomingServiceConnectionManagerLegacy) CanQueue_description() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_description"))
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) _queue_eventDeliveryObserverServiceForAuditToken(token objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_eventDeliveryObserverServiceForAuditToken:"), token)
	return objectivec.Object{ID: rv}
}

// Queue_eventDeliveryObserverServiceForAuditToken is an exported wrapper for the private method _queue_eventDeliveryObserverServiceForAuditToken.
func (w WSHIDIncomingServiceConnectionManagerLegacy) Queue_eventDeliveryObserverServiceForAuditToken(token objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_eventDeliveryObserverServiceForAuditToken:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_eventDeliveryObserverServiceForAuditToken:"}
		return nil, err
	}
	return w._queue_eventDeliveryObserverServiceForAuditToken(token), nil
}

// CanQueue_eventDeliveryObserverServiceForAuditToken reports whether the receiver responds to the private selector _queue_eventDeliveryObserverServiceForAuditToken:.
func (w WSHIDIncomingServiceConnectionManagerLegacy) CanQueue_eventDeliveryObserverServiceForAuditToken() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_eventDeliveryObserverServiceForAuditToken:"))
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) _queue_modernSidecarEventProcessorForAuditToken(token objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_modernSidecarEventProcessorForAuditToken:"), token)
	return objectivec.Object{ID: rv}
}

// Queue_modernSidecarEventProcessorForAuditToken is an exported wrapper for the private method _queue_modernSidecarEventProcessorForAuditToken.
func (w WSHIDIncomingServiceConnectionManagerLegacy) Queue_modernSidecarEventProcessorForAuditToken(token objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_modernSidecarEventProcessorForAuditToken:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_modernSidecarEventProcessorForAuditToken:"}
		return nil, err
	}
	return w._queue_modernSidecarEventProcessorForAuditToken(token), nil
}

// CanQueue_modernSidecarEventProcessorForAuditToken reports whether the receiver responds to the private selector _queue_modernSidecarEventProcessorForAuditToken:.
func (w WSHIDIncomingServiceConnectionManagerLegacy) CanQueue_modernSidecarEventProcessorForAuditToken() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_modernSidecarEventProcessorForAuditToken:"))
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) _queue_touchDeliveryObserverForAuditToken(token objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_queue_touchDeliveryObserverForAuditToken:"), token)
	return objectivec.Object{ID: rv}
}

// Queue_touchDeliveryObserverForAuditToken is an exported wrapper for the private method _queue_touchDeliveryObserverForAuditToken.
func (w WSHIDIncomingServiceConnectionManagerLegacy) Queue_touchDeliveryObserverForAuditToken(token objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_queue_touchDeliveryObserverForAuditToken:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queue_touchDeliveryObserverForAuditToken:"}
		return nil, err
	}
	return w._queue_touchDeliveryObserverForAuditToken(token), nil
}

// CanQueue_touchDeliveryObserverForAuditToken reports whether the receiver responds to the private selector _queue_touchDeliveryObserverForAuditToken:.
func (w WSHIDIncomingServiceConnectionManagerLegacy) CanQueue_touchDeliveryObserverForAuditToken() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_queue_touchDeliveryObserverForAuditToken:"))
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) AppendDescriptionToStream(stream objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("appendDescriptionToStream:"), stream)
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) HandleIncomingDeliveryManagerConnection(connection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("handleIncomingDeliveryManagerConnection:"), connection)
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) HandleIncomingDeliveryObserverConnection(connection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("handleIncomingDeliveryObserverConnection:"), connection)
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) IncomingServiceConnectionDidRevoke(revoke objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("incomingServiceConnectionDidRevoke:"), revoke)
}

func (w WSHIDIncomingServiceConnectionManagerLegacy) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSHIDIncomingServiceConnectionManagerLegacy) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

// _queue_acceptIncomingConnectionsMappedObjectFetcherSync is a synchronous wrapper around [WSHIDIncomingServiceConnectionManagerLegacy._queue_acceptIncomingConnectionsMappedObjectFetcher].
// It blocks until the completion handler fires or the context is cancelled.
func (w WSHIDIncomingServiceConnectionManagerLegacy) _queue_acceptIncomingConnectionsMappedObjectFetcherSync(ctx context.Context, connections objectivec.IObject) error {
	done := make(chan struct{}, 1)
	w._queue_acceptIncomingConnectionsMappedObjectFetcher(connections, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

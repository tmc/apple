// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
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
//   - [WSHIDIncomingServiceConnectionManager.DidUpdateEventDeliveryManagerForSession]
//   - [WSHIDIncomingServiceConnectionManager.DidUpdateModernSidecarEventProcessorForSession]
//   - [WSHIDIncomingServiceConnectionManager.HandleIncomingDeliveryManagerConnection]
//   - [WSHIDIncomingServiceConnectionManager.HandleIncomingDeliveryObserverConnection]
//   - [WSHIDIncomingServiceConnectionManager.HandleIncomingTouchDeliveryObservationConnection]
//   - [WSHIDIncomingServiceConnectionManager.HandleIncomingTouchEventConnection]
//   - [WSHIDIncomingServiceConnectionManager.HandleIncomingTouchStreamConnection]
//   - [WSHIDIncomingServiceConnectionManager.IncomingServiceConnectionDidRevoke]
//   - [WSHIDIncomingServiceConnectionManager.DebugDescription]
//   - [WSHIDIncomingServiceConnectionManager.Description]
//   - [WSHIDIncomingServiceConnectionManager.Hash]
//   - [WSHIDIncomingServiceConnectionManager.Superclass]
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
//   - [IWSHIDIncomingServiceConnectionManager.DidUpdateEventDeliveryManagerForSession]
//   - [IWSHIDIncomingServiceConnectionManager.DidUpdateModernSidecarEventProcessorForSession]
//   - [IWSHIDIncomingServiceConnectionManager.HandleIncomingDeliveryManagerConnection]
//   - [IWSHIDIncomingServiceConnectionManager.HandleIncomingDeliveryObserverConnection]
//   - [IWSHIDIncomingServiceConnectionManager.HandleIncomingTouchDeliveryObservationConnection]
//   - [IWSHIDIncomingServiceConnectionManager.HandleIncomingTouchEventConnection]
//   - [IWSHIDIncomingServiceConnectionManager.HandleIncomingTouchStreamConnection]
//   - [IWSHIDIncomingServiceConnectionManager.IncomingServiceConnectionDidRevoke]
//   - [IWSHIDIncomingServiceConnectionManager.DebugDescription]
//   - [IWSHIDIncomingServiceConnectionManager.Description]
//   - [IWSHIDIncomingServiceConnectionManager.Hash]
//   - [IWSHIDIncomingServiceConnectionManager.Superclass]
type IWSHIDIncomingServiceConnectionManager interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	DidUpdateEventDeliveryManagerForSession()
	DidUpdateModernSidecarEventProcessorForSession()
	HandleIncomingDeliveryManagerConnection(connection objectivec.IObject)
	HandleIncomingDeliveryObserverConnection(connection objectivec.IObject)
	HandleIncomingTouchDeliveryObservationConnection(connection objectivec.IObject)
	HandleIncomingTouchEventConnection(connection objectivec.IObject)
	HandleIncomingTouchStreamConnection(connection objectivec.IObject)
	IncomingServiceConnectionDidRevoke(revoke objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
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
func (w WSHIDIncomingServiceConnectionManager) DidUpdateEventDeliveryManagerForSession() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("didUpdateEventDeliveryManagerForSession"))
}
func (w WSHIDIncomingServiceConnectionManager) DidUpdateModernSidecarEventProcessorForSession() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("didUpdateModernSidecarEventProcessorForSession"))
}
func (w WSHIDIncomingServiceConnectionManager) HandleIncomingDeliveryManagerConnection(connection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("handleIncomingDeliveryManagerConnection:"), connection)
}
func (w WSHIDIncomingServiceConnectionManager) HandleIncomingDeliveryObserverConnection(connection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("handleIncomingDeliveryObserverConnection:"), connection)
}
func (w WSHIDIncomingServiceConnectionManager) HandleIncomingTouchDeliveryObservationConnection(connection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("handleIncomingTouchDeliveryObservationConnection:"), connection)
}
func (w WSHIDIncomingServiceConnectionManager) HandleIncomingTouchEventConnection(connection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("handleIncomingTouchEventConnection:"), connection)
}
func (w WSHIDIncomingServiceConnectionManager) HandleIncomingTouchStreamConnection(connection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("handleIncomingTouchStreamConnection:"), connection)
}
func (w WSHIDIncomingServiceConnectionManager) IncomingServiceConnectionDidRevoke(revoke objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("incomingServiceConnectionDidRevoke:"), revoke)
}

func (_WSHIDIncomingServiceConnectionManagerClass WSHIDIncomingServiceConnectionManagerClass) SharedInstance() WSHIDIncomingServiceConnectionManager {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_WSHIDIncomingServiceConnectionManagerClass.class), objc.Sel("sharedInstance"))
	return WSHIDIncomingServiceConnectionManagerFromID(rv)
}

func (w WSHIDIncomingServiceConnectionManager) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSHIDIncomingServiceConnectionManager) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSHIDIncomingServiceConnectionManager) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSHIDIncomingServiceConnectionManager) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSHIDEventDeliveryManagerServer] class.
var (
	_WSHIDEventDeliveryManagerServerClass     WSHIDEventDeliveryManagerServerClass
	_WSHIDEventDeliveryManagerServerClassOnce sync.Once
)

func getWSHIDEventDeliveryManagerServerClass() WSHIDEventDeliveryManagerServerClass {
	_WSHIDEventDeliveryManagerServerClassOnce.Do(func() {
		_WSHIDEventDeliveryManagerServerClass = WSHIDEventDeliveryManagerServerClass{class: objc.GetClass("WSHIDEventDeliveryManagerServer")}
	})
	return _WSHIDEventDeliveryManagerServerClass
}

// GetWSHIDEventDeliveryManagerServerClass returns the class object for WSHIDEventDeliveryManagerServer.
func GetWSHIDEventDeliveryManagerServerClass() WSHIDEventDeliveryManagerServerClass {
	return getWSHIDEventDeliveryManagerServerClass()
}

type WSHIDEventDeliveryManagerServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSHIDEventDeliveryManagerServerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSHIDEventDeliveryManagerServerClass) Alloc() WSHIDEventDeliveryManagerServer {
	rv := objc.SendIfResponds[WSHIDEventDeliveryManagerServer](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSHIDEventDeliveryManagerServer._init]
//   - [WSHIDEventDeliveryManagerServer.Activate]
//   - [WSHIDEventDeliveryManagerServer.AppendDescriptionToStream]
//   - [WSHIDEventDeliveryManagerServer.DeliveryManagerForAuditToken]
//   - [WSHIDEventDeliveryManagerServer.PermittedRuleChangeMaskForAuditToken]
//   - [WSHIDEventDeliveryManagerServer.Server]
//   - [WSHIDEventDeliveryManagerServer.DebugDescription]
//   - [WSHIDEventDeliveryManagerServer.Description]
//   - [WSHIDEventDeliveryManagerServer.Hash]
//   - [WSHIDEventDeliveryManagerServer.Superclass]
type WSHIDEventDeliveryManagerServer struct {
	objectivec.Object
}

// WSHIDEventDeliveryManagerServerFromID constructs a [WSHIDEventDeliveryManagerServer] from an objc.ID.
func WSHIDEventDeliveryManagerServerFromID(id objc.ID) WSHIDEventDeliveryManagerServer {
	return WSHIDEventDeliveryManagerServer{objectivec.Object{ID: id}}
}

// Ensure WSHIDEventDeliveryManagerServer implements IWSHIDEventDeliveryManagerServer.
var _ IWSHIDEventDeliveryManagerServer = WSHIDEventDeliveryManagerServer{}

// An interface definition for the [WSHIDEventDeliveryManagerServer] class.
//
// # Methods
//
//   - [IWSHIDEventDeliveryManagerServer._init]
//   - [IWSHIDEventDeliveryManagerServer.Activate]
//   - [IWSHIDEventDeliveryManagerServer.AppendDescriptionToStream]
//   - [IWSHIDEventDeliveryManagerServer.DeliveryManagerForAuditToken]
//   - [IWSHIDEventDeliveryManagerServer.PermittedRuleChangeMaskForAuditToken]
//   - [IWSHIDEventDeliveryManagerServer.Server]
//   - [IWSHIDEventDeliveryManagerServer.DebugDescription]
//   - [IWSHIDEventDeliveryManagerServer.Description]
//   - [IWSHIDEventDeliveryManagerServer.Hash]
//   - [IWSHIDEventDeliveryManagerServer.Superclass]
type IWSHIDEventDeliveryManagerServer interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	Activate()
	AppendDescriptionToStream(stream objectivec.IObject)
	DeliveryManagerForAuditToken(token objectivec.IObject) objectivec.IObject
	PermittedRuleChangeMaskForAuditToken(token objectivec.IObject) uint64
	Server() unsafe.Pointer
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSHIDEventDeliveryManagerServer) Init() WSHIDEventDeliveryManagerServer {
	rv := objc.SendIfResponds[WSHIDEventDeliveryManagerServer](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSHIDEventDeliveryManagerServer) Autorelease() WSHIDEventDeliveryManagerServer {
	rv := objc.SendIfResponds[WSHIDEventDeliveryManagerServer](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSHIDEventDeliveryManagerServer creates a new WSHIDEventDeliveryManagerServer instance.
func NewWSHIDEventDeliveryManagerServer() WSHIDEventDeliveryManagerServer {
	class := getWSHIDEventDeliveryManagerServerClass()
	rv := objc.SendIfResponds[WSHIDEventDeliveryManagerServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSHIDEventDeliveryManagerServer) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (w WSHIDEventDeliveryManagerServer) Activate() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("activate"))
}
func (w WSHIDEventDeliveryManagerServer) AppendDescriptionToStream(stream objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("appendDescriptionToStream:"), stream)
}
func (w WSHIDEventDeliveryManagerServer) DeliveryManagerForAuditToken(token objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("deliveryManagerForAuditToken:"), token)
	return objectivec.Object{ID: rv}
}
func (w WSHIDEventDeliveryManagerServer) PermittedRuleChangeMaskForAuditToken(token objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("permittedRuleChangeMaskForAuditToken:"), token)
	return rv
}

func (_WSHIDEventDeliveryManagerServerClass WSHIDEventDeliveryManagerServerClass) SharedInstance() WSHIDEventDeliveryManagerServer {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_WSHIDEventDeliveryManagerServerClass.class), objc.Sel("sharedInstance"))
	return WSHIDEventDeliveryManagerServerFromID(rv)
}

func (w WSHIDEventDeliveryManagerServer) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSHIDEventDeliveryManagerServer) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSHIDEventDeliveryManagerServer) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSHIDEventDeliveryManagerServer) Server() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("server"))
	return rv
}
func (w WSHIDEventDeliveryManagerServer) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

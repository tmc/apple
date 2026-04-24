// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSHIDEventDeliveryObserverServer] class.
var (
	_WSHIDEventDeliveryObserverServerClass     WSHIDEventDeliveryObserverServerClass
	_WSHIDEventDeliveryObserverServerClassOnce sync.Once
)

func getWSHIDEventDeliveryObserverServerClass() WSHIDEventDeliveryObserverServerClass {
	_WSHIDEventDeliveryObserverServerClassOnce.Do(func() {
		_WSHIDEventDeliveryObserverServerClass = WSHIDEventDeliveryObserverServerClass{class: objc.GetClass("WSHIDEventDeliveryObserverServer")}
	})
	return _WSHIDEventDeliveryObserverServerClass
}

// GetWSHIDEventDeliveryObserverServerClass returns the class object for WSHIDEventDeliveryObserverServer.
func GetWSHIDEventDeliveryObserverServerClass() WSHIDEventDeliveryObserverServerClass {
	return getWSHIDEventDeliveryObserverServerClass()
}

type WSHIDEventDeliveryObserverServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSHIDEventDeliveryObserverServerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSHIDEventDeliveryObserverServerClass) Alloc() WSHIDEventDeliveryObserverServer {
	rv := objc.Send[WSHIDEventDeliveryObserverServer](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSHIDEventDeliveryObserverServer._init]
//   - [WSHIDEventDeliveryObserverServer.Activate]
//   - [WSHIDEventDeliveryObserverServer.AppendDescriptionToStream]
//   - [WSHIDEventDeliveryObserverServer.BkServer]
//   - [WSHIDEventDeliveryObserverServer.DeliveryObserverServiceForAuditToken]
//   - [WSHIDEventDeliveryObserverServer.DebugDescription]
//   - [WSHIDEventDeliveryObserverServer.Description]
//   - [WSHIDEventDeliveryObserverServer.Hash]
//   - [WSHIDEventDeliveryObserverServer.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer
type WSHIDEventDeliveryObserverServer struct {
	objectivec.Object
}

// WSHIDEventDeliveryObserverServerFromID constructs a [WSHIDEventDeliveryObserverServer] from an objc.ID.
func WSHIDEventDeliveryObserverServerFromID(id objc.ID) WSHIDEventDeliveryObserverServer {
	return WSHIDEventDeliveryObserverServer{objectivec.Object{ID: id}}
}

// Ensure WSHIDEventDeliveryObserverServer implements IWSHIDEventDeliveryObserverServer.
var _ IWSHIDEventDeliveryObserverServer = WSHIDEventDeliveryObserverServer{}

// An interface definition for the [WSHIDEventDeliveryObserverServer] class.
//
// # Methods
//
//   - [IWSHIDEventDeliveryObserverServer._init]
//   - [IWSHIDEventDeliveryObserverServer.Activate]
//   - [IWSHIDEventDeliveryObserverServer.AppendDescriptionToStream]
//   - [IWSHIDEventDeliveryObserverServer.BkServer]
//   - [IWSHIDEventDeliveryObserverServer.DeliveryObserverServiceForAuditToken]
//   - [IWSHIDEventDeliveryObserverServer.DebugDescription]
//   - [IWSHIDEventDeliveryObserverServer.Description]
//   - [IWSHIDEventDeliveryObserverServer.Hash]
//   - [IWSHIDEventDeliveryObserverServer.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer
type IWSHIDEventDeliveryObserverServer interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	Activate()
	AppendDescriptionToStream(stream objectivec.IObject)
	BkServer() unsafe.Pointer
	DeliveryObserverServiceForAuditToken(token objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (w WSHIDEventDeliveryObserverServer) Init() WSHIDEventDeliveryObserverServer {
	rv := objc.Send[WSHIDEventDeliveryObserverServer](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSHIDEventDeliveryObserverServer) Autorelease() WSHIDEventDeliveryObserverServer {
	rv := objc.Send[WSHIDEventDeliveryObserverServer](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSHIDEventDeliveryObserverServer creates a new WSHIDEventDeliveryObserverServer instance.
func NewWSHIDEventDeliveryObserverServer() WSHIDEventDeliveryObserverServer {
	class := getWSHIDEventDeliveryObserverServerClass()
	rv := objc.Send[WSHIDEventDeliveryObserverServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer/_init
func (w WSHIDEventDeliveryObserverServer) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer/activate
func (w WSHIDEventDeliveryObserverServer) Activate() {
	objc.Send[objc.ID](w.ID, objc.Sel("activate"))
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer/appendDescriptionToStream:
func (w WSHIDEventDeliveryObserverServer) AppendDescriptionToStream(stream objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("appendDescriptionToStream:"), stream)
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer/deliveryObserverServiceForAuditToken:
func (w WSHIDEventDeliveryObserverServer) DeliveryObserverServiceForAuditToken(token objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("deliveryObserverServiceForAuditToken:"), token)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer/sharedInstance
func (_WSHIDEventDeliveryObserverServerClass WSHIDEventDeliveryObserverServerClass) SharedInstance() WSHIDEventDeliveryObserverServer {
	rv := objc.Send[objc.ID](objc.ID(_WSHIDEventDeliveryObserverServerClass.class), objc.Sel("sharedInstance"))
	return WSHIDEventDeliveryObserverServerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer/bkServer
func (w WSHIDEventDeliveryObserverServer) BkServer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w.ID, objc.Sel("bkServer"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer/debugDescription
func (w WSHIDEventDeliveryObserverServer) DebugDescription() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer/description
func (w WSHIDEventDeliveryObserverServer) Description() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer/hash
func (w WSHIDEventDeliveryObserverServer) Hash() uint64 {
	rv := objc.Send[uint64](w.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSHIDEventDeliveryObserverServer/superclass
func (w WSHIDEventDeliveryObserverServer) Superclass() objc.Class {
	rv := objc.Send[objc.Class](w.ID, objc.Sel("superclass"))
	return rv
}

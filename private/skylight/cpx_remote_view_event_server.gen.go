// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXRemoteViewEventServer] class.
var (
	_CPXRemoteViewEventServerClass     CPXRemoteViewEventServerClass
	_CPXRemoteViewEventServerClassOnce sync.Once
)

func getCPXRemoteViewEventServerClass() CPXRemoteViewEventServerClass {
	_CPXRemoteViewEventServerClassOnce.Do(func() {
		_CPXRemoteViewEventServerClass = CPXRemoteViewEventServerClass{class: objc.GetClass("CPXRemoteViewEventServer")}
	})
	return _CPXRemoteViewEventServerClass
}

// GetCPXRemoteViewEventServerClass returns the class object for CPXRemoteViewEventServer.
func GetCPXRemoteViewEventServerClass() CPXRemoteViewEventServerClass {
	return getCPXRemoteViewEventServerClass()
}

type CPXRemoteViewEventServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXRemoteViewEventServerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXRemoteViewEventServerClass) Alloc() CPXRemoteViewEventServer {
	rv := objc.SendIfResponds[CPXRemoteViewEventServer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXRemoteViewEventServer.DidUpdateRemoteViewEventManagerForSession]
//   - [CPXRemoteViewEventServer.Endpoint]
//   - [CPXRemoteViewEventServer.Invalidate]
//   - [CPXRemoteViewEventServer.ListenerDidReceiveConnectionWithContext]
//   - [CPXRemoteViewEventServer.PendingConnectionsCount]
//   - [CPXRemoteViewEventServer.InitWithConfig]
//   - [CPXRemoteViewEventServer.DebugDescription]
//   - [CPXRemoteViewEventServer.Description]
//   - [CPXRemoteViewEventServer.Hash]
//   - [CPXRemoteViewEventServer.Superclass]
type CPXRemoteViewEventServer struct {
	objectivec.Object
}

// CPXRemoteViewEventServerFromID constructs a [CPXRemoteViewEventServer] from an objc.ID.
func CPXRemoteViewEventServerFromID(id objc.ID) CPXRemoteViewEventServer {
	return CPXRemoteViewEventServer{objectivec.Object{ID: id}}
}

// Ensure CPXRemoteViewEventServer implements ICPXRemoteViewEventServer.
var _ ICPXRemoteViewEventServer = CPXRemoteViewEventServer{}

// An interface definition for the [CPXRemoteViewEventServer] class.
//
// # Methods
//
//   - [ICPXRemoteViewEventServer.DidUpdateRemoteViewEventManagerForSession]
//   - [ICPXRemoteViewEventServer.Endpoint]
//   - [ICPXRemoteViewEventServer.Invalidate]
//   - [ICPXRemoteViewEventServer.ListenerDidReceiveConnectionWithContext]
//   - [ICPXRemoteViewEventServer.PendingConnectionsCount]
//   - [ICPXRemoteViewEventServer.InitWithConfig]
//   - [ICPXRemoteViewEventServer.DebugDescription]
//   - [ICPXRemoteViewEventServer.Description]
//   - [ICPXRemoteViewEventServer.Hash]
//   - [ICPXRemoteViewEventServer.Superclass]
type ICPXRemoteViewEventServer interface {
	objectivec.IObject

	// Topic: Methods

	DidUpdateRemoteViewEventManagerForSession()
	Endpoint() objectivec.IObject
	Invalidate()
	ListenerDidReceiveConnectionWithContext(listener objectivec.IObject, connection objectivec.IObject, context objectivec.IObject)
	PendingConnectionsCount() uint64
	InitWithConfig(config objectivec.IObject) CPXRemoteViewEventServer
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (c CPXRemoteViewEventServer) Init() CPXRemoteViewEventServer {
	rv := objc.SendIfResponds[CPXRemoteViewEventServer](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXRemoteViewEventServer) Autorelease() CPXRemoteViewEventServer {
	rv := objc.SendIfResponds[CPXRemoteViewEventServer](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXRemoteViewEventServer creates a new CPXRemoteViewEventServer instance.
func NewCPXRemoteViewEventServer() CPXRemoteViewEventServer {
	class := getCPXRemoteViewEventServerClass()
	rv := objc.SendIfResponds[CPXRemoteViewEventServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCPXRemoteViewEventServerWithConfig(config objectivec.IObject) CPXRemoteViewEventServer {
	instance := getCPXRemoteViewEventServerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithConfig:"), config)
	return CPXRemoteViewEventServerFromID(rv)
}

func (c CPXRemoteViewEventServer) DidUpdateRemoteViewEventManagerForSession() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("didUpdateRemoteViewEventManagerForSession"))
}
func (c CPXRemoteViewEventServer) Endpoint() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("endpoint"))
	return objectivec.Object{ID: rv}
}
func (c CPXRemoteViewEventServer) Invalidate() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("invalidate"))
}
func (c CPXRemoteViewEventServer) ListenerDidReceiveConnectionWithContext(listener objectivec.IObject, connection objectivec.IObject, context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("listener:didReceiveConnection:withContext:"), listener, connection, context)
}
func (c CPXRemoteViewEventServer) PendingConnectionsCount() uint64 {
	rv := objc.SendIfResponds[uint64](c.ID, objc.Sel("pendingConnectionsCount"))
	return rv
}
func (c CPXRemoteViewEventServer) InitWithConfig(config objectivec.IObject) CPXRemoteViewEventServer {
	rv := objc.SendIfResponds[CPXRemoteViewEventServer](c.ID, objc.Sel("initWithConfig:"), config)
	return rv
}

func (_CPXRemoteViewEventServerClass CPXRemoteViewEventServerClass) SharedInstance() CPXRemoteViewEventServer {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_CPXRemoteViewEventServerClass.class), objc.Sel("sharedInstance"))
	return CPXRemoteViewEventServerFromID(rv)
}

func (c CPXRemoteViewEventServer) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXRemoteViewEventServer) Description() string {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXRemoteViewEventServer) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](c.ID, objc.Sel("hash"))
	return rv
}
func (c CPXRemoteViewEventServer) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](c.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWConnection] class.
var (
	_NWConnectionClass     NWConnectionClass
	_NWConnectionClassOnce sync.Once
)

func getNWConnectionClass() NWConnectionClass {
	_NWConnectionClassOnce.Do(func() {
		_NWConnectionClass = NWConnectionClass{class: objc.GetClass("NWConnection")}
	})
	return _NWConnectionClass
}

// GetNWConnectionClass returns the class object for NWConnection.
func GetNWConnectionClass() NWConnectionClass {
	return getNWConnectionClass()
}

type NWConnectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWConnectionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWConnectionClass) Alloc() NWConnection {
	rv := objc.Send[NWConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [NWConnection.Cancel]
//   - [NWConnection.CancelCurrentEndpoint]
//   - [NWConnection.ConnectedLocalEndpoint]
//   - [NWConnection.ConnectedRemoteEndpoint]
//   - [NWConnection.ConnectionState]
//   - [NWConnection.CopyCurrentPath]
//   - [NWConnection.CopyError]
//   - [NWConnection.CurrentPath]
//   - [NWConnection.Endpoint]
//   - [NWConnection.Error]
//   - [NWConnection.ForceCancel]
//   - [NWConnection.GetConnectedSocket]
//   - [NWConnection.HasBetterPath]
//   - [NWConnection.InternalConnection]
//   - [NWConnection.SetInternalConnection]
//   - [NWConnection.InternalConnectionState]
//   - [NWConnection.SetInternalConnectionState]
//   - [NWConnection.InternalError]
//   - [NWConnection.SetInternalError]
//   - [NWConnection.InternalHasBetterPath]
//   - [NWConnection.SetInternalHasBetterPath]
//   - [NWConnection.InternalIsViable]
//   - [NWConnection.SetInternalIsViable]
//   - [NWConnection.InternalPath]
//   - [NWConnection.SetInternalPath]
//   - [NWConnection.IsViable]
//   - [NWConnection.Parameters]
//   - [NWConnection.Start]
//   - [NWConnection.TlsConnectionTime]
//   - [NWConnection.InitWithConnectedSocket]
//   - [NWConnection.InitWithEndpointParameters]
//   - [NWConnection.InitWithInternalConnection]
//   - [NWConnection.Viable]
type NWConnection struct {
	objectivec.Object
}

// NWConnectionFromID constructs a [NWConnection] from an objc.ID.
func NWConnectionFromID(id objc.ID) NWConnection {
	return NWConnection{objectivec.Object{ID: id}}
}

// Ensure NWConnection implements INWConnection.
var _ INWConnection = NWConnection{}

// An interface definition for the [NWConnection] class.
//
// # Methods
//
//   - [INWConnection.Cancel]
//   - [INWConnection.CancelCurrentEndpoint]
//   - [INWConnection.ConnectedLocalEndpoint]
//   - [INWConnection.ConnectedRemoteEndpoint]
//   - [INWConnection.ConnectionState]
//   - [INWConnection.CopyCurrentPath]
//   - [INWConnection.CopyError]
//   - [INWConnection.CurrentPath]
//   - [INWConnection.Endpoint]
//   - [INWConnection.Error]
//   - [INWConnection.ForceCancel]
//   - [INWConnection.GetConnectedSocket]
//   - [INWConnection.HasBetterPath]
//   - [INWConnection.InternalConnection]
//   - [INWConnection.SetInternalConnection]
//   - [INWConnection.InternalConnectionState]
//   - [INWConnection.SetInternalConnectionState]
//   - [INWConnection.InternalError]
//   - [INWConnection.SetInternalError]
//   - [INWConnection.InternalHasBetterPath]
//   - [INWConnection.SetInternalHasBetterPath]
//   - [INWConnection.InternalIsViable]
//   - [INWConnection.SetInternalIsViable]
//   - [INWConnection.InternalPath]
//   - [INWConnection.SetInternalPath]
//   - [INWConnection.IsViable]
//   - [INWConnection.Parameters]
//   - [INWConnection.Start]
//   - [INWConnection.TlsConnectionTime]
//   - [INWConnection.InitWithConnectedSocket]
//   - [INWConnection.InitWithEndpointParameters]
//   - [INWConnection.InitWithInternalConnection]
//   - [INWConnection.Viable]
type INWConnection interface {
	objectivec.IObject

	// Topic: Methods

	Cancel()
	CancelCurrentEndpoint()
	ConnectedLocalEndpoint() INWEndpoint
	ConnectedRemoteEndpoint() INWEndpoint
	ConnectionState() int64
	CopyCurrentPath() objectivec.IObject
	CopyError() objectivec.IObject
	CurrentPath() INWPath
	Endpoint() INWEndpoint
	Error() foundation.NSError
	ForceCancel()
	GetConnectedSocket() int
	HasBetterPath() bool
	InternalConnection() unsafe.Pointer
	SetInternalConnection(value unsafe.Pointer)
	InternalConnectionState() int64
	SetInternalConnectionState(value int64)
	InternalError() foundation.NSError
	SetInternalError(value foundation.NSError)
	InternalHasBetterPath() bool
	SetInternalHasBetterPath(value bool)
	InternalIsViable() bool
	SetInternalIsViable(value bool)
	InternalPath() INWPath
	SetInternalPath(value INWPath)
	IsViable() bool
	Parameters() INWParameters
	Start()
	TlsConnectionTime() uint32
	InitWithConnectedSocket(socket int) NWConnection
	InitWithEndpointParameters(endpoint objectivec.IObject, parameters objectivec.IObject) NWConnection
	InitWithInternalConnection(connection objectivec.IObject) NWConnection
	Viable() bool
}

// Init initializes the instance.
func (n NWConnection) Init() NWConnection {
	rv := objc.Send[NWConnection](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWConnection) Autorelease() NWConnection {
	rv := objc.Send[NWConnection](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWConnection creates a new NWConnection instance.
func NewNWConnection() NWConnection {
	class := getNWConnectionClass()
	rv := objc.Send[NWConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNWConnectionWithConnectedSocket(socket int) NWConnection {
	instance := getNWConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConnectedSocket:"), socket)
	return NWConnectionFromID(rv)
}

func NewNWConnectionWithEndpointParameters(endpoint objectivec.IObject, parameters objectivec.IObject) NWConnection {
	instance := getNWConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEndpoint:parameters:"), endpoint, parameters)
	return NWConnectionFromID(rv)
}

func NewNWConnectionWithInternalConnection(connection objectivec.IObject) NWConnection {
	instance := getNWConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInternalConnection:"), connection)
	return NWConnectionFromID(rv)
}

func (n NWConnection) Cancel() {
	objc.Send[objc.ID](n.ID, objc.Sel("cancel"))
}
func (n NWConnection) CancelCurrentEndpoint() {
	objc.Send[objc.ID](n.ID, objc.Sel("cancelCurrentEndpoint"))
}
func (n NWConnection) CopyCurrentPath() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyCurrentPath"))
	return objectivec.Object{ID: rv}
}
func (n NWConnection) CopyError() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("copyError"))
	return objectivec.Object{ID: rv}
}
func (n NWConnection) ForceCancel() {
	objc.Send[objc.ID](n.ID, objc.Sel("forceCancel"))
}
func (n NWConnection) GetConnectedSocket() int {
	rv := objc.Send[int](n.ID, objc.Sel("getConnectedSocket"))
	return rv
}
func (n NWConnection) IsViable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isViable"))
	return rv
}
func (n NWConnection) Start() {
	objc.Send[objc.ID](n.ID, objc.Sel("start"))
}
func (n NWConnection) InitWithConnectedSocket(socket int) NWConnection {
	rv := objc.Send[NWConnection](n.ID, objc.Sel("initWithConnectedSocket:"), socket)
	return rv
}
func (n NWConnection) InitWithEndpointParameters(endpoint objectivec.IObject, parameters objectivec.IObject) NWConnection {
	rv := objc.Send[NWConnection](n.ID, objc.Sel("initWithEndpoint:parameters:"), endpoint, parameters)
	return rv
}
func (n NWConnection) InitWithInternalConnection(connection objectivec.IObject) NWConnection {
	rv := objc.Send[NWConnection](n.ID, objc.Sel("initWithInternalConnection:"), connection)
	return rv
}

func (_NWConnectionClass NWConnectionClass) AutomaticallyNotifiesObserversForKey(key objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_NWConnectionClass.class), objc.Sel("automaticallyNotifiesObserversForKey:"), key)
	return rv
}
func (_NWConnectionClass NWConnectionClass) ConnectionWithConnectedSocket(socket int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWConnectionClass.class), objc.Sel("connectionWithConnectedSocket:"), socket)
	return objectivec.Object{ID: rv}
}
func (_NWConnectionClass NWConnectionClass) ConnectionWithEndpointParameters(endpoint objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWConnectionClass.class), objc.Sel("connectionWithEndpoint:parameters:"), endpoint, parameters)
	return objectivec.Object{ID: rv}
}
func (_NWConnectionClass NWConnectionClass) ConnectionWithInternalConnection(connection objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NWConnectionClass.class), objc.Sel("connectionWithInternalConnection:"), connection)
	return objectivec.Object{ID: rv}
}

func (n NWConnection) ConnectedLocalEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("connectedLocalEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}
func (n NWConnection) ConnectedRemoteEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("connectedRemoteEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}
func (n NWConnection) ConnectionState() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("connectionState"))
	return rv
}
func (n NWConnection) CurrentPath() INWPath {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("currentPath"))
	return NWPathFromID(objc.ID(rv))
}
func (n NWConnection) Endpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("endpoint"))
	return NWEndpointFromID(objc.ID(rv))
}
func (n NWConnection) Error() foundation.NSError {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (n NWConnection) HasBetterPath() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasBetterPath"))
	return rv
}
func (n NWConnection) InternalConnection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n.ID, objc.Sel("internalConnection"))
	return rv
}
func (n NWConnection) SetInternalConnection(value unsafe.Pointer) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalConnection:"), value)
}
func (n NWConnection) InternalConnectionState() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("internalConnectionState"))
	return rv
}
func (n NWConnection) SetInternalConnectionState(value int64) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalConnectionState:"), value)
}
func (n NWConnection) InternalError() foundation.NSError {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("internalError"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (n NWConnection) SetInternalError(value foundation.NSError) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalError:"), value)
}
func (n NWConnection) InternalHasBetterPath() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("internalHasBetterPath"))
	return rv
}
func (n NWConnection) SetInternalHasBetterPath(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalHasBetterPath:"), value)
}
func (n NWConnection) InternalIsViable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("internalIsViable"))
	return rv
}
func (n NWConnection) SetInternalIsViable(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalIsViable:"), value)
}
func (n NWConnection) InternalPath() INWPath {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("internalPath"))
	return NWPathFromID(objc.ID(rv))
}
func (n NWConnection) SetInternalPath(value INWPath) {
	objc.Send[struct{}](n.ID, objc.Sel("setInternalPath:"), value)
}
func (n NWConnection) Parameters() INWParameters {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parameters"))
	return NWParametersFromID(objc.ID(rv))
}
func (n NWConnection) TlsConnectionTime() uint32 {
	rv := objc.Send[uint32](n.ID, objc.Sel("tlsConnectionTime"))
	return rv
}
func (n NWConnection) Viable() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("viable"))
	return rv
}

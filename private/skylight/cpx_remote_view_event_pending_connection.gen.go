// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXRemoteViewEventPendingConnection] class.
var (
	_CPXRemoteViewEventPendingConnectionClass     CPXRemoteViewEventPendingConnectionClass
	_CPXRemoteViewEventPendingConnectionClassOnce sync.Once
)

func getCPXRemoteViewEventPendingConnectionClass() CPXRemoteViewEventPendingConnectionClass {
	_CPXRemoteViewEventPendingConnectionClassOnce.Do(func() {
		_CPXRemoteViewEventPendingConnectionClass = CPXRemoteViewEventPendingConnectionClass{class: objc.GetClass("_CPXRemoteViewEventPendingConnection")}
	})
	return _CPXRemoteViewEventPendingConnectionClass
}

// GetCPXRemoteViewEventPendingConnectionClass returns the class object for _CPXRemoteViewEventPendingConnection.
func GetCPXRemoteViewEventPendingConnectionClass() CPXRemoteViewEventPendingConnectionClass {
	return getCPXRemoteViewEventPendingConnectionClass()
}

type CPXRemoteViewEventPendingConnectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXRemoteViewEventPendingConnectionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXRemoteViewEventPendingConnectionClass) Alloc() CPXRemoteViewEventPendingConnection {
	rv := objc.Send[CPXRemoteViewEventPendingConnection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXRemoteViewEventPendingConnection.AcceptConnection]
//   - [CPXRemoteViewEventPendingConnection.Activate]
//   - [CPXRemoteViewEventPendingConnection.Connection]
//   - [CPXRemoteViewEventPendingConnection.ConnectionRevokedWithEvent]
//   - [CPXRemoteViewEventPendingConnection.Handler]
//   - [CPXRemoteViewEventPendingConnection.SetHandler]
//   - [CPXRemoteViewEventPendingConnection.RejectConnection]
//   - [CPXRemoteViewEventPendingConnection.InitWithConnectionHandler]
//   - [CPXRemoteViewEventPendingConnection.DebugDescription]
//   - [CPXRemoteViewEventPendingConnection.Description]
//   - [CPXRemoteViewEventPendingConnection.Hash]
//   - [CPXRemoteViewEventPendingConnection.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection
type CPXRemoteViewEventPendingConnection struct {
	objectivec.Object
}

// CPXRemoteViewEventPendingConnectionFromID constructs a [CPXRemoteViewEventPendingConnection] from an objc.ID.
func CPXRemoteViewEventPendingConnectionFromID(id objc.ID) CPXRemoteViewEventPendingConnection {
	return CPXRemoteViewEventPendingConnection{objectivec.Object{ID: id}}
}

// Ensure CPXRemoteViewEventPendingConnection implements ICPXRemoteViewEventPendingConnection.
var _ ICPXRemoteViewEventPendingConnection = CPXRemoteViewEventPendingConnection{}

// An interface definition for the [CPXRemoteViewEventPendingConnection] class.
//
// # Methods
//
//   - [ICPXRemoteViewEventPendingConnection.AcceptConnection]
//   - [ICPXRemoteViewEventPendingConnection.Activate]
//   - [ICPXRemoteViewEventPendingConnection.Connection]
//   - [ICPXRemoteViewEventPendingConnection.ConnectionRevokedWithEvent]
//   - [ICPXRemoteViewEventPendingConnection.Handler]
//   - [ICPXRemoteViewEventPendingConnection.SetHandler]
//   - [ICPXRemoteViewEventPendingConnection.RejectConnection]
//   - [ICPXRemoteViewEventPendingConnection.InitWithConnectionHandler]
//   - [ICPXRemoteViewEventPendingConnection.DebugDescription]
//   - [ICPXRemoteViewEventPendingConnection.Description]
//   - [ICPXRemoteViewEventPendingConnection.Hash]
//   - [ICPXRemoteViewEventPendingConnection.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection
type ICPXRemoteViewEventPendingConnection interface {
	objectivec.IObject

	// Topic: Methods

	AcceptConnection()
	Activate() bool
	Connection() unsafe.Pointer
	ConnectionRevokedWithEvent(connection objectivec.IObject, event objectivec.IObject)
	Handler() ICPXRemoteViewEventServer
	SetHandler(value ICPXRemoteViewEventServer)
	RejectConnection()
	InitWithConnectionHandler(connection objectivec.IObject, handler ErrorHandler) CPXRemoteViewEventPendingConnection
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXRemoteViewEventPendingConnection) Init() CPXRemoteViewEventPendingConnection {
	rv := objc.Send[CPXRemoteViewEventPendingConnection](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXRemoteViewEventPendingConnection) Autorelease() CPXRemoteViewEventPendingConnection {
	rv := objc.Send[CPXRemoteViewEventPendingConnection](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXRemoteViewEventPendingConnection creates a new CPXRemoteViewEventPendingConnection instance.
func NewCPXRemoteViewEventPendingConnection() CPXRemoteViewEventPendingConnection {
	class := getCPXRemoteViewEventPendingConnectionClass()
	rv := objc.Send[CPXRemoteViewEventPendingConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/initWithConnection:handler:
func NewCPXRemoteViewEventPendingConnectionWithConnectionHandler(connection objectivec.IObject, handler objectivec.IObject) CPXRemoteViewEventPendingConnection {
	instance := getCPXRemoteViewEventPendingConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConnection:handler:"), connection, handler)
	return CPXRemoteViewEventPendingConnectionFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/acceptConnection
func (c CPXRemoteViewEventPendingConnection) AcceptConnection() {
	objc.Send[objc.ID](c.ID, objc.Sel("acceptConnection"))
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/activate
func (c CPXRemoteViewEventPendingConnection) Activate() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("activate"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/connection:revokedWithEvent:
func (c CPXRemoteViewEventPendingConnection) ConnectionRevokedWithEvent(connection objectivec.IObject, event objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("connection:revokedWithEvent:"), connection, event)
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/rejectConnection
func (c CPXRemoteViewEventPendingConnection) RejectConnection() {
	objc.Send[objc.ID](c.ID, objc.Sel("rejectConnection"))
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/initWithConnection:handler:
func (c CPXRemoteViewEventPendingConnection) InitWithConnectionHandler(connection objectivec.IObject, handler ErrorHandler) CPXRemoteViewEventPendingConnection {
	_block1, _ := NewErrorBlock(handler)
	rv := objc.Send[CPXRemoteViewEventPendingConnection](c.ID, objc.Sel("initWithConnection:handler:"), connection, _block1)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/connection
func (c CPXRemoteViewEventPendingConnection) Connection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("connection"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/debugDescription
func (c CPXRemoteViewEventPendingConnection) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/description
func (c CPXRemoteViewEventPendingConnection) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/handler
func (c CPXRemoteViewEventPendingConnection) Handler() ICPXRemoteViewEventServer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("handler"))
	return CPXRemoteViewEventServerFromID(objc.ID(rv))
}
func (c CPXRemoteViewEventPendingConnection) SetHandler(value ICPXRemoteViewEventServer) {
	objc.Send[struct{}](c.ID, objc.Sel("setHandler:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/hash
func (c CPXRemoteViewEventPendingConnection) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXRemoteViewEventPendingConnection/superclass
func (c CPXRemoteViewEventPendingConnection) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

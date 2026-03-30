// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIBaseXPCHandler] class.
var (
	_DIBaseXPCHandlerClass     DIBaseXPCHandlerClass
	_DIBaseXPCHandlerClassOnce sync.Once
)

func getDIBaseXPCHandlerClass() DIBaseXPCHandlerClass {
	_DIBaseXPCHandlerClassOnce.Do(func() {
		_DIBaseXPCHandlerClass = DIBaseXPCHandlerClass{class: objc.GetClass("DIBaseXPCHandler")}
	})
	return _DIBaseXPCHandlerClass
}

// GetDIBaseXPCHandlerClass returns the class object for DIBaseXPCHandler.
func GetDIBaseXPCHandlerClass() DIBaseXPCHandlerClass {
	return getDIBaseXPCHandlerClass()
}

type DIBaseXPCHandlerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIBaseXPCHandlerClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIBaseXPCHandlerClass) Alloc() DIBaseXPCHandler {
	rv := objc.Send[DIBaseXPCHandler](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIBaseXPCHandler.CloseConnection]
//   - [DIBaseXPCHandler.CompleteCommandWithError]
//   - [DIBaseXPCHandler.ConnectWithError]
//   - [DIBaseXPCHandler.Connection]
//   - [DIBaseXPCHandler.SetConnection]
//   - [DIBaseXPCHandler.CreateConnection]
//   - [DIBaseXPCHandler.DupStderrWithError]
//   - [DIBaseXPCHandler.IsPrivileged]
//   - [DIBaseXPCHandler.SetIsPrivileged]
//   - [DIBaseXPCHandler.RemoteObjectInterface]
//   - [DIBaseXPCHandler.RemoteProxy]
//   - [DIBaseXPCHandler.SetRemoteProxy]
//   - [DIBaseXPCHandler.Semaphore]
//   - [DIBaseXPCHandler.SetSemaphore]
//   - [DIBaseXPCHandler.ServiceName]
//   - [DIBaseXPCHandler.SignalCommandCompletedWithXpcError]
//   - [DIBaseXPCHandler.XpcError]
//   - [DIBaseXPCHandler.SetXpcError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler
type DIBaseXPCHandler struct {
	objectivec.Object
}

// DIBaseXPCHandlerFromID constructs a [DIBaseXPCHandler] from an objc.ID.
func DIBaseXPCHandlerFromID(id objc.ID) DIBaseXPCHandler {
	return DIBaseXPCHandler{objectivec.Object{ID: id}}
}

// Ensure DIBaseXPCHandler implements IDIBaseXPCHandler.
var _ IDIBaseXPCHandler = DIBaseXPCHandler{}

// An interface definition for the [DIBaseXPCHandler] class.
//
// # Methods
//
//   - [IDIBaseXPCHandler.CloseConnection]
//   - [IDIBaseXPCHandler.CompleteCommandWithError]
//   - [IDIBaseXPCHandler.ConnectWithError]
//   - [IDIBaseXPCHandler.Connection]
//   - [IDIBaseXPCHandler.SetConnection]
//   - [IDIBaseXPCHandler.CreateConnection]
//   - [IDIBaseXPCHandler.DupStderrWithError]
//   - [IDIBaseXPCHandler.IsPrivileged]
//   - [IDIBaseXPCHandler.SetIsPrivileged]
//   - [IDIBaseXPCHandler.RemoteObjectInterface]
//   - [IDIBaseXPCHandler.RemoteProxy]
//   - [IDIBaseXPCHandler.SetRemoteProxy]
//   - [IDIBaseXPCHandler.Semaphore]
//   - [IDIBaseXPCHandler.SetSemaphore]
//   - [IDIBaseXPCHandler.ServiceName]
//   - [IDIBaseXPCHandler.SignalCommandCompletedWithXpcError]
//   - [IDIBaseXPCHandler.XpcError]
//   - [IDIBaseXPCHandler.SetXpcError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler
type IDIBaseXPCHandler interface {
	objectivec.IObject

	// Topic: Methods

	CloseConnection()
	CompleteCommandWithError() (bool, error)
	ConnectWithError() (bool, error)
	Connection() foundation.NSXPCConnection
	SetConnection(value foundation.NSXPCConnection)
	CreateConnection()
	DupStderrWithError() (bool, error)
	IsPrivileged() bool
	SetIsPrivileged(value bool)
	RemoteObjectInterface() objectivec.IObject
	RemoteProxy() objectivec.IObject
	SetRemoteProxy(value objectivec.IObject)
	Semaphore() objectivec.Object
	SetSemaphore(value objectivec.Object)
	ServiceName() objectivec.IObject
	SignalCommandCompletedWithXpcError(error_ objectivec.IObject)
	XpcError() foundation.INSError
	SetXpcError(value foundation.INSError)
}

// Init initializes the instance.
func (d DIBaseXPCHandler) Init() DIBaseXPCHandler {
	rv := objc.Send[DIBaseXPCHandler](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIBaseXPCHandler) Autorelease() DIBaseXPCHandler {
	rv := objc.Send[DIBaseXPCHandler](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIBaseXPCHandler creates a new DIBaseXPCHandler instance.
func NewDIBaseXPCHandler() DIBaseXPCHandler {
	class := getDIBaseXPCHandlerClass()
	rv := objc.Send[DIBaseXPCHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/closeConnection
func (d DIBaseXPCHandler) CloseConnection() {
	objc.Send[objc.ID](d.ID, objc.Sel("closeConnection"))
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/completeCommandWithError:
func (d DIBaseXPCHandler) CompleteCommandWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("completeCommandWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("completeCommandWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/connectWithError:
func (d DIBaseXPCHandler) ConnectWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("connectWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("connectWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/createConnection
func (d DIBaseXPCHandler) CreateConnection() {
	objc.Send[objc.ID](d.ID, objc.Sel("createConnection"))
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/dupStderrWithError:
func (d DIBaseXPCHandler) DupStderrWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("dupStderrWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("dupStderrWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/remoteObjectInterface
func (d DIBaseXPCHandler) RemoteObjectInterface() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("remoteObjectInterface"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/serviceName
func (d DIBaseXPCHandler) ServiceName() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("serviceName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/signalCommandCompletedWithXpcError:
func (d DIBaseXPCHandler) SignalCommandCompletedWithXpcError(error_ objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("signalCommandCompletedWithXpcError:"), error_)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/connection
func (d DIBaseXPCHandler) Connection() foundation.NSXPCConnection {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("connection"))
	return foundation.NSXPCConnectionFromID(objc.ID(rv))
}
func (d DIBaseXPCHandler) SetConnection(value foundation.NSXPCConnection) {
	objc.Send[struct{}](d.ID, objc.Sel("setConnection:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/isPrivileged
func (d DIBaseXPCHandler) IsPrivileged() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isPrivileged"))
	return rv
}
func (d DIBaseXPCHandler) SetIsPrivileged(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setIsPrivileged:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/remoteProxy
func (d DIBaseXPCHandler) RemoteProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("remoteProxy"))
	return objectivec.Object{ID: rv}
}
func (d DIBaseXPCHandler) SetRemoteProxy(value objectivec.IObject) {
	objc.Send[struct{}](d.ID, objc.Sel("setRemoteProxy:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/semaphore
func (d DIBaseXPCHandler) Semaphore() objectivec.Object {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("semaphore"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (d DIBaseXPCHandler) SetSemaphore(value objectivec.Object) {
	objc.Send[struct{}](d.ID, objc.Sel("setSemaphore:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseXPCHandler/xpcError
func (d DIBaseXPCHandler) XpcError() foundation.INSError {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("xpcError"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (d DIBaseXPCHandler) SetXpcError(value foundation.INSError) {
	objc.Send[struct{}](d.ID, objc.Sel("setXpcError:"), value)
}

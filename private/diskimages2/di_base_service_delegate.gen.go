// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIBaseServiceDelegate] class.
var (
	_DIBaseServiceDelegateClass     DIBaseServiceDelegateClass
	_DIBaseServiceDelegateClassOnce sync.Once
)

func getDIBaseServiceDelegateClass() DIBaseServiceDelegateClass {
	_DIBaseServiceDelegateClassOnce.Do(func() {
		_DIBaseServiceDelegateClass = DIBaseServiceDelegateClass{class: objc.GetClass("DIBaseServiceDelegate")}
	})
	return _DIBaseServiceDelegateClass
}

// GetDIBaseServiceDelegateClass returns the class object for DIBaseServiceDelegate.
func GetDIBaseServiceDelegateClass() DIBaseServiceDelegateClass {
	return getDIBaseServiceDelegateClass()
}

type DIBaseServiceDelegateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIBaseServiceDelegateClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIBaseServiceDelegateClass) Alloc() DIBaseServiceDelegate {
	rv := objc.Send[DIBaseServiceDelegate](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIBaseServiceDelegate.CreateListener]
//   - [DIBaseServiceDelegate.DispatchQueue]
//   - [DIBaseServiceDelegate.EnterSandbox]
//   - [DIBaseServiceDelegate.Listener]
//   - [DIBaseServiceDelegate.SetListener]
//   - [DIBaseServiceDelegate.ListenerShouldAcceptNewConnection]
//   - [DIBaseServiceDelegate.SandboxProfile]
//   - [DIBaseServiceDelegate.ServiceName]
//   - [DIBaseServiceDelegate.SetupNewConnection]
//   - [DIBaseServiceDelegate.StartXPClistener]
//   - [DIBaseServiceDelegate.ValidateConnection]
//   - [DIBaseServiceDelegate.DebugDescription]
//   - [DIBaseServiceDelegate.Description]
//   - [DIBaseServiceDelegate.Hash]
//   - [DIBaseServiceDelegate.Superclass]
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate
type DIBaseServiceDelegate struct {
	objectivec.Object
}

// DIBaseServiceDelegateFromID constructs a [DIBaseServiceDelegate] from an objc.ID.
func DIBaseServiceDelegateFromID(id objc.ID) DIBaseServiceDelegate {
	return DIBaseServiceDelegate{objectivec.Object{ID: id}}
}
// Ensure DIBaseServiceDelegate implements IDIBaseServiceDelegate.
var _ IDIBaseServiceDelegate = DIBaseServiceDelegate{}

// An interface definition for the [DIBaseServiceDelegate] class.
//
// # Methods
//
//   - [IDIBaseServiceDelegate.CreateListener]
//   - [IDIBaseServiceDelegate.DispatchQueue]
//   - [IDIBaseServiceDelegate.EnterSandbox]
//   - [IDIBaseServiceDelegate.Listener]
//   - [IDIBaseServiceDelegate.SetListener]
//   - [IDIBaseServiceDelegate.ListenerShouldAcceptNewConnection]
//   - [IDIBaseServiceDelegate.SandboxProfile]
//   - [IDIBaseServiceDelegate.ServiceName]
//   - [IDIBaseServiceDelegate.SetupNewConnection]
//   - [IDIBaseServiceDelegate.StartXPClistener]
//   - [IDIBaseServiceDelegate.ValidateConnection]
//   - [IDIBaseServiceDelegate.DebugDescription]
//   - [IDIBaseServiceDelegate.Description]
//   - [IDIBaseServiceDelegate.Hash]
//   - [IDIBaseServiceDelegate.Superclass]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate
type IDIBaseServiceDelegate interface {
	objectivec.IObject

	// Topic: Methods

	CreateListener()
	DispatchQueue() objectivec.Object
	EnterSandbox()
	Listener() foundation.NSXPCListener
	SetListener(value foundation.NSXPCListener)
	ListenerShouldAcceptNewConnection(listener objectivec.IObject, connection objectivec.IObject) bool
	SandboxProfile() objectivec.IObject
	ServiceName() objectivec.IObject
	SetupNewConnection(connection objectivec.IObject) bool
	StartXPClistener()
	ValidateConnection()
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (d DIBaseServiceDelegate) Init() DIBaseServiceDelegate {
	rv := objc.Send[DIBaseServiceDelegate](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIBaseServiceDelegate) Autorelease() DIBaseServiceDelegate {
	rv := objc.Send[DIBaseServiceDelegate](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIBaseServiceDelegate creates a new DIBaseServiceDelegate instance.
func NewDIBaseServiceDelegate() DIBaseServiceDelegate {
	class := getDIBaseServiceDelegateClass()
	rv := objc.Send[DIBaseServiceDelegate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/createListener
func (d DIBaseServiceDelegate) CreateListener() {
	objc.Send[objc.ID](d.ID, objc.Sel("createListener"))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/enterSandbox
func (d DIBaseServiceDelegate) EnterSandbox() {
	objc.Send[objc.ID](d.ID, objc.Sel("enterSandbox"))
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/listener:shouldAcceptNewConnection:
func (d DIBaseServiceDelegate) ListenerShouldAcceptNewConnection(listener objectivec.IObject, connection objectivec.IObject) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("listener:shouldAcceptNewConnection:"), listener, connection)
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/sandboxProfile
func (d DIBaseServiceDelegate) SandboxProfile() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("sandboxProfile"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/serviceName
func (d DIBaseServiceDelegate) ServiceName() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("serviceName"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/setupNewConnection:
func (d DIBaseServiceDelegate) SetupNewConnection(connection objectivec.IObject) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("setupNewConnection:"), connection)
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/startXPClistener
func (d DIBaseServiceDelegate) StartXPClistener() {
	objc.Send[objc.ID](d.ID, objc.Sel("startXPClistener"))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/validateConnection
func (d DIBaseServiceDelegate) ValidateConnection() {
	objc.Send[objc.ID](d.ID, objc.Sel("validateConnection"))
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/debugDescription
func (d DIBaseServiceDelegate) DebugDescription() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/description
func (d DIBaseServiceDelegate) Description() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/dispatchQueue
func (d DIBaseServiceDelegate) DispatchQueue() objectivec.Object {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/hash
func (d DIBaseServiceDelegate) Hash() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/listener
func (d DIBaseServiceDelegate) Listener() foundation.NSXPCListener {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("listener"))
	return foundation.NSXPCListenerFromID(objc.ID(rv))
}
func (d DIBaseServiceDelegate) SetListener(value foundation.NSXPCListener) {
	objc.Send[struct{}](d.ID, objc.Sel("setListener:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseServiceDelegate/superclass
func (d DIBaseServiceDelegate) Superclass() objc.Class {
	rv := objc.Send[objc.Class](d.ID, objc.Sel("superclass"))
	return rv
}


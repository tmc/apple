// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVNCServer] class.
var (
	_VZVNCServerClass     VZVNCServerClass
	_VZVNCServerClassOnce sync.Once
)

func getVZVNCServerClass() VZVNCServerClass {
	_VZVNCServerClassOnce.Do(func() {
		_VZVNCServerClass = VZVNCServerClass{class: objc.GetClass("_VZVNCServer")}
	})
	return _VZVNCServerClass
}

// GetVZVNCServerClass returns the class object for _VZVNCServer.
func GetVZVNCServerClass() VZVNCServerClass {
	return getVZVNCServerClass()
}

type VZVNCServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVNCServerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVNCServerClass) Alloc() VZVNCServer {
	rv := objc.Send[VZVNCServer](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVNCServer.Delegate]
//   - [VZVNCServer.SetDelegate]
//   - [VZVNCServer.FramebufferDidUpdateGraphicsOrientation]
//   - [VZVNCServer.FramebufferDidUpdateColorSpace]
//   - [VZVNCServer.GetDisplayProtectionOptions]
//   - [VZVNCServer.GraphicsDisplay]
//   - [VZVNCServer.SetGraphicsDisplay]
//   - [VZVNCServer.Port]
//   - [VZVNCServer.Queue]
//   - [VZVNCServer.SecurityConfiguration]
//   - [VZVNCServer.Start]
//   - [VZVNCServer.State]
//   - [VZVNCServer.SetState]
//   - [VZVNCServer.Stop]
//   - [VZVNCServer.VirtualMachine]
//   - [VZVNCServer.SetVirtualMachine]
//   - [VZVNCServer.InitWithBonjourServiceName]
//   - [VZVNCServer.InitWithBonjourServiceNameQueue]
//   - [VZVNCServer.InitWithBonjourServiceNameQueueSecurityConfiguration]
//   - [VZVNCServer.InitWithPort]
//   - [VZVNCServer.InitWithPortQueue]
//   - [VZVNCServer.InitWithPortQueueSecurityConfiguration]
//   - [VZVNCServer.DebugDescription]
//   - [VZVNCServer.Description]
//   - [VZVNCServer.Hash]
//   - [VZVNCServer.Superclass]
type VZVNCServer struct {
	objectivec.Object
}

// VZVNCServerFromID constructs a [VZVNCServer] from an objc.ID.
func VZVNCServerFromID(id objc.ID) VZVNCServer {
	return VZVNCServer{objectivec.Object{ID: id}}
}

// Ensure VZVNCServer implements IVZVNCServer.
var _ IVZVNCServer = VZVNCServer{}

// An interface definition for the [VZVNCServer] class.
//
// # Methods
//
//   - [IVZVNCServer.Delegate]
//   - [IVZVNCServer.SetDelegate]
//   - [IVZVNCServer.FramebufferDidUpdateGraphicsOrientation]
//   - [IVZVNCServer.FramebufferDidUpdateColorSpace]
//   - [IVZVNCServer.GetDisplayProtectionOptions]
//   - [IVZVNCServer.GraphicsDisplay]
//   - [IVZVNCServer.SetGraphicsDisplay]
//   - [IVZVNCServer.Port]
//   - [IVZVNCServer.Queue]
//   - [IVZVNCServer.SecurityConfiguration]
//   - [IVZVNCServer.Start]
//   - [IVZVNCServer.State]
//   - [IVZVNCServer.SetState]
//   - [IVZVNCServer.Stop]
//   - [IVZVNCServer.VirtualMachine]
//   - [IVZVNCServer.SetVirtualMachine]
//   - [IVZVNCServer.InitWithBonjourServiceName]
//   - [IVZVNCServer.InitWithBonjourServiceNameQueue]
//   - [IVZVNCServer.InitWithBonjourServiceNameQueueSecurityConfiguration]
//   - [IVZVNCServer.InitWithPort]
//   - [IVZVNCServer.InitWithPortQueue]
//   - [IVZVNCServer.InitWithPortQueueSecurityConfiguration]
//   - [IVZVNCServer.DebugDescription]
//   - [IVZVNCServer.Description]
//   - [IVZVNCServer.Hash]
//   - [IVZVNCServer.Superclass]
type IVZVNCServer interface {
	objectivec.IObject

	// Topic: Methods

	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	FramebufferDidUpdateGraphicsOrientation(framebuffer objectivec.IObject, orientation int64)
	FramebufferDidUpdateColorSpace(space objectivec.IObject)
	GetDisplayProtectionOptions() unsafe.Pointer
	GraphicsDisplay() IVZGraphicsDisplay
	SetGraphicsDisplay(value IVZGraphicsDisplay)
	Port() uint16
	Queue() objectivec.Object
	SecurityConfiguration() IVZVNCSecurityConfiguration
	Start()
	State() int64
	SetState(value int64)
	Stop()
	VirtualMachine() IVZVirtualMachine
	SetVirtualMachine(value IVZVirtualMachine)
	InitWithBonjourServiceName(name objectivec.IObject) VZVNCServer
	InitWithBonjourServiceNameQueue(name objectivec.IObject, queue objectivec.IObject) VZVNCServer
	InitWithBonjourServiceNameQueueSecurityConfiguration(name objectivec.IObject, queue objectivec.IObject, configuration objectivec.IObject) VZVNCServer
	InitWithPort(port uint16) VZVNCServer
	InitWithPortQueue(port uint16, queue objectivec.IObject) VZVNCServer
	InitWithPortQueueSecurityConfiguration(port uint16, queue objectivec.IObject, configuration objectivec.IObject) VZVNCServer
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZVNCServer) Init() VZVNCServer {
	rv := objc.Send[VZVNCServer](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVNCServer) Autorelease() VZVNCServer {
	rv := objc.Send[VZVNCServer](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVNCServer creates a new VZVNCServer instance.
func NewVZVNCServer() VZVNCServer {
	class := getVZVNCServerClass()
	rv := objc.Send[VZVNCServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZVNCServerWithBonjourServiceName(name objectivec.IObject) VZVNCServer {
	instance := getVZVNCServerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBonjourServiceName:"), name)
	return VZVNCServerFromID(rv)
}

func NewVZVNCServerWithBonjourServiceNameQueue(name objectivec.IObject, queue objectivec.IObject) VZVNCServer {
	instance := getVZVNCServerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBonjourServiceName:queue:"), name, queue)
	return VZVNCServerFromID(rv)
}

func NewVZVNCServerWithBonjourServiceNameQueueSecurityConfiguration(name objectivec.IObject, queue objectivec.IObject, configuration objectivec.IObject) VZVNCServer {
	instance := getVZVNCServerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBonjourServiceName:queue:securityConfiguration:"), name, queue, configuration)
	return VZVNCServerFromID(rv)
}

func NewVZVNCServerWithPort(port uint16) VZVNCServer {
	instance := getVZVNCServerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPort:"), port)
	return VZVNCServerFromID(rv)
}

func NewVZVNCServerWithPortQueue(port uint16, queue objectivec.IObject) VZVNCServer {
	instance := getVZVNCServerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPort:queue:"), port, queue)
	return VZVNCServerFromID(rv)
}

func NewVZVNCServerWithPortQueueSecurityConfiguration(port uint16, queue objectivec.IObject, configuration objectivec.IObject) VZVNCServer {
	instance := getVZVNCServerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPort:queue:securityConfiguration:"), port, queue, configuration)
	return VZVNCServerFromID(rv)
}

func (v VZVNCServer) FramebufferDidUpdateGraphicsOrientation(framebuffer objectivec.IObject, orientation int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("framebuffer:didUpdateGraphicsOrientation:"), framebuffer, orientation)
}
func (v VZVNCServer) FramebufferDidUpdateColorSpace(space objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("framebufferDidUpdateColorSpace:"), space)
}
func (v VZVNCServer) GetDisplayProtectionOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("getDisplayProtectionOptions"))
	return rv
}
func (v VZVNCServer) Start() {
	objc.Send[objc.ID](v.ID, objc.Sel("start"))
}
func (v VZVNCServer) Stop() {
	objc.Send[objc.ID](v.ID, objc.Sel("stop"))
}
func (v VZVNCServer) InitWithBonjourServiceName(name objectivec.IObject) VZVNCServer {
	rv := objc.Send[VZVNCServer](v.ID, objc.Sel("initWithBonjourServiceName:"), name)
	return rv
}
func (v VZVNCServer) InitWithBonjourServiceNameQueue(name objectivec.IObject, queue objectivec.IObject) VZVNCServer {
	rv := objc.Send[VZVNCServer](v.ID, objc.Sel("initWithBonjourServiceName:queue:"), name, queue)
	return rv
}
func (v VZVNCServer) InitWithBonjourServiceNameQueueSecurityConfiguration(name objectivec.IObject, queue objectivec.IObject, configuration objectivec.IObject) VZVNCServer {
	rv := objc.Send[VZVNCServer](v.ID, objc.Sel("initWithBonjourServiceName:queue:securityConfiguration:"), name, queue, configuration)
	return rv
}
func (v VZVNCServer) InitWithPort(port uint16) VZVNCServer {
	rv := objc.Send[VZVNCServer](v.ID, objc.Sel("initWithPort:"), port)
	return rv
}
func (v VZVNCServer) InitWithPortQueue(port uint16, queue objectivec.IObject) VZVNCServer {
	rv := objc.Send[VZVNCServer](v.ID, objc.Sel("initWithPort:queue:"), port, queue)
	return rv
}
func (v VZVNCServer) InitWithPortQueueSecurityConfiguration(port uint16, queue objectivec.IObject, configuration objectivec.IObject) VZVNCServer {
	rv := objc.Send[VZVNCServer](v.ID, objc.Sel("initWithPort:queue:securityConfiguration:"), port, queue, configuration)
	return rv
}

func (v VZVNCServer) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVNCServer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("delegate"))
	return rv
}
func (v VZVNCServer) SetDelegate(value unsafe.Pointer) {
	objc.Send[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}
func (v VZVNCServer) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVNCServer) GraphicsDisplay() IVZGraphicsDisplay {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("graphicsDisplay"))
	return VZGraphicsDisplayFromID(objc.ID(rv))
}
func (v VZVNCServer) SetGraphicsDisplay(value IVZGraphicsDisplay) {
	objc.Send[struct{}](v.ID, objc.Sel("setGraphicsDisplay:"), value)
}
func (v VZVNCServer) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZVNCServer) Port() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("port"))
	return rv
}
func (v VZVNCServer) Queue() objectivec.Object {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (v VZVNCServer) SecurityConfiguration() IVZVNCSecurityConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("securityConfiguration"))
	return VZVNCSecurityConfigurationFromID(objc.ID(rv))
}
func (v VZVNCServer) State() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("state"))
	return rv
}
func (v VZVNCServer) SetState(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setState:"), value)
}
func (v VZVNCServer) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (v VZVNCServer) VirtualMachine() IVZVirtualMachine {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("virtualMachine"))
	return VZVirtualMachineFromID(objc.ID(rv))
}
func (v VZVNCServer) SetVirtualMachine(value IVZVirtualMachine) {
	objc.Send[struct{}](v.ID, objc.Sel("setVirtualMachine:"), value)
}

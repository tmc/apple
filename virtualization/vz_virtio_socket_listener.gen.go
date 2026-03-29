// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioSocketListener] class.
var (
	_VZVirtioSocketListenerClass     VZVirtioSocketListenerClass
	_VZVirtioSocketListenerClassOnce sync.Once
)

func getVZVirtioSocketListenerClass() VZVirtioSocketListenerClass {
	_VZVirtioSocketListenerClassOnce.Do(func() {
		_VZVirtioSocketListenerClass = VZVirtioSocketListenerClass{class: objc.GetClass("VZVirtioSocketListener")}
	})
	return _VZVirtioSocketListenerClass
}

// GetVZVirtioSocketListenerClass returns the class object for VZVirtioSocketListener.
func GetVZVirtioSocketListenerClass() VZVirtioSocketListenerClass {
	return getVZVirtioSocketListenerClass()
}

type VZVirtioSocketListenerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSocketListenerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSocketListenerClass) Alloc() VZVirtioSocketListener {
	rv := objc.Send[VZVirtioSocketListener](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that listens for port-based connection requests from the guest
// operating system.
//
// # Overview
// 
// Use a [VZVirtioSocketListener] object to route connection requests to your
// associated delegate object. The socket listener object handles incoming
// connection requests from the guest operating system and directs them to the
// methods of its associated [VZVirtioSocketListener.Delegate] object. You may use the same listener
// object to monitor connections on multiple ports.
// 
// After creating a [VZVirtioSocketListener] object, assign a custom object to
// its [VZVirtioSocketListener.Delegate] property. The delegate must implement the
// [VZVirtioSocketListenerDelegate] protocol. To connect the listener to a
// port, call the [SetSocketListenerForPort] method of your virtual
// machine’s [VZVirtioSocketDevice] object.
//
// # Responding to new connections
//
//   - [VZVirtioSocketListener.Delegate]: The custom object you use to respond to port-based connection attempts.
//   - [VZVirtioSocketListener.SetDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketListener
type VZVirtioSocketListener struct {
	objectivec.Object
}

// VZVirtioSocketListenerFromID constructs a [VZVirtioSocketListener] from an objc.ID.
//
// An object that listens for port-based connection requests from the guest
// operating system.
func VZVirtioSocketListenerFromID(id objc.ID) VZVirtioSocketListener {
	return VZVirtioSocketListener{objectivec.Object{ID: id}}
}
// NOTE: VZVirtioSocketListener adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioSocketListener] class.
//
// # Responding to new connections
//
//   - [IVZVirtioSocketListener.Delegate]: The custom object you use to respond to port-based connection attempts.
//   - [IVZVirtioSocketListener.SetDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketListener
type IVZVirtioSocketListener interface {
	objectivec.IObject

	// Topic: Responding to new connections

	// The custom object you use to respond to port-based connection attempts.
	Delegate() VZVirtioSocketListenerDelegate
	SetDelegate(value VZVirtioSocketListenerDelegate)
}

// Init initializes the instance.
func (v VZVirtioSocketListener) Init() VZVirtioSocketListener {
	rv := objc.Send[VZVirtioSocketListener](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSocketListener) Autorelease() VZVirtioSocketListener {
	rv := objc.Send[VZVirtioSocketListener](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketListener creates a new VZVirtioSocketListener instance.
func NewVZVirtioSocketListener() VZVirtioSocketListener {
	class := getVZVirtioSocketListenerClass()
	rv := objc.Send[VZVirtioSocketListener](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The custom object you use to respond to port-based connection attempts.
//
// # Discussion
// 
// Your delegate object must conform to the [VZVirtioSocketListenerDelegate]
// protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketListener/delegate
func (v VZVirtioSocketListener) Delegate() VZVirtioSocketListenerDelegate {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return VZVirtioSocketListenerDelegateObjectFromID(rv)
}
func (v VZVirtioSocketListener) SetDelegate(value VZVirtioSocketListenerDelegate) {
	objc.Send[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}


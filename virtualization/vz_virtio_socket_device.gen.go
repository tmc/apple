// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioSocketDevice] class.
var (
	_VZVirtioSocketDeviceClass     VZVirtioSocketDeviceClass
	_VZVirtioSocketDeviceClassOnce sync.Once
)

func getVZVirtioSocketDeviceClass() VZVirtioSocketDeviceClass {
	_VZVirtioSocketDeviceClassOnce.Do(func() {
		_VZVirtioSocketDeviceClass = VZVirtioSocketDeviceClass{class: objc.GetClass("VZVirtioSocketDevice")}
	})
	return _VZVirtioSocketDeviceClass
}

// GetVZVirtioSocketDeviceClass returns the class object for VZVirtioSocketDevice.
func GetVZVirtioSocketDeviceClass() VZVirtioSocketDeviceClass {
	return getVZVirtioSocketDeviceClass()
}

type VZVirtioSocketDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSocketDeviceClass) Alloc() VZVirtioSocketDevice {
	rv := objc.Send[VZVirtioSocketDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A device that manages port-based connections between the guest system and
// the host computer.
//
// # Overview
// 
// Use a [VZVirtioSocketDevice] object to configure services and other
// communication end points in your virtual machine. Host computers make
// services available using ports, which identify the type of service and the
// protocol to use when transmitting data. Use this object to specify the
// ports available to your guest operating system, and to register handlers to
// manage the communication on those ports.
// 
// Don’t create a [VZVirtioSocketDevice] object directly. Instead, when you
// request a socket device in your configuration, the virtual machine creates
// it and stores it in the [VZVirtioSocketDevice.SocketDevices] property. For each port you want to
// make available in your virtual machine, call the [VZVirtioSocketDevice.SetSocketListenerForPort]
// method and provide an object to manage the port connections.
//
// # Configuring Port Listeners
//
//   - [VZVirtioSocketDevice.SetSocketListenerForPort]: Configures an object to monitor the specified port for new connections.
//   - [VZVirtioSocketDevice.RemoveSocketListenerForPort]: Removes the listener object from the specfied port.
//
// # Connecting to Guest System Ports
//
//   - [VZVirtioSocketDevice.ConnectToPortCompletionHandler]: Initiates a connection to the specified port of the guest operating system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDevice
type VZVirtioSocketDevice struct {
	VZSocketDevice
}

// VZVirtioSocketDeviceFromID constructs a [VZVirtioSocketDevice] from an objc.ID.
//
// A device that manages port-based connections between the guest system and
// the host computer.
func VZVirtioSocketDeviceFromID(id objc.ID) VZVirtioSocketDevice {
	return VZVirtioSocketDevice{VZSocketDevice: VZSocketDeviceFromID(id)}
}
// NOTE: VZVirtioSocketDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZVirtioSocketDevice] class.
//
// # Configuring Port Listeners
//
//   - [IVZVirtioSocketDevice.SetSocketListenerForPort]: Configures an object to monitor the specified port for new connections.
//   - [IVZVirtioSocketDevice.RemoveSocketListenerForPort]: Removes the listener object from the specfied port.
//
// # Connecting to Guest System Ports
//
//   - [IVZVirtioSocketDevice.ConnectToPortCompletionHandler]: Initiates a connection to the specified port of the guest operating system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDevice
type IVZVirtioSocketDevice interface {
	IVZSocketDevice

	// Topic: Configuring Port Listeners

	// Configures an object to monitor the specified port for new connections.
	SetSocketListenerForPort(listener IVZVirtioSocketListener, port uint32)
	// Removes the listener object from the specfied port.
	RemoveSocketListenerForPort(port uint32)

	// Topic: Connecting to Guest System Ports

	// Initiates a connection to the specified port of the guest operating system.
	ConnectToPortCompletionHandler(port uint32, completionHandler VirtioSocketConnectionErrorHandler)
}





// Init initializes the instance.
func (v VZVirtioSocketDevice) Init() VZVirtioSocketDevice {
	rv := objc.Send[VZVirtioSocketDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSocketDevice) Autorelease() VZVirtioSocketDevice {
	rv := objc.Send[VZVirtioSocketDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketDevice creates a new VZVirtioSocketDevice instance.
func NewVZVirtioSocketDevice() VZVirtioSocketDevice {
	class := getVZVirtioSocketDeviceClass()
	rv := objc.Send[VZVirtioSocketDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Configures an object to monitor the specified port for new connections.
//
// listener: The [VZVirtioSocketListener] object to monitor the port. This object
// replaces the previous listener object, if any.
//
// port: The port number to monitor.
//
// # Discussion
// 
// You can register the same listener object on multiple ports. When the guest
// operating system opens a connection to the port, the listener object
// notifies its associated delegate.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDevice/setSocketListener(_:forPort:)
func (v VZVirtioSocketDevice) SetSocketListenerForPort(listener IVZVirtioSocketListener, port uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("setSocketListener:forPort:"), listener, port)
}

// Removes the listener object from the specfied port.
//
// port: The port number to clear. If the specified port doesn’t have a listener
// object, this method does nothing.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDevice/removeSocketListener(forPort:)
func (v VZVirtioSocketDevice) RemoveSocketListenerForPort(port uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("removeSocketListenerForPort:"), port)
}

// Initiates a connection to the specified port of the guest operating system.
//
// port: The destination port number in the guest operating system.
//
// # Discussion
// 
// This method initiates the connection asynchronously, and executes the
// completion handler when the results are available. If the guest operating
// system doesn’t listen for connections to the specifed port, this method
// does nothing.
// 
// For a successful connection, this method sets the [SourcePort] property of
// the resulting [VZVirtioSocketConnection] object to a random port number.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDevice/connect(toPort:)
func (v VZVirtioSocketDevice) ConnectToPortCompletionHandler(port uint32, completionHandler VirtioSocketConnectionErrorHandler) {
		_block1, _cleanup1 := NewVirtioSocketConnectionErrorBlock(completionHandler)
	defer _cleanup1()
		objc.Send[objc.ID](v.ID, objc.Sel("connectToPort:completionHandler:"), port, _block1)
}



























// ConnectToPort is a synchronous wrapper around [VZVirtioSocketDevice.ConnectToPortCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtioSocketDevice) ConnectToPort(ctx context.Context, port uint32) (*VZVirtioSocketConnection, error) {
	type result struct {
		val *VZVirtioSocketConnection
		err error
	}
	done := make(chan result, 1)
	v.ConnectToPortCompletionHandler(port, func(val *VZVirtioSocketConnection, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}







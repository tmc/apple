// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioSocketConnection] class.
var (
	_VZVirtioSocketConnectionClass     VZVirtioSocketConnectionClass
	_VZVirtioSocketConnectionClassOnce sync.Once
)

func getVZVirtioSocketConnectionClass() VZVirtioSocketConnectionClass {
	_VZVirtioSocketConnectionClassOnce.Do(func() {
		_VZVirtioSocketConnectionClass = VZVirtioSocketConnectionClass{class: objc.GetClass("VZVirtioSocketConnection")}
	})
	return _VZVirtioSocketConnectionClass
}

// GetVZVirtioSocketConnectionClass returns the class object for VZVirtioSocketConnection.
func GetVZVirtioSocketConnectionClass() VZVirtioSocketConnectionClass {
	return getVZVirtioSocketConnectionClass()
}

type VZVirtioSocketConnectionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSocketConnectionClass) Alloc() VZVirtioSocketConnection {
	rv := objc.Send[VZVirtioSocketConnection](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A port-based connection between the guest operating system and the host
// computer.
//
// # Overview
// 
// A [VZVirtioSocketConnection] object contains the port information for the
// guest operating system and host computer. You don’t create connection
// objects directly. When the guest operating system initiates a connection,
// the virtual machine creates the connection object and passes it to the
// appropriate [VZVirtioSocketListener] object, which forwards the object to
// its delegate. When the virtual machine opens a connection to a guest port,
// the [ConnectToPortCompletionHandler] method (Objective-C) or
// [connect(toPort:completionHandler:)] method (Swift) pass the connection
// object to your completion handler.
//
// [connect(toPort:completionHandler:)]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDevice/connect(toPort:completionHandler:)
//
// # Getting the connection details
//
//   - [VZVirtioSocketConnection.SourcePort]: The port number of the system that opened the connection.
//   - [VZVirtioSocketConnection.DestinationPort]: The destination port number of the connection.
//   - [VZVirtioSocketConnection.FileDescriptor]: The file descriptor to use when sending data.
//
// # Closing the connection
//
//   - [VZVirtioSocketConnection.Close]: Close the file descriptor associated with the socket.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection
type VZVirtioSocketConnection struct {
	objectivec.Object
}

// VZVirtioSocketConnectionFromID constructs a [VZVirtioSocketConnection] from an objc.ID.
//
// A port-based connection between the guest operating system and the host
// computer.
func VZVirtioSocketConnectionFromID(id objc.ID) VZVirtioSocketConnection {
	return VZVirtioSocketConnection{objectivec.Object{id}}
}
// NOTE: VZVirtioSocketConnection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZVirtioSocketConnection] class.
//
// # Getting the connection details
//
//   - [IVZVirtioSocketConnection.SourcePort]: The port number of the system that opened the connection.
//   - [IVZVirtioSocketConnection.DestinationPort]: The destination port number of the connection.
//   - [IVZVirtioSocketConnection.FileDescriptor]: The file descriptor to use when sending data.
//
// # Closing the connection
//
//   - [IVZVirtioSocketConnection.Close]: Close the file descriptor associated with the socket.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection
type IVZVirtioSocketConnection interface {
	objectivec.IObject

	// Topic: Getting the connection details

	// The port number of the system that opened the connection.
	SourcePort() uint32
	// The destination port number of the connection.
	DestinationPort() uint32
	// The file descriptor to use when sending data.
	FileDescriptor() int

	// Topic: Closing the connection

	// Close the file descriptor associated with the socket.
	Close()
}





// Init initializes the instance.
func (v VZVirtioSocketConnection) Init() VZVirtioSocketConnection {
	rv := objc.Send[VZVirtioSocketConnection](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSocketConnection) Autorelease() VZVirtioSocketConnection {
	rv := objc.Send[VZVirtioSocketConnection](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketConnection creates a new VZVirtioSocketConnection instance.
func NewVZVirtioSocketConnection() VZVirtioSocketConnection {
	class := getVZVirtioSocketConnectionClass()
	rv := objc.Send[VZVirtioSocketConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Close the file descriptor associated with the socket.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/close()
func (v VZVirtioSocketConnection) Close() {
	objc.Send[objc.ID](v.ID, objc.Sel("close"))
}












// The port number of the system that opened the connection.
//
// # Discussion
// 
// When the guest operating system opens a connection, this property contains
// the port number that the guest specified. When you open a connection to the
// guest operating system from your [VZVirtioSocketDevice] object, this
// property contains a randomly generated port number.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/sourcePort
func (v VZVirtioSocketConnection) SourcePort() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("sourcePort"))
	return rv
}



// The destination port number of the connection.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/destinationPort
func (v VZVirtioSocketConnection) DestinationPort() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("destinationPort"))
	return rv
}



// The file descriptor to use when sending data.
//
// # Discussion
// 
// To send data to the other side of the connection, write to the file
// descriptor. To read data from connection, read from the file descriptor. If
// the socket connection is closed, the value of this property is `-1`.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/fileDescriptor
func (v VZVirtioSocketConnection) FileDescriptor() int {
	rv := objc.Send[int](v.ID, objc.Sel("fileDescriptor"))
	return rv
}
























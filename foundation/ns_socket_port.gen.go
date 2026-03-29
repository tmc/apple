// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SocketPort] class.
var (
	_SocketPortClass     SocketPortClass
	_SocketPortClassOnce sync.Once
)

func getSocketPortClass() SocketPortClass {
	_SocketPortClassOnce.Do(func() {
		_SocketPortClass = SocketPortClass{class: objc.GetClass("NSSocketPort")}
	})
	return _SocketPortClass
}

// GetSocketPortClass returns the class object for NSSocketPort.
func GetSocketPortClass() SocketPortClass {
	return getSocketPortClass()
}

type SocketPortClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SocketPortClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SocketPortClass) Alloc() SocketPort {
	rv := objc.Send[SocketPort](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A port that represents a BSD socket.
//
// # Overview
// 
// A [NSSocketPort] object can be used as an endpoint for distributed object
// connections. Companion classes, [NSMachPort] and [NSMessagePort], allow for
// local (on the same machine) communication only. The [NSSocketPort] class
// allows for both local and remote communication, but may be more expensive
// than the others for the local case.
//
// # Creating Instances
//
//   - [SocketPort.InitWithTCPPort]: Initializes the receiver as a local TCP/IP socket of type `SOCK_STREAM`, listening on a specified port number.
//   - [SocketPort.InitWithProtocolFamilySocketTypeProtocolAddress]: Initializes the receiver as a local socket with the provided arguments.
//   - [SocketPort.InitWithProtocolFamilySocketTypeProtocolSocket]: Initializes the receiver with a previously created local socket.
//   - [SocketPort.InitRemoteWithTCPPortHost]: Initializes the receiver as a TCP/IP socket of type `SOCK_STREAM` that can connect to a remote host on a specified port.
//   - [SocketPort.InitRemoteWithProtocolFamilySocketTypeProtocolAddress]: Initializes the receiver as a remote socket with the provided arguments.
//
// # Getting Information
//
//   - [SocketPort.Address]: The receiver’s socket address structure stored inside an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object.
//   - [SocketPort.Protocol]: The protocol that the receiver uses for communication.
//   - [SocketPort.ProtocolFamily]: The protocol family that the receiver uses for communication.
//   - [SocketPort.Socket]: The receiver’s native socket identifier on the platform.
//   - [SocketPort.SocketType]: The receiver’s socket type.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort
type SocketPort struct {
	NSPort
}

// SocketPortFromID constructs a [SocketPort] from an objc.ID.
//
// A port that represents a BSD socket.
func SocketPortFromID(id objc.ID) SocketPort {
	return NSSocketPort{NSPort: NSPortFromID(id)}
}

// NSSocketPortFromID is an alias for [SocketPortFromID] for cross-framework compatibility.
func NSSocketPortFromID(id objc.ID) SocketPort { return SocketPortFromID(id) }
// NOTE: SocketPort adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SocketPort] class.
//
// # Creating Instances
//
//   - [ISocketPort.InitWithTCPPort]: Initializes the receiver as a local TCP/IP socket of type `SOCK_STREAM`, listening on a specified port number.
//   - [ISocketPort.InitWithProtocolFamilySocketTypeProtocolAddress]: Initializes the receiver as a local socket with the provided arguments.
//   - [ISocketPort.InitWithProtocolFamilySocketTypeProtocolSocket]: Initializes the receiver with a previously created local socket.
//   - [ISocketPort.InitRemoteWithTCPPortHost]: Initializes the receiver as a TCP/IP socket of type `SOCK_STREAM` that can connect to a remote host on a specified port.
//   - [ISocketPort.InitRemoteWithProtocolFamilySocketTypeProtocolAddress]: Initializes the receiver as a remote socket with the provided arguments.
//
// # Getting Information
//
//   - [ISocketPort.Address]: The receiver’s socket address structure stored inside an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object.
//   - [ISocketPort.Protocol]: The protocol that the receiver uses for communication.
//   - [ISocketPort.ProtocolFamily]: The protocol family that the receiver uses for communication.
//   - [ISocketPort.Socket]: The receiver’s native socket identifier on the platform.
//   - [ISocketPort.SocketType]: The receiver’s socket type.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort
type ISocketPort interface {
	INSPort

	// Topic: Creating Instances

	// Initializes the receiver as a local TCP/IP socket of type `SOCK_STREAM`, listening on a specified port number.
	InitWithTCPPort(port uint16) SocketPort
	// Initializes the receiver as a local socket with the provided arguments.
	InitWithProtocolFamilySocketTypeProtocolAddress(family int, type_ int, protocol_ int, address INSData) SocketPort
	// Initializes the receiver with a previously created local socket.
	InitWithProtocolFamilySocketTypeProtocolSocket(family int, type_ int, protocol_ int, sock NSSocketNativeHandle) SocketPort
	// Initializes the receiver as a TCP/IP socket of type `SOCK_STREAM` that can connect to a remote host on a specified port.
	InitRemoteWithTCPPortHost(port uint16, hostName string) SocketPort
	// Initializes the receiver as a remote socket with the provided arguments.
	InitRemoteWithProtocolFamilySocketTypeProtocolAddress(family int, type_ int, protocol_ int, address INSData) SocketPort

	// Topic: Getting Information

	// The receiver’s socket address structure stored inside an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object.
	Address() INSData
	// The protocol that the receiver uses for communication.
	Protocol() int
	// The protocol family that the receiver uses for communication.
	ProtocolFamily() int
	// The receiver’s native socket identifier on the platform.
	Socket() NSSocketNativeHandle
	// The receiver’s socket type.
	SocketType() int
}

// Init initializes the instance.
func (s SocketPort) Init() SocketPort {
	rv := objc.Send[SocketPort](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SocketPort) Autorelease() SocketPort {
	rv := objc.Send[SocketPort](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSocketPort creates a new SocketPort instance.
func NewSocketPort() SocketPort {
	class := getSocketPortClass()
	rv := objc.Send[SocketPort](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the receiver as a remote socket with the provided arguments.
//
// family: The protocol family for the socket port. Possible values are defined in ``,
// such as `AF_LOCAL`, `AF_INET`, and `AF_INET6`.
//
// type: The type of socket.
//
// protocol: The specific protocol to use from the protocol family.
//
// address: The family-specific socket address for the receiver copied into an [NSData]
// object.
//
// # Discussion
// 
// A connection is not opened to the remote address until data is sent.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/init(remoteWithProtocolFamily:socketType:protocol:address:)
func NewSocketPortRemoteWithProtocolFamilySocketTypeProtocolAddress(family int, type_ int, protocol_ int, address INSData) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initRemoteWithProtocolFamily:socketType:protocol:address:"), family, type_, protocol_, address)
	return SocketPortFromID(rv)
}

// Initializes the receiver as a TCP/IP socket of type `SOCK_STREAM` that can
// connect to a remote host on a specified port.
//
// port: The port to connect to.
//
// hostName: The host name to connect to. `hostName` may be either a host name or an
// IPv4-style address.
//
// # Return Value
// 
// A TCP/IP socket port of type `SOCK_STREAM` that can connect to the remote
// host `hostName` on port `port`.
//
// # Discussion
// 
// A connection is not opened to the remote host until data is sent.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/init(remoteWithTCPPort:host:)
func NewSocketPortRemoteWithTCPPortHost(port uint16, hostName string) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initRemoteWithTCPPort:host:"), port, objc.String(hostName))
	return SocketPortFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewSocketPortWithCoder(coder INSCoder) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SocketPortFromID(rv)
}

// Initializes the receiver as a local socket with the provided arguments.
//
// family: The protocol family for the socket port. Possible values are defined in ``,
// such as `AF_LOCAL`, `AF_INET`, and `AF_INET6`.
//
// type: The type of socket.
//
// protocol: The specific protocol to use from the protocol family.
//
// address: The family-specific socket address for the receiver copied into an [NSData]
// object.
//
// # Return Value
// 
// A local socket port initialized with the provided arguments.
//
// # Discussion
// 
// The receiver must be added to a run loop before it can accept connections
// or receive messages. Incoming messages are passed to the receiver’s
// delegate method handlePortMessage:.
// 
// To create a standard TCP/IP socket, use [InitWithTCPPort].
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/init(protocolFamily:socketType:protocol:address:)
func NewSocketPortWithProtocolFamilySocketTypeProtocolAddress(family int, type_ int, protocol_ int, address INSData) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProtocolFamily:socketType:protocol:address:"), family, type_, protocol_, address)
	return SocketPortFromID(rv)
}

// Initializes the receiver with a previously created local socket.
//
// family: The protocol family for the provided socket. Possible values are defined in
// ``, such as `AF_LOCAL`, `AF_INET`, and `AF_INET6`.
//
// type: The type of the provided socket.
//
// protocol: The specific protocol the provided socket uses.
//
// sock: The previously created socket.
//
// # Return Value
// 
// A local socket port initialized with the provided socket.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/init(protocolFamily:socketType:protocol:socket:)
func NewSocketPortWithProtocolFamilySocketTypeProtocolSocket(family int, type_ int, protocol_ int, sock NSSocketNativeHandle) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProtocolFamily:socketType:protocol:socket:"), family, type_, protocol_, sock)
	return SocketPortFromID(rv)
}

// Initializes the receiver as a local TCP/IP socket of type `SOCK_STREAM`,
// listening on a specified port number.
//
// port: The port number for the newly created socket port to listen on. If `port`
// is 0, the system will assign a port number.
//
// # Return Value
// 
// An initialized local TCP/IP socket of type `SOCK_STREAM`, listening on port
// `port`.
//
// # Discussion
// 
// This method creates an IPv4 port, not an IPv6 port.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/init(tcpPort:)
func NewSocketPortWithTCPPort(port uint16) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTCPPort:"), port)
	return SocketPortFromID(rv)
}

// Initializes the receiver as a local TCP/IP socket of type `SOCK_STREAM`,
// listening on a specified port number.
//
// port: The port number for the newly created socket port to listen on. If `port`
// is 0, the system will assign a port number.
//
// # Return Value
// 
// An initialized local TCP/IP socket of type `SOCK_STREAM`, listening on port
// `port`.
//
// # Discussion
// 
// This method creates an IPv4 port, not an IPv6 port.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/init(tcpPort:)
func (s SocketPort) InitWithTCPPort(port uint16) SocketPort {
	rv := objc.Send[SocketPort](s.ID, objc.Sel("initWithTCPPort:"), port)
	return rv
}
// Initializes the receiver as a local socket with the provided arguments.
//
// family: The protocol family for the socket port. Possible values are defined in ``,
// such as `AF_LOCAL`, `AF_INET`, and `AF_INET6`.
//
// type: The type of socket.
//
// protocol: The specific protocol to use from the protocol family.
//
// address: The family-specific socket address for the receiver copied into an [NSData]
// object.
//
// # Return Value
// 
// A local socket port initialized with the provided arguments.
//
// # Discussion
// 
// The receiver must be added to a run loop before it can accept connections
// or receive messages. Incoming messages are passed to the receiver’s
// delegate method handlePortMessage:.
// 
// To create a standard TCP/IP socket, use [InitWithTCPPort].
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/init(protocolFamily:socketType:protocol:address:)
func (s SocketPort) InitWithProtocolFamilySocketTypeProtocolAddress(family int, type_ int, protocol_ int, address INSData) SocketPort {
	rv := objc.Send[SocketPort](s.ID, objc.Sel("initWithProtocolFamily:socketType:protocol:address:"), family, type_, protocol_, address)
	return rv
}
// Initializes the receiver with a previously created local socket.
//
// family: The protocol family for the provided socket. Possible values are defined in
// ``, such as `AF_LOCAL`, `AF_INET`, and `AF_INET6`.
//
// type: The type of the provided socket.
//
// protocol: The specific protocol the provided socket uses.
//
// sock: The previously created socket.
//
// # Return Value
// 
// A local socket port initialized with the provided socket.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/init(protocolFamily:socketType:protocol:socket:)
func (s SocketPort) InitWithProtocolFamilySocketTypeProtocolSocket(family int, type_ int, protocol_ int, sock NSSocketNativeHandle) SocketPort {
	rv := objc.Send[SocketPort](s.ID, objc.Sel("initWithProtocolFamily:socketType:protocol:socket:"), family, type_, protocol_, sock)
	return rv
}
// Initializes the receiver as a TCP/IP socket of type `SOCK_STREAM` that can
// connect to a remote host on a specified port.
//
// port: The port to connect to.
//
// hostName: The host name to connect to. `hostName` may be either a host name or an
// IPv4-style address.
//
// # Return Value
// 
// A TCP/IP socket port of type `SOCK_STREAM` that can connect to the remote
// host `hostName` on port `port`.
//
// # Discussion
// 
// A connection is not opened to the remote host until data is sent.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/init(remoteWithTCPPort:host:)
func (s SocketPort) InitRemoteWithTCPPortHost(port uint16, hostName string) SocketPort {
	rv := objc.Send[SocketPort](s.ID, objc.Sel("initRemoteWithTCPPort:host:"), port, objc.String(hostName))
	return rv
}
// Initializes the receiver as a remote socket with the provided arguments.
//
// family: The protocol family for the socket port. Possible values are defined in ``,
// such as `AF_LOCAL`, `AF_INET`, and `AF_INET6`.
//
// type: The type of socket.
//
// protocol: The specific protocol to use from the protocol family.
//
// address: The family-specific socket address for the receiver copied into an [NSData]
// object.
//
// # Discussion
// 
// A connection is not opened to the remote address until data is sent.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/init(remoteWithProtocolFamily:socketType:protocol:address:)
func (s SocketPort) InitRemoteWithProtocolFamilySocketTypeProtocolAddress(family int, type_ int, protocol_ int, address INSData) SocketPort {
	rv := objc.Send[SocketPort](s.ID, objc.Sel("initRemoteWithProtocolFamily:socketType:protocol:address:"), family, type_, protocol_, address)
	return rv
}

// The receiver’s socket address structure stored inside an [NSData] object.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/address
func (s SocketPort) Address() INSData {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("address"))
	return NSDataFromID(objc.ID(rv))
}
// The protocol that the receiver uses for communication.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/protocol
func (s SocketPort) Protocol() int {
	rv := objc.Send[int](s.ID, objc.Sel("protocol"))
	return rv
}
// The protocol family that the receiver uses for communication.
//
// # Discussion
// 
// Possible values are defined in ``, such as `AF_LOCAL`, `AF_INET`, and
// `AF_INET6`.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/protocolFamily
func (s SocketPort) ProtocolFamily() int {
	rv := objc.Send[int](s.ID, objc.Sel("protocolFamily"))
	return rv
}
// The receiver’s native socket identifier on the platform.
//
// # Discussion
// 
// In macOS, the native socket identifier is an integer file descriptor.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/socket
func (s SocketPort) Socket() NSSocketNativeHandle {
	rv := objc.Send[NSSocketNativeHandle](s.ID, objc.Sel("socket"))
	return NSSocketNativeHandle(rv)
}
// The receiver’s socket type.
//
// See: https://developer.apple.com/documentation/Foundation/SocketPort/socketType
func (s SocketPort) SocketType() int {
	rv := objc.Send[int](s.ID, objc.Sel("socketType"))
	return rv
}


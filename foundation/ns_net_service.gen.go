// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NetService] class.
var (
	_NetServiceClass     NetServiceClass
	_NetServiceClassOnce sync.Once
)

func getNetServiceClass() NetServiceClass {
	_NetServiceClassOnce.Do(func() {
		_NetServiceClass = NetServiceClass{class: objc.GetClass("NSNetService")}
	})
	return _NetServiceClass
}

// GetNetServiceClass returns the class object for NSNetService.
func GetNetServiceClass() NetServiceClass {
	return getNetServiceClass()
}

type NetServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NetServiceClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NetServiceClass) Alloc() NetService {
	rv := objc.Send[NetService](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A network service that broadcasts its availability using multicast DNS.
//
// # Overview
// 
// The [NSNetService] class represents a network service, either one your
// application publishes or is a client of. This class and the
// [NSNetServiceBrowser] class use multicast DNS to convey information about
// network services to and from your application. The API of [NSNetService]
// provides a convenient way to publish the services offered by your
// application and to resolve the socket address for a service.
// 
// The types of services you access using [NSNetService] are the same types
// that you access directly using BSD sockets. HTTP and FTP are two services
// commonly provided by systems. (For a list of common services and the ports
// used by those services, see the file `/etc/services`.) Applications can
// also define their own custom services to provide specific data to clients.
// 
// You can use the [NSNetService] class as either a publisher of a service or
// a client of a service. If your application publishes a service, your code
// must acquire a port and prepare a socket to communicate with clients. Once
// your socket is ready, you use the [NSNetService] class to notify clients
// that your service is ready. If your application is the client of a network
// service, you can either create an [NSNetService] object directly (if you
// know the exact host and port information) or use an [NSNetServiceBrowser]
// object to browse for services.
// 
// To publish a service, initialize your [NSNetService] object with the
// service name, domain, type, and port information. All of this information
// must be valid for the socket created by your application. Once initialized,
// call the [Publish] method to broadcast your service information to the
// network.
// 
// When connecting to a service, use the [NSNetServiceBrowser] class to locate
// the service on the network and obtain the corresponding [NSNetService]
// object. Once you have the object, call the [ResolveWithTimeout] method to
// verify that the service is available and ready for your application. If it
// is, the [Addresses] property provides the socket information you can use to
// connect to the service.
// 
// The methods of [NSNetService] operate asynchronously so your application is
// not impacted by the speed of the network. All information about a service
// is returned to your application through the [NSNetService] object’s
// delegate. You must provide a delegate object to respond to messages and to
// handle errors appropriately.
//
// # Configuring Network Services
//
//   - [NetService.Addresses]: A read-only array containing [NSData] objects, each of which contains a socket address for the service.
//   - [NetService.Domain]: A string containing the domain for this service.
//   - [NetService.IncludesPeerToPeer]: Specifies whether to also publish, resolve, or monitor this service over peer-to-peer Bluetooth and Wi-Fi, if available.
//   - [NetService.SetIncludesPeerToPeer]
//   - [NetService.Name]: A string containing the name of this service.
//   - [NetService.Type]: The type of the published service.
//   - [NetService.Delegate]: The delegate for the receiver.
//   - [NetService.SetDelegate]
//
// # Using Network Services
//
//   - [NetService.Port]: The port on which the service is listening for connections.
//
// # Obtaining the DNS Hostname
//
//   - [NetService.HostName]: A string containing the DNS hostname for this service.
//
// See: https://developer.apple.com/documentation/Foundation/NetService
type NetService struct {
	objectivec.Object
}

// NetServiceFromID constructs a [NetService] from an objc.ID.
//
// A network service that broadcasts its availability using multicast DNS.
func NetServiceFromID(id objc.ID) NetService {
	return NSNetService{objectivec.Object{ID: id}}
}

// NSNetServiceFromID is an alias for [NetServiceFromID] for cross-framework compatibility.
func NSNetServiceFromID(id objc.ID) NetService { return NetServiceFromID(id) }
// NOTE: NetService adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NetService] class.
//
// # Configuring Network Services
//
//   - [INetService.Addresses]: A read-only array containing [NSData] objects, each of which contains a socket address for the service.
//   - [INetService.Domain]: A string containing the domain for this service.
//   - [INetService.IncludesPeerToPeer]: Specifies whether to also publish, resolve, or monitor this service over peer-to-peer Bluetooth and Wi-Fi, if available.
//   - [INetService.SetIncludesPeerToPeer]
//   - [INetService.Name]: A string containing the name of this service.
//   - [INetService.Type]: The type of the published service.
//   - [INetService.Delegate]: The delegate for the receiver.
//   - [INetService.SetDelegate]
//
// # Using Network Services
//
//   - [INetService.Port]: The port on which the service is listening for connections.
//
// # Obtaining the DNS Hostname
//
//   - [INetService.HostName]: A string containing the DNS hostname for this service.
//
// See: https://developer.apple.com/documentation/Foundation/NetService
type INetService interface {
	objectivec.IObject

	// Topic: Configuring Network Services

	// A read-only array containing [NSData] objects, each of which contains a socket address for the service.
	Addresses() []NSData
	// A string containing the domain for this service.
	Domain() string
	// Specifies whether to also publish, resolve, or monitor this service over peer-to-peer Bluetooth and Wi-Fi, if available.
	IncludesPeerToPeer() bool
	SetIncludesPeerToPeer(value bool)
	// A string containing the name of this service.
	Name() string
	// The type of the published service.
	Type() string
	// The delegate for the receiver.
	Delegate() NSNetServiceDelegate
	SetDelegate(value NSNetServiceDelegate)

	// Topic: Using Network Services

	// The port on which the service is listening for connections.
	Port() int

	// Topic: Obtaining the DNS Hostname

	// A string containing the DNS hostname for this service.
	HostName() string
}

// Init initializes the instance.
func (n NetService) Init() NetService {
	rv := objc.Send[NetService](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NetService) Autorelease() NetService {
	rv := objc.Send[NetService](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNetService creates a new NetService instance.
func NewNetService() NetService {
	class := getNetServiceClass()
	rv := objc.Send[NetService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the receiver, initialized as a network service of a given type and
// sets the initial host information.
//
// domain: The domain for the service. To resolve in the default domains, pass in an
// empty string (`@""`). To limit resolution to the local domain, use
// `@"local."`.
// 
// If you are creating this object to resolve a service whose information your
// app stored previously, you should set this to the domain in which the
// service was originally discovered.
// 
// You can also use a [NSNetServiceBrowser] object to obtain a list of
// possible domains in which you can discover and resolve services.
//
// type: The network service type.
// 
// `type` must contain both the service type and transport layer information.
// To ensure that the mDNS responder searches for services, as opposed to
// hosts, prefix both the service name and transport layer name with an
// underscore character (”_”). For example, to search for an HTTP service
// on TCP, you would use the type string “`_http._tcp.`”. Note that the
// period character at the end of the string, which indicates that the domain
// name is an absolute name, is required.
//
// name: The name of the service to resolve.
//
// # Return Value
// 
// The receiver, initialized as a network service named `name` of type `type`
// in the domain `domain`.
//
// # Discussion
// 
// This method is the appropriate initializer to use to resolve a service—to
// publish a service, use [InitWithDomainTypeNamePort].
// 
// If you know the values for `domain`, `type`, and `name` of the service you
// wish to connect to, you can create an [NSNetService] object using this
// initializer and call [ResolveWithTimeout] on the result.
// 
// You cannot use this initializer to publish a service. This initializer
// passes an invalid port number to the designated initializer, which prevents
// the service from being registered. Calling [Publish] on an [NSNetService]
// object initialized with this method generates a call to your delegate’s
// [NetServiceDidNotPublish] method with an [NetServicesBadArgumentError]
// error.
//
// See: https://developer.apple.com/documentation/Foundation/NetService/init(domain:type:name:)
func NewNetServiceWithDomainTypeName(domain string, type_ string, name string) NetService {
	instance := getNetServiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDomain:type:name:"), objc.String(domain), objc.String(type_), objc.String(name))
	return NetServiceFromID(rv)
}

// Initializes the receiver for publishing a network service of type `type` at
// the socket location specified by `domain`, `name`, and `port`.
//
// domain: The domain for the service. To use the default registration domains, pass
// in an empty string (`@""`). To limit registration to the local domain, use
// `@"local."`.
// 
// You can also use a [NSNetServiceBrowser] object to obtain a list of
// possible domains in which you can publish your service.
//
// type: The network service type.
// 
// `type` must contain both the service type and transport layer information.
// To ensure that the mDNS responder searches for services, as opposed to
// hosts, prefix both the service name and transport layer name with an
// underscore character (”_”). For example, to search for an HTTP service
// on TCP, you would use the type string “`_http._tcp.`”. Note that the
// period character at the end of the string, which indicates that the domain
// name is an absolute name, is required.
//
// name: The name by which the service is identified to the network. The name must
// be unique. If you pass the empty string (`@""`), the system automatically
// advertises your service using the computer name as the service name.
//
// port: The port on which the service is published.
// 
// If you specify the [NSNetServiceListenForConnections] flag, you may pass
// zero (`0`), in which case the service automatically allocates an arbitrary
// (ephemeral) port for your service. When the delegate’s
// [NetServiceDidPublish] is called, you can determine the actual port chosen
// by calling the service object’s [NSNetService] method or accessing the
// corresponding property.
// 
// If your app is listening for connections on its own, the value of `port`
// must be a port number acquired by your application for the service.
//
// # Discussion
// 
// You use this method to create a service that you wish to publish on the
// network. Although you can also use this method to create a service you wish
// to resolve on the network, it is generally more appropriate to use the
// [InitWithDomainTypeName] method instead.
// 
// When publishing a service, you must provide valid arguments in order to
// advertise your service correctly. If the host computer has access to
// multiple registration domains, you must create separate [NSNetService]
// objects for each domain. If you attempt to publish in a domain for which
// you do not have registration authority, your request may be denied.
// 
// It is acceptable to use an empty string for the `domain` argument when
// publishing or browsing a service, but do not rely on this for resolution.
// 
// This method is the designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/NetService/init(domain:type:name:port:)
func NewNetServiceWithDomainTypeNamePort(domain string, type_ string, name string, port int) NetService {
	instance := getNetServiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDomain:type:name:port:"), objc.String(domain), objc.String(type_), objc.String(name), port)
	return NetServiceFromID(rv)
}

// A read-only array containing [NSData] objects, each of which contains a
// socket address for the service.
//
// # Discussion
// 
// An array containing [NSData] objects, each of which contains a socket
// address for the service. Each [NSData] object in the returned array
// contains an appropriate `sockaddr` structure that you can use to connect to
// the socket. The exact type of this structure depends on the service to
// which you are connecting. If no addresses were resolved for the service,
// the returned array contains zero elements.
// 
// It is possible for a single service to resolve to more than one address or
// not resolve to any addresses. A service might resolve to multiple addresses
// if the computer publishing the service is currently multihoming.
//
// See: https://developer.apple.com/documentation/Foundation/NetService/addresses
func (n NetService) Addresses() []NSData {
	rv := objc.Send[[]objc.ID](n.ID, objc.Sel("addresses"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSData {
		return NSDataFromID(id)
	})
}
// A string containing the domain for this service.
//
// # Discussion
// 
// This can be an explicit domain name or it can contain the generic local
// domain name, `@"local."` (note the trailing period, which indicates an
// absolute name).
// 
// This property’s value is set when the object is first initialized,
// whether by your code or by a browser object. See [InitWithDomainTypeName]
// for more information.
//
// See: https://developer.apple.com/documentation/Foundation/NetService/domain
func (n NetService) Domain() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("domain"))
	return NSStringFromID(rv).String()
}
// Specifies whether to also publish, resolve, or monitor this service over
// peer-to-peer Bluetooth and Wi-Fi, if available.
//
// # Discussion
// 
// This property must be set before calling [Publish] or [PublishWithOptions],
// [ResolveWithTimeout]`, or [StartMonitoring] in order to take effect.
//
// See: https://developer.apple.com/documentation/Foundation/NetService/includesPeerToPeer
func (n NetService) IncludesPeerToPeer() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("includesPeerToPeer"))
	return rv
}
func (n NetService) SetIncludesPeerToPeer(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setIncludesPeerToPeer:"), value)
}
// A string containing the name of this service.
//
// # Discussion
// 
// This value is set when the object is first initialized, whether by your
// code or by a browser object. See [InitWithDomainTypeName] for more
// information.
//
// See: https://developer.apple.com/documentation/Foundation/NetService/name
func (n NetService) Name() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
// The type of the published service.
//
// # Discussion
// 
// This value is set when the object is first initialized, whether by your
// code or by a browser object. See [InitWithDomainTypeName] for more
// information.
//
// See: https://developer.apple.com/documentation/Foundation/NetService/type
func (n NetService) Type() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("type"))
	return NSStringFromID(rv).String()
}
// The delegate for the receiver.
//
// # Discussion
// 
// The delegate must conform to the [NSNetServiceDelegate] protocol, and is
// not retained.
//
// See: https://developer.apple.com/documentation/Foundation/NetService/delegate
func (n NetService) Delegate() NSNetServiceDelegate {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("delegate"))
	return NSNetServiceDelegateObjectFromID(rv)
}
func (n NetService) SetDelegate(value NSNetServiceDelegate) {
	objc.Send[struct{}](n.ID, objc.Sel("setDelegate:"), value)
}
// The port on which the service is listening for connections.
//
// # Discussion
// 
// If the object was initialized by calling [InitWithDomainTypeNamePort]
// (whether by your code or by a browser object), then the value was set when
// the object was first initialized.
// 
// If the object was initialized by calling [InitWithDomainTypeName], the
// value of this property is not valid (`-1`) until after the service has
// successfully been resolved (when `addresses` is non-`nil`).
//
// See: https://developer.apple.com/documentation/Foundation/NetService/port
func (n NetService) Port() int {
	rv := objc.Send[int](n.ID, objc.Sel("port"))
	return rv
}
// A string containing the DNS hostname for this service.
//
// # Discussion
// 
// This value is `nil` until the service has been resolved (when `addresses`
// is non-`nil`).
//
// See: https://developer.apple.com/documentation/Foundation/NetService/hostName
func (n NetService) HostName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("hostName"))
	return NSStringFromID(rv).String()
}


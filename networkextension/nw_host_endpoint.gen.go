// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NWHostEndpoint] class.
var (
	_NWHostEndpointClass     NWHostEndpointClass
	_NWHostEndpointClassOnce sync.Once
)

func getNWHostEndpointClass() NWHostEndpointClass {
	_NWHostEndpointClassOnce.Do(func() {
		_NWHostEndpointClass = NWHostEndpointClass{class: objc.GetClass("NWHostEndpoint")}
	})
	return _NWHostEndpointClass
}

// GetNWHostEndpointClass returns the class object for NWHostEndpoint.
func GetNWHostEndpointClass() NWHostEndpointClass {
	return getNWHostEndpointClass()
}

type NWHostEndpointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWHostEndpointClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWHostEndpointClass) Alloc() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A network endpoint specified by DNS name (or IP address) and port.
//
// # Getting endpoint properties
//
//   - [NWHostEndpoint.Hostname]: The endpoint’s hostname.
//   - [NWHostEndpoint.Port]: The endpoint’s port, represented as a string.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint
type NWHostEndpoint struct {
	NWEndpoint
}

// NWHostEndpointFromID constructs a [NWHostEndpoint] from an objc.ID.
//
// A network endpoint specified by DNS name (or IP address) and port.
func NWHostEndpointFromID(id objc.ID) NWHostEndpoint {
	return NWHostEndpoint{NWEndpoint: NWEndpointFromID(id)}
}

// NOTE: NWHostEndpoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NWHostEndpoint] class.
//
// # Getting endpoint properties
//
//   - [INWHostEndpoint.Hostname]: The endpoint’s hostname.
//   - [INWHostEndpoint.Port]: The endpoint’s port, represented as a string.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint
type INWHostEndpoint interface {
	INWEndpoint

	// Topic: Getting endpoint properties

	// The endpoint’s hostname.
	Hostname() string
	// The endpoint’s port, represented as a string.
	Port() string
}

// Init initializes the instance.
func (n NWHostEndpoint) Init() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWHostEndpoint) Autorelease() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWHostEndpoint creates a new NWHostEndpoint instance.
func NewNWHostEndpoint() NWHostEndpoint {
	class := getNWHostEndpointClass()
	rv := objc.Send[NWHostEndpoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Create a host endpoint with a hostname and port.
//
// hostname: A string representation of the hostname or address, such as
// `www.ExampleXCUIElementTypeCom()` or `10.0.0.1`.
//
// port: A string containing the port on the host, such as `80`.
//
// # Discussion
//
// If the hostname is a domain name, such as
// `www.ExampleXCUIElementTypeCom()`, starting a connection to the host
// endpoint causes the hostname to be resolved to an address during the
// connection process. If the hostname is an IPv4 or IPv6 address, such as
// `10.0.0.1` or `:1`, starting a connection to the host endpoint will cause
// the address to be used directly.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/init(hostname:port:)
func NewNWHostEndpointWithHostnamePort(hostname string, port string) NWHostEndpoint {
	rv := objc.Send[objc.ID](objc.ID(getNWHostEndpointClass().class), objc.Sel("endpointWithHostname:port:"), objc.String(hostname), objc.String(port))
	return NWHostEndpointFromID(rv)
}

// The endpoint’s hostname.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/hostname
func (n NWHostEndpoint) Hostname() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("hostname"))
	return foundation.NSStringFromID(rv).String()
}

// The endpoint’s port, represented as a string.
//
// # Discussion
//
// Since the port is represented as a string, it is always represented in host
// byte order. If converting between byte fields and strings, make sure to use
// host byte ordering.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWHostEndpoint/port
func (n NWHostEndpoint) Port() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("port"))
	return foundation.NSStringFromID(rv).String()
}

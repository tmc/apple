// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDINetworkHost] class.
var (
	_MIDINetworkHostClass     MIDINetworkHostClass
	_MIDINetworkHostClassOnce sync.Once
)

func getMIDINetworkHostClass() MIDINetworkHostClass {
	_MIDINetworkHostClassOnce.Do(func() {
		_MIDINetworkHostClass = MIDINetworkHostClass{class: objc.GetClass("MIDINetworkHost")}
	})
	return _MIDINetworkHostClass
}

// GetMIDINetworkHostClass returns the class object for MIDINetworkHost.
func GetMIDINetworkHostClass() MIDINetworkHostClass {
	return getMIDINetworkHostClass()
}

type MIDINetworkHostClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDINetworkHostClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDINetworkHostClass) Alloc() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the host’s network address.
//
// # Inspecting Host Properties
//
//   - [MIDINetworkHost.Name]: The host name.
//   - [MIDINetworkHost.NetServiceName]: The net service name.
//   - [MIDINetworkHost.NetServiceDomain]: The net service domain.
//   - [MIDINetworkHost.Address]: The host address.
//   - [MIDINetworkHost.Port]: The host port.
//
// # Comparing Hosts
//
//   - [MIDINetworkHost.HasSameAddressAs]: Compares this host instance with another to see if they share the same address value.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost
type MIDINetworkHost struct {
	objectivec.Object
}

// MIDINetworkHostFromID constructs a [MIDINetworkHost] from an objc.ID.
//
// An object that represents the host’s network address.
func MIDINetworkHostFromID(id objc.ID) MIDINetworkHost {
	return MIDINetworkHost{objectivec.Object{ID: id}}
}

// NOTE: MIDINetworkHost adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDINetworkHost] class.
//
// # Inspecting Host Properties
//
//   - [IMIDINetworkHost.Name]: The host name.
//   - [IMIDINetworkHost.NetServiceName]: The net service name.
//   - [IMIDINetworkHost.NetServiceDomain]: The net service domain.
//   - [IMIDINetworkHost.Address]: The host address.
//   - [IMIDINetworkHost.Port]: The host port.
//
// # Comparing Hosts
//
//   - [IMIDINetworkHost.HasSameAddressAs]: Compares this host instance with another to see if they share the same address value.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost
type IMIDINetworkHost interface {
	objectivec.IObject

	// Topic: Inspecting Host Properties

	// The host name.
	Name() string
	// The net service name.
	NetServiceName() string
	// The net service domain.
	NetServiceDomain() string
	// The host address.
	Address() string
	// The host port.
	Port() uint

	// Topic: Comparing Hosts

	// Compares this host instance with another to see if they share the same address value.
	HasSameAddressAs(other IMIDINetworkHost) bool
}

// Init initializes the instance.
func (m MIDINetworkHost) Init() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDINetworkHost) Autorelease() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDINetworkHost creates a new MIDINetworkHost instance.
func NewMIDINetworkHost() MIDINetworkHost {
	class := getMIDINetworkHostClass()
	rv := objc.Send[MIDINetworkHost](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a host with the specified name, adress, and port.
//
// name: The host name.
//
// address: The host’s IP address or hostname.
//
// port: The host’s UDP port.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:address:port:)
func NewMIDINetworkHostWithNameAddressPort(name string, address string, port uint) MIDINetworkHost {
	rv := objc.Send[objc.ID](objc.ID(getMIDINetworkHostClass().class), objc.Sel("hostWithName:address:port:"), objc.String(name), objc.String(address), port)
	return MIDINetworkHostFromID(rv)
}

// Creates a host with the specified name and net service.
//
// name: The host name.
//
// netService: The net service.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:netService:)
func NewMIDINetworkHostWithNameNetService(name string, netService foundation.NSNetService) MIDINetworkHost {
	rv := objc.Send[objc.ID](objc.ID(getMIDINetworkHostClass().class), objc.Sel("hostWithName:netService:"), objc.String(name), netService)
	return MIDINetworkHostFromID(rv)
}

// Creates a host with the specified name, net service name, and domain.
//
// name: The host name.
//
// netServiceName: The net service name.
//
// netServiceDomain: The net service domain.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:netServiceName:netServiceDomain:)
func NewMIDINetworkHostWithNameNetServiceNameNetServiceDomain(name string, netServiceName string, netServiceDomain string) MIDINetworkHost {
	rv := objc.Send[objc.ID](objc.ID(getMIDINetworkHostClass().class), objc.Sel("hostWithName:netServiceName:netServiceDomain:"), objc.String(name), objc.String(netServiceName), objc.String(netServiceDomain))
	return MIDINetworkHostFromID(rv)
}

// Compares this host instance with another to see if they share the same
// address value.
//
// other: The other host instance to compare.
//
// # Return Value
//
// A Boolean value that indicates whether the hosts have the same address
// value.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/hasSameAddress(as:)
func (m MIDINetworkHost) HasSameAddressAs(other IMIDINetworkHost) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasSameAddressAs:"), other)
	return rv
}

// The host name.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/name
func (m MIDINetworkHost) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The net service name.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/netServiceName
func (m MIDINetworkHost) NetServiceName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("netServiceName"))
	return foundation.NSStringFromID(rv).String()
}

// The net service domain.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/netServiceDomain
func (m MIDINetworkHost) NetServiceDomain() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("netServiceDomain"))
	return foundation.NSStringFromID(rv).String()
}

// The host address.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/address
func (m MIDINetworkHost) Address() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("address"))
	return foundation.NSStringFromID(rv).String()
}

// The host port.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/port
func (m MIDINetworkHost) Port() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("port"))
	return rv
}

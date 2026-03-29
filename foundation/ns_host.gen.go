// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Host] class.
var (
	_HostClass     HostClass
	_HostClassOnce sync.Once
)

func getHostClass() HostClass {
	_HostClassOnce.Do(func() {
		_HostClass = HostClass{class: objc.GetClass("NSHost")}
	})
	return _HostClass
}

// GetHostClass returns the class object for NSHost.
func GetHostClass() HostClass {
	return getHostClass()
}

type HostClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (hc HostClass) Class() objc.Class {
	return hc.class
}

// Alloc allocates memory for a new instance of the class.
func (hc HostClass) Alloc() Host {
	rv := objc.Send[Host](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// A representation of an individual host on the network.
//
// # Overview
// 
// The [NSHost] class provides methods to access the network name and address
// information for a host. Instances of the [NSHost] class represent
// individual on a network. Use [NSHost] objects to get the current host’s
// names and addresses and to look up other hosts by name or by address.
// 
// To create an [NSHost] object, use the [CurrentHost], [HostWithAddress], or
// [HostWithName] class methods (don’t use `alloc` and `init`). These
// methods use available network administration services to discover all names
// and addresses for the host requested. They don’t attempt to contact the
// host itself, however. This approach avoids untimely delays due to a host
// being unavailable, but it may result in incomplete information about the
// host.
// 
// An [NSHost] object contains all of the network addresses and names
// discovered for a given host by the network administration services. Each
// [NSHost] object may contain several addresses and have more than one name.
// If an [NSHost] object has more than one name, the additional names are
// variations on the same name, typically the basic host name plus the fully
// qualified domain name. For example, with a host name `"sales"` in the
// domain `"anycorp.Com()"`, an [NSHost] object can hold both the names
// `"sales"` and `"sales.AnycorpXCUIElementTypeCom()"`.
// 
// [NSHost] methods are thread-safe.
//
// # Getting Host Information
//
//   - [Host.Address]: Returns one of the network addresses of the receiver.
//   - [Host.Addresses]: Returns all the network addresses of the receiver.
//   - [Host.Name]: Returns one of the hostnames of the receiver.
//   - [Host.LocalizedName]: Returns the name used as by default when publishing [NSNetServices].
//   - [Host.Names]: Returns all the hostnames of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Host
type Host struct {
	objectivec.Object
}

// HostFromID constructs a [Host] from an objc.ID.
//
// A representation of an individual host on the network.
func HostFromID(id objc.ID) Host {
	return NSHost{objectivec.Object{ID: id}}
}

// NSHostFromID is an alias for [HostFromID] for cross-framework compatibility.
func NSHostFromID(id objc.ID) Host { return HostFromID(id) }
// NOTE: Host adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [Host] class.
//
// # Getting Host Information
//
//   - [IHost.Address]: Returns one of the network addresses of the receiver.
//   - [IHost.Addresses]: Returns all the network addresses of the receiver.
//   - [IHost.Name]: Returns one of the hostnames of the receiver.
//   - [IHost.LocalizedName]: Returns the name used as by default when publishing [NSNetServices].
//   - [IHost.Names]: Returns all the hostnames of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Host
type IHost interface {
	objectivec.IObject

	// Topic: Getting Host Information

	// Returns one of the network addresses of the receiver.
	Address() string
	// Returns all the network addresses of the receiver.
	Addresses() []string
	// Returns one of the hostnames of the receiver.
	Name() string
	// Returns the name used as by default when publishing [NSNetServices].
	LocalizedName() string
	// Returns all the hostnames of the receiver.
	Names() []string
}

// Init initializes the instance.
func (h Host) Init() Host {
	rv := objc.Send[Host](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h Host) Autorelease() Host {
	rv := objc.Send[Host](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHost creates a new Host instance.
func NewHost() Host {
	class := getHostClass()
	rv := objc.Send[Host](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the [NSHost] with the Internet address `address`.
//
// address: Network address to look up. For example, `"127.0.0.1"` or `":1"`.
//
// # Return Value
// 
// The host for `address`.
//
// See: https://developer.apple.com/documentation/Foundation/Host/init(address:)
func NewHostWithAddress(address string) Host {
	rv := objc.Send[objc.ID](objc.ID(getHostClass().class), objc.Sel("hostWithAddress:"), objc.String(address))
	return HostFromID(rv)
}

// Returns a host with a specific name.
//
// name: Name of the host to look up. Can be either a simple hostname, such as
// `"sales"`, or a fully qualified domain name, such as
// `"sales.AnycorpXCUIElementTypeCom()"`.
//
// # Return Value
// 
// The host named `hostname`.
//
// See: https://developer.apple.com/documentation/Foundation/Host/init(name:)
func NewHostWithName(name string) Host {
	rv := objc.Send[objc.ID](objc.ID(getHostClass().class), objc.Sel("hostWithName:"), objc.String(name))
	return HostFromID(rv)
}

// Returns one of the network addresses of the receiver.
//
// # Return Value
// 
// One of the network address for the receiver. For example, `"192.42.172.1"`
// or `":1"`.
//
// See: https://developer.apple.com/documentation/Foundation/Host/address
func (h Host) Address() string {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("address"))
	return NSStringFromID(rv).String()
}
// Returns all the network addresses of the receiver.
//
// # Return Value
// 
// All the network addresses of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Host/addresses
func (h Host) Addresses() []string {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("addresses"))
	return objc.ConvertSliceToStrings(rv)
}
// Returns one of the hostnames of the receiver.
//
// # Return Value
// 
// One of the hostnames of the receiver. Can be either a simple hostname, such
// as `"sales"`, or a fully qualified domain name, such as
// `"sales.AnycorpXCUIElementTypeCom()"`.
//
// See: https://developer.apple.com/documentation/Foundation/Host/name
func (h Host) Name() string {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
// Returns the name used as by default when publishing [NSNetServices].
//
// # Return Value
// 
// A string containing the computer name.
// 
// # Discussion
// 
// This is the name displayed in the Finder sidebar, as well as in the Sharing
// preference panel.
// 
// This method only returns an [NSString] when sent to the [CurrentHost]
// instance, all other instances currently return `nil`.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/Foundation/Host/localizedName
func (h Host) LocalizedName() string {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("localizedName"))
	return NSStringFromID(rv).String()
}
// Returns all the hostnames of the receiver.
//
// # Return Value
// 
// All the hostnames of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Host/names
func (h Host) Names() []string {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("names"))
	return objc.ConvertSliceToStrings(rv)
}


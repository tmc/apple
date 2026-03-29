// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEIPv6Route] class.
var (
	_NEIPv6RouteClass     NEIPv6RouteClass
	_NEIPv6RouteClassOnce sync.Once
)

func getNEIPv6RouteClass() NEIPv6RouteClass {
	_NEIPv6RouteClassOnce.Do(func() {
		_NEIPv6RouteClass = NEIPv6RouteClass{class: objc.GetClass("NEIPv6Route")}
	})
	return _NEIPv6RouteClass
}

// GetNEIPv6RouteClass returns the class object for NEIPv6Route.
func GetNEIPv6RouteClass() NEIPv6RouteClass {
	return getNEIPv6RouteClass()
}

type NEIPv6RouteClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEIPv6RouteClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEIPv6RouteClass) Alloc() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The settings for an IPv6 route.
//
// # Creating an IPv6 Route
//
//   - [NEIPv6Route.InitWithDestinationAddressNetworkPrefixLength]: Initialize the NEIPv6Route
//
// # Accessing IPv6 Route Properties
//
//   - [NEIPv6Route.DestinationAddress]: The destination network address of the route.
//   - [NEIPv6Route.DestinationNetworkPrefixLength]: The destination network prefix length of the route.
//   - [NEIPv6Route.GatewayAddress]: The address of the next-hop gateway of the route.
//   - [NEIPv6Route.SetGatewayAddress]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route
type NEIPv6Route struct {
	objectivec.Object
}

// NEIPv6RouteFromID constructs a [NEIPv6Route] from an objc.ID.
//
// The settings for an IPv6 route.
func NEIPv6RouteFromID(id objc.ID) NEIPv6Route {
	return NEIPv6Route{objectivec.Object{ID: id}}
}
// NOTE: NEIPv6Route adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEIPv6Route] class.
//
// # Creating an IPv6 Route
//
//   - [INEIPv6Route.InitWithDestinationAddressNetworkPrefixLength]: Initialize the NEIPv6Route
//
// # Accessing IPv6 Route Properties
//
//   - [INEIPv6Route.DestinationAddress]: The destination network address of the route.
//   - [INEIPv6Route.DestinationNetworkPrefixLength]: The destination network prefix length of the route.
//   - [INEIPv6Route.GatewayAddress]: The address of the next-hop gateway of the route.
//   - [INEIPv6Route.SetGatewayAddress]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route
type INEIPv6Route interface {
	objectivec.IObject

	// Topic: Creating an IPv6 Route

	// Initialize the NEIPv6Route
	InitWithDestinationAddressNetworkPrefixLength(address string, networkPrefixLength foundation.NSNumber) NEIPv6Route

	// Topic: Accessing IPv6 Route Properties

	// The destination network address of the route.
	DestinationAddress() string
	// The destination network prefix length of the route.
	DestinationNetworkPrefixLength() foundation.NSNumber
	// The address of the next-hop gateway of the route.
	GatewayAddress() string
	SetGatewayAddress(value string)

	// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
	ExcludedRoutes() INEIPv6Route
	SetExcludedRoutes(value INEIPv6Route)
	// The IPv6 network traffic that the system routes to the TUN interface.
	IncludedRoutes() INEIPv6Route
	SetIncludedRoutes(value INEIPv6Route)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (i NEIPv6Route) Init() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NEIPv6Route) Autorelease() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv6Route creates a new NEIPv6Route instance.
func NewNEIPv6Route() NEIPv6Route {
	class := getNEIPv6RouteClass()
	rv := objc.Send[NEIPv6Route](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize the NEIPv6Route
//
// address: An IPv6 address string. This string is combined with `networkPrefixLength`
// to specify the destination network of the route.
//
// networkPrefixLength: An IPv6 network prefix length. This number is combined with `address` to
// specify the destination network of the route. The network prefix length
// must be an integer between 0 and 128.
//
// # Return Value
// 
// The initialized [NEIPv6Route] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/init(destinationAddress:networkPrefixLength:)
func NewIPv6RouteWithDestinationAddressNetworkPrefixLength(address string, networkPrefixLength foundation.NSNumber) NEIPv6Route {
	instance := getNEIPv6RouteClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDestinationAddress:networkPrefixLength:"), objc.String(address), networkPrefixLength)
	return NEIPv6RouteFromID(rv)
}

// Initialize the NEIPv6Route
//
// address: An IPv6 address string. This string is combined with `networkPrefixLength`
// to specify the destination network of the route.
//
// networkPrefixLength: An IPv6 network prefix length. This number is combined with `address` to
// specify the destination network of the route. The network prefix length
// must be an integer between 0 and 128.
//
// # Return Value
// 
// The initialized [NEIPv6Route] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/init(destinationAddress:networkPrefixLength:)
func (i NEIPv6Route) InitWithDestinationAddressNetworkPrefixLength(address string, networkPrefixLength foundation.NSNumber) NEIPv6Route {
	rv := objc.Send[NEIPv6Route](i.ID, objc.Sel("initWithDestinationAddress:networkPrefixLength:"), objc.String(address), networkPrefixLength)
	return rv
}
func (i NEIPv6Route) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A convenience method for creating the default IPv4 route.
//
// # Return Value
// 
// A [NEIPv6Route] object containing the default IPv6 route.
//
// # Discussion
// 
// Set this route in the `includedRoutes` array in [NEIPv6Settings] to specify
// that all IPv6 network traffic be routed to the TUN interface by default.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/default()
func (_NEIPv6RouteClass NEIPv6RouteClass) DefaultRoute() NEIPv6Route {
	rv := objc.Send[objc.ID](objc.ID(_NEIPv6RouteClass.class), objc.Sel("defaultRoute"))
	return NEIPv6RouteFromID(rv)
}

// The destination network address of the route.
//
// # Discussion
// 
// This string is combined with `destinationNetworkPrefixLength` to specify
// the destination network of the route.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/destinationAddress
func (i NEIPv6Route) DestinationAddress() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("destinationAddress"))
	return foundation.NSStringFromID(rv).String()
}
// The destination network prefix length of the route.
//
// # Discussion
// 
// This string is combined with `destinationAddress` to specify the
// destination network of the route.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/destinationNetworkPrefixLength
func (i NEIPv6Route) DestinationNetworkPrefixLength() foundation.NSNumber {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("destinationNetworkPrefixLength"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// The address of the next-hop gateway of the route.
//
// # Discussion
// 
// The default value of this property is nil. When this property is nil, the
// route’s next-hop gateway will be set to the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/gatewayAddress
func (i NEIPv6Route) GatewayAddress() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("gatewayAddress"))
	return foundation.NSStringFromID(rv).String()
}
func (i NEIPv6Route) SetGatewayAddress(value string) {
	objc.Send[struct{}](i.ID, objc.Sel("setGatewayAddress:"), objc.String(value))
}
// The IPv6 network traffic that the system routes to the primary physical
// interface, not the TUN interface.
//
// See: https://developer.apple.com/documentation/networkextension/neipv6settings/excludedroutes
func (i NEIPv6Route) ExcludedRoutes() INEIPv6Route {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("excludedRoutes"))
	return NEIPv6RouteFromID(objc.ID(rv))
}
func (i NEIPv6Route) SetExcludedRoutes(value INEIPv6Route) {
	objc.Send[struct{}](i.ID, objc.Sel("setExcludedRoutes:"), value)
}
// The IPv6 network traffic that the system routes to the TUN interface.
//
// See: https://developer.apple.com/documentation/networkextension/neipv6settings/includedroutes
func (i NEIPv6Route) IncludedRoutes() INEIPv6Route {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("includedRoutes"))
	return NEIPv6RouteFromID(objc.ID(rv))
}
func (i NEIPv6Route) SetIncludedRoutes(value INEIPv6Route) {
	objc.Send[struct{}](i.ID, objc.Sel("setIncludedRoutes:"), value)
}


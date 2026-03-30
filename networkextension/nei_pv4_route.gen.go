// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEIPv4Route] class.
var (
	_NEIPv4RouteClass     NEIPv4RouteClass
	_NEIPv4RouteClassOnce sync.Once
)

func getNEIPv4RouteClass() NEIPv4RouteClass {
	_NEIPv4RouteClassOnce.Do(func() {
		_NEIPv4RouteClass = NEIPv4RouteClass{class: objc.GetClass("NEIPv4Route")}
	})
	return _NEIPv4RouteClass
}

// GetNEIPv4RouteClass returns the class object for NEIPv4Route.
func GetNEIPv4RouteClass() NEIPv4RouteClass {
	return getNEIPv4RouteClass()
}

type NEIPv4RouteClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEIPv4RouteClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEIPv4RouteClass) Alloc() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The settings for an IPv4 route.
//
// # Creating an IPv4 Route
//
//   - [NEIPv4Route.InitWithDestinationAddressSubnetMask]: Initialize the [NEIPv4Route](<doc://com.apple.networkextension/documentation/NetworkExtension/NEIPv4Route>) object.
//
// # Accessing IPv4 Route Properties
//
//   - [NEIPv4Route.DestinationAddress]: The destination network address of the route.
//   - [NEIPv4Route.DestinationSubnetMask]: The destination network mask of the route.
//   - [NEIPv4Route.GatewayAddress]: The address of the next-hop gateway of the route.
//   - [NEIPv4Route.SetGatewayAddress]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route
type NEIPv4Route struct {
	objectivec.Object
}

// NEIPv4RouteFromID constructs a [NEIPv4Route] from an objc.ID.
//
// The settings for an IPv4 route.
func NEIPv4RouteFromID(id objc.ID) NEIPv4Route {
	return NEIPv4Route{objectivec.Object{ID: id}}
}

// NOTE: NEIPv4Route adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEIPv4Route] class.
//
// # Creating an IPv4 Route
//
//   - [INEIPv4Route.InitWithDestinationAddressSubnetMask]: Initialize the [NEIPv4Route](<doc://com.apple.networkextension/documentation/NetworkExtension/NEIPv4Route>) object.
//
// # Accessing IPv4 Route Properties
//
//   - [INEIPv4Route.DestinationAddress]: The destination network address of the route.
//   - [INEIPv4Route.DestinationSubnetMask]: The destination network mask of the route.
//   - [INEIPv4Route.GatewayAddress]: The address of the next-hop gateway of the route.
//   - [INEIPv4Route.SetGatewayAddress]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route
type INEIPv4Route interface {
	objectivec.IObject

	// Topic: Creating an IPv4 Route

	// Initialize the [NEIPv4Route](<doc://com.apple.networkextension/documentation/NetworkExtension/NEIPv4Route>) object.
	InitWithDestinationAddressSubnetMask(address string, subnetMask string) NEIPv4Route

	// Topic: Accessing IPv4 Route Properties

	// The destination network address of the route.
	DestinationAddress() string
	// The destination network mask of the route.
	DestinationSubnetMask() string
	// The address of the next-hop gateway of the route.
	GatewayAddress() string
	SetGatewayAddress(value string)

	// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
	ExcludedRoutes() INEIPv4Route
	SetExcludedRoutes(value INEIPv4Route)
	// The IPv4 network traffic that the system routes to the TUN interface.
	IncludedRoutes() INEIPv4Route
	SetIncludedRoutes(value INEIPv4Route)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (i NEIPv4Route) Init() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NEIPv4Route) Autorelease() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv4Route creates a new NEIPv4Route instance.
func NewNEIPv4Route() NEIPv4Route {
	class := getNEIPv4RouteClass()
	rv := objc.Send[NEIPv4Route](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize the [NEIPv4Route] object.
//
// address: An IPv4 address string. This string is combined with `subnetMask` to
// specify the destination network of the route.
//
// subnetMask: An IPv4 network mask string. This string is combined with `address` to
// specify the destination network of the route.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/init(destinationAddress:subnetMask:)
func NewIPv4RouteWithDestinationAddressSubnetMask(address string, subnetMask string) NEIPv4Route {
	instance := getNEIPv4RouteClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDestinationAddress:subnetMask:"), objc.String(address), objc.String(subnetMask))
	return NEIPv4RouteFromID(rv)
}

// Initialize the [NEIPv4Route] object.
//
// address: An IPv4 address string. This string is combined with `subnetMask` to
// specify the destination network of the route.
//
// subnetMask: An IPv4 network mask string. This string is combined with `address` to
// specify the destination network of the route.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/init(destinationAddress:subnetMask:)
func (i NEIPv4Route) InitWithDestinationAddressSubnetMask(address string, subnetMask string) NEIPv4Route {
	rv := objc.Send[NEIPv4Route](i.ID, objc.Sel("initWithDestinationAddress:subnetMask:"), objc.String(address), objc.String(subnetMask))
	return rv
}
func (i NEIPv4Route) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A convenience method for creating the default IPv4 route.
//
// # Return Value
//
// An [NEIPv4Route] object containing the default IPv4 route.
//
// # Discussion
//
// Set this route in the `includedRoutes` array in the [NEIPv4Settings] object
// to specify that all IPv4 network traffic be routed to the TUN interface by
// default.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/default()
func (_NEIPv4RouteClass NEIPv4RouteClass) DefaultRoute() NEIPv4Route {
	rv := objc.Send[objc.ID](objc.ID(_NEIPv4RouteClass.class), objc.Sel("defaultRoute"))
	return NEIPv4RouteFromID(rv)
}

// The destination network address of the route.
//
// # Discussion
//
// This string is combined with `destinationSubnetMask` to specify the
// destination network of the route.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/destinationAddress
func (i NEIPv4Route) DestinationAddress() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("destinationAddress"))
	return foundation.NSStringFromID(rv).String()
}

// The destination network mask of the route.
//
// # Discussion
//
// This string is combined with `destinationAddress` to specify the
// destination network of the route.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/destinationSubnetMask
func (i NEIPv4Route) DestinationSubnetMask() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("destinationSubnetMask"))
	return foundation.NSStringFromID(rv).String()
}

// The address of the next-hop gateway of the route.
//
// # Discussion
//
// The default value of this property is nil. When this property is nil, the
// route’s next-hop gateway will be set to the TUN interface unless it is a
// Split Exclude route.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/gatewayAddress
func (i NEIPv4Route) GatewayAddress() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("gatewayAddress"))
	return foundation.NSStringFromID(rv).String()
}
func (i NEIPv4Route) SetGatewayAddress(value string) {
	objc.Send[struct{}](i.ID, objc.Sel("setGatewayAddress:"), objc.String(value))
}

// The IPv4 network traffic that the system routes to the primary physical
// interface, not the TUN interface.
//
// See: https://developer.apple.com/documentation/networkextension/neipv4settings/excludedroutes
func (i NEIPv4Route) ExcludedRoutes() INEIPv4Route {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("excludedRoutes"))
	return NEIPv4RouteFromID(objc.ID(rv))
}
func (i NEIPv4Route) SetExcludedRoutes(value INEIPv4Route) {
	objc.Send[struct{}](i.ID, objc.Sel("setExcludedRoutes:"), value)
}

// The IPv4 network traffic that the system routes to the TUN interface.
//
// See: https://developer.apple.com/documentation/networkextension/neipv4settings/includedroutes
func (i NEIPv4Route) IncludedRoutes() INEIPv4Route {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("includedRoutes"))
	return NEIPv4RouteFromID(objc.ID(rv))
}
func (i NEIPv4Route) SetIncludedRoutes(value INEIPv4Route) {
	objc.Send[struct{}](i.ID, objc.Sel("setIncludedRoutes:"), value)
}

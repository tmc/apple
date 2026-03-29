// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NEEthernetTunnelProvider] class.
var (
	_NEEthernetTunnelProviderClass     NEEthernetTunnelProviderClass
	_NEEthernetTunnelProviderClassOnce sync.Once
)

func getNEEthernetTunnelProviderClass() NEEthernetTunnelProviderClass {
	_NEEthernetTunnelProviderClassOnce.Do(func() {
		_NEEthernetTunnelProviderClass = NEEthernetTunnelProviderClass{class: objc.GetClass("NEEthernetTunnelProvider")}
	})
	return _NEEthernetTunnelProviderClass
}

// GetNEEthernetTunnelProviderClass returns the class object for NEEthernetTunnelProvider.
func GetNEEthernetTunnelProviderClass() NEEthernetTunnelProviderClass {
	return getNEEthernetTunnelProviderClass()
}

type NEEthernetTunnelProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEEthernetTunnelProviderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEEthernetTunnelProviderClass) Alloc() NEEthernetTunnelProvider {
	rv := objc.Send[NEEthernetTunnelProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A type that implements the client side of a custom link-layer packet
// tunneling protocol.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelProvider
type NEEthernetTunnelProvider struct {
	NEPacketTunnelProvider
}

// NEEthernetTunnelProviderFromID constructs a [NEEthernetTunnelProvider] from an objc.ID.
//
// A type that implements the client side of a custom link-layer packet
// tunneling protocol.
func NEEthernetTunnelProviderFromID(id objc.ID) NEEthernetTunnelProvider {
	return NEEthernetTunnelProvider{NEPacketTunnelProvider: NEPacketTunnelProviderFromID(id)}
}
// NOTE: NEEthernetTunnelProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEEthernetTunnelProvider] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelProvider
type INEEthernetTunnelProvider interface {
	INEPacketTunnelProvider
}

// Init initializes the instance.
func (e NEEthernetTunnelProvider) Init() NEEthernetTunnelProvider {
	rv := objc.Send[NEEthernetTunnelProvider](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NEEthernetTunnelProvider) Autorelease() NEEthernetTunnelProvider {
	rv := objc.Send[NEEthernetTunnelProvider](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEEthernetTunnelProvider creates a new NEEthernetTunnelProvider instance.
func NewNEEthernetTunnelProvider() NEEthernetTunnelProvider {
	class := getNEEthernetTunnelProviderClass()
	rv := objc.Send[NEEthernetTunnelProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}


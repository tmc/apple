// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NETransparentProxyProvider] class.
var (
	_NETransparentProxyProviderClass     NETransparentProxyProviderClass
	_NETransparentProxyProviderClassOnce sync.Once
)

func getNETransparentProxyProviderClass() NETransparentProxyProviderClass {
	_NETransparentProxyProviderClassOnce.Do(func() {
		_NETransparentProxyProviderClass = NETransparentProxyProviderClass{class: objc.GetClass("NETransparentProxyProvider")}
	})
	return _NETransparentProxyProviderClass
}

// GetNETransparentProxyProviderClass returns the class object for NETransparentProxyProvider.
func GetNETransparentProxyProviderClass() NETransparentProxyProviderClass {
	return getNETransparentProxyProviderClass()
}

type NETransparentProxyProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NETransparentProxyProviderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NETransparentProxyProviderClass) Alloc() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that implements the client side of a custom transparent network
// proxy solution.
//
// # Overview
//
// The [NETransparentProxyProvider] class has the following behavior
// differences from its superclass [NEAppProxyProvider]:
//
// - Returning [NO] from [HandleNewFlow] and
// [HandleNewUDPFlowInitialRemoteEndpoint] causes the flow to proceed to
// communicate directly with the flow’s ultimate destination, instead of
// closing the flow with a “Connection Refused” error. - This provider
// ignores [NEDNSSettings] and [NEProxySettings] specified within
// [NETransparentProxyNetworkSettings]. Flows that match the
// [NETransparentProxyProvider.IncludedNetworkRules] within [NETransparentProxyNetworkSettings] use the
// same DNS and proxy settings that other flows on the system currently use. -
// Flows that are created using a “connect by name” API (such as [Network]
// framework or [URLSession]) that match the [NETransparentProxyProvider.IncludedNetworkRules] don’t
// bypass DNS resolution.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyProvider
//
// [Network]: https://developer.apple.com/documentation/Network
// [URLSession]: https://developer.apple.com/documentation/Foundation/URLSession
type NETransparentProxyProvider struct {
	NEAppProxyProvider
}

// NETransparentProxyProviderFromID constructs a [NETransparentProxyProvider] from an objc.ID.
//
// An object that implements the client side of a custom transparent network
// proxy solution.
func NETransparentProxyProviderFromID(id objc.ID) NETransparentProxyProvider {
	return NETransparentProxyProvider{NEAppProxyProvider: NEAppProxyProviderFromID(id)}
}

// NOTE: NETransparentProxyProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NETransparentProxyProvider] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyProvider
type INETransparentProxyProvider interface {
	INEAppProxyProvider

	// An array of rules that collectively specify what traffic to route through the transparent proxy.
	IncludedNetworkRules() INENetworkRule
	SetIncludedNetworkRules(value INENetworkRule)
}

// Init initializes the instance.
func (t NETransparentProxyProvider) Init() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NETransparentProxyProvider) Autorelease() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETransparentProxyProvider creates a new NETransparentProxyProvider instance.
func NewNETransparentProxyProvider() NETransparentProxyProvider {
	class := getNETransparentProxyProviderClass()
	rv := objc.Send[NETransparentProxyProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array of rules that collectively specify what traffic to route through
// the transparent proxy.
//
// See: https://developer.apple.com/documentation/networkextension/netransparentproxynetworksettings/includednetworkrules
func (t NETransparentProxyProvider) IncludedNetworkRules() INENetworkRule {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("includedNetworkRules"))
	return NENetworkRuleFromID(objc.ID(rv))
}
func (t NETransparentProxyProvider) SetIncludedNetworkRules(value INENetworkRule) {
	objc.Send[struct{}](t.ID, objc.Sel("setIncludedNetworkRules:"), value)
}

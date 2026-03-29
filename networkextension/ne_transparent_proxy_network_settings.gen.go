// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NETransparentProxyNetworkSettings] class.
var (
	_NETransparentProxyNetworkSettingsClass     NETransparentProxyNetworkSettingsClass
	_NETransparentProxyNetworkSettingsClassOnce sync.Once
)

func getNETransparentProxyNetworkSettingsClass() NETransparentProxyNetworkSettingsClass {
	_NETransparentProxyNetworkSettingsClassOnce.Do(func() {
		_NETransparentProxyNetworkSettingsClass = NETransparentProxyNetworkSettingsClass{class: objc.GetClass("NETransparentProxyNetworkSettings")}
	})
	return _NETransparentProxyNetworkSettingsClass
}

// GetNETransparentProxyNetworkSettingsClass returns the class object for NETransparentProxyNetworkSettings.
func GetNETransparentProxyNetworkSettingsClass() NETransparentProxyNetworkSettingsClass {
	return getNETransparentProxyNetworkSettingsClass()
}

type NETransparentProxyNetworkSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NETransparentProxyNetworkSettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NETransparentProxyNetworkSettingsClass) Alloc() NETransparentProxyNetworkSettings {
	rv := objc.Send[NETransparentProxyNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specification of what traffic to route through a transparent proxy.
//
// # Overview
// 
// A proxy network settings object contains two properties: an array of rules
// to include traffic ([NETransparentProxyNetworkSettings.IncludedNetworkRules]) and an array of rules to
// exclude traffic ([NETransparentProxyNetworkSettings.ExcludedNetworkRules]). The exclusion rules take
// prirority. Therefore, if a given flow matches any of the
// [NETransparentProxyNetworkSettings.ExcludedNetworkRules], evaluation ends and the flow doesn’t route to the
// proxy. If there’s no match, then evaluation continues and attempts to
// match the flow against the [NETransparentProxyNetworkSettings.IncludedNetworkRules].
//
// # Traffic routing rules
//
//   - [NETransparentProxyNetworkSettings.IncludedNetworkRules]: An array of rules that collectively specify what traffic to route through the transparent proxy.
//   - [NETransparentProxyNetworkSettings.SetIncludedNetworkRules]
//   - [NETransparentProxyNetworkSettings.ExcludedNetworkRules]: An array of rules that collectively specify what traffic to not route through the transparent proxy.
//   - [NETransparentProxyNetworkSettings.SetExcludedNetworkRules]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings
type NETransparentProxyNetworkSettings struct {
	NETunnelNetworkSettings
}

// NETransparentProxyNetworkSettingsFromID constructs a [NETransparentProxyNetworkSettings] from an objc.ID.
//
// A specification of what traffic to route through a transparent proxy.
func NETransparentProxyNetworkSettingsFromID(id objc.ID) NETransparentProxyNetworkSettings {
	return NETransparentProxyNetworkSettings{NETunnelNetworkSettings: NETunnelNetworkSettingsFromID(id)}
}
// NOTE: NETransparentProxyNetworkSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NETransparentProxyNetworkSettings] class.
//
// # Traffic routing rules
//
//   - [INETransparentProxyNetworkSettings.IncludedNetworkRules]: An array of rules that collectively specify what traffic to route through the transparent proxy.
//   - [INETransparentProxyNetworkSettings.SetIncludedNetworkRules]
//   - [INETransparentProxyNetworkSettings.ExcludedNetworkRules]: An array of rules that collectively specify what traffic to not route through the transparent proxy.
//   - [INETransparentProxyNetworkSettings.SetExcludedNetworkRules]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings
type INETransparentProxyNetworkSettings interface {
	INETunnelNetworkSettings

	// Topic: Traffic routing rules

	// An array of rules that collectively specify what traffic to route through the transparent proxy.
	IncludedNetworkRules() []NENetworkRule
	SetIncludedNetworkRules(value []NENetworkRule)
	// An array of rules that collectively specify what traffic to not route through the transparent proxy.
	ExcludedNetworkRules() []NENetworkRule
	SetExcludedNetworkRules(value []NENetworkRule)
}

// Init initializes the instance.
func (t NETransparentProxyNetworkSettings) Init() NETransparentProxyNetworkSettings {
	rv := objc.Send[NETransparentProxyNetworkSettings](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NETransparentProxyNetworkSettings) Autorelease() NETransparentProxyNetworkSettings {
	rv := objc.Send[NETransparentProxyNetworkSettings](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETransparentProxyNetworkSettings creates a new NETransparentProxyNetworkSettings instance.
func NewNETransparentProxyNetworkSettings() NETransparentProxyNetworkSettings {
	class := getNETransparentProxyNetworkSettingsClass()
	rv := objc.Send[NETransparentProxyNetworkSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize a [NETunnelNetworkSettings] object.
//
// address: The IP address of the tunnel server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/init(tunnelRemoteAddress:)
func NewTransparentProxyNetworkSettingsWithTunnelRemoteAddress(address string) NETransparentProxyNetworkSettings {
	instance := getNETransparentProxyNetworkSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTunnelRemoteAddress:"), objc.String(address))
	return NETransparentProxyNetworkSettingsFromID(rv)
}

// An array of rules that collectively specify what traffic to route through
// the transparent proxy.
//
// # Discussion
// 
// The following restrictions apply to each rule in the array:
// 
// - If the port string of the endpoint is `0` or is the empty string, then
// the address of the endpoint must be a non-wildcard address, such as
// `0.0.0.0` or `::`. - If the address is a wildcard address (such as
// `0.0.0.0` or `::)`, then the port string of the endpoint must be non-empty
// and must not be `0`. - A port string of `53` is not allowed. Use
// Destination Domain-based rules to match DNS traffic. - The
// [MatchLocalNetwork] property must be `nil`. - The [MatchDirection] property
// must be [NETrafficDirection.outbound].
//
// [NETrafficDirection.outbound]: https://developer.apple.com/documentation/NetworkExtension/NETrafficDirection/outbound
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings/includedNetworkRules
func (t NETransparentProxyNetworkSettings) IncludedNetworkRules() []NENetworkRule {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("includedNetworkRules"))
	return objc.ConvertSlice(rv, func(id objc.ID) NENetworkRule {
		return NENetworkRuleFromID(id)
	})
}
func (t NETransparentProxyNetworkSettings) SetIncludedNetworkRules(value []NENetworkRule) {
	objc.Send[struct{}](t.ID, objc.Sel("setIncludedNetworkRules:"), objectivec.IObjectSliceToNSArray(value))
}
// An array of rules that collectively specify what traffic to not route
// through the transparent proxy.
//
// # Discussion
// 
// The following restrictions apply to each rule in the array:
// 
// - If the port string of the endpoint is `0` or is the empty string, then
// the address of the endpoint must be a non-wildcard address, such as
// `0.0.0.0` or `::`. - If the address is a wildcard address (such as
// `0.0.0.0` or `::)`, then the port string of the endpoint must be non-empty
// and must not be `0`. - A port string of `53` is not allowed. Use
// Destination Domain-based rules to match DNS traffic. - The
// [MatchLocalNetwork] property must be `nil`. - The [MatchDirection] property
// must be [NETrafficDirection.outbound].
//
// [NETrafficDirection.outbound]: https://developer.apple.com/documentation/NetworkExtension/NETrafficDirection/outbound
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings/excludedNetworkRules
func (t NETransparentProxyNetworkSettings) ExcludedNetworkRules() []NENetworkRule {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("excludedNetworkRules"))
	return objc.ConvertSlice(rv, func(id objc.ID) NENetworkRule {
		return NENetworkRuleFromID(id)
	})
}
func (t NETransparentProxyNetworkSettings) SetExcludedNetworkRules(value []NENetworkRule) {
	objc.Send[struct{}](t.ID, objc.Sel("setExcludedNetworkRules:"), objectivec.IObjectSliceToNSArray(value))
}


// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEOnDemandRule] class.
var (
	_NEOnDemandRuleClass     NEOnDemandRuleClass
	_NEOnDemandRuleClassOnce sync.Once
)

func getNEOnDemandRuleClass() NEOnDemandRuleClass {
	_NEOnDemandRuleClassOnce.Do(func() {
		_NEOnDemandRuleClass = NEOnDemandRuleClass{class: objc.GetClass("NEOnDemandRule")}
	})
	return _NEOnDemandRuleClass
}

// GetNEOnDemandRuleClass returns the class object for NEOnDemandRule.
func GetNEOnDemandRuleClass() NEOnDemandRuleClass {
	return getNEOnDemandRuleClass()
}

type NEOnDemandRuleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEOnDemandRuleClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEOnDemandRuleClass) Alloc() NEOnDemandRule {
	rv := objc.Send[NEOnDemandRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A base class shared by all VPN On Demand rules.
//
// # Overview
// 
// Each rule is defined by a single action and a set of optional matching
// conditions. The action defines how the system should trigger the VPN when
// the conditions are met, such as connecting automatically for all
// connections, connecting conditionally, or disconnecting. The optional
// conditions describe parameters of a network. Some common rules include
// disconnecting the VPN on a trusted, internal network, and triggering on all
// other networks. When rules are defined in an array, they are evaluated in
// order and the action of the first rule to match all conditions is chosen.
// 
// Instances of the [NEOnDemandRule] class should be created through one of
// its subclasses: [NEOnDemandRuleConnect], [NEOnDemandRuleDisconnect],
// [NEOnDemandRuleEvaluateConnection], or [NEOnDemandRuleIgnore].
//
// # Accessing match parameters
//
//   - [NEOnDemandRule.DNSSearchDomainMatch]: DNS search domains that identify a network.
//   - [NEOnDemandRule.SetDNSSearchDomainMatch]
//   - [NEOnDemandRule.DNSServerAddressMatch]: DNS server addresses that identify a network.
//   - [NEOnDemandRule.SetDNSServerAddressMatch]
//   - [NEOnDemandRule.InterfaceTypeMatch]: An interface type to identify a network.
//   - [NEOnDemandRule.SetInterfaceTypeMatch]
//   - [NEOnDemandRule.SSIDMatch]: SSIDs that identify a network.
//   - [NEOnDemandRule.SetSSIDMatch]
//   - [NEOnDemandRule.ProbeURL]: A URL to probe when all other network identifiers match to validate that an expected resource is available.
//   - [NEOnDemandRule.SetProbeURL]
//
// # Accessing the rule action
//
//   - [NEOnDemandRule.Action]: The action of the On Demand Rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule
type NEOnDemandRule struct {
	objectivec.Object
}

// NEOnDemandRuleFromID constructs a [NEOnDemandRule] from an objc.ID.
//
// A base class shared by all VPN On Demand rules.
func NEOnDemandRuleFromID(id objc.ID) NEOnDemandRule {
	return NEOnDemandRule{objectivec.Object{ID: id}}
}
// NOTE: NEOnDemandRule adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEOnDemandRule] class.
//
// # Accessing match parameters
//
//   - [INEOnDemandRule.DNSSearchDomainMatch]: DNS search domains that identify a network.
//   - [INEOnDemandRule.SetDNSSearchDomainMatch]
//   - [INEOnDemandRule.DNSServerAddressMatch]: DNS server addresses that identify a network.
//   - [INEOnDemandRule.SetDNSServerAddressMatch]
//   - [INEOnDemandRule.InterfaceTypeMatch]: An interface type to identify a network.
//   - [INEOnDemandRule.SetInterfaceTypeMatch]
//   - [INEOnDemandRule.SSIDMatch]: SSIDs that identify a network.
//   - [INEOnDemandRule.SetSSIDMatch]
//   - [INEOnDemandRule.ProbeURL]: A URL to probe when all other network identifiers match to validate that an expected resource is available.
//   - [INEOnDemandRule.SetProbeURL]
//
// # Accessing the rule action
//
//   - [INEOnDemandRule.Action]: The action of the On Demand Rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule
type INEOnDemandRule interface {
	objectivec.IObject

	// Topic: Accessing match parameters

	// DNS search domains that identify a network.
	DNSSearchDomainMatch() []string
	SetDNSSearchDomainMatch(value []string)
	// DNS server addresses that identify a network.
	DNSServerAddressMatch() []string
	SetDNSServerAddressMatch(value []string)
	// An interface type to identify a network.
	InterfaceTypeMatch() NEOnDemandRuleInterfaceType
	SetInterfaceTypeMatch(value NEOnDemandRuleInterfaceType)
	// SSIDs that identify a network.
	SSIDMatch() []string
	SetSSIDMatch(value []string)
	// A URL to probe when all other network identifiers match to validate that an expected resource is available.
	ProbeURL() foundation.INSURL
	SetProbeURL(value foundation.INSURL)

	// Topic: Accessing the rule action

	// The action of the On Demand Rule.
	Action() NEOnDemandRuleAction

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (o NEOnDemandRule) Init() NEOnDemandRule {
	rv := objc.Send[NEOnDemandRule](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NEOnDemandRule) Autorelease() NEOnDemandRule {
	rv := objc.Send[NEOnDemandRule](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEOnDemandRule creates a new NEOnDemandRule instance.
func NewNEOnDemandRule() NEOnDemandRule {
	class := getNEOnDemandRuleClass()
	rv := objc.Send[NEOnDemandRule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (o NEOnDemandRule) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeWithCoder:"), coder)
}

// DNS search domains that identify a network.
//
// # Discussion
// 
// An array of [NSString] objects. If the current default search domain is
// equal to one of the strings in this array and all of the other conditions
// in the rule match, then the rule matches. If this property is nil (the
// default), then the current default search domain does not factor into the
// rule match.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/dnsSearchDomainMatch
func (o NEOnDemandRule) DNSSearchDomainMatch() []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("DNSSearchDomainMatch"))
	return objc.ConvertSliceToStrings(rv)
}
func (o NEOnDemandRule) SetDNSSearchDomainMatch(value []string) {
	objc.Send[struct{}](o.ID, objc.Sel("setDNSSearchDomainMatch:"), objectivec.StringSliceToNSArray(value))
}
// DNS server addresses that identify a network.
//
// # Discussion
// 
// An array of DNS server IP addresses represented as [NSString] objects. If
// each of the current default DNS servers is equal to one of the strings in
// this array and all of the other conditions in the rule match, then the rule
// matches. If this property is nil (the default), then the default DNS
// servers do not factor into the rule match.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/dnsServerAddressMatch
func (o NEOnDemandRule) DNSServerAddressMatch() []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("DNSServerAddressMatch"))
	return objc.ConvertSliceToStrings(rv)
}
func (o NEOnDemandRule) SetDNSServerAddressMatch(value []string) {
	objc.Send[struct{}](o.ID, objc.Sel("setDNSServerAddressMatch:"), objectivec.StringSliceToNSArray(value))
}
// An interface type to identify a network.
//
// # Discussion
// 
// The type of interface that this rule matches. If the current primary
// network interface is of this type and all of the other conditions in the
// rule match, then the rule matches. If this property is
// [NEOnDemandRuleInterfaceTypeAny] (the default), then the current primary
// interface type does not factor into the rule match.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/interfaceTypeMatch
func (o NEOnDemandRule) InterfaceTypeMatch() NEOnDemandRuleInterfaceType {
	rv := objc.Send[NEOnDemandRuleInterfaceType](o.ID, objc.Sel("interfaceTypeMatch"))
	return NEOnDemandRuleInterfaceType(rv)
}
func (o NEOnDemandRule) SetInterfaceTypeMatch(value NEOnDemandRuleInterfaceType) {
	objc.Send[struct{}](o.ID, objc.Sel("setInterfaceTypeMatch:"), value)
}
// SSIDs that identify a network.
//
// # Discussion
// 
// An array of [NSString] objects. If the Service Set Identifier (SSID) of the
// current primary connected network matches one of the strings in this array
// and all of the other conditions in the rule match, then the rule matches.
// If this property is nil (the default), then the current primary connected
// network SSID does not factor into the rule match.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/ssidMatch
func (o NEOnDemandRule) SSIDMatch() []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("SSIDMatch"))
	return objc.ConvertSliceToStrings(rv)
}
func (o NEOnDemandRule) SetSSIDMatch(value []string) {
	objc.Send[struct{}](o.ID, objc.Sel("setSSIDMatch:"), objectivec.StringSliceToNSArray(value))
}
// A URL to probe when all other network identifiers match to validate that an
// expected resource is available.
//
// # Discussion
// 
// An HTTP or HTTPS URL. If a request sent to this URL results in a HTTP 200
// OK response and all of the other conditions in the rule match, then then
// rule matches. If this property is nil (the default), then an HTTP request
// does not factor into the rule match.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/probeURL
func (o NEOnDemandRule) ProbeURL() foundation.INSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("probeURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (o NEOnDemandRule) SetProbeURL(value foundation.INSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("setProbeURL:"), value)
}
// The action of the On Demand Rule.
//
// # Discussion
// 
// The action of the On Demand Rule represents the behavior for triggering the
// corresponding VPN when the rule conditions are matched.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/action
func (o NEOnDemandRule) Action() NEOnDemandRuleAction {
	rv := objc.Send[NEOnDemandRuleAction](o.ID, objc.Sel("action"))
	return NEOnDemandRuleAction(rv)
}


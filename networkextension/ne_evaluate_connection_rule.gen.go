// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEEvaluateConnectionRule] class.
var (
	_NEEvaluateConnectionRuleClass     NEEvaluateConnectionRuleClass
	_NEEvaluateConnectionRuleClassOnce sync.Once
)

func getNEEvaluateConnectionRuleClass() NEEvaluateConnectionRuleClass {
	_NEEvaluateConnectionRuleClassOnce.Do(func() {
		_NEEvaluateConnectionRuleClass = NEEvaluateConnectionRuleClass{class: objc.GetClass("NEEvaluateConnectionRule")}
	})
	return _NEEvaluateConnectionRuleClass
}

// GetNEEvaluateConnectionRuleClass returns the class object for NEEvaluateConnectionRule.
func GetNEEvaluateConnectionRuleClass() NEEvaluateConnectionRuleClass {
	return getNEEvaluateConnectionRuleClass()
}

type NEEvaluateConnectionRuleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEEvaluateConnectionRuleClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEEvaluateConnectionRuleClass) Alloc() NEEvaluateConnectionRule {
	rv := objc.Send[NEEvaluateConnectionRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// [NEEvaluateConnectionRule] associates properties of network connections
// with an action.
//
// # Initializing a Rule
//
//   - [NEEvaluateConnectionRule.InitWithMatchDomainsAndAction]: Initialize an [NEEvaluateConnectionRule] instance with a list of destination host domains and an action.
//
// # Accessing Rule Match Properties
//
//   - [NEEvaluateConnectionRule.MatchDomains]: An array of domains used to match the destination hostname of connections. If the destination hostname of a connection matches any of the domains in the array, then the connection matches the rule. Each domain is matched against the destination hostname using suffix matching, and each label in the domain must match an entire label in the hostname. For example, the domain `example.Com()` will match the hostname `www.ExampleXCUIElementTypeCom()` but not `www.AnotherexampleXCUIElementTypeCom()`.
//   - [NEEvaluateConnectionRule.UseDNSServers]: If the rule matches the connection being established and the action is [NEEvaluateConnectionRuleActionConnectIfNeeded], the DNS servers specified in this array are used to resolve the destination hostname of the connection while evaluating connectivity to the destination of the connection. If the resolution fails for any reason, the VPN is started.
//   - [NEEvaluateConnectionRule.SetUseDNSServers]
//   - [NEEvaluateConnectionRule.ProbeURL]: An HTTP or HTTPS URL. If the rule matches the connection being established and the action is [NEEvaluateConnectionRuleActionConnectIfNeeded] and a request sent to this URL results in a response with an HTTP response code other than 200, then the VPN is started.
//   - [NEEvaluateConnectionRule.SetProbeURL]
//
// # Accessing the Rule Action
//
//   - [NEEvaluateConnectionRule.Action]: The action to take if the properties of the network connection being established match the rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule
type NEEvaluateConnectionRule struct {
	objectivec.Object
}

// NEEvaluateConnectionRuleFromID constructs a [NEEvaluateConnectionRule] from an objc.ID.
//
// [NEEvaluateConnectionRule] associates properties of network connections
// with an action.
func NEEvaluateConnectionRuleFromID(id objc.ID) NEEvaluateConnectionRule {
	return NEEvaluateConnectionRule{objectivec.Object{ID: id}}
}
// NOTE: NEEvaluateConnectionRule adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEEvaluateConnectionRule] class.
//
// # Initializing a Rule
//
//   - [INEEvaluateConnectionRule.InitWithMatchDomainsAndAction]: Initialize an [NEEvaluateConnectionRule] instance with a list of destination host domains and an action.
//
// # Accessing Rule Match Properties
//
//   - [INEEvaluateConnectionRule.MatchDomains]: An array of domains used to match the destination hostname of connections. If the destination hostname of a connection matches any of the domains in the array, then the connection matches the rule. Each domain is matched against the destination hostname using suffix matching, and each label in the domain must match an entire label in the hostname. For example, the domain `example.Com()` will match the hostname `www.ExampleXCUIElementTypeCom()` but not `www.AnotherexampleXCUIElementTypeCom()`.
//   - [INEEvaluateConnectionRule.UseDNSServers]: If the rule matches the connection being established and the action is [NEEvaluateConnectionRuleActionConnectIfNeeded], the DNS servers specified in this array are used to resolve the destination hostname of the connection while evaluating connectivity to the destination of the connection. If the resolution fails for any reason, the VPN is started.
//   - [INEEvaluateConnectionRule.SetUseDNSServers]
//   - [INEEvaluateConnectionRule.ProbeURL]: An HTTP or HTTPS URL. If the rule matches the connection being established and the action is [NEEvaluateConnectionRuleActionConnectIfNeeded] and a request sent to this URL results in a response with an HTTP response code other than 200, then the VPN is started.
//   - [INEEvaluateConnectionRule.SetProbeURL]
//
// # Accessing the Rule Action
//
//   - [INEEvaluateConnectionRule.Action]: The action to take if the properties of the network connection being established match the rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule
type INEEvaluateConnectionRule interface {
	objectivec.IObject

	// Topic: Initializing a Rule

	// Initialize an [NEEvaluateConnectionRule] instance with a list of destination host domains and an action.
	InitWithMatchDomainsAndAction(domains []string, action NEEvaluateConnectionRuleAction) NEEvaluateConnectionRule

	// Topic: Accessing Rule Match Properties

	// An array of domains used to match the destination hostname of connections. If the destination hostname of a connection matches any of the domains in the array, then the connection matches the rule. Each domain is matched against the destination hostname using suffix matching, and each label in the domain must match an entire label in the hostname. For example, the domain `example.Com()` will match the hostname `www.ExampleXCUIElementTypeCom()` but not `www.AnotherexampleXCUIElementTypeCom()`.
	MatchDomains() []string
	// If the rule matches the connection being established and the action is [NEEvaluateConnectionRuleActionConnectIfNeeded], the DNS servers specified in this array are used to resolve the destination hostname of the connection while evaluating connectivity to the destination of the connection. If the resolution fails for any reason, the VPN is started.
	UseDNSServers() []string
	SetUseDNSServers(value []string)
	// An HTTP or HTTPS URL. If the rule matches the connection being established and the action is [NEEvaluateConnectionRuleActionConnectIfNeeded] and a request sent to this URL results in a response with an HTTP response code other than 200, then the VPN is started.
	ProbeURL() foundation.INSURL
	SetProbeURL(value foundation.INSURL)

	// Topic: Accessing the Rule Action

	// The action to take if the properties of the network connection being established match the rule.
	Action() NEEvaluateConnectionRuleAction

	// An array of 
	ConnectionRules() INEEvaluateConnectionRule
	SetConnectionRules(value INEEvaluateConnectionRule)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (e NEEvaluateConnectionRule) Init() NEEvaluateConnectionRule {
	rv := objc.Send[NEEvaluateConnectionRule](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NEEvaluateConnectionRule) Autorelease() NEEvaluateConnectionRule {
	rv := objc.Send[NEEvaluateConnectionRule](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEEvaluateConnectionRule creates a new NEEvaluateConnectionRule instance.
func NewNEEvaluateConnectionRule() NEEvaluateConnectionRule {
	class := getNEEvaluateConnectionRuleClass()
	rv := objc.Send[NEEvaluateConnectionRule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize an [NEEvaluateConnectionRule] instance with a list of
// destination host domains and an action.
//
// domains: An array of domains used to match the destination hostname of connections.
// If the destination hostname of a connection matches any of the domains in
// the array, then the connection matches the rule. Each domain is matched
// against the destination hostname using suffix matching, and each label in
// the domain must match an entire label in the hostname. For example, the
// domain `example.Com()` will match the hostname
// `www.ExampleXCUIElementTypeCom()` but not
// `www.AnotherexampleXCUIElementTypeCom()`.
//
// action: The action to apply for connections matching the rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/init(matchDomains:andAction:)
func NewEvaluateConnectionRuleWithMatchDomainsAndAction(domains []string, action NEEvaluateConnectionRuleAction) NEEvaluateConnectionRule {
	instance := getNEEvaluateConnectionRuleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMatchDomains:andAction:"), objectivec.StringSliceToNSArray(domains), action)
	return NEEvaluateConnectionRuleFromID(rv)
}

// Initialize an [NEEvaluateConnectionRule] instance with a list of
// destination host domains and an action.
//
// domains: An array of domains used to match the destination hostname of connections.
// If the destination hostname of a connection matches any of the domains in
// the array, then the connection matches the rule. Each domain is matched
// against the destination hostname using suffix matching, and each label in
// the domain must match an entire label in the hostname. For example, the
// domain `example.Com()` will match the hostname
// `www.ExampleXCUIElementTypeCom()` but not
// `www.AnotherexampleXCUIElementTypeCom()`.
//
// action: The action to apply for connections matching the rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/init(matchDomains:andAction:)
func (e NEEvaluateConnectionRule) InitWithMatchDomainsAndAction(domains []string, action NEEvaluateConnectionRuleAction) NEEvaluateConnectionRule {
	rv := objc.Send[NEEvaluateConnectionRule](e.ID, objc.Sel("initWithMatchDomains:andAction:"), objectivec.StringSliceToNSArray(domains), action)
	return rv
}
func (e NEEvaluateConnectionRule) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}

// An array of domains used to match the destination hostname of connections.
// If the destination hostname of a connection matches any of the domains in
// the array, then the connection matches the rule. Each domain is matched
// against the destination hostname using suffix matching, and each label in
// the domain must match an entire label in the hostname. For example, the
// domain `example.Com()` will match the hostname
// `www.ExampleXCUIElementTypeCom()` but not
// `www.AnotherexampleXCUIElementTypeCom()`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/matchDomains
func (e NEEvaluateConnectionRule) MatchDomains() []string {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("matchDomains"))
	return objc.ConvertSliceToStrings(rv)
}
// If the rule matches the connection being established and the action is
// [NEEvaluateConnectionRuleActionConnectIfNeeded], the DNS servers specified
// in this array are used to resolve the destination hostname of the
// connection while evaluating connectivity to the destination of the
// connection. If the resolution fails for any reason, the VPN is started.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/useDNSServers
func (e NEEvaluateConnectionRule) UseDNSServers() []string {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("useDNSServers"))
	return objc.ConvertSliceToStrings(rv)
}
func (e NEEvaluateConnectionRule) SetUseDNSServers(value []string) {
	objc.Send[struct{}](e.ID, objc.Sel("setUseDNSServers:"), objectivec.StringSliceToNSArray(value))
}
// An HTTP or HTTPS URL. If the rule matches the connection being established
// and the action is [NEEvaluateConnectionRuleActionConnectIfNeeded] and a
// request sent to this URL results in a response with an HTTP response code
// other than 200, then the VPN is started.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/probeURL
func (e NEEvaluateConnectionRule) ProbeURL() foundation.INSURL {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("probeURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (e NEEvaluateConnectionRule) SetProbeURL(value foundation.INSURL) {
	objc.Send[struct{}](e.ID, objc.Sel("setProbeURL:"), value)
}
// The action to take if the properties of the network connection being
// established match the rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/action
func (e NEEvaluateConnectionRule) Action() NEEvaluateConnectionRuleAction {
	rv := objc.Send[NEEvaluateConnectionRuleAction](e.ID, objc.Sel("action"))
	return NEEvaluateConnectionRuleAction(rv)
}
// An array of
//
// See: https://developer.apple.com/documentation/networkextension/neondemandruleevaluateconnection/connectionrules
func (e NEEvaluateConnectionRule) ConnectionRules() INEEvaluateConnectionRule {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("connectionRules"))
	return NEEvaluateConnectionRuleFromID(objc.ID(rv))
}
func (e NEEvaluateConnectionRule) SetConnectionRules(value INEEvaluateConnectionRule) {
	objc.Send[struct{}](e.ID, objc.Sel("setConnectionRules:"), value)
}


// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEOnDemandRuleEvaluateConnection] class.
var (
	_NEOnDemandRuleEvaluateConnectionClass     NEOnDemandRuleEvaluateConnectionClass
	_NEOnDemandRuleEvaluateConnectionClassOnce sync.Once
)

func getNEOnDemandRuleEvaluateConnectionClass() NEOnDemandRuleEvaluateConnectionClass {
	_NEOnDemandRuleEvaluateConnectionClassOnce.Do(func() {
		_NEOnDemandRuleEvaluateConnectionClass = NEOnDemandRuleEvaluateConnectionClass{class: objc.GetClass("NEOnDemandRuleEvaluateConnection")}
	})
	return _NEOnDemandRuleEvaluateConnectionClass
}

// GetNEOnDemandRuleEvaluateConnectionClass returns the class object for NEOnDemandRuleEvaluateConnection.
func GetNEOnDemandRuleEvaluateConnectionClass() NEOnDemandRuleEvaluateConnectionClass {
	return getNEOnDemandRuleEvaluateConnectionClass()
}

type NEOnDemandRuleEvaluateConnectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEOnDemandRuleEvaluateConnectionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEOnDemandRuleEvaluateConnectionClass) Alloc() NEOnDemandRuleEvaluateConnection {
	rv := objc.Send[NEOnDemandRuleEvaluateConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A VPN On Demand rule that evaluate the app’s connection to determine
// whether to run its action.
//
// # Overview
//
// When rules of this class match, the properties of the network connection
// being established are matched against a set of connection rules. The action
// of the matched rule (if any) is used to determine whether or not the VPN
// will be started.
//
// # Accessing connection rules
//
//   - [NEOnDemandRuleEvaluateConnection.ConnectionRules]: An array of [NEEvaluateConnectionRule](<doc://com.apple.networkextension/documentation/NetworkExtension/NEEvaluateConnectionRule>) objects
//   - [NEOnDemandRuleEvaluateConnection.SetConnectionRules]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleEvaluateConnection
type NEOnDemandRuleEvaluateConnection struct {
	NEOnDemandRule
}

// NEOnDemandRuleEvaluateConnectionFromID constructs a [NEOnDemandRuleEvaluateConnection] from an objc.ID.
//
// A VPN On Demand rule that evaluate the app’s connection to determine
// whether to run its action.
func NEOnDemandRuleEvaluateConnectionFromID(id objc.ID) NEOnDemandRuleEvaluateConnection {
	return NEOnDemandRuleEvaluateConnection{NEOnDemandRule: NEOnDemandRuleFromID(id)}
}

// NOTE: NEOnDemandRuleEvaluateConnection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEOnDemandRuleEvaluateConnection] class.
//
// # Accessing connection rules
//
//   - [INEOnDemandRuleEvaluateConnection.ConnectionRules]: An array of [NEEvaluateConnectionRule](<doc://com.apple.networkextension/documentation/NetworkExtension/NEEvaluateConnectionRule>) objects
//   - [INEOnDemandRuleEvaluateConnection.SetConnectionRules]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleEvaluateConnection
type INEOnDemandRuleEvaluateConnection interface {
	INEOnDemandRule

	// Topic: Accessing connection rules

	// An array of [NEEvaluateConnectionRule](<doc://com.apple.networkextension/documentation/NetworkExtension/NEEvaluateConnectionRule>) objects
	ConnectionRules() []NEEvaluateConnectionRule
	SetConnectionRules(value []NEEvaluateConnectionRule)
}

// Init initializes the instance.
func (o NEOnDemandRuleEvaluateConnection) Init() NEOnDemandRuleEvaluateConnection {
	rv := objc.Send[NEOnDemandRuleEvaluateConnection](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NEOnDemandRuleEvaluateConnection) Autorelease() NEOnDemandRuleEvaluateConnection {
	rv := objc.Send[NEOnDemandRuleEvaluateConnection](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEOnDemandRuleEvaluateConnection creates a new NEOnDemandRuleEvaluateConnection instance.
func NewNEOnDemandRuleEvaluateConnection() NEOnDemandRuleEvaluateConnection {
	class := getNEOnDemandRuleEvaluateConnectionClass()
	rv := objc.Send[NEOnDemandRuleEvaluateConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array of [NEEvaluateConnectionRule] objects
//
// # Discussion
//
// Each [NEEvaluateConnectionRule] object defines a behavior to take for
// connections that match the domain of the rule. Each rule is evaluated in
// order against the properties of a network connection being established. An
// example configuration has two connection rules: a rule matching
// `myserver.ExampleXCUIElementTypeCom()` with the domain action
// [NEEvaluateConnectionRuleActionNeverConnect], followed by a rule matching
// `example.Com()` with the domain action
// [NEEvaluateConnectionRuleActionConnectIfNeeded]. This configuration would
// cause all connections to hostnames in `example.Com()` that do not resolve
// on the current network to trigger the VPN, except for
// `myserver.ExampleXCUIElementTypeCom()`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleEvaluateConnection/connectionRules
func (o NEOnDemandRuleEvaluateConnection) ConnectionRules() []NEEvaluateConnectionRule {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("connectionRules"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEEvaluateConnectionRule {
		return NEEvaluateConnectionRuleFromID(id)
	})
}
func (o NEOnDemandRuleEvaluateConnection) SetConnectionRules(value []NEEvaluateConnectionRule) {
	objc.Send[struct{}](o.ID, objc.Sel("setConnectionRules:"), objectivec.IObjectSliceToNSArray(value))
}

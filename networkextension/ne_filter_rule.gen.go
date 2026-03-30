// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEFilterRule] class.
var (
	_NEFilterRuleClass     NEFilterRuleClass
	_NEFilterRuleClassOnce sync.Once
)

func getNEFilterRuleClass() NEFilterRuleClass {
	_NEFilterRuleClassOnce.Do(func() {
		_NEFilterRuleClass = NEFilterRuleClass{class: objc.GetClass("NEFilterRule")}
	})
	return _NEFilterRuleClass
}

// GetNEFilterRuleClass returns the class object for NEFilterRule.
func GetNEFilterRuleClass() NEFilterRuleClass {
	return getNEFilterRuleClass()
}

type NEFilterRuleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterRuleClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterRuleClass) Alloc() NEFilterRule {
	rv := objc.Send[NEFilterRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A rule for filters that combines a rule to match network traffic and an
// action to take when the rule matches.
//
// # Creating a Filter Rule
//
//   - [NEFilterRule.InitWithNetworkRuleAction]: Creates a new filter rule from a network rule and an action to take when network traffic matches.
//
// # Inspecting Filter Rule Properties
//
//   - [NEFilterRule.NetworkRule]: The network rule that defines the network traffic characteristics that this filter rule matches.
//   - [NEFilterRule.Action]: The action to take when this rule matches network traffic.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule
type NEFilterRule struct {
	objectivec.Object
}

// NEFilterRuleFromID constructs a [NEFilterRule] from an objc.ID.
//
// A rule for filters that combines a rule to match network traffic and an
// action to take when the rule matches.
func NEFilterRuleFromID(id objc.ID) NEFilterRule {
	return NEFilterRule{objectivec.Object{ID: id}}
}

// NOTE: NEFilterRule adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterRule] class.
//
// # Creating a Filter Rule
//
//   - [INEFilterRule.InitWithNetworkRuleAction]: Creates a new filter rule from a network rule and an action to take when network traffic matches.
//
// # Inspecting Filter Rule Properties
//
//   - [INEFilterRule.NetworkRule]: The network rule that defines the network traffic characteristics that this filter rule matches.
//   - [INEFilterRule.Action]: The action to take when this rule matches network traffic.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule
type INEFilterRule interface {
	objectivec.IObject

	// Topic: Creating a Filter Rule

	// Creates a new filter rule from a network rule and an action to take when network traffic matches.
	InitWithNetworkRuleAction(networkRule INENetworkRule, action NEFilterAction) NEFilterRule

	// Topic: Inspecting Filter Rule Properties

	// The network rule that defines the network traffic characteristics that this filter rule matches.
	NetworkRule() INENetworkRule
	// The action to take when this rule matches network traffic.
	Action() NEFilterAction

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NEFilterRule) Init() NEFilterRule {
	rv := objc.Send[NEFilterRule](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterRule) Autorelease() NEFilterRule {
	rv := objc.Send[NEFilterRule](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterRule creates a new NEFilterRule instance.
func NewNEFilterRule() NEFilterRule {
	class := getNEFilterRuleClass()
	rv := objc.Send[NEFilterRule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new filter rule from a network rule and an action to take when
// network traffic matches.
//
// networkRule: An [NENetworkRule] object that defines the network traffic characteristics
// that this rule matches.
//
// action: The action to take when the network rule matches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule/init(networkRule:action:)
func NewFilterRuleWithNetworkRuleAction(networkRule INENetworkRule, action NEFilterAction) NEFilterRule {
	instance := getNEFilterRuleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetworkRule:action:"), networkRule, action)
	return NEFilterRuleFromID(rv)
}

// Creates a new filter rule from a network rule and an action to take when
// network traffic matches.
//
// networkRule: An [NENetworkRule] object that defines the network traffic characteristics
// that this rule matches.
//
// action: The action to take when the network rule matches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule/init(networkRule:action:)
func (f NEFilterRule) InitWithNetworkRuleAction(networkRule INENetworkRule, action NEFilterAction) NEFilterRule {
	rv := objc.Send[NEFilterRule](f.ID, objc.Sel("initWithNetworkRule:action:"), networkRule, action)
	return rv
}
func (f NEFilterRule) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The network rule that defines the network traffic characteristics that this
// filter rule matches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule/networkRule
func (f NEFilterRule) NetworkRule() INENetworkRule {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("networkRule"))
	return NENetworkRuleFromID(objc.ID(rv))
}

// The action to take when this rule matches network traffic.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule/action
func (f NEFilterRule) Action() NEFilterAction {
	rv := objc.Send[NEFilterAction](f.ID, objc.Sel("action"))
	return NEFilterAction(rv)
}

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEAppRule] class.
var (
	_NEAppRuleClass     NEAppRuleClass
	_NEAppRuleClassOnce sync.Once
)

func getNEAppRuleClass() NEAppRuleClass {
	_NEAppRuleClassOnce.Do(func() {
		_NEAppRuleClass = NEAppRuleClass{class: objc.GetClass("NEAppRule")}
	})
	return _NEAppRuleClass
}

// GetNEAppRuleClass returns the class object for NEAppRule.
func GetNEAppRuleClass() NEAppRuleClass {
	return getNEAppRuleClass()
}

type NEAppRuleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEAppRuleClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEAppRuleClass) Alloc() NEAppRule {
	rv := objc.Send[NEAppRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The identity of an app whose traffic is to be routed through the tunnel.
//
// # Initializing an app rule
//
//   - [NEAppRule.InitWithSigningIdentifierDesignatedRequirement]: Create an app rule that matches an app with a given signing identifier and a given designated requirement.
//
// # Accessing app rule properties
//
//   - [NEAppRule.MatchSigningIdentifier]: The signing identifier of the app that matches the rule.
//   - [NEAppRule.MatchDesignatedRequirement]: The designated requirement of the app that matches the rule.
//   - [NEAppRule.MatchPath]: The file system path of the app that matches the rule.
//   - [NEAppRule.SetMatchPath]
//   - [NEAppRule.MatchDomains]: The hostname domains that match the rule.
//   - [NEAppRule.SetMatchDomains]
//   - [NEAppRule.MatchTools]: An array of app rule objects that restrict the rule so it only matches network traffic generated from helper processes.
//   - [NEAppRule.SetMatchTools]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule
type NEAppRule struct {
	objectivec.Object
}

// NEAppRuleFromID constructs a [NEAppRule] from an objc.ID.
//
// The identity of an app whose traffic is to be routed through the tunnel.
func NEAppRuleFromID(id objc.ID) NEAppRule {
	return NEAppRule{objectivec.Object{ID: id}}
}

// NOTE: NEAppRule adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEAppRule] class.
//
// # Initializing an app rule
//
//   - [INEAppRule.InitWithSigningIdentifierDesignatedRequirement]: Create an app rule that matches an app with a given signing identifier and a given designated requirement.
//
// # Accessing app rule properties
//
//   - [INEAppRule.MatchSigningIdentifier]: The signing identifier of the app that matches the rule.
//   - [INEAppRule.MatchDesignatedRequirement]: The designated requirement of the app that matches the rule.
//   - [INEAppRule.MatchPath]: The file system path of the app that matches the rule.
//   - [INEAppRule.SetMatchPath]
//   - [INEAppRule.MatchDomains]: The hostname domains that match the rule.
//   - [INEAppRule.SetMatchDomains]
//   - [INEAppRule.MatchTools]: An array of app rule objects that restrict the rule so it only matches network traffic generated from helper processes.
//   - [INEAppRule.SetMatchTools]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule
type INEAppRule interface {
	objectivec.IObject

	// Topic: Initializing an app rule

	// Create an app rule that matches an app with a given signing identifier and a given designated requirement.
	InitWithSigningIdentifierDesignatedRequirement(signingIdentifier string, designatedRequirement string) NEAppRule

	// Topic: Accessing app rule properties

	// The signing identifier of the app that matches the rule.
	MatchSigningIdentifier() string
	// The designated requirement of the app that matches the rule.
	MatchDesignatedRequirement() string
	// The file system path of the app that matches the rule.
	MatchPath() string
	SetMatchPath(value string)
	// The hostname domains that match the rule.
	MatchDomains() foundation.INSArray
	SetMatchDomains(value foundation.INSArray)
	// An array of app rule objects that restrict the rule so it only matches network traffic generated from helper processes.
	MatchTools() []NEAppRule
	SetMatchTools(value []NEAppRule)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a NEAppRule) Init() NEAppRule {
	rv := objc.Send[NEAppRule](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NEAppRule) Autorelease() NEAppRule {
	rv := objc.Send[NEAppRule](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppRule creates a new NEAppRule instance.
func NewNEAppRule() NEAppRule {
	class := getNEAppRuleClass()
	rv := objc.Send[NEAppRule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Create an app rule that matches an app with a given signing identifier.
//
// signingIdentifier: The signing identifier of the app that matches the rule. For apps that are
// signed using Xcode, the app’s signing identifier is equivalent to the
// app’s bundle identifier.
//
// # Return Value
//
// A newly-initialized [NEAppRule] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/init(signingIdentifier:)
func NewAppRuleWithSigningIdentifier(signingIdentifier string) NEAppRule {
	instance := getNEAppRuleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSigningIdentifier:"), objc.String(signingIdentifier))
	return NEAppRuleFromID(rv)
}

// Create an app rule that matches an app with a given signing identifier and
// a given designated requirement.
//
// signingIdentifier: The signing identifier of the app that matches the rule. For apps that are
// signed using Xcode, the app’s signing identifier is equivalent to the
// app’s bundle identifier.
//
// designatedRequirement: The designated requirement of the app that matches the rule. The designated
// requirement for an app can be obtained using the `codesign` command-line
// developer tool.
//
// # Return Value
//
// A newly-initialized [NEAppRule] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/init(signingIdentifier:designatedRequirement:)
func NewAppRuleWithSigningIdentifierDesignatedRequirement(signingIdentifier string, designatedRequirement string) NEAppRule {
	instance := getNEAppRuleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSigningIdentifier:designatedRequirement:"), objc.String(signingIdentifier), objc.String(designatedRequirement))
	return NEAppRuleFromID(rv)
}

// Create an app rule that matches an app with a given signing identifier and
// a given designated requirement.
//
// signingIdentifier: The signing identifier of the app that matches the rule. For apps that are
// signed using Xcode, the app’s signing identifier is equivalent to the
// app’s bundle identifier.
//
// designatedRequirement: The designated requirement of the app that matches the rule. The designated
// requirement for an app can be obtained using the `codesign` command-line
// developer tool.
//
// # Return Value
//
// A newly-initialized [NEAppRule] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/init(signingIdentifier:designatedRequirement:)
func (a NEAppRule) InitWithSigningIdentifierDesignatedRequirement(signingIdentifier string, designatedRequirement string) NEAppRule {
	rv := objc.Send[NEAppRule](a.ID, objc.Sel("initWithSigningIdentifier:designatedRequirement:"), objc.String(signingIdentifier), objc.String(designatedRequirement))
	return rv
}
func (a NEAppRule) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The signing identifier of the app that matches the rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchSigningIdentifier
func (a NEAppRule) MatchSigningIdentifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("matchSigningIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// The designated requirement of the app that matches the rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDesignatedRequirement
func (a NEAppRule) MatchDesignatedRequirement() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("matchDesignatedRequirement"))
	return foundation.NSStringFromID(rv).String()
}

// The file system path of the app that matches the rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchPath
func (a NEAppRule) MatchPath() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("matchPath"))
	return foundation.NSStringFromID(rv).String()
}
func (a NEAppRule) SetMatchPath(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setMatchPath:"), objc.String(value))
}

// The hostname domains that match the rule.
//
// # Discussion
//
// If this property is set to a nonempty array, then only connections to
// destinations in the domains specified in the array will use the VPN.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDomains
func (a NEAppRule) MatchDomains() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("matchDomains"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a NEAppRule) SetMatchDomains(value foundation.INSArray) {
	objc.Send[struct{}](a.ID, objc.Sel("setMatchDomains:"), value)
}

// An array of app rule objects that restrict the rule so it only matches
// network traffic generated from helper processes.
//
// # Discussion
//
// Use this property to restrict this rule so it only matches network traffic
// that the matching app generates and all helper tool processes that the
// matching app spawns.
//
// For example, to match network traffic generated by the `curl` command line
// tool run from `Terminal.App()`, do the following:
//
// - Create an [NEAppRule] for `Terminal.App()`. - Set the app rule’s
// [MatchTools] property to an array that contains an [NEAppRule] for the
// `curl` command line tool.
//
// Set this property to `nil` (the default) to match all network traffic
// generated by the matching app and all helper tool processes spawned by the
// matching app.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchTools
func (a NEAppRule) MatchTools() []NEAppRule {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("matchTools"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEAppRule {
		return NEAppRuleFromID(id)
	})
}
func (a NEAppRule) SetMatchTools(value []NEAppRule) {
	objc.Send[struct{}](a.ID, objc.Sel("setMatchTools:"), objectivec.IObjectSliceToNSArray(value))
}

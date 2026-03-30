// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEFilterSettings] class.
var (
	_NEFilterSettingsClass     NEFilterSettingsClass
	_NEFilterSettingsClassOnce sync.Once
)

func getNEFilterSettingsClass() NEFilterSettingsClass {
	_NEFilterSettingsClassOnce.Do(func() {
		_NEFilterSettingsClass = NEFilterSettingsClass{class: objc.GetClass("NEFilterSettings")}
	})
	return _NEFilterSettingsClass
}

// GetNEFilterSettingsClass returns the class object for NEFilterSettings.
func GetNEFilterSettingsClass() NEFilterSettingsClass {
	return getNEFilterSettingsClass()
}

type NEFilterSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterSettingsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterSettingsClass) Alloc() NEFilterSettings {
	rv := objc.Send[NEFilterSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The rules and other settings that define the operation of a filter.
//
// # Overview
//
// [NEFilterDataProvider] instances use [NEFilterSettings] to communicate the
// desired settings for the filter to the framework. The framework takes care
// of applying the contained settings to the system.
//
// # Creating Filter Settings
//
//   - [NEFilterSettings.InitWithRulesDefaultAction]: Creates a new settings instance from an array of rules and a default action.
//
// # Inspecting Filter Settings
//
//   - [NEFilterSettings.Rules]: An ordered list of rules that define the filter’s operation.
//   - [NEFilterSettings.DefaultAction]: The default action to take for flows of network data that don’t match any of the specified rules.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSettings
type NEFilterSettings struct {
	objectivec.Object
}

// NEFilterSettingsFromID constructs a [NEFilterSettings] from an objc.ID.
//
// The rules and other settings that define the operation of a filter.
func NEFilterSettingsFromID(id objc.ID) NEFilterSettings {
	return NEFilterSettings{objectivec.Object{ID: id}}
}

// NOTE: NEFilterSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterSettings] class.
//
// # Creating Filter Settings
//
//   - [INEFilterSettings.InitWithRulesDefaultAction]: Creates a new settings instance from an array of rules and a default action.
//
// # Inspecting Filter Settings
//
//   - [INEFilterSettings.Rules]: An ordered list of rules that define the filter’s operation.
//   - [INEFilterSettings.DefaultAction]: The default action to take for flows of network data that don’t match any of the specified rules.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSettings
type INEFilterSettings interface {
	objectivec.IObject

	// Topic: Creating Filter Settings

	// Creates a new settings instance from an array of rules and a default action.
	InitWithRulesDefaultAction(rules []NEFilterRule, defaultAction NEFilterAction) NEFilterSettings

	// Topic: Inspecting Filter Settings

	// An ordered list of rules that define the filter’s operation.
	Rules() []NEFilterRule
	// The default action to take for flows of network data that don’t match any of the specified rules.
	DefaultAction() NEFilterAction

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NEFilterSettings) Init() NEFilterSettings {
	rv := objc.Send[NEFilterSettings](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterSettings) Autorelease() NEFilterSettings {
	rv := objc.Send[NEFilterSettings](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterSettings creates a new NEFilterSettings instance.
func NewNEFilterSettings() NEFilterSettings {
	class := getNEFilterSettingsClass()
	rv := objc.Send[NEFilterSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new settings instance from an array of rules and a default
// action.
//
// rules: An array containing an ordered list of [NEFilterRule] objects. The maximum
// number of rules that this array can contain is 1000.
//
// defaultAction: The [NEFilterAction] to take for flows of network data that don’t match
// any of the specified rules. The default `defaultAction` is
// [NEFilterActionFilterData]. If `defaultAction` is [NEFilterActionAllow] or
// [NEFilterActionDrop], then the `rules` array must contain at least one
// [NEFilterRule].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSettings/init(rules:defaultAction:)
//
// [NEFilterAction]: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction
func NewFilterSettingsWithRulesDefaultAction(rules []NEFilterRule, defaultAction NEFilterAction) NEFilterSettings {
	instance := getNEFilterSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRules:defaultAction:"), objectivec.IObjectSliceToNSArray(rules), defaultAction)
	return NEFilterSettingsFromID(rv)
}

// Creates a new settings instance from an array of rules and a default
// action.
//
// rules: An array containing an ordered list of [NEFilterRule] objects. The maximum
// number of rules that this array can contain is 1000.
//
// defaultAction: The [NEFilterAction] to take for flows of network data that don’t match
// any of the specified rules. The default `defaultAction` is
// [NEFilterActionFilterData]. If `defaultAction` is [NEFilterActionAllow] or
// [NEFilterActionDrop], then the `rules` array must contain at least one
// [NEFilterRule].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSettings/init(rules:defaultAction:)
//
// [NEFilterAction]: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction
func (f NEFilterSettings) InitWithRulesDefaultAction(rules []NEFilterRule, defaultAction NEFilterAction) NEFilterSettings {
	rv := objc.Send[NEFilterSettings](f.ID, objc.Sel("initWithRules:defaultAction:"), objectivec.IObjectSliceToNSArray(rules), defaultAction)
	return rv
}
func (f NEFilterSettings) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// An ordered list of rules that define the filter’s operation.
//
// # Discussion
//
// After applying the [NEFilterSettings], the system compares each network
// flow against these rules in order, and acts on the rule of the first
// [NEFilterAction] that matches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSettings/rules
//
// [NEFilterAction]: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction
func (f NEFilterSettings) Rules() []NEFilterRule {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("rules"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEFilterRule {
		return NEFilterRuleFromID(id)
	})
}

// The default action to take for flows of network data that don’t match any
// of the specified rules.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterSettings/defaultAction
func (f NEFilterSettings) DefaultAction() NEFilterAction {
	rv := objc.Send[NEFilterAction](f.ID, objc.Sel("defaultAction"))
	return NEFilterAction(rv)
}

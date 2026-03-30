// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEDNSSettingsManager] class.
var (
	_NEDNSSettingsManagerClass     NEDNSSettingsManagerClass
	_NEDNSSettingsManagerClassOnce sync.Once
)

func getNEDNSSettingsManagerClass() NEDNSSettingsManagerClass {
	_NEDNSSettingsManagerClassOnce.Do(func() {
		_NEDNSSettingsManagerClass = NEDNSSettingsManagerClass{class: objc.GetClass("NEDNSSettingsManager")}
	})
	return _NEDNSSettingsManagerClass
}

// GetNEDNSSettingsManagerClass returns the class object for NEDNSSettingsManager.
func GetNEDNSSettingsManagerClass() NEDNSSettingsManagerClass {
	return getNEDNSSettingsManagerClass()
}

type NEDNSSettingsManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEDNSSettingsManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEDNSSettingsManagerClass) Alloc() NEDNSSettingsManager {
	rv := objc.Send[NEDNSSettingsManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object you use to create and manage a DNS settings configuration.
//
// # Overview
//
// When your app starts up, access the shared instance of the DNS settings
// manager, and load existing settings from the preferences using
// [NEDNSSettingsManager.LoadFromPreferencesWithCompletionHandler]. You can define your DNS server
// configuration, and persist it by calling
// [NEDNSSettingsManager.SaveToPreferencesWithCompletionHandler].
//
// In order to use your DNS settings, the user needs to enable it in the
// Settings app on iOS or in System Preferences on macOS.
//
// # Managing DNS configurations
//
//   - [NEDNSSettingsManager.LoadFromPreferencesWithCompletionHandler]: Load your DNS settings configuration from the system networking preferences.
//   - [NEDNSSettingsManager.SaveToPreferencesWithCompletionHandler]: Save your DNS settings configuration to the system networking preferences.
//   - [NEDNSSettingsManager.RemoveFromPreferencesWithCompletionHandler]: Remove your DNS settings configuration from the system networking preferences.
//
// # Accessing DNS configuration properties
//
//   - [NEDNSSettingsManager.Enabled]: A Boolean you use to query the enabled state of the DNS settings configuration.
//   - [NEDNSSettingsManager.DnsSettings]: An object that contains the configuration settings for a DNS server.
//   - [NEDNSSettingsManager.SetDnsSettings]
//   - [NEDNSSettingsManager.LocalizedDescription]: A string that contains the display name of the DNS settings configuration.
//   - [NEDNSSettingsManager.SetLocalizedDescription]
//   - [NEDNSSettingsManager.OnDemandRules]: A list of ordered rules that defines the networks on which the DNS settings will apply.
//   - [NEDNSSettingsManager.SetOnDemandRules]
//
// # Handling errors
//
//   - [NEDNSSettingsManager.NEDNSSettingsErrorDomain]: The domain for errors resulting from calls to the DNS settings manager.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager
type NEDNSSettingsManager struct {
	objectivec.Object
}

// NEDNSSettingsManagerFromID constructs a [NEDNSSettingsManager] from an objc.ID.
//
// An object you use to create and manage a DNS settings configuration.
func NEDNSSettingsManagerFromID(id objc.ID) NEDNSSettingsManager {
	return NEDNSSettingsManager{objectivec.Object{ID: id}}
}

// NOTE: NEDNSSettingsManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEDNSSettingsManager] class.
//
// # Managing DNS configurations
//
//   - [INEDNSSettingsManager.LoadFromPreferencesWithCompletionHandler]: Load your DNS settings configuration from the system networking preferences.
//   - [INEDNSSettingsManager.SaveToPreferencesWithCompletionHandler]: Save your DNS settings configuration to the system networking preferences.
//   - [INEDNSSettingsManager.RemoveFromPreferencesWithCompletionHandler]: Remove your DNS settings configuration from the system networking preferences.
//
// # Accessing DNS configuration properties
//
//   - [INEDNSSettingsManager.Enabled]: A Boolean you use to query the enabled state of the DNS settings configuration.
//   - [INEDNSSettingsManager.DnsSettings]: An object that contains the configuration settings for a DNS server.
//   - [INEDNSSettingsManager.SetDnsSettings]
//   - [INEDNSSettingsManager.LocalizedDescription]: A string that contains the display name of the DNS settings configuration.
//   - [INEDNSSettingsManager.SetLocalizedDescription]
//   - [INEDNSSettingsManager.OnDemandRules]: A list of ordered rules that defines the networks on which the DNS settings will apply.
//   - [INEDNSSettingsManager.SetOnDemandRules]
//
// # Handling errors
//
//   - [INEDNSSettingsManager.NEDNSSettingsErrorDomain]: The domain for errors resulting from calls to the DNS settings manager.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager
type INEDNSSettingsManager interface {
	objectivec.IObject

	// Topic: Managing DNS configurations

	// Load your DNS settings configuration from the system networking preferences.
	LoadFromPreferencesWithCompletionHandler(completionHandler ErrorHandler)
	// Save your DNS settings configuration to the system networking preferences.
	SaveToPreferencesWithCompletionHandler(completionHandler ErrorHandler)
	// Remove your DNS settings configuration from the system networking preferences.
	RemoveFromPreferencesWithCompletionHandler(completionHandler ErrorHandler)

	// Topic: Accessing DNS configuration properties

	// A Boolean you use to query the enabled state of the DNS settings configuration.
	Enabled() bool
	// An object that contains the configuration settings for a DNS server.
	DnsSettings() INEDNSSettings
	SetDnsSettings(value INEDNSSettings)
	// A string that contains the display name of the DNS settings configuration.
	LocalizedDescription() string
	SetLocalizedDescription(value string)
	// A list of ordered rules that defines the networks on which the DNS settings will apply.
	OnDemandRules() []NEOnDemandRule
	SetOnDemandRules(value []NEOnDemandRule)

	// Topic: Handling errors

	// The domain for errors resulting from calls to the DNS settings manager.
	NEDNSSettingsErrorDomain() string
}

// Init initializes the instance.
func (d NEDNSSettingsManager) Init() NEDNSSettingsManager {
	rv := objc.Send[NEDNSSettingsManager](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NEDNSSettingsManager) Autorelease() NEDNSSettingsManager {
	rv := objc.Send[NEDNSSettingsManager](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSSettingsManager creates a new NEDNSSettingsManager instance.
func NewNEDNSSettingsManager() NEDNSSettingsManager {
	class := getNEDNSSettingsManagerClass()
	rv := objc.Send[NEDNSSettingsManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Load your DNS settings configuration from the system networking
// preferences.
//
// completionHandler: A block that takes an [NSError] object. This block runs on your
// application’s main thread after the load operation is complete. If an
// error occurs while loading the configuration, the block returns an
// [NSError] object.
//
// # Discussion
//
// You must call this method at least once before calling
// [SaveToPreferencesWithCompletionHandler] for the first time after your app
// launches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/loadFromPreferences(completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (d NEDNSSettingsManager) LoadFromPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), _block0)
}

// Save your DNS settings configuration to the system networking preferences.
//
// completionHandler: An optional block that takes an [NSError] object. If specified, this block
// runs on your application’s main thread after the save operation
// completes. If an error occurs while saving the configuration, the block
// returns an [NSError] object.
//
// # Discussion
//
// You must call [LoadFromPreferencesWithCompletionHandler] at least once
// before calling this method the first time after your app launches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/saveToPreferences(completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (d NEDNSSettingsManager) SaveToPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), _block0)
}

// Remove your DNS settings configuration from the system networking
// preferences.
//
// completionHandler: An optional block that takes an [NSError] object. If specified, this block
// runs on your application’s main thread after your configuration is
// removed. If an error occurs while removing the configuration, the block
// returns an [NSError] object.
//
// # Discussion
//
// After you remove your configuration, the [NEDNSSettingsManager] object
// still contains the configuration parameters. Calling
// [LoadFromPreferencesWithCompletionHandler] clears out the configuration
// parameters from the [NEDNSSettingsManager] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/removeFromPreferences(completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (d NEDNSSettingsManager) RemoveFromPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), _block0)
}

// Access the single instance of a DNS settings manager.
//
// # Return Value
//
// The DNS settings manager instance for the calling application.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/shared()
func (_NEDNSSettingsManagerClass NEDNSSettingsManagerClass) SharedManager() NEDNSSettingsManager {
	rv := objc.Send[objc.ID](objc.ID(_NEDNSSettingsManagerClass.class), objc.Sel("sharedManager"))
	return NEDNSSettingsManagerFromID(rv)
}

// A Boolean you use to query the enabled state of the DNS settings
// configuration.
//
// # Discussion
//
// A user must enable your DNS settings configuration in order to apply it to
// the system. By default, configurations are disabled until the user enables
// the configuration in the Settings app on iOS or in System Preferences on
// macOS.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/isEnabled
func (d NEDNSSettingsManager) Enabled() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isEnabled"))
	return rv
}

// An object that contains the configuration settings for a DNS server.
//
// # Discussion
//
// This property can be set to either an [NEDNSOverHTTPSSettings] object or an
// [NEDNSOverTLSSettings] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/dnsSettings
func (d NEDNSSettingsManager) DnsSettings() INEDNSSettings {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dnsSettings"))
	return NEDNSSettingsFromID(objc.ID(rv))
}
func (d NEDNSSettingsManager) SetDnsSettings(value INEDNSSettings) {
	objc.Send[struct{}](d.ID, objc.Sel("setDnsSettings:"), value)
}

// A string that contains the display name of the DNS settings configuration.
//
// # Discussion
//
// This string is used as the display name of the DNS settings configuration
// in the system’s settings UI. If this property is set to `nil` at the time
// that the configuration is created, it is automatically set to the display
// name of the calling app.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/localizedDescription
func (d NEDNSSettingsManager) LocalizedDescription() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("localizedDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (d NEDNSSettingsManager) SetLocalizedDescription(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setLocalizedDescription:"), objc.String(value))
}

// A list of ordered rules that defines the networks on which the DNS settings
// will apply.
//
// # Discussion
//
// An On Demand rule with the action [NEOnDemandRuleActionConnect] defines a
// network on which the DNS settings apply. An On Demand rule with the action
// [NEOnDemandRuleActionDisconnect] causes DNS settings to not apply. An On
// Demand rule with the action of [NEOnDemandRuleActionEvaluateConnection] can
// be used to enable the DNS settings on a network with excluded domains, as
// specified using a [NEEvaluateConnectionRuleActionNeverConnect] rule.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/onDemandRules
func (d NEDNSSettingsManager) OnDemandRules() []NEOnDemandRule {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("onDemandRules"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEOnDemandRule {
		return NEOnDemandRuleFromID(id)
	})
}
func (d NEDNSSettingsManager) SetOnDemandRules(value []NEOnDemandRule) {
	objc.Send[struct{}](d.ID, objc.Sel("setOnDemandRules:"), objectivec.IObjectSliceToNSArray(value))
}

// The domain for errors resulting from calls to the DNS settings manager.
//
// See: https://developer.apple.com/documentation/networkextension/nednssettingserrordomain
func (d NEDNSSettingsManager) NEDNSSettingsErrorDomain() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("NEDNSSettingsErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// LoadFromPreferences is a synchronous wrapper around [NEDNSSettingsManager.LoadFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NEDNSSettingsManager) LoadFromPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	d.LoadFromPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SaveToPreferences is a synchronous wrapper around [NEDNSSettingsManager.SaveToPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NEDNSSettingsManager) SaveToPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	d.SaveToPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RemoveFromPreferences is a synchronous wrapper around [NEDNSSettingsManager.RemoveFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NEDNSSettingsManager) RemoveFromPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	d.RemoveFromPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

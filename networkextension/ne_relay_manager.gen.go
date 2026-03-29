// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NERelayManager] class.
var (
	_NERelayManagerClass     NERelayManagerClass
	_NERelayManagerClassOnce sync.Once
)

func getNERelayManagerClass() NERelayManagerClass {
	_NERelayManagerClassOnce.Do(func() {
		_NERelayManagerClass = NERelayManagerClass{class: objc.GetClass("NERelayManager")}
	})
	return _NERelayManagerClass
}

// GetNERelayManagerClass returns the class object for NERelayManager.
func GetNERelayManagerClass() NERelayManagerClass {
	return getNERelayManagerClass()
}

type NERelayManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NERelayManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NERelayManagerClass) Alloc() NERelayManager {
	rv := objc.Send[NERelayManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object you use to create and manage a network relay configuration.
//
// # Overview
// 
// When your app starts up, access the shared instance of the relay manager,
// and load existing settings from the preferences using
// [NERelayManager.LoadFromPreferencesWithCompletionHandler]. You can define your relay
// server configuration, and persist it by calling
// [NERelayManager.SaveToPreferencesWithCompletionHandler].
//
// # Managing relay configurations
//
//   - [NERelayManager.LoadFromPreferencesWithCompletionHandler]: Load your relay configuration from the system networking preferences.
//   - [NERelayManager.SaveToPreferencesWithCompletionHandler]: Save your relay configuration to the system networking preferences.
//   - [NERelayManager.RemoveFromPreferencesWithCompletionHandler]: Remove your relay configuration from the system networking preferences.
//
// # Accessing relay configuration properties
//
//   - [NERelayManager.Enabled]: A Boolean used to toggle the enabled state of the relay configuration.
//   - [NERelayManager.SetEnabled]
//   - [NERelayManager.Relays]: An array of one or two relay server configurations. If multiple relays are configured, application traffic routes through both of them in the order they appear in the array.
//   - [NERelayManager.SetRelays]
//   - [NERelayManager.MatchDomains]: A list of domain strings used to determine which connections will use the relay configuration contained in this object.
//   - [NERelayManager.SetMatchDomains]
//   - [NERelayManager.ExcludedDomains]: A list of domain strings used to determine which connections won’t use the relay configuration contained in this object.
//   - [NERelayManager.SetExcludedDomains]
//   - [NERelayManager.LocalizedDescription]: A string that contains the display name of the relay configuration.
//   - [NERelayManager.SetLocalizedDescription]
//   - [NERelayManager.OnDemandRules]: An array of rules you use to determine which networks the relay uses.
//   - [NERelayManager.SetOnDemandRules]
//
// # Handling errors
//
//   - [NERelayManager.NERelayErrorDomain]: The domain for errors resulting from calls to the relay manager.
//
// # Instance Properties
//
//   - [NERelayManager.ExcludedFQDNs]
//   - [NERelayManager.SetExcludedFQDNs]
//   - [NERelayManager.AllowDNSFailover]
//   - [NERelayManager.SetAllowDNSFailover]
//   - [NERelayManager.UIToggleEnabled]
//   - [NERelayManager.SetUIToggleEnabled]
//   - [NERelayManager.MatchFQDNs]
//   - [NERelayManager.SetMatchFQDNs]
//
// # Instance Methods
//
//   - [NERelayManager.GetLastClientErrorsCompletionHandler]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager
type NERelayManager struct {
	objectivec.Object
}

// NERelayManagerFromID constructs a [NERelayManager] from an objc.ID.
//
// An object you use to create and manage a network relay configuration.
func NERelayManagerFromID(id objc.ID) NERelayManager {
	return NERelayManager{objectivec.Object{ID: id}}
}
// NOTE: NERelayManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NERelayManager] class.
//
// # Managing relay configurations
//
//   - [INERelayManager.LoadFromPreferencesWithCompletionHandler]: Load your relay configuration from the system networking preferences.
//   - [INERelayManager.SaveToPreferencesWithCompletionHandler]: Save your relay configuration to the system networking preferences.
//   - [INERelayManager.RemoveFromPreferencesWithCompletionHandler]: Remove your relay configuration from the system networking preferences.
//
// # Accessing relay configuration properties
//
//   - [INERelayManager.Enabled]: A Boolean used to toggle the enabled state of the relay configuration.
//   - [INERelayManager.SetEnabled]
//   - [INERelayManager.Relays]: An array of one or two relay server configurations. If multiple relays are configured, application traffic routes through both of them in the order they appear in the array.
//   - [INERelayManager.SetRelays]
//   - [INERelayManager.MatchDomains]: A list of domain strings used to determine which connections will use the relay configuration contained in this object.
//   - [INERelayManager.SetMatchDomains]
//   - [INERelayManager.ExcludedDomains]: A list of domain strings used to determine which connections won’t use the relay configuration contained in this object.
//   - [INERelayManager.SetExcludedDomains]
//   - [INERelayManager.LocalizedDescription]: A string that contains the display name of the relay configuration.
//   - [INERelayManager.SetLocalizedDescription]
//   - [INERelayManager.OnDemandRules]: An array of rules you use to determine which networks the relay uses.
//   - [INERelayManager.SetOnDemandRules]
//
// # Handling errors
//
//   - [INERelayManager.NERelayErrorDomain]: The domain for errors resulting from calls to the relay manager.
//
// # Instance Properties
//
//   - [INERelayManager.ExcludedFQDNs]
//   - [INERelayManager.SetExcludedFQDNs]
//   - [INERelayManager.AllowDNSFailover]
//   - [INERelayManager.SetAllowDNSFailover]
//   - [INERelayManager.UIToggleEnabled]
//   - [INERelayManager.SetUIToggleEnabled]
//   - [INERelayManager.MatchFQDNs]
//   - [INERelayManager.SetMatchFQDNs]
//
// # Instance Methods
//
//   - [INERelayManager.GetLastClientErrorsCompletionHandler]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager
type INERelayManager interface {
	objectivec.IObject

	// Topic: Managing relay configurations

	// Load your relay configuration from the system networking preferences.
	LoadFromPreferencesWithCompletionHandler(completionHandler ErrorHandler)
	// Save your relay configuration to the system networking preferences.
	SaveToPreferencesWithCompletionHandler(completionHandler ErrorHandler)
	// Remove your relay configuration from the system networking preferences.
	RemoveFromPreferencesWithCompletionHandler(completionHandler ErrorHandler)

	// Topic: Accessing relay configuration properties

	// A Boolean used to toggle the enabled state of the relay configuration.
	Enabled() bool
	SetEnabled(value bool)
	// An array of one or two relay server configurations. If multiple relays are configured, application traffic routes through both of them in the order they appear in the array.
	Relays() []NERelay
	SetRelays(value []NERelay)
	// A list of domain strings used to determine which connections will use the relay configuration contained in this object.
	MatchDomains() []string
	SetMatchDomains(value []string)
	// A list of domain strings used to determine which connections won’t use the relay configuration contained in this object.
	ExcludedDomains() []string
	SetExcludedDomains(value []string)
	// A string that contains the display name of the relay configuration.
	LocalizedDescription() string
	SetLocalizedDescription(value string)
	// An array of rules you use to determine which networks the relay uses.
	OnDemandRules() []NEOnDemandRule
	SetOnDemandRules(value []NEOnDemandRule)

	// Topic: Handling errors

	// The domain for errors resulting from calls to the relay manager.
	NERelayErrorDomain() string

	// Topic: Instance Properties

	ExcludedFQDNs() []string
	SetExcludedFQDNs(value []string)
	AllowDNSFailover() bool
	SetAllowDNSFailover(value bool)
	UIToggleEnabled() bool
	SetUIToggleEnabled(value bool)
	MatchFQDNs() []string
	SetMatchFQDNs(value []string)

	// Topic: Instance Methods

	GetLastClientErrorsCompletionHandler(seconds float64, completionHandler ErrorHandler)
}

// Init initializes the instance.
func (r NERelayManager) Init() NERelayManager {
	rv := objc.Send[NERelayManager](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NERelayManager) Autorelease() NERelayManager {
	rv := objc.Send[NERelayManager](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNERelayManager creates a new NERelayManager instance.
func NewNERelayManager() NERelayManager {
	class := getNERelayManagerClass()
	rv := objc.Send[NERelayManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Load your relay configuration from the system networking preferences.
//
// completionHandler: A block that takes an [NSError] object. This block runs on your
// application’s main thread after the load operation is complete. If an
// error occurs while loading the configuration, the block returns an
// [NSError] object.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// You must call this method at least once before calling
// [SaveToPreferencesWithCompletionHandler] for the first time after your app
// launches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/loadFromPreferences(completionHandler:)
func (r NERelayManager) LoadFromPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](r.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), _block0)
}
// Save your relay configuration to the system networking preferences.
//
// completionHandler: An optional block that takes an [NSError] object. If specified, this block
// runs on your application’s main thread after the save operation
// completes. If an error occurs while saving the configuration, the block
// returns an [NSError] object.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// You must call [LoadFromPreferencesWithCompletionHandler] at least once
// before calling this method the first time after your app launches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/saveToPreferences(completionHandler:)
func (r NERelayManager) SaveToPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](r.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), _block0)
}
// Remove your relay configuration from the system networking preferences.
//
// completionHandler: An optional block that takes an [NSError] object. If specified, this block
// runs on your application’s main thread after your configuration is
// removed. If an error occurs while removing the configuration, the block
// returns an [NSError] object.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// After you remove your configuration, the [NERelayManager] object still
// contains the configuration parameters. Calling
// [LoadFromPreferencesWithCompletionHandler] clears out the configuration
// parameters from the [NERelayManager] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/removeFromPreferences(completionHandler:)
func (r NERelayManager) RemoveFromPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](r.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/getLastClientErrors(_:completionHandler:)
func (r NERelayManager) GetLastClientErrorsCompletionHandler(seconds float64, completionHandler ErrorHandler) {
_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](r.ID, objc.Sel("getLastClientErrors:completionHandler:"), seconds, _block1)
}

// Access the single instance of a network relay manager.
//
// # Return Value
// 
// The network relay manager instance for the calling application.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/shared()
func (_NERelayManagerClass NERelayManagerClass) SharedManager() NERelayManager {
	rv := objc.Send[objc.ID](objc.ID(_NERelayManagerClass.class), objc.Sel("sharedManager"))
	return NERelayManagerFromID(rv)
}

// A Boolean used to toggle the enabled state of the relay configuration.
//
// # Discussion
// 
// A relay configuration must be enabled before it can be used to proxy
// application traffic.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/isEnabled
func (r NERelayManager) Enabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isEnabled"))
	return rv
}
func (r NERelayManager) SetEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setEnabled:"), value)
}
// An array of one or two relay server configurations. If multiple relays are
// configured, application traffic routes through both of them in the order
// they appear in the array.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/relays
func (r NERelayManager) Relays() []NERelay {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("relays"))
	return objc.ConvertSlice(rv, func(id objc.ID) NERelay {
		return NERelayFromID(id)
	})
}
func (r NERelayManager) SetRelays(value []NERelay) {
	objc.Send[struct{}](r.ID, objc.Sel("setRelays:"), objectivec.IObjectSliceToNSArray(value))
}
// A list of domain strings used to determine which connections will use the
// relay configuration contained in this object.
//
// # Discussion
// 
// This property is used to create a “split DNS” configuration, where only
// hosts in certain domains route through the relays.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/matchDomains
func (r NERelayManager) MatchDomains() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("matchDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (r NERelayManager) SetMatchDomains(value []string) {
	objc.Send[struct{}](r.ID, objc.Sel("setMatchDomains:"), objectivec.StringSliceToNSArray(value))
}
// A list of domain strings used to determine which connections won’t use
// the relay configuration contained in this object.
//
// # Discussion
// 
// Excluded domains take precedence over domains listed in [MatchDomains].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/excludedDomains
func (r NERelayManager) ExcludedDomains() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("excludedDomains"))
	return objc.ConvertSliceToStrings(rv)
}
func (r NERelayManager) SetExcludedDomains(value []string) {
	objc.Send[struct{}](r.ID, objc.Sel("setExcludedDomains:"), objectivec.StringSliceToNSArray(value))
}
// A string that contains the display name of the relay configuration.
//
// # Discussion
// 
// This string is used as the display name of the relay configuration in the
// system’s settings UI. If this property is set to `nil` at the time that
// the configuration is created, it is automatically set to the display name
// of the calling app.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/localizedDescription
func (r NERelayManager) LocalizedDescription() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("localizedDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (r NERelayManager) SetLocalizedDescription(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setLocalizedDescription:"), objc.String(value))
}
// An array of rules you use to determine which networks the relay uses.
//
// # Discussion
// 
// If this value is `nil`, the associated relay always applies. If non-`nil`,
// the array describes the networks to which the relay applies.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/onDemandRules
func (r NERelayManager) OnDemandRules() []NEOnDemandRule {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("onDemandRules"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEOnDemandRule {
		return NEOnDemandRuleFromID(id)
	})
}
func (r NERelayManager) SetOnDemandRules(value []NEOnDemandRule) {
	objc.Send[struct{}](r.ID, objc.Sel("setOnDemandRules:"), objectivec.IObjectSliceToNSArray(value))
}
// The domain for errors resulting from calls to the relay manager.
//
// See: https://developer.apple.com/documentation/networkextension/nerelayerrordomain
func (r NERelayManager) NERelayErrorDomain() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("NERelayErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}
//
// # Discussion
// 
// An array of strings containing Fully Qualified Domain Names (FQDNs). If the
// destination host matches one of these strings then the relay will not be
// used. An excluded FQDN takes priority over the matchDomain property. This
// means the relay will not be used if the hostname matches an FQDN in this
// array even if the matchDomains contains a domain that would have been
// considered a match.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/excludedFQDNs
func (r NERelayManager) ExcludedFQDNs() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("excludedFQDNs"))
	return objc.ConvertSliceToStrings(rv)
}
func (r NERelayManager) SetExcludedFQDNs(value []string) {
	objc.Send[struct{}](r.ID, objc.Sel("setExcludedFQDNs:"), objectivec.StringSliceToNSArray(value))
}
//
// # Discussion
// 
// Determines if DNS queries that fail over relay can fallback to default DNS
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/isDNSFailoverAllowed
func (r NERelayManager) AllowDNSFailover() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isDNSFailoverAllowed"))
	return rv
}
func (r NERelayManager) SetAllowDNSFailover(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setAllowDNSFailover:"), value)
}
//
// # Discussion
// 
// Determines if the user will have the ability to enable and disable the
// relay
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/isUIToggleEnabled
func (r NERelayManager) UIToggleEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isUIToggleEnabled"))
	return rv
}
func (r NERelayManager) SetUIToggleEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setUIToggleEnabled:"), value)
}
//
// # Discussion
// 
// An array of strings containing Fully Qualified Domain Names (FQDNs). If
// this property is non-nil, the relay will be used to access the specified
// hosts. If this and the matchDomains property is nil, the relay will be used
// for all domains.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/matchFQDNs
func (r NERelayManager) MatchFQDNs() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("matchFQDNs"))
	return objc.ConvertSliceToStrings(rv)
}
func (r NERelayManager) SetMatchFQDNs(value []string) {
	objc.Send[struct{}](r.ID, objc.Sel("setMatchFQDNs:"), objectivec.StringSliceToNSArray(value))
}

// LoadFromPreferences is a synchronous wrapper around [NERelayManager.LoadFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (r NERelayManager) LoadFromPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	r.LoadFromPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SaveToPreferences is a synchronous wrapper around [NERelayManager.SaveToPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (r NERelayManager) SaveToPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	r.SaveToPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RemoveFromPreferences is a synchronous wrapper around [NERelayManager.RemoveFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (r NERelayManager) RemoveFromPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	r.RemoveFromPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetLastClientErrors is a synchronous wrapper around [NERelayManager.GetLastClientErrorsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (r NERelayManager) GetLastClientErrors(ctx context.Context, seconds float64) error {
	done := make(chan error, 1)
	r.GetLastClientErrorsCompletionHandler(seconds, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}


// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEVPNManager] class.
var (
	_NEVPNManagerClass     NEVPNManagerClass
	_NEVPNManagerClassOnce sync.Once
)

func getNEVPNManagerClass() NEVPNManagerClass {
	_NEVPNManagerClassOnce.Do(func() {
		_NEVPNManagerClass = NEVPNManagerClass{class: objc.GetClass("NEVPNManager")}
	})
	return _NEVPNManagerClass
}

// GetNEVPNManagerClass returns the class object for NEVPNManager.
func GetNEVPNManagerClass() NEVPNManagerClass {
	return getNEVPNManagerClass()
}

type NEVPNManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEVPNManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEVPNManagerClass) Alloc() NEVPNManager {
	rv := objc.Send[NEVPNManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object to create and manage a Personal VPN configuration.
//
// # Overview
// 
// The [NEVPNManager] API gives apps the ability to create and manage a
// Personal VPN configuration on iOS and macOS. Personal VPN configurations
// are typically used to provide a service to users that protects their
// Internet browsing activity on insecure networks such as public Wi-Fi
// networks.
//
// # Managing VPN configurations
//
//   - [NEVPNManager.LoadFromPreferencesWithCompletionHandler]: Load the VPN configuration from the Network Extension preferences.
//   - [NEVPNManager.SaveToPreferencesWithCompletionHandler]: Save the VPN configuration in the Network Extension preferences.
//   - [NEVPNManager.SetAuthorization]
//   - [NEVPNManager.RemoveFromPreferencesWithCompletionHandler]: Remove the VPN configuration from the Network Extension preferences.
//
// # Accessing VPN configuration properties
//
//   - [NEVPNManager.Enabled]: A Boolean used to toggle the enabled state of the VPN configuration.
//   - [NEVPNManager.SetEnabled]
//   - [NEVPNManager.ProtocolConfiguration]: An [NEVPNProtocol](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNProtocol>) object containing the configuration settings of the VPN tunneling protocol.
//   - [NEVPNManager.SetProtocolConfiguration]
//   - [NEVPNManager.LocalizedDescription]: A string containing the display name of the VPN configuration.
//   - [NEVPNManager.SetLocalizedDescription]
//   - [NEVPNManager.OnDemandEnabled]: A Boolean used to toggle the Connect On Demand capability.
//   - [NEVPNManager.SetOnDemandEnabled]
//   - [NEVPNManager.OnDemandRules]: An ordered list of Connect On Demand rules.
//   - [NEVPNManager.SetOnDemandRules]
//
// # Connecting and disconnecting VPN
//
//   - [NEVPNManager.Connection]: An [NEVPNConnection](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNConnection>) object that is used to control the VPN tunnel specified by the VPN configuration.
//
// # Errors
//
//   - [NEVPNManager.NEVPNErrorDomain]
//
// # Notifications
//
//   - [NEVPNManager.NEVPNConfigurationChange]: Posted after the VPN configuration stored in the Network Extension preferences changes.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager
type NEVPNManager struct {
	objectivec.Object
}

// NEVPNManagerFromID constructs a [NEVPNManager] from an objc.ID.
//
// An object to create and manage a Personal VPN configuration.
func NEVPNManagerFromID(id objc.ID) NEVPNManager {
	return NEVPNManager{objectivec.Object{ID: id}}
}
// NOTE: NEVPNManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEVPNManager] class.
//
// # Managing VPN configurations
//
//   - [INEVPNManager.LoadFromPreferencesWithCompletionHandler]: Load the VPN configuration from the Network Extension preferences.
//   - [INEVPNManager.SaveToPreferencesWithCompletionHandler]: Save the VPN configuration in the Network Extension preferences.
//   - [INEVPNManager.SetAuthorization]
//   - [INEVPNManager.RemoveFromPreferencesWithCompletionHandler]: Remove the VPN configuration from the Network Extension preferences.
//
// # Accessing VPN configuration properties
//
//   - [INEVPNManager.Enabled]: A Boolean used to toggle the enabled state of the VPN configuration.
//   - [INEVPNManager.SetEnabled]
//   - [INEVPNManager.ProtocolConfiguration]: An [NEVPNProtocol](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNProtocol>) object containing the configuration settings of the VPN tunneling protocol.
//   - [INEVPNManager.SetProtocolConfiguration]
//   - [INEVPNManager.LocalizedDescription]: A string containing the display name of the VPN configuration.
//   - [INEVPNManager.SetLocalizedDescription]
//   - [INEVPNManager.OnDemandEnabled]: A Boolean used to toggle the Connect On Demand capability.
//   - [INEVPNManager.SetOnDemandEnabled]
//   - [INEVPNManager.OnDemandRules]: An ordered list of Connect On Demand rules.
//   - [INEVPNManager.SetOnDemandRules]
//
// # Connecting and disconnecting VPN
//
//   - [INEVPNManager.Connection]: An [NEVPNConnection](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNConnection>) object that is used to control the VPN tunnel specified by the VPN configuration.
//
// # Errors
//
//   - [INEVPNManager.NEVPNErrorDomain]
//
// # Notifications
//
//   - [INEVPNManager.NEVPNConfigurationChange]: Posted after the VPN configuration stored in the Network Extension preferences changes.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager
type INEVPNManager interface {
	objectivec.IObject

	// Topic: Managing VPN configurations

	// Load the VPN configuration from the Network Extension preferences.
	LoadFromPreferencesWithCompletionHandler(completionHandler ErrorHandler)
	// Save the VPN configuration in the Network Extension preferences.
	SaveToPreferencesWithCompletionHandler(completionHandler ErrorHandler)
	SetAuthorization(authorization objectivec.IObject)
	// Remove the VPN configuration from the Network Extension preferences.
	RemoveFromPreferencesWithCompletionHandler(completionHandler ErrorHandler)

	// Topic: Accessing VPN configuration properties

	// A Boolean used to toggle the enabled state of the VPN configuration.
	Enabled() bool
	SetEnabled(value bool)
	// An [NEVPNProtocol](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNProtocol>) object containing the configuration settings of the VPN tunneling protocol.
	ProtocolConfiguration() INEVPNProtocol
	SetProtocolConfiguration(value INEVPNProtocol)
	// A string containing the display name of the VPN configuration.
	LocalizedDescription() string
	SetLocalizedDescription(value string)
	// A Boolean used to toggle the Connect On Demand capability.
	OnDemandEnabled() bool
	SetOnDemandEnabled(value bool)
	// An ordered list of Connect On Demand rules.
	OnDemandRules() []NEOnDemandRule
	SetOnDemandRules(value []NEOnDemandRule)

	// Topic: Connecting and disconnecting VPN

	// An [NEVPNConnection](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNConnection>) object that is used to control the VPN tunnel specified by the VPN configuration.
	Connection() INEVPNConnection

	// Topic: Errors

	NEVPNErrorDomain() string

	// Topic: Notifications

	// Posted after the VPN configuration stored in the Network Extension preferences changes.
	NEVPNConfigurationChange() foundation.NSString
}

// Init initializes the instance.
func (v NEVPNManager) Init() NEVPNManager {
	rv := objc.Send[NEVPNManager](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NEVPNManager) Autorelease() NEVPNManager {
	rv := objc.Send[NEVPNManager](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNManager creates a new NEVPNManager instance.
func NewNEVPNManager() NEVPNManager {
	class := getNEVPNManagerClass()
	rv := objc.Send[NEVPNManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Load the VPN configuration from the Network Extension preferences.
//
// completionHandler: A block that takes an [NSError] object. This block will be executed on the
// caller’s main thread after the load operation is complete. If the
// configuration does not exist in the preferences or is loaded successfully,
// the error parameter will be nil. If an error occurred while loading the
// configuration, the error parameter will be set to an [NSError] object
// containing details about the error. See `NEVPN Errors` for a list of
// possible errors.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// You must call this method at least once before calling `` for the first
// time after your app launches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/loadFromPreferences(completionHandler:)
func (v NEVPNManager) LoadFromPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](v.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), _block0)
}
// Save the VPN configuration in the Network Extension preferences.
//
// completionHandler: An optional block that takes an [NSError] object. If specified, this block
// will be executed on the caller’s main thread after the save operation is
// complete. If the configuration could not be saved to the preferences, the
// error parameter will be set to an [NSError] object containing details about
// the error. See `NEVPN Errors` for a list of possible errors. If the
// configuration is saved successfully then the error parameter will be set to
// nil.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// You must call [LoadFromPreferencesWithCompletionHandler]: at least once
// before calling this method the first time after your app launches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/saveToPreferences(completionHandler:)
func (v NEVPNManager) SaveToPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](v.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), _block0)
}
//
// authorization is a [systemconfiguration.AuthorizationRef].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/setAuthorization(_:)
// authorization is a [systemconfiguration.AuthorizationRef].
func (v NEVPNManager) SetAuthorization(authorization objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("setAuthorization:"), authorization)
}
// Remove the VPN configuration from the Network Extension preferences.
//
// completionHandler: An optional block that takes an [NSError] object. If specified, this block
// will be executed on the caller’s main thread after the removal operation
// is complete. If the configuration does not exist or an error occurs while
// removing it, the error parameter will be set to an [NSError] object
// containing details about the error. See `NEVPN Errors` for a list of
// possible errors. If the configuration is removed successfully then the
// error parameter will be set to nil.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// After the configuration is removed from the preferences the [NEVPNManager]
// object will still contain the configuration parameters. Calling
// [LoadFromPreferencesWithCompletionHandler]: will clear out the
// configuration parameters from the [NEVPNManager] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/removeFromPreferences(completionHandler:)
func (v NEVPNManager) RemoveFromPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](v.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), _block0)
}

// Access the single instance of [NEVPNManager].
//
// # Return Value
// 
// The [NEVPNManager] instance for the calling application.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/shared()
func (_NEVPNManagerClass NEVPNManagerClass) SharedManager() NEVPNManager {
	rv := objc.Send[objc.ID](objc.ID(_NEVPNManagerClass.class), objc.Sel("sharedManager"))
	return NEVPNManagerFromID(rv)
}

// A Boolean used to toggle the enabled state of the VPN configuration.
//
// # Discussion
// 
// A VPN configuration must be enabled before it can be used to bring up a VPN
// tunnel. Only one Personal VPN configuration can be enabled simultaneously
// on the system. If another Personal VPN configuration is enabled, then this
// property will be automatically set to [false] in the Network Extension
// preferences. Note that you will need to re-load the VPN configuration from
// the preferences in order to see the change in value. You can register with
// [NotificationCenter] to observe the [NEVPNConfigurationChangeNotification]
// notification for the [NEVPNManager] object so that your code can detect
// when the VPN configuration has been disabled.
//
// [NEVPNConfigurationChangeNotification]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConfigurationChangeNotification
// [NotificationCenter]: https://developer.apple.com/documentation/Foundation/NotificationCenter
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/isEnabled
func (v NEVPNManager) Enabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isEnabled"))
	return rv
}
func (v NEVPNManager) SetEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setEnabled:"), value)
}
// An [NEVPNProtocol] object containing the configuration settings of the VPN
// tunneling protocol.
//
// # Discussion
// 
// For [NEVPNManager] objects, this property can be set to either an
// [NEVPNProtocolIPSec] object or an [NEVPNProtocolIKEv2] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/protocolConfiguration
func (v NEVPNManager) ProtocolConfiguration() INEVPNProtocol {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("protocolConfiguration"))
	return NEVPNProtocolFromID(objc.ID(rv))
}
func (v NEVPNManager) SetProtocolConfiguration(value INEVPNProtocol) {
	objc.Send[struct{}](v.ID, objc.Sel("setProtocolConfiguration:"), value)
}
// A string containing the display name of the VPN configuration.
//
// # Discussion
// 
// This string is used as the display name of the VPN configuration in the
// system’s VPN settings UI. If this property is set to nil at the time that
// the configuration is created, it will be automatically set to the display
// name of the calling app.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/localizedDescription
func (v NEVPNManager) LocalizedDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("localizedDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v NEVPNManager) SetLocalizedDescription(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setLocalizedDescription:"), objc.String(value))
}
// A Boolean used to toggle the Connect On Demand capability.
//
// # Discussion
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/isOnDemandEnabled
func (v NEVPNManager) OnDemandEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isOnDemandEnabled"))
	return rv
}
func (v NEVPNManager) SetOnDemandEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setOnDemandEnabled:"), value)
}
// An ordered list of Connect On Demand rules.
//
// # Discussion
// 
// The VPN configuration can optionally be configured to connect automatically
// based on a variety of criteria specified in [NEOnDemandRule] objects. The
// [OnDemandRules] property contains the current set of Connect On Demand
// rules for the VPN configuration. Each rule is evaluated in order, and the
// first rule that matches all criteria on the current network is applied.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/onDemandRules
func (v NEVPNManager) OnDemandRules() []NEOnDemandRule {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("onDemandRules"))
	return objc.ConvertSlice(rv, func(id objc.ID) NEOnDemandRule {
		return NEOnDemandRuleFromID(id)
	})
}
func (v NEVPNManager) SetOnDemandRules(value []NEOnDemandRule) {
	objc.Send[struct{}](v.ID, objc.Sel("setOnDemandRules:"), objectivec.IObjectSliceToNSArray(value))
}
// An [NEVPNConnection] object that is used to control the VPN tunnel
// specified by the VPN configuration.
//
// # Discussion
// 
// The connection object is used to manually start and stop the VPN tunnel,
// and introspect the current status of the VPN tunnel. If the VPN
// configuration does not exist in the Network Extension preferences then the
// connection’s status is set to [NEVPNStatusInvalid].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/connection
func (v NEVPNManager) Connection() INEVPNConnection {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("connection"))
	return NEVPNConnectionFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/networkextension/nevpnerrordomain
func (v NEVPNManager) NEVPNErrorDomain() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("NEVPNErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}
// Posted after the VPN configuration stored in the Network Extension
// preferences changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NEVPNConfigurationChange
func (v NEVPNManager) NEVPNConfigurationChange() foundation.NSString {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("NEVPNConfigurationChange"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// LoadFromPreferences is a synchronous wrapper around [NEVPNManager.LoadFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v NEVPNManager) LoadFromPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	v.LoadFromPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SaveToPreferences is a synchronous wrapper around [NEVPNManager.SaveToPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v NEVPNManager) SaveToPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	v.SaveToPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RemoveFromPreferences is a synchronous wrapper around [NEVPNManager.RemoveFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v NEVPNManager) RemoveFromPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	v.RemoveFromPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}


// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEDNSProxyManager] class.
var (
	_NEDNSProxyManagerClass     NEDNSProxyManagerClass
	_NEDNSProxyManagerClassOnce sync.Once
)

func getNEDNSProxyManagerClass() NEDNSProxyManagerClass {
	_NEDNSProxyManagerClassOnce.Do(func() {
		_NEDNSProxyManagerClass = NEDNSProxyManagerClass{class: objc.GetClass("NEDNSProxyManager")}
	})
	return _NEDNSProxyManagerClass
}

// GetNEDNSProxyManagerClass returns the class object for NEDNSProxyManager.
func GetNEDNSProxyManagerClass() NEDNSProxyManagerClass {
	return getNEDNSProxyManagerClass()
}

type NEDNSProxyManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEDNSProxyManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEDNSProxyManagerClass) Alloc() NEDNSProxyManager {
	rv := objc.Send[NEDNSProxyManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object to create and manage an DNS proxy provider’s configuration.
//
// # Overview
// 
// A DNS proxy allows your app to intercept all DNS traffic generated on a
// device. You can use this capability to provide services like DNS traffic
// encryption, typically by redirecting DNS traffic to your own server. You
// usually do this in the context of managed devices, such as those owned by a
// school or an enterprise.
// 
// You create a DNS proxy as an app extension based on a custom subclass of
// the [NEDNSProxyProvider] class. You enable and configure this proxy from
// within your app using the singleton proxy manager instance provided by the
// [NEDNSProxyManager.SharedManager] type method of the [NEDNSProxyManager] class. For example,
// for a proxy that performs a simple redirect, you can use the proxy manager
// to define and dynamically configure the destination IP address of the
// redirected traffic.
// 
// Instances of the proxy manager are thread safe.
//
// # Managing the DNS proxy configuration
//
//   - [NEDNSProxyManager.LoadFromPreferencesWithCompletionHandler]: Loads the current DNS proxy configuration from the caller’s DNS proxy preferences.
//   - [NEDNSProxyManager.SaveToPreferencesWithCompletionHandler]: Saves the DNS proxy configuration in the caller’s DNS proxy preferences.
//   - [NEDNSProxyManager.RemoveFromPreferencesWithCompletionHandler]: Removes the DNS proxy configuration from the caller’s DNS proxy preferences.
//
// # Accessing DNS proxy configuration properties
//
//   - [NEDNSProxyManager.Enabled]: The status of a DNS proxy.
//   - [NEDNSProxyManager.SetEnabled]
//   - [NEDNSProxyManager.ProviderProtocol]: The provider-specific portion of the DNS proxy configuration.
//   - [NEDNSProxyManager.SetProviderProtocol]
//   - [NEDNSProxyManager.LocalizedDescription]: A description of the DNS proxy.
//   - [NEDNSProxyManager.SetLocalizedDescription]
//
// # Notifications
//
//   - [NEDNSProxyManager.NEDNSProxyConfigurationDidChange]: A notification that is posted when the DNS proxy configuration changes.
//
// # Errors
//
//   - [NEDNSProxyManager.NEDNSProxyErrorDomain]: The DNS proxy error domain.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager
type NEDNSProxyManager struct {
	objectivec.Object
}

// NEDNSProxyManagerFromID constructs a [NEDNSProxyManager] from an objc.ID.
//
// An object to create and manage an DNS proxy provider’s configuration.
func NEDNSProxyManagerFromID(id objc.ID) NEDNSProxyManager {
	return NEDNSProxyManager{objectivec.Object{ID: id}}
}
// NOTE: NEDNSProxyManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEDNSProxyManager] class.
//
// # Managing the DNS proxy configuration
//
//   - [INEDNSProxyManager.LoadFromPreferencesWithCompletionHandler]: Loads the current DNS proxy configuration from the caller’s DNS proxy preferences.
//   - [INEDNSProxyManager.SaveToPreferencesWithCompletionHandler]: Saves the DNS proxy configuration in the caller’s DNS proxy preferences.
//   - [INEDNSProxyManager.RemoveFromPreferencesWithCompletionHandler]: Removes the DNS proxy configuration from the caller’s DNS proxy preferences.
//
// # Accessing DNS proxy configuration properties
//
//   - [INEDNSProxyManager.Enabled]: The status of a DNS proxy.
//   - [INEDNSProxyManager.SetEnabled]
//   - [INEDNSProxyManager.ProviderProtocol]: The provider-specific portion of the DNS proxy configuration.
//   - [INEDNSProxyManager.SetProviderProtocol]
//   - [INEDNSProxyManager.LocalizedDescription]: A description of the DNS proxy.
//   - [INEDNSProxyManager.SetLocalizedDescription]
//
// # Notifications
//
//   - [INEDNSProxyManager.NEDNSProxyConfigurationDidChange]: A notification that is posted when the DNS proxy configuration changes.
//
// # Errors
//
//   - [INEDNSProxyManager.NEDNSProxyErrorDomain]: The DNS proxy error domain.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager
type INEDNSProxyManager interface {
	objectivec.IObject

	// Topic: Managing the DNS proxy configuration

	// Loads the current DNS proxy configuration from the caller’s DNS proxy preferences.
	LoadFromPreferencesWithCompletionHandler(completionHandler ErrorHandler)
	// Saves the DNS proxy configuration in the caller’s DNS proxy preferences.
	SaveToPreferencesWithCompletionHandler(completionHandler ErrorHandler)
	// Removes the DNS proxy configuration from the caller’s DNS proxy preferences.
	RemoveFromPreferencesWithCompletionHandler(completionHandler ErrorHandler)

	// Topic: Accessing DNS proxy configuration properties

	// The status of a DNS proxy.
	Enabled() bool
	SetEnabled(value bool)
	// The provider-specific portion of the DNS proxy configuration.
	ProviderProtocol() INEDNSProxyProviderProtocol
	SetProviderProtocol(value INEDNSProxyProviderProtocol)
	// A description of the DNS proxy.
	LocalizedDescription() string
	SetLocalizedDescription(value string)

	// Topic: Notifications

	// A notification that is posted when the DNS proxy configuration changes.
	NEDNSProxyConfigurationDidChange() foundation.NSString

	// Topic: Errors

	// The DNS proxy error domain.
	NEDNSProxyErrorDomain() string
}

// Init initializes the instance.
func (d NEDNSProxyManager) Init() NEDNSProxyManager {
	rv := objc.Send[NEDNSProxyManager](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NEDNSProxyManager) Autorelease() NEDNSProxyManager {
	rv := objc.Send[NEDNSProxyManager](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSProxyManager creates a new NEDNSProxyManager instance.
func NewNEDNSProxyManager() NEDNSProxyManager {
	class := getNEDNSProxyManagerClass()
	rv := objc.Send[NEDNSProxyManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Loads the current DNS proxy configuration from the caller’s DNS proxy
// preferences.
//
// completionHandler: A block called when the load operation completes. If the operation fails,
// an error instance passed to this block describes the problem. Otherwise,
// the error is `nil`. See [NEDNSProxyManagerError] for the list of possible
// errors.
// //
// [NEDNSProxyManagerError]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError
//
// # Discussion
// 
// Initially, the DNS proxy configuration comes from a configuration profile
// stored on the device in a managed environment, as described in
// [Configuration Profile Reference].
// 
// When you want to inspect or make changes to the configuration, you call the
// proxy manager’s [LoadFromPreferencesWithCompletionHandler] method. This
// causes the system to load the configuration into the manager’s
// [ProviderProtocol] and [Enabled] properties.
// 
// If you modify the configuration stored in these properties, you must then
// call the [SaveToPreferencesWithCompletionHandler] method to make the
// changes take effect. Saving the preferences also stores the modified
// configuration on disk for use the next time the proxy is started or the
// configuration is loaded.
//
// [Configuration Profile Reference]: https://developer.apple.com/library/archive/featuredarticles/iPhoneConfigurationProfileRef/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010206
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/loadFromPreferences(completionHandler:)
func (d NEDNSProxyManager) LoadFromPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), _block0)
}
// Saves the DNS proxy configuration in the caller’s DNS proxy preferences.
//
// completionHandler: A block called when the save operation completes. If the operation fails,
// an error instance passed to this block describes the problem. Otherwise,
// the error is `nil`. See [NEDNSProxyManagerError] for the list of possible
// errors.
// //
// [NEDNSProxyManagerError]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError
//
// # Discussion
// 
// If you alter the DNS proxy configuration that you load into the proxy
// manager’s properties using a call to the
// [LoadFromPreferencesWithCompletionHandler] method, you must then call the
// [SaveToPreferencesWithCompletionHandler] method to make the changes take
// effect. Saving also stores the modified configuration for the next time the
// proxy is started or the configuration loaded.
// 
// Trying to save preferences before loading them produces an error.
// 
// If the DNS proxy is enabled, it becomes active as a result of this call.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/saveToPreferences(completionHandler:)
func (d NEDNSProxyManager) SaveToPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), _block0)
}
// Removes the DNS proxy configuration from the caller’s DNS proxy
// preferences.
//
// completionHandler: A block called when the remove operation completes. If the operation fails,
// an error instance passed to this block describes the problem. Otherwise,
// the error is `nil`. See [NEDNSProxyManagerError] for the list of possible
// errors.
// //
// [NEDNSProxyManagerError]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError
//
// # Discussion
// 
// If you use a device without an installed configuration profile during
// development, your app can create the DNS proxy configuration from scratch.
// You first call the [LoadFromPreferencesWithCompletionHandler] method to
// retrieve the empty configuration. You then make updates and call the
// [SaveToPreferencesWithCompletionHandler] method to store them. To remove
// the configuration, call the [RemoveFromPreferencesWithCompletionHandler]
// method. This allows you to restore the device to a clean, unconfigured
// state.
// 
// In a production environment, however, a configuration profile placed in the
// system by an external process typically provides the baseline DNS proxy
// configuration. Your app can modify this configuration at runtime using the
// same load-modify-save steps, but cannot remove the configuration entirely.
// An attempt to remove the configuration when a configuration profile is
// present on the device results in a
// [NEDNSProxyManagerError.configurationCannotBeRemoved] error.
// 
// If the DNS proxy is enabled, it becomes disabled as a result of this call.
//
// [NEDNSProxyManagerError.configurationCannotBeRemoved]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError/configurationCannotBeRemoved
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/removeFromPreferences(completionHandler:)
func (d NEDNSProxyManager) RemoveFromPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), _block0)
}

// Returns a singleton DNS proxy manager instance.
//
// # Return Value
// 
// The [NEDNSProxyManager] instance for the app.
//
// # Discussion
// 
// Each app is allowed to create a single DNS proxy manager. The
// [SharedManager] type method returns a singleton [NEDNSProxyManager]
// instance that your app can use to manage any DNS proxy instances that it
// creates.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/shared()
func (_NEDNSProxyManagerClass NEDNSProxyManagerClass) SharedManager() NEDNSProxyManager {
	rv := objc.Send[objc.ID](objc.ID(_NEDNSProxyManagerClass.class), objc.Sel("sharedManager"))
	return NEDNSProxyManagerFromID(rv)
}

// The status of a DNS proxy.
//
// # Discussion
// 
// Only one DNS proxy can be active in the system at a time. Therefore,
// setting this property to [true] disables any DNS proxy configurations of
// other apps. Similarly, the system sets this property to [false] when any
// other DNS proxy configuration is enabled.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/isEnabled
func (d NEDNSProxyManager) Enabled() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isEnabled"))
	return rv
}
func (d NEDNSProxyManager) SetEnabled(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setEnabled:"), value)
}
// The provider-specific portion of the DNS proxy configuration.
//
// # Discussion
// 
// As the author of the DNS proxy, you decide what configuration the proxy
// needs. For example, if your proxy requires the IP addresses of servers to
// which DNS traffic can be redirected, you can use an array of strings to
// hold these values.
// 
// Initially, you store this array in the configuration profile, as described
// in [Configuration Profile Reference]. When you want to inspect or modify
// this data, you call [LoadFromPreferencesWithCompletionHandler] to pull the
// configuration into memory. You access this memory through the proxy
// manager’s [ProviderProtocol] property.
//
// [Configuration Profile Reference]: https://developer.apple.com/library/archive/featuredarticles/iPhoneConfigurationProfileRef/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010206
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/providerProtocol
func (d NEDNSProxyManager) ProviderProtocol() INEDNSProxyProviderProtocol {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("providerProtocol"))
	return NEDNSProxyProviderProtocolFromID(objc.ID(rv))
}
func (d NEDNSProxyManager) SetProviderProtocol(value INEDNSProxyProviderProtocol) {
	objc.Send[struct{}](d.ID, objc.Sel("setProviderProtocol:"), value)
}
// A description of the DNS proxy.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/localizedDescription
func (d NEDNSProxyManager) LocalizedDescription() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("localizedDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (d NEDNSProxyManager) SetLocalizedDescription(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setLocalizedDescription:"), objc.String(value))
}
// A notification that is posted when the DNS proxy configuration changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NEDNSProxyConfigurationDidChange
func (d NEDNSProxyManager) NEDNSProxyConfigurationDidChange() foundation.NSString {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("NEDNSProxyConfigurationDidChange"))
	return foundation.NSStringFromID(objc.ID(rv))
}
// The DNS proxy error domain.
//
// See: https://developer.apple.com/documentation/networkextension/nednsproxyerrordomain
func (d NEDNSProxyManager) NEDNSProxyErrorDomain() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("NEDNSProxyErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// LoadFromPreferences is a synchronous wrapper around [NEDNSProxyManager.LoadFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NEDNSProxyManager) LoadFromPreferences(ctx context.Context) error {
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

// SaveToPreferences is a synchronous wrapper around [NEDNSProxyManager.SaveToPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NEDNSProxyManager) SaveToPreferences(ctx context.Context) error {
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

// RemoveFromPreferences is a synchronous wrapper around [NEDNSProxyManager.RemoveFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NEDNSProxyManager) RemoveFromPreferences(ctx context.Context) error {
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


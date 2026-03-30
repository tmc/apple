// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEFilterManager] class.
var (
	_NEFilterManagerClass     NEFilterManagerClass
	_NEFilterManagerClassOnce sync.Once
)

func getNEFilterManagerClass() NEFilterManagerClass {
	_NEFilterManagerClassOnce.Do(func() {
		_NEFilterManagerClass = NEFilterManagerClass{class: objc.GetClass("NEFilterManager")}
	})
	return _NEFilterManagerClass
}

// GetNEFilterManagerClass returns the class object for NEFilterManager.
func GetNEFilterManagerClass() NEFilterManagerClass {
	return getNEFilterManagerClass()
}

type NEFilterManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterManagerClass) Alloc() NEFilterManager {
	rv := objc.Send[NEFilterManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object to create and manage a content filter’s configuration.
//
// # Overview
//
// Each app is allowed to create a single filter configuration. The
// [NEFilterManager] class has a class method ([NEFilterManager.SharedManager]) that provides
// access to a single [NEFilterManager] instance. This single instance
// corresponds to a single filter configuration.
//
// The filter configuration is stored in the Network Extension preferences
// which are managed by the Network Extension framework. The filter
// configuration must be explicitly loaded into memory from the Network
// Extension preferences before it can be used, and any changes must be
// explicitly saved to the Network Extension preferences before taking effect
// on the system.
//
// # Profile Configuration
//
// Filter configurations are created using configuration profiles. See
// [WebContentFilter] for more information. To specify that a filter
// configuration created via a profile payload is associated with a particular
// app (and therefore allow the app to use [NEFilterManager] to manage the
// configuration), the app’s bundle identifier must be set as the value of
// the [PluginBundleID] field in the profile payload.
//
// # Filter Provider Extensions
//
// Apps that use [NEFilterManager] are required to contain two Filter Provider
// extensions that together perform the task of examining network content and
// making pass and block decisions. See the [NEFilterControlProvider] and
// [NEFilterDataProvider] classes for more details about these extensions.
//
// # Managing the filter configuration
//
//   - [NEFilterManager.LoadFromPreferencesWithCompletionHandler]: Load the filter configuration from the Network Extension preferences.
//   - [NEFilterManager.SaveToPreferencesWithCompletionHandler]: Save the filter configuration in the Network Extension preferences.
//   - [NEFilterManager.RemoveFromPreferencesWithCompletionHandler]: Remove the filter configuration from the Network Extension preferences.
//
// # Accessing filter configuration properties
//
//   - [NEFilterManager.Enabled]: A Boolean used to toggle the enabled state of the filter.
//   - [NEFilterManager.SetEnabled]
//   - [NEFilterManager.ProviderConfiguration]: A [NEFilterProviderConfiguration](<doc://com.apple.networkextension/documentation/NetworkExtension/NEFilterProviderConfiguration>) object containing the filter configuration settings.
//   - [NEFilterManager.SetProviderConfiguration]
//   - [NEFilterManager.LocalizedDescription]: A string containing a description of the filter configuration.
//   - [NEFilterManager.SetLocalizedDescription]
//
// # Prioritizing filters
//
//   - [NEFilterManager.Grade]: The grade of the filter, which determines when it acts relative to other filters.
//   - [NEFilterManager.SetGrade]
//
// # Errors
//
//   - [NEFilterManager.NEFilterErrorDomain]: The domain for errors resulting from calls to the filter manager.
//
// # Notifications
//
//   - [NEFilterManager.NEFilterConfigurationDidChange]: Posted after the filter configuration stored in the Network Extension preferences changes.
//
// # Instance Properties
//
//   - [NEFilterManager.DisableEncryptedDNSSettings]
//   - [NEFilterManager.SetDisableEncryptedDNSSettings]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager
//
// [NEFilterControlProvider]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider
// [WebContentFilter]: https://developer.apple.com/documentation/DeviceManagement/WebContentFilter
type NEFilterManager struct {
	objectivec.Object
}

// NEFilterManagerFromID constructs a [NEFilterManager] from an objc.ID.
//
// An object to create and manage a content filter’s configuration.
func NEFilterManagerFromID(id objc.ID) NEFilterManager {
	return NEFilterManager{objectivec.Object{ID: id}}
}

// NOTE: NEFilterManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterManager] class.
//
// # Managing the filter configuration
//
//   - [INEFilterManager.LoadFromPreferencesWithCompletionHandler]: Load the filter configuration from the Network Extension preferences.
//   - [INEFilterManager.SaveToPreferencesWithCompletionHandler]: Save the filter configuration in the Network Extension preferences.
//   - [INEFilterManager.RemoveFromPreferencesWithCompletionHandler]: Remove the filter configuration from the Network Extension preferences.
//
// # Accessing filter configuration properties
//
//   - [INEFilterManager.Enabled]: A Boolean used to toggle the enabled state of the filter.
//   - [INEFilterManager.SetEnabled]
//   - [INEFilterManager.ProviderConfiguration]: A [NEFilterProviderConfiguration](<doc://com.apple.networkextension/documentation/NetworkExtension/NEFilterProviderConfiguration>) object containing the filter configuration settings.
//   - [INEFilterManager.SetProviderConfiguration]
//   - [INEFilterManager.LocalizedDescription]: A string containing a description of the filter configuration.
//   - [INEFilterManager.SetLocalizedDescription]
//
// # Prioritizing filters
//
//   - [INEFilterManager.Grade]: The grade of the filter, which determines when it acts relative to other filters.
//   - [INEFilterManager.SetGrade]
//
// # Errors
//
//   - [INEFilterManager.NEFilterErrorDomain]: The domain for errors resulting from calls to the filter manager.
//
// # Notifications
//
//   - [INEFilterManager.NEFilterConfigurationDidChange]: Posted after the filter configuration stored in the Network Extension preferences changes.
//
// # Instance Properties
//
//   - [INEFilterManager.DisableEncryptedDNSSettings]
//   - [INEFilterManager.SetDisableEncryptedDNSSettings]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager
type INEFilterManager interface {
	objectivec.IObject

	// Topic: Managing the filter configuration

	// Load the filter configuration from the Network Extension preferences.
	LoadFromPreferencesWithCompletionHandler(completionHandler ErrorHandler)
	// Save the filter configuration in the Network Extension preferences.
	SaveToPreferencesWithCompletionHandler(completionHandler ErrorHandler)
	// Remove the filter configuration from the Network Extension preferences.
	RemoveFromPreferencesWithCompletionHandler(completionHandler ErrorHandler)

	// Topic: Accessing filter configuration properties

	// A Boolean used to toggle the enabled state of the filter.
	Enabled() bool
	SetEnabled(value bool)
	// A [NEFilterProviderConfiguration](<doc://com.apple.networkextension/documentation/NetworkExtension/NEFilterProviderConfiguration>) object containing the filter configuration settings.
	ProviderConfiguration() INEFilterProviderConfiguration
	SetProviderConfiguration(value INEFilterProviderConfiguration)
	// A string containing a description of the filter configuration.
	LocalizedDescription() string
	SetLocalizedDescription(value string)

	// Topic: Prioritizing filters

	// The grade of the filter, which determines when it acts relative to other filters.
	Grade() NEFilterManagerGrade
	SetGrade(value NEFilterManagerGrade)

	// Topic: Errors

	// The domain for errors resulting from calls to the filter manager.
	NEFilterErrorDomain() string

	// Topic: Notifications

	// Posted after the filter configuration stored in the Network Extension preferences changes.
	NEFilterConfigurationDidChange() foundation.NSString

	// Topic: Instance Properties

	DisableEncryptedDNSSettings() bool
	SetDisableEncryptedDNSSettings(value bool)
}

// Init initializes the instance.
func (f NEFilterManager) Init() NEFilterManager {
	rv := objc.Send[NEFilterManager](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterManager) Autorelease() NEFilterManager {
	rv := objc.Send[NEFilterManager](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterManager creates a new NEFilterManager instance.
func NewNEFilterManager() NEFilterManager {
	class := getNEFilterManagerClass()
	rv := objc.Send[NEFilterManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Load the filter configuration from the Network Extension preferences.
//
// completionHandler: A block that takes an [NSError] object. This block will be executed on the
// caller’s main thread after the load operation is complete. If the
// configuration does not exist in the Network Extension preferences or is
// loaded successfully, the error parameter will be nil. If an error occurred
// while loading the configuration, the error parameter will be set to an
// [NSError] object containing details about the error. See
// [NEFilterManagerError] for a list of possible errors.
//
// # Discussion
//
// You must call this method at least once before calling
// `saveToPreferencesWithCompletionHandler`: for the first time after your app
// launches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/loadFromPreferences(completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (f NEFilterManager) LoadFromPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), _block0)
}

// Save the filter configuration in the Network Extension preferences.
//
// completionHandler: A block that takes an [NSError] object. This block will be executed on the
// caller’s main thread after the save operation is complete. If the
// configuration could not be saved to the preferences, the error parameter
// will be set to an [NSError] object containing details about the error. See
// [NEFilterManagerError] for a list of possible errors. If the configuration
// is saved successfully then the error parameter will be set to nil.
//
// # Discussion
//
// You must call “ at least once before calling this method the first time
// after your app launches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/saveToPreferences(completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (f NEFilterManager) SaveToPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), _block0)
}

// Remove the filter configuration from the Network Extension preferences.
//
// completionHandler: A block that takes an [NSError] object. This block will be executed on the
// caller’s main thread after the removal operation is complete. If the
// configuration does not exist in the Network Extension preferences or an
// error occurs while removing it, the error parameter will be set to an
// [NSError] object containing details about the error. See
// [NEFilterManagerError] for a list of possible errors. If the configuration
// is removed successfully the error parameter will be set to nil.
//
// # Discussion
//
// After the configuration is removed from the preferences the
// [NEFilterManager] object will still contain the configuration parameters.
// Calling “ will clear out the configuration parameters from the
// [NEFilterManager] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/removeFromPreferences(completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (f NEFilterManager) RemoveFromPreferencesWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), _block0)
}

// Access the single instance of [NEFilterManager].
//
// # Return Value
//
// The [NEFilterManager] instance for the calling application.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/shared()
func (_NEFilterManagerClass NEFilterManagerClass) SharedManager() NEFilterManager {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterManagerClass.class), objc.Sel("sharedManager"))
	return NEFilterManagerFromID(rv)
}

// A Boolean used to toggle the enabled state of the filter.
//
// # Discussion
//
// Setting this property to true and saving the configuration will disable all
// other network content filters on the system, and will start the filter’s
// Filter Provider extensions. Setting this property to false and saving the
// configuration will disable the filter and stop the filter’s Filter
// Provider extensions.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/isEnabled
func (f NEFilterManager) Enabled() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isEnabled"))
	return rv
}
func (f NEFilterManager) SetEnabled(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setEnabled:"), value)
}

// A [NEFilterProviderConfiguration] object containing the filter
// configuration settings.
//
// # Discussion
//
// If this property is nil after calling “, then the filter configuration
// does not exist in the Network Extension preferences.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/providerConfiguration
func (f NEFilterManager) ProviderConfiguration() INEFilterProviderConfiguration {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("providerConfiguration"))
	return NEFilterProviderConfigurationFromID(objc.ID(rv))
}
func (f NEFilterManager) SetProviderConfiguration(value INEFilterProviderConfiguration) {
	objc.Send[struct{}](f.ID, objc.Sel("setProviderConfiguration:"), value)
}

// A string containing a description of the filter configuration.
//
// # Discussion
//
// If this property is set to nil at the time that the configuration is
// created, it will be automatically set to the display name of the calling
// app.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/localizedDescription
func (f NEFilterManager) LocalizedDescription() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("localizedDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (f NEFilterManager) SetLocalizedDescription(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setLocalizedDescription:"), objc.String(value))
}

// The grade of the filter, which determines when it acts relative to other
// filters.
//
// # Discussion
//
// The default grade is [NEFilterManagerGradeFirewall].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/grade-swift.property
func (f NEFilterManager) Grade() NEFilterManagerGrade {
	rv := objc.Send[NEFilterManagerGrade](f.ID, objc.Sel("grade"))
	return NEFilterManagerGrade(rv)
}
func (f NEFilterManager) SetGrade(value NEFilterManagerGrade) {
	objc.Send[struct{}](f.ID, objc.Sel("setGrade:"), value)
}

// The domain for errors resulting from calls to the filter manager.
//
// See: https://developer.apple.com/documentation/networkextension/nefiltererrordomain
func (f NEFilterManager) NEFilterErrorDomain() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("NEFilterErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// Posted after the filter configuration stored in the Network Extension
// preferences changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NEFilterConfigurationDidChange
func (f NEFilterManager) NEFilterConfigurationDidChange() foundation.NSString {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("NEFilterConfigurationDidChange"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/disableEncryptedDNSSettings
func (f NEFilterManager) DisableEncryptedDNSSettings() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("disableEncryptedDNSSettings"))
	return rv
}
func (f NEFilterManager) SetDisableEncryptedDNSSettings(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setDisableEncryptedDNSSettings:"), value)
}

// LoadFromPreferences is a synchronous wrapper around [NEFilterManager.LoadFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NEFilterManager) LoadFromPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	f.LoadFromPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SaveToPreferences is a synchronous wrapper around [NEFilterManager.SaveToPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NEFilterManager) SaveToPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	f.SaveToPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RemoveFromPreferences is a synchronous wrapper around [NEFilterManager.RemoveFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NEFilterManager) RemoveFromPreferences(ctx context.Context) error {
	done := make(chan error, 1)
	f.RemoveFromPreferencesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

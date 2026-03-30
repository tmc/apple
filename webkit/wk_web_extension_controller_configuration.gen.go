// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtensionControllerConfiguration] class.
var (
	_WKWebExtensionControllerConfigurationClass     WKWebExtensionControllerConfigurationClass
	_WKWebExtensionControllerConfigurationClassOnce sync.Once
)

func getWKWebExtensionControllerConfigurationClass() WKWebExtensionControllerConfigurationClass {
	_WKWebExtensionControllerConfigurationClassOnce.Do(func() {
		_WKWebExtensionControllerConfigurationClass = WKWebExtensionControllerConfigurationClass{class: objc.GetClass("WKWebExtensionControllerConfiguration")}
	})
	return _WKWebExtensionControllerConfigurationClass
}

// GetWKWebExtensionControllerConfigurationClass returns the class object for WKWebExtensionControllerConfiguration.
func GetWKWebExtensionControllerConfigurationClass() WKWebExtensionControllerConfigurationClass {
	return getWKWebExtensionControllerConfigurationClass()
}

type WKWebExtensionControllerConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionControllerConfigurationClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionControllerConfigurationClass) Alloc() WKWebExtensionControllerConfiguration {
	rv := objc.Send[WKWebExtensionControllerConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// A [WKWebExtensionControllerConfiguration] object with which to initialize a
// web extension controller.
//
// # Overview
//
// Contains properties used to configure a [WKWebExtensionController].
//
// # Instance Properties
//
//   - [WKWebExtensionControllerConfiguration.DefaultWebsiteDataStore]: The default data store for website data and cookie access in extension contexts.
//   - [WKWebExtensionControllerConfiguration.SetDefaultWebsiteDataStore]
//   - [WKWebExtensionControllerConfiguration.Identifier]: The unique identifier used for persistent configuration storage, or `nil` when it is the default or not persistent.
//   - [WKWebExtensionControllerConfiguration.Persistent]: A Boolean value indicating if this context will write data to the the file system.
//   - [WKWebExtensionControllerConfiguration.WebViewConfiguration]: The web view configuration to be used as a basis for configuring web views in extension contexts.
//   - [WKWebExtensionControllerConfiguration.SetWebViewConfiguration]
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class
type WKWebExtensionControllerConfiguration struct {
	objectivec.Object
}

// WKWebExtensionControllerConfigurationFromID constructs a [WKWebExtensionControllerConfiguration] from an objc.ID.
//
// A [WKWebExtensionControllerConfiguration] object with which to initialize a
// web extension controller.
func WKWebExtensionControllerConfigurationFromID(id objc.ID) WKWebExtensionControllerConfiguration {
	return WKWebExtensionControllerConfiguration{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtensionControllerConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtensionControllerConfiguration] class.
//
// # Instance Properties
//
//   - [IWKWebExtensionControllerConfiguration.DefaultWebsiteDataStore]: The default data store for website data and cookie access in extension contexts.
//   - [IWKWebExtensionControllerConfiguration.SetDefaultWebsiteDataStore]
//   - [IWKWebExtensionControllerConfiguration.Identifier]: The unique identifier used for persistent configuration storage, or `nil` when it is the default or not persistent.
//   - [IWKWebExtensionControllerConfiguration.Persistent]: A Boolean value indicating if this context will write data to the the file system.
//   - [IWKWebExtensionControllerConfiguration.WebViewConfiguration]: The web view configuration to be used as a basis for configuring web views in extension contexts.
//   - [IWKWebExtensionControllerConfiguration.SetWebViewConfiguration]
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class
type IWKWebExtensionControllerConfiguration interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The default data store for website data and cookie access in extension contexts.
	DefaultWebsiteDataStore() IWKWebsiteDataStore
	SetDefaultWebsiteDataStore(value IWKWebsiteDataStore)
	// The unique identifier used for persistent configuration storage, or `nil` when it is the default or not persistent.
	Identifier() foundation.NSUUID
	// A Boolean value indicating if this context will write data to the the file system.
	Persistent() bool
	// The web view configuration to be used as a basis for configuring web views in extension contexts.
	WebViewConfiguration() IWKWebViewConfiguration
	SetWebViewConfiguration(value IWKWebViewConfiguration)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (w WKWebExtensionControllerConfiguration) Init() WKWebExtensionControllerConfiguration {
	rv := objc.Send[WKWebExtensionControllerConfiguration](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtensionControllerConfiguration) Autorelease() WKWebExtensionControllerConfiguration {
	rv := objc.Send[WKWebExtensionControllerConfiguration](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtensionControllerConfiguration creates a new WKWebExtensionControllerConfiguration instance.
func NewWKWebExtensionControllerConfiguration() WKWebExtensionControllerConfiguration {
	class := getWKWebExtensionControllerConfigurationClass()
	rv := objc.Send[WKWebExtensionControllerConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a new configuration that is persistent and unique for the specified
// identifier.
//
// # Discussion
//
// If a [WKWebExtensionController] is associated with a unique persistent
// configuration, data will be written to the file system in a unique location
// based on the specified identifier.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/init(identifier:)
func NewWebExtensionControllerConfigurationWithIdentifier(identifier foundation.NSUUID) WKWebExtensionControllerConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getWKWebExtensionControllerConfigurationClass().class), objc.Sel("configurationWithIdentifier:"), identifier)
	return WKWebExtensionControllerConfigurationFromID(rv)
}

func (w WKWebExtensionControllerConfiguration) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](w.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns a new default configuration that is persistent and not unique.
//
// # Discussion
//
// If a [WKWebExtensionController] is associated with a persistent
// configuration, data will be written to the file system in a common
// location. When using multiple extension controllers, each controller should
// use a unique configuration to avoid conflicts.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/default()
func (_WKWebExtensionControllerConfigurationClass WKWebExtensionControllerConfigurationClass) DefaultConfiguration() WKWebExtensionControllerConfiguration {
	rv := objc.Send[objc.ID](objc.ID(_WKWebExtensionControllerConfigurationClass.class), objc.Sel("defaultConfiguration"))
	return WKWebExtensionControllerConfigurationFromID(rv)
}

// Returns a new non-persistent configuration.
//
// # Discussion
//
// If a [WKWebExtensionController] is associated with a non-persistent
// configuration, no data will be written to the file system. This is useful
// for extensions in “private browsing” situations.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/nonPersistent()
func (_WKWebExtensionControllerConfigurationClass WKWebExtensionControllerConfigurationClass) NonPersistentConfiguration() WKWebExtensionControllerConfiguration {
	rv := objc.Send[objc.ID](objc.ID(_WKWebExtensionControllerConfigurationClass.class), objc.Sel("nonPersistentConfiguration"))
	return WKWebExtensionControllerConfigurationFromID(rv)
}

// The default data store for website data and cookie access in extension
// contexts.
//
// # Discussion
//
// This property sets the primary data store for managing website data,
// including cookies, which extensions can access, subject to the granted
// permissions within the extension contexts. Defaults to [DefaultDataStore].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/defaultWebsiteDataStore
func (w WKWebExtensionControllerConfiguration) DefaultWebsiteDataStore() IWKWebsiteDataStore {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("defaultWebsiteDataStore"))
	return WKWebsiteDataStoreFromID(objc.ID(rv))
}
func (w WKWebExtensionControllerConfiguration) SetDefaultWebsiteDataStore(value IWKWebsiteDataStore) {
	objc.Send[struct{}](w.ID, objc.Sel("setDefaultWebsiteDataStore:"), value)
}

// The unique identifier used for persistent configuration storage, or `nil`
// when it is the default or not persistent.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/identifier
func (w WKWebExtensionControllerConfiguration) Identifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("identifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// A Boolean value indicating if this context will write data to the the file
// system.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/isPersistent
func (w WKWebExtensionControllerConfiguration) Persistent() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isPersistent"))
	return rv
}

// The web view configuration to be used as a basis for configuring web views
// in extension contexts.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/Configuration-swift.class/webViewConfiguration
func (w WKWebExtensionControllerConfiguration) WebViewConfiguration() IWKWebViewConfiguration {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("webViewConfiguration"))
	return WKWebViewConfigurationFromID(objc.ID(rv))
}
func (w WKWebExtensionControllerConfiguration) SetWebViewConfiguration(value IWKWebViewConfiguration) {
	objc.Send[struct{}](w.ID, objc.Sel("setWebViewConfiguration:"), value)
}

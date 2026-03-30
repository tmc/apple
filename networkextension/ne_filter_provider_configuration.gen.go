// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEFilterProviderConfiguration] class.
var (
	_NEFilterProviderConfigurationClass     NEFilterProviderConfigurationClass
	_NEFilterProviderConfigurationClassOnce sync.Once
)

func getNEFilterProviderConfigurationClass() NEFilterProviderConfigurationClass {
	_NEFilterProviderConfigurationClassOnce.Do(func() {
		_NEFilterProviderConfigurationClass = NEFilterProviderConfigurationClass{class: objc.GetClass("NEFilterProviderConfiguration")}
	})
	return _NEFilterProviderConfigurationClass
}

// GetNEFilterProviderConfigurationClass returns the class object for NEFilterProviderConfiguration.
func GetNEFilterProviderConfigurationClass() NEFilterProviderConfigurationClass {
	return getNEFilterProviderConfigurationClass()
}

type NEFilterProviderConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterProviderConfigurationClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterProviderConfigurationClass) Alloc() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Configuration parameters for a content filter.
//
// # Configuring filter behavior
//
//   - [NEFilterProviderConfiguration.FilterSockets]: A Boolean value that indicates that the system applies the filter to flows of network data originated from sockets.
//   - [NEFilterProviderConfiguration.SetFilterSockets]
//   - [NEFilterProviderConfiguration.FilterPackets]: A Boolean value that indicates that the system applies the filter to packets of network data.
//   - [NEFilterProviderConfiguration.SetFilterPackets]
//
// # Accessing the filter configuration
//
//   - [NEFilterProviderConfiguration.VendorConfiguration]: A dictionary of provider-specific configuration settings.
//   - [NEFilterProviderConfiguration.SetVendorConfiguration]
//   - [NEFilterProviderConfiguration.ServerAddress]: The address of a server that the Filter Control Provider may contact for rules and other configuration information.
//   - [NEFilterProviderConfiguration.SetServerAddress]
//   - [NEFilterProviderConfiguration.Username]: A string that identifies the user.
//   - [NEFilterProviderConfiguration.SetUsername]
//   - [NEFilterProviderConfiguration.Organization]: A string that identifies the organization that administers the filter.
//   - [NEFilterProviderConfiguration.SetOrganization]
//   - [NEFilterProviderConfiguration.PasswordReference]: A persistent reference to a keychain item containing a password associated with the filter.
//   - [NEFilterProviderConfiguration.SetPasswordReference]
//   - [NEFilterProviderConfiguration.IdentityReference]: A persistent reference to a keychain item containing a certificate and private key associated with the filter.
//   - [NEFilterProviderConfiguration.SetIdentityReference]
//
// # Accessing bundle identifiers
//
//   - [NEFilterProviderConfiguration.FilterDataProviderBundleIdentifier]: The bundle identifier of the filter data provider system extension.
//   - [NEFilterProviderConfiguration.SetFilterDataProviderBundleIdentifier]
//   - [NEFilterProviderConfiguration.FilterPacketProviderBundleIdentifier]: The bundle identifier of the filter packet provider system extension.
//   - [NEFilterProviderConfiguration.SetFilterPacketProviderBundleIdentifier]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration
type NEFilterProviderConfiguration struct {
	objectivec.Object
}

// NEFilterProviderConfigurationFromID constructs a [NEFilterProviderConfiguration] from an objc.ID.
//
// Configuration parameters for a content filter.
func NEFilterProviderConfigurationFromID(id objc.ID) NEFilterProviderConfiguration {
	return NEFilterProviderConfiguration{objectivec.Object{ID: id}}
}

// NOTE: NEFilterProviderConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterProviderConfiguration] class.
//
// # Configuring filter behavior
//
//   - [INEFilterProviderConfiguration.FilterSockets]: A Boolean value that indicates that the system applies the filter to flows of network data originated from sockets.
//   - [INEFilterProviderConfiguration.SetFilterSockets]
//   - [INEFilterProviderConfiguration.FilterPackets]: A Boolean value that indicates that the system applies the filter to packets of network data.
//   - [INEFilterProviderConfiguration.SetFilterPackets]
//
// # Accessing the filter configuration
//
//   - [INEFilterProviderConfiguration.VendorConfiguration]: A dictionary of provider-specific configuration settings.
//   - [INEFilterProviderConfiguration.SetVendorConfiguration]
//   - [INEFilterProviderConfiguration.ServerAddress]: The address of a server that the Filter Control Provider may contact for rules and other configuration information.
//   - [INEFilterProviderConfiguration.SetServerAddress]
//   - [INEFilterProviderConfiguration.Username]: A string that identifies the user.
//   - [INEFilterProviderConfiguration.SetUsername]
//   - [INEFilterProviderConfiguration.Organization]: A string that identifies the organization that administers the filter.
//   - [INEFilterProviderConfiguration.SetOrganization]
//   - [INEFilterProviderConfiguration.PasswordReference]: A persistent reference to a keychain item containing a password associated with the filter.
//   - [INEFilterProviderConfiguration.SetPasswordReference]
//   - [INEFilterProviderConfiguration.IdentityReference]: A persistent reference to a keychain item containing a certificate and private key associated with the filter.
//   - [INEFilterProviderConfiguration.SetIdentityReference]
//
// # Accessing bundle identifiers
//
//   - [INEFilterProviderConfiguration.FilterDataProviderBundleIdentifier]: The bundle identifier of the filter data provider system extension.
//   - [INEFilterProviderConfiguration.SetFilterDataProviderBundleIdentifier]
//   - [INEFilterProviderConfiguration.FilterPacketProviderBundleIdentifier]: The bundle identifier of the filter packet provider system extension.
//   - [INEFilterProviderConfiguration.SetFilterPacketProviderBundleIdentifier]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration
type INEFilterProviderConfiguration interface {
	objectivec.IObject

	// Topic: Configuring filter behavior

	// A Boolean value that indicates that the system applies the filter to flows of network data originated from sockets.
	FilterSockets() bool
	SetFilterSockets(value bool)
	// A Boolean value that indicates that the system applies the filter to packets of network data.
	FilterPackets() bool
	SetFilterPackets(value bool)

	// Topic: Accessing the filter configuration

	// A dictionary of provider-specific configuration settings.
	VendorConfiguration() foundation.INSDictionary
	SetVendorConfiguration(value foundation.INSDictionary)
	// The address of a server that the Filter Control Provider may contact for rules and other configuration information.
	ServerAddress() string
	SetServerAddress(value string)
	// A string that identifies the user.
	Username() string
	SetUsername(value string)
	// A string that identifies the organization that administers the filter.
	Organization() string
	SetOrganization(value string)
	// A persistent reference to a keychain item containing a password associated with the filter.
	PasswordReference() foundation.INSData
	SetPasswordReference(value foundation.INSData)
	// A persistent reference to a keychain item containing a certificate and private key associated with the filter.
	IdentityReference() foundation.INSData
	SetIdentityReference(value foundation.INSData)

	// Topic: Accessing bundle identifiers

	// The bundle identifier of the filter data provider system extension.
	FilterDataProviderBundleIdentifier() string
	SetFilterDataProviderBundleIdentifier(value string)
	// The bundle identifier of the filter packet provider system extension.
	FilterPacketProviderBundleIdentifier() string
	SetFilterPacketProviderBundleIdentifier(value string)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NEFilterProviderConfiguration) Init() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterProviderConfiguration) Autorelease() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterProviderConfiguration creates a new NEFilterProviderConfiguration instance.
func NewNEFilterProviderConfiguration() NEFilterProviderConfiguration {
	class := getNEFilterProviderConfigurationClass()
	rv := objc.Send[NEFilterProviderConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (f NEFilterProviderConfiguration) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value that indicates that the system applies the filter to flows
// of network data originated from sockets.
//
// # Discussion
//
// The default value of this property is false.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterSockets
func (f NEFilterProviderConfiguration) FilterSockets() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("filterSockets"))
	return rv
}
func (f NEFilterProviderConfiguration) SetFilterSockets(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setFilterSockets:"), value)
}

// A Boolean value that indicates that the system applies the filter to
// packets of network data.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterPackets
func (f NEFilterProviderConfiguration) FilterPackets() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("filterPackets"))
	return rv
}
func (f NEFilterProviderConfiguration) SetFilterPackets(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setFilterPackets:"), value)
}

// A dictionary of provider-specific configuration settings.
//
// # Discussion
//
// All of the values in this dictionary must be [NSSecureCoding]-compliant.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/vendorConfiguration
//
// [NSSecureCoding]: https://developer.apple.com/documentation/Foundation/NSSecureCoding
func (f NEFilterProviderConfiguration) VendorConfiguration() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("vendorConfiguration"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (f NEFilterProviderConfiguration) SetVendorConfiguration(value foundation.INSDictionary) {
	objc.Send[struct{}](f.ID, objc.Sel("setVendorConfiguration:"), value)
}

// The address of a server that the Filter Control Provider may contact for
// rules and other configuration information.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/serverAddress
func (f NEFilterProviderConfiguration) ServerAddress() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("serverAddress"))
	return foundation.NSStringFromID(rv).String()
}
func (f NEFilterProviderConfiguration) SetServerAddress(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setServerAddress:"), objc.String(value))
}

// A string that identifies the user.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/username
func (f NEFilterProviderConfiguration) Username() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("username"))
	return foundation.NSStringFromID(rv).String()
}
func (f NEFilterProviderConfiguration) SetUsername(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setUsername:"), objc.String(value))
}

// A string that identifies the organization that administers the filter.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/organization
func (f NEFilterProviderConfiguration) Organization() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("organization"))
	return foundation.NSStringFromID(rv).String()
}
func (f NEFilterProviderConfiguration) SetOrganization(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setOrganization:"), objc.String(value))
}

// A persistent reference to a keychain item containing a password associated
// with the filter.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/passwordReference
func (f NEFilterProviderConfiguration) PasswordReference() foundation.INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("passwordReference"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (f NEFilterProviderConfiguration) SetPasswordReference(value foundation.INSData) {
	objc.Send[struct{}](f.ID, objc.Sel("setPasswordReference:"), value)
}

// A persistent reference to a keychain item containing a certificate and
// private key associated with the filter.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/identityReference
func (f NEFilterProviderConfiguration) IdentityReference() foundation.INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("identityReference"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (f NEFilterProviderConfiguration) SetIdentityReference(value foundation.INSData) {
	objc.Send[struct{}](f.ID, objc.Sel("setIdentityReference:"), value)
}

// The bundle identifier of the filter data provider system extension.
//
// # Discussion
//
// If this property is `nil`, then the framework uses the bundle identifier of
// the [NEFilterDataProvider] extension in the calling app’s bundle. In this
// case, make sure the calling app’s bundle contains only one
// [NEFilterDataProvider], so there’s no ambiguity about which one to use.
//
// This property only applies to system extensions, since macOS doesn’t
// support implementing a filter data provider as an app extension.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterDataProviderBundleIdentifier
func (f NEFilterProviderConfiguration) FilterDataProviderBundleIdentifier() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("filterDataProviderBundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (f NEFilterProviderConfiguration) SetFilterDataProviderBundleIdentifier(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setFilterDataProviderBundleIdentifier:"), objc.String(value))
}

// The bundle identifier of the filter packet provider system extension.
//
// # Discussion
//
// If this property is `nil`, then the framework uses the bundle identifier of
// the [NEFilterPacketProvider] extension in the calling app’s bundle. In
// this case, make sure the calling app’s bundle contains only one
// [NEFilterPacketProvider], so there’s no ambiguity about which one to use.
//
// This property only applies to system extensions, since macOS doesn’t
// support implementing a filter packet provider as an app extension.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterPacketProviderBundleIdentifier
func (f NEFilterProviderConfiguration) FilterPacketProviderBundleIdentifier() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("filterPacketProviderBundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (f NEFilterProviderConfiguration) SetFilterPacketProviderBundleIdentifier(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setFilterPacketProviderBundleIdentifier:"), objc.String(value))
}

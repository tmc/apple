// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEVPNIKEv2PPKConfiguration] class.
var (
	_NEVPNIKEv2PPKConfigurationClass     NEVPNIKEv2PPKConfigurationClass
	_NEVPNIKEv2PPKConfigurationClassOnce sync.Once
)

func getNEVPNIKEv2PPKConfigurationClass() NEVPNIKEv2PPKConfigurationClass {
	_NEVPNIKEv2PPKConfigurationClassOnce.Do(func() {
		_NEVPNIKEv2PPKConfigurationClass = NEVPNIKEv2PPKConfigurationClass{class: objc.GetClass("NEVPNIKEv2PPKConfiguration")}
	})
	return _NEVPNIKEv2PPKConfigurationClass
}

// GetNEVPNIKEv2PPKConfigurationClass returns the class object for NEVPNIKEv2PPKConfiguration.
func GetNEVPNIKEv2PPKConfigurationClass() NEVPNIKEv2PPKConfigurationClass {
	return getNEVPNIKEv2PPKConfigurationClass()
}

type NEVPNIKEv2PPKConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEVPNIKEv2PPKConfigurationClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEVPNIKEv2PPKConfigurationClass) Alloc() NEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A class that manages parameters of a post-quantum pre-shared key (PPK).
//
// # Discussion
//
// Instances of this class are thread safe. The class conforms to RFC 8784.
//
// # Creating a PPK configuration
//
//   - [NEVPNIKEv2PPKConfiguration.InitWithIdentifierKeychainReference]: Initializes a quantum-secure pre-shared key (PPK) configuration.
//
// # Accessing the configuration parameters
//
//   - [NEVPNIKEv2PPKConfiguration.Identifier]: The identifier for the PPK.
//   - [NEVPNIKEv2PPKConfiguration.KeychainReference]: A persistent reference to the key in the keychain.
//   - [NEVPNIKEv2PPKConfiguration.IsMandatory]: A Boolean value that indicates whether it’s mandatory for the VPN server to use this PPK.
//   - [NEVPNIKEv2PPKConfiguration.SetIsMandatory]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration
type NEVPNIKEv2PPKConfiguration struct {
	objectivec.Object
}

// NEVPNIKEv2PPKConfigurationFromID constructs a [NEVPNIKEv2PPKConfiguration] from an objc.ID.
//
// A class that manages parameters of a post-quantum pre-shared key (PPK).
func NEVPNIKEv2PPKConfigurationFromID(id objc.ID) NEVPNIKEv2PPKConfiguration {
	return NEVPNIKEv2PPKConfiguration{objectivec.Object{ID: id}}
}

// NOTE: NEVPNIKEv2PPKConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEVPNIKEv2PPKConfiguration] class.
//
// # Creating a PPK configuration
//
//   - [INEVPNIKEv2PPKConfiguration.InitWithIdentifierKeychainReference]: Initializes a quantum-secure pre-shared key (PPK) configuration.
//
// # Accessing the configuration parameters
//
//   - [INEVPNIKEv2PPKConfiguration.Identifier]: The identifier for the PPK.
//   - [INEVPNIKEv2PPKConfiguration.KeychainReference]: A persistent reference to the key in the keychain.
//   - [INEVPNIKEv2PPKConfiguration.IsMandatory]: A Boolean value that indicates whether it’s mandatory for the VPN server to use this PPK.
//   - [INEVPNIKEv2PPKConfiguration.SetIsMandatory]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration
type INEVPNIKEv2PPKConfiguration interface {
	objectivec.IObject

	// Topic: Creating a PPK configuration

	// Initializes a quantum-secure pre-shared key (PPK) configuration.
	InitWithIdentifierKeychainReference(identifier string, keychainReference foundation.INSData) NEVPNIKEv2PPKConfiguration

	// Topic: Accessing the configuration parameters

	// The identifier for the PPK.
	Identifier() string
	// A persistent reference to the key in the keychain.
	KeychainReference() foundation.INSData
	// A Boolean value that indicates whether it’s mandatory for the VPN server to use this PPK.
	IsMandatory() bool
	SetIsMandatory(value bool)

	// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
	AllowPostQuantumKeyExchangeFallback() bool
	SetAllowPostQuantumKeyExchangeFallback(value bool)
	// The configuration for a post-quantum pre-shared key (PPK).
	PpkConfiguration() INEVPNIKEv2PPKConfiguration
	SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration)
}

// Init initializes the instance.
func (v NEVPNIKEv2PPKConfiguration) Init() NEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NEVPNIKEv2PPKConfiguration) Autorelease() NEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNIKEv2PPKConfiguration creates a new NEVPNIKEv2PPKConfiguration instance.
func NewNEVPNIKEv2PPKConfiguration() NEVPNIKEv2PPKConfiguration {
	class := getNEVPNIKEv2PPKConfigurationClass()
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a quantum-secure pre-shared key (PPK) configuration.
//
// identifier: The identifier for the PPK.
//
// keychainReference: A persistent reference to a keychain item with the class
// [kSecClassGenericPassword] that contains the PPK.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/init(identifier:keychainReference:)
//
// [kSecClassGenericPassword]: https://developer.apple.com/documentation/Security/kSecClassGenericPassword
func NewVPNIKEv2PPKConfigurationWithIdentifierKeychainReference(identifier string, keychainReference foundation.INSData) NEVPNIKEv2PPKConfiguration {
	instance := getNEVPNIKEv2PPKConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:keychainReference:"), objc.String(identifier), keychainReference)
	return NEVPNIKEv2PPKConfigurationFromID(rv)
}

// Initializes a quantum-secure pre-shared key (PPK) configuration.
//
// identifier: The identifier for the PPK.
//
// keychainReference: A persistent reference to a keychain item with the class
// [kSecClassGenericPassword] that contains the PPK.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/init(identifier:keychainReference:)
//
// [kSecClassGenericPassword]: https://developer.apple.com/documentation/Security/kSecClassGenericPassword
func (v NEVPNIKEv2PPKConfiguration) InitWithIdentifierKeychainReference(identifier string, keychainReference foundation.INSData) NEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](v.ID, objc.Sel("initWithIdentifier:keychainReference:"), objc.String(identifier), keychainReference)
	return rv
}

// The identifier for the PPK.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/identifier
func (v NEVPNIKEv2PPKConfiguration) Identifier() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// A persistent reference to the key in the keychain.
//
// # Discussion
//
// The keychain item needs to have the class [kSecClassGenericPassword].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/keychainReference
//
// [kSecClassGenericPassword]: https://developer.apple.com/documentation/Security/kSecClassGenericPassword
func (v NEVPNIKEv2PPKConfiguration) KeychainReference() foundation.INSData {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("keychainReference"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// A Boolean value that indicates whether it’s mandatory for the VPN server
// to use this PPK.
//
// # Discussion
//
// The default value is `true`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/isMandatory
func (v NEVPNIKEv2PPKConfiguration) IsMandatory() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isMandatory"))
	return rv
}
func (v NEVPNIKEv2PPKConfiguration) SetIsMandatory(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setIsMandatory:"), value)
}

// A Boolean value that indicates whether servers that don’t support
// post-quantum key exchanges can skip them.
//
// See: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/allowpostquantumkeyexchangefallback
func (v NEVPNIKEv2PPKConfiguration) AllowPostQuantumKeyExchangeFallback() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("allowPostQuantumKeyExchangeFallback"))
	return rv
}
func (v NEVPNIKEv2PPKConfiguration) SetAllowPostQuantumKeyExchangeFallback(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setAllowPostQuantumKeyExchangeFallback:"), value)
}

// The configuration for a post-quantum pre-shared key (PPK).
//
// See: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ppkconfiguration
func (v NEVPNIKEv2PPKConfiguration) PpkConfiguration() INEVPNIKEv2PPKConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("ppkConfiguration"))
	return NEVPNIKEv2PPKConfigurationFromID(objc.ID(rv))
}
func (v NEVPNIKEv2PPKConfiguration) SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setPpkConfiguration:"), value)
}

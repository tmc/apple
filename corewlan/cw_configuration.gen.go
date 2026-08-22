// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CWConfiguration] class.
var (
	_CWConfigurationClass     CWConfigurationClass
	_CWConfigurationClassOnce sync.Once
)

func getCWConfigurationClass() CWConfigurationClass {
	_CWConfigurationClassOnce.Do(func() {
		_CWConfigurationClass = CWConfigurationClass{class: objc.GetClass("CWConfiguration")}
	})
	return _CWConfigurationClass
}

// GetCWConfigurationClass returns the class object for CWConfiguration.
func GetCWConfigurationClass() CWConfigurationClass {
	return getCWConfigurationClass()
}

type CWConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CWConfigurationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CWConfigurationClass) Alloc() CWConfiguration {
	rv := objc.Send[CWConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Encapsulates an immutable configuration for an AirPort WLAN interface.
//
// # Creating a configuration
//
//   - [CWConfiguration.InitWithConfiguration]: Creates and returns a CWConfiguration object initialized with the given CWConfiguration object.
//
// # Comparing configurations
//
//   - [CWConfiguration.IsEqualToConfiguration]: Determine CWConfiguration object equality.
//
// # Instance Properties
//
//   - [CWConfiguration.NetworkProfiles]: An array of remembered CWNetworkProfile objects.
//   - [CWConfiguration.RememberJoinedNetworks]: AirPort client will remember all joined networks.
//   - [CWConfiguration.RequireAdministratorForAssociation]: Require an administrator password to change networks.
//   - [CWConfiguration.RequireAdministratorForIBSSMode]: Require an administrator password to create a computer-to-computer network.
//   - [CWConfiguration.RequireAdministratorForPower]: Require an administrator password to change the interface power state.
//
// # Initializers
//
//   - [CWConfiguration.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration
type CWConfiguration struct {
	objectivec.Object
}

// CWConfigurationFromID constructs a [CWConfiguration] from an objc.ID.
//
// Encapsulates an immutable configuration for an AirPort WLAN interface.
func CWConfigurationFromID(id objc.ID) CWConfiguration {
	return CWConfiguration{objectivec.Object{ID: id}}
}

// NOTE: CWConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CWConfiguration] class.
//
// # Creating a configuration
//
//   - [ICWConfiguration.InitWithConfiguration]: Creates and returns a CWConfiguration object initialized with the given CWConfiguration object.
//
// # Comparing configurations
//
//   - [ICWConfiguration.IsEqualToConfiguration]: Determine CWConfiguration object equality.
//
// # Instance Properties
//
//   - [ICWConfiguration.NetworkProfiles]: An array of remembered CWNetworkProfile objects.
//   - [ICWConfiguration.RememberJoinedNetworks]: AirPort client will remember all joined networks.
//   - [ICWConfiguration.RequireAdministratorForAssociation]: Require an administrator password to change networks.
//   - [ICWConfiguration.RequireAdministratorForIBSSMode]: Require an administrator password to create a computer-to-computer network.
//   - [ICWConfiguration.RequireAdministratorForPower]: Require an administrator password to change the interface power state.
//
// # Initializers
//
//   - [ICWConfiguration.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration
type ICWConfiguration interface {
	objectivec.IObject

	// Topic: Creating a configuration

	// Creates and returns a CWConfiguration object initialized with the given CWConfiguration object.
	InitWithConfiguration(configuration ICWConfiguration) CWConfiguration

	// Topic: Comparing configurations

	// Determine CWConfiguration object equality.
	IsEqualToConfiguration(configuration ICWConfiguration) bool

	// Topic: Instance Properties

	// An array of remembered CWNetworkProfile objects.
	NetworkProfiles() foundation.INSOrderedSet
	// AirPort client will remember all joined networks.
	RememberJoinedNetworks() bool
	// Require an administrator password to change networks.
	RequireAdministratorForAssociation() bool
	// Require an administrator password to create a computer-to-computer network.
	RequireAdministratorForIBSSMode() bool
	// Require an administrator password to change the interface power state.
	RequireAdministratorForPower() bool

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CWConfiguration

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CWConfiguration) Init() CWConfiguration {
	rv := objc.Send[CWConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CWConfiguration) Autorelease() CWConfiguration {
	rv := objc.Send[CWConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWConfiguration creates a new CWConfiguration instance.
func NewCWConfiguration() CWConfiguration {
	class := getCWConfigurationClass()
	rv := objc.Send[CWConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/init(coder:)
func NewCWConfigurationWithCoder(coder foundation.INSCoder) CWConfiguration {
	instance := getCWConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CWConfigurationFromID(rv)
}

// Creates and returns a CWConfiguration object initialized with the given
// CWConfiguration object.
//
// configuration: The CWConfiguration object to use to initialize a new CWConfiguration
// object.
//
// # Return Value
//
// A CWConfiguration object.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/init(configuration:)
func NewCWConfigurationWithConfiguration(configuration ICWConfiguration) CWConfiguration {
	instance := getCWConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return CWConfigurationFromID(rv)
}

// Creates and returns a CWConfiguration object initialized with the given
// CWConfiguration object.
//
// configuration: The CWConfiguration object to use to initialize a new CWConfiguration
// object.
//
// # Return Value
//
// A CWConfiguration object.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/init(configuration:)
func (c CWConfiguration) InitWithConfiguration(configuration ICWConfiguration) CWConfiguration {
	rv := objc.Send[CWConfiguration](c.ID, objc.Sel("initWithConfiguration:"), configuration)
	return rv
}

// Determine CWConfiguration object equality.
//
// configuration: The CWConfiguration object with which to compare the receiver.
//
// # Return Value
//
// YES if the objects are equal.
//
// # Discussion
//
// CWConfiguration objects are considered equal if all their corresponding
// properties are equal.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/isEqual(to:)
func (c CWConfiguration) IsEqualToConfiguration(configuration ICWConfiguration) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEqualToConfiguration:"), configuration)
	return rv
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/init(coder:)
func (c CWConfiguration) InitWithCoder(coder foundation.INSCoder) CWConfiguration {
	rv := objc.Send[CWConfiguration](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CWConfiguration) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// An array of remembered CWNetworkProfile objects.
//
// # Discussion
//
// The order of this array corresponds to the order in which the the
// CWNetworkProfile objects participate in the auto-join process.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/networkProfiles
func (c CWConfiguration) NetworkProfiles() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("networkProfiles"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}

// AirPort client will remember all joined networks.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/rememberJoinedNetworks
func (c CWConfiguration) RememberJoinedNetworks() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("rememberJoinedNetworks"))
	return rv
}

// Require an administrator password to change networks.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/requireAdministratorForAssociation
func (c CWConfiguration) RequireAdministratorForAssociation() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("requireAdministratorForAssociation"))
	return rv
}

// Require an administrator password to create a computer-to-computer network.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/requireAdministratorForIBSSMode
func (c CWConfiguration) RequireAdministratorForIBSSMode() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("requireAdministratorForIBSSMode"))
	return rv
}

// Require an administrator password to change the interface power state.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/requireAdministratorForPower
func (c CWConfiguration) RequireAdministratorForPower() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("requireAdministratorForPower"))
	return rv
}

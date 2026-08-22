// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenDriverConfiguration] class.
var (
	_TKTokenDriverConfigurationClass     TKTokenDriverConfigurationClass
	_TKTokenDriverConfigurationClassOnce sync.Once
)

func getTKTokenDriverConfigurationClass() TKTokenDriverConfigurationClass {
	_TKTokenDriverConfigurationClassOnce.Do(func() {
		_TKTokenDriverConfigurationClass = TKTokenDriverConfigurationClass{class: objc.GetClass("TKTokenDriverConfiguration")}
	})
	return _TKTokenDriverConfigurationClass
}

// GetTKTokenDriverConfigurationClass returns the class object for TKTokenDriverConfiguration.
func GetTKTokenDriverConfigurationClass() TKTokenDriverConfigurationClass {
	return getTKTokenDriverConfigurationClass()
}

type TKTokenDriverConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenDriverConfigurationClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenDriverConfigurationClass) Alloc() TKTokenDriverConfiguration {
	rv := objc.Send[TKTokenDriverConfiguration](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A configuration for one class of token.
//
// # Reporting Configuration Information
//
//   - [TKTokenDriverConfiguration.ClassID]: The class identifier of the token driver.
//   - [TKTokenDriverConfiguration.TokenConfigurations]: A dictionary of all currently configured tokens for this token class, which the token instance identifier keys.
//
// # Adding and Removing Configurations
//
//   - [TKTokenDriverConfiguration.AddTokenConfigurationForTokenInstanceID]: Creates a configuration object for a token with the token instance identifier you specify.
//   - [TKTokenDriverConfiguration.RemoveTokenConfigurationForTokenInstanceID]: Removes a configuration for a token with the token instance identifier you specify.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration
type TKTokenDriverConfiguration struct {
	objectivec.Object
}

// TKTokenDriverConfigurationFromID constructs a [TKTokenDriverConfiguration] from an objc.ID.
//
// A configuration for one class of token.
func TKTokenDriverConfigurationFromID(id objc.ID) TKTokenDriverConfiguration {
	return TKTokenDriverConfiguration{objectivec.Object{ID: id}}
}

// NOTE: TKTokenDriverConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenDriverConfiguration] class.
//
// # Reporting Configuration Information
//
//   - [ITKTokenDriverConfiguration.ClassID]: The class identifier of the token driver.
//   - [ITKTokenDriverConfiguration.TokenConfigurations]: A dictionary of all currently configured tokens for this token class, which the token instance identifier keys.
//
// # Adding and Removing Configurations
//
//   - [ITKTokenDriverConfiguration.AddTokenConfigurationForTokenInstanceID]: Creates a configuration object for a token with the token instance identifier you specify.
//   - [ITKTokenDriverConfiguration.RemoveTokenConfigurationForTokenInstanceID]: Removes a configuration for a token with the token instance identifier you specify.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration
type ITKTokenDriverConfiguration interface {
	objectivec.IObject

	// Topic: Reporting Configuration Information

	// The class identifier of the token driver.
	ClassID() TKTokenDriverClassID
	// A dictionary of all currently configured tokens for this token class, which the token instance identifier keys.
	TokenConfigurations() foundation.INSDictionary

	// Topic: Adding and Removing Configurations

	// Creates a configuration object for a token with the token instance identifier you specify.
	AddTokenConfigurationForTokenInstanceID(instanceID TKTokenInstanceID) ITKTokenConfiguration
	// Removes a configuration for a token with the token instance identifier you specify.
	RemoveTokenConfigurationForTokenInstanceID(instanceID TKTokenInstanceID)
}

// Init initializes the instance.
func (t TKTokenDriverConfiguration) Init() TKTokenDriverConfiguration {
	rv := objc.Send[TKTokenDriverConfiguration](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenDriverConfiguration) Autorelease() TKTokenDriverConfiguration {
	rv := objc.Send[TKTokenDriverConfiguration](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenDriverConfiguration creates a new TKTokenDriverConfiguration instance.
func NewTKTokenDriverConfiguration() TKTokenDriverConfiguration {
	class := getTKTokenDriverConfigurationClass()
	rv := objc.Send[TKTokenDriverConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a configuration object for a token with the token instance
// identifier you specify.
//
// instanceID: The token’s instance identifier.
//
// # Return Value
//
// The configuration class for the token.
//
// # Discussion
//
// This method adds the created configuration into the
// [TKTokenDriverConfiguration.TokenConfigurations] dictionary. Adding a
// configuration with an `instanceID` that already exists replaces the
// existing configuration with a new empty configuration.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/addTokenConfiguration(for:)
func (t TKTokenDriverConfiguration) AddTokenConfigurationForTokenInstanceID(instanceID TKTokenInstanceID) ITKTokenConfiguration {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("addTokenConfigurationForTokenInstanceID:"), objc.String(string(instanceID)))
	return TKTokenConfigurationFromID(rv)
}

// Removes a configuration for a token with the token instance identifier you
// specify.
//
// instanceID: The token’s instance identifier.
//
// # Discussion
//
// The method does nothing if the token configuration you specify doesn’t
// exist.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/removeTokenConfiguration(for:)
func (t TKTokenDriverConfiguration) RemoveTokenConfigurationForTokenInstanceID(instanceID TKTokenInstanceID) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeTokenConfigurationForTokenInstanceID:"), objc.String(string(instanceID)))
}

// The class identifier of the token driver.
//
// # Discussion
//
// The identifier uses a reverse-DNS format, such as
// `com.ExampleXCUIElementTypeId()`.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/classID
func (t TKTokenDriverConfiguration) ClassID() TKTokenDriverClassID {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("classID"))
	return TKTokenDriverClassID(foundation.NSStringFromID(rv).String())
}

// A dictionary of all currently configured tokens for this token class, which
// the token instance identifier keys.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/tokenConfigurations
func (t TKTokenDriverConfiguration) TokenConfigurations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tokenConfigurations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A dictionary of token class configurations which the class identifier of
// the token driver keys.
//
// # Discussion
//
// If the app hosting the token extension calls this method, it returns a list
// of configurations for hosted token extensions. Otherwise, this method
// returns an empty array.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/driverConfigurations
func (_TKTokenDriverConfigurationClass TKTokenDriverConfigurationClass) DriverConfigurations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_TKTokenDriverConfigurationClass.class), objc.Sel("driverConfigurations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

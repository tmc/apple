// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CWMutableConfiguration] class.
var (
	_CWMutableConfigurationClass     CWMutableConfigurationClass
	_CWMutableConfigurationClassOnce sync.Once
)

func getCWMutableConfigurationClass() CWMutableConfigurationClass {
	_CWMutableConfigurationClassOnce.Do(func() {
		_CWMutableConfigurationClass = CWMutableConfigurationClass{class: objc.GetClass("CWMutableConfiguration")}
	})
	return _CWMutableConfigurationClass
}

// GetCWMutableConfigurationClass returns the class object for CWMutableConfiguration.
func GetCWMutableConfigurationClass() CWMutableConfigurationClass {
	return getCWMutableConfigurationClass()
}

type CWMutableConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CWMutableConfigurationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CWMutableConfigurationClass) Alloc() CWMutableConfiguration {
	rv := objc.Send[CWMutableConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Encapsulates a mutable configuration for an AirPort WLAN interface.
//
// # Overview
//
// Use this class to change configuration settings or the preferred networks
// list. To commit configuration changes, use
// [CWInterface.CommitConfigurationAuthorizationError].
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration
type CWMutableConfiguration struct {
	CWConfiguration
}

// CWMutableConfigurationFromID constructs a [CWMutableConfiguration] from an objc.ID.
//
// Encapsulates a mutable configuration for an AirPort WLAN interface.
func CWMutableConfigurationFromID(id objc.ID) CWMutableConfiguration {
	return CWMutableConfiguration{CWConfiguration: CWConfigurationFromID(id)}
}

// NOTE: CWMutableConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CWMutableConfiguration] class.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration
type ICWMutableConfiguration interface {
	ICWConfiguration
}

// Init initializes the instance.
func (c CWMutableConfiguration) Init() CWMutableConfiguration {
	rv := objc.Send[CWMutableConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CWMutableConfiguration) Autorelease() CWMutableConfiguration {
	rv := objc.Send[CWMutableConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWMutableConfiguration creates a new CWMutableConfiguration instance.
func NewCWMutableConfiguration() CWMutableConfiguration {
	class := getCWMutableConfigurationClass()
	rv := objc.Send[CWMutableConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/init(coder:)
func NewCWMutableConfigurationWithCoder(coder foundation.INSCoder) CWMutableConfiguration {
	instance := getCWMutableConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CWMutableConfigurationFromID(rv)
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
func NewCWMutableConfigurationWithConfiguration(configuration ICWConfiguration) CWMutableConfiguration {
	instance := getCWMutableConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return CWMutableConfigurationFromID(rv)
}

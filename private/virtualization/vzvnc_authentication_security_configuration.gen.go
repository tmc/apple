// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVNCAuthenticationSecurityConfiguration] class.
var (
	_VZVNCAuthenticationSecurityConfigurationClass     VZVNCAuthenticationSecurityConfigurationClass
	_VZVNCAuthenticationSecurityConfigurationClassOnce sync.Once
)

func getVZVNCAuthenticationSecurityConfigurationClass() VZVNCAuthenticationSecurityConfigurationClass {
	_VZVNCAuthenticationSecurityConfigurationClassOnce.Do(func() {
		_VZVNCAuthenticationSecurityConfigurationClass = VZVNCAuthenticationSecurityConfigurationClass{class: objc.GetClass("_VZVNCAuthenticationSecurityConfiguration")}
	})
	return _VZVNCAuthenticationSecurityConfigurationClass
}

// GetVZVNCAuthenticationSecurityConfigurationClass returns the class object for _VZVNCAuthenticationSecurityConfiguration.
func GetVZVNCAuthenticationSecurityConfigurationClass() VZVNCAuthenticationSecurityConfigurationClass {
	return getVZVNCAuthenticationSecurityConfigurationClass()
}

type VZVNCAuthenticationSecurityConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVNCAuthenticationSecurityConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVNCAuthenticationSecurityConfigurationClass) Alloc() VZVNCAuthenticationSecurityConfiguration {
	rv := objc.Send[VZVNCAuthenticationSecurityConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVNCAuthenticationSecurityConfiguration.Password]
//   - [VZVNCAuthenticationSecurityConfiguration.InitWithPassword]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVNCAuthenticationSecurityConfiguration
type VZVNCAuthenticationSecurityConfiguration struct {
	VZVNCSecurityConfiguration
}

// VZVNCAuthenticationSecurityConfigurationFromID constructs a [VZVNCAuthenticationSecurityConfiguration] from an objc.ID.
func VZVNCAuthenticationSecurityConfigurationFromID(id objc.ID) VZVNCAuthenticationSecurityConfiguration {
	return VZVNCAuthenticationSecurityConfiguration{VZVNCSecurityConfiguration: VZVNCSecurityConfigurationFromID(id)}
}

// Ensure VZVNCAuthenticationSecurityConfiguration implements IVZVNCAuthenticationSecurityConfiguration.
var _ IVZVNCAuthenticationSecurityConfiguration = VZVNCAuthenticationSecurityConfiguration{}

// An interface definition for the [VZVNCAuthenticationSecurityConfiguration] class.
//
// # Methods
//
//   - [IVZVNCAuthenticationSecurityConfiguration.Password]
//   - [IVZVNCAuthenticationSecurityConfiguration.InitWithPassword]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVNCAuthenticationSecurityConfiguration
type IVZVNCAuthenticationSecurityConfiguration interface {
	IVZVNCSecurityConfiguration

	// Topic: Methods

	Password() string
	InitWithPassword(password objectivec.IObject) VZVNCAuthenticationSecurityConfiguration
}

// Init initializes the instance.
func (v VZVNCAuthenticationSecurityConfiguration) Init() VZVNCAuthenticationSecurityConfiguration {
	rv := objc.Send[VZVNCAuthenticationSecurityConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVNCAuthenticationSecurityConfiguration) Autorelease() VZVNCAuthenticationSecurityConfiguration {
	rv := objc.Send[VZVNCAuthenticationSecurityConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVNCAuthenticationSecurityConfiguration creates a new VZVNCAuthenticationSecurityConfiguration instance.
func NewVZVNCAuthenticationSecurityConfiguration() VZVNCAuthenticationSecurityConfiguration {
	class := getVZVNCAuthenticationSecurityConfigurationClass()
	rv := objc.Send[VZVNCAuthenticationSecurityConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVNCAuthenticationSecurityConfiguration/initWithPassword:
func NewVZVNCAuthenticationSecurityConfigurationWithPassword(password objectivec.IObject) VZVNCAuthenticationSecurityConfiguration {
	instance := getVZVNCAuthenticationSecurityConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPassword:"), password)
	return VZVNCAuthenticationSecurityConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVNCAuthenticationSecurityConfiguration/initWithPassword:
func (v VZVNCAuthenticationSecurityConfiguration) InitWithPassword(password objectivec.IObject) VZVNCAuthenticationSecurityConfiguration {
	rv := objc.Send[VZVNCAuthenticationSecurityConfiguration](v.ID, objc.Sel("initWithPassword:"), password)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVNCAuthenticationSecurityConfiguration/password
func (v VZVNCAuthenticationSecurityConfiguration) Password() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("password"))
	return foundation.NSStringFromID(rv).String()
}

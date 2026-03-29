// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVNCNoSecuritySecurityConfiguration] class.
var (
	_VZVNCNoSecuritySecurityConfigurationClass     VZVNCNoSecuritySecurityConfigurationClass
	_VZVNCNoSecuritySecurityConfigurationClassOnce sync.Once
)

func getVZVNCNoSecuritySecurityConfigurationClass() VZVNCNoSecuritySecurityConfigurationClass {
	_VZVNCNoSecuritySecurityConfigurationClassOnce.Do(func() {
		_VZVNCNoSecuritySecurityConfigurationClass = VZVNCNoSecuritySecurityConfigurationClass{class: objc.GetClass("_VZVNCNoSecuritySecurityConfiguration")}
	})
	return _VZVNCNoSecuritySecurityConfigurationClass
}

// GetVZVNCNoSecuritySecurityConfigurationClass returns the class object for _VZVNCNoSecuritySecurityConfiguration.
func GetVZVNCNoSecuritySecurityConfigurationClass() VZVNCNoSecuritySecurityConfigurationClass {
	return getVZVNCNoSecuritySecurityConfigurationClass()
}

type VZVNCNoSecuritySecurityConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVNCNoSecuritySecurityConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVNCNoSecuritySecurityConfigurationClass) Alloc() VZVNCNoSecuritySecurityConfiguration {
	rv := objc.Send[VZVNCNoSecuritySecurityConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVNCNoSecuritySecurityConfiguration
type VZVNCNoSecuritySecurityConfiguration struct {
	VZVNCSecurityConfiguration
}

// VZVNCNoSecuritySecurityConfigurationFromID constructs a [VZVNCNoSecuritySecurityConfiguration] from an objc.ID.
func VZVNCNoSecuritySecurityConfigurationFromID(id objc.ID) VZVNCNoSecuritySecurityConfiguration {
	return VZVNCNoSecuritySecurityConfiguration{VZVNCSecurityConfiguration: VZVNCSecurityConfigurationFromID(id)}
}
// Ensure VZVNCNoSecuritySecurityConfiguration implements IVZVNCNoSecuritySecurityConfiguration.
var _ IVZVNCNoSecuritySecurityConfiguration = VZVNCNoSecuritySecurityConfiguration{}

// An interface definition for the [VZVNCNoSecuritySecurityConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVNCNoSecuritySecurityConfiguration
type IVZVNCNoSecuritySecurityConfiguration interface {
	IVZVNCSecurityConfiguration
}

// Init initializes the instance.
func (v VZVNCNoSecuritySecurityConfiguration) Init() VZVNCNoSecuritySecurityConfiguration {
	rv := objc.Send[VZVNCNoSecuritySecurityConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVNCNoSecuritySecurityConfiguration) Autorelease() VZVNCNoSecuritySecurityConfiguration {
	rv := objc.Send[VZVNCNoSecuritySecurityConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVNCNoSecuritySecurityConfiguration creates a new VZVNCNoSecuritySecurityConfiguration instance.
func NewVZVNCNoSecuritySecurityConfiguration() VZVNCNoSecuritySecurityConfiguration {
	class := getVZVNCNoSecuritySecurityConfigurationClass()
	rv := objc.Send[VZVNCNoSecuritySecurityConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}


// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVNCSecurityConfiguration] class.
var (
	_VZVNCSecurityConfigurationClass     VZVNCSecurityConfigurationClass
	_VZVNCSecurityConfigurationClassOnce sync.Once
)

func getVZVNCSecurityConfigurationClass() VZVNCSecurityConfigurationClass {
	_VZVNCSecurityConfigurationClassOnce.Do(func() {
		_VZVNCSecurityConfigurationClass = VZVNCSecurityConfigurationClass{class: objc.GetClass("_VZVNCSecurityConfiguration")}
	})
	return _VZVNCSecurityConfigurationClass
}

// GetVZVNCSecurityConfigurationClass returns the class object for _VZVNCSecurityConfiguration.
func GetVZVNCSecurityConfigurationClass() VZVNCSecurityConfigurationClass {
	return getVZVNCSecurityConfigurationClass()
}

type VZVNCSecurityConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVNCSecurityConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVNCSecurityConfigurationClass) Alloc() VZVNCSecurityConfiguration {
	rv := objc.Send[VZVNCSecurityConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVNCSecurityConfiguration._init]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVNCSecurityConfiguration
type VZVNCSecurityConfiguration struct {
	objectivec.Object
}

// VZVNCSecurityConfigurationFromID constructs a [VZVNCSecurityConfiguration] from an objc.ID.
func VZVNCSecurityConfigurationFromID(id objc.ID) VZVNCSecurityConfiguration {
	return VZVNCSecurityConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZVNCSecurityConfiguration implements IVZVNCSecurityConfiguration.
var _ IVZVNCSecurityConfiguration = VZVNCSecurityConfiguration{}

// An interface definition for the [VZVNCSecurityConfiguration] class.
//
// # Methods
//
//   - [IVZVNCSecurityConfiguration._init]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVNCSecurityConfiguration
type IVZVNCSecurityConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
}

// Init initializes the instance.
func (v VZVNCSecurityConfiguration) Init() VZVNCSecurityConfiguration {
	rv := objc.Send[VZVNCSecurityConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVNCSecurityConfiguration) Autorelease() VZVNCSecurityConfiguration {
	rv := objc.Send[VZVNCSecurityConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVNCSecurityConfiguration creates a new VZVNCSecurityConfiguration instance.
func NewVZVNCSecurityConfiguration() VZVNCSecurityConfiguration {
	class := getVZVNCSecurityConfigurationClass()
	rv := objc.Send[VZVNCSecurityConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVNCSecurityConfiguration/_init
func (v VZVNCSecurityConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacRemoteServiceDiscoveryConfiguration] class.
var (
	_VZMacRemoteServiceDiscoveryConfigurationClass     VZMacRemoteServiceDiscoveryConfigurationClass
	_VZMacRemoteServiceDiscoveryConfigurationClassOnce sync.Once
)

func getVZMacRemoteServiceDiscoveryConfigurationClass() VZMacRemoteServiceDiscoveryConfigurationClass {
	_VZMacRemoteServiceDiscoveryConfigurationClassOnce.Do(func() {
		_VZMacRemoteServiceDiscoveryConfigurationClass = VZMacRemoteServiceDiscoveryConfigurationClass{class: objc.GetClass("_VZMacRemoteServiceDiscoveryConfiguration")}
	})
	return _VZMacRemoteServiceDiscoveryConfigurationClass
}

// GetVZMacRemoteServiceDiscoveryConfigurationClass returns the class object for _VZMacRemoteServiceDiscoveryConfiguration.
func GetVZMacRemoteServiceDiscoveryConfigurationClass() VZMacRemoteServiceDiscoveryConfigurationClass {
	return getVZMacRemoteServiceDiscoveryConfigurationClass()
}

type VZMacRemoteServiceDiscoveryConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacRemoteServiceDiscoveryConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacRemoteServiceDiscoveryConfigurationClass) Alloc() VZMacRemoteServiceDiscoveryConfiguration {
	rv := objc.Send[VZMacRemoteServiceDiscoveryConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMacRemoteServiceDiscoveryConfiguration.GuestServiceAllowList]
//   - [VZMacRemoteServiceDiscoveryConfiguration.SetGuestServiceAllowList]
//   - [VZMacRemoteServiceDiscoveryConfiguration.HostServiceAllowList]
//   - [VZMacRemoteServiceDiscoveryConfiguration.SetHostServiceAllowList]
// See: https://developer.apple.com/documentation/Virtualization/_VZMacRemoteServiceDiscoveryConfiguration
type VZMacRemoteServiceDiscoveryConfiguration struct {
	objectivec.Object
}

// VZMacRemoteServiceDiscoveryConfigurationFromID constructs a [VZMacRemoteServiceDiscoveryConfiguration] from an objc.ID.
func VZMacRemoteServiceDiscoveryConfigurationFromID(id objc.ID) VZMacRemoteServiceDiscoveryConfiguration {
	return VZMacRemoteServiceDiscoveryConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZMacRemoteServiceDiscoveryConfiguration implements IVZMacRemoteServiceDiscoveryConfiguration.
var _ IVZMacRemoteServiceDiscoveryConfiguration = VZMacRemoteServiceDiscoveryConfiguration{}

// An interface definition for the [VZMacRemoteServiceDiscoveryConfiguration] class.
//
// # Methods
//
//   - [IVZMacRemoteServiceDiscoveryConfiguration.GuestServiceAllowList]
//   - [IVZMacRemoteServiceDiscoveryConfiguration.SetGuestServiceAllowList]
//   - [IVZMacRemoteServiceDiscoveryConfiguration.HostServiceAllowList]
//   - [IVZMacRemoteServiceDiscoveryConfiguration.SetHostServiceAllowList]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacRemoteServiceDiscoveryConfiguration
type IVZMacRemoteServiceDiscoveryConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	GuestServiceAllowList() foundation.INSArray
	SetGuestServiceAllowList(value foundation.INSArray)
	HostServiceAllowList() foundation.INSArray
	SetHostServiceAllowList(value foundation.INSArray)
}

// Init initializes the instance.
func (v VZMacRemoteServiceDiscoveryConfiguration) Init() VZMacRemoteServiceDiscoveryConfiguration {
	rv := objc.Send[VZMacRemoteServiceDiscoveryConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacRemoteServiceDiscoveryConfiguration) Autorelease() VZMacRemoteServiceDiscoveryConfiguration {
	rv := objc.Send[VZMacRemoteServiceDiscoveryConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacRemoteServiceDiscoveryConfiguration creates a new VZMacRemoteServiceDiscoveryConfiguration instance.
func NewVZMacRemoteServiceDiscoveryConfiguration() VZMacRemoteServiceDiscoveryConfiguration {
	class := getVZMacRemoteServiceDiscoveryConfigurationClass()
	rv := objc.Send[VZMacRemoteServiceDiscoveryConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacRemoteServiceDiscoveryConfiguration/guestServiceAllowList
func (v VZMacRemoteServiceDiscoveryConfiguration) GuestServiceAllowList() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("guestServiceAllowList"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZMacRemoteServiceDiscoveryConfiguration) SetGuestServiceAllowList(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("setGuestServiceAllowList:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZMacRemoteServiceDiscoveryConfiguration/hostServiceAllowList
func (v VZMacRemoteServiceDiscoveryConfiguration) HostServiceAllowList() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("hostServiceAllowList"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZMacRemoteServiceDiscoveryConfiguration) SetHostServiceAllowList(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("setHostServiceAllowList:"), value)
}


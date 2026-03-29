// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZMacHIDDeviceConfiguration] class.
var (
	_VZMacHIDDeviceConfigurationClass     VZMacHIDDeviceConfigurationClass
	_VZMacHIDDeviceConfigurationClassOnce sync.Once
)

func getVZMacHIDDeviceConfigurationClass() VZMacHIDDeviceConfigurationClass {
	_VZMacHIDDeviceConfigurationClassOnce.Do(func() {
		_VZMacHIDDeviceConfigurationClass = VZMacHIDDeviceConfigurationClass{class: objc.GetClass("_VZMacHIDDeviceConfiguration")}
	})
	return _VZMacHIDDeviceConfigurationClass
}

// GetVZMacHIDDeviceConfigurationClass returns the class object for _VZMacHIDDeviceConfiguration.
func GetVZMacHIDDeviceConfigurationClass() VZMacHIDDeviceConfigurationClass {
	return getVZMacHIDDeviceConfigurationClass()
}

type VZMacHIDDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacHIDDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacHIDDeviceConfigurationClass) Alloc() VZMacHIDDeviceConfiguration {
	rv := objc.Send[VZMacHIDDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMacHIDDeviceConfiguration.ProductID]
//   - [VZMacHIDDeviceConfiguration.SetProductID]
//   - [VZMacHIDDeviceConfiguration.RegistryProperties]
//   - [VZMacHIDDeviceConfiguration.SetRegistryProperties]
//   - [VZMacHIDDeviceConfiguration.Usage]
//   - [VZMacHIDDeviceConfiguration.SetUsage]
//   - [VZMacHIDDeviceConfiguration.UsagePage]
//   - [VZMacHIDDeviceConfiguration.SetUsagePage]
//   - [VZMacHIDDeviceConfiguration.VendorID]
//   - [VZMacHIDDeviceConfiguration.SetVendorID]
//   - [VZMacHIDDeviceConfiguration.InitWithVendorIDProductIDUsagePageUsage]
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHIDDeviceConfiguration
type VZMacHIDDeviceConfiguration struct {
	VZHIDDeviceConfiguration
}

// VZMacHIDDeviceConfigurationFromID constructs a [VZMacHIDDeviceConfiguration] from an objc.ID.
func VZMacHIDDeviceConfigurationFromID(id objc.ID) VZMacHIDDeviceConfiguration {
	return VZMacHIDDeviceConfiguration{VZHIDDeviceConfiguration: VZHIDDeviceConfigurationFromID(id)}
}
// Ensure VZMacHIDDeviceConfiguration implements IVZMacHIDDeviceConfiguration.
var _ IVZMacHIDDeviceConfiguration = VZMacHIDDeviceConfiguration{}

// An interface definition for the [VZMacHIDDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZMacHIDDeviceConfiguration.ProductID]
//   - [IVZMacHIDDeviceConfiguration.SetProductID]
//   - [IVZMacHIDDeviceConfiguration.RegistryProperties]
//   - [IVZMacHIDDeviceConfiguration.SetRegistryProperties]
//   - [IVZMacHIDDeviceConfiguration.Usage]
//   - [IVZMacHIDDeviceConfiguration.SetUsage]
//   - [IVZMacHIDDeviceConfiguration.UsagePage]
//   - [IVZMacHIDDeviceConfiguration.SetUsagePage]
//   - [IVZMacHIDDeviceConfiguration.VendorID]
//   - [IVZMacHIDDeviceConfiguration.SetVendorID]
//   - [IVZMacHIDDeviceConfiguration.InitWithVendorIDProductIDUsagePageUsage]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHIDDeviceConfiguration
type IVZMacHIDDeviceConfiguration interface {
	IVZHIDDeviceConfiguration

	// Topic: Methods

	ProductID() uint16
	SetProductID(value uint16)
	RegistryProperties() foundation.INSDictionary
	SetRegistryProperties(value foundation.INSDictionary)
	Usage() uint32
	SetUsage(value uint32)
	UsagePage() uint32
	SetUsagePage(value uint32)
	VendorID() uint16
	SetVendorID(value uint16)
	InitWithVendorIDProductIDUsagePageUsage(id uint16, id2 uint16, page uint32, usage uint32) VZMacHIDDeviceConfiguration
}

// Init initializes the instance.
func (v VZMacHIDDeviceConfiguration) Init() VZMacHIDDeviceConfiguration {
	rv := objc.Send[VZMacHIDDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacHIDDeviceConfiguration) Autorelease() VZMacHIDDeviceConfiguration {
	rv := objc.Send[VZMacHIDDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacHIDDeviceConfiguration creates a new VZMacHIDDeviceConfiguration instance.
func NewVZMacHIDDeviceConfiguration() VZMacHIDDeviceConfiguration {
	class := getVZMacHIDDeviceConfigurationClass()
	rv := objc.Send[VZMacHIDDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHIDDeviceConfiguration/initWithVendorID:productID:usagePage:usage:
func NewVZMacHIDDeviceConfigurationWithVendorIDProductIDUsagePageUsage(id uint16, id2 uint16, page uint32, usage uint32) VZMacHIDDeviceConfiguration {
	instance := getVZMacHIDDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVendorID:productID:usagePage:usage:"), id, id2, page, usage)
	return VZMacHIDDeviceConfigurationFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHIDDeviceConfiguration/initWithVendorID:productID:usagePage:usage:
func (v VZMacHIDDeviceConfiguration) InitWithVendorIDProductIDUsagePageUsage(id uint16, id2 uint16, page uint32, usage uint32) VZMacHIDDeviceConfiguration {
	rv := objc.Send[VZMacHIDDeviceConfiguration](v.ID, objc.Sel("initWithVendorID:productID:usagePage:usage:"), id, id2, page, usage)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacHIDDeviceConfiguration/productID
func (v VZMacHIDDeviceConfiguration) ProductID() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("productID"))
	return rv
}
func (v VZMacHIDDeviceConfiguration) SetProductID(value uint16) {
	objc.Send[struct{}](v.ID, objc.Sel("setProductID:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHIDDeviceConfiguration/registryProperties
func (v VZMacHIDDeviceConfiguration) RegistryProperties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("registryProperties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v VZMacHIDDeviceConfiguration) SetRegistryProperties(value foundation.INSDictionary) {
	objc.Send[struct{}](v.ID, objc.Sel("setRegistryProperties:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHIDDeviceConfiguration/usage
func (v VZMacHIDDeviceConfiguration) Usage() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("usage"))
	return rv
}
func (v VZMacHIDDeviceConfiguration) SetUsage(value uint32) {
	objc.Send[struct{}](v.ID, objc.Sel("setUsage:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHIDDeviceConfiguration/usagePage
func (v VZMacHIDDeviceConfiguration) UsagePage() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("usagePage"))
	return rv
}
func (v VZMacHIDDeviceConfiguration) SetUsagePage(value uint32) {
	objc.Send[struct{}](v.ID, objc.Sel("setUsagePage:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHIDDeviceConfiguration/vendorID
func (v VZMacHIDDeviceConfiguration) VendorID() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("vendorID"))
	return rv
}
func (v VZMacHIDDeviceConfiguration) SetVendorID(value uint16) {
	objc.Send[struct{}](v.ID, objc.Sel("setVendorID:"), value)
}


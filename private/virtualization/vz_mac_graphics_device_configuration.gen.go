// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacGraphicsDeviceConfiguration] class.
var (
	_VZMacGraphicsDeviceConfigurationClass     VZMacGraphicsDeviceConfigurationClass
	_VZMacGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZMacGraphicsDeviceConfigurationClass() VZMacGraphicsDeviceConfigurationClass {
	_VZMacGraphicsDeviceConfigurationClassOnce.Do(func() {
		_VZMacGraphicsDeviceConfigurationClass = VZMacGraphicsDeviceConfigurationClass{class: objc.GetClass("VZMacGraphicsDeviceConfiguration")}
	})
	return _VZMacGraphicsDeviceConfigurationClass
}

// GetVZMacGraphicsDeviceConfigurationClass returns the class object for VZMacGraphicsDeviceConfiguration.
func GetVZMacGraphicsDeviceConfigurationClass() VZMacGraphicsDeviceConfigurationClass {
	return getVZMacGraphicsDeviceConfigurationClass()
}

type VZMacGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacGraphicsDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacGraphicsDeviceConfigurationClass) Alloc() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacGraphicsDeviceConfiguration._deviceFeatureLevel]
//   - [VZMacGraphicsDeviceConfiguration.Set_deviceFeatureLevel]
//   - [VZMacGraphicsDeviceConfiguration._displayPortCount]
//   - [VZMacGraphicsDeviceConfiguration.Set_displayPortCount]
//   - [VZMacGraphicsDeviceConfiguration._enableProcessIsolation]
//   - [VZMacGraphicsDeviceConfiguration.Set_enableProcessIsolation]
//   - [VZMacGraphicsDeviceConfiguration._implicitlyAddsVideoToolboxDevice]
//   - [VZMacGraphicsDeviceConfiguration.Set_implicitlyAddsVideoToolboxDevice]
//   - [VZMacGraphicsDeviceConfiguration._prefersLowPower]
//   - [VZMacGraphicsDeviceConfiguration.Set_prefersLowPower]
//   - [VZMacGraphicsDeviceConfiguration._setDeviceFeatureLevel]
//   - [VZMacGraphicsDeviceConfiguration._setDisplayPortCount]
//   - [VZMacGraphicsDeviceConfiguration._setEnableProcessIsolation]
//   - [VZMacGraphicsDeviceConfiguration._setImplicitlyAddsVideoToolboxDevice]
//   - [VZMacGraphicsDeviceConfiguration._setPrefersLowPower]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration
type VZMacGraphicsDeviceConfiguration struct {
	VZGraphicsDeviceConfiguration
}

// VZMacGraphicsDeviceConfigurationFromID constructs a [VZMacGraphicsDeviceConfiguration] from an objc.ID.
func VZMacGraphicsDeviceConfigurationFromID(id objc.ID) VZMacGraphicsDeviceConfiguration {
	return VZMacGraphicsDeviceConfiguration{VZGraphicsDeviceConfiguration: VZGraphicsDeviceConfigurationFromID(id)}
}

// Ensure VZMacGraphicsDeviceConfiguration implements IVZMacGraphicsDeviceConfiguration.
var _ IVZMacGraphicsDeviceConfiguration = VZMacGraphicsDeviceConfiguration{}

// An interface definition for the [VZMacGraphicsDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZMacGraphicsDeviceConfiguration._deviceFeatureLevel]
//   - [IVZMacGraphicsDeviceConfiguration.Set_deviceFeatureLevel]
//   - [IVZMacGraphicsDeviceConfiguration._displayPortCount]
//   - [IVZMacGraphicsDeviceConfiguration.Set_displayPortCount]
//   - [IVZMacGraphicsDeviceConfiguration._enableProcessIsolation]
//   - [IVZMacGraphicsDeviceConfiguration.Set_enableProcessIsolation]
//   - [IVZMacGraphicsDeviceConfiguration._implicitlyAddsVideoToolboxDevice]
//   - [IVZMacGraphicsDeviceConfiguration.Set_implicitlyAddsVideoToolboxDevice]
//   - [IVZMacGraphicsDeviceConfiguration._prefersLowPower]
//   - [IVZMacGraphicsDeviceConfiguration.Set_prefersLowPower]
//   - [IVZMacGraphicsDeviceConfiguration._setDeviceFeatureLevel]
//   - [IVZMacGraphicsDeviceConfiguration._setDisplayPortCount]
//   - [IVZMacGraphicsDeviceConfiguration._setEnableProcessIsolation]
//   - [IVZMacGraphicsDeviceConfiguration._setImplicitlyAddsVideoToolboxDevice]
//   - [IVZMacGraphicsDeviceConfiguration._setPrefersLowPower]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration
type IVZMacGraphicsDeviceConfiguration interface {
	IVZGraphicsDeviceConfiguration

	// Topic: Methods

	_deviceFeatureLevel() int64
	Set_deviceFeatureLevel(value int64)
	_displayPortCount() uint64
	Set_displayPortCount(value uint64)
	_enableProcessIsolation() bool
	Set_enableProcessIsolation(value bool)
	_implicitlyAddsVideoToolboxDevice() bool
	Set_implicitlyAddsVideoToolboxDevice(value bool)
	_prefersLowPower() bool
	Set_prefersLowPower(value bool)
	_setDeviceFeatureLevel(level int64)
	_setDisplayPortCount(count uint64)
	_setEnableProcessIsolation(isolation bool)
	_setImplicitlyAddsVideoToolboxDevice(device bool)
	_setPrefersLowPower(power bool)
}

// Init initializes the instance.
func (m VZMacGraphicsDeviceConfiguration) Init() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacGraphicsDeviceConfiguration) Autorelease() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDeviceConfiguration creates a new VZMacGraphicsDeviceConfiguration instance.
func NewVZMacGraphicsDeviceConfiguration() VZMacGraphicsDeviceConfiguration {
	class := getVZMacGraphicsDeviceConfigurationClass()
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_setDeviceFeatureLevel:
func (m VZMacGraphicsDeviceConfiguration) _setDeviceFeatureLevel(level int64) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setDeviceFeatureLevel:"), level)
}

// SetDeviceFeatureLevel is an exported wrapper for the private method _setDeviceFeatureLevel.
func (m VZMacGraphicsDeviceConfiguration) SetDeviceFeatureLevel(level int64) {
	m._setDeviceFeatureLevel(level)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_setDisplayPortCount:
func (m VZMacGraphicsDeviceConfiguration) _setDisplayPortCount(count uint64) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setDisplayPortCount:"), count)
}

// SetDisplayPortCount is an exported wrapper for the private method _setDisplayPortCount.
func (m VZMacGraphicsDeviceConfiguration) SetDisplayPortCount(count uint64) {
	m._setDisplayPortCount(count)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_setEnableProcessIsolation:
func (m VZMacGraphicsDeviceConfiguration) _setEnableProcessIsolation(isolation bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setEnableProcessIsolation:"), isolation)
}

// SetEnableProcessIsolation is an exported wrapper for the private method _setEnableProcessIsolation.
func (m VZMacGraphicsDeviceConfiguration) SetEnableProcessIsolation(isolation bool) {
	m._setEnableProcessIsolation(isolation)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_setImplicitlyAddsVideoToolboxDevice:
func (m VZMacGraphicsDeviceConfiguration) _setImplicitlyAddsVideoToolboxDevice(device bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setImplicitlyAddsVideoToolboxDevice:"), device)
}

// SetImplicitlyAddsVideoToolboxDevice is an exported wrapper for the private method _setImplicitlyAddsVideoToolboxDevice.
func (m VZMacGraphicsDeviceConfiguration) SetImplicitlyAddsVideoToolboxDevice(device bool) {
	m._setImplicitlyAddsVideoToolboxDevice(device)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_setPrefersLowPower:
func (m VZMacGraphicsDeviceConfiguration) _setPrefersLowPower(power bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setPrefersLowPower:"), power)
}

// SetPrefersLowPower is an exported wrapper for the private method _setPrefersLowPower.
func (m VZMacGraphicsDeviceConfiguration) SetPrefersLowPower(power bool) {
	m._setPrefersLowPower(power)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_maximumAllowedDisplayCount
func (_VZMacGraphicsDeviceConfigurationClass VZMacGraphicsDeviceConfigurationClass) _maximumAllowedDisplayCount() uint64 {
	rv := objc.Send[uint64](objc.ID(_VZMacGraphicsDeviceConfigurationClass.class), objc.Sel("_maximumAllowedDisplayCount"))
	return rv
}

// MaximumAllowedDisplayCount is an exported wrapper for the private method _maximumAllowedDisplayCount.
func (_VZMacGraphicsDeviceConfigurationClass VZMacGraphicsDeviceConfigurationClass) MaximumAllowedDisplayCount() uint64 {
	return _VZMacGraphicsDeviceConfigurationClass._maximumAllowedDisplayCount()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_deviceFeatureLevel
func (m VZMacGraphicsDeviceConfiguration) _deviceFeatureLevel() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("_deviceFeatureLevel"))
	return rv
}
func (m VZMacGraphicsDeviceConfiguration) Set_deviceFeatureLevel(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("set_deviceFeatureLevel:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_displayPortCount
func (m VZMacGraphicsDeviceConfiguration) _displayPortCount() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("_displayPortCount"))
	return rv
}
func (m VZMacGraphicsDeviceConfiguration) Set_displayPortCount(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("set_displayPortCount:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_enableProcessIsolation
func (m VZMacGraphicsDeviceConfiguration) _enableProcessIsolation() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_enableProcessIsolation"))
	return rv
}
func (m VZMacGraphicsDeviceConfiguration) Set_enableProcessIsolation(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_enableProcessIsolation:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_implicitlyAddsVideoToolboxDevice
func (m VZMacGraphicsDeviceConfiguration) _implicitlyAddsVideoToolboxDevice() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_implicitlyAddsVideoToolboxDevice"))
	return rv
}
func (m VZMacGraphicsDeviceConfiguration) Set_implicitlyAddsVideoToolboxDevice(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_implicitlyAddsVideoToolboxDevice:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/_prefersLowPower
func (m VZMacGraphicsDeviceConfiguration) _prefersLowPower() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_prefersLowPower"))
	return rv
}
func (m VZMacGraphicsDeviceConfiguration) Set_prefersLowPower(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_prefersLowPower:"), value)
}

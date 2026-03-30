// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacNeuralEngineDeviceConfiguration] class.
var (
	_VZMacNeuralEngineDeviceConfigurationClass     VZMacNeuralEngineDeviceConfigurationClass
	_VZMacNeuralEngineDeviceConfigurationClassOnce sync.Once
)

func getVZMacNeuralEngineDeviceConfigurationClass() VZMacNeuralEngineDeviceConfigurationClass {
	_VZMacNeuralEngineDeviceConfigurationClassOnce.Do(func() {
		_VZMacNeuralEngineDeviceConfigurationClass = VZMacNeuralEngineDeviceConfigurationClass{class: objc.GetClass("_VZMacNeuralEngineDeviceConfiguration")}
	})
	return _VZMacNeuralEngineDeviceConfigurationClass
}

// GetVZMacNeuralEngineDeviceConfigurationClass returns the class object for _VZMacNeuralEngineDeviceConfiguration.
func GetVZMacNeuralEngineDeviceConfigurationClass() VZMacNeuralEngineDeviceConfigurationClass {
	return getVZMacNeuralEngineDeviceConfigurationClass()
}

type VZMacNeuralEngineDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacNeuralEngineDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacNeuralEngineDeviceConfigurationClass) Alloc() VZMacNeuralEngineDeviceConfiguration {
	rv := objc.Send[VZMacNeuralEngineDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacNeuralEngineDeviceConfiguration._isSignatureMismatchAllowed]
//   - [VZMacNeuralEngineDeviceConfiguration._setSignatureMismatchAllowed]
//   - [VZMacNeuralEngineDeviceConfiguration._signatureMismatchAllowed]
//   - [VZMacNeuralEngineDeviceConfiguration.Set_signatureMismatchAllowed]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacNeuralEngineDeviceConfiguration
type VZMacNeuralEngineDeviceConfiguration struct {
	VZAcceleratorDeviceConfiguration
}

// VZMacNeuralEngineDeviceConfigurationFromID constructs a [VZMacNeuralEngineDeviceConfiguration] from an objc.ID.
func VZMacNeuralEngineDeviceConfigurationFromID(id objc.ID) VZMacNeuralEngineDeviceConfiguration {
	return VZMacNeuralEngineDeviceConfiguration{VZAcceleratorDeviceConfiguration: VZAcceleratorDeviceConfigurationFromID(id)}
}

// Ensure VZMacNeuralEngineDeviceConfiguration implements IVZMacNeuralEngineDeviceConfiguration.
var _ IVZMacNeuralEngineDeviceConfiguration = VZMacNeuralEngineDeviceConfiguration{}

// An interface definition for the [VZMacNeuralEngineDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZMacNeuralEngineDeviceConfiguration._isSignatureMismatchAllowed]
//   - [IVZMacNeuralEngineDeviceConfiguration._setSignatureMismatchAllowed]
//   - [IVZMacNeuralEngineDeviceConfiguration._signatureMismatchAllowed]
//   - [IVZMacNeuralEngineDeviceConfiguration.Set_signatureMismatchAllowed]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacNeuralEngineDeviceConfiguration
type IVZMacNeuralEngineDeviceConfiguration interface {
	IVZAcceleratorDeviceConfiguration

	// Topic: Methods

	_isSignatureMismatchAllowed() bool
	_setSignatureMismatchAllowed(allowed bool)
	_signatureMismatchAllowed() bool
	Set_signatureMismatchAllowed(value bool)
}

// Init initializes the instance.
func (v VZMacNeuralEngineDeviceConfiguration) Init() VZMacNeuralEngineDeviceConfiguration {
	rv := objc.Send[VZMacNeuralEngineDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacNeuralEngineDeviceConfiguration) Autorelease() VZMacNeuralEngineDeviceConfiguration {
	rv := objc.Send[VZMacNeuralEngineDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacNeuralEngineDeviceConfiguration creates a new VZMacNeuralEngineDeviceConfiguration instance.
func NewVZMacNeuralEngineDeviceConfiguration() VZMacNeuralEngineDeviceConfiguration {
	class := getVZMacNeuralEngineDeviceConfigurationClass()
	rv := objc.Send[VZMacNeuralEngineDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacNeuralEngineDeviceConfiguration/_isSignatureMismatchAllowed
func (v VZMacNeuralEngineDeviceConfiguration) _isSignatureMismatchAllowed() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_isSignatureMismatchAllowed"))
	return rv
}

// IsSignatureMismatchAllowed is an exported wrapper for the private method _isSignatureMismatchAllowed.
func (v VZMacNeuralEngineDeviceConfiguration) IsSignatureMismatchAllowed() bool {
	return v._isSignatureMismatchAllowed()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacNeuralEngineDeviceConfiguration/_setSignatureMismatchAllowed:
func (v VZMacNeuralEngineDeviceConfiguration) _setSignatureMismatchAllowed(allowed bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setSignatureMismatchAllowed:"), allowed)
}

// SetSignatureMismatchAllowed is an exported wrapper for the private method _setSignatureMismatchAllowed.
func (v VZMacNeuralEngineDeviceConfiguration) SetSignatureMismatchAllowed(allowed bool) {
	v._setSignatureMismatchAllowed(allowed)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacNeuralEngineDeviceConfiguration/_signatureMismatchAllowed
func (v VZMacNeuralEngineDeviceConfiguration) _signatureMismatchAllowed() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_signatureMismatchAllowed"))
	return rv
}
func (v VZMacNeuralEngineDeviceConfiguration) Set_signatureMismatchAllowed(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_signatureMismatchAllowed:"), value)
}

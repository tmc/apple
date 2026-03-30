// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZBiometricDeviceConfiguration] class.
var (
	_VZBiometricDeviceConfigurationClass     VZBiometricDeviceConfigurationClass
	_VZBiometricDeviceConfigurationClassOnce sync.Once
)

func getVZBiometricDeviceConfigurationClass() VZBiometricDeviceConfigurationClass {
	_VZBiometricDeviceConfigurationClassOnce.Do(func() {
		_VZBiometricDeviceConfigurationClass = VZBiometricDeviceConfigurationClass{class: objc.GetClass("_VZBiometricDeviceConfiguration")}
	})
	return _VZBiometricDeviceConfigurationClass
}

// GetVZBiometricDeviceConfigurationClass returns the class object for _VZBiometricDeviceConfiguration.
func GetVZBiometricDeviceConfigurationClass() VZBiometricDeviceConfigurationClass {
	return getVZBiometricDeviceConfigurationClass()
}

type VZBiometricDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZBiometricDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBiometricDeviceConfigurationClass) Alloc() VZBiometricDeviceConfiguration {
	rv := objc.Send[VZBiometricDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZBiometricDeviceConfiguration._biometricDeviceWithPlatform]
//   - [VZBiometricDeviceConfiguration._init]
//   - [VZBiometricDeviceConfiguration.EncodeWithEncoder]
//   - [VZBiometricDeviceConfiguration.DebugDescription]
//   - [VZBiometricDeviceConfiguration.Description]
//   - [VZBiometricDeviceConfiguration.Hash]
//   - [VZBiometricDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBiometricDeviceConfiguration
type VZBiometricDeviceConfiguration struct {
	objectivec.Object
}

// VZBiometricDeviceConfigurationFromID constructs a [VZBiometricDeviceConfiguration] from an objc.ID.
func VZBiometricDeviceConfigurationFromID(id objc.ID) VZBiometricDeviceConfiguration {
	return VZBiometricDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZBiometricDeviceConfiguration implements IVZBiometricDeviceConfiguration.
var _ IVZBiometricDeviceConfiguration = VZBiometricDeviceConfiguration{}

// An interface definition for the [VZBiometricDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZBiometricDeviceConfiguration._biometricDeviceWithPlatform]
//   - [IVZBiometricDeviceConfiguration._init]
//   - [IVZBiometricDeviceConfiguration.EncodeWithEncoder]
//   - [IVZBiometricDeviceConfiguration.DebugDescription]
//   - [IVZBiometricDeviceConfiguration.Description]
//   - [IVZBiometricDeviceConfiguration.Hash]
//   - [IVZBiometricDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBiometricDeviceConfiguration
type IVZBiometricDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_biometricDeviceWithPlatform(platform objectivec.IObject) objectivec.IObject
	_init() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZBiometricDeviceConfiguration) Init() VZBiometricDeviceConfiguration {
	rv := objc.Send[VZBiometricDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZBiometricDeviceConfiguration) Autorelease() VZBiometricDeviceConfiguration {
	rv := objc.Send[VZBiometricDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBiometricDeviceConfiguration creates a new VZBiometricDeviceConfiguration instance.
func NewVZBiometricDeviceConfiguration() VZBiometricDeviceConfiguration {
	class := getVZBiometricDeviceConfigurationClass()
	rv := objc.Send[VZBiometricDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBiometricDeviceConfiguration/_biometricDeviceWithPlatform:
func (v VZBiometricDeviceConfiguration) _biometricDeviceWithPlatform(platform objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_biometricDeviceWithPlatform:"), platform)
	return objectivec.Object{ID: rv}
}

// BiometricDeviceWithPlatform is an exported wrapper for the private method _biometricDeviceWithPlatform.
func (v VZBiometricDeviceConfiguration) BiometricDeviceWithPlatform(platform objectivec.IObject) objectivec.IObject {
	return v._biometricDeviceWithPlatform(platform)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBiometricDeviceConfiguration/_init
func (v VZBiometricDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBiometricDeviceConfiguration/encodeWithEncoder:
func (v VZBiometricDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBiometricDeviceConfiguration/debugDescription
func (v VZBiometricDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBiometricDeviceConfiguration/description
func (v VZBiometricDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBiometricDeviceConfiguration/hash
func (v VZBiometricDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBiometricDeviceConfiguration/superclass
func (v VZBiometricDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

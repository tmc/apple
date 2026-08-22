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
	rv := objc.SendIfResponds[VZBiometricDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZBiometricDeviceConfiguration._init]
//   - [VZBiometricDeviceConfiguration.DebugDescription]
//   - [VZBiometricDeviceConfiguration.Description]
//   - [VZBiometricDeviceConfiguration.Hash]
//   - [VZBiometricDeviceConfiguration.Superclass]
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
//   - [IVZBiometricDeviceConfiguration._init]
//   - [IVZBiometricDeviceConfiguration.DebugDescription]
//   - [IVZBiometricDeviceConfiguration.Description]
//   - [IVZBiometricDeviceConfiguration.Hash]
//   - [IVZBiometricDeviceConfiguration.Superclass]
type IVZBiometricDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZBiometricDeviceConfiguration) Init() VZBiometricDeviceConfiguration {
	rv := objc.SendIfResponds[VZBiometricDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZBiometricDeviceConfiguration) Autorelease() VZBiometricDeviceConfiguration {
	rv := objc.SendIfResponds[VZBiometricDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBiometricDeviceConfiguration creates a new VZBiometricDeviceConfiguration instance.
func NewVZBiometricDeviceConfiguration() VZBiometricDeviceConfiguration {
	class := getVZBiometricDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZBiometricDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZBiometricDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZBiometricDeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZBiometricDeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZBiometricDeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZBiometricDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

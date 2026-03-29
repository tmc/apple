// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZAcceleratorDeviceConfiguration] class.
var (
	_VZAcceleratorDeviceConfigurationClass     VZAcceleratorDeviceConfigurationClass
	_VZAcceleratorDeviceConfigurationClassOnce sync.Once
)

func getVZAcceleratorDeviceConfigurationClass() VZAcceleratorDeviceConfigurationClass {
	_VZAcceleratorDeviceConfigurationClassOnce.Do(func() {
		_VZAcceleratorDeviceConfigurationClass = VZAcceleratorDeviceConfigurationClass{class: objc.GetClass("_VZAcceleratorDeviceConfiguration")}
	})
	return _VZAcceleratorDeviceConfigurationClass
}

// GetVZAcceleratorDeviceConfigurationClass returns the class object for _VZAcceleratorDeviceConfiguration.
func GetVZAcceleratorDeviceConfigurationClass() VZAcceleratorDeviceConfigurationClass {
	return getVZAcceleratorDeviceConfigurationClass()
}

type VZAcceleratorDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZAcceleratorDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZAcceleratorDeviceConfigurationClass) Alloc() VZAcceleratorDeviceConfiguration {
	rv := objc.Send[VZAcceleratorDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZAcceleratorDeviceConfiguration._acceleratorDeviceWithPlatform]
//   - [VZAcceleratorDeviceConfiguration._init]
//   - [VZAcceleratorDeviceConfiguration.EncodeWithEncoder]
//   - [VZAcceleratorDeviceConfiguration.DebugDescription]
//   - [VZAcceleratorDeviceConfiguration.Description]
//   - [VZAcceleratorDeviceConfiguration.Hash]
//   - [VZAcceleratorDeviceConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZAcceleratorDeviceConfiguration
type VZAcceleratorDeviceConfiguration struct {
	objectivec.Object
}

// VZAcceleratorDeviceConfigurationFromID constructs a [VZAcceleratorDeviceConfiguration] from an objc.ID.
func VZAcceleratorDeviceConfigurationFromID(id objc.ID) VZAcceleratorDeviceConfiguration {
	return VZAcceleratorDeviceConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZAcceleratorDeviceConfiguration implements IVZAcceleratorDeviceConfiguration.
var _ IVZAcceleratorDeviceConfiguration = VZAcceleratorDeviceConfiguration{}

// An interface definition for the [VZAcceleratorDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZAcceleratorDeviceConfiguration._acceleratorDeviceWithPlatform]
//   - [IVZAcceleratorDeviceConfiguration._init]
//   - [IVZAcceleratorDeviceConfiguration.EncodeWithEncoder]
//   - [IVZAcceleratorDeviceConfiguration.DebugDescription]
//   - [IVZAcceleratorDeviceConfiguration.Description]
//   - [IVZAcceleratorDeviceConfiguration.Hash]
//   - [IVZAcceleratorDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZAcceleratorDeviceConfiguration
type IVZAcceleratorDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_acceleratorDeviceWithPlatform(platform objectivec.IObject) objectivec.IObject
	_init() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZAcceleratorDeviceConfiguration) Init() VZAcceleratorDeviceConfiguration {
	rv := objc.Send[VZAcceleratorDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZAcceleratorDeviceConfiguration) Autorelease() VZAcceleratorDeviceConfiguration {
	rv := objc.Send[VZAcceleratorDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAcceleratorDeviceConfiguration creates a new VZAcceleratorDeviceConfiguration instance.
func NewVZAcceleratorDeviceConfiguration() VZAcceleratorDeviceConfiguration {
	class := getVZAcceleratorDeviceConfigurationClass()
	rv := objc.Send[VZAcceleratorDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZAcceleratorDeviceConfiguration/_acceleratorDeviceWithPlatform:
func (v VZAcceleratorDeviceConfiguration) _acceleratorDeviceWithPlatform(platform objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_acceleratorDeviceWithPlatform:"), platform)
	return objectivec.Object{ID: rv}
}

// AcceleratorDeviceWithPlatform is an exported wrapper for the private method _acceleratorDeviceWithPlatform.
func (v VZAcceleratorDeviceConfiguration) AcceleratorDeviceWithPlatform(platform objectivec.IObject) objectivec.IObject {
	return v._acceleratorDeviceWithPlatform(platform)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZAcceleratorDeviceConfiguration/_init
func (v VZAcceleratorDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZAcceleratorDeviceConfiguration/encodeWithEncoder:
func (v VZAcceleratorDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZAcceleratorDeviceConfiguration/debugDescription
func (v VZAcceleratorDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZAcceleratorDeviceConfiguration/description
func (v VZAcceleratorDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZAcceleratorDeviceConfiguration/hash
func (v VZAcceleratorDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZAcceleratorDeviceConfiguration/superclass
func (v VZAcceleratorDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}


// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPanicDeviceConfiguration] class.
var (
	_VZPanicDeviceConfigurationClass     VZPanicDeviceConfigurationClass
	_VZPanicDeviceConfigurationClassOnce sync.Once
)

func getVZPanicDeviceConfigurationClass() VZPanicDeviceConfigurationClass {
	_VZPanicDeviceConfigurationClassOnce.Do(func() {
		_VZPanicDeviceConfigurationClass = VZPanicDeviceConfigurationClass{class: objc.GetClass("_VZPanicDeviceConfiguration")}
	})
	return _VZPanicDeviceConfigurationClass
}

// GetVZPanicDeviceConfigurationClass returns the class object for _VZPanicDeviceConfiguration.
func GetVZPanicDeviceConfigurationClass() VZPanicDeviceConfigurationClass {
	return getVZPanicDeviceConfigurationClass()
}

type VZPanicDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPanicDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPanicDeviceConfigurationClass) Alloc() VZPanicDeviceConfiguration {
	rv := objc.Send[VZPanicDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZPanicDeviceConfiguration._init]
//   - [VZPanicDeviceConfiguration._panicDevice]
//   - [VZPanicDeviceConfiguration.EncodeWithEncoder]
//   - [VZPanicDeviceConfiguration.DebugDescription]
//   - [VZPanicDeviceConfiguration.Description]
//   - [VZPanicDeviceConfiguration.Hash]
//   - [VZPanicDeviceConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZPanicDeviceConfiguration
type VZPanicDeviceConfiguration struct {
	objectivec.Object
}

// VZPanicDeviceConfigurationFromID constructs a [VZPanicDeviceConfiguration] from an objc.ID.
func VZPanicDeviceConfigurationFromID(id objc.ID) VZPanicDeviceConfiguration {
	return VZPanicDeviceConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZPanicDeviceConfiguration implements IVZPanicDeviceConfiguration.
var _ IVZPanicDeviceConfiguration = VZPanicDeviceConfiguration{}

// An interface definition for the [VZPanicDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZPanicDeviceConfiguration._init]
//   - [IVZPanicDeviceConfiguration._panicDevice]
//   - [IVZPanicDeviceConfiguration.EncodeWithEncoder]
//   - [IVZPanicDeviceConfiguration.DebugDescription]
//   - [IVZPanicDeviceConfiguration.Description]
//   - [IVZPanicDeviceConfiguration.Hash]
//   - [IVZPanicDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPanicDeviceConfiguration
type IVZPanicDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_panicDevice() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZPanicDeviceConfiguration) Init() VZPanicDeviceConfiguration {
	rv := objc.Send[VZPanicDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZPanicDeviceConfiguration) Autorelease() VZPanicDeviceConfiguration {
	rv := objc.Send[VZPanicDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPanicDeviceConfiguration creates a new VZPanicDeviceConfiguration instance.
func NewVZPanicDeviceConfiguration() VZPanicDeviceConfiguration {
	class := getVZPanicDeviceConfigurationClass()
	rv := objc.Send[VZPanicDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPanicDeviceConfiguration/_init
func (v VZPanicDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/_VZPanicDeviceConfiguration/_panicDevice
func (v VZPanicDeviceConfiguration) _panicDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_panicDevice"))
	return objectivec.Object{ID: rv}
}

// PanicDevice is an exported wrapper for the private method _panicDevice.
func (v VZPanicDeviceConfiguration) PanicDevice() objectivec.IObject {
	return v._panicDevice()
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPanicDeviceConfiguration/encodeWithEncoder:
func (v VZPanicDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPanicDeviceConfiguration/debugDescription
func (v VZPanicDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZPanicDeviceConfiguration/description
func (v VZPanicDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZPanicDeviceConfiguration/hash
func (v VZPanicDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZPanicDeviceConfiguration/superclass
func (v VZPanicDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}


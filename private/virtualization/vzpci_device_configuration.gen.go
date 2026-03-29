// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPCIDeviceConfiguration] class.
var (
	_VZPCIDeviceConfigurationClass     VZPCIDeviceConfigurationClass
	_VZPCIDeviceConfigurationClassOnce sync.Once
)

func getVZPCIDeviceConfigurationClass() VZPCIDeviceConfigurationClass {
	_VZPCIDeviceConfigurationClassOnce.Do(func() {
		_VZPCIDeviceConfigurationClass = VZPCIDeviceConfigurationClass{class: objc.GetClass("_VZPCIDeviceConfiguration")}
	})
	return _VZPCIDeviceConfigurationClass
}

// GetVZPCIDeviceConfigurationClass returns the class object for _VZPCIDeviceConfiguration.
func GetVZPCIDeviceConfigurationClass() VZPCIDeviceConfigurationClass {
	return getVZPCIDeviceConfigurationClass()
}

type VZPCIDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPCIDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPCIDeviceConfigurationClass) Alloc() VZPCIDeviceConfiguration {
	rv := objc.Send[VZPCIDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZPCIDeviceConfiguration._init]
//   - [VZPCIDeviceConfiguration._pciDevice]
//   - [VZPCIDeviceConfiguration.EncodeWithEncoder]
//   - [VZPCIDeviceConfiguration.DebugDescription]
//   - [VZPCIDeviceConfiguration.Description]
//   - [VZPCIDeviceConfiguration.Hash]
//   - [VZPCIDeviceConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZPCIDeviceConfiguration
type VZPCIDeviceConfiguration struct {
	objectivec.Object
}

// VZPCIDeviceConfigurationFromID constructs a [VZPCIDeviceConfiguration] from an objc.ID.
func VZPCIDeviceConfigurationFromID(id objc.ID) VZPCIDeviceConfiguration {
	return VZPCIDeviceConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZPCIDeviceConfiguration implements IVZPCIDeviceConfiguration.
var _ IVZPCIDeviceConfiguration = VZPCIDeviceConfiguration{}

// An interface definition for the [VZPCIDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZPCIDeviceConfiguration._init]
//   - [IVZPCIDeviceConfiguration._pciDevice]
//   - [IVZPCIDeviceConfiguration.EncodeWithEncoder]
//   - [IVZPCIDeviceConfiguration.DebugDescription]
//   - [IVZPCIDeviceConfiguration.Description]
//   - [IVZPCIDeviceConfiguration.Hash]
//   - [IVZPCIDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPCIDeviceConfiguration
type IVZPCIDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_pciDevice() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZPCIDeviceConfiguration) Init() VZPCIDeviceConfiguration {
	rv := objc.Send[VZPCIDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZPCIDeviceConfiguration) Autorelease() VZPCIDeviceConfiguration {
	rv := objc.Send[VZPCIDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPCIDeviceConfiguration creates a new VZPCIDeviceConfiguration instance.
func NewVZPCIDeviceConfiguration() VZPCIDeviceConfiguration {
	class := getVZPCIDeviceConfigurationClass()
	rv := objc.Send[VZPCIDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPCIDeviceConfiguration/_init
func (v VZPCIDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/_VZPCIDeviceConfiguration/_pciDevice
func (v VZPCIDeviceConfiguration) _pciDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_pciDevice"))
	return objectivec.Object{ID: rv}
}

// PciDevice is an exported wrapper for the private method _pciDevice.
func (v VZPCIDeviceConfiguration) PciDevice() objectivec.IObject {
	return v._pciDevice()
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPCIDeviceConfiguration/encodeWithEncoder:
func (v VZPCIDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPCIDeviceConfiguration/debugDescription
func (v VZPCIDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZPCIDeviceConfiguration/description
func (v VZPCIDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZPCIDeviceConfiguration/hash
func (v VZPCIDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZPCIDeviceConfiguration/superclass
func (v VZPCIDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}


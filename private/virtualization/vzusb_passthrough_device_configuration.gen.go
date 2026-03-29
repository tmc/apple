// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBPassthroughDeviceConfiguration] class.
var (
	_VZUSBPassthroughDeviceConfigurationClass     VZUSBPassthroughDeviceConfigurationClass
	_VZUSBPassthroughDeviceConfigurationClassOnce sync.Once
)

func getVZUSBPassthroughDeviceConfigurationClass() VZUSBPassthroughDeviceConfigurationClass {
	_VZUSBPassthroughDeviceConfigurationClassOnce.Do(func() {
		_VZUSBPassthroughDeviceConfigurationClass = VZUSBPassthroughDeviceConfigurationClass{class: objc.GetClass("_VZUSBPassthroughDeviceConfiguration")}
	})
	return _VZUSBPassthroughDeviceConfigurationClass
}

// GetVZUSBPassthroughDeviceConfigurationClass returns the class object for _VZUSBPassthroughDeviceConfiguration.
func GetVZUSBPassthroughDeviceConfigurationClass() VZUSBPassthroughDeviceConfigurationClass {
	return getVZUSBPassthroughDeviceConfigurationClass()
}

type VZUSBPassthroughDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBPassthroughDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBPassthroughDeviceConfigurationClass) Alloc() VZUSBPassthroughDeviceConfiguration {
	rv := objc.Send[VZUSBPassthroughDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZUSBPassthroughDeviceConfiguration.Accessory]
//   - [VZUSBPassthroughDeviceConfiguration.EncodeWithEncoder]
//   - [VZUSBPassthroughDeviceConfiguration.IsDuplicateConfiguration]
//   - [VZUSBPassthroughDeviceConfiguration.MakeUSBDeviceWithVirtualMachine]
//   - [VZUSBPassthroughDeviceConfiguration.Signature]
//   - [VZUSBPassthroughDeviceConfiguration.Uuid]
//   - [VZUSBPassthroughDeviceConfiguration.SetUuid]
//   - [VZUSBPassthroughDeviceConfiguration.InitWithDevice]
//   - [VZUSBPassthroughDeviceConfiguration.DebugDescription]
//   - [VZUSBPassthroughDeviceConfiguration.Description]
//   - [VZUSBPassthroughDeviceConfiguration.Hash]
//   - [VZUSBPassthroughDeviceConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration
type VZUSBPassthroughDeviceConfiguration struct {
	objectivec.Object
}

// VZUSBPassthroughDeviceConfigurationFromID constructs a [VZUSBPassthroughDeviceConfiguration] from an objc.ID.
func VZUSBPassthroughDeviceConfigurationFromID(id objc.ID) VZUSBPassthroughDeviceConfiguration {
	return VZUSBPassthroughDeviceConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZUSBPassthroughDeviceConfiguration implements IVZUSBPassthroughDeviceConfiguration.
var _ IVZUSBPassthroughDeviceConfiguration = VZUSBPassthroughDeviceConfiguration{}

// An interface definition for the [VZUSBPassthroughDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZUSBPassthroughDeviceConfiguration.Accessory]
//   - [IVZUSBPassthroughDeviceConfiguration.EncodeWithEncoder]
//   - [IVZUSBPassthroughDeviceConfiguration.IsDuplicateConfiguration]
//   - [IVZUSBPassthroughDeviceConfiguration.MakeUSBDeviceWithVirtualMachine]
//   - [IVZUSBPassthroughDeviceConfiguration.Signature]
//   - [IVZUSBPassthroughDeviceConfiguration.Uuid]
//   - [IVZUSBPassthroughDeviceConfiguration.SetUuid]
//   - [IVZUSBPassthroughDeviceConfiguration.InitWithDevice]
//   - [IVZUSBPassthroughDeviceConfiguration.DebugDescription]
//   - [IVZUSBPassthroughDeviceConfiguration.Description]
//   - [IVZUSBPassthroughDeviceConfiguration.Hash]
//   - [IVZUSBPassthroughDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration
type IVZUSBPassthroughDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	Accessory() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	IsDuplicateConfiguration(configuration objectivec.IObject) bool
	MakeUSBDeviceWithVirtualMachine(machine objectivec.IObject) objectivec.IObject
	Signature() foundation.INSData
	Uuid() foundation.NSUUID
	SetUuid(value foundation.NSUUID)
	InitWithDevice(device objectivec.IObject) VZUSBPassthroughDeviceConfiguration
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZUSBPassthroughDeviceConfiguration) Init() VZUSBPassthroughDeviceConfiguration {
	rv := objc.Send[VZUSBPassthroughDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBPassthroughDeviceConfiguration) Autorelease() VZUSBPassthroughDeviceConfiguration {
	rv := objc.Send[VZUSBPassthroughDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBPassthroughDeviceConfiguration creates a new VZUSBPassthroughDeviceConfiguration instance.
func NewVZUSBPassthroughDeviceConfiguration() VZUSBPassthroughDeviceConfiguration {
	class := getVZUSBPassthroughDeviceConfigurationClass()
	rv := objc.Send[VZUSBPassthroughDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/initWithDevice:
func NewVZUSBPassthroughDeviceConfigurationWithDevice(device objectivec.IObject) VZUSBPassthroughDeviceConfiguration {
	instance := getVZUSBPassthroughDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return VZUSBPassthroughDeviceConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/accessory
func (v VZUSBPassthroughDeviceConfiguration) Accessory() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("accessory"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/encodeWithEncoder:
func (v VZUSBPassthroughDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/isDuplicateConfiguration:
func (v VZUSBPassthroughDeviceConfiguration) IsDuplicateConfiguration(configuration objectivec.IObject) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isDuplicateConfiguration:"), configuration)
	return rv
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/makeUSBDeviceWithVirtualMachine:
func (v VZUSBPassthroughDeviceConfiguration) MakeUSBDeviceWithVirtualMachine(machine objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeUSBDeviceWithVirtualMachine:"), machine)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/initWithDevice:
func (v VZUSBPassthroughDeviceConfiguration) InitWithDevice(device objectivec.IObject) VZUSBPassthroughDeviceConfiguration {
	rv := objc.Send[VZUSBPassthroughDeviceConfiguration](v.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/debugDescription
func (v VZUSBPassthroughDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/description
func (v VZUSBPassthroughDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/hash
func (v VZUSBPassthroughDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/signature
func (v VZUSBPassthroughDeviceConfiguration) Signature() foundation.INSData {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("signature"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/superclass
func (v VZUSBPassthroughDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDeviceConfiguration/uuid
func (v VZUSBPassthroughDeviceConfiguration) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (v VZUSBPassthroughDeviceConfiguration) SetUuid(value foundation.NSUUID) {
	objc.Send[struct{}](v.ID, objc.Sel("setUuid:"), value)
}


// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZIOUSBHostPassthroughDeviceConfiguration] class.
var (
	_VZIOUSBHostPassthroughDeviceConfigurationClass     VZIOUSBHostPassthroughDeviceConfigurationClass
	_VZIOUSBHostPassthroughDeviceConfigurationClassOnce sync.Once
)

func getVZIOUSBHostPassthroughDeviceConfigurationClass() VZIOUSBHostPassthroughDeviceConfigurationClass {
	_VZIOUSBHostPassthroughDeviceConfigurationClassOnce.Do(func() {
		_VZIOUSBHostPassthroughDeviceConfigurationClass = VZIOUSBHostPassthroughDeviceConfigurationClass{class: objc.GetClass("_VZIOUSBHostPassthroughDeviceConfiguration")}
	})
	return _VZIOUSBHostPassthroughDeviceConfigurationClass
}

// GetVZIOUSBHostPassthroughDeviceConfigurationClass returns the class object for _VZIOUSBHostPassthroughDeviceConfiguration.
func GetVZIOUSBHostPassthroughDeviceConfigurationClass() VZIOUSBHostPassthroughDeviceConfigurationClass {
	return getVZIOUSBHostPassthroughDeviceConfigurationClass()
}

type VZIOUSBHostPassthroughDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZIOUSBHostPassthroughDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZIOUSBHostPassthroughDeviceConfigurationClass) Alloc() VZIOUSBHostPassthroughDeviceConfiguration {
	rv := objc.Send[VZIOUSBHostPassthroughDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZIOUSBHostPassthroughDeviceConfiguration.EncodeWithEncoder]
//   - [VZIOUSBHostPassthroughDeviceConfiguration.IsDuplicateConfiguration]
//   - [VZIOUSBHostPassthroughDeviceConfiguration.MakeUSBDeviceWithVirtualMachine]
//   - [VZIOUSBHostPassthroughDeviceConfiguration.Signature]
//   - [VZIOUSBHostPassthroughDeviceConfiguration.Uuid]
//   - [VZIOUSBHostPassthroughDeviceConfiguration.SetUuid]
//   - [VZIOUSBHostPassthroughDeviceConfiguration.InitWithServiceError]
//   - [VZIOUSBHostPassthroughDeviceConfiguration.DebugDescription]
//   - [VZIOUSBHostPassthroughDeviceConfiguration.Description]
//   - [VZIOUSBHostPassthroughDeviceConfiguration.Hash]
//   - [VZIOUSBHostPassthroughDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration
type VZIOUSBHostPassthroughDeviceConfiguration struct {
	objectivec.Object
}

// VZIOUSBHostPassthroughDeviceConfigurationFromID constructs a [VZIOUSBHostPassthroughDeviceConfiguration] from an objc.ID.
func VZIOUSBHostPassthroughDeviceConfigurationFromID(id objc.ID) VZIOUSBHostPassthroughDeviceConfiguration {
	return VZIOUSBHostPassthroughDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZIOUSBHostPassthroughDeviceConfiguration implements IVZIOUSBHostPassthroughDeviceConfiguration.
var _ IVZIOUSBHostPassthroughDeviceConfiguration = VZIOUSBHostPassthroughDeviceConfiguration{}

// An interface definition for the [VZIOUSBHostPassthroughDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.EncodeWithEncoder]
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.IsDuplicateConfiguration]
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.MakeUSBDeviceWithVirtualMachine]
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.Signature]
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.Uuid]
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.SetUuid]
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.InitWithServiceError]
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.DebugDescription]
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.Description]
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.Hash]
//   - [IVZIOUSBHostPassthroughDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration
type IVZIOUSBHostPassthroughDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	IsDuplicateConfiguration(configuration objectivec.IObject) bool
	MakeUSBDeviceWithVirtualMachine(machine objectivec.IObject) objectivec.IObject
	Signature() foundation.INSData
	Uuid() foundation.NSUUID
	SetUuid(value foundation.NSUUID)
	InitWithServiceError(service uint32) (VZIOUSBHostPassthroughDeviceConfiguration, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZIOUSBHostPassthroughDeviceConfiguration) Init() VZIOUSBHostPassthroughDeviceConfiguration {
	rv := objc.Send[VZIOUSBHostPassthroughDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZIOUSBHostPassthroughDeviceConfiguration) Autorelease() VZIOUSBHostPassthroughDeviceConfiguration {
	rv := objc.Send[VZIOUSBHostPassthroughDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZIOUSBHostPassthroughDeviceConfiguration creates a new VZIOUSBHostPassthroughDeviceConfiguration instance.
func NewVZIOUSBHostPassthroughDeviceConfiguration() VZIOUSBHostPassthroughDeviceConfiguration {
	class := getVZIOUSBHostPassthroughDeviceConfigurationClass()
	rv := objc.Send[VZIOUSBHostPassthroughDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/initWithService:error:
func NewVZIOUSBHostPassthroughDeviceConfigurationWithServiceError(service uint32) (VZIOUSBHostPassthroughDeviceConfiguration, error) {
	var errorPtr objc.ID
	instance := getVZIOUSBHostPassthroughDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithService:error:"), service, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZIOUSBHostPassthroughDeviceConfiguration{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZIOUSBHostPassthroughDeviceConfigurationFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/encodeWithEncoder:
func (v VZIOUSBHostPassthroughDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/isDuplicateConfiguration:
func (v VZIOUSBHostPassthroughDeviceConfiguration) IsDuplicateConfiguration(configuration objectivec.IObject) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isDuplicateConfiguration:"), configuration)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/makeUSBDeviceWithVirtualMachine:
func (v VZIOUSBHostPassthroughDeviceConfiguration) MakeUSBDeviceWithVirtualMachine(machine objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeUSBDeviceWithVirtualMachine:"), machine)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/initWithService:error:
func (v VZIOUSBHostPassthroughDeviceConfiguration) InitWithServiceError(service uint32) (VZIOUSBHostPassthroughDeviceConfiguration, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithService:error:"), service, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZIOUSBHostPassthroughDeviceConfiguration{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZIOUSBHostPassthroughDeviceConfigurationFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/fromLocationID:error:
func (_VZIOUSBHostPassthroughDeviceConfigurationClass VZIOUSBHostPassthroughDeviceConfigurationClass) FromLocationIDError(id uint32) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_VZIOUSBHostPassthroughDeviceConfigurationClass.class), objc.Sel("fromLocationID:error:"), id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/debugDescription
func (v VZIOUSBHostPassthroughDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/description
func (v VZIOUSBHostPassthroughDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/hash
func (v VZIOUSBHostPassthroughDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/signature
func (v VZIOUSBHostPassthroughDeviceConfiguration) Signature() foundation.INSData {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("signature"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/superclass
func (v VZIOUSBHostPassthroughDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDeviceConfiguration/uuid
func (v VZIOUSBHostPassthroughDeviceConfiguration) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (v VZIOUSBHostPassthroughDeviceConfiguration) SetUuid(value foundation.NSUUID) {
	objc.Send[struct{}](v.ID, objc.Sel("setUuid:"), value)
}

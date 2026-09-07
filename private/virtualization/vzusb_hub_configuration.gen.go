// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBHubConfiguration] class.
var (
	_VZUSBHubConfigurationClass     VZUSBHubConfigurationClass
	_VZUSBHubConfigurationClassOnce sync.Once
)

func getVZUSBHubConfigurationClass() VZUSBHubConfigurationClass {
	_VZUSBHubConfigurationClassOnce.Do(func() {
		_VZUSBHubConfigurationClass = VZUSBHubConfigurationClass{class: objc.GetClass("_VZUSBHubConfiguration")}
	})
	return _VZUSBHubConfigurationClass
}

// GetVZUSBHubConfigurationClass returns the class object for _VZUSBHubConfiguration.
func GetVZUSBHubConfigurationClass() VZUSBHubConfigurationClass {
	return getVZUSBHubConfigurationClass()
}

type VZUSBHubConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBHubConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBHubConfigurationClass) Alloc() VZUSBHubConfiguration {
	rv := objc.SendIfResponds[VZUSBHubConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZUSBHubConfiguration.IsDuplicateConfiguration]
//   - [VZUSBHubConfiguration.MakeUSBDeviceWithVirtualMachine]
//   - [VZUSBHubConfiguration.Signature]
//   - [VZUSBHubConfiguration.Uuid]
//   - [VZUSBHubConfiguration.SetUuid]
//   - [VZUSBHubConfiguration.ValidateWithError]
//   - [VZUSBHubConfiguration.DebugDescription]
//   - [VZUSBHubConfiguration.Description]
//   - [VZUSBHubConfiguration.Hash]
//   - [VZUSBHubConfiguration.Superclass]
type VZUSBHubConfiguration struct {
	objectivec.Object
}

// VZUSBHubConfigurationFromID constructs a [VZUSBHubConfiguration] from an objc.ID.
func VZUSBHubConfigurationFromID(id objc.ID) VZUSBHubConfiguration {
	return VZUSBHubConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZUSBHubConfiguration implements IVZUSBHubConfiguration.
var _ IVZUSBHubConfiguration = VZUSBHubConfiguration{}

// An interface definition for the [VZUSBHubConfiguration] class.
//
// # Methods
//
//   - [IVZUSBHubConfiguration.IsDuplicateConfiguration]
//   - [IVZUSBHubConfiguration.MakeUSBDeviceWithVirtualMachine]
//   - [IVZUSBHubConfiguration.Signature]
//   - [IVZUSBHubConfiguration.Uuid]
//   - [IVZUSBHubConfiguration.SetUuid]
//   - [IVZUSBHubConfiguration.ValidateWithError]
//   - [IVZUSBHubConfiguration.DebugDescription]
//   - [IVZUSBHubConfiguration.Description]
//   - [IVZUSBHubConfiguration.Hash]
//   - [IVZUSBHubConfiguration.Superclass]
type IVZUSBHubConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	IsDuplicateConfiguration(configuration objectivec.IObject) bool
	MakeUSBDeviceWithVirtualMachine(machine objectivec.IObject) objectivec.IObject
	Signature() foundation.NSData
	Uuid() foundation.NSUUID
	SetUuid(value foundation.NSUUID)
	ValidateWithError() (bool, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZUSBHubConfiguration) Init() VZUSBHubConfiguration {
	rv := objc.SendIfResponds[VZUSBHubConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBHubConfiguration) Autorelease() VZUSBHubConfiguration {
	rv := objc.SendIfResponds[VZUSBHubConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBHubConfiguration creates a new VZUSBHubConfiguration instance.
func NewVZUSBHubConfiguration() VZUSBHubConfiguration {
	class := getVZUSBHubConfigurationClass()
	rv := objc.SendIfResponds[VZUSBHubConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZUSBHubConfiguration) IsDuplicateConfiguration(configuration objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("isDuplicateConfiguration:"), configuration)
	return rv
}
func (v VZUSBHubConfiguration) MakeUSBDeviceWithVirtualMachine(machine objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("makeUSBDeviceWithVirtualMachine:"), machine)
	return objectivec.Object{ID: rv}
}
func (v VZUSBHubConfiguration) ValidateWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("validateWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateWithError: returned NO with nil NSError")
	}
	return rv, nil

}

func (_VZUSBHubConfigurationClass VZUSBHubConfigurationClass) FromLocationIDError(id uint32) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_VZUSBHubConfigurationClass.class), objc.Sel("fromLocationID:error:"), id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_VZUSBHubConfigurationClass VZUSBHubConfigurationClass) PassthroughConfigurationWithServiceError(service uint32) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_VZUSBHubConfigurationClass.class), objc.Sel("passthroughConfigurationWithService:error:"), service, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (v VZUSBHubConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZUSBHubConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZUSBHubConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZUSBHubConfiguration) Signature() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](v.ID, objc.Sel("signature"))
	return foundation.NSData(rv)
}
func (v VZUSBHubConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (v VZUSBHubConfiguration) Uuid() foundation.NSUUID {
	rv := objc.SendIfResponds[foundation.NSUUID](v.ID, objc.Sel("uuid"))
	return foundation.NSUUID(rv)
}
func (v VZUSBHubConfiguration) SetUuid(value foundation.NSUUID) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setUuid:"), value)
}

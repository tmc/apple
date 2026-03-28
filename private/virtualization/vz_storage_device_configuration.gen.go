// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZStorageDeviceConfiguration] class.
var (
	_VZStorageDeviceConfigurationClass     VZStorageDeviceConfigurationClass
	_VZStorageDeviceConfigurationClassOnce sync.Once
)

func getVZStorageDeviceConfigurationClass() VZStorageDeviceConfigurationClass {
	_VZStorageDeviceConfigurationClassOnce.Do(func() {
		_VZStorageDeviceConfigurationClass = VZStorageDeviceConfigurationClass{class: objc.GetClass("VZStorageDeviceConfiguration")}
	})
	return _VZStorageDeviceConfigurationClass
}

// GetVZStorageDeviceConfigurationClass returns the class object for VZStorageDeviceConfiguration.
func GetVZStorageDeviceConfigurationClass() VZStorageDeviceConfigurationClass {
	return getVZStorageDeviceConfigurationClass()
}

type VZStorageDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZStorageDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZStorageDeviceConfigurationClass) Alloc() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZStorageDeviceConfiguration._initWithAttachment]
//   - [VZStorageDeviceConfiguration._setAttachment]
//   - [VZStorageDeviceConfiguration.MakeStorageDeviceForVirtualMachineStorageDeviceIndex]
//   - [VZStorageDeviceConfiguration.DebugDescription]
//   - [VZStorageDeviceConfiguration.Description]
//   - [VZStorageDeviceConfiguration.Hash]
//   - [VZStorageDeviceConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration
type VZStorageDeviceConfiguration struct {
	objectivec.Object
}

// VZStorageDeviceConfigurationFromID constructs a [VZStorageDeviceConfiguration] from an objc.ID.
func VZStorageDeviceConfigurationFromID(id objc.ID) VZStorageDeviceConfiguration {
	return VZStorageDeviceConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZStorageDeviceConfiguration implements IVZStorageDeviceConfiguration.
var _ IVZStorageDeviceConfiguration = VZStorageDeviceConfiguration{}

// An interface definition for the [VZStorageDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZStorageDeviceConfiguration._initWithAttachment]
//   - [IVZStorageDeviceConfiguration._setAttachment]
//   - [IVZStorageDeviceConfiguration.MakeStorageDeviceForVirtualMachineStorageDeviceIndex]
//   - [IVZStorageDeviceConfiguration.DebugDescription]
//   - [IVZStorageDeviceConfiguration.Description]
//   - [IVZStorageDeviceConfiguration.Hash]
//   - [IVZStorageDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration
type IVZStorageDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_initWithAttachment(attachment objectivec.IObject) objectivec.IObject
	_setAttachment(attachment objectivec.IObject)
	MakeStorageDeviceForVirtualMachineStorageDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s VZStorageDeviceConfiguration) Init() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZStorageDeviceConfiguration) Autorelease() VZStorageDeviceConfiguration {
	rv := objc.Send[VZStorageDeviceConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDeviceConfiguration creates a new VZStorageDeviceConfiguration instance.
func NewVZStorageDeviceConfiguration() VZStorageDeviceConfiguration {
	class := getVZStorageDeviceConfigurationClass()
	rv := objc.Send[VZStorageDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration/_initWithAttachment:
func (s VZStorageDeviceConfiguration) _initWithAttachment(attachment objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_initWithAttachment:"), attachment)
	return objectivec.Object{ID: rv}
}

// InitWithAttachment is an exported wrapper for the private method _initWithAttachment.
func (s VZStorageDeviceConfiguration) InitWithAttachment(attachment objectivec.IObject) objectivec.IObject {
	return s._initWithAttachment(attachment)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration/_setAttachment:
func (s VZStorageDeviceConfiguration) _setAttachment(attachment objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_setAttachment:"), attachment)
}

// SetAttachment is an exported wrapper for the private method _setAttachment.
func (s VZStorageDeviceConfiguration) SetAttachment(attachment objectivec.IObject) {
	s._setAttachment(attachment)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration/makeStorageDeviceForVirtualMachine:storageDeviceIndex:
func (s VZStorageDeviceConfiguration) MakeStorageDeviceForVirtualMachineStorageDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeStorageDeviceForVirtualMachine:storageDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration/debugDescription
func (s VZStorageDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration/description
func (s VZStorageDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration/hash
func (s VZStorageDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceConfiguration/superclass
func (s VZStorageDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}


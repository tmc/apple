// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDirectorySharingDeviceConfiguration] class.
var (
	_VZDirectorySharingDeviceConfigurationClass     VZDirectorySharingDeviceConfigurationClass
	_VZDirectorySharingDeviceConfigurationClassOnce sync.Once
)

func getVZDirectorySharingDeviceConfigurationClass() VZDirectorySharingDeviceConfigurationClass {
	_VZDirectorySharingDeviceConfigurationClassOnce.Do(func() {
		_VZDirectorySharingDeviceConfigurationClass = VZDirectorySharingDeviceConfigurationClass{class: objc.GetClass("VZDirectorySharingDeviceConfiguration")}
	})
	return _VZDirectorySharingDeviceConfigurationClass
}

// GetVZDirectorySharingDeviceConfigurationClass returns the class object for VZDirectorySharingDeviceConfiguration.
func GetVZDirectorySharingDeviceConfigurationClass() VZDirectorySharingDeviceConfigurationClass {
	return getVZDirectorySharingDeviceConfigurationClass()
}

type VZDirectorySharingDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDirectorySharingDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDirectorySharingDeviceConfigurationClass) Alloc() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZDirectorySharingDeviceConfiguration._init]
//   - [VZDirectorySharingDeviceConfiguration._makeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex]
//   - [VZDirectorySharingDeviceConfiguration._directorySharingDevice]
//   - [VZDirectorySharingDeviceConfiguration.DebugDescription]
//   - [VZDirectorySharingDeviceConfiguration.Description]
//   - [VZDirectorySharingDeviceConfiguration.Hash]
//   - [VZDirectorySharingDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration
type VZDirectorySharingDeviceConfiguration struct {
	objectivec.Object
}

// VZDirectorySharingDeviceConfigurationFromID constructs a [VZDirectorySharingDeviceConfiguration] from an objc.ID.
func VZDirectorySharingDeviceConfigurationFromID(id objc.ID) VZDirectorySharingDeviceConfiguration {
	return VZDirectorySharingDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZDirectorySharingDeviceConfiguration implements IVZDirectorySharingDeviceConfiguration.
var _ IVZDirectorySharingDeviceConfiguration = VZDirectorySharingDeviceConfiguration{}

// An interface definition for the [VZDirectorySharingDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZDirectorySharingDeviceConfiguration._init]
//   - [IVZDirectorySharingDeviceConfiguration._makeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex]
//   - [IVZDirectorySharingDeviceConfiguration._directorySharingDevice]
//   - [IVZDirectorySharingDeviceConfiguration.DebugDescription]
//   - [IVZDirectorySharingDeviceConfiguration.Description]
//   - [IVZDirectorySharingDeviceConfiguration.Hash]
//   - [IVZDirectorySharingDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration
type IVZDirectorySharingDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_makeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	_directorySharingDevice() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZDirectorySharingDeviceConfiguration) Init() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDirectorySharingDeviceConfiguration) Autorelease() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectorySharingDeviceConfiguration creates a new VZDirectorySharingDeviceConfiguration instance.
func NewVZDirectorySharingDeviceConfiguration() VZDirectorySharingDeviceConfiguration {
	class := getVZDirectorySharingDeviceConfigurationClass()
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration/_init
func (v VZDirectorySharingDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration/_makeDirectorySharingDeviceForVirtualMachine:directorySharingDeviceIndex:
func (v VZDirectorySharingDeviceConfiguration) _makeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_makeDirectorySharingDeviceForVirtualMachine:directorySharingDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// MakeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex is an exported wrapper for the private method _makeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex.
func (v VZDirectorySharingDeviceConfiguration) MakeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_makeDirectorySharingDeviceForVirtualMachine:directorySharingDeviceIndex:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_makeDirectorySharingDeviceForVirtualMachine:directorySharingDeviceIndex:"}
		return nil, err
	}
	return v._makeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex(machine, index), nil
}

// CanMakeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex reports whether the receiver responds to the private selector _makeDirectorySharingDeviceForVirtualMachine:directorySharingDeviceIndex:.
func (v VZDirectorySharingDeviceConfiguration) CanMakeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_makeDirectorySharingDeviceForVirtualMachine:directorySharingDeviceIndex:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration/_directorySharingDevice
func (v VZDirectorySharingDeviceConfiguration) _directorySharingDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_directorySharingDevice"))
	return objectivec.Object{ID: rv}
}

// CanDirectorySharingDevice reports whether the receiver responds to the private selector _directorySharingDevice.
func (v VZDirectorySharingDeviceConfiguration) CanDirectorySharingDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_directorySharingDevice"))
}

// DirectorySharingDevice is an exported wrapper for the private property _directorySharingDevice.
func (v VZDirectorySharingDeviceConfiguration) DirectorySharingDevice() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_directorySharingDevice")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_directorySharingDevice"}
	}
	return v._directorySharingDevice(), nil
}

// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration/debugDescription
func (v VZDirectorySharingDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration/description
func (v VZDirectorySharingDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration/hash
func (v VZDirectorySharingDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration/superclass
func (v VZDirectorySharingDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

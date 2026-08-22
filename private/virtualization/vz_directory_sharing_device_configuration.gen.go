// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[VZDirectorySharingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZDirectorySharingDeviceConfiguration._directorySharingDevice]
//   - [VZDirectorySharingDeviceConfiguration._init]
//   - [VZDirectorySharingDeviceConfiguration._makeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex]
//   - [VZDirectorySharingDeviceConfiguration.DebugDescription]
//   - [VZDirectorySharingDeviceConfiguration.Description]
//   - [VZDirectorySharingDeviceConfiguration.Hash]
//   - [VZDirectorySharingDeviceConfiguration.Superclass]
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
//   - [IVZDirectorySharingDeviceConfiguration._directorySharingDevice]
//   - [IVZDirectorySharingDeviceConfiguration._init]
//   - [IVZDirectorySharingDeviceConfiguration._makeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex]
//   - [IVZDirectorySharingDeviceConfiguration.DebugDescription]
//   - [IVZDirectorySharingDeviceConfiguration.Description]
//   - [IVZDirectorySharingDeviceConfiguration.Hash]
//   - [IVZDirectorySharingDeviceConfiguration.Superclass]
type IVZDirectorySharingDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_directorySharingDevice() unsafe.Pointer
	_init() objectivec.IObject
	_makeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZDirectorySharingDeviceConfiguration) Init() VZDirectorySharingDeviceConfiguration {
	rv := objc.SendIfResponds[VZDirectorySharingDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDirectorySharingDeviceConfiguration) Autorelease() VZDirectorySharingDeviceConfiguration {
	rv := objc.SendIfResponds[VZDirectorySharingDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectorySharingDeviceConfiguration creates a new VZDirectorySharingDeviceConfiguration instance.
func NewVZDirectorySharingDeviceConfiguration() VZDirectorySharingDeviceConfiguration {
	class := getVZDirectorySharingDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZDirectorySharingDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZDirectorySharingDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZDirectorySharingDeviceConfiguration) _makeDirectorySharingDeviceForVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_makeDirectorySharingDeviceForVirtualMachine:directorySharingDeviceIndex:"), machine, index)
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

func (v VZDirectorySharingDeviceConfiguration) _directorySharingDevice() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_directorySharingDevice"))
	return rv
}

// CanDirectorySharingDevice reports whether the receiver responds to the private selector _directorySharingDevice.
func (v VZDirectorySharingDeviceConfiguration) CanDirectorySharingDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_directorySharingDevice"))
}

// DirectorySharingDevice is an exported wrapper for the private property _directorySharingDevice.
func (v VZDirectorySharingDeviceConfiguration) DirectorySharingDevice() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_directorySharingDevice")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_directorySharingDevice"}
	}
	return v._directorySharingDevice(), nil
}
func (v VZDirectorySharingDeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZDirectorySharingDeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZDirectorySharingDeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZDirectorySharingDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

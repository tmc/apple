// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSocketDeviceConfiguration] class.
var (
	_VZSocketDeviceConfigurationClass     VZSocketDeviceConfigurationClass
	_VZSocketDeviceConfigurationClassOnce sync.Once
)

func getVZSocketDeviceConfigurationClass() VZSocketDeviceConfigurationClass {
	_VZSocketDeviceConfigurationClassOnce.Do(func() {
		_VZSocketDeviceConfigurationClass = VZSocketDeviceConfigurationClass{class: objc.GetClass("VZSocketDeviceConfiguration")}
	})
	return _VZSocketDeviceConfigurationClass
}

// GetVZSocketDeviceConfigurationClass returns the class object for VZSocketDeviceConfiguration.
func GetVZSocketDeviceConfigurationClass() VZSocketDeviceConfigurationClass {
	return getVZSocketDeviceConfigurationClass()
}

type VZSocketDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSocketDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSocketDeviceConfigurationClass) Alloc() VZSocketDeviceConfiguration {
	rv := objc.SendIfResponds[VZSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSocketDeviceConfiguration._init]
//   - [VZSocketDeviceConfiguration.MakeSocketDeviceForVirtualMachineIdentifier]
//   - [VZSocketDeviceConfiguration.DebugDescription]
//   - [VZSocketDeviceConfiguration.Description]
//   - [VZSocketDeviceConfiguration.Hash]
//   - [VZSocketDeviceConfiguration.Superclass]
type VZSocketDeviceConfiguration struct {
	objectivec.Object
}

// VZSocketDeviceConfigurationFromID constructs a [VZSocketDeviceConfiguration] from an objc.ID.
func VZSocketDeviceConfigurationFromID(id objc.ID) VZSocketDeviceConfiguration {
	return VZSocketDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZSocketDeviceConfiguration implements IVZSocketDeviceConfiguration.
var _ IVZSocketDeviceConfiguration = VZSocketDeviceConfiguration{}

// An interface definition for the [VZSocketDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZSocketDeviceConfiguration._init]
//   - [IVZSocketDeviceConfiguration.MakeSocketDeviceForVirtualMachineIdentifier]
//   - [IVZSocketDeviceConfiguration.DebugDescription]
//   - [IVZSocketDeviceConfiguration.Description]
//   - [IVZSocketDeviceConfiguration.Hash]
//   - [IVZSocketDeviceConfiguration.Superclass]
type IVZSocketDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	MakeSocketDeviceForVirtualMachineIdentifier(machine objectivec.IObject, identifier uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZSocketDeviceConfiguration) Init() VZSocketDeviceConfiguration {
	rv := objc.SendIfResponds[VZSocketDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSocketDeviceConfiguration) Autorelease() VZSocketDeviceConfiguration {
	rv := objc.SendIfResponds[VZSocketDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSocketDeviceConfiguration creates a new VZSocketDeviceConfiguration instance.
func NewVZSocketDeviceConfiguration() VZSocketDeviceConfiguration {
	class := getVZSocketDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZSocketDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZSocketDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZSocketDeviceConfiguration) MakeSocketDeviceForVirtualMachineIdentifier(machine objectivec.IObject, identifier uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("makeSocketDeviceForVirtualMachine:identifier:"), machine, identifier)
	return objectivec.Object{ID: rv}
}

func (v VZSocketDeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZSocketDeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZSocketDeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZSocketDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

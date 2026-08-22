// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMemoryBalloonDeviceConfiguration] class.
var (
	_VZMemoryBalloonDeviceConfigurationClass     VZMemoryBalloonDeviceConfigurationClass
	_VZMemoryBalloonDeviceConfigurationClassOnce sync.Once
)

func getVZMemoryBalloonDeviceConfigurationClass() VZMemoryBalloonDeviceConfigurationClass {
	_VZMemoryBalloonDeviceConfigurationClassOnce.Do(func() {
		_VZMemoryBalloonDeviceConfigurationClass = VZMemoryBalloonDeviceConfigurationClass{class: objc.GetClass("VZMemoryBalloonDeviceConfiguration")}
	})
	return _VZMemoryBalloonDeviceConfigurationClass
}

// GetVZMemoryBalloonDeviceConfigurationClass returns the class object for VZMemoryBalloonDeviceConfiguration.
func GetVZMemoryBalloonDeviceConfigurationClass() VZMemoryBalloonDeviceConfigurationClass {
	return getVZMemoryBalloonDeviceConfigurationClass()
}

type VZMemoryBalloonDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMemoryBalloonDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMemoryBalloonDeviceConfigurationClass) Alloc() VZMemoryBalloonDeviceConfiguration {
	rv := objc.SendIfResponds[VZMemoryBalloonDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMemoryBalloonDeviceConfiguration._init]
//   - [VZMemoryBalloonDeviceConfiguration.MakeMemoryBalloonDeviceForVirtualMachineMemoryBalloonDeviceIndexMaxTargetMemorySize]
//   - [VZMemoryBalloonDeviceConfiguration.DebugDescription]
//   - [VZMemoryBalloonDeviceConfiguration.Description]
//   - [VZMemoryBalloonDeviceConfiguration.Hash]
//   - [VZMemoryBalloonDeviceConfiguration.Superclass]
type VZMemoryBalloonDeviceConfiguration struct {
	objectivec.Object
}

// VZMemoryBalloonDeviceConfigurationFromID constructs a [VZMemoryBalloonDeviceConfiguration] from an objc.ID.
func VZMemoryBalloonDeviceConfigurationFromID(id objc.ID) VZMemoryBalloonDeviceConfiguration {
	return VZMemoryBalloonDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZMemoryBalloonDeviceConfiguration implements IVZMemoryBalloonDeviceConfiguration.
var _ IVZMemoryBalloonDeviceConfiguration = VZMemoryBalloonDeviceConfiguration{}

// An interface definition for the [VZMemoryBalloonDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZMemoryBalloonDeviceConfiguration._init]
//   - [IVZMemoryBalloonDeviceConfiguration.MakeMemoryBalloonDeviceForVirtualMachineMemoryBalloonDeviceIndexMaxTargetMemorySize]
//   - [IVZMemoryBalloonDeviceConfiguration.DebugDescription]
//   - [IVZMemoryBalloonDeviceConfiguration.Description]
//   - [IVZMemoryBalloonDeviceConfiguration.Hash]
//   - [IVZMemoryBalloonDeviceConfiguration.Superclass]
type IVZMemoryBalloonDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	MakeMemoryBalloonDeviceForVirtualMachineMemoryBalloonDeviceIndexMaxTargetMemorySize(machine objectivec.IObject, index uint64, size uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZMemoryBalloonDeviceConfiguration) Init() VZMemoryBalloonDeviceConfiguration {
	rv := objc.SendIfResponds[VZMemoryBalloonDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMemoryBalloonDeviceConfiguration) Autorelease() VZMemoryBalloonDeviceConfiguration {
	rv := objc.SendIfResponds[VZMemoryBalloonDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMemoryBalloonDeviceConfiguration creates a new VZMemoryBalloonDeviceConfiguration instance.
func NewVZMemoryBalloonDeviceConfiguration() VZMemoryBalloonDeviceConfiguration {
	class := getVZMemoryBalloonDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZMemoryBalloonDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMemoryBalloonDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZMemoryBalloonDeviceConfiguration) MakeMemoryBalloonDeviceForVirtualMachineMemoryBalloonDeviceIndexMaxTargetMemorySize(machine objectivec.IObject, index uint64, size uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("makeMemoryBalloonDeviceForVirtualMachine:memoryBalloonDeviceIndex:maxTargetMemorySize:"), machine, index, size)
	return objectivec.Object{ID: rv}
}

func (v VZMemoryBalloonDeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMemoryBalloonDeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMemoryBalloonDeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZMemoryBalloonDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

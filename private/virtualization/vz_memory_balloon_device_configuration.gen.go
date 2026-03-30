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
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration
type IVZMemoryBalloonDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	MakeMemoryBalloonDeviceForVirtualMachineMemoryBalloonDeviceIndexMaxTargetMemorySize(machine objectivec.IObject, index uint64, size uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (m VZMemoryBalloonDeviceConfiguration) Init() VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMemoryBalloonDeviceConfiguration) Autorelease() VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMemoryBalloonDeviceConfiguration creates a new VZMemoryBalloonDeviceConfiguration instance.
func NewVZMemoryBalloonDeviceConfiguration() VZMemoryBalloonDeviceConfiguration {
	class := getVZMemoryBalloonDeviceConfigurationClass()
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration/_init
func (m VZMemoryBalloonDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration/makeMemoryBalloonDeviceForVirtualMachine:memoryBalloonDeviceIndex:maxTargetMemorySize:
func (m VZMemoryBalloonDeviceConfiguration) MakeMemoryBalloonDeviceForVirtualMachineMemoryBalloonDeviceIndexMaxTargetMemorySize(machine objectivec.IObject, index uint64, size uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("makeMemoryBalloonDeviceForVirtualMachine:memoryBalloonDeviceIndex:maxTargetMemorySize:"), machine, index, size)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration/debugDescription
func (m VZMemoryBalloonDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration/description
func (m VZMemoryBalloonDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration/hash
func (m VZMemoryBalloonDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration/superclass
func (m VZMemoryBalloonDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}

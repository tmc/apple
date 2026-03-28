// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZAudioDeviceConfiguration] class.
var (
	_VZAudioDeviceConfigurationClass     VZAudioDeviceConfigurationClass
	_VZAudioDeviceConfigurationClassOnce sync.Once
)

func getVZAudioDeviceConfigurationClass() VZAudioDeviceConfigurationClass {
	_VZAudioDeviceConfigurationClassOnce.Do(func() {
		_VZAudioDeviceConfigurationClass = VZAudioDeviceConfigurationClass{class: objc.GetClass("VZAudioDeviceConfiguration")}
	})
	return _VZAudioDeviceConfigurationClass
}

// GetVZAudioDeviceConfigurationClass returns the class object for VZAudioDeviceConfiguration.
func GetVZAudioDeviceConfigurationClass() VZAudioDeviceConfigurationClass {
	return getVZAudioDeviceConfigurationClass()
}

type VZAudioDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZAudioDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZAudioDeviceConfigurationClass) Alloc() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZAudioDeviceConfiguration._init]
//   - [VZAudioDeviceConfiguration._makeAudioDeviceForVirtualMachineAudioDeviceIndex]
//   - [VZAudioDeviceConfiguration._role]
//   - [VZAudioDeviceConfiguration.Set_role]
//   - [VZAudioDeviceConfiguration._setRole]
//   - [VZAudioDeviceConfiguration._audioDevice]
//   - [VZAudioDeviceConfiguration.DebugDescription]
//   - [VZAudioDeviceConfiguration.Description]
//   - [VZAudioDeviceConfiguration.Hash]
//   - [VZAudioDeviceConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration
type VZAudioDeviceConfiguration struct {
	objectivec.Object
}

// VZAudioDeviceConfigurationFromID constructs a [VZAudioDeviceConfiguration] from an objc.ID.
func VZAudioDeviceConfigurationFromID(id objc.ID) VZAudioDeviceConfiguration {
	return VZAudioDeviceConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZAudioDeviceConfiguration implements IVZAudioDeviceConfiguration.
var _ IVZAudioDeviceConfiguration = VZAudioDeviceConfiguration{}

// An interface definition for the [VZAudioDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZAudioDeviceConfiguration._init]
//   - [IVZAudioDeviceConfiguration._makeAudioDeviceForVirtualMachineAudioDeviceIndex]
//   - [IVZAudioDeviceConfiguration._role]
//   - [IVZAudioDeviceConfiguration.Set_role]
//   - [IVZAudioDeviceConfiguration._setRole]
//   - [IVZAudioDeviceConfiguration._audioDevice]
//   - [IVZAudioDeviceConfiguration.DebugDescription]
//   - [IVZAudioDeviceConfiguration.Description]
//   - [IVZAudioDeviceConfiguration.Hash]
//   - [IVZAudioDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration
type IVZAudioDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_makeAudioDeviceForVirtualMachineAudioDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	_role() int64
	Set_role(value int64)
	_setRole(role int64)
	_audioDevice() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (a VZAudioDeviceConfiguration) Init() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a VZAudioDeviceConfiguration) Autorelease() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioDeviceConfiguration creates a new VZAudioDeviceConfiguration instance.
func NewVZAudioDeviceConfiguration() VZAudioDeviceConfiguration {
	class := getVZAudioDeviceConfigurationClass()
	rv := objc.Send[VZAudioDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/_init
func (a VZAudioDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/_makeAudioDeviceForVirtualMachine:audioDeviceIndex:
func (a VZAudioDeviceConfiguration) _makeAudioDeviceForVirtualMachineAudioDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("_makeAudioDeviceForVirtualMachine:audioDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// MakeAudioDeviceForVirtualMachineAudioDeviceIndex is an exported wrapper for the private method _makeAudioDeviceForVirtualMachineAudioDeviceIndex.
func (a VZAudioDeviceConfiguration) MakeAudioDeviceForVirtualMachineAudioDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	return a._makeAudioDeviceForVirtualMachineAudioDeviceIndex(machine, index)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/_setRole:
func (a VZAudioDeviceConfiguration) _setRole(role int64) {
	objc.Send[objc.ID](a.ID, objc.Sel("_setRole:"), role)
}

// SetRole is an exported wrapper for the private method _setRole.
func (a VZAudioDeviceConfiguration) SetRole(role int64) {
	a._setRole(role)
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/_audioDevice
func (a VZAudioDeviceConfiguration) _audioDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("_audioDevice"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/_role
func (a VZAudioDeviceConfiguration) _role() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("_role"))
	return rv
}
func (a VZAudioDeviceConfiguration) Set_role(value int64) {
	objc.Send[struct{}](a.ID, objc.Sel("set_role:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/debugDescription
func (a VZAudioDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/description
func (a VZAudioDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/hash
func (a VZAudioDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/superclass
func (a VZAudioDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](a.ID, objc.Sel("superclass"))
	return rv
}


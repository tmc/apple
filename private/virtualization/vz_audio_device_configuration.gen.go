// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
//
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
func (v VZAudioDeviceConfiguration) Init() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZAudioDeviceConfiguration) Autorelease() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioDeviceConfiguration creates a new VZAudioDeviceConfiguration instance.
func NewVZAudioDeviceConfiguration() VZAudioDeviceConfiguration {
	class := getVZAudioDeviceConfigurationClass()
	rv := objc.Send[VZAudioDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/_init
func (v VZAudioDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/_makeAudioDeviceForVirtualMachine:audioDeviceIndex:
func (v VZAudioDeviceConfiguration) _makeAudioDeviceForVirtualMachineAudioDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_makeAudioDeviceForVirtualMachine:audioDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// MakeAudioDeviceForVirtualMachineAudioDeviceIndex is an exported wrapper for the private method _makeAudioDeviceForVirtualMachineAudioDeviceIndex.
func (v VZAudioDeviceConfiguration) MakeAudioDeviceForVirtualMachineAudioDeviceIndex(machine objectivec.IObject, index uint64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_makeAudioDeviceForVirtualMachine:audioDeviceIndex:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_makeAudioDeviceForVirtualMachine:audioDeviceIndex:"}
		return nil, err
	}
	return v._makeAudioDeviceForVirtualMachineAudioDeviceIndex(machine, index), nil
}

// CanMakeAudioDeviceForVirtualMachineAudioDeviceIndex reports whether the receiver responds to the private selector _makeAudioDeviceForVirtualMachine:audioDeviceIndex:.
func (v VZAudioDeviceConfiguration) CanMakeAudioDeviceForVirtualMachineAudioDeviceIndex() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_makeAudioDeviceForVirtualMachine:audioDeviceIndex:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/_setRole:
func (v VZAudioDeviceConfiguration) _setRole(role int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setRole:"), role)
}

// SetRole is an exported wrapper for the private method _setRole.
func (v VZAudioDeviceConfiguration) SetRole(role int64) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setRole:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setRole:"}
		return err
	}
	v._setRole(role)
	return nil
}

// CanSetRole reports whether the receiver responds to the private selector _setRole:.
func (v VZAudioDeviceConfiguration) CanSetRole() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setRole:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/_audioDevice
func (v VZAudioDeviceConfiguration) _audioDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_audioDevice"))
	return objectivec.Object{ID: rv}
}

// CanAudioDevice reports whether the receiver responds to the private selector _audioDevice.
func (v VZAudioDeviceConfiguration) CanAudioDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_audioDevice"))
}

// AudioDevice is an exported wrapper for the private property _audioDevice.
func (v VZAudioDeviceConfiguration) AudioDevice() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_audioDevice")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_audioDevice"}
	}
	return v._audioDevice(), nil
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/_role
func (v VZAudioDeviceConfiguration) _role() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_role"))
	return rv
}

// CanRole reports whether the receiver responds to the private selector _role.
func (v VZAudioDeviceConfiguration) CanRole() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_role"))
}

// Role is an exported wrapper for the private property _role.
func (v VZAudioDeviceConfiguration) Role() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_role")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_role"}
	}
	return v._role(), nil
}
func (v VZAudioDeviceConfiguration) Set_role(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("set_role:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/debugDescription
func (v VZAudioDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/description
func (v VZAudioDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/hash
func (v VZAudioDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration/superclass
func (v VZAudioDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

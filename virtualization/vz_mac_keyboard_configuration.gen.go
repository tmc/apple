// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacKeyboardConfiguration] class.
var (
	_VZMacKeyboardConfigurationClass     VZMacKeyboardConfigurationClass
	_VZMacKeyboardConfigurationClassOnce sync.Once
)

func getVZMacKeyboardConfigurationClass() VZMacKeyboardConfigurationClass {
	_VZMacKeyboardConfigurationClassOnce.Do(func() {
		_VZMacKeyboardConfigurationClass = VZMacKeyboardConfigurationClass{class: objc.GetClass("VZMacKeyboardConfiguration")}
	})
	return _VZMacKeyboardConfigurationClass
}

// GetVZMacKeyboardConfigurationClass returns the class object for VZMacKeyboardConfiguration.
func GetVZMacKeyboardConfigurationClass() VZMacKeyboardConfigurationClass {
	return getVZMacKeyboardConfigurationClass()
}

type VZMacKeyboardConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacKeyboardConfigurationClass) Alloc() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A device that defines the configuration for a Mac keyboard.
//
// # Overview
// 
// Use this configuration to attach a Mac keyboard configuration to a VM. A
// [VZVirtualMachineView] can use this device to send key events to the VM,
// including the Mac-specific key events, such as the Globe key.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacKeyboardConfiguration
type VZMacKeyboardConfiguration struct {
	VZKeyboardConfiguration
}

// VZMacKeyboardConfigurationFromID constructs a [VZMacKeyboardConfiguration] from an objc.ID.
//
// A device that defines the configuration for a Mac keyboard.
func VZMacKeyboardConfigurationFromID(id objc.ID) VZMacKeyboardConfiguration {
	return VZMacKeyboardConfiguration{VZKeyboardConfiguration: VZKeyboardConfigurationFromID(id)}
}
// NOTE: VZMacKeyboardConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMacKeyboardConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacKeyboardConfiguration
type IVZMacKeyboardConfiguration interface {
	IVZKeyboardConfiguration
}

// Init initializes the instance.
func (m VZMacKeyboardConfiguration) Init() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacKeyboardConfiguration) Autorelease() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacKeyboardConfiguration creates a new VZMacKeyboardConfiguration instance.
func NewVZMacKeyboardConfiguration() VZMacKeyboardConfiguration {
	class := getVZMacKeyboardConfigurationClass()
	rv := objc.Send[VZMacKeyboardConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}


// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZGenericPlatformConfiguration] class.
var (
	_VZGenericPlatformConfigurationClass     VZGenericPlatformConfigurationClass
	_VZGenericPlatformConfigurationClassOnce sync.Once
)

func getVZGenericPlatformConfigurationClass() VZGenericPlatformConfigurationClass {
	_VZGenericPlatformConfigurationClassOnce.Do(func() {
		_VZGenericPlatformConfigurationClass = VZGenericPlatformConfigurationClass{class: objc.GetClass("VZGenericPlatformConfiguration")}
	})
	return _VZGenericPlatformConfigurationClass
}

// GetVZGenericPlatformConfigurationClass returns the class object for VZGenericPlatformConfiguration.
func GetVZGenericPlatformConfigurationClass() VZGenericPlatformConfigurationClass {
	return getVZGenericPlatformConfigurationClass()
}

type VZGenericPlatformConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGenericPlatformConfigurationClass) Alloc() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The platform configuration for a generic Intel or ARM virtual machine.
//
// # Identifying the platform configuration
//
//   - [VZGenericPlatformConfiguration.MachineIdentifier]: A value that represents a unique identifier for the virtual machine.
//   - [VZGenericPlatformConfiguration.SetMachineIdentifier]
//   - [VZGenericPlatformConfiguration.NestedVirtualizationEnabled]: A Boolean value that indicates whether nested virtualization is in an enabled state.
//   - [VZGenericPlatformConfiguration.SetNestedVirtualizationEnabled]
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration
type VZGenericPlatformConfiguration struct {
	VZPlatformConfiguration
}

// VZGenericPlatformConfigurationFromID constructs a [VZGenericPlatformConfiguration] from an objc.ID.
//
// The platform configuration for a generic Intel or ARM virtual machine.
func VZGenericPlatformConfigurationFromID(id objc.ID) VZGenericPlatformConfiguration {
	return VZGenericPlatformConfiguration{VZPlatformConfiguration: VZPlatformConfigurationFromID(id)}
}
// NOTE: VZGenericPlatformConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZGenericPlatformConfiguration] class.
//
// # Identifying the platform configuration
//
//   - [IVZGenericPlatformConfiguration.MachineIdentifier]: A value that represents a unique identifier for the virtual machine.
//   - [IVZGenericPlatformConfiguration.SetMachineIdentifier]
//   - [IVZGenericPlatformConfiguration.NestedVirtualizationEnabled]: A Boolean value that indicates whether nested virtualization is in an enabled state.
//   - [IVZGenericPlatformConfiguration.SetNestedVirtualizationEnabled]
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration
type IVZGenericPlatformConfiguration interface {
	IVZPlatformConfiguration

	// Topic: Identifying the platform configuration

	// A value that represents a unique identifier for the virtual machine.
	MachineIdentifier() IVZGenericMachineIdentifier
	SetMachineIdentifier(value IVZGenericMachineIdentifier)
	// A Boolean value that indicates whether nested virtualization is in an enabled state.
	NestedVirtualizationEnabled() bool
	SetNestedVirtualizationEnabled(value bool)
}

// Init initializes the instance.
func (g VZGenericPlatformConfiguration) Init() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VZGenericPlatformConfiguration) Autorelease() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGenericPlatformConfiguration creates a new VZGenericPlatformConfiguration instance.
func NewVZGenericPlatformConfiguration() VZGenericPlatformConfiguration {
	class := getVZGenericPlatformConfigurationClass()
	rv := objc.Send[VZGenericPlatformConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A value that represents a unique identifier for the virtual machine.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/machineIdentifier
func (g VZGenericPlatformConfiguration) MachineIdentifier() IVZGenericMachineIdentifier {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("machineIdentifier"))
	return VZGenericMachineIdentifierFromID(objc.ID(rv))
}
func (g VZGenericPlatformConfiguration) SetMachineIdentifier(value IVZGenericMachineIdentifier) {
	objc.Send[struct{}](g.ID, objc.Sel("setMachineIdentifier:"), value)
}
// A Boolean value that indicates whether nested virtualization is in an
// enabled state.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/isNestedVirtualizationEnabled
func (g VZGenericPlatformConfiguration) NestedVirtualizationEnabled() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isNestedVirtualizationEnabled"))
	return rv
}
func (g VZGenericPlatformConfiguration) SetNestedVirtualizationEnabled(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setNestedVirtualizationEnabled:"), value)
}

// A Boolean value that describes whether the platform configuration supports
// nested virtualization.
//
// # Discussion
// 
// Use this property to check whether support is available for the platform.
// As the following example shows, if the framework supports nested
// virtualization on the host, use [NestedVirtualizationEnabled] to enable the
// feature:
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/isNestedVirtualizationSupported
func (_VZGenericPlatformConfigurationClass VZGenericPlatformConfigurationClass) NestedVirtualizationSupported() bool {
	rv := objc.Send[bool](objc.ID(_VZGenericPlatformConfigurationClass.class), objc.Sel("isNestedVirtualizationSupported"))
	return rv
}


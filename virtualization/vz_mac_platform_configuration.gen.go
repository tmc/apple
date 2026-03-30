// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacPlatformConfiguration] class.
var (
	_VZMacPlatformConfigurationClass     VZMacPlatformConfigurationClass
	_VZMacPlatformConfigurationClassOnce sync.Once
)

func getVZMacPlatformConfigurationClass() VZMacPlatformConfigurationClass {
	_VZMacPlatformConfigurationClassOnce.Do(func() {
		_VZMacPlatformConfigurationClass = VZMacPlatformConfigurationClass{class: objc.GetClass("VZMacPlatformConfiguration")}
	})
	return _VZMacPlatformConfigurationClass
}

// GetVZMacPlatformConfigurationClass returns the class object for VZMacPlatformConfiguration.
func GetVZMacPlatformConfigurationClass() VZMacPlatformConfigurationClass {
	return getVZMacPlatformConfigurationClass()
}

type VZMacPlatformConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacPlatformConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacPlatformConfigurationClass) Alloc() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The platform configuration for booting macOS on Apple silicon.
//
// # Overview
//
// When creating a VM, the [VZMacPlatformConfiguration.HardwareModel] and [VZMacPlatformConfiguration.AuxiliaryStorage] depend on
// the restore image that you use to install macOS.
//
// To choose the hardware model, start from
// [VZMacOSRestoreImage].[VZMacPlatformConfiguration.MostFeaturefulSupportedConfiguration] to get a
// supported configuration, then use its
// [VZMacOSConfigurationRequirements].[VZMacPlatformConfiguration.HardwareModel] property to get the
// hardware model.
//
// Use the hardware model to set up [VZMacPlatformConfiguration] and to
// initialize a new auxiliary storage with
// [InitCreatingStorageAtURLHardwareModelOptionsError].
//
// When you save a VM to disk and load it again, you must restore the
// [VZMacPlatformConfiguration.HardwareModel], [VZMacPlatformConfiguration.MachineIdentifier] and [VZMacPlatformConfiguration.AuxiliaryStorage] properties to
// their original values.
//
// If you create multiple VMs from the same configuration, each should have a
// unique `auxiliaryStorage` and `machineIdentifier`.
//
// # Getting platform properties
//
//   - [VZMacPlatformConfiguration.AuxiliaryStorage]: The Mac auxiliary storage.
//   - [VZMacPlatformConfiguration.SetAuxiliaryStorage]
//   - [VZMacPlatformConfiguration.HardwareModel]: The Mac hardware model.
//   - [VZMacPlatformConfiguration.SetHardwareModel]
//   - [VZMacPlatformConfiguration.MachineIdentifier]: The Mac machine identifier.
//   - [VZMacPlatformConfiguration.SetMachineIdentifier]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration
type VZMacPlatformConfiguration struct {
	VZPlatformConfiguration
}

// VZMacPlatformConfigurationFromID constructs a [VZMacPlatformConfiguration] from an objc.ID.
//
// The platform configuration for booting macOS on Apple silicon.
func VZMacPlatformConfigurationFromID(id objc.ID) VZMacPlatformConfiguration {
	return VZMacPlatformConfiguration{VZPlatformConfiguration: VZPlatformConfigurationFromID(id)}
}

// NOTE: VZMacPlatformConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMacPlatformConfiguration] class.
//
// # Getting platform properties
//
//   - [IVZMacPlatformConfiguration.AuxiliaryStorage]: The Mac auxiliary storage.
//   - [IVZMacPlatformConfiguration.SetAuxiliaryStorage]
//   - [IVZMacPlatformConfiguration.HardwareModel]: The Mac hardware model.
//   - [IVZMacPlatformConfiguration.SetHardwareModel]
//   - [IVZMacPlatformConfiguration.MachineIdentifier]: The Mac machine identifier.
//   - [IVZMacPlatformConfiguration.SetMachineIdentifier]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration
type IVZMacPlatformConfiguration interface {
	IVZPlatformConfiguration

	// Topic: Getting platform properties

	// The Mac auxiliary storage.
	AuxiliaryStorage() IVZMacAuxiliaryStorage
	SetAuxiliaryStorage(value IVZMacAuxiliaryStorage)
	// The Mac hardware model.
	HardwareModel() IVZMacHardwareModel
	SetHardwareModel(value IVZMacHardwareModel)
	// The Mac machine identifier.
	MachineIdentifier() IVZMacMachineIdentifier
	SetMachineIdentifier(value IVZMacMachineIdentifier)

	// This object represents the most fully featured configuration that’s supported by both the current host and by this restore image.
	MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements
	SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements)
}

// Init initializes the instance.
func (m VZMacPlatformConfiguration) Init() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacPlatformConfiguration) Autorelease() VZMacPlatformConfiguration {
	rv := objc.Send[VZMacPlatformConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacPlatformConfiguration creates a new VZMacPlatformConfiguration instance.
func NewVZMacPlatformConfiguration() VZMacPlatformConfiguration {
	class := getVZMacPlatformConfigurationClass()
	rv := objc.Send[VZMacPlatformConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The Mac auxiliary storage.
//
// # Discussion
//
// When creating a VM, the hardware model of the `auxiliaryStorage` must match
// the hardware model of the `hardwareModel` property. Defaults to `nil`, but
// you must set a value for a configuration to be valid.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/auxiliaryStorage
func (m VZMacPlatformConfiguration) AuxiliaryStorage() IVZMacAuxiliaryStorage {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("auxiliaryStorage"))
	return VZMacAuxiliaryStorageFromID(objc.ID(rv))
}
func (m VZMacPlatformConfiguration) SetAuxiliaryStorage(value IVZMacAuxiliaryStorage) {
	objc.Send[struct{}](m.ID, objc.Sel("setAuxiliaryStorage:"), value)
}

// The Mac hardware model.
//
// # Discussion
//
// When creating a VM, the [HardwareModel] depends on the restore image that
// you use to install macOS.
//
// To choose the hardware model, start from
// [VZMacOSRestoreImage].[MostFeaturefulSupportedConfiguration] to get a
// supported configuration, then use its
// [VZMacOSConfigurationRequirements].[HardwareModel] property to get the
// hardware model.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/hardwareModel
func (m VZMacPlatformConfiguration) HardwareModel() IVZMacHardwareModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("hardwareModel"))
	return VZMacHardwareModelFromID(objc.ID(rv))
}
func (m VZMacPlatformConfiguration) SetHardwareModel(value IVZMacHardwareModel) {
	objc.Send[struct{}](m.ID, objc.Sel("setHardwareModel:"), value)
}

// The Mac machine identifier.
//
// # Discussion
//
// This value uniquely identifies an instance of a VM. Running two VMs
// concurrently with the same identifier results in undefined behavior in the
// guest operating system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacPlatformConfiguration/machineIdentifier
func (m VZMacPlatformConfiguration) MachineIdentifier() IVZMacMachineIdentifier {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("machineIdentifier"))
	return VZMacMachineIdentifierFromID(objc.ID(rv))
}
func (m VZMacPlatformConfiguration) SetMachineIdentifier(value IVZMacMachineIdentifier) {
	objc.Send[struct{}](m.ID, objc.Sel("setMachineIdentifier:"), value)
}

// This object represents the most fully featured configuration that’s
// supported by both the current host and by this restore image.
//
// See: https://developer.apple.com/documentation/virtualization/vzmacosrestoreimage/mostfeaturefulsupportedconfiguration
func (m VZMacPlatformConfiguration) MostFeaturefulSupportedConfiguration() IVZMacOSConfigurationRequirements {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mostFeaturefulSupportedConfiguration"))
	return VZMacOSConfigurationRequirementsFromID(objc.ID(rv))
}
func (m VZMacPlatformConfiguration) SetMostFeaturefulSupportedConfiguration(value IVZMacOSConfigurationRequirements) {
	objc.Send[struct{}](m.ID, objc.Sel("setMostFeaturefulSupportedConfiguration:"), value)
}

// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacOSConfigurationRequirements] class.
var (
	_VZMacOSConfigurationRequirementsClass     VZMacOSConfigurationRequirementsClass
	_VZMacOSConfigurationRequirementsClassOnce sync.Once
)

func getVZMacOSConfigurationRequirementsClass() VZMacOSConfigurationRequirementsClass {
	_VZMacOSConfigurationRequirementsClassOnce.Do(func() {
		_VZMacOSConfigurationRequirementsClass = VZMacOSConfigurationRequirementsClass{class: objc.GetClass("VZMacOSConfigurationRequirements")}
	})
	return _VZMacOSConfigurationRequirementsClass
}

// GetVZMacOSConfigurationRequirementsClass returns the class object for VZMacOSConfigurationRequirements.
func GetVZMacOSConfigurationRequirementsClass() VZMacOSConfigurationRequirementsClass {
	return getVZMacOSConfigurationRequirementsClass()
}

type VZMacOSConfigurationRequirementsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacOSConfigurationRequirementsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSConfigurationRequirementsClass) Alloc() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes the parameter constraints required by a specific
// configuration of macOS.
//
// # Configuration Requirements
//
//   - [VZMacOSConfigurationRequirements.HardwareModel]: The hardware model for this configuration.
//   - [VZMacOSConfigurationRequirements.MinimumSupportedCPUCount]: The minimum supported number of CPUs for this configuration.
//   - [VZMacOSConfigurationRequirements.MinimumSupportedMemorySize]: The minimum supported memory size for this configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements
type VZMacOSConfigurationRequirements struct {
	objectivec.Object
}

// VZMacOSConfigurationRequirementsFromID constructs a [VZMacOSConfigurationRequirements] from an objc.ID.
//
// An object that describes the parameter constraints required by a specific
// configuration of macOS.
func VZMacOSConfigurationRequirementsFromID(id objc.ID) VZMacOSConfigurationRequirements {
	return VZMacOSConfigurationRequirements{objectivec.Object{ID: id}}
}

// NOTE: VZMacOSConfigurationRequirements adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMacOSConfigurationRequirements] class.
//
// # Configuration Requirements
//
//   - [IVZMacOSConfigurationRequirements.HardwareModel]: The hardware model for this configuration.
//   - [IVZMacOSConfigurationRequirements.MinimumSupportedCPUCount]: The minimum supported number of CPUs for this configuration.
//   - [IVZMacOSConfigurationRequirements.MinimumSupportedMemorySize]: The minimum supported memory size for this configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements
type IVZMacOSConfigurationRequirements interface {
	objectivec.IObject

	// Topic: Configuration Requirements

	// The hardware model for this configuration.
	HardwareModel() IVZMacHardwareModel
	// The minimum supported number of CPUs for this configuration.
	MinimumSupportedCPUCount() uint
	// The minimum supported memory size for this configuration.
	MinimumSupportedMemorySize() uint64
}

// Init initializes the instance.
func (m VZMacOSConfigurationRequirements) Init() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacOSConfigurationRequirements) Autorelease() VZMacOSConfigurationRequirements {
	rv := objc.Send[VZMacOSConfigurationRequirements](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSConfigurationRequirements creates a new VZMacOSConfigurationRequirements instance.
func NewVZMacOSConfigurationRequirements() VZMacOSConfigurationRequirements {
	class := getVZMacOSConfigurationRequirementsClass()
	rv := objc.Send[VZMacOSConfigurationRequirements](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The hardware model for this configuration.
//
// # Discussion
//
// Use a hardware model to configure a new VM that meets a set of specific
// requirements.
//
// After creating the hardware model, use [VZMacPlatformConfiguration]
// [HardwareModel] to configure the Mac platform, and
// [InitCreatingStorageAtURLHardwareModelOptionsError] to create its auxiliary
// storage.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements/hardwareModel
func (m VZMacOSConfigurationRequirements) HardwareModel() IVZMacHardwareModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("hardwareModel"))
	return VZMacHardwareModelFromID(objc.ID(rv))
}

// The minimum supported number of CPUs for this configuration.
//
// # Discussion
//
// This property specifies the minimum number of CPUs required by the
// associated macOS configuration.
//
// You associate a [VZMacOSConfigurationRequirements] with a specific
// [VZMacOSRestoreImage] object, which results in a specific macOS
// configuration.
//
// Installing or running the associated configuration of macOS on a virtual
// machine with fewer than the specified number of CPUs results in undefined
// behavior.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements/minimumSupportedCPUCount
func (m VZMacOSConfigurationRequirements) MinimumSupportedCPUCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("minimumSupportedCPUCount"))
	return rv
}

// The minimum supported memory size for this configuration.
//
// # Discussion
//
// This property specifies the minimum amount of memory required by the
// associated macOS configuration.
//
// You associate a [VZMacOSConfigurationRequirements] with a specific
// [VZMacOSRestoreImage] object, which results in a specific macOS
// configuration.
//
// Installing or running the associated configuration of macOS on a VM with
// less than this amount of memory results in undefined behavior.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSConfigurationRequirements/minimumSupportedMemorySize
func (m VZMacOSConfigurationRequirements) MinimumSupportedMemorySize() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("minimumSupportedMemorySize"))
	return rv
}

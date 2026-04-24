// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomVirtioDeviceConfiguration] class.
var (
	_VZCustomVirtioDeviceConfigurationClass     VZCustomVirtioDeviceConfigurationClass
	_VZCustomVirtioDeviceConfigurationClassOnce sync.Once
)

func getVZCustomVirtioDeviceConfigurationClass() VZCustomVirtioDeviceConfigurationClass {
	_VZCustomVirtioDeviceConfigurationClassOnce.Do(func() {
		_VZCustomVirtioDeviceConfigurationClass = VZCustomVirtioDeviceConfigurationClass{class: objc.GetClass("_VZCustomVirtioDeviceConfiguration")}
	})
	return _VZCustomVirtioDeviceConfigurationClass
}

// GetVZCustomVirtioDeviceConfigurationClass returns the class object for _VZCustomVirtioDeviceConfiguration.
func GetVZCustomVirtioDeviceConfigurationClass() VZCustomVirtioDeviceConfigurationClass {
	return getVZCustomVirtioDeviceConfigurationClass()
}

type VZCustomVirtioDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomVirtioDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomVirtioDeviceConfigurationClass) Alloc() VZCustomVirtioDeviceConfiguration {
	rv := objc.Send[VZCustomVirtioDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomVirtioDeviceConfiguration.PCIClassID]
//   - [VZCustomVirtioDeviceConfiguration.SetPCIClassID]
//   - [VZCustomVirtioDeviceConfiguration.PCISubclassID]
//   - [VZCustomVirtioDeviceConfiguration.SetPCISubclassID]
//   - [VZCustomVirtioDeviceConfiguration._PCIDeviceID]
//   - [VZCustomVirtioDeviceConfiguration._PCISubsystemID]
//   - [VZCustomVirtioDeviceConfiguration._PCISubsystemVendorID]
//   - [VZCustomVirtioDeviceConfiguration._PCIVendorID]
//   - [VZCustomVirtioDeviceConfiguration._pluginName]
//   - [VZCustomVirtioDeviceConfiguration.Set_pluginName]
//   - [VZCustomVirtioDeviceConfiguration._pluginPersonality]
//   - [VZCustomVirtioDeviceConfiguration.Set_pluginPersonality]
//   - [VZCustomVirtioDeviceConfiguration._setPCIDeviceID]
//   - [VZCustomVirtioDeviceConfiguration._setPCISubsystemID]
//   - [VZCustomVirtioDeviceConfiguration._setPCISubsystemVendorID]
//   - [VZCustomVirtioDeviceConfiguration._setPCIVendorID]
//   - [VZCustomVirtioDeviceConfiguration._setPluginName]
//   - [VZCustomVirtioDeviceConfiguration._setPluginPersonality]
//   - [VZCustomVirtioDeviceConfiguration._setSupportsSaveRestore]
//   - [VZCustomVirtioDeviceConfiguration._supportsSaveRestore]
//   - [VZCustomVirtioDeviceConfiguration.DeviceID]
//   - [VZCustomVirtioDeviceConfiguration.SetDeviceID]
//   - [VZCustomVirtioDeviceConfiguration.DeviceSpecificConfiguration]
//   - [VZCustomVirtioDeviceConfiguration.SetDeviceSpecificConfiguration]
//   - [VZCustomVirtioDeviceConfiguration.MandatoryFeaturesAtIndex]
//   - [VZCustomVirtioDeviceConfiguration.OptionalFeaturesAtIndex]
//   - [VZCustomVirtioDeviceConfiguration.Provider]
//   - [VZCustomVirtioDeviceConfiguration.SetProvider]
//   - [VZCustomVirtioDeviceConfiguration.SetMandatoryFeaturesAtIndex]
//   - [VZCustomVirtioDeviceConfiguration.SetOptionalFeaturesAtIndex]
//   - [VZCustomVirtioDeviceConfiguration.VirtioQueueCount]
//   - [VZCustomVirtioDeviceConfiguration.SetVirtioQueueCount]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration
type VZCustomVirtioDeviceConfiguration struct {
	objectivec.Object
}

// VZCustomVirtioDeviceConfigurationFromID constructs a [VZCustomVirtioDeviceConfiguration] from an objc.ID.
func VZCustomVirtioDeviceConfigurationFromID(id objc.ID) VZCustomVirtioDeviceConfiguration {
	return VZCustomVirtioDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZCustomVirtioDeviceConfiguration implements IVZCustomVirtioDeviceConfiguration.
var _ IVZCustomVirtioDeviceConfiguration = VZCustomVirtioDeviceConfiguration{}

// An interface definition for the [VZCustomVirtioDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZCustomVirtioDeviceConfiguration.PCIClassID]
//   - [IVZCustomVirtioDeviceConfiguration.SetPCIClassID]
//   - [IVZCustomVirtioDeviceConfiguration.PCISubclassID]
//   - [IVZCustomVirtioDeviceConfiguration.SetPCISubclassID]
//   - [IVZCustomVirtioDeviceConfiguration._PCIDeviceID]
//   - [IVZCustomVirtioDeviceConfiguration._PCISubsystemID]
//   - [IVZCustomVirtioDeviceConfiguration._PCISubsystemVendorID]
//   - [IVZCustomVirtioDeviceConfiguration._PCIVendorID]
//   - [IVZCustomVirtioDeviceConfiguration._pluginName]
//   - [IVZCustomVirtioDeviceConfiguration.Set_pluginName]
//   - [IVZCustomVirtioDeviceConfiguration._pluginPersonality]
//   - [IVZCustomVirtioDeviceConfiguration.Set_pluginPersonality]
//   - [IVZCustomVirtioDeviceConfiguration._setPCIDeviceID]
//   - [IVZCustomVirtioDeviceConfiguration._setPCISubsystemID]
//   - [IVZCustomVirtioDeviceConfiguration._setPCISubsystemVendorID]
//   - [IVZCustomVirtioDeviceConfiguration._setPCIVendorID]
//   - [IVZCustomVirtioDeviceConfiguration._setPluginName]
//   - [IVZCustomVirtioDeviceConfiguration._setPluginPersonality]
//   - [IVZCustomVirtioDeviceConfiguration._setSupportsSaveRestore]
//   - [IVZCustomVirtioDeviceConfiguration._supportsSaveRestore]
//   - [IVZCustomVirtioDeviceConfiguration.DeviceID]
//   - [IVZCustomVirtioDeviceConfiguration.SetDeviceID]
//   - [IVZCustomVirtioDeviceConfiguration.DeviceSpecificConfiguration]
//   - [IVZCustomVirtioDeviceConfiguration.SetDeviceSpecificConfiguration]
//   - [IVZCustomVirtioDeviceConfiguration.MandatoryFeaturesAtIndex]
//   - [IVZCustomVirtioDeviceConfiguration.OptionalFeaturesAtIndex]
//   - [IVZCustomVirtioDeviceConfiguration.Provider]
//   - [IVZCustomVirtioDeviceConfiguration.SetProvider]
//   - [IVZCustomVirtioDeviceConfiguration.SetMandatoryFeaturesAtIndex]
//   - [IVZCustomVirtioDeviceConfiguration.SetOptionalFeaturesAtIndex]
//   - [IVZCustomVirtioDeviceConfiguration.VirtioQueueCount]
//   - [IVZCustomVirtioDeviceConfiguration.SetVirtioQueueCount]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration
type IVZCustomVirtioDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	PCIClassID() byte
	SetPCIClassID(value byte)
	PCISubclassID() byte
	SetPCISubclassID(value byte)
	_PCIDeviceID() uint16
	_PCISubsystemID() uint16
	_PCISubsystemVendorID() uint16
	_PCIVendorID() uint16
	_pluginName() string
	Set_pluginName(value string)
	_pluginPersonality() string
	Set_pluginPersonality(value string)
	_setPCIDeviceID(id uint16)
	_setPCISubsystemID(id uint16)
	_setPCISubsystemVendorID(id uint16)
	_setPCIVendorID(id uint16)
	_setPluginName(name objectivec.IObject)
	_setPluginPersonality(personality objectivec.IObject)
	_setSupportsSaveRestore(restore bool)
	_supportsSaveRestore() bool
	DeviceID() uint16
	SetDeviceID(value uint16)
	DeviceSpecificConfiguration() IVZVirtioDeviceSpecificConfiguration
	SetDeviceSpecificConfiguration(value IVZVirtioDeviceSpecificConfiguration)
	MandatoryFeaturesAtIndex(index uint64) uint32
	OptionalFeaturesAtIndex(index uint64) uint32
	Provider() IVZCustomVirtioDeviceProvider
	SetProvider(value IVZCustomVirtioDeviceProvider)
	SetMandatoryFeaturesAtIndex(features uint32, index uint64)
	SetOptionalFeaturesAtIndex(features uint32, index uint64)
	VirtioQueueCount() uint16
	SetVirtioQueueCount(value uint16)
}

// Init initializes the instance.
func (v VZCustomVirtioDeviceConfiguration) Init() VZCustomVirtioDeviceConfiguration {
	rv := objc.Send[VZCustomVirtioDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomVirtioDeviceConfiguration) Autorelease() VZCustomVirtioDeviceConfiguration {
	rv := objc.Send[VZCustomVirtioDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomVirtioDeviceConfiguration creates a new VZCustomVirtioDeviceConfiguration instance.
func NewVZCustomVirtioDeviceConfiguration() VZCustomVirtioDeviceConfiguration {
	class := getVZCustomVirtioDeviceConfigurationClass()
	rv := objc.Send[VZCustomVirtioDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_PCIDeviceID
func (v VZCustomVirtioDeviceConfiguration) _PCIDeviceID() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("_PCIDeviceID"))
	return rv
}

// PCIDeviceID is an exported wrapper for the private method _PCIDeviceID.
func (v VZCustomVirtioDeviceConfiguration) PCIDeviceID() uint16 {
	return v._PCIDeviceID()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_PCISubsystemID
func (v VZCustomVirtioDeviceConfiguration) _PCISubsystemID() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("_PCISubsystemID"))
	return rv
}

// PCISubsystemID is an exported wrapper for the private method _PCISubsystemID.
func (v VZCustomVirtioDeviceConfiguration) PCISubsystemID() uint16 {
	return v._PCISubsystemID()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_PCISubsystemVendorID
func (v VZCustomVirtioDeviceConfiguration) _PCISubsystemVendorID() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("_PCISubsystemVendorID"))
	return rv
}

// PCISubsystemVendorID is an exported wrapper for the private method _PCISubsystemVendorID.
func (v VZCustomVirtioDeviceConfiguration) PCISubsystemVendorID() uint16 {
	return v._PCISubsystemVendorID()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_PCIVendorID
func (v VZCustomVirtioDeviceConfiguration) _PCIVendorID() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("_PCIVendorID"))
	return rv
}

// PCIVendorID is an exported wrapper for the private method _PCIVendorID.
func (v VZCustomVirtioDeviceConfiguration) PCIVendorID() uint16 {
	return v._PCIVendorID()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_setPCIDeviceID:
func (v VZCustomVirtioDeviceConfiguration) _setPCIDeviceID(id uint16) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPCIDeviceID:"), id)
}

// SetPCIDeviceID is an exported wrapper for the private method _setPCIDeviceID.
func (v VZCustomVirtioDeviceConfiguration) SetPCIDeviceID(id uint16) {
	v._setPCIDeviceID(id)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_setPCISubsystemID:
func (v VZCustomVirtioDeviceConfiguration) _setPCISubsystemID(id uint16) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPCISubsystemID:"), id)
}

// SetPCISubsystemID is an exported wrapper for the private method _setPCISubsystemID.
func (v VZCustomVirtioDeviceConfiguration) SetPCISubsystemID(id uint16) {
	v._setPCISubsystemID(id)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_setPCISubsystemVendorID:
func (v VZCustomVirtioDeviceConfiguration) _setPCISubsystemVendorID(id uint16) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPCISubsystemVendorID:"), id)
}

// SetPCISubsystemVendorID is an exported wrapper for the private method _setPCISubsystemVendorID.
func (v VZCustomVirtioDeviceConfiguration) SetPCISubsystemVendorID(id uint16) {
	v._setPCISubsystemVendorID(id)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_setPCIVendorID:
func (v VZCustomVirtioDeviceConfiguration) _setPCIVendorID(id uint16) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPCIVendorID:"), id)
}

// SetPCIVendorID is an exported wrapper for the private method _setPCIVendorID.
func (v VZCustomVirtioDeviceConfiguration) SetPCIVendorID(id uint16) {
	v._setPCIVendorID(id)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_setPluginName:
func (v VZCustomVirtioDeviceConfiguration) _setPluginName(name objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPluginName:"), name)
}

// SetPluginName is an exported wrapper for the private method _setPluginName.
func (v VZCustomVirtioDeviceConfiguration) SetPluginName(name objectivec.IObject) {
	v._setPluginName(name)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_setPluginPersonality:
func (v VZCustomVirtioDeviceConfiguration) _setPluginPersonality(personality objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPluginPersonality:"), personality)
}

// SetPluginPersonality is an exported wrapper for the private method _setPluginPersonality.
func (v VZCustomVirtioDeviceConfiguration) SetPluginPersonality(personality objectivec.IObject) {
	v._setPluginPersonality(personality)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_setSupportsSaveRestore:
func (v VZCustomVirtioDeviceConfiguration) _setSupportsSaveRestore(restore bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setSupportsSaveRestore:"), restore)
}

// SetSupportsSaveRestore is an exported wrapper for the private method _setSupportsSaveRestore.
func (v VZCustomVirtioDeviceConfiguration) SetSupportsSaveRestore(restore bool) {
	v._setSupportsSaveRestore(restore)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_supportsSaveRestore
func (v VZCustomVirtioDeviceConfiguration) _supportsSaveRestore() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_supportsSaveRestore"))
	return rv
}

// SupportsSaveRestore is an exported wrapper for the private method _supportsSaveRestore.
func (v VZCustomVirtioDeviceConfiguration) SupportsSaveRestore() bool {
	return v._supportsSaveRestore()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/mandatoryFeaturesAtIndex:
func (v VZCustomVirtioDeviceConfiguration) MandatoryFeaturesAtIndex(index uint64) uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("mandatoryFeaturesAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/optionalFeaturesAtIndex:
func (v VZCustomVirtioDeviceConfiguration) OptionalFeaturesAtIndex(index uint64) uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("optionalFeaturesAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/setMandatoryFeatures:atIndex:
func (v VZCustomVirtioDeviceConfiguration) SetMandatoryFeaturesAtIndex(features uint32, index uint64) {
	objc.Send[objc.ID](v.ID, objc.Sel("setMandatoryFeatures:atIndex:"), features, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/setOptionalFeatures:atIndex:
func (v VZCustomVirtioDeviceConfiguration) SetOptionalFeaturesAtIndex(features uint32, index uint64) {
	objc.Send[objc.ID](v.ID, objc.Sel("setOptionalFeatures:atIndex:"), features, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/PCIClassID
func (v VZCustomVirtioDeviceConfiguration) PCIClassID() byte {
	rv := objc.Send[byte](v.ID, objc.Sel("PCIClassID"))
	return rv
}
func (v VZCustomVirtioDeviceConfiguration) SetPCIClassID(value byte) {
	objc.Send[struct{}](v.ID, objc.Sel("setPCIClassID:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/PCISubclassID
func (v VZCustomVirtioDeviceConfiguration) PCISubclassID() byte {
	rv := objc.Send[byte](v.ID, objc.Sel("PCISubclassID"))
	return rv
}
func (v VZCustomVirtioDeviceConfiguration) SetPCISubclassID(value byte) {
	objc.Send[struct{}](v.ID, objc.Sel("setPCISubclassID:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_pluginName
func (v VZCustomVirtioDeviceConfiguration) _pluginName() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_pluginName"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCustomVirtioDeviceConfiguration) Set_pluginName(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("set_pluginName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/_pluginPersonality
func (v VZCustomVirtioDeviceConfiguration) _pluginPersonality() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_pluginPersonality"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCustomVirtioDeviceConfiguration) Set_pluginPersonality(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("set_pluginPersonality:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/deviceID
func (v VZCustomVirtioDeviceConfiguration) DeviceID() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("deviceID"))
	return rv
}
func (v VZCustomVirtioDeviceConfiguration) SetDeviceID(value uint16) {
	objc.Send[struct{}](v.ID, objc.Sel("setDeviceID:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/deviceSpecificConfiguration
func (v VZCustomVirtioDeviceConfiguration) DeviceSpecificConfiguration() IVZVirtioDeviceSpecificConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("deviceSpecificConfiguration"))
	return VZVirtioDeviceSpecificConfigurationFromID(objc.ID(rv))
}
func (v VZCustomVirtioDeviceConfiguration) SetDeviceSpecificConfiguration(value IVZVirtioDeviceSpecificConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setDeviceSpecificConfiguration:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/provider
func (v VZCustomVirtioDeviceConfiguration) Provider() IVZCustomVirtioDeviceProvider {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("provider"))
	return VZCustomVirtioDeviceProviderFromID(objc.ID(rv))
}
func (v VZCustomVirtioDeviceConfiguration) SetProvider(value IVZCustomVirtioDeviceProvider) {
	objc.Send[struct{}](v.ID, objc.Sel("setProvider:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceConfiguration/virtioQueueCount
func (v VZCustomVirtioDeviceConfiguration) VirtioQueueCount() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("virtioQueueCount"))
	return rv
}
func (v VZCustomVirtioDeviceConfiguration) SetVirtioQueueCount(value uint16) {
	objc.Send[struct{}](v.ID, objc.Sel("setVirtioQueueCount:"), value)
}

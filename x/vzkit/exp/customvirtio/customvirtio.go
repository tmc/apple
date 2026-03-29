package customvirtio

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
)

// Config describes a custom Virtio device configuration.
type Config struct {
	PCIClassID                 byte
	PCISubclassID              byte
	PCIDeviceID                uint16
	PCISubsystemID             uint16
	PCISubsystemVendorID       uint16
	PCIVendorID                uint16
	PluginName                 string
	PluginPersonality          string
	DeviceID                   uint16
	DeviceSpecificConfiguration objectivec.IObject
	MandatoryFeatures          []uint32
	OptionalFeatures           []uint32
	Provider                   objectivec.IObject
	VirtioQueueCount           uint16
	SupportsSaveRestore        bool
}

// NewConfig creates a custom Virtio device configuration and applies cfg.
func NewConfig(cfg Config) (pvz.VZCustomVirtioDeviceConfiguration, error) {
	device := pvz.NewVZCustomVirtioDeviceConfiguration()
	if device.ID == 0 {
		return device, fmt.Errorf("create custom virtio device configuration")
	}
	device.SetPCIClassID(cfg.PCIClassID)
	device.SetPCISubclassID(cfg.PCISubclassID)
	device.SetPCIDeviceID(cfg.PCIDeviceID)
	device.SetPCISubsystemID(cfg.PCISubsystemID)
	device.SetPCISubsystemVendorID(cfg.PCISubsystemVendorID)
	device.SetPCIVendorID(cfg.PCIVendorID)
	device.Set_pluginName(cfg.PluginName)
	device.Set_pluginPersonality(cfg.PluginPersonality)
	device.SetDeviceID(cfg.DeviceID)
	if cfg.DeviceSpecificConfiguration != nil && cfg.DeviceSpecificConfiguration.GetID() != 0 {
		spec := pvz.VZVirtioDeviceSpecificConfigurationFromID(cfg.DeviceSpecificConfiguration.GetID())
		device.SetDeviceSpecificConfiguration(&spec)
	}
	for i, features := range cfg.MandatoryFeatures {
		device.SetMandatoryFeaturesAtIndex(features, uint64(i))
	}
	for i, features := range cfg.OptionalFeatures {
		device.SetOptionalFeaturesAtIndex(features, uint64(i))
	}
	if cfg.Provider != nil && cfg.Provider.GetID() != 0 {
		provider := pvz.VZCustomVirtioDeviceProviderFromID(cfg.Provider.GetID())
		device.SetProvider(&provider)
	}
	device.SetVirtioQueueCount(cfg.VirtioQueueCount)
	device.SetSupportsSaveRestore(cfg.SupportsSaveRestore)
	return device, nil
}

// NewProvider creates the base Virtio provider object.
func NewProvider() (pvz.VZCustomVirtioDeviceProvider, error) {
	provider := pvz.NewVZCustomVirtioDeviceProvider()
	if provider.ID == 0 {
		return provider, fmt.Errorf("create custom virtio provider")
	}
	return provider, nil
}

// NewPluginProvider creates a plugin-backed Virtio provider.
func NewPluginProvider(name, personality string) (pvz.VZCustomVirtioDevicePluginProvider, error) {
	provider := pvz.NewVZCustomVirtioDevicePluginProviderWithPluginNamePluginPersonality(
		foundation.NewStringWithString(name),
		foundation.NewStringWithString(personality),
	)
	if provider.ID == 0 {
		return provider, fmt.Errorf("create custom virtio plugin provider")
	}
	return provider, nil
}

// NewDelegateProvider creates a delegate-backed Virtio provider.
func NewDelegateProvider(queue, delegate objectivec.IObject) (pvz.VZCustomVirtioDeviceDelegateProvider, error) {
	provider := pvz.NewVZCustomVirtioDeviceDelegateProviderWithDeviceQueueDelegate(queue, delegate)
	if provider.ID == 0 {
		return provider, fmt.Errorf("create custom virtio delegate provider")
	}
	return provider, nil
}

package custommmio

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
)

// Config describes a custom MMIO device configuration.
type Config struct {
	MMIORegions         foundation.INSArray
	AdditionalProperties foundation.INSDictionary
	AdditionalXPCProperties objectivec.IObject
	IRQs                foundation.INSArray
	Provider            objectivec.IObject
	SupportsSaveRestore bool
}

// NewConfig creates a custom MMIO device configuration and applies cfg.
func NewConfig(cfg Config) (pvz.VZCustomMMIODeviceConfiguration, error) {
	device := pvz.NewVZCustomMMIODeviceConfiguration()
	if device.ID == 0 {
		return device, fmt.Errorf("create custom mmio device configuration")
	}
	device.SetMMIORegions(cfg.MMIORegions)
	device.SetAdditionalProperties(cfg.AdditionalProperties)
	if cfg.AdditionalXPCProperties != nil && cfg.AdditionalXPCProperties.GetID() != 0 {
		device.SetAdditionalXPCProperties(objectivec.ObjectFromID(cfg.AdditionalXPCProperties.GetID()))
	}
	device.SetIrqs(cfg.IRQs)
	if cfg.Provider != nil && cfg.Provider.GetID() != 0 {
		provider := pvz.VZCustomMMIODeviceProviderFromID(cfg.Provider.GetID())
		device.SetProvider(&provider)
	}
	device.SetSupportsSaveRestore(cfg.SupportsSaveRestore)
	return device, nil
}

// NewProvider creates the base MMIO provider object.
func NewProvider() (pvz.VZCustomMMIODeviceProvider, error) {
	provider := pvz.NewVZCustomMMIODeviceProvider()
	if provider.ID == 0 {
		return provider, fmt.Errorf("create custom mmio provider")
	}
	return provider, nil
}

// NewPluginProvider creates a plugin-backed MMIO provider.
func NewPluginProvider(name, personality string) (pvz.VZCustomMMIODevicePluginProvider, error) {
	provider := pvz.NewVZCustomMMIODevicePluginProviderWithPluginNamePluginPersonality(
		foundation.NewStringWithString(name),
		foundation.NewStringWithString(personality),
	)
	if provider.ID == 0 {
		return provider, fmt.Errorf("create custom mmio plugin provider")
	}
	return provider, nil
}

// NewDelegateProvider creates a delegate-backed MMIO provider.
func NewDelegateProvider(queue, delegate objectivec.IObject) (pvz.VZCustomMMIODeviceDelegateProvider, error) {
	provider := pvz.NewVZCustomMMIODeviceDelegateProviderWithDeviceQueueDelegate(queue, delegate)
	if provider.ID == 0 {
		return provider, fmt.Errorf("create custom mmio delegate provider")
	}
	return provider, nil
}

package accelerator

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
)

// NewConfiguration creates the base accelerator configuration.
func NewConfiguration() (pvz.VZAcceleratorDeviceConfiguration, error) {
	cfg := pvz.NewVZAcceleratorDeviceConfiguration()
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create accelerator device configuration")
	}
	return cfg, nil
}

// NewDevice creates the runtime accelerator device for a platform object.
func NewDevice(platform objectivec.IObject) (objectivec.IObject, error) {
	_ = platform
	return nil, fmt.Errorf("create accelerator device: unsupported C++ ABI")
}

// NewBifrostConfiguration creates a Bifrost configuration with the given attachment and MMIO size.
func NewBifrostConfiguration(attachment objectivec.IObject, mmioSize uint64) (pvz.VZBifrostDeviceConfiguration, error) {
	cfg := pvz.NewVZBifrostDeviceConfiguration()
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create bifrost device configuration")
	}
	created, err := cfg.InitWithAttachmentMMIOSize(attachment, mmioSize)
	if err != nil {
		return cfg, fmt.Errorf("create bifrost device configuration: %w", err)
	}
	cfg = pvz.VZBifrostDeviceConfigurationFromID(created.GetID())
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create bifrost device configuration")
	}
	return cfg, nil
}

// NewMacBifrostConfiguration creates a macOS Bifrost configuration with the given attachment and MMIO size.
func NewMacBifrostConfiguration(attachment objectivec.IObject, mmioSize uint64) (pvz.VZMacBifrostDeviceConfiguration, error) {
	cfg := pvz.NewVZMacBifrostDeviceConfigurationWithAttachmentMMIOSize(attachment, mmioSize)
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create mac bifrost device configuration")
	}
	return cfg, nil
}

// NewScalerConfiguration creates the mac scaler accelerator configuration.
func NewScalerConfiguration() (pvz.VZMacScalerAcceleratorDeviceConfiguration, error) {
	cfg := pvz.NewVZMacScalerAcceleratorDeviceConfiguration()
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create mac scaler accelerator device configuration")
	}
	return cfg, nil
}

// NewMacVideoToolboxConfiguration creates the mac video toolbox accelerator configuration.
func NewMacVideoToolboxConfiguration() (pvz.VZMacVideoToolboxDeviceConfiguration, error) {
	supported, err := pvz.GetVZMacVideoToolboxDeviceConfigurationClass().IsSupported()
	if err != nil {
		return pvz.VZMacVideoToolboxDeviceConfiguration{}, fmt.Errorf("mac video toolbox accelerator support: %w", err)
	}
	if !supported {
		return pvz.VZMacVideoToolboxDeviceConfiguration{}, fmt.Errorf("mac video toolbox accelerator unsupported")
	}
	cfg := pvz.NewVZMacVideoToolboxDeviceConfiguration()
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create mac video toolbox device configuration")
	}
	return cfg, nil
}

// NewCoprocessorConfiguration creates the base coprocessor configuration.
func NewCoprocessorConfiguration() (pvz.VZCoprocessorConfiguration, error) {
	cfg := pvz.NewVZCoprocessorConfiguration()
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create coprocessor configuration")
	}
	return cfg, nil
}

// NewSEPCoprocessorConfigurationWithStorage creates a SEP coprocessor configuration from a storage object.
func NewSEPCoprocessorConfigurationWithStorage(storage objectivec.IObject) (pvz.VZSEPCoprocessorConfiguration, error) {
	cfg := pvz.NewVZSEPCoprocessorConfigurationWithStorage(storage)
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create sep coprocessor configuration")
	}
	return cfg, nil
}

// NewSEPCoprocessorConfigurationWithStorageURL creates a SEP coprocessor configuration from a storage URL.
func NewSEPCoprocessorConfigurationWithStorageURL(url foundation.INSURL) (pvz.VZSEPCoprocessorConfiguration, error) {
	cfg := pvz.NewVZSEPCoprocessorConfigurationWithStorageURL(foundation.NSURLFromID(url.GetID()))
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create sep coprocessor configuration")
	}
	return cfg, nil
}

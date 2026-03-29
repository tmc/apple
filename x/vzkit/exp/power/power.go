package power

import (
	"fmt"

	pvz "github.com/tmc/apple/private/virtualization"
)

// Config describes the synthetic battery state to apply.
type Config struct {
	Charge       float64
	Connectivity int64
}

// NewSource creates a synthetic battery source with the requested values.
func NewSource(cfg Config) (pvz.VZMacSyntheticBatterySource, error) {
	source := pvz.NewVZMacSyntheticBatterySource()
	if source.ID == 0 {
		return source, fmt.Errorf("create synthetic battery source")
	}
	source.SetCharge(cfg.Charge)
	source.SetConnectivity(cfg.Connectivity)
	return source, nil
}

// NewDevice creates a macOS battery power source device backed by a synthetic source.
func NewDevice(cfg Config) (pvz.VZMacBatteryPowerSourceDeviceConfiguration, error) {
	device := pvz.NewVZMacBatteryPowerSourceDeviceConfiguration()
	if device.ID == 0 {
		return device, fmt.Errorf("create mac battery power source device configuration")
	}
	source, err := NewSource(cfg)
	if err != nil {
		return device, err
	}
	device.SetSource(&source.VZMacBatterySource)
	return device, nil
}

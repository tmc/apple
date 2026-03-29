package biometric

import (
	"fmt"

	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
)

// NewConfiguration creates the generic biometric configuration.
func NewConfiguration() (pvz.VZBiometricDeviceConfiguration, error) {
	cfg := pvz.NewVZBiometricDeviceConfiguration()
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create biometric device configuration")
	}
	return cfg, nil
}

// NewTouchIDConfiguration creates the macOS Touch ID configuration.
func NewTouchIDConfiguration() (pvz.VZMacTouchIDDeviceConfiguration, error) {
	cfg := pvz.NewVZMacTouchIDDeviceConfiguration()
	if cfg.ID == 0 {
		return cfg, fmt.Errorf("create mac touch id device configuration")
	}
	return cfg, nil
}

// NewDevice creates a biometric runtime device for the given platform object.
func NewDevice(platform objectivec.IObject) (objectivec.IObject, error) {
	cfg, err := NewConfiguration()
	if err != nil {
		return nil, err
	}
	device := cfg.BiometricDeviceWithPlatform(platform)
	if device == nil || device.GetID() == 0 {
		return nil, fmt.Errorf("create biometric device")
	}
	return device, nil
}

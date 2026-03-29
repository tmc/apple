package vzkit

import (
	vz "github.com/tmc/apple/virtualization"
	audiox "github.com/tmc/apple/x/vzkit/audio"
)

// AudioConfig holds audio device configuration settings.
type AudioConfig = audiox.Config

// CreateAudioDevice creates a VZVirtioSoundDeviceConfiguration.
func CreateAudioDevice(cfg AudioConfig) (vz.VZVirtioSoundDeviceConfiguration, error) {
	return audiox.CreateDevice(cfg)
}

// AddAudioDevice creates an audio device and adds it to the VM configuration.
func AddAudioDevice(vmConfig vz.VZVirtualMachineConfiguration, cfg AudioConfig) error {
	return audiox.AddDevice(vmConfig, cfg)
}

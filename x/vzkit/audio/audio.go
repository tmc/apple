package audio

import (
	"fmt"

	vz "github.com/tmc/apple/virtualization"
)

// Config holds audio device configuration settings.
type Config struct {
	InputEnabled  bool
	OutputEnabled bool
}

// CreateDevice creates a Virtio sound device configuration.
func CreateDevice(cfg Config) (vz.VZVirtioSoundDeviceConfiguration, error) {
	if !cfg.InputEnabled && !cfg.OutputEnabled {
		return vz.VZVirtioSoundDeviceConfiguration{}, fmt.Errorf("audio device requires at least one of input or output enabled")
	}

	soundConfig := vz.NewVZVirtioSoundDeviceConfiguration()
	if soundConfig.ID == 0 {
		return soundConfig, fmt.Errorf("create sound device configuration")
	}

	var streams []vz.VZVirtioSoundDeviceStreamConfiguration
	if cfg.OutputEnabled {
		outputStream := vz.NewVZVirtioSoundDeviceOutputStreamConfiguration()
		if outputStream.ID == 0 {
			return soundConfig, fmt.Errorf("create audio output stream configuration")
		}
		sink := vz.NewVZHostAudioOutputStreamSink()
		if sink.ID == 0 {
			return soundConfig, fmt.Errorf("create host audio output stream sink")
		}
		outputStream.SetSink(&sink)
		streams = append(streams, outputStream.VZVirtioSoundDeviceStreamConfiguration)
	}
	if cfg.InputEnabled {
		inputStream := vz.NewVZVirtioSoundDeviceInputStreamConfiguration()
		if inputStream.ID == 0 {
			return soundConfig, fmt.Errorf("create audio input stream configuration")
		}
		source := vz.NewVZHostAudioInputStreamSource()
		if source.ID == 0 {
			return soundConfig, fmt.Errorf("create host audio input stream source")
		}
		inputStream.SetSource(&source)
		streams = append(streams, inputStream.VZVirtioSoundDeviceStreamConfiguration)
	}
	soundConfig.SetStreams(streams)
	return soundConfig, nil
}

// AddDevice adds an audio device to a VM configuration.
func AddDevice(vmConfig vz.VZVirtualMachineConfiguration, cfg Config) error {
	soundDevice, err := CreateDevice(cfg)
	if err != nil {
		return fmt.Errorf("create audio device: %w", err)
	}
	vmConfig.SetAudioDevices([]vz.VZAudioDeviceConfiguration{
		soundDevice.VZAudioDeviceConfiguration,
	})
	return nil
}

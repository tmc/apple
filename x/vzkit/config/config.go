package config

import vz "github.com/tmc/apple/virtualization"

// SetStorageDevices sets a single Virtio block storage device.
func SetStorageDevices(config vz.VZVirtualMachineConfiguration, device vz.VZVirtioBlockDeviceConfiguration) {
	config.SetStorageDevices([]vz.VZStorageDeviceConfiguration{
		device.VZStorageDeviceConfiguration,
	})
}

// SetMacGraphicsDevices sets a single macOS graphics device.
func SetMacGraphicsDevices(config vz.VZVirtualMachineConfiguration, device vz.VZMacGraphicsDeviceConfiguration) {
	config.SetGraphicsDevices([]vz.VZGraphicsDeviceConfiguration{
		device.VZGraphicsDeviceConfiguration,
	})
}

// SetMacGraphicsDisplays sets a single macOS graphics display.
func SetMacGraphicsDisplays(config vz.VZMacGraphicsDeviceConfiguration, display vz.VZMacGraphicsDisplayConfiguration) {
	config.SetDisplays([]vz.VZMacGraphicsDisplayConfiguration{display})
}

// SetVirtioGraphicsDevices sets a single Virtio graphics device.
func SetVirtioGraphicsDevices(config vz.VZVirtualMachineConfiguration, device vz.VZVirtioGraphicsDeviceConfiguration) {
	config.SetGraphicsDevices([]vz.VZGraphicsDeviceConfiguration{
		vz.VZGraphicsDeviceConfigurationFromID(device.ID),
	})
}

// SetVirtioScanouts sets a single Virtio graphics scanout.
func SetVirtioScanouts(config vz.VZVirtioGraphicsDeviceConfiguration, scanout vz.VZVirtioGraphicsScanoutConfiguration) {
	config.SetScanouts([]vz.VZVirtioGraphicsScanoutConfiguration{scanout})
}

// SetNetworkDevices sets a single Virtio network device.
func SetNetworkDevices(config vz.VZVirtualMachineConfiguration, device vz.VZVirtioNetworkDeviceConfiguration) {
	config.SetNetworkDevices([]vz.VZNetworkDeviceConfiguration{
		device.VZNetworkDeviceConfiguration,
	})
}

// SetKeyboards sets a single keyboard device.
func SetKeyboards(config vz.VZVirtualMachineConfiguration, device vz.IVZKeyboardConfiguration) {
	config.SetKeyboards([]vz.VZKeyboardConfiguration{
		vz.VZKeyboardConfigurationFromID(device.GetID()),
	})
}

// SetPointingDevices sets pointing devices.
func SetPointingDevices(config vz.VZVirtualMachineConfiguration, devices []vz.IVZPointingDeviceConfiguration) {
	pointing := make([]vz.VZPointingDeviceConfiguration, 0, len(devices))
	for _, device := range devices {
		pointing = append(pointing, vz.VZPointingDeviceConfigurationFromID(device.GetID()))
	}
	config.SetPointingDevices(pointing)
}

// SetEntropyDevices sets a single Virtio entropy device.
func SetEntropyDevices(config vz.VZVirtualMachineConfiguration, device vz.VZVirtioEntropyDeviceConfiguration) {
	config.SetEntropyDevices([]vz.VZEntropyDeviceConfiguration{
		device.VZEntropyDeviceConfiguration,
	})
}

// SetAudioDevices sets a single Virtio sound device.
func SetAudioDevices(config vz.VZVirtualMachineConfiguration, device vz.VZVirtioSoundDeviceConfiguration) {
	config.SetAudioDevices([]vz.VZAudioDeviceConfiguration{
		device.VZAudioDeviceConfiguration,
	})
}

// SetSerialPorts sets a single serial port.
func SetSerialPorts(config vz.VZVirtualMachineConfiguration, device vz.VZSerialPortConfiguration) {
	config.SetSerialPorts([]vz.VZSerialPortConfiguration{device})
}

// SetMemoryBalloonDevices sets a single Virtio memory balloon device.
func SetMemoryBalloonDevices(config vz.VZVirtualMachineConfiguration, device vz.VZVirtioTraditionalMemoryBalloonDeviceConfiguration) {
	config.SetMemoryBalloonDevices([]vz.VZMemoryBalloonDeviceConfiguration{
		device.VZMemoryBalloonDeviceConfiguration,
	})
}

// SetSocketDevices sets a single Virtio socket device.
func SetSocketDevices(config vz.VZVirtualMachineConfiguration, device vz.VZVirtioSocketDeviceConfiguration) {
	config.SetSocketDevices([]vz.VZSocketDeviceConfiguration{
		device.VZSocketDeviceConfiguration,
	})
}

// SetDirectorySharingDevices sets VirtioFS directory sharing devices.
func SetDirectorySharingDevices(config vz.VZVirtualMachineConfiguration, devices []vz.VZVirtioFileSystemDeviceConfiguration) {
	sharing := make([]vz.VZDirectorySharingDeviceConfiguration, 0, len(devices))
	for _, device := range devices {
		sharing = append(sharing, vz.VZDirectorySharingDeviceConfigurationFromID(device.ID))
	}
	config.SetDirectorySharingDevices(sharing)
}

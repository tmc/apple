package windowsconfig

import (
	"fmt"

	vz "github.com/tmc/apple/virtualization"
	configx "github.com/tmc/apple/x/vzkit/config"
	displayx "github.com/tmc/apple/x/vzkit/display"
	networkx "github.com/tmc/apple/x/vzkit/network"
	storagex "github.com/tmc/apple/x/vzkit/storage"
)

// Config describes the mechanical Windows VM devices that do not own Cove
// policy such as boot media, private graphics, or serial mode.
type Config struct {
	CPUCount uint
	MemoryGB uint64

	Display []displayx.Config
	Network Network

	Keyboard      bool
	Pointing      bool
	Entropy       bool
	Sound         bool
	USBController bool
	MemoryBalloon bool
	Socket        bool
}

// Network describes the primary Virtio network device.
type Network struct {
	Config networkx.Config
	MAC    *vz.VZMACAddress
}

// Validate checks Config for shape errors.
func (c Config) Validate() error {
	if c.CPUCount == 0 {
		return fmt.Errorf("cpu count must be at least 1")
	}
	if c.MemoryGB == 0 {
		return fmt.Errorf("memory must be at least 1 GB")
	}
	for i, d := range c.Display {
		if d.Width <= 0 || d.Height <= 0 {
			return fmt.Errorf("display[%d]: width and height must be positive", i)
		}
	}
	return nil
}

// Build creates a VZVirtualMachineConfiguration and applies Config devices.
func Build(c Config) (vz.VZVirtualMachineConfiguration, error) {
	if err := c.Validate(); err != nil {
		return vz.VZVirtualMachineConfiguration{}, err
	}

	config := vz.NewVZVirtualMachineConfiguration()
	config.SetCPUCount(c.CPUCount)
	config.SetMemorySize(c.MemoryGB * 1024 * 1024 * 1024)

	if len(c.Display) > 0 {
		graphics, err := displayx.CreateVirtioGraphicsConfig(c.Display)
		if err != nil {
			return config, fmt.Errorf("create graphics config: %w", err)
		}
		configx.SetVirtioGraphicsDevices(config, graphics)
	}

	if c.Network.Config.Mode != "" && c.Network.Config.Mode != networkx.ModeNone {
		networkDevice, err := networkx.CreateDevice(c.Network.Config)
		if err != nil {
			return config, fmt.Errorf("create network device: %w", err)
		}
		if c.Network.MAC != nil && c.Network.MAC.ID != 0 {
			networkDevice.SetMACAddress(c.Network.MAC)
		}
		configx.SetNetworkDevices(config, networkDevice)
	}

	if c.Keyboard {
		configx.SetKeyboards(config, vz.NewVZUSBKeyboardConfiguration())
	}
	if c.Pointing {
		configx.SetPointingDevices(config, []vz.IVZPointingDeviceConfiguration{
			vz.NewVZUSBScreenCoordinatePointingDeviceConfiguration(),
		})
	}
	if c.Entropy {
		entropy := vz.NewVZVirtioEntropyDeviceConfiguration()
		if entropy.ID != 0 {
			configx.SetEntropyDevices(config, entropy)
		}
	}
	if c.Sound {
		sound := vz.NewVZVirtioSoundDeviceConfiguration()
		if sound.ID != 0 {
			configx.SetAudioDevices(config, sound)
		}
	}
	if c.USBController {
		storagex.EnsureUSBController(config)
	}
	if c.MemoryBalloon {
		balloon := vz.NewVZVirtioTraditionalMemoryBalloonDeviceConfiguration()
		if balloon.ID != 0 {
			configx.SetMemoryBalloonDevices(config, balloon)
		}
	}
	if c.Socket {
		socket := vz.NewVZVirtioSocketDeviceConfiguration()
		if socket.ID != 0 {
			configx.SetSocketDevices(config, socket)
		}
	}

	return config, nil
}

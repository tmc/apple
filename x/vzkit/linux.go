package vzkit

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
)

// LinuxVMConfig describes the configuration for a Linux virtual machine.
type LinuxVMConfig struct {
	CPUs     uint   // Number of CPUs
	MemoryGB uint64 // Memory in gigabytes

	DiskPath string // Path to the main disk image
	ISOPath  string // Optional ISO to attach as second block device (read-only)

	// Boot mode: set KernelPath for direct boot, leave empty for EFI boot.
	KernelPath string
	InitrdPath string
	CmdLine    string

	// EFI state directory (for efi.nvram and linux-machine.id).
	StateDir string

	// Peripherals
	Volumes  []VolumeMount // VirtioFS mounts
	Network  NetworkConfig // Network configuration
	Headless bool          // Skip graphics if true
}

// BuildLinuxVMConfig creates a VZVirtualMachineConfiguration from a LinuxVMConfig.
// It sets up platform, boot loader, storage, network, serial console, entropy,
// memory balloon, and vsock devices.
func BuildLinuxVMConfig(cfg LinuxVMConfig) (vz.VZVirtualMachineConfiguration, error) {
	config := vz.NewVZVirtualMachineConfiguration()

	config.SetCPUCount(cfg.CPUs)
	config.SetMemorySize(cfg.MemoryGB * 1024 * 1024 * 1024)

	// Platform
	platformConfig := vz.NewVZGenericPlatformConfiguration()
	machineID, err := loadOrCreateGenericMachineID(cfg.StateDir)
	if err != nil {
		return config, err
	}
	platformConfig.SetMachineIdentifier(&machineID)
	config.SetPlatform(&platformConfig.VZPlatformConfiguration)

	// Boot loader
	if cfg.KernelPath != "" {
		bl, err := createLinuxBootLoader(cfg.KernelPath, cfg.InitrdPath, cfg.CmdLine)
		if err != nil {
			return config, err
		}
		config.SetBootLoader(&bl.VZBootLoader)
	} else {
		bl, err := createEFIBootLoader(cfg.StateDir)
		if err != nil {
			return config, err
		}
		config.SetBootLoader(&bl.VZBootLoader)
	}

	// Main disk
	diskAttachment, err := CreateDiskAttachment(cfg.DiskPath, false)
	if err != nil {
		return vz.VZVirtualMachineConfiguration{}, err
	}
	mainDisk := CreateBlockDevice(diskAttachment)

	if cfg.ISOPath != "" {
		isoAttachment, err := CreateDiskAttachment(cfg.ISOPath, true)
		if err != nil {
			return config, fmt.Errorf("create ISO attachment: %w", err)
		}
		config.SetStorageDevices([]vz.VZStorageDeviceConfiguration{
			mainDisk.VZStorageDeviceConfiguration,
			CreateBlockDevice(isoAttachment).VZStorageDeviceConfiguration,
		})
	} else {
		config.SetStorageDevices([]vz.VZStorageDeviceConfiguration{
			mainDisk.VZStorageDeviceConfiguration,
		})
	}

	// Network
	if cfg.Network.Mode != NetworkModeNone {
		netDev, err := CreateNetworkDevice(cfg.Network)
		if err != nil {
			return config, fmt.Errorf("create network device: %w", err)
		}
		config.SetNetworkDevices([]vz.VZNetworkDeviceConfiguration{
			vz.VZNetworkDeviceConfigurationFromID(netDev.ID),
		})
	}

	// Entropy
	entropyConfig := vz.NewVZVirtioEntropyDeviceConfiguration()
	config.SetEntropyDevices([]vz.VZEntropyDeviceConfiguration{
		entropyConfig.VZEntropyDeviceConfiguration,
	})

	// Memory balloon
	balloonConfig := vz.NewVZVirtioTraditionalMemoryBalloonDeviceConfiguration()
	if balloonConfig.ID != 0 {
		config.SetMemoryBalloonDevices([]vz.VZMemoryBalloonDeviceConfiguration{
			balloonConfig.VZMemoryBalloonDeviceConfiguration,
		})
	}

	// Vsock
	vsockConfig := vz.NewVZVirtioSocketDeviceConfiguration()
	if vsockConfig.ID != 0 {
		config.SetSocketDevices([]vz.VZSocketDeviceConfiguration{
			vsockConfig.VZSocketDeviceConfiguration,
		})
	}

	// Serial console (stdin/stdout)
	serialConfig, err := CreateStdioSerialConsole()
	if err == nil && serialConfig.ID != 0 {
		config.SetSerialPorts([]vz.VZSerialPortConfiguration{
			serialConfig.VZSerialPortConfiguration,
		})
	}

	// VirtioFS volumes
	if len(cfg.Volumes) > 0 {
		volumeConfigs, err := CreateVirtioFSDevices(cfg.Volumes)
		if err != nil {
			return config, fmt.Errorf("create volumes: %w", err)
		}
		var fsDevices []vz.VZDirectorySharingDeviceConfiguration
		for _, vc := range volumeConfigs {
			fsDevices = append(fsDevices, vc.VZDirectorySharingDeviceConfiguration)
		}
		config.SetDirectorySharingDevices(fsDevices)
	}

	return config, nil
}

func createLinuxBootLoader(kernelPath, initrdPath, cmdLine string) (vz.VZLinuxBootLoader, error) {
	absKernel, err := filepath.Abs(kernelPath)
	if err != nil {
		return vz.VZLinuxBootLoader{}, fmt.Errorf("resolve kernel path: %w", err)
	}
	if _, err := os.Stat(absKernel); err != nil {
		return vz.VZLinuxBootLoader{}, fmt.Errorf("kernel not found: %s", absKernel)
	}

	kernelURL := foundation.NewURLFileURLWithPath(absKernel)
	bl := vz.NewLinuxBootLoaderWithKernelURL(kernelURL)
	if bl.ID == 0 {
		return vz.VZLinuxBootLoader{}, fmt.Errorf("create Linux boot loader")
	}

	if initrdPath != "" {
		absInitrd, err := filepath.Abs(initrdPath)
		if err != nil {
			return vz.VZLinuxBootLoader{}, fmt.Errorf("resolve initrd path: %w", err)
		}
		if _, err := os.Stat(absInitrd); err != nil {
			return vz.VZLinuxBootLoader{}, fmt.Errorf("initrd not found: %s", absInitrd)
		}
		initrdURL := foundation.NewURLFileURLWithPath(absInitrd)
		bl.SetInitialRamdiskURL(initrdURL)
	}

	if cmdLine == "" {
		cmdLine = "console=tty0 console=hvc0 root=/dev/vda"
	}
	bl.SetCommandLine(cmdLine)

	return bl, nil
}

func createEFIBootLoader(stateDir string) (vz.VZEFIBootLoader, error) {
	bl := vz.NewVZEFIBootLoader()
	if bl.ID == 0 {
		return bl, fmt.Errorf("create EFI boot loader")
	}

	efiStorePath := filepath.Join(stateDir, "efi.nvram")
	efiStoreURL := foundation.NewURLFileURLWithPath(efiStorePath)

	var efiStore vz.VZEFIVariableStore
	if _, err := os.Stat(efiStorePath); os.IsNotExist(err) {
		var createErr error
		efiStore, createErr = vz.NewEFIVariableStoreCreatingVariableStoreAtURLOptionsError(
			efiStoreURL, vz.VZEFIVariableStoreInitializationOptionAllowOverwrite)
		if createErr != nil {
			return bl, fmt.Errorf("create EFI variable store: %w", createErr)
		}
	} else {
		efiStore = vz.NewEFIVariableStoreWithURL(efiStoreURL)
	}

	if efiStore.ID != 0 {
		efiStore.Retain()
		bl.SetVariableStore(efiStore)
	}

	return bl, nil
}

func loadOrCreateGenericMachineID(stateDir string) (vz.VZGenericMachineIdentifier, error) {
	machineIDPath := filepath.Join(stateDir, "linux-machine.id")

	if data, err := os.ReadFile(machineIDPath); err == nil && len(data) > 0 {
		nsData := NSDataFromBytes(data)
		if nsData.ID != 0 {
			machineID := vz.NewGenericMachineIdentifierWithDataRepresentation(&nsData)
			if machineID.ID != 0 {
				return machineID, nil
			}
		}
	}

	machineID := vz.NewVZGenericMachineIdentifier()

	// Save for future use.
	dataRep := machineID.DataRepresentation()
	if dataRep.GetID() != 0 {
		nsData := foundation.NSDataFromID(dataRep.GetID())
		bytes := NSDataToBytes(nsData)
		if len(bytes) > 0 {
			os.MkdirAll(stateDir, 0755)
			os.WriteFile(machineIDPath, bytes, 0644)
		}
	}

	return machineID, nil
}

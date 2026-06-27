package vzkit

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
)

// RootfsMode selects how a Linux VM obtains its root filesystem.
type RootfsMode int

const (
	// DiskRootfs boots from a block-device disk image at DiskPath. This is the
	// zero value, so a LinuxVMConfig left unset keeps the historical behavior:
	// a main disk is always attached.
	DiskRootfs RootfsMode = iota

	// VirtioFSRootfs boots from a virtiofs-shared host directory. No main disk
	// is attached; instead one of Volumes must carry the root, tagged with
	// RootVolumeTag (default "containerfs"). BuildLinuxVMConfig errors if no
	// such volume is present, rather than silently producing a rootless boot.
	VirtioFSRootfs

	// NoRootfs attaches no disk and requires no root volume. The guest boots
	// from the kernel/initrd alone (an initrd-only image).
	NoRootfs
)

// DefaultRootVolumeTag is the virtiofs tag a VirtioFSRootfs VM uses for its
// root share when LinuxVMConfig.RootVolumeTag is empty.
const DefaultRootVolumeTag = "containerfs"

// LinuxVMConfig describes the configuration for a Linux virtual machine.
type LinuxVMConfig struct {
	CPUs     uint   // Number of CPUs
	MemoryGB uint64 // Memory in gigabytes

	// MemoryBytes, when non-zero, sets the memory size exactly and takes
	// precedence over MemoryGB. It exists for callers that compute a size with
	// sub-gigabyte granularity (e.g. a base size plus a fixed overhead).
	MemoryBytes uint64

	// RootfsMode selects disk vs. virtiofs vs. initrd-only root. The zero value
	// (DiskRootfs) preserves the original disk-image behavior.
	RootfsMode RootfsMode

	// RootVolumeTag is the virtiofs tag identifying the root share when
	// RootfsMode is VirtioFSRootfs. Empty means DefaultRootVolumeTag.
	RootVolumeTag string

	DiskPath string // Path to the main disk image (RootfsMode == DiskRootfs)
	ISOPath  string // Optional ISO to attach as second block device (read-only)

	// Boot mode: set KernelPath for direct boot, leave empty for EFI boot.
	KernelPath string
	InitrdPath string
	CmdLine    string

	// EFI state directory (for efi.nvram and linux-machine.id).
	StateDir string

	// Nested virtualization (requires macOS 15+, M3+ chip).
	// Defaults to true; set to false to disable.
	NestedVirtualization *bool

	// Peripherals
	Volumes  []VolumeMount // VirtioFS mounts
	Network  NetworkConfig // Network configuration
	Audio    *AudioConfig  // Audio device configuration (nil for no audio)
	Headless bool          // Skip graphics if true
}

// BuildLinuxVMConfig creates a VZVirtualMachineConfiguration from a LinuxVMConfig.
// It sets up platform, boot loader, storage, network, serial console, entropy,
// memory balloon, and vsock devices.
func BuildLinuxVMConfig(cfg LinuxVMConfig) (vz.VZVirtualMachineConfiguration, error) {
	config := vz.NewVZVirtualMachineConfiguration()

	config.SetCPUCount(cfg.CPUs)
	memBytes := cfg.MemoryBytes
	if memBytes == 0 {
		memBytes = cfg.MemoryGB * 1024 * 1024 * 1024
	}
	config.SetMemorySize(memBytes)

	// Platform
	platformConfig := vz.NewVZGenericPlatformConfiguration()
	machineID, err := loadOrCreateGenericMachineID(cfg.StateDir)
	if err != nil {
		return config, err
	}
	platformConfig.SetMachineIdentifier(&machineID)

	// Nested virtualization: default on, opt-out with NestedVirtualization=false.
	enableNested := cfg.NestedVirtualization == nil || *cfg.NestedVirtualization
	if enableNested && vz.GetVZGenericPlatformConfigurationClass().IsNestedVirtualizationSupported() {
		platformConfig.SetNestedVirtualizationEnabled(true)
	}

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

	// Storage. The main disk is attached only for DiskRootfs; VirtioFSRootfs and
	// NoRootfs boot without one (the root comes from a virtiofs share or the
	// initrd). An optional ISO is attached in every mode.
	var storageDevices []vz.VZStorageDeviceConfiguration

	switch cfg.RootfsMode {
	case DiskRootfs:
		diskAttachment, err := CreateDiskAttachment(cfg.DiskPath, false)
		if err != nil {
			return vz.VZVirtualMachineConfiguration{}, err
		}
		storageDevices = append(storageDevices, CreateBlockDevice(diskAttachment).VZStorageDeviceConfiguration)
	case VirtioFSRootfs:
		if err := requireRootVolume(cfg); err != nil {
			return config, err
		}
	case NoRootfs:
		// No disk, no root volume; the guest boots from kernel/initrd alone.
	default:
		return config, fmt.Errorf("unknown RootfsMode %d", cfg.RootfsMode)
	}

	if cfg.ISOPath != "" {
		isoAttachment, err := CreateDiskAttachment(cfg.ISOPath, true)
		if err != nil {
			return config, fmt.Errorf("create ISO attachment: %w", err)
		}
		storageDevices = append(storageDevices, CreateBlockDevice(isoAttachment).VZStorageDeviceConfiguration)
	}

	if len(storageDevices) > 0 {
		config.SetStorageDevices(storageDevices)
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

	// Audio
	if cfg.Audio != nil {
		if err := AddAudioDevice(config, *cfg.Audio); err != nil {
			return config, fmt.Errorf("create audio device: %w", err)
		}
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

// requireRootVolume verifies that a VirtioFSRootfs config carries a root share
// tagged with cfg.RootVolumeTag (or DefaultRootVolumeTag when empty). Without
// it the VM would boot with no root filesystem and hang in the guest, so this
// is a hard error rather than a silent misconfiguration.
func requireRootVolume(cfg LinuxVMConfig) error {
	tag := cfg.RootVolumeTag
	if tag == "" {
		tag = DefaultRootVolumeTag
	}
	for _, v := range cfg.Volumes {
		if v.Tag == tag {
			return nil
		}
	}
	return fmt.Errorf("vzkit: RootfsMode VirtioFSRootfs requires a volume tagged %q for the root filesystem", tag)
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
			machineID := vz.NewGenericMachineIdentifierWithDataRepresentation(nsData)
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

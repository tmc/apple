package vzkit

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
)

// MacVMConfig describes the configuration for a macOS virtual machine.
type MacVMConfig struct {
	CPUs     uint   // Number of CPUs
	MemoryGB uint64 // Memory in gigabytes

	DiskPath string // Path to the main disk image

	// StateDir stores persistent machine state (hardware model, machine
	// identifier, auxiliary storage). Required.
	StateDir string

	// HardwareModel is the hardware model data representation. If nil, it is
	// loaded from StateDir. For first-time setup, obtain a hardware model
	// from a VZMacOSRestoreImage and save it to StateDir before calling
	// BuildMacVMConfig.
	HardwareModel []byte

	// Peripherals
	Displays []DisplayConfig // Graphics displays (default: 1920x1200 @ 144 PPI)
	Volumes  []VolumeMount   // VirtioFS mounts
	Network  NetworkConfig   // Network configuration
	Audio    *AudioConfig    // Audio device configuration (nil for no audio)
	Headless bool            // Skip graphics, keyboard, and trackpad if true
}

// BuildMacVMConfig creates a VZVirtualMachineConfiguration from a MacVMConfig.
func BuildMacVMConfig(cfg MacVMConfig) (vz.VZVirtualMachineConfiguration, error) {
	if cfg.StateDir == "" {
		return vz.VZVirtualMachineConfiguration{}, fmt.Errorf("state directory required for macOS VM")
	}

	config := vz.NewVZVirtualMachineConfiguration()
	config.SetCPUCount(cfg.CPUs)
	config.SetMemorySize(cfg.MemoryGB * 1024 * 1024 * 1024)

	// Platform
	platformConfig, err := createMacPlatform(cfg)
	if err != nil {
		return config, fmt.Errorf("mac platform: %w", err)
	}
	config.SetPlatform(&platformConfig.VZPlatformConfiguration)

	// Boot loader
	bootLoader := vz.NewVZMacOSBootLoader()
	if bootLoader.ID == 0 {
		return config, fmt.Errorf("create macOS boot loader")
	}
	config.SetBootLoader(&bootLoader.VZBootLoader)

	// Main disk
	diskAttachment, err := CreateDiskAttachment(cfg.DiskPath, false)
	if err != nil {
		return config, err
	}
	mainDisk := CreateBlockDevice(diskAttachment)
	config.SetStorageDevices([]vz.VZStorageDeviceConfiguration{
		mainDisk.VZStorageDeviceConfiguration,
	})

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

	// Graphics, keyboard, trackpad
	if !cfg.Headless {
		graphicsConfig, err := CreateMacGraphicsConfig(cfg.Displays)
		if err != nil {
			return config, fmt.Errorf("create graphics: %w", err)
		}
		config.SetGraphicsDevices([]vz.VZGraphicsDeviceConfiguration{
			graphicsConfig.VZGraphicsDeviceConfiguration,
		})

		keyboard := vz.NewVZMacKeyboardConfiguration()
		if keyboard.ID != 0 {
			config.SetKeyboards([]vz.VZKeyboardConfiguration{
				keyboard.VZKeyboardConfiguration,
			})
		}

		trackpad := vz.NewVZMacTrackpadConfiguration()
		if trackpad.ID != 0 {
			config.SetPointingDevices([]vz.VZPointingDeviceConfiguration{
				trackpad.VZPointingDeviceConfiguration,
			})
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

func createMacPlatform(cfg MacVMConfig) (vz.VZMacPlatformConfiguration, error) {
	platformConfig := vz.NewVZMacPlatformConfiguration()
	if platformConfig.ID == 0 {
		return platformConfig, fmt.Errorf("create platform configuration")
	}

	// Hardware model
	hardwareModel, err := loadOrCreateMacHardwareModel(cfg.StateDir, cfg.HardwareModel)
	if err != nil {
		return platformConfig, fmt.Errorf("hardware model: %w", err)
	}
	platformConfig.SetHardwareModel(&hardwareModel)

	// Machine identifier
	machineID, err := loadOrCreateMacMachineID(cfg.StateDir)
	if err != nil {
		return platformConfig, fmt.Errorf("machine identifier: %w", err)
	}
	platformConfig.SetMachineIdentifier(&machineID)

	// Auxiliary storage
	auxStorage, err := loadOrCreateMacAuxiliaryStorage(cfg.StateDir, hardwareModel)
	if err != nil {
		return platformConfig, fmt.Errorf("auxiliary storage: %w", err)
	}
	platformConfig.SetAuxiliaryStorage(&auxStorage)

	return platformConfig, nil
}

func loadOrCreateMacHardwareModel(stateDir string, override []byte) (vz.VZMacHardwareModel, error) {
	hwModelPath := filepath.Join(stateDir, "mac-hardware-model.bin")

	data := override
	if data == nil {
		var err error
		data, err = os.ReadFile(hwModelPath)
		if err != nil {
			return vz.VZMacHardwareModel{}, fmt.Errorf("read %s: %w (obtain hardware model from a restore image first)", hwModelPath, err)
		}
	}

	nsData := NSDataFromBytes(data)
	if nsData.ID == 0 {
		return vz.VZMacHardwareModel{}, fmt.Errorf("create NSData for hardware model")
	}
	hwModel := vz.NewMacHardwareModelWithDataRepresentation(&nsData)
	if hwModel.ID == 0 {
		return vz.VZMacHardwareModel{}, fmt.Errorf("invalid hardware model data")
	}
	if !hwModel.Supported() {
		return vz.VZMacHardwareModel{}, fmt.Errorf("hardware model not supported on this host")
	}

	// Persist if loaded from override.
	if override != nil {
		os.MkdirAll(stateDir, 0755)
		os.WriteFile(hwModelPath, data, 0644)
	}

	return hwModel, nil
}

func loadOrCreateMacMachineID(stateDir string) (vz.VZMacMachineIdentifier, error) {
	machineIDPath := filepath.Join(stateDir, "mac-machine.id")

	if data, err := os.ReadFile(machineIDPath); err == nil && len(data) > 0 {
		nsData := NSDataFromBytes(data)
		if nsData.ID != 0 {
			machineID := vz.NewMacMachineIdentifierWithDataRepresentation(&nsData)
			if machineID.ID != 0 {
				return machineID, nil
			}
		}
	}

	machineID := vz.NewVZMacMachineIdentifier()
	if machineID.ID == 0 {
		return machineID, fmt.Errorf("create machine identifier")
	}

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

func loadOrCreateMacAuxiliaryStorage(stateDir string, hwModel vz.VZMacHardwareModel) (vz.VZMacAuxiliaryStorage, error) {
	auxStoragePath := filepath.Join(stateDir, "mac-auxiliary-storage.bin")
	auxURL := foundation.NewURLFileURLWithPath(auxStoragePath)

	if _, err := os.Stat(auxStoragePath); err == nil {
		auxStorage := vz.NewMacAuxiliaryStorageWithContentsOfURL(auxURL)
		if auxStorage.ID != 0 {
			return auxStorage, nil
		}
	}

	os.MkdirAll(stateDir, 0755)
	auxStorage, err := vz.NewMacAuxiliaryStorageCreatingStorageAtURLHardwareModelOptionsError(
		auxURL, &hwModel, vz.VZMacAuxiliaryStorageInitializationOptionAllowOverwrite)
	if err != nil {
		return vz.VZMacAuxiliaryStorage{}, fmt.Errorf("create auxiliary storage: %w", err)
	}
	if auxStorage.ID == 0 {
		return vz.VZMacAuxiliaryStorage{}, fmt.Errorf("create auxiliary storage: nil result")
	}

	return auxStorage, nil
}

// SaveMacHardwareModel saves a hardware model's data representation to the
// state directory for later use by BuildMacVMConfig.
func SaveMacHardwareModel(stateDir string, hwModel vz.VZMacHardwareModel) error {
	dataRep := hwModel.DataRepresentation()
	if dataRep.GetID() == 0 {
		return fmt.Errorf("hardware model has no data representation")
	}
	nsData := foundation.NSDataFromID(dataRep.GetID())
	bytes := NSDataToBytes(nsData)
	if len(bytes) == 0 {
		return fmt.Errorf("hardware model data is empty")
	}
	if err := os.MkdirAll(stateDir, 0755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(stateDir, "mac-hardware-model.bin"), bytes, 0644)
}

// SetupMacStateFromRestoreImage extracts the hardware model from a
// RestoreImageInfo (obtained via FetchLatestRestoreImage or
// LoadRestoreImage) and initializes the state directory for
// BuildMacVMConfig. It saves the hardware model, creates a machine
// identifier, and creates the auxiliary storage.
//
// Call this once before the first BuildMacVMConfig for a new VM. On
// subsequent boots, BuildMacVMConfig loads the persisted state directly.
func SetupMacStateFromRestoreImage(stateDir string, info RestoreImageInfo) error {
	if !info.Supported {
		return fmt.Errorf("restore image not supported on this host")
	}

	cfg := info.Image.MostFeaturefulSupportedConfiguration()
	if cfg == nil || cfg.GetID() == 0 {
		return fmt.Errorf("restore image has no supported configuration")
	}
	reqs := vz.VZMacOSConfigurationRequirementsFromID(cfg.GetID())

	hwModelIface := reqs.HardwareModel()
	if hwModelIface == nil || hwModelIface.GetID() == 0 {
		return fmt.Errorf("configuration requirements have no hardware model")
	}
	hwModel := vz.VZMacHardwareModelFromID(hwModelIface.GetID())

	if err := SaveMacHardwareModel(stateDir, hwModel); err != nil {
		return fmt.Errorf("save hardware model: %w", err)
	}

	if _, err := loadOrCreateMacMachineID(stateDir); err != nil {
		return fmt.Errorf("create machine identifier: %w", err)
	}

	if _, err := loadOrCreateMacAuxiliaryStorage(stateDir, hwModel); err != nil {
		return fmt.Errorf("create auxiliary storage: %w", err)
	}

	return nil
}

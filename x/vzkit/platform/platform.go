package platform

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/x/vzkit/storage"
)

// LoadOrCreateGenericMachineIdentifier loads a generic machine identifier from
// path or creates and saves a new one.
func LoadOrCreateGenericMachineIdentifier(path string) (vz.VZGenericMachineIdentifier, bool, error) {
	if data, err := os.ReadFile(path); err == nil && len(data) > 0 {
		nsData := storage.NSDataFromBytes(data)
		if nsData.ID != 0 {
			machineID := vz.NewGenericMachineIdentifierWithDataRepresentation(&nsData)
			if machineID.ID != 0 {
				return machineID, false, nil
			}
		}
	}

	machineID := vz.NewVZGenericMachineIdentifier()
	if machineID.ID == 0 {
		return machineID, false, fmt.Errorf("create generic machine identifier")
	}
	if err := SaveGenericMachineIdentifier(machineID, path); err != nil {
		return machineID, true, err
	}
	return machineID, true, nil
}

// SaveGenericMachineIdentifier saves a generic machine identifier to path.
func SaveGenericMachineIdentifier(machineID vz.VZGenericMachineIdentifier, path string) error {
	data := machineID.DataRepresentation()
	if data.GetID() == 0 {
		return fmt.Errorf("machine identifier has no data representation")
	}
	bytes := storage.NSDataToBytes(foundation.NSDataFromID(data.GetID()))
	if len(bytes) == 0 {
		return fmt.Errorf("machine identifier has empty data representation")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("create machine identifier directory: %w", err)
	}
	if err := os.WriteFile(path, bytes, 0644); err != nil {
		return fmt.Errorf("write machine identifier: %w", err)
	}
	return nil
}

// CreateEFIBootLoader creates an EFI boot loader with a variable store at path.
func CreateEFIBootLoader(path string) (vz.VZEFIBootLoader, bool, error) {
	bootloader := vz.NewVZEFIBootLoader()
	if bootloader.ID == 0 {
		return bootloader, false, fmt.Errorf("create EFI boot loader")
	}

	url := foundation.NewURLFileURLWithPath(path)
	if url.ID == 0 {
		return bootloader, false, fmt.Errorf("create EFI variable store URL")
	}
	url.Retain()

	var store vz.VZEFIVariableStore
	created := false
	if _, err := os.Stat(path); os.IsNotExist(err) {
		var createErr error
		store, createErr = vz.NewEFIVariableStoreCreatingVariableStoreAtURLOptionsError(
			url, vz.VZEFIVariableStoreInitializationOptionAllowOverwrite)
		if createErr != nil {
			return bootloader, false, fmt.Errorf("create EFI variable store: %w", createErr)
		}
		created = true
	} else {
		store = vz.NewEFIVariableStoreWithURL(url)
	}
	if store.ID == 0 {
		return bootloader, created, fmt.Errorf("create EFI variable store")
	}
	store.Retain()
	bootloader.SetVariableStore(&store)
	return bootloader, created, nil
}

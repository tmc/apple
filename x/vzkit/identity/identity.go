package identity

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/x/vzkit/storage"
)

// LoadMacMachineIdentifier loads a macOS machine identifier from path.
func LoadMacMachineIdentifier(path string) (vz.VZMacMachineIdentifier, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return vz.VZMacMachineIdentifier{}, fmt.Errorf("read machine identifier: %w", err)
	}
	if len(data) == 0 {
		return vz.VZMacMachineIdentifier{}, fmt.Errorf("machine identifier file is empty")
	}
	nsData := storage.NSDataFromBytes(data)
	if nsData.ID == 0 {
		return vz.VZMacMachineIdentifier{}, fmt.Errorf("machine identifier file is invalid")
	}
	machineID := vz.NewMacMachineIdentifierWithDataRepresentation(&nsData)
	if machineID.ID == 0 {
		return vz.VZMacMachineIdentifier{}, fmt.Errorf("machine identifier file is invalid")
	}
	return machineID, nil
}

// SaveMacMachineIdentifier saves a macOS machine identifier to path.
func SaveMacMachineIdentifier(machineID vz.VZMacMachineIdentifier, path string) error {
	data := machineID.DataRepresentation()
	if data.GetID() == 0 {
		return fmt.Errorf("machine identifier has no data representation")
	}
	return writeNSData(path, data.GetID(), "machine identifier")
}

// LoadOrCreateMacMachineIdentifier loads a macOS machine identifier or creates one.
func LoadOrCreateMacMachineIdentifier(path string) (vz.VZMacMachineIdentifier, bool, error) {
	machineID, err := LoadMacMachineIdentifier(path)
	if err == nil {
		return machineID, false, nil
	}
	if !os.IsNotExist(err) {
		return vz.VZMacMachineIdentifier{}, false, err
	}
	machineID = vz.NewVZMacMachineIdentifier()
	if machineID.ID == 0 {
		return machineID, false, fmt.Errorf("create machine identifier")
	}
	if err := SaveMacMachineIdentifier(machineID, path); err != nil {
		return machineID, true, err
	}
	return machineID, true, nil
}

// CreateMacMachineIdentifier creates and saves a new macOS machine identifier.
func CreateMacMachineIdentifier(path string) error {
	machineID := vz.NewVZMacMachineIdentifier()
	if machineID.ID == 0 {
		return fmt.Errorf("create machine identifier")
	}
	return SaveMacMachineIdentifier(machineID, path)
}

// LoadMacHardwareModel loads a macOS hardware model from path.
func LoadMacHardwareModel(path string) (vz.VZMacHardwareModel, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return vz.VZMacHardwareModel{}, fmt.Errorf("read hardware model: %w", err)
	}
	if len(data) == 0 {
		return vz.VZMacHardwareModel{}, fmt.Errorf("hardware model file is empty")
	}
	nsData := storage.NSDataFromBytes(data)
	if nsData.ID == 0 {
		return vz.VZMacHardwareModel{}, fmt.Errorf("hardware model file is invalid")
	}
	model := vz.NewMacHardwareModelWithDataRepresentation(&nsData)
	if model.ID == 0 {
		return vz.VZMacHardwareModel{}, fmt.Errorf("hardware model file is invalid")
	}
	if !model.Supported() {
		return vz.VZMacHardwareModel{}, fmt.Errorf("hardware model is not supported on this host")
	}
	return model, nil
}

// SaveMacHardwareModel saves a hardware model to path.
func SaveMacHardwareModel(model vz.VZMacHardwareModel, path string) error {
	data := model.DataRepresentation()
	if data == nil || data.GetID() == 0 {
		return fmt.Errorf("hardware model has no data representation")
	}
	return writeNSData(path, data.GetID(), "hardware model")
}

// LoadMacAuxiliaryStorage loads auxiliary storage from path.
func LoadMacAuxiliaryStorage(path string) (vz.VZMacAuxiliaryStorage, error) {
	url := foundation.NewURLFileURLWithPath(path)
	if url.ID == 0 {
		return vz.VZMacAuxiliaryStorage{}, fmt.Errorf("create auxiliary storage URL")
	}
	url.Retain()
	aux := vz.NewMacAuxiliaryStorageWithContentsOfURL(url)
	if aux.ID == 0 {
		return vz.VZMacAuxiliaryStorage{}, fmt.Errorf("failed to load auxiliary storage: %s", path)
	}
	aux.Retain()
	return aux, nil
}

// CreateMacAuxiliaryStorage creates auxiliary storage at path for model.
func CreateMacAuxiliaryStorage(path string, model vz.VZMacHardwareModel, overwrite bool) (vz.VZMacAuxiliaryStorage, error) {
	url := foundation.NewURLFileURLWithPath(path)
	if url.ID == 0 {
		return vz.VZMacAuxiliaryStorage{}, fmt.Errorf("create auxiliary storage URL")
	}
	url.Retain()
	options := vz.VZMacAuxiliaryStorageInitializationOptions(0)
	if overwrite {
		options = vz.VZMacAuxiliaryStorageInitializationOptionAllowOverwrite
	}
	aux, err := vz.NewMacAuxiliaryStorageCreatingStorageAtURLHardwareModelOptionsError(url, model, options)
	if err != nil {
		return vz.VZMacAuxiliaryStorage{}, fmt.Errorf("create auxiliary storage: %w", err)
	}
	if aux.ID == 0 {
		return vz.VZMacAuxiliaryStorage{}, fmt.Errorf("create auxiliary storage: nil result")
	}
	aux.Retain()
	return aux, nil
}

func writeNSData(path string, id objc.ID, name string) error {
	bytes := storage.NSDataToBytes(foundation.NSDataFromID(id))
	if len(bytes) == 0 {
		return fmt.Errorf("%s data is empty", name)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("create %s directory: %w", name, err)
	}
	if err := os.WriteFile(path, bytes, 0644); err != nil {
		return fmt.Errorf("write %s: %w", name, err)
	}
	return nil
}

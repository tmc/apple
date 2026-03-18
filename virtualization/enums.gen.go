// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageCachingMode
type VZDiskImageCachingMode int

const (
	// VZDiskImageCachingModeAutomatic: Allows the virtualization framework to automatically determine whether to enable data caching.
	VZDiskImageCachingModeAutomatic VZDiskImageCachingMode = 0
	// VZDiskImageCachingModeCached: Enables data caching.
	VZDiskImageCachingModeCached VZDiskImageCachingMode = 2
	// VZDiskImageCachingModeUncached: Disables data caching.
	VZDiskImageCachingModeUncached VZDiskImageCachingMode = 1
)

func (e VZDiskImageCachingMode) String() string {
	switch e {
	case VZDiskImageCachingModeAutomatic:
		return "VZDiskImageCachingModeAutomatic"
	case VZDiskImageCachingModeCached:
		return "VZDiskImageCachingModeCached"
	case VZDiskImageCachingModeUncached:
		return "VZDiskImageCachingModeUncached"
	default:
		return fmt.Sprintf("VZDiskImageCachingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageSynchronizationMode
type VZDiskImageSynchronizationMode int

const (
	// VZDiskImageSynchronizationModeFsync: Synchronizes data to the drive using the system’s best-effort synchronization mode.
	VZDiskImageSynchronizationModeFsync VZDiskImageSynchronizationMode = 2
	// VZDiskImageSynchronizationModeFull: Synchronizes data to the permanent storage holding the disk image.
	VZDiskImageSynchronizationModeFull VZDiskImageSynchronizationMode = 1
	// VZDiskImageSynchronizationModeNone: Disables data synchronization with the permanent storage.
	VZDiskImageSynchronizationModeNone VZDiskImageSynchronizationMode = 3
)

func (e VZDiskImageSynchronizationMode) String() string {
	switch e {
	case VZDiskImageSynchronizationModeFsync:
		return "VZDiskImageSynchronizationModeFsync"
	case VZDiskImageSynchronizationModeFull:
		return "VZDiskImageSynchronizationModeFull"
	case VZDiskImageSynchronizationModeNone:
		return "VZDiskImageSynchronizationModeNone"
	default:
		return fmt.Sprintf("VZDiskImageSynchronizationMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Virtualization/VZDiskSynchronizationMode
type VZDiskSynchronizationMode int

const (
	// VZDiskSynchronizationModeFull: Perform all synchronization operations as requested by the guest OS.
	VZDiskSynchronizationModeFull VZDiskSynchronizationMode = 0
	// VZDiskSynchronizationModeNone: Don’t synchronize the data with the permanent storage.
	VZDiskSynchronizationModeNone VZDiskSynchronizationMode = 1
)

func (e VZDiskSynchronizationMode) String() string {
	switch e {
	case VZDiskSynchronizationModeFull:
		return "VZDiskSynchronizationModeFull"
	case VZDiskSynchronizationModeNone:
		return "VZDiskSynchronizationModeNone"
	default:
		return fmt.Sprintf("VZDiskSynchronizationMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/InitializationOptions
type VZEFIVariableStoreInitializationOptions int

const (
	// VZEFIVariableStoreInitializationOptionAllowOverwrite: A Boolean value that indicates whether the framework can overwrite the EFI variable store.
	VZEFIVariableStoreInitializationOptionAllowOverwrite VZEFIVariableStoreInitializationOptions = 1
)

func (e VZEFIVariableStoreInitializationOptions) String() string {
	switch e {
	case VZEFIVariableStoreInitializationOptionAllowOverwrite:
		return "VZEFIVariableStoreInitializationOptionAllowOverwrite"
	default:
		return fmt.Sprintf("VZEFIVariableStoreInitializationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Virtualization/VZError/Code
type VZErrorCode int

const (
	// VZErrorDeviceAlreadyAttached: The device already has an attachment to the VM.
	VZErrorDeviceAlreadyAttached VZErrorCode = 30002
	// VZErrorDeviceInitializationFailure: A device initialization failure.
	VZErrorDeviceInitializationFailure VZErrorCode = 30003
	// VZErrorDeviceNotFound: The framework can’t find the device.
	VZErrorDeviceNotFound VZErrorCode = 30004
	// VZErrorInstallationFailed: An error occurred during installation.
	VZErrorInstallationFailed VZErrorCode = 10007
	// VZErrorInstallationRequiresUpdate: The VM requires a software update in order to complete the installation.
	VZErrorInstallationRequiresUpdate VZErrorCode = 10006
	// VZErrorInternal: An internal error, such as the VM unexpectedly stopping.
	VZErrorInternal VZErrorCode = 1
	// VZErrorInvalidDiskImage: An invalid disk-image error.
	VZErrorInvalidDiskImage VZErrorCode = 5
	// VZErrorInvalidRestoreImage: The restore image is invalid.
	VZErrorInvalidRestoreImage VZErrorCode = 10005
	// VZErrorInvalidRestoreImageCatalog: The restore image catalog is invalid.
	VZErrorInvalidRestoreImageCatalog VZErrorCode = 10002
	// VZErrorInvalidVirtualMachineConfiguration: An invalid configuration error.
	VZErrorInvalidVirtualMachineConfiguration VZErrorCode = 2
	// VZErrorInvalidVirtualMachineState: An invalid state error.
	VZErrorInvalidVirtualMachineState VZErrorCode = 3
	// VZErrorInvalidVirtualMachineStateTransition: An invalid state transition error.
	VZErrorInvalidVirtualMachineStateTransition VZErrorCode = 4
	// VZErrorNetworkBlockDeviceDisconnected: The network block device client disconnected from the server.
	VZErrorNetworkBlockDeviceDisconnected VZErrorCode = 20002
	// VZErrorNetworkBlockDeviceNegotiationFailed: The connection or the negotiation with the network block device server failed.
	VZErrorNetworkBlockDeviceNegotiationFailed VZErrorCode = 20001
	// VZErrorNetworkError: A network error, such as a failed connection error, occurred.
	VZErrorNetworkError VZErrorCode = 7
	// VZErrorNoSupportedRestoreImagesInCatalog: The restore image catalog has no supported restore images.
	VZErrorNoSupportedRestoreImagesInCatalog VZErrorCode = 10003
	// VZErrorNotSupported: The host computer or operating system isn’t supported.
	VZErrorNotSupported VZErrorCode = 10
	// VZErrorOperationCancelled: The code that indicates user canceled the installation of Rosetta or the app canceled the installation of a guest OS.
	VZErrorOperationCancelled VZErrorCode = 9
	// VZErrorOutOfDiskSpace: The host is out of disk space.
	VZErrorOutOfDiskSpace VZErrorCode = 8
	// VZErrorRestore: The VM failed to restore from save file.
	VZErrorRestore VZErrorCode = 12
	// VZErrorRestoreImageCatalogLoadFailed: The restore image catalog failed to load.
	VZErrorRestoreImageCatalogLoadFailed VZErrorCode = 10001
	// VZErrorRestoreImageLoadFailed: The restore image failed to load.
	VZErrorRestoreImageLoadFailed VZErrorCode = 10004
	// VZErrorSave: The VM failed to save to the save file.
	VZErrorSave VZErrorCode = 11
	// VZErrorUSBControllerNotFound: The framework can’t find the controller.
	VZErrorUSBControllerNotFound VZErrorCode = 30001
	// VZErrorVirtualMachineLimitExceeded: Unable to create an additional VM.
	VZErrorVirtualMachineLimitExceeded VZErrorCode = 6
)

func (e VZErrorCode) String() string {
	switch e {
	case VZErrorDeviceAlreadyAttached:
		return "VZErrorDeviceAlreadyAttached"
	case VZErrorDeviceInitializationFailure:
		return "VZErrorDeviceInitializationFailure"
	case VZErrorDeviceNotFound:
		return "VZErrorDeviceNotFound"
	case VZErrorInstallationFailed:
		return "VZErrorInstallationFailed"
	case VZErrorInstallationRequiresUpdate:
		return "VZErrorInstallationRequiresUpdate"
	case VZErrorInternal:
		return "VZErrorInternal"
	case VZErrorInvalidDiskImage:
		return "VZErrorInvalidDiskImage"
	case VZErrorInvalidRestoreImage:
		return "VZErrorInvalidRestoreImage"
	case VZErrorInvalidRestoreImageCatalog:
		return "VZErrorInvalidRestoreImageCatalog"
	case VZErrorInvalidVirtualMachineConfiguration:
		return "VZErrorInvalidVirtualMachineConfiguration"
	case VZErrorInvalidVirtualMachineState:
		return "VZErrorInvalidVirtualMachineState"
	case VZErrorInvalidVirtualMachineStateTransition:
		return "VZErrorInvalidVirtualMachineStateTransition"
	case VZErrorNetworkBlockDeviceDisconnected:
		return "VZErrorNetworkBlockDeviceDisconnected"
	case VZErrorNetworkBlockDeviceNegotiationFailed:
		return "VZErrorNetworkBlockDeviceNegotiationFailed"
	case VZErrorNetworkError:
		return "VZErrorNetworkError"
	case VZErrorNoSupportedRestoreImagesInCatalog:
		return "VZErrorNoSupportedRestoreImagesInCatalog"
	case VZErrorNotSupported:
		return "VZErrorNotSupported"
	case VZErrorOperationCancelled:
		return "VZErrorOperationCancelled"
	case VZErrorOutOfDiskSpace:
		return "VZErrorOutOfDiskSpace"
	case VZErrorRestore:
		return "VZErrorRestore"
	case VZErrorRestoreImageCatalogLoadFailed:
		return "VZErrorRestoreImageCatalogLoadFailed"
	case VZErrorRestoreImageLoadFailed:
		return "VZErrorRestoreImageLoadFailed"
	case VZErrorSave:
		return "VZErrorSave"
	case VZErrorUSBControllerNotFound:
		return "VZErrorUSBControllerNotFound"
	case VZErrorVirtualMachineLimitExceeded:
		return "VZErrorVirtualMachineLimitExceeded"
	default:
		return fmt.Sprintf("VZErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAvailability
type VZLinuxRosettaAvailability int

const (
	// VZLinuxRosettaAvailabilityInstalled: Rosetta is available on the host system.
	VZLinuxRosettaAvailabilityInstalled VZLinuxRosettaAvailability = 2
	// VZLinuxRosettaAvailabilityNotInstalled: Rosetta isn’t installed.
	VZLinuxRosettaAvailabilityNotInstalled VZLinuxRosettaAvailability = 1
	// VZLinuxRosettaAvailabilityNotSupported: The current hardware or software configuration doesn’t support Rosetta.
	VZLinuxRosettaAvailabilityNotSupported VZLinuxRosettaAvailability = 0
)

func (e VZLinuxRosettaAvailability) String() string {
	switch e {
	case VZLinuxRosettaAvailabilityInstalled:
		return "VZLinuxRosettaAvailabilityInstalled"
	case VZLinuxRosettaAvailabilityNotInstalled:
		return "VZLinuxRosettaAvailabilityNotInstalled"
	case VZLinuxRosettaAvailabilityNotSupported:
		return "VZLinuxRosettaAvailabilityNotSupported"
	default:
		return fmt.Sprintf("VZLinuxRosettaAvailability(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/InitializationOptions
type VZMacAuxiliaryStorageInitializationOptions int

const (
	// VZMacAuxiliaryStorageInitializationOptionAllowOverwrite: A Boolean value that indicates whether the VM can overwrite an existing auxiliary storage file.
	VZMacAuxiliaryStorageInitializationOptionAllowOverwrite VZMacAuxiliaryStorageInitializationOptions = 1
)

func (e VZMacAuxiliaryStorageInitializationOptions) String() string {
	switch e {
	case VZMacAuxiliaryStorageInitializationOptionAllowOverwrite:
		return "VZMacAuxiliaryStorageInitializationOptionAllowOverwrite"
	default:
		return fmt.Sprintf("VZMacAuxiliaryStorageInitializationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum
type VZVirtualMachineState int

const (
	// VZVirtualMachineStateError: The VM encountered an unrecoverable error.
	VZVirtualMachineStateError VZVirtualMachineState = 3
	// VZVirtualMachineStatePaused: The framework has paused a started VM.
	VZVirtualMachineStatePaused VZVirtualMachineState = 2
	// VZVirtualMachineStatePausing: The VM is pausing.
	VZVirtualMachineStatePausing VZVirtualMachineState = 5
	// VZVirtualMachineStateRestoring: The VM is restoring from a saved state.
	VZVirtualMachineStateRestoring VZVirtualMachineState = 9
	// VZVirtualMachineStateResuming: The VM is resuming from a paused state.
	VZVirtualMachineStateResuming VZVirtualMachineState = 6
	// VZVirtualMachineStateRunning: The VM is running.
	VZVirtualMachineStateRunning VZVirtualMachineState = 1
	// VZVirtualMachineStateSaving: The VM is saving its state.
	VZVirtualMachineStateSaving VZVirtualMachineState = 8
	// VZVirtualMachineStateStarting: The VM is configuring the hardware preparing to run.
	VZVirtualMachineStateStarting VZVirtualMachineState = 4
	// VZVirtualMachineStateStopped: The VM isn’t running.
	VZVirtualMachineStateStopped VZVirtualMachineState = 0
	// VZVirtualMachineStateStopping: The VM is stopping.
	VZVirtualMachineStateStopping VZVirtualMachineState = 7
)

func (e VZVirtualMachineState) String() string {
	switch e {
	case VZVirtualMachineStateError:
		return "VZVirtualMachineStateError"
	case VZVirtualMachineStatePaused:
		return "VZVirtualMachineStatePaused"
	case VZVirtualMachineStatePausing:
		return "VZVirtualMachineStatePausing"
	case VZVirtualMachineStateRestoring:
		return "VZVirtualMachineStateRestoring"
	case VZVirtualMachineStateResuming:
		return "VZVirtualMachineStateResuming"
	case VZVirtualMachineStateRunning:
		return "VZVirtualMachineStateRunning"
	case VZVirtualMachineStateSaving:
		return "VZVirtualMachineStateSaving"
	case VZVirtualMachineStateStarting:
		return "VZVirtualMachineStateStarting"
	case VZVirtualMachineStateStopped:
		return "VZVirtualMachineStateStopped"
	case VZVirtualMachineStateStopping:
		return "VZVirtualMachineStateStopping"
	default:
		return fmt.Sprintf("VZVirtualMachineState(%d)", e)
	}
}


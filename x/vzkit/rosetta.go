package vzkit

import (
	"context"
	"fmt"

	vz "github.com/tmc/apple/virtualization"
)

// RosettaAvailable reports whether Rosetta is available for Linux guests.
// Returns true only if Rosetta is installed on the host.
func RosettaAvailable() bool {
	cls := vz.GetVZLinuxRosettaDirectoryShareClass()
	return cls.Availability() == vz.VZLinuxRosettaAvailabilityInstalled
}

// RosettaSupported reports whether the current host supports Rosetta.
// Returns false on Intel Macs or other unsupported configurations.
func RosettaSupported() bool {
	cls := vz.GetVZLinuxRosettaDirectoryShareClass()
	return cls.Availability() != vz.VZLinuxRosettaAvailabilityNotSupported
}

// InstallRosetta installs Rosetta on the host. It blocks until the
// installation completes or the context is cancelled.
func InstallRosetta(ctx context.Context) error {
	cls := vz.GetVZLinuxRosettaDirectoryShareClass()
	return cls.InstallRosetta(ctx)
}

// CreateRosettaShare creates a VZLinuxRosettaDirectoryShare for use with
// Linux virtual machines. Returns an error if Rosetta is not installed.
func CreateRosettaShare() (vz.VZLinuxRosettaDirectoryShare, error) {
	share, err := vz.NewLinuxRosettaDirectoryShareWithError()
	if err != nil {
		return vz.VZLinuxRosettaDirectoryShare{}, fmt.Errorf("create rosetta share: %w", err)
	}
	if share.ID == 0 {
		return vz.VZLinuxRosettaDirectoryShare{}, fmt.Errorf("create rosetta share: nil result")
	}
	share.Retain()
	return share, nil
}

// AddRosettaSupport adds a Rosetta directory sharing device to the given
// VM configuration. The Rosetta share is exposed via VirtioFS with the tag
// "rosetta". In the Linux guest, mount it with:
//
//	mount -t virtiofs rosetta /mnt/rosetta
//
// Returns an error if Rosetta is not available on the host.
func AddRosettaSupport(config vz.VZVirtualMachineConfiguration) error {
	share, err := CreateRosettaShare()
	if err != nil {
		return err
	}

	fsConfig := vz.NewVirtioFileSystemDeviceConfigurationWithTag("rosetta")
	if fsConfig.ID == 0 {
		return fmt.Errorf("create rosetta filesystem device")
	}
	fsConfig.Retain()
	fsConfig.SetShare(&share.VZDirectoryShare)

	// Append to any existing directory sharing devices.
	existing := config.DirectorySharingDevices()
	existing = append(existing, fsConfig.VZDirectorySharingDeviceConfiguration)
	config.SetDirectorySharingDevices(existing)

	return nil
}

package vzkit

import (
	vz "github.com/tmc/apple/virtualization"
	virtiofsx "github.com/tmc/apple/x/vzkit/virtiofs"
)

// VolumeMount represents a host-to-guest volume mount configuration.
type VolumeMount = virtiofsx.Mount

// ParseVolumeSpec parses a docker-style volume specification.
func ParseVolumeSpec(spec string) (VolumeMount, error) { return virtiofsx.ParseMount(spec) }

// CreateVirtioFSDevices creates VirtioFS device configurations for the given mounts.
func CreateVirtioFSDevices(mounts []VolumeMount) ([]vz.VZVirtioFileSystemDeviceConfiguration, error) {
	return virtiofsx.CreateDevices(mounts)
}

// MacOSAutomountTag requests the macOS guest automount tag.
const MacOSAutomountTag = virtiofsx.MacOSAutomountTag

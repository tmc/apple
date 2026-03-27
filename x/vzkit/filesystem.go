package vzkit

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
)

// VolumeMount represents a host-to-guest volume mount configuration.
// Format: host_path[:tag][:ro|rw]
type VolumeMount struct {
	HostPath string // Absolute path on host
	Tag      string // VirtioFS tag (empty = auto from dir name; MacOSAutomountTag for macOS automount)
	ReadOnly bool
}

// ParseVolumeSpec parses a docker-style volume specification.
// Format: host_path[:tag][:ro|rw]
//
// Examples:
//
//	/path/to/dir                    -> auto-mount tag, read-write
//	/path/to/dir:ro                 -> auto-mount tag, read-only
//	/path/to/dir:MyVolume           -> /Volumes/MyVolume, read-write
//	/path/to/dir:MyVolume:ro        -> /Volumes/MyVolume, read-only
func ParseVolumeSpec(spec string) (VolumeMount, error) {
	parts := strings.Split(spec, ":")
	if len(parts) == 0 || parts[0] == "" {
		return VolumeMount{}, fmt.Errorf("volume spec cannot be empty")
	}

	mount := VolumeMount{
		HostPath: parts[0],
	}

	absPath, err := filepath.Abs(mount.HostPath)
	if err != nil {
		return VolumeMount{}, fmt.Errorf("invalid host path: %w", err)
	}
	// Resolve symlinks.
	if real, err := filepath.EvalSymlinks(absPath); err == nil {
		absPath = real
	}
	mount.HostPath = absPath

	info, err := os.Stat(mount.HostPath)
	if err != nil {
		return VolumeMount{}, fmt.Errorf("host path not found: %w", err)
	}
	if !info.IsDir() {
		return VolumeMount{}, fmt.Errorf("host path is not a directory: %s", mount.HostPath)
	}

	for i := 1; i < len(parts); i++ {
		part := strings.ToLower(parts[i])
		switch part {
		case "ro", "readonly":
			mount.ReadOnly = true
		case "rw", "readwrite":
			mount.ReadOnly = false
		default:
			if mount.Tag != "" {
				return VolumeMount{}, fmt.Errorf("multiple tags specified: %s and %s", mount.Tag, parts[i])
			}
			mount.Tag = parts[i]
		}
	}

	return mount, nil
}

// CreateVirtioFSDevices creates VirtioFS device configurations for the given mounts.
func CreateVirtioFSDevices(mounts []VolumeMount) ([]vz.VZVirtioFileSystemDeviceConfiguration, error) {
	if len(mounts) == 0 {
		return nil, nil
	}

	var configs []vz.VZVirtioFileSystemDeviceConfiguration
	usedTags := make(map[string]bool)

	for _, mount := range mounts {
		config, err := createVirtioFSDevice(mount, usedTags)
		if err != nil {
			return nil, fmt.Errorf("volume %s: %w", mount.HostPath, err)
		}
		configs = append(configs, config)
	}

	return configs, nil
}

// MacOSAutomountTag is a sentinel value that requests the macOS guest
// automount tag. Pass this as VolumeMount.Tag to use macOS auto-mounting
// under /Volumes/My Shared Files/. For Linux VMs, use an explicit tag
// or leave Tag empty for an auto-generated name based on the directory.
const MacOSAutomountTag = "\x00macos-automount"

func createVirtioFSDevice(mount VolumeMount, usedTags map[string]bool) (vz.VZVirtioFileSystemDeviceConfiguration, error) {
	tag := mount.Tag
	if tag == MacOSAutomountTag {
		tag = vz.GetVZVirtioFileSystemDeviceConfigurationClass().MacOSGuestAutomountTag()
		if tag == "" {
			return vz.VZVirtioFileSystemDeviceConfiguration{}, fmt.Errorf("failed to get macOSGuestAutomountTag")
		}
	} else if tag == "" {
		// Generate a tag from the directory basename for cross-platform compatibility.
		base := filepath.Base(mount.HostPath)
		tag = base
		// Deduplicate if needed.
		for i := 1; usedTags[tag]; i++ {
			tag = fmt.Sprintf("%s-%d", base, i)
		}
	}

	if usedTags[tag] {
		return vz.VZVirtioFileSystemDeviceConfiguration{}, fmt.Errorf("duplicate volume tag: %s", tag)
	}
	usedTags[tag] = true

	fsConfig := vz.NewVirtioFileSystemDeviceConfigurationWithTag(tag)
	if fsConfig.ID == 0 {
		return vz.VZVirtioFileSystemDeviceConfiguration{}, fmt.Errorf("create file system device configuration")
	}
	fsConfig.Retain()

	sharedURL := foundation.NewURLFileURLWithPath(mount.HostPath)
	sharedURL.Retain()

	sharedDir := vz.NewSharedDirectoryWithURLReadOnly(sharedURL, mount.ReadOnly)
	if sharedDir.ID == 0 {
		return vz.VZVirtioFileSystemDeviceConfiguration{}, fmt.Errorf("create shared directory")
	}
	sharedDir.Retain()

	singleShare := vz.NewSingleDirectoryShareWithDirectory(&sharedDir)
	if singleShare.ID == 0 {
		return vz.VZVirtioFileSystemDeviceConfiguration{}, fmt.Errorf("create single directory share")
	}
	singleShare.Retain()

	fsConfig.SetShare(&singleShare.VZDirectoryShare)
	return fsConfig, nil
}

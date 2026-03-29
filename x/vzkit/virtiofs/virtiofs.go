package virtiofs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
)

// Mount represents a host-to-guest volume mount configuration.
type Mount struct {
	HostPath  string
	Tag       string
	ReadOnly  bool
	MountOpts []string
}

// MacOSAutomountTag requests the macOS guest automount tag.
const MacOSAutomountTag = "\x00macos-automount"

// ParseMount parses a docker-style volume specification.
func ParseMount(spec string) (Mount, error) {
	parts := strings.Split(spec, ":")
	if len(parts) == 0 || parts[0] == "" {
		return Mount{}, fmt.Errorf("volume spec cannot be empty")
	}

	mount := Mount{HostPath: parts[0]}
	absPath, err := filepath.Abs(mount.HostPath)
	if err != nil {
		return Mount{}, fmt.Errorf("invalid host path: %w", err)
	}
	if real, err := filepath.EvalSymlinks(absPath); err == nil {
		absPath = real
	}
	mount.HostPath = absPath

	info, err := os.Stat(mount.HostPath)
	if err != nil {
		return Mount{}, fmt.Errorf("host path not found: %w", err)
	}
	if !info.IsDir() {
		return Mount{}, fmt.Errorf("host path is not a directory: %s", mount.HostPath)
	}

	for i := 1; i < len(parts); i++ {
		part := parts[i]
		lower := strings.ToLower(part)
		switch lower {
		case "ro", "readonly":
			mount.ReadOnly = true
		case "rw", "readwrite":
			mount.ReadOnly = false
		default:
			if strings.Contains(part, "=") {
				for _, opt := range strings.Split(part, ",") {
					opt = strings.TrimSpace(opt)
					if opt != "" {
						mount.MountOpts = append(mount.MountOpts, opt)
					}
				}
				continue
			}
			if mount.Tag != "" {
				return Mount{}, fmt.Errorf("multiple tags specified: %s and %s", mount.Tag, parts[i])
			}
			mount.Tag = parts[i]
		}
	}
	return mount, nil
}

// CreateDevices creates VirtioFS device configurations for mounts.
func CreateDevices(mounts []Mount) ([]vz.VZVirtioFileSystemDeviceConfiguration, error) {
	if len(mounts) == 0 {
		return nil, nil
	}
	var configs []vz.VZVirtioFileSystemDeviceConfiguration
	usedTags := make(map[string]bool)
	for _, mount := range mounts {
		config, err := createDevice(mount, usedTags)
		if err != nil {
			return nil, fmt.Errorf("volume %s: %w", mount.HostPath, err)
		}
		configs = append(configs, config)
	}
	return configs, nil
}

func createDevice(mount Mount, usedTags map[string]bool) (vz.VZVirtioFileSystemDeviceConfiguration, error) {
	tag := mount.Tag
	if tag == MacOSAutomountTag {
		tag = vz.GetVZVirtioFileSystemDeviceConfigurationClass().MacOSGuestAutomountTag()
		if tag == "" {
			return vz.VZVirtioFileSystemDeviceConfiguration{}, fmt.Errorf("failed to get macOSGuestAutomountTag")
		}
	} else if tag == "" {
		base := filepath.Base(mount.HostPath)
		tag = base
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

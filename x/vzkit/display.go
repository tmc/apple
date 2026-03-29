package vzkit

import (
	vz "github.com/tmc/apple/virtualization"
	displayx "github.com/tmc/apple/x/vzkit/display"
)

// DisplayConfig describes a single display.
type DisplayConfig = displayx.Config

// DefaultDisplayConfig returns the default macOS display configuration.
func DefaultDisplayConfig() DisplayConfig { return displayx.DefaultConfig() }

// DefaultLinuxDisplayConfig returns the default Linux display configuration.
func DefaultLinuxDisplayConfig() DisplayConfig { return displayx.DefaultLinuxConfig() }

// ParseDisplaySpec parses a display specification.
func ParseDisplaySpec(s string) (DisplayConfig, error) { return displayx.ParseSpec(s) }

// DisplaySlice is a convenience wrapper for repeated display flags.
type DisplaySlice = displayx.Slice

// CreateMacGraphicsConfig creates a macOS graphics device configuration.
func CreateMacGraphicsConfig(displays []DisplayConfig) (vz.VZMacGraphicsDeviceConfiguration, error) {
	return displayx.CreateMacGraphicsConfig(displays)
}

// CreateVirtioGraphicsConfig creates a Virtio graphics device configuration.
func CreateVirtioGraphicsConfig(displays []DisplayConfig) (vz.VZVirtioGraphicsDeviceConfiguration, error) {
	return displayx.CreateVirtioGraphicsConfig(displays)
}

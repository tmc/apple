package display

import (
	"fmt"
	"strconv"
	"strings"

	vz "github.com/tmc/apple/virtualization"
)

// Config describes a single display.
type Config struct {
	Width  int
	Height int
	PPI    int
}

// Default returns the default display configuration (1920x1200 @ 144 PPI).
func Default() Config {
	return Config{
		Width:  1920,
		Height: 1200,
		PPI:    144,
	}
}

// DefaultConfig is kept as a compatibility alias.
func DefaultConfig() Config { return Default() }

// DefaultLinux returns the default Linux display configuration (1920x1080 @ 144 PPI).
func DefaultLinux() Config {
	return Config{
		Width:  1920,
		Height: 1080,
		PPI:    144,
	}
}

// DefaultLinuxConfig is kept as a compatibility alias.
func DefaultLinuxConfig() Config { return DefaultLinux() }

// ParseSpec parses a display specification.
func ParseSpec(s string) (Config, error) {
	s = strings.ToLower(strings.TrimSpace(s))

	switch s {
	case "4k", "uhd":
		return Config{Width: 3840, Height: 2160, PPI: 144}, nil
	case "1080p", "fhd":
		return Config{Width: 1920, Height: 1080, PPI: 144}, nil
	case "720p", "hd":
		return Config{Width: 1280, Height: 720, PPI: 144}, nil
	case "retina":
		return Config{Width: 2560, Height: 1600, PPI: 227}, nil
	}

	config := Config{PPI: 144}

	if idx := strings.Index(s, "@"); idx > 0 {
		ppiStr := s[idx+1:]
		ppi, err := strconv.Atoi(ppiStr)
		if err != nil || ppi <= 0 {
			return config, fmt.Errorf("invalid PPI value: %s", ppiStr)
		}
		config.PPI = ppi
		s = s[:idx]
	}

	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return config, fmt.Errorf("invalid display format: %s (expected WIDTHxHEIGHT)", s)
	}

	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return config, fmt.Errorf("invalid width: %s", parts[0])
	}

	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return config, fmt.Errorf("invalid height: %s", parts[1])
	}

	if width < 640 || width > 7680 {
		return config, fmt.Errorf("width must be between 640 and 7680, got %d", width)
	}
	if height < 480 || height > 4320 {
		return config, fmt.Errorf("height must be between 480 and 4320, got %d", height)
	}

	config.Width = width
	config.Height = height
	return config, nil
}

// Slice is a convenience wrapper for repeated display flags.
type Slice []Config

// String implements fmt.Stringer.
func (s *Slice) String() string {
	if s == nil || len(*s) == 0 {
		return ""
	}
	var parts []string
	for _, cfg := range *s {
		parts = append(parts, fmt.Sprintf("%dx%d@%d", cfg.Width, cfg.Height, cfg.PPI))
	}
	return strings.Join(parts, ",")
}

// Set parses and appends a display spec.
func (s *Slice) Set(value string) error {
	cfg, err := ParseSpec(value)
	if err != nil {
		return err
	}
	*s = append(*s, cfg)
	return nil
}

// CreateMacGraphicsConfig creates a macOS graphics device configuration.
func CreateMacGraphicsConfig(displays []Config) (vz.VZMacGraphicsDeviceConfiguration, error) {
	if len(displays) == 0 {
		displays = []Config{Default()}
	}

	config := vz.NewVZMacGraphicsDeviceConfiguration()
	if config.ID == 0 {
		return vz.VZMacGraphicsDeviceConfiguration{}, fmt.Errorf("create mac graphics device configuration")
	}

	var graphicsDisplays []vz.VZMacGraphicsDisplayConfiguration
	for _, display := range displays {
		gd := vz.NewMacGraphicsDisplayConfigurationWithWidthInPixelsHeightInPixelsPixelsPerInch(
			display.Width, display.Height, display.PPI,
		)
		if gd.ID == 0 {
			return vz.VZMacGraphicsDeviceConfiguration{}, fmt.Errorf("create mac graphics display configuration")
		}
		graphicsDisplays = append(graphicsDisplays, gd)
	}
	config.SetDisplays(graphicsDisplays)
	return config, nil
}

// CreateVirtioGraphicsConfig creates a Virtio graphics device configuration.
func CreateVirtioGraphicsConfig(displays []Config) (vz.VZVirtioGraphicsDeviceConfiguration, error) {
	if len(displays) == 0 {
		displays = []Config{DefaultLinux()}
	}

	config := vz.NewVZVirtioGraphicsDeviceConfiguration()
	if config.ID == 0 {
		return vz.VZVirtioGraphicsDeviceConfiguration{}, fmt.Errorf("create virtio graphics device configuration")
	}

	var scanouts []vz.VZVirtioGraphicsScanoutConfiguration
	for _, display := range displays {
		scanout := vz.NewVirtioGraphicsScanoutConfigurationWithWidthInPixelsHeightInPixels(
			display.Width, display.Height,
		)
		if scanout.ID == 0 {
			return vz.VZVirtioGraphicsDeviceConfiguration{}, fmt.Errorf("create virtio scanout configuration")
		}
		scanouts = append(scanouts, scanout)
	}
	config.SetScanouts(scanouts)
	return config, nil
}

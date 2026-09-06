package jaccl

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// ConfigFromEnv reads the standalone JACCL configuration variables. It accepts
// JACCL_* names first, then the MLX_* aliases used by the upstream library.
//
// The returned Config retains the full device matrix for parity tooling. The
// native backend uses the matrix during setup: a mesh selects its first
// directed device and a ring retains each directed wire.
func ConfigFromEnv() (Config, error) {
	devicesPath, haveDevices := firstEnv("JACCL_IBV_DEVICES", "MLX_IBV_DEVICES")
	coordinator, haveCoordinator := firstEnv("JACCL_COORDINATOR", "MLX_JACCL_COORDINATOR")
	rankText, haveRank := firstEnv("JACCL_RANK", "MLX_RANK")
	if !haveDevices || !haveCoordinator || !haveRank {
		return Config{}, fmt.Errorf("JACCL_IBV_DEVICES, JACCL_COORDINATOR, and JACCL_RANK: %w", ErrInvalidConfig)
	}
	rank, err := parseJACCLRank(rankText)
	if err != nil {
		return Config{}, fmt.Errorf("JACCL_RANK %q: %w", rankText, ErrInvalidConfig)
	}
	data, err := os.ReadFile(devicesPath)
	if err != nil {
		return Config{}, fmt.Errorf("read JACCL_IBV_DEVICES: %w", err)
	}
	devices, err := parseDeviceMatrix(data)
	if err != nil {
		return Config{}, err
	}
	if rank < 0 || rank >= len(devices) {
		return Config{}, fmt.Errorf("JACCL_RANK %d for %d devices: %w", rank, len(devices), ErrInvalidConfig)
	}
	_, preferRing := firstEnv("JACCL_RING", "MLX_JACCL_RING")
	topology, err := topologyFromDevices(devices, preferRing)
	if err != nil {
		return Config{}, err
	}
	return Config{
		Rank:        rank,
		Size:        len(devices),
		GroupID:     coordinator,
		Coordinator: coordinator,
		Device:      firstLocalDevice(devices[rank]),
		Port:        1,
		Topology:    topology,
		Devices:     devices,
		PreferRing:  preferRing,
	}, nil
}

func firstEnv(names ...string) (string, bool) {
	for _, name := range names {
		if value, ok := os.LookupEnv(name); ok {
			return value, true
		}
	}
	return "", false
}

// parseJACCLRank follows the C atoi parsing used by the standalone library.
func parseJACCLRank(text string) (int, error) {
	text = strings.TrimLeft(text, " \t\n\r\v\f")
	sign := 1
	if strings.HasPrefix(text, "+") {
		text = text[1:]
	} else if strings.HasPrefix(text, "-") {
		sign = -1
		text = text[1:]
	}
	end := 0
	for end < len(text) && text[end] >= '0' && text[end] <= '9' {
		end++
	}
	if end == 0 {
		return 0, nil
	}
	value := 0
	for _, digit := range text[:end] {
		if value > (int(^uint(0)>>1)-int(digit-'0'))/10 {
			return 0, fmt.Errorf("JACCL_RANK %q: %w", text, ErrInvalidConfig)
		}
		value = value*10 + int(digit-'0')
	}
	return sign * value, nil
}

func parseDeviceMatrix(data []byte) ([][][]string, error) {
	var rows [][]json.RawMessage
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, fmt.Errorf("decode JACCL_IBV_DEVICES: %w", err)
	}
	devices := make([][][]string, len(rows))
	for source, row := range rows {
		if len(row) != len(rows) {
			return nil, fmt.Errorf("device matrix row %d has %d columns, want %d: %w", source, len(row), len(rows), ErrInvalidConfig)
		}
		devices[source] = make([][]string, len(row))
		for destination, raw := range row {
			if string(raw) == "null" {
				continue
			}
			var one string
			if err := json.Unmarshal(raw, &one); err == nil {
				devices[source][destination] = []string{one}
				continue
			}
			if err := json.Unmarshal(raw, &devices[source][destination]); err != nil {
				return nil, fmt.Errorf("device matrix %d->%d: %w", source, destination, ErrInvalidConfig)
			}
		}
	}
	return devices, nil
}

func topologyFromDevices(devices [][][]string, preferRing bool) (Topology, error) {
	if validDeviceRing(devices) && preferRing {
		return Ring(len(devices)), nil
	}
	if validDeviceMesh(devices) {
		return Mesh(len(devices)), nil
	}
	if validDeviceRing(devices) {
		return Ring(len(devices)), nil
	}
	return nil, fmt.Errorf("device matrix is neither a JACCL mesh nor ring: %w", ErrInvalidConfig)
}

func validDeviceMesh(devices [][][]string) bool {
	if len(devices) < 2 {
		return false
	}
	for source, row := range devices {
		if len(row) != len(devices) {
			return false
		}
		for destination, names := range row {
			if source == destination && len(names) != 0 || source != destination && len(names) == 0 {
				return false
			}
		}
	}
	return true
}

func validDeviceRing(devices [][][]string) bool {
	if len(devices) < 2 {
		return false
	}
	wires := len(devices[0][1])
	if wires == 0 {
		return false
	}
	for source, row := range devices {
		if len(row) != len(devices) {
			return false
		}
		left, right := (source+len(devices)-1)%len(devices), (source+1)%len(devices)
		if len(row[left]) != wires || len(row[right]) != wires {
			return false
		}
	}
	return true
}

func firstLocalDevice(row [][]string) string {
	for _, names := range row {
		if len(names) != 0 {
			return names[0]
		}
	}
	return ""
}

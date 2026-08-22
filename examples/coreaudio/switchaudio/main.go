// Command switchaudio lists and switches macOS audio input, output, and system devices.
//
// Usage:
//
//	switchaudio [-a list|get|set] [-t input|output|system] [-n device_name]
package main

import (
	"flag"
	"fmt"
	"os"
	"unsafe"

	"github.com/tmc/apple/coreaudio"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

type DeviceInfo struct {
	ID           uint32
	Name         string
	Manufacturer string
	UID          string
}

func main() {
	action := flag.String("a", "list", "Action to perform: list, get, set")
	devType := flag.String("t", "output", "Device type: input, output, system")
	name := flag.String("n", "", "Device name to set (used with -a set)")
	flag.Parse()

	if err := run(*action, *devType, *name); err != nil {
		fmt.Fprintf(os.Stderr, "switchaudio: %v\n", err)
		os.Exit(1)
	}
}

func run(action, devType, name string) error {
	pool := foundation.NewNSAutoreleasePool()
	defer pool.Drain()

	switch action {
	case "list":
		devs, err := getDevices()
		if err != nil {
			return err
		}
		for _, d := range devs {
			fmt.Printf("Device %d: %s (%s) [UID: %s]\n", d.ID, d.Name, d.Manufacturer, d.UID)
		}
	case "get":
		selector := getSelectorForType(devType)
		devID, err := getDefaultDevice(selector)
		if err != nil {
			return err
		}
		info, err := getDeviceInfo(devID)
		if err != nil {
			return err
		}
		fmt.Printf("Default %s device: %s (ID: %d)\n", devType, info.Name, devID)
	case "set":
		if name == "" {
			return fmt.Errorf("device name required for set action (-n)")
		}
		devs, err := getDevices()
		if err != nil {
			return err
		}
		var targetID uint32
		for _, d := range devs {
			if d.Name == name {
				targetID = d.ID
				break
			}
		}
		if targetID == 0 {
			return fmt.Errorf("device %q not found", name)
		}
		selector := getSelectorForType(devType)
		if err := setDefaultDevice(selector, targetID); err != nil {
			return err
		}
		fmt.Printf("Set default %s device to %s (ID: %d)\n", devType, name, targetID)
	default:
		return fmt.Errorf("unknown action %q", action)
	}
	return nil
}

func getSelectorForType(devType string) uint32 {
	switch devType {
	case "input":
		return uint32(coreaudio.KAudioHardwarePropertyDefaultInputDevice)
	case "system":
		return uint32(coreaudio.KAudioHardwarePropertyDefaultSystemOutputDevice)
	default:
		return uint32(coreaudio.KAudioHardwarePropertyDefaultOutputDevice)
	}
}

func getDefaultDevice(selector uint32) (uint32, error) {
	var devID uint32
	size := uint32(unsafe.Sizeof(devID))
	addr := coreaudio.AudioObjectPropertyAddress{
		MSelector: selector,
		MScope:    uint32(coreaudio.KAudioObjectPropertyScopeGlobalValue),
		MElement:  uint32(coreaudio.KAudioObjectPropertyElementMain),
	}
	status := coreaudio.AudioObjectGetPropertyData(uint32(coreaudio.KAudioObjectSystemObject), &addr, 0, nil, &size, unsafe.Pointer(&devID))
	if status != 0 {
		return 0, fmt.Errorf("AudioObjectGetPropertyData failed: 0x%x", status)
	}
	return devID, nil
}

func setDefaultDevice(selector uint32, devID uint32) error {
	size := uint32(unsafe.Sizeof(devID))
	addr := coreaudio.AudioObjectPropertyAddress{
		MSelector: selector,
		MScope:    uint32(coreaudio.KAudioObjectPropertyScopeGlobalValue),
		MElement:  uint32(coreaudio.KAudioObjectPropertyElementMain),
	}
	status := coreaudio.AudioObjectSetPropertyData(uint32(coreaudio.KAudioObjectSystemObject), &addr, 0, nil, size, unsafe.Pointer(&devID))
	if status != 0 {
		return fmt.Errorf("AudioObjectSetPropertyData failed: 0x%x", status)
	}
	return nil
}

func getDevices() ([]DeviceInfo, error) {
	addr := coreaudio.AudioObjectPropertyAddress{
		MSelector: uint32(coreaudio.KAudioHardwarePropertyDevicesValue),
		MScope:    uint32(coreaudio.KAudioObjectPropertyScopeGlobalValue),
		MElement:  uint32(coreaudio.KAudioObjectPropertyElementMain),
	}

	var dataSize uint32
	status := coreaudio.AudioObjectGetPropertyDataSize(uint32(coreaudio.KAudioObjectSystemObject), &addr, 0, nil, &dataSize)
	if status != 0 {
		return nil, fmt.Errorf("AudioObjectGetPropertyDataSize failed: 0x%x", status)
	}

	count := dataSize / uint32(unsafe.Sizeof(uint32(0)))
	deviceIDs := make([]uint32, count)
	status = coreaudio.AudioObjectGetPropertyData(uint32(coreaudio.KAudioObjectSystemObject), &addr, 0, nil, &dataSize, unsafe.Pointer(&deviceIDs[0]))
	if status != 0 {
		return nil, fmt.Errorf("AudioObjectGetPropertyData failed: 0x%x", status)
	}

	var list []DeviceInfo
	for _, id := range deviceIDs {
		info, err := getDeviceInfo(id)
		if err == nil {
			list = append(list, info)
		}
	}
	return list, nil
}

func getDeviceInfo(devID uint32) (DeviceInfo, error) {
	info := DeviceInfo{ID: devID}
	info.Name = getStringProperty(devID, uint32(coreaudio.KAudioObjectPropertyName))
	info.Manufacturer = getStringProperty(devID, uint32(coreaudio.KAudioObjectPropertyManufacturer))
	info.UID = getStringProperty(devID, uint32(coreaudio.KAudioDevicePropertyDeviceUID))
	return info, nil
}

func getStringProperty(devID uint32, selector uint32) string {
	addr := coreaudio.AudioObjectPropertyAddress{
		MSelector: selector,
		MScope:    uint32(coreaudio.KAudioObjectPropertyScopeGlobalValue),
		MElement:  uint32(coreaudio.KAudioObjectPropertyElementMain),
	}
	var cfStr corefoundation.CFStringRef
	size := uint32(unsafe.Sizeof(cfStr))
	status := coreaudio.AudioObjectGetPropertyData(devID, &addr, 0, nil, &size, unsafe.Pointer(&cfStr))
	if status != 0 || cfStr == 0 {
		return ""
	}
	nsStr := foundation.NSStringFromID(objc.ID(cfStr))
	return nsStr.String()
}

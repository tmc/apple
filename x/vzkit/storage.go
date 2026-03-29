package vzkit

import (
	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
	storagex "github.com/tmc/apple/x/vzkit/storage"
)

// CreateDiskAttachment creates a VZDiskImageStorageDeviceAttachment from a file path.
func CreateDiskAttachment(path string, readOnly bool) (vz.VZDiskImageStorageDeviceAttachment, error) {
	return storagex.CreateDiskAttachment(path, readOnly)
}

// CreateDiskImage creates a sparse disk image of the given size in gigabytes.
func CreateDiskImage(path string, sizeGB uint64) error { return storagex.CreateDiskImage(path, sizeGB) }

// CreateSerialConsole creates a VZVirtioConsoleDeviceSerialPortConfiguration.
func CreateSerialConsole(readFd, writeFd int) (vz.VZVirtioConsoleDeviceSerialPortConfiguration, error) {
	return storagex.CreateSerialConsole(readFd, writeFd)
}

// CreateStdioSerialConsole creates a serial console connected to stdin/stdout.
func CreateStdioSerialConsole() (vz.VZVirtioConsoleDeviceSerialPortConfiguration, error) {
	return storagex.CreateStdioSerialConsole()
}

// CreateBlockDevice creates a VZVirtioBlockDeviceConfiguration from a disk attachment.
func CreateBlockDevice(attachment vz.VZDiskImageStorageDeviceAttachment) vz.VZVirtioBlockDeviceConfiguration {
	return storagex.CreateBlockDevice(attachment)
}

// CreateDirectoryShare creates a VZSingleDirectoryShare for VirtioFS.
func CreateDirectoryShare(path string, readOnly bool) (vz.VZSingleDirectoryShare, error) {
	return storagex.CreateDirectoryShare(path, readOnly)
}

// NSDataToBytes extracts the bytes from an NSData object into a Go-owned slice.
func NSDataToBytes(data foundation.NSData) []byte { return storagex.NSDataToBytes(data) }

// NSDataFromBytes creates an NSData object from Go bytes.
func NSDataFromBytes(data []byte) foundation.NSData { return storagex.NSDataFromBytes(data) }

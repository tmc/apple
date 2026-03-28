package vzkit

import (
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
)

// CreateDiskAttachment creates a VZDiskImageStorageDeviceAttachment from a file path.
func CreateDiskAttachment(path string, readOnly bool) (vz.VZDiskImageStorageDeviceAttachment, error) {
	url := foundation.NewURLFileURLWithPath(path)
	rv, err := vz.NewDiskImageStorageDeviceAttachmentWithURLReadOnlyError(url, readOnly)
	if err != nil {
		return vz.VZDiskImageStorageDeviceAttachment{}, fmt.Errorf("create disk attachment: %w", err)
	}
	if rv.ID != 0 {
		rv.Retain()
	}
	return rv, nil
}

// CreateDiskImage creates a sparse disk image of the given size in gigabytes.
func CreateDiskImage(path string, sizeGB uint64) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()
	return f.Truncate(int64(sizeGB) * 1024 * 1024 * 1024)
}

// CreateSerialConsole creates a VZVirtioConsoleDeviceSerialPortConfiguration
// connected to the given read and write file descriptors.
func CreateSerialConsole(readFd, writeFd int) (vz.VZVirtioConsoleDeviceSerialPortConfiguration, error) {
	readHandle := foundation.NewFileHandleWithFileDescriptor(readFd)
	readHandle.Retain()
	writeHandle := foundation.NewFileHandleWithFileDescriptor(writeFd)
	writeHandle.Retain()

	attachment := vz.NewFileHandleSerialPortAttachmentWithFileHandleForReadingFileHandleForWriting(readHandle, writeHandle)
	if attachment.ID == 0 {
		return vz.VZVirtioConsoleDeviceSerialPortConfiguration{}, fmt.Errorf("create serial port attachment")
	}
	attachment.Retain()

	serialConfig := vz.NewVZVirtioConsoleDeviceSerialPortConfiguration()
	if serialConfig.ID == 0 {
		return vz.VZVirtioConsoleDeviceSerialPortConfiguration{}, fmt.Errorf("create serial port configuration")
	}
	serialConfig.SetAttachment(attachment)
	return serialConfig, nil
}

// CreateStdioSerialConsole creates a serial console connected to stdin/stdout.
func CreateStdioSerialConsole() (vz.VZVirtioConsoleDeviceSerialPortConfiguration, error) {
	return CreateSerialConsole(int(os.Stdin.Fd()), int(os.Stdout.Fd()))
}

// CreateBlockDevice creates a VZVirtioBlockDeviceConfiguration from a disk attachment.
func CreateBlockDevice(attachment vz.VZDiskImageStorageDeviceAttachment) vz.VZVirtioBlockDeviceConfiguration {
	config := vz.NewVirtioBlockDeviceConfigurationWithAttachment(attachment)
	config.Retain()
	return config
}

// CreateDirectoryShare creates a VZSingleDirectoryShare for VirtioFS.
func CreateDirectoryShare(path string, readOnly bool) (vz.VZSingleDirectoryShare, error) {
	dirURL := foundation.NewURLFileURLWithPath(path)
	if dirURL.ID == 0 {
		return vz.VZSingleDirectoryShare{}, fmt.Errorf("create URL for path %s", path)
	}

	sharedDir := vz.NewSharedDirectoryWithURLReadOnly(dirURL, readOnly)
	if sharedDir.ID == 0 {
		return vz.VZSingleDirectoryShare{}, fmt.Errorf("create shared directory for %s", path)
	}
	sharedDir.Retain()

	share := vz.NewSingleDirectoryShareWithDirectory(sharedDir)
	if share.ID != 0 {
		share.Retain()
	}
	return share, nil
}

// NSDataToBytes extracts the bytes from an NSData object into a Go-owned slice.
func NSDataToBytes(data foundation.NSData) []byte {
	length := data.Length()
	if length == 0 {
		return nil
	}
	bytesPtr := data.Bytes()
	if bytesPtr == nil {
		return nil
	}
	src := unsafe.Slice((*byte)(bytesPtr), length)
	dst := make([]byte, length)
	copy(dst, src)
	runtime.KeepAlive(data)
	return dst
}

// NSDataFromBytes creates an NSData object from Go bytes.
func NSDataFromBytes(data []byte) foundation.NSData {
	if len(data) == 0 {
		return foundation.NSData{}
	}
	return foundation.NewDataWithBytesLength(data)
}

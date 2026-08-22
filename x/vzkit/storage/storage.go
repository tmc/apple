package storage

import (
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
)

type CachePolicy int

const (
	CacheDurable CachePolicy = iota
	CacheEphemeral
	CacheReadOnly
)

func DiskImageModes(policy CachePolicy) (vz.VZDiskImageCachingMode, vz.VZDiskImageSynchronizationMode, error) {
	switch policy {
	case CacheDurable:
		return vz.VZDiskImageCachingModeCached, vz.VZDiskImageSynchronizationModeFsync, nil
	case CacheEphemeral:
		return vz.VZDiskImageCachingModeCached, vz.VZDiskImageSynchronizationModeNone, nil
	case CacheReadOnly:
		return vz.VZDiskImageCachingModeAutomatic, vz.VZDiskImageSynchronizationModeFull, nil
	default:
		return 0, 0, fmt.Errorf("unknown disk cache policy %d", policy)
	}
}

func NewDiskImageAttachment(url foundation.INSURL, readOnly bool, policy CachePolicy) (vz.VZDiskImageStorageDeviceAttachment, error) {
	caching, sync, err := DiskImageModes(policy)
	if err != nil {
		return vz.VZDiskImageStorageDeviceAttachment{}, err
	}
	return NewDiskImageAttachmentWithSynchronizationMode(url, readOnly, caching, sync)
}

func NewDiskImageAttachmentWithSynchronizationMode(url foundation.INSURL, readOnly bool, caching vz.VZDiskImageCachingMode, sync vz.VZDiskImageSynchronizationMode) (vz.VZDiskImageStorageDeviceAttachment, error) {
	return vz.NewDiskImageStorageDeviceAttachmentWithURLReadOnlyCachingModeSynchronizationModeError(foundation.NSURLFromID(url.GetID()), readOnly, caching, sync)
}

// CreateDiskAttachment creates a disk image attachment for path.
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

// CreateSerialConsole creates a serial console wired to readFd and writeFd.
func CreateSerialConsole(readFd, writeFd int) (vz.VZVirtioConsoleDeviceSerialPortConfiguration, error) {
	readHandle := foundation.NewFileHandleWithFileDescriptor(int32(readFd))
	readHandle.Retain()
	writeHandle := foundation.NewFileHandleWithFileDescriptor(int32(writeFd))
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

// CreateBlockDevice creates a Virtio block device configuration.
func CreateBlockDevice(attachment vz.VZDiskImageStorageDeviceAttachment) vz.VZVirtioBlockDeviceConfiguration {
	config := vz.NewVirtioBlockDeviceConfigurationWithAttachment(attachment)
	config.Retain()
	return config
}

// CreateBlockDeviceWithAttachment creates a Virtio block device configuration.
func CreateBlockDeviceWithAttachment(attachment vz.VZStorageDeviceAttachment) (vz.VZVirtioBlockDeviceConfiguration, error) {
	config := vz.NewVirtioBlockDeviceConfigurationWithAttachment(&attachment)
	if config.ID == 0 {
		return vz.VZVirtioBlockDeviceConfiguration{}, fmt.Errorf("create virtio block storage device")
	}
	config.Retain()
	return config, nil
}

// CreateNVMeDeviceWithAttachment creates an NVMe storage device configuration.
func CreateNVMeDeviceWithAttachment(attachment vz.VZStorageDeviceAttachment) (vz.VZNVMExpressControllerDeviceConfiguration, error) {
	config := vz.NewNVMExpressControllerDeviceConfigurationWithAttachment(attachment)
	if config.ID == 0 {
		return vz.VZNVMExpressControllerDeviceConfiguration{}, fmt.Errorf("create NVMe storage device")
	}
	config.Retain()
	return config, nil
}

// CreateUSBMassStorageDeviceWithAttachment creates a USB mass-storage device configuration.
func CreateUSBMassStorageDeviceWithAttachment(attachment vz.VZStorageDeviceAttachment) (vz.VZUSBMassStorageDeviceConfiguration, error) {
	config := vz.NewUSBMassStorageDeviceConfigurationWithAttachment(&attachment)
	if config.ID == 0 {
		return vz.VZUSBMassStorageDeviceConfiguration{}, fmt.Errorf("create USB mass storage device")
	}
	config.Retain()
	return config, nil
}

// AppendStorageDevices appends devices to a VM configuration's storage devices.
func AppendStorageDevices(config vz.VZVirtualMachineConfiguration, devices ...vz.VZStorageDeviceConfiguration) {
	existing := config.StorageDevices()
	all := make([]vz.VZStorageDeviceConfiguration, 0, len(existing)+len(devices))
	for _, dev := range existing {
		all = append(all, vz.VZStorageDeviceConfigurationFromID(dev.GetID()))
	}
	all = append(all, devices...)
	config.SetStorageDevices(all)
}

// EnsureUSBController adds a default XHCI controller when none is configured.
func EnsureUSBController(config vz.VZVirtualMachineConfiguration) {
	if len(config.UsbControllers()) > 0 {
		return
	}
	controller := vz.NewVZXHCIControllerConfiguration()
	if controller.ID == 0 {
		return
	}
	config.SetUsbControllers([]vz.VZUSBControllerConfiguration{
		vz.VZUSBControllerConfigurationFromID(controller.ID),
	})
}

// CreateDirectoryShare creates a single-directory share for VirtioFS.
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

// NSDataToBytes copies bytes from an NSData object into a Go slice.
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

// NSDataFromBytes copies Go bytes into a new NSData object.
func NSDataFromBytes(data []byte) foundation.NSData {
	if len(data) == 0 {
		return foundation.NSData{}
	}
	return foundation.NewDataWithBytesLength(data)
}

// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZDiskImageStorageDeviceAttachment] class.
var (
	_VZDiskImageStorageDeviceAttachmentClass     VZDiskImageStorageDeviceAttachmentClass
	_VZDiskImageStorageDeviceAttachmentClassOnce sync.Once
)

func getVZDiskImageStorageDeviceAttachmentClass() VZDiskImageStorageDeviceAttachmentClass {
	_VZDiskImageStorageDeviceAttachmentClassOnce.Do(func() {
		_VZDiskImageStorageDeviceAttachmentClass = VZDiskImageStorageDeviceAttachmentClass{class: objc.GetClass("VZDiskImageStorageDeviceAttachment")}
	})
	return _VZDiskImageStorageDeviceAttachmentClass
}

// GetVZDiskImageStorageDeviceAttachmentClass returns the class object for VZDiskImageStorageDeviceAttachment.
func GetVZDiskImageStorageDeviceAttachmentClass() VZDiskImageStorageDeviceAttachmentClass {
	return getVZDiskImageStorageDeviceAttachmentClass()
}

type VZDiskImageStorageDeviceAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDiskImageStorageDeviceAttachmentClass) Alloc() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A device that stores content in a disk image.
//
// # Overview
// 
// Use a [VZDiskImageStorageDeviceAttachment] object to manage the storage for
// a disk in a virtual machine (VM). The guest operating system sees the
// storage as a disk, and when the guest operating system writes files to the
// disk, the virtual machine stores the files in the disk image you provide.
// 
// The virtualization framework supports two disk image formats:
// 
// RAW disk images: A file that’s the requested size of the VM disk, this
// format results in a 1-to-1 mapping between the offsets in the file and the
// offsets in the VM disk. ASIF disk images: (ASIF) files transfer more
// efficiently between hosts or disks because their intrinsic structure
// doesn’t depend on the host file system’s capabilities. The size the
// ASIF file takes on the file system is proportional to the actual data
// stored in the disk image.
// 
// # Create the disk image
// 
// Create disk images using `diskutil`, a command-line utility you can access
// through the Terminal app. The `diskutil` utility performs a number of
// functions. The command uses a specific structure that tells the app what
// function it’s performing, including many arguments documented in its
// online manual page documents. For more information, see the `diskutil`
// manual page by using the command `man diskutil` in Terminal.
// 
// The command to creates disk image files you use with VMs starts with
// `diskutil image create blank`, followed by three required arguments that
// describe the structure, size, and location of the disk image. The general
// format of the command follows the pattern shown here:
// 
// Use these parameters to specify the configuration of the disk image:
// 
// `--format`: An argument that describes the specific file-system format to
// use, either [RAW] or [ASIF]. `diskutil` supports other file-system formats
// as well, not all of which Virtualization supports. For more information,
// see the `diskutil` manual page by using the command `man diskutil` in
// Terminal. `--fs`: An argument that specifies the file system to use when
// creating the disk image. `--size`: An argument that describes the size of
// the file system to create, usually specified in gigabytes with the suffix
// [GiB]; for example, `256GiB` specifies a 256 GB image file size. You can
// also specify sizes in megabytes, using the MiB notation; for example,
// `256MiB`. `IMAGE_PATH`: The file system path that represents the location
// where `diskutil` creates the [RAW] or [ASIF] file.
// 
// To create a RAW disk image in the file system of the host computer, open
// Terminal and run the following command, replacing the placeholder values
// with your own, for example:
// 
// A best practice is to give disk images a common and easily recognizable
// file extension, such as `XCUIElementTypeImg`; for example `VM_Image.Img()`.
// 
// To create an ASIF image, replace [RAW] with [ASIF], as shown here:
// 
// You can execute `diskutil` commands under app control to create disk images
// programmatically.
// 
// You can also create [ASIF] or [RAW] disk images interactively using Disk
// Utility, a utility app that Apple provides with macOS. To create volumes
// with Disk Utility, open the app, then select File > New Image > Blank
// Image. Configure the form by selecting the disk format (either [ASIF] or
// [RAW]), volume size, and location for the new volume, and then click Save.
// For more information on using Disk Utility, see [Disk Utility User Guide].
// 
// Alternatively, you can create a raw disk image programmatically using the
// UNIX `open()` and `truncate()` standard library functions as shown here:
// 
// # Initialize the disk image
// 
// After creating the disk image, you use it to initialize a VM’s
// [VZDiskImageStorageDeviceAttachment] object. Use the attachment object to
// configure the [VZVirtioBlockDeviceConfiguration] object that you add to
// your virtual machine’s configuration, as shown here:
//
// [Disk Utility User Guide]: https://support.apple.com/guide/disk-utility/welcome/mac
//
// # Creating the attachment point
//
//   - [VZDiskImageStorageDeviceAttachment.InitWithURLReadOnlyError]: Creates the attachment object from the specified disk image.
//   - [VZDiskImageStorageDeviceAttachment.InitWithURLReadOnlyCachingModeSynchronizationModeError]: Initialize the attachment from a local file URL.
//
// # Getting the disk image details
//
//   - [VZDiskImageStorageDeviceAttachment.URL]: The URL of the underlying disk image.
//   - [VZDiskImageStorageDeviceAttachment.ReadOnly]: A Boolean value that indicates whether the underlying disk image is read-only.
//   - [VZDiskImageStorageDeviceAttachment.CachingMode]: The current cacheing mode for the virtual disk image.
//   - [VZDiskImageStorageDeviceAttachment.SynchronizationMode]: The mode in which the disk image synchronizes data with the underlying storage device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment
type VZDiskImageStorageDeviceAttachment struct {
	VZStorageDeviceAttachment
}

// VZDiskImageStorageDeviceAttachmentFromID constructs a [VZDiskImageStorageDeviceAttachment] from an objc.ID.
//
// A device that stores content in a disk image.
func VZDiskImageStorageDeviceAttachmentFromID(id objc.ID) VZDiskImageStorageDeviceAttachment {
	return VZDiskImageStorageDeviceAttachment{VZStorageDeviceAttachment: VZStorageDeviceAttachmentFromID(id)}
}
// NOTE: VZDiskImageStorageDeviceAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZDiskImageStorageDeviceAttachment] class.
//
// # Creating the attachment point
//
//   - [IVZDiskImageStorageDeviceAttachment.InitWithURLReadOnlyError]: Creates the attachment object from the specified disk image.
//   - [IVZDiskImageStorageDeviceAttachment.InitWithURLReadOnlyCachingModeSynchronizationModeError]: Initialize the attachment from a local file URL.
//
// # Getting the disk image details
//
//   - [IVZDiskImageStorageDeviceAttachment.URL]: The URL of the underlying disk image.
//   - [IVZDiskImageStorageDeviceAttachment.ReadOnly]: A Boolean value that indicates whether the underlying disk image is read-only.
//   - [IVZDiskImageStorageDeviceAttachment.CachingMode]: The current cacheing mode for the virtual disk image.
//   - [IVZDiskImageStorageDeviceAttachment.SynchronizationMode]: The mode in which the disk image synchronizes data with the underlying storage device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment
type IVZDiskImageStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment

	// Topic: Creating the attachment point

	// Creates the attachment object from the specified disk image.
	InitWithURLReadOnlyError(url foundation.INSURL, readOnly bool) (VZDiskImageStorageDeviceAttachment, error)
	// Initialize the attachment from a local file URL.
	InitWithURLReadOnlyCachingModeSynchronizationModeError(url foundation.INSURL, readOnly bool, cachingMode VZDiskImageCachingMode, synchronizationMode VZDiskImageSynchronizationMode) (VZDiskImageStorageDeviceAttachment, error)

	// Topic: Getting the disk image details

	// The URL of the underlying disk image.
	URL() foundation.INSURL
	// A Boolean value that indicates whether the underlying disk image is read-only.
	ReadOnly() bool
	// The current cacheing mode for the virtual disk image.
	CachingMode() VZDiskImageCachingMode
	// The mode in which the disk image synchronizes data with the underlying storage device.
	SynchronizationMode() VZDiskImageSynchronizationMode
}

// Init initializes the instance.
func (d VZDiskImageStorageDeviceAttachment) Init() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VZDiskImageStorageDeviceAttachment) Autorelease() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDiskImageStorageDeviceAttachment creates a new VZDiskImageStorageDeviceAttachment instance.
func NewVZDiskImageStorageDeviceAttachment() VZDiskImageStorageDeviceAttachment {
	class := getVZDiskImageStorageDeviceAttachmentClass()
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize the attachment from a local file URL.
//
// url: Local file URL to the disk image in RAW format.
//
// readOnly: If `true`, the device attachment is read-only, otherwise the device can
// write data to the disk image.
//
// cachingMode: The cacheing mode from one of the available [VZDiskImageCachingMode]
// options.
// //
// [VZDiskImageCachingMode]: https://developer.apple.com/documentation/Virtualization/VZDiskImageCachingMode
//
// synchronizationMode: How the disk image synchronizes with the underlying storage when the guest
// operating system flushes data, described by one of the available
// [VZDiskImageSynchronizationMode] modes.
// //
// [VZDiskImageSynchronizationMode]: https://developer.apple.com/documentation/Virtualization/VZDiskImageSynchronizationMode
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/init(url:readOnly:cachingMode:synchronizationMode:)
func NewDiskImageStorageDeviceAttachmentWithURLReadOnlyCachingModeSynchronizationModeError(url foundation.INSURL, readOnly bool, cachingMode VZDiskImageCachingMode, synchronizationMode VZDiskImageSynchronizationMode) (VZDiskImageStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	instance := getVZDiskImageStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:readOnly:cachingMode:synchronizationMode:error:"), url, readOnly, cachingMode, synchronizationMode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZDiskImageStorageDeviceAttachmentFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return VZDiskImageStorageDeviceAttachmentFromID(rv), nil
}

// Creates the attachment object from the specified disk image.
//
// url: A URL that points to a local disk image in RAW format.
//
// readOnly: A Boolean that indicates whether to configure the disk image as read-only.
// Specify [true] to prevent the guest operating system from writing to the
// disk image, and [false] to allow writing.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// In Swift the methods returns an attachment object; in Objective-C the
// methods returns an attachment object on success, or `nil` if an error
// occurred
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/init(url:readOnly:)
func NewDiskImageStorageDeviceAttachmentWithURLReadOnlyError(url foundation.INSURL, readOnly bool) (VZDiskImageStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	instance := getVZDiskImageStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:readOnly:error:"), url, readOnly, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZDiskImageStorageDeviceAttachmentFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return VZDiskImageStorageDeviceAttachmentFromID(rv), nil
}

// Creates the attachment object from the specified disk image.
//
// url: A URL that points to a local disk image in RAW format.
//
// readOnly: A Boolean that indicates whether to configure the disk image as read-only.
// Specify [true] to prevent the guest operating system from writing to the
// disk image, and [false] to allow writing.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// In Swift the methods returns an attachment object; in Objective-C the
// methods returns an attachment object on success, or `nil` if an error
// occurred
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/init(url:readOnly:)
func (d VZDiskImageStorageDeviceAttachment) InitWithURLReadOnlyError(url foundation.INSURL, readOnly bool) (VZDiskImageStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:readOnly:error:"), url, readOnly, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZDiskImageStorageDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZDiskImageStorageDeviceAttachmentFromID(rv), nil

}

// Initialize the attachment from a local file URL.
//
// url: Local file URL to the disk image in RAW format.
//
// readOnly: If `true`, the device attachment is read-only, otherwise the device can
// write data to the disk image.
//
// cachingMode: The cacheing mode from one of the available [VZDiskImageCachingMode]
// options.
// //
// [VZDiskImageCachingMode]: https://developer.apple.com/documentation/Virtualization/VZDiskImageCachingMode
//
// synchronizationMode: How the disk image synchronizes with the underlying storage when the guest
// operating system flushes data, described by one of the available
// [VZDiskImageSynchronizationMode] modes.
// //
// [VZDiskImageSynchronizationMode]: https://developer.apple.com/documentation/Virtualization/VZDiskImageSynchronizationMode
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/init(url:readOnly:cachingMode:synchronizationMode:)
func (d VZDiskImageStorageDeviceAttachment) InitWithURLReadOnlyCachingModeSynchronizationModeError(url foundation.INSURL, readOnly bool, cachingMode VZDiskImageCachingMode, synchronizationMode VZDiskImageSynchronizationMode) (VZDiskImageStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:readOnly:cachingMode:synchronizationMode:error:"), url, readOnly, cachingMode, synchronizationMode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZDiskImageStorageDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZDiskImageStorageDeviceAttachmentFromID(rv), nil

}

// The URL of the underlying disk image.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/url
func (d VZDiskImageStorageDeviceAttachment) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the underlying disk image is
// read-only.
//
// # Discussion
// 
// If the value of this property is [true], the guest operating system may
// read the contents of the disk image, but may not write to it.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/isReadOnly
func (d VZDiskImageStorageDeviceAttachment) ReadOnly() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isReadOnly"))
	return rv
}

// The current cacheing mode for the virtual disk image.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/cachingMode
func (d VZDiskImageStorageDeviceAttachment) CachingMode() VZDiskImageCachingMode {
	rv := objc.Send[VZDiskImageCachingMode](d.ID, objc.Sel("cachingMode"))
	return VZDiskImageCachingMode(rv)
}

// The mode in which the disk image synchronizes data with the underlying
// storage device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/synchronizationMode
func (d VZDiskImageStorageDeviceAttachment) SynchronizationMode() VZDiskImageSynchronizationMode {
	rv := objc.Send[VZDiskImageSynchronizationMode](d.ID, objc.Sel("synchronizationMode"))
	return VZDiskImageSynchronizationMode(rv)
}


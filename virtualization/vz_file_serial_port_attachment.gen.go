// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZFileSerialPortAttachment] class.
var (
	_VZFileSerialPortAttachmentClass     VZFileSerialPortAttachmentClass
	_VZFileSerialPortAttachmentClassOnce sync.Once
)

func getVZFileSerialPortAttachmentClass() VZFileSerialPortAttachmentClass {
	_VZFileSerialPortAttachmentClassOnce.Do(func() {
		_VZFileSerialPortAttachmentClass = VZFileSerialPortAttachmentClass{class: objc.GetClass("VZFileSerialPortAttachment")}
	})
	return _VZFileSerialPortAttachmentClass
}

// GetVZFileSerialPortAttachmentClass returns the class object for VZFileSerialPortAttachment.
func GetVZFileSerialPortAttachmentClass() VZFileSerialPortAttachmentClass {
	return getVZFileSerialPortAttachmentClass()
}

type VZFileSerialPortAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZFileSerialPortAttachmentClass) Alloc() VZFileSerialPortAttachment {
	rv := objc.Send[VZFileSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An attachment point that writes data from the guest system to a file.
//
// # Overview
// 
// Use a [VZFileSerialPortAttachment] object to configure a one-way serial
// port from the guest operating system to the virtual machine. When the guest
// sends data to the serial port, the virtual machine writes that data to the
// specified file. You can’t use this serial port to send data back to the
// guest.
// 
// Create a [VZSerialPortAttachment] object and assign it to an appropriate
// subclass of [VZSerialPortConfiguration] object, such as
// [VZVirtioConsoleDeviceConfiguration]. The file you use to create this
// object must be writable.
//
// # Creating the attachment point
//
//   - [VZFileSerialPortAttachment.InitWithURLAppendError]: Creates a file-based serial port attachment object.
//
// # Getting the file details
//
//   - [VZFileSerialPortAttachment.URL]: The URL of a file on the local file system.
//   - [VZFileSerialPortAttachment.Append]: A Boolean that indicates whether the virtual machine appends data to the file.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileSerialPortAttachment
type VZFileSerialPortAttachment struct {
	VZSerialPortAttachment
}

// VZFileSerialPortAttachmentFromID constructs a [VZFileSerialPortAttachment] from an objc.ID.
//
// An attachment point that writes data from the guest system to a file.
func VZFileSerialPortAttachmentFromID(id objc.ID) VZFileSerialPortAttachment {
	return VZFileSerialPortAttachment{VZSerialPortAttachment: VZSerialPortAttachmentFromID(id)}
}
// NOTE: VZFileSerialPortAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZFileSerialPortAttachment] class.
//
// # Creating the attachment point
//
//   - [IVZFileSerialPortAttachment.InitWithURLAppendError]: Creates a file-based serial port attachment object.
//
// # Getting the file details
//
//   - [IVZFileSerialPortAttachment.URL]: The URL of a file on the local file system.
//   - [IVZFileSerialPortAttachment.Append]: A Boolean that indicates whether the virtual machine appends data to the file.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileSerialPortAttachment
type IVZFileSerialPortAttachment interface {
	IVZSerialPortAttachment

	// Topic: Creating the attachment point

	// Creates a file-based serial port attachment object.
	InitWithURLAppendError(url foundation.INSURL, shouldAppend bool) (VZFileSerialPortAttachment, error)

	// Topic: Getting the file details

	// The URL of a file on the local file system.
	URL() foundation.INSURL
	// A Boolean that indicates whether the virtual machine appends data to the file.
	Append() bool
}

// Init initializes the instance.
func (f VZFileSerialPortAttachment) Init() VZFileSerialPortAttachment {
	rv := objc.Send[VZFileSerialPortAttachment](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f VZFileSerialPortAttachment) Autorelease() VZFileSerialPortAttachment {
	rv := objc.Send[VZFileSerialPortAttachment](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFileSerialPortAttachment creates a new VZFileSerialPortAttachment instance.
func NewVZFileSerialPortAttachment() VZFileSerialPortAttachment {
	class := getVZFileSerialPortAttachmentClass()
	rv := objc.Send[VZFileSerialPortAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a file-based serial port attachment object.
//
// url: The URL of a file on the local file system. The specified file must be
// writable by the virtual machine.
//
// shouldAppend: A Boolean that indicates whether the virtual machine opens the file in
// append mode. Specify [true] to append data to the file, and specify [false]
// to replace the contents of the file with any new data.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A file-based serial port attachment on success, or `nil` if initialization
// failed.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileSerialPortAttachment/init(url:append:)
func NewFileSerialPortAttachmentWithURLAppendError(url foundation.INSURL, shouldAppend bool) (VZFileSerialPortAttachment, error) {
	var errorPtr objc.ID
	instance := getVZFileSerialPortAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:append:error:"), url, shouldAppend, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZFileSerialPortAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZFileSerialPortAttachmentFromID(rv), nil
}

// Creates a file-based serial port attachment object.
//
// url: The URL of a file on the local file system. The specified file must be
// writable by the virtual machine.
//
// shouldAppend: A Boolean that indicates whether the virtual machine opens the file in
// append mode. Specify [true] to append data to the file, and specify [false]
// to replace the contents of the file with any new data.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A file-based serial port attachment on success, or `nil` if initialization
// failed.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileSerialPortAttachment/init(url:append:)
func (f VZFileSerialPortAttachment) InitWithURLAppendError(url foundation.INSURL, shouldAppend bool) (VZFileSerialPortAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("initWithURL:append:error:"), url, shouldAppend, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZFileSerialPortAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZFileSerialPortAttachmentFromID(rv), nil

}

// The URL of a file on the local file system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileSerialPortAttachment/url
func (f VZFileSerialPortAttachment) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// A Boolean that indicates whether the virtual machine appends data to the
// file.
//
// # Discussion
// 
// When the value of this property is [true], the virtual machine appends new
// data to the file; otherwise, it replaces the existing contents of the file
// before writing new data to it.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileSerialPortAttachment/append
func (f VZFileSerialPortAttachment) Append() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("append"))
	return rv
}


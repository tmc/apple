// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZFileHandleNetworkDeviceAttachment] class.
var (
	_VZFileHandleNetworkDeviceAttachmentClass     VZFileHandleNetworkDeviceAttachmentClass
	_VZFileHandleNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZFileHandleNetworkDeviceAttachmentClass() VZFileHandleNetworkDeviceAttachmentClass {
	_VZFileHandleNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZFileHandleNetworkDeviceAttachmentClass = VZFileHandleNetworkDeviceAttachmentClass{class: objc.GetClass("VZFileHandleNetworkDeviceAttachment")}
	})
	return _VZFileHandleNetworkDeviceAttachmentClass
}

// GetVZFileHandleNetworkDeviceAttachmentClass returns the class object for VZFileHandleNetworkDeviceAttachment.
func GetVZFileHandleNetworkDeviceAttachmentClass() VZFileHandleNetworkDeviceAttachmentClass {
	return getVZFileHandleNetworkDeviceAttachmentClass()
}

type VZFileHandleNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZFileHandleNetworkDeviceAttachmentClass) Alloc() VZFileHandleNetworkDeviceAttachment {
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A network device that transmits raw network packets and frames using a
// datagram socket.
//
// # Overview
// 
// A [VZFileHandleNetworkDeviceAttachment] object maps a network interface to
// a connected datagram socket. This attachment transmits data at the data
// link layer. You configure and manage the socket in your app, and manage the
// corresponding data transfers.
// 
// To configure a network device with a socket-based file handle:
// 
// - Create a socket with the `SOCK_DGRAM` type in your app. - Create a
// [VZFileHandleNetworkDeviceAttachment.FileHandle] from the socket’s file descriptor. - Create the
// [VZFileHandleNetworkDeviceAttachment] object using the file handle. -
// Assign the attachment object to the [VZFileHandleNetworkDeviceAttachment.Attachment] property of a
// [VZVirtioNetworkDeviceConfiguration] object. - Add the
// [VZVirtioNetworkDeviceConfiguration] object to the [VZFileHandleNetworkDeviceAttachment.NetworkDevices]
// property of your [VZVirtualMachineConfiguration].
// 
// This attachment doesn’t require your app to have the
// [com.apple.vm.networking] entitlement.
//
// [VZFileHandleNetworkDeviceAttachment.FileHandle]: https://developer.apple.com/documentation/Foundation/FileHandle
// [com.apple.vm.networking]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.vm.networking
//
// # Creating the attachment point
//
//   - [VZFileHandleNetworkDeviceAttachment.InitWithFileHandle]: Creates the attachment from a file handle that contains a connected datagram socket.
//
// # Getting the file handle
//
//   - [VZFileHandleNetworkDeviceAttachment.FileHandle]: The file handle assigned to this attachment.
//
// # Specifying the network packet size
//
//   - [VZFileHandleNetworkDeviceAttachment.MaximumTransmissionUnit]: An integer value that indicates the maximum transmission unit (MTU) associated with this attachment.
//   - [VZFileHandleNetworkDeviceAttachment.SetMaximumTransmissionUnit]
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment
type VZFileHandleNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZFileHandleNetworkDeviceAttachmentFromID constructs a [VZFileHandleNetworkDeviceAttachment] from an objc.ID.
//
// A network device that transmits raw network packets and frames using a
// datagram socket.
func VZFileHandleNetworkDeviceAttachmentFromID(id objc.ID) VZFileHandleNetworkDeviceAttachment {
	return VZFileHandleNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}
// NOTE: VZFileHandleNetworkDeviceAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZFileHandleNetworkDeviceAttachment] class.
//
// # Creating the attachment point
//
//   - [IVZFileHandleNetworkDeviceAttachment.InitWithFileHandle]: Creates the attachment from a file handle that contains a connected datagram socket.
//
// # Getting the file handle
//
//   - [IVZFileHandleNetworkDeviceAttachment.FileHandle]: The file handle assigned to this attachment.
//
// # Specifying the network packet size
//
//   - [IVZFileHandleNetworkDeviceAttachment.MaximumTransmissionUnit]: An integer value that indicates the maximum transmission unit (MTU) associated with this attachment.
//   - [IVZFileHandleNetworkDeviceAttachment.SetMaximumTransmissionUnit]
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment
type IVZFileHandleNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	// Topic: Creating the attachment point

	// Creates the attachment from a file handle that contains a connected datagram socket.
	InitWithFileHandle(fileHandle foundation.NSFileHandle) VZFileHandleNetworkDeviceAttachment

	// Topic: Getting the file handle

	// The file handle assigned to this attachment.
	FileHandle() foundation.NSFileHandle

	// Topic: Specifying the network packet size

	// An integer value that indicates the maximum transmission unit (MTU) associated with this attachment.
	MaximumTransmissionUnit() int
	SetMaximumTransmissionUnit(value int)

	// The object that defines how the virtual network device communicates with the host system.
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	// The array of network devices that you expose to the guest operating system.
	NetworkDevices() IVZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
}





// Init initializes the instance.
func (f VZFileHandleNetworkDeviceAttachment) Init() VZFileHandleNetworkDeviceAttachment {
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f VZFileHandleNetworkDeviceAttachment) Autorelease() VZFileHandleNetworkDeviceAttachment {
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFileHandleNetworkDeviceAttachment creates a new VZFileHandleNetworkDeviceAttachment instance.
func NewVZFileHandleNetworkDeviceAttachment() VZFileHandleNetworkDeviceAttachment {
	class := getVZFileHandleNetworkDeviceAttachmentClass()
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates the attachment from a file handle that contains a connected
// datagram socket.
//
// fileHandle: A file handle to a connected datagram socket.
//
// # Return Value
// 
// An attachment object for the specified file handle.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/init(fileHandle:)
func NewFileHandleNetworkDeviceAttachmentWithFileHandle(fileHandle foundation.NSFileHandle) VZFileHandleNetworkDeviceAttachment {
	instance := getVZFileHandleNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileHandle:"), fileHandle)
	return VZFileHandleNetworkDeviceAttachmentFromID(rv)
}







// Creates the attachment from a file handle that contains a connected
// datagram socket.
//
// fileHandle: A file handle to a connected datagram socket.
//
// # Return Value
// 
// An attachment object for the specified file handle.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/init(fileHandle:)
func (f VZFileHandleNetworkDeviceAttachment) InitWithFileHandle(fileHandle foundation.NSFileHandle) VZFileHandleNetworkDeviceAttachment {
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](f.ID, objc.Sel("initWithFileHandle:"), fileHandle)
	return rv
}












// The file handle assigned to this attachment.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/fileHandle
func (f VZFileHandleNetworkDeviceAttachment) FileHandle() foundation.NSFileHandle {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fileHandle"))
	return foundation.NSFileHandleFromID(objc.ID(rv))
}



// An integer value that indicates the maximum transmission unit (MTU)
// associated with this attachment.
//
// # Discussion
// 
// The client side of the associated datagram socket must be properly
// configured with the appropriate values for `SO_SNDBUF`, and `SO_RCVBUF`.
// Set these using the `setsockopt(_:_:_:_:_:)` system call. The system
// expects the value of `SO_RCVBUF` to be at least double the value of
// `SO_SNDBUF`, and for optimal performance, the recommended value of
// `SO_RCVBUF` is four times the value of `SO_SNDBUF`.
// 
// The default MTU is 1500. The maximum MTU allowed is 65535, and the minimum
// MTU allowed is 1500. An invalid MTU value results in an invalid virtual
// machine configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/maximumTransmissionUnit
func (f VZFileHandleNetworkDeviceAttachment) MaximumTransmissionUnit() int {
	rv := objc.Send[int](f.ID, objc.Sel("maximumTransmissionUnit"))
	return rv
}
func (f VZFileHandleNetworkDeviceAttachment) SetMaximumTransmissionUnit(value int) {
	objc.Send[struct{}](f.ID, objc.Sel("setMaximumTransmissionUnit:"), value)
}



// The object that defines how the virtual network device communicates with
// the host system.
//
// See: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (f VZFileHandleNetworkDeviceAttachment) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("attachment"))
	return VZNetworkDeviceAttachmentFromID(objc.ID(rv))
}
func (f VZFileHandleNetworkDeviceAttachment) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[struct{}](f.ID, objc.Sel("setAttachment:"), value)
}



// The array of network devices that you expose to the guest operating system.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (f VZFileHandleNetworkDeviceAttachment) NetworkDevices() IVZNetworkDeviceConfiguration {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("networkDevices"))
	return VZNetworkDeviceConfigurationFromID(objc.ID(rv))
}
func (f VZFileHandleNetworkDeviceAttachment) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[struct{}](f.ID, objc.Sel("setNetworkDevices:"), value)
}
























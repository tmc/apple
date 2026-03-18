// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZNetworkBlockDeviceStorageDeviceAttachment] class.
var (
	_VZNetworkBlockDeviceStorageDeviceAttachmentClass     VZNetworkBlockDeviceStorageDeviceAttachmentClass
	_VZNetworkBlockDeviceStorageDeviceAttachmentClassOnce sync.Once
)

func getVZNetworkBlockDeviceStorageDeviceAttachmentClass() VZNetworkBlockDeviceStorageDeviceAttachmentClass {
	_VZNetworkBlockDeviceStorageDeviceAttachmentClassOnce.Do(func() {
		_VZNetworkBlockDeviceStorageDeviceAttachmentClass = VZNetworkBlockDeviceStorageDeviceAttachmentClass{class: objc.GetClass("VZNetworkBlockDeviceStorageDeviceAttachment")}
	})
	return _VZNetworkBlockDeviceStorageDeviceAttachmentClass
}

// GetVZNetworkBlockDeviceStorageDeviceAttachmentClass returns the class object for VZNetworkBlockDeviceStorageDeviceAttachment.
func GetVZNetworkBlockDeviceStorageDeviceAttachmentClass() VZNetworkBlockDeviceStorageDeviceAttachmentClass {
	return getVZNetworkBlockDeviceStorageDeviceAttachmentClass()
}

type VZNetworkBlockDeviceStorageDeviceAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNetworkBlockDeviceStorageDeviceAttachmentClass) Alloc() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A storage device attachment backed by a Network Block Device (NBD) client.
//
// # Overview
// 
// This storage device attachment provides a Network Block Device (NBD) client
// implementation. The NBD client connects to an NBD server referred to by an
// NBD Uniform Resource Indicator (URI), represented as an URL in this API.
// The NBD server runs outside of and isn’t controlled by the Virtualization
// framework. The NBD client forwards the guest’s I/O operations to the NBD
// server, which handles the I/O operations.
// 
// The NBD client attempts to connect to the NBD server referred to by the URL
// used when you started the VM with [StartWithCompletionHandler]. However,
// it’s important to note that a connection attempt isn’t made when the
// framework initializes the attachment object.
// 
// Reconnection attempts take place throughout the life cycle of the VM when
// the NBD client encounters a recoverable error such as connection timeout
// and unexpected connection errors. The NBD client disconnects from the
// server when the VM shuts down.
// 
// Using this attachment requires the app to have the
// [com.apple.security.network.client] entitlement because this attachment
// opens an outgoing network connection.
// 
// To create a device that uses an NBD service, first initialize a
// [VZNetworkBlockDeviceStorageDeviceAttachment] with the URI of an NBD
// server, then use the attachment to configure a
// [VZStorageDeviceConfiguration] as shown in the example below (the
// attachment works with any subclass of [VZStorageDeviceConfiguration], not
// just [VZVirtioBlockDeviceConfiguration]):
// 
// For more information about Network Block Devices, see the [Network Block
// Device Specification] on GitHub.
// 
// For more information about the NBD URL format, see the [Network Block
// Device URL specification] on GitHub.
//
// [Network Block Device Specification]: https://github.com/NetworkBlockDevice/nbd/blob/master/doc/proto.md
// [Network Block Device URL specification]: https://github.com/NetworkBlockDevice/nbd/blob/master/doc/uri.md
// [com.apple.security.network.client]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.security.network.client
//
// # Creating network block device attachments
//
//   - [VZNetworkBlockDeviceStorageDeviceAttachment.InitWithURLError]: Creates a new network block device (NBD) storage attachment from an NDB Uniform Resource Indicator (URI) represented as a URL that you provide.
//   - [VZNetworkBlockDeviceStorageDeviceAttachment.InitWithURLTimeoutForcedReadOnlySynchronizationModeError]: Creates a new network block device storage attachment from an NBD Uniform Resource Indicator (URI) represented as a URL, timeout value, and read-only and synchronization modes that you provide.
//
// # Getting information about the attachment point
//
//   - [VZNetworkBlockDeviceStorageDeviceAttachment.ForcedReadOnly]: Returns a Boolean value that indicates whether the underlying disk attachment network is in a read-only state.
//   - [VZNetworkBlockDeviceStorageDeviceAttachment.SynchronizationMode]: The mode in which the NBD client synchronizes data with the NBD server.
//   - [VZNetworkBlockDeviceStorageDeviceAttachment.Timeout]: The timeout value in seconds for the connection between the client and server.
//   - [VZNetworkBlockDeviceStorageDeviceAttachment.URL]: The URL that refers to the NBD server to which the NBD client will connect.
//
// # Observing changes to the network block device
//
//   - [VZNetworkBlockDeviceStorageDeviceAttachment.Delegate]: The object that receives messages about changes to the network block device attachment.
//   - [VZNetworkBlockDeviceStorageDeviceAttachment.SetDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment
type VZNetworkBlockDeviceStorageDeviceAttachment struct {
	VZStorageDeviceAttachment
}

// VZNetworkBlockDeviceStorageDeviceAttachmentFromID constructs a [VZNetworkBlockDeviceStorageDeviceAttachment] from an objc.ID.
//
// A storage device attachment backed by a Network Block Device (NBD) client.
func VZNetworkBlockDeviceStorageDeviceAttachmentFromID(id objc.ID) VZNetworkBlockDeviceStorageDeviceAttachment {
	return VZNetworkBlockDeviceStorageDeviceAttachment{VZStorageDeviceAttachment: VZStorageDeviceAttachmentFromID(id)}
}
// NOTE: VZNetworkBlockDeviceStorageDeviceAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZNetworkBlockDeviceStorageDeviceAttachment] class.
//
// # Creating network block device attachments
//
//   - [IVZNetworkBlockDeviceStorageDeviceAttachment.InitWithURLError]: Creates a new network block device (NBD) storage attachment from an NDB Uniform Resource Indicator (URI) represented as a URL that you provide.
//   - [IVZNetworkBlockDeviceStorageDeviceAttachment.InitWithURLTimeoutForcedReadOnlySynchronizationModeError]: Creates a new network block device storage attachment from an NBD Uniform Resource Indicator (URI) represented as a URL, timeout value, and read-only and synchronization modes that you provide.
//
// # Getting information about the attachment point
//
//   - [IVZNetworkBlockDeviceStorageDeviceAttachment.ForcedReadOnly]: Returns a Boolean value that indicates whether the underlying disk attachment network is in a read-only state.
//   - [IVZNetworkBlockDeviceStorageDeviceAttachment.SynchronizationMode]: The mode in which the NBD client synchronizes data with the NBD server.
//   - [IVZNetworkBlockDeviceStorageDeviceAttachment.Timeout]: The timeout value in seconds for the connection between the client and server.
//   - [IVZNetworkBlockDeviceStorageDeviceAttachment.URL]: The URL that refers to the NBD server to which the NBD client will connect.
//
// # Observing changes to the network block device
//
//   - [IVZNetworkBlockDeviceStorageDeviceAttachment.Delegate]: The object that receives messages about changes to the network block device attachment.
//   - [IVZNetworkBlockDeviceStorageDeviceAttachment.SetDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment
type IVZNetworkBlockDeviceStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment

	// Topic: Creating network block device attachments

	// Creates a new network block device (NBD) storage attachment from an NDB Uniform Resource Indicator (URI) represented as a URL that you provide.
	InitWithURLError(URL foundation.INSURL) (VZNetworkBlockDeviceStorageDeviceAttachment, error)
	// Creates a new network block device storage attachment from an NBD Uniform Resource Indicator (URI) represented as a URL, timeout value, and read-only and synchronization modes that you provide.
	InitWithURLTimeoutForcedReadOnlySynchronizationModeError(URL foundation.INSURL, timeout float64, forcedReadOnly bool, synchronizationMode VZDiskSynchronizationMode) (VZNetworkBlockDeviceStorageDeviceAttachment, error)

	// Topic: Getting information about the attachment point

	// Returns a Boolean value that indicates whether the underlying disk attachment network is in a read-only state.
	ForcedReadOnly() bool
	// The mode in which the NBD client synchronizes data with the NBD server.
	SynchronizationMode() VZDiskSynchronizationMode
	// The timeout value in seconds for the connection between the client and server.
	Timeout() float64
	// The URL that refers to the NBD server to which the NBD client will connect.
	URL() foundation.INSURL

	// Topic: Observing changes to the network block device

	// The object that receives messages about changes to the network block device attachment.
	Delegate() VZNetworkBlockDeviceStorageDeviceAttachmentDelegate
	SetDelegate(value VZNetworkBlockDeviceStorageDeviceAttachmentDelegate)
}

// Init initializes the instance.
func (n VZNetworkBlockDeviceStorageDeviceAttachment) Init() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n VZNetworkBlockDeviceStorageDeviceAttachment) Autorelease() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkBlockDeviceStorageDeviceAttachment creates a new VZNetworkBlockDeviceStorageDeviceAttachment instance.
func NewVZNetworkBlockDeviceStorageDeviceAttachment() VZNetworkBlockDeviceStorageDeviceAttachment {
	class := getVZNetworkBlockDeviceStorageDeviceAttachmentClass()
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new network block device (NBD) storage attachment from an NDB
// Uniform Resource Indicator (URI) represented as a URL that you provide.
//
// URL: The NBD’s URI represented as a URL.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/init(url:)
func NewNetworkBlockDeviceStorageDeviceAttachmentWithURLError(URL foundation.INSURL) (VZNetworkBlockDeviceStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	instance := getVZNetworkBlockDeviceStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), URL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZNetworkBlockDeviceStorageDeviceAttachmentFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return VZNetworkBlockDeviceStorageDeviceAttachmentFromID(rv), nil
}

// Creates a new network block device storage attachment from an NBD Uniform
// Resource Indicator (URI) represented as a URL, timeout value, and read-only
// and synchronization modes that you provide.
//
// URL: The NBD’s URI represented as a URL.
//
// timeout: The timeout value in seconds for the connection between the client and
// server. When the timeout expires, an attempt to reconnect with the server
// takes place.
//
// forcedReadOnly: If [true], the framework forces the disk attachment to be read-only,
// regardless of whether or not the NBD server supports write requests.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// synchronizationMode: The mode in which the disk attachment synchronizes data with the underlying
// storage device.
//
// # Discussion
// 
// The `forcedReadOnly` parameter affects how framework exposes the NBD client
// to the guest operating system by the storage controller. As part of the NBD
// protocol, the NBD server advertises whether or not the disk exposed by the
// NBD client is read-only during the handshake phase of the protocol. Setting
// `forcedReadOnly` to [true] forces the NBD client to show up as read-only to
// the guest regardless of whether or not the NBD server advertises itself as
// read-only.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/init(url:timeout:isForcedReadOnly:synchronizationMode:)
func NewNetworkBlockDeviceStorageDeviceAttachmentWithURLTimeoutForcedReadOnlySynchronizationModeError(URL foundation.INSURL, timeout float64, forcedReadOnly bool, synchronizationMode VZDiskSynchronizationMode) (VZNetworkBlockDeviceStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	instance := getVZNetworkBlockDeviceStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:timeout:forcedReadOnly:synchronizationMode:error:"), URL, timeout, forcedReadOnly, synchronizationMode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZNetworkBlockDeviceStorageDeviceAttachmentFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return VZNetworkBlockDeviceStorageDeviceAttachmentFromID(rv), nil
}

// Creates a new network block device (NBD) storage attachment from an NDB
// Uniform Resource Indicator (URI) represented as a URL that you provide.
//
// URL: The NBD’s URI represented as a URL.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/init(url:)
func (n VZNetworkBlockDeviceStorageDeviceAttachment) InitWithURLError(URL foundation.INSURL) (VZNetworkBlockDeviceStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("initWithURL:error:"), URL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZNetworkBlockDeviceStorageDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZNetworkBlockDeviceStorageDeviceAttachmentFromID(rv), nil

}

// Creates a new network block device storage attachment from an NBD Uniform
// Resource Indicator (URI) represented as a URL, timeout value, and read-only
// and synchronization modes that you provide.
//
// URL: The NBD’s URI represented as a URL.
//
// timeout: The timeout value in seconds for the connection between the client and
// server. When the timeout expires, an attempt to reconnect with the server
// takes place.
//
// forcedReadOnly: If [true], the framework forces the disk attachment to be read-only,
// regardless of whether or not the NBD server supports write requests.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// synchronizationMode: The mode in which the disk attachment synchronizes data with the underlying
// storage device.
//
// # Discussion
// 
// The `forcedReadOnly` parameter affects how framework exposes the NBD client
// to the guest operating system by the storage controller. As part of the NBD
// protocol, the NBD server advertises whether or not the disk exposed by the
// NBD client is read-only during the handshake phase of the protocol. Setting
// `forcedReadOnly` to [true] forces the NBD client to show up as read-only to
// the guest regardless of whether or not the NBD server advertises itself as
// read-only.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/init(url:timeout:isForcedReadOnly:synchronizationMode:)
func (n VZNetworkBlockDeviceStorageDeviceAttachment) InitWithURLTimeoutForcedReadOnlySynchronizationModeError(URL foundation.INSURL, timeout float64, forcedReadOnly bool, synchronizationMode VZDiskSynchronizationMode) (VZNetworkBlockDeviceStorageDeviceAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("initWithURL:timeout:forcedReadOnly:synchronizationMode:error:"), URL, timeout, forcedReadOnly, synchronizationMode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZNetworkBlockDeviceStorageDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZNetworkBlockDeviceStorageDeviceAttachmentFromID(rv), nil

}

// Checks if the URL is a valid network block device URL.
//
// URL: The NBD URL to validate.
//
// # Discussion
// 
// This method checks that the URL is well-formed; however, it doesn’t
// attempt to access the URL. See the [NBD URL specification] on GitHub for
// more detailed descriptions of valid URIs.
//
// [NBD URL specification]: https://github.com/NetworkBlockDevice/nbd/blob/master/doc/uri.md
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/validate(_:)
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) ValidateURLError(URL foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("validateURL:error:"), URL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns a Boolean value that indicates whether the underlying disk
// attachment network is in a read-only state.
//
// # Discussion
// 
// The `forcedReadOnly` parameter affects how the Virtualization framework
// exposes the network block device (NBD) client to the guest operating system
// by the storage controller.
// 
// As part of the NBD protocol, during the handshake phase, the server
// advertises whether or not the disk the server exposes is read-only. Setting
// `forcedReadOnly` to [true] forces the NBD client to show up as read-only to
// the guest regardless of whether or not the NBD server advertises itself as
// read-only.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/isForcedReadOnly
func (n VZNetworkBlockDeviceStorageDeviceAttachment) ForcedReadOnly() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isForcedReadOnly"))
	return rv
}

// The mode in which the NBD client synchronizes data with the NBD server.
//
// # Discussion
// 
// See [VZDiskSynchronizationMode] for details on how the specific mode
// affects data synchronization between the NBD client and server.
//
// [VZDiskSynchronizationMode]: https://developer.apple.com/documentation/Virtualization/VZDiskSynchronizationMode
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/synchronizationMode
func (n VZNetworkBlockDeviceStorageDeviceAttachment) SynchronizationMode() VZDiskSynchronizationMode {
	rv := objc.Send[VZDiskSynchronizationMode](n.ID, objc.Sel("synchronizationMode"))
	return VZDiskSynchronizationMode(rv)
}

// The timeout value in seconds for the connection between the client and
// server.
//
// # Discussion
// 
// When the timeout expires, the client attempts to reconnect with the server.
// If after several retries, the client can’t reestablish a connection to
// the server, the framework invokes the [AttachmentDidEncounterError]
// delegate method.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/timeout
func (n VZNetworkBlockDeviceStorageDeviceAttachment) Timeout() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("timeout"))
	return rv
}

// The URL that refers to the NBD server to which the NBD client will connect.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/url
func (n VZNetworkBlockDeviceStorageDeviceAttachment) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The object that receives messages about changes to the network block device
// attachment.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/delegate
func (n VZNetworkBlockDeviceStorageDeviceAttachment) Delegate() VZNetworkBlockDeviceStorageDeviceAttachmentDelegate {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("delegate"))
	return VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObjectFromID(rv)
}
func (n VZNetworkBlockDeviceStorageDeviceAttachment) SetDelegate(value VZNetworkBlockDeviceStorageDeviceAttachmentDelegate) {
	objc.Send[struct{}](n.ID, objc.Sel("setDelegate:"), value)
}

